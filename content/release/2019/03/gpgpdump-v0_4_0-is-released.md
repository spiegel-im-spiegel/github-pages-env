+++
title = "gpgpdump v0.4.0 をリリースした"
date = "2019-03-10T20:10:20+09:00"
description = "マイナー・バージョンアップだが特に機能追加があるわけではなく，ちょっとした仕様変更のみ。"
image = "/images/attention/openpgp.png"
tags = ["tools", "openpgp", "golang", "gpgpdump"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP] パケットの内容を視覚化する [gpgpdump] の v0.4.0 をリリースした。

- [Release v0.4.0 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.4.0)

マイナー・バージョンアップだが特に機能追加があるわけではなく

- エラー・ハンドリングの変更
- [`gpgpdump`]`/options` サブパッケージのインタフェース変更

という感じで内部の仕様を変更したので。

エラー・ハンドリングについては以前に書いた

- [エラー・ハンドリング再考]({{< ref "/golang/error-handling-again.md" >}})

を実践している。

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[`gpgpdump`]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[pgpdump]: http://www.mew.org/~kazu/proj/pgpdump/ "pgpdump"
[OpenPGP]: http://openpgp.org/
