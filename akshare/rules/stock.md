# 股票数据 HTTP API

## 概述

本文档基于 AKShare 实际使用的数据源 API，提供可以直接通过 curl 调用的股票数据接口。所有接口均来自东方财富网等公开数据源。

## 重要说明

- 所有接口均为 HTTP GET 请求
- 返回格式为 JSON
- 数据来源：东方财富网公开 API
- 数据仅供学术研究使用，不构成投资建议

---

## 核心接口

### 1. A股历史行情数据（日/周/月线）

**接口名称**: 股票K线数据

**数据源**: 东方财富网

**目标地址**: https://push2his.eastmoney.com/api/qt/stock/kline/get

**描述**: 获取沪深京 A 股的历史K线数据，支持日线、周线、月线，支持前复权、后复权

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| fields1 | string | 是 | 字段1 | f1,f2,f3,f4,f5,f6 |
| fields2 | string | 是 | 字段2 | f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f116 |
| ut | string | 是 | 用户标识 | 7eea3edcaed734bea9cbfc24409ed989 |
| klt | string | 是 | K线类型 | 101(日线)/102(周线)/103(月线) |
| fqt | string | 是 | 复权类型 | 0(不复权)/1(前复权)/2(后复权) |
| secid | string | 是 | 证券ID | 市场代码.股票代码，如：0.000657 |
| beg | string | 是 | 开始日期 | 20240101 |
| end | string | 是 | 结束日期 | 20241231 |

**secid 说明**:
- 深圳市场（0/3开头）：市场代码为 0，如 0.000657
- 上海市场（6开头）：市场代码为 1，如 1.600519

**返回字段**:
```json
{
  "rc": 0,
  "rt": 6,
  "svr": 182482261,
  "lt": 1,
  "full": 1,
  "data": {
    "code": "000657",
    "market": 0,
    "name": "中钨高新",
    "klines": [
      "2024-02-20,10.50,10.80,10.90,10.45,125000,135000000,4.29,2.86,0.30,1.25",
      "2024-02-21,10.82,10.75,10.95,10.70,98000,106000000,2.31,-0.46,-0.05,0.98"
    ]
  }
}
```

**klines 字段说明**（逗号分隔）:
1. 日期
2. 开盘价
3. 收盘价
4. 最高价
5. 最低价
6. 成交量（手）
7. 成交额（元）
8. 振幅（%）
9. 涨跌幅（%）
10. 涨跌额（元）
11. 换手率（%）

**curl 调用示例**:

```bash
# 查询中钨高新(000657)最近的日线数据（前复权）
curl -X GET "https://push2his.eastmoney.com/api/qt/stock/kline/get?fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f116&ut=7eea3edcaed734bea9cbfc24409ed989&klt=101&fqt=1&secid=0.000657&beg=20240220&end=20260301"

# 查询贵州茅台(600519)周线数据（不复权）
curl -X GET "https://push2his.eastmoney.com/api/qt/stock/kline/get?fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f116&ut=7eea3edcaed734bea9cbfc24409ed989&klt=102&fqt=0&secid=1.600519&beg=20240101&end=20241231"

# 查询平安银行(000001)月线数据（后复权）
curl -X GET "https://push2his.eastmoney.com/api/qt/stock/kline/get?fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f116&ut=7eea3edcaed734bea9cbfc24409ed989&klt=103&fqt=2&secid=0.000001&beg=20230101&end=20241231"
```

**数据处理示例**（使用 jq）:

```bash
# 提取最近5天的收盘价
curl -s "https://push2his.eastmoney.com/api/qt/stock/kline/get?fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f116&ut=7eea3edcaed734bea9cbfc24409ed989&klt=101&fqt=1&secid=0.000657&beg=20240220&end=20260301" | jq -r '.data.klines[-5:] | .[]'

# 提取股票名称和代码
curl -s "https://push2his.eastmoney.com/api/qt/stock/kline/get?fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f116&ut=7eea3edcaed734bea9cbfc24409ed989&klt=101&fqt=1&secid=0.000657&beg=20240220&end=20260301" | jq '.data | {code, name}'
```

---

### 2. A股实时行情数据

**接口名称**: 沪深京A股实时行情

**数据源**: 东方财富网

**目标地址**: https://82.push2.eastmoney.com/api/qt/clist/get

