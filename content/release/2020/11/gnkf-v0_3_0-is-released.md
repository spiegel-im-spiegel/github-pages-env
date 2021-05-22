+++
title = "GNKF: NKF ぽいなにか の v0.3.0 をリリースした"
date =  "2020-11-28T17:57:55+09:00"
description = "このバージョンで UTF-8 テキスト中の BOM を除去する機能を追加した。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "tools", "gnkf", "golang", "unicode" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go 言語][Go]における文字列処理の習作 [gnkf] の v0.3.0 をリリースした。

- [Release v0.3.0 · spiegel-im-spiegel/gnkf · GitHub](https://github.com/spiegel-im-spiegel/gnkf/releases/tag/v0.3.0)

このバージョンで UTF-8 テキスト中の BOM (Byte Order Mark; U+FEFF) を除去する機能を追加した。

```text
$ gnkf remove-bom -h
Remove BOM character in UTF-8 string.

Usage:
  gnkf remove-bom [flags]

Aliases:
  remove-bom, rbom, rb

Flags:
  -f, --file string     path of input text file
  -h, --help            help for remove-bom
  -o, --output string   path of output file

Global Flags:
      --debug   for debug
```

たとえば

```text
$ echo ﻿Hello | gnkf dump
0xef, 0xbb, 0xbf, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x0a
```

こんな感じに BOM が付いてるテキストに対し

```text
$ echo ﻿Hello | gnkf rb | gnkf dump
0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x0a
```

てな感じで BOM を除去できる。

（[東京都の COVID-2019 オープンデータ](https://catalog.data.metro.tokyo.lg.jp/dataset/t000010d0000000068 "東京都 新型コロナウイルス陽性患者発表詳細 - データセット - 東京都オープンデータカタログサイト")とか） CSV ファイル等に BOM が付いてるものがあって，前々から「誰がこんな要らんことしてんねん」と思っていたが，仕事で再び MS Office ドキュメントを弄るようになって気がついた。
**おまえか！**

ことあるごとに言っているが， UTF-8 はオクテット単位のストリームなのでエンディアンを指示する BOM は不要である。
てか，付けるべきではない。
大きさを持たないかつ意味のない文字コードを付加するのはリスクである。

昔は文字列の先頭に BOM があるかどうかで文字エンコーディングを判定する手抜き実装もあったが，流石にそんな頭の悪いシステムはなくなっただろう。

なお `enc`, `norm`, `width`, `kana` の各サブコマンドでも `--remove-bom` オプション（短縮名 `-b`）を付けることで，前処理として BOM の除去を行うようにした。

全文をいったんメモリに読み込んで変換処理を行っているので，巨大テキストを処理する際はご注意を。
習作（study）として作ったツールを仕事で使う羽目になるとは `orz`

ホンマ UTF-8 の BOM は滅びればいいのに。

[Go]: https://golang.org/ "The Go Programming Language"
[gnkf]: https://github.com/spiegel-im-spiegel/gnkf "spiegel-im-spiegel/gnkf: Network Kanji Filter by Golang"

## ブックマーク

- [GNKF: Network Kanji Filter by Golang]({{< ref "/release/gnkf.md" >}})

## 参考図書

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
