+++
title = "暦の改訂と天文単位の再定義（再掲載）"
date = "2018-12-12T15:48:49+09:00"
description = "この記事は本家サイトで 2015-05-02 に書いた「暦の改訂（DE405 から DE430 へ）」 を再構成したものです。"
image = "/images/attention/kitten.jpg"
tags = [ "astronomy", "ephemeris" ]

[scripts]
  mathjax = true
  mermaidjs = false
+++

この記事は[本家サイト](https://baldanders.info/ "Baldanders.info")で 2015-05-02 に書いた「[暦の改訂（DE405 から DE430 へ）](https://baldanders.info/blog/000840/)」 を再構成したものです。

----

暦を決定するには太陽や月などの天体の位置や速度をできるだけ精確に測定することが不可欠である。
国立天文台ではこれらの天体情報をまとめた「暦象年表」を毎年発行している（祝日などをまとめた「年暦要項」は「暦象年表」の抜粋）。

現在「暦象年表」のベースとなっているのが米国 JPL（Jet Propulsion Laboratory; ジェット推進研究所）が惑星探査用に編纂した月・惑星の暦 DE（Development Ephemeris）で，国際的に広く使われている。
2003年から2015年までは [DE405 が採用](http://eco.mtk.nao.ac.jp/koyomi/topics/html/topics2003.html "暦の改訂について（2003） - 国立天文台暦計算室")されてきたが，2016年版からは（2012年 IAU（International Astronomical Union; 国際天文学連合）勧告を組み込んだ） DE430 が採用されることになったらしい。

- [暦の改訂について（2016） - 国立天文台暦計算室](http://eco.mtk.nao.ac.jp/koyomi/topics/html/topics2016_1.html)

DE430 では DE405 以後の観測値が反映されている。
特に惑星観測に関しては惑星探査機による精密な測距観測が行われており，その成果が反映されている。
地球の月についても[レーザ測距（Lunar Laser Ranging; LLR）](http://ja.wikipedia.org/wiki/%E6%9C%88%E3%83%AC%E3%83%BC%E3%82%B6%E3%83%BC%E6%B8%AC%E8%B7%9D%E5%AE%9F%E9%A8%93 "月レーザー測距実験 - Wikipedia")による継続観測や [GRAIL ミッション](http://ja.wikipedia.org/wiki/GRAIL "GRAIL - Wikipedia")などにより精度が向上している。

更に，これが多分いちばんインパクトが大きいと思うが，2012年 IAU 勧告に基づき「天文単位 $\mathrm{au}$（astronomical unit）」が定数として再定義された。
すなわち（時刻系に依らず）

{{< fig-math class="box" >}}
\[
  A = 1\,\mathrm{au} = 149,597,870,700\,\mathrm{m}
  \tag{1}\label{eq:au}
\]
{{< /fig-math >}}

となる（定数なので「誤差」はない）。

もともと「天文単位」は測距観測技術が未熟だった時代に考えだされた単位で「1日（$D=1^{\mathrm{d}}=86400^{\mathrm{s}}$）あたり $0.01720209895\\,\mathrm{rad}$ の角速度で太陽の周りを等速円運動する質点の軌道半径」と定義されていた。
ちなみに $0.01720209895$ はガウス引力定数 $k$ とよばれ，19世紀にガウスが定めたものである。

{{< fig-math class="box" >}}
\[
  k = 0.01720209895
  \tag{2}\label{eq:k}
\]
{{< /fig-math >}}

「1日あたり $k\\,\mathrm{rad}$ の角速度で太陽の周りを等速円運動する質点」の周期を $T$ とすると

{{< fig-math class="box" >}}
\[
  T = \frac{2\pi}{k} \simeq 365.256898^{\mathrm{d}} \simeq 1^{\mathrm{y}}
  \tag{3}\label{eq:T}
\]
{{< /fig-math >}}

となり（さぁ検算しよう！） $1\\,\mathrm{au}$ がほぼ地球の公転軌道半径に等しいことがわかる[^au1]。

[^au1]: 実際，古い（？）理科の教科書には「1天文単位は地球の公転軌道半径と同じ」と説明されているものが多いと思う。ていうか本来の定義が「1天文単位は地球の公転軌道半径と同じ」なのだが，色々あって定義が変遷している。

なんでこんな面倒臭いことになっていたかというと，この定義を使うことにより天文単位の具体的な値が分からなくても日心重力定数 $GM_{S}$ が

{{< fig-math class="box" >}}
\[
  GM_{S} = A^{3} \times \left( \frac{k}{D} \right)^{2} = k^{2} \times \frac{1\,\mathrm{au^{3}}}{1\,\mathrm{day^{2}}} = k^{2}
  \tag{4}\label{eq:GMs}
\]
{{< /fig-math >}}

と一定の値になるのである。
もちろん $GM_{S}$ を SI 単位系に変換するには，何らかの方法で天文単位の値を推定または測定する必要がある。

しかし，現代の測距観測技術の向上により天文単位を考慮せずとも日心重力定数を SI 単位系で直接観測（！）することが可能になった。
また式 ($\ref{eq:GMs}$) の方法では天文単位や日心重力定数の変動が何によるものかが分かりにくく，相対論的な考慮も含まれていない。

そこで，2012年の IAU 勧告で天文単位を定数系に組み入れ，代わりにガウス引力定数 $k$ は外されることとなったのだ。

あぁ，そろそろ新しい天文学の本を調達しないとなぁ。
なにかいい教科書ありません？

## 参考図書

{{% review-paapi "4805202254" %}} <!-- 天体の位置計算 -->
{{% review-paapi "4627275110" %}} <!-- 天体物理学 -->