**描述**: 获取所有沪深京 A 股的实时行情数据

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
| fs | string | 是 | 市场筛选 | m:0 t:6,m:0 t:80,m:1 t:2,m:1 t:23,m:0 t:81 s:2048 |
| fields | string | 是 | 返回字段 | f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f11,f62,f128,f136,f115,f152 |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| f12 | 股票代码 |
| f14 | 股票名称 |
| f2 | 最新价 |
| f3 | 涨跌幅（%） |
| f4 | 涨跌额 |
| f5 | 成交量（手） |
| f6 | 成交额（元） |
| f7 | 振幅（%） |
| f8 | 换手率（%） |
| f9 | 市盈率-动态 |
| f15 | 最高价 |
| f16 | 最低价 |
| f17 | 今开价 |
| f18 | 昨收价 |
| f10 | 量比 |
| f23 | 市净率 |
| f20 | 总市值 |
| f21 | 流通市值 |

**curl 调用示例**:

```bash
# 获取所有A股实时行情
curl -X GET "https://82.push2.eastmoney.com/api/qt/clist/get?pn=1&pz=5000&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f12&fs=m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23,m:0+t:81+s:2048&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f11,f62,f128,f136,f115,f152"

# 使用 jq 查找中钨高新的实时行情
curl -s "https://82.push2.eastmoney.com/api/qt/clist/get?pn=1&pz=5000&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f12&fs=m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23,m:0+t:81+s:2048&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f11,f62,f128,f136,f115,f152" | jq '.data.diff[] | select(.f14 == "中钨高新")'
```

---

### 3. 单只股票实时行情

**接口名称**: 个股实时行情

**目标地址**: https://push2.eastmoney.com/api/qt/stock/get

**描述**: 获取单只股票的实时行情详情

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| ut | string | 是 | 用户标识 | fa5fd1943c7b386f172d6893dbfba10b |
| invt | string | 是 | 参数 | 2 |
| fltt | string | 是 | 过滤类型 | 2 |
| fields | string | 是 | 返回字段 | f43,f57,f58,f169,f170,f46,f44,f51,f168,f47,f164,f163,f116,f60,f45,f52,f50,f48,f167,f117,f71,f161,f49,f530,f135,f136,f137,f138,f139,f141,f142,f144,f145,f147,f148,f140,f143,f146,f149,f55,f62,f162,f92,f173,f104,f105,f84,f85,f183,f184,f185,f186,f187,f188,f189,f190,f191,f192,f107,f111,f86,f177,f78,f110,f262,f263,f264,f267,f268,f250,f251,f252,f253,f254,f255,f256,f257,f258,f266,f269,f270,f271,f273,f274,f275,f127,f199,f128,f198,f259,f260,f261,f171,f277,f278,f279,f288,f152,f250,f251,f252,f253,f254,f255,f256,f257,f258 |
| secid | string | 是 | 证券ID | 0.000657 |

**curl 调用示例**:

```bash
# 获取中钨高新实时行情
curl -X GET "https://push2.eastmoney.com/api/qt/stock/get?ut=fa5fd1943c7b386f172d6893dbfba10b&invt=2&fltt=2&fields=f43,f57,f58,f169,f170,f46,f44,f51,f168,f47,f164,f163,f116,f60,f45,f52,f50,f48,f167,f117,f71,f161,f49,f530,f135,f136,f137,f138,f139,f141,f142,f144,f145,f147,f148,f140,f143,f146,f149,f55,f62,f162,f92,f173,f104,f105,f84,f85,f183,f184,f185,f186,f187,f188,f189,f190,f191,f192,f107,f111,f86,f177,f78,f110,f262,f263,f264,f267,f268&secid=0.000657"
```

---

### 4. 个股新闻数据

**接口名称**: 个股新闻

**数据源**: 东方财富网

**目标地址**: https://search-api-web.eastmoney.com/search/jsonp

**描述**: 获取指定股票的最新新闻资讯

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| cb | string | 是 | JSONP回调函数 | jQuery35101792940631092459_1764599530165 |
| param | string | 是 | 查询参数JSON | 见下方说明 |
| _ | string | 是 | 时间戳 | 1764599530176 |

