+++
date = "2017-04-07T20:01:34+09:00"
update = "2018-04-10T19:45:34+09:00"
title = "ソートを使う"
description = "ソートをアルゴリズムまで言及すると非常に深いテーマになるのだが，今回は標準の sort パッケージの使い方に絞って「こんな感じ」で説明していく。"
draft = false
tags = ["golang", "programming", "sort"]

[author]
  flickr = "spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  avatar = "/images/avatar.jpg"
  name = "Spiegel"
  facebook = "spiegel.im.spiegel"
  github = "spiegel-im-spiegel"
  url = "https://baldanders.info/spiegel/profile/"
  linkedin = "spiegelimspiegel"
  flattr = ""
  tumblr = ""
  twitter = "spiegel_2007"
+++

今回はソート（sort）のお話。

プログラマでソートを知らない人はいないだろうが，一応説明しておくと，あるデータの集合を一定の規則に従って並べ替えることを指す。
日本語では「整列」と呼んだりするらしい。
ソートをアルゴリズムまで言及すると非常に深いテーマになるのだが，今回は標準の [`sort`] パッケージの使い方に絞って「こんな感じ」で説明していく。

なお，この記事で紹介するコードは [`sort`] パッケージのドキュメントに書かれているものを流用している。
[Go 言語]のコンパイラ・コードは MIT ライセンスで提供されているのでご注意を。

## 基本型データ列のソート

[`sort`] パッケージでは基本型の int, float64, string についてはソート関数が用意されている。

たとえば `{0.055, 0.815, 1.0, 0.107}` というデータ列があるとしよう。
これを昇順（小さい値から大きい値へ順に並べること）で並べることを考える。
この場合は [`sort`].`Float64s()` 関数を使えば簡単である。
コードにするとこんな感じ。

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    fset := []float64{0.055, 0.815, 1.0, 0.107}
    for _, f := range fset {
        fmt.Printf("%f ", f)
    }
    fmt.Print("\n")
    sort.Float64s(fset)
    for _, f := range fset {
        fmt.Printf("%f ", f)
    }
}
```

結果はこんな感じになる。

```text
$ go run sort1.go
0.055000 0.815000 1.000000 0.107000
0.055000 0.107000 0.815000 1.000000
```

では，降順（大きい値から小さい値へ順に並べること）で並べるにはどうすればいいだろう。
これはちょっとだけ面倒くさくなる。

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    fset := []float64{0.055, 0.815, 1.0, 0.107}
    for _, f := range fset {
        fmt.Printf("%f ", f)
    }
    fmt.Print("\n")
    sort.Sort(sort.Reverse(sort.Float64Slice(fset)))
    for _, f := range fset {
        fmt.Printf("%f ", f)
    }
}
```

まず [`sort`].`Float64Slice` は `[]float64` を示す型である。

```go
type Float64Slice []float64
```

この型が示すデータ集合を [`sort`].`Sort()` 関数で並べ替えるのだが，並べ替えの規則を [`sort`].`Reverse()` 関数で反転させている。
実行結果はこんな感じでちゃんと降順になっているのが分かるだろう。

```text
$ go run sort2.go
0.055000 0.815000 1.000000 0.107000
1.000000 0.815000 0.107000 0.055000
```

実は最初に出た [`sort`].`Float64s()` 関数は内部で [`sort`].`Sort()` 関数を呼んでいる。

```go
// Float64s sorts a slice of float64s in increasing order.
func Float64s(a []float64) { Sort(Float64Slice(a)) }
```

で， [`sort`].`Sort()` 関数の内部では [`sort`].`Float64Slice` に紐付く `Len()`, `Less()`, `Swap()` 各メソッドが呼ばれている。
`Len()`, `Less()`, `Swap()` 各メソッドを持つ [interface] は以下のように定義されている。

```go
// A type, typically a collection, that satisfies sort.Interface can be
// sorted by the routines in this package. The methods require that the
// elements of the collection be enumerated by an integer index.
type Interface interface {
    // Len is the number of elements in the collection.
    Len() int
    // Less reports whether the element with
    // index i should sort before the element with index j.
    Less(i, j int) bool
    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)
}
```

## オブジェクトのソート

つまり，この [`sort`].`Interface` インタフェースを持つ型であれば [`sort`].`Sort()` 関数でソート可能ということになる。
たとえば以下のオブジェクト集合を考える。

```go
// A Planet defines the properties of a solar system object.
type Planet struct {
    Name     string
    Mass     float64
    Distance float64
}

var planets = []Planet{
    {"Mercury", 0.055, 0.4},
    {"Venus", 0.815, 0.7},
    {"Earth", 1.0, 1.0},
    {"Mars", 0.107, 1.5},
}
```

`Planet` オブジェクトの集合に対する [`sort`].`Interface` インタフェースはこんな感じにする[^str]。

[^str]: 今回は簡単のため [slice] を使っているが，データ集合は [slice] である必要はなく [`sort`].`Interface` インタフェースを持つ任意のオブジェクトであればよい。

