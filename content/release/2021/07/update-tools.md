+++
title = "手持ちの #golang 製ツールをアップデート"
date =  "2021-07-13T22:26:49+09:00"
description = "gpgpdump / depm / books-data / ml / cov19jpn"
image = "/images/attention/go-logo_blue.png"
tags  = [ "tools", "golang", "gpgpdump", "books-data" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.16.6 で [HTTPS クライアントに関連する脆弱性の修正]({{< relref "./go-1_16_6-is-released.md" >}})が行われたので，念のため，私のリポジトリで公開している [Go] 製ツール（実行バイナリをリリースしているもの）のいくつかをアップデートした。

- [Release v0.12.4 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.12.4)
- [Release v0.5.10 · spiegel-im-spiegel/books-data · GitHub](https://github.com/spiegel-im-spiegel/books-data/releases/tag/v0.5.10)
- [Release v0.4.4 · spiegel-im-spiegel/depm · GitHub](https://github.com/spiegel-im-spiegel/depm/releases/tag/v0.4.4)
- [Release v0.5.1 · spiegel-im-spiegel/ml · GitHub](https://github.com/spiegel-im-spiegel/ml/releases/tag/v0.5.1)
- [Release v0.2.5 · spiegel-im-spiegel/cov19jpn · GitHub](https://github.com/spiegel-im-spiegel/cov19jpn/releases/tag/v0.2.5)

とりあえず，私がよく使うツールのみ。
ついでに外部依存パッケージのバージョンも上げておいた。
リポジトリに投げっぱなしにしているものも結構あるが，そっちは放置で（笑）

アップデートは計画的に。

## ブックマーク

- [OpenPGP パケットを可視化する gpgpdump]({{< ref "/release/gpgpdump.md" >}})
- [書籍データ取得ツール books-data]({{< ref "/release/books-data.md" >}})
- [Depm: Go 言語用モジュール依存関係可視化ツール]({{< ref "/release/dependency-graph-for-golang-modules.md" >}})
- [spiegel-im-spiegel/ml v0.4.1 をリリースした]({{< ref "/release/2021/01/makelink-0_4_1-is-released.md" >}})
- [手遊びで日本版 Google COVID-19 Forecast データを取得するツールを作ってみた]({{< ref "/release/2021/02/cov19jpn.md" >}})

[Go]: https://go.dev/

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
