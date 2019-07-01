+++
date = "2016-02-09T20:42:43+09:00"
description = "そろそろこういうバカバカしい話はやめて欲しいよね。そのための CC Licenses なんだが。"
draft = false
tags = ["creative-commons", "google"]
title = "Google 八分とファッションとしての CC License"

[scripts]
  mathjax = false
  mermaidjs = false
+++

結城浩さんが懐かしい記事を tweet しておられる。

<blockquote class="twitter-tweet" data-lang="ja"><p lang="ja" dir="ltr">“クリエイティブ・コモンズのライセンスをWeblogツールで使うことの危険性” <a href="https://t.co/XQK35Sj2VH">https://t.co/XQK35Sj2VH</a></p>&mdash; 結城浩 (@hyuki) <a href="https://twitter.com/hyuki/status/696895800601739265">2016, 2月 9</a></blockquote>

- [クリエイティブ・コモンズのライセンスをWeblogツールで使うことの危険性](http://www.hyuki.com/trans/blogtrap.html)
- [で、東浩紀は相手にちゃんと問題を指摘したのか？ - 2004-05-10 - YAMDAS現更新履歴](http://d.hatena.ne.jp/yomoyomo/20040510#p3)

昔「ファッションとしての Creative Commons」が云々と言った人がいて上の記事のような議論になったのだが，その後 [Creative Commons] 本家は Free Culture に舵を切ったため，この議論自体がどうでもいいものになってしまった。
まぁ日本では [Creative Commons] 以上に Free Culture は知名度が低いのだが（笑）

色々上がってくる記事とかを見ていると，どうも DMCA を盾に Google 八分に遭った人がいるらしい。
いやぁ，これも懐かしいな。

ちなみに Google の検索結果やその順位はフェアでも何でもない。
Google はこの手の係争を避けようとするので，誰かにとって気に入らないページが検索結果に上がって Google にクレームが来れば，基本的に削除するかランクを大幅に下げる操作を行う。

でもこれは Google に限らない。
たとえば Facebook はタイムラインの順番をよく分からないロジックで勝手に入れ替える。
Twitter は気に入らないアカウントを平気で削除する。
最近も[テロリストだと主張して12万アカウントを削除](http://www.mdn.co.jp/di/newstopics/44028/)した。

ここで言論の自由が云々とか hate speech が云々とか言うつもりはないが，これらのサービスは所詮商業サービスであり，彼らにとってそれがフェアであるかどうかは優先順位としてそれほど高くないということだ。
言い換えれば「安心」というのは極めてアンフェアなものだということである[^0]。

[^0]: 某国首相は「安心・安全」という言葉を安直に口にするが，自分の言っていることの意味をよく考えてほしいものである。

ところで私は「[CC Licenses] をファッションで付けて何が悪い」と思ってる人なので，バンバン付けて欲しい。
そもそも権利のないコンテンツに勝手にライセンスを付けても無効になるだけのことである[^a]。
リスクヘッジを考えるのならフリー素材であってもオリジナルを指示するようクレジット表示を行えばよい。

[^a]: ちなみに [CC Licenses] はサブライセンスを禁止している。複製・配布したものに別のライセンスを付けることは出来ない。

ちなみに [CC Licenses] のクレジット表示は 4.0 International で緩和された。

{{< fig-quote title="What's New in 4.0 - Creative Commons - Creative Commons" link="https://creativecommons.org/Version4/" lang="en" >}}
<q>Version 4.0 includes a slight change to <a href="http://wiki.creativecommons.org/Frequently_Asked_Questions#How_do_I_properly_attribute_material_offered_under_a_Creative_Commons_license.3F">attribution requirements</a>, designed to better reflect accepted practices. The licenses explicitly permit licensees to satisfy the attribution requirement with a link to a separate page for attribution information. This was already common practice on the internet and possible under earlier versions of the licenses, and Version 4.0 alleviates any uncertainty about its use.</q>
{{< /fig-quote >}}

Legal Code でも

{{< fig-quote title="クリエイティブ・コモンズ (Creative Commons) — 表示-継承 4.0 国際" link="https://creativecommons.org/licenses/by-sa/4.0/legalcode.ja" >}}
<q>あなたがライセンス対象物を共有する媒体・方法・文脈に照らして、いかなる合理的な方法でも満たすことができます。
例えば、必要とされる情報を含むリソースのURIやハイパーリンクを付すことで条件を満たすことが合理的な場合があります。</q>
{{< /fig-quote >}}

の文言が追加されている。
要は他者の作品に敬意を払い，原典にあたれる手段をきちんと示すことが重要なのだ。
それができていれば割とラフな運用でも構わないと思う。

あと， [CC Licenses] では一度設定したライセンスを撤回することは出来ない[^b]。
なのでフリーで公開しているものを「フリーやめたから掲載するな」とか理不尽な事にはならないのでご安心を。
条件がフリーに見えても撤回の余地を残しているのなら，それは「フリー」とは言わない。

[^b]: ただし許諾者が未成年の場合は法定代理人によって取り消される場合がある。

そろそろこういうバカバカしい話はやめて欲しいよね。
そのための [CC Licenses] なんだが。

そうそう。
このサイトで [CC Licenses] の不定期連載やってます。

- [改訂3版： CC Licenses について — text.Baldanders.info](/cc-licenses/)

[人格権]({{< ref "/cc-licenses/03-moral-rights.md" >}})と[二次的著作物]({{< ref "/cc-licenses/04-derivative-works.md" >}})の話までは書いた[^c]。
今後は DRM （技術的保護手段）についてと CC0 および Free Culture Licenses について書く予定。
予定は未定。

[^c]: 私は法律の専門家ではないので，勘違いや間違いがあれば指摘していただけると嬉しいです。

以上，宣伝でした。

[Creative Commons]: https://creativecommons.org/ "Creative Commons"
[CC Licenses]: https://creativecommons.org/licenses/ "ライセンスについて - Creative Commons"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%82%BB%E3%83%9E%E3%83%B3%E3%83%86%E3%82%A3%E3%83%83%E3%82%AF-HTML-XHTML-%E7%A5%9E%E5%B4%8E-%E6%AD%A3%E8%8B%B1/dp/483993195X?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=483993195X"><img src="https://images-fe.ssl-images-amazon.com/images/I/51oaN2iq9xL._SL160_.jpg" width="124" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%82%BB%E3%83%9E%E3%83%B3%E3%83%86%E3%82%A3%E3%83%83%E3%82%AF-HTML-XHTML-%E7%A5%9E%E5%B4%8E-%E6%AD%A3%E8%8B%B1/dp/483993195X?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=483993195X">セマンティック HTML/XHTML</a></dt>
	<dd>神崎 正英</dd>
    <dd>毎日コミュニケーションズ 2009-05-28</dd>
    <dd>Book 単行本（ソフトカバー）</dd>
    <dd>ASIN: 483993195X, EAN: 9784839931957</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">残念ながら紙の本は実質的に絶版なんですよねぇ。是非デジタル化を希望します。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2014-08-17">2014-08-17</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%82%AA%E3%83%BC%E3%83%97%E3%83%B3%E5%8C%96%E3%81%99%E3%82%8B%E5%89%B5%E9%80%A0%E3%81%AE%E6%99%82%E4%BB%A3-%E8%91%97%E4%BD%9C%E6%A8%A9%E3%82%92%E6%8B%A1%E5%BC%B5%E3%81%99%E3%82%8B%E3%82%AF%E3%83%AA%E3%82%A8%E3%82%A4%E3%83%86%E3%82%A3%E3%83%96%E3%83%BB%E3%82%B3%E3%83%A2%E3%83%B3%E3%82%BA%E3%81%AE%E6%96%B9%E6%B3%95%E8%AB%96-%E3%82%AB%E3%83%89%E3%82%AB%E3%83%AF%E3%83%BB%E3%83%9F%E3%83%8B%E3%83%83%E3%83%84%E3%83%96%E3%83%83%E3%82%AF-%E3%83%89%E3%83%9F%E3%83%8B%E3%82%AF%E3%83%BB%E3%83%81%E3%82%A7%E3%83%B3-ebook/dp/B00DI8TMPU?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00DI8TMPU"><img src="https://images-fe.ssl-images-amazon.com/images/I/51zmlOAOaFL._SL160_.jpg" width="106" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%82%AA%E3%83%BC%E3%83%97%E3%83%B3%E5%8C%96%E3%81%99%E3%82%8B%E5%89%B5%E9%80%A0%E3%81%AE%E6%99%82%E4%BB%A3-%E8%91%97%E4%BD%9C%E6%A8%A9%E3%82%92%E6%8B%A1%E5%BC%B5%E3%81%99%E3%82%8B%E3%82%AF%E3%83%AA%E3%82%A8%E3%82%A4%E3%83%86%E3%82%A3%E3%83%96%E3%83%BB%E3%82%B3%E3%83%A2%E3%83%B3%E3%82%BA%E3%81%AE%E6%96%B9%E6%B3%95%E8%AB%96-%E3%82%AB%E3%83%89%E3%82%AB%E3%83%AF%E3%83%BB%E3%83%9F%E3%83%8B%E3%83%83%E3%83%84%E3%83%96%E3%83%83%E3%82%AF-%E3%83%89%E3%83%9F%E3%83%8B%E3%82%AF%E3%83%BB%E3%83%81%E3%82%A7%E3%83%B3-ebook/dp/B00DI8TMPU?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00DI8TMPU">オープン化する創造の時代　著作権を拡張するクリエイティブ・コモンズの方法論 (カドカワ・ミニッツブック)</a></dt>
	<dd>ドミニク・チェン</dd>
    <dd>ブックウォーカー 2013-06-25 (Release 2013-06-27)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00DI8TMPU</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description"><a href='https://baldanders.info/blog/000643/'>手軽に読める</a>。お薦め。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2014-09-13">2014-09-13</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>
