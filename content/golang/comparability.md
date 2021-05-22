+++
title = "インスタンスの比較可能性"
date =  "2020-02-02T17:52:58+09:00"
description = "少なくとも == および != 演算子が使えることを「比較可能」であると言う。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "programming", "interface" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[次のイベント](https://shimane-go.connpass.com/event/165192/ "Shimane.go#04 - connpass")に向けてネタの整理をしているところだが，その中でインスタンスの{{< ruby "comparability" >}}比較可能性{{< /ruby >}}についてきちんと整理したほうがよさそうな気がしたので，小ネタ記事として公開する。

## 用語の定義（暫定）

本題に入る前に，以下の2つの単語を，この記事限りの用語として定義する。
他所で使わないように（笑）

| 用語 | 意味                                          |
|:----:| --------------------------------------------- |
| 等値 | インスタンスの値が等しい（equal）こと         |
| 等価 | インスタンスの型が同一（identical）であること |

プログラミングの世界で等値と等価に関して議論があることは知っているが，今回はまるっと無視する。
だって鬱陶しいんだもん。

## 比較演算子

[Go 言語]ではインスタンス同士の比較演算子として

| 演算子 | 名称             |
|:------:| ---------------- |
|  `==`  | equal            |
|  `!=`  | not equal        |
|  `<`   | less             |
|  `<=`  | less or equal    |
|  `>`   | greater          |
|  `>=`  | greater or equal |

の5つが定義されている。
このうち少なくとも `==` および `!=` が使えることを「比較可能（comparable）」であると言う。

## 型の比較可能性

インスタンス同士が比較可能であるためには以下の2つの条件がが必要である。

1. インスタンスの型が同一（等価）であること
2. インスタンスの型が比較可能であること

たとえばある型を別の型に再定義しただけの場合でも等価とは見なされず，コンパイルエラーになる。

```go
package main

import "fmt"

type Number int

func main() {
    var c1 int = 1
	var c2 Number = 1
	fmt.Println(c1 == c2) //compile error: mismatched types int and Number
}
```

ただし，等価な型にキャスト可能であれば

```go {hl_lines=[10]}
package main

import "fmt"

type Number int

func main() {
	var c1 int = 1
	var c2 Number = 1
	fmt.Println(c1 == int(c2)) //true
}
```

などとできる。
また type alias であれば等価とみなされる。

```go {hl_lines=[5]}
package main

import "fmt"

type Number = int

func main() {
	var c1 int = 1
	var c2 Number = 1
	fmt.Println(c1 == c2) //true
}
```


型の比較可能性については以下の通り。

|             型 | 等値比較 | 大小比較 |
| --------------:|:--------:|:--------:|
|         整数型 |    可    |    可    |
| 浮動小数点数型 |    可    |    可    |
|       複素数型 |    可    |   不可   |
|         真偽型 |    可    |   不可   |
|         構造体 |    可    |   不可   |
|           配列 |    可    |   不可   |
|         文字列 |    可    |    可    |
|       Slice 型 |   不可   |   不可   |
|         Map 型 |   不可   |   不可   |
|         関数型 |   不可   |   不可   |
|     Channel 型 |    可    |   不可   |
|   Interface 型 |    可    |   不可   |
|       ポインタ |    可    |   不可   |

以下，補足。

### NaN は比較可能だが比較できない

NaN (Not a Number) は浮動小数点数型における（ゼロ除算などの）特別な状態を示す。
NaN 自体は比較可能なのだが，常に同じ結果を返すので，比較演算子は使えない。
浮動小数点数の値が NaN かどうか調べるには [`math`]`.IsNaN()` 関数を使う。

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var z float64
	nan := z / z
	fmt.Println(nan)                                   //NaN
	fmt.Println("NaN == NaN ->", nan == math.NaN())    //false
	fmt.Println("NaN != NaN ->", nan != math.NaN())    //true
	fmt.Println("NaN > NaN ->", nan > math.NaN())      //false
	fmt.Println("NaN < NaN ->", nan < math.NaN())      //false
	fmt.Println("math.IsNaN(NaN) ->", math.IsNaN(nan)) //true

}
```

### 構造体要素の型

構造体は，要素の型が全て比較可能であれば，比較可能である。

```go
package main

import "fmt"

type Number struct{ num int }

func main() {
	c1 := Number{num: 1}
	c2 := Number{num: 1}
	fmt.Println(c1 == c2) //true
}
```

### 配列要素の型

配列は，要素の型が比較可能であれば，比較可能である。
 
```go
package main

import "fmt"

func main() {
	b1 := [1]byte{1}
    b2 := b1
    fmt.Println(b1 == b2) //true

    b3 := [2]byte{1, 2}
	fmt.Println(b1 == b3) //compile error: mismatched types [1]byte and [2]byte
}
```

ちなみに上のコードの `[1]byte` と `[2]byte` は等価ではないのでご注意を。
配列と Slice の関係については拙文「[配列と Slice]({{< relref "./array-and-slice.md" >}})」を参考にどうぞ。

### ポインタの型

ポインタはインスタンスのアドレッシングを指すものだが，等価なインスタンスのポインタであれば比較可能である。

```go
package main

import "fmt"

type Number int

func main() {
	var c1 int = 1
	fmt.Println(&c1 == &c1) //true

	var c2 Number = 1
	fmt.Println(&c1 == &c2) //compile error: mismatched types *int and *Number
}
```

たとえば Slice 型や Map 型は比較可能ではないがポインタは比較できる（内容を比較しているわけではない）。

```go
package main

