+++
title = "gpgpdump v0.3.9 他をリリースした"
date = "2019-02-11T16:42:15+09:00"
description = "logf v0.2.5, mklink v0.1.13, gpgpdump v0.3.9 をリリースした。 golangci-lint サマサマという感じである。"
image = "/images/attention/tools.png"
tags  = [ "tools", "package", "golang", "gpgpdump", "openpgp", "logger" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"

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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253"><img src="https://images-fe.ssl-images-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan, Brian W. Kernighan</dd>
    <dd>柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>Book 単行本（ソフトカバー）</dd>
    <dd>ASIN: 4621300253, EAN: 9784621300251</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
