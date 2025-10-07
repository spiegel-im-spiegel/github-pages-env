+++
title = "Go 1.25 のリリース"
date =  "2025-08-13T10:46:18+09:00"
description = "ついに encoding/json/v2 パッケージが実装された。ただし実験的実装"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

予定通り8月に [Go] 1.25 がリリースされた。

- [Go 1.25.0 is released](https://groups.google.com/g/golang-announce/c/BVrdugXW05c)
- [Go 1.25 Release Notes - The Go Programming Language](https://go.dev/doc/go1.25)
- [Go 1.25 is released - The Go Programming Language](https://go.dev/blog/go1.25)

[リリースノート](https://go.dev/doc/go1.25 "Go 1.25 Release Notes - The Go Programming Language")をざっと眺めて個人的に気になった部分をピックアップする。

## go.mod に ignore ディレクティブを追加

[`go.mod`](https://go.dev/ref/mod "Go Modules Reference - The Go Programming Language") ファイルの記述に `ignore` ディレクティブが追加された。
これを使うことで [Go] コマンドの解析対象から除外するディレクトリを指定できる。
たとえば

```text
module example.com/my/module

go 1.25

ignore ./node_modules
ignore (
    static
    content/html
    ./third_party/javascript
)
```

みたいな感じ。
[Go] のコードとは関係ないディレクトリに大量のファイルがある場合は `ignore` ディレクティブで対象ディレクトリを指定することによって [Go] コマンドの解析効率を上げる，というのが目的らしい。

## nil ポインタに関するバグの修正

たとえば

```go
package main

import "os"

func main() {
    f, err := os.Open("nonExistentFile")
    name := f.Name()
    if err != nil {
        return
    }
    println(name)
}
```

というコードで [`os`]`.Open()` 関数で失敗した場合，本来は `f` に `nil` が入ってるはずなので `name := f.Name()` の部分で panic になる筈なのだが，1.21 以降 1.24 以前のバージョンでは `err != nil` のチェックが行われるまで評価が遅延されてしまい panic にならずに通ってしまっていたらしい。
上述のコードの場合は実害なさげだが，もっと複雑なコードでは問題になるかもしれない。

というわけで 1.25 では上述のコードの場合，正しく panic を吐くよう修正された。
併せて，エラーチェックに関しては

```go {hl_lines=[7]}
package main

import "os"

func main() {
    f, err := os.Open("nonExistentFile")
    if err != nil {
        return
    }
    name := f.Name()
    println(name)
}
```

のように，可能な限り `error` を返す関数の直後に `err != nil` のチェックを行うよう推奨している。

## encoding/json/v2 パッケージの実験的実装

ついに [`encoding/json/v2`] パッケージが実装された。

ただし，今のところは experimental implementation であり，今後のバージョンとの互換性は保証されていないため，取り扱いには注意が必要である。
新しいパッケージを有効にするには，ビルド時に環境変数で `GOEXPERIMENT=jsonv2` と指定する必要がある。
また，この環境変数をセットすることで従来の [`encoding/json`] パッケージの内部実装も [`v2`][`encoding/json/v2`] ベースに切り換わるらしい。

というわけで，今の [`encoding/json`] パッケージで問題がないのであれば [`v2`][`encoding/json/v2`] パッケージへの移行は今のところは保留しておいたほうがいいかもしれない。
ただし [`encoding/json`] パッケージの最新ドキュメントには

{{< fig-quote type="markdown" title="json package - encoding/json - Go Packages" link="https://pkg.go.dev/encoding/json" >}}
For historical reasons, the default behavior of v1 [`encoding/json`](https://pkg.go.dev/encoding/json) unfortunately operates with less secure defaults. New usages of JSON in Go are encouraged to use [encoding/json/v2](https://pkg.go.dev/encoding/json/v2) instead.
{{< /fig-quote >}}

などと記されており，将来的には [`v2`][`encoding/json/v2`] パッケージへの移行が必要になるだろう。

## その他

- コンテナ対応の `GOMAXPROCS`
- 新しい実験的な GC
  - ビルド時に環境変数に `GOEXPERIMENT=greenteagc` を指定
- [`runtime/trace.FlightRecorder`](https://go.dev/pkg/runtime/trace#FlightRecorder) API
- [`crypto`](https://pkg.go.dev/crypto) パッケージの改良
- Windows で [`os`]`.NewFile()` の改良。 non blocking I/O のサポート

他にも様々な変更・改良がある。

## Go 1.25.2 の導入 【2025-10-08 更新】 {#update}

{{< div-box type="markdown" >}}
[Go] 1.25.2 がリリースされた。
`golang.org/x/net` パッケージの脆弱性報告も併せて行われている。

- [[security] Go 1.25.2 and Go 1.24.8 are released](https://groups.google.com/g/golang-announce/c/4Emdl2iQ_bI)
- [[security] Vulnerabilities in golang.org/x/net](https://groups.google.com/g/golang-announce/c/jnQcOYpiR2c)

CVE ID ベースで10件の脆弱性修正がある（1.25.0 からの累積で11件）。

[Go]: https://go.dev/
{{< /div-box >}}

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.25.2.linux-amd64.tar.gz`](https://go.dev/dl/go1.25.2.linux-amd64.tar.gz)）を取ってきてインストールすることを推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -LO "https://go.dev/dl/go1.25.2.linux-amd64.tar.gz"
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.25.2.linux-amd64.tar.gz
$ sudo mv go go1.25.2
$ sudo ln -s go1.25.2 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.25.2 linux/amd64
```

Windows はインストールパッケージを取ってきて直接インストールする。
[Scoop] などのパッケージ管理ツール経由でも OK

複数バージョンの [Go] コンパイラを扱いたい場合は

```text
$ go install golang.org/dl/25.2@latest
$ go1.25.2 download
$ go1.25.2 version
go version go1.25.2 linux/amd64
```

てな感じに導入できる。

## ブックマーク

- [Container-aware GOMAXPROCS - The Go Programming Language](https://go.dev/blog/container-aware-gomaxprocs)
- [Testing Time (and other asyncronicities) - The Go Programming Language](https://go.dev/blog/testing-time)
- [A new experimental Go API for JSON - The Go Programming Language](https://go.dev/blog/jsonv2-exp)
- [Flight Recorder in Go 1.25 - The Go Programming Language](https://go.dev/blog/flight-recorder)

- [[security] Go 1.25.1 and Go 1.24.7 are released](https://groups.google.com/g/golang-announce/c/PtW9VW21NPs)

### 日本語情報

- [Go1.25 New Features](https://zenn.dev/koya_iwamura/articles/ea2cf191cdcb2a)
- [Go 1.25、GOMAXPROCS自動設定の恩恵](https://zenn.dev/drsprime/articles/944e7a6c3e990f)
- [Go 1.25以降でのslogについて](https://zenn.dev/drsprime/articles/f8cc13820e6a93)
- [[go1.25]json/v1→v2で何が変わったん？〜気になった変更点をお届け〜](https://zenn.dev/go_izumin/articles/json_v1_v2_update_details)

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/
[`encoding/json/v2`]: https://pkg.go.dev/encoding/json/v2 "json package - encoding/json/v2 - Go Packages"
[`encoding/json`]: https://pkg.go.dev/encoding/json "json package - encoding/json - Go Packages"
[`os`]: https://pkg.go.dev/os "os package - os - Go Packages"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400535" %}} <!-- 効率的なGo : Efficient Go -->
{{% review-paapi "B0DNYMMBBQ" %}} <!-- Go言語で学ぶ並行プログラミング -->
