+++
title = "π は間違ってる？"
date =  "2023-03-15T12:43:29+09:00"
description = "円周率の日だったので，ネタに走ってみた。"
image = "/images/attention/kitten.jpg"
tags = [ "math", "astronomy", "golang" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

昨日3月14日は「円周率の日」で「数学の日」でアインシュタイン博士の誕生日でした。
昔は物理学イベントもあった気がするけど ...まぁいいや。

## π は間違ってる？

Mastodon の TL を眺めてたら

{{< fig-quote class="nobox center" >}}
<iframe src="https://social.tinygo.org/@deadprogram/110020519869289793/embed" class="mastodon-embed" style="max-width: 100%; border: 0" width="400" allowfullscreen="allowfullscreen"></iframe>
{{< /fig-quote >}}

という投稿を見かける。
ここでいう $\tau$ (tau) は[^tao1]，半径 $r$，円周 $C$ の円

[^tao1]: 道（tao）じゃなくてギリシア文字の $\tau$ (tau) ね（笑）

{{< fig-img-quote src="./circle.png" title="The Tau Manifesto by Michael Hartl" link="https://tauday.com/tau-manifesto" lang="en" width="540" >}}

に対して

{{< fig-math >}}
\[
    \tau \equiv \dfrac{C}{r} = 6.283185307179586...
\]
{{< /fig-math >}}

と定義されたものだ。
“[The Tau Manifesto]” では，この $\tau$ を “the true circle constant” (真円定数？) と呼んでいるようだ。
言うまでもなくこれは

{{< fig-math >}}
\[
    \tau = 2\pi
\]
{{< /fig-math >}}

である。
$\pi$ を $\tau$ に置き換えることにより，たとえば角度について

{{< fig-img-quote src="./tau-angles.png" title="The Tau Manifesto by Michael Hartl" link="https://tauday.com/tau-manifesto" lang="en" >}}

などと（$\pi$ に比べて）ちょっとシンプルな表現にできる。
また，さきほどの定義を使って円周の長さは

{{< fig-math >}}
\[
    C = {\tau}{r}
\]
{{< /fig-math >}}

と記述でき，そこからの積分により円の面積 $S$ は

{{< fig-math >}}
\[
    S = \dfrac{1}{2}{\tau}r^2
\]
{{< /fig-math >}}

と導ける。
こっちのほうが，たとえば運動エネルギー $\tfrac{1}{2}mv^2$ や自由落下する物体の移動距離 $\tfrac{1}{2}gt^2$ などと同じ体裁で分かりやすいよね。

あるいは，有名なオイラーの公式

{{< fig-math >}}
\[
    e^{i\theta} = \cos{\theta} + i\sin{\theta}
\]
{{< /fig-math >}}

およびそこから導き出されるオイラーの等式は $\tau$ を使って

{{< fig-img-quote src="./tau_euler_circle.png" title="The Tau Manifesto by Michael Hartl" link="https://tauday.com/tau-manifesto" lang="en" width="521" >}}

つまり

{{< fig-math >}}
\[
    e^{i\tau} = 1
\]
{{< /fig-math >}}

となり

{{< fig-quote class="center" type="markdown" title="The Tau Manifesto by Michael Hartl" link="https://tauday.com/tau-manifesto" lang="en" >}}
A rotation by one turn is 1.
{{< /fig-quote >}}

と，より直感的な主張になる。
クールだろ！

というわけで，次は「$\tau$ の日」である6月28日にお祝いしましょう（笑）

## 【おまけ】 天文計算で円周率の精度は何桁まで？

同じく Mastodon の TL で

- [How Many Decimals of Pi Do We Really Need? - Edu News | NASA/JPL Edu](https://www.jpl.nasa.gov/edu/news/2016/3/16/how-many-decimals-of-pi-do-we-really-need/)

という記事を[教えてもらった](https://mastodon.social/@mondinspace/110022315123771073)。
これによると

{{< fig-quote type="markdown" title="How Many Decimals of Pi Do We Really Need?" link="https://www.jpl.nasa.gov/edu/news/2016/3/16/how-many-decimals-of-pi-do-we-really-need/" lang="en" >}}
For JPL's highest accuracy calculations, which are for interplanetary navigation, we use 3.141592653589793.
{{< /fig-quote >}}

なんだそうだ（2016年に書かれたものという但し書きがある）。
これっていわゆる double 型の浮動小数点数の有効桁数かな。
Go で書くと[こんな感じ](https://go.dev/play/p/6FqQkYp0hGq)になる。

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Printf("%g\n", math.Pi)
    fmt.Printf("%b\n", math.Pi)
    fmt.Printf("%x\n", math.Pi)
    //Output:
    //3.141592653589793
    //7074237752028440p-51
    //0x1.921fb54442d18p+01
}
```

うしろのふたつは浮動小数点数（IEEE 754）の内部表現で記述されている。
桁落ちなどの計算誤差に気をつける必要があるけど，基本型の浮動小数点数の演算で行けそうな感じではある。

なお [`math`]`.Pi` 定数は

```go
const (
    ...
    Pi  = 3.14159265358979323846264338327950288419716939937510582097494459 // https://oeis.org/A000796
    ...
)
```

と定義されている[^c1]。
残念ながら $\tau$ は定義されてなかった。
まぁ $2\pi$ でいけるからな（笑）

[^c1]: Go の定数の仕様については拙文「[リテラル定数](https://zenn.dev/spiegel/articles/20220904-literal-constants)」を参考にどうぞ。

## ブックマーク

- {{< pdf-file title="π Is Wrong" link="http://www.math.utah.edu/~palais/pi.pdf" >}}
- [Why Pi Matters | The New Yorker](https://www.newyorker.com/tech/annals-of-technology/pi-day-why-pi-matters)

[The Tau Manifesto]: https://tauday.com/tau-manifesto "Tau Day | No, really, pi is wrong: The Tau Manifesto by Michael Hartl"
[`math`]: https://pkg.go.dev/math "math package - math - Go Packages"

## 参考図書

{{% review-paapi "B00W6NCLJM" %}} <!-- 数学ガールの秘密ノート／丸い三角関数 -->
{{% review-paapi "B093PZLQMQ" %}} <!-- 数学ガールの物理ノート／ニュートン力学 -->
{{% review-paapi "4627275110" %}} <!-- 天体物理学 -->
