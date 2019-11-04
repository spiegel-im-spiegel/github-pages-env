+++
title = "Google Books の Library Book Scan すら Fair Use と言われたのに..."
date = "2015-11-01T21:37:18+09:00"
description = "今回の判決について，大原ケイさんの「グーグル勝訴で浮き彫りになるフェア・ユースと著作権の問題」を参照しながら簡単に紹介する。"
tags = ["code", "law", "intellectual-property", "copyright", "fair-use", "google", "book", "search", "library"]

[scripts]
  mathjax = false
  mermaidjs = false
+++

## 事の発端は...

Google Books の [Library Project] において米国内の図書館（ハーバード大学図書館など）の蔵書を著作権の有無にかかわらずデジタル化し，検索できるようにしようとしたことである。

当然のように米国内の作家ギルド（The Authors Guild）や全米出版社協会は反発し，集団訴訟を起こした。
この和解案が示されたのが2009年2月だが，その内容に世界中が騒然となった。
それは Google にとって大幅に有利な内容であるとともに米国内の作家のみならず世界中の作家に影響をおよぼすものであったからだ。
もちろんそれには日本の作家も含まれる。

この辺の話は名和小太郎さんの「[グーグル・ブック・サーチ，あるいはバベルの図書館 新しいぶどう酒は新しい革袋に](https://www.jstage.jst.go.jp/article/johokanri/53/3/53_3_131/_article/-char/ja/)」で詳しく解説されている。
つまり Google が提示する著作権システムは「新しい革袋」を要求しているということだ。
「古い革袋」と「新しい革袋」を対比した表があるので以下に引用してみる[^a]。

[^a]: 余談だがこの表は [TablesGenerator.com](http://www.tablesgenerator.com/) で作図した。ちょー便利！

{{< fig-gen title="古い革袋 対 新しい革袋" >}}
<table>
  <tr>
    <th colspan="2"></th>
    <th>古い革袋</th>
    <th>新しい革袋</th>
  </tr>
  <tr>
    <td rowspan="3">原則</td>
    <td>人格的な権利</td>
    <td>不可欠</td>
    <td>有効利用が前提<br></td>
  </tr>
  <tr>
    <td>経済的な権利</td>
    <td>許諾権</td>
    <td>報酬請求権</td>
  </tr>
  <tr>
    <td>権利の取得</td>
    <td>無条件</td>
    <td>登録が条件</td>
  </tr>
  <tr>
    <td rowspan="3">その他</td>
    <td rowspan="2">保護の対象</td>
    <td>エリートの閃き</td>
    <td>投資も<br></td>
  </tr>
  <tr>
    <td>原著作物が主</td>
    <td>二次的著作物が重要<br></td>
  </tr>
  <tr>
    <td>対価<br></td>
    <td>ユーザ負担</td>
    <td>第三者負担も，無償も</td>
  </tr>
</table>
{{< /fig-gen >}}

特に「経済的な権利」について「新しい革袋」では「報酬請求権」となっているが，これは要するに著作物の利用をオプトアウトで行いましょうというもので，これまでの著作権システムと真っ向から対立する。
それだけ [Library Project] はインパクトのある内容だったのである。

しかし2011年にニューヨーク地方裁判所がこの和解の承認を拒否。
訴訟は振り出しに戻ってしまった。

{{< fig-quote title="グーグル勝訴で浮き彫りになるフェア・ユースと著作権の問題" link="http://magazine-k.jp/2015/10/29/google-books-wins-over-the-authors-guild/" >}}
<q>そして2013年にようやく発表された判決は、著者協会側の言い分を全面的に却下、本をスキャンしてデジタル化するのはフェア・ユースの範囲内なので、グーグルは無罪、というもの。
電子書籍黎明期にあって、原告が主張するように、心配されていたことも多々あったけれど、とりあえずこれだけの時間が経ってようやく、デジタル化された本のデータが実際にどう使われるかが見えてきたところで、グーグルのプロジェクトは「フェア・ユース」の範囲内であり、著作権侵害に当たらないという判断が下せるようになった、ということでしょう。</q>
{{< /fig-quote >}}

要するに，和解案で言われていたオプトアウト云々以前に（Google が主張した通り） fair use でええやん，ということになった。
これはある意味すごいことで，和解案で天地がひっくり返りそうになっていたところを fair use として認めることで最小限のインパクトに抑えているのである。

駄菓子菓子，作家ギルドはこれを不服とし第2巡回区連邦控訴裁判所（日本でいうところの高等裁判所）に上訴した。
その結果が今回の判決である。
だぁ，前置きが長すぎた。

## 本当に Library Project は Fair Use か

以降は大原ケイさんの「[グーグル勝訴で浮き彫りになるフェア・ユースと著作権の問題](http://magazine-k.jp/2015/10/29/google-books-wins-over-the-authors-guild/)」を参照しながら簡単に紹介する。

米国で fair use が成立するためには以下の観点において「公正」であることが認められなければならない。

1. 利用の目的や本質
2. 原作品の本質
3. 抜粋の量や実質性
4. 原作品の価値への影響

**利用の目的や本質**： [Library Project] の目的は本の内容をまるっと見せることではなく「本についての情報」を検索可能にすることである。したがって著作権の切れていない著作物については全文を見れるようにはなっていない。
紙の本で出来なかった「書籍を横断的に検索する」という機能が transformative であるということらしい。

**原作品の本質**： 集団訴訟の代表者の作品はいずれもノンフィクションだったけど，「事実を書いてるだけだから著作権はない」なんてことはなく，原作品にもちゃんと著作権はあるという確認。

**抜粋の量や実質性**： 先ほど書いたように[Library Project] では本の全文を見ることは出来ない。
でも検索を可能にするためには全文スキャンしなきゃいけないわけで，それを以って違法とはいえないだろうという判断。
ちなみに

{{< fig-quote title="グーグル勝訴で浮き彫りになるフェア・ユースと著作権の問題" link="http://magazine-k.jp/2015/10/29/google-books-wins-over-the-authors-guild/" >}}
<q>短すぎて、数行表示されただけで用が済んでしまうコンテンツ、例えば料理のレシピとか、辞書とか、詩歌や俳句などはブックスキャンのsnippetサービスから除外されています。</q>
{{< /fig-quote >}}

なんだそうだ。

**原作品の価値への影響**： 今回の件では，ここが一番重要だろう。

{{< fig-quote title="グーグル勝訴で浮き彫りになるフェア・ユースと著作権の問題" link="http://magazine-k.jp/2015/10/29/google-books-wins-over-the-authors-guild/" >}}
<q>つまり、グーグルで検索できることで、調べ物をするのにこの本は要らないや、という判断になることもあり、その分、著者の儲けが減るという可能性もあるだろうけど、全体的に見れば、一部閲覧という形でその本についての情報が得られれば、その情報に基いて本を買う、という判断もあるわけで、原告が主張するように、グーグルで見られるから買わなくなるとばかりは言えないよね、という判断。</q>
{{< /fig-quote >}}

というわけで， Google の全面勝訴とあいなった。
作家ギルドは最高裁までいこうとしているようだが，まぁもう決まりだろう。

記事中で大原ケイさんが書かれている

{{< fig-quote title="グーグル勝訴で浮き彫りになるフェア・ユースと著作権の問題" link="http://magazine-k.jp/2015/10/29/google-books-wins-over-the-authors-guild/" >}}
<q>今回の控訴審の判決文を読んで、なるほどなぁ、と思わされたのが「なぜ著作権というものがあって、それを法律で保護するのか」を憲法に基いて再確認しているところ。
もちろん、一義的には、「何かしらクリエイティブなものを生み出した人（＝本を書いた人）が、それを広めるときにそれなりの見返りが得られるようにして、そのクリエイティブな活動を奨励するため」なのですが、広義的・根本的には「すべての人々が知の恩恵を受けられるように、何かを生み出した当人の著作権を認めてその知を広める」ということで、あくまでも「公益」を守るためなんだなぁということがわかります。</q>
{{< /fig-quote >}}

という部分はその通りで，表現やアイデアは基本的に「公衆のもの」であるという点は外すことなく，その上で表現やアイデアを生み出す人たちも幸せになれるにはどうしたらいいか，を考えていかなければならない。
だから表現やアイデアを「私有化」するという方向は基本的にありえないのである。

ちなみに日本ではこのような判決にはならないだろう。
あらゆる本のスキャンとその検索サービスの提供は国立国会図書館にだけ与えられた特権であり，他の法人・個人には許容されない。
許容するためには著作権法を修正しなければならない。
Fair use のない日本が如何に時代に取り残されつつあるか分かると思う。

## 一方，その頃日本では

まだ「図書館のせいで本が売れない」とか駄々をこねてるようで。

- [「本が売れぬのは図書館のせい」というニュースを見たのでデータを確かめてみました - CNET Japan](http://japan.cnet.com/sp/t_hayashi/35072745/)

ちなみに[政策研究大学院大学の修士論文](http://www3.grips.ac.jp/~ip/paper.html)に「{{< pdf-file title="公立図書館における書籍の貸出が売上に与える影響について" link="http://www3.grips.ac.jp/~ip/pdf/paper2011/MJI11004nakase.pdf" >}}」という論文があるが，分析結果の解釈として

{{< fig-quote title="公立図書館における書籍の貸出が売上に与える影響について" link="http://www3.grips.ac.jp/~ip/pdf/paper2011/MJI11004nakase.pdf" >}}
<q>分析結果から、図書館による書籍の貸出は、売上に対して、正の影響を与えていることが実証された。
この結果を踏まえれば、貸出を減少させるような行為は、却って売上を減少させることになるため、著作者にとっても不本意な結果をもたらすことになるだろう。</q>
{{< /fig-quote >}}

と結ばれている。
まぁ，この論文には保留されている問題もあるようなので，鵜呑みには出来ないけど。

本をよく買う人ってのは，自分で買う以上に本を読む人なんだよね。
もしくは蒐集家か。
私なんか子どもの頃は一晩で文庫本8冊くらいなら読める人だったけど（もっと読む人もいた），子どもがそんなに本買えるわけないじゃん。
だからほとんど毎日学校の図書館に通ってたし，ほぼ毎週県立図書館にも通っていた[^c]。
本を読む習慣を身につけられるのは子どものうちなんだよ[^b]。

[^b]: そういや，最近の子どもは漫画の読み方を知らないと聞く。まぁ読んだことないものは読めないよね。
[^c]: そうして大人になって自分で本や CD が買えるようになると部屋が魔窟になるわけやね（笑）

本を読む習慣のない人は本を買わないし，音楽を聴く習慣のない人は音楽を買わないのだ。
そんな当たり前のことが何故わからない。

## 参考文献

### 参考になるページ

- [ニュース - Google Booksめぐる集団訴訟、連邦地裁が修正和解案を認めず：ITpro](http://itpro.nikkeibp.co.jp/article/NEWS/20110323/358605/)
- [ニュース - Google、「Google Books」を巡る訴訟でフェアユースを主張---米英メディアの報道：ITpro](http://itpro.nikkeibp.co.jp/article/NEWS/20130924/506405/)
- [ニュース - 「Google Books」を巡る係争でGoogleが勝利、フェアユースの主張が認められる：ITpro](http://itpro.nikkeibp.co.jp/article/NEWS/20131115/518222/)
- [ニュース - 「Google Booksは著作権法に違反せず」米控訴裁、地裁判断を支持：ITpro](http://itpro.nikkeibp.co.jp/atcl/news/15/101903421/)
- [グーグル勝訴で浮き彫りになる「フェア・ユース」と著作権（とたぶんTPP）の問題｜りんがる｜note](https://note.mu/lingualina/n/n8e6589a8262b)
- [グーグル勝訴で浮き彫りになるフェア・ユースと著作権の問題 « マガジン航](http://magazine-k.jp/2015/10/29/google-books-wins-over-the-authors-guild/)
- [米最高裁、Googleブックスの書籍スキャンを公正使用と認定 | TechCrunch Japan](http://jp.techcrunch.com/2016/04/19/20160418supreme-court-affirms-google-books-scans-of-copyrighted-works-are-fair-use/) : どうやら確定したようです

### 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4757102852?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41YkbcP5IyL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4757102852?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">著作権２．０ ウェブ時代の文化発展をめざして (NTT出版ライブラリー―レゾナント)</a></dt>
    <dd>名和 小太郎 (著)</dd>
    <dd>NTT出版 2010-06-24</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4757102852 (ASIN), 9784757102859 (EAN), 4757102852 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">名著です。今すぐ買うべきです。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-11-13">2018-11-13</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>

[Library Project]: https://www.google.com/googlebooks/library/ "Google Books Library Project – Google Books"
