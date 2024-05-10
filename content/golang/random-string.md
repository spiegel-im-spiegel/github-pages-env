+++
title = "ランダムな文字列を生成する"
date =  "2020-10-16T14:24:30+09:00"
description = "というわけで math/rand と crypto/rand はトレードオフの関係にある。"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "random", "math" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

疑似乱数に関して面白い記事を見つけたので紹介しつつ，こちらでも試してみる。

- [Go言語でランダムな文字列を作ってみる](https://zenn.dev/najeira/articles/2017-12-28-qiita-5974c6b2c59ecc02fad2)
- [najeira/randstr: Generate random string using crypto/rand and math/rand for Go](https://github.com/najeira/randstr)

お題はこんな感じ：

- 英数字62文字（`a-zA-Z0-9`）からランダムに1文字ずつとって指定の長さの文字列を作成する
- 同じ文字を何度使ってもよい

また，この記事における前提として，以下の interface 型および定数が定義済みであるとする。

```go
package randstr

type Random interface {
    String(int) string
}

const (
    letterBytes    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    letterBytesLen = len(letterBytes)
    letterIdxBits  = 6                    // 6 bits to represent a letter index
    letterIdxMask  = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
    letterIdxMax   = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)
```

じゃあ，いってみようか

## [math/rand][`math/rand`] パッケージによる実装

まずは [`math/rand`] パッケージを使った実装から。
これについては以下のページの議論が参考になる。

- [How to generate a random string of a fixed length in Go? - Stack Overflow](https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go)

これを参照しつつ，過程をすっ飛ばして最終的に以下のコードにしてみた。

```go
package randstr

import (
    "math/rand"
    "unsafe"
)

type MathRandom struct {
    src rand.Source
}

func NewMathRandom(seed int64) Random {
    return &MathRandom{src: rand.NewSource(seed)}
}

func (mr *MathRandom) String(len int) string {
    b := make([]byte, len)
    for i, cache, rest := 0, mr.src.Int63(), letterIdxMax; i < len; rest-- {
        if rest <= 0 {
            cache, rest = mr.src.Int63(), letterIdxMax
        }
        if idx := int(cache & letterIdxMask); idx < letterBytesLen {
            b[i] = letterBytes[idx]
            i++
        }
        cache >>= letterIdxBits
    }
    return *(*string)(unsafe.Pointer(&b))
}
```

`Int63()` メソッドで63ビット分の乱数を取って6ビットずつ切り出して使うイメージ。
ただし参照する `letterBytes` が62文字分なのに対し6ビット整数では `0-63` まであるので，値がはみ出る場合は取得した6ビット乱数を捨てている。
`letterBytes` に適当な記号を2文字足してやればロスは無くなるだろうが，お題から外れるので今回は割愛する。

最後の

```go
return *(*string)(unsafe.Pointer(&b))
```

は少々トリッキーだが `[]byte` インスタンスをコピーなしに `string` 型にキャストするための「{{< ruby "おまじない" >}}お呪い{{< /ruby >}}」だと思えばいい[^cast1]。

[^cast1]: `*(*string)(unsafe.Pointer(&b))` なキャストは [`strings`]`.Builder` の `String()` メソッドで使われている手法の丸パクリだったりする（笑）

文字通り [`unsafe`] な操作なので，乱用して「{{< ruby "のろい" >}}呪い{{< /ruby >}}」にならないようご注意を（笑）

## [crypto/rand][`crypto/rand`] パッケージによる実装

次は [`crypto/rand`] パッケージを使った実装。

[`crypto/rand`] パッケージでは乱数の生成に専用デバイスを使う。

{{< fig-quote type="markdown" title="rand - The Go Programming Language" link="https://golang.org/pkg/crypto/rand/" lang="en" >}}
{{% quote %}}On Linux and FreeBSD, Reader uses `getrandom(2)` if available, `/dev/urandom` otherwise. On OpenBSD, Reader uses `getentropy(2)`. On other Unix-like systems, Reader reads from `/dev/urandom`. On Windows systems, Reader uses the `CryptGenRandom` API. On Wasm, Reader uses the Web Crypto API{{% /quote %}}.
{{< /fig-quote >}}

そのため [`math/rand`] と比べてどうしても処理速度が遅くなる。
したがって [`rand`][`crypto/rand`]`.Read()` 関数の呼び出し回数を抑えるよう実装するのがコツである。

最初に挙げた記事を参考にしつつ，こんな感じでどうだろう。

```go
package randstr

import (
    "crypto/rand"
    "unsafe"
)

type CryptoRandom struct{}

func NewCryptoRandom() Random {
    return &CryptoRandom{}
}

func (cr *CryptoRandom) String(len int) string {
    b := make([]byte, len)
    for i, offset, size, rest := 0, 0, 0, 0; i < len; rest-- {
        //fmt.Println(i, offset, size, rest)
        if rest <= 0 {
            offset = i
            var err error
            size, err = rand.Read(b[offset:])
            if err != nil || size <= 0 {
                return ""
            }
            rest = size
        }
        if idx := int(b[offset+(size-rest)] & letterIdxMask); idx < letterBytesLen {
            b[i] = letterBytes[idx]
            i++
        }
    }
    return *(*string)(unsafe.Pointer(&b))
}
```

インタフェースを合わせるためにエラーハンドリングをサボっているが，ご容赦を。

[`rand`][`crypto/rand`]`.Read()` 関数で乱数をいったんバッファに展開し，その後文字に置き換えていく。
ただし `letterBytes` からはみ出る場合はその値を捨てて，捨てた分をまとめて [`rand`][`crypto/rand`]`.Read()` 関数で再取得する，という動作を繰り返している。

## ベンチマークをとってみる

んじゃあ，これらのコードを使ってベンチマークをとってみよう。
こんなテスト・コードでどうだろう。

```go
package randstr_test

import (
    "randstr"
    "testing"
    "time"
)

const (
    len64   = 64
    len128  = 128
    max512  = len64 * 8
    max1024 = len128 * 8
)

func BenchmarkRandomStringMath64t8(b *testing.B) {
    r := randstr.NewMathRandom(time.Now().UnixNano())
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        for n := 0; n < max512/len64; n++ {
            _ = r.String(len64)
        }
    }
}

func BenchmarkRandomStringMath512(b *testing.B) {
    r := randstr.NewMathRandom(time.Now().UnixNano())
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = r.String(max512)
    }
}

func BenchmarkRandomStringMath128t8(b *testing.B) {
    r := randstr.NewMathRandom(time.Now().UnixNano())
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        for n := 0; n < max1024/len128; n++ {
            _ = r.String(len128)
        }
    }
}

func BenchmarkRandomStringMath1024(b *testing.B) {
    r := randstr.NewMathRandom(time.Now().UnixNano())
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = r.String(max1024)
    }
}

func BenchmarkRandomStringCrypto64t8(b *testing.B) {
    r := randstr.NewCryptoRandom()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        for n := 0; n < max512/len64; n++ {
            _ = r.String(len64)
        }
    }
}

func BenchmarkRandomStringCrypto512(b *testing.B) {
    r := randstr.NewCryptoRandom()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = r.String(max512)
    }
}

func BenchmarkRandomStringCrypto128t8(b *testing.B) {
    r := randstr.NewCryptoRandom()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        for n := 0; n < max1024/len128; n++ {
            _ = r.String(len128)
        }
    }
}

func BenchmarkRandomStringCrypto1024(b *testing.B) {
    r := randstr.NewCryptoRandom()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = r.String(max1024)
    }
}
```

処理内容を表にするとこんな感じ。

| テスト名                           | 使用パッケージ  | 処理内容               |
| ---------------------------------- | --------------- | ---------------------- |
| `BenchmarkRandomStringMath64t8`    | [`math/rand`]   | 64バイト文字列生成×8  |
| `BenchmarkRandomStringMath512`     | [`math/rand`]   | 512バイト文字列生成    |
| `BenchmarkRandomStringMath128t8`   | [`math/rand`]   | 128バイト文字列生成×8 |
| `BenchmarkRandomStringMath1024`    | [`math/rand`]   | 1,024バイト文字列生成  |
| `BenchmarkRandomStringCrypto64t8`  | [`crypto/rand`] | 64バイト文字列生成×8  |
| `BenchmarkRandomStringCrypto512`   | [`crypto/rand`] | 512バイト文字列生成    |
| `BenchmarkRandomStringCrypto128t8` | [`crypto/rand`] | 128バイト文字列生成×8 |
| `BenchmarkRandomStringCrypto1024`  | [`crypto/rand`] | 1,024バイト文字列生成  |

では実際に動かしてみよう。

```text
$ go test -bench RandomString -benchmem
goos: linux
goarch: amd64
pkg: randstr
BenchmarkRandomStringMath64t8-4            641556          1616 ns/op         512 B/op           8 allocs/op
BenchmarkRandomStringMath512-4             899421          1386 ns/op         512 B/op           1 allocs/op
BenchmarkRandomStringMath128t8-4           357760          3093 ns/op        1024 B/op           8 allocs/op
BenchmarkRandomStringMath1024-4            407550          2820 ns/op        1024 B/op           1 allocs/op
BenchmarkRandomStringCrypto64t8-4           81285         14320 ns/op         512 B/op           8 allocs/op
BenchmarkRandomStringCrypto512-4           241180          4827 ns/op         512 B/op           1 allocs/op
BenchmarkRandomStringCrypto128t8-4          64815         18358 ns/op        1024 B/op           8 allocs/op
BenchmarkRandomStringCrypto1024-4          149160          8212 ns/op        1024 B/op           1 allocs/op
PASS
ok      randstr    10.851s
```

これも表にまとめてみる。
処理回数でソートしているのでご注意を。

|  使用パッケージ | 処理内容               |  ns/op | Alloc<br>Size | Alloc<br>Count | Ratio |
| ---------------:|:---------------------- | ------:| -------------:| --------------:| -----:|
|   [`math/rand`] | 64バイト文字列生成×8  |  1,616 |           512 |              8 |   1.0 |
|   [`math/rand`] | 128バイト文字列生成×8 |  3,093 |          1024 |              8 |   1.9 |
| [`crypto/rand`] | 64バイト文字列生成×8  | 14,329 |           512 |              8 |   8.9 |
| [`crypto/rand`] | 128バイト文字列生成×8 | 18,358 |          1024 |              8 |  11.3 |
|                 |                        |        |               |                |       |
|   [`math/rand`] | 512バイト文字列生成    |  1,386 |           512 |              1 |   1.0 |
|   [`math/rand`] | 1,024バイト文字列生成  |  2,820 |          1024 |              1 |   2.0 |
| [`crypto/rand`] | 512バイト文字列生成    |  4,827 |           512 |              1 |   3.5 |
| [`crypto/rand`] | 1,024バイト文字列生成  |  8,212 |          1024 |              1 |   5.9 |

[`math/rand`] パッケージを使った実装は分かりやすい。
文字列が長くなると処理時間が長くなり，アロケーション回数が多いと更に時間がかかる。

[`crypto/rand`] パッケージについては，やはりメソッドの呼び出し回数がボトルネックになっているようだ。
文字列の長さやアロケーション回数の影響を大きく上回っている。

{{< div-box type="markdown" >}}
[`math/rand`] パッケージを使った実装でも `Read()` メソッドでまとめて乱数を取得したほうが速くなるんじゃね？ って思うよね。
私も思った。
ので，実際に試してベンチマークもとったのだが， `Int63()` メソッドで63ビットずつ取るほうが速いのよ，これが。

まぁ，中身を見れば分かるが，[`math/rand`] パッケージの `Read()` メソッドは中で `Int63()` メソッドを呼び出して8ビットずつ切り分けているだけなので，そのオーバヘッド分だけ遅くなってしまうようだ。
残念！

[`math/rand`]: https://golang.org/pkg/math/rand/ "rand - The Go Programming Language"
{{< /div-box >}}

## 科学技術用の疑似乱数生成器と暗号技術用の乱数生成器

科学技術用の疑似乱数生成器と暗号技術用の乱数生成器では求められる要件が異なる。

科学技術用の疑似乱数生成器で最重要なのは「高次元均等分布」な乱数を生成できることで，その次に重要なのは高速に乱数が生成できることである。

一方，暗号技術用の乱数生成器で最重要なのは「予測困難性」である。

たとえば，科学技術用の疑似乱数生成器の多くは，アルゴリズムで乱数を生成するため，起点となる seed が決まれば生成される値が確定してしまう。
これが科学技術用の疑似乱数生成器が暗号技術には向かないとされる理由だ。

しかし，現時点の技術で「予測困難」な乱数を作るためには何らかの方法で外乱要素（またはエントロピー源）を組み込む必要があるため[^dr]，どうしても乱数の生成速度が遅くなる。
大量の乱数を必要とする科学技術計算には向かないわけだ。

[^dr]: `/dev/urandom` はハードウェア・デバイスから十分なエントロピー源が得られない場合は内部で疑似乱数生成器を使用する。このため一時は `/dev/urandom` の脆弱性が疑われたが，現時点では事実上は問題ないとされている。一方で，スマートデバイスのような場合はハードウェア・デバイスからのエントロピー源だけでは外部から推測され易いため，性能のよい疑似乱数生成器を組み合わせるほうが有効になる場合もあるようだ。

というわけで [`math/rand`] と [`crypto/rand`] はトレードオフの関係にある。
上手く使い分けて欲しい。

## ブックマーク

- [モンテカルロ法による円周率の推定（その4 PRNG）]({{< relref "./estimate-of-pi-4-prng.md" >}})
- [Go の疑似乱数生成器は Goroutine-Safe ではないらしい（追記あり）]({{< relref "./pseudo-random-number-generator.md" >}})
- [疑似乱数生成器 spiegel-im-spiegel/mt]({{< ref "/release/mersenne-twister-by-golang.md" >}})

[Go]: https://go.dev/
[`math/rand`]: https://golang.org/pkg/math/rand/ "rand - The Go Programming Language"
[`crypto/rand`]: https://golang.org/pkg/crypto/rand/ "rand - The Go Programming Language"
[`unsafe`]: https://golang.org/pkg/unsafe/ "unsafe - The Go Programming Language"
[`strings`]: https://golang.org/pkg/strings/ "strings - The Go Programming Language"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4908686033" %}} <!-- Goならわかるシステムプログラミング -->
{{% review-paapi "B07VPSXF6N" %}} <!-- 改訂2版 みんなのGo言語 -->
