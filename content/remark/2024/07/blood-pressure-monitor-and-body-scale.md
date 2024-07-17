+++
title = "血圧計と体重計の計測値を自動入力したい"
date =  "2024-07-17T23:56:00+09:00"
description = "スマート体重計の Wi-Fi 接続で苦労する / スマート血圧計のセットアップ / Google Fit との連携に注意 / Garmin Connect にデータが送れない / ECLEAR plus へ歩数データが上手く入らない"
image = "/images/attention/kitten.jpg"
tags = [ "disease", "healthcare", "tools", "gear" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = true
+++

## 血圧計と体重計の計測値を自動入力したい

きっかけは Amazon の Prime Day で [Anker 製のスマート体重（体組成）計](https://www.amazon.co.jp/dp/B0BG5FJ33T?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Amazon.co.jp: Anker Eufy (ユーフィ) Smart Scale P2 Pro（体重体組成計）【アプリ対応/Fitbit連携/体脂肪率/BMI/心拍数/筋肉量/基礎代謝量/水分量/体脂肪量/骨量/内臓脂肪/タンパク質/骨格筋量/皮下脂肪/体内年齢/ボディタイプ / 3Dモデル】ブラック : ホーム＆キッチン")が安売りされてるのを見かけたこと。

専用のスマホアプリと連携して体重管理するのだが Google Fit (Android), Apple ヘルスケア (iOS), および Fitbit (Android/iOS) とも連携できるらしい。
いいなぁ。
でも，私は Garmin ユーザなので [Garmin Connect] と連携して欲しいわけですよ。
もっと個人的なことを言うと，ダイエットが一段落してリバウンドもないため体重のモニタリングの優先順位はさほど高くなくて，どちらかというとスマホと（できれば [Garmin Connect] と）連携できる血圧計のほうが嬉しいんだけどなぁ。

スマホアプリと連携できる血圧計というと[オムロンの血圧計](https://www.amazon.co.jp/dp/B07ZK8VZMG?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Amazon | オムロン上腕式血圧計 HCR-7502T | オムロン(OMRON) | 上腕式血圧計")が有名だけど，そこそこいいお値段なのと Google Fit や Fitbit といった他サービスとの連携ができないのがネックで手を出しかねていた。

で，「他にいい製品はないの？」と何となく Amazon サイト内を彷徨ってたらエレコムが，この手の製品を出してるのね。
知らなかったよ。

{{% review-paapi "B084ZDCVLJ" %}} <!-- スマート血圧計 エレコム ELECOM -->
{{% review-paapi "B0CFTT1L2M" %}} <!-- スマート体重体組成計 エレコム ELECOM -->

これらの製品で測定したデータを専用アプリ [ECLEAR plus] で統合的に管理できる。
測定データは CSV または PDF レポートとして出力可能（クラウド・ストレージに直接保存できる）。
[ECLEAR plus] 経由で Google Fit または Apple ヘルスケアとも連携可能。
しかも安い！

2022年末に[心筋梗塞]({{< ref "/remark/2022/12/heart-attack.md" >}} "ハライタだと思った？ 残念！ 心筋梗塞でした")をやらかして以来，毎朝「高血圧管理手帳」に血圧と体重を記録し続けてるんだけど

{{< fig-img src="/remark/2022/12/heart-attack/52583275820_69191cfcae_e.jpg" title="私も高血圧管理手帳持ちになりました | Flickr" link="https://www.flickr.com/photos/spiegel/52583275820/">}}

手書きなんだよね，これ。
今年に入って（いつか再利用できるかと思って）スプレッドシート（LibreOffice Calc）でも記録してるんだけど，手作業には変わりない。
いい加減に飽きてきたのよ。
インプットもアウトプットもできるだけ自動化したい！

連携サービスが Google Fit と Apple ヘルスケアだけというのは惜しいが，上述の製品なら [ECLEAR plus] で入力もレポート化も自動化できるので，アリなのか？

というわけで，うーやーたー！ と気合を入れて発注してみた。
今年も Prime Day に踊らされてしまったな（しかも安売り対象製品は買ってない）。

## スマート体重計の Wi-Fi 接続で苦労する

先にスマート体重計が来たので，まずは体重計から。

{{< fig-img src="./53860851485_8f5537d695_e.jpg" title="スマート体重計キタ！ | Flickr" link="https://www.flickr.com/photos/spiegel/53860851485/" >}}

その前にスマホに [ECLEAR plus] をインストールしてサインアップする。
セットアップは[オンラインマニュアル](https://app.elecom.co.jp/eclear_plus/manual.html "ECLEAR plus オンラインマニュアル")を見ながらどうぞ。

何故か Wi-Fi 接続。
挙動を見るにデータを直にクラウドにアップロードしているように見える。
つまり体重計とスマホは直接連携しているわけではなく，あくまでクラウドを介してということのようだ。
データの評価やレポート作成もクラウド側でやっているのだろう。
スマホアプリは結果を表示してるだけって感じ。
大丈夫か？ これ。

Wi-Fi は5G帯に対応してない。
Wi-Fi 接続設定をスマホ側で行うんだけど，スマホも一時的に2.4G帯に繋ぎ直さないといけない。
更に中継機でメッシュを構成していると上手く設定できないようだ。
つか，できなかった。
不良品か思うて返品しそうになったわ（笑） 中継機の電源を切って対処。
接続設定が完了したらネット環境は全部元に戻して構わない。

## スマート血圧計のセットアップ

スマート血圧計も到着。

{{< fig-img src="./53862684506_c8e52a8846_e.jpg" title="スマート血圧計キタ（充電ケーブルが極悪w） | Flickr" link="https://www.flickr.com/photos/spiegel/53862684506/" >}}

バッテリ駆動なのはいいとして，なんだこの充電ケーブル。
エレコムってこんなことする会社だっけ？

スマホとの連携は問題なく出来た。

{{< fig-img src="./eclear-plus.png" title="ECLEAR plus" link="./eclear-plus.png" >}}

上腕式なのはありがたいが[^h1] カフが巻きにくい。
買ったばっかりで馴染まないのかな。
でも，実際に計測してみたらそれっぽい値が出たので問題はなかろう。

[^h1]: 病院では自宅で血圧を測るなら上腕式の血圧計を使うようにと指導を受けている。

前節のスマート体重計は Wi-Fi 接続で勝手にクラウドにデータがアップロードされていたが，
こちらは Bluetooth 接続なので，データ転送時にスマホ側を待ち受け状態にしないといけない。
まぁ，大した手間ではないか。

これで体重と血圧のデータをスマホアプリで管理できるようになった。
体重と血圧の測定は継続的に行うことで身体異常を早期発見できるようになる。
というわけで，これからもデータの蓄積を行っていこう。

## Google Fit との連携に注意

Android 版 [ECLEAR plus] の場合，上のスナップショットのように「Google Fit と連携」を有効にすれば Google Fit に体重・血圧データが転送される。
また Google Fit にある歩数データが [ECLEAR plus] に転送される。
「Google Fit と連携」を有効にする際は全ての項目へのアクセスを許可すること。

[ECLEAR plus] では一度連携を行うと取り消しできない。
連携を許可した状態で「Google Fit と連携」を無効にしてもデータ転送が行われないだけで連携状態は維持しているようだ。

また Google Fit 側で [ECLEAR plus] との連携を解除しても [ECLEAR plus] 側はそれを認識しないみたいで，連携のやり直しが出来ない。
私はアプリを弄ってるうちにこの状態になってしまい，結局 [ECLEAR plus] をいったんアンインストールした後にインストールし直して，全ての設定をやり直す羽目になった。
この辺の作りがなんとも微妙なんだよな。

## Garmin Connect にデータが送れない

最初の方で「できれば [Garmin Connect] と」とこっそり書いたが，私の場合ウォーキングなどのアクティビティの管理を [Garmin Connect] で行ってるので，体重や血圧といった情報は Garmin 側に集約したい。
色々とググって[^s1] みるとサービス間でデータの受け渡しができるハブ的なサービスもあるらしい。
以下に2つほど挙げてみる。

[^s1]: 実際に検索に使うのは [Kagi Search]({{< ref "/remark/2024/06/kagi-search.md" >}} "Kagi Search を試してみる 〜 検索サービスも有料の時代？") だが（笑）

### [MyFitnessPal]

食事管理用のサービスおよびアプリらしい。
他サービスからアクティビティ情報や体重などの身体データを集約または転送できるみたい。

- [【Garmin Connect】”体重”を自動連携させる一番”コスパ”良い方法（お薦めです） | よしッくすＣｈ](https://55544aki.com/body-composition-meter/)

この記事を見て試してみた。
Google Fit と連携できるものの，体重等のデータが [MyFitnessPal] へ転送されなかった。
逆向きで Google Fit へカロリーデータ等は送れるみたい。

[MyFitnessPal] では今回の要件に合わないため諦めた。
残念。

### [Health Sync]

アクティビティや体重などの身体データをサービス間でやりとりできるアプリ。
継続的に利用するためにはお金を払う必要があるが，480円一括払いでOKだった。

こちらも残念なことに [Garmin Connect] へ身体データを転送することが出来なかった。
ただし [Garmin Connect] から Google Fit (または Health Connect) に転送することは可能。

というわけで，泣く泣く以下の設定にした。

{{< fig-img src="./health-sync.png" title="Health Sync" link="./health-sync.png" >}}

これで全てのデータは（[Garmin Connect] ではなく） Google Fit に集約されることになった `orz`

どなたか体重と血圧データを Google Fit から [Garmin Connect] へ転送する方法をご存じないですかねぇ...

## [ECLEAR plus] へ歩数データが上手く入らない

さて [Health Sync] を採用したことで歩数データの流れは以下のようになった。

{{< fig-mermaid >}}
graph TB
  garmin["Garmin vívosmart 5"]
  gconnect["Garmin connect"]
  healthsync["Health Sync"]
  fit["Google Fit"]
  eclear["ECLEAR plus"]

  garmin -- write --> gconnect
  gconnect -- write --> healthsync
  healthsync -- write (Health Connect) --> fit
  fit -- write --> eclear
{{< /fig-mermaid >}}

これで Google Fit までは問題なく歩数データが転送されるんだけど，何故か [ECLEAR plus] に上手くデータが渡らない。
具体的には歩数がかなり少なめに記録されてしまうのだ。

どうもウォーキングなどのアクティビティで発生する歩数を [ECLEAR plus] では除外しているっぽい。
というか，アクティビティデータは Google Fit から [ECLEAR plus] へ転送されないため，必然的にそうなってしまうのだろう？

なんだかなぁ。
ままならないものである。

## ブックマーク

- [GARMIN(ガーミン)とGooglefit(グーグルフィット)の接続方法 | うつ抜けライダー闘病記](https://ameblo.jp/kdm-12151018/entry-12793409703.html)

[Garmin Connect]: https://connect.garmin.com/
[ECLEAR plus]: https://app.elecom.co.jp/eclear_plus/ "ECLEAR plus | 血圧や体重、歩数の管理ができるヘルスケアアプリ"
[MyFitnessPal]: https://www.myfitnesspal.com/ "無料カロリーカウンター、ダイエット & エクササイズ日記 | MyFitnessPal"
[Health Sync]: https://healthsync.app/ "Health Sync"

## 参考

{{% review-paapi "B08BZ5T9NZ" %}} <!-- GARMIN EDGE 130 PLUS サイクルコンピュータ -->
{{% review-paapi "B09XGYX7JF" %}} <!-- GARMIN vívosmart 5 -->
{{% review-paapi "B08VRTDBW9" %}} <!-- 無線 KAN 中継機 -->
