+++
title = "「改憲」という要求開発"
date =  "2017-10-23T17:45:41+09:00"
description = "やぁ，選挙が終わりましたよ。ようやくこれで胡乱なことが喋れるよ（笑）"
tags        = [ "engineering", "politics", "code", "requirement-revelopment" ]

[scripts]
  mathjax = true
  mermaidjs = false
+++

やぁ，選挙が終わりましたよ。
ようやくこれで胡乱なことが喋れるよ（笑） みんなよく選挙期間中に公衆空間で具体的な候補者を名指しでこき下ろしたり礼賛したりできるな。
私は選管や選管に密告するネットユーザが怖いので無理です。

さっそく頭のなかでトグロを巻いてる色々な雑念を [Scrapbox] に吐き出していく（外部記憶に書いて忘れるため）。
そこで半分無意識に書いた

{{< fig-quote title="改憲 - Spiegel's Branch - Scrapbox" link="https://scrapbox.io/spiegel-branch/%E6%94%B9%E6%86%B2" >}} 
<q>普通は逆。国家の主権者のアイデンティティを担保するものとして憲法がある。太平洋戦争で何もかもぶっ壊された当時の日本人には自分たちを肯定するものが天皇と憲法しかなかったのかもしれないが。改憲議論を含めて憲法について考えるなら，自分たちの行動規範や慣習の源泉（ルーツ）について突き詰めていく作業が不可欠だが，突き詰めていくと当の憲法にぶち当たって，そこより以前とは断絶状態になっている</q>
{{< /fig-quote >}} 

を眺めていて（我ながら珍しいことを書くなぁ）ふと思いついた。

「これってコーランのことなんじゃね[^quran1]？」

[^quran1]: いや，イスラーム教徒の方は怒らないでね。この手の「思考の横滑り」は私にはいつものことなので笑って許してください。私はエンジニアだけど，いつも論理的に考えているわけじゃあないんだよ。

なんでこんな明後日の方向に思考が飛んだかというと，以下の記事を思い出したからだ。