```go
// ByMass implements sort.Interface for []Planet based on the Mass field.
type ByMass []Planet

func (a ByMass) Len() int           { return len(a) }
func (a ByMass) Less(i, j int) bool { return a[i].Mass < a[j].Mass }
func (a ByMass) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
```

つまり `Mass` フィールド値の昇順に並べるわけだ。

全体ではこんな感じになるだろう。

```go
package main

import (
    "fmt"
    "sort"
)

// A Planet defines the properties of a solar system object.
type Planet struct {
    Name     string
    Mass     float64
    Distance float64
}

func (p Planet) String() string {
    return p.Name
}

// ByMass implements sort.Interface for []Planet based on the Mass field.
type ByMass []Planet

func (a ByMass) Len() int           { return len(a) }
func (a ByMass) Less(i, j int) bool { return a[i].Mass < a[j].Mass }
func (a ByMass) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
    planets := []Planet{
        {"Mercury", 0.055, 0.4},
        {"Venus", 0.815, 0.7},
        {"Earth", 1.0, 1.0},
        {"Mars", 0.107, 1.5},
    }

    for _, p := range planets {
        fmt.Printf("%v ", p)
    }
    fmt.Print("\n")
    sort.Sort(ByMass(planets))
    for _, p := range planets {
        fmt.Printf("%v ", p)
    }
}
```

結果は以下の通りで意図通りの動作になっているのが分かるだろう。

```text
$ go run sort3.go
Mercury Venus Earth Mars
Mercury Mars Venus Earth
```

## `sort.Slice()` 関数を使う場合 

[slice] 限定であるが， [`sort`].`Slice()` 関数を使えば [`sort`].`Interface` インタフェースを定義しなくてもソートを行うことができる。
[`sort`].`Slice()` 関数の定義は以下の通り。

```go
func Slice(slice interface{}, less func(i, j int) bool)
```

実際のコードはこんな感じになる。

```go
package main

import (
    "fmt"
    "sort"
)

// A Planet defines the properties of a solar system object.
type Planet struct {
    Name     string
    Mass     float64
    Distance float64
}

func (p Planet) String() string {
    return p.Name
}

func main() {
    planets := []Planet{
        {"Mercury", 0.055, 0.4},
        {"Venus", 0.815, 0.7},
        {"Earth", 1.0, 1.0},
        {"Mars", 0.107, 1.5},
    }

    for _, p := range planets {
        fmt.Printf("%v ", p)
    }
    fmt.Print("\n")
    sort.Slice(planets, func(i, j int) bool {
        return planets[i].Mass < planets[j].Mass
    })
    for _, p := range planets {
        fmt.Printf("%v ", p)
    }
}
```

[`sort`].`Slice()` 関数の第2引数が関数閉包（closure）になっている点に注意[^c]。
これなら第2引数の関数の内容を変えれば任意の規則でソートを行うことができる。

[^c]: つか， [Go 言語]の関数は全て関数閉包として動作するんだけどね。

結果は [`sort`].`Interface` インタフェースがある場合と同じく

```text
$ go run sort4.go
Mercury Venus Earth Mars
Mercury Mars Venus Earth
```

となった。

さて，実際に [`sort`].`Slice()` 関数を覗いてみよう。

