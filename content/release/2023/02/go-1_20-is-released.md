+++
title = "Go 1.20 がリリースされた"
date =  "2023-02-02T20:15:43+09:00"
description = "とりあえずマルチエラーに関してはちゃんと調査しないと。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

まさか，2月早々に出るとか！

- [Go 1.20 is released](https://groups.google.com/g/golang-announce/c/QMK8IQALDvA)
- [Go 1.20 is released! - The Go Programming Language](https://go.dev/blog/go1.20)
- [Go 1.20 Release Notes - The Go Programming Language](https://go.dev/doc/go1.20)

その他のリンクについては後日に補完する。
とりあえず[マルチエラー]({{< ref "/golang/wrapping-multiple-errors.md" >}} "【Go 1.20 の予習】複数 error を返す Unwrap メソッドについて")に関してはちゃんと調査しないと。

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.20.linux-amd64.tar.gz`](https://go.dev/dl/go1.20.linux-amd64.tar.gz)）を取ってきてインストールすることを強く推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.20.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.20.linux-amd64.tar.gz
$ sudo mv go go1.20
$ sudo ln -s go1.20 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.20 linux/amd64
```

Windows は [Scoop] 経由で OK

アップデートは計画的に。

## ブックマーク

- [Go 1.20 リリースパーティ - connpass](https://gocon.connpass.com/event/273096/)

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
{{% review-paapi "4873119693" %}} <!-- 実用 Go 言語 -->
{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
{{% review-paapi "4908686122" %}} <!-- Goならわかるシステムプログラミング 第2版 -->
{{% review-paapi "B09C2XBC2F" %}} <!-- Golang Tシャツ -->
