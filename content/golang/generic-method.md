+++
title = "Generic Method について"
date =  "2026-11-22T13:42:41+09:00"
description = "Go 1.27 から Generic Method がサポートされた。"
isCJKLanguage = true
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "generics", "language" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先日行われた『[実用 Go言語 第2版]』の[読書会](https://yokohama-go-reading.connpass.com/event/410042/ "第81回横浜Go読書会（オンライン） - connpass")で

{{< fig-quote type="markdown" title="『実用 Go言語 第2版』5章 ジェネリクス" link="https://www.oreilly.co.jp/books/9784814401369/" >}}
Goのジェネリクスの残念な点が、メソッドで新たな型変数が追加できないという点です。
{{< /fig-quote >}}

という記述に対して，実は [Go] 1.27 から Generic Method がサポートされたんだよん，という雑談に発展した。
この記事では，遅ればせながら，復習がてら Go 1.27 で導入された Generic Method について解説する。

## Go 1.27 より前の Generics

せっかくなので『[実用 Go言語 第2版]』のサンプルコードを例に書いていこう。

たとえば，数値型 `T` の値の平均値を取る `Average[T]()` 関数を定義するとこんな感じになる。

{{< fig-quote class="nobox" type="markdown" title="『実用 Go言語 第2版』リスト 5-11：ジェネリクス化" link="https://www.oreilly.co.jp/books/9784814401369/" >}}
```go
package main

func Average[T int | float64](values ...T) T {
    var total T
    for _, v := range values {
        total += v
    }
    if len(values) == 0 {
        return 0
    }
    return total / T(len(values))
}
```
{{< /fig-quote >}}

えーっと。
コードに対するツッコミはなしで[^p1]。
私が書いたわけじゃないし（笑）

[^p1]: 読書会では「`if` 文は最初に持ってくるべきでは？」とか「要素数が `0` のときにゼロ値を返すのは間違ってるのでは？」とか散々にツッコまれていた（笑）

別の実装として，型に対しても型パラメータを指定できる。

{{< fig-quote class="nobox" type="markdown" title="『実用 Go言語 第2版』リスト 5-15：メソッドに型引数を追加したい" link="https://www.oreilly.co.jp/books/9784814401369/" >}}
```go
type Average[T int | float64] struct {
    total T
    count int
}

func (a *Average[T]) Add(v T) {
    a.total += v
    a.count += 1
}

func (a *Average[T]) Result() T {
    return a.total / T(a.count)
}
```
{{< /fig-quote >}}

`Average[T]` 型では `Add(T)` メソッドで値を加えていって `Result()` メソッドで結果を取り出す構成になっている（コードに対するツッコミはなしでw）。
ここで `Add()` メソッドの引数には `T` 型とは異なる数値型を指定し，内部でキャストして加算したいところであるが，1.27 より前までは，メソッドで独自の型パラメータを指定することはできなかった。
どうしてもキャストして加算したいのであれば，専用の Generic 関数を用意していた。

{{< fig-quote class="nobox" type="markdown" title="『実用 Go言語 第2版』リスト 5-16：関数にして渡す" link="https://www.oreilly.co.jp/books/9784814401369/" >}}
```go
func Add[T int | float64, RV int | float64](a *Average[T], v RV) {
    a.Add(T(v))
}
```
{{< /fig-quote >}}

ただ，この方法だと Generic 型に関連する Generic 関数が山ほど生えてくる可能性があり，特にコンテナ操作に関してはパッケージの構成が煩雑になると思われる。
記述自体も煩雑になりがちだし。

こういった理由で，メソッドに独自の型パラメータを追加できるようにして欲しいという要望が多かったようだ。

## Go 1.27 の Generic Method

そこで Generic Method の登場である。
前節の `Average[T]` 型を Generic Method を含む形に書き換えてみよう。

```go {hl_lines=["13-17"]}
package main

import (
	"fmt"
	"math"
)

type Average[T int | float64] struct {
	total T
	count int
}

func (a *Average[T]) Add[U int | float64](v U) *Average[T] {
	a.total += T(v)
	a.count += 1
	return a
}

func (a *Average[T]) Result() T {
	if a.count == 0 {
		// When T is int, the result of converting NaN to an integer is implementation-specific.
		return T(math.NaN())
	}
	return a.total / T(a.count)
}

func main() {
	fmt.Println("Average:", (&Average[float64]{}).Add(1).Add(2.0).Add(4).Result())
	// Average: 2.3333333333333335
}
```

オリジナルから少し細工をして`Add[U]()` メソッドを連結して記述できるよう返り値を自身のポインタにしている。
`main()` 関数では `Average[float64]` 型のインスタンスを生成し 1, 2.0, 4 と値を追加して平均値を取り出している。

### Generic Method の制約

Generic Method は `interface` のメソッドとして宣言できない。
たとえば

```go
type MyInterface interface { // interface type with generic method (not allowed)
	DoSomething[T any]()
}
```

と書いても "`interface method must have no type parameters`" とコンパイルエラーになる。

それじゃあ，というので `interface` 型のメソッド定義から型パラメータを外した上で

```go {hl_lines=["1-3"]}
type MyInterface interface { // interface type with non-generic method
	DoSomething()
}

type MyStruct struct{} // concrete type implementing MyInterface

func (m *MyStruct) DoSomething[T any]() { /* ... */ }

func New() MyInterface {
	return &MyStruct{} // compile error due to signature mismatch
}
```

と書いても，`interface` 型と concrete 型の間でシグニチャが合わないため，以下のエラーになる。

```text
cannot use &MyStruct{} (value of type *MyStruct) as MyInterface value in return statement: *MyStruct does not implement MyInterface (wrong type for method DoSomething)
		have DoSomething[T any]()
		want DoSomething()
```

`interface` 型のメソッドとして Generic Method を許すと，インターフェース値を介した呼び出しで実行され得る具象メソッドを一般には特定できず，必要な型引数の組み合わせをコンパイラがあらかじめ展開するのが難しい。

{{< fig-quote type="markdown" title="Generic Methods - The Go Programming Language" link="https://go.dev/blog/generic-methods" lang="en" >}}
But in more complex programs, the relation between an interface value method call and its possible targets is significantly less obvious. This relation typically spans across packages, making it difficult or impossible to deduce in the general case.
{{< /fig-quote >}}

というわけで `interface` 型では Generic Method 定義を使えなくしているようだ。

### Generic Method を使った実装

たとえば [Go] 1.27 の [`math/rand/v2`] 標準パッケージでは，Generic Method を使った乱数生成が可能になっている。

```go {hl_lines=[11]}
package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	rng := rand.New(rand.NewPCG(1, 2))
	for i := range 5 {
		fmt.Println(i+1, rng.N[int64](100)) // generate a random int64 number less than 100
	}
}
```

またサードパーティでは，以下のパッケージが Generic Method を使った実装を提供しているようだ。

- [ahmetb/go-linq: .NET LINQ capabilities in Go](https://github.com/ahmetb/go-linq)
- [goforj/collection: Fluent generic Go collections for explicit slice pipelines, transformations, filtering, aggregation, and collection helpers.](https://github.com/goforj/collection) : v4 以降
- [grandper/go-monad: A monadic types library for Go, providing Option, Result, Either, Promise, Lazy and IO monads with full Map, FlatMap, and Fold support.](https://github.com/grandper/go-monad)

参考にしていただきたい。

## ブックマーク

- [Generic Methods - The Go Programming Language](https://go.dev/blog/generic-methods)
- [Go 1.27 リリース連載：ジェネリクスメソッド (generic methods) | フューチャー技術ブログ](https://future-architect.github.io/articles/20260730a/)

[Go]: https://go.dev/
[実用 Go言語 第2版]: https://www.oreilly.co.jp/books/9784814401369/ "実用 Go言語 第2版 - O'Reilly Japan"
[`math/rand/v2`]: https://pkg.go.dev/math/rand/v2 "rand package - math/rand/v2 - Go Packages"

## 参考図書

{{< linkcard "51769bc6eecda48305121bbc56350f35b5599135" >}} <!-- 実用 Go言語 第2版 -->
{{< linkcard "9e8d8f717b1d1d85b0da7f07cb10a83b78d4227f" >}} <!-- プログラミング言語Go -->
{{< review-paapi "B0CFL1DK8Q" >}} <!-- Go言語 100Tips -->
{{< review-paapi "B0DNYMMBBQ" >}} <!-- Go言語で学ぶ並行プログラミング -->
