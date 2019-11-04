+++
title = "Go 1.13 がリリースされた"
date =  "2019-09-04T22:34:29+09:00"
description = "9月にずれ込んでしまったが，ようやくリリースされた。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

9月にずれ込んでしまったが，ようやく [Go] 1.13 がリリースされた[^google1]。

[^google1]: [同じタイミングで Android 10 がリリースされていた](https://forest.watch.impress.co.jp/docs/news/1205164.html "Google、「Android 10」を正式リリース - 窓の杜")ので関連を疑ったが（笑），どうやら関係ないらしい。 Twitter で教えてもらいました。まぁ Android のバージョンアップスケジュールに付き合わされてはたまらんしねぇ。

- [Go 1.13 is released - Google Group](https://groups.google.com/forum/#!topic/golang-announce/wQqVlKJiMBA)
- [Go 1.13 Release Notes - The Go Programming Language](https://golang.org/doc/go1.13)
- [Go 1.13 is released - The Go Blog](https://blog.golang.org/go1.13)

主な変更点は以下の通り。

{{< fig-quote type="md" title="Go 1.13 is released" link="https://blog.golang.org/go1.13" lang="en" >}}
- The go command now downloads and authenticates modules [using the Go module mirror and Go checksum database by default](https://golang.org/doc/go1.13#introduction)
- [Improvements to number literals](https://golang.org/doc/go1.13#language)
- [Error wrapping](https://golang.org/doc/go1.13#error_wrapping)
- [TLS 1.3 on by default](https://golang.org/doc/go1.13#tls_1_3)
- [Improved modules support](https://golang.org/doc/go1.13#modules)
{{< /fig-quote >}}

この内のいくつかはこのブログでも既に紹介しているので，参考にどうぞ。

- [Go 1.13 のエラー・ハンドリング]({{< ref "/golang/error-handling-in-go-1_3.md" >}})
- [Go 言語の環境変数管理]({{< ref "/golang/go-env.md" >}})
- [Go モジュールのミラーリング・サービス【正式版】]({{< ref "/golang/mirror-index-and-checksum-database-for-go-module.md" >}})

そろそろ [WebAssembly](https://developer.mozilla.org/ja/docs/WebAssembly) にも手を出す季節だろうか。

例によって [Ubuntu] の APT が提供する [Go] コンパイラは3世代も古くて使いものにならないため[^gosup1]，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")から [`go1.13.linux-amd64.tar.gz`](https://dl.google.com/go/go1.13.linux-amd64.tar.gz) を取ってきて任意の場所に展開する。

[^gosup1]: 提供される [Go] コンパイラのサポートは1世代前まで。

```text
$ cd /usr/local/src
$ sudo curl "https://dl.google.com/go/go1.13.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.13.linux-amd64.tar.gz
$ sudo mv go go1.13
$ sudo ln -s go1.13 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.13 linux/amd64
```

ほい。
ひと仕事終わり。

アップデートは計画的に。

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## ブックマーク

- [スライドまとめ on Go 1.13 Release Party in Tokyo（非公式） - Qiita](https://qiita.com/usk81/items/b2803b47e658a905af98)

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
