+++
title = "gpgpdump v0.13.0 をリリースした"
date =  "2021-11-10T21:43:49+09:00"
description = "ようやく draft-ietf-openpgp-crypto-refresh-04 に追従できた。"
image = "/images/attention/openpgp.png"
tags = ["tools", "openpgp", "gpgpdump", "golang"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP] パケットの内容を可視化する [gpgpdump] の v0.13.0 をリリースした。

- [Release v0.13.0 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.13.0)

先月（2021-10-18）に [draft-ietf-openpgp-crypto-refresh-04](https://datatracker.ietf.org/doc/draft-ietf-openpgp-crypto-refresh/04/ "draft-ietf-openpgp-crypto-refresh-04 - OpenPGP Message Format") が出てたんだけど，ようやく追従できた。
つってもまだ精査できてないんだよね。
細かいところはおいおい直していきます。

というか，そろそろ出力をちゃんと構造化テキストで出せるようにしたいんだけど，本業ががが...

ちなみに [draft-ietf-openpgp-crypto-refresh](https://datatracker.ietf.org/doc/draft-ietf-openpgp-crypto-refresh/) は [draft-ietf-openpgp-rfc4880bis](https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/) を再構成したもので2021年から検討が始まっている。
現行の [RFC 4880] の互換性は維持しつつ AEAD に対応したりと機能拡張が検討されている。

## ブックマーク

- [OpenPGP の実装]({{< rlnk "openpgp/" >}})
- [OpenPGP パケットを可視化する gpgpdump]({{< ref "/release/gpgpdump.md" >}})

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[OpenPGP]: http://openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[RFC 4880]: https://tools.ietf.org/html/rfc4880
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/
[Go]: https://go.dev/

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
