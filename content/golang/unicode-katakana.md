+++
title = "Unicode のカタカナを調べる"
date =  "2021-09-18T13:15:18+09:00"
description = "unicode.Letter で絞り込めば記号類は除外できる"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "unicode", "character" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Qiita で

- [Goでカタカナバリデーションを実装しようとしたが、unicodeパッケージのKatakanaテーブルには記号なども含まれていた。 - Qiita](https://qiita.com/mitsuaki1229/items/50c6da607b62f138c32f)

という記事を見かけたので，自分でも調べてみる。

上の記事にもある通り [`unicode`]`.Katakana` は以下の範囲で定義されている。

```go
var _Katakana = &RangeTable{
	R16: []Range16{
		{0x30a1, 0x30fa, 1},
		{0x30fd, 0x30ff, 1},
		{0x31f0, 0x31ff, 1},
		{0x32d0, 0x32fe, 1},
		{0x3300, 0x3357, 1},
		{0xff66, 0xff6f, 1},
		{0xff71, 0xff9d, 1},
	},
	R32: []Range32{
		{0x1b000, 0x1b164, 356},
		{0x1b165, 0x1b167, 1},
	},
}
```

面倒なので `0x30a1` から `0x1b167` まで片っ端からスキャンしてしまおう（笑）

```go
//go:build run
// +build run

package main

import (
	"fmt"
	"unicode"
)

func main() {
	i := 0
	for c := rune(0x30a1); c <= rune(0x1b167); c++ {
		if unicode.In(c, unicode.Katakana) {
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

これを実行すると以下の通り。

```text
$ go run sample1.go
U+30A1 'ァ' U+30A2 'ア' U+30A3 'ィ' U+30A4 'イ' U+30A5 'ゥ' U+30A6 'ウ' U+30A7 'ェ' U+30A8 'エ' U+30A9 'ォ' U+30AA 'オ' U+30AB 'カ' U+30AC 'ガ' U+30AD 'キ' U+30AE 'ギ' U+30AF 'ク' U+30B0 'グ' 
U+30B1 'ケ' U+30B2 'ゲ' U+30B3 'コ' U+30B4 'ゴ' U+30B5 'サ' U+30B6 'ザ' U+30B7 'シ' U+30B8 'ジ' U+30B9 'ス' U+30BA 'ズ' U+30BB 'セ' U+30BC 'ゼ' U+30BD 'ソ' U+30BE 'ゾ' U+30BF 'タ' U+30C0 'ダ' 
U+30C1 'チ' U+30C2 'ヂ' U+30C3 'ッ' U+30C4 'ツ' U+30C5 'ヅ' U+30C6 'テ' U+30C7 'デ' U+30C8 'ト' U+30C9 'ド' U+30CA 'ナ' U+30CB 'ニ' U+30CC 'ヌ' U+30CD 'ネ' U+30CE 'ノ' U+30CF 'ハ' U+30D0 'バ' 
U+30D1 'パ' U+30D2 'ヒ' U+30D3 'ビ' U+30D4 'ピ' U+30D5 'フ' U+30D6 'ブ' U+30D7 'プ' U+30D8 'ヘ' U+30D9 'ベ' U+30DA 'ペ' U+30DB 'ホ' U+30DC 'ボ' U+30DD 'ポ' U+30DE 'マ' U+30DF 'ミ' U+30E0 'ム' 
U+30E1 'メ' U+30E2 'モ' U+30E3 'ャ' U+30E4 'ヤ' U+30E5 'ュ' U+30E6 'ユ' U+30E7 'ョ' U+30E8 'ヨ' U+30E9 'ラ' U+30EA 'リ' U+30EB 'ル' U+30EC 'レ' U+30ED 'ロ' U+30EE 'ヮ' U+30EF 'ワ' U+30F0 'ヰ' 
U+30F1 'ヱ' U+30F2 'ヲ' U+30F3 'ン' U+30F4 'ヴ' U+30F5 'ヵ' U+30F6 'ヶ' U+30F7 'ヷ' U+30F8 'ヸ' U+30F9 'ヹ' U+30FA 'ヺ' U+30FD 'ヽ' U+30FE 'ヾ' U+30FF 'ヿ' U+31F0 'ㇰ' U+31F1 'ㇱ' U+31F2 'ㇲ' 
U+31F3 'ㇳ' U+31F4 'ㇴ' U+31F5 'ㇵ' U+31F6 'ㇶ' U+31F7 'ㇷ' U+31F8 'ㇸ' U+31F9 'ㇹ' U+31FA 'ㇺ' U+31FB 'ㇻ' U+31FC 'ㇼ' U+31FD 'ㇽ' U+31FE 'ㇾ' U+31FF 'ㇿ' U+32D0 '㋐' U+32D1 '㋑' U+32D2 '㋒' 
U+32D3 '㋓' U+32D4 '㋔' U+32D5 '㋕' U+32D6 '㋖' U+32D7 '㋗' U+32D8 '㋘' U+32D9 '㋙' U+32DA '㋚' U+32DB '㋛' U+32DC '㋜' U+32DD '㋝' U+32DE '㋞' U+32DF '㋟' U+32E0 '㋠' U+32E1 '㋡' U+32E2 '㋢' 
U+32E3 '㋣' U+32E4 '㋤' U+32E5 '㋥' U+32E6 '㋦' U+32E7 '㋧' U+32E8 '㋨' U+32E9 '㋩' U+32EA '㋪' U+32EB '㋫' U+32EC '㋬' U+32ED '㋭' U+32EE '㋮' U+32EF '㋯' U+32F0 '㋰' U+32F1 '㋱' U+32F2 '㋲' 
U+32F3 '㋳' U+32F4 '㋴' U+32F5 '㋵' U+32F6 '㋶' U+32F7 '㋷' U+32F8 '㋸' U+32F9 '㋹' U+32FA '㋺' U+32FB '㋻' U+32FC '㋼' U+32FD '㋽' U+32FE '㋾' U+3300 '㌀' U+3301 '㌁' U+3302 '㌂' U+3303 '㌃' 
U+3304 '㌄' U+3305 '㌅' U+3306 '㌆' U+3307 '㌇' U+3308 '㌈' U+3309 '㌉' U+330A '㌊' U+330B '㌋' U+330C '㌌' U+330D '㌍' U+330E '㌎' U+330F '㌏' U+3310 '㌐' U+3311 '㌑' U+3312 '㌒' U+3313 '㌓' 
U+3314 '㌔' U+3315 '㌕' U+3316 '㌖' U+3317 '㌗' U+3318 '㌘' U+3319 '㌙' U+331A '㌚' U+331B '㌛' U+331C '㌜' U+331D '㌝' U+331E '㌞' U+331F '㌟' U+3320 '㌠' U+3321 '㌡' U+3322 '㌢' U+3323 '㌣' 
U+3324 '㌤' U+3325 '㌥' U+3326 '㌦' U+3327 '㌧' U+3328 '㌨' U+3329 '㌩' U+332A '㌪' U+332B '㌫' U+332C '㌬' U+332D '㌭' U+332E '㌮' U+332F '㌯' U+3330 '㌰' U+3331 '㌱' U+3332 '㌲' U+3333 '㌳' 
U+3334 '㌴' U+3335 '㌵' U+3336 '㌶' U+3337 '㌷' U+3338 '㌸' U+3339 '㌹' U+333A '㌺' U+333B '㌻' U+333C '㌼' U+333D '㌽' U+333E '㌾' U+333F '㌿' U+3340 '㍀' U+3341 '㍁' U+3342 '㍂' U+3343 '㍃' 
U+3344 '㍄' U+3345 '㍅' U+3346 '㍆' U+3347 '㍇' U+3348 '㍈' U+3349 '㍉' U+334A '㍊' U+334B '㍋' U+334C '㍌' U+334D '㍍' U+334E '㍎' U+334F '㍏' U+3350 '㍐' U+3351 '㍑' U+3352 '㍒' U+3353 '㍓' 
U+3354 '㍔' U+3355 '㍕' U+3356 '㍖' U+3357 '㍗' U+FF66 'ｦ' U+FF67 'ｧ' U+FF68 'ｨ' U+FF69 'ｩ' U+FF6A 'ｪ' U+FF6B 'ｫ' U+FF6C 'ｬ' U+FF6D 'ｭ' U+FF6E 'ｮ' U+FF6F 'ｯ' U+FF71 'ｱ' U+FF72 'ｲ' 
U+FF73 'ｳ' U+FF74 'ｴ' U+FF75 'ｵ' U+FF76 'ｶ' U+FF77 'ｷ' U+FF78 'ｸ' U+FF79 'ｹ' U+FF7A 'ｺ' U+FF7B 'ｻ' U+FF7C 'ｼ' U+FF7D 'ｽ' U+FF7E 'ｾ' U+FF7F 'ｿ' U+FF80 'ﾀ' U+FF81 'ﾁ' U+FF82 'ﾂ' 
U+FF83 'ﾃ' U+FF84 'ﾄ' U+FF85 'ﾅ' U+FF86 'ﾆ' U+FF87 'ﾇ' U+FF88 'ﾈ' U+FF89 'ﾉ' U+FF8A 'ﾊ' U+FF8B 'ﾋ' U+FF8C 'ﾌ' U+FF8D 'ﾍ' U+FF8E 'ﾎ' U+FF8F 'ﾏ' U+FF90 'ﾐ' U+FF91 'ﾑ' U+FF92 'ﾒ' 
U+FF93 'ﾓ' U+FF94 'ﾔ' U+FF95 'ﾕ' U+FF96 'ﾖ' U+FF97 'ﾗ' U+FF98 'ﾘ' U+FF99 'ﾙ' U+FF9A 'ﾚ' U+FF9B 'ﾛ' U+FF9C 'ﾜ' U+FF9D 'ﾝ' U+1B000 '𛀀' U+1B164 '𛅤' U+1B165 '𛅥' U+1B166 '𛅦' U+1B167 '𛅧' 
```

まぢか。
これは困るわ。

`U+1B000` は変体仮名の「れ（連）」を指す文字らしい。
これだけカタカナ扱いなのか。
また `U+1B164` 以降の4文字は小書きの「ヰ」「ヱ」「ヲ」「ン」だそうな。

じゃあ，次。
上のカタカナを [`unicode`]`.Letter` 種別で絞り込んだらどうなるか。

```go {hl_lines=[14]}
//go:build run
// +build run

package main

import (
	"fmt"
	"unicode"
)

func main() {
	i := 0
	for c := rune(0x30a1); c <= rune(0x1b167); c++ {
		if unicode.In(c, unicode.Katakana) && unicode.IsLetter(c) {
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

これを実行すると以下の通り。

```text
$ go run sample2.go
U+30A1 'ァ' U+30A2 'ア' U+30A3 'ィ' U+30A4 'イ' U+30A5 'ゥ' U+30A6 'ウ' U+30A7 'ェ' U+30A8 'エ' U+30A9 'ォ' U+30AA 'オ' U+30AB 'カ' U+30AC 'ガ' U+30AD 'キ' U+30AE 'ギ' U+30AF 'ク' U+30B0 'グ' 
U+30B1 'ケ' U+30B2 'ゲ' U+30B3 'コ' U+30B4 'ゴ' U+30B5 'サ' U+30B6 'ザ' U+30B7 'シ' U+30B8 'ジ' U+30B9 'ス' U+30BA 'ズ' U+30BB 'セ' U+30BC 'ゼ' U+30BD 'ソ' U+30BE 'ゾ' U+30BF 'タ' U+30C0 'ダ' 
U+30C1 'チ' U+30C2 'ヂ' U+30C3 'ッ' U+30C4 'ツ' U+30C5 'ヅ' U+30C6 'テ' U+30C7 'デ' U+30C8 'ト' U+30C9 'ド' U+30CA 'ナ' U+30CB 'ニ' U+30CC 'ヌ' U+30CD 'ネ' U+30CE 'ノ' U+30CF 'ハ' U+30D0 'バ' 
U+30D1 'パ' U+30D2 'ヒ' U+30D3 'ビ' U+30D4 'ピ' U+30D5 'フ' U+30D6 'ブ' U+30D7 'プ' U+30D8 'ヘ' U+30D9 'ベ' U+30DA 'ペ' U+30DB 'ホ' U+30DC 'ボ' U+30DD 'ポ' U+30DE 'マ' U+30DF 'ミ' U+30E0 'ム' 
U+30E1 'メ' U+30E2 'モ' U+30E3 'ャ' U+30E4 'ヤ' U+30E5 'ュ' U+30E6 'ユ' U+30E7 'ョ' U+30E8 'ヨ' U+30E9 'ラ' U+30EA 'リ' U+30EB 'ル' U+30EC 'レ' U+30ED 'ロ' U+30EE 'ヮ' U+30EF 'ワ' U+30F0 'ヰ' 
U+30F1 'ヱ' U+30F2 'ヲ' U+30F3 'ン' U+30F4 'ヴ' U+30F5 'ヵ' U+30F6 'ヶ' U+30F7 'ヷ' U+30F8 'ヸ' U+30F9 'ヹ' U+30FA 'ヺ' U+30FD 'ヽ' U+30FE 'ヾ' U+30FF 'ヿ' U+31F0 'ㇰ' U+31F1 'ㇱ' U+31F2 'ㇲ' 
U+31F3 'ㇳ' U+31F4 'ㇴ' U+31F5 'ㇵ' U+31F6 'ㇶ' U+31F7 'ㇷ' U+31F8 'ㇸ' U+31F9 'ㇹ' U+31FA 'ㇺ' U+31FB 'ㇻ' U+31FC 'ㇼ' U+31FD 'ㇽ' U+31FE 'ㇾ' U+31FF 'ㇿ' U+FF66 'ｦ' U+FF67 'ｧ' U+FF68 'ｨ' 
U+FF69 'ｩ' U+FF6A 'ｪ' U+FF6B 'ｫ' U+FF6C 'ｬ' U+FF6D 'ｭ' U+FF6E 'ｮ' U+FF6F 'ｯ' U+FF71 'ｱ' U+FF72 'ｲ' U+FF73 'ｳ' U+FF74 'ｴ' U+FF75 'ｵ' U+FF76 'ｶ' U+FF77 'ｷ' U+FF78 'ｸ' U+FF79 'ｹ' 
U+FF7A 'ｺ' U+FF7B 'ｻ' U+FF7C 'ｼ' U+FF7D 'ｽ' U+FF7E 'ｾ' U+FF7F 'ｿ' U+FF80 'ﾀ' U+FF81 'ﾁ' U+FF82 'ﾂ' U+FF83 'ﾃ' U+FF84 'ﾄ' U+FF85 'ﾅ' U+FF86 'ﾆ' U+FF87 'ﾇ' U+FF88 'ﾈ' U+FF89 'ﾉ' 
U+FF8A 'ﾊ' U+FF8B 'ﾋ' U+FF8C 'ﾌ' U+FF8D 'ﾍ' U+FF8E 'ﾎ' U+FF8F 'ﾏ' U+FF90 'ﾐ' U+FF91 'ﾑ' U+FF92 'ﾒ' U+FF93 'ﾓ' U+FF94 'ﾔ' U+FF95 'ﾕ' U+FF96 'ﾖ' U+FF97 'ﾗ' U+FF98 'ﾘ' U+FF99 'ﾙ' 
U+FF9A 'ﾚ' U+FF9B 'ﾛ' U+FF9C 'ﾜ' U+FF9D 'ﾝ' U+1B000 '𛀀' U+1B164 '𛅤' U+1B165 '𛅥' U+1B166 '𛅦' U+1B167 '𛅧' 
```

これで記号はちゃんと排除できることが分かった。
よーし，うむうむ，よーし。

注意点としては [`unicode`]`.Katakana` では濁点・半濁点の結合文字が除外されるということだろうか。
あと半角カナを識別する [`unicode`]`.RangeTable` テーブルはなさそうだ。
まぁ，単純だから自前で作ればいいんだけど。
こんな感じかな。

```go {hl_lines=["11-16", 21]}
//go:build run
// +build run

package main

import (
	"fmt"
	"unicode"
)

var halfWidthKatakana = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0xff61, 0xff9f, 1},
	},
	R32: []unicode.Range32{},
}

func main() {
	i := 0
	for c := rune(0x30a1); c <= rune(0x1b167); c++ {
		if unicode.In(c, unicode.Katakana) && unicode.In(c, halfWidthKatakana) {
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

`halfWidthKatakana` は JIS X 0201 の JIS カナ領域を（記号を含めて）全てカバーしている点に注意。
上のコードのように，これと [`unicode`]`.Katakana` を組み合わせれば

```text
$ go run sample3.go
U+FF66 'ｦ' U+FF67 'ｧ' U+FF68 'ｨ' U+FF69 'ｩ' U+FF6A 'ｪ' U+FF6B 'ｫ' U+FF6C 'ｬ' U+FF6D 'ｭ' U+FF6E 'ｮ' U+FF6F 'ｯ' U+FF71 'ｱ' U+FF72 'ｲ' U+FF73 'ｳ' U+FF74 'ｴ' U+FF75 'ｵ' U+FF76 'ｶ' 
U+FF77 'ｷ' U+FF78 'ｸ' U+FF79 'ｹ' U+FF7A 'ｺ' U+FF7B 'ｻ' U+FF7C 'ｼ' U+FF7D 'ｽ' U+FF7E 'ｾ' U+FF7F 'ｿ' U+FF80 'ﾀ' U+FF81 'ﾁ' U+FF82 'ﾂ' U+FF83 'ﾃ' U+FF84 'ﾄ' U+FF85 'ﾅ' U+FF86 'ﾆ' 
U+FF87 'ﾇ' U+FF88 'ﾈ' U+FF89 'ﾉ' U+FF8A 'ﾊ' U+FF8B 'ﾋ' U+FF8C 'ﾌ' U+FF8D 'ﾍ' U+FF8E 'ﾎ' U+FF8F 'ﾏ' U+FF90 'ﾐ' U+FF91 'ﾑ' U+FF92 'ﾒ' U+FF93 'ﾓ' U+FF94 'ﾔ' U+FF95 'ﾕ' U+FF96 'ﾖ' 
U+FF97 'ﾗ' U+FF98 'ﾘ' U+FF99 'ﾙ' U+FF9A 'ﾚ' U+FF9B 'ﾛ' U+FF9C 'ﾜ' U+FF9D 'ﾝ' 
```

といい感じに半角カナ文字だけ取り出せる。

ついでに {{% span lang="en" %}}NFKC (Normalization Form Compatibility Composition){{% /span %}} 正規化で記号がどのように変換されるか見てみよう。
以下は変換されるコードのみを列挙したもの。

```go
//go:build run
// +build run

package main

import (
	"fmt"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func main() {
	for c := rune(0x30a1); c <= rune(0x1b167); c++ {
		if unicode.In(c, unicode.Katakana) {
			n := norm.NFKC.String(string(c))
			if n != string(c) {
				fmt.Printf("%#U -> %s\n", c, n)
			}
		}
	}
}
```

これを実行する。
ちょっと多いよ。

```text
$ go run sample4.go
U+30FF 'ヿ' -> コト
U+32D0 '㋐' -> ア
U+32D1 '㋑' -> イ
U+32D2 '㋒' -> ウ
U+32D3 '㋓' -> エ
U+32D4 '㋔' -> オ
U+32D5 '㋕' -> カ
U+32D6 '㋖' -> キ
U+32D7 '㋗' -> ク
U+32D8 '㋘' -> ケ
U+32D9 '㋙' -> コ
U+32DA '㋚' -> サ
U+32DB '㋛' -> シ
U+32DC '㋜' -> ス
U+32DD '㋝' -> セ
U+32DE '㋞' -> ソ
U+32DF '㋟' -> タ
U+32E0 '㋠' -> チ
U+32E1 '㋡' -> ツ
U+32E2 '㋢' -> テ
U+32E3 '㋣' -> ト
U+32E4 '㋤' -> ナ
U+32E5 '㋥' -> ニ
U+32E6 '㋦' -> ヌ
U+32E7 '㋧' -> ネ
U+32E8 '㋨' -> ノ
U+32E9 '㋩' -> ハ
U+32EA '㋪' -> ヒ
U+32EB '㋫' -> フ
U+32EC '㋬' -> ヘ
U+32ED '㋭' -> ホ
U+32EE '㋮' -> マ
U+32EF '㋯' -> ミ
U+32F0 '㋰' -> ム
U+32F1 '㋱' -> メ
U+32F2 '㋲' -> モ
U+32F3 '㋳' -> ヤ
U+32F4 '㋴' -> ユ
U+32F5 '㋵' -> ヨ
U+32F6 '㋶' -> ラ
U+32F7 '㋷' -> リ
U+32F8 '㋸' -> ル
U+32F9 '㋹' -> レ
U+32FA '㋺' -> ロ
U+32FB '㋻' -> ワ
U+32FC '㋼' -> ヰ
U+32FD '㋽' -> ヱ
U+32FE '㋾' -> ヲ
U+3300 '㌀' -> アパート
U+3301 '㌁' -> アルファ
U+3302 '㌂' -> アンペア
U+3303 '㌃' -> アール
U+3304 '㌄' -> イニング
U+3305 '㌅' -> インチ
U+3306 '㌆' -> ウォン
U+3307 '㌇' -> エスクード
U+3308 '㌈' -> エーカー
U+3309 '㌉' -> オンス
U+330A '㌊' -> オーム
U+330B '㌋' -> カイリ
U+330C '㌌' -> カラット
U+330D '㌍' -> カロリー
U+330E '㌎' -> ガロン
U+330F '㌏' -> ガンマ
U+3310 '㌐' -> ギガ
U+3311 '㌑' -> ギニー
U+3312 '㌒' -> キュリー
U+3313 '㌓' -> ギルダー
U+3314 '㌔' -> キロ
U+3315 '㌕' -> キログラム
U+3316 '㌖' -> キロメートル
U+3317 '㌗' -> キロワット
U+3318 '㌘' -> グラム
U+3319 '㌙' -> グラムトン
U+331A '㌚' -> クルゼイロ
U+331B '㌛' -> クローネ
U+331C '㌜' -> ケース
U+331D '㌝' -> コルナ
U+331E '㌞' -> コーポ
U+331F '㌟' -> サイクル
U+3320 '㌠' -> サンチーム
U+3321 '㌡' -> シリング
U+3322 '㌢' -> センチ
U+3323 '㌣' -> セント
U+3324 '㌤' -> ダース
U+3325 '㌥' -> デシ
U+3326 '㌦' -> ドル
U+3327 '㌧' -> トン
U+3328 '㌨' -> ナノ
U+3329 '㌩' -> ノット
U+332A '㌪' -> ハイツ
U+332B '㌫' -> パーセント
U+332C '㌬' -> パーツ
U+332D '㌭' -> バーレル
U+332E '㌮' -> ピアストル
U+332F '㌯' -> ピクル
U+3330 '㌰' -> ピコ
U+3331 '㌱' -> ビル
U+3332 '㌲' -> ファラッド
U+3333 '㌳' -> フィート
U+3334 '㌴' -> ブッシェル
U+3335 '㌵' -> フラン
U+3336 '㌶' -> ヘクタール
U+3337 '㌷' -> ペソ
U+3338 '㌸' -> ペニヒ
U+3339 '㌹' -> ヘルツ
U+333A '㌺' -> ペンス
U+333B '㌻' -> ページ
U+333C '㌼' -> ベータ
U+333D '㌽' -> ポイント
U+333E '㌾' -> ボルト
U+333F '㌿' -> ホン
U+3340 '㍀' -> ポンド
U+3341 '㍁' -> ホール
U+3342 '㍂' -> ホーン
U+3343 '㍃' -> マイクロ
U+3344 '㍄' -> マイル
U+3345 '㍅' -> マッハ
U+3346 '㍆' -> マルク
U+3347 '㍇' -> マンション
U+3348 '㍈' -> ミクロン
U+3349 '㍉' -> ミリ
U+334A '㍊' -> ミリバール
U+334B '㍋' -> メガ
U+334C '㍌' -> メガトン
U+334D '㍍' -> メートル
U+334E '㍎' -> ヤード
U+334F '㍏' -> ヤール
U+3350 '㍐' -> ユアン
U+3351 '㍑' -> リットル
U+3352 '㍒' -> リラ
U+3353 '㍓' -> ルピー
U+3354 '㍔' -> ルーブル
U+3355 '㍕' -> レム
U+3356 '㍖' -> レントゲン
U+3357 '㍗' -> ワット
U+FF66 'ｦ' -> ヲ
U+FF67 'ｧ' -> ァ
U+FF68 'ｨ' -> ィ
U+FF69 'ｩ' -> ゥ
U+FF6A 'ｪ' -> ェ
U+FF6B 'ｫ' -> ォ
U+FF6C 'ｬ' -> ャ
U+FF6D 'ｭ' -> ュ
U+FF6E 'ｮ' -> ョ
U+FF6F 'ｯ' -> ッ
U+FF71 'ｱ' -> ア
U+FF72 'ｲ' -> イ
U+FF73 'ｳ' -> ウ
U+FF74 'ｴ' -> エ
U+FF75 'ｵ' -> オ
U+FF76 'ｶ' -> カ
U+FF77 'ｷ' -> キ
U+FF78 'ｸ' -> ク
U+FF79 'ｹ' -> ケ
U+FF7A 'ｺ' -> コ
U+FF7B 'ｻ' -> サ
U+FF7C 'ｼ' -> シ
U+FF7D 'ｽ' -> ス
U+FF7E 'ｾ' -> セ
U+FF7F 'ｿ' -> ソ
U+FF80 'ﾀ' -> タ
U+FF81 'ﾁ' -> チ
U+FF82 'ﾂ' -> ツ
U+FF83 'ﾃ' -> テ
U+FF84 'ﾄ' -> ト
U+FF85 'ﾅ' -> ナ
U+FF86 'ﾆ' -> ニ
U+FF87 'ﾇ' -> ヌ
U+FF88 'ﾈ' -> ネ
U+FF89 'ﾉ' -> ノ
U+FF8A 'ﾊ' -> ハ
U+FF8B 'ﾋ' -> ヒ
U+FF8C 'ﾌ' -> フ
U+FF8D 'ﾍ' -> ヘ
U+FF8E 'ﾎ' -> ホ
U+FF8F 'ﾏ' -> マ
U+FF90 'ﾐ' -> ミ
U+FF91 'ﾑ' -> ム
U+FF92 'ﾒ' -> メ
U+FF93 'ﾓ' -> モ
U+FF94 'ﾔ' -> ヤ
U+FF95 'ﾕ' -> ユ
U+FF96 'ﾖ' -> ヨ
U+FF97 'ﾗ' -> ラ
U+FF98 'ﾘ' -> リ
U+FF99 'ﾙ' -> ル
U+FF9A 'ﾚ' -> レ
U+FF9B 'ﾛ' -> ロ
U+FF9C 'ﾜ' -> ワ
U+FF9D 'ﾝ' -> ン
```

ふむむ。
変体仮名は NFKC 正規化の対象にならないんだな。

更についでに，平仮名もやってしまおう。
平仮名は

```go
var _Hiragana = &RangeTable{
	R16: []Range16{
		{0x3041, 0x3096, 1},
		{0x309d, 0x309f, 1},
	},
	R32: []Range32{
		{0x1b001, 0x1b11e, 1},
		{0x1b150, 0x1b152, 1},
		{0x1f200, 0x1f200, 1},
	},
}
```

と定義されている。
なので `0x3041` から `0x1f200` の範囲でスキャンすればいいかな。
とりあえず

```go
//go:build run
// +build run

package main

import (
	"fmt"
	"unicode"
)

func main() {
	i := 0
	for c := rune(0x3041); c <= rune(0x1f200); c++ {
		if unicode.In(c, unicode.Hiragana) {
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

でいってみよう。

```text
$ go run sample5.go
U+3041 'ぁ' U+3042 'あ' U+3043 'ぃ' U+3044 'い' U+3045 'ぅ' U+3046 'う' U+3047 'ぇ' U+3048 'え' U+3049 'ぉ' U+304A 'お' U+304B 'か' U+304C 'が' U+304D 'き' U+304E 'ぎ' U+304F 'く' U+3050 'ぐ' 
U+3051 'け' U+3052 'げ' U+3053 'こ' U+3054 'ご' U+3055 'さ' U+3056 'ざ' U+3057 'し' U+3058 'じ' U+3059 'す' U+305A 'ず' U+305B 'せ' U+305C 'ぜ' U+305D 'そ' U+305E 'ぞ' U+305F 'た' U+3060 'だ' 
U+3061 'ち' U+3062 'ぢ' U+3063 'っ' U+3064 'つ' U+3065 'づ' U+3066 'て' U+3067 'で' U+3068 'と' U+3069 'ど' U+306A 'な' U+306B 'に' U+306C 'ぬ' U+306D 'ね' U+306E 'の' U+306F 'は' U+3070 'ば' 
U+3071 'ぱ' U+3072 'ひ' U+3073 'び' U+3074 'ぴ' U+3075 'ふ' U+3076 'ぶ' U+3077 'ぷ' U+3078 'へ' U+3079 'べ' U+307A 'ぺ' U+307B 'ほ' U+307C 'ぼ' U+307D 'ぽ' U+307E 'ま' U+307F 'み' U+3080 'む' 
U+3081 'め' U+3082 'も' U+3083 'ゃ' U+3084 'や' U+3085 'ゅ' U+3086 'ゆ' U+3087 'ょ' U+3088 'よ' U+3089 'ら' U+308A 'り' U+308B 'る' U+308C 'れ' U+308D 'ろ' U+308E 'ゎ' U+308F 'わ' U+3090 'ゐ' 
U+3091 'ゑ' U+3092 'を' U+3093 'ん' U+3094 'ゔ' U+3095 'ゕ' U+3096 'ゖ' U+309D 'ゝ' U+309E 'ゞ' U+309F 'ゟ' U+1B001 '𛀁' U+1B002 '𛀂' U+1B003 '𛀃' U+1B004 '𛀄' U+1B005 '𛀅' U+1B006 '𛀆' U+1B007 '𛀇' 
U+1B008 '𛀈' U+1B009 '𛀉' U+1B00A '𛀊' U+1B00B '𛀋' U+1B00C '𛀌' U+1B00D '𛀍' U+1B00E '𛀎' U+1B00F '𛀏' U+1B010 '𛀐' U+1B011 '𛀑' U+1B012 '𛀒' U+1B013 '𛀓' U+1B014 '𛀔' U+1B015 '𛀕' U+1B016 '𛀖' U+1B017 '𛀗' 
U+1B018 '𛀘' U+1B019 '𛀙' U+1B01A '𛀚' U+1B01B '𛀛' U+1B01C '𛀜' U+1B01D '𛀝' U+1B01E '𛀞' U+1B01F '𛀟' U+1B020 '𛀠' U+1B021 '𛀡' U+1B022 '𛀢' U+1B023 '𛀣' U+1B024 '𛀤' U+1B025 '𛀥' U+1B026 '𛀦' U+1B027 '𛀧' 
U+1B028 '𛀨' U+1B029 '𛀩' U+1B02A '𛀪' U+1B02B '𛀫' U+1B02C '𛀬' U+1B02D '𛀭' U+1B02E '𛀮' U+1B02F '𛀯' U+1B030 '𛀰' U+1B031 '𛀱' U+1B032 '𛀲' U+1B033 '𛀳' U+1B034 '𛀴' U+1B035 '𛀵' U+1B036 '𛀶' U+1B037 '𛀷' 
U+1B038 '𛀸' U+1B039 '𛀹' U+1B03A '𛀺' U+1B03B '𛀻' U+1B03C '𛀼' U+1B03D '𛀽' U+1B03E '𛀾' U+1B03F '𛀿' U+1B040 '𛁀' U+1B041 '𛁁' U+1B042 '𛁂' U+1B043 '𛁃' U+1B044 '𛁄' U+1B045 '𛁅' U+1B046 '𛁆' U+1B047 '𛁇' 
U+1B048 '𛁈' U+1B049 '𛁉' U+1B04A '𛁊' U+1B04B '𛁋' U+1B04C '𛁌' U+1B04D '𛁍' U+1B04E '𛁎' U+1B04F '𛁏' U+1B050 '𛁐' U+1B051 '𛁑' U+1B052 '𛁒' U+1B053 '𛁓' U+1B054 '𛁔' U+1B055 '𛁕' U+1B056 '𛁖' U+1B057 '𛁗' 
U+1B058 '𛁘' U+1B059 '𛁙' U+1B05A '𛁚' U+1B05B '𛁛' U+1B05C '𛁜' U+1B05D '𛁝' U+1B05E '𛁞' U+1B05F '𛁟' U+1B060 '𛁠' U+1B061 '𛁡' U+1B062 '𛁢' U+1B063 '𛁣' U+1B064 '𛁤' U+1B065 '𛁥' U+1B066 '𛁦' U+1B067 '𛁧' 
U+1B068 '𛁨' U+1B069 '𛁩' U+1B06A '𛁪' U+1B06B '𛁫' U+1B06C '𛁬' U+1B06D '𛁭' U+1B06E '𛁮' U+1B06F '𛁯' U+1B070 '𛁰' U+1B071 '𛁱' U+1B072 '𛁲' U+1B073 '𛁳' U+1B074 '𛁴' U+1B075 '𛁵' U+1B076 '𛁶' U+1B077 '𛁷' 
U+1B078 '𛁸' U+1B079 '𛁹' U+1B07A '𛁺' U+1B07B '𛁻' U+1B07C '𛁼' U+1B07D '𛁽' U+1B07E '𛁾' U+1B07F '𛁿' U+1B080 '𛂀' U+1B081 '𛂁' U+1B082 '𛂂' U+1B083 '𛂃' U+1B084 '𛂄' U+1B085 '𛂅' U+1B086 '𛂆' U+1B087 '𛂇' 
U+1B088 '𛂈' U+1B089 '𛂉' U+1B08A '𛂊' U+1B08B '𛂋' U+1B08C '𛂌' U+1B08D '𛂍' U+1B08E '𛂎' U+1B08F '𛂏' U+1B090 '𛂐' U+1B091 '𛂑' U+1B092 '𛂒' U+1B093 '𛂓' U+1B094 '𛂔' U+1B095 '𛂕' U+1B096 '𛂖' U+1B097 '𛂗' 
U+1B098 '𛂘' U+1B099 '𛂙' U+1B09A '𛂚' U+1B09B '𛂛' U+1B09C '𛂜' U+1B09D '𛂝' U+1B09E '𛂞' U+1B09F '𛂟' U+1B0A0 '𛂠' U+1B0A1 '𛂡' U+1B0A2 '𛂢' U+1B0A3 '𛂣' U+1B0A4 '𛂤' U+1B0A5 '𛂥' U+1B0A6 '𛂦' U+1B0A7 '𛂧' 
U+1B0A8 '𛂨' U+1B0A9 '𛂩' U+1B0AA '𛂪' U+1B0AB '𛂫' U+1B0AC '𛂬' U+1B0AD '𛂭' U+1B0AE '𛂮' U+1B0AF '𛂯' U+1B0B0 '𛂰' U+1B0B1 '𛂱' U+1B0B2 '𛂲' U+1B0B3 '𛂳' U+1B0B4 '𛂴' U+1B0B5 '𛂵' U+1B0B6 '𛂶' U+1B0B7 '𛂷' 
U+1B0B8 '𛂸' U+1B0B9 '𛂹' U+1B0BA '𛂺' U+1B0BB '𛂻' U+1B0BC '𛂼' U+1B0BD '𛂽' U+1B0BE '𛂾' U+1B0BF '𛂿' U+1B0C0 '𛃀' U+1B0C1 '𛃁' U+1B0C2 '𛃂' U+1B0C3 '𛃃' U+1B0C4 '𛃄' U+1B0C5 '𛃅' U+1B0C6 '𛃆' U+1B0C7 '𛃇' 
U+1B0C8 '𛃈' U+1B0C9 '𛃉' U+1B0CA '𛃊' U+1B0CB '𛃋' U+1B0CC '𛃌' U+1B0CD '𛃍' U+1B0CE '𛃎' U+1B0CF '𛃏' U+1B0D0 '𛃐' U+1B0D1 '𛃑' U+1B0D2 '𛃒' U+1B0D3 '𛃓' U+1B0D4 '𛃔' U+1B0D5 '𛃕' U+1B0D6 '𛃖' U+1B0D7 '𛃗' 
U+1B0D8 '𛃘' U+1B0D9 '𛃙' U+1B0DA '𛃚' U+1B0DB '𛃛' U+1B0DC '𛃜' U+1B0DD '𛃝' U+1B0DE '𛃞' U+1B0DF '𛃟' U+1B0E0 '𛃠' U+1B0E1 '𛃡' U+1B0E2 '𛃢' U+1B0E3 '𛃣' U+1B0E4 '𛃤' U+1B0E5 '𛃥' U+1B0E6 '𛃦' U+1B0E7 '𛃧' 
U+1B0E8 '𛃨' U+1B0E9 '𛃩' U+1B0EA '𛃪' U+1B0EB '𛃫' U+1B0EC '𛃬' U+1B0ED '𛃭' U+1B0EE '𛃮' U+1B0EF '𛃯' U+1B0F0 '𛃰' U+1B0F1 '𛃱' U+1B0F2 '𛃲' U+1B0F3 '𛃳' U+1B0F4 '𛃴' U+1B0F5 '𛃵' U+1B0F6 '𛃶' U+1B0F7 '𛃷' 
U+1B0F8 '𛃸' U+1B0F9 '𛃹' U+1B0FA '𛃺' U+1B0FB '𛃻' U+1B0FC '𛃼' U+1B0FD '𛃽' U+1B0FE '𛃾' U+1B0FF '𛃿' U+1B100 '𛄀' U+1B101 '𛄁' U+1B102 '𛄂' U+1B103 '𛄃' U+1B104 '𛄄' U+1B105 '𛄅' U+1B106 '𛄆' U+1B107 '𛄇' 
U+1B108 '𛄈' U+1B109 '𛄉' U+1B10A '𛄊' U+1B10B '𛄋' U+1B10C '𛄌' U+1B10D '𛄍' U+1B10E '𛄎' U+1B10F '𛄏' U+1B110 '𛄐' U+1B111 '𛄑' U+1B112 '𛄒' U+1B113 '𛄓' U+1B114 '𛄔' U+1B115 '𛄕' U+1B116 '𛄖' U+1B117 '𛄗' 
U+1B118 '𛄘' U+1B119 '𛄙' U+1B11A '𛄚' U+1B11B '𛄛' U+1B11C '𛄜' U+1B11D '𛄝' U+1B11E '𛄞' U+1B150 '𛅐' U+1B151 '𛅑' U+1B152 '𛅒' U+1F200 '🈀' 
```

うわ。
`U+1B001` 以降は変体仮名か。
`U+1B150` から `U+1B152` までは小書きの「ゐ」「ゑ」「を」とのこと。
ちなみに小書きの「ん」はプロポーザルが通らなかったらしい。
よーわからん（笑）

これも NFKC 正規化をかけてみよっか。

```go
//go:build run
// +build run

package main

import (
	"fmt"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func main() {
	for c := rune(0x3041); c <= rune(0x1f200); c++ {
		if unicode.In(c, unicode.Hiragana) {
			n := norm.NFKC.String(string(c))
			if n != string(c) {
				fmt.Printf("%#U -> %s\n", c, n)
			}
		}
	}
}
```

今度は少ないはず（笑）

```text
$ go run sample6.go
U+309F 'ゟ' -> より
U+1F200 '🈀' -> ほか
```

うんうん。
こんな感じだよね。

はぅぅ。
拙作の [gnkf](https://github.com/spiegel-im-spiegel/gnkf "spiegel-im-spiegel/gnkf: Network Kanji Filter by Golang") も直さなきゃ...

## ブックマーク

- [Go 言語と Unicode 正規化]({{< relref "./unicode-normalization.md" >}})

[Go]: https://go.dev/
[`unicode`]: https://pkg.go.dev/unicode "unicode package - unicode - pkg.go.dev"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
