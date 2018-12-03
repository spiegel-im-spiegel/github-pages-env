+++
tags = ["golang", "programming", "slice"]
description = "配列と slice との関係について。あくまでも私のための覚え書きである。"
draft = false
date = "2017-03-15T00:31:48+09:00"
update = "2018-02-07T15:29:58+09:00"
title = "配列と Slice"

[author]
  license = "by-sa"
  url = "http://www.baldanders.info/spiegel/profile/"
  linkedin = "spiegelimspiegel"
  twitter = "spiegel_2007"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  avatar = "/images/avatar.jpg"
  flattr = "spiegel"
  name = "Spiegel"
  facebook = "spiegel.im.spiegel"
  flickr = "spiegel"
  tumblr = ""

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

{{< highlight go "hl_lines=24" >}}
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
{{< /highlight >}}

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

{{% div-box %}}
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


{{< highlight go "hl_lines=15-16" >}}
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
{{< /highlight >}}

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

{{< highlight go "hl_lines=18-22" >}}
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
{{< /highlight >}}

とやっても

```text
prog.go:18:10: invalid operation: slc1 == slc2 (slice can only be compared to nil)
```

という感じでコンパイルエラーになる[^ce1]。

[^ce1]: 同様に比較演算子が使えない基本型としては [map] と関数値（function value）がある。

[slice] の内容を比較したいのであれば [`reflect`]`.DeepEqual()` 関数が使える[^cmp1]。

[^cmp1]: byte 型の [slice] であれば [`bytes`]`.Compare()` を使って比較できる。

{{< highlight go "hl_lines=21-25" >}}
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
{{< /highlight >}}

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

{{< highlight go "hl_lines=21-25 28-32" >}}
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
{{< /highlight >}}

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

{{< highlight go "hl_lines=17-21" >}}
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
        fmt.Println("slc1 == slc2")
    } else {
        fmt.Println("slc1 != slc2")
    }
}
{{< /highlight >}}

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
slc1 != slc2
```

## ブックマーク

- [Go Slices: usage and internals - The Go Blog](https://blog.golang.org/go-slices-usage-and-internals)
- [Arrays, slices (and strings): The mechanics of 'append' - The Go Blog](https://blog.golang.org/slices)
- [Go のスライスでハマッたところ - Block Rockin’ Codes](http://jxck.hatenablog.com/entry/golang-slice-internals2)
- [golang で string を []byte にキャストしてもメモリコピーが走らない方法を考えてみる - Qiita](http://qiita.com/mattn/items/176459728ff4f854b165) : ネタ記事
- [golangのequalityの評価について - podhmo's diary](http://pod.hatenablog.com/entry/2016/07/30/204357)
- [スライスを返すときはcapに注意しよう - Qiita](https://qiita.com/methane/items/2cc4e4a23172f6f9b993)

- [Map の話]({{< relref "map.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[slice]: http://golang.org/ref/spec#Slice_types
[map]: http://golang.org/ref/spec#Map_types
[`reflect`]: https://golang.org/pkg/reflect/ "reflect - The Go Programming Language"
[`bytes`]: https://golang.org/pkg/bytes/ "bytes - The Go Programming Language"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-20</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4865940391/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4865940391.09._SCTHUMBZZZ_.jpg"  alt="Kotlinスタートブック -新しいAndroidプログラミング"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4839959234/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4839959234.09._SCTHUMBZZZ_.jpg"  alt="Docker実戦活用ガイド"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274218961/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274218961.09._SCTHUMBZZZ_.jpg"  alt="グッド・マス ギークのための数・論理・計算機科学"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-13">2016-07-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
