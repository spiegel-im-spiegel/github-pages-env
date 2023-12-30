+++
title = "KeePassXC に TOTP を設定する"
date =  "2023-08-17T21:48:05+09:00"
description = "Google Authenticator を捨てて TOTP 管理を KeePassXC に一元化する。"
image = "/images/attention/kitten.jpg"
tags = [ "security", "management", "authentication", "tools" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

私が愛用しているパスワードマネージャ [KeePassXC] の 2.7.6 がリリースされていた。

- [Release Release 2.7.6 · keepassxreboot/keepassxc · GitHub](https://github.com/keepassxreboot/keepassxc/releases/tag/2.7.6)

これで思い出したのだが [Google Authenticator] を捨てて TOTP 管理を [KeePassXC] に一元化しようとしてたのだった。
すっかり忘れてたよ。

[KeePassXC] は [KeePass] 2.x 互換のアプリケーションで Windows, Linux, macOS で動作するマルチプラットフォームのアプリケーションである。
特に [KeePassXC] は標準機能で TOTP のシークレットを登録してワンタイムパスワードを振りだすことができる[^kp2a]。

[^kp2a]: [KeePassXC] の機能については[リポジトリページ](https://github.com/keepassxreboot/keepassxc "keepassxreboot/keepassxc: KeePassXC is a cross-platform community-driven port of the Windows application “Keepass Password Safe”.")を参照のこと。他に [KeePass] 2.x 互換のアプリケーションとして Android 用の [Keepass2Android] というのもあって，こちらも TOTP に対応している。 iPhone 版は知らん。

（以下は [Ubuntu] 環境でのセットアップ）

TOTP の設定はコンテキストメニューから可能。

{{< fig-img src="./context-menu-1.png" title="KeePassXC コンテキストメニュー" link="./context-menu-1.png" width="731" >}}

「TOTP の設定」を選択すると以下のウィンドウがポップアップする。

{{< fig-img src="./setting-totp.png" title="KeePassXC TOTP の設定" link="./setting-totp.png" >}}

ここの「秘密鍵」の項目に各サービスで振り出されるシークレット（あるいは秘密鍵）を入力する。
他は「既定の設定 (RFC 6238)」のままで大抵は問題ない。

たとえば {{% emoji "X" %}} (旧 Twitter) の場合なら，設定で セキュリティとアカウントアクセス → セキュリティ → 2要素認証 と進み

{{< fig-img src="./twitter-secret-1.png" title="Twitter  2要素認証" link="./twitter-secret-1.png" width="607" >}}

QRコード（上の図は塗り潰してる）の下の “Can't scan the QR code?” のリンクをクリックすると

{{< fig-img src="./twitter-secret-2.png" title="Can't scan the QR code?" link="./twitter-secret-1.png" width="607" >}}

という感じにシークレットが表示されるので（上の図の塗り潰してる部分），そのままコピペすればよい。
これでコンテキストメニューから TOTP のワンタイムパスワードが取得可能になる。

{{< fig-img src="./context-menu-2.png" title="KeePassXC コンテキストメニュー" link="./context-menu-2.png" width="731" >}}

たとえば「TOTP を表示」を選択すると，こんなウィンドウが表示される。

{{< fig-img src="./totp.png" title="KeePassXC TOTP を表示" link="./totp.png" >}}

TOTP の設定を紛失すると，最悪の場合，そのサービスに二度と入れくなくなる。
サービスで2要素認証を設定したなら，必ずリカバリーコード（あるいはバックアップコード）をダウンロードし保存しておくこと。
紙に出力しておくのもいいだろう。
ただしリカバリーコードの管理は慎重に。

...よしよし。
これでまたひとつ Google 依存が減ったな。

## ブックマーク

- [Google Online Security Blog: Google Authenticator now supports Google Account synchronization](https://security.googleblog.com/2023/04/google-authenticator-now-supports.html) : この記事を見て [Google Authenticator] を捨てようと思ったのだった

- [Ubuntu に KeePassXC を導入する]({{< ref "/remark/2019/08/installing-keepassxc-in-ubuntu.md" >}})
- [Authenticator と AAL]({{< ref "/remark/2020/09/authenticator-and-aal.md" >}})

[KeePassXC]: https://keepassxc.org/ "KeePassXC Password Manager"
[Google Authenticator]: https://play.google.com/store/apps/details?id=com.google.android.apps.authenticator2 "Google Authenticator - Apps on Google Play"
[KeePass]: https://keepass.info/ "KeePass Password Safe"
[Keepass2Android]: https://play.google.com/store/apps/details?id=keepass2android.keepass2android&hl=en_US "Keepass2Android Password Safe - Apps on Google Play"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
{{% review-paapi "4757143044" %}} <!-- 信頼と裏切りの社会 -->
