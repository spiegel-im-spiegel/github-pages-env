+++
title = "Go 言語による Unicode 半角/全角変換"
date =  "2017-12-08T13:47:48+09:00"
update = "2018-01-21T11:19:13+09:00"
description = "以前「Go 言語と Unicode 正規化」の脚注でこっそり「単に全角・半角変換ができればいいのなら golang.org/x/text/width パッケージをお勧めする」と書いていたのだが，今回はその話。"
image = "/images/attention/go-code.png"
tags = ["golang", "unicode", "transform", "character"]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

もう2年近く前の記事だが，「[Go 言語と Unicode 正規化]({{< relref "unicode-normalization.md" >}})」の脚注でこっそり「単に全角・半角変換ができればいいのなら `golang.org/x/text/`[`width`] パッケージをお勧めする」と書いていたのだが，今回はその話。

## golang.org/x/text/width パッケージ

`golang.org/x/text/`[`width`] パッケージ[^wdth] では半角/全角変換用に以下の [`width`].`Transformer` が用意されている。

[^wdth]: “width” って日本人には発音しにくいよね。「[IT業界で横行する恥ずかしい英語発音](https://qiita.com/ryounagaoka/items/290885ee3291b393fe1f "IT業界で横行する恥ずかしい英語発音 - Qiita")」でもネタになるほど（笑）

```go
var (
	// Fold is a transform that maps all runes to their canonical width.
	//
	// Note that the NFKC and NFKD transforms in golang.org/x/text/unicode/norm
	// provide a more generic folding mechanism.
	Fold Transformer = Transformer{foldTransform{}}

	// Widen is a transform that maps runes to their wide variant, if
	// available.
	Widen Transformer = Transformer{wideTransform{}}

	// Narrow is a transform that maps runes to their narrow variant, if
	// available.
	Narrow Transformer = Transformer{narrowTransform{}}
)
```

[`width`].`Widen` が全角変換用， [`width`].`Narrow` が半角変換用である。 

たとえば，半角変換なら

```go
fmt.Println(width.Narrow.String("１２３４５６７８９０アイウエオカキクケコＡＢＣＤＥＦＧＨＩＪＫ"))
//Outpput
//1234567890ｱｲｳｴｵｶｷｸｹｺABCDEFGHIJK
```

全角変換なら

```go
fmt.Println(width.Narrow.String("1234567890ｱｲｳｴｵｶｷｸｹｺABCDEFGHIJK"))
//Outpput
//１２３４５６７８９０アイウエオカキクケコＡＢＣＤＥＦＧＨＩＪＫ
```

のように変換される。
[`io`].`Reader` を使うなら

{{< highlight go "hl_lines=2" >}}
r := bytes.NewBufferString("1234567890ｱｲｳｴｵｶｷｸｹｺABCDEFGHIJK")
tr := transform.NewReader(r, wdth.Widen)
buf := new(bytes.Buffer)
io.Copy(buf, tr)
fmt.Println(buf)
//Outpput
//１２３４５６７８９０アイウエオカキクケコＡＢＣＤＥＦＧＨＩＪＫ
{{< /highlight >}}

ってな感じで書ける。

面白いのは [`width`].`Fold` で，半角カナ文字は全角に，全角英数字は半角に，いい感じに変換してくれるのだ。

```go
fmt.Println(width.Fold.String("１２３４５６７８９０ｱｲｳｴｵｶｷｸｹｺＡＢＣＤＥＦＧＨＩＪＫ"))
//Outpput
//1234567890アイウエオカキクケコABCDEFGHIJK
```

[`width`].`Fold` は， Unicode 正規化と異なり，「神と神[^fnt1]」のような異体字は

[^fnt1]: ブラウザやブラウザの設定によっては「神」が上手く表示できないようです。 最初の「神」は「`U+FA19 '神'`」で CJK 互換文字です。２番目の「神」が通常の「`U+795E '神'`」です。

```go
fmt.Println(width.Fold.String("神と神"))
//Outpput
//神と神
```

と影響を受けない[^un1]。
ブラボー！

[^un1]: Unicode 正規化と異体字の問題については「[Go 言語と Unicode 正規化]({{< relref "unicode-normalization.md" >}})」を参照のこと。

ちなみに Unicode の円記号（YEN SIGN）は，通常文字（半角相当）が「¥：U+00A5[^yen1]」で全角文字が「￥：U+FFE5」である。
JIS ラテン文字（JIS X 0201）の半角円記号（0x5C）は機械的に変換しても同じ値の「`\` (REVERSE SOLIDUS)」のままなので（当たり前だけど）ご注意を[^ws]。
そろそろ JIS ラテン文字のことは忘れていいかもねー。

[^ws]: この辺の事情は韓国語の WON SIGN「₩：U+20A9」とかでも同じらしい。

[^yen1]: 円記号「¥：U+00A5」のコード・ポイントは ISO-8859-1 (Latin-1) で割り当てられている円記号のコードと同じ。というか，そもそも Unicode は ISO-8859-1 と親和するように設計されているのでこうなっているのだが。円記号が ISO-8859-1 に収録されていて（8ビット空間に収まっていて）よかったね，というべきか（笑）

## [spiegel-im-spiegel/text] に組み込んでみた

というわけで，この Unicode 半角/全角変換機能を私作の [spiegel-im-spiegel/text] v0.5.0 に組み込んでみた。
こんな感じに使える。

```go
res := width.Reader(bytes.NewBufferString("１２３４５６７８９０ｱｲｳｴｵｶｷｸｹｺＡＢＣＤＥＦＧＨＩＪＫ"), width.Fold)
buf := new(bytes.Buffer)
io.Copy(buf, res)
fmt.Println(buf)
// Output:
// 1234567890アイウエオカキクケコABCDEFGHIJK
```

あと，このパッケージの CLI 実装用に作った「[nkf っぽいなにか]({{< ref "/remark/2017/12/like-nkf.md" >}})」にもサブコマンド `width` として組み込んでいる。

```text
$ gonkf width -h
Convert character width of text (UTF-8 text only)

Usage:
  gonkf width [flags] [text file]

Flags:
  -f, --form string     form of width [fold|narrow|widen] (default "fold")
  -h, --help            help for width
  -o, --output string   output file path

$ echo １２３４５６７８９０ｱｲｳｴｵｶｷｸｹｺＡＢＣＤＥＦＧＨＩＪＫ | gonkf width -f fold
1234567890アイウエオカキクケコABCDEFGHIJK
```

たとえば Sjift-JIS テキストに対して機能をパイプで繋いで

```text
$ cat sjis.txt
１２３４５６７８９０ｱｲｳｴｵｶｷｸｹｺＡＢＣＤＥＦＧＨＩＪＫ

$ gonkf conv -s sjis sjis.txt | gonkf nwline -f lf | gonkf width -f fold > utf8.txt

$ cat utf8.txt
1234567890アイウエオカキクケコABCDEFGHIJK
```

みたいな変換もできる（画面はイメージです`w`）。

これで概ね [spiegel-im-spiegel/text] でやりたいことはやり尽くしたかな。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`norm`]: https://godoc.org/golang.org/x/text/unicode/norm "norm - GoDoc"
[`width`]: https://godoc.org/golang.org/x/text/width "width - GoDoc"
[`transform`]: https://godoc.org/golang.org/x/text/transform "transform - GoDoc"
[`io`]: https://golang.org/pkg/io/ "io - The Go Programming Language"
[spiegel-im-spiegel/text]: https://github.com/spiegel-im-spiegel/text "spiegel-im-spiegel/text: Encoding/Decoding Text Package by Golang"

## ブックマーク

- [Goで全角/半角変換する - Qiita](https://qiita.com/ktnyt/items/ee56d5425e1c0ad42094) : こちらは内部に変換辞書を持っているらしい

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-20</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4865940391/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4865940391.09._SCTHUMBZZZ_.jpg"  alt="Kotlinスタートブック -新しいAndroidプログラミング"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4839959234/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4839959234.09._SCTHUMBZZZ_.jpg"  alt="Docker実戦活用ガイド"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274218961/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274218961.09._SCTHUMBZZZ_.jpg"  alt="グッド・マス ギークのための数・論理・計算機科学"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-13">2016-07-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
