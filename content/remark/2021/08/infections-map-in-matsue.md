+++
title = "松江市における SARS-CoV-2 感染者マップを描いてみた"
date =  "2021-08-29T13:20:04+09:00"
description = "こういったグラフを使ってクラスタ型感染の兆候・推移が可視化できればいいかな，と思ったのだ。"
image = "/images/attention/kitten.jpg"
tags = [ "infection" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

昨日ふと思い付いて，松江市における SARS-CoV-2 感染者マップを描いてみた。
こんな感じ。

{{< fig-img src="./infections.png" title="covid-2019-report/report/matsue at master · spiegel-im-spiegel/covid-2019-report" link="https://github.com/spiegel-im-spiegel/covid-2019-report/tree/master/report/matsue" width="2219" >}}

内訳はこんな感じ。

- 2020年10月25日（125例目）からの[報告](http://www1.city.matsue.shimane.jp/kenkou/kenkoudukuri/kansensyo_yobou/coronavirus-disease/coronahasseijoukyou.html "松江市:暮らしのガイド:発生状況")をプロットしている
- 黒丸は松江市内在住者
- 赤丸は松江市外在住者または松江市外在住者からの二次感染者
- 背景が濃いものは直近1週間以内の松江市内在住の感染者

松江市は（昨年夏の集団感染を除けば）1日あたり（多いときでも）20人ほどの感染規模である。
たとえばこれを東京や大阪のような大都市でやっても，ゴチャゴチャして意味不明のグラフになるだけだろうが，いい感じに過疎ってる松江ならではというところか（笑）

見方を変えると松江市では（母集団が小さいので）量（＝人数）による統計評価が難しい。
なので，こういったグラフを使ってクラスタ型感染の兆候・推移が可視化できればいいかな，と思ったのだ。

何となく使えそうな気がするので，情報収集用に運用している [spiegel-im-spiegel/covid-2019-report](https://github.com/spiegel-im-spiegel/covid-2019-report "spiegel-im-spiegel/covid-2019-report: 日本における COVID-2019 確認発症者のレポート") のワークフローに組み込んでみることにした。
といっても毎日更新する気はない。
いつものように週末にサクっとまとめる程度である。
専門家じゃないんだから，大体の傾向が分かればいいのよ。

「あんしん」に耽溺することなく，取るべきリスクは取って，その上で安全に暮らしましょう。
