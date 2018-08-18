+++
date = "2015-09-23T20:04:00+09:00"
update = "2017-12-08T13:47:07+09:00"
description = "文字エンコーディング変換に関してはあちこちに記事があるのだが，微妙に古い気がするので，メモとして書き記しておく。"
draft = false
tags = ["golang", "character", "encoding", "transform"]
title = "文字エンコーディング変換"

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

（初出： [Golang による文字エンコーディング変換 - Qiita](http://qiita.com/spiegel-im-spiegel/items/2e475b48226330aa5570)）

文字エンコーディング変換に関してはあちこちに記事があるのだが，微妙に古い気がするので，メモとして書き記しておく。

## Go 言語の文字エンコーディング変換

Go 言語では Unicode が既定となっている。
そもそもソースコードが UTF-8 前提になっているし，文字の単位である [rune] の実体は Unicode 符号点（code point）である[^a]。

[^a]: 「[String と Rune]」参照。

したがって UTF-8 以外の文字エンコーディングを扱う場合は何らかの変換処理を挟む必要がある。
そのためのパッケージが [`transform`] である。
また[エンコーディング](https://godoc.org/golang.org/x/text/encoding)についても各種そろっていて，日本語の場合は [`encoding/japanese`] パッケージを使う。
[`encoding/japanese`] パッケージでは Shift-JIS, EUC-JP, ISO-2202-JP を扱える。

### パッケージの導入

日本語が必要なだけなら `golang.org/x/text/encoding/japanese`[^b] を `go get` すれば全てインストールされる。

[^b]: かつてパッケージの場所は `code.google.com/p/go.text/transform` および `code.google.com/p/go.text/encoding/japanese` だったが，ここの repository は今は存在しないので注意。

```
C:>go get -v golang.org/x/text/encoding/japanese
Fetching https://golang.org/x/text/encoding/japanese?go-get=1
Parsing meta tags from https://golang.org/x/text/encoding/japanese?go-get=1 (status code 200)
get "golang.org/x/text/encoding/japanese": found meta tag main.metaImport{Prefix:"golang.org/x/text", VCS:"git", RepoRoot:"https://go.googlesource.com/text"} at https://golang.org/x/text/encoding/japanese?go-get=1
get "golang.org/x/text/encoding/japanese": verifying non-authoritative meta tag
Fetching https://golang.org/x/text?go-get=1
Parsing meta tags from https://golang.org/x/text?go-get=1 (status code 200)
golang.org/x/text (download)
golang.org/x/text/transform
golang.org/x/text/encoding/internal/identifier
golang.org/x/text/encoding
golang.org/x/text/encoding/internal
golang.org/x/text/encoding/japanese
```

### 変換ロジック（サンプルコード）

変換ロジックのサンプルを以下に示す（thanks [@mattn](http://qiita.com/mattn)）。

```go
package main

import (
    "fmt"
    "io"
    "os"

    "golang.org/x/text/encoding/japanese"
)

func main() {
    reader := japanese.ShiftJIS.NewDecoder().Reader(os.Stdin)
    writer := japanese.EUCJP.NewEncoder().Writer(os.Stdout)
    if _, err := io.Copy(writer, reader); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
}
```

UTF-8 以外の文字エンコーディングから UTF-8 への変換は decode と呼ばれている。
一方， UTF-8 から UTF-8 以外の文字エンコーディングへの変換は encode と呼ばれている。

上のサンプルでは Shift-JIS →（Decoder）→ UTF-8 →（Encoder）→ EUC-JP の手順で変換していることがお分かりであろうか。
では実際に動かしてみよう。

```
C:>echo Go言語で行こう | go run transform.go > euc.txt
```

これで `euc.txt` に「Go言語で行こう」と EUC-JP で書き込まれていたら成功である[^c]。
今回は標準入出力を使ったが， Reader/Writer の stream で表せるものなら同様の処理でできるはずである。

[^c]: もちろんこれは Windows のコマンドプロンプトで動かした場合の話。

## 文字エンコーディング変換に関する注意点

Shift-JIS/EUC-JP と UTF-8/UTF-32 ではベースとなっている文字集合（文字エンコーディングではない）が異なる。
Shift-JIS/EUC-JP の文字集合は基本的に JIS 規格だが UTF-8/UTF-32 は Unicode であり，両者は非対称の関係である。
このため，今回のような異なる文字集合を跨ぐ変換を行うと変換が正しく行われない場合もあり得る（要検証[^d]）。

[^d]: たとえば「[Shift_JIS文化からUTF-8への移行ガイド - Qiita](http://qiita.com/kawasima/items/41632dbd423dc0445e14)」などが参考になる。

更に困ったことに，歴史的経緯（便利な言葉だw）から Shift-JIS や EUC-JP の実装にはいくつかバリエーションがあるため，実装間の差異が問題になる場合もある。
[`encoding/japanese`] パッケージはこの実装上の差異を考慮してはいないようである。

更に更に言えば，汎用機などは旧JIS＋外字の構成になっていることが多く，このような需要に応えるなら独自の変換ロジックを開発するしかない。

### 変換ロジックの別解

[`transform`] を使った変換でうまくいかない場合は [`djimenez/iconv-go`] パッケージを使う手もある。
ただし， [`djimenez/iconv-go`] パッケージのビルドには `libiconv` および `glibc` が必要である（クロス環境では注意）。
このパッケージを使うことで `iconv` 相当の処理が可能になる。

（別の変換パッケージとして [`mahonia`](https://godoc.org/code.google.com/p/mahonia) を紹介しているところがいくつか見られたが，ドキュメントを見る限り DEPRECATED となっていて使えないようだ。また repository にもアクセスできない）

## ブックマーク

- [golang - Go言語で文字コード変換 - Qiita](http://qiita.com/uchiko/items/1810ddacd23fd4d3c934)
- [go - Goで[]byteをshift-jisの文字列に変換する - スタック・オーバーフロー](http://ja.stackoverflow.com/questions/6120/go%E3%81%A7byte%E3%82%92shift-jis%E3%81%AE%E6%96%87%E5%AD%97%E5%88%97%E3%81%AB%E5%A4%89%E6%8F%9B%E3%81%99%E3%82%8B)
- [GO言語で文字コードを扱うライブラリの使用例 - Qiita](http://qiita.com/irugo/items/390bd187871c7716a1e1)
- [Golangで文字コード判定 - Qiita](http://qiita.com/nobuhito/items/ff782f64e32f7ed95e43)

[Go 言語に関するブックマーク集はこちら]({{< relref "bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[String と Rune]: {{< relref "string-and-rune.md" >}}
[string]: http://golang.org/ref/spec#String_types
[rune]: http://blog.golang.org/strings "Strings, bytes, runes and characters in Go - The Go Blog"
[`transform`]: https://godoc.org/golang.org/x/text/transform "transform - GoDoc"
[`encoding/japanese`]: https://godoc.org/golang.org/x/text/encoding/japanese "japanese - GoDoc"
[`djimenez/iconv-go`]: https://github.com/djimenez/iconv-go "djimenez/iconv-go"
