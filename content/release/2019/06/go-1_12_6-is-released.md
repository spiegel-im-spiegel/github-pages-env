+++
title = "Go 1.12.6 がリリースされた"
date =  "2019-06-12T22:01:59+09:00"
description = "今回もセキュリティ・アップデートはなし。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.12.6 がリリースされた。
セキュリティ・アップデートはなし。

- [Go 1.12.6 and Go 1.11.11 are released - Google Group](https://groups.google.com/forum/#!topic/golang-announce/dNU0sAdX65I)

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release.html#go1.12.minor" lang="en" >}}
{{< quote >}}go1.12.6 (released 2019/06/11) includes fixes to the compiler, the linker, the go command, and the `crypto/x509`, `net/http`, and `os` packages. See the [Go 1.12.6 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.12.6) on our issue tracker for details.{{< /quote >}}
{{< /fig-quote >}}

[Ubuntu] の場合は（APT で提供される [Go] コンパイラは古すぎるので）[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")から [`go1.12.6.linux-amd64.tar.gz`](https://dl.google.com/go/go1.12.6.linux-amd64.tar.gz) を取ってきて任意の場所に手動で展開するほうがお勧めである。

たとえば，こんな感じ。

```text
$ cd /usr/local/src
$ sudo curl https://dl.google.com/go/go1.12.6.linux-amd64.tar.gz -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.12.6.linux-amd64.tar.gz
$ sudo mv go go1.12.6
$ sudo ln -s go1.12.6 go
$ ./go/bin/go version
go version go1.12.6 linux/amd64
```

アップデートは計画的に。

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
