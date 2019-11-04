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
