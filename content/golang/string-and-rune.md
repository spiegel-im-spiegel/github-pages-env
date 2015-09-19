+++
date = "2015-09-19T23:45:56+09:00"
description = "今回は文字列について。短めにさくっと。"
draft = false
tags = ["golang", "string", "rune", "character"]
title = "String と Rune"

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

（初出： [はじめての Go 言語 (on Windows) その4 - Qiita](http://qiita.com/spiegel-im-spiegel/items/556166b6631c0369754f)）

文字列を示す [string] は不変（immutable）なオブジェクトだが，中身は byte 配列である。
したがって以下のように

```go
package main

import "fmt"

func main() {
	nihongo := "日本語"
	size := len(nihongo)

	fmt.Printf("nihongo = %d bytes :", size)
	for i := 0; i < size; i++ {
		fmt.Printf(" %x", nihongo[i])
	}
	fmt.Print("\n")
}
```

[string] をダンプすると以下の結果になる。

```
C:>go run string01.go
nihongo = 9 bytes : e6 97 a5 e6 9c ac e8 aa 9e
```

（ちなみに [Go 言語]で取り扱う文字列の文字エンコーディングは UTF-8 が既定である。
他の文字エンコーディングで書かれた文字列を扱うには，一度 UTF-8 に変換する処理が必要になる。
文字エンコーディングの変換については項を改めて紹介する）

（[string] なんて名前なのに）文字単位で情報を保持しているわけではないため，最初の2文字を取り出すつもりでうっかり

```go
package main

import "fmt"

func main() {
	nihongo := "日本語"

	fmt.Printf("nihongo = %s\n", nihongo)
	fmt.Printf("nippon = %s\n", nihongo[:2])
}
```

なんてなコードを書くと以下の結果になる。

```
C:>go run string02.go
nihongo = 日本語
nippon = ��
```

文字列を文字単位で扱うには [rune] を使う。
いや，ルーンってどんだけ厨二... ゲフンゲフン。

```go
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
C:>go run string03.go
nihongo = 3 characters : U+65E5 '日' U+672C '本' U+8A9E '語'
```

または， [string] に対して for range を使ってループを回すと文字（[rune]）単位で取得できる。

```go
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
C:>go run string03b.go
U+65E5 '日' starts at byte position 0
U+672C '本' starts at byte position 3
U+8A9E '語' starts at byte position 6
```

[rune] の実体は int32 で，内部的には Unicode 表現になっている。
[string] と [rune] は相互変換できるので，文字列を切り取る場合は

```go
package main

import "fmt"

func main() {
	nihongo := "日本語"

	fmt.Printf("nihongo = %s\n", nihongo)
	fmt.Printf("nippon = %s\n", string([]rune(nihongo)[:2]))
}
```

のように [string] → [][rune] → [string] と変換していけば安全に処理できる。

```
C:>go run string02b.go
nihongo = 日本語
nippon = 日本
```

もう少し細かい処理が必要なら [`unicode/utf8`](http://golang.org/pkg/unicode/utf8/) パッケージを使う手もある[^1]。

[^1]: ちなみに [`strings`](http://golang.org/pkg/strings/) パッケージは内部で [`unicode/utf8`](http://golang.org/pkg/unicode/utf8/) パッケージを使っているようだ。

## ブックマーク

- [Strings, bytes, runes and characters in Go - The Go Blog](http://blog.golang.org/strings)
- [Go言語のstring, runeの正体とは？ - golang - The Round](http://knightso.hateblo.jp/entry/2014/06/24/090719)
- [golang - Goでマルチバイトが混在した文字列をtruncateする - Qiita](http://qiita.com/hokaccha/items/3d3f45b5927b4584dbac)
- [golang - Go言語で文字コード変換 - Qiita](http://qiita.com/uchiko/items/1810ddacd23fd4d3c934)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[string]: http://golang.org/ref/spec#String_types
[rune]: http://blog.golang.org/strings "Strings, bytes, runes and characters in Go - The Go Blog"
