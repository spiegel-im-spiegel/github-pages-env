# はじめての Go 言語 (on Windows) その10

（[これまでの記事の目次](http://qiita.com/spiegel-im-spiegel/items/98d49ac456485b007a15#%E3%81%AF%E3%81%98%E3%82%81%E3%81%A6%E3%81%AE-go-%E8%A8%80%E8%AA%9E-on-windows)。同ページのブックマークも参考にどうぞ）

ご無沙汰でした。今回でこのシリーズを最終回にします。ここまでくるともう「はじめて」じゃないし（笑）

## コマンドライン・ツールを作ろう

Go 言語の使いどころはいろいろあると思いますが，今回はコマンドライン・ツールを作ります。一般的なコマンドライン・ツールの要件は以下のとおり。

- shell（bash やコマンドプロンプトなど）からの起動を前提とする
- コマンドラインの引数または標準入力からデータや条件を入力する
- 結果を標準出力に出力する
- 結果以外の情報は標準エラー出力に出力する
- shell へ実行結果の状態を返す

「shell へ実行結果の状態を返す」というのは分かりにくいかもしれませんが，これはコマンドライン・ツールが shell に対して返す値で， bash などであれば `$?`，コマンドプロンプトであれば `ERRORLEVEL` に格納される値です。正常終了する場合は 0 を，異常終了であれば 0 以外をセットするのがお約束になっています。

題材を何にしようか考えていましたが，先日面白い題材を見つけたので，これを利用したいと思います。

- [Git.io 短縮 URL を golang コードで取得してみる - Qiita](http://qiita.com/spiegel-im-spiegel/items/042751d98e315e4e3382)

## Git.io 短縮 URL 取得ツールを作る

「[Git.io 短縮 URL を golang コードで取得してみる](http://qiita.com/spiegel-im-spiegel/items/042751d98e315e4e3382)」で最終的に提示したコードがこれ。

```go:gitio3.go
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	resp, err := http.PostForm("http://git.io", url.Values{"url": {"https://github.com/spiegel-im-spiegel"}})
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("  Status: ", resp.Header.Get("Status"))
	log.Println("Location: ", resp.Header.Get("Location"))
	log.Println("    Body: ", string(body))
}
```

このコードをベースにして，引数で渡された URL を短縮 URL に変換して標準出力に返すツールを作ってみます。

まずは素直に変形。

```go:gitio4.go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	//arguments
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
		return
	}
	urlStr := os.Args[1]

	//shortening url
	resp, err := http.PostForm("http://git.io", url.Values{"url": {urlStr}})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Fprintln(os.Stderr, resp.Header.Get("Status"))
	if string(body) != urlStr {
		fmt.Fprintln(os.Stderr, string(body))
	}
	fmt.Fprint(os.Stdout, resp.Header.Get("Location"))
}
```

このうち変換処理の部分を [spiegel-im-spiegel/gitioapi](https://github.com/spiegel-im-spiegel/gitioapi) に外出しします。引数解析の部分も [flag](https://golang.org/pkg/flag/) パッケージを使った方法に変えてみます（Thanx @mattn）。

```go:gitio5.go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spiegel-im-spiegel/gitioapi"
)

func main() {
	//arguments
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
		return
	}
	urlStr := flag.Arg(0)

	//shortening url
	shortUrl, err := gitioapi.Encode(&gitioapi.Param{Url: urlStr})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Print(shortUrl)
}
```

だいぶスッキリしました。これでも一応動作します。

```shell
C:>go run gitio5.go https://github.com/spiegel-im-spiegel
http://git.io/vOj52
```

### このコードの問題点

このコードの問題点は大きく2つあります。

1. `main` 関数に直接ロジックを書いているためテストが出来ない（特に引数の解析部分）
2. `main` 関数の `return` では shell に値を返せない

このうち2番目は `os.Exit()` 関数を使うことで解決しますが， `os.Exit()` 関数は `defer` 構文と相性が悪いのが難点です。いずれにしろ `main` 関数にロジックを直書きするというのは実用的なコードではありえない話なので，ここからもう少し変形していきます。

### Go のエラーハンドリング

ここで，ちょろんと寄り道。

なのだが，エラーハンドリングについては以下の記事にて全面的に書きなおした。

- [エラー・ハンドリングについて - プログラミング言語 Go | text.Baldanders.info](http://text.baldanders.info/golang/error-handling/)

## tcnksm/gcli を使ってコマンドライン・ツール用のコードを自動生成する

話を元に戻しましょう。

コマンドライン・ツールに関しては同様の悩みを持っておられる方が多いらしく，様々な支援パッケージが公開されています。この中で，今回は [tcnksm/gcli] を紹介します。

[tcnksm/gcli] は，以前は cli-init という名前で公開されていたもののようです。

- [高速にGo言語のCLIツールをつくるcli-initというツールをつくった | SOTA](http://deeeet.com/writing/2014/06/22/cli-init/)

[tcnksm/gcli] では更に機能が強化されていて， [codegangsta/cli](http://deeeet.com/writing/2014/06/22/cli-init/) 以外にも，標準パッケージの [flag](https://golang.org/pkg/flag/) や [mitchellh/cli](https://github.com/mitchellh/cli) にも対応しているようです。

### tcnksm/gcli のインストール

[tcnksm/gcli] のインストールは少々コツがいるっぽいです。また UNIX 系の環境を前提としているようで Windows 環境ではインストールが少々面倒くさい感じです。ここでは手順としてコマンドラインの流れを紹介します。

```shell
C:>go get -d -v github.com/tcnksm/gcli
C:>cd %GOPATH%\src\github.com\tcnksm\gcli
C:\path\to\gopath\src\github.com\tcnksm\gcli>go get -v golang.org/x/tools/cmd/vet
C:\path\to\gopath\src\github.com\tcnksm\gcli>go get -v github.com/golang/lint/golint
C:\path\to\gopath\src\github.com\tcnksm\gcli>go get -v github.com/jteeuwen/go-bindata/...
C:\path\to\gopath\src\github.com\tcnksm\gcli>go get -v -d -t ./...
C:\path\to\gopath\src\github.com\tcnksm\gcli>git describe --always
3ff629e
C:\path\to\gopath\src\github.com\tcnksm\gcli>pushd skeleton
C:\path\to\gopath\src\github.com\tcnksm\gcli\skeleton>go-bindata -pkg="skeleton" resource/...
C:\path\to\gopath\src\github.com\tcnksm\gcli\skeleton>popd
C:\path\to\gopath\src\github.com\tcnksm\gcli>go install -ldflags "-X main.GitCommit \"3ff629e\""
```

`go get` コマンドの `-d` オプションは GOPATH へのダウンロードのみでインストールを行わないようにするためのオプションです。また `git describe --always` で取得した commit ID を `go install` 時にバイナリに埋め込んでいます。

これで `gcli` コマンドがインストールされました。最後に動作確認します。

```shell
C:>gcli version
[0;0mgcli version v0.2.0 (3ff629e)[0m
[0;31m
Your versin of gcli is out of date! The latest version is 0.2.1.[0m
```

おおっと，ターミナル前提か。まぁいいや。バージョンが古いと怒られているのは `version.go` ファイルの記述が古いままになってるからのようです。これを直せば

```shell
C:>gcli version
gcli version v0.2.1 (3ff629e)
```

となります（エスケープシーケンスは削除しています，見た目がウザいので）。あんまりメンテされてないのかなぁ。

### tcnksm/gcli によるコード自動生成

今回は標準の [flag](https://golang.org/pkg/flag/) パッケージを使った一番簡単なコードを生成してみます。

```shell
C:>gcli new -F flag -flag=c:String gitio6
Created gitio6\main.go
Created gitio6\CHANGELOG.md
Created gitio6\version.go
Created gitio6\cli_test.go
Created gitio6\README.md
Created gitio6\cli.go
====> Successfully generated gitio6
```

ここではオプションとして `-c` を追加してみました。生成されたファイルの内訳は以下のとおりです。

- `main.go` : `main` 関数
- `cli.go` : 処理本体
- `cli_test.go` : `cli.go` のテスト
- `version.go` : バージョン情報

いやぁ，テストまで自動生成してくれるってありがたいです。ちなみに `main.go` は以下のようになっています。

```go:main.go
package main

import "os"

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
```

`cli` インスタンスに標準（エラー）出力および引数情報を渡しているのがわかると思います。テストの際はこれらにテスト用の情報をセットしてテストするわけです。これはこれで完成形なので，ほぼ弄る必要はありません（標準入力も対象にするなら手直しする必要あり）。

`cli.go` が実際にロジックを書くところです。ここでは `cli.go` の最終形のみ示します。

```go:cli.go
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/spiegel-im-spiegel/gitioapi"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		c string
		url string
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.StringVar(&c, "c", "", "'code' parameter.")

	flVersion := flags.Bool("version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if *flVersion {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	// Parse argument
	switch flags.NArg() {
	case 0 :
		fmt.Fprintln(cli.errStream, os.ErrInvalid, "No GitHub URL")
		return ExitCodeError
	case 1 :
		url = flags.Arg(0)
	default :
		fmt.Fprintln(cli.errStream, os.ErrInvalid, flags.Arg(1))
		return ExitCodeError
	}

	// shortening URL
	shortUrl, err := gitioapi.Encode(&gitioapi.Param{Url: url, Code: c})
	if err != nil {
		fmt.Fprintln(cli.errStream, err)
		return ExitCodeError
	}
	fmt.Fprint(cli.outStream, shortUrl)

	return ExitCodeOK
}
```

`// Parse argument` 以下の行が追記したロジックです。この程度の処理ならほとんど書くことありませんよね（笑）

一応，動作確認。

```shell
C:>go build
C:>gitio6.exe https://github.com/spiegel-im-spiegel
http://git.io/vOj52
C:>gitio6.exe -c t https://github.com/technoweenie
http://git.io/t
```

今回の構成を Domain-Driven 的に考えるのなら， `CLI` クラスが Application Service で `gitioapi` パッケージが Domain Service に相当するのかな。だいぶ無理矢理ですけど。

### 自動生成されたコードの著作権

今回使用した [tcnksm/gcli] を含めいわゆる CASE（Computer Aided Software Engineering）が生成するコードは誰に帰属するのでしょう。

- [著作権審議会第9小委員会（コンピュータ創作物関係）報告書 | 著作権審議会/文化審議会分科会報告 | 著作権データベース | 公益社団法人著作権情報センター　CRIC](http://www.cric.or.jp/db/report/h5_11_2/h5_11_2_main.html)
- [コンピュータ自動生成物は著作物ではない - ものがたり](http://d.hatena.ne.jp/atsushieno/20060903/p1)
- [自動生成物の著作権について、すでに判決が下された判例が（あれ… - 人力検索はてな](http://q.hatena.ne.jp/1339209099)
- [コンピュータが自動生成した創作物 - Footprints](http://d.hatena.ne.jp/redips/20141124/1416825529)

「自動生成」がどこまでのレベルを指すのかにもよると思いますし，個々の判例では「自動生成されたコードに創作性はない」とするものもあるようですが，一般的に「こう」と言えるものはない感じです。

自動生成ツールを公開する場合は，生成されたコードの扱いについて明記しておくのが（今のところは）安全だと思います。

この記事で紹介したコードは，個人的には「実証コード」レベルの品質だと思っているので， [CC0](http://creativecommons.org/publicdomain/zero/1.0/deed.ja) で公開しています。ただし自動生成されたコードについては，ツールの製作者の方が何らかのライセンスが必要であるとするのであれば [CC0](http://creativecommons.org/publicdomain/zero/1.0/deed.ja) を撤回する場合もあると思います。

これ AI が進歩して完動品のコードを生成するようになったら，どうなるんでしょうねぇ。

## 最後に

最初に述べたように，「はじめての Go 言語 (on Windows)」シリーズは今回で最終回にします。とはいえ，個人的には Go 言語はかなり気に入ってるので，仕事で使えるレベルにはもう少し勉強していきたいと思います。そのときは別途記事を立てることもあるかと思います。

では，また。

## ブックマーク

- [Go言語によるCLIツール開発とUNIX哲学について - ゆううきブログ](http://yuuki.hatenablog.com/entry/go-cli-unix)
- [Go Conference 2015 summer - Qiita](http://qiita.com/t-sato/items/a5d1a309733e765533ce)
- [開発者から見た UNIX 哲学とコマンドラインツールと Go言語 - TELLME.TOKYO](http://tellme.tokyo/post/2015/06/23/unix_cli_tool_go/) （[Qiita 版](http://qiita.com/b4b4r07/items/df660d82e2de715acda5)）
- [Go言語でテストしやすいコマンドラインツールをつくる | SOTA](http://deeeet.com/writing/2014/12/18/golang-cli-test/)

[tcnksm/gcli]: https://github.com/tcnksm/gcli "The easy way to build Golang command-line application."
