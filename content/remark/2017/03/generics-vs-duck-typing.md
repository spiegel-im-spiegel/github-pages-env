+++
description = "これはどちらが正しいかという問題ではない。"
date = "2017-03-11T14:55:06+09:00"
update = "2018-11-11T15:28:29+09:00"
title = "きみは Generics がとくいなフレンズなんだね，または「制約は構造を生む」"
draft = false
tags = ["golang", "object-oriented", "java", "design", "refactoring", "generics"]

[author]
  license = "by-sa"
  flickr = "spiegel"
  flattr = "spiegel"
  avatar = "/images/avatar.jpg"
  instagram = "spiegel_2007"
  tumblr = "spiegel-im-spiegel"
  url = "http://www.baldanders.info/spiegel/profile/"
  name = "Spiegel"
  linkedin = "spiegelimspiegel"
  facebook = "spiegel.im.spiegel"
  twitter = "spiegel_2007"
  github = "spiegel-im-spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

{{< fig-quote title="数学ガール／フェルマーの最終定理" link="http://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1CM/baldandersinf-22/" >}}
<q>公理によって与えられる暗黙の制約。この制約が集合の要素同士をしっかり結びつける。単純にしばるのではない、相互に秩序ある関係を結ぶ。言い換えれば――公理によって与えられる制約が構造を生み出しているのだ</q>
{{< /fig-quote >}}

今回は戯れ言モードなので「[プログラミング言語 Go](/golang/)」ではなくこちらで書いてみる。
コードは1行も書かないのでご安心を（笑）

