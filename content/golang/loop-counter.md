+++
tags = ["golang", "programming", "type", "floating-point"]
description = "浮動小数点数型の変数をループカウンタにするのは止めましょうね。約束だよ！"
date = "2017-01-18T21:45:30+09:00"
update = "2018-11-23T10:29:41+09:00"
title = "1を1億回足して1億にならない場合"
draft = false

[author]
  license = "by-sa"
  name = "Spiegel"
  flickr = "spiegel"
  flattr = ""
  facebook = "spiegel.im.spiegel"
  twitter = "spiegel_2007"
  github = "spiegel-im-spiegel"
  avatar = "/images/avatar.jpg"
  instagram = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
  linkedin = "spiegelimspiegel"
  tumblr = ""

[scripts]
  mathjax = true
  mermaidjs = false
+++

（この記事は [Qiita に投稿した記事](http://qiita.com/spiegel-im-spiegel/items/74a49773413c62721189 "1を1億回足して1億にならない場合 - Qiita")の転載です）

今回は軽めのネタで。

- [C > 浮動小数点型変数はループカウンタとして使用しない - Qiita](http://qiita.com/7of9/items/438a43bf53d60eab59e3)

まぁ浮動小数点数型の仕様を知れば当たり前の話なのだが，面白そうなので「1を1億回足す」ってのを [Go 言語]でも書いてみる。

```go
package main

import "fmt"

func main() {
    var d float32 = 0.0
    for i := 0; i < 100000000; i++ {
        d += 1.0
    }
    fmt.Println(d)
}
```

実行結果は予想通り

```text
$ go run loop1.go
1.6777216e+07
```

となる[^f32]。
念のため `float64` でも試してみよう。

[^f32]: `float32` は32ビットサイズの浮動小数点数型で，符号部1ビット，指数部8ビット，仮数部23ビット，という内訳になっている（仮数部は仮数の小数点以下を表す）。つまり有効桁数が24ビット（10進数で約7桁）しかない。したがって今回のような「1づつ加算する動作を繰り返す」処理では16,777,216（`=0xffffff+1`）以降は「情報落ち」が発生する。ちなみに `float64` は64ビットサイズで仮数部は52ビットあり，10進数にして約15桁の有効桁数になる。


```go
package main

import "fmt"

func main() {
    var d float64 = 0.0
    for i := 0; i < 100000000; i++ {
        d += 1.0
    }
    fmt.Println(d)
}
```

結果は

```text
$ go run loop2.go
1e+08
```

で，ちゃんと1億になる。
[Go 言語]では基本型のサイズが厳密に決まってるので（int, uint, uintptr は除く），浮動小数点数型の計算誤差についてもきちんと見積もれるはずである。

ちなみに

```go
package main

import "fmt"

func main() {
    for d := 0.0; d < 1.0; d += 0.1 {
        fmt.Println(d)
    }
}
```

とすると[^var]

[^var]: “`d := 0.0`” と記述した場合，変数 `d` は `float64` として宣言・初期化される。厳密には定数 “`0.0`” は，いったん「型付けなし」の浮動小数点数として評価された後，変数宣言時に `float64` に暗黙的に変換される。 [Go 言語]におけるこの定数の機能は何かと便利なので覚えておくとよいだろう。

```text
$ go run loop3.go
0
0.1
0.2
0.30000000000000004
0.4
0.5
0.6
0.7
0.7999999999999999
0.8999999999999999
0.9999999999999999
```

ってなことになる[^r] ので浮動小数点数型の変数をループカウンタにするのは止めましょうね。
約束だよ！

[^r]: このような結果になるのは `float32`/`float64` の浮動小数点数型の内部表現が2進数になっているため。たとえば 0.1 を2進数で表すと「0.000110011...」と循環しキリのいい値にならない。このため 0.1 を加算していくと「丸め誤差」が蓄積していくのである。

## math/big パッケージ（追記）

浮動小数点数演算の「情報落ち」や「丸め誤差」等を緩和する方法として [`math/big`] パッケージの [`Float`] 型を使う手がある。
[`Float`] 型では有効桁数を指定できる。
たとえば先程の 0.1 ずつカウントアップさせる処理ならこんなコードになる。

```go
package main

import (
    "fmt"
    "math/big"
)

func main() {
    var x, y, z big.Float //zero initialize
    y.SetFloat64(0.1)     //53bit precision
    x.SetPrec(64)
    y.SetPrec(64)

    for i := 0; i < 10; i++ {
        z.Add(&x, &y)
        x.Set(&z)
        fmt.Printf("x = %v (prec = %d bits)\n", &x, x.Prec())
    }
}
```

このコードでは有効桁数を64ビットにそろえて計算している。
これを実行するとこうなる。

```text
$ go run big2.go
x = 0.10000000000000000555 (prec = 64 bits)
x = 0.2000000000000000111 (prec = 64 bits)
x = 0.30000000000000001665 (prec = 64 bits)
x = 0.4000000000000000222 (prec = 64 bits)
x = 0.50000000000000002776 (prec = 64 bits)
x = 0.6000000000000000333 (prec = 64 bits)
x = 0.70000000000000003886 (prec = 64 bits)
x = 0.8000000000000000444 (prec = 64 bits)
x = 0.90000000000000004996 (prec = 64 bits)
x = 1.0000000000000000555 (prec = 64 bits)
```

もうひとつ。
[`Rat`] 型を使う手もある。
[`Rat`] 型は有理数の内部表現で値を保持するため記述によっては誤差を小さくできる。
たとえばこんな感じ。

```go
package main

import (
    "fmt"
    "math/big"
)

func main() {
    var x, y, z big.Rat //zero initialize
    var a, b big.Int
    a.SetInt64(1)
    b.SetInt64(10)
    y.SetFrac(&a, &b)

    for i := 0; i < 10; i++ {
        z.Add(&x, &y)
        x.Set(&z)
        fmt.Printf("x = %s (%v)\n", x.FloatString(20), &x)
    }
}
```

実行結果は以下の通り。

```text
$ go run big3.go
x = 0.10000000000000000000 (1/10)
x = 0.20000000000000000000 (1/5)
x = 0.30000000000000000000 (3/10)
x = 0.40000000000000000000 (2/5)
x = 0.50000000000000000000 (1/2)
x = 0.60000000000000000000 (3/5)
x = 0.70000000000000000000 (7/10)
x = 0.80000000000000000000 (4/5)
x = 0.90000000000000000000 (9/10)
x = 1.00000000000000000000 (1/1)
```

## Decimal 型（追記）

残念ながら [Go 言語]の標準パッケージには Java で言うところの [BigDecimal](http://docs.oracle.com/javase/8/docs/api/java/math/BigDecimal.html) に相当するものがない。
ただし似たパッケージを提供している人はいるようだ。

- [shopspring/decimal](https://github.com/shopspring/decimal "shopspring/decimal: Arbitrary-precision fixed-point decimal numbers in go")

たとえばこんな感じに記述する。

```go
package main

import (
    "fmt"

    "github.com/shopspring/decimal"
)

func main() {
    x := decimal.NewFromFloat(0)
    y, _ := decimal.NewFromString("0.1")

    for i := 0; i < 10; i++ {
        x = x.Add(y)
        fmt.Printf("x = %s\n", x.StringFixed(20))
    }
}
```

実行結果はこんな感じ。

```text
$ go run big4.go
x = 0.10000000000000000000
x = 0.20000000000000000000
x = 0.30000000000000000000
x = 0.40000000000000000000
x = 0.50000000000000000000
x = 0.60000000000000000000
x = 0.70000000000000000000
x = 0.80000000000000000000
x = 0.90000000000000000000
x = 1.00000000000000000000
```

じゃあ試しに1を1億回足してみよう。

```go
package main

import (
    "fmt"

    "github.com/shopspring/decimal"
)

func main() {
    x := decimal.NewFromFloat(0)
    y, _ := decimal.NewFromString("1.0")

    for i := 0; i < 100000000; i++ {
        x = x.Add(y)
    }
    fmt.Printf("x = %s\n", x.StringFixed(20))
}
```

実行結果はこうなる。

```text
$ go run big4b.go
z = 100000000.00000000000000000000
```

結構な時間がかかった。
でも「情報落ち」もなく綺麗に1億になったようだ。

## ブックマーク

- [浮動小数点数型と誤差](http://www.cc.kyoto-su.ac.jp/~yamada/programming/float.html)
- [情報落ち、桁落ち、丸め誤差、打切り誤差の違い](http://tooljp.com/jyosho/docs/ketaochi-jyohoochi/ketaochi-jyohoochi.html)
- [Better C - Go言語と小数 - Qiita](https://qiita.com/sonatard/items/eac6fb35dcc8e052a293)
- [Go言語の浮動小数点数のお話](https://shogo82148.github.io/blog/2017/10/28/golang-floating-point-number/)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`math/big`]: https://golang.org/pkg/math/big/ "big - The Go Programming Language"
[`Float`]: https://golang.org/pkg/math/big/#Float
[`Rat`]: https://golang.org/pkg/math/big/#Rat

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22"><img src="https://images-fe.ssl-images-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" height="160" alt="プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)"></a></div>
	<dl class="fn">
      <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
      <dd>Alan A.A. Donovan, Brian W. Kernighan</dd>
      <dd>丸善出版</dd>
	  <dd>評価&nbsp;<span class="fa-sm" style="color:goldenrod;">
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
      </span></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed">2018.10.19</abbr> (powered by <a href="https://dadadadone.com/amakuri/" >Amakuri</a>)</p>
</div>
