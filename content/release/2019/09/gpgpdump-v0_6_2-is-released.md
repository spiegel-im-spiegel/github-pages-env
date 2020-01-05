+++
title = "gpgpdump v0.6.2 をリリースした"
date =  "2019-09-09T19:38:32+09:00"
description = "今回は細かい調整のみ。 draft-ietf-openpgp-rfc4880bis-08 への追随も行った。"
image = "/images/attention/openpgp.png"
tags = ["tools", "openpgp", "golang", "gpgpdump"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP] パケットの内容を可視化する [gpgpdump] の v0.6.2 をリリースした。

- [Release v0.6.2 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.6.2)

といっても今回は細かい調整のみ。
[Go] 1.13 に対応したのと新しいエラーハンドリング・パッケージ [spiegel-im-spiegel/errs] を組み込んだだけ。
でも，これでデバッグが格段にし易くなった。

あ，そうそう。
[draft-ietf-openpgp-rfc4880bis-08](https://tools.ietf.org/html/draft-ietf-openpgp-rfc4880bis-08 "draft-ietf-openpgp-rfc4880bis-08 - OpenPGP Message Format") への追随も行った。
でも試す対象がないからなぁ（笑）

[gpgpdump] の使い方等は以下を参照のこと。

- [OpenPGP パケットを可視化する gpgpdump]({{< ref "/release/gpgpdump.md" >}})

## ブックマーク

- [Go 言語用エラーハンドリング・パッケージをリリースした]({{< ref "/release/2019/09/errs-package-is-released.md" >}})
- [spiegel-im-spiegel/gocli v0.10.1 のリリース]({{< ref "/release/2019/09/gocli-v0_10_1-is-released.md" >}})

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[pgpdump]: http://www.mew.org/~kazu/proj/pgpdump/ "pgpdump"
[OpenPGP]: http://openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Go]: https://golang.org/ "The Go Programming Language"
[spiegel-im-spiegel/errs]: https://github.com/spiegel-im-spiegel/errs "spiegel-im-spiegel/errs: Error handling for Golang"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
