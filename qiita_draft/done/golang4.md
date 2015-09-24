# はじめての Go 言語 (on Windows) その4

（[これまでの記事の目次](http://qiita.com/spiegel-im-spiegel/items/dca0df389df1470bdbfa#%E7%9B%AE%E6%AC%A1)）

今回は文字列について。

## string と rune

文字列を示す string は immutable なオブジェクトだが，中身は byte 配列である。したがって

```go:string1.go
package main

import "fmt"

func main() {
	nihongo := "日本語"
	size := len(nihongo)

	fmt.Printf("nihongo = %d bytes : ", size)
	for i := 0; i < size; i++ {
		fmt.Printf("%x ", nihongo[i])
	}
	fmt.Printf("\n")
}
```

とすると以下の結果になる。

```shell
C:>go run string1.go
nihongo = 9 bytes : e6 97 a5 e6 9c ac e8 aa 9e
```

（ちなみに Go 言語の文字エンコーディングは UTF-8 が既定である。他の文字エンコーディングを扱うには何らかの変換パッケージが必要になる。文字エンコーディングの変換は言語に依らず相当 sensitive な分野なので，ここでは踏み込まない）

（string なんて名前なのに）文字単位で情報を保持しているわけではないため，最初の2文字を取り出すつもりでうっかり

```go:string2.go
package main

import "fmt"

func main() {
	nihongo := "日本語"

	fmt.Printf("nihongo = %s\n", nihongo)
	fmt.Printf("nippon = %s\n", nihongo[:2])
}
```

なんてなコードを書くと以下の結果になる。

```shell
C:>go run string2.go
nihongo = 日本語
nippon = ��
```

文字列を文字単位で扱うには rune を使う。いや，ルーンってどんだけ厨二... ゲフンゲフン。

```go:string3.go
package main

import "fmt"

func main() {
	nihongo := "日本語"
	nihongoRune := []rune(nihongo)
	size := len(nihongoRune)

	fmt.Printf("nihongo = %d characters : ", size)
	for i := 0; i < size; i++ {
		fmt.Printf("%#U ", nihongoRune[i])
	}
	fmt.Printf("\n")
}
```

```shell
C:>go run string3.go
nihongo = 3 characters : U+65E5 '日' U+672C '本' U+8A9E '語'
```

または， string に対して for range を使ってループを回すと文字（rune）単位で取得できる。

```go:string3b.go
package main

import "fmt"

func main() {
	nihongo := "日本語"

	for pos, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, pos)
	}
}
```

```shell
C:>go run string3b.go
U+65E5 '日' starts at byte position 0
U+672C '本' starts at byte position 3
U+8A9E '語' starts at byte position 6
```

rune の実体は int32 で，内部的には Unicode（UTF-32 相当？）表現になっている。 Go 言語式で書くなら

```go
type rune int32
```

って感じだろうか。

string と rune は相互変換できるので，文字列を切り取る場合は

```go:string4.go
package main

import "fmt"

func main() {
	nihongo := "日本語"

	fmt.Printf("nihongo = %s\n", nihongo)
	fmt.Printf("nippon = %s\n", string([]rune(nihongo)[:2]))
}
```

のように string → []rune → string と変換していけば安全に処理できる。

```shell
C:>go run string4.go
nihongo = 日本語
nippon = 日本
```

関数化するならこんな感じ？

```go:string4.go
package main

import "fmt"

func main() {
	nihongo := "日本語"

	fmt.Printf("nihongo = %s\n", nihongo)
	fmt.Printf("nippon = %s\n", substring(nihongo, 0, 2))
}

func substring(s string, start, length int) string {
	runes := []rune(s)
	if start >= len(runes) {
		return "";
	} else if start < 0 {
		start = 0
	}
	if length < 0 {
		length = 0
	}
	end := start + length
	if end >= len(runes) {
		return string(runes[start:])
	} else {
		return string(runes[start:end])
	}
}
}
```

もう少し細かい処理が必要なら [unicode/utf8](http://golang.org/pkg/unicode/utf8/) パッケージを使う手もある。ちなみに [strings](http://golang.org/pkg/strings/) パッケージは内部で [unicode/utf8](http://golang.org/pkg/unicode/utf8/) パッケージを使っているようだ。

## ブックマーク

- [Strings, bytes, runes and characters in Go - The Go Blog](http://blog.golang.org/strings)
- [Go言語のstring, runeの正体とは？ - golang - The Round](http://knightso.hateblo.jp/entry/2014/06/24/090719)
- [golang - Goでマルチバイトが混在した文字列をtruncateする - Qiita](http://qiita.com/hokaccha/items/3d3f45b5927b4584dbac)
- [golang - Go言語で文字コード変換 - Qiita](http://qiita.com/uchiko/items/1810ddacd23fd4d3c934)
