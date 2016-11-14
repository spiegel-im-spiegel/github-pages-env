+++
title = "モンテカルロ法による円周率の推定（その3 Gaussian）"
description = "さて，前回書いたコードを利用していよいよ円周率の推定結果を評価してみる。"
tags = [
  "golang",
  "math",
  "random",
  "circle-ratio",
  "gaussian",
]
draft = true
date = "2016-11-14T12:32:18+09:00"

[author]
  instagram = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  avatar = "/images/avatar.jpg"
  license = "by-sa"
  github = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  tumblr = "spiegel-im-spiegel"
+++

1. [モンテカルロ法による円周率の推定（その1）]({{< relref "golang/estimate-of-pi.md" >}})
2. [モンテカルロ法による円周率の推定（その2 CLI）]({{< relref "golang/estimate-of-pi-2-cli.md" >}})
3. [モンテカルロ法による円周率の推定（その3 Gaussian）]({{< relref "golang/estimate-of-pi-3-gaussian.md" >}}) ← イマココ

## 推定結果の分布

さて，[前回]書いたコードを利用していよいよ円周率の推定結果を評価してみる。

- [spiegel-im-spiegel/pi: Estimate of Pi with Monte Carlo method.](https://github.com/spiegel-im-spiegel/pi)

CLI はこんな感じになった。

```text
$ go run main.go estmt --help
Estimate of Pi with Monte Carlo method.

Usage:
  pi estmt [flags]

Flags:
  -e, --ecount int   Count of estimate (default 100)
  -p, --pcount int   Count of points (default 10000)

Global Flags:
      --config string   config file (default is $HOME/.pi.yaml)
```

まずは円周率の推定処理を10,000回繰り返してみる。
また推定処理のためのランダムな点の数 $n$ を1,000個，10,000個，100,000個と変えて実行してみようか。

```text
$ go run main.go estmt -e 10000 -p 1000 > estmt1k.dat
minimum value: 2.94800
maximum value: 3.31600
average value: 3.14199
standard deviation: 0.05104 (69.9%)

$ go run main.go estmt -e 10000 -p 10000 > estmt10k.dat
minimum value: 3.07240
maximum value: 3.20360
average value: 3.14178
standard deviation: 0.01654 (68.0%)

$ go run main.go estmt -e 10000 -p 100000 > estmt100k.dat
minimum value: 3.12360
maximum value: 3.16184
average value: 3.14163
standard deviation: 0.00518 (68.3%)
```

最後のはさすがに時間がかかるので，お茶でも飲みながら優雅に待ちましょう（笑）

標準エラー出力に最小値，最大値，平均値（$E$），標準偏差（$\sigma$）を出力してみた。
標準偏差の後ろの括弧は $\left[ E-\sigma, E+\sigma \right]$ の範囲にある推定値の割合を示したものだ。

円周率の推定処理の試行回数が十分大きいなら推定値の分布は正規分布（またはガウス分布）に近似できる筈である。
（以下の[図は Wikimedia Commons のもの](https://commons.wikimedia.org/wiki/File:Standard_deviation_diagram.svg "File:Standard deviation diagram.svg - Wikimedia Commons")を拝借した。 [CC-BY-2.5 Generic](https://creativecommons.org/licenses/by/2.5/ "Creative Commons — Attribution 2.5 Generic — CC BY 2.5") で公開されている）
{{< fig-img src="https://upload.wikimedia.org/wikipedia/commons/thumb/8/8c/Standard_deviation_diagram.svg/640px-Standard_deviation_diagram.svg.png" width="640" link="https://commons.wikimedia.org/wiki/File:Standard_deviation_diagram.svg" title="normal distribution from Wikimedia" >}}

そこで $n=100,000$ のときの推定結果についてヒストグラムを描いてみる。
幸いなことに [gnuplot] では簡単にヒストグラムを作図できる。
こんな感じ（階級幅を0.001としている）。

```text
gnuplot> filter(x,y)=int(x/y)*y
gnuplot> plot "estmt100k.dat" u (filter($1,0.001)):(1) smooth frequency with boxes lc rgb "black"
```

[gnuplot] の出力結果はこんな感じ。

{{< fig-img src="/images/histogram.png" link="/images/histogram.png" width="640" >}}

んー。
まぁ正規分布っぽい？

もうひとつ，正規確率の分布を調べてみよう。
これも [gnuplot] で描こうと思ったけど，少し面倒そうなので，ズルして以下を参考に Excel で描くことにした。

- [正規確率プロット（Normal Q-Q Plot）の作成 with Excel](http://hitorimarketing.net/tools/normal-quantile-quantile-plot.html)

とりあえず結果だけ。

{{< fig-img src="/images/qq-plot.png" link="/images/qq-plot.png" width="614" >}}

プロットが直線状に並んでいれば正規分布であると言える。
図から見る限り，概ね正規分布になっているようである。

## モンテカルロ法の誤差評価

モンテカルロ法を使ってどの程度の精度で円周率が求まるかの考察については以下が参考になる。

- [モンテカルロ法の誤差を考える](http://ruby.kyoto-wu.ac.jp/info-com/NumericalModels/RandomProcess/estimateMCmodel.html)

これも横着して結果だけを拝借すると， $n=100,000$ で推定を行った場合の値の分布は，中央値を $\pi$，99%信頼区間を $\frac{4.230}{\sqrt{100,000}} = 0.013$ として， $\left[ \pi - 0.013, \pi + 0.013 \right]$ の範囲になるようだ。

しまった。
今回は [Go 言語]が全然出てこなかった。
まぁ，いいや。
多分あと1回続きます。

## ブックマーク

- [gnuplot でヒストグラム（頻度分布図）を描画する - Qiita](http://qiita.com/iwiwi/items/4c7635d4c84bc785e47a)
- [【統計学】Q-Qプロットの仕組みをアニメーションで理解する。 - Qiita](http://qiita.com/kenmatsu4/items/59605dc745707e8701e0)

[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。

[前回]: {{< relref "golang/estimate-of-pi-2-cli.md" >}} "モンテカルロ法による円周率の推定（その2 CLI）"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[gnuplot]: http://www.gnuplot.info/ "gnuplot homepage"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-20</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4865940391/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4865940391.09._SCTHUMBZZZ_.jpg"  alt="Kotlinスタートブック -新しいAndroidプログラミング"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4839959234/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4839959234.09._SCTHUMBZZZ_.jpg"  alt="Docker実戦活用ガイド"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274218961/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274218961.09._SCTHUMBZZZ_.jpg"  alt="グッド・マス ギークのための数・論理・計算機科学"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-13">2016-07-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1FO/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/416jAxVU4NL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1FO/baldandersinf-22/">数学ガール／乱択アルゴリズム</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ株式会社 2014-02-14</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1D6/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00I8AT1D6.09._SCTHUMBZZZ_.jpg"  alt="数学ガール／ゲーデルの不完全性定理"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1CM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00I8AT1CM.09._SCTHUMBZZZ_.jpg"  alt="数学ガール／フェルマーの最終定理"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00L0PDMK4/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00L0PDMK4.09._SCTHUMBZZZ_.jpg"  alt="数学ガール／ガロア理論"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00EYXMA9I/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00EYXMA9I.09._SCTHUMBZZZ_.jpg"  alt="数学ガール"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00L0PDMIQ/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00L0PDMIQ.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート／式とグラフ"  /></a> </p>
<p class="description" >工学ガール，リサちゃん登場！</p>
<p class="gtools" >reviewed by <a href="#maker" class="reviewer">Spiegel</a> on <abbr class="dtreviewed" title="2015-04-19">2015/04/19</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html">G-Tools</a>)</p>
</div>
