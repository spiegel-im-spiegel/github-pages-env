# そろそろ真面目に Golang 開発環境について考える — Internal Packages と Vendoring

[前回]の続き。保留していた Go 1.5 の Vendoring 機能について。

今回の調査も含めて以下のパッケージを作った

[spiegel-im-spiegel/gcat]: https://github.com/spiegel-im-spiegel/gcat "spiegel-im-spiegel/gcat"
[spiegel-im-spiegel/gutil]: https://github.com/spiegel-im-spiegel/gutil "spiegel-im-spiegel/gutil"

- [spiegel-im-spiegel/gcat]
- [spiegel-im-spiegel/gutil]

[spiegel-im-spiegel/gcat] は `cat` コマンド相当の機能を持つコマンドラインツールで，内部で [spiegel-im-spiegel/gutil] パッケージを呼んでいる。 [spiegel-im-spiegel/gutil] パッケージは雑多な処理を集めたものだが，まだ作りかけ。とりあえず [spiegel-im-spiegel/gcat] で使うものだけを置いている。

まずはこれを導入してみる。 [spiegel-im-spiegel/gutil] を submodule にしてるのでちょっとウザいけどご勘弁を。あと，ユーザ名が長いのはホンマごめん。こんなことならもっと短い名前にするんだった。

```shell
C:>mkdir C:\workspace\gcat

C:>cd C:\workspace\gcat

C:\workspace\gcat>SET GO15VENDOREXPERIMENT=1

C:\workspace\gcat>SET GOPATH=C:\workspace\gcat

C:\workspace\gcat>go get -d github.com/spiegel-im-spiegel/gcat
# cd .; git --git-dir=C:\workspace\gcat\src\github.com\spiegel-im-spiegel\gcat/.git submodule update --init --recursive
No submodule mapping found in .gitmodules for path 'vendor/github.com/spiegel-im-spiegel/gutil'
package github.com/spiegel-im-spiegel/gcat: exit status 1

C:\workspace\gcat>pushd src\github.com\spiegel-im-spiegel\gcat

C:\workspace\gcat\src\github.com\spiegel-im-spiegel\gcat>git submodule init
Submodule 'vendor/github.com/spiegel-im-spiegel/gutil' (https://github.com/spiegel-im-spiegel/gutil.git) registered for path 'vendor/github.com/spiegel-im-spiegel/gutil'

C:\workspace\gcat\src\github.com\spiegel-im-spiegel\gcat>git submodule update
Cloning into 'vendor/github.com/spiegel-im-spiegel/gutil'...
remote: Counting objects: 25, done.
remote: Compressing objects: 100% (20/20), done.
remote: Total 25 (delta 10), reused 15 (delta 4), pack-reused 0
Unpacking objects: 100% (25/25), done.
Checking connectivity... done.
Submodule path 'vendor/github.com/spiegel-im-spiegel/gutil': checked out '7d271650d9937ef0e7b447aff5a55f410f2c9f89'

C:\workspace\gcat\src\github.com\spiegel-im-spiegel\gcat>popd

C:\workspace\gcat>go install -v ./...
github.com/spiegel-im-spiegel/gcat/vendor/github.com/spiegel-im-spiegel/gutil
github.com/spiegel-im-spiegel/gcat/internal/gcat
github.com/spiegel-im-spiegel/gcat/internal/facade
github.com/spiegel-im-spiegel/gcat

C:\workspace\gcat>echo Take the Go-lang! | bin\gcat.exe
Take the Go-lang!
```

導入時のフォルダ構成はこうなっている。

```shell
C:\workspace\gcat>tree .
C:\WORKSPACE\GCAT
├─bin
├─pkg
│  └─windows_amd64
│      └─github.com
│          └─spiegel-im-spiegel
│              └─gcat
│                  ├─internal
│                  └─vendor
│                      └─github.com
│                          └─spiegel-im-spiegel
└─src
    └─github.com
        └─spiegel-im-spiegel
            └─gcat
                ├─internal
                │  ├─facade
                │  └─gcat
                └─vendor
                    └─github.com
                        └─spiegel-im-spiegel
                            └─gutil
```

