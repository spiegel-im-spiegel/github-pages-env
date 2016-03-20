+++
date = "2016-03-19T18:26:28+09:00"
description = "description"
draft = true
tags = ["security", "cryptography", "e-mail"]
title = "週末スペシャル： ProtonMail のアカウントを作りました"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

## ProtonMail のアカウントを作りました

- [暗号化電子メールのProtonMailがベータ段階を卒業--「iOS」「Android」アプリを公開 - CNET Japan](http://japan.cnet.com/news/service/35079779/)
- [Announcement: ProtonMail has launched worldwide! - ProtonMail Blog](https://protonmail.com/blog/protonmail-launch-worldwide/)
- [ProtonMail](https://github.com/ProtonMail) (GitHub)
    - [OpenPGP.js](https://openpgpjs.org/ "OpenPGP.js | OpenPGP JavaScript Implementation") : フロントエンドで使われている暗号化ライブラリ

というわけで [ProtonMail] のアカウントを作った。
よさげなら本サイトの連絡先アドレスとして使おうかと思っている。

個人で利用するなら無料で使える。
ただしストレージサイズや1日あたりの流量に制限がある。
カスタム・ドメインを設定したり複数のアドレスが必要な場合は 5 ユーロ/月から利用できる。

MUA (Messaging User Agent) は [ProtonMail] が提供するもの以外は使えないようだ。
今のところ Android および iOS 用のアプリが提供されているほか， Web サービスもある。

- [ProtonMail - Encrypted Email - Google Play](https://play.google.com/store/apps/details?id=ch.protonmail.android)
- [ProtonMail Encrypted Email - Easy To Use, Free For All on the App Store](https://itunes.apple.com/app/protonmail-encrypted-email/id979659905)

[ProtonMail] ユーザ間のメールは End-to-End で暗号化される[^pm1]。
[ProtonMail] 以外からのメールは当然暗号化されていないが，ストレージへの保存は暗号化されるようだ。

[^pm1]: アカウント作成時にアカウントのパスワードと mailbox のパスワードを設定する必要がある。 mailbox のパスワードは復号に使われていて（つまり公開鍵暗号方式を使っている） [ProtonMail] 側では保持していない（保持していたら [ProtonMail] 側も復号できることになる）。またメールには有効期限を付けることができる。

- [ProtonMail | Security Features Overview](https://protonmail.com/security-details)

[ProtonMail] 以外のユーザにメールを送る場合は，あらかじめパスワードを決めておけば，暗号化メールを作成することができる。
メール受信者は以下のような内容のメールを受け取る。

```text
I am using ProtonMail to send and receive secure emails. Click the link below to decrypt and view my message:
View Secure Message
Message expires
2016-04-16 08:48:06 GMT
(672 hours after this message was sent.)
```

“View Secure Message” のリンクから [ProtonMail] のサイトに移動しパスワードを入力すれば該当のメッセージを見ることができる（アカウントは不要）。
更に，そのまま返信すれば暗号化メールで返信できる。
これって Phishing のリスクがあるし，どうやって安全にパスワードを交換するかという問題もあるのだが，まぁしょうがないのかな。

ちなみに [ProtonMail] へ PGP/MIME 署名したメールを送ってみたが解釈してくれなかった。
ただし

{{< fig-quote title="ProtonMail Open Source Cryptography - ProtonMail Blog" link="https://protonmail.com/blog/protonmail-open-source-crytography/" lang="en" >}}
<q>In the future, we will be adding to ProtonMail the ability to import and export PGP keys. By complying with OpenPGP, it will be possible to do things like, download ProtonMail messages and decrypt them locally using your own PGP software.</q>
{{< /fig-quote >}}

ということらしいので，楽しみに待つことにする。
[OpenPGP はセキュリティやプライバシーに関心のない人からは「面倒」と思われている]({{< relref "remark/2015/use-the-signal-luke.md" >}})らしいが，この辺のギャップを [ProtonMail] が埋めてくれることを期待したい。

- [なぜジョニーは今もやっぱり暗号化できないのか：現在のPGPクライアントの使いやすさ評価 - YAMDAS現更新履歴](http://d.hatena.ne.jp/yomoyomo/20151112/jonnycantencrypt)

## その他の気になる記事{#other}


[ProtonMail]: https://protonmail.com/ "Secure email: ProtonMail is free encrypted email."