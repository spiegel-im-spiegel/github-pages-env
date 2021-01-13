+++
title = "GnuPG 2.2.27 がリリースされた"
date =  "2021-01-13T18:41:38+09:00"
description = "主に Windows 環境における不具合の修正のようだ。"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg", "ubuntu", "windows"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] 2.2.27 がリリースされた。

- [[Announce] GnuPG 2.2.27 released](https://lists.gnupg.org/pipermail/gnupg-announce/2021q1/000452.html)

このバージョンは主に Windows 環境における不具合の修正のようだ。
併せて Gpg4win 3.1.15 もリリースされている。

- [[Gpg4win-users-en] [Gpg4win-announce] Gpg4win 3.1.15 released](https://lists.wald.intevation.org/pipermail/gpg4win-users-en/2021-January/001635.html)

詳細はこちら。

{{< fig-quote type="markdown" title="GnuPG 2.2.27 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2021q1/000452.html" lang="en" >}}
* gpg: Fix regression in 2.2.24 for gnupg_remove function under Windows.  [#5230]
* gpgconf: Fix case with neither local nor global gpg.conf.  [9f37d3e6f3]
* gpgconf: Fix description of two new options.  [#5221]
* Build Windows installer without timestamps.  Note that the Authenticode signatures still carry a timestamp.

Release-info: https://dev.gnupg.org/T5234
{{< /fig-quote >}}

[自前でビルド]({{< ref "/openpgp/build-gnupg-in-ubuntu.md" >}} "Ubuntu で最新版 GnuPG をビルドする")する際の対象パッケージは以下の通り。

|   # | パッケージ名                                             | バージョン | 公開日     |         更新          |
| ---:| -------------------------------------------------------- | ---------- | ---------- |:---------------------:|
|   1 | [Libgpg-error](https://gnupg.org/software/libgpg-error/) | 1.41       | 2020-12-21 |                       |
|   2 | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.8.7      | 2020-10-23 |                       |
|   3 | [Libassuan](https://gnupg.org/software/libassuan/)       | 2.5.4      | 2020-10-23 |                       |
|   4 | [Libksba](https://gnupg.org/software/libksba/)           | 1.5.0      | 2020-11-18 |                       |
|   5 | [nPth](https://gnupg.org/software/npth/)                 | 1.6        | 2018-07-16 |                       |
|   6 | [ntbTLS](https://gnupg.org/software/ntbtls/)             | 0.2.0      | 2020-08-27 |                       |
|   7 | [GnuPG](https://gnupg.org/software/)                     | 2.2.27     | 2021-01-11 | {{< icons "check" >}} |

`gpg-agent` を Linux のサービスで起動している場合は

```text
$ systemctl --user restart gpg-agent
```

とサービスを再起動するのを忘れずに。
これで

```text
$ gpg --version
gpg (GnuPG) 2.2.27
libgcrypt 1.8.7
Copyright (C) 2021 Free Software Foundation, Inc.
License GNU GPL-3.0-or-later <https://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Home: /home/username/.gnupg
サポートしているアルゴリズム:
公開鍵: RSA, ELG, DSA, ECDH, ECDSA, EDDSA
暗号方式: IDEA, 3DES, CAST5, BLOWFISH, AES, AES192, AES256,
      TWOFISH, CAMELLIA128, CAMELLIA192, CAMELLIA256
ハッシュ: SHA1, RIPEMD160, SHA256, SHA384, SHA512, SHA224
圧縮: 無圧縮, ZIP, ZLIB, BZIP2
```

となる。

アップデートは計画的に。

## ブックマーク

- [OpenPGP の実装]({{< rlnk "/openpgp/" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenPGP]: http://openpgp.org/

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
