+++
title = "Garmin Edge Explore 2 でルート探索"
date =  "2024-08-14T00:25:03+09:00"
description = "ルート探索・ナビゲーション機能特化。タッチパネル操作ができて Garmin Edge 540 よりちょっとだけ安い。"
image = "/remark/2024/08/garmin-edge-explore-2/53920400824_a5bd46f3e6_o.jpg"
tags = [ "osanpo-camera", "photography", "matsue", "bike", "tools", "gear" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

以前にサイコンを [Garmin 製に買い換えた]({{< ref "/remark/2024/06/cyclocomputer-by-garmin.md" >}} "サイクルコンピュータを買い換えた")話をしたが，このとき購入した Garmin Edge 130 plus はナビゲーション機能が貧弱なため，ナビに関しては[復活した trimm ROLLIN]({{< ref "/remark/2024/07/20-osanpo-camera.md" >}} "お散歩カメラ 2024-07-20") を使っていた。

ただ trimm ROLLIN は [Strava] と相性が悪くてですね。
[Strava] で作ったルートを trimm ROLLIN にインポートしてもまともにナビゲーションしてくれないのですよ。

というわけで，年末にも Garmin Edge の上位機種に買い換えようかなぁ，とぼんやり考えていたのだが，恐ろしい動画記事を見てしまった。

{{< fig-youtube id="njP7a-txTnk" title="【タイプ別】おすすめのGarmin教えます - YouTube" >}}

は？ Garmin Edge シリーズって盆明けに値上げするの？ えー，今のうちに買っとかなきゃ（←扇動される消費者）。

Garmin Edge は，最安の 130 plus の他に500シリーズ（530, 540），800シリーズ（840），1000シリーズ（1040, 1050）とあって，それとは別に [Explore 2][Garmin Edge Explore 2] というルート探索・ナビゲーション機能に特化した製品もあるらしい。
個人的にはガチのトレーニングをするつもりはないので [Garmin Edge Explore 2] でええんちゃうん？ というわけで発注した。

{{% review-paapi "B0BD7FGVR6" %}} <!-- GARMIN EDGE Explore 2 サイクルコンピュータ -->

## Garmin Edge Explore 2 を使ってみる

### セットアップで往生する

早速来たので開梱してセットアップしようとしたが，完全放電しとるがな。
充電するところからかい！ 充電用のポートが USB-C なのは助かるが，残念ながら Power Delivery ではないようだ。

ある程度充電したら起動したぽいので，スマホアプリの [Garmin Connect] とペアリングする。
このとき 130 plus に登録してたセンサのペアリング情報を [Explore 2][Garmin Edge Explore 2] にコピーしてくれた。
マジ助かる。
でもこのあとソフトウェアの更新が始まり，完全に使えるようになったのは数時間後だった。
こういうのは出荷時にベンダ側でやっておいてほしい `orz`

### サイズ

大きさの比較はこんな感じ。

{{< fig-img src="./53920552964_f797dfdee7_e.jpg" title="サイコンサイズ比較（真ん中は6インチスマホ） | Flickr" link="https://www.flickr.com/photos/spiegel/53920552964/" >}}

左が 130 plus，右が [Explore 2][Garmin Edge Explore 2]，真ん中が Galaxy S22 (6.1インチスマホ) である。
ちなみに [Explore 2][Garmin Edge Explore 2] はスマホの倍以上の厚みがある。

### 機能上の特徴

最初に挙げた動画記事では [Explore 2][Garmin Edge Explore 2] の機能として 130 plus レベルのトレーニング機能とナビゲーション機能を併せたようなものと述べていたが，正直に言ってトレーニング機能は 130 plus よりも更に劣る。

まずライドタイプが「ロード」「室内」「グラベル」の3つしかなく，変更も追加もできない。
たとえば「通勤」とかは選べない。

また [Garmin Connect] で作成したワークアウトを利用することができない。
私は心臓リハビリ用に「ウォームアップ30秒＋バイク30分（ケイデンス70〜80の範囲を外れると警告）＋クールダウン1分」みたいな感じでワークアウトを定義して[フィットネスバイク]({{< ref "/remark/2023/01/stationary-bike.md" >}} "フィットネスバイク買うた")で使っているが [Explore 2][Garmin Edge Explore 2] でこの定義を使うことはできなかった。

Live Segment に対応していない。
Garmin Edge では [Strava] と [Garmin Connect] を連携させて「セグメント」ルート走行中の状態（ラップタイムや速度など）をリアルタイムでモニタできるのだが [Explore 2][Garmin Edge Explore 2] だけがこれに対応していなかった。
まぁ，リアルタイムでモニタできなくても記録はちゃんと残るので致命的な問題ではないのだが。

一方で [Explore 2][Garmin Edge Explore 2] の特徴のひとつはディスプレイがタッチパネルになっていることである。
このため物理ボタンは最小限になっている。
他の Garmin Edge 製品でタッチパネルを備えているのは800シリーズと1000シリーズで，500シリーズはタッチパネルじゃない。
これで500シリーズより（ちょっとだけ）安いんだぜ。

ナビ機能は流石によくできている。
タッチパネルなので指2本のピンチ操作で地図のズームが出来る。
地図上の建物とかの情報が古い気がするが（既に潰れたコンビニとか載ってた）致命的ではないし，いいか。

地図の詳細度を指定できるが，これを「最高」にすると街中では使い物にならないくらいゴチャゴチャした見た目になった。
特に松江市の中心部は（城下町だったせいか）古い路地と新しく整備された道路が錯綜していてホンマに分かりにくいのだ。
逆に中心部から外れたところは何もなさすぎて地図に何も表示されなくなる。
熊野大社に行く途中の道とか一本道みたいになっていた。
ままならないものである（笑） なお，ナビゲーション中に地図の詳細度を変えると地図表示がおかしくなるみたい。
気ぃつけなはれや！

スマホの [Garmin Connect] と連携させてナビゲーション中の（右に曲がれとか左に曲がれとかの）指示を音声でアナウンスできる（声はスマホから鳴る）。
サイコンの画面ばかり見て走るわけに行かないので音声アナウンス機能はありがたい。
ただ日本語がカタコトで聞き取りにくい。
この辺は trimm ROLLIN はよくできてるよなぁ。

トレーニング機能とナビゲーション機能の両方が欲しいなら800シリーズまたは1000シリーズを強くおすすめする。
500シリーズは中途半端な感じ。
やっぱタッチパネルは要るですよ。

私は，当面の間，平日の通勤や心臓リハビリでは 130 plus を，週末のサイクリングでは [Explore 2][Garmin Edge Explore 2] を，という感じに使い分けることにした。
まぁ，ちょっと高めのお勉強代と割り切るか。

### パワーマウントを使った給電

今回買った[セット][Garmin Edge Explore 2]にはパワーマウントが付いている。
これは eBike からの給電を想定していて，走行中も充電しながらサイコンを使うことが出来る。

別売りの [USB-A への変換ケーブル](https://www.amazon.co.jp/dp/B0BD6SH1LH?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Amazon | ガーミン(GARMIN) GARMIN(ガーミン) Edge PowerMountケーブル USB-A 400mm【日本正規品】 | ガーミン(GARMIN) | サイクルコンピューター・GPS用アクセサリ")を使うとモバイルバッテリ（USB-A の口がついてること）に接続して充電することが出来る。
こんな感じ。

{{< fig-img src="./53920400824_31c56fcbdc_e.jpg" title="パワーマウント経由でモバイルバッテリと接続（仮） | Flickr" link="https://www.flickr.com/photos/spiegel/53920400824/" >}}

ちょうど手元にコンパクトサイズのモバイルバッテリが余ってたので利用する。
モバイルバッテリはフレームバッグに格納して，バッグのケーブル穴からケーブルを通せば配線がスッキリする。

モバイルバッテリを繋いだ状態で近所を走ってみたが [Explore 2][Garmin Edge Explore 2] を使いつつバッテリの充電状態もちゃんと回復していた。
これならサイコンのバッテリ状態を気にしなくて済むな。
よーし，うむうむ，よーし。

ちなみに[拡張バッテリ](https://www.amazon.co.jp/dp/B077CZ31V3?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Amazon | Garmin(ガーミン) Edge1030用拡張バッテリーパック 010-12562-30【GARMIN純正品】 | ガーミン(GARMIN) | サイクルコンピューター・GPS用アクセサリ")も[使える](https://www.riteway-jp.com/itemblog/%E3%83%96%E3%83%A9%E3%83%B3%E3%83%89-44966/2022/11/_kamata "GARMIN EXPLORE2に付属してくるパワーマウントに拡張バッテリーをつけることができるのかどうか。")らしい。
でも，なんか不格好だな。
私はいいや（笑）

## 縁結び大橋を渡ってみたい

さて，サイコンも新しくなったし，今まで行ったことのない場所へ行ってみよう。
以前から考えていたが，広島在住の頃にできたという縁結び大橋を自転車で渡れると聞いたので，この機会に渡ってみたい。
今回のルートは松江城前を起点としてこんな感じに組んでみた。

{{< fig-img-quote src="./cycling-route.png" title="縁結び大橋を渡りたい" link="https://www.strava.com/routes/3258147846190661044" width="874" >}}

大雑把に

1. 松江城前
2. くにびき大橋南下
3. 縁結び大橋北上
4. 楽山公園
5. 島大通り
6. 惣門橋通り
6. 松江城前

という巡回ルート。
早速行ってみよう。

{{< fig-img src="./53919972534_491b398ec7_e.jpg" title="自転車で縁結び大橋に上がる | Flickr" link="https://www.flickr.com/photos/spiegel/53919972534/" >}}

おお。
ここから入るのか。
では！

{{< fig-img src="./53919618611_3a10b57471_e.jpg" title="縁結び大橋 歩行者自転車道 | Flickr" link="https://www.flickr.com/photos/spiegel/53919618611/" >}}

{{< fig-img src="./53918733097_6732ba8acf_e.jpg" title="縁結び大橋 歩行者自転車道 | Flickr" link="https://www.flickr.com/photos/spiegel/53918733097/" >}}

橋からの眺めはこんな感じ。

{{< fig-img src="./53918734187_29a338ccf5_e.jpg" title="大橋川 from 縁結び大橋 | Flickr" link="https://www.flickr.com/photos/spiegel/53918734187/" >}}

おお！ これはなかなか絶景だ。

{{< fig-img src="./53918735077_12c0e2545b_e.jpg" title="縁結び大橋 歩行者自転車道 | Flickr" link="https://www.flickr.com/photos/spiegel/53918735077/" >}}

勾配はさほどキツくなく，私の貧弱な脚力でも問題なく走れた。
これは楽しい。

松江市の東は私の（遊び場としての）テリトリから外れているので知らない場所が多いんだよな。
[Strava] と [Explore 2][Garmin Edge Explore 2] を組み合わせて，これからも楽しくサイクリングとお散歩カメラを続けよう。

## ブックマーク

- [Strava サブスクリプションに加入した（お散歩カメラ 2024-08-03）]({{< relref "./03-osanpo-camera.md" >}})

[Strava]: https://www.strava.com/ "Strava | アプリで簡単ラン、サイクリング、ハイキング - トレーニングの結果を追跡・シェア"
[Garmin Edge Explore 2]: https://www.amazon.co.jp/dp/B0BD7FGVR6?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Amazon | GARMIN(ガーミン)Edge Explore 2 Power サイクルコンピューター【日本正規品】 | ガーミン(GARMIN) | サイクルコンピューター"
[Garmin Connect]: https://connect.garmin.com/ "Garmin Connect"
[熊野大社]: http://www.kumanotaisha.or.jp/ "出雲國一之宮　熊野大社"

## 参考

{{% review-paapi "B0BD6SH1LH" %}} <!-- Garmin Edge Explore 2 パワーマウント USB変換ケーブル -->
{{% review-paapi "B0C5HLDYSM" %}} <!-- モバイルバッテリ（小） -->
{{% review-paapi "B08L4WKDZ7" %}} <!-- PowerShot ZOOM -->
{{% review-paapi "B08BZ5T9NZ" %}} <!-- GARMIN EDGE 130 PLUS サイクルコンピュータ -->
{{% review-paapi "B09XGYX7JF" %}} <!-- GARMIN vívosmart 5 -->
{{% review-paapi "B09TVLHJ1X" %}} <!-- Shokz OpenRun Mini 骨伝導ヘッドセット -->
{{% review-paapi "B0B8D1S61W" %}} <!-- 仮面ライダーW 風都探偵 Let’s go ahead -->
{{% review-paapi "B0DB5N3HJL" %}} <!-- ダンジョンの中の人 マイクロレボリューション TrySail -->
