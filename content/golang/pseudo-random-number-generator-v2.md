+++
title = "Go 1.22 における疑似乱数生成器"
date =  "2024-03-07T22:20:40+09:00"
description = "時代は math/rand/v2 かな"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "programming", "random" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

[Go] 1.22 の [`math/rand`] パッケージと追加された [`math/rand/v2`] パッケージを眺めている。

おそらく [Go] 1.22 における疑似乱数関連の最大のトピックは ChaCha8 がランタイムに組み込まれ，疑似乱数生成器の既定アルゴリズムになったことだろう。

- {{< pdf-file title="The Salsa20 family of stream ciphers" link="https://cr.yp.to/snuffle/salsafamily-20071225.pdf" >}} : これがベースになる論文かな
- [Salsa20 - Wikipedia](https://en.wikipedia.org/wiki/Salsa20)
- [C2SP/chacha8rand.md at main · C2SP/C2SP · GitHub](https://github.com/C2SP/C2SP/blob/main/chacha8rand.md)
- [つくって理解するストリーム暗号 ChaCha20 - ちりもつもればミルキーウェイ](https://convto.hatenablog.com/entry/2024/02/26/121013) : [Go] でサンプルコードを書いておられる。ありがたや

ChaCha はストリーム暗号の一種で，簡単に言うと，疑似乱数を生成してそれを平文と XOR するというものらしい。
このうちの疑似乱数を生成する部分を切り出しているようだ。
ストリーム暗号に使うものなので，暗号技術的にセキュアでかつ速いというのが特徴になるだろうか。
ちなみに ChaCha の後ろについている 20 とか 8 とかはラウンド数を示しているそうな。

ChaCha は OpenSSL だか OpenSSH だかでも見かけたような（うろ覚え）。
もし結城浩さんが『[暗号技術入門](https://www.amazon.co.jp/dp/B015643CPE?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Amazon.co.jp: 暗号技術入門 第3版　秘密の国のアリス eBook : 結城 浩: Kindleストア")』の第4版を出される機会があれば，付録でいいので是非 ChaCha にも言及して欲しい。

## ランタイムに ChaCha8 疑似乱数生成器を組み込む

ChaCha8 疑似乱数生成器のアルゴリズムは [`internal/chacha8rand`] パッケージに実装されている。
中身については割愛させてもらう。
Internal パッケージなので，サードパーティのパッケージからは直接参照できない。

まずは [`runtime`] パッケージに組み込んでいる部分を見てみる。
ちょっと長いけどご容赦。

{{< fig-quote class="nobox" type="markdown" title="runtime/rand.go" link="https://pkg.go.dev/runtime" lang="en" >}}
```go
// OS-specific startup can set startupRand if the OS passes
// random data to the process at startup time.
// For example Linux passes 16 bytes in the auxv vector.
var startupRand []byte

// globalRand holds the global random state.
// It is only used at startup and for creating new m's.
// Otherwise the per-m random state should be used
// by calling goodrand.
var globalRand struct {
    lock  mutex
    seed  [32]byte
    state chacha8rand.State
    init  bool
}

var readRandomFailed bool

// randinit initializes the global random state.
// It must be called before any use of grand.
func randinit() {
    lock(&globalRand.lock)
    if globalRand.init {
        fatal("randinit twice")
    }

    seed := &globalRand.seed
    if startupRand != nil {
        for i, c := range startupRand {
            seed[i%len(seed)] ^= c
        }
        clear(startupRand)
        startupRand = nil
    } else {
        if readRandom(seed[:]) != len(seed) {
            // readRandom should never fail, but if it does we'd rather
            // not make Go binaries completely unusable, so make up
            // some random data based on the current time.
            readRandomFailed = true
            readTimeRandom(seed[:])
        }
    }
    globalRand.state.Init(*seed)
    clear(seed[:])
    globalRand.init = true
    unlock(&globalRand.lock)
}

// readTimeRandom stretches any entropy in the current time
// into entropy the length of r and XORs it into r.
// This is a fallback for when readRandom does not read
// the full requested amount.
// Whatever entropy r already contained is preserved.
func readTimeRandom(r []byte) {
    // Inspired by wyrand.
    // An earlier version of this code used getg().m.procid as well,
    // but note that this is called so early in startup that procid
    // is not initialized yet.
    v := uint64(nanotime())
    for len(r) > 0 {
        v ^= 0xa0761d6478bd642f
        v *= 0xe7037ed1a0b428db
        size := 8
        if len(r) < 8 {
            size = len(r)
        }
        for i := 0; i < size; i++ {
            r[i] ^= byte(v >> (8 * i))
        }
        r = r[size:]
        v = v>>32 | v<<32
    }
}
```
{{< /fig-quote >}}

これは疑似乱数生成器の状態（主に seed）を管理してる部分かな。
最初の seed は乱数デバイスから取ってるんだね。
これに失敗すると時刻から生成する，と。
ユーザ側は明示的に seed を指定する必要がなくなるということやね。

{{< fig-quote class="nobox" type="markdown" title="runtime/rand.go" link="https://pkg.go.dev/runtime" lang="en" >}}
```go
// readTimeRandom stretches any entropy in the current time
// into entropy the length of r and XORs it into r.
// This is a fallback for when readRandom does not read
// the full requested amount.
// Whatever entropy r already contained is preserved.
func readTimeRandom(r []byte) {
    // Inspired by wyrand.
    // An earlier version of this code used getg().m.procid as well,
    // but note that this is called so early in startup that procid
    // is not initialized yet.
    v := uint64(nanotime())
    for len(r) > 0 {
        v ^= 0xa0761d6478bd642f
        v *= 0xe7037ed1a0b428db
        size := 8
        if len(r) < 8 {
            size = len(r)
        }
        for i := 0; i < size; i++ {
            r[i] ^= byte(v >> (8 * i))
        }
        r = r[size:]
        v = v>>32 | v<<32
    }
}

// bootstrapRand returns a random uint64 from the global random generator.
func bootstrapRand() uint64 {
    lock(&globalRand.lock)
    if !globalRand.init {
        fatal("randinit missed")
    }
    for {
        if x, ok := globalRand.state.Next(); ok {
            unlock(&globalRand.lock)
            return x
        }
        globalRand.state.Refill()
    }
}

// bootstrapRandReseed reseeds the bootstrap random number generator,
// clearing from memory any trace of previously returned random numbers.
func bootstrapRandReseed() {
    lock(&globalRand.lock)
    if !globalRand.init {
        fatal("randinit missed")
    }
    globalRand.state.Reseed()
    unlock(&globalRand.lock)
}

// rand32 is uint32(rand()), called from compiler-generated code.
//go:nosplit
func rand32() uint32 {
    return uint32(rand())
}

// rand returns a random uint64 from the per-m chacha8 state.
// Do not change signature: used via linkname from other packages.
//go:nosplit
//go:linkname rand
func rand() uint64 {
    // Note: We avoid acquirem here so that in the fast path
    // there is just a getg, an inlined c.Next, and a return.
    // The performance difference on a 16-core AMD is
    // 3.7ns/call this way versus 4.3ns/call with acquirem (+16%).
    mp := getg().m
    c := &mp.chacha8
    for {
        // Note: c.Next is marked nosplit,
        // so we don't need to use mp.locks
        // on the fast path, which is that the
        // first attempt succeeds.
        x, ok := c.Next()
        if ok {
            return x
        }
        mp.locks++ // hold m even though c.Refill may do stack split checks
        c.Refill()
        mp.locks--
    }
}

// mrandinit initializes the random state of an m.
func mrandinit(mp *m) {
    var seed [4]uint64
    for i := range seed {
        seed[i] = bootstrapRand()
    }
    bootstrapRandReseed() // erase key we just extracted
    mp.chacha8.Init64(seed)
    mp.cheaprand = rand()
}
```
{{< /fig-quote >}}

`mrandinit()` 関数でランタイムを初期化して，それを使って実際に乱数を取得してるのが `rand()` 関数だね。
ふむふむ。

## math/rand パッケージのトップレベル関数群にランタイムの ChaCha8 を組み込む

それじゃあ [`math/rand`] パッケージの方を見てみよう。

{{< fig-quote class="nobox" type="markdown" title="math/rand/rand.go" link="https://pkg.go.dev/math/rand" lang="en" >}}
```go
//go:linkname runtime_rand runtime.rand
func runtime_rand() uint64
```
{{< /fig-quote >}}

`go:linkname` ディレクティブの説明は割愛する。
こうやってリンクしてるということで飲み込んでいただければ（笑） そうそう。
`go:linkname` ディレクティブは [`unsafe`] パッケージを要求するので，真似するときは要注意だよ。

ともかく，この `runtime_rand()` 関数を使って [`rand`][`math/rand`]`.Source` インタフェース互換の構造体 `runtimeSource` を定義している。

{{< fig-quote class="nobox" type="markdown" title="math/rand/rand.go" link="https://pkg.go.dev/math/rand" lang="en" >}}
```go
// runtimeSource is an implementation of Source64 that uses the runtime
// fastrand functions.
type runtimeSource struct {
    // The mutex is used to avoid race conditions in Read.
    mu sync.Mutex
}

func (*runtimeSource) Int63() int64 {
    return int64(runtime_rand() & rngMask)
}

func (*runtimeSource) Seed(int64) {
    panic("internal error: call to runtimeSource.Seed")
}

func (*runtimeSource) Uint64() uint64 {
    return runtime_rand()
}
```
{{< /fig-quote >}}

`Seed()` メソッドを呼び出したら panic 吐くとか容赦ないな（笑） `runtimeSource` はこんな風に使う。

{{< fig-quote class="nobox" type="markdown" title="math/rand/rand.go" link="https://pkg.go.dev/math/rand" lang="en" >}}
```go
// globalRandGenerator is the source of random numbers for the top-level
// convenience functions. When possible it uses the runtime fastrand64
// function to avoid locking. This is not possible if the user called Seed,
// either explicitly or implicitly via GODEBUG=randautoseed=0.
var globalRandGenerator atomic.Pointer[Rand]

var randautoseed = godebug.New("randautoseed")

// globalRand returns the generator to use for the top-level convenience
// functions.
func globalRand() *Rand {
    if r := globalRandGenerator.Load(); r != nil {
        return r
    }

    // This is the first call. Initialize based on GODEBUG.
    var r *Rand
    if randautoseed.Value() == "0" {
        randautoseed.IncNonDefault()
        r = New(new(lockedSource))
        r.Seed(1)
    } else {
        r = &Rand{
            src: &runtimeSource{},
            s64: &runtimeSource{},
        }
    }

    if !globalRandGenerator.CompareAndSwap(nil, r) {
        // Two different goroutines called some top-level
        // function at the same time. While the results in
        // that case are unpredictable, if we just use r here,
        // and we are using a seed, we will most likely return
        // the same value for both calls. That doesn't seem ideal.
        // Just use the first one to get in.
        return globalRandGenerator.Load()
    }

    return r
}
```
{{< /fig-quote >}}

{{< fig-quote class="nobox" type="markdown" title="math/rand/rand.go" link="https://pkg.go.dev/math/rand" lang="en" >}}
```go
// Seed uses the provided seed value to initialize the default Source to a
// deterministic state. Seed values that have the same remainder when
// divided by 2³¹-1 generate the same pseudo-random sequence.
// Seed, unlike the [Rand.Seed] method, is safe for concurrent use.
//
// If Seed is not called, the generator is seeded randomly at program startup.
//
// Prior to Go 1.20, the generator was seeded like Seed(1) at program startup.
// To force the old behavior, call Seed(1) at program startup.
// Alternately, set GODEBUG=randautoseed=0 in the environment
// before making any calls to functions in this package.
//
// Deprecated: As of Go 1.20 there is no reason to call Seed with
// a random value. Programs that call Seed with a known value to get
// a specific sequence of results should use New(NewSource(seed)) to
// obtain a local random generator.
func Seed(seed int64) {
    orig := globalRandGenerator.Load()

    // If we are already using a lockedSource, we can just re-seed it.
    if orig != nil {
        if _, ok := orig.src.(*lockedSource); ok {
            orig.Seed(seed)
            return
        }
    }

    // Otherwise either
    // 1) orig == nil, which is the normal case when Seed is the first
    // top-level function to be called, or
    // 2) orig is already a runtimeSource, in which case we need to change
    // to a lockedSource.
    // Either way we do the same thing.

    r := New(new(lockedSource))
    r.Seed(seed)

    if !globalRandGenerator.CompareAndSwap(orig, r) {
        // Something changed underfoot. Retry to be safe.
        Seed(seed)
    }
}
```
{{< /fig-quote >}}

つまり，環境変数 `GODEBUG` で明示的に指定（`randautoseed=0`）するか最初に [`rand`][`math/rand`]`.Seed()` 関数を呼び出すかしない限りランタイムに組み込んだ ChaCha8 疑似乱数生成器が有効になるっちうわけだ。
ちなみに `lockedSource` は [`math/rand`] パッケージに従来からある疑似乱数生成器で，名前の通り，ちゃんと mutex で排他処理している。

## math/rand/v2 パッケージにおける Source インタフェースの定義

では，いよいよ [Go] 1.22 で追加された [`math/rand/v2`] パッケージを見てみよう。

[`math/rand`] パッケージと [`math/rand/v2`] パッケージとの大きな違いは `rand.Source` インタフェースが非互換になっていることだろう。

[`math/rand`] パッケージの `Source` インタフェースの定義は以下の通り。

{{< fig-quote class="nobox" type="markdown" title="math/rand/rand.go" link="https://pkg.go.dev/math/rand" lang="en" >}}
```go
// A Source represents a source of uniformly-distributed
// pseudo-random int64 values in the range [0, 1<<63).
//
// A Source is not safe for concurrent use by multiple goroutines.
type Source interface {
    Int63() int64
    Seed(seed int64)
}

// A Source64 is a [Source] that can also generate
// uniformly-distributed pseudo-random uint64 values in
// the range [0, 1<<64) directly.
// If a [Rand] r's underlying [Source] s implements Source64,
// then r.Uint64 returns the result of one call to s.Uint64
// instead of making two calls to s.Int63.
type Source64 interface {
    Source
    Uint64() uint64
}
```
{{< /fig-quote >}}

これに対して [`math/rand/v2`] ではこう定義されている。

{{< fig-quote class="nobox" type="markdown" title="math/rand/v2/rand.go" link="https://pkg.go.dev/math/rand/v2" lang="en" >}}
```go
// A Source is a source of uniformly-distributed
// pseudo-random uint64 values in the range [0, 1<<64).
//
// A Source is not safe for concurrent use by multiple goroutines.
type Source interface {
    Uint64() uint64
}
```
{{< /fig-quote >}}

どえらシンプル！ `Seed()` メソッドがなくなったのは大きいね。
これによって `runtimeSource` やトップレベル関数群が参照する `globalRand` の定義もめっさシンプルになった。

{{< fig-quote class="nobox" type="markdown" title="math/rand/v2/rand.go" link="https://pkg.go.dev/math/rand/v2" lang="en" >}}
```go
// globalRand is the source of random numbers for the top-level
// convenience functions.
var globalRand = &Rand{src: &runtimeSource{}}

//go:linkname runtime_rand runtime.rand
func runtime_rand() uint64

// runtimeSource is a Source that uses the runtime fastrand functions.
type runtimeSource struct{}

func (*runtimeSource) Uint64() uint64 {
    return runtime_rand()
}
```
{{< /fig-quote >}}

うんうん。
シンプルが一番だね。

## ChaCha8 を rand.Source にする

ChaCha8 疑似乱数生成器を疑似乱数の `Source` として明示的に組み込む場合は， [`rand`][`math/rand/v2`]`.NewChaCha8()` 関数を使って生成する。

{{< fig-quote class="nobox" type="markdown" title="math/rand/v2/chacha8.go" link="https://pkg.go.dev/math/rand/v2" lang="en" >}}
```go
import "internal/chacha8rand"

// A ChaCha8 is a ChaCha8-based cryptographically strong
// random number generator.
type ChaCha8 struct {
    state chacha8rand.State
}

// NewChaCha8 returns a new ChaCha8 seeded with the given seed.
func NewChaCha8(seed [32]byte) *ChaCha8 {
    c := new(ChaCha8)
    c.state.Init(seed)
    return c
}

// Seed resets the ChaCha8 to behave the same way as NewChaCha8(seed).
func (c *ChaCha8) Seed(seed [32]byte) {
    c.state.Init(seed)
}

// Uint64 returns a uniformly distributed random uint64 value.
func (c *ChaCha8) Uint64() uint64 {
    for {
        x, ok := c.state.Next()
        if ok {
            return x
        }
        c.state.Refill()
    }
}
```
{{< /fig-quote >}}

ランタイムに組み込まれているものと違って，こちらは排他処理を行っていない。
並行的に安全（concurrency safe）ではないわけだ。
なので平行処理下で [`math/rand/v2`] の ChaCha8 を扱う場合は要注意である。
つか，平行処理下で ChaCha8 疑似乱数生成器を使うならトップレベル関数群を使うべきだろう。

## PCG を rand.Source にする

[`math/rand/v2`] にはもうひとつ疑似乱数生成器が用意されている。
PCG (Permuted Congruential Generator) というそうな。

{{< fig-quote class="nobox" type="markdown" title="math/rand/v2/pcg.go" link="https://pkg.go.dev/math/rand/v2" lang="en" >}}
```go
// A PCG is a PCG generator with 128 bits of internal state.
// A zero PCG is equivalent to NewPCG(0, 0).
type PCG struct {
    hi uint64
    lo uint64
}

// NewPCG returns a new PCG seeded with the given values.
func NewPCG(seed1, seed2 uint64) *PCG {
    return &PCG{seed1, seed2}
}
```
{{< /fig-quote >}}

PCG は線形合同法（LCG）のバリエーションなんだそうで， LCG の統計学上の欠点を改善したものらしい。
2014年に発表された比較的新しいアルゴリズムのようだが，今のところは欠点のようなものは特に指摘されてないとか。
当然ながら暗号技術分野では使えない。

- [PCG, A Family of Better Random Number Generators | PCG, A Better Random Number Generator](http://www.pcg-random.org/index.html)
- [Permuted congruential generator - Wikipedia](https://en.wikipedia.org/wiki/Permuted_congruential_generator)
- [Permuted congruential generator - Wikipedia](https://ja.wikipedia.org/wiki/Permuted_congruential_generator) : 日本語

こちらも並行的に安全ではないのでご注意を。

## これから math/rand を使う理由はないかな

今までは [`math/rand`] で用意されている疑似乱数生成器は暗号技術的にセキュアではない（要件である「予測困難性」を満たさない）ため使いどころを考えなければならなかったが [Go] 1.22 から [`math/rand`], [`math/rand/v2`] ともに ChaCha8 が既定の疑似乱数生成器なったため用途を選ばずカジュアルに使えるようになるだろう。
そうなると，これから疑似乱数生成器を使おうというとききに，わざわざ [`math/rand`] を使う理由はないかな。
[`math/rand`] って無駄に複雑になってる感じだもんなぁ。

ちなみに [`math/rand/v2`] では `Read()` 関数もなくなっているが， [`math/rand`] でも [`rand`][`math/rand`]`.Read()` 関数は Deprecated になってるし，[リリースノート](https://go.dev/doc/go1.22#math_rand_v2 "Go 1.22 Release Notes - The Go Programming Language")を見ると， `Read()` 関数が使いたきゃ [`crypto/rand`] パッケージを使え，みたいなことが書いてあるので，まぁそういうことなんだろう。

## ブックマーク

- [Go 1.22 Release Notes - The Go Programming Language](https://go.dev/doc/go1.22)
  - [Evolving the Go Standard Library with math/rand/v2 - The Go Programming Language](https://go.dev/blog/randv2)
  - [Secure Randomness in Go 1.22 - The Go Programming Language](https://go.dev/blog/chacha8rand)
- [RFC 4086 - Randomness Requirements for Security](https://tools.ietf.org/html/rfc4086) （[日本語訳](https://www.nic.ad.jp/ja/tech/ipa/RFC4086JA.html "セキュリティのための乱雑性についての要件")）
- [Go の疑似乱数生成器は並行的に安全ではないらしい（追記あり）]({{< ref "/golang/pseudo-random-number-generator.md" >}})
- [Go 1.22 の math/rand/v2 を使ってみる](https://zenn.dev/spiegel/articles/20240309-golang-math-rand-v2)

[Go]: https://go.dev/
[`internal/chacha8rand`]: https://pkg.go.dev/internal/chacha8rand "chacha8rand package - internal/chacha8rand - Go Packages"
[`runtime`]: https://pkg.go.dev/runtime "runtime package - runtime - Go Packages"
[`unsafe`]: https://pkg.go.dev/unsafe "unsafe package - unsafe - Go Packages"
[`math/rand`]: https://pkg.go.dev/math/rand "rand package - math/rand - Go Packages"
[`math/rand/v2`]: https://pkg.go.dev/math/rand/v2 "rand package - math/rand/v2 - Go Packages"
[`crypto/rand`]: https://pkg.go.dev/crypto/rand "rand package - crypto/rand - Go Packages"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B0CFL1DK8Q" %}} <!-- Go言語 100Tips -->
{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
{{% review-paapi "4814400535" %}} <!-- 効率的なGo : Efficient Go -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
{{% review-paapi "B00I8ETG96" %}} <!-- 赤ずきんチャチャ -->
