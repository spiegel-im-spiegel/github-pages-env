+++
title = "「オブジェクト指向」の黒歴史"
date = "2018-10-11T18:54:19+09:00"
update = "2018-10-21T13:34:31+09:00"
description = "オブジェクト指向の価値を利便性に置くという考え方には激しく同意する。"
image = "/images/attention/kitten.jpg"
tags = [ "engineering", "object-oriented", "design", "programming", "language" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

（もちろんタイトルは釣りです）

Qiita で面白い記事を見かけた。

- [まだ「オブジェクト指向はこうあるべき」で消耗してるの?](https://qiita.com/gyu-don/items/09db0a298136debfe757)

この記事にあるオブジェクト指向の価値を利便性に置くという考え方には激しく同意する。
その上で，私の黒歴史を交えて，戯れ言をいくつか書いてみよう。

## 「プログラミング」で最初に何を学びましたか？

最近の子らは「プログラミング」で何を学ぶのだろう。

私が学生の頃（昭和時代）， FORTRAN の授業で真っ先にやらされたのがフローチャートを書くことだった[^p1]。
ちなみに授業の最初のお題は素数を求めるプログラムだった（ありがち`w`）。

[^p1]: 今どきの子は知らないかも知れないので念のために解説すると，「フローチャート」というのはコンピュータの処理の流れ（工程）を図式した一種のモデリング言語と思っていただいて構わない。

そもそも「コンピュータ」っていうのは，名前の通り，「計算をする機械」なわけですよ。
つまりフローチャートというのは「計算する」ことをモデル化したものなわけ。
だから「計算する」ことに関してはフローチャートで全て表現できる。

そう思ってた時代がありました。

それから，まぁ，[紆余曲折](https://baldanders.info/blog/000529/ "私はこうしてプログラミングを覚えた — Baldanders.info")あって（バブル最盛期に）某システムハウスに潜り込んだのだが，初仕事の設計書で書かせられたのはフローチャートではなく状態遷移表だった。
むしろ「フローチャートなんか要らん（コードを見れば分かる）」と言われましたよ。
これが「社会の現実」ってやつですね，分かります。

ここで社会に出たての小僧は気づくわけですよ，「コンピュータ」っていうのは「**計算** をする機械」ではなく「**情報** を処理する機械」なんだということを。
そして少年は「オブジェクト指向」と出会う（笑）

## 「情報」ってなに？

少なくともノイマン型コンピュータにおいては，基本的に以下の3つの機能しかない。

- 指定したアドレスから命令をフェッチして計算する
- 指定したアドレスからデータを読む
- 指定したアドレスへデータを書く

指定したアドレス（メモリとは限らない）に何があるのか決めるのはコードを書く人間側の責務であり，それを可能とするプログラミング言語であれば「情報を処理する」ことに関しては何でもできる。

じゃあ「情報」ってなに？

「オブジェクト指向」以前の「構造化プログラミング」の時代において，情報とは「データ」と「機能」だった。
しかし「データ」と「機能」だけではコードは簡単に破綻する。
以前に「[ハード屋が書く C ソースコードが凄まじかった思い出]({{< ref "/remark/2016/06/code-by-hardware-engineer.md" >}})」という記事を書いたが，これは情報を「データ」と「機能」と考えた場合の極端例と言える。

では「構造化プログラミング」時代のソフトウェア・エンジニアのコードは何故破綻しなかったのか[^p2]。
ソフトウェア・エンジニアが暗黙的に考えていたこととはなにか。
その答えのひとつが情報を「オブジェクト」と考える「オブジェクト指向」設計ないしはプログラミングである。

[^p2]: いや「動かないコンピュータ」なんてザラにあったけど。むしろ今も見かけるけど（笑）

「オブジェクト指向」設計ではオブジェクトを以下の3つの組み合わせと考えた。

- 名前
- 状態（属性）
- 機能（操作，手段）

特に重要なのが「名前」である。
「名前」とは自と他を区別する識別子で，区別することでそこに「関係」が生まれる。

たとえば「犬は動物の一種である」というのは，「犬」と「動物」の関係が「一種である[^isa1]（is-a）」ことを示す。
C++ や Java では「一種である」という関係を「クラス・オブジェクトの継承」という形で記述するが，継承は「オブジェクト指向」に必須の記述ではない。
Go 言語の構造的部分型（structural subtyping）のような記述だってあるのだ[^exc1]。

[^isa1]: ちなみに「A は B の一種である」という関係を「汎化（または特化）」と言う。 “A is a B” と当て嵌めて考えれば分かりやすいだろう。
[^exc1]: むしろ継承にこだわって菱形継承のような弱点や例外処理のような歪な構造を生み出したことは初期のオブジェクト指向プログラミング言語の黒歴史と言えるだろう。

つまり「オブジェクト指向」においてはオブジェクトとその関係をモデル化することが重要で，それを記述することができるのであればどんな言語でも構わないのである。

ちなみに「オブジェクト指向」プログラミングはアセンブラや C 言語でも記述できるし実際にそういうプロジェクトに関わったこともある（大昔の話だよ）。
プログラミング言語はプログラミングの手段に過ぎないということだ。

## 「オブジェクト指向」の先へ

しかし今さらアセンブラや C 言語で「オブジェクト指向」なコードを書こうという人は少ない（いや，ほぼいない？）だろう。
では，当時オブジェクト指向プログラミング言語と呼ばれた C++ や Java によって私達ソフトウェア・エンジニアは幸せになれたのだろうか？ 否である！

これは私見だけど，「オブジェクト指向」の恩恵は「オブジェクト指向」で設計・記述できるようになったことではなく「オブジェクト指向」からの派生でコード記述が文芸的[^bp1] になり意図や文脈を記述するようになってきたことだと思う。

こうした発展はプログラミング言語のトレンドがオブジェクト指向言語だけではなく関数型言語など複数のパラダイムをブレンドした「マルチパラダイム・プログラミング言語」へシフトしていることからも言えるんじゃないだろうか。
また FOSS が主流となりプログラムコードがエンジニア同士の対話手段として使われるようになったことも影響しているかもしれない[^cr1]。

[^cr1]: まさしく著作権による知的独占からの解放がもたらした好例のひとつですな。
[^bp1]: これはいわゆる「文芸的プログラミング」とは意味が異なる。紛らわしくてゴメンペコン。

そういう意味で「オブジェクト指向はこうあるべき」という議論は不毛で非生産的な行為と言える。
私達はもはやその先に足を踏み入れているのだから。

## ブックマーク

- [バッドノウハウと「奥が深い症候群」](http://0xcc.net/misc/bad-knowhow.html)
- [なぜPythonはこんなにも遅いのか？ | POSTD](https://postd.cc/why-is-python-so-slow/)

- [「null 安全」について]({{< ref "/remark/2016/11/null-safety.md" >}})
- [きみは Generics がとくいなフレンズなんだね，または「制約は構造を生む」]({{< ref "/remark/2017/03/generics-vs-duck-typing.md" >}})
- [「自動販売機の気持ちになって考える」]({{< ref "/remark/2017/11/do-you-understand-what-a-computer-is.md" >}})

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%81%82%E3%81%AA%E3%81%9F%E3%81%AF%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%82%BF%E3%82%92%E7%90%86%E8%A7%A3%E3%81%97%E3%81%A6%E3%81%84%E3%81%BE%E3%81%99%E3%81%8B-10%E5%B9%B4%E5%BE%8C%E3%80%8120%E5%B9%B4%E5%BE%8C%E3%81%BE%E3%81%A7%E5%BF%85%E3%81%9A%E5%BD%B9%E7%AB%8B%E3%81%A4%E6%A0%B9%E3%81%A3%E3%81%93%E3%81%AE%E9%83%A8%E5%88%86%E3%81%8C%E3%81%8D%E3%81%A3%E3%81%A1%E3%82%8A%E3%82%8F%E3%81%8B%E3%82%8B%EF%BC%81-%E3%82%B5%E3%82%A4%E3%82%A8%E3%83%B3%E3%82%B9%EF%BD%A5%E3%82%A2%E3%82%A4%E6%96%B0%E6%9B%B8-%E6%A2%85%E6%B4%A5-%E4%BF%A1%E5%B9%B8/dp/4797339497?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4797339497"><img src="https://images-fe.ssl-images-amazon.com/images/I/51W3fP3Q%2BtL._SL160_.jpg" width="102" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%81%82%E3%81%AA%E3%81%9F%E3%81%AF%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%82%BF%E3%82%92%E7%90%86%E8%A7%A3%E3%81%97%E3%81%A6%E3%81%84%E3%81%BE%E3%81%99%E3%81%8B-10%E5%B9%B4%E5%BE%8C%E3%80%8120%E5%B9%B4%E5%BE%8C%E3%81%BE%E3%81%A7%E5%BF%85%E3%81%9A%E5%BD%B9%E7%AB%8B%E3%81%A4%E6%A0%B9%E3%81%A3%E3%81%93%E3%81%AE%E9%83%A8%E5%88%86%E3%81%8C%E3%81%8D%E3%81%A3%E3%81%A1%E3%82%8A%E3%82%8F%E3%81%8B%E3%82%8B%EF%BC%81-%E3%82%B5%E3%82%A4%E3%82%A8%E3%83%B3%E3%82%B9%EF%BD%A5%E3%82%A2%E3%82%A4%E6%96%B0%E6%9B%B8-%E6%A2%85%E6%B4%A5-%E4%BF%A1%E5%B9%B8/dp/4797339497?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4797339497">あなたはコンピュータを理解していますか? 10年後、20年後まで必ず役立つ根っこの部分がきっちりわかる！ (サイエンス･アイ新書)</a></dt>
	<dd>梅津 信幸</dd>
    <dd>ソフトバンククリエイティブ 2007-03-16</dd>
    <dd>Book 新書</dd>
    <dd>ASIN: 4797339497, EAN: 9784797339499</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">2002年に技術評論社から出た同名タイトルのリニューアルらしい。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-11-24">2017-11-24</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%95%B0%E5%AD%A6%E3%82%AC%E3%83%BC%E3%83%AB%EF%BC%8F%E4%B9%B1%E6%8A%9E%E3%82%A2%E3%83%AB%E3%82%B4%E3%83%AA%E3%82%BA%E3%83%A0-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B00I8AT1FO?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00I8AT1FO"><img src="https://images-fe.ssl-images-amazon.com/images/I/41353H%2BBzFL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%95%B0%E5%AD%A6%E3%82%AC%E3%83%BC%E3%83%AB%EF%BC%8F%E4%B9%B1%E6%8A%9E%E3%82%A2%E3%83%AB%E3%82%B4%E3%83%AA%E3%82%BA%E3%83%A0-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B00I8AT1FO?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00I8AT1FO">数学ガール／乱択アルゴリズム</a></dt>
	<dd>結城 浩</dd>
    <dd>SBクリエイティブ 2011-02-25 (Release 2014-03-12)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00I8AT1FO</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">工学ガール，リサちゃん登場！</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-04-19">2015-04-19</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
