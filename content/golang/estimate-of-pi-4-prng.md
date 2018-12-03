+++
tags = [
 "golang",
  "math",
  "random",
  "circle-ratio",
]
draft = false
date = "2016-11-20T23:33:55+09:00"
update = "2018-11-13T10:00:52+09:00"
title = "モンテカルロ法による円周率の推定（その4 PRNG）"
description = "math/rand パッケージでは rand.Source interface を持つ別の擬似乱数生成器を使うことができる。"

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
  tumblr = ""
  github = "spiegel-im-spiegel"

[scripts]
  mathjax = true
  mermaidjs = false
+++

1. [モンテカルロ法による円周率の推定（その1）]({{< relref "estimate-of-pi.md" >}})
2. [モンテカルロ法による円周率の推定（その2 CLI）]({{< relref "estimate-of-pi-2-cli.md" >}})
3. [モンテカルロ法による円周率の推定（その3 Gaussian）]({{< relref "estimate-of-pi-3-gaussian.md" >}})
4. [モンテカルロ法による円周率の推定（その4 PRNG）]({{< relref "estimate-of-pi-4-prng.md" >}}) ← イマココ

「モンテカルロ法による円周率の推定」もひととおり終わったので，今回は擬似乱数生成器（pseudo random number generator）の話。

## math/rand の擬似乱数生成器

[Go 言語]では [`math/rand`] パッケージで擬似乱数を取り扱えることは「[その1]」で紹介した通り。
[`math/rand`] パッケージに実装されている擬似乱数生成器はラグ付フィボナッチ法（Lagged Fibonacci Generator）のバリエーションらしい。

{{< fig-quote title="[Announce] A rand package for high quality 64bit random numbers (possibly go2)" link="https://groups.google.com/forum/#!topic/golang-nuts/RZ1G3_cxMcM" lang="en" >}}
<q>If I am not mistaken again, the generator is an ALFG (Additive Lagged Fibonacci Generator, thats what Wikipedia calls it). Knuth describes the algorithm in Volume 2 of The art of computer programming in section 3.2.2 (around equation 7). Both Wikipedia and Knuth state the parameter combination 607,273 as possible combination with a period length of 2^(e-1) * (2^607-1) where e is the length of the random number in bits.<br>
I actually found a few references examining its properties and it seems to be a good rng so faar, but there is still seems to be a lack of mathematical background and it is fairly easy to get into trouble by not seeding properly.</q>
{{< /fig-quote >}}

ラグ付フィボナッチ法は[線形合同法（後述）]({{< relref "#lcg" >}})を改善することを目的としたものでフィボナッチ数の生成法を元にしている。

- [Lagged Fibonacci generator - Wikipedia](https://en.wikipedia.org/wiki/Lagged_Fibonacci_generator)

ラグ付フィボナッチ法は以下の式で表される。

{{< fig-quote >}}
\begin{alignat*}{2}
S_{n} \equiv S_{n-j} * S_{n-k} \pmod{m}, & \; & 0 \lt j \lt k
\end{alignat*}
{{< /fig-quote >}}

ラグ付フィボナッチ法は $ * $ 演算子によってバリエーションがあるが [`math/rand`] パッケージの実装では加算を使うため “**Additive** Lagged Fibonacci Generator” ということらしい。
ソースコードで言うとこの部分かな。

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

## 擬似乱数生成器のバリエーション

[`math/rand`] パッケージでは以下の [interface] を持つ別の擬似乱数生成器を使うことができる。

```go
// A Source represents a source of uniformly-distributed
// pseudo-random int64 values in the range [0, 1<<63).
type Source interface {
    Int63() int64
    Seed(seed int64)
}
```

以下に2つほど紹介する。

### 線形合同法{#lcg}

線形合同法（Linear Congruential Generator）は昔の擬似乱数ライブラリでよく使われていたアルゴリズムで，以下の式で表される。

{{< fig-quote >}}
\[
X_{n+1} = (A \times X_{n} + B) \bmod M
\]
{{< /fig-quote >}}

定数 $A$ および $B$ の与え方により幾つかバリエーションがある。

線形合同法のメリットは実装サイズが小さく計算量も少ない点だろうか。
一方デメリットとしては，多次元で疎に分布する性質があり，周期も小さいため乱数を大量に発生させる必要がある科学技術シミュレーションなどには向かないと言われている。
このためメモリサイズが限られるマイクロ・コントローラのようなものでもない限り線形合同法が使われることはなくなった。

[Go 言語]でわざわざ線形合同法を実装しているパッケージは少ないのだが，たとえば以下のパッケージがある[^dgl]。

[^dgl]: ただし [`github.com/davidminor/gorand`]`/lcg` には不具合があって `Int63()` 関数で負の値を出力する場合がある。とりあえず fork 版の [`github.com/spiegel-im-spiegel/gorand`]`/lcg` で修正している。 Pull request も出したけど，古いコードだし，もうメンテしてないかなぁ。

- [davidminor/gorand: Basic golang implementation of a permuted congruential generator for pseudorandom number generation](https://github.com/davidminor/gorand)

### Mersenne Twister{#mt}

[Mersenne Twister] とは[松本眞](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/ "Makoto Matsumoto Home Page")・西村拓士両氏によって1996年に発表された擬似乱数生成アルゴリズムである。
他の擬似乱数生成アルゴリズムと比べて以下の特徴があるそうだ。
（「[Mersenne Twister とは?](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/what-is-mt.html "What & how is MT?")」より）

- 従来の様々な生成法の欠点を考慮して設計されています。
- 従来にない長周期，高次元均等分布を持ちます。（周期が $2^{19937}-1$ で[^mt1]、623次元超立方体の中に 均等に分布することが証明されています。）
- 生成速度がかなり速い。（処理系にもよりますが、パイプライン処理やキャッシュメモリ のあるシステムでは、Cの標準ライブラリの `rand()` より高速なこと もあります。なお、開発当時には cokus 版は `rand()` より4倍程度高速でしたが、その後 ANSI-C の `rand()` が LCG 法から lagged-fibonacci に 変更されたこともあり、2002年現在 rand と MT の速度差はあまりありません。）
- メモリ効率が良い。（32ビット以上のマシン用に設計された `mt19937.c` は、 624ワードのワーキングメモリを消費するだけです。 1ワードは32ビット長とします。

[^mt1]: $2^{19937}-1$ はメルセンヌ素数で [Mersenne Twister] の名前の由来になっている。 [Mersenne Twister] では周期サイズごとに複数の実装があるが， $2^{19937}-1$ がポピュラーな実装として広く使われているようだ。

ちなみに [Mersenne Twister] の[オリジナル・コードは BSD ライセンスで提供](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/MT2002/license.html "Mersenne Twisterの商用について")されている。

[公式のリポジトリ](https://github.com/MersenneTwister-Lab "MersenneTwister-Lab")には C/C++ による実装のみのようだが， [Go 言語]で実装している人もいるようである。

- [seehuhn/mt19937: An implementation of Takuji Nishimura's and Makoto Matsumoto's Mersenne Twister pseudo random number generator in Go.](https://github.com/seehuhn/mt19937)
- [nutterts/randgen: Pseudo Random Number Generators implementing the Go(lang) math/rand.Source Interface](https://github.com/nutterts/randgen)
- [farces/mt19937_64: Mersenne Twister (int64) for Go](https://github.com/farces/mt19937_64)
- [cuixin/goalg: golang algorithms](https://github.com/cuixin/goalg)
- [cpmech/gosl: Go scientific library](https://github.com/cpmech/gosl) : [SFMT](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/SFMT/ "SIMD-oriented Fast Mersenne Twister") や [TinyMT](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/TINYMT/ "Tiny Mersenne Twister") に対応。オリジナルのコードを [cgo](https://golang.org/cmd/cgo/ "cgo - The Go Programming Language") で結合しているのでクロス環境では注意

## 擬似乱数生成器を組み込む

では，先ほど紹介した擬似乱数生成器を今回のコードに組み込んでみることにしよう。
こんな感じ。

```go
package gencmplx

import (
    "math/rand"

    "github.com/davidminor/gorand/lcg"
    "github.com/seehuhn/mt19937"
)

//RNGs is kind of RNG
type RNGs int

const (
    NULL RNGs = iota
    GO
    LCG
    MT
)

//NewRndSource returns Source of random numbers
func NewRndSource(rng RNGs, seed int64) rand.Source {
    switch rng {
    case LCG:
        return lcg.NewLcg64(uint64(seed))
    case MT:
        mt := mt19937.New()
        mt.Seed(seed)
        return mt
    default:
        return rand.NewSource(seed)
    }
}
```

`gencmplx.NewRndSource()` 関数で `rand.Source` オブジェクトを生成する[^pm]。
これを「[その1]」で作った `gencmplx.New()` 関数に渡せばよい。

[^pm]: [Go 言語]におけるオブジェクトの多態性については「[Go 言語における「オブジェクト」]({{< relref "object-oriented-programming.md" >}})」を参考にどうぞ。

CLI は以下のように調整してみた。

```text
$ go run main.go estmt --help
Estimate of Pi with Monte Carlo method.

Usage:
  pi estmt [flags]

Flags:
  -e, --ecount int   Count of estimate (default 100)

Global Flags:
  -p, --pcount int       Count of points (default 10000)
  -r, --rsource string   Source of RNG (GO/LCG/MT) (default "GO")
```

で，それぞれの擬似乱数生成器で評価を行ってみようと思ったのだが，今回のケースに限ってはあまり違いが出ないようである。

まずは[線形合同法]({{< relref "#lcg" >}})の場合。

```text
$ go run main.go estmt -e 10000 -p 100000 -r LCG > estmt100k-lcg.dat
random number generator: LCG
minimum value: 3.12204
maximum value: 3.16224
average value: 3.14164
standard deviation: 0.00524 (68.3%)
```

{{< fig-img src="./histogram-lcg.png" link="./histogram-lcg.png" width="611" >}}

{{< fig-img src="./qq100k-plot-lcg.png" link="./qq100k-plot-lcg.png" width="611" >}}

次は [Mersenne Twister]({{< relref "#mt" >}}) の場合。

```text
$ go run main.go estmt -e 10000 -p 100000 -r MT > estmt100k-mt.dat
random number generator: MT
minimum value: 3.12380
maximum value: 3.16140
average value: 3.14165
standard deviation: 0.00517 (67.8%)
```

{{< fig-img src="./histogram-mt.png" link="./histogram-mt.png" width="611" >}}

{{< fig-img src="./qq100k-plot-mt.png" link="./qq100k-plot-mt.png" width="611" >}}

もっと多次元だったりすると変わってくるのかなぁ。

## 暗号技術用途の乱数生成器

[Go 言語]では暗号技術用途の乱数として [`crypto/rand`] パッケージが用意されている。
これは [`math/rand`] とは互換性がない。

具体的には，UNIX 系のプラットフォームでは乱数生成に `/dev/urandom` デバイスを参照している[^dr]。
また Windows プラットフォームでは [CryptoAPI 2.0](https://technet.microsoft.com/ja-jp/library/cc734124.aspx) の [`CryptGenRandom`] 関数を使っている[^win]。

[^dr]: `/dev/urandom` はハードウェア・デバイスから十分なエントロピー源が得られない場合は内部で疑似乱数生成器を使用する。このため一時は `/dev/urandom` の脆弱性が疑われたが，現時点では事実上は問題ないとされている。一方で，スマートデバイスのような場合はハードウェア・デバイスからのエントロピー源だけでは外部から推測され易いため，性能のよい疑似乱数生成器を組み合わせるほうが有効になる場合もあるようだ。
[^win]: [`CryptGenRandom`] 関数の内部実装は公開されていないが，やはりキーボードやマウス等のデバイスの挙動をエントロピー源とし， NIST の SP800-90 勧告に従った実装をしているようである。余談だが SP800-90 は乱数生成の一部のアルゴリズムで脆弱性が発見され（これがまた NSA 絡みだったものだから大騒ぎになった），現在は修正版の [SP800-90A Revision 1](http://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-90Ar1.pdf "Recommendation for Random Number Generation Using Deterministic Random Bit Generators")が発行されている。（参考：[擬似乱数生成アルゴリズム Dual_EC_DRBG について](http://www.cryptrec.go.jp/topics/cryptrec_20131106_dual_ec_drbg.html)）

そもそも暗号技術用途の乱数生成器は科学技術シミュレーションやゲームで使う擬似乱数生成器とは要件が異なる。

- [RFC 4086 - Randomness Requirements for Security](https://tools.ietf.org/html/rfc4086) （[IPA による日本語訳](https://www.ipa.go.jp/security/rfc/RFC4086JA.html "セキュリティのための乱雑性についての要件")）
- {{< pdf-file title="NIST Special Publication 800-90A Revision 1: Recommendation for Random Number Generation Using Deterministic Random Bit Generators" link="http://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-90Ar1.pdf" >}}

{{< fig-quote title="RFC 4086 - セキュリティのための乱雑性についての要件" link="https://www.ipa.go.jp/security/rfc/RFC4086JA.html" >}}
<q>従前の観点から統計的にテストされた乱雑性は、セキュリティ用途に要求される予測困難性と同等では<strong>ありません</strong>。<br>
例えば、（CRC Standard Mathematical Tables からのランダムテーブルのような）広く利用可能な一定のシーケンスの利用は、攻撃者に対して非常に弱いです。これを学習したり、推測する攻撃者は、容易に（過去・未来を問わず）そのシーケンス [CRC] に基づいて、すべてのセキュリティを破ることができます。他の例として、AES を 1, 2, 3 ... のような連続した整数を暗号化する一定の鍵と共に使うことは、優れた統計的乱雑性をもつが予測可能な出力を作り出します。他方、6 面のサイコロを連続して転がして、その結果の値を ASCII にエンコードすることは、実質的に予測困難なコンポーネントをもちながらも「統計的に貧弱な出力」を作り出します。それゆえ、「統計的テストの合否は、『何かが予測不可能であるか否か、あるいは、予測可能であるか否か』を表さないこと」に注意してください。</q>
{{< /fig-quote >}}

暗号技術用途の乱数生成器は，暗号分野においては中核技術のひとつであるが，一度に大量の乱数を生成させる必要のある科学技術シミュレーションなどの用途には向かない。
上手く使い分けてほしい。

## ブックマーク

- [Mersenne Twister に関する覚え書き]({{< ref "/remark/2016/03/mersenne-twister.md" >}})
- [PCG, A Family of Better Random Number Generators | PCG, A Better Random Number Generator](http://www.pcg-random.org/)
- [/dev/randomではなく/dev/urandomを使うべき理由? | マイナビニュース](http://news.mynavi.jp/news/2014/03/11/037/)

[Go 言語に関するブックマーク集はこちら]({{< relref "bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`math/rand`]: https://golang.org/pkg/math/rand/ "rand - The Go Programming Language"
[`crypto/rand`]: https://golang.org/pkg/crypto/rand/ "rand - The Go Programming Language"
[その1]: {{< relref "estimate-of-pi.md" >}} "モンテカルロ法による円周率の推定（その1）"
[interface]: https://golang.org/doc/effective_go.html#interfaces_and_types "Effective Go - The Go Programming Language"
[`github.com/davidminor/gorand`]: https://github.com/davidminor/gorand "davidminor/gorand: Basic golang implementation of a permuted congruential generator for pseudorandom number generation"
[`github.com/spiegel-im-spiegel/gorand`]: https://github.com/davidminor/gorand "spiegel-im-spiegel/gorand: Basic golang implementation of a permuted congruential generator for pseudorandom number generation"
[Mersenne Twister]: http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/mt.html "Mersenne Twister: A random number generator (since 1997/10)"
[`CryptGenRandom`]: https://msdn.microsoft.com/ja-jp/library/windows/desktop/aa379942.aspx "CryptGenRandom function"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-20</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4865940391/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4865940391.09._SCTHUMBZZZ_.jpg"  alt="Kotlinスタートブック -新しいAndroidプログラミング"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4839959234/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4839959234.09._SCTHUMBZZZ_.jpg"  alt="Docker実戦活用ガイド"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274218961/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274218961.09._SCTHUMBZZZ_.jpg"  alt="グッド・マス ギークのための数・論理・計算機科学"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-13">2016-07-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
