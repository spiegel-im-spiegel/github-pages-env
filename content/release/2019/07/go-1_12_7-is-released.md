+++
title = "Go 1.12.7 がリリースされた"
date =  "2019-07-09T20:23:45+09:00"
description = "今回もセキュリティ・アップデートはなし。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.12.7 がリリースされた。
セキュリティ・アップデートはなし。

[Go 1.12.7 and Go 1.11.12 are released - Google Group](https://groups.google.com/forum/#!topic/golang-announce/-JidGVRIVEc)

{{< fig-quote type="md" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release.html#go1.12.minor" lang="en" >}}
{{< quote >}}go1.12.7 (released 2019/07/08) includes fixes to cgo, the compiler, and the linker. See the [Go 1.12.7 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.12.7) on our issue tracker for details.{{< /quote >}}
{{< /fig-quote >}}

さらっと書いてるが[マイルストーン](https://github.com/golang/go/issues?q=milestone%3AGo1.12.7)を見ると結構ヤバげな内容なんだよねぇ。


例によって [Ubuntu] の APT が提供する [Go] コンパイラは2世代も古くて使いものにならないため[^gosup1]，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")から [`go1.12.7.linux-amd64.tar.gz`](https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz) を取ってきて任意の場所に展開する。

[^gosup1]: 提供される [Go] コンパイラのサポートは1世代前まで。

```text
$ cd /usr/local/src
$ sudo curl "https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.12.7.linux-amd64.tar.gz
$ sudo mv go go1.12.7
$ sudo ln -s go1.12.7 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.12.7 linux/amd64
```

ほい。
ひと仕事終わり。

アップデートは計画的に。

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

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