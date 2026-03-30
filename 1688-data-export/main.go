package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/spf13/cobra"
)

type CRMData struct {
	Timestamp          string                  `json:"timestamp"`
	Platform           string                  `json:"platform"`
	MarketingCRM       *MarketingCRMData       `json:"marketingCRM"`
	CustomerManagement *CustomerManagementData `json:"customerManagement"`
}

type MarketingCRMData struct {
	PotentialCustomers [][]string `json:"potentialCustomers"`
	PotentialPublic    [][]string `json:"potentialPublic"`
	OldCustomersFirst  [][]string `json:"oldCustomersFirst"`
	OldCustomersSecond [][]string `json:"oldCustomersSecond"`
}

type CustomerManagementData struct {
	URL    string  `json:"url"`
	Tabs   []Tab   `json:"tabs"`
	Tables []Table `json:"tables"`
}

type Tab struct {
	Index    int    `json:"index"`
	Text     string `json:"text"`
	IsActive bool   `json:"isActive"`
}

type Table struct {
	TableIndex int        `json:"tableIndex"`
	RowCount   int        `json:"rowCount"`
	Data       [][]string `json:"data"`
}

var (
	outputDir string
	headless  bool
	timeout   int
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "crm-export",
		Short: "1688 CRM数据导出工具",
		Long:  "从1688 CRM系统导出客户数据到Excel格式文件",
		Run:   runExport,
	}

	rootCmd.Flags().StringVarP(&outputDir, "output", "o", ".", "输出目录")
	rootCmd.Flags().BoolVarP(&headless, "headless", "H", false, "无头模式运行")
	rootCmd.Flags().IntVarP(&timeout, "timeout", "t", 300, "登录超时时间(秒)")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runExport(cmd *cobra.Command, args []string) {
	fmt.Println("====================================")
	fmt.Println("1688 CRM数据导出工具 (Go版)")
	fmt.Printf("操作系统: %s\n", runtime.GOOS)
	fmt.Println("====================================\n")

	baseDir, _ := os.Getwd()
	userDataDir := filepath.Join(baseDir, "browser_data")
	os.MkdirAll(userDataDir, 0755)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", headless),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("start-maximized", true),
		chromedp.WindowSize(1920, 1080),
		chromedp.UserDataDir(userDataDir),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	silentLogger := func(string, ...interface{}) {}
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithErrorf(silentLogger))
	defer cancel()

	data := &CRMData{
		Timestamp: time.Now().Format(time.RFC3339),
		Platform:  runtime.GOOS,
	}

	if err := loginAndWait(ctx); err != nil {
		log.Printf("登录失败: %v", err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	var marketingErr, customerErr error
	var marketingData *MarketingCRMData
	var customerData *CustomerManagementData

	ctx1, cancel1 := chromedp.NewContext(allocCtx, chromedp.WithErrorf(silentLogger))
	defer cancel1()

	ctx2, cancel2 := chromedp.NewContext(allocCtx, chromedp.WithErrorf(silentLogger))
	defer cancel2()

	go func() {
		defer wg.Done()
		marketingData, marketingErr = exportMarketingCRM(ctx1)
		if marketingErr != nil {
			log.Printf("营销平台CRM导出失败: %v", marketingErr)
		}
	}()

	go func() {
		defer wg.Done()
		customerData, customerErr = exportCustomerManagement(ctx2)
		if customerErr != nil {
			log.Printf("客户管理系统导出失败: %v", customerErr)
		}
	}()

	wg.Wait()

	data.MarketingCRM = marketingData
	data.CustomerManagement = customerData

	saveData(data, outputDir)
}

func loginAndWait(ctx context.Context) error {
	fmt.Println("\n【登录检测】")

	var currentURL string
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://p4p.1688.com/main.html"),
		chromedp.Sleep(3*time.Second),
		chromedp.Location(&currentURL),
	)
	if err != nil {
		return err
	}

	if strings.Contains(currentURL, "login") {
		fmt.Println("需要登录，请扫码...")
		fmt.Println("请在浏览器中完成登录...")

		loginCtx, loginCancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
		defer loginCancel()

		for {
			select {
			case <-loginCtx.Done():
				return fmt.Errorf("登录超时")
			default:
				time.Sleep(3 * time.Second)
				var checkURL string
				chromedp.Run(ctx, chromedp.Location(&checkURL))
				if !strings.Contains(checkURL, "login") {
					fmt.Println("登录成功！")
					return nil
				}
			}
		}
	}

	fmt.Println("已登录，继续执行...")
	return nil
}

