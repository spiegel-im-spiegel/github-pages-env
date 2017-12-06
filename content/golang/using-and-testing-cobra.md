+++
title = "Cobra の使い方とテスト"
date =  "2017-12-06T21:01:33+09:00"
description = "spf13/cobra そのものについてちゃんと書いてない気がするので，今回はコードの書き方からテストまでをひと通り紹介していく。"
image = "/images/attention/go-code.png"
tags = ["golang", "cli", "programming", "testing"]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = true
  mermaidjs = true
+++

だいぶ [spf13/cobra] の扱いに慣れてきたので，そろそろブログにまとめておこうかなと。
以前に CLI (Command-Line Interface) とファサード・パターンについては以下の記事に

- [コマンドライン・インタフェースとファサード・パターン]({{< relref "golang/cli-and-facade-pattern.md" >}})

[spf13/cobra] の使い方についてさわりの部分は以下の記事に書いた。

- [モンテカルロ法による円周率の推定（その2 CLI）]({{< relref "golang/estimate-of-pi-2-cli.md" >}})
- [Codic API を利用するパッケージを作ってみた]({{< relref "golang/codic-api.md" >}}) （[spf13/viper] との連携について）

ただ [spf13/cobra] そのものについてちゃんと書いてない気がするので，今回はコードの書き方からテストまでをひと通り紹介していく。
漫然と紹介するのもアレなので，今回は以下の CLI を作ることを目標にする。

```text
$ cli-demo show -i 123 -s 文字列 -b
Integer option value: 123
 String option value: 文字列
Boolean option value: true
```

## [spf13/cobra] のインストール

[spf13/cobra] には CLI ツールもあるのだがバイナリは提供されていないので，自前で取ってきてコンパイルする。
といっても

```text
$ go get -u github.com/spf13/cobra/cobra
```

とするだけだが。
`github.com/spf13/cobra` で get しようとすると CLI のコンパイルを行わないので注意が必要である。
CLI の Usage はこんな感じ。

```text
$ cobra -h
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  cobra [command]

Available Commands:
  add         Add a command to a Cobra Application
  help        Help about any command
  init        Initialize a Cobra Application

Flags:
  -a, --author string    author name for copyright attribution (default "YOUR NAME")
      --config string    config file (default is $HOME/.cobra.yaml)
  -h, --help             help for cobra
  -l, --license string   name of license for the project
      --viper            use Viper for configuration (default true)

Use "cobra [command] --help" for more information about a command.
```

## 初期化処理

まずは `init` サブコマンドで雛形となるソースコードを展開する。
ソースを展開したいフォルダで

```text
cobra --viper=false init .
```

とする。
今回は [spf13/viper] は使わないので `--viper` オプションは無効にしておく。

これで対象のフォルダに3つのファイルが出力される

- `main.go`
- `LICENSE`
- `cmd/root.go`

`LICENSE` はライセンス・ファイルなのでそのまま利用するなり他のファイルに差し替えるなりすればいい。
`main.go` と `cmd/root.go` が実際の雛形になるのだが，やたらコメントが多いので，整理したものを以下に示す。

まずは `main.go`。

```go
package main

import "github.com/spiegel-im-spiegel/cli-demo/cmd"

func main() {
	cmd.Execute()
}
```

次に `cmd/root.go`。

```go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cli-demo",
	Short: "Short comment",
	Long:  "Long comment",
	//Run: func(cmd *cobra.Command, args []string) { },
}

//Execute is called from main function
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}
```

呼び出し関係は以下の通り。

{{< fig-mermaid >}}
graph LR
  main["main()"]
  exec1["cmd.Execute()"]
  exec2["cmd.RootCmd.Execute()"]

  main-->exec1
  exec1-->exec2
{{< /fig-mermaid >}}

もちろん `main()` 関数から直接 `cmd.RootCmd.Execute()` 関数を呼んでも構わない（このままであれば）。
この状態でコマンドを起動すると以下のようになる。

```text
$ go run main.go -h
Long comment
```

まぁ，当然だよね。
現状での問題は3つある。

1. `cmd.RootCmd` を他パッケージから直接操作できてしまう
2. 引数と標準入出力を `cmd.RootCmd` が握っていてコントロールできない（ように見える）
3. `cmd.RootCmd.Execute()` 関数がエラー時に `os.Exit()` 関数で強制終了してしまう

特に2番目が致命的。
何故ならこのままではテストができないからだ。
そこで `cmd.Execute()` 関数を改造することにする。

## [spiegel-im-spiegel/gocli] の導入

その前に標準入出力の取り回しを簡単にするために，手前味噌で申し訳ないが， [spiegel-im-spiegel/gocli] を導入する。
これを使えば標準入出力をひとつのインスタンスで取り回しできる。

