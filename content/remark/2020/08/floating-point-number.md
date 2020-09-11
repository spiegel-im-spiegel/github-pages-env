+++
title = "0.1 たす 0.2 が 0.3 にならない場合"
date =  "2020-08-19T11:43:53+09:00"
description = "浮動小数点数の仕様はプログラマなら知ってて当然だよね（笑）"
image = "/images/attention/kitten.jpg"
tags = [ "math", "floating-point", "golang", "engineering" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

面白い tweet を[見かけた](https://twitter.com/prithvirathor99/status/1295728692316917761)ので。

以下のコード

```go
package main

import "fmt"

func main() {
    a := 0.1
    b := 0.2
    c := a + b
    fmt.Println(c == 0.3)
}
```

を実行したら `true` になるか `false` になるか，というもの。
分からなければ実際に書いて実行してみれば分かる。

ちなみに [Go 言語][Go]では小数点数のリテラル表現を使って `a := 0.1` の様に記述すると暗黙的に `float64` (倍精度浮動小数点数) 型として宣言・初期化される[^var]。

[^var]: 厳密には定数 “`0.1`” は，いったん「型付けなし」の浮動小数点数として評価された後，変数宣言時に `float64` に暗黙的に変換される。 

まぁ，浮動小数点数の仕様はプログラマなら知ってて当然だよね（笑）

念のために説明すると，これは（言語に関係なく）浮動小数点数の標準である IEEE 754 の仕様によるもの。
IEEE 754 は浮動小数点数を2進数で保持している。

たとえば `0.1` を [Go] の書式 `%b` で表すと

```go
package main

import "fmt"

func main() {
	fmt.Printf("%b\n", 0.1)
    //Output:
    //7205759403792794p-56
}
```

となる。
仮数部が10進数になっているので分かりにくいが，無理やり2進数で表すと

```go {hl_lines=[7,10]}
package main

import "fmt"

func main() {
	fmt.Printf("%b\n", 0.1)
    fmt.Printf("%bp-56\n", 7205759403792794)
    //Output:
    //7205759403792794p-56
    //11001100110011001100110011001100110011001100110011010p-56

}
```

と循環し，丸め誤差が発生していることが分かる。
これが「0.1 たす 0.2 が 0.3 にならない」カラクリである。

といっても数値計算であれば（計算誤差を含めた誤差評価ができているなら）特に問題はないのだが，たとえばお金の計算をする場合は小数点以下で計算誤差が発生しては困ってしまう。

そこで，大抵の言語では「通貨型」と呼ばれる型またはクラスが用意されている。
中身は10の冪でスケーリングした固定小数点数であることが多い。
こういった型やクラスを上手く使い分けることがプログラム設計の最初の1フィートとなるだろう。

なお [Go 言語][Go]では，通貨型あるいは固定小数点数型に相当するものが基本型にも標準ライブラリにも存在しない[^dec1]。
その代わりと言ってはナニだが，標準の [`math/big`][`big`] パッケージに有理数型 [`big`]`.Rat` がある。
これを使えば（ちょっと面倒くさい記述になるが）先ほどのコードは

[^dec1]: ただし [shopspring/decimal](https://github.com/shopspring/decimal "shopspring/decimal: Arbitrary-precision fixed-point decimal numbers in go") のようなパッケージを公開されている方はいる。感謝！

```go
package main

import (
    "fmt"
    "math/big"
)

func main() {
    a := big.NewRat(1, 10)
    b := big.NewRat(2, 10)
    c := (&big.Rat{}).Add(a, b)
    fmt.Println(big.NewRat(3, 10).Cmp(c) == 0)
}
```

のように書き換えることができる。
実行結果は各自でどうぞ（笑）

## ブックマーク

- [1を1億回足して1億にならない場合]({{< ref "/golang/loop-counter.md" >}})
- [Go 言語で浮動小数点数をいろいろな書式で表示してみる]({{< ref "/golang/floating-point-number.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[`big`]: https://pkg.go.dev/math/big "big package · pkg.go.dev"

## 参考図書

{{% review-paapi "B079JLW5YN" %}} <!-- プログラマの数学 第2版 -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
