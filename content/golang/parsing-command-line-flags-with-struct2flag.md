+++
title = "struct2flag でコマンドライン解析【追記あり】"
date =  "2025-10-13T03:28:30+09:00"
description = "flag を素のまま使うよりは気軽にコマンドラインオプションを構成できるので，ちょっとしたツールにはいい感じである。"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

## flag 標準パッケージによるコマンドライン解析

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
  -n    omit trailing newline
  -s string
        separator (default " ")
```

これらオプションの解析を行っているのが標準パッケージ [`flag`] である。
上述のコード例は簡単なのでアレでいいのだが，オプションに紐付けられた変数をひとつずつ定義して取り回しするのは面倒なので，普通は構造体にまとめる。
例えばこんな感じ（エラーハンドリングをサボっているがご容赦）。

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

## struct2flag でコマンドライン解析

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

## 【2025-10-18 追記】オプションの既定値を設定ファイルから取得する

前節では `Flags` 構造体の初期値を固定値にしていたが，これを設定ファイルから取得できるようにしてみる。
設定ファイルの内容は以下の JSON 形式とする。

```json
{
  "no_newline": true,
  "separator": "/"
}
```

さらに `Flags` 構造体のタグを以下のように書き換える。

```go
type Flags struct {
    N    bool   `flag:"n,omit trailing newline" json:"no_newline"`
    Sep  string `flag:"s,separator" json:"separator"`
    Strs []string
}
```

読み込む処理は以下の通り（`main()` 関数のみ表示）。

```go {hl_lines=[1, "5-12"]}
const configPath = "echo/config.json"

func main() {
    f := NewFlags()
    if conf, err := os.UserConfigDir(); err == nil {
        if data, err := os.ReadFile(filepath.Join(conf, configPath)); err == nil {
            if err := json.Unmarshal(data, f); err != nil {
                fmt.Fprintln(os.Stderr, err)
                return
            }
        }
    }
    struct2flag.BindDefault(f)
    flag.Parse()
    f.Strs = flag.Args()

    Echo(f)
}
```

これで設定ファイルが存在する場合はその内容を規定値とし，存在しない場合（または読み込みに失敗した場合）はスルーする。
[`os`]`.UserConfigDir()` って何？ って方は（ちょっと古いけど）拙文「[go-homedir はもう要らない]({{< relref "./no-need-go-homedir.md" >}})」あたりを参考にどうぞ。

実行結果はこんな感じ。

```text
$ go run echo4e.go May the Force be with you
May/the/Force/be/with/you
```

オプション値を設定ファイルで取得できるようにすると何が嬉しいかというと，バッチ処理を構成しやすくなるのよ。
設定ファイル自体が「どういう条件でコマンドを実行したか」という証跡になるし，コマンドラインを短くできるのも嬉しい。

## 【2025-10-18 追記】オプションの既定値を環境変数から取得する

[`struct2flag`][`github.com/hymkor/struct2flag`] の作者の方が構造体のタグと環境変数を紐付けて取得するパッケージ [`github.com/hymkor/struct2env`] も公開された。
構造体のタグはこんな感じに定義する。

```go
type Flags struct {
    N    bool   `flag:"n,omit trailing newline" env:"NO_NEWLINE"`
    Sep  string `flag:"s,separator" env:"SEPARATOR"`
    Strs []string
}
```

読み込む処理は以下の通り（`main()` 関数のみ表示）。

```go {hl_lines=[3]}
func main() {
    f := NewFlags()
    struct2env.Bind(f)
    struct2flag.BindDefault(f)
    flag.Parse()
    f.Strs = flag.Args()

    Echo(f)
}
```

実行結果はこんな感じ。

```text
$ env SEPARATOR=/ go run echo4f.go May the Force be with you
May/the/Force/be/with/you
```

まぁ，今回みたいにコマンドラインオプションと環境変数で同じ情報を取得するのはメリットが薄いかもしれないけど（Docker 環境では重宝するかも），ある値は環境変数から取得し別の値はコマンドラインから取得する，みたいな使い方はアリかなと思ったりする。
ともかく，環境変数の情報を（構造体のタグ付けで）一度に取得できるのはかなり便利である。

### 【2025-10-19 追記】別解

[`struct2env`][`github.com/hymkor/struct2env`] ほど便利ではないが，一応別解もある。

```go {hl_lines=["4-8"]}
func main() {
    f := NewFlags()
    struct2flag.BindDefault(f)
    flag.VisitAll(func(f *flag.Flag) {
        if s := os.Getenv(strings.ToUpper(f.Name)); s != "" {
            _ = f.Value.Set(s)
        }
    })
    flag.Parse()
    f.Strs = flag.Args()

    Echo(f)
}
```

これはオプション名と同名の環境変数を探して既定値としてセットするもので，実行するとこんな感じになる。

```text
$ env S=/ go run echo4g.go May the Force be with you
May/the/Force/be/with/you
```

標準パッケージのみで構成できるのがメリットだが，このままだと1文字のオプション名もそのまま環境変数名として検索されることになるので，実際に使うには変換ルールを工夫する必要があるかもしれない。
それを考えると [`struct2env`][`github.com/hymkor/struct2env`] のほうが自由度が高くていいかも。

## ブックマーク

- [Goの構造体タグによるコマンドラインオプション・設定ファイル・環境変数の一元管理 - 標準愚痴出力](https://zetamatta.hatenablog.com/entry/2025/10/17/190842) : 作者による公式解説
- [Big Sky :: Re: Goでコマンドライン引数と環境変数の両方からflagを設定したい](https://mattn.kaoriya.net/software/lang/go/20170609110526.htm)

[Go]: https://go.dev/
[`flag`]: https://pkg.go.dev/flag "flag package - flag - Go Packages"
[`os`]: https://pkg.go.dev/os "os package - os - Go Packages"
[`github.com/hymkor/struct2flag`]: https://github.com/hymkor/struct2flag "hymkor/struct2flag: `struct2flag` automatically registers struct fields as flags for your Go command-line tools."
[`github.com/hymkor/struct2env`]: https://github.com/hymkor/struct2env "hymkor/struct2env: `struct2env` binds environment variables to struct fields according to their env tags."
[`github.com/spf13/cobra`]: https://github.com/spf13/cobra "spf13/cobra: A Commander for modern Go CLI interactions"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
