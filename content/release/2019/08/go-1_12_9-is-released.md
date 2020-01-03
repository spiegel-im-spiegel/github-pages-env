+++
title = "Go 1.12.9 がリリースされた"
date =  "2019-08-16T20:56:08+09:00"
description = "今回はセキュリティ・アップデートなし。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.12.9 がリリースされた。
今回はセキュリティ・アップデートはなし。

- [Go 1.12.9 is released - Google Group](https://groups.google.com/forum/#!topic/golang-announce/oeMaeUnkvVE)

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release.html#go1.12.minor" lang="en" >}}
{{< quote >}}go1.12.9 (released 2019/08/15) includes fixes to the linker, and the os and math/big packages. See the [Go 1.12.9 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.12.9) on our issue tracker for details.{{< /quote >}}
{{< /fig-quote >}}

例によって [Ubuntu] の APT が提供する [Go] コンパイラは2世代も古くて使いものにならないため[^gosup1]，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")から [`go1.12.9.linux-amd64.tar.gz`](https://dl.google.com/go/go1.12.9.linux-amd64.tar.gz) を取ってきて任意の場所に展開する。

[^gosup1]: 提供される [Go] コンパイラのサポートは1世代前まで。

```text
$ cd /usr/local/src
$ sudo curl "https://dl.google.com/go/go1.12.9.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.12.9.linux-amd64.tar.gz
$ sudo mv go go1.12.9
$ sudo ln -s go1.12.9 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.12.9 linux/amd64
```

ほい。
ひと仕事終わり。

アップデートは計画的に。

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/B07VPSXF6N?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/51jif840ScL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/B07VPSXF6N?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">改訂2版 みんなのGo言語</a></dt>
    <dd>松木 雅幸 (著), mattn (著), 藤原 俊一郎 (著), 中島 大一 (著), 上田 拓也 (著), 牧 大輔 (著), 鈴木 健太 (著)</dd>
    <dd>技術評論社 2019-08-01 (Release 2019-08-01)</dd>
    <dd>Kindle版</dd>
    <dd>B07VPSXF6N (ASIN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">改訂2版の目玉は7章の「データベースの扱い方」が追加されたことだろう。他の章では，大まかな構成は1版と同じだが細かい部分が変わっていて Go 1.12 への言及まであるのには驚いた。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-08-12">2019-08-12</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
