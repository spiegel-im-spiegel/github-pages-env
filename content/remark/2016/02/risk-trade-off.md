+++
date = "2016-02-03T22:02:21+09:00"
description = "リスクと利便性はトレードオフできない。リスクとトレードオフできるのはリスクのみ。リスクをトレードオフして「全体最適化」することが重要なのである。"
draft = false
tags = ["security", "risk", "management"]
title = "リスク認知とトレードオフ"

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

これとよく似た表があったな，と思ったら吉川肇子さんの『[リスクとつきあう](http://www.amazon.co.jp/exec/obidos/ASIN/4641280304/baldandersinf-22/)』という本の中にあった。
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
営利企業は個人と似ているが，賢い企業はリスクをお金に換算できる（これは個人も見習うべきだと思うが）。
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

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4822283100/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51-pZ52JsUL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4822283100/baldandersinf-22/">セキュリティはなぜやぶられたのか</a></dt><dd>ブルース・シュナイアー 井口 耕二 </dd><dd>日経BP社 2007-02-15</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4881359967/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4881359967.09._SCTHUMBZZZ_.jpg"  alt="暗号の秘密とウソ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797350997/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797350997.09._SCTHUMBZZZ_.jpg"  alt="新版暗号技術入門 秘密の国のアリス"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4594070507/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4594070507.09._SCTHUMBZZZ_.jpg"  alt="チャイナ・ハッカーズ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4105393022/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4105393022.09._SCTHUMBZZZ_.jpg"  alt="暗号解読―ロゼッタストーンから量子暗号まで"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4102159746/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4102159746.09._SCTHUMBZZZ_.jpg"  alt="宇宙創成〈上〉 (新潮文庫)"  /></a> </p>
<p class="description" >日本語のタイトルはアレだが中身は名著。とりあえず読んどきなはれ。</p>
<p class="gtools" >reviewed by <a href="#maker" class="reviewer">Spiegel</a> on <abbr class="dtreviewed" title="2014-09-14">2014/09/14</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html">G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B00E7HMI7U/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51I9C7cFl2L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B00E7HMI7U/baldandersinf-22/">環境リスク学</a></dt><dd>中西準子 </dd><dd>日本評論社 2013-08-01</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00E7HMIB6/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00E7HMIB6.09._SCTHUMBZZZ_.jpg"  alt="食のリスク学"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00M98XGDO/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00M98XGDO.09._SCTHUMBZZZ_.jpg"  alt="基準値のからくり　安全はこうして数字になった (ブルーバックス)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00VQ75FAQ/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00VQ75FAQ.09._SCTHUMBZZZ_.jpg"  alt="21世紀の資本"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B014II6012/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B014II6012.09._SCTHUMBZZZ_.jpg"  alt="丸山眞男セレクション (平凡社ライブラリー700)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00BWI0U0O/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00BWI0U0O.09._SCTHUMBZZZ_.jpg"  alt="「ゼロリスク社会」の罠～「怖い」が判断を狂わせる～"  /></a> </p>
<p class="description">環境リスクのみならず「リスク」全体に目配せした良書。著者の自伝的作品でもある。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-02-03">2016-02-03</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4641280304/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/519S1SM2S4L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4641280304/baldandersinf-22/">リスクとつきあう―危険な時代のコミュニケーション (有斐閣選書)</a></dt><dd>吉川 肇子 </dd><dd>有斐閣 2000-03</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4885554241/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4885554241.09._SCTHUMBZZZ_.jpg"  alt="リスクコミュニケーション (エネルギーフォーラム新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4779502357/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4779502357.09._SCTHUMBZZZ_.jpg"  alt="健康リスク・コミュニケーションの手引き"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873262526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873262526.09._SCTHUMBZZZ_.jpg"  alt="リスクコミュニケーション―前進への提言"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4480066845/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4480066845.09._SCTHUMBZZZ_.jpg"  alt="「リスク」の食べ方―食の安全・安心を考える (ちくま新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4872592840/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4872592840.09._SCTHUMBZZZ_.jpg"  alt="リスクコミュニケーション論 (シリーズ環境リスクマネジメント)"  /></a> </p>
<p class="description">リスク・コミュニケーションについて。内容は古いがまだまだ使える。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-02-03">2016-02-03</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

