+++
date = "2016-02-22T19:22:33+09:00"
update = "2016-02-22T20:53:30+09:00"
description = "Interface には落とし穴がある。"
draft = false
tags = ["golang", "interface"]
title = "Interface の謎"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

今回も軽めの小ネタで。

まず，文字列の配列を表示するだけの単純なコードを書いてみる。

```go
package main

import "fmt"

func main() {
    strlist := []string{"alpha", "beta", "gamma"}
    fmt.Println(strlist)
}
```

結果は

```text
[alpha beta gamma]
```

となる。
配列[^s] の中身をそのままダンプ出力しているだけなので，まぁ当たり前っちゃあ当たり前。
では，配列のダンプではなくきちんと項目を列挙したいとしよう。
やり方は色々あるが簡単に "`...`” トークンを使って

[^s]: 厳密には [slice]。分かってますよ，もちろん。

```go
package main

import "fmt"

func main() {
    strlist := []string{"alpha", "beta", "gamma"}
    fmt.Println(strlist...)
}
```

と配列を展開すればいんじゃね？ って思うよね，普通。
[`fmt`].`Println()` 関数の定義を見ても

```go
func Println(a ...interface{}) (n int, err error)
```

となっているし，問題ないように見える。

でもこれはうまくいかない。
これを実行しようとすると

```text
prog.go:7: cannot use strlist (type []string) as type []interface {} in argument to fmt.Println
```

とエラーになる。

実は `[]string` 型の `strlist` は [`fmt`].`Println()` 関数に渡す際に `[]interface{}` 型ではなく `interface{}` 型に**必ず**キャストされる。
だから `strlist...` と展開しようとしても「そりゃあ無理（←超意訳）」と怒られてしまうわけだ。
[Go 言語]の型（[type]）は

```go
type Msg []string
```

のように配列やポインタも型として定義できてしまうことを[思い出して]({{< relref "object-oriented-programming.md" >}})欲しい。

じゃあ，明示的なキャストならいけるのかと思ったが

```go
package main

import "fmt"

type Msg []string

func main() {
    strlist := []string{"alpha", "beta", "gamma"}
    fmt.Println(([]interface{})(strlist)...)
}
```

結果は

```text
prog.go:9: cannot convert strlist (type []string) to type []interface {}
```

と，これもエラーになった。

ではどうすればいいのかというと `[]interface{}` 型の配列を用意してそこに値をコピーする。
先程のコードであれば

```go
package main

import "fmt"

type Msg []string

func main() {
    strlist := []string{"alpha", "beta", "gamma"}
    var list = make([]interface{}, 0)
    for _, str := range strlist {
        list = append(list, str)
    }
    fmt.Println(list...)
}
```

とすれば

```text
alpha beta gamma
```

のようにめでたく列挙される。
この問題は [`fmt`].`Println()` 関数だけじゃなく，ある型の配列を `[]interface{}` 型にキャストしようとする際は必ず発生するようだ[^c]。

[^c]: 例えば `list` に `strlist` の内容をコピーする際に for 文で回すのではなく `list = append(list, strlist...)` でできるかどうか試してみればいい。

いや，「“`cannot use strlist (type []string) as type []interface {} in argument to fmt.Println`” なんてコンパイルエラーを出せるならコンパイラ側でなんとかしてよ」と思うのだが，どうも無理らしい。

やれやれ。

## ブックマーク

- [Why Go is a poorly designed language — Medium](https://medium.com/@tucnak/why-go-is-a-poorly-designed-language-1cc04e5daf2#.ucutrogyz) （[日本語訳](http://postd.cc/why-go-is-a-poorly-designed-language/)）
- [InterfaceSlice · golang/go Wiki](https://github.com/golang/go/wiki/InterfaceSlice)

[Go 言語に関するブックマーク集はこちら]({{< relref "bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[slice]: http://golang.org/ref/spec#Slice_types
[`fmt`]: https://golang.org/pkg/fmt/ "fmt - The Go Programming Language"
[type]: https://golang.org/ref/spec#Properties_of_types_and_values "Properties of types and values"