**param参数说明**（需URL编码的JSON）:
```json
{
  "uid": "",
  "keyword": "股票代码",
  "type": ["cmsArticleWebOld"],
  "client": "web",
  "clientType": "web",
  "clientVersion": "curr",
  "param": {
    "cmsArticleWebOld": {
      "searchScope": "default",
      "sort": "default",
      "pageIndex": 1,
      "pageSize": 10,
      "preTag": "<em>",
      "postTag": "</em>"
    }
  }
}
```

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| title | 新闻标题 |
| content | 新闻内容摘要 |
| date | 发布时间 |
| mediaName | 文章来源 |
| code | 文章代码 |

**curl 调用示例**:

```bash
# 获取中钨高新(000657)个股新闻
curl -s "https://search-api-web.eastmoney.com/search/jsonp?cb=jQuery35101792940631092459_1764599530165&param=%7B%22uid%22%3A%22%22%2C%22keyword%22%3A%22000657%22%2C%22type%22%3A%5B%22cmsArticleWebOld%22%5D%2C%22client%22%3A%22web%22%2C%22clientType%22%3A%22web%22%2C%22clientVersion%22%3A%22curr%22%2C%22param%22%3A%7B%22cmsArticleWebOld%22%3A%7B%22searchScope%22%3A%22default%22%2C%22sort%22%3A%22default%22%2C%22pageIndex%22%3A1%2C%22pageSize%22%3A10%2C%22preTag%22%3A%22%3Cem%3E%22%2C%22postTag%22%3A%22%3C%2Fem%3E%22%7D%7D%7D&_=1764599530176"
```

---

### 5. 龙虎榜数据

**接口名称**: 龙虎榜详情

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取龙虎榜交易详情数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| sortColumns | string | 是 | 排序列 | SECURITY_CODE,TRADE_DATE |
| sortTypes | string | 是 | 排序类型 | 1,-1 |
| pageSize | string | 是 | 每页数量 | 5000 |
| pageNumber | string | 是 | 页码 | 1 |
| reportName | string | 是 | 报告名称 | RPT_DAILYBILLBOARD_DETAILSNEW |
| columns | string | 是 | 字段 | ALL |
| source | string | 是 | 来源 | WEB |
| client | string | 是 | 客户端 | WEB |
| filter | string | 是 | 筛选条件 | 日期范围 |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| SECURITY_CODE | 股票代码 |
| SECURITY_NAME_ABBR | 股票名称 |
| TRADE_DATE | 上榜日期 |
| CLOSE_PRICE | 收盘价 |
| CHANGE_RATE | 涨跌幅 |
| BILLBOARD_NET_AMT | 龙虎榜净买额 |
| BILLBOARD_BUY_AMT | 龙虎榜买入额 |
| BILLBOARD_SELL_AMT | 龙虎榜卖出额 |
| TURNOVERRATE | 换手率 |

**curl 调用示例**:

```bash
# 获取2026年2月龙虎榜数据
curl -s "https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=SECURITY_CODE,TRADE_DATE&sortTypes=1,-1&pageSize=5000&pageNumber=1&reportName=RPT_DAILYBILLBOARD_DETAILSNEW&columns=ALL&source=WEB&client=WEB&filter=(TRADE_DATE<='2026-02-28')(TRADE_DATE>='2026-02-01')"
```

---

### 6. 分红送配数据

**接口名称**: 分红送配

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取股票分红送配数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| sortColumns | string | 是 | 排序列 | PLAN_NOTICE_DATE |
| sortTypes | string | 是 | 排序类型 | -1 |
| pageSize | string | 是 | 每页数量 | 500 |
| pageNumber | string | 是 | 页码 | 1 |
| reportName | string | 是 | 报告名称 | RPT_SHAREBONUS_DET |
| columns | string | 是 | 字段 | ALL |
| source | string | 是 | 来源 | WEB |
| client | string | 是 | 客户端 | WEB |
| filter | string | 是 | 筛选条件 | 报告期 |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 代码 | 股票代码 |
| 名称 | 股票名称 |
| 送转股份-送转总比例 | 每10股送转比例 |
| 现金分红-现金分红比例 | 每10股分红金额 |
| 现金分红-股息率 | 股息率 |
| 预案公告日 | 公告日期 |
| 方案进度 | 方案状态 |

**curl 调用示例**:

```bash
# 获取2025年报分红送配数据
curl -s "https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=PLAN_NOTICE_DATE&sortTypes=-1&pageSize=500&pageNumber=1&reportName=RPT_SHAREBONUS_DET&columns=ALL&source=WEB&client=WEB&filter=(REPORT_DATE='2025-12-31')"
```

