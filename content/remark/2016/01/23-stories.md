+++
date = "2016-01-23T13:54:34+09:00"
update = "2016-01-23T21:34:36+09:00"
description = "Planet Nine / 次期X線国際天文衛星 ASTRO-H は 2月12日に打ち上げ予定 / 2^74207281-1 is Prime! / いまのところ「秀丸」への依存度は1割程度 / Bitcoin は失敗したか / TeX Wiki が移転しとるがな / 「『いま読むべき本』３冊」"
draft = false
tags = ["astronomy", "planet9", "jaxa", "hidemaru", "atom", "editor", "bitcoin", "tex", "pandoc", "book", "math", "prime-number"]
title = "週末スペシャル： Planet Nine"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

1. [Planet Nine]({{< relref "#planet9" >}})
1. [次期X線国際天文衛星 ASTRO-H は 2月12日に打ち上げ予定]({{< relref "#astro-h" >}})
1. [2^74207281-1 is Prime!]({{< relref "#prime" >}})
1. [いまのところ「秀丸」への依存度は1割程度]({{< relref "#hidemaru" >}})
1. [Bitcoin は失敗したか]({{< relref "#bitcoin" >}})
1. [TeX Wiki が移転しとるがな]({{< relref "#tex" >}})
1. [「『いま読むべき本』３冊」]({{< relref "#books" >}})

## Planet Nine{#planet9}

冥王星が9番目の惑星だった頃からこの手の話は尽きないが，どうも今回はマジらしい。

