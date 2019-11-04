+++
description = "キーワードは「公開されたプライバシー」である。つまり，公開された私的情報であってもプライバシー権は失われることはなく，そのコントロールは情報へのアクセス性にある，という考え方だ。"
date = "2016-11-05T11:46:10+09:00"
title = "「ネットが蝕むプライバシー」を読む（再掲載）"
tags = [
  "internet",
  "privacy",
  "book",
  "grigori",
]

[scripts]
  mathjax = false
  mermaidjs = false
+++

この記事は[本家サイト](https://baldanders.info/ "Baldanders.info")で2008年に公開した「[「ネットが蝕むプライバシー」を読む](https://baldanders.info/blog/000415/)」を加筆・修正して再掲載したものです。
なお，本文中で紹介している記事「[プライバシーに無分別な若者]」は今年出た[別冊日経サイエンス『サイバーセキュリティ』](http://www.nikkei-science.com/page/sci_book/bessatu/51212.html)に「プライバシーの終焉？」と原題に近いタイトルで再録されています。
興味のある方は是非どうぞ。

そして「公開されたプライバシー」というキーワードに興味を持っていただけたなら，次に Danah Boyd さんの『[つながりっぱなしの日常を生きる]』を読むことを強くお勧めします。

- [『つながりっぱなしの日常を生きる』を読む]({{< ref "/remark/2016/05/its-complicated.md" >}})

----

[『日経サイエンス』 2008年12月号]の特集記事「ネットが蝕むプライバシー」を読む。
『日経サイエンス』とプライバシー特集ってなんか毛色が違う感じがしたが，この特集記事ではプライバシーの問題をコンピュータ･サイエンスとして扱っているようだ。
したがって，内容的にはコンピュータやネットワークのセキュリティに関する記事が多い（のでプライバシー問題の特集だと思って読むと拍子抜けするかもしれない）。
この分野に関心のある方は手にとって読んでみることをお薦めする。
個人的には最初の「プライバシー2.0を考える[^1]」以外は面白かった。

[^1]: 「[サイバーセキュリティ]」には「プライバシー2.0を考える」は収録されていない。さすがに内容がアレだったか（笑） その代わりなのか Jaron Lanier さんの「[プライバシーをどう考えるべきか（How Should We Think about Privacy ?）](http://www.nikkei-science.com/201411_086.html)」が収録されている。こちらはお勧め。

一番面白かったのは「[プライバシーに無分別な若者]」という記事。
結構アレなタイトルだが

{{< fig-quote title="プライバシーに無分別な若者" link="http://www.nikkei-science.com/page/magazine/0812/200812_088.html" >}}
<q>多くの若者が，日常生活のきわめて私的な内容を<br />
ソーシャルネットワーキングのウェブサイトで公開している<br />
これはプライベートとパブリックの境界線が大きく変化する前触れだ</q>
{{< /fig-quote >}}

という要約文が示す通り特定の世代をバッシングする記事ではない（原題は “The End of Privacy?” なので，これは日本語タイトルが悪い）。
この記事は Facebook を例にして書かれているため日本ではピンと来ないかもしれないが，底にあるものは共通していると思う。

キーワードは「公開されたプライバシー」である。
つまり，公開された私的情報であってもプライバシー権は失われることはなく，そのコントロールは情報へのアクセス性にある，という考え方だ。

もともとのプライバシーは空間概念と一体になっている。
つまり私的領域と公的領域を区別することである。
しかしインターネット（特に Web 2.0 以降）は情報は空間概念から切り離される。
場所が不定になるのだ。
だから，それをコントロールしたければ情報へのアクセス性を対象にせざるを得なくなる。
感覚的には分かっていても，改めて「公開されたプライバシー」という言葉を与えられると少し見方が変わってくる。

例えば（2008年当時の）旬の GSV 問題だが，これまで私は「GSV にはプライバシー問題はない」と書いてきた[^2]。
日本は私的領域と公的領域の区別が曖昧なことが多く日常生活が（たとえそれが公道であっても）生活道路にまで漏れ出てしまう（日本以外はどうかというのは[高木浩光さんがまとめておられる](http://takagi-hiromitsu.jp/diary/20081026.html#p01 "高木浩光＠自宅の日記 - 住宅ストリートビューの国際比較 アメリカ・フランス・日本, 修正（27日）")）。
これを「公開されたプライバシー」の問題として捉えるのなら前言は撤回しなければならない。

[^2]: 「[Street View のアレ](https://baldanders.info/blog/000404/)」および「[Street View のアレ 2](https://baldanders.info/blog/000406/)」。そういえば当時 GSV を使った窃盗団がどうとかほざいていた人がいたが，あの人は今（笑）

あと，ちょっとぎょっとなったのが特集記事のあとの茂木健一郎さんと木村忠正さんの対談記事（茂木健一郎さんのこの連載記事自体はあまり面白くないんだけどね）。
この部分も少し引用してみる。

{{< fig-quote title="『日経サイエンス』 2008年12月号 「日本のネット文化を変えるには」" link="http://www.nikkei-science.com/page/magazine/0812/200812_088.html" >}}
<q>日本，韓国，フィンランドで行った調査項目で「インターネット空間というのはパブリックなものか，プライベートなものか」と質問しました。
するとパブリックだという答えは日本の男性がいちばん少ない。
フィンランドでは男女とも8割以上がパブリックだと答えていますが，日本の男性では38.6%と4割を切っています。</q> （p.99-100）
{{< /fig-quote >}}

まぁ調査方法とか分からないので鵜呑みにしていいかどうか分からないけど，これってやっぱりケータイというデバイスの影響とかあるのだろうか。
WIRED VISION に「[実はネットブックだった『iPhone』？――平均以下所得層の購入が急増](http://wired.jp/2008/10/31/%e5%ae%9f%e3%81%af%e3%83%8d%e3%83%83%e3%83%88%e3%83%96%e3%83%83%e3%82%af%e3%81%a0%e3%81%a3%e3%81%9f%e3%80%8eiphone%e3%80%8f%ef%bc%9f%e2%80%95%e2%80%95%e5%b9%b3%e5%9d%87%e4%bb%a5%e4%b8%8b%e6%89%80/)」という記事があるけど，向こうで iPhone のみでネットに繋がる人が増えるとネットを公衆空間だと思わない人が増えたりするのだろうか。

[『日経サイエンス』 2008年12月号]: http://www.nikkei-science.com/page/magazine/200812.html "2008年12月号 | 日経サイエンス"
[プライバシーに無分別な若者]: http://www.nikkei-science.com/page/magazine/0812/200812_088.html "プライバシーに無分別な若者 | 日経サイエンス"
[サイバーセキュリティ]: http://www.nikkei-science.com/page/sci_book/bessatu/51212.html "サイバーセキュリティー | 日経サイエンス"
[つながりっぱなしの日常を生きる]: https://www.amazon.co.jp/exec/obidos/ASIN/B0125TZSZ0/baldandersinf-22/ "Amazon.co.jp: つながりっぱなしの日常を生きる：ソーシャルメディアが若者にもたらしたもの 電子書籍: ダナ・ボイド, 野中 モモ: Kindleストア"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%82%B5%E3%82%A4%E3%83%90%E3%83%BC%E3%82%BB%E3%82%AD%E3%83%A5%E3%83%AA%E3%83%86%E3%82%A3-%E5%88%A5%E5%86%8A%E6%97%A5%E7%B5%8C%E3%82%B5%E3%82%A4%E3%82%A8%E3%83%B3%E3%82%B9-%E6%97%A5%E7%B5%8C%E3%82%B5%E3%82%A4%E3%82%A8%E3%83%B3%E3%82%B9%E7%B7%A8%E9%9B%86%E9%83%A8/dp/4532512123?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4532512123"><img src="https://images-fe.ssl-images-amazon.com/images/I/51gurnOqhiL._SL160_.jpg" width="120" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%82%B5%E3%82%A4%E3%83%90%E3%83%BC%E3%82%BB%E3%82%AD%E3%83%A5%E3%83%AA%E3%83%86%E3%82%A3-%E5%88%A5%E5%86%8A%E6%97%A5%E7%B5%8C%E3%82%B5%E3%82%A4%E3%82%A8%E3%83%B3%E3%82%B9-%E6%97%A5%E7%B5%8C%E3%82%B5%E3%82%A4%E3%82%A8%E3%83%B3%E3%82%B9%E7%B7%A8%E9%9B%86%E9%83%A8/dp/4532512123?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4532512123">サイバーセキュリティ (別冊日経サイエンス)</a></dt>
    <dd>日経サイエンス編集部 (編集)</dd>
    <dd>日本経済新聞出版社 2016-04-22 (Release 2016-04-22)</dd>
    <dd>ムック</dd>
    <dd>4532512123 (ASIN), 9784532512125 (EAN)</dd>
    <dd>Rating<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">『日経サイエンス』2012年6月号の「介入されないもうひとつのウェブ」が収録されている。その他にも2010年代前半におけるセキュリティ問題についてよくまとめられている。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-11-05">2016-11-05</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home" >PA-API</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/B0125TZSZ0?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/616sjle5ITL._SL160_.jpg" width="109" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/B0125TZSZ0?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">つながりっぱなしの日常を生きる：ソーシャルメディアが若者にもたらしたもの</a></dt>
    <dd>ダナ・ボイド (著), 野中 モモ (翻訳)</dd>
    <dd>草思社 2014-10-09 (Release 2015-07-21)</dd>
    <dd>Kindle版</dd>
    <dd>B0125TZSZ0 (ASIN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">読むのに1年半以上かかってしまった。ネット，特に SNS 上で自身のアイデンティティやプライバシーを保つにはどうすればいいか。豊富な事例を交えて考察する。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-31">2018-12-31</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
