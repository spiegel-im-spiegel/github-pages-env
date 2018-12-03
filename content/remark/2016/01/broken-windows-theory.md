+++
date = "2016-01-06T21:04:43+09:00"
description = "結局これは他の犯罪と同じで，「割れ窓」をせっせと補修して回るだけではきりがなく，犯罪を起こす「社会の物質的条件」を変えていくことこそが本命なのである。"
draft = false
tags = ["security", "risk", "management", "botnet"]
title = "ボットネット・テイクダウンと割れ窓理論"

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
  tumblr = ""
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

「[管子](https://ja.wikipedia.org/wiki/%E7%AE%A1%E5%AD%90)」の中の有名な言葉に「倉廩満ちて礼節を知り，衣食足りて栄辱を知る」というのがある。
いわゆる「割れ窓理論（broken windows theory）」は，これと比較すると因果が逆であり，全く以って的外れである[^bw]。
メリケンのインテリよりも古代人の言葉のほうが洞察が深いというのはなかなか皮肉が効いている。

[^bw]: 「割れ窓理論」では米国ニューヨーク市警の事例が有名だろう。1993年から1996年にかけてニューヨークの犯罪発生率が 36% も減少した。しかし当時の警視総監は後の講演で，当時のニューヨークで「ゼロ・トレランス政策」は実施されていなかったと証言した。また当時の米国の主要都市の犯罪発生率は，犯罪対策の有無や警察体制に関わらず，全体的に減少傾向にあった。「ゼロ・トレランス政策とは、大量の人々が周縁に追いやられ、貧困にあえぎ、抑圧されるような社会をつくりだす排除の過程の一部をなすものであり、さらにいえば、それはたんなる保険統計的な処理にすぎないもので、司法というより公共衛生と呼ぶべき政策である」（『[排除型社会](http://www.amazon.co.jp/exec/obidos/ASIN/4903127044/baldandersinf-22/)』より）。

- [エフセキュアブログ : Dridexの解体](http://blog.f-secure.jp/archives/50756643.html)
- [エフセキュアブログ : ボットネットテイクダウン狂想曲](http://blog.f-secure.jp/archives/50757427.html)

「ボットネット・テイクダウン（botnet take down）」とは，詐欺等の犯罪やサイバー攻撃の温床となっている[ボットネット](https://ja.wikipedia.org/wiki/%E3%83%9C%E3%83%83%E3%83%88%E3%83%8D%E3%83%83%E3%83%88)の壊滅を目標として各国警察やセキュリティ企業などを中心に行われている活動を指す。

{{< fig-quote title="国際的なボットネットのテイクダウン作戦" link="https://www.npa.go.jp/cyber/goz/" >}}
<q>本作戦は、関連サーバを押収し、当該ネットワークの管理者を起訴するとともに、より多くの感染端末を特定し、プロバイダ等を通じて感染端末の利用者に対して不正プログラムの駆除を促すことにより、感染端末を減少させることとしている。</q>
{{< /fig-quote >}}

{{< fig-quote title="警察による国際的ボットネットの壊滅作戦" link="http://www.is702.jp/news/1591/partner/101_g/" >}}
<q>この作戦は、インターネットバンキングに係る不正送金事犯に使用されているとみられる不正プログラム「Game Over Zeus」（GOZ）のネットワークを崩壊させる“ボットネットのテイクダウン作戦”と呼ばれています。FBI、ユーロポールが中心となり、日本の警察を含む協力国の法執行機関が連携して、関連サーバを押収し、当該ネットワークの管理者を起訴しています。さらには、感染端末を特定し、プロバイダ等を通じてユーザに駆除を促すことにより、感染端末を減少させることも行っています。</q>
{{< /fig-quote >}}

GOZ のテイクダウン作戦は有名だが，これ以外のボットネットについても順次テイクダウンが行われている。
しかし実際の効果はまちまちで場当たり的な印象がぬぐえず，この点が「割れ窓理論」が引き合いに出される理由なようだ。

テイクダウン作戦のようなやり方は「電撃戦」としては効果があるかもれないが，長期的に継続しても効果が薄くなるだけのように思う。
むしろテイクダウン作戦の長期化はコストの浪費を招き，そのことが攻撃者に意味を与える（テロと同じ）。
その辺が分からないほど警察やセキュリティ企業は馬鹿ではないだろう。
しかし他にもっと効果的な方法がないのだから（今のところは）続けるしかないというのが現状かもしれない。

{{< fig-quote title="ボットネットテイクダウン狂想曲" link="http://blog.f-secure.jp/archives/50757427.html" >}}
<q>それでも私個人的にはボットネットのテイクダウンに賛成派だったりします。賛成というのは、テイクダウンをやれば万事解決という意味ではなく、テイクダウンもボットネット一掃作戦の一部として有効だという意味です。実際、テイクダウンをきっかけとして捜査が進展することはよくあります。</q>
{{< /fig-quote >}}

結局これは他の犯罪と同じで，「割れ窓」をせっせと補修して回るだけではきりがなく，犯罪を起こす「社会の物質的条件」を変えていくことこそが本命なのである。
もっとも，そこまでくれば警察の仕事というより政治家の仕事なのだが。
いや，経済学的センスのない日本の政治家には無理かな（笑）

## 参考文献

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4903127044/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/417iD4x5N%2BL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4903127044/baldandersinf-22/">排除型社会―後期近代における犯罪・雇用・差異</a></dt><dd>ジョック ヤング Jock Young </dd><dd>洛北出版 2007-03</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4791764331/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4791764331.09._SCTHUMBZZZ_.jpg"  alt="後期近代の眩暈―排除から過剰包摂へ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4255008515/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4255008515.09._SCTHUMBZZZ_.jpg"  alt="断片的なものの社会学"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4796700439/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4796700439.09._SCTHUMBZZZ_.jpg"  alt="スティグマの社会学―烙印を押されたアイデンティティ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4791764242/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4791764242.09._SCTHUMBZZZ_.jpg"  alt="新しい貧困 労働消費主義ニュープア"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4062881357/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4062881357.09._SCTHUMBZZZ_.jpg"  alt="弱者の居場所がない社会――貧困・格差と社会的包摂 (講談社現代新書)"  /></a> </p>
<p class="description">「割れ窓理論」についても言及あり。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-15">2015-09-15</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4822283100/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51-pZ52JsUL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4822283100/baldandersinf-22/">セキュリティはなぜやぶられたのか</a></dt><dd>ブルース・シュナイアー 井口 耕二 </dd><dd>日経BP社 2007-02-15</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4881359967/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4881359967.09._SCTHUMBZZZ_.jpg"  alt="暗号の秘密とウソ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797350997/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797350997.09._SCTHUMBZZZ_.jpg"  alt="新版暗号技術入門 秘密の国のアリス"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4594070507/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4594070507.09._SCTHUMBZZZ_.jpg"  alt="チャイナ・ハッカーズ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4105393022/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4105393022.09._SCTHUMBZZZ_.jpg"  alt="暗号解読―ロゼッタストーンから量子暗号まで"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4102159746/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4102159746.09._SCTHUMBZZZ_.jpg"  alt="宇宙創成〈上〉 (新潮文庫)"  /></a> </p>
<p class="description" >日本語のタイトルはアレだが中身は名著。とりあえず読んどきなはれ。</p>
<p class="gtools" >reviewed by <a href="#maker" class="reviewer">Spiegel</a> on <abbr class="dtreviewed" title="2014-09-14">2014/09/14</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html">G-Tools</a>)</p>
</div>
