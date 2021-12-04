+++
title = "GnuPG 2.2.33 (LTS) がリリースされた"
date =  "2021-12-04T15:16:28+09:00"
description = "セキュリティ・アップデートはなし。"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

遅まきながらで申し訳ないが [GnuPG] LTS 版 2.2.33 がリリースされている。

- [[Announce] GnuPG 2.2.33 (LTS) released](https://lists.gnupg.org/pipermail/gnupg-announce/2021q4/000467.html)

セキュリティ・アップデートはなし。
詳細はこちら。
かなりあるぞ。

{{< fig-quote type="markdown" title="GnuPG 2.2.33 (LTS) released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2021q4/000467.html" lang="en" >}}
* gpg: New option `--min-rsa-length`.  [rG6ee01c1d26]
* gpg: New option `--forbid-gen-key`.  [rG985fb25c46]
* gpg: New option `--override-compliance-check`.  [T5655]
* gpgconf: New command `--show-configs`.  [rG8fe3f57643]
* agent,dirmngr: New option `--steal-socket`.  [rG6507c6ab10]
* scd: Improve the selection of the default PC/SC reader.  [T5644]
* gpg: Fix printing of binary notations.  [T5667]
* gpg: Remove stale ultimately trusted keys from the trustdb.  [T5685]
* gpgsm: Detect circular chains in `--list-chain`.  [rGc9343bec83]
* gpgconf: Create the local option file even if the global file exists.  [T5650]
* dirmngr: Make reading resolv.conf more robust.  [T5657]
* gpg-wks-server: Fix created file permissions.  [rGf54feb4470]
* scd: Support longer data for ssh-agent authentication with openpgp cards.  [T5682]
* Support `gpgconf.ctl` for NetBSD and Solaris.  [T5656,T5671]
* Silence "Garbled console data" warning under Windows in most cases.
* Silence warning about the rootdir under Unices w/o a mounted /proc file system.
* Fix possible build problems about missing include files.  [T5592]
* i18n: Replace the term "PIN-Cache" by "Passswort-Cache" in the German translation. [rgf453d52e53]
* i18n: Update the Russian translation.

Release-info: https://dev.gnupg.org/T5641
{{< /fig-quote >}}

おっ NetBSD と Solaris でも `gpgconf.ctl` をサポートするのか。
`gpgconf.ctl` については以下を参照のこと。

- [GnuPG の HOME はどこにある？]({{< ref "/openpgp/gnupg-home-in-windows.md" >}})

ちなみに 2.3 系ではテスト用として UNIX プラットフォームでも使えるようになっているらしい。

{{< fig-quote type="markdown" title="GnuPG 2.3.3 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2021q4/000466.html" lang="en" >}}
{{% quote %}}Support a `gpgconf.ctl` file under Unix and use this for the regression tests{{% /quote %}}
{{< /fig-quote >}}


[GnuPG] 関連の各パッケージのバージョンは以下の通り（数字は大体の[ビルド]({{< ref "/openpgp/build-gnupg-in-ubuntu.md" >}} "Ubuntu で最新版 GnuPG をビルドする")順）。

|    # | パッケージ名                                             | バージョン   | 公開日     |         更新          |
| ---: | -------------------------------------------------------- | ------------ | ---------- | :-------------------: |
|    1 | [Libgpg-error](https://gnupg.org/software/libgpg-error/) | 1.43         | 2021-11-03 | {{< icons "check" >}} |
|    2 | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.8.8 (LTS)  | 2021-06-02 |                       |
|      | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.9.4        | 2021-08-22 |                       |
|    3 | [Libassuan](https://gnupg.org/software/libassuan/)       | 2.5.5        | 2021-03-22 |                       |
|    4 | [Libksba](https://gnupg.org/software/libksba/)           | 1.6.0        | 2021-06-10 |                       |
|    5 | [nPth](https://gnupg.org/software/npth/)                 | 1.6          | 2018-07-16 |                       |
|    6 | [ntbTLS](https://gnupg.org/software/ntbtls/)             | 0.2.0        | 2020-08-27 |                       |
|    7 | [GnuPG](https://gnupg.org/software/)                     | 2.2.33 (LTS) | 2021-11-23 | {{< icons "check" >}} |
|      | [GnuPG](https://gnupg.org/software/)                     | 2.3.3        | 2021-10-12 |                       |

[Libgpg-error](https://gnupg.org/software/libgpg-error/) が 1.43 に上がっているが，これについて

{{< fig-quote type="markdown" title="GnuPG 2.2.33 (LTS) released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2021q4/000467.html" lang="en" >}}
{{% quote %}}The provided Windows installer is now build against Libgpg-error 1.43 and thus features versatile meta options for the config file.  For example, it is possible to read and use option values from the Registry or from environment variables.  More on this is in a not published written blog article{{% /quote %}}.
{{< /fig-quote >}}

とある。
これからドキュメント化されるのかな？

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
