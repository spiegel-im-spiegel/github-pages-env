+++
title = "次期 Go 言語で導入される総称型について予習する（その2）"
date =  "2020-06-17T16:33:04+09:00"
description = "総称型について仕様が変わったらしい。早ければ2021年8月のリリースでお披露目されるかもしれないとのこと。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "language", "programming", "generics", "type", "contract", "constraint" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

以前書いた

- [次期 Go 言語で導入される（かもしれない）総称型について予習する]({{< relref "./generics-in-go-2.md" >}})

に関して，仕様が変わったらしい。

- [The Next Step for Generics - The Go Blog](https://blog.golang.org/generics-next-step)
- [Type Parameters - Draft Design](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md)

そこで，これまでのおさらいと変更点について簡単に紹介する。

## 「総称型」のおさらい

次期 [Go 言語]で導入される「総称型（generics）」について，これまで提案されていた仕様は以下の通り。

```go
contract ordered(t T) {
    t < t
}

func Max(type T ordered)(x, y T) T {
    if x < y {
        return y
    }
    return x
}
```

`Max()` 関数の `(type T ordered)` の部分が「型パラメータ（type parameter）」で，これによって型 `T` が総称型であることを示す。
また `ordered` は「型コントラクト（type contract）」と呼ばれ，総称型に対する制約（constraint）として機能している。

## 制約の定義

新しい提案では，型コントラクトを捨て，制約の定義を以下のように変更する。

```go {hl_lines=["1-6"]}
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

以前の型コントラクトとの違いは **既存のキーワードがそのまま使える** ことである。
文法も `interface` 型の宣言とほぼ同じだが `type int, int8, ...` といった感じに型の列挙ができる点がこれまでと異なる。

これにより，制約 `ordered` を満たす型であれば

```go
func main() {
	fmt.Println(Max(Max(1, 2), 3))               //Output: 3
    fmt.Printf("%#U\n", Max(Max('あ', 'い'), 0)) //Output: U+3044 'い'
}
```

などと記述できる。

なお，これまでの `interface` 型も制約として使える。
たとえば，こんな感じ。

```go
func SayError(type E error)(e E) string {
	return e.Error()
}

func main() {
	fmt.Println(SayError(io.EOF)) //Output: EOF
}
```

更に `interface` 型を入れ子にすることで複雑な制約条件を構成することもできる。

## 注入するか単態化するか

ところで先程の

```go
func SayError(type E error)(e E) string {
	return e.Error()
}
```

は，わざわざ総称型を使わなくとも今までの書き方で

```go
func SayError(e error) string {
	return e.Error()
}
```

としても見た目の結果は同じになる。

しかし，今までの書き方が実行時に動的ディスパッチを行う「依存の注入（depencency injection）」であるのに対し，総称型を使った記述はコンパイル時に型ごとに別々の関数として展開される。
これを「単態化（monomorphization）」と呼ぶ。

依存の注入と単態化ではリソース管理の戦略がまるで異なる。

たとえば上述の `SayError()` 関数を総称型で記述するのは，どう考えても効率が悪い。
`error` に包摂される型は必要に応じていくらでも定義可能であり，それらの型ごとに単態化されるなどぞんぞがさばる。

逆に `Max()` 関数なんかは，基本型の振る舞いを `interface` で括るのが難しいこともあり，コストの高い `float64` 型の [`math`]`.Max()` くらいしか標準では実装されてなかったりする。
こういったものを総称型で記述できるようになれば（OAOO 原則の観点から考えても）コード管理しやすくなるだろう。

「[Go プログラマは息をするように依存を注入する](https://slide.baldanders.info/shimane-go-2020-01-23/)」。
他のプログラミング言語が総称型で書いているかなりの部分を依存の注入で簡単に纏めてしまえるのが [Go 言語]の強みであるが，使える道具は多いことに越したことはない。

## [Go] 2 Playground

次期 [Go 言語]用に新しい Playground が公開されている。

- [`https://go2goplay.golang.org/`](https://go2goplay.golang.org/)

総称型についてもここで色々と試せるようなので，遊びに行ってみてはいかがだろうか。

## 早ければ [Go] 1.17 でお披露目？

だいぶ仕様が固まってきた総称型であるが，早ければ2021年8月にリリース予定の [Go] 1.17 で[お披露目されるかもしれない](https://blog.golang.org/generics-next-step "The Next Step for Generics - The Go Blog")とのこと。

たのしみ！

## ブックマーク

- [Why no max/min function for integer in GoLang | Pixelstech.net](https://www.pixelstech.net/article/1559993656-Why-no-max-min-function-for-integer-in-GoLang)
- [それは Duck Typing ぢゃない（らしい）]({{< ref "/remark/2020/04/subtyping.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`math`]: https://pkg.go.dev/math "math package · pkg.go.dev"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
