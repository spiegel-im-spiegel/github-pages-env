+++
title = "Go 製のツールとパッケージをまとめてアップデートした"
date =  "2020-10-02T16:40:54+09:00"
description = "Update gpgpdump, books-data, gnkf, pa-api, aozora-api, openbd-api, gocli, and errs packages"
image = "/images/attention/tools.png"
tags  = [ "tools", "package", "golang", "books-data", "pa-api", "openbd-api", "aozora-api", "error", "gnkf", "cli", "gpgpdump" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GitHub] リポジトリにある [Go 言語][Go]製のツールやパッケージについて [GitHub] Actions で CI (Continuous Integration) を回す目処が立ったので，設定を変更しリリースし直すことにした。
今回リリースしたツール・パッケージのうち主なものは以下の通り。

- [Release v0.9.1 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.9.1)
- [Release v0.5.8 · spiegel-im-spiegel/books-data · GitHub](https://github.com/spiegel-im-spiegel/books-data/releases/tag/v0.5.8)
- [Release v0.1.3 · spiegel-im-spiegel/gnkf · GitHub](https://github.com/spiegel-im-spiegel/gnkf/releases/tag/v0.1.3)
- [Release v0.7.2 · spiegel-im-spiegel/pa-api · GitHub](https://github.com/spiegel-im-spiegel/pa-api/releases/tag/v0.7.2)
- [Release v0.2.6 · spiegel-im-spiegel/aozora-api · GitHub](https://github.com/spiegel-im-spiegel/aozora-api/releases/tag/v0.2.6)
- [Release v0.2.6 · spiegel-im-spiegel/openbd-api · GitHub](https://github.com/spiegel-im-spiegel/openbd-api/releases/tag/v0.2.6)
- [Release v0.10.3 · spiegel-im-spiegel/gocli · GitHub](https://github.com/spiegel-im-spiegel/gocli/releases/tag/v0.10.3)
- [Release v1.0.2 · spiegel-im-spiegel/errs · GitHub](https://github.com/spiegel-im-spiegel/errs/releases/tag/v1.0.2)

若干 lint に叱られたり軽微なバグを指摘されたりしたものもあるが，機能上の追加・変更はない。

それじゃあ，次のステージに行きましょう♪

## ブックマーク

- [OpenPGP パケットを可視化する gpgpdump]({{< ref "/release/gpgpdump.md" >}})
- [書籍データ取得ツール books-data]({{< ref "/release/books-data.md" >}})
- [GNKF: Network Kanji Filter by Golang]({{< ref "/release/gnkf.md" >}})
- [Go 言語用 PA-API v5 クライアント・パッケージ]({{< ref "/release/pa-api-v5.md" >}})
- [Go 言語用青空文庫 API クライアント・パッケージ]({{< ref "/release/aozora-api-package-for-golang.md" >}})
- [Go 言語用 openBD クライアント・パッケージ]({{< ref "/release/openbd-api-package-for-golang.md" >}})
- [Go 言語用 CLI プログラミング支援パッケージ]({{< ref "/release/gocli-package-for-golang.md" >}})
- [Go 言語用エラーハンドリング・パッケージ]({{< ref "/release/errs-package-for-golang.md" >}})

- [golangci-lint を GitHub Actions で使う]({{< ref "/golang/using-golangci-lint-action.md" >}})
- [Go 依存パッケージの脆弱性検査]({{< ref "/golang/check-for-vulns-in-golang-dependencies.md" >}})
- [GitHub Actions でクロス・コンパイル（GoReleaser 編）]({{< ref "/golang/cross-compiling-in-github-actions-with-goreleaser.md" >}})
- [Go のコードでも GitHub Code Scanning が使えるらしい]({{< ref "/remark/2020/10/github-code-scanning-with-golang.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[GitHub]: https://github.com/

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B01MZ8UA8O" %}} <!-- 射手座☆午後九時 Don't be late -->
