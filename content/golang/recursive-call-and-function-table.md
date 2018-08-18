+++
date = "2016-01-18T22:59:07+09:00"
update = "2018-05-28T20:08:28+09:00"
description = "今回は再帰呼び出しの話。"
draft = false
tags = ["golang", "recursion"]
title = "再帰呼び出しと関数テーブル"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

今回は再帰呼び出しの話。

再帰呼び出しのサンプルとして典型的なのは[フィボナッチ数]かな。
[フィボナッチ数]の定義を愚直にコードにするとこんな感じになる。

```go
package main

import "fmt"

func fibonacciNumber(n int) int {
    switch n {
    case 0:
        return 0
    case 1:
        return 1
    default:
        return fibonacciNumber(n-2) + fibonacciNumber(n-1)
    }
}

func main() {
    fmt.Println(fibonacciNumber(40))
}
```

一般に手続き型言語は再帰呼び出しに弱いと言われている（関数型のほうが有利）。
特に [Go 言語]の場合は [goroutine] に最適化を割り振っている関係で，関数呼び出しやその戻り[^r] のパフォーマンスが冷遇されているようだ。
したがって，再帰呼び出し部分のパフォーマンスを改善したければ，なるべく呼び出し回数を減らすようにするのがコツである。
たとえば上の[フィボナッチ数]の場合は

[^r]: 再帰呼び出しが「末尾呼び出し（tail call）」の場合は，戻りの最適化でパフォーマンスの向上が期待できるが， [Go 言語]のコンパイラはこの辺の最適化も行っていないらしい。

```go
package main

import "fmt"

var fibonacciNumbers = make(map[int]int)

func fibonacciNumber(n int) int {
    switch n {
    case 0:
        return 0
    case 1:
        return 1
    default:
        if fn, ok := fibonacciNumbers[n]; ok {
            return fn
        }
        fn := fibonacciNumber(n-2) + fibonacciNumber(n-1)
        fibonacciNumbers[n] = fn
        return fn
    }
}

func main() {
    fmt.Println(fibonacciNumber(40))
}
```

という感じに計算結果を保持っておくことでかなり改善する。

ところで，再帰呼び出しで怖いのが無限呼び出しに陥るパターンである。
[Go 言語]では関数値（function value）を介す場合であれば再帰呼び出しを禁止する。

たとえば先ほどのコードを以下のように書き換えてみる。

```go
package main

import "fmt"

var fibonacciNumbers = make(map[int]int)

func fibonacciNumber(n int) int {
    switch n {
    case 0:
        return 0
    case 1:
        return 1
    default:
        if fn, ok := fibonacciNumbers[n]; ok {
            return fn
        }
        fn := fib(n-2) + fib(n-1)
        fibonacciNumbers[n] = fn
        return fn
    }
}

type ff func(int) int

var fib ff = fibonacciNumber

func main() {
    fmt.Println(fib(40))
}
```

`fibonacciNumber()` を `fib()` で別名定義しているだけだが，これを実行しようとすると以下のコンパイルエラーになる。

```
prog.go:25: initialization loop:
    prog.go:25 fib refers to
    prog.go:7 fibonacciNumber refers to
    prog.go:25 fib
```

関数値 `fib` の部分でエラーになっている点に注目してほしい。

本当に「ついうっかり」再帰呼び出しになってしまう場合はエラーではじいてもらってありがたいのだが，そうでない場合もある。
あまり例示がうまくなくて申し訳ないのだが，以下の簡単なステート・マシンを考えてみる。

```go
package main

import "fmt"

func f0(evt int) int {
    fmt.Println("processing f0")
    return 1
}

func f1(evt int) int {
    fmt.Println("processing f1")
    return 2
}

func f2(evt int) int {
    fmt.Println("processing f2")
    return 3
}

func f3(evt int) int {
    fmt.Println("processing f3")
    return 0
}

type fn func(int) int

var fs = []fn{
    f0,
    f1,
    f2,
    f3,
}

func StateMachin(stat, evt int) int {
    return fs[stat](evt)
}

func main() {
    s := 0
    for e := 0; e < 10; e++ {
        s = StateMachin(s, e)
        if s == 0 {
            break
        }
    }
    fmt.Println("end")
}
```

`fs` が関数テーブルになっていて，状態 `s` とイベント `e` に対する処理 `fs[s](e)` を呼び出して処理後の状態を返してもらう。
一応 `StateMachine()` 関数で詳細を隠蔽しているつもりである[^b]。

[^b]: 実際には別パッケージにしてちゃんとクラス設計すべきだろうけど色々端折っている。ゴメンペコン。

ここで `StateMachine(3, evt)` の処理に続けて `StateMachine(1, evt)` の処理がしたくなり

```go
func f3(evt int) int {
    fmt.Println("processing f3")
    return StateMachin(1, evt)
}
```

などと書いたらどうなるか。
もちろんこれもコンパイルエラーになる。

```
prog.go:32: initialization loop:
    prog.go:32 fs refers to
    prog.go:20 f3 refers to
    prog.go:34 StateMachin refers to
    prog.go:32 fs
```

しかし実際には `f3` を無限に呼び出しているわけではないので，このコンパイルエラーでは困ってしまう[^a]。
このエラーを回避するには関数テーブル `fs` を介さなければよい。
たとえば

```
var fs = []fn{
    f0,
    f1,
    f2,
    //f3,
}

func StateMachin(stat, evt int) int {
    if stat == 3 {
        return f3(evt)
    }
    return fs[stat](evt)
}
```

とすれば，めでたく `f1` → `f2` → `f3` → `f1` とエンドレスにつながる。
ちなみにここでうっかり

```go
func f3(evt int) int {
    fmt.Println("processing f3")
    return StateMachin(3, evt)
}
```

などと書くと，コンパイルエラーにもならず無限呼び出しが発生する。

再帰呼び出しは計画的に。

[^a]: このコードに限れば `StateMachine(1, evt)` ではなく `f1(evt)` を呼び出せば済む話なのでこれは言いがかりであるが，「話の都合」ということで軽く流していただけるとありがたい。

## ブックマーク

- [Goで再帰使うと遅くなりますがそれが何だ - YAMAGUCHI::weblog](http://ymotongpoo.hatenablog.com/entry/2015/02/23/165341)
- [.\hoge.go:7: initialization loop - Qiita](http://qiita.com/zetamatta/items/cc0f29441b16d63472ed)
- [Fibonacci exercise (go tour) - Qiita](http://qiita.com/penguin_dream/items/f1bdeb4c621a3e8d8990)
- [Golangでゴルーチンにより再帰関数を並列処理](https://qiita.com/hiroykam/items/fdbb68ea21e5c67b8225) : 再帰処理を goroutine で並行処理するという業なコード。 goroutine が大量発生して面白い

[Go 言語に関するブックマーク集はこちら]({{< relref "bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[フィボナッチ数]: https://en.wikipedia.org/wiki/Fibonacci_number "Fibonacci number - Wikipedia, the free encyclopedia"
[goroutine]: http://golang.org/ref/spec#Go_statements
