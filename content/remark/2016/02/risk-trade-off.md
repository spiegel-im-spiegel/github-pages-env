+++
date = "2016-02-03T22:02:21+09:00"
description = "リスクと利便性はトレードオフできない。リスクとトレードオフできるのはリスクのみ。リスクをトレードオフして「全体最適化」することが重要なのである。"
draft = false
tags = ["security", "risk", "management"]
title = "リスク認知とトレードオフ"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"
+++

繰り返し語られていることではあるけれど。

- {{< pdf-file title="“...no one can hack my mind”: Comparing Expert and Non-Expert Security Practices" link="https://www.usenix.org/system/files/conference/soups2015/soups15-paper-ion.pdf" >}}
- [Google Japan Blog: サイバーセキュリティ月間：エキスパートたちのセキュリティ対策](http://googlejapan.blogspot.jp/2016/02/blog-post.html)
- [セキュリティの専門家とそうでない人のセキュリティ対策の違いとは、Googleが論文紹介 -INTERNET Watch](http://internet.watch.impress.co.jp/docs/news/20160202_741876.html)

面白いのはこの表。

{{< fig-quote title="エキスパートたちのセキュリティ対策" link="http://googlejapan.blogspot.jp/2016/02/blog-post.html" >}}
<q>本論文では、オンラインにおける安全性を確保するために行っている対策について、セキュリティ専門家 231 人と一般のインターネットユーザー 294 人に対して行った調査結果から、異なる 2 つのグループからの回答をもとに、それぞれの対策や判断の違いと理由をまとめたものです。<br>
以下は、最も回答が多かったセキュリティ対策を、エキスパートと一般ユーザーで分類した表です。</q>
{{< /fig-quote >}}

{{< fig-gen title="「エキスパートたちのセキュリティ対策」より" link="http://googlejapan.blogspot.jp/2016/02/blog-post.html" >}}
<table>
  <tr>
    <th>一般ユーザ</th>
    <th>順位</th>
    <th>セキュリティ・エキスパート</th>
  </tr>
  <tr>
    <td>アンチウイルスソフトを使う</td>
    <th>1</th>
    <td>ソフトウェアアップデートをインストールする</td>
  </tr>
  <tr>
    <td>強固なパスワードを使う</td>
    <th>2</th>
    <td>サービスごとにユニークなパスワードを使う</td>
  </tr>
  <tr>
    <td>頻繁にパスワードを変える</td>
    <th>3</th>
    <td>2段階認証を使う</td>
  </tr>
  <tr>
    <td>知ってるウェブサイトしか利用しない</td>
    <th>4</th>
    <td>強固なパスワードを使う</td>
  </tr>
  <tr>
    <td>個人情報は公開しない</td>
    <th>5</th>
    <td>パスワード管理ツールを使う</td>
  </tr>
</table>
{{< /fig-gen >}}

これとよく似た表があったな，と思ったら吉川肇子さんの『[リスクとつきあう](https://www.amazon.co.jp/exec/obidos/ASIN/4641280304/baldandersinf-22/)』という本の中にあった。
それは「危険と感じられる活動や科学技術」を30項目挙げ順位をつけたものであるが，全部を載せるのは無理なので，大学生のワースト3とトップ3を挙げ，それに対する専門家の順位を並べてみる。

{{< fig-gen title="吉川肇子著『リスクとつきあう』 p.77" >}}
<table>
  <tr>
    <th>活動や科学技術</th>
    <th>大学生順位</th>
    <th>専門家順位</th>
  </tr>
  <tr>
    <td>原子力</td>
    <td>1</td>
    <td>20</td>
  </tr>
  <tr>
    <td>拳銃</td>
    <td>2</td>
    <td>4</td>
  </tr>
  <tr>
    <td>喫煙</td>
    <td>3</td>
    <td>2</td>
  </tr>
  <tr>
    <td>動力草刈り機</td>
    <td>28</td>
    <td>28</td>
  </tr>
  <tr>
    <td>予防接種</td>
    <td>29</td>
    <td>25</td>
  </tr>
  <tr>
    <td>水泳</td>
    <td>30</td>
    <td>10</td>
  </tr>
</table>
{{< /fig-gen >}}

ちなみに専門家の1位は「自動車」である。
もっともこの表は1987年の論文のもので現在とは知見が異なると思うけど[^aa]，注目すべきは一般の人と専門家とのリスク「感覚」のギャップである。

[^aa]: たとえば専門家の間では長い間，飛行機よりも自動車のほうがリスクが高いとみなされていたが，近年では飛行機のほうがリスクが高いのではないかという意見もある。飛行機事故の生起確率に対して事故時の被害の大きさのほうがずっと大きいかららしい。

リスクは「生起確率×脅威の大きさ」で測るが，私を含め一般の人はこの「感覚」に慣れていない。
何らかの認知バイアスが必ずかかってしまう。
それがこのギャップとなって顕われるのだと思う。

だからこの表を見て「専門家の言うとおりにしなきゃ」と考えるのは早計で[^a]，自身のリスク「感覚」を磨いていくことが重要である。
本当はそういうのは義務教育のうちにやっとくべきなんだろうけどね。

[^a]: いや専門家の意見を参考にするのはいいことだと思うけど，今時は専門家もピンキリだからなぁ。

さらに言うなら，個人と営利企業と国家とではリスクに対してそれぞれ執るべき行動が異なる。
何故なら，それぞれの「守るべきもの」が異なるから。

個人にとってまっさきに重要なのはプライバシーと財産だろう。
営利企業は個人と似ているが，賢い企業はリスクをお金に換算できる。
故に「リスクはチャンス」なのだ。
そして国家にとっては「軍事」を含めた包括的なものであり，特定の個人や企業との利害には（基本的には）興味が無い[^b]。

[^b]: つまり安倍首相が「国民の安心安全のために」というときの「国民」は個人を指しているわけではない。そしてこれらの思惑は時に衝突する。

そういえば最近こんな記事があった。

- [エフセキュアブログ : 今年度のCTFを無双した韓国チーム、その強さの秘密](http://blog.f-secure.jp/archives/50762469.html)

国家にとって「サイバー・セキュリティ」の防衛力の高さは「サイバー攻撃」力の高さに直結する（たとえそれを行使する気がなくとも）。
だから各国は「サイバー・セキュリティ」の向上を目指さざるをえない。
「サイバー・セキュリティ」防衛力の高さを誇示することが他国に対する示威行動になる。
この記事の最後の文章を挙げておこう。

{{< fig-quote title="今年度のCTFを無双した韓国チーム、その強さの秘密" link="http://blog.f-secure.jp/archives/50762469.html" >}}
<q>日本で〇〇万人のセキュリティ技術者が不足しているという報道があるのと同じように、韓国でも〇〇万人(韓国では10万人説が主流)のセキュリティ技術者が不足しているとマスコミは報道したがるそうです。<br>
それでも中の人は、サイバーセキュリティは数ではなく質だということをきちんとわかっていて、韓国トップクラスの技術者が世界トップクラスになるための工夫を続けています。</q>
{{< /fig-quote >}}

ホンマ，日本年金機構の不始末にかこつけて「セキュリティ戦略」を利権化しようという意図は分かりやすいが，単に動員に頼るんじゃなく，本当の意味での「戦略」をきちんと企図して実行しないと，まぁた世界から取り残されるよ。

- [サイバー対策強化へ法案＝政府 - WSJ](http://jp.wsj.com/articles/JJ11757514300003864480318027226602632645960)

愚痴はともかく，ただ政府やセキュリティ企業や（自称を含む）専門家に従えばいいというわけではなく，自分にとって何が大事で何を守るべきかを客観的に把握できていなければならない。
「彼を知り己を知れば百戦殆からず」である。
孫子だっけ？

最後にリスクと利便性はトレードオフできない。
リスクとトレードオフできるのはリスクのみ。
リスクをトレードオフして「全体最適化」することが重要なのである[^c]。
故に優先順位がとても重要になる。
そこで最初の表に戻るわけだ。

[^c]: だからリスクをお金に換算するというのは理に適っていると言える。

「辞書攻撃」で簡単に解読できるようなパスワードを「強固」と思い込んで必死で暗記したのに，何かの拍子で忘れてしまってサービスが使えなくなった上に攻撃者に乗っ取られる。
なんてお馬鹿なことにならないように。
使える道具はきっちり使って人間の弱さや曖昧さを許容しつつ排除するよう構成するなら，個人にとっての「サイバー・セキュリティ」はそれほど難しいことではない。

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
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E7%92%B0%E5%A2%83%E3%83%AA%E3%82%B9%E3%82%AF%E5%AD%A6-%E4%B8%AD%E8%A5%BF%E6%BA%96%E5%AD%90-ebook/dp/B00E7HMI7U?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00E7HMI7U"><img src="https://images-fe.ssl-images-amazon.com/images/I/51I9C7cFl2L._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E7%92%B0%E5%A2%83%E3%83%AA%E3%82%B9%E3%82%AF%E5%AD%A6-%E4%B8%AD%E8%A5%BF%E6%BA%96%E5%AD%90-ebook/dp/B00E7HMI7U?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00E7HMI7U">環境リスク学</a></dt>
	<dd>中西準子</dd>
    <dd>日本評論社 (Release 2013-08-01)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00E7HMI7U</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">環境リスクのみならず「リスク」全体に目配せした良書。著者の自伝的作品でもある。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-02-03">2016-02-03</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%AA%E3%82%B9%E3%82%AF%E3%81%A8%E3%81%A4%E3%81%8D%E3%81%82%E3%81%86%E2%80%95%E5%8D%B1%E9%99%BA%E3%81%AA%E6%99%82%E4%BB%A3%E3%81%AE%E3%82%B3%E3%83%9F%E3%83%A5%E3%83%8B%E3%82%B1%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3-%E6%9C%89%E6%96%90%E9%96%A3%E9%81%B8%E6%9B%B8-%E5%90%89%E5%B7%9D-%E8%82%87%E5%AD%90/dp/4641280304?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4641280304"><img src="https://images-fe.ssl-images-amazon.com/images/I/519S1SM2S4L._SL160_.jpg" width="110" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%AA%E3%82%B9%E3%82%AF%E3%81%A8%E3%81%A4%E3%81%8D%E3%81%82%E3%81%86%E2%80%95%E5%8D%B1%E9%99%BA%E3%81%AA%E6%99%82%E4%BB%A3%E3%81%AE%E3%82%B3%E3%83%9F%E3%83%A5%E3%83%8B%E3%82%B1%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3-%E6%9C%89%E6%96%90%E9%96%A3%E9%81%B8%E6%9B%B8-%E5%90%89%E5%B7%9D-%E8%82%87%E5%AD%90/dp/4641280304?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4641280304">リスクとつきあう―危険な時代のコミュニケーション (有斐閣選書)</a></dt>
	<dd>吉川 肇子</dd>
    <dd>有斐閣 2000-03</dd>
    <dd>Book 単行本</dd>
    <dd>ASIN: 4641280304, EAN: 9784641280304</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">リスク・コミュニケーションについて。内容は古いがまだまだ使える。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-02-03">2016-02-03</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>
