+++
title = "GnuPG 2.2.9 がリリース"
date = "2018-08-18T15:43:11+09:00"
description = "今回はセキュリティ・アップデートはなし。"
image = "/images/attention/tools.png"
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

1ヶ月以上前の話で恐縮だが [GnuPG] 2.2.9 がリリースされた。

- [[Announce] GnuPG 2.2.9 released](https://lists.gnupg.org/pipermail/gnupg-announce/2018q3/000427.html)

今回はセキュリティ・アップデートはなし。
主な修正点は以下の通り。

{{% fig-quote type="markdown" title="GnuPG 2.2.9 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2018q3/000427.html" lang="en" %}}
* dirmngr: Fix recursive resolver mode and other bugs in the libdns code.  [#3374,#3803,#3610] 
* dirmngr: When using libgpg-error 1.32 or later a GnuPG build with NTBTLS support (e.g. the standard Windows installer) does not anymore block for dozens of seconds before returning data.  If you still have problems on Windows, please consider to use one of the options `disable-ipv4` or `disable-ipv6`.
* gpg: Fix bug in `--show-keys` which actually imported revocation certificates.  [#4017]
* gpg: Ignore too long user-ID and comment packets.  [#4022]
* gpg: Fix crash due to bad German translation.  Improved printf format compile time check.
* gpg: Handle missing `ISSUER` sub packet gracefully in the presence of the new `ISSUER_FPR`.  [#4046]
* gpg: Allow decryption using several passphrases in most cases. [#3795,#4050]
* gpg: Command `--show-keys` now enables the list options `show-unusable-uids`, `show-unusable-subkeys`, `show-notations` and `show-policy-urls` by default.
* gpg: Command `--show-keys` now prints revocation certificates. [#4018]
* gpg: Add revocation reason to the "rev" and "rvs" records of the option `--with-colons`.  [#1173]
* gpg: Export option `export-clean` does now remove certain expired subkeys; `export-minimal` removes all expired subkeys.  [#3622]
* gpg: New "usage" property for the `drop-subkey` filters.  [#4019]
{{% /fig-quote %}}

最新版をインストールすると以下のようになる。

```text
$ gpg --version
gpg (GnuPG) 2.2.9
libgcrypt 1.8.3
Copyright (C) 2018 Free Software Foundation, Inc.
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

## ブックマーク

- [OpenPGP の実装]({{< ref "/openpgp" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Gpg4win]: https://www.gpg4win.org/ "Gpg4win - Secure email and file encryption with GnuPG for Windows"

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
