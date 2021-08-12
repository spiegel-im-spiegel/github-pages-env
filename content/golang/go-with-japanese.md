+++
title = "Go と日本語"
date =  "2021-06-19T14:51:29+09:00"
description = "Go では型名や関数名や変数名といった識別名として Unicode 文字が使える。 "
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "language", "japanese", "unicode" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Twitter でちょろんと書いた](https://twitter.com/spiegel_2007/status/1406010487565783040)が，なんか勝手に NFC 正規化されてしまうみたいだし，このブログではもう少し掘り下げて書いてみる。

[Go] では型名や関数名や変数名といった識別名（identifiers name）として Unicode 文字が使える。
ということは名前を日本語にすることも可能である。
まぁ，どんな文字でも使えるわけではないが。

たとえば以下のようなコードを考える。

```go
// +build run

package main

type ペンギン struct{}

func (p *ペンギン) 鳴く() string {
    return "ふるるー！"
}

func main() {
    println((&ペンギン{}).鳴く())
}
```

型名や関数名に Unicode 文字が使えるなら上のコードはうまく動きそうだが，実際にはコンパイルエラーになる。

```text
$ go run sample1.go
# command-line-arguments
./sample.go:12:15: undefined: ペンギン
./sample.go:12:18: invalid character U+309A '゚' in identifier
./sample.go:12:27: invalid character U+3099 '゙' in identifier
```

コンパイルエラーになる理由は2つ。

1. `type` で宣言した型名と `main()` 関数内で使われている型名が異なる「文字列」である
2. 半濁点および濁点を表す結合文字 `U+309A` および `U+3099` は識別名として使えない

なんだか面倒くさそうである。

## 識別名として使える文字

[Go の仕様書][Go Specs] を読むと識別名として使える文字が示されている。

{{< fig-quote class="nobox" type="markdown" title="The Go Programming Language Specification" link="https://golang.org/ref/spec" lang="en" >}}
```
identifier = letter { letter | unicode_digit } .
```
{{< /fig-quote >}}

このうち `letter` は

{{< fig-quote class="nobox" type="markdown" title="The Go Programming Language Specification" link="https://golang.org/ref/spec" lang="en" >}}
```
letter = unicode_letter | "_" .
```
{{< /fig-quote >}}

と定義され Unicode 文字が使えることを示している。

`unicode_digit` や `unicode_letter` を含む文字の定義は以下の通り。

{{< fig-quote class="nobox" type="markdown" title="The Go Programming Language Specification" link="https://golang.org/ref/spec" lang="en" >}}
```
newline        = /* the Unicode code point U+000A */ .
unicode_char   = /* an arbitrary Unicode code point except newline */ .
unicode_letter = /* a Unicode code point classified as "Letter" */ .
unicode_digit  = /* a Unicode code point classified as "Number, decimal digit" */ .
```
{{< /fig-quote >}}

`unicode_digit` は全角数字も含んでいる点に注意。
改行コードも LF なんだねぇ（Windows の CRLF は早く捨てたい）。

では Letter や Number の文字種が具体的にどの文字を指しているのか。
幸い [`unicode`] 標準パッケージを使って文字種をチェックできるので，簡単なチェッカーを書いてみよう。

```go
// +build run

package main

import (
    "flag"
    "fmt"
    "unicode"
)

var mapRangeTable = map[*unicode.RangeTable]string{
    unicode.Variation_Selector: "Variation Selector",
    unicode.Sc:                 "Symbol/currency",
    unicode.Sk:                 "Symbol/modifier",
    unicode.Sm:                 "Symbol/math",
    unicode.Radical:            "Radical",
    unicode.So:                 "Symbol/other",
    unicode.Lm:                 "Letter/modifier",
    unicode.Ideographic:        "Ideographic",
    unicode.Lo:                 "Letter/other",
    unicode.Nl:                 "Number/letter",
    unicode.No:                 "Number/other",
    unicode.Mc:                 "Mark/spacing combining",
    unicode.Me:                 "Mark/enclosing",
    unicode.Mn:                 "Mark/nonspacing",
    unicode.Pc:                 "Punctuation/connector",
    unicode.Pd:                 "Punctuation/dash",
    unicode.Pe:                 "Punctuation/close",
    unicode.Pf:                 "Punctuation/final quote",
    unicode.Pi:                 "Punctuation/initial quote",
    unicode.Ps:                 "Punctuation/open",
    unicode.Po:                 "Punctuation/other",
    unicode.Zl:                 "Separator/line",
    unicode.Zp:                 "Separator/paragraph",
    unicode.Zs:                 "Separator/space",
    unicode.Join_Control:       "Join Control",
    unicode.Cc:                 "Control/control",
    unicode.Cf:                 "Control/format",
    unicode.Cs:                 "Control/surrogate",
    unicode.Co:                 "Control/private use",
}

var arryRangeTable = []*unicode.RangeTable{
    unicode.Variation_Selector,
    unicode.Sc,
    unicode.Sk,
    unicode.Sm,
    unicode.Radical,
    unicode.So,
    unicode.Lm,
    unicode.Ideographic,
    unicode.Lo,
    unicode.Nl,
    unicode.No,
    unicode.Mc,
    unicode.Me,
    unicode.Mn,
    unicode.Pc,
    unicode.Pd,
    unicode.Pe,
    unicode.Pf,
    unicode.Pi,
    unicode.Ps,
    unicode.Po,
    unicode.Zl,
    unicode.Zp,
    unicode.Zs,
    unicode.Join_Control,
    unicode.Cc,
    unicode.Cf,
    unicode.Cs,
    unicode.Co,
}

type unicodeChecker struct {
    mapRangeTable  map[*unicode.RangeTable]string
    arryRangeTable []*unicode.RangeTable
}

func newChecker() *unicodeChecker {
    return &unicodeChecker{mapRangeTable: mapRangeTable, arryRangeTable: arryRangeTable}
}

func (c *unicodeChecker) string(rt *unicode.RangeTable) string {
    if c == nil || rt == nil {
        return "Unknown"
    }
    if s, ok := c.mapRangeTable[rt]; ok {
        return s
    }
    return "Unknown"
}

func (c *unicodeChecker) check(r rune) string {
    if c == nil {
        return c.string(nil)
    }
    for _, rt := range c.arryRangeTable {
        if unicode.Is(rt, r) {
            return c.string(rt)
        }
    }
    return c.string(nil)
}

func main() {
    flag.Parse()
    args := flag.Args()
    checker := newChecker()
    for _, arg := range args {
        fmt.Println(arg)
        for i, r := range arg {
            switch {
            case unicode.IsLetter(r), r == '_':
                fmt.Printf("\t%#U : OK : letter\n", r)
            case unicode.IsNumber(r):
                okng := "OK"
                if i == 0 {
                    okng = "NG"
                }
                fmt.Printf("\t%#U : %v : unicode_digit\n", r, okng)
            default:
                fmt.Printf("\t%#U : NG : %v\n", r, checker.check(r))
            }
        }
    }
}
```

ここで `unicodeChecker.check()` メソッドは「[Unicode 文字種の判別]({{< relref "./unicode-rangetables.md" >}})」で書いた `check()` 関数の改良版で，かなり細かくチェックできるようにしている。
実際に動かしてみよう。

```text
$ go run sample2.go ペンギン ペンギン
ペンギン
	U+30DA 'ペ' : OK : letter
	U+30F3 'ン' : OK : letter
	U+30AE 'ギ' : OK : letter
	U+30F3 'ン' : OK : letter
ペンギン
	U+30D8 'ヘ' : OK : letter
	U+309A '゚' : NG : Mark/nonspacing
	U+30F3 'ン' : OK : letter
	U+30AD 'キ' : OK : letter
	U+3099 '゙' : NG : Mark/nonspacing
	U+30F3 'ン' : OK : letter
```

うんうん。
ちなみに絵文字は，複数のコードポイントで構成されている場合は当然として，単一のコードポイントでも `letter` とは見なされないらしい。

```text
$ go run sample2.go ⌛
⌛
	U+231B '⌛' : NG : Symbol/other
```

あと康熙部首とかもアカン感じ？

```text
$ go run sample2.go 埼⽟ 埼玉
埼⽟
	U+57FC '埼' : OK : letter
	U+2F5F '⽟' : NG : Radical
埼玉
	U+57FC '埼' : OK : letter
	U+7389 '玉' : OK : letter
```

## やっぱ英語にしとこうや

どこで見たのか失念してしまったが，メジャーなプログラミング言語，特に手続き型の言語は英語の文法や語感が前提になっていて，日本人はそこから既にハンデがあると読んだことがある。

そういう言語体系の中に日本語を混ぜ込むと，日本語が刷り込まれている人ほど奇妙に見えてしまうんじゃないだろうか。
だから私のように英語不得手でも，プログラム・コードは「英語」で書くのが一番自然だろう。

「[悲しいけどこれ現実なのよね](https://dic.pixiv.net/a/%E6%82%B2%E3%81%97%E3%81%84%E3%81%91%E3%81%A9%E3%81%93%E3%82%8C%E3%80%81%E7%8F%BE%E5%AE%9F%E3%81%AA%E3%81%AE%E3%82%88%E3%81%AD)」

[Go]: https://golang.org/ "The Go Programming Language"
[Go Specs]: https://golang.org/ref/spec "The Go Programming Language Specification - The Go Programming Language"
[`unicode`]: https://golang.org/pkg/unicode/ "unicode - The Go Programming Language"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B07143DX2F" %}} <!-- with #前置詞といっしょ！ -->
{{% review-paapi "4906985165" %}} <!-- coffee ＃名詞のヒミツ！ -->
{{% review-paapi "4906985262" %}} <!-- #動詞と踏み出そう！ -->
