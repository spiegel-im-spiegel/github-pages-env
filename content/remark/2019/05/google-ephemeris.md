+++
title = "カレンダーに祝日を入れたいなら国立天文台へ行けばいいじゃない"
date =  "2019-05-14T22:37:13+09:00"
description = "日本の暦はこれが国家公式データである。"
image = "/images/attention/kitten.jpg"
tags = [ "astronomy", "ephemeris", "japanese", "calendar" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[前回のネタ]({{< relref "./tokyo-2020.md" >}})は個人的に色々と不本意だったので，今回は軽い小ネタで。
よく考えたらこの記事が今年初めての天文ネタだよ `orz`

- [Google カレンダーに日本の祝祭日を表示したい！ カレンダーを追加する方法 - 窓の杜](https://forest.watch.impress.co.jp/docs/serial/chrometips/1184245.html)

という記事を見かけたのだが，日本の祝日は毎年国立天文台で発表してるんだから国立天文台のデータを使えばいいじゃない，と思ったり。

日本の暦は[国立天文台]の[暦計算室]で見ることができる。
こんな感じのページ。

{{< fig-img src="./google-ephemeris.png" title="今月のこよみ powered by Google Calendar - 国立天文台暦計算室" link="https://eco.mtk.nao.ac.jp/koyomi/cande/calendar.html" width="824" >}}

しかも Google カレンダーへインポート可能で，上のページの「＋Google Calendar」の部分をクリックすれば自身の Google カレンダーにインポートできる。

祝日だけでなく二十四節気や朔望月，あるいは日食・月食といった情報が取得できる。
必要に応じてオン・オフを切り替えればいいだろう。

なお，上記データで見れるのは「祝日および休日」で「祭日」のデータはない。
つか，祭日は国家が規定するものではないので，あるわけないのだが。
祭日が知りたければ[高島暦](https://www.amazon.co.jp/exec/obidos/ASIN/B07JKP4CG8/baldandersinf-22/ "平成31年神宮館高島暦 | 神宮館編集部 | 占い | Kindleストア | Amazon")とかを購入することをオススメする（笑）

ちなみに国立天文台では毎年2月1日に翌年の「[暦要項](https://eco.mtk.nao.ac.jp/koyomi/yoko/ "暦要項 - 国立天文台暦計算室")」が官報で発表される。

- [平成31（2019）年暦要項の発表 | 国立天文台(NAOJ)](https://www.nao.ac.jp/news/topics/2018/20180201-rekiyoko.html)
- [平成32（2020）年暦要項の発表 | 国立天文台(NAOJ)](https://www.nao.ac.jp/news/topics/2019/20190201-rekiyoko.html)

日本の暦はこれが国家公式データである。

[国立天文台]: https://www.nao.ac.jp/ "国立天文台(NAOJ)"
[暦計算室]: https://eco.mtk.nao.ac.jp/koyomi/ "国立天文台 天文情報センター 暦計算室"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E5%B9%B3%E6%88%9031%E5%B9%B4%E7%A5%9E%E5%AE%AE%E9%A4%A8%E9%AB%98%E5%B3%B6%E6%9A%A6-%E7%A5%9E%E5%AE%AE%E9%A4%A8%E7%B7%A8%E9%9B%86%E9%83%A8-ebook/dp/B07JKP4CG8?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B07JKP4CG8"><img src="https://images-fe.ssl-images-amazon.com/images/I/51wFsGhTlrL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E5%B9%B3%E6%88%9031%E5%B9%B4%E7%A5%9E%E5%AE%AE%E9%A4%A8%E9%AB%98%E5%B3%B6%E6%9A%A6-%E7%A5%9E%E5%AE%AE%E9%A4%A8%E7%B7%A8%E9%9B%86%E9%83%A8-ebook/dp/B07JKP4CG8?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B07JKP4CG8">平成31年神宮館高島暦</a></dt>
	<dd>神宮館編集部</dd>
    <dd>神宮館 2018-07-10 (Release 2018-10-22)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B07JKP4CG8</dd>
  </dl>
  <p class="description">Kindle 版があるのか。思ったよりデジタル化が進んでるんだな（笑） 昔は親が毎年買っていたのだが近年は実家で見かけないな。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-05-14">2019-05-14</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4805202254?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/51mQCyP04rL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4805202254?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">天体の位置計算</a></dt>
    <dd>長沢 工 (著)</dd>
    <dd>地人書館 1985-09-01</dd>
    <dd>単行本</dd>
    <dd>4805202254 (ASIN), 9784805202258 (EAN), 4805202254 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">B1950.0 分点から J2000.0 分点への過渡期に書かれた本なので情報が古いものもあるが，基本的な内容は位置天文学の教科書として充分通用する。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4627275110?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/51UOq7TlGyL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4627275110?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">天体物理学</a></dt>
    <dd>Arnab Rai Choudhuri (著), 森 正樹 (翻訳)</dd>
    <dd>森北出版 2019-05-28</dd>
    <dd>単行本</dd>
    <dd>4627275110 (ASIN), 9784627275119 (EAN), 4627275110 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">興味本位で買うにはちょっとビビる値段なので図書館で借りて読んでいる。まえがきによると，この手のタイプの教科書はあまりないらしい。内容は非常に堅実で分かりやすい。理系の学部生レベルなら問題なく読めるかな。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-11-13">2019-11-13</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
