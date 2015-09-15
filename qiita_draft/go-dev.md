# そろそろ真面目に Golang 開発環境について考える — GOPATH 汚染問題

「[はじめての Go 言語](http://qiita.com/spiegel-im-spiegel/items/98d49ac456485b007a15#%E3%81%AF%E3%81%98%E3%82%81%E3%81%A6%E3%81%AE-go-%E8%A8%80%E8%AA%9E-on-windows)」シリーズも終わった（終わらせた）ので，そろそろ真面目に開発環境を作って色々やってみたいと思う。まずは環境を整えるところから

## GOPATH 汚染問題

`go get` コマンドはとても強力な機能で，私のように Windows と UNIX 系環境の間を渡り歩いてる身としては， curl や make などの環境依存のツールを使わずに， `go get` コマンドだけで repository の fetch から build/install まで出来てしまうのは非常にありがたい（もちろん make 等を使って細かいコントロールをすることもできる）。

ひとつのマシン（仮想マシンを含めて）でひとつのプロジェクトを管理していくのならこれで充分だが，残念ながら，仕事でそんな状況はあまりない。プロジェクトごとに仮想マシンを用意できるのはまだマシな方で，そんなことすら出来ないぎりぎりのスペックで複数のプロジェクトの面倒を同時に見ていく事のほうが多い（Google 様とは違うのだよ）。で，困ったことに `GOPATH` 環境変数は複数のプロジェクト管理を想定していない。

`GOPATH` 環境変数には複数のパスをセットできる。 Windows 環境ならこんな感じ。

```shell
SET GOPATH=C:\golib;C:\workspace\project1;C:\workspace\project1;...
```

しかし，これらのパスが全て有効になるのは Go コンパイラが外部パッケージを探す場合であり， `go get` コマンドで repogitory を fetch する場合には `GOPATH` で指定する最初のパス（上述の例なら `C:\golib`）と決められている。これでは折角プロジェクトごとにフォルダを分けても，外部パッケージはひとつのフォルダに集約されてしまうので管理が煩雑になってしまう。

### 【対策1】 プロジェクトごとに GOPATH を設定し直す

この問題に対する一番安直な答えは「プロジェクトごとに GOPATH を設定し直す」である。例えば「[その9](http://qiita.com/spiegel-im-spiegel/items/ef15a48542e043b32c99)」で紹介した [gb] を build する場合は以下のようにする。

```shell
C:>SET GOPATH=C:\workspace\gb

C:>go get -v github.com/constabulary/gb/...
github.com/constabulary/gb (download)
github.com/constabulary/gb
github.com/constabulary/gb/cmd
github.com/constabulary/gb/vendor
github.com/constabulary/gb/cmd/gb
github.com/constabulary/gb/cmd/gb-vendor

C:>cd C:\workspace\gb

C:\workspace\gb>tree /f
C:.
├─bin
│      gb-vendor.exe
│      gb.exe
│
├─pkg
│  └─windows_amd64
│      └─github.com
│          └─constabulary
│              └─gb
│
└─src
    └─github.com
        └─constabulary
            └─gb
                ├─cmd
                │  ├─gb
                │  └─gb-vendor
                ├─testdata
                └─vendor
                    └─_testdata
```

（`src` および `pkg` フォルダ以下は一部省略している）

あとは `bin` フォルダに `PATH` を通すか， `PATH` の通ってるフォルダに実行ファイルをコピーすればよい。実行履歴はバッチファイル（または shell スクリプト）に保存しておけばいつでも復元できる。

確かに毎回環境をセットアップしないといけないのは面倒だが，プロジェクト管理のためのツールも必要なく， Go コンパイラの標準機能のみで管理できる。標準機能のみで管理できるというのは結構重要で，たとえば CI ツールを使っている場合は，設定を単純にできるので管理しやすいといえる。

### 【対策2】 プロジェクト・ベースの管理ツールを使う

もうひとつは [gb] のようなプロジェクト・ベースでコード管理のできるツールを使う方法である。例として「[その10](http://qiita.com/spiegel-im-spiegel/items/5d2878596360af8dd753)」で紹介した [tcnksm/gcli] の build 環境を [gb] を使って構築してみる。

```shell
C:> cd C:\workspace\gcli

C:\workspace\gcli>git clone https://github.com/tcnksm/gcli.git src\github.com/tcnksm/gcli
Cloning into 'src\github.com/tcnksm/gcli'...
remote: Counting objects: 766, done.
remote: Total 766 (delta 0), reused 0 (delta 0), pack-reused 766Receiving objects:  90% (690/766), 2.11 MiB | 828.00 KiB/s
Receiving objects: 100% (766/766), 2.50 MiB | 828.00 KiB/s, done.
Resolving deltas: 100% (415/415), done.
Checking connectivity... done.

C:\workspace\gcli>gb vendor fetch github.com/mitchellh/cli
fetching recursive dependency golang.org/x/crypto/ssh/terminal

C:\workspace\gcli>gb vendor fetch github.com/olekukonko/tablewriter

C:\workspace\gcli>gb vendor fetch github.com/tcnksm/go-gitconfig

C:\workspace\gcli>gb vendor fetch github.com/tcnksm/go-latest
fetching recursive dependency github.com/google/go-github/github
fetching recursive dependency github.com/google/go-querystring/query
fetching recursive dependency github.com/hashicorp/go-version
fetching recursive dependency golang.org/x/net/html

C:\workspace\gcli>pushd src\github.com\tcnksm\gcli\skeleton

C:\workspace\gcli\src\github.com\tcnksm\gcli\skeleton>go-bindata -pkg="skeleton" resource/...

C:\workspace\gcli\src\github.com\tcnksm\gcli\skeleton>popd

C:\workspace\gcli>gb build
github.com/tcnksm/gcli/helper
github.com/tcnksm/go-gitconfig
github.com/google/go-querystring/query
golang.org/x/crypto/ssh/terminal
github.com/hashicorp/go-version
golang.org/x/net/html/atom
github.com/olekukonko/tablewriter
github.com/google/go-github/github
github.com/tcnksm/gcli/skeleton
github.com/mitchellh/cli
golang.org/x/net/html
github.com/tcnksm/go-latest
github.com/tcnksm/gcli/command
github.com/tcnksm/gcli

C:\workspace\gcli>bin\gcli.exe version
[0;0mgcli version v0.2.0[0m
[0;31m
Your versin of gcli is out of date! The latest version is 0.2.1.[0m

C:\workspace\gcli>
C:.
├─bin
│      gcli.exe
│
├─pkg
│  └─windows
│      └─amd64
│          ├─github.com
│          │  ├─google
│          │  │  ├─go-github
│          │  │  └─go-querystring
│          │  ├─hashicorp
│          │  ├─mitchellh
│          │  ├─olekukonko
│          │  └─tcnksm
│          │      └─gcli
│          └─golang.org
│              └─x
│                  ├─crypto
│                  │  └─ssh
│                  └─net
│                      └─html
├─src
│  └─github.com
│      └─tcnksm
│          └─gcli
│              ├─command
│              ├─helper
│              ├─skeleton
│              │  └─resource
│              └─tests
└─vendor
    │  manifest
    │  
    └─src
        ├─github.com
        │  ├─google
        │  │  ├─go-github
        │  │  │  └─github
        │  │  └─go-querystring
        │  │      └─query
        │  ├─hashicorp
        │  │  └─go-version
        │  ├─mitchellh
        │  │  └─cli
        │  ├─olekukonko
        │  │  └─tablewriter
        │  │      └─csv2table
        │  └─tcnksm
        │      ├─go-gitconfig
        │      └─go-latest
        │          └─latest
        └─golang.org
            └─x
                ├─crypto
                │  └─ssh
                │      └─terminal
                └─net
                    └─html
                        ├─atom
                        ├─charset
                        └─testdata
```

（`src` および `pkg` フォルダ以下は一部省略している）

[gb] では外部パッケージは `gb vendor fetch` コマンドで導入するが repository を clone するのではなく単にコピーするだけだ。外部パッケージの情報は `vendor/manifest` ファイルに格納されている。

```json:vendor/manifest
{
	"version": 0,
	"dependencies": [
		{
			"importpath": "github.com/google/go-github/github",
			"repository": "https://github.com/google/go-github",
			"revision": "7277108aa3e8823e0e028f6c74aea2f4ce4a1b5a",
			"branch": "master",
			"path": "/github"
		},
		{
			"importpath": "github.com/google/go-querystring/query",
			"repository": "https://github.com/google/go-querystring",
			"revision": "547ef5ac979778feb2f760cdb5f4eae1a2207b86",
			"branch": "master",
			"path": "/query"
		},
		{
			"importpath": "github.com/hashicorp/go-version",
			"repository": "https://github.com/hashicorp/go-version",
			"revision": "999359b6b7a041ce16e695d51e92145b83f01087",
			"branch": "master"
		},
		{
			"importpath": "github.com/mitchellh/cli",
			"repository": "https://github.com/mitchellh/cli",
			"revision": "8102d0ed5ea2709ade1243798785888175f6e415",
			"branch": "master"
		},
		{
			"importpath": "github.com/olekukonko/tablewriter",
			"repository": "https://github.com/olekukonko/tablewriter",
			"revision": "b9346ac189c55dd419f85c7ad2cd56f810bf19d6",
			"branch": "master"
		},
		{
			"importpath": "github.com/tcnksm/go-gitconfig",
			"repository": "https://github.com/tcnksm/go-gitconfig",
			"revision": "6411ba19847f20afe47f603328d97aaeca6def6f",
			"branch": "master"
		},
		{
			"importpath": "github.com/tcnksm/go-latest",
			"repository": "https://github.com/tcnksm/go-latest",
			"revision": "ef81df8e23895f6e86f9bdfea0576b9c17b9f1f4",
			"branch": "master"
		},
		{
			"importpath": "golang.org/x/crypto/ssh/terminal",
			"repository": "https://go.googlesource.com/crypto",
			"revision": "81bf7719a6b7ce9b665598222362b50122dfc13b",
			"branch": "master",
			"path": "/ssh/terminal"
		},
		{
			"importpath": "golang.org/x/net/html",
			"repository": "https://go.googlesource.com/net",
			"revision": "7654728e381988afd88e58cabfd6363a5ea91810",
			"branch": "master",
			"path": "/html"
		}
	]
}
```

つまり [gb] で作った開発環境はフォルダごと開発メンバに配布・同期することが可能になる。（[gb] で作った開発環境を git 等で管理する際に， `src` フォルダ内のファイルが別の repository で管理されている場合は，上述のように単純に clone するのではなく submodule として追加する必要がある）

[gb] の欠点は `go test` が使えないことである。一応 `gb test` というのがあって，ほぼ `go test` と互換なのだが，特殊なフォルダ構成にしている場合はテスト結果が変わる場合がある。このため CI ツールを使う際には注意が必要となる。

### 【対策3】 Go 1.5 の vendoring 機能を使う

これについては調査中。よい記事があればあとでリンクしておく。とりあえず，以下を参考に。

- [Go 1.5 Release Notes - The Go Programming Language](https://golang.org/doc/go1.5)
    - [Go 1.5 Vendor Experiment](https://docs.google.com/document/d/1Bz5-UB7g2uPBdOx-rw5t9MxJwkfpx90cqG9AFL0JAYo/edit)

*追記：*

記事を書きました： [Internal Packages と Vendoring](http://qiita.com/spiegel-im-spiegel/items/baa3671c7e1b8a6594a9)

### じゃあ GOPATH には何を入れればいいの？

godoc や golint などの標準的なツールやプロジェクト・メンバ間で共通に使うツールは `go get` で導入したほうがいいだろう。

```shell
C:>go get -v golang.org/x/tools/cmd/godoc
golang.org/x/tools/blog/atom
golang.org/x/tools/present
golang.org/x/tools/go/ast/astutil
golang.org/x/tools/go/types/typeutil
golang.org/x/tools/go/buildutil
golang.org/x/tools/container/intsets
golang.org/x/tools/blog
golang.org/x/tools/go/ssa
golang.org/x/tools/go/loader
golang.org/x/tools/godoc/vfs
golang.org/x/tools/godoc/redirect
golang.org/x/tools/godoc/static
golang.org/x/tools/go/callgraph
golang.org/x/tools/go/ssa/ssautil
golang.org/x/tools/godoc/util
golang.org/x/tools/godoc/vfs/httpfs
golang.org/x/tools/godoc/vfs/gatefs
golang.org/x/tools/go/pointer
golang.org/x/tools/godoc/vfs/mapfs
golang.org/x/tools/godoc/vfs/zipfs
golang.org/x/tools/playground
golang.org/x/tools/godoc/analysis
golang.org/x/tools/godoc
golang.org/x/tools/cmd/godoc

C:>go get -v golang.org/x/tools/cmd/vet
Fetching https://golang.org/x/tools/cmd/vet?go-get=1
Parsing meta tags from https://golang.org/x/tools/cmd/vet?go-get=1 (status code 200)
get "golang.org/x/tools/cmd/vet": found meta tag main.metaImport{Prefix:"golang.org/x/tools", VCS:"git", RepoRoot:"https://go.googlesource.com/tools"} at https://golang.org/x/tools/cmd/vet?go-get=1
get "golang.org/x/tools/cmd/vet": verifying non-authoritative meta tag
Fetching https://golang.org/x/tools?go-get=1
Parsing meta tags from https://golang.org/x/tools?go-get=1 (status code 200)
golang.org/x/tools (download)
golang.org/x/tools/go/exact
golang.org/x/tools/cmd/vet/whitelist
golang.org/x/tools/go/types
golang.org/x/tools/go/gcimporter
golang.org/x/tools/cmd/vet

C:>go get -v github.com/golang/lint/golint
github.com/golang/lint (download)
github.com/golang/lint
github.com/golang/lint/golint

C:>go get -v github.com/jteeuwen/go-bindata/...
github.com/jteeuwen/go-bindata (download)
github.com/jteeuwen/go-bindata
github.com/jteeuwen/go-bindata/go-bindata
```

ただし Windows を開発プラットフォームにしている場合は，ツールの管理担当者を決めて，管理担当者が配布するバイナリを使うほうが安全かもしれない。独りで作業している場合はどうとでもなるが，実際には業務システムを独りで開発するケースのほうが稀だと思うので，メンバ間で同じ環境になるように調整していくことが重要である。

では，楽しくお仕事しましょう。

## ブックマーク

- [Go言語のDependency/Vendoringの問題と今後．gbあるいはGo1.5 | SOTA](http://deeeet.com/writing/2015/06/26/golang-dependency-vendoring/)

[gb]: http://getgb.io/ "gb - A project based build tool for Go"
[tcnksm/gcli]: https://github.com/tcnksm/gcli "The easy way to build Golang command-line application."