func exportMarketingCRM(ctx context.Context) (*MarketingCRMData, error) {
	fmt.Println("\n【数据源一】营销平台CRM (并发)")

	result := &MarketingCRMData{}

	fmt.Println("\n1. 潜客机会")
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://p4p.1688.com/main.html#!/crmbuyer?tab=potential&potentialType=potential_real"),
		chromedp.Sleep(3*time.Second),
	)
	if err != nil {
		return nil, err
	}

	result.PotentialCustomers = extractTables(ctx)
	fmt.Printf("  找到 %d 行数据\n", len(result.PotentialCustomers))

	fmt.Println("\n2. 潜客公海")
	err = chromedp.Run(ctx,
		chromedp.Navigate("https://p4p.1688.com/main.html#!/crmbuyer?tab=potential&potentialType=potential_no_real"),
		chromedp.Sleep(3*time.Second),
	)
	if err != nil {
		return nil, err
	}
	result.PotentialPublic = extractTables(ctx)
	fmt.Printf("  找到 %d 行数据\n", len(result.PotentialPublic))

	fmt.Println("\n3. 老客管理")
	err = chromedp.Run(ctx,
		chromedp.Navigate("https://p4p.1688.com/main.html#!/crmbuyer?tab=old"),
		chromedp.Sleep(3*time.Second),
	)
	if err != nil {
		return nil, err
	}

	fmt.Println("\n3.1 老客管理 - 第一次获取数据")
	result.OldCustomersFirst = extractTables(ctx)
	fmt.Printf("  找到 %d 行数据\n", len(result.OldCustomersFirst))

	fmt.Println("\n3.2 老客管理 - 点击重置按钮")

	var resetButtonFound bool
	err = chromedp.Run(ctx,
		chromedp.Evaluate(`
			(() => {
				const resetBtn = document.querySelector('#app > div > div.bp-new-layout-wrapper > div.bp-new-layout-right > div.bp-new-layout-main.bp-zhizhou-layout-main > div > div.custom_tab_content > div > div.next-tabs-content > div > div.poten_range button');
				if (resetBtn) {
					resetBtn.click();
					return true;
				}
				return false;
			})()
		`, &resetButtonFound),
	)
	if err != nil {
		log.Printf("重置按钮操作失败: %v", err)
	}

	if resetButtonFound {
		fmt.Println("  ✓ 已点击重置按钮")
	} else {
		fmt.Println("  ⚠ 未找到重置按钮")
	}

	chromedp.Run(ctx, chromedp.Sleep(5*time.Second))

	result.OldCustomersSecond = extractTables(ctx)
	fmt.Printf("  找到 %d 行数据\n", len(result.OldCustomersSecond))

	fmt.Println("\n✓ 营销平台CRM数据获取完成")
	return result, nil
}

