+++
date = "2015-11-21T14:19:02+09:00"
update = "2016-03-19T14:39:05+09:00"
description = "ATOM Editor で Go 言語のコーディング環境を整える。 go-plus パッケージの導入について。"
draft = false
tags = ["golang", "engineering", "tools", "atom", "editor"]
title = "ATOM で Go"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

今回もまた横道に逸れてツールの話。
[ATOM] Editor で [Go 言語]のコーディング環境を整える。
環境を作る度に「どうだったっけ」とあちこちサイトを巡るので，覚え書きとしてまとめておく。

例によって Windows 環境を前提にしているので，他の環境の方は適当に脳内補完してください。

（2016年3月： [go-plus] バージョン 4 以降で大きく構成が変わったので改訂した）

## 開発支援ツールの導入

まずは [Go 言語]用の支援ツールを導入する（更新時には `-u` オプションを付ける）。

```text
$ go get -v github.com/alecthomas/gometalinter
$ go get -v golang.org/x/tools/cmd/gorename
$ go get -v github.com/nsf/gocode
$ go get -v github.com/rogpeppe/godef
$ go get -v golang.org/x/tools/cmd/oracle
```

[gometalinter] は所謂 lint ツールなのだが，単独で動作するのではなく，巷にいくつかある lint ツール（標準の [vet] を含む）を統合的に管理することができる。
以下のコマンドで [gometalinter] が使用する lint ツールをまとめてインストールする。

```text
$ gometalinter --install --update
Installing:
  structcheck
  interfacer
  goconst
  golint
  goimports
  dupl
  errcheck
  aligncheck
  gocyclo
  ineffassign
  unconvert
  gotype
  varcheck
  deadcode
  lll
```

[gorename] は関数や変数の名前を変更したい時に使うツールで，文法を解釈してくれるため副作用が少ないのが特徴。
[gocode] は入力補完ツール。
[godef] は指定したシンボルの定義定義元情報を出力するツール（出力を使って定義元へジャンプできる。実際には [oracle] と併用するらしい）。
いずれも vim や emacs などでは有名だが [ATOM] でも使える。

言わずもがなだが，これらのツールには PATH を通しておくこと。
`%GOPATH%\bin` フォルダにパスを通しておけばいいだろう。

## go-plus パッケージの導入

では，上述のツールを操作できる [go-plus] パッケージを導入する。
やり方は Settings 画面（`ctrl+,` で起動）で Install してもいいし `apm` コマンドを使ってもよい。

- [ATOM Editor に関するメモ]({{< relref "remark/2015/atom-editor.md" >}})

最近の [go-plus] は複数のサブ・パッケージで構成されているらしい。
[go-plus] を導入すると以下のサブ・パッケージも自動的に導入される。

