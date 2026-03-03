# 学术研究数据 HTTP API

## 概述

本文档基于 AKShare 封装好的 Python 库，提供可以通过 Python 调用的学术研究数据接口。所有接口均来自 Policy Uncertainty、Fama-French 等公开数据源。

## 重要说明

- 所有接口均为 Python 函数调用
- 返回格式为 pandas.DataFrame
- 数据来源：Policy Uncertainty、Dartmouth 等公开 API
- 数据仅供学术研究使用，不构成投资建议

---

## 核心接口

### 1. 经济政策不确定性指数(EPU)

**接口名称**: 经济政策不确定性指数

**数据源**: Policy Uncertainty

**目标地址**: http://www.policyuncertainty.com/media/{Country}_Policy_Uncertainty_Data.csv

**描述**: 获取各国经济政策不确定性指数数据

**请求方式**: GET

**支持的国家/地区**:
| 国家/地区 | 参数值 | 说明 |
|---------|--------|------|
| 中国 | SCMP_China | 基于南华早报 |
| 美国 | US | 美国 |
| 香港 | HK | 香港 |
| 欧洲 | Europe | 欧洲(德国/法国/意大利) |
| 日本 | Japan | 日本 |
| 韩国 | Korea | 韩国 |
| 英国 | UK | 英国 |
| 俄罗斯 | Russia | 俄罗斯 |
| 印度 | India | 印度 |
| 巴西 | Brazil | 巴西 |
| 墨西哥 | Mexico | 墨西哥 |
| 澳大利亚 | Australia | 澳大利亚 |
| 加拿大 | Canada | 加拿大 |
| 智利 | Chile | 智利 |
| 哥伦比亚 | Colombia | 哥伦比亚 |
| 希腊 | Greece | 希腊 |
| 爱尔兰 | Ireland | 爱尔兰 |
| 荷兰 | Netherlands | 荷兰 |
| 新加坡 | Singapore | 新加坡 |
| 西班牙 | Spain | 西班牙 |
| 瑞典 | Sweden | 瑞典 |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| Year | 年份 |
| Month | 月份 |
| EPU | 经济政策不确定性指数 |
| ... | 其他相关字段 |

**Python 调用示例**:

```python
import akshare as ak

# 获取中国经济政策不确定性指数
df = ak.article_epu_index(symbol="China")
print(df)

# 获取美国经济政策不确定性指数
df = ak.article_epu_index(symbol="US")
print(df)

# 获取欧洲经济政策不确定性指数
df = ak.article_epu_index(symbol="Europe")
print(df)
```

---

### 2. 中国经济政策不确定性指数(Excel版)

**接口名称**: 中国经济政策不确定性指数

**数据源**: Policy Uncertainty

**目标地址**: http://www.policyuncertainty.com/media/HK_EPU_Data_Annotated.xlsx

**描述**: 获取中国经济政策不确定性指数详细数据(Excel格式)

**请求方式**: GET

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| Year | 年份 |
| Month | 月份 |
| EPU | 经济政策不确定性指数 |
| ... | 其他相关字段 |

**Python 调用示例**:

```python
import akshare as ak

# 获取Fama-French多因子模型数据
df = ak.article_ff_crr()
print(df)
```

---

### 4. FRED-MD宏观经济数据库

**接口名称**: FRED-MD宏观经济数据

**数据源**: Federal Reserve Bank of St. Louis

**目标地址**: https://fred.stlouisfed.org/

**描述**: 获取美国宏观经济数据

**请求方式**: GET

**Python 调用示例**:

```python
import akshare as ak

# 获取FRED波动率指数数据
df = ak.article_oman_rv(symbol="VIX")
print(df)
```

---

## 常用学术数据指标说明

### 经济政策不确定性指数(EPU)

EPU指数通过统计报纸中与经济政策不确定性相关的文章数量来衡量政策不确定性程度。指数越高，表示政策不确定性越大。

| 指数 | 说明 |
|------|------|
| EPU | 经济政策不确定性指数 |
| News_Based_Policy_Uncert_Index | 基于新闻的政策不确定性指数 |
| Three_Component_Index | 三成分指数 |

### Fama-French因子

| 因子 | 说明 |
|------|------|
| Mkt-RF | 市场风险溢价 |
| SMB | 规模因子(Small Minus Big) |
| HML | 价值因子(High Minus Low) |
| RMW | 盈利能力因子(Robust Minus Weak) |
| CMA | 投资因子(Conservative Minus Aggressive) |
| RF | 无风险利率 |

### VIX波动率指数

| 指标 | 说明 |
|------|------|
| VIX | 芝加哥期权交易所波动率指数 |
| VXN | 纳斯达克100波动率指数 |
| VXD | 道琼斯波动率指数 |

---

## 数据应用场景

| 应用场景 | 推荐数据 |
|---------|---------|
| 政策影响研究 | EPU指数 |
| 资产定价研究 | Fama-French因子 |
| 市场波动研究 | VIX指数 |
| 宏观经济研究 | FRED-MD数据 |
| 国际比较研究 | 各国EPU指数 |

---

## 注意事项

1. 学术数据通常按月或按日更新
2. 不同国家的EPU指数计算方法可能略有不同
3. Fama-French因子数据需要从官网下载CSV文件
4. 数据仅供学术研究，不构成投资建议
5. 使用数据时请引用原始数据来源
