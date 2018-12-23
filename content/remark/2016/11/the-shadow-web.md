+++
date = "2016-11-05T14:06:16+09:00"
title = "“The Shadow Web” （再掲載）"
description = "「インターネットは検閲をダメージであると解釈し，それを回避する」（John Gilmore）"
tags = [
  "book",
  "internet",
  "politics",
]
draft = false

[author]
  avatar = "/images/avatar.jpg"
  tumblr = ""
  flattr = ""
  instagram = "spiegel_2007"
  facebook = "spiegel.im.spiegel"
  url = "http://www.baldanders.info/spiegel/profile/"
  license = "by"
  github = "spiegel-im-spiegel"
  flickr = "spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  linkedin = "spiegelimspiegel"
+++

この記事は[本家サイト](http://www.baldanders.info/ "Baldanders.info")で2012年に公開した「[“The Shadow Web”](http://www.baldanders.info/spiegel/log2/000599.shtml)」を加筆・修正して再掲載したものです。
なお，本文中で紹介している記事「[介入されないもうひとつのウェブ]」は今年出た[別冊日経サイエンス『サイバーセキュリティ』](http://www.nikkei-science.com/page/sci_book/bessatu/51212.html)に再録されています。
興味のある方は是非どうぞ。

----

[『日経サイエンス』2012年6月号]の「[介入されないもうひとつのウェブ]」が面白そうだったので買ってみた。
原題は “The Shadow Web” （SCIENTIFIC AMERICAN March 2012）で著者はサイエンスライターの Julian Dibbell 氏。

インターネットはもともと障害や（国家などによる）検閲に強いシステムとして開発されてきた。
「インターネットは検閲をダメージであると解釈し，それを回避する」（John Gilmore）のである。
しかし現状のインターネットはこのようには機能していない。
昨年（2011年）1月末。市民運動で揺れるエジプトではたった5本の電話で8000万人のインターネット・アクセスが遮断された。
中国ではもう長いこと「グレート・ファイアウォール」が稼働している。
日本でだって無線通信業者がユーザ端末に勝手に ID を付与したりユーザの行動を追跡したり特定のサービスをブロックしたりしている。

{{< fig-quote title="介入されないもうひとつのウェブ" link="http://www.nikkei-science.com/201206_074.html" >}}
<q>もし今日のインターネットが実際にこの理論に近い状態であれば，メッシュネットワークは余計ものだったろう。
だがインターネットが当初の学術目的から踏み出して現在のような誰でも使える商業サービスになってから20年以上が経つうちに，そうした蓄積伝送の原理が果たす役割は，一貫して縮小していった。<br />
　この間，ネットワークに加わる新たなノードの圧倒的多数はISPを介してネットに接続する家庭や企業のコンピューターだった。
ISPの接続モデルでは，利用者のコンピューターはデータの中継はしない。
それはネットワークの端末，つまりデータの送受信だけを，常にISPのコンピューターを介して行うターミナル・ノードだ。
言い換えれば，インターネットの爆発的な成長はネットワーク地図に行き止まりのルートを増やしただけで，新たなルートを加えることはほとんどなかった。<br />
　そしてISPなど大量の情報ルートを持つ者は，彼らがルートを提供している何百万ものノードを支配下におくこととなった。
これらのノードは，もしISPがダウンしたり，ネットから遮断されたりすると，その障害を回避する方法がない。
ISPはインターネットが停止しないようにするどころか，実効上は停止スイッチになってしまった。</q>
（p.77）
{{< /fig-quote >}}

こうした問題に対処するために「無線メッシュネットワーク」の利用が進んでいる。
「無線メッシュネットワーク」ではユーザのノードは「ターミナル・ノード」ではなく中継ノードとして機能する。
このノードが増えれば増えるほどネットワークは自律的に機能し，遮断や検閲を回避できるようになる。

どうも日本ではメッシュネットワークはスマートグリッドの文脈で語られることが多いようである。
しかし，メッシュネットワークの本領は，少数の ISP による中央集権型ネットワークを補完する手段，まさに “The Shadow Web” として機能しうる点にあるように思える。

メッシュネットワークがうまくいくかどうかは参加するユーザの規模にかかっている。
メッシュネットワークが社会的に効果のあるレベルまで普及するには市場の15%を大幅に超える浸透率が必要になるそうだ。
しかし実際にそこまで行くかどうかは疑問らしい。
いっぽう為政者から見れば，現状でさえインターネットのコントロールを握ろうと躍起になってるのに，それをチャラにしてしまうような仕組みに対していい顔はしないだろう。
少なくともメッシュネットワークの構築について何らかの介入をしてくる可能性はある。
となれば市民運動レベルで普及を図っていくしかないということになるだろう。

{{< fig-quote title="介入されないもうひとつのウェブ" link="http://www.nikkei-science.com/201206_074.html" >}}
<q>人はとかく，プライバシーの死という環境破壊によってどれだけの被害をうけるかを過小評価してしまう。
通常の環境破壊行動，例えばゴミのポイ捨てや環境汚染が引き起こす，倍々に増えていく桁外れの被害を過小評価しがちなのとよく似ている</q>
（p.81）
{{< /fig-quote >}}

個人的には，日本で考えた場合，メッシュネットワークがどの程度効果的なのかよく分からない。
例えば災害時にネットワークを早く再構築するためにメッシュネットワークを使うというのはあるだろう（海外では既に事例がある）。
しかし “The Shadow Web” としてのメッシュネットワークがうまくいくのかどうかは疑問がある。
この辺は今後も情報を追いかけていきたいと思う。

[『日経サイエンス』2012年6月号]: http://www.nikkei-science.com/page/magazine/201206.html "2012年6月号 | 日経サイエンス"
[介入されないもうひとつのウェブ]: http://www.nikkei-science.com/201206_074.html "介入されないもうひとつのウェブ | 日経サイエンス"
[サイバーセキュリティ]: http://www.nikkei-science.com/page/sci_book/bessatu/51212.html "サイバーセキュリティー | 日経サイエンス"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4532512123/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51gurnOqhiL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4532512123/baldandersinf-22/">サイバーセキュリティ (別冊日経サイエンス)</a></dt><dd>日経サイエンス編集部 </dd><dd>日本経済新聞出版社 2016-04-22</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4320009061/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4320009061.09._SCTHUMBZZZ_.jpg"  alt="サイバーセキュリティ入門: 私たちを取り巻く光と闇 (共立スマートセレクション)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4905318416/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4905318416.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2016"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4822237656/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4822237656.09._SCTHUMBZZZ_.jpg"  alt="すべてわかるセキュリティ大全2017 (日経BPムック)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01G5SQLQC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B01G5SQLQC.09._SCTHUMBZZZ_.jpg"  alt="日経サイエンス2016年9月号"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4478083908/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4478083908.09._SCTHUMBZZZ_.jpg"  alt="CxO(経営層)のための情報セキュリティ―――経営判断に必要な知識と心得"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4822237982/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4822237982.09._SCTHUMBZZZ_.jpg"  alt="あなたのセキュリティ対応間違っています"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4339028533/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4339028533.09._SCTHUMBZZZ_.jpg"  alt="実践サイバーセキュリティモニタリング"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4802090927/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4802090927.09._SCTHUMBZZZ_.jpg"  alt="IT管理者のための情報セキュリティガイド (NextPublishing)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4822237788/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4822237788.09._SCTHUMBZZZ_.jpg"  alt="セキュリティ 最強の指南書(日経BPムック) (日経ITエンジニアスクール)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798145629/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798145629.09._SCTHUMBZZZ_.jpg"  alt="Webセキュリティ担当者のための脆弱性診断スタートガイド 上野宣が教える情報漏えいを防ぐ技術"  /></a> </p>
<p class="description">ここ5,6年ほどのセキュリティ・プライバシー関連の記事を集めたもの。俯瞰するにはちょうどいいと思う。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-11-05">2016-11-05</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
