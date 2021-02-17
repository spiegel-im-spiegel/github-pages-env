+++
title = "Go 1.16 がリリースされた"
date =  "2021-02-17T19:16:04+09:00"
description = "毎度のごとく多岐にわたる変更あり。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.16 がリリースされた。

- [Go 1.16 is released](https://groups.google.com/g/golang-announce/c/q-exD-8mG3w)

毎度のごとく変更は多岐にわたるが，個人的に気になった部分を挙げておく。

1. macOS の 64bit ARM アーキテクチャをサポート。 `darwin/arm64` で指定可能。またこれにより iOS ポートが `ios/arm64` または `ios/amd64` にリネームされた（後者は iOS シミュレータ用）
1. `GO111MODULE` 環境変数の既定値が `on` になった。以前の状態に戻すには `auto` を指定すること
1. `go install` コマンドの引数がバージョン番号サフィックスを受け入れるようになった
    - 例： `go install example.com/cmd@v1.0.0`
1. `go build` や `go test` などのビルド系コマンド実行時に go.mod` や go.sum` ファイルの変更を行わなくなった
1. `go.mod` によるモジュール設定で `retract` ディレクティブが追加された
1. ラベルのない for ループ，メソッド値，型スイッチのある関数をインライン化できるようになった
1. [`embed`] パッケージおよび `//go:embed`  ディレクティブの追加
1. [`io/fs`] パッケージの追加によりファイルシステムの汎化が容易になった
1. [`io/ioutil`] パッケージの非推奨化
1. [`crypto/dsa`] パッケージの非推奨化。これに伴い [`crypto/x509`] で DSA による電子署名検証機能がドロップされた（電子署名生成は以前からないらしい）
1. [`os/signal`]`.NotifyContext()` 関数の追加。これにより SIGNAL 到着時に [`context`]`.Context` によるキャンセルを容易に実装できる
1. [`strconv`] のパフォーマンス改善（最大で2倍らしい）

例によって [Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.16.linux-amd64.tar.gz`](https://golang.org/dl/go1.16.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する。
以下は手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://golang.org/dl/go1.16.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.16.linux-amd64.tar.gz
$ sudo mv go go1.16
$ sudo ln -s go1.16 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.16 linux/amd64
```

アップデートは計画的に。

## ブックマーク

- [Go 1.16 is released - The Go Blog](https://blog.golang.org/go1.16)
- [Go 1.16 Release Notes - The Go Programming Language](https://golang.org/doc/go1.16)
- [Go 1.16 リリースノート 日本語訳  - Qiita](https://qiita.com/c-yan/items/f6f504d5e1c1caceace4)

[Go]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[`embed`]: https://golang.org/pkg/embed/ "embed - The Go Programming Language"
[`io/fs`]: https://golang.org/pkg/io/fs/ "fs - The Go Programming Language"
[`io/ioutil`]: https://golang.org/pkg/io/ioutil/ "ioutil - The Go Programming Language"
[`crypto/dsa`]: https://golang.org/pkg/crypto/dsa/ "dsa - The Go Programming Language"
[`crypto/x509`]: https://golang.org/pkg/crypto/x509/ "x509 - The Go Programming Language"
[`os/signal`]: https://golang.org/pkg/os/signal/ "signal - The Go Programming Language"
[`context`]: https://golang.org/pkg/context/ "context - The Go Programming Language"
[`strconv`]: https://golang.org/pkg/strconv/ "strconv - The Go Programming Language"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
