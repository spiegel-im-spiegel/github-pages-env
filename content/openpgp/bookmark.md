+++
title = "OpenPGP に関するブックマーク"
date =  "2017-12-01T15:15:33+09:00"
update = "2018-02-04T21:04:43+09:00"
description = "OpenPGP に関するリンクをブックマークとして集めてみた（時事ネタは除く）。"
image = "/images/attention/openpgp.png"
tags        = [ "openpgp", "gnupg", "bookmark" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP] に関するリンクをブックマークとして集めてみた（時事ネタは除く）。

## Request for Comments

- [RFC 1847 - Security Multiparts for MIME: Multipart/Signed and Multipart/Encrypted](https://tools.ietf.org/html/rfc1847)
- [RFC 3156 - MIME Security with OpenPGP](https://tools.ietf.org/html/rfc3156) (update [RFC 2015])
- [RFC 4880 - OpenPGP Message Format](https://tools.ietf.org/html/rfc4880) (obsoletes [RFC 1991] and [RFC 2440])
    - [RFC 5581 - The Camellia Cipher in OpenPGP](https://tools.ietf.org/html/rfc5581) (update [RFC 4880])
    - [RFC 6637 - Elliptic Curve Cryptography (ECC) in OpenPGP](https://tools.ietf.org/html/rfc6637) (update [RFC 4880])

- [draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format](https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/)

[RFC 1847]: https://tools.ietf.org/html/rfc1847 "RFC 1847 - Security Multiparts for MIME: Multipart/Signed and Multipart/Encrypted"
[RFC 1991]: https://tools.ietf.org/html/rfc1991 "RFC 1991 - PGP Message Exchange Formats"
[RFC 2015]: https://tools.ietf.org/html/rfc2015 "RFC 2015 - MIME Security with Pretty Good Privacy (PGP)"
[RFC 2440]: https://tools.ietf.org/html/rfc2440 "RFC 2440 - OpenPGP Message Format"
[RFC 3156]: https://tools.ietf.org/html/rfc3156 "RFC 3156 - MIME Security with OpenPGP"
[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 5581]: https://tools.ietf.org/html/rfc5581 "RFC 5581 - The Camellia Cipher in OpenPGP"
[RFC 6637]: https://tools.ietf.org/html/rfc6637 "RFC 6637 - Elliptic Curve Cryptography (ECC) in OpenPGP"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"

## [OpenPGP] の実装および関連ツール

- [Encryption Solutions | Symantec](https://www.symantec.com/products/encryption) : PGP は現在 Symantec 社の暗号化ソリューションの一部となっていて，[商標権も Symantec 社が保持](https://www.symantec.com/about/legal/trademark-policies "Our Current Trademarks | Symantec")している
    - [日本語ページ](https://www.symantec.com/ja/jp/products/encryption)
- [The GNU Privacy Guard](https://gnupg.org/)
    - [Gpg4win - Secure email and file encryption with GnuPG for Windows](https://www.gpg4win.org/)
    - [GPG Suite](https://gpgtools.org/) : macOS 用 GPG Suite
- [OpenPGP.js | OpenPGP JavaScript Implementation](http://openpgpjs.org/)
    - [openpgpjs/openpgpjs: OpenPGP implementation for JavaScript](https://github.com/openpgpjs/openpgpjs)
    - [Mailvelope](https://www.mailvelope.com/)
- [APG - Google Play](https://play.google.com/store/apps/details?id=org.thialfihar.android.apg) ([Thialfihar](http://thialfihar.org/))
    - [thialfihar/apg: OpenPGP for Android](https://github.com/thialfihar/apg)
- [OpenKeychain](https://www.openkeychain.org/)
    - [OpenKeychain: Easy PGP - Google Play](https://play.google.com/store/apps/details?id=org.sufficientlysecure.keychain)
    - [open-keychain/open-keychain: OpenKeychain is an OpenPGP implementation for Android.](https://github.com/open-keychain/open-keychain)
- [iPGMail](https://ipgmail.com/) : PGP for iOS
- [Keybase](https://keybase.io/)
    - [kbpgp.js](https://keybase.io/kbpgp)
    - [[メモ]: keybaseとgit/githubでgpg keyを共有する - Qiita](https://qiita.com/joemphilips/items/7e4d2941448807c4d431)
- [pgpdump]
    - [kazu-yamamoto/pgpdump: A PGP packet visualizer](https://github.com/kazu-yamamoto/pgpdump)
    - [PGPdump Interface](http://www.pgpdump.net/) : Web サービス

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[pgpdump]: http://www.mew.org/~kazu/proj/pgpdump/

## [OpenPGP] 鍵サーバ

- [keys.gnupg.net](http://keys.gnupg.net/ "Nebraska Wesleyan University - OpenPGP Keyserver")
- [pgp.mit.edu](https://pgp.mit.edu/ "MIT PGP Key Server")
- [pgp.nic.ad.jp](http://pgp.nic.ad.jp/ "PGP KEYSERVER")
- [pool.sks-keyservers.net](https://sks-keyservers.net/ "SKS Keyservers")
- [keys.mailvelope.com ](https://keys.mailvelope.com/ "Mailvelope Key Server")

## 参考になる（かもしれない） Web ページ

- [openpgp -- Ongoing discussion of OpenPGP issues.](https://www.ietf.org/mailman/listinfo/openpgp) : ここで次期 OpenPGP の議論などが交わされている
- [PGP User's Manual for Windows](http://www.cla-ri.net/pgp/) : 最近は更新されてないが PGP を巡る歴史資料としての価値は高い
- [GNU Privacy Guard講座](https://gnupg.hclippr.com/) : 最近の [GnuPG] に関する日本語の情報はこちらがお薦め
- [GnuPG - ArchWiki](https://wiki.archlinuxjp.org/index.php/GnuPG)
- [はじめての暗号化メール (Thunderbird編)](https://www.jpcert.or.jp/magazine/security/pgpquick.html) by [JPCERT/CC](https://www.jpcert.or.jp/ "JPCERT コーディネーションセンター")
- [CRYPTREC](http://www.cryptrec.go.jp/)
    - [CRYPTREC暗号リスト(電子政府推奨暗号リスト)](http://www.cryptrec.go.jp/list.html)
    - [CRYPTREC暗号の仕様書](http://www.cryptrec.go.jp/method.html)
    - [CRYPTREC報告書](http://www.cryptrec.go.jp/report.html)
- 旧 OpePKSD.org（[Internet Archive](https://web.archive.org/web/20110907063003/http://www.openpksd.org/)）からドキュメントをサルベージしてみた
    - {{< pdf-file title="PGP/GPG インターネットで広く使われている暗号技術（白黒バージョン）" link="http://www.baldanders.info/spiegel/archive/pgpdump/About_PGP_2002.pdf" >}} （2002年）
    - {{< pdf-file title="PGP/GPG インターネットで広く使われている暗号技術（スライド＋ノート）" link="http://www.baldanders.info/spiegel/archive/pgpdump/About_PGP_2002_NOTE.pdf" >}} （2002年）
    - {{< pdf-file title="OpenPGPとPKI" link="http://www.baldanders.info/spiegel/archive/pgpdump/PGP-001.pdf" >}} （2002年）
- [OpenPGP(実際はGnuPGなど)の活用事例 - Qiita](https://qiita.com/tsuyoshi_cho/items/b5b37be4f52f0adb0f1b)

## 自作ツール

- [spiegel-im-spiegel/BkGnuPG: GNU Privacy Guard Plug-in for Becky! 2](https://github.com/spiegel-im-spiegel/BkGnuPG) : 過去の技術的負債（笑）。（セキュリティ的に）危ないので使ってはいけません
- [spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer](https://github.com/spiegel-im-spiegel/gpgpdump) : [pgpdump] を Go 言語で relpace してみた

## [本家サイト]の記事

- [わかる！ OpenPGP 暗号 — Baldanders.info](http://www.baldanders.info/spiegel/archive/pgpdump/openpgp.shtml)
- [pgpdump (patched version) — Baldanders.info](http://www.baldanders.info/spiegel/archive/pgpdump/) : [pgpdump] の Windows 用バイナリを公開中
- [PGP/GnuPG 用語一覧](http://www.baldanders.info/spiegel/archive/pgp-vocabulary/) : 作りかけ放置プレイ中
- [OpenPGP で利用可能なアルゴリズム一覧（RFC4880 対応版） — Baldanders.info](http://www.baldanders.info/spiegel/log2/000452.shtml)
- [Google による OpenPGP 鍵配送の解決提案 — Baldanders.info](http://www.baldanders.info/spiegel/log2/000785.shtml)

[OpenPGP]: https://www.openpgp.org/
[本家サイト]: http://www.baldanders.info/ "Baldanders.info"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4314009071/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51ZRZ62WKCL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4314009071/baldandersinf-22/">暗号化 プライバシーを救った反乱者たち</a></dt><dd>スティーブン・レビー 斉藤 隆央 </dd><dd>紀伊國屋書店 2002-02-16</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/487593100X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/487593100X.09._SCTHUMBZZZ_.jpg"  alt="ハッカーズ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4105393022/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4105393022.09._SCTHUMBZZZ_.jpg"  alt="暗号解読―ロゼッタストーンから量子暗号まで"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4484111160/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4484111160.09._SCTHUMBZZZ_.jpg"  alt="グーグル ネット覇者の真実 追われる立場から追う立場へ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/410215972X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/410215972X.09._SCTHUMBZZZ_.jpg"  alt="暗号解読〈上〉 (新潮文庫)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4102159738/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4102159738.09._SCTHUMBZZZ_.jpg"  alt="暗号解読 下巻 (新潮文庫 シ 37-3)"  /></a> </p>
<p class="description">20世紀末，暗号技術の世界で何があったのか。知りたかったらこちらを読むべし！</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-03-09">2015/03/09</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4900900028/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/5132396FFQL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4900900028/baldandersinf-22/">PGP―暗号メールと電子署名</a></dt><dd>シムソン ガーフィンケル Simson Garfinkel </dd><dd>オライリー・ジャパン 1996-04</dd><dd>評価<abbr class="rating" title="3"><img src="http://g-images.amazon.com/images/G/01/detail/stars-3-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4756136494/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4756136494.09._SCTHUMBZZZ_.jpg"  alt="プログラミング作法"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4320026926/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4320026926.09._SCTHUMBZZZ_.jpg"  alt="プログラミング言語C 第2版 ANSI規格準拠"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797350997/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797350997.09._SCTHUMBZZZ_.jpg"  alt="新版暗号技術入門 秘密の国のアリス"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798132608/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798132608.09._SCTHUMBZZZ_.jpg"  alt="情報処理教科書 高度試験午後II論述 春期・秋期 (EXAMPRESS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798105538/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798105538.09._SCTHUMBZZZ_.jpg"  alt="エンタープライズ アプリケーションアーキテクチャパターン (Object Oriented Selection)"  /></a> </p>
<p class="description" >良書なのだが，残念ながら内容が古すぎた。 PGP の歴史資料として読むならいいかもしれない。</p>
<p class="gtools" >reviewed by <a href="#maker" class="reviewer">Spiegel</a> on <abbr class="dtreviewed" title="2014-10-16">2014/10/16</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html">G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