func exportCustomerManagement(ctx context.Context) (*CustomerManagementData, error) {
	fmt.Println("\n【数据源二】客户管理系统 (并发)")

	result := &CustomerManagementData{
		URL: "https://air.1688.com/app/CSBC-modules/csbc-page-member-crm/index.html",
	}

	fmt.Println("\n直接导航到CRM页面...")
	err := chromedp.Run(ctx,
		chromedp.Navigate(result.URL),
		chromedp.Sleep(5*time.Second),
	)
	if err != nil {
		return nil, err
	}
	fmt.Println("✓ 页面加载完成")

	var tabs []Tab
	err = chromedp.Run(ctx,
		chromedp.Evaluate(`
			(() => {
				const tabItems = document.querySelectorAll('.csbc-page-member-crm-tabs-item');
				return Array.from(tabItems).map((item, index) => ({
					index: index,
					text: item.innerText.trim(),
					isActive: item.className.includes('active')
				}));
			})()
		`, &tabs),
	)

	if err == nil && len(tabs) > 0 {
		fmt.Printf("标签: %s\n", strings.Join(func() []string {
			var names []string
			for _, t := range tabs {
				names = append(names, t.Text)
			}
			return names
		}(), ", "))
	}
	result.Tabs = tabs

	fmt.Println("\n点击\"全部客户明细\"标签...")
	var clicked bool
	err = chromedp.Run(ctx,
		chromedp.Evaluate(`
			(() => {
				let clicked = false;
				const tabItems = document.querySelectorAll('.csbc-page-member-crm-tabs-item');
				tabItems.forEach(item => {
					if (item.innerText.includes('全部客户明细')) {
						item.click();
						clicked = true;
					}
				});
				return clicked;
			})()
		`, &clicked),
	)

	if clicked {
		fmt.Println("  ✓ 成功点击\"全部客户明细\"标签")
	} else {
		fmt.Println("  ⚠ 未找到\"全部客户明细\"标签")
	}

	chromedp.Run(ctx, chromedp.Sleep(5*time.Second))

	tables := extractTables(ctx)
	if len(tables) > 0 {
		result.Tables = []Table{{
			TableIndex: 0,
			RowCount:   len(tables),
			Data:       tables,
		}}
	}
	fmt.Printf("提取到 %d 行数据\n", len(tables))

	fmt.Println("✓ 客户管理系统数据获取完成")
	return result, nil
}

func extractTables(ctx context.Context) [][]string {
	var result [][]string
	chromedp.Run(ctx,
		chromedp.Evaluate(`
			(() => {
				const tables = document.querySelectorAll('table');
				const results = [];
				tables.forEach(table => {
					const rows = table.querySelectorAll('tr');
					rows.forEach(row => {
						const cells = row.querySelectorAll('td, th');
						const rowData = Array.from(cells).map(cell => cell.innerText.trim());
						if (rowData.length > 0 && rowData.some(d => d !== '')) {
							results.push(rowData);
						}
					});
				});
				return results;
			})()
		`, &result),
	)
	return result
}

func saveData(data *CRMData, dir string) {
	fmt.Println("\n====================================")
	fmt.Println("保存数据文件")
	fmt.Println("====================================\n")

	os.MkdirAll(dir, 0755)

	timestamp := time.Now().Format("20060102150405")
	jsonFile := filepath.Join(dir, fmt.Sprintf("all_crm_data_%s.json", timestamp))
	csvFile := filepath.Join(dir, fmt.Sprintf("all_crm_data_%s.csv", timestamp))

	jsonData, _ := json.MarshalIndent(data, "", "  ")
	os.WriteFile(jsonFile, jsonData, 0644)
	fmt.Printf("✓ JSON数据已保存: %s\n", jsonFile)

	saveCSV(data, csvFile)
	fmt.Printf("✓ CSV数据已保存: %s\n", csvFile)

	fmt.Println("\n====================================")
	fmt.Println("数据导出完成！")
	fmt.Println("====================================")
}

func saveCSV(data *CRMData, filename string) {
	file, _ := os.Create(filename)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"========================================"})
	writer.Write([]string{"数据源一：营销平台CRM"})
	writer.Write([]string{"========================================"})
	writer.Write([]string{})

	writer.Write([]string{"潜客机会"})
	for _, row := range data.MarketingCRM.PotentialCustomers {
		writer.Write(row)
	}

	writer.Write([]string{})
	writer.Write([]string{"潜客公海"})
	for _, row := range data.MarketingCRM.PotentialPublic {
		writer.Write(row)
	}

	writer.Write([]string{})
	writer.Write([]string{"老客管理 - 第一次获取"})
	for _, row := range data.MarketingCRM.OldCustomersFirst {
		writer.Write(row)
	}

	writer.Write([]string{})
	writer.Write([]string{"老客管理 - 重置后获取"})
	for _, row := range data.MarketingCRM.OldCustomersSecond {
		writer.Write(row)
	}

	writer.Write([]string{})
	writer.Write([]string{"========================================"})
	writer.Write([]string{"数据源二：客户管理系统"})
	writer.Write([]string{"========================================"})
	writer.Write([]string{})

	writer.Write([]string{"全部客户明细"})
	for _, table := range data.CustomerManagement.Tables {
		for _, row := range table.Data {
			writer.Write(row)
		}
	}
}
