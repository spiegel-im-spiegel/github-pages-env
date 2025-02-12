+++
title = "Go 1.24 のリリース"
date =  "2025-02-12T19:28:13+09:00"
description = "パッと見たところ暗号周りが派手に変わってる感じ。 あとは encoding.TextAppender と encoding.BinaryAppender かな。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

予定通り [Go] 1.24 がリリースされた。
以下，簡単にブックマークだけ記しておく（随時追加予定）。

- [Go 1.24.0 is released](https://groups.google.com/g/golang-announce/c/_G2hEiKx8SE)
- [Go 1.24 is released! - The Go Programming Language](https://go.dev/blog/go1.24)
- [Go 1.24 Release Notes - The Go Programming Language](https://go.dev/doc/go1.24)
- [FIPS 140-3 Compliance - The Go Programming Language](https://go.dev/doc/security/fips140)

パッと見たところ暗号周りが派手に変わってる感じ。
あとは [`encoding`]`.TextAppender` と [`encoding`]`.BinaryAppender` かな。
これを使ったリファクタリングが広範に行われている印象。
この辺は後で調べてみよう。

## Go 1.24.0 の導入

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.24.0.linux-amd64.tar.gz`](https://go.dev/dl/go1.24.0.linux-amd64.tar.gz)）を取ってきてインストールすることを推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.24.0.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.24.0.linux-amd64.tar.gz
$ sudo mv go go1.24.0
$ sudo ln -s go1.24.0 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.24.0 linux/amd64
```

Windows はインストールパッケージを取ってきて直接インストールする。
[Scoop] 経由でも OK

複数バージョンの Go コンパイラを扱いたい場合は

```text
$ go install golang.org/dl/24.0@latest
$ go1.24.0 download
$ go1.24.0 version
go version go1.24.0 linux/amd64
```

てな感じに導入できる。

## ブックマーク

- [渋川よしき氏・mattn氏に聞く、春のGoお悩み相談会 - connpass](https://levtechlab.connpass.com/event/339320/)

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/
[`encoding`]: https://pkg.go.dev/encoding "encoding package - encoding - Go Packages"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400535" %}} <!-- 効率的なGo : Efficient Go -->
{{% review-paapi "B0DNYMMBBQ" %}} <!-- Go言語で学ぶ並行プログラミング -->
