+++
title = "Project Euler で遊ぶ"
date = "2018-03-19T21:30:42+09:00"
update = "2018-03-26T15:18:48+09:00"
description = "Project Euler のポイントは「問題を解く」ことではなく「問題を理解する」ことにあると思う。 "
image = "/images/attention/kitten.jpg"
tags = [ "math", "games", "programming" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%81%84%E3%81%8B%E3%81%AB%E3%81%97%E3%81%A6%E5%95%8F%E9%A1%8C%E3%82%92%E3%81%A8%E3%81%8F%E3%81%8B-G-%E3%83%9D%E3%83%AA%E3%82%A2/dp/4621045938?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621045938"><img src="https://images-fe.ssl-images-amazon.com/images/I/51XGP8AFX2L._SL160_.jpg" width="112" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%81%84%E3%81%8B%E3%81%AB%E3%81%97%E3%81%A6%E5%95%8F%E9%A1%8C%E3%82%92%E3%81%A8%E3%81%8F%E3%81%8B-G-%E3%83%9D%E3%83%AA%E3%82%A2/dp/4621045938?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621045938">いかにして問題をとくか</a></dt>
	<dd>G. ポリア</dd>
	<dd>G. Polya (原著), 柿内 賢信 (翻訳)</dd>
    <dd>丸善 1975-04-01</dd>
    <dd>Book 単行本</dd>
    <dd>ASIN: 4621045938, EAN: 9784621045930</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">数学書。というか問いの立てかたやものの考え方についての指南書。のようなものかな。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