```go
// maxDepth returns a threshold at which quicksort should switch
// to heapsort. It returns 2*ceil(lg(n+1)).
func maxDepth(n int) int {
    var depth int
    for i := n; i > 0; i >>= 1 {
        depth++
    }
    return depth * 2
}

// lessSwap is a pair of Less and Swap function for use with the
// auto-generated func-optimized variant of sort.go in
// zfuncversion.go.
type lessSwap struct {
    Less func(i, j int) bool
    Swap func(i, j int)
}

// Slice sorts the provided slice given the provided less function.
//
// The sort is not guaranteed to be stable. For a stable sort, use
// SliceStable.
//
// The function panics if the provided interface is not a slice.
func Slice(slice interface{}, less func(i, j int) bool) {
    rv := reflect.ValueOf(slice)
    swap := reflect.Swapper(slice)
    length := rv.Len()
    quickSort_func(lessSwap{less, swap}, 0, length, maxDepth(length))
}
```
[`reflect`].`ValueOf()` 関数は [`reflect`].`Value` を取得する関数だ[^rf1]。
その次の [`reflect`].`Swapper()` 関数がポイント。
この関数は先程の  [`sort`].`Interface` インタフェースでいうところの `Swap()` 関数に相当するものを返す[^rf2]。
なのでこんなこともできる。

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    s := []int{1, 2, 3}
    fmt.Println(s) // [1 2 3]
    reflect.Swapper(s)(0, 2)
    fmt.Println(s) // [3 2 1]
}
```

残りの `Len()` 関数に相当するものは [`reflect`].`Value` で用意されているし， `Less()` 関数に相当するものは [`sort`].`Slice()` 関数の引数として与えられる。
これでソートに必要な3つの関数が揃うわけだ。

[^rf1]: [`reflect`] パッケージについての詳細は割愛する。簡単に言うと， [Go 言語]において [interface] 型のインスタンスは型情報と値への参照の2つを保持していて，これに対応するのが [`reflect`].`Type` と [`reflect`].`Value` である（参考： [research!rsc: Go Data Structures: Interfaces](https://research.swtch.com/interfaces)）。
[^rf2]: [`reflect`].`Swapper()` 関数は引数の型が [slice] であることを前提にしていて， [slice] でない場合は panic が返る。

ちなみに `quickSort_func()` 関数は，名前の通り，クイックソートである。
ただしクイックソートでは安定ソートにならないため，安定ソートを実行するための [`sort`].`SliceStable()` 関数も用意されている。
[`sort`].`SliceStable()` 関数ではアルゴリズムに挿入ソートを用いる[^s1]。

[^s1]: [`sort`].`Sort()` 関数も同じくクイックソートである。これの安定ソート版が [`sort`].`Stable()` 関数で同じく挿入ソートを用いている。

## ブックマーク

- [sliceのシャッフル - Qiita](http://qiita.com/sugyan/items/fd7138a756c1a409f5fd) : [Fisher–Yates shuffle](https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle) というアルゴリズムらしい
- [Go言語でバイトニックソート実装してみた - Qiita](http://qiita.com/nyamadandan/items/2c82011801b148c98e52)
- [Goでバケットソートアルゴリズム(ビット列を使用) - Qiita](http://qiita.com/ohkawa/items/269507985b3ae10cbff9)
- [Goでのアルゴリズムクイックリファレンス第2版(4.1.1 挿入ソート) - Qiita](http://qiita.com/amesho/items/64dcd231038c96345848)
- [interface{} をソートする - Qiita](http://qiita.com/tchssk/items/b61f1f06d22a6232d4c8)
- [Big Sky :: golang の sort インタフェース難しい問題が解決した](http://mattn.kaoriya.net/software/lang/go/20161004092237.htm)
- [Go 言語 1.8 がリリース]({{< ref "/remark/2017/02/go-1_8-released.md" >}}) : [`sort`].`Slice()` 関数はバージョン 1.8 で導入された
- [sort.Slice に学ぶ高速化のヒント - Qiita](https://qiita.com/chimatter/items/f908507287fe2c7030e9)
    - [sort.Sort と sort.Slice の速度比較 - Qiita](https://qiita.com/chimatter/items/bc8f3ab3e617211b9a24)
- [golang でクイックソートを並列化してみる - 長文書くところ](http://zenito9970.hatenablog.com/entry/2015/05/30/155726)
- [「ソート」を極める！ 〜 なぜソートを学ぶのか 〜 - Qiita](https://qiita.com/drken/items/44c60118ab3703f7727f) : ソート・アルゴリズムの概説。よくまとまっている

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`sort`]: https://golang.org/pkg/sort/ "sort - The Go Programming Language"
[`reflect`]: https://golang.org/pkg/reflect/ "reflect - The Go Programming Language"
[slice]: http://golang.org/ref/spec#Slice_types
[interface]: https://golang.org/doc/effective_go.html#interfaces_and_types "Effective Go - The Go Programming Language"

## 参考図書

<div class="hreview" ><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-20</dd><dd>評価<abbr class="rating" title="5"><img src="https://images-fe.ssl-images-amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4865940391/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4865940391.09._SCTHUMBZZZ_.jpg"  alt="Kotlinスタートブック -新しいAndroidプログラミング"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4839959234/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4839959234.09._SCTHUMBZZZ_.jpg"  alt="Docker実戦活用ガイド"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4274218961/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4274218961.09._SCTHUMBZZZ_.jpg"  alt="グッド・マス ギークのための数・論理・計算機科学"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-13">2016-07-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1FO/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/416jAxVU4NL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1FO/baldandersinf-22/">数学ガール／乱択アルゴリズム</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ株式会社 2014-02-14</dd><dd>評価<abbr class="rating" title="5"><img src="https://images-fe.ssl-images-amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1D6/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B00I8AT1D6.09._SCTHUMBZZZ_.jpg"  alt="数学ガール／ゲーデルの不完全性定理"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1CM/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B00I8AT1CM.09._SCTHUMBZZZ_.jpg"  alt="数学ガール／フェルマーの最終定理"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00L0PDMK4/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B00L0PDMK4.09._SCTHUMBZZZ_.jpg"  alt="数学ガール／ガロア理論"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00EYXMA9I/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B00EYXMA9I.09._SCTHUMBZZZ_.jpg"  alt="数学ガール"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00L0PDMIQ/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B00L0PDMIQ.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート／式とグラフ"  /></a> </p>
<p class="description" >工学ガール，リサちゃん登場！</p>
<p class="gtools" >reviewed by <a href="#maker" class="reviewer">Spiegel</a> on <abbr class="dtreviewed" title="2015-04-19">2015/04/19</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html">G-Tools</a>)</p>
</div>
