+++
title = "指定桁の数字列の先頭をゼロで埋める遊び"
date =  "2017-10-09T17:05:06+09:00"
description = "今回は 文字列⇔数値 変換と正規表現（のさわり）について書けたからよしとするか。"
tags        = [ "golang", "programming", "cli", "regexp" ]

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
  mermaidjs = false
+++

一応予防線を張っておくけど，この記事のコードをそのまま業務に使わないように。
まんず，そがな人はおらんじゃろけど。

さて，今日は体育の日だし頭の体操ということで「指定桁数の数字列の先頭をゼロで埋める遊び」をやってみる。
たとえば指定桁が4桁で，与えられる数字列が “123” なら， “0123” に整形するということである。
以下に要求をまとめよう。

1. 桁数は1以上の数値，数字列は長さ0以上の文字列で与えられる
1. 数字列について `0-9` の数字のみ許可（`+` / `-` 等の符号は考慮しない）
1. 与えられた数字列の長さが桁数より小さい場合は先頭をゼロ “`0`” で埋めて返す（4桁： `123` → `0123`）
1. 与えられた数字列の長さが桁数より大きい場合は下の桁を桁数分返す（4桁： `12345` → `2345`）

この機能を CLI (Command Line Interface) で表現することを考える。
こんな感じでどうだろう。

```text
$ fillzero -h
Usage: fillzero -n <number of digits> [digit string]
```

このインタフェースを実装する `main()` 関数はこんな感じかな。

```go
package main

import (
    "flag"
    "fmt"
    "os"
)

//Run returns error status in proess.
func Run(args []string) error {
    //initialize flag options
    f := flag.NewFlagSet("fillzero", flag.ContinueOnError)
    f.Usage = func() {
        fmt.Fprintln(os.Stderr, "Usage: fillzero -n <number of digits> [digit string]")
    }
    n := f.Int("n", 0, "number of digits")

    //parse arguments
    if err := f.Parse(args); err != nil {
        return os.ErrInvalid
    }
    if f.NArg() > 1 {
        f.Usage()
        return errors.New("Too many arguments.")
    }

    // get digit string
    var s string
    if f.NArg() == 1 {
        s = f.Arg(0)
    }

    // Fill in digits...

    fmt.Println(s)
    return nil
}

func main() {
    if err := Run(os.Args[1:]); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    os.Exit(0)
}
```

これであとは `s = fillzero(*n, s)` みたいな関数をでっち上げればいいのだが，私は天邪鬼なので別のことを考えた。
つまり「指定桁数の数字列の先頭をゼロで埋める」機能を持つ value object を作ったほうが面白いよね。

たとえばこんな感じ。

```go
//Identity is digit string class
type Identity struct {
    ds string // digit string
    nd int    // number of digits
}
```

これに対してインスタンスを生成する `New()` 関数と，整形済みの文字列を吐き出す [Stringer] を組み込んであげればいい。
最終的にはこんな感じになる。

```go
package main

import (
    "errors"
    "flag"
    "fmt"
    "os"
    "regexp"
    "strings"
)

//Identity is digit string class
type Identity struct {
    ds string // digit string
    nd int    // number of digits
}

//IdentityNew returns Identity instance.
func IdentityNew(n int, s string) (*Identity, error) {
    id := &Identity{}
    if n <= 0 {
        return id, errors.New("Number of digits is greater than zero.")
    }
    if len(s) > 0 {
        r, err := regexp.Compile(`^[0-9]+$`)
        if err != nil {
            return id, err
        }
        if !r.MatchString(s) {
            return id, errors.New(s + " is not digit string.")
        }
    }
    id.nd = n
    id.ds = s
    return id, nil
}

//String is Stringer method
func (id *Identity) String() string {
    if id.nd <= 0 {
        return ""
    }
    l := len(id.ds)
    if id.nd == l {
        return id.ds
    } else if id.nd < l {
        return id.ds[l-id.nd:]
    }
    return strings.Repeat("0", id.nd-l) + id.ds
}

//Run returns error status in proess.
func Run(args []string) error {
    //initialize flag options
    f := flag.NewFlagSet("fillzero", flag.ContinueOnError)
    f.Usage = func() {
        fmt.Fprintln(os.Stderr, "Usage: fillzero -n <number of digits> [digit string]")
    }
    n := f.Int("n", 0, "number of digits")

    //parse arguments
    if err := f.Parse(args); err != nil {
        return os.ErrInvalid
    }
    if f.NArg() > 1 {
        f.Usage()
        return errors.New("Too many arguments.")
    }

    // get digit string
    var s string
    if f.NArg() == 1 {
        s = f.Arg(0)
    }

    d, err := IdentityNew(*n, s)
    if err != nil {
        return err
    }
    fmt.Println(d)

    return nil
}

func main() {
    if err := Run(os.Args[1:]); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    os.Exit(0)
}
```

