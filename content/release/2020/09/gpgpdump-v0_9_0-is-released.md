+++
title = "gpgpdump v0.9.0 をリリースした"
date =  "2020-09-01T13:35:34+09:00"
description = "draft-ietf-openpgp-rfc4880bis-10 がリリースされたので，その反映。"
image = "/images/attention/openpgp.png"
tags = ["tools", "openpgp", "gpgpdump", "golang"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP] パケットの内容を可視化する [gpgpdump] の v0.9.0 をリリースした。

- [Release v0.9.0 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.9.0)

使い方は以下を参照のこと。

- [OpenPGP パケットを可視化する gpgpdump]({{< ref "/release/gpgpdump.md" >}})

今回の主な変更は， [draft-ietf-openpgp-rfc4880bis-10](https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/10/ "draft-ietf-openpgp-rfc4880bis-10 - OpenPGP Message Format") がリリースされたので，その反映。

大きい変更としては “Key Block” サブパケットの追加があるのだが

{{< fig-quote type="markdown" title="draft-ietf-openpgp-rfc4880bis-10" link="https://www.ietf.org/id/draft-ietf-openpgp-rfc4880bis-10.html" lang="en" >}}
{{% quote %}}This subpacket MAY be used to convey key data along with a signature of class 0x00, 0x01, or 0x02. It MUST contain the key used to create the signature; either as the primary key or as a subkey. The key SHOULD contain a primary or subkey capable of encryption and the entire key must be a valid OpenPGP key including at least one User ID packet and the corresponding self-signatures{{% /quote %}}.
{{< /fig-quote >}}

例によってサンプルがないので「こんな感じかなぁ」と適当に実装している。
再帰処理になるので，もしかしたらパニクるかもしれない。
あらかじめゴメンペコン。


## ブックマーク

- [OpenPGP の実装]({{< rlnk "openpgp/" >}})

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[OpenPGP]: http://openpgp.org/
[RFC 4880]: https://tools.ietf.org/html/rfc4880
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/
[Go]: https://go.dev/

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
