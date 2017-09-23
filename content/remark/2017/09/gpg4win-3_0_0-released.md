+++
title = "GnuPG 2.2 に対応した Gpg4win 3.0.0 がリリース"
date =  "2017-09-22T23:13:02+09:00"
description = "Gpg4win (GNU Privacy Guard for Windows) は GnuPG を含む Windows 用のパッケージ群である。"
tags = [
  "security",
  "cryptography",
  "tools",
  "openpgp",
  "gnupg",
  "windows",
]

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
  highlightjs = false
  mathjax = false
  mermaidjs = false
+++

[GnuPG] 2.2 に対応した [Gpg4win] 3.0.0 がリリースされた。

- [[Gpg4win-announce] Gpg4win 3.0.0 released](http://lists.wald.intevation.org/pipermail/gpg4win-announce/2017-September/000073.html)
    - [README](https://files.gpg4win.org/README-3.0.0.en.txt)
- [Gpg4win - Whats new - Version 3](https://www.gpg4win.org/version3.html)

[Gpg4win] は [GnuPG] を含む Windows 用のパッケージ群で，以下の製品で構成されている。

- [GnuPG] 2.2.1
- [Kleopatra] 3.0.0
- [GPA] 0.9.9
- [GpgOL] 2.0.1
- [GpgEX] 1.0.5

これまで [Gpg4win] は [GnuPG] 2.0.x をベースに製品を提供してきたが，ようやく最新版においついた感じだ。

[Kleopatra] は [GnuPG] の GUI ラッパーで OpenPGP 鍵および X.509 電子証明書の管理を行う。
[KMail] と連動してメールの暗号化や復号を行うこともできるが Windows 用の [KMail] は同梱されていない。

[GPA] も同じく GUI ラッパーで，暗号化や復号を行うことができる。

[GpgOL] は Microsoft Outlook 用の拡張機能で，Outlook 上で OpenPGP/MIME および S/MIME フォーマットによるメールの暗号化や復号を行うことができる。

[GpgEX] は Windows Explorer の拡張機能で， Explorer 上からファイルの暗号化や復号を行うことができる。
個人的に [GpgEX] はかなり使い勝手がいいのでお勧めである。

ちなみに [Gpg4win] 以外で Windows 用で動作する製品は以下の通り。

- メールの暗号化や復号については [Thunderbird](https://www.mozilla.org/thunderbird/)＋[Enigmail](https://addons.mozilla.org/thunderbird/addon/enigmail/) でも行うことができる
- Gmail や Outlook.com といった Web メール用に [Mailvelope](https://www.mailvelope.com/) という製品がある。こちらは [GnuPG] ではなく [OpenPGP.js] を使っている。 Chrome や Firefox の拡張機能として機能し，メールの暗号化や復号を行う。 Google も Chrome 拡張として [E2EMail](https://github.com/e2email-org/e2email "e2email-org/e2email: E2EMail is a simple Chrome application - a Gmail client that exchanges OpenPGP mail.") を公開しているが [Mailvelope](https://www.mailvelope.com/) のほうが先んじている印象だ
- [Git for Windows](https://git-for-windows.github.io/) は commit への電子署名に内部で [GnuPG] を呼び出す（「[Git Commit で OpenPGP 署名を行う]({{< relref "remark/2016/04/git-commit-with-openpgp-signature.md" >}})」を参照）

とまぁ，地味ながら Windows でも [GnuPG] は活躍してますよ，ということで。

## ブックマーク

- [グーグルのメール暗号化Chromeアプリケーション「E2EMail」がオープンソースに - ZDNet Japan](https://japan.zdnet.com/article/35097359/)

- [わかる！ OpenPGP 暗号](http://www.baldanders.info/spiegel/archive/pgpdump/openpgp.shtml)
- [GnuPG Modern Version for Windows ― インストール編]({{< relref "remark/2016/03/using-gnupg-modern-version-1.md" >}})
- [GnuPG Modern Version for Windows ― gpg-agent について]({{< relref "remark/2016/03/using-gnupg-modern-version-2.md" >}})

[GnuPG]: https://www.gnupg.org/ "The GNU Privacy Guard"
[Gpg4win]: https://www.gpg4win.org/ "Gpg4win - Secure email and file encryption with GnuPG for Windows"
[Kleopatra]: http://openpgp.org/software/kleopatra/ "Kleopatra - OpenPGP"
[KMail]: https://www.kde.org/applications/internet/kmail/ "KDE - KMail - Mail Client"
[GPA]: https://gnupg.org/software/gpa/ "GPA - The Gnu Privacy Assistant"
[GpgEX]: https://github.com/gpg/gpgex "gpg/gpgex: GnupG extension for the Windows Explorer"
[GpgOL]: https://wiki.gnupg.org/GpgOL "GpgOL - GnuPG wiki"
[OpenPGP.js]: https://openpgpjs.org/ "OpenPGP.js | OpenPGP JavaScript Implementation"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
