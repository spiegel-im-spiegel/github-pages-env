+++
title = "Typst のドキュメント要素"
date =  "2025-02-23T13:01:15+09:00"
description = "コードモードとルール定義 / 見出し / 段落 / 引用 / 打ち込んだ通りに出力する / 箇条書き / 図表"
image = "/images/attention/tools.png"
tags  = [ "typst" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

[前回]でようやく環境が整ったので [Typst] のドキュメント構造を見ていこう。

## コードモードとルール定義

[Typst] では `set` キーワードで設定ルールを `show` キーワードで表示ルールを定義する。
[前回]の記事で言うと文書全体のフォントを指定している

```typst
// for main text
#set text(font: (
  (
    name: "New Computer Modern",
    covers: "latin-in-cjk",
  ),
  "BIZ UDMincho"
))
```

が（文書全体の）設定ルールを定義している部分である。
文字コードの範囲とフォント名をセットで指定するのは [Typst] 0.13 からの仕様らしい。
`latin-in-cjk` はラテン文字の範囲を示す [Typst] の予約語。
コード範囲の指定に正規表現も使えるそうな。

同じく[前回]記事の

```typst
// for headings
#show heading: it => {
    set text(font: (
      (
        name: "Liberation Sans",
        covers: "latin-in-cjk",
      ),
      "BIZ UDGothic"
    ))
	it.body
}
```

は表示ルールについての記述で，見出し（[`heading`]）要素に対して表示時のフォントを指定している。
表示ルール `show` の中で設定ルール `set` が使われていることに注目してほしい。
ここでの `set` は `show` による定義内でのみ効力がある。

キーワードの頭に付いている `#` はコードモードの開始を示している。
他にも `$ ... $` で囲まれている部分は数式モード， `[ ... ]` で囲まれている部分はマークアップモードといった感じにモードが切り替わる。
指定されている部分以外はマークアップモード。
$\mathrm{\LaTeX}$ では文書を記述する領域（環境）が明示的に示され，それ以前の部分をプリアンブル（preamble; 前口上）として設定やルールを纏めて記述するが [Typst] ではそういった区別はなく，コードモードを任意の場所に記述可能で逐次処理される。

コードモードは基本的に文単位で `#` を付けるが `#{ ... }` とブロック単位で指定することもできる。
`set` や `show` はコードモードで使えるキーワードである。
コードモードで使えるキーワードについては[公式ドキュメント][Typst Documentation]の “[Syntax](https://typst.app/docs/reference/syntax/ "Syntax – Typst Documentation")” のページを参照のこと。

[`text`] はテキストの外観とレイアウトを定義・変更する関数で，ルール設定以外にも頻繁に登場する。
引数で複数のプロパティ値を連想配列で指定する。

```typst
#set text(font: "Noto Sans CJK JP", size: 14pt, lang: "jp")

空が#text(fill:blue)[青い] 。
```

{{< fig-img src="./pdf-sample-code-01.png" title="文字の色を変える" link="./pdf-sample-code-01.png" width="713" >}}

## 見出し

ついでなので見出し（[`heading`]）の話からしよう。

マークアップモードで行頭に `=` を付けると見出しになる。

```typst
#set text(font: "Noto Sans CJK JP")

= 見出し1
== 見出し1.1
=== 見出し1.1.1
==== 見出し1.1.1.1
= 見出し2
```

フォント指定部分はスルーしてほしい。
`=` と見出しの文言の間に空白文字を入れるのがポイント。
これをコンパイルすると以下の出力になる。

{{< fig-img src="./pdf-sample-hd-01.png" title="見出し (1)" link="./pdf-sample-hd-01.png" width="713" >}}

`=` を重ねることで見出しのレベルが深くなるが，レベルの制限は（仕様上は）ないみたい。
レベルが深くなると見た目で判別できなくなるけど。

[`heading`] 関数を使ってコードモードで見出しを表示することもできる。

```typst
#set text(font: "Noto Sans CJK JP")

#outline(title: [目次])

#heading(level:1)[見出し1]
#heading(level:2)[見出し1-1]
#heading(
  level:1,
  outlined: false
)[目次に乗らない付録]
```

{{< fig-img src="./pdf-sample-hd-02.png" title="見出し (2)" link="./pdf-sample-hd-02.png" width="713" >}}

[`outline`] は目次を出力する関数。

コードモードで見出しを書くメリットはあまりないと思うけど，上のように特定の見出しを目次に含めないとかいった制御が必要な場合があるかもしれない。

見出し番号のフォーマットは `set` キーワードでルールを定義するのがよい。

```typst {hl_lines=[2]}
#set text(font: "Noto Sans CJK JP")
#set heading(numbering: "1.1")

= 見出し1
== 見出し1.1
=== 見出し1.1.1
== 見出し1.2
= 見出し2
== 見出し2.1
```

{{< fig-img src="./pdf-sample-hd-03.png" title="見出し (3)" link="./pdf-sample-hd-03.png" width="713" >}}

[`numbering`] の表現はかなり強力で

```typst {hl_lines=[2,7,8]}
#set text(font: "Noto Sans CJK JP")
#set heading(numbering: "ア.")

= 見出し1
= 見出し2

#let unary(.., last) = "†" * last
#set heading(numbering: unary)
#counter(heading).update(0) //見出しカウンタのリセット

= 付録1
= 付録2
```

{{< fig-img src="./pdf-sample-hd-04.png" title="見出し (4)" link="./pdf-sample-hd-04.png" width="713" >}}

みたいなこともできる。

{{< fig-quote type="markdown" title="Numbering Function – Typst Documentation" link="https://typst.app/docs/reference/model/numbering/" lang="en" >}}
*Counting symbols* are 1, a, A, i, I, α, Α, 一, 壹, あ, い, ア, イ, א, 가, ㄱ, *, ١, ۱, १, ১, ক, ①, and ⓵. They are replaced by the number in the sequence, preserving the original case.

The * character means that symbols should be used to count, in the order of *, †, ‡, §, ¶, ‖. If there are more than six items, the number is represented using repeated symbols.
{{< /fig-quote >}}

ちなみに [`counter`] は関数ではなく型。
[`heading`] 関数を [`counter`] 型にキャストして `update` メソッドを呼び出すイメージか？

## 段落

マークアップモードでは，連続したひとかたまりの文字列を段落としてあつかう。
改行は無視される。
段落を改めるには空行を挟む。

```typst
#set text(font: "Noto Serif CJK JP")

「何人ものニュートンがいた（There were several Newtons）」
と言ったのは，科学史家ハイルブロンである．同様にコーヘンは
「ニュートンはつねに二つの貌を持っていた
（Newton was always ambivalent）」と語っている．

近代物理学史上でもっとも傑出しもっとも影響の大きな人物が
ニュートンであることは，誰しも頷くことであろう．
しかしハイルブロンやコーヘンの言うように，
ニュートンは様々な，ときには相矛盾した顔を持ち，
その影響もまた時代とともに大きく変っていった．
```

{{< fig-img src="./pdf-sample-par-01.png" title="段落 (1)" link="./pdf-sample-par-01.png" width="713" >}}

{{< div-gen class="center" >}}
〈山本義隆『熱学思想の史的展開』（現代数学社，1987年）より〉
{{< /div-gen >}}

この辺は $\mathrm{\LaTeX}$ や Markdown と同じなので馴染みがあるだろう。

段落内で強制的に改行したい場合は行末に `\` を付ける（または `#`[`linebreak`]`()` 関数を使う）。

```typst {hl_lines=[4]}
#set text(font: "Noto Serif CJK JP")

近代物理学史上でもっとも傑出しもっとも影響の大きな人物が
ニュートンであることは，誰しも頷くことであろう． \
しかしハイルブロンやコーヘンの言うように，
ニュートンは様々な，ときには相矛盾した顔を持ち，
その影響もまた時代とともに大きく変っていった．
```

{{< fig-img src="./pdf-sample-par-02.png" title="段落 (2)" link="./pdf-sample-par-02.png" width="713" >}}

[`par`] 関数を使ってコードモードで段落を表示することもできる。

```typst
#set text(font: "Noto Serif CJK JP", lang: "jp")

#par(justify: false)[
「何人ものニュートンがいた（There were several Newtons）」
と言ったのは，科学史家ハイルブロンである．同様にコーヘンは
「ニュートンはつねに二つの貌を持っていた
（Newton was always ambivalent）」と語っている．
]

#par(justify: true)[
「何人ものニュートンがいた（There were several Newtons）」
と言ったのは，科学史家ハイルブロンである．同様にコーヘンは
「ニュートンはつねに二つの貌を持っていた
（Newton was always ambivalent）」と語っている．
]
```

{{< fig-img src="./pdf-sample-par-03.png" title="段落 (3)" link="./pdf-sample-par-03.png" width="713" >}}

`justify` 指定の有無でレイアウトがどう変わるか見て欲しい。

実際にはコードモードで個別に指定するより `set` を使って段落の設定ルールを定義することのほうが多いだろう。

```typst {hl_lines=["2-10"]}
#set text(font: "Noto Serif CJK JP", lang: "jp")
#set par(
  first-line-indent: (
    amount: 1em,
    all: true,
  ),
  leading: 0.9em,
  spacing: 0.9em,
  justify: true,
)

「何人ものニュートンがいた（There were several Newtons）」
と言ったのは，科学史家ハイルブロンである．同様にコーヘンは
「ニュートンはつねに二つの貌を持っていた
（Newton was always ambivalent）」と語っている．

近代物理学史上でもっとも傑出しもっとも影響の大きな人物が
ニュートンであることは，誰しも頷くことであろう．
しかしハイルブロンやコーヘンの言うように，
ニュートンは様々な，ときには相矛盾した顔を持ち，
その影響もまた時代とともに大きく変っていった．
```

{{< fig-img src="./pdf-sample-par-04.png" title="段落 (4)" link="./pdf-sample-par-04.png" width="713" >}}

[Typst] の以前のバージョンでは最初の段落の最初の文字の字下げが出来ない問題があったが，バージョン 0.13 から `first-line-indent` の `all` プロパティを `true` にすることで字下げできるようになった。
めでたい！

{{< div-box type="markdown" >}}
**【字下げが効かない問題は解消したため以下の記述は obsolete となった】**

上の例で段落の最初の文字を字下げする設定にしてみたのだが，どうも最初の段落だけ字下げが効かないようだ。
欧文では[最初の段落は字下げしないのが主流](https://note.com/jinbunxshakai/n/n6a72f690f11c "最初のパラグラフはインデントしない!? 欧文組版の傾向は？｜『人文×社会』の中の人")とかいう話もあるらしく，バグじゃなくて意図的な仕様じゃないかと言われている。
一応[回避策](https://adbird.hatenablog.com/entry/2024/03/21/015335#%E5%AD%97%E4%B8%8B%E3%81%92%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6 "Typst 備忘録 - adbird（広告鳥） 備忘録")はあるそうなのだが，副作用もあるみたいだし，積極的に使いたい方法ではないかなぁ。
オンライン・テキストに慣れてしまうと字下げ云々はあまり気にならなくなるけど，書籍出版や論文を書く人など決められた書式で書く人にとっては問題だろう。

[Typst]: https://typst.app/ "Typst: Compose papers faster"
{{< /div-box >}}

## 引用

引用（[`quote`]）の表現はコードモードのみ対応しているようだ。

```typst
#set text(font: "Noto Serif CJK JP", lang: "jp")

引用文：

#quote(
  block: true,
  attribution: [山本義隆『熱学思想の史的展開』（現代数学社，1987年）]
)[
#quote(block: false)[何人ものニュートンがいた（There were several Newtons）]
と言ったのは，科学史家ハイルブロンである．同様にコーヘンは
#quote(block: false)[ニュートンはつねに二つの貌を持っていた
（Newton was always ambivalent）]と語っている．

近代物理学史上でもっとも傑出しもっとも影響の大きな人物が
ニュートンであることは，誰しも頷くことであろう．
しかしハイルブロンやコーヘンの言うように，
ニュートンは様々な，ときには相矛盾した顔を持ち，
その影響もまた時代とともに大きく変っていった．
]
```

{{< fig-img src="./pdf-sample-quote-01.png" title="引用文" link="./pdf-sample-quote-01.png" width="713" >}}

インラインの [`quote`] は言語によって括弧を変えてくれるらしいのだが，日本語の「 ... 」には対応していない模様。
残念。

## 打ち込んだ通りに出力する

$\mathrm{\LaTeX}$ で言うところの `verbatim` 環境（または [`listings`](https://ctan.org/tex-archive/macros/latex/contrib/listings "CTAN: /tex-archive/macros/latex/contrib/listings") パッケージ）， Markdown で言うところのコードフェンスかな。
マークアップモードでの書式は Markdown と同じ。

{{< highlight typst >}}
#show raw: body => {
    set text(font: (
      (
        name: "Inconsolata",
        covers: "latin-in-cjk",
      ),
      "Noto Sans CJK JP"
    ))
    body
}

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```
{{< /highlight >}}

{{< fig-img src="./pdf-sample-raw-01.png" title="打ち込んだ通りに出力する (1)" link="./pdf-sample-raw-01.png" width="713" >}}

おー。
ちゃんとフォントも指定できた。

一応 [`raw`] 関数を使った記述もできるけど実用的じゃないかな。

```typst
#raw("package main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello, 世界\")\n}", block: true, lang: "go")
```

[`raw`] はインラインで使うと面白い。

```typst
#set text(font: "Noto Serif CJK JP", lang: "jp")
#show raw: body => {
    set text(font: (
      (
        name: "Inconsolata",
        covers: "latin-in-cjk",
      ),
      "Noto Sans CJK JP"
    ))
    body
}

Go 言語では最初に ```go package main``` と指定する。

Go 言語では最初に #raw("package main", lang: "go") と指定する。
```

{{< fig-img src="./pdf-sample-raw-02.png" title="打ち込んだ通りに出力する (2)" link="./pdf-sample-raw-02.png" width="713" >}}

こんな感じにインラインでも syntax highlight が効く。

## 箇条書き

マークアップモードでは `+` および `-` で箇条書きを表現できる。

```typst
#set text(font: "Noto Serif CJK JP", lang: "jp")

+ 番号付き箇条書き1
  + 番号付き箇条書き1-1
  + 番号付き箇条書き1-2
+ 番号付き箇条書き2
  - 番号なし箇条書き2-1
  - 番号なし箇条書き2-2

- 番号なし箇条書き1
  + 番号付き箇条書き1-1
  + 番号付き箇条書き1-2
- 番号なし箇条書き2
  - 番号なし箇条書き2-1
  - 番号なし箇条書き2-2

5. 番号付き箇条書き5 (番号指定)
+ 番号付き箇条書き6
```

{{< fig-img src="./pdf-sample-list-01.png" title="箇条書き (1)" link="./pdf-sample-list-01.png" width="713" >}}

またはコードモードで [`enum`] または [`list`] 関数を使って

```typst
#set text(font: "Noto Serif CJK JP", lang: "jp")

#enum(
  numbering: "①",
  enum.item[番号付き箇条書き1
    #enum(numbering: "ア)")[番号付き箇条書き1-1][番号付き箇条書き1-2]
  ],
  enum.item[番号付き箇条書き2],
)

#list(
  marker: [‣],
  list.item[番号なし箇条書き1
    #list(marker: [--])[番号なし箇条書き1-1][番号なし箇条書き1-2]
  ],
  list.item[番号なし箇条書き2],
)
```

{{< fig-img src="./pdf-sample-list-02.png" title="箇条書き (2)" link="./pdf-sample-list-02.png" width="713" >}}

てな感じに記述できる。

箇条書きの番号やシンボルは `set` であらかじめルールを設定しておくのがいいだろう。

```typst
#set text(font: "Noto Serif CJK JP", lang: "jp")

#set enum(numbering: "①.ア")
#set list(marker: ([‣], [--]))

+ 番号付き箇条書き1
  + 番号付き箇条書き1-1
  + 番号付き箇条書き1-2
+ 番号付き箇条書き2

- 番号なし箇条書き1
  - 番号なし箇条書き1-1
  - 番号なし箇条書き1-2
- 番号なし箇条書き2
```

{{< fig-img src="./pdf-sample-list-03.png" title="箇条書き (3)" link="./pdf-sample-list-03.png" width="713" >}}

## 図表

表（[`table`]）に関してはコードモードによる記述のみ提供されている。
こんな感じ。

```typst
#set text(font: "Noto Serif CJK JP", lang: "jp")

#table(
  columns: (auto, auto, auto),
  align: (right, center, left),
  table.header([*日付*], [*曜日*], [*内容*]),
  [2025年5月3日], [土], [憲法記念日],
  [2025年5月4日], [日], [みどりの日],
  [2025年5月5日], [月], [こどもの日],
  [2025年5月6日], [火], [休日],
)
```

{{< fig-img src="./pdf-sample-table-01.png" title="表 (1)" link="./pdf-sample-table-01.png" width="713" >}}

[`table`] を [`figure`] で囲むことでキャプションを付けることができる。

```typst
#set text(font: "Noto Serif CJK JP", lang: "jp")

