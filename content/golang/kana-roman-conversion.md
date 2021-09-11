+++
title = "仮名・ローマ字変換する Go パッケージを作ってみた"
date =  "2021-09-11T12:31:49+09:00"
description = "ついカッとなってやった。"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "unicode", "character", "package" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[@mattn_jp さんの tweet](https://twitter.com/mattn_jp/status/1436388661231177730) で偽名や偽住所を生成できる [mattn/go-gimei][go-gimei] という [Go] パッケージを知る。
なにそれ便利。
さっそく使わせてもらおう。

ところでこれローマ字があるといいよね。

というわけで，仮名からローマ字に変換するパッケージを作ってしまった（笑）

- [spiegel-im-spiegel/krconv: Convert kana-character to roman-alphabet][krconv]

こんな感じで使うことができる。

```go
package krconv_test

import (
	"fmt"

	"github.com/spiegel-im-spiegel/krconv"
)

func ExampleConvert() {
	s := "マツエ テッペイ　めっちゃほりでぃ ﾅﾝﾊﾞかげつで まんざい みるんだょっ"
	fmt.Println(krconv.Convert(s))
	//Output:
	//matsue teppei metchahoridei nambakagetsude manzai mirundayotsu
}
```

ちなみに UTF-8 文字列が前提ね。

これと先程の [mattn/go-gimei][go-gimei] を組み合わせてみよう。

```go
//go:build run
// +build run

package main

import (
	"fmt"
	"strings"

	"github.com/mattn/go-gimei"
	"github.com/spiegel-im-spiegel/krconv"
)

func main() {
	p := gimei.NewName()
	fmt.Println("氏名：", p.Kanji())
	fmt.Println("カナ：", p.Katakana())
	fmt.Println("ローマ字：", strings.ToTitle(krconv.Convert(p.Last.Hiragana())), strings.Title(krconv.Convert(p.First.Hiragana())))
	fmt.Println("Email：", string([]rune(krconv.Convert(p.First.Hiragana()))[0:1])+"."+krconv.Convert(p.Last.Hiragana())+"@example.com")
}
```

出力結果はこんな感じ。

```text
$ go run sample.go 
氏名： 上原 弥璃
カナ： ウエハラ イヨリ
ローマ字： UEHARA Iyori
Email： i.uehara@example.com
```

どやさ！

変換については，以下のページを参考に，ヘボン式にしている。

- [ヘボン式ローマ字｜神奈川県パスポートセンター公式サイト](http://www.pref.kanagawa.jp/osirase/02/2315/hepburn.html)

ただし，長音の扱い（大野（おおの）→ [☓] oono , [○] ono）については判定が難しそうなので無視している[^lv1]。
そうそう，これに関連して長音の記号（ー）は変換後の文字列から削除することにした。

[^lv1]: ヘボン式では基本的に長音は表記しないらしい。[例示](http://www.pref.kanagawa.jp/osirase/02/2315/hepburn.html "ヘボン式ローマ字｜神奈川県パスポートセンター公式サイト")では「佐藤（さとう） 」の「う」は長音で「SATO」と表記するけど「松浦（まつうら）」の「う」は長音ではないので「MATSUURA」と書けとある。分かるか！ こんなもん。要するに仮名文字からそれが長音かどうか判定するのは不可能なんだよ。せいぜい「ー」のような明確な長音記号を省くくらいしかできないわけ。機械的な処理はこれが限界と諦めた。

全体の処理手順としては

1. 半角全角変換（全角英数は半角に，半角カナは全角に）
2. カタカナ→平仮名変換
3. 文字単位（`rune` 単位ではない）に分割
4. 文字単位でローマ字への置き換え
5. 拗音の変換
6. 撥音・促音の変換

という感じ。
仮名文字以外の漢字や記号（「々」等）は素通しなのであしからず。
ヷ，ヸ，ヹ，ヺ はどう変換していいか分からなかったので，これも素通ししている。
「あ&#x3099;」のように標準的でない結合文字がくっついている場合も素通し。

あとヒープを潤沢に使ってループもぐるぐる回してかなり富豪的なコードになっているので，大規模文字列やクリティカルな処理には向いてない。

というわけで，こんなんでよろしければどうぞ。

## ブックマーク

- [Unicode 文字列を「文字」単位に分離する](https://zenn.dev/spiegel/articles/20210408-unicode-characters)

[Go]: https://golang.org/ "The Go Programming Language"
[go-gimei]: https://github.com/mattn/go-gimei "mattn/go-gimei"
[krconv]: https://github.com/spiegel-im-spiegel/krconv "spiegel-im-spiegel/krconv: Convert kana-character to roman-alphabet"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
