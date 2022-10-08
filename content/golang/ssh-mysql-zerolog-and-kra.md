+++
title = "SSH, MySQL, Zerolog, そして Kra"
date =  "2022-09-22T23:05:54+09:00"
description = "LOAD DATA INFILE 文を駆動させるところまで。"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "sql", "orm", "mysql" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

## SSH, MySQL, Zerolog

VPS 上に構築された MySQL サービスに大量のデータを送り込む必要がありまして。
[Go] でバッチ処理を組もうと考えたわけだ。
当然 MySQL サービスは VPS の外から直接アクセスできないので SSH トンネルをくぐる必要がある。

というわけで最初に作ったのが [`github.com/goark/sshql`][`sshql`] だった。

- [SSH 越しに DB サーバにアクセスする]({{< ref "/release/2022/09/sql-over-ssh.md" >}})

再掲載すると，こんな感じに書ける。

```go
package main

import (
    "database/sql"
    "fmt"
    "os"

    "github.com/goark/sshql"
    "github.com/goark/sshql/mysqldrv"
)

func main() {
    dialer := &sshql.Dialer{
        Hostname:   "sshserver",
        Port:       22,
        Username:   "remoteuser",
        Password:   "passphraseforauthkey",
        PrivateKey: "/home/username/.ssh/id_eddsa",
    }
    mysqldrv.New(dialer).RegisterDial("ssh+tcp")

    db, err := sql.Open("mysql", "dbuser:dbpassword@ssh+tcp(localhost:3306)/dbname")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    defer dialer.Close()
    defer db.Close()

    rows, err := db.Query("SELECT id, name FROM example ORDER BY id")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    for rows.Next() {
        var id int64
        var name string
        if err := rows.Scan(&id, &name); err != nil {
            fmt.Fprintln(os.Stderr, err)
            break
        }
        fmt.Printf("ID: %d  Name: %s\n", id, name)
    }
    rows.Close()
}
```

`ssh+tcp` という名前で dialer を登録し，この名前を使って DSN (Data Source Name) を構成するというのがポイント。
Dialer のクローズを忘れずに（笑）

ただ，これだとログが取れない。
んで，どうせログを取るなら [`github.com/rs/zerolog`][zerolog] パッケージを使いたいわけですよ。

そこで登場するのが [`github.com/simukti/sqldb-logger`][`sqldb-logger`] パッケージ。
これを使えば標準の [`sql`]`.DB` に [zerolog] などのサードパーティ製 logger を仕込むことができる。
こんな感じ。

```go {hl_lines=["10-12","25-26","31-36"]}
package main

import (
    "database/sql"
    "fmt"
    "os"

    "github.com/goark/sshql"
    "github.com/goark/sshql/mysqldrv"
    "github.com/rs/zerolog"
    sqldblogger "github.com/simukti/sqldb-logger"
    "github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

func main() {
    dialer := &sshql.Dialer{
        Hostname:   "sshserver",
        Port:       22,
        Username:   "remoteuser",
        Password:   "passphraseforauthkey",
        PrivateKey: "/home/username/.ssh/id_eddsa",
    }
    mysqldrv.New(dialer).RegisterDial("ssh+tcp")

    dsn := "dbuser:dbpassword@ssh+tcp(localhost:3306)/dbname"
    mysqlDb, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    db := sqldblogger.OpenDriver(
        dsn,
        mysqlDb.Driver(),
        zerologadapter.New(zerolog.New(os.Stderr)),
        sqldblogger.WithMinimumLevel(sqldblogger.LevelDebug),
    )
    defer dialer.Close()
    defer db.Close()

    rows, err := db.Query("SELECT id, name FROM example ORDER BY id")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    for rows.Next() {
        var id int64
        var name string
        if err := rows.Scan(&id, &name); err != nil {
            fmt.Fprintln(os.Stderr, err)
            break
        }
        fmt.Printf("ID: %d  Name: %s\n", id, name)
    }
    rows.Close()
}
```

DSN を2回指定しないといけないのが若干鬱陶しいが，ともかくこれで [zerolog] による構造化ログが出るようになった。

## さらに Kra を仕込む

