+++
title = "Go 1.17 がリリースされた"
date =  "2021-08-17T21:31:29+09:00"
description = "Zenn 記事の「Go のモジュール管理」は書き直さないとなぁ。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.1７ がリリースされた。

- [Go 1.17 Release Notes - The Go Programming Language](https://golang.org/doc/go1.17)
- [Go 1.17 is released - The Go Blog](https://blog.golang.org/go1.17)

毎度のごとく変更は多岐にわたるが，個人的に気になった部分を挙げておく。

- スライス `[]T` から配列ポインタ `*[N]T` への変換ができるようになったそうな

```go
s := make([]byte, 2, 4)
s0 := (*[0]byte)(s)      // s0 != nil
s1 := (*[1]byte)(s[1:])  // &s1[0] == &s[1]
s2 := (*[2]byte)(s)      // &s2[0] == &s[0]
s4 := (*[4]byte)(s)      // panics: len([4]byte) > len(s)
```

- `unsafe.Add` および `unsafe.Slice` の導入
  - [Big Sky :: unsafe.Add と unsafe.Slice が入った。](https://mattn.kaoriya.net/software/lang/go/20210504002438.htm)
- `windows/arm64` アーキテクチャのサポート
- `go.nmod` で `go 1.17` 以降が指定されている場合，間接的な依存モジュールも記述するよう要求されるようになった。とりあえず `go mod tidy -go=1.17` とすれば `go` ディレクティブごといい感じに更新してくれるらしい
  - [Go 1.17連載が始まります: コンパイラとgo mod | フューチャー技術ブログ](https://future-architect.github.io/articles/20210810a/)
- `go.nmod` で `module` ディレクティブの直前行に `// Deprecated: comment` とつけることで，非推奨のワーニングを出せるようになった

```go
// Deprecated: use example.com/mod/v2 instead.
module example.com/mod
```

- ビルドの制御を行う `// +build` が `//go:build` に変更になる。とりあえず 1.17 では両方有効らしい。サンプルファイルとか `// +build run` とか山ほど書いてるんだが，どうしよう...
- `go run` コマンドがモジュール・バージョンを受け入れるようになった。これ GitHub Actions とかで重宝しそうだな（笑）

```text
$ go run github.com/spiegel-im-spiegel/gpgpdump@latest -h
go: downloading github.com/spiegel-im-spiegel/gpgpdump v0.12.4
...
OpenPGP (RFC 4880) packet visualizer by golang.

Usage:
  gpgpdump [flags]
...
```

- スタックではなくレジスタ・ベースの関数呼び出しがサポートされた。5%くらい速くなるらしい。サイズもちょびっと小さくなるようだ
- 関数閉包（closure）を含む関数をインライン展開できるようになった
- `rune` 型の変数が負値の場合に関連するメソッドがエラーとなるよう挙動が変わったようだ
  - [Go 1.17からの負のruneの扱い | フューチャー技術ブログ](https://future-architect.github.io/articles/20210817a/)
- `go test` に `-shuffle` オプションが追加された。テスト順序をシャッフルしてくれるらしい
  - [Go 1.17のtesting新機能 | フューチャー技術ブログ](https://future-architect.github.io/articles/20210812a/)
- [`time`]`.Time` 型に `GoString()` メソッドが追加された。なんだこれ（笑）

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("%#v\n", time.Now())
	// Output:
	// time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local)
}
```

例によって [Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.17.linux-amd64.tar.gz`](https://golang.org/dl/go1.17.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する。
以下は手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://golang.org/dl/go1.17.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.17.linux-amd64.tar.gz
$ sudo mv go go1.17
$ sudo ln -s go1.17 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.17 linux/amd64
```

アップデートは計画的に。
Zenn 記事の「[Go のモジュール管理](https://zenn.dev/spiegel/articles/20210223-go-module-aware-mode)」は書き直さないとなぁ。

## ブックマーク

- [Go1.17のencoding/csv | フューチャー技術ブログ](https://future-architect.github.io/articles/20210811a/)
- [Go1.17における go get の変更点 | フューチャー技術ブログ](https://future-architect.github.io/articles/20210818a/)
- [Go 1.17 Release祝い 細かすぎて伝わらないMinor change in flag package](https://zenn.dev/hgsgtk/articles/9f662a4c96fa3f)

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[`time`]: https://pkg.go.dev/time "time · pkg.go.dev"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