本来なら別パッケージにして `identity.Identity` みたいにするのがいいんだろうけど（テストも書きやすいし），ファイルを分けたり面倒なことになるので端折っている。
また，文字列のチェックに正規表現を扱える [`regexp`] パッケージを使ったが，正直に言ってこの程度の処理に正規表現パッケージを使うのが効率的かどうかは分からない（コードはすっきりするけど）。

実際に動かしてみると，こんな感じになる。

```text
$ fillzero -n 4
0000
$ fillzero -n 4 123
0123
$ fillzero -n 4 1234
1234
$ fillzero -n 4 12345
2345
```

要求は満たしているな。
よしよし。

値と機能を value object にパッケージングすると，別の実装も考えることができる。
たとえば数字列を文字列のまま保持するのではなく数値で保持するとか。

```go
//Identity is digit string class
type Identity struct {
    id uint64 // identity
    nd int    // number of digits
}

//IdentityNew returns Identity instance.
func IdentityNew(n int, s string) (*Identity, error) {
    id := &Identity{}
    if n <= 0 {
        return id, errors.New("Number of digits is greater than zero.")
    }
    if len(s) > 0 {
        i, err := strconv.ParseUint(s, 10, 64)
        if err != nil {
            return id, err
        }
        id.id = i
    }
    id.nd = n
    return id, nil
}

//String is Stringer method
func (id *Identity) String() string {
    if id.nd <= 0 {
        return ""
    }
    s := strconv.FormatUint(id.id, 10)
    l := len(s)
    if id.nd == l {
        return s
    } else if id.nd < l {
        return s[l-id.nd:]
    }
    return strings.Repeat("0", id.nd-l) + s
}
```

これにまるっと入れ替えても概ね同じ結果になる[^p1]。
数字列に対して何らかの数値演算を行う機能が必要ならこちらのほうがいいかもしれない。
ちなみに `uint64` を使ったことに特に意味はなくて，ぶっちゃけ「大は小を兼ねる」ってやつだが，このサイズで足りなければ [`big`].`Int` を使う必要がある。

[^p1]: 厳密に言うと数字列をいきなり `strconv.ParseUint()` 関数にかけているので数値として表現可能なものは通ってしまう。たとえば `+` / `-` の符号などだ。そういう意味では要求からの逸脱（deviation）があるのだが，所詮お遊びなので目をつむって欲しい。実際には数字列のチェック処理は別関数にした方がいいと思う。

って，ここまで考えたけど，やっぱ具体的な業務の中で考えないと意味ないよね（笑）
まぁ，今回は 文字列⇔数値 変換と正規表現（のさわり）について書けたからよしとするか。

## ブックマーク

- [regexpとの付き合い方 〜 Go言語標準の正規表現ライブラリのパフォーマンスとアルゴリズム〜 | eureka tech blog](https://developers.eure.jp/tech/golang-regexp/)
- [Go言語でのString・Int間の変換速度について - Qiita](https://qiita.com/evalphobia/items/ab6aefdaa576217ef8fa)
- [golang 数値（文字列）を指定桁数の先頭0詰めにして返す - Qiita](https://qiita.com/dolpher/items/021583168b7c13c66536)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[Stringer]: https://tour.golang.org/methods/17 "Stringers - A Tour of Go"
[`regexp`]: https://golang.org/pkg/regexp/ "regexp - The Go Programming Language"
[`big`]: https://golang.org/pkg/math/big/ "big - The Go Programming Language"
