+++
title = "GnuPG 2.3.0 がリリースされた"
date =  "2021-04-11T13:33:09+09:00"
description = "2.3 系は devel の位置付け。最終的には 2.4 系安定版としてリリース予定。2.2 系でもあと3年は戦える（笑）"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] 2.3.0 がリリースされた。

- [[Announce] GnuPG 2.3.0 released](https://lists.gnupg.org/pipermail/gnupg-announce/2021q2/000458.html)
- [Using a TPM with GnuPG 2.3](https://gnupg.org/blog/20210315-using-tpm-with-gnupg-2.3.html)

[GnuPG] 2.3 系はパブリック・テスト・リリースとして位置付けられ，最終的には 2.4 系安定版としてリリースされる予定である。

{{< fig-quote type="markdown" title="GnuPG 2.3.0 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2021q2/000458.html" lang="en" >}}
{{% quote %}}We are pleased to announce the availability of a new GnuPG release: version 2.3.0.  This release marks the start of public testing releases eventually leading to a new stable version 2.4{{% /quote %}}.
{{< /fig-quote >}}

このため 2.3 系ではバグやリグレッションを頻繁に含む可能性があるので，クリティカルな運用では使用しないほうがいいだろう。

{{< fig-quote type="markdown" title="GnuPG 2.3.0 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2021q2/000458.html" lang="en" >}}
{{% quote %}}Although some bugs might linger in the 2.3 versions, they are intended to replace the 2.2 series.  2.3 may even be used for production purposes if either the risk of minor regressions is acceptable or the new features are important{{% /quote %}}.
{{< /fig-quote >}}

個人的には [GnuPG] 2.3 系は運用しない方向で行く。
試す暇も環境もないので。
その辺の余裕ができれば挑戦してみたいところではあるのだが。

現在のリリース状況は以下の通り。

|    # | パッケージ名                                             | バージョン    | 公開日     |         更新          |
| ---: | -------------------------------------------------------- | ------------- | ---------- | :-------------------: |
|    1 | [Libgpg-error](https://gnupg.org/software/libgpg-error/) | 1.42          | 2021-03-22 | {{< icons "check" >}} |
|    2 | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.8.7 (LTS)   | 2020-10-23 |                       |
|      | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.9.2         | 2021-02-17 | {{< icons "check" >}} |
|    3 | [Libassuan](https://gnupg.org/software/libassuan/)       | 2.5.4         | 2020-10-23 |                       |
|    4 | [Libksba](https://gnupg.org/software/libksba/)           | 1.5.1         | 2021-04-06 | {{< icons "check" >}} |
|    5 | [nPth](https://gnupg.org/software/npth/)                 | 1.6           | 2018-07-16 |                       |
|    6 | [ntbTLS](https://gnupg.org/software/ntbtls/)             | 0.2.0         | 2020-08-27 |                       |
|    7 | [GnuPG](https://gnupg.org/software/)                     | 2.2.27 (LTS)  | 2021-01-11 |                       |
|      | [GnuPG](https://gnupg.org/software/)                     | 2.3.0 (devel) | 2021-04-07 | {{< icons "check" >}} |

[GnuPG](https://gnupg.org/software/) および [Libgcrypt](https://gnupg.org/software/libgcrypt/) のライフサイクルは以下のようになっているみたい。

| Package                                            | Ver. | End-of-life | Birth      |
| -------------------------------------------------- | ---: | ----------- | ---------- |
| [GnuPG](https://gnupg.org/software/)               |  1.4 | none        | 2004-12-16 |
|                                                    |  2.2 | 2024-12-31  | 2014-11-06 |
|                                                    |  2.3 | TBA         | 2021-04-07 |
| [Libgcrypt](https://gnupg.org/software/libgcrypt/) |  1.8 | 2024-12-31  | 2017-07-18 |
|                                                    |  1.9 | 2024-03-31  | 2021-01-19 |

2.2 系でもあと3年は戦える，ということで（笑）

## ブックマーク

- [OpenPGP の実装]({{< rlnk "/openpgp/" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenPGP]: http://openpgp.org/

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
