+++
title = "Go 1.14 リリース候補版"
date =  "2020-02-08T15:33:22+09:00"
description = "この記事では個人的に気になった点をかいつまんで紹介する。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "interface", "module", "concurrency" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.14 のリリース候補版が出た。
これに伴ってリリースノートのドラフト版も更新されたようだ。

- [Go 1.14 Release Candidate 1 is released - Google group](https://groups.google.com/forum/#!topic/golang-announce/mB1Mp9RlQw8)
- [Go 1.14 Release Notes - The Go Programming Language](https://tip.golang.org/doc/go1.14)

更にブログでは以下の記事で 1.15 についても言及されている。

- [Proposals for Go 1.15 - The Go Blog](https://blog.golang.org/go1.15-proposals)

詳しくはそれぞれの記事を読んでもらうとして，この記事では個人的に気になった点をかいつまんで紹介する。

## Try 終了のお知らせ

「[Go 1.13 と 1.14]({{< ref "/release/2019/06/next-steps-toward-go-2.md" >}} "Go 1.13 と 1.14 （Go 2 へ向けて）")」で紹介した `try()` 組み込み関数の導入は見送られたらしい。

{{< fig-quote type="markdown" title="Proposals for Go 1.15" link="https://blog.golang.org/go1.15-proposals" lang="en" >}}
{{< quote >}}Our attempt seven months ago at providing a better error handling mechanism, the [try proposal](https://golang.org/issue/32437), met good support but also strong opposition and we decided to abandon it. In its aftermath there were many follow-up proposals, but none of them seemed convincing enough, clearly superior to the try proposal, or less likely to cause similar controversy{{< /quote >}}.
{{< /fig-quote >}}

というわけで，エラー・ハンドリング周りはこれ以上の仕様追加・変更は（1.x の間は）なさそうである。

## 埋め込み Interface の改善

[Go] 1.14 では埋め込み interface の仕様が一部変更になる。

{{< fig-quote type="markdown" title="Go 1.14 Release Notes (draft)" link="https://tip.golang.org/doc/go1.14" lang="en" >}}
{{< quote >}}Per the [overlapping interfaces proposal](https://github.com/golang/proposal/blob/master/design/6977-overlapping-interfaces.md), Go 1.14 now permits embedding of interfaces with overlapping method sets: methods from an embedded interface may have the same names and identical signatures as methods already present in the (embedding) interface. This solves problems that typically (but not exclusively) occur with diamond-shaped embedding graphs. Explicitly declared methods in an interface must remain unique, as before{{< /quote >}}.
{{< /fig-quote >}}

[Go] では interface は振る舞いのみを定義する型だが，入れ子にすることができる。
こんな感じ。

```go
type Person interface {
	Name() string
	Age() int
}

type Employee interface {
	Person
	Level() int
    String() string
}
```

この例では `Employee` interface 型に `Person` interface 型が埋め込まれている。
つまり `Employee` 型では `Name()`, `Age()`, `Level()`, `String()` 各メソッドを要求しているわけだ。

ここで `Person` 型に `String()` メソッドを付けることを考える。

```go {hl_lines=[4]}
type Person interface {
	Name() string
	Age() int
    String() string
}
```

修正された `Person` 型を使って `Employee` 型を定義しようとしても {{% quote lang="en" %}}`duplicate method String`{{% /quote %}} とコンパイルエラーになる。
Interface 型の間で定義するメソッドを調整すればいいのだが，他パッケージの interface 型を埋め込む場合は，そのパッケージの仕様変更の影響をモロに受けることになる。

更に

{{< fig-img src="./diamond.png" link="./diamond.puml" width="946" >}}

のようなひし形構造になっている場合はより複雑になる。

[Go] 1.14 では（関数型の同一性も含めて）同じメソッドについては重複を許容する。
上述のコード例でもコンパイル・エラーにならないわけだ。

ただし

```go
type E1 interface{ M(x int) bool }
type E2 interface{ M(x float32) bool }
type I interface {
	E1
	E2
}
```

では（メソッド `M()` の型が同一ではないので）相変わらずコンパイル・エラーになるようだ。

## Preemptive なスケジューリング

これは [Shimane.go](https://shimane-go.connpass.com/ "Shimane.go - connpass") の Slack で教えてもらったのだが， [Go] 1.14 では preemptive (非協調的) なスケジューリング実装になるようだ。

{{< fig-quote type="markdown" title="Go 1.14 Release Notes (draft)" link="https://tip.golang.org/doc/go1.14" lang="en" >}}
{{< quote >}}Goroutines are now asynchronously preemptible. As a result, loops without function calls no longer potentially deadlock the scheduler or significantly delay garbage collection{{< /quote >}}.
{{< /fig-quote >}}

もともと [Go 言語]には，ランタイムによって並列処理の実装詳細を隠蔽することにより，コード記述としての平行処理に注力できるというメリットがあるが preemptive なスケジューリングによって強化されることになる。

ただし全てのプラットフォームで有効になるのではなく

- `windows/arm`
- `darwin/arm`
- `js/wasm`
- `plan9/*`

は例外となるらしい。
WebAssembly や Plan 9 が non-preemptive になるのは分かるが， ARM アーキテクチャって実装が難しいのか？

## モジュール対応モード

モジュール対応モード（module-aware mode）も色々と機能追加されるようだ。
特に `vendor` ディレクトリとの組み合わせは色々とできそうだ。

{{< fig-quote type="markdown" title="Go 1.14 Release Notes (draft)" link="https://tip.golang.org/doc/go1.14" lang="en" >}}
{{< quote >}}When the main module contains a top-level `vendor` directory and its `go.mod` file specifies go 1.14 or higher, the go command now defaults to `-mod=vendor` for operations that accept that flag. A new value for that flag, `-mod=mod`, causes the go command to instead load modules from the module cache (as when no vendor directory is present){{< /quote >}}.
{{< /fig-quote >}}

この辺は [Go] 1.14 正式版がリリースされてから試してみよう。

## Launched [pkg.go.dev]

- [Next steps for pkg.go.dev - The Go Blog](https://blog.golang.org/pkg.go.dev-2020)

[pkg.go.dev] は従来の [godoc.org] から置き換えることができる。

{{< fig-quote type="markdown" title="Next steps for pkg.go.dev" link="https://blog.golang.org/pkg.go.dev-2020" lang="en" >}}
{{< quote >}}Like [godoc.org](https://godoc.org/), pkg.go.dev serves Go documentation. However, it also understands modules and has information about past versions of a package!{{< /quote >}}
{{< /fig-quote >}}

実際に2020年後半には [godoc.org] へのリクエストを [pkg.go.dev] にリダイレクトする計画があるらしい。

{{< fig-quote type="markdown" title="Next steps for pkg.go.dev" link="https://blog.golang.org/pkg.go.dev-2020" lang="en" >}}
{{< quote >}}To minimize confusion about which site to use, later this year we are planning to redirect traffic from [godoc.org](https://godoc.org/) to the corresponding page on [pkg.go.dev](https://pkg.go.dev/){{< /quote >}}.
{{< /fig-quote >}}

標準パッケージもサードパーティのパッケージも同等に扱えるので，今後は [pkg.go.dev] を参照するのがいいかもしれない。

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[pkg.go.dev]: https://pkg.go.dev/
[godoc.org]: https://godoc.org/

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
{{% review-paapi "4908686033" %}} <!-- Goならわかるシステムプログラミング -->
