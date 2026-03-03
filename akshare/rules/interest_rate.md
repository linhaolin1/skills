# 利率数据 HTTP API

## 概述

本文档基于 AKShare 封装好的 Python 库，提供可以通过 Python 调用的利率数据接口。所有接口均来自东方财富网、中国外汇交易中心等公开数据源。

## 重要说明

- 所有接口均为 Python 函数调用
- 返回格式为 pandas.DataFrame
- 数据来源：东方财富网、中国外汇交易中心公开 API
- 数据仅供学术研究使用，不构成投资建议

---

## 核心接口

### 1. 银行间拆借利率

**接口名称**: 银行间拆借利率数据

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取各市场的拆借利率数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| reportName | string | 是 | 报告名称 | RPT_IMP_INTRESTRATEN |
| columns | string | 是 | 字段 | REPORT_DATE,REPORT_PERIOD,IR_RATE,CHANGE_RATE,INDICATOR_ID,LATEST_RECORD,MARKET,MARKET_CODE,CURRENCY,CURRENCY_CODE |
| filter | string | 是 | 筛选条件 | 市场代码、货币代码、指标代码 |
| pageNumber | string | 是 | 页码 | 1 |
| pageSize | string | 是 | 每页数量 | 500 |
| sortTypes | string | 是 | 排序方向 | -1 |
| sortColumns | string | 是 | 排序字段 | REPORT_DATE |
| source | string | 是 | 来源 | WEB |
| client | string | 是 | 客户端 | WEB |

**市场代码对照表**:
| 市场 | 代码 |
|------|------|
| 上海银行同业拆借市场 | 001 |
| 中国银行同业拆借市场 | 002 |
| 伦敦银行同业拆借市场 | 003 |
| 欧洲银行同业拆借市场 | 004 |
| 香港银行同业拆借市场 | 005 |
| 新加坡银行同业拆借市场 | 006 |

**货币代码对照表**:
| 货币 | 代码 |
|------|------|
| 人民币 | CNY |
| 美元 | USD |
| 欧元 | EUR |
| 英镑 | GBP |
| 日元 | JPY |
| 港币 | HKD |
| 离岸人民币 | CNH |
| 新加坡元 | SGD |

**期限代码对照表**:
| 期限 | 代码 |
|------|------|
| 隔夜 | 001 |
| 1周 | 101 |
| 2周 | 102 |
| 3周 | 103 |
| 1月 | 201 |
| 2月 | 202 |
| 3月 | 203 |
| 6月 | 206 |
| 9月 | 209 |
| 1年 | 301 |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 报告日 | 统计日期 |
| 利率 | 利率值(%) |
| 涨跌 | 涨跌幅度 |

**Python 调用示例**:

```python
import akshare as ak

# 获取银行间拆借利率数据
df = ak.macro_china_shibor_all()
print(df)
```

---

### 2. 回购定盘利率

**接口名称**: 回购定盘利率历史数据

**数据源**: 中国外汇交易中心

**目标地址**: https://www.chinamoney.com.cn/r/cms/www/chinamoney/data/currency/frr-chrt.csv

**描述**: 获取回购定盘利率历史数据

**请求方式**: GET

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| date | 日期 |
| FR001 | 1天回购定盘利率(%) |
| FR007 | 7天回购定盘利率(%) |
| FR014 | 14天回购定盘利率(%) |

**Python 调用示例**:

```python
import akshare as ak

# 获取回购定盘利率数据
df = ak.forex_spot_em()
print(df)
```

---

### 3. 银银间回购定盘利率

**接口名称**: 银银间回购定盘利率历史数据

**数据源**: 中国外汇交易中心

**目标地址**: https://www.chinamoney.com.cn/r/cms/www/chinamoney/data/currency/fdr-chrt.csv

**描述**: 获取银银间回购定盘利率历史数据

**请求方式**: GET

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| date | 日期 |
| FDR001 | 银银间1天回购定盘利率(%) |
| FDR007 | 银银间7天回购定盘利率(%) |
| FDR014 | 银银间14天回购定盘利率(%) |

**Python 调用示例**:

```python
import akshare as ak

# 获取银银间回购定盘利率数据
df = ak.forex_spot_em()
print(df)
```

---

### 4. 人民币LPR利率

**接口名称**: 回购定盘利率历史数据查询

**数据源**: 中国外汇交易中心

**目标地址**: https://www.chinamoney.com.cn/ags/ms/cm-u-bk-currency/FrrHis

**描述**: 查询指定时间段的回购定盘利率历史数据

**请求方式**: POST

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| lang | string | 是 | 语言 | CN |
| startDate | string | 是 | 开始日期 | 2024-01-01 |
| endDate | string | 是 | 结束日期 | 2024-01-31 |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| date | 日期 |
| FR001 | 1天回购定盘利率(%) |
| FR007 | 7天回购定盘利率(%) |
| FR014 | 14天回购定盘利率(%) |
| FDR001 | 银银间1天回购定盘利率(%) |
| FDR007 | 银银间7天回购定盘利率(%) |
| FDR014 | 银银间14天回购定盘利率(%) |

**Python 调用示例**:

```python
import akshare as ak

# 获取LPR利率数据
df = ak.macro_china_lpr()
print(df)
```

---

## 常用利率指标说明

| 指标名称 | 英文缩写 | 说明 |
|---------|---------|------|
| 上海银行间同业拆放利率 | Shibor | 上海银行间同业拆放利率 |
| 中国银行间同业拆借利率 | Chibor | 中国银行间同业拆借利率 |
| 伦敦银行间同业拆借利率 | Libor | 伦敦银行间同业拆借利率 |
| 欧元银行间同业拆借利率 | Euribor | 欧元银行间同业拆借利率 |
| 香港银行间同业拆借利率 | Hibor | 香港银行间同业拆借利率 |
| 新加坡银行间同业拆借利率 | Sibor | 新加坡银行间同业拆借利率 |
| 回购定盘利率 | FR | 回购定盘利率 |
| 银银间回购定盘利率 | FDR | 银银间回购定盘利率 |

---

## 注意事项

1. 利率数据通常每个工作日更新
2. Shibor是中国货币市场的重要基准利率
3. 回购利率反映银行间市场资金面松紧程度
4. 数据仅供学术研究，不构成投资建议
5. 利率变动对债券、股票市场有重要影响
