# 其他数据 HTTP API

## 概述

本文档基于 AKShare 实际使用的数据源 API，提供可以直接通过 curl 调用的其他类型数据接口，包括新闻、公告、交易日历等。所有接口均来自东方财富网等公开数据源。

## 重要说明

- 所有接口均为 HTTP GET 请求
- 返回格式为 JSON
- 数据来源：东方财富网等公开 API
- 数据仅供学术研究使用，不构成投资建议

---

## 核心接口

### 1. 股票账户统计

**接口名称**: 股票账户统计

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取股票账户开户统计数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| reportName | string | 是 | 报告名称 | RPT_DATA_ACCOUNT_STATISTIC |
| columns | string | 是 | 字段 | ALL |
| source | string | 是 | 来源 | WEB |
| client | string | 是 | 客户端 | WEB |

**curl 调用示例**:

```bash
# 获取股票账户统计
curl -s "https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_DATA_ACCOUNT_STATISTIC&columns=ALL&source=WEB&client=WEB"
```

---

### 2. 龙虎榜数据

**接口名称**: 龙虎榜数据

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取龙虎榜交易数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| reportName | string | 是 | 报告名称 | RPT_BILLBOARD_DAILYDETAILS |
| columns | string | 是 | 字段 | ALL |
| source | string | 是 | 来源 | WEB |
| client | string | 是 | 客户端 | WEB |
| filter | string | 是 | 筛选条件 | 日期等 |

**curl 调用示例**:

```bash
# 获取龙虎榜数据
curl -s "https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_BILLBOARD_DAILYDETAILS&columns=ALL&source=WEB&client=WEB"
```

---

### 3. 股权质押数据

**接口名称**: 股权质押数据

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取股权质押统计数据

**请求方式**: GET

**curl 调用示例**:

```bash
# 获取股权质押市场概况
curl -s "https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_GPZY_MARKETPROFILE&columns=ALL&source=WEB&client=WEB"
```

---

### 4. 融资融券数据

**接口名称**: 融资融券数据

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取融资融券余额数据

**请求方式**: GET

**curl 调用示例**:

```bash
# 获取融资融券数据
curl -s "https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_RZRQ_LSHJ&columns=ALL&source=WEB&client=WEB"
```

---

### 5. 机构调研数据

**接口名称**: 机构调研数据

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取机构调研统计数据

**请求方式**: GET

**curl 调用示例**:

```bash
# 获取机构调研数据
curl -s "https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_ORG_SURVEY&columns=ALL&source=WEB&client=WEB"
```

---

### 6. 港股通资金流向

**接口名称**: 港股通资金流向

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取港股通资金流向数据

**请求方式**: GET

**curl 调用示例**:

```bash
# 获取港股通资金流向
curl -s "https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_HKSTOCK_FLOW&columns=ALL&source=WEB&client=WEB"
```

---

### 7. 涨跌停统计

**接口名称**: 涨跌停统计

**数据源**: 东方财富网

**目标地址**: https://push2.eastmoney.com/api/qt/clist/get

**描述**: 获取涨跌停股票列表

**请求方式**: GET

**curl 调用示例**:

```bash
# 获取涨停股列表
curl -s "https://push2.eastmoney.com/api/qt/clist/get?pn=1&pz=500&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f3&fs=m:0 t:6 f:!2,m:0 t:80 f:!2,m:1 t:2 f:!2,m:1 t:23 f:!2&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18"
```

---

### 8. 交易日历

**接口名称**: 交易日历

**数据源**: 各交易所

**目标地址**: 各交易所官网

**描述**: 获取交易日历数据

**请求方式**: GET

**curl 调用示例**:

```bash
# 上交所交易日历
curl -s "http://www.sse.com.cn/market/others/tradecalendar/"

# 深交所交易日历
curl -s "http://www.szse.cn/market/calendar/index.html"
```

---

## 注意事项

1. 新闻数据实时更新
2. 龙虎榜数据每日收盘后更新
3. 融资融券数据每日更新
4. 数据仅供学术研究，不构成投资建议
