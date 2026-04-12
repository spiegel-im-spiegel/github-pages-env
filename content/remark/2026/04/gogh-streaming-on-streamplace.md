+++
title = "Streamplace で Gogh 作業配信を行う"
date =  "2026-04-12T15:23:32+09:00"
description = "Streamplace とは / Streamplace 側の準備 / OBS Studio 側の準備と配信の開始 / スマホ用アプリ"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "atproto", "media", "tools", "bluesky", "web", "streaming" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

[前回]の最後の方で「[Streamplace] を使うことを考えてみようかな」と書いたが，実際に試してみることにした。

## Streamplace とは

[Streamplace] は分散型 SNS 向けのビデオレイヤおよびオープンソースのライブビデオ配信プラットフォームで [AT Protocol (Authenticated Transfer Protocol)][atproto] エコシステム下で構築されているのが特徴らしい。

[Kagi Assistant] に [Streamplace] と YouTube の比較表を作ってもらった（YouTube 以外の動画配信サービスでも似たようなものだろう）。

{{< div-box type="markdown" >}}
| 比較項目 | Streamplace | YouTube |
| :--- | :--- | :--- |
| **運営形態** | 分散型（特定の企業に依存しない） | 中央集権型（Google が運営） |
| **基盤技術** | AT Protocol / Livepeer ネットワーク | Google 独自のクローズドなインフラ |
| **ソースコード** | オープンソース（誰でも利用・改善可能） | 非公開（プロプライエタリ） |
| **データの所有権** | ユーザー自身（自己主権型アイデンティティ） | 運営企業（Google） |
| **主な目的** | 分散型 SNS 向けのビデオインフラ提供 | 広告収益を主とした動画共有サービス |
| **検閲・制限** | プロトコルレベルでの検閲は困難 | 運営企業のポリシーにより削除・停止あり |
| **主な利用者層** | Web3 開発者，分散型 SNS ユーザー | 一般消費者，クリエイター，広告主 |
{{< /div-box >}}

だいたい合ってるかな。

[Streamplace] および基盤技術である [Livepeer] については[別記事][次回]でもう少し掘り下げて紹介する予定である。
まずは（細かいことは考えず）アカウントを取得するところから始めよう。

## Streamplace 側の準備

公式サイトの[トップページ][Streamplace]の右上にある {{% keytop %}}Sign Up{{% /keytop %}} をクリックする。

{{< fig-img-quote src="./streamplace-signup.png" title="Streamplace Signup" link="https://stream.place/" lang="en">}}

サインアップには Bluesky などのアカウントが必要らしい。
アカウントを入力して（上の図では隠れているが）下の方にある {{% keytop %}}Continue{{% /keytop %}} をクリックすると OAuth の認証画面になるので確認して承認すればサインアップは完了。

{{< fig-img-quote src="./pds-oauth.png" title="OAuth Authentication" link="./pds-oauth.png" width="1130" lang="en">}}

ドメイン等をきちんと確認すること。

サインアップ（あるいはサインイン）したら Live Dashboard に移動する。

{{< fig-img-quote src="./live-dashboard.png" title="Live Dashboard" link="https://stream.place/live" width="1550" lang="en">}}

ここで {{% keytop %}}Stream from OBS{{% /keytop %}} をクリックする。

{{< fig-img-quote src="./stream-from-obs-1.png" title="Stream from OBS" link="https://stream.place/live" width="639" lang="en">}}

プロトコルは RTMPS を選択。
この状態で {{% keytop %}}Generate Stream Key{{% /keytop %}} をクリックする。

{{< fig-img-quote src="./stream-from-obs-2.png" title="Stream from OBS" link="https://stream.place/live" width="639" lang="en">}}

Stream Key が伏せ字で表示されるので右側のボタンをクリックしてクリップボードにコピーする。

OBS 側の設定に必要なので以下の情報はどこかにメモっておこう。

