# はじめての Go 言語 (on Windows) その3

[これまで](http://qiita.com/spiegel-im-spiegel/items/047a9bd6436e6391ddd4)は Go 言語そのものについてさわって遊んでみたのだが，今回はその周辺機能について試してみる。

## get コマンドを使ってアプリケーションをインストールする

`get` コマンドを使って公開 repository からアプリケーションおよびその依存パッケージをインストールすることができる。ここでは最近「ポスト Jekyll」として注目されている [Hugo] をインストールしてみる。

- [Hugo :: A fast and modern static website engine](http://gohugo.io/)

[Hugo] はバイナリでも配布されているので，別にソースコードからコンパイルする必要はないのだが，それでも中身は見てみたいし， get コマンドの練習として使えるかな，と思ったり。あと，私自身，メインで使ってる Movable Type を何とかしないと MTOS のサポートも今年で切れちゃうし。

[Hugo]: http://gohugo.io/ "Hugo :: A fast and modern static website engine"


## get コマンドの要件

### GOPATH

`get` コマンドでは，環境変数 `GOPATH` に取得したソースコードやコンパイル後のモジュールを格納する。 `GOPATH` がないと以下のように怒られる。

```shell
C:>go get
package github.com/spf13/hugo/commands: cannot download, $GOPATH not set. For more details see: go help gopath
```

`GOPATH` の詳しい解説は以下のとおり。
あとは実際に `get` コマンド後の `GOPATH` フォルダ内を覗いてみてみたほうが早いだろう。

```shell
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

The src/ directory holds source code.  The path below 'src'
determines the import path or executable name.

The pkg/ directory holds installed package objects.
As in the Go tree, each target operating system and
architecture pair has its own subdirectory of pkg
(pkg/GOOS_GOARCH).

If DIR is a directory listed in the GOPATH, a package with
source in DIR/src/foo/bar can be imported as "foo/bar" and
has its compiled form installed to "DIR/pkg/GOOS_GOARCH/foo/bar.a".

The bin/ directory holds compiled commands.
Each command is named for its source directory, but only
the final element, not the entire path.  That is, the
command with source in DIR/src/foo/quux is installed into
DIR/bin/quux, not DIR/bin/foo/quux.  The foo/ is stripped
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
```

### 内部コマンド

`get` コマンドは内部で以下のバージョン管理ツールを呼び出す。

- svn : [Subversion]
- hg : [Mercurial]
- git : [Git]
- bzr : [Bazaar]

たとえば get コマンド実行時に

```shell
C:>go get
go: missing Mercurial command. See http://golang.org/s/gogetcmd
package github.com/spf13/hugo/commands
        imports bitbucket.org/pkg/inflect: exec: "hg": executable file not found in %PATH%
```

などとエラーが出たらいずれかのバージョン管理ツール（この場合は Mercurial）がないことになる。いずれのツールも Windows 版が存在するので，必要に応じてインストールしておくとよい。

ちなみに [Hugo] では git と Mercurial が必要である。

[Subversion]: http://subversion.apache.org/ "Apache Subversion"
[Mercurial]: http://mercurial.selenic.com/ "Mercurial SCM"
[Git]: http://git-scm.com/ "Git"
[Bazaar]: http://bazaar.canonical.com/ "Bazaar"

### Proxy 設定

GitHub や Bitbucket 等のよく知られたリポジトリから取得する場合は直に上記のコマンドを起動するようだけど，それ以外のサイトでは HTTP/HTTPS でいったん fetch してから，どのコマンドで取得するか判断してるみたい。このとき Firewall/Proxy で阻まれている環境ではエラーになってしまう。

この場合は Proxy の設定が必要。設定には `http_proxy` および `https_proxy` 環境変数をセットする。たとえばこんな感じ。

```shell
SET http_proxy=http://username:password@proxy.exsample.com:8080/
SET https_proxy=%http_proxy%
```

Proxy 設定に username:password を含む場合はレジストリに環境変数を設定するのではなく（コントロールパネル → システム → システムん詳細設定 で［環境変数(N)...］ボタンを押すと環境変数の設定画面が開く）、コマンドプロンプトなどから都度手作業で設定するほうが安全... なんだけど面倒くさいよね，やっぱ。

## Hugo のインストール

以上の準備ができていたら [Hugo] のインストールはコマンド一発である。

[Hugo] の repository は [spf13/hugo](https://github.com/spf13/hugo) にあるので，以下のようにコマンドを起動する。

```shell
C:>go get -v github.com/spf13/hugo
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
github.com/inconshreveable/mousetrap
github.com/spf13/hugo/bufferpool
github.com/kr/text
github.com/kardianos/osext
github.com/spf13/afero
github.com/spf13/jwalterweatherman
github.com/spf13/pflag
github.com/russross/blackfriday
github.com/opennota/urlesc
github.com/BurntSushi/toml
gopkg.in/yaml.v2
github.com/kr/pretty
github.com/spf13/cast
# github.com/opennota/urlesc
runtime: failed to create new OS thread (have 5 already; errno=5)
fatal error: runtime.newosproc
github.com/PuerkitoBio/purell
github.com/spf13/fsync
github.com/spf13/hugo/hugofs
github.com/magiconair/properties
github.com/mitchellh/mapstructure
github.com/cpuguy83/go-md2man/md2man
golang.org/x/text/transform
bitbucket.org/pkg/inflect
github.com/miekg/mmark
github.com/dchest/cssmin
github.com/spf13/cobra
github.com/eknkc/amber/parser
golang.org/x/text/unicode/norm
github.com/yosssi/ace
github.com/spf13/nitro
github.com/spf13/hugo/parser
github.com/spf13/viper
github.com/eknkc/amber
github.com/gorilla/websocket
github.com/spf13/hugo/utils
gopkg.in/fsnotify.v1
github.com/spf13/hugo/transform
github.com/spf13/hugo/watcher
github.com/spf13/hugo/helpers
github.com/spf13/hugo/livereload
github.com/spf13/hugo/source
github.com/spf13/hugo/target
github.com/spf13/hugo/tpl
github.com/spf13/hugo/hugolib
github.com/spf13/hugo/create
github.com/spf13/hugo/commands
github.com/spf13/hugo

C:>hugo version
Hugo Static Site Generator v0.15-DEV BuildDate: 2015-08-24T10:24:01+09:00
```

`-v` オプションを付ければ依存関係も全部表示してくれるので分かりやすい。

アップデートが必要なときは `-u` オプションを使って

```shell
C:>go get -u -v github.com/spf13/hugo
```

でよい。

実行モジュールは `%GOPATH%\bin` フォルダに出力される。そのままこのフォルダにパスを通してもいいし別の場所にコピーしてもよい。

## さぁて

ゴールデンウィークだし，遊ぶぞ！

## ブックマーク

- [Jekyllが許されるのは小学生までだよね - MOL](http://t32k.me/mol/log/hugo/)
- [CSS - HUGOを使ってサイトを立ち上げる方法 - Qiita](http://qiita.com/syui/items/869538099551f24acbbf)
- [静的サイトジェネレーターHugoでgithubページ - Qiita](http://qiita.com/kmry@github/items/01f68550880dcc6d01fe)
