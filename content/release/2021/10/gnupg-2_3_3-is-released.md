+++
title = "GnuPG 2.3.3 がリリースされた"
date =  "2021-10-16T13:45:19+09:00"
description = "Windows 版で keyboxd の提供が始まったようだ。"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

遅まきながらで申し訳ないが [GnuPG] 2.3.3 がリリースされている。

- [[Announce] GnuPG 2.3.3 released](https://lists.gnupg.org/pipermail/gnupg-announce/2021q4/000466.html)

セキュリティ・アップデートはなし。
詳細はこちら。
かなりあるぞ。

{{< fig-quote type="markdown" title="GnuPG 2.3.3 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2021q4/000466.html" lang="en" >}}
* agent: Fix segv in `GET_PASSPHRASE` (regression).  [#5577]
* dirmngr: Fix Let's Encrypt certificate chain validation.  [#5639]
* gpg: Change default and maximum AEAD chunk size to 4 MiB. [ad3dabc9fb]
* gpg: Print a warning when importing a bad cv25519 secret key. [#5464]
* gpg: Fix `--list-packets` for undecryptable AEAD packets.  [#5584]
* gpg: Verify backsigs for v5 keys correctly.  [#5628]
* keyboxd: Fix checksum computation for no UBID entry on disk. [#5573]
* keyboxd: Fix "invalid object" error with cv448 keys.  [#5609]
* dirmngr: New option `--ignore-cert`.  [4b3e9a44b5]
* agent: Fix `calibrate_get_time` use of `clock_gettime`.  [#5623]
* Silence process spawning diagnostics on Windows. [f2b01025c3]
* Support a `gpgconf.ctl` file under Unix and use this for the regression tests.  [#5999]
* The Windows installer now also installs the new keyboxd. (Put "`use-keyboxd`" into `common.conf` to use a fast SQLite database instead of the `pubring.kbx` file.)

Release-info: https://dev.gnupg.org/T5565
{{< /fig-quote >}}

keyboxd ってのは [GnuPG 2.3.1]({{< ref "/release/2021/04/gnupg-2_3_1-is-released.md" >}} "GnuPG 2.3.1 がリリースされた") で導入された鍵管理用の daemon らしい。
あんまり安定してるように見えないけど， Windows 版も対応し始めたみたいだし，いずれはこれに移行するんだろうな。

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
|    7 | [GnuPG](https://gnupg.org/software/)                     | 2.2.32 (LTS) | 2021-10-06 |                       |
|      | [GnuPG](https://gnupg.org/software/)                     | 2.3.3        | 2021-10-12 | {{< icons "check" >}} |

現在 [GnuPG] には2.2系と2.3系があり[^gpg14]，2.2系は LTS 版に位置付けられている。
2.3系では AEAD (Authenticated Encryption with Associated Data) 等 [RFC 4880bis] で検討されている機能が実装されているので，最新機能を試したいのであればこちらを入れるとよいだろう。
なお2.2系は少なくとも2024年末まではサポートが続けられる予定である。
通常運用であれば，当面は2.2系でも問題ない（[ECC も対応]({{< ref "/openpgp/using-ecc-with-gnupg.md" >}} "そろそろ GnuPG でも ECC を標準で使うのがいいんじゃないかな")してるよ）。

[^gpg14]: 厳密には1.4系もあるが，これは legacy 版と位置付けられており，よほどのバグか脆弱性がない限りは更新されない。もし今だに1.4系（あるいは既にサポートされていない2.0/2.1系）を使っているのなら2.2系以降にアップグレードすることを強くお勧めする。

アップデートは計画的に。

## ブックマーク

- [Using a TPM with GnuPG 2.3](https://gnupg.org/blog/20210315-using-tpm-with-gnupg-2.3.html)
- [OpenPGP の実装]({{< rlnk "/openpgp/" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenPGP]: http://openpgp.org/

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
