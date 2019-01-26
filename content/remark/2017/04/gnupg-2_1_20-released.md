+++
draft = false
tags = [
  "security",
  "cryptography",
  "tools",
  "openpgp",
  "gnupg",
]
description = "今回もセキュリティ・アップデートはなし。"
date = "2017-04-04T20:30:11+09:00"
update = "2017-04-04T20:45:16+09:00"
title = "GnuPG 2.1.20 がリリース"

[author]
  avatar = "/images/avatar.jpg"
  flattr = ""
  linkedin = "spiegelimspiegel"
  flickr = "spiegel"
  instagram = "spiegel_2007"
  github = "spiegel-im-spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  license = "by-sa"
  facebook = "spiegel.im.spiegel"
  url = "https://baldanders.info/spiegel/profile/"
  name = "Spiegel"
+++

GnuPG 2.1.20 がリリースされている。

- [[Announce] GnuPG 2.1.20 released](https://lists.gnupg.org/pipermail/gnupg-announce/2017q2/000404.html)

今回もセキュリティ・アップデートはなし。
主な修正点は以下の通り。

* gpg: New properties '`expired`', '`revoked`', and '`disabled`' for the import and export filters.
* gpg: New command `--quick-set-primary-uid`.
* gpg: New compliance field for the `--with-colon` key listing.
* gpg: Changed the key parser to generalize the processing of local meta data packets.
* gpg: Fixed assertion failure in the TOFU trust model.
* gpg: Fixed exporting of zero length user ID packets.
* scd: Improved support for multiple readers.
* scd: Fixed timeout handling for key generation.
* agent: New option `--enable-extended-key-format`.
* dirmngr: Do not add a keyserver to a new dirmngr.conf.  Dirmngr uses a default keyserver.
* dimngr: Do not treat TLS warning alerts as severe error when building with GNUTLS.
* dirmngr: Actually take `/etc/hosts` in account.
* wks: Fixed client problems on Windows.  Published keys are now set to world-readable.
* tests: Fixed creation of temporary directories.
* A socket directory for a non standard GNUGHOME is now created on the fly under `/run/user`.  Thus "`gpgconf --create-socketdir`" is now optional.  The use of "`gpgconf --remove-socketdir`" to clean up obsolete socket directories is however recommended to avoid cluttering `/run/user` with useless directories.
* Fixed build problems on some platforms.

```text
$ gpg --version
gpg (GnuPG) 2.1.20
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
