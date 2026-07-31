+++
title = "紙芝居用の簡易 Web サーバを Go 言語で書く【叱られ編】"
date =  "2022-10-17T22:28:28+09:00"
description = "またまた lint に叱られる"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "web", "security", "risk" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

このネタも3回目なので強引にシリーズ化（笑）

1. [紙芝居用の簡易 Web サーバを Go 言語で書く]({{< ref "/remark/2018/02/simple-web-server-with-golang.md" >}})
2. [紙芝居用の簡易サーバを書く【Go 1.16 版】]({{< ref "/golang/embeded-filesystem.md" >}})
3. [紙芝居用の簡易 Web サーバを Go 言語で書く【叱られ編】]({{< ref "/golang/simple-web-server-with-golang-3.md" >}}) ←イマココ

## net.JoinHostPort 関数を使え！

これまでの2回の記事を受けて，今回はこのコードからスタート。

```go {hl_lines=[20]}
package main

import (
    "embed"
    "flag"
    "fmt"
    "io/fs"
    "net/http"
    "os"
)

//go:embed html
var assets embed.FS

func main() {
    port := flag.Int("p", 3000, "port number")
    host := flag.String("host", "", "hostname")
    flag.Parse()

    addr := fmt.Sprintf("%s:%d", *host, *port)
    if len(*host) == 0 {
        fmt.Printf("Open http://localhost%s/\n", addr)
    } else {
        fmt.Printf("Open http://%s/\n", addr)
    }
    fmt.Println("Press ctrl+c to stop")

    root, err := fs.Sub(assets, "html")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }

    http.Handle("/", http.FileServer(http.FS(root)))
    if err := http.ListenAndServe(addr, nil); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
}
```

このうち，色付きの行に関連する以下の tweet を見かけた。

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">Goで、pathの結合はfilepath.Joinを使えというのは広まっている気がしているけど、同様にホストとポートを:で結合するのはnet.JoinHostPortを使ってほしい。fmt.Sprintfの場合、IPv6アドレスには:が含まれるので困ることになる。</p>&mdash; kadota (@plan9user) <a href="https://twitter.com/plan9user/status/1581870129423405056?ref_src=twsrc%5Etfw">October 17, 2022</a></blockquote>
{{< /fig-gen >}}

なるほど。
というわけで，先程の部分を

```go
addr := net.JoinHostPort(*host, strconv.Itoa(*port))
```

に置き換えた。
これを実行すると

```text
$ go run main.go -host "127.0.0.1" -p 8080
Open http://127.0.0.1:8080/
Press ctrl+c to stop
```

うんうん。
ちゃんと動くな。

## またまた lint に叱られる

今回のコードに対して念のため lint をかけてみる。

```text
$ golangci-lint run --enable gosec
main.go:35:12: G114: Use of net/http serve function that has no support for setting timeouts (gosec)
    if err := http.ListenAndServe(addr, nil); err != nil {
              ^
```

おぅふ。
そっちかよ `orz`

[`http`]`.ListenAndServe()` 関数の中身を見ると

```go
// ListenAndServe listens on the TCP network address addr and then calls
// Serve with handler to handle requests on incoming connections.
// Accepted connections are configured to enable TCP keep-alives.
//
// The handler is typically nil, in which case the DefaultServeMux is used.
//
// ListenAndServe always returns a non-nil error.
func ListenAndServe(addr string, handler Handler) error {
    server := &Server{Addr: addr, Handler: handler}
    return server.ListenAndServe()
}
```

となっている。
たしかにタイムアウト関連のフィールドがまるっと無視（つまりゼロ値が設定）されてるな。

試しに

```go
if err := http.ListenAndServe(addr, nil); err != nil {
    fmt.Fprintln(os.Stderr, err)
}
```

を以下に置き換えてみる。

```go
server := &http.Server{
    Addr:    addr,
    Handler: nil,
}
if err := server.ListenAndServe(); err != nil {
    fmt.Fprintln(os.Stderr, err)
}
```

これで機能は全く同じになる。
これを lint にかけてみる。

```text
$ golangci-lint run --enable gosec
main.go:35:13: G112: Potential Slowloris Attack because ReadHeaderTimeout is not configured in the http.Server (gosec)
    server := &http.Server{
        Addr:    addr,
        Handler: nil,
    }
```

おおっ。
内容が変わった。
ふむふむ。
[`http`]`.Server.ReadHeaderTimeout` フィールドに値を設定しろということか。

