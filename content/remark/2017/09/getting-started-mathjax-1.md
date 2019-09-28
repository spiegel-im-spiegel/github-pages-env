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
エネルギーと質量には $E=mc^2$ の関係がある。
```

と書くとブラウザ側では

> エネルギーと質量には $E=mc^2$ の関係がある。

と適切に表示してくれる[^mj1]。

この記事ではまず Web ページ上で [MathJax] が動くところまでを説明していこう。
数式の書き方については[次回]以降に解説していく予定である。

[^mj1]: `$E=mc^2$` という入力に対して $E=mc^2$ と，各文字間を適切に空けたり詰めたりしてくれるのがお分かりだろうか。このように $\mathrm{\TeX}$ では数式を半自動的かつ適切に「組版」してくれるのが特徴である。ただし万能ではない。

{{< div-box type="md" >}}
## 2019-09-28 追記

[MathJax v3 がリリース]({{< ref "/release/2019/09/mathjax-v3-is-released.md" >}} "MathJax v3 がリリースされていた")された。
v2.7.x までとは大きく変わったためこの記事も改訂する予定である。
{{< /div-box >}}

## [MathJax] のインストール {#install}

[MathJax] は JavaScript パッケージなので `<script>` 要素で指定する。
CDN (Content Delivery Network) で配布されているので HTML の `<head>` 要素内に

```html
<script async src="https://cdnjs.cloudflare.com/ajax/libs/mathjax/2.7.6/latest.js?config=TeX-AMS_HTML&amp;locale=ja"></script>
```

と記述すればよい（`async` を忘れずに）。
なお URL の “`2.7.6`” 部分はバージョン番号で，2019年8月21日時点の最新バージョンは 2.7.6 である[^cdn1]。

[^cdn1]: ちなみに2017年5月から CDN の[配布 URL が変わった](https://www.mathjax.org/cdn-shutting-down/ "MathJax CDN shutting down on April 30, 2017.  Alternatives available.")ので，設定が古いままの方は注意が必要である。

- [MathJax v2.7.6 now available | MathJax](https://www.mathjax.org/mathjax-v2-7-6-now-available/)

URL パラメータ部の `config=TeX-AMS_HTML` については[指定可能なコンフィギュレーション](http://docs.mathjax.org/en/latest/config-files.html "Combined Configurations — MathJax 2.7 documentation")がいくつかあるが，  $\mathrm{\TeX}$ 記法を使うのであれば `TeX-AMS_CHTML` または `TeX-AMS_HTML` を指定するのがいいだろう[^html1]。
さらにパラメータ部に `locale=ja` を追加すると，数式部分で表示されるコンテキスト・メニューが日本語になる。

[^html1]: `TeX-AMS_CHTML` を指定すると `HTML-CSS` オプションをまるっと無視してしまい Web フォントの指定ができない。逆に `TeX-AMS_HTML` では `CommonHTML` オプションを無視してしまうようだ。 [MathJax] 側は `TeX-AMS_HTML` を古いブラウザ向けのレガシーなコンフィギュレーションと位置付けているようだが，やはり `TeX-AMS_CHTML` では Web フォントのカスタマイズができない（現在は `"TeX"` のみサポート）のが致命的だろう。この辺は今後のバージョンアップに期待したいところ。

{{< fig-img src="https://photo.baldanders.info/flickr/image/37316621442_m.png" title="MathJax: context menu" link="https://photo.baldanders.info/flickr/37316621442/" >}}

ブラボー！

## [MathJax] のオプション {#options}

[MathJax] にはいくつかのオプションを設定できる。
オプションの設定には JavaScript で `MathJax.Hub.Config()` 関数を使う。
先程の `<script>` 指定と併せて `<head>` 要素内に以下のように記述する。

```html
<script type="text/x-mathjax-config">
MathJax.Hub.Config({ ... });
</script>
<script async src="https://cdnjs.cloudflare.com/ajax/libs/mathjax/2.7.2/MathJax.js?config=TeX-AMS_HTML&locale=ja"></script>
```

オプション指定部のメディア・タイプに `text/x-mathjax-config` を指定すること（`<script>` の順番にも注意）。

`{ ... }` のオブジェクトに具体的なオプションを記述していく。
全部を説明するのは大変なので，よく使いそうなものを幾つか紹介しよう。
なお，[最後の節]({{< relref "#mysetting" >}})に[本ブログ]におけるオプションの設定例を挙げているので，以降の解説がウザい方は丸写しでも OK です（笑）

（次節以降に出てくる「インライン数式」および「別行立て数式」については[第3回]で説明する）

### [Core Configuration Options](http://docs.mathjax.org/en/latest/options/hub.html "The Core Configuration Options — MathJax 2.7 documentation") {#core}

```html
<script type="text/x-mathjax-config">
MathJax.Hub.Config({
  displayAlign: "left",
  displayIndent: "2em"
});
</script>
```

`displayAlign` は別行立て数式の位置を何処に寄せるか指定する。
左寄せ（`"left"`），右寄せ（`"right"`），中央寄せ（`"center"`）を指定できる。
既定値は `"center"`。
また中央寄せ以外のときは `displayIndent` でインデント幅を指定できる。

左寄せで2文字分インデントさせた場合は以下のように表示される。

> エネルギーと質量には $$E=mc^2$$ の関係がある。

### [“tex2jax” Preprocessor Options](http://docs.mathjax.org/en/latest/options/preprocessors/tex2jax.html "The tex2jax Preprocessor — MathJax 2.7 documentation") {#tex2jax}

```html
<script type="text/x-mathjax-config">
MathJax.Hub.Config({
  tex2jax: {
    inlineMath: [['$','$'], ['\\(','\\)']],
    // displayMath: [['$$','$$'], ['\[','\]']], // default
    processEscapes: true
  }
});
</script>
```

`inlineMath` はインライン数式の開始・終了デリミタを指定する。
複数列挙できる。
上記の設定であれば `$...$` または `\(...\)` で囲まれた部分が [MathJax] の処理対象となる。
`inlineMath` の既定値は `[['\(','\)']]` のみである。

```html
エネルギーと質量には $E=mc^2$ の関係がある。
```

> エネルギーと質量には $E=mc^2$ の関係がある。

`displayMath` は別行立て数式の開始・終了デリミタを指定する。
こちらも複数列挙できる。
インライン数式と同じように `$$...$$` または `\[...\]` で囲まれた部分が [MathJax] の処理対象となる。
`displayMath` の既定値は `[['$$','$$'], ['\[','\]']]` である。

```html
エネルギーと質量には $$E=mc^2$$ の関係がある。
```

> エネルギーと質量には $$E=mc^2$$ の関係がある。

`processEscapes` を true にすると，上述の数式開始・終了デリミタを `\` 記号でエスケープする[^esc1]。

[^esc1]: `processEscapes` オプションを有効にすると `\(...\)` までエスケープされてただの `(...)` になってしまうので注意すること。というか `processEscapes` オプションを有効にするなら `\(...\)` は使わないほうがいいかも。また `processEscapes` オプションはパラグラフ `<p>...</p>` の中でのみ効いているようだ。

```html
エネルギーと質量には \$E=mc^2\$ の関係がある。
```

> エネルギーと質量には \$E=mc^2\$ の関係がある。

既定値は false。

### [“TeX” Input Processor Options](http://docs.mathjax.org/en/latest/options/input-processors/TeX.html "The TeX input processor — MathJax 2.7 documentation") {#tex}

```html
<script type="text/x-mathjax-config">
MathJax.Hub.Config({
  TeX: {
    equationNumbers: { autoNumber: "AMS" },
    extensions: ["mhchem.js"]
 }
});
</script>
```

`equationNumbers` で別行立て数式の採番を制御する。
このうち自動採番については `autoNumber` で有効・無効を指定する。
指定可能な値は `"none"`, `"all"`, `"AMS"` で，既定値は `"none"`。

自動採番を無効にする場合は `"none"` をセットする。
ページ内の全ての別行立て数式に対して自動採番を有効にする場合は `"all"` をセットする。
`"AMS"` をセットした場合の動作については[第3回]を参照のこと。

`extensions` は文字通り[拡張機能の指定](http://docs.mathjax.org/en/latest/tex.html#tex-and-latex-extensions)で，複数の拡張機能を列挙することができる。
このうち [`mhchem.js`](http://docs.mathjax.org/en/latest/tex.html#mhchem) は [MathJax] で化学式や化学反応式を記述するための拡張である。

たとえば

```html
経済成長と $\ce{CO2}$ 排出量は比例しなくなっている。
```

> 経済成長と $\ce{CO2}$ 排出量は比例しなくなっている[^co2]。

[^co2]: 「[経済成長とCO2排出量は「比例しなくなっている」：IEA報告書](http://wired.jp/2017/03/29/global-carbon-emissions/ "経済成長とCO2排出量は「比例しなくなっている」：IEA報告書｜WIRED.jp")」より。

なんてな感じに書くことができる。

### [“CommonHTML” Output Processor Options](http://docs.mathjax.org/en/latest/options/output-processors/CommonHTML.html "The CommonHTML output processor — MathJax 2.7 documentation") {#chtml}

コンフィギュレーションに `TeX-AMS_CHTML` を指定した場合に有効になるオプション。

```html
<script type="text/x-mathjax-config">
MathJax.Hub.Config({
  CommonHTML: {
    matchFontHeight: false,
    mtextFontInherit: true
  }
});
</script>
```

`matchFontHeight` が true であれば本文の文字の大きさにマッチするよう数式の文字の高さを調節してくれる。
ただし本文が日本語だと却ってバランスが悪いようだ。
したがって（既定値は true だが） false にしておくのがお薦めである。

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

{{< fig-quote >}}
\begin{alignat*}{2}
    (a+b)^2 &= a^2 + 2ab + b^2 & \qquad & \text{展開する} \\
            &= a(a + 2b) + b^2 &        & \text{a でくくる}
\end{alignat*}
{{< /fig-quote >}}

と表示される。
「`a でくくる`」の a が数式用の書体でないことに注意。
さらに「`a でくくる`」を「`$a$ でくくる`」とすると

```html
\begin{alignat*}{2}
    (a+b)^2 &= a^2 + 2ab + b^2 & \qquad & \text{展開する} \\
            &= a(a + 2b) + b^2 &        & \text{$a$ でくくる}
