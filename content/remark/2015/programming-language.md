+++
date = "2015-11-17T21:52:50+09:00"
update = "2016-02-18T12:56:03+09:00"
description = "こういう話は好きなので便乗してみる。"
draft = false
tags = ["programming", "language"]
title = "プログラミング言語との付き合い方"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

- [考えてみると結城はC, Perl, Javaの本は書いたけれど... - 結城浩の連ツイ](http://rentwi.textfile.org/?666213569055166464s)

今朝見かけたこれ。
こういう話は好きなので便乗してみる。

ちなみに，結城浩さんの通称「デザパタ本」はずい分昔に買っている。
お世話になってます。

## プログラミング言語の「母国語」

{{< fig-quote title="考えてみると結城はC, Perl, Javaの本は書いたけれど..." link="http://rentwi.textfile.org/?666213569055166464s" >}}
<q>プログラミング言語との付き合い方というのはいろいろあってですね。自分の母国語という言語はある。それから現在学んでいる最中の言語というのもある。そして、仕事用の言語やら、他の人とのコミュニケーション用言語というのもある。そのあたりは、自然言語とちょっと似ている。</q>
{{< /fig-quote >}}

個人的に母国語と言えるのはアセンブラとC言語。
私の場合はコードを脳内でインストラクションに翻訳する。
手続き型の言語ならこの「翻訳」をほとんど無意識でできる[^0]。
だから手続き型の言語であれば知らない言語でも見ればだいたい理解できる。

[^0]: たとえば，C言語では「ポインタ」の概念でよく躓くと言われているが，インストラクションで考えれば実にシンプル。

逆に，関数型のような非手続き型言語はあまり得意ではないのだが，簡単なものであれば手続き型に翻訳できるので，簡単なものを組み合わせて考えることで，まぁ何とか理解することはできる。

### ちまりまわるつ

そういえば，竹本泉さんの作品に『ちまりまわるつ』シリーズというのがある[^1]。
このシリーズの世界では「魔法」が科学的（？）に体系化されていて，工場で生産できるようになっている。なので，人力（？）で魔法を使う「魔法使い」は（工場で生産できない高度な魔法を使うことのできる一部の大魔法使いを除いては）古臭い職業として子供たちからは敬遠されている。
魔法使い達は能力でランク分けされていて，低ランクの魔法使いは簡単な呪文で簡単な魔法しか使えない。

でもここからが竹本泉作品らしいところで，簡単な呪文しか使えない魔法使いも呪文を組み合わせることで高度な魔法を使うことができるのだ（ただし高度な魔法は制御が難しいので，ふつうは使わせてもらえない）。
プログラムに例えるなら，高級言語で1行で書けるプログラムをアセンブラで数十ステップで書くような感じだと思えばよい。

こういう世界設定を少女漫画でさらっと描いてしまうところが竹本泉さんの凄いところである。

[^1]: 多分もう絶版。

### 目的別の言語

話が逸れた。

そういうわけなので個人的に「学んでいる最中の言語」というのはない。
見て理解できるかできないか。
いや，ぺーぺーの新人の頃はアセンブラやC言語を必死こいて学んでいたが，一度基礎ができればあとは全部「応用編」なのである。
そういう意味じゃ「今現在学んでない言語なんかない」とも言えるか[^a1] [^a2]。

[^a1]: 評価を中断している言語はある。 Erlang とか Scala とか Haskell とか。仕事で絡むようなことがあれば是非やりたい。
[^a2]: そう考えると，私にとって [Go 言語]はちょっと例外的かも。今のところ身近に [Go 言語]を使った仕事がないということもあるが，仕事抜きで完全に[「遊び」でやってる](/golang)。はっきり言って C/C++ や初期の Java 以来「これでなに作ろうかな」って思わせる言語に久しぶりに出会った感じ。もっとも，今は仕事が忙しすぎて全然かまってあげられないのだが orz

私は職業プログラマなので当然「仕事用の言語」というのは存在するが，「コミュニケーション用言語」というのはないな。
『数学ガール』冒頭のミルカさんの登場シーンはなかなかインパクトがあるが，あんな感じだろうか。

プログラマにとって最も信頼できる言葉は「動くコード」なので，ある意味で「仕事用の言語」が「コミュニケーション用言語」と言えるかもしれない。
ただ，職業プログラマは非プログラマとも話ができないといけない。
というか，大抵の顧客はそう。
顧客の「言語化されない意図」をいかに聞き取れるかが重要。
ホンマ「コミュニケーション用言語」なるものがあるなら欲しいよ。

あぁでも，[要求開発](http://www.baldanders.info/spiegel/log2/000869.shtml)で使う「概念モデリング」は「コミュニケーション用言語」と言えなくもない？

## 言語を巡る愛憎

{{< fig-quote title="考えてみると結城はC, Perl, Javaの本は書いたけれど..." link="http://rentwi.textfile.org/?666213569055166464s" >}}
<q>自然言語と同じようにプログラミング言語を使う人（要はプログラマ）には、その言語に対する愛情がこもる（愛憎がこもる）。なので、エンジニアリングや効率の話題と思っているのにいつのまにか忠誠心や貢献度みたいな話になることも。</q>
{{< /fig-quote >}}

仕事でプログラミングを行う際に言語を選べることはほとんどない。
顧客が指定してくることもあるし（顧客がコード資産を持っている），プロジェクト管理者が指定してくることもある（すでにある資産を使おうとする）。
そうじゃない場合でも要求と予算と期間とプロジェクトの面子によって（つまりそれを使うメリットとか関係なしに）言語が決まってしまう。

もし，そういうのが一切ないのなら，今なら [Go 言語]か JavaScript/[Electron] がいいなぁ。
これらで十分だよね。

一時期は Ruby も好きだったが，きれいさっぱり忘れてしまった。
今では Ruby の最新バージョンがいくつかさえ知らない。
なんか Rails 以降，凄い面倒くさいイメージがあるんだよね。
ある機能を Ruby で実装した記事を見かけたら，同じ機能を他の言語でもっと簡単にできないか，つい探してしまう。
こういうのが「宗教的」って言われるんだろうな，きっと（笑）

まぁ，でも，上で書いたように仕事で言語を選べることはほとんどないので，「グダグダ言わずにコード書け」って感じだけど。

**【12月21日 追記】**
フィードバックで Python について言及があったので。
Python は Linux 等では実質的に（少なくとも Ruby よりは）標準言語のようになっているし，資産も豊富なのでちゃんと覚えなきゃなぁ，とは思ってる。
しかし構文にインデントが必須な言語構造はどうしても慣れない。
同じ理由で Haskell や CoffeeScript とかも馴染めない。

プログラムコードをもっと human-understandable に，という考え方は分からなくもないけど，それなら [Go 言語]のようにフォーマッタやドキュメントフレームワークを標準で用意するほうが賢いと思う。
私はもうプログラミングで（タブだの全角空白だのも含めた）空白文字に煩わされたくない。

とはいえ，こういうのは「慣れ」の問題なので，仕事でやれと言われれば喜んでやりますよ。

## ひとつのプログラミング言語に縛られることの恐さ

{{< fig-quote title="考えてみると結城はC, Perl, Javaの本は書いたけれど..." link="http://rentwi.textfile.org/?666213569055166464s" >}}
<q>一つの技術に縛られることの恐さは、エンジニアなら誰でも知っている。では一つのプログラミング言語に縛られることの恐さは知っているか。一つのプログラミング言語がパーフェクトなことはない。時代が変われば要請も変わる。リソース配分は時々刻々変わる。そんな中で何にコミットするか。</q>
{{< /fig-quote >}}

手続き型言語の弱点は「ノイマン型コンピュータ」に最適化されていることだ。

ノイマン型コンピュータの構造はプロセッサとメモリが分かれているのが特徴で[^b]，メモリから命令を順番にプロセッサにフェッチして実行していく。

[^b]: キャッシュやパイプライン等の話はとりあえずチャイしてね。

もちろん「非ノイマン型コンピュータ」というのもある。
典型的なのは，いわゆるニューロチップである。

- [脳を模したチップ「TrueNorth」でコンピューティング革命を模索するIBM - ZDNet Japan](http://japan.zdnet.com/article/35073012/)

こういうタイプのコンピュータは（ノイマン型コンピュータで言うところの）プロセッサとメモリがセットでひとつの素子になっているのが特徴で，しかも素子同士はお互いに非同期で動く。
こういうのが本気で市場に台頭してきたら私のようなロートル・エンジニアはお払い箱だ。

## 若い人が正しい

{{< fig-quote title="考えてみると結城はC, Perl, Javaの本は書いたけれど..." link="http://rentwi.textfile.org/?666213569055166464s" >}}
<q>いつの時代でも、若い人は自由だ。何でも選べる。何でも言える。ときにはバカにされるかもしれない。わかってないよと侮られるかもしれない。でも、ほとんどの場合、若い人が正しい。</q>
{{< /fig-quote >}}

先のことは誰にもわからない。
特に今の時代は。
だから「より多くの未来を持っている人」が正しいと言える。
私のように人生の残り時間を勘定している人ではない。

願わくば「何でも選べる。何でも言える」未来を引き渡したいものである。

## 【おまけ】 TensorFlow に関するブックマーク

[フィードバック](#feedback)で [TensorFlow] について言及があったので，ついでに。

（この項は「[TensorFlow に関するブックマーク]({{< ref "/remark/2016/02/tensorflow.md" >}})」に移動しました）

[Go 言語]: https://golang.org/ "The Go Programming Language"
[Electron]: http://electron.atom.io/ "Electron"
[TensorFlow]: http://tensorflow.org/ "TensorFlow is an Open Source Software Library for Machine Intelligence"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B00EYXMA9I/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/41ETMZ7i9qL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B00EYXMA9I/baldandersinf-22/">数学ガール</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ株式会社 2014-02-14</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1CM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00I8AT1CM.09._SCTHUMBZZZ_.jpg"  alt="数学ガール／フェルマーの最終定理"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1D6/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00I8AT1D6.09._SCTHUMBZZZ_.jpg"  alt="数学ガール／ゲーデルの不完全性定理"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> </p>
<p class="description" >「数学ガール」シリーズ第1作目。ミルカさん衝撃の登場から分割数まで。</p>
<p class="gtools" >reviewed by <a href="#maker" class="reviewer">Spiegel</a> on <abbr class="dtreviewed" title="2015-06-06">2015/06/06</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html">G-Tools</a>)</p>
</div>
