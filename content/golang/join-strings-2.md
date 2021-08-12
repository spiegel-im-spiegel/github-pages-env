+++
title = "【改訂版】文字列連結はどれが速い？"
date =  "2019-10-13T17:15:36+09:00"
description = "つまり []byte 配列への append() と strings.Builder への追記と strings.Join() は実質的に同じ処理と言える。"
image = "/images/attention/go-logo_blue.png"
tags = ["golang", "string", "join", "benchmark"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今回も小ネタでお送りしております。

2015年に「[文字列連結はどれが速い？]({{< relref "./join-strings.md" >}})」という記事を書いた。
あれから文字連結に関してどう変わったのか。
特に [Go] 1.10 で [`strings`]`.Builder` が追加されているので，その辺も含めて再検証してみる。

今回検証するコードは以下の通り。

```go
package join

import (
    "bytes"
    "strings"
)

var sz8k = 8 * 1024

func JoinStringPlus(ss []string) {
    var str string
    for _, s := range ss {
        str += s + "\n"
    }
}

func JoinStringJoin(ss []string) {
    strings.Join(ss, "\n")
}

func JoinStringByteAppend(ss []string) {
    b := []byte{}
    for _, s := range ss {
        b = append(b, s...)
        b = append(b, "\n"...)
    }
}

func JoinStringByteAppend8K(ss []string) {
    b := make([]byte, 0, sz8k)
    for _, s := range ss {
        b = append(b, s...)
        b = append(b, "\n"...)
    }
}

func JoinStringBuilder(ss []string) {
    b := &strings.Builder{}
    for _, s := range ss {
        b.WriteString(s)
        b.WriteString("\n")
    }
}

func JoinStringBuilder8K(ss []string) {
    b := &strings.Builder{}
    b.Grow(sz8k)
    for _, s := range ss {
        b.WriteString(s)
        b.WriteString("\n")
    }
}

func JoinStringBuffer(ss []string) {
    b := &bytes.Buffer{}
    for _, s := range ss {
        b.WriteString(s)
        b.WriteString("\n")
    }
}

func JoinStringBuffer8K(ss []string) {
    b := bytes.NewBuffer(make([]byte, 0, sz8k))
    for _, s := range ss {
        b.WriteString(s)
        b.WriteString("\n")
    }
}
```

各関数の内容は以下の通り。

| 関数名                   | 内容                                                   |
| ------------------------ | ------------------------------------------------------ |
| `JoinStringPlus`         | `+` 演算子で連結する                                   |
| `JoinStringJoin`         | [`strings`]`.Join` 関数で連結する                      |
| `JoinStringByteAppend`   | `[]byte` 配列に追記する                                |
| `JoinStringByteAppend8K` | `[]byte` 配列に追記する（8KB アロケーション）          |
| `JoinStringBuilder`      | [`strings`]`.Builder` に追記する                       |
| `JoinStringBuilder8K`    | [`strings`]`.Builder` に追記する（8KB アロケーション） |
| `JoinStringBuffer`       | [`bytes`]`.Buffer` に追記する                          |
| `JoinStringBuffer8K`     | [`bytes`]`.Buffer` に追記する（8KB アロケーション）    |

使うメソッドによって出力する型が異なるが（`string` or `[]byte`），今回は無視することにした[^str1]。

[^str1]: `string` 型は不変オブジェクトなので，通常は `[]byte` 型との相互変換の際にメモリ・アロケーションとデータ・コピーが発生する。ちなみに [`strings`]`.Builder` の `String()` メソッドでは [`unsafe`] パッケージを使って無理やりキャスティングしている。

ベンチマーク用のコードは以下の通り。

```go
package join

import (
    "bufio"
    "os"
    "testing"
)

func ReadAll(path string) []string {
    file, err := os.Open(path) //maybe file path
    if err != nil {
        return nil
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    list := []string{}
    for scanner.Scan() {
        list = append(list, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        return nil
    }
    return list
}

var content = ReadAll("CollisionsForHashFunctions.txt")

func BenchmarkJoinStringPlus(b *testing.B) {
    for i := 0; i < b.N; i++ {
        JoinStringPlus(content)
    }
}

func BenchmarkJoinStringJoin(b *testing.B) {
    for i := 0; i < b.N; i++ {
        JoinStringJoin(content)
    }
}

func BenchmarkJoinStringByteAppend(b *testing.B) {
    for i := 0; i < b.N; i++ {
        JoinStringByteAppend(content)
    }
}

func BenchmarkJoinStringByteAppend8K(b *testing.B) {
    for i := 0; i < b.N; i++ {
        JoinStringByteAppend8K(content)
    }
}

func BenchmarkJoinStringBuilder(b *testing.B) {
    for i := 0; i < b.N; i++ {
        JoinStringBuilder(content)
    }
}

func BenchmarkJoinStringBuilder8K(b *testing.B) {
    for i := 0; i < b.N; i++ {
        JoinStringBuilder8K(content)
    }
}

func BenchmarkJoinStringBuffer(b *testing.B) {
    for i := 0; i < b.N; i++ {
        JoinStringBuffer(content)
    }
}

func BenchmarkJoinStringBuffer8K(b *testing.B) {
    for i := 0; i < b.N; i++ {
        JoinStringBuffer8K(content)
    }
}
```

入力テキストは[前回]と同じ [`CollisionsForHashFunctions.txt`](https://baldanders.info/spiegel/archive/CollisionsForHashFunctions.txt) を使用した。
8KB ほどのサイズがある。
つまりコピー先バッファに 8KB の容量があれば追加のアロケーションは発生しないことになる。

では，さっそく実行してみる。

```text
$ go test -bench JoinString -benchmem
goos: linux
goarch: amd64
pkg: join
BenchmarkJoinStringPlus-4                  19484         65256 ns/op      272160 B/op          69 allocs/op
BenchmarkJoinStringJoin-4                 371649          3087 ns/op        8192 B/op           1 allocs/op
BenchmarkJoinStringByteAppend-4           151417          8339 ns/op       35376 B/op          12 allocs/op
BenchmarkJoinStringByteAppend8K-4         502942          2544 ns/op        8192 B/op           1 allocs/op
BenchmarkJoinStringBuilder-4              130408          8434 ns/op       35376 B/op          12 allocs/op
BenchmarkJoinStringBuilder8K-4            418900          2781 ns/op        8192 B/op           1 allocs/op
BenchmarkJoinStringBuffer-4               133052          9545 ns/op       32240 B/op           8 allocs/op
BenchmarkJoinStringBuffer8K-4             351681          3526 ns/op        8192 B/op           1 allocs/op
PASS
ok      join    12.695s
```

結果を表にまとめておこう。

| 関数名                   |  実行時間 |    Alloc サイズ | Alloc 回数 |
| ------------------------ | ---------:| ---------------:| ----------:|
| `JoinStringPlus`         | 65,256 ns | 2,702,160 bytes |         69 |
| `JoinStringJoin`         |  3,087 ns |     8,192 bytes |          1 |
| `JoinStringByteAppend`   |  8,339 ns |    35,376 bytes |         12 |
| `JoinStringByteAppend8K` |  2,544 ns |      8192 bytes |          1 |
| `JoinStringBuilder`      |  8,434 ns |    35,376 bytes |         12 |
| `JoinStringBuilder8K`    |  2,781 ns |     8,192 bytes |          1 |
| `JoinStringBuffer`       |  9,545 ns |    32,240 bytes |          8 |
| `JoinStringBuffer8K`     |  3,526 ns |      8192 bytes |          1 |

まず [`strings`]`.Join()` 関数を使った連結がめっさ速くなってアロケーション回数も1回のみになっていることにビックリした。
ソースコードを見てみたら，やっぱり [`strings`]`.Join()` 関数内部で [`strings`]`.Builder` を使っていた。

```go
// Join concatenates the elements of a to create a single string. The separator string
// sep is placed between elements in the resulting string.
func Join(a []string, sep string) string {
    switch len(a) {
    case 0:
        return ""
    case 1:
        return a[0]
    }
    n := len(sep) * (len(a) - 1)
    for i := 0; i < len(a); i++ {
        n += len(a[i])
    }

    var b Builder
    b.Grow(n)
    b.WriteString(a[0])
    for _, s := range a[1:] {
        b.WriteString(sep)
        b.WriteString(s)
    }
    return b.String()
}
```

ちなみに [`strings`]`.Builder` への追記処理は以下のようになっている。

```go
// WriteString appends the contents of s to b's buffer.
// It returns the length of s and a nil error.
func (b *Builder) WriteString(s string) (int, error) {
    b.copyCheck()
    b.buf = append(b.buf, s...)
    return len(s), nil
}
```

つまり `[]byte` 配列への `append()` と [`strings`]`.Builder` への追記と [`strings`]`.Join()` は実質的に同じ処理で，それぞれの前処理分だけ差が出ているということになる。

今回の検証では

1. やっぱり `+` 演算子による連結はダメダメ[^plus1]
1. よほどの最適化が要求されない限り `[]byte` 配列への `append()` は [`strings`]`.Builder` へ代替可能[^ts1]
1. [`strings`]`.Join()` 関数のパフォーマンスは十分なので気軽に使ってよい
1. 文字列連結に限るなら，もはや [`bytes`]`.Buffer` は有利とは言えない

[^plus1]: リテラル文字列同士の連結はコンパイラが処理するので `+` 演算子で無問題。
[^ts1]: 内部で `append()` 関数を使っていることから分かる通り [`strings`]`.Builder` のインスタンスはコピーして使えないので注意が必要である（インスタンスのポインタを渡せばOK）。当然ながら goroutine-safe ではないので複数の goroutine 間で共有できない。

といったところだろうか。

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[前回]: {{< relref "./join-strings.md" >}} "文字列連結はどれが速い？"
[`strings`]: https://golang.org/pkg/strings/ "strings - The Go Programming Language"
[`bytes`]: https://golang.org/pkg/bytes/ "bytes - The Go Programming Language"
[`unsafe`]: https://golang.org/pkg/unsafe/ "unsafe - The Go Programming Language"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
