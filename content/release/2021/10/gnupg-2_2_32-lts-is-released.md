+++
title = "GnuPG 2.2.32 (LTS) がリリースされた"
date =  "2021-10-07T21:01:12+09:00"
description = "今回は Let’s Encrypt 絡みの障害対応のようだ。"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] LTS 版 2.2.32 がリリースされた。

- [[Announce] GnuPG 2.2.32 (LTS) fixes a problem with Let's Encrypt](https://lists.gnupg.org/pipermail/gnupg-announce/2021q4/000465.html)

タイトルにある通り，今回は Let's Encrypt 絡みの障害対応のようだ。

{{< fig-quote type="markdown" title="⚓ T5639 dirmngr uses the wrong Let's encrypt chain" link="https://dev.gnupg.org/T5639" lang="en" >}}
{{% quote %}}dirmngr's certificate chain validation does not handle the new let's Encrypt root certificate correctly. When looking for the issuer of the intermediate certificate the first match ing certificate is used which might be the old second intermediate certificate then leading to the Root which expired on 2021-09-21. What we need to do is the same as what can be done with OpenSSL: Prefer trusted certificates ober the first found.. This way the old second intermediate certificate is not used but the new root{{% /quote %}}.
{{< /fig-quote >}}

その他，詳細はこちら。

{{< fig-quote type="markdown" title="GnuPG 2.2.32 (LTS) fixes a problem with Let's Encrypt" link="https://lists.gnupg.org/pipermail/gnupg-announce/2021q4/000465.html" lang="en" >}}
* dirmngr: Fix Let's Encrypt certificate chain validation.  [T5639]<br>(See https://dev.gnupg.org/T5639)
* dirmngr: New option `--ignore-cert`.  [323a20399d]
* gpg: Fix `--list-packets` for AEAD packets with unknown key.  [T5584]

Release-info: https://dev.gnupg.org/T5601
{{< /fig-quote >}}

## [GnuPG] 関連パッケージのバージョン

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
|    7 | [GnuPG](https://gnupg.org/software/)                     | 2.2.32 (LTS) | 2021-10-06 | {{< icons "check" >}} |
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
