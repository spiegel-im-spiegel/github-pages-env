+++
title = "ちょこっと MathJax 番外編： V4 へのアップグレード"
date =  "2025-09-01T17:55:24+09:00"
description = "URL に注意 / 利用可能なフォントが増えた / MathJax 設定例"
image = "/images/attention/kitten.jpg"
tags = [ "math", "tex", "mathjax", "javascript" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

1. [ちょこっと MathJax： 初期設定]({{< ref "/remark/2017/09/getting-started-mathjax-1.md" >}})
2. [ちょこっと MathJax： 基本的な数式表現]({{< ref "/remark/2017/09/getting-started-mathjax-2.md" >}})
3. [ちょこっと MathJax： インライン数式と別行立て数式]({{< ref "/remark/2017/10/getting-started-mathjax-3.md" >}})
4. [ちょこっと MathJax 番外編： mathcomp パッケージの代替え]({{< ref "/remark/2017/12/mathcomp-in-mathjax.md" >}})
5. [ちょこっと MathJax 番外編： V4 へのアップグレード]({{< ref "/remark/2025/09/upgrade-mathjax-v4.md" >}}) ← イマココ

先月の話だが [MathJax] の v4 がリリースされていた。

- [MathJax v4.0.0 now available | MathJax](https://www.mathjax.org/MathJax-v4.0.0-available/)

主な内容は以下の通り：

{{< fig-quote type="markdown" title="MathJax v4.0.0 now available" link="https://www.mathjax.org/MathJax-v4.0.0-available/" lang="en" >}}
- Extended font support (11 new fonts)
- A new default font with much greater character coverage.
- Line-breaking support (for in-line as well as displayed equations)
- An updated expression explorer that is now on by default.
- Separation of speech generation into a webworker thread for better performance.
- Support for HTML embedded in TeX and MathML expressions.
- Availability as ES6 modules as well as CommonJS modules.
- New TeX extensions, including the begingroup extension from v2 (with more features), and several font-related extensions.
- Inclusion of the textmacros extension in all combined component files.
- Updated mathtools package to include the changes from 2022 and 2024.
- Improvements in the baseline alignment of text in CHTML output in WebKit-based browsers.
- Improved stretchy delimiters in CHTML output.
- More promise-based conversion and typesetting calls at the document level.
{{< /fig-quote >}}

本来なら [Release]({{< rlnk "/release/" >}}) セクションに書くべきなんだろうけど，今回は「ちょこっと MathJax」シリーズの番外編として主にアップグレード時の注意点をいくつか記しておく。
といっても v3 から v4 へのアップグレードで考慮すべき点は多くない。
詳しくは以下を参照のこと：

- [Upgrading from v3 to v4 — MathJax 4.0 documentation](https://docs.mathjax.org/en/latest/upgrading/v3.html)

## URL に注意

v4 では CDN の URL は以下の通りになっている。

```html
<script defer src="https://cdn.jsdelivr.net/npm/mathjax@4/tex-mml-chtml.js"></script>
```

v4 は ES6 モジュールとして実装されているため，パスから `es5` が除かれている。
あと `async` から `defer` に変わっている。
`id` 指定もなくなってるな。

バージョンを厳密に指示したい場合は

```html
<script defer src="https://cdn.jsdelivr.net/npm/mathjax@4.0.0/tex-mml-chtml.js"></script>
```

といった感じに指定する。

利用可能なコンポーネントは以下の通り。

| Component | Input | Output | Font|
| --- | --- | --- | :---: |
| `tex-chtml` | TeX | CHTML | {{< emoji "チェック" >}} |
| `tex-svg` | TeX | SVG | {{< emoji "チェック" >}} |
| `tex-mml-chtml` | TeX, MathML | CHTML | {{< emoji "チェック" >}} |
| `tex-mml-svg` | TeX, MathML | SVG | {{< emoji "チェック" >}} |
| `mml-chtml` | MathML | CHTML | {{< emoji "チェック" >}} |
| `mml-svg` | MathML | SVG | {{< emoji "チェック" >}} |
| `tex-chtml-nofont` | TeX | CHTML |  |
| `tex-svg-nofont` | TeX | SVG |  |
| `tex-mml-chtml-nofont` | TeX, MathML | CHTML |  |
| `tex-mml-svg-nofont` | TeX, MathML | SVG |  |
| `mml-chtml-nofont` | MathML | CHTML |  |
| `mml-svg-nofont` | MathML | SVG |  |

v3 にあった `-full` 付きのコンポーネントはなくなった。
その代わり [`autoload`] 拡張機能で必要な[機能](https://docs.mathjax.org/en/latest/input/tex/extensions/ "The TeX/LaTeX Extension List — MathJax 4.0 documentation")を自動的に読み込むため，通常は気にする必要はない[^a1]。

[^a1]: `ams`, `bbm`, `bboldx`, `begingroup`, `cases`, `centernot`, `colortbl`, `dsfont`, `empheq`, `mathtools`, `physics`, `require`, `setoptions`, `textcomp`, `textmacros`, `units`, `upgreek` 各機能は [`autoload`] の対象にならない。新しく追加される拡張機能は [`autoload`] の対象にならない場合があるそうで，将来的に [`autoload`] が廃止される可能性もあるらしい。どうなるやら。

## 利用可能なフォントが増えた

v4 では11種類のフォントが利用可能になった。

{{< fig-quote class="nobox" type="markdown" title="MathJax Font Support — MathJax 4.0 documentation" link="https://docs.mathjax.org/en/latest/output/fonts.html" lang="en" >}}
| Font Name | Original Source |
| --- | --- |
| `mathjax-newcm` | Based on New Computer Modern (now the default font) |
| `mathjax-asana` | A version of the Asana-Math font |
| `mathjax-bonum` | A version of the Gyre Bonum font |
| `mathjax-dejavu` | A version of the Gyre DejaVu font |
| `mathjax-fira` | A version of the Fira and Fira-Math fonts |
| `mathjax-modern` | A version of Latin-Modern |
| `mathjax-pagella` | A version of the Gyre Pagella font |
| `mathjax-schola` | A version of the Gyre Schola font |
| `mathjax-stix2` | A version of the STIX2 font |
| `mathjax-termes` | A version of the Gyre Termes font |
| `mathjax-tex` | The original MathJax TeX font |
{{< /fig-quote >}}

このうち `mathjax-newcm` については前節で {{< emoji "チェック" >}} が付いているコンポーネントに含まれている。
`mathjax-newcm` 以外のフォントを利用したい場合は `-nofont` が付いているコンポーネントを指定した上で，以下のように設定する。

```html
<script>
MathJax = {
  output: {
    font: 'mathjax-stix2'
  }
};
</script>
<script defer src="https://cdn.jsdelivr.net/npm/mathjax@4/tex-mml-chtml-nofont.js"></script>
```

ちなみに `mathjax-tex` は v3 の既定フォントなので， v3 の環境に合わせたい場合はこれを指定する。

残念ながら，みんな大好き Euler フォントはない[^e1]。
また複数のフォントを組み合わせることもできなさそうだ。

[^e1]: “[Font Extensions](https://docs.mathjax.org/en/latest/output/fonts.html#font-extensions)” に Euler フォントの指定方法が書かれていたがサンプルのままでは上手く行かなかった。拡張用のフォントは自前で用意しろってことらしい。

## MathJax 設定例

設定例として，このブログサイトにおける設定を挙げておく（2025-09-01 時点）。

```html
<script>
MathJax = {
  loader: {
    load: ['[tex]/ams', '[tex]/textcomp']
  },
  tex: {
    packages: {'[+]': ['ams', 'textcomp']},
    inlineMath: {'[+]': [['$', '$']]},
    tags: 'ams',
    macros: {
      ssqrt: ['\\sqrt{\\smash[b]{\\mathstrut #1}}', 1],
      tcdegree: ['\\unicode{xb0}'],
      tccelsius: ['\\unicode{x2103}'],
      tcperthousand: ['\\unicode{x2030}'],
      tcmu: ['\\unicode{x3bc}'],
      tcohm: ['\\unicode{x3a9}']
    }
  },
  chtml: {
    displayAlign: "left",
    displayIndent: "2em"
  },
  output: {
    font: 'mathjax-stix2'
  }
};
</script>
<script defer src="//cdn.jsdelivr.net/npm/mathjax@4/tex-chtml-nofont.js"></script>
```

[MathJax]: https://www.mathjax.org/
[`autoload`]: https://docs.mathjax.org/en/latest/input/tex/extensions/autoload.html "autoload — MathJax 4.0 documentation"

## 参考図書 {#books}

{{% review-paapi "4297138891" %}} <!-- ［改訂第9版］LaTeX美文書作成入門 -->
