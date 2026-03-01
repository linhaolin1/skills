---
name: AKShare 财经数据 HTTP API
description: 提供股票、期货、期权、基金、债券、外汇、加密货币等金融产品的 HTTP API 数据接口，支持通过 curl 直接调用
version: 2.0.0
author: AKShare Community
tags: [finance, stock, futures, fund, bond, forex, crypto, api, http, rest, curl]
---

# AKShare 财经数据 HTTP API 接口

## 概述

AKShare 财经数据 HTTP API 提供股票、期货、期权、基金、债券、外汇、加密货币等金融产品的数据获取能力。本文档基于 AKShare 实际使用的数据源 API，提供可以直接通过 curl 调用的接口说明。

## 特点

- **直接调用**: 使用 curl 命令直接访问数据源 API，无需安装任何库
- **数据源可靠**: 基于东方财富网、新浪财经等可信任数据源
- **覆盖全面**: 涵盖股票、期货、基金、债券、外汇等多个金融市场
- **实时更新**: 提供实时行情和历史数据

## 使用方式

所有接口均为 HTTP GET 请求，可以直接使用 curl 调用：

```bash
# 基本调用示例：获取股票历史数据
curl -s "https://push2his.eastmoney.com/api/qt/stock/kline/get?fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f116&ut=7eea3edcaed734bea9cbfc24409ed989&klt=101&fqt=1&secid=0.000657&beg=20240101&end=20241231"
```

### 数据格式说明
- 返回格式: JSON
- 日期字段: 字符串格式 "YYYY-MM-DD" 或 "YYYYMMDD"
- 数值字段: 数字类型（整数或浮点数）
- 文本字段: 字符串类型

## 数据分类

AKShare API 提供以下主要数据类别：

1. **股票数据** (Stock) - A股、港股、美股等股票市场数据
2. **期货数据** (Futures) - 国内外期货市场数据
3. **期权数据** (Option) - 股票期权和商品期权数据
4. **基金数据** (Fund) - 公募基金、私募基金、ETF等
5. **债券数据** (Bond) - 国债、企业债、可转债等
6. **指数数据** (Index) - 各类市场指数数据
7. **外汇数据** (Forex/FX) - 外汇汇率和行情数据
8. **宏观经济数据** (Macro) - 国内外宏观经济指标
9. **加密货币数据** (Crypto) - 数字货币行情数据
10. **现货数据** (Spot) - 大宗商品现货价格

## 使用指南

### 基本调用示例

```bash
# 获取A股历史行情（中钨高新）
curl -s "https://push2his.eastmoney.com/api/qt/stock/kline/get?fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f116&ut=7eea3edcaed734bea9cbfc24409ed989&klt=101&fqt=1&secid=0.000657&beg=20240101&end=20241231"

# 获取所有A股实时行情
curl -s "https://82.push2.eastmoney.com/api/qt/clist/get?pn=1&pz=5000&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f12&fs=m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23,m:0+t:81+s:2048&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23"
```

### 数据处理

返回的 JSON 数据可以使用各种工具处理：

```bash
# 使用 python 处理
curl -s "API_URL" | python3 -m json.tool

# 使用 jq 处理（如果已安装）
curl -s "API_URL" | jq .

# 保存到文件
curl -s "API_URL" > data.json
```

### 注意事项

1. 所有接口均为公开数据源，无需认证
2. 建议合理控制请求频率，避免被限流
3. 数据仅供学术研究使用，不构成投资建议
4. 建议添加错误处理和重试机制

## API 接口分类

详细的接口描述请参考 `rules/` 目录下的分类文档：

- [股票数据接口](rules/stock.md) - 股票行情、财务、公告等
- [期货数据接口](rules/futures.md) - 期货行情、持仓、库存等
- [期权数据接口](rules/option.md) - 期权行情、持仓等
- [基金数据接口](rules/fund.md) - 基金净值、持仓、评级等
- [债券数据接口](rules/bond.md) - 债券行情、发行、评级等
- [指数数据接口](rules/index.md) - 各类指数行情数据
- [外汇数据接口](rules/forex.md) - 外汇汇率和行情
- [宏观经济接口](rules/macro.md) - 宏观经济指标
- [加密货币接口](rules/crypto.md) - 数字货币行情
- [现货数据接口](rules/spot.md) - 大宗商品现货价格
- [其他数据接口](rules/others.md) - 新闻、事件、工具等

## 快速开始

### 示例1：查询股票最近行情

```bash
#!/bin/bash

# 查询中钨高新(000657)最近行情
STOCK_CODE="000657"
MARKET_CODE="0"  # 深圳市场
END_DATE=$(date +%Y%m%d)
START_DATE=$(date -v-15d +%Y%m%d)  # macOS

curl -s "https://push2his.eastmoney.com/api/qt/stock/kline/get?fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f116&ut=7eea3edcaed734bea9cbfc24409ed989&klt=101&fqt=1&secid=${MARKET_CODE}.${STOCK_CODE}&beg=${START_DATE}&end=${END_DATE}"
```

### 示例2：查询实时行情

```bash
#!/bin/bash

# 获取所有A股实时行情
curl -s "https://82.push2.eastmoney.com/api/qt/clist/get?pn=1&pz=5000&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f12&fs=m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23,m:0+t:81+s:2048&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23"
```

## 认证与限流

### 访问限制
- 直接访问公开数据源，无需 API Key
- 部分数据源可能有访问频率限制
- 建议合理控制请求频率，避免被数据源封禁

### 错误处理建议

```bash
#!/bin/bash

# 带重试机制的请求
function fetch_with_retry() {
  local url=$1
  local max_retries=3
  local retry=0
  
  while [ $retry -lt $max_retries ]; do
    response=$(curl -s -w "\n%{http_code}" "$url")
    http_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | sed '$d')
    
    if [ "$http_code" -eq 200 ]; then
      echo "$body"
      return 0
    else
      echo "请求失败，HTTP状态码: $http_code，重试 $((retry+1))/$max_retries" >&2
      retry=$((retry+1))
      sleep 2
    fi
  done
  
  return 1
}

# 使用示例
fetch_with_retry "https://push2his.eastmoney.com/api/qt/stock/kline/get?..."
```

## 相关资源

- **AKShare 官方文档**: https://akshare.akfamily.xyz/
- **GitHub**: https://github.com/akfamily/akshare
- **问题反馈**: https://github.com/akfamily/akshare/issues

## 许可声明

1. 所有数据仅供学术研究使用
2. 数据仅供参考，不构成任何投资建议
3. 投资者应自行承担投资风险
4. 遵循 MIT 开源协议