---

### 7. 沪深港通资金流向

**接口名称**: 沪深港通资金流向

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取沪深港通资金流向数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| reportName | string | 是 | 报告名称 | RPT_MUTUAL_QUOTA |
| columns | string | 是 | 字段 | 见下方说明 |
| quoteColumns | string | 是 | 行情字段 | 见下方说明 |
| quoteType | string | 是 | 行情类型 | 0 |
| pageNumber | string | 是 | 页码 | 1 |
| pageSize | string | 是 | 每页数量 | 2000 |
| sortTypes | string | 是 | 排序类型 | 1 |
| sortColumns | string | 是 | 排序列 | MUTUAL_TYPE |
| source | string | 是 | 来源 | WEB |
| client | string | 是 | 客户端 | WEB |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 交易日 | 交易日期 |
| 类型 | 沪港通/深港通 |
| 板块 | 板块名称 |
| 资金方向 | 流入/流出 |
| 成交净买额 | 净买入金额 |
| 资金净流入 | 净流入金额 |
| 当日资金余额 | 剩余额度 |

**curl 调用示例**:

```bash
# 获取沪深港通资金流向
curl -s "https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_MUTUAL_QUOTA&columns=TRADE_DATE,MUTUAL_TYPE,BOARD_TYPE,MUTUAL_TYPE_NAME,FUNDS_DIRECTION,INDEX_CODE,INDEX_NAME,BOARD_CODE&quoteColumns=status~07~BOARD_CODE,dayNetAmtIn~07~BOARD_CODE,dayAmtRemain~07~BOARD_CODE,dayAmtThreshold~07~BOARD_CODE,f104~07~BOARD_CODE,f105~07~BOARD_CODE,f106~07~BOARD_CODE,f3~03~INDEX_CODE~INDEX_f3,netBuyAmt~07~BOARD_CODE&quoteType=0&pageNumber=1&pageSize=2000&sortTypes=1&sortColumns=MUTUAL_TYPE&source=WEB&client=WEB"
```

---

### 8. 行业板块行情

**接口名称**: 行业板块实时行情

**数据源**: 东方财富网

**目标地址**: https://17.push2.eastmoney.com/api/qt/clist/get

**描述**: 获取行业板块实时行情数据

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
| fs | string | 是 | 市场筛选 | m:90 t:2 f:!50 |
| fields | string | 是 | 返回字段 | 见下方说明 |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 板块代码 | 板块代码 |
| 板块名称 | 板块名称 |
| 最新价 | 当前价格 |
| 涨跌幅 | 涨跌幅度(%) |
| 总市值 | 板块总市值 |
| 换手率 | 板块换手率 |
| 上涨家数 | 板块内上涨股票数 |
| 下跌家数 | 板块内下跌股票数 |
| 领涨股票 | 板块领涨股票名称 |

**curl 调用示例**:

```bash
# 获取行业板块实时行情
curl -s "https://17.push2.eastmoney.com/api/qt/clist/get?pn=1&pz=100&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f3&fs=m:90+t:2+f:!50&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f26,f22,f33,f11,f62,f128,f136,f115,f152,f124,f107,f104,f105,f140,f141,f207,f208,f209,f222"
```

---

### 9. 概念板块行情

**接口名称**: 概念板块实时行情

**数据源**: 东方财富网

**目标地址**: https://79.push2.eastmoney.com/api/qt/clist/get

**描述**: 获取概念板块实时行情数据

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
| fid | string | 是 | 字段ID | f12 |
| fs | string | 是 | 市场筛选 | m:90 t:3 f:!50 |
| fields | string | 是 | 返回字段 | 见下方说明 |

**curl 调用示例**:

```bash
# 获取概念板块实时行情
curl -s "https://79.push2.eastmoney.com/api/qt/clist/get?pn=1&pz=100&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f12&fs=m:90+t:3+f:!50&fields=f2,f3,f4,f8,f12,f14,f15,f16,f17,f18,f20,f21,f24,f25,f22,f33,f11,f62,f128,f124,f107,f104,f105,f136"
```

---

### 10. 融资融券数据

