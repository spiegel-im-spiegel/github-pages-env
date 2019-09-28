+++
title = "MathJax v3 がリリースされていた"
date =  "2019-09-28T18:29:26+09:00"
description = "v2.7.x までとはだいぶ変わったので注意。"
image = "/images/attention/tools.png"
tags = [ "math", "tex", "mathjax", "javascript", "site" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

いや，確かに8月末にはリリースするって予告されてたけどさ。
その後のアナウンスがないから気づかなかったよ。

ちうわけで [MathJax] v3.0.0 がリリースされていた。
よーやくベータが取れたよ（笑）

ドキュメントも v3 に対応しているようだ。

- [MathJax Documentation — MathJax 3.0 documentation](https://docs.mathjax.org/en/latest/)

Web ページで [MathJax] を使えるようにするには以下のコードを記述する。

```html
<script>
MathJax = {
  ...
};
</script>
<script src="https://polyfill.io/v3/polyfill.min.js?features=es6"></script>
<script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-mml-chtml.js"></script>
```

[MathJax] v3 の特定のバージョンを指定するには

```html
<script>
MathJax = {
  ...
};
</script>
<script src="https://polyfill.io/v3/polyfill.min.js?features=es6"></script>
<script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3.0.0/es5/tex-mml-chtml.js"></script>
```

などとする。

ちなみに IE (Internet Explorer) はサポートから外れるようだ。
まぁ IE は Web ではもう「要らない子」だからね。

## [MathJax] v3 の設定

`MathJax = { ... };` の部分には [MathJax] の設定を記述する。
v2.7.x までとはフォーマットが異なるので注意。

主な設定項目は以下の通り。

### [TeX Input Processor Options](https://docs.mathjax.org/en/latest/options/input/tex.html "TeX Input Processor Options — MathJax 3.0 documentation")

設定項目と規定値は以下の通り。

```js
MathJax = {
  tex: {
    packages: ['base'],        // extensions to use
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

インライン数式を囲む記号として `$ .. $` を有効にするには `inlineMath` 項目を以下のように変更する。

```js
inlineMath: [['$', '$'], ['\\(', '\\)']]
```

別行立て数式に対して自動採番を行うには `tags` 項目に `ams` または `all` をセットする。

自作のマクロを組み込む際には `macro` 項目をセットする。

```js
macros: {
  ssqrt: ['\\sqrt{\\smash[b]{\\mathstrut #1}}', 1]
}
```

### [CommonHTML Output Processor Options](https://docs.mathjax.org/en/latest/options/output/chtml.html "CommonHTML Output Processor Options — MathJax 3.0 documentation")

設定項目と規定値は以下の通り。

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
    displayIndent: '0'             // default for indentshift when set to 'auto'
  }
};
```

日本語の文章に数式を埋め込む場合，文字サイズのバランスがよくないので `matchFontHeight` 項目を `false` にする。

別行立て数式を左寄せで表示するには `displayAlign` 項目を `left` にする。
併せて `displayIndent` 項目でインデント幅もセットするとよい。

### 拡張パッケージを組み込む

たとえば `color` という拡張パッケージを組み込むには以下のように記述する。

```js
MathJax = {
  loader: {
    load: ['[tex]/color']
  },
  tex: {
    packages: {'[+]': ['color']}
  }
};
```

ただし現時点で使えそうなパッケージはなさそう。
[mathcomp](https://ctan.org/pkg/mathcomp "CTAN: Package mathcomp") があるといいのに。

v2.7.x で外部パッケージだったものの一部は標準で組み込まれているようだ。
たとえば `mhchem.js` は明示的に組み込まなくても

```html
経済成長と $\ce{CO2}$ 排出量は比例しなくなっている。
```

> 経済成長と $\ce{CO2}$ 排出量は比例しなくなっている[^co2]。

[^co2]: 「[経済成長とCO2排出量は「比例しなくなっている」：IEA報告書](http://wired.jp/2017/03/29/global-carbon-emissions/ "経済成長とCO2排出量は「比例しなくなっている」：IEA報告書｜WIRED.jp")」より。

てな感じに書くことができる。

### このサイトの設定内容

このサイトの設定内容は今のところ以下の通り。

```js
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
```

[MathJax] v3 では利用できる Web フォントは1種類のみで他のフォントを選択することができない。
これは将来バージョンで改善するらしい。
Web フォントのサポートが組み込まればまた紹介することもあるだろう。

## ブックマーク

- [ちょこっと MathJax]({{< ref "/remark/2017/09/getting-started-mathjax-1.md" >}})

[MathJax]: https://www.mathjax.org/ "MathJax | Beautiful math in all browsers."

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%94%B9%E8%A8%82%E7%AC%AC7%E7%89%88-LaTeX2%CE%B5%E7%BE%8E%E6%96%87%E6%9B%B8%E4%BD%9C%E6%88%90%E5%85%A5%E9%96%80-%E5%A5%A5%E6%9D%91-%E6%99%B4%E5%BD%A6/dp/4774187054?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4774187054"><img src="https://images-fe.ssl-images-amazon.com/images/I/51E5K7B53aL._SL160_.jpg" width="127" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%94%B9%E8%A8%82%E7%AC%AC7%E7%89%88-LaTeX2%CE%B5%E7%BE%8E%E6%96%87%E6%9B%B8%E4%BD%9C%E6%88%90%E5%85%A5%E9%96%80-%E5%A5%A5%E6%9D%91-%E6%99%B4%E5%BD%A6/dp/4774187054?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4774187054">[改訂第7版]LaTeX2ε美文書作成入門</a></dt>
    <dd>奥村 晴彦, 黒木 裕介</dd>
    <dd>技術評論社 2017-01-24</dd>
    <dd>大型本</dd>
    <dd>4774187054 (ASIN), 9784774187051 (EAN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">ついに第7版が登場。紙の本で買って常に側に置いておくのが吉。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-09-27">2017-09-27</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-API</a>)</p>
</div>
