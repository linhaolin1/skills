# 宏观经济数据 HTTP API

## 概述

本文档基于 AKShare 实际使用的数据源 API，提供可以直接通过 curl 调用的宏观经济数据接口。所有接口均来自东方财富网、金十数据等公开数据源。

## 重要说明

- 所有接口均为 HTTP GET 请求
- 返回格式为 JSON
- 数据来源：东方财富网、金十数据公开 API
- 数据仅供学术研究使用，不构成投资建议

---

## 核心接口

### 1. 中国宏观数据

**接口名称**: 中国宏观经济指标

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取中国宏观经济数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| reportName | string | 是 | 报告名称 | 具体报告名见下方 |
| columns | string | 是 | 字段 | ALL |
| source | string | 是 | 来源 | WEB |
| client | string | 是 | 客户端 | WEB |

**常用报告名称**:
| 报告名称 | 说明 |
|---------|------|
| RPT_ECONOMY_GDP | GDP数据 |
| RPT_ECONOMY_CPI | CPI数据 |
| RPT_ECONOMY_PPI | PPI数据 |
| RPT_ECONOMY_PMI | PMI数据 |

**curl 调用示例**:

```bash
# 获取GDP数据
curl -s "https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_ECONOMY_GDP&columns=ALL&source=WEB&client=WEB"

# 获取CPI数据
curl -s "https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_ECONOMY_CPI&columns=ALL&source=WEB&client=WEB"
```

---

### 2. 美国宏观数据

**接口名称**: 美国宏观经济指标

**数据源**: 金十数据

**目标地址**: https://datacenter-api.jin10.com/

**描述**: 获取美国宏观经济数据

**请求方式**: GET

**curl 调用示例**:

```bash
# 获取美国非农数据（示例）
curl -s "https://datacenter-api.jin10.com/data_center/economic/us_nonfarm" \
  -H "X-App-Id: lnFP5lxse24wPgtY" \
  -H "X-Version: 1.0.0"
```

---

### 3. 货币供应量

**接口名称**: 中国货币供应量

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取M0、M1、M2货币供应量数据

**请求方式**: GET

**curl 调用示例**:

```bash
# 获取货币供应量数据
curl -s "https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_ECONOMY_MONEY_SUPPLY&columns=ALL&source=WEB&client=WEB"
```

---

### 4. 社会融资规模

**接口名称**: 社会融资规模数据

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取社会融资规模增量数据

**请求方式**: GET

**curl 调用示例**:

```bash
# 获取社会融资规模数据
curl -s "https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_ECONOMY_SOCIAL_FINANCING&columns=ALL&source=WEB&client=WEB"
```

---

## 常用宏观经济指标

| 指标名称 | 说明 | 发布频率 |
|---------|------|---------|
| GDP | 国内生产总值 | 季度 |
| CPI | 消费者物价指数 | 月度 |
| PPI | 生产者物价指数 | 月度 |
| PMI | 采购经理人指数 | 月度 |
| M2 | 广义货币供应量 | 月度 |
| 社会融资规模 | 社会融资增量 | 月度 |
| 外汇储备 | 外汇储备规模 | 月度 |

---

## 注意事项

1. 宏观数据通常按月或按季度发布
2. 数据可能有修订
3. 数据仅供学术研究，不构成投资建议
