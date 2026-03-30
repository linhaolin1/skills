#!/usr/bin/env python3
"""
实时买入建议 - 使用腾讯API
"""
import urllib.request
import ssl
from datetime import datetime
import os
import time

ssl_context = ssl.create_default_context()
ssl_context.check_hostname = False
ssl_context.verify_mode = ssl.CERT_NONE

def fetch_url(url, encoding='gbk'):
    try:
        req = urllib.request.Request(url, headers={
            'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36',
        })
        with urllib.request.urlopen(req, context=ssl_context, timeout=10) as response:
            return response.read().decode(encoding)
    except:
        return None

def get_stock_data(code):
    market = 'sh' if code.startswith('6') else 'sz'
    url = f"http://qt.gtimg.cn/q={market}{code}"
    data = fetch_url(url)
    if data and '~' in data:
        parts = data.split('~')
        if len(parts) > 40:
            return {
                'name': parts[1],
                'price': float(parts[3]),
                'change': float(parts[32]),
                'high': float(parts[33]),
                'low': float(parts[34]),
                'amount': float(parts[37]) * float(parts[3]) * 100,
            }
    return None

stocks = {
    '300750': ('宁德时代', '储能'),
    '601012': ('隆基绿能', '光伏'),
    '300014': ('亿纬锂能', '锂电池'),
    '300760': ('迈瑞医疗', '医药'),
    '600519': ('贵州茅台', '白酒'),
    '688256': ('寒武纪', 'AI芯片'),
    '002594': ('比亚迪', '锂电池'),
    '601398': ('工商银行', '银行'),
}

report_time = datetime.now()
report_time_str = report_time.strftime('%Y-%m-%d %H:%M:%S')
report_date = report_time.strftime('%Y%m%d')
report_hour = report_time.hour

print(f"正在获取实时数据... ({report_time_str})")

results = []
for code, (name, industry) in stocks.items():
    data = get_stock_data(code)
    if data:
        results.append({
            'code': code,
            'name': name,
            'industry': industry,
            'price': data['price'],
            'change': data['change'],
            'amount': data['amount'],
        })
        print(f"{name}: {data['price']:.2f}元 ({data['change']:+.2f}%)")
    time.sleep(0.1)

print(f"\n成功获取 {len(results)} 只股票数据")

report = f"""# A股实时买入建议报告

**报告生成时间**: {report_time_str}
**报告日期**: {report_time.strftime('%Y年%m月%d日')}
**当前时段**: {report_hour}:00
**交易状态**: {'交易时段' if 9 <= report_hour < 15 else '非交易时段'}

---

## 一、实时行情数据

| 代码 | 名称 | 行业 | 现价(元) | 涨跌 | 成交额(亿) | 评估 |
|------|------|------|----------|------|------------|------|
"""

for r in sorted(results, key=lambda x: -x['change']):
    amount_yi = r['amount'] / 100000000
    if -3 <= r['change'] <= -1:
       评估 = "回调充分"
    elif -1 < r['change'] < 0:
       评估 = "小幅调整"
    elif 0 <= r['change'] < 3:
       评估 = "温和上涨"
    elif r['change'] >= 3:
       评估 = "强势上涨"
    else:
       评估 = "大幅回调"
    report += f"| {r['code']} | {r['name']} | {r['industry']} | {r['price']:.2f} | {r['change']:+.2f}% | {amount_yi:.2f} | {评估} |\n"

report += f"""
---

## 二、买入建议分析

### 2.1 评分标准

| 条件 | 评分 | 说明 |
|------|------|------|
| 回调-5%~-3% | +30分 | 大幅回调 |
| 回调-3%~-1% | +25分 | 适度回调 |
| 回调-1%~0% | +20分 | 小幅调整 |
| 成交>10亿 | +20分 | 成交活跃 |
| 成长性行业 | +20分 | 光伏/锂电/AI |
| 防御性行业 | +15分 | 医药/白酒/银行 |

### 2.2 综合评分与建议

| 排名 | 代码 | 名称 | 现价 | 涨跌 | 评分 | 建议 | 买入价 | 止损位 | 目标价 |
|------|------|------|------|------|------|------|--------|--------|--------|
"""

scored = []
for r in results:
    score = 0
    reasons = []
    
    if -5 <= r['change'] <= -3:
        score += 30
        reasons.append("大幅回调")
    elif -3 < r['change'] <= -1:
        score += 25
        reasons.append("适度回调")
    elif -1 < r['change'] < 0:
        score += 20
        reasons.append("小幅调整")
    
    if r['amount'] > 10000000000:
        score += 20
        reasons.append("成交活跃")
    
    if r['industry'] in ['光伏', '锂电池', 'AI芯片']:
        score += 20
        reasons.append("成长性行业")
    elif r['industry'] in ['医药', '白酒']:
        score += 15
        reasons.append("防御性行业")
    
    buy_price = r['price'] * 0.99 if r['change'] < 0 else r['price'] * 0.98
    stop_loss = r['price'] * 0.93
    target_price = r['price'] * 1.08
    
    scored.append({
        **r,
        'score': score,
        'reasons': reasons,
        'buy_price': buy_price,
        'stop_loss': stop_loss,
        'target_price': target_price,
    })

scored.sort(key=lambda x: -x['score'])

for i, s in enumerate(scored, 1):
    if s['score'] >= 60:
       建议 = "强烈推荐"
    elif s['score'] >= 50:
       建议 = "推荐买入"
    elif s['score'] >= 40:
       建议 = "可考虑"
    else:
       建议 = "观望"
    report += f"| {i} | {s['code']} | {s['name']} | {s['price']:.2f} | {s['change']:+.2f}% | {s['score']} | {建议} | {s['buy_price']:.2f}元 | {s['stop_loss']:.2f}元 | {s['target_price']:.2f}元 |\n"

report += f"""
---

## 三、重点推荐

"""

top = scored[0] if scored else None
if top:
    report += f"### {top['name']}({top['code']}) - {top['industry']}\n\n"
    report += f"- **评分**: {top['score']}分\n"
    report += f"- **现价**: {top['price']:.2f}元\n"
    report += f"- **涨跌**: {top['change']:+.2f}%\n"
    report += f"- **建议买入价**: **{top['buy_price']:.2f}元**\n"
    report += f"- **止损位**: {top['stop_loss']:.2f}元\n"
    report += f"- **目标价**: {top['target_price']:.2f}元\n"
    report += f"- **推荐理由**: {', '.join(top['reasons'])}\n"

report += f"""
---

## 四、风险提示

1. 市场波动较大，投资需谨慎
2. 建议严格执行止损纪律
3. 仓位控制在30%-50%
4. 数据来源于公开渠道，仅供参考

---

*报告生成时间: {report_time_str}*
"""

output_file = f"/Users/linhaolin1/Documents/code/skills/doc/实时买入建议_{report_date}_{report_hour}时.md"
with open(output_file, 'w', encoding='utf-8') as f:
    f.write(report)

print(f"\n✅ 报告已生成: {output_file}")
