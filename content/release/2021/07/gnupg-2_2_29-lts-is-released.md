+++
title = "GnuPG 2.2.29 (LTS) がリリースされた"
date =  "2021-07-07T12:30:18+09:00"
description = "セキュリティ・アップデートはなし。"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] 2.2.29 がリリースされた。

- [[Announce] GnuPG 2.2.29 (LTS) released](https://lists.gnupg.org/pipermail/gnupg-announce/2021q3/000461.html)

セキュリティ・アップデートはなし。
詳細はこちら。

{{< fig-quote type="markdown" title="GnuPG 2.2.29 (LTS) released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2021q3/000461.html" lang="en" >}}
* Fix regression in 2.2.28 for Yubikey NEO.  [#5487]
* Change the default keyserver to `keyserver.ubuntu.com`.  This is a temporary change due to the shutdown of the SKS keyserver pools. [47c4e3e00a]
* gpg: Let `--fetch-key` return an exit code on failure.  [#5376]
* dirmngr: Fix regression in `KS_GET` for mail address pattern. [#5497]
* Add fallback in case the Windows console can't cope with Unicode. [#5491]
* Improve initialization of SPR532 in the CCID driver and make the driver more robust.  [#5297,b90c55fa66db]
* Make test suite work in presence of a broken Libgcrypt installation.  [#5502]
* Make configure option `--disable-ldap` work again.

Release-info: https://dev.gnupg.org/T5498
{{< /fig-quote >}}

んー。
鍵サーバの既定が `keyserver.ubuntu.com` になってるな。
一時的な措置みたいだけど。

現在 [GnuPG] には2.2系と2.3系があり，2.2系は LTS 版に位置づけられている。
2.3系では AEAD (Authenticated Encryption with Associated Data) 等 [RFC 4880bis] で検討されている機能が実装されているので，最新機能を試したいのであればこちらを入れるとよいだろう。
通常運用であれば2.2系で問題ない（[ECC も対応]({{< ref "/openpgp/using-ecc-with-gnupg.md" >}} "そろそろ GnuPG でも ECC を標準で使うのがいいんじゃないかな")してるよ）。

[自前でビルド]({{< ref "/openpgp/build-gnupg-in-ubuntu.md" >}} "Ubuntu で最新版 GnuPG をビルドする")する際の対象パッケージは以下の通り。

|    # | パッケージ名                                             | バージョン   | 公開日     |         更新          |
| ---: | -------------------------------------------------------- | ------------ | ---------- | :-------------------: |
|    1 | [Libgpg-error](https://gnupg.org/software/libgpg-error/) | 1.42         | 2021-03-22 |                       |
|    2 | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.8.8 (LTS)  | 2021-06-02 |                       |
|      | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.9.3        | 2021-04-19 |                       |
|    3 | [Libassuan](https://gnupg.org/software/libassuan/)       | 2.5.5        | 2021-03-22 |                       |
|    4 | [Libksba](https://gnupg.org/software/libksba/)           | 1.6.0        | 2021-06-10 |                       |
|    5 | [nPth](https://gnupg.org/software/npth/)                 | 1.6          | 2018-07-16 |                       |
|    6 | [ntbTLS](https://gnupg.org/software/ntbtls/)             | 0.2.0        | 2020-08-27 |                       |
|    7 | [GnuPG](https://gnupg.org/software/)                     | 2.2.29 (LTS) | 2021-07-04 | {{< icons "check" >}} |
|      | [GnuPG](https://gnupg.org/software/)                     | 2.3.1        | 2021-04-20 |                       |

アップデートは計画的に。

## ブックマーク

- [OpenPGP の実装]({{< rlnk "/openpgp/" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenPGP]: http://openpgp.org/
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
