+++
title = "GnuPG 2.3.7 のリリース【セキュリティ・アップデート】"
date =  "2022-08-14T10:10:01+09:00"
description = "1件の脆弱性が修正されている。"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg", "vulnerability"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

サボってました。
まじすんません。

毎度遅まきながらで申し訳ないが [GnuPG] 2.3.6 および 2.3.7 がリリースされている。

- [[Announce] GnuPG 2.3.6 released](https://lists.gnupg.org/pipermail/gnupg-announce/2022q2/000473.html)
- [[Announce] GnuPG 2.3.7 released](https://lists.gnupg.org/pipermail/gnupg-announce/2022q3/000474.html)

[GnuPG] 2.3.7 については1件の脆弱性が修正されている。

{{< fig-quote type="markdown" title="GnuPG 2.3.7 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2022q2/000472.html" lang="en" >}}
This release fixes CVE-2022-34903 which could be used to inject wrong status information in signatures.  The status information could then be abused to display a wrong validity in Kleopatra and other users of GPGME.
{{< /fig-quote >}}

## [CVE-2022-34903]

{{< fig-quote type="markdown" title="GnuPG 2.3.7 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2022q2/000472.html" lang="en" >}}
Fix possibly garbled status messages in `NOTATION_DATA`.  This bug could trick GPGME and other parsers to accept faked status lines.
{{< /fig-quote >}}

- `CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:H/I:L/A:N`
- 深刻度: 警告 (Score: 6.5)

| 基本評価基準     | 評価値       |
| ---------------- | ------------ |
| 攻撃元区分       | ネットワーク |
| 攻撃条件の複雑さ | 高           |
| 必要な特権レベル | 不要         |
| ユーザ関与レベル | 不要         |
| スコープ         | 変更なし     |
| 機密性への影響   | 高           |
| 完全性への影響   | 低           |
| 可用性への影響   | なし         |

## [GnuPG] 関連パッケージ

[GnuPG] 関連の各パッケージのバージョンは以下の通り（数字は大体の[ビルド]({{< ref "/openpgp/build-gnupg-in-ubuntu.md" >}} "Ubuntu で最新版 GnuPG をビルドする")順）。

|    # | パッケージ名                                             | バージョン   | 公開日     |         更新          |
| ---: | -------------------------------------------------------- | ------------ | ---------- | :-------------------: |
|    1 | [Libgpg-error](https://gnupg.org/software/libgpg-error/) | 1.45         | 2022-04-21 |                       |
|    2 | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.8.9 (LTS)  | 2022-02-07 |                       |
|      | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.10.1       | 2022-03-28 |                       |
|    3 | [Libassuan](https://gnupg.org/software/libassuan/)       | 2.5.5        | 2021-03-22 |                       |
|    4 | [Libksba](https://gnupg.org/software/libksba/)           | 1.6.0        | 2021-06-10 |                       |
|    5 | [nPth](https://gnupg.org/software/npth/)                 | 1.6          | 2018-07-16 |                       |
|    6 | [ntbTLS](https://gnupg.org/software/ntbtls/)             | 0.3.1        | 2022-04-07 |                       |
|    7 | [GnuPG](https://gnupg.org/software/)                     | 2.2.36 (LTS) | 2022-07-06 | {{< icons "check" >}} |
|      | [GnuPG](https://gnupg.org/software/)                     | 2.3.7        | 2022-07-11 | {{< icons "check" >}} |

うわっ！ 2.2系もバージョンが上がってるぢゃん。
アナウンスがないぞ。

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
[CVE-2022-34903]: https://nvd.nist.gov/vuln/detail/CVE-2022-34903

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
