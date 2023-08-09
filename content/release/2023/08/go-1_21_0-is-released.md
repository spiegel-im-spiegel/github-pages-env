+++
title = "Go 1.21.0 がリリースされた"
date =  "2023-08-09T13:58:32+09:00"
description = "log/slog, slices, maps, cmp パッケージが追加された！"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "logger", "math", "map", "slice" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

予定通り8月に [Go] 1.21.0 がリリースされた。

- [Go 1.21.0 is released](https://groups.google.com/g/golang-announce/c/Mk0Jar6hfhI)
- [Go 1.21 Release Notes - The Go Programming Language](https://go.dev/doc/go1.21)
- [Go 1.21 is released! - The Go Programming Language](https://go.dev/blog/go1.21)

お気づきだろうか。
今回のバージョン番号は 1.21 ではなく 1.21.0 なのですよ。
これからは2月8月の初期リリースでも後ろにちゃんと 1.xx.0 って付けるんだってさ。
これで [Go] の機械的なバージョン管理がしやすくなると思う。

## Windows 7, 8, Server 2008, Server 2012 サポート終了のお知らせ

1.21 からは Windows 10 および Windows Server 2016 が最低要件となる。
それより前の Windows バージョンに対応するには [1.20](https://go.dev/doc/go1.20#windows) 系が必要。
それでも2024年2月でサポートが終わるんだよな。

## log/slog パッケージ

個人的に最大のトピックはこれ。
[`log/slog`] 構造化ログパッケージが追加されたこと。
これでサードパーティの logger も [`log/slog`] パッケージの下に統合されることを期待したい。

ログなんて「詳細」ですよ。
偉い人には分からんのです。

## min/max 組み込み関数

多分，他の言語の人には「今までなかったんかい！」ってツッコまれるんだろうなぁ。
マジでなかったんスよ。
つか，いわゆる Generics が導入されてようやく実現したというか。

## slices, maps, cmp パッケージ

[`slices`], [`maps`], [`cmp`] 各パッケージの追加も Generics 導入の恩恵だよね。
何せ，今までは配列要素の反転（reverse）でさえ自前で

{{< fig-quote class="nobox" type="markdown" title="プログラミング言語Go" link="https://www.amazon.co.jp/dp/B099928SJD?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
```go
func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}
```
{{< /fig-quote >}}

てな感じに書いてたんだぜ。
クールだろ（笑）

こうやって [Go] の標準ライブラリのイケてないところが徐々に解消されていくといいねぇ。
[`math/rand/v2` のプロポーザル](https://github.com/golang/go/issues/61716)とか見てても，着々と [Go] 2 に進んでる気がする。

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.21.0.linux-amd64.tar.gz`](https://go.dev/dl/go1.21.0.linux-amd64.tar.gz)）を取ってきてインストールすることを強く推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.21.0.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.21.0.linux-amd64.tar.gz
$ sudo mv go go1.21.0
$ sudo ln -s go1.21.0 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.21.0 linux/amd64
```

Windows は [Scoop] 経由で OK

複数バージョンの Go コンパイラを扱いたい場合は

```text
$ go install golang.org/dl/go1.21.0@latest
$ go1.21.0 download
$ go1.21.0 version
go version go1.21.0 linux/amd64
```

てな感じで導入できる。

アップデートは計画的に。

## ブックマーク

（随時追加予定）

- [公式の構造化 Logger (になるかもしれない) slog パッケージ]({{< ref "/golang/maybe-official-structured-logger-package.md" >}})

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/
[`log/slog`]: https://pkg.go.dev/log/slog "slog package - log/slog - Go Packages"
[`slices`]: https://pkg.go.dev/slices "slices package - slices - Go Packages"
[`maps`]: https://pkg.go.dev/maps "maps package - maps - Go Packages"
[`cmp`]: https://pkg.go.dev/cmp "cmp package - cmp - Go Packages"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
{{% review-paapi "4873119693" %}} <!-- 実用 Go 言語 -->
{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
{{% review-paapi "4873119979" %}} <!-- Go言語による分散サービス -->
{{% review-paapi "4908686122" %}} <!-- Goならわかるシステムプログラミング 第2版 -->
{{% review-paapi "B09C2XBC2F" %}} <!-- Golang Tシャツ -->
