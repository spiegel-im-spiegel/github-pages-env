# はじめての Go 言語 (on Windows) その9

（[これまでの記事の目次](http://qiita.com/spiegel-im-spiegel/items/98d49ac456485b007a15#%E3%81%AF%E3%81%98%E3%82%81%E3%81%A6%E3%81%AE-go-%E8%A8%80%E8%AA%9E-on-windows)。同ページのブックマークも参考にどうぞ）

このシリーズではしばらくぶり。今回は [gb] を使ってプロジェクト・ベースで Golang のソースを管理してみる。

- [gb]
- [constabulary/gb](https://github.com/constabulary/gb)

## gb の導入

[gb] の導入は `go get` コマンド一発である（`go get` コマンドの使い方は「[その3](http://qiita.com/spiegel-im-spiegel/items/a52a47942fd3946bb583)」をどうぞ）。

```shell
C:>go get -v github.com/constabulary/gb/...
github.com/constabulary/gb (download)
github.com/constabulary/gb
github.com/constabulary/gb/cmd
github.com/constabulary/gb/cmd/gb-vendor/vendor
github.com/constabulary/gb/cmd/gb
github.com/constabulary/gb/cmd/gb-vendor
```

環境変数 `GOPATH` で指定するフォルダ配下の `bin` フォルダに `gb.exe` および `gb-vendor.exe` が生成される。このフォルダに `PATH` を通しておく。

## プロジェクトの構築とビルド

### ソースコードの配置

プロジェクト・フォルダを `C:\path\to\project` とし，ソース・ファイル用のフォルダを作成する。

```shell
C:\path\to\project>mkdir src\gbdemo
C:\path\to\project>tree . /A
C:\PATH\TO\PROJECT
└─src
    └─gbdemo
```

更に `C:\path\to\project\src\gbdemo` フォルダにソース・ファイルを作成する。今回は「[その6](http://qiita.com/spiegel-im-spiegel/items/404871d2bafd22bdbb90)」で最後に作ったファイルを流用する。

```go:gbdemo/main.go
package main

import (
	"flag"
	"fmt"
	"github.com/spiegel-im-spiegel/astrocalc/modjulian"
	"strconv"
	"time"
)

func main() {
	//引数のチェック
	flag.Parse()
	argsStr := flag.Args()
	if len(argsStr) < 3 {
		fmt.Printf("年月日を指定してください")
		return
	}
	args := make([]int, 3)
	for i := 0; i < 3; i++ {
		num, err := strconv.Atoi(argsStr[i])
		if err != nil {
			if enum, ok := err.(*strconv.NumError); ok {
				switch enum.Err {
				case strconv.ErrRange:
					fmt.Printf("Bad Range Error")
				case strconv.ErrSyntax:
					fmt.Printf("Syntax Error")
				}
			}
			return
		} else {
			args[i] = num
		}
	}
	tm := time.Date(args[0], time.Month(args[1]), args[2], 0, 0, 0, 0, time.UTC)
	fmt.Printf("%v\n", tm)
	fmt.Printf("MJD = %d日\n", modjulian.DayNumber(tm))
}
```

あとは `gb build` コマンドでビルドすればいいのだが，ここで `github.com/spiegel-im-spiegel/astrocalc/modjulian` パッケージがないとエラーを吐いてしまう。

```shell
C:\path\to\project>gb build
FATAL command "build" failed: failed to resolve import path "gbdemo": cannot find package "github.com/spiegel-im-spiegel/astrocalc/modjulian" in any of:
        C:\Go\src\github.com\spiegel-im-spiegel\astrocalc\modjulian (from $GOROOT)
        C:\path\to\project\src\github.com\spiegel-im-spiegel\astrocalc\modjulian (from $GOPATH)
        C:\path\to\project\vendor\src\github.com\spiegel-im-spiegel\astrocalc\modjulian
```

`GOPATH` にプロジェクト・フォルダが設定されていることにお気づきだろうか。 [gb] では既存の `GOPATH` 設定は使わないようだ。

### 外部パッケージの導入

外部パッケージを導入するには `gb vendor fetch` コマンドを使う。

```shell
C:\path\to\project>gb vendor fetch github.com/spiegel-im-spiegel/astrocalc/modjulian
C:\path\to\project>tree . /A
C:\PATH\TO\PROJECT
├─src
│  └─gbdemo
│          main.go
│
└─vendor
    │  manifest
    │
    └─src
        └─github.com
            └─spiegel-im-spiegel
                └─astrocalc
                    └─modjulian
                            example_test.go
                            modjulian.go
                            modjulian_test.go
```

プロジェクト・フォルダ以下に `vendor` フォルダが作成され，パッケージのソースファイルが展開できているのがお分かりだろうか。ちなみに今回 `gb vendor fetch` で取得したパッケージは [GitHub] のリポジトリから取ってきたものだが， clone されたものではなく，（少なくとも見た目は）単純に export してきたもののようである（プロジェクト・フォルダ全体をソースコード管理する場合に `vendor` フォルダ以下に別（システム）のリポジトリがあると管理が面倒臭くなるからかな？）。

ちなみに外部パッケージをアップデートする場合は `gb vendor update` コマンドを使う。

```shell
C:\path\to\project>gb vendor update github.com/spiegel-im-spiegel/astrocalc/modjulian
```

または

```shell
C:\path\to\project>gb vendor update -all
```

### プロジェクトのビルド

では，この状態でもう一回ビルドしてみる。

```shell
C:\path\to\project>gb build
github.com/spiegel-im-spiegel/astrocalc/modjulian
gbdemo

C:\path\to\project>tree . /A
C:\PATH\TO\PROJECT
├─bin
│      gbdemo.exe
│
├─pkg
│  └─windows
│      └─amd64
│          └─github.com
│              └─spiegel-im-spiegel
│                  └─astrocalc
│                          modjulian.a
│
├─src
│  └─gbdemo
│          main.go
│
└─vendor
    │  manifest
    │
    └─src
        └─github.com
            └─spiegel-im-spiegel
                └─astrocalc
                    └─modjulian
                            example_test.go
                            modjulian.go
                            modjulian_test.go

C:\path\to\project>bin\gbdemo 2015 1 1
2015-01-01 00:00:00 +0000 UTC
MJD = 57023日
```

今度は上手くいったようだ。

ちなみに `gb build` コマンドのオプションは以下の通り。

```shell
C:\path\to\project>gb help build
usage: gb build [build flags] [packages]

Build compiles the packages named by the import paths, along with their dependencies.

The build flags are

        -f
                ignore cached packages if present, new packages built will overwrite any cached packages.
                This effectively disables incremental compilation.
        -F
                do not cache packages, cached packages will still be used for in cremental compilation.
                -f -F is advised to disable the package caching system.
        -q
                decreases verbosity, effectively raising the output level to ERROR.
                In a successful build, no output will be displayed.
        -R
                sets the base of the project root search path from the current working directory to the value supplied.
                Effectively gb changes working directory to this path before searching for the project root.
        -v
                increases verbosity, effectively lowering the output level from INFO to DEBUG.
        -ldflags 'flag list'
                arguments to pass on each linker invocation.
        -gcflags 'arg list'
                arguments to pass on each go tool compile invocation.

The list flags accept a space-separated list of strings. To embed spaces in an element in the list, surround it with either single or double quotes.

For more about specifying packages, see 'gb help packages'. For more about where packages and binaries are installed, run 'gb help project'.
```

`-ldflags` や `-gcflags` が使えるのはありがたいかな。


## 複数パッケージを含めたプロジェクト管理

いくつかのパッケージを含んだプロジェクト管理を行いたい場合もある。例えば以下のような構成を考えてみる。

```shell
C:\path\to\project>tree . /A
C:\PATH\TO\PROJECT
└─src
    └─gbdemo
            main.go

C:\path\to\project>pushd src
C:\path\to\project\src>mkdir github.com\spiegel-im-spiegel
C:\path\to\project\src>cd github.com\spiegel-im-spiegel
C:\path\to\project\src\github.com\spiegel-im-spiegel>git clone git@github.com:spiegel-im-spiegel/astrocalc.git
Cloning into 'astrocalc'...
remote: Counting objects: 35, done.
remote: Total 35 (delta 0), reused 0 (delta 0), pack-reused 34
Unpacking objects: 100% (35/35), done.
Checking connectivity... done.

C:\path\to\project\src\github.com\spiegel-im-spiegel>popd
C:\path\to\project>tree . /A
C:\PATH\TO\PROJECT
└─src
    ├─gbdemo
    │      main.go
    │
    └─github.com
        └─spiegel-im-spiegel
            └─astrocalc
                │  .gitignore
                │  .travis.yml
                │  LICENSE
                │  README.md
                │
                └─modjulian
                        example_test.go
                        modjulian.go
                        modjulian_test.go
```

この状態でビルドを実行してみる。

```shell
C:\path\to\project>gb build
github.com/spiegel-im-spiegel/astrocalc/modjulian
gbdemo

C:\path\to\project>tree . /A
C:\PATH\TO\PROJECT
├─bin
│      gbdemo.exe
│
├─pkg
│  └─windows
│      └─amd64
│          └─github.com
│              └─spiegel-im-spiegel
│                  └─astrocalc
│                          modjulian.a
│
└─src
    ├─gbdemo
    │      main.go
    │
    └─github.com
        └─spiegel-im-spiegel
            └─astrocalc
                │  .gitignore
                │  .travis.yml
                │  LICENSE
                │  README.md
                │
                └─modjulian
                        example_test.go
                        modjulian.go
                        modjulian_test.go

C:\path\to\project>bin\gbdemo 2015 1 1
2015-01-01 00:00:00 +0000 UTC
MJD = 57023日
```

さらにテストもできる（テストの書き方は「[その7](http://qiita.com/spiegel-im-spiegel/items/64224f22ef17d916dc2d)」をどうぞ）。

```shell
C:\path\to\project>gb test
github.com/spiegel-im-spiegel/astrocalc/modjulian
testing: warning: no tests to run
PASS
PASS
```

（ワーニングが出るのは `gbdemo` パッケージにテストがないため）

## ブックマーク

- [golang - gbを知ったのでgojiを使ったWEBアプリケーションプロジェクトを管理してみた - Qiita](http://qiita.com/shinofara/items/ac0591fef95c2c6e936e)
- [Go言語のDependency/Vendoringの問題と今後．gbあるいはGo1.5 | SOTA](http://deeeet.com/writing/2015/06/26/golang-dependency-vendoring/)

[gb]: http://getgb.io/ "gb - A project based build tool for Go"
[GitHub]: https://github.com/ "GitHub"
