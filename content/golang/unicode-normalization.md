+++
date = "2016-01-30T04:05:52+09:00"
description = "今回は少し目先を変えて「Unicode 正規化」のお話。"
draft = false
tags = ["golang", "unicode", "normalization", "character"]
title = "Go 言語と Unicode 正規化"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

今回は少し目先を変えて「Unicode 正規化（normalization）」のお話。

## 2羽の「ペンギン」

まず「ペンギン」という文字列を思い浮かべてみる。
この文字列を Unicode のコードポイントで表すと以下のようになる。

- ペ：U+30DA
- ン：U+30F3
- ギ：U+30AE
- ン：U+30F3

ところでペンギンの「ペ」と「ギ」は半濁点および濁点を含む。
Unicode は「ペ」と「ギ」をそれぞれ2つの要素に分解できる。

- ペ：U+30D8 + U+309A
- ン：U+30F3
- ギ：U+30AD + U+3099
- ン：U+30F3

U+309A および U+3099 はそれぞれ半濁点と濁点を表す「結合文字（combining character）」である。
「ヘ」や「キ」のような「基底文字（base character）」に結合文字を1つ以上[^cs] 付加した文字を「合成列（composite sequence）」と呼ぶ。
これに対して「ペ：U+30DA」や「ギ：U+30AE」のような文字を「事前合成形（precomposed）」と呼ぶ。

[^cs]: ちなみに結合文字はひとつの基底文字に対して複数付加される場合もある。しかもこの場合に結合文字同士の順序は不定である。

つまり同じ文字を同じ文字集合[^ccs] で表しているのにもかかわらず複数の符号化[^ecd] が存在するわけだ。
これを「重複符号化」と言う。
文字集合に重複符号化があるというのは，はっきり言って「クソ仕様」である。

[^ccs]: 正しくは「符号化文字集合（coded character set）」である。
[^ecd]: これは UTF-8 などの「文字エンコーディング」とは異なるものだ。

もちろんこれは Unicode を作った連中がヘボいのではなく（いや，ヘボいのかもしれないが），いわゆる「歴史的経緯」というやつである[^jis]。
だからこれはこういうものだと諦めるしかない。

[^jis]: 日本の JIS 規格にも「歴史的経緯」による重複符号化がある。言わずと知れた「半角」「全角」文字である。異体字も一種の重複符号化と言える。

しかし情報処理を行う上では，この2羽の「ペンギン」が等価（equivalance）であることを示す手立てを考えなければならない。

## 正規等価

2羽の「ペンギン」が等価であることを示す一番簡単な方法は，文字列を事前合成形あるいは合成列のどちらかに統一（＝正規化）してしまえばいい。
これを「正規等価（canonical equivalance）」と呼ぶ。
このうち，事前合成形に正規化する方法を “NFC（Normalization Form Composition）”，合成列に正規化する方法を “NFD（Normalization Form Decomposition）” と呼ぶ。

[Go 言語]では `golang.org/x/text/unicode/`[`norm`] パッケージで Unicode 文字列を正規化できる。
まぁ，コードで書いたほうがはやいか。

```go
package main

import (
	"fmt"

	"golang.org/x/text/unicode/norm"
)

func main() {
	penguin := "ペンギン"
	for pos, runeValue := range penguin {
		fmt.Printf("penguin[%d] = %#U\n", pos, runeValue)
	}

	penguin2 := string(norm.NFD.Bytes([]byte(penguin)))
	for pos, runeValue := range penguin2 {
		fmt.Printf("penguin2[%d] = %#U\n", pos, runeValue)
	}

	penguin3 := string(norm.NFC.Bytes([]byte(penguin2)))
	for pos, runeValue := range penguin3 {
		fmt.Printf("penguin3[%d] = %#U\n", pos, runeValue)
	}
}
```

実行結果はこんな感じ。

```
penguin[0] = U+30DA 'ペ'
penguin[3] = U+30F3 'ン'
penguin[6] = U+30AE 'ギ'
penguin[9] = U+30F3 'ン'
penguin2[0] = U+30D8 'ヘ'
penguin2[3] = U+309A '゚'
penguin2[6] = U+30F3 'ン'
penguin2[9] = U+30AD 'キ'
penguin2[12] = U+3099 '゙'
penguin2[15] = U+30F3 'ン'
penguin3[0] = U+30DA 'ペ'
penguin3[3] = U+30F3 'ン'
penguin3[6] = U+30AE 'ギ'
penguin3[9] = U+30F3 'ン'
```

NFC と NFD が交換可能であることがわかると思う。

## 3羽目の「ペンギン」と互換等価