```go
ui := gocli.NewUI(
    gocli.Reader(os.Stdin),
    gocli.Writer(os.Stdout),
    gocli.ErrorWriter(os.Stderr),
)
ui.Outputln("Hello world")
```

`gocli.NewUI()` の初期化は Functional Options パターンを使っている。
Functional Options パターンについて詳しくは以下を参照のこと。

- [インスタンスの生成と Functional Options パターン]({{< relref "golang/functional-options-pattern.md" >}})

## cmd.Execute() 関数を改造する

では `cmd.Execute()` 関数の改造を行おう。
途中を省いて結果は以下の通り。

```go
var (
	cui = gocli.NewUI() //CUI instance
)

//Execute is called from main function
func Execute(ui *gocli.UI, args []string) (exit ExitCode) {
    defer func() {
		//panic hundling
		if r := recover(); r != nil {
			cui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, src, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				cui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ":", src, ":", line)
			}
			exit = ExitAbnormal
		}
	}()

	//execution
	cui = ui
	rootCmd.SetArgs(args)
	rootCmd.SetOutput(ui.ErrorWriter())
	exit = ExitNormal
	if err := rootCmd.Execute(); err != nil {
		exit = ExitAbnormal
	}
	return
}
```

前半は [panic] 時のハンドリングで，スタック追跡できるようにしている。
[panic] 時のハンドリングについて詳しくは以下を参照のこと。

- [スタック追跡とパニック・ハンドリング]({{< relref "golang/stack-trace-and-panic-handling.md" >}})

後半では `cmd.rootCmd` に引数と標準出力をセットしてから `cmd.rootCmd.Execute()` 関数を起動している[^so1]。
エラー時は `os.Exit()` 関数で強制終了するのではなく，ちゃんとステータスを返すようにした。

[^so1]: [spf13/cobra] では出力先がひとつしかなく何故か標準出力と標準エラー出力を区別していない。そこで `cmd.rootCmd` には標準エラー出力をセットして標準出力へのアクセスは `cui` を `cmd` パッケージ内のどのメソッドからも参照できるようにした。 [spf13/cobra] 側が標準出力と標準エラー出力をちゃんと区別してくれれば，もう少しスマートにできるんだけどねぇ。
そうそう。
`cmd.RootCmd` ではパッケージ外から直接操作できてしまうので `cmd.rootCmd` と小文字にしている。
小さなことからコツコツと（笑）

これで `main()` 関数もこのように書き換えることができる。

```go
func main() {
	os.Exit(cmd.Execute(
		gocli.NewUI(
			gocli.Reader(os.Stdin),
			gocli.Writer(os.Stdout),
			gocli.ErrorWriter(os.Stderr),
		),
		os.Args[1:],
	).Int())
}
```

## サブコマンドの追加

ではいよいよサブコマンド `show` を追加する。
サブコマンドの追加は `main.go` のあるフォルダで以下のコマンドを実行する。

```text
$ cobra add show
```

これにより `cmd/show.go` ファイルが追加された。
これも中身はコメントが山ほどあるので，整理したものを以下に示す。

```go
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Short comment for show sub-command",
	Long:  "Long comment for show sub-command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("show called")
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
```

この状態で実際に動かしてみるとこうなる。

```text
$ go run main.go -h
Long comment

Usage:
  cli-demo [flags]
  cli-demo [command]

Available Commands:
  help        Help about any command
  show        Short comment for show sub-command

Flags:
  -h, --help   help for cli-demo

Use "cli-demo [command] --help" for more information about a command.

$ go run main.go show -h
Long comment for show sub-command

Usage:
  cli-demo show [flags]

Flags:
  -h, --help   help for show

$ go run main.go show
show called
```

### オプションの追加

このコードに対してオプションを追加していく。
こんな感じでいいだろう。

```go
package cmd

import (
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Short comment for show sub-command",
	Long:  "Long comment for show sub-command",
	RunE: func(cmd *cobra.Command, args []string) error {
		i, err := cmd.Flags().GetInt("integer")
		if err != nil {
			return err
		}
		b, err := cmd.Flags().GetBool("boolean")
		if err != nil {
			return err
		}
		s, err := cmd.Flags().GetString("string")
		if err != nil {
			return err
		}
		cui.Outputln("Integer option value:", i)
		cui.Outputln(" String option value:", s)
		cui.Outputln("Boolean option value:", b)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	showCmd.Flags().IntP("integer", "i", 0, "integer option")
	showCmd.Flags().BoolP("boolean", "b", false, "boolean option")
	showCmd.Flags().StringP("string", "s", "", "string option")
}
```

`showCmd` 内の `Run` が `RunE` に変わってる点に注目。
これにより，関数内で error が発生した場合に，それを返り値で渡すことができる。 

