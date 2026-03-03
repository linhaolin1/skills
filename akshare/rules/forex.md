# 外汇数据 HTTP API

## 概述

本文档基于 AKShare 封装好的 Python 库，提供可以通过 Python 调用的外汇数据接口。所有接口均来自东方财富网等公开数据源。

## 重要说明

- 所有接口均为 Python 函数调用
- 返回格式为 pandas.DataFrame
- 数据来源：东方财富网公开 API
- 数据仅供学术研究使用，不构成投资建议

---

## 核心接口

### 1. 人民币外汇即期报价

**接口名称**: 人民币外汇即期报价

**数据源**: 中国外汇交易中心

**目标地址**: http://www.chinamoney.com.cn/r/cms/www/chinamoney/data/fx/rfx-sp-quot.json

**描述**: 获取人民币外汇即期买卖报价

**请求方式**: GET

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| ccyPair | 货币对 |
| bidPrc | 买报价 |
| askPrc | 卖报价 |
| midprice | 中间价 |
| time | 时间 |

**Python 调用示例**:

```python
import akshare as ak

# 获取人民币外汇即期报价
df = ak.forex_spot_em()
print(df)
```

---

### 2. 人民币外汇远掉报价

**接口名称**: 人民币外汇远掉报价

**数据源**: 中国外汇交易中心

**目标地址**: http://www.chinamoney.com.cn/r/cms/www/chinamoney/data/fx/rfx-sw-quot.json

**描述**: 获取人民币外汇远期掉期报价

**请求方式**: GET

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| ccyPair | 货币对 |
| label_1W | 1周期限 |
| label_1M | 1月期限 |
| label_3M | 3月期限 |
| label_6M | 6月期限 |
| label_9M | 9月期限 |
| label_1Y | 1年期限 |

**Python 调用示例**:

```python
import akshare as ak

# 获取人民币外汇远掉报价
df = ak.forex_spot_em()
print(df)
```

---

### 3. 外币对即期报价

**接口名称**: 外币对即期报价

**数据源**: 中国外汇交易中心

**目标地址**: http://www.chinamoney.com.cn/r/cms/www/chinamoney/data/fx/cpair-quot.json

**描述**: 获取外币对即期买卖报价

**请求方式**: GET

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| ccyPair | 货币对 |
| bidPrc | 买报价 |
| askPrc | 卖报价 |
| midprice | 中间价 |
| time | 时间 |

**Python 调用示例**:

```python
import akshare as ak

# 获取外币对即期报价
df = ak.forex_spot_em()
print(df)
```

---

### 4. 人民币汇率中间价

**接口名称**: 人民币汇率中间价

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取人民币汇率中间价数据

**请求方式**: GET

**Python 调用示例**:

```python
import akshare as ak

# 获取人民币汇率中间价
df = ak.macro_china_rmb()
print(df)
```

---

## 常用货币对代码

| 货币对 | 说明 |
|--------|------|
| USD/CNY | 美元/人民币 |
| EUR/CNY | 欧元/人民币 |
| JPY/CNY | 日元/人民币 |
| GBP/CNY | 英镑/人民币 |
| EUR/USD | 欧元/美元 |
| GBP/USD | 英镑/美元 |
| USD/JPY | 美元/日元 |

---

## 注意事项

1. 外汇市场24小时交易
2. 数据更新频率通常为每日一次
3. 数据仅供学术研究，不构成投资建议
4. 外汇有风险，投资需谨慎
