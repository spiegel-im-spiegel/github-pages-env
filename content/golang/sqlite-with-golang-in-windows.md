+++
title = "Go 言語で SQLite を使う（Windows 向けの紹介）"
date = "2018-03-31T21:12:23+09:00"
description = "今回はパッケージの紹介のみ。つか，Windows で cgo を使うための覚え書きのようなものか。"
tags  = [ "golang", "package", "sqlite", "gcc", "windows" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

今回はパッケージの紹介のみ。
つか， Windows で cgo を使うための覚え書きのようなものか。

[SQLite] はアプリケーションに組み込み可能な簡易 RDBMS で，公有（public domain）のソフトウェアとして公開されている。
C 言語で書かれているため多くのプラットフォームまたは（C 言語とバインド可能な）多くのプログラミング言語で利用可能である。
コア部分のコードが（他の製品に比べて）小さいため組込みソフトウェアで使われることもある。

## mattn/go-sqlite3

[Go 言語]で利用可能な [SQLite] パッケージはいくつか存在するが，標準の database/[sql] に対応しているのは [mattn/go-sqlite3] のみのようだ。

- [mattn/go-sqlite3: sqlite3 driver for go using database/sql](https://github.com/mattn/go-sqlite3)

このパッケージを利用するには [GCC] が必要である（内部で C 言語コードのコンパイルを行うため）。
[GCC] がない状態で `go get` しようとすると以下のようにエラーになる。

```text
$ go get -v github.com/mattn/go-sqlite3
github.com/mattn/go-sqlite3
# github.com/mattn/go-sqlite3
exec: "gcc": executable file not found in %PATH%
```

なお，必要なコードは [mattn/go-sqlite3] に組み込まれているため [SQLite] サイトからソースコードや DLL などのバイナリを別途取ってくる必要はない[^ver1]。

[^ver1]: [mattn/go-sqlite3] に組み込まれている [SQLite] のバージョンは 2018-03-31 時点で 3.22.0 のようだ。

### GCC の導入

Linux などのプラットフォームには最初から [GCC] が入っているが， Windows 環境には残念ながら入ってないので別途用意する必要がある[^sub1]。
Windows 環境で [GCC] **のみ** が必要なのであれば [MinGW-w64] から Windows 用のバイナリを取得するのがお勧めである[^gcc1]。

[^sub1]: Windows 10 用の Linux 系サブシステムになら入ってると思うが，私は Windows 10 を使ったことがないのでよく知らない。
[^gcc1]: [GCC] だけでなく autotools などの周辺ツールも必要なら [MSYS2] を導入するほうがいいかもしれない（参考： [MSYS2 による gcc 開発環境の構築 ― gcc パッケージ群の導入]({{< ref "/remark/2016/03/gcc-msys2-2.md" >}})）。今回は [MinGW-w64] で全く問題ない。

- [Mingw-w64 を導入する]({{< ref "/remark/2018/03/mingw-w64.md" >}})

## サンプルをコンパイルしてみる

[mattn/go-sqlite3] に `_example/simple/simple.go` というサンプルファイルがあるので，これを動かして動作確認してみる。

```text
$ go run simple.go
```

これで `foo.db` ファイルができていたら成功だ。
適当なブラウザツールで中身を確認してみるといいだろう。

あぁ [SQLite] 用の SQL 方言を覚えないと。
なんで製品ごとに SQL の方言がこんなに微妙な感じなんだろうねぇ。
特定の製品にロックインさせるための陰謀なのだろうか（笑）

## ブックマーク

- [DB Browser for SQLite](http://sqlitebrowser.org/)
    - [「DB Browser for SQLite」“SQLite”のデータベースを管理できるソフト - 窓の杜ライブラリ](https://forest.watch.impress.co.jp/library/software/sqldbbrowser/)
- [SQLiteのVACUUMメモ | Siguniang's Blog](https://siguniang.wordpress.com/2012/11/10/notes-on-sqlite-vacuum/) : pragma を使って自動で vacuum できるらしい

[Go 言語]: https://golang.org/ "The Go Programming Language"
[SQLite]: https://www.sqlite.org/
[GCC]: https://gcc.gnu.org/ "GCC, the GNU Compiler Collection - GNU Project - Free Software Foundation (FSF)"
[MinGW-w64]: http://mingw-w64.org/ "Mingw-w64 - GCC for Windows 64 & 32 bits [mingw-w64]"
[MSYS2]: http://www.msys2.org/
[sql]: https://golang.org/pkg/database/sql/ "sql - The Go Programming Language"
[mattn/go-sqlite3]: https://github.com/mattn/go-sqlite3 "mattn/go-sqlite3: sqlite3 driver for go using database/sql"
