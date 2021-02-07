+++
title = "Unicode 文字種の判別"
date =  "2021-01-22T20:26:18+09:00"
description = "Unicode 文字種を判別するには unicode 標準パッケージが使える。"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "unicode", "character" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Zenn で「[やっかいな日本語](https://zenn.dev/spiegel/articles/20210118-characters)」なる{{< ruby "ポエム" >}}記事{{< /ruby >}}を書いたが，このブログでは [Go] における Unicode 文字種の判別に話題を絞って紹介してみる。

Unicode 文字種を判別するには [`unicode`] 標準パッケージが使える。

判別用の [`unicode`]`.RangeTable` を用意し，これを参照することで文字種を判別することができる。
このパッケージの中身を見ると分かるが，かなりの数の定義済み [`unicode`]`.RangeTable` テーブルが取り揃えられている。
今回はこの定義済みテーブルのみ使うことにしよう。

## 図形文字と制御文字

まずは大雑把に「図形文字」と「制御文字」を判別してみよう。

図形文字の判別には [`unicode`]`.IsGraphic()` 関数が，制御文字の判別には [`unicode`]`.IsControl()` 関数が使える。

ただし [`unicode`]`.IsControl()` 関数では `U+00FF` 以下の ISO 8859[^iso8859] で定義されている制御文字領域しか判別してくれないようで BOM (`U+FEFF`) などの Unicode 独自の制御文字も含めて判別するのであれば [`unicode`]`.C` テーブルを使う必要がある。

[^iso8859]: 8ビット空間の符号化文字集合および文字エンコーディング。国や言語ごとにいくつかのバリエーションがある。最も有名なのはドイツ語やフランス語の文字を含む ISO 8859-1，通称 Latin-1 だろう。日本の JIS X 0201 も ISO 8859 のバリエーションと言える。

そこで，こんな関数を考えてみる。

```go
import "unicode"

func check(r rune) string {
    switch {
    case unicode.IsGraphic(r):
        return "Graphic"
    case unicode.IsControl(r):
        return "Latin1 Control"
    case unicode.Is(unicode.C, r):
        return "Unicode Control"
    }
    return "Unknown"
}
```

これを使って実際に文字列をチェックしてみよう。

```go:sample1.go
func main() {
    text := string([]byte{0xef, 0xbb, 0xbf, 0xe3, 0x82, 0x84, 0x09, 0xe3, 0x81, 0x82})
    fmt.Println(text)
    for _, c := range text {
        fmt.Printf("%#U (%v)\n", c, check(c))
    }
}
```

これを実行すると

```text
$ go run sample1.go
﻿や     あ
U+FEFF (Unicode Control)
U+3084 'や' (Graphic)
U+0009 (Latin1 Control)
U+3042 'あ' (Graphic)
```

となった。うんうん。

## 結合子と異体字セレクタ

上述の `check()` 関数を使って，今度は絵文字の中身を見てみる。

```go
func main() {
    text := "🙇‍♂️"
    for _, c := range text {
        fmt.Printf("%#U (%v)\n", c, check(c))
    }
}
```

これを実行すると

```text
$ go run sample2.go
U+1F647 '🙇' (Graphic)
U+200D (Unicode Control)
U+2642 '♂' (Graphic)
U+FE0F '️' (Graphic)
```

となった。

ありゃ。
ZWJ はともかく異体字セレクタって図形文字あつかいなんだ。

これでは大雑把すぎるので `check()` 関数にいくつか条件を足して

```go {hl_lines=["3-14"]}
func check(r rune) string {
    switch {
    case unicode.Is(unicode.Sc, r):
        return "Symbol/currency"
    case unicode.Is(unicode.Sk, r):
        return "Symbol/modifier"
    case unicode.Is(unicode.Sm, r):
        return "Symbol/math"
    case unicode.Is(unicode.So, r):
        return "Symbol/other"
    case unicode.Is(unicode.Variation_Selector, r):
        return "Variation Selector"
    case unicode.Is(unicode.Join_Control, r):
        return "Join Control"
    case unicode.IsGraphic(r):
        return "Graphic"
    case unicode.IsControl(r):
        return "Latin1 Control"
    case unicode.Is(unicode.C, r):
        return "Unicode Control"
    }
    return "Unknown"
}
```

と書き換えてみる。
これを使ってもう一度実行してみると

```text
$ go run sample2.go
U+1F647 '🙇' (Symbol/other)
U+200D (Join Control)
U+2642 '♂' (Symbol/other)
U+FE0F '️' (Variation Selector)
```

となった。これで結合子や異体字セレクタをきちんと判別できる。
なお，シンボルについて細かく区別しなくていいのなら [`unicode`]`.IsSymbol()` 関数を使う手もある。

## 漢字と部首

以前「[こんな埼「玉」修正してやるぅ]({{< relref "./unicode-kangxi-radical.md" >}})」でも書いたが， Unicode では漢字の部首にもコードポイントが割り当てられている。
しかし，幸いなことに [`unicode`] 標準パッケージの定義済み [`unicode`]`.RangeTable` テーブルで部首を判別可能である。

具体的には `check()` 関数を以下のように書き換える。

```go {hl_lines=["3-6"]}
func check(r rune) string {
    switch {
    case unicode.Is(unicode.Radical, r):
        return "Radical"
    case unicode.Is(unicode.Ideographic, r):
        return "Ideographic"
    case unicode.IsGraphic(r):
        return "Graphic"
    case unicode.IsControl(r):
        return "Latin1 Control"
    case unicode.Is(unicode.C, r):
        return "Unicode Control"
    }
    return "Unknown"
}
```

これを使えば

```go
func main() {
    text := "⽟玉"
    for _, c := range text {
        fmt.Printf("%#U (%v)\n", c, check(c))
    }
}
```

の実行結果が

```text
$ go run sample3.go 
U+2F5F '⽟' (Radical)
U+7389 '玉' (Ideographic)
```

となる。

なお， [`unicode`]`.Ideographic` テーブルで判別できるのは本当に漢字だけなので，全角の英数字・かな文字・記号は，これにかからない。
ちなみに，部首は絵文字と同じくシンボル扱いなので [`unicode`]`.IsSymbol()` 関数でも一応は区別できる。

## 3羽の「ペンギン」

次は `check()` 関数をかな文字を判別するよう書き換える。
こんな感じ。

```go {hl_lines=["3-16"]}
func check(r rune) string {
    switch {
    case unicode.Is(unicode.Katakana, r):
        return "Katakana"
    case unicode.Is(unicode.Hiragana, r):
        return "Hiragana"
    case unicode.Is(unicode.Lm, r):
        return "Letter/modifier"
    case unicode.Is(unicode.Lo, r):
        return "Letter"
    case unicode.Is(unicode.Mc, r):
        return "Mark/spacing combining"
    case unicode.Is(unicode.Me, r):
        return "Mark/enclosing"
    case unicode.Is(unicode.Mn, r):
        return "Mark/nonspacing"
    case unicode.IsGraphic(r):
        return "Graphic"
    case unicode.IsControl(r):
        return "Latin1 Control"
    case unicode.Is(unicode.C, r):
        return "Unicode Control"
    }
    return "Unknown"
}
```

これを使って以下の文字列を判別してみる。

```go
func main() {
    text := "ペンギンペンギンﾍﾟﾝｷﾞﾝ"
    for _, c := range text {
        fmt.Printf("%#U (%v)\n", c, check(c))
    }
}
```

実行結果は以下の通り。

```text
$ go run sample4.go 
U+30DA 'ペ' (Katakana)
U+30F3 'ン' (Katakana)
U+30AE 'ギ' (Katakana)
U+30F3 'ン' (Katakana)
U+30D8 'ヘ' (Katakana)
U+309A '゚' (Mark/nonspacing)
U+30F3 'ン' (Katakana)
U+30AD 'キ' (Katakana)
U+3099 '゙' (Mark/nonspacing)
U+30F3 'ン' (Katakana)
U+FF8D 'ﾍ' (Katakana)
U+FF9F 'ﾟ' (Letter/modifier)
U+FF9D 'ﾝ' (Katakana)
U+FF77 'ｷ' (Katakana)
U+FF9E 'ﾞ' (Letter/modifier)
U+FF9D 'ﾝ' (Katakana)
```

濁点や半濁点の文字種が全角と半角で異なっている点に注意。
ホンマ，面倒くさいったら。

## 面倒な Unicode

[`unicode`] 標準パッケージにある定義済み [`unicode`]`.RangeTable` テーブルはよくできてるし，ある程度は日本語も考慮されているけど，細かい制御を行うのであれば用途に応じて専用の [`unicode`]`.RangeTable` テーブルを用意したほうがいいだろう。
量が多くて面倒くさいけどね。

## ブックマーク

- [その文字が JIS X 0208 に含まれるか？ あるいは unicode.RangeTable の使い方](https://zenn.dev/ikawaha/articles/20210116-ab1ac4a692ae8bb4d9cf)
- [Golangでひらがな、カタカナ、漢字を判定する - Qiita](https://qiita.com/tomtwinkle/items/d52a01d5a020b00b4b8e)

- [かなカナ変換]({{< relref "./kana-conversion.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[`unicode`]: https://golang.org/pkg/unicode/ "unicode - The Go Programming Language"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
