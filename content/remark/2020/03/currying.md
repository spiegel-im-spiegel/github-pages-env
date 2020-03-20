+++
title = "カリー化に関する覚え書き"
date =  "2020-03-20T14:30:42+09:00"
description = "関数型プログラミング言語への馴染みが薄いせいですぐ忘れちゃうのよ。"
image = "/images/attention/kitten.jpg"
tags = [ "engineering", "currying" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

どこぞの某ウイルスのせいでメチャメチャ忙しい。
しかもここ1ヶ月くらいですっかり脅威扇動型ビジネス・モデルへと変貌したようで，ホンマにいい迷惑である。
もはやため息しか出ない。
ので，この件は無視することに決めた。

と，まぁ近況はこれくらいにして，今回は「カリー化」の話。
いや，関数型プログラミング言語への馴染みが薄いせいですぐ忘れちゃうのよ。

というわけで，覚え書きとして記しておく。

## カリーは ~~飲みもの~~ 動詞

Wikipedia によると「カリー化（currying）」とは

{{< fig-quote type="markdown" title="カリー化 - Wikipedia" link="https://ja.wikipedia.org/wiki/%E3%82%AB%E3%83%AA%E3%83%BC%E5%8C%96" >}}
{{% quote %}}複数の引数をとる関数を、引数が「もとの関数の最初の引数」で戻り値が「もとの関数の残りの引数を取り結果を返す関数」であるような関数にすること{{% /quote %}}
{{< /fig-quote >}}

とある。
「カリー」は偉い数学者である Haskell B. Curry の名前から拝借したものらしい。
名前が動詞化しちゃってるのね（笑）

詳しくは近所の数学オタクに訊きなはれ。

## 関数型言語におけるカリー化

ガチの関数型プログラミング言語 [Haskell] ではカリー化は言語仕様に組み込まれていて，たとえば関数 `add` の定義

```haskell
add x y = x + y
```

は実際にはカリー化表現

```haskell
add = \x -> \y -> x + y
```

の糖衣構文となっている[^tpl1]。

[^tpl1]: [Haskell] では関数の引数は1つしかとれないためカリー化は必須の要件となる。意図的にカリー化を避けたいのであれば `add (x, y) = x + y` のように引数を組（tuple）にすればよい。

カリー化のメリットは関数の部分適用（partial application）が作れることで[^pa1]，たとえば

[^pa1]: 部分適用を上手く使えば，いわゆる OAOO (Once And Only Once) 原則に基づいて効率的なコードにできる。念のために言うと，部分適用を構成するのにカリー化は必要条件ではない。

```haskell
increment = add 1
```

とすれば `add` を実引数 `1` で部分適用とした新しい関数 `increment` をシンプルに記述できる。
もちろん，わざわざ名前を付けなくても無名関数として使えばいいのだが。

## 関数型じゃなくてもカリー化はできる

ガチの関数型プログラミング言語じゃなくても第一級関数（first-class function）をサポートするプログラミング言語であればカリー化の記述自体は可能である。

たとえば [Go 言語]なら

```go { hl_lines=["5-9", 13]}
package main

import "fmt"

func add(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	fmt.Println(add(1)(2)) //Output: 3
	increment := add(1) //partial application
	fmt.Println(increment(2)) //Output: 3
}
```

のように書ける。
JavaScript でも

```javascript { hl_lines=["1-5", 8]}
function add(x) {
    return function(y) {
        return x + y;
    };
}

console.log(add(1)(2)); //Output: 3
let increment = add(1); //partial application
console.log(increment(2)); //Output: 3
```

と書くことができる。
さらに JavaScript では[アロー関数式]が使えるので，関数 `add` の定義を

```javascript
const add = x => y => x + y;
```

などと書くことも可能である。
ここまでくると，だいぶ関数型っぽいよね。

## 「それができる」ことと「そのように作られている」ことには天と地ほどの違いがある

この記事を書いて思い出したが，随分前に脊髄反射で

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">そんなこと言ってたら Go でだって関数型モドキな記述はできる。「それができる」ことと「そのように作られている」ことには天と地ほどの違いがある。何言ってるんだ、こいつw <a href="https://t.co/6YsZ0ouDQY">https://t.co/6YsZ0ouDQY</a></p>&mdash; Der Spiegel im Spiegel (@spiegel_2007) <a href="https://twitter.com/spiegel_2007/status/1230417545447870465?ref_src=twsrc%5Etfw">February 20, 2020</a></blockquote>
{{< /fig-gen >}}

と呟いた。
今回の話はまさにそれ。

まぁ，そもそも [Go 言語]の場合はシンプルを旨とする思想な上に構文（statement）による制約が強いため，関数型っぽい記述には（書けるとしても）向いてない。

JavaScript は ES5 以降から関数型の要素を大幅に取り込んでいるが， [Haskell] と比較すれば分かるとおり，「関数」に対する考え方の根本が異なっている。

これは良し悪しの問題ではない。
まさに「制約は構造を生む」で，そうして生み出される構造と実装するシステムとの間で無理なくバランスし続けることがシステムを上手に運用するコツで，それこそが言語を選択する最重要ポイントだと思う（仕事ならね）。

{{< fig-quote type="markdown" title="数学ガール／フェルマーの最終定理" link="https://www.amazon.co.jp/dp/B00I8AT1CM?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
{{% quote %}}公理によって与えられる暗黙の制約。この制約が集合の要素同士をしっかり結びつける。単純にしばるのではない、相互に秩序ある関係を結ぶ。言い換えれば――公理によって与えられる制約が構造を生み出しているのだ{{% /quote %}}
{{< /fig-quote >}}

システムを維持するために遺産や負債を抱え続けなければならない場合もあるが（それでも限度というか寿命はあるけど），そうでないならわざわざレガシーを選択する必然性は微塵もない。

## ブックマーク

- [Haskell 超入門 - Qiita](https://qiita.com/7shi/items/145f1234f8ec2af923ef)
- [カリー化と部分適用（JavaScriptとHaskell） - Qiita](https://qiita.com/7shi/items/a0143daac77a205e7962)

[Haskell]: https://www.haskell.org/ "Haskell Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[アロー関数式]: https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Functions/Arrow_functions "アロー関数 - JavaScript | MDN"
