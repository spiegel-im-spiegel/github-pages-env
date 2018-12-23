+++
title = "2018年お気に入り ATOM パッケージ"
date = "2018-12-17T22:45:24+09:00"
update = "2018-12-18T08:08:33+09:00"
description = "年末なので ATOM エディタの整理を。 "
image = "/images/attention/kitten.jpg"
tags = ["atom", "editor", "tools", "golang"]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

年末なので [ATOM] エディタの整理を。
といっても最近はあまりいじるところがないけど。

## gocode を巡るあれこれ

本当のことを言うと [Go 言語]用の支援パッケージを LSP (Language Server Protocol) ベースにしたいんだけど，最近流行りの [golsp] や [bingo] を扱えるパッケージはなさそう。

以前からある [go-langserver] はバックエンドに gocode を利用しているのだが（今は違う？），本家の [nsf/gocode] が [Go] 1.10 以降で組み込まれたビルドキャッシュ機能に追従できず fork が乱立して混沌としている[^gc1]。

[^gc1]: gocode を巡るゴタゴタについては「[gocode やめます(そして Language Server へ)](https://mattn.kaoriya.net/software/lang/go/20181217000056.htm)」が参考になる。

結局 [ATOM] で [Go 言語]用の支援パッケージを利用するなら [go-plus] 一択で

1. [atom-ide-ui] をインストール，またはインストール済みであることを確認する
2. [go-plus] をインストールする。この際，サブパッケージおよび各種支援ツールが自動インストールされるのを確認する
3. [stamblerre/gocode] を `go get` コマンドで上書きインストールする[^gp1]

[^gp1]: [go-plus] では [mdempsky/gocode] が自動インストールされるが， [mdempsky/gocode] は[モジュール対応モード]に対応していないらしいので [stamblerre/gocode] を手動でインストールする必要がある。実は [ide-go] で [go-langserver] を駆動させていたのだが [ide-go] 自体が [go-plus] を要求するので意味がなかったっぽい。

という手順になるだろう。
[golsp] や [bingo] が使えれば将来的に楽なんだけどねぇ。

それ以前に 1.12 で[モジュール対応モード]に本格的に移行した場合に [go-plus] が追従できるかどうかなんだけど。
最悪の場合は [Visual Studio Code] に乗り換えるか（vim は若い頃のデスマーチなトラウマがフラッシュバックするので使いたくない）。

## [atom-ide-ui] と連携する [document-outline]

今まで markdown テキストのアウトライン表示は [nav-panel-plus] を使ってたんだけど [atom-ide-ui] に対応していないのが欠点だった。
で，色々探してみたら [document-outline] が良さげである。
[document-outline] の設定で “show by default” 項目を無効にしておけば [atom-ide-ui] のアウトライン表示のタブだけが表示される。

[document-outline] は以下の構造化テキストに対応している

- Markdown (Commonmark)
- LaTeX
- ReStructuredText
- AsciiDoc
- Knitr

らしいんだけど LaTeX ファイルのアウトラインが上手く表示されないんだよなぁ。
うーむ。

他にも [ide-html] で HTML を，[ide-css] で CSS を，[ide-yaml] で YAML を，[ide-json] で JSON を扱える。
これでメジャーな構造化テキストは大体扱えるかな[^om1]。

[^om1]: 私は使わないが [org-mode] パッケージもあるらしい。

## ブックマーク

- [年末なので ATOM Editor を掃除しましょう（もしくは2017年お気に入り ATOM パッケージ）]({{< ref "/remark/2017/12/favorite-atom-packages-2017.md" >}})
- [ATOM エディタを使った作図（PlantUML 編）]({{< ref "/remark/2017/12/plantuml-with-atom.md" >}})

[ATOM]: https://atom.io/ "Atom"
[Visual Studio Code]: https://code.visualstudio.com/ "Visual Studio Code - Code Editing. Redefined"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Go]: https://golang.org/ "The Go Programming Language"
[golsp]: https://godoc.org/golang.org/x/tools/cmd/golsp "golsp - GoDoc"
[bingo]: https://github.com/saibing/bingo "saibing/bingo: Bingo is a Go language server that speaks Language Server Protocol."
[go-langserver]: https://github.com/sourcegraph/go-langserver "sourcegraph/go-langserver: Go language server to add Go support to editors and other tools that use the Language Server Protocol (LSP)"
[nsf/gocode]: https://github.com/nsf/gocode "nsf/gocode: An autocompletion daemon for the Go programming language"
[mdempsky/gocode]: https://github.com/mdempsky/gocode "mdempsky/gocode: An autocompletion daemon for the Go programming language"
[stamblerre/gocode]: https://github.com/stamblerre/gocode "stamblerre/gocode: An autocompletion daemon for the Go programming language"
[atom-ide-ui]: https://atom.io/packages/atom-ide-ui
[go-plus]: https://atom.io/packages/go-plus
[ide-go]: https://atom.io/packages/ide-go
[モジュール対応モード]: {{< ref "/golang/go-module-aware-mode.md" >}} "モジュール対応モードへの移行を検討する"
[nav-panel-plus]: https://atom.io/packages/nav-panel-plus
[document-outline]: https://atom.io/packages/document-outline
[ide-html]: https://atom.io/packages/ide-html
[ide-css]: https://atom.io/packages/ide-css
[ide-yaml]: https://atom.io/packages/ide-yaml
[ide-json]: https://atom.io/packages/ide-json
[org-mode]: https://atom.io/packages/org-mode
