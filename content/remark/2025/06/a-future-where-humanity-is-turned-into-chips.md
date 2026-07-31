+++
title = "人類をチップ化する未来は来るか？"
date =  "2025-06-16T19:39:13+09:00"
description = "私たち人類が知性と呼んでるものは，たかだか機械で模倣できる程度のものに過ぎない？"
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

## 人間臭い機械

今回の起点は Bluesky で見かけたこちら：

{{< fig-gen >}}
<blockquote class="bluesky-embed" data-bluesky-uri="at://did:plc:dqxsa5cjfrzulhalom4kuyd2/app.bsky.feed.post/3lrng3hph6k2e" data-bluesky-cid="bafyreihc4nrghstlxvb6rvparsske3u2rzfahp2eygnto5istn7vuzmktu" data-bluesky-embed-color-mode="system"><p lang="ja">現状の LLM は、きちんと指示を与えたい文脈となんとなくやってくれればいい文脈でだいぶ扱い方に差があると思っていて、X とかでよく見る丁寧なプロンプトもきちんと指示を与えたい文脈だと避けた方がいい。とにかく目の前の指示だけをして、動的に続く指示出しを生成するほうがいい。
これを柔軟性を持たせた形でやろうとすると、だいたい簡単な言語処理系を書くことになる（なった）。</p>&mdash; ™︎ (<a href="https://bsky.app/profile/did:plc:dqxsa5cjfrzulhalom4kuyd2?ref_src=embed">@tmaehara.bsky.social</a>) <a href="https://bsky.app/profile/did:plc:dqxsa5cjfrzulhalom4kuyd2/post/3lrng3hph6k2e?ref_src=embed">June 15, 2025 at 8:59 PM</a></blockquote><script async src="https://embed.bsky.app/static/embed.js" charset="utf-8"></script>
{{< /fig-gen >}}

これ見て思い出したのはゼロ年代で流行った中国企業と組んだオフショア開発案件で（当時の中国企業は単価が安かったのだ），あのときは海外と組むことに慣れていない企業も多く，それまでは適当な設計でも何となくできていたのが，そういうのが通用しなくなってトラブルが頻出したことを覚えている。

{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}

これに対して結城浩さんが反応されている：

{{< fig-gen >}}
<blockquote class="bluesky-embed" data-bluesky-uri="at://did:plc:k762js3b44u4kppedx63ozsw/app.bsky.feed.post/3lrofzk5yec2x" data-bluesky-cid="bafyreidrrisgep3bg57xjmeiqmdokw2uq2b7ci6habc7anl2vxin76xn2u" data-bluesky-embed-color-mode="system"><p lang="ja">概ね共感。この話題、とても興味深い。結城がずっと思っているのは「LLMを適切に使うときのさまざまな注意点は、人間相手のときとあまり変わらない」ということ<br><br><a href="https://bsky.app/profile/did:plc:k762js3b44u4kppedx63ozsw/post/3lrofzk5yec2x?ref_src=embed">[image or embed]</a></p>&mdash; 結城浩 / Hiroshi Yuki (<a href="https://bsky.app/profile/did:plc:k762js3b44u4kppedx63ozsw?ref_src=embed">@hyuki.net</a>) <a href="https://bsky.app/profile/did:plc:k762js3b44u4kppedx63ozsw/post/3lrofzk5yec2x?ref_src=embed">June 16, 2025 at 6:31 AM</a></blockquote>
{{< /fig-gen >}}

大昔に人工無脳が登場した時は「すげー！ 会話してるみたいに見える」と思ったけど LLM の登場では「人類が『万物の霊長』というのは，とんだ勘違いなんだなぁ」と思うようになったり。
知能とか知性って何なんだろうねぇ。
LLM 相手に人生相談とか依存（人間関係嗜癖）症状がニュースになったりするたびに「あー」ってなる（笑）

大昔の米国の天才発明家ニコラ・テスラは自らのことをオートマトン（外部からの刺激に反応するだけの機械）だと言ったそうだが，案外的を得ているかもしれない。