- [Big Sky :: golang と Generics と私](http://mattn.kaoriya.net/software/lang/go/20170309201506.htm)
- [golang と Generics と吾 - Qiita](http://qiita.com/yuroyoro/items/6bf33f3cd4bb35469e0b)
- [Java の Generics にもの思い - Qiita](http://qiita.com/t2y/items/139c6a38173d7750ddfc)

私は出自が組込みエンジニアで（今は何でも屋），アセンブラや C/C++ から始まり Java などの制御に向いていると言われる言語を遍歴している（PHP を機器制御に使うとかいうこともやったが）。
[Go 言語]もその流れから興味を持っているが，あいにく私が住んでいる地方都市で [Go 言語]の出番はまだない。

そういう経歴を持つ私から見て [Go 言語]が特異だと思ったのは以下の2点である。

1. 例外処理がない
2. 明示的なクラス定義構文がない

私だけでなく C++ や Java などから来た人は大抵これで面食らうらしい。

このうち1番目については「[プログラミング言語 Go](/golang/)」で[記事にした]({{< ref "/golang/error-handling.md" >}} "エラー・ハンドリングについて")ので割愛する。

さて，2番目の「明示的なクラス定義構文がない」について。

そもそも「クラス」とはなにか。
クラスとは以下の要素をひとまとめの「モノ（object）」として定義したものである。

- 名前（必ず1個）
- 属性（0個以上）
- 操作（0個以上）

[Go 言語]では明示的なクラス定義構文がない代わりに [type] と [struct]，およびメソッド・レシーバを組み合わせることでクラスの要素である名前，属性，操作を定義できる。

そしてクラス定義で重要なのは「クラス間の関係」を定義することである。
クラス間の関係としては大雑把に以下の2つがある。

1. 汎化・特化（継承 等）
2. 関連（集約，依存 等）

このうち2番目の関連は定義しやすい。
あるクラスの属性として別のクラスを定義するか，操作によって関連付けるかすればいいからだ。
問題は1番目の汎化・特化をどうやって定義するかである。

[Go 言語]の場合は [interface] を使った「構造的部分型（structural subtyping）[^dt]」を採用した。
構造的部分型とはクラスの振る舞いに注目してクラス間の汎化・特化関係を帰納法的に定義することである。
例を挙げると，それが「にゃーん」と鳴くのなら机器猫だろうが猫耳メイドだろうがサーバルキャットだろうが全部「猫」である，ということだ。

[^dt]: URL を見るとわかる通り最初は [duck typing] と表記していたが正しくは「構造的部分型」と言うらしい。 [duck typing] は主に動的型付け言語における型推論方式（のひとつ）で，クラス間の関係を記述するものではないようだ。ちなみに構造的部分型に対する言葉として「公称型（nominal subtyping）」というのがあって， C++ や Java におけるテンプレート・クラスやインタフェース・クラスを使った汎化・特化関係を指す場合に使うそうだ。

クラス間の関係を定義するのは意外に大変である。
皆さんは「クラス設計」をどのように行っているだろうか。
まずは具体的なクラスを列挙していき，それらの関係を考察していくのではないだろうか（「ユーザ」や「管理者」を定義するのに 動物→人間→... と考えていく人はいないだろう）。
考察する過程で（クラスとクラスを繋ぐ）不可視のクラスを発見したり複数のクラスがひとつの概念で括れることに気づいたりすることもある。
つまり設計する過程では「具象→抽象」へと遡っていく。

一方，実装する際には， C++ や Java では最初にテンプレート・クラスやインタフェース・クラスを作ってからインプリメント・クラスに落とし込む。

たとえば，最初に「猫」という抽象クラスを作っておいて，それを継承する形で机器猫や猫耳メイドやサーバルキャットといった具体的なクラスを実装していく。
つまり「抽象→具象」へと作業していくわけだ。
そしてその過程において Generics[^g1] は，ほとんど必須と言えるほど利用価値の高い機能と言える。

[^g1]: 知らない人のために Generics について簡単に説明しておくと，変数の型あるいはインスタンス（instance）に対するクラス（class）に関係なく単一の記述で変数ないしインスタンスを扱うことのできる仕組みである。汎化の一種と考えてもよい。いわゆる多態性（polymorphism）とは異なり，継承関係の異なるクラスでも一緒くたに扱うことが可能なかなり強力な仕組みである。 Generics は特にコンテナ（container; オブジェクトの集まりを表現するデータ構造，配列など）操作で威力を発揮する。

これが [Go 言語]による実装ではひっくり返る。
たとえば，最初に机器猫や猫耳メイドやサーバルキャットといった具体的なクラスを作っていって「これってみんな『にゃーん』って鳴くじゃん」と気がつけば後付けで「猫」という抽象クラスを実装できるのである。

どういうことかというと， [Go 言語]においては設計と実装を同時進行で「具象→抽象」へと考察していくことができる，ということである。
このような思考過程においては Generics の有無はさして重要ではなくなる。
だって具象化されたオブジェクトから作り始めるのだから。

「抽象→具象」へと実装する人にとっては Generics のない [Go 言語]はとてもまだるこしく見えるかもしれない。
「なんで Generics がねーんだよ。いちいち全部書かせる気か。このポンコツ言語が！」となること請け合いである。
しかし一度構造的部分型に慣れた人にとっては抽象クラスから書かなければならない C++ や Java こそが面倒くさい。
何故なら，脳内では「具象→抽象」で思考していくのに実際に書くときには「考え終わらないと書けない[^cd1]」からである。
[Go 言語]なら「考えながら書ける」のに。

[^cd1]: 私はこれを「写経」と呼んでいる。はっきり言ってプログラミングでもっとも苦痛なのがコーディング＝写経だったりする。ちなみに一番好きなのはデバッグ。特に他人の書いたコードをデバッグするのは大好物。あれは極上の数理パズルである（締切さえなければね）。

これはどちらが正しいかという問題ではない。

たとえばウォータフォール型[^wf] の開発スタイルでは実装を開始するまでに設計が終わることが（建前上は）保証されているため「抽象→具象」へと書き進めることが容易な言語が向いている。
一方，要件が絶えず変わったり実験的な製品の場合は設計が終わるまで待っていられないため [Go 言語]のような言語が向いてるかもしれない。
まぁ設計と実装を同時にやろうとするとリファクタリングが頻繁に発生するのでコピペ・プログラマにはキツい作業になるかもしれないが。

[^wf]: 「ウォータフォール型」とは滝の水が上から下へと落ちていくように 要件定義→設計→製造 と上流工程から下流工程へ順番にプロセスを進めていく開発スタイル。工程ごとにマイルストーンを設けてチェックを行い，各工程が完了しないと先に進めないようにする。まぁ実際にはスケジュールやらの関係でチェックを端折って先に進めてしまうことが多く，下流工程に入ってから致命的な欠陥に気づいて抜き差しならない状況に陥ることもしばしばある（笑）

個人的には「プログラマは要件定義の段階から参加してコードを書くべき」と思ってるので，これを容易にするであろう [Go 言語]には注目している。

## 2018-10-11 追記

“[Go 2 Draft Designs](https://blog.golang.org/go2draft)” において将来バージョンにおける Generics 対応の言及がある。

{{< fig-youtube id="6wIP3rO6On8" title="Go 2 Drafts Announcement - YouTube" >}}

- [Go 2のgenerics/contract簡易まとめ](https://qiita.com/lufia/items/242d25e8c93d88e22a2e)
- [次期 Go 言語で導入される（かもしれない）総称型について予習する]({{< ref "/golang/generics-in-go-2.md" >}})

## ブックマーク

- [Why Everyone Hates Go · npf.io](Why Everyone Hates Go · npf.io)
    - [[翻訳]なんでGoってみんなに嫌われてるの？ - Qiita](http://qiita.com/hirokidaichi/items/adccebb41f77eaa6132f)
- [Why Go? | Dave Cheney](https://dave.cheney.net/2017/03/20/why-go)
    - [[翻訳] Why Go? - Qiita](http://qiita.com/methane/items/b627f20457873a504638)
- [Go 言語における「オブジェクト」]({{< ref "/golang/object-oriented-programming.md" >}})
- [JavaScriptで継承を使わないプログラミングスタイル - JavaScript勉強会](http://jsstudy.hatenablog.com/entry/2017/03/29/214931) : オブジェクト指向設計について上手くまとめている
- [Go にジェネリクスがなくても構わない人たちに対する批判について - methaneのブログ](http://methane.hatenablog.jp/entry/2017/09/19/Go_%E3%81%AB%E3%82%B8%E3%82%A7%E3%83%8D%E3%83%AA%E3%82%AF%E3%82%B9%E3%81%8C%E3%81%AA%E3%81%8F%E3%81%A6%E3%82%82%E6%A7%8B%E3%82%8F%E3%81%AA%E3%81%84%E4%BA%BA%E3%81%9F%E3%81%A1%E3%81%AB%E5%AF%BE%E3%81%99)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[struct]: https://golang.org/ref/spec#Struct_types "Struct types"
[type]: https://golang.org/ref/spec#Properties_of_types_and_values "Properties of types and values"
[interface]: https://golang.org/ref/spec#Interface_types "Interface types"
[duck typing]: https://en.wikipedia.org/wiki/Duck_typing "Duck typing - Wikipedia"
[duck test]: https://en.wikipedia.org/wiki/Duck_test "Duck test - Wikipedia"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22"><img src="https://images-fe.ssl-images-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" height="160" alt="プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)"></a></div>
	<dl class="fn">
      <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
      <dd>Alan A.A. Donovan, Brian W. Kernighan</dd>
      <dd>丸善出版</dd>
	  <dd>評価&nbsp;<span class="fa-sm" style="color:goldenrod;">
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
      </span></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed">2018.10.19</abbr> (powered by <a href="https://dadadadone.com/amakuri/" >Amakuri</a>)</p>
</div>