ちなみに `ReadHeaderTimeout` は，名前の通り，リクエストヘッダ読み込み時のタイムアウト時間を指定する [`time`]`.Duration` 型のフィールドで，ゼロ値がセットされているとタイムアウトが設定されないらしい。
つまり `ReadHeaderTimeout` フィールドは，悪意を持った巨大リクエストヘッダを読み込まされることで処理全体が stall しないようにするための措置のようだ。

というわけで書き直す。

```go {hl_lines=[4]}
server := &http.Server{
    Addr:              addr,
    Handler:           nil,
    ReadHeaderTimeout: 3 * time.Second,
}
if err := server.ListenAndServe(); err != nil {
    fmt.Fprintln(os.Stderr, err)
}
```

3秒という値には特に意味はない。
どのくらいが適当なのかねぇ。

とにかく，これで問題なく動作することを確認した後，3たび lint をかけてみる。

```go
$ golangci-lint run --enable gosec
```

なんとか lint は通ったようだ。

余談だが [`http`]`.Server.Handler` フィールドにゼロ値（＝`nil`）がセットされていると [`http`]`.DefaultServeMux` が既定のハンドラとして使われる。
また [`http`]`.Handle()` 関数の中身は

```go
// Handle registers the handler for the given pattern
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
func Handle(pattern string, handler Handler) { DefaultServeMux.Handle(pattern, handler) }
```

となっていて [`http`]`.DefaultServeMux` にハンドラを登録していることが分かる。
さらに [`http`]`.DefaultServeMux` は

```go
// DefaultServeMux is the default ServeMux used by Serve.
var DefaultServeMux = &defaultServeMux

var defaultServeMux ServeMux
```

と定義されている。
なので，先程のコードは

```go {hl_lines=[1,2,5]}
serverMux := http.NewServeMux()
serverMux.Handle("/", http.FileServer(http.FS(root)))
server := &http.Server{
    Addr:              addr,
    Handler:           serverMux,
    ReadHeaderTimeout: 3 * time.Second,
}
if err := server.ListenAndServe(); err != nil {
    fmt.Fprintln(os.Stderr, err)
}
```

と等価だ。
こちらのほうが却って分かりやすいかもしれない。

## Ctrl+C でサーバを Graceful にシャットダウンする

[`http`]`.Server` のドキュメントに `Ctrl+C` の SIGNAL を受信したら `Shutdown()` メソッドを走らせて graceful にシャットダウンするサンプルが載っていたので，それを参考に組み込んで今回の最終コードとしてみた。
全体としてはこんな感じでどうだろう。

```go {hl_lines=["48-68"]}
package main

import (
    "context"
    "embed"
    "errors"
    "flag"
    "fmt"
    "io/fs"
    "net"
    "net/http"
    "os"
    "os/signal"
    "strconv"
    "time"
)

//go:embed html
var assets embed.FS

func main() {
    port := flag.Int("p", 3000, "port number")
    host := flag.String("host", "", "hostname")
    flag.Parse()

    addr := net.JoinHostPort(*host, strconv.Itoa(*port))
    if len(*host) == 0 {
        fmt.Printf("Open http://localhost%s/\n", addr)
    } else {
        fmt.Printf("Open http://%s/\n", addr)
    }
    fmt.Println("Press ctrl+c to stop")

    root, err := fs.Sub(assets, "html")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }

    serverMux := http.NewServeMux()
    serverMux.Handle("/", http.FileServer(http.FS(root)))
    server := &http.Server{
        Addr:              addr,
        Handler:           serverMux,
        ReadHeaderTimeout: 3 * time.Second,
    }

    idleConnsClosed := make(chan struct{})
    go func() {
        defer close(idleConnsClosed)
        sigint := make(chan os.Signal, 1)
        signal.Notify(sigint, os.Interrupt)
        <-sigint

        if err := server.Shutdown(context.Background()); err != nil {
            fmt.Fprintln(os.Stderr, "shutdown error:", err)
            return
        }
        fmt.Println("\ngraceful shutdown")
    }()

    if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
        if err != nil {
            fmt.Fprintln(os.Stderr, "server error:", err)
        }
        return
    }
    <-idleConnsClosed
}
```

これを実行する。

```text
$ go run main.go -host "127.0.0.1" -p 8080
Open http://127.0.0.1:8080/
Press ctrl+c to stop
^C
graceful shutdown
```

`Ctrl+C` でちゃんとシャットダウンしてるかな。
よしよし。

## ブックマーク

- [Goのnet/httpのtimeoutについて - Carpe Diem](https://christina04.hatenablog.com/entry/go-timeouts)

[Go]: https://go.dev/
[`http`]: https://pkg.go.dev/net/http "http package - net/http - Go Packages"
[`time`]: https://pkg.go.dev/time "time package - time - Go Packages"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
