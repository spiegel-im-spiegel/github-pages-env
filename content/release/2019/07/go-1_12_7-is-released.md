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

- [Go 1.12.7 and Go 1.11.12 are released - Google Group](https://groups.google.com/forum/#!topic/golang-announce/-JidGVRIVEc)

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release.html#go1.12.minor" lang="en" >}}
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

[Go]: https://go.dev/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
