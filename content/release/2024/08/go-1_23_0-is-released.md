+++
title = "Go 1.23 のリリース"
date =  "2024-08-14T08:49:49+09:00"
description = "とりあえずブックマークだけ記しておく。最近いろいろと放ったらかしだなぁ。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

予定通り [Go] 1.23 がリリースされた。
以下，簡単にブックマークだけ記しておく。
最近いろいろと放ったらかしだなぁ。

- [Go 1.23.0 is released](https://groups.google.com/g/golang-announce/c/RQjbRNOcV74)
- [Go 1.23 is released - The Go Programming Language](https://go.dev/blog/go1.23)
- [Go 1.23 Release Notes - The Go Programming Language](https://go.dev/doc/go1.23)

- [Go1.23で導入予定のイテレータを完全理解する✌️](https://zenn.dev/kkkxxx/articles/d9505540581b5d)
- [Go界隈で巻き起こった go:linkname 騒動について - ANDPAD Tech Blog](https://tech.andpad.co.jp/entry/2024/06/20/140000)
  - [cmd/link: lock down future uses of linkname · Issue #67401 · golang/go · GitHub](https://github.com/golang/go/issues/67401)
- [Reduce allocations and comparison performance with the new unique package in Go 1.23 - Joseph Woodward | Software Engineer & Go lover based in Somerset, England](https://josephwoodward.co.uk/2024/08/performance-improvements-unique-package-go-1-23)
- [Range Over Function Types - The Go Programming Language](https://go.dev/blog/range-functions)
- [New unique package - The Go Programming Language](https://go.dev/blog/unique)
- [Go の iter パッケージを使ってみよう](https://zenn.dev/mattn/articles/641f1d86fffdc9)

## Go 1.23.0 の導入

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.23.0.linux-amd64.tar.gz`](https://go.dev/dl/go1.23.0.linux-amd64.tar.gz)）を取ってきてインストールすることを推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.23.0.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.23.0.linux-amd64.tar.gz
$ sudo mv go go1.23.0
$ sudo ln -s go1.23.0 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.23.0 linux/amd64
```

Windows はインストールパッケージを取ってきて直接インストールする。
[Scoop] 経由でも OK

複数バージョンの Go コンパイラを扱いたい場合は

```text
$ go install golang.org/dl/23.0@latest
$ go1.23.0 download
$ go1.23.0 version
go version go1.23.0 linux/amd64
```

てな感じに導入できる。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
{{% review-paapi "B0CFL1DK8Q" %}} <!-- Go言語 100Tips -->
{{% review-paapi "4814400535" %}} <!-- 効率的なGo : Efficient Go -->
