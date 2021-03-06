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

{{% review-paapi "B00FIX664S" %}} <!-- 時計じかけのオレンジ -->
{{% review-paapi "B01J1I8PRQ" %}} <!-- 社会は情報化の夢を見る -->
{{% review-paapi "B00CHIFA1M" %}} <!-- Gene Mapper -->
{{% review-paapi "4488711022" %}} <!-- 万物理論 -->
