+++
date = "2016-03-20T11:45:24+09:00"
update = "2016-04-09T21:39:04+09:00"
description = "というわけで ProtonMail のアカウントを作った。よさげなら本サイトの連絡先アドレスとして使おうかと思っている。"
draft = false
tags = ["security", "privacy", "cryptography", "tools", "messaging", "openpgp"]
title = "ProtonMail のアカウントを作りました"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/profile/"
+++

- [暗号化電子メールのProtonMailがベータ段階を卒業--「iOS」「Android」アプリを公開 - CNET Japan](http://japan.cnet.com/news/service/35079779/)
- [Announcement: ProtonMail has launched worldwide! - ProtonMail Blog](https://protonmail.com/blog/protonmail-launch-worldwide/)
- [ProtonMail](https://github.com/ProtonMail) (GitHub)
    - [OpenPGP.js](https://openpgpjs.org/ "OpenPGP.js | OpenPGP JavaScript Implementation") : フロントエンドで使われている暗号化ライブラリ

というわけで [ProtonMail] のアカウントを作った。
よさげなら本サイトの連絡先アドレスとして使うことも検討してみる。

個人で利用するなら無料で使える。
ただしストレージサイズや1日あたりの流量に制限がある。
カスタム・ドメインを設定したり複数のアドレスが必要な場合は 48 ユーロ/年から利用できる[^ch]。

[^ch]: [ProtonMail] の開発主体はスイスにある。なお米ドルでの決済も可能。

MUA (Messaging User Agent) は [ProtonMail] が提供するもの以外は使えないようだ。
今のところ Android および iOS 用のアプリが提供されているほか， Web サービスもある。

- [ProtonMail - Encrypted Email - Google Play](https://play.google.com/store/apps/details?id=ch.protonmail.android)
- [ProtonMail Encrypted Email - Easy To Use, Free For All on the App Store](https://itunes.apple.com/app/protonmail-encrypted-email/id979659905)
- [エンドツーエンドでの暗号化で第三者の解読を防ぐWebメールサービス「ProtonMail」 - 知っ得！旬のネットサービス - 窓の杜](http://www.forest.impress.co.jp/docs/serial/netserv/20160331_750622.html)

[ProtonMail] ユーザ間のメールは自動的に End-to-End で暗号化される[^pm1]。
[ProtonMail] 以外からのメールについてもストレージへの保存は暗号化されるようだ。

[^pm1]: アカウント作成時にアカウント認証用のパスワードと mailbox のパスワードを設定する必要がある。 mailbox のパスワードは復号に使われていて（つまり公開鍵暗号方式を使っている） [ProtonMail] 側では保持していない（保持していたら [ProtonMail] 側も復号できることになる）。またメールには有効期限を付けることができる。

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
[OpenPGP はセキュリティやプライバシーに関心のない人からは「面倒」と思われている]({{< ref "/remark/2015/use-the-signal-luke.md" >}})らしいが，この辺のギャップを [ProtonMail] が埋めてくれることを期待したい[^m]。

[^m]: 電子メールを含む全てのメッセージング・サービスは，プライバシー要件の有無にかかわらず，全て End-to-End で暗号化すべきだし，そうでないサービス（Gmail 等も含めて）は将来的には排除すべきと思う。ただ，現状としてはメッセージング・サービスには本当に個人的なやり取りから広告や Phishing まであらゆるメッセージが乗ってしまう。[暗号化されたメッセージから悪意を排除するのは難しい]({{< ref "/remark/2015/cipher-risk.md" >}})。

- [なぜジョニーは今もやっぱり暗号化できないのか：現在のPGPクライアントの使いやすさ評価 - YAMDAS現更新履歴](http://d.hatena.ne.jp/yomoyomo/20151112/jonnycantencrypt)

[ProtonMail]: https://protonmail.com/ "Secure email: ProtonMail is free encrypted email."

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E5%8C%96-%E3%83%97%E3%83%A9%E3%82%A4%E3%83%90%E3%82%B7%E3%83%BC%E3%82%92%E6%95%91%E3%81%A3%E3%81%9F%E5%8F%8D%E4%B9%B1%E8%80%85%E3%81%9F%E3%81%A1-%E3%82%B9%E3%83%86%E3%82%A3%E3%83%BC%E3%83%96%E3%83%B3%E3%83%BB%E3%83%AC%E3%83%93%E3%83%BC/dp/4314009071?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4314009071"><img src="https://images-fe.ssl-images-amazon.com/images/I/51ZRZ62WKCL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E5%8C%96-%E3%83%97%E3%83%A9%E3%82%A4%E3%83%90%E3%82%B7%E3%83%BC%E3%82%92%E6%95%91%E3%81%A3%E3%81%9F%E5%8F%8D%E4%B9%B1%E8%80%85%E3%81%9F%E3%81%A1-%E3%82%B9%E3%83%86%E3%82%A3%E3%83%BC%E3%83%96%E3%83%B3%E3%83%BB%E3%83%AC%E3%83%93%E3%83%BC/dp/4314009071?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4314009071">暗号化 プライバシーを救った反乱者たち</a></dt>
	<dd>スティーブン・レビー</dd>
	<dd>斉藤 隆央 (翻訳)</dd>
    <dd>紀伊國屋書店 2002-02-16</dd>
    <dd>Book 単行本</dd>
    <dd>ASIN: 4314009071, EAN: 9784314009072</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">20世紀末，暗号技術の世界で何があったのか。知りたかったらこちらを読むべし！</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-16">2018-12-16</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
