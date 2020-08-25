# GOPATH モードからモジュール対応モードへ移行せよ

[Go 言語]バージョン 1.11 からパッケージとそのバージョン情報を一括管理できる「モジュール対応モード（module-aware mode）」が組み込まれた。 1.11 では preliminary support だが，2019年2月頃にリリースされる 1.12 以降で正式なサポートになる予定である（予定は未定）。

モジュール対応モードの利点は以下の2つであろう。

1. パッケージとそのバージョン情報を一括管理できる
2. 「[GOPATH 汚染問題](https://text.baldanders.info/golang/gopath-pollution/ "GOPATH 汚染問題 — プログラミング言語 Go | text.Baldanders.info")」の解消

特に2番目が重要だろう。私達はようやく GOPATH の呪いから解放されるのである（笑）

1番目については [dep] などの外部ツールでも対応可能だが，やはり公式にサポートされたというのは大きい。

モジュール対応モードでの各種操作は主に `go mod` コマンドで実装される。

```
$ go help mod
Go mod provides access to operations on modules.

Note that support for modules is built into all the go commands,
not just 'go mod'. For example, day-to-day adding, removing, upgrading,
and downgrading of dependencies should be done using 'go get'.
See 'go help modules' for an overview of module functionality.

Usage:

        go mod <command> [arguments]

The commands are:

        download    download modules to local cache
        edit        edit go.mod from tools or scripts
        graph       print module requirement graph
        init        initialize new module in current directory
        tidy        add missing and remove unused modules
        vendor      make vendored copy of dependencies
        verify      verify dependencies have expected content
        why         explain why packages or modules are needed

Use "go help mod <command>" for more information about a command.
```

以降から，簡単な例を交えて，モジュール対応モードについて解説してみる。

## モジュール対応モード

“[A Tour of Versioned Go (vgo)](https://research.swtch.com/vgo-tour)” を参考に以下のコードを書いてみる。ファイル名は `hello.go` とでもしておこう。

```go:hello.go
package main

import (
    "fmt"

    "rsc.io/quote"
)

func main() {
    fmt.Println(quote.Hello())
}
```

このファイルを任意の場所に置く（ココ重要）。もし外部パッケージ `rsc.io/quote` が GOPATH 配下にインストールされていなければコンパイル時に “`cannot find package "rsc.io/quote"`” と怒られエラーになる。ここまでは従来どおりの動作。

ここで `hello.go` のコードを `hello` モジュールとして定義する。モジュール定義は，ソースコードを変更する必要はなく， `go.mod` ファイルを新たに作成する。

`hello.go` のあるフォルダ上で以下のコマンドを実行する。

```
$ go mod init hello
go: creating new go.mod: module hello
```

このコマンドで作成した `go.mod` ファイルの中身を見てみよう。中身はテキストで1行だけ

```text:go.mod
module hello
```

と記されている。これで `hello.go` のコードは `hello` モジュールとして定義された。

この状態でコンパイル&実行してみよう。

```
$ go run hello.go
go: finding rsc.io/quote v1.5.2
go: downloading rsc.io/quote v1.5.2
go: finding rsc.io/sampler v1.3.0
go: finding golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: downloading rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
Hello, world.
```

`rsc.io/quote` パッケージとこのパッケージが依存する他の外部パッケージがダウンロードされ，コンパイル&実行が成功しているのが分かると思う。なお，コンパイル時に `go.mod` ファイルが以下の内容に書き換えられる。

```text:go.mod
module hello

require rsc.io/quote v1.5.2 // indirect
```

依存関係を見てみると

```
$ go mod graph
hello rsc.io/quote@v1.5.2
rsc.io/quote@v1.5.2 rsc.io/sampler@v1.3.0
rsc.io/sampler@v1.3.0 golang.org/x/text@v0.0.0-20170915032832-14c0d48ead0c
```

hello → rsc.io/quote → rsc.io/sampler → golang.org/x/text と依存関係を掘り下げていっているのが分かるだろう。

各パッケージの実体は `$GOPATH/mod` フォルダ以下に格納される。たとえば `rsc.io/quote` パッケージの v1.5.2 のコードは `$GOPATH/mod/rsc.io/quote@v1.5.2` に格納されている。

`rsc.io/quote@v1.5.2` フォルダを見おると `go.mod` ファイルがあるのが分かる。中身を見てみると

```text:go.mod
module "rsc.io/quote"

require "rsc.io/sampler" v1.3.0
```

と記されている。さらに `rsc.io/sampler@v1.3.0` フォルダの `go.mod` ファイルを見ると

```text:go.mod
module "rsc.io/sampler"

require "golang.org/x/text" v0.0.0-20170915032832-14c0d48ead0c
```

となっている。

このようにパッケージごとに `go.mod` ファイルを記述することで（バージョンを含めた）依存関係を階層的に構築できるようになる。

やったね！

### dep からの移行

[dep] を使って外部パッケージを管理している場合は `go mod init` コマンド実行時に `Gopkg.lock` ファイルの内容を自動的にインポートする。

```
$ go mod init github.com/spiegel-im-spiegel/gpgpdump
go: creating new go.mod: module github.com/spiegel-im-spiegel/gpgpdump
go: copying requirements from Gopkg.lock
```

外部パッケージが更新されているかどうか確認するには `go list -m -u all` コマンドを実行する。

```text
$ go list -m -u all
go: finding github.com/pkg/errors v0.8.0
go: finding github.com/spf13/pflag v1.0.2
go: finding github.com/spf13/cobra v0.0.3
go: finding github.com/BurntSushi/toml v0.3.0
go: finding github.com/spiegel-im-spiegel/gocli v0.8.0
go: finding github.com/inconshreveable/mousetrap v1.0.0
go: finding golang.org/x/crypto v0.0.0-20180816225734-aabede6cba87
go: finding golang.org/x/crypto latest
go: finding github.com/BurntSushi/toml v0.3.1
github.com/spiegel-im-spiegel/gpgpdump
github.com/BurntSushi/toml v0.3.0 [v0.3.1]
github.com/inconshreveable/mousetrap v1.0.0
github.com/pkg/errors v0.8.0
github.com/spf13/cobra v0.0.3
github.com/spf13/pflag v1.0.2
github.com/spiegel-im-spiegel/gocli v0.8.0
golang.org/x/crypto v0.0.0-20180816225734-aabede6cba87 [v0.0.0-20180910181607-0e37d006457b]
```

実行結果を見ると `github.com/BurntSushi/toml` と `golang.org/x/crypto` の各パッケージが更新されていることが分かる。

## Semantic Versioning

モジュール対応モードのバージョン管理は [Semantic Versioning] をベースに以下のような考え方になっているようだ。

{{< fig-img-quote src="https://research.swtch.com/impver.png" title="Semantic Import Versioning" link="https://research.swtch.com/vgo-import" lang="en" class="lightmode" >}}

なお `go.mod` ファイルを含まない外部パッケージで v2 以降のバージョンタグが付いているものについては，自パッケージの `go.mod` ファイルに

```text:go.mod
require github.com/lestrrat-go/file-rotatelogs v2.2.0+incompatible
```

のように “`+incompatible`” が付加される。ただし上の図のようにディレクトリを分けたりメジャーバージョンごとにブランチを切るなどすれば問題ないようだ。

リポジトリにバージョン・タグが切られてない場合は

```text:go.mod
require golang.org/x/sys v0.0.0-20181004145325-8469e314837c
```

のように日付とリビジョンを含む仮のバージョン番号を使う。とにかく最新版を指定するのであれば

```text:go.mod
require golang.org/x/sys latest
```

のように `latest` を指定する。`latest` 記述はコンパイル等を行う際に実際のバージョン番号に書き換えられる。

## モジュール・パスとパッケージ・パス

`go.mod` ファイルで定義するモジュール・パスとパッケージ・パスは合わせておくほうがよい。たとえば，先程の `hello.go` をリポジトリ `https://github.com/you/hello` に置くのであれば，モジュール・パスを

```text:go.mod
module github.com/you/hello
```

としておけばコンパイラが自動的にリポジトリからダウンロードしてくれる。

### gopkg.in/russross/blackfriday.v2 の場合

特殊な例だが `gopkg.in/russross/blackfriday.v2` パッケージを例にモジュール・パスとパッケージ・パスが一致しない場合の対応を紹介する。

（ちなみに `gopkg.in/russross/blackfriday.v2` パッケージは Markdown コードを HTML 等の書式に変換するパッケージである。参考：[Markdown パーサ blackfriday.v2 で遊ぶ](https://text.baldanders.info/golang/using-blackfriday-v2/ "Markdown パーサ blackfriday.v2 で遊ぶ — プログラミング言語 Go | text.Baldanders.info")）

まず以下のコードを書く。ファイル名は `md2html.go` とでもしておこう。

```go:md2html.go
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/russross/blackfriday.v2"
)

func main() {
	md, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	html := blackfriday.Run(md)
	fmt.Println(string(html))
}
```

`md2html.go` に対して以下のようにモジュール定義を行う。

```text:go.mod
module md2html

require gopkg.in/russross/blackfriday.v2 latest
```

では実際にコンパイル&実行してみよう。

```
$ go run md2html.go hello.md
go: finding gopkg.in/russross/blackfriday.v2 v2.0.1
go: gopkg.in/russross/blackfriday.v2@v2.0.1: go.mod has non-....v2 module path "github.com/russross/blackfriday/v2" at revision v2.0.1
go: error loading module requirements
```

おおう。エラーになってしまった。 `github.com/russross/blackfriday/v2` なんてパスはねーよ！ とお怒りのようだ。

パス名 `gopkg.in/russross/blackfriday.v2` は `github.com/russross/blackfriday` の alias で v2 のコードは `v2` ブランチ上で管理されている。 `v2` ブランチ上にある `go.mod` ファイルを見てみると

```text:go.mod
module github.com/russross/blackfriday/v2
```

と記されているが，実際に v2 ブランチの [`github.com/russross/blackfriday`](https://github.com/russross/blackfriday/tree/v2) に `v2` ディレクトリは存在しない。このパスの不一致がエラーの原因のようである。

この不一致を解消するためには `replace` directive を追加する。具体的には以下のように記述する。

```text:go.mod
module md2html

require gopkg.in/russross/blackfriday.v2 latest

replace gopkg.in/russross/blackfriday.v2 latest => github.com/russross/blackfriday/v2 latest
```

では，もう一度コンパイル&実行してみよう。

```
$ go run md2html.go hello.md
go: finding github.com/russross/blackfriday/v2 v2.0.1
go: downloading github.com/russross/blackfriday/v2 v2.0.1
go: finding github.com/shurcooL/sanitized_anchor_name latest
go: downloading github.com/shurcooL/sanitized_anchor_name v0.0.0-20170918181015-86672fcb3f95
<h1>Hello</h1>

<p>Hello, World!</p>
```

というわけでうまく動作したようだ。ちなみに書き換えられた `go.mod` ファイルの内容は以下の通り。

```text:go.mod
module md2html

require (
	github.com/shurcooL/sanitized_anchor_name v0.0.0-20170918181015-86672fcb3f95 // indirect
	gopkg.in/russross/blackfriday.v2 v2.0.1
)

replace gopkg.in/russross/blackfriday.v2 v2.0.1 => github.com/russross/blackfriday/v2 v2.0.1
```

`github.com/shurcooL/sanitized_anchor_name` パッケージは `gopkg.in/russross/blackfriday.v2` 内で使用されるのだが， `gopkg.in/russross/blackfriday.v2` の `go.mod` ファイルに依存関係の記述がないため `md2html` 側に記述されてしまった。残念です！

## GO111MODULE 環境変数

バージョン 1.11 でモジュール対応モードを有効にするためには環境変数 `GO111MODULE` を指定する。環境変数 `GO111MODULE` には以下の3つの値を指定できる。

- **on** : コードの置かれている場所にかかわらず常にモジュール対応モードで動作する
- **off** : モジュール対応モードを無効とし従来の GOPATH モード（GOPATH mode）で動作する（バージョン 1.10 までの挙動と同等）
- **auto** : `$GOPATH/src` 以下にあるコードについては従来どおり GOPATH モードで，それ以外の場所にあるコードについてはモジュール対応モードで動作する

既定では `auto` がセットされている。

## モジュール対応モード with Travis CI

CI (Continuous Integration) サービスのひとつである [Travis CI] は [Go 言語]にも対応しているが，処理対象のパッケージを `$GOPATH/src` 以下に展開するため GOPATH モードで処理が走ってしまう。そこで .travis.yml を以下のように書き換えて強引にモジュール対応モードにする。

```yaml:.travis.yml
language: go

go:
  - "1.11.x"

install:
  - env GO111MODULE=on go mod download

script:
  - env GO111MODULE=on go test ./...

after_success:
  - test -n "$TRAVIS_TAG" && curl -sL https://git.io/goreleaser | bash
```

最後の行のように [GoReleaser] を使って[複数の実行モジュールをビルド](https://text.baldanders.info/golang/cross-compiling-in-travis-ci-with-goreleaser/ "Travis CI でクロス・コンパイル（GoReleaser 編） — プログラミング言語 Go | text.Baldanders.info")する場合には `.goreleaser.yml` ファイルの `builds` 項目に以下の記述を追加する。

```yaml:.goreleaser.yml
builds:
- env:
  - GO111MODULE=on
```

まぁ，こんな感じかな。

## ブックマーク

- [go - The Go Programming Language](https://golang.org/cmd/go/)
- [Go 1.11 リリースノート（和訳） - Qiita](https://qiita.com/pokeh/items/c6511ca15c9a33b48fcc)
- [Go 1.11 の modules・vgo を試す - 実際に使っていく上で考えないといけないこと #golang | Wantedly Engineer Blog](https://www.wantedly.com/companies/wantedly/post_articles/132270)
- [Go 1.11 のリリースと「モジュール」機能の実験的サポート — リリース情報 | text.Baldanders.info](https://text.baldanders.info/release/2018/09/go-1_11-ise-released/)
- [Using Go modules with Travis CI | Dave Cheney](https://dave.cheney.net/2018/07/16/using-go-modules-with-travis-ci)


[Go 言語]: https://golang.org/ "The Go Programming Language"
[dep]: https://golang.github.io/dep/ "dep · Dependency management for Go"
[Semantic Versioning]: http://semver.org/ "Semantic Versioning 2.0.0 | Semantic Versioning"
[Travis CI]: https://travis-ci.org/ "Travis CI - Test and Deploy Your Code with Confidence"
[GoReleaser]: https://goreleaser.com/ "GoReleaser | Deliver Go binaries as fast and easily as possible"
