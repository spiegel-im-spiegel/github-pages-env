+++
title = "書式 %v のカスタマイズ"
date =  "2019-09-12T23:19:22+09:00"
description = "Stringer, GoStringer および Formatter インタフェースを使ったカスタマイズ。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "programming", "format" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今回も小ネタで。

お馴染みの [`fmt`].`Printf()` 関数などで使われる書式（verb）のうち，今回は `%v` の出力をカスタマイズすることを考えてみる。

## 基本型における %v 書式の出力

まずは `%v` の定義から

| Verb  | Description                                                                                    |
|:-----:| ---------------------------------------------------------------------------------------------- |
| `%v`  | the value in a default format<br>when printing structs, the plus flag (`%+v`) adds field names |
| `%#v` | a Go-syntax representation of the value                                                        |

更に基本型については `%v` は以下の書式と対応している。

| Type                        | Default Verb                      |
| --------------------------- | --------------------------------- |
| `bool`                      | `%t`                              |
| `int`, `int8`, ...          | `%d`                              |
| `uint`, `uint8`, ...        | `%d`, `%#x` if printed with `%#v` |
| `float32`, `complex64`, ... | `%g`                              |
| `string`                    | `%s`                              |
| `chan`                      | `%p`                              |
| pointer                     | `%p`                              |

構造体や配列などの複合オブジェクトについては以下のように展開される。

| Compound Object  | Format                             |
| ---------------- | ---------------------------------- |
| struct           | `{field0 field1 ...}`              |
| array, slice     | `[elem0 elem1 ...]`                |
| maps             | `map[key1:value1 key2:value2 ...]` |
| pointer to above | `&{}`, `&[]`, `&map[]`             |

ちょっと試し書きをしてみよう。
たとえば，以下のような構造体とデータを考えてみる。

```go
type Planet struct {
    Name string
    Mass float64
}

var planets = []Planet{
    {Name: "Mercury", Mass: 0.055},
    {Name: "Venus", Mass: 0.815},
    {Name: "Earth", Mass: 1.0},
    {Name: "Mars", Mass: 0.107},
}
```

この `planets` を `%v` を使って出力してみよう。
こんな感じ。

```go
fmt.Printf("%v", planets)
// Output:
// [{Mercury 0.055} {Venus 0.815} {Earth 1} {Mars 0.107}]
```

```go
fmt.Printf("%+v", planets)
// Output:
// [{Name:Mercury Mass:0.055} {Name:Venus Mass:0.815} {Name:Earth Mass:1} {Name:Mars Mass:0.107}]
```

```go
fmt.Printf("%#v", planets)
// Output:
// []main.Planet{main.Planet{Name:"Mercury", Mass:0.055}, main.Planet{Name:"Venus", Mass:0.815}, main.Planet{Name:"Earth", Mass:1}, main.Planet{Name:"Mars", Mass:0.107}}
```

## Stringer および GoStringer インタフェース

[`fmt`].`Stringer` および [`fmt`].`GoStringer` インタフェースを持つ型であれば `%v` の出力をカスタマイズできる。
[`fmt`].`Stringer` および [`fmt`].`GoStringer` インタフェースの定義は以下の通り。

```go
// *.go is implemented by any value that has a String method,
// which defines the ``native'' format for that value.
// The String method is used to print values passed as an operand
// to any format that accepts a string or to an unformatted printer
// such as Print.
type Stringer interface {
    String() string
}

// GoStringer is implemented by any value that has a GoString method,
// which defines the Go syntax for that value.
// The GoString method is used to print values passed as an operand
// to a %#v format.
type GoStringer interface {
    GoString() string
}
```

先ほどの `Planet` 型に [`fmt`].`Stringer` および [`fmt`].`GoStringer` インタフェースを組み込んでみよう。

```go
func (p Planet) String() string {
    return fmt.Sprintf("%s (%.3f)", p.Name, p.Mass)
}

func (p Planet) GoString() string {
    return fmt.Sprintf(`main.Planet{Name:%s, Mass:%.3f}`, strconv.Quote(p.Name), p.Mass)
}
```

これで `%v` の出力は以下のように変わる。

```go
fmt.Printf("%v", planets)
// Output:
// [Mercury (0.055) Venus (0.815) Earth (1.000) Mars (0.107)]
```

```go
fmt.Printf("%+v", planets)
// Output:
// [Mercury (0.055) Venus (0.815) Earth (1.000) Mars (0.107)]
```

