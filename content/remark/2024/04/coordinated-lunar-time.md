+++
title = "月の標準時を決めよう"
date =  "2024-04-05T20:32:22+09:00"
description = "LTC (Coordinated Lunar Time) というらしい"
image = "/images/attention/kitten.jpg"
tags = [ "astronomy", "ephemeris", "science" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

この時期にこんな話題とか体のいい政治宣伝にしか見えないけど，そういう話も出てるよ，ということで。

- [White House Office of Science and Technology Policy Releases Celestial Time Standardization Policy | OSTP | The White House](https://www.whitehouse.gov/ostp/news-updates/2024/04/02/white-house-office-of-science-and-technology-policy-releases-celestial-time-standardization-policy/)

これによると

{{< fig-quote type="markdown" title="White House Office of Science and Technology Policy Releases Celestial Time Standardization Policy" link="https://www.whitehouse.gov/ostp/news-updates/2024/04/02/white-house-office-of-science-and-technology-policy-releases-celestial-time-standardization-policy/" lang="en" >}}
This policy directs NASA to work with the Departments of Commerce, Defense, State, and Transportation to deliver a strategy for the implementation of LTC no later than December 31, 2026. NASA will also coordinate with other federal agencies as appropriate and international partners through existing international forums, including [Artemis Accords](https://www.nasa.gov/artemis-accords/) partner nations.
{{< /fig-quote >}}

ということで，2026年中に NASA 主導で LTC (Coordinated Lunar Time) の実施に向けた戦略を策定する，とのこと。
なんとも気の長い話である。

地球では標準時として UTC (Coordinated Universal Time; 協定世界時) というのがある。
これは時間の刻みは TAI (International Atomic Time; 国際原子時) に合わせつつ，時刻は UT (Universal Time; 世界時) との差が1秒以内になるよう調整された時刻系である。
ちなみに TAI は座標時系の一種である TT (Terrestrial Time; 地球時) と同義で「地球表面上で歩度SI秒となる一様な時刻」と定義されている[^tt1]。

[^tt1]: 厳密な TT 定義はもう少し複雑で地球重心座標時（TCG）を基にしたジオイド面で考えるのだが，実際にはジオイド面では十分な精度を得られないそうで，ちょっと面倒そうな変換式をかませるようだ。この辺の話は最近の「理科年表」に解説があるので興味のある方はそちらをどうぞ。

座標時系は相対論効果を含んでいるため，当然ながら，どの天体を基準にするかで時間の流れが変わる。
身近な話でいうと，地表の原子時計と GPS 人工衛星の原子時計の間にはほんの僅かな進み遅れがあるし，国際宇宙ステーション（ISS）と地表との間でも差異がある。
もちろん地球と月の間にも差異はある。
人類が宇宙開発を進め地球外に進出するのであれば，こうした差異を精密に測定し各天体ごとの時刻系を決める必要があるというわけだ。

というわけで，最初の紹介記事に戻る。
まぁ，今年の大統領選挙で共和党候補が勝ったらなかったことにされそうな気がする（笑）

## ブックマーク

- [暦の改訂と天文単位の再定義（再掲載）]({{< ref "/remark/2018/12/from-de405-to-de430.md" >}})
- [UNIX 時刻に関する四方山話]({{< ref "/remark/2018/10/unix-time.md" >}})
- [20年ぶりに買った「理科年表」は「けもフレ」とコラボしていた]({{< ref "/remark/2019/12/chronological-scientific-tables-2020-and-kemono-friends.md" >}})

## 参考

{{% review-paapi "4621304259" %}} <!-- 理科年表 2020 -->
{{% review-paapi "4416623410" %}} <!-- 天文年鑑 2024年版 -->
{{% review-paapi "4627275110" %}} <!-- 天体物理学 -->
