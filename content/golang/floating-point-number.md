+++
title = "Go 言語で浮動小数点数をいろいろな書式で表示してみる"
date =  "2019-09-28T01:27:23+09:00"
description = "Go 1.13 で浮動小数点数の内部構造が簡単に見れるようになって，より理解が進むというものである。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "programming", "floating-point", "json", "format" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今日も面白い記事を見つけた。

- [go で float32 を JSON にするとちょっと意外なことになる - Qiita](https://qiita.com/Nabetani/items/6cea56acb71f111aebc8)

今回は [Go] 1.13 になった記念として 浮動小数点数をいろいろな書式で表示してみることにする。

## その前に...

上の記事について補足をしておくが[^fp1]， `0xa0000000` をいわゆる倍精度浮動小数点数（[Go 言語]で言うところの `float64`）にキャストした場合と，いったん単精度浮動小数点数（[Go 言語]で言うところの `float32`）にキャストした後に倍精度で再度キャストした場合で結果は同じ値になるが，これは倍精度へのキャストで精度が復元しているわけではない点に注意が必要である。

[^fp1]: [件の記事]の著者は浮動小数点数の性質については分かった上で [Go 言語]以外の言語との表現の違いについて問題にしておられるようだ。

このことは `0xa0000000` を2進数に展開するとよく分かる。
[Go] 1.13 から `%b` で整数値を2進数で表現できるようになったので試してみよう。

```go
fmt.Printf("%b\n", int64(0xa0000000))
//Output
//10100000000000000000000000000000
```

このとおり `0xa0000000` の実質的な有効桁数は3ビットしかない。
なので，単精度と倍精度の間でキャストを繰り返しても実質的な精度に影響はないのである。

もし有効桁数の違いを見たいのであればサンプルの値を `0xa0000001` に変えてみるといいだろう。
たとえば[件の記事]に出てくる JavaScript コード

```js
o = {
  foo: new Float32Array([0xa0000000]),
  bar: new Float64Array([0xa0000000])
}
process.stdout.write(JSON.stringify(o));
```

を

{{< highlight js "hl_lines=2-3" >}}
o = {
  foo: new Float32Array([0xa0000001]),
  bar: new Float64Array([0xa0000001])
}
process.stdout.write(JSON.stringify(o));
{{< /highlight >}}

に書き換えて実行してみると

```text
$ node json.js 
{"foo":{"0":2684354560},"bar":{"0":2684354561}}
```

という感じに異なる値で表示される。
これで JavaScript コードでも明示的に単精度浮動小数点数にすれば，ちゃんと精度が削られること（情報落ち）が分かるだろう。

JSON ([RFC 7159]) における数値（numbers）の内部表現は倍精度浮動小数点数のみだが JavaScript の仕様に合わせて整数で表現可能な場合はできるだけ整数で表そうとする。
故に無理やり単精度浮動小数点数に押し込めばこんな妙ちきりんなことも起きる。
幸い JSON は `1.23E+4` みたいな表現も許容するので，有効桁数を明示したいならこの表記がいいだろう。

ちなみに [Go 言語]で `float32` を JSON 形式に展開する際に `1.23E+4` のような表記にしたければ `float32` を wrap する型を作って

```go
type Float32 float32

func (f Float32) MarshalJSON() ([]byte, error) {
    return []byte(fmt.Sprintf("%v", f)), nil
}
```

などとすればよい。

## 0xa0000001 を様々な書式で表記してみる

ようやく本題（笑）

あらかじめ `0xa0000001` を

```go
const value = 0xa0000001

var (
    i64 = int64(value)
    f64 = float64(value)
    f32 = float32(value)
)
```

と定義して [`fmt`]`.Printf()` 関数の各書式（verb）を使って表記してみよう。
こんな感じ。

| Format                           | Expression                         |
| -------------------------------- | ---------------------------------- |
| `fmt.Printf("%d", i64)`          | `2684354561`                       |
| `fmt.Printf("%b", i64)`          | `10100000000000000000000000000001` |
| `fmt.Printf("%f", f64)`          | `2684354561.000000`                |
| `fmt.Printf("%f", f32)`          | `2684354560.000000`                |
| `fmt.Printf("%e", f64)`          | `2.684355e+09`                     |
| `fmt.Printf("%e", f32)`          | `2.684355e+09`                     |
| `fmt.Printf("%g", f64)`          | `2.684354561e+09`                  |
| `fmt.Printf("%g", f32)`          | `2.6843546e+09`                    |
| `fmt.Printf("%g", float64(f32))` | `2.68435456e+09`                   |

更に [Go] 1.13 から浮動小数点数に対して `%b` が使えるようになった。
 `%b` を使うと仮数部と指数部を分離した表記になる。
こんな感じ。

| Format                           | Expression             |
| -------------------------------- | ---------------------- |
| `fmt.Printf("%b", f64)`          | `5629499536310272p-21` |
| `fmt.Printf("%b", f32)`          | `10485760p+8`          |
| `fmt.Printf("%b", float64(f32))` | `5629499534213120p-21` |

仮数部が10進数表記で分かりにくいので，強引に2進数に置き換えるとこんな感じ。

| Value          | Expression                                                  |
| -------------- | ----------------------------------------------------------- |
| `f64`          | `10100000000000000000000000000001000000000000000000000p-21` |
| `f32`          | `101000000000000000000000p+8`                               |
| `float64(f32)` | `10100000000000000000000000000000000000000000000000000p-21` |

`float32` で8ビット分の情報落ちが発生していることがよく分かる。
覆水盆に返らず（笑）

[Go] 1.13 で浮動小数点数の内部構造が簡単に見れるようになって，浮動小数点数対する理解がより進むというものである。

## 【おまけの追記】 encoding/json パッケージにおける浮動小数点数の扱いと json.Number 型

[Go 言語]の標準パッケージである [encoding/json] で構造体の要素に `float32`/`float64` を割り当てた際の JSON へのエンコードでは，最終的に [`strconv`]`.AppendFloat()` 関数で文字列に変換される。

[`strconv`]`.AppendFloat()` 関数とほぼ同じ機能を持つ [`strconv`]`.FormatFloat()` 関数で出力を確認してみよう。

| Conversion                                       | Output       |
| ------------------------------------------------ | ------------ |
| `strconv.FormatFloat(f64, 'f', -1, 64)`          | `2684354561` |
| `strconv.FormatFloat(f64, 'f', -1, 32)`          | `2684354600` |
| `strconv.FormatFloat(float64(f32), 'f', -1, 64)` | `2684354560` |
| `strconv.FormatFloat(float64(f32), 'f', -1, 32)` | `2684354600` |

[encoding/json] では1番目および4番目のフォーマットを使っている。
理由は不明だが間違いではない。
単精度浮動小数点数の有効桁数は10進数換算で7桁程度しかないため `2684354560` と `2684354600` の違いに意味はないのだ。

それでも JavaScript に近い表記が欲しいのであれば，たとえば先程の `float32` の wrapper として定義した `Float32` 型を以下のように書き直すことで実現できる。

{{< highlight go "hl_lines=4" >}}
type Float32 float32

func (f Float32) MarshalJSON() ([]byte, error) {
    return []byte(strconv.FormatFloat(float64(f), 'f', -1, 64)), nil
}
{{< /highlight >}}

ところで，ちょっと反則的（？）かもしれないが [encoding/json] パッケージには [`json`]`.Number` という型が用意されている。
[`json`]`.Number` 型は名前に反して `string` 型の wrapper になっている。

```go
// A Number represents a JSON number literal.
type Number string

// String returns the literal text of the number.
func (n Number) String() string { return string(n) }

// Float64 returns the number as a float64.
func (n Number) Float64() (float64, error) {
	return strconv.ParseFloat(string(n), 64)
}

// Int64 returns the number as an int64.
func (n Number) Int64() (int64, error) {
	return strconv.ParseInt(string(n), 10, 64)
}
```

つまり JSON との `Marshal`/`Unmarshal` は文字列と同じ扱いだが，必要に応じて数値（`int64` または `float64` 型）に変換できるというわけだ。

これまで述べたように JSON の numbers を浮動小数点数に変換すると計算誤差が発生するため破壊的な変換になりがちだが [`json`]`.Number` 型であればそれを最小限に抑えられるだろう。

## ブックマーク

- [1を1億回足して1億にならない場合]({{< relref "./loop-counter.md" >}})
- [書式 %v のカスタマイズ]({{< relref "./formatting.md" >}})

[件の記事]: https://qiita.com/Nabetani/items/6cea56acb71f111aebc8 "go で float32 を JSON にするとちょっと意外なことになる - Qiita"
[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[RFC 7159]: https://tools.ietf.org/html/rfc7159 "RFC 7159 - The JavaScript Object Notation (JSON) Data Interchange Format"
[encoding/json]: https://golang.org/pkg/encoding/json/ "json - The Go Programming Language"
[`json`]: https://golang.org/pkg/encoding/json/ "json - The Go Programming Language"
[`fmt`]: https://golang.org/pkg/encoding/json/ "json - The Go Programming Language"
[`strconv`]: https://golang.org/pkg/strconv/ "strconv - The Go Programming Language"

## 参考図書

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
