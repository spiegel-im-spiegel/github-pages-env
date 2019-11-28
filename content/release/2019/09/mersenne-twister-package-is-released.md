+++
title = "疑似乱数生成器 spiegel-im-spiegel/mt をリリースした"
date =  "2019-09-22T17:37:19+09:00"
description = "ついカッとなって書いた。反省はしていない。"
image = "/images/attention/go-logo_blue.png"
tags = [ "package", "module", "golang", "random" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

ついカッとなって書いた。
反省はしていない。

- [spiegel-im-spiegel/mt: Mersenne Twister; Pseudo Random Number Generator, Implemented by Golang](https://github.com/spiegel-im-spiegel/mt)

[Mersenne Twister] の [Go 言語]実装はいくつかあるのだが，やっぱ他人が作る道具は使いにくいよね，というわけで自分で書いてしまった（笑）
[spiegel-im-spiegel/mt] の特徴は以下の通り。

- [math/rand] 互換で [`rand`].`Rand` のソースとして利用できる
- goroutine-safe な構成にできる（[`mt`].`PRNG` 型を利用した場合）

使い方は以下を参照のこと。

- [疑似乱数生成器 spiegel-im-spiegel/mt]({{< ref "/release/mersenne-twister-by-golang.md" >}})

一応，ベンチマークテストもしてみた。

```go
package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/spiegel-im-spiegel/mt"
	"github.com/spiegel-im-spiegel/mt/mt19937"
)

const count = 10000000

func BenchmarkRandomALFG(b *testing.B) {
	rnd := rand.NewSource(time.Now().UnixNano()).(rand.Source64)
	b.ResetTimer()
	for i := 0; i < count; i++ {
		rnd.Uint64()
	}
}

func BenchmarkRandomMT19917(b *testing.B) {
	rnd := mt19937.NewSource(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < count; i++ {
		rnd.Uint64()
	}
}

func BenchmarkRandomALFGRand(b *testing.B) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	b.ResetTimer()
	for i := 0; i < count; i++ {
		rnd.Uint64()
	}
}

func BenchmarkRandomMT19917Rand(b *testing.B) {
	rnd := rand.New(mt19937.NewSource(time.Now().UnixNano()))
	b.ResetTimer()
	for i := 0; i < count; i++ {
		rnd.Uint64()
	}
}

func BenchmarkRandomALFGLocked(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < count; i++ {
		rand.Uint64()
	}
}

func BenchmarkRandomMT19917Locked(b *testing.B) {
	rnd := mt.New(mt19937.NewSource(time.Now().UnixNano()))
	b.ResetTimer()
	for i := 0; i < count; i++ {
		rnd.Uint64()
	}
}
```

テスト対象は以下の通り。

| テスト名                       | 対象                                        |
| ------------------------------ | ------------------------------------------- |
| `BenchmarkRandomALFG`          | [math/rand] 標準アルゴリズム[^rnd1]         |
| `BenchmarkRandomMT19917`       | [`mt`]`/mt19937` パッケージ                 |
| `BenchmarkRandomALFGRand`      | [math/rand] （[`rand`].`Rand` ラッパ）      |
| `BenchmarkRandomMT19917Rand`   | [`mt`]`/mt19937` （[`rand`].`Rand` ラッパ） |
| `BenchmarkRandomALFGLocked`    | [math/rand] Sync バージョン                 |
| `BenchmarkRandomMT19917Locked` | [`mt`]`/mt19937` ＋ [`mt`].`PRNG`           |

[^rnd1]: [math/rand] パッケージに実装されている擬似乱数生成器は[ラグ付フィボナッチ法（Lagged Fibonacci Generator）のバリエーション](https://groups.google.com/forum/#!topic/golang-nuts/RZ1G3_cxMcM)らしい。

このベンチマークテストの実行結果は以下の通り。

```text
$ go test -bench Random -benchmem
goos: linux
goarch: amd64
pkg: github.com/spiegel-im-spiegel/mt/benchmark
BenchmarkRandomALFG-4            	1000000000	         0.0492 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandomMT19917-4         	1000000000	         0.0651 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandomALFGRand-4        	1000000000	         0.0749 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandomMT19917Rand-4     	1000000000	         0.0846 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandomALFGLocked-4      	1000000000	         0.176 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandomMT19917Locked-4   	1000000000	         0.192 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/spiegel-im-spiegel/mt/benchmark	7.081s
```

というわけで [math/rand] のほうが若干速いかな。
乱数としての性能は別の機会に。

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Mersenne Twister]: http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/mt.html "Mersenne Twister: A random number generator (since 1997/10)"
[math/rand]: https://golang.org/pkg/math/rand/ "rand - The Go Programming Language"
[`rand`]: https://golang.org/pkg/math/rand/ "rand - The Go Programming Language"
[spiegel-im-spiegel/mt]: https://github.com/spiegel-im-spiegel/mt "spiegel-im-spiegel/mt: Mersenne Twister; Pseudo Random Number Generator, Implemented by Golang"
[`mt`]: https://github.com/spiegel-im-spiegel/mt "spiegel-im-spiegel/mt: Mersenne Twister; Pseudo Random Number Generator, Implemented by Golang"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN), 4621300253 (ISBN), 9784621300251 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%95%B0%E5%AD%A6%E3%82%AC%E3%83%BC%E3%83%AB%EF%BC%8F%E4%B9%B1%E6%8A%9E%E3%82%A2%E3%83%AB%E3%82%B4%E3%83%AA%E3%82%BA%E3%83%A0-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B00I8AT1FO?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00I8AT1FO"><img src="https://images-fe.ssl-images-amazon.com/images/I/41353H%2BBzFL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%95%B0%E5%AD%A6%E3%82%AC%E3%83%BC%E3%83%AB%EF%BC%8F%E4%B9%B1%E6%8A%9E%E3%82%A2%E3%83%AB%E3%82%B4%E3%83%AA%E3%82%BA%E3%83%A0-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B00I8AT1FO?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00I8AT1FO">数学ガール／乱択アルゴリズム</a></dt>
    <dd>結城 浩</dd>
    <dd>SBクリエイティブ 2011-02-25 (Release 2014-03-12)</dd>
    <dd>Kindle版</dd>
    <dd>B00I8AT1FO (ASIN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">工学ガール，リサちゃん登場！</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-04-19">2015-04-19</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-API</a>)</p>
</div>
