+++
title = "GnuPG 2.4.4 / 2.4.5 のリリース【セキュリティ・アップデートを含む】"
date =  "2024-03-20T11:03:46+09:00"
description = "2.4.4 では脆弱性の修正がある"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg", "security", "vulnerability"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

仕事の忙しさにかまけて色々と放っぽり出してたので，少しずつ回復中。

遅まきながらであるが， 2024-01-25 に [GnuPG] 2.4.4 がリリースされた。
さらに先日 2024-03-12 に 2.4.5 がリリースされている。

- [[Announce] GnuPG 2.4.4 released](https://lists.gnupg.org/pipermail/gnupg-announce/2024q1/000481.html)
- [[Announce] GnuPG 2.4.5 released](https://lists.gnupg.org/pipermail/gnupg-announce/2024q1/000482.html)

詳細については割愛する。
[GnuPG] 2.4.4 については次節で紹介する脆弱性の修正がある。

## 【脆弱性】スマートカード生成時に保護されていないバックアップ鍵が PC 上に残る

[GnuPG] 2.4.4 より前の一部のバージョンでスマートカード生成時に保護されていないバックアップ鍵が PC 上に残る脆弱性があるらしい。

{{< fig-quote type="markdown" title="Smartcard key generation keeps an unprotected backup key on disk" link="https://gnupg.org/blog/20240125-smartcard-backup-key.html" lang="en" >}}
The standard way to generate keys on a smartcard with GnuPG is to create the encryption subkey with gpg and to move this key to the smartcard. A password protected backup file named `sk_<keyid>.gpg` is also created so that in the case of a lost or broken smartcard, the key can be restored to a new smartcard to allow decryption of existing data. Unfortunately with some versions of GnuPG an additional unprotected copy of the encryption subkey is also kept on disk.

All possibly affected users should check whether such an unintended copy of a smartcard key exists and delete it.
{{< /fig-quote >}}

該当するバージョンは以下の通り。

{{< fig-quote type="markdown" title="Smartcard key generation keeps an unprotected backup key on disk" link="https://gnupg.org/blog/20240125-smartcard-backup-key.html" lang="en" >}}
- Gpg4win version 4.2.0
- GnuPG versions 2.4.2 and 2.4.3 iff the card generation was done with the command `gpg --card-edit`.
- GnuPG version 2.2.42 iff the card generation was done with the command `gpg --card-edit`.
- GnuPG VS-Desktop version 3.2.0 and 3.2.1 iff the card generation was done with the non-approved command `gpg --card-edit`. The documented way to create keys on OpenPGP cards and Yubikeys is not affected.
{{< /fig-quote >}}

確認方法等については以下を参照のこと。

- [Smartcard key generation keeps an unprotected backup key on disk](https://gnupg.org/blog/20240125-smartcard-backup-key.html)

## [GnuPG] 関連パッケージ

[GnuPG] 関連の各パッケージのバージョンは以下の通り（数字は大体の[ビルド]({{< ref "/openpgp/build-gnupg-in-ubuntu.md" >}} "Ubuntu で最新版 GnuPG をビルドする")順）。

|    # | パッケージ名                                             | バージョン   | 公開日     |         更新          |
| ---: | -------------------------------------------------------- | ------------ | ---------- | :-------------------: |
|    1 | [Libgpg-error](https://gnupg.org/software/libgpg-error/) | 1.48         | 2024-02-23 | {{< icons "check" >}} |
|    2 | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.8.11 (LTS) | 2023-11-16 | {{< icons "check" >}} |
|      | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.10.3       | 2023-11-14 | {{< icons "check" >}} |
|    3 | [Libassuan](https://gnupg.org/software/libassuan/)       | 2.5.7        | 2024-03-06 | {{< icons "check" >}} |
|    4 | [Libksba](https://gnupg.org/software/libksba/)           | 1.6.6        | 2024-02-23 | {{< icons "check" >}} |
|    5 | [nPth](https://gnupg.org/software/npth/)                 | 1.7          | 2024-02-23 | {{< icons "check" >}} |
|    6 | [ntbTLS](https://gnupg.org/software/ntbtls/)             | 0.3.2        | 2024-01-12 | {{< icons "check" >}} |
|    7 | [GnuPG](https://gnupg.org/software/)                     | 2.4.5        | 2024-03-07 | {{< icons "check" >}} |

ありゃ。
2.2 系のコードが一覧から消えてるな。

現在 [GnuPG] には 2.2 系と 2.4 系があり[^gpg14]， 2.4 系では AEAD (Authenticated Encryption with Associated Data) 等 [RFC 4880bis] で検討されている機能が実装されている。
2.2 系は 2.4 系のサブセットという位置づけで，少なくとも2024年末まではサポートが続けられる予定である。

[^gpg14]: 厳密には1.4系もあるが，これは legacy 版と位置付けられており，よほどのバグか脆弱性がない限りは更新されない。もし今だに1.4系（あるいは既にサポートされていない2.0/2.1系）を使っているのなら2.4系以降にアップグレードすることを強くお勧めする。

別記事でも書いたが，どうも [Ubuntu] は真面目に [GnuPG] のメンテナンスをやる気がないようである。
なので，自前でビルドを行おうかと考えているが，遅々として進まず...

- [Installing GnuPG 2.4 on Ubuntu 22.04 | Pro Custodibus](https://www.procustodibus.com/blog/2023/02/gpg-2-4-on-ubuntu-22-04/)

## ブックマーク

- [iquiw/pinentry-w32-ncg-binary: pinentry-w32, no characters garbled](https://github.com/iquiw/pinentry-w32-ncg-binary)
- [ADSK: The Additional Decryption Subkey](https://gnupg.org/blog/20230321-adsk.html)

- [OpenPGP の実装]({{< rlnk "/openpgp/" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Gpg4win]: https://gpg4win.org/ "Gpg4win - Secure email and file encryption with GnuPG for Windows"
[OpenPGP]: http://openpgp.org/
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
{{% review-paapi "B0CW1DZS1W" %}} <!-- 岩波「科学」2024年3月号 現代暗号の展開と応用 -->
