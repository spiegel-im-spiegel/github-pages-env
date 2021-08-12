+++
title = "かなカナ変換"
date =  "2020-07-13T18:45:08+09:00"
description = "全角⇔半角 / ひらがな⇔カタカナ / 拗音・促音⇔直音 各種変換"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "unicode", "character", "transform" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今回は仮名文字を色々と変換することを考えてみる。

## 全角⇔半角 変換

いわゆる全角文字と半角文字の変換は `golang.org/x/text/`[`width`] パッケージを使えばいいのだが，仮名文字の場合は少しクセがある。
たとえば以下のように「ペンギン」を半角カナに変換しようとすると

```go
package main

import (
    "fmt"

    "golang.org/x/text/width"
)

func unicodePrint(s string) {
    sep := ""
    for _, r := range s {
        fmt.Printf("%s{%#U}", sep, r)
        sep = " "
    }
    fmt.Println()
}

func main() {
    fullwidth := "ペンギン"

    unicodePrint(fullwidth)
    unicodePrint(width.Narrow.String(fullwidth))
}

```

実行結果は

```text
$ go run sample1.go
{U+30DA 'ペ'} {U+30F3 'ン'} {U+30AE 'ギ'} {U+30F3 'ン'}
{U+30DA 'ペ'} {U+FF9D 'ﾝ'} {U+30AE 'ギ'} {U+FF9D 'ﾝ'}
```

のように濁点・半濁点を上手く処理できないようだ。

これを解消するには，安直な手段だが，いったん NFD 正規化で合成列に変換してから変換するとよい。

```go {hl_lines=[5]}
func main() {
    fullwidth := "ペンギン"

    unicodePrint(fullwidth)
    unicodePrint(width.Narrow.String(norm.NFD.String(fullwidth)))
}
```

これで実行すると

```text {hl_lines=[3]}
$ go run sample1b.go
{U+30DA 'ペ'} {U+30F3 'ン'} {U+30AE 'ギ'} {U+30F3 'ン'}
{U+FF8D 'ﾍ'} {U+FF9F 'ﾟ'} {U+FF9D 'ﾝ'} {U+FF77 'ｷ'} {U+FF9E 'ﾞ'} {U+FF9D 'ﾝ'}
```

と，綺麗に変換してくれる[^norm1]。

[^norm1]: Unicode 正規化には副作用があり，実際のところコード変換用途にはお勧めできない。詳しくは拙文「[Go 言語と Unicode 正規化]({{< relref "./unicode-normalization.md" >}})」を参照のこと。

逆に `golang.org/x/text/`[`width`] パッケージで半角カナから全角カナに変換しようとすると

```go {hl_lines=[7]}
func main() {
    fullwidth := "ペンギン"
    halfwidth := width.Narrow.String(norm.NFD.String(fullwidth))

    unicodePrint(fullwidth)
    unicodePrint(halfwidth)
    unicodePrint(width.Widen.String(halfwidth))
}
```

実行結果は当然ながら

```text {hl_lines=[4]}
$ go run sample1c.go
{U+30DA 'ペ'} {U+30F3 'ン'} {U+30AE 'ギ'} {U+30F3 'ン'}
{U+FF8D 'ﾍ'} {U+FF9F 'ﾟ'} {U+FF9D 'ﾝ'} {U+FF77 'ｷ'} {U+FF9E 'ﾞ'} {U+FF9D 'ﾝ'}
{U+30D8 'ヘ'} {U+309A '゚'} {U+30F3 'ン'} {U+30AD 'キ'} {U+3099 '゙'} {U+30F3 'ン'}
```

合成列となる。
これを事前合成形にするのであれば

```go {hl_lines=[7]}
func main() {
    fullwidth := "ペンギン"
    halfwidth := width.Narrow.String(norm.NFD.String(fullwidth))

    unicodePrint(fullwidth)
    unicodePrint(halfwidth)
    unicodePrint(norm.NFC.String(width.Widen.String(halfwidth)))
}
```

と NFC 正規化をすることで

```text {hl_lines=[4]}
$ go run sample1c.go
{U+30DA 'ペ'} {U+30F3 'ン'} {U+30AE 'ギ'} {U+30F3 'ン'}
{U+FF8D 'ﾍ'} {U+FF9F 'ﾟ'} {U+FF9D 'ﾝ'} {U+FF77 'ｷ'} {U+FF9E 'ﾞ'} {U+FF9D 'ﾝ'}
{U+30DA 'ペ'} {U+30F3 'ン'} {U+30AE 'ギ'} {U+30F3 'ン'}
```

