+++
tags = ["golang", "programming", "slice"]
description = "配列と slice との関係について。あくまでも私のための覚え書きである。"
draft = false
date = "2017-03-15T00:31:48+09:00"
update = "2017-04-04T00:51:45+09:00"
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
  tumblr = "spiegel-im-spiegel"
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
ary = 0x1040a124
0x1040a124: 0
0x1040a125: 1
0x1040a126: 2
0x1040a127: 3
```

まぁ分かりやすいよね。
今度はダンプ表示部分を別関数にしてみる。

```go
package main

import "fmt"

type Array4 [4]int8

func dump(ary Array4) {
	fmt.Printf("ary(dump) = %p\n", &ary)
	for i := 0; i < len(ary); i++ {
		fmt.Printf("%p: %v\n", &ary[i], ary[i])
	}
}

func main() {
	ary := Array4{0, 1, 2, 3}
	fmt.Printf("ary(org) = %p\n", &ary)
	dump(ary)
}
```

このコードの実行結果は以下の通り。

```text
ary(org) = 0x1040a124
ary(dump) = 0x1040a128
0x1040a128: 0
0x1040a129: 1
0x1040a12a: 2
0x1040a12b: 3
```

`dump()` 関数の引数として渡される配列がオリジナルのものと異なることが分かるだろう。
[Go 言語]では関数の引数は原則として「値渡し（call by value）」であるため，配列を渡す場合でも配列のコピーを作って渡すことになる。
配列を値渡しではなく「参照渡し（call by reference）」にしたい場合はポインタを使う。

```go
package main

import "fmt"

type Array4 [4]int8

func dump(ary *Array4) {
	fmt.Printf("ary(dump) = %p\n", ary)
	for i := 0; i < len(ary); i++ {
		fmt.Printf("%p: %v\n", &ary[i], ary[i])
	}
}

func main() {
	ary := Array4{0, 1, 2, 3}
	fmt.Printf("ary(org) = %p\n", &ary)
	dump(&ary)
}
```

結果は以下の通り。

```text
ary(org) = 0x1040a124
ary(dump) = 0x1040a124
0x1040a124: 0
0x1040a125: 1
0x1040a126: 2
0x1040a127: 3
```

## Slice は配列への参照である

次は配列を [slice] に置き換えてみよう。

```go
package main

import "fmt"

func dump(slc []int8) {
	fmt.Printf("slc(dump) = %p\n", slc)
	for i := 0; i < len(slc); i++ {
		fmt.Printf("%p: %v\n", &slc[i], slc[i])
	}
}

func main() {
	slc := []int8{0, 1, 2, 3}
	fmt.Printf("slc(org) = %p\n", slc)
	dump(slc)
}
```

配列の場合の記述の違いが分かるだろうか。
この場合 `slc` には配列 `{0, 1, 2, 3}` へのポインタがセットされる。
したがって `dump()` 関数の引数には（見かけ上）配列 `{0, 1, 2, 3}` への参照がセットされていることになる。

```text
slc(org) = 0x1040a124
slc(dump) = 0x1040a124
0x1040a124: 0
0x1040a125: 1
0x1040a126: 2
0x1040a127: 3
```

では応用として今度はこんなコードを考えてみよう。

```go
package main

import "fmt"

type Array4 [4]int8

func dumpA(ary Array4) {
	fmt.Printf("ary(dumpA) = %p\n", &ary)
	for i := 0; i < len(ary); i++ {
		fmt.Printf("%p: %v\n", &ary[i], ary[i])
	}
}

func dumpS(slc []int8) {
	fmt.Printf("slc(dumpS) = %p\n", slc)
	for i := 0; i < len(slc); i++ {
		fmt.Printf("%p: %v\n", &slc[i], slc[i])
	}
}

func main() {
	ary := Array4{0, 1, 2, 3}
	fmt.Printf("ary(org) = %p\n", &ary)
	dumpA(ary)
	slc := ary[:]
	dumpS(slc)
}
```

`slc := ary[:]` で配列 `ary` が [slice] `slc` にキャストされているのがポイントである。
実行結果は以下の通り。

```text
ary(org) = 0x1040a124
ary(dumpA) = 0x1040a128
0x1040a128: 0
0x1040a129: 1
0x1040a12a: 2
0x1040a12b: 3
slc(dumpS) = 0x1040a124
0x1040a124: 0
0x1040a125: 1
0x1040a126: 2
0x1040a127: 3
```

配列 `ary` のオリジナルのポインタ値がそのまま `slc` の値になっているのが分かると思う。
この「[slice] は配列への参照である」ということを踏まえると，こんな面白いコードも書ける。

```go
package main

import "fmt"

type Array4 [4]int8

func dumpA(ary Array4) {
	fmt.Printf("ary(dumpA) = %p\n", &ary)
	for i := 0; i < len(ary); i++ {
		fmt.Printf("%p: %v\n", &ary[i], ary[i])
	}
}

