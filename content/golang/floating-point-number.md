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

上の記事について補足をしておくが `0xa0000000` をいわゆる倍精度浮動小数点数（[Go 言語]で言うところの `float64`）にキャストした場合と，いったん単精度浮動小数点数（[Go 言語]で言うところの `float32`）にキャストした後に倍精度で再度キャストした場合で結果は同じ値になるが，これは倍精度へのキャストで精度が復元しているわけではない点に注意が必要である。
これは `0xa0000000` を2進数に展開するとよく分かる。

```go
fmt.Printf("%b\n", int64(0xa0000000))
//Output
//10100000000000000000000000000000
```

つまり `0xa0000000` の実質的な有効桁数は3ビットしかない。
なので単精度と倍精度の間でキャストを繰り返しても実質的な精度に影響はないのである。

この話はサンプルの値を `0xa0000000` から `0xa0000001` に変えてみればよく分かる。
たとえば件の記事の JavaScript コード

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

という具合に値が異なっている。
これで JavaScript コードでも明示的に単精度浮動小数点数にすれば精度が削られること（情報落ち）が分かるだろう。

JSON ([RFC 7159]) における数値（numbers）の内部表現は倍精度浮動小数点数のみだが JavaScript の仕様に合わせて整数で表現可能な場合はできるだけ整数で表そうとする。
故にこんな妙ちきりんなことも起きる。
幸い JSON は `1.23E+4` みたいな表現も許容するので，有効桁数を明示したいならこの表記がいいだろう。

ちなみに [Go 言語]で `float32` を JSON 形式に展開する際に `1.23E+4` のような表記にしたければ `float32` をラップする型を作って

```go
type Float32 float32

func (f Float32) MarshalJSON() ([]byte, error) {
    return []byte(fmt.Sprintf("%v", f)), nil
}
```

などとすればよい。

## 0xa0000001 を様々な書式で表記してみる

では本題。

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
 `%b` を使うと仮数部と指数部が分けて表記される。
こんな感じ。

| Format                           | Expression             |
| -------------------------------- | ---------------------- |
| `fmt.Printf("%b", f64)`          | `5629499536310272p-21` |
| `fmt.Printf("%b", f32)`          | `10485760p+8`          |
| `fmt.Printf("%b", float64(f32))` | `5629499534213120p-21` |

仮数部が10進数表記で分かりにくいので強引に2進数に置き換えるとこんな感じ。

| Value          | Expression                                                    |
| -------------- | ------------------------------------------------------------- |
| `f64`          | `0b10100000000000000000000000000001000000000000000000000p-21` |
| `f32`          | `0b101000000000000000000000p+8`                               |
| `float64(f32)` | `0b10100000000000000000000000000000000000000000000000000p-21` |

というわけで `float32` にキャストした瞬間に見事に末尾の1が削れて情報落ちが発生していることが分かるだろう。
覆水盆に返らず（笑）

[Go] 1.13 で浮動小数点数の内部構造が簡単に見れるようになって，より理解が進むというものである。

## 【おまけの追記】 encoding/json パッケージにおける浮動小数点数の扱いと json.Number 型

[Go 言語]の標準パッケージである [encoding/json] で構造体の要素に `float64` を割り当てた際の JSON へのエンコードでは，最終的に [`strconv`]`.AppendFloat()` 関数で文字列に変換される。

[`strconv`]`.AppendFloat()` 関数とほぼ同じ機能を持つ [`strconv`]`.FormatFloat()` 関数で出力を確認してみよう。

| Conversion                                       | Output       |
| ------------------------------------------------ | ------------ |
| `strconv.FormatFloat(f64, 'f', -1, 64)`          | `2684354561` |
| `strconv.FormatFloat(f64, 'f', -1, 32)`          | `2684354600` |
| `strconv.FormatFloat(float64(f32), 'f', -1, 64)` | `2684354560` |
| `strconv.FormatFloat(float64(f32), 'f', -1, 32)` | `2684354600` |

これを使えば JavaScript に近い表現になるだろう。

例えば，先程の `float32` のラッパとして定義した `Float32` 型は以下のように書き直せる。

{{< highlight go "hl_lines=4" >}}
type Float32 float32

func (f Float32) MarshalJSON() ([]byte, error) {
    return []byte(strconv.FormatFloat(float64(f), 'f', -1, 64)), nil
}
{{< /highlight >}}

ところで，ちょっと反則的（？）かもしれないが [encoding/json] パッケージには [`json`]`.Number` という型が用意されている。
[`json`]`.Number` 型は名前に反して `string` 型のラッパになっている。

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

つまり JSON としては文字列と同じ扱いだが，必要に応じて数値（`int64` または `float64` 型）に変換できるというわけだ。

これまで述べたように JSON の number を浮動小数点数に変換すると計算誤差が発生するため破壊的な変換になりがちだが [`json`]`.Number` 型であれば最小限に抑えられるだろう。

## ブックマーク

- [1を1億回足して1億にならない場合]({{< relref "./loop-counter.md" >}})
- [書式 %v のカスタマイズ]({{< relref "./formatting.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[RFC 7159]: https://tools.ietf.org/html/rfc7159 "RFC 7159 - The JavaScript Object Notation (JSON) Data Interchange Format"
[encoding/json]: https://golang.org/pkg/encoding/json/ "json - The Go Programming Language"
[`json`]: https://golang.org/pkg/encoding/json/ "json - The Go Programming Language"
[`fmt`]: https://golang.org/pkg/encoding/json/ "json - The Go Programming Language"
[`strconv`]: https://golang.org/pkg/strconv/ "strconv - The Go Programming Language"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253"><img src="https://images-fe.ssl-images-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan, Brian W. Kernighan</dd>
    <dd>柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-API</a>)</p>
</div>
