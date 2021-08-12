+++
title = "Go 1.14 がリリースされた"
date =  "2020-02-26T12:28:32+09:00"
description = "アップデートは計画的に。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.14 がリリースされた。
今回は2月中にリリースできてよかったね（笑）

- [Go 1.14 is released - Google group](https://groups.google.com/forum/#!topic/golang-announce/AYd4cXYG8qk)
- [Go 1.14 is released - The Go Blog](https://blog.golang.org/go1.14)
- [Go 1.14 Release Notes - The Go Programming Language](https://golang.org/doc/go1.14)

主な変更点は以下の通り。

{{< fig-quote type="markdown" title="Go 1.14 is released - The Go Blog" link="https://blog.golang.org/go1.14" lang="en" >}}
- Module support in the go command is now ready for production use. We encourage all users to [migrate to go modules for dependency management](https://golang.org/doc/go1.14#introduction).
- [Embedding interfaces with overlapping method sets](https://golang.org/doc/go1.14#language)
- [Improved defer performance](https://golang.org/doc/go1.14#runtime)
- [Goroutines are asynchronously preemptible](https://golang.org/doc/go1.14#runtime)
- [The page allocator is more efficient](https://golang.org/doc/go1.14#runtime)
- [Internal timers are more efficient](https://golang.org/doc/go1.14#runtime)
{{< /fig-quote >}}

個人的に気になっている点は以前「[Go 1.14 リリース候補版]({{< relref "./go1_14-rc.md" >}})」に書いているので，参考にどうぞ。

例によって [Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.14.linux-amd64.tar.gz`](https://dl.google.com/go/go1.14.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する。

```text
$ cd /usr/local/src
$ sudo curl "https://dl.google.com/go/go1.14.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.14.linux-amd64.tar.gz
$ sudo mv go go1.14
$ sudo ln -s go1.14 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.14 linux/amd64
```

アップデートは計画的に。



## ブックマーク

- [Changes to interfaces in Go1.14 - Dylan Meeus - Medium](https://medium.com/@meeusdylan/changes-to-interfaces-in-go1-14-821ab533f62)
- [A Go Module Testbed « null program](https://nullprogram.com/blog/2020/02/13/)
- [Inlined defers in Go · Go, the unwritten parts](https://rakyll.org/inlined-defers/)
- [Go1.14で来るGo Modules関連の変更を見てみる - Qiita](https://qiita.com/pi9min/items/9dac44c6663e0e933968)
- [Go 1.14 リリースノート 日本語訳 - Qiita](https://qiita.com/c-yan/items/3dd0c63ea4c3041728cc)
- [Go1.14のcontextは何が変わるのか - Qiita](https://qiita.com/tutuz/items/963a6118cec63a4cd2f3)
- [What's new Go 1.14?](https://talks.godoc.org/github.com/qt-luigi/talks/2020/whats-new-go-114.slide)

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
