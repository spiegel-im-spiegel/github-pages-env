+++
date = "2015-09-28T20:38:45+09:00"
update = "2015-12-08T15:44:37+09:00"
description = "今回は gb を使ってプロジェクト・ベースで Golang のコードを管理してみる。"
draft = false
tags = ["golang", "engineering", "tools", "gb"]
title = "gb でプロジェクト・ベースの開発環境をつくる"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

（初出： [はじめての Go 言語 (on Windows) その9 - Qiita](http://qiita.com/spiegel-im-spiegel/items/ef15a48542e043b32c99)）

今回は [gb] を使ってプロジェクト・ベースで Golang のコードを管理してみる。

- [gb - A project based build tool for Go](http://getgb.io/)

## gb の導入

[gb] の導入は `go get` でできる[^b]。

[^b]: `go get` の使い方については「[go get コマンドでパッケージを管理する]({{< relref "golang/go-get-package.md" >}})」を参照のこと。

```
C:>go get -v github.com/constabulary/gb/...
github.com/constabulary/gb (download)
github.com/constabulary/gb/log
github.com/constabulary/gb
github.com/constabulary/gb/vendor
github.com/constabulary/gb/cmd
github.com/constabulary/gb/cmd/gb
github.com/constabulary/gb/cmd/gb-vendor
```

Windows の場合，環境変数 `GOPATH` で指定するフォルダ配下の `bin` フォルダに `gb.exe` および `gb-vendor.exe` が生成される。
このフォルダにパスを通しておく（またはパスの通っているフォルダに実行ファイルをコピーする）。

## プロジェクトの構築とビルド

「[機能のパッケージ化]({{< relref "golang/packaging.md" >}})」で最後に作ったコードを使って実際に [gb] でプロジェクトを作成しビルドを行ってみる。

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
        } else {
            args[i] = num
        }
    }
    tm := time.Date(args[0], time.Month(args[1]), args[2], 0, 0, 0, 0, time.UTC)
    fmt.Printf("%v\n", tm)
    fmt.Printf("MJD = %d日\n", modjulian.DayNumber(tm))
}
```

### ソース・ファイルの配置

プロジェクト・フォルダを `C:\workspace\gbdemo` とし，ソース・ファイル用のフォルダ `src\julian-day` を作成する。
このフォルダに上述のコードを記述したソース・ファイルを配置する。
フォルダ構成は以下の通り。

```
C:\workspace\gbdemo>tree /f .
C:\WORKSPACE\GBDEMO
└─src
    └─julian-day
            julian-day.go
```

ビルドするには `gb build` コマンドを実行すればいいのだが，このままでは `modjulian` パッケージがないため失敗する。

```
C:\workspace\gbdemo>gb build
FATAL command "build" failed: failed to resolve import path "julian-day": cannot find package "github.com/spiegel-im-spiegel/astrocalc/modjulian" in any of:
        C:\Go\src\github.com\spiegel-im-spiegel\astrocalc\modjulian (from $GOROOT)
        C:\workspace\gbdemo\src\github.com\spiegel-im-spiegel\astrocalc\modjulian (from $GOPATH)
        C:\workspace\gbdemo\vendor\src\github.com\spiegel-im-spiegel\astrocalc\modjulian
```

プロジェクト・フォルダ以下を `GOPATH` として `modjulian` パッケージを探しているのがお分かりだろうか。
[gb] では，実行時に既存の `GOPATH` を上書きするようである。
またプロジェクト・フォルダ配下の `vendor` フォルダを探しているのにも注目してほしい。
            
### 外部パッケージの導入

[gb] では外部パッケージを `gb vendor` コマンドで管理できる。
外部パッケージの導入には `gb vendor fetch` コマンドを使う。

```
C:\workspace\gbdemo>gb vendor fetch github.com/spiegel-im-spiegel/astrocalc/modjulian

