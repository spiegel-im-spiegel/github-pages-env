+++
title = "GnuPG 2.2.23 のリリース【セキュリティ・アップデート】"
date =  "2020-09-04T12:45:27+09:00"
description = "知らない人の鍵はインポートしないように（笑）"
image = "/images/attention/openpgp.png"
tags = [
  "security",
  "vulnerability",
  "cryptography",
  "tools",
  "openpgp",
  "gnupg",
]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] 2.2.23 がリリースされた。

- [[Announce] [security fix] GnuPG 2.2.23 released](https://lists.gnupg.org/pipermail/gnupg-announce/2020q3/000448.html)

以下の脆弱性の改修を含む。

{{< fig-quote type="markdown" title="GnuPG 2.2.23 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2020q3/000448.html" lang="en" >}}
{{% quote %}}Importing an OpenPGP key having a preference list for AEAD algorithms will lead to an array overflow and thus often to a crash or other undefined behaviour{{% /quote %}}.
{{< /fig-quote >}}

CVE ID は（まだ？）割り振られていないようだ。
ちなみに 2.2.20 以下のバージョンには影響しない。

AEAD (Authenticated Encryption with Associated Data; 認証付き暗号) は次期 OpenPGP ([RFC 4880bis]) で導入予定の暗号モードで，先行して [GnuPG] に組み込まれている。
インパクトも可用性リスクのみと思われる。

なので，殆どの人には影響はないと思うが，実験的に最新 [GnuPG] を試しておられる人はご注意を。
また，知らない人の鍵はインポートしないように（笑）

アップデートは計画的に。

## ブックマーク

- [OpenPGP の実装]({{< rlnk "/openpgp/" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenPGP]: http://openpgp.org/
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
