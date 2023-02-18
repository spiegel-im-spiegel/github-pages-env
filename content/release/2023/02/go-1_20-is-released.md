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

- [Big Sky :: errors.Join が入った。](https://mattn.kaoriya.net/software/lang/go/20221001015441.htm)
- [Providing context to cancellations in Go 1.20 with the new context WithCause API - Joseph Woodward | Software Engineer & Go lover based in Somerset, England](https://josephwoodward.co.uk/2023/01/context-cancellation-cause-with-cancel-cause)
- [Go 1.20 Cryptography](https://words.filippo.io/dispatches/go-1-20-cryptography/)
- [What’s New in Go 1.20, Part I: Language Changes · The Ethically-Trained Programmer](https://blog.carlmjohnson.net/post/2023/golang-120-language-changes/)
  - [Scheduling In Go : Part II - Go Scheduler](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html)
- [What’s New in Go: The Developer Ecosystem Report 2022 | The GoLand Blog](https://blog.jetbrains.com/go/2023/01/17/what-s-new-in-go-the-developer-ecosystem-report-2022/)
- [Go 1.20リリース連載が始まります＆メモリアリーナの紹介＆落ち穂拾い | フューチャー技術ブログ](https://future-architect.github.io/articles/20230123a/)
  - [Go1.20リリース連載 contextパッケージのWithCancelCauseとCause | フューチャー技術ブログ](https://future-architect.github.io/articles/20230125a/)
  - [Go 1.20 Wrapping multiple errors | フューチャー技術ブログ](https://future-architect.github.io/articles/20230126a/)
  - [Go 1.20 timeパッケージのアップデート | フューチャー技術ブログ](https://future-architect.github.io/articles/20230127a/)
  - [Go 1.20 HTTP ResponseController | フューチャー技術ブログ](https://future-architect.github.io/articles/20230128a/)
  - [New ReverseProxy Rewrite hook を動かしながら理解する | フューチャー技術ブログ](https://future-architect.github.io/articles/20230131a/)
  - [Go 1.20 vetのアップデート | フューチャー技術ブログ](https://future-architect.github.io/articles/20230202a/)
  - [Go 1.20 リリース連載 go build に追加される cover オプション（利用例付き） | フューチャー技術ブログ](https://future-architect.github.io/articles/20230203a/)
- [Go 1.20 リリースパーティ - connpass](https://gocon.connpass.com/event/273096/)
- [Go 1.20: Profile-Guided Optimization](https://zenn.dev/mjhd/articles/a09cb5905b7848)
  - [Go言語が実行時のプロファイラ情報でコンパイルを最適化する「Profile-guided optimization」パブリックプレビュー － Publickey](https://www.publickey1.jp/blog/23/goprofile-guided_optimization.html)
- [Go1.20リリースお知らせ記事を翻訳する(前編)](https://zenn.dev/nii/articles/read-go-1-20-release-article)
  - [Go1.20リリースお知らせ記事を翻訳する(後編)](https://zenn.dev/nii/articles/read-go-1-20-release-article2)
- [All your comparable types - The Go Programming Language](https://go.dev/blog/comparable)
  - [Go言語のBasic Interfaceはcomparableを満たすようになる(でも実装するようにはならない)](https://zenn.dev/nobishii/articles/basic-interface-is-comparable)
  - [Goの比較可能性（comparable） - Qiita](https://qiita.com/tenntenn/items/e15cc0c54b3bbfddb04e)

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
