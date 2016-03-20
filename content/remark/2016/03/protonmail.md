+++
date = "2016-03-20T11:45:24+09:00"
description = "というわけで ProtonMail のアカウントを作った。よさげなら本サイトの連絡先アドレスとして使おうかと思っている。"
draft = false
tags = ["security", "privacy", "cryptography", "tools", "messaging", "openpgp"]
title = "ProtonMail のアカウントを作りました"

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
[OpenPGP はセキュリティやプライバシーに関心のない人からは「面倒」と思われている]({{< relref "remark/2015/use-the-signal-luke.md" >}})らしいが，この辺のギャップを [ProtonMail] が埋めてくれることを期待したい[^m]。

[^m]: 電子メールを含む全てのメッセージング・サービスは，プライバシー要件の有無にかかわらず，全て End-to-End で暗号化すべきだし，そうでないサービス（Gmail 等も含めて）は将来的には排除すべきと思う。ただ，現状としてはメッセージング・サービスには本当に個人的なやり取りから広告や Phishing まであらゆるメッセージが乗ってしまう。[暗号化されたメッセージから悪意を排除するのは難しい]({{< relref "remark/2015/cipher-risk.md" >}})。

- [なぜジョニーは今もやっぱり暗号化できないのか：現在のPGPクライアントの使いやすさ評価 - YAMDAS現更新履歴](http://d.hatena.ne.jp/yomoyomo/20151112/jonnycantencrypt)

[ProtonMail]: https://protonmail.com/ "Secure email: ProtonMail is free encrypted email."

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4314009071/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51ZRZ62WKCL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4314009071/baldandersinf-22/">暗号化 プライバシーを救った反乱者たち</a></dt><dd>スティーブン・レビー 斉藤 隆央 </dd><dd>紀伊國屋書店 2002-02-16</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/487593100X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/487593100X.09._SCTHUMBZZZ_.jpg"  alt="ハッカーズ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4105393022/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4105393022.09._SCTHUMBZZZ_.jpg"  alt="暗号解読―ロゼッタストーンから量子暗号まで"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4484111160/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4484111160.09._SCTHUMBZZZ_.jpg"  alt="グーグル ネット覇者の真実 追われる立場から追う立場へ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/410215972X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/410215972X.09._SCTHUMBZZZ_.jpg"  alt="暗号解読〈上〉 (新潮文庫)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4102159738/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4102159738.09._SCTHUMBZZZ_.jpg"  alt="暗号解読 下巻 (新潮文庫 シ 37-3)"  /></a> </p>
<p class="description">20世紀末，暗号技術の世界で何があったのか。知りたかったらこちらを読むべし！</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-03-09">2015/03/09</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
