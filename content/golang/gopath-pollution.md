+++
date = "2015-09-29T23:40:59+09:00"
update = "2018-09-26T13:52:50+09:00"
description = "go get コマンドは外部パッケージの revision 等をコントロールできず，常に repository の最新コードを取ってこようとする。GOPATH 内に複数のプロジェクトが同居している場合は同じ外部パッケージでもプロジェクトごとに異なるリビジョンを要求する場合があり，管理が煩雑になってしまう。"
draft = false
tags = ["golang", "engineering", "environment", "pollution", "vendoring", "package"]
title = "GOPATH 汚染問題"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "https://baldanders.info/profile/"
+++

{{% div-box type="md" %}}
**【注意 2018-09-26】**
この問題はバージョン 1.11 からサポートされる「モジュール」機能によって解消可能です。
もはやこの記事の内容は古いものであり「こんな時代もあったね」と生暖かい気持ちで読んでいただければ幸いです。
{{% /div-box %}}

（初出： [そろそろ真面目に Golang 開発環境について考える — GOPATH 汚染問題 - Qiita](http://qiita.com/spiegel-im-spiegel/items/73ebc684b5807277b7e2)，[そろそろ真面目に Golang 開発環境について考える — Internal Packages と Vendoring - Qiita](http://qiita.com/spiegel-im-spiegel/items/baa3671c7e1b8a6594a9)）

`go get` コマンドはとても強力な機能で，私のように Windows と UNIX 系環境の間を渡り歩いてる身としては， make などの tool chain に大きく依存することなく， `go get` コマンドだけで repository の fetch からビルド・インストールまで出来てしまうのは非常にありがたい[^a]。

[^a]: それでも git などのコード管理ツールへの依存はどうしても残るのだけれど。

しかし， `go get` コマンドは外部パッケージの revision 等をコントロールできず，常に repository の最新コードを取ってこようとする。
ひとつの環境でひとつのプロジェクトを管理していくのならこれでも何とかならないこともないが， `GOPATH` 内に複数のプロジェクトが同居している場合は同じ外部パッケージでもプロジェクトごとに異なるリビジョンを要求する可能性があり，管理が煩雑になってしまう。

しかも困ったことに `GOPATH` 環境変数は複数のプロジェクト管理を想定していないため，全てのパッケージをひとつのフォルダに入れようとする[^c] [^d]。

[^c]: 具体的には `GOPATH` で列挙されるパスのリストのうち先頭のパスにインストールされる。
[^d]: [Go 言語]の開発・管理主体は Google だが，こんな構成で Google は困らないのかと思ったのだが，実は Google は全てのコードを単一の repository で管理しているらしい。（参考： [20億行のコードを保存し、毎日4万5000回のコミットを発行しているGoogleが、単一のリポジトリで全社のソースコードを管理している理由](http://www.publickey1.jp/blog/15/2045000google.html)）

## 【対策1】 プロジェクトごとに GOPATH を設定し直す

この問題に対する一番安直な答えは「プロジェクトごとに `GOPATH` を設定し直す」である。例えば[前回]紹介した [gb] をビルドする場合は以下のようにする。

```
C:>mkdir C:\workspace\gb

C:>SET GOPATH=C:\workspace\gb

C:>go get -v github.com/constabulary/gb/...
github.com/constabulary/gb (download)
github.com/constabulary/gb/log
github.com/constabulary/gb
github.com/constabulary/gb/vendor
github.com/constabulary/gb/cmd
github.com/constabulary/gb/cmd/gb
github.com/constabulary/gb/cmd/gb-vendor
```

あとは `GOPATH` 直下の `bin` フォルダにパスを通すか，パスの通ってるフォルダに実行ファイルをコピーすればよい。
実行履歴はバッチファイル（UNIX 系なら shell スクリプト）に保存しておけばいつでも復元できる。

毎回環境をセットアップしないといけないのは面倒だが，プロジェクト管理のためのツールも必要なく， Go コンパイラの標準機能のみで管理できる。
標準機能のみで管理できるというのは結構重要で，たとえば CI ツールを使っている場合は，設定を単純にできるので管理しやすいといえる。

UNIX 系の環境であれば [direnv] を使う手もある[^b]。
[direnv] は `cd` をフックし，ディレクトリごとに環境変数を書き換えることができる。
この機能を使ってプロジェクト・フォルダごとに `GOPATH` を設定できる。

[^b]: [direnv] は [Go 言語]で組まれている。

## 【対策2】 プロジェクト・ベースの管理ツールを使う

もうひとつは [gb] のようなプロジェクト・ベースでコード管理のできるツールを使う方法である。
[gb] については[前回]紹介したので，そちらを参照のこと。

[gb] で作った開発環境はフォルダ構成を丸ごと開発メンバに配布・同期することが可能になるため，複数人で環境を合わせることが容易になる。

## 【対策3】 Go 1.5 の Vendoring 機能を使う

[Go 言語]のバージョン 1.5 から Vendoring 機能が使えるようになった。

Vendoring 機能を使うと，外部パッケージを `GOPATH` とは独立に管理できるようになる。
この機能を使うには環境変数 `GO15VENDOREXPERIMENT` に 1 をセットする。

（**追記** 当初の予告通り Vendoring 機能は 1.6 から既定の機能になった。環境変数 `GO15VENDOREXPERIMENT` をセットしなくても有効になる）

- [Go 1.5 Vendor Experiment](https://docs.google.com/document/d/1Bz5-UB7g2uPBdOx-rw5t9MxJwkfpx90cqG9AFL0JAYo/edit)

Vendoring 機能が有効な状態では `vendor` フォルダが特別な意味を持つ。
たとえば `mypackage` パッケージに対して `mypackage/vendor/vpackage` と配置した場合， `import "vpackage"` と記述すれば `mypackage/vendor` フォルダ以下の `vpackage` も探してくれる。

では，[前回]作ったコードを流用して確かめてみる。

```
C:\workspace\vdemo>SET GOPATH=C:\workspace\vdemo

C:\workspace\vdemo>SET GO15VENDOREXPERIMENT=1

C:\workspace\vdemo>tree /f .
C:\WORKSPACE\VDEMO
└─src
    └─julian-day
            julian-day.go

C:\workspace\vdemo>go build ./...
src\julian-day\julian-day.go:10:2: cannot find package "github.com/spiegel-im-spiegel/astrocalc/modjulian" in any of:
        C:\Go\src\github.com\spiegel-im-spiegel\astrocalc\modjulian (from $GOROOT)
        C:\workspace\vdemo\src\github.com\spiegel-im-spiegel\astrocalc\modjulian (from $GOPATH)

C:\workspace\vdemo>mkdir src\julian-day\vendor

C:\workspace\vdemo>tree /f .
C:\WORKSPACE\VDEMO
└─src
    └─julian-day
        │  julian-day.go
        │
        └─vendor


C:\workspace\vdemo>go build ./...
src\julian-day\julian-day.go:10:2: cannot find package "github.com/spiegel-im-spiegel/astrocalc/modjulian" in any of:
        C:\workspace\vdemo\src\julian-day\vendor\github.com\spiegel-im-spiegel\astrocalc\modjulian (vendor tree)
        C:\Go\src\github.com\spiegel-im-spiegel\astrocalc\modjulian (from $GOROOT)
        C:\workspace\vdemo\src\github.com\spiegel-im-spiegel\astrocalc\modjulian (from $GOPATH)
```

`vendor` フォルダを追加したことで Go コンパイラの挙動が変わったことがお分かりだろうか。
目的のパッケージを vendor tree → `GOROOT` → `GOPATH` の順で捜索している。

では `vendor` フォルダに外部パッケージを導入してビルドしてみよう。

```
C:\workspace\vdemo>pushd src\julian-day\vendor

C:\workspace\vdemo\src\julian-day\vendor>git clone https://github.com/spiegel-im-spiegel/astrocalc.git github.com/spiegel-im-spiegel/astrocalc
Cloning into 'github.com/spiegel-im-spiegel/astrocalc'...
remote: Counting objects: 43, done.
remote: Total 43 (delta 0), reused 0 (delta 0), pack-reused 43
Unpacking objects: 100% (43/43), done.
Checking connectivity... done.

C:\workspace\vdemo\src\julian-day\vendor>popd

C:\workspace\vdemo>tree /f .
C:\WORKSPACE\VDEMO
└─src
    └─julian-day
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

C:\workspace\vdemo>go install -v ./...
julian-day/vendor/github.com/spiegel-im-spiegel/astrocalc/modjulian
julian-day

C:\workspace\vdemo>tree /f .
C:\WORKSPACE\VDEMO
├─bin
│      julian-day.exe
│
├─pkg
│  └─windows_amd64
│      └─julian-day
│          └─vendor
│              └─github.com
│                  └─spiegel-im-spiegel
│                      └─astrocalc
│                              modjulian.a
│
└─src
    └─julian-day
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

C:\workspace\vdemo>bin\julian-day.exe 2015 1 1
2015-01-01 00:00:00 +0000 UTC
MJD = 57023日
```

`vendor` フォルダ以下にパッケージがフルパスで入ってしまうため階層が深くなりがちなのが「玉に瑕」だが，それ以外は特に問題はない。
あるいは `vendor` フォルダ以下のパッケージは `go get` の制約から外れているので，呼び出し側を

```go
import (
    "flag"
    "fmt"
    "os"
    "strconv"
    "time"

    "astrocalc/modjulian"
)
```

として以下のフォルダ構成にする手もある[^e]。

[^e]: パッケージのパスが変わるとテストが通らなくなる場合があるので注意。

```
C:\workspace\vdemo>tree /f .
C:\WORKSPACE\VDEMO
└─src
    └─julian-day
        │  julian-day.go
        │
        └─vendor
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


C:\workspace\vdemo>go install -v ./...
julian-day/vendor/astrocalc/modjulian
julian-day

C:\workspace\vdemo>bin\julian-day.exe 2015 1 1
2015-01-01 00:00:00 +0000 UTC
MJD = 57023日
```

注意が必要なのは， `go get` は git の submodule を上手く扱えないため， `vendor` フォルダ以下のパッケージを submodule として配置している場合はビルドに失敗することだ。
この場合は `-d` オプションで `go get` がビルドまで行わないようにし，手動で submodule の `init` と `update` を行う必要がある。

```
C:>go get -d project/...
C:>git submodule init
C:>git submodule update
C:>go install ./...
```

（「[Glide で Vendoring]({{< relref "vendoring-with-glide.md" >}})」に続く）

## ブックマーク

- [Go言語のDependency/Vendoringの問題と今後．gbあるいはGo1.5 | SOTA](http://deeeet.com/writing/2015/06/26/golang-dependency-vendoring/)
- [direnvで解決するGOPATHの3つの問題点 - None is None is None](http://doloopwhile.hatenablog.com/entry/2014/06/18/010449)
- [改めて、direnvを使いましょう！ - HDE BLOG](http://blog.hde.co.jp/entry/2015/02/27/182117)
- [さくら - homeにgolang, direnv とvirtualenvを入れて動かす - Qiita](http://qiita.com/aminamid/items/5a0e9461385c80d0c8a6)

- [Go 1.11 のリリースと「モジュール」機能の実験的サポート]({{< ref "/release/2018/09/go-1_11-ise-released.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[前回]: {{< relref "project-based-development.md" >}} "プロジェクト・ベースの開発環境をつくる"
[gb]: http://getgb.io/ "gb - A project based build tool for Go"
[direnv]: http://direnv.net/ "direnv - unclutter your .profile"
