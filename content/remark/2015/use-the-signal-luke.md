+++
date = "2015-11-14T06:25:40+09:00"
description = "TextSecure は Signal に統合された。Signal は SMS を置き換えることができ，ローカルストレージにあるログもちゃんと暗号化される。オススメである。"
draft = false
tags = ["security", "cryptography", "messaging", "risk", "openpgp", "tools", "signal", "otr"]
title = "ルークよ， Signal を使え！"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"
+++

「[なぜジョニーは今もやっぱり暗号化できないのか：現在のPGPクライアントの使いやすさ評価 - YAMDAS現更新履歴](http://d.hatena.ne.jp/yomoyomo/20151112/jonnycantencrypt)」経由。

- {{< pdf-file title="Why Johnny Still, Still Can’t Encrypt: Evaluating the Usability of a Modern PGP Client" link="http://arxiv.org/pdf/1510.08555.pdf" >}}
- [Why do encryption tools suck? / Boing Boing](http://boingboing.net/2015/11/06/why-do-encryption-tools-suck.html)

タイトルだけ見て最初に思ったのは「PGP ってゆーな[^a]」だったが，それはともかく

[^a]: 昔はともかく，現在 [PGP] は Symantic 社の製品名であり，いくつかある [OpenPGP] 実装のひとつにすぎない。余談だが，確かに [OpenPGP] 用の「公開鍵サーバ」というのは存在しているが， [GnuPG] も [PGP] も，今回紹介する [Mailvelope] も別にクライアントというわけではなく，スタンドアロンで動作するプログラムである。

{{< fig-quote title="Why Johnny Still, Still Can’t Encrypt: Evaluating the Usability of a Modern PGP Client" link="http://arxiv.org/pdf/1510.08555.pdf" lang="en" >}}
<q>In our study of 20 participants, grouped into 10 pairs of participants who attempted to exchange encrypted email, only one pair was able to successfully complete the assigned tasks using Mailvelope. All other participants were unable to complete the assigned task in the one hour allotted to the study.
This demonstrates that encrypting email with PGP, as implemented in Mailvelope, is still unusable for the masses.</q>
{{< /fig-quote >}}

10% 以下かよ orz

1990年代からカジュアルに PGP を使ってて，あまつさえ [Becky! 用のプラグイン](http://hp.vector.co.jp/authors/VA023900/gpg-pin/)まで自作してしまった[^a2] 身としては，かなり衝撃的ではあったが，まぁそんなもんかなぁ。

[^a2]: 現在，このプラグインはメンテナンスされてない。使わないように。 Windows パソコンでメールの暗号化がしたいなら [Thunderbird](https://www.mozilla.org/thunderbird/)＋[Enigmail](https://addons.mozilla.org/thunderbird/addon/enigmail/) がオススメである。

ちなみに [Mailvelope] は Chrome や Firefox の拡張機能として使える MUA（Mail User Agent）で， Gmail などと連携して使える。
[以前にも紹介](https://baldanders.info/spiegel/log2/000782.shtml "安全なメッセージング・アプリとは（追記あり） — Baldanders.info")したが， [EFF の Secure Messaging Scorecard](https://www.eff.org/secure-messaging-scorecard) で MUA の中では（PFS の項目を除いて[^b]）唯一満点だった製品である。
[Mailvelope] は [OpenPGP] 実装を含んでいるため [GnuPG] や [PGP] のような製品を必要としない利点がある。

[^b]: [前にも書いた](https://baldanders.info/spiegel/log2/000782.shtml "安全なメッセージング・アプリとは（追記あり） — Baldanders.info")が， MUA はその性質上 PFS（Perfect Forward Secrecy）を満たせない。 PFS は通信経路の暗号化や IM（Instant Messaging）など使い捨てのメッセージを扱う場合には必要な要件だが，電子メールには向いていない。電子メールは過去のやり取りに対して「否認防止（non-repudiation）」できなければならないからだ。

- [mailvelopeの使い方 | セキュリティの意識と知識](http://security.hondaclinic.jp/%E6%9A%97%E5%8F%B7%E3%81%AE%E3%81%99%E3%81%99%E3%82%81/mailvelope%E3%81%AE%E4%BD%BF%E3%81%84%E6%96%B9/)

この論文については Bruce Schneier さんも言及している。

{{< fig-quote title="Testing the Usability of PGP Encryption Tools - Schneier on Security" link="https://www.schneier.com/blog/archives/2015/11/testing_the_usa.html" lang="en" >}}
<q>I have recently come to the conclusion that e-mail is fundamentally unsecurable. The things we want out of e-mail, and an e-mail system, are not readily compatible with encryption. I advise people who want communications security to not use e-mail, but instead use an encrypted message client like <a href="https://otr.cypherpunks.ca/">OTR</a> or <a href="https://whispersystems.org/">Signal</a>.</q>
{{< /fig-quote >}}

正論ですね。
そういうことです。

そういえば， [Open Whisper Systems](https://whispersystems.org/) の TextSecure は Signal に統合された。

- [Open Whisper Systems >> Blog >> Just Signal](https://whispersystems.org/blog/just-signal/)
- [Signal Private Messenger - Google Play](https://play.google.com/store/apps/details?id=org.thoughtcrime.securesms)

暗号化通話アプリの RedPhone も Signal に統合された。
Signal は SMS を置き換えることができ，ローカルストレージにあるログもちゃんと暗号化される。
オススメである。

いや，もう，マジ LINE とか捨てて欲しい。
Facebook もくだらない「[セキュリティ劇場](http://www.itmedia.co.jp/news/articles/1511/13/news077.html)」で遊んでないで， OTR くらい組み込めや。

メッセージングの主体が PC からスマホなどの携帯端末にシフトしている状態で電子メールの有効性は下がってきていると思う。
もちろん仕事などではまだまだ PC や Workstation を使うことも多いし，対外的には電子メールが主体になると思うけど，たとえばグループウェアに XMPP＋OTR を組み込んでいくなら，やはり電子メールは無用の長物になっていくだろう。

{{< fig-youtube id="o2we_B6hDrY" width="500" height="375" >}}

[OpenPGP]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[PGP]: https://www.symantec.com/encryption/ "PGP Encryption Software | Symantec"
[GnuPG]: https://www.gnupg.org/ "The GNU Privacy Guard"
[Mailvelope]: https://www.mailvelope.com/ "Mailvelope"

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">第3版出た！ てか，もう Kindle 版出てるのか。紙の本買うのはやまったかなぁ。 SHA-3 や BitCoin/BlockChain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
