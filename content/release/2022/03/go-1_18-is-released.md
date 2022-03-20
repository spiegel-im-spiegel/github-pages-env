+++
title = "Go 1.18 がリリースされた"
date =  "2022-03-19T13:38:06+09:00"
description = "個人的には workspace mode についてはちゃんと調べて使えるようにしたい"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

既に Gopher たちの間では話題沸騰ですが（笑）， [Go] 1.18 がリリースされました。

- [Go 1.18 is released](https://groups.google.com/g/golang-announce/c/6gJm7mgF6rw)
- [Go 1.18 is released! - The Go Programming Language](https://go.dev/blog/go1.18)
- [Go 1.18 Release Notes - The Go Programming Language](https://go.dev/doc/go1.18)

[ブログ記事](https://go.dev/blog/go1.18 "Go 1.18 is released! - The Go Programming Language")を参考にポイントを紹介すると

- Generics
- Fuzzing
- Workspaces
- 20% Performance Improvements

といったところか。
個人的には [workspace mode](https://go.dev/doc/tutorial/workspaces "Tutorial: Getting started with multi-module workspaces - The Go Programming Language") についてはちゃんと調べて使えるようにしたいと考えている。

{{< fig-quote type="markdown" title="Go 1.18 is released!" link="https://go.dev/blog/go1.18" lang="en" >}}
{{% quote %}}Go modules have been almost universally adopted, and Go users have reported very high satisfaction scores in our annual surveys. In our 2021 user survey, the most common challenge users identified with modules was working across multiple modules. In Go 1.18, we’ve addressed this with a new [Go workspace mode](https://go.dev/doc/tutorial/workspaces), which makes it simple to work with multiple modules{{% /quote %}}.
{{< /fig-quote >}}

Generics は自分では積極的に使おうという気は起きないので，他の人のパッケージに含まれていればお世話になる感じかなぁ。
とりあえず，この辺はいじってみたい。

- [GavinClarke0/lockless-generic-ring-buffer: Single producer and multi-reader lockless ring buffer in go using generics from the go 1.18beta release. It is significantly faster than channels with the added type safety of generics compared to ring buffers using interfaces.](https://github.com/GavinClarke0/lockless-generic-ring-buffer)

例によって [Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.18.linux-amd64.tar.gz`](https://go.dev/dl/go1.18.linux-amd64.tar.gz)）を取ってきてインストールすることを強く推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.18.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.18.linux-amd64.tar.gz
$ sudo mv go go1.18
$ sudo ln -s go1.18 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.18 linux/amd64
```

アップデートは計画的に。

## ブックマーク

- [Big Sky :: Go の http パッケージに MaxBytesHandler が入った。](https://mattn.kaoriya.net/software/lang/go/20211224005655.htm)
- [go1.18で入ったhttp.MaxBytesHandlerの中身を見てみた](https://zenn.dev/hiroyukim/articles/4b4f5b482c0c2d)
- [Go1.18から導入されるnetip package/netip-package - Speaker Deck](https://speakerdeck.com/sonatard/netip-package)
- [What's new in Go 1.18? - Speaker Deck](https://speakerdeck.com/syumai/whats-new-in-go-1-dot-18)
- [What Is the Go Workspace Mode - Speaker Deck](https://speakerdeck.com/110y/what-is-the-go-workspace-mode)
- [[shared] 20220218 Go 1.18 Fuzzing - Google スライド](https://docs.google.com/presentation/d/e/2PACX-1vQ2PX-s-As01o_fvGi9qdx9ZCpQS6RePDWw6rkznN-lI3z8bc4JJ601HLzb1fujo4uf0wSK0Wzl_oc1/pub?resourcekey=0-R72bI85HvGnbMfYmAP_77g#slide=id.g405a9dc47b_0_0)
- [Go 1.18で追加されるstrings/bytes.Cutとsync.Mutex.TryLockについて - Google スライド](https://docs.google.com/presentation/d/1iaEMhXHQa5chIK7Zqqcv6sugXoOEYDQnvldlZlxhJjw/edit#slide=id.p)
- [Go言語のジェネリクス入門(1)](https://zenn.dev/nobishii/articles/type_param_intro)
  - [Go言語のジェネリクス入門(2) インスタンス化と型推論](https://zenn.dev/nobishii/articles/type_param_intro_2)
- [strings.Cut と strings.SplitN はどっちが速いか](https://zenn.dev/mattn/articles/01f258a5127ef8)
- [Big Sky :: text/template と html/template に continue/break が入った。](https://mattn.kaoriya.net/software/lang/go/20210924011409.htm)
- [Go言語がGenericsを導入、過去最大の変更となる「Go 1.18」正式版リリース － Publickey](https://www.publickey1.jp/blog/22/gogenericsgo_118.html)
- [Go 1.18で導入されたnet/netip package](https://zenn.dev/sonatard/articles/92b3ce38e28ee8)
- [Go1.18からのWorkspace modeをさっそく使ってみた](https://zenn.dev/kimuson13/articles/go-workspace-mode-impressions)

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
