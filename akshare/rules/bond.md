# 债券数据 HTTP API

## 概述

本文档基于 AKShare 实际使用的数据源 API，提供可以直接通过 curl 调用的债券数据接口。所有接口均来自东方财富网、中国债券信息网等公开数据源。

## 重要说明

- 所有接口均为 HTTP GET 请求
- 返回格式为 JSON
- 数据来源：东方财富网、中国债券信息网公开 API
- 数据仅供学术研究使用，不构成投资建议

---

## 核心接口

### 1. 中美国债收益率

**接口名称**: 中美国债收益率数据

**数据源**: 东方财富网

**目标地址**: https://datacenter.eastmoney.com/api/data/get

**描述**: 获取中国和美国各期限国债收益率数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| type | string | 是 | 数据类型 | RPTA_WEB_TREASURYYIELD |
| sty | string | 是 | 样式 | ALL |
| st | string | 是 | 排序字段 | SOLAR_DATE |
| sr | string | 是 | 排序方向 | -1 |
| token | string | 是 | 访问令牌 | 894050c76af8597a853f5b408b759f5d |
| p | string | 是 | 页码 | 1 |
| ps | string | 是 | 每页数量 | 500 |
| pageNo | string | 是 | 页码 | 1 |
| pageNum | string | 是 | 页码 | 1 |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| SOLAR_DATE | 日期 |
| EMM00166462 | 中国国债收益率5年 |
| EMM00166466 | 中国国债收益率10年 |
| EMM00166469 | 中国国债收益率30年 |
| EMM00588704 | 中国国债收益率2年 |
| EMG00001306 | 美国国债收益率2年 |
| EMG00001308 | 美国国债收益率5年 |
| EMG00001310 | 美国国债收益率10年 |
| EMG00001312 | 美国国债收益率30年 |

**curl 调用示例**:

```bash
# 获取中美国债收益率数据
curl -s "https://datacenter.eastmoney.com/api/data/get?type=RPTA_WEB_TREASURYYIELD&sty=ALL&st=SOLAR_DATE&sr=-1&token=894050c76af8597a853f5b408b759f5d&p=1&ps=500&pageNo=1&pageNum=1"
```

---

### 2. 可转债实时行情

**接口名称**: 可转债实时行情列表

**数据源**: 东方财富网

**目标地址**: https://datacenter-web.eastmoney.com/api/data/v1/get

**描述**: 获取可转债实时行情数据

**请求方式**: GET

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| sortColumns | string | 是 | 排序列 | SECUCODE |
| sortTypes | string | 是 | 排序类型 | -1 |
| pageSize | string | 是 | 每页数量 | 500 |
| pageNumber | string | 是 | 页码 | 1 |
| reportName | string | 是 | 报告名称 | RPT_BOND_CB_LIST |
| columns | string | 是 | 字段列表 | 具体字段见下方 |
| source | string | 是 | 来源 | WEB |
| client | string | 是 | 客户端 | WEB |

**curl 调用示例**:

```bash
# 获取可转债实时行情
curl -s "https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=SECUCODE&sortTypes=-1&pageSize=500&pageNumber=1&reportName=RPT_BOND_CB_LIST&columns=ALL&source=WEB&client=WEB"
```

---

### 3. 国债期货实时行情

**接口名称**: 国债期货行情

**数据源**: 东方财富网

**目标地址**: https://push2.eastmoney.com/api/qt/clist/get

**描述**: 获取国债期货实时行情数据

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
| fs | string | 是 | 市场筛选 | m:144 |
| fields | string | 是 | 返回字段 | f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18 |

**curl 调用示例**:

```bash
# 获取国债期货实时行情
curl -s "https://push2.eastmoney.com/api/qt/clist/get?pn=1&pz=100&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f12&fs=m:144&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18"
```

---

## 注意事项

1. 债券数据更新频率因品种而异
2. 可转债交易时间与股票相同
3. 国债期货交易时间与商品期货不同
4. 数据仅供学术研究，不构成投资建议
5. 债券有风险，投资需谨慎
