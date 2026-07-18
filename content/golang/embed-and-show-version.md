+++
title = "バージョンの埋め込みと表示"
date =  "2026-07-17T18:15:52+09:00"
description = "ビルド時に埋め込む方法とビルド情報から取得する方法の2つを紹介する。"
isCJKLanguage = true
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "module", "versioning" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

まずは，こんな感じにバージョン情報を表示するコードを考える。

```go
package main

var Version = "v1.2.3"

func main() {
    println(Version)
}
```

これを実行すれば `v1.2.3` と表示される。

実際の運用としてバージョンが上がるたびに変数 `Version` の内容を書き換える必要がある。
また Git でソースコード管理する場合はリリースバージョンに対してバージョンタグを付けることが多いので，それとの整合を保つ必要がある。
実はこれが割と煩雑な作業なのである。

そこで [Go] でビルドしたツールのバージョン表記について，以下の2つの方法を記しておく。

## ビルド時にバージョン文字列を埋め込む

これは割と昔からある方法で， `go build` コマンドの `-ldflags` オプションを使って変数に値を埋め込む。

```text
$ go build -ldflags="-X main.Version=v9.9.9" -o a.out

$ ./a.out
v9.9.9
```

Git でバージョンタグを打っている場合は

```text
$ go build -ldflags="-X main.Version=`git describe --tags`" -o a.out
```

などと Git コマンドの出力を埋め込む手もある（bash シェルの場合）。

もしビルドに [GoReleaser] を使っているのであれば `.goreleaser.yml` の中で

```yml
builds:
-
  ldflags: -X main.Version=v{{ .Version }}
```

のように書くことで対応できる。

