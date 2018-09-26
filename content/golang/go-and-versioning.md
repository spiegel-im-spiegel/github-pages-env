+++
title = "vgo (Versioned Go) に関する覚え書き"
date = "2018-02-24T16:10:44+09:00"
update = "2018-03-01T09:31:39+09:00"
description = "Go 言語の次のバージョン（v1.11）から vgo (Versioned Go) を実装する計画があるようで， vgo 関連のドキュメントが公開されている。"
image = "https://research.swtch.com/impver.png"
tags        = [ "golang", "engineering", "versioning", "module" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go 言語]の次のバージョン（v1.11）から vgo (Versioned Go) を実装する計画があるようで， vgo 関連のドキュメントが公開されている。

- [research!rsc: Go & Versioning](https://research.swtch.com/vgo)
    - [research!rsc: Go += Package Versioning (Go & Versioning, Part 1)](https://research.swtch.com/vgo-intro)
    - [research!rsc: A Tour of Versioned Go (vgo) (Go & Versioning, Part 2)](https://research.swtch.com/vgo-tour)
    - [research!rsc: Semantic Import Versioning (Go & Versioning, Part 3)](https://research.swtch.com/vgo-import)
    - [research!rsc: Minimal Version Selection (Go & Versioning, Part 4)](https://research.swtch.com/vgo-mvs)
    - [research!rsc: Reproducible, Verifiable, Verified Builds (Go & Versioning, Part 5)](https://research.swtch.com/vgo-repro)
    - [research!rsc: Defining Go Modules (Go & Versioning, Part 6)](https://research.swtch.com/vgo-module)
    - [research!rsc: Versioned Go Commands (Go & Versioning, Part 7)](https://research.swtch.com/vgo-cmd)

vgo は新しいパッケージのバージョン管理機能で，vendoring 機能を使った [dep] や [glide] のような仕組みとは異なるアプローチらしい。
まず v1.11 で試験的に導入し， v1.12 で正式に導入することを目指しているようだ。
最終的に vgo が従来の go コマンドから完全に置き換えられることになれば `go get` を削除することも考えてるみたい。

{{< fig-quote title="Go += Package Versioning" link="https://research.swtch.com/vgo-intro" lang="en" >}}
<q>I would like Go 1.11 to ship with preliminary support for Go modules, as a kind of technology preview, and then I'd like Go 1.12 to ship with official support. In some later release, we'll remove support for the old, unversioned go get. That's an aggressive schedule, though, and if getting the functionality right means waiting for later releases, we will.</q>
{{< /fig-quote >}}

## お試し vgo

vgo のプロトタイプ版があるようなのでちょっとだけ試してみる。
なお，以下は “[A Tour of Versioned Go]” からの拝借なのでご注意を。

vgo のプロトタイプ版は `go get` で取得できる。

```text
$ go get -u golang.org/x/vgo
```

次に以下のコードを用意する[^imp1]。

[^imp1]: ソースコードの先頭部分 `package main` に続くコメント `// import ...` を正しく書かないとビルド時にエラーになる。

```go
package main // import "github.com/spiegel-im-spiegel/hello"

import (
    "fmt"

    "rsc.io/quote"
)

func main() {
    fmt.Println(quote.Hello())
}
```

ファイル名は `hello.go` とする。
次に `hello.go` を置いているフォルダに空の `go.mod` ファイルを作る。

```text
$ echo>go.mod
```

これで準備完了。
それじゃあ，いきなり `vgo build` してみよう。

```text
$ vgo build
vgo: resolving import "rsc.io/quote"
vgo: finding rsc.io/quote (latest)
vgo: adding rsc.io/quote v1.5.2
vgo: finding rsc.io/quote v1.5.2
vgo: finding rsc.io/sampler v1.3.0
vgo: finding golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
vgo: downloading rsc.io/quote v1.5.2
vgo: downloading rsc.io/sampler v1.3.0
vgo: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
```

`v1.5.2` といったバージョンはパッケージのリポジトリのタグ情報から取得する。
バージョンを示すタグ情報がない場合は `v0.0.0-20170915032832-14c0d48ead0c` のような感じで仮バージョンが付与される。

では，作成した実行バイナリを動かしてみる（この記事では Windows 環境なので悪しからず）。

```text
$ hello.exe
Hello, world.

$ SET LANG=ja

$ hello.exe
こんにちは世界。
```

おおー。
ちゃんと動いてる。

ビルド後，空の `go.mod` ファイルに以下のように記述が加えられた。

```text
module "github.com/spiegel-im-spiegel/hello"

require "rsc.io/quote" v1.5.2
```

### vgo のパッケージ管理

ビルドの様子を見ればわかると思うが，依存関係を調べて各パッケージを全てダウンロードしている。
実は `golang.org/x/text` パッケージは `GOPATH` 配下にダウンロード済みだったのだが，これを使ってはいないようだ。

じゃあ，ダウンロードしたパッケージは何処にあるかというと `$GOPATH/src/v` フォルダ以下に展開されていた[^gpth]。

[^gpth]: このパスは正式版までに変更されると考えられる。 `go.mod` ファイルにパッケージへのフルパスが記述されるため，わざわざ `GOPATH` 配下にパス構成を統合する必要が無いからである。最終的には `GOPATH` の削除も視野に入れているかもしれない。

```text
$GOPATH/src/v
├─cache
│  ├─golang.org
│  │  └─x
│  │      └─text
│  │          └─@v
│  └─rsc.io
│      ├─quote
│      │  └─@v
│      └─sampler
│          └─@v
├─golang.org
│  └─x
│      └─text@v0.0.0-20170915032832-14c0d48ead0c
└─rsc.io
    ├─quote@v1.5.2
    └─sampler@v1.3.0
```

更に `quote@v1.5.2` フォルダと `sampler@v1.3.0` フォルダにもそれぞれ `go.mod` ファイルがあって，以下のような記述になっている。

```text
module "rsc.io/quote"

require "rsc.io/sampler" v1.3.0
```

```text
module "rsc.io/sampler"

require "golang.org/x/text" v0.0.0-20170915032832-14c0d48ead0c
```

このように `go.mod` ファイルの情報を元にして依存パッケージのバージョンを管理するわけだ。
ちなみに各パッケージの全てのバージョンを列挙するには以下のコマンドが使える。

```text
$ vgo list -t rsc.io/sampler
rsc.io/sampler
        v1.0.0
        v1.2.0
        v1.2.1
        v1.3.0
        v1.3.1
        v1.99.99
```

また `hello.go ` 以下の依存パッケージのバージョン情報は以下のコマンドで見ることができる。

```text
$ vgo list -m -u
MODULE                               VERSION                             LATEST
github.com/spiegel-im-spiegel/hello  -                                   -
golang.org/x/text                    v0.0.0-20170915032832-14c0d48ead0c  v0.0.0-20171214130843-f21a4dfb5e38
rsc.io/quote                         v1.5.2 (2018-02-15 00:44)           -
rsc.io/sampler                       v1.3.0 (2018-02-14 04:05)           v1.99.99 (2018-02-14 07:20)
```

[dep] みたいに依存関係を可視化できるといいんだけどねえ。


## [Semantic Versioning] と後方互換性

vgo が管理するバージョンは [Semantic Versioning] に従うことが期待されている。
また同じ import パスで取得するパッケージは後方互換性を持つことも期待されている。

たとえば， `my/thing` パッケージの v2 が後方互換性のない構成になっていた場合は `my/thing/v2` という感じに import パスを変えるわけだ[^gopkg1]。

[^gopkg1]: あるいは [gopkg.in](http://labix.org/gopkg.in "gopkg.in - Stable APIs for the Go language") のような API を使う手もある。

{{< fig-img src="https://research.swtch.com/impver.png" title="Semantic Import Versioning" link="https://research.swtch.com/vgo-import" >}}

現行の [Go 言語]コンパイラはパッケージのバージョンを意識していないが（バージョン管理は [dep] のような外部ツールが担っている）， vgo が正式に組み込まれればより厳密な（[Semantic Versioning] に基づいた）バージョン管理が要求されることになると思う。
なので，今からそれを意識した運用を考えておくべきかもしれない。

## とはいえ，まだ先の話

とはいえ，次の v1.11 が出るのは早くても半年後（2018年8月頃）だし，正式対応するという v.1.12 など鬼が笑う話である。
今後 [Semantic Versioning] は意識したほうがいいかもしれないが，当面は [dep] などを用いた運用ができていればいいと思う。

## ブックマーク

- [Go & Versioning(vgo)を読んで大きな変更が入ったなと思った - Qiita](https://qiita.com/lufia/items/67701e2f927c77a75d6e)
- [和訳: Go & Versioning - Qiita](https://qiita.com/nekketsuuu/items/36f00484ff7c30fd2007)
    - [和訳: Go += Package Versioning (Go & Versioning, Part 1) - Qiita](https://qiita.com/nekketsuuu/items/60634417e6279ccfd95b)
    - [和訳: A Tour of Versioned Go (vgo) (Go & Versioning, Part 2) - Qiita](https://qiita.com/nekketsuuu/items/589bc29f00b507492a96)
    - [和訳: Semantic Import Versioning (Go & Versioning, Part 3) - Qiita](https://qiita.com/nekketsuuu/items/2dcad7dde29171e1fe3f)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[glide]: https://github.com/Masterminds/glide "Masterminds/glide"
[dep]: https://golang.github.io/dep/ "dep · Dependency management for Go"
[A Tour of Versioned Go]: https://research.swtch.com/vgo-tour "research!rsc: A Tour of Versioned Go (vgo) (Go & Versioning, Part 2)"
[Semantic Versioning]: http://semver.org/ "Semantic Versioning 2.0.0 | Semantic Versioning"
