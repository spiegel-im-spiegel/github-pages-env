+++
title = "最大公約数と関数型プログラミング"
date =  "2017-09-23T23:32:56+09:00"
description = "そうだ。最大公約数（greatest common divisor）の話をしよう。"
tags = [ "golang", "function", "programming", "math", "greatest-common-divisor", "recursion" ]

[scripts]
  mathjax = true
  mermaidjs = false
+++

- [配列の全ての要素の最大公約数を求める (Java 8, Python) - Qiita](https://qiita.com/rsky/items/a39070208eaea38394c5)

この記事を見て思いついた。
そうだ。
最大公約数（greatest common divisor）の話をしよう。

## 最大公約数を求める

まずは定義から。
最大公約数の定義は以下の通り。

{{% fig-quote type="markdown" %}}
2つ以上の正の整数に共通な約数（公約数）のうち最大のもの
{{% /fig-quote %}}

折角なので何か例題を立ててみよう。

{{% fig-quote type="markdown" %}}
**例題1**

20 と 32 の最大公約数を求めよ。
{{% /fig-quote %}}

簡単な数だし，まずは暗算で解いてみる。
それぞれの値を素因数分解すると以下のようになる。

{{< fig-math >}}
\begin{align*}
20 &= 2^2 \times 5 \\
32 &= 2^5
\end{align*}
{{< /fig-math >}}

これにより最大公約数は $2^2 = 4$ だということが分かる。
簡単でよかったね。

## ユークリッドの互除法

さて，最大公約数を求める機械向けの計算方法としては「[ユークリッドの互除法]」が有名である。
具体的な手順は以下の通り。

1. 32 を 20 で割った余りは 12
2. 20 を 12 で割った余りは 8
3. 12 を 8 で割った余りは 4
4. 8 を 4 で割った余りは 0
5. したがって 32 と 20 の最大公約数は 4 である

これを図形で表すとこんな感じになる。
（以下の[図は Wikipedia のもの](https://ja.wikipedia.org/wiki/File:GCM_Of_20_And_32.gif "File_GCM Of 20 And 32.gif - Wikipedia")を拝借した。 [CC-BY-SA-3.0](https://creativecommons.org/licenses/by-sa/3.0/ "Creative Commons — Attribution-ShareAlike 3.0 Unported — CC BY-SA 3.0") で公開されている）

{{< fig-img src="https://upload.wikimedia.org/wikipedia/commons/thumb/1/1a/GCM_Of_20_And_32.gif/640px-GCM_Of_20_And_32.gif" title="From Wikipedia" link="https://ja.wikipedia.org/wiki/File:GCM_Of_20_And_32.gif" width="640" >}}


今回は[ユークリッドの互除法]の証明は割愛するとして[^pf1]，上記の手順の 1 から 4 は再帰処理になっていることが分かる。
というわけで，こんな感じのコードを組んでみる。
（だいぶ端折ったコードでゴメン）

[^pf1]: 結城浩さんが連載しておられる「[数学ガールの秘密ノート]」に[ユークリッドの互除法が出て来る](https://cakes.mu/posts/16292 "第195回　ユークリッドの互除法（前編）｜数学ガールの秘密ノート｜結城浩｜cakes（ケイクス）")。はやく本にならないかなぁ。

```go
package main

import "fmt"

func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a%b)
}

func main() {
    fmt.Println(gcd(32, 20)) // 4
}
```

これで実行結果は 4 になる。

実は [Go 言語]にも最大公約数を求める関数が標準パッケージ [`math/big`] に用意されている。
こんな感じで使える。

```go
package main

import (
    "fmt"
    "math/big"
)

func gcd(m, n uint64) uint64 {
    x := new(big.Int)
    y := new(big.Int)
    z := new(big.Int)
    a := new(big.Int).SetUint64(m)
    b := new(big.Int).SetUint64(n)
    return z.GCD(x, y, a, b).Uint64()
}

func main() {
    fmt.Println(gcd(20, 32)) // 4
}
```

[`big`]`.Int.GCD()` 関数も[ユークリッドの互除法]の一種で，正の整数 $a$, $b$ に対する最大公約数を $\mathrm{gcd}(a, b)$ とすると

{{< fig-gen >}}
\[
\mathrm{gcd}(a, b) = ax + by
\]
{{< /fig-gen >}}

となる $x$, $y$ の組み合わせを探すものだ。

## 3つ以上の数の最大公約数

では3つ以上の数の最大公約数を求めるにはどうすればいいか。
ちょっと考えれば分かるが，たとえば $a$, $b$, $c$ の最大公約数を求めたいなら $\mathrm{gcd}(\mathrm{gcd}(a, b), c)$ とすればいい。

では例題を立ててみよう。
これは「[配列の全ての要素の最大公約数を求める]」の設問と同等と言える。

{{% fig-quote type="markdown" %}}
**例題2**

(290021904, 927964716, 826824516, 817140688) の最大公約数を求めよ。
{{% /fig-quote %}}

まずはベタに for 文を回してベタに解いてみる。

```go
package main

import (
    "fmt"
    "math/big"
)

func gcd(m, n uint64) uint64 {
    x := new(big.Int)
    y := new(big.Int)
    z := new(big.Int)
    a := new(big.Int).SetUint64(m)
    b := new(big.Int).SetUint64(n)
    return z.GCD(x, y, a, b).Uint64()
}

func main() {
    values := []uint64{290021904, 927964716, 826824516, 817140688}

    z := values[0]
    for _, n := range values {
        z = gcd(n, z)
    }
    fmt.Println(z) // 92
}
```

ふむふむ。
イケてそうだな[^gcd1]。

[^gcd1]: 繰り返し処理が1回余分に回っているが $gcd(a, a) = a$ で影響はないのでご容赦。

## Go で関数型プログラミング

「[配列の全ての要素の最大公約数を求める]」で紹介されているコードは高階関数（higher-order function）である `reduce()` による関数型プログラミングになっている。

[Go 言語]の関数は第一級関数（first-class function）なので関数型っぽいプログラミングも可能なのだが[^fp1]， `reduce()` のような関数は標準では用意されていない。
ただし，似たような機能を持つパッケージを公開しておられる人はいる。
わざわざ自作するのもナニなので今回は以下のパッケージを利用させてもらおう。

[^fp1]: あくまでも「ぽい」である。たとえば [Go 言語]では if 文や for 文などは関数でも演算子でもない単なる制御[構文（stateement）]({{< relref "operators-and-statements.md" >}} "演算子とステートメント")であり，直接ロジックにコンパイルされる。これまでの手続き型言語と関数型言語の利点を合わせたような言語は「マルチパラダイム・プログラミング言語（multiparadigm programming language）」などと呼ばれたりする。 Python や Swift といった近頃流行りの言語もマルチパラダイムの流れのひとつである。

- [robpike/filter: Simple apply/filter/reduce package.](https://github.com/robpike/filter)

これを使って先程の例題を以下のコードで解くことができる。

```go
package main

import (
    "fmt"
    "math/big"

    "github.com/robpike/filter"
)

func gcd(m, n uint64) uint64 {
    x := new(big.Int)
    y := new(big.Int)
    z := new(big.Int)
    a := new(big.Int).SetUint64(m)
    b := new(big.Int).SetUint64(n)
    return z.GCD(x, y, a, b).Uint64()
}

func main() {
    values := []uint64{290021904, 927964716, 826824516, 817140688}

    fmt.Println(filter.Reduce(values, gcd, 1).(uint64)) // 92
}
```

もしくは `gcd()` 関数自体を引数に組み込んで

```go
package main

import (
    "fmt"
    "math/big"

    "github.com/robpike/filter"
)

func main() {
    values := []uint64{290021904, 927964716, 826824516, 817140688}

    fmt.Println(filter.Reduce(values,
        func(m, n uint64) uint64 {
            x := new(big.Int)
            y := new(big.Int)
            z := new(big.Int)
            a := new(big.Int).SetUint64(m)
            b := new(big.Int).SetUint64(n)
            return z.GCD(x, y, a, b).Uint64()
        },
        1).(uint64)) // 92
}
```

としてもよい。
ほら，これなら「実質的にワンライナー」と呼べないこともない（笑）

[`robpike/filter`] パッケージの作者も書いておられるが

{{< fig-quote title="robpike/filter" link="https://github.com/robpike/filter" lang="en" >}}
<q>I wanted to see how hard it was to implement this sort of thing in Go, with as nice an API as I could manage. It wasn't hard.<br>
<br>
Having written it a couple of years ago, I haven't had occasion to use it once. Instead, I just use "for" loops.<br>
<br>
You shouldn't use it either.</q>
{{< /fig-quote >}}

[`filter`]`.Reduce()` 関数を駆動するコストを考えれば[^rf1] 普通に for 文で回したほうが安上がりだよね。
イマドキっぽく関数型言語の利点をいくつか取り込んでいるとはいえ Haskell のようなガッツリした関数型言語とは役割が異なるので，無理に関数型にこだわらなくてもいいということである。

[^rf1]: [Go 言語]には総称型（Generics）がないため [`filter`]`.Reduce()` 関数内部で [`reflect`] パッケージを駆使することになるが，その分はどうしてもパフォーマンスに影響を与えてしまう（参考： [きみは Generics がとくいなフレンズなんだね，または「制約は構造を生む」]({{< ref "/remark/2017/03/generics-vs-duck-typing.md" >}})）。しかし [`reflect`] パッケージが悪というわけではなく，たとえば[先日紹介]({{< relref "sort.md" >}} "ソートを使う")した [`sort`]`.Slice()` 関数は [`reflect`] パッケージを効果的に使用した例といえる。高階関数は見た目はクールだが，それに惑わされることなく目的に適うコードを選ぶことが重要である。とはいえ，勉強のために [Go 言語]で高階関数を組んでみようというのは悪くないと思う。

## ブックマーク

- [3つ以上の数の最大公約数と最小公倍数 - Qiita](https://qiita.com/tawatawa/items/408b872a7092be0d7b3c)
- [Goで関数型プログラミング - Qiita](https://qiita.com/taksatou@github/items/d721a62158f554b8e399)
- [Go言語では高階関数よりforループを使え(by Rob Pike氏) - Qiita](https://qiita.com/yohhoy/items/d3c12361bb5eed3cbede)
- [関数型言語のウソとホント - Qiita](https://qiita.com/hiruberuto/items/26a813ab2b188ca39019)

- [再帰呼び出しと関数テーブル]({{< relref "recursive-call-and-function-table.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[配列の全ての要素の最大公約数を求める]: https://qiita.com/rsky/items/a39070208eaea38394c5
[ユークリッドの互除法]: https://ja.wikipedia.org/wiki/%E3%83%A6%E3%83%BC%E3%82%AF%E3%83%AA%E3%83%83%E3%83%89%E3%81%AE%E4%BA%92%E9%99%A4%E6%B3%95 "ユークリッドの互除法 - Wikipedia"
[数学ガールの秘密ノート]: https://cakes.mu/series/339 "数学ガールの秘密ノート｜結城浩｜cakes（ケイクス）"
[`big`]: https://golang.org/pkg/math/big/ "big - The Go Programming Language"
[`math/big`]: https://golang.org/pkg/math/big/ "big - The Go Programming Language"
[`reflect`]: https://golang.org/pkg/reflect/ "reflect - The Go Programming Language"
[`sort`]: https://golang.org/pkg/sort/ "sort - The Go Programming Language"
[`robpike/filter`]: https://github.com/robpike/filter "robpike/filter: Simple apply/filter/reduce package."
[`filter`]: https://github.com/robpike/filter "robpike/filter: Simple apply/filter/reduce package."

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B07VPSXF6N" %}} <!-- 改訂2版 みんなのGo言語 -->
{{% review-paapi "B00L0PDMJ0" %}} <!-- 数学ガールの秘密ノート／整数で遊ぼう -->
{{% review-paapi "B00I8AT1CM" %}} <!-- 数学ガール／フェルマーの最終定理 -->
