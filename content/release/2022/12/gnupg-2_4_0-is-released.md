+++
title = "GnuPG 2.4.0 のリリース【セキュリティ・アップデートを含む】"
date =  "2022-12-25T17:26:41+09:00"
description = "Libksba の脆弱性修正を含んでいる。そして，25周年おめでとう！"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg", "vulnerability"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[入院中]({{< ref "/remark/2022/12/heart-attack.md" >}} "ハライタだと思った？ 残念！ 心筋梗塞でした")だったためリアクションが遅れたけど [GnuPG] 2.4.0 がリリースされた。
そして，25周年おめでとう！

- [[Announce] GnuPG 2.4.0 released (silver anniversary)](https://lists.gnupg.org/pipermail/gnupg-announce/2022q4/000477.html)

そういえば昨年は[030周年]({{< ref "/remark/2021/12/gnupg-030-years-anniversary.md" >}} "GnuPG 030周年（笑）と寄付の話")だったな（笑）

2.4 系がリリースされたことで，今後は 2.4 系が最新 stable 版としてメンテナンスされていく。
LTS 版である 2.2 系は 2.4 系の不具合や脆弱性の修正のみバックポートされることになるだろう。

{{< fig-quote type="markdown" title="GnuPG 2.4.0 released (silver anniversary)" link="https://lists.gnupg.org/pipermail/gnupg-announce/2022q4/000477.html" lang="en" >}}
Version 2.2 is our LTS (long term support) version and guaranteed to be maintained at least until the end of 2024. Only a small subset of features from 2.4 has been back-ported to this series.
{{< /fig-quote >}}

そして，このバージョンから

{{< fig-quote type="markdown" title="GnuPG 2.4.0 released (silver anniversary)" link="https://lists.gnupg.org/pipermail/gnupg-announce/2022q4/000477.html" lang="en" >}}
The key database daemon is now a fully supported feature.  Keys are stored in a SQLite database to make key lookups much faster.  Enable it by adding `"use-keyboxd"` o `common.conf`.  See also the `README` file.
{{< /fig-quote >}}

となるそうだ。
うーん，試す環境がない。
Ubuntu は 2.2 系のまま突き進むのかなぁ。

## [CVE-2022-3515](https://nvd.nist.gov/vuln/detail/CVE-2022-3515) GnuPG / Libksba Security Advisory

今回のリリースには [Libksba](https://gnupg.org/software/libksba/) のセキュリティ・アップデートを含んでいる。

- [[Announce] [CVE-2022-3515] GnuPG / Libksba Security Advisory](https://lists.gnupg.org/pipermail/gnupg-announce/2022q4/000475.html)

これによると

{{< fig-quote type="markdown" title="[CVE-2022-3515] GnuPG / Libksba Security Advisory" link="https://lists.gnupg.org/pipermail/gnupg-announce/2022q4/000475.html" lang="en" >}}
A severe bug has been found in [Libksba] , the library used by GnuPG for parsing the ASN.1 structures as used by S/MIME.  The bug affects all versions of [Libksba] before 1.6.2 and may be used for remote code execution.
{{< /fig-quote >}}

だそうだ。
結構ヤバい。
“*Updating this library is thus important*.” とあるので早めのアップデートを。

[Libksba](https://gnupg.org/software/libksba/) のバージョンを確認するには

```text
$ gpgconf --show-versions | grep KSBA
```

とすればいいらしい。
これで

```text
* KSBA 1.6.3 (xxxxx)
```

てな風になっていれば OK かな。

（以下未稿）

## [GnuPG] 関連パッケージ

[GnuPG] 関連の各パッケージのバージョンは以下の通り（数字は大体の[ビルド]({{< ref "/openpgp/build-gnupg-in-ubuntu.md" >}} "Ubuntu で最新版 GnuPG をビルドする")順）。

|    # | パッケージ名                                             | バージョン   | 公開日     |         更新          |
| ---: | -------------------------------------------------------- | ------------ | ---------- | :-------------------: |
|    1 | [Libgpg-error](https://gnupg.org/software/libgpg-error/) | 1.46         | 2022-10-07 | {{< icons "check" >}} |
|    2 | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.8.9 (LTS)  | 2022-02-07 |                       |
|      | [Libgcrypt](https://gnupg.org/software/libgcrypt/)       | 1.10.1       | 2022-03-28 |                       |
|    3 | [Libassuan](https://gnupg.org/software/libassuan/)       | 2.5.5        | 2021-03-22 |                       |
|    4 | [Libksba](https://gnupg.org/software/libksba/)           | 1.6.3        | 2022-12-06 | {{< icons "check" >}} |
|    5 | [nPth](https://gnupg.org/software/npth/)                 | 1.6          | 2018-07-16 |                       |
|    6 | [ntbTLS](https://gnupg.org/software/ntbtls/)             | 0.3.1        | 2022-04-07 |                       |
|    7 | [GnuPG](https://gnupg.org/software/)                     | 2.2.41 (LTS) | 2022-12-09 | {{< icons "check" >}} |
|      | [GnuPG](https://gnupg.org/software/)                     | 2.4.0        | 2022-12-16 | {{< icons "check" >}} |

現在 [GnuPG] には 2.2 系と 2.4 系があり[^gpg14]， 2.4 系では AEAD (Authenticated Encryption with Associated Data) 等 [RFC 4880bis] で検討されている機能が実装されている。
2.2 系は 2.4 系のサブセットという位置づけで，少なくとも2024年末まではサポートが続けられる予定である。

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
