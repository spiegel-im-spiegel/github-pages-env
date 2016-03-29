+++
date = "2016-03-29T22:16:41+09:00"
description = "Go 言語の引数は基本的に「値渡し（call by value）」である。「参照渡し（call by reference）」にしたい場合はポインタを使う。"
draft = false
tags = ["golang", "function", "pointer"]
title = "関数とポインタ"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

いまさらな内容なのだが覚え書きとして記しておく。

## Go 言語における Calling Sequence

まずは簡単な足し算の関数を定義してみる。

```go
func add(x int, y int) int {
	return x + y
}
```

`add` に続く括弧内が引数を定義していて，括弧の後ろの `int` は返り値の型[^t] を示している。
`add()` 関数を呼び出すには以下のように記述する。

[^t]: 型については「[Go 言語における「オブジェクト」]({{< relref "golang/object-oriented-programming.md" >}})」を参照のこと。

```go
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	ans := add(42, 13)
	fmt.Println(ans)
}
```

`x` と `y` は同じ int 型なので以下のように記述することもできる。

```go
func add(x, y int) int {
	return x + y
}
```

返り値として複数の値を定義することもできる。

```go
func split(sum int) (int, int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}
```

また返り値は以下のように名前をつけることもできる。

```go
func add(x, y int) (ans int) {
	ans = x + y
	return
}
```

最後の `return` がないとコンパイル・エラーになるので注意。
この書き方は [defer] 構文と組み合わせるときに威力を発揮する。

```go
package main

import "fmt"

func main() {
	err := r()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Normal End.")
	}
}

func r() (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("Recovered from: %v", rec)
		}
	}()

	f()
	err = nil
	return
}

func f() {
	panic("Panic!")
}
```

このコード[^p] では `r()` 関数内で [panic] を捕まえ， 返り値の `err` に値をセットしなおしている。

[^p]: このコードについては「[エラー・ハンドリングについて]({{< relref "golang/error-handling.md" >}})」で解説している。ちなみに [panic] を潰して error を返すのはエラー・ハンドリングとしてはいいやり方ではない。

### Go 言語の引数は「値渡し」

[Go 言語]の引数は基本的に「値渡し（call by value）」である。
たとえば先程の足し算を

```go
func add(x, y int) int {
	x += y
	return x
}
```

と定義した場合でも

```go
package main

import "fmt"

func add(x, y int) int {
	x += y
	return x
}

func main() {
    x := 42
    y := 13
	ans := add(x, y)
	fmt.Printf("%d + %d = %d\n", x, y, ans) //output: 42 + 13 = 55
}
```

呼び出し元で `add()` 関数の引数に渡した instance は関数実行後も変化しない。
このため「値渡し」は thread safe なコードに向いている。
たとえば value object を構成する際には関連する関数は「値渡し」のほうが安全である。
ただし関数呼び出し時に instance の値が常にコピーされるため[^copy]，サイズの大きな instance の場合は呼び出し時のコストが高くなる。

[^copy]: 値がどこにコピーされるかは型によって異なる。 [string] 以外の基本型は値がスタックに積まれる。 [string] および基本型以外はヒープ領域に値がコピーされそのポインタがスタックに積まれる。

引数を「参照渡し（call by reference）」にしたい場合はポインタを使う。
つまり instance のポインタ値を渡すのである。

```go
package main

import "fmt"

func add(x, y *int) int {
	*x += *y
	return *x
}

func main() {
	x := 42
	y := 13
	ans := add(&x, &y)
	fmt.Printf("%d + %d = %d\n", x, y, ans) //output: 55 + 13 = 55
}
```

このコードでは `add()` 関数実行後の `x` の値が変更されている。
内部状態を持つ instance を引数に指定する場合は参照渡しにする必要がある。
しかし引数を参照渡しにすると関数実行が thread safe でなくなる可能性がある。
また引数の値が nil の場合も考慮する必要がある。

ちなみに [Go 言語]では通常の方法ではポインタ演算ができない。
たとえば，ついうっかり

```go
func add(x, y *int) int {
	x += y
	return *x
}
```

とか書いてしまっても

```text
invalid operation: x += y (operator + not defined on pointer)
```

とコンパイル・エラーになる。
ポインタ演算が必要な場合は [`unsafe`] パッケージを使う。

### Slice, Map, Channel は常に「参照渡し」