import "fmt"

func main() {
	c1 := []int{1}
	c2 := []int{1}
    fmt.Println(c1 == c2)   //compile error: slice can only be compared to nil
	fmt.Println(&c1 == &c2) //false
}
```

### nil と比較可能な型

`nil` は（null 参照など）ポインタ値の特別な状態を示す。
なので `nil` はポインタと比較可能である。
他に `nil` と比較可能な型は以下の通り。

- Slice 型
- Map 型
- 関数型
- Channel 型
- Interface 型

Slice 型, Map 型, 関数型は比較可能ではないが `nil` とは比較可能である。

## Interface 型の比較可能性

Interface 型は型情報と値への参照を属性として持っている。
Interface 型が参照している型を動的な型（dynamic type），参照値を動的な値（dynamic value）と呼ぶ。
動的な型も値も実行時に決まるからだ。

たとえば

```go
type Binary uint64

func (i Binary) String() string {
	return strconv.FormatUint(uint64(i), 2)
}
```

という型の定義に対して [`fmt`]`.Stringer` インタフェースを適用してみる。

```go
func main() {
	b := Binary(200)
	s := fmt.Stringer(b)
	fmt.Println(s.String())
	// Output:
	// 11001000
}
```

このときの [`fmt`]`.Stringer` インタフェースは以下のような構造になっている。

{{< fig-quote type="markdown" title="Go Data Structures: Interfaces" link="https://research.swtch.com/interfaces" lang="en" >}}
{{< fig-img src="https://research.swtch.com/gointer2.png" link="https://research.swtch.com/interfaces" >}}
{{< /fig-quote >}}

Interface 型は，動的な型が比較可能であれば，比較可能である。

### Interface 型の比較

Interface 型インスタンスの動的な型が等価で比較可能あれば値の等値性を調べられる。
更に Interface 型インスタンスの動的な型と等価な型のインスタンスとも（比較可能な型であれば）比較できる。

```go
package main

import "fmt"

func main() {
	c1 := (interface{})(1)
	c2 := (interface{})(1)
	c3 := 1
	fmt.Println(c1 == c2) //true
	fmt.Println(c1 == c3) //true
}
```

通常の型では等価でない型同士の比較はコンパイルエラーになるが Interface 型同士であれば単に `false` なるだけである。

```go
package main

import "fmt"

func main() {
	c1 := (interface{})(int(1))
	c2 := (interface{})(byte(1))
	c3 := (interface{})([]int{1})
	fmt.Println(c1 == c2) //false
	fmt.Println(c1 == c3) //false
}
```

ただし動的な型が等価でも比較可能ではない場合には（コンパイルは通るが）実行時 panic になる。

```go
package main

import "fmt"

func main() {
	c1 := (interface{})([]int{1})
	c2 := (interface{})([]int{1})
	fmt.Println(c1 == c2) //panic: runtime error: comparing uncomparable type []int
}
```

### Interface 型インスタンスが比較可能か検証する

比較結果が `false` になるのはまだしも，実行時 panic はいただけない。
Panic を回避するには比較する前に動的な型が比較可能かどうか調べる必要がある。

動的な型が比較可能かどうか調べるには [`reflect`] パッケージが使える。
たとえば，こんな感じでどうだろう。

```go {hl_lines=["9-11"]}
package main

import (
	"fmt"
	"reflect"
)

func Compare(left, right interface{}) bool {
	if !reflect.TypeOf(left).Comparable() && !reflect.TypeOf(right).Comparable() {
		return false
	}
	return left == right
}

func main() {
	c1 := (interface{})([]int{1})
	c2 := (interface{})([]int{1})
	fmt.Println(Compare(c1, c2)) //false
}
```

まぁ，実際に組み込むには（動的な型がポインタの際の処理など）もう少し工夫がいるだろう。
また，一般的に [`reflect`] は遅いと言われてるので，他の手段も考える必要があるかもしれない。

たとえば標準の [`errors`]`.Is()` 関数は

```go {hl_lines=[7]}
// Is reports whether any error in err's chain matches target.
func Is(err, target error) bool {
	if target == nil {
		return err == target
	}

	isComparable := reflectlite.TypeOf(target).Comparable()
	for {
		if isComparable && err == target {
			return true
		}
		if x, ok := err.(interface{ Is(error) bool }); ok && x.Is(target) {
			return true
		}
		// TODO: consider supporing target.Is(err). This would allow
		// user-definable predicates, but also may allow for coping with sloppy
		// APIs, thereby making it easier to get away with them.
		if err = Unwrap(err); err == nil {
			return false
		}
	}
}
```

となっていて，独自の internal package を使っているようだ。

## ブックマーク

- [Go で interface {} の中身がポインタならその参照先を取得する - Qiita](http://qiita.com/chimatter/items/b0879401d6666589ab71)
- [Sliceを含んだ構造体が等値演算子（==）でpanicを引き起こすのを回避する #golang - My External Storage](https://budougumi0617.github.io/2019/07/07/prevent-runtime-error-by-pointer/)

- [nil は nil]({{< relref "./nil-is-nil.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`math`]: https://golang.org/pkg/math/ "math - The Go Programming Language"
[`fmt`]: https://golang.org/pkg/fmt/ "fmt - The Go Programming Language"
[`reflect`]: https://golang.org/pkg/reflect/ "reflect - The Go Programming Language"
[`errors`]: https://golang.org/pkg/errors/ "errors - The Go Programming Language"

## 参考図書

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
