+++
title = "go get コマンドでパッケージを管理する"
date = "2015-09-13T22:16:34+09:00"
description = "今回は Go 言語のパッケージ管理について。"
tags = ["golang", "package", "engineering" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

（初出： [はじめての Go 言語 (on Windows) その3 - Qiita](http://qiita.com/spiegel-im-spiegel/items/a52a47942fd3946bb583)）

{{< div-box type="markdown" >}}
【追記 2019-12-20】 バージョン 1.11 からパッケージの管理は「モジュール」をベースにしたものに移行しつつある。
詳しくは以下を参照のこと。

- [モジュール対応モードへの移行を検討する]({{< relref "./go-module-aware-mode.md" >}})
- [Go モジュールのバージョン管理]({{< relref "./versioning-of-go-modules.md" >}})
- [パッケージの管理（モジュール対応版）]({{< relref "./manage-packages.md" >}})
{{< /div-box >}}

## Go のパッケージ{#package}

[前回]の最後に載せた hello.go をもう一度を挙げてみる。

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```

最初の宣言文 `package main` はこのファイルが `main` パッケージに属することを示している。
[Go 言語]では `main` パッケージにある `main` 関数がプログラム実行時の起点となる。
C 言語などと異なり， [Go 言語]では `main` 関数に渡す引数はなく  `main` 関数が返す返り値もない[^a]。

[^a]: コマンドライン引数を取得するには [`os`]`.Args` を，システムの `$?` や `%ERRORLEVEL%` に値を返す場合は [`os`]`.Exit()` 関数を使う。これらはプラットフォーム依存なので（そもそも組み込みなら引数も返り値もなくて当たり前），言語仕様に含めるのではなくパッケージとして実装しているようだ。ちなみに [`os`]`.Exit()` 関数はプロセスを強制終了してしまうため [Defer] 構文で予約された処理が起動しないので注意が必要。

`import "fmt"` はパッケージ [`fmt`] を呼び出すもので，ソースファイルの最初の方でまとめて呼び出す。
複数のパッケージを呼び出す場合は以下のように

```go
import (
    "flag"
    "fmt"
)
```

などと `( )` で囲む。
ちなみに，パッケージ [`flag`] はコマンドライン引数を処理するためのパッケージである。

パッケージの関数等を使う際は `fmt.Println` のようにパッケージ名を関数等の名前の前に付ける。

### gofmt{#gofmt}

余談だが [Go 言語]ではソースコードを整形するための [`gofmt`] コマンドが存在する。
このコマンドはコンパイラからも `go fmt` で呼び出すことができる。

```
C:>go help fmt
usage: go fmt [-n] [-x] [packages]

Fmt runs the command 'gofmt -l -w' on the packages named
by the import paths.  It prints the names of the files that are modified.

For more about gofmt, see 'go doc cmd/gofmt'.
For more about specifying packages, see 'go help packages'.

The -n flag prints commands that would be executed.
The -x flag prints commands as they are executed.

To run gofmt with specific options, run gofmt itself.

See also: go fix, go vet.
```

[Go 言語]は比較的冗長な表現を許容している。
こういうタイプの言語にはとっつきやすい利点はあるが，記述形式を巡って宗教論争が起きることも多い（あるいは[品質を落とすコーディングというのも存在する](https://baldanders.info/blog/000195/)）。
そこで `gofmt` コマンドを使ってある程度記述形式を統一することで，この手の混乱を避ける狙いがある。

## 外部パッケージ{#external}

パッケージの import は [`fmt`] のような標準ライブラリの他に任意のライブラリをパッケージとして含めることができる。
またインターネット上の GitHub などにある repository からパッケージを取得することもできる。
外部パッケージを取り込むために `go get` コマンドが用意されている。

```
C:>go help get
usage: go get [-d] [-f] [-fix] [-insecure] [-t] [-u] [build flags] [packages]

Get downloads and installs the packages named by the import paths,
along with their dependencies.

The -d flag instructs get to stop after downloading the packages; that is,
it instructs get not to install the packages.

The -f flag, valid only when -u is set, forces get -u not to verify that
each package has been checked out from the source control repository
implied by its import path. This can be useful if the source is a local fork
of the original.

The -fix flag instructs get to run the fix tool on the downloaded packages
before resolving dependencies or building the code.

The -insecure flag permits fetching from repositories and resolving
custom domains using insecure schemes such as HTTP. Use with caution.

The -t flag instructs get to also download the packages required to build
the tests for the specified packages.

The -u flag instructs get to use the network to update the named packages
and their dependencies.  By default, get uses the network to check out
missing packages but does not use it to look for updates to existing packages.

Get also accepts build flags to control the installation. See 'go help build'.

When checking out or updating a package, get looks for a branch or tag
that matches the locally installed version of Go. The most important
rule is that if the local installation is running version "go1", get
searches for a branch or tag named "go1". If no such version exists it
retrieves the most recent version of the package.

If the vendoring experiment is enabled (see 'go help gopath'),
then when go get checks out or updates a Git repository,
it also updates any git submodules referenced by the repository.

For more about specifying packages, see 'go help packages'.

For more about how 'go get' finds source code to
download, see 'go help importpath'.

See also: go build, go install, go clean.
```

たとえば [https://github.com/mitchellh/cli](https://github.com/mitchellh/cli) にあるパッケージを取得する場合には以下のようにする。

```
C:>go get -v github.com/mitchellh/cli
github.com/mitchellh/cli (download)
Fetching https://golang.org/x/crypto/ssh/terminal?go-get=1
Parsing meta tags from https://golang.org/x/crypto/ssh/terminal?go-get=1 (status code 200)
get "golang.org/x/crypto/ssh/terminal": found meta tag main.metaImport{Prefix:"golang.org/x/crypto", VCS:"git", RepoRoot:"https://go.googlesource.com/crypto"} at https://golang.org/x/crypto/ssh/terminal?go-get=1
get "golang.org/x/crypto/ssh/terminal": verifying non-authoritative meta tag
Fetching https://golang.org/x/crypto?go-get=1
Parsing meta tags from https://golang.org/x/crypto?go-get=1 (status code 200)
golang.org/x/crypto (download)
golang.org/x/crypto/ssh/terminal
github.com/mitchellh/cli
```

もし取得するパッケージが別のパッケージを呼び出している場合でも，依存関係ごとまとめて取得できる。
取得したパッケージをソースコード上で呼び出すには

```go
import "github.com/mitchellh/cli"
```

と記述する。
URL（の scheme を除いた部分）をそのまま記述するのがポイントである。

[Go 言語]では， `import` で指定するパッケージの path と repository の URL がそのまま連動していて， `go get` コマンドでパッケージ内の依存関係を解決する際にもパッケージの path を URL と解釈してパッケージを取得しようとする。

実際にコード上でパッケージを使用する際はパッケージの path のベース名がコード上のパッケージ名になる。
`github.com/mitchellh/cli` パッケージなら `cli.Run()` のように記述する。
このパッケージ名が被る場合などは

```go
import mcli "github.com/mitchellh/cli"
```

などとする。

## GOPATH 環境変数

`go get` コマンドでは，環境変数 `GOPATH` に取得したソースコードやコンパイル後のモジュールを格納する。
`GOPATH` 環境変数がないと以下のように怒られる。

```shell
C:>go get -v github.com/mitchellh/cli
package github.com/mitchellh/cli: cannot download, $GOPATH not set. For more details see: go help gopath
```

と怒られる。
`GOPATH` の詳しい解説は以下のとおり。

```
C:>go help gopath
The Go path is used to resolve import statements.
It is implemented by and documented in the go/build package.

The GOPATH environment variable lists places to look for Go code.
On Unix, the value is a colon-separated string.
On Windows, the value is a semicolon-separated string.
On Plan 9, the value is a list.

GOPATH must be set to get, build and install packages outside the
standard Go tree.

Each directory listed in GOPATH must have a prescribed structure:

The src directory holds source code.  The path below src
determines the import path or executable name.

The pkg directory holds installed package objects.
As in the Go tree, each target operating system and
architecture pair has its own subdirectory of pkg
(pkg/GOOS_GOARCH).

If DIR is a directory listed in the GOPATH, a package with
source in DIR/src/foo/bar can be imported as "foo/bar" and
has its compiled form installed to "DIR/pkg/GOOS_GOARCH/foo/bar.a".

The bin directory holds compiled commands.
Each command is named for its source directory, but only
the final element, not the entire path.  That is, the
command with source in DIR/src/foo/quux is installed into
DIR/bin/quux, not DIR/bin/foo/quux.  The "foo/" prefix is stripped
so that you can add DIR/bin to your PATH to get at the
installed commands.  If the GOBIN environment variable is
set, commands are installed to the directory it names instead
of DIR/bin.

Here's an example directory layout:

    GOPATH=/home/user/gocode

    /home/user/gocode/
        src/
            foo/
                bar/               (go code in package bar)
                    x.go
                quux/              (go code in package main)
                    y.go
        bin/
            quux                   (installed command)
        pkg/
            linux_amd64/
                foo/
                    bar.a          (installed package object)

Go searches each directory listed in GOPATH to find source code,
but new packages are always downloaded into the first directory
in the list.

See https://golang.org/doc/code.html for an example.

Internal Directories

Code in or below a directory named "internal" is importable only
by code in the directory tree rooted at the parent of "internal".
Here's an extended version of the directory layout above:

    /home/user/gocode/
        src/
            crash/
                bang/              (go code in package bang)
                    b.go
            foo/                   (go code in package foo)
                f.go
                bar/               (go code in package bar)
                    x.go
                internal/
                    baz/           (go code in package baz)
                        z.go
                quux/              (go code in package main)
                    y.go


The code in z.go is imported as "foo/internal/baz", but that
import statement can only appear in source files in the subtree
rooted at foo. The source files foo/f.go, foo/bar/x.go, and
foo/quux/y.go can all import "foo/internal/baz", but the source file
crash/bang/b.go cannot.

See https://golang.org/s/go14internal for details.

Vendor Directories

Go 1.5 includes experimental support for using local copies
of external dependencies to satisfy imports of those dependencies,
often referred to as vendoring. Setting the environment variable
GO15VENDOREXPERIMENT=1 enables that experimental support.

When the vendor experiment is enabled,
code below a directory named "vendor" is importable only
by code in the directory tree rooted at the parent of "vendor",
and only using an import path that omits the prefix up to and
including the vendor element.

Here's the example from the previous section,
but with the "internal" directory renamed to "vendor"
and a new foo/vendor/crash/bang directory added:

    /home/user/gocode/
        src/
            crash/
                bang/              (go code in package bang)
                    b.go
            foo/                   (go code in package foo)
                f.go
                bar/               (go code in package bar)
                    x.go
                vendor/
                    crash/
                        bang/      (go code in package bang)
                            b.go
                    baz/           (go code in package baz)
                        z.go
                quux/              (go code in package main)
                    y.go

The same visibility rules apply as for internal, but the code
in z.go is imported as "baz", not as "foo/vendor/baz".

Code in vendor directories deeper in the source tree shadows
code in higher directories. Within the subtree rooted at foo, an import
of "crash/bang" resolves to "foo/vendor/crash/bang", not the
top-level "crash/bang".

Code in vendor directories is not subject to import path
checking (see 'go help importpath').

When the vendor experiment is enabled, 'go get' checks out
submodules when checking out or updating a git repository
(see 'go help get').

The vendoring semantics are an experiment, and they may change
in future releases. Once settled, they will be on by default.

See https://golang.org/s/go15vendor for details.
```

（[Go 言語]のバージョン 1.5 から Internal Packages や Vendoring の機能が追加された。
これについては「[GOPATH 汚染問題]({{< relref "gopath-pollution.md" >}})」および「[パッケージ外部からの呼び出しを禁止する Internal Packages]({{< relref "internal-packages.md" >}})」で解説している）

また `go get` コマンドでは内部でソースコード管理ツールを呼び出す。
呼び出される可能性のあるツールは以下のとおり。

- svn : [Subversion]
- hg : [Mercurial]
- git : [Git]
- bzr : [Bazaar]

たとえば get コマンド実行時に

```shell
C:>go get -v github.com/spf13/hugo
go: missing Mercurial command. See http://golang.org/s/gogetcmd
package github.com/spf13/hugo/commands
        imports bitbucket.org/pkg/inflect: exec: "hg": executable file not found in %PATH%
```

などとエラーが出たらいずれかのバージョン管理ツール（この場合は [Mercurial]）がないことになる。
必要に応じてインストールしておくとよい。

### Proxy 設定

GitHub や Bitbucket 等のよく知られた repository から取得する場合は直に上記のコマンドを起動するが，それ以外のサイトでは HTTP/HTTPS でいったん fetch してから，どのコマンドで取得するか判断してるようだ。
このとき Firewall/Proxy で阻まれている環境ではエラーになってしまう。

この場合は Proxy の設定が必要。
設定には `http_proxy` および `https_proxy` 環境変数をセットする。たとえばこんな感じ。

```shell
SET http_proxy=http://username:password@proxy.exsample.com:8080/
SET https_proxy=%http_proxy%
```

## go get コマンドでツールをインストールする

`go get` ではライブラリだけではなく，ツールそのものもパッケージとしてダウンロード→インストールできる。

たとえば静的サイト・ジェネレータの [Hugo] [^b] は以下のように `go get` コマンド一発で依存関係ごとパッケージを取得しインストールまで自動で行う。

[^b]: [Hugo] については「[ゼロから始める Hugo](/hugo)」で解説している。以上，宣伝でした（笑）

```
C:\workspace>mkdir hugo

C:\workspace>cd hugo

C:\workspace\hugo>SET GOPATH=C:\workspace\hugo

C:\workspace\hugo>go get -v github.com/spf13/hugo
github.com/spf13/hugo (download)
github.com/kardianos/osext (download)
github.com/spf13/afero (download)
github.com/spf13/cast (download)
github.com/spf13/jwalterweatherman (download)
github.com/spf13/cobra (download)
github.com/cpuguy83/go-md2man (download)
github.com/russross/blackfriday (download)
github.com/shurcooL/sanitized_anchor_name (download)
github.com/inconshreveable/mousetrap (download)
github.com/spf13/pflag (download)
github.com/spf13/fsync (download)
github.com/PuerkitoBio/purell (download)
github.com/opennota/urlesc (download)
github.com/miekg/mmark (download)
github.com/BurntSushi/toml (download)
Fetching https://gopkg.in/yaml.v2?go-get=1
Parsing meta tags from https://gopkg.in/yaml.v2?go-get=1 (status code 200)
get "gopkg.in/yaml.v2": found meta tag main.metaImport{Prefix:"gopkg.in/yaml.v2", VCS:"git", RepoRoot:"https://gopkg.in/yaml.v2"} at https://gopkg.in/yaml.v2?go-get=1
gopkg.in/yaml.v2 (download)
github.com/spf13/viper (download)
github.com/kr/pretty (download)
github.com/kr/text (download)
github.com/magiconair/properties (download)
github.com/mitchellh/mapstructure (download)
Fetching https://golang.org/x/text/transform?go-get=1
Parsing meta tags from https://golang.org/x/text/transform?go-get=1 (status code 200)
get "golang.org/x/text/transform": found meta tag main.metaImport{Prefix:"golang.org/x/text", VCS:"git", RepoRoot:"https://go.googlesource.com/text"} at https://golang.org/x/text/transform?go-get=1
get "golang.org/x/text/transform": verifying non-authoritative meta tag
Fetching https://golang.org/x/text?go-get=1
Parsing meta tags from https://golang.org/x/text?go-get=1 (status code 200)
golang.org/x/text (download)
Fetching https://golang.org/x/text/unicode/norm?go-get=1
Parsing meta tags from https://golang.org/x/text/unicode/norm?go-get=1 (status code 200)
get "golang.org/x/text/unicode/norm": found meta tag main.metaImport{Prefix:"golang.org/x/text", VCS:"git", RepoRoot:"https://go.googlesource.com/text"} at https://golang.org/x/text/unicode/norm?go-get=1
get "golang.org/x/text/unicode/norm": verifying non-authoritative meta tag
bitbucket.org/pkg/inflect (download)
github.com/dchest/cssmin (download)
github.com/eknkc/amber (download)
github.com/yosssi/ace (download)
github.com/spf13/nitro (download)
github.com/gorilla/websocket (download)
Fetching https://gopkg.in/fsnotify.v1?go-get=1
Parsing meta tags from https://gopkg.in/fsnotify.v1?go-get=1 (status code 200)
get "gopkg.in/fsnotify.v1": found meta tag main.metaImport{Prefix:"gopkg.in/fsnotify.v1", VCS:"git", RepoRoot:"https://gopkg.in/fsnotify.v1"} at https://gopkg.in/fsnotify.v1?go-get=1
gopkg.in/fsnotify.v1 (download)
github.com/shurcooL/sanitized_anchor_name
github.com/spf13/afero
github.com/inconshreveable/mousetrap
github.com/spf13/hugo/bufferpool
github.com/kr/text
github.com/kardianos/osext
github.com/spf13/jwalterweatherman
github.com/spf13/pflag
github.com/russross/blackfriday
github.com/opennota/urlesc
github.com/BurntSushi/toml
github.com/PuerkitoBio/purell
gopkg.in/yaml.v2
github.com/spf13/cast
github.com/kr/pretty
github.com/magiconair/properties
github.com/spf13/fsync
github.com/cpuguy83/go-md2man/md2man
github.com/spf13/hugo/hugofs
github.com/mitchellh/mapstructure
golang.org/x/text/transform
bitbucket.org/pkg/inflect
github.com/dchest/cssmin
github.com/miekg/mmark
github.com/eknkc/amber/parser
github.com/spf13/cobra
github.com/yosssi/ace
golang.org/x/text/unicode/norm
github.com/spf13/nitro
github.com/spf13/hugo/parser
github.com/spf13/viper
github.com/eknkc/amber
github.com/gorilla/websocket
github.com/spf13/hugo/utils
gopkg.in/fsnotify.v1
github.com/spf13/hugo/transform
github.com/spf13/hugo/watcher
github.com/spf13/hugo/livereload
github.com/spf13/hugo/helpers
github.com/spf13/hugo/source
github.com/spf13/hugo/target
github.com/spf13/hugo/tpl
github.com/spf13/hugo/hugolib
github.com/spf13/hugo/create
github.com/spf13/hugo/commands
github.com/spf13/hugo

C:\workspace\hugo>bin\hugo.exe version
Hugo Static Site Generator v0.15-DEV BuildDate: 2015-09-05T13:45:44+09:00
```

実行モジュールは `%GOPATH%\bin` フォルダに置かれる。
実際に使う場合は， `%GOPATH%\bin` フォルダに `PATH` を通すか， `PATH` の通った場所にコピーすればよい。

パッケージをアップデートする場合は `-u` オプションを組み合わせる。

```
C:\workspace\hugo>go get -u -v github.com/spf13/hugo
```

## ブックマーク

- [How to Write Go Code - The Go Programming Language](http://golang.org/doc/code.html) （[日本語版](http://golang-jp.org/doc/code.html)は情報が少し古いので注意）

[Go 言語]: https://golang.org/ "The Go Programming Language"
[前回]: {{< relref "hello.md" >}} "インストールから Hello World まで"
[`os`]: https://golang.org/pkg/os/ "os - The Go Programming Language"
[`fmt`]: https://golang.org/pkg/fmt/ "fmt - The Go Programming Language"
[`flag`]: https://golang.org/pkg/flag/ "flag - The Go Programming Language"
[Defer]: http://blog.golang.org/defer-panic-and-recover "Defer, Panic, and Recover - The Go Blog"
[`gofmt`]: https://golang.org/cmd/gofmt/ "gofmt - The Go Programming Language"
[Subversion]: http://subversion.apache.org/ "Apache Subversion"
[Mercurial]: http://mercurial.selenic.com/ "Mercurial SCM"
[Git]: http://git-scm.com/ "Git"
[Bazaar]: http://bazaar.canonical.com/ "Bazaar"
[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN), 4621300253 (ISBN), 9784621300251 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
