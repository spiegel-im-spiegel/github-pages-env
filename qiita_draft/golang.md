# はじめての Go 言語 (on Windows)

- [The Go Programming Language](https://golang.org/)
- [golang.jp - プログラミング言語Goの情報サイト](http://golang.jp/)

## まずはインストールと動作確認

Go言語はコンパイル言語である。プラットフォームとして FreeBSD, Lunux, Mac OS X, Windows などがある。コンパイラの実装としてはネイティブの gc コンパイラと gcc 版 go コンパイラ（gccgo）が存在する（当然ネイティブのほうが安定している、らしい）。 Windows 用には MinGW 版の gccgo がある（ただし MSYS 等の周辺ツールは必要ない）。

現時点（2015-04-09）での最新版は 1.4.2。 Windows 版では[ダウンロードページ](https://golang.org/dl/)にパッケージが用意されているのでダウンロードしてインストールすればよい。

Windows 版の場合、ルート直下の `C:\Go` フォルダとかとんでもないところにインストールしようとする。これが嫌なら、インストール時に任意のフォルダを指定すること。

インストールしたら環境変数 `GOROOT` にインストールフォルダ（例：`C:\Go`）をセットする。また環境変数 `PATH` に実行モジュールのあるフォルダ（例：`C:\Go\bin`）を追記しておく。（ただし、 `*.msi` のインストールパッケージでインストールする際はインストーラが環境変数を正しくセットしてくれる）

セットアップできたら動作確認する。

```shell
C:>go version
go version go1.4.2 windows/amd64
```

まずは、みんな大好き Hello World。

```go:hello.go
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```

このソースコード hello.go を compile and run。

```shell
C:>go run hello.go
hello, world
```

おおっ、動いた動いた。じゃあ、これを build して、実行モジュールを起動してみる。

```shell
C:>go build hello.go

C:>hello.exe
hello, world
```

よし。ちゃんと動くようだな。では、[次回](http://qiita.com/spiegel-im-spiegel/items/047a9bd6436e6391ddd4)へ続く。

## 目次

[目次はこちら](http://qiita.com/spiegel-im-spiegel/items/98d49ac456485b007a15#%E3%81%AF%E3%81%98%E3%82%81%E3%81%A6%E3%81%AE-go-%E8%A8%80%E8%AA%9E-on-windows)に移動。
