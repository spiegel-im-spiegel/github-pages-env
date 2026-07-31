+++
title = "GnuPG 2.2.31 (LTS) がリリースされた"
date =  "2021-09-16T20:39:59+09:00"
description = "Libgcrypt の脆弱性について追記。GnuPG 自身のセキュリティ・アップデートはなし。"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg", "vulnerability"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] LTS 版 2.2.31 がリリースされた。

- [[Announce] GnuPG 2.2.31 (LTS) released](https://lists.gnupg.org/pipermail/gnupg-announce/2021q3/000464.html)

セキュリティ・アップデートはなし。
詳細はこちら。

{{< fig-quote type="markdown" title="GnuPG 2.2.31 (LTS) released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2021q3/000464.html" lang="en" >}}
* agent: Fix a regression in `GET_PASSPHRASE`.  [#5577]
* scd: Fix an assertion failure in `close_pcsc_reader`.  [67e1834ad4]
* scd: Add support for PC/SC in "`GETINFO reader_list`".

Release-info: https://dev.gnupg.org/T5571
{{< /fig-quote >}}

## 【2021-09-17 追記】 Libgcrypt の脆弱性について

Ubuntu が [Libgcrypt の脆弱性報告](https://ubuntu.com/security/notices/USN-5080-1 "USN-5080-1: Libgcrypt vulnerabilities | Ubuntu security notices | Ubuntu")をしていて気が付いたが，[Libgcrypt] 1.8.8 および 1.9 系の 1.9.4 より前のバージョンは脆弱性を含んでいるらしい。
[GnuPG] 側のアナウンスがなくいつの間にかアップデートしてたので単純なバグ修正かと思ったらそうじゃなかったようだ。
そういうのはちゃんと通知してほしい。

てか， Ubuntu は（OpenSSL もそうだが）ちまちまバックポートパッチを当てるんじゃなくて普通にバージョンアップしてくれよ。
誰がそのパッチの安全性を保証するんだよ `orz`

### [CVE-2021-33560]

{{< fig-quote type="markdown" title="NVD - CVE-2021-33560" link="https://nvd.nist.gov/vuln/detail/CVE-2021-33560" lang="en" >}}
{{% quote %}}Libgcrypt before 1.8.8 and 1.9.x before 1.9.3 mishandles ElGamal encryption because it lacks exponent blinding to address a side-channel attack against mpi_powm, and the window size is not chosen appropriately. This, for example, affects use of ElGamal in OpenPGP{{% /quote %}}.
{{< /fig-quote >}}

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:N/A:N`
- 深刻度: 重要 (Score: 7.5)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更なし |
| 機密性への影響 | 高 |
| 完全性への影響 | なし |
| 可用性への影響 | なし |

### [CVE-2021-40528]

{{< fig-quote type="markdown" title="NVD - CVE-2021-40528" link="https://nvd.nist.gov/vuln/detail/CVE-2021-40528" lang="en" >}}
{{% quote %}}The ElGamal implementation in Libgcrypt before 1.9.4 allows plaintext recovery because, during interaction between two cryptographic libraries, a certain dangerous combination of the prime defined by the receiver's public key, the generator defined by the receiver's public key, and the sender's ephemeral exponents can lead to a cross-configuration attack against OpenPGP{{% /quote %}}.
{{< /fig-quote >}}

- `CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:H/I:N/A:N`
- 深刻度: 警告 (Score: 5.9)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 高 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更なし |
| 機密性への影響 | 高 |
| 完全性への影響 | なし |
| 可用性への影響 | なし |

- [On the (in)security of ElGamal in OpenPGP - Part I                  -                     Syssec@IBM Research](https://ibm.github.io/system-security-research-updates/2021/07/20/insecurity-elgamal-pt1)
- [On the (in)security of ElGamal in OpenPGP - Part II                  -                     Syssec@IBM Research](https://ibm.github.io/system-security-research-updates/2021/09/06/insecurity-elgamal-pt2)

[Libgcrypt]: https://gnupg.org/software/libgcrypt/
[CVE-2021-33560]: https://nvd.nist.gov/vuln/detail/CVE-2021-33560
[CVE-2021-40528]: https://nvd.nist.gov/vuln/detail/CVE-2021-40528

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
|    7 | [GnuPG](https://gnupg.org/software/)                     | 2.2.31 (LTS) | 2021-09-15 | {{< icons "check" >}} |
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