さてここで3羽目の「ペンギン」に登場してもらおう。

- ﾍ：U+FF8D
- ﾟ：U+FF9F
- ﾝ：U+FF9D
- ｷ：U+FF77
- ﾞ：U+FF9E
- ﾝ：U+FF9D

これはいわゆる「半角カナ」である。
半角カナの濁点および半濁点は結合文字ではない。
そのためこの文字列を NFC で事前合成形に正規化しようとしても

```go
package main

import (
	"fmt"

	"golang.org/x/text/unicode/norm"
)

func main() {
	penguin := "ﾍﾟﾝｷﾞﾝ"
	for pos, runeValue := range penguin {
		fmt.Printf("penguin[%d] = %#U\n", pos, runeValue)
	}

	penguin2 := string(norm.NFC.Bytes([]byte(penguin)))
	for pos, runeValue := range penguin2 {
		fmt.Printf("penguin2[%d] = %#U\n", pos, runeValue)
	}
}
```

```
penguin[0] = U+FF8D 'ﾍ'
penguin[3] = U+FF9F 'ﾟ'
penguin[6] = U+FF9D 'ﾝ'
penguin[9] = U+FF77 'ｷ'
penguin[12] = U+FF9E 'ﾞ'
penguin[15] = U+FF9D 'ﾝ'
penguin2[0] = U+FF8D 'ﾍ'
penguin2[3] = U+FF9F 'ﾟ'
penguin2[6] = U+FF9D 'ﾝ'
penguin2[9] = U+FF77 'ｷ'
penguin2[12] = U+FF9E 'ﾞ'
penguin2[15] = U+FF9D 'ﾝ'
```

何も変わらないことが分かるだろう。

このような場合は半角の「ﾍﾟﾝｷﾞﾝ」と互換性のある文字列に正規化できるとよい。
これを「互換等価（compatibility equivalance）」と呼ぶ。
具体的には NFKC（Normalization Form Compatibility Composition）および NFKD（Normalization Form Compatibility Decomposition）の2つの正規化がある。

早速 [`norm`] パッケージを使ってコードを書いてみる。

```go
package main

import (
	"fmt"

	"golang.org/x/text/unicode/norm"
)

func main() {
	penguin := "ﾍﾟﾝｷﾞﾝ"
	for pos, runeValue := range penguin {
		fmt.Printf("penguin[%d] = %#U\n", pos, runeValue)
	}

	penguin2 := string(norm.NFKC.Bytes([]byte(penguin)))
	for pos, runeValue := range penguin2 {
		fmt.Printf("penguin2[%d] = %#U\n", pos, runeValue)
	}
}
```

これを実行すると

```
penguin[0] = U+FF8D 'ﾍ'
penguin[3] = U+FF9F 'ﾟ'
penguin[6] = U+FF9D 'ﾝ'
penguin[9] = U+FF77 'ｷ'
penguin[12] = U+FF9E 'ﾞ'
penguin[15] = U+FF9D 'ﾝ'
penguin2[0] = U+30DA 'ペ'
penguin2[3] = U+30F3 'ン'
penguin2[6] = U+30AE 'ギ'
penguin2[9] = U+30F3 'ン'
```

となり， NFC で正規化した「ペンギン」と等価であることがわかる。

互換等価による正規化は応用範囲が広い。
たとえば「㈱」（U+3231）は「（株）」（U+0028 + U+682A + U+0029）に変換される。
文字列検索の前に互換等価による正規化を行っておくことで処理がやりやすくなるというのはあるかもしれない。
ただし， NFC と NFD は交換可能だが， NFKC や NFKD で正規化した文字列を元のオリジナルに戻す方法はないので注意が必要である[^hw]。

[^hw]: 単に全角・半角変換ができればいいのなら `golang.org/x/text/`[`width`] パッケージをお勧めする。

## 恐怖の CJK 互換文字

正規等価については注意すべき点がある。
有名な「神」を例に挙げよう。
これを NFC / NFD で正規化する。

```go
package main

import (
	"fmt"

	"golang.org/x/text/unicode/norm"
)

//神様の遊び
func main() {
	penguin := "神"
	for pos, runeValue := range penguin {
		fmt.Printf("penguin[%d] = %#U\n", pos, runeValue)
	}

	penguin2 := string(norm.NFKC.Bytes([]byte(penguin)))
	for pos, runeValue := range penguin2 {
		fmt.Printf("penguin2[%d] = %#U\n", pos, runeValue)
	}

	penguin3 := string(norm.NFKD.Bytes([]byte(penguin)))
	for pos, runeValue := range penguin3 {
		fmt.Printf("penguin3[%d] = %#U\n", pos, runeValue)
	}
}
```

