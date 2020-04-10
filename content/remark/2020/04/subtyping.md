+++
title = "それは Duck Typing ぢゃない（らしい）"
date =  "2020-04-10T19:37:23+09:00"
description = "今回は Go と Rust との比較をちょっとポエミーに語ってみる（笑）"
image = "/images/attention/kitten.jpg"
tags = [ "programming", "rust", "golang", "type", "object-oriented" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今回は [Go] と [Rust] との比較をちょっとポエミーに語ってみる（笑）

そもそも [duck typing] は Ruby のような動的型付け言語における型推論の手法（のひとつ）である。
その由来は [duck test] から来ていて

{{< fig-quote type="markdown" title="Duck test - Wikipedia" link="https://en.wikipedia.org/wiki/Duck_test" lang="en" >}}
{{< quote >}}If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck.{{< /quote >}}
{{< /fig-quote >}}

というフレーズに集約されている。

静的型付け言語である [Go] や [Rust] における抽象型を使った型推論を [duck typing] と呼ぶのは厳密には正しくない，らしい。
[Go] や [Rust] における `interface` や `trait` といった抽象型を用いた型推論は「部分型付け（subtyping）」と呼ばれる。
ただし [Go] と [Rust] では全く異なる戦略をとる。

## Cat コマンドもどき（[Go] 版）

ここで簡単なプログラムを書いてみよう。
UNIX 系のプラットフォームではおなじみの `cat` コマンドの「もどき」を書いてみる。

本来の `cat` コマンドは複数の入力を結合（con**cat**enate）して出力するものだが，真面目な実装をし始めるとキリがないので，今回は以下の2つの機能のみ実装する。

1. コマンドライン引数で指定したファイルを1つのみ標準出力に出力する
2. ファイルの指定がない場合は標準入力をそのまま標準出力に出力する

ぶっちゃけ，ただの「土管」である（笑） これを [Go] で書いたのが以下のコードだ。

```go
package main

import (
    "fmt"
    "io"
    "os"
)

func concatenate(w io.Writer, r io.Reader) error {
    _, err := io.Copy(w, r)
    return err
}

func main() {
    if len(os.Args) > 1 {
        file, err := os.Open(os.Args[1])
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            return
        }
        defer file.Close()

        if err := concatenate(os.Stdout, file); err != nil {
            fmt.Fprintln(os.Stderr, err)
            return
        }
    } else {
        if err := concatenate(os.Stdout, os.Stdin); err != nil {
            fmt.Fprintln(os.Stderr, err)
            return
        }
    }
}
```

`concatenate()` 関数がメインのロジックで，引数の [`io`]`.Writer`, [`io`]`.Reader` および返り値の `error` は全て `interface` 型である。
まぁ `concatenate()` 関数を括り出す必然性は全くないのだが，後述の [Rust] のコードと比較しやすいよう敢えて分けている。

`concatenate()` 関数の呼び出しで，最初の

```go
if err := concatenate(os.Stdout, file); err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
}
```

と次の

```go
if err := concatenate(os.Stdout, os.Stdin); err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
}
```

は（当然ながら）同じ関数で，引数や返り値にどのようなインスタンスが入るかは実行時に決まる。
コンパイル時に決まるのは注入するインスタンスの構造が受け入れる `interface` 型の構造と合致していることだけだ（合致しなければコンパイル・エラー）。

すンごい簡単に書かれているけど，これは「依存の注入（depencency injection）」の典型例であり「[Go] では [duck typing] ができる」とか言われる所以である。

では，これをリファレンスとして，今度は [Rust] を使って書いてみる。

## Cat コマンドもどき（[Rust] 版，総称型編）

とりあえず，えいやっで書いたコードがこちら。

```rust
fn concatenate<W, R>(w: &mut W, r: &mut R) -> Result<(), std::io::Error>
where
    W: std::io::Write,
    R: std::io::Read,
{
    let mut buf = Vec::new();
    r.read_to_end(&mut buf)?;
    w.write_all(&buf)?;
    Ok(())
}

fn main() -> Result<(), std::io::Error> {
    let args = std::env::args();
    if args.len() > 1 {
        for s in args.skip(1).take(1) {
            concatenate(
                &mut std::io::stdout(),
                &mut std::io::BufReader::new(std::fs::File::open(s)?),
            )?;
        }
    } else {
        concatenate(&mut std::io::stdout(), &mut std::io::stdin())?;
    }
    Ok(())
}
```

[`std::io::Write`] と [`std::io::Read`] が `trait` 型なのだが，各 `trait` は総称型 `W`, `R` の制約条件として書かれているだけで実行時に機能するわけではない。
つまり最初の

```rust
concatenate(
    &mut std::io::stdout(),
    &mut std::io::BufReader::new(std::fs::File::open(s)?),
)?;
```

と次の

```rust
concatenate(&mut std::io::stdout(), &mut std::io::stdin())?;
```

はコンパイル時に別の関数として展開される[^drop1]。
これを（多態化（polymorphization）に対する）単態化（monomorphization）と呼ぶ。

[^drop1]: 余談だが [Rust] では「ファイルを閉じる」操作は変数の生存期間満了時に暗黙的に行われるようだ。明示的に閉じるには `drop` 関数を使う。

じゃあ [Rust] では依存の注入は書けないのかというと，勿論そんなことはない。

## Cat コマンドもどき（[Rust] 版，依存注入編）

依存の注入ができるように書き換えたバージョンがこれ。

```rust
fn concatenate(
    w: &mut Box<dyn std::io::Write>,
    r: &mut Box<dyn std::io::Read>,
) -> Result<(), Box<dyn std::error::Error>> {
    let mut buf = Vec::new();
    r.read_to_end(&mut buf)?;
    w.write_all(&buf)?;
    Ok(())
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let args = std::env::args();
    let mut r: Box<dyn std::io::Read> = if args.len() > 1 {
        let fnam = match args.skip(1).next() {
            Some(s) => s,
            _ => "".to_string(),
        };
        Box::new(std::io::BufReader::new(std::fs::File::open(fnam)?))
    } else {
        Box::new(std::io::stdin())
    };
    let mut w: Box<dyn std::io::Write> = Box::new(std::io::stdout());
    concatenate(&mut w, &mut r)?;
    Ok(())
}
```

`concatenate()` 関数が同一のものであることを強調するために呼び出しをひとつに纏めているので少しまだるこしい書き方になっているが，ご容赦。
このように [Rust] では `trait` 型を [`Box`]`<dyn Trait>` の形式に落とし込むことで実行時の動的ディスパッチを可能にしている。

## Accept Interfaces, Return Structs

[Go] の設計指針で有名な言葉に {{< quote lang="en" >}}accept interfaces, return structs{{< /quote >}} というのがある。
私自身は必ずしもこれに賛同しないが（システム内部のコンテキスト境界は `interface` にすべき），この指針は [Go] の特徴をよく表している。

たとえば [`io`]`.Reader` と [`os`]`.File` は同じ `Read()` 関数を持つという点で関連しているけど，両者の間に明示された記述は存在しない。
それでも，その関係を以って [`io`]`.Reader` に [`os`]`.File` インスタンスを注入可能である。
[Go] プログラマは息をするように依存を注入するのだ。

このような関係を構造型の部分型付け（structural subtyping）と呼ぶそうな。


## 構造型と公称型

[Go] の `interface` 型が構造型の部分型付けであるのに対し [Rust] の `trait` 型は公称型の部分型付け（nominal subtyping）に分類されるだろう。
たとえば [`std::io::Read`] と [`std::fs::File`] との間にはコード上で明示された関係がある。
その「明示された関係」がなければ，たとえ同じ構造を持っていたとしても，両者の間に関係があるとは見なされないのだ。

[Rust] の言語仕様がこのような制約を構成しているのには，勿論ちゃんとした理由がある。

[Go] においてはメモリ管理や並列処理[^ccrcy1] をランタイム・モジュールに「丸投げ」している。
なので，プログラマは富豪的な記述に専念できるが，バイナリは肥大化してしまうしコンパイル時の最適化にも限度がある[^tg1]。

[^ccrcy1]: [Go] における並行処理と並列処理の違いについては[『Go言語による並行処理』を読む]({{< ref "/remark/2018/11/concurrency-in-go.md" >}} "『Go 言語による並行処理』は Go 言語プログラマ必読書だろう")ことを強くおすすめする。
[^tg1]: 近年，特に組込み用途で注目されている [TinyGo] は [LLVM] 上で動作することを前提としていて，本家 [Go] に比べてかなり小さい実行バイナリを吐けるらしい。

[Rust] はリソース管理等についてプログラマ側でかなり面倒を見なければならないが（それでも C/C++ などに比べれば全然楽だし安全），言い換えればコード上でのコントロールがし易くコンパイル時の最適化についてもかなり期待できる。
上述の cat コマンドもどきでも，コンパイル時の単態化を避けるコードをわざわざ書く理由はないだろう。

これはプログラム設計時の重要なトレードオフとなる。
まぁ「[Go] か [Rust] か」みたいな究極の選択をする状況はないと思うが，複数のプログラミング言語からどれかを選ぶ際にはこういったことも考慮していくべきだ（選ぶ余裕もない事案のほうが多いだろうけど`w`）。

[前にも書いた]({{< ref "/remark/2020/03/currying.md" >}})が，「それができる」ことと「そのように作られている」ことには天と地ほどの違いがある。
どうせ「書く」なら無茶せず楽しく書きたいものである。

## ブックマーク

- [Rustでファイルの入出力 - Qiita](https://qiita.com/fujitayy/items/12a80560a356607da637)
- [RustのファイルI/OにはBufReader, BufWriterを使いましょう、という話 - Qiita](https://qiita.com/gyu-don/items/50f4239fc856bed00dc4)
- [Go言語のInterfaceの考え方、Accept interfaces,return structs - Qiita](https://qiita.com/weloan/items/de3b1bcabd329ec61709)

- [継承できないなら注入すればいいじゃない！](https://slide.baldanders.info/shimane-go-2020-01-23/) : [Go] のイベント用に作ったスライド

[Rust]: https://www.rust-lang.org/ "Rust Programming Language"
[Go]: https://golang.org/ "The Go Programming Language"
[TinyGo]: https://tinygo.org/ "TinyGo - Go on Microcontrollers and WASM"
[LLVM]: https://llvm.org/ "The LLVM Compiler Infrastructure Project"
[duck typing]: https://en.wikipedia.org/wiki/Duck_typing "Duck typing - Wikipedia"
[duck test]: https://en.wikipedia.org/wiki/Duck_test "Duck test - Wikipedia"
[`std::io::Write`]: https://doc.rust-lang.org/nightly/std/io/trait.Write.html "std::io::Write - Rust"
[`std::io::Read`]: https://doc.rust-lang.org/nightly/std/io/trait.Read.html "std::io::Read - Rust"
[`std::fs::File`]: https://doc.rust-lang.org/std/fs/struct.File.html "std::fs::File - Rust"
[`Box`]: https://doc.rust-lang.org/std/boxed/struct.Box.html "std::boxed::Box - Rust"
[`io`]: https://pkg.go.dev/io "io package · go.dev"
[`os`]: https://pkg.go.dev/os "os package · go.dev"

## 参考図書

{{% review-paapi "4048930702" %}} <!-- プログラミング言語Rust 公式ガイド -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
