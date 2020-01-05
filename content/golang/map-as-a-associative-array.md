+++
title = "Map は連想配列ではなく連想配列への「参照」である"
date =  "2019-06-07T23:07:39+09:00"
description = "いかに参照っぽく振る舞っていようとも Go 言語でやり取りできるのはあくまでも「値」であり，その「値」が何を指しているかを考えながらコードを書いていく必要がある。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "programming", "map" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今回は小ネタで。
つか，だいぶフワッとした話になるので，その辺は割引いて読んでいただけるとありがたい。

いやね。

- [Goのスコープに苦しんだ事例集 - Qiita](https://qiita.com/bushiyama/items/fafc0d2d64e4e1b1d0ae)

の「事例その１」で

```go
type hogehoge map[string]string

func (h *hogehoge) Seter(p string) {
    h = &hogehoge{
        "key": p,
    }
}

func main() {
    var f hogehoge
    f.Seter("ddd")

    fmt.Println(f) // <- nil
}
```

というコードを書いておられて，言わんとすることは分かるけど事例に [map] を使うのは混乱を助長しないかなぁ，と思ってしまったのだ。
そこでこの記事では主に [map] の振る舞いにピントを合わせて説明していこう。

## [Map] は連想配列ではなく連想配列への「参照」である

のっけからぶっちゃけるが，要するにそういうことである。
もう少し厳密に言うと「[Map] は連想配列への「参照」のように振る舞う」といったところか。

これを端的に表すコードがこれ。

```go
package main

import "fmt"

type KeyValues map[string]string

func main() {
	kv := KeyValues{"foo": "bar"}
	fmt.Printf("%p: %p: %v\n", &kv, kv, kv)
}
```

これを[実行する](https://play.golang.org/p/TXrq6yJmbjK)と

```text
0x40c128: 0x43e260: map[foo:bar]
```

などと表示される。

最初の値がインスタンス `kv` のポインタ値，次の値が `kv` が参照する連想配列のポインタ値，最後が `kv` の内容と考えれば分かりやすいだろう。

## [Go 言語]において関数の引数は「値渡し」なので...

型 `KeyValues` にひとつ関数を追加してみよう。

```go
package main

import "fmt"

type KeyValues map[string]string

func (kv KeyValues) Set(k, v string) {
	kv[k] = v
	fmt.Printf("%p: %p: %v\n", &kv, kv, kv)
}

func main() {
	kv := KeyValues{}
	fmt.Printf("%p: %p: %v\n", &kv, kv, kv)
	kv.Set("foo", "bar")
	fmt.Printf("%p: %p: %v\n", &kv, kv, kv)
}
```

これを[実行する](https://play.golang.org/p/aIf6wNNsMQK)と

```text
0x40c128: 0x43e260: map[]
0x40c138: 0x43e260: map[foo:bar]
0x40c128: 0x43e260: map[foo:bar]
```

となる。

ポイントは2行目で `main()` 関数内のインスタンス `kv` と `KeyValues.Set()` 関数のメソッド・レシーバの `kv` は異なるポインタ値になっているので異なるインスタンスだと分かるが，参照している連想配列（へのポインタ）は同一である。

もっと簡単なコードで示そうか。

```go
package main

import "fmt"

type KeyValues map[string]string

func main() {
	kv := KeyValues{"foo": "bar"}
	fmt.Printf("%p: %p: %v\n", &kv, kv, kv)
	cpy := KeyValues{}
	fmt.Printf("%p: %p: %v\n", &cpy, cpy, cpy)
	cpy = kv
	fmt.Printf("%p: %p: %v\n", &cpy, cpy, cpy)
}
```

これを[実行する](https://play.golang.org/p/vLaMjYxo9Y4)と

```text
0x40c128: 0x43e260: map[foo:bar]
0x40c148: 0x43e2a0: map[]
0x40c148: 0x43e260: map[foo:bar]
```

となる。
つまり [map] は連想配列への参照なので，連想配列そのものは宣言構文（`:=`）や代入構文（`=`）では複製できない，ということである[^state1]。

[^state1]: [Go 言語]においては宣言や代入は式（expression）ではなく構文（statement）であり `:=` や `=` は演算子ではなく構文を構成する（`var` とかと同じ）トークンに過ぎない。ちなみに [`++` や `--` も演算子ではなく代入構文のトークン]({{< relref "./operators-and-statements.md" >}} "演算子とステートメント")である。

## それでもポインタは「参照」ではない

じゃあ，先ほどの `KeyValues.Set()` 関数のメソッド・レシーバをポインタ型にするとどうなるか，やってみよう。

```go
package main

import "fmt"

type KeyValues map[string]string

func (kv *KeyValues) Set(k, v string) {
	(*kv)[k] = v
	fmt.Printf("%p: %p: %v\n", kv, *kv, *kv)
}

func main() {
	kv := KeyValues{}
	fmt.Printf("%p: %p: %v\n", &kv, kv, kv)
	kv.Set("foo", "bar")
	fmt.Printf("%p: %p: %v\n", &kv, kv, kv)
}
```

これを[実行する](https://play.golang.org/p/dLjVZ9rUjtJ)と

```text
0x40c128: 0x43e260: map[]
0x40c128: 0x43e260: map[foo:bar]
0x40c128: 0x43e260: map[foo:bar]
```

となる。

`main()` 内のインスタンス `kv` はポインタ型ではないが `KeyValues.Set()` 関数呼び出し時に暗黙的な変換が行われる。
このため `KeyValues.Set()` 関数のメソッド・レシーバの `kv` は `main()` 内のインスタンス `kv` と同一のインスタンスになるわけだ。

では調子に乗ってこんな関数を作ってみよう。

```go
func (kv *KeyValues) Initialize(k, v string) {
	kv = &KeyValues{k: v}
}
```

`KeyValues.Initialize()` 関数のメソッド・レシーバは呼び出し元のインスタンスと同一なんだから，これで初期化できるんじゃね？ というわけだ。

プログラム全体はこんな感じ。
これでようやく最初に紹介したコードにほぼ近い形になっただろう。

```go
package main

import "fmt"

type KeyValues map[string]string

func (kv *KeyValues) Initialize(k, v string) {
	kv = &KeyValues{k: v}
	fmt.Printf("%p: %p: %v\n", kv, *kv, *kv)
}

func main() {
	kv := KeyValues{}
	fmt.Printf("%p: %p: %v\n", &kv, kv, kv)
	kv.Initialize("foo", "bar")
	fmt.Printf("%p: %p: %v\n", &kv, kv, kv)
}
```

まぁ[動かして](https://play.golang.org/p/gUFL_j5NzLa)みよう（笑）

```text
0x40c128: 0x43e260: map[]
0x40c138: 0x43e2a0: map[foo:bar]
0x40c128: 0x43e260: map[]
```

ちょっと考えれば分かるのだが， `KeyValues.Initialize()` 関数のメソッド・レシーバに渡されるのはポインタ「値」なので，それを関数内で上書きしたところで呼び出し元の `main()` 関数には全く影響ないのである。

これが「[Go 言語に『参照』は存在しない]({{< relref "./for-range-statement.md" >}} "For-Range 構文の話")」ということの意味である。
いかに参照っぽく振る舞っていようとも [Go 言語]でやり取りできるのはあくまでも「値」であり，その「値」が何を指しているかを考えながらコードを書いていく必要がある。

たとえば先ほどの `KeyValues.Initialize()` 関数を以下のように書き換えればまた挙動が変わる。

{{< highlight go "hl_lines=8" >}}
package main

import "fmt"

type KeyValues map[string]string

func (kv * KeyValues) Initialize(k, v string) {
	* kv = KeyValues{k: v}
	fmt.Printf("%p: %p: %v\n", kv, * kv, * kv)
}

func main() {
	kv := KeyValues{}
	fmt.Printf("%p: %p: %v\n", &kv, kv, kv)
	kv.Initialize("foo", "bar")
	fmt.Printf("%p: %p: %v\n", &kv, kv, kv)
}
{{< /highlight >}}

これを[実行する](https://play.golang.org/p/HT7ApB21ka0)と

```text
0x40c128: 0x43e260: map[]
0x40c128: 0x43e2a0: map[foo:bar]
0x40c128: 0x43e2a0: map[foo:bar]
```

となる。
なにがどう違うのか考えてみよう。
ここまでくれば簡単だよね（笑）

## ブックマーク

- [Map の話]({{< relref "./map.md" >}})


[Go 言語]: https://golang.org/ "The Go Programming Language"
[Map]: http://golang.org/ref/spec#Map_types
[map]: http://golang.org/ref/spec#Map_types

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
