+++
title = "「自動販売機の気持ちになって考える」"
date =  "2017-11-24T12:07:09+09:00"
description = "機械は人間のようには考えない。「機械（＝コンピュータ）が考える」というのはどういうことなのか。これを過不足なく説明するのは意外に難しい。"
image = "/images/attention/remark.jpg"
tags        = [ "engineering", "design", "programming" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

[結城浩さんの tweet](https://twitter.com/hyuki/status/933850105614032897) で

{{< div-gen >}}
<blockquote class="twitter-tweet" data-lang="ja"><p lang="ja" dir="ltr">プログラミングの課題で「自動販売機のシミュレータを作れ」というのはいい問題ではないかとよく思う。易しいものから、難しいものまで、仕様しだいでいろいろできる。GUIの有無、商品の種類、在庫管理、金種管理、各種アラート…</p>&mdash; 結城浩 (@hyuki) <a href="https://twitter.com/hyuki/status/933850105614032897?ref_src=twsrc%5Etfw">2017年11月24日</a></blockquote>
{{< /div-gen >}}

というのがあって，続きの tweets から見ても（プログラミングというよりは）システム設計（もしくは base design）のことを指しているのは明らかなんだけど，私が連想したのは梅津信幸さんの『[あなたはコンピュータを理解していますか？](https://www.amazon.co.jp/exec/obidos/ASIN/4774116009/baldandersinf-22/)』だった。

この本は[2002年に出版](https://baldanders.info/spiegel/log/nikki-s/200211.html#1706)されていて，これ自体は恐らく絶版なのだが，[2007年にリニュール本が出ている](https://www.amazon.co.jp/exec/obidos/ASIN/4797339497/baldandersinf-22/)ようだ。
さっそく発注してしまった（笑）

この本で一番印象に残っていたのが第3章の「自動販売機はコンピュータ理解の始まり」である。
この章では

1. 客がお金を入れる
1. 客が商品（ジュース1本）を注文する
1. 自動販売機が商品を出す
1. 自動販売機がお釣りを返す

という流れを人間がシミュレーションしていて，特に「機械がお金を計算する」方法について図解で説明されているのが秀逸だった。
つまり「自動販売機の気持ちになって考える」わけだ。

機械は人間のようには考えない。
「機械（＝コンピュータ）が考える」というのはどういうことなのか。
これを過不足なく説明するのは意外に難しい。
用語を並べ立てて煙に巻くか，頭の悪い擬人化でお茶を濁すのがせいぜいだろう。
私もちゃんと説明できるか自信がない。
でも，それを考えるのがプログラミングの最初の1フィートである。

どうせ学校教育でプログラミングを教えるのなら，「自動販売機の気持ちになって考える」からやっていただきたいものだ。
「プログラミング言語」なんか必要ないのである。
機械が人間の書いたプログラムを「理解」して実行する，というのは大いなる誤解である。
たとえば，そうした誤解が「[シンギュラリティ神話]({{< ref "/remark/2017/09/the-myth-of-the-singularity.md" >}} "『シンギュラリティの神話』を読む")」へ向かわせるのなら「プログラミング言語」を習うとかむしろ有害かもしれない。

学生の時に読んだ「日経サイエンス」の記事にティンカートイで3目並べを解くプログラムを組む話があったが[^ttc1]，どんなにコンピュータ技術が発達しても（量子コンピュータが登場しても）スタートラインは常に「ここ」なので，それを忘れないでほしいのですよ。

[^ttc1]: この話「三目並べをするティンカートイ・コンピュータ」は『[別冊日経サイエンス コンピューターレクリエーション 4 遊びの展開](https://www.amazon.co.jp/exec/obidos/ASIN/4532511135/baldandersinf-22/)』に収録されていたのだが，絶版本とは言え，どえらい値段がついてるな（笑）

まぁ，本気で AI が発達してノイマン型コンピュータが廃れたら説明の根本が変わるだろうけど，それは当分先の話（多分）。

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
