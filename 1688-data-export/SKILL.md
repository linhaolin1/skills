---
name: "1688-data-export"
description: "导出1688 CRM客户数据（潜客机会、老客管理、AI老客跟进、AI新客拓展、全部客户明细）到Excel格式文件。Invoke when user wants to export 1688 CRM data, customer data, or get customer management information."
---

# 1688数据导出工具

Go语言实现的1688 CRM数据导出工具，使用chromedp进行浏览器自动化，支持跨平台运行。

## 快速开始

### 编译

```bash
cd .trae/skills/1688-data-export
go build -ldflags="-s -w" -o crm-export .
```

### 运行

```bash
./crm-export
```

### 命令行参数

```bash
./crm-export [选项]

选项:
  -o, --output string   输出目录 (默认当前目录)
  -H, --headless        无头模式运行（不显示浏览器窗口）
  -t, --timeout int     登录超时时间(秒) (默认 300)
```

### 跨平台编译

```bash
# macOS
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o crm-export-darwin-amd64 .
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o crm-export-darwin-arm64 .

# Linux
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o crm-export-linux-amd64 .

# Windows
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o crm-export-windows-amd64.exe .
```

---

## 数据源概览

### 数据源一：营销平台CRM
- **主入口**: `https://p4p.1688.com/main.html`
- **数据内容**:
  - 潜客机会（潜在客户询盘数据）
  - 潜客公海（潜在客户公海数据）
  - 老客管理（老客户采购数据）

### 数据源二：客户管理系统
- **URL**: `https://air.1688.com/app/CSBC-modules/csbc-page-member-crm/index.html`
- **数据内容**:
  - AI老客跟进
  - AI新客拓展
  - 全部客户明细

---

## 执行流程

```
启动程序
    ↓
打开营销平台CRM页面
    ↓
检测登录状态
    ├─ 未登录 → 等待用户扫码登录（最长300秒）
    └─ 已登录 → 继续
    ↓
获取"潜客机会"数据
    ↓
访问"潜客公海"URL并获取数据
    ↓
访问"老客管理"URL
    ↓
第一次获取"老客管理"数据（默认筛选）
    ↓
点击重置按钮
    ↓
第二次获取"老客管理"数据（全部数据）
    ↓
访问客户管理系统
    ↓
点击"全部客户明细"标签
    ↓
提取表格数据
    ↓
保存JSON和CSV文件
```

---

## 数据表结构

### 潜客机会表格
| 字段 | 说明 |
|------|------|
| 买家昵称 | 客户昵称 |
| 买家行为 | 浏览、询盘等 |
| 最近询盘时间 | 最后询盘日期 |
| 渠道 | 来源渠道 |
| 表单留言 | 留言内容 |
| 联系电话 | 电话号码 |
| 询盘记录 | 询盘次数 |
| 买家身份 | 企业/个人 |
| 买家等级 | L1-L6 |
| 所在地 | 地址 |
| 小记内容 | 备注信息 |

### 老客管理表格
| 字段 | 说明 |
|------|------|
| 买家昵称 | 客户昵称 |
| 首单是否广告引导 | 是/否 |
| 买家身份 | 客户类型 |
| 采购次数 | 历史采购次数 |
| 累计采购金额（元） | 总金额 |
| 距上次采购（天） | 时间间隔 |
| 最近采购日期 | 最后采购时间 |
| 首次采购日期 | 第一次采购时间 |
| 收货地址电话 | 联系方式 |
| 买家等级 | L1-L6 |
| 所在地 | 地址 |
| 买家姓名 | 姓名 |
| 联系方式 | 电话 |
| 跟进方式 | 跟进渠道 |
| 重要等级 | 优先级 |
| 小记内容 | 备注信息 |

### 全部客户明细表格
| 字段 | 说明 |
|------|------|
| 客户信息 | 客户昵称 |
| 客户身份 | 客户类型 |
| 月采购频率 | 每月采购次数 |
| 站内月采购金额 | 月度金额 |
| 采购偏好 | 偏好类别 |
| 本店最后询盘时间 | 最后询盘 |
| 询盘总结 | 总结信息 |

---

## 输出文件

程序运行后会生成两个文件：

- `all_crm_data_YYYYMMDDHHMMSS.json` - JSON格式完整数据
- `all_crm_data_YYYYMMDDHHMMSS.csv` - CSV格式数据（Excel兼容）

### JSON数据结构

```json
{
  "timestamp": "2026-03-31T00:34:22+08:00",
  "platform": "darwin",
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

### CSV文件结构

```
========================================
数据源一：营销平台CRM
========================================

潜客机会
[数据行...]

潜客公海
[数据行...]

老客管理 - 第一次获取
[数据行...]

老客管理 - 重置后获取
[数据行...]

========================================
数据源二：客户管理系统
========================================

全部客户明细
[数据行...]
```

---

## 登录状态持久化

程序会在当前工作目录创建 `browser_data` 文件夹保存浏览器用户数据：

- 首次运行需要扫码登录
- 登录状态会被保存
- 后续运行无需重复登录
- 如果登录过期，会再次显示登录页面

---

## 注意事项

### 1. 登录处理
- 默认显示浏览器窗口（headless: false）
- 检测到登录页面会等待用户扫码
- 最长等待时间：300秒（可通过 -t 参数调整）

### 2. 数据加载
- 页面导航后等待3-5秒确保数据加载
- 点击重置按钮后等待5秒
- 点击标签后等待5秒

### 3. 老客管理数据
- 第一次获取：默认筛选条件下的数据（可能为空）
- 第二次获取：点击重置按钮后的全部数据

### 4. 空数据处理
- 表格可能只包含表头
- CSV文件仍会包含表头信息

---

## 故障排查

### 问题1：找不到Chrome浏览器
**解决**: 确保系统已安装Chrome或Chromium浏览器

### 问题2：登录超时
**解决**: 使用 `-t` 参数增加超时时间
```bash
./crm-export -t 600
```

### 问题3：数据为空
**原因**: 
- 页面未完全加载
- 登录状态过期

**解决**:
- 删除 `browser_data` 文件夹重新登录
- 检查网络连接

### 问题4：权限问题（macOS）
**解决**: 
```bash
chmod +x crm-export
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

- **语言**: Go 1.21+
- **浏览器自动化**: chromedp
- **命令行**: cobra
- **输出格式**: JSON, CSV
