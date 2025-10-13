+++
title = "struct2flag でコマンドライン解析"
date =  "2025-10-13T03:28:30+09:00"
description = "flag を素のまま使うよりは気軽にコマンドラインオプションを構成できるので，ちょっとしたツールにはいい感じである。"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

いきなりだが，以下のコードから始める。

{{< fig-quote type="markdown" class="nobox" title="『プログラミング言語Go』p.36-37" link="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
```go
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
```
{{< /fig-quote >}}

これはコマンドラインの文字列をそのまま出力するプログラムで，こんなふうに動作する。

```text
$ go run echo4.go May the Force be with you
May the Force be with you
```

`-s` オプションで区切り文字を指定できる。
こんな感じ。

```text
$ go run echo4.go -s / May the Force be with you
May/the/Force/be/with/you
```

また `-n` オプションを指定すると，末尾の改行が抑制される。
さらに `-h` オプションでヘルプが表示される。

```text
$ go run echo4.go -h
Usage of /home/spiegel/.cache/go-build/06/06a6e71bb093bd1ebbb176c5042329730592597ae86dcb2ca99b3759e1aecb18-d/echo4:
  -n	omit trailing newline
  -s string
    	separator (default " ")
```

これらオプションの制御を行っているのが標準パッケージ [`flag`] である。
上述のコード例は簡単なのでアレでいいのだが，オプションに紐付けられた変数をひとつずつ定義して取り回しするのは面倒なので，普通は構造体にまとめる。
例えばこんな感じ。

```go {hl_lines=["9-17","20-21"]}
package main

import (
	"flag"
	"fmt"
	"strings"
)

type Flags struct {
	N   bool
	Sep string
}

func (f *Flags) Bind() {
	flag.BoolVar(&f.N, "n", false, "omit trailing newline")
	flag.StringVar(&f.Sep, "s", " ", "separator")
}

func main() {
	f := &Flags{}
	f.Bind()
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), f.Sep))
	if !f.N {
		fmt.Println()
	}
}
```

オプションを構造体にまとめるメリットはもうひとつあって，それは facade パターンを構成してコマンドライン制御とロジックを簡単に分離できることである。
例えばこんな感じ。

```go {hl_lines=[12,"20-25",31,33]}
package main

import (
	"flag"
	"fmt"
	"strings"
)

type Flags struct {
	N    bool
	Sep  string
	Strs []string
}

func (f *Flags) Bind() {
	flag.BoolVar(&f.N, "n", false, "omit trailing newline")
	flag.StringVar(&f.Sep, "s", " ", "separator")
}

func Echo(f *Flags) {
	fmt.Print(strings.Join(f.Strs, f.Sep))
	if !f.N {
		fmt.Println()
	}
}

func main() {
	f := &Flags{}
	f.Bind()
	flag.Parse()
	f.Strs = flag.Args()

	Echo(f)
}
```

これで `main()` 関数はコマンドライン制御に徹してロジックを `Echo()` 関数に任せることができる。
一方で `Echo()` 関数は渡される `Flags` 構造体にのみ依存しているのでコマンドライン制御の詳細を知らなくて済む。
この例では全てが `main` パッケージにまとまってるので分かりにくいが，ロジックを別パッケージにすることはよくあるので，テストのしやすさも考慮してこのパターンを意識するのは大事である。

とはいえ，ロジックが変われば要求される情報も変わるし，上の例でいう `Flags` 構造体の中身も変わる。
そうなると `Flags.Bind()` メソッドの中身も変える必要があるが，このメソッドで構造体の要素をひとつづつ `flag` パッケージに紐付けているので，これを書くのが地味に面倒くさいのである。
そもそも `flag` パッケージに紐付ける処理が構造体側にあるのもイマイチだよね。

というところで [`github.com/hymkor/struct2flag`] パッケージの登場である。
これを使ってコードを書き直してみる。

```go {hl_lines=["12-13","17-19","29-30"]}
package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/hymkor/struct2flag"
)

type Flags struct {
	N    bool   `flag:"n,omit trailing newline"`
	Sep  string `flag:"s,separator"`
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
	struct2flag.BindDefault(f)
	flag.Parse()
	f.Strs = flag.Args()

	Echo(f)
}
```

[`struct2flag`][`github.com/hymkor/struct2flag`] を使うことにより [`flag`] との紐づけ情報は `Flags` 構造体定義のタグとインスタンス生成・初期化を行う `NewFlags()` 関数に集約される（インスタンスの値が各オプションの既定値として登録されるため）。

[`struct2flag`][`github.com/hymkor/struct2flag`] を使うメリットはいくつかあって

- `main()` 関数側はオプション値を格納する構造体の詳細を知らなくてもよい
- 構造体のタグが各要素の簡単な説明になっている
- [`struct2flag`][`github.com/hymkor/struct2flag`] の中身がシンプルで標準以外の依存パッケージがない

といったところであろう。

私は，ちゃんとしたコマンドラインツールを作るときはサブコマンドや GNU 拡張シンタックスが使える [`cobra`][`github.com/spf13/cobra`] を使ったりしているのだが，ちょっとしたツールには大げさなんだよね，これ。
かといって [`flag`] を素のまま使うのは微妙に鬱陶しくて，結局コマンドラインオプション無しで `main()` 関数に直接値を埋め込んでしまうことが多い。

[`struct2flag`][`github.com/hymkor/struct2flag`] なら構造体にタグを書くだけで済むし [`flag`] を素のまま使うよりは気軽にコマンドラインオプションを構成できるので，ちょっとしたツールにはいい感じである。

公開してくださった [@zetamatta@mstdn.jp ](https://mstdn.jp/@zetamatta "\"ζ\" (@zetamatta@mstdn.jp) - mstdn.jp") さんに感謝 {{% emoji "ペコン" %}}

[Go]: https://go.dev/
[`flag`]: https://pkg.go.dev/flag "flag package - flag - Go Packages"
[`github.com/hymkor/struct2flag`]: https://github.com/hymkor/struct2flag "hymkor/struct2flag: `struct2flag` automatically registers struct fields as flags for your Go command-line tools."
[`github.com/spf13/cobra`]: https://github.com/spf13/cobra "spf13/cobra: A Commander for modern Go CLI interactions"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