[slice], [map], [channel] は組み込み型だが内部状態を持つ[^make]。
したがって，これらの型の instance を引数に渡す場合はつねに「参照渡し」になる（つまり instance のコピーは発生しない）。

[^make]: [slice], [map], [channel] は内部状態を持つため `new()` 関数ではなく `make()` 関数で instance を生成する。

```go
package main

import "fmt"

func setItem(ary map[int]int, index, item int) {
	ary[index] = item
}

func main() {
    ary := map[int]int{0: 0}
	fmt.Println(ary) //output: map[0:0]
	setItem(ary, 0, 1)
	fmt.Println(ary) //output: map[0:1]
	setItem(ary, 10, 10)
	fmt.Println(ary) //output: map[0:1 10:10]
}
```

ただし固定の配列や [string] 型[^str] の instance は「値」として振る舞うため[^n]，引数に指定した場合も「値渡し」になる。
[slice] とは挙動が異なるためテキトーなコードを書いていると混乱しやすい。

[^str]: [string] 型の実体は `[]byte` 型である。
[^n]: たとえば固定の配列や [string] 型の instance は nil 値を持たない。 [string] 型のゼロ値は空文字列である。

```go
package main

import "fmt"

func setItem(ary [4]int, index, item int) {
	ary[index] = item
}

func main() {
	ary := [4]int{0, 1, 2, 3}
	fmt.Println(ary) //output: [0 1 2 3]
	setItem(ary, 1, 10)
	fmt.Println(ary) //output: [0 1 2 3]
	ary[2] = 200
	fmt.Println(ary) //output: [0 1 200 3]
}
```

固定配列や [string] 型を「参照渡し」にしたい場合はやはりポインタ値を渡す。

```go
package main

import "fmt"

func setItem(ary *[4]int, index, item int) {
	(*ary)[index] = item
}

func main() {
	ary := [4]int{0, 1, 2, 3}
	fmt.Println(ary) //output: [0 1 2 3]
	setItem(&ary, 1, 10)
	fmt.Println(ary) //output: [0 10 2 3]
}
```

実際には [string] 型の instance は「不変（immutable）」なので「参照渡し」が必要な局面はほとんど無いと思われる。
固定配列は不変ではないが，配列を操作するのであれば固定配列ではなく [slice] のほうが扱いやすい。
たとえば上のコードでは `ary := []int{0, 1, 2, 3}` と初期化すれば [slice] として扱える。

## Method Receiver

ある型に関数を関連付ける場合は method receiver を使う。

```go
type Vertex struct {
	X int
	Y int
}

func (v Vertex) Add(dv Vertex) Vertex {
	v.X += dv.X
	v.Y += dv.Y
	return v
}
```

`(v Vertex)` の部分が method receiver である。
`Add()` 関数を呼び出すには以下のように記述する。

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func (v Vertex) String() string {
	return fmt.Sprint("X = ", v.X, ", Y = ", v.Y)
}

func (v Vertex) Add(dv Vertex) Vertex {
	v.X += dv.X
	v.Y += dv.Y
	return v
}

func main() {
	v := Vertex{X: 1, Y: 2}
	vv := v.Add(Vertex{X: 3, Y: 4})
	fmt.Println(v)  //output: X = 1, Y = 2
	fmt.Println(vv) //output: X = 4, Y = 6
}
```

関数の calling sequence としては `v.Add(dv)` と `Vertex.Add(v, dv)` は等価である。
つまり `v` は `Add()` 関数の0番目の引数として振る舞い，「値渡し」でセットされる。

Method receiver の型をポインタ型にすれば「参照渡し」にできる。

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func (v Vertex) String() string {
	return fmt.Sprint("X = ", v.X, ", Y = ", v.Y)
}

func (v *Vertex) Add(dv Vertex) {
	if v == nil {
		return
	}
	v.X += dv.X
	v.Y += dv.Y
}

func main() {
	v := &Vertex{X: 1, Y: 2}
	v.Add(Vertex{X: 3, Y: 4})
	fmt.Println(v) //output: X = 4, Y = 6
}
```

この場合も calling sequence としては `v.Add(dv)` と `(*Vertex).Add(v, dv)` は等価である。

### Method Receiver の暗黙的変換

Method receiver を「値渡し」にした場合，呼び出し元の instance がポインタ型であっても暗黙的に「値渡し」に変換される。

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func (v Vertex) String() string {
	return fmt.Sprint("X = ", v.X, ", Y = ", v.Y)
}

