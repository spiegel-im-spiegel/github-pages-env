+++
date = "2017-01-24T12:59:29+09:00"
title = "別に貴方のことが嫌いなわけではない"
draft = true
tags = ["communication", "engineering", "management"]
description = "プロジェクト全体を見回して，いったい何がボトルネックになっているのか正しく把握すること。"

[author]
  url = "https://baldanders.info/spiegel/profile/"
  linkedin = "spiegelimspiegel"
  flickr = "spiegel"
  instagram = "spiegel_2007"
  github = "spiegel-im-spiegel"
  avatar = "/images/avatar.jpg"
  tumblr = ""
  license = "by-sa"
  flattr = ""
  name = "Spiegel"
  facebook = "spiegel.im.spiegel"
  twitter = "spiegel_2007"
+++

（ツンデレおやじがデレているわけではない`w`）

- [なぜプログラマはあなたの事が嫌いなのか - megamouthの葬列](http://megamouth.hateblo.jp/entry/2017/01/19/053801)

<p>別に貴方のことは嫌いじゃない。<br>
強いて言うなら「貴方は背中を預けるに足る人ですか」ってこと。<br>
誰だって身内に切り殺されたくはないからね。<br>
貴方もそうでしょう？</p>

もちろん顧客には顧客の立場があり，営業やチーム・マネージャにもそれぞれ立場がある。
そしてもちろんエンジニアにも立場がある。
みんなそれぞれの責務を背負っているのだから。

顧客も営業もマネージャもエンジニアも，みんな対等で peer じゃないと互いの信頼は醸成されないし仕事も完遂できない[^job]。
たとえば，顧客に対して嘘の進捗を報告する営業やマネージャの言うことを現場のエンジニアが聞くと思うかい？
そんな人の言うことを聞いたところでエンジニアにも嘘をつかれた挙句に都合が悪くなれば責任をなすりつけられて排除されるのがオチだ。
プロジェクト全体を見回して，いったい（誰がではなく）何がボトルネックになっているのか正しく把握すること。

[^job]: まっ，無理くりやれば出来ないことはないかもしれないが，それは間違いなく「人を壊すプロジェクト」になる。

しかし，実際問題として「みんな対等で peer」な環境を作るのは簡単ではない。
これに関して面白い記事があった。

- [組織全体でGitHubを使うようになるまで / Everyone Can Use GitHub Issues // Speaker Deck](https://speakerdeck.com/osa/everyone-can-use-github-issues)
- [組織全体でGitHubを使うようになるまで（前編）～ 使い方が分からない？ 使うのが怖い？ Cookpad TechConf 2017 － Publickey](http://www.publickey1.jp/blog/17/github_cookpad_techconf_2017.html)
- [組織全体でGitHubを使うようになるまで（後編）～ エンジニアでなくてもGitHubは便利に使える。Cookpad TechConf 2017 － Publickey](http://www.publickey1.jp/blog/17/github_githubcookpad_techconf_2017.html)

この事例は企業内の内省プロジェクトなので請負開発や業務委託とはプレイヤーが異なるのだが基本は同じ。
で，この事例のポイントは「GitHub を使う」ことではないのだ[^gh1]。
少し引用してみよう。

[^gh1]: 記事の中でも GitHub を選択した理由として「身近にあったから」「エンジニアはGitHubのIssueでコミュニケーションしちゃうので、そこで分断して見えなくなるところができるのがイヤだった」と述べている。つまり GitHub 自体が魔法のツールというわけではないのだ。

{{< fig-quote title="組織全体でGitHubを使うようになるまで（前編）" link="http://www.publickey1.jp/blog/17/github_cookpad_techconf_2017.html" >}}
<q>絵文字を使うのはエンジニアには普通だと思うのですが、非エンジニアでは「仕事で絵文字使っていいんですか？」という人がけっこういて、文化というか意識が違ったりするので、いいんだよと。ハードルを下げることをしてきました。</q>
{{< /fig-quote >}}

{{< fig-quote title="組織全体でGitHubを使うようになるまで（後編）" link="http://www.publickey1.jp/blog/17/github_githubcookpad_techconf_2017.html" >}}
<q>GitHubは最初はエンジニアがコミュニケーションをする、そいういうツールでした。<br>
けれども、チーム全員がGitHubでコミュニケーションをしようと決めたのだったら、そこはエンジニアの場じゃなくてみんなの場なんですね。だから自分たちも振る舞いを変える必要があります。</q>
{{< /fig-quote >}}

{{< fig-quote title="組織全体でGitHubを使うようになるまで（後編）" link="http://www.publickey1.jp/blog/17/github_githubcookpad_techconf_2017.html" >}}
<q>サポートエンジニアから「こういう対応をお願いします」とIssueが投げられると、投げた人はそのあとの過程も気にして見るんですね。見ると詳細は分らなくても、こうしたことをやってるんだなあと、だったら自分たちでこういうこともできるんじゃないかなあと、隣のチームへの関心が沸いてくるという、思いがけない効果がありました。</q>
{{< /fig-quote >}}

たとえば顧客は顧客の「言葉」を使って話す。
工場のおじさんやコールセンターのお姉さんや大学の先生や... それぞれみんな異なる言葉で話すのだから，それをきちんと理解し更に相手に分かるように提案するにはコミュニケーションを繰り返すしかない。
そのためにはお互いの立ち位置からそれぞれ一歩ずつ「降りる」ことが必要になる。
それが上手く回ったのが「[組織全体でGitHubを使うようになるまで](https://speakerdeck.com/osa/everyone-can-use-github-issues)」なんだと思う。

「なぜプログラマはあなたの事が嫌いなのか」なんて問いの立て方はそれ自体がコミュニケーションを拒絶しているも同然であり，そこから何らかの解が得られる見込みはないと言っていいだろう。
仕事なんだから相手におもねる必要はないんだけど，コミュニケーションを拒絶している相手に話しかける人などいない。

道具はなんだっていい。
もちろん GitHub でもいいし [Redmine](http://www.redmine.org/) みたいなチケット駆動の ITS (Issue Tracking System) でもいいかもしれない（ある職場ではプロジェクトごとに ITS を顧客にまで開放して運用していた。これなら営業が嘘をつく余地はない`w`）。
いっそのこと [Slack](https://slack.com/) で「これからは [ChatOps の時代](http://www.nttdata.com/jp/ja/insights/trend_keyword/2016070701.html "ChatOpsで加速するOps効率化 | NTTデータ")だZ！」とかやったって構わない。

目的はあくまで「みんな対等で peer」な環境を作ることであり，その手段としてツールやルールを整備するのである。

<div class="hreview" ><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4822282686/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/512Y77Y5WDL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4822282686/baldandersinf-22/">要求開発~価値ある要求を導き出すプロセスとモデリング</a></dt><dd>山岸 耕二 安井 昌男 萩本 順三 河野 正幸 野田 伊佐夫 平鍋 健児 細川 努 依田 智夫 ［要求開発アライアンス］ </dd><dd>日経BP社 2006-03-02</dd></dl><p class="similar"><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4822283585/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4822283585.09._SCTHUMBZZZ_.jpg"  alt="UMLモデリング入門"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4798121967/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4798121967.09._SCTHUMBZZZ_.jpg"  alt="エリック・エヴァンスのドメイン駆動設計 (IT Architects’Archive ソフトウェア開発の実践)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4492961143/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4492961143.09._SCTHUMBZZZ_.jpg"  alt="ビジネスプロセスの教科書"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4822283496/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4822283496.09._SCTHUMBZZZ_.jpg"  alt="UMLモデリングレッスン"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4274505219/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4274505219.09._SCTHUMBZZZ_.jpg"  alt="ソフトウェアエンジニアリング基礎知識体系 ―SWEBOK V3.0―"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4320023528/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4320023528.09._SCTHUMBZZZ_.jpg"  alt="要求仕様の探検学―設計に先立つ品質の作り込み"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4764904047/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4764904047.09._SCTHUMBZZZ_.jpg"  alt="要求工学知識体系 第1版"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4798114456/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4798114456.09._SCTHUMBZZZ_.jpg"  alt="ユースケース駆動開発実践ガイド (OOP Foundations)"  /></a> </p>
<p class="description">要求は開発するものらしい。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-01-24">2017-01-24</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4621045938/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/51XGP8AFX2L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4621045938/baldandersinf-22/">いかにして問題をとくか</a></dt><dd>G. ポリア 柿内 賢信 </dd><dd>丸善 1975-04-01</dd><dd>評価<abbr class="rating" title="4"><img src="https://images-fe.ssl-images-amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4621085298/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4621085298.09._SCTHUMBZZZ_.jpg"  alt="いかにして問題をとくか・実践活用編"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4061497863/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4061497863.09._SCTHUMBZZZ_.jpg"  alt="数学的思考法―説明力を鍛えるヒント  講談社現代新書"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/462108819X/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/462108819X.09._SCTHUMBZZZ_.jpg"  alt="数学×思考=ざっくりと  いかにして問題をとくか"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4797375698/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4797375698.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4185086180/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4185086180.09._SCTHUMBZZZ_.jpg"  alt="授業研究に学ぶ高校新数学科の在り方"  /></a> </p>
<p class="description" >数学書。というか問いの立てかたやものの考え方についての指南書。のようなものかな。</p>
<p class="gtools" >reviewed by <a href="#maker" class="reviewer">Spiegel</a> on <abbr class="dtreviewed" title="2014-09-26">2014/09/26</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html">G-Tools</a>)</p>
</div>