**接口名称**: 融资融券账户统计

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取融资融券账户统计数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| reportName | string | 是 | 报告名称 | RPTA_WEB_MARGIN_DAILYTRADE |
| columns | string | 是 | 字段 | ALL |
| pageNumber | string | 是 | 页码 | 1 |
| pageSize | string | 是 | 每页数量 | 500 |
| sortColumns | string | 是 | 排序列 | STATISTICS_DATE |
| sortTypes | string | 是 | 排序类型 | -1 |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 日期 | 统计日期 |
| 融资余额 | 融资余额 |
| 融券余额 | 融券余额 |
| 融资买入额 | 当日融资买入额 |
| 融券卖出额 | 当日融券卖出额 |
| 个人投资者数量 | 参与融资融券的个人投资者数 |
| 机构投资者数量 | 参与融资融券的机构投资者数 |
| 担保物总价值 | 担保物总价值 |
| 平均维持担保比例 | 平均维持担保比例 |

**curl 调用示例**:

```bash
# 获取融资融券账户统计
curl -s "https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPTA_WEB_MARGIN_DAILYTRADE&columns=ALL&pageNumber=1&pageSize=500&sortColumns=STATISTICS_DATE&sortTypes=-1"
```

---

### 11. 业绩快报数据

**接口名称**: 业绩快报

**数据源**: 东方财富网

**目标地址**: https://datacenter.eastmoney.com/securities/api/data/v1/get

**描述**: 获取上市公司业绩快报数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| sortColumns | string | 是 | 排序列 | UPDATE_DATE,SECURITY_CODE |
| sortTypes | string | 是 | 排序类型 | -1,-1 |
| pageSize | string | 是 | 每页数量 | 500 |
| pageNumber | string | 是 | 页码 | 1 |
| reportName | string | 是 | 报告名称 | RPT_FCI_PERFORMANCEE |
| columns | string | 是 | 字段 | ALL |
| filter | string | 是 | 筛选条件 | 报告期 |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 股票代码 | 证券代码 |
| 股票简称 | 证券名称 |
| 每股收益 | 每股收益 |
| 营业收入 | 营业收入 |
| 净利润 | 净利润 |
| 净利润同比增长 | 净利润同比增长率 |
| 净资产收益率 | ROE |

**curl 调用示例**:

```bash
# 获取2025年报业绩快报
curl -s "https://datacenter.eastmoney.com/securities/api/data/v1/get?sortColumns=UPDATE_DATE,SECURITY_CODE&sortTypes=-1,-1&pageSize=500&pageNumber=1&reportName=RPT_FCI_PERFORMANCEE&columns=ALL&filter=(SECURITY_TYPE_CODE%20in%20(%22058001001%22,%22058001008%22))(TRADE_MARKET_CODE!%3D%22069001017%22)(REPORT_DATE%3D'2025-12-31')"
```

---

### 12. 股东持股统计

**接口名称**: 十大流通股东统计

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取十大流通股东持股统计数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| sortColumns | string | 是 | 排序列 | STATISTICS_TIMES,COOPERATION_HOLDER_MARK |
| sortTypes | string | 是 | 排序类型 | -1,-1 |
| pageSize | string | 是 | 每页数量 | 500 |
| pageNumber | string | 是 | 页码 | 1 |
| reportName | string | 是 | 报告名称 | RPT_COOPFREEHOLDERS_ANALYSIS |
| columns | string | 是 | 字段 | ALL |
| source | string | 是 | 来源 | WEB |
| client | string | 是 | 客户端 | WEB |
| filter | string | 是 | 筛选条件 | 报告期 |

**curl 调用示例**:

```bash
# 获取十大流通股东统计
curl -s "https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=STATISTICS_TIMES,COOPERATION_HOLDER_MARK&sortTypes=-1,-1&pageSize=500&pageNumber=1&reportName=RPT_COOPFREEHOLDERS_ANALYSIS&columns=ALL&source=WEB&client=WEB&filter=(HOLDNUM_CHANGE_TYPE%3D%22001%22)(END_DATE%3D'2025-06-30')"
```

---

## 完整查询示例

### 示例1：查询中钨高新最近5天行情

