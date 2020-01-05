+++
title = "PA-API v5 クライアント・パッケージ v0.5.0 をリリースした"
date =  "2019-10-20T20:29:51+09:00"
description = "コメントとかドキュメントはまだまだ整理中なので，長い目で見てやってください。"
image = "/images/attention/go-logo_blue.png"
tags = [ "package", "module", "golang", "amazon", "pa-api", "programming" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go 言語]用 PA-API v5 クライアント・パッケージ [spiegel-im-spiegel/pa-api] v0.5.0 をリリースした。

- [Release v0.5.0 · spiegel-im-spiegel/pa-api · GitHub](https://github.com/spiegel-im-spiegel/pa-api/releases/tag/v0.5.0)

[以前いただいた pull request]({{< relref "./pa-api-v5-package-v0_4_0-is-released.md" >}}) のコードがなかなかよかったので [`paapi5`]`/query` サブパッケージを全面的に書き直した。

もともと [`paapi5`]`/query` サブパッケージは [`paapi5`]`.Query` インタフェースの実装サンプル・コード程度にしか考えてなかったのだが， [PA-API] の `GetItems` および `SearchItems` オペレーションに限れば，そのまま使えそうな感じになってしまった。
PR を送っていただいた方にはホンマに感謝である。

コメントとかドキュメントはまだまだ整理中なので，長い目で見てやってください。
あっ  PR はいつでも歓迎です。

## ブックマーク

- [PA-API v5 への移行]({{< ref "/remark/2019/10/pa-api-v5.md" >}})
- [Go 言語用 PA-API v5 クライアント・パッケージ]({{< ref "/release/pa-api-v5.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[PA-API]: https://affiliate.amazon.co.jp/assoc_credentials/home "Product Advertising API"
[spiegel-im-spiegel/pa-api]: https://github.com/spiegel-im-spiegel/pa-api "spiegel-im-spiegel/pa-api: APIs for Amazon Product Advertising API v5 by Golang"
[`paapi5`]: https://github.com/spiegel-im-spiegel/pa-api "spiegel-im-spiegel/pa-api: APIs for Amazon Product Advertising API v5 by Golang"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
