+++
title = "モンテカルロ法による円周率の推定（その2 Gaussian）"
description = "description"
tags = [
  "remark",
]
draft = true
date = "2016-11-09T20:41:57+09:00"

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
3. [モンテカルロ法による円周率の推定（その2 Gaussian）]({{< relref "golang/estimate-of-pi-3-gaussian.md" >}}) ← イマココ

さて，いよいよ円周率の推定結果を評価してみよう。

まずは「[その1]」で求めた推定方法をもう一度おさらいしてみる。

$0 \le y \le 1$ および $0 \le y \le 1$ の領域内にランダムにプロットされた点が原点を中心とする単位円（半径 1 の円）の内側に入る確率 $p$ は面積比から

{{< fig-quote >}}
\[
p = \frac{1}{4}{\pi} = 0.785398 \ldots
\]
{{< /fig-quote >}}

と考えられる。
したがって，ランダムにプロットされた点の総数を $n$ とすると単位円に入る点の数 $m$ は

{{< fig-quote >}}
\[
m \simeq np
\]
{{< /fig-quote >}}

となることが期待できる。
これを期待値（expected value） $\mu$ とする。

{{< fig-quote >}}
\[
\mu = np = \frac{1}{4}{\pi}n
\]
{{< /fig-quote >}}

$m$ を求める実験を何度も繰り返したときの $m$ の分布は二項分布となり，その標準偏差（standard deviation） $\sigma$ は

{{< fig-quote >}}
\[
\sigma = \sqrt{np(1-p)}
\]
{{< /fig-quote >}}

となる。
さらに $n$ が大きければ $m$ の分布は正規分布（またはガウス分布）に近似できる。
円周率の推定値は ${4m}/{n}$ なので，円周率の推定値の分布もまた正規分布に近似できると考えられる。

というわけでヒストグラムを作ってみることにする。
CLI はこんな感じで

```text
$ go run main.go estmt --help
Estimate of Pi with Monte Carlo method.

Usage:
  pi estmt [flags]

Flags:
  -e, --ecount int     Count of estimate (default 100)
  -c, --hclass float   Class interval of histogram
  -p, --pcount int     Count of points (default 10000)

Global Flags:
      --config string   config file (default is $HOME/.pi.yaml)
```

`-c` オプションで階級幅を指定するとヒストグラム用のデータを出力するようにした。
実際の処理はこんな感じ

```go
//Execute returns estimate of Pi.
func Execute(cxt *Context) error {
	if cxt.pointCount <= 0 {
		return fmt.Errorf("invalid argument \"%v\" for pcount option", cxt.pointCount)
	}
	if cxt.estimateCount <= 0 {
		return fmt.Errorf("invalid argument \"%v\" for ecount option", cxt.estimateCount)
	}

	ch := genpi.New(cxt.pointCount, cxt.estimateCount)
	min := float64(10)
	max := float64(0)
	sum := float64(0)
	sum2 := float64(0)
	pis := make([]float64, 0, cxt.estimateCount)
	for pi := range ch {
		pis = append(pis, pi)
		if pi < min {
			min = pi
		}
		if pi > max {
			max = pi
		}
		sum += pi
		sum2 += pi * pi
	}
	cxt.ui.OutputErrln(fmt.Sprintf("minimum value: %7.5f", min))
	cxt.ui.OutputErrln(fmt.Sprintf("maximum value: %7.5f", max))
	ave := sum / float64(cxt.estimateCount)
	cxt.ui.OutputErrln(fmt.Sprintf("average: %7.5f", ave))
	vari := sum2/float64(cxt.estimateCount) - ave*ave
	cxt.ui.OutputErrln(fmt.Sprintf("standard deviation: %7.5f", math.Sqrt(vari)))

	if cxt.histClass <= 0.0 {
		for _, pi := range pis {
			cxt.ui.Outputln(fmt.Sprintf("%7.5f", pi))
		}
		return nil
	}

	classCount := int((max-min)/cxt.histClass) + 1
	freqs := make([]int, classCount)
	for _, pi := range pis {
		class := 0
		for i := 0; i < classCount; i++ {
			if min+cxt.histClass*float64(i) <= pi {
				class = i
			} else {
				break
			}
		}
		freqs[class]++
	}
	for i, freq := range freqs {
		mid := min + cxt.histClass*float64(i) + cxt.histClass/2.0
		cxt.ui.Outputln(fmt.Sprintf("%7.5f\t%v", mid, freq))
	}

	return nil
}
```

では推定処理を10,000回，階級幅を0.001で試してみる。

```text
$ go run main.go estmt -e 10000 -p 100000 -c 0.001 > hist.dat
minimum value: 3.12176
maximum value: 3.15972
average value: 3.14163
standard deviation: 0.00520
```

で `hist.dat` を [gnuplot] に食わせてみるとこんな感じ。

{{< fig-img src="/images/histogram.png" link="/images/histogram.png" width="753" >}}

んー，どうかなぁ。
平均値や標準偏差はそれっぽいけど正規分布というには...

## ブックマーク

- [モンテカルロ法の誤差を考える](http://ruby.kyoto-wu.ac.jp/info-com/NumericalModels/RandomProcess/estimateMCmodel.html)
- [gnuplot でヒストグラム・度数分布多角形づくり: 個人的健忘録 from 2013](http://bluearth.cocolog-nifty.com/blog/2014/03/gnuplot-d8cc-1.html)

[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[その1]: {{< relref "golang/estimate-of-pi.md" >}} "モンテカルロ法による円周率の推定（その1）"
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
