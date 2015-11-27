+++
date = "2015-11-21T14:19:02+09:00"
update = "2015-11-27T14:31:41+09:00"
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

## 開発支援ツールの導入

まずは [Go 言語]用の支援ツールを導入する。

```bash
C:> go get -v golang.org/x/tools/cmd/vet
C:> go get -v golang.org/x/tools/cmd/goimports
C:> go get -v golang.org/x/tools/cmd/oracle
C:> go get -v github.com/golang/lint/golint
C:> go get -v github.com/nsf/gocode
C:> go get -v github.com/rogpeppe/godef
```

[golint] は，いわゆる lint ツール。
[vet] もコードの静的検査ツール。
両方あると幸せになれる。

[goimport] はコード整形ツールで，標準の [gofmt] を置き換えることができ，かつ [gofmt] よりも若干かしこい。
[gocode] は入力補完ツール。
[godef] は指定したシンボルの定義定義元情報を出力するツール（出力を使って定義元へジャンプできる。実際には [oracle] と併用するらしい）。
いずれも vim や emacs などでは有名だが [ATOM] でも使える。

言わずもがなだが，これらのツールには PATH を通しておくこと。
`%GOPATH%\bin` フォルダにパスを通しておけばいいだろう。

## go-plus パッケージの導入

では，上述のツールを操作できる [go-plus] パッケージを導入する。
やり方は Settings 画面（`ctrl+,` で起動）で Install してもいいし `apm` コマンドを使ってもよい。

- [ATOM Editor に関するメモ]({{< relref "remark/2015/atom-editor.md" >}})

[go-plus] には設定項目がいくつかあるが，ほとんど既定値のままで使える。
コマンドパレットから `golang` をキーワードに検索すると山程機能があるのが分かるだろう。

{{< fig-img flickr="true" src="https://farm6.staticflickr.com/5794/22710708563_3d4aca2709.jpg" title="menu of go-plus (ATOM)" link="https://www.flickr.com/photos/spiegel/22710708563/" >}}

既定ではソースファイルを保存する度にコード整形や lint 等が走る。
これを制御したい場合はパッケージの Setting で以下の項目を調整すればよい。

{{< fig-img flickr="true" src="https://farm1.staticflickr.com/735/22767398347_ed9329653a.jpg" title="settings for go-plus (ATOM)" link="https://www.flickr.com/photos/spiegel/22767398347/" >}}

定義ファイルへのジャンプと復帰は `alt-cmd-g` および `alt-shift-cmd-G` にバインドされているが Windows 環境では動かないので（コマンドパレットから起動してもいいのだが）適当なキーに再割当てするといいだろう。
ファンクションキーは結構空いてるので，たとえば

| Keystroke   | Command               | Selector |
|:------------|:----------------------|:---------|
| `f12`       | `golang:godef`        | `atom-text-editor[data-grammar="source go"]:not(.mini)` |
| `shift-f12` | `golang:godef-return` | `atom-text-editor[data-grammar="source go"]:not(.mini)` |

とアサインするなら `%USERPROFILE%\.atom\keymap.cson` に

```cson
'atom-text-editor[data-grammar="source go"]:not(.mini)':
  'f12': 'golang:godef'
  'shift-f12': 'golang:godef-return'
```

と設定すればいい。

lint や定義ファイルのジャンプは `GOPATH` や `GOROOT` を見て外部パッケージを判断しているのだが， [gb] のようなツールでは `GOPATH` をコマンド内部で書き換えて実行するので lint ツールとは整合性が取れなくなる。
[go-plus] の設定では `GOPATH` を上書きすることも可能なので，とりあえずこれで回避する方法もある[^a]。

[^a]: `GOPATH` を [go-plus] の設定で上書きする場合は “Environment Overrides Config” を**無効にする**こと。なんでかこれ，毎回ハマるんだよなぁ。

{{< fig-img flickr="true" src="https://farm6.staticflickr.com/5821/23233956325_0ddf55e61a.jpg" title="settings for go-plus (ATOM)" link="https://www.flickr.com/photos/spiegel/23233956325/" >}}

[gb] への対応は “Planned Features” に挙がってるので，将来的には小細工しなくても [gb] ベースの開発ができるようになるかもしれない。
てか，なってほしい。

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

## go-find-references パッケージが惜しい

[go-find-references] パッケージは [redefiance/go-find-references](https://github.com/redefiance/go-find-references) を使って指定したシンボルを参照しているファイルを列挙してくれる便利ツールだが， Windows 環境ではタグジャンプが上手く動かない。
どうやら `C:` などのドライブレターを上手く処理できないようだ。
とほほ。

## ブックマーク{#bookmark}

- [struct にアノテーションつけてたら go vet . すべき - Qiita](http://qiita.com/amanoiverse/items/fcd25db64f341ad2471f)
- [これからGo言語を書く人への三種の神器 - Qiita](http://qiita.com/osamingo/items/d5ec42fb8587d857310a) : `go vet`, `goimports`, `golint` で正しいコードを書きましょう。

[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[ATOM]: https://atom.io/ "Atom"
[golint]: https://github.com/golang/lint "golang/lint"
[vet]: https://golang.org/cmd/vet/ "vet - The Go Programming Language"
[goimport]: https://godoc.org/golang.org/x/tools/cmd/goimports "goimports - GoDoc"
[gofmt]: https://golang.org/cmd/gofmt/ "gofmt - The Go Programming Language"
[gocode]: https://github.com/nsf/gocode "nsf/gocode"
[godef]: https://github.com/rogpeppe/godef "rogpeppe/godef"
[oracle]: https://godoc.org/golang.org/x/tools/cmd/oracle "oracle - GoDoc"
[go-plus]: https://atom.io/packages/go-plus "go-plus"
[gb]: http://getgb.io/ "gb - A project based build tool for Go"
[language-go]: https://atom.io/packages/language-go "language-go"
[go-find-references]: https://atom.io/packages/go-find-references "go-find-references"