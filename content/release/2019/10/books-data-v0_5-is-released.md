+++
title = "書籍データ取得ツール books-data v0.5 をリリースした"
date =  "2019-10-06T02:02:59+09:00"
description = "今回のメインは PA-API v5 に対応したことだろう。なんか2019年11月末までに v5 に切り替えないといけないらしい。"
image = "/images/attention/tools.png"
tags = [ "tools", "book", "books-data", "pa-api" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

書籍データ取得ツール [books-data] v0.5 をリリースした。

- [Release v0.5.0 · spiegel-im-spiegel/books-data · GitHub](https://github.com/spiegel-im-spiegel/books-data/releases/tag/v0.5.0)
- [Release v0.5.1 · spiegel-im-spiegel/books-data · GitHub](https://github.com/spiegel-im-spiegel/books-data/releases/tag/v0.5.1)
- [Release v0.5.2 · spiegel-im-spiegel/books-data · GitHub](https://github.com/spiegel-im-spiegel/books-data/releases/tag/v0.5.2)
- [Release v0.5.3 · spiegel-im-spiegel/books-data · GitHub](https://github.com/spiegel-im-spiegel/books-data/releases/tag/v0.5.3)

パッチバージョンが妙に進んでいるが気にしないように（笑）

今回のメインは [PA-API] v5 に対応したことだろう。
ざっと見回したが  [PA-API] v5 に対応した [Go 言語]パッケージが見当たらなかったので自作した。

- [Go 言語用 PA-API v5 クライアント・パッケージ]({{< ref "/release/pa-api-v5.md" >}})

なんか2019年11月末までに v5 に切り替えないといけないらしい。
Java と PHP は SDK が用意されているが，他の言語は見当たらなかったので必要なら自作するしかないだろう。

- [PA-API V5 への移行]({{< ref "/remark/2019/10/pa-api-v5.md" >}})

他にはまた少しデータ構成を変えた。
後方互換性が崩れるが，私以外に使っている人はおるまぁ（笑）

詳しくは以下の記事を参照のこと。

- [書籍データ取得ツール books-data]({{< ref "/release/books-data.md" >}})

これで API 周りは全部自作パッケージに換装してしまったよ。
そんなつもりじゃなかったんだけどなぁ。

[books-data]: https://github.com/spiegel-im-spiegel/books-data "spiegel-im-spiegel/books-data: Search for Books Data"
[PA-API]: https://affiliate.amazon.co.jp/assoc_credentials/home "Product Advertising API"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[spiegel-im-spiegel/pa-api]: https://github.com/spiegel-im-spiegel/pa-api "spiegel-im-spiegel/pa-api: APIs for Amazon Product Advertising API v5 by Golang"
