+++
title = "Go 1.19 がリリースされた"
date =  "2022-08-03T08:11:29+09:00"
description = "お盆過ぎに出ればいいかと油断してた（笑）"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

続いて [Go] 1.19 リリースの話。
いや，マジで「お盆過ぎに出ればいいか」と完全に油断してたよ（笑）

- [Go 1.19 is released](https://groups.google.com/g/golang-announce/c/hwjGDuSVZJo)
- [Go 1.19 is released! - The Go Programming Language](https://go.dev/blog/go1.19)
- [Go 1.19 Release Notes - The Go Programming Language](https://go.dev/doc/go1.19)

ここ半年ほど [Go] から離れ気味だったので 1.19 については全く情報収集してなかった。
なので，後日追記することになるだろう。
とりあえずリリパには間に合ってよかったね。

- [Go 1.19 リリースパーティ - connpass](https://gocon.connpass.com/event/253355/)


## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.19.linux-amd64.tar.gz`](https://go.dev/dl/go1.19.linux-amd64.tar.gz)）を取ってきてインストールすることを強く推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.19.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.19.linux-amd64.tar.gz
$ sudo mv go go1.19
$ sudo ln -s go1.19 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.19 linux/amd64
```

Windows は [Scoop] 経由で OK

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/

## ブックマーク

- [The Go Memory Model - The Go Programming Language](https://tip.golang.org/ref/mem)
- [Go 1.19リリース連載始まります GoDoc/ツール周りのアップデート | フューチャー技術ブログ](https://future-architect.github.io/articles/20220801a/)
  - [Go1.19 encoding/csv のアップデート | フューチャー技術ブログ](https://future-architect.github.io/articles/20220802a/)
  - [Go1.19で追加されたAppend系メソッド | フューチャー技術ブログ](https://future-architect.github.io/articles/20220803a/)
  - [Go1.19 net/http のアップデート | フューチャー技術ブログ](https://future-architect.github.io/articles/20220804a/)
  - [Go 1.19 Genericsのアップデート | フューチャー技術ブログ](https://future-architect.github.io/articles/20220805a/)
  - [Go 1.19のメモリ周りの更新 | フューチャー技術ブログ](https://future-architect.github.io/articles/20220808a/)
- [Go1.19で採用された Pattern-defeating Quicksort の紹介 - Speaker Deck](https://speakerdeck.com/po3rin/go1-dot-19decai-yong-sareta-pattern-defeating-quicksort-falseshao-jie)
- [[shared] 20220815 What's new in Go 1.19? - Google スライド](https://docs.google.com/presentation/d/1FkXdI9oR8mUCzh-woca7O3K_T5iZCirp7QcoJY3d4Wk/edit)

- [Go 1.19 で os/exec パッケージの挙動が変わった話]({{< ref "/golang/exec-package-in-go119.md" >}})

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B09P4PH63R" %}} <!-- エキスパートたちのGo言語 -->
{{% review-paapi "4873119693" %}} <!-- 実用 Go 言語 -->
{{% review-paapi "B0B62K55SL" %}} <!-- 詳解Go言語Webアプリケーション開発 -->
{{% review-paapi "4873119979" %}} <!-- Go言語による分散サービス -->
{{% review-paapi "B09C2XBC2F" %}} <!-- Golang Tシャツ -->
