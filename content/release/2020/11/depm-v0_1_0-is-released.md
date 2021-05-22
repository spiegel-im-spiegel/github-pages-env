+++
title = "Go パッケージ／モジュールの依存関係可視化ツールを作ってみた"
date =  "2020-11-01T20:47:14+09:00"
description = "以前自作したツールがイマイチで，他の方が作ったツールも微妙だったので，最初から作り直してみた。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "tools", "golang", "package", "module" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

ついカッとなってやった。
反省はしない。

この手のツールって誰もが一度は書いてみたくなると思うけど，以前[自作したツール][ggm]がイマイチで，他の方が作ったツールも微妙だったので，最初から作り直してみた。

- [spiegel-im-spiegel/depm: Visualize depndency packages and modules](https://github.com/spiegel-im-spiegel/depm)

[前のツール][ggm] では `go mod graph` コマンドの出力をパイプで繋いで処理してたけど，今回はツール内で `go list` コマンドを呼び出している。
したがって [Go] コンパイラがインストール済みであることが前提条件となる。

`-h` オプションで簡単なヘルプを表示する。

```text
$ depm -h
Visualize depndency packages and modules.

Usage:
  depm [flags]
  depm [command]

Available Commands:
  help        Help about any command
  module      analyze depndency modules
  package     analyze depndency packages
  version     print the version number

Flags:
      --cgo-enabled string   set CGO_ENABLED environment variable
      --debug                for debug
      --dot                  output by DOT language
      --dot-config string    config file for DOT language
      --goarch string        set GOARCH environment variable
      --goos string          set GOOS environment variable
  -h, --help                 help for depm

Use "depm [command] --help" for more information about a command.
```

`depm module` コマンドならこんな感じに使える。

```text
$ depm module github.com/spiegel-im-spiegel/depm | jq .
[
  {
    "Module": {
      "Path": "github.com/spf13/cobra@v1.1.1",
      "Packages": [
        "github.com/spf13/cobra"
      ]
    },
    "Deps": [
      {
        "Path": "github.com/spf13/pflag@v1.0.5",
        "Packages": [
          "github.com/spf13/pflag"
        ]
      }
    ]
  },
  {
    "Module": {
      "Path": "github.com/spiegel-im-spiegel/depm@v0.1.0",
      "Packages": [
        "github.com/spiegel-im-spiegel/depm",
        "github.com/spiegel-im-spiegel/depm/dependency",
        "github.com/spiegel-im-spiegel/depm/dependency/modjson",
        "github.com/spiegel-im-spiegel/depm/dependency/pkgjson",
        "github.com/spiegel-im-spiegel/depm/dotenc",
        "github.com/spiegel-im-spiegel/depm/ecode",
        "github.com/spiegel-im-spiegel/depm/facade",
        "github.com/spiegel-im-spiegel/depm/golist",
        "github.com/spiegel-im-spiegel/depm/modules",
        "github.com/spiegel-im-spiegel/depm/packages"
      ]
    },
    "Deps": [
      {
        "Path": "github.com/spiegel-im-spiegel/gocli@v0.10.3",
        "Packages": [
          "github.com/spiegel-im-spiegel/gocli/exitcode",
          "github.com/spiegel-im-spiegel/gocli/rwi"
        ]
      },
      {
        "Path": "github.com/spf13/cobra@v1.1.1",
        "Packages": [
          "github.com/spf13/cobra"
        ]
      },
      {
        "Path": "github.com/spiegel-im-spiegel/errs@v1.0.2",
        "Packages": [
          "github.com/spiegel-im-spiegel/errs"
        ]
      },
      {
        "Path": "golang.org/x/tools@v0.0.0-20201031021630-582c62ec74d0",
        "Packages": [
          "golang.org/x/tools/go/ast/astutil",
          "golang.org/x/tools/imports",
          "golang.org/x/tools/internal/event",
          "golang.org/x/tools/internal/event/core",
          "golang.org/x/tools/internal/event/keys",
          "golang.org/x/tools/internal/event/label",
          "golang.org/x/tools/internal/fastwalk",
          "golang.org/x/tools/internal/gocommand",
          "golang.org/x/tools/internal/gopathwalk",
          "golang.org/x/tools/internal/imports"
        ]
      },
      {
        "Path": "github.com/BurntSushi/toml@v0.3.1",
        "Packages": [
          "github.com/BurntSushi/toml"
        ]
      },
      {
        "Path": "github.com/emicklei/dot@v0.15.0",
        "Packages": [
          "github.com/emicklei/dot"
        ]
      }
    ]
  },
  {
    "Module": {
      "Path": "golang.org/x/mod@v0.3.0",
      "Packages": [
        "golang.org/x/mod/module",
        "golang.org/x/mod/semver"
      ]
    },
    "Deps": [
      {
        "Path": "golang.org/x/xerrors@v0.0.0-20200804184101-5ec99f83aff1",
        "Packages": [
          "golang.org/x/xerrors",
          "golang.org/x/xerrors/internal"
        ]
      }
    ]
  },
  {
    "Module": {
      "Path": "golang.org/x/tools@v0.0.0-20201031021630-582c62ec74d0",
      "Packages": [
        "golang.org/x/tools/go/ast/astutil",
        "golang.org/x/tools/imports",
        "golang.org/x/tools/internal/event",
        "golang.org/x/tools/internal/event/core",
        "golang.org/x/tools/internal/event/keys",
        "golang.org/x/tools/internal/event/label",
        "golang.org/x/tools/internal/fastwalk",
        "golang.org/x/tools/internal/gocommand",
        "golang.org/x/tools/internal/gopathwalk",
        "golang.org/x/tools/internal/imports"
      ]
    },
    "Deps": [
      {
        "Path": "golang.org/x/mod@v0.3.0",
        "Packages": [
          "golang.org/x/mod/module",
          "golang.org/x/mod/semver"
        ]
      }
    ]
  }
]
```

出力は JSON 形式で。
パッケージ（モジュール）パスを省略するとカレント・ディレクトリを調べる[^l1]。

[^l1]: パッケージ（モジュール）パスを省略した場合は，ツール内部で `go list` コマンドに `all` を渡している。

`--dot` オプションを付けると [DOT] 言語形式で出力するので，そのまま `dot` コマンドに渡して関連図を作成できる。

```text
$ depm module --dot github.com/spiegel-im-spiegel/depm | dot -Tpng -o output.png
```

出力はこんな感じ。

{{< fig-img src="./output.png" link="./output.png" title="output.png" width="2818" >}}

TOML 形式で以下のような設定ファイルを作れば（[DOT] 言語の仕様にしたがって）見た目を多少変えることができる。

```toml
[node]
  fontname = "Inconsolata"
[edge]
  color = "red"
```

これで

```text
$ depm module --dot --dot-config dot-config.toml github.com/spiegel-im-spiegel/depm | dot -Tpng -o output2.png
```

とかすれば

{{< fig-img src="./output2.png" link="./output2.png" title="output2.png" width="2818" >}}

のような見た目にできる。
あと `-c` オプションでモジュールの最新バージョンの取得もできたりする。

`depm package` コマンドにするとモジュール単位ではなくパッケージ単位で依存関係を整理する。

```text
$ depm package github.com/spiegel-im-spiegel/depm
```

結構スゴい出力になるので，結果は割愛する（笑）
`depm package` コマンドに `-s` や `-i` オプションを付けると [Go] の標準ライブラリや internal パッケージも対象になるので，本当にワケワカラン出力になる。

{{< fig-img src="./output3.png" link="./output3.png" title="output3.png" width="9282" >}}

勢いで書いてろくにテストもしないでリリースしたが，あとはノンビリ手を入れていこう。

[Go]: https://golang.org/ "The Go Programming Language"
[graphviz]: https://www.graphviz.org/ "Graphviz - Graph Visualization Software"
[DOT]: https://graphviz.gitlab.io/_pages/doc/info/lang.html "The DOT Language"
[github.com/emicklei/dot]: https://github.com/emicklei/dot "emicklei/dot: Go package for writing descriptions using the Graphviz DOT language"
[ggm]: {{< ref "/release/2019/05/ggm.md" >}} "Go モジュールの依存関係を可視化するツールを作った"


## 参考図書

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
