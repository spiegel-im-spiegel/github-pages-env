+++
title = "gpgpdump v0.6.5 をリリースした"
date =  "2019-11-07T21:06:11+09:00"
description = "ドラフト版の仕様に追随できてない部分があったので，署名サブパケットの解析処理をいくつか修正した。"
image = "/images/attention/openpgp.png"
tags = ["tools", "openpgp", "gpgpdump"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP] パケットの内容を可視化する [gpgpdump] の v0.6.5 をリリースした。

- [Release v0.6.5 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.6.5)

ドラフト版 [RFC 4880bis] の仕様に追随できてない部分があったので，署名パケット（Tag 2）内のサブパケットの解析処理をいくつか修正した。

そろそろちゃんとテストを書かないとなぁ。
一応 [RFC 4880bis] が正式版になるまでには何とかしようという気はある。

まぁ鼻の先は動いているので気長にお付き合いください。

## ブックマーク

- [OpenPGP パケットを可視化する gpgpdump]({{< ref "/release/gpgpdump.md" >}})
- [GitHub に登録されている OpenPGP 公開鍵の情報を取得する]({{< ref "/remark/2019/10/openpgp-public-keys-in-github.md" >}})

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[OpenPGP]: http://openpgp.org/
[RFC 4880]: https://tools.ietf.org/html/rfc4880
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/
[RFC 5581]: http://tools.ietf.org/html/rfc5581
[RFC 6637]: http://tools.ietf.org/html/rfc6637
[Go]: https://golang.org/ "The Go Programming Language"
[spiegel-im-spiegel/errs]: https://github.com/spiegel-im-spiegel/errs "spiegel-im-spiegel/errs: Error handling for Golang"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
