+++
title = "SQLite 3 データを A5:SQL Mk-2 で開く"
date = "2018-04-25T13:54:05+09:00"
description = "ネイティブで SQLite に対応してるのか，これ。"
image = "/images/attention/kitten.jpg"
tags = [ "sqlite", "tools" ]

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

最近，ン年ぶりに [PostgreSQL] 周りの仕事をやってるんだけど [pgAdmin] がアレでねぇ。
普段使いならテーブルやビューやストアド・プロシージャが見ればよく，ついでに SQL 文のチューニングが手軽に試せるならラッキー！ くらいの機能でいいのよ。
なので [A5:SQL Mk-2] に乗り換えたのさ。

で，よく見たらネイティブで [SQLite] に対応してるのか，これ。

{{< fig-img src="https://farm1.staticflickr.com/967/26816747937_2b4195bd3a_o.png" title="A5:SQL Mk-2" link="https://www.flickr.com/photos/spiegel/26816747937/" >}}

追加もDBファイルを指定するだけなので簡単（sqlite3.dll が別途必要）。

{{< fig-img src="https://farm1.staticflickr.com/979/27815800168_0fa56d9c07.jpg" title="SQLite 3 w/ A5:SQL Mk-2" link="https://www.flickr.com/photos/spiegel/27815800168/" >}}

ただしDBファイルを暗号化している場合は [SQLite ODBC Driver](http://www.ch-werner.de/sqliteodbc/) を経由する必要があるようだ。

- [System.Data.SQLiteで暗号化したSQLiteファイルをA5:SQL Mk-２で | ひまブログ](https://ameblo.jp/hirokun-marichan/entry-12168092949.html)
- [暗号化されたSQLiteデータベースファイルに「A5:SQL Mk-2」で接続する | ドラブロ](https://www.doraxdora.com/blog/2017/10/27/post-2888/)

私は今のところ必要ないけどね。

## ブックマーク

- [Go 言語で SQLite を使う（Windows 向けの紹介）]({{< relref "golang/sqlite-with-golang-in-windows.md" >}})
- [xeodou/go-sqlcipher: Golang SQLCipher driver conforming to the built-in database/sql interface and using the latest sqlite3 code.](https://github.com/xeodou/go-sqlcipher)

[SQLite]: https://www.sqlite.org/
[PostgreSQL]: https://www.postgresql.org/ "PostgreSQL: The world's most advanced open source database"
[pgAdmin]: https://www.pgadmin.org/ "pgAdmin - PostgreSQL Tools"
[A5:SQL Mk-2]: https://a5m2.mmatsubara.com/ "A5:SQL Mk-2 - フリーの汎用SQL開発ツール/ER図ツール .. 松原正和"
