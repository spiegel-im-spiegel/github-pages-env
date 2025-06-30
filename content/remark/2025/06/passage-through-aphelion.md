+++
title = "2025-07-04 に地球が遠日点を通過する話"
date =  "2025-06-30T20:58:43+09:00"
description = "二十四節気も閏年の関係で1〜2日程度の変動が起きるが，それどころじゃないねぇ。"
image = "/images/attention/kitten.jpg"
tags = [ "astronomy", "ephemeris" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
  jsx = false
+++

かねてより「夏至付近ってパッとした天文イベントがないよなぁ」と思っていたが， Bluesky で自 TL を眺めてたときに見つけたこのネタは面白そうだ。

{{< fig-gen >}}
<blockquote class="bluesky-embed" data-bluesky-uri="at://did:plc:bgfky4yz4n46vokquob7kkey/app.bsky.feed.post/3lsqvg3f33222" data-bluesky-cid="bafyreidoazb2wcc4zil6ra2ufhttnbebcbren4ggpb4yyi6kx5vwv3rs34" data-bluesky-embed-color-mode="system"><p lang="ja">2025年7月4日、地球は太陽から最も遠ざかる「遠日点通過」を迎えます。<br>
<br>
「遠日点通過」と太陽に最も近づく「近日点通過」、実は月との関係で、数日単位で動くのだそうです。<br>
eco.mtk.nao.ac.jp/koyomi/wiki/...<br>
<br>
実際にどれくらい動くのか調べてみました。<br>
【参考】ほしぞら情報(国立天文台)<br>
www.nao.ac.jp/astro/sky/<br>
<br>
2020年7月4日<br>
2021年7月6日<br>
2022年7月4日<br>
2023年7月7日<br>
2024年7月5日<br>
2025年7月4日<br>
<br>
二十四節気も、うるう年の関係などで1日前後動きますが、それどころじゃない複雑で大きな動き方をしています。<br><br><a href="https://bsky.app/profile/did:plc:bgfky4yz4n46vokquob7kkey/post/3lsqvg3f33222?ref_src=embed">[image or embed]</a></p>&mdash; 星が好きな人のための新着情報 (<a href="https://bsky.app/profile/did:plc:bgfky4yz4n46vokquob7kkey?ref_src=embed">@localgroupjp.bsky.social</a>) <a href="https://bsky.app/profile/did:plc:bgfky4yz4n46vokquob7kkey/post/3lsqvg3f33222?ref_src=embed">June 29, 2025 at 11:37 PM</a></blockquote><script async src="https://embed.bsky.app/static/embed.js" charset="utf-8"></script>
{{< /fig-gen >}}

「[星が好きな人のための新着情報](https://news.local-group.jp/ "星が好きな人のための新着情報")」さん，いつもお世話になっています。

とういうわけで今回は，地球公転軌道における近日点と遠日点の話。

2025年の近日点と遠日点通過日は以下の通り。

| 日付 | 曜日 | 内容 |
| ---- |:----:| ---- |
| 2025-01-04 | 土 | 地球が近日点通過 |
| 2025-07-04 | 金 | 地球が遠日点通過 |

「えー？ 夏に遠日点を通過するのかよ。しかも冬至や夏至とズレてるし」って思わなかった？ 私は子どもの頃にそう思った（南半球は冬だけど）。

実は冬至や夏至といった二十四節気は、地球の公転面（黄道）に対する自転軸の傾きで決まる。

{{< fig-img-quote src="/remark/2025/02/setsubun/Celestial_sphere_in_Japanese.png" title="Explanation about a celestial sphere in Japanese" link="https://commons.wikimedia.org/wiki/File:Celestial_sphere(in_Japanese).png" lang="en" >}}

一方で近日点と遠日点は地球の公転軌道が関係している。

{{< fig-img-quote src="/remark/2025/06/passage-through-aphelion/A5B1A5D7A5E9A1BCA4CECBA1C2A7kepler1_2.png" title="暦Wiki/ケプラーの法則 - 国立天文台暦計算室" link="https://eco.mtk.nao.ac.jp/koyomi/wiki/A5B1A5D7A5E9A1BCA4CECBA1C2A7.html" >}}

つまり（[ケプラーの法則](https://eco.mtk.nao.ac.jp/koyomi/wiki/A5B1A5D7A5E9A1BCA4CECBA1C2A7.html "暦Wiki/ケプラーの法則 - 国立天文台暦計算室")に基づき）地球の公転軌道を楕円と見立てたときに楕円の焦点にある太陽から最も近づいた瞬間が近日点であり、最も遠ざかった瞬間が遠日点となる。

ちなみに，上の図のように軌道長半径を $a$ 軌道短半径を $b$ としたとき，離心率 $e$ は以下で表される[^e1]。

{{< fig-math >}}
$$
e = \sqrt{1 - \left(\frac{b}{a}\right)^2}
$$
{{< /fig-math >}}

離心率 $e$ が $0$ に近いほど円軌道に近くなる[^e2]。
ちなみに地球公転軌道の離心率は $0.01670$ とされている[^e3]。
まぁ，ほぼ円やね[^e4]。

[^e1]: 厳密には「軌道離心率」と呼び，幾何学で言う離心率とは異なる。この記事では「離心率」で統一する。
[^e2]: 離心率が $0$ で円軌道となり $1$ で放物線軌道となる。さらに $1$ より大きくなると双曲線軌道となる。
[^e3]: ここで挙げた地球公転軌道の離心率は[「天文年鑑」2025年版](https://www.amazon.co.jp/dp/4416723660?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "天文年鑑 2025年版 | 天文年鑑編集委員会 |本 | 通販 | Amazon")のデータを元にしている。この本によると「2025年1月1日0時 TT の平均軌道要素と，これから得られたもの」とある（年の単位は365.25日，J2000.0の黄道と平均春分点準拠）。ちなみに TT は座標時系のひとつである地球時（Terrestrial Time）を指す。また TAI (International Atomic Time; 国際原子時) も座標時系として再定義されていて $TT = TAI + 32.184\\,\mathrm{sec}$ の関係がある。さらに $TT = UTC + 69.184\\,\mathrm{sec}$ (2025年初の時点) である。
[^e4]: 有名なハレー彗星は76年周期で回帰するとされているが（次の近日点通過は2061年8月23日頃），離心率は $0.967786$ と見積もられている。2017年に発見された恒星間天体 ['Oumuamua](https://science.nasa.gov/solar-system/comets/oumuamua/ "'Oumuamua - NASA Science") は離心率が $1$ より大きい双曲線軌道である。

太陽と地球だけを考えれば概ねこれで OK だけど，実際には地球の周りには月が周回している。
なので地球-月の間の重心で考える必要がある[^o1]。
これにより太陽-地球-月の位置関係により（地球から見た）近日点および遠日点の通過日が微妙に変わるらしい。

[^o1]: 大きなタイムスケールで見れば太陽系の他の天体との相互作用も考慮する必要があるが，この記事では割愛する。

{{< fig-img-quote src="/remark/2025/06/passage-through-aphelion/CFC7C0B12FB6E1C6FCC5C0C4CCB2E1C3CFB5E5B7EEBDC5BFB4.png" title="暦Wiki/惑星/近日点通過 - 国立天文台暦計算室" link="https://eco.mtk.nao.ac.jp/koyomi/wiki/CFC7C0B12FB6E1C6FCC5C0C4CCB2E1.html" >}}

最初に紹介したポストにあるように2020年以降に限っても数日単位で遠日点通過日がシフトしている。

{{< fig-quote title="2025年7月4日、地球は太陽から最も遠ざかる「遠日点通過」を迎えます..." link="https://bsky.app/profile/localgroupjp.bsky.social/post/3lsqvg3f33222" >}}
2020年7月4日<br>
2021年7月6日<br>
2022年7月4日<br>
2023年7月7日<br>
2024年7月5日<br>
2025年7月4日<br>
{{< /fig-quote >}}

結構激しいな。
二十四節気も閏年の関係で1〜2日程度の変動が起きるが，それどころじゃないねぇ。

というわけで，近日点・遠日点の話でした。
これで，こっぽし。

## ブックマーク

- [「暦」日本史 （再掲載）]({{< ref "/remark/2015/japanese-koyomi.md" >}})
- [『猫暦』を読んだ]({{< ref "/remark/2016/05/nekoyomi.md" >}})
- [冬至に関する与太話]({{< ref "/remark/2017/12/winter-solstice.md" >}})
- [「新暦七夕」なるものは存在しない]({{< ref "/remark/2019/06/traditional-tanabata.md" >}})
- [立春も動き出す]({{< ref "/remark/2021/01/the-beginning-of-spring.md" >}})
- [お彼岸]({{< ref "/remark/2022/03/ohigan.md" >}})
- [第五の季節：土用]({{< ref "/remark/2022/04/doyo-period.md" >}})
- [春夏秋冬は「四季」ではない？]({{< ref "/remark/2025/03/four-seasons.md" >}})

{{< fig-youtube id="97vI07JmqGE" title="【🌟 星学手簡】江戸時代の天文学者の手紙が重要文化財に！？ VTuberと辿る暦の歴史【星見まどか / 国立天文台】 - YouTube" >}}

{{< fig-youtube id="zsiAknJHcos" title="【🌟 10時間目】江戸時代の天文学🔭望遠鏡や地動説はいつから？【#スクールオブチューブ江戸時代講座 / 星見まどか】 - YouTube" >}}

## 参考

{{% review-paapi "4416723660" %}} <!-- 天文年鑑 2025年版 -->
{{% review-paapi "4627275110" %}} <!-- 天体物理学 -->
{{% review-paapi "4254102372" %}} <!-- 暦の大事典 -->
{{% review-paapi "4805202254" %}} <!-- 天体の位置計算 -->
{{% review-paapi "B01BHGVLOY" %}} <!-- 猫暦 -->