[2022年春の Go Conference](https://gocon.connpass.com/event/212162/ "Go Conference 2022 Spring Online - connpass") でとても[感銘を受けた](https://gocon.jp/2022spring/ja/sessions/b7-l/ "Go Conference 2022 Spring | Go で RDB に SQL でアクセスするためのライブラリ Kra の紹介")のが [`github.com/taichi/kra`][`kra`] パッケージ。

- [DBアクセスライブラリ Kra](https://drive.google.com/file/d/1lsF7q74o0Akewk-5rUb9Jt9nA2MqpDDw/view)

いや [ent] とかってデータ構造とその関係を最初から構築するならいい道具だと思うけど（コードで設計できるのは素晴らしい！），既にある RDBMS 環境を使う場合には必ずしもマッチしないのよね。
更に言うと，既存 ORM やクエリビルダとかの中途半端な SQL 抽象化・隠蔽にはウンザリしてるのよ（この辺は [Go] に限らないけど）。
それなら最初からガチで SQL 文を書いて，クエリプランをチェックしつつ評価・最適化して，それから実装を進めるべきだと常々思っていた[^sql1]。

[^sql1]: SQL はひとつの独立した言語で（チューリング完全），宣言型プログラミング言語と考えるのが分かりやすい（異論は認めるw）。宣言型で分かりやすいのは正規表現だろう（関数型言語には宣言型が多い。 Lisp とか Haskell とか）。特に [Go] は宣言型の言語とはあまり相性がよくない。たとえば，正規表現の（構文解析やコンパイラではなく）ビルダを作ろうと考える人は少ないだろう（労力に見合わない）。どの言語・フレームワークでも同じことだが ORM やクエリビルダを使って頑張って抽象化や隠蔽をしても上手くマッチしない局面が多く，結局は「ガチの SQL でいいぢゃん。 PREPARE 構文で事前準備して変数部分はプレースホルダ経由で渡せば安全は確保される」となる。そういう意味じゃ [Go] 標準の [`database/sql`][`sql`] パッケージは，かなり妥当な割り切りをしてると思う。まぁ，後方互換性を保つためにちょっとアレな感じになっているのは否めないけど（笑）

というわけで [`github.com/taichi/kra`][`kra`] を使うチャンスを伺っていたのだが，今回はお試しにはちょうどいいサイズだったので採用した。
とはいえ，今回のような構成ではどうすればいいのか分からなくて [`kra`] パッケージのソースコードやサンプルコードを眺めながら，まずは以下のように書いてみる。

```go {hl_lines=["13-14", "33-41", 45, 50,"52-60"]}
package main

import (
    "database/sql"
    "fmt"
    "os"

    "github.com/goark/sshql"
    "github.com/goark/sshql/mysqldrv"
    "github.com/rs/zerolog"
    sqldblogger "github.com/simukti/sqldb-logger"
    "github.com/simukti/sqldb-logger/logadapter/zerologadapter"
    "github.com/taichi/kra"
    kraSql "github.com/taichi/kra/sql"
)

func main() {
    dialer := &sshql.Dialer{
        Hostname:   "sshserver",
        Port:       22,
        Username:   "remoteuser",
        Password:   "passphraseforauthkey",
        PrivateKey: "/home/username/.ssh/id_eddsa",
    }
    mysqldrv.New(dialer).RegisterDial("ssh+tcp")

    dsn := "dbuser:dbpassword@ssh+tcp(localhost:3306)/dbname"
    mysqlDb, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    db := kraSql.NewDB(
        sqldblogger.OpenDriver(
            dsn,
            mysqlDb.Driver(),
            zerologadapter.New(zerolog.New(os.Stderr)),
            sqldblogger.WithMinimumLevel(sqldblogger.LevelDebug),
        ),
        kraSql.NewCore(kra.NewCore(kra.MySQL)),
    )
    defer dialer.Close()
    defer db.Close()

    rs, err := db.Query(context.TODO(), "SELECT id, name FROM example ORDER BY id")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    rows := rs.Rows()
    for rows.Next() {
        rec := struct {
            Id   int64  `db:"id"`
            Name string `db:"name"`
        }{}
        if err := kraSql.NewRows(kraSql.NewCore(kra.NewCore(kra.MySQL)), rows).Scan(&rec); err != nil {
            fmt.Fprintln(os.Stderr, err)
            break
        }
        fmt.Printf("%#v\n", rec)
    }
    rows.Close()
}
```

クエリ結果の取り出しが若干まどろこしいが，これは [`kra`]`/sql.Rows` には何故か `Next()` メソッドがないため[^pgx1]。
しょうがないので 標準の [`sql`]`.Rows` を取り出して，そちらの `Next()` メソッドで回している。

[^pgx1]: [`github.com/jackc/pgx`][`pgx`] 専用の [`kra`]`/pgx.Rows` にはちゃんと `Next()` メソッドが付いている。提供されているメソッドが微妙に違う理由はよく分からないが，ソースコードを眺めるに，何となく意図的にそうなってる気がする。

これでちゃんと動いてログも取れているのを確認できた。

## トランザクション制御

トランザクション制御用に以下のような関数を用意する。

```go
func Transaction(ctx context.Context, db *kraSql.DB, opts *sql.TxOptions, fn func(tx *kraSql.Tx) error) error {
    tx, err := db.BeginTx(ctx, opts)
    if err != nil {
        return errs.Wrap(err)
    }
    defer func() {
        if v := recover(); v != nil {
            _ = tx.Rollback()
            panic(v)
        }
    }()

    if err := fn(tx); err != nil {
        if rErr := tx.Rollback(); rErr != nil {
            return errs.Wrap(rErr, errs.WithCause(err))
        }
        return errs.Wrap(err)
    }

    if err := tx.Commit(); err != nil {
        return errs.Wrap(err)
    }
    return nil
}
```

なお，エラーハンドリングには自作の [`github.com/goark/errs`](https://github.com/goark/errs "goark/errs: Error handling for Golang") パッケージを使っている。
[zerolog] と組み合わせて[エラーを構造化してログに吐ける](https://zenn.dev/spiegel/books/error-handling-in-golang/viewer/error-logging "ぼくがかんがえたさいきょうのえらーろぐ｜Go のエラーハンドリング")のが利点。

実際にトランザクション処理を行う場合は，たとえば

```go
// logger := zerolog.New(os.Stderr)
// ctx := context.TODO()
values := struct {
    Id   int64  `db:"id"`
    Name string `db:"name"`
}{
    Id:   100,
    Name: "Alice",
}
if err := Transaction(ctx, db, &sql.TxOptions{}, func(tx *kraSql.Tx) error {
    stmt, err := tx.Prepare(ctx, "INSERT INTO example(id,name) VALUES (:id,:name)")
    if err != nil {
        return errs.Wrap(err)
    }
    defer stmt.Close()

    res, err := stmt.Exec(ctx, &values)
    if err != nil {
        return errs.Wrap(err)
    }
    count, err := res.RowsAffected()
    if err != nil {
        return errs.Wrap(err)
    }
    logger.Info().Int64("affected", count).Send()
    return nil
}); err != nil {
    logger.Error().Interface("error", err).Send()
    ...
}
```

てな感じに書ける。
こういうのが [`kra`] にあるとめがっさ便利なんだけどねぇ（それを言い出すとパッケージがどんどん膨れてしまうのだがw）。

## LOAD DATA INFILE 文で大量のデータを突っ込む

さて，いよいよ MySQL の テーブルに大量のデータを突っ込むのだが， `INSERT` 文でちまちまやってたら日が暮れてしまうので（実際に試して日が暮れた） `LOAD DATA INFILE` 文を使うことにする。

- [MySQL :: MySQL 5.6 リファレンスマニュアル :: 13.2.6 LOAD DATA INFILE 構文](https://dev.mysql.com/doc/refman/8.0/ja/load-data.html)

こんな感じの SQL 文。

```sql
LOAD DATA LOCAL
  INFILE 'input.file'
  INTO TABLE exsample_table
  CHARACTER SET utf8mb4
  FIELDS TERMINATED BY '\t'
  LINES TERMINATED BY '\n'
  (
    field1,
    field2,
    ...
  )
```

これでローカルにある `input.file` ファイルの内容をリモートの MySQL の `exsample_table` テーブルに送り込める。
`(field1, field2, ...)` の並びと入力ファイルの要素の並びが同じであることが前提。
またサーバ側の MySQL サービスが `--local_infile` オプション付きで起動されていること[^local1]。

[^local1]: `local_infile` オプションをセットすると mysqldump を使用する際に「そんなフラグは知らん」と怒られる場合がある。これを回避するには `loose-local-infile` にするといいらしい。当然ではあるが `LOAD DATA LOCAL ...` 文は[セキュリティ・リスクがある](https://docs.oracle.com/cd/E17952_01/mysql-8.0-ja/load-data-local-security.html "6.1.6 LOAD DATA LOCAL のセキュリティー上の考慮事項")ので取り扱い注意である。

`CHARACTER SET` 句はファイルの文字エンコードディングがDBサービスのデフォルトと異なる場合に設定する。
`FIELDS` 句および `LINES` を省略した場合のデフォルト値はこうなっているそうな。

```sql
FIELDS TERMINATED BY '\t' ENCLOSED BY '' ESCAPED BY '\\'
LINES TERMINATED BY '\n' STARTING BY ''
```

いわゆる TSV (Tab Separated Value) 形式のレコードだね。
これ以外の形式なら明示的に設定する必要がある。

で，これを [`github.com/go-sql-driver/mysql`][`mysql`] パッケージで実装するには，3つのステップが必要。

ひとつ目はデータ読み込みハンドラを登録する。
こんな感じ[^ld1]。

[^ld1]: 実は [`mysql`]`.RegisterLocalFile()` 関数を使えば直接ファイルパスを登録することができる。ハンドラ登録で [`io`]`.Reader` interface 型で渡すほうが応用が効きやすいので，今回はこちら。

```go
file, err := os.Open("input.file")
if err != nil {
    return err
}
defer file.Close()

mysql.RegisterReaderHandler("data", func() io.Reader {
    return file
})
```

次に実際の SQL 文を発行する。

```go
// logger := zerolog.New(os.Stderr)
// ctx := context.TODO()
if err := Transaction(ctx, db, &sql.TxOptions{}, func(tx *kraSql.Tx) error {
    res, err := tx.Exec(ctx, `LOAD DATA LOCAL INFILE 'Reader::data' INTO TABLE exsample_table (field1, field2, ...)`)
    if err != nil {
        return errs.Wrap(err)
    }
    count, err := res.RowsAffected()
    if err != nil {
        return errs.Wrap(err)
    }
    logger.Info().Int64("affected", count).Send()
    return nil
}); err != nil {
    logger.Error().Interface("error", err).Send()
    ...
}
```

ファイルを指定する部分に先ほど登録したハンドラの名前を使って `'Reader::data'` と指定する。
爆速でした。

ちなみに，うっかり `tx.Prepare()` で前準備しようとすると「そんな構文はサポートしてない」（←超意訳）と怒られる。
`PREPARE` で対応している構文は以下のページが参考になる。

- [MySQL :: MySQL 8.0 リファレンスマニュアル :: 13.5 プリペアドステートメント](https://dev.mysql.com/doc/refman/8.0/ja/sql-prepared-statements.html)

最後に [`mysql`]`.DeregisterReaderHandler()` 関数で登録を解除する。
後始末はきちんとね。

## ブックマーク

- [【自分用のメモ】MySQL8でインポート・エクスポート - Qiita](https://qiita.com/yes_dog/items/57c1a8c03aff8bb177de)

[Go]: https://go.dev/
[ent]: https://entgo.io/
[`sql`]: https://pkg.go.dev/database/sql "sql package - database/sql - Go Packages"
[`io`]: https://pkg.go.dev/io "io package - io - Go Packages"
[`mysql`]: https://github.com/go-sql-driver/mysql "go-sql-driver/mysql: Go MySQL Driver is a MySQL driver for Go's (golang) database/sql package"
[`pgx`]: https://github.com/jackc/pgx "jackc/pgx: PostgreSQL driver and toolkit for Go"
[`sshql`]: https://github.com/goark/sshql "goark/sshql: Go SQL drivers over SSH"
[zerolog]: https://github.com/rs/zerolog "rs/zerolog: Zero Allocation JSON Logger"
[`sqldb-logger`]: https://github.com/simukti/sqldb-logger "simukti/sqldb-logger: A logger for Go SQL database driver without modifying existing *sql.DB stdlib usage."
[`kra`]: https://github.com/taichi/kra "taichi/kra: relational database access helper library"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B09C2XBC2F" %}} <!-- Golang Tシャツ -->
