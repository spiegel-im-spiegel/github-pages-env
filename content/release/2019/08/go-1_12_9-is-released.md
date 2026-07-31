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

[Go]: https://go.dev/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->

{{% review-paapi "B07VPSXF6N" %}} <!-- 改訂2版 みんなのGo言語 -->