[GoReleaser] はマルチプラットフォームのコマンドラインツールで，複数プラットフォーム向けの実行バイナリを一度に作ってくれるので重宝している。
また [GitHub Actions](https://github.com/features/actions "GitHub Actions") にも[対応][GoReleaser Action]していて，ビルドしたバイナリを GitHub Release に自動でアップロードし，リリースノートまで書いてくれる。
ホンマ助かる。

ただし，この方法は `go install` では使えないし，そもそもユーザに「`go build` 時に `-ldflags` でバージョンを埋め込んでね」とお願いするのもおかしな話である[^m1]。

[^m1]: プロダクトによっては `Makefile` やビルド用のスクリプトを用意してるものもある。まぁ，大抵はそれなりの環境を用意しないとビルドできないことが多いのだが。

## ビルド情報を表示する

これは [Go] 1.18 から追加されたらしいのだが [`runtime/debug`][`debug`] パッケージの `ReadBuildInfo()` 関数を使うと，ビルド時の情報を取得できる。
たとえば，こんな感じ。

```go {hl_lines=[11]}
package main

import "runtime/debug"

var Version = ""

func replaceVersion(v string) string {
    if v != "" {
        return v
    }
    if info, ok := debug.ReadBuildInfo(); ok {
        if info.Main.Version != "" {
            return info.Main.Version
        }
        var revision, dirty string
        for _, v := range info.Settings {
            switch v.Key {
            case "vcs.revision":
                revision = v.Value
            case "vcs.modified":
                if v.Value == "true" {
                    dirty = "+dirty"
                }
            }
        }
        if revision != "" {
            return revision + dirty
        }
    }
    return "(version not set)"
}

func main() {
    println(replaceVersion(Version))
}
```

[`debug`]`.ReadBuildInfo()` 関数の返り値である `BuildInfo` 構造体の内容は以下の通り。

```go
type BuildInfo struct {
    // GoVersion is the version of the Go toolchain that built the binary
    // (for example, "go1.19.2").
    GoVersion string `json:",omitempty"`

    // Path is the package path of the main package for the binary
    // (for example, "golang.org/x/tools/cmd/stringer").
    Path string `json:",omitempty"`

    // Main describes the module that contains the main package for the binary.
    Main Module `json:""`

    // Deps describes all the dependency modules, both direct and indirect,
    // that contributed packages to the build of this binary.
    Deps []*Module `json:",omitempty"`

    // Settings describes the build settings used to build the binary.
    Settings []BuildSetting `json:",omitempty"`
}
```

`Module` 構造体はこんな感じかな。

```go
type Module struct {
    Path    string  `json:",omitempty"` // module path
    Version string  `json:",omitempty"` // module version
    Sum     string  `json:",omitempty"` // checksum
    Replace *Module `json:",omitempty"` // replaced by this module
}
```

つまり `go.mod` でモジュールとして定義・管理されている場合は `BuildInfo.Main.Version` にバージョンが入っている。

作成した `replaceVersion()` 関数を含むコードを Git で commit し `v0.1.0` のバージョンタグを打った状態でビルド&実行すると以下のように表示された。

```text
$ go build -o a.out && ./a.out
v0.1.0
```

ちなみに `go.mod` でモジュールとして定義されている場合でも，バージョン管理されていない場合は `(devel)` という文字列が返ってくるようだ。

```text
$ go build -o a.out && ./a.out
(devel)
```

`replaceVersion()` 関数では `BuildInfo.Main.Version` で値を取れない場合には `BuildInfo.Settings` 配列から情報を取得している。
`BuildSetting` 構造体は以下の通り key-value 形式になっている。

```go
type BuildSetting struct {
    // Key and Value describe the build setting.
    // Key must not contain an equals sign, space, tab, or newline.
    Key string `json:",omitempty"`
    // Value must not contain newlines ('\n').
    Value string `json:",omitempty"`
}
```

このうち `Key` には以下の値が入ってるらしい。

| Key の値 | Value の内容 |
|---|---|
| `-buildmode` | the buildmode flag used (typically `"exe"`) |
| `-compiler` | the compiler toolchain flag used (typically `"gc"`) |
| `CGO_ENABLED` | the effective `CGO_ENABLED` environment variable |
| `CGO_CFLAGS` | the effective `CGO_CFLAGS` environment variable |
| `CGO_CPPFLAGS` | the effective `CGO_CPPFLAGS` environment variable |
| `CGO_CXXFLAGS` | the effective `CGO_CXXFLAGS` environment variable |
| `CGO_LDFLAGS` | the effective `CGO_LDFLAGS` environment variable |
| `DefaultGODEBUG` | the effective `GODEBUG` settings |
| `GOARCH` | the architecture target |
| `GOAMD64/GOARM/GO386/etc` | the architecture feature level for `GOARCH` |
| `GOOS` | the operating system target |
| `GOFIPS140` | the frozen FIPS 140-3 module version, if any |
| `vcs` | the version control system for the source tree where the build ran |
| `vcs.revision` | the revision identifier for the current commit or checkout |
| `vcs.time` | the modification time associated with `vcs.revision`, in RFC3339 format |
| `vcs.modified` | `"true"` or `"false"` indicating whether the source tree had local modifications |

`replaceVersion()` 関数では VCS (Version Control System) のリビジョン情報（`vcs.revision` および `vcs.modified`）を使ってバージョンを表示している。
ここまで至れり尽くせりなら，大抵のパターンで使えるだろう。

### go run ではビルド情報がセットされない

以上は `go build` あるいは `go install` でビルドした場合の話だが `go run` で即時実行した場合はビルド情報がセットされないようだ。
先程のコード例の場合は

```text
$ go run main.go
(version not set)
```

と表示されビルド情報からバージョンを取得できなかったことが分かる。
まぁ `go run` による実行でバージョンを気にする人はいないと思うので，これはこれで問題ないだろう。

## ブックマーク

- [Go 1.18 集中連載 実行ファイルのメタデータに関するアップデート(コミットID追加等) | フューチャー技術ブログ](https://future-architect.github.io/articles/20220217a/)
- [Go で作成した CLI ツールにバージョンを埋め込む方法](https://zenn.dev/otakakot/articles/d4448f208c91f7)

[Go]: https://go.dev/
[GoReleaser]: https://goreleaser.com/ "GoReleaser"
[GoReleaser Action]: https://github.com/marketplace/actions/goreleaser-action "GoReleaser Action - GitHub Marketplace"
[`debug`]: https://pkg.go.dev/runtime/debug "debug package - runtime/debug - Go Packages"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B0CFL1DK8Q" %}} <!-- Go言語 100Tips -->
{{% review-paapi "B0DNYMMBBQ" %}} <!-- Go言語で学ぶ並行プログラミング -->
{{% linkcard "51769bc6eecda48305121bbc56350f35b5599135" %}} <!-- 実用 Go 言語 第2版 -->
