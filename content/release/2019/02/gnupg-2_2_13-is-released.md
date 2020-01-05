+++
title = "GnuPG 2.2.13 がリリースされた"
date = "2019-02-13T22:07:18+09:00"
description = "今回もセキュリティ・アップデートはなし。"
image = "/images/attention/openpgp.png"
tags = [
  "security",
  "cryptography",
  "tools",
  "openpgp",
  "gnupg",
]

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] 2.2.13 がリリースされた。

- [[Announce] GnuPG 2.2.13 released](https://lists.gnupg.org/pipermail/gnupg-announce/2019q1/000434.html)

今回もセキュリティ・アップデートはなし。
主な機能追加・修正点は以下の通り。

{{% fig-quote type="markdown" title="GnuPG 2.2.13 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2019q1/000434.html" lang="en" %}}
* gpg: Implement key lookup via keygrip (using the & prefix).
* gpg: Allow generating Ed25519 key from existing key.
* gpg: Emit an ERROR status line if no key was found with `-k`.
* gpg: Stop early when trying to create a primary Elgamal key.  [#4329]
* gpgsm: Print the card's key algorithms along with their keygrips in interactive key generation.
* agent: Clear bogus pinentry cache in the error case.  [#4348]
* scd: Support "acknowledge button" feature.
* scd: Fix for USB INTERRUPT transfer.  [#4308]
* wks: Do no use compression for the the encrypted challenge and response.

Release-info: https://dev.gnupg.org/T4290
{{% /fig-quote %}}

最新版をインストールすると以下のようになる。

```text
$ gpg --version
gpg (GnuPG) 2.2.13
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
