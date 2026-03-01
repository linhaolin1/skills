# AKShare 财经数据 HTTP API

基于 AKShare 实际数据源的 HTTP API 接口文档，支持通过 curl 直接调用。

## 快速开始

```bash
# 查询股票历史数据
curl -s "https://push2his.eastmoney.com/api/qt/stock/kline/get?fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f116&ut=7eea3edcaed734bea9cbfc24409ed989&klt=101&fqt=1&secid=0.000657&beg=20240101&end=20241231"
```

## 文档结构

- [SKILL.md](SKILL.md) - 完整的技能文档和使用指南
- [rules/stock.md](rules/stock.md) - 股票数据接口
- [rules/futures.md](rules/futures.md) - 期货数据接口
- [rules/fund.md](rules/fund.md) - 基金数据接口
- [rules/bond.md](rules/bond.md) - 债券数据接口
- [rules/index.md](rules/index.md) - 指数数据接口
- [rules/forex.md](rules/forex.md) - 外汇数据接口
- [rules/option.md](rules/option.md) - 期权数据接口
- [rules/macro.md](rules/macro.md) - 宏观经济数据接口
- [rules/crypto.md](rules/crypto.md) - 加密货币数据接口
- [rules/spot.md](rules/spot.md) - 现货数据接口
- [rules/others.md](rules/others.md) - 其他数据接口

## 主要特点

- ✅ 直接使用 curl 调用，无需安装任何库
- ✅ 基于可靠的公开数据源（东方财富网等）
- ✅ 覆盖股票、期货、基金、债券等多个市场
- ✅ 提供实时行情和历史数据

## 使用示例

### 查询股票行情

```bash
# 中钨高新(000657)最近行情
curl -s "https://push2his.eastmoney.com/api/qt/stock/kline/get?fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f116&ut=7eea3edcaed734bea9cbfc24409ed989&klt=101&fqt=1&secid=0.000657&beg=20260220&end=20260301"
```

### 查询实时行情

```bash
# 所有A股实时行情
curl -s "https://82.push2.eastmoney.com/api/qt/clist/get?pn=1&pz=5000&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f12&fs=m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23,m:0+t:81+s:2048&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23"
```

## 注意事项

1. 所有接口均为公开数据源，无需认证
2. 建议合理控制请求频率
3. 数据仅供学术研究，不构成投资建议
4. 股市有风险，投资需谨慎

## 相关链接

- [AKShare 官方文档](https://akshare.akfamily.xyz/)
- [GitHub 仓库](https://github.com/akfamily/akshare)
