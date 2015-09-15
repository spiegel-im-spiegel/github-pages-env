# はじめての Go 言語 (on Windows) その2

[前回](http://qiita.com/spiegel-im-spiegel/items/dca0df389df1470bdbfa)の続き。

Go 言語（golang）はネットにあるドキュメントがかなりしっかりしているので言語仕様で悩むところはあまりない。

- [golang.jp - プログラミング言語Goの情報サイト](http://golang.jp/)

なので、体系的に説明するのではなく、具体例で遊びながらこの言語の癖等を調べていくことにする。仕事に使うなら厳密な評価が必要だけど、今のところはそんな予定もないし、まずはテキトーで（笑）

では、これまたみんな大好き素数探索アルゴリズムで遊ぶ。

## 素数の定義

このサイトに訪れるような方には「釈迦に説法」だと思うが、素数（prime number）の定義を以下に示す。

> 1 と自分自身以外に正の約数を持たない 1 より大きい自然数

ここで自然数（natural number）は「ペアノの公理」に従う。（0 が自然数に含まれるかどうかについては色々あるみたいだが、素数の定義には影響がないので、ここでは無視する）

## 素数探索アルゴリズム（その1）

素数の定義を愚直にコードで表すなら以下のようになる。

```go:prime01.go
package main

import (
    "fmt"
    "time"
)

func main() {
    max := 100
    fmt.Printf("%v 以下の素数:", max)

    start := time.Now() //Start
    for n := 2; n <= max; n++ {
        flag := true
        for m := 2; m < n; m++ {
            if (n % m) == 0 { // n が m で割り切れる → 素数ではない
                flag = false
                break
            }
        }
        if flag {
            fmt.Printf(" %v", n)
        }
    }
    goal := time.Now() //Goal
    fmt.Printf("\n%v 経過", goal.Sub(start)) //経過時間を表示
}
```

実行結果はこんな感じ。

```shell
C:>go run prime01.go
100 以下の素数: 2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97
5.0002ms 経過
```

## 素数探索アルゴリズム（その2: エラトステネスの篩）

もう少しだけ効率的に素数を探すアルゴリズムとして「エラトステネスの篩」と呼ばれる方法がある。これは素数の定義から導かれる以下の特徴を最初のアルゴリズムに加えたものである。

1. 2 以上の全ての自然数はひとつ以上の素数の積で構成される（この素数の集合を素因数（prime factor）という）。したがってある数が素数か否かの判定は、その数より小さい素数のみで調べればよい
1. 積の可換則（commutative property あるいは交換法則）により 自然数 $n$ が素数か否か判定する場合は $\sqrt{n}$ 以下の素数で調べればよい  
   （たとえば 35 の素因数は 5 と 7 だが、 $5 \times 7 = 7 \times 5 = 35$ なので、 31 まで回さずとも 3 および 5 ($\le \sqrt{35}$) まで調べれば判定できる）

そして、定義から 2 が素数であることは自明なので（1 と 2 の間に自然数は存在しない）、 2 より大きい 2 の倍数（すなわち偶数）については判定しなくてもよく、対象となる自然数は 3 以上の奇数のみでいいことになる。

```go:prime02.go
package main

import (
    "fmt"
    "math"
    "time"
)

func main() {
    primes := make([]int64, 1)     //素数のリスト
    primes_f := make([]float64, 1) //素数のリスト（浮動小数点へのキャスト）
    primes[0] = 2
    primes_f[0] = 2.0
    var max int64 = 100

    start := time.Now() //Start
    var n int64 = 3
    for n = 3; n < max; n += 2 {
        flag := true
        f := float64(n)
        rf := math.Sqrt(f) // n に対して √n をとる
        for i := 1; i < len(primes); i++ {
            if primes_f[i] > rf { // n に対して √n まで探索すればよい
                break
            } else if (n % primes[i]) == 0 {
                flag = false
                break
            }
        }
        if flag {
            primes = append(primes, n)
            primes_f = append(primes_f, f)
        }
    }
    goal := time.Now() //Goal
    fmt.Printf("%v 以下の素数: %v\n", max, primes)
    fmt.Printf("%v 経過", goal.Sub(start)) //経過時間を表示
}
```

実行結果はこんな感じ。

```shell
C:>go run prime02.go
100 以下の素数: [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97]
0 経過
```

### Slice について

コード中

```go
primes := make([]int64, 1) //素数のリスト
primes[0] = 2
```

の変数 `primes` は slice と呼ばれる可変長のリスト型である。更に組み込み関数 `make()` は slice, map, channel のみ使用可能なメモリ割り当て関数である（slice, map, channel 以外は `new()` を使う）。

slice に要素を追加する場合は `append()` 関数を使えばいいのだが、これが結構クセがある。 `append()` 関数では slice の容量がいっぱいになると新たにメモリを確保してオリジナルの内容をコピーする。つまりポインタが変わってしまうのだ。（メモリの割り当て方のパターンにも注目）

```go:slice.go
package main

import (
	"fmt"
)

func main() {
	a := make([]int, 0)                                         //空の配列を用意
	fmt.Printf("Slice(%02d) : %p : %v (%v)\n", 0, a, a, cap(a)) //配列の表示
	for num := 1; num <= 17; num++ {
		a = append(a, num)                                            //配列要素の追加
		fmt.Printf("Slice(%02d) : %p : %v (%v)\n", num, a, a, cap(a)) //配列の表示
	}
}
```

```shell
C:>go run slice.go
Slice(00) : 0x570560 : [] (0)
Slice(01) : 0xc082002280 : [1] (1)
Slice(02) : 0xc0820022a0 : [1 2] (2)
Slice(03) : 0xc082006620 : [1 2 3] (4)
Slice(04) : 0xc082006620 : [1 2 3 4] (4)
Slice(05) : 0xc082004340 : [1 2 3 4 5] (8)
...
Slice(09) : 0xc08204a000 : [1 2 3 4 5 6 7 8 9] (16)
...
Slice(17) : 0xc082050000 : [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17] (32)
```

ちなみに Go コンパイラは、返値を無視するコーディングに対してエラーを吐く。

```go
primes = append(primes, n)
    ↓
append(primes, n)
```

とすると

```shell
C:>go run prime02.go
# command-line-arguments
.\prime02.go:31: append(primes, n) evaluated but not used
```

とエラーになる。

## 100万個目の素数

上のコードを少し修正して $x$ 個目の素数を調べることにしよう。

```go:prime01b.go
package main

import (
    "flag"
    "fmt"
    "strconv"
    "time"
)

func main() {
    var n, m, max, count int64

    //引数のチェック
    flag.Parse();
    args := flag.Args()
    if len(args) == 0 {
        fmt.Printf("No Argument")
        return
    }
    max, err := strconv.ParseInt(flag.Args()[0], 10, 64)
    if err != nil {
        if enum, ok := err.(*strconv.NumError); ok {
            switch enum.Err {
            case strconv.ErrRange:
                fmt.Printf("Bad Range Error")
            case strconv.ErrSyntax:
                fmt.Printf("Syntax Error")
            }
        }
        return
    }
    if max <= 0 {
        max = 1
    }
    fmt.Printf("%v 個目の素数: ", max)

    //素数探索
    start := time.Now() //Start
    count = 0
    for n = 2; ; n++ {
        flag := true
        for m = 2; m < n; m++ {
            if (n % m) == 0 { // n が m で割り切れる → 素数ではない
                flag = false
                break
            }
        }
        if flag {
            count++;
            if max <= count {
                fmt.Printf("%v", n)
                break;
            }
        }
    }
    goal := time.Now() //Goal
    fmt.Printf("\n%v 経過", goal.Sub(start)) //経過時間を表示
}
```

```shell
C:>go run prime01b.go 100
100 個目の素数: 541
0 経過
C:>go run prime01b.go 10000
10000 個目の素数: 104729
4.3932513s 経過
```

100万個目の素数は有意の時間で見つかりませんでした orz

```go:prime02b.go
package main

import (
    "flag"
    "fmt"
    "math"
    "strconv"
    "time"
)

func main() {
    var n, max, count int64

    //引数のチェック
    flag.Parse()
    args := flag.Args()
    if len(args) == 0 {
        fmt.Printf("No Argument")
        return
    }
    max, err := strconv.ParseInt(flag.Args()[0], 10, 64)
    if err != nil {
        if enum, ok := err.(*strconv.NumError); ok {
            switch enum.Err {
            case strconv.ErrRange:
                fmt.Printf("Bad Range Error")
            case strconv.ErrSyntax:
                fmt.Printf("Syntax Error")
            }
        }
        return
    }
    if max <= 0 {
        max = 1
    }
    fmt.Printf("%v 個目の素数: ", max)

    //素数を格納する slice を生成
    primes := make([]int64, 1, max)     //素数のリスト
    primes_f := make([]float64, 1, max) //素数のリスト（浮動小数点へのキャスト）
    primes[0] = 2
    primes_f[0] = 2.0

    //素数探索
    start := time.Now() //Start
    count = 1
    if max > count {
        for n = 3; ; n += 2 {
            flag := true
            f := float64(n)
            rf := math.Sqrt(f) // n に対して √n をとる
            for i := 1; i < len(primes); i++ {
                if primes_f[i] > rf { // n に対して √n まで探索すればよい
                    break
                } else if (n % primes[i]) == 0 {
                    flag = false
                    break
                }
            }
            if flag {
                count++
                if max > count {
                    primes = append(primes, n)
                    primes_f = append(primes_f, f)
                } else {
                    fmt.Printf("%v", n)
                    break
                }
            }
        }
    } else {
        fmt.Printf("%v", primes[0])
    }
    goal := time.Now()                     //Goal
    fmt.Printf("\n%v 経過", goal.Sub(start)) //経過時間を表示
}
```

```shell
C:>go run prime02b.go 100
100 個目の素数: 541
0 経過
C:>go run prime02b.go 10000
10000 個目の素数: 104729
8.0005ms 経過
C:>go run prime02b.go 1000000
1000000 個目の素数: 15485863
5.0312878s 経過
C:>go run prime02b.go 10000000
10000000 個目の素数: 179424673
2m17.821883s 経過
```

というわけで、100万個目の素数探索に5秒、1000万個目の素数探索に2分半近くかかってしまった。
まぁ、でも、こんなもんか。

## 素数探索アルゴリズム（その3: エラトステネスの篩を並列処理で）

これまでのアルゴリズムは基本的に2重のループで値を順番に付き合わせているだけだったが、この部分を並列処理で行えば速いんじゃないかと思うよね。で、実際に [Go 言語のチュートリアルには並列処理を使った素数探索アルゴリズムが紹介](http://golang.jp/go_tutorial#index12)されている。そこで紹介されているコードを流用させてもらうことにしよう。

並列処理にするには goroutine（「ゴルーチン」と読むらしい）および channel を使う。

```go:prime03.go
package main

import (
    "fmt"
    "time"
)

// 素数候補の数を生成する
// チャネル 'ch' に、3 以降の奇数を送信（2 以外の偶数は素数ではない）
func generate() chan int64 {
    ch := make(chan int64)
    go func() {
        var n int64
        for n = 3; ; n += 2 {
            ch <- n
        }
    }()
    return ch
}

// チャネル 'in' からチャネル 'out' に値をコピー
// その際 'prime' で割り切れる値を取り除く
func filter(in chan int64, prime int64) chan int64 {
    out := make(chan int64)
    go func() {
        for {
            n := <-in
            if (n % prime) != 0 {
                out <- n
            }
        }
    }()
    return out
}

func sieve() chan int64 {
    out := make(chan int64)
    go func() {
        ch := generate()
        for {
            prime := <-ch
            out <- prime
            ch = filter(ch, prime)
        }
    }()
    return out
}

func main() {
    var max, count int64
    max = 100
    count = 0
    fmt.Printf("%v 以下の素数: 2", max) // 2 は自明

    start := time.Now() //Start
    primes := sieve()
    for {
        prime := <-primes
        count++
        if prime <= max {
            fmt.Printf(" %v", prime)
        } else {
            break
        }
    }
    goal := time.Now()                     //Goal
    fmt.Printf("\n%v 経過", goal.Sub(start)) //経過時間を表示
}
```

```shell
C:>go run prime03.go
100 以下の素数: 2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97
5.0003ms 経過
```

あれ？ えーっと、じゃあ、今までと同じように $x$ 個目の素数を調べるコードに書き換えてみる。

```go:prime03b.go
package main

import (
    "flag"
    "fmt"
    "strconv"
    "time"
)

// 素数候補の数を生成する
// チャネル 'ch' に、3 以降の奇数を送信（2 以外の偶数は素数ではない）
func generate() chan int64 {
    ch := make(chan int64)
    go func() {
        var n int64
        for n = 3; ; n += 2 {
            ch <- n
        }
    }()
    return ch
}

// チャネル 'in' からチャネル 'out' に値をコピー
// その際 'prime' で割り切れる値を取り除く
func filter(in chan int64, prime int64) chan int64 {
    out := make(chan int64)
    go func() {
        for {
            n := <-in
            if (n % prime) != 0 {
                out <- n
            }
        }
    }()
    return out
}

func sieve() chan int64 {
    out := make(chan int64)
    go func() {
        ch := generate()
        for {
            prime := <-ch
            out <- prime
            ch = filter(ch, prime)
        }
    }()
    return out
}

func main() {
    var max, count int64

    //引数のチェック
    flag.Parse()
    args := flag.Args()
    if len(args) == 0 {
        fmt.Printf("No Argument")
        return
    }
    max, err := strconv.ParseInt(flag.Args()[0], 10, 64)
    if err != nil {
        if enum, ok := err.(*strconv.NumError); ok {
            switch enum.Err {
            case strconv.ErrRange:
                fmt.Printf("Bad Range Error")
            case strconv.ErrSyntax:
                fmt.Printf("Syntax Error")
            }
        }
        return
    }
    if max <= 0 {
        max = 1
    }
    fmt.Printf("%v 個目の素数: ", max)

    start := time.Now() //Start
    count = 1
    primes := sieve()
    for {
        prime := <-primes
        count++
        if count >= max {
            fmt.Printf(" %v", prime)
            break
        }
    }
    goal := time.Now()                     //Goal
    fmt.Printf("\n%v 経過", goal.Sub(start)) //経過時間を表示
}
```

```shell
C:>go run prime03b.go 100
100 個目の素数:  541
3.0002ms 経過
C:>go run prime03b.go 10000
10000 個目の素数:  104729
18.8920805s 経過
```

こちらも100万個目の素数は有意の時間で見つかりませんでした orz  
いや、一番最初のべた書きより遅いとか、どうなんかいのぅ。

まぁ、遅い理由は分かっていて、ある素数の `filter()` を作るには、それより小さい全ての素数の `filter()` を通らないといけない。「その2」で述べた2番目の特徴（自然数 $n$ が素数か否か判定する場合は $\sqrt{n}$ 以下の素数で調べればよい）が使えないのだ、このやり方では。

というところで、今回はここまで。遊び疲れた。


## ブックマーク

- [Golangでの文字列・数値変換 - 小野マトペの納豆ペペロンチーノ日記](http://matope.hatenablog.com/entry/2014/04/22/101127)
- [Go - コマンドライン引数 - Qiita](http://qiita.com/uokada/items/f0e069a751679dcf616d)
- [Go の並行処理 - Block Rockin’ Codes](http://jxck.hatenablog.com/entry/20130414/1365960707)
