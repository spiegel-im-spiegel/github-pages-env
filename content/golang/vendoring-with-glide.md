+++
date = "2015-12-07T23:30:21+09:00"
update = "2017-10-10T18:12:56+09:00"
description = "Go 言語 1.5 の vendoring 機能をサポートするツールが glide である。"
draft = false
tags = ["golang", "engineering", "vendoring", "package", "tools", "glide", "gb"]
title = "Glide で Vendoring"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

**【2017-10-10 追記】**
[glide] プロジェクトは Vendoring ツールとして公式ツールの [dep] を推し始めたようである。

- [Glide から Dep への移行を検討する]({{< relref "golang/consider-switching-from-glide-to-dep.md" >}})

そこで当セクションでも今後は [dep] を推していくことにする。
この記事はメンテナンスされない可能性があるが，あしからずご了承の程を。

[dep]: https://github.com/golang/dep "golang/dep: Go dependency management tool"

----

以前「[GOPATH 汚染問題]({{< relref "golang/gopath-pollution.md" >}})」で [Go 言語] 1.5 の vendoring 機能を紹介したが，この vendoring のヘルパ・ツールと言えるのが [glide] である。
[glide] では依存する外部パッケージの情報を YAML 形式の定義ファイルに記述し，この定義ファイルを基に外部パッケージの管理を行う。

（[glide] 0.8 から挙動が少し変わった。
この記事では 0.8.2 を使っている）

## Glide のインストール

