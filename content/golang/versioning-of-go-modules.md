+++
title = "Go モジュールのバージョン管理"
date =  "2019-05-04T13:08:19+09:00"
description = "試して壊して試して壊して... を繰り返した成果が今回の記事である。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "golang", "engineering", "module", "versioning" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今回の長期休暇を利用して今まで公開したツールやパッケージ類をチューニングしているのだが， [Go] 1.11 以降から実装されているモジュール対応モード（module-aware mode）のバージョン管理の挙動が（ドキュメントを読んだだけでは）ピンとこなかったので，この際いろいろと試してみることにした。

試して壊して試して壊して... を繰り返した成果が今回の記事である[^e1]。
まとめは[最後に書いておく]({{< relref "#digest" >}})のであしからず。

[^e1]: 実際にはこの記事で書いた量の三倍くらいは試して壊して... を繰り返している。

## みんな大好き Hello World

まずは以下の簡単なパッケージを作ってみる。

```text
hello/
├── go.mod
└── hello.go
```

`go.mod` ファイルの内容は以下の通り。
今回の記事では先頭行の `module` ディレクティブに注目する。
`module` ディレクティブはパッケージのモジュール・パスを定義するもので，このモジュールパスとバージョンのセットがモジュールの IDentity となる。

```text
module github.com/spiegel-im-spiegel/hello

go 1.12
```

`hello.go` ファイルの内容は以下の通り。

```go
package hello

import "fmt"

func Hello() {
	fmt.Println("Hello World")
}
```

このパッケージをリポジトリに push してバージョンタグ `v1.0.0` を付ける。

パッケージを使う側のコードも書いておこう。

```go
package main

import "github.com/spiegel-im-spiegel/hello"

func main() {
	hello.Hello()
}
```

これを実行すると以下のようになる。

```text
$ go run main.go 
go: finding github.com/spiegel-im-spiegel/hello v1.0.0
go: downloading github.com/spiegel-im-spiegel/hello v1.0.0
go: extracting github.com/spiegel-im-spiegel/hello v1.0.0
Hello World
```

このとき，パッケージを使う側の `go.mod` は以下のようになっているはずである（モジュール名は適当）。

```text
module work

go 1.12

require github.com/spiegel-im-spiegel/hello v1.0.0
```

前準備はこれで OK

## パッケージのバージョンを v2 にアップグレードする

ではこの `hello` パッケージを少し弄ってみよう。
まずは安直に `hello.go` 関数を以下のように変更する。

{{< highlight go "hl_lines=5-7" >}}
package hello

import "fmt"

func Hello(name string) {
	fmt.Println("Hello", name, "by v2")
}
{{< /highlight >}}

`Hello()` 関数の後方互換性が失われたのでメジャーバージョンを上げることにしよう。
このコードを push してバージョンタグ `v2.0.0` を付ける。

この新しいパッケージで使う側のコードを修正してみる。

{{< highlight go "hl_lines=6" >}}
package main

import "github.com/spiegel-im-spiegel/hello"

func main() {
	hello.Hello("Golang")
}
{{< /highlight >}}

`go.mod` ファイルも直さないとね。

{{< highlight text "hl_lines=5" >}}
module work

go 1.12

require github.com/spiegel-im-spiegel/hello v2.0.0
{{< /highlight >}}

これを実行すると以下のようになる。

```text
$ go run main.go 
go: finding github.com/spiegel-im-spiegel/hello v2.0.0
go: downloading github.com/spiegel-im-spiegel/hello v0.0.0-20190503134808-f31e6a72de0f
go: extracting github.com/spiegel-im-spiegel/hello v0.0.0-20190503134808-f31e6a72de0f
Hello Golang by v2
```

ありゃりゃ。
`v2.0.0` のモジュールを見つけたまではよかったが，ダウンロード時にバージョンタグを認識していない？

ここで思い出したのが [Semantic Versioning] のルールである。

{{< fig-img src="https://research.swtch.com/impver.png" title="Semantic Import Versioning" link="https://research.swtch.com/vgo-import" >}}

ひょっとして `v2` ディレクトリを切ったらいいのか？ 試してみよう[^sv1]。

[^sv1]: ちなみに `v0` から `v1` へのアップグレード時にはこのようなことは起きない。一般的に `v0` 系はベータ版と認識されていて後方互換性については煩くない。 [Go 言語]のモジュール対応モードでもチェックが入らないようだ。言い方を変えると `v1` 以降は（[Semantic Versioning] に従うなら）後方互換性についてちゃんと考えないといけないってこともであるのだが。バージョン設計と運用は意外と難しい？

## v2 ディレクトリによる分離

