+++
title = "ポインタが指し示す意味を考える"
date =  "2020-12-26T15:08:29+09:00"
description = "Go では goroutine や interface 型を使った抽象化で並列処理やヒープ管理などの面倒くさい部分をランタイム・モジュールに丸投げする。"
image = "/images/attention/go-logo_blue.png"
tags = [ "engineering", "golang", "pointer" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

以下の記事がちょっと面白かったのでこの記事でも試してみる。

- [Go: Should I Use a Pointer instead of a Copy of my Struct? | by Vincent Blanchon | A Journey With Go | Medium](https://medium.com/a-journey-with-go/go-should-i-use-a-pointer-instead-of-a-copy-of-my-struct-44b43b104963)
- [Goにおけるポインタの使いどころ](https://zenn.dev/uji/articles/f6ab9a06320294146733)

なお，大元の {{% quote lang="en" %}}[Should I Use a Pointer instead of a Copy of my Struct?](https://medium.com/a-journey-with-go/go-should-i-use-a-pointer-instead-of-a-copy-of-my-struct-44b43b104963 "Go: Should I Use a Pointer instead of a Copy of my Struct? | by Vincent Blanchon | A Journey With Go | Medium"){{% /quote %}} が書かれたのは2019年5月で，おそらく [Go] のバージョンも 1.12 あたりだと思うので，その辺を考慮して読むといいだろう。
ちなみに {{% quote lang="en" %}}[A Journey With Go](https://medium.com/a-journey-with-go){{% /quote %}} は [Go] の内部動作について割と詳しく解説されていてオススメの読み物である。

## ヒープのコスト

{{% ruby "トリガー" %}}起点{{% /ruby %}}はこの構造体型。

{{< fig-quote type="markdown" title="Should I Use a Pointer instead of a Copy of my Struct?" link="https://medium.com/a-journey-with-go/-44b43b104963" lang="en" class="nobox" >}}
```go
type S struct {
   a, b, c int64
   d, e, f string
   g, h, i float64
}
```
{{< /fig-quote >}}

そして，この型のインスタンスを生成する（実質的な）構築子を2つ用意する。

{{< fig-quote type="markdown" title="Should I Use a Pointer instead of a Copy of my Struct?" link="https://medium.com/a-journey-with-go/-44b43b104963" lang="en" class="nobox" >}}
```go
func byCopy() S {
   return S{
      a: 1, b: 1, c: 1,
      e: "foo", f: "foo",
      g: 1.0, h: 1.0, i: 1.0,
   }
}

func byPointer() *S {
   return &S{
      a: 1, b: 1, c: 1,
      e: "foo", f: "foo",
      g: 1.0, h: 1.0, i: 1.0,
   }
}
```
{{< /fig-quote >}}

2つの関数はいずれもリテラル表現で指定された内容のインスタンスを返すが， `byCopy()` 関数は値を `byPointer()` 関数はポインタを返すという違いがある。
また `byCopy()` 関数ではインスタンスをスタック上に置くが `byPointer()` 関数ではインスタンスをヒープ上に生成する[^heap1]。

[^heap1]: よく勘違いされるが（というか私も最初の頃は勘違いしていたが）リテラル表現で `&S{ ... }` と記述する場合は，どっかに固定のインスタンスがあって，その固定インスタンスへのポインタを示しているのではなく，暗黙的にヒープ上にインスタンスを生成してリテラルの内容で初期化している。つまり `&S{}` は `new(S)` と等価である。むしろリテラルで初期値を指定できる分だけ `new()` 関数より簡潔で優れている。詳しくは『[プログラミング言語Go](https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』の4.4.1章を参照のこと。これを知ってから組み込みの `new()` 関数はほとんど使わなくなった（笑）

これらの関数の呼び出しコストを計測するベンチマーク・テストは以下の通り。

```go
func BenchmarkMemoryStack(b *testing.B) {
    var s S
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        s = byCopy()
    }
    b.StopTimer()
    _ = fmt.Sprintf("%v", s.a)
}

func BenchmarkMemoryHeap(b *testing.B) {
    var s *S
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        s = byPointer()
    }
    b.StopTimer()
    _ = fmt.Sprintf("%v", s.a)
}
```

[元記事]のコードでは GC (Garbage Collection) の挙動を検証するために色々と仕込んでいるが，今回はコストだけを測ればいいので単純な構成にしてある。

結果はこんな感じ。

```text
$ go test ./... -bench Memory -benchmem
goos: linux
goarch: amd64
pkg: pointer
BenchmarkMemoryStack-4       132169167             9.04 ns/op           0 B/op           0 allocs/op
BenchmarkMemoryHeap-4        15257716            71.6 ns/op          96 B/op           1 allocs/op
PASS
ok      pointer    3.233s
```

まぁ，[元記事]とだいたい同じ結果かな。
見にくいので表にまとめておこう。

| 関数名                 | 実行時間<br>(ナノ秒) | Alloc<br>サイズ | Alloc<br>回数 |
| ---------------------- | --------------------:| ---------------:| -------------:|
| `BenchmarkMemoryStack` |                  9.0 |               0 |             0 |
| `BenchmarkMemoryHeap`  |                 71.6 |              96 |             1 |

言うまでもないが `s = byCopy()` は[代入文](https://golang.org/ref/spec#Assignments "The Go Programming Language Specification - The Go Programming Language")で [Go] では代入時に必ずコピーが発生する。
ただし `s = byCopy()` がインスタンス自体のコピーなのに対し `s = byPointer()` ではポインタ値のみコピーされる。

つまり上の結果はヒープ領域の割当と解放にかかる（GC を含む）時間コスト（の平均）がインスタンスのコピーよりもかなり大きいことを示している。
それでも（GC のオーバーヘッドを含めても）平均で100ナノ秒未満で済んでいるなら十分に優秀だと思うけどね。

[元記事]でも解説されているが [Go] の GC は独立の goroutine で駆動するため，アーキテクチャ[^a1] や使用するコア数の影響を大きく受ける。
GC を含めてシビアな評価が必要なのであれば，その辺の環境を含めて考えるべきだろう。

[^a1]: 最近の goroutine はプリエンプティブ・マルチタスクが可能になったが，アーキテクチャによっては対応していない場合がある。

## コピーのコスト

[元記事]には続きがある。
さきほどの構造体 `S` に対し

{{< fig-quote type="markdown" title="Should I Use a Pointer instead of a Copy of my Struct?" link="https://medium.com/a-journey-with-go/-44b43b104963" lang="en" class="nobox" >}}
```go
func (s S) stack(s1 S) {}

func (s *S) heap(s1 *S) {}
```
{{< /fig-quote >}}

というメソッドを用意してベンチマークテストを以下のように書き直す。

{{< fig-quote type="markdown" title="Should I Use a Pointer instead of a Copy of my Struct?" link="https://medium.com/a-journey-with-go/-44b43b104963" lang="en" class="nobox" >}}
```go
func BenchmarkMemoryStack(b *testing.B) {
    var s S
    var s1 S

    s = byCopy()
    s1 = byCopy()
    for i := 0; i < b.N; i++ {
        for i := 0; i < 1000000; i++ {
            s.stack(s1)
        }
    }
}

func BenchmarkMemoryHeap(b *testing.B) {
    var s *S
    var s1 *S

    s = byPointer()
    s1 = byPointer()
    for i := 0; i < b.N; i++ {
        for i := 0; i < 1000000; i++ {
            s.heap(s1)
        }
    }
}
```
{{< /fig-quote >}}

[Go] の関数引数は値渡し（call by value）なので引数として渡す時点でコピーが発生するが `s.heap(s1)` はポインタ値がコピーされるだけなので，単純に考えれば `s.stack(s1)` のほうがコストが大きいように思える。

しかし，これを実行すると

```text
$ go test ./... -bench Memory -benchmem
goos: linux
goarch: amd64
pkg: pointer
BenchmarkMemoryStack-4           4005        284168 ns/op           0 B/op           0 allocs/op
BenchmarkMemoryHeap-4            4066        283066 ns/op           0 B/op           0 allocs/op
PASS
ok      pointer    2.354s
```

てな感じになる。

んー。
[元記事]とはだいぶ様子が違うねぇ。
両者に有意の差はないようだ。
これも表にまとめておこう。

| 関数名                 | 実行時間<br>(ナノ秒) | Alloc<br>サイズ | Alloc<br>回数 |
| ---------------------- | --------------------:| ---------------:| -------------:|
| `BenchmarkMemoryStack` |              284,168 |               0 |             0 |
| `BenchmarkMemoryHeap`  |              283,066 |               0 |             0 |


確か最近のバージョンアップでインスタンスのコピーがかなり改善されたと聞いているので，その辺が影響しているかもしれない。

## Interface のコスト

ではここで[元記事]にはなかったテストを考えてみよう。

構造体 `S` に以下のメソッドを追加し

```go
func (s S) ValueA() int64 { return s.a }
```

このメソッドを有効にする interface 型

```go
type IS interface {
    ValueA() int64
}
```

と，この型を返す構築子

```go
func byInterface() IS {
    return S{
        a: 1, b: 1, c: 1,
        e: "foo", f: "foo",
        g: 1.0, h: 1.0, i: 1.0,
    }
}
```

を定義する。
この構築子を使ったベンチマークテストも書いておこう。

```go
func BenchmarkMemoryBox(b *testing.B) {
    var s IS
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        s = byInterface()
    }
    b.StopTimer()
    _ = fmt.Sprintf("%v", s.ValueA())
}
```

これを最初のベンチマークテストと比較してみる。
結果はこんな感じ。

```text
$ go test ./... -bench Memory -benchmem
goos: linux
goarch: amd64
pkg: pointer
BenchmarkMemoryStack-4       132085750             9.08 ns/op           0 B/op           0 allocs/op
BenchmarkMemoryHeap-4        15357787            70.0 ns/op          96 B/op           1 allocs/op
BenchmarkMemoryBox-4         14711439            76.0 ns/op          96 B/op           1 allocs/op
PASS
ok      pointer    4.392s
```

表にまとめておこう。

| 関数名                 | 実行時間<br>(ナノ秒) | Alloc<br>サイズ | Alloc<br>回数 |
| ---------------------- | --------------------:| ---------------:| -------------:|
| `BenchmarkMemoryStack` |                  9.1 |               0 |             0 |
| `BenchmarkMemoryHeap`  |                 70.0 |              96 |             1 |
| `BenchmarkMemoryBox`   |                 76.0 |              96 |             1 |

時間コストについて `byCopy()` 関数と `byPointer()` 関数を足したよりちょっと小さい，って感じだろうか。

Interface 型の機能とはボックス化（boxing）である。
ボックス化されたインスタンスは必ずヒープ領域に置かれる。
その意味で `byPointer()` 関数と `byInterface()` 関数がメモリ管理で似たような挙動になるのは納得できるのではないだろうか。

## ヒープを恐れるな

ヒープメモリ操作が高コストなのは汎用 OS 下で動くアプリケーションであれば自明であり，そこに GC のオーバーヘッドが加わるのだから，そりゃあもう「あたり前田のクラッカー」という奴である。

私のようなロートル世代ではヒープ管理は（可能であれば）忌避したい代物だった。
上述したように操作自体が高コストなのに加えて割当と解放を漏れなく矛盾なく記述しきらなければならないのだから面倒くさいことこの上ない。

[Go] では goroutine や interface 型を使った抽象化と引き換えに並列処理やヒープ管理などの面倒くさい部分をランタイム・モジュールに丸投げする。
しかもその「面倒くさい部分」を細かく制御できず，これが [Go] プログラミングにおける重要なトレードオフとなっているのである。

もしヒープ管理をテッペンから見下ろして完全掌握したいと考えるのなら GC は邪魔なだけだし，そもそも [Go] で書くインセンティブがない。
それこそ近ごろ流行りの Rust とかで書くべきだろう。

今回の記事のような話を知識として知っておくのはいいことだと思うが，設計上の重要なポイントではない（むしろチューニングの話だ）。
ポインタを「概念」で捉えることができれば「ポインタが指し示す意味」について深く考察できるようになる。
それこそが本来の「プログラム設計」というやつである。

## ブックマーク

- [Design Philosophy On Data And Semantics](https://www.ardanlabs.com/blog/2017/06/design-philosophy-on-data-and-semantics.html)
- [Big Sky :: Go のポインタの躓きやすい点](https://mattn.kaoriya.net/software/lang/go/20190516095124.htm)

[Go]: https://golang.org/ "The Go Programming Language"
[元記事]: https://medium.com/a-journey-with-go/-44b43b104963 "Go: Should I Use a Pointer instead of a Copy of my Struct? | by Vincent Blanchon | A Journey With Go | Medium"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
{{% review-paapi "B00HY73M16" %}} <!-- Be mine! -->
