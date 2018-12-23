+++
title = "「情報弱者」を再々定義する"
date = "2018-05-29T23:01:50+09:00"
update = "2018-05-30T09:52:34+09:00"
description = "検閲や注目の搾取がもたらす「情報の非対称性」の影響を受ける人達こそ「情報弱者」と呼ぶにふさわしいし，その意味で私たちユーザの大半は「情報弱者」なのである。"
image = "/images/attention/kitten.jpg"
tags = [ "code", "internet", "censorship", "grigori" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
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

## 「フィルターバブル」再考

最近メディア記事などで今更ながら「フィルターバブル」という単語を目にするようになったが大いなる誤解があるような気がする。

たとえばこんな感じ。

{{< fig-quote title="Facebookのフェイクニュース対策を悲観する理由--根深いフィルターバブルの壁" link="https://japan.cnet.com/article/35119789/" >}}
<q>結局、人間というものは「フィルターバブル」の心地よさが大好きなのだ。そして、Facebookはフィルターバブルのために設計されたサービスとも言える。</q>
{{< /fig-quote >}}

確かに Facebook や Google 検索サービス等はヒトの認知バイアス（cognitive bias）を巧妙に使っているが「人は信じたいものを信じる[^cb1]」ことと「Facebook の設計がフィルターバブルを生む」ことは別の話である。

[^cb1]: 確証バイアス（confirmation bias）の一例。確証バイアスを含む認知バイアスは動物が外部の情報を処理する際の適応反応（または行動）であり意識的に制御するのは難しい。たとえば「それ」が確証バイアスであると知り回避しようとする行動自体が別の認知バイアスを生む。科学・技術分野では認知バイアスによる偏りや誤りを多様な「目」による査読（peer review）によって解消しようとする。

Eli Pariser 氏による “THE FILTER BUBBLE” については私も（随分と遅ればせながらではあったが）[感想文を書いた]({{< ref "/remark/2016/05/filter-bubble.md" >}} "『フィルターバブル』を読む")[^fb0]。
日本では「閉じこもるインターネット」などというセンスの欠片もない邦題[^fb1] のせいでフィルターバブルが「情弱の SNS 引きこもり」みたいに思われている向きもあるが，向こうの記事（の翻訳）でも同じように受け止められているのは少々驚きである。

[^fb0]: [邦訳本](http://www.amazon.co.jp/exec/obidos/ASIN/4150504598/baldandersinf-22/ "フィルターバブル──インターネットが隠していること")のほうね。
[^fb1]: このタイトルは後の[文庫版](http://www.amazon.co.jp/exec/obidos/ASIN/4150504598/baldandersinf-22/ "フィルターバブル──インターネットが隠していること")でほぼ原書タイトルに近い名前に改題された。「閉じこもるインターネット」と付けた人（翻訳者ではないだろう）はきっと中身をちゃんと読んでなかったのだろう。

[私が書いた感想文]({{< ref "/remark/2016/05/filter-bubble.md" >}} "『フィルターバブル』を読む")でまとめた内容を以下に再掲載してみる。

1. ユーザによって制御できないフィルタリングは目的や運営主体にかかわらず全て「検閲」である。「フィルターバブル」は検閲の問題といえる
2. 「フィルターバブル」は都合の悪いものを隠し，受け手にとって都合のよいまたは心地よいものだけを「連呼」する。「フィルターバブル」は「注目の搾取」の問題といえる
3. 「フィルターバブル」は人の認知バイアスを利用し背景に擬態している。「フィルターバブル」は目に見えない壁であり手が届く範囲のほんのちょっと先にあるため存在することすら気付かない場合がある
4. 「フィルターバブル」は「アーキテクチャ[^ac1]」によって駆動するため一見すると公正に見える。「フィルターバブル」について考えることは「アーキテクチャ」の本質的な問題を炙りだす

[^ac1]: ここで言う「アーキテクチャ」とは新シカゴ学派の理論大系で出てくる「4つの規制」のひとつである。ちなみに「4つの規制」とは法（Law）・規範（Norm）・市場（Market）・アーキテクチャ（Architecture）の関係を指す。拙文「[『法のデザイン』を斜め読みした]({{< ref "/remark/2017/05/legal-design-book.md" >}})」でちょろんと紹介しているので宜しければどうぞ。

重要なのは最初の2つ。
つまりフィルターバブルは，ユーザ（あるいは消費者）側の問題ではなく，検閲の問題であり注目の搾取の問題である。
もう少し細かく言うなら「技術的ゲートキーパーによる検閲と搾取」なのである。

### 技術的ゲートキーパーによる検閲と搾取

「技術的ゲートキーパー（technological gatekeepers）」については Jonathan ZITTRAIN 氏の以下の論文が参考になるだろう。

- {{< pdf-file title="オンライン上のゲートキーピングの歴史 (1)" link="http://www.juris.hokudai.ac.jp/gcoe/journal/IP_vol28/28_4.pdf" >}}
- {{< pdf-file title="オンライン上のゲートキーピングの歴史 (2)" link="http://lex.juris.hokudai.ac.jp/gcoe/journal/IP_vol29/29_4.pdf" >}}
- {{< pdf-file title="オンライン上のゲートキーピングの歴史 (3・完)" link="http://lex.juris.hokudai.ac.jp/gcoe/journal/IP_vol30/30_5.pdf" >}}

この論文によると，（オンライン上の）ゲートキーパーとは「他者のコンテンツを伝達し、ホスティングし、インデックス化するなど、さまざまな形で媒介を行う者」を指し，特に「技術的ゲートキーパー」については「個人を特定したり個人の行動を規制したりすることを促すように技術自体を変化させる試み」と定義している。

「技術的ゲートキーパー」で括るのは，対象となるのが Facebook や Google といった特定のサービスだけではなく，あらゆるキャリア・ISP・OSP（Online Service Provider）に及ぶからである。
もちろん「メディア」も例外ではない。

確かに Facebook は [Campbridge Analytica の件]({{< ref "/remark/2018/03/name-identification.md" >}} "善悪の葛藤")以来さまざまな点で批判の矢面に立っているが Facebook に論点を絞るのは問題の矮小化といえる。
もしメディアが意図的に矮小化を行っているのなら，これこそがフィルターバブルの典型例と言えるだろう[^fn1]。

[^fn1]: いわゆる fake news というのは「矮小化」「先鋭化」「すり替え」といったもの（これも注目の搾取）の延長線上にあるものであり「メディア」が，それこそ新聞の時代から，根源的に抱える宿業のようなものだ。故に「fake news をなくす」というアプローチには殆ど意味がない。そのアプローチ自体が “fake” と言える。

## 「情報弱者」を再々定義する

もともと「情報弱者」という言葉は，情報へのアクセスについて世代・貧富・地理等の要因によって狭められる個人，すなわち「情報格差（digital divide）」を指す言葉だった。
ただし，ここ10年程は「情弱」という略語とともに情報の収集・分析の能力が劣る人あるいは単なる無知を指す差別語として再定義されている。
最初に引用した翻訳記事も「情弱」を揶揄する内容に見える。

しかし（これまで述べたように，あるいは[激化するサービス・ブロッキング]({{< ref "/remark/2018/04/japanese-gfw.md" >}} "日本版「グレート・ファイアウォール」に関するブックマーク集")等に見るように）そもそもネット上の情報は技術的ゲートキーパーとそれを利用するユーザとの間で非対称になっており，個人の格差や能力差以前の問題と言える。

なんでインターネットがこんなに「後退」してしまったのかについては別の議論に譲るとして，検閲や注目の搾取がもたらす「情報の非対称性」の影響を受ける人達こそ「情報弱者」と呼ぶにふさわしいし，その意味で私たちユーザの大半は「情報弱者」なのである。

この観点で言えば「fake news を検閲する」というのは完全なる悪手であり，私たち「情報弱者」を掬い上げることにはならない。
また [GDPR] 関連で再び注目されている「忘れられる権利」についても「情報の非対称性」という観点で見ればまた違った議論があるかもしれない。

## ブックマーク

- [フェイクニュースの需要と供給 | mhatta's mumbo jumbo](http://www.mhatta.org/wp/blog/2018/03/08/supply-and-demand-of-fakenews/)
    - [民主主義のその先へ（1） | mhatta's mumbo jumbo](http://www.mhatta.org/wp/blog/2018/03/15/beyond-democracy-1/)

- [監視をコントロールする — Baldanders.info](http://www.baldanders.info/spiegel/log2/000490.shtml)

[GDPR]: https://en.wikipedia.org/wiki/General_Data_Protection_Regulation "General Data Protection Regulation - Wikipedia"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4150504598/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/41UdjkE4OpL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4150504598/baldandersinf-22/">フィルターバブル──インターネットが隠していること (ハヤカワ文庫NF)</a></dt><dd>イーライ・パリサー 井口耕二 </dd><dd>早川書房 2016-03-09</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4569762468/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4569762468.09._SCTHUMBZZZ_.jpg"  alt="インターネット的 (PHP文庫)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4140912073/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4140912073.09._SCTHUMBZZZ_.jpg"  alt="ウェブ社会のゆくえ―<多孔化>した現実のなかで (NHKブックス No.1207)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4122033985/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4122033985.09._SCTHUMBZZZ_.jpg"  alt="情報の文明学 (中公文庫)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4480062858/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4480062858.09._SCTHUMBZZZ_.jpg"  alt="ウェブ進化論 本当の大変化はこれから始まる (ちくま新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4152096098/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4152096098.09._SCTHUMBZZZ_.jpg"  alt="五〇億年の孤独:宇宙に生命を探す天文学者たち"  /></a> </p>
<p class="description">ネットにおいて私たちは自由ではなく，それと知らず「フィルターバブル」に捕らわれている。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-05-07">2016-05-07</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