このフォルダ構成に対して `main.go` の記述は以下のとおりである。

```go:main.go
package main

import (
	"os"

	"github.com/spiegel-im-spiegel/gcat/internal/facade"
	"github.com/spiegel-im-spiegel/gutil"
)

func main() {
	cli := &gutil.CliContext{Reader: os.Stdin, Writer: os.Stdout, ErrorWriter: os.Stderr}
	facadeCxt := &facade.Context{Cli: cli, CommandName: Name, Version: Version}
	os.Exit(facadeCxt.Run(os.Args))
}
```

ポイントになるのは `internal` フォルダと `vendor` フォルダだ。もう少し詳しく見ていこう。

## パッケージ外部からの呼び出しを禁止する Internal Packages

Internal Packages の仕組みは 1.4 のころから存在したが， 1.5 から `GOPATH` 配下のパッケージまで拡張された。

- [Go 1.4 "Internal" Packages](https://docs.google.com/document/d/1e8kOo3r51b2BWtTs_1uADIA5djfXhPT36s6eHVRIvaU/edit)

要するに `internal` フォルダ以下のパッケージは外部から参照できない。これは再利用の難しいビジネスロジックを含むパッケージを配置する場合にはよい仕掛けである。 Internal Packages の制約から外すには `internal` フォルダの外側にパッケージを再配置すればよい。

## Vendoring 機能

環境変数 `GO15VENDOREXPERIMENT` に 1 をセットすると Go 1.5 の Vendoring 機能が使える。

- [Go 1.5 Vendor Experiment](https://docs.google.com/document/d/1Bz5-UB7g2uPBdOx-rw5t9MxJwkfpx90cqG9AFL0JAYo/edit)

Vendoring 機能が有効な状態では `vendor` フォルダが特別な意味を持つ。たとえば `mypackage` パッケージに対して `mypackage/vendor/vpackage` と配置した場合， `import "vpackage"` と記述すれば `mypackage/vendor` フォルダ以下の `vpackage` を探してくれる。上述の [spiegel-im-spiegel/gcat] の場合は `github.com/spiegel-im-spiegel/gutil` がこれにあたり，実体は `github.com/spiegel-im-spiegel/gcat/vendor/github.com/spiegel-im-spiegel/gutil` にある。パッケージが見つからない場合は

```shell
C:\workspace\gcat>go install -v ./...
src\github.com\spiegel-im-spiegel\gcat\internal\gcat\catenate.go:4:2: cannot find package "github.com/spiegel-im-spiegel/gutil" in any of:
        C:\workspace\gcat\src\github.com\spiegel-im-spiegel\gcat\vendor\github.com\spiegel-im-spiegel\gutil (vendor tree)
        C:\Go\src\github.com\spiegel-im-spiegel\gutil (from $GOROOT)
        C:\workspace\gcat\src\github.com\spiegel-im-spiegel\gutil (from $GOPATH)
```

という感じのエラーになり， Vendor tree → `GOROOT` → `GOPATH` の順でパッケージを探していることが分かる。

ただしこの Vendoring 機能は実験段階であり，上手くいけば次かその次のバージョンでは既定で有効になるようだが，当面は様子見というところだろうか。

## ブックマーク

- [Go言語のDependency/Vendoringの問題と今後．gbあるいはGo1.5 | SOTA](http://deeeet.com/writing/2015/06/26/golang-dependency-vendoring/)
- [Big Sky :: golang 1.5 の internal パッケージの使い方。](http://mattn.kaoriya.net/software/lang/go/20150820102400.htm)
- [「golang 1.5 の internal パッケージの使い方。」を試してみた - Qiita](http://qiita.com/qt-luigi/items/d0f52b3b0906b35e6027)

[前回]: http://qiita.com/spiegel-im-spiegel/items/73ebc684b5807277b7e2 "そろそろ真面目に Golang 開発環境について考える（その1） - Qiita"
