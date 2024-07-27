+++
title = "お散歩カメラ 2024-07-20"
date =  "2024-07-21T21:55:44+09:00"
description = "祝♪ 復活のナビ先生 / Strava オススメのルートを試してみる / 定点観測 / そうだ。温泉に行こう"
image = "/remark/2024/07/20-osanpo-camera/53867532427_d353040ac3_o.jpg"
tags = [ "osanpo-camera", "photography", "matsue", "bike", "tools", "gear" ]
license = "by"
pageType = "photo"

[scripts]
  mathjax = false
  mermaidjs = false
+++

## 祝♪ 復活のナビ先生

[前回]「もしアプリが復帰したらサイコン同時使用とかできないだろうか」と書いたが [trimm ROLLIN] の[専用アプリ][trimm Cycling Center]の故障がようやく直った模様。
めでたい！ とはいえ既にメインのサイコンは [GARMIN Edge 130 plus] で運用を始めているので [trimm ROLLIN] はナビゲーション専用機として使うことにしよう。

試してみたが，心拍数をモニタする[スマートバンド]({{< ref "/remark/2023/08/monitoring-heart-rate-2.md" >}} "Garmin vívosmart 5 を買っちまった（心臓リハビリ＠がんばらない 2）")は同時に2つのサイコンと接続できるが，~~手持ちのケイデンスセンサでは2つ同時に接続できないようだ~~。
まぁ，ナビ動作中はケイデンスはさほど気にしないし。
[大丈夫だ問題ない](https://www.j-cast.com/2021/06/26414470.html "「大丈夫だ問題ない」大流行に「罪悪感しかない」　伝説のゲーム「エルシャダイ」生みの親の「懺悔」: J-CAST ニュース")。

{{< div-box type="markdown" >}}
### 【2024-07-27 追記】

後日に分かったが [trimm ROLLIN] 側でケイデンスセンサを認識できなかったのは単にペアリングが外れてたから，だった。
[trimm ROLLIN] 側でペアリングし直したら [GARMIN Edge 130 plus] と同時にモニタできることを確認できた。
まぁ，どのみちナビ中はケイデンスとか見ないのでいいんだけど（笑）

[trimm ROLLIN]: https://www.amazon.co.jp/dp/B0BLNFPWTQ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Amazon.co.jp: trimm ROLLIN サイクルコンピュータ GPS 自転車 速度計 ワイヤレス ナビゲーション ANT+センサー対応 Bluetooth 心拍数 高度計 2.7インチ スピードセンサー(device only) : スポーツ＆アウトドア"
[GARMIN Edge 130 plus]: https://www.amazon.co.jp/dp/B08BZ5T9NZ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Amazon | GARMIN ガーミン エッジ(Edge) 130plus 日本版 本体のみ GPS ブルートゥース Android/iOS対応 (010-02385-05)【日本正規品】 | ガーミン(GARMIN) | トレッキング用GPS・アクセサリ"
{{< /div-box >}}

## Strava オススメのルートを試してみる

最近 [Strava] が30日間の「サブスクリプションプレビュー」ってのをやっていて，プレミアム機能を試せる。
特にパーソナルヒートマップとルート探索機能が面白い。
他の [Strava] ユーザが公開しているアクティビティ情報を元にルート候補を推薦してくれるらしい。

その候補の中から面白いルートがあったので試してみることにした。
具体的には宍道湖大橋北詰から湖沿いの国道431号を西進し，免許センターから一畑電車沿いの道を通って宍道湖北部広域農道へ出る。
そこから更に松江西部農免農道に入り，松江城を回って宍道湖大橋北詰に戻るルート。
具体的には以下の通り。
約17kmほどの道のりかな。

{{< fig-img src="./cycling-route.png" title="国道431号-宍道湖北部広域農道-松江西部農免農道" link="./cycling-route.png" width="1810" >}}

これを GPX 形式のファイルで出力して [trimm Cycling Center] にインポートして実際にナビしてもらうことにした。

まずはスタート地点へ。

{{< fig-img src="./53867328402_6298ac2b73_e.jpg" title="宍道湖大橋北詰バス停付近からスタート | Flickr" link="https://www.flickr.com/photos/spiegel/53867328402/" >}}

なんだこの絵。
なぜ頭にスイカ？？？ まぁいいや。

[GARMIN Edge 130 plus] ではそのままモニタリングを続け，新たに [trimm ROLLIN] を起動してマウントにセットしナビを開始する。

{{< fig-img src="./53868647680_16082b97f4_e.jpg" title="復活のナビ先生 | Flickr" link="https://www.flickr.com/photos/spiegel/53868647680/" >}}

[GARMIN Edge 130 plus] のほうはステムバッグに放り込んでおけばいいだろう。

宍道湖沿いを西に向かって走り出す。

{{< fig-img src="./53868587804_79b89dd30a_e.jpg" title="宍道湖沿いの道は自転車に優しくない | Flickr" link="https://www.flickr.com/photos/spiegel/53868587804/" >}}

前にフォーゲルパークの近くに[行った]({{< ref "/remark/2023/09/matsue-vogel-park.md" >}} "そうだ，松江フォーゲルパーク（の近く）へ行こう")ときも思ったが，宍道湖沿いの道はホンマに自転車に優しくないよなぁ。
まぁ，ロードバイクでバビューンと飛ばしていくならどうってことないんだろうけど，私のはただのクロスバイクだし，そもそも脚力ないからなぁ[^c1]。

[^c1]: [Strava] のセグメント公開記録を見ると，トップ目の人は国道431号線を時速40km以上で走ってる。これなら原チャリとさして変わらん。私はその半分の速度が精一杯だけど。

[免許センター][島根県 運転免許センター]から[一畑電車]沿いの道へ。
ちょっと寄り道して朝日が丘駅の写真を撮っておくか。

{{< fig-img src="./53867337417_017d8786bf_e.jpg" title="一畑電車 朝日が丘駅 | Flickr" link="https://www.flickr.com/photos/spiegel/53867337417/" >}}

前に[免許の更新]({{< ref "/remark/2022/02/moving-2022.md#lic" >}} "引っ越し 2022")をしたときにこの駅を利用したんだよな。

ここから宍道湖北部広域農道に入り，さらに松江西部農免農道に入る。

農免農道の正式名称は「農林漁業用揮発油税財源身替農道」と言うそうな。
なんか農林漁業用の機械にかかる油（揮発油）に対する税金を免除する代わりに，その揮発油税で作ったのが農免農道ってことらしい。

松江西部農免農道は2008年に開通までこぎつけたが，農免農道事業自体は当時の民主党の「事業仕分け」の対象になったんだそうで，2010年以降は新規着手されてないんだとか。
なので，割とレアな道路ということになるだろうか。

農道と言うだけあって，道路の周りは一面田んぼでなかなか壮観だ。
以下は松江・出雲の{{% pdf-file title="郷土史" link="https://www.city.matsue.lg.jp/material/files/group/32/sh.pdf" %}} には必ず出てくる佐陀川の眺めである。

{{< fig-img src="./53868595734_c977dc2240_e.jpg" title="佐陀川 | Flickr" link="https://www.flickr.com/photos/spiegel/53868595734/" >}}

宍道湖北部広域農道も松江西部農免農道も自転車に優しい道だった。
自転車も通行可能な広い歩道。
勾配も殆どなくて快適にペダルを漕げる。
休日だったからかもしれないけど，自動車の流量も少な目。
いい事尽くしである。

今度免許の更新に行くときは宍道湖沿いの国道431号じゃなくて，遠回りでもこっちの農道を通っていこう。
それまで農道とその周辺の道を探検しておきたい。

### 喋らないナビ先生

今回は [Strava] のデータ（GPX 形式）を [trimm Cycling Center] にインポートして利用したんだけど，この方法だと上手くナビゲーションしてくれないようだ。
要するに，単純に道を塗りつぶしたデータを地図に貼り付けただけみたいな状態で，ナビゲーションに必要な道路分岐の情報が脱落してしまってるらしい。
なので，ルートから外れてもだんまりで，ルートの再検索もしてくれない。
これに気が付かなくて何度かコースアウトしてしまった（常にナビ画面を見てるわけじゃないからね）。

[trimm ROLLIN] でナビゲーションしてもらいたいならルートは [trimm Cycling Center] で自前で作らないといけないみたい。
まぁ，そのくらいは大したことないし，次回以降はそれで行こう。

## 定点観測

今日は午後から市内で用事があったので，上述のサイクリングのあと，用事を済ませながらいつもの定点観測。

{{< fig-img src="./53868284781_8a1d1fa1fc_e.jpg" title="今日の大山 from 松江城 | Flickr" link="https://www.flickr.com/photos/spiegel/53868284781/" >}}

{{< fig-img src="./53868636249_7c6f920dfd_e.jpg" title="今日の松江城 | Flickr" link="https://www.flickr.com/photos/spiegel/53868636249/" >}}

ついでに松江城の喫茶店で昼飯。

{{< fig-img src="./53867381207_41cf43933b_e.jpg" title="松江城 興雲閣 亀田山喫茶室 | Flickr" link="https://www.flickr.com/photos/spiegel/53867381207/" >}}

## そうだ。温泉に行こう

思ったより早く用事が済んだので，暑い時間帯であるが[八雲温泉][八雲温泉ゆうあい熊野館]に行くことにした。

{{< fig-img src="./53867532427_a7139ff897_e.jpg" title="八雲温泉へGo | Flickr" link="https://www.flickr.com/photos/spiegel/53867532427/" >}}

{{< fig-img src="./53868850895_d6e3d99c3d_e.jpg" title="八雲温泉上がりのコーヒー牛乳＆日焼け止め | Flickr" link="https://www.flickr.com/photos/spiegel/53868850895/" >}}

夕方近くだと結構人が多いのな。
やっぱ地元の銭湯なんだな，ここ（笑）

さて帰宅するか。
その前にコメダでクールダウン。

{{< fig-img src="./53868749053_b12c5780b2_e.jpg" title="クールダウン＆糖分補給中 | Flickr" link="https://www.flickr.com/photos/spiegel/53868749053/" >}}

夏場のかき氷は偉大である！

[前回]: {{< ref "/remark/2024/07/13-osanpo-camera.md#navi" >}} "スマホアプリのナビゲーションは地図が見れてナンボ？"
[trimm ROLLIN]: https://www.amazon.co.jp/dp/B0BLNFPWTQ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Amazon.co.jp: trimm ROLLIN サイクルコンピュータ GPS 自転車 速度計 ワイヤレス ナビゲーション ANT+センサー対応 Bluetooth 心拍数 高度計 2.7インチ スピードセンサー(device only) : スポーツ＆アウトドア"
[trimm Cycling Center]: https://play.google.com/store/apps/details?id=bike.trimm.rideWithMe&hl=en_US "trimm Cycling Center - Apps on Google Play"
[GARMIN Edge 130 plus]: https://www.amazon.co.jp/dp/B08BZ5T9NZ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Amazon | GARMIN ガーミン エッジ(Edge) 130plus 日本版 本体のみ GPS ブルートゥース Android/iOS対応 (010-02385-05)【日本正規品】 | ガーミン(GARMIN) | トレッキング用GPS・アクセサリ"
[Garmin Connect]: https://connect.garmin.com/ "Garmin Connect"
[Strava]: https://www.strava.com/ "Strava"
[島根県 運転免許センター]: https://maps.app.goo.gl/ULutdQbLNaKKpYLF6
[一畑電車]: https://www.ichibata.co.jp/railway/ "ばたでん【一畑電車株式会社】"
[八雲温泉ゆうあい熊野館]: https://www.kumanokan.jp/ "八雲温泉ゆうあい熊野館"

## 参考

{{% review-paapi "B08BZ5T9NZ" %}} <!-- GARMIN EDGE 130 PLUS サイクルコンピュータ -->
{{% review-paapi "B0BLNFPWTQ" %}} <!-- trimm ROLLIN サイクルコンピュータ -->
{{% review-paapi "B09XGYX7JF" %}} <!-- GARMIN vívosmart 5 -->
{{% review-paapi "B09TVLHJ1X" %}} <!-- Shokz OpenRun Mini 骨伝導ヘッドセット -->
{{% review-paapi "B08L4WKDZ7" %}} <!-- PowerShot ZOOM -->
{{% review-paapi "B08K34WLXD" %}} <!-- ステムバッグ（stem bag） -->
{{% review-paapi "B0BT1X9H7B" %}} <!-- 日焼け止め ミストタイプ -->
{{% review-paapi "B09NRGCBY5" %}} <!-- 日焼け止め 乳液タイプ -->
{{% review-paapi "B09Z6F7MGT" %}} <!-- OS-1 経口補水液 -->
{{% review-paapi "B0CZQ9YF7Y" %}} <!-- Tシャツ シャーロック・ホームズ 221B BAKER STREET -->
{{% review-paapi "B08M23DNWH" %}} <!-- 復活のイデオン -->