とできる[^norm1]。

ついでに， `'ヰ'`, `'ヱ'` 文字， `'ヽ'`, `'ヾ'` といった繰り返し記号，あるいは `'ヵ'`, `'ヶ'` といった拗音の一部は半角カナにはないので，これらを含む文字列の変換には注意が必要である。

## ひらがな⇔カタカナ 変換

[Go] 言語では，ひらがなとカタカナを相互変換するパッケージは標準では用意されていないので自作する必要がある。

まずは，ひらがなとカタカナの Unicode 符号点を眺めてみる[^dg1]。

[^dg1]: `'ゟ'` や `'ヿ'` といった合略仮名文字については，今回は無視する。片仮名拡張や仮名補助も同様に扱わない。

{{< fig-gen type="markdown" title="Unicode 符号点：ひらがな" >}}
{{% div-gen class="smaller" %}}

|        | +0  | +1  | +2  | +3  | +4  | +5  | +6  | +7  | +8  | +9  | +A  | +B  | +C  | +D  | +E  | +E  |
| ------:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| U+3040 |     | ぁ  | あ  | ぃ  | い  | ぅ  | う  | ぇ  | え  | ぉ  | お  | か  | が  | き  | ぎ  | く  |
| U+3050 | ぐ  | け  | げ  | こ  | ご  | さ  | ざ  | し  | じ  | す  | ず  | せ  | ぜ  | そ  | ぞ  | た  |
| U+3060 | だ  | ち  | ぢ  | っ  | つ  | づ  | て  | で  | と  | ど  | な  | に  | ぬ  | ね  | の  | は  |
| U+3070 | ば  | ぱ  | ひ  | び  | ぴ  | ふ  | ぶ  | ぷ  | へ  | べ  | ぺ  | ほ  | ぼ  | ぽ  | ま  | み  |
| U+3080 | む  | め  | も  | ゃ  | や  | ゅ  | ゆ  | ょ  | よ  | ら  | り  | る  | れ  | ろ  | ゎ  | わ  |
| U+3090 | ゐ  | ゑ  | を  | ん  | ゔ  | ゕ  | ゖ  |     |     | ゙  | ゚  | ゛  | ゜  | ゝ  | ゞ  |     |

{{% /div-gen %}}
{{< /fig-gen >}}

{{< fig-gen type="markdown" title="Unicode 符号点：カタカナ" >}}
{{% div-gen class="smaller" %}}

|        | +0  | +1  | +2  | +3  | +4  | +5  | +6  | +7  | +8  | +9  | +A  | +B  | +C  | +D  | +E  | +E  |
| ------:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| U+30A0 |     | ァ  | ア  | ィ  | イ  | ゥ  | ウ  | ェ  | エ  | ォ  | オ  | カ  | ガ  | キ  | ギ  | ク  |
| U+30B0 | グ  | ケ  | ゲ  | コ  | ゴ  | サ  | ザ  | シ  | ジ  | ス  | ズ  | セ  | ゼ  | ソ  | ゾ  | タ  |
| U+30C0 | ダ  | チ  | ヂ  | ッ  | ツ  | ヅ  | テ  | デ  | ト  | ド  | ナ  | ニ  | ヌ  | ネ  | ノ  | ハ  |
| U+30D0 | バ  | パ  | ヒ  | ビ  | ピ  | フ  | ブ  | プ  | ヘ  | ベ  | ペ  | ホ  | ボ  | ポ  | マ  | ミ  |
| U+30E0 | ム  | メ  | モ  | ャ  | ヤ  | ュ  | ユ  | ョ  | ヨ  | ラ  | リ  | ル  | レ  | ロ  | ヮ  | ワ  |
| U+30F0 | ヰ  | ヱ  | ヲ  | ン  | ヴ  | ヵ  | ヶ  | ヷ  | ヸ  | ヹ  | ヺ  | ・  | ー  | ヽ  | ヾ  |     |

{{% /div-gen %}}
{{< /fig-gen >}}

<!--
ついでに半角カナの Unicode 符号点も挙げておく。

