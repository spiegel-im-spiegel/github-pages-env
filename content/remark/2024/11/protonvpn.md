+++
title = "ブラックフライデーなので Proton VPN に加入する"
date =  "2024-11-26T21:39:13+09:00"
description = "今どきの VPN って WireGuard で繋いでるんだねぇ"
image = "/images/attention/kitten.jpg"
tags = [ "security", "privacy", "risk", "management", "tools", "vpn" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

まぁ確かに「闇バイト」なるものには困ったものだが，それでまた通信傍受が云々とか言い出す某政治家には本当に困ったものである。
ぶっちゃけ「サイバー犯罪」に関して警察にできることは犯罪者にもできるし，法で規制しようとしても（法を破るから犯罪者なのであって）割りを食うのは善良で従順な庶民なんだよな。

{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}。

個人的な話で申し訳ないが，スマホで今使ってる VPN アプリである Freedome は来年で契約が切れて以降は使えなくなるので別のアプリを検討しないといけないのだが，後継の [F‑Secure VPN](https://www.f-secure.com/vpn) はどう見ても値段に対して機能が釣り合ってない。
まぁ F-Secure は WithSecure に[社名変更](https://cloud.watch.impress.co.jp/docs/news/1397156.html "F-Secureの企業向け部門、新ブランド「WithSecure」として分社化へ - クラウド Watch")して B2B にシフトしてしまったので（F-Secure のブランド自体は個人向け用に残すみたいだが），しょうがないところである。

一応，代替手段として以前から [Proton VPN] を考えていたのだが，折角の機会だし乗り換えてしまおう。

実は [Proton Mail] のほうは大昔に無料アカウントをとっていたのだが，コミュニケーション手段として（暗号化）電子メール自体をあまり使わなくなったため放置していた。
さっき確認したらアカウントがまだ生きてるみたいなので，これを使って [Proton VPN] を設定してみよう。
どうやら [Proton Mail] の無料アカウントで（追加料金なしで）ひとつのデバイスに [Proton VPN] を設定できるらしい。
さっそく手元のスマホにインストールしてみる。

- [Proton VPN: 速くて安全なVPN - Apps on Google Play](https://play.google.com/store/apps/details?id=ch.protonvpn.android)

予想されたことだが，無料版では本当に必要最小限の機能しか提供されないようだ。

- 日本リージョンのAPには繋がる
- マルウェアやトラッカーやアプリ上の広告をブロックする NetShield が使えない
- スプリット（アプリや IP アドレスによって経路を分ける）が使えない
- LAN 上の機器に（VPN を避けて）アクセスできない
- プロトコルの指定ができない

まぁそりゃそうだろう，ということで無料版は諦めてサブスクリプションの料金表を眺めてみる。

{{< fig-img-quote src="./price-of-protonvpn.png" title="Pricing | Proton VPN" link="https://protonvpn.com/pricing" width="935" lang="en" >}}

Google の各サービスから置き換えたいなら Proton Unlimited もアリかも知れないが，今のところは VPN Plus でいいだろう。
支払い処理を進めていくとブラックフライデーの割引があることに気がつく。

{{< fig-img-quote src="./black-friday.png" title="Proton BLACK FRIDAY" link="./black-friday.png" lang="en" >}}

よしよし。
これで買ってしまえ！ 来年のことは来年の私が考える。

というわけで無事に購入し，スマホアプリもフル機能が使えるようになった。

{{< fig-img-quote src="./protonvon-android.png" title="Proton VPN for Android" link="./protonvon-android.png" lang="en" >}}

ステータスはこんな感じ。

{{< fig-img-quote src="./protonvpn-status.png" title="Proton VPN for Android" link="./protonvpn-status.png" lang="en" >}}

おー。
[WireGuard] で繋いでるのか。
今どきの VPN ってこんな感じなんだねぇ。
細かい機能とかはおいおい調べていこう。

とりあえず，来年 Freedome の契約が切れるまでに手持ちの全マシンに仕込んでいかないと。

[Proton VPN]: https://protonvpn.com/ "Proton VPN: Secure, fast VPN service in 110+ countries"
[Proton Mail]: https://proton.me/mail "Proton Mail: Get a private, secure, and encrypted email account | Proton"
[WireGuard]: https://www.wireguard.com/ "WireGuard: fast, modern, secure VPN tunnel"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
{{% review-paapi "B01MZGVHOA" %}} <!-- 超監視社会 -->
