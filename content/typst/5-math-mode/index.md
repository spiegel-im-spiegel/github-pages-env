+++
title = "Typst の数式モード"
date =  "2025-02-26T22:34:05+09:00"
description = "LaTeX 記法とは似て非なる… / フォントの設定 / 立体か斜体か / 番号付けと参照"
image = "/images/attention/tools.png"
tags  = [ "typst", "tex", "math", "font" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

今回は今まで敢えて後回しにしていた数式モード（math mode）[^m1] の話。
パッと見は $\mathrm{\LaTeX}$ 記法に似ているのだが色々と違っている。

[^m1]: “Math mode” をどう訳すのか悩んだが（英語不得手なもので），「[非公式日本語ドキュメント](https://typst-jp.github.io/docs/ "Typstについて – Typstドキュメント日本語版")」を見ると「数学」ではなく「数式」となっていたので「数式モード」とした。

## LaTeX 記法とは似て非なる...

たとえば $\mathrm{\LaTeX}$ 記法では

```latex
エネルギーと質量には $E=mc^2$ の関係がある。
```

と書けば

{{< div-box type="markdown" >}}
エネルギーと質量には $E=mc^2$ の関係がある。
{{< /div-box >}}

という感じに組版される。
これをそのまま [Typst] で組版しようとすると

```text
$ typst compile math.typ
error: unknown variable: mc
  ┌─ math.typ:1:14
  │
1 │ エネルギーと質量には $E=mc^2$ の関係がある。
  │                         ^^
  │
  = hint: if you meant to display multiple letters as is, try adding spaces between each letter: `m c`
  = hint: or if you meant to display this as text, try placing it in quotes: `"mc"`
```

てな感じで怒られる。
どうも2つ以上の文字の綴りは何らかのシンボルと解釈されるらしい。
たとえば $\mathrm{\LaTeX}$ 記法の以下の数式は

{{< fig-quote class="nobox" type="markdown" title="『［改訂第9版］LaTeX美文書作成入門』p.221" link="https://www.amazon.co.jp/dp/4297138891?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
```latex
\[
  \left( \int_0^\infty \frac{\sin x}{\sqrt x} dx \right)^2 =
  \sum_{k=0}^\infty \frac{(2k)!}{2^{2k}(k!)^2} \frac{1}{2k+1} =
  \prod_{k=1}^\infty \frac{4k^2}{4k^2-1} =
  \frac{\pi}{2}
\]
```
{{< /fig-quote >}}

{{< fig-img src="https://photo.baldanders.info/flickr/image/36788335984_m.png" title="equation (STIX2 Math)" link="https://photo.baldanders.info/flickr/36788335984/" >}}

[Typst] 記法だと以下のようになる。

```typst
$
  (integral_0^infinity (sin x) / sqrt(x) dif x)^2 =
  sum_(k=0)^infinity (2k)! / (2^(2k) (k!)^2) 1 / (2k+1) =
  product_(k=1)^(infinity) (4k^2) / (4k^2-1) =
  pi / 2
$
```

{{< fig-img src="./math-02.png" title="equation by Typst (STIX2 Math)" link="./math-02.pdf" width="693" >}}

$\mathrm{\LaTeX}$ 記法の `{ ... }` に相当するものが [Typst] 記法だと `( ... )` になってるのが紛らわしい。
おそらく文脈で判別してるんだろうけど。
あとインライン数式も別行立て数式もいずれも `$ ... $` なのも紛らわしい。
こちらは `$E=m c^2$` と `$ E=m c^2 $` の違い。
数式の前後に空白文字や改行が入ってると別行立て数式と認識するらしい。

これは $\mathrm{\LaTeX}$ 記法に慣れている人はちょっと大変かもしれない。
[Pandoc] には $\mathrm{\LaTeX}$ 記法 → [Typst] 記法への変換機能があるそうだが，そもそも [Typst] を使うような人はわざわざ [Pandoc] を導入しないと思う。

## フォントの設定

数式モードのフォントを指定する場合は以下のように記述する。

```typst
#show math.equation: set text(font: "New Computer Modern Math")
```

既定は New Computer Modern Math。
数式モード用のフォントには通常のフォントファイルは使えないので注意。

{{< fig-quote type="markdown" title="Math fonts" link="https://typst.app/docs/reference/math/#math-fonts" lang="en" >}}
Note that only special OpenType math fonts are suitable for typesetting maths.
{{< /fig-quote >}}

試しに [Euler](https://github.com/aliftype/euler-otf "aliftype/euler-otf: An abandoned OpenType/Unicode math port of AMS Euler font") フォントで試そうと思ったが，そういやこれってグリフが全部揃ってなかったわ。
[STIX2](https://github.com/stipub/stixfonts "stipub/stixfonts: OpenType Unicode fonts for Scientific, Technical, and Mathematical texts") フォントのほうはこんな感じで上手くいった[^m2]。

[^m2]: [Typst] のクラウドサービスではフォントアップロードなしに [STIX]2 フォントが使える。というか [TeX Live] に収録されている TTF/OTF ファイルは全部使えるっぽい。

```typst
#set text(font: (
  (
    name: "STIX Two Text",
    covers: "latin-in-cjk",
  ),
  "Noto Serif CJK JP"
), lang: "ja")
#show math.equation: set text(font: ("STIX Two Math"))

アルベルト・アインシュタインは1905年に発表した特殊相対性理論を通して以下の関係を導き出した。

$
  E = m c^2 "（エネルギーと質量の等価性）"
$
```

{{< fig-img src="./math-01.png" title="数式表現（STIX2 フォント）" link="./math-01.pdf" width="693" >}}

ただし，上の例では数式モードに日本語が入った場合に問題が出る。
どうも最初の `#set text( ... )` のルール設定が数式モードでは適用されないみたいで

{{< fig-img src="./math-01-property.png" title="数式表現（STIX2 フォント）プロパティ" link="./math-01-property.png" >}}

という感じで予期しないフォントが埋め込まれてしまった。

これを回避するには

```typst
#show math.equation: set text(font: ("New Computer Modern Math", "NOTO Serif CJK JP"))
```

などと指定する。
これで `New Computer Modern Math` ＞ `NOTO Serif CJK JP` という優先順位でフォントを選択する。

{{< fig-img src="./math-01b-property.png" title="数式表現（STIX2 フォント）プロパティ" link="./math-01b-property.png" >}}

ありゃ。
Medium フォントが選択されてる。
なんで？ うーん。
とりあえずこれは置いておこう。

## 立体か斜体か

数式の場合は基本的にイタリック体または斜体で記述するのが原則だが，以下の例外がある。

{{< fig-quote type="markdown" title="『［改訂第9版］LaTeX美文書作成入門』p.94" link="https://www.amazon.co.jp/dp/4297138891?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
- 数字はローマン体にする（$\mathit{3.14}$ ではなく $3.14$）
- 複数文字からなる名前はローマン体にする（$sin\\,x$ ではなく $\sin x$）
- 単位記号はローマン体にする（$3\\,m$ ではなく $3\\,\mathrm{m}$）
{{< /fig-quote >}}

ただ立体（upright）か斜体（slant）[^s1] かの規準は学会等によっても違うようで，たとえば積分等に出てくる `dx` とかでも斜体の $dx$ とするか立体の $\mathrm{d}x$ とするか流儀があるらしい。
ちなみに [Typst] では積分の `dx` を立体で表現するために `dif x` と記述する。

[^s1]: 厳密にはイタリック体と斜体は異なるのだが（イタリック体にも立体と斜体がある），この記事ではまとめて「斜体」と呼ぶことにする。

$\mathrm{\LaTeX}$ には `unicode-math` パッケージというのがあって，数式の書体についてはこのパッケージである程度コントロールできる。
詳しくは拙文「[数式用フォントで遊ぶ]({{< ref "/remark/2017/10/math-fonts.md" >}} "数式用フォントで遊ぶ")」を参考にどうぞ。
[Typst] には `unicode-math` パッケージに相当するものはないっぽい？

[Typst] では $\sin$ などの一般的な演算子（関数）はあらかじめ定義されていて，これらは立体で表示される。

{{< fig-quote type="markdown" title="Text Operator Function – Typst Documentation" link="https://typst.app/docs/reference/math/op/" lang="en" >}}
Typst predefines the operators `arccos`, `arcsin`, `arctan`, `arg`, `cos`, `cosh`, `cot`, `coth`, `csc`, `csch`, `ctg`, `deg`, `det`, `dim`, `exp`, `gcd`, `lcm`, `hom`, `id`, `im`, `inf`, `ker`, `lg`, `lim`, `liminf`, `limsup`, `ln`, `log`, `max`, `min`, `mod`, `Pr`, `sec`, `sech`, `sin`, `sinc`, `sinh`, `sup`, `tan`, `tanh`, `tg` and `tr`.
{{< /fig-quote >}}

この一覧にない演算子（関数）は自分で定義できる。

```typst {hl_lines=[4]}
#set text(font: "Noto Serif CJK JP", lang: "ja")
#show math.equation: set text(font: "STIX Two Math")

#let GF = math.op("GF")

位数が $q$ である有限体を $GF(q)$ などと表記します。
```

{{< fig-img src="./math-03.png" title="カスタム演算子（関数）" link="./math-03.pdf" width="651" >}}

他に $\int$ などのシンボルは “[Symbols]” のページまたは [`sym`] の項を参照のこと。
ちなみに [`emoji`] も使える。

```typst
#show math.equation: set text(font: "STIX Two Math")

$
  sum_(k=0)^infinity #text(red)[#emoji.suit.heart]^k
$
```

{{< fig-img src="./math-04.png" title="数学ガールの秘密ノート" link="./math-04.pdf" width="651" >}}

シンボルについては $\mathrm{\LaTeX}$ の `physica` パッケージに近い [`physica`] パッケージがあるらしい。

メートル（$\mathrm{m}$）などの単位も立体で表記する。
単位等の表現については $\mathrm{\LaTeX}$ の `siunitx` パッケージに近い [`metro`] または [`unify`] パッケージを使う手がある。
[`unify`] のほうが活況？


```latex
\[
  1\,\mathrm{atm} = 1.01325\,\mathrm{bar} = 101,325\,\mathrm{Pa}
\]
```

{{< fig-math class="box" >}}
\[
  1\,\mathrm{atm} = 1.01325\,\mathrm{bar} = 101,325\,\mathrm{Pa}
\]
{{< /fig-math >}}

MathJax だと泥臭い記述になるのはご容赦（MathJax 用の動く `siunitx` パッケージはないっぽい）。

```typst
#import "@preview/unify:0.7.1": qty
#show math.equation: set text(font: "STIX Two Math")

$
  qty("1", "atm") = qty("1.01325", "bar") = qty("101,325", "Pa")
$
```

{{< fig-img src="./math-05.png" title="気圧単位の換算" link="./math-05.pdf" width="651" >}}

同じような感じになったかな。
単位表記は値と単位の間のアキを調整する必要があるが [`unify`] パッケージなら上手いことやってくれているようだ。

最終手段として，強制的に立体または斜体にする場合は [`upright`] または [`italic`] 関数を使う。

```typst
#show math.equation: set text(font: "STIX Two Math")

$
  upright("Upright, not Italic") \
  "normal string" \
  italic("Italic, not Upright") \
  bold("Bold") \
  italic(bold("Bold Italic")) \
$
```

{{< fig-img src="./math-06.png" title="立体と斜体" link="./math-06.pdf" width="651" >}}

おー。
ちゃんと（単なる斜体でない）イタリック体になるんだな。

## 番号付けと参照

数式の番号付けは [`equation`] の設定で行う。

```typst
#set text(font: "Noto Serif CJK JP", lang: "ja")
#set math.equation(numbering: "(1)")
#show math.equation: set text(font: "STIX Two Math")

エネルギーと質量には以下の関係がある。

$
  E = m c^2
$
```

{{< fig-img src="./math-07.png" title="番号付け" link="./math-07.pdf" width="905" >}}

また数式には以下のように `<label_name>` でラベルを付けることができ `@label_name` で参照できる。

```typst {hl_lines=[5,9]}
#set text(font: "Noto Serif CJK JP", lang: "ja")
#set math.equation(numbering: "(1)")
#show math.equation: set text(font: "STIX Two Math")

エネルギーと質量には@eq.1 の関係がある。

$
  E = m c^2
$ <eq.1>
```

{{< fig-img src="./math-08.png" title="ラベルと参照" link="./math-08.pdf" width="905" >}}

ラベルはあくまでもマークアップモードでのみ付けることができるみたい。
数式モードの中では付けられなかった。
参照については後日に改めて扱うことにする。

## ブックマーク

ブックマークは「[Typst に関するブックマーク]({{< relref "./0-bookmark.md" >}})」にてまとめています。

[Typst]: https://typst.app/ "Typst: Compose papers faster"
[Typst Documentation]: https://typst.app/docs/ "Typst Documentation"
[Tutorial]: https://typst.app/docs/tutorial "Tutorial – Typst Documentation"
[`text`]: https://typst.app/docs/reference/text/text/ "Text Function – Typst Documentation"
[Symbols]: https://typst.app/docs/reference/symbols/ "Symbols – Typst Documentation"
[`sym`]: https://typst.app/docs/reference/symbols/sym/ "General Symbols – Typst Documentation"
[`emoji`]: https://typst.app/docs/reference/symbols/emoji/ "Emoji Symbols – Typst Documentation"
[`physica`]: https://typst.app/universe/package/physica/ "physica – Typst Universe"
[`metro`]: https://typst.app/universe/package/metro "metro – Typst Universe"
[`unify`]: https://typst.app/universe/package/unify/ "unify – Typst Universe"
[`upright`]: https://typst.app/docs/reference/math/styles#functions-upright "Styles Functions – Typst Documentation"
[`italic`]: https://typst.app/docs/reference/math/styles#functions-italic "Styles Functions – Typst Documentation"
[`equation`]: https://typst.app/docs/reference/math/equation/ "Equation Function – Typst Documentation"


[Pandoc]: https://pandoc.org/ "Pandoc"
[STIX]: https://www.stixfonts.org/ "Welcome to stixfonts | Scientific and Technical Information Exchange (STIX)"
[TeX Live]: https://www.tug.org/texlive/ "TeX Live - TeX Users Group"

## 参考文献

{{% review-paapi "B0DPXBNTRS" %}} <!-- Typst完全入門-->
{{% review-paapi "4297138891" %}} <!-- ［改訂第9版］LaTeX美文書作成入門 -->