C:\workspace\gbdemo>tree /f .
C:\WORKSPACE\GBDEMO
├─src
│  └─julian-day
│          julian-day.go
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
                            LICENSE
                            modjulian.go
                            modjulian_test.go

```

プロジェクト・フォルダ以下に `vendor` フォルダが作成され，パッケージのソースファイルが展開されている。

今回 `gb vendor fetch` で取得したパッケージは [GitHub] のリポジトリから取ってきたものだが， `git clone` ではなく，フォルダ・ファイル構成ごとコピーしてきたもののようである。

`gb vendor fetch` コマンドでは `-branch` や `-tag` や `-revision` オプションでリポジトリのブランチやタグまたはリビジョンを指定できる。
このとき，導入したパッケージのリポジトリ情報は `vender\manifest` ファイルに格納されている（中身は JSON 形式）。

```json
{
    "version": 0,
    "dependencies": [
        {
            "importpath": "github.com/spiegel-im-spiegel/astrocalc/modjulian",
            "repository": "https://github.com/spiegel-im-spiegel/astrocalc",
            "revision": "c9f5fb495e67b868a2b3f0e16c38282095fe5033",
            "branch": "master",
            "path": "/modjulian"
        }
    ]
}
```

ちなみに外部パッケージをアップデートする場合は `gb vendor update` コマンドを使う。

```
C:\workspace\gbdemo>gb vendor update github.com/spiegel-im-spiegel/astrocalc/modjulian
```

または

```
C:\workspace\gbdemo>gb vendor update -all
```

### プロジェクトのビルド

では，この状態でもう一回ビルドしてみる。

```
C:\workspace\gbdemo>gb build
julian-day

C:\workspace\gbdemo>tree /f .
C:\WORKSPACE\GBDEMO
├─bin
│      julian-day.exe
│
├─pkg
│  └─windows-amd64
│      │  julian-day.a
│      │
│      └─github.com
│          └─spiegel-im-spiegel
│              └─astrocalc
│                      modjulian.a
│
├─src
│  └─julian-day
│          julian-day.go
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
                            LICENSE
                            modjulian.go
                            modjulian_test.go

C:\workspace\gbdemo>bin\julian-day.exe 2015 1 1
2015-01-01 00:00:00 +0000 UTC
MJD = 57023日
```

今度は上手くいったようだ。
`gb build` コマンドのオプションは以下の通り。

```
C:\workspace\gbdemo>gb help build
usage: gb build [build flags] [packages]

Build compiles the packages named by the import paths, along with their
dependencies.

Flags:

        -f
                ignore cached packages if present, new packages built will overwrite
                any cached packages. This effectively disables incremental
                compilation.
        -F
                do not cache packages, cached packages will still be used for
                incremental compilation. -f -F is advised to disable the package

                caching system.
        -q
                decreases verbosity, effectively raising the output level to ERROR.
                In a successful build, no output will be displayed.
        -P
                The number of build jobs to run in parallel, including test execution.
                By default this is the number of CPUs visible to gb.
        -R
                sets the base of the project root search path from the current working
                directory to the value supplied. Effectively gb changes working
                directory to this path before searching for the project root.
        -v
                increases verbosity, effectively lowering the output level from INFO
                to DEBUG.
        -dotfile
                if provided, gb will output a dot formatted file of the build steps to
                be performed.
        -ldflags 'flag list'
                arguments to pass on each linker invocation.
        -gcflags 'arg list'
                arguments to pass on each compile invocation.
        -tags 'tag list'
                additional build tags.

The list flags accept a space-separated list of strings. To embed spaces in an
element in the list, surround it with either single or double quotes.

For more about specifying packages, see 'gb help packages'. For more about
where packages and binaries are installed, run 'gb help project'.
```

`-ldflags` や `-gcflags` オプションが使えるのはありがたいかな。

## 複数パッケージを含めたプロジェクト管理

複数のパッケージをまとめて管理したい場合もある。
例えば以下のような構成を考えてみる。

```
C:\workspace\gbdemo>tree /f .
C:\WORKSPACE\GBDEMO
└─src
    └─julian-day
            julian-day.go