{{< fig-gen type="markdown" title="Unicode 符号点：半角カナ" >}}
{{% div-gen class="smaller" %}}

|        | +0  | +1  | +2  | +3  | +4  | +5  | +6  | +7  | +8  | +9  | +A  | +B  | +C  | +D  | +E  | +E  |
| ------:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| U+FF60 |     |     |     |     |     |  ･  |  ｦ  |  ｧ  |  ｨ  |  ｩ  |  ｪ  |  ｫ  |  ｬ  |  ｭ  |  ｮ  |  ｯ  |
| U+FF70 |  ｰ  |  ｱ  |  ｲ  |  ｳ  |  ｴ  |  ｵ  |  ｶ  |  ｷ  |  ｸ  |  ｹ  |  ｺ  |  ｻ  |  ｼ  |  ｽ  |  ｾ  |  ｿ  |
| U+FF80 |  ﾀ  |  ﾁ  |  ﾂ  |  ﾃ  |  ﾄ  |  ﾅ  |  ﾆ  |  ﾇ  |  ﾈ  |  ﾉ  |  ﾊ  |  ﾋ  |  ﾌ  |  ﾍ  |  ﾎ  |  ﾏ  |
| U+FF90 |  ﾐ  |  ﾑ  |  ﾒ  |  ﾓ  |  ﾔ  |  ﾕ  |  ﾖ  |  ﾗ  |  ﾘ  |  ﾙ  |  ﾚ  |  ﾛ  |  ﾜ  |  ﾝ  |  ﾞ  |  ﾟ  |

{{% /div-gen %}}
{{< /fig-gen >}}
-->

これを見るとひらがなの U+3041 〜 U+3096 と片仮名の U+30A1 〜 U+30F6 の領域が1対1で対応していることが分かる。
繰り返し記号の `'ゝ'`, `'ゞ'`, `'ヽ'`, `'ヾ'` も同様。
これなら標準の [`strings`] パッケージを使って変換できそうだ。

たとえば [`strings`]`.ToUpperSpecial()` 関数でひらがな→カタカナ変換を， [`strings`]`.ToLowerSpecial()` 関数でカタカナ→ひらがな変換を行うように [`unicode`]`.SpecialCase` の値を設定すればよい。

[`unicode`]`.SpecialCase` 構造体の定義はこうなっている。

```go
// CaseRange represents a range of Unicode code points for simple (one
// code point to one code point) case conversion.
// The range runs from Lo to Hi inclusive, with a fixed stride of 1. Deltas
// are the number to add to the code point to reach the code point for a
// different case for that character. They may be negative. If zero, it
// means the character is in the corresponding case. There is a special
// case representing sequences of alternating corresponding Upper and Lower
// pairs. It appears with a fixed Delta of
//    {UpperLower, UpperLower, UpperLower}
// The constant UpperLower has an otherwise impossible delta value.
type CaseRange struct {
    Lo    uint32
    Hi    uint32
    Delta d
}

// SpecialCase represents language-specific case mappings such as Turkish.
// Methods of SpecialCase customize (by overriding) the standard mappings.
type SpecialCase []CaseRange

// BUG(r): There is no mechanism for full case folding, that is, for
// characters that involve multiple runes in the input or output.

// Indices into the Delta arrays inside CaseRanges for case mapping.
const (
    UpperCase = iota
    LowerCase
    TitleCase
    MaxCase
)

type d [MaxCase]rune // to make the CaseRanges text shorter
```

ほんじゃあ，さくっとコードを書いてみよう。
こんな感じかな。

```go
package main

import (
    "fmt"
    "strings"
    "unicode"
)

var kanaCase = unicode.SpecialCase{
    unicode.CaseRange{'ぁ', 'ゖ', [unicode.MaxCase]rune{'ァ' - 'ぁ', 0, 0}},
    unicode.CaseRange{'ゝ', 'ゞ', [unicode.MaxCase]rune{'ヽ' - 'ゝ', 0, 0}},
    unicode.CaseRange{'ァ', 'ヶ', [unicode.MaxCase]rune{0, 'ぁ' - 'ァ', 0}},
    unicode.CaseRange{'ヽ', 'ヾ', [unicode.MaxCase]rune{0, 'ゝ' - 'ヽ', 0}},
}

func main() {
    kana := "あいうえおわゐゑをんゔゕゖゝゞアイウエオワヰヱヲンヴヵヶヽヾ"
    fmt.Println(strings.ToUpperSpecial(kanaCase, kana))
    fmt.Println(strings.ToLowerSpecial(kanaCase, kana))
}
```

