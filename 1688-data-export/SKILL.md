---
name: "1688-data-export"
description: "导出1688 CRM客户数据（潜客机会、老客管理、AI老客跟进、AI新客拓展、全部客户明细）到Excel格式文件。Invoke when user wants to export 1688 CRM data, customer data, or get customer management information."
---

# 1688数据导出工具

Go语言实现的1688 CRM数据导出工具，使用chromedp进行浏览器自动化，专为Linux服务器部署优化。

---

## Linux服务器部署

### 1. 安装依赖

```bash
# Ubuntu/Debian
sudo apt update
sudo apt install -y chromium-browser

# CentOS/RHEL
sudo yum install -y chromium

# Alpine
apk add chromium
```

### 2. 下载程序

```bash
# 下载 Linux 版本
wget https://github.com/linhaolin1/skills/raw/main/1688-data-export/bin/crm-export-linux-amd64 -O crm-export
chmod +x crm-export
```

### 3. 首次登录（重要）

**Linux服务器无头模式无法扫码登录，需要先在本地登录获取登录状态：**

```bash
# 方法一：在本地Mac/Windows登录后上传browser_data
# 1. 在本地运行一次程序（会打开浏览器窗口）
./crm-export

# 2. 登录成功后，打包browser_data文件夹
tar -czvf browser_data.tar.gz browser_data

# 3. 上传到服务器并解压
scp browser_data.tar.gz user@server:/path/to/
ssh user@server "cd /path/to/ && tar -xzvf browser_data.tar.gz"
```

### 4. 运行

```bash
# 无头模式运行
./crm-export -H -o /data/crm

# 参数说明
# -H, --headless    无头模式（服务器必须）
# -o, --output      输出目录
# -t, --timeout     登录超时时间(秒)
```

### 5. 定时任务

```bash
# 编辑 crontab
crontab -e

# 每天早上8点执行
0 8 * * * cd /path/to/crm-export && ./crm-export -H -o /data/crm >> /var/log/crm-export.log 2>&1

# 每周一早上8点执行
0 8 * * 1 cd /path/to/crm-export && ./crm-export -H -o /data/crm >> /var/log/crm-export.log 2>&1
```

---

## 命令行参数

```bash
./crm-export [选项]

选项:
  -o, --output string   输出目录 (默认当前目录)
  -H, --headless        无头模式运行（服务器必须启用）
  -t, --timeout int     登录超时时间(秒) (默认 300)
  -h, --help            帮助信息
```

---

## 跨平台编译

```bash
# Linux AMD64（推荐）
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o crm-export-linux-amd64 .

# Linux ARM64
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o crm-export-linux-arm64 .

# macOS
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o crm-export-darwin-amd64 .
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o crm-export-darwin-arm64 .

# Windows
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o crm-export-windows-amd64.exe .
```

---

## 数据源

### 数据源一：营销平台CRM
- 潜客机会（潜在客户询盘数据）
- 潜客公海（潜在客户公海数据）
- 老客管理（老客户采购数据）

### 数据源二：客户管理系统
- AI老客跟进
- AI新客拓展
- 全部客户明细

---

## 输出文件

```
/data/crm/
├── all_crm_data_20260331080000.json
└── all_crm_data_20260331080000.csv
```

### JSON结构

```json
{
  "timestamp": "2026-03-31T08:00:00+08:00",
  "platform": "linux",
  "marketingCRM": {
    "potentialCustomers": [[...]],
    "potentialPublic": [[...]],
    "oldCustomersFirst": [[...]],
    "oldCustomersSecond": [[...]]
  },
  "customerManagement": {
    "url": "https://air.1688.com/...",
    "tabs": [...],
    "tables": [...]
  }
}
```

---

## 登录状态管理

程序在当前目录创建 `browser_data` 保存登录状态：

| 状态 | 说明 |
|------|------|
| 首次运行 | 需要扫码登录 |
| 登录成功 | 状态保存到 browser_data |
| 后续运行 | 自动使用已保存的登录状态 |
| 登录过期 | 需要重新登录 |

**服务器部署建议：**
1. 本地登录后上传 `browser_data`
2. 定期检查登录状态是否过期
3. 过期后重新上传新的 `browser_data`

---

## 故障排查

### 找不到Chrome/Chromium
```bash
# 检查是否安装
which chromium-browser || which chromium || which google-chrome

# 安装
sudo apt install -y chromium-browser
```

### 登录状态过期
```bash
# 删除旧的登录数据
rm -rf browser_data

# 重新从本地上传
```

### 权限不足
```bash
chmod +x crm-export
```

### 内存不足
```bash
# Chromium 占用较多内存，建议服务器至少 2GB
# 可添加 swap
sudo fallocate -l 2G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
```

### 查看日志
```bash
# 查看最近的日志
tail -100 /var/log/crm-export.log

# 实时查看
tail -f /var/log/crm-export.log
```

---

## 文件结构

```
1688-data-export/
├── main.go        # 主程序源码
├── go.mod         # Go模块配置
├── go.sum         # 依赖校验
├── build.sh       # 跨平台编译脚本
├── SKILL.md       # 本文档
└── .gitignore     # Git忽略配置
```

---

## 技术栈

- Go 1.21+
- chromedp（浏览器自动化）
- cobra（命令行）
- 输出格式：JSON, CSV
