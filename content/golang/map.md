+++
title = "Map の話"
date =  "2018-02-07T13:36:46+09:00"
description = "description"
image = "/images/attention/go-code.png"
tags        = [ "golang" ]
draft = true

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = true
  mermaidjs = false
+++

今回も覚え書き用の小ネタ。

## [map] による連想配列の定義と基本操作

まず， [map] を使って以下の連想配列を考える。

```go
type platnets map[string]float64

p := platnets{
    "Mercury": 0.055,
    "Venus":   0.815,
    "Earth":   1.0,
}
```

この連想配列の各要素を操作する記述は以下のような感じになる。

```go
package main

import (
    "fmt"
)

type platnets map[string]float64

func main() {
    p := platnets{
        "Mercury": 0.055,
        "Venus":   0.815,
        "Earth":   1.0,
    }

    if v, ok := p["Mars"]; ok {
        fmt.Println("Mars =", v)
    } else {
        fmt.Println("Mars is not exist.")
    }

    p["Mars"] = 0.107 //add or set
    if v, ok := p["Mars"]; ok {
        fmt.Println("Mars =", v)
    } else {
        fmt.Println("Mars is not exist.")
    }

    delete(p, "Mars")
    if v, ok := p["Mars"]; ok {
        fmt.Println("Mars =", v)
    } else {
        fmt.Println("Mars is not exist.")
    }
}
```

このコードの実行結果は以下の通り。

```text
Mars is not exist.
Mars = 0.107
Mars is not exist.
```

この連想配列の要素を全て取り出す。
これは `for range` 構文を使ってこんな感じに記述する。

{{< highlight go "hl_lines=13-15" >}}
package main

import "fmt"

type platnets map[string]float64

func main() {
    p := platnets{
        "Mercury": 0.055,
        "Venus":   0.815,
        "Earth":   1.0,
    }
    for k, v := range p {
        fmt.Println(k, v)
    }
}
{{< /highlight >}}

結果は以下の通り。

```text
Mercury 0.055
Venus 0.815
Earth 1
```

ちなみに `for range` 構文で [map] を使う場合，取り出しの順番は不定になる[^rnd]。
例えば，アルファベット順に取り出したければ以下のようにキーを[ソート]({{< relref "golang/sort.md" >}} "ソートを使う")した配列を用意する。

[^rnd]: 言語仕様上定義されていないという意味ではなく，意図的に乱択されるらしい。

{{< highlight go "hl_lines=18-24" >}}
package main

import (
    "fmt"
    "sort"
    "strings"
)

type platnets map[string]float64

func main() {
    p := platnets{
        "Mercury": 0.055,
        "Venus":   0.815,
        "Earth":   1.0,
    }

    keys := []string{}
    for k, _ := range p {
        keys = append(keys, k)
    }
    sort.Slice(keys, func(i, j int) bool {
        return strings.Compare(keys[i], keys[j]) < 0
    })
    for _ , k := range keys {
        fmt.Println(k, p[k])
    }
}
{{< /highlight >}}

結果は以下の通り

```text
Earth 1
Mercury 0.055
Venus 0.815
```

## [map] は連想配列を参照するオブジェクトである

[map] は**連想配列への参照を属性として持つオブジェクト**である。
したがって [map] インスタンスを引数として渡した場合は見かけ上「参照渡し」として機能する。

{{< highlight go "hl_lines=13-15" >}}
package main

import "fmt"

type platnets map[string]float64

func (p platnets) print() {
    for k, v := range p {
        fmt.Println(k, v)
    }
}

func add(p platnets, k string, v float64) {
    p[k] = v
}

func main() {
    p := platnets{
        "Mercury": 0.055,
        "Venus":   0.815,
        "Earth":   1.0,
    }
    p.print()
    fmt.Println()

    add(p, "Mars", 0.107)
    p.print()
}
{{< /highlight >}}

このコードの実行結果は以下の通り。

```text
Mercury 0.055
Venus 0.815
Earth 1

Mercury 0.055
Venus 0.815
Earth 1
Mars 0.107
```

## 連想配列の複製

[map] 内の連想配列の複製を行う関数は用意されていない[^cpy] ため `for range` 構文で地道に処理する。

[^cpy]: 代入構文では [map] インスタンスの複製ができるだけで，参照している連想配列は同じものになる。

{{< highlight go "hl_lines=15-17" >}}
package main

import "fmt"

type platnets map[string]float64

func (p platnets) print() {
    for k, v := range p {
        fmt.Println(k, v)
    }
}

