+++
date = "2016-11-03T22:18:00+09:00"
title = "鬼（もの）のインターネット"
description = "つまり，今のインターネットでは，結果的に，より多くの「モノ」を自分たちの陣営に「包摂」した者こそが「覇権」を握るのである。量こそ正義！ まさに民主主義の典型ではないか（笑）"
tags = [
  "internet",
  "engineering",
  "politics",
]
draft = false

[author]
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  flattr = "spiegel"
  flickr = "spiegel"
  facebook = "spiegel.im.spiegel"
  linkedin = "spiegelimspiegel"
  instagram = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
  avatar = "/images/avatar.jpg"
  license = "by-sa"
  github = "spiegel-im-spiegel"
+++

もともとの Internet of Things はいわゆる [Sensorware](http://www.yamdas.org/column/technique/sensor.html) の文脈で語られたもののようだが，2010年代に入って定義が変質する（よくある話）。

たとえば「[特定通信・放送開発事業実施円滑化法](http://law.e-gov.go.jp/htmldata/H02/H02HO035.html)」には以下の記述がある。

{{< fig-quote title="特定通信・放送開発事業実施円滑化法" link="http://law.e-gov.go.jp/htmldata/H02/H02HO035.html" >}}
<q>インターネット・オブ・シングスの実現（インターネットに多様かつ多数の物が接続され、及びそれらの物から送信され、又はそれらの物に送信される大量の情報の円滑な流通が国民生活及び経済活動の基盤となる社会の実現をいう。）</q>
（第五条 ２ 一）
{{< /fig-quote >}}

また経産省の「{{< pdf-file title="IoTセキュリティガイドライン ver1.0" link="http://www.meti.go.jp/press/2016/07/20160705002/20160705002-1.pdf" >}}」では ITU の定義を引いて

{{< fig-quote title="IoTセキュリティガイドライン ver1.0" link="http://www.meti.go.jp/press/2016/07/20160705002/20160705002-1.pdf" >}}
<q>IoT とは”Internet of Things”の略であり、ITU（国際電気通信連合）の勧告（ITU-T Y.2060(Y.4000)）では、「情報社会のために、既存もしくは開発中の相互運用可能な情報通信技術により、物理的もしくは仮想的なモノを接続し、高度なサービスを実現するグローバルインフラ」とされ</q>
（p.7，「1.3.1 IoT とは」）
{{< /fig-quote >}}

と記してある。

ポイントは「モノ（things）」自体は単なる PC (Programmable Controller) に過ぎず（intelligent ではない），その制御は主にインターネット側の「（intelligent な）誰か」が行う点にある。
問題は「モノ」を制御する「誰か」も programmable （というか hackable）であることで，更にその「モノ」が無数（それこそ何億という単位で）にネットに繋がってしまっていることだ。
まるで「蛍火の光く神」の如く（笑）

つまり，今のインターネットでは，手段の如何に関わらず，より多くの「モノ」を自分たちの陣営に「包摂」した者こそが「覇権」を握るのである。
量こそ正義！ まさに近代の夢，民主主義の典型ではないか（笑）

先日の DDoS のように，包摂された「モノ」が犯罪に使われれば確かにセキュリティ問題と言えるが，一方でこれが政治宣伝に利用されれば政治問題だし，軍事作戦に使われれば軍事問題である。
あるいはもっとつつましく家中のスマート家電を操って特定のお店でしか買い物をしないよう仕向ける，なんてなこともできるかもしれない。

まったく “[The Democratization of Censorship](https://krebsonsecurity.com/2016/09/the-democratization-of-censorship/)” とはよく言ったものである。

## ブックマーク

- [「モノのインターネット」--定義はどこまで拡散するのか - ZDNet Japan](http://japan.zdnet.com/article/35051376/)
- [Sensorware（前編）](http://www.yamdas.org/column/technique/sensor.html)
    - [Sensorware（後編）](http://www.yamdas.org/column/technique/sensor2.html)
    - [Sensorwareふたたび | ワイアードビジョン アーカイブ](http://archive.wiredvision.co.jp/blog/yomoyomo/200905/200905141600.html)
- [エフセキュアブログ : スマートホームの安全を保つ方法](http://blog.f-secure.jp/archives/50744439.html)
- [モノのインターネットはすさまじく危険だ - そして多くはパッチ不可能である by Bruce Schneier (The Internet of Things Is Wildly Insecure - And Often Unpatchable)](http://www.unixuser.org/~euske/doc/tiotiwiaou/index.html)
- [IoTデバイスがハッキングされやすい簡単な理由 | Business Security Insider Japan](https://jp.business.f-secure.com/the-simple-reason-iot-devices-are-so-hackable/)
- [我々は「モノのインターネット」からインターネットを守る必要がある、ところまで来てしまったのか - YAMDAS現更新履歴](http://d.hatena.ne.jp/yomoyomo/20161030/iotsecurity) ： この記事のリンク先も必見

## 参考図書？

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/456008193X/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/315iNBEKHLL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/456008193X/baldandersinf-22/">鬼と天皇（新装版）</a></dt><dd>大和 岩雄 </dd><dd>白水社 2012-01-18</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4584393788/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4584393788.09._SCTHUMBZZZ_.jpg"  alt="もうひとつの日本史 闇の修験道 異端の古代史5 (ワニ文庫)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4309226159/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4309226159.09._SCTHUMBZZZ_.jpg"  alt="諏訪の神: 封印された縄文の血祭り"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4908117039/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4908117039.09._SCTHUMBZZZ_.jpg"  alt="天皇と鬼"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/499065692X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/499065692X.09._SCTHUMBZZZ_.jpg"  alt="失われた十部族の足跡　イスラエルの地から日本まで　－新書版－"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4480083774/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4480083774.09._SCTHUMBZZZ_.jpg"  alt="雨月物語 (ちくま学芸文庫)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4054061605/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4054061605.09._SCTHUMBZZZ_.jpg"  alt="日本とユダヤ 聖徳太子の謎 (ムー・スーパー・ミステリー・ブックス)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4884698207/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4884698207.09._SCTHUMBZZZ_.jpg"  alt="富士山、2200年の秘密 なぜ日本最大の霊山は古事記に無視されたのか"  /></a> </p>
<p class="description">「鬼は天皇の影法師であり，両者の関係は「かくれんぼう遊び」に喩えることができる」（当時の帯より）</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-11-03">2016-11-03</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