この実行結果は

```text
$ go run sample2.go 
アイウエオワヰヱヲンヴヵヶヽヾアイウエオワヰヱヲンヴヵヶヽヾ
あいうえおわゐゑをんゔゕゖゝゞあいうえおわゐゑをんゔゕゖゝゞ
```

となる。
よーし，うむうむ，よーし。

なお，カタカナの `'ヷ'`, `'ヸ'`, `'ヹ'`, `'ヺ'` 文字は対応する事前合成形のひらがな文字がないため，必要なら個別に処理する必要がある。


## 拗音・促音⇔直音 変換

{{< ruby "ようおん" >}}拗音{{< /ruby >}}（小さい `'ゃ'`, `'ゅ'`, `'ょ'` など）や促音（小さい `'っ'`）と直音（`'や'`, `'ゆ'`, `'よ'`, `'つ'`）とを相互変換することを考える。
この機能も標準では用意されてないけど，前節と同じように [`unicode`]`.SpecialCase` の値を決めればいいかな。

長いコードだけど，ご容赦。

```go
package main

import (
    "fmt"
    "sort"
    "strings"
    "unicode"
)

var kanaCase2 = unicode.SpecialCase{
    unicode.CaseRange{'あ', 'あ', [unicode.MaxCase]rune{0, 'ぁ' - 'あ', 0}},
    unicode.CaseRange{'い', 'い', [unicode.MaxCase]rune{0, 'ぃ' - 'い', 0}},
    unicode.CaseRange{'う', 'う', [unicode.MaxCase]rune{0, 'ぅ' - 'う', 0}},
    unicode.CaseRange{'え', 'え', [unicode.MaxCase]rune{0, 'ぇ' - 'え', 0}},
    unicode.CaseRange{'お', 'お', [unicode.MaxCase]rune{0, 'ぉ' - 'お', 0}},
    unicode.CaseRange{'か', 'か', [unicode.MaxCase]rune{0, 'ゕ' - 'か', 0}},
    unicode.CaseRange{'け', 'け', [unicode.MaxCase]rune{0, 'ゖ' - 'け', 0}},
    unicode.CaseRange{'つ', 'つ', [unicode.MaxCase]rune{0, 'っ' - 'つ', 0}},
    unicode.CaseRange{'や', 'や', [unicode.MaxCase]rune{0, 'ゃ' - 'や', 0}},
    unicode.CaseRange{'ゆ', 'ゆ', [unicode.MaxCase]rune{0, 'ゅ' - 'ゆ', 0}},
    unicode.CaseRange{'よ', 'よ', [unicode.MaxCase]rune{0, 'ょ' - 'よ', 0}},
    unicode.CaseRange{'わ', 'わ', [unicode.MaxCase]rune{0, 'ゎ' - 'わ', 0}},

    unicode.CaseRange{'ぁ', 'ぁ', [unicode.MaxCase]rune{'あ' - 'ぁ', 0, 0}},
    unicode.CaseRange{'ぃ', 'ぃ', [unicode.MaxCase]rune{'い' - 'ぃ', 0, 0}},
    unicode.CaseRange{'ぅ', 'ぅ', [unicode.MaxCase]rune{'う' - 'ぅ', 0, 0}},
    unicode.CaseRange{'ぇ', 'ぇ', [unicode.MaxCase]rune{'え' - 'ぇ', 0, 0}},
    unicode.CaseRange{'ぉ', 'ぉ', [unicode.MaxCase]rune{'お' - 'ぉ', 0, 0}},
    unicode.CaseRange{'ゕ', 'ゕ', [unicode.MaxCase]rune{'か' - 'ゕ', 0, 0}},
    unicode.CaseRange{'ゖ', 'ゖ', [unicode.MaxCase]rune{'け' - 'ゖ', 0, 0}},
    unicode.CaseRange{'っ', 'っ', [unicode.MaxCase]rune{'つ' - 'っ', 0, 0}},
    unicode.CaseRange{'ゃ', 'ゃ', [unicode.MaxCase]rune{'や' - 'ゃ', 0, 0}},
    unicode.CaseRange{'ゅ', 'ゅ', [unicode.MaxCase]rune{'ゆ' - 'ゅ', 0, 0}},
    unicode.CaseRange{'ょ', 'ょ', [unicode.MaxCase]rune{'よ' - 'ょ', 0, 0}},
    unicode.CaseRange{'ゎ', 'ゎ', [unicode.MaxCase]rune{'わ' - 'ゎ', 0, 0}},

    unicode.CaseRange{'ア', 'ア', [unicode.MaxCase]rune{0, 'ァ' - 'ア', 0}},
    unicode.CaseRange{'イ', 'イ', [unicode.MaxCase]rune{0, 'ィ' - 'イ', 0}},
    unicode.CaseRange{'ウ', 'ウ', [unicode.MaxCase]rune{0, 'ゥ' - 'ウ', 0}},
    unicode.CaseRange{'エ', 'エ', [unicode.MaxCase]rune{0, 'ェ' - 'エ', 0}},
    unicode.CaseRange{'オ', 'オ', [unicode.MaxCase]rune{0, 'ォ' - 'オ', 0}},
    unicode.CaseRange{'カ', 'カ', [unicode.MaxCase]rune{0, 'ヵ' - 'カ', 0}},
    unicode.CaseRange{'ケ', 'ケ', [unicode.MaxCase]rune{0, 'ヶ' - 'ケ', 0}},
    unicode.CaseRange{'ツ', 'ツ', [unicode.MaxCase]rune{0, 'ッ' - 'ツ', 0}},
    unicode.CaseRange{'ヤ', 'ヤ', [unicode.MaxCase]rune{0, 'ャ' - 'ヤ', 0}},
    unicode.CaseRange{'ユ', 'ユ', [unicode.MaxCase]rune{0, 'ュ' - 'ユ', 0}},
    unicode.CaseRange{'ヨ', 'ヨ', [unicode.MaxCase]rune{0, 'ョ' - 'ヨ', 0}},
    unicode.CaseRange{'ワ', 'ワ', [unicode.MaxCase]rune{0, 'ヮ' - 'ワ', 0}},

    unicode.CaseRange{'ァ', 'ァ', [unicode.MaxCase]rune{'ア' - 'ァ', 0, 0}},
    unicode.CaseRange{'ィ', 'ィ', [unicode.MaxCase]rune{'イ' - 'ィ', 0, 0}},
    unicode.CaseRange{'ゥ', 'ゥ', [unicode.MaxCase]rune{'ウ' - 'ゥ', 0, 0}},
    unicode.CaseRange{'ェ', 'ェ', [unicode.MaxCase]rune{'エ' - 'ェ', 0, 0}},
    unicode.CaseRange{'ォ', 'ォ', [unicode.MaxCase]rune{'オ' - 'ォ', 0, 0}},
    unicode.CaseRange{'ヵ', 'ヵ', [unicode.MaxCase]rune{'カ' - 'ヵ', 0, 0}},
    unicode.CaseRange{'ヶ', 'ヶ', [unicode.MaxCase]rune{'ケ' - 'ヶ', 0, 0}},
    unicode.CaseRange{'ッ', 'ッ', [unicode.MaxCase]rune{'ツ' - 'ッ', 0, 0}},
    unicode.CaseRange{'ャ', 'ャ', [unicode.MaxCase]rune{'ヤ' - 'ャ', 0, 0}},
    unicode.CaseRange{'ュ', 'ュ', [unicode.MaxCase]rune{'ユ' - 'ュ', 0, 0}},
    unicode.CaseRange{'ョ', 'ョ', [unicode.MaxCase]rune{'ヨ' - 'ョ', 0, 0}},
    unicode.CaseRange{'ヮ', 'ヮ', [unicode.MaxCase]rune{'ワ' - 'ヮ', 0, 0}},

    unicode.CaseRange{'ｱ', 'ｱ', [unicode.MaxCase]rune{0, 'ｧ' - 'ｱ', 0}},
    unicode.CaseRange{'ｲ', 'ｲ', [unicode.MaxCase]rune{0, 'ｨ' - 'ｲ', 0}},
    unicode.CaseRange{'ｳ', 'ｳ', [unicode.MaxCase]rune{0, 'ｩ' - 'ｳ', 0}},
    unicode.CaseRange{'ｴ', 'ｴ', [unicode.MaxCase]rune{0, 'ｪ' - 'ｴ', 0}},
    unicode.CaseRange{'ｵ', 'ｵ', [unicode.MaxCase]rune{0, 'ｫ' - 'ｵ', 0}},
    unicode.CaseRange{'ﾂ', 'ﾂ', [unicode.MaxCase]rune{0, 'ｯ' - 'ﾂ', 0}},
    unicode.CaseRange{'ﾔ', 'ﾔ', [unicode.MaxCase]rune{0, 'ｬ' - 'ﾔ', 0}},
    unicode.CaseRange{'ﾕ', 'ﾕ', [unicode.MaxCase]rune{0, 'ｭ' - 'ﾕ', 0}},
    unicode.CaseRange{'ﾖ', 'ﾖ', [unicode.MaxCase]rune{0, 'ｮ' - 'ﾖ', 0}},

    unicode.CaseRange{'ｧ', 'ｧ', [unicode.MaxCase]rune{'ｱ' - 'ｧ', 0, 0}},
    unicode.CaseRange{'ｨ', 'ｨ', [unicode.MaxCase]rune{'ｲ' - 'ｨ', 0, 0}},
    unicode.CaseRange{'ｩ', 'ｩ', [unicode.MaxCase]rune{'ｳ' - 'ｩ', 0, 0}},
    unicode.CaseRange{'ｪ', 'ｪ', [unicode.MaxCase]rune{'ｴ' - 'ｪ', 0, 0}},
    unicode.CaseRange{'ｫ', 'ｫ', [unicode.MaxCase]rune{'ｵ' - 'ｫ', 0, 0}},
    unicode.CaseRange{'ｯ', 'ｯ', [unicode.MaxCase]rune{'ﾂ' - 'ｯ', 0, 0}},
    unicode.CaseRange{'ｬ', 'ｬ', [unicode.MaxCase]rune{'ﾔ' - 'ｬ', 0, 0}},
    unicode.CaseRange{'ｭ', 'ｭ', [unicode.MaxCase]rune{'ﾕ' - 'ｭ', 0, 0}},
    unicode.CaseRange{'ｮ', 'ｮ', [unicode.MaxCase]rune{'ﾖ' - 'ｮ', 0, 0}},
}

func init() {
    sort.Slice(kanaCase2, func(i, j int) bool {
        return kanaCase2[i].Lo < kanaCase2[j].Lo
    })
}

func main() {
    kanaLower := "ぁぃぅぇぉゕゖっゃゅょゎ ァィゥェォヵヶッャュョヮ ｧｨｩｪｫｯｬｭｮ"
    kanaUpper := strings.ToUpperSpecial(kanaCase2, kanaLower)
    fmt.Println(kanaLower)
    fmt.Println(kanaUpper)
    fmt.Println(strings.ToLowerSpecial(kanaCase2, kanaUpper))
}
```

