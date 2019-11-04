+++
title = "文字エンコーディングの検出"
date =  "2017-12-04T12:42:55+09:00"
update =  "2017-12-06T00:20:33+09:00"
description = "文字エンコーディングを変換するのはいいのだが，そのためには元の文字列の文字エンコーディングが分かってないといけない。あぁ，これみんな苦労してるやつだよね。"
image = "/images/attention/go-code.png"
tags = ["golang", "character", "encoding", "transform"]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/profile/"
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

以前に洒落で作った「[Markdown 形式のリンクを生成するツール]({{< relref "make-link-with-markdown-format.md" >}} "Markdown 形式のリンクを生成するツールを作ってみた")」が思いのほか便利で重宝しているのだが，実際に使ってみると EUC-JP や Shift-JIS な日本語サイトが今だにあって，最初は手で修正していたのだが，だんだん面倒臭くなってきたので変換ロジックを入れることにした。

[文字エンコーディングを変換する]({{< relref "transform-character-encoding.md" >}} "文字エンコーディング変換")のはいいのだが，そのためには元の文字列の文字エンコーディングが分かってないといけない。
HTML の中の "charset” を見て変換する手もあるのだが， HTML のバージョンによってお作法が違うし，そもそもそれが正しいとは限らない。
あぁ，これみんな苦労してるやつだよね。

文字列が UTF-8 かどうかは [`utf8`].`ValidString()` 関数で分かるので，そこから強引に各種文字エンコーディング変換を試すという力技もあるが，あんまりスマートじゃない。
それは最後の手段に取っておくとして，まずは文字エンコーディング検出ができるパッケージを誰か公開してないか探してみた。
したら，ありましたよ。

