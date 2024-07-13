+++
title = "お散歩カメラ 2024-07-13"
date =  "2024-07-13T22:29:43+09:00"
description = "暑熱順化 / スマホのナビゲーションアプリを試す / 八雲温泉で暑熱順化だ（笑） / 定点観測 / おつかい"
image = "/remark/2024/07/13-osanpo-camera/53853592160_1679503d14_o.jpg"
tags = [ "osanpo-camera", "photography", "matsue", "bike", "tools", "gear" ]
license = "by"
pageType = "photo"

[scripts]
  mathjax = false
  mermaidjs = false
+++

## 暑熱順化

無教養な私は「暑熱順化」なる言葉を初めて知ったのだが，知ったきっかけがこれ。

{{< fig-youtube id="sO-9_JLEgdc" title="「暑熱順化」で熱中症対策　尾道市消防局「市民も適度な運動を」 - YouTube" >}}

いやぁ，あの装備で階段の上り下りとか腕立て伏せとかないわ（笑） 消防士さんスゲー！

「暑熱順化」については以下のページが参考になるようだ。

- [暑熱順化 | 熱中症ゼロへ - 日本気象協会推進](https://www.netsuzero.jp/learning/le15)

このページの下の方に「暑熱順化チェック」ってのがあって

{{< fig-img-quote src="checking-heat-adaptation.png" title="暑熱順化 | 熱中症ゼロへ - 日本気象協会推進" link="https://www.netsuzero.jp/learning/le15" width="685" >}}

「入浴（シャワーだけでなく、湯船に入るもの）」「運動（汗をかく程度のもの）」「その他の汗をかく行動（運動・入浴以外の外出など）」の3つのうち2つ以上を継続的に行っていれば効果があるらしい。

入浴といえば温泉っスよ。
というわけで，今回も温泉へGo。

## スマホのナビゲーションアプリを試す

[前回] Garmin 先生のナビがイマイチという話を書いたんだけど，今回は音声ナビゲーションが可能なサービスを2つ試してみた。

### [OsmAnd]

[OsmAnd] はマップをダウンロードしてオフラインでナビゲーションが可能なサービスおよびアプリ。
Android 版のアプリは以下で取得できる。

- [OsmAnd —マップと GPS オフライン - Google Play](https://play.google.com/store/apps/details?id=net.osmand)

地図は7つまでは無料でダウンロードできる。
つっても中国地方まるごとひとつの地図データとしてダウンロードできるので，私の行動範囲で7つ使い切ることはなかろう。

{{< fig-img src="./osmand.png" title="OsmAnd" link="./osmand.png" >}}

ルート作成やナビゲーションは無料でできる。
...のだが，何故かやたら裏道を通りたがるのでルート作成はギブアップした。
[Garmin Connect] で作成したルートを GPX 形式のファイルでエクスポートしたものを [OsmAnd] にインポートできるので，今回はそれで。

さっそく[熊野大社]までの道のりを音声案内でナビゲートさせたのだが，すンげぇ分かりにくい。
ただのカーブで左に行けだの右に行けだの行ってくる。
しかも数百メートル手前からアナウンスするので「何言ってるの？」って感じになる。
特に街中の案内がダメダメで，知らない道だと（音声だけでは）確実に迷う。

アナウンスの音が小さい。
スマホの音量を上限近くまで上げないと周囲のノイズに埋もれて聞こえない。
スマホの音量を上げると（音が大きすぎて）音楽が聴けないというのも不便。

### [Ride with GPS]

自転車等のアクティビティの共有と評価を行うサービス。
スマホアプリを使って音声ナビゲーションが可能。

- [Ride with GPS: Bike Navigation - Google Play](https://play.google.com/store/apps/details?id=com.ridewithgps.mobile)

ただし，ルートの保存やナビゲーション機能を使う場合はサブスクリプション契約を行う必要がある。
7日間のお試し期間あり。
今回はお試し期間を利用した。

[Ride with GPS] のルート作成（ルートプランナー）も微妙に使いにくい。
まぁ [OsmAnd] よりはだいぶマシだけど。
[熊野大社]から市内へのルートを [Garmin Connect] で作成し GPX ファイルにエクスポートしたものを [Ride with GPS] に読み込ませることで対応した。
やっぱ [Garmin Connect] のルート作成機能が一番使いやすいな（個人の感想です）。

で，実際に音声案内でナビゲートさせたのだが，これが [OsmAnd] より分かりにくい。
音声で「たんれふと」とか「たんらいと」とか言ってて，最初はなに言ってるか分からなかった。
あぁ “turn left” “turn right” のことか。
英語不得手だと苦労する。
しかもアナウンスで「たんれふと，それからひだりにまがる」とか言うのよ。
2回曲がるのかと思ったよ。

↑のアナウンスが分岐ポイントの数百メートル手前で鳴るのだが，実際にルートの分岐点に到達しても何のアナウンスもなくただジングルが鳴るだけ。
しゃべってくれ！ 何のジングルか分からんじゃろ。

[OsmAnd] と同じで，ただのカーブで左に曲がれとか右に曲がれとか言ってくさる。
そうかと思うと「道なりに行け」と言ってしばらくだんまりだったり。
アナウンスの基準が分からん。
丁字路やY字路でのアナウンスも不親切だし。

[Ride with GPS] もアナウンスの音が小さい。
スマホの音量を上限近くまで上げないと周囲のノイズに埋もれて聞こえない。
しかも通知音を有効にしないとそもそも音声アナウンスが鳴らない。
スマホの通知音を有効にすると要らんジングルまで鳴り始めるので鬱陶しい。

ぶっちゃけ，お金払ってこの品質なら要らん（アクティビティの共有と評価なら [Strava] のほうがよく出来ている）。
というわけでサブスクリプション契約は速攻で解除した（お試し期間中は有効みたい）。

### スマホアプリのナビゲーションは地図が見れてナンボ？

[OsmAnd] も [Ride with GPS] も音声アナウンスのみでナビゲートさせるのは無理があることが分かった。
強いてどっちか選ぶとするなら [OsmAnd] のほうかな。

そう考えると [trimm ROLLIN] のナビゲーション機能はホンマによく出来てる。
つくづく[アプリがぶっ壊れた件]({{< ref "/remark/2024/06/29-osanpo-camera.md#cc" >}} "サイコンの制御アプリが起動しなくなった")は惜しい。
もしアプリが復帰したらサイコン同時使用とかできないだろうか。

スマホを自転車に取り付けるのは抵抗があるんだよな。
でも，もし買うなら [Peak Design](https://jp.peakdesign.com/ "Bags + Camera Gear | Peak Design Official Site") の[アダプタ](https://jp.peakdesign.com/collections/mobile/products/everyday-case-samsung?variant=39603148161101 "Everyday Case for Samsung | Peak Design Official Site")と[マウント](https://jp.peakdesign.com/products/universal-bar-mount "Universal Bar Mount | Peak Design Official Site")かな。

{{< fig-youtube id="-LAfJFocW3A" title="自転車との相性が抜群！カメラマン御用達の peak design のスマホケースがとても素晴らしいのです。 - YouTube" >}}

## 八雲温泉で暑熱順化だ（笑）

気を取り直してお散歩カメラを始めよう。
まずはいつものように[八雲温泉][八雲温泉ゆうあい熊野館]へ。

あれに見えるは白い~~モビルスーツ~~案山子か。

{{< fig-img src="./53853373819_4ed35abf12_e.jpg" title="白の監視者（カカシ） | Flickr" link="https://www.flickr.com/photos/spiegel/53853373819/" >}}

案山子なんて久しぶりに見たな。

[八雲温泉][八雲温泉ゆうあい熊野館]に到着して，まずはサッパリ。

{{< fig-img src="./53853374804_b3216c297c_e.jpg" title="八雲温泉上がりのコーヒー牛乳＆日焼け止め | Flickr" link="https://www.flickr.com/photos/spiegel/53853374804/" >}}

[木次乳業]さん，いつもお世話になっています。
日焼け止めは塗り直さないと。

お昼は何にしようかとメニューを見たら夏限定の盛岡冷麺ってのがあった。
さっそく注文する。

{{< fig-img src="./53853470695_5536e2b744_e.jpg" title="夏季限定盛岡冷麺（熊野館） | Flickr" link="https://www.flickr.com/photos/spiegel/53853470695/" >}}

スープが旨辛で美味しかった。
もっとも私は本場の盛岡冷麺を知らんので，この味で正解なのかは分からない（笑）

腹ごなしというわけではないが，久しぶりに[熊野大社]にも詣でておくか。

{{< fig-img src="./53853157881_182b98909e_e.jpg" title="熊野大社へ参る | Flickr" link="https://www.flickr.com/photos/spiegel/53853157881/" >}}

{{< fig-img src="./53853592160_47f5914307_e.jpg" title="お稲荷さんへも参る | Flickr" link="https://www.flickr.com/photos/spiegel/53853592160/" >}}

最近，ちょっとずつ筋力が付いてきたせいか，うっかりスピードを出し過ぎることがあるんだよね。
田んぼの真ん中の見通しのいい道ならともかく，街中でスピードを出すのは危ない。
実際に最近ヒヤリとしたことがあるし。
というわけで，[熊野大社]では交通安全を誓ってみた。
街中ではギアを落として（ケイデンスは落とさない）安全運転。

## 定点観測

さて腹もこなれたし，市内へとって返そう。
いつもの定点観測。

{{< fig-img src="./53853711670_fe5d83bcca_e.jpg" title="今日の宍道湖 | Flickr" link="https://www.flickr.com/photos/spiegel/53853711670/" >}}

{{< fig-img src="./53853278396_0c916e320a_e.jpg" title="今日の松江城 | Flickr" link="https://www.flickr.com/photos/spiegel/53853278396/" >}}

今日は基本的に曇りで大山も見れなかったんだけど，ときどき日が差すんだよね。
日焼け止めを何度も塗り直したり。

## おつかい

定点観測のあとは[県立図書館][島根県立図書館]へ。

{{< fig-img src="./53853279521_980ffb161a_e.jpg" title="島根県立図書館は外壁工事中（年内） | Flickr" link="https://www.flickr.com/photos/spiegel/53853279521/" >}}

おっ。
なんか足場が組んであるな。
年末まで外壁工事なのかな。

このまま夕方までまったりするつもりだったのだが，実家から電話で「Icoca カード買ってきてくれ」とのリクエスト。
しょうがない。
JR 松江駅横の[テルサ][松江テルサ]まで行くか。

[テルサ][松江テルサ]の2階で Icoca カードを発行してくれる。

{{< fig-img src="./53853763385_ef157908e7_e.jpg" title="松江テルサでおつかい | Flickr" link="https://www.flickr.com/photos/spiegel/53853763385/" >}}

もう[図書館][島根県立図書館]に戻るのも面倒なので駅ビルの喫茶店で涼みつつ夕方まで Kindle 読書をしていた。



## ブックマーク

- [自作したサイクリングルートで音声ナビゲーション！STRAVA+OsmAnd Maps(iOS版) | 東京ライド](https://tokyo-ride.jp/strava_osmand/)

[前回]: {{< ref "/remark/2024/07/06-osanpo-camera.md" >}} "お散歩カメラ 2024-07-06"
[OsmAnd]: https://osmand.net/
[Ride with GPS]: https://ridewithgps.com/
[Strava]: https://www.strava.com/
[Garmin Connect]: https://connect.garmin.com/
[OpenStreetMap]: https://www.openstreetmap.org/
[Ride with GPS]: https://ridewithgps.com/
[trimm ROLLIN]: https://www.amazon.co.jp/dp/B0BLNFPWTQ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Amazon.co.jp: trimm ROLLIN サイクルコンピュータ GPS 自転車 速度計 ワイヤレス ナビゲーション ANT+センサー対応 Bluetooth 心拍数 高度計 2.7インチ スピードセンサー(device only) : スポーツ＆アウトドア"
[熊野大社]: http://www.kumanotaisha.or.jp/ "出雲國一之宮　熊野大社"
[八雲温泉ゆうあい熊野館]: https://www.kumanokan.jp/ "八雲温泉ゆうあい熊野館"
[木次乳業]: https://www.kisuki-milk.co.jp/ "木次乳業"
[島根県立図書館]: https://www.library.pref.shimane.lg.jp/ "島根県立図書館"
[松江テルサ]: https://www.matsue-terrsa.jp/ "松江テルサ"

## 参考

{{% review-paapi "B08BZ5T9NZ" %}} <!-- GARMIN EDGE 130 PLUS サイクルコンピュータ -->
{{% review-paapi "B0BLNFPWTQ" %}} <!-- trimm ROLLIN サイクルコンピュータ -->
{{% review-paapi "B09TVLHJ1X" %}} <!-- Shokz OpenRun Mini 骨伝導ヘッドセット -->
{{% review-paapi "B08L4WKDZ7" %}} <!-- PowerShot ZOOM -->
{{% review-paapi "B0D5LQBL2P" %}} <!-- 余りモノ異世界人の自由生活 7 -->
<!-- eof -->
