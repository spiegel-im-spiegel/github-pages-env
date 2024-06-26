+++
title = "Go の疑似乱数生成器は並行的に安全ではないらしい（追記あり）"
date =  "2019-09-17T23:27:18+09:00"
description = "件の記事では解決方法が（具体的には）示されていないので，いくつか対策を考えてみよう。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "programming", "concurrency", "goroutine", "channel", "random" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

面白い記事みっけ！

- [【Go】rand.Sourceを並列で使いまわすなんて何事だ - Qiita](https://qiita.com/hiromichi_n/items/d0636b9444dca18ef357)

折角なので便乗記事を書いてみる。

まぁ，内部状態を持つオブジェクトは，状態が変わらない（immutable）か操作が並行的に安全（concurrency safe）であることが仕様・設計として明確であるものでない限り，複数の goroutine 間でインスタンスを共有してはいけない，というのは基本中の基本である。

ましてや標準の [math/rand] パッケージは [`rand`]`.Source` インタフェースを満たすのであればユーザ側で任意のアルゴリズムを用意することもできるので，並行的に安全であることを期待するほうが間違っているとも言える。

まずは，[件の記事]で書かれているコードを挙げておこう。
ただし動作に直接関係ないコードは極力省いている。

```go
package main

import (
    "math/rand"
    "sync"
    "time"
)

var randSource = NewRandSource()

func NewRandSource() *rand.Rand {
    return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func calcRand() {
    for i := 0; i < 10000; i++ {
        randSource.Intn(1000)
    }
}

func main() {
    wg := sync.WaitGroup{}
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            calcRand()
            wg.Done()
        }()
    }
    wg.Wait()
}
```

ポイントは [`rand`]`.Rand` インスタンスを初期化時にひとつだけ生成し，複数の goroutine で使い回している点である。
これを実行すると以下のように panic を吐く。

```text
$ go run -trimpath sample.go
panic: runtime error: index out of range [-1]

goroutine 94 [running]:
math/rand.(*rngSource).Uint64(...)
    math/rand/rng.go:249
math/rand.(*rngSource).Int63(0xc000083500, 0x50321535775976c1)
    math/rand/rng.go:234 +0x93
math/rand.(*Rand).Int63(...)
    math/rand/rand.go:85
math/rand.(*Rand).Int31(...)
    math/rand/rand.go:99
math/rand.(*Rand).Int31n(0xc000088090, 0x3e8, 0x1fd)
    math/rand/rand.go:134 +0x5f
math/rand.(*Rand).Intn(0xc000088090, 0x3e8, 0x1fd)
    math/rand/rand.go:172 +0x45
main.calcRand()
    sample@/sample.go:17 +0x3f
main.main.func1(0xc000098000)
    sample@/sample.go:26 +0x22
created by main.main
    sample@/sample.go:25 +0x78
exit status 2
```

panic が発生する仕組みは[件の記事]に分かりやすく解説されているので参照のこと。

## goroutine ごとにインスタンスを生成する

[件の記事]では解決方法が（具体的には）示されていないので，こちらでいくつか考えてみよう。

一番簡単なのは goroutine ごとに [`rand`]`.Rand` インスタンスを生成することだ。
こんな感じに変えたらどうだろう。

{{< highlight go "hl_lines=1 3 12" >}}
func calcRand(rnd* rand.Rand) {
    for i := 0; i < 10000; i++ {
        rnd.Intn(1000)
    }
}

func main() {
    wg := sync.WaitGroup{}
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            calcRand(NewRandSource())
            wg.Done()
        }()
    }
    wg.Wait()
}
{{< /highlight >}}

これで必要十分に機能するし，少なくとも panic は起こらない。
このやり方の欠点は（goroutine ごとに [`rand`]`.Rand` インスタンスが生成されるため）元のコードより（僅かだが）高コストになることと，疑似乱数生成器の性能がアルゴリズムだけでなく seed の選び方にも依存する，というあたりだろうか。

まぁ [math/rand] の標準アルゴリズム[^rnd1] であれば性能に関してはさしたる問題にはならないだろう。

