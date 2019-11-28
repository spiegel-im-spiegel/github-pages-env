+++
title = "Map の話"
date = "2018-02-07T21:03:54+09:00"
description = "連想配列： map の複製とか比較とか。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "programming", "map" ]

[scripts]
  mathjax = false
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
例えば，アルファベット順に取り出したければ以下のようにキーを[ソート]({{< relref "sort.md" >}} "ソートを使う")した配列を用意する。

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
したがって [map] インスタンスを引数として渡した場合は[見かけ上「参照渡し」として機能する]({{< relref "./map-as-a-associative-array.md" >}} "Map は連想配列ではなく連想配列への「参照」である")。

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

### Nil に注意

当たり前の話ではあるが [map] では空の連想配列と nil では挙動が異なる。
たとえば nil [map] に要素を追加しようとすると

```go
package main

type platnets map[string]float64

func main() {
	var p platnets
	p["Mars"] = 0.107 //add or set
}
```

以下のように実行時 panic になる。

```text
panic: assignment to entry in nil map
```

したがって有効な [map] インスタンスを宣言する場合は `make()` 関数で初期化するか nil 以外の初期値を与える必要がある。

```go
package main

type platnets map[string]float64

func main() {
	p := platnets{}
	p["Mars"] = 0.107 //add or set
}
```

ちなみに nil [map] に対して `len()` 関数を使っても実行時 panic にならない[^len1]。

[^len1]: これは nil [slice] も同様で， `len()` 関数を使っても実行時 panic にならず 0 が返る。

```go
package main

import "fmt"

type platnets map[string]float64

func main() {
	var p platnets
	fmt.Println(len(p)) // 0
}
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

[map] が参照している連想配列の内容を比較したいのであれば [`reflect`].`DeepEqual()` 関数が使える。

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

[map] が参照している連想配列のインスタンスが同一であるかどうか調べるには [`reflect`].`ValueOf()` 関数で値（＝連想配列）を取得し，そのポインタ値を `==` 演算子で比較する。

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
- [Goでmapをjson文字列に変換する - Qiita](https://qiita.com/konchan_____/items/dce130f79c49e04e9931)

- [配列と Slice]({{< relref "array-and-slice.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[map]: http://golang.org/ref/spec#Map_types
[slice]: http://golang.org/ref/spec#Slice_types
[`reflect`]: https://golang.org/pkg/reflect/ "reflect - The Go Programming Language"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN), 4621300253 (ISBN), 9784621300251 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