```go
fmt.Printf("%#v", planets)
// Output:
// []main.Planet{main.Planet{Name:"Mercury", Mass:0.055}, main.Planet{Name:"Venus", Mass:0.815}, main.Planet{Name:"Earth", Mass:1.000}, main.Planet{Name:"Mars", Mass:0.107}}
```

`%v` および `%+v` が [`fmt`].`Stringer` に `%#v` が [`fmt`].`GoStringer` に対応しているのが分かると思う。

## Formatter インタフェース

[`fmt`].`Stringer` インタフェースを使ったカスタマイズの欠点は `%v` と `%+v` を区別できないことだ。
`%v` と `%+v` を区別できるよう詳細な操作を行いたいのであれば [`fmt`].`Formatter` インタフェースを組み込む。
[`fmt`].`Formatter` インタフェースの定義は以下の通り。

```go
// Formatter is the interface implemented by values with a custom formatter.
// The implementation of Format may call Sprint(f) or Fprint(f) etc.
// to generate its output.
type Formatter interface {
    Format(f State, c rune)
}
```

更に引数の [`fmt`].`State` もインタフェース型で以下のように定義されている。

```go
// State represents the printer state passed to custom formatters.
// It provides access to the io.Writer interface plus information about
// the flags and options for the operand's format specifier.
type State interface {
    // Write is the function to call to emit formatted output to be printed.
    Write(b []byte) (n int, err error)
    // Width returns the value of the width option and whether it has been set.
    Width() (wid int, ok bool)
    // Precision returns the value of the precision option and whether it has been set.
    Precision() (prec int, ok bool)

    // Flag reports whether the flag c, a character, has been set.
    Flag(c int) bool
}
```

つまり自作の `Format()` メソッド内では `State.Write()`,  `State.Width()`,  `State.Precision()`, `State.Flag()` 各メソッドが使える。 これらを使って出力の整形を行えるわけだ（`State.Write()` は [`io`].`Writer` インタフェースとマッチしている点にも注目）。

では `Planet` 型に [`fmt`].`Formatter` インタフェースを組み込んでみる。
こんな感じでどうだろう。

```go
func (p Planet) Format(s fmt.State, verb rune) {
    switch verb {
    case 'v':
        switch {
        case s.Flag('#'):
            io.Copy(s, strings.NewReader(p.GoString()))
        case s.Flag('+'):
            fmt.Fprintf(s, `{"Name":%s,"Mass":%.3f}`, strconv.Quote(p.Name), p.Mass)
        default:
            io.Copy(s, strings.NewReader(p.String()))
        }
    case 's':
        io.Copy(s, strings.NewReader(p.String()))
    default: //bad verb
        fmt.Fprintf(s, `%%!%c(%s)`, verb, p.GoString())
    }
}
```

これで `%v` の出力は以下のように変わる。

```go
fmt.Printf("%v", planets)
// Output:
// [Mercury (0.055) Venus (0.815) Earth (1.000) Mars (0.107)]
```

```go
fmt.Printf("%+v", planets)
// Output:
// [{"Name":"Mercury","Mass":0.055} {"Name":"Venus","Mass":0.815} {"Name":"Earth","Mass":1.000} {"Name":"Mars","Mass":0.107}]
```

```go
fmt.Printf("%#v", planets)
// Output:
// []main.Planet{main.Planet{Name:"Mercury", Mass:0.055}, main.Planet{Name:"Venus", Mass:0.815}, main.Planet{Name:"Earth", Mass:1.000}, main.Planet{Name:"Mars", Mass:0.107}}
```

[`fmt`].`Formatter` インタフェースを組み込めば細かい制御ができるようになるが，取りうる書式を全て記述しないといけないのが面倒である[^f1]。
状況によって使い分けるのがいいだろう。

[^f1]: 型名（`%T`）とポインタ値（`%p`）は [`fmt`].`Formatter` の制御外になるようだ。

## ブックマーク

- [fmt.Formatterを実装して%vや%+vをカスタマイズしたり、%3🍺みたいな書式をつくってみよう #golang - Qiita](https://qiita.com/tenntenn/items/453a09c4c6d7f580d0ab)
- [[Go] 独自型にfmtパッケージのインターフェースを実装して出力を制御する - My External Storage](https://budougumi0617.github.io/2019/10/12/confirm-print-with-fmt-interfaces/)

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`fmt`]: https://golang.org/pkg/fmt/ "fmt - The Go Programming Language"
[`io`]: https://golang.org/pkg/io/ "io - The Go Programming Language"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN), 4621300253 (ISBN), 9784621300251 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
