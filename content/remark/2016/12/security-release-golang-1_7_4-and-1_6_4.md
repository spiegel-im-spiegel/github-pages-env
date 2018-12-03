+++
description = "数日前から予告されていたが， Go 言語の 1.7.4 と 1.6.4 がリリースされた。詳細が分かり次第，ここに追記する。"
tags = [
  "security",
  "vulnerability",
  "golang",
  "x509",
]
draft = false
date = "2016-12-02T20:32:33+09:00"
title = "Security Release Go 1.7.4 and 1.6.4"

[author]
  instagram = "spiegel_2007"
  flickr = "spiegel"
  linkedin = "spiegelimspiegel"
  tumblr = ""
  facebook = "spiegel.im.spiegel"
  github = "spiegel-im-spiegel"
  license = "by-sa"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
  flattr = "spiegel"
  name = "Spiegel"
  avatar = "/images/avatar.jpg"
+++

数日前から[予告](https://groups.google.com/forum/#!topic/golang-announce/YOqTqcJtiJI)されていたが， [Go 言語]の 1.7.4 と 1.6.4 がリリースされた。

- [[security] Go 1.7.4 and Go 1.6.4 are released](https://groups.google.com/forum/#!topic/golang-announce/2lP5z9i9ySY)

セキュリティ脆弱性を含むので（特に Web 関連で [Go 言語]を使ってる方は）必ずアップデートすること。
なお，特に理由がない限り 1.7 系を使うことをお勧めする。

ひとつは [`crypto/x509`] パケージに関するもの

- [Change If681c514: crypto/x509: read Darwin trust settings for root CAs | go-review.googlesource Code Review](https://go-review.googlesource.com/#/c/33721/)
- [crypto/x509: honor OS X certificate trust settings · Issue #18141 · golang/go](https://github.com/golang/go/issues/18141)

もうひとつは [`net/http`] パッケージに関するものだ。

- [Change Ib394655b: net/http: multipart ReadForm close file after copy | go-review.googlesource Code Review](https://go-review.googlesource.com/#/c/30410/)
- [net/http: backport "multipart ReadForm close file after copy" to 1.7 · Issue #17965 · golang/go](https://github.com/golang/go/issues/17965)

CVE 番号とかはまだ付いてないのかな？
詳細が分かり次第，ここに追記する。

そうそう。
1.8 ベータ版も登場している。

- [Go 1.8 Beta 1 is released](https://groups.google.com/forum/#!topic/golang-announce/Wgv6NGcntlQ)
- [Go 1.8 Release Notes - The Go Programming Language](https://beta.golang.org/doc/go1.8)

`GOPATH` 環境変数にデフォルト値ができたらしい。
あと `go bug` コマンドってなんだ？

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`crypto/x509`]: https://golang.org/pkg/crypto/x509/ "x509 - The Go Programming Language"
[`net/http`]: https://golang.org/pkg/net/http/ "http - The Go Programming Language"
