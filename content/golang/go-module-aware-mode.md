+++
title = "モジュール対応モードへの移行を検討する"
date = "2018-10-22T15:55:11+09:00"
update = "2018-10-23T13:48:39+09:00"
description = "みんな。なにはさておき go.mod ファイルを作成するんだ！"
image = "/images/attention/go-logo_blue.png"
tags  = [ "golang", "engineering", "module", "versioning" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = true
+++

[Go 言語]コンパイラの[バージョン 1.11]({{< relref "/release/2018/09/go-1_11-ise-released.md" >}} "Go 1.11 のリリースと「モジュール」機能の実験的サポート")から搭載された「モジュール対応モード」だが，少し試してみてそれなりに使えそうなので，この記事である程度まとめておくことにした。
これからも何かあればこの記事に加筆・修正していく予定である。

なお，モジュール対応モードは 1.11 時点で preliminary support に過ぎないため，以降のバージョンで大きな変更が行われるかもしれない。
したがってこの記事はバージョン 1.12 以降で大幅に書き替える可能性がる。

## 用語の整理

まず最初に用語の定義をしておく。

### GOPATH モードとモジュール対応モード

バージョン 1.11 以降からは [Go 言語]コンパイラは以下の2つのモードのどちらかで動作する[^cmd1]。

[^cmd1]: 「GOPATH モード」および「モジュール対応モード」の名称は “[Command go](https://golang.org/cmd/go/ "go - The Go Programming Language")” の記述から抜き出した。 [Go 言語]って用語の表記に微妙なブレがあるのがイマイチだよなぁ。

- **GOPATH モード (GOPATH mode)** : バージョン 1.10 までの動作モード。標準ライブラリを除く全てのパッケージの管理とビルドを `$GOPATH` 以下のディレクトリで行う。パッケージの管理はリポジトリの最新リビジョンのみが対象となる
- **モジュール対応モード (module-aware mode)** : 標準ライブラリを除く全てのパッケージをモジュールとして管理する。モジュールの管理とビルドは任意のディレクトリで可能で，モジュールはリポジトリのバージョンタグまたはリビジョン毎に管理される

### 「モジュール」とは

モジュール対応モードでは，標準ライブラリを除くパッケージを「モジュール（module）」として管理する。
パッケージが [git] 等のバージョン管理ツールで管理されている場合はバージョン毎に異なるモジュールと見なされる。
つまりモジュールの実体は「パッケージ＋バージョン」ということになる。
 
ただしコード上ではパッケージとモジュールの間に区別はなく，したがってソースコードを書き換える必要はない。
モジュールはソースコードではなく `go.mod` ファイルで管理される。

## 環境変数 $GO111MODULE

バージョン 1.11 では2つのモードの切り替えのために環境変数 `$GO111MODULE` でモードを指定する。
環境変数 `$GO111MODULE` は以下の値をとり得る。

| 値     | 内容                                                                                                     |
| ------ | -------------------------------------------------------------------------------------------------------- |
| `on`   | 常にモジュール対応モードで動作する                                                                       |
| `off`  | 常に GOPATH モードで動作する                                                                             |
| `auto` | `$GOPATH` 以下のディレクトリにあるパッケージは GOPATH モードで，それ以外はモジュール対応モードで動作する |

バージョン 1.11 では，環境変数 `$GO111MODULE` の既定値には `auto` が設定されている。
独自の開発環境や IDE を使用しているためモジュール対応モードへの移行が難しい場合には `$GO111MODULE` を `off` にするとよいだろう。

## モジュール対応モード移行への準備

本格的にモジュール対応モードへ移行する前に，以下の準備を行う。

### [Semantic Versioning] によるバージョン管理

モジュールのバージョンははリポジトリのリビジョン番号またはバージョンタグによって管理されるが，バージョンタグに関しては [Semantic Versioning] のルールに則ってバージョン番号を設定することが推奨されている[^sv1]。

[^sv1]: リビジョン番号によるバージョン管理も可能だが， `v0.0.0-20180816225734-aabede6cba87` のような擬似バージョン番号に置き換えられるため，モジュールのインポート時の運用が煩雑になり推奨できない。

{{< fig-img src="https://research.swtch.com/impver.png" title="Semantic Import Versioning" link="https://research.swtch.com/vgo-import" >}}

このように後方互換性のない変更がある場合にはメジャーバージョンを，後方互換性が担保された変更や追加についてはマイナーバージョンを，不具合や脆弱性の修正については第3位のパッチバージョンを上げるようにする。
またメジャーバージョンを上げる際には，図のようにディレクトリを分離するか，旧バージョン用にブランチを切るのがよいだろう[^sv2]。

[^sv2]: 異なるメジャーバージョンを同一ディレクトリの同一ブランチで管理していると `go.mod` のバージョン番号部分に勝手に “`+incompatible`” が付加されてめっさカッコ悪くなる（笑）

バージョンタグによってバージョン管理を行うのであればきちんとルールを決めて管理する必要がある。
リポジトリ管理に [Git] を使うのであれば Git Flow や GitHub Flow などが参考になるかも知れない。

- [A successful Git branching model » nvie.com](http://nvie.com/posts/a-successful-git-branching-model/)
- [git-flow cheatsheet](http://danielkummer.github.io/git-flow-cheatsheet/) （[日本語](http://danielkummer.github.io/git-flow-cheatsheet/index.ja_JP.html)）
- [GitHub Flow – Scott Chacon](http://scottchacon.com/2011/08/31/github-flow.html) （[日本語訳](https://gist.github.com/Gab-km/3705015)）

[GitHub] 等の SaaS ではバージョンタグを使ったリリース管理がやりやすいように構成されているし，上手く使ってほしい。

### なにはさておき go.mod ファイルを作成する

モジュール対応モードでは `go.mod` ファイルでモジュール管理を行う。
そういうわけで，なにはさておき `go.mod` ファイルを作ってしまおう。

適当なディレクトリで以下のコマンドを実行する。

```text
$ go mod init hello
go: creating new go.mod: module hello
```

ここで `hello` というのがモジュール名（厳密にはモジュール・パス）になる。
モジュール名は [Go 言語]の名前として使えるものであれば何でもいいが，パッケージ名（厳密にはパッケージ・パス）と合わせておくと面倒がない。
たとえば [GitHub] に `github.com/my/package` というパッケージを公開するのであれば，モジュール名も

```text
$ go mod init github.com/my/package
go: creating new go.mod: module github.com/my/package
```

としておくのがいいだろう。

コンパイル時に参照する `go.mod` ファイルがどこにあるかは環境変数 `$GOMOD` を見ればよい。

```text
$ go env GOMOD
```

カレント・ディレクトリに `go.mod` ファイルがない場合は親ディレクトリに遡って探すようだ。
まぁ，サブパッケージはディレクトリで階層化されているから当然なんだけど。

#### [dep] からの移行

既存のパッケージに `go.mod` ファイルを追加する場合も同様の操作でよい。
なお [dep] を使ってパッケージ管理を行っている場合は `Gopkg.lock` ファイルを読んで `go.mod` ファイルに組み込んでくれる。

```text
$ go mod init github.com/spiegel-im-spiegel/gpgpdump
go: creating new go.mod: module github.com/spiegel-im-spiegel/gpgpdump
go: copying requirements from Gopkg.lock
```

なんて便利！

#### go.mod の内容

ここで作成した `go.mod` ファイルの中身を見てみよう。

`github.com/spiegel-im-spiegel/gpgpdump` パッケージのリポジトリに対して `go mod init` した結果は以下の通りだった。

```text
module github.com/spiegel-im-spiegel/gpgpdump

require (
    github.com/BurntSushi/toml v0.3.0
    github.com/inconshreveable/mousetrap v1.0.0
    github.com/pkg/errors v0.8.0
    github.com/spf13/cobra v0.0.3
    github.com/spf13/pflag v1.0.2
    github.com/spiegel-im-spiegel/gocli v0.8.0
    golang.org/x/crypto v0.0.0-20180816225734-aabede6cba87
)
```

`module` や `require` は命令（directive）と呼ばれるもので，これらの命令を使ってモジュールの管理を行う。
`go.mod` ファイルの記述で使える命令は以下の通り。

| 命令      | 記述例                                          |
| --------- | ----------------------------------------------- |
| `module`  | `module my/thing`                               |
| `require` | `require other/thing v1.0.2`                    |
| `exclude` | `exclude old/thing v1.2.3`                      |
| `replace` | `replace bad/thing v1.4.5 => good/thing v1.4.5` |

`module` はカレント以下のディレクトリにあるパッケージに対するモジュール名を定義する。
前述したようにモジュール名（モジュール・パス）はパッケージ名（パッケージ・パス）と合わせておくほうが面倒がない。

`require` はカレントのモジュールが要求するモジュール名とバージョンを指定する。
つまり `require` で指定したモジュールが，カレント・モジュールが依存する外部モジュールとなるわけだ。

`exclude` は管理から除外するモジュールを指定する。
不具合等で特定のバージョンのモジュールを使いたくないときなどに使える。

`replace` はモジュール名の置き換えである。
パッケージ・パスのリダイレクトなどで名前と実体が一致しないときなどに使える。
たとえば以下のような感じ[^ltst1]。

[^ltst1]: バージョンに `latest` を指定すると，コンパイラ側で最新バージョン（またはリビジョン）を探して，最新の番号に置き換えてくれる。 `go build` や `go test` などを実行するたびに `go.mod` ファイルの `latest` 表記が書き換えられるので注意すること。

```text
module md2html

require gopkg.in/russross/blackfriday.v2 latest
replace gopkg.in/russross/blackfriday.v2 latest => github.com/russross/blackfriday/v2 latest
```

この件については後ほど詳しく説明する。

## モジュール対応モードで使う主なコマンド

それではモジュール対応モードで色々と操作してみよう。

`go build` や `go test` や `go run` のようなコマンドは，モジュール対応モードでも変わりなく使える。
ビルド時に必要なモジュールは自動的にダウンロードされるため，事前に `go get -u ./...` みたいなことはしなくていいようだ。

```text
$ go run hello.go
go: finding rsc.io/quote v1.5.2
go: downloading rsc.io/quote v1.5.2
go: finding rsc.io/sampler v1.3.0
go: finding golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: downloading rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
Hello, world.
```

なお `go get` の挙動については後ほど詳しく紹介する。

`go list` コマンドは `-m` オプションを付けることでモジュールに対応した。

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

実行結果から `github.com/BurntSushi/toml` と `golang.org/x/crypto` には新しいバージョンがあることが分かる。
この情報を基に `go.mod` を調整していけばいいだろう。

新しく追加されたモジュール対応モード用のコマンド `go mod` はこんな感じ。

```text
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

`go mod init` は既に紹介したとおり。
他によく使うものといえば `go mod download` と `go mod graph` だろうか。

`go mod download` は `go get -u ./...` の代わりに使える。
`go mod graph` はモジュール間の依存関係を調べるときに使えるだろう。

`go mod edit` は `go.mod` ファイルを編集するためのコマンドだが手で書いたほうが早い（笑） まぁ何らかのバッチ処理で使える感じだろうか。

## モジュール間の依存関係の構造化

以下のコードに対して

```go
package main

import (
    "fmt"

    "rsc.io/quote"
)

func main() {
    fmt.Println(quote.Hello())
}
```

`go.mod` ファイルの内容が以下の通りだとする。

```text
module hello

require rsc.io/quote v1.5.2
```

これに対して `go mod graph` コマンドを実行すると以下のようになる。

```text
$ go mod graph
hello rsc.io/quote@v1.5.2
rsc.io/quote@v1.5.2 rsc.io/sampler@v1.3.0
rsc.io/sampler@v1.3.0 golang.org/x/text@v0.0.0-20170915032832-14c0d48ead0c
```

これを図式化すると以下のようになるだろう。

{{< fig-mermaid >}}
graph TD
  hello["hello"]
  quote["rsc.io/quote@v1.5.2"]
  sampler["rsc.io/sampler@v1.3.0"]
  text["golang.org/x/text@v0.0.0-20170915032832-14c0d48ead0c"]

  hello--"hello/go.mod"-->quote
  quote--"rsc.io/quote@v1.5.2/go.mod"-->sampler
  sampler--"rsc.io/sampler@v1.3.0/go.mod"-->text
{{< /fig-mermaid >}}

`hello` モジュールは `go.mod` より `rsc.io/quote@v1.5.2` モジュールを参照している。
`rsc.io/quote@v1.5.2` モジュールは自身の `go.mod` より `rsc.io/sampler@v1.3.0` モジュールを参照している。
同じように `rsc.io/sampler@v1.3.0` モジュールからも依存モジュールを参照しているわけだ。

`hello/go.mod` には `rsc.io/sampler@v1.3.0` 等のモジュールについての記述はないが， `rsc.io/quote@v1.5.2/go.mod` や `rsc.io/sampler@v1.3.0/go.mod` の記述から依存関係を階層的に取得することができる。

このようにモジュールごとに `go.mod` ファイルを記述していくことで依存関係の構造化を実現できる。

- [research!rsc: Minimal Version Selection (Go & Versioning, Part 4)](https://research.swtch.com/vgo-mvs)

だからみんな。なにはさておき `go.mod` ファイルを作成するんだ！

## モジュールのキャッシュとビルド・キャッシュ

パッケージのコンパイル結果は環境変数 `$GOCACHE` の示すディレクトリ[^bc1] にキャッシュされる。
モジュール対応モードではモジュール単位でキャッシュされるようだ。
したがって `$GOCACHE` の値は `off` にせず，正しいディレクトリを指定する必要がある[^bc2]。

[^bc1]: Windows 環境では `$GOCACHE` の既定値は `%USERPROFILE%\AppData\Local\go-build` となっているようだ。
[^bc2]: バージョン 1.12 以降では `$GOCACHE` の値を `off` にできなくなる予定である。

`$GOCACHE` の値は `go env` コマンドで確認できる。

```text
$ go env GOCACHE
```

ビルド・キャッシュをクリアするには以下のコマンドを実行する。

```text
$ go clean -cache
```

バージョン 1.11 ではダウンロードしたモジュールのソースコードは `$GOPATH/pkg/mod` 以下に格納される。
モジュールのソースコードを含めて全てのキャッシュをクリアするには以下のコマンドを実行する。

```text
$ go clean -modcache
```

[Go 言語]コンパイラの将来バージョンでは `$GOPATH/pkg` を廃止する予定があるようで，モジュールのキャッシュの置き場が変更される可能性がある。

## モジュール対応モード in Travis CI

CI (Continuous Integration) サービスのひとつである [Travis CI] は [Go 言語]にも対応しているが，処理対象のパッケージを `$GOPATH/src` 以下に展開するため GOPATH モードで処理が走ってしまう。
そこで `.travis.yml` ファイルを以下のように書き換えて強引にモジュール対応モードにする。

{{< highlight yaml "hl_lines=6-8" >}}
language: go

go:
- "1.11.x"

env:
  global:
  - GO111MODULE=on

install:
- go mod download

script:
- go test ./...
{{< /highlight >}}

環境変数の設定部分に注目してほしい。

### [GoReleaser] による Deploy

[GoReleaser] を使って[複数の実行モジュールをビルド]({{< ref "/golang/cross-compiling-in-travis-ci-with-goreleaser.md" >}} "Travis CI でクロス・コンパイル（GoReleaser 編）")するには `.goreleaser.yml` ファイルの `builds` 項目に環境変数の設定を追加しておけば安心である。

```yaml
builds:
- env:
  - GO111MODULE=on
```

[Travis CI] で [GoReleaser] を使って deploy を行うには `.travis.yml` ファイルを以下のように記述すれば良い。

{{< highlight yaml "hl_lines=16-22" >}}
language: go

go:
- "1.11.x"

env:
  global:
  - GO111MODULE=on

install:
- go mod download

script:
- go test ./...

deploy:
- provider: script
  skip_cleanup: true
  script: curl -sL https://git.io/goreleaser | bash
  on:
    tags: true
    condition: $TRAVIS_OS_NAME = linux
{{< /highlight >}}

## モジュール・パスとパッケージ・パスが一致しない場合

先ほど少し紹介したが `gopkg.in/russross/blackfriday.v2` パッケージをモジュールとして記述する例を挙げてみる[^bf1]。

[^bf1]: `gopkg.in/russross/blackfriday.v2` パッケージは Markdown コードを HTML 等の書式に変換するパッケージである（参考：[Markdown パーサ blackfriday.v2 で遊ぶ]({{< ref "/golang/using-blackfriday-v2.md" >}})）。

まず以下のコードを書く。ファイル名は `md2html.go` とでもしておこう。

```go
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

```text
module md2html

require gopkg.in/russross/blackfriday.v2 latest
```

では実際にコンパイル&実行してみよう。

```text
$ go run md2html.go hello.md
go: finding gopkg.in/russross/blackfriday.v2 v2.0.1
go: gopkg.in/russross/blackfriday.v2@v2.0.1: go.mod has non-....v2 module path "github.com/russross/blackfriday/v2" at revision v2.0.1
go: error loading module requirements
```

おおう。
エラーになってしまった。
`github.com/russross/blackfriday/v2` なんてパスはねーよ！ とお怒りのようだ。

パス名 `gopkg.in/russross/blackfriday.v2` は `github.com/russross/blackfriday` の alias で v2 のコードは v2 ブランチで管理されている。
v2 ブランチにある `go.mod` ファイルを見てみると

```text
module github.com/russross/blackfriday/v2
```

と記されているが，実際に v2 ブランチの [`github.com/russross/blackfriday`](https://github.com/russross/blackfriday/tree/v2) に `v2` ディレクトリは存在しない。
このパスの不一致がエラーの原因のようである。

この不一致を解消するために `replace` を追加する。
具体的には以下のように記述する。

{{< highlight tex "hl_lines=4" >}}
module md2html

require gopkg.in/russross/blackfriday.v2 latest
replace gopkg.in/russross/blackfriday.v2 latest => github.com/russross/blackfriday/v2 latest
{{< /highlight >}}

では，この状態でコンパイル&実行してみよう。

```text
$ go run md2html.go hello.md
go: finding github.com/russross/blackfriday/v2 v2.0.1
go: downloading github.com/russross/blackfriday/v2 v2.0.1
go: finding github.com/shurcooL/sanitized_anchor_name latest
go: downloading github.com/shurcooL/sanitized_anchor_name v0.0.0-20170918181015-86672fcb3f95
<h1>Hello</h1>

<p>Hello, World!</p>
```

というわけでうまく動作したようだ。

## バージョンを指定して go get を実行する

{{% div-box %}}
【追記 2019-02-26】[Go 1.12]({{< ref "/release/2019/02/go-1_12-is-released.md" >}}) からはダミーの `go.mod` ファイルは必要なくなった。
任意の場所で `go get` コマンドを起動できる。
ブラボー！
{{% /div-box %}}

たとえば私の自作ツールである [gpgpdump] を自前でビルドすることを考える。

GOPATH モードでパッケージをダウンロードし実行ファイルをビルドするには以下のコマンドを実行すればよい。

```text
$ go get -v github.com/spiegel-im-spiegel/gpgpdump/cli/gpgpdump
github.com/spiegel-im-spiegel/gpgpdump (download)
github.com/spiegel-im-spiegel/gocli (download)
vendor/golang_org/x/net/dns/dnsmessage
github.com/spiegel-im-spiegel/gocli/exitcode
github.com/spiegel-im-spiegel/gocli/rwi
github.com/inconshreveable/mousetrap
net
github.com/spf13/pflag
github.com/spf13/cobra
github.com/spiegel-im-spiegel/gpgpdump/cli/gpgpdump/facade
github.com/spiegel-im-spiegel/gpgpdump/cli/gpgpdump
```

ダウンロードした `github.com/spiegel-im-spiegel/gpgpdump` パッケージのソースコードとコンパイル結果（実行ファイルを含む）は `$GOPATH` 以下に格納される。

では，モジュール対応モードでバージョンを指定して [gpgpdump] をビルドしてみよう。

```text
$ export GO111MODULE=on

$ go get github.com/spiegel-im-spiegel/gpgpdump/cli/gpgpdump@latest
go: finding github.com/spiegel-im-spiegel/gpgpdump/cli/gpgpdump latest
go: finding github.com/spiegel-im-spiegel/gpgpdump/cli latest
go: finding github.com/spiegel-im-spiegel/gpgpdump v0.3.9
go: downloading github.com/spiegel-im-spiegel/gpgpdump v0.3.9
go: extracting github.com/spiegel-im-spiegel/gpgpdump v0.3.9
go: finding github.com/pkg/errors v0.8.1
go: finding github.com/spf13/pflag v1.0.3
go: finding github.com/BurntSushi/toml v0.3.1
go: finding github.com/spiegel-im-spiegel/gocli v0.9.1
go: finding github.com/spf13/cobra v0.0.3
go: finding github.com/inconshreveable/mousetrap v1.0.0
go: finding golang.org/x/crypto v0.0.0-20190208162236-193df9c0f06f
go: finding github.com/mattn/go-isatty v0.0.4
go: downloading github.com/spiegel-im-spiegel/gocli v0.9.1
go: downloading github.com/spf13/cobra v0.0.3
go: downloading github.com/pkg/errors v0.8.1
go: downloading golang.org/x/crypto v0.0.0-20190208162236-193df9c0f06f
go: downloading github.com/BurntSushi/toml v0.3.1
go: extracting github.com/spiegel-im-spiegel/gocli v0.9.1
go: extracting github.com/pkg/errors v0.8.1
go: extracting github.com/BurntSushi/toml v0.3.1
go: extracting github.com/spf13/cobra v0.0.3
go: downloading github.com/spf13/pflag v1.0.3
go: downloading github.com/inconshreveable/mousetrap v1.0.0
go: extracting golang.org/x/crypto v0.0.0-20190208162236-193df9c0f06f
go: extracting github.com/inconshreveable/mousetrap v1.0.0
go: extracting github.com/spf13/pflag v1.0.3
```

おおっ，上手くいった（ちなみにモジュール名に意味はない）。
ソースコードは `$GOPATH/pkg/mod` 以下に，ビルド結果の実行ファイルは `$GOPATH/bin` に格納されるようだ。

ちなみに `path@version` 形式でモジュールを指定できるのは `go get` コマンドのみらしい。

```text
$ go run github.com/spiegel-im-spiegel/gpgpdump/cli/gpgpdump@latest
package github.com/spiegel-im-spiegel/gpgpdump/cli/gpgpdump@latest: can only use path@version syntax with 'go get'
```

ただし，バージョン指定を省略して latest バージョンで `go build` や `go run` コマンドを使うことはできるようだ。

```text
$ go clean -modcache

$ go run github.com/spiegel-im-spiegel/gpgpdump/cli/gpgpdump -h
go: finding github.com/spiegel-im-spiegel/gpgpdump/cli/gpgpdump latest
go: finding github.com/spiegel-im-spiegel/gpgpdump/cli latest
go: finding github.com/spiegel-im-spiegel/gpgpdump v0.3.9
go: downloading github.com/spiegel-im-spiegel/gpgpdump v0.3.9
go: extracting github.com/spiegel-im-spiegel/gpgpdump v0.3.9
go: finding github.com/spiegel-im-spiegel/gocli v0.9.1
go: finding github.com/spf13/cobra v0.0.3
go: finding github.com/spf13/pflag v1.0.3
go: finding github.com/BurntSushi/toml v0.3.1
go: finding github.com/pkg/errors v0.8.1
go: finding github.com/inconshreveable/mousetrap v1.0.0
go: finding golang.org/x/crypto v0.0.0-20190208162236-193df9c0f06f
go: finding github.com/mattn/go-isatty v0.0.4
go: downloading github.com/spiegel-im-spiegel/gocli v0.9.1
go: downloading github.com/pkg/errors v0.8.1
go: downloading golang.org/x/crypto v0.0.0-20190208162236-193df9c0f06f
go: downloading github.com/spf13/cobra v0.0.3
go: downloading github.com/BurntSushi/toml v0.3.1
go: extracting github.com/pkg/errors v0.8.1
go: extracting github.com/spiegel-im-spiegel/gocli v0.9.1
go: extracting github.com/BurntSushi/toml v0.3.1
go: extracting github.com/spf13/cobra v0.0.3
go: extracting golang.org/x/crypto v0.0.0-20190208162236-193df9c0f06f
go: downloading github.com/spf13/pflag v1.0.3
go: downloading github.com/inconshreveable/mousetrap v1.0.0
go: extracting github.com/inconshreveable/mousetrap v1.0.0
go: extracting github.com/spf13/pflag v1.0.3
Usage:
  gpgpdump [flags] [OpenPGP file]

Flags:
  -a, --armor     accepts ASCII input only
      --debug     for debug
  -h, --help      help for gpgpdump
  -i, --int       dumps multi-precision integers
  -j, --json      output with JSON format
  -l, --literal   dumps literal packets (tag 11)
  -m, --marker    dumps marker packets (tag 10)
  -p, --private   dumps private packets (tag 60-63)
  -t, --toml      output with TOML format
  -u, --utc       output with UTC time
  -v, --version   output version of gpgpdump
```

こりゃあ，便利だ。
上手く活用しましょう。

## ブックマーク

- [Go 1.11 のリリースと「モジュール」機能の実験的サポート]({{< relref "/release/2018/09/go-1_11-ise-released.md" >}})
- [GOPATH モードからモジュール対応モードへ移行せよ](https://qiita.com/spiegel-im-spiegel/items/5cb1587cb55d6f6a34d7)
- [Go 1.11 の modules・vgo を試す - 実際に使っていく上で考えないといけないこと #golang | Wantedly Engineer Blog](https://www.wantedly.com/companies/wantedly/post_articles/132270)

- [GOPATH 汚染問題]({{< ref "/golang/gopath-pollution.md" >}})
- [vgo (Versioned Go) に関する覚え書き]({{< ref "/golang/go-and-versioning.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[dep]: https://golang.github.io/dep/ "dep · Dependency management for Go"
[Semantic Versioning]: http://semver.org/ "Semantic Versioning 2.0.0 | Semantic Versioning"
[Travis CI]: https://travis-ci.org/ "Travis CI - Test and Deploy Your Code with Confidence"
[GoReleaser]: https://goreleaser.com/ "GoReleaser | Deliver Go binaries as fast and easily as possible"
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[Git]: https://git-scm.com/
[git]: https://git-scm.com/ "Git"
[GitHub]: https://github.com/

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253"><img src="https://images-fe.ssl-images-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
	<dd>Alan A.A. Donovan, Brian W. Kernighan</dd>
	<dd>柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>Book 単行本（ソフトカバー）</dd>
    <dd>ASIN: 4621300253, EAN: 9784621300251</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
