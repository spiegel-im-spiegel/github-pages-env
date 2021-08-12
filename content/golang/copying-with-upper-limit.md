+++
title = "io.Copy を上限付きで実行する"
date =  "2020-12-06T14:10:36+09:00"
description = "「データ解凍爆弾」への対処"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "compress" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

「[もうlintに怒られない！Goでより安全にzipを扱うために](https://qiita.com/wagi0716/items/362700f84881ca27521a "もうlintに怒られない！Goでより安全にzipを扱うために（Potential DoS vulnerability via decompression bombエラーへの対処法） - Qiita")」を読んで「なるほど！」と思ったが，（おそらくコードをかなり端折ってるんだと思うが）サンプル・コードがツッコミどころ満載なので，うちのブログでも覚え書として記しておく。
[自作ツール](https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer")でも圧縮データの解凍処理を使ってるしね。

## [gosec] について

ところで，上述のリンク先で使われている [gosec] は統合 linter である [golangci-lint] から呼び出すことができるが，既定では無効になっている。
有効にするには `--enable` オプションで [gosec] を指定すればよい。

```text
$ golangci-lint run --enable gosec ./...
```

[gosec] が既定で有効になっていないのは，特に [`unsafe`] 標準パッケージを使っている場合の誤検知が多すぎるかららしい。
私の環境で試してみたが処理速度的には問題ないようなので，一度は試してみてもいいかもしれない。

## 最初のサンプル・コード

まずは起点となるサンプル・コードを挙げておこう。
こんな感じでどうだろう。

```go
package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"strings"
)

func compress(dst io.Writer, src io.Reader) error {
	zw := zlib.NewWriter(dst)
	defer zw.Close()

	if _, err := io.Copy(zw, src); err != nil {
		return err
	}
	return nil
}

func extract(dst io.Writer, src io.Reader) error {
	r, err := zlib.NewReader(src)
	if err != nil {
		return err
	}
	if _, err := io.Copy(dst, r); err != nil {
		return err
	}
	return nil
}

func main() {
	content := "Hello world\n" //raw data
	zbuf := &bytes.Buffer{}
	//compress raw data
	if err := compress(zbuf, strings.NewReader(content)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	//display compressed data
	b := zbuf.Bytes()
	fmt.Printf("%d bytes: %v\n", len(b), b)
	//extract from compressed data
	buf := &bytes.Buffer{}
	err := extract(buf, bytes.NewReader(b))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	//display extracted data
	if _, err := io.Copy(os.Stdout, buf); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
```

最初に紹介した記事で使っている [`archive/zip`] パッケージはファイルシステムを含んでいるので，今回はもっとシンプルに [`compress/zlib`] パッケージ（[RFC 1950] 準拠）を使ってみた。
これならデータの圧縮と解凍のみ注視できるだろう。

このコードを実行すると

```text
$ go run sample.go 
24 bytes: [120 156 242 72 205 201 201 87 40 207 47 202 73 225 2 4 0 0 255 255 28 242 4 71]
Hello world
```

と一応は問題なく動く。
これで準備OK。

## Decompression Bomb Vulnerability

では，このコードに対して lint をかけてみる。

```text
$ golangci-lint run --enable gosec ./...
sample1.go:27:15: G110: Potential DoS vulnerability via decompression bomb (gosec)
	if _, err := io.Copy(dst, r); err != nil {
	             ^
```

おー，出た出た。

“Decompression bomb” というのは

{{< fig-quote type="markdown" title="Zip bomb - Wikipedia" link="https://en.wikipedia.org/wiki/Zip_bomb" lang="en" >}}
{{% quote %}}A **zip bomb**, also known as a **zip of death** or **decompression bomb**, is a malicious [archive file](https://en.wikipedia.org/wiki/Archive_file) designed to crash or render useless the program or system reading it. It is often employed to disable [antivirus software](https://en.wikipedia.org/wiki/Antivirus_software), in order to create an opening for more traditional viruses{{% /quote %}}.
{{< /fig-quote >}}

ということで，主にウイルス対策ツールをターゲットにした攻撃手段らしい。
まぁ `42.zip` みたいなファイルを解凍したら普通にシステムが死ぬと思うけど（笑）

## 解凍処理の改修

“Decompression bomb” を緩和する方法はいくつかあるようだが，いちばん安直なのは [`io`]`.Copy()` 関数に上限を設けることだろう。
そのための関数として [`io`]`.CopyN()` が用意されている。

```go
func io.CopyN(dst io.Writer, src io.Reader, n int64) (written int64, err error)
```

機能としては

{{< fig-quote type="markdown" title="io - The Go Programming Language" link="https://golang.org/pkg/io/#CopyN" lang="en" >}}
{{% quote %}}`CopyN` copies `n` bytes (or until an error) from `src` to `dst`. It returns the number of bytes copied and the earliest error encountered while copying. On return, written `== n` if and only if `err == nil`{{% /quote %}}.
{{< /fig-quote >}}

ということで

1. `err == nil` なら上限いっぱいまで読み込み完了
2. `err` の値が [`io`]`.EOF` ならストリームの終端まで読み込み完了
3. `err` の値が [`io`]`.EOF` 以外ならエラー

と見なすことができそうだ。
そこで最初のサンプルコードの `extract()` 関数を以下のように書き直す。

```go {hl_lines=["10-16"]}
const maxSize = 1024 * 1024 * 1024 //1GB

var ErrTooLarge = errors.New("too laege decompressed data")

func extract(dst io.Writer, src io.Reader) error {
	r, err := zlib.NewReader(src)
	if err != nil {
		return err
	}
	if _, err := io.CopyN(dst, r, maxSize); err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return err
	}
	return ErrTooLarge
}
```

とりあえずコピーの上限は1GBとした（深い意味はない）。
これで少なくとも [gosec] に怒られることはなくなる。

これ仕込むのかぁ。
頑張ろう。

## ブックマーク

- [Secure Go · A project devoted to secure programming in the Go language](https://securego.io/)

- [Go 言語でデータ圧縮と解凍]({{< relref "./compress-data.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[gosec]: https://github.com/securego/gosec "securego/gosec: Golang security checker"
[golangci-lint]: https://github.com/golangci/golangci-lint "golangci/golangci-lint: Fast linters Runner for Go"
[`unsafe`]: https://golang.org/pkg/unsafe/ "unsafe - The Go Programming Language"
[`io`]: https://golang.org/pkg/io/ "io - The Go Programming Language"
[`archive/zip`]: https://golang.org/pkg/archive/zip/ "zip - The Go Programming Language"
[`compress/zlib`]: https://golang.org/pkg/compress/zlib/ "zlib - The Go Programming Language"
[RFC 1950]: https://tools.ietf.org/html/rfc1950 "RFC 1950 - ZLIB Compressed Data Format Specification version 3.3"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
