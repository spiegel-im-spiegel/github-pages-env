+++
title = "spiegel-im-spiegel/ml v0.4.1 をリリースした"
date =  "2021-01-17T13:32:11+09:00"
description = "ヒストリ機能を簡易リングバッファに切り替えた。"
image = "/images/attention/tools.png"
tags  = [ "tools", "golang" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Markdown 形式のリンクを生成する [spiegel-im-spiegel/ml][ml] v0.4.0 および v0.4.1 をリリースした。

- [Release v0.4.0 · spiegel-im-spiegel/ml · GitHub](https://github.com/spiegel-im-spiegel/ml/releases/tag/v0.4.0)
- [Release v0.4.1 · spiegel-im-spiegel/ml · GitHub](https://github.com/spiegel-im-spiegel/ml/releases/tag/v0.4.1)

いや， v0.4.0 をリリースした後にエラーハンドリングをしくじってるのに気が付いて出し直し（笑）

対話モード（`-i` オプション）のヒストリ機能を簡易リングバッファに切り替えた。
あと `-l` オプションでヒストリ数を指定できる。
既定は `0`，つまりヒストリ機能を無効にしている。
またヒストリ情報を `$XDG_CONFIG_HOME/ml/history` ファイルに保存して再利用できるようにした。
ちなみに Windows 版では `%APPDATA%\ml\hisotry` ファイル， macOS では `/Library/Application Support/ml/history` ファイルとなる（筈）。

[ml]: https://github.com/spiegel-im-spiegel/ml "spiegel-im-spiegel/ml: Make Link with Markdown Format"

## 参考図書

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
