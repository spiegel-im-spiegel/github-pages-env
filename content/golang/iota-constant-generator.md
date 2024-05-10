+++
title = "定数生成器 iota についてちゃんと書く"
date =  "2021-09-05T07:28:46+09:00"
description = "ポイントは iota はインデックス値であってカウンタではないということだ。"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

「[第16回『プログラミング言語Go』オンライン読書会](https://gpl-reading.connpass.com/event/221591/)」で出た話題を過去の失敗を含めて [Zenn で書いた][失敗談]が，微妙に反応がある気がするので，今後の自分のためにも `iota`[^iota1] についてちゃんと書いておいたほうがいいかもしれない。

[^iota1]: ちなみに iota は「イオタ」と読む。ギリシャ文字のイオタ（&Iota; / &iota;）を指す言葉である。

まずは用語から入るとしよう。

定数の中でも `rune`，整数（`int` など），浮動小数点数（`float64` など），複素数（`complex128` など）の定数をまとめて「数値定数（{{% span lang="en" %}}numeric constants{{% /span %}}）」と呼ぶ。
数値定数の定数式（{{% span lang="en" %}}constant expression{{% /span %}}）は以下のようにリテラルで表すことができる。

```go
package main

import "fmt"

const (
    Rune         = 'a'
    Integer      = 1 << 2
    FloatingPont = 3.4
    Complex      = 5.6 + 7.8i
)

func main() {
    fmt.Println(Rune, Integer, FloatingPont, Complex) // 97 4 3.4 (5.6+7.8i)
}
```

（型付けなし定数（{{% span lang="en" %}}untyped constant{{% /span %}}）として宣言されている点に注意）

定数生成器（{{% span lang="en" %}}constant generator{{% /span %}}）`iota` は数値定数の中で事前に宣言された型付けなし整数を表す識別子で，以下の性質を持つ。

{{< fig-quote type="markdown" title="The Go Programming Language Specification - The Go Programming Language" link="https://go.dev/ref/spec#Iota" lang="en" >}}
{{% quote %}}Within a [constant declaration](https://go.dev/ref/spec#Constant_declarations), the predeclared identifier `iota` represents successive untyped integer [constants](https://go.dev/ref/spec#Constants). Its value is the **index of the respective ConstSpec** in that constant declaration, starting at zero{{% /quote %}}.
{{< /fig-quote >}}

強調は私がやりました。
ポイントは `iota` は（`const (...)` 内の）インデックス値であってカウンタではないということだ。

私が過去にやらかした[勘違い][失敗談]について敢えて解説すると，もともと定数宣言は

{{< fig-quote type="markdown" title="The Go Programming Language Specification - The Go Programming Language" link="https://go.dev/ref/spec#Constant_declarations" lang="en" >}}
{{% quote %}}Within a parenthesized `const` declaration list the expression list may be omitted from any but the first ConstSpec. Such an empty list is equivalent to the textual substitution of the first preceding non-empty expression list and its type if any. Omitting the list of expressions is therefore equivalent to repeating the previous list{{% /quote %}}.
{{< /fig-quote >}}

という性質があり，同じ定数式の繰り返しを省略することができる。

```go
package main

import "fmt"

const (
    c1 = 1
    c2
    c3
)

func main() {
    fmt.Println(c1, c2, c3) // 1 1 1
}
```

これとインデックス値 `iota` を組み合わせて

```go
package main

import "fmt"

const (
    c1 = 1 + iota
    c2
    c3
)

func main() {
    fmt.Println(c1, c2, c3) // 1 2 3
}
```

のように書けるのだが，この字面から何となく `iota` をカウンタと思い込んでいたのだった。
嗚呼，恥ずかしい...

上の定数宣言は以下と等価である。

```go
const (
    c1 = 1 + iota // c1 == 1 (iota == 0)
    c2 = 1 + iota // c2 == 2 (iota == 1)
    c3 = 1 + iota // c3 == 3 (iota == 2)
)
```

また ConstSpec 単位のインデックス値なので，こんな変態的な記述もできる（笑）

{{< fig-quote class="nobox" type="markdown" title="The Go Programming Language Specification - The Go Programming Language" link="https://go.dev/ref/spec#Iota" lang="en" >}}
```go
const (
    bit0, mask0 = 1 << iota, 1<<iota - 1 // bit0 == 1, mask0 == 0  (iota == 0)
    bit1, mask1                          // bit1 == 2, mask1 == 1  (iota == 1)
    _, _                                 //                        (iota == 2, unused)
    bit3, mask3                          // bit3 == 8, mask3 == 7  (iota == 3)
)
```
{{< /fig-quote >}}

というわけで，どっとはらい。

[Go]: https://go.dev/
[失敗談]: https://zenn.dev/spiegel/articles/20210904-value-of-iota "iota 出現時の値はゼロとは限らない"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