func dumpS(slc []int8) {
	fmt.Printf("pointer(dumpS) = %p\n", slc)
	fmt.Printf("size(dumpS) = %v\n", len(slc))
	fmt.Printf("capacity(dumpS) = %v\n", cap(slc))
	for i := 0; i < len(slc); i++ {
		fmt.Printf("%p: %v\n", &slc[i], slc[i])
	}
}

func main() {
	ary := Array4{0, 1, 2, 3}
	fmt.Printf("ary(org) = %p\n", &ary)
	dumpA(ary)
	slc1 := ary[0:2]
	dumpS(slc1)
	slc2 := slc1[0:4]
	dumpS(slc2)
}
```

サイズ 2 の `slc1` からサイズ4の `slc2` を取得しているのがポイント。
このコードの実行結果は以下の通り。

```text
ary(org) = 0x1040a124
ary(dumpA) = 0x1040a128
0x1040a128: 0
0x1040a129: 1
0x1040a12a: 2
0x1040a12b: 3
pointer(dumpS) = 0x1040a124
size(dumpS) = 2
capacity(dumpS) = 4
0x1040a124: 0
0x1040a125: 1
pointer(dumpS) = 0x1040a124
size(dumpS) = 4
capacity(dumpS) = 4
0x1040a124: 0
0x1040a125: 1
0x1040a126: 2
0x1040a127: 3
```

実は [slice] は配列に対してポインタとサイズと容量の3つの属性を持つオブジェクトである。
上述のコードでは配列 `ary` を反映し， `slc1` の容量が4となるため `slc2` ではサイズ4の [slice] が作れるわけだ。

たとえば `slc1 := ary[2:4]` と書き換えると

```go
package main

import "fmt"

type Array4 [4]int8

func dumpA(ary Array4) {
	fmt.Printf("ary(dumpA) = %p\n", &ary)
	for i := 0; i < len(ary); i++ {
		fmt.Printf("%p: %v\n", &ary[i], ary[i])
	}
}

func dumpS(slc []int8) {
	fmt.Printf("pointer(dumpS) = %p\n", slc)
	fmt.Printf("size(dumpS) = %v\n", len(slc))
	fmt.Printf("capacity(dumpS) = %v\n", cap(slc))
	for i := 0; i < len(slc); i++ {
		fmt.Printf("%p: %v\n", &slc[i], slc[i])
	}
}

func main() {
	ary := Array4{0, 1, 2, 3}
	fmt.Printf("ary(org) = %p\n", &ary)
	dumpA(ary)
	slc1 := ary[2:4]
	dumpS(slc1)
	slc2 := slc1[0:4]
	dumpS(slc2)
}
```
`slc1` の容量が変わるため，以下のように実行時 panic になる。

```text
ary(org) = 0x1040a124
ary(dumpA) = 0x1040a128
0x1040a128: 0
0x1040a129: 1
0x1040a12a: 2
0x1040a12b: 3
pointer(dumpS) = 0x1040a126
size(dumpS) = 2
capacity(dumpS) = 2
0x1040a126: 2
0x1040a127: 3
panic: runtime error: slice bounds out of range
```

このように配列と [slice] の関係が分かると `append()` 関数の挙動も理解しやすくなる。

ところで先ほど [slice] は「ポインタとサイズと容量の3つの属性を持つオブジェクト」と書いた。
つまり厳密に言えば，関数の引数に [slice] をセットするということは「ポインタとサイズと容量の3つの属性を持つオブジェクト」を値渡しでセットしているということになる。
たとえば以下のようなコードを考える。

```go
package main

import "fmt"

type Array4 [4]int8

