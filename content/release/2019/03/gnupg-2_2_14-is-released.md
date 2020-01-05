+++
title = "GnuPG 2.2.14 がリリースされた"
date = "2019-03-20T22:37:08+09:00"
description = "今回もセキュリティ・アップデートはなし。"
image = "/images/attention/openpgp.png"
tags = [
  "security",
  "cryptography",
  "tools",
  "openpgp",
  "gnupg",
]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] 2.2.14 がリリースされた。

- [[Announce] GnuPG 2.2.14 released](https://lists.gnupg.org/pipermail/gnupg-announce/2019q1/000435.html)


今回もセキュリティ・アップデートはなし。
主な機能追加・修正点は以下の通り。

{{% fig-quote type="markdown" title="GnuPG 2.2.14 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2019q1/000435.html" lang="en" %}}
* gpg: Allow import of PGP desktop exported secret keys.  Also avoid importing secret keys if the secret keyblock is not valid.  [#4392]
* gpg: Do not error out on version 5 keys in the local keyring.
* gpg: Make invalid primary key algo obvious in key listings.
* sm: Do not mark a certificate in a key listing as de-vs compliant if its use for a signature will not be possible.
* sm: Fix certificate creation with key on card.
* sm: Create rsa3072 bit certificates by default.
* sm: Print Yubikey attestation extensions with `--dump-cert`.
* agent: Fix cancellation handling for scdaemon.
* agent: Support `--mode=ssh` option for CLEAR_PASSPHRASE.  [#4340]
* scd: Fix flushing of the CA-FPR DOs in app-openpgp.
* scd: Avoid a conflict error with the "undefined" app.
* dirmngr: Add CSRF protection exception for protonmail.
* dirmngr: Fix build problems with gcc 9 in libdns.
* gpgconf: New option `--show-socket` for use wity `--launch`.
* gpgtar: Make option `-C` work for archive creation.

Release-info: https://dev.gnupg.org/T4412
{{% /fig-quote %}}

最新版をインストールすると以下のようになる。

```text
$ gpg --version
gpg (GnuPG) 2.2.14
libgcrypt 1.8.4
Copyright (C) 2019 Free Software Foundation, Inc.
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
[Libgcrypt]: https://gnupg.org/software/libgcrypt/

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
