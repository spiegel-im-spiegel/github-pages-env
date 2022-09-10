+++
title = "SSH 越しに DB サーバにアクセスする"
date =  "2022-09-10T18:04:56+09:00"
description = "mattn さんが公開されているパッケージを参考に組んでみた"
image = "/images/attention/go-logo_blue.png"
tags  = [ "golang", "programming", "package", "module", "sql", "openssh" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

PostgreSQL や MySQL などの RDBMS サービスにアクセスするために [Go] では標準で [`database/sql`][`sql`] パッケージを用意している
（実際にサービスにアクセスするためには [`github.com/lib/pq`] や [`github.com/go-sql-driver/mysql`] といったドライバ・パッケージを使う必要がある）。
たとえばこんな感じ。

```go
db, err := sql.Open("postgres", "postgres://dbuser:dbpassword@dbserver:5432/example?sslmode=require")
```

ただし，これはクライアントからサービスに直結する場合で，たとえば VPS から SSH 経由でサービスにアクセスする必要がある場合は少し工夫が必要である。
ありがたいことに PostgreSQL を SSH 経由でアクセスするためのパッケージを [mattn](https://github.com/mattn) さんが公開して下さっている。

- [mattn/pqssh](https://github.com/mattn/pqssh)
- [Go で SSH 超しに PostgreSQL に接続できる database/sql ドライバを作った。](https://zenn.dev/mattn/articles/d1b114e2d4a421)

ありがたや {{< emoji "ペコン" >}}

で，実は MySQL サービスを SSH 経由でアクセスする必要ができたので，上のパッケージを参考に自作してみた。

- [goark/sshql: Go SQL drivers over SSH](https://github.com/goark/sshql)

このパッケージを使ってこんな感じに書ける。

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
    mysqldrv.New(&sshql.Dialer{
        Hostname:   "sshserver",
        Port:       22,
        Username:   "remoteuser",
        Password:   "passphraseforauthkey",
        PrivateKey: "/home/username/.ssh/id_eddsa",
    }).RegisterDial()

    db, err := sql.Open("mysql", fmt.Sprintf("dbuser:dbpassword@%s(localhost:3306)/dbname", mysqldrv.DialName))
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
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

MySQL の場合 SSH でアクセスするための Dialer を登録して，登録文字列を DSN に含める必要がある。

なお，このパッケージを使った PostgreSQL への SSH 越しのアクセスはこんな感じに書ける。

```go
package main

import (
    "database/sql"
    "fmt"
    "os"

    "github.com/goark/sshql"
    "github.com/goark/sshql/pgdrv"
)

func main() {
    pgdrv.New(&sshql.Dialer{
        Hostname:   "sshserver",
        Port:       22,
        Username:   "remoteuser",
        Password:   "passphraseforauthkey",
        PrivateKey: "/home/username/.ssh/id_eddsa",
    }).Register()

    db, err := sql.Open(pgdrv.DriverName, "postgres://dbuser:dbpassword@localhost:5432/example?sslmode=disable")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
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

やっぱ [`sql`]`.Open()` 関数の第1引数で専用ドライバを指定するほうがシンプルだよなぁ。
DSN 文字列をいじらなくて済むし。

ここでごめんペコン {{< emoji "ペコン" >}}

SSH 越しにアクセスできる適当な PostgreSQL 環境を持ってなくて，上のコードはテストしてません。
「あんだ，動かねーよ！」とかありましたら pull request いただけるとありがたいです。

## pgx ドライバでも SSH 越しにアクセスしたいのだが...

PostgreSQL にアクセスするなら [`github.com/jackc/pgx`] パッケージを使いたいところだが [`github.com/lib/pq`] のように専用の Dialer を登録できるインタフェースはないっぽい？ 「◯◯すればできるよ」という情報がありましたら是非教えてください。
まぁ，試す環境がないんだけどね。

## InsecureIgnoreHostKey() 関数で叱られる

mattn さんの [`github.com/mattn/pqssh`] パッケージの中で

```go
sshConfig := &ssh.ClientConfig{
    User:            d.Username,
    Auth:            []ssh.AuthMethod{},
    HostKeyCallback: ssh.InsecureIgnoreHostKey(),
}
```

という記述があり，最初はそのまま真似してたのだが，例によって lint で「あかんがな！」と叱られた。

`HostKeyCallback` 項目は SSH ログイン時のホスト認証の動作を定義しているもので， [`ssh`]`.InsecureIgnoreHostKey()` は何もしないで `nil` を返す関数を返している。

```go
// InsecureIgnoreHostKey returns a function that can be used for
// ClientConfig.HostKeyCallback to accept any host key. It should
// not be used for production code.
func InsecureIgnoreHostKey() HostKeyCallback {
    return func(hostname string, remote net.Addr, key PublicKey) error {
        return nil
    }
}
```

こりゃあ，確かにあかんわ（笑）

今回の [`github.com/goark/sshql`] パッケージでは一応ホスト認証を行っているが `~/.ssh/known_hosts` ファイルに登録されていないホストや登録されている鍵が異なる場合は問答無用でエラーを返すようにしている。
まぁ，こういうパッケージはバッチ処理とかにしか使わないだろうし，ええじゃろう。

なお [`sshql`][`github.com/goark/sshql`]`.Dialer` 構造体は

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

と定義しているけど `IgnoreHostKey` 要素に `true` をセットするとホスト認証をすっ飛ばしてくれる。

## ブックマーク

- [Using MySQL / MariaDB via SSH in Golang · GitHub](https://gist.github.com/vinzenz/d8e6834d9e25bbd422c14326f357cce0)
- [Golang – How to write ssh.HostKeyCallback – cyruslab](https://cyruslab.net/2020/10/23/golang-how-to-write-ssh-hostkeycallback/)
- [simukti/sqldb-logger: A logger for Go SQL database driver without modifying existing *sql.DB stdlib usage.](https://github.com/simukti/sqldb-logger)

[Go]: https://go.dev/
[`github.com/goark/sshql`]: https://github.com/goark/sshql "goark/sshql: Go SQL drivers over SSH"
[`sql`]: https://pkg.go.dev/database/sql "sql package - database/sql - Go Packages"
[`github.com/lib/pq`]: https://github.com/lib/pq "lib/pq: Pure Go Postgres driver for database/sql"
[`github.com/go-sql-driver/mysql`]: https://github.com/go-sql-driver/mysql "go-sql-driver/mysql: Go MySQL Driver is a MySQL driver for Go's (golang) database/sql package"
[`github.com/jackc/pgx`]: https://github.com/jackc/pgx "jackc/pgx: PostgreSQL driver and toolkit for Go"
[`github.com/mattn/pqssh`]: https://github.com/mattn/pqssh
[`ssh`]: https://pkg.go.dev/golang.org/x/crypto/ssh "ssh package - golang.org/x/crypto/ssh - Go Packages"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4873119979" %}} <!-- Go言語による分散サービス -->
{{% review-paapi "B09C2XBC2F" %}} <!-- Golang Tシャツ -->
