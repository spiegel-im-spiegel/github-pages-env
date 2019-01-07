+++
date = "2015-11-21T14:19:02+09:00"
update = "2017-12-17T08:37:00+09:00"
description = "ATOM エディタで Go 言語のコーディング環境を整える。 go-plus パッケージの導入について。"
draft = false
tags = ["golang", "engineering", "tools", "atom", "editor"]
title = "ATOM で Go"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"
+++

今回もまた横道に逸れてツールの話。
[ATOM] エディタで [Go 言語]のコーディング環境を整える。
環境を作る度に「どうだったっけ」とあちこちサイトを巡るので，覚え書きとしてまとめておく。

例によって Windows 環境を前提にしているので，他の環境の方は適当に脳内補完してください。

{{% div-box %}}
- *2016年3月：* [go-plus](https://atom.io/packages/go-plus) バージョン 4 以降で大きく構成が変わったので改訂した
- *2017年12月：* [go-plus](https://atom.io/packages/go-plus) がかなり大きく変わってるので改訂
{{% /div-box %}}

## [go-plus] パッケージの導入

現在は [go-plus] パッケージにより [Go 言語]用の支援ツール（[alecthomas/gometalinter] 他）を含めて一括で導入・管理できる。
[go-plus] パッケージ自身の導入については [ATOM] の Settings 画面（`ctrl+,` で起動）で Install してもいいし `apm` コマンドでも可能である。

```text
$ apm install go-plus
```

[go-plus] から以下のパッケージの導入も勧められる。
こちらもお好みでどうぞ。

- [atom-ide-ui] : [ATOM 用統合環境](https://ide.atom.io/ "Atom IDE")（統合環境の機能自体は使わない）
- [go-debug] : [derekparker/delve] をバックエンドにしたデバッガ
    - [Golangのデバッガdelveの使い方 - Qiita](https://qiita.com/minamijoyo/items/4da68467c1c5d94c8cd7)
- [go-signature-statusbar] : 入力支援
- [hyperclick] : マウス・クリックで定義元へジャンプできる（[atom-ide-ui] が入ってる場合は不要？）

あらかじめ [Go 言語]支援ツール用に環境変数 `GOPATH` の設定を済ませておくこと。
また `%GOPATH%/bin` フォルダに `PATH` を通しておくこと。

## [go-plus] パッケージの設定

[go-plus] には支援ツールごとに設定項目がいくつかあるが，ほとんど既定値のままで使える。
たとえばフォーマッタの設定画面は以下のようになっている。

{{< fig-img flickr="true" src="https://farm1.staticflickr.com/735/22767398347_86d14e29f9.jpg" title="settings for gofmt@go-plus (ATOM)" link="https://www.flickr.com/photos/spiegel/22767398347/" >}}

コマンドパレットからも機能を呼び出しできる。

{{< fig-img flickr="true" src="https://farm6.staticflickr.com/5794/22710708563_f49bdbb61c.jpg" title="menu for go-plus (ATOM)" link="https://www.flickr.com/photos/spiegel/22710708563/" >}}

## シンボルの定義元へジャンプする

定義元へのジャンプと復帰は `ctrl-alt-g` および `ctrl-alt-shift-G` にバインドされているが， `alt` キーは個人的に使いにくいので，以下のように割り当て直した。

| Keystroke   | Command               | Selector                                           |
|:----------- |:--------------------- |:-------------------------------------------------- |
| `f12`       | `golang:godef`        | `atom-text-editor[data-grammar~="go"]:not([mini])` |
| `shift-f12` | `golang:godef-return` | `atom-text-editor[data-grammar~="go"]:not([mini])` |

具体的には `%USERPROFILE%\.atom\keymap.cson` に以下の記述を追加する。

```cson
'atom-text-editor[data-grammar~="go"]:not([mini])':
  'f12': 'golang:godef'
  'shift-f12': 'golang:godef-return'
```

なお [hyperclick] パッケージがあればマウス操作で定義元へジャンプできる。
まぁ，コーディング中は殆どマウスは使わないだろうけど。

## language-go パッケージは同梱済み

[language-go] は Core パッケージに入っているためインストール時点で既に入っている[^b]。
[language-go] の機能で目を引くのはやはり snippets だろう。
以下はその一部（パッケージの Settings に一覧がある）。

[^b]: `language` でインストール済みパッケージを検索するとメジャーな言語は大抵入っているのが分かる。

{{< fig-img flickr="true" src="https://farm6.staticflickr.com/5675/22712254763_f8fb9f6735.jpg" title="Snippets for golang (ATOM)" link="https://www.flickr.com/photos/spiegel/22712254763/" >}}

こんなよぅけ覚えれるか！ まぁとりあえず，よく使うものだけ覚えておけばいいのだろうけど。

使い方は，トリガーとなる文字列を入力して tab キーを押す。
たとえば `func` と入力して tab キーを押すと

```go
func ()  {

}
```

と展開される。
`iferr` なら

```go
if err != nil {
    return
}
```

となる。

なお [go-signature-statusbar] パッケージがあれば snippets の入力支援をしてくれるので，多少は気楽にできるかも。

## ブックマーク

- [struct にアノテーションつけてたら go vet . すべき - Qiita](http://qiita.com/amanoiverse/items/fcd25db64f341ad2471f)
- [これからGo言語を書く人への三種の神器 - Qiita](http://qiita.com/osamingo/items/d5ec42fb8587d857310a) : `go vet`, `goimports`, `golint` で正しいコードを書きましょう。
- [Big Sky :: golang のリファクタリングには gofmt ではなく、gorename を使おう。](http://mattn.kaoriya.net/software/lang/go/20150113141338.htm)
- [gometalinter で楽々 lint - Qiita](http://qiita.com/spiegel-im-spiegel/items/238f6f0ee27bdf1de2a0)

[Go 言語に関するブックマーク集はこちら]({{< relref "bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[ATOM]: https://atom.io/ "Atom"
[go-debug]: https://atom.io/packages/go-debug
[go-signature-statusbar]: https://atom.io/packages/go-signature-statusbar
[go-plus]: https://atom.io/packages/go-plus "go-plus"
[hyperclick]: https://atom.io/packages/hyperclick
[atom-ide-ui]: https://atom.io/packages/atom-ide-ui
[language-go]: https://atom.io/packages/language-go "language-go"
[alecthomas/gometalinter]: https://github.com/alecthomas/gometalinter "GitHub - alecthomas/gometalinter: Concurrently run Go lint tools and normalise their output"
[derekparker/delve]: https://github.com/derekparker/delve "GitHub - derekparker/delve: Delve is a debugger for the Go programming language."

[golint]: https://github.com/golang/lint "golang/lint"
[vet]: https://golang.org/cmd/vet/ "vet - The Go Programming Language"
[gometalinter]: https://github.com/alecthomas/gometalinter "alecthomas/gometalinter: Concurrently run Go lint tools and normalise their output"
[goimport]: https://godoc.org/golang.org/x/tools/cmd/goimports "goimports - GoDoc"
[gofmt]: https://golang.org/cmd/gofmt/ "gofmt - The Go Programming Language"
[gorename]: https://godoc.org/golang.org/x/tools/cmd/gorename "gorename - GoDoc"
[gocode]: https://github.com/nsf/gocode "nsf/gocode"
[godef]: https://github.com/rogpeppe/godef "rogpeppe/godef"
[oracle]: https://godoc.org/golang.org/x/tools/cmd/oracle "oracle - GoDoc"
[gb]: http://getgb.io/ "gb - A project based build tool for Go"
[go-find-references]: https://atom.io/packages/go-find-references "go-find-references"
