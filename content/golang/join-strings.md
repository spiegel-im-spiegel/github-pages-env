+++
date = "2015-09-21T23:08:07+09:00"
update = "2015-12-20T17:27:35+09:00"
description = "Go 言語で文字列の連結を行う際にどうやるのが一番速いか，という話。"
draft = false
tags = ["golang", "string", "join", "benchmark"]
title = "文字列連結はどれが速い？"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

（初出： [Golang の文字列連結はどちらが速い？ - Qiita](http://qiita.com/spiegel-im-spiegel/items/16ab7dabbd0749281227)）

[前回]につづき [string] の話題。
[Go 言語]で文字列の連結を行う際にどうやるのが一番速いか，という話。

## 文字列連結を行う4つの方法

[Go 言語]で文字列の連結を行う際には概ね以下の4つの方法がある。

1. “`+`” 演算子で連結する
1. [`strings`](http://golang.org/pkg/strings/)`.Join` で連結する
1. [`bytes`](http://golang.org/pkg/bytes/)`.Buffer` に追記する
1. `[]byte` に `append` する

[string] は「不変（immutable）」なので，最初の2つが高コストになるだろうことはすぐに想像がつく。

- [Goでは文字列連結はコストの高い操作 - Qiita](http://qiita.com/ruiu/items/2bb83b29baeae2433a79)
- [Go言語で効率良く文字列を連結する話 #golang - memoメモ](http://atotto.hatenadiary.jp/entry/2013/04/26/202701)

では残りの2つはどうなのかというと

- [Goの文字列結合のパフォーマンス - Qiita](http://qiita.com/ono_matope/items/d5e70d8a9ff2b54d5c37)

によると最後のが一番速いらしい。ほんじゃまぁ，確かめてみるか。

## サンプルコードを用意

以下のコード `join.go` を使って評価してみる。

```go
package main

import (
    "bufio"
    "bytes"
    "io"
)

//Read content (text data) from buffer
func ContentText(inStream io.Reader) ([]string, error) {
    scanner := bufio.NewScanner(inStream)
    list := make([]string, 0)
    for scanner.Scan() {
        list = append(list, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        return nil, err
    }
    return list, nil
}

//Write content (text data) to buffer
func WriteBuffer1(lines []string) []byte {
    //write to byte buffer
    content := make([]byte, 0)
    recode := "\r\n"
    for _, line := range lines {
        content = append(content, line...)
        content = append(content, recode...)
    }
    return content
}

//Write content (text data) to buffer
func WriteBuffer1Cap128(lines []string) []byte {
    //write to byte buffer
    content := make([]byte, 0, 128) //128 bytes capacity
    recode := "\r\n"
    for _, line := range lines {
        content = append(content, line...)
        content = append(content, recode...)
    }
    return content
}

//Write content (text data) to buffer
func WriteBuffer1Cap1K(lines []string) []byte {
    //write to byte buffer
    content := make([]byte, 0, 1024) //1K bytes capacity
    recode := "\r\n"
    for _, line := range lines {
        content = append(content, line...)
        content = append(content, recode...)
    }
    return content
}

//Write content (text data) to buffer (buffered I/O)
func WriteBuffer2(lines []string) []byte {
    //write to byte buffer
    content := bytes.NewBuffer(make([]byte, 0))
    recode := "\r\n"
    for _, line := range lines {
        content.WriteString(line)
        content.WriteString(recode)
    }
    return content.Bytes()
}

//Write content (text data) to buffer (buffered I/O)
func WriteBuffer2Cap128(lines []string) []byte {
    //write to byte buffer
    content := bytes.NewBuffer(make([]byte, 0, 128)) //128 bytes capacity
    recode := "\r\n"
    for _, line := range lines {
        content.WriteString(line)
        content.WriteString(recode)
    }
    return content.Bytes()
}

//Write content (text data) to buffer (buffered I/O)
func WriteBuffer2Cap1K(lines []string) []byte {
    //write to byte buffer
    content := bytes.NewBuffer(make([]byte, 0, 1024)) //1K bytes capacity
    recode := "\r\n"
    for _, line := range lines {
        content.WriteString(line)
        content.WriteString(recode)
    }
    return content.Bytes()
}
```

テストコード `join_test.go` はこんな感じ。

```go
package main

import (
    "os"
    "testing"
)

func readFile() []string {
    file, err := os.Open("CollisionsForHashFunctions.txt") //maybe file path
    if err != nil {
        panic(err)
    }
    defer file.Close()
    list, err := ContentText(file)
    if err != nil {
        panic(err)
    }
    return list
}

func BenchmarkWriteBuffer1(b *testing.B) {
    list := readFile()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        content := WriteBuffer1(list)
        _ = string(content)
    }
}

func BenchmarkWriteBuffer1Cap128(b *testing.B) {
    list := readFile()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        content := WriteBuffer1Cap128(list)
        _ = string(content)
    }
}

func BenchmarkWriteBuffer1Cap1K(b *testing.B) {
    list := readFile()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        content := WriteBuffer1Cap1K(list)
        _ = string(content)
    }
}

func BenchmarkWriteBuffer2(b *testing.B) {
    list := readFile()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        content := WriteBuffer2(list)
        _ = string(content)
    }
}

func BenchmarkWriteBuffer2Cap128(b *testing.B) {
    list := readFile()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        content := WriteBuffer2Cap128(list)
        _ = string(content)
    }
}

func BenchmarkWriteBuffer2Cap1K(b *testing.B) {
    list := readFile()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        content := WriteBuffer2Cap1K(list)
        _ = string(content)
    }
}
```

[Go 言語]のテストについては[以前紹介した]({{< ref "golang/testing.md" >}})が，同じ要領で `Benchmark` から始まる名前の関数を作るとベンチマーク用のコードとして認識される。
引数には `b *testing.B` を指定する。

ベンチマークの内訳は以下のとおり。

| ベンチマーク名  | 処理内容                     |
|:----------------|:-----------------------------|
| `BenchmarkWriteBuffer1`       | `[]byte` に `append` する |
| `BenchmarkWriteBuffer1Cap128` | `[]byte` に `append` する（ capacity 128B） |
| `BenchmarkWriteBuffer1Cap1K`  | `[]byte` に `append` する（ capacity 1KB） |
| `BenchmarkWriteBuffer2`       | [`bytes`]`.Buffer` に追記する |
| `BenchmarkWriteBuffer2Cap128` | [`bytes`]`.Buffer` に追記する（ capacity 128B） |
| `BenchmarkWriteBuffer2Cap1K`  | [`bytes`]`.Buffer` に追記する（ capacity 1KB） |

入力テキストだが，小さいファイルではテストにならない気がしたので，大昔に書いたテキスト [`CollisionsForHashFunctions.txt`](http://www.baldanders.info/spiegel/archive/CollisionsForHashFunctions.txt) を使うことにした。
サイズは70行，7KB ほど。

## テスト結果

結果は以下のとおり。

```
C:>go test -bench WriteBuffer -benchmem
testing: warning: no tests to run
PASS
BenchmarkWriteBuffer1-8           100000         17831 ns/op       37056 B/op     12 allocs/op
BenchmarkWriteBuffer1Cap128-8     100000         20321 ns/op       36992 B/op     11 allocs/op
BenchmarkWriteBuffer1Cap1K-8      100000         19301 ns/op       36096 B/op      8 allocs/op
BenchmarkWriteBuffer2-8           100000         17300 ns/op       33760 B/op     10 allocs/op
BenchmarkWriteBuffer2Cap128-8     100000         19451 ns/op       34992 B/op      9 allocs/op
BenchmarkWriteBuffer2Cap1K-8      100000         15490 ns/op       25712 B/op      6 allocs/op
ok      join    12.659s
```

ありゃりゃ。 [`bytes`](http://golang.org/pkg/bytes/)`.Buffer` を使ったほうが速いみたい（capacity を大きくとれば）。

それなら，入力テキストを切り詰めて10行，0.3KB にしてやってみる。

```
C:>go test -bench WriteBuffer -benchmem
testing: warning: no tests to run
PASS
BenchmarkWriteBuffer1-8          2000000           859 ns/op        1312 B/op      5 allocs/op
BenchmarkWriteBuffer1Cap128-8    2000000           707 ns/op        1248 B/op      4 allocs/op
BenchmarkWriteBuffer1Cap1K-8     2000000           796 ns/op        1376 B/op      2 allocs/op
BenchmarkWriteBuffer2-8          1000000          1686 ns/op        1600 B/op      6 allocs/op
BenchmarkWriteBuffer2Cap128-8    1000000          1411 ns/op        1680 B/op      5 allocs/op
BenchmarkWriteBuffer2Cap1K-8     2000000           980 ns/op        1488 B/op      3 allocs/op
ok      join    13.589s
```


今度は `[]byte` の方が速くなった。

まぁでも予想通りかな。
データのサイズが大きくなればバッファ操作のほうが有利になるのは分かりやすいっちゃあ分かりやすい。

注目すべきは `BenchmarkWriteBuffer1Cap128` と `BenchmarkWriteBuffer1Cap1K` で， capacity を 1KB 取ったほうが若干遅くなっている。この辺のチューニングをどうするか，というところなのだろう（実はこれ，環境によって微妙に順位が変わるんだよなぁ）。

## ブックマーク

- [Go でベンチマーク - Block Rockin’ Codes](http://jxck.hatenablog.com/entry/20131123/1385189088)
- [go言語でベンチマーク - Qiita](http://qiita.com/Mulyu/items/ed585f2777496f29a725)

[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[前回]: {{< ref "golang/string-and-rune.md" >}} "String と Rune"
[string]: http://golang.org/ref/spec#String_types
[`bytes`]: http://golang.org/pkg/bytes/
