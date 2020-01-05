+++
title = "GnuPG 2.2.15 がリリースされた"
date = "2019-03-27T23:11:03+09:00"
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

[GnuPG] 2.2.15 がリリースされた。

- [[Announce] GnuPG 2.2.15 released](https://lists.gnupg.org/pipermail/gnupg-announce/2019q1/000436.html)


今回もセキュリティ・アップデートはなし。
主な機能追加・修正点は以下の通り。

{{% fig-quote type="markdown" title="GnuPG 2.2.15 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2019q1/000436.html" lang="en" %}}
* sm: Fix `--logger-fd` and `--status-fd` on Windows for non-standard file descriptors.
* sm: Allow decryption even if expired keys are configured.  [#4431]
* agent: Change command KEYINFO to print ssh fingerprints with other hash algos.
* dirmngr: Fix build problems on Solaris due to the use of reserved symbol names.  [#4420]
* wkd: New commands `--print-wkd-hash` and `--print-wkd-url` for gpg-wks-client.

Release-info: https://dev.gnupg.org/T4434
{{% /fig-quote %}}

最新版をインストールすると以下のようになる。

```text
$ gpg --version
gpg (GnuPG) 2.2.15
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

ついでだが GPGME と Gpg4win もアップデートが出ている。

- [[Announce] GnuPG Made Easy (GPGME) 1.13.0 released](https://lists.gnupg.org/pipermail/gnupg-announce/2019q1/000437.html)
- [Gpg4win version 3.1.6 is released!](https://files.gpg4win.org/README-3.1.6.en.txt)

アップデートは計画的に。

## ブックマーク

- [OpenPGP の実装]({{< ref "/openpgp" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Libgcrypt]: https://gnupg.org/software/libgcrypt/

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
