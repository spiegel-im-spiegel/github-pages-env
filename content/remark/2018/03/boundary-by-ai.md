+++
title = "機械によって引かれた境界線"
date = "2018-03-18T15:34:11+09:00"
description = "あぁ，久しぶりにとりとめのない戯れ言を書いてる気がする。"
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "code", "grigori" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

あぁ，久しぶりにとりとめのない戯れ言を書いてる気がする。

## 機械を使う人と使われる人の境界線

最近読んだ記事で面白いものがあった。
有料記事で会員[^nid1] 以外はもう読めないのだが一応リンクだけ挙げておく。

[^nid1]: [日経IDでは以前に酷い目にあった](http://www.baldanders.info/spiegel/log2/000709.shtml "NIKKEI is Worst of Worsts. もしくは「無料（ただ）より高くつくものはない」 — Baldanders.info")ので私は絶対に会員にはならない。

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

[^ait1]: 「AI」ではなく「AI 技術」と呼ぶ理由については，以前に書いた「[『AI vs. 教科書が読めない子どもたち』を読む]({{< relref "remark/2018/02/artificial-intelligence-book.md" >}})」を参照のこと。

でも，本当はそれは間違いだと思う。
「機械 vs 人間」という構図は「機械社会」の中では意味を成すかも知れないが，実際に私達が暮らすのは「人間社会」である。

機械が賢くなって，それまで人間がやっていたことの多くを肩代わりできるようになる（なった）かも知れないが，それでも機械を使うのは人間だし，機械に使われるのも人間なのである。
機械が自ら考え創造する未来など（少なくとも現代の延長線上においては）やってこない。
機械は情報という名の富を集約・増幅・移転するだけで生み出すわけではないのだ。

であるなら，人間と人間の関係が（機械が差し込まれることによって）どう変化していくかという観点で考えなければならない。
機械によって引かれた（越えられない）境界線で明確に分離される新たな社会階級システムが誕生しつつある。

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B0791XCYQG/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/51KFIJ%2BqpkL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B0791XCYQG/baldandersinf-22/">ＡＩ　ｖｓ．　教科書が読めない子どもたち</a></dt><dd>新井 紀子 </dd><dd>東洋経済新報社 2018-02-02</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0797K44CH/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0797K44CH.09._SCTHUMBZZZ_.jpg"  alt="日本再興戦略 (NewsPicks Book)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0791CLWH8/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0791CLWH8.09._SCTHUMBZZZ_.jpg"  alt="英語教育の危機 (ちくま新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07919S1LQ/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B07919S1LQ.09._SCTHUMBZZZ_.jpg"  alt="総合教育技術 2018年 2月号 [雑誌] 教育技術シリーズ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B079JRSVVQ/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B079JRSVVQ.09._SCTHUMBZZZ_.jpg"  alt="ファンベース　──支持され、愛され、長く売れ続けるために (ちくま新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B072Z81MHK/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B072Z81MHK.09._SCTHUMBZZZ_.jpg"  alt="働きたくないイタチと言葉がわかるロボット 人工知能から考える「人と言葉」"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0798BNCYG/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0798BNCYG.09._SCTHUMBZZZ_.jpg"  alt="プラットフォーム革命――経済を支配するビジネスモデルはどう機能し、どう作られるのか"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B078YLH4W2/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B078YLH4W2.09._SCTHUMBZZZ_.jpg"  alt="働く大人のための「学び」の教科書"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0798QLBVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0798QLBVC.09._SCTHUMBZZZ_.jpg"  alt="アルテミス 下 (ハヤカワ文庫SF)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07919W1YX/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B07919W1YX.09._SCTHUMBZZZ_.jpg"  alt="Rで楽しむベイズ統計入門［しくみから理解するベイズ推定の基礎］ Data Science Library"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0798S7N12/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0798S7N12.09._SCTHUMBZZZ_.jpg"  alt="最後にして最初のアイドル (ハヤカワ文庫JA)"  /></a> </p>
<p class="description">4章以外は面白かった。感想文は<a href="/remark/2018/02/artificial-intelligence-book/">こちら</a>。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-02-11">2018-02-11</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B078TNC8RW/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/51ijGO2LR3L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B078TNC8RW/baldandersinf-22/">新・日本の階級社会 (講談社現代新書)</a></dt><dd>橋本健二 </dd><dd>講談社 2018-01-16</dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B079T1WBP6/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B079T1WBP6.09._SCTHUMBZZZ_.jpg"  alt="週刊文春　3月22日号[雑誌]"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07B9Q4NXQ/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B07B9Q4NXQ.09._SCTHUMBZZZ_.jpg"  alt="文藝春秋2018年4月号[雑誌]"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B078TN4G25/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B078TN4G25.09._SCTHUMBZZZ_.jpg"  alt="９．１１後の現代史 (講談社現代新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B078TJZWNC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B078TJZWNC.09._SCTHUMBZZZ_.jpg"  alt="自民党秘史　過ぎ去りし政治家の面影 (講談社現代新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07653L7LM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B07653L7LM.09._SCTHUMBZZZ_.jpg"  alt="底辺への競争　格差放置社会ニッポンの末路 (朝日新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0792VXR45/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0792VXR45.09._SCTHUMBZZZ_.jpg"  alt="1985年の無条件降伏～プラザ合意とバブル～ (光文社新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B077D22GFX/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B077D22GFX.09._SCTHUMBZZZ_.jpg"  alt="戦争調査会　幻の政府文書を読み解く (講談社現代新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0776RFRZ4/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0776RFRZ4.09._SCTHUMBZZZ_.jpg"  alt="健康格差　あなたの寿命は社会が決める (講談社現代新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0791XCYQG/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0791XCYQG.09._SCTHUMBZZZ_.jpg"  alt="AI vs. 教科書が読めない子どもたち"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07B454T89/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B07B454T89.09._SCTHUMBZZZ_.jpg"  alt="近代日本一五〇年－科学技術総力戦体制の破綻 (岩波新書)"  /></a> </p>
<p class="description">Facebook で紹介されているのを見て買ってみたのだが，微妙に読みにくくて積ん読中。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-03-18">2018-03-18</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
