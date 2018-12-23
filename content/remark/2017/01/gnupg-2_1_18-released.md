+++
title = "GnuPG 2.1.18 がリリース"
tags = [
  "security",
  "cryptography",
  "tools",
  "openpgp",
  "gnupg",
]
description = "今回もセキュリティ・アップデートはなし。"
draft = false
date = "2017-01-25T18:16:00+09:00"

[author]
  instagram = "spiegel_2007"
  github = "spiegel-im-spiegel"
  facebook = "spiegel.im.spiegel"
  flickr = "spiegel"
  license = "by-sa"
  url = "http://www.baldanders.info/spiegel/profile/"
  tumblr = ""
  avatar = "/images/avatar.jpg"
  name = "Spiegel"
  flattr = ""
  twitter = "spiegel_2007"
  linkedin = "spiegelimspiegel"
+++

GnuPG 2.1.18 がリリースされている。

- [[Announce] GnuPG 2.1.18 released](https://lists.gnupg.org/pipermail/gnupg-announce/2017q1/000401.html)

今回もセキュリティ・アップデートはなし。
主な修正点は以下の通り。

* gpg: Remove bogus subkey signature while cleaning a key (with `export-clean`, `import-clean`, or `--edit-key`'s sub-command clean)
* gpg: Allow freezing the clock with `--faked-system-time`.
* gpg: New `--export-option` flag "`backup`", new `--import-option` flag "`restore`".
* gpg-agent: Fixed long delay due to a regression in the progress callback code.
* scd: Lots of code cleanup and internal changes.
* scd: Improved the internal CCID driver.
* dirmngr: Fixed problem with the DNS glue code (removal of the trailing dot in domain names).
* dirmngr: Make sure that Tor is actually enabled after changing the conf file and sending SIGHUP or "`gpgconf --reload dirmngr`".
* dirmngr: Fixed Tor access to IPv6 addresses.  Note that current versions of Tor may require that the flag "`IPv6Traffic`" is used with the option "`SocksPort`" in torrc to actually allow IPv6 traffic.
* dirmngr: Fixed HKP for literally given IPv6 addresses.
* dirmngr: Enabled reverse DNS lookups via Tor.
* dirmngr: Added experimental SRV record lookup for WKD.  See commit `88dc3af3d4ae1afe1d5e136bc4c38bc4e7d4cd10` for details.
* dirmngr: For HKP use "`pgpkey-hkps`" and "`pgpkey-hkp`" in SRV record lookups.  Avoid SRV record lookup when a port is explicitly specified.  This fixes a regression from the 1.4 and 2.0 behavior.
* dirmngr: Gracefully handle a missing `/etc/nsswitch.conf`.  Ignore negation terms (e.g. "`[!UNAVAIL=return]`" instead of bailing out.
* dirmngr: Better debug output for flags "`dns`" and "`network`".
* dirmngr: On reload mark all known HKP servers alive.
* gpgconf: Allow keyword "`all`" for `--launch`, `--kill`, and `--reload`.
* tools: gpg-wks-client now ignores a missing policy file on the server.
* Avoid unnecessary ambiguity error message in the option parsing.
* Further improvements of the regression test suite.
* Fixed building with `--disable-libdns` configure option.
* Fixed a crash running the tests on 32 bit architectures.
* Fixed spurious failures on BSD system in the spawn functions.  This affected for example gpg-wks-client and gpgconf.

んー。
IPv6 や Tor 周りが改善されてるって感じ？

```text
$ gpg --version
gpg (GnuPG) 2.1.18
libgcrypt 1.7.6
Copyright (C) 2017 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Home: ********
サポートしているアルゴリズム:
公開鍵: RSA, ELG, DSA, ECDH, ECDSA, EDDSA
暗号方式: IDEA, 3DES, CAST5, BLOWFISH, AES, AES192, AES256, TWOFISH,
    CAMELLIA128, CAMELLIA192, CAMELLIA256
ハッシュ: SHA1, RIPEMD160, SHA256, SHA384, SHA512, SHA224
圧縮: 無圧縮, ZIP, ZLIB, BZIP2
```

あれ？ Libgcrypt も 1.7.6 になってるなぁ。
[メーリングリスト](https://lists.gnupg.org/mailman/listinfo/gnupg-announce)にアナウンスがなかったぞ。

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
