+++
date = "2016-03-12T22:12:53+09:00"
update = "2016-03-16T17:35:35+09:00"
description = "というわけで Go 言語で実装することを考えてみる。"
draft = false
tags = ["programming", "golang", "goroutine", "channel", "slice", "benchmark", "random"]
title = "「ズンドコチェック」なるものが流行っているらしい"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"
+++

{{< fig-gen >}}
<blockquote class="twitter-tweet" data-lang="ja"><p lang="ja" dir="ltr">Javaの講義、試験が「自作関数を作り記述しなさい」って問題だったから<br>「ズン」「ドコ」のいずれかをランダムで出力し続けて「ズン」「ズン」「ズン」「ズン」「ドコ」の配列が出たら「キ・ヨ・シ！」って出力した後終了って関数作ったら満点で単位貰ってた</p>&mdash; てくも (@kumiromilk) <a href="https://twitter.com/kumiromilk/status/707437861881180160">2016年3月9日</a></blockquote>
{{< /fig-gen >}}

「習作（study）」としてはなかなか秀逸なアイデアだと思う。
これで満点くれる教官も流石だが（笑） 巷では「ズンドコキヨシ」とか「キヨシチェック」とか「ズンドコチェック」とか呼ばれているらしい。

というわけで

> 「ズン」「ドコ」のいずれかをランダムで出力し続けて「ズン」「ズン」「ズン」「ズン」「ドコ」の配列が出たら「キ・ヨ・シ！」って出力した後終了って関数

を [Go 言語]で実装することを考えてみる。
私はコレを「ズンドコ・コール（zundoko-choir）」と呼ぶことにする。

とはいえ，ズンドコ・コールを実装する事自体はそう難しくない。
たとえばこんな感じ。

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

const (
    zun     = "ズン"
    doko    = "ドコ"
    kiyoshi = "キ・ヨ・シ！"
)

func generate() chan string {
    ch := make(chan string)
    go func() {
        var zundoko = [2]string{zun, doko}
        rand.Seed(time.Now().UnixNano())
        for {
            ch <- zundoko[rand.Intn(2)]
        }
    }()
    return ch
}

func main() {
    zundoko := generate()
    zcount := 0
    for {
        zd := <-zundoko
        fmt.Print(zd)
        if zd == zun {
            zcount++
        } else if zcount >= 4 {
            break
        } else {
            zcount = 0
        }
    }
    fmt.Print(kiyoshi)
}
```

「ズン」および「ドコ」をランダムに生成する部分は [channel] と [goroutine] を使えばいいだろう（`generate()` 関数内の処理）。
擬似乱数は厳密でなくてもいいので安直に [`math/rand`] を使うことにした[^rand]。
さらに「ズン」「ズン」「ズン」「ズン」「ドコ」の配列パターンのチェックだが，「ズン」が4回以上連続で来た直後に「ドコ」が来たら OK としてみた。
まぁ，これがもっとも素朴な実装でパフォーマンスとしてもそれほど遜色ない筈。

[^rand]: [`math/rand`] の乱数生成アルゴリズムの既定は線形合同法らしい。[線形合同法は性能が良くなく](http://www001.upp.so-net.ne.jp/isaku/rand.html "良い乱数・悪い乱数")ゲームや暗号等では使えない。[`math/rand`] の乱数生成アルゴリズムは他のものに入れ替えることができる。たとえば [`seehuhn/mt19937`](https://github.com/seehuhn/mt19937 "seehuhn/mt19937: An implementation of Takuji Nishimura's and Makoto Matsumoto's Mersenne Twister pseudo random number generator in Go.") パッケージが使える。

と，ここまで考えてハタと気づいた。
問題は「自作関数を作り記述しなさい」なんだからメイン関数にロジック書いたらアカンやん！

というわけでまたもゴリゴリとコードを書いてパッケージにしてしまった。
アホだ，私（笑）

- [spiegel-im-spiegel/zundoko: Zundoko-choirs](https://github.com/spiegel-im-spiegel/zundoko)

出力は標準出力に直書きするのではなく [string] の [slice] に `append()` することで実現する。
この出力先を `Choirs` 型として定義した。

```go
//Choirs - zundoko-choirs list
type Choirs struct {
    c []string
}

//Push is append choirs
func (c *Choirs) Push(s string) {
    c.c = append(c.c, s) //maybe panic if c is nil.
}

func (c *Choirs) String() string {
    if c == nil {
        return ""
    }
    content := make([]byte, 0, 128)
    for _, s := range c.c {
        content = append(content, s...)
    }
    return string(content)
}
```

ちなみに文字列の連結は [`strings`].`Join()` 関数は使わず「[文字列連結はどれが速い？]」で紹介した方法を使っている。

これで最初のコードは

```go
func generate() chan string {
    ch := make(chan string)
    go func() {
        var zd = [2]string{Zun, Doko}
        rand.Seed(time.Now().UnixNano())
        for {
            ch <- zd[rand.Intn(2)]
        }
    }()
    return ch
}

//Run zundoko-choirs
func Run() *Choirs {
    zd := generate()
    c := &Choirs{make([]string, 0)}
    zcount := 0
    for {
        s := <-zd
        c.Push(s)
        if s == Zun {
            zcount++
        } else if zcount >= 4 {
            break
        } else {
            zcount = 0
        }
    }
    c.Push(Kiyoshi)
    return c
}
```

と書き換えることができる。
このパッケージを呼び出すメイン側は例えば

```go
package main

