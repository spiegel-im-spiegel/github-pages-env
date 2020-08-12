+++
title = "Go 1.15 がリリースされた"
date =  "2020-08-12T11:04:16+09:00"
description = "今回はエラい早いリリースだったねぇ。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.15 がリリースされた。
今回はエラい早いリリースだったねぇ（1.13 のときは[8月中のリリースに間に合わなかった]({{< ref "/release/2019/09/go-1_13-is-released.md" >}} "Go 1.13 がリリースされた")）。

- [Go 1.15 is released](https://groups.google.com/g/golang-announce/c/Z-cY6ZdGdEU)
- [Go 1.15 Release Notes - The Go Programming Language](https://golang.org/doc/go1.15)
- [Go 1.15 is released - The Go Blog](https://blog.golang.org/go1.15)

主な変更点は以下の通り。

{{< fig-quote type="markdown" title="Go 1.15 is released" link="https://blog.golang.org/go1.15" lang="en" >}}
- [Substantial improvements to the Go linker](https://golang.org/doc/go1.15#linker)
- [Improved allocation for small objects at high core counts](https://golang.org/doc/go1.15#runtime)
- [X.509 CommonName deprecation](https://golang.org/doc/go1.15#commonname)
- [GOPROXY supports skipping proxies that return errors](https://golang.org/doc/go1.15#go-command)
- [New embedded tzdata package](https://golang.org/doc/go1.15#time/tzdata)
- [A number of Core Library improvements](https://golang.org/doc/go1.15#library)
{{< /fig-quote >}}

例によって [Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.15.linux-amd64.tar.gz`](https://golang.org/dl/go1.15.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを推奨する。

```text
$ cd /usr/local/src
$ sudo curl -L "https://golang.org/dl/go1.15.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.15.linux-amd64.tar.gz
$ sudo mv go go1.15
$ sudo ln -s go1.15 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.15 linux/amd64
```

アップデートは計画的に。

[Go]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
