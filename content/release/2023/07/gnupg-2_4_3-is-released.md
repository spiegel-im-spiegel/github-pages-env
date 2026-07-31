+++
title = "GnuPG 2.4.3 のリリース"
date =  "2023-07-05T19:42:04+09:00"
description = "脆弱性の修正はなし。"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] 2.4.3 がリリースされた。

- [[Announce] GnuPG 2.4.3 released](https://lists.gnupg.org/pipermail/gnupg-announce/2023q3/000480.html)
- [[Announce] GnuPG for OS X 2.4.3](https://lists.gnupg.org/pipermail/gnupg-users/2023-July/066614.html)

Mastodon でもアナウンスがあった。

{{< fig-quote  class="nobox center" >}}
<iframe src="https://mstdn.social/@GnuPG/110660553674699342/embed" class="mastodon-embed" style="max-width: 100%; border: 0" width="400" allowfullscreen="allowfullscreen"></iframe>
{{< /fig-quote >}}

脆弱性の修正はなし。
その他の改修ポイントは以下の通り。

{{< fig-quote type="markdown" title="GnuPG 2.4.3 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2023q3/000480.html" lang="en" >}}
* gpg: Set default expiration date to 3 years.  [T2701]
* gpg: Add `--list-filter` properties `"key_expires"` and `"key_expires_d"`.  [T6529]
* gpg: Emit status line and proper diagnostics for write errors.  [T6528]
* gpg: Make progress work for large files on Windows.  [T6534]
* gpg: New option `--no-compress` as alias for `-z0`.
* gpgsm: Print `PROGRESS` status lines.  Add new `--input-size-hint`.  [T6534]
* gpgsm: Support `SENDCERT_SKI` for `--call-dirmngr`.  [rG701a8b30f0]
* gpgsm: Major rewrite of the PKCS#12 parser.  [T6536]
* gpgtar: New option `--no-compress`.
* dirmngr: Extend the `AD_QUERY` command.  [rG207c99567c]
* dirmngr: Disable the HTTP redirect rewriting.  [T6477]
* dirmngr: New option `--compatibility-flags`.  [rGbf04b07327]
* dirmngr: New option `--ignore-crl-extensions`.  [T6545]
* wkd: Use export-clean for `gpg-wks-client`'s `--mirror` and `--create` commands.  [rG2c7f7a5a27]
* wkd: Make `--add-revocs` the default in `gpg-wks-client`.  New option `--no-add-revocs`.  [rG10c937ee68]
* scd: Make signing work for Nexus cards.  [rGb83d86b988]
* scd: Fix authentication with Administration Key for PIV.  [rG25b59cf6ce]

Release-info: https://dev.gnupg.org/T6509
{{< /fig-quote >}}

## [GnuPG] 関連パッケージ

[GnuPG] 関連の各パッケージのバージョンは以下の通り（数字は大体の[ビルド]({{< ref "/openpgp/build-gnupg-in-ubuntu.md" >}} "Ubuntu で最新版 GnuPG をビルドする")順）。

|    # | パッケージ名                                             | バージョン   | 公開日     |         更新          |
| ---: | -------------------------------------------------------- | ------------ | ---------- | :-------------------: |
|    1 | [Libgpg-error](https://gnupg.org/software/libgpg-error/) | 1.47         | 2023-04-06 |                       |
|    2 | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.8.10 (LTS) | 2023-01-05 |                       |
|      | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.10.2       | 2023-04-06 |                       |
|    3 | [Libassuan](https://gnupg.org/software/libassuan/)       | 2.5.6        | 2023-06-19 | {{< icons "check" >}} |
|    4 | [Libksba](https://gnupg.org/software/libksba/)           | 1.6.4        | 2023-06-19 | {{< icons "check" >}} |
|    5 | [nPth](https://gnupg.org/software/npth/)                 | 1.6          | 2018-07-16 |                       |
|    6 | [ntbTLS](https://gnupg.org/software/ntbtls/)             | 0.3.1        | 2022-04-07 |                       |
|    7 | [GnuPG](https://gnupg.org/software/)                     | 2.2.41 (LTS) | 2022-12-09 |                       |
|      | [GnuPG](https://gnupg.org/software/)                     | 2.4.3        | 2023-07-04 | {{< icons "check" >}} |

現在 [GnuPG] には 2.2 系と 2.4 系があり[^gpg14]， 2.4 系では AEAD (Authenticated Encryption with Associated Data) 等 [RFC 4880bis] で検討されている機能が実装されている。
2.2 系は 2.4 系のサブセットという位置づけで，少なくとも2024年末まではサポートが続けられる予定である。

[^gpg14]: 厳密には1.4系もあるが，これは legacy 版と位置付けられており，よほどのバグか脆弱性がない限りは更新されない。もし今だに1.4系（あるいは既にサポートされていない2.0/2.1系）を使っているのなら2.2系以降にアップグレードすることを強くお勧めする。

別記事でも書いたが，どうも [Ubuntu] は真面目に [GnuPG] のメンテナンスをやる気がないようである。
なので，近々自前でビルドを行おうかと考えているが，遅々として進まず...

- [Installing GnuPG 2.4 on Ubuntu 22.04 | Pro Custodibus](https://www.procustodibus.com/blog/2023/02/gpg-2-4-on-ubuntu-22-04/)

## ブックマーク

- [Gpg4win - Whats new - Version 4.2?](https://gpg4win.org/version4.2.html)
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
