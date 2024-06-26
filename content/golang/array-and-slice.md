+++
tags = ["golang", "programming", "slice"]
description = "配列と slice との関係について。あくまでも私のための覚え書きである。"
date = "2017-03-15T00:31:48+09:00"
title = "配列と Slice"

[scripts]
  mathjax = false
  mermaidjs = false
+++

以下の記事を見て思わず膝を打った。

- [Goのarrayとsliceを理解するときがきた - Qiita](http://qiita.com/seihmd/items/d9bc98a4f4f606ecaef7)

なるほど！ こういう風に説明すればいいのか。
というわけで，私も便乗してみる。
あくまでも私のための覚え書きである。

## 配列は常に「値」である

まずはこんなコードを書いてみる。

```go
package main

import "fmt"

type Array4 [4]int8

func main() {
    ary := Array4{0, 1, 2, 3}
    fmt.Printf("ary = %p\n", &ary)
    for i := 0; i < len(ary); i++ {
        fmt.Printf("%p: %v\n", &ary[i], ary[i])
    }
}
```

`[4]int8` が「型」であることを意識してもらうために敢えて `Array4` という型を宣言している。
実行結果は以下の通り。

```text
ary = 0x10410020
0x10410020: 0
0x10410021: 1
0x10410022: 2
0x10410023: 3
```

まぁ分かりやすいよね。
今度はダンプ表示部分を別関数にしてみる。

```go
package main

import "fmt"

type Array4 [4]int8

func dumpA(ary Array4) {
    fmt.Printf("dump(ary) = %p\n", &ary)
    for i := 0; i < len(ary); i++ {
        fmt.Printf("%p: %v\n", &ary[i], ary[i])
    }
}

func main() {
    ary := Array4{0, 1, 2, 3}
    fmt.Printf("ary = %p\n", &ary)
    dumpA(ary)
}
```

このコードの実行結果は以下の通り。

```text
ary = 0x10410020
dump(ary) = 0x10410040
0x10410040: 0
0x10410041: 1
0x10410042: 2
0x10410043: 3
```

`dumpA()` 関数の引数として渡される配列（のポインタ）がオリジナルのものと異なることが分かるだろう。
[Go 言語]では関数の引数は原則として「値渡し（call by value）」であるため，配列を渡す場合でも配列のコピーを作って渡すことになる。
配列の値（＝コピー）を渡すのではなく配列の実体を渡すにはポインタを使う。

```go
package main

import "fmt"

type Array4 [4]int8

func dumpA(ary *Array4) {
    fmt.Printf("dump(ary) = %p\n", ary)
    for i := 0; i < len(ary); i++ {
        fmt.Printf("%p: %v\n", &ary[i], ary[i])
    }
}

func main() {
    ary := Array4{0, 1, 2, 3}
    fmt.Printf("ary = %p\n", &ary)
    dumpA(&ary)
}
```

結果は以下の通り。

```text
ary = 0x10410020
dump(ary) = 0x10410020
0x10410020: 0
0x10410021: 1
0x10410022: 2
0x10410023: 3
```

## [slice] は配列を参照するオブジェクトである

次は配列を [slice] に置き換えてみよう。

```go
package main

import "fmt"

func dumpS(slc []int8) {
    fmt.Printf("dump(slc) = %p\n", slc)
    for i := 0; i < len(slc); i++ {
        fmt.Printf("%p: %v\n", &slc[i], slc[i])
    }
}

func main() {
    slc := []int8{0, 1, 2, 3}
    fmt.Printf("slc = %p\n", slc)
    dumpS(slc)
}
```

配列の場合の記述の違いが分かるだろうか。
この場合 `slc` には配列 `{0, 1, 2, 3}` へのポインタがセットされる。
したがって `dumpS()` 関数の引数には（見かけ上）配列 `{0, 1, 2, 3}` への参照がセットされていることになる。

```text
slc = 0x10410020
dump(slc) = 0x10410020
0x10410020: 0
0x10410021: 1
0x10410022: 2
0x10410023: 3
```

では応用として今度はこんなコードを考えてみよう。

```go
package main

import "fmt"

type Array4 [4]int8

func dumpS(slc []int8) {
    fmt.Printf("dump(slc) = %p\n", slc)
    for i := 0; i < len(slc); i++ {
        fmt.Printf("%p: %v\n", &slc[i], slc[i])
    }
}

func main() {
    ary := Array4{0, 1, 2, 3}
    fmt.Printf("ary = %p\n", &ary)
    slc := ary[:]
    dumpS(slc)
}
```

`slc := ary[:]` で配列 `ary` が [slice] `slc` にキャストされているのがポイントである。
実行結果は以下の通り。

```text
aary = 0x10410020
dump(slc) = 0x10410020
0x10410020: 0
0x10410021: 1
0x10410022: 2
0x10410023: 3
```

配列 `ary` のオリジナルのポインタ値がそのまま `slc` の配列への参照値になっているのが分かると思う。
これを踏まえると以下の面白いコードも書ける。

```go
package main

import "fmt"

type Array4 [4]int8

func dumpA(ary *Array4) {
    fmt.Printf("dump(ary) = %p\n", ary)
    for i := 0; i < len(ary); i++ {
        fmt.Printf("%p: %v\n", &ary[i], ary[i])
    }
}

func dumpS(slc []int8) {
    fmt.Printf("dump(slc, sz, cap) = %p, %d, %d\n", slc, len(slc), cap(slc))
    for i := 0; i < len(slc); i++ {
        fmt.Printf("%p: %v\n", &slc[i], slc[i])
    }
}

func main() {
    ary := Array4{0, 1, 2, 3}
    dumpA(&ary)
    slc1 := ary[0:2]
    dumpS(slc1)
    slc2 := slc1[0:4]
    dumpS(slc2)
}
```

サイズ 2 の `slc1` からサイズ4の `slc2` を取得しているのがポイント。
このコードの実行結果は以下の通り。

```text
ddump(ary) = 0x10410020
0x10410020: 0
0x10410021: 1
0x10410022: 2
0x10410023: 3
dump(slc, sz, cap) = 0x10410020, 2, 4
0x10410020: 0
0x10410021: 1
dump(slc, sz, cap) = 0x10410020, 4, 4
0x10410020: 0
0x10410021: 1
0x10410022: 2
0x10410023: 3
```

実は [slice] は配列に対する**ポインタとサイズと容量の3つの属性を持つオブジェクト**である。
上述のコードでは配列 `ary` のサイズが反映されて `slc1` の容量が4となるため `slc2` ではサイズ4の [slice] を作ることができるわけだ。

ここで，たとえば `slc1 := ary[2:4]` と書き換えると

```go {hl_lines=[24]}
package main

import "fmt"

type Array4 [4]int8

func dumpA(ary * Array4) {
    fmt.Printf("dump(ary) = %p\n", ary)
    for i := 0; i < len(ary); i++ {
        fmt.Printf("%p: %v\n", &ary[i], ary[i])
    }
}

func dumpS(slc []int8) {
    fmt.Printf("dump(slc, sz, cap) = %p, %d, %d\n", slc, len(slc), cap(slc))
    for i := 0; i < len(slc); i++ {
        fmt.Printf("%p: %v\n", &slc[i], slc[i])
    }
}

func main() {
    ary := Array4{0, 1, 2, 3}
    dumpA(&ary)
    slc1 := ary[2:4]
    dumpS(slc1)
    slc2 := slc1[0:4]
    dumpS(slc2)
}
```

`slc1` の容量が変わるため，以下のように実行時 panic になる。

```text
dump(ary) = 0x10410020
0x10410020: 0
0x10410021: 1
0x10410022: 2
0x10410023: 3
dump(slc, sz, cap) = 0x10410022, 2, 2
0x10410022: 2
0x10410023: 3
panic: runtime error: slice bounds out of range
```

このように配列と [slice] の関係が分かると `append()` 関数の挙動も理解しやすくなる。

{{% div-box type="markdown" %}}
**【追記】** たとえば，配列 `ary := Array4{0, 1, 2, 3}` の先頭2要素を slice として切り出す場合，普通は `ary[0:2]` でいいのだが，容量を含めて2要素のみとしたい場合は `ary[0:2:2]` と記述する。

```go
package main

import "fmt"

type Array4 [4]int8

func dumpS(slc []int8) {
    fmt.Printf("dumpS(slc, sz, cap) = %p, %d, %d\n", slc, len(slc), cap(slc))
    for i := 0; i < len(slc); i++ {
        fmt.Printf("%p: %v\n", &slc[i], slc[i])
    }
}

func main() {
    ary := Array4{0, 1, 2, 3}
    slc1 := ary[0:2:2]
    dumpS(slc1)
}
```

このコードの実行結果は以下の通り。

```text
dumpS(slc, sz, cap) = 0x10414020, 2, 2
0x10414020: 0
0x10414021: 1
```

`ary[0:2:2]` の3番目の要素は（容量の）サイズではなくインデックス値である点に注意。
たとえば `ary` の2番目と3番目の要素を切り出そうとして `ary[1:3:2]` などと書くとコンパイルエラーになる。

```text
invalid slice index: 3 > 2
```
{{% /div-box %}}


ところで先ほど [slice] は「ポインタとサイズと容量の3つの属性を持つオブジェクト」と書いた。
つまり厳密に言えば，関数の引数に [slice] をセットするということは「ポインタとサイズと容量の3つの属性を持つオブジェクト」を値渡しでセットしているということになる。
たとえば以下のようなコードを考える。

```go
package main

import "fmt"

func dumpS(slc []int8) {
    fmt.Printf("dump(slc) = %p\n", slc)
    for i := 0; i < len(slc); i++ {
        fmt.Printf("%p: %v\n", &slc[i], slc[i])
    }
}

func addS(slc []int8, e int8) {
    slc = append(slc, e)
    dumpS(slc)
}

func main() {
    slc := make([]int8, 4, 8)
    slc[0] = 0
    slc[1] = 1
    slc[2] = 2
    slc[3] = 3
    dumpS(slc)
    addS(slc, 4)
    dumpS(slc)
    slc2 := slc[0:5]
    dumpS(slc2)
}
```

杜撰なコードではあるが，サイズ4容量8の配列に5番目の要素を `append()` しても内部の配列そのものは更新されないため動きとしては問題ないように見える。
しかし実際には以下のような実行結果になる。

```text
dump(slc) = 0x10410020
0x10410020: 0
0x10410021: 1
0x10410022: 2
0x10410023: 3
dump(slc) = 0x10410020
0x10410020: 0
0x10410021: 1
0x10410022: 2
0x10410023: 3
0x10410024: 4
dump(slc) = 0x10410020
0x10410020: 0
0x10410021: 1
0x10410022: 2
0x10410023: 3
dump(slc) = 0x10410020
0x10410020: 0
0x10410021: 1
0x10410022: 2
0x10410023: 3
0x10410024: 4
```

つまり `addS()` 関数に渡す `slc` は値渡しなので `addS()` 関数内で `slc` のサイズが変わっても関数の呼び出し元には反映されないことになる（ただし配列自体には値がセットされている）。
`append()` 関数実行後は必ず状態が変わるため正しく [slice] の「値」を更新する必要がある。

## 配列の複製と比較

[Go 言語]では配列は値であるため，代入時に自動的に複製が発生する。
また同じ型であれば `==` 演算子で配列の内容を比較できる。

```go
package main

import "fmt"

type Array4 [4]int8

func dumpA(ary *Array4) {
    fmt.Printf("dump(ary) = %p\n", ary)
    for i := 0; i < len(ary); i++ {
        fmt.Printf("%p: %v\n", &ary[i], ary[i])
    }
}

func main() {
    ary := Array4{0, 1, 2, 3}
    var ary2 Array4
    dumpA(&ary)
    ary2 = ary
    dumpA(&ary2)
    if ary == ary2 {
        fmt.Println("ary == ary2")
    } else {
        fmt.Println("ary != ary2")
    }
}
```

実行結果は以下の通り。

```text
dump(ary) = 0x10410020
0x10410020: 0
0x10410021: 1
0x10410022: 2
0x10410023: 3
dump(ary) = 0x10410024
0x10410024: 0
0x10410025: 1
0x10410026: 2
0x10410027: 3
ary == ary2
```

ただし配列の型が異なる場合（たとえば `[3]int8` と `[4]int8`）は単純比較はできないので注意が必要である。

## [slice] の複製と比較

一方， [slice] は配列のポインタを属性値として持っているだけなので代入を行っても配列自体は複製されない。
[slice] の複製が欲しい場合は `copy()` 関数を使う。

```go {hl_lines=["15-16"]}
package main

import "fmt"

func dumpS(slc []int8) {
    fmt.Printf("dump(slc) = %p\n", slc)
    for i := 0; i < len(slc); i++ {
        fmt.Printf("%p: %v\n", &slc[i], slc[i])
    }
}

func main() {
    slc1 := []int8{0, 1, 2, 3}
    dumpS(slc1)
    slc2 := make([]int8, len(slc1), cap(slc1))
    copy(slc2, slc1)
    dumpS(slc2)
}
```

コピー先の `slc2` について `make()` 関数であらかじめサイズと容量を確保しておくのがポイント。
実行結果は以下の通り。

```text
dump(slc) = 0x10410020
0x10410020: 0
0x10410021: 1
0x10410022: 2
0x10410023: 3
dump(slc) = 0x10410024
0x10410024: 0
0x10410025: 1
0x10410026: 2
0x10410027: 3
```

なお [slice] インスタンス同士は `==` 演算子による比較ができない。

```go {hl_lines=["18-22"]}
package main

import "fmt"

func dumpS(slc []int8) {
    fmt.Printf("dump(slc) = %p\n", slc)
    for i := 0; i < len(slc); i++ {
        fmt.Printf("%p: %v\n", &slc[i], slc[i])
    }
}

func main() {
    slc1 := []int8{0, 1, 2, 3}
    dumpS(slc1)
    slc2 := make([]int8, len(slc1), cap(slc1))
    copy(slc2, slc1)
    dumpS(slc2)
    if slc1 == slc2 {
        fmt.Println("slc1 == slc2")
    } else {
        fmt.Println("slc1 != slc2")
    }
}
```

とやっても

```text
prog.go:18:10: invalid operation: slc1 == slc2 (slice can only be compared to nil)
```

という感じでコンパイルエラーになる[^ce1]。

[^ce1]: 同様に比較演算子が使えない基本型としては [map] と関数値（function value）がある。

[slice] の内容を比較したいのであれば [`reflect`]`.DeepEqual()` 関数が使える[^cmp1]。

[^cmp1]: byte 型の [slice] であれば [`bytes`]`.Compare()` を使って比較できる。

```go {hl_lines=["21-25"]}
package main

import (
    "fmt"
    "reflect"
)

func dumpS(slc []int8) {
    fmt.Printf("dump(slc) = %p\n", slc)
    for i := 0; i < len(slc); i++ {
        fmt.Printf("%p: %v\n", &slc[i], slc[i])
    }
}

func main() {
    slc1 := []int8{0, 1, 2, 3}
    dumpS(slc1)
    slc2 := make([]int8, len(slc1), cap(slc1))
    copy(slc2, slc1)
    dumpS(slc2)
    if reflect.DeepEqual(slc1, slc2) {
        fmt.Println("slc1 == slc2")
    } else {
        fmt.Println("slc1 != slc2")
    }

}
```

とすれば以下の結果になる。

```text
dump(slc) = 0x10410020
0x10410020: 0
0x10410021: 1
0x10410022: 2
0x10410023: 3
dump(slc) = 0x10410024
0x10410024: 0
0x10410025: 1
0x10410026: 2
0x10410027: 3
slc1 == slc2
```

需要があるかどうか分からないが， [slice] が参照している配列のインスタンスが同一であるかどうか調べるには [`reflect`]`.ValueOf()` 関数で値（＝配列）を取得し，そのポインタ値を `==` 演算子で比較する。

```go {hl_lines=["21-25", "28-32"]}
package main

import (
    "fmt"
    "reflect"
)

func dumpS(slc []int8) {
    fmt.Printf("dump(slc) = %p\n", slc)
    for i := 0; i < len(slc); i++ {
        fmt.Printf("%p: %v\n", &slc[i], slc[i])
    }
}

func main() {
    slc1 := []int8{0, 1, 2, 3}
    dumpS(slc1)
    slc2 := make([]int8, len(slc1), cap(slc1))
    copy(slc2, slc1)
    dumpS(slc2)
    if reflect.ValueOf(slc1).Pointer() == reflect.ValueOf(slc2).Pointer() {
        fmt.Println("slc1 == slc2")
    } else {
        fmt.Println("slc1 != slc2")
    }
    slc3 := slc1
    dumpS(slc3)
    if reflect.ValueOf(slc1).Pointer() == reflect.ValueOf(slc3).Pointer() {
        fmt.Println("slc1 == slc3")
    } else {
        fmt.Println("slc1 != slc3")
    }
}
```

結果は以下の通り。

```text
dump(slc) = 0x10410020
0x10410020: 0
0x10410021: 1
0x10410022: 2
0x10410023: 3
dump(slc) = 0x10410024
0x10410024: 0
0x10410025: 1
0x10410026: 2
0x10410027: 3
slc1 != slc2
dump(slc) = 0x10410020
0x10410020: 0
0x10410021: 1
0x10410022: 2
0x10410023: 3
slc1 == slc3
```

[slice] のポインタ値を比較すればいいぢゃん，と思うかもしれないが，その場合は単に [slice] インスタンスが同一かどうかを比較しているに過ぎない。

```go {hl_lines=["17-21"]}
package main

import "fmt"

func dumpS(slc []int8) {
    fmt.Printf("dump(slc) = %p\n", slc)
    for i := 0; i < len(slc); i++ {
        fmt.Printf("%p: %v\n", &slc[i], slc[i])
    }
}

func main() {
    slc1 := []int8{0, 1, 2, 3}
    dumpS(slc1)
    slc2 := slc1
    dumpS(slc2)
    if &slc1 == &slc2 {
        fmt.Println("&slc1 == &slc2")
    } else {
        fmt.Println("&slc1 != &slc2")
    }
}
```

なので，結果は以下の通り。

```text
dump(slc) = 0x10410020
0x10410020: 0
0x10410021: 1
0x10410022: 2
0x10410023: 3
dump(slc) = 0x10410020
0x10410020: 0
0x10410021: 1
0x10410022: 2
0x10410023: 3
&slc1 != &slc2
```

### 【2023-08-10 追記】 slices 標準パッケージを使う

[Go][Go 言語] 1.21 から [`slices`] 標準パッケージが追加された。
これは [slice] の操作を Generics を使って定義したもので，たとえば [slice] の複製や比較を行うメソッドは

```go
// Clone returns a copy of the slice.
// The elements are copied using assignment, so this is a shallow clone.
func Clone[S ~[]E, E any](s S) S {
    // Preserve nil in case it matters.
    if s == nil {
        return nil
    }
    return append(S([]E{}), s...)
}
```

```go
// Equal reports whether two slices are equal: the same length and all
// elements equal. If the lengths are different, Equal returns false.
// Otherwise, the elements are compared in increasing index order, and the
// comparison stops at the first unequal pair.
// Floating point NaNs are not considered equal.
func Equal[S ~[]E, E comparable](s1, s2 S) bool {
    if len(s1) != len(s2) {
        return false
    }
    for i := range s1 {
        if s1[i] != s2[i] {
            return false
        }
    }
    return true
}
```

といった感じに定義されている。
これを使えば前節のコードは

```go {hl_lines=[6,19,"26-30"]}
package main

import (
    "fmt"
    "reflect"
    "slices"
)

func dumpS(slc []int8) {
    fmt.Printf("dump(slc) = %p\n", slc)
    for i := 0; i < len(slc); i++ {
        fmt.Printf("%p: %v\n", &slc[i], slc[i])
    }
}

func main() {
    slc1 := []int8{0, 1, 2, 3}
    dumpS(slc1)
    slc2 := slices.Clone(slc1)
    dumpS(slc2)
    if reflect.ValueOf(slc1).Pointer() == reflect.ValueOf(slc2).Pointer() {
        fmt.Println("slc1.Pointer == slc2.Pointer")
    } else {
        fmt.Println("slc1.Pointer != slc2.Pointer")
    }
    if slices.Equal(slc1, slc2) {
        fmt.Println("slc1 == slc2")
    } else {
        fmt.Println("slc1 != slc2")
    }
}
```

てな感じに書き直すことができる。
これを実行すると

```text
dump(slc) = 0xc000012028
0xc000012028: 0
0xc000012029: 1
0xc00001202a: 2
0xc00001202b: 3
dump(slc) = 0xc000012050
0xc000012050: 0
0xc000012051: 1
0xc000012052: 2
0xc000012053: 3
slc1.Pointer != slc2.Pointer
slc1 == slc2
```

といった出力になる。
[`slices`]`.Clone()` 関数によって新たにコピーが生成されているのが確認できるだろう。
また [`slices`]`.Equal()` 関数が配列の内容（値）を比較している点にも注目してほしい。

## ブックマーク

- [Go Slices: usage and internals - The Go Blog](https://blog.golang.org/go-slices-usage-and-internals)
- [Arrays, slices (and strings): The mechanics of 'append' - The Go Blog](https://blog.golang.org/slices)
- [Go のスライスでハマッたところ - Block Rockin’ Codes](http://jxck.hatenablog.com/entry/golang-slice-internals2)
- [golang で string を []byte にキャストしてもメモリコピーが走らない方法を考えてみる - Qiita](http://qiita.com/mattn/items/176459728ff4f854b165) : ネタ記事
- [golangのequalityの評価について - podhmo's diary](http://pod.hatenablog.com/entry/2016/07/30/204357)
- [スライスを返すときはcapに注意しよう - Qiita](https://qiita.com/methane/items/2cc4e4a23172f6f9b993)
- [SliceTricks · golang/go Wiki · GitHub](https://github.com/golang/go/wiki/SliceTricks)
- [Go 言語でスライスから要素を消すには](https://zenn.dev/mattn/articles/31dfed3c89956d)

- [Map の話]({{< relref "map.md" >}})

[Go 言語]: https://go.dev/ "The Go Programming Language"
[slice]: https://go.dev/ref/spec#Slice_types
[map]: https://go.dev/ref/spec#Map_types
[`reflect`]: https://pkg.go.dev/reflect/ "reflect - The Go Programming Language"
[`bytes`]: https://pkg.go.dev/bytes/ "bytes - The Go Programming Language"
[`slices`]: https://pkg.go.dev/slices "slices package - slices - Go Packages"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
