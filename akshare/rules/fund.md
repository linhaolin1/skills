# 基金数据 HTTP API

## 概述

本文档基于 AKShare 封装好的 Python 库，提供可以通过 Python 调用的基金数据接口。所有接口均来自东方财富网、天天基金网等公开数据源。

## 重要说明

- 所有接口均为 Python 函数调用
- 返回格式为 pandas.DataFrame
- 数据来源：东方财富网、天天基金网公开 API
- 数据仅供学术研究使用，不构成投资建议

---

## 核心接口

### 1. ETF实时行情

**接口名称**: ETF实时行情列表

**数据源**: 东方财富网

**目标地址**: https://88.push2.eastmoney.com/api/qt/clist/get

**描述**: 获取ETF基金的实时行情数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| pn | string | 是 | 页码 | 1 |
| pz | string | 是 | 每页数量 | 5000 |
| po | string | 是 | 排序 | 1 |
| np | string | 是 | 参数 | 1 |
| ut | string | 是 | 用户标识 | bd1d9ddb04089700cf9c27f6f7426281 |
| fltt | string | 是 | 过滤类型 | 2 |
| invt | string | 是 | 参数 | 2 |
| fid | string | 是 | 字段ID | f12 |
| fs | string | 是 | 市场筛选 | b:MK0404 |
| fields | string | 是 | 返回字段 | f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23 |

**Python 调用示例**:

```python
import akshare as ak

# 获取ETF实时行情数据
df = ak.fund_etf_spot_em()
print(df)
```

---

### 2. ETF历史K线数据

**接口名称**: ETF历史行情

**数据源**: 东方财富网

**目标地址**: https://push2his.eastmoney.com/api/qt/stock/kline/get

**描述**: 获取ETF基金的历史K线数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| secid | string | 是 | 证券ID | 市场代码.ETF代码，如：1.510300 |
| klt | string | 是 | K线类型 | 101(日线)/102(周线)/103(月线) |
| fqt | string | 是 | 复权类型 | 1 |
| lmt | string | 是 | 限制条数 | 10000 |
| end | string | 是 | 结束日期 | 20500000 |
| fields1 | string | 是 | 字段1 | f1,f2,f3,f4,f5,f6 |
| fields2 | string | 是 | 字段2 | f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f116 |
| ut | string | 是 | 用户标识 | 7eea3edcaed734bea9cbfc24409ed989 |

**Python 调用示例**:

```python
import akshare as ak

# 获取沪深300ETF(510300)历史K线数据
df = ak.fund_etf_hist_em(symbol="510300", period="daily", start_date="20230101", end_date="20250101")
print(df)
```

---

### 3. LOF基金实时行情

**接口名称**: LOF实时行情列表

**数据源**: 东方财富网

**目标地址**: https://2.push2.eastmoney.com/api/qt/clist/get

**描述**: 获取LOF基金的实时行情数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| pn | string | 是 | 页码 | 1 |
| pz | string | 是 | 每页数量 | 5000 |
| po | string | 是 | 排序 | 1 |
| np | string | 是 | 参数 | 1 |
| ut | string | 是 | 用户标识 | bd1d9ddb04089700cf9c27f6f7426281 |
| fltt | string | 是 | 过滤类型 | 2 |
| invt | string | 是 | 参数 | 2 |
| fid | string | 是 | 字段ID | f12 |
| fs | string | 是 | 市场筛选 | b:MK0405 |
| fields | string | 是 | 返回字段 | f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23 |

**Python 调用示例**:

```python
import akshare as ak

# 获取LOF基金实时行情数据
df = ak.fund_lof_spot_em()
print(df)
```

---

### 4. 开放式基金净值

**接口名称**: 开放式基金每日净值

**数据源**: 天天基金网

**目标地址**: https://fund.eastmoney.com/Data/Fund_JJJZ_Data.aspx

**描述**: 获取开放式基金的每日净值数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| t | string | 是 | 类型 | 1 |
| lx | string | 是 | 类型 | 1 |
| letter | string | 是 | 字母筛选 | 空 |
| gsid | string | 是 | 公司ID | 空 |
| text | string | 是 | 文本 | 空 |
| sort | string | 是 | 排序 | zdf,desc |
| page | string | 是 | 分页 | 1,50000 |
| dt | string | 是 | 时间戳 | 当前时间戳 |
| atfc | string | 是 | 参数 | 空 |
| onlySale | string | 是 | 仅销售 | 0 |

**Python 调用示例**:

```python
import akshare as ak

# 获取开放式基金每日净值数据
df = ak.fund_open_fund_daily_em()
print(df)
```

---

### 5. 基金基本信息

**接口名称**: 基金名称和类型

**数据源**: 天天基金网

**目标地址**: https://fund.eastmoney.com/js/fundcode_search.js

**描述**: 获取所有基金的代码、名称和类型

**请求方式**: GET

**Python 调用示例**:

```python
import akshare as ak

# 获取基金基本信息
df = ak.fund_overview_em(symbol="000001")
print(df)
```

---

### 6. 基金详细信息

**接口名称**: 单只基金详情

**数据源**: 天天基金网

**目标地址**: https://fund.eastmoney.com/pingzhongdata/{基金代码}.js

**描述**: 获取单只基金的详细信息，包括净值走势、累计净值等

**请求方式**: GET

**Python 调用示例**:

```python
import akshare as ak

# 获取单只基金的详细信息
df = ak.fund_info_index_em(symbol="000001")
print(df)
```

---

### 7. 基金累计收益率

**接口名称**: 基金累计收益率走势

**数据源**: 天天基金网

**目标地址**: https://api.fund.eastmoney.com/pinzhong/LJSYLZS

**描述**: 获取基金累计收益率走势数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| fundCode | string | 是 | 基金代码 | 000001 |
| indexcode | string | 是 | 指数代码 | 000300 |
| type | string | 是 | 周期类型 | m(1月)/q(3月)/hy(6月)/y(1年)/try(3年)/fiy(5年)/sy(今年来)/se(成立来) |

**Python 调用示例**:

```python
import akshare as ak

# 获取基金累计收益率走势数据
df = ak.fund_info_index_em(symbol="000001")
print(df)
```

---

## 常用ETF代码对照表

| ETF名称 | ETF代码 | secid | 跟踪指数 |
|---------|---------|-------|---------|
| 沪深300ETF | 510300 | 1.510300 | 沪深300 |
| 上证50ETF | 510050 | 1.510050 | 上证50 |
| 创业板ETF | 159915 | 0.159915 | 创业板指 |
| 中证500ETF | 510500 | 1.510500 | 中证500 |
| 科创50ETF | 588000 | 1.588000 | 科创50 |

---

## 注意事项

1. 基金净值通常每日更新一次
2. ETF可以像股票一样实时交易
3. 数据仅供学术研究，不构成投资建议
4. 基金有风险，投资需谨慎
