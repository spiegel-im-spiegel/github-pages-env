+++
title = "ちょこっと MathJax： 初期設定"
date =  "2017-09-27T20:02:09+09:00"
description = "この記事ではまず Web ページ上で MathJax が動くところまで説明していこう。"
tags = [ "math", "tex", "mathjax", "javascript", "blog", "site" ]

[scripts]
  mathjax = true
  mermaidjs = false
+++

ちょっと思いついたので，これから何回かに分けて [MathJax] について簡単に紹介していきたいと思う。

1. [ちょこっと MathJax： 初期設定]({{< ref "/remark/2017/09/getting-started-mathjax-1.md" >}}) ← イマココ
2. [ちょこっと MathJax： 基本的な数式表現]({{< ref "/remark/2017/09/getting-started-mathjax-2.md" >}})
3. [ちょこっと MathJax： インライン数式と別行立て数式]({{< ref "/remark/2017/10/getting-started-mathjax-3.md" >}})
4. [ちょこっと MathJax 番外編： mathcomp パッケージの代替え]({{< ref "/remark/2017/12/mathcomp-in-mathjax.md" >}})

[MathJax] は Web ブラウザ上で数学論文等でも使える高品質な数式表現を行うための JavaScript パッケージで GitHub にリポジトリがある。

- [MathJax リポジトリ](https://github.com/mathjax)

数式表現として $\mathrm{\TeX}$ 記法[^mj0] が使えるのが特徴で，たとえば HTML ソースに

[^mj0]: 厳密には $\mathrm{\TeX}$ 記法ではなく $\mathrm{\LaTeX}$ 記法である。が，ここでは両者を区別することにあまり意味が無いので「$\mathrm{\TeX}$ 記法」で通すことにする。

```html
エネルギーと質量には \( E=mc^2 \) の関係がある。
```

と書くとブラウザ側では

{{< div-box >}}
エネルギーと質量には \( E=mc^2 \) の関係がある。
{{< /div-box >}}

と適切に表示してくれる[^mj1]。

[^mj1]: `$E=mc^2$` という入力に対して $E=mc^2$ と，各文字間を適切に空けたり詰めたりしてくれるのがお分かりだろうか。このように $\mathrm{\TeX}$ では数式を半自動的かつ適切に「組版」してくれるのが特徴である。ただし万能ではない。

この記事ではまず Web ページ上で [MathJax] が動くところまでを説明していこう。
数式の書き方については[次回]以降に解説していく予定である。

## [MathJax] の組み込み {#install}

[MathJax] は v3 より完全に node.js ベースでの開発になった。
したがってサーバ側に組み込むこともできる。
今回は Web ページごとにクライアント側の JavaScript として組み込む方法を紹介する。

といっても組み込み自体は簡単で [MathJax] は CDN (Content Delivery Network) で配布されているので HTML の `<head>` 要素内に以下の行を追加するだけである。

```html
<script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-mml-chtml.js"></script>
```

[MathJax] v3 の特定バージョンを指定するには以下のようにバージョンを明記する。

```html
<script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3.1.2/es5/tex-mml-chtml.js"></script>
```

なお 2020-09-25 時点の最新バージョンは 3.1.2 である。

[MathJax] v3 は，そのままでは IE (Internet Explorer) に対応していない。
ただし IE11 に対応するのであれば，直前に1行足して

```html {hl_lines=[1]}
<script src="https://polyfill.io/v3/polyfill.min.js?features=es6"></script>
<script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-mml-chtml.js"></script>
```

などとすればいいらしい。
ちなみに IE11 より前は（Microsoft も [MathJax] v3 も）既にサポート外なのでご注意を。

<!-- MathJax v3 ではメニューの多言語化はサポートしていないっぽい？
さらにパラメータ部に `locale=ja` を追加すると，数式部分で表示されるコンテキスト・メニューが日本語になる。

{{< fig-img src="https://photo.baldanders.info/flickr/image/37316621442_m.png" title="MathJax: context menu" link="https://photo.baldanders.info/flickr/37316621442/" >}}

ブラボー！
-->

CDN から利用する場合，プロトコルは HTTPS のみで HTTP を明示的に指定しても HTTPS にリダイレクトされるようだ。
最近のブラウザは HTTP と HTTPS が混在するページでは（セキュリティの関係で）上手く表示できない場合があるので注意が必要である。
どうしても HTTP を使いたいなら CDN を使わず自前で環境を用意するほうがいいだろう。

組み込む JavaScript は `tex-mml-chtml.js` 以外に以下のものがある。

| JavaScript ファイル | 内容                                                  |
| ------------------- | ----------------------------------------------------- |
| `tex-chtml.js`      | 入力：$\mathrm{\TeX}$ 記法 , 出力： HTML              |
| `tex-chtml-full.js` | 入力：$\mathrm{\TeX}$ 記法（フル機能）, 出力： HTML  |
| `tex-svg.js`        | 入力：$\mathrm{\TeX}$ 記法 , 出力： SVG               |
| `tex-svg-full.js`   | 入力：$\mathrm{\TeX}$ 記法（フル機能）, 出力： SVG   |
| `tex-mml-chtml.js`  | 入力：$\mathrm{\TeX}$ 記法または MathML , 出力： HTML |
| `tex-mml-svg.js`    | 入力：$\mathrm{\TeX}$ 記法または MathML , 出力： SVG  |
| `mml-chtml.js`      | 入力：MathML 記法 , 出力： HTML                       |
| `mml-svg.js`        | 入力：MathML , 出力： SVG                             |

MathML による入力は本記事では割愛する。

{{< div-box type="markdown" >}}
### 2020-09-25 追記

[MathJax] を CDN から読み込む際に Chrome や Edge だと上手く行かないらしい。

- [Texによる数式表現36～MathJax, KaTeXのトラブル - つれづれなる備忘録](https://atatat.hatenablog.com/entry/tex36_mathjax_err)
- [Texによる数式表現37～MathJax, KaTeX表示トラブルの要因・解決 - つれづれなる備忘録](https://atatat.hatenablog.com/entry/tex37_katex_resolve)
- [Texによる数式表現38～MathJaxの表示トラブル解決法 - つれづれなる備忘録](https://atatat.hatenablog.com/entry/tex38_mathjax_resolve)

証明書云々ってこれのことかなぁ。

- [9月1日から、398日間を超えるSSL/TLS証明書は信頼性を失うため要注意 | マイナビニュース](https://news.mynavi.jp/article/20200819-1233144/)
- [2020年9月よりAppleがSSL証明書の有効期間を13か月に短縮！詳細や対策とは？ | さくらのSSL](https://ssl.sakura.ad.jp/column/safari-shortening/)

でも別に期限切れてないし期間も1年なので問題ないよなぁ。
むしろ `cdnjs.cloudflare.com` の証明書のほうが（今は問題ないが）将来的にヤバくね？ って感じなのだが（笑）

実は自マシンでは既に Chrome を捨てていて（つか Ubuntu では Chrome は既定で入ってない）スマホかタブレットの Chrome で確認するしかないんだけど，そっちは問題なく見れてるんだよなぁ？

対策としては，以下のように `cdnjs.cloudflare.com` で指定すればいいらしい。

```html
<script id="MathJax-script" async src="https://cdnjs.cloudflare.com/ajax/libs/mathjax/3.1.2/es5/tex-mml-chtml.js"></script>
```

強いて言えば `cdn.jsdelivr.net` の証明書を発行しているのが Cloudflare 社の中間 CA ってとこだが，それってブラウザ側の証明書ストアの問題なんじゃ...

まぁとにかく，現象が確認できない以上[本ブログ]では対応しないけど，見れない人が頻出するようなら考えます（笑）

[本ブログ]: {{< rlnk "/" >}} "text.Baldanders.info"
[MathJax]: https://www.mathjax.org/ "MathJax | Beautiful math in all browsers."
{{< /div-box >}}

## [MathJax] のオプション {#options}

[MathJax] にはいくつかのオプションを設定できる。
オプションの設定には  `<head>` 要素内に以下のように `MathJax` インスタンスを作成する（スクリプトの順番に注意）。

```html {hl_lines=["1-3"]}
<script>
MathJax = { ... };
</script>
<script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-mml-chtml.js"></script>
```

`{ ... }` の部分に具体的なオプションを記述していく。
全部を説明するのは大変なので，よく使いそうなものを幾つか紹介しよう。
なお，[最後の節]({{< relref "#mysetting" >}})に[本ブログ]におけるオプションの設定例を挙げているので，以降の解説がウザい方は丸写しでも OK です（笑）

（次節以降に出てくる「インライン数式」および「別行立て数式」については[第3回]で説明する）

### [TeX Input Processor Options](https://docs.mathjax.org/en/latest/options/input/tex.html "TeX Input Processor Options — MathJax 3.0 documentation")

設定項目と既定値は以下の通り。

```js
MathJax = {
  tex: {
    inlineMath: [              // start/end delimiter pairs for in-line math
      ['\\(', '\\)']
    ],
    displayMath: [             // start/end delimiter pairs for display math
      ['$$', '$$'],
      ['\\[', '\\]']
    ],
    processEscapes: true,      // use \$ to produce a literal dollar sign
    processEnvironments: true, // process \begin{xxx}...\end{xxx} outside math mode
    processRefs: true,         // process \ref{...} outside of math mode
    digits: /^(?:[0-9]+(?:\{,\}[0-9]{3})*(?:\.[0-9]*)?|\.[0-9]+)/,
                               // pattern for recognizing numbers
    tags: 'none',              // or 'ams' or 'all'
    tagSide: 'right',          // side for \tag macros
    tagIndent: '0.8em',        // amount to indent tags
    useLabelIds: true,         // use label name rather than tag for ids
    multlineWidth: '85%',      // width of multline environment
    maxMacros: 1000,           // maximum number of macro substitutions per expression
    maxBuffer: 5 * 1024,       // maximum size for the internal TeX string (5K)
    baseURL:                   // URL for use with links to tags (when there is a <base> tag in effect)
       (document.getElementsByTagName('base').length === 0) ?
        '' : String(document.location).replace(/#.*$/, ''))
  }
};
```

`inlineMath` はインライン数式の開始・終了デリミタを指定する。
複数列挙できる。

$\mathrm{\LaTeX}$ と同じく `$...$` 記述を有効にしたいのであれば

```js
inlineMath: [
  ['$', '$'],
  ['\\(', '\\)']
],
```

などとする。
これで

```html
エネルギーと質量には $E=mc^2$ の関係がある。
```
{{< div-box >}}
エネルギーと質量には $E=mc^2$ の関係がある。
{{< /div-box >}}

と記述できる。

`processEscapes` を `true` にすると（既定値），上述の数式開始・終了デリミタを `\` 記号でエスケープする[^esc1]。
たとえば `$` 文字を表示する場合には `\$` と記述すればよい。

[^esc1]: `processEscapes` オプションを有効にすると `\(...\)` までエスケープされてただの `(...)` になってしまうので注意すること。というか `processEscapes` オプションを有効にするなら `\(...\)` は使わないほうがいいかも。また `processEscapes` オプションはパラグラフ `<p>...</p>` の中でのみ効いているようだ。

`tags` で別行立て数式の採番を制御する。
規定値の `"none"` では自動採番が無効になっている。
ページ内の全ての別行立て数式に対して自動採番を有効にする場合は `"all"` をセットする。
`"ams"` をセットした場合の動作については[第3回]を参照のこと。

`macros` 項目を追加して自作のマクロを組み込むこともできる。
こんな感じ[^ssqrt1]。

[^ssqrt1]: `ssqrt` マクロについては[第3回]で紹介している。

```js
macros: {
  ssqrt: ['\\sqrt{\\smash[b]{\\mathstrut #1}}', 1]
}
```

これで

```html
平方根の高さを揃えるには \mathstrut と \smash コマンドを使って $\ssqrt{g}$ と $\ssqrt{h}$ のように表示できる。
```

{{< div-box >}}
平方根の高さを揃えるには <code>\mathstrut</code> と <code>\smash</code> コマンドを使って $\ssqrt{g}$ と $\ssqrt{h}$ のように表示できる。
{{< /div-box >}}

のように使うことができる。

### [CommonHTML Output Processor Options](https://docs.mathjax.org/en/latest/options/output/chtml.html "CommonHTML Output Processor Options — MathJax 3.0 documentation")

設定項目と既定値は以下の通り。

```js
MathJax = {
  chtml: {
    scale: 1,                      // global scaling factor for all expressions
    minScale: .5,                  // smallest scaling factor to use
    matchFontHeight: true,         // true to match ex-height of surrounding font
    mtextInheritFont: false,       // true to make mtext elements use surrounding font
    merrorInheritFont: true,       // true to make merror text use surrounding font
    mathmlSpacing: false,          // true for MathML spacing rules, false for TeX rules
    skipAttributes: {},            // RFDa and other attributes NOT to copy to the output
    exFactor: .5,                  // default size of ex in em units
    displayAlign: 'center',        // default for indentalign when set to 'auto'
    displayIndent: '0',            // default for indentshift when set to 'auto'
    fontURL: '[mathjax]/components/output/chtml/fonts/woff-v2',   // The URL where the fonts are found
    adaptiveCSS: true              // true means only produce CSS that is used in the processed equations
  }
};
```

`matchFontHeight` が `true` であれば本文の文字の大きさにマッチするよう数式の文字の高さを調節してくれるが，本文が日本語だと却ってバランスが悪いようだ。
したがって `false` にしておくのがお薦めである。

<!-- MathJax v3 で mtextFontInherit の挙動が変わった？ よーわからんので，とりあえずコメントアウトしておく
`mtextFontInherit` は `\text` コマンドで囲まれた文字列の組版規則と書体を指定する。
false なら数式の規則のままだが true であれば数式の周囲の地文（大抵は本文）の組版規則[^rl1]と書体が継承される。
既定値は false。

[^rl1]: 「地文の組版規則」とは要するに HTML/CSS の規則ということだが，`\text` コマンドで囲まれた部分は HTML の要素タグ等（`<code>` タグ等）は使えないようだ。試しにやってみたがエラーになってしまう。

たとえば `mtextFontInherit` を true にして

```html
\begin{alignat*}{2}
    (a+b)^2 &= a^2 + 2ab + b^2 & \qquad & \text{展開する} \\
            &= a(a + 2b) + b^2 &        & \text{a でくくる}
\end{alignat*}
```

と記述した場合は（『[LaTeX2ε美文書作成入門]』より引用）

{{< div-box >}}
\begin{alignat*}{2}
    (a+b)^2 &= a^2 + 2ab + b^2 & \qquad & \text{展開する} \\
            &= a(a + 2b) + b^2 &        & \text{a でくくる}
\end{alignat*}
{{< /div-box >}}

と表示される。
「`a でくくる`」の a が数式用の書体でないことに注意。
さらに「`a でくくる`」を「`$a$ でくくる`」とすると

```html
\begin{alignat*}{2}
    (a+b)^2 &= a^2 + 2ab + b^2 & \qquad & \text{展開する} \\
            &= a(a + 2b) + b^2 &        & \text{$a$ でくくる}
\end{alignat*}
```

{{< div-box >}}
\begin{alignat*}{2}
    (a+b)^2 &= a^2 + 2ab + b^2 & \qquad & \text{展開する} \\
            &= a(a + 2b) + b^2 &        & \text{$a$ でくくる}
\end{alignat*}
{{< /div-box >}}

と $a$ が数式用の書体になる。
なお，日本語（和文）部分は数式内でも `\text` コマンドの有無に関係なく影響を受けない。
-->

`displayAlign` は別行立て数式の位置を何処に寄せるか指定する。
左寄せ（`"left"`），右寄せ（`"right"`），中央寄せ（`"center"`）を指定できる。
中央寄せ以外のときは `displayIndent` でインデント幅を指定できる。

たとえば左寄せで2文字分インデントさせた場合は

```js
MathJax = {
  chtml: {
    displayAlign: 'left',
    displayIndent: '2em'
  }
};
```

以下のように表示される。

{{< div-box >}}
エネルギーと質量には $$E=mc^2$$ の関係がある。
{{< /div-box >}}

### [SVG Output Processor Options](https://docs.mathjax.org/en/latest/options/output/svg.html "SVG Output Processor Options — MathJax 3.0 documentation")

設定項目と既定値は以下の通り。

```js
MathJax = {
  svg: {
    scale: 1,                      // global scaling factor for all expressions
    minScale: .5,                  // smallest scaling factor to use
    matchFontHeight: true,         // true to match ex-height of surrounding font
    mtextInheritFont: false,       // true to make mtext elements use surrounding font
    merrorInheritFont: true,       // true to make merror text use surrounding font
    mathmlSpacing: false,          // true for MathML spacing rules, false for TeX rules
    skipAttributes: {},            // RFDa and other attributes NOT to copy to the output
    exFactor: .5,                  // default size of ex in em units
    displayAlign: 'center',        // default for indentalign when set to 'auto'
    displayIndent: '0',            // default for indentshift when set to 'auto'
    fontCache: 'local',            // or 'global' or 'none'
    localID: null,                 // ID to use for local font cache (for single equation processing)
    internalSpeechTitles: true,    // insert <title> tags with speech content
    titleID: 0                     // initial id number to use for aria-labeledby titles
  }
};
```

内容については前節とほぼ同じなので割愛する。
なお `matchFontHeight` 項目については `false` にしても日本語の文章と上手くマッチしない。
残念。

### 機能の拡張

[`physics`]: https://docs.mathjax.org/en/latest/input/tex/extensions/physics.html "physics — MathJax 3.0 documentation"
[`colorV2`]: https://docs.mathjax.org/en/latest/input/tex/extensions/colorV2.html "colorV2 — MathJax 3.0 documentation"

たとえば [`physics`] 拡張を組み込む場合は以下のように記述する。

```js
MathJax = {
  loader: {load: ['[tex]/physics']},
  tex: {
    packages: {
      '[+]': ['physics']
    },
    ...
  }
};
```

組み込み可能な拡張機能については以下のページを参照のこと。

- [The TeX/LaTeX Extension List — MathJax 3.0 documentation](https://docs.mathjax.org/en/latest/input/tex/extensions/index.html)

ただし [`physics`], [`colorV2`] 以外は標準で組み込まれているようで[^full1]，たとえば

[^full1]: `tex-chtml-full.js` でフル機能を組み込んだ場合は [`physics`], [`colorV2`] も組み込まれるようだ。

```html
経済成長と $\ce{CO2}$ 排出量は比例しなくなっている。
```

{{< div-box >}}
経済成長と $\ce{CO2}$ 排出量は比例しなくなっている。
{{< /div-box >}}

なんてな感じに書くことができる[^co2]。

[^co2]: 「[経済成長とCO2排出量は「比例しなくなっている」：IEA報告書](http://wired.jp/2017/03/29/global-carbon-emissions/ "経済成長とCO2排出量は「比例しなくなっている」：IEA報告書｜WIRED.jp")」より。

### Web フォントの指定

- [MathJax Font Support — MathJax 3.0 documentation](https://docs.mathjax.org/en/latest/output/fonts.html)

今のところ [MathJax] v3 の CDN では TeX フォントしか対応していない。
将来バージョンで対応するとあるが，フォント群を指定する仕組みだけは用意されているようだ。

```js
MathJax = {
  chtml: {
    fontURL: '[mathjax]/components/output/chtml/fonts/woff-v2' // The URL where the fonts are found
  }
};
```

ただ [MathJax] v3 で利用可能なフォント群を用意するのは（今のところ）簡単ではなさそうなので「今後に期待」といったところだろうか。


## このサイトでの設定例 {#mysetting}

以上を踏まえて，[本ブログ]における [MathJax] オプションの設定内容を以下に示す。

```html
<script>
MathJax = {
  tex: {
    inlineMath: [['$', '$'], ['\\(', '\\)']],
	processEscapes: true,
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
    matchFontHeight: false,
    displayAlign: "left",
    displayIndent: "2em"
  }
};
</script>
<script id="MathJax-script" async src="//cdn.jsdelivr.net/npm/mathjax@3/es5/tex-chtml.js"></script>
```

これでようやく準備が整った。

## ブックマーク {#bookmark}

- [MathJax Documentation — MathJax 3.0 documentation](https://docs.mathjax.org/en/latest/)
- [MathJaxによる数式表示](https://oku.edu.mie-u.ac.jp/~okumura/javascript/mathjax.html)
- [MathJaxの使い方](http://gilbert.ninja-web.net/math/mathjax1.html)
    - [MathJaxの使い方〈化学編〉](http://gilbert.ninja-web.net/math/mathjax3.html)
- [世界標準が期待される数式用フォント「STIX Fonts」 - 窓の杜](http://forest.watch.impress.co.jp/docs/news/373370.html)

- [MathJax v3 がリリースされていた]({{< ref "/release/2019/09/mathjax-v3-is-released.md" >}})

[本ブログ]: {{< rlnk "/" >}} "text.Baldanders.info"
[MathJax]: https://www.mathjax.org/ "MathJax | Beautiful math in all browsers."
[次回]: {{< ref "/remark/2017/09/getting-started-mathjax-2.md" >}} "ちょこっと MathJax： 基本的な数式表現"
[第3回]: {{< ref "/remark/2017/10/getting-started-mathjax-3.md" >}} "ちょこっと MathJax： インライン数式と別行立て数式"
[LaTeX2ε美文書作成入門]: https://www.amazon.co.jp/dp/4774187054?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Amazon | [改訂第7版]LaTeX2ε美文書作成入門 | 奥村 晴彦, 黒木 裕介 通販"

## 参考図書 {#books}

{{% review-paapi "4774187054" %}} <!-- [改訂第7版]LaTeX2ε美文書作成入門 -->