試しにちょろんと動かしてみよう。

```text
$ go run main.go show -h
Long comment for show sub-command

Usage:
  cli-demo show [flags]

Flags:
  -b, --boolean         boolean option
  -h, --help            help for show
  -i, --integer int     integer option
  -s, --string string   string option

$ go run main.go show -i 123 -s 日本語 -b
Integer option value: 123
 String option value: 日本語
Boolean option value: true
```

おおっ。
なんか動いてるような気がする。

じゃあテストを始めようか。

## cobra.Command を返す関数を作る。

と言いたいところだけど，このままではまだテストができない。
何故かというと， `rootCmd` も `showCmd` も static な変数として定義されているので，そのままテストを繰り返すと前回の状態が残ってしまって正しいテストにならないからだ。

じゃあどうすればいいかというと， [`cobra`].`Command` を返す関数を作って，その中で `rootCmd` や `showCmd` に相当するインスタンスを作ればいいのである。

じゃあ，まずは `cmd/root.go` の完全版から。

```go
package cmd

import (
	"runtime"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/gocli"
)

var (
	cui = gocli.NewUI() //CUI instance
)

//newRootCmd returns cobra.Command instance for root command
func newRootCmd(ui *gocli.UI, args []string) *cobra.Command {
	cui = ui
	rootCmd := &cobra.Command{
		Use:   "cli-demo",
		Short: "Short comment",
		Long:  "Long comment",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("no command")
		},
	}
	rootCmd.SetArgs(args)
	rootCmd.SetOutput(ui.ErrorWriter())

	rootCmd.AddCommand(newShowCmd())

	return rootCmd
}

//Execute is called from main function
func Execute(ui *gocli.UI, args []string) (exit ExitCode) {
	defer func() {
		//panic hundling
		if r := recover(); r != nil {
			cui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, src, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				cui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ":", src, ":", line)
			}
			exit = ExitAbnormal
		}
	}()

	//execution
	exit = ExitNormal
	if err := newRootCmd(ui, args).Execute(); err != nil {
		exit = ExitAbnormal
	}
	return
}
```

次は `cmd/show.go` の完全版。

```go
package cmd

import (
	"github.com/spf13/cobra"
)

//newShowCmd returns cobra.Command instance for show sub-command
func newShowCmd() *cobra.Command {
	showCmd := &cobra.Command{
		Use:   "show",
		Short: "Short comment for show sub-command",
		Long:  "Long comment for show sub-command",
		RunE: func(cmd *cobra.Command, args []string) error {
			i, err := cmd.Flags().GetInt("integer")
			if err != nil {
				return err
			}
			b, err := cmd.Flags().GetBool("boolean")
			if err != nil {
				return err
			}
			s, err := cmd.Flags().GetString("string")
			if err != nil {
				return err
			}
			cui.Outputln("Integer option value:", i)
			cui.Outputln(" String option value:", s)
			cui.Outputln("Boolean option value:", b)

			return nil
		},
	}

	showCmd.Flags().IntP("integer", "i", 0, "integer option")
	showCmd.Flags().BoolP("boolean", "b", false, "boolean option")
	showCmd.Flags().StringP("string", "s", "", "string option")

	return showCmd
}
```

なんで [spf13/cobra] が最初からこういう雛形を作ってくれないかというと，おそらく `cobra add` コマンドで `cmd/root.go` を触るわけにはいかないため， static 変数にして `init()` 関数でコマンドを繋げるような組み方しかできなかったんだと思う。
でも手作業でこういう雛形を作ってしまえば，以後はコピペでいくらでも量産できるので，面倒でも最初に手間を掛けたほうがいいかもしれない。

## テストを書く

ようやくテストが書けるよ。
とりあえず正常系のみ。

```go
package cmd

import (
	"bytes"
	"testing"

	"github.com/spiegel-im-spiegel/gocli"
)

func TestShowNormal(t *testing.T) {
	testCases := []struct {
		args []string
		want string
	}{
        {args: []string{"show", "-i", "123", "-s", "日本語", "-b"}, want: "Integer option value: 123\n String option value: 日本語\nBoolean option value: true\n"},
		{args: []string{"show", "-i", "123", "-s", "日本語"}, want: "Integer option value: 123\n String option value: 日本語\nBoolean option value: false\n"},
		{args: []string{"show", "-i", "123"}, want: "Integer option value: 123\n String option value: \nBoolean option value: false\n"},
		{args: []string{"show"}, want: "Integer option value: 0\n String option value: \nBoolean option value: false\n"},
	}

	for _, c := range testCases {
		out := new(bytes.Buffer)
		errOut := new(bytes.Buffer)
		ui := gocli.NewUI(
			gocli.Writer(out),
			gocli.ErrorWriter(errOut),
		)
		exit := Execute(ui, c.args)
		if exit != ExitNormal {
			t.Errorf("Execute() err = \"%v\", want \"%v\".", exit, ExitNormal)
		}
		if out.String() != c.want {
			t.Errorf("Execute() Stdout = \"%v\", want \"%v\".", out.String(), c.want)
		}
		if errOut.String() != "" {
			t.Errorf("Execute() Stderr = \"%v\", want \"%v\".", errOut.String(), "")
		}
	}
}
```

