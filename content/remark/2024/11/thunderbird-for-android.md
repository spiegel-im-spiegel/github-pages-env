+++
title = "Android 版 Thunderbird を導入する"
date =  "2024-11-02T11:09:23+09:00"
description = "K-9 ベースの Andorid 版 Thunderbird がようやく正式リリースとなったようだ。"
image = "/images/attention/kitten.jpg"
tags = [ "communication", "tools", "thunderbird", "android" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[K-9] ベースの Andorid 版 [Thunderbird] がようやく正式リリースとなったようだ。

- [Thunderbird for Android 8.0 Takes Flight - The Thunderbird Blog](https://blog.thunderbird.net/2024/10/thunderbird-for-android-8-0-takes-flight/)

Android 版のメールソフトは大昔に [K-9] や初期の Thunderbird も使ったことがあるけど，見た目や使い勝手がイマイチで削除したんだよな。
更にコミュニケーション手段としての電子メールの需要が落ちてきたこともあって Android 機では標準の Gmail 以外は使わなくなった。

あれから状況が変わってきて，特に Google/Gmail への（プライバシー方面の）信頼性が地に落ちたこともあり，現在 Gmail アカウントへの依存を減らして他の ISP のメールサービスに切り替えつつある。
まぁ，思いついたときにぼちぼちという感じではあるが。

そこで Andorid 機でも使える Gmail 以外のメールソフトが欲しいなぁ，と思ってたのよ。

[K-9] ベースの Andorid 版 [Thunderbird] についてはだいぶ前から注目していたが進捗が捗々しくなく最近は諦め気味であったが（ベータ版には興味がない），ようやく正式リリースに到達したようで，まずはめでたし。

さっそくインストールしてみよう。

{{< fig-img-quote src="./thunderbird-for-android.png" title="Thunderbird for Android" link="./thunderbird-for-android.png" >}}

すでに PC 版 Thunderbird (v128.4 以降) を使っているのであれば QR コードを使ってアカウント情報を Andorid 版 [Thunderbird] にインポートできる。

PC 版 Thunderbird から設定をエクスポートするにはメニューの「編集」→「設定」を開いて「モバイル向けのエクスポート」を開く。

{{< fig-img-quote src="./settings.png" title="Thunderbird 設定（1）" link="./settings.png" width="1174" >}}

ここでエクスポートしたいアカウントを選択して `[エクスポート]` ボタンを押すと QR コードが表示される。

{{< fig-img-quote src="./settings-2.png" title="Thunderbird 設定（2）" link="./settings-2.png" width="1127" >}}

これを Andorid 版 [Thunderbird] 側で読み込めばよい（設定 {{% icons "settings" %}} → 設定をインポート でインポートできる）。
アカウントのインポートは仕組み上は複数一度にできるが，アカウントの数が増えると QR コードの読み取りに失敗するみたいなので，ひとつずつインポートしていくのが無難だろう。

これで準備完了。
使い勝手についてはおいおい検証していこう。
とりあえず OpenPGP 暗号化機能が組み込まれていることは（設定から）確認できた。

## ブックマーク

- [Android版「Thunderbird」がようやく正式リリース - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1635729.html)
- [QRコードでアカウント設定をモバイル版に ～「Thunderbird」v128.4.0esrがリリース - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1635393.html)

[K-9]: https://play.google.com/store/apps/details?id=com.fsck.k9 "K-9 Mail - Google Play"
[Thunderbird]: https://play.google.com/store/apps/details?id=net.thunderbird.android "Thunderbird: Free Your Inbox - Google Play"