func (v Vertex) Add(dv Vertex) Vertex {
	v.X += dv.X
	v.Y += dv.Y
	return v
}

func main() {
	v := &Vertex{X: 1, Y: 2} //pointer
	vv := v.Add(Vertex{X: 3, Y: 4})
	fmt.Println(v)  //output: X = 1, Y = 2
	fmt.Println(vv) //output: X = 4, Y = 6
}
```

Method receiver を「参照渡し」にした場合も暗黙的に「参照渡し」に変換される。

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func (v Vertex) String() string {
	return fmt.Sprint("X = ", v.X, ", Y = ", v.Y)
}

func (v *Vertex) Add(dv Vertex) {
	if v == nil {
		return
	}
	v.X += dv.X
	v.Y += dv.Y
}

func main() {
	v := Vertex{X: 1, Y: 2} //not pointer
	v.Add(Vertex{X: 3, Y: 4})
	fmt.Println(v) //output: X = 4, Y = 6
}
```

### Method Receiver の値が nil の場合

Method receiver の値が nil の場合はどうなるか。
まずは「値渡し」の場合。

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func (v Vertex) String() string {
	return fmt.Sprint("X = ", v.X, ", Y = ", v.Y)
}

func (v Vertex) Add(dv Vertex) Vertex {
	v.X += dv.X
	v.Y += dv.Y
	return v
}

func main() {
	v := (*Vertex)(nil) //nil
	vv := v.Add(Vertex{X: 3, Y: 4})
	fmt.Println(v)
	fmt.Println(vv)
}
```

この場合は `Add()` 関数呼び出し時に [panic] になる。

```text
panic: runtime error: invalid memory address or nil pointer dereference
```

まぁこれは分かりやすいよね。
では「参照渡し」の場合はどうなるか。

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func (v Vertex) String() string {
	return fmt.Sprint("X = ", v.X, ", Y = ", v.Y)
}

func (v *Vertex) Add(dv Vertex) {
	if v == nil {
		return
	}
	v.X += dv.X
	v.Y += dv.Y
}

func main() {
	v := (*Vertex)(nil) //nil
	v.Add(Vertex{X: 3, Y: 4})
	fmt.Println(v) //output: <nil>
}
```

実は `Add()` 関数呼び出し時点では [panic] にはならない。
上のコードでは `v` に nil が渡される。
したがって `Add()` 関数内の条件文を削除すると

```go
func (v *Vertex) Add(dv Vertex) {
	v.X += dv.X
	v.Y += dv.Y
}
```

`v` 内の要素を参照としたところで [panic] になる。
Method receiver を「参照渡し」にする場合は nil 値に注意する必要がある。

## for-range 構文も「値渡し」

余談だが for-range 構文も「値渡し」（つまりコピーが発生する）なので注意が必要である。
たとえば以下のコードで

```go
package main

import "fmt"

func main() {
	ary := []int{0, 1, 2, 3}
	fmt.Println(ary) //output: [0 1 2 3]
	for _, item := range ary {
		item += 10
	}
	fmt.Println(ary) //output: [0 1 2 3]
}
```

for-range 構文内の `item` は `ary` 内の要素を指すのではなく要素のコピーである。
したがって `item` を操作しても `ary` には影響しない。
`ary` 内の要素を操作するのであれば素朴に

```go
package main

import "fmt"

func main() {
	ary := []int{0, 1, 2, 3}
	fmt.Println(ary) //output: [0 1 2 3]
	for i := 0; i < len(ary); i++ {
		ary[i] += 10
	}
	fmt.Println(ary) //output: [10 11 12 13]
}
```

とするしかない。

## ブックマーク

- [Go 言語の値レシーバとポインタレシーバ | Step by Step](https://skatsuta.github.io/2015/12/29/value-receiver-pointer-receiver/)

[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[defer]: http://blog.golang.org/defer-panic-and-recover "Defer, Panic, and Recover - The Go Blog"
[panic]: http://blog.golang.org/defer-panic-and-recover "Defer, Panic, and Recover - The Go Blog"
[`unsafe`]: https://golang.org/pkg/unsafe/ "unsafe - The Go Programming Language"
[slice]: http://golang.org/ref/spec#Slice_types
[map]: http://golang.org/ref/spec#Map_types
[channel]: http://golang.org/ref/spec#Channel_types
[string]: http://golang.org/ref/spec#String_types
