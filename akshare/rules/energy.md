# 能源数据 HTTP API

## 概述

本文档基于 AKShare 实际使用的数据源 API，提供可以直接通过 curl 调用的能源数据接口。所有接口均来自东方财富网、碳排放交易所等公开数据源。

## 重要说明

- 所有接口均为 HTTP GET 请求
- 返回格式为 JSON
- 数据来源：东方财富网、各碳排放交易所公开 API
- 数据仅供学术研究使用，不构成投资建议

---

## 核心接口

### 1. 汽柴油历史调价信息

**接口名称**: 汽柴油历史调价信息

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取汽柴油历史调价信息

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| reportName | string | 是 | 报告名称 | RPTA_WEB_YJ_BD |
| columns | string | 是 | 字段 | ALL |
| sortColumns | string | 是 | 排序字段 | dim_date |
| sortTypes | string | 是 | 排序方向 | -1 |
| token | string | 是 | 访问令牌 | 894050c76af8597a853f5b408b759f5d |
| pageNumber | string | 是 | 页码 | 1 |
| pageSize | string | 是 | 每页数量 | 1000 |
| source | string | 是 | 来源 | WEB |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 调整日期 | 油价调整日期 |
| 汽油价格 | 汽油价格(元/吨) |
| 柴油价格 | 柴油价格(元/吨) |
| 汽油涨跌 | 汽油涨跌幅度 |
| 柴油涨跌 | 柴油涨跌幅度 |

**curl 调用示例**:

```bash
curl -s "https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPTA_WEB_YJ_BD&columns=ALL&sortColumns=dim_date&sortTypes=-1&token=894050c76af8597a853f5b408b759f5d&pageNumber=1&pageSize=1000&source=WEB"
```

---

### 2. 全国各地区油价

**接口名称**: 全国各地区的汽油和柴油油价

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取全国各地区的汽油和柴油油价详情

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| reportName | string | 是 | 报告名称 | RPTA_WEB_YJ_JH |
| columns | string | 是 | 字段 | ALL |
| filter | string | 是 | 筛选条件 | (dim_date='2024-01-18') |
| sortColumns | string | 是 | 排序字段 | cityname |
| sortTypes | string | 是 | 排序方向 | 1 |
| token | string | 是 | 访问令牌 | 894050c76af8597a853f5b408b759f5d |
| pageNumber | string | 是 | 页码 | 1 |
| pageSize | string | 是 | 每页数量 | 1000 |
| source | string | 是 | 来源 | WEB |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 日期 | 调价日期 |
| 地区 | 省份/地区名称 |
| V_0 | 0号柴油价格 |
| V_92 | 92号汽油价格 |
| V_95 | 95号汽油价格 |
| V_89 | 89号汽油价格 |
| ZDE_0 | 0号柴油涨跌 |
| ZDE_92 | 92号汽油涨跌 |
| ZDE_95 | 95号汽油涨跌 |
| ZDE_89 | 89号汽油涨跌 |

**curl 调用示例**:

```bash
curl -s "https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPTA_WEB_YJ_JH&columns=ALL&filter=(dim_date='2024-01-18')&sortColumns=cityname&sortTypes=1&token=894050c76af8597a853f5b408b759f5d&pageNumber=1&pageSize=1000&source=WEB"
```

---

### 3. 碳交易行情信息

**接口名称**: 碳交易网行情信息

**数据源**: 碳交易网

**目标地址**: http://k.tanjiaoyi.com:8080/KDataController/getHouseDatasInAverage.do

**描述**: 获取各碳排放交易所行情数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| lcnK | string | 是 | 密钥 | 53f75bfcefff58e4046ccfa42171636c |
| brand | string | 是 | 品牌 | TAN |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 日期 | 交易日期 |
| 成交价 | 成交价格(元/吨) |
| 成交量 | 成交量(吨) |
| 成交额 | 成交金额 |
| 地点 | 交易所地点 |

**支持的交易所**:
- 湖北
- 上海
- 北京
- 重庆
- 广东
- 天津
- 深圳
- 福建

**curl 调用示例**:

```bash
curl -s "http://k.tanjiaoyi.com:8080/KDataController/getHouseDatasInAverage.do?lcnK=53f75bfcefff58e4046ccfa42171636c&brand=TAN"
```

---

### 4. 北京碳排放权交易行情

**接口名称**: 北京市碳排放权公开交易行情

**数据源**: 北京市碳排放权电子交易平台

**目标地址**: https://www.bjets.com.cn/article/jyxx/