先ほどのコミットはなかったことにして， `hello` パッケージの構成を以下のように変える。

```text
hello/
├── go.mod
├── hello.go
└── v2/
    └── hello.go
```

`hello.go` が `v1` のコードで `v2/hello.go` が `v2` のコードである。

このパッケージを使う側のコードも以下のように変える。

{{< highlight go "hl_lines=3" >}}
package main

import "github.com/spiegel-im-spiegel/hello/v2"

func main() {
	hello.Hello("Golang")
}
{{< /highlight >}}

`go.mod` はこんな感じ？

{{< highlight text "hl_lines=5" >}}
module work

go 1.12

require github.com/spiegel-im-spiegel/hello/v2 v2.0.0
{{< /highlight >}}

これで実行してみよう。

```text
$ go run main.go 
go: finding github.com/spiegel-im-spiegel/hello/v2 v2.0.0
go: github.com/spiegel-im-spiegel/hello/v2@v2.0.0: go.mod has non-.../v2 module path "github.com/spiegel-im-spiegel/hello" (and .../v2/go.mod does not exist) at revision v2.0.0
go: error loading module requirements
```

ええつと？ あぁ，そうか。
パッケージ側の`v2/` ディレクトリにも `go.mod` ファイルがいるのか。

んじゃあ，以下の内容の `v2/go.mod` ファイルを追加して `v2.0.1` タグを付ける。

```text
module github.com/spiegel-im-spiegel/hello/v2

go 1.12
```

これでパッケージの構成は以下のようになった。

```text
hello/
├── go.mod
├── hello.go
└── v2/
    ├── go.mod
    └── hello.go
```

では，このパッケージを使って先ほどのコードを動かしてみよう。

```text
$ go run main.go 
go: finding github.com/spiegel-im-spiegel/hello/v2 v2.0.1
go: downloading github.com/spiegel-im-spiegel/hello/v2 v2.0.1
go: extracting github.com/spiegel-im-spiegel/hello/v2 v2.0.1
Hello Golang by v2
```

ようやく動いたよ... `orz`

## インポートパスをリダイレクトしたかったのだが...

パッケージ側の構成はこれでいいとして，パッケージをインポートする側は

```go
import "github.com/spiegel-im-spiegel/hello"
```

で `v2` のコードを動かしたいよね。
というわけで `go.mod` を以下のように書いてみる。

```text
module work

go 1.12

require github.com/spiegel-im-spiegel/hello/v2 v2.0.1

replace github.com/spiegel-im-spiegel/hello v2.0.1 => github.com/spiegel-im-spiegel/hello/v2 v2.0.1
```

これで動かすとどうなるか。

```text
$ go run main.go 
go: finding github.com/spiegel-im-spiegel/hello v2.0.1
go: finding github.com/spiegel-im-spiegel/hello/v2 v2.0.1
go: downloading github.com/spiegel-im-spiegel/hello/v2 v2.0.1
go: extracting github.com/spiegel-im-spiegel/hello/v2 v2.0.1
Hello Golang by v2
```

おっ，うまくいったっぽい？ でも `go.mod` ファイルを見てみると

```text
module work

go 1.12

require (
	github.com/spiegel-im-spiegel/hello v0.0.0-20190503144136-a8f02ef988d2 // indirect
	github.com/spiegel-im-spiegel/hello/v2 v2.0.1
)

replace github.com/spiegel-im-spiegel/hello v0.0.0-20190503144136-a8f02ef988d2 => github.com/spiegel-im-spiegel/hello/v2 v2.0.1
```

てな感じに書き換えられてしまった。
ふむむむむ？

どうもパッケージ内のディレクトリ名とバージョンタグを暗黙的に関連付けているようだ。
なので `v2.x` タグは `hello/v2/` ディレクトリに関連付けられてしまう。

{{< fig-img src="./hello-graph.png" link="./hello-graph.dot" >}}

たとえば同じリビジョンに `v1.0.1` タグを付ければ

{{< fig-img src="./hello-graph2.png" link="./hello-graph2.dot" >}}

という感じで `hello/` ディレクトリにもバージョンタグが割り当てられる。
もっともそれで

```text
module work

go 1.12

require (
	github.com/spiegel-im-spiegel/hello v1.0.1
	github.com/spiegel-im-spiegel/hello/v2 v2.0.1
)

replace github.com/spiegel-im-spiegel/hello v1.0.1 => github.com/spiegel-im-spiegel/hello/v2 v2.0.1
```

としたところで更なる混乱を招くだけだけどね。

### “Malformed Module Path”

ならば，旧い `v1` の方を別ディレクトリに移動すればいんじゃね？ って思うよね。
私は思った。

