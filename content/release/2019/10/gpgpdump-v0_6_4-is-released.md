+++
title = "gpgpdump v0.6.4 をリリースした"
date =  "2019-10-18T23:33:08+09:00"
description = "主目的は Go 1.13.3 でリコンパイルすることだったり。というわけで，よろしければ更新しておいてください。"
image = "/images/attention/openpgp.png"
tags = ["tools", "openpgp", "golang", "gpgpdump"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP] パケットの内容を可視化する [gpgpdump] の v0.6.4 をリリースした。

- [Release v0.6.4 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.6.4)

といっても，今回は機能的な変更・修正はほとんどなくて（[spiegel-im-spiegel/errs] パッケージを更新したくらい），主目的は [Go] 1.13.3 でリコンパイルすることだったり。

実は

- [openpgp-wg / openpgp-samples · GitLab](https://gitlab.com/openpgp-wg/openpgp-samples)

というリポジトリができていて，これを使って検証とテストを行おうとしたのだが，テストするまでもなく全く問題なく表示されたという（笑）

というわけで，よろしければ更新しておいてください。

## ブックマーク

- [Go 1.13.2 および Go 1.13.3 のリリース【セキュリティ・アップデート】]({{< ref "/release/2019/10/go-1_13_2-is-released.md" >}})

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[OpenPGP]: http://openpgp.org/
[Go]: https://golang.org/ "The Go Programming Language"
[spiegel-im-spiegel/errs]: https://github.com/spiegel-im-spiegel/errs "spiegel-im-spiegel/errs: Error handling for Golang"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
