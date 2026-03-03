# 电影票房数据 HTTP API

## 概述

本文档基于 AKShare 封装好的 Python 库，提供可以通过 Python 调用的电影票房数据接口。所有接口均来自艺恩数据等公开数据源。

## 重要说明

- 接口为 Python 函数调用
- 返回格式为 pandas.DataFrame
- 数据来源：艺恩数据公开 API
- 数据仅供学术研究使用，不构成投资建议

---

## 核心接口

### 1. 实时票房

**接口名称**: 电影实时票房

**数据源**: 艺恩数据

**目标地址**: https://www.endata.com.cn/API/GetData.ashx

**描述**: 获取电影实时票房数据

**请求方式**: POST

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| showDate | string | 否 | 日期 | 空 |
| tdate | string | 是 | 当前日期 | 2024-02-20 |
| MethodName | string | 是 | 方法名 | BoxOffice_GetHourBoxOffice |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 排序 | 排名序号 |
| 影片名称 | 电影名称 |
| 实时票房 | 实时票房(万元) |
| 票房占比 | 票房占比(%) |
| 上映天数 | 上映天数 |
| 累计票房 | 累计票房(万元) |

**Python 调用示例**:

```python
import akshare as ak

# 获取电影实时票房数据
df = ak.movie_boxoffice_realtime()
print(df)
```

---

### 2. 单日票房

**接口名称**: 电影单日票房

**数据源**: 艺恩数据

**目标地址**: https://www.endata.com.cn/API/GetData.ashx

**描述**: 获取电影单日票房数据

**请求方式**: POST

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| sdate | string | 是 | 开始日期 | 2024-02-19 |
| edate | string | 是 | 结束日期 | 2024-02-18 |
| MethodName | string | 是 | 方法名 | BoxOffice_GetDayBoxOffice |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 排序 | 排名序号 |
| 影片名称 | 电影名称 |
| 单日票房 | 单日票房(万元) |
| 环比变化 | 环比变化(%) |
| 累计票房 | 累计票房(万元) |
| 平均票价 | 平均票价(元) |
| 场均人次 | 场均人次 |
| 口碑指数 | 口碑指数 |
| 上映天数 | 上映天数 |

**Python 调用示例**:

```python
import akshare as ak

# 获取电影单日票房数据
df = ak.movie_boxoffice_daily(date="20240219")
print(df)
```

---

### 3. 单周票房

**接口名称**: 电影单周票房

**数据源**: 艺恩数据

**目标地址**: https://www.endata.com.cn/API/GetData.ashx

**描述**: 获取电影单周票房数据

**请求方式**: POST

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| sdate | string | 是 | 周一日期 | 2024-02-12 |
| MethodName | string | 是 | 方法名 | BoxOffice_GetWeekInfoData |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 排序 | 排名序号 |
| 影片名称 | 电影名称 |
| 排名变化 | 与上周排名变化 |
| 单周票房 | 单周票房(万元) |
| 环比变化 | 环比变化(%) |
| 累计票房 | 累计票房(万元) |
| 平均票价 | 平均票价(元) |
| 场均人次 | 场均人次 |
| 口碑指数 | 口碑指数 |
| 上映天数 | 上映天数 |

**Python 调用示例**:

```python
import akshare as ak

# 获取电影单周票房数据
df = ak.movie_boxoffice_weekly(date="20240218")
print(df)
```

---

### 4. 单月票房

**接口名称**: 电影单月票房

**数据源**: 艺恩数据

**目标地址**: https://www.endata.com.cn/API/GetData.ashx

**描述**: 获取电影单月票房数据

**请求方式**: POST

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| startTime | string | 是 | 月初日期 | 2024-02-01 |
| MethodName | string | 是 | 方法名 | BoxOffice_GetMonthBox |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 排序 | 排名序号 |
| 影片名称 | 电影名称 |
| 单月票房 | 单月票房(万元) |
| 月度占比 | 月度票房占比(%) |
| 平均票价 | 平均票价(元) |
| 场均人次 | 场均人次 |
| 上映日期 | 上映日期 |
| 口碑指数 | 口碑指数 |
| 月内天数 | 月内上映天数 |

**Python 调用示例**:

```python
import akshare as ak

# 获取电影单月票房数据
df = ak.movie_boxoffice_monthly(date="20240218")
print(df)
```

---

### 5. 年度票房

