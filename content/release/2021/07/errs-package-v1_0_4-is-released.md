+++
title = "spiegel-im-spiegel/errs パッケージ v1.0.4 をリリースした"
date =  "2021-07-15T20:42:27+09:00"
description = "不具合修正。マジすんません 🙇"
image = "/images/attention/go-logo_blue.png"
tags  = [ "golang", "package", "module", "error" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Zenn で [pkg/errors] の不具合（？）について[教えてもらった](https://zenn.dev/nekoshita/articles/097e00c6d3d1c9#comment-1dd0f100389e4e "今goのエラーハンドリングを無難にしておく方法（2021.07現在）")。

- [Unwrap doesn't return the base error · Issue #223 · pkg/errors · GitHub](https://github.com/pkg/errors/issues/223)

これを見たときは完全に他人事だったのだが，よく考えたら拙作の [spiegel-im-spiegel/errs][`errs`] では [pkg/errors] の設計コンセプトを参考にしているのでもしかして... と思って調べたら別のところで不具合を見つけてしまった `orz`

具体的には [`errs`]`.New()` 関数で生成したインスタンスに標準の [`errors`]`.Unwrap()` をかけても 返り値が `nil` にならないというもの。
ちうわけで，修正版をリリースした。

- [Release v1.0.4 · spiegel-im-spiegel/errs · GitHub](https://github.com/spiegel-im-spiegel/errs/releases/tag/v1.0.4)

言い訳するなら，実際のエラーハンドリングで [`errors`]`.Unwrap()` を使って手動でひとつずつ階層を遡るとかしないよね，普通。
いや，マジすんません {{% emoji "ペコン" %}}

エラーハンドリングについて，最近読んで面白いと思ったのが以下の記事。

- [Goエラーハンドリング戦略](https://zenn.dev/nobonobo/articles/0b722c9c2b18d5)

特にクラウド環境下でロギングと絡めたハンドリングは参考になるだろう。

この記事の中で「サードパーティのロガーやエラーハンドラは極力使わない」というのは半分くらいは同意。
でも，サードパーティの汎用パッケージじゃなくても，結局は対象となるツール・サービスの設計に合わせて「自分で作ったほうが早い」ってのはあるんだよなぁ。
だから私も自前で作ってるんだけど（笑）

## ブックマーク

- [Go 言語用エラーハンドリング・パッケージ]({{< ref "/release/errs-package-for-golang.md" >}})
- [Go のエラーハンドリング](https://zenn.dev/spiegel/books/error-handling-in-golang) : Zenn 本書きました

[Go]: https://golang.org/ "The Go Programming Language"
[pkg/errors]: https://github.com/pkg/errors "pkg/errors: Simple error handling primitives"
[`errs`]: https://github.com/spiegel-im-spiegel/errs "spiegel-im-spiegel/errs: Error handling for Golang"
[`errors`]: https://pkg.go.dev/errors "errors · pkg.go.dev"

## 参考図書

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
