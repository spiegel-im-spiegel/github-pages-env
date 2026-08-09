+++
title = "算数のお時間"
date =  "2026-08-09T15:32:35+09:00"
description = "合同式，数学的帰納法，オイラーの定理"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "math" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
  jsx = false
+++

今日は珍しく Facebook の TL から。

たぶんお子さんの夏休みの宿題かな。
対象は中学2年生らしい。
問題はこんな感じ。

{{< fig-quote type="markdown" >}}
3の累乗の計算の結果は，

{{< fig-math >}}
\[
    3^1 = 3,\ 3^2 = 9,\ 3^3 = 27,\ 3^4 = 81,\ 3^5 = 243,\ 3^6 = 729,\ \cdots
\]
{{< /fig-math >}}

$3^{100}$ を計算した数の一の位の数を答えなさい。
{{< /fig-quote >}}

例を眺めて1の位が 3, 9, 7, **1**, 3, 9, ... と4つの値で循環してることに気づけば， $3^{100} = 3^{(4\cdot25)}$ なので，1の位が1になることは直感的に分かる。
中学生の回答としてはこれでOK？ でも，もう少し数学っぽく記述してみたいよね。

## 合同式を使って書く

計算結果の1の位に注目するということは，10で割った余りを求めることと同じである。
つまり問題は

{{< fig-math >}}
\[
    3^{n} \bmod{10}
\]
{{< /fig-math >}}

を解くことと同じである。
3の100乗は

{{< fig-math >}}
\[
    3^{100} = 3^{(4\cdot25)} = (3^4)^{25}
\]
{{< /fig-math >}}

であり，このうち $3^4$ に注目すると，問題文より

{{< fig-math >}}
\[
    3^4 = 81 \equiv 1 \pmod{10}
\]
{{< /fig-math >}}

なので

{{< fig-math >}}
\[
    3^{100} = (3^4)^{25} \equiv 1^{25} \equiv 1 \pmod{10}
\]
{{< /fig-math >}}

と書ける。

## 数学的帰納法を使って書く

この問題を更に一般化してみよう。
つまり $0$ 以上の任意の整数 $k$ に対して

| $n$ | $3^n \bmod {10}$ |
|---|:---:|
| $4k$ | $1$ |
| $4k+1$ | $3$ |
| $4k+2$ | $9$ |
| $4k+3$ | $7$ |

が成り立つことを示せればいい。

まず $k=0$ のときは問題文より

| $n$ | $3^n \bmod {10}$ |
|---|:---:|
| $0$ | $1$ |
| $1$ | $3$ |
| $2$ | $9$ |
| $3$ | $7$ |

となることは明らか。

ある $k$ で上が成立すると仮定した場合， $k+1$ のときも成立することを示す。
まず $3^{4(k+1)}$ について

{{< fig-math >}}
\[
    3^{4(k+1)} = 3^{4k+4} = 3^{4k} \cdot 3^4 \equiv 1 \cdot 1 \equiv 1 \pmod{10}
\]
{{< /fig-math >}}

となり $3^{4(k+1)} \equiv 1 \pmod{10}$ も成り立つことが分かる。
続けて

{{< fig-math >}}
\[ \begin{align*}
    3^{4(k+1)+1} &= 3^{4(k+1)} \cdot 3 \equiv 1 \cdot 3 \equiv 3 \pmod{10} \\
    3^{4(k+1)+2} &= 3^{4(k+1)+1} \cdot 3 \equiv 3 \cdot 3 \equiv 9 \pmod{10} \\
    3^{4(k+1)+3} &= 3^{4(k+1)+2} \cdot 3 \equiv 9 \cdot 3 \equiv 7 \pmod{10}
\end{align*} \]
{{< /fig-math >}}

と残りのケースも成り立つことに分かる（直前の合同式の結果を用いて計算している点に注目）。
これにより $0$ 以上の任意の整数 $k$ に対して

| $n$ | $3^n \bmod {10}$ |
|---|:---:|
| $4k$ | $1$ |
| $4k+1$ | $3$ |
| $4k+2$ | $9$ |
| $4k+3$ | $7$ |

が成り立つと言える。
以上より $3^{100} = 3^{4\cdot25} \equiv 1 \pmod{10}$ と計算できる。

## オイラーの定理を使って書く

他に面白い解き方がないかなぁ，と [Kagi Assistant] に訊いてみたら，[オイラーの定理][Euler's theorem]を使った解法を教えてもらった。

オイラーの定理とは，互いに素（最大公約数が $1$ のみ）な整数 $a$ と $n$ に対して

{{< fig-math >}}
\[
    a^{\varphi(n)} \equiv 1 \pmod{n}
\]
{{< /fig-math >}}

が成り立つというもの[^e1]。
ここで $\varphi(n)$ は[オイラーの totient 関数][Euler's totient function]と呼ばれるもので，$1$ から $n$ までの整数のうち $n$ と互いに素なものの個数を表す。

[^e1]: [オイラーの定理][Euler's theorem]の証明については [Wikipedia][Euler's theorem] 等を参考にどうぞ。

$3$ と $10$ は互いに素で，かつ $\varphi(10) = \varphi(2)\varphi(5) = 4$ なので，オイラーの定理より

{{< fig-math >}}
\[
    3^{\varphi(10)} = 3^4 \equiv 1 \pmod{10}
\]
{{< /fig-math >}}

が保証されているということらしい。
あとは前節までの記述に従って

{{< fig-math >}}
\[
    3^{100} = 3^{(4\cdot25)} \equiv 1 \pmod{10}
\]
{{< /fig-math >}}

と書ける。

[オイラーの定理][Euler's theorem]は思いつかんかったわ。

## 中学生はどうやって書く？

[オイラーの定理][Euler's theorem]は中学生は習ってないよね，たぶん。
合同式や数学的帰納法は習ってるっけ？ まぁ，いいか。
習ってないなら，これを夏休みの自由研究にしてしまえ（笑）

[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"
[Euler's theorem]: https://en.wikipedia.org/wiki/Euler%27s_theorem "Euler's theorem - Wikipedia"
[Euler's totient function]: https://en.wikipedia.org/wiki/Euler%27s_totient_function "Euler's totient function - Wikipedia"

## 参考

{{% review-paapi "B00L0PDMJ0" %}} <!-- 数学ガールの秘密ノート／整数で遊ぼう -->
{{% review-paapi "B00I8AT1CM" %}} <!-- 数学ガール／フェルマーの最終定理 -->
