# 指数数据 HTTP API

## 概述

本文档基于 AKShare 封装好的 Python 库，提供可以通过 Python 调用的指数数据接口。所有接口均来自东方财富网、新浪财经等公开数据源。

## 重要说明

- 所有接口均为 Python 函数调用
- 返回格式为 pandas.DataFrame
- 数据来源：东方财富网、新浪财经公开 API
- 数据仅供学术研究使用，不构成投资建议

---

## 核心接口

### 1. 指数实时行情

**接口名称**: 指数实时行情列表

**数据源**: 东方财富网

**目标地址**: https://push2.eastmoney.com/api/qt/clist/get

**描述**: 获取各类指数的实时行情数据

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
| fs | string | 是 | 市场筛选 | m:1 s:000001,m:1 s:000300 等 |
| fields | string | 是 | 返回字段 | f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18 |

**Python 调用示例**:

```python
import akshare as ak

# 获取上证系列指数实时行情
df = ak.stock_zh_index_spot_em(symbol="上证系列指数")
print(df)

# 获取深证系列指数
df = ak.stock_zh_index_spot_em(symbol="深证系列指数")
print(df)
```

---

### 2. 指数历史K线数据

**接口名称**: 指数历史行情

**数据源**: 东方财富网

**目标地址**: https://push2his.eastmoney.com/api/qt/stock/kline/get

**描述**: 获取指数的历史K线数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| secid | string | 是 | 证券ID | 市场代码.指数代码，如：1.000001 |
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

# 获取上证指数历史K线数据
df = ak.stock_zh_index_daily_em(symbol="000001")
print(df)
```

---

### 3. 单只指数实时详情

**接口名称**: 指数实时详情

**数据源**: 东方财富网

**目标地址**: https://push2.eastmoney.com/api/qt/stock/get

**描述**: 获取单只指数的实时详情

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| ut | string | 是 | 用户标识 | fa5fd1943c7b386f172d6893dbfba10b |
| invt | string | 是 | 参数 | 2 |
| fltt | string | 是 | 过滤类型 | 2 |
| fields | string | 是 | 返回字段 | f43,f57,f58,f169,f170,f46,f44,f51,f168,f47,f164,f163 |
| secid | string | 是 | 证券ID | 1.000001 |

**Python 调用示例**:

```python
import akshare as ak

# 获取指数实时详情数据
df = ak.index_stock_info()
print(df)
```

---

## 常用指数代码对照表

| 指数名称 | 指数代码 | secid | 市场 |
|---------|---------|-------|------|
| 上证指数 | 000001 | 1.000001 | 上海 |
| 深证成指 | 399001 | 0.399001 | 深圳 |
| 沪深300 | 000300 | 1.000300 | 上海 |
| 创业板指 | 399006 | 0.399006 | 深圳 |
| 科创50 | 000688 | 1.000688 | 上海 |
| 上证50 | 000016 | 1.000016 | 上海 |
| 中证500 | 000905 | 1.000905 | 上海 |
| 中证1000 | 000852 | 1.000852 | 上海 |

---

## 注意事项

1. 指数数据通常实时更新
2. 不同指数的计算方法不同
3. 数据仅供学术研究，不构成投资建议
