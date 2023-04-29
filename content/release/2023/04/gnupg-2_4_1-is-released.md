+++
title = "GnuPG 2.4.1 のリリース"
date =  "2023-04-29T18:29:31+09:00"
description = "脆弱性の修正はなし。機能の追加と不具合の修正"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] 2.4.1 がリリースされた。

- [[Announce] GnuPG 2.4.1 released](https://lists.gnupg.org/pipermail/gnupg-announce/2023q2/000478.html)

脆弱性の修正はなし。
その他の改修ポイントは以下の通り。

{{< fig-quote type="markdown" title="GnuPG 2.4.1 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2023q2/000478.html" lang="en" >}}
* If the ~/.gnupg directory does not exist, the keyboxd is now automagically enabled.  [rGd9e7488b17]
* gpg: New option `--add-desig-revoker`.  [rG3d094e2bcf]
* gpg: New option `--assert-signer`.  [rGc9e95b8dee]
* gpg: New command `--quick-add-adsk` and other ADSK features.  [T6395, https://gnupg.org/blog/20230321-adsk.html]
* gpg: New list-option "`show-unusable-sigs`".  Also show "`[self-signature]`" instead of the `user-id` in key signature listings.  [rG103acfe9ca]
* gpg: For symmetric encryption the default S2K hash is now SHA256.  [T6367]
* gpg: Detect already compressed data also when using a pipe.  Also detect JPEG and PNG file formats.  [T6332]
* gpg: New subcommand "`openpgp`" for `--card-edit`.  [T6462]
* gpgsm: Verification of detached signatures does now strip trailing zeroes from the input if `--assume-binary` is used.  [rG2a13f7f9dc]
* gpgsm: Non-armored detached signature are now created without using indefinite form length octets.  This improves compatibility with some PDF signature verification software.  [rG8996b0b655]
* gpgtar: Emit progress status lines in create mode.  [T6363]
* dirmngr: The LDAP modifyTimestamp is now returned by some keyserver commands.  [rG56d309133f]
* ssh: Allow specification of the order keys are presented to ssh. See the man page entry for `--enable-ssh-support`.  [T5996, T6212]
* gpg: Make list-options "`show-sig-subpackets`" work again. Fixes regression in 2.4.0.  [rG5a223303d7]
* gpg: Fix the keytocard command for Yubikeys.  [T6378]
* gpg: Do not continue an export after a cancel for the primary key.  [T6093]
* gpg: Replace the `--override-compliance-check` hack by a real fix.  [T5655]
* gpgtar: Fix decryption with input taken from stdin.  [T6355]

Release-info: https://dev.gnupg.org/T6454
{{< /fig-quote >}}

とりあえず `--quick-add-adsk` オプションと `--enable-ssh-support` オプションが気になるところではある。
`--quick-add-adsk` オプションおよび ADSK については以下の記事が参考になる。

- [ADSK: The Additional Decryption Subkey](https://gnupg.org/blog/20230321-adsk.html)

## [GnuPG] 関連パッケージ

[GnuPG] 関連の各パッケージのバージョンは以下の通り（数字は大体の[ビルド]({{< ref "/openpgp/build-gnupg-in-ubuntu.md" >}} "Ubuntu で最新版 GnuPG をビルドする")順）。

|    # | パッケージ名                                             | バージョン   | 公開日     |         更新          |
| ---: | -------------------------------------------------------- | ------------ | ---------- | :-------------------: |
|    1 | [Libgpg-error](https://gnupg.org/software/libgpg-error/) | 1.47         | 2023-04-06 | {{< icons "check" >}} |
|    2 | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.8.10 (LTS) | 2023-01-05 | {{< icons "check" >}} |
|      | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.10.2       | 2023-04-06 | {{< icons "check" >}} |
|    3 | [Libassuan](https://gnupg.org/software/libassuan/)       | 2.5.5        | 2021-03-22 |                       |
|    4 | [Libksba](https://gnupg.org/software/libksba/)           | 1.6.3        | 2022-12-06 |                       |
|    5 | [nPth](https://gnupg.org/software/npth/)                 | 1.6          | 2018-07-16 |                       |
|    6 | [ntbTLS](https://gnupg.org/software/ntbtls/)             | 0.3.1        | 2022-04-07 |                       |
|    7 | [GnuPG](https://gnupg.org/software/)                     | 2.2.41 (LTS) | 2022-12-09 |                       |
|      | [GnuPG](https://gnupg.org/software/)                     | 2.4.1        | 2023-04-28 | {{< icons "check" >}} |

現在 [GnuPG] には 2.2 系と 2.4 系があり[^gpg14]， 2.4 系では AEAD (Authenticated Encryption with Associated Data) 等 [RFC 4880bis] で検討されている機能が実装されている。
2.2 系は 2.4 系のサブセットという位置づけで，少なくとも2024年末まではサポートが続けられる予定である。

[^gpg14]: 厳密には1.4系もあるが，これは legacy 版と位置付けられており，よほどのバグか脆弱性がない限りは更新されない。もし今だに1.4系（あるいは既にサポートされていない2.0/2.1系）を使っているのなら2.2系以降にアップグレードすることを強くお勧めする。

別記事でも書いたが，どうも [Ubuntu] は真面目に [GnuPG] のメンテナンスをやる気がないようである。
なので，せっかく GW で時間に余裕があるし，近々自前でビルドを行おうかと考えている。

- [Installing GnuPG 2.4 on Ubuntu 22.04 | Pro Custodibus](https://www.procustodibus.com/blog/2023/02/gpg-2-4-on-ubuntu-22-04/)

## ブックマーク

- [iquiw/pinentry-w32-ncg-binary: pinentry-w32, no characters garbled](https://github.com/iquiw/pinentry-w32-ncg-binary)

- [OpenPGP の実装]({{< rlnk "/openpgp/" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Gpg4win]: https://gpg4win.org/ "Gpg4win - Secure email and file encryption with GnuPG for Windows"
[OpenPGP]: http://openpgp.org/
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
