+++
title = "次期 Go 言語で導入される総称型について予習する（その3）"
date =  "2020-09-26T23:21:12+09:00"
description = "Go が総称型を得ることで更に独自に発展することを期待している。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "language", "programming", "generics", "type", "constraint" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

{{< div-box type="markdown" >}}
**【2021-02-21 追記】**
型パラメータの提案（「総称型」というのは厳密には正しくないらしい）は2021年2月に承認され，製造フェーズに入ったようだ。
{{< /div-box >}}

1. [次期 Go 言語で導入される（かもしれない）総称型について予習する]({{< relref "./generics-in-go-2.md" >}})
1. [次期 Go 言語で導入される総称型について予習する（その2）]({{< relref "./generics-in-go-2-update.md" >}})
1. [次期 Go 言語で導入される総称型について予習する（その3）]({{< relref "./generics-in-go-2-update-update.md" >}}) ←イマココ

----

[前回]紹介した “[Type Parameters - Draft Design](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md)” が 2020-09-21 にアップデートされたらしい（最終更新：2020-11-25）。
どうやらこれが最終案になりそうだ。

早ければ2021年8月にリリース予定の [Go] 1.17 で[お披露目](https://blog.golang.org/generics-next-step "The Next Step for Generics - The Go Blog")というスケジュールは変わらない模様。
また総称型（generics）に対応した Playground も最終案の仕様で稼働中である。

- [`https://go2goplay.golang.org/`](https://go2goplay.golang.org/)

## 前回からの変更点

[前回]は，たとえば

```go
type ordered interface {
	type int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		float32, float64,
		string
}

func Max(type T ordered)(x, y T) T {
    if x < y {
        return y
    }
    return x
}
```

のような記述だったが，今回は型パラメータ（type parameter）の記述がちょっと変わったようで


```go {hl_lines=[8]}
type ordered interface {
	type int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		float32, float64,
		string
}

func Max[T ordered](x, y T) T {
	if x < y {
		return y
	}
	return x
}
```

と角括弧 `[...]` で囲まれている。
どうしても山括弧 `<...>` はイヤみたい（笑）

`type` キーワードは無用となった。

## 組込み済みの制約

組込み済みの interface 型として `error` が定義されているように，組込み済みの制約（constraint）として `any` と `comparable` が追加されるようだ。

`any` はあらゆる型を包摂する。
むしろ「制約がない」ことを示す制約というべきか。
強いて書くなら

```go
type any interface{}
```

といったところだろうか。
たとえば

{{< fig-quote class="nobox" type="markdown" title="Type Parameters - Draft Design" link="https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md" lang="en" >}}
```go {hl_lines=[4]}
// Print prints the elements of any slice.
// Print has a type parameter T and has a single (non-type)
// parameter s which is a slice of that type parameter.
func Print[T any](s []T) {
    for _, v := range s {
		fmt.Println(v)
	}
}
```
{{< /fig-quote >}}

みたいな感じに使える。
どうやら型パラメータの指定で制約を省略することはできないようなので，制約がない場合は `any` を付けることになりそうだ。

`comparable` は演算子 `==` および `!==` が使える型を示す制約だ。
これを使って

{{< fig-quote class="nobox" type="markdown" title="Go Generics draft design Final" link="https://blog.yongweilun.me/go-generics-draft-design-final" lang="en" >}}
```go {hl_lines=[5]}
package main

import "fmt"

func Contains[T comparable](col []T, item T) bool {
	for _, e := range col {
		if e == item {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(Contains([]string{"coffee", "tea"}, "coffee")) //true
	fmt.Println(Contains([]int{1, 2, 3}, 11))                  //false
	fmt.Println(Contains([]int{1, 2, 3}, 1))                   //true
}
```
{{< /fig-quote >}}


のように書ける。

`any` も `comparable` も制約としてのみ使用可能で，通常の inteface 型としては使えないようだ。

なお， [Go] の総称型について詳しくは以下の「本」がオススメである。
色々な事例が載っているのでかなり参考になる。

- [Go 言語にやってくる Generics は我々に何をもたらすのか | Zenn](https://zenn.dev/mattn/books/4c7de85ec42cb44cf285)

この本を買うために [Zenn](https://zenn.dev/) のアカウントを取ってもいいかも（クレカが通ればだけど）。

## [Go] は Nearly Java になるか

上述の「本」に

{{< fig-quote title="Go 言語にやってくる Generics は我々に何をもたらすのか" link="https://zenn.dev/mattn/books/4c7de85ec42cb44cf285" >}}
{{% quote %}}その慎重な Go が今、Generics を取り入れ、Better C から Nearly Java へと変わろうとしています。Java がそうであった様に、Go は Generics の導入により多くのアルゴリズムが汎用的に実装され便利になっていくでしょう{{% /quote %}}
{{< /fig-quote >}}

とある。
標準の [`sort`] パッケージのように型ごとに似たようなコードをゴリゴリ書くのではなく，汎化されたアルゴリズムに集約されていく期待感はあるが，それが Nearly Java になるかについては懐疑的である。

[Go 言語][Go]の特徴のひとつは「汎化」の実装として Java や Rust のような公称型の部分型付け（nominal subtyping）ではなく構造型の部分型付け（structural subtyping）を採用していることにある。
当然ながらこれは総称型を使ったアルゴリズムの実装にも大きな影響を与える筈である。

私は [Go] が総称型を得ることで更に独自に発展することを期待している。

## ブックマーク

- [Go Generics draft design Final](https://blog.yongweilun.me/go-generics-draft-design-final?guid=none&deviceId=0389b3ed-c102-4f48-ba65-49e5f54124a4)
- [Go: Type Parameters - Draft Design 抄訳](https://zenn.dev/shuyo/books/4536b976e169ca)

- [それは Duck Typing ぢゃない（らしい）]({{< ref "/remark/2020/04/subtyping.md" >}})

[前回]: {{< relref "./generics-in-go-2-update.md" >}} "次期 Go 言語で導入される総称型について予習する（その2）"
[Go]: https://go.dev/
[`sort`]: https://golang.org/pkg/sort/ "sort - The Go Programming Language"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
