+++
title = "Logf Package v0.2.1 Released"
date = "2018-03-03T20:11:38+09:00"
description = "本当は JSON 形式の構造化ログに挑戦しようと思ったのだが，既に出来のいいパッケージが存在したので止めた。"
image = "/images/attention/tools.png"
tags  = [ "golang", "package", "logger" ]

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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253"><img src="https://images-fe.ssl-images-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
	<dd>Alan A.A. Donovan, Brian W. Kernighan</dd>
	<dd>柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>Book 単行本（ソフトカバー）</dd>
    <dd>ASIN: 4621300253, EAN: 9784621300251</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