- [chardet/detector_test.go at master · saintfish/chardet](https://github.com/saintfish/chardet)

しかも，これを使って簡易 nkf を [Go 言語]で公開しておられる方もいた。

- [Go言語でnkfみたいなやつ - How to spend the terminal](http://moxtsuan.hatenablog.com/entry/nkf-go)
    - [moxtsuan/go-nkf: tiny nkf(Only JIS, SJIS, EUC-JP, UTF-8)](https://github.com/moxtsuan/go-nkf)

[moxtsuan/go-nkf] は今回の目的にはオーバースペックなので，参考にしつつも自前で [saintfish/chardet] を組み込んでみることにした。
具体的には [`charencode.go`](https://github.com/spiegel-im-spiegel/mklink/blob/master/charencode.go "mklink/charencode.go at master · spiegel-im-spiegel/mklink") というファイルだ。

まず文字エンコーディングの検出部分はこんな感じに書ける。

{{< highlight go "hl_lines=5-6" >}}
import "github.com/saintfish/chardet"

//DetectCharEncode returns character encoding
func DetectCharEncode(body []byte) CharEncode {
	det := chardet.NewTextDetector()
	res, err := det.DetectBest(body)
	if err != nil {
		return CharUnknown
	}
	return TypeofCharEncode(res.Charset)
}
{{< /highlight >}}

関数内の最初の2行で [saintfish/chardet] による文字エンコーディング検出を行っている。
結果は `res.Charset` に文字列（！）で格納されるのだが，文字列で取り回すのはあんまりなので

```go
//CharEncode is type of character encoding
type CharEncode int

const (
	//CharUnknown is unknown character
	CharUnknown CharEncode = iota
	//CharUTF8 is UTF-8
	CharUTF8
	//CharISO8859_1 is ISO-8859-1
	CharISO8859_1
	//CharShiftJIS is Shift-JIS
	CharShiftJIS
	//CharEUCJP is EUC-JP
	CharEUCJP
	//CharISO2022JP is ISO-2022-JP
	CharISO2022JP
)
```

という enum 型を作って，最終的にはこれを返すようにしている。
これを見て分かるように今のところは Shift-JIS, EUC-JP, ISO-2022-JP のみ対応している。
中国語とか韓国語とかは知らんぷりする（笑）

`DetectCharEncode()` 関数の結果を使って実際に UTF-8 へ変換する。
こんな感じ。

```go
import (
	"bytes"
	"io"

	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

//ToUTF8 returns string with UTF-8 encoding
func ToUTF8(body []byte) string {
	var trans transform.Transformer
	switch DetectCharEncode(body) {
	case CharUTF8, CharISO8859_1:
		return string(body)
	case CharShiftJIS:
		trans = japanese.ShiftJIS.NewDecoder()
	case CharEUCJP:
		trans = japanese.EUCJP.NewDecoder()
	case CharISO2022JP:
		trans = japanese.ISO2022JP.NewDecoder()
	default:
		return ""
	}
	buf := new(bytes.Buffer)
	io.Copy(buf, transform.NewReader(bytes.NewReader(body), trans))
	return buf.String()
}
```

UTF-8 および ISO-8859-1 (Latin-1) は素通し。
他は decoder を作って変換を行っている。

これを組み込んで実際に動かしてみた。
とりあえず ITmedia (Shift-JIS) とはてなダイアリー（EUC-JP）のサイトで確認したが，ちゃんと動いてるっぽい。
ただ，やっぱり時々失敗するんだよねー。
まぁこれはしょうがないか。

失敗が多いようなら最終手段（片っ端から変換を試す）を執ることにしよう[^cnv1]。
あっ，コードの再利用は（こんなんでよければ）ご自由に。

[^cnv1]: よく考えたら，この手は使えない。たとえば ISO-2022-JP は7ビット・エンコーディングなので [`utf8`].`ValidString()` 関数でチェックしても valid になるに決まってるし（制御コードは無視するようだ），間違ったエンコーディングで無理やり UTF-8 に decode した結果も valid になるに決まっている `orz`

## 追記（2017-12-04）

この記事を書いた後いろいろ試してみたのだが，やっぱり Shift-JIS と EUC-JP の誤判定が多い。
で，上の検出と変換部分を別パッケージにしていろいろ調べてみた。

- [spiegel-im-spiegel/text: Encoding/Decoding Text Package by Golang](https://github.com/spiegel-im-spiegel/text)

このパッケージでは `detect`, `encode`, `decode`, `convert` と機能毎にサブパッケージに分けたので多少使いやすいのではないかと思う。

で，どうも短い文字列だと誤判定する確率が爆上がりする。
[saintfish/chardet] では可能性の高い（？）文字エンコーディング候補から順に配列の形で返すのだが（上のコードの `DetectBest()` 関数では配列の最上位のものを返す），短い文字列では候補を絞りきれず誤判定になってしまうようだ。

苦肉の策として文字エンコーディング候補のうち日本語のエンコーディングを優先して選ぶ関数を作った。
これで [`charencode.go`](https://github.com/spiegel-im-spiegel/mklink/blob/master/charencode.go "mklink/charencode.go at master · spiegel-im-spiegel/mklink") はこんな感じになった。

{{< highlight go "hl_lines=12" >}}
package mklink

import (
	"github.com/spiegel-im-spiegel/text/decode"
)

//ToUTF8 returns string with UTF-8 encoding
func ToUTF8(body []byte) string {
	if len(body) == 0 {
		return ""
	}
	utf8Text, err := decode.ToUTF8ja(body)
	if err != nil {
		return ""
	}
	return string(utf8Text)
}
{{< /highlight >}}

これで他言語の文字エンコーディングと間違える確率は減ったが，今回はページのタイトルだけを対象にしていたので，やっぱり Shift-JIS と EUC-JP とで判定がとっちらかってしまう。
というところで力尽きた。

そういや UTF-8 から ISO-2022-JP への encode でバグ見つけちゃったよ。

ISO-2022-JP のルールでは行末で必ず US-ASCII に戻さないといけない（1BH 28H 42H のシーケンスを出力する）のだが，文字列の末尾に改行がない場合に US-ASCII に戻していない。

まっ，日本語圏の人しか使わないエンコーディングだし，今は日本語圏の人でも使ってる人は殆どおらんじゃろう。
メールも今は UTF-8 が主流だし。
メールメッセージで末尾に改行がないという状況もないしね（実際にはメールメッセージへの電子署名で影響が出るのだが）。

というわけで，標準パッケージでもないし，実害はないので放置する（これが15年くらい前なら煩く言うところだったろうけど）。

## 追記（2017-12-06）

タイムリーなことに Web ページの文字エンコーディングを簡単に取得する方法について書かれた記事が上がっていた。

- [Big Sky :: Golang で HTTP コンテンツの charset 判定をするには](https://mattn.kaoriya.net/software/lang/go/20171205164150.htm)

HTML の charset 設定だけを見るんじゃなくてレスポンスのヘッダとか色々見て総合判断しているらしい。
具体的にはこんな感じになる。

{{< highlight go "hl_lines=9-20" >}}
resp, err := http.Get(url)
if err != nil {
    return nil, err
}
defer resp.Body.Close()

link.Location = resp.Request.URL.String()

br := bufio.NewReader(resp.Body)
var r io.Reader = br
if data, err2 := br.Peek(1024); err2 == nil { //next 1024 bytes without advancing the reader.
    enc, name, _ := charset.DetermineEncoding(data, resp.Header.Get("content-type"))
    if enc != nil {
        r = enc.NewDecoder().Reader(br)
    } else if len(name) > 0 {
        if enc := encoding.GetEncoding(name); enc != nil {
            r = enc.NewDecoder().Reader(br)
        }
    }
}
doc, err := goquery.NewDocumentFromReader(r)
if err != nil {
    return link, err
}
{{< /highlight >}}

これを自前でやると相当に面倒臭いので忌避していたのだが，パッケージがちゃんとあるんだねぇ。
助かりました。

これで [saintfish/chardet] を使って文字エンコーディングを検出するコードがまるっと要らなくなったけど，いろいろと勉強になったので良しとしよう。
そして [mattn/go-encoding] を導入したおかげで何気に対応する文字エンコードディングが増えてしまった（笑）

というわけで「[Markdown 形式のリンクを生成するツール]({{< relref "make-link-with-markdown-format.md" >}} "Markdown 形式のリンクを生成するツールを作ってみた")」の v0.1.10 が絶賛リリース中です。

- [Release v0.1.10 · spiegel-im-spiegel/mklink · GitHub](https://github.com/spiegel-im-spiegel/mklink/releases/tag/v0.1.10)

[spiegel-im-spiegel/text] は他のことに再利用しているので，こっちも問題なし。
やっぱ文字エンコーディングの変換は苦労するなぁ。
旧 JIS とかイラわんだけマシだけど。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`utf8`]: https://golang.org/pkg/unicode/utf8/ "utf8 - The Go Programming Language"
[saintfish/chardet]: https://github.com/saintfish/chardet "saintfish/chardet: Charset detector library for golang derived from ICU"
[moxtsuan/go-nkf]: https://github.com/moxtsuan/go-nkf "moxtsuan/go-nkf: tiny nkf(Only JIS, SJIS, EUC-JP, UTF-8)"
[spiegel-im-spiegel/text]: https://github.com/spiegel-im-spiegel/text "spiegel-im-spiegel/text: Encoding/Decoding Text Package by Golang"
[mattn/go-encoding]: https://github.com/mattn/go-encoding

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
