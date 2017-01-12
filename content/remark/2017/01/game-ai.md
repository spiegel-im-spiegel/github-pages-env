+++
tags = ["games", "artificial-intelligence"]
description = "これは2016年に Medium に書いた記事の再掲載です。"
date = "2017-01-12T15:47:01+09:00"
title = "ゲーム AI の本領"
draft = true

[author]
  flickr = "spiegel"
  twitter = "spiegel_2007"
  avatar = "/images/avatar.jpg"
  license = "by-sa"
  url = "http://www.baldanders.info/spiegel/profile/"
  linkedin = "spiegelimspiegel"
  flattr = "spiegel"
  tumblr = "spiegel-im-spiegel"
  github = "spiegel-im-spiegel"
  name = "Spiegel"
  facebook = "spiegel.im.spiegel"
  instagram = "spiegel_2007"
+++

（これは[2016年に Medium に書いた記事](https://medium.com/@spiegel/-67c0fc849272 "ゲーム AI の本領 – Medium")の再掲載です。
更に大元をたどると，当時 Facebook の TL に書き散らした記事らしい`w` 今回はこの記事に加えて最近の状況を追記しています）

----

## ゲーム AI の本領（[2016年3月14日の記事](https://medium.com/@spiegel/-67c0fc849272 "ゲーム AI の本領 – Medium")を再掲載 ）

- [DeepMindのAIに負けた囲碁の世界チャンピオンが最終戦直前のゲームで勝利…AlphaGoを上回る妙手で | TechCrunch Japan](http://jp.techcrunch.com/2016/03/13/20160313defeated-go-world-champion-beats-deepmind-ai-in-penultimate-match/)

この手のストラテジー・ゲームって膨大な手筋があって機械でも全てを把握することは不可能。
だからコンピュータも人と同じようにその時々の局面で有効な手筋を探索するという戦略を取らざるをえない（もちろん思考の仕方は人とコンピュータでは違うのだが）。

だから機械が充分優秀であれば機械のとる戦略から人が学ぶということも起こりうる。
多分「人 vs 機械」対局の本当の面白さはそこにあるんじゃないだろうか。
ある対局の勝ち負けのみを追いかけてたら，この重要なポイントを見逃してしまう。

日本では将棋 AI が有名だが，日本の将棋 AI は「既知の手筋を模倣し組み合わせる」というレベルに留まっていて（エキスパートシステムとしてはこれで充分優秀だが） DeepMind/AlphaGo のように未知の手筋を探索するというところまで達していない。

対局というのはお互いに闘っているわけだけど，一方でプレイヤー同士の相互作用で戦局を構築していくという傍から見るとまるで協働作業（collaboration）のような側面もある。
そこがストラテジー・ゲームの面白さであり，ゲーム AI の本領はそこにあると思う。
機械が単なる「道具」であることを超えて人と関わってくることで，今まで想像もしなかった「未知の手筋」が見えてくるかもしれない。

巷では「人工知能が人の仕事を奪う」みたいな言説が流布されているけど，機械の本領は人の仕事を肩代わりすることではない。

海外では（日本の一部の法専門家も），「ターミネーター」のトラウマでもあるのか，人工知能に恐怖し排除するような動きがある。
人工知能を「スマートな道具」に押しとどめるならそこから先に進むことはないし，結局は「人工知能が人の仕事を奪う」だけになるだろう。
そこでリスクを取って「人と機械が対等に関わりあう」ことができるかどうかがポイントだ。

ゲームのように「特定のルール下で人と機械を対等に配置する」というのは将来の人工知能開発におけるひとつの方向性を指してるように思える。

- [「囲碁の謎」を解いたグーグルの超知能は、人工知能の進化を10年早めた｜WIRED.jp](http://wired.jp/2016/01/31/huge-breakthrough-google-ai/)
- [「またこれから学ぶことが増えました」AlphaGoとイ・セドルが、囲碁にもたらしたもの、AIにもたらしたもの｜WIRED.jp](http://wired.jp/2016/03/16/final-round/)

## 「人 vs 機械」という思考停止

- [ネットの謎棋士60連勝、熱狂生んだ陰の主役：日経ビジネスオンライン](http://business.nikkeibp.co.jp/atcl/opinion/15/221102/010500382/?rt=nocnt)

へぇ。
日経ビジネスで（細かい点はいろいろあるだろうが[^a]）このレベルの記事が書けることのほうが驚きだよ。
個人的には昨年のプロ将棋界の不祥事（だよね？）には鼻くそほどの興味もないけど[^b]，この記事は面白い。

[^a]: 最初「アルファ碁」が何かわからなかった。 AlphaGo のことか！ カタカナで書くな。分かりにくいわ！
[^b]: まぁ「将棋指しは頭が固い」という印象を植え付けることには貢献しているかもしれないが。

AlphaGo が他の “AI” （と敢えて言うけど）と一線を画すのは機械自ら（対局を通して）ゲーム空間を探索できることだ。
これは今までの「エキスパート・システム」の延長線上の存在にすぎなかった “AI” が，一歩先に進んだことを意味する。
だから人間が AlphaGo と対戦するのは「楽しい」のである。

- [囲碁のトップ棋士を連破した謎のアカウントの正体は新版「AlphaGo」～ハサビス氏が明かす - 窓の杜](http://forest.watch.impress.co.jp/docs/news/1037627.html)
- [“謎の囲碁棋士”「Master」、正体はGoogleのAI「AlphaGO」改良版　「年内にも公式戦を」 - ITmedia ニュース](http://www.itmedia.co.jp/news/articles/1701/05/news060.html)
- [謎の囲碁棋士「Master」は「AlphaGo」と判明--トップ棋士ら相手に60連勝 - CNET Japan](http://japan.cnet.com/news/service/35094593/)

## ブックマーク

- [「AIを搭載」は「全て自然」同様の技術的ナンセンスだ | TechCrunch Japan](http://jp.techcrunch.com/2017/01/11/20170110ai-powered-is-techs-meaningless-equivalent-of-all-natural/)

- [人工知能は「ハイル・ヒトラー」と叫ぶか]({{< relref "remark/2015/artificial-intelligence.md" >}})
- [「ピノキオ」と AI]({{< relref "remark/2016/08/pinocchio.md" >}})
- [AI は人（の良心）を殺すか？]({{< relref "remark/2016/10/artificial-intelligence.md" >}})
- [AI の読解力，人の読解力]({{< relref "remark/2016/11/reading-comprehension.md" >}})
