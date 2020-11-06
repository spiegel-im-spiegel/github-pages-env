+++
title = "PA-APIv5 クライアント・パッケージ v0.8.0 をリリースした"
date =  "2020-11-06T21:01:53+09:00"
description = "今回は Marketplace 定義の追加のみ。"
image = "/images/attention/go-logo_blue.png"
tags = [ "package", "module", "golang", "amazon", "pa-api" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

ありがたいことに pull request をいただきまして， Amazon [Product Advertising API][PA-API] v5 へアクセスできる [Go 言語][Go]用パッケージ [spiegel-im-spiegel/pa-api] の v0.8.0 をリリースした。

- [Release v0.8.0 · spiegel-im-spiegel/pa-api · GitHub](https://github.com/spiegel-im-spiegel/pa-api/releases/tag/v0.8.0)

今回は Marketplace 定義の追加のみ。

[spiegel-im-spiegel/pa-api] では Marketplace を[列挙型]({{< ref "/golang/enumeration.md" >}} "列挙型での遊び方")で定義しているのだが，頻繁に追加や変更があるなら上手いやり方ではないかもしれない。
いっそ interface 型で汎化してしまうか。
手を入れるのは簡単だが，今のところは考え中。

[Go]: https://golang.org/ "The Go Programming Language"
[spiegel-im-spiegel/pa-api]: https://github.com/spiegel-im-spiegel/pa-api "spiegel-im-spiegel/pa-api: APIs for Amazon Product Advertising API v5 by Golang"
[PA-API]: https://affiliate.amazon.co.jp/assoc_credentials/home "Product Advertising API"
<!-- eof -->