[glide] は自身も [glide] で外部パッケージを管理している。
なので最初は `go get` ではなく [Releases](https://github.com/Masterminds/glide/releases) ページからビルド済みのものを取得することをお勧めする[^ins]。

[^ins]: Mac 環境なら brew でインストールできるらしい。 Linux 等の環境であれば `make` コマンドで各種プラットフォームの実行ファイルをビルドできる。一応 `go get` でもビルドできるが， revision を制御できないので失敗する可能性もある（これは `make` コマンドでビルドする場合でも同じだけど）。

既に [glide] が利用可能な状態なら，以下の要領でビルドできる。

```
C:\workspace\glide>SET GOPATH=C:\workspace\glide

C:\workspace\glide>SET GO15VENDOREXPERIMENT=1

C:\workspace\glide>git clone git@github.com:Masterminds/glide.git src\github.com\Masterminds\glide
Cloning into 'src\github.com\Masterminds\glide'...
remote: Counting objects: 2292, done.
remote: Compressing objects: 100% (65/65), done.
remote: Total 2292 (delta 34), reused 0 (delta 0), pack-reused 2292
Receiving objects: 100% (2292/2292), 468.01 KiB | 397.00
Receiving objects: 100% (2292/2292), 537.97 KiB | 397.00 KiB/s, done.
Resolving deltas: 100% (1578/1578), done.
Checking connectivity... done.

C:\workspace\glide>pushd src\github.com\Masterminds\glide

C:\workspace\glide\src\github.com\Masterminds\glide>glide install
[INFO] Fetching updates for github.com/codegangsta/cli.
[INFO] Fetching updates for github.com/Masterminds/cookoo.
[INFO] Fetching updates for github.com/Masterminds/vcs.
[INFO] Fetching updates for gopkg.in/yaml.v2.
[INFO] Fetching updates for github.com/Masterminds/semver.
[INFO] Setting version for github.com/Masterminds/cookoo to 78aa11ce75e257c51be7ea945edb84cf19c4a6de.
[INFO] Setting version for github.com/Masterminds/semver to 6333b7bd29aad1d79898ff568fd90a8aa533ae82.
[INFO] Setting version for github.com/codegangsta/cli to b5232bb2934f606f9f27a1305f1eea224e8e8b88.
[INFO] Setting version for github.com/Masterminds/vcs to eaee272c8fa4514e1572e182faecff5be20e792a.
[INFO] Setting version for gopkg.in/yaml.v2 to f7716cbe52baa25d2e9b0d0da546fcf909fc16b4.

C:\workspace\glide\src\github.com\Masterminds\glide>popd

C:\workspace\glide>go install -v ./...
github.com/Masterminds/glide/vendor/github.com/Masterminds/cookoo/io
github.com/Masterminds/glide/msg
github.com/Masterminds/glide/gb
github.com/Masterminds/glide/vendor/gopkg.in/yaml.v2
github.com/Masterminds/glide/vendor/github.com/Masterminds/semver
github.com/Masterminds/glide/vendor/github.com/codegangsta/cli
github.com/Masterminds/glide/vendor/github.com/Masterminds/cookoo/safely
github.com/Masterminds/glide/vendor/github.com/Masterminds/vcs
github.com/Masterminds/glide/vendor/github.com/Masterminds/cookoo
github.com/Masterminds/glide/vendor/github.com/Masterminds/cookoo/cli
github.com/Masterminds/glide/vendor/github.com/Masterminds/cookoo/database/sql
github.com/Masterminds/glide/vendor/github.com/Masterminds/cookoo/database/active
github.com/Masterminds/glide/vendor/github.com/Masterminds/cookoo/convert
github.com/Masterminds/glide/vendor/github.com/Masterminds/cookoo/fmt
github.com/Masterminds/glide/vendor/github.com/Masterminds/cookoo/web
github.com/Masterminds/glide/vendor/github.com/Masterminds/cookoo/example
github.com/Masterminds/glide/vendor/github.com/Masterminds/cookoo/log
github.com/Masterminds/glide/vendor/github.com/Masterminds/cookoo/web/auth
github.com/Masterminds/glide/util
github.com/Masterminds/glide/vendor/github.com/Masterminds/cookoo/doc
github.com/Masterminds/glide/cfg
github.com/Masterminds/glide/dependency
github.com/Masterminds/glide/cmd
github.com/Masterminds/glide

C:\workspace\glide>bin\glide.exe -v
glide version dev
```

簡単な使い方は以下の通り。

```
C:> glide -h
NAME:
   glide - The lightweight vendor package manager for your Go projects.

Each project should have a 'glide.yaml' file in the project directory. Files
look something like this:

        package: github.com/Masterminds/glide
        imports:
                - package: github.com/Masterminds/cookoo
                  vcs: git
                  ref: 1.1.0
                  subpackages: **
                - package: github.com/kylelemons/go-gypsy
                  subpackages: yaml
                        flatten: true

NOTE: As of Glide 0.5, the commands 'in', 'into', 'gopath', 'status', and 'env'
no longer exist.


USAGE:
   glide [global options] command [command options] [arguments...]

VERSION:
   0.8.2

COMMANDS:
   create, init Initialize a new project, creating a glide.yaml file
   get          Install one or more packages into `vendor/` and add dependency to glide.yaml.
   import       Import files from other dependency management systems.
   name         Print the name of this project.
   novendor, nv List all non-vendor paths in a directory.
   rebuild      Rebuild ('go build') the dependencies
   install, i   Install a project's dependencies
   update, up   Update a project's dependencies
   tree         Tree prints the dependencies of this project as a tree.
   list         List prints all dependencies that Glide could discover.
   about        Learn about Glide
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --yaml, -y "glide.yaml"              Set a YAML configuration file.
   --quiet, -q                          Quiet (no info or debug messages)
   --debug                              Print Debug messages (verbose)
   --home "C:\Users\username\.glide"    The location of Glide files [$GLIDE_HOME]
   --no-color                           Turn off colored output for log messages
   --help, -h                           show help
   --version, -v                        print the version
```

## 開発環境の準備

動作検証用に「[GOPATH 汚染問題]({{< relref "golang/gopath-pollution.md" >}})」で使ったコードを利用する。
まず，以下の環境を作る。

```bash
C:\workspace\vdemo2>SET GOPATH=C:\workspace\vdemo2

C:\workspace\vdemo2>SET GO15VENDOREXPERIMENT=1

C:\workspace\vdemo2>tree /f .
C:\WORKSPACE\VDEMO2
└─src
    └─julian-day
            julian-day.go
```

`julian-day.go` の内容は以下のとおりである。

```go
package main

import (
    "flag"
    "fmt"
    "os"
    "strconv"
    "time"

    "github.com/spiegel-im-spiegel/astrocalc/modjulian"
)

func main() {
    //引数のチェック
    flag.Parse()
    argsStr := flag.Args()
    if len(argsStr) < 3 {
        fmt.Fprintln(os.Stderr, "年月日を指定してください")
        return
    }
    args := make([]int, 3)
    for i := 0; i < 3; i++ {
        num, err := strconv.Atoi(argsStr[i])
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            return
        }
        args[i] = num
    }
    tm := time.Date(args[0], time.Month(args[1]), args[2], 0, 0, 0, 0, time.UTC)
    fmt.Printf("%v\n", tm)
    fmt.Printf("MJD = %d日\n", modjulian.DayNumber(tm))
}
```

当然ながら，このままビルドしても外部パッケージがないため失敗する。

```
C:\workspace\vdemo2>go install ./...
src\julian-day\julian-day.go:10:2: cannot find package "github.com/spiegel-im-spiegel/astrocalc/modjulian" in any of:
        C:\Go\src\github.com\spiegel-im-spiegel\astrocalc\modjulian (from $GOROOT)
        C:\workspace\vdemo2\src\github.com\spiegel-im-spiegel\astrocalc\modjulian (from $GOPATH)
```

## 依存関係を定義する

開発環境ができたら，パッケージのフォルダ（今回は `src/julian-day`）に移動し， `glide create` コマンドで依存関係を定義する `glide.yaml` ファイルを生成する。

```
C:\workspace\vdemo2>pushd src\julian-day

C:\workspace\vdemo2\src\julian-day>glide create
[INFO] Generating a YAML configuration file and guessing the dependencies
[INFO] Attempting to import from other package managers (use --skip-import to skip)
[INFO] Found reference to github.com\spiegel-im-spiegel\astrocalc\modjulian

C:\workspace\vdemo2\src\julian-day>tree /f C:\workspace\vdemo2
C:\WORKSPACE\VDEMO2
└─src
    └─julian-day
            glide.yaml
            julian-day.go
```

ソースコードを読んで `glide.yaml` に必要な情報を推測して書いてくれているようだが[^cre]，デリミタが “/” じゃなくて “\” になってる。
Windows 環境だからかな。
残念 orz

[^cre]: `glide create` 時の依存パッケージの推測を抑止するには `--skip-import` オプションを付ける。のだが，効いてないな。

```yaml
package: julian-day
import:
- package: github.com\spiegel-im-spiegel\astrocalc\modjulian
```

[astrocalc]/modjulian パッケージのパスを修正して、更にバージョン情報を追加する。

```yaml
package: julian-day
import:
- package: github.com/spiegel-im-spiegel/astrocalc/modjulian
  version: v0.1.0
```

これで `go get` コマンドと同じように， `package` のパスから自動的に repository を判別してパッケージを取得できる。
`go get` コマンドと異なるのは， `glide.yaml` ファイルで指定した `version` 情報から適切な revision を選択できる点である[^semv]。

[^semv]: 今回であれば repository の [`v0.1.0`](https://github.com/spiegel-im-spiegel/astrocalc/releases/tag/v0.1.0) タグに対応する revision を選択する。バージョンの記述形式は `package.json` と同じように記述でき，バージョンの解釈は [Semantic Versioning](http://semver.org/) に従っている。ちなみに revision ID を直接指定することもできる。

また，以下のように VCS (Version Control System) の種類[^vcs] と URI を明示的に指定することもできる（`vcs` と `repo` は必ずセットで指定する）。

[^vcs]: [glide] では [git](http://git-scm.com/) のほか svn ([Subversion](http://subversion.apache.org/)), hg ([Mercurial](http://mercurial.selenic.com/)), bzr ([Bazaar](http://bazaar.canonical.com/)) が利用可能である。

```yaml
package: julian-day
import:
- package: github.com/spiegel-im-spiegel/astrocalc/modjulian
  vcs:     git
  repo:    git@github.com:spiegel-im-spiegel/astrocalc.git
  version: v0.1.0
```

たとえば，プライベートな bare repository からインポートする場合には，この方法が有効である。

## パッケージの取得とビルド

パッケージの取得には `glide update` または `glide up` コマンドを起動する。

```
C:\workspace\vdemo2\src\julian-day>glide up
[INFO] Fetching updates for github.com/spiegel-im-spiegel/astrocalc.
[INFO] Setting version for github.com/spiegel-im-spiegel/astrocalc to v0.1.0.
[INFO] Scanning github.com/spiegel-im-spiegel/astrocalc for dependencies.
[INFO] Setting version for github.com/spiegel-im-spiegel/astrocalc to v0.1.0.
[INFO] Project relies on 1 dependencies.
[INFO] Writing glide.lock file
C:\workspace\vdemo2\src\julian-day>tree /f C:\workspace\vdemo2
C:\WORKSPACE\VDEMO2
└─src
    └─julian-day
        │  glide.lock
        │  glide.yaml
        │  julian-day.go
        │
        └─vendor
            └─github.com
                └─spiegel-im-spiegel
                    └─astrocalc
                        │  .editorconfig
                        │  .gitignore
                        │  .travis.yml
                        │  LICENSE
                        │  README.md
                        │
                        └─modjulian
                                example_test.go
                                LICENSE
                                modjulian.go
                                modjulian_test.go
```

これでビルドが可能になった。
ではビルドしようかな。

```
C:\workspace\vdemo2\src\julian-day>popd

C:\workspace\vdemo2>go install -v  ./...
julian-day/vendor/github.com/spiegel-im-spiegel/astrocalc/modjulian
julian-day

C:\workspace\vdemo2>bin\julian-day.exe 2015 1 1
2015-01-01 00:00:00 +0000 UTC
MJD = 57023日
```

よーし，うむうむ，よーし。

## Vendor フォルダの管理

`glide.lock` ファイルには外部パッケージの状態が記述されている。

```yaml
hash: b1eecd43769aa1225f7b8922a35566008561c64f4ce11180f299888ed153ad69
updated: 2015-12-23T10:47:22.4140434+09:00
imports:
- name: github.com/spiegel-im-spiegel/astrocalc
  version: 256cd3bd4dc63441d3f4759d761593a147d51d88
  subpackages:
  - modjulian
devImports: []
```

`glide.lock` があれば `glide install`（または `glide i`）コマンドで `vendor` フォルダ以下の状態を復元できる。

[glide] では外部パッケージを vendor フォルダ以下に repository 構造ごと展開する。
この場合，開発対象のパッケージも repository で管理しているのだから， repository が入れ子になり具合が悪い。
その辺，当の [glide] はどうしてるのかなぁと思ったら `.gitignore` ファイルで `vendor/` を除外対象にしていた。
なるほど，そりゃそうか。

`glide.yaml` および `glide.lock` ファイルの管理さえちゃんとしていれば `glide install` コマンドでいつでも復元できるのだから `vendor` フォルダ以下を除外しても問題ないわけだ[^v]。
これは複数メンバで開発環境を共有する際には便利な機能である。

また vendoring に対応していない（Go 1.4 以下の）環境や [glide] がない環境では `go get` で外部パッケージを取ってくることで（revision 等の問題はあるけど）一応ビルドは通る。

[^v]: `vendor` フォルダ以下は外部パッケージなので通常はさわることはない。

更に言うと， [glide] は [Go 言語]の標準機能に準拠しているため，他のサポートツールとの相性がいいのも利点だろう。
たとえば， [ATOM ベースの開発環境]({{< relref "golang/golang-with-atom.md" >}})は [glide] と相性がいい[^gov]。
あと，（多少強引な手を使っているが[^tci]） [Travis CI](https://travis-ci.org/) のような CI (Continuous Integration) と組み合わせることも難しくない。

[^gov]: 残念ながら，「[パッケージの依存状況の視覚化]({{< relref "golang/package-visualization-tool.md" >}})」ツールは vendoring 機能に対応していないため上手く表示できない。なお， [glide] では `glide list` および `glide tree` で依存パッケージを見ることができる。
[^tci]: [glide] の [`.travis.yml`](https://github.com/Masterminds/glide/blob/master/.travis.yml) や [`Makefile`](https://github.com/Masterminds/glide/blob/master/Makefile) を参照。

こう考えると [glide] は[前に紹介]({{< relref "golang/project-based-development.md" >}})した [gb](http://getgb.io/) よりも筋がいいツールといえるかもしれない。

## おまけ： vender フォルダ以下をテストから除外する

テストを行う際に普通に

```text
$ go test -v ./...
```

とかやると `vendor` フォルダ以下も対象になってしまう。
`vendor` フォルダ以下を除外したいのであれば

```text
$ go test -v $(glide novendor)
```

とすればよい。

ただこれ Windows のコマンドプロンプトでは使えないんだよねぇ。
パイプで `go test` に渡してもうまくいかない感じ。
そういう場合は `glide novendor` の実行結果を整形してバッチファイルにするしかないのだろう。
やれやれ。

## ブックマーク

- [glide - パッケージ管理のお困りの方へ - - Qiita](http://qiita.com/tienlen/items/8e192e68d6b18bec3b4a)
- [Golangでプロジェクト内のテストを全件実行する - Qiita](http://qiita.com/ktsujichan/items/c78e2515c459316cb1f6)

[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[GOPATH 汚染問題]: {{< relref "golang/gopath-pollution.md" >}} "GOPATH 汚染問題 — プログラミング言語 Go | text.Baldanders.info"
[glide]: https://github.com/Masterminds/glide "Masterminds/glide"
[astrocalc]: https://github.com/spiegel-im-spiegel/astrocalc "spiegel-im-spiegel/astrocalc"
[GitHub]: https://github.com/ "GitHub"
