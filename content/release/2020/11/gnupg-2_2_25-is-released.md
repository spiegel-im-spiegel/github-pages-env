+++
title = "GnuPG 2.2.25 がリリースされた"
date =  "2020-11-24T08:24:09+09:00"
description = "どうやら 2.2.24 に regression があったようだ。"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "tools", "openpgp", "gnupg"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

このまえ 2.2.24 が出たばかりだが [GnuPG] 2.2.25 がリリースされた。

- [[Announce] GnuPG 2.2.25 released](https://lists.gnupg.org/pipermail/gnupg-announce/2020q4/000450.html)

どうやら前回の 2.2.24 に regression があったようだ。

詳細はこちら。

{{< fig-quote type="markdown" title="GnuPG 2.2.24 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2020q4/000449.html" lang="en" >}}
* scd: Fix regression in 2.2.24 requiring `gpg --card-status` before signing or decrypting.  [#5065]
* gpgsm: Using Libksba 1.5.0 signatures with a rarely used combination of attributes can now be verified.  [#5146]

Release-info: https://dev.gnupg.org/T5140
{{< /fig-quote >}}

アップデートは計画的に。

## ブックマーク

- [OpenPGP の実装]({{< rlnk "/openpgp/" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenPGP]: http://openpgp.org/

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
