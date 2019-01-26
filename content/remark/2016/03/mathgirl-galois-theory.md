+++
date = "2016-03-22T23:19:00+09:00"
description = "というわけで，以前に「ガロア理論」の中身がよく分からなかったって方は「数学ガールの秘密ノート」シリーズを読んでから再読してみると理解が進むかもしれない。"
draft = false
tags = ["book", "math", "galois-theory"]
title = "いまさら『数学ガール／ガロア理論』を読んだ時の話"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"
+++

[Qiita](http://qiita.com/) で面白い記事を見つける。

- [Swiftで代数学入門 〜 1. 数とは何か？ - Qiita](http://qiita.com/taketo1024/items/bd356c59dc0559ee9a0b)

[前にも書いた]({{< ref "/remark/2015/programming-language.md" >}})が，私の母国語はアセンブラとC言語である。
これは私の出自が「制御システム」だからである。
なので，組み込み系はもちろん業務アプリケーションでも「いかに対象を制御するか」を念頭に置いてプログラミングを行っている。
しかし「関数型言語」は「制御」に関してはどうにも具合が悪い。
いや，もちろん Erlang のような優れた言語もあるし，決して「関数型言語」が制御に向かないわけではないのだが，シャンタッ君並に小さい私の脳みそが「関数型言語」のコードをノイマン型コンピュータのインストラクションに上手く「コンパイル」してくれないのである。

でもようやく近年になって少し見方が変わってきて，「対象を制御する」のではなく「対象を記述する」手段としてプログラミングを考えるようになってきた。
そういう意味で「数学を記述する」この連載は教科書としてうってつけである。
Swift は Apple 製ということで Objective-C 並に毛嫌いしていたが（仕事ならやりますよ，もちろん），オープンソースになったことだし， iOS 以外にも使えるのならちょっと勉強してみようかと思っていたところなのだった。

そういえば「ガロア理論」って「数学ガール」シリーズでもやってたなぁ... と思って[昔のブログ記事](https://baldanders.info/archives.shtml)を探しまわったが，ない。
あれ？ 読書感想文を書いてないじゃん，私。
そういや2012年頃ってエラい仕事が忙しくて「ガロア理論」もすぐには読めなかったんだっけ。
やっちまったなぁ。

というわけで前置きが長くなったが，今更『数学ガール／ガロア理論』の感想文を書いてみる。
もう多少ネタバレしててもいいよね。

----

今回はやはりミルカさん回か。
いやユーリちゃんも今回はいい感じだったんだけどね。
終盤のミルカさんの「教師」っぷりには思わず「ミルカさんも大人になったなぁ」とホロリとしてしまった。
しかし，それよりもリサちゃん（「ちゃん」は不要）の出番が少ないです！ もっと出してください。

今回の話は「ガロア理論」だけあって「教師と生徒」を一方の軸として物語が進んでいく。
さらに主要キャラクター達が教師役と生徒役とで立ち位置を入れ替え（置換）ていくのが面白い。
また「教師と生徒」というのは著者の結城浩さんの言う「[対話](http://rentwi.textfile.org/?708541967391547392s)」形式のひとつだ。
こういう幾重にも重なった「協奏（concerto）」は音楽的でもある。

実は群論についてはあまり真面目に勉強しなかったんだよなぁ。
大学の講義でちょろんと習った覚えはあるけど，それだけ。
私にとって数学は単なる「道具」で，道具の中身を弄くり回すことについては興味が薄かった。

だから「あみだくじ」から話が始まって「これはラグランジュ分解式だな」と分かっても，そこから先にどう展開するか読めない。
でも逆にそれが私にとって面白かったようだ。
今回この記事を書くために再読したが，やっぱりワクワクしながら読めた。

改めて読んで気づいたんだけど，「数学ガールの秘密ノート」シリーズは「数学ガール」シリーズの「補題」になってるんだね。
とくに昨年出た「[ベクトルの真実]({{< ref "/remark/2016/03/mathgirl-note-vector.md" >}})」やさらにその前の「[丸い三角関数](https://baldanders.info/spiegel/log2/000685.shtml)」なんかは今回の「ガロア理論」と直接リンクしている。

というわけで，以前に「ガロア理論」の中身がよく分からなかったって方は「数学ガールの秘密ノート」シリーズを読んでから再読してみると理解が進むかもしれない。
そんで「ガロア理論」を読んだら改めて最初のリンク先の記事へ GO！

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%95%B0%E5%AD%A6%E3%82%AC%E3%83%BC%E3%83%AB%EF%BC%8F%E3%82%AC%E3%83%AD%E3%82%A2%E7%90%86%E8%AB%96-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B00L0PDMK4?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00L0PDMK4"><img src="https://images-fe.ssl-images-amazon.com/images/I/41szGJIR-jL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%95%B0%E5%AD%A6%E3%82%AC%E3%83%BC%E3%83%AB%EF%BC%8F%E3%82%AC%E3%83%AD%E3%82%A2%E7%90%86%E8%AB%96-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B00L0PDMK4?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00L0PDMK4">数学ガール／ガロア理論</a></dt>
	<dd>結城 浩</dd>
    <dd>SBクリエイティブ 2012-05-29 (Release 2014-07-24)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00L0PDMK4</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">あみだくじからガロア理論へ。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-03-22">2016-03-22</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
