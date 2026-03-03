# 其他数据 HTTP API

## 概述

本文档基于 AKShare 封装好的 Python 库，提供可以通过 Python 调用的其他类型数据接口，包括新闻、公告、交易日历等。所有接口均来自东方财富网等公开数据源。

## 重要说明

- 所有接口均为 Python 函数调用
- 返回格式为 pandas.DataFrame
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

**Python 调用示例**:

```python
import akshare as ak

# 获取股票账户统计数据
df = ak.stock_account_statistics_em()
print(df)
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

**Python 调用示例**:

```python
import akshare as ak

# 获取龙虎榜详细数据
df = ak.stock_lhb_detail_em()
print(df)
```

---

### 3. 股权质押数据

**接口名称**: 股权质押数据

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取股权质押统计数据

**请求方式**: GET

**Python 调用示例**:

```python
import akshare as ak

# 获取股权质押数据
df = ak.stock_gpzy_pledge_ratio_em()
print(df)
```

---

### 4. 融资融券数据

**接口名称**: 融资融券数据

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取融资融券余额数据

**请求方式**: GET

**Python 调用示例**:

```python
import akshare as ak

# 获取融资融券数据
df = ak.stock_margin_account_info()
print(df)
```

---

### 5. 机构调研数据

**接口名称**: 机构调研数据

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取机构调研统计数据

**请求方式**: GET

**Python 调用示例**:

```python
import akshare as ak

# 获取机构调研数据
df = ak.stock_jgdy_tj_em()
print(df)
```

---

### 6. 港股通资金流向

**接口名称**: 港股通资金流向

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取港股通资金流向数据

**请求方式**: GET

**Python 调用示例**:

```python
import akshare as ak

# 获取港股通资金流向数据
df = ak.stock_hsgt_fund_flow_summary_em()
print(df)
```

---

### 7. 涨跌停统计

**接口名称**: 涨跌停统计

**数据源**: 东方财富网

**目标地址**: https://push2.eastmoney.com/api/qt/clist/get

**描述**: 获取涨跌停股票列表

**请求方式**: GET

**Python 调用示例**:

```python
import akshare as ak

# 获取涨停股列表数据
df = ak.stock_zt_pool_em()
print(df)
```

---

### 8. 交易日历

**接口名称**: 交易日历

**数据源**: 各交易所

**目标地址**: 各交易所官网

**描述**: 获取交易日历数据

**请求方式**: GET

**Python 调用示例**:

```python
import akshare as ak

# 获取交易日历数据
df = ak.tool_trade_date_hist_sina()
print(df)
```

---

## 注意事项

1. 新闻数据实时更新
2. 龙虎榜数据每日收盘后更新
3. 融资融券数据每日更新
4. 数据仅供学术研究，不构成投资建议
