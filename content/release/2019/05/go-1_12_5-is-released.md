+++
title = "Go 1.12.5 がリリースされた"
date =  "2019-05-07T22:05:00+09:00"
description = "セキュリティ・アップデートはなし。ついでに Ubuntu へのインストールについても言及しておく。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "install" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.12.5 がリリースされた。
セキュリティ・アップデートはなし。

- [Go 1.12.5 and Go 1.11.10 are released - Google Group](https://groups.google.com/forum/#!topic/golang-announce/GhnThAAITFI)

{{< fig-quote title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release.html#go1.12.minor" lang="en" >}}
<q>go1.12.5 (released 2019/05/06) includes fixes to the compiler, the linker, the go command, the runtime, and the os package. See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.12.5">Go 1.12.5 milestone</a> on our issue tracker for details.</q>
{{< /fig-quote >}}

ちなみに [Ubuntu] の APT (Advanced Package Tool) で [Go] コンパイラをインストールすると2世代も古いのが入る（モジュール・モードが使えないしサポートからも外れていると思う）。

```text
$ apt show golang
Package: golang
Version: 2:1.10~4ubuntu1
Priority: optional
Section: devel
Source: golang-defaults
Origin: Ubuntu
...
```

ちうわけで[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")から [`go1.12.5.linux-amd64.tar.gz`](https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz) とかを取ってきて任意の場所に手動で展開するほうが吉である。
たとえば，こんな感じ。

```text
$ cd /usr/local/src
$ sudo curl https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.12.5.linux-amd64.tar.gz
$ sudo mv go go1.12.5
$ sudo ln -s go1.12.5 go
$ ./go/bin/go version
go version go1.12.5 linux/amd64
```

これで `/usr/local/go/bin` にパスを通してやればよい。
たとえば `/etc/profile.d` ディレクトリに `golang-bin-path.sh` とかいった名前でファイルを作って（名前は適当）

```bash
# shellcheck shell=sh

# Expand $PATH to include the directory where golang applications go.
golang_bin_path="/usr/local/go/bin"
if [ -n "${PATH##*${golang_bin_path}}" -a -n "${PATH##*${golang_bin_path}:*}" ]; then
    export PATH=$PATH:${golang_bin_path}
fi
```

とでも書いておけば次回ログイン時にはパスが通っている。

```text
$ go version
go version go1.12.5 linux/amd64
```

環境変数については `GO111MODULE` 以外は設定する必要はない。
ちなみに `GOPATH` は既定で

```text
$ go env | grep GOPATH
GOPATH="/home/username/go"
```

となっている。

[Go]: https://go.dev/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
