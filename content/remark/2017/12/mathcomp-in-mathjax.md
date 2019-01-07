+++
title = "ちょこっと MathJax 番外編： mathcomp パッケージの代替え"
date =  "2017-12-22T19:15:52+09:00"
description = "MathJax では \\tcdegree コマンドを持っていないようだ。ググってみたら「Unicode を使え」という身も蓋もない回答が見つかった。"
image = "/images/attention/remark.jpg"
tags        = [ "math", "tex", "mathjax", "javascript" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = true
  mermaidjs = false
+++

今回は横道に外れて小ネタでいきます。

1. [ちょこっと MathJax： 初期設定]({{< ref "/remark/2017/09/getting-started-mathjax-1.md" >}})
2. [ちょこっと MathJax： 基本的な数式表現]({{< ref "/remark/2017/09/getting-started-mathjax-2.md" >}})
3. [ちょこっと MathJax： インライン数式と別行立て数式]({{< ref "/remark/2017/10/getting-started-mathjax-3.md" >}})
4. [ちょこっと MathJax 番外編： mathcomp パッケージの代替え]({{< ref "/remark/2017/12/mathcomp-in-mathjax.md" >}}) ← イマココ

たとえば

> 直角は ±90°。

をインライン数式で書くことを考える。

昔の $\mathrm{\LaTeX}$ は「°」に相当する記号がなく， `90^{\circ}` などと表記していた。
すなわち

```html
直角は $\pm90^{\circ}$。
```

と書いて

> 直角は $\pm90^{\circ}$。

と表示させていたわけだ。
しかし `\circ` ($\circ$) は角度を表す記号ではなく $+$ や $-$ と同じ2項演算子の記号である。
そこで今は [mathcomp] パッケージの `\tcdegree` コマンドを使って

```html
直角は $\pm90\tcdegree$。
```

とするのが正しいそうだ[^ang1]。

[^ang1]: 他の方法としては [siunitx] パッケージを使って `\ang{90}` のように記述する手がある。

ところが [MathJax] では `\tcdegree` コマンドを持っていないようで，上の記述ではエラーになる。
[MathJax] の Extensions でもそれらしいものは見当たらず，誰かが公開してるってわけでもないっぽい。

みんなどうやってるんだ？

ググってみたら「Unicode の「°」文字を使うか `\unicode{xb0}` で指定しろ」という身も蓋もない回答が見つかった。
どうもそれしかないらしい... sigh

そこで $\mathrm{\LaTeX}$ との互換性を保つため

```html
<script type="text/x-mathjax-config">
MathJax.Hub.Config({
  TeX: {
    Macros: {
        tcdegree: ['\\unicode{xb0}']
    }
 }
});
</script>
```

という感じでマクロを組んでみた。
これで，以下のように表示可能になる。

> 直角は $\pm90\tcdegree$。

[mathcomp] パッケージで使える記号はかなりあるのだが，私がよく使いそうなもののみ以下に挙げる。

| Macros                                | $\mathrm{\LaTeX}$ Format       | Output                |
|:------------------------------------- |:----------------------- |:--------------------- |
| `tcdegree: ['\\unicode{xb0}']`        | `$35\tcdegree$`         | $35\tcdegree$         |
| `tccelsius: ['\\unicode{x2103}']`     | `$35\,\tccelsius$`      | $35\,\tccelsius$      |
| `tcperthousand: ['\\unicode{x2030}']` | `$35\,\tcperthousand$`  | $35\,\tcperthousand$  |
| `tcmu: ['\\unicode{x3bc}']`           | `$35\,\tcmu\mathrm{s}$` | $35\,\tcmu\mathrm{s}$ |
| `tcohm: ['\\unicode{x3a9}']`          | `$35\,\tcohm$`          | $35\,\tcohm$          | 

- `\tccelsius` は `\tcdegree\mathrm{C}` でもいける（`$35\,\tcdegree\mathrm{C}$` →  $35\,\tcdegree\mathrm{C}$）。これなら華氏にも応用できる（`$35\,\tcdegree\mathrm{F}$` →  $35\,\tcdegree\mathrm{F}$）
- `\tcmu` は `\symup{\mu}` みたいにしたいのだが `\symup` コマンドに相当するものが [MathJax] にないっぽい
- `\tcohm` は `\Omega` と等価，というかフォントのバランスを考えると `\Omega` を使ったほうがいいと思う（`$35\,\Omega$` →  $35\,\Omega$）

## ブックマーク

- [siunitx - TeX Wiki](https://texwiki.texjp.org/?siunitx)

[MathJax]: https://www.mathjax.org/
[mathcomp]: https://ctan.org/pkg/mathcomp "https://ctan.org/pkg/mathcomp"
[siunitx]: https://ctan.org/pkg/siunitx "CTAN: Package siunitx"

## 参考図書 {#books}

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4774187054/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/51E5K7B53aL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4774187054/baldandersinf-22/">[改訂第7版]LaTeX2ε美文書作成入門</a></dt><dd>奥村 晴彦 黒木 裕介 </dd><dd>技術評論社 2017-01-24</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798118141/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798118141.09._SCTHUMBZZZ_.jpg"  alt="LaTeX2e辞典~用法・用例逆引きリファレンス (DESKTOP REFERENCE)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4535558752/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4535558752.09._SCTHUMBZZZ_.jpg"  alt="公共政策入門 ミクロ経済学的アプローチ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4320112415/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4320112415.09._SCTHUMBZZZ_.jpg"  alt="Rで楽しむ統計 (Wonderful R 1)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4000298550/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4000298550.09._SCTHUMBZZZ_.jpg"  alt="岩波データサイエンス Vol.5"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797391383/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797391383.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/積分を見つめて (数学ガールの秘密ノートシリーズ)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4000298569/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4000298569.09._SCTHUMBZZZ_.jpg"  alt="岩波データサイエンス Vol.6"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798115363/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798115363.09._SCTHUMBZZZ_.jpg"  alt="独習 LaTeX2ε"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4785315717/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4785315717.09._SCTHUMBZZZ_.jpg"  alt="具体例から学ぶ 多様体"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774193046/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774193046.09._SCTHUMBZZZ_.jpg"  alt="【改訂第3版】基礎からわかる情報リテラシー"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4768704700/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4768704700.09._SCTHUMBZZZ_.jpg"  alt="はじめて学ぶリー群 ―線型代数から始めよう"  /></a> </p>
<p class="description">ついに第7版が登場。紙の本で買って常に側に置いておくのが吉。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-09-27">2017-09-27</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