C:\workspace\gbdemo>pushd src

C:\workspace\gbdemo\src>git clone https://github.com/spiegel-im-spiegel/astrocalc.git github.com/spiegel-im-spiegel/astrocalc
Cloning into 'github.com/spiegel-im-spiegel/astrocalc'...
remote: Counting objects: 43, done.
remote: Total 43 (delta 0), reused 0 (delta 0), pack-reused 43
Unpacking objects: 100% (43/43), done.
Checking connectivity... done.

C:\workspace\gbdemo\src>popd

C:\workspace\gbdemo>tree /f .
C:\WORKSPACE\GBDEMO
└─src
    ├─github.com
    │  └─spiegel-im-spiegel
    │      └─astrocalc
    │          │  .editorconfig
    │          │  .gitignore
    │          │  .travis.yml
    │          │  LICENSE
    │          │  README.md
    │          │
    │          └─modjulian
    │                  example_test.go
    │                  LICENSE
    │                  modjulian.go
    │                  modjulian_test.go
    │
    └─julian-day
            julian-day.go
```

この状態でビルドを実行してみる。

```
C:\workspace\gbdemo>gb build
github.com/spiegel-im-spiegel/astrocalc/modjulian
julian-day

C:\workspace\gbdemo>tree /f .
C:\WORKSPACE\GBDEMO
├─bin
│      julian-day.exe
│
├─pkg
│  └─windows-amd64
│      │  julian-day.a
│      │
│      └─github.com
│          └─spiegel-im-spiegel
│              └─astrocalc
│                      modjulian.a
│
└─src
    ├─github.com
    │  └─spiegel-im-spiegel
    │      └─astrocalc
    │          │  .editorconfig
    │          │  .gitignore
    │          │  .travis.yml
    │          │  LICENSE
    │          │  README.md
    │          │
    │          └─modjulian
    │                  example_test.go
    │                  LICENSE
    │                  modjulian.go
    │                  modjulian_test.go
    │
    └─julian-day
            julian-day.go

C:\workspace\gbdemo>bin\julian-day.exe 2015 1 1
2015-01-01 00:00:00 +0000 UTC
MJD = 57023日
```
[gb] ではプロジェクト・フォルダ以下にあるパッケージを自動で検索してビルドしてくれる。
もちろんパッケージを指定してビルドすることも可能。

```
C:\workspace\gbdemo>gb build github.com/spiegel-im-spiegel/astrocalc/modjulian
github.com/spiegel-im-spiegel/astrocalc/modjulian

C:\workspace\gbdemo>gb build julian-day
julian-day
```

さらにテストもできる[^a]。

[^a]: テストについては「[Go 言語のテスト・フレームワーク]({{< relref "golang/testing.md" >}})」を参照のこと。

```shell
C:\workspace\gbdemo>gb test -v github.com/spiegel-im-spiegel/astrocalc/modjulian
=== RUN   TestDayNumber
--- PASS: TestDayNumber (0.00s)
=== RUN   ExampleDayNumber
--- PASS: ExampleDayNumber (0.00s)                  
PASS
```

パッケージによっては `go test` の結果と `gb test` の結果が異なる場合があるので注意が必要。

## ブックマーク

- [golang - gbを知ったのでgojiを使ったWEBアプリケーションプロジェクトを管理してみた - Qiita](http://qiita.com/shinofara/items/ac0591fef95c2c6e936e)
- [Go言語のDependency/Vendoringの問題と今後．gbあるいはGo1.5 | SOTA](http://deeeet.com/writing/2015/06/26/golang-dependency-vendoring/)

[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[gb]: http://getgb.io/ "gb - A project based build tool for Go"
[GitHub]: https://github.com/ "GitHub"
