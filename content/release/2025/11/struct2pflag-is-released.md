+++
title = "hymkor/struct2flag から fork した goark/struct2pflag をリリースした"
date =  "2025-11-23T18:33:12+09:00"
description = "どうしても GNU 拡張シンタックスが使える spf13/pflag が使いたかったのよ。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "package", "module" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

以前 [golang]({{< rlnk "/golang/" >}} "プログラミング言語 Go") セクションで[紹介]({{< ref "/golang/parsing-command-line-flags-with-struct2flag.md" >}} "struct2flag でコマンドライン解析【追記あり】")した [`hymkor/struct2flag`] だが，これは標準の [`flag`] パッケージを使っている。
他の類似の flag パッケージついては

{{< fig-quote type="markdown" title="Goの構造体タグによるコマンドラインオプション・設定ファイル・環境変数の一元管理 - 標準愚痴出力" link="https://zetamatta.hatenablog.com/entry/2025/10/17/190842" >}}
オプション体系が GNU っぽくなってしまったり、`"flag"` の機能まで内包していて、そこまで要らないということで選定候補から外した
{{< /fig-quote >}}

とのことで，それはそれで妥当な選択だと思う。

ただ，私は [GNU 拡張シンタックス](https://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html "Argument Syntax (The GNU C Library)")に対応した [`spf13/pflag`] が使いたかったのよ。

というわけで，オリジナルの [`hymkor/struct2flag`] から fork して [`spf13/pflag`] を使うよう改造した [`goark/struct2pflag`] をリリースした[^f1]。
使い方は [`hymkor/struct2flag`] とほぼ同じ。
以前書いた「[struct2flag でコマンドライン解析]({{< ref "/golang/parsing-command-line-flags-with-struct2flag.md" >}} "struct2flag でコマンドライン解析【追記あり】")」のコードを [`goark/struct2pflag`] を使って書き換えてみる。

[^f1]: 最初は GitHub の fork 機能をそのまま使っていたが自己 pull request の際に，うっかりオリジナルの [`hymkor/struct2flag`] の方に PR しそうになったため fork 関係を切り離した。やり方は「[GithubでForkしたリポジトリをFork元から切り離して管理する](https://zenn.dev/misora/articles/62f36f175e5416)」を参考にした。ありがたや {{% emoji "ペコン" %}}

```go {hl_lines=["7-8", "12-13", "30-32"]}
package main

import (
	"fmt"
	"strings"

	"github.com/goark/struct2pflag"
	"github.com/spf13/pflag"
)

type Flags struct {
	N    bool   `pflag:"no-newline,n,omit trailing newline"`
	Sep  string `pflag:"sep,s,separator"`
	Strs []string
}

func NewFlags() *Flags {
	return &Flags{N: false, Sep: " ", Strs: []string{}}
}

func Echo(f *Flags) {
	fmt.Print(strings.Join(f.Strs, f.Sep))
	if !f.N {
		fmt.Println()
	}
}

func main() {
	f := NewFlags()
	struct2pflag.BindDefault(f)
	pflag.Parse()
	f.Strs = pflag.Args()

	Echo(f)
}
```

[`flag`] が [`spf13/pflag`] に [`hymkor/struct2flag`] が [`goark/struct2pflag`] に置き換わったほか，構造体タグのプレフィックスが `flag:` から `pflag:` に変わっている点に注目。
ヘルプを表示させると

```text
$ go run echo4a.go -h
Usage of /tmp/go-build2259633867/b001/exe/echo4a:
  -n, --no-newline   omit trailing newline
  -s, --sep string   separator (default " ")
```

てな感じに長いオプション名と一文字の短いオプション名が併記されているのが分かる。
実際に動かしてみると

```text
$ go run echo4a.go --sep / May the Force be with you
May/the/Force/be/with/you

$ go run echo4a.go -s / May the Force be with you
May/the/Force/be/with/you
```

と長いオプション名でも短いオプション名でも同じ動作になることが分かる。
[`spf13/pflag`] の挙動については拙文「[標準 flag パッケージを pflag パッケージに置き換える](https://zenn.dev/spiegel/articles/20221008-customize-flag-package)」で簡単に紹介しているので，そちらも参考にどうぞ。

[`hymkor/struct2flag`] で使う構造体タグ `flag:` も受け入れるようにした（`flag:` と `pflag:` は混在可能）。

```go
type Flags struct {
	N    bool   `flag:"n,omit trailing newline"`
	Sep  string `flag:"s,separator"`
	Strs []string
}
```

ただし，この場合は長いオプション名のみ定義されるため

```text
$ go run echo4b.go -h
Usage of /tmp/go-build1365887832/b001/exe/echo4b:
      --n          omit trailing newline
      --s string   separator (default " ")
```

という感じに，あくまでも [`spf13/pflag`] のルールで構成されるのでご注意を。

一応 [`hymkor/struct2flag`] と同程度の機能は持たせている。
今後は浮動小数点数や時間を指定できたらいいな，と思っている。
他に [`spf13/pflag`] ではオプション値を格納する変数側を slice にすることで同じオプションを複数回指定できる機能もあるので，それもサポートしたいところ。
まぁ，あくまでも私が使う際に必要に応じて拡張するということで。
リクエスト等ありましたら PR いただけると助かります。

[`hymkor/struct2flag`]: https://github.com/hymkor/struct2flag "hymkor/struct2flag: `struct2flag` automatically registers struct fields as flags for your Go command-line tools."
[`goark/struct2pflag`]: https://github.com/goark/struct2pflag "goark/struct2pflag: `struct2pflag` automatically registers struct fields as flags for your Go command-line tools."
[`flag`]: https://pkg.go.dev/flag "flag package - flag - Go Packages"
[`spf13/pflag`]: https://github.com/spf13/pflag "spf13/pflag: Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags."
<!-- eof -->
