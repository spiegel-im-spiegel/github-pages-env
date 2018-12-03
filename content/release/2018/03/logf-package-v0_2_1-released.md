+++
title = "Logf Package v0.2.1 Released"
date = "2018-03-03T20:11:38+09:00"
description = "本当は JSON 形式の構造化ログに挑戦しようと思ったのだが，既に出来のいいパッケージが存在したので止めた。"
image = "/images/attention/tools.png"
tags  = [ "golang", "package", "logger" ]

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
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

レベル付き logger パッケージ [spiegel-im-spiegel/logf] v0.2.1 をリリースした。

- [spiegel-im-spiegel/logf: Simple leveled logging package by Golang](https://github.com/spiegel-im-spiegel/logf)

標準の [log] パッケージをラップする形で実装している。
以下のように [lestrrat-go/file-rotatelogs] のようなローテーション機能付きの Writer と組み合わせて使うといい感じになる。

```go
package main

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/spiegel-im-spiegel/logf"
)

func main() {
	rl, err := rotatelogs.New("./log.%Y%m%d%H%M.txt")
	if err != nil {
		logf.Fatal(err)
		return
	}
	logger := logf.New(
		logf.WithFlags(logf.LstdFlags|logf.Lshortfile),
		logf.WithPrefix("[Sample] "),
		logf.WithWriter(rl),
		logf.WithMinLevel(logf.INFO),
	)
	logger.Print("Information")
	//Output:
	//[Sample] 2009/11/10 23:00:00 sample.go:20: [INFO] Information
}
```

本当は JSON 形式の構造化ログに挑戦しようと思ったのだが，既に出来のいいパッケージが存在したので止めた。

- [rs/zerolog: Zero Allocation JSON Logger](https://github.com/rs/zerolog)

なので，これ以上 [spiegel-im-spiegel/logf] を弄ることはないと思う。

## ブックマーク

- [Log パッケージで遊ぶ]({{< ref "/golang/logger.md" >}})

[spiegel-im-spiegel/logf]: https://github.com/spiegel-im-spiegel/logf "spiegel-im-spiegel/logf: Simple leveled logging package by Golang"
[log]: https://golang.org/pkg/log/ "log - The Go Programming Language"
[lestrrat-go/file-rotatelogs]: https://github.com/lestrrat-go/file-rotatelogs "lestrrat-go/file-rotatelogs: Port of perl5 File::RotateLogs to Go"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-20</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4865940391/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4865940391.09._SCTHUMBZZZ_.jpg"  alt="Kotlinスタートブック -新しいAndroidプログラミング"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4839959234/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4839959234.09._SCTHUMBZZZ_.jpg"  alt="Docker実戦活用ガイド"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274218961/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274218961.09._SCTHUMBZZZ_.jpg"  alt="グッド・マス ギークのための数・論理・計算機科学"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-13">2016-07-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