- [対話する「未来論」：イスラームの宗教と脳の機能は交差する。](http://30th.rcast.u-tokyo.ac.jp/future/future01.html)

正直な話，最初にこれを読んだときにはうまく脳に馴染まなかった。
ただ，この中で西欧（ギリシア哲学）的な思考とイスラーム教的な思考の違いが分かりやすく例示されている。

{{< fig-quote title="対話する「未来論」：イスラームの宗教と脳の機能は交差する。" link="http://30th.rcast.u-tokyo.ac.jp/future/future01.html" >}} 
<q>例えば高校で数学の問題集をわれわれは渡されて、解答は渡されずに、数学的な真理を見出していく。後で先生が解答集から答えを見せてくれて、それ以前の人間がすでに見出していた真理と照らし合わせて理解する、そういうやり方でわれわれは数学を学びますよね。これはギリシア哲学の時代の、つまり近代的な自我を課題にするようになる以前の、科学としての哲学の方法論です。それに対して、コーランは神が啓示で真理を下した、いわば「解答集」であって、人間はそこから逆算して世界に存在する問題を認識する。問題を認識すると同時に答えも与えられている。そもそも答えが先にあって、答えにあるように世界に存在する問題読み取るのですから、答えがあるのは当たり前なのですが、しかしとにかくそのように解答と問いを同時に受け取って、世界と人間の生命に対する確信を得る。究極のマニュアルなんですね。</q>
{{< /fig-quote >}} 

具体的には算数で「$1+1=\,?$」という問題と「$2$」という解答のセットがあるとする[^math1]。
そこで「$1+1$」について考察し解答である「$2$」を導き出すのが西欧的思考で，「$1+1=2$ である」を絶対として無条件に受け入れるのがイスラーム教的思考と考えればいいのだろうか（算数に関しては日本の学校教育は後者な気もするが[^ed1]）。

[^math1]: 厳密には $1+1=2$ というのは「$1+1$ と $2$ は常に等しい」という意味の恒等式であるが，細かいことは言わないでおく。
[^ed1]: そうでなければ「[掛け算は順序が大事](https://baldanders.info/blog/000744/ "日本の「算数」は壊れてる？ — Baldanders.info")」みたいな大惨事にはならないだろう。

実はコンピュータ・エンジニアも似たようなことをする。
それは TDD (Test-Driven Development) である。

TDD では先に「テスト」を書く。
つまり「問題と解答」のセットを先に作るのである。
そして入力した問題に対して必ず「正しい」解答が出力されるよう手順（algorithm）を記述するのがプログラミングである。

たとえば「20と32の最大公約数は4」を導く手順としては[「ユークリッド互除法」が有名]({{< ref "/golang/greatest-common-divisor.md" >}})だが，なぜ「ユークリッド互除法」で最大公約数が解けるのかエンジニアは考えない[^gcd1]。
テストが要求する $gcd(20,32)=4$ を実装できることが重要なのである。
そもそも大抵の数学ライブラリに入ってるしね（笑）

[^gcd1]: いや，さすがに現場のエンジニアは学生時代に一度くらいは「ユークリッドの互除法」を証明したことがあるだろけど。やったことないって人や忘れてしまった人は[結城浩さんの「数学ガールの秘密ノート」の連載で「ユークリッドの互除法」が登場する](https://cakes.mu/posts/16292)ので読むといいだろう。

現在の日本国憲法が日本人にとって「最後の憲法[^quran2]」であるなら「改憲」議論そのものがナンセンスだろう。
実際そのように考えている政治家（やその支持者）は多そうだ。

[^quran2]: 「[イスラームの宗教と脳の機能は交差する。](http://30th.rcast.u-tokyo.ac.jp/future/future01.html)」には「最後」について「宗教的に画期的な「発明」」と書かれている： 「イスラーム教徒は「最後の預言者であるムハンマドに託された最後の啓示の言葉」としてコーランを認識することで、「最後の次」の啓示というものが出てくることを、単に認識しません。誰かがどれだけよく考え抜いて、現代のグローバル化した人類社会の新たな環境に適合した新たな啓示の法はこうだ、と新しい宗教を提示しても、コーランの内容と比べてその新しい宗教が優れているか否か、ということをイスラーム教徒の側は誰も論じません。論じる必然性を全く感じないからです」

憲法に関しては，私はそっちに与しないが。

たとえば，上で挙げた TDD は「テストは正しい」ことが必要条件である。
「テストは正しい」と信じられなければ，そもそも TDD は成立しないのだ。

エンジニアがアルゴリズム偏重になることについては苦言を呈する人もいる。

{{< fig-quote title="頼むからプログラミングを学ばないでくれ | TechCrunch Japan" link="http://jp.techcrunch.com/2016/05/17/20160510please-dont-learn-to-code/" >}}
<q>プログラミングにおける何らかの問題に取り組むとき、まず私たちはその問題が何であるのか、そしてそれは本当に問題であるのかを見極めなければならない。その問題が本当にプログラミングで解決できる問題かどうかを考慮せず、プログラミングで解決することに固執し、「なぜ問題なのか」という視点を失ってしまっては、そこから何も得ることはできない。それがプログラミングで解決できる問題であろうと、なかろうとだ。</q>
{{< /fig-quote >}}

これは全くもってそのとおり。
エンジニアリングの世界ではこれを「要求定義」もしくはもう少し推し進めて「要求開発」と呼ぶ。

私たち日本人に日本国憲法を「最後の憲法」たらしめんとする圧力の源は何なのか。
そこから考え始めるべきではないのだろうか。
それも「要求開発」である。

## ブックマーク

- [護憲派は改憲案に賛成すべきだ : 「平和構築」を専門にする国際政治学者](http://shinodahideaki.blog.jp/archives/16834872.html)
    - [立憲民主党の未来は実は改憲にある : 「平和構築」を専門にする国際政治学者](http://shinodahideaki.blog.jp/archives/21571730.html)
    - [長谷部恭男教授の「立憲主義」は、集団的自衛権の違憲性を説明しない : 「平和構築」を専門にする国際政治学者](http://shinodahideaki.blog.jp/archives/21632477.html)

- [第48回衆議院議員総選挙]({{< ref "/remark/2017/10/the-48th-general-election.md" >}})

[Scrapbox]: https://scrapbox.io/spiegel-branch/ "Spiegel's Branch - Scrapbox"
[ATOM]: https://atom.io/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[go-plus]: https://atom.io/packages/go-plus

## 参考図書

<div class="hreview" ><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4822282686/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/512Y77Y5WDL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4822282686/baldandersinf-22/">要求開発~価値ある要求を導き出すプロセスとモデリング</a></dt><dd>山岸 耕二 安井 昌男 萩本 順三 河野 正幸 野田 伊佐夫 平鍋 健児 細川 努 依田 智夫 ［要求開発アライアンス］ </dd><dd>日経BP社 2006-03-02</dd></dl><p class="similar"><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4822283585/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4822283585.09._SCTHUMBZZZ_.jpg"  alt="UMLモデリング入門"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4798121967/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4798121967.09._SCTHUMBZZZ_.jpg"  alt="エリック・エヴァンスのドメイン駆動設計 (IT Architects’Archive ソフトウェア開発の実践)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4492961143/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4492961143.09._SCTHUMBZZZ_.jpg"  alt="ビジネスプロセスの教科書"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4822283496/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4822283496.09._SCTHUMBZZZ_.jpg"  alt="UMLモデリングレッスン"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4274505219/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4274505219.09._SCTHUMBZZZ_.jpg"  alt="ソフトウェアエンジニアリング基礎知識体系 ―SWEBOK V3.0―"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4320023528/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4320023528.09._SCTHUMBZZZ_.jpg"  alt="要求仕様の探検学―設計に先立つ品質の作り込み"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4764904047/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4764904047.09._SCTHUMBZZZ_.jpg"  alt="要求工学知識体系 第1版"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4798114456/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4798114456.09._SCTHUMBZZZ_.jpg"  alt="ユースケース駆動開発実践ガイド (OOP Foundations)"  /></a> </p>
<p class="description">要求は開発するものらしい。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-01-24">2017-01-24</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%81%84%E3%81%8B%E3%81%AB%E3%81%97%E3%81%A6%E5%95%8F%E9%A1%8C%E3%82%92%E3%81%A8%E3%81%8F%E3%81%8B-G-%E3%83%9D%E3%83%AA%E3%82%A2/dp/4621045938?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621045938"><img src="https://images-fe.ssl-images-amazon.com/images/I/51XGP8AFX2L._SL160_.jpg" width="112" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%81%84%E3%81%8B%E3%81%AB%E3%81%97%E3%81%A6%E5%95%8F%E9%A1%8C%E3%82%92%E3%81%A8%E3%81%8F%E3%81%8B-G-%E3%83%9D%E3%83%AA%E3%82%A2/dp/4621045938?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621045938">いかにして問題をとくか</a></dt>
	<dd>G. ポリア</dd>
	<dd>G. Polya (原著), 柿内 賢信 (翻訳)</dd>
    <dd>丸善 1975-04-01</dd>
    <dd>Book 単行本</dd>
    <dd>ASIN: 4621045938, EAN: 9784621045930</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">数学書。というか問いの立てかたやものの考え方についての指南書。のようなものかな。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
