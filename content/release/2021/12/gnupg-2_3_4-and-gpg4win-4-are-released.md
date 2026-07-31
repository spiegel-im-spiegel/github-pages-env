+++
title = "GnuPG 2.3.4 および Gpg4win 4 のリリース"
date =  "2021-12-25T13:57:55+09:00"
description = "セキュリティ・アップデートはなし。"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg", "windows"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

毎度遅まきながらで申し訳ないが [GnuPG] 2.3.4 がリリースされている。

- [[Announce] GnuPG 2.3.4 released](https://lists.gnupg.org/pipermail/gnupg-announce/2021q4/000468.html)

セキュリティ・アップデートはなし。
詳細はこちら。

{{< fig-quote type="markdown" title="GnuPG 2.3.4 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2021q4/000468.html" lang="en" >}}
* gpg: New option `--min-rsa-length`.  [rG5f39db70c0]
* gpg: New option `--forbid-gen-key`.  [rGc397ba3ac0]
* gpg: New option `--override-compliance-check`.  [T5655]
* gpgconf: New command `--show-configs`.  [rGa0fb78ee0f]
* agent,dirmngr,keyboxd: New option `--steal-socket`. [rGb0079ab39d,rGdd708f60d5]
* gpg: Fix printing of binary notations.  [T5667]
* gpg: Remove stale ultimately trusted keys from the trustdb. [T5685,T5742]
* gpg: Fix indentation of `--print-mds` and `--print-md sha512`.  [T5679]
* gpg: Emit gpg 2.2 compatible Ed25519 signature.  [T5331]
* gpgsm: Detect circular chains in `--list-chain`.  [rG74c5b35062]
* dirmngr: Make reading resolv.conf more robust.  [T5657]
* dirmngr: Ask keyservers to provide the key fingerprints.  [T5741]
* gpgconf: Allow changing gpg's deprecated keyserver option.  [T5462]
* gpg-wks-server: Fix created file permissions.  [rG60be00b033]
* scd: Support longer data for ssh-agent authentication with openpgp cards.  [T5682]
* scd: Modify `DEVINFO` behavior to support looping forever.  [T5359]
* Support gpgconf.ctl for NetBSD and Solaris.  [T5656,T5671]
* Silence "Garbled console data" warning under Windows in most cases.  [rGe293da3b21]
* Silence warning about the rootdir under Unices w/o a mounted /proc file system.  [T5656]
* Fix possible build problems about missing include files.  [T5592]

Release-info: https://dev.gnupg.org/T5654
{{< /fig-quote >}}

[GnuPG] 関連の各パッケージのバージョンは以下の通り（数字は大体の[ビルド]({{< ref "/openpgp/build-gnupg-in-ubuntu.md" >}} "Ubuntu で最新版 GnuPG をビルドする")順）。

|    # | パッケージ名                                             | バージョン   | 公開日     |         更新          |
| ---: | -------------------------------------------------------- | ------------ | ---------- | :-------------------: |
|    1 | [Libgpg-error](https://gnupg.org/software/libgpg-error/) | 1.43         | 2021-11-03 |                       |
|    2 | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.8.8 (LTS)  | 2021-06-02 |                       |
|      | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.9.4        | 2021-08-22 |                       |
|    3 | [Libassuan](https://gnupg.org/software/libassuan/)       | 2.5.5        | 2021-03-22 |                       |
|    4 | [Libksba](https://gnupg.org/software/libksba/)           | 1.6.0        | 2021-06-10 |                       |
|    5 | [nPth](https://gnupg.org/software/npth/)                 | 1.6          | 2018-07-16 |                       |
|    6 | [ntbTLS](https://gnupg.org/software/ntbtls/)             | 0.2.0        | 2020-08-27 |                       |
|    7 | [GnuPG](https://gnupg.org/software/)                     | 2.2.33 (LTS) | 2021-11-23 |                       |
|      | [GnuPG](https://gnupg.org/software/)                     | 2.3.4        | 2021-12-20 | {{< icons "check" >}} |

現在 [GnuPG] には2.2系と2.3系があり[^gpg14]，2.2系は LTS 版に位置付けられている。
2.3系では AEAD (Authenticated Encryption with Associated Data) 等 [RFC 4880bis] で検討されている機能が実装されているので，最新機能を試したいのであればこちらを入れるとよいだろう。
なお2.2系は少なくとも2024年末まではサポートが続けられる予定である。
通常運用であれば，当面は2.2系でも問題ない（[ECC も対応]({{< ref "/openpgp/using-ecc-with-gnupg.md" >}} "そろそろ GnuPG でも ECC を標準で使うのがいいんじゃないかな")してるよ）。

[^gpg14]: 厳密には1.4系もあるが，これは legacy 版と位置付けられており，よほどのバグか脆弱性がない限りは更新されない。もし今だに1.4系（あるいは既にサポートされていない2.0/2.1系）を使っているのなら2.2系以降にアップグレードすることを強くお勧めする。

さて [GnuPG] 2.3.4 のリリースとほぼ同じタイミングで [Gpg4win] 4 系最初のバージョン 4.0.0 がリリースされた。
[Gpg4win] 4 系では [GnuPG] 2.3 系が導入された。

- [[Gpg4win-users-en] [Gpg4win-announce] Gpg4win 4.0.0 released](https://lists.wald.intevation.org/pipermail/gpg4win-users-en/2021-December/001704.html)
- [Gpg4win - Whats new - Version 4?](https://gpg4win.org/version4.html)

これまでの [Gpg4win] 3.1 系も引き続きメンテナンスされているようなので，慌てずにタイミングを見計らってアップグレードしていくのがいいだろう。

アップデートは計画的に。

## ブックマーク

- [OpenPGP の実装]({{< rlnk "/openpgp/" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Gpg4win]: https://gpg4win.org/ "Gpg4win - Secure email and file encryption with GnuPG for Windows"
[OpenPGP]: http://openpgp.org/
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
