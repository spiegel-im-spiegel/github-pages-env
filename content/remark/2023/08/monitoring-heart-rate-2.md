+++
title = "Garmin vívosmart 5 を買っちまった（心臓リハビリ＠がんばらない 2）"
date =  "2023-08-02T22:04:53+09:00"
description = "スマートバンドをサイクルコンピュータに繋いで心拍数をモニタする。"
image = "/images/attention/kitten.jpg"
tags = [ "bike", "disease", "tools" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

いや，[以前買った Fitbit]({{< ref "/monitoring-heart-rate.md" >}} "心拍数を監視する") が駄目になったわけじゃないのよ。
Fitbit のエコシステムの中で活動するなら，それはそれで必要十分だと思うし[^fb1]。
でも，やっぱり[先日買ったサイクルコンピュータ]({{< ref "/remark/2023/07/cyclocomputer.md">}} "サイクルコンピュータ買うた（お散歩カメラ 2023-07-29）")で心拍数をモニタしたかったのよ。
そうすればアクティビティとして記録を取れるし。

[^fb1]: Fitbit のデバイスは Bluetooth で心拍データを外部システムに出力する仕組みがないそうな。もちろん ANT+ にも対応してないためサイクルコンピュータと連携できない。

心拍数をモニタできてサイクルコンピュータと Bluetooth SMART または ANT+ で連携できるセンサは色々ある。

でも，いくら正確に測れるつっても{{% ruby "ちち" %}}胸{{% /ruby %}}バンド型はちょっと...
あと手首じゃなくて腕に巻くやつとかもあるのね。
[Wahoo の](https://www.amazon.co.jp/dp/B078GRMFSN?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Amazon | Wahoo TICKR FITハートレートアームバンド、Bluetooth/ANT + | WAHOO(ワフー) | 心拍計")とか。
実は殆どこれに決めかけてたんだけど，手首にスマートバンド着けて腕に心拍計着けてって絵面を想像してしまい諦めた。

というわけで，手首に巻くスマートバンド（活動量計）を買い換えることにした。

{{% review-paapi "B09XGYX7JF" %}} <!-- GARMIN vívosmart 5 -->

嗚呼，ついに Garmin 製品に手を出すことになるとは。
でもいっちゃん安いやつね。
というわけで，到着。

{{< fig-img src="./53087395226_949a35cec7_e.jpg" title="ついに Garmin に手を出してしまった | Flickr" link="https://www.flickr.com/photos/spiegel/53087395226/">}}

初期化にはスマホアプリ “[Garmin Connect Mobile](https://www.garmin.co.jp/products/apps/garmin-connect-mobile/ "Garmin Connect™ Mobile App")” が必要。
最近は何でもスマホがないと動かないんだねぇ。
便利だけど不便。

で，アプリを入れただけでは駄目で [Garmin Connect](https://connect.garmin.com/) サービスにサインアップが必要。
アプリをサインアップしてようやく Garmin デバイスとペアリングできた。
これが分からなくてしばらく悩んだんだよ。
ちゃんとマニュアルに書いとけっての！

ひとしきり弄り倒したあと，いよいよサイクルコンピュータに繋いでみる。
どうやら自動で接続できるわけではなく， Garmin デバイスを操作して「心拍転送モード」にしないといけないらしい。

{{< fig-img-quote src="./from-manual.png" title="VÍVOSMART 5 操作マニュアル" link="https://download.garmin.com/jp/download/manuals/vivosmart5_OM_JA.pdf" width="1413" >}}

「心拍転送モード」の間だけペアリングした機器に心拍データを送ってくれる。
ANT+ はごく近傍しか電波が届かないとはいえ，四六時中そこらにデータを撒き散らすわけにはいかないもんな。

実際に「心拍転送モード」にしてサイクルコンピュータに繋いで自転車に乗ってみたが，いい感じにデータが取れた。
うんうん。
これならよかろう。
ただ「心拍転送モード」にすると Garmin デバイスのバッテリ消費が激しいみたい。
この辺はもう少し検証が必要か。

フィットネスバイクでもサイクルコンピュータを使ってケイデンスと心拍データを「記録」することができるようになった。
Garmin デバイス側でもフィットネスバイクを漕いでる間の心拍データ等は取れるのだが，何故かウォーキングとして記録されるんだよなぁ。
まぁ，ええけど。

なんだかんだと散財してるなぁ。
そうでなくとも今月は現金が要るのに。
しばらく大人しくしていよう。

## 今日のお言葉

垂直統合すればベンダーロックインのリスクがある。
分離分散させればサプライチェーンのリスクがある。

## 参考

{{% review-paapi "B0BLNFPWTQ" %}} <!-- trimm ROLLIN サイクルコンピュータ -->
