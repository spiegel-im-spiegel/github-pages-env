+++
title = "スマホで高解像度音源ファイルを鳴らしてみる"
date =  "2024-03-05T12:46:15+09:00"
description = "違いが分からん orz"
image = "/images/attention/kitten.jpg"
tags = [ "music", "k-tai" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[前回]({{< relref "./ktai-and-music.md" >}} "スマホと音楽")の続き。

スマホの [Onkyo HF Player] アプリと USB-DAC + BOSE 卓上スピーカーの構成で高解像度版の音源ファイルを鳴らしたらどうなるか気になったので，以下の2つの方法を試してみた。

1. [mora] で購入した音源ファイルと Amazon で購入した MP3 のファイルを比較してみる
2. [mora] で購入した音源ファイルと CD からリッピングしたファイルを比較してみる

## 高解像度版（High Resolution）音源ファイル

知ってる人には今更な話だが，音楽データの場合 CD が基準になっていてそれ以上の解像度のデータを「ハイレゾ（High Resolution）」と呼んでいるらしい。
解像度を示すパラメータは2つあって，ひとつはサンプリングレート，もうひとつは量子化ビット数である[^pcm1]。
ものすごく端折って言うと，サンプリングレートは時間に対する分解能で，量子化ビット数はダイナミックレンジに対する分解能と考えてよい。
CD の場合は44.1kHz/16bitの解像度である。
そして，サンプリングレートまたは量子化ビット数がCDより大きければ「ハイレゾ」なんだそうな。

[^pcm1]: この記事では PCM (Pulse Code Modulation) を前提に書いている。これとは別に DSD (Direct Stream Digial) というのもあるのだが，仕組みが全然違うので割愛する。 DSD の音源ファイルは `.dff` とか `.dsf` とかいった拡張子になることが多い。

実は「ハイレゾ」か否かの指標にはもうひとつあって，それがデータの圧縮方法である。
MP3 ファイルは非可逆圧縮方式なので元に戻した際にどうしても情報落ちが発生する。
このため MP3 ファイルは数値上の解像度に関わらず「ハイレゾ」とは見なされないようだ。

一方，可逆圧縮方式で普及している（大抵のメジャーな音楽配信サービスで対応している）のが FLAC (Free Lossless Audio Codec) である。
つまり FLAC ファイルで44.1kHz/16bit以上の解像度であれば「ハイレゾ」と言えるわけだ。

とまぁ，話を単純化したところで，早速比べてみよう。

## アップサンプリング

実は少し前に MindaRyn さんの「[HIBANA](https://www.amazon.co.jp/dp/B0CPM8V21D?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)」を Amazon で購入したのだが，同じタイトルを [mora] で購入してみた。

- [HIBANA／MindaRyn｜音楽ダウンロード・音楽配信サイト　mora ～WALKMAN®公式ミュージックストア～](https://mora.jp/package/43000100/LAXX24518B00Z_96/)

ここで買える音源ファイルは FLAC で解像度は 96.0kHz/24bit である。
さっそく聴き比べてみたのだが... 違いが分からん `orz`

実は [Onkyo HF Player] 有料版の機能に「アップサンプリング機能」ってのがあって USB-DAC を繋いだ状態だと最大で384kHzまでレートを上げてくれる。
実際に Amazon で買った MP3 版は352.8kHz， [mora] で買った FLAC 版は384kHzになっていた。
ここまでレートを上げると量子化ビット数はあまり問題にならなくなる。
私ごときの耳では聴き分けられないわけだ。

## 「組曲 惑星」をリッピングしてみる

今度は，以前に CD で買った「[組曲 惑星](https://www.amazon.co.jp/dp/B0009N2VDM?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)（小澤征爾さんの指揮）」をリッピングして FLAC ファイルに落としてみる。
Ubuntu では昔に[紹介]({{< ref "/remark/2019/06/ripping-cd-music-in-ubuntu.md" >}} "Ubuntu で音楽 CD のリッピング")した [Asunder](http://www.littlesvr.ca/asunder/) が FLAC に対応している。

{{< fig-img src="./asunder-settings.png" title="Asunder settings" link="asunder-settings.png" >}}

早速これを使って FLAC ファイルを作成してみる。
でも，これって CD 以上の解像度にはならないよね，多分。

[mora] でも「組曲 惑星」を探してみたのだが，こちらは小澤征爾さんの指揮によるものはなくて，以下のものを購入した。

- [ホルスト: 組曲「惑星」/神秘のトランペッター Op. 18(スコティッシュ・ナショナル管/ロイド=ジョーンズ)／ロイヤル・スコティッシュ・ナショナル管弦楽団/デイヴィッド・ロイド=ジョーンズ(指揮)｜音楽ダウンロード・音楽配信サイト　mora ～WALKMAN®公式ミュージックストア～](https://mora.jp/package/43000069/8555776h/)

こちらは44.1kHz/24bitの解像度。

[Onkyo HF Player] で聴くと，両者とも352.8kHzにアップサンプリングされた。
やっぱり違いは分からず。

アップサンプリング機能を外して比較することも考えたが，わざとアプリのスペックを落として比較することに意義を感じなかったので止めた。
要するに MP3 だろうが FLAC だろうが [Onkyo HF Player] から USB-DAC 経由で BOSE の卓上スピーカーで聴く限り，少なくとも私の耳では，違いは分からんということだ。

まぁ，でも，折角 [mora] にサインアップしたし，可逆圧縮の FLAC で曲が手に入るなら，これからは FLAC を優先して調達してもいいかなぁ。

## ブックマーク

- [「ハイレゾ」の基本をおさらい。DSDやFLAC、MQAなどの違いとは? - AV Watch](https://av.watch.impress.co.jp/docs/topic/1068831.html)
- [【Ubuntu日和】【第23回】ハイレゾは？Spotifyは？LDACは？Ubuntuで音楽を聴くあれこれ  - PC Watch](https://pc.watch.impress.co.jp/docs/column/ubuntu/1486550.html)
- [Bluetooth SBCコーデックは本当に音質が悪いのか - 謝花ミカ - Medium](https://mikajabana.medium.com/bluetooth-sbc%E3%82%B3%E3%83%BC%E3%83%87%E3%83%83%E3%82%AF%E3%81%AF%E6%9C%AC%E5%BD%93%E3%81%AB%E9%9F%B3%E8%B3%AA%E3%81%8C%E6%82%AA%E3%81%84%E3%81%AE%E3%81%8B-64ef74727bad)

[Onkyo HF Player]: https://www.jp.onkyo.com/support/hfplayer/ "オーディオ&ビジュアル製品情報：Onkyo HF Player"
[mora]: https://mora.jp/ "音楽ダウンロード・音楽配信サイト　mora ～WALKMAN®公式ミュージックストア～"

## 参考

{{% review-paapi "B00CD1PTF0" %}} <!-- BOSE Conpanion 2 -->
{{% review-paapi "B0BL76SJL6" %}} <!-- USB-DAC -->
{{% review-paapi "B09TVLHJ1X" %}} <!-- Shokz OpenRun Mini 骨伝導ヘッドセット -->
{{% review-paapi "B0CPM8V21D" %}} <!-- HIBANA : MindaRyn -->
{{% review-paapi "B0009N2VDM" %}} <!-- ホルスト 組曲 惑星 -->
{{% review-paapi "B09V3D51LF" %}} <!-- ビッグブリッヂの死闘 -->