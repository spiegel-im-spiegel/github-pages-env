+++
title = "GnuPG 2.4.2 のリリースと Mastodon 公式アカウントができた話"
date =  "2023-06-11T10:28:34+09:00"
description = "脆弱性の修正はなし。"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

毎度遅まきながらで申し訳ないが [GnuPG] 2.4.2 がリリースされた。

- [[Announce] GnuPG 2.4.2 released](https://lists.gnupg.org/pipermail/gnupg-announce/2023q2/000479.html)

脆弱性の修正はなし。
その他の改修ポイントは以下の通り。

{{< fig-quote type="markdown" title="GnuPG 2.4.2 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2023q2/000479.html" lang="en" >}}
* gpg: Print a warning if no more encryption subkeys are left over after changing the expiration date.  [rGef2c3d50fa]
* gpg: Fix searching for the ADSK key when adding an ADSK.  [T6504]
* gpgsm: Speed up key listings on Windows.  [rG08ff55bd44]
* gpgsm: Reduce the number of "failed to open policy file" diagnostics.  [rG68613a6a9d]
* agent: Make updating of private key files more robust and track display S/N.  [T6135]
* keyboxd: Avoid longish delays on Windows when listing keys.  [rG6944aefa3c]
* gpgtar: Emit extra status lines to help GPGME.  [T6497]
* w32: Avoid using the VirtualStore.  [T6403]
{{< /fig-quote >}}

## [GnuPG] 関連パッケージ

[GnuPG] 関連の各パッケージのバージョンは以下の通り（数字は大体の[ビルド]({{< ref "/openpgp/build-gnupg-in-ubuntu.md" >}} "Ubuntu で最新版 GnuPG をビルドする")順）。

|    # | パッケージ名                                             | バージョン   | 公開日     |         更新          |
| ---: | -------------------------------------------------------- | ------------ | ---------- | :-------------------: |
|    1 | [Libgpg-error](https://gnupg.org/software/libgpg-error/) | 1.47         | 2023-04-06 |                       |
|    2 | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.8.10 (LTS) | 2023-01-05 |                       |
|      | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.10.2       | 2023-04-06 |                       |
|    3 | [Libassuan](https://gnupg.org/software/libassuan/)       | 2.5.5        | 2021-03-22 |                       |
|    4 | [Libksba](https://gnupg.org/software/libksba/)           | 1.6.3        | 2022-12-06 |                       |
|    5 | [nPth](https://gnupg.org/software/npth/)                 | 1.6          | 2018-07-16 |                       |
|    6 | [ntbTLS](https://gnupg.org/software/ntbtls/)             | 0.3.1        | 2022-04-07 |                       |
|    7 | [GnuPG](https://gnupg.org/software/)                     | 2.2.41 (LTS) | 2022-12-09 |                       |
|      | [GnuPG](https://gnupg.org/software/)                     | 2.4.2        | 2023-05-30 | {{< icons "check" >}} |

現在 [GnuPG] には 2.2 系と 2.4 系があり[^gpg14]， 2.4 系では AEAD (Authenticated Encryption with Associated Data) 等 [RFC 4880bis] で検討されている機能が実装されている。
2.2 系は 2.4 系のサブセットという位置づけで，少なくとも2024年末まではサポートが続けられる予定である。

[^gpg14]: 厳密には1.4系もあるが，これは legacy 版と位置付けられており，よほどのバグか脆弱性がない限りは更新されない。もし今だに1.4系（あるいは既にサポートされていない2.0/2.1系）を使っているのなら2.2系以降にアップグレードすることを強くお勧めする。

別記事でも書いたが，どうも [Ubuntu] は真面目に [GnuPG] のメンテナンスをやる気がないようである。
なので，近々自前でビルドを行おうかと考えているが，遅々として進まず...

- [Installing GnuPG 2.4 on Ubuntu 22.04 | Pro Custodibus](https://www.procustodibus.com/blog/2023/02/gpg-2-4-on-ubuntu-22-04/)

## 【余談】 GnuPG の Mastodon 公式アカウントができたらしい

この前 Twitter で見かけたのだが

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="en" dir="ltr">Now that we lost our verification mark, we will likely move to a Mastodon: @GnuPG@mstdn.social</p>&mdash; GNU Privacy Guard (@gnupg) <a href="https://twitter.com/gnupg/status/1664251891230941184?ref_src=twsrc%5Etfw">June 1, 2023</a></blockquote>
{{< /fig-gen >}}

なんだそうで，本当に Mastodon にアカウントができていた。

- [GnuPG (@GnuPG@mstdn.social) - Mastodon 🐘](https://mstdn.social/@GnuPG)

Twitter アカウントはマークを外された状態で継続するのかな？
まぁ認証マークつってもアレって「Twitter 教にお布施しました」マークだろ，金で買えるらしいし（笑） それよりも早く [GnuPG.org][GnuPG] サイトに `<link rel="me">` 要素で連携して認証してもらいなはれ！

こうやって Twitter を参照する理由がひとつずつ減っていくんだねぇ。

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