```bash
#!/bin/bash

# 股票信息
STOCK_CODE="000657"
STOCK_NAME="中钨高新"
MARKET_CODE="0"  # 深圳市场

# 日期范围（最近15天，确保包含5个交易日）
END_DATE=$(date +%Y%m%d)
START_DATE=$(date -v-15d +%Y%m%d)  # macOS
# START_DATE=$(date -d "15 days ago" +%Y%m%d)  # Linux

echo "查询 ${STOCK_NAME}(${STOCK_CODE}) 最近行情..."
echo "日期范围: ${START_DATE} - ${END_DATE}"
echo ""

# 调用API
curl -s "https://push2his.eastmoney.com/api/qt/stock/kline/get?fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f116&ut=7eea3edcaed734bea9cbfc24409ed989&klt=101&fqt=1&secid=${MARKET_CODE}.${STOCK_CODE}&beg=${START_DATE}&end=${END_DATE}" | jq -r '
  .data | 
  "股票: \(.name) (\(.code))",
  "",
  "最近5天行情:",
  "日期,开盘,收盘,最高,最低,成交量,成交额,振幅,涨跌幅,涨跌额,换手率",
  (.klines[-5:] | .[])
'
```

### 示例2：查询实时价格并计算涨跌

```bash
#!/bin/bash

# 获取中钨高新实时行情
echo "查询中钨高新实时行情..."

curl -s "https://82.push2.eastmoney.com/api/qt/clist/get?pn=1&pz=5000&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f12&fs=m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23,m:0+t:81+s:2048&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23" | jq -r '
  .data.diff[] | 
  select(.f14 == "中钨高新") | 
  "股票代码: \(.f12)",
  "股票名称: \(.f14)",
  "最新价: \(.f2)",
  "涨跌幅: \(.f3)%",
  "涨跌额: \(.f4)",
  "今开: \(.f17)",
  "昨收: \(.f18)",
  "最高: \(.f15)",
  "最低: \(.f16)",
  "成交量: \(.f5)手",
  "成交额: \(.f6)元",
  "换手率: \(.f8)%",
  "市盈率: \(.f9)",
  "市净率: \(.f23)"
'
```

### 示例3：对比多只股票

```bash
#!/bin/bash

# 对比中钨高新和贵州茅台
echo "对比股票行情..."

for stock in "0.000657:中钨高新" "1.600519:贵州茅台"; do
  IFS=':' read -r secid name <<< "$stock"
  echo ""
  echo "=== ${name} ==="
  
  curl -s "https://push2his.eastmoney.com/api/qt/stock/kline/get?fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f116&ut=7eea3edcaed734bea9cbfc24409ed989&klt=101&fqt=1&secid=${secid}&beg=20240220&end=20260301" | jq -r '
    .data.klines[-1] | 
    split(",") | 
    "日期: \(.[0])",
    "收盘价: \(.[2])",
    "涨跌幅: \(.[8])%"
  '
done
```

---

## 常用股票代码对照表

| 股票名称 | 股票代码 | secid | 市场 |
|---------|---------|-------|------|
| 平安银行 | 000001 | 0.000001 | 深圳 |
| 万科A | 000002 | 0.000002 | 深圳 |
| 中钨高新 | 000657 | 0.000657 | 深圳 |
| 贵州茅台 | 600519 | 1.600519 | 上海 |
| 中国平安 | 601318 | 1.601318 | 上海 |
| 工商银行 | 601398 | 1.601398 | 上海 |

---

## 注意事项

1. 所有接口均为公开数据源，无需认证
2. 建议合理控制请求频率，避免被限流
3. 股票代码格式：6位数字
4. secid 格式：市场代码.股票代码（深圳为0，上海为1）
5. 日期格式：YYYYMMDD（如：20240101）
6. 数据仅供学术研究，不构成投资建议
7. 股市有风险，投资需谨慎

---

## 错误处理

如果 API 返回错误或无数据，检查：
- 股票代码是否正确
- secid 的市场代码是否正确（深圳0，上海1）
- 日期范围是否合理
- 网络连接是否正常

```bash
# 检查API响应状态
response=$(curl -s -w "\n%{http_code}" "https://push2his.eastmoney.com/api/qt/stock/kline/get?fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f116&ut=7eea3edcaed734bea9cbfc24409ed989&klt=101&fqt=1&secid=0.000657&beg=20240220&end=20260301")
http_code=$(echo "$response" | tail -n1)
body=$(echo "$response" | sed '$d')

if [ "$http_code" -eq 200 ]; then
  echo "请求成功"
  echo "$body" | jq .
else
  echo "请求失败，HTTP状态码: $http_code"
fi
```
