+++
title = "こんな埼「玉」修正してやるぅ"
date =  "2020-07-14T13:12:59+09:00"
description = "問題となっているのは「康熙部首（kangxi radicals）」と呼ばれる漢字の部首を集めた以下の Unicode 符号点領域である。"
image = "/images/attention/go-logo_blue.png"
tags = ["golang", "unicode", "normalization", "character", "transform", "programming"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Twitter を眺めていたら

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">総務省のマイナンバーカード交付状況をデータ化していて、データの結合がうまくいかないなーと思ったら、なんと、同じ埼玉でも文字コードが違うという落とし穴が・・<br>3月8日では \u2f5f が使われていて、それ以降では\u7389…<a href="https://t.co/AOkV3iojaz">https://t.co/AOkV3iojaz</a> <a href="https://t.co/jU4P583Ad5">pic.twitter.com/jU4P583Ad5</a></p>&mdash; Hal Seki (@hal_sk) <a href="https://twitter.com/hal_sk/status/1281853581218336768?ref_src=twsrc%5Etfw">July 11, 2020</a></blockquote> {{< /fig-gen >}}

という tweet を見かけた。
これは Adobe Acrobat Distiller の不具合なんだそうで，2019年9月には既に話題に登っているのだが，2020年7月の時点でも修正されていないようだ。

- [Create PDF, why KANJI 9AD8(高) will be changed to 2... - Adobe Support Community - 10625575](https://community.adobe.com/t5/acrobat/create-pdf-why-kanji-9ad8-%E9%AB%98-will-be-changed-to-2fbc-when-meiryo-ui/td-p/10625575)

Adobe Acrobat Distiller が見捨てられてるのか，それとも「日本語」が見捨てられているのか...

{{< div-box type="markdown" >}}
**【2021-0129 追記】**
PostScript プリンタドライバ経由で PostScript データを吐き出すと似たようなことが起きるらしい。もしかしたら [Distiller の不具合ってこれが原因？](https://twitter.com/trueroad_jp/status/1354445342461235202) かも。
{{< /div-box >}}

{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}。
問題となっているのは「康熙部首（kangxi radicals）」と呼ばれる漢字の部首を集めた以下の Unicode 符号点領域である。

{{< fig-img src="./unicode-kangxi-radical.png" title="“Kangxi radical - Wikipedia” より抜粋" link="https://en.wikipedia.org/wiki/Kangxi_radical" width="696" >}}

要するに，これらの領域の文字を本来の符号点に変換してやればいいわけだ。
件の tweet のスレッドを見ると，幸いにも Unicode の NFKC 正規化で変換可能らしい。

試しに以下のコードを組んで

```go
package main

import (
    "fmt"

    "golang.org/x/text/unicode/norm"
)

func main() {
    for r := rune(0x2f00); r <= 0x2fd5; r++ {
        rr := []rune(norm.NFKC.String(string([]rune{r})))
        if r != rr[0] {
            fmt.Printf("%#U -(NFKC)-> %#U\n", r, rr[0])
        }
    }
}
```

実行してみると

```text
$ go run sample1.go 
U+2F00 '⼀' -(NFKC)-> U+4E00 '一'
U+2F01 '⼁' -(NFKC)-> U+4E28 '丨'
U+2F02 '⼂' -(NFKC)-> U+4E36 '丶'
...
U+2F5F '⽟' -(NFKC)-> U+7389 '玉'
...
U+2FD4 '⿔' -(NFKC)-> U+9F9C '龜'
U+2FD5 '⿕' -(NFKC)-> U+9FA0 '龠'
```

てな出力（一部割愛）になって，ちゃんと変換されていることが分かる。

ただし [Unicode 正規化は副作用がある]({{< relref "./unicode-normalization.md" >}} "Go 言語と Unicode 正規化")ので安直には使えない。
となると，前回の「[かなカナ変換]({{< relref "./kana-conversion.md" >}})」で紹介した方法が使えるかな。

変換後の符号点の値は散らばっていて且つ数も多く手作業でコードを書くのは不毛なので，まずは [`unicode`]`.SpecialCase` を生成するコードを書いてみよう（笑）

```go
package main

import (
    "fmt"
    "strconv"

    "golang.org/x/text/unicode/norm"
)

func main() {
    fmt.Println("var KangxiRadicals = unicode.SpecialCase{")
    for kr := rune(0x2f00); kr <= 0x2fd5; kr++ {
        rr := []rune(norm.NFKC.String(string([]rune{kr})))
        if kr != rr[0] {
            fmt.Printf("\tunicode.CaseRange{%#[1]x, %#[1]x, [unicode.MaxCase]rune{%#[2]x - %#[1]x, 0, 0}}, // %#[1]U -> %#[2]U\n", kr, rr[0])
        }
    }
    fmt.Println("}")
}
```

これを実行するとこんなコードが得られる（一部割愛）。

```text
$ go run sample1b.go
var KangxiRadicals = unicode.SpecialCase{
    unicode.CaseRange{0x2f00, 0x2f00, [unicode.MaxCase]rune{0x4e00 - 0x2f00, 0, 0}}, // U+2F00 '⼀' -> U+4E00 '一'
    unicode.CaseRange{0x2f01, 0x2f01, [unicode.MaxCase]rune{0x4e28 - 0x2f01, 0, 0}}, // U+2F01 '⼁' -> U+4E28 '丨'
    unicode.CaseRange{0x2f02, 0x2f02, [unicode.MaxCase]rune{0x4e36 - 0x2f02, 0, 0}}, // U+2F02 '⼂' -> U+4E36 '丶'
    ...
    unicode.CaseRange{0x2f5f, 0x2f5f, [unicode.MaxCase]rune{0x7389 - 0x2f5f, 0, 0}}, // U+2F5F '⽟' -> U+7389 '玉'
    ...
    unicode.CaseRange{0x2fd4, 0x2fd4, [unicode.MaxCase]rune{0x9f9c - 0x2fd4, 0, 0}}, // U+2FD4 '⿔' -> U+9F9C '龜'
    unicode.CaseRange{0x2fd5, 0x2fd5, [unicode.MaxCase]rune{0x9fa0 - 0x2fd5, 0, 0}}, // U+2FD5 '⿕' -> U+9FA0 '龠'
}
```

あとはこれを組み込んで使えばいいだけ。
たとえばこんな感じに使える。

```go {hl_lines=[12]}
func unicodePrint(s string) {
    ss := []string{}
    for _, r := range s {
        ss = append(ss, fmt.Sprintf("{%#U}", r))
    }
    fmt.Println(strings.Join(ss, " "))
}

func main() {
    saitama := "埼⽟"
    unicodePrint(saitama)
    unicodePrint(strings.ToUpperSpecial(KangxiRadicals, saitama))
}
```

これを実行すると

```text
go run sample2.go
{U+57FC '埼'} {U+2F5F '⽟'}
{U+57FC '埼'} {U+7389 '玉'}
```

となる。
よーし，うむうむ，よーし。

## ブックマーク

- [[BOD供養寺] スクレイピングしてきたデータの文字コードがおかしかったので修正した - Qiita](https://qiita.com/hal_sk/items/8a95e9daa17b500f3f27)
- [Go 言語と Unicode 正規化]({{< relref "./unicode-normalization.md" >}})

[Go]: https://go.dev/
[`strings`]: https://pkg.go.dev/strings "strings package · pkg.go.dev"
[`unicode`]: https://pkg.go.dev/unicode "unicode package · pkg.go.dev"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
