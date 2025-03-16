+++
title = "Go 1.24 からのベンチマークテスト"
date =  "2025-03-15T21:08:09+09:00"
description = "for b.Loop() { ... } 内のコードは最適化されないことを保証する"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "book", "testing", "benchmark" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

現在『[効率的なGo]』の読書会に参加してるんだけど，今回はその10回目の話。

- [第67回横浜Go読書会（オンライン） - connpass](https://yokohama-go-reading.connpass.com/event/346187/)

今回は「8章 ベンチマーク」の最初の部分まで読み進めた。
この辺からようやく [Go] のツールを使った具体的な話になってくる。

内容については本を読んでもらうとして，今回はベンチマーク・コードの話。
[件の本][効率的なGo]では「ほとんどのベンチマークに必要な小さなボイラープレートの生テンプレート」として以下のコードを挙げている。

{{< fig-quote class="nobox" type="markdown" title="『効率的なGo』 リスト 8-1：Go のベンチマークの核となる要素" link="https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
```go {hl_lines=["6-8"]}
func BenchmarkSum(b *testing.B) {
    b.ReportAllocs()
    // TODO(bwplotka): 必要な初期化処理を追加

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        // TODO(bwplotka): テストされた機能を追加
    }
}
```
{{< /fig-quote >}}

ちなみに `b.ReportAllocs()` はヒープメモリの割り当て回数と総量をレポートするための関数で， `go test -bench` コマンドでベンチマークテストを起動する際の `-benchmem` オプションと同じ効果がある。

このコード中の `for` 文でカウンターを回しているのだが [Go] 1.24 から [`testing`]`.B.Loop()` メソッドが追加され `for` ループの記述が大きく改善された。

{{< fig-quote type="markdown" title="Go 1.24 Release Notes - The Go Programming Language" link="https://go.dev/doc/go1.24#new-benchmark-function" lang="en" >}}
Benchmarks may now use the faster and less error-prone [`testing.B.Loop`](https://go.dev/pkg/testing#B.Loop) method to perform benchmark iterations like `for b.Loop() { ... }` in place of the typical loop structures involving `b.N` like for range `b.N`. This offers two significant advantages:

- The benchmark function will execute exactly once per `-count`, so expensive setup and cleanup steps execute only once.
- Function call parameters and results are kept alive, preventing the compiler from fully optimizing away the loop body.
{{< /fig-quote >}}

最初のほうの「コストのかかるセットアップおよびクリーンアップ手順は1回だけ実行される」というのは分かりにくいが，実は [`testing`]`.B.Loop()` メソッドを使うと，先のテンプレートコードにある `b.ResetTimer()` を省略できて，純粋に `for b.Loop() { ... }` 内の処理だけ計測される。

もうひとつはもっと重要で `for b.Loop() { ... }` 内のコードは最適化されないことを保証するというものだ。
たとえば

```go
func Add(a, b int) int {
	return a + b
}
```

という関数のベンチマークを取ろうとして素朴に

```go
for range b.N {
    _ = Add(1, 2)
}
```

なんてなコードを書いたら，最悪の場合，最適化でループ内の処理がまるごと無くなりかねない[^b1]。
[`testing`]`.B.Loop()` メソッドを使うことでそうした事態を回避できるというわけだ。

[^b1]: 実際に試してみたが10億回ループしてたので，やはり中身が消えてると思われる。ちなみにループ部分を `for b.Loop() { ... }` にしたら6億回のループになったので，多分そういうことなんだろう。

お試しに何か書いてみよう[^c1]。
まずはこんな関数を書いてみる。

[^c1]: 今回の記事のコードは [`testing` パッケージ][`testing`]のドキュメントにあるサンプルコードを参考にしている。

```go
package sample

func Sum(data []int) int {
	total := 0
	for _, value := range data {
		total += value
	}
	return total
}
```

よくある合計値を返す関数ですな。
作成した `sample.Sum()` 関数のベンチマークテストのコードは以下の通り。

```go {hl_lines=["21-23"]}
package sample_test

import (
	"math/rand/v2"
	"sample"
	"testing"
)

func intList(n int) []int {
	list := make([]int, n)
	for i := range list {
		list[i] = rand.Int()
	}
	return list
}

func BenchmarkSum(b *testing.B) {
	b.ReportAllocs()
	input := intList(128 << 10)

	for b.Loop() {
		sample.Sum(input)
	}
}
```

さっそく実行してみよう[^B2]。

[^B2]: `-run '^$'` オプションはひとつもテストを実行しないことを示す。 `-bench '^BenchmarkSum$'` オプションは指定したベンチマークテストコードを実行することを示す。 `-count 6` は指定したベンチマークテストを6回繰り返すことを示す。ベンチマークテストを繰り返すのは統計処理を行うため。

```text
$ go test -run '^$' -bench '^BenchmarkSum$' -count 6 | tee sample_sum_bench_v1.txt
goos: linux
goarch: amd64
pkg: sample
cpu: AMD Ryzen 5 PRO 4650G with Radeon Graphics
BenchmarkSum-12        37249         32302 ns/op           0 B/op          0 allocs/op
BenchmarkSum-12        36802         32570 ns/op           0 B/op          0 allocs/op
BenchmarkSum-12        36495         34515 ns/op           0 B/op          0 allocs/op
BenchmarkSum-12        36825         32516 ns/op           0 B/op          0 allocs/op
BenchmarkSum-12        37101         32142 ns/op           0 B/op          0 allocs/op
BenchmarkSum-12        37298         32072 ns/op           0 B/op          0 allocs/op
PASS
ok      sample	7.260s
```

ヒープの使用もカウントされてないし，上手くいってるっぽいな。

もうひとつ。
『[効率的なGo]』では [benchstat] コマンドの紹介もされていた。
上のベンチマーク結果について簡単な統計処理をしてくれるらしい。
これも試してみよう。

まずはコマンドのインストール。

```
$ go install golang.org/x/perf/cmd/benchstat@latest
```

先程のベンチマーク結果を [benchstat] コマンドに食わせてみる。

```text
$ benchstat sample_sum_bench_v1.txt
goos: linux
goarch: amd64
pkg: sample
cpu: AMD Ryzen 5 PRO 4650G with Radeon Graphics
       │ sample_sum_bench_v1.txt │
       │         sec/op          │
Sum-12               32.41µ ± 6%

       │ sample_sum_bench_v1.txt │
       │          B/op           │
Sum-12                0.000 ± 0%

       │ sample_sum_bench_v1.txt │
       │        allocs/op        │
Sum-12                0.000 ± 0%
```

んー。
$6\\,\\%$ の分散はちょっと大きいかなぁ。
『[効率的なGo]』によると，分散が $5\\,\\%$ 以上ある場合は環境ノイズ（バックグラウンドのプロセスとかメモリスワップとか）が大きい可能性があるらしい。
まぁ，今回はこのまま進めよう。

[benchstat] で統計処理を行う場合は `-count` オプションを使って少なくとも6回以上は繰り返すべきと書かれている。
これによって環境ノイズを検出しやすくなる。

[benchstat] コマンドは複数のベンチマーク結果を比較することもできる。

たとえば，先ほどの `Sum()` 関数のループ回数を半分にしたら速くなるだろうか。
試してみよう。
まず `Sum()` 関数を以下のように書き直す（ほとんど GitHub Copilot が書いたけど`w`）。

```go
func Sum(data []int) int {
	total := 0
	l := len(data)
	h := l / 2
	for i := range h {
		total += data[i] + data[l-i-1]
	}
	if h*2 != l {
		total += data[h]
	}
	return total
}
```

これに対して先ほどと同じ条件でベンチマークテストを行う。

```text
$ go test -run '^$' -bench '^BenchmarkSum$' -count 6 | tee sample_sum_bench_v2.txt
goos: linux
goarch: amd64
pkg: sample
cpu: AMD Ryzen 5 PRO 4650G with Radeon Graphics
BenchmarkSum-12    	   30920	     39051 ns/op	       0 B/op	       0 allocs/op
BenchmarkSum-12    	   30186	     39492 ns/op	       0 B/op	       0 allocs/op
BenchmarkSum-12    	   30282	     39615 ns/op	       0 B/op	       0 allocs/op
BenchmarkSum-12    	   30373	     39535 ns/op	       0 B/op	       0 allocs/op
BenchmarkSum-12    	   30525	     39413 ns/op	       0 B/op	       0 allocs/op
BenchmarkSum-12    	   30379	     39540 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	sample	7.215s
```

ありゃ。
さっきより遅くなっちゃった？ [benchstat] コマンドで確認してみよう。

```text
$ benchstat sample_sum_bench_v1.txt sample_sum_bench_v2.txt
goos: linux
goarch: amd64
pkg: sample
cpu: AMD Ryzen 5 PRO 4650G with Radeon Graphics
       │ sample_sum_bench_v1.txt │      sample_sum_bench_v2.txt       │
       │         sec/op          │   sec/op     vs base               │
Sum-12               32.41µ ± 6%   39.51µ ± 1%  +21.92% (p=0.002 n=6)

       │ sample_sum_bench_v1.txt │    sample_sum_bench_v2.txt    │
       │          B/op           │    B/op     vs base           │
Sum-12                0.000 ± 0%   0.000 ± 0%  ~ (p=1.000 n=6) ¹
¹ all samples are equal

       │ sample_sum_bench_v1.txt │    sample_sum_bench_v2.txt    │
       │        allocs/op        │ allocs/op   vs base           │
Sum-12                0.000 ± 0%   0.000 ± 0%  ~ (p=1.000 n=6) ¹
¹ all samples are equal

```

$22\\,\\%$ 近く遅くなっちゃったよ `orz`

`p` の値は統計的有意性を表す [$p\\,値$](https://ja.wikipedia.org/wiki/P%E5%80%A4 "p値 - Wikipedia")のことで，既定では $\alpha=0.05$ に設定されている（$\alpha$ の値は `-alpha` オプションで変更できる）。
つまり `p` の値が 0.05 より小さければ有意な値であると見なすことができる。

それはともかく，とりあえずループ回数を畳み込むのはなしの方向で。

...てな感じで [benchstat] コマンドを使えば統計学や誤差論の知識がなくてもある程度の判断を下すことができる。

[Go]: https://go.dev/
[効率的なGo]: https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Amazon.co.jp: 効率的なGo ―データ指向によるGoアプリケーションの性能最適化 : Bartłomiej Płotka, 山口 能迪: 本"

[`testing`]: https://pkg.go.dev/testing "testing package - testing - Go Packages"
[benchstat]: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat "benchstat command - golang.org/x/perf/cmd/benchstat - Go Packages"

## 参考図書

{{% review-paapi "4814400535" %}} <!-- 効率的なGo : Efficient Go -->
{{% review-paapi "B0DNYMMBBQ" %}} <!-- Go言語で学ぶ並行プログラミング -->
{{% review-paapi "B0CFL1DK8Q" %}} <!-- Go言語 100Tips -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
