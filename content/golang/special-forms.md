+++
title = "特殊形式による式評価について"
date =  "2021-03-04T19:27:13+09:00"
description = "Go の言語仕様を読みましょう（笑）"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "map", "type", "channel" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[元ネタが Zenn の記事](https://zenn.dev/dqneo/articles/436bb59d565be7 "Go言語における式の評価文脈を理解する")だし小ネタだし，どちらで書こうか悩んだが，今まで書いたことがない切り口な気がするので，こっち側で。

いきなりだが，以下のコードを起点に話を始めよう。

```go {hl_lines=[10]}
package main

import "fmt"

func main() {
    ageMap := map[string]uint{
        "Alice": 18,
        "Teles": 20,
    }
    fmt.Println(ageMap["Alice"])
    // Output
    // 18
}
```

このコードをみて分かるように，インデックス式 `ageMap[x]` の評価結果は `uint` 型になっている（インデックス値が存在しない場合はゼロ値がセットされる）。
まぁ，当たり前だよね。

ところが `ageMap[x]` の評価結果を `(uint, bool)` の組（tuple）で受けることもできるのだ[^tuple1]。
これによってコードを

[^tuple1]: 厳密には，特殊形式 `(T, bool)` の2要素目は型付けなしの真偽値（untyped boolean）に評価される。「[型付けなし（untyped）](https://zenn.dev/hsaki/articles/gospecdictionary#untyped "Goの言語仕様書精読のススメ & 英語彙集")」というのは [Go] 特有の概念だそうで，具体的な型が決定される前の状態を指す。たとえば，定数 [`math`]`.Pi` は10進数64桁の小数点数で定義されている。型付けなし定数については『[プログラミング言語Go](https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』の3.6.2章にも解説がある。

```go {hl_lines=["10-14"]}
package main

import "fmt"

func main() {
    ageMap := map[string]uint{
        "Alice": 18,
        "Teles": 20,
    }
    if age, ok := ageMap["Selene"]; ok {
        fmt.Println(age)
    } else {
        fmt.Println("unknown")
    }
    // Output
    // unknown
}
```

と書き換えることでき， `ageMap[x]` の評価に失敗した際の挙動を記述できる。
[Go] の[言語仕様]では，これを「特殊形式（special form）」と呼んでいる（日本語が怪しいのはご容赦）。

[言語仕様]を眺めてみると，この特殊形式が適用可能なのは以下の3つのみのようだ。

|    # | Expressions             | Normal Form      | Special Form         |
| ---: | ----------------------- | ---------------- | -------------------- |
|    1 | Index Expression on Map | `foo := bar[x]`  | `foo, ok := bar[x]`  |
|    2 | Type Assertion          | `foo := bar.(T)` | `foo, ok := bar.(T)` |
|    3 | Receive Expression      | `foo := <-ch`    | `foo, ok := <-ch`    |

最初のはインデックス値 `x` が map にない場合に `false` になる。
2番目は型 `T` でのアサーションに失敗した場合に（panic を吐くことなく） `false` になる。
最後のはチャネル `ch` が閉じている場合に `false` になる。

他の式評価や関数の返り値ではこのようなことはできない。
たとえば slice のインデックス式評価に特殊形式はない。
インデックス値が範囲外なら単に panic が投げられるだけだ。

実は，特殊形式が上の3つの式評価でしか適用されないというのが分からなくて「どうして普通の関数の返り値とかではできないんだろう」とひとしきり悩んだことがあったのだった。
[言語仕様]を読めっての（笑）

## ブックマーク

- [The Go Programming Language Specification - The Go Programming Language](https://golang.org/ref/spec)
- [Goの言語仕様書精読のススメ & 英語彙集](https://zenn.dev/hsaki/articles/gospecdictionary)
- [Go言語における式の評価文脈を理解する](https://zenn.dev/dqneo/articles/436bb59d565be7)

[Go]: https://golang.org/ "The Go Programming Language"
[言語仕様]: https://golang.org/ref/spec "The Go Programming Language Specification - The Go Programming Language"
[`math`]: https://golang.org/pkg/math/ "math - The Go Programming Language"

## 参考図書

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
