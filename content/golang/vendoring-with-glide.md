+++
date = "2015-12-07T14:07:20+09:00"
description = "description"
draft = true
tags = ["golang", "engineering", "vendoring", "package", "tools", "glide"]
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

以前「[GOPATH 汚染問題]({{< relref "golang/gopath-pollution.md" >}})」で [Go 言語]の 1.5 からの vendoring 機能を紹介したが，この vendoring のヘルパ・ツールと言えるのが [glide] である。

## Glide のインストール

[glide] は自身も [glide] でパッケージ管理している。
なので， `go get` ではなく [Releases](https://github.com/Masterminds/glide/releases) ページからビルド済みのものを取得することをお勧めする[^0]。

[^0]: 一応 `go get` でもビルドできるが， revision を制御できないので失敗する可能性もある。 Linux 等の環境であれば `make` コマンドでもビルドできる。

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
   0.7.2

COMMANDS:
   create, init         Initialize a new project, creating a template glide.yaml
   get                  Install one or more package into `vendor/` and add dependency to glide.yaml.

   import               Import files from other dependency management systems.
   name                 Print the name of this project.
   novendor, nv         List all non-vendor paths in a directory.
   pin                  Print a YAML file with all of the packages pinned to the current version
   rebuild              Rebuild ('go build') the dependencies
   update, up, install  Update or install a project's dependencies
   tree                 Tree prints the dependencies of this project as a tree.
   list                 List prints all dependencies that Glide could discover.
   guess                Guess dependencies for existing source.
   about                Learn about Glide
   help, h              Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --yaml, -y "glide.yaml"                      Set a YAML configuration file.
   --quiet, -q                                  Quiet (no info or debug messages)
   --debug                                      Print Debug messages (verbose)
   --home "C:\Users\Administrator\.glide"       The location of Glide files [$GLIDE_HOME]
   --no-color                                   Turn off colored output for log messages
   --help, -h                                   show help
   --version, -v                                print the version
```

## 開発環境の準備

今回も「[GOPATH 汚染問題]({{< relref "golang/gopath-pollution.md" >}})」で使ったコードを利用する。
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

## 定義ファイルの作成

開発環境ができたら，パッケージのフォルダ（今回は `src/julian-day`）に移動し， `glide create` コマンドで `glide.yaml` ファイルを生成する。

```
C:\workspace\vdemo2>pushd src\julian-day

C:\workspace\vdemo2\src\julian-day>glide create
[INFO] Initialized. You can now edit 'glide.yaml'

C:\workspace\vdemo2\src\julian-day>tree /f C:\workspace\vdemo2
C:\WORKSPACE\VDEMO2
└─src
    └─julian-day
             glide.yaml
            julian-day.go
```

この時点での `glide.yaml` の中身は以下の通り空っぽの状態。

```yaml
# Glide YAML configuration file
# Set this to your fully qualified package name, e.g.
# github.com/Masterminds/foo. This should be the
# top level package.
package: main

# Declare your project's dependencies.
import:
  # Fetch package similar to 'go get':
  #- package: github.com/Masterminds/cookoo
  # Get and manage a package with Git:
  #- package: github.com/Masterminds/cookoo
  #  # The repository URL
  #  repo: git@github.com:Masterminds/cookoo.git
  #  # A tag, branch, or SHA
  #  ref: 1.1.0
  #  # the VCS type (compare to bzr, hg, svn). You should
  #  # set this if you know it.
  #  vcs: git
```

これに [astrocalc]/modjulian の記述を追加する。

```yaml
# Glide YAML configuration file
# Set this to your fully qualified package name, e.g.
# github.com/Masterminds/foo. This should be the
# top level package.
package: main

# Declare your project's dependencies.
import:
  - package: github.com/spiegel-im-spiegel/astrocalc/modjulian
    version: 0.1.0
```

これで [astrocalc]/modjulian パッケージの v0.1.0 を指定できた。

## パッケージの取得とビルド

パッケージの取得には `glide install` コマンドを起動する。

```
C:\workspace\vdemo2\src\julian-day>glide install
[INFO] Fetching updates for github.com/spiegel-im-spiegel/astrocalc.
[INFO] Detected semantic version. Setting version for github.com/spiegel-im-spiegel/astrocalc to v0.1.0.
[INFO] Package github.com/spiegel-im-spiegel/astrocalc manages its own dependencies
[INFO] Detected semantic version. Setting version for github.com/spiegel-im-spiegel/astrocalc to v0.1.0.
[INFO] Project relies on 1 dependencies.

C:\workspace\vdemo2\src\julian-day>tree /f C:\workspace\vdemo2
C:\WORKSPACE\VDEMO2
└─src
    └─julian-day
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

`go get` コマンドと同じく，パッケージのパスから自動的に repository を判別して clone する。
`go get` コマンドと異なる点は， `glide.yaml` ファイルで指定した `version` から適切な revision を取得している点である[^semv]。

[^semv]: このバージョンの解釈は [Semantic Versioning](http://semver.org/) に従っている。バージョンの記述は `package.json` と同様にできるようだ。

VCS (Version Control System) の種類と URI を明示的に指定することもできる。

```yaml
# Glide YAML configuration file
# Set this to your fully qualified package name, e.g.
# github.com/Masterminds/foo. This should be the
# top level package.
package: main

# Declare your project's dependencies.
import:
  - package: github.com/spiegel-im-spiegel/astrocalc/modjulian
    vcs:     git
    repo:    git@github.com:spiegel-im-spiegel/astrocalc.git
    version: 0.1.0
```

たとえばプライベートな repository をインポートする場合には，この方法が有効である。

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


## ブックマーク

- [glide - パッケージ管理のお困りの方へ - - Qiita](http://qiita.com/tienlen/items/8e192e68d6b18bec3b4a)

[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[GOPATH 汚染問題]: {{< relref "golang/gopath-pollution.md" >}} "GOPATH 汚染問題 — プログラミング言語 Go | text.Baldanders.info"
[glide]: https://github.com/Masterminds/glide "Masterminds/glide"
[astrocalc]: https://github.com/spiegel-im-spiegel/astrocalc "spiegel-im-spiegel/astrocalc"
[GitHub]: https://github.com/ "GitHub"
