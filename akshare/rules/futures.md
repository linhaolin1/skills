# 期货数据 HTTP API

## 概述

本文档基于 AKShare 封装好的 Python 库，提供可以通过 Python 调用的期货数据接口。所有接口均来自东方财富网等公开数据源。

## 重要说明

- 所有接口均为 Python 函数调用
- 返回格式为 pandas.DataFrame
- 数据来源：东方财富网公开 API
- 数据仅供学术研究使用，不构成投资建议

---

## 核心接口

### 1. 期货历史行情数据

**接口名称**: 期货K线数据

**数据源**: 东方财富网

**目标地址**: https://push2his.eastmoney.com/api/qt/stock/kline/get

**描述**: 获取期货合约的历史K线数据，支持日线、周线、月线

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| secid | string | 是 | 证券ID | 市场代码.合约代码 |
| klt | string | 是 | K线类型 | 101(日线)/102(周线)/103(月线) |
| fqt | string | 是 | 复权类型 | 1 |
| lmt | string | 是 | 限制条数 | 10000 |
| end | string | 是 | 结束日期 | 20500000 |
| iscca | string | 是 | 参数 | 1 |
| fields1 | string | 是 | 字段1 | f1,f2,f3,f4,f5,f6,f7,f8 |
| fields2 | string | 是 | 字段2 | f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f62,f63,f64 |
| ut | string | 是 | 用户标识 | 7eea3edcaed734bea9cbfc24409ed989 |
| forcect | string | 是 | 参数 | 1 |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 时间 | 日期 |
| 开盘 | 开盘价 |
| 收盘 | 收盘价 |
| 最高 | 最高价 |
| 最低 | 最低价 |
| 成交量 | 成交量 |
| 成交额 | 成交额 |
| 涨跌幅 | 涨跌幅(%) |
| 涨跌 | 涨跌额 |
| 持仓量 | 持仓量 |

**Python 调用示例**:

```python
import akshare as ak

# 获取期货历史K线数据
df = ak.futures_zh_daily_sina(symbol="RB0")
print(df)
```

---

### 2. 期货品种对照表

**接口名称**: 期货品种映射

**数据源**: 东方财富网

**目标地址**: https://futsse-static.eastmoney.com/redis

**描述**: 获取期货品种和合约代码的映射关系

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| msgid | string | 是 | 消息ID | gnweb 或具体市场ID |

**Python 调用示例**:

```python
import akshare as ak

# 获取期货品种映射数据
df = ak.futures_symbol_mark()
print(df)
```

---

### 3. 期货实时行情（分时数据）

**接口名称**: 期货分时行情

**数据源**: 东方财富网

**目标地址**: https://push2his.eastmoney.com/api/qt/stock/trends2/get

**描述**: 获取期货合约的分时行情数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| secid | string | 是 | 证券ID | 市场代码.合约代码 |
| fields1 | string | 是 | 字段1 | f1,f2,f3,f4,f5,f6,f7,f8 |
| fields2 | string | 是 | 字段2 | f51,f52,f53,f54,f55,f56,f57,f58 |
| ut | string | 是 | 用户标识 | 7eea3edcaed734bea9cbfc24409ed989 |

**Python 调用示例**:

```python
import akshare as ak

# 获取期货分时数据
df = ak.futures_zh_minute_sina(symbol="IF2008")
print(df)
```

---

### 4. 期货库存数据

**接口名称**: 期货库存数据

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取期货品种的库存数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| reportName | string | 是 | 报告名称 | RPT_FUTU_STOCKDATA |
| columns | string | 是 | 字段 | SECURITY_CODE,TRADE_DATE,ON_WARRANT_NUM,ADDCHANGE |
| filter | string | 是 | 筛选条件 | 品种代码和日期 |
| pageNumber | string | 是 | 页码 | 1 |
| pageSize | string | 是 | 每页数量 | 500 |
| sortTypes | string | 是 | 排序类型 | -1 |
| sortColumns | string | 是 | 排序列 | TRADE_DATE |
| source | string | 是 | 来源 | WEB |
| client | string | 是 | 客户端 | WEB |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 日期 | 统计日期 |
| 库存 | 库存数量 |
| 增减 | 库存增减 |

**Python 调用示例**:

```python
import akshare as ak

# 获取期货库存数据
df = ak.futures_inventory_em(symbol="a")
print(df)
```

---

### 5. 国际期货行情

**接口名称**: 国际期货实时行情

**数据源**: 东方财富网

**目标地址**: https://quote.eastmoney.com/center/gridlist.html#futures_global

**描述**: 获取国际期货品种的实时行情数据

**请求方式**: GET

**Python 调用示例**:

```python
import akshare as ak

# 获取国际期货实时行情数据
df = ak.futures_global_spot_em()
print(df)
```

---

## 常用期货品种代码对照表

| 品种名称 | 合约代码 | 交易所 |
|---------|---------|--------|
| 螺纹钢 | rb | 上期所 |
| 热卷 | hc | 上期所 |
| 铁矿石 | i | 大商所 |
| 焦煤 | jm | 大商所 |
| 焦炭 | j | 大商所 |
| 铜 | cu | 上期所 |
| 铝 | al | 上期所 |
| 黄金 | au | 上期所 |
| 原油 | sc | 上期能源 |
| 沪深300股指 | IF | 中金所 |

---

## 注意事项

1. 期货合约代码需要加上交易所前缀
2. 主力合约使用"主连"后缀
3. 数据仅供学术研究，不构成投资建议
4. 期货有风险，投资需谨慎
