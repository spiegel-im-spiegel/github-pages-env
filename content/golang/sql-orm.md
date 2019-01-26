+++
title = "Go 言語で使える ORM と SQL ビルダー"
date = "2018-05-13T19:43:28+09:00"
description = "私は SQL 文を手書きで書くのは苦にならない性質なのだが，途中までコードを書いて流石に煩わしくなってきたので Go 言語製の ORM および SQL Builder パッケージを探すことにした。"
image = "/images/attention/go-logo_blue.png"
tags        = [ "golang", "programming", "sqlite", "sql", "orm" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = true
  mermaidjs = false
+++

先月あたりから余暇にコツコツ作ってた脆弱性情報の収集・管理ツール [jvnman] の最初のリリースを行った。

- [spiegel-im-spiegel/jvnman: JVN Vulnerability Data Management](https://github.com/spiegel-im-spiegel/jvnman)
    - [Release v0.1.0 · spiegel-im-spiegel/jvnman](https://github.com/spiegel-im-spiegel/jvnman/releases/tag/v0.1.0)
- [脆弱性情報の収集・管理ツール jvnman の最初のリリース]({{< ref "/release/2018/05/jvnman-0_1_0-is-released.md" >}})

[jvnman] についての説明は上の記事を見ていただくとして，この記事ではツールの内部で行っている SQL 文のハンドリングについて。

[jvnman] は収集した脆弱性情報を [SQLite] データベース・ファイルに格納している。
[jvnman] には簡単な帳票出力機能も付いているが，メインは脆弱性情報の収集・蓄積である。
一度 [SQLite] データベース・ファイルを作っておけば作成したデータベースのハンドリングは他のツール（たとえば Office ツール）でもできる。

（ゆえに最初は軽く考えていて [spiegel-im-spiegel/go-myjvn] の{{< ruby "おまけ" >}}サンプル{{< /ruby >}}として [jvnman] を組み込む予定だった。思ったより依存パッケージが多いことが分かってリポジトリを別にしたんだけど）

私は SQL 文を手書きで書くのは苦にならない性質なのだが（そもそも最初は [A5:SQL Mk-2 みたいなツール]({{< ref "/remark/2018/04/sqlite3.md" >}} "SQLite 3 データを A5:SQL Mk-2 で開く")で試して最適化を行うものだし），途中までコードを書いて流石に煩わしくなってきたので [Go 言語]製の ORM (Object-Relational Mapping) および SQL Builder パッケージを探すことにした。

というわけで，以下のパッケージを紹介。

- [go-gorp/gorp: Go Relational Persistence - an ORM-ish library for Go](https://github.com/go-gorp/gorp)
    - [Big Sky :: Go言語向けの ORM、gorp がなかなか良い](https://mattn.kaoriya.net/software/lang/go/20120914222828.htm)
    - [gorp(go-mysql-driver)で独自に定義した型をカラムに割り当てる - Qiita](https://qiita.com/itoudium/items/e599daa93ff24a15f5f6)
- [Masterminds/squirrel: Fluent SQL generation for golang](https://github.com/Masterminds/squirrel)

[go-gorp/gorp] では [Go 言語] の構造体（struct）を SQL のクエリ出力に関連付けることができる。
たとえば

```go
type Vulnlist struct {
	ID          sql.NullString `db:"id,primarykey"`
	Title       sql.NullString `db:"title"`
	Description sql.NullString `db:"description"`
	URI         sql.NullString `db:"uri"`
	Creator     sql.NullString `db:"creator"`
	Impact      sql.NullString `db:"impact"`
	Solution    sql.NullString `db:"solution"`
	DatePublic  sql.NullInt64  `db:"date_public"`
	DatePublish sql.NullInt64  `db:"date_publish"`
	DateUpdate  sql.NullInt64  `db:"date_update"`
}
```

のように struct タグを使って関連付けるわけだ。
これで `gorp.DbMap` インスタンス生成時に

```go
db, err := sql.Open("sqlite3", "./jvndb.sqlite3")
if err != nil {
    return err
}
dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
dbmap.AddTableWithName(Vulnlist{}, "vulnlist")
```

としておけば

```go
obj, err := dbmap.Get(Vulnlist{}, "JVNDB-2018-003082")
if err != nil {
    return err
}
if obj == nil {
    fmt.Println("mp data")
} else if ds, ok := obj.(*Vulnlist); ok {
    if ds.Title.Valid {
        fmt.Println("Title =", ds.Title.String)
    } else {
        fmt.Println("Title is NULL")
    }
}
```

のように書くことができる。
INSERT や UPDATE についても同じようにできる。

[Masterminds/squirrel] は簡易 SQL Builder で，あまりガチガチの抽象化をしないのが気に入っている。
たとえば

```go
var ds struct {
    dateUpdate int64 `db:"date_update"`
}
if psql, args, err := squirrel.Select("date_update").From("vulnlist").Where(squirrel.Eq{"id": "JVNDB-2018-003082"}).ToSql(); err != nil {
    return err
} else if err := dbmap.SelectOne(&ds, psql, args...); err != nil {
    return err
} else {
    fmt.Println(time.Unix(ds.dateUpdate, 0))
}
```

のように prepared statement とパラメータをちゃんと分離してくれる。

注意しないといけないのは `From()` メソッドや `Where()` メソッドなどで連結するたびにインスタンスのコピーが発生している点だろう[^bld1]。
[jvnman] のようなツールなら気にするまでもないが，短時間に大量のトランザクションが発生する可能性がある場合は注意したほうがいいかもしれない。

[^bld1]: Value object として見るなら正しい動きなんだけどね。 Builder ツールとしてその動きはアリなんだろうか，とつい考えてしまう。

## ブックマーク

- [GitHub - govert/SQLiteForExcel: A lightweight wrapper to give access to the SQLite3 library from VBA.](https://github.com/govert/SQLiteForExcel)
    - [ExcelからSQLiteを使う方法 | Gabekore Garage](https://gabekore.org/sqlite-for-excel-vba)
    - [SQLite for Excel | crossframe](http://crossframe.iiv.jp/201603051181/)
- [Datatypes In SQLite Version 3](http://www.sqlite.org/datatype3.html)
- [Base: how to connect to an SQLite database? [closed] - Ask LibreOffice](https://ask.libreoffice.org/en/question/139052/base-how-to-connect-to-an-sqlite-database/)
    - [LibreOffice Base をSQLite のフロントエンドにしてみよう](https://www.slideshare.net/78tch/libreoffice-base-sqlite)

- [Go 言語で SQLite を使う（Windows 向けの紹介）]({{< relref "localhost:1313/golang/sqlite-with-golang-in-windows.md" >}})
- [spiegel-im-spiegel/go-myjvn: Handling MyJVN RESTful API by Golang](https://github.com/spiegel-im-spiegel/go-myjvn)
    - [go-myjvn パッケージを作ってみた]({{< ref "/release/2018/03/go-myjvn-v0_1_0-released.md" >}})
    - [MyJVN API に関する覚え書き]({{< ref "/remark/2018/03/myjvn-api.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[jvnman]: https://github.com/spiegel-im-spiegel/jvnman "spiegel-im-spiegel/jvnman: JVN Vulnerability Data Management"
[SQLite]: https://www.sqlite.org/
[spiegel-im-spiegel/go-myjvn]: https://github.com/spiegel-im-spiegel/go-myjvn "spiegel-im-spiegel/go-myjvn: Handling MyJVN RESTful API by Golang"
[go-gorp/gorp]: https://github.com/go-gorp/gorp "go-gorp/gorp: Go Relational Persistence - an ORM-ish library for Go"
[Masterminds/squirrel]: https://github.com/Masterminds/squirrel "Masterminds/squirrel: Fluent SQL generation for golang"

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
