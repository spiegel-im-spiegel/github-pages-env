+++
title = "strings.Title 関数は非推奨になった"
date =  "2022-07-24T21:24:01+09:00"
description = "lint は用法・用量を守って正しく使いましょう"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "lint" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

たとえば，以下のコードがあるとする。

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    msg := "hello, world!"
    fmt.Println(msg, "->", strings.Title(msg))
}
```

これを実行すると

```text
$ go run sample.1.go
hello, world! -> Hello, World!
```

と単語の先頭が大文字に残りは小文字に変換される。

この [`strings`]`.Title()` 関数について [Go] 1.18 から

{{< fig-quote type="markdown" title="Go 1.18 Release Notes - The Go Programming Language" link="https://go.dev/doc/go1.18" lang="en" >}}
The `Title` function is now deprecated. It doesn't handle Unicode punctuation and language-specific capitalization rules, and is superseded by the [`golang.org/x/text/cases`](https://golang.org/x/text/cases) package.
{{< /fig-quote >}}

という感じに非推奨になった。
なお，このコードを [golangci-lint] にかけると

```text
$ golangci-lint run
sample1.go:10:25: SA1019: strings.Title has been deprecated since Go 1.18 and an alternative has been available since Go 1.0: The rule Title uses for word boundaries does not handle Unicode punctuation properly. Use golang.org/x/text/cases instead. (staticcheck)
    fmt.Println(msg, "->", strings.Title(msg))
                           ^
```

と言う感じに警告を出してくれる。
つか，この警告で気がついたんだけどね。
相変わらず lint に叱られっぱなしである `orz`

ちなみに似た関数名に [`strings`]`.ToTitle()` というのがあるが，これを使うと

```text
$ go run sample.1b.go
hello, world! -> HELLO, WORLD!
```

という感じに全部大文字になる。
これって [`strings`]`.ToUpper()` と何が違うんか分からん（笑）

さて， [`strings`]`.Title()` 関数を使う代わりに [`golang.org/x/text/cases`] を使えとあるようなので，早速コードを書き換えてみる。
こんな感じかな。

```go
package main

import (
    "fmt"

    "golang.org/x/text/cases"
    "golang.org/x/text/language"
)

func main() {
    msg := "hello, world!"
    fmt.Println(msg, "->", cases.Title(language.Und).String(msg))
}
```

これを実行すると

```text
$ go run sample2.go
hello, world! -> Hello, World!
```

と同じ結果が得られた。

[`language`][`golang.org/x/text/language`]`.Und` の部分は特定の言語（[`language`][`golang.org/x/text/language`]`.Japanese` とか）を指定できるのだが，たとえば役物が絡む場合：

```go
package main

import (
    "fmt"

    "golang.org/x/text/cases"
    "golang.org/x/text/language"
)

func main() {
    msg := "'n"
    tags := []language.Tag{
        language.English,
        language.Dutch,
    }
    for _, tag := range tags {
        fmt.Println(tag, ":", msg, "->", cases.Title(tag).String(msg))
    }
}
```

これを実行すると

```text
$ go run sample3.go
en : 'n -> 'N
nl : 'n -> 'n
```

という感じに言語によって違いが出るようだ？

## Lint は用法・用量を守って正しく使いましょう

話は変わるが，「失敗」には大きく2つある。
「受動的失敗」と「能動的失敗」だ。
このフレーズは Bruce Schneier さんの『[セキュリティはなぜやぶられたのか（Beyond Fear）](https://www.amazon.co.jp/dp/4822283100?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』にセキュリティ用語として出てくる。

{{< fig-quote type="markdown" title="『セキュリティはなぜやぶられたのか』p.77" link="https://www.amazon.co.jp/dp/4822283100?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
セキュリティシステムの問題は大きく分けてふたつのパターンがある。
ひとつ目は攻撃に対する防御が失敗するもの。
（中略）
とるべき対策が実行されない受動的な失敗である。
ふたつ目は間違ったときに対策を実行して防御が失敗するもの。
（中略）
とるべき対策を実行したがゆえの能動的な失敗である。
{{< /fig-quote >}}

受動的失敗は批判の対象になりやすいが，受動的失敗を恐れるあまり能動的失敗に陥るというのもありがちなパターンである。

いや Twitter の TL で [golangci-lint] をフルオプションで起動して山ほど警告が出てきても平気みたいな記述を見かけたので。
これって典型的な能動的失敗だよなぁ。
オオカミが来た！

もとより lint のような静的コード解析は万能ではない。
なれば，より気にすべきなのは受動的失敗より能動的失敗である。

[golangci-lint] はオプション無しでもかなりの部分を網羅できる。
せいぜい gosec を追加するくらい。
まずはこれで**警告が出ない**ことを目指した上で，プロダクトによって過不足があれば少しずつ調整していけばいいのだ。

たとえば，拙作の [gnkf] は MD5 や SHA-1 のハッシュ値を出力する機能があるが gosec を含めて lint をかけると「弱いハッシュ関数を使うな（←超意訳）」みたいな警告が出てしまう。
そこで

```text
$ golangci-lint run --enable gosec --exclude "G501|G505" ./...
```

のようにチェック対象を調整している。
仕事で使うならメンバ間で `.golangci.yaml` 等を使って設定を合わせておけばいいだろう。

lint は用法・用量を守って正しく使いましょう。

## ブックマーク

- [逆引き Goによる静的解析](https://zenn.dev/tenntenn/books/d168faebb1a739)

[Go]: https://go.dev/
[golangci-lint]: https://golangci-lint.run/
[`strings`]: https://pkg.go.dev/strings "strings package - strings - Go Packages"
[`golang.org/x/text/cases`]: https://pkg.go.dev/golang.org/x/text/cases "cases package - golang.org/x/text/cases - Go Packages"
[`golang.org/x/text/language`]: https://pkg.go.dev/golang.org/x/text/language "language package - golang.org/x/text/language - Go Packages"
[gnkf]: https://github.com/goark/gnkf "goark/gnkf: Network Kanji Filter by Golang"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
