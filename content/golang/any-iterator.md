+++
title = "任意のイテレータを作る"
date =  "2026-07-06T22:53:06+09:00"
description = "配列や連想配列以外で for-range 構文が使えれば何かと便利だろう。"
isCJKLanguage = true
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "map", "slice", "koyomi", "random", "time" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] でイテレータを組む機会があったので，思い出しつつ，簡単なメモを残しておく。
なお [Go] のイテレータは 1.23 から導入されている。

たとえば配列（または `slice`）や連想配列（`map`）であれば for-range 構文で以下のように書ける。

```go
package main

func main() {
	nums := []int{1, 2, 3}
	for _, num := range nums {
		println(num)
	}

	planets := map[string]float64{
		"Mercury": 0.055,
		"Venus":   0.815,
		"Earth":   1.0,
	}
	for planet, mass := range planets {
		println(planet, ":", mass)
	}
}
```

配列や連想配列以外でも for-range 構文が使えれば何かと便利だろう。
[Go] ではこれを range over function という仕組みで実現している。

- [Range Over Function Types - The Go Programming Language](https://go.dev/blog/range-functions)

## Range Over Functions

具体的には以下の3つの関数型を使う。

{{< fig-quote type="markdown" title="Range Over Function Types" link="https://go.dev/blog/range-functions" lang="en" >}}
The improved `for/range` statement doesn’t support arbitrary function types. As of Go 1.23 it now supports ranging over functions that take a single argument. The single argument must itself be a function that takes zero to two arguments and returns a `bool`; by convention, we call it the yield function.

```go
func(yield func() bool)
func(yield func(V) bool)
func(yield func(K, V) bool)
```
{{< /fig-quote >}}

例として，3つの乱数を返すイテレータを作ってみよう。
こんな感じかな。

```go {hl_lines=["7-11"]}
package main

import "math/rand/v2"

func randoms(count int) func(func(int) bool) {
	return func(yield func(int) bool) {
		for range count { // range over int (from 0 to count-1)
			if !yield(rand.Int()) {
				return
			}
		}
	}
}

func main() {
	for n := range randoms(3) {
		println(n)
	}
}
```

強調部分がイテレータの本体である。
これを実行すると，以下のような出力になる。

```text
$ go run sample1a.go
27723168966634698
2343875658410196265
4739277610031835691
```

このように `yield()` 関数を呼び出した回数だけ for-range ループが回っていることが分かる。
`yield()` 関数は何をしているかというと，引数の値をループに push して，処理を継続する場合は `true` を返している。

ちなみに `yield()` 関数の返り値のチェックをサボるとどうなるかというと，たとえば

```go {hl_lines=[8, "16-18"]}
package main

import "math/rand/v2"

func randoms2(count int) func(func(int, int) bool) {
	return func(yield func(int, int) bool) {
		for i := range count { // range over int (from 0 to count-1)
			yield(i+1, rand.Int())
		}
	}
}

func main() {
	for i, n := range randoms2(3) {
		println(i, n)
		if i >= 2 {
			break
		}
	}
}
```

のように，3回出力するイテレータに対して何らかの理由で2回で中断したとすると

```text
$ go run sample1b.go
1 5146054532611686492
2 7588979459646265126
panic: runtime error: range function continued iteration after function for loop body returned false

goroutine 1 [running]:
main.main-range1(...)
	/path/to/sample1b.go:16
main.main.randoms2.func1(...)
	/path/to/sample1b.go:10
main.main()
	/path/to/sample1b.go:16 +0xcd
exit status 2
```

などと panic を吐く。
怖いですねぇ（笑）

試しに上のコードの `randoms2()` 関数を

```go {hl_lines=["4-8"]}
func randoms2(count int) func(func(int, int) bool) {
	return func(yield func(int, int) bool) {
		for i := range count { // range over int (from 0 to count-1)
			println("calling", i+1)
			if !yield(i+1, rand.Int()) {
				println("return false from yield")
				return
			}
		}
	}
}
```

などとデバッグ文を入れて書き直すと

```text
$ go run sample1c.go
calling 1
1 682324887426579321
calling 2
2 5758210549060056534
return false from yield
```

のように出力される。
`yield()` 関数の呼び出しは2回までで，2回目に `false` を返していることが分かるだろう。

## iter パッケージ

イテレータを比較的簡単に取り扱えるようにするため [`iter`] パッケージが用意されていて，イテレータを以下のように定義している。

```go
type Seq[V any] func(yield func(V) bool)
```

```go
type Seq2[K, V any] func(yield func(K, V) bool)
```

前節の `randoms()` 関数であれば

```go {hl_lines=[6]}
import (
	"iter"
	"math/rand/v2"
)

func randoms(count int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for range count { // range over int (from 0 to count-1)
			if !yield(rand.Int()) {
				return
			}
		}
	}
}
```

のように書き直せる。
ん？ 大して変わらんな（笑）

もうひとつ。
[`iter`] パッケージには `iter.Pull()` および `iter.Pull2()` 関数が用意されていて，たとえば先程の `randoms()` 関数の呼び出し側で，以下のように，push-style から pull-style に変形できる。

```go {hl_lines=[2]}
func main() {
	next, stop := iter.Pull(randoms(3))
	defer stop()
	for {
		n, ok := next()
		if !ok {
			break
		}
		println(n)
	}
}
```

これはアリだろう。

## slices パッケージ

[`slices`] パッケージにもイテレータ関連の関数が追加された。
たとえば `All()` 関数を使って

```go {hl_lines=[7]}
package main

import "slices"

func main() {
	nums := []int{1, 2, 3}
	for i, num := range slices.All(nums) {
		println(i, ":", num)
	}
}
```

のように，配列や slice をイテレータに変換する。
記述が増えただけで見た目は変わらんやんけ！ と思うかも知れないが

```go {hl_lines=[10]}
package main

import (
	"iter"
	"slices"
)

func main() {
	nums := []int{1, 2, 3}
	next, stop := iter.Pull(slices.Values(nums))
	defer stop()
	for {
		num, ok := next()
		if !ok {
			break
		}
		println(num)
	}
}
```

のように pull-style に変形できるのはいいかもしれない。
また，イテレータに変換する際に `Backward()` 関数で逆順に並び替えたり `Chunk()` 関数で変形することもできる。

イテレータを `slice` に変換することもできる。
前節の `randoms()` 関数を使って

```go
func main() {
	randlist := slices.Collect(randoms(3))
	fmt.Println("randlist:", randlist)
}
```

のように書く。
また `AppendSeq()` 関数で既にある `slice` インスタンスにイテレータの出力を追加したり，`Sorted()` 関数等を使った並び替えや `Chunk()` 関数を使った変形もできる。

## maps パッケージ

[`maps`] パッケージにもイテレータ関連の関数が追加された。
ここでは簡単に関数名だけ紹介する。

- `All` : `map` 型を [`iter`]`.Seq2` 型に変換する
- `Keys` : `map` 型のキーを [`iter`]`.Seq` 型に変換する
- `Values` : `map` 型の値を [`iter`]`.Seq` 型に変換する
- `Collect` : [`iter`]`.Seq2` 型を `map` 型に変換する
- `Insert` : 既にある `map` インスタンスに [`iter`]`.Seq2` 型の出力を追加する

## 改訂版： 土用期間中の干支を数え上げる

以前書いた「[第五の季節：土用]({{< ref "/remark/2022/04/doyo-period.md" >}})」という記事で，拙作の [`goark/koyomi`] パッケージを使って，土用期間中の干支を数え上げるコードを紹介した。
こんな感じ。

```go
package main

import (
	"fmt"
	"time"

	"github.com/goark/koyomi/value"
	"github.com/goark/koyomi/zodiac"
)

func main() {
	start := value.NewDateYMD(2026, time.July, 20)   //土用の入り
	end := value.NewDateYMD(2026, time.August, 7)    //立秋
	for d := start; d.Before(end); d = d.AddDay(1) { //1日単位で立秋の前日まで出力する
		干, 支 := zodiac.ZodiacDayNumber(d)
		fmt.Printf("Day %v is %v%v\n", d, 干, 支)
	}
}
```

実行結果は以下の通り。

```text
$ go run sample6a.go
Day 2026-07-20 is 乙未
Day 2026-07-21 is 丙申
Day 2026-07-22 is 丁酉
Day 2026-07-23 is 戊戌
Day 2026-07-24 is 己亥
Day 2026-07-25 is 庚子
Day 2026-07-26 is 辛丑
Day 2026-07-27 is 壬寅
Day 2026-07-28 is 癸卯
Day 2026-07-29 is 甲辰
Day 2026-07-30 is 乙巳
Day 2026-07-31 is 丙午
Day 2026-08-01 is 丁未
Day 2026-08-02 is 戊申
Day 2026-08-03 is 己酉
Day 2026-08-04 is 庚戌
Day 2026-08-05 is 辛亥
Day 2026-08-06 is 壬子
```

次の土用の丑の日は7月26日か。

それはともかく，コード中の for 文がカッコ悪いなぁと以前から思ってたので，今回の勉強ついでに，指定期間の日付を出力するイテレータを実装してみた。
これを使うと，こんな感じに書き直せる。

```go {hl_lines=["12-20"]}
package main

import (
	"fmt"
	"time"

	"github.com/goark/koyomi/value"
	"github.com/goark/koyomi/zodiac"
)

func main() {
	seq, err := value.NewDateYMD(2026, time.July, 20).IterDay( //土用の入りから
		1, //1日単位で
		value.NewDateYMD(2026, time.August, 7).AddDay(-1), //立秋の前日まで出力する
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range seq {
		干, 支 := zodiac.ZodiacDayNumber(d)
		fmt.Printf("Day %v is %v%v\n", d, 干, 支)
	}
}
```

どうよ。
こっちのほうが少し見やすいべ。
`Before()` とか `After()` とか[直感的に分かりにくいメソッド]({{< relref "./order-by-time.md" >}} "time.Time の比較が覚えれん！")が隠蔽できるだけでもだいぶありがたい。


## ブックマーク

- [Go1.23のイテレータを完全理解する✌️](https://zenn.dev/kkkxxx/articles/d9505540581b5d)
- [Go1.23のイテレータを動かして学ぶ](https://zenn.dev/team_soda/articles/understanding-iterators-in-go)

- [十干十二支を数え上げるパッケージを作ってみた]({{< ref "/release/2021/07/japanese-zodiac.md" >}})
- [Go 1.22 における疑似乱数生成器]({{< relref "./pseudo-random-number-generator-v2.md" >}})

[Go]: https://go.dev/
[`iter`]: https://pkg.go.dev/iter "iter package - iter - Go Packages"
[`slices`]: https://pkg.go.dev/slices "slices package - slices - Go Packages"
[`maps`]: https://pkg.go.dev/maps "maps package - maps - Go Packages"
[`goark/koyomi`]: https://github.com/goark/koyomi "goark/koyomi: 日本のこよみ"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400535" %}} <!-- 効率的なGo : Efficient Go -->
{{% review-paapi "B0DNYMMBBQ" %}} <!-- Go言語で学ぶ並行プログラミング -->
