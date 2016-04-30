+++
date = "2016-04-30T16:40:25+09:00"
description = "「なぜ、いま「著作権」について考えなければならないのか？」 / Apache Struts 2 の脆弱性 / 「ひとみ（ASTRO-H）」はやっぱりダメらしい / その他の気になる記事"
draft = false
tags = ["code", "intellectual-property", "copyright", "security", "vulnerability", "astronomy", "telescope", "jaxa", "astro-h"]
title = "週末スペシャル： 「なぜ、いま「著作権」について考えなければならないのか？」"

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

1. [「なぜ、いま「著作権」について考えなければならないのか？」]({{< relref "#eu" >}})
1. [Apache Struts 2 の脆弱性]({{< relref "#struts" >}})
1. [「ひとみ（ASTRO-H）」はやっぱりダメらしい]({{< relref "#astro" >}})
1. [その他の気になる記事]({{< relref "#other" >}})

## 「なぜ、いま「著作権」について考えなければならないのか？」{#eu}

- [なぜ、いま「著作権」について考えなければならないのか？―ヨーロッパの現場から | ワールド | 最新記事 | ニューズウィーク日本版 オフィシャルサイト](http://www.newsweekjapan.jp/stories/world/2016/04/post-5010.php)

面白い視点。
EU という「市場」ならではの問題であるともいえる。
まぁでも，所詮は EU 域内での話であって，言い換えれば EU 域外に対してどのように presence を示せるかという話だよね。

{{< fig-quote title="なぜ、いま「著作権」について考えなければならないのか？―ヨーロッパの現場から" link="http://www.newsweekjapan.jp/stories/world/2016/04/post-5010.php" >}}
<q>たとえば、サービス事業者が国ごとに著作権のライセンス契約を結ばなければならないために、あるコンテンツが特定の国では見られるのに別の国では見られない、通称「ジオブロッキング」と呼ばれる問題がある。日本でも、Youtubeなどにアクセスしようとしたときに「この動画はお住まいの地域ではご覧になることができません」という警告が出て、動画にアクセスできなかったりすることがある。ヨーロッパではこれは日常茶飯事の光景だ。ヨーロッパには物理的な国境はないので、日本で言うと、たとえば東京都から神奈川県に移動したらいきなり動画が見られなくなった、と考えるとわかりやすいかもしれない。</q>
{{< /fig-quote >}}

結局は「知的財産権」の世界市場で誰が覇権をとるか。
それでも

{{< fig-quote title="なぜ、いま「著作権」について考えなければならないのか？―ヨーロッパの現場から" link="http://www.newsweekjapan.jp/stories/world/2016/04/post-5010.php" >}}
<q>つまり、「変わらなければいけないのは自明だが、どう変えるかを決めるのが決定的に難しい」、のがEUの著作権法の実情なのである。しかし、この一連の動きによって、「著作権」の本質とはなんなのか、特にコピーが無限にできるという性質をもつデジタルコンテンツに付帯する著作権とはなんなのか、といった問いに対する議論が活発に交わされることになった。</q>
{{< /fig-quote >}}

という感じに議論になるだけまだマシな方だろう。
少なくとも TPP よりは。

## Apache Struts 2 の脆弱性{#struts}

- [JVNVU#91375252: Apache Struts2 に任意のコード実行の脆弱性](http://jvn.jp/vu/JVNVU91375252/)
- [Apache Struts2 の脆弱性対策について(CVE-2016-3081)(S2-032)：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/ciadr/vul/20160427-struts.html)
- [Apache Struts 2 の脆弱性 (S2-032) に関する注意喚起](https://www.jpcert.or.jp/at/2016/at160020.html)
- [Apache Struts2 の脆弱性 (CVE-2016-3081)についてまとめてみた - piyolog](http://d.hatena.ne.jp/Kango/20160427/1461771099)
- [Struts2のDMI(Dynamic Method Invocation)とは何か - Qiita](http://qiita.com/alpha_pz/items/e6b41be70b12174dabda)

CVSSv3 基本評価値は 7.3 (CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:L/I:L/A:L)。

| 基本評価基準                            | 評価値            |
|----------------------------------------:|:------------------|
| 攻撃元区分（AV）                        | ネットワーク（N） |
| 攻撃条件の複雑さ（AC）                  | 低（L）           |
| 必要な特権レベル（PR）                  | 不要（N）         |
| ユーザ関与レベル（UI）                  | 不要（N）         |
| スコープ（S）                           | 変更なし（U）     |
| 情報漏えいの可能性（機密性への影響, C） | 低（L）           |
| 情報改ざんの可能性（完全性への影響, I） | 低（L）           |
| 業務停止の可能性（可用性への影響, A）   | 低（L）           |

対策としては各バージョンの更新版を導入するか DMI (Dynamic Method Invocation) を無効にすればいいようだ。
もっとも 2.3.15.2 以降 DMI は既定で無効らしいので，そのまま無効にしておくのがよろしかろう。
Struts1 への影響は（すでに Apaches のサポート外なので）不明だが， DMI は 2 からの機能と聞いているので関係ないかな？

## 「ひとみ（ASTRO-H）」はやっぱりダメらしい{#astro}

- [JAXA | X線天文衛星ASTRO-H「ひとみ」の今後の運用について](http://www.jaxa.jp/press/2016/04/20160428_hitomi_j.html)

こういうこともあるんだねぇ。
でもまぁ

{{< fig-quote title="Software error doomed Japanese Hitomi spacecraft : Nature News & Comment" link="http://www.nature.com/news/software-error-doomed-japanese-hitomi-spacecraft-1.19835" lang="en" >}}
<q>But Hitomi could still contribute to science. Because of the early failure with Suzaku, Hitomi scientists planned one important early observation. About eight days after launch, Hitomi turned its X-ray gaze on the Perseus cluster, about 250 million light years from Earth. By measuring the speed of gas flowing from the cluster, Hitomi can reveal how the mass of galaxy clusters changes over time as stars are born and die — a test of the crucial cosmological parameter known as dark energy.</q>
{{< /fig-quote >}}

というわけで，初期観測が成果になる可能性もあるわけで，そうすれば「次」に繋がる可能性もあるわけだ。

最近話題になった重力波観測と X 線観測は相補的な役割を果すことが期待されている。
今回の事故で萎縮することなく次へ進むことを祈る。

## その他の気になる記事{#other}

- [キャンペーン · 川内原発を止めてください。 · Change.org](https://www.change.org/p/%E5%B7%9D%E5%86%85%E5%8E%9F%E7%99%BA%E3%82%92%E6%AD%A2%E3%82%81%E3%81%A6%E3%81%8F%E3%81%A0%E3%81%95%E3%81%84)
    - [キャンペーン · 川内原発を止めないでください！ · Change.org](https://www.change.org/p/%E5%B7%9D%E5%86%85%E5%8E%9F%E7%99%BA%E3%82%92%E6%AD%A2%E3%82%81%E3%81%AA%E3%81%84%E3%81%A7%E3%81%8F%E3%81%A0%E3%81%95%E3%81%84)
- [Go 1.6.2 is released - Google グループ](https://groups.google.com/forum/#!topic/golang-announce/8FwSHbMTEjQ)
- [警察庁が「自動走行の公道実証実験」でガイドライン案を公表｜ニュース/アーカイブ｜準天頂衛星システム（QZSS）公式サイト - 内閣府](http://qzss.go.jp/news/archive/npa_160422.html)
- [エフセキュアブログ : Windows Script Hostを無効にする方法](http://blog.f-secure.jp/archives/50766909.html)
- [視覚障碍者プログラマのためのツール | キャリア・働き方 | POSTD](http://postd.cc/tools-of-blind-programmer/)
- [ProtonMail Android BETA v1.3.0 Release Notes - ProtonMail Blog](https://protonmail.com/blog/android-v130/)
    - [ProtonMail iOS BETA v1.2.3 Release Notes - ProtonMail Blog](https://protonmail.com/blog/protonmail-ios-v1-2-3-release-notes/)
- [恐竜が絶滅したのは「隕石の衝突」のせいじゃない：研究結果｜WIRED.jp](http://wired.jp/2016/04/21/dinosaurs-werent-wiped-out-by/)
- [ドイツ原発、燃料棒監視システムにマルウェアが見つかる｜WIRED.jp](http://wired.jp/2016/04/30/german-nuclear-plants-fuel-rod-system-swarming/)

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4757102852/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/41YkbcP5IyL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4757102852/baldandersinf-22/">著作権２．０ ウェブ時代の文化発展をめざして (NTT出版ライブラリー―レゾナント)</a></dt><dd>名和 小太郎 </dd><dd>エヌティティ出版 2010-06-24</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4569812902/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4569812902.09._SCTHUMBZZZ_.jpg"  alt="著作権法がソーシャルメディアを殺す (PHP新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4334037070/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4334037070.09._SCTHUMBZZZ_.jpg"  alt="「ネットの自由」vs.著作権: TPPは、終わりの始まりなのか (光文社新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4087202941/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4087202941.09._SCTHUMBZZZ_.jpg"  alt="著作権とは何か―文化と創造のゆくえ (集英社新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4166608347/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4166608347.09._SCTHUMBZZZ_.jpg"  alt="ビジネスパーソンのための契約の教科書 (文春新書 834)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798119806/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798119806.09._SCTHUMBZZZ_.jpg"  alt="REMIX ハイブリッド経済で栄える文化と商業のあり方"  /></a> </p>
<p class="description">名著です。今すぐ買うべきです。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2014-08-02">2014/08/02</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
