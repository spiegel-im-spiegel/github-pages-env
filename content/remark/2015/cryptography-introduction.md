+++
date = "2015-09-20T21:43:17+09:00"
update = "2016-01-07T13:09:45+09:00"
description = "今回は大幅改訂版なので，以前のを持ってる人も買っておいて損はない。"
draft = false
tags = ["book", "review", "cryptography", "security", "trust", "sha-3", "blockchain", "ecc"]
title = "『暗号技術入門 第3版』をななめ読み"

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

結城浩さんの『[暗号技術入門 第3版](http://www.hyuki.com/cr/)』がついに登場。
前の第2版のときは細々した追記が主だったような気がするが，今回は大幅改訂版なので，以前のを持ってる人も買っておいて損はない。
主な改訂内容としては

- SHA-3 について詳しく解説
- HeartBleed や POODLE など，最近の攻撃手法について言及
- 付録で楕円曲線暗号（Elliptic Curve Cryptography）について詳しく解説
- Bitcoin というか Bitcoin の中の重要な技術要素である Block Chain について詳しく解説

他にも [GnuPG](https://www.gnupg.org/) の記述が modern version に対応してたり，認証つき暗号（AEAD; Authenticated Encryption with Associated Data）およびその実装である GCM (Galois/Counter Mode) への言及があったり，いろいろ細かく手直しされている。

特に楕円曲線暗号の解説は秀逸で，入門レベルでの解説の中では一番分かりやすかった。
あと Block Chain の解説もお勧め。
Bitcoin や Block Chain に関する解説本はすでにたくさん出ているが，暗号技術の観点からきちんと解説しているものは少なく，「信用モデル」にまで話を展開しているものは更に少ない。

結局，暗号システムの実装というのは究極的には「信用モデル」の開発であると言える。
問題は「信用モデル」はロジックだけでは成立しない，ということだ。
『[暗号技術入門 第3版](http://www.hyuki.com/cr/)』では信用モデルの例として hierarchical PKI の典型である X.509 と OpenPGP の Web of Trust，そして Block Chain を挙げているが，それぞれ特徴があって比較すると面白い。
たとえば Block Chain はユーザ間に働く経済的 incentive を巧妙に使うが，それだけにパラメータの調整が難しいし， Mt. Gox のような経済リスクも考慮しなくてはならない。

そもそも信用というのは過去の事実に対してのみ評価可能なのに，実際に評価したいのは現在及び未来についてなのだ。
これって本来は不能解だよね。
でも信用が評価できなくては世の中は回らない。
だから，どうにかして実用可能なレベルにまで近似できないか，と専門家やエンジニアは頭を悩ますわけ。

そういったことを頭の隅に入れながら読めば，この本は単なるリファレンス本ではないことに気づくと思う。

最後にちょっとだけ注文をつけるなら「前方秘匿性（PFS; Perfect Forward Secrecy）」および OTR (Off-the-Record) Messaging の肝である「否認可能（Deniability）」についても言及が欲しかった[^1]。
メッセージングの世界においてはこのふたつが重要な要件になってきているからだ。

[^1]: PFS についてはもしかしたら見落としてるかもしれないが。なにせ斜め読みだったから。

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">第3版出た！ てか，もう Kindle 版出てるのか。紙の本買うのはやまったかなぁ。 SHA-3 や BitCoin/BlockChain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4757143044/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/413qoSjODUL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4757143044/baldandersinf-22/">信頼と裏切りの社会</a></dt><dd>ブルース・シュナイアー 山形 浩生 </dd><dd>エヌティティ出版 2013-12-24</dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4622078007/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4622078007.09._SCTHUMBZZZ_.jpg"  alt="殺人ザルはいかにして経済に目覚めたか?―― ヒトの進化からみた経済学"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4620323098/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4620323098.09._SCTHUMBZZZ_.jpg"  alt="リベラルのことは嫌いでも、リベラリズムは嫌いにならないでください－－井上達夫の法哲学入門"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4152094362/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4152094362.09._SCTHUMBZZZ_.jpg"  alt="楽観主義者の未来予測(上): テクノロジーの爆発的進化が世界を豊かにする"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/413052027X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/413052027X.09._SCTHUMBZZZ_.jpg"  alt="社会調査の考え方　下"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4326302402/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4326302402.09._SCTHUMBZZZ_.jpg"  alt="歴史から理論を創造する方法: 社会科学と歴史学を統合する"  /></a> </p>
<p class="description">ゴメン。途中まで読んで積読中。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
