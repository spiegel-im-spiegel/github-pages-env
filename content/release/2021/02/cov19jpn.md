+++
title = "手遊びで日本版 Google COVID-19 Forecast データを取得するツールを作ってみた"
date =  "2021-02-14T19:38:03+09:00"
description = "大した機能はないが，よろしければどうぞ。"
image = "/images/attention/tools.png"
tags  = [ "tools", "golang", "infection" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

天下無敵の無職の頃は自分の住んでる地域以外はあまり気にする必要がなかったが，仕事をするようになれば客先地域の感染状況とかも気にしないといけない。
ただ，公表されている情報に関しては自治体によってバラバラで統一性がなく，国レベルで整理して公表しているというわけでもないっぽい？

で，軽くググってみたが Google が公開している日本版 COVID-19 Forecast データが使いやすいようだ。

- [日本版 Google COVID-19 Forecast データを眺める](https://zenn.dev/spiegel/scraps/e992be8b03eeb7)

予測値はともかく，実績値だけでもこのレベルで国が取り扱ってくれんもんかねぇ。

愚痴はともかく，このデータを使って簡単なツールを作ってみた。

- [spiegel-im-spiegel/cov19jpn: COVID-2019 in Japan; Importing Google COVID-19 Public Forecasts](https://github.com/spiegel-im-spiegel/cov19jpn)

たとえばコマンドラインで

```text
$ cov19jpn plot tokyo
```

とかやれば

{{< fig-img src="./tokyo-cov19-chart.png" link="./tokyo-cov19-chart.png" title="tokyo-cov19-chart.png" >}}

てな感じに棒グラフを出力してくれる。
ちなみに7日単位の集計で，過去4週間分の実績値と，それに続く1週間分の予測値を表示している。
個人的には AI 予測は「当たるも八卦」とあまり信用してないが，1週間分なら多少は参考になるかもしれない。

なお，コマンドラインの都道府県名を省略すると47都道府県のグラフをそれぞれ出力するのでご注意を（笑）
    
あと，

```text
$ cov19jpn csv tokyo
japan_prefecture_code,prefecture_name,target_prediction_date,cumulative_confirmed,cumulative_confirmed_q0025,cumulative_confirmed_q0975,cumulative_deaths,cumulative_deaths_q0025,cumulative_deaths_q0975,hospitalized_patients,hospitalized_patients_q0025,hospitalized_patients_q0975,recovered,recovered_q0025,recovered_q0975,cumulative_confirmed_ground_truth,cumulative_deaths_ground_truth,hospitalized_patients_ground_truth,recovered_ground_truth,forecast_date,new_deaths,new_confirmed,new_deaths_ground_truth,new_confirmed_ground_truth,prefecture_name_kanji
JP-13,TOKYO,2021-01-15,,,,,,,,,,,,,82069,717,19797,61555,2021-02-11,,,,,東京都
JP-13,TOKYO,2021-01-16,,,,,,,,,,,,,83878,720,19879,63279,2021-02-11,,,3,1809,東京都
JP-13,TOKYO,2021-01-17,,,,,,,,,,,,,85470,725,20646,64099,2021-02-11,,,5,1592,東京都
JP-13,TOKYO,2021-01-18,,,,,,,,,,,,,86674,728,20777,65169,2021-02-11,,,3,1204,東京都
JP-13,TOKYO,2021-01-19,,,,,,,,,,,,,87914,744,20289,66881,2021-02-11,,,16,1240,東京都
...
```

などと都道府県単位で絞って CSV 出力することも可能。

手遊びで作ったものなので大した機能はないが，よろしければどうぞ。

## ブックマーク

- [WHO COVID-2019 データを取得するパッケージを作ってみた]({{< ref "/release/2020/09/cov19data-package-is-released.md" >}})

<!-- eof -->
