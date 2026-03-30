---
name: "1688-data-export"
description: "导出1688 CRM客户数据（潜客机会、老客管理、AI老客跟进、AI新客拓展、全部客户明细）到Excel格式文件。Invoke when user wants to export 1688 CRM data, customer data, or get customer management information."
---

# 1688数据导出工具

使用puppeteer或agent-browser工具打开1688客户管理系统，获取CRM客户数据，并导出为Excel兼容格式。

## 数据源概览

本工具支持从两个不同的1688页面获取客户数据：

### 数据源一：营销平台CRM
- **主入口**: `https://p4p.1688.com/main.html`
- **数据内容**:
  - 潜客机会（潜在客户询盘数据）
  - 潜客公海（潜在客户公海数据）
  - 老客管理（老客户采购数据）

### 数据源二：客户管理系统
- **URL**: `https://work.1688.com/home/page/index.htm?_path_=sellerPro/2017Sellerbase_Customer/mycrm`
- **数据内容**:
  - AI老客跟进
  - AI新客拓展
  - 全部客户明细

---

## 数据源一：营销平台CRM

### 页面URL
```
潜客机会: https://p4p.1688.com/main.html#!/crmbuyer?tab=potential&potentialType=potential_real
潜客公海: https://p4p.1688.com/main.html#!/crmbuyer?tab=potential&potentialType=potential_no_real
老客管理: https://p4p.1688.com/main.html#!/crmbuyer?tab=old
```

### 页面结构
页面是一个单页应用（SPA），使用hash路由，包含三个主要标签：
1. **潜客机会** - 显示潜在客户的询盘数据
2. **潜客公海** - 显示潜在客户公海数据
3. **老客管理** - 显示老客户的采购数据

### 请求链路
```
用户请求
    ↓
打开营销平台CRM页面（潜客机会）
    ↓
检测登录状态
    ├─ 未登录 → 等待用户扫码登录
    └─ 已登录 → 继续
    ↓
等待页面加载完成
    ↓
获取"潜客机会"数据
    ↓
访问"潜客公海"URL
    ↓
获取"潜客公海"数据
    ↓
访问"老客管理"URL
    ↓
第一次获取"老客管理"数据
    ↓
点击重置按钮
    ↓
第二次获取"老客管理"数据
    ↓
保存数据文件
    ↓
发送文件给用户
```

### 数据表结构

#### 潜客机会表格
**表头字段：**
- 买家昵称
- 买家行为
- 最近询盘时间
- 渠道
- 表单留言
- 联系电话
- 询盘记录
- 买家身份
- 买家等级
- 所在地
- 小记内容
- 操作

**表格选择器：**
```javascript
// 标签选择器
潜客机会标签: 'a[href="#!/crmbuyer"]' 或包含"潜客机会"文本的元素

// 表格选择器
table selector: 'table'
row selector: 'tr'
cell selector: 'td, th'
```

#### 老客管理表格
**表头字段：**
- 买家昵称
- 首单是否广告引导
- 买家身份
- 采购次数
- 累计采购金额（元）
- 距上次采购（天）
- 最近采购日期
- 首次采购日期
- 收货地址电话
- 买家等级
- 所在地
- 买家姓名
- 联系方式
- 跟进方式
- 重要等级
- 小记内容
- 操作

**数据获取方式：**
老客管理数据获取两次：
1. 第一次获取：页面加载后的默认数据
2. 第二次获取：点击重置按钮后的数据

**重置按钮选择器：**
```javascript
// 重置按钮
selector: '#app > div > div.bp-new-layout-wrapper > div.bp-new-layout-right > div.bp-new-layout-main.bp-zhizhou-layout-main > div > div.custom_tab_content > div > div.next-tabs-content > div > div.poten_range button'
```

**重置按钮点击方法：**
```javascript
// 点击重置按钮
const resetButton = await page.$('#app > div > div.bp-new-layout-wrapper > div.bp-new-layout-right > div.bp-new-layout-main.bp-zhizhou-layout-main > div > div.custom_tab_content > div > div.next-tabs-content > div > div.poten_range button');
if (resetButton) {
    await resetButton.click();
    console.log('已点击重置按钮');
    await new Promise(resolve => setTimeout(resolve, 3000));
}
```

**注意事项：**
- 不要使用 `body > div.next-overlay-wrapper.opened > ul > li` 选择器，因为下拉菜单可能不在预期的DOM位置
- 直接遍历所有 `li` 元素，找到包含目标文本的选项
- 点击后需要等待数据加载（建议3-5秒）

