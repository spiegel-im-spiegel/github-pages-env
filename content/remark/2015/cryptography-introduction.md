+++
date = "2015-09-20T21:43:17+09:00"
update = "2016-01-07T13:09:45+09:00"
description = "今回は大幅改訂版なので，以前のを持ってる人も買っておいて損はない。"
draft = false
tags = ["book", "cryptography", "security", "trust", "sha-3", "blockchain", "ecc"]
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
  url = "https://baldanders.info/spiegel/profile/"
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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE"><img src="https://images-fe.ssl-images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE">暗号技術入門 第3版　秘密の国のアリス</a></dt>
	<dd>結城 浩</dd>
    <dd>SBクリエイティブ 2015-08-25 (Release 2015-09-17)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B015643CPE</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E4%BF%A1%E9%A0%BC%E3%81%A8%E8%A3%8F%E5%88%87%E3%82%8A%E3%81%AE%E7%A4%BE%E4%BC%9A-%E3%83%96%E3%83%AB%E3%83%BC%E3%82%B9%E3%83%BB%E3%82%B7%E3%83%A5%E3%83%8A%E3%82%A4%E3%82%A2%E3%83%BC/dp/4757143044?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4757143044"><img src="https://images-fe.ssl-images-amazon.com/images/I/413qoSjODUL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E4%BF%A1%E9%A0%BC%E3%81%A8%E8%A3%8F%E5%88%87%E3%82%8A%E3%81%AE%E7%A4%BE%E4%BC%9A-%E3%83%96%E3%83%AB%E3%83%BC%E3%82%B9%E3%83%BB%E3%82%B7%E3%83%A5%E3%83%8A%E3%82%A4%E3%82%A2%E3%83%BC/dp/4757143044?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4757143044">信頼と裏切りの社会</a></dt>
	<dd>ブルース・シュナイアー</dd>
	<dd>山形 浩生 (翻訳)</dd>
    <dd>エヌティティ出版 2013-12-24</dd>
    <dd>Book 単行本（ソフトカバー）</dd>
    <dd>ASIN: 4757143044, EAN: 9784757143043</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">社会における「信頼」とは。そういいえば『<a href="https://www.amazon.co.jp/exec/obidos/ASIN/B01MZGVHOA/baldandersinf-22">超監視社会</a>』は積ん読のまま読んでない。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-31">2018-12-31</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
