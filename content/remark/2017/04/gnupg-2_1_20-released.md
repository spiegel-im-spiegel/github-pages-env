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
  url = "https://baldanders.info/profile/"
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

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