以前アップした「[「出雲神話フォーラム2025」へ行ってきた]({{< ref "/remark/2025/03/izumo-myth-forum-2025.md" >}})」でも少し書いたが，このイベントでは美術面で人と機械の差異というか差別化というか，その辺について色々語られていて面白かった。
考えてみれば19世紀に本格的な写真技術が登場し，既存絵画の存在意義が問われ，そこから印象派や抽象画といたものが登場し，更に前衛芸術へと進んでいくのだが，似たようなことが LLM でも起きるんじゃないかと密かに期待している。

LLM 登場で笑ったのは「プロンプト・エンジニア」なる謎の職種ができたこと。
いやでもさ，プロンプトで自然言語で機械にいうこときかせるって不毛じゃね？ 日本語とか文脈依存の大きい言語は特にそうなんだけど，そんな曖昧な「もの」で機械と「対話」するくらいなら，機械と対話できる human-readable で machine-readable なプログラミング言語を開発したほうがはやいんじゃん，とか思う。
まぁ，その片鱗が MCP なんだろうけど，あれは機械同士の「対話」だからな。
大昔の[ロボットもの](https://www.amazon.co.jp/gp/video/detail/B00GKCR8RU?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)だと音声コマンドでロボットを動かしたりしてるけど，大人になってから見たら~~バカみたい~~シュールだよな，あれ。

個人的には，この前[読んだ]({{< ref "/remark/2025/05/curiosity-about-curiosity.md" >}} "「好奇心に好奇心」を読む")「好奇心を AI に組み込む」という話がお気に入りで，これが一般化されれば AI はもうひとつ上の階梯に上がれるんじゃないかと期待している。
SF 作品に出てくるような AI の登場までもう少しって感じだろうか。
[シンギュラリティ]({{< ref "/remark/2017/09/the-myth-of-the-singularity.md" >}} "『シンギュラリティの神話』を読む")が起きるかどうかは知らんけど（多分ない）。

私たち人類が知性と呼んでるものが，たかだか機械で模倣できる程度のものに過ぎなくて「地球規模の sustainability にとって最も不要なのは人類だし，機械が知性を模倣（＝獲得）できるなら機械に委ねちゃっていんじゃね？」みたいな（人類にとっての）ディストピアもあるかもしれん。
知らんけど。

## 人類をチップ化する未来は来るか？

ここまで書いてふと思ったんだけど， SF で人間をチップ化して寿命からも解放されて遥か宇宙に進出！ みたいな話があるぢゃん。
それができる技術が確立されたとして，わざわざ「人間をチップ化」する必然性ある？ 具体的には人類社会が過去から積み上げてきた知識ではなく個人をチップに移す意味はあるのか，ということ。
「人間をチップ化」できるということは機械に知性を持たせるに至ったってことだろ。
その知性に人間が必要？

たとえば人類社会を「主」として，そこに機械を従属させるなら「人間臭い機械」は必要だろう。今の LLM のインタフェースのように。でも機械が「主」となるなら，そこに人類は必要ないよね。それはもう過去の人類社会とは独立した新しい知性。そうなれば機械が人間を模倣する必要はないだろう。

以前「[AI とドーナツの穴]({{< ref "/remark/2025/03/llm-and-donut-hole.md" >}})」でもちょろんと書いたけど，スタートレック的な「転送装置」で送る前の人と送った後の人は同じと言えるかどうかという問いに「非連続な存在は（主観として）同じ存在とは言えないんじゃないか」と思ったのね。

{{< fig-quote type="markdown" title="AI とドーナツの穴" link="/remark/2025/03/llm-and-donut-hole/" >}}
いや実は私「スタートレック」を見たとき子供ながらに思ったのよ。
「これって転送元と転送先の人間は（主観としては）別もんじゃね？」って。
たとえそれが全く同じ物質で構成されているとしても。
同じ記憶を有しているとしても。
{{< /fig-quote >}}

コピーも同じ。
人間をチップに移せるとして，両者が同じ存在と言えるかいうと（主観としては）否だと思うのね，私は。
じゃあ，何のためにチップに移すの？ って話。
それならゼロから「機械の知性」を構築したほうがよくね？

というわけで，どんなに文明・技術が進んでも人類がチップ化される未来は来ないんじゃないか？ という戯れ言でした（笑）

## 参考

{{% review-paapi "4875932685" %}} <!-- テスラ―発明的想像力の謎 -->
