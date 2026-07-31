+++
title = "Anker Smart Pouch と USB Cable Checker 2 買うた"
date =  "2024-11-16T23:29:41+09:00"
description = "ノートPCの付属品をまとめるポーチ / 古い規格の USB ケーブルは断捨離だw"
image = "/images/attention/kitten.jpg"
tags = [ "tools", "gear" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

あー。また散財しちまったよ。
でも反省はしない。

## Anker Smart Pouch Supported by KOKUYO

ノートPCを外に持ち出そうとするとケーブルとか充電器とかマウスとか付属品も一緒に持ち出す必要があるぢゃん。
今までは巾着袋に入れてたんだけど，中でぐちゃぐちゃになるんだよね，当たり前だけど。
ノートPCの付属品を上手くまとめることができてバッグ・イン・バッグにできるものないかなぁ，と探してたんだけど，よさげなものがあった。

{{% review-paapi "B0D8TJJGNJ" %}} <!-- Anker Smart Pouch Supported by KOKUYO ポーチ -->

見た目より収納できる。
ケーブルと充電器と（外部モニタとつなぐための）メディアハブとモバイルバッテリと薄型マウスが余裕で入った。
この前買った Rollbahn × HAYAKAWA FACTORY 製[ホームズ柄のノート](https://www.amazon.co.jp/dp/B0CGZLNVD5?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Amazon.co.jp: 【HAYAKAWA FACTORY】デルフォニックス ロルバーン ポケット付メモL ホームズ チェック : 文房具・オフィス用品")も入ったし，フリクションペンを挿すところもある。

これでこの前買った[スリングバッグ]({{< ref "/remark/2024/09/chrome-kadet-max.md" >}} "スリングバッグ買うた")や仕事用の[バックパック]({{< ref "/remark/2023/10/my-new-gear.md" >}} "バックパックを買い替えた")に問題なく入る。
ホンマ助かる。

## USB Cable Checker 2 買うた

以下の動画を見て

{{< fig-youtube id="nfLPE12Ewd4" title="ケーブルチェッカーで調べたらゴミUSB-Cケーブルが大量に出てきた。。。 - YouTube">}}

私も真似して自宅にある USB ケーブルの山を整理することにした。
さっそく件の USB ケーブルチェッカを購入。

{{% review-paapi "B07Y8BPVV4" %}} <!-- USB Cable Checker 2 ケーブルチェッカ -->

今回は Amazon のマーケットプレイスを利用したんだけどさ， Amazon で「USB Cable Checker 2」で検索すると「保護フィルム」なる謎の製品がやたらとヒットするのよ。
で，肝心の製品は検索結果のずうっと下のほうにあったり。
アカンやろこれ。
そういうとこやぞ  Amazon！

心配な方は[楽天で買える](https://item.rakuten.co.jp/r-kojima/4562469772981/?scid=af_pc_etc&sc2id=af_101_0_0 "【楽天市場】ビットトレードワン　USBケーブルの性能を確認できる検証デバイス USB CABLE CHECKER 2　ADUSBCIM：コジマ楽天市場店")らしいのでそっちでどうぞ。
私は楽天のアカウントを持ってないので関係ないけど。

USB 3.2 まで対応している。
公式マニュアルが GitHub の[リポジトリ](https://github.com/bit-trade-one/ADUSBCIM-USBCableChecker2 "bit-trade-one/ADUSBCIM-USBCableChecker2: ADUSBCIM")にある。

トラブルもなく無事に届いたので，さっそく充電に使ってるケーブルで試してみる。

{{< fig-img src="./54143133076_bc2b96eb1b_e.jpg" title="USBケーブルチェッカ買うた | Flickr" link="https://www.flickr.com/photos/spiegel/54143133076/" >}}

CC (Configuration Channel) ワイヤが有効になってるので PD (Power Delivery) 対応であることが分かる。
まぁ充電用ケーブルとしては十分だろう。

USB-A コネクタがあるケーブルはいいのよ。
USB-A は USB 2 と USB 3 で色違いになっていて分かりやすいから。
いい機会なので USB-A で USB 3 に対応していないケーブルは（使用中のもの以外）全部廃棄することにした。
スッキリ！

問題は USB-C ⇔ USB-C ケーブル。
同じ見た目なのに色々ありすぎてワケワカメだよ。

手持ちのケーブルで100W給電に対応しているケーブルは2本しかなかった。

{{< fig-img src="./54143782045_0b6a611fe4_e.jpg" title="おっ、E Marker IC 内蔵ケーブルだった | Flickr" link="https://www.flickr.com/photos/spiegel/54143782045/" >}}

この2本はとっておくとして CC ワイヤが有効でないケーブルは全部廃棄。
残りのケーブルのうち短めのケーブルのみ予備として残した。

D ワイヤが有効でないケーブルは流石になかったが USB-C ⇔ USB-C ケーブルで TX1/RX1, TX2/RX2, SBU1/SBU2 (Side Band Use; 音声や映像をやり取りするためのワイヤ) が有効なケーブルは1本もなかった。
ありゃりゃー。
1本くらいフルスペックのケーブルを買っておくか。

これからお金を出して買うなら Thunderbolt 4 ケーブルを買うといいらしい。
Thunderbolt 4 なら USB 4 の仕様を内包しているので安心である。
その代わりお高いんだけど。
Thunderbolt 規格のケーブルはコネクタ部分に「{{% emoji "雷" %}} 4」みたいな刻印があるので比較的分かりやすい。

あとは電気製品を買ったときにオマケでついてくる USB ケーブルは，今回買ったチェッカで調べて古い規格のものは積極的に廃棄だな。

## ブックマーク

- [コクヨ×アンカーの「Anker Smart Pouch」でテック小物の持ち運びを解決【いつモノコト】-Impress Watch](https://www.watch.impress.co.jp/docs/series/itsmo/1635647.html)
- [身近なのに複雑怪奇、USB規格のおさらい【第3回】USB Type-Cケーブル編 | 株式会社アスク](https://www.ask-corp.jp/guide/usb-part3.html)
- [USB Type-C - Wikipedia](https://ja.wikipedia.org/wiki/USB_Type-C)
- [フリクションボールノックゾーンはなかなか良い – mhatta's mumbo jumbo](https://www.mhatta.org/wp/2023/01/07/frixion-ball-knock-zone-is-good/)

## 参考

{{% review-paapi "B0BQQMVWBB" %}} <!-- ボディバッグ スリングバッグ CHROME KADET MAX -->
{{% review-paapi "B07TXY9KBJ" %}} <!-- バックパック CHROME -->
{{% review-paapi "B07KJWYQJW" %}} <!-- ANKER PowerExpand USB メディアハブ -->
{{% review-paapi "B088R6SV4Z" %}} <!-- ANKER 充電器 Power Delivery (PD) 対応。65W -->
{{% review-paapi "B093L6298K" %}} <!-- Anker USB-C ⇔ USB-C ケーブル 100W対応 PD 対応 -->
{{% review-paapi "B095RZKWRV" %}} <!-- Anker USB-C & Thunderbolt 4 ケーブル -->
{{% review-paapi "B08LMYWKZD" %}} <!-- Bluetooth 無線静音マウス -->
{{% review-paapi "B0CGZLNVD5" %}} <!-- Rollbahn HAYAKAWA FACTORY ホームズ メモ ノート -->
{{% review-paapi "B0BQR8RZHB" %}} <!-- パイロット フリクションボールノックゾーン 0.5mm ミッドナイトネイビー -->
{{% review-paapi "B0BQX9QTRH" %}} <!-- パイロット フリクションボールノックゾーン替芯 0.5mm ブルー -->


<!-- eof -->
