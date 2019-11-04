+++
date = "2016-01-06T21:04:43+09:00"
description = "結局これは他の犯罪と同じで，「割れ窓」をせっせと補修して回るだけではきりがなく，犯罪を起こす「社会の物質的条件」を変えていくことこそが本命なのである。"
tags = ["security", "risk", "management", "botnet"]
title = "ボットネット・テイクダウンと割れ窓理論"

[scripts]
  mathjax = false
  mermaidjs = false
+++

「[管子](https://ja.wikipedia.org/wiki/%E7%AE%A1%E5%AD%90)」の中の有名な言葉に「倉廩満ちて礼節を知り，衣食足りて栄辱を知る」というのがある。
いわゆる「割れ窓理論（broken windows theory）」は，これと比較すると因果が逆であり，全く以って的外れである[^bw]。
メリケンのインテリよりも古代人の言葉のほうが洞察が深いというのはなかなか皮肉が効いている。

[^bw]: 「割れ窓理論」では米国ニューヨーク市警の事例が有名だろう。1993年から1996年にかけてニューヨークの犯罪発生率が 36% も減少した。しかし当時の警視総監は後の講演で，当時のニューヨークで「ゼロ・トレランス政策」は実施されていなかったと証言した。また当時の米国の主要都市の犯罪発生率は，犯罪対策の有無や警察体制に関わらず，全体的に減少傾向にあった。「ゼロ・トレランス政策とは、大量の人々が周縁に追いやられ、貧困にあえぎ、抑圧されるような社会をつくりだす排除の過程の一部をなすものであり、さらにいえば、それはたんなる保険統計的な処理にすぎないもので、司法というより公共衛生と呼ぶべき政策である」（『[排除型社会](https://www.amazon.co.jp/exec/obidos/ASIN/4903127044/baldandersinf-22/)』より）。

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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4903127044?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41uBRNdBygL._SL160_.jpg" width="110" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4903127044?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">排除型社会―後期近代における犯罪・雇用・差異</a></dt>
    <dd>ジョック ヤング (著), Young,Jock (原著), 秀男, 青木 (翻訳), 泰郎, 伊藤 (翻訳), 政彦, 岸 (翻訳), 真保呂, 村澤 (翻訳)</dd>
    <dd>洛北出版 2007-03-01</dd>
    <dd>単行本</dd>
    <dd>4903127044 (ASIN), 9784903127040 (EAN), 4903127044 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description"><a href="https://baldanders.info/blog/000410/">感想はこちら</a>。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-07">2018-12-07</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4822283100?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/51-pZ52JsUL._SL160_.jpg" width="107" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4822283100?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">セキュリティはなぜやぶられたのか</a></dt>
    <dd>ブルース・シュナイアー (著), 井口 耕二 (翻訳)</dd>
    <dd>日経BP 2007-02-15</dd>
    <dd>単行本</dd>
    <dd>4822283100 (ASIN), 9784822283100 (EAN), 4822283100 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">原書のタイトルが “<a href="https://www.amazon.co.jp/dp/B000PY3NB4?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">Beyond Fear: Thinking Sensibly About Security in an Uncertain World</a>” なのに対して日本語タイトルがどうしようもなくヘボいが中身は名著。とりあえず読んどきなはれ。ゼロ年代当時 9.11 およびその後の米国のセキュリティ政策と深く関連している内容なので，そのへんを加味して読むとよい。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-02-11">2019-02-11</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
