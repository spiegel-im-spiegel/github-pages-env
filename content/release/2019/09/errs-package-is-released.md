+++
title = "Go 言語用エラーハンドリング・パッケージをリリースした"
date =  "2019-09-05T23:54:45+09:00"
description = "実は以前にこっそり v0.1.0 をリリースして自作ツールのエラーハンドリングに用いていたのだが， Go 1.13 のリリースに合わせて中身を作り直した。"
image = "/images/attention/go-logo_blue.png"
tags = [ "package", "module", "golang", "error" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Go 言語用エラーハンドリング・パッケージ [`errs`] の v0.2.0 をリリースした。

- [Release v0.2.0 · spiegel-im-spiegel/errs · GitHub](https://github.com/spiegel-im-spiegel/errs/releases/tag/v0.2.0)

実は以前にこっそり v0.1.0 をリリースして自作ツールのエラーハンドリングに用いていたのだが， [Go] 1.13 のリリースに合わせて中身を作り直した。

使い方は以下を参照のこと。

- [Go 言語用エラーハンドリング・パッケージ]({{< ref "/release/errs-package-for-golang.md" >}})

これで先に進めるな。

## 【2019-09-06 追記】 v0.2.1 をリリースした

- [Release v0.2.1 · spiegel-im-spiegel/errs · GitHub](https://github.com/spiegel-im-spiegel/errs/releases/tag/v0.2.1)

つか，実はバージョンタグを付け間違えただけなのだが，バージョンタグを付け換えるとチェックサム・データベース [`sum.golang.org`] が怒って

```text
SECURITY ERROR
This download does NOT match the one reported by the checksum server.
The bits may have been replaced on the origin server, or an attacker may
have intercepted the download attempt.

For more information, see 'go help module-auth'.
```

とか言ってくさるので，しょうが無しにバージョン番号を上げることにした。
やれやれ。

バージョンタグの管理は慎重に。

[`sum.golang.org`]: https://sum.golang.org/

## 【2019-09-15 追記】 v0.2.2 をリリースした

- [Release v0.2.2 · spiegel-im-spiegel/errs · GitHub](https://github.com/spiegel-im-spiegel/errs/releases/tag/v0.2.2)

Formatting に関するバグを修正した。

## ブックマーク

- [Go 1.13 のエラー・ハンドリング]({{< ref "/golang/error-handling-in-go-1_3.md" >}})

[Go]: https://go.dev/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`errs`]: https://github.com/spiegel-im-spiegel/errs "spiegel-im-spiegel/errs: Error handling for Golang"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