**表格选择器：**
```javascript
// 标签选择器
老客管理标签: 'a[href="#!/oldbuyer"]' 或包含"老客管理"文本的元素

// 表格选择器
table selector: 'table'
row selector: 'tr'
cell selector: 'td, th'
```

### 使用方法

#### 方法一：使用puppeteer

```javascript
const puppeteer = require('puppeteer-core');
const fs = require('fs');

async function exportMarketingCRM() {
    const browser = await puppeteer.launch({
        executablePath: '/Applications/Google Chrome.app/Contents/MacOS/Google Chrome',
        headless: false,
        defaultViewport: { width: 1920, height: 1080 }
    });

    const page = await browser.newPage();
    
    // 1. 打开营销平台CRM页面
    await page.goto('https://p4p.1688.com/main.html#!/crmbuyer?from=tools', {
        waitUntil: 'networkidle2',
        timeout: 60000
    });

    // 2. 检测登录
    if (page.url().includes('login')) {
        console.log('请完成登录...');
        // 等待登录完成
    }

    // 3. 等待页面加载
    await new Promise(resolve => setTimeout(resolve, 5000));

    // 4. 获取潜客机会数据
    const potentialCustomers = await page.evaluate(() => {
        const table = document.querySelector('table');
        if (!table) return [];
        
        const rows = table.querySelectorAll('tr');
        const data = [];
        
        rows.forEach(row => {
            const cells = row.querySelectorAll('td, th');
            const rowData = Array.from(cells).map(cell => cell.innerText.trim());
            if (rowData.length > 0) {
                data.push(rowData);
            }
        });
        
        return data;
    });

    // 5. 切换到老客管理标签
    await page.evaluate(() => {
        const tabs = document.querySelectorAll('a[href^="#!"]');
        for (const tab of tabs) {
            if (tab.innerText.includes('老客管理')) {
                tab.click();
                break;
            }
        }
    });

    // 6. 等待数据加载
    await new Promise(resolve => setTimeout(resolve, 3000));

    // 7. 获取老客管理数据
    const oldCustomers = await page.evaluate(() => {
        const table = document.querySelector('table');
        if (!table) return [];
        
        const rows = table.querySelectorAll('tr');
        const data = [];
        
        rows.forEach(row => {
            const cells = row.querySelectorAll('td, th');
            const rowData = Array.from(cells).map(cell => cell.innerText.trim());
            if (rowData.length > 0) {
                data.push(rowData);
            }
        });
        
        return data;
    });

    // 8. 保存数据
    const result = {
        timestamp: new Date().toISOString(),
        url: page.url(),
        potentialCustomers: potentialCustomers,
        oldCustomers: oldCustomers
    };
    
    fs.writeFileSync('crm_marketing_data.json', JSON.stringify(result, null, 2));
    
    await browser.close();
}
```

---

## 数据源二：客户管理系统

### 主页面URL
```
https://work.1688.com/home/page/index.htm?_path_=sellerPro/2017Sellerbase_Customer/mycrm
```

### iframe结构
页面包含3个iframe：
1. **CRM iframe**（目标iframe）
   - URL: `https://air.1688.com/app/CSBC-modules/csbc-page-member-crm/index.html`
   - 包含客户管理数据
2. 洞察iframe
   - URL: `https://xstore.insights.1688.com/index.html?at_iframe=1&versionId=...`
3. 桥接iframe
   - URL: `https://widget.1688.com/front/ajax/bridge.html?target=brg-0-...`

### CRM iframe标签结构
```
├── AI老客跟进（默认激活）
├── AI新客拓展
└── 全部客户明细
```

### 请求链路
```
用户请求
    ↓
打开主页面URL
    ↓
检测登录状态
    ├─ 未登录 → 等待用户扫码登录
    └─ 已登录 → 继续
    ↓
查找CRM iframe（通过URL匹配 'csbc-page-member-crm'）
    ↓
访问iframe内容
    ↓
识别标签结构（.csbc-page-member-crm-tabs-item）
    ↓
点击目标标签（如"全部客户明细"）
    ↓
等待数据加载（5秒）
    ↓
提取表格数据
    ├─ 表格选择器: table
    └─ ant-design表格: .ant-table
    ↓
保存数据文件
    ├─ crm_data_local.json
    ├─ crm_data_local.csv
    └─ crm_iframe_content.html
```

### 数据表结构