これを実行すると

```text
$ go run sample3.go 
ぁぃぅぇぉゕゖっゃゅょゎ ァィゥェォヵヶッャュョヮ ｧｨｩｪｫｯｬｭｮ
あいうえおかけつやゆよわ アイウエオカケツヤユヨワ ｱｲｳｴｵﾂﾔﾕﾖ
ぁぃぅぇぉゕゖっゃゅょゎ ァィゥェォヵヶッャュョヮ ｧｨｩｪｫｯｬｭｮ
```

となる。

## ブックマーク

- [Go言語で文字列の変換(全角・半角、ひらがな・カタカナ)をする  :  Serendip – Webデザイン・プログラミング](https://www.serendip.ws/archives/6307)
- [Go 言語と Unicode 正規化]({{< relref "./unicode-normalization.md" >}})
- [Go 言語による Unicode 半角/全角変換]({{< relref "./character-width-for-unicode.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[`norm`]: https://pkg.go.dev/golang.org/x/text/unicode/norm "norm package · pkg.go.dev"
[`width`]: https://pkg.go.dev/golang.org/x/text/width "width package · pkg.go.dev"
[`transform`]: https://pkg.go.dev/golang.org/x/text/transform "transform package · pkg.go.dev"
[`strings`]: https://pkg.go.dev/strings "strings package · pkg.go.dev"
[`unicode`]: https://pkg.go.dev/unicode "unicode package · pkg.go.dev"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
