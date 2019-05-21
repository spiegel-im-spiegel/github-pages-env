+++
title = "時計じかけのプログラミング"
date =  "2019-05-20T23:09:39+09:00"
description = "むしろ人工知能が書いたコードを人間が読めるように「リバース・エンジニアリング」するのに苦労するんじゃないのかな。"
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "engineering" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

実家に{{< ruby "かえ" >}}帰郷{{< /ruby >}}って以前と変わったことは色々あるが，そのうちのひとつが「テレビを観ている人」を観察するようになったことだ。
いや，うちの親父殿と甥っ子がどえらテレビっ子になってるのよ。

テレビ番組というのは視聴者の感情をコントロールすることで成立している。
視聴者は番組が笑ってほしいところで笑い，怒ってほしいところで怒り，驚いてほしいところで驚き，泣いてほしいところで泣く。
そりゃあもう見事にそのとおりに行動するのよ，「テレビを観ている人」は。

独り暮らしだと「テレビを観ている人」を観ることはないからねぇ。
私自身もそうした視聴者に違いないのだが，「テレビを観ている人」でフィルタリングされたテレビ番組は違う意味で滑稽に見える。

昭和時代の「知識人」はテレビの台頭を指して「一億総白痴化」と揶揄したが，もしかしたら「白痴化」はテレビのみならず（Web を含む）あらゆるメディアに及んでいるのかもしれない。

