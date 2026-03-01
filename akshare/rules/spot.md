# 现货数据 HTTP API

## 概述

本文档基于 AKShare 实际使用的数据源 API，提供可以直接通过 curl 调用的现货数据接口。所有接口均来自99期货、东方财富网等公开数据源。

## 重要说明

- 所有接口均为 HTTP GET 请求
- 返回格式为 JSON
- 数据来源：99期货、东方财富网公开 API
- 数据仅供学术研究使用，不构成投资建议

---

## 核心接口

### 1. 现货价格走势

**接口名称**: 大宗商品现货价格走势

**数据源**: 99期货

**目标地址**: https://centerapi.fx168api.com/app/qh/api/spot/trend

**描述**: 获取大宗商品现货价格历史数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| productId | string | 是 | 品种ID | 从品种对照表获取 |
| pageNo | string | 是 | 页码 | 1 |
| pageSize | string | 是 | 每页数量 | 50000 |
| startDate | string | 是 | 开始日期 | 空 |
| endDate | string | 是 | 结束日期 | 2050-01-01 |
| appCategory | string | 是 | 应用类别 | web |

**请求头**:
| 参数名 | 说明 |
|--------|------|
| User-Agent | 浏览器标识 |
| _pcc | 访问令牌（需要先获取） |
| Origin | https://www.99qh.com |
| Referer | https://www.99qh.com |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| date | 日期 |
| fp | 期货收盘价 |
| sp | 现货价格 |

**curl 调用示例**:

```bash
# 获取螺纹钢现货价格（需要先获取token）
curl -s "https://centerapi.fx168api.com/app/qh/api/spot/trend?productId=rb&pageNo=1&pageSize=50000&startDate=&endDate=2050-01-01&appCategory=web" \
  -H "User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36" \
  -H "Origin: https://www.99qh.com" \
  -H "Referer: https://www.99qh.com"
```

---

### 2. 现货与股票关联

**接口名称**: 现货与股票上下游

**数据源**: 东方财富网

**目标地址**: https://data.eastmoney.com/ifdata/xhgp.html

**描述**: 获取现货商品与相关股票的上下游关系

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| symbol | string | 是 | 行业分类 | 能源/化工/塑料/纺织/有色/钢铁/建材/农副 |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 商品名称 | 商品名称 |
| 最新价格 | 最新现货价格 |
| 近半年涨跌幅 | 近半年涨跌幅 |
| 生产商 | 相关生产商 |
| 下游用户 | 下游用户 |

**curl 调用示例**:

```bash
# 获取能源类现货与股票关联
curl -s "https://data.eastmoney.com/ifdata/xhgp.html"
```

---

### 3. 现货品种对照表

**接口名称**: 现货品种ID对照

**数据源**: 99期货

**目标地址**: https://www.99qh.com/data/spotTrend

**描述**: 获取现货品种名称和ID的对照关系

**请求方式**: GET

**curl 调用示例**:

```bash
# 获取现货品种对照表
curl -s "https://www.99qh.com/data/spotTrend"
```

---

## 常用现货品种代码

| 品种名称 | 代码 | 分类 |
|---------|------|------|
| 螺纹钢 | rb | 钢铁 |
| 热卷 | hc | 钢铁 |
| 铁矿石 | i | 有色 |
| 铜 | cu | 有色 |
| 铝 | al | 有色 |
| 原油 | sc | 能源 |
| 焦煤 | jm | 能源 |
| 焦炭 | j | 能源 |
| PTA | TA | 化工 |
| 甲醇 | MA | 化工 |

---

## 注意事项

1. 现货价格通常每日更新
2. 不同品种更新时间可能不同
3. 数据仅供学术研究，不构成投资建议
4. 现货价格受供需关系影响较大