これを実行すると

```
penguin[0] = U+FA19 '神'
penguin2[0] = U+795E '神'
penguin3[0] = U+795E '神'
```

となり，NFC でも NFD でも違う文字になってしまった。
ちなみに「神」から「神」へ正規化する方法はない。
困ったね。

実は「神」は「CJK 互換文字」と呼ばれるグループに属し，「神」とは異体字の関係にある。
故に「神」を「神」に正規化してしまったのである。

しかし，これは明らかに仕様ミスである。
「神」と「神」のような異体字の関係は本来なら正規等価ではなく互換等価であるべきだからだ。

...やっぱり Unicode はクソ仕様だ。

しかし，これが実際の場面で問題になることは少ないと思われる。
なぜなら，正規化を行うのは「2つの文字列が等価であるか？」を調べるための手段にすぎないからだ。
普通はね。

ところが，普通でないことをする馬鹿がいるのである。

### 独自路線に走る Apple

Apple の OS X （iOS も？）のファイルシステムである HFS+ はファイル名を NFD 相当に正規化するという恐ろしい仕様になっている[^fs]。
しかしそれでは，先ほどの例のように異体字に関しては正規化によって別の文字に変えられてしまうため困ったことになってしまう。

そこで Apple は CJK 互換文字を含むいくつかの文字を正規化の対象から外すという蛮行に出た。
俗に “NFD-mac” などと呼ばれる独自路線に走ってしまったわけだ。

[^fs]: ちなみに Windows のファイルシステムはフォルダ・ファイルの名前を正規化するとかいうアホなことはしない。事前合成形も合成列も受け入れる。見かけ同じ名前のフォルダ・ファイルが複数できる可能性はあるが，それはそれ。多分，ほとんどの OS のファイルシステムは名前の正規化なんてしてないはず。

これにより様々な（特にマルチプラットフォームな）アプリケーションが多大なる迷惑を被ることになるが[^mac]，深くはツッコむまい。

[^mac]: たとえば Linus Torvalds は HFS+ に起因する git の脆弱性問題で[激怒](http://cpplover.blogspot.jp/2015/01/blog-post_14.html)している。

## Unicode 正規化に関するまとめ

以上， Unicode 正規化の4つの方式をまとめると以下のようになる。

{{< fig-gen title="Text normalization in Go" link="https://blog.golang.org/normalization" >}}
<table>
  <tr>
    <th></th>
    <th>Composing</th>
    <th>Decomposing</th>
  </tr>
  <tr>
    <th>Canonical equivalence</th>
    <td>NFC</td>
    <td>NFD<br></td>
  </tr>
  <tr>
    <th>Compatibility equivalence</th>
    <td>NFKC</td>
    <td>NFKD</td>
  </tr>
</table>
{{< /fig-gen >}}

Unicode 文字列の等価属性を調べる際には是非参考にどうぞ。

## ブックマーク

- [特別編24 JIS X 0213の改正は、文字コードにどんな未来をもたらすか（7）　番外編：改正JIS X 0213とUnicodeの等価属性／正規化について（上）](http://internet.watch.impress.co.jp/www/column/ogata/sp24.htm)
- [特別編24 JIS X 0213の改正は、文字コードにどんな未来をもたらすか（7）　番外編：改正JIS X 0213とUnicodeの等価属性／正規化について（下）](http://internet.watch.impress.co.jp/www/column/ogata/sp25.htm)
- [Unicode正規化](http://nomenclator.la.coocan.jp/unicode/normalization.htm)
- [Text normalization in Go - The Go Blog](https://blog.golang.org/normalization)
- [Go で UTF-8 の文字列を扱う - Qiita](http://qiita.com/masakielastic/items/01a4fb691c572dd71a19)
- [文字コード地獄秘話 第3話：後戻りの効かないUnicode正規化 - ALBERT Engineer Blog](http://tech.albert2005.co.jp/blog/2014/11/21/mco-normalize/)
- [本の虫: Linus Torvalds、HFS+に激怒](http://cpplover.blogspot.jp/2015/01/blog-post_14.html)

[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`norm`]: https://godoc.org/golang.org/x/text/unicode/norm "norm - GoDoc"
[`width`]: https://godoc.org/golang.org/x/text/width "width - GoDoc"
[String と Rune]: {{< ref "golang/string-and-rune.md" >}}
[string]: http://golang.org/ref/spec#String_types
[rune]: http://blog.golang.org/strings "Strings, bytes, runes and characters in Go - The Go Blog"
