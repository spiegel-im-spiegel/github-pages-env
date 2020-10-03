+++
title = "バイト列の同値性（『プログラミング言語 Go』読書会より）"
date =  "2020-10-03T19:27:34+09:00"
description = "色んな方法で比較してみる"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "benchmark", "string" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

「[第5回『プログラミング言語Go』オンライン読書会](https://gpl-reading.connpass.com/event/188380/)」に参加してみたです。

以前 Discord で某エベントに参加したことがあったけど，やっぱオンラインイベントはしんどい。
それでもメインで喋る人が決まってるから，今回は楽なほうだったかな。
Discord にしろ Zoom にしろ，オンラインでフリートーク・イベントはかなり難しいと思うのだが，みんなどうしてるんだろう。

{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}。

今回，のっけから面白い話を聞いた。
『[プログラミング言語 Go]』4.2章「スライス」に書かれている

{{< fig-quote type="markdown" title="プログラミング言語 Go" link="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
{{% quote %}}配列と異なりスライスは比較可能ではありませんので，二つのスライスが同じ要素で構成されているかを検査するために `==` は使えません。標準ライブラリは二つのバイトスライス（`[]byte`）を比較するために**高度に最適化**された `bytes.Equal` 関数を提供しています。しかし...{{% /quote %}}
{{< /fig-quote >}}

という部分（強調は私がやりました）。
実際にコードを見てみると

```go
// Equal reports whether a and b
// are the same length and contain the same bytes.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []byte) bool {
	// Neither cmd/compile nor gccgo allocates for these string conversions.
	return string(a) == string(b)
}
```

てなことになっている。
で，「これのどこが『高度に最適化』なん？」という話があったらしい。

バージョンを遡ってみると 1.12.7 までは

```go
// Equal returns a boolean reporting whether a and b
// are the same length and contain the same bytes.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []byte) bool {
	return bytealg.Equal(a, b)
}
```

となっていた。
ちなみに `internal/bytealg` は内部パッケージで，中身はほぼ（アーキテクチャ毎に）アセンブラで記述されている。
実は [Go] 1.13 では `string` 周りが大幅に強化されていて，その辺の影響が出ているのかもしれない。

それなら，どの程度のパフォーマンスか気になるよね。
というわけで，こんなコードを用意してみた[^pp1]。

[^pp1]: ちなみに “I could tell you my pass phrase, but then I would have to kill you.” という物騒なフレーズは Simson Garfinkel 氏の『[PGP]』に載っていたパスフレーズの事例である（笑）

```go
package eq

import (
	"bytes"
	"fmt"
	"testing"
)

var (
	hello1  = "I could tell you my pass phrase, but then I would have to kill you."
	hello2  = "I could tell you my pass phrase, but then I would have to kill you."
	helloA1 = [67]byte{0x49, 0x20, 0x63, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x74, 0x65, 0x6c, 0x6c, 0x20, 0x79, 0x6f, 0x75, 0x20, 0x6d, 0x79, 0x20, 0x70, 0x61, 0x73, 0x73, 0x20, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x2c, 0x20, 0x62, 0x75, 0x74, 0x20, 0x74, 0x68, 0x65, 0x6e, 0x20, 0x49, 0x20, 0x77, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x6b, 0x69, 0x6c, 0x6c, 0x20, 0x79, 0x6f, 0x75, 0x2e}
	helloA2 = [67]byte{0x49, 0x20, 0x63, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x74, 0x65, 0x6c, 0x6c, 0x20, 0x79, 0x6f, 0x75, 0x20, 0x6d, 0x79, 0x20, 0x70, 0x61, 0x73, 0x73, 0x20, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x2c, 0x20, 0x62, 0x75, 0x74, 0x20, 0x74, 0x68, 0x65, 0x6e, 0x20, 0x49, 0x20, 0x77, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x6b, 0x69, 0x6c, 0x6c, 0x20, 0x79, 0x6f, 0x75, 0x2e}
	helloS1 = []byte{0x49, 0x20, 0x63, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x74, 0x65, 0x6c, 0x6c, 0x20, 0x79, 0x6f, 0x75, 0x20, 0x6d, 0x79, 0x20, 0x70, 0x61, 0x73, 0x73, 0x20, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x2c, 0x20, 0x62, 0x75, 0x74, 0x20, 0x74, 0x68, 0x65, 0x6e, 0x20, 0x49, 0x20, 0x77, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x6b, 0x69, 0x6c, 0x6c, 0x20, 0x79, 0x6f, 0x75, 0x2e}
	helloS2 = []byte{0x49, 0x20, 0x63, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x74, 0x65, 0x6c, 0x6c, 0x20, 0x79, 0x6f, 0x75, 0x20, 0x6d, 0x79, 0x20, 0x70, 0x61, 0x73, 0x73, 0x20, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x2c, 0x20, 0x62, 0x75, 0x74, 0x20, 0x74, 0x68, 0x65, 0x6e, 0x20, 0x49, 0x20, 0x77, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x6b, 0x69, 0x6c, 0x6c, 0x20, 0x79, 0x6f, 0x75, 0x2e}
)

func BenchmarkByteEqual1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if hello1 != hello2 {
			fmt.Printf("false")
		}
	}
}

func BenchmarkByteEqual2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !bytes.Equal([]byte(hello1), []byte(hello2)) {
			fmt.Printf("false")
		}
	}
}

func BenchmarkByteEqual3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if helloA1 != helloA2 {
			fmt.Printf("false")
		}
	}
}

func BenchmarkByteEqual4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if string(helloA1[:]) != string(helloA2[:]) {
			fmt.Printf("false")
		}
	}
}

func BenchmarkByteEqual5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !bytes.Equal(helloA1[:], helloA2[:]) {
			fmt.Printf("false")
		}
	}
}

func BenchmarkByteEqual6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if string(helloS1) != string(helloS2) {
			fmt.Printf("false")
		}
	}
}

func BenchmarkByteEqual7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !bytes.Equal(helloS1, helloS2) {
			fmt.Printf("false")
		}
	}
}
```

`hello1`, `hello2` は `string` 型，`helloA1`, `helloA2` は `byte` 配列，`helloS1`, `helloS2` は `[]byte` 型で中身はみんな同じ。
これらを使って同値性（equallity）をチェックするのだが，内訳はこんな感じ。

| 関数名                | 比較手順                                      |
| --------------------- | --------------------------------------------- |
| `BenchmarkByteEqual1` | `string == string`                            |
| `BenchmarkByteEqual2` | `bytes.Equal([]byte(string), []byte(string))` |
| `BenchmarkByteEqual3` | `<byte array> == <byte array>`                |
| `BenchmarkByteEqual4` | `string(<array>[:]) == string(<array>[:])`    |
| `BenchmarkByteEqual5` | `bytes.Equal(<array>[:], <array>[:])`         |
| `BenchmarkByteEqual6` | `string([]byte) == string([]byte)`            |
| `BenchmarkByteEqual7` | `bytes.Equal([]byte, []byte)`                 |

では，実際に動かしてみようか。

```text
$ go test -bench ByteEqual -benchmem
goos: linux
goarch: amd64
pkg: sample
BenchmarkByteEqual1-4   	301173402	         4.04 ns/op	       0 B/op	       0 allocs/op
BenchmarkByteEqual2-4   	 8552673	       122 ns/op	     160 B/op	       2 allocs/op
BenchmarkByteEqual3-4   	182939372	         6.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkByteEqual4-4   	191359716	         6.29 ns/op	       0 B/op	       0 allocs/op
BenchmarkByteEqual5-4   	191511163	         6.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkByteEqual6-4   	168382664	         7.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkByteEqual7-4   	167058468	         7.06 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	sample	12.140s
```

分かりやすく表にしておく。

| 比較手順                                      | 処理時間 (ns/op) | Alloc 回数 |
| --------------------------------------------- | ----------------:| ----------:|
| `string == string`                            |             4.04 |          0 |
| `bytes.Equal([]byte(string), []byte(string))` |           122.00 |          2 |
| `<byte array> == <byte array>`                |             6.27 |          0 |
| `string(<array>[:]) == string(array>[:])`     |             6.29 |          0 |
| `bytes.Equal(array>[:], <array>[:])`          |             6.27 |          0 |
| `string([]byte) == string([]byte)`            |             7.16 |          0 |
| `bytes.Equal([]byte, []byte)`                 |             7.06 |          0 |


`string` → `[]byte` へのキャスト時にアロケーションが発生している点に注意。

つか `string` 同士の比較処理が速いな！ 配列と slice で若干差が出るのは仕方ないが，元が同じ型なら殆ど差がないようだ。
まぁ，これなら確かに

```go
func Equal(a, b []byte) bool { return string(a) == string(b) }
```

でもいっか，って気になるよな。

## ブックマーク

- [【改訂版】文字列連結はどれが速い？]({{< relref "./join-strings-2.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[プログラミング言語 Go]: https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1
[PGP]: https://www.amazon.co.jp/exec/obidos/ASIN/4900900028/baldandersinf-22/ "Amazon | PGP―暗号メールと電子署名 | シムソン ガーフィンケル, Simson Garfinkel, ユニテック 通販"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
