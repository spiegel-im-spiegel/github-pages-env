+++
title = "GnuPG 2.2.31 (LTS) がリリースされた"
date =  "2021-09-16T20:39:59+09:00"
description = "セキュリティ・アップデートはなし。"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] LTS 版 2.2.31 がリリースされた。

- [[Announce] GnuPG 2.2.31 (LTS) released](https://lists.gnupg.org/pipermail/gnupg-announce/2021q3/000464.html)

セキュリティ・アップデートはなし。
詳細はこちら。

{{< fig-quote type="markdown" title="GnuPG 2.2.31 (LTS) released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2021q3/000464.html" lang="en" >}}
* agent: Fix a regression in `GET_PASSPHRASE`.  [#5577]
* scd: Fix an assertion failure in `close_pcsc_reader`.  [67e1834ad4]
* scd: Add support for PC/SC in "`GETINFO reader_list`".

Release-info: https://dev.gnupg.org/T5571
{{< /fig-quote >}}

[GnuPG] 関連の各パッケージのバージョンは以下の通り（数字は大体の[ビルド]({{< ref "/openpgp/build-gnupg-in-ubuntu.md" >}} "Ubuntu で最新版 GnuPG をビルドする")順）。

|    # | パッケージ名                                             | バージョン   | 公開日     |         更新          |
| ---: | -------------------------------------------------------- | ------------ | ---------- | :-------------------: |
|    1 | [Libgpg-error](https://gnupg.org/software/libgpg-error/) | 1.42         | 2021-03-22 |                       |
|    2 | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.8.8 (LTS)  | 2021-06-02 |                       |
|      | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.9.4        | 2021-08-22 |                       |
|    3 | [Libassuan](https://gnupg.org/software/libassuan/)       | 2.5.5        | 2021-03-22 |                       |
|    4 | [Libksba](https://gnupg.org/software/libksba/)           | 1.6.0        | 2021-06-10 |                       |
|    5 | [nPth](https://gnupg.org/software/npth/)                 | 1.6          | 2018-07-16 |                       |
|    6 | [ntbTLS](https://gnupg.org/software/ntbtls/)             | 0.2.0        | 2020-08-27 |                       |
|    7 | [GnuPG](https://gnupg.org/software/)                     | 2.2.31 (LTS) | 2021-09-15 | {{< icons "check" >}} |
|      | [GnuPG](https://gnupg.org/software/)                     | 2.3.2        | 2021-08-24 |                       |

現在 [GnuPG] には2.2系と2.3系があり[^gpg14]，2.2系は LTS 版に位置付けられている。
2.3系では AEAD (Authenticated Encryption with Associated Data) 等 [RFC 4880bis] で検討されている機能が実装されているので，最新機能を試したいのであればこちらを入れるとよいだろう。
なお2.2系は少なくとも2024年末まではサポートが続けられる予定である。
通常運用であれば，当面は2.2系でも問題ない（[ECC も対応]({{< ref "/openpgp/using-ecc-with-gnupg.md" >}} "そろそろ GnuPG でも ECC を標準で使うのがいいんじゃないかな")してるよ）。

[^gpg14]: 厳密には1.4系もあるが，これは legacy 版と位置付けられており，よほどのバグか脆弱性がない限りは更新されない。もし今だに1.4系（あるいは既にサポートされていない2.0/2.1系）を使っているのなら2.2系以降にアップグレードすることを強くお勧めする。

アップデートは計画的に。

## ブックマーク

- [OpenPGP の実装]({{< rlnk "/openpgp/" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenPGP]: http://openpgp.org/
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
