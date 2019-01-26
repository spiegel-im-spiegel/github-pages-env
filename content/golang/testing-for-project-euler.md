+++
title = "テスト・フレームワークで Project Euler を解く"
date = "2018-03-26T19:27:50+09:00"
description = "ベンチマーク・テストもできるよ。"
image = "/images/attention/go-code2.png"
tags = ["golang", "testing", "benchmark"]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
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

別のセクションで「[Project Euler で遊ぶ]({{< ref "/remark/2018/03/project-euler.md" >}})」話を書いたのだが，この記事のコメントで「ユニットテストを書く練習になる」というのがあって「なるほど！」と膝を打った。
つか，自力で気づけよ，自分。
まだまだ TDD が身につかないなぁ。

ちうわけで，以下の問題（Problem）をテスト・フレームワークで解いてみる。

{{< fig-quote title="Problem 1 - Project Euler" link="https://projecteuler.net/problem=1" lang="en" >}}
<q>If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.<br>
Find the sum of all the multiples of 3 or 5 below 1000.</q>
{{< /fig-quote >}}

{{< fig-quote title="Problem 1 - PukiWiki" link="http://odz.sakura.ne.jp/projecteuler/index.php?cmd=read&page=Problem%201" >}}
<q>10未満の自然数のうち, 3 もしくは 5 の倍数になっているものは 3, 5, 6, 9 の4つがあり, これらの合計は 23 になる.<br>
同じようにして, 1000 未満の 3 か 5 の倍数になっている数字の合計を求めよ.</q>
{{< /fig-quote >}}

まず，ベタに解いたコード。

```go
//Answer0 returns answer to this problem
func Answer0(max int) int {
    sum := 0
    for n := 1; n < max; n++ {
        if (n%3) == 0 || (n%5) == 0 {
            sum += n
        }
    }
    return sum
}
```

これをテストにかける。

```go
package problem1

import "testing"

func TestAnswer(t *testing.T) {
    testCases := []struct {
        max, answer int
    }{
        {max: 10, answer: 23},
    }

    for _, tc := range testCases {
        a := Answer0(tc.max)
        if a != tc.answer {
            t.Errorf("Answer0(%v) = %v, want %v.", tc.max, a, tc.answer)
        }
    }
}
```

実行すると以下の通り。

```text
$ go test -v .
=== RUN   TestAnswer
--- PASS: TestAnswer (0.00s)
PASS
ok      project-euler/problem-1   0.128s
```

これで最初の例示について正しい答えを出していることがわかる。
では1,000未満についてもやってみようか。
以下のようにテスト項目を追加する。

{{< highlight go "hl_lines=6" >}}
func TestAnswer(t * testing.T) {
    testCases := []struct {
        max, answer int
    }{
        {max: 10, answer: 23},
        {max: 1000, answer: 0},
    }

    for _ , tc := range testCases {
        a := Answer0(tc.max)
        if a != tc.answer {
            t.Errorf("Answer0(%v) = %v, want %v.", tc.max, a, tc.answer)
        }
    }
}
{{< /highlight >}}

まだ正解が分からないので，とりあえず `0` をぶちこんでいる。
これを実行する。

```text
$ go test -v .
=== RUN   TestAnswer
--- FAIL: TestAnswer (0.00s)
        answer_test.go:16: Answer0(1000) = ******, want 0.
FAIL
FAIL    project-euler/problem-1   0.106s
exit status 1
```

テストは当然失敗するが，回答が表示されるので[^ans1] これを [Project Euler] のサイトで入力してチェックする。
正解が分かったら改めて正解をテストにセットしなおす。

[^ans1]: 一応ネタバレ防止用に伏字にしている。ご容赦。

{{< highlight go "hl_lines=6" >}}
func TestAnswer(t * testing.T) {
    testCases := []struct {
        max, answer int
    }{
        {max: 10, answer: 23},
        {max: 1000, answer: ******},
    }

    for _ , tc := range testCases {
        a := Answer0(tc.max)
        if a != tc.answer {
            t.Errorf("Answer0(%v) = %v, want %v.", tc.max, a, tc.answer)
        }
    }
}
{{< /highlight >}}

これで再テストすると

```text
$ go test -v .
=== RUN   TestAnswer
--- PASS: TestAnswer (0.00s)
PASS
ok      project-euler/problem-1   0.128s
```

ちゃんと PASS していることが確認できた。

## 改良版

ここで解説を参考に改良版のコードを組む。
こんな感じ。

```go
//Answer1 returns answer to this problem (refactoring version)
func Answer1(max int) int {
    return sumDivisibleBy(max, 3) + sumDivisibleBy(max, 5) - sumDivisibleBy(max, 3*5)
}

func sumDivisibleBy(max, n int) int {
    m := (max - 1) / n
    return n * (m * (m + 1)) / 2
}
```

このコードも同じテストにかける。

{{< highlight go "hl_lines=14-17" >}}
func TestAnswer(t * testing.T) {
    testCases := []struct {
        max, answer int
    }{
        {max: 10, answer: 23},
        {max: 1000, answer: ******},
    }

    for _ , tc := range testCases {
        a := Answer0(tc.max)
        if a != tc.answer {
            t.Errorf("Answer0(%v) = %v, want %v.", tc.max, a, tc.answer)
        }
        a = Answer1(tc.max)
        if a != tc.answer {
            t.Errorf("Answer1(%v) = %v, want %v.", tc.max, a, tc.answer)
        }
    }
}
{{< /highlight >}}

全てのテストケースでパスすればOK。

## ベンチマーク・テスト

更に `Answer0()`, `Answer1()` 両関数をベンチマーク・テストで比較してみることにする。
ベンチマーク・テスト用のコードは以下の通り[^bm1]。

[^bm1]: ベンチマーク用のコードは `***_test.go` ファイルに含めること。また関数名は `Benchmark` から始まる名前にする。

```go
func BenchmarkAnswer0(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = Answer0(1000)
    }
}

func BenchmarkAnswer1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = Answer1(1000)
    }
}
```

ではベンチマーク・テストを実行してみる。

```text
$ go test -bench . -benchmem
goos: windows
goarch: amd64
pkg: project-euler/problem-1
BenchmarkAnswer0-4        500000              2415 ns/op               0 B/op          0 allocs/op
BenchmarkAnswer1-4      2000000000               0.29 ns/op            0 B/op          0 allocs/op
PASS
ok      project-euler/problem-1   1.952s
```

おおっ。
万倍ちがうぜ（笑）

このようにテストフレームワークを使って楽しく [Project Euler] で遊ぶことができる。
テストを書く習慣づけになればもっといいんだろうけど。

## ブックマーク

- [Go 言語のテスト・フレームワーク]({{<relref "testing.md" >}})
- [文字列連結はどれが速い？]({{< relref "join-strings.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[Project Euler]: https://projecteuler.net/
[`testing`]: http://golang.org/pkg/testing/

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253"><img src="https://images-fe.ssl-images-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
	<dd>Alan A.A. Donovan, Brian W. Kernighan</dd>
	<dd>柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>Book 単行本（ソフトカバー）</dd>
    <dd>ASIN: 4621300253, EAN: 9784621300251</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