#### 全部客户明细表格
**表头字段：**
- 客户信息（包含客户昵称）
- 客户身份
- 月采购频率
- 站内月采购金额
- 采购偏好
- 本店最后询盘时间
- 询盘总结
- 操作

**表格选择器：**
```javascript
// 主表格
table selector: 'table'

// ant-design表格
table selector: '.ant-table'

// 表格行
row selector: 'tr' 或 '.ant-table-row'

// 表格单元格
cell selector: 'td, th' 或 '.ant-table-cell'
```

### 使用方法

#### 方法一：使用puppeteer（推荐）

```javascript
const puppeteer = require('puppeteer-core');
const fs = require('fs');

async function exportCustomerManagement() {
    // 1. 启动浏览器
    const browser = await puppeteer.launch({
        executablePath: '/Applications/Google Chrome.app/Contents/MacOS/Google Chrome',
        headless: false,
        defaultViewport: { width: 1920, height: 1080 }
    });

    const page = await browser.newPage();
    
    // 2. 打开主页面
    await page.goto('https://work.1688.com/home/page/index.htm?_path_=sellerPro/2017Sellerbase_Customer/mycrm', {
        waitUntil: 'networkidle2',
        timeout: 60000
    });

    // 3. 检测登录
    if (page.url().includes('login')) {
        console.log('请完成登录...');
        // 等待登录完成（最多5分钟）
        // 登录成功后URL会变化
    }

    // 4. 查找CRM iframe
    const iframes = await page.$$('iframe');
    let crmFrame = null;
    
    for (let i = 0; i < iframes.length; i++) {
        const frame = await iframes[i].contentFrame();
        if (frame && frame.url().includes('csbc-page-member-crm')) {
            crmFrame = frame;
            break;
        }
    }

    // 5. 点击"全部客户明细"标签
    await crmFrame.evaluate(() => {
        const tabItems = document.querySelectorAll('.csbc-page-member-crm-tabs-item');
        for (const item of tabItems) {
            if (item.innerText.includes('全部客户明细')) {
                item.click();
                break;
            }
        }
    });

    // 6. 等待数据加载
    await new Promise(resolve => setTimeout(resolve, 5000));

    // 7. 提取表格数据
    const tables = await crmFrame.evaluate(() => {
        const results = [];
        const tables = document.querySelectorAll('table');
        
        tables.forEach(table => {
            const rows = table.querySelectorAll('tr');
            const tableData = [];
            
            rows.forEach(row => {
                const cells = row.querySelectorAll('td, th');
                const rowData = Array.from(cells).map(cell => cell.innerText.trim());
                if (rowData.length > 0) {
                    tableData.push(rowData);
                }
            });
            
            if (tableData.length > 0) {
                results.push(tableData);
            }
        });
        
        return results;
    });

    // 8. 保存数据
    fs.writeFileSync('crm_management_data.json', JSON.stringify(tables, null, 2));
    
    await browser.close();
}
```

---

## 完整执行流程

### 步骤1：导出营销平台CRM数据
```bash
# 执行脚本
node export_marketing_crm.js

# 输出文件
crm_marketing_data.json
crm_marketing_data.csv
```

### 步骤2：导出客户管理数据
```bash
# 执行脚本
node export_customer_management.js

# 输出文件
crm_management_data.json
crm_management_data.csv
```

---

## 输出文件

### 营销平台CRM数据文件
- `crm_marketing_data.json` - JSON格式完整数据
- `crm_marketing_data.csv` - CSV格式数据

### 客户管理数据文件
- `crm_management_data.json` - JSON格式完整数据
- `crm_management_data.csv` - CSV格式数据
- `crm_iframe_content.html` - iframe HTML内容

---

## 数据格式示例

### 营销平台CRM - 潜客机会
```json
{
  "timestamp": "2026-03-28T08:17:24.033Z",
  "url": "https://p4p.1688.com/main.html#!/crmbuyer",
  "potentialCustomers": [
    ["买家昵称", "买家行为", "最近询盘时间", "渠道", "表单留言", "联系电话", "询盘记录", "买家身份", "买家等级", "所在地", "小记内容", "操作"],
    ["张三", "浏览商品", "2026-03-27 10:30", "搜索", "询价", "13800138000", "3次", "企业买家", "金牌", "浙江杭州", "", ""]
  ]
}
```

