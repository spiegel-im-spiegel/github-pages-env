+++
title = "「ヒト」こそがセキュリティの最強点"
date =  "2019-05-30T21:51:57+09:00"
description = "こんな15年くらい前の腐ったセキュリティ事例を目の当たりにできるとは思わなかったが，「他山の石」とでも思って今後の私達の活動に活かしていきたいものである。"
image = "/images/attention/kitten.jpg"
tags = [ "security", "risk", "management" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今夜はリリースされたばかりの『[賢者の弟子を名乗る賢者 11](https://www.amazon.co.jp/exec/obidos/ASIN/B07RWSCKKG/baldandersinf-22/)』でも読んでのんびり過ごそうと思っていたのだが，面白そうな話が上がっているので便乗してみる。

- [ヤマダ電機通販サイトの不正アクセスについてまとめてみた - piyolog](https://piyolog.hatenadiary.jp/entry/2019/05/30/063417)

さすが [piyolog](https://piyolog.hatenadiary.jp/) さんは早いなぁ。

この記事を見かけたのが朝の忙しい時間帯だったこともあり，タイトルだけを見て「やらかしてら」くらいの感想しかなかったのだが，どうも「ヤマダ電機はクレカのセキュリティ・コードまで保持っていて丸ごとそれらを抜かれたらしい」という噂になっていたようだ。

が，実際にはこれはマスメディアの報道からくる誤解らしい。

- [ヤマダ電機のクレジットカード情報流出事件、「セキュリティコードを保存」の誤解広がる(篠原修司) - 個人 - Yahoo!ニュース](https://news.yahoo.co.jp/byline/shinoharashuji/20190530-00127999/)

例によってマスメディアの中途半端で杜撰な報道をネットメディア（Twitter 等を含む）が増幅しフェイク化するといういつもの展開なわけだが，語るのも面倒なので，こちらは無視する。

問題はヤマダ電機側がインシデントを認知してから実際にクレカ運用を止めるまでに10日かかっていること，そこから警察に届けるまで（例の10連休を挟んで）10日，そこからユーザに告知するまでに20日以上もかかっていることだ。

{{< fig-quote type="md" title="ヤマダ電機通販サイトの不正アクセスについてまとめてみた" link="https://piyolog.hatenadiary.jp/entry/2019/05/30/063417" >}}
|        日時         | 出来事                                                                                         |
|:-------------------:| ---------------------------------------------------------------------------------------------- |
| 2019年3月18日までに | ヤマダウェブコム・ヤマダモールが不正アクセスを受け、ペイメントアプリケーションが改ざんされる。 |
|         ：          | カード情報流出の可能性がある対象期間                                                           |
|    2019年4月16日    | クレジットカード会社より情報流出の疑義について連絡                                             |
|        同日         | ヤマダ電機が情報漏えいの可能性を把握。                                                         |
|         ：          | カード情報流出の可能性がある対象期間                                                           |
|    2019年4月26日    | ヤマダウェブコム・ヤマダモールでシステムメンテナンス。カード登録を休止する措置。               |
|        同日         | ヤマダ電機がP.C.F.FRONTEOへ不正アクセスの調査を依頼。                                          |
|    2019年5月7日     | ヤマダ電機が警察へ被害相談。                                                                   |
|    2019年5月20日    | P.C.F.FRONTEOによる調査が完了。                                                                |
|    2019年5月22日    | ヤマダ電機が警察へ被害届を提出。                                                               |
|    2019年5月28日    | ヤマダ電機が個人情報保護委員会へ不正アクセス被害を報告。                                       |
|    2019年5月29日    | ヤマダ電機が不正アクセス被害とクレジットカード情報流出の可能性を発表。                         |
{{< /fig-quote >}}

いやいやいや。
あり得ない愚鈍さである。
今時こんなテキトーなインシデント・レスポンスをかます企業があるとは思わざりき。

上のタイムラインを見れば分かるが，情報漏洩の可能性を認知した4月16日の時点でサービスを止めユーザに告知した上で警察と調査会社との三者で事に当たれば被害（の可能性）を抑え少なくとも10連休前には解決できただろう。
しかも10連休直前にえんやらやっと調査を依頼して自分たちは10連休を楽しんだ後に警察へ通報かい？ 安倍政権の中の人たち，あなた達の進める「働き方改革」は順調なようですよ（笑）

「ヒトはセキュリティの最弱点である」とはよく言われることだが，まさに今回のケースはヤマダ電機側の怠慢による「人災」だろう（もちろんクレカ情報を盗んだ犯罪者が一番悪いのは言うまでもないが）。

しかし本当は「ヒトはセキュリティの最強点になり得る」のである（「である」ではなく「なり得る」のがポイント）。
何故なら，予期しないイレギュラーが起きた時，危機的状況に見舞われた時，そういったときにこそ「ヒト」の真価が問われるからだ。

こんな15年くらい前の腐ったセキュリティ事例を目の当たりにできるとは思わなかったが，「他山の石」とでも思って今後の私達の活動に活かしていきたいものである。

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%82%BB%E3%82%AD%E3%83%A5%E3%83%AA%E3%83%86%E3%82%A3%E3%81%AF%E3%81%AA%E3%81%9C%E3%82%84%E3%81%B6%E3%82%89%E3%82%8C%E3%81%9F%E3%81%AE%E3%81%8B-%E3%83%96%E3%83%AB%E3%83%BC%E3%82%B9%E3%83%BB%E3%82%B7%E3%83%A5%E3%83%8A%E3%82%A4%E3%82%A2%E3%83%BC/dp/4822283100?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4822283100"><img src="https://images-fe.ssl-images-amazon.com/images/I/51-pZ52JsUL._SL160_.jpg" width="107" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%82%BB%E3%82%AD%E3%83%A5%E3%83%AA%E3%83%86%E3%82%A3%E3%81%AF%E3%81%AA%E3%81%9C%E3%82%84%E3%81%B6%E3%82%89%E3%82%8C%E3%81%9F%E3%81%AE%E3%81%8B-%E3%83%96%E3%83%AB%E3%83%BC%E3%82%B9%E3%83%BB%E3%82%B7%E3%83%A5%E3%83%8A%E3%82%A4%E3%82%A2%E3%83%BC/dp/4822283100?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4822283100">セキュリティはなぜやぶられたのか</a></dt>
	<dd>ブルース・シュナイアー</dd>
	<dd>井口 耕二 (翻訳)</dd>
    <dd>日経BP社 2007-02-15</dd>
    <dd>Book 単行本</dd>
    <dd>ASIN: 4822283100, EAN: 9784822283100</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">原書のタイトルが “<a href="https://www.amazon.co.jp/Beyond-Fear-Thinking-Sensibly-Uncertain-ebook/dp/B000PY3NB4?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B000PY3NB4">Beyond Fear: Thinking Sensibly About Security in an Uncertain World</a>” なのに対して日本語タイトルがどうしようもなくヘボいが中身は名著。とりあえず読んどきなはれ。ゼロ年代当時 9.11 およびその後の米国のセキュリティ政策と深く関連している内容なので，そのへんを加味して読むとよい。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-02-11">2019-02-11</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E5%BE%A9%E6%B4%BB%E3%81%AE%E5%9C%B0%EF%BC%91-%E5%B0%8F%E5%B7%9D%E4%B8%80%E6%B0%B4-ebook/dp/B00GJOESS6?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00GJOESS6"><img src="https://images-fe.ssl-images-amazon.com/images/I/51ymtvyHUmL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E5%BE%A9%E6%B4%BB%E3%81%AE%E5%9C%B0%EF%BC%91-%E5%B0%8F%E5%B7%9D%E4%B8%80%E6%B0%B4-ebook/dp/B00GJOESS6?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00GJOESS6">復活の地１</a></dt>
	<dd>小川一水</dd>
    <dd>早川書房 2012-09-15 (Release 2013-11-15)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00GJOESS6</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">コミカライズ版もある。てか，コミカライズ版を最初に読んだ（笑） 大きな災害がある度にこの作品を思い出す。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-04-17">2016-04-17</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%82%BD%E3%83%BC%E3%82%B7%E3%83%A3%E3%83%AB%E3%82%B7%E3%83%95%E3%83%88-%E3%81%93%E3%82%8C%E3%81%8B%E3%82%89%E3%81%AE%E4%BC%81%E6%A5%AD%E3%81%AB%E3%81%A8%E3%81%A3%E3%81%A6%E4%B8%80%E7%95%AA%E5%A4%A7%E5%88%87%E3%81%AA%E3%81%93%E3%81%A8-%E6%96%89%E8%97%A4-%E5%BE%B9-ebook/dp/B009S7CDP6?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B009S7CDP6"><img src="https://images-fe.ssl-images-amazon.com/images/I/51b1VUql4DL._SL160_.jpg" width="107" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%82%BD%E3%83%BC%E3%82%B7%E3%83%A3%E3%83%AB%E3%82%B7%E3%83%95%E3%83%88-%E3%81%93%E3%82%8C%E3%81%8B%E3%82%89%E3%81%AE%E4%BC%81%E6%A5%AD%E3%81%AB%E3%81%A8%E3%81%A3%E3%81%A6%E4%B8%80%E7%95%AA%E5%A4%A7%E5%88%87%E3%81%AA%E3%81%93%E3%81%A8-%E6%96%89%E8%97%A4-%E5%BE%B9-ebook/dp/B009S7CDP6?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B009S7CDP6">ソーシャルシフト　これからの企業にとって一番大切なこと</a></dt>
	<dd>斉藤 徹</dd>
    <dd>日本経済新聞出版社 2011-11-11 (Release 2012-10-18)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B009S7CDP6</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">随分前に出版された本だが，企業がソーシャル・メディアと付き合うための基本的な事柄が載っている。これ読んで出直してきなはれ。ちなみに巻末の spesial thanks に私の名前が載っているのは密かな自慢である（笑）</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-05-30">2019-05-30</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E8%B3%A2%E8%80%85%E3%81%AE%E5%BC%9F%E5%AD%90%E3%82%92%E5%90%8D%E4%B9%97%E3%82%8B%E8%B3%A2%E8%80%85-11-GC%E3%83%8E%E3%83%99%E3%83%AB%E3%82%BA-%E3%82%8A%E3%82%85%E3%81%86%E3%81%9B%E3%82%93%E3%81%B2%E3%82%8D%E3%81%A4%E3%81%90-ebook/dp/B07RWSCKKG?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B07RWSCKKG"><img src="https://images-fe.ssl-images-amazon.com/images/I/51nMoOn8kkL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E8%B3%A2%E8%80%85%E3%81%AE%E5%BC%9F%E5%AD%90%E3%82%92%E5%90%8D%E4%B9%97%E3%82%8B%E8%B3%A2%E8%80%85-11-GC%E3%83%8E%E3%83%99%E3%83%AB%E3%82%BA-%E3%82%8A%E3%82%85%E3%81%86%E3%81%9B%E3%82%93%E3%81%B2%E3%82%8D%E3%81%A4%E3%81%90-ebook/dp/B07RWSCKKG?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B07RWSCKKG">賢者の弟子を名乗る賢者 11 (GCノベルズ)</a></dt>
	<dd>りゅうせんひろつぐ, 藤ちょこ</dd>
    <dd>マイクロマガジン社 2019-05-30 (Release 2019-05-30)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B07RWSCKKG</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">ソウルハウル編のクライマックス。今度はレイドボスだ！</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-05-30">2019-05-30</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>
