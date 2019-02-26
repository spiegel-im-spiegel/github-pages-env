+++
title = "Go 1.12 がリリースされた"
date = "2019-02-26T23:03:40+09:00"
description = "主な変更点としては TLS 1.3 の暫定的なサポートとモジュール・モードの挙動の一部が変わったことだろうか。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "engineering", "module", "tls", "regular-expression", "concurrency" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Go 1.12 がリリースされた。
そういえば2月ももう終わりか。

- [Go 1.12 is released - The Go Blog](https://blog.golang.org/go1.12)
- [Go 1.12 Release Notes - The Go Programming Language](https://golang.org/doc/go1.12)

主な変更点としては TLS 1.3 の暫定的なサポート（有効にするには環境変数の設定が必要）と[モジュール・モード]({{< ref "/golang/go-module-aware-mode.md" >}} "モジュール対応モードへの移行を検討する")の挙動の一部が変わったことだろうか。

たとえば環境変数 `$GO111MODULE` を `on` にしておけば [mattn](https://github.com/mattn) さんの [jvgrep](https://github.com/mattn/jvgrep "mattn/jvgrep: grep for japanese vimmer") をインストールする際にも任意のフォルダで

```text
$ go get github.com/mattn/jvgrep@latest
```

とすればよい。
ダミーの `go.mod` ファイルを作る必要はなくなった。
ブラボー！

Go 1.13 からはモジュール・モードが既定になるようで，「GOPATH モードじゃないと困る」とかじゃなければ環境変数 `$GO111MODULE` は `on` のままにしてしまえばいいんじゃないのかなぁ。

あと地味だが嬉しい変更としては，並行処理下で正規表現パッケージ [`regexp`] を使う際に [`regexp`]`.Regexp.Copy()` でクローンを作らなくてもブロッキングが起きないようになった。

[Go 言語] はこの 1.12 から 1.13 にかけて大きく変わる予感がする（今のところ後方互換性は確保されるだろうが）。
色々と試していって慣れていくのがいいかもしれない。

## ブックマーク

- [モジュール対応モードへの移行を検討する]({{< ref "/golang/go-module-aware-mode.md" >}})
- [正規表現に関する戯れ言]({{< ref "/golang/regular-expression.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`regexp`]: https://golang.org/pkg/regexp/ "regexp - The Go Programming Language"

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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/Go%E8%A8%80%E8%AA%9E%E3%81%AB%E3%82%88%E3%82%8B%E4%B8%A6%E8%A1%8C%E5%87%A6%E7%90%86-Katherine-Cox-Buday/dp/4873118468?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4873118468"><img src="https://images-fe.ssl-images-amazon.com/images/I/51pUKQajnaL._SL160_.jpg" width="125" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/Go%E8%A8%80%E8%AA%9E%E3%81%AB%E3%82%88%E3%82%8B%E4%B8%A6%E8%A1%8C%E5%87%A6%E7%90%86-Katherine-Cox-Buday/dp/4873118468?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4873118468">Go言語による並行処理</a></dt>
	<dd>Katherine Cox-Buday</dd>
	<dd>山口 能迪 (翻訳)</dd>
    <dd>オライリージャパン 2018-10-26</dd>
    <dd>Book 単行本（ソフトカバー）</dd>
    <dd>ASIN: 4873118468, EAN: 9784873118468</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">この本を読んで「よっしゃ，明日から立派な goroutine 使いだ！」とはならないと思うけど，有象無象なコピペ・プログラマじゃなく，きちんと Go 言語のプログラミングを勉強したいのであれば，この本は必読書になると思う。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-11-03">2018-11-03</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
