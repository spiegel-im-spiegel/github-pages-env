+++
title = "本当に夏至は「太陽が一番高い日」か？"
date =  "2026-06-21T23:55:27+09:00"
description = "本当に？"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "astronomy", "ephemeris", "calendar", "koyomi" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
  jsx = false
+++

その辺の Web ページとかを見ると夏至の説明として「太陽の南中高度が一番高い日」あるいはもっと簡単に「太陽が一番高い日」と書かれていることが多い。
それを見るたびに「違うだろ！」とツッコみそうになるが，定義はともかく，現象としては間違ってないので思いとどまっている。

いや，待て。
本当に夏至は「太陽が一番高い日」なのか？

## 二十四節気と七十二候

まずは定義から。

夏至は二十四節気のひとつである。
もともと二十四節気は冬至を起点に1年を24分割したもので，更に二十四節気をそれぞれ3分割したものを七十二候と呼ぶ。
なので，二十四節気や七十二候は本来期間を指すもので，ある時点を指すものではない。

七十二候は明治の改暦で廃止になったが，二十四節気については少し定義を変えて現在でも公式に運用されている[^st1]。
このうち夏至については **「太陽黄経が90度となる瞬間を含む日」** と定義されている。

[^st1]: 二十四節気は2016年にユネスコの無形文化遺産に[登録](https://ich.unesco.org/en/RL/the-twenty-four-solar-terms-knowledge-in-china-of-time-and-practices-developed-through-observation-of-the-suns-annual-motion-00647 "The Twenty-Four Solar Terms, knowledge in China of time and practices developed through observation of the sun’s annual motion - UNESCO Intangible Cultural Heritage")されたそうな。

{{< fig-img-quote src="/remark/2025/02/setsubun/Celestial_sphere_in_Japanese.png" title="Explanation about a celestial sphere in Japanese" link="https://commons.wikimedia.org/wiki/File:Celestial_sphere(in_Japanese).png" lang="en" >}}

## 太陽位置の座標系を変換する

ある時点の太陽高度を求めるには天球[^cs1] 上の太陽の位置（黄道座標系[^ecs1] および赤道座標系[^ecs2]）を観測者から見た位置（地平座標系[^hcs1]）に変換する必要がある。

[^cs1]: 「天球」とは地球を中心とした無限遠の仮想球面である。「無限遠」というのは具体的な長さを指すものではない。3次元空間上の座標を表すには通常は3つのパラメータが必要だが（方位・高度・距離 など），天球を想定することで地球と天体との「距離」のパラメータを省くことができる。これにより天球上の天体を2つのパラメータ（方位と高度あるいは赤経と赤緯）のみで表すことが可能になる。つまり天球は観測者から見た天体の2次元写像であると見なすこともできる。もちろんこれは地球と天体との距離が（人間の認知レベルでは）十分遠いが故に可能な想定である。しかし現代天文学においてアインシュタインの業績や観測技術の飛躍的な向上により「天体までの距離」が重要な意味を持つようになったのは，皆さんご承知のとおりである。余談だが，神話的には神々は「天球」の外側に存在すると考えられている。中世の星座絵が反転してるのは「神の視点」を意識しているからである。3次元の「無限遠の球」の外側に神が居るとか面白すぎる！
[^ecs1]: 地球の太陽公転面（厳密には地球-月重心の平均軌道角運動量ベクトルに垂直な面）を黄道と呼び，これを基準にした座標系を黄道座標系 (ecliptic coordinate system) と呼ぶ。地球の座標系のように黄経 (longitude; $\lambda$)・黄緯 (latitude; $\beta$)で表す。黄道に対して天の赤道は約 $23.4\tcdegree$ ほど傾いていて（公転面に対する地軸の傾き，黄道傾斜角），黄道と天の赤道が交わる点が春分点および秋分点となる。現代天文学では春分点が基準になっていて，$春分点 = 黄経\\,0\tcdegree = 赤経\\,0\tcdegree$ である。太陽黄経は太陽の位置を黄道座標系で表したときの黄経の値を指す。
[^ecs2]: 地球の北極・南極・赤道を延長した方向を，それぞれ天の北極・天の南極・天の赤道と呼ぶ。このとき地球の経緯度と同様に，天体の位置を赤経 (right ascension; $\alpha$)・赤緯 (declination; $\delta$) で表すことができる。これを赤道座標系 (equatorial coordinate system) と呼ぶ。赤緯は $天の赤道 = 0\tcdegree$ を基準に、北または南へ $90\tcdegree$ までの角度で表される。
[^hcs1]: 観測者から見て真上の方向を天頂 (zenith) と天頂を通り南北を結ぶ大円（子午線）を基準にした座標系を地平座標系 (horizontal coordinate system) と呼ぶ。地平座標系では東西南北の方位 (azimuth) と高度 (altitude または height) で表すことができる。ちなみにここで言う高度は地平面からの仰角を指し，地平面が $0\tcdegree$，天頂が $90\tcdegree$ となる。太陽が子午線上を通過する瞬間を南中と呼ぶ。南中時の太陽高度が南中高度である。南半球は太陽が北側にあるから南中じゃないな（笑） 正確に「極上正中 (upper culmination)」と呼ぶべきか。この記事では南中で押し通す。

黄道座標系 ($\lambda, \beta$) と赤道座標系 ($\alpha, \delta$) との関係は，黄道傾斜角を $\epsilon$ として，以下の通りである。

{{< fig-img-quote src="/remark/2026/06/is-summer-solstice-the-highest-sun/BAC2C9B8B7CF2FC0D6C6BBBAC2C9B8B7CFA4C8B2ABC6BBBAC2C9B8B7CFB2ABC6BBA4C8C0D6C6BB.png" title="暦Wiki/座標系/赤道座標系と黄道座標系 - 国立天文台暦計算室" link="https://eco.mtk.nao.ac.jp/koyomi/wiki/BAC2C9B8B7CF2FC0D6C6BBBAC2C9B8B7CFA4C8B2ABC6BBBAC2C9B8B7CF.html" >}}

{{< fig-math >}}
\begin{align*}
\cos\delta\cos\alpha &= \cos\beta\cos\lambda \\
\cos\delta\sin\alpha &= \cos\epsilon\cos\beta\sin\lambda - \sin\epsilon\sin\beta \\
\sin\delta &= \sin\epsilon\cos\beta\sin\lambda + \cos\epsilon\sin\beta
\end{align*}
{{< /fig-math >}}

太陽は常に黄緯 $β_\mathrm{sun} = 0\tcdegree$ と見なせるので

{{< fig-math >}}
\begin{align*}
\sin\delta_\mathrm{sun} &= \sin\epsilon\sin\lambda_\mathrm{sun}
\end{align*}
{{< /fig-math >}}

となる。
つまり（$\epsilon$ が変化しないなら[^e1]）太陽黄経 $\lambda_\mathrm{sun} = 90\tcdegree = 夏至$ のとき赤緯 $\delta_\mathrm{sun} = \epsilon$ で最大値になる。

[^e1]: 厳密には黄道傾斜角は章動の影響で毎日 $0''.05$ くらい変化しているようだ。が，太陽黄緯の変化に比べればごく僅かと言える。

さらに赤道座標系 ($\alpha, \delta$) と地平座標系 ($A, h$) との関係は，天の北極の高度（＝観測地点の北緯）を $\phi$，天体の子午線からの時角を $H$ として，以下の通りである。

{{< fig-img-quote src="/remark/2026/06/is-summer-solstice-the-highest-sun/BAC2C9B8B7CF2FC3CFCABFBAC2C9B8B7CFA4C8C0D6C6BBBAC2C9B8B7CFhorequ.png" title="暦Wiki/座標系/地平座標系と赤道座標系 - 国立天文台暦計算室" link="https://eco.mtk.nao.ac.jp/koyomi/wiki/BAC2C9B8B7CF2FC3CFCABFBAC2C9B8B7CFA4C8C0D6C6BBBAC2C9B8B7CF.html" >}}

{{< fig-math >}}
\begin{align*}
\cos h\sin A &= -\cos\delta\sin H \\
\cos h\cos A &= \cos\phi\sin\delta - \sin\phi\cos\delta\cos H \\
\sin h &= \sin\phi\sin\delta + \cos\phi\cos\delta\cos H
\end{align*}
{{< /fig-math >}}

この式を元に季節ごとに太陽の方位と高度を視覚化したグラフが以下である。

{{< fig-img-quote src="/remark/2026/06/is-summer-solstice-the-highest-sun/topics2005a.png" title="太陽の高度，方位角および影の位置の概略値の求め方 - 国立天文台暦計算室" link="https://eco.mtk.nao.ac.jp/koyomi/topics/html/topics2005.html" width="627" >}}

おー。
アナレンマだ。

{{< fig-img-quote src="/remark/2022/05/mathgirl-note-circular-functions/11815465285_58c30d7bf2_e.jpg" title="Analemma of the Sun | Flickr (György Soponyai, CC-BY-NC)" link="https://www.flickr.com/photos/vanamonde81/11815465285" lang="en" >}}

## 夏至の日の南中高度（北半球）

前節の式より，ある時点の太陽の南中高度は，もっと簡単に以下の式で表される。

{{< fig-math >}}
\[
h_\mathrm{sun} = 90\,\tcdegree - \phi + \delta_\mathrm{sun}
\]
{{< /fig-math >}}

ここまでくれば（観測地点が同じであれば）夏至の瞬間に太陽の南中高度が最大になることが納得できるだろう。

ただし，沖ノ鳥島など北回帰線（北緯約23.4度）より南では夏至の南中高度が $90\tcdegree$ を超える。
つまり天頂より北側に太陽が来る。
この場合は夏至の日の南中高度が最大とは言えなくなる。

いや，待て待て待て。

これは夏至（$\lambda_\mathrm{sun} = 90\tcdegree$）の瞬間に太陽が南中した場合の話であって，夏至の日の南中高度が最大になるとは限らないのでは？

こういうときは AI に教えてもらおう！

というわけで [Kagi Assistant] にお願いしてみた。
最初は過去10年の夏至前後の南中高度を調べてもらいたかったのだが，なんか無理っぽい感じだったので，ピンポイントに，松江市における2019年の夏至（2019-06-22 00:54 JST）前後1日で[調べて](https://eco.mtk.nao.ac.jp/koyomi/dni/2019/s3306.html "日の出入り＠松江(島根県) 令和元年(2019)06月 - 国立天文台暦計算室")もらった。

| 日付 | 南中時刻 | 南中高度[°] |
|---|---:|---:|
| 2019-06-21 | 12:09 | 78.0° |
| 2019-06-22 | 12:10 | 78.0° |
| 2019-06-23 | 12:10 | 78.0° |

同じく松江市における2023年の夏至（2023-06-21 23:58 JST）前後1日に[ついて](https://eco.mtk.nao.ac.jp/koyomi/dni/2023/s3306.html "日の出入り＠松江(島根県) 令和 5年(2023)06月 - 国立天文台暦計算室")も以下に示す。

| 日付 | 南中時刻 | 南中高度[°] |
|---|---:|---:|
| 2023-06-20 | 12:09 | 78.0° |
| 2023-06-21 | 12:09 | 78.0° |
| 2023-06-22 | 12:10 | 78.0° |

なるほど。
夏至前後の数日程度では南中高度はほとんど変わらないのか。
もっと有効桁数を大きく取れれば違いが分かるのかもしれないが，公開されているデータではこれが限度かな[^cande1]。
先程の太陽の方位と高度の図を見ても真南を中心に対称になっているっぽいし，太陽の最大南中高度が夏至の日から外れることはないのかも？

[^cande1]: 国立天文台から発行されている[暦象年表](https://eco.mtk.nao.ac.jp/koyomi/cande/ "暦象年表 - 国立天文台暦計算室")には地球時0時の太陽視赤緯が $0.01''$ の精度で掲載されているが PDF 版は2024年以降しか公開されていないようだ（少なくとも2019年と2023年のはなかった）。紙であれば[国会図書館にある](https://ndlsearch.ndl.go.jp/search?cs=bib&collapse=none&display=panel&from=0&size=20&f-tid=R100000002-I000000050761 "NDLサーチ | 国立国会図書館")が，ネットでは閲覧できないんだよなぁ。{{<pdf-file title="2026年の暦象年表" link="https://eco.mtk.nao.ac.jp/koyomi/cande/pdf/cande2026.pdf" >}} で夏至（2026-06-21 17:25 JST）の前後1日の太陽視赤緯を見てみると，1日に $10'' \simeq 0.0028\tcdegree$ 程度の変化量になっている。東京における日南中高度のデータもあるが毎日ではなく精度も $1' \simeq 0.017\tcdegree$。このため夏至前後1日の値は分からない。まぁ，視赤緯の変化量が $10''$ 程度なら $1'$ 精度の南中高度データでは変化は分かるまい。

## というわけで

とりあえず北回帰線以北という条件付きで，夏至は「太陽の南中高度が一番高い日」と言ってもよさそうだ。

やー，遊んだ遊んだ（笑）

## ブックマーク

- [暦Wiki/季節/二十四節気とは？ - 国立天文台暦計算室](https://eco.mtk.nao.ac.jp/koyomi/wiki/B5A8C0E12FC6F3BDBDBBCDC0E1B5A4A4C8A4CFA1A9.html)
- [暦Wiki/季節/七十二候 - 国立天文台暦計算室](https://eco.mtk.nao.ac.jp/koyomi/wiki/B5A8C0E12FBCB7BDBDC6F3B8F5.html)
- [暦Wiki/黄道座標系 - 国立天文台暦計算室](https://eco.mtk.nao.ac.jp/koyomi/wiki/B2ABC6BBBAC2C9B8B7CF.html)
- [暦Wiki/赤道座標系 - 国立天文台暦計算室](https://eco.mtk.nao.ac.jp/koyomi/wiki/C0D6C6BBBAC2C9B8B7CF.html)
- [暦Wiki/地平座標系 - 国立天文台暦計算室](https://eco.mtk.nao.ac.jp/koyomi/wiki/C3CFCABFBAC2C9B8B7CF.html)
- [太陽の南中高度はどうやって計算する？ | 国立天文台(NAOJ)](https://www.nao.ac.jp/faq/a0109.html)
- [太陽の南中時刻はどうすればわかる？ | 国立天文台(NAOJ)](https://www.nao.ac.jp/faq/a0108.html)

※ 日本独自の暦・暦象に関する話題は [koyomi]({{< rlnk "/tags/koyomi/" >}}) を参照のこと。

[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"

## 参考

{{% linkcard "6b8ffab8684f95758ee5d25b952682638ba7889e" %}} <!-- 天文年鑑 2026年版 -->

{{% review-paapi "4621304259" %}} <!-- 理科年表 2020 -->
{{% review-paapi "4627275110" %}} <!-- 天体物理学 -->
{{% review-paapi "4805202254" %}} <!-- 天体の位置計算 -->
