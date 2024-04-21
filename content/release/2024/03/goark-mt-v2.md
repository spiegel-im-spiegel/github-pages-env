+++
title = "github.com/goark/mt/v2 をリリースした"
date =  "2024-03-09T20:45:42+09:00"
description = " Mersenne Twister 疑似乱数生成器を実装した拙作のパッケージを math/rand/v2 に対応することにした"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "programming", "random", "package", "module", "mersenne-twister" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.22 で [`math/rand/v2`] パッケージが登場したため， [Mersenne Twister] 疑似乱数生成器を実装した拙作の [`github.com/goark/mt`][`github.com/goark/mt/v2`] パッケージも [`math/rand/v2`] に対応することにした。

まずはバージョンを v2 に上げて，以下のインポート・パスに変更した。

```go
import "github.com/goark/mt/v2"
```

また `go.mod` も同様に

```text
require github.com/goark/mt/v2 v2.0.1
```

とする。

例として [`github.com/goark/mt/v2`] パッケージと [`math/rand/v2`] パッケージを組み合わせて標準正規分布する値を1万個生成してみる。
こんな感じ。

```go
package main

import (
    "fmt"
    "math"
    "math/rand/v2"

    "github.com/goark/mt/v2/mt19937"
)

func main() {
    rnd := rand.New(mt19937.New(rand.Int64()))
    points := []float64{}
    max := 0.0
    min := 1.0
    sum := 0.0
    for range 10000 {
        point := rnd.NormFloat64()
        points = append(points, point)
        min = math.Min(min, point)
        max = math.Max(max, point)
        sum += point
    }
    n := float64(len(points))
    ave := sum / n
    d2 := 0.0
    for _, p := range points {
        d2 += (p - ave) * (p - ave)
    }
    fmt.Println("           minimum: ", min)
    fmt.Println("           maximum: ", max)
    fmt.Println("           average: ", ave)
    fmt.Println("standard deviation: ", math.Sqrt(d2/n))
}
```

これを実行するとこんな感じになる。

```text
$ go run sample.go
           minimum:  -4.465497509270884
           maximum:  4.409945906326592
           average:  0.010399867661332784
standard deviation:  1.0027323703801945
```

まぁまぁ妥当な感じ？

[`math/rand`] および [`math/rand/v2`] パッケージのトップレベル関数群の疑似乱数生成器が ChaCha8 になったおかげで seed を与えるのがめっちゃ楽になった。
これだけでもありがたい。
ともかくこれで，乱数生成周りの調査と対応は一通り完了かな。

## ブックマーク

- [疑似乱数生成器 goark/mt]({{< ref "/release/mersenne-twister-by-golang.md" >}})
- [Go 1.22 における疑似乱数生成器]({{< ref "/golang/pseudo-random-number-generator-v2.md" >}})
- [Go 1.22 の math/rand/v2 を使ってみる](https://zenn.dev/spiegel/articles/20240309-golang-math-rand-v2)

[Go]: https://go.dev/
[Mersenne Twister]: http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/mt.html "Mersenne Twister: A random number generator (since 1997/10)"
[`math/rand/v2`]: https://pkg.go.dev/math/rand/v2 "rand package - math/rand - pkg.go.dev"
[`math/rand`]: https://pkg.go.dev/math/rand "rand package - math/rand - Go Packages"
[`io`]: https://golang.org/pkg/io/ "io - The Go Programming Language"
[`github.com/goark/mt/v2`]: https://github.com/goark/mt "goark/mt: Mersenne Twister; Pseudo Random Number Generator, Implemented by Golang"
[`crypto/rand`]: https://pkg.go.dev/crypto/rand "rand package - crypto/rand - pkg.go.dev"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B0CFL1DK8Q" %}} <!-- Go言語 100Tips -->
{{% review-paapi "4814400535" %}} <!-- 効率的なGo : Efficient Go -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
{{% review-paapi "B00I8AT1FO" %}} <!-- 数学ガール／乱択アルゴリズム -->