### 营销平台CRM - 老客管理
```json
{
  "oldCustomersFirst": [
    ["", "买家昵称", "首单是否广告引导", "买家身份", "采购次数", "累计采购金额（元）", "距上次采购（天）", "最近采购日期", "首次采购日期", "收货地址电话", "买家等级", "所在地", "买家姓名", "联系方式", "跟进方式", "重要等级", "小记内容", "操作"]
  ],
  "oldCustomersSecond": [
    ["", "买家昵称", "首单是否广告引导", "买家身份", "采购次数", "累计采购金额（元）", "距上次采购（天）", "最近采购日期", "首次采购日期", "收货地址电话", "买家等级", "所在地", "买家姓名", "联系方式", "跟进方式", "重要等级", "小记内容", "操作"],
    ["", "tb579057721", "否", "", "4", "6851", "109", "2025-12-08 17:25:28", "2025-11-06 17:57:42", "获取号码", "L3", "广东省 肇庆市", "", "", "", "", "- - -", "记录旺旺"],
    ["", "qq13642743538", "否", "", "3", "4875", "116", "2025-12-01 16:31:12", "2025-11-03 10:18:20", "获取号码", "L3", "广东省 广州市", "", "", "", "", "- - -", "记录旺旺"]
  ]
}
```

**CSV格式示例：**
```csv
老客管理 - 重置后获取
表格 1 (共1行)
"","买家昵称","首单是否广告引导","买家身份","采购次数","累计采购金额（元）","距上次采购（天）","最近采购日期","首次采购日期","收货地址电话","买家等级","所在地","买家姓名","联系方式","跟进方式","重要等级","小记内容","操作"

表格 2 (共10行)
"","tb579057721","否","","4","6851","109","2025-12-08 17:25:28","2025-11-06 17:57:42","获取号码","L3","广东省 肇庆市","","","","","- - -","记录旺旺"
"","qq13642743538","否","","3","4875","116","2025-12-01 16:31:12","2025-11-03 10:18:20","获取号码","L3","广东省 广州市","","","","","- - -","记录旺旺"
```

### 客户管理 - 全部客户明细
```json
{
  "timestamp": "2026-03-28T09:29:18.425Z",
  "url": "https://air.1688.com/app/CSBC-modules/csbc-page-member-crm/index.html",
  "title": "测试页面",
  "tabs": [
    {
      "index": 0,
      "text": "AI老客跟进",
      "isActive": true
    },
    {
      "index": 1,
      "text": "AI新客拓展",
      "isActive": false
    },
    {
      "index": 2,
      "text": "全部客户明细",
      "isActive": false
    }
  ],
  "tables": [
    {
      "tableIndex": 0,
      "rowCount": 1,
      "data": [
        ["", "客户身份", "月采购频率", "站内月采购金额", "采购偏好", "本店最后询盘时间", "询盘总结", ""]
      ]
    }
  ]
}
```

---

## 注意事项

### 1. 登录处理
- 检测URL是否包含 `login`
- 使用 `headless: false` 模式让用户看到登录页面
- 等待URL变化确认登录成功
- 最长等待时间：5分钟

### 2. 页面加载
- 营销平台CRM：单页应用，需要等待hash路由加载
- 客户管理：iframe嵌套，需要等待iframe加载

### 3. iframe访问（仅客户管理系统）
- 必须使用 `contentFrame()` 方法获取iframe内容
- 通过URL匹配确认正确的iframe
- iframe可能跨域，需要正确处理

### 4. 数据加载
- 点击标签后需要等待数据加载
- 建议等待时间：3-5秒
- 可以通过检查表格行数确认数据是否加载完成

### 5. 空数据处理
- 表格可能只包含表头，没有数据行
- HTML中会显示"暂无数据"或"没有数据"
- CSV文件仍会包含表头

### 6. 页面结构变化
- 标签名称可能变化
- 表格字段可能调整
- 建议定期检查HTML文件确认结构

---

## 故障排查

### 问题1：营销平台CRM找不到标签
**原因：** 页面未完全加载或hash路由未生效
**解决：** 增加等待时间，确保hash路由已加载

### 问题2：客户管理系统找不到iframe
**原因：** 页面未完全加载
**解决：** 增加等待时间，确保iframe已加载

### 问题3：表格数据为空
**原因：** 
- 数据确实为空
- 数据未加载完成
- 标签未正确点击

**解决：**
- 检查HTML文件确认数据状态
- 增加等待时间
- 确认标签点击成功

### 问题4：登录失败
**原因：** 登录超时或网络问题
**解决：** 
- 增加等待时间
- 检查网络连接
- 手动完成登录后继续