#figure(
  table(
    columns: (auto, auto, auto),
    align: (right, center, left),
    table.header([*日付*], [*曜日*], [*内容*]),
    [2025年5月3日], [土], [憲法記念日],
    [2025年5月4日], [日], [みどりの日],
    [2025年5月5日], [月], [こどもの日],
    [2025年5月6日], [火], [休日],
  ),
  supplement: [表],
  caption: [2025年5月の祝日・休日],
) <holiday>

@holiday は国立天文台で公開されている暦データから抽出したもの。詳しくは拙文「#link("http://localhost:1313/remark/2019/05/google-ephemeris/")[カレンダーに祝日を入れたいなら国立天文台へ行けばいいじゃない]」を参照のこと。
```

{{< fig-img src="./pdf-sample-table-02.png" title="表 (2)" link="./pdf-sample-table-02.png" width="713" >}}

表が中央寄せになった。
`<holiday>` および `@holiday` は参照構造を表すもの。
[`link`] 関数を含む参照については後日に改めて扱うことにする。

キャプションを表の上側に表示させることもできる。

```typst {hl_lines=["14-17"]}
#set text(font: "Noto Serif CJK JP", lang: "jp")

#figure(
  table(
    columns: (auto, auto, auto),
    align: (right, center, left),
    table.header([*日付*], [*曜日*], [*内容*]),
    [2025年5月3日], [土], [憲法記念日],
    [2025年5月4日], [日], [みどりの日],
    [2025年5月5日], [月], [こどもの日],
    [2025年5月6日], [火], [休日],
  ),
  supplement: [表],
  caption: figure.caption(
    position: top,
    [2025年5月の祝日・休日]
  ),
)
```

{{< fig-img src="./pdf-sample-table-03.png" title="表 (3)" link="./pdf-sample-table-03.png" width="713" >}}

あるいは最初からルール化してしまうか。

```typst
#show figure.where(
  kind: table
): set figure.caption(position: top)
```

表に関してはかなり複雑な表現が可能なようだ。
[公式ドキュメント][Typst Documentation]の “[Table guide]” も参考にどうぞ。

図に関しては [Typst] の機能を使って[作画][Visualize]もできるみたいだけど，今回は [`image`] 関数を使って既存の画像ファイルを埋め込んでみた。

```typst
#set text(font: "Noto Serif CJK JP", lang: "jp")