\end{alignat*}
```

{{< fig-quote >}}
\begin{alignat*}{2}
    (a+b)^2 &= a^2 + 2ab + b^2 & \qquad & \text{展開する} \\
            &= a(a + 2b) + b^2 &        & \text{$a$ でくくる}
\end{alignat*}
{{< /fig-quote >}}

と $a$ が数式用の書体になる。
なお，日本語（和文）部分は数式内でも `\text` コマンドの有無に関係なく影響を受けない。

### [“HTML-CSS” Output Processor Options](http://docs.mathjax.org/en/latest/options/output-processors/HTML-CSS.html "The HTML-CSS output processor — MathJax 2.7 documentation") {#html}

コンフィギュレーションに `TeX-AMS_HTML` を指定した場合に有効になるオプション。

```html
<script type="text/x-mathjax-config">
MathJax.Hub.Config({
  "HTML-CSS": {
    availableFonts: ["STIX"],
    preferredFont: "STIX",
    webFont: "STIX-Web",
    matchFontHeight: false,
    mtextFontInherit: true
  }
});
</script>
```

`availableFonts` で指定した書体がローカル環境にある[^fnt0] 場合は，その書体を使用する。
複数の書体を列挙できるが [MathJax/jax/output/HTML-CSS/fonts](https://github.com/mathjax/MathJax/tree/master/jax/output/HTML-CSS/fonts "MathJax/jax/output/HTML-CSS/fonts at master · mathjax/MathJax") のいずれかである必要がある。
ローカルの書体を使用しない場合は空の配列 `[]` を指定する。
既定値は `["STIX","TeX"]`。

[^fnt0]: $\mathrm{\TeX}$ 用フォントは $\mathrm{\TeX}$ 作業環境を整えた PC 以外にはインストールされていないのが普通である。 $\mathrm{\TeX}$ 用フォントは複数の OpenType フォントファイルで提供されるが，構成が特殊なため $\mathrm{\TeX}$ 以外での使用はおすすめできない。

`availableFonts` で指定した書体のうち優先して使う書体を `preferredFont` で指定する。
 [MathJax/jax/output/HTML-CSS/fonts](https://github.com/mathjax/MathJax/tree/master/jax/output/HTML-CSS/fonts "MathJax/jax/output/HTML-CSS/fonts at master · mathjax/MathJax") のいずれかである必要がある。
ローカルの書体を使用しない場合は null を指定する。
既定値は `"TeX"`。

Web フォントを使用する場合は `webFont` で指定する。
指定可能な書体は `"TeX"`, `"STIX-Web"`, `"Asana-Math"`, `"Neo-Euler"`, `"Gyre-Pagella"`, `"Gyre-Termes"`, `"Latin-Modern"`。
Web フォントを使用しない場合は null を指定する。
既定値は `"TeX"`。

ちなみに [STIX (Scientific and Technical Information Exchange)](http://www.stixfonts.org/ "STIX Fonts Project Website") は Times 系の書体のひとつで，長い開発期間を経て2010年に正式リリースされた。
Microsoft Office や macOS などには既に同梱されているらしい。

`"Neo-Euler"` は黒板手書き風の Euler フォントで数式好きの方には人気が高い[^fnt1] が， [MathJax] で利用する際には制限があるため取り扱いには若干の注意が必要である（[次回]で解説）。

[^fnt1]: 結城浩さんの「数学ガール」シリーズでは数式表現に Euler フォントを使用している。

{{< fig-quote title="The HTML-CSS output processor" link="http://docs.mathjax.org/en/latest/options/output-processors/HTML-CSS.html" lang="en" >}}
<q>Note that not all mathematical characters are available in all fonts (e.g., Neo-Euler does not include italic characters), so some mathematics may work better in some fonts than in others. The <code>STIX-Web</code> font is the most complete.</q>
{{< /fig-quote >}}

特にこだわりがなければ既定値どおり `"TeX"` にするか `"STIX"` (`"STIX-Web"`) を選択するのが無難だと思う（`"TeX"` にするならコンフィギュレーションを `TeX-AMS_CHTML` にすることをお薦めする）。

`matchFontHeight` および `mtextFontInherit` については [`CommonHTML` オプションの節]({{< relref "#chtml" >}})を参照のこと。

### 最終的なオプション設定 {#mysetting}

以上を踏まえて，[本ブログ]における [MathJax] オプションの設定内容を以下に示す。

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
      equationNumbers: { autoNumber: "AMS" },
      extensions: ["mhchem.js"]
  },
  "HTML-CSS": {
    availableFonts: ["STIX"],
    preferredFont: "STIX",
    webFont: "STIX-Web",
    matchFontHeight: false,
    mtextFontInherit: true
  }
});
</script>
<script async src="https://cdnjs.cloudflare.com/ajax/libs/mathjax/2.7.2/MathJax.js?config=TeX-AMS_HTML&locale=ja"></script>
```

