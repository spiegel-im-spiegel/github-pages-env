+++
title = "改訂第8版『LaTeX2ε美文書作成入門』を眺める"
date =  "2021-09-05T17:42:03+09:00"
description = "2020年11月に出てたの？ 完全に見落としてたよ orz"
image = "/images/attention/kitten.jpg"
tags = [ "tex", "book" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

Twitter の TL を見て気が付いた。
あれ？ 『[LaTeX2ε美文書作成入門]』の第8版って出てるの？ 完全に見落としてたよ `orz`

『[LaTeX2ε美文書作成入門]』は Kindle 版もあるし[版元]で買えば PDF 版もあるんだけど，今回は紙の本を買った。
この手のリファレンス本は，一覧性という観点では，まだ紙のほうが有利。
デジタル版が好みなら PDF 版を買うのがいいだろう。
悪いが，この手の本で Kindle 版はない。

第7版からの3年間で大きく変わったことといえば $\mathrm{Lua\LaTeX}$ 日本語環境の台頭だろう。
序文でも

{{< fig-quote type="markdown" title="［改訂第8版］LaTeX2ε美文書作成入門" link="https://www.amazon.co.jp/dp/B08MZ98Z1Q?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
{{% quote %}}本書も $\mathrm{Lua\LaTeX}$ と新しい jleq ドキュメントクラスと原の味で組版しました{{% /quote %}}
{{< /fig-quote >}}

と書かれている。
えっ，これ原の味フォントなの？ そういや巻末に原の味フォントの全グリフが載ってるな（笑）

入力環境（第2章）についても真っ先に [Cloud LaTeX] が紹介されているし[^clatex1]，コラム内とはいえ VS Code 用の [LaTeX Workshop](https://marketplace.visualstudio.com/items?itemName=James-Yu.latex-workshop "LaTeX Workshop - Visual Studio Marketplace") が紹介されていたりする。
また jleq ドキュメントクラスについても14章で詳しく解説している。
時代が変わったんだねぇ。

[^clatex1]: [Cloud LaTeX] 連携用の [VS Code 拡張機能](https://marketplace.visualstudio.com/items?itemName=cloudlatex.cloudlatex "Cloud LaTeX - Visual Studio Marketplace")があるんだな。

グラフィック周りについては第7版に引き続き付録Dで [$\mathrm{Ti}k\mathrm{Z}$/PGF][TikZ] を詳しく解説している。

{{< fig-quote type="markdown" title="［改訂第8版］LaTeX2ε美文書作成入門" link="https://www.amazon.co.jp/dp/B08MZ98Z1Q?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
{{% quote %}}$\mathrm{Ti}k\mathrm{Z}$ に慣れると `picture` 環境は使いたくなくなります{{% /quote %}}
{{< /fig-quote >}}

とか書かれてあって，ちょっと笑ってしまう。

第8版では R, gnuplot 以外に Python を使った例を挙げている。
ちなみに Go 言語でも [gonum.org/v1/plot](https://github.com/gonum/plot "gonum/plot: A repository for plotting and visualizing data") パッケージで $\mathrm{Ti}k\mathrm{Z}$ 出力をサポートしている。
また，みんな大好き [DOT 言語](https://graphviz.org/doc/info/lang.html "DOT Language | Graphviz")でも [dot2tex] というツール（Python 製）を使うと $\mathrm{Ti}k\mathrm{Z}$ 形式に変換してくれるらしい。
もっとも [dot2tex] には [`dot2texi` $\mathrm{\LaTeX}$ パッケージ](https://dot2tex.readthedocs.io/en/latest/tipsandtricks.html#the-dot2texi-latex-package)もあるようなので，こっちのほうが便利かもしれないが。
あと [PlantUML](https://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw beautiful UML diagrams.") でもベータ版ながら [$\mathrm{Ti}k\mathrm{Z}$ 形式に対応](https://plantuml.com/ja/latex)しているようだ。

ぶっちゃけ PDF を最終出力とするのなら $\mathrm{Lua\LaTeX}$ で必要十分だよな。
まぁ，論文を書くとなると各学会用のドキュメントクラスを使わないといけないので簡単ではないだろうが。

最近は $\mathrm{Lua\LaTeX}$ を直にゴリゴリ書く機会は無くなったが『[LaTeX2ε美文書作成入門]』を眺めていると色々遊びたくなってくる。
今の仕事が落ち着いたら試してみたいところである。

[LaTeX2ε美文書作成入門]: https://www.amazon.co.jp/dp/4297117126?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "[改訂第8版]LaTeX2ε美文書作成入門 | 奥村晴彦, 黒木裕介 |本 | 通販 | Amazon"
[版元]: https://gihyo.jp/book/2020/978-4-297-11712-2 "［改訂第8版］LaTeX2ε美文書作成入門：書籍案内｜技術評論社"
[Cloud LaTeX]: https://cloudlatex.io/ "Cloud LaTeX | Build your own LaTeX environment, in seconds"
[TikZ]: https://github.com/pgf-tikz "pgf-tikz"
[dot2tex]: https://dot2tex.readthedocs.io/ "dot2tex - A Graphviz to LaTeX converter — dot2tex 2.11.3 documentation"

## ブックマーク

- [TeX Wiki](https://texwiki.texjp.org/)
- [LaTeX Workshopってなに？どうやって使うの？調べてみた！ - Qiita](https://qiita.com/moinslut/items/bc1d1b1e13cb38377406)
- [TikZ実例集〜2Dグラフ編 - Notes_JP](https://www.mynote-jp.com/entry/TikZ-examples-2D-graph)
- [TikZ実例集〜3D編 - Notes_JP](https://www.mynote-jp.com/entry/TikZ-examples)

## 参考図書

{{% review-paapi "4297117126" %}} <!-- LaTeX2ε美文書作成入門 -->
