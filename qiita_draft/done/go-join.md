# Golang の文字列連結はどちらが速い？

Go 言語で文字列の連結を行う際にどうやるのが一番速いか，という話。

## 文字列連結を行う4つの方法

Go 言語で文字列の連結を行う際には概ね以下の4つの方法がある。

1. “`+`” 演算子で連結する
1. [`strings`](http://golang.org/pkg/strings/)`.Join` で連結する
1. [`bytes`](http://golang.org/pkg/bytes/)`.Buffer` で追記する
1. `[]byte` に append する

`string` 型は「不変（immutable）」なので，最初の2つが高コストになるだろうことはすぐに想像がつく。

- [Goでは文字列連結はコストの高い操作 - Qiita](http://qiita.com/ruiu/items/2bb83b29baeae2433a79)
- [Go言語で効率良く文字列を連結する話 #golang - memoメモ](http://atotto.hatenadiary.jp/entry/2013/04/26/202701)

では残りの2つはどうなのかというと

- [Goの文字列結合のパフォーマンス - Qiita](http://qiita.com/ono_matope/items/d5e70d8a9ff2b54d5c37)

によると最後のが一番速いらしい。ほんじゃまぁ，確かめてみるか。

## サンプルコードを用意

以下のコードを使って評価してみる。

```go:join.go
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

テストコードはこんな感じ。

```go:join_test.go
package main

import (
	"os"
	"testing"
)

var list []string

func readFile() {
	file, err := os.Open("CollisionsForHashFunctions.txt") //maybe file path
	if err != nil {
		panic(err)
	}
	defer file.Close()
	list, err = ContentText(file)
	if err != nil {
		panic(err)
	}
}

func BenchmarkWriteBuffer1(b *testing.B) {
	readFile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		content := WriteBuffer1(list)
		_ = content
	}
}

func BenchmarkWriteBuffer1Cap128(b *testing.B) {
	readFile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		content := WriteBuffer1Cap128(list)
		_ = content
	}
}

func BenchmarkWriteBuffer1Cap1K(b *testing.B) {
	readFile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		content := WriteBuffer1Cap1K(list)
		_ = content
	}
}

func BenchmarkWriteBuffer2(b *testing.B) {
	readFile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		content := WriteBuffer2(list)
		_ = content
	}
}

func BenchmarkWriteBuffer2Cap128(b *testing.B) {
	readFile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		content := WriteBuffer2Cap128(list)
		_ = content
	}
}

func BenchmarkWriteBuffer2Cap1K(b *testing.B) {
	readFile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		content := WriteBuffer2Cap1K(list)
		_ = content
	}
}
```

入力テキストだが，小さいファイルではテストにならない気がしたので，大昔に書いたテキスト [`CollisionsForHashFunctions.txt`](http://www.baldanders.info/spiegel/archive/CollisionsForHashFunctions.txt) を使うことにした。サイズは70行，8KB ほど。

## テスト結果

結果は以下のとおり。

```shell
C:>go test -bench WriteBuffer -benchmen
testing: warning: no tests to run
PASS
BenchmarkWriteBuffer1-4           100000     12220 ns/op  28864 B/op  11 allocs/op
BenchmarkWriteBuffer1Cap128-4     100000     11620 ns/op  28800 B/op  10 allocs/op
BenchmarkWriteBuffer1Cap1K-4      200000     11605 ns/op  27904 B/op   7 allocs/op
BenchmarkWriteBuffer2-4           100000     14200 ns/op  25568 B/op   9 allocs/op
BenchmarkWriteBuffer2Cap128-4     100000     15790 ns/op  26800 B/op   8 allocs/op
BenchmarkWriteBuffer2Cap1K-4      200000     10305 ns/op  17520 B/op   5 allocs/op
ok      join    13.260s
```

ありゃりゃ。 [`bytes`](http://golang.org/pkg/bytes/)`.Buffer` を使ったほうが速いみたい（capacity を大きくとれば）。

それなら，ファイルサイズを一気に小さくして9行，1KB にしてやってみる。

```shell
C:>go test -bench WriteBuffer -benchmen
testing: warning: no tests to run
PASS
BenchmarkWriteBuffer1-4          3000000       443 ns/op    448 B/op   3 allocs/op
BenchmarkWriteBuffer1Cap128-4    5000000       345 ns/op    384 B/op   2 allocs/op
BenchmarkWriteBuffer1Cap1K-4     3000000       561 ns/op   1024 B/op   1 allocs/op
BenchmarkWriteBuffer2-4          1000000      1079 ns/op    544 B/op   4 allocs/op
BenchmarkWriteBuffer2Cap128-4    2000000       727 ns/op    560 B/op   3 allocs/op
BenchmarkWriteBuffer2Cap1K-4     2000000       663 ns/op   1136 B/op   2 allocs/op
ok      join    11.987s
```

今度は `[]bytes` の方が速くなった。

まぁでも予想通りかな。データのサイズが大きくなればバッファ操作のほうが有利になるのは分かりやすいっちゃあ分かりやすい。注目すべきは `BenchmarkWriteBuffer1Cap128` と `BenchmarkWriteBuffer1Cap1K` で， capacity を 1KB 取ったほうが若干遅くなっている。この辺のチューニングをどうするか，というところなのだろう。
