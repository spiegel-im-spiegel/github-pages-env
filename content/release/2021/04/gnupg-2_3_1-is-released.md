+++
title = "GnuPG 2.3.1 がリリースされた"
date =  "2021-04-21T21:48:55+09:00"
description = "セキュリティ・アップデートはなし"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] 2.3.1 がリリースされた。

- [[Announce] GnuPG 2.3.1 released](https://lists.gnupg.org/pipermail/gnupg-announce/2021q2/000459.html)

セキュリティ・アップデートはなし。
詳細はこちら。

{{< fig-quote type="markdown" title="GnuPG 2.3.1 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2021q2/000459.html" lang="en" >}}
* The new configuration file common.conf is now used to enable the use of the key database daemon with "`use-keyboxd`".  Using this option in gpg.conf and gpgsm.conf is supported for a transitional period.  See `doc/example/common.conf` for more.
* gpg: Force version 5 key creation for ed448 and cv448 algorithms.
* gpg: By default do not use the `self-sigs-only` option when importing from an LDAP keyserver.  [#5387]
* gpg: Lookup a missing public key of the active card via LDAP. [d7e707170f]
* gpgsm: New command `--show-certs`.  [51419d6341]
* scd: Fix CCID driver for SCM SPR332/SPR532.  [#5297]
* scd: Further improvements for PKCS#15 cards.
* Fix build problems on Fedora.  [#5389]
* Fix build problems on macOS.  [#5400]
* New configure option --with-tss to allow the selection of the TSS library.  [93c88d0af3]

Release-info: https://dev.gnupg.org/T5386
{{< /fig-quote >}}

[GnuPG] 関連の各パッケージのバージョンは以下の通り。

|    # | パッケージ名                                             | バージョン    | 公開日     |         更新          |
| ---: | -------------------------------------------------------- | ------------- | ---------- | :-------------------: |
|    1 | [Libgpg-error](https://gnupg.org/software/libgpg-error/) | 1.42          | 2021-03-22 |                       |
|    2 | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.8.7 (LTS)   | 2020-10-23 |                       |
|      | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.9.3         | 2021-04-19 | {{< icons "check" >}} |
|    3 | [Libassuan](https://gnupg.org/software/libassuan/)       | 2.5.5         | 2021-03-22 | {{< icons "check" >}} |
|    4 | [Libksba](https://gnupg.org/software/libksba/)           | 1.5.1         | 2021-04-06 |                       |
|    5 | [nPth](https://gnupg.org/software/npth/)                 | 1.6           | 2018-07-16 |                       |
|    6 | [ntbTLS](https://gnupg.org/software/ntbtls/)             | 0.2.0         | 2020-08-27 |                       |
|    7 | [GnuPG](https://gnupg.org/software/)                     | 2.2.27 (LTS)  | 2021-01-11 |                       |
|      | [GnuPG](https://gnupg.org/software/)                     | 2.3.1 (devel) | 2021-04-20 | {{< icons "check" >}} |

[GnuPG] 2.3 系は最終的に 2.4 系安定版としてリリースされる予定である。
2.2 系は少なくとも2024年末まではサポートが続けられるので，今すぐ 2.3 系にバージョンアップする必要はない。

けど，どっかの時点で環境を作って試したいなぁ...

アップデートは計画的に。

## ブックマーク

- [Using a TPM with GnuPG 2.3](https://gnupg.org/blog/20210315-using-tpm-with-gnupg-2.3.html)
- [OpenPGP の実装]({{< rlnk "/openpgp/" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenPGP]: http://openpgp.org/

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
