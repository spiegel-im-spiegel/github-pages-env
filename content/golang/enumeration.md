+++
title = "列挙型での遊び方"
date =  "2019-08-31T23:49:53+09:00"
description = "実際に Go 言語で列挙を記述する場合は列挙専用の型を定義するのが定石である。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "programming", "enumeration" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

どうも [Go] 1.13 は8月中には出なさそうなので，軽く小話を。

## 列挙型と定数生成器

[Go 言語]にはいわゆる「列挙型」や「列挙クラス」といったものは存在しない。
シンボルの定義には `const` を使えばいいので敢えて「列挙型」のようなものを考える必要はないわけだ。
それでも

```go
const (
    ONE   = 1
    TWO   = 2
    THREE = 3
)
```

などといちいち書いていくのは面倒である。
そこで [Go 言語]には `iota` と呼ばれる定数生成器が用意されている。

先ほどのコードであれば

{{< highlight go "hl_lines=5-9" >}}
package main

import "fmt"

const (
    ONE = iota + 1
    TWO
    THREE
)

func main() {
    fmt.Println(ONE, TWO, THREE)
    //Output:
    //1 2 3
}
{{< /highlight >}}

と記述できる。
整数値に限れば `iota` の応用範囲は結構広くて

```go
const (
    BIT0 = 1 << iota
    BIT1
    BIT2
    BIT3
)
```

てな感じにビットフラグのシンボルを定義することもできる。

## 列挙専用の型を定義する

実際に [Go 言語]で列挙を記述する場合は列挙専用の型を定義するのが定石である。
たとえば

```go
type BitFlag uint
```

という型を定義して

```go
const (
    BIT0 BitFlag = 1 << iota
    BIT1
    BIT2
    BIT3
)
```

という感じに書ける。
型にはメソッドを関連付けられるので，上の `BitFlag` 型に対して

```go
func (f BitFlag) String() string {
    return fmt.Sprintf("%#02x", uint(f))
}
```

のような Stringer を用意すると

```go
func main() {
    fmt.Println(BIT0, BIT1, BIT2, BIT3)
    //Output:
    //0x01 0x02 0x04 0x08
}
```

てな感じに使える。

## 列挙シンボルに「値」を関連付ける

ここまで説明すれば列挙シンボルに「値」を関連付けるのはそんなに難しくないと気づくだろう。

以下のような列挙シンボルを考える。

```go
type CharEncoding int

const (
    Unknown CharEncoding = iota
    UTF8
    ShiftJIS
    EUCJP
    ISO2022JP
)
```

これらの列挙シンボルに対応する文字列を定義する。
こんな感じ。

```go
var encodingNameMap = map[CharEncoding]string{
    UTF8:      "UTF-8",
    ShiftJIS:  "Shift_JIS",
    EUCJP:     "EUC-JP",
    ISO2022JP: "ISO-2022-JP",
}
```

これで Stringer を

```go
func (c CharEncoding) String() string {
    if s, ok := encodingNameMap[c]; ok {
        return s
    }
    return "Unknown"
}
```

と書けば

```go
func main() {
    fmt.Println(UTF8)
    //Output:
    //UTF-8
}
```

てな感じにできる。

あるいは文字列から列挙シンボルに変換する

```go
func GetCharEncoding(s string) CharEncoding {
    for key, value := range encodingNameMap {
        if strings.EqualFold(value, s) {
            return key
        }
    }
    return Unknown
}
```

という関数を作れば

```go
func main() {
    fmt.Println(GetCharEncoding("utf-8"))
    //Output:
    //UTF-8
}
```

といった感じで簡易バリデーションみたいなこともできるだろう。

もちろん「値」との関連付けは文字列に限るわけではなく，たとえば

```go
import (
    "golang.org/x/text/encoding"
    "golang.org/x/text/encoding/japanese"
    "golang.org/x/text/encoding/unicode"
)

var encodingMap = map[CharEncoding]encoding.Encoding{
    UTF8:      unicode.UTF8,
    ShiftJIS:  japanese.ShiftJIS,
    EUCJP:     japanese.EUCJP,
    ISO2022JP: japanese.ISO2022JP,
}

func (c CharEncoding) Encoding() encoding.Encoding {
    if e, ok := encodingMap[c]; ok {
        return e
    }
    return nil
}
```

などと幾らでも追加できる。

[Go]: https://go.dev/
[Go 言語]: https://golang.org/ "The Go Programming Language"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