[^rnd1]: [math/rand] パッケージに実装されている擬似乱数生成器は[ラグ付フィボナッチ法（Lagged Fibonacci Generator）のバリエーション](https://groups.google.com/forum/#!topic/golang-nuts/RZ1G3_cxMcM)らしい。

## Generator Pattern を使う

今回の例ではあまりオススメではないのだが，並行処理の Generator Pattern を使う手もある。

まず `NewRandSource()` 関数を以下の関数で置き換える。

```go
func NewGenerator() <-chan int {
    rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
    ch := make(chan int)
    go func() {
        defer close(ch)
        for {
            ch <- rnd.Intn(1000)
        }
    }()
    return ch
}
```

こうすれば [`rand`]`.Rand` インスタンスはひとつで済むし（seed もひとつ），持ち回すインスタンスは channel のみなので並行的に安全にできる。
乱数の取り出し側はこう書き換えればよい。

{{< highlight go "hl_lines=1 3-5 11 15" >}}
func calcRand(gen <-chan int) {
    for i := 0; i < 10000; i++ {
        if _ , ok := <-gen; !ok {
            return
        }
    }
}

func main() {
    wg := sync.WaitGroup{}
    gen := NewGenerator()
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            calcRand(gen)
            wg.Done()
        }()
    }
    wg.Wait()
}
{{< /highlight >}}

このコードの欠点は「遅い」ことに尽きる。
まぁ channel の読み書きで同期を取る必要があるから遅くなって当たり前だけど。

今回のようなケースではなく，例えば generator がハードウェア制御を伴うものだったり [singleton]({{< relref "./singleton-pattern.md" >}}) を含む処理だったり channel の読み書きにかかるコストに対して他の処理が相対的に大きくなったり ...などなど，状況によっては Generator Pattern のほうが有利になる場合もあるだろう。

Generator Pattern は平行処理のデザインパターンの中では比較的単純なものだが応用範囲が広い。
[Go 言語]の goroutine 自体は（OS スレッドなどと比べて）かなり安価で手軽に構成できるので，積極的に試してみるといいと思う。

### おまけの追記

そうそう。
上の `NewGenerator()` 関数で生成・駆動される goroutine は自力で終了できない。
なので，以下のように

```go
func NewGenerator(ctx context.Context) <-chan int {
    rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
    ch := make(chan int)
    go func() {
        defer close(ch)
        for {
            select {
            case <-ctx.Done():
                return
            default:
                ch <- rnd.Intn(1000)
            }
        }
    }()
    return ch
}
```

外部からキャンセルイベントを流し込めるようにするといいかもしれない。

## 【2019-09-20 追記】 実は標準で並行的に安全な疑似乱数生成器が用意されていた

あれから [math/rand] のソースコードを眺めてて気がついたのだが，実は並行的に安全な疑似乱数生成器が標準で用意されていた。

たとえば [`rand`]`.Intn()` 関数を見ると

```go
// Intn returns, as an int, a non-negative pseudo-random number in [0,n)
// from the default Source.
// It panics if n <= 0.
func Intn(n int) int { return globalRand.Intn(n) }
```

とか書かれていて，じゃあ `globalRand` って何なん？ と思って見てみたら

```go
type lockedSource struct {
	lk  sync.Mutex
	src Source64
}


func (r *lockedSource) Int63() (n int64) {
	r.lk.Lock()
	n = r.src.Int63()
	r.lk.Unlock()
	return
}

func (r *lockedSource) Uint64() (n uint64) {
	r.lk.Lock()
	n = r.src.Uint64()
	r.lk.Unlock()
	return
}

func (r *lockedSource) Seed(seed int64) {
	r.lk.Lock()
	r.src.Seed(seed)
	r.lk.Unlock()
}

...

var globalRand = New(&lockedSource{src: NewSource(1).(Source64)})
```

とか書かれているわけですよ。
なんだ，ちゃんと [`sync`]`.Mutex` で排他制御してるんぢゃん。

というわけで，最初のコードは

```go
package main

import (
	"math/rand"
	"sync"
	"time"
)

func calcRnad() {
	for i := 0; i < 10000; i++ {
		rand.Intn(1000)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			calcRnad()
			wg.Done()
		}()
	}
	wg.Wait()
}
```

と書けば panic を吐くことなくちゃんと終了する。
若干遅くはなるけど，それでも Generator Pattern を使うよりは全然速い。

## ブックマーク

- [Go の channel 処理パターン集](https://hori-ryota.com/blog/golang-channel-pattern/)

- [モンテカルロ法による円周率の推定（その4 PRNG）]({{< ref "http://localhost:1313/golang/estimate-of-pi-4-prng.md" >}})

[Go]: https://go.dev/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[件の記事]: https://qiita.com/hiromichi_n/items/d0636b9444dca18ef357 "【Go】rand.Sourceを並列で使いまわすなんて何事だ - Qiita"
[math/rand]: https://golang.org/pkg/math/rand/ "rand - The Go Programming Language"
[`rand`]: https://golang.org/pkg/math/rand/ "rand - The Go Programming Language"
[`sync`]: https://golang.org/pkg/sync/ "sync - The Go Programming Language"

## 参考図書

{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