import (
    "fmt"

    "github.com/spiegel-im-spiegel/zundoko"
)

func main() {
    c := zundoko.Run()
    fmt.Println(c)
}
```

と書けばいい。

ところで「ズン」「ドコ」の出力は `Choirs` 型で保持られているので，末尾の5要素のパターンを調べる別の方法もあると気づく。
たとえばこんな感じ。

```go
var matchingPattern = []string{Zun, Zun, Zun, Zun, Doko}

func (c *Choirs) match() bool {
    if c == nil {
        return false
    }
    if len(c.c) < 5 {
        return false
    }
    return reflect.DeepEqual(c.c[len(c.c)-5:], matchingPattern)
}
```

この関数を使えば `Run()` 関数は

```go
//Run2 zundoko-choirs (another logic)
func Run2() *Choirs {
    zd := generate()
    c := &Choirs{make([]string, 0)}
    for {
        s := <-zd
        c.Push(s)
        if c.match() {
            break
        }
    }
    c.Push(Kiyoshi)
    return c
}
```

となり随分すっきりする。
ただこれコストが高くつきそうである。
というわけで，これも調べてみた。
まず以下のベンチマークを書く。

```go
package zundoko

import "testing"

func BenchmarkRun1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Run()
    }
}

func BenchmarkRun2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Run2()
    }
}
```

`BenchmarkRun1` が従来のもの， `BenchmarkRun2` が先程の `match()` 関数を使ったバージョンである。
結果は以下のとおり。

```text
$ go test -bench Run -benchmem
testing: warning: no tests to run
PASS
BenchmarkRun1-4    50000     28141 ns/op    1800 B/op     9 allocs/op
BenchmarkRun2-4    30000     40102 ns/op    3912 B/op   115 allocs/op
ok      github.com/spiegel-im-spiegel/zundoko   4.261s
```

乱数の要素が絡むので毎回同じ値ではないが，傾向としてはこんな感じ。
`BenchmarkRun2` のほうが allocation 回数が圧倒的に多いのが分かるだろう。
これがスピードにもダイレクトに反映されている感じである。

今回は「「ズン」が4回以上連続で来た直後に「ドコ」が来たら OK」という単純なロジックだったが，もっと複雑なパターンが要求される場合は工夫が必要かもしれない[^lr]。

[^lr]: たとえば [`container/list`] や [`container/ring`] といったパッケージを使う手がある。

「ズン」と「ドコ」の出現回数を数える関数も作ってみた。

```go
//CountZunDoko returns count of "ZUN" and "DOKO" choirs
func (c *Choirs) CountZunDoko() (int, int) {
    z := 0
    d := 0
    if c == nil {
        return z, d
    }
    for _, s := range c.c {
        switch s {
        case Zun:
            z++
        case Doko:
            d++
        }
    }
    return z, d
}
```

たとえば `generate()` 関数内で使っている擬似乱数パッケージを別のものに換えた時に統計処理で簡単な性能評価ができるかもしれない。
今回はそこまではしなけど（擬似乱数の話はいずれやりたい）。

こうやって手遊びでコードを弄るのは楽しいものである。

## ブックマーク

- [ズンドコキヨシまとめ - Qiita](http://qiita.com/shunsugai@github/items/971a15461de29563bf90)
- [ズンドコキヨシ with Go (n番煎じ) - Qiita](http://qiita.com/shinderuman@github/items/2ff67c2404647d2b7ea6)
- [ズンドコキヨシGolang並列版 - Qiita](http://qiita.com/Rompei/items/bfa03fbc9a94a37703bb) : 「ズン」「ドコ」の生成部分を CPU の数だけ並列処理で行わせてひとつの [channel] に入力するというユニークな実装
- [「ズンドコキヨシ」と擬似乱数 - Qiita](http://qiita.com/spiegel-im-spiegel/items/6a5bc07dbfa46a328e26) : Qiita で擬似乱数について簡単にまとめてみた。整理できたらこちらでも記事にするかも

[Go 言語]: https://golang.org/ "The Go Programming Language"
[channel]: http://golang.org/ref/spec#Channel_types
[goroutine]: http://golang.org/ref/spec#Go_statements
[`math/rand`]: https://golang.org/pkg/math/rand/ "rand - The Go Programming Language"
[string]: http://golang.org/ref/spec#String_types
[slice]: http://golang.org/ref/spec#Slice_types
[`strings`]: https://golang.org/pkg/strings/ "strings - The Go Programming Language"
[文字列連結はどれが速い？]: {{< relref "join-strings.md" >}} "文字列連結はどれが速い？ — プログラミング言語 Go"
[`container/list`]: https://golang.org/pkg/container/list/ "list - The Go Programming Language"
[`container/ring`]: https://golang.org/pkg/container/ring/ "ring - The Go Programming Language"

## 参考図書

<div class="hreview" ><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/41aCueik45L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-15</dd></dl><p class="similar"><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4873117607/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4873117607.09._SCTHUMBZZZ_.jpg"  alt="マイクロサービスアーキテクチャ"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4873117402/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4873117402.09._SCTHUMBZZZ_.jpg"  alt="ハイパフォーマンスPython"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/0134190440/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/0134190440.09._SCTHUMBZZZ_.jpg"  alt="The Go Programming Language (Addison-Wesley Professional Computing Series)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4774166340/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4774166340.09._SCTHUMBZZZ_.jpg"  alt="Vim script テクニックバイブル ~Vim使いの魔法の杖"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。買おうかどうか悩み中。目次があればなぁ。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-03-12">2016-03-12</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
