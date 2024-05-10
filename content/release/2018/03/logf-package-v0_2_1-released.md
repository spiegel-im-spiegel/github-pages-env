+++
title = "Logf Package v0.2.1 Released"
date = "2018-03-03T20:11:38+09:00"
description = "本当は JSON 形式の構造化ログに挑戦しようと思ったのだが，既に出来のいいパッケージが存在したので止めた。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "golang", "package", "logger" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

レベル付き logger パッケージ [spiegel-im-spiegel/logf] v0.2.1 をリリースした。

- [spiegel-im-spiegel/logf: Simple leveled logging package by Golang](https://github.com/spiegel-im-spiegel/logf)

標準の [log] パッケージをラップする形で実装している。
以下のように [lestrrat-go/file-rotatelogs] のようなローテーション機能付きの Writer と組み合わせて使うといい感じになる。

```go
package main

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/spiegel-im-spiegel/logf"
)

func main() {
	rl, err := rotatelogs.New("./log.%Y%m%d%H%M.txt")
	if err != nil {
		logf.Fatal(err)
		return
	}
	logger := logf.New(
		logf.WithFlags(logf.LstdFlags|logf.Lshortfile),
		logf.WithPrefix("[Sample] "),
		logf.WithWriter(rl),
		logf.WithMinLevel(logf.INFO),
	)
	logger.Print("Information")
	//Output:
	//[Sample] 2009/11/10 23:00:00 sample.go:20: [INFO] Information
}
```

本当は JSON 形式の構造化ログに挑戦しようと思ったのだが，既に出来のいいパッケージが存在したので止めた。

- [rs/zerolog: Zero Allocation JSON Logger](https://github.com/rs/zerolog)

なので，これ以上 [spiegel-im-spiegel/logf] を弄ることはないと思う。

## ブックマーク

- [Log パッケージで遊ぶ]({{< ref "/golang/logger.md" >}})

[spiegel-im-spiegel/logf]: https://github.com/spiegel-im-spiegel/logf "spiegel-im-spiegel/logf: Simple leveled logging package by Golang"
[log]: https://golang.org/pkg/log/ "log - The Go Programming Language"
[lestrrat-go/file-rotatelogs]: https://github.com/lestrrat-go/file-rotatelogs "lestrrat-go/file-rotatelogs: Port of perl5 File::RotateLogs to Go"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
