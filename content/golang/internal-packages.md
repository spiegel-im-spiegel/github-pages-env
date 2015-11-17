+++
date = "2015-11-07T10:38:20+09:00"
update = "2015-11-17T10:46:12+09:00"
description = "Internal Packages の仕組みは 1.4 から存在したが標準パッケージのみの適用だった。 1.5 からは GOPATH 配下のパッケージまで拡張される。"
draft = false
tags = ["golang", "package" ]
title = "パッケージ外部からの呼び出しを禁止する Internal Packages"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

（初出： [そろそろ真面目に Golang 開発環境について考える — Internal Packages と Vendoring - Qiita](http://qiita.com/spiegel-im-spiegel/items/baa3671c7e1b8a6594a9)）

「[GOPATH 汚染問題]({{< relref "golang/gopath-pollution.md" >}})」で言及しそこねたので，今回は軽く。

Internal Packages の仕組みは 1.4 から存在したが標準パッケージのみの適用だった。
1.5 からは `GOPATH` 配下のパッケージまで拡張される。

- [Go 1.4 "Internal" Packages](https://docs.google.com/document/d/1e8kOo3r51b2BWtTs_1uADIA5djfXhPT36s6eHVRIvaU/edit)

要するに `internal` フォルダ以下のパッケージは外部から参照できない。
例として [`net`] パッケージを挙げてみる。
[`net`] パッケージのソースコードの構成は以下のようになっている。

```
C:\Go\src\net> tree .
C:\GO\SRC\NET
├─http
│  ├─cgi
│  │  └─testdata
│  ├─cookiejar
│  ├─fcgi
│  ├─httptest
│  ├─httputil
│  ├─internal
│  ├─pprof
│  └─testdata
├─internal
│  └─socktest
├─mail
├─rpc
│  └─jsonrpc
├─smtp
├─testdata
├─textproto
└─url
```

[`net/http/internal`](https://golang.org/pkg/net/http/internal/) パッケージには `chunked.go` ファイルが含まれる。
`chunked.go` は [chunked transfer encoding](https://en.wikipedia.org/wiki/Chunked_transfer_encoding) の仕組みを実装したもののようだが， [`net/http`] およびその配下のパッケージ以外では使えない。
また `net/internal` フォルダ以下には socket テスト用の [`net/internal/socktest`] パッケージがあるが，これも外部から再利用できない。

これは再利用の難しいパッケージを配置する場合にはよい仕掛けである。
ビジネスロジックには再利用が難しいものが多いので，そういったものを `internal` フォルダ以下に置けば，うっかり他所で使われるといったこともなく安全である。
なお， Internal Packages の制約から外すには `internal` フォルダの外側にパッケージを再配置すればよい。

## ブックマーク

[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`net`]: https://golang.org/pkg/net/ "net - The Go Programming Language"
[`net/internal/socktest`]: https://golang.org/pkg/net/internal/socktest/ "socktest - The Go Programming Language"
[`net/http`]: https://golang.org/pkg/net/http/ "http - The Go Programming Language"
[`net/http/internal`]: https://golang.org/pkg/net/http/internal/ "internal - The Go Programming Language"