func (p platnets) copy() platnets {
    c := platnets{}
    for k, v := range p {
        c[k] = v
    }
    return c
}

func main() {
    p1 := platnets{
        "Mercury": 0.055,
        "Venus":   0.815,
        "Earth":   1.0,
    }

    p2 := p1.copy()
    p2["Mars"] = 0.107

    p1.print()
    fmt.Println()
    p2.print()
}
{{< /highlight >}}

このコードの実行結果は以下の通り。

```text
Mercury 0.055
Venus 0.815
Earth 1

Mercury 0.055
Venus 0.815
Earth 1
Mars 0.107
```

## 連想配列の比較

[map] インスタンス同士は `==` 演算子による比較ができない。
以下のコードはコンパイルエラーになる[^cmp1]。

[^cmp1]: `&p1 == &p2` ならコンパイルエラーにはならないが，やってることは [map] インスタンスのポインタ値を比較しているだけである。

{{< highlight go "hl_lines=23" >}}
package main

import "fmt"

type platnets map[string]float64

func (p platnets) copy() platnets {
    c := platnets{}
    for k, v := range p {
        c[k] = v
    }
    return c
}

func main() {
    p1 := platnets{
        "Mercury": 0.055,
        "Venus":   0.815,
        "Earth":   1.0,
    }

    p2 := p1.copy()
    if p1 == p2 {
        fmt.Println("p1 == p2")
    } else {
        fmt.Println("p1 != p2")
    }
}
{{< /highlight >}}

```text
prog.go:23:8: invalid operation: p1 == p2 (map can only be compared to nil)
```

[map] が参照している連想配列の内容を比較したいのであれば [`reflect`]`.DeepEqual()` 関数が使える。

{{< highlight go "hl_lines=26-30" >}}
package main

import (
    "fmt"
    "reflect"
)

type platnets map[string]float64

func (p platnets) copy() platnets {
    c := platnets{}
    for k, v := range p {
        c[k] = v
    }
    return c
}

func main() {
    p1 := platnets{
        "Mercury": 0.055,
        "Venus":   0.815,
        "Earth":   1.0,
    }

    p2 := p1.copy()
    if reflect.DeepEqual(p1, p2) {
        fmt.Println("p1 == p2")
    } else {
        fmt.Println("p1 != p2")
    }
}
{{< /highlight >}}

このコードの実行結果は以下の通り。

```text
p1 == p2
```

[map] が参照している連想配列のインスタンスが同一であるかどうか調べるには [`reflect`]`.ValueOf()` 関数で値（＝連想配列）を取得し，そのポインタ値を `==` 演算子で比較する。

{{< highlight go "hl_lines=26-30 33-37" >}}
package main

import (
    "fmt"
    "reflect"
)

type platnets map[string]float64

func (p platnets) copy() platnets {
    c := platnets{}
    for k, v := range p {
        c[k] = v
    }
    return c
}

func main() {
    p1 := platnets{
        "Mercury": 0.055,
        "Venus":   0.815,
        "Earth":   1.0,
    }

    p2 := p1
    if reflect.ValueOf(p1).Pointer() == reflect.ValueOf(p2).Pointer() {
        fmt.Println("p1 == p2")
    } else {
        fmt.Println("p1 != p2")
    }

    p2 = p1.copy()
    if reflect.ValueOf(p1).Pointer() == reflect.ValueOf(p2).Pointer() {
        fmt.Println("p1 == p2")
    } else {
        fmt.Println("p1 != p2")
    }
}
{{< /highlight >}}

このコードの実行結果は以下の通り。

```text
p1 == p2
p1 != p2
```

## ブックマーク

- [golangのequalityの評価について - podhmo's diary](http://pod.hatenablog.com/entry/2016/07/30/204357)
- [Goで違うmapであることをテストする - Qiita](https://qiita.com/karupanerura/items/03d6766fd8568c15fc90)

- [配列と Slice]({{< relref "golang/array-and-slice.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[map]: http://golang.org/ref/spec#Map_types
[`reflect`]: https://golang.org/pkg/reflect/ "reflect - The Go Programming Language"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-20</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4865940391/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4865940391.09._SCTHUMBZZZ_.jpg"  alt="Kotlinスタートブック -新しいAndroidプログラミング"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4839959234/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4839959234.09._SCTHUMBZZZ_.jpg"  alt="Docker実戦活用ガイド"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274218961/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274218961.09._SCTHUMBZZZ_.jpg"  alt="グッド・マス ギークのための数・論理・計算機科学"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-13">2016-07-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
