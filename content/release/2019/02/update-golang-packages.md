+++
title = "gpgpdump v0.3.9 他をリリースした"
date = "2019-02-11T16:42:15+09:00"
description = "logf v0.2.5, mklink v0.1.13, gpgpdump v0.3.9 をリリースした。 golangci-lint サマサマという感じである。"
image = "/images/attention/tools.png"
tags  = [ "tools", "package", "golang", "gpgpdump", "openpgp", "logger" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

いやぁ，この3連休は捗ったよ。

[spiegel-im-spiegel/gocli] v0.9.1 のリリースについては[前回お知らせした]({{< ref "/release/2019/02/gocli-v0_9_1-is-released.md" >}} "spiegel-im-spiegel/gocli v0.9.1 をリリースした")が，続けて以下のツールおよび [Go 言語]パッケージをリリースした。

- [Release v0.2.5 · spiegel-im-spiegel/logf · GitHub](https://github.com/spiegel-im-spiegel/logf/releases/tag/v0.2.5)
- [Release v0.1.13 · spiegel-im-spiegel/mklink · GitHub](https://github.com/spiegel-im-spiegel/mklink/releases/tag/v0.1.13)
- [Release v0.3.9 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.3.9)

これらは特に機能追加はないが [lint を golangci-lint に替えた]({{< ref "/golang/donot-sleep-through-life.md" >}} "golangci-lint に叱られる")ので [golangci-lint] の指摘を参考にリファクタリングというかデバッグを行った。
特に [gpgpdump] については結構ヤバい潜在バグも見つかったので [golangci-lint] サマサマという感じである。
やっぱ（たとえ網羅的でなくても書けるところまででも）テストは書いておくべきだよね。

そうそう。
[mklink] の対話モードは [`gocli`]`/prompt` パッケージで全面的に書き換えた。

Flickr の写真を [photo.Baldanders.info](https://photo.baldanders.info/) へ移転するための内部ツールも完成し，とりあえずここのブログの写真（へのリンク）は全部貼り替えた。
これでいよいよ[本家サイト](https://baldanders.info/ "Baldanders.info")の改造に取り掛かれるな。
その前に [Go 言語]コンパイラ v1.12 のリリースがあるが。

## ブックマーク

- [gpgpdump - OpenPGP packet visualizer]({{< ref "/remark/2016/02/gpgpdump-released.md" >}})
    - [gpgpdump 0.3.0 をリリースした]({{< ref "/remark/2017/11/gpgpdump-0_3_0-released.md" >}})
- [Logf Package v0.2.1 Released]({{< ref "/release/2018/03/logf-package-v0_2_1-released.md" >}})
- [Markdown 形式のリンクを生成するツールを作ってみた]({{< ref "/golang/make-link-with-markdown-format.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[golangci-lint]: https://github.com/golangci/golangci-lint "golangci/golangci-lint: Linters Runner for Go. 5x faster than gometalinter. Nice colored output. Can report only new issues. Fewer false-positives. Yaml/toml config."
[spiegel-im-spiegel/gocli]: https://github.com/spiegel-im-spiegel/gocli "spiegel-im-spiegel/gocli: Minimal Packages for Command-Line Interface"
[`gocli`]: https://github.com/spiegel-im-spiegel/gocli "spiegel-im-spiegel/gocli: Minimal Packages for Command-Line Interface"
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[mklink]: https://github.com/spiegel-im-spiegel/mklink "spiegel-im-spiegel/mklink: Make Link with Markdown Format"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