で，パッケージ側を

```text
hello/
├── go.mod
├── hello.go
└── v1/
    ├── go.mod
    └── hello.go
```

という構成にし，呼び出す側の `go.mod` ファイルを

```text
module work

go 1.12

require github.com/spiegel-im-spiegel/hello v1.0.1

replace github.com/spiegel-im-spiegel/hello v1.0.1 => github.com/spiegel-im-spiegel/hello/v1 v1.0.1
```

とかやってみたんだけど

```text
invalid module version github.com/spiegel-im-spiegel/hello/v1: malformed module path: github.com/spiegel-im-spiegel/hello/v1
```

とか言われたですよ。
いや “malformed module path” て `orz`

結局 *モジュール対応モード下でメジャー・バージョンを上げたならモジュール・パスも変えるしかない* ということらしい。

## v2 ブランチを切って運用する

とはいえバージョンごとに物理的にディレクトリを切って運用するというのは今時ありえないダサさである。
そこで物理的にディレクトリを切るのではなくリポジトリ上でブランチを切って運用することを考える。

パッケージのディレクトリ構成は `v1` と同じ。

```text
hello/
├── go.mod
└── hello.go
```

これに対して `v2` ブランチを切り， `v2` ブランチ上で `go.mod` を以下のように変更する。

{{< highlight text "hl_lines=1" >}}
module github.com/spiegel-im-spiegel/hello2/v2

go 1.12
{{< /highlight >}}

*モジュールのパスと物理パスが異なっている* が気にしないで先に進む。
`hello.go` を

```go
package hello

import "fmt"

func Hello(name string) {
	fmt.Println("Hello", name, "by v2")
}
```

として `go.mod` とともに `v2` ブランチに commit & push し，バージョンタグ `v2.0.0` を付与する。

パッケージを使用する側のコードは以下の通り。

```go
package main

import "github.com/spiegel-im-spiegel/hello/v2"

func main() {
	hello2.Hello("Golang")
}
```

これを実行すると

```text
$ go run main.go 
go: finding github.com/spiegel-im-spiegel/hello/v2 v2.0.0
go: downloading github.com/spiegel-im-spiegel/hello/v2 v2.0.0
go: extracting github.com/spiegel-im-spiegel/hello/v2 v2.0.0
Hello Golang by v2
```

という感じでうまく動いたようだ。
`go.mod` の内容も

```text
module work

go 1.12

require github.com/spiegel-im-spiegel/hello/v2 v2.0.0 // indirect
```

となっていた。
よーし，うむうむ，よーし。

ブランチとモジュール・パスの関係は以下のような感じだろうか。

{{< fig-img src="./hello-graph3.png" link="./hello-graph3.dot" >}}

## まとめると...{#digest}

1. `v1` 以降，メジャーバージョンを上げる度にモジュール・パスを変更して管理を分ける
    - `v2.x` なら `path/to/module/v2` などとする。最後の `v2` がポイント
    - パスの最後がバージョン番号（`v2` など）になっていれば，暗黙的にバージョンタグが対応する
2. モジュール・パスを変更するには `go.mod` ファイルの `module` ディレクティブを変更する
    - 物理的にディレクトリを切るのであれば `go.mod` ファイルも含める
    - バージョンごとにブランチを切って管理するのであれば，各ブランチの `go.mod` ファイルで指定するモジュール・パスに注意する
3. パッケージを利用する側はリポジトリの物理パスとモジュール・パスが異なる場合があるため `go.mod` ファイルに記述されているモジュール・パスを確認する
4. 同一パッケージの異なるメジャー・バージョンのモジュール・パスを `replace` で繋がないこと。更に分かりにくくなるか指定によってはエラーになる

といったところだろうか。

バージョンごとにパッケージのパスを分けるため [gopkg.in] といったサービスが使われることがあるが，リポジトリの物理パスとモジュール・パスが異なる場合は注意が必要である。
うまくパッケージをダウンロードできない場合は `go.mod` ファイル内に

```text
replace gopkg.in/russross/blackfriday.v2 v2.0.1 => github.com/russross/blackfriday/v2 v2.0.1
```

といった記述が必要になるかもしれない（というかそれが元々の `replace` ディレクティブの機能）。

## ブックマーク

- [モジュール対応モードへの移行を検討する]({{< relref "./go-module-aware-mode.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[Go]: https://golang.org/ "The Go Programming Language"
[Semantic Versioning]: http://semver.org/ "Semantic Versioning 2.0.0 | Semantic Versioning"
[gopkg.in]: https://labix.org/gopkg.in "gopkg.in - Stable APIs for the Go language"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