- [Caltech Researchers Find Evidence of a Real Ninth Planet | Caltech](http://www.caltech.edu/news/caltech-researchers-find-evidence-real-ninth-planet-49523)
- [EVIDENCE FOR A DISTANT GIANT PLANET IN THE SOLAR SYSTEM - pdf](http://iopscience.iop.org/article/10.3847/0004-6256/151/2/22/pdf)
- [Evidence grows for giant planet on fringes of Solar System : Nature News & Comment](http://www.nature.com/news/evidence-grows-for-giant-planet-on-fringes-of-solar-system-1.19182)
- [Planet Nine: Are We Not That Special? | SETI Institute](http://www.seti.org/seti-institute/planet-nine-are-we-not-that-special)
- [2016年1月21日ニュース「太陽系に9番目の惑星か 米チームが予測」 | SciencePortal](http://scienceportal.jst.go.jp/news/newsflash_review/newsflash/2016/01/20160121_02.html)
- [シミュレーションで推測、太陽系第9惑星存在の可能性 - アストロアーツ](http://www.astroarts.co.jp/news/2016/01/21planet9/index-j.shtml)
- [太陽系に「第9惑星」存在か　米天文学者らが論文発表 | sorae.jp : 宇宙（そら）へのポータルサイト](http://sorae.jp/030201/2016_01_21_planet-nine.html)
- [9番目の惑星の存在を示す証拠が発見された | TechCrunch Japan](http://jp.techcrunch.com/2016/01/21/20160120astronomers-find-evidence-of-a-ninth-planet/)

実際に発見されたわけではなく，いわゆる「海王星以遠天体（Trans-Neptunian Object; TNO または Edgeworth-Kuiper Belt Object; EKBO）」の軌道の偏りからの推測らしい。
よくある与太話ではなく，割と確からしい話のようだ。

こういう話を聞くといつも思うが，子供のころ聞かされていた「太陽系」のイメージが実は太陽のほんの近傍のものに過ぎないことに気付かされる。
「水金地火木...」などと受験用の念仏を唱えている場合ではないのですよ。

## 次期X線国際天文衛星 ASTRO-H は 2月12日に打ち上げ予定{#astro-h}

[ASTRO-H](http://astro-h.isas.jaxa.jp/) て広島大学の CORE-U（極限宇宙研究拠点）も参画してるのか。
母校のことなのに全然知らなかったよ。

- [X線天文衛星「ASTRO-H」2月12日打ち上げへ広島大学極限宇宙研究拠点 CORE-U](http://core-u.hiroshima-u.ac.jp/blog/2016/01/22/x%E7%B7%9A%E5%A4%A9%E6%96%87%E8%A1%9B%E6%98%9F%E3%80%8Castro-h%E3%80%8D2%E6%9C%8812%E6%97%A5%E6%89%93%E3%81%A1%E4%B8%8A%E3%81%92%E3%81%B8/)

## 2^74207281-1 is Prime!{#prime}

- [Mersenne Prime Discovery - 2^74207281-1 is Prime!](http://www.mersenne.org/primes/?press=M74207281)
- [ASCII.jp：「新たな世界最大の素数」昨年9月に発見](http://ascii.jp/elem/000/001/107/1107594/)
- [「史上最大の素数」、更新される « WIRED.jp](http://wired.jp/2016/01/22/discover-your-own-prime-number/)

[GIMPS (Great Internet Mersenne Prime Search)](http://www.mersenne.org/primes/) プロジェクトによる分散コンピューティングを使った素数探索で新たなメルセンヌ素数 $2^{74,207,281}-1$ が見つかったようだ。

「メルセンヌ素数」というのは $M_p = 2^p-1$ で表されるメルセンヌ数 $M_p$ のうち素数であるものを指す。
$M_p$ が素数なら $p$ も素数であるという面白い性質がある（逆は成り立たない）。
また $M_p = 2^p-1$ が素数なら $2^{p-1}(2^p-1)$ は完全数[^pn] になる。

[^pn]: 「完全数（perfect number）」とは「その数自身を除く約数の和がその数自身と等しい自然数」を指す。たとえば $6$ は $1+2+3 = 1\times2\times3$ で完全数になる。ちなみに $3$ は $2^2-1$ でメルセンヌ素数である。

メルセンヌ数に対しては効率的なな素数判定法が知られており分散コンピューティング向きの題材である。

### 参考

- [巨大素数と計算機](http://mailsrv.nara-edu.ac.jp/~asait/prime.htm)

## いまのところ「秀丸」への依存度は1割程度{#hidemaru}

- [WindowsがMacより優れているのは秀丸エディタが動くことだけ（暴論） - YAMDAS現更新履歴](http://d.hatena.ne.jp/yomoyomo/20160119/hidemaru)

あー，わかるわかる。
私もかなりの「[秀丸]依存症」なので（笑）

その「[秀丸]依存症」から脱却すべく [ATOM に乗り換え中]({{< relref "remark/2015/atom-editor.md" >}})だが，今のところ9割くらいは [ATOM] で賄えている。
とはいえ，私はコード書きで文章書きではない。
文章を大量に書く人に [ATOM] がいいかどうかはなんとも言えない。

[ATOM] の欠点は大きく2つ。

1. 起動が遅く，全体的にもっさりしている
2. 大きなテキストを読み込めない（または読み込みに非常に時間がかかる）

たかだか20万行程度のテキスト（ぶっちゃけ CSV なんだけど）を読み込むのにフリーズしないで欲しい。
起動が遅いのは仕方がないと諦めた。
秀丸みたいに常駐モードがあればいいんだけどねぇ。
まぁ [ATOM] を立ち上げっぱなしにしておけばいいか。

この辺は多分 [Electron](http://electron.atom.io/) の限界なんだろう。
そういや[秀丸]も登場したての頃は動作がもっさりしていて Vz と併用してたっけ（Vz は DOS 窓から使ってた）。
気長に性能向上を待つとしよう。

テキストエディタは製品云々より手に馴染むかどうかが絶対的に重要で，そのため必然的に hackable にならざるを得ない。
故に宗教論争も起こりやすい。
それに人は慣れ親しんだ手順や workflow からは簡単に抜け出せないものである。
私は emacs なんか触りたくもないし， vi/vim は若い頃のトラウマがフラッシュバックするので可能なかぎり使いたくない。

そういえば，結城浩さんは Vim を使われているような...

## Bitcoin は失敗したか{#bitcoin}

- [ビットコインは「失敗した」　離脱を表明した主要開発者が語る、その問題点 (1/2) - ITmedia ニュース](http://www.itmedia.co.jp/news/articles/1601/20/news120.html)

Bitcoin ってどうしてもかつての「地域通貨」を連想しちゃうんだよなぁ。
かつての「地域通貨」で上手くいったところはないはず。

通貨は血液のように国や企業や人々の間で循環していくことが重要。
特に総量が決まっている[^bc] 補完通貨は誰かがガメて抱え込んでしまえばそこでお終いなのだ。

[^bc]: たとえば Bitcoin の総量は約2,100万BTCで，どのように振り出されるかは数学的に決まっており恣意的な制御はできないようになっている。

そういえば「世界の富豪62人が保有する資産は、「36億人分の富」に相当する」のだそうだ。
さらに

{{< fig-quote title="世界人口の約半数は「より貧しく」なっていく « WIRED.jp" link="http://wired.jp/2016/01/21/global-poverty-oxfam/" >}}
<q>オックスファムによれば、世界の人口のうち「最も豊かな1パーセント」がもつ富と、「最も貧しい50パーセント」がもつ富が同じになるという見通しは、予想より1年早く、現実になったという</q>
{{< /fig-quote >}}

のだそうだ。
まぁ今の世界のシステムがそうなるよう構築されているのだから，ある意味当然の帰結と言えるけど。
「アベノミクス」だって結局はそのシステムに乗っかったものなんだから地方在住の貧乏人にカネが回ってくる道理がないのだ。

でも Bitcoin のような通貨システムはこの動きを加速させてしまうんじゃないだろうか。
どうなんだろう。
そろそろ本当に総括が必要なのだろうか。

### 参考

- [HotWired Japan_Altbiz  山形浩生の『ケイザイ2.0』 第23回 地域通貨って、そんなにいいの？](http://cruel.org/hotwired/hotwired23_01.html)
- [特集：FinTech入門（3）：ブロックチェーンは「取引コストゼロ」の世界を実現しようとしている (1/3) - ＠IT](http://www.atmarkit.co.jp/ait/articles/1601/21/news024.html)

## TeX Wiki が移転しとるがな{#tex}

1月20日から [https://texwiki.texjp.org/](https://texwiki.texjp.org/ "TeX Wiki") に移転したらしい。

つらつら眺めてたら

- [Atom - TeX Wiki](https://texwiki.texjp.org/?Atom)

というページがあった。

そういや [Pandoc] も気になってるんだよな。

- [Pandoc - About pandoc](http://pandoc.org/)
- [jgm/pandoc: Universal markup converter](https://github.com/jgm/pandoc)
- [Pandocに関する投稿 - Qiita](http://qiita.com/tags/Pandoc)
- [PandocでMarkdownからPDF化を試してみた - ARMERIA](http://betrue12.hateblo.jp/entry/2015/04/10/003830)

[Pandoc] は一度導入しかけて挫折したんだけど，最近のバージョンはとっつきやすくなってるみたい。
けど試す隙がない。

## 「『いま読むべき本』３冊」{#books}

- [出版の「初心」を思い出すための３冊 « マガジン航[kɔː]](http://magazine-k.jp/2016/01/22/3-books-for-new-editors/)

うーむ，「『いま読むべき本』３冊」かぁ。
私が人様に本を勧めるなどおこがましい話ではあるが，敢えて3冊選べというなら，今のところこれかな。

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4798115002/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/41bC8pdM2iL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4798115002/baldandersinf-22/">CODE VERSION 2.0</a></dt><dd>ローレンス・レッシグ Lawrence Lessig </dd><dd>翔泳社 2007-12-20</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798102040/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798102040.09._SCTHUMBZZZ_.jpg"  alt="コモンズ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798119806/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798119806.09._SCTHUMBZZZ_.jpg"  alt="REMIX ハイブリッド経済で栄える文化と商業のあり方"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798106801/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798106801.09._SCTHUMBZZZ_.jpg"  alt="Free Culture"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4480431837/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4480431837.09._SCTHUMBZZZ_.jpg"  alt="アーキテクチャの生態系: 情報環境はいかに設計されてきたか (ちくま文庫)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4757102453/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4757102453.09._SCTHUMBZZZ_.jpg"  alt="アーキテクチャの生態系――情報環境はいかに設計されてきたか"  /></a> </p>
<p class="description">前著『CODE』改訂版。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-15">2015-09-15</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4757143044/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/413qoSjODUL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4757143044/baldandersinf-22/">信頼と裏切りの社会</a></dt><dd>ブルース・シュナイアー 山形 浩生 </dd><dd>エヌティティ出版 2013-12-24</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4622078007/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4622078007.09._SCTHUMBZZZ_.jpg"  alt="殺人ザルはいかにして経済に目覚めたか?―― ヒトの進化からみた経済学"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4140816872/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4140816872.09._SCTHUMBZZZ_.jpg"  alt="限界費用ゼロ社会―<モノのインターネット>と共有型経済の台頭"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4478017883/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4478017883.09._SCTHUMBZZZ_.jpg"  alt="第五の権力---Googleには見えている未来"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4621089188/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4621089188.09._SCTHUMBZZZ_.jpg"  alt="リスク 不確実性の中での意思決定 (サイエンス・パレット)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/412102138X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/412102138X.09._SCTHUMBZZZ_.jpg"  alt="ソーシャル・キャピタル入門　- 孤立から絆へ (中公新書)"  /></a> </p>
<p class="description">社会における「信頼」とは。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-01-23">2016-01-23</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4903127044/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/417iD4x5N%2BL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4903127044/baldandersinf-22/">排除型社会―後期近代における犯罪・雇用・差異</a></dt><dd>ジョック ヤング Jock Young </dd><dd>洛北出版 2007-03</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4791764331/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4791764331.09._SCTHUMBZZZ_.jpg"  alt="後期近代の眩暈―排除から過剰包摂へ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4255008515/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4255008515.09._SCTHUMBZZZ_.jpg"  alt="断片的なものの社会学"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4796700439/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4796700439.09._SCTHUMBZZZ_.jpg"  alt="スティグマの社会学―烙印を押されたアイデンティティ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4791764242/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4791764242.09._SCTHUMBZZZ_.jpg"  alt="新しい貧困 労働消費主義ニュープア"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4062881357/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4062881357.09._SCTHUMBZZZ_.jpg"  alt="弱者の居場所がない社会――貧困・格差と社会的包摂 (講談社現代新書)"  /></a> </p>
<p class="description"><a href="http://www.baldanders.info/spiegel/log2/000410.shtml">感想はこちら</a>。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-15">2015-09-15</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

もっとも，ちゃんと最後まで読んだのは『排除型社会』だけであとは途中まで読んで積ん読状態。
まぁ『CODE』は Version 1 は既読なのでそのうちゆっくり。
面白い本ほど読むのに時間が掛かるし考える時間も増える。

というわけで，今年は「本を買わない」ことにした[^b]。
厳密には「2016年に新規に買うのは5冊まで」。
その代わり積ん読状態になってる本を消化することに注力する。

[^b]: マンガやラノベは別ね。あれは娯楽だから。

[秀丸]: http://hide.maruo.co.jp/software/hidemaru.html "秀まるおのホームページ(サイトー企画)－秀丸エディタ"
[ATOM]: https://atom.io/ "Atom"
[Pandoc]:http://pandoc.org/ "Pandoc - About pandoc"