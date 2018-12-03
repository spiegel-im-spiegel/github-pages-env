+++
date = "2016-03-29T22:16:41+09:00"
update = "2018-02-07T15:29:42+09:00"
description = "Go 言語の引数は基本的に「値渡し（call by value）」である。 Instance の値ではなく実体を渡したいときにしたい場合はポインタを使う。"
draft = false
tags = ["golang", "function", "pointer", "defer"]
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
  tumblr = ""
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

[^t]: 型については「[Go 言語における「オブジェクト」]({{< relref "object-oriented-programming.md" >}})」を参照のこと。

```go
package main

import "fmt"

func add(x int, y int) int {
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

返り値を組（tuple）で定義することもできる。

```go
func split(sum int) (int, int) {
    x = sum * 4 / 9
    y = sum - x
    return x, y
}
```

また返り値には，以下に示すように，あらかじめ名前をつけることもできる。

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

[^p]: このコードについては「[エラー・ハンドリングについて]({{< relref "error-handling.md" >}})」で解説している。ちなみに [panic] を潰して error を返すのはエラー・ハンドリングとしては推奨できない。

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
このため引数の値渡しは thread safe なコードに向いている。
ただし関数呼び出し時に instance の値が常にコピーされるため[^copy]，サイズの大きな instance の場合は呼び出し時のコストが高くなる。
引数の値渡しが有利な例としては value object を構成する場合などが考えられる。

[^copy]: 値がどこにコピーされるかは型によって異なる。 [string] 以外の基本型は値がスタックに積まれる。 [string] および基本型以外はヒープ領域に値がコピーされその参照（＝ポインタ）がスタックに積まれる。

Instance の値ではなく実体を渡したいときがある。
この場合はポインタを使う。
つまり instance へのポインタ値を渡すのである。

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
このように instance へのポインタ値を引数として渡すやり方をこの記事では，通常の「値渡し」とは区別して，「ポインタ渡し」と呼ぶことにする[^cbp1]。

[^cbp1]: 以前は，「値渡し」に対置する言葉として「参照渡し（call by reference）」と書いていたが，最近は「もう参照渡しとは言わせない」とか言う原理主義者がいるそうで，言いがかりに巻き込まれないよう「ポインタ渡し」と言い回しを変えることにした。やれやれ。

内部状態を持つ instance を引数に指定する場合はポインタ渡しにする必要がある。
しかし引数をポインタ渡しにすると関数実行が thread safe でなくなる可能性がある。
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
したがって通常は「[Go 言語]のポインタは nullable 参照[^nr] と同じ」と考えてよい。

[^nr]: nullable 参照は「null を許容する参照」くらいの意味。 [Go 言語]なら nil 値。 [Go 言語]は「null 安全（null safty）」ではないので null 参照（＝無効な参照）の始末について instance を参照する側が責務を負うことになる。（参考： [「null 安全」について]({{< ref "/remark/2016/11/null-safety.md" >}})）

なお，ポインタ演算が必要な場合は [`unsafe`] パッケージを使う。

### Slice, Map, Channel は参照渡しのように振る舞う

[slice], [map], [channel] は組み込み型だが内部状態を持つ[^make]。
これらの型の instance を引数に渡す場合は参照渡しのように振る舞う[^slc]。

[^make]: [slice], [map], [channel] は内部状態を持つため `new()` 関数ではなく `make()` 関数を使う。  `make()` 関数は生成した instance への参照（実体はポインタ）を返す。
[^slc]: 詳しくは「[配列と Slice]({{< relref "array-and-slice.md" >}})」および「[Map の話]({{< relref "map.md" >}})」を参照のこと。

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

ただし固定の配列や [string] 型[^str] の instance は「値」として振る舞うため[^n]，引数に指定した場合も値渡しのように振る舞う。
[slice] とは挙動が異なるためテキトーなコードを書いていると混乱しやすい。

[^str]: [string] 型の実体は `[]byte` 型である。
[^n]: 固定の配列や [string] 型の instance は nil 値を持たない「non-null 参照」と言える。ちなみに [string] 型のゼロ値は空文字列である。

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

固定配列や [string] 型をポインタ渡しすることはできる。
もっとも [string] 型の instance は「不変（immutable）」なのでポインタ渡しが必要な局面はほとんど無いと思うが[^s]。

[^s]: このような需要としては文字列操作で「NULL 状態」が必要な場合であろう。たとえば DBMS にアクセスする場合は NULL 状態を扱う必要がある。なお [Go 言語]のコア・パッケージには [`database/sql`] があり `NullString` を使うことにより NULL 状態を扱える。このように NULL 状態を扱う必要がある場合は，直にポインタ操作するのではなく，何らかの value object を用意してカプセル化するほうが安全である。

固定配列をポインタ渡しする例は以下の通り。

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

固定配列は不変ではないが，配列を操作するのであれば固定配列ではなく [slice] のほうが扱いやすい。
たとえば上のコードでは `slc := ary[:] ` といった感じにキャストするか最初から `ary := []int{0, 1, 2, 3}` と初期化すれば [slice] として扱える。

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
つまり `v` は `Add()` 関数の0番目の引数として振る舞い，値渡しでセットされる。

Method receiver の型をポインタ型にすればポインタ渡しにできる。

{{< highlight go "hl_lines=14" >}}
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
{{< /highlight >}}

この場合も calling sequence としては `v.Add(dv)` と `(*Vertex).Add(v, dv)` は等価である。

### Method Receiver の暗黙的変換

Method receiver を値にした場合，呼び出し元の instance がポインタ型であっても暗黙的にコピーが発生し値渡しに変換される。

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

Method receiver をポインタにした場合も，やはり暗黙的にポインタ渡しに変換される。

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
まずは値渡しの場合。

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
    vv := v.Add(Vertex{X: 3, Y: 4}) //panic!
    fmt.Println(v)
    fmt.Println(vv)
}
```

この場合は `Add()` 関数呼び出し時に [panic] になる。

```text
panic: runtime error: invalid memory address or nil pointer dereference
```

まぁこれは分かりやすいよね。
ではポインタ渡しの場合はどうなるか。

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
    v.X += dv.X //panic!
    v.Y += dv.Y
}
```

`v` 内の要素を参照しようとしたところで [panic] になる。
Method receiver をポインタ渡しにする場合は nil 値に注意する必要がある。

## for-range 構文も値渡し

余談だが for-range 構文も値渡し（つまりコピーが発生する）ので注意が必要である。
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

{{< highlight go "hl_lines=8-10" >}}
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
{{< /highlight >}}

とするしかない。

## ブックマーク

- [Go 言語の値レシーバとポインタレシーバ | Step by Step](https://skatsuta.github.io/2015/12/29/value-receiver-pointer-receiver/)

[Go 言語に関するブックマーク集はこちら]({{< relref "bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[defer]: http://blog.golang.org/defer-panic-and-recover "Defer, Panic, and Recover - The Go Blog"
[panic]: http://blog.golang.org/defer-panic-and-recover "Defer, Panic, and Recover - The Go Blog"
[`unsafe`]: https://golang.org/pkg/unsafe/ "unsafe - The Go Programming Language"
[slice]: http://golang.org/ref/spec#Slice_types
[map]: http://golang.org/ref/spec#Map_types
[channel]: http://golang.org/ref/spec#Channel_types
[string]: http://golang.org/ref/spec#String_types
[`database/sql`]: https://golang.org/pkg/database/sql/ "sql - The Go Programming Language"
