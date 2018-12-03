+++
title = "Project Euler で遊ぶ"
date = "2018-03-19T21:30:42+09:00"
update = "2018-03-26T15:18:48+09:00"
description = "Project Euler のポイントは「問題を解く」ことではなく「問題を理解する」ことにあると思う。 "
image = "/images/attention/kitten.jpg"
tags = [ "math", "games", "programming" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

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
 * http://creativecommons.org/licenses/by-nc-sa/2.0/uk/
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

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B079JLW5YN/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/51QDhrqqEtL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B079JLW5YN/baldandersinf-22/">プログラマの数学 第2版</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2018-01-16</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B015643CPE.09._SCTHUMBZZZ_.jpg"  alt="暗号技術入門 第3版　秘密の国のアリス"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00L0PDMK4/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00L0PDMK4.09._SCTHUMBZZZ_.jpg"  alt="数学ガール／ガロア理論"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1CM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00I8AT1CM.09._SCTHUMBZZZ_.jpg"  alt="数学ガール／フェルマーの最終定理"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1D6/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00I8AT1D6.09._SCTHUMBZZZ_.jpg"  alt="数学ガール／ゲーデルの不完全性定理"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1FO/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00I8AT1FO.09._SCTHUMBZZZ_.jpg"  alt="数学ガール／乱択アルゴリズム"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート／数列の広場"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B018VE46YW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B018VE46YW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート／ベクトルの真実"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B073F45B97/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B073F45B97.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート／積分を見つめて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート／微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00L0PDMJ0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00L0PDMJ0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート／整数で遊ぼう"  /></a> </p>
<p class="description">タイトル通りプログラマ必読書。第2版では機械学習に関する章が付録に追加された。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-03-19">2018-03-19</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621045938/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51XGP8AFX2L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621045938/baldandersinf-22/">いかにして問題をとくか</a></dt><dd>G. ポリア 柿内 賢信 </dd><dd>丸善 1975-04-01</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4621085298/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4621085298.09._SCTHUMBZZZ_.jpg"  alt="いかにして問題をとくか・実践活用編"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4061497863/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4061497863.09._SCTHUMBZZZ_.jpg"  alt="数学的思考法―説明力を鍛えるヒント  講談社現代新書"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/462108819X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/462108819X.09._SCTHUMBZZZ_.jpg"  alt="数学×思考=ざっくりと  いかにして問題をとくか"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797375698/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797375698.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4185086180/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4185086180.09._SCTHUMBZZZ_.jpg"  alt="授業研究に学ぶ高校新数学科の在り方"  /></a> </p>
<p class="description" >数学書。というか問いの立てかたやものの考え方についての指南書。のようなものかな。</p>
<p class="gtools" >reviewed by <a href="#maker" class="reviewer">Spiegel</a> on <abbr class="dtreviewed" title="2014-09-26">2014/09/26</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html">G-Tools</a>)</p>
</div>
