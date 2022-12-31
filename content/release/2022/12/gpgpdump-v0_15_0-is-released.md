+++
title = "gpgpdump v0.15.0 をリリースした"
date =  "2022-12-31T11:42:08+09:00"
description = "ようやく draft-ietf-openpgp-crypto-refresh-07 に追従できた。"
image = "/images/attention/openpgp.png"
tags = ["tools", "openpgp", "gpgpdump", "golang"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP] パケットの内容を可視化する [gpgpdump] の v0.15.0 をリリースした。

- [Release v0.15.0 · goark/gpgpdump · GitHub](https://github.com/goark/gpgpdump/releases/tag/v0.15.0)

2022-10-23 に出た [draft-ietf-openpgp-crypto-refresh-07](https://datatracker.ietf.org/doc/draft-ietf-openpgp-crypto-refresh/07/ "draft-ietf-openpgp-crypto-refresh-07 - OpenPGP Message Format") に追従した。
つっても，あんまりテストしてないんだよね。
とりあえずメーリング・リストに上がってくるサンプルデータは読めるようにした。
ギリギリだけど年内にリリースできてよかったよ。

すンごい直したい。
技術的負債ががが（笑） まぁ[次期 OpenPGP][RFC 4880bis] は（特に AEAD 周りで）アクティブな議論が続いてるので，それらが一段落してからじゃないと諸々に手を付けられない感じ。

まぁ，来年も緩々とやっていきましょう。

## ブックマーク

- [OpenPGP の実装]({{< rlnk "openpgp/" >}})
- [OpenPGP パケットを可視化する gpgpdump]({{< ref "/release/gpgpdump.md" >}})

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[OpenPGP]: http://openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[RFC 4880]: https://tools.ietf.org/html/rfc4880
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-crypto-refresh/
[Go]: https://go.dev/

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
