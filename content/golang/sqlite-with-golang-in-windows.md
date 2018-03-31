+++
title = "Go 言語で SQLite を使う（Windows 向けの紹介）"
date =  "2018-03-30T16:03:11+09:00"
description = "「XXX」他"
image = "/images/attention/go-code2.png"
tags  = [ "golang", "package", "sqlite", "gcc" ]
draft = true

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

[SQLite] はアプリケーションに組み込み可能な簡易 RDBMS で，公有（public domain）のソフトウェアとして公開されている。

C 言語で書かれているため多くのプラットフォームまたは（C 言語とバインド可能な）多くのプログラミング言語で利用可能である。
コア部分のコードが小さいため組込みソフトウェアでも使われたりする。

[Go 言語]でも利用可能だがコンパイル時に [GCC] の C コンパイラが必要である。
Linux などのプラットフォームには当然 [GCC] が入っているが，残念ながら Windows 環境には入ってないので別途用意する必要がある[^sub1]。

[^sub1]: Windows 10 用の Linux 系サブシステムになら入ってると思うが，私は Windows 10 を使ったことがないのでよく知らない。

Windows 環境で [GCC] **のみ** が必要なのであれば [TDM-GCC] がおススメである[^gcc1]。

[^gcc1]: ただし [TDM-GCC] のバージョンは 5.1 系でかなり古いので注意が必要である（どうも [TDM-GCC] は開発が止まってるっぽい？）。なお，今回は [TDM-GCC] で全く問題ない。ちなみに [MSYS2] で提供される [Mingw-w64] 版 gcc は32ビット版・64ビット版ともに 7.3 が提供されている（2018-02-31 時点）。また [GCC] だけでなく autotools などの周辺ツールも必要なら [MSYS2] を使うべきだろう。参照： [MSYS2 による gcc 開発環境の構築 ― gcc パッケージ群の導入]({{< relref "remark/2016/03/gcc-msys2-2.md" >}})

まぁそれを言ったら [MSYS2] に含まれる [GCC] は 4.8 系で更に古いけどね（笑）

[Go 言語]用の [SQLite] パッケージはいくつか存在するが標準の database/[sql] に対応しているのは [mattn/go-sqlite3] のみのようだ。

- [mattn/go-sqlite3: sqlite3 driver for go using database/sql](https://github.com/mattn/go-sqlite3)


```text
$ go get -v github.com/mattn/go-sqlite3
github.com/mattn/go-sqlite3
# github.com/mattn/go-sqlite3
exec: "gcc": executable file not found in %PATH%
```


## ブックマーク

- [DB Browser for SQLite](http://sqlitebrowser.org/)
    - [「DB Browser for SQLite」“SQLite”のデータベースを管理できるソフト - 窓の杜ライブラリ](https://forest.watch.impress.co.jp/library/software/sqldbbrowser/)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[SQLite]: https://www.sqlite.org/
[GCC]: https://gcc.gnu.org/ "GCC, the GNU Compiler Collection - GNU Project - Free Software Foundation (FSF)"
[Mingw-w64]: http://mingw-w64.org/doku.php "Mingw-w64 - GCC for Windows 64 & 32 bits [mingw-w64]"
[TDM-GCC]: http://tdm-gcc.tdragon.net/
[MSYS2]: http://www.msys2.org/
[sql]: https://golang.org/pkg/database/sql/ "sql - The Go Programming Language"
[mattn/go-sqlite3]: https://github.com/mattn/go-sqlite3 "mattn/go-sqlite3: sqlite3 driver for go using database/sql"
