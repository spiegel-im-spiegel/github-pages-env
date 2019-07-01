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
  url = "https://baldanders.info/profile/"
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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE"><img src="https://images-fe.ssl-images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE">暗号技術入門 第3版　秘密の国のアリス</a></dt>
	<dd>結城 浩</dd>
    <dd>SBクリエイティブ 2015-08-25 (Release 2015-09-17)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B015643CPE</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
