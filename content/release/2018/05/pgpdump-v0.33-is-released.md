+++
title = "pgpdump 0.33 がリリース"
date = "2018-05-07T22:49:18+09:00"
description = "例によってソースコードのみのリリースだが Windows 用にコンパイルしたものを用意した。"
image = "/images/attention/tools.png"
tags  = [ "openpgp", "tools", "pgpdump" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[pgpdump] は山本和彦さんによる [OpenPGP] パケットの[視覚化ツール](http://www.mew.org/~kazu/proj/pgpdump/ja/)である。

v0.33 での変更点を [git のログからさらう]({{< ref "/remark/2018/03/git-log.md" >}} "バージョン間のコミット・ログを取得する")とこんな感じ。

- Make the haskell test driver use UTC. 6118d4a
- Convert tests to UTC so they work in not-JST locales. c03c140
- Add simple shell test harness without massive haskell dependencies. 2acddc6
- Updated public key algorithms 66ee31c

楕円曲線暗号に一部対応しているようだ。
あとはテスト周りかな。

例によってソースコードのみのリリースだが，こちらで Windows 用にコンパイルしたものを置いている。

- [pgpdump (patched version) — Baldanders.info](http://www.baldanders.info/spiegel/archive/pgpdump/)

とりあえず暗号アルゴリズムの名前だけ表示できるようにしたって感じだが，作者の方は既に Haskell の人なので積極的にメンテナンスする気はないんだろうな。

## ブックマーク

- [MSYS2 による gcc 開発環境の構築 ― pgpdump をビルドする]({{< ref "/remark/2016/03/gcc-msys2-3.md" >}})
- [gpgpdump 0.3.0 をリリースした]({{< ref "/remark/2017/11/gpgpdump-0_3_0-released.md" >}})

[OpenPGP]: http://openpgp.org/
[pgpdump]: https://github.com/kazu-yamamoto/pgpdump "kazu-yamamoto/pgpdump: A PGP packet visualizer"