**描述**: 获取北京市碳排放权公开交易行情

**请求方式**: GET

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 日期 | 交易日期 |
| 成交量 | 成交量(吨) |
| 成交均价 | 成交均价(元/吨) |
| 成交额 | 成交金额 |
| 成交单位 | 成交单位 |

**curl 调用示例**:

```bash
curl -s "https://www.bjets.com.cn/article/jyxx/" -H "User-Agent: Mozilla/5.0"
```

---

### 5. 深圳碳排放交易所国内碳情

**接口名称**: 深圳碳排放交易所国内碳情

**数据源**: 深圳碳排放交易所

**目标地址**: http://www.cerx.cn/dailynewsCN/index.htm

**描述**: 获取国内碳情每日行情数据

**请求方式**: GET

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 交易日期 | 交易日期 |
| 开盘价 | 开盘价(元/吨) |
| 最高价 | 最高价(元/吨) |
| 最低价 | 最低价(元/吨) |
| 成交均价 | 成交均价(元/吨) |
| 收盘价 | 收盘价(元/吨) |
| 成交量 | 成交量(吨) |
| 成交额 | 成交金额 |

**curl 调用示例**:

```bash
curl -s "http://www.cerx.cn/dailynewsCN/index.htm" -H "User-Agent: Mozilla/5.0"
```

---

### 6. 深圳碳排放交易所国际碳情

**接口名称**: 深圳碳排放交易所国际碳情

**数据源**: 深圳碳排放交易所

**目标地址**: http://www.cerx.cn/dailynewsOuter/index.htm

**描述**: 获取国际碳情每日行情数据

**请求方式**: GET

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 交易日期 | 交易日期 |
| 开盘价 | 开盘价 |
| 最高价 | 最高价 |
| 最低价 | 最低价 |
| 成交均价 | 成交均价 |
| 收盘价 | 收盘价 |
| 成交量 | 成交量 |
| 成交额 | 成交额 |

**curl 调用示例**:

```bash
curl -s "http://www.cerx.cn/dailynewsOuter/index.htm" -H "User-Agent: Mozilla/5.0"
```

---

### 7. 湖北碳排放权交易中心行情

**接口名称**: 湖北碳排放权交易中心现货交易数据

**数据源**: 湖北碳排放权交易中心

**目标地址**: https://www.hbets.cn/

**描述**: 获取湖北碳排放权交易中心配额每日概况

**请求方式**: GET

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 日期 | 交易日期 |
| 成交价 | 成交价格(元/吨) |
| 成交量 | 成交量(吨) |
| 最新 | 最新价 |
| 涨跌 | 涨跌幅度 |

**curl 调用示例**:

```bash
curl -s "https://www.hbets.cn/" -H "User-Agent: Mozilla/5.0"
```

---

### 8. 广州碳排放权交易中心行情

**接口名称**: 广州碳排放权交易中心行情信息

**数据源**: 广州碳排放权交易中心

**目标地址**: http://ets.cnemission.com/carbon/portalIndex/markethistory

**描述**: 获取广州碳排放权交易中心行情历史数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| Top | string | 是 | 排序 | 1 |
| beginTime | string | 是 | 开始时间 | 2010-01-01 |
| endTime | string | 是 | 结束时间 | 2030-09-12 |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 日期 | 交易日期 |
| 品种 | 交易品种 |
| 开盘价 | 开盘价(元/吨) |
| 收盘价 | 收盘价(元/吨) |
| 最高价 | 最高价(元/吨) |
| 最低价 | 最低价(元/吨) |
| 涨跌 | 涨跌额 |
| 涨跌幅 | 涨跌幅(%) |
| 成交数量 | 成交量(吨) |
| 成交金额 | 成交金额 |

**curl 调用示例**:

```bash
curl -s "http://ets.cnemission.com/carbon/portalIndex/markethistory?Top=1&beginTime=2010-01-01&endTime=2030-09-12"
```

---

## 常用能源品种代码

| 品种名称 | 代码 | 单位 |
|---------|------|------|
| 92号汽油 | V_92 | 元/升 |
| 95号汽油 | V_95 | 元/升 |
| 0号柴油 | V_0 | 元/升 |
| 碳排放配额 | CEA | 元/吨 |
| 国家核证自愿减排量 | CCER | 元/吨 |

---

## 注意事项

1. 油价调整周期通常为10个工作日
2. 碳排放交易时间与股票交易时间类似
3. 不同碳排放交易所价格可能存在差异
4. 数据仅供学术研究，不构成投资建议
5. 能源价格受国际市场影响较大
