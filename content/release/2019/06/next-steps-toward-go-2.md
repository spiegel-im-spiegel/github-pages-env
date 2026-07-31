+++
title = "Go 1.13 と 1.14 （Go 2 へ向けて）"
date =  "2019-06-29T16:50:58+09:00"
description = "Go 1.13 のベータ版が登場したようだ。リリースノートも併せて公開されている。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

## 8月に正式リリースされる [Go] 1.13 の主な機能

[Go] 1.13 のベータ版が登場したようだ。
リリースノートも併せて公開されている。

- [Go 1.13 Beta 1 is released - Google Group](https://groups.google.com/forum/#!topic/golang-announce/fmU5r5NfHZs)
- [Go 1.13 Release Notes - The Go Programming Language](https://tip.golang.org/doc/go1.13)

[Go] 1.13 では数値のリテラル表現（2進数表現や浮動小数点数の16進数表現）など色々と重要な機能追加があるが，主なものは以下の通り。

### errors パッケージへの機能追加

以前に紹介した `golang.org/x/xerrors` の機能が正式に `errors` パッケージに組み込まれるようだ。

- [階層化 Error パッケージ “xerrors” を試してみる]({{< ref "/golang/xerrors.md" >}})

`golang.org/x/xerrors` の機能をほぼ踏襲しているみたいなので，既に `golang.org/x/xerrors` を使っている人は僅かな変更で対応できるだろう。

- [エラー・ハンドリング再考]({{< ref "/golang/error-handling-again.md" >}})

### 環境変数 GO111MODULE の変更

環境変数 `GO111MODULE` の値が `auto` の際の挙動が変わるようだ。
具体的には

{{< fig-quote type="markdown" title="Go 1.13 Release Notes" link="https://tip.golang.org/doc/go1.13" lang="en" >}}
{{< quote >}}The [`GO111MODULE`](https://tip.golang.org/cmd/go/#hdr-Module_support) environment variable continues to default to `auto`, but the `auto` setting now activates the module-aware mode of the go command whenever the current working directory contains, or is below a directory containing, a `go.mod` file — even if the current directory is within `GOPATH/src`.{{< /quote >}}
{{< /fig-quote >}}

ということで `GOPATH` 以下にあるソースコードでも `go.mod` ファイルがあればモジュール対応モードで管理が可能なようだ。

- [モジュール対応モードへの移行を検討する]({{< ref "/golang/go-module-aware-mode.md" >}})

### GOPRIVATE, GOPROXY/GONOPROXY および GOSUMDB/GONOSUMDB

以前

- [Go モジュールのミラーリングとインデックス化]({{< ref "/golang/mirror-index-and-checksum-database-for-go-module.md" >}})

を紹介したが，この機能を制御するために `GOPRIVATE`, `GOPROXY`/`GONOPROXY` および `GOSUMDB`/`GONOSUMDB` 各環境変数が追加される。

`GOPROXY` の既定値は `https://proxy.golang.org,direct` となっている。
これを無効にする場合は `direct` のみをセットすればよい。

ちなみに 1.13 からは `go env` コマンドが拡張され

```text
$ go env -w GOPROXY=direct
```

という感じに設定できるらしい。
これでシステムの環境変数を汚さずに済む（笑）

`GOSUMDB` については

{{< fig-quote type="markdown" title="Go 1.13 Release Notes" link="https://tip.golang.org/doc/go1.13" lang="en" >}}
{{< quote >}}The new [`GOSUMDB`](https://tip.golang.org/cmd/go/#hdr-Module_authentication_failures) environment variable identifies the name, and optionally the public key and server URL, of the database to consult for checksums of modules that are not yet listed in the main module's `go.sum` file. If `GOSUMDB` does not include an explicit URL, the URL is chosen by probing the `GOPROXY` URLs for an endpoint indicating support for the checksum database, falling back to a direct connection to the named database if it is not supported by any proxy.{{< /quote >}}
{{< /fig-quote >}}

ということらしい。

また `GOPRIVATE` 環境変数を使えばミラーリングやチェックサム・データベースの対象から外すモジュールを指定できるようだ。

{{< fig-quote type="markdown" title="Go 1.13 Release Notes" link="https://tip.golang.org/doc/go1.13" lang="en" >}}
{{< quote >}}The new [`GOPRIVATE`](https://tip.golang.org/cmd/go/#hdr-Module_configuration_for_non_public_modules) environment variable indicates module paths that are not publicly available. It serves as the default value for the lower-level `GONOPROXY` and `GONOSUMDB` variables, which provide finer-grained control over which modules are fetched via proxy and verified using the checksum database.{{< /quote >}}
{{< /fig-quote >}}

Google はミラーリング・サービスとして [`proxy.golang.org`](https://proxy.golang.org) を，データベース・サービスとして [`sum.golang.org`](https://sum.golang.org/) を提供しているが，個人的には

{{< fig-quote type="markdown" title="Go モジュールのミラーリングとインデックス化" link="/golang/mirror-index-and-checksum-database-for-go-module/" >}}
{{< quote >}}[プライバシーに敵対的な企業]({{< ref "/remark/2018/04/handling-privacy.md" >}} "誰がプライバシーを支配するのか")のサービスだと思うとあまり利用したくない気分{{< /quote >}}
{{< /fig-quote >}}

なので

```text
$ go env -w GOPROXY=direct
$ go env -w GOSUMDB=off
```

として無効にするのがいいかな。
まぁ，8月に正式版が出てから様子を見て方針を決めればいいか。

## [Go] 1.14 へ向けて： try() 組み込み関数によるエラー検査

以下のブログ記事では [Go] 1.14 および最終的な [Go] 2 へ向けての方針が書かれている。

- [Next steps toward Go 2 - The Go Blog](https://blog.golang.org/go2-next-steps)

この中で注目したいのは [Go] 1.14 で追加されるというエラー検査機能だ。

エラーの検査については以前

- [次期 Go 言語で導入される（かもしれない）新しいエラー・ハンドリングについて予習する]({{< ref "/golang/error-handling-in-go-2.md" >}})

において `check` 式（`check` expression）と `handle` 構文（`handle` statement）の提案を紹介したが，最終的には `try()` 組み込み関数を導入することにしたようだ。

- [Proposal: A built-in Go error check function,](https://go.googlesource.com/proposal/+/master/design/32437-try-builtin.md)

具体的には

```go
func foo() (T1, T2, ..., Tn, error) { ... }
```

という関数に対して

```go
v1, v2, ..., vn := try(foo())
```

のように記述できる。
`try()` 組み込み関数は `foo()` 関数の返り値の `error` 値を見て `nil` でなければ値をセットして `return` する。
概念的にはこんな感じ

```go
var err error
v1, v2, ..., vn, te := foo()
if te != nil {
    err = te
    return 
}
```

セットされた `error` は `defer` 構文で捕まえてまとめて処理できる。

```go
func bar() (err error) {
    defer func() {
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
    }()

    v1, v2, ..., vn, := try(foo())
    ...
    return
}
```


実際のコードで考えてみよう。
たとえばファイルのコピーを行う関数 `copyFile()` は

```go
func copyFile(src, dst string) error {
    r, err := os.Open(src)
    if err != nil {
        return err
    }
    defer r.Close()

    w, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer w.Close()

    if _, err := io.Copy(w, r); err != nil {
        return err
    }
    return nil
}
```

`try()` 組み込み関数と `defer` 構文を使って以下のように書き直せる。

```go
func copyFile(src, dst string) (err error) {
    defer func() {
        if err != nil {
            err = fmt.Errorf("copyFile %s %s: %v", src, dst, err)
        }
    }()

    r := try(os.Open(src))
    defer r.Close()

    w := try(os.Create(dst))
    defer w.Close()

    try(io.Copy(w, r))
    return nil
}
```

`try()` 組み込み関数を導入することにより，演算子や構文を追加することなく後方互換性を確保しつつ仕様を拡張できるというのは大きい。
来年の冬が楽しみだなぁ。



## ブックマーク

- [try - Go の新しいエラーハンドリング (Go 1.14で導入予定) - Qiita](https://qiita.com/worldhello/items/9f2fe358d57a8835706e)

[Go]: https://go.dev/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
