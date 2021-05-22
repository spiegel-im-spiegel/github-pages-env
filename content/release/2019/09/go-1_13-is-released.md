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

{{< fig-quote type="markdown" title="Go 1.13 is released" link="https://blog.golang.org/go1.13" lang="en" >}}
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

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