- **Service：** Custom...
- **Server：** `rtmps://rtmps.stream.place:1935/live`
- **Stream Key：** 上で取得した値
- **Output Settings**
  - **Output Mode：** Advanced
  - **Keyframe Interval：** 1s
  - **x264 Options：** `bframes=0`

## OBS Studio 側の準備と配信の開始

当然ながら [gogh] はあらかじめ起動しておいてね。

[OBS] Studio のインストールおよびシーンやソースの設定については割愛する（[前回]記事の後半でもちょこっと説明している）。
各自でいい感じに設定して欲しい。

配信設定については，設定（Settings）画面を表示し配信（Stream）タブで，上の情報を参考に，以下のように設定する。

{{< fig-img-quote src="./obsstudio-settings-stream.png" title="OBS Studio ➢ Settings ➢ Stream" link="./obsstudio-settings-stream.png" width="1003" lang="en">}}

さらに出力（Output）タブでも，以下のように設定する。

{{< fig-img-quote src="./obsstudio-settings-output.png" title="OBS Studio ➢ Settings ➢ Output" link="./obsstudio-settings-output.png" width="1003" lang="en">}}

これであとは{{% keytop %}}Start Streaming{{% /keytop %}} ボタンを押せば [Streamplace] 側に配信が流れる。
[Streamplace] の Live Dashboard が以下のような状態になっていれば OK（まだプレビュー状態）。

{{< fig-img-quote src="./live-dashboard-preview.png" title="Live Dashboard" link="https://stream.place/live" width="1296" lang="en">}}

タイトル等を入力して {{% keytop %}}Start Livestream!{{% /keytop %}} をクリックすれば配信が開始される。

{{< fig-img-quote src="./live-dashboard-onlive.png" title="Live Dashboard" link="https://stream.place/live" width="1296" lang="en">}}

ちなみに，最近の Bluesky はライブ中フラグを立てることができる。

{{< fig-img-quote src="./bluesky-edit-live-status.png" title="Edit live status" link="./bluesky-edit-live-status.png" lang="en">}}

ライブ中状態にすると以下のようにアバターアイコンに赤枠が付く。

{{< fig-img-quote src="./bluesky-avater-icon.png" title="Profile" link="https://bsky.app/profile/baldanders.info" width="613" lang="en">}}

## スマホ用アプリ

[Streamplace] の配信は Web ブラウザで見れるが，スマホ用のアプリも提供されている。

- [Streamplace - Apps on Google Play](https://play.google.com/store/apps/details?id=tv.aquareum)
- [Streamplace App - App Store](https://apps.apple.com/us/app/streamplace/id6535653195)

Android だとこんな感じ（横持ちの場合）。

{{< fig-img-quote src="./android-app.jpg" title="Streamplace App" link="https://play.google.com/store/apps/details?id=tv.aquareum" width="2197" lang="en">}}

見るだけならサインインしなくても OK なようだ。
これならスマホ版の [gogh] は要らんな（笑） 配信[^s1] 仕掛けてポモドーロ・タイマーを起動しておけば出先でも使えるぢゃん。

[^s1]: [Streamplace] 配信は特定のユーザにだけ見えるようにすることもできる。

...[後半][次回]へ続く。

[前回]: {{< ref "./tekipaki-working-with-gogh.md" >}} "Gogh で てきぱきワーキング"
[次回]: {{< ref "./streamplace-and-livepeer.md" >}} "Streamplace と Livepeer"
[gogh]: https://gogh.gg/ "gogh - Focus with Your Avatar"
[OBS]: https://obsproject.com/ "Open Broadcaster Software | OBS"
[Steam]: https://store.steampowered.com/ "Welcome to Steam"
[Proton]: https://github.com/ValveSoftware/Proton "ValveSoftware/Proton: Compatibility tool for Steam Play based on Wine and additional components"
[Streamplace]: https://stream.place/ "Streamplace"
[atproto]: https://atproto.com/ "The AT Protocol"
[Livepeer]: https://livepeer.org/ "Livepeer — The open network for real-time AI video"
[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"
