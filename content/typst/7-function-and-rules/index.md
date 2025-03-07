+++
title = "Typst における関数とルール"
date =  "2025-03-07T19:16:25+09:00"
description = "関数の定義と呼び出し / 名前付き引数 / Show ルール / Set ルール / 文書ファイルを分割した際のルール設定"
image = "/images/attention/tools.png"
tags  = [ "typst", "programming" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

そろそろ
[Typst] の関数とルール設定について，そろそろちゃんと調べようと思って。
本当は「[Typst に関する雑多な話]({{< relref "./x-miscellaneous.md" >}})」で軽く触れるだけにするつもりだったが，思ったより量があったので記事を立てることにした。

## 関数の定義と呼び出し

まずは一番簡単な関数を考えてみる。

```typst
#let fnc(it) = {
  it
};
```

これは引数をそのまま返すだけの関数である。
この関数の引数に `"Hello"` を渡して呼び出すと

```typst
#fnc("Hello")
```

引数の内容がそのまま出力される。

{{< fig-img src="./func-1.png" title="関数の定義と呼び出し (1)" link="./func-1.png" width="647" >}}

引数には（出力可能なものであれば）なんでも渡せる。
数値を入れることもできるし，何なら関数を入れ子にしてもよい。

```typst
#fnc(fnc("Hello"))
```

もうひとつ。
[Typst] の関数呼び出しではコンテントブロック（角括弧 `[ ... ]` で囲まれる領域）を関数呼び出しの後ろに付けることができる。
これをコンテント引数（content argument）と呼ぶ。

```typst
#fnc[Hello]
```

コンテント引数（content argument）は呼び出された側では [`content`] 型の引数としてセットされる。
つまり `#fnc("Hello")` と `#fnc[Hello]` は（[`str`] と [`content`] の型の違い以外は）ほぼ同じ機能である。

ついでの話として，呼び出す側はコンテント引数を複数並べることができる。

```typst
#fnc2[Hello][world]
```

この場合，呼び出される側は以下のように定義する。

```typst {hl_lines=[1]}
#let fnc2(..it) = {
  it
}

#fnc2[Hello][world]
```

このときの出力結果は以下の通り。

{{< fig-img src="./func-1b.png" title="関数の定義と呼び出し (2)" link="./func-1b.png" width="647" >}}

`it` に [`arguments`] 型で受けているのが分かる。
[`arguments`] 型は `pos` メソッドで [`array`] 型に変換できる。

## 名前付き引数

次は文字列を色付きで出力することを考えてみる。
関数定義は以下の通り。

```typst
#let colorText(color: red, it) = {
  text(fill: color)[#it]
}
```

`color: red` は名前付き引数（named parameter）と呼ばれる。
名前付き引数には既定値がつく（上述のコードでは `red` が既定値）。
この関数の呼び出しは以下のように書く。

```typst
#colorText[Hello] #colorText(color: blue)[world]
```

名前付き引数は省略可能で省略した場合は既定値になる。
もちろん

```typst
#colorText("Hello") #colorText(color: blue, "world")
```

などと書くこともできる。
出力結果は同じで以下の通り：

{{< fig-img src="./color-text-1.png" title="関数の定義と呼び出し (2)" link="./color-text-1.png" width="647" >}}

## Show ルール

`show` キーワードを使って指定した対象に対し show ルールを設定できる。

```typst
#let colorText(color: red, it) = {
  text(fill: color)[#it]
}

#show heading: it => colorText(color: blue, it)

= Heading 1

== Heading 1.1
```

上のコードの `it` は `show` キーワードで指定した対象を無名関数のコンテント引数として渡したもの。
全ての関数で共通に用意されている `with` メソッドを使うことで以下のようにコンテント引数を省略できる。

```typst {hl_lines=[5]}
#let colorText(color: red, it) = {
  text(fill: color)[#it]
}

#show heading: colorText.with(color: blue)

= Heading 1

== Heading 1.1
```

出力結果はいずれも同じでこんな感じ。

{{< fig-img src="./show-rule-1.png" title="Show ルール設定 (1)" link="./show-rule-1.png" width="647" >}}

Show ルールの対象は（出力可能なものなら）なんでもよくて，たとえば

```typst
#let colorText(color: red, it) = {
  text(fill: color)[#it]
}

#show "Hello": colorText

Hello World
```

みたいな記述もできる。
なお，コンテント引数以外に引数がない（または省略できる）場合は，上のように関数名だけを指定できる。

このコードの出力結果は以下の通り：

{{< fig-img src="./show-rule-2.png" title="Show ルール設定 (2)" link="./show-rule-2.png" width="647" >}}

Show ルールの対象には [`selector`] も指定できる。
たとえば `where` メソッドを使って以下のように記述できる。

```typst {hl_lines=[5]}
#let colorText(color: red, it) = {
  text(fill: color)[#it]
}

#show heading.where(level: 1).or(heading.where(level: 2)): colorText

= Heading 1

== Heading 1.1

=== Heading 1.1.1
```

出力結果は以下の通り：

{{< fig-img src="./show-rule-3.png" title="Show ルール設定 (3)" link="./show-rule-3.png" width="647" >}}

Show ルールの対象がない場合，文書全体が対象となる。

```typst {hl_lines=[5]}
#let colorText(color: red, it) = {
  text(fill: color)[#it]
}

#show: colorText.with(color: blue)

#lorem(40)
```

ちなみに [`lorem`] は指定した単語数のランダムな文章（欧文[^l1]）を生成する関数である。

[^l1]: 日本語の文を生成する [`roremu` パッケージ](https://zenn.dev/mkpoli/articles/7e54c1c780ff43 "Typstの日本語Lipsumパッケージを作ってみた件")を公開されている方もいる。感謝！

出力結果は以下の通り：

{{< fig-img src="./show-rule-4.png" title="Show ルール設定 (4)" link="./show-rule-4.png" width="647" >}}

## Set ルール

[`text`] や [`image`] あるいは [`heading`] や [`par`] などドキュメント要素に紐づく組込み関数は「要素関数（element function）」と言うそうな。
要素関数は `set` キーワードを使い set ルールで名前付き引数の既定値を変更することができる。

{{< fig-quote type="markdown" title="Function Type – Typst Documentation" link="https://typst.app/docs/reference/foundations/function/#element-functions" lang="en" >}}
Some functions are associated with elements like [headings](https://typst.app/docs/reference/model/heading/) or [tables](https://typst.app/docs/reference/model/table/). When called, these create an element of their respective kind. In contrast to normal functions, they can further be used in [set rules](https://typst.app/docs/reference/styling/#set-rules), [show rules](https://typst.app/docs/reference/styling/#show-rules), and [selectors](https://typst.app/docs/reference/foundations/selector/).
{{< /fig-quote >}}

ドキュメントのルートで set ルールを指定すると，逐次処理で指定位置以降に適用される。

```typst
#set text(font: "Noto Serif CJK JP", lang: "ja")

明朝体

#set text(font: "Noto Sans CJK JP", lang: "ja")

ゴシック体
```

出力結果は以下の通り：

{{< fig-img src="./set-rule-1.png" title="Set ルール設定 (1)" link="./set-rule-1.png" width="647" >}}

Show ルールの中で set ルールを指定した場合，その show ルールの中でのみ set ルールが適用される。

```typst
#set text(font: "Noto Serif CJK JP", lang: "ja")
#show heading: it => {
  set text(font: "Noto Sans CJK JP")
  it
}

= ゴシック体

明朝体
```

`set` キーワードのみの記述ならコンテント引数を省略して

```typst {hl_lines=[2]}
#set text(font: "Noto Serif CJK JP", lang: "ja")
#show heading: set text(font: "Noto Sans CJK JP")

= ゴシック体

明朝体
```

などと記述することもできる。

出力結果はいずれも同じでこんな感じ。

{{< fig-img src="./set-rule-2.png" title="Set ルール設定 (2)" link="./set-rule-2.png" width="647" >}}

なお `let` キーワードで定義したユーザ関数は要素関数ではないので **set ルールは使えない**。

## 文書ファイルを分割した際のルール設定

文書ファイルを分割して [`import` や `include`][module] で読み込む場合，子ドキュメントで設定した set および show ルールは親ドキュメントには適用されないので注意（親ドキュメントから子ドキュメントへはルールが継承される）。
子ドキュメントで記述した set および show ルールを親ドキュメントに適用させるには（[テンプレート][template]で使う手法）ルールを記述した関数を [`import`][module] で読み込み，親ドキュメントの show ルールを使って関数を呼び出す。

{{< fig-gen type="markdown" title="子ドキュメント" >}}
```typst
#let initFonts(font-name: "New Computer Modern", body) = {
  set text(
    lang: "ja",
    font: font-name,
    size: 10pt,
  )
  body
}
```
{{< /fig-gen >}}

{{< fig-gen type="markdown" title="親ドキュメント (1)" >}}
```typst
#import "child.typ": initFonts

#show: it => initFonts(
  font-name: "Noto Sans CJK JP",
  it
)

こんにちわ、世界！
```
{{< /fig-gen >}}

または `with` メソッドを使って

{{< fig-gen type="markdown" title="親ドキュメント (2)" >}}
```typst {hl_lines=["3-5"]}
#import "child.typ": initFonts

#show: initFonts.with(
  font-name: "Noto Sans CJK JP",
)

こんにちわ、世界！
```
{{< /fig-gen >}}

などとコンテント引数を省略できる。








## ブックマーク

ブックマークは「[Typst に関するブックマーク]({{< relref "./0-bookmark.md" >}})」にてまとめています。

[Typst]: https://typst.app/ "Typst: Compose papers faster"
[Typst Documentation]: https://typst.app/docs/ "Typst Documentation"
[Tutorial]: https://typst.app/docs/tutorial "Tutorial – Typst Documentation"
[module]: https://typst.app/docs/reference/scripting/#modules "Scripting – Typst Documentation"
[template]: https://typst.app/docs/tutorial/making-a-template/ "Making a Template – Typst Documentation"
[`content`]: https://typst.app/docs/reference/foundations/content/ "Content Type – Typst Documentation"
[`str`]: https://typst.app/docs/reference/foundations/str/ "String Type – Typst Documentation"
[`arguments`]: https://typst.app/docs/reference/foundations/arguments/ "Arguments Type – Typst Documentation"
[`selector`]: https://typst.app/docs/reference/foundations/selector/ "Selector Type – Typst Documentation"
[`array`]: https://typst.app/docs/reference/foundations/array/ "Array Type – Typst Documentation"
[`dictionary`]: https://typst.app/docs/reference/foundations/dictionary/ "Dictionary Type – Typst Documentation"
[`image`]: https://typst.app/docs/reference/visualize/image/ "Image Function – Typst Documentation"
[`text`]: https://typst.app/docs/reference/text/text/ "Text Function – Typst Documentation"
[`lorem`]: https://typst.app/docs/reference/text/lorem/ "Lorem Function – Typst Documentation"
[`heading`]: https://typst.app/docs/reference/model/heading/ "Heading Function – Typst Documentation"
[`par`]: https://typst.app/docs/reference/model/par/ "Paragraph Function – Typst Documentation"

## 参考文献

{{% review-paapi "B0DPXBNTRS" %}} <!-- Typst完全入門-->
