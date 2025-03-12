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
- [Extensible Wasm Applications with Go - The Go Programming Language](https://go.dev/blog/wasmexport)
- [Testing concurrent code with testing/synctest - The Go Programming Language](https://go.dev/blog/synctest)
- [Faster Go maps with Swiss Tables - The Go Programming Language](https://go.dev/blog/swisstable)
- [Traversal-resistant file APIs - The Go Programming Language](https://go.dev/blog/osroot)

パッと見たところ暗号周りが派手に変わってる感じ。
あとは [`encoding`]`.TextAppender` と [`encoding`]`.BinaryAppender` かな。
これを使ったリファクタリングが広範に行われている印象。
この辺は後で調べてみよう。

## Go 1.24.1 の導入 【2025-03-05 更新】

{{< div-box type="markdown" >}}
[Go] 1.24.1 がリリースされた。

- [[security] Go 1.24.1 and Go 1.23.7 are released](https://groups.google.com/g/golang-announce/c/4t3lzH3I0eI)

CVE ID ベースで1件の脆弱性修正がある。

[Go]: https://go.dev/
{{< /div-box >}}

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.24.1.linux-amd64.tar.gz`](https://go.dev/dl/go1.24.1.linux-amd64.tar.gz)）を取ってきてインストールすることを推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.24.1.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.24.1.linux-amd64.tar.gz
$ sudo mv go go1.24.1
$ sudo ln -s go1.24.1 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.24.1 linux/amd64
```

Windows はインストールパッケージを取ってきて直接インストールする。
[Scoop] 経由でも OK

複数バージョンの Go コンパイラを扱いたい場合は

```text
$ go install golang.org/dl/24.1@latest
$ go1.24.1 download
$ go1.24.1 version
go version go1.24.1 linux/amd64
```

てな感じに導入できる。

## ブックマーク

- [CA.go ~ Deep dive into Go1.24 ~ - connpass](https://cyberagent.connpass.com/event/342451/)
- [Go 1.24 リリースパーティ - connpass](https://gocon.connpass.com/event/345795/)
- [渋川よしき氏・mattn氏に聞く、春のGoお悩み相談会 - connpass](https://levtechlab.connpass.com/event/339320/)
- [Go 1.24でジェネリックになった型エイリアスの紹介 - Speaker Deck](https://speakerdeck.com/syumai/go-1-dot-24dezieneritukuninatutaxing-eiriasunoshao-jie)
- [Go 1.24で入ったGo製ツールの管理機能が便利だったのでおすすめしたい - 焼売飯店](https://blog.syum.ai/entry/2025/03/01/235814)

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/
[`encoding`]: https://pkg.go.dev/encoding "encoding package - encoding - Go Packages"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400535" %}} <!-- 効率的なGo : Efficient Go -->
{{% review-paapi "B0DNYMMBBQ" %}} <!-- Go言語で学ぶ並行プログラミング -->
