+++
title = "Go 1.16 からのモジュール管理"
date =  "2021-02-21T16:15:24+09:00"
description = "覚え書きとして記しておく"
image = "/images/attention/go-logo_blue.png"
tags = [ "module", "golang", "engineering", "management" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先日リリースされた [Go] 1.16 でモジュール管理がいくつか変更になったので，覚え書きとして記しておく。
なお，このブログで書き散らした内容をまとめる形で 以下の Zenn 記事を書いた。
こちらも併せてどうぞ。

- [Go のモジュール管理【バージョン 1.16 改訂版】](https://zenn.dev/spiegel/articles/20210223-go-module-aware-mode)

## GO111MODULE 既定値の変更

環境変数 `GO111MODULE` の既定値が `auto` から `on` に変更になった。
`GO111MODULE` の取りうる値は以下の通り。

| 値     | 内容 |
| ------ | ---- |
| `on`   | 常にモジュール対応モードで動作する |
| `off`  | 常に GOPATH モードで動作する  |
| `auto` | `$GOPATH/src` 以下のディレクトリに配置され `go.mod` を含まないパッケージは GOPATH モードで，それ以外はモジュール対応モードで動作する |

`GO111MODULE` の値を `auto` に戻すのであれば `go env` コマンドで

```text
$ go env -w GO111MODULE=auto
```

とする。
[Go] の環境変数の取り扱いについては，拙文「[Go 言語の環境変数管理]({{< relref "./go-env.md" >}})」を参照のこと。

## go.mod および go.sum の自動更新の抑制

[Go] 1.15 までは `go build` や `go test` などのコマンドで `go.mod` や `go.sum` の内容が勝手に更新されていたが， 1.16 からは自動では更新されなくなった。

なので，コード上の `import` で新しい外部パッケージを追加しても `go.mod` や `go.sum` に記述がないと

```text
$ go test ./...
main.go:9:2: no required module provides package github.com/spiegel-im-spiegel/cov19jpn/chart; to add it:
	go get github.com/spiegel-im-spiegel/cov19jpn/chart
```

とか

```text
$ go test ./...
go: github.com/spiegel-im-spiegel/cov19jpn@v0.2.0: missing go.sum entry; to add it:
	go mod download github.com/spiegel-im-spiegel/cov19jpn
```

みたいなエラーが出たりする。

`go.mod` や `go.sum` をいい感じに更新したいのであれば

```text
$ go mod tidy
```

とするとよい。

## バージョン付きの go install

`go install` コマンドで指定するパッケージパスにバージョン番号サフィックスを付けることができるようになった。

```text
$ go install golang.org/x/tools/gopls@v0.6.5
```

これでバージョンを指定してモジュールのビルド&インストールができる。
ただし `go.mod` ファイルが `replace` や `exclude` ディレクティブを含んでいると `go install` が失敗するみたい。

```text
$ go install github.com/spiegel-im-spiegel/cov19jpn@v0.2.0
go install github.com/spiegel-im-spiegel/cov19jpn@v0.2.0: github.com/spiegel-im-spiegel/cov19jpn@v0.2.0
	The go.mod file for the module providing named packages contains one or
	more replace directives. It must not contain directives that would cause
	it to be interpreted differently than if it were the main module.
```

この場合は今までどおり `go get` コマンドでビルドまでやってくれるが， `go get` コマンドによるビルドは[将来的に廃止](https://github.com/golang/go/issues/43684 "cmd/go: deprecate installing binaries using 'go get' in Go 1.17 and make 'go get -d' the default behavior · Issue #43684 · golang/go")になるみたいなので，早めになんとかしたいものである。

でもなぁ，私の場合「[Go 依存パッケージの脆弱性検査]({{< relref "//check-for-vulns-in-golang-dependencies.md" >}})」の誤検出対策に `replace` ディレクティブを使ってるから悩ましいんだよなぁ...

## 特定バージョンのモジュールの撤回

`go.mod` で `retract` ディレクティブを使って特定バージョンのモジュールを撤回できるようになった（`go` ディレクティブが `1.16` 以上の場合）。

こんな感じにコメントとともに指定するといいらしい（コメント付けられたのか）。

```text
// Remote-triggered crash in package foo. See CVE-2021-01234.
retract v1.0.5
```

これでこのバージョンを使おうとしても

```text
$ go list -m -u all
example.com/lib v1.0.0 (retracted)
$ go get .
go: warning: example.com/lib@v1.0.5: retracted by module author:
    Remote-triggered crash in package foo. See CVE-2021-01234.
go: to switch to the latest unretracted version, run:
    go get example.com/lib@latest
```

てな感じにコメントの内容で警告が表示されるそうだ。

## GOVCS によるバージョン管理ツールの制御

環境変数 `GOVCS` を使ってリポジトリとバージョン管理ツールを関連付けることができる。
これを使って悪意のあるリポジトリ・サーバへのアクセスを制限することを意図しているようだ。

環境変数の設定は `go env -w` コマンドを使うとよいだろう。
たとえば

```text
$ go env -w "GOVCS=github.com:git,evil.com:off,*:git|hg"
```

とすれば， github.com サイトでは git のみ許容し evil.com は使用禁止，その他のサイトでは git および mercurial のみ許容する，といった指定ができるらしい。

なお `GOVCS` の既定値は `public:git|hg,private:all` となっている。

## ブックマーク

- [New module changes in Go 1.16 - The Go Blog](https://blog.golang.org/go116-module-changes)

- [Go モジュールのバージョン管理]({{< relref "./versioning-of-go-modules.md" >}})
- [Go 1.16 がリリースされた]({{< ref "/release/2021/02/go-1_16-is-released.md" >}})

[Go]: https://go.dev/

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
