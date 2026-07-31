+++
title = "sshql — SSH 越しに RDBMS にアクセスする"
date =  "2022-09-30T21:17:13+09:00"
description = "拙作 github.com/goark/sshql は SSH 経由で MySQL や PostgreSQL といった RDBMS に接続するための Go パッケージである。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "package", "module", "sql", "openssh", "mysql", "postgresql" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

拙作 [`github.com/goark/sshql`][`sshql`] は SSH 経由で MySQL や PostgreSQL といった RDBMS に接続するための [Go] パッケージである。

- [GitHub - goark/sshql: Go SQL drivers over SSH](https://github.com/goark/sshql)

[`github.com/goark/sshql`][`sshql`] パッケージでは以下の Dialer 型の構造体を用意し，これを使って SSH に接続する[^mattn1]。

[^mattn1]: 今回のパッケージは [`github.com/mattn/pqssh`](https://github.com/mattn/pqssh) の事実上の fork である。同パッケージに敬意を表して拙作の方も MIT ライセンスで提供している。感謝！

```go
// Dialer is authentication provider information.
type Dialer struct {
    Hostname      string `json:"hostname"`
    Port          int    `json:"port"`
    Username      string `json:"username"`
    Password      string `json:"password"`
    PrivateKey    string `json:"privateKey"`
    IgnoreHostKey bool   `json:"IgnoreHostKey"`
    client        *ssh.Client
}
```

パスワード認証の場合は `Username` と `Password` にそれぞれ値をセットする。
`PrivateKey` に認証用の秘密鍵へのパスを指定する場合は `Password` に秘密鍵のパスフレーズをセットする。
また `ssh-agent` 等を使って `$SSH_AUTH_SOCK` ソケットから秘密鍵を取得できる場合は，こちらを優先するようになっている。

`IgnoreHostKey` に `true` をセットするとホスト認証を無視するようになっているが，セキュリティ上お勧めできないのでご注意を。
ホスト認証は `$HOME/.ssh/known_hosts` ファイルにホスト鍵が登録されていることが前提で， `known_hosts` ファイルがない，または `known_hosts` にホスト鍵がない場合は

```text
$ go run main.go
ssh: handshake failed: knownhosts: key is unknown
```

という感じにエラーになる。

では，これを使って実際にコードを書いてみる。

## [github.com/lib/pq][`pq`] + [github.com/goark/sshql][`sshql`]

まずは PostgreSQL 用ドライバ [`github.com/lib/pq`][`pq`] と [`github.com/goark/sshql`][`sshql`] を組み合わせるところから。
こんな感じに書ける。

```go {hl_lines=["13-28"]}
package main

import (
    "database/sql"
    "fmt"
    "os"

    "github.com/goark/sshql"
    "github.com/goark/sshql/pgdrv"
)

func main() {
    dialer := &sshql.Dialer{
        Hostname:   "sshserver",
        Port:       22,
        Username:   "remoteuser",
        Password:   "passphraseforauthkey",
        PrivateKey: "/home/username/.ssh/id_eddsa",
    }
    pgdrv.New(dialer).Register("postgres+ssh")

    db, err := sql.Open("postgres+ssh", "postgres://dbuser:dbpassword@localhost:5432/example?sslmode=disable")
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

`"postgres+ssh"` としてドライバを登録し（名前は何でもOK），この名前を [`sql`]`.Open()` 関数で指定しているのがポイントである。
あとは [`database/sql`][`sql`] の機能をそのまま使うことができる。

[`github.com/goark/sshql`][`sshql`]`/pgdrv` パッケージ内部で [`github.com/lib/pq`][`pq`] パッケージを呼び出しているため

```go
import _ "github.com/lib/pq"
```

のようなインポートは不要である。

## [github.com/jackc/pgx][`pgx`] + [github.com/goark/sshql][`sshql`]

[`github.com/jackc/pgx`][`pgx`] は [`github.com/lib/pq`][`pq`] の後継とも言えるパッケージである。
標準 [`database/sql`][`sql`] パッケージと組み合わせて使うことも可能だが，自身が [`database/sql`][`sql`] 互換のインタフェースを持っていて更に独自の機能も搭載している。
PostgreSQL にアクセスするなら，個人的にはこちらのほうがお勧めである。

拙作 [`github.com/goark/sshql`][`sshql`] と組み合わせて使うにはちょっと変則的な初期化コードを書く。
こんな感じ。

```go {hl_lines=["13-39"]}
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/goark/sshql"
    "github.com/jackc/pgx/v4/pgxpool"
)

func main() {
    dialer := &sshql.Dialer{
        Hostname:   "sshserver",
        Port:       22,
        Username:   "remoteuser",
        Password:   "passphraseforauthkey",
        PrivateKey: "/home/username/.ssh/id_eddsa",
    }

    cfg, err := pgxpool.ParseConfig("postgres://dbuser:dbpassword@localhost:5432/example?sslmode=disable")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    cfg.ConnConfig.DialFunc = dialer.DialContext

    if err := dialer.Connect(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    defer dialer.Close()

    pool, err := pgxpool.ConnectConfig(context.TODO(), cfg)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    defer pool.Close()

    rows, err := pool.Query(context.TODO(), "SELECT id, name FROM example ORDER BY id")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    rows.Close()

    for rows.Next() {
        var id int64
        var name string
        if err := rows.Scan(&id, &name); err != nil {
            fmt.Fprintln(os.Stderr, err)
            break
        }
        fmt.Printf("ID: %d  Name: %s\n", id, name)
    }
}
```

まず [`pgx`]`.ConnConfig` を含む [`pgx`]`/pgxpool.Config` 構造体を作って，そこに `dialer.DialContext()` メソッドを登録するイメージ。
[`pgx`] ドライバーをオープンする前に `dialer.Connect()` つまり SSH 接続を明示的に記述する必要がある。

一見まどろこしく見えるが， [`pgx`]`/pgxpool.Config` 構造体を使って logger 登録を含むカスタマイズができるので，実用上はそれほど迂遠なコードではなかったりする。
接続ごとにインスタンスを作れるので，グローバルに名前をつけてドライバーを登録するより合理的かもしれない。

## [github.com/go-sql-driver/mysql][`mysql`] + [github.com/goark/sshql][`sshql`]

MySQL 用ドライバ [`github.com/go-sql-driver/mysql`][`mysql`] と [`github.com/goark/sshql`][`sshql`] との組み合わせはこんな感じ。

```go {hl_lines=["13-28"]}
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

[`github.com/lib/pq`][`pq`] のときとは違いドライバー名ではなくプロトコル名（通常は `"tcp"` など）として登録し DSN に登録したプロトコル名（ここでは `ssh+tcp`）を含める。
あとは [`database/sql`][`sql`] の機能をそのまま使うことができる。

[`github.com/goark/sshql`][`sshql`]`/mysqldrv` パッケージ内部で [`github.com/go-sql-driver/mysql`][`mysql`] パッケージを呼び出しているため

```go
import _ "github.com/go-sql-driver/mysql"
```

のようなインポートは不要である。

[Go]: https://go.dev/
[`sshql`]: https://github.com/goark/sshql "goark/sshql: Go SQL drivers over SSH"
[`sql`]: https://pkg.go.dev/database/sql "sql package - database/sql - Go Packages"
[`pq`]: https://github.com/lib/pq "lib/pq: Pure Go Postgres driver for database/sql"
[`pgx`]: https://github.com/jackc/pgx "jackc/pgx: PostgreSQL driver and toolkit for Go"
[`mysql`]: https://github.com/go-sql-driver/mysql "go-sql-driver/mysql: Go MySQL Driver is a MySQL driver for Go's (golang) database/sql package"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
