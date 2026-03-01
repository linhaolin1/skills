# 期权数据 HTTP API

## 概述

本文档基于 AKShare 实际使用的数据源 API，提供可以直接通过 curl 调用的期权数据接口。所有接口均来自东方财富网、各期权交易所等公开数据源。

## 重要说明

- 所有接口均为 HTTP GET 请求
- 返回格式为 JSON
- 数据来源：东方财富网、各期权交易所公开 API
- 数据仅供学术研究使用，不构成投资建议

---

## 核心接口

### 1. 股票期权实时行情

**接口名称**: 股票期权实时行情列表

**数据源**: 东方财富网

**目标地址**: https://push2.eastmoney.com/api/qt/clist/get

**描述**: 获取股票期权的实时行情数据

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
| fs | string | 是 | 市场筛选 | 期权市场代码 |
| fields | string | 是 | 返回字段 | f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18 |

**curl 调用示例**:

```bash
# 获取上证50ETF期权
curl -s "https://push2.eastmoney.com/api/qt/clist/get?pn=1&pz=5000&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f12&fs=m:8&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18"
```

---

### 2. 期权历史K线数据

**接口名称**: 期权历史行情

**数据源**: 东方财富网

**目标地址**: https://push2his.eastmoney.com/api/qt/stock/kline/get

**描述**: 获取期权合约的历史K线数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| secid | string | 是 | 证券ID | 市场代码.期权代码 |
| klt | string | 是 | K线类型 | 101(日线) |
| fqt | string | 是 | 复权类型 | 1 |
| lmt | string | 是 | 限制条数 | 10000 |
| end | string | 是 | 结束日期 | 20500000 |
| fields1 | string | 是 | 字段1 | f1,f2,f3,f4,f5,f6 |
| fields2 | string | 是 | 字段2 | f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f116 |
| ut | string | 是 | 用户标识 | 7eea3edcaed734bea9cbfc24409ed989 |

**curl 调用示例**:

```bash
# 获取期权历史数据（需要具体的期权代码）
curl -s "https://push2his.eastmoney.com/api/qt/stock/kline/get?secid=8.10003720&klt=101&fqt=1&lmt=10000&end=20500000&fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f116&ut=7eea3edcaed734bea9cbfc24409ed989"
```

---

### 3. 期权综合行情

**接口名称**: 期权市场综合行情

**数据源**: 东方财富网

**目标地址**: https://23.push2.eastmoney.com/api/qt/clist/get

**描述**: 获取期权市场所有品种的综合行情数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| pn | string | 是 | 页码 | 1 |
| pz | string | 是 | 每页数量 | 100 |
| po | string | 是 | 排序 | 1 |
| np | string | 是 | 参数 | 1 |
| ut | string | 是 | 用户标识 | bd1d9ddb04089700cf9c27f6f7426281 |
| fltt | string | 是 | 过滤类型 | 2 |
| invt | string | 是 | 参数 | 2 |
| fid | string | 是 | 字段ID | f3 |
| fs | string | 是 | 市场筛选 | m:10,m:12,m:140,m:141,m:151,m:163,m:226 |
| fields | string | 是 | 返回字段 | f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f28,f11,f62,f128,f136,f115,f152,f133,f108,f163,f161,f162 |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 代码 | 期权代码 |
| 名称 | 期权名称 |
| 最新价 | 当前价格 |
| 涨跌额 | 涨跌金额 |
| 涨跌幅 | 涨跌幅度(%) |
| 成交量 | 成交量(手) |
| 成交额 | 成交金额 |
| 持仓量 | 持仓量 |
| 行权价 | 行权价格 |
| 剩余日 | 距离到期天数 |
| 日增 | 持仓量日变化 |
| 昨结 | 昨日结算价 |
| 今开 | 今日开盘价 |

**curl 调用示例**:

```bash
# 获取期权市场综合行情
curl -s "https://23.push2.eastmoney.com/api/qt/clist/get?pn=1&pz=100&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f3&fs=m:10,m:12,m:140,m:141,m:151,m:163,m:226&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f28,f11,f62,f128,f136,f115,f152,f133,f108,f163,f161,f162"
```

---

### 4. 期权希腊字母

**接口名称**: 期权风险指标

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取期权希腊字母等风险指标数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| reportName | string | 是 | 报告名称 | RPT_OPTION_RISK_INDICATOR |
| columns | string | 是 | 字段 | ALL |
| source | string | 是 | 来源 | WEB |
| client | string | 是 | 客户端 | WEB |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| Delta | 标的价格变化对期权价格的影响 |
| Gamma | 标的价格变化对Delta的影响 |
| Theta | 时间流逝对期权价格的影响 |
| Vega | 波动率变化对期权价格的影响 |
| Rho | 利率变化对期权价格的影响 |

**curl 调用示例**:

```bash
# 获取期权风险指标
curl -s "https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_OPTION_RISK_INDICATOR&columns=ALL&source=WEB&client=WEB"
```

---

### 5. 期权隐含波动率

**接口名称**: 期权隐含波动率分析

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取期权隐含波动率数据

**请求方式**: GET

**curl 调用示例**:

```bash
# 获取期权隐含波动率
curl -s "https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_OPTION_IMPLIED_VOLATILITY&columns=ALL&source=WEB&client=WEB"
```

---

### 6. 商品期权行情

**接口名称**: 商品期权实时行情

**数据源**: 东方财富网

**目标地址**: https://push2.eastmoney.com/api/qt/clist/get

**描述**: 获取商品期权的实时行情数据

**请求方式**: GET

**curl 调用示例**:

```bash
# 获取商品期权行情
curl -s "https://push2.eastmoney.com/api/qt/clist/get?pn=1&pz=500&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f3&fs=m:140,m:141,m:151&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18"
```

---

## 常用期权品种

| 期权类型 | 标的 | 交易所 |
|---------|------|--------|
| 50ETF期权 | 上证50ETF | 上交所 |
| 沪深300ETF期权 | 沪深300ETF | 上交所/深交所 |
| 股指期权 | 沪深300指数 | 中金所 |
| 铜期权 | 铜 | 上期所 |
| 豆粕期权 | 豆粕 | 大商所 |
| 白糖期权 | 白糖 | 郑商所 |
| 棉花期权 | 棉花 | 郑商所 |

---

## 期权基本概念

| 概念 | 说明 |
|------|------|
| 看涨期权(Call) | 买方有权以行权价买入标的 |
| 看跌期权(Put) | 买方有权以行权价卖出标的 |
| 行权价 | 期权合约约定的买卖价格 |
| 到期日 | 期权合约到期的日期 |
| 权利金 | 期权的价格 |
| 内在价值 | 标的价格与行权价的差额 |
| 时间价值 | 权利金减去内在价值 |

---

## 注意事项

1. 期权有到期日，到期后归零
2. 期权价格受标的价格、波动率、时间等多因素影响
3. 期权交易风险较高
4. 数据仅供学术研究，不构成投资建议
5. 期权有风险，投资需谨慎