#figure(
  image("lake-shinjiko.jpg", width: 60%),
  supplement: [図],
  caption: "チャリで来た",
)
```

{{< fig-img src="./pdf-sample-image-01.png" title="図" link="./pdf-sample-image-01.png" width="713" >}}

これも [`figure`] 関数で囲んでいる。

データフォーマットとしては `"png"`, `"jpg"`, `"gif"`, `"svg"` のみ対応しているようだ。
特に `format` を指定しなくても自動で認識してくれるらしい。
個人的には（$\mathrm{\LaTeX}$ からの移行を考えるなら） EPS や PDF も対応して欲しいところではある。

...これでよく使う文書要素に関しては一通り調べたかな。
今回はここまで。

## ブックマーク

ブックマークは「[Typst に関するブックマーク]({{< relref "./0-bookmark.md" >}})」へ移動しました。

[Typst]: https://typst.app/ "Typst: Compose papers faster"
[Typst Documentation]: https://typst.app/docs/ "Typst Documentation"
[Tutorial]: https://typst.app/docs/tutorial "Tutorial – Typst Documentation"
[Table guide]: https://typst.app/docs/guides/table-guide/ "Table guide – Typst Documentation"
[Visualize]: https://typst.app/docs/reference/visualize/ "Visualize – Typst Documentation"
[`image`]: https://typst.app/docs/reference/visualize/image/ "Image Function – Typst Documentation"
[`linebreak`]: https://typst.app/docs/reference/text/linebreak/ "Line Break Function – Typst Documentation"
[`heading`]: https://typst.app/docs/reference/model/heading/ "Heading Function – Typst Documentation"
[`par`]: https://typst.app/docs/reference/model/par/ "Paragraph Function – Typst Documentation"
[`quote`]: https://typst.app/docs/reference/model/quote/ "Quote Function – Typst Documentation"
[`outline`]: https://typst.app/docs/reference/model/outline/ "Outline Function – Typst Documentation"
[`numbering`]: https://typst.app/docs/reference/model/numbering/ "Numbering Function – Typst Documentation"
[`enum`]: https://typst.app/docs/reference/model/enum/ "Numbered List Function – Typst Documentation"
[`list`]: https://typst.app/docs/reference/model/list/ "Bullet List Function – Typst Documentation"
[`table`]: https://typst.app/docs/reference/model/table/ "Table Function – Typst Documentation"
[`figure`]: https://typst.app/docs/reference/model/figure/ "Figure Function – Typst Documentation"
[`link`]: https://typst.app/docs/reference/model/link/ "Link Function – Typst Documentation"
[`counter`]: https://typst.app/docs/reference/introspection/counter/ "Counter Type – Typst Documentation"
[`text`]: https://typst.app/docs/reference/text/text/ "Text Function – Typst Documentation"
[`raw`]: https://typst.app/docs/reference/text/raw/ "Raw Text / Code Function – Typst Documentation"

[前回]: {{< relref "./2-setup-typst-in-local-machine.md" >}} "ローカルで Typst 環境を整える"

## 参考文献

{{% review-paapi "B0DPXBNTRS" %}} <!-- Typst完全入門-->
{{% review-paapi "4297138891" %}} <!-- ［改訂第9版］LaTeX美文書作成入門 -->
{{% review-paapi "B071CWZ2NM" %}} <!-- 熱学思想の史的展開 山本義隆 -->
