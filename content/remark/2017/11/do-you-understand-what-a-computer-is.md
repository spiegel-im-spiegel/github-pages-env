+++
title = "「自動販売機の気持ちになって考える」"
date =  "2017-11-24T12:07:09+09:00"
description = "機械は人間のようには考えない。「機械（＝コンピュータ）が考える」というのはどういうことなのか。これを過不足なく説明するのは意外に難しい。"
image = "/images/attention/remark.jpg"
tags        = [ "engineering", "design", "programming" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[結城浩さんの tweet](https://twitter.com/hyuki/status/933850105614032897) で

{{< div-gen >}}
<blockquote class="twitter-tweet" data-lang="ja"><p lang="ja" dir="ltr">プログラミングの課題で「自動販売機のシミュレータを作れ」というのはいい問題ではないかとよく思う。易しいものから、難しいものまで、仕様しだいでいろいろできる。GUIの有無、商品の種類、在庫管理、金種管理、各種アラート…</p>&mdash; 結城浩 (@hyuki) <a href="https://twitter.com/hyuki/status/933850105614032897?ref_src=twsrc%5Etfw">2017年11月24日</a></blockquote>
{{< /div-gen >}}

というのがあって，続きの tweets から見ても（プログラミングというよりは）システム設計（もしくは base design）のことを指しているのは明らかなんだけど，私が連想したのは梅津信幸さんの『[あなたはコンピュータを理解していますか？](http://www.amazon.co.jp/exec/obidos/ASIN/4774116009/baldandersinf-22/)』だった。

この本は[2002年に出版](http://www.baldanders.info/spiegel/log/nikki-s/200211.html#1706)されていて，これ自体は恐らく絶版なのだが，[2007年にリニュール本が出ている](http://www.amazon.co.jp/exec/obidos/ASIN/4797339497/baldandersinf-22/)ようだ。
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

[^ttc1]: この話「三目並べをするティンカートイ・コンピュータ」は『[別冊日経サイエンス コンピューターレクリエーション 4 遊びの展開](http://www.amazon.co.jp/exec/obidos/ASIN/4532511135/baldandersinf-22/)』に収録されていたのだが，絶版本とは言え，どえらい値段がついてるな（笑）

まぁ，本気で AI が発達してノイマン型コンピュータが廃れたら説明の根本が変わるだろうけど，それは当分先の話（多分）。

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4797339497/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/51W3fP3Q%2BtL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4797339497/baldandersinf-22/">あなたはコンピュータを理解していますか? 10年後、20年後まで必ず役立つ根っこの部分がきっちりわかる！ (サイエンス･アイ新書)</a></dt><dd>梅津 信幸 </dd><dd>ソフトバンク クリエイティブ 2007-03-16</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797354690/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797354690.09._SCTHUMBZZZ_.jpg"  alt="あなたはネットワークを理解していますか? インターネット時代に欠かせない根っこの知識が確実に身につく! (サイエンス・アイ新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4087474283/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4087474283.09._SCTHUMBZZZ_.jpg"  alt="痛快!コンピュータ学 (集英社文庫)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774124222/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774124222.09._SCTHUMBZZZ_.jpg"  alt="コンピュータのしくみを理解するための10章"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797348747/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797348747.09._SCTHUMBZZZ_.jpg"  alt="カラー図解でわかる通信のしくみ あなたはインターネット&モバイル通信をどこまで理解していますか? (サイエンス・アイ新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797370939/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797370939.09._SCTHUMBZZZ_.jpg"  alt="図解でかんたんアルゴリズム 情報処理のかなめとなる考え方が手に取るようにわかる! (サイエンス・アイ新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4822281655/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4822281655.09._SCTHUMBZZZ_.jpg"  alt="コンピュータはなぜ動くのか～知っておきたいハードウエア＆ソフトウエアの基礎知識～"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797384298/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797384298.09._SCTHUMBZZZ_.jpg"  alt="コンピューター&テクノロジー解体新書"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4822282708/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4822282708.09._SCTHUMBZZZ_.jpg"  alt="情報はなぜビットなのか 知っておきたいコンピュータと情報処理の基礎知識"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4816352481/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4816352481.09._SCTHUMBZZZ_.jpg"  alt="史上最強カラー図解 プロが教えるパソコンのすべてがわかる本"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4794220588/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4794220588.09._SCTHUMBZZZ_.jpg"  alt="文庫 思考する機械コンピュータ (草思社文庫)"  /></a> </p>
<p class="description">2002年に技術評論社から出た同名タイトルのリニューアルらしい。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-11-24">2017-11-24</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