func dumpS(slc []int8) {
	fmt.Printf("pointer(dumpS) = %p\n", slc)
	fmt.Printf("size(dumpS) = %v\n", len(slc))
	fmt.Printf("capacity(dumpS) = %v\n", cap(slc))
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
pointer(dumpS) = 0x1040a128
size(dumpS) = 4
capacity(dumpS) = 8
0x1040a128: 0
0x1040a129: 1
0x1040a12a: 2
0x1040a12b: 3
pointer(dumpS) = 0x1040a128
size(dumpS) = 5
capacity(dumpS) = 8
0x1040a128: 0
0x1040a129: 1
0x1040a12a: 2
0x1040a12b: 3
0x1040a12c: 4
pointer(dumpS) = 0x1040a128
size(dumpS) = 4
capacity(dumpS) = 8
0x1040a128: 0
0x1040a129: 1
0x1040a12a: 2
0x1040a12b: 3
pointer(dumpS) = 0x1040a128
size(dumpS) = 5
capacity(dumpS) = 8
0x1040a128: 0
0x1040a129: 1
0x1040a12a: 2
0x1040a12b: 3
0x1040a12c: 4
```

つまり `addS()` 関数に渡す `slc` は値渡しなので `addS()` 関数内で `slc` のサイズが変わっても関数の呼び出し元には反映されないことになる（配列自体には値がセットされている）。
`append()` 関数実行後は必ず状態が変わるため正しく [slice] の「値」を更新する必要がある。

## 配列の複製

配列を明示的に複製して使いたい場合がある。
[Go 言語]では配列の複製はとても簡単である。

```go
package main

import "fmt"

type Array4 [4]int8

func dump(ary *Array4) {
	fmt.Printf("ary(dump) = %p\n", ary)
	for i := 0; i < len(ary); i++ {
		fmt.Printf("%p: %v\n", &ary[i], ary[i])
	}
}

func main() {
	ary := Array4{0, 1, 2, 3}
	var ary2 Array4
	fmt.Printf("ary(org) = %p\n", &ary)
	dump(&ary)
	ary2 = ary
	dump(&ary2)
    if ary == ary2 {
		fmt.Println("ary == ary2")
	} else {
		fmt.Println("ary != ary2")
	}
}
```

実行結果は以下の通り。

```text
ary(org) = 0x1040a124
ary(dump) = 0x1040a124
0x1040a124: 0
0x1040a125: 1
0x1040a126: 2
0x1040a127: 3
ary(dump) = 0x1040a128
0x1040a128: 0
0x1040a129: 1
0x1040a12a: 2
0x1040a12b: 3
ary == ary2
```

`ary` と `ary2` が同じ内容の異なるインスタンス（instance）であることが分かると思う。
また配列同士の比較も同じ型であれば単純である。
`[3]int8` と `[4]int8` は異なる型と見なされるため単純比較はできない。

一方， [slice] の複製が欲しい場合は `copy()` 関数を使う。

```go
package main

import (
	"fmt"
	"reflect"
)

type Array4 [4]int8

func dumpS(slc []int8) {
	fmt.Printf("pointer(dumpS) = %p\n", slc)
	fmt.Printf("size(dumpS) = %v\n", len(slc))
	fmt.Printf("capacity(dumpS) = %v\n", cap(slc))
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

コピー先の `slc2` について `make()` 関数であらかじめサイズと容量を確保しておくのがポイント。
実行結果は以下の通り。

```text
pointer(dumpS) = 0x1040a124
size(dumpS) = 4
capacity(dumpS) = 4
0x1040a124: 0
0x1040a125: 1
0x1040a126: 2
0x1040a127: 3
pointer(dumpS) = 0x1040a144
size(dumpS) = 4
capacity(dumpS) = 4
0x1040a144: 0
0x1040a145: 1
0x1040a146: 2
0x1040a147: 3
slc1 == slc2
```

[slice] 同士を比較するのも単純ではないが， [`reflect`].`DeepEqual()` 関数が使える。
ちなみに 宣言構文を使ってもっと単純に

```go
package main

import (
	"fmt"
	"reflect"
)

type Array4 [4]int8

func dumpS(slc []int8) {
	fmt.Printf("pointer(dumpS) = %p\n", slc)
	fmt.Printf("size(dumpS) = %v\n", len(slc))
	fmt.Printf("capacity(dumpS) = %v\n", cap(slc))
	for i := 0; i < len(slc); i++ {
		fmt.Printf("%p: %v\n", &slc[i], slc[i])
	}
}

func main() {
	slc1 := []int8{0, 1, 2, 3}
	dumpS(slc1)
    slc2 := slc1
	dumpS(slc2)
	if reflect.DeepEqual(slc1, slc2) {
		fmt.Println("slc1 == slc2")
	} else {
		fmt.Println("slc1 != slc2")
	}
}
```

とすればいいじゃない，と思われるかもしれないが，結果は以下の通り。

```text
pointer(dumpS) = 0x1040a124
size(dumpS) = 4
capacity(dumpS) = 4
0x1040a124: 0
0x1040a125: 1
0x1040a126: 2
0x1040a127: 3
pointer(dumpS) = 0x1040a124
size(dumpS) = 4
capacity(dumpS) = 4
0x1040a124: 0
0x1040a125: 1
0x1040a126: 2
0x1040a127: 3
slc1 == slc2
```

`slc1` と `slc2` の指す配列が同じになり複製できない。

## ブックマーク

- [Go Slices: usage and internals - The Go Blog](https://blog.golang.org/go-slices-usage-and-internals)
- [Arrays, slices (and strings): The mechanics of 'append' - The Go Blog](https://blog.golang.org/slices)
- [Go のスライスでハマッたところ - Block Rockin’ Codes](http://jxck.hatenablog.com/entry/golang-slice-internals2)
- [golang で string を []byte にキャストしてもメモリコピーが走らない方法を考えてみる - Qiita](http://qiita.com/mattn/items/176459728ff4f854b165) : ネタ記事
- [関数とポインタ]({{< relref "golang/function-and-pointer.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[slice]: http://golang.org/ref/spec#Slice_types
[`reflect`]: https://golang.org/pkg/reflect/ "reflect - The Go Programming Language"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-20</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4865940391/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4865940391.09._SCTHUMBZZZ_.jpg"  alt="Kotlinスタートブック -新しいAndroidプログラミング"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4839959234/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4839959234.09._SCTHUMBZZZ_.jpg"  alt="Docker実戦活用ガイド"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274218961/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274218961.09._SCTHUMBZZZ_.jpg"  alt="グッド・マス ギークのための数・論理・計算機科学"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-13">2016-07-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
