+++
title = "GnuPG 2.2.1 がリリース"
date =  "2017-09-20T10:46:45+09:00"
description = "今回はセキュリティ・アップデートはなし。"
tags = [
  "security",
  "cryptography",
  "tools",
  "openpgp",
  "gnupg",
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
+++

[GnuPG] 2.2.1 がリリースされた。

- [[Announce] GnuPG 2.2.1 released](https://lists.gnupg.org/pipermail/gnupg-announce/2017q3/000415.html)

今回はセキュリティ・アップデートはなし。
主な修正点は以下の通り。

* gpg: Fix formatting of the user id in batch mode key generation if only "name-email" is given.
* gpgv: Fix annoying "not suitable for" warnings.
* wks: Convey only the newest user id to the provider.  This is the case if different names are used with the same addr-spec.
* wks: Create a complying user id for provider policy mailbox-only.
* wks: Add workaround for [posteo.de](https://posteo.de/ "Email green, secure, simple and ad-free - posteo.de -").
* scd: Fix the use of large ECC keys with an OpenPGP card.
* dirmngr: Use system provided root certificates if no specific HKP certificates are configured.  If build with GNUTLS, this was already the case.

最新版をインストールすると以下のようになる。

```text
$ gpg --version
gpg (GnuPG) 2.2.1
libgcrypt 1.8.1
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

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
