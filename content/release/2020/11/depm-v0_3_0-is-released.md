+++
title = "Go パッケージ／モジュールの依存関係可視化ツール Depm v0.3.0 をリリースした"
date =  "2020-11-08T17:14:01+09:00"
description = "これで個人的に欲しい機能は揃ったかな。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "tools", "golang", "package", "module" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

気になる部分の修正といくつか機能追加を行った [spiegel-im-spiegel/depm][depm] の v0.3.0 をリリースした。

- [Release v0.3.0 · spiegel-im-spiegel/depm · GitHub](https://github.com/spiegel-im-spiegel/depm/releases/tag/v0.3.0)

使い方については以下の記事を参照のこと。

- [Depm: Go 言語用モジュール依存関係可視化ツール]({{< ref "/release/dependency-graph-for-golang-modules.md" >}})

気になる部分というのは，ツール内から `go list` コマンドを呼び出している部分で，以下の問題に対応している。

- [Go でサブプロセスを起動する際は LookPath に気をつけろ！](https://zenn.dev/spiegel/articles/20201107-lookpath-by-golang)

これで個人的に欲しい機能は揃ったので少しペースを落とすか。
つか，ほとんどテストを書いてないんだよなぁ...

[Go]: https://golang.org/ "The Go Programming Language"
[depm]: https://github.com/spiegel-im-spiegel/depm "spiegel-im-spiegel/depm: Visualize depndency packages and modules"

## 参考図書

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
