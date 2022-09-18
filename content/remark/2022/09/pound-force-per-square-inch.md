+++
title = "はじめての空気圧チェック"
date =  "2022-09-17T10:28:18+09:00"
description = "おぅふ！ 思ったより抜けている。"
image = "/remark/2022/09/pound-force-per-square-inch/52363096943_46d213da58_e.jpg"
tags = [ "bike" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

先週，[自転車を受け取った]({{< relref "./x-road-bike.md" >}})際に色々とアドバイスを頂いたのだが（[専門店](https://giant-store.jp/matsue/ "ジャイアントストア松江")はこういうところがありがたい），そのひとつが「通勤で毎日乗るのなら週に1度は空気圧チェックをすること」だった。

というわけで，はじめての空気圧チェック（笑）

{{< fig-img src="./52363096943_46d213da58_e.jpg" title="空気圧チェック | Flickr" link="https://www.flickr.com/photos/spiegel/52363096943/" >}}

おぅふ！ 思ったより抜けている。
確かにこれなら毎週チェックせんとあかんな。

ちなみに今回買った自転車はオフロード寄りのクロスバイクで太いタイヤで空気圧指定も高めになっているそうな。
タイヤの刻印をみると 50-85 psi (pound-force per square inch) とあるが，店員さんには「上限ギリギリまで入れてはいけない」とも言われているので 75 psi で調整[^psi1]（これも指導してもらった）。

[^psi1]: “psi” は「ぴーえすあい」呼びでいいらしい。「ぷさい」と呼ぶ人もいるそうだがギリシャ文字の $\psi$ とは関係ないようだ（笑）

ちなみに

{{< div-box >}}
$$
  1\,\mathrm{atm} = 1.01325\,\mathrm{bar} = 101,325\,\mathrm{Pa}
$$
{{< /div-box >}}

で[^atm]，さらに

[^atm]: “atm” は標準大気圧（standard atmosphere）を 1 とした単位。 SI 単位系では $1\\,\mathrm{atm}=101,325\\,\mathrm{Pa}$ と**定義**されている。“bar” は「ばーる」と読む。大昔の天気予報で気圧を mb (ミリバール) で表記していたあの「バール」である。本文を見れば分かる通り SI 単位系では $1\\,\mathrm{bar}=100,000\\,\mathrm{Pa}$ で換算できる。ちなみに $\mathrm{Pa}$ は $\mathrm{N}/\mathrm{m}^2$ または $\mathrm{kg}\\,\mathrm{m}^{-1}\\,\mathrm{s}^{-2}$ と等価。

{{< div-box >}}
$$
  7\,\mathrm{bar} = 101.526\,\mathrm{psi}
$$
{{< /div-box >}}

なんだそうな（空気入れのメータを見てもそんな感じ）。
つまり $75\\,\mathrm{psi}$ は概ね $5.2\\,\mathrm{bar}=5.2\times10^2\\,\mathrm{kPa}$ ($5.1\\,\mathrm{atm}$) ってところか。

実はフレンチ・バルブ自体がはじめての経験で，空気入れも専用のものを購入している。
店頭ハンズオンでも一度試させてもらっているのだが，結構気を使うのな（乱暴に扱うとバルブが折れるらしい。バルブが折れたら全取っ替え）。
これも慣れないとなー。

## ブックマーク

- [自転車タイヤの空気圧単位 最も使われている単位と変換方法について解説します | CBN Blog](https://blog.cbnanashi.net/2019/01/6564)

## 参考図書

{{% review-paapi "4621304259" %}} <!-- 理科年表 2020 -->
