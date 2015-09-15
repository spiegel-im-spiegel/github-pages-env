# Golang による文字エンコーディング変換

文字エンコーディング変換に関してはあちこちに記事があるのだが，微妙に古い気がするので，メモとして書き記しておく。

## Go 言語の文字エンコーディング

Go 言語では Unicode が既定となっている。文字の単位である `rune` は `int32` でおそらく UTF-32 相当だし，コード上の文字エンコーディングも UTF-8 である（「[その4](http://qiita.com/spiegel-im-spiegel/items/556166b6631c0369754f)」参照）。

したがって UTF-8 以外の文字エンコーディングを扱う場合は何らかの変換処理を挟む必要がある。そのためのパッケージが [transform] である。また[エンコーディング](https://godoc.org/golang.org/x/text/encoding)についても各種そろっていて，日本語の場合は [encoding/japanese] パッケージを使う。 [encoding/japanese] パッケージでは Shift-JIS, EUC-JP, ISO-2202-JP を扱える。

```go:encoding/japanese/all.go
// All is a list of all defined encodings in this package.
var All = []encoding.Encoding{EUCJP, ISO2022JP, ShiftJIS}
```

## パッケージの導入

日本語が必要なだけなら `golang.org/x/text/encoding/japanese` を `go get` すれば全てインストールされる。

```shell
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

（かつてパッケージの場所は `code.google.com/p/go.text/transform` および `code.google.com/p/go.text/encoding/japanese` だったが，ここの repository は今は存在しないので注意）

## 変換ロジック（サンプルコード）

変換ロジックのサンプルを以下に示す。

```go:conversion.go
package main

import (
	"bufio"
	"fmt"
	"io"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

//Conversion
func Conversion(inStream io.Reader, outStream io.Writer) error {
	//read from stream (Shift-JIS to UTF-8)
	scanner := bufio.NewScanner(transform.NewReader(inStream, japanese.ShiftJIS.NewDecoder()))
	list := make([]string, 0)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	//write to stream (UTF-8 to EUC-JP)
	writer := bufio.NewWriter(transform.NewWriter(outStream, japanese.EUCJP.NewEncoder()))
	for _, line := range list {
		_, err = fmt.Fprintln(writer, line)
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
```

Shift-JIS ─（Decoder）→ UTF-8 ─（Encoder）→ EUC-JP の手順で変換していることがお分かりであろうか。

このように文字エンコーディング変換は stream 処理として行う。サンプルコードはテキストファイルを想定しているが， stream として扱えるならファイル以外でも使えるはずである。

コメントでファイルを丸ごと一気に読み込んで変換するコード例を紹介していただきました。こちらもどうぞ。

- [ShiftJIS -> EUC-JP by golang](https://gist.github.com/tyochiai/239c5433872e4c9cb517)

## 文字エンコーディング変換に関する注意点

Shift-JIS/EUC-JP と UTF-8/32 ではベースとなっている文字集合（文字エンコーディングではない）が異なる。 Shift-JIS/EUC-JP の文字集合は基本的に JIS 規格だが UTF-8/32 は Unicode であり，両者は非対称の関係である。このため，今回のような異なる文字集合を跨ぐ変換を行うと変換が正しく行われない場合もあり得る。

更に困ったことに，歴史的経緯（便利な言葉だw）から Shift-JIS や EUC-JP の実装にはいくつかバリエーションがあるため，実装間の差異が問題になる場合もある。

更に更に言えば，汎用機などは旧JIS＋外字の構成になっていることが多く，このような需要に応えるなら独自の変換ロジックを開発するしかない。

### 変換ロジックの別解

これまで述べた方法でうまく変換ができない場合は [djimenez/iconv-go] パッケージを使う手もある。ただし， [djimenez/iconv-go] パッケージのビルドには libiconv および glibc が必要である（クロス環境では注意）。このパッケージを使うことで iconv 相当の処理が可能になる。

（別の変換パッケージとして [mahonia](https://godoc.org/code.google.com/p/mahonia) を紹介しているところがいくつか見られたが，ドキュメントを見る限り DEPRECATED となっていて使えないようだ。また repository にもアクセスできない）

## ブックマーク

- [golang - Go言語で文字コード変換 - Qiita](http://qiita.com/uchiko/items/1810ddacd23fd4d3c934)
- [go - Goで[]byteをshift-jisの文字列に変換する - スタック・オーバーフロー](http://ja.stackoverflow.com/questions/6120/go%E3%81%A7byte%E3%82%92shift-jis%E3%81%AE%E6%96%87%E5%AD%97%E5%88%97%E3%81%AB%E5%A4%89%E6%8F%9B%E3%81%99%E3%82%8B)
- [GO言語で文字コードを扱うライブラリの使用例 - Qiita](http://qiita.com/irugo/items/390bd187871c7716a1e1)
- [Golangで文字コード判定 - Qiita](http://qiita.com/nobuhito/items/ff782f64e32f7ed95e43)

[transform]: https://godoc.org/golang.org/x/text/transform "transform - GoDoc"
[encoding/japanese]: https://godoc.org/golang.org/x/text/encoding/japanese "japanese - GoDoc"
[djimenez/iconv-go]: https://github.com/djimenez/iconv-go "djimenez/iconv-go"
