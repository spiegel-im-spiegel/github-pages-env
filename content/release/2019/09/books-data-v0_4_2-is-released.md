+++
title = "書籍データ取得ツール books-data v0.4.2 のリリース，他"
date =  "2019-09-08T15:22:41+09:00"
description = "description"
image = "/images/attention/tools.png"
tags = [ "tools", "book", "books-data", "openbd-api", "aozora-api" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

書籍データ取得ツール [books-data] v0.4.2 をリリースした。

- [Release v0.4.2 · spiegel-im-spiegel/books-data · GitHub](https://github.com/spiegel-im-spiegel/books-data/releases/tag/v0.4.2)

パッチバージョンが妙に進んでいるが気にしないように（笑）

機能的な変更は殆どないが [Go] 1.13 のリリースに併せて内部をいろいろ変更している。
リファクタリングというやつである。

そうそう。
設定ファイルの既定の置き場所を変更して [XDG Base Directory] に対応した。
詳しくは以下の記事を参照のこと。

- [書籍データ取得ツール]({{< ref "/release/books-data.md" >}})

他には [openBD] クライアント側のハンドリング・パッケージを自前のもの  [spiegel-im-spiegel/openbd-api]  に換装した。
[spiegel-im-spiegel/aozora-api] もバージョンアップしている。

- [Go 言語用青空文庫 API クライアント・パッケージ]({{< ref "/release/aozora-api-package-for-golang.md" >}})
- [Go 言語用 openBD クライアント・パッケージ]({{< ref "/release/openbd-api-package-for-golang.md" >}})

## ブックマーク

- [Go 言語用エラーハンドリング・パッケージをリリースした]({{< ref "/release/2019/09/errs-package-is-released.md" >}})
- [spiegel-im-spiegel/gocli v0.10.1 のリリース]({{< ref "/release/2019/09/gocli-v0_10_1-is-released.md" >}})

[books-data]: https://github.com/spiegel-im-spiegel/books-data "spiegel-im-spiegel/books-data: Search for Books Data"
[Go]: https://golang.org/ "The Go Programming Language"
[XDG Base Directory]: https://standards.freedesktop.org/basedir-spec/latest/ "XDG Base Directory Specification"
[openBD]: https://openbd.jp/ "openBD | 書誌情報・書影を自由に"
[spiegel-im-spiegel/aozora-api]: https://github.com/spiegel-im-spiegel/aozora-api "spiegel-im-spiegel/aozora-api: APIs for Aozora-bunko RESTful Service by Golang"
[spiegel-im-spiegel/openbd-api]: https://github.com/spiegel-im-spiegel/openbd-api "spiegel-im-spiegel/openbd-api: APIs for openBD by Golang"
