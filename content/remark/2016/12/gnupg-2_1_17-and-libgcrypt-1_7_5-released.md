+++
date = "2016-12-26T20:45:46+09:00"
update = "2016-12-27T06:46:48+09:00"
title = "GnuPG 2.1.17 and Libgcrypt 1.7.5 released"
tags = [
  "security",
  "cryptography",
  "tools",
  "openpgp",
  "gnupg",
]
draft = false
description = "Libgcrypt はバグフィックスがメイン， GnuPG 2.1 系 はちょこちょこと機能変更が行われている。"

[author]
  flickr = "spiegel"
  twitter = "spiegel_2007"
  tumblr = ""
  avatar = "/images/avatar.jpg"
  url = "http://www.baldanders.info/spiegel/profile/"
  license = "by-sa"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  facebook = "spiegel.im.spiegel"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  flattr = "spiegel"
+++

Libgcrypt 1.7.5 および GnuPG 2.1.17 がリリースされている。
Libgcrypt はバグフィックスがメイン， GnuPG 2.1 系 はちょこちょこと機能変更が行われている。

- [[Announce] Libgcrypt 1.7.5 released](https://lists.gnupg.org/pipermail/gnupg-announce/2016q4/000399.html)
- [[Announce] GnuPG 2.1.17 released](https://lists.gnupg.org/pipermail/gnupg-announce/2016q4/000400.html)

以下は GnuPG 2.1.17 の変更履歴。

* gpg: By default new keys expire after 2 years.
* gpg: New command `--quick-set-expire` to conveniently change the expiration date of keys.
* gpg: Option and command names have been changed for easier comprehension.  The old names are still available as aliases.
* gpg: Improved the TOFU trust model.
* gpg: New option `--default-new-key-algo`.
* scd: Support OpenPGP card V3 for RSA.
* dirmngr: Support for the ADNS library has been removed.  Instead William Ahern's Libdns is now source included and used on all platforms.  This enables Tor support on all platforms.  The new option `--standard-resolver` can be used to disable this code at runtime.  In case of build problems the new configure option `--disable-libdns` can be used to build without Libdns.
* dirmngr: Lazily launch ldap reaper thread.
* tools: New options `--check` and `--status-fd` for gpg-wks-client.
* The UTF-8 byte order mark is now skipped when reading conf files.
* Fixed many bugs and regressions.
* Major improvements to the test suite.  For example it is possible to run the external test suite of GPGME.

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