- [autocomplete-go](https://atom.io/packages/autocomplete-go) : [gocode] を使って入力補完
- [go-config](https://atom.io/packages/go-config) : [Go 言語]用ツール等のチェック
- [go-get](https://atom.io/packages/go-get) : [Go 言語]用ツールを取得する
- [gofmt](https://atom.io/packages/gofmt) : [gofmt] または [goimport] を使用したフォーマッタ[^gf]
- [gometalinter-linter](https://atom.io/packages/gometalinter-linter) : [gometalinter] を使った lint
- [gorename](https://atom.io/packages/gorename) : [gorename] を使ってリネーム
- [navigator-godef](https://atom.io/packages/navigator-godef) : [godef] を使って定義元へジャンプ
- [tester-go](https://atom.io/packages/tester-go) : テストの実行

[^gf]: [goimport] はコード整形ツールで，標準の [gofmt] を置き換えることができ，かつ [gofmt] よりも若干かしこい。 [gometalinter] からインストールされる。

コマンドパレットから呼び出されるコマンドもかなり整理されているようだ。

{{< fig-img flickr="true" src="https://farm6.staticflickr.com/5794/22710708563_f49bdbb61c.jpg" title="menu for go-plus (ATOM)" link="https://www.flickr.com/photos/spiegel/22710708563/" >}}

[go-plus] ではサブ・パッケージごとに設定項目がいくつかあるが，ほとんど既定値のままで使える。
たとえば [gofmt](https://atom.io/packages/gofmt) サブ・パッケージの設定画面は以下のようになっている。

{{< fig-img flickr="true" src="https://farm1.staticflickr.com/735/22767398347_86d14e29f9.jpg" title="settings for gofmt@go-plus (ATOM)" link="https://www.flickr.com/photos/spiegel/22767398347/" >}}

定義元へのジャンプと復帰は `alt-cmd-g` および `alt-shift-cmd-G` にバインドされているが Windows 環境では動かないので（コマンドパレットから起動してもいいのだが）適当なキーに再割当てするといいだろう。
ファンクションキーは結構空いてるので，たとえば

| Keystroke   | Command               | Selector |
|:------------|:----------------------|:---------|
| `f12`       | `golang:godef`        | `atom-text-editor[data-grammar~="go"]:not([mini])` |
| `shift-f12` | `golang:godef-return` | `atom-text-editor[data-grammar~="go"]:not([mini])` |

とアサインするなら `%USERPROFILE%\.atom\keymap.cson` に

```cson
'atom-text-editor[data-grammar~="go"]:not([mini])':
  'f12': 'golang:godef'
  'shift-f12': 'golang:godef-return'
```

と設定すればいい。

lint や定義元のジャンプは `GOPATH` や `GOROOT` を見て外部パッケージを判断しているのだが， [gb] のようなツールでは `GOPATH` をコマンド内部で書き換えて実行するので lint ツールとは整合性が取れなくなる。
[go-plus] の設定では `GOPATH` を上書きすることも可能なので，とりあえずこれで回避する方法もある[^a]。

[^a]: `GOPATH` を [go-plus] の設定で上書きする場合は “Environment Overrides Config” を**無効にする**こと。なんでかこれ，毎回ハマるんだよなぁ。

{{< fig-img src="https://farm6.staticflickr.com/5821/23233956325_0d13c7379f.jpg" title="settings for go-plus (ATOM)" link="https://www.flickr.com/photos/spiegel/23233956325/" >}}

（[go-plus] 4.0.1 および [gometalinter-linter](https://atom.io/packages/gometalinter-linter) 1.0.2 で上記の設定が効いてない模様。コマンドプロンプトななどで環境変数 `GOPATH` を上書き設定してからプロジェクト・フォルダ上で `atom.com .` と起動すれば上手くいくようだ）

## language-go パッケージは同梱済み

[language-go] は Core パッケージに入っているためインストール時点で既に入っている[^b]。
[language-go] の機能で目を引くのはやはり Snippets だろう。
以下はその一部（パッケージの Settings に一覧がある）。

{{< fig-img flickr="true" src="https://farm6.staticflickr.com/5675/22712254763_f8fb9f6735.jpg" title="Snippets for golang (ATOM)" link="https://www.flickr.com/photos/spiegel/22712254763/" >}}

こんなよぅけ覚えれるか！ まぁとりあえず，よく使うものだけ覚えておけばいいのだろうけど。

[^b]: `language` でインストール済みパッケージを検索するとメジャーな言語は大抵入っているのが分かる。

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

## ブックマーク

- [struct にアノテーションつけてたら go vet . すべき - Qiita](http://qiita.com/amanoiverse/items/fcd25db64f341ad2471f)
- [これからGo言語を書く人への三種の神器 - Qiita](http://qiita.com/osamingo/items/d5ec42fb8587d857310a) : `go vet`, `goimports`, `golint` で正しいコードを書きましょう。
- [Big Sky :: golang のリファクタリングには gofmt ではなく、gorename を使おう。](http://mattn.kaoriya.net/software/lang/go/20150113141338.htm)
- [gometalinter で楽々 lint - Qiita](http://qiita.com/spiegel-im-spiegel/items/238f6f0ee27bdf1de2a0)

[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[ATOM]: https://atom.io/ "Atom"
[golint]: https://github.com/golang/lint "golang/lint"
[vet]: https://golang.org/cmd/vet/ "vet - The Go Programming Language"
[gometalinter]: https://github.com/alecthomas/gometalinter "alecthomas/gometalinter: Concurrently run Go lint tools and normalise their output"
[goimport]: https://godoc.org/golang.org/x/tools/cmd/goimports "goimports - GoDoc"
[gofmt]: https://golang.org/cmd/gofmt/ "gofmt - The Go Programming Language"
[gorename]: https://godoc.org/golang.org/x/tools/cmd/gorename "gorename - GoDoc"
[gocode]: https://github.com/nsf/gocode "nsf/gocode"
[godef]: https://github.com/rogpeppe/godef "rogpeppe/godef"
[oracle]: https://godoc.org/golang.org/x/tools/cmd/oracle "oracle - GoDoc"
[go-plus]: https://atom.io/packages/go-plus "go-plus"
[gb]: http://getgb.io/ "gb - A project based build tool for Go"
[language-go]: https://atom.io/packages/language-go "language-go"
[go-find-references]: https://atom.io/packages/go-find-references "go-find-references"