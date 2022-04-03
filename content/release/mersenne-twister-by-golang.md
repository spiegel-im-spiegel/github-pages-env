+++
title = "疑似乱数生成器 goark/mt"
date =  "2019-09-22T17:37:19+09:00"
description = "goark/mt は64bit版 Mersenne Twister を元に pure Go で書き直したものである。"
image = "/images/attention/go-logo_blue.png"
tags = [ "package", "module", "golang", "random" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

[Mersenne Twister] とは[松本眞](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/ "Makoto Matsumoto Home Page")・西村拓士両氏によって1996年に発表された擬似乱数生成アルゴリズムである。
他の擬似乱数生成アルゴリズムと比べて以下の特徴があるそうだ。

{{< fig-quote type="markdown" title="Mersenne Twister とは?" link="http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/what-is-mt.html" >}}
- 従来の様々な生成法の欠点を考慮して設計されています
- 従来にない長周期, 高次元均等分布を持ちます（周期が $2^{19937}-1$ で、623次元超立方体の中に 均等に分布することが証明されています）
- 生成速度がかなり速い
- メモリ効率が良い
{{< /fig-quote >}}

特に2番目が重要で，モンテカルロ法などの科学技術計算に向いている。
Ruby などの一部のプログラミング言語では標準の疑似乱数生成器として組み込まれているらしい。

[goark/mt][`mt`] は [Mersenne Twister] のオリジナルコード（C/C++）を pure [Go] で書き直したものである。

[![check vulns](https://github.com/goark/mt/workflows/vulns/badge.svg)](https://github.com/goark/mt/actions)
[![lint status](https://github.com/goark/mt/workflows/lint/badge.svg)](https://github.com/goark/mt/actions)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/goark/mt/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/goark/mt.svg)](https://github.com/goark/mt/releases/latest)

[goark/mt][`mt`] の特徴は以下の通り。

- [math/rand][`rand`] 互換で [`rand`]`.Rand` のソースとして利用できる
- 並行的に安全（concurrency safe）な構成にできる（[`mt`]`.PRNG` 型を利用した場合）

## mt/mt19937.Source の機能

[`mt`]`/mt19937` パッケージは [64bit版 Mersenne Twister](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/mt64.html) を元に pure [Go] で書き直したものである。

[`mt`]`/mt19937.Source` はそのまま疑似乱数生成器として使える。
たとえば以下のように記述する。

```go
import (
    "fmt"

    "github.com/goark/mt/mt19937"
)

fmt.Println(mt19937.New(19650218).Uint64())
//Output:
//13735441942630277712
```

提供するメソッドは以下の通り。

| メソッド                     | 機能                                              |
| ---------------------------- | ------------------------------------------------- |
| `Source.Seed(int64)`         | 乱数のシードをセットする                          |
| `Source.SeedArray([]uint64)` | 乱数のシード（配列）をセットする                  |
| `Source.Uint64() uint64`     | 乱数として範囲 $[0, 2^{64}-1]$ の整数値を生成する |
| `Source.Int63() int64`       | 乱数として範囲 $[0, 2^{63}-1]$ の整数値を生成する |
| `Source.Real(int) float64`   | 乱数として浮動小数点数値を生成する                |

`Source.Real()` 関数の引数による乱数の出力範囲は以下の通り。

|   引数   | 生成範囲                       |
| :------: | ------------------------------ |
|    1     | 範囲 $[0, 1)$ の浮動小数点数値 |
|    2     | 範囲 $(0, 1)$ の浮動小数点数値 |
| 上記以外 | 範囲 $[0, 1]$ の浮動小数点数値 |

なお [`mt`]`/mt19937.Source` は並行的に安全ではないので goroutine 間でインスタンスを共有できない。

## [math/rand][`rand`] と組み合わせる

[`mt`]`/mt19937.Source` を [`rand`]`.Rand` のソースとして利用するには以下のように記述すればよい。

```go
import (
    "fmt"
    "math/rand"

    "github.com/goark/mt/mt19937"
)

fmt.Println(rand.New(mt19937.New(19650218)).Uint64())
//Output:
//13735441942630277712
```

これで [`rand`]`.Rand` が提供するメソッドはすべて使える。
ただし [`rand`]`.Rand` 自体も並行的に安全ではないので，取り扱いにはやはり注意が必要である。

## mt.PRNG と組み合わせる

[`mt`]`/mt19937.Source` 型を [`mt`]`.PRNG` 型と組み合わせることで並行的に安全な構成にできる。
たとえばこんな感じに記述できる。

```go {hl_lines=[13,19]}
package main

import (
    "sync"
    "time"

    "github.com/goark/mt"
    "github.com/goark/mt/mt19937"
)

func main() {
    wg := sync.WaitGroup{}
    prng := mt.New(mt19937.New(time.Now().UnixNano()))
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for i := 0; i < 10000; i++ {
                prng.Uint64()
            }
        }()
    }
    wg.Wait()
}
```

[`mt`]`.PRNG` 型は [`mt`]`/mt19937.Source` のラッパーになっていて [`rand`]`.Rand` と組み合わせることも可能だが， [`rand`]`.Rand` の内部構造の問題で並行的に安全にならない。
ご注意を。

### io.Reader 互換の疑似乱数生成器

[`mt`]`.PRNG` のインスタンスから [`mt`]`.Reader` 型のインスタンスを生成できる。
こんな感じに記述できる。

```go {hl_lines=[14,19,22]}
package main

import (
    "fmt"
    "sync"
    "time"

    "github.com/goark/mt"
    "github.com/goark/mt/mt19937"
)

func main() {
    wg := sync.WaitGroup{}
    prng := mt.New(mt19937.New(time.Now().UnixNano()))
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            r := prng.NewReader()
            buf := [8]byte{}
            for i := 0; i < 10000; i++ {
                ct , err := r.Read(buf[:])
                if err != nil {
                    return
                }
                fmt.Println(binary.LittleEndian.Uint64(buf[:ct]))
            }
        }()
    }
    wg.Wait()
}
```

[`mt`]`.Reader` 型は [`io`]`.Reader` インタフェースと互換性がある。
また [`mt`]`.Reader` インスタンスも並行的に安全なので goroutine 間で共有可能である。

## ライセンスについて

[goark/mt][`mt`] は MIT ライセンスで提供している。

オリジナルの [Mersenne Twister] コードは GPL または BSD ライセンスで提供されているが MIT ライセンスに書き換えてもいいらしい。

- [Mersenne Twisterの商業利用について](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/MT2002/license.html)

というわけで [goark/mt][`mt`] は MIT ライセンスで提供することにした。
ご利用はお気軽に。

## 【おまけ】 secure.Source について

暗号処理などで使われる [crypto/rand] 標準パッケージを [math/rand][`rand`] 標準パッケージのソースとして使うための [`mt`]`/secure` サブパッケージを作った。
今回追加した `secure.Source` 型は以下のメソッドを持つ。

| メソッド                           | 機能                                                       |
| ---------------------------------- | ---------------------------------------------------------- |
| `Source.Seed(int64)`               | 何もしないダミー・メソッド                                 |
| `Source.SeedArray([]uint64)`       | 何もしないダミー・メソッド                                 |
| `Source.Read([]byte) (int, error)` | オクテット単位で乱数を生成する<br>（[`io`]`.Reader` 互換） |
| `Source.Uint64() uint64`           | 乱数として範囲 $[0, 2^{64}-1]$ の整数値を生成する          |
| `Source.Int63() int64`             | 乱数として範囲 $[0, 2^{63}-1]$ の整数値を生成する          |
| `Source.Real(int) float64`         | 乱数として浮動小数点数値を生成する                         |

これで [`rand`]`.Source` / [`rand`]`.Source64` および [`mt`]`.Source` として [`rand`]`.Rand` 型や [`mt`]`.PRNG` 型の機能を利用することができる。

```go
//go:build run
// +build run

package main

import (
    "fmt"
    "math/rand"

    "github.com/goark/mt"
    "github.com/goark/mt/secure"
)

func main() {
    fmt.Println(rand.New(secure.Source{}).Uint64())
    fmt.Println(mt.New(secure.Source{}).Uint64())
}
```

[Go]: https://go.dev/
[Mersenne Twister]: http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/mt.html "Mersenne Twister: A random number generator (since 1997/10)"
[`rand`]: https://pkg.go.dev/math/rand "rand package - math/rand - pkg.go.dev"
[`io`]: https://golang.org/pkg/io/ "io - The Go Programming Language"
[`mt`]: https://github.com/goark/mt "goark/mt: Mersenne Twister; Pseudo Random Number Generator, Implemented by Golang"
[crypto/rand]: https://pkg.go.dev/crypto/rand "rand package - crypto/rand - pkg.go.dev"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B00I8AT1FO" %}} <!-- 数学ガール／乱択アルゴリズム -->
