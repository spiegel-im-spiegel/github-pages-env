+++
title = "Depm: Go 言語用モジュール依存関係可視化ツール"
date =  "2020-11-08T16:55:50+09:00"
description = " Go言語のパッケージまたはモジュール間の依存関係を可視化するコマンドラインツール"
image = "/images/attention/go-logo_blue.png"
tags  = [ "tools", "golang", "package", "module" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[spiegel-im-spiegel/depm][depm] は [Go]言語のパッケージまたはモジュール間の依存関係を可視化するコマンドラインツールである。
以下のように依存関係を図にすることも可能である。

{{< fig-img src="./output3.png" link="./output3.png" width="2534" >}}

以降で [spiegel-im-spiegel/depm][depm] について簡単に紹介する。

[![check vulns](https://github.com/spiegel-im-spiegel/depm/workflows/vulns/badge.svg)](https://github.com/spiegel-im-spiegel/depm/actions)
[![lint status](https://github.com/spiegel-im-spiegel/depm/workflows/lint/badge.svg)](https://github.com/spiegel-im-spiegel/depm/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/depm/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/spiegel-im-spiegel/depm.svg)](https://github.com/spiegel-im-spiegel/depm/releases/latest)

## 簡単な使い方

`-h` オプションで簡単な使い方が表示される。

```text
$ depm -h
Visualize depndency packages and modules.

Usage:
  depm [flags]
  depm [command]

Available Commands:
  help        Help about any command
  list        list modules
  module      analyze depndency modules
  package     analyze depndency packages
  version     print the version number

Flags:
      --cgo-enabled string   set CGO_ENABLED environment variable
      --debug                for debug
      --goarch string        set GOARCH environment variable
      --goos string          set GOOS environment variable
  -h, --help                 help for depm

Use "depm [command] --help" for more information about a command.
```

では，コマンドごとに見ていく。

### パッケージ間の依存関係を調べる

`depm package` コマンドはパッケージ単位で依存関係を調べて結果を出力する。

```text
$ depm package -h
analyze depndency packages.

Usage:
  depm package [flags] [package import path]

Aliases:
  package, pkg, p

Flags:
      --dot                 output by DOT language
      --dot-config string   config file for DOT language
  -h, --help                help for package
  -i, --include-internal    include internal packages
  -s, --include-standard    include standard Go library

Global Flags:
      --cgo-enabled string   set CGO_ENABLED environment variable
      --debug                for debug
      --goarch string        set GOARCH environment variable
      --goos string          set GOOS environment variable
```

`github.com/spiegel-im-spiegel/depm` パッケージを調べる場合は以下のコマンドラインとなる。

```text
$ depm package github.com/spiegel-im-spiegel/depm | jq .
[
  {
    "Package": {
      "ImportPath": "github.com/google/licenseclassifier",
      "Module": {
        "Path": "github.com/google/licenseclassifier",
        "Version": "v0.0.0-20210324205846-148b633b0534",
        "License": "Apache-2.0"
      }
    },
    "Deps": [
      {
        "ImportPath": "github.com/google/licenseclassifier/stringclassifier",
        "Module": {
          "Path": "github.com/google/licenseclassifier",
          "Version": "v0.0.0-20210324205846-148b633b0534",
          "License": "Apache-2.0"
        }
      },
      {
        "ImportPath": "github.com/google/licenseclassifier/stringclassifier/searchset",
        "Module": {
          "Path": "github.com/google/licenseclassifier",
          "Version": "v0.0.0-20210324205846-148b633b0534",
          "License": "Apache-2.0"
        }
      }
    ]
  },
...
```

出力は JSON 形式なので [jq] コマンド等で加工可能である。
JSON 出力の構造は以下の構造体 `nodeJSON` の配列で表される。

```go
type nodeJSON struct {
    Package *packageJSON
    Deps    []*packageJSON `json:",omitempty"`
}
type packageJSON struct {
    ImportPath string
    Internal   bool        `json:",omitempty"`
    CGO        bool        `json:",omitempty"`
    Unsafe     bool        `json:",omitempty"`
    Module     *moduleJSON `json:",omitempty"`
}
type moduleJSON struct {
    Path    string `json:",omitempty"`
    Version string `json:",omitempty"`
    License string `json:",omitempty"`
}
```

CGO に依存するパッケージは `CGO` フラグが `true` になる。
同様に [`unsafe`] を使っているパッケージでも `Unsafe` が `true` になる。
また `-i` オプションを付けてコマンドを起動すると internal パッケージも結果に含まれるが，この場合は `Internal` フラグが `true` になる。

たとえば有名な [`github.com/mattn/go-sqlite3`](https://github.com/mattn/go-sqlite3 "mattn/go-sqlite3: sqlite3 driver for go using database/sql") パッケージなら

```text
$ depm package "github.com/mattn/go-sqlite3" | jq .
[
  {
    "Package": {
      "ImportPath": "github.com/mattn/go-sqlite3",
      "CGO": true,
      "Unsafe": true,
      "Module": {
        "Path": "github.com/mattn/go-sqlite3",
        "Version": "v1.14.6"
      }
    }
  }
]
```

のようになる。

`--dot` オプションで [DOT] 言語形式の出力にもできるため [Graphviz] の dot コマンドを使って図に変換することができる。

```text
$ depm package --dot github.com/spiegel-im-spiegel/depm | dot -Tpng -o output1.png
```

実行結果はこんな感じ。

{{< fig-img src="./output1.png" title="output1.png" link="./output1.png" width="4578" >}}

図の node と edge の属性については [DOT 言語の仕様](http://www.graphviz.org/doc/info/attrs.html)に従って簡単な装飾ができる。
TOML 形式で

```toml
[node]
  fontname = "Inconsolata"
[edge]
  color = "red"
```

のような設定ファイルを作って

```text
$ depm package --dot --dot-config sample.toml "github.com/spiegel-im-spiegel/depm" | dot -Tpng -o output2.png
```

{{< fig-img src="./output2.png" title="output2.png" link="./output2.png" width="4259" >}}

のようにできる。

パッケージのインポート・パスを省略すると，カレント・ディレクトリのパッケージを走査する（`go list` コマンドで `all` を指定したときと同じ動作）。

### モジュール間の依存関係を調べる

`depm module` コマンドはモジュール単位で依存関係を調べて結果を出力する。

```text
$ depm module -h
analyze depndency modules.

Usage:
  depm module [flags] [package import path]

Aliases:
  module, mod, m

Flags:
  -u, --check-update        check updating module
      --dot                 output by DOT language
      --dot-config string   config file for DOT language
  -h, --help                help for module
  -i, --include-internal    include internal packages

Global Flags:
      --cgo-enabled string   set CGO_ENABLED environment variable
      --debug                for debug
      --goarch string        set GOARCH environment variable
      --goos string          set GOOS environment variable

$ depm module "github.com/spiegel-im-spiegel/depm" | jq .
[
  {
    "Module": {
      "Path": "github.com/google/licenseclassifier@v0.0.0-20210324205846-148b633b0534",
      "License": "Apache-2.0",
      "Packages": [
        "github.com/google/licenseclassifier",
        "github.com/google/licenseclassifier/stringclassifier",
        "github.com/google/licenseclassifier/stringclassifier/searchset",
        "github.com/google/licenseclassifier/stringclassifier/searchset/tokenizer"
      ]
    },
    "Deps": [
      {
        "Path": "github.com/sergi/go-diff@v1.0.0",
        "License": "MIT",
        "Packages": [
          "github.com/sergi/go-diff/diffmatchpatch"
        ]
      }
    ]
  },
...
```

出力は同じく JSON 形式で，以下の構造体 `nodeJSON` の配列で表される。

```go
type nodeJSON struct {
    Module *moduleJSON
    Deps   []*moduleJSON `json:",omitempty"`
}
type moduleJSON struct {
    Path     string
    Replace  string   `json:",omitempty"`
    Latest   string   `json:",omitempty"`
    Main     bool     `json:",omitempty"`
    CGO      bool     `json:",omitempty"`
    Unsafe   bool     `json:",omitempty"`
    License  string   `json:",omitempty"`
    Packages []string `json:",omitempty"`
}
```

CGO や [`unsafe`] パッケージに関するフラグは `depm package` コマンドのときと同様である。
`-u` オプションを付けると，新しいバージョンのモジュールがある場合に `Latest` に情報がセットされる。

`depm package` コマンドのときと同じく `-dot` オプションで [DOT] 言語形式の出力にできる。

```text
$ depm module --dot --dot-config sample.toml "github.com/spiegel-im-spiegel/depm" | dot -Tpng -o output3.png
```

{{< fig-img src="./output3.png" title="output3.png" link="./output3.png" width="2534" >}}

パッケージのインポート・パスを省略すると，カレント・ディレクトリのパッケージを走査する。
`go list -m` コマンドで `all` を指定したときと同じ動作だが，実際にはリンクしない形式的な依存モジュールは出力されない。
基本的にコード内の `import` に基づく依存関係のみ表示される。

### モジュールの列挙

`depm list` コマンドは `go list -m` コマンドと同じ形式でモジュールの列挙を行う。

```text
$ depm list -h
list modules, compatible 'go list -m' command

Usage:
  depm list [flags] [package import path]

Aliases:
  list, lst, l

Flags:
  -u, --check-update   check updating module
  -h, --help           help for list
  -j, --json           output by JSON format

Global Flags:
      --cgo-enabled string   set CGO_ENABLED environment variable
      --debug                for debug
      --goarch string        set GOARCH environment variable
      --goos string          set GOOS environment variable

$ depm list "github.com/spiegel-im-spiegel/depm"
github.com/BurntSushi/toml v0.3.1
github.com/emicklei/dot v0.15.0
github.com/google/licenseclassifier v0.0.0-20210324205846-148b633b0534
github.com/sergi/go-diff v1.0.0
github.com/spf13/cobra v1.1.3
github.com/spf13/pflag v1.0.5
github.com/spiegel-im-spiegel/depm v0.4.0
github.com/spiegel-im-spiegel/errs v1.0.2
github.com/spiegel-im-spiegel/gocli v0.10.4
golang.org/x/mod v0.3.0
golang.org/x/net v0.0.0-20201021035429-f5854403a974
golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4
golang.org/x/tools v0.1.0
golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
```

`go list -m` コマンドでは実際にはリンクしない形式上の依存モジュールまで表示してしまうが， `depm list` コマンドは基本的にコード内の `import` に基づく依存関係のみ表示されるため，最小限のリストに抑えられる。

`-u` オプションを付けると `go list -m -u` と同等の出力にできる。

```text
$ depm list -u "github.com/spiegel-im-spiegel/depm"
github.com/BurntSushi/toml v0.3.1
github.com/emicklei/dot v0.15.0
github.com/google/licenseclassifier v0.0.0-20210324205846-148b633b0534
github.com/sergi/go-diff v1.0.0 [v1.1.0]
github.com/spf13/cobra v1.1.3
github.com/spf13/pflag v1.0.5
github.com/spiegel-im-spiegel/depm v0.4.0
github.com/spiegel-im-spiegel/errs v1.0.2
github.com/spiegel-im-spiegel/gocli v0.10.4
golang.org/x/mod v0.3.0 [v0.4.2]
golang.org/x/net v0.0.0-20201021035429-f5854403a974 [v0.0.0-20210324205630-d1beb07c2056]
golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 [v0.0.0-20210324051608-47abb6519492]
golang.org/x/tools v0.1.0
golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
```

`-j` オプションを付けると `go list -m -json` と同等の JSON 出力にできる。

```text
$ depm list -j "github.com/spiegel-im-spiegel/depm"
{
    "Path": "github.com/BurntSushi/toml",
    "Version": "v0.3.1",
    "Time": "2018-08-15T10:47:33Z",
    "Indirect": true,
    "Dir": "/home/username/go/pkg/mod/github.com/!burnt!sushi/toml@v0.3.1",
    "GoMod": "/home/username/go/pkg/mod/cache/download/github.com/!burnt!sushi/toml/@v/v0.3.1.mod"
}
{
    "Path": "github.com/emicklei/dot",
    "Version": "v0.15.0",
    "Time": "2020-10-30T08:43:19Z",
    "Indirect": true,
    "Dir": "/home/username/go/pkg/mod/github.com/emicklei/dot@v0.15.0",
    "GoMod": "/home/username/go/pkg/mod/cache/download/github.com/emicklei/dot/@v/v0.15.0.mod",
    "GoVersion": "1.13"
}
...
```

`go list -m` コマンドに比べると遅くなりがちなのが欠点だが，他のツールと組み合わせることを考えると使い勝手はさほど悪くないと思う。

[Go]: https://golang.org/ "The Go Programming Language"
[depm]: https://github.com/spiegel-im-spiegel/depm "spiegel-im-spiegel/depm: Visualize depndency packages and modules"
[jq]: https://stedolan.github.io/jq/ "jq"
[DOT]: https://graphviz.gitlab.io/_pages/doc/info/lang.html "The DOT Language"
[Graphviz]: http://www.graphviz.org/ "Graphviz - Graph Visualization Software"
[`unsafe`]: https://golang.org/pkg/unsafe/ "unsafe - The Go Programming Language"

## 参考文献

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