標準出力と標準エラー出力を [`bytes`].`Buffer` で代替えしているのがお分かりだろうか。
これなら CLI でもかなりの部分をテストでカバーできる。
ここには挙げていないが，当然パイプのテストも可能である。

## [dep] で [spf13/cobra] を管理する

[spf13/cobra] をはじめとする外部パッケージを [dep] を使って `vendor/` 配下に置くのなら， `Gopkg.toml` の記述は以下でいいだろう。

```toml
[[constraint]]
  name = "github.com/spf13/cobra"
  version = "~0.0.1"

[[constraint]]
  name = "github.com/spiegel-im-spiegel/gocli"
  version = "~0.4.0"

[[constraint]]
  name = "github.com/pkg/errors"
  version = "~0.8.0"
```

ちなみに今回のデモのパッケージ依存関係はこんな感じになっている。

```text
$ dep status -dot | dot -Tpng -o dependency.png
```
 
{{< fig-img src="/images/cli-demo-dependency.png" width="732" >}}


## ブックマーク

- [Golangのコマンドライブラリcobraを使って少しうまく実装する - Qiita](https://qiita.com/tkit/items/3cdeafcde2bd98612428)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[panic]: http://blog.golang.org/defer-panic-and-recover "Defer, Panic, and Recover - The Go Blog"
[pkg/errors]: https://github.com/pkg/errors "pkg/errors: Simple error handling primitives"
[spf13/cobra]: https://github.com/spf13/cobra "spf13/cobra: A Commander for modern Go CLI interactions"
[`cobra`]: https://github.com/spf13/cobra "spf13/cobra: A Commander for modern Go CLI interactions"
[spf13/viper]: https://github.com/spf13/viper "spf13/viper: Go configuration with fangs"
[spf13/pflag]: https://github.com/spf13/pflag "spf13/pflag: Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags."
[spiegel-im-spiegel/gocli]: https://github.com/spiegel-im-spiegel/gocli "spiegel-im-spiegel/gocli: Command line interface"
[`bytes`]: https://golang.org/pkg/bytes/ "bytes - The Go Programming Language"
[dep]: https://github.com/golang/dep "golang/dep: Go dependency management tool"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-20</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4865940391/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4865940391.09._SCTHUMBZZZ_.jpg"  alt="Kotlinスタートブック -新しいAndroidプログラミング"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4839959234/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4839959234.09._SCTHUMBZZZ_.jpg"  alt="Docker実戦活用ガイド"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274218961/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274218961.09._SCTHUMBZZZ_.jpg"  alt="グッド・マス ギークのための数・論理・計算機科学"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-13">2016-07-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/477418392X/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/61EL3Dc95dL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/477418392X/baldandersinf-22/">みんなのGo言語【現場で使える実践テクニック】</a></dt><dd>松木雅幸 mattn 藤原俊一郎 中島大一 牧 大輔 鈴木健太 稲葉貴洋 </dd><dd>技術評論社 2016-09-09</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774184322/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774184322.09._SCTHUMBZZZ_.jpg"  alt="WEB+DB PRESS Vol.95"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4621300253.09._SCTHUMBZZZ_.jpg"  alt="プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274219151/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274219151.09._SCTHUMBZZZ_.jpg"  alt="プログラミングElixir"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798147400/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798147400.09._SCTHUMBZZZ_.jpg"  alt="詳解MySQL 5.7 止まらぬ進化に乗り遅れないためのテクニカルガイド (NEXT ONE)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774182869/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774182869.09._SCTHUMBZZZ_.jpg"  alt="WEB+DB PRESS Vol.94"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117763/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117763.09._SCTHUMBZZZ_.jpg"  alt="Docker"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/477418361X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/477418361X.09._SCTHUMBZZZ_.jpg"  alt="オブジェクト指向設計実践ガイド ~Rubyでわかる 進化しつづける柔軟なアプリケーションの育て方"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4295000337/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4295000337.09._SCTHUMBZZZ_.jpg"  alt="WebデベロッパーのためのReact開発入門 JavaScript UIライブラリの基本と活用"  /></a> </p>
<p class="description">リファレンス本なのに索引が貧弱。これなら Kindle で買ってもよかったか。 1.7 への言及があるのはよかった。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-11-17">2016-11-17</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
