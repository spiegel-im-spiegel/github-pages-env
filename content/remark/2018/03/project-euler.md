+++
title = "Project Euler で遊ぶ"
date = "2018-03-19T21:30:42+09:00"
description = "Project Euler のポイントは「問題を解く」ことではなく「問題を理解する」ことにあると思う。 "
image = "/images/attention/kitten.jpg"
tags = [ "math", "games", "programming" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

私は社会人になってからコンピュータやプログラミングを習ったので，仕事以外でプログラミングをするとか，ましてや競技プログラミングなどには全く食指が動かなかったのだが，これは面白いと思った。

- [Project Euler]

なんというか数理パズルで遊んでる感じ。
Level や Award といったゲーム要素もあり自分のペースで楽しめるのが嬉しい。
[Go 言語]を触るようになって「プログラミングで遊ぶ」ことを楽しいと思えるようになったことも大きいかもしれない。
年を食ってから新しい趣味を見出すというのはいいものである。

アカウント登録が必要（問題を見るだけならアカウントなしでも大丈夫）。

問題（Problem）は既に600以上あるが好きなものから手を付けてよい。
解き方も自由だしどんな言語を使ってもよい。
つか，手作業で解いても全然OK。

答え（Answer）を入力して `[Check]` すれば正解かどうか教えてくれる。
正解した問題は unlock され，問題ごとのディスカッションに参加したり PDF による解説も見れたりする。

私のような英語不得手の方には日本語訳を公開している奇特な方もおられるので参考になるだろう。

- [Project Euler - PukiWiki](http://odz.sakura.ne.jp/projecteuler/index.php?Project%20Euler)

GitHub で解答集（Solutions）を公開しておられる人も多い。

- [Topic: project-euler](https://github.com/topics/project-euler)

というわけで私も参加してみようかなと。
方針としては

1. 問題を解くのは1日5問まで（ハマりすぎるとサルになる性格なので）
2. Go 言語によるプログラミング
3. 下手でも泥臭いコードでも，とにかく正解を書いてみる
4. 正解したら解説を読むなどしてコードを改善してみる

という感じだろうか。
ちなみに [Problem 1](https://projecteuler.net/problem=1 "Problem 1 - Project Euler") はこんな感じに解いてみた。

```go
package main

import "fmt"

/**
 * Multiples of 3 and 5
 *
 * https://projecteuler.net/problem=1
 * If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
 * Find the sum of all the multiples of 3 or 5 below 1000.
 *
 * http://odz.sakura.ne.jp/projecteuler/index.php?cmd=read&page=Problem%201
 * 10未満の自然数のうち, 3 もしくは 5 の倍数になっているものは 3, 5, 6, 9 の4つがあり, これらの合計は 23 になる.
 * 同じようにして, 1000 未満の 3 か 5 の倍数になっている数字の合計を求めよ.
 *
 * Contents of Project Euler are licenced under a Creative Commons Licence: Attribution-NonCommercial-ShareAlike 2.0 UK: England & Wales.
 * https://creativecommons.org/licenses/by-nc-sa/2.0/uk/
 */

func answer0(max int) int {
    sum := 0
    for n := 1; n < max; n++ {
        if (n%3) == 0 || (n%5) == 0 {
            sum += n
        }
    }
    return sum
}

func sumDivisibleBy(max, n int) int {
    m := (max - 1) / n
    return n * (m * (m + 1)) / 2
}
func answer1(max int) int {
    return sumDivisibleBy(max, 3) + sumDivisibleBy(max, 5) - sumDivisibleBy(max, 3*5)
}

func main() {
    fmt.Println(answer0(1000))
    fmt.Println(answer1(1000))
}
```

`answer0()` が最初に書いたコードで `answer1()` が改善後のコードである。
`answer0()` はなるべく問題の記述に忠実に書くようにして，改善コード `answer1()` の検算用に残している。

[Project Euler] のポイントは「問題を解く」ことではなく「問題を理解する」ことにあると思う。
これは数理パズルで遊ぶときも同様で[^sudoku1]，問題を解くことそのものよりも問題の背後にある「何か」を見つけたときの喜びのほうが大きい。

[^sudoku1]: たとえば世界的に有名な「数独（sudoku）」では，これをテーマにした論文とかあるらしい： [数独の科学 | 日経サイエンス](http://www.nikkei-science.com/page/magazine/0609/sudoku.html)

まぁ，仕事ならこんな悠長なことはやってられないが，遊びだからこそできることもあるのである。

## ブックマーク

- [テスト・フレームワークで Project Euler を解く]({{< ref "/golang/testing-for-project-euler.md" >}})

[Project Euler]: https://projecteuler.net/
[Go 言語]: https://golang.org/ "The Go Programming Language"

## 参考図書

{{% review-paapi "B079JLW5YN" %}} <!-- プログラマの数学 第2版 -->
{{% review-paapi "4621045938" %}} <!-- いかにして問題をとくか -->
