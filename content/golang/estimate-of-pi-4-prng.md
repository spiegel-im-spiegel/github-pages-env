+++
tags = [
 "golang",
  "math",
  "random",
  "circle-ratio",
]
draft = true
date = "2016-11-20T01:05:27+09:00"
title = "モンテカルロ法による円周率の推定（その4 PRNG）"
description = "description"

[author]
  flickr = "spiegel"
  linkedin = "spiegelimspiegel"
  license = "by-sa"
  instagram = "spiegel_2007"
  flattr = "spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  facebook = "spiegel.im.spiegel"
  url = "http://www.baldanders.info/spiegel/profile/"
  avatar = "/images/avatar.jpg"
  tumblr = "spiegel-im-spiegel"
  github = "spiegel-im-spiegel"
+++

1. [モンテカルロ法による円周率の推定（その1）]({{< relref "golang/estimate-of-pi.md" >}})
2. [モンテカルロ法による円周率の推定（その2 CLI）]({{< relref "golang/estimate-of-pi-2-cli.md" >}})
3. [モンテカルロ法による円周率の推定（その3 Gaussian）]({{< relref "golang/estimate-of-pi-3-gaussian.md" >}})
4. [モンテカルロ法による円周率の推定（その4 PRNG）]({{< relref "golang/estimate-of-pi-4-prng.md" >}}) ← イマココ

「モンテカルロ法による円周率の推定」もひととおり終わったので，今回は擬似乱数生成器（pseudo random number generator）の話。

## math/rand の擬似乱数生成器

[Go 言語]では [`math/rand`] パッケージで擬似乱数を取り扱えることは「[その1]」で紹介した通り。
[`math/rand`] パッケージに実装されている擬似乱数生成器は ALFG (Additive Lagged Fibonacci Generator) のバリエーションらしい。

{{< fig-quote title="[Announce] A rand package for high quality 64bit random numbers (possibly go2)" link="https://groups.google.com/forum/#!topic/golang-nuts/RZ1G3_cxMcM" lang="en" >}}
<q>If I am not mistaken again, the generator is an ALFG (Additive Lagged Fibonacci Generator, thats what Wikipedia calls it). Knuth describes the algorithm in Volume 2 of The art of computer programming in section 3.2.2 (around equation 7). Both Wikipedia and Knuth state the parameter combination 607,273 as possible combination with a period length of 2^(e-1)*(2^607-1) where e is the length of the random number in bits.<br>
I actually found a few references examining its properties and it seems to be a good rng so faar, but there is still seems to be a lack of mathematical background and it is fairly easy to get into trouble by not seeding properly.</q>
{{< /fig-quote >}}

ラグ付フィボナッチ法は[線形合同法（後述）]({{< relref "#lcg" >}})を改善することを目的としたものでフィボナッチ数の生成法を元にしている。

- [Lagged Fibonacci generator - Wikipedia](https://en.wikipedia.org/wiki/Lagged_Fibonacci_generator)

ラグ付フィボナッチ法は以下の式で表される。

{{< fig-quote >}}
\[ \begin{array}{ll}
S_{n} \equiv S_{n-j} * S_{n-k} \pmod{m}, & 0 \lt j \lt k
\end{array} \]
{{< /fig-quote >}}

ラグ付フィボナッチ法は $*$ 演算子によってバリエーションがあるが [`math/rand`] パッケージの実装では加算を使うため **Additive** Lagged Fibonacci Generator ということらしい。
ソースコードで言うとこの部分。

```go
// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (rng *rngSource) Int63() int64 {
	rng.tap--
	if rng.tap < 0 {
		rng.tap += _LEN
	}

	rng.feed--
	if rng.feed < 0 {
		rng.feed += _LEN
	}

	x := (rng.vec[rng.feed] + rng.vec[rng.tap]) & _MASK
	rng.vec[rng.feed] = x
	return x
}
```

## 別の擬似乱数生成器を使う

[`math/rand`] パッケージでは以下の [interface] を持つ別の擬似乱数生成器を使うことができる。

```go
// A Source represents a source of uniformly-distributed
// pseudo-random int64 values in the range [0, 1<<63).
type Source interface {
	Int63() int64
	Seed(seed int64)
}
```

以下にいくつか紹介する。

### 線形合同法{#lcg}

線形合同法（Linear Congruential Generator）は昔の擬似乱数ライブラリでよく使われていたアルゴリズムで，以下の式で表される。

{{< fig-quote >}}
\[
X_{n+1} = (A \times X_{n} + B) \bmod M
\]
{{< /fig-quote >}}

定数 $A$ および $B$ の与え方により幾つかバリエーションがある。

線形合同法は多次元で疎に分布する性質があり，周期も小さいため科学技術シミュレーションなどには向かないと言われている。









## ブックマーク


[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`math/rand`]: https://golang.org/pkg/math/rand/ "rand - The Go Programming Language"
[その1]: {{< relref "golang/estimate-of-pi.md" >}} "モンテカルロ法による円周率の推定（その1）"
[interface]: https://golang.org/doc/effective_go.html#interfaces_and_types "Effective Go - The Go Programming Language"

[complex]: https://golang.org/ref/spec#Complex_numbers "Manipulating complex numbers"
[type]: https://golang.org/ref/spec#Properties_of_types_and_values "Properties of types and values"
[channel]: http://golang.org/ref/spec#Channel_types
[goroutine]: http://golang.org/ref/spec#Go_statements
[gnuplot]: http://www.gnuplot.info/ "gnuplot homepage"
