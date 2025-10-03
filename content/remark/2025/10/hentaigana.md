+++
title = "変体仮名をサポートした"
date =  "2025-10-03T13:08:59+09:00"
description = "このブログサイトで変体仮名を表示できるようにしてみた。"
image = "/images/attention/tools.png"
tags = [ "site", "web", "font", "character", "unicode" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

昨年の古い話題になるのだが Google NOTO フォントが変体仮名[^h1] をリリースした。

[^h1]: 変体仮名とは，簡単に言うと仮名の異体字のこと。変体仮名は明治時代に廃止されたが，店の看板などで依然使われている。コンピュータ・システムでは，地名や古い戸籍の人名などで使われていることがあるそうで「住民基本台帳収録変体仮名」として一部の変体仮名が運用されていたらしい。その後 Unicode 10 から本格的に変体仮名がサポートされるようになった。フォントとしては「[IPAmj明朝フォント](https://moji.or.jp/mojikiban/font/ "IPAmj明朝フォント | 一般社団法人 文字情報技術促進協議会")」が変体仮名を含む異体字を幅広くカバーしている。

- [Noto Serif Hentaigana - Google Fonts](https://fonts.google.com/noto/specimen/Noto+Serif+Hentaigana)
- [Googleが変体仮名フォント「Noto Hentaigana」をリリース ～蕎麦屋の看板などを再現可能 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1581369.html)

というわけで変体仮名を Web フォントとして組み込むことができる。
このサイトでの設定はこんな感じ。

```html
<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link href="https://fonts.googleapis.com/css2?family=BIZ+UDGothic:wght@400;700&family=BIZ+UDMincho:wght@400;700&family=Intel+One+Mono:wght@400;700&family=Noto+Color+Emoji&family=Noto+Sans:wght@400;700&family=Noto+Serif+Hentaigana:wght@400;700&family=Noto+Serif:wght@400;700&display=swap" rel="stylesheet">
```

これで CSS の `font-family` に `'Noto Serif Hentaigana'` を追加すればよい。

```css
body {
  font-family: 'Noto Serif', 'BIZ UDMincho', 'Noto Serif Hentaigana', serif;
}
```

ちょっと表示させてみよう。
以前「[Unicode のカタカナを調べる]({{< ref "/golang/unicode-katakana/" >}})」で書いたコードを流用して Unicode で変体仮名を含む仮名補助（Kana Supplement; `U+1B000` - `U+1B0FF`）および仮名拡張A（Kana Extended-A; `U+1B100` – `U+1B12F`）を列挙してみる。
こんな感じ。

```go
package main

import (
    "fmt"
    "unicode"
)

func main() {
    i := 0
    for c := rune(0x1b000); c <= rune(0x1b12f); c++ {
        if unicode.In(c, unicode.Hiragana) || unicode.In(c, unicode.Katakana) {
            fmt.Printf("%#U ", c)
            i++
            i &= 0x0f
            if i == 0 {
                fmt.Println()
            }
        }
    }
    fmt.Println()
}
```

実行結果は以下の通り。

```text
$ go run main.go
U+1B000 '𛀀' U+1B001 '𛀁' U+1B002 '𛀂' U+1B003 '𛀃' U+1B004 '𛀄' U+1B005 '𛀅' U+1B006 '𛀆' U+1B007 '𛀇' U+1B008 '𛀈' U+1B009 '𛀉' U+1B00A '𛀊' U+1B00B '𛀋' U+1B00C '𛀌' U+1B00D '𛀍' U+1B00E '𛀎' U+1B00F '𛀏' 
U+1B010 '𛀐' U+1B011 '𛀑' U+1B012 '𛀒' U+1B013 '𛀓' U+1B014 '𛀔' U+1B015 '𛀕' U+1B016 '𛀖' U+1B017 '𛀗' U+1B018 '𛀘' U+1B019 '𛀙' U+1B01A '𛀚' U+1B01B '𛀛' U+1B01C '𛀜' U+1B01D '𛀝' U+1B01E '𛀞' U+1B01F '𛀟' 
U+1B020 '𛀠' U+1B021 '𛀡' U+1B022 '𛀢' U+1B023 '𛀣' U+1B024 '𛀤' U+1B025 '𛀥' U+1B026 '𛀦' U+1B027 '𛀧' U+1B028 '𛀨' U+1B029 '𛀩' U+1B02A '𛀪' U+1B02B '𛀫' U+1B02C '𛀬' U+1B02D '𛀭' U+1B02E '𛀮' U+1B02F '𛀯' 
U+1B030 '𛀰' U+1B031 '𛀱' U+1B032 '𛀲' U+1B033 '𛀳' U+1B034 '𛀴' U+1B035 '𛀵' U+1B036 '𛀶' U+1B037 '𛀷' U+1B038 '𛀸' U+1B039 '𛀹' U+1B03A '𛀺' U+1B03B '𛀻' U+1B03C '𛀼' U+1B03D '𛀽' U+1B03E '𛀾' U+1B03F '𛀿' 
U+1B040 '𛁀' U+1B041 '𛁁' U+1B042 '𛁂' U+1B043 '𛁃' U+1B044 '𛁄' U+1B045 '𛁅' U+1B046 '𛁆' U+1B047 '𛁇' U+1B048 '𛁈' U+1B049 '𛁉' U+1B04A '𛁊' U+1B04B '𛁋' U+1B04C '𛁌' U+1B04D '𛁍' U+1B04E '𛁎' U+1B04F '𛁏' 
U+1B050 '𛁐' U+1B051 '𛁑' U+1B052 '𛁒' U+1B053 '𛁓' U+1B054 '𛁔' U+1B055 '𛁕' U+1B056 '𛁖' U+1B057 '𛁗' U+1B058 '𛁘' U+1B059 '𛁙' U+1B05A '𛁚' U+1B05B '𛁛' U+1B05C '𛁜' U+1B05D '𛁝' U+1B05E '𛁞' U+1B05F '𛁟' 
U+1B060 '𛁠' U+1B061 '𛁡' U+1B062 '𛁢' U+1B063 '𛁣' U+1B064 '𛁤' U+1B065 '𛁥' U+1B066 '𛁦' U+1B067 '𛁧' U+1B068 '𛁨' U+1B069 '𛁩' U+1B06A '𛁪' U+1B06B '𛁫' U+1B06C '𛁬' U+1B06D '𛁭' U+1B06E '𛁮' U+1B06F '𛁯' 
U+1B070 '𛁰' U+1B071 '𛁱' U+1B072 '𛁲' U+1B073 '𛁳' U+1B074 '𛁴' U+1B075 '𛁵' U+1B076 '𛁶' U+1B077 '𛁷' U+1B078 '𛁸' U+1B079 '𛁹' U+1B07A '𛁺' U+1B07B '𛁻' U+1B07C '𛁼' U+1B07D '𛁽' U+1B07E '𛁾' U+1B07F '𛁿' 
U+1B080 '𛂀' U+1B081 '𛂁' U+1B082 '𛂂' U+1B083 '𛂃' U+1B084 '𛂄' U+1B085 '𛂅' U+1B086 '𛂆' U+1B087 '𛂇' U+1B088 '𛂈' U+1B089 '𛂉' U+1B08A '𛂊' U+1B08B '𛂋' U+1B08C '𛂌' U+1B08D '𛂍' U+1B08E '𛂎' U+1B08F '𛂏' 
U+1B090 '𛂐' U+1B091 '𛂑' U+1B092 '𛂒' U+1B093 '𛂓' U+1B094 '𛂔' U+1B095 '𛂕' U+1B096 '𛂖' U+1B097 '𛂗' U+1B098 '𛂘' U+1B099 '𛂙' U+1B09A '𛂚' U+1B09B '𛂛' U+1B09C '𛂜' U+1B09D '𛂝' U+1B09E '𛂞' U+1B09F '𛂟' 
U+1B0A0 '𛂠' U+1B0A1 '𛂡' U+1B0A2 '𛂢' U+1B0A3 '𛂣' U+1B0A4 '𛂤' U+1B0A5 '𛂥' U+1B0A6 '𛂦' U+1B0A7 '𛂧' U+1B0A8 '𛂨' U+1B0A9 '𛂩' U+1B0AA '𛂪' U+1B0AB '𛂫' U+1B0AC '𛂬' U+1B0AD '𛂭' U+1B0AE '𛂮' U+1B0AF '𛂯' 
U+1B0B0 '𛂰' U+1B0B1 '𛂱' U+1B0B2 '𛂲' U+1B0B3 '𛂳' U+1B0B4 '𛂴' U+1B0B5 '𛂵' U+1B0B6 '𛂶' U+1B0B7 '𛂷' U+1B0B8 '𛂸' U+1B0B9 '𛂹' U+1B0BA '𛂺' U+1B0BB '𛂻' U+1B0BC '𛂼' U+1B0BD '𛂽' U+1B0BE '𛂾' U+1B0BF '𛂿' 
U+1B0C0 '𛃀' U+1B0C1 '𛃁' U+1B0C2 '𛃂' U+1B0C3 '𛃃' U+1B0C4 '𛃄' U+1B0C5 '𛃅' U+1B0C6 '𛃆' U+1B0C7 '𛃇' U+1B0C8 '𛃈' U+1B0C9 '𛃉' U+1B0CA '𛃊' U+1B0CB '𛃋' U+1B0CC '𛃌' U+1B0CD '𛃍' U+1B0CE '𛃎' U+1B0CF '𛃏' 
U+1B0D0 '𛃐' U+1B0D1 '𛃑' U+1B0D2 '𛃒' U+1B0D3 '𛃓' U+1B0D4 '𛃔' U+1B0D5 '𛃕' U+1B0D6 '𛃖' U+1B0D7 '𛃗' U+1B0D8 '𛃘' U+1B0D9 '𛃙' U+1B0DA '𛃚' U+1B0DB '𛃛' U+1B0DC '𛃜' U+1B0DD '𛃝' U+1B0DE '𛃞' U+1B0DF '𛃟' 
U+1B0E0 '𛃠' U+1B0E1 '𛃡' U+1B0E2 '𛃢' U+1B0E3 '𛃣' U+1B0E4 '𛃤' U+1B0E5 '𛃥' U+1B0E6 '𛃦' U+1B0E7 '𛃧' U+1B0E8 '𛃨' U+1B0E9 '𛃩' U+1B0EA '𛃪' U+1B0EB '𛃫' U+1B0EC '𛃬' U+1B0ED '𛃭' U+1B0EE '𛃮' U+1B0EF '𛃯' 
U+1B0F0 '𛃰' U+1B0F1 '𛃱' U+1B0F2 '𛃲' U+1B0F3 '𛃳' U+1B0F4 '𛃴' U+1B0F5 '𛃵' U+1B0F6 '𛃶' U+1B0F7 '𛃷' U+1B0F8 '𛃸' U+1B0F9 '𛃹' U+1B0FA '𛃺' U+1B0FB '𛃻' U+1B0FC '𛃼' U+1B0FD '𛃽' U+1B0FE '𛃾' U+1B0FF '𛃿' 
U+1B100 '𛄀' U+1B101 '𛄁' U+1B102 '𛄂' U+1B103 '𛄃' U+1B104 '𛄄' U+1B105 '𛄅' U+1B106 '𛄆' U+1B107 '𛄇' U+1B108 '𛄈' U+1B109 '𛄉' U+1B10A '𛄊' U+1B10B '𛄋' U+1B10C '𛄌' U+1B10D '𛄍' U+1B10E '𛄎' U+1B10F '𛄏' 
U+1B110 '𛄐' U+1B111 '𛄑' U+1B112 '𛄒' U+1B113 '𛄓' U+1B114 '𛄔' U+1B115 '𛄕' U+1B116 '𛄖' U+1B117 '𛄗' U+1B118 '𛄘' U+1B119 '𛄙' U+1B11A '𛄚' U+1B11B '𛄛' U+1B11C '𛄜' U+1B11D '𛄝' U+1B11E '𛄞' U+1B11F '𛄟' 
U+1B120 '𛄠' U+1B121 '𛄡' U+1B122 '𛄢'
```

皆さんの環境では見えるだろうか。

ちなみに `U+1B000` (&#x1B000;) および `U+1B001` (&#x1B001;) は Unicode 6 で “Historic Katakana” および “Historic Hiragana” として収録されたそうで，これまでの NOTO 日本語フォント等で表示できる。
また `U+1B120` (&#x1B120;), `U+1B121` (&#x1B121;), `U+1B122` (&#x1B122;) は “Historic Katakana” として Unicode 14 で追加されたそうだ。

同じく Unicode 14 で追加された `U+1B11F` は表示されないのか？ いや `U+1B11F` は “Historic Hiragana” に分類されているから “Hentaigana” ではないという判定なのかな？ でもでも，上述の “Historic Katakana” は `U+1B122` まで表示できるんだよな。
うーむ，よく分からん。

まっ，細かいことは置いておいて，例として蕎麦屋の看板や暖簾に書かれている変体仮名「&#x1B05B;&#x1B0A6;&#x3099;」を挙げてみる。
あれって元は「楚（そ）」「者（は）」なんだそうで

楚 → &#x1B05B; (`U+1B05B`)<br>
者 → &#x1B0A6;&#x3099; (`U+1B0A6` + `U+3099`; 濁点が付いてる)<br>

ってな感じらしい。
勉強になりました。

「[変体仮名一覧]({{< ref "/hentaigana-table.md" >}})」に変体仮名の各文字に対応する読みと元になる漢字を表にまとめてみた。
参考にどうぞ。
&#x1b099;（乃）や &#x1b0c0;（本）といった文字は今でもたまに見かけるな。

## ブックマーク

- {{< pdf-file title="Kana Supplement Range: 1B000–1B0FF" link="https://www.unicode.org/charts/PDF/U1B000.pdf" >}}
- {{< pdf-file title="Kana Extended-A Range: 1B100–1B12F" link="https://www.unicode.org/charts/PDF/U1B100.pdf" >}}
- [変体仮名 - Wikipedia](https://ja.wikipedia.org/wiki/%E5%A4%89%E4%BD%93%E4%BB%AE%E5%90%8D)
- [蕎麦屋の暖簾に書いてあるあの難しい字・・・なにあれ？＝Sobapedia＝ | 出張そば打ち体験～SOBAUCHI 楽常～](https://rakujyo.com/blog/sobaya-anoji/)

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
