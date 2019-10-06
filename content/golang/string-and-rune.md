+++
date = "2015-09-19T23:45:56+09:00"
description = "今回は文字列について。短めにさくっと。"
tags = ["golang", "string", "rune", "character", "unicode"]
title = "String と Rune"

[scripts]
  mathjax = false
  mermaidjs = false
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
        fmt.Printf(" %02x", nihongo[i])
    }
    fmt.Println()
}
```

[string] をダンプすると以下の[結果](https://play.golang.org/p/acsS35_ZtHl)になる[^1]。

```text
nihongo = 9 bytes : e6 97 a5 e6 9c ac e8 aa 9e
```

[^1]: ちなみに [Go 言語]で取り扱う文字列の文字エンコーディングは UTF-8 が既定である。ソースコードがそもそも UTF-8 を要求しているし（つまりリテラルの文字列はかならず UTF-8 になる）， [string] と [rune] の相互変換も文字列が UTF-8 であることを前提にしている。しかし，実際には [string] の実体はただの byte 配列であり，中身が UTF-8 文字列であることを保証しているわけではない。通常は，未知の文字列についてはいったん byte 配列に格納しておいて，何らかの方法で UTF-8 に変換してから [string] にキャストするのが安全である（または各文字エンコーディング用の独自 type を作るか）。なお，文字エンコーディングの変換については[別の記事]({{< relref "./transform-character-encoding.md" >}} "文字エンコーディング変換")で改めて紹介する。

このように（[string] なんて名前なのに）文字単位で情報を保持しているわけではないため，最初の2文字を取り出すつもりでうっかり

```go
package main

import "fmt"

func main() {
    nihongo := "日本語"

    fmt.Printf("nihongo = %s\n", nihongo)
    fmt.Printf("nippon = %s\n", nihongo[:2])
}
```

なんてなコードを書くと以下の[結果](https://play.golang.org/p/s6FEJg9bCh_D)になる。

```text
nihongo = 日本語
nippon = ��
```

文字列を文字単位で扱うには [rune] 型を使う[^rune1]。
いや，ルーンってどんだけ厨二... ゲフンゲフン。

[^rune1]: 厳密には文字単位ではなく Unicode の符号点単位と言ったほうがいいかもしれない。 Unicode では結合文字（濁点記号など）や異体字セレクタにもコードポイントが割り当てられているので，特に日本語で文字を意識した文字列操作を行う場合は注意が必要である。ぶっちゃけ面倒くさい。

ええつと，たとえばこんな感じ。

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

[実行結果](https://play.golang.org/p/gJqGozG4rOs)はこんな感じになる。

```text
nihongo = 3 characters : U+65E5 '日' U+672C '本' U+8A9E '語' 
```

または， [string] に対して [for range 構文](http://golang.org/ref/spec#For_statements)を使ってループを回すと [rune] 単位で取得できる。

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

このコードの[実行結果](https://play.golang.org/p/tEtE-FEmoRx)はこんな感じ。

```text
U+65E5 '日' starts at byte position 0
U+672C '本' starts at byte position 3
U+8A9E '語' starts at byte position 6
```

[rune] 型の実体は int32 で，内部表現は Unicode の符号点（code point）になっている。
[string] と [rune] 配列は相互変換できるので，文字列を切り取る場合は

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
上のコードの[実行結果](https://play.golang.org/p/EQEUIgkriHr)はこんな感じ。

```text
nihongo = 日本語
nippon = 日本
```

もう少し細かい処理が必要なら [`unicode/utf8`](http://golang.org/pkg/unicode/utf8/) パッケージを使う手もある[^2]。

[^2]: ちなみに [`strings`](http://golang.org/pkg/strings/) パッケージは内部で [`unicode/utf8`](http://golang.org/pkg/unicode/utf8/) パッケージを使っているようだ。

## ブックマーク

- [Strings, bytes, runes and characters in Go - The Go Blog](http://blog.golang.org/strings)
- [Go言語のstring, runeの正体とは？ - golang - The Round](http://knightso.hateblo.jp/entry/2014/06/24/090719)
- [Goでマルチバイトが混在した文字列をtruncateする - Qiita](http://qiita.com/hokaccha/items/3d3f45b5927b4584dbac)
- [Go で UTF-8 の文字列を扱う - Qiita](http://qiita.com/masakielastic/items/01a4fb691c572dd71a19)

[Go 言語に関するブックマーク集はこちら]({{< relref "bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[string]: http://golang.org/ref/spec#String_types
[rune]: http://blog.golang.org/strings "Strings, bytes, runes and characters in Go - The Go Blog"

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