**接口名称**: 电影年度票房

**数据源**: 艺恩数据

**目标地址**: https://www.endata.com.cn/API/GetData.ashx

**描述**: 获取电影年度票房数据

**请求方式**: POST

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| year | string | 是 | 年份 | 2024 |
| MethodName | string | 是 | 方法名 | BoxOffice_GetYearInfoData |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 排序 | 排名序号 |
| 影片名称 | 电影名称 |
| 类型 | 电影类型 |
| 总票房 | 总票房(万元) |
| 平均票价 | 平均票价(元) |
| 场均人次 | 场均人次 |
| 国家及地区 | 制片国家/地区 |
| 上映日期 | 上映日期 |

**Python 调用示例**:

```python
import akshare as ak

# 获取电影年度票房数据
df = ak.movie_boxoffice_yearly(date="20240218")
print(df)
```

---

### 6. 年度首周票房

**接口名称**: 电影年度首周票房

**数据源**: 艺恩数据

**目标地址**: https://www.endata.com.cn/API/GetData.ashx

**描述**: 获取电影年度首周票房数据

**请求方式**: POST

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| year | string | 是 | 年份 | 2024 |
| MethodName | string | 是 | 方法名 | BoxOffice_getYearInfo_fData |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 排序 | 排名序号 |
| 影片名称 | 电影名称 |
| 类型 | 电影类型 |
| 首周票房 | 首周票房(万元) |
| 占总票房比重 | 首周票房占比(%) |
| 场均人次 | 场均人次 |
| 国家及地区 | 制片国家/地区 |
| 上映日期 | 上映日期 |
| 首周天数 | 首周天数 |

**Python 调用示例**:

```python
import akshare as ak

# 获取电影年度首周票房数据
df = ak.movie_boxoffice_yearly_first_week(date="20201018")
print(df)
```

---

### 7. 影院日票房排行

**接口名称**: 影院日票房排行

**数据源**: 艺恩数据

**目标地址**: https://www.endata.com.cn/API/GetData.ashx

**描述**: 获取影院日票房排行榜

**请求方式**: POST

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| rowNum1 | string | 是 | 起始行 | 1 |
| rowNum2 | string | 是 | 结束行 | 100 |
| date | string | 是 | 日期 | 20240219 |
| MethodName | string | 是 | 方法名 | BoxOffice_GetCinemaDayBoxOffice |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 排序 | 排名序号 |
| 影院名称 | 影院名称 |
| 单日票房 | 单日票房(元) |
| 单日场次 | 单日场次 |
| 场均人次 | 场均人次 |
| 场均票价 | 场均票价(元) |
| 上座率 | 上座率(%) |

**Python 调用示例**:

```python
import akshare as ak

# 获取影院日票房排行数据
df = ak.movie_boxoffice_cinema_daily(date="20240219")
print(df)
```

---

### 8. 影院周票房排行

**接口名称**: 影院周票房排行

**数据源**: 艺恩数据

**目标地址**: https://www.endata.com.cn/API/GetData.ashx

**描述**: 获取影院周票房排行榜

**请求方式**: POST

**输入参数**:
| 参数名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|------|------|--------|
| dateID | string | 是 | 周ID | 1088 |
| rowNum1 | string | 是 | 起始行 | 1 |
| rowNum2 | string | 是 | 结束行 | 100 |
| MethodName | string | 是 | 方法名 | BoxOffice_GetCinemaWeekBoxOffice |

**返回字段说明**:
| 字段 | 说明 |
|------|------|
| 排序 | 排名序号 |
| 影院名称 | 影院名称 |
| 当周票房 | 当周票房(元) |
| 单银幕票房 | 单银幕票房(元) |
| 场均人次 | 场均人次 |
| 单日单厅票房 | 单日单厅票房(元) |
| 单日单厅场次 | 单日单厅场次 |

**Python 调用示例**:

```python
import akshare as ak

# 获取影院周票房排行数据
df = ak.movie_boxoffice_cinema_weekly(date="20240219")
print(df)
```

---

## 注意事项

1. 艺恩数据接口返回的数据为加密格式，需要使用JavaScript解密
2. 实时票房数据每小时更新一次
3. 单日票房数据通常在次日上午更新
4. 票房数据单位通常为万元
5. 数据仅供学术研究，不构成投资建议
6. 影院票房排行仅显示前100名
