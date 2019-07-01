+++
title = "機械によって引かれた境界線"
date = "2018-03-18T15:34:11+09:00"
description = "あぁ，久しぶりにとりとめのない戯れ言を書いてる気がする。"
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "code", "grigori" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

あぁ，久しぶりにとりとめのない戯れ言を書いてる気がする。

## 機械を使う人と使われる人の境界線

最近読んだ記事で面白いものがあった。
有料記事で会員[^nid1] 以外はもう読めないのだが一応リンクだけ挙げておく。

[^nid1]: [日経IDでは以前に酷い目にあった](https://baldanders.info/blog/000709/ "NIKKEI is Worst of Worsts. もしくは「無料（ただ）より高くつくものはない」 — Baldanders.info")ので私は絶対に会員にはならない。

- [トヨタやJapanTaxiなどAI活用するタクシー配車を検証、ドライバーの売上高2割増の効果 | 日経 xTECH（クロステック）](http://tech.nikkeibp.co.jp/atcl/nxt/news/18/00408/)

これを考案した人は Uber のようなカー・シェアリングへの対抗と思ったのかも知れないが，考えてみれば酷いシステムである。

知り合いにタクシーの運ちゃんがいるのだが，「できるタクシー・ドライバー」ってのはだいたい常連の顧客を持っているし街や郊外における時間帯ごとの人の流れのようなものはおおよそ把握しているらしい。
その上で（売上を計算に入れながら）計画的に仕事をして計画的にサボるわけだ。
知り合いの運ちゃんはタクシーに GPS が配備されて仕事がやりにくくなったとボヤいていた。
ずいぶん昔の話である。

「ドライバーの売上高2割増」というのはそれだけ（ドライバーの思惑とは関係なく）余分に仕事をさせられているということなのだ。
経営者目線なら文字通り「馬車馬のように働け！」ってことかもしれないが，現場で実際に働いている人から見れば有難迷惑な話である。
合理的な最適化が常に正しいとは限らないのだ。

このシステムの特徴は，タクシー・ドライバーを機械が **効率よく** 監視し成果を吸い上げ経営者にアウトプットすることである（**効率よく** の部分が画期的なわけだ）。
この構造はタクシー・ドライバーが機械に依存すればするほど強化され容易に覆せなくなる。

## 自動生成コンテンツの何に価値があるのか

もうひとつ面白い記事。

- [AIが作ったコンテンツの著作権はどうなる？--福井弁護士が解説する知財戦略 - CNET Japan](https://japan.cnet.com/article/35115900/)

「AIが作ったコンテンツの著作権はどうなる」かはリンク先を読んでいただきたい。
著作権法の解釈はともかく，著作権の根底は「出版規制」であり，メディアで流通する情報をその「出版規制」で制限するのはそもそも無理筋であるという点は押さえておくべきだろう。

それはともかく，この記事で面白いと思ったのは以下の部分。

{{< fig-quote title="AIが作ったコンテンツの著作権はどうなる？--福井弁護士が解説する知財戦略" link="https://japan.cnet.com/article/35115900/" >}}
<q>こうやってAIが自動生成した文書は、文脈にゆらぎがないため、自動翻訳との相性も良い。AIが生成した日本語文書を精度良く英語に翻訳できれば、さらにその他言語へ容易に翻訳できる。「つまり、AIで低コストで作った記事は、数十カ国語で簡単に配信できる。</q>
{{< /fig-quote >}}

自動生成と言ってもゼロ知識で生成できるわけではない。
元ネタとなるデータとロジック（＝情報）があり，それを基に文字通り機械的に吐き出しているに過ぎない。
本当に価値があるのは吐き出された大量の「◯◯風の作品」ではなく「元ネタとなるデータとロジック」である。

## 機械は機械とお話したい

Google 翻訳 AI は翻訳のベース言語として（英語ではなく）独自の言語を獲得しているという記事がある。

- [グーグルの翻訳AIが「独自の言語」を生み出したといえる根拠｜WIRED.jp](https://wired.jp/2016/11/24/google-ai-language-create/)

つまり「独自の言語」で構築された情報があれば（翻訳 AI が対応できる）どんな言語でも出力できるわけだ。
そして，もし「独自の言語で構築された情報」を他のサービスと共有できるなら，機械同士でかなり効率よく情報交換できるはずである。

- [人工知能は「機械同士で会話する」独自の言語を覚え始めている｜WIRED.jp](https://wired.jp/2017/03/30/bots-learn-speak-language/)

たとえばメッセージング・サービスを考えてみよう。

電子メールにしろ LINE のようなメッセージング・アプリにしろ，基本的には人間同士で通じる「人間語」でやり取りされる。
当然だよね。
でも人間とメッセージング・サービスの間に AI アシスタントのようなものが差し込まれたらどうなるか。
この場合，直接メッセージをやり取りしているのは機械なので，もはや「人間語」である必要はない。
人間は機械同士の会話を「人間語」に翻訳したものを受け取ることになるだるう。

これなら，相手が異なる言語圏の人でも関係ないし，下らないメールマナーなるものも考えなくて済む。
万々歳（笑）

これを機械側から見ると非合理的な「人間語」の情報は価値が低い。
サービス・プロバイダは機械から吐き出される「人間語」よりも機械同士の会話に用いられる「機械語」に重きを置いてサービスを最適化していくかもしれない。
ここに機械を境界線とする情報の不均衡が発生する。

## 機械によって引かれた境界線

AI 技術[^ait1] について考える時，私たちは「機械 vs 人間」という構図で考えることが多い。
この構図の上で機械が人間に及ぼす影響を肯定的にあるいは否定的に論ずるわけだ。

[^ait1]: 「AI」ではなく「AI 技術」と呼ぶ理由については，以前に書いた「[『AI vs. 教科書が読めない子どもたち』を読む]({{< ref "/remark/2018/02/artificial-intelligence-book.md" >}})」を参照のこと。

でも，本当はそれは間違いだと思う。
「機械 vs 人間」という構図は「機械社会」の中では意味を成すかも知れないが，実際に私達が暮らすのは「人間社会」である。

機械が賢くなって，それまで人間がやっていたことの多くを肩代わりできるようになる（なった）かも知れないが，それでも機械を使うのは人間だし，機械に使われるのも人間なのである。
機械が自ら考え創造する未来など（少なくとも現代の延長線上においては）やってこない。
機械は情報という名の富を集約・増幅・移転するだけで生み出すわけではないのだ。

であるなら，人間と人間の関係が（機械が差し込まれることによって）どう変化していくかという観点で考えなければならない。
機械によって引かれた（越えられない）境界線で明確に分離される新たな社会階級システムが誕生しつつある。

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/AI-vs-%E6%95%99%E7%A7%91%E6%9B%B8%E3%81%8C%E8%AA%AD%E3%82%81%E3%81%AA%E3%81%84%E5%AD%90%E3%81%A9%E3%82%82%E3%81%9F%E3%81%A1-%E6%96%B0%E4%BA%95-%E7%B4%80%E5%AD%90-ebook/dp/B0791XCYQG?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B0791XCYQG"><img src="https://images-fe.ssl-images-amazon.com/images/I/51KFIJ%2BqpkL._SL160_.jpg" width="111" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/AI-vs-%E6%95%99%E7%A7%91%E6%9B%B8%E3%81%8C%E8%AA%AD%E3%82%81%E3%81%AA%E3%81%84%E5%AD%90%E3%81%A9%E3%82%82%E3%81%9F%E3%81%A1-%E6%96%B0%E4%BA%95-%E7%B4%80%E5%AD%90-ebook/dp/B0791XCYQG?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B0791XCYQG">AI vs. 教科書が読めない子どもたち</a></dt>
	<dd>新井 紀子</dd>
    <dd>東洋経済新報社 2018-02-02 (Release 2018-02-02)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B0791XCYQG</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">4章以外は面白かった。感想文は<a href="/remark/2018/02/artificial-intelligence-book/">こちら</a>。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-02-11">2018-02-11</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%96%B0%E3%83%BB%E6%97%A5%E6%9C%AC%E3%81%AE%E9%9A%8E%E7%B4%9A%E7%A4%BE%E4%BC%9A-%E8%AC%9B%E8%AB%87%E7%A4%BE%E7%8F%BE%E4%BB%A3%E6%96%B0%E6%9B%B8-%E6%A9%8B%E6%9C%AC%E5%81%A5%E4%BA%8C-ebook/dp/B078TNC8RW?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B078TNC8RW"><img src="https://images-fe.ssl-images-amazon.com/images/I/51ijGO2LR3L._SL160_.jpg" width="98" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%96%B0%E3%83%BB%E6%97%A5%E6%9C%AC%E3%81%AE%E9%9A%8E%E7%B4%9A%E7%A4%BE%E4%BC%9A-%E8%AC%9B%E8%AB%87%E7%A4%BE%E7%8F%BE%E4%BB%A3%E6%96%B0%E6%9B%B8-%E6%A9%8B%E6%9C%AC%E5%81%A5%E4%BA%8C-ebook/dp/B078TNC8RW?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B078TNC8RW">新・日本の階級社会 (講談社現代新書)</a></dt>
	<dd>橋本健二</dd>
    <dd>講談社 2018-01-16 (Release 2018-01-19)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B078TNC8RW</dd>
    <dd>評価<abbr class="rating fa-sm" title="3">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">Facebook で紹介されているのを見て買ってみたのだが，微妙に読みにくくて積ん読中。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-03-18">2018-03-18</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
