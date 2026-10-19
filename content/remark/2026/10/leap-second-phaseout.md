+++
title = "国際度量衡総会，閏秒やめるってよ"
date =  "2026-10-19T15:45:35+09:00"
description = "エンジニア的には，今後，閏秒の度にサーバに張り付いて監視しなくていいというのは朗報だろう。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "leap-second", "astronomy", "engineering", "calendar" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
  jsx = false
+++

先週の話であるが，10月15日の第28回国際度量衡総会（Conférence générale des poids et mesures; CGPM）において UTC (協定世界時) の閏秒運用を実質廃止することが決議された。

- [「うるう秒」、2027年5月に実質廃止へ 国際会議が決議](https://japan.cnet.com/article/35253693/)

決議の内容がこれ：

- [Résolution 5 (2026) - BIPM](https://www.doi.org/10.59161/CGPM2026RES5F)

オリジナルはフランス語だが，[英語版](https://www.bipm.org/en/cgpm-2026/resolution-5)もある。
今回は英語版をベースにして紹介する。

## 閏秒廃止の経緯

閏秒に関する議論は国立天文台の「暦Wiki」にまとめられている。

- [暦Wiki/時刻/うるう秒がなくなる？ - 国立天文台暦計算室](https://eco.mtk.nao.ac.jp/koyomi/wiki/BBFEB9EF2FA4A6A4EBA4A6C9C3A4ACA4CAA4AFA4CAA4EBA1A9.html)

これによると国際天文学連合（International Astronomical Union; IAU）では随分前から議論があったが，自分たちでは決められないということになったらしい。

{{< fig-quote type="markdown" title="うるう秒がなくなる？" link="https://eco.mtk.nao.ac.jp/koyomi/wiki/BBFEB9EF2FA4A6A4EBA4A6C9C3A4ACA4CAA4AFA4CAA4EBA1A9.html" >}}
- IAUのUTC作業部会でもうるう秒廃止については意見がまとまらず ([IERS](https://hpiers.obspm.fr/eop-pc/earthor/utc/report_WG_UTC_2014.pdf))、
  - 廃止については、賛成も反対もできない。
  - 廃止なら、地球の自転とリンクする"Universal Time"(世界時)という名称を使うべきではない。
  - 廃止なら、対応のため、最低5年の猶予を設けること。
{{< /fig-quote >}}

その後，舞台は国際電気通信連合 無線通信部門 (International Telecommunication Union (ITU) Radiocommunication Sector; ITU-R) へ移った。
しかし，ここでも意見がまとまらず，2015年の世界無線通信会議（WRC-15）でも色々と案は提示されるものの，結論を出せず2023年に先送りになった。

{{< fig-quote type="markdown" title="うるう秒がなくなる？" link="https://eco.mtk.nao.ac.jp/koyomi/wiki/BBFEB9EF2FA4A6A4EBA4A6C9C3A4ACA4CAA4AFA4CAA4EBA1A9.html" >}}
- 2015-11-19 結局、うるう秒や[協定世界時](https://eco.mtk.nao.ac.jp/koyomi/wiki/B6A8C4EAC0A4B3A6BBFE.html)といった現在のシステムは変更せず、現在及び将来の時刻系について更なる研究を行い、2023年の世界無線通信会議(WRC-23)に提案することになりました。
  - 記者発表 ([ITU](https://www.itu.int/net/pressoffice/press_releases/2015/53.aspx#.Vk6G-HbhBB8))
  - [WRC-15 Resolution 655](https://www.itu.int/dms_pub/itu-r/oth/0C/0A/R0C0A00000F00135PDFE.pdf)
  - 情報通信研究機構によるまとめ ([NICT](https://www.nict.go.jp/publication/shuppan/kihou-journal/houkoku65-2_HTML/2019S-07-02(06-02).pdf))
{{< /fig-quote >}}

ここでさらに舞台は国際度量衡総会（CGPM）へ移った。
2020年に ITU と国際度量衡局（Bureau International des Poids et Mesures; BIPM）は UTC 関連の相互協力協定を締結。
2022年の第27回 CGPM で

- 2035年までに $|UT1 - UTC|$ の許容値を0.9秒から引き上げること
- ITU 等の関連機関とも協議の上，新しい許容値は，少なくとも1世紀は UTC の調整が不要となるよう定めること

を決めたらしい。
これを受けて，2023年の世界無線通信会議（WRC-23）では，2022年の CGPM の決議を支持し<q>許容値や適用年については、2026年のCGPMに向けて準備を進める</q>こととなった。

## 2026年 CGPM 決議の内容

そもそも，地球の自転速度は長期的には徐々に遅くなっているのだが，どうも地球温暖化の影響で一時的に自転速度が速まってるらしい。

{{< fig-quote type="markdown" title="「うるう秒は時代遅れ？」" link="/remark/2025/01/should-we-abandon-the-leap-second/" >}}
加速している理由は地球温暖化により極地で地殻を押さえつけていた氷が融けて地球の形状が（より球に近く）変わったせいなんだとか。でもこれは一時的なもので，融解した水が赤道付近に集まれば再び自転速度は遅くなるそうな。
{{< /fig-quote >}}

このままだと2035年までに負の閏秒（1秒を引く）が必要になる確率が30%に達するそうな。
負の閏秒はこれまで経験がなく，可用性リスクがかなり高いと考えられている。

{{< fig-quote type="markdown" title="Resolution 5 (2026) - BIPM" link="https://www.bipm.org/en/cgpm-2026/resolution-5" lang="en" >}}
- a negative leap second has not previously been applied, and it is considered to pose a high risk of causing anomalies and disruption to critical infrastructures that are largely unprepared, and the preparation would create, for the industries that rely on time synchronization, a need for financial investment, whose amount is estimated to be similar to their preparations for the millennium bug,
- the growing influence of IoT (Internet of Things) technologies, dispersed multi-sensor systems, smart grids, and virtual services for contemporary society, creates an increasing risk of malfunction in operation of systems and services with the insertion of any leap second, especially with a negative one,
- the use of different mitigation methods for leap second introductions by global service providers in their server systems, in order to avoid time steps, leads to inconsistency in timing between different systems and in the disseminated time and frequency, which do not conform to the SI unit,
{{< /fig-quote >}}

そこで今回の決定である。

{{< fig-quote type="markdown" title="Resolution 5 (2026) - BIPM" link="https://www.bipm.org/en/cgpm-2026/resolution-5" lang="en" >}}
- continuous UTC will become effective on 20 May 2027,
- the maximum value for the difference |UT1-UTC| will be 3 600 seconds (1 hour), ensuring the long term continuity for UTC for several centuries,
{{< /fig-quote >}}

つまり $|UT1 - UTC|$ の許容値を具体的に1時間（3,600秒）に引き上げることを決め，実施開始日を2027年5月20日とすることにしたようだ（continuous UTC ってのは正式名称になるのかね？）。

1時間を超えた場合にどうするかについては何も決まっていない。
でも，今後 UT とのギャップをスルーし続けたとしてもその差は100年に1分程度のスケールらしいので，1時間の許容値であれば上述の引用の通り数世紀は余裕で保つだろうと思われる。
たぶん，数世紀後なら地球の自転を基準とした時刻系は意味が無くなるんじゃないかな（希望的観測）。
むしろ数世紀後に1日や1年の長さが再定義される可能性すらある（恐竜時代は1日の長さが23.5時間程度だったらしい）。

今後はこの決定が妥当かどうか $|UT1 - UTC|$ の値を監視していくことになる。
まぁ，天文学に於いては UT と TAI (国際原子時) とのギャップを観測することは重要なので，今まで通りっちゃあ今まで通りか。

エンジニア的には，今後，閏秒の度にサーバに張り付いて監視しなくていいというのは朗報だろう。

## 参考文献

{{< linkcard "6b8ffab8684f95758ee5d25b952682638ba7889e" >}} <!-- 天文年鑑 2026年版 -->
{{% review-paapi "B0D47FY7Q3" %}} <!-- 「日経サイエンス」2025年2月号 -->
{{% review-paapi "4621304259" %}} <!-- 理科年表 2020 -->
{{% review-paapi "4627275110" %}} <!-- 天体物理学 -->