そういったものを観察していると，人工知能の「シンギュラリティ」はひょっとしたら起こるのかもしれない，と思うようになった。
それは「人工知能が人類を凌駕する」という意味ではなく「人工知能が人類を白痴化する」という意味で（「[マトリックス](https://www.amazon.co.jp/exec/obidos/ASIN/B00FIWIMOG/baldandersinf-22/)」だな`w`）。
案外，人工知能によって真っ先に職を奪われるのはメディアの「中の人」なのかもねぇ。
人間の感情をコントロールする術が（確率・統計的に）分かっているのなら，それを最もうまく使いこなすのは機械だろうから。

{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}。

- [次世代のプログラミングツール、未来のプログラミング言語の方向性について - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20190519/programmingtools)

20世紀末の話なのでもう時効だと思うけど，通信機械のシミュレータのプロジェクトに参加してたことがあって，そのプロジェクトが面白かったのはプロトコルスタックを GUI の CASE (Computer Aided Software Engineering) ツールで組むことができるというものだった。

20世紀末はいわゆる「オブジェクト指向」が産業分野でも台頭し始めた時期であり，その「オブジェクト指向」の究極の夢が図像化したオブジェウトを線でつなぐだけでプログラムができるという CASE ツールだったわけ。

でも，これらがメインストリームになることは殆どなかった。
私の知ってる範囲だと Yahoo! Pipes くらいなもんじゃないのかな（Yahoo! Pipes が登場し廃れるのはゼロ年代だけど）。
あるいは [Astah*] のように UML 図からプログラム・コードを生成するツールとか[^m1]。

[^m1]: そういや数式を書いたらプログラムに変換してくれるっていうツールもあったな。いや，今もあるか。 AI 絡みならそちらのほうがよほど現実味がある（笑）

でも，たとえば UML 図にしたって実際に図を描いてみるとかったるくてやってられないのですよ。
で，結局はテキストで [PlantUML] で「書いた」ほうが早いという話になり，それならいっそ [PlantUML] のコードからプログラム・コードを生成すればいいじゃない，という話になってくる。

「図を描く」ってのは案外に不自由なものである。
何故なら「図を描く」ことによってそれまでの思考が図に固定されてしまうから。
その先はない，やり直さない限り。
だからかったるいのだ。
図像化と符号化には天地ほどのひらきがある。

テキストデータのプログラムコードが紙テープやパンチカードと決定的に違うのは記述コストが圧倒的に低くてしかも human-readable である点だろう。
考えながら書き，書きながら考えることができる。
これを超える「発明」はそれこそ数十年単位で出てこないんじゃないだろうか。

たとえば「プログラムコード」を（人間が指示したとおりに）人工知能が書くことができるようになったとしよう。
その場合，コードがテキストや図である必然性は微塵もなく machine-readable でありさえすればいい。
塩基配列でコードを構成したって構わないだろう。
むしろ人工知能が書いたコードを人間が読めるように「リバース・エンジニアリング」するのに苦労するんじゃないのかな。

グレッグ・イーガンの『[万物理論](https://www.amazon.co.jp/exec/obidos/ASIN/4488711022/baldandersinf-22/)』では「万物理論」を人工知能が書いている。
人間は理論の骨子とデータを入力するだけで，あとは人工知能が論文を書き上げるのをバカンスでもしながら待つだけである。
まさに理想の研究者生活。
未来のプログラマもかく在りたいものである（笑） まぁ，その時代まで私は生きていないだろうけど。

[Astah*]: http://astah.net/ "Astah - Software Design Tools for Agile teams with UML, ER Diagram, Flowchart, Mindmap and More | Astah.net"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw beautiful UML diagrams."

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%99%82%E8%A8%88%E3%81%98%E3%81%8B%E3%81%91%E3%81%AE%E3%82%AA%E3%83%AC%E3%83%B3%E3%82%B8-%E5%AD%97%E5%B9%95%E7%89%88-%E3%83%9E%E3%83%AB%E3%82%B3%E3%83%A0%E3%83%BB%E3%83%9E%E3%82%AF%E3%83%89%E3%82%A6%E3%82%A7%E3%83%AB/dp/B00FIX664S?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00FIX664S"><img src="https://images-fe.ssl-images-amazon.com/images/I/41qjhvuarDL._SL160_.jpg" width="120" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%99%82%E8%A8%88%E3%81%98%E3%81%8B%E3%81%91%E3%81%AE%E3%82%AA%E3%83%AC%E3%83%B3%E3%82%B8-%E5%AD%97%E5%B9%95%E7%89%88-%E3%83%9E%E3%83%AB%E3%82%B3%E3%83%A0%E3%83%BB%E3%83%9E%E3%82%AF%E3%83%89%E3%82%A6%E3%82%A7%E3%83%AB/dp/B00FIX664S?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00FIX664S">時計じかけのオレンジ (字幕版)</a></dt>
	<dd>スタンリー・キューブリック (プロデュース)</dd>
    <dd> (Release 2013-11-26)</dd>
    <dd>Movie Prime Video</dd>
    <dd>ASIN: B00FIX664S</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">この映画を観たのって大学生の頃なんだけど，女性の先輩と2人で観に行ったんだよなぁ。後で考えたらものすごいシチュエーションだったんだけど，その時は映画が面白すぎて状況を全く把握してなかった（笑）</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-05-20">2019-05-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E7%A4%BE%E4%BC%9A%E3%81%AF%E6%83%85%E5%A0%B1%E5%8C%96%E3%81%AE%E5%A4%A2%E3%82%92%E8%A6%8B%E3%82%8B-%E6%B2%B3%E5%87%BA%E6%96%87%E5%BA%AB-%E4%BD%90%E8%97%A4%E4%BF%8A%E6%A8%B9-ebook/dp/B01J1I8PRQ?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B01J1I8PRQ"><img src="https://images-fe.ssl-images-amazon.com/images/I/5132MaXTkGL._SL160_.jpg" width="112" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E7%A4%BE%E4%BC%9A%E3%81%AF%E6%83%85%E5%A0%B1%E5%8C%96%E3%81%AE%E5%A4%A2%E3%82%92%E8%A6%8B%E3%82%8B-%E6%B2%B3%E5%87%BA%E6%96%87%E5%BA%AB-%E4%BD%90%E8%97%A4%E4%BF%8A%E6%A8%B9-ebook/dp/B01J1I8PRQ?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B01J1I8PRQ">社会は情報化の夢を見る (河出文庫)</a></dt>
	<dd>佐藤俊樹</dd>
    <dd>河出書房新社 2010-09-03 (Release 2016-07-29)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B01J1I8PRQ</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">1996年に出版された『ノイマンの夢・近代の欲望―情報化社会を解体する』の改訂新装版。しかし内容はこれまでと変わりなく，繰り返し語られる技術決定論を前提とする安易な未来予測を「情報化」社会論だとして批判する。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-31">2018-12-31</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/Gene-Mapper-full-build-%E8%97%A4%E4%BA%95-ebook/dp/B00CHIFA1M?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00CHIFA1M"><img src="https://images-fe.ssl-images-amazon.com/images/I/516s6S%2Bhv1L._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/Gene-Mapper-full-build-%E8%97%A4%E4%BA%95-ebook/dp/B00CHIFA1M?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00CHIFA1M">Gene Mapper -full build-</a></dt>
	<dd>藤井 太洋</dd>
    <dd>早川書房 2013-04-25 (Release 2013-04-24)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00CHIFA1M</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">SF が特に好きというわけではないのだが（子供の頃は好んで読んでたけど），たまにこうして良質の SF を読むのは楽しいね。日本でEブックによる「自己出版」の魁となったのが作品。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-05-20">2019-05-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E4%B8%87%E7%89%A9%E7%90%86%E8%AB%96-%E5%89%B5%E5%85%83SF%E6%96%87%E5%BA%AB-%E3%82%B0%E3%83%AC%E3%83%83%E3%82%B0%E3%83%BB%E3%82%A4%E3%83%BC%E3%82%AC%E3%83%B3/dp/4488711022?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4488711022"><img src="https://images-fe.ssl-images-amazon.com/images/I/51J3DEJJ1TL._SL160_.jpg" width="112" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E4%B8%87%E7%89%A9%E7%90%86%E8%AB%96-%E5%89%B5%E5%85%83SF%E6%96%87%E5%BA%AB-%E3%82%B0%E3%83%AC%E3%83%83%E3%82%B0%E3%83%BB%E3%82%A4%E3%83%BC%E3%82%AC%E3%83%B3/dp/4488711022?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4488711022">万物理論 (創元SF文庫)</a></dt>
	<dd>グレッグ・イーガン</dd>
	<dd>山岸 真 (翻訳)</dd>
    <dd>東京創元社 2004-10-28</dd>
    <dd>Book 文庫</dd>
    <dd>ASIN: 4488711022, EAN: 9784488711023</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">グレッグ・イーガンの名作。これも singularity を巡る物語だな。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-09-18">2017-09-18</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
