+++
title = "For-Range 構文の話"
date =  "2019-05-18T00:44:20+09:00"
description = "Go 言語においては「参照」は忘れよう。コーディングを行う際はインスタンスの値とアドレッシングに注意しながら進めるとハマりにくい。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "engineering", "pointer" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今回は小ネタでお送りします。

つか，以下の記事

- [Big Sky :: Go のポインタの躓きやすい点](https://mattn.kaoriya.net/software/lang/go/20190516095124.htm)

の後半部分が何を問題にしているのか分からずしばらく悩んでしまった。
頭が悪くてゴメンペコン。

まず前提として [Go 言語]において「参照」のことは忘れよう。
見た目は参照ぽく振る舞う場合もあるしこのブログでも比喩表現として参照という言葉をよく使うが，言語仕様としては [Go 言語]に「参照」は存在しない。
[Go 言語]のインスタンスを表すものは「値」と「アドレッシング」であり，その「アドレッシング」の表現として（C 言語ではお馴染みの）ポインタを使うわけだ[^ptr1]。
したがって，コーディングを行う際はインスタンスの値とアドレッシングを常に頭に入れながら進めるとハマりにくい。
分からないならコードを書いてみればいいのだ。

[^ptr1]: 参照とポインタの決定的な違いは，参照は「機能」だがポインタは「値」である，という点であろう。値なのでポインタ自身をインスタンスとして表現し得るし「ポンタへのポインタ」みたいな記述もできる。この辺の感覚が掴めないと苦労するかもしれない。大昔の C 言語が全盛だった頃もポインタで躓く人は割といたからなぁ。

では，最初に紹介した記事を参考に実際にコードを書いてみようか。

まず，以下のような型 `A` を考える。

```go
type A struct {
    N string
}

func (a *A) GenFunc() func() {
    return func() {
        fmt.Printf("%s : %p\n", a.N, a)
    }
}
```

`A.GenFunc()` 関数は型 `A` の内容を標準出力に出力する関数を返す。
Method receiver が `A` ではなく `*A` とポインタ型になっているのがポイントである（洒落ちゃったてへぺろ）。

これを使って以下のコードを書いてみる。

```go
package main

import "fmt"

type A struct {
    N string
}

func (a *A) GenFunc() func() {
    return func() {
        fmt.Printf("%s : %p\n", a.N, a)
    }
}

func main() {
    as := []A{
        {"foo"},
        {"bar"},
    }
    for i := 0; i < len(as); i++ {
        fmt.Printf("instance: %s : %p\n", as[i].N, &as[i])
        as[i].GenFunc()()
    }
}
```

このコードを実行すると

```text
instance: foo : 0x40a0e0
foo : 0x40a0e0
instance: bar : 0x40a0e8
bar : 0x40a0e8
```

などと出力される。
`GenFunc()` メソッドを呼び出す際に method receiver としての `as[i]` の型が暗黙的に変換されていることに注意してほしい。
ここまでは OK かな。

では for-range 構文を使うとどうなるか。
試しに `main` 関数に for-range の制御ブロックを追加して比較してみよう。

{{< highlight go "hl_lines=11-14" >}}
func main() {
    as := []A{
        {"foo"},
        {"bar"},
    }
    for i := 0; i < len(as); i++ {
        fmt.Printf("instance: %s : %p\n", as[i].N, &as[i])
        as[i].GenFunc()()
    }
    fmt.Println()
    for _ , a := range as {
        fmt.Printf("instance: %s : %p\n", a.N, &a)
        a.GenFunc()()
    }
}
{{< /highlight >}}

これを実行するとこうなる。

```text
instance: foo : 0x40a0e0
foo : 0x40a0e0
instance: bar : 0x40a0e8
bar : 0x40a0e8

instance: foo : 0x40c160
foo : 0x40c160
instance: bar : 0x40c160
bar : 0x40c160
```

ポインタ値からインスタンス `a` はインスタンス `as[i]` そのものではないことが分かる。
しかも `a` は for-range ループの中で使い回されていていることも分かる。

つまり上の for-range 構文は

```go
for i, a := 0, as[0]; i < len(as); i, a = i+1, as[i+1] {
    fmt.Printf("instance: %s : %p\n", a.N, &a)
    a.GenFunc()()
}
```

と実質的に同じなのだ[^f1]。

[^f1]: 厳密には違う。つか，このコードを実際に書いて動かしてみれば分かるが “`index out of range`” の panic を吐いて落ちる。理由は各自で考えてみよう（笑）

ここまで来れば

```go
func main() {
    as := []A{
        {"foo"},
        {"bar"},
    }
    fs := []func(){}
    for _, a := range as {
        fs = append(fs, a.GenFunc())
    }
    for _, f := range fs {
        f()
    }
}
```

の実行結果がどうなるか容易に想像がつくだろう。
以下の通りである。

```text
bar : 0x40c128
bar : 0x40c128
```

つまり `fs` に格納される関数は全て前半の for-range の中のインスタンス `a` に帰属するメソッドであり，そのインスタンス `a` の値は for-range ループの中で上書きされているのである。

じゃあどうすればいいかというと，一番簡単で場当たりな対処としては for-range ループの中でインスタンス `a` のコピーを作ればよい。

{{< highlight go "hl_lines=8-9" >}}
func main() {
    as := []A{
        {"foo"},
        {"bar"},
    }
    fs := []func(){}
    for _ , a := range as {
        aa := a //create copy instance
        fs = append(fs, aa.GenFunc())
    }
    for _ , f := range fs {
        f()
    }
}
{{< /highlight >}}


これで実行結果は

```text
foo : 0x40c128
bar : 0x40c140
```

となる。
あるいは最初から素直に

{{< highlight go "hl_lines=7-9" >}}
func main() {
    as := []A{
        {"foo"},
        {"bar"},
    }
    fs := []func(){}
    for i := 0; i < len(as); i++ {
        fs = append(fs, as[i].GenFunc())
    }
    for _ , f := range fs {
        f()
    }
}
{{< /highlight >}}

とすれば要らんコピーも発生せず，実行結果も

```text
foo : 0x40a0e0
bar : 0x40a0e8
```

となる。

最初に紹介した記事にも書かれている通り，インスタンスの値とポインタを混ぜるから危険なのであって混ぜなければ大丈夫（笑）

[Go 言語]: https://golang.org/ "The Go Programming Language"

## ブックマーク

- [関数とポインタ]({{< ref "./function-and-pointer.md" >}})

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->

{{% review-paapi "4320026926" %}} <!-- プログラミング言語C -->
