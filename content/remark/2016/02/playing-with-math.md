+++
date = "2016-02-21T22:33:47+09:00"
description = "最近の子どもは知らないが，私はこれ（誤差論）を大学に入ってから正式に習った。「小学生には無理」とは言わないが，有効桁数を含め誤差について論じるのならきちんと手順を踏んで教える必要がある。"
draft = false
tags = ["math"]
title = "算数で遊ぼう または 11 × 11 × 3.14 = 379.94 は誤りか"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

- [算数の問題「円周率を3.14とするとき、半径11の円の面積を求めよ」の解を379.94とするのは誤り？ - Togetterまとめ](http://togetter.com/li/940931)

問題はこう。

> 円周率を 3.14 とするとき、半径 11 の円の面積を求めよ

この答えとして $11 \times 11 \times 3.14 = 379.94$ とするのは誤りか？ という話。

結論から言おう。

**馬鹿じゃねーの？**

そもそもこれを正解か不正解かという観点で論じる時点で阿呆である。
日本の学校教育の悪影響がいかんなく発揮されている例と言えよう。
だから「[掛け算は順序が大事](http://www.baldanders.info/spiegel/log2/000744.shtml)」とか言う馬鹿がはびこるんだよ。

これは計算問題である。
上の問題を言い換えるのなら

> $r = 11$, $\pi = 3.14$ として $\pi{r}^2$ を計算せよ

と言っているのと同じ。
ただの計算結果だから答えは 379.94 で間違いない。
ここに議論の余地はない。

（国外は知らないが）日本の小学校は何故か代数を教えない。
だから算数の問題もこんな冗長な文章になってしまうし意味のないものに意味を付加する馬鹿野郎も現れる。
ちゃんと小学生の時に代数を教えれば鶴亀算とかワケワカメなものも教えなくて済むのに。

閑話休題。

問題を，文章を糞真面目に解釈して，「（誤差を含めて）円の面積を求めよ」とするなら話は変わる。
なぜなら 11 も 3.14 も「数」ではなく「値」になるから。

円周率は超越数である（今の小学生が「超越数」という言葉を習うかどうかは知らないが）。
ある決まった定数だが，いくらでも桁を大きくとれる。
「半径11」を固定にして円周率の桁を変えて計算してみる。
暗算が苦手なら電卓を使えばいいし，電卓もかったるいというのならプログラムを組めばいい。
私は最近 [Go 言語を勉強している](/golang)ので，これでプログラムを組んでみよう。

```go
package main

import "fmt"

func main() {
	r := 11.0
	pilist := []float64{3.0, 3.1, 3.14, 3.142, 3.1416, 3.14159}
	for _, pi := range pilist {
		fmt.Printf("r = %f , pi = %f , pi*r^2 = %f\n", r, pi, pi*r*r)
	}
}
```

この計算結果を表にして整理してみる。

{{< fig-gen >}}
<table>
  <tr>
    <th>$r$</th>
    <th>$\pi$</th>
    <th>$\pi{r}^2$</th>
  </tr>
  <tr>
    <td rowspan="6">11</td>
    <td>3<br></td>
    <td>363.000000</td>
  </tr>
  <tr>
    <td>3.1</td>
    <td>375.100000</td>
  </tr>
  <tr>
    <td>3.14</td>
    <td>379.940000</td>
  </tr>
  <tr>
    <td>3.142</td>
    <td>380.182000</td>
  </tr>
  <tr>
    <td>3.1416</td>
    <td>380.133600</td>
  </tr>
  <tr>
    <td>3.14159</td>
    <td>380.132390</td>
  </tr>
</table>
{{< /fig-gen >}}

最初の 3 を除いてほぼ 380 に収斂しているのがわかると思う。
ここで「有効桁数」の概念が出てくる。
つまり有効桁数を考慮するなら約 380 または $3.8 \times 10^2$ となる。

一方，半径 11 は測定値である。
ここでは半径を $11.00\pm0.40\,\mathrm{cm}$ として考える（値なら単位がないとね。この範囲なら四捨五入して 11 である）。
この条件で計算結果がどう変わるか同じように計算してみる。
ちなみに円周率は先ほどの計算結果を考慮して大きめに 3.1416 としてみた。

```go
package main

import "fmt"

func main() {
	rlist := []float64{10.6, 10.7, 10.8, 10.9, 11.0, 11.1, 11.2, 11.3, 11.4}
	pi := 3.1416
	for _, r := range rlist {
		fmt.Printf("r = %f , pi = %f , pi*r^2 = %f\n", r, pi, pi*r*r)
	}
}
```

以下が計算結果。

{{< fig-gen >}}
<table>
  <tr>
    <th>$r$</th>
    <th>$\pi$</th>
    <th>$\pi{r}^2$</th>
  </tr>
  <tr>
    <td>10.60<br></td>
    <td rowspan="9">3.1416</td>
    <td>352.990176</td>
  </tr>
  <tr>
    <td>10.70</td>
    <td>359.681784</td>
  </tr>
  <tr>
    <td>10.80</td>
    <td>366.436224</td>
  </tr>
  <tr>
    <td>10.90</td>
    <td>373.253496</td>
  </tr>
  <tr>
    <td>11.00</td>
    <td>380.133600</td>
  </tr>
  <tr>
    <td>11.10</td>
    <td>387.076536</td>
  </tr>
  <tr>
    <td>11.20</td>
    <td>394.082304</td>
  </tr>
  <tr>
    <td>11.30</td>
    <td>401.150904</td>
  </tr>
  <tr>
    <td>11.40</td>
    <td>408.282336</td>
  </tr>
</table>
{{< /fig-gen >}}

同じ「半径11」でもこんなに計算結果が違うのが分かるだろう。

実は半径の誤差[^a] の $\pm0.40\,\mathrm{cm}$ というのはかなり大きい。
詳しい説明は面倒くさいので省くが， $A$, $r$ の関係式

[^a]: ここでいう誤差は平均2乗誤差とみなす。

{{< fig-quote >}}
\[
    A = \pi{r}^2
\]
{{< /fig-quote >}}

の誤差をそれぞれ $r_A$, $r_r$ とすると

{{< fig-quote >}}
\[
    \left(\frac{r_A}{A}\right)^2 = 4\left(\frac{r_r}{r}\right)^2
\]
{{< /fig-quote >}}

の関係があるため[^b]，この場合の面積は $380.1\pm27.6\,\mathrm{cm}^2$ ということになる。
これがもし半径  $11.00\pm0.10\,\mathrm{cm}$ 程度の誤差なら面積は $380.1\pm6.9\,\mathrm{cm}^2$ となる。

[^b]: これを「誤差の伝播」の法則と呼ぶ。

最近の子どもは知らないが，私はこれ（誤差論）を大学に入ってから正式に習った。
「小学生には無理」とは言わないが，有効桁数を含め誤差について論じるのならきちんと手順を踏んで教える必要がある。
「正解」「不正解」でぶった切るなんてもっともやってはいけないことである。

なお，誤差論を習うなら微積分と統計学の基礎知識が必要である。
高校の数学レベルでいけるはず。
あるいは結城浩さんの「数学ガール」シリーズを読む手もある。

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4797382317/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/411g%2BaRmCzL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4797382317/baldandersinf-22/">数学ガールの秘密ノート/微分を追いかけて</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-04-18</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797375698/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797375698.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/479737568X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/479737568X.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797374152/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797374152.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/整数で遊ぼう"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4480095268/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4480095268.09._SCTHUMBZZZ_.jpg"  alt="数学文章作法 推敲編 (ちくま学芸文庫)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/434402740X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/434402740X.09._SCTHUMBZZZ_.jpg"  alt="数学の言葉で世界を見たら 父から娘に贈る数学"  /></a> </p>
<p class="description" >三角関数や微分方程式は物理学，特に天文学を理解するには必須の道具。学校で教えてくれるのを待ってる暇はないのだよ。そして三角関数や微分（と積分）を理解すると理科も数学も抜群に面白くなる。</p>
<p class="gtools" >reviewed by <a href="#maker" class="reviewer">Spiegel</a> on <abbr class="dtreviewed" title="2015-04-30">2015/04/30</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html">G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4320030885/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/41n4-gFkFPL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4320030885/baldandersinf-22/">基礎物理学実験 増訂版</a></dt><dd>下村 健次 </dd><dd>共立出版 1977-10-05</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl>
<p class="description">学生時代，私はコレで実験の基礎を習いました。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-02-21">2016-02-21</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