これでようやく準備が整った。

## ブックマーク {#bookmark}

- [MathJaxによる数式表示](https://oku.edu.mie-u.ac.jp/~okumura/javascript/mathjax.html)
- [MathJaxの使い方](http://gilbert.ninja-web.net/math/mathjax1.html)
    - [MathJaxの使い方〈化学編〉](http://gilbert.ninja-web.net/math/mathjax3.html)
- [世界標準が期待される数式用フォント「STIX Fonts」 - 窓の杜](http://forest.watch.impress.co.jp/docs/news/373370.html)

[本ブログ]: / "text.Baldanders.info"
[MathJax]: https://www.mathjax.org/ "MathJax | Beautiful math in all browsers."
[“CommonHTML” オプション]: http://docs.mathjax.org/en/latest/options/output-processors/CommonHTML.html "The CommonHTML output processor — MathJax 2.7 documentation"
[“HTML-CSS” オプション]: http://docs.mathjax.org/en/latest/options/output-processors/HTML-CSS.html "The HTML-CSS output processor — MathJax 2.7 documentation"
[次回]: {{< ref "/remark/2017/09/getting-started-mathjax-2.md" >}} "ちょこっと MathJax： 基本的な数式表現"
[第3回]: {{< ref "/remark/2017/10/getting-started-mathjax-3.md" >}} "ちょこっと MathJax： インライン数式と別行立て数式"
[LaTeX2ε美文書作成入門]: https://www.amazon.co.jp/exec/obidos/ASIN/4774187054/baldandersinf-22/ "Amazon | [改訂第7版]LaTeX2ε美文書作成入門 | 奥村 晴彦, 黒木 裕介 通販"

## 参考図書 {#books}

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
