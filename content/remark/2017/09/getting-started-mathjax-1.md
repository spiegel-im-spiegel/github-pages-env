+++
title = "ちょこっと MathJax： 初期設定"
date =  "2017-09-27T20:02:09+09:00"
description = "この記事ではまず Web ページ上で MathJax が動くところまで説明していこう。"
tags        = [ "math", "tex", "mathjax", "blog", "site" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = true
  mermaidjs = false
+++

ちょっと思いついたので，これから何回かに分けて [MathJax] について簡単に紹介していきたいと思う。

1. [ちょこっと MathJax： 初期設定]({{< relref "remark/2017/09/getting-started-mathjax-1.md" >}}) ← イマココ

[MathJax] は Web ブラウザで数式表現を行うための JavaScript パッケージで GutHub 上で開発が行われている。

- [MathJax リポジトリ](https://github.com/mathjax)

数式表現として $\mathrm{\TeX}$ 記法[^mj0] が使えるのが特徴で，たとえば HTML ソースに

[^mj0]: 厳密には $\mathrm{\TeX}$ 記法ではなく $\mathrm{\LaTeX}$ 記法である。が，ここでは両者を区別することにあまり意味が無いので「$\mathrm{\TeX}$ 記法」で通すことにする。

```html
エネルギーと質量には $E=mc^2$ の関係がある。
```

と書くとブラウザ側では

> エネルギーと質量には $E=mc^2$ の関係がある。

と適切に表示してくれる[^mj1]。

この記事ではまず Web ページ上で [MathJax] が動くところまで説明していこう。

[^mj1]: `$E=mc^2$` という入力に対して $E=mc^2$ と，各文字間を適切に空けたり詰めたりしてくれるのがお分かりだろうか。このように $\mathrm{\TeX}$ では数式を半自動的かつ適切に「組版」してくれるのが特徴である。ただし万能ではない。

## [MathJax] のインストール

[MathJax] は JavaScript パッケージなので `<script>` タグで指定する。
[MathJax] パッケージは CDN (Content Delivery Network) で配布されているので HTML の `<head>` 要素内に

```html
<script async src="https://cdnjs.cloudflare.com/ajax/libs/mathjax/2.7.2/MathJax.js?config=TeX-AMS_CHTML&locale=ja"></script>
```

と指定すればよい（`async` を忘れずに）。
`2.7.2` の部分はバージョン番号で 2017-08-28 現在のバージョンは 2.7.2 である。
ちなみに2017年5月以降 CDN の配布場所が変わったので，古い設定のままの方は注意が必要である。

- [MathJax CDN shutting down on April 30, 2017.  Alternatives available.](https://www.mathjax.org/cdn-shutting-down/)

パラメータ部分の `config=TeX-AMS_CHTML` 部分は今のところ固定のままにしておくのが無難なようだ。
さらにパラメータとして `locale=ja` を追加すると，数式部分で表示されるコンテキスト・メニューが日本語になる。

{{< fig-img src="https://farm5.staticflickr.com/4498/37316621442_fc895ffc22_o.png" title="MathJax: context menu"  link="https://www.flickr.com/photos/spiegel/37316621442/" >}}

ブラボー！

## [MathJax] のオプション

[MathJax] にはいくつかのオプションを設定できる。
オプションの設定には `MathJax.Hub.Config()` 関数を使う。
先程の `<script>` 指定と併せて `<head>` 要素内に以下のように記述する。

```html
<script type="text/x-mathjax-config">
MathJax.Hub.Config({ ... });
</script>
<script async src="https://cdnjs.cloudflare.com/ajax/libs/mathjax/2.7.2/MathJax.js?config=TeX-AMS_CHTML&locale=ja"></script>
```

オプション指定部のメディア・タイプに `text/x-mathjax-config` を指定すること（`<script>` の順番にも注意）。
`{ ... }` のオブジェクトに具体的なオプションを指定していく。
全部を説明するのは大変なので，よく使いそうなものを幾つか紹介しよう。

### [Core Configuration Options](http://docs.mathjax.org/en/latest/options/hub.html "The Core Configuration Options — MathJax 2.7 documentation")

```html
<script type="text/x-mathjax-config">
MathJax.Hub.Config({
  displayAlign: "left",
  displayIndent: "2em"
});
</script>
```

`displayAlign` は，以下のような別行立ての場合に，数式を何処に寄せるか指定する[^eq1]。

[^eq1]: ちなみにこの数式は線形合同法の漸化式である。一番簡単な擬似乱数生成器ですね。

{{< fig-quote >}}
\[
  X_{n+1}=\left(A\times\,X_{n}+B\right)\bmod\,M
\]
{{< /fig-quote >}}

左寄せ（`left`），右寄せ（`right`），中央寄せ（`center`）を指定できる。
既定値は `center`。
また中央寄せ以外のときは `displayIndent` でインデント幅を指定できる。

### [“tex2jax” Preprocessor Options](http://docs.mathjax.org/en/latest/options/preprocessors/tex2jax.html "The tex2jax Preprocessor — MathJax 2.7 documentation")

```html
<script type="text/x-mathjax-config">
MathJax.Hub.Config({
  tex2jax: {
    inlineMath: [['$','$'], ['\\(','\\)']],
    // displayMath: [['$$','$$'], ['\\[','\\]']], // default
    processEscapes: true
  }
});
</script>
```

`inlineMath` はインライン数式の開始・終了デリミタを指定する。
複数指定可能。
上記の設定であれば `$ ... $` または `\( ... \)` で囲まれた部分が [MathJax] の処理対象となる。
`inlineMath` の既定値は `[['\\(','\\)']]` のみである。

`displayMath` は別行立て数式の開始・終了デリミタを指定する。
同じように `$$ ... $$` または `\[ ... \]` で囲まれた部分が [MathJax] の処理対象となる。
`displayMath` の既定値は `[['$$','$$'], ['\\[','\\]']]` である。

`processEscapes` は上述の数式開始・終了記号をエスケープするかどうかを指定する。
このオプションが true なら「`\$E=mc^2\$`」と記述すれば「\$E=mc^2\$」と表示される。
既定値は false。

### [“TeX” Input Processor Options](http://docs.mathjax.org/en/latest/options/input-processors/TeX.html "The TeX input processor — MathJax 2.7 documentation")

```html
<script type="text/x-mathjax-config">
MathJax.Hub.Config({
  TeX: {
    equationNumbers: { autoNumber: "AMS" }
 }
});
</script>
```

`equationNumbers` 項目で別行立て数式の採番を制御する。
このうち自動採番については `autoNumber` で指定する。
指定可能な値は `"none"`, `"all"`, `"AMS"`。
自動採番を行わない場合（既定値）は `"none"` をセットする。
ページ内の全ての別行立て数式に採番を行う場合は `"all"` をセットする。

`"AMS"` をセットすると `\begin{equation} ... \end{equation}` で囲まれた数式のみ採番を行う。

```html
\[ \begin{equation}
  \frac{\pi}{2} =
  \left( \int_{0}^{\infty} \frac{\sin x}{\sqrt{x}} dx \right)^2 =
  \sum_{k=0}^{\infty} \frac{(2k)!}{2^{2k}(k!)^2} \frac{1}{2k+1} =
  \prod_{k=1}^{\infty} \frac{4k^2}{4k^2 - 1}
  \label{eq:1}
\end{equation} \]
```

{{< fig-quote >}}
\[ \begin{equation}
  \frac{\pi}{2} =
  \left( \int_{0}^{\infty} \frac{\sin x}{\sqrt{x}} dx \right)^2 =
  \sum_{k=0}^{\infty} \frac{(2k)!}{2^{2k}(k!)^2} \frac{1}{2k+1} =
  \prod_{k=1}^{\infty} \frac{4k^2}{4k^2 - 1}
  \label{eq:1}
\end{equation} \]
{{< /fig-quote >}}

このとき，ラベル `\label{eq:1}` に対する `\eqref{eq:1}` は “\eqref{eq:1}” と展開される（`$ ... $` で囲む必要はない）。

### [“CommonHTML” Output Processor Options](http://docs.mathjax.org/en/latest/options/output-processors/CommonHTML.html "The CommonHTML output processor — MathJax 2.7 documentation")

```html
<script type="text/x-mathjax-config">
MathJax.Hub.Config({
  CommonHTML: {
    matchFontHeight: false
  }
});
</script>
```

`matchFontHeight` が true であれば本文の文字の大きさにマッチするように数式フォントの高さを調節してくれる。
ただし本文が日本語だと却ってバランスが悪いようだ。
したがって（既定値は true だが） false にしておくのが安全である[^op1]。

[^op1]: `matchFontHeight` 項目は [“HTML-CSS” オプション]にもあって，マニュアルには [“HTML-CSS” オプション]の方に説明があるのだが，実際には [“CommonHTML” オプション]の方が有効なようである。

### [“HTML-CSS” Output Processor Options](http://docs.mathjax.org/en/latest/options/output-processors/HTML-CSS.html "The HTML-CSS output processor — MathJax 2.7 documentation")

```html
<script type="text/x-mathjax-config">
MathJax.Hub.Config({
  "HTML-CSS": {
    availableFonts: [],
    preferredFont: null,
    webFont: "STIX-Web"
  }
});
</script>
```

`availableFonts` で指定したフォントがローカル環境にある場合は，そのフォントを使用する。
`availableFonts` では複数のフォントを指定可能。
またローカルのフォントを使用しない場合は空の配列 `[]` を指定する。
既定値は `["STIX","TeX"]`。

`availableFonts` で指定したフォントのうち優先して使うフォントを `preferredFont` で指定する。
ローカルのフォントを使用しない場合は null を指定する。
既定値は `"TeX"`。

Web フォントを使用する場合は `webFont` で指定する。
指定可能なフォントは `"TeX"`, `"STIX-Web"`, `"Asana-Math"`, `"Neo-Euler"`, `"Gyre-Pagella"`, `"Gyre-Termes"`, `"Latin-Modern"`。
Web フォントを使用しない場合は null を指定する。
既定値は `"TeX"`。

指定可能な Web フォントのうち，最も完成度が高いのは `"STIX-Web"` だそうだ。
ちなみに [STIX (Scientific and Technical Information Exchange)](http://www.stixfonts.org/ "STIX Fonts Project Website") フォントは Times 系フォントと相性のいいフォントと言われている。
一方 `"Neo-Euler"` は黒板手書き風のフォントで人気が高いのだが[^fnt1] 今のところイタリック体が同梱されてない模様。

[^fnt1]: 結城浩さんの「数学ガール」シリーズでは数式に Euler フォントを使用している。

### 最終的なオプション設定

[本ブログ](/ "text.Baldanders.info")における [MathJax] オプションの設定内容を以下に示す。

```html
<script type="text/x-mathjax-config">
MathJax.Hub.Config({
  displayAlign: "left",
  displayIndent: "2em",
  tex2jax: {
    inlineMath: [['$','$'], ['\\(','\\)']],
    processEscapes: true
  },
  TeX: {
    equationNumbers: { autoNumber: "AMS" }
  },
  CommonHTML: {
    matchFontHeight: false
  },
  "HTML-CSS": {
    availableFonts: [],
    preferredFont: null,
    webFont: "STIX-Web"
  }
});
</script>
```

これでようやく準備が整った。

## ブックマーク

- [MathJaxによる数式表示](https://oku.edu.mie-u.ac.jp/~okumura/javascript/mathjax.html)

[MathJax]: https://www.mathjax.org/
[“CommonHTML” オプション]: http://docs.mathjax.org/en/latest/options/output-processors/CommonHTML.html "The CommonHTML output processor — MathJax 2.7 documentation"
[“HTML-CSS” オプション]: http://docs.mathjax.org/en/latest/options/output-processors/HTML-CSS.html "The HTML-CSS output processor — MathJax 2.7 documentation"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4774187054/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/51E5K7B53aL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4774187054/baldandersinf-22/">[改訂第7版]LaTeX2ε美文書作成入門</a></dt><dd>奥村 晴彦 黒木 裕介 </dd><dd>技術評論社 2017-01-24</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798118141/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798118141.09._SCTHUMBZZZ_.jpg"  alt="LaTeX2e辞典~用法・用例逆引きリファレンス (DESKTOP REFERENCE)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4535558752/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4535558752.09._SCTHUMBZZZ_.jpg"  alt="公共政策入門 ミクロ経済学的アプローチ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4320112415/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4320112415.09._SCTHUMBZZZ_.jpg"  alt="Rで楽しむ統計 (Wonderful R 1)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4000298550/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4000298550.09._SCTHUMBZZZ_.jpg"  alt="岩波データサイエンス Vol.5"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797391383/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797391383.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/積分を見つめて (数学ガールの秘密ノートシリーズ)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4000298569/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4000298569.09._SCTHUMBZZZ_.jpg"  alt="岩波データサイエンス Vol.6"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798115363/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798115363.09._SCTHUMBZZZ_.jpg"  alt="独習 LaTeX2ε"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4785315717/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4785315717.09._SCTHUMBZZZ_.jpg"  alt="具体例から学ぶ 多様体"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774193046/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774193046.09._SCTHUMBZZZ_.jpg"  alt="【改訂第3版】基礎からわかる情報リテラシー"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4768704700/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4768704700.09._SCTHUMBZZZ_.jpg"  alt="はじめて学ぶリー群 ―線型代数から始めよう"  /></a> </p>
<p class="description">ついに第7版が登場。紙の本で買って常に側に置いておくのが吉。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-09-27">2017-09-27</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
