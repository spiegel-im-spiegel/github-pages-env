+++
title = "GnuPG 2.3.8 のリリース【セキュリティ・アップデート】"
date =  "2022-11-04T10:12:53+09:00"
description = "1件の脆弱性が修正されている。"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg", "vulnerability"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

毎度遅まきながらで申し訳ないが [GnuPG] 2.3.8 がリリースされている。

- [[Announce] [CVE-2022-3515] GnuPG / Libksba Security Advisory](https://lists.gnupg.org/pipermail/gnupg-announce/2022q4/000475.html)
- [[Announce] GnuPG 2.3.8 released](https://lists.gnupg.org/pipermail/gnupg-announce/2022q4/000476.html)

今回は1件の脆弱性修正を含んでいる。

## [CVE-2022-3515](https://nvd.nist.gov/vuln/detail/CVE-2022-3515)

{{< fig-quote type="markdown" title="[CVE-2022-3515] GnuPG / Libksba Security Advisory" link="https://lists.gnupg.org/pipermail/gnupg-announce/2022q4/000475.html" lang="en" >}}
A severe bug has been found in [Libksba] , the library used by GnuPG for parsing the ASN.1 structures as used by S/MIME.  The bug affects all versions of [Libksba] before 1.6.2 and may be used for remote code execution.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="[CVE-2022-3515] GnuPG / Libksba Security Advisory" link="https://lists.gnupg.org/pipermail/gnupg-announce/2022q4/000475.html" lang="en" >}}
The major user of Libksba is /gpgsm/, the S/MIME cousin of /gpg/. There it is used to parse all kind of input data, in particular signed or encrypted data in files or in mails.  Feeding a user with malicious data can thus be easily achieved.

A second user of Libksba is /dirmngr/, which is responsible for loading and parsing Certificate Revocation Lists (CRLs) and for verifying certificates used by TLS (i.e. https connections).  Mounting an attack is a bit more complex but can anyway be easily done using a rogue web server to serve a Web Key Directory, certificates, or CRLs.

An exploit is not yet publicly known but very straightforward to create for experienced crooks.
{{< /fig-quote >}}

今回の脆弱性について [GnuPG] 側は以下のように見積もっている。

- `CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:H/I:H/A:H`
- 深刻度: 重要 (Score: 8.1)

| 基本評価基準     | 評価値       |
| ---------------- | ------------ |
| 攻撃元区分       | ネットワーク |
| 攻撃条件の複雑さ | 高           |
| 必要な特権レベル | 不要         |
| ユーザ関与レベル | 不要         |
| スコープ         | 変更なし     |
| 機密性への影響   | 高           |
| 完全性への影響   | 高           |
| 可用性への影響   | 高           |

## [GnuPG] 関連パッケージ

[GnuPG] 関連の各パッケージのバージョンは以下の通り（数字は大体の[ビルド]({{< ref "/openpgp/build-gnupg-in-ubuntu.md" >}} "Ubuntu で最新版 GnuPG をビルドする")順）。

|    # | パッケージ名                                             | バージョン   | 公開日     |         更新          |
| ---: | -------------------------------------------------------- | ------------ | ---------- | :-------------------: |
|    1 | [Libgpg-error](https://gnupg.org/software/libgpg-error/) | 1.45         | 2022-10-07 | {{< icons "check" >}} |
|    2 | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.8.9 (LTS)  | 2022-02-07 |                       |
|      | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.10.1       | 2022-03-28 |                       |
|    3 | [Libassuan](https://gnupg.org/software/libassuan/)       | 2.5.5        | 2021-03-22 |                       |
|    4 | [Libksba](https://gnupg.org/software/libksba/)           | 1.6.2        | 2021-10-07 | {{< icons "check" >}} |
|    5 | [nPth](https://gnupg.org/software/npth/)                 | 1.6          | 2018-07-16 |                       |
|    6 | [ntbTLS](https://gnupg.org/software/ntbtls/)             | 0.3.1        | 2022-04-07 |                       |
|    7 | [GnuPG](https://gnupg.org/software/)                     | 2.2.40 (LTS) | 2022-10-10 | {{< icons "check" >}} |
|      | [GnuPG](https://gnupg.org/software/)                     | 2.3.8        | 2022-10-13 | {{< icons "check" >}} |

現在 [GnuPG] には2.2系と2.3系があり[^gpg14]，2.2系は LTS 版に位置付けられている。
2.3系では AEAD (Authenticated Encryption with Associated Data) 等 [RFC 4880bis] で検討されている機能が実装されているので，最新機能を試したいのであればこちらを入れるとよいだろう。
なお2.2系は少なくとも2024年末まではサポートが続けられる予定である。
通常運用であれば，当面は2.2系でも問題ない（[ECC も対応]({{< ref "/openpgp/using-ecc-with-gnupg.md" >}} "そろそろ GnuPG でも ECC を標準で使うのがいいんじゃないかな")してるよ）。

[^gpg14]: 厳密には1.4系もあるが，これは legacy 版と位置付けられており，よほどのバグか脆弱性がない限りは更新されない。もし今だに1.4系（あるいは既にサポートされていない2.0/2.1系）を使っているのなら2.2系以降にアップグレードすることを強くお勧めする。

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
