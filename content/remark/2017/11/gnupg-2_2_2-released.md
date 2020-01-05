+++
title = "GnuPG 2.2.2 がリリースされた"
date =  "2017-11-07T23:39:38+09:00"
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
  url       = "https://baldanders.info/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] 2.2.2 がリリースされた。

- [[Announce] GnuPG 2.2.2 released](https://lists.gnupg.org/pipermail/gnupg-announce/2017q4/000416.html)

今回はセキュリティ・アップデートはなし。
主な修正点は以下の通り。

* gpg: Avoid duplicate key imports by concurrently running gpg processes. [#3446]
* gpg: Fix creating on-disk subkey with on-card primary key. [#3280]
* gpg: Fix validity retrieval for multiple keyrings. [Debian#878812]
* gpg: Fix `--dry-run` and import option show-only for secret keys.
* gpg: Print "sec" or "sbb" for secret keys with import option import-show. [#3431]
* gpg: Make import less verbose. [#3397]
* gpg: Add alias "Key-Grip" for parameter "Keygrip" and new parameter "Subkey-Grip" to unattended key generation.  [#3478]
* gpg: Improve "factory-reset" command for OpenPGP cards.  [#3286]
* gpg: Ease switching Gnuk tokens into ECC mode by using the magic keysize value 25519.
* gpgsm: Fix `--with-colon` listing in crt records for fields > 12.
* gpgsm: Do not expect X.509 keyids to be unique.  [#1644]
* agent: Fix stucked Pinentry when using --max-passphrase-days. [#3190]
* agent: New option `--s2k-count`.  [#3276 (workaround)]
* dirmngr: Do not follow https-to-http redirects. [#3436]
* dirmngr: Reduce default LDAP timeout from 100 to 15 seconds. [#3487]
* gpgconf: Ignore non-installed components for commands `--apply-profile` and `--apply-defaults`. [#3313]
* Add configure option `--enable-werror`.  [#2423]


最新版をインストールすると以下のようになる。

```text
$ gpg --version
gpg (GnuPG) 2.2.2
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

アップデートは計画的に。

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
