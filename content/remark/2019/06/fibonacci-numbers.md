+++
title = "スプレッド・シートでフィボナッチ数列を列挙する"
date =  "2019-06-23T15:57:34+09:00"
description = "というわけで（大きな有効桁数が要件となる）数値計算をスプレッドシートで行うのはオススメできない。"
image = "/images/attention/kitten.jpg"
tags = [ "math", "floating-point" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

あっ，ブログネタ見っけ（笑）

- [LibreOffice CalcとExcelの計算結果誤差をみなさんで検証していただいた結果まとめ - Togetter](https://togetter.com/li/1367733)

はっきり言って前半の議論以下の脊髄反射的やり取りは鼻糞ほどの参考にもならないので，最後に紹介されているリンクを読むことを強くオススメする。

- [Excelの演算誤差](https://oku.edu.mie-u.ac.jp/~okumura/software/excel/roundoff.html)

いや，これ読んでて大昔に特定用途のバンドパスフィルタのシミュレーションを Excel で書かされたことを思い出したよ。
トラウマになるほどではなかったが，あれはなかなか酷い仕事だった（笑）

## 【事前準備】 フィボナッチ数列の一般項

{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}，問題を整理しよう。

まずはフィボナッチ数列の定義から。
$n$ 番目のフィボナッチ数を $F_n$ で表すと

{{< div-box >}}
\begin{align*}
F_0 &= 0 \\
F_1 &= 1 \\
F_{n+2} &= F_n + F_{n+1} & (n = 0,1,2,3 \cdots) 
\end{align*}
{{< /div-box >}}

で定義できる。
ちなみに，この定義を[素朴に Go 言語コードにした]({{< ref "/golang/recursive-call-and-function-table.md" >}} "再帰呼び出しと関数テーブル")のが以下。

```go
package main

import "fmt"

var fibonacciNumbers = make(map[int64]int64)

func fibonacciNumber0(n int64) int64 {
    switch n {
    case 0:
        return 0
    case 1:
        return 1
    default:
        if fn, ok := fibonacciNumbers[n]; ok {
            return fn
        }
        fn := fibonacciNumber0(n-2) + fibonacciNumber0(n-1)
        fibonacciNumbers[n] = fn
        return fn
    }
}

func main() {
    fmt.Println("| $n$ | $F_n$ |")
    fmt.Println("| ---:| ---:|")
    for n := int64(1); n <= 75; n++ {
        fmt.Printf("| %d | %d |\n", n, fibonacciNumber0(n))
    }
}
```

このコードを実行して1番目から75番目までのフィボナッチ数列をつくりリファレンス値としよう。
一部を挙げておく。

| $n$ |            $F_n$ |
| ---:| ----------------:|
|   1 |                1 |
|   2 |                1 |
|   3 |                2 |
|   4 |                3 |
|   5 |                5 |
| ... |              ... |
|  69 |  117669030460994 |
|  70 |  190392490709135 |
|  71 |  308061521170129 |
|  72 |  498454011879264 |
|  73 |  806515533049393 |
|  74 | 1304969544928657 |
|  75 | 2111485077978050 |

ところで定義で挙げた数式は「漸化式」と呼ばれるものだが，一般項で表すこともできる。
面倒なので結果だけ [Wikipedia から引用](https://en.wikipedia.org/wiki/Fibonacci_number "Fibonacci number - Wikipedia")してしまおう。
すなわち

{{< div-box >}}
\begin{align*}
    \varphi &= \frac{1+\sqrt{5}}{2} \\
    \psi &= \frac{1-\sqrt{5}}{2} = 1 - \varphi
\end{align*}
{{< /div-box >}}

とするなら $F_n$ は

{{< div-box >}}
\begin{align*}
    F_n &= \frac{\varphi^n - \psi^n}{\varphi - \psi} = \frac{\varphi^n - \psi^n}{\sqrt{5}} & (n = 0,1,2,3 \cdots)
\end{align*}
{{< /div-box >}}

と書ける[^gr1]。
これを[Go 言語]で書くとこんな感じ。

[^gr1]: ちなみに今回出てきた $\varphi$ の値は「黄金比（golden ratio）」などと呼ばれているらしい。黄金比については与太話も含めて色々と面白い話もあるが今回は割愛する。

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("| $n$ | $\\varphi^n$ | $\\psi^n$ | $F_n$ |")
    fmt.Println("| ---:| ---:| ---:| ---:|")
    r5 := math.Sqrt(5)
    phi := (1 + math.Sqrt(5)) / 2
    psi := 1 - phi
    for n := int64(1); n <= 75; n++ {
        phin := math.Pow(phi, float64(n))
        psin := math.Pow(psi, float64(n))
        f := (phin - psin) / r5
        fmt.Printf("| %d | %.15f | %.15f | %.15f |\n", n, phin, psin, f)
    }
}
```

結果も載せておこう。

| $n$ |                      $\varphi^n$ |           $\psi^n$ |                            $F_n$ |
| ---:| --------------------------------:| ------------------:| --------------------------------:|
|   1 |                1.618033988749895 | -0.618033988749895 |                1.000000000000000 |
|   2 |                2.618033988749895 |  0.381966011250105 |                1.000000000000000 |
|   3 |                4.236067977499790 | -0.236067977499790 |                2.000000000000000 |
|   4 |                6.854101966249685 |  0.145898033750315 |                3.000000000000000 |
|   5 |               11.090169943749475 | -0.090169943749474 |                5.000000000000000 |
| ... |                              ... |                ... |                              ... |
|  69 |  263115950957275.968750000000000 | -0.000000000000004 |  117669030460993.984375000000000 |
|  70 |  425730551631122.937500000000000 |  0.000000000000002 |  190392490709134.968750000000000 |
|  71 |  688846502588399.000000000000000 | -0.000000000000001 |  308061521170129.000000000000000 |
|  72 | 1114577054219521.750000000000000 |  0.000000000000001 |  498454011879263.875000000000000 |
|  73 | 1803423556807920.750000000000000 | -0.000000000000001 |  806515533049392.875000000000000 |
|  74 | 2918000611027442.500000000000000 |  0.000000000000000 | 1304969544928656.750000000000000 |
|  75 | 4721424167835363.000000000000000 | -0.000000000000000 | 2111485077978049.500000000000000 |

$\psi^n$ の値が小さすぎてガッツリ情報落ちしているが $F_n$ として必要なのは整数部分のみなので小数点以下を丸めた値が正しければ無問題。
実際のところ

{{< div-box >}}
\begin{align*}
    \left| \frac{\psi^n}{\sqrt{5}} \right| &\lt \frac{1}{2} & (n = 0,1,2,3 \cdots)
\end{align*}
{{< /div-box >}}

と言えるので

{{< div-box >}}
\begin{align*}
    F_n &= \left[ \frac{\varphi^n}{\sqrt{5}} \right] = \left\lfloor \frac{\varphi^n}{\sqrt{5}} + \frac{1}{2} \right\rfloor & (n = 0,1,2,3 \cdots)
\end{align*}
{{< /div-box >}}

でも十分いけるのだ。

ていうかね。
$\sqrt{5}$ を無理やり浮動小数点数に展開して計算してるんだから誤差が出て当然なのであり，それが不可視化されている各スプレッドシートのほうが相当怪しげな内部処理をしていると言わざるを得ない。

## スプレッドシートでフィボナッチ数列の一般項を計算する

さて，ようやく本題。

つまるところ，この一般項を使ってスプレッドシートでフィボナッチ数列を計算させたらどうなるか，というのが今回のお題だ。

1. まずA列に 1, 2, 3, 4,... と整数値を入れておく
1. B列には式 `=((1+SQRT(5))/2)^A1` (1行目の場合) をセットする（$\varphi^n$ 相当）
1. C列には式 `=((1-SQRT(5))/2)^A1` (1行目の場合) をセットする（$\psi^n$ 相当）
1. D列には式 `=((((1+SQRT(5))/2)^A1)-(((1-SQRT(5))/2)^A1))/SQRT(5)` (1行目の場合) をセットする

これで B列 → C列 → D列 と計算過程を追うことができる。

まずは [LibreOffice] Calc で実行してみた。
ちなみにバージョン6.2系である。

- [fibonacci-numbers.ods](./fibonacci-numbers.ods)

結果の一部，70番目以降のみ挙げておく。

| $n$ |                      $\varphi^n$ |           $\psi^n$ |                            $F_n$ |
| ---:| --------------------------------:| ------------------:| --------------------------------:|
|  70 |  425730551631124.000000000000000 |  0.000000000000002 |  190392490709135.000000000000000 |
|  71 |  688846502588401.000000000000000 | -0.000000000000001 |  308061521170130.000000000000000 |
|  72 | 1114577054219520.000000000000000 |  0.000000000000001 |  498454011879265.000000000000000 |
|  73 | 1803423556807930.000000000000000 | -0.000000000000001 |  806515533049395.000000000000000 |
|  74 | 2918000611027450.000000000000000 |  0.000000000000000 | 1304969544928660.000000000000000 |
|  75 | 4721424167835376.000000000000000 |  0.000000000000000 | 2111485077978060.000000000000000 |

まず小数点以下の値が潰れているのが懸念点なのだが（これについては[後述]({{< relref "#me">}})する），表からは71番目からはリファレンス値に対して明らかにズレが生じていることが分かる。

同じことを Google スプレッドシートでもやってみた。

- [fibonacci-numbers](https://docs.google.com/spreadsheets/d/1feWsGdqbmmXbchLP4kofM3RL8FAdP2jimCttJLRZ3YA/edit?usp=sharing)

これも70番目以降のみ挙げておこう。

| $n$ |                      $\varphi^n$ |           $\psi^n$ |                            $F_n$ |
| ---:| --------------------------------:| ------------------:| --------------------------------:|
|  70 |  425730551631124.000000000000000 |  0.000000000000002 |  190392490709135.000000000000000 |
|  71 |  688846502588401.000000000000000 | -0.000000000000001 |  308061521170130.000000000000000 |
|  72 | 1114577054219520.000000000000000 |  0.000000000000001 |  498454011879265.000000000000000 |
|  73 | 1803423556807930.000000000000000 | -0.000000000000001 |  806515533049395.000000000000000 |
|  74 | 2918000611027450.000000000000000 |  0.000000000000000 | 1304969544928660.000000000000000 |
|  75 | 4721424167835380.000000000000000 |  0.000000000000000 | 2111485077978060.000000000000000 |

途中計算に若干の違いはあれど最終的な結果は同じになった。

Excel のことは知らない[^e1]。

[^e1]: もはや Microsoft Office は個人ユーザが使うものではないだろう。あれはお役所や企業等のレガシー環境で仕方なく使う道具である。

## Machine Epsilon{#me}

ある処理系における浮動小数点数の精度を調べるには $1+\epsilon-1$ がゼロにならない最小の $\epsilon$ つまり machine epsilon を調べるのがよい。
浮動小数点数の基数は2進数なので $\epsilon$ の値を

{{< div-box >}}
\begin{align*}
    \epsilon &= \frac{1}{2^{n}} & (n = 1,2,3 \cdots)
\end{align*}
{{< /div-box >}}

として調べていけばいいだろう。
たとえば [Go 言語]であれば

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("| $n$ | $\\epsilon$ | $1+\\epsilon-1$ |")
    fmt.Println("| ---:| ---:| ---:|")
    for n := int64(1); n <= 55; n++ {
        e := math.Pow(2.0, float64(-n))
        me := 1.0 + e - 1.0
        fmt.Printf("| %d | %v | %v |\n", n, e, me)
    }
}
```

といったコード。
このコードの実行結果（の一部）は以下の通り。

| $n$ |             $\epsilon$ |         $1+\epsilon-1$ |
| ---:| ----------------------:| ----------------------:|
|  45 |  2.842170943040401e-14 |  2.842170943040401e-14 |
|  46 | 1.4210854715202004e-14 | 1.4210854715202004e-14 |
|  47 |  7.105427357601002e-15 |  7.105427357601002e-15 |
|  48 |  3.552713678800501e-15 |  3.552713678800501e-15 |
|  49 | 1.7763568394002505e-15 | 1.7763568394002505e-15 |
|  50 |  8.881784197001252e-16 |  8.881784197001252e-16 |
|  51 |  4.440892098500626e-16 |  4.440892098500626e-16 |
|  52 |  2.220446049250313e-16 |  2.220446049250313e-16 |
|  53 | 1.1102230246251565e-16 |                      0 |

64ビット浮動小数点数の仮数部のサイズは52ビットなので，これは妥当な結果と言える。

同じことを [LibreOffice] Calc で実行した結果は以下の通り。

| $n$ |           $\epsilon$ |      $1+\epsilon-1$ |
| ---:| --------------------:| -------------------:|
|  45 |  2.8421709430404E-14 | 2.8421709430404E-14 |
|  46 |  1.4210854715202E-14 | 1.4210854715202E-14 |
|  47 |   7.105427357601E-15 |  7.105427357601E-15 |
|  48 |  3.5527136788005E-15 | 3.5527136788005E-15 |
|  49 | 1.77635683940025E-15 |                   0 |
|  50 | 8.88178419700125E-16 |                   0 |
|  51 | 4.44089209850063E-16 |                   0 |
|  52 | 2.22044604925031E-16 |                   0 |
|  53 | 1.11022302462516E-16 |                   0 |

Excel で聞かれる括弧の付け方で結果が変わる云々はなかったが $\epsilon=2^{-48}$ までの精度しかないようだ。
これなら先のフィボナッチ数列の一般項計算の結果も納得である。

Google スプレッドシートはちょっと複雑で，まず冪乗計算（`POW(x,y)` または `x^y`）が $2^{-33}$ までしか値が取れない。
しょうがないので，計算値ではなく，生値をセルにコピペして検証してみたところ $\epsilon=2^{-52}$ までの精度があることが確認できた。
しかし関数ごとに精度が異なるのであれば，これはこれでヒッジョーに面倒くさい話になる。

## というわけで

結論として（大きな有効桁数が要件となる）数値計算をスプレッドシートで行うのはオススメできない。
やるのであれば身元の確かなツールを使って数値計算を行い，その結果をスプレッドシートにインポートしてデータの整理を行うことであろうか。

## ブックマーク

- [1を1億回足して1億にならない場合]({{< ref "/golang/loop-counter.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[LibreOffice]: https://www.libreoffice.org/ "LibreOffice - Free Office Suite - Fun Project - Fantastic People"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%95%B0%E5%AD%A6%E3%82%AC%E3%83%BC%E3%83%AB-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B00EYXMA9I?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00EYXMA9I"><img src="https://images-fe.ssl-images-amazon.com/images/I/41ETMZ7i9qL._SL160_.jpg" width="114" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%95%B0%E5%AD%A6%E3%82%AC%E3%83%BC%E3%83%AB-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B00EYXMA9I?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00EYXMA9I">数学ガール</a></dt>
    <dd>結城 浩</dd>
    <dd>SBクリエイティブ 2007-06-26 (Release 2014-03-12)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00EYXMA9I</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">ミルカさんとの衝撃の encounter。数学ガールがワルツを踊る。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9E%E3%81%AE%E6%95%B0%E5%AD%A6-%E7%AC%AC2%E7%89%88-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B079JLW5YN?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B079JLW5YN"><img src="https://images-fe.ssl-images-amazon.com/images/I/51QDhrqqEtL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9E%E3%81%AE%E6%95%B0%E5%AD%A6-%E7%AC%AC2%E7%89%88-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B079JLW5YN?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B079JLW5YN">プログラマの数学 第2版</a></dt>
    <dd>結城 浩</dd>
    <dd>SBクリエイティブ 2018-01-16 (Release 2018-02-08)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B079JLW5YN</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">タイトル通りプログラマ必読書。第2版では機械学習に関する章が付録に追加された。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-03-19">2018-03-19</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