### 问题5：重置按钮点击失败
**症状：** 
- 日志显示"未找到重置按钮"
- 数据未更新

**原因：** 
- 页面结构变化
- 选择器不正确

**解决方案：**
```javascript
// 检查重置按钮是否存在
const resetButton = await page.$('#app > div > div.bp-new-layout-wrapper > div.bp-new-layout-right > div.bp-new-layout-main.bp-zhizhou-layout-main > div > div.custom_tab_content > div > div.next-tabs-content > div > div.poten_range button');
if (resetButton) {
    await resetButton.click();
    console.log('已点击重置按钮');
}
```

---

## 完整示例脚本

脚本位于当前工作目录下：

```
.trae/skills/1688-data-export/scripts/export_all_crm.js
```

**注意：** 脚本使用相对路径，请在项目根目录下执行。

### 统一脚本（自动适配操作系统）
- `export_all_crm.js` - 完整数据导出脚本，自动检测操作系统并选择合适的浏览器配置

### 操作系统适配
脚本会自动检测当前操作系统：
- **macOS**: 使用 `/Applications/Google Chrome.app`，显示浏览器窗口（headless: false）
- **Linux**: 使用 `/usr/bin/chromium-browser`，无头模式运行（headless: 'new'），**登录状态持久化**

**Linux服务器登录状态持久化：**
- 用户数据保存在当前工作目录的 `browser_data` 目录中
- 首次登录后，登录状态会被保存
- 后续执行无需重复登录
- 如果登录过期，会自动生成新的二维码截图

### 执行方式（异步执行）

脚本必须以异步方式运行，启动后需要持续监控输出：

```bash
# 异步启动脚本（在项目根目录下执行）
node .trae/skills/1688-data-export/scripts/export_all_crm.js &

# 或者使用 RunCommand 工具，设置 blocking: false
```

**执行流程：**
1. 异步启动 `export_all_crm.js` 脚本
2. 每5秒检查一次脚本输出
3. 如果检测到需要登录（输出包含"需要登录"或"登录页面截图已保存"）：
   - 下载登录二维码截图 `login_qrcode_all.png`
   - 发送给用户扫码登录
   - 继续监控输出
4. 持续监控直到脚本运行完毕（输出包含"数据导出完成"）

**监控输出示例：**
```
====================================
1688 CRM数据导出工具
操作系统: darwin
====================================

【数据源一】营销平台CRM
...
需要登录！
登录页面截图已保存: login_qrcode_all.png
...（此时需要发送二维码给用户）
登录成功！
...
数据导出完成！
```

**关键输出标识：**
- `需要登录！` - 需要用户扫码登录
- `登录页面截图已保存` - 二维码已生成
- `登录成功！` - 登录完成
- `✓ 已点击重置按钮` - 重置按钮点击成功
- `数据导出完成！` - 脚本执行完毕

### 脚本功能：
1. 自动检测操作系统并配置浏览器
2. 打开浏览器（macOS显示窗口，Linux无头模式）
3. 访问营销平台CRM页面
4. 等待用户扫码登录（最长5分钟）
5. 获取"潜客机会"数据
6. 访问"潜客公海"URL并获取数据
7. 访问"老客管理"URL
8. 第一次获取"老客管理"数据
9. 点击重置按钮
10. 第二次获取"老客管理"数据
11. 访问客户管理系统
12. 查找CRM iframe
13. 点击"全部客户明细"标签
14. 提取表格数据
15. 保存JSON和CSV文件
16. 关闭浏览器
17. 发送文件给用户

### 发送文件给用户
脚本执行完成后，将CSV文件发送给用户：

**发送的文件：**
- `all_crm_data_YYYYMMDDHHMMSS.csv` - CSV格式数据（Excel兼容，带时间戳）

**时间戳格式：** `YYYYMMDDHHMMSS`（年月日时分秒）

**其他生成的文件（本地保存）：**
- `all_crm_data_YYYYMMDDHHMMSS.json` - JSON格式完整数据
- `marketing_crm_final.png` - 营销平台CRM截图
- `customer_management_final.png` - 客户管理系统截图

**发送方式：**
1. 使用文件读取工具读取CSV文件
2. 向用户展示文件路径，让用户自行下载

---

## 数据更新时间

### 营销平台CRM
- 潜客机会：当日实时数据
- 老客管理：当日实时数据

### 客户管理系统
- AI老客跟进：实时数据
- AI新客拓展：实时数据
- 全部客户明细：实时询盘客户数据，更新时延约15分钟
