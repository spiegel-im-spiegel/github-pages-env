+++
title = "strings.EqualFold 関数を使え"
date =  "2019-12-15T18:01:19+09:00"
description = "大文字小文字を無視した文字列比較では素直に strings.EqualFold() 関数を使いましょう。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "string", "unicode", "benchmark" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GolangCI] が吐く[レビュー結果]({{< relref "./code-review-with-golangci.md" >}} "GolangCI でコード・レビューを自動化する")を基にチマチマとコードを直していたのだが，その中で

```text
SA6005: should use strings.EqualFold(a, b) instead of strings.ToLower(a) == strings.ToLower(b)

if strings.ToLower(left) == strings.ToLower(right) {
```

という指摘があった。
いや，もの知らずでゴメンペコン。

[`strings`].`EqualFold()` 関数ってなんじゃら？ と思ってソースコードを見たら

```go
// EqualFold reports whether s and t, interpreted as UTF-8 strings,
// are equal under Unicode case-folding.
func EqualFold(s, t string) bool {
    ...
}
```

と書かれている。

ふむふむ。
では試してみよう。
こんな感じのコードを書いて

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	lefts := []string{"go", "ｇｏ"}
	rights := []string{"Go", "GO", "go", "Ｇｏ", "ＧＯ", "ｇｏ"}

	for _, left := range lefts {
		for _, right := range rights {
			fmt.Printf("%s == %s : %v\n", left, right, strings.EqualFold(left, right))
		}
	}
}
```

実行してみると

```text
$ go run sample1.go 
go == Go : true
go == GO : true
go == go : true
go == Ｇｏ : false
go == ＧＯ : false
go == ｇｏ : false
ｇｏ == Go : false
ｇｏ == GO : false
ｇｏ == go : false
ｇｏ == Ｇｏ : true
ｇｏ == ＧＯ : true
ｇｏ == ｇｏ : true
```

ってな感じになった。
全角と半角は区別してくれるらしい。
Unicode の文字種をきちんと判別しているということだ。

ちなみに [`strings`].`ToLower()` 関数を使って

```go {hl_lines=[14]}
package main

import (
	"fmt"
	"strings"
)

func main() {
	lefts := []string{"go", "ｇｏ"}
	rights := []string{"Go", "GO", "go", "Ｇｏ", "ＧＯ", "ｇｏ"}

	for _, left := range lefts {
		for _, right := range rights {
			fmt.Printf("%s == %s : %v\n", left, right, (left == strings.ToLower(right)))
		}
	}
}
```

とやっても同じ結果になる。

[`strings`].`EqualFold()` 関数と [`strings`].`ToLower()` 関数でどっちが速いかなんてのは考えるまでもないのだが，いちおう試しておこう。
こんな感じのコードでいいかな。

```go {hl_lines=[18, 32, 46]}
package equalfold

import (
	"strings"
	"testing"
)

var (
	lefts   = []string{"go", "ｇｏ"}
	rights  = []string{"Go", "GO", "go", "Ｇｏ", "ＧＯ", "ｇｏ"}
	rights2 = []string{"go", "go", "go", "ｇｏ", "ｇｏ", "ｇｏ"}
)

func BenchmarkEqualCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, left := range lefts {
			for _, right := range rights2 {
				if left == right {
					_ = left
				} else {
					_ = right
				}
			}
		}
	}
}

func BenchmarkEqualLower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, left := range lefts {
			for _, right := range rights {
				if left == strings.ToLower(right) {
					_ = left
				} else {
					_ = right
				}
			}
		}
	}
}

func BenchmarkEqualFold(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, left := range lefts {
			for _, right := range rights {
				if strings.EqualFold(left, right) {
					_ = left
				} else {
					_ = right
				}
			}
		}
	}
}
```

`BenchmarkEqualCase` は他の2つのコードとの比較用に書いてみた。
実行結果はこんな感じ。

```text
$ go test -bench Equal -benchmem
goos: linux
goarch: amd64
pkg: sample
BenchmarkEqualCase-4    	32061360	        36.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqualLower-4   	 1367802	       869 ns/op	      64 B/op	       8 allocs/op
BenchmarkEqualFold-4    	 3149362	       378 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	sample	4.748s
```

表にまとめておこう。

| 関数名                | 実行時間 | Alloc サイズ | Alloc 回数 |
| --------------------- | --------:| ------------:| ----------:|
| `BenchmarkEqualCase`  |  36.2 ns |      0 bytes |          0 |
| `BenchmarkEqualLower` |   869 ns |     64 bytes |          8 |
| `BenchmarkEqualFold`  |   378 ns |      0 bytes |          0 |

`BenchmarkEqualCase` と `BenchmarkEqualFold` の比較では `BenchmarkEqualFold` のほうが10倍の時間がかかっているが，それよりも `BenchmarkEqualLower` の処理のほうが圧倒的に遅いことが分かる。
まぁメモリ・アロケーションが絡むとねぇ。

というわけで，大文字小文字を無視した文字列比較では素直に [`strings`].`EqualFold()` 関数を使いましょう，という話でした。

## 【付録】 “NUL” 文字の比較

まるきし余談ではあるが

- [Big Sky :: Go で大文字小文字無視の文字列比較ベンチマーク](https://mattn.kaoriya.net/software/lang/go/20190806152526.htm)

この記事にある `isDevNull3` 関数について

```go
func isDevNull3(name string) bool {
    return strings.ToLower(name) == "nul"
}
```

 [`strings`].`EqualFold()` 関数を使うよう書き換えてみる。

 ```go {hl_lines=[2]}
 func isDevNull3(name string) bool {
 	return strings.EqualFold(name, "nul")
 }
 ```

これでベンチマークテストを実行すると

```text
goos: linux
goarch: amd64
pkg: sample/lowercase
BenchmarkS1-4   	41640913	        27.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkS2-4   	35464141	        30.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkS3-4   	12628962	        94.4 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	sample/lowercase	3.578s
```

メモリ・アロケーションが発生しなくなり，かなり速くなる。
まぁ，それでもいっちゃん遅いのだが（笑）

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[GolangCI]: https://golangci.com/ "Automated code review for Go"
[`strings`]: https://golang.org/pkg/strings/ "strings - The Go Programming Language"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN), 4621300253 (ISBN), 9784621300251 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
