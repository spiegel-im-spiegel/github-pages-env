+++
title = "パッケージの管理（モジュール対応版）"
date =  "2019-12-20T18:01:26+09:00"
description = "モジュール対応モードでは $GOPATH/src 下にある外部パッケージのコードは参照しない。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "package", "module", "engineering", "design" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

随分前に「[go get コマンドでパッケージを管理する]({{< relref "./go-get-package.md" >}})」を書いたのだが，内容が古すぎて使い物にならなくなっている。
この記事を全面改訂してもいいのだが，個人的には当時の試行錯誤っぷりに懐かしい気分になったので，これはそのまま手を加えず新たに記事を起こすことにした。
特にパッケージとモジュールの関係に注意して読んでいただければ幸いである。

{{< div-box type="markdown" >}}
**【2021-02-25 追記】**
なお，このブログで書き散らした内容をまとめる形で 以下の Zenn 記事を書いた。
こちらも併せてどうぞ。

- [Go のモジュール管理【バージョン 1.16 改訂版】](https://zenn.dev/spiegel/articles/20210223-go-module-aware-mode)
{{< /div-box >}}

## 前提条件

作業プラットフォームは [Ubuntu] で [Go] コンパイラのバージョンは 1.13.x (またはそれ以上) とする。
環境構築手順は以下を参考にどうぞ。

- [中古 PC に Ubuntu 環境を導入する]({{< ref "/remark/2019/12/install-ubuntu-to-second-hand-pc.md" >}})

各環境変数の値（一部）は以下の通り。
 
```text
$ go env GOPATH
GO111MODULE=""
GOARCH="amd64"
GOCACHE="/home/username/.cache/go-build"
GOENV="/home/username/.config/go/env"
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH="/home/username/go"
...
```

プラットフォームによって環境変数の値は異なるが凡その振る舞いは同じなので，適当に読み替えていただきたい。

ちなみに `GO111MODULE` の値として以下を設定できる。

| 値     | 内容                                                                                                     |
| ------ | -------------------------------------------------------------------------------------------------------- |
| `on`   | 常にモジュール対応モードで動作する                                                                       |
| `off`  | 常に GOPATH モードで動作する                                                                             |
| `auto` | `$GOPATH` 以下のディレクトリにあるパッケージは GOPATH モードで，それ以外はモジュール対応モードで動作する |

`GO111MODULE` 未定義時の既定は `auto`。
なお，バージョン 1.13 からは `auto` で `$GOPATH` 以下のディレクトリで作業していても `go.mod` によるモジュール定義がされている場合はモジュール対応モードで動作するようになった。

## ありのまま今おこった事を話すぜ！

### パッケージ hello を作って実行してみる

まずは `$GOPATH/src/hello` ディレクトリを作成し，以下のソースコード `hello.go` を記述する。

```go
package hello

import "fmt"

func Hello(s string) {
	if s == "" {
		s = "there"
	}
	fmt.Println("Hello", s)
}
```

ついでに動作確認を兼ねてテスト用に `example_test.go` ファイルも作っておこう。
中身はこんな感じ。

```go
package hello_test

import "hello"

func ExampleHello() {
	hello.Hello("World")
	//Output:
	//Hello World
}
```

これで

```text
$ go test ./...
ok  	hello	0.001s
```

とかなればパッケージ `hello` の完成である。

次にパッケージ `hello` を使うコードを書いてみる。
`$GOPATH/src/sample` ディレクトリを作成し，以下のソースコード `sample.go` を記述する。

```go
package main

import "hello"

func main() {
	hello.Hello("World")
}
```

先程のサンプル関数 `ExampleHello()` と同じ内容なので出力結果は同じになる筈である。
このコードをコンパイル&実行してみる。

```text
$ go run sample.go 
Hello World
```

うんうん。
ここまでは問題なし。

### パッケージをモジュール化する

では，パッケージ `hello` と `sample` をモジュール化してみよう。

まずはパッケージ `hello` から。

```text
$ go mod init hello
go: creating new go.mod: module hello

$ go test ./...
ok  	hello	(cached)
```

よーし，うむうむ，よーし。

つぎはパッケージ `sample`。

```text
$ go mod init sample
go: creating new go.mod: module sample

$ go run sample.go 
build command-line-arguments: cannot load hello: cannot find module providing package hello
```

ありゃりゃーん。
コンパイルに失敗しちゃったよ。
パッケージ `hello` をロードできないと言っている。
ロード？

ここで `$GOPATH/pkg/linux_amd64` ディレクトリを見るも `hello` に対応するコンパイル済みバイナリは存在しない。
それもその筈で，最近の [Go] コンパイラは外部パッケージのロードとコンパイルをモジュール毎に `$GOPATH/pkg/mod` および `$GOCACHE` ディレクトリ下で行っている[^env2]。
したがってモジュール対応モードでは `$GOPATH/src` 下にある外部パッケージのコードは参照しないのだ。

[^env2]: 古いバージョンの [Go] コンパイラでは環境変数 `GOCACHE` に `off` をセットすることでキャッシュ利用を無効化することができたが，最近のバージョンでは `off` は設定できなくなっている。

さて，困ったね。

## 解決法1： 強制的に GOPATH モードにする

環境変数 `GO111MODULE` の値を `off` にすることにより， `go.mod` の有無に依らず強制的に GOPATH モードでビルドすることができる。
上述の例であれば

```text
$ go env -w GO111MODULE=off

$ go run sample.go 
Hello World
```

とすればよい[^env1]。

[^env1]: コマンド `go env -w` で設定した環境変数の値を削除するには `-u` オプションを使う。例えば `go env -u GO111MODULE` のように使う。

ただし [Go] コンパイラの開発側は GOPATH を将来的に無くす方向で議論および開発を進めているため（後方互換性の観点からバージョン 1.x のうちは大丈夫だろうが）このやり方はあまりオススメできない。
暫定措置というやつである。

## 解決法2： replace ディレクティブを使う

`sample` パッケージの `go.mod` に以下のディレクティブを追記することで，ローカルにある `hello` パッケージのコードを参照するようになる。

```text
replace hello => ../hello
```

`replace` ディレクティブは非常に便利なのだが，どうしても環境依存の記述になってしまうのが欠点である。
たとえば外部の CI/CD サービスと組み合わせる場合はローカルの環境と同じになるように設定をチューニングする必要があるかもしれない。 
これも暫定措置だよね。

## 解決法3： モジュール・パッケージ構成を再設計する

まぁ，身も蓋もない話だが，最終的にはモジュールおよびパッケージの構成を再設計するしかないだろう。

今回の例で言えば `sample` パッケージの構成を

```text
$ tree sample
sample
├── go.mod
├── hello
│   ├── example_test.go
│   └── hello.go
└── sample.go

1 directory, 4 files
```

として `sample.go` を

```go {hl_lines=[3]}
package main

import "sample/hello"

func main() {
	hello.Hello("World")
}
```

と書き換えればいい。
`hello` パッケージを `sample` モジュール内のサブ・パッケージとして再構成するのである。

[Go 言語]における「モジュール」は「パッケージ（群）＋バージョン」であり，しかもバージョン管理は git 等 VCS の機能に依るところが大きい。
故に「1モジュール＝1リポジトリ」を目安にすべきだろう（でないとバージョン管理が煩雑になる）。
その上で「モジュール＝リポジトリ」内に関係の密なパッケージ（群）を組み込んでいくというのが [Go 言語]プログラミング設計の基本的な進め方になると思うのだが，どうだろう。

## ブックマーク

- [Goモジュールモードでモジュール内に作ったモジュールを扱う - Qiita](https://qiita.com/k-motoyan/items/4213d5ef09963ffea489) : モジュール内にサブ・モジュールを構成すること自体は可能である
- [cmd/go: default to GO111MODULE=on · Issue #41330 · golang/go · GitHub](https://github.com/golang/go/issues/41330) : Go 1.17 より `GO111MODULE` の既定が `on` になるらしい

- [モジュール対応モードへの移行を検討する]({{< relref "./go-module-aware-mode.md" >}})
- [Go モジュールのバージョン管理]({{< relref "./versioning-of-go-modules.md" >}})

[Go]: https://go.dev/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
