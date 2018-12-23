+++
title = "「暗号通貨」ってゆーな！"
date =  "2018-01-10T22:21:55+09:00"
update = "2018-03-06T19:35:50+09:00"
description = "このまま進めば間違いなく「仮想通貨」は FinTech の中央集権化に向けた覇権競争に突入するだろう。"
image = "/images/attention/remark.jpg"
tags = ["blockchain", "engineering", "fintech", "market", "grigori", "hacker-ethic"]

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

断っておくが，私だけが言ってるんじゃないからね。

- [Cryptocurrencies Aren't 'Crypto' - Motherboard](https://motherboard.vice.com/en_us/article/43nk9b/cryptocurrency-are-not-crypto-bitcoin)
- ["Crypto" Is Being Redefined as Cryptocurrencies - Schneier on Security](https://www.schneier.com/blog/archives/2017/12/crypto_is_being.html)

{{< fig-quote title="Cryptocurrencies Aren't 'Crypto'" link="https://motherboard.vice.com/en_us/article/43nk9b/cryptocurrency-are-not-crypto-bitcoin" lang="en" >}}
<q>“Most cryptocurrency barely has anything to do with serious cryptography,” <a href="https://isi.jhu.edu/~mgreen/">Matthew Green</a>, a renowned computer scientist who studies cryptography, told me via email. “Aside from the trivial use of digital signatures and hash functions, it’s a stupid name.”</q>
{{< /fig-quote >}}

私の場合は「暗号」をキーワードにしている Google Alert がクソみたい記事ばかり引っ掛けるのにいい加減腹が立ってるからなのだが[^crypt1]。
なんだよ「暗号通貨長者」って。

[^crypt1]: ついでに言うなら，文章で禄を食んでる者なら「暗号」と「暗合」と「符号」と「符牒」をちゃんと書き分けて欲しい。なんでも「暗号」ゆーな！

ほんじゃあ，なんて言えばいいかというと

{{< fig-quote title="Cryptocurrencies Aren't 'Crypto'" link="https://motherboard.vice.com/en_us/article/43nk9b/cryptocurrency-are-not-crypto-bitcoin" lang="en" >}}
<q>So if you care about this, please politely correct people who incorrectly use the word “crypto.” Or maybe make fun of it, as Ryan Stortz, a security researcher in New York suggested. In a chat, he joked that he wants to start trolling people by referring to cryptocurrencies as “Block,” short for “blockchain technologies.”</q>
{{< /fig-quote >}}

日本語なら後ろに「通貨」を付けて「ブロックチェーン通貨」とか「ブロック通貨」といったところだろうか（笑） まぁこの記事では，大雑把すぎる名前ではあるが，無難に括弧書きで「仮想通貨」としておくか[^crypt2]。

[^crypt2]: “cryptocurrency” をわざと「仮想通貨」と訳してるページもあるみたい。

さて，今まで書いた Blockchain に関する記事は以下の通り

- [そろそろ Blockchain について勉強を始めるか — Baldanders.info](http://www.baldanders.info/spiegel/log2/000827.shtml)
- [Blockchain と Smart Contract]({{< ref "/remark/2016/01/blockchain-and-smart-contract.md" >}})

なのだが，2017年で大きく状況が変わったようなので再々勉強。

まずはこれまでのおさらいを再々掲載しておく。

1. Blockchain は「鎖」で繋がれた追記型データベース。「鎖」の途中のデータは取り消しも変更（改竄）もできない
2. Blockchain の追記プロセスには不正の余地がないよう何らかの仕掛けが必要。 Bitcoin の場合は「作業証明（proof-of-work）」がそれにあたる
3. Blockchain は P2P による分散型かつ fault tolerant（過失を許容する）なシステムだが最終的には fork も merge も許容しない
4. Bitcoin のアドレス（実体は公開鍵）の帰属先について Bitcoin/Blockchain は関知しない。Bitcoin が気にするのは Blockchain に記載されるアドレスの一貫性と無矛盾性である。アドレスの証明が必要な場合は外部の PKI を利用するか新たに組み込む必要がある

このうち大きく変わったのが 3 番目である。

## ICO と Hard Fork

私は最初，名前だけ見て東京オリンピックの話でもしてるのかと思ったよ（それは IOC だ`w`）。 

ICO (Initial Coin Offering) とは「仮想通貨」を使った資金調達を指すらしい。
建前上は「仮想通貨による生態系を創る」ということのようだが，実際には[厨二まるだしの趣意](https://www.shin-sasaki.com/single-post/2017/12/29/GACKT%25E3%2582%25B3%25E3%2582%25A4%25E3%2583%25B3%25EF%25BC%2588SPINDLE%25EF%25BC%2589%25E3%2581%25AE%25E4%25B8%25AD%25E8%25BA%25AB%25E3%2581%25AF%25E3%2581%259F%25E3%2581%25A0%25E3%2581%25AE%25E5%258E%25A8%25E4%25BA%258C%25E7%2597%2585 "GACKTコイン（SPINDLE）の中身はただの厨二病 | しんの公式ブログ")だったり[^gkt1]，挙句には資金を集めるだけ集めてトンズラする輩までいて，ぶっちゃけ山師の巣窟と成り果てている。
そもそも「仮想通貨による生態系を創る」という発想はゼロ年代に流行した「地域通貨」そのものであり，今もって[ミヒャエル・エンデの呪い](https://www.amazon.co.jp/exec/obidos/ASIN/B008YOHIAY/baldandersinf-22/ "エンデの遺言「根源からお金を問うこと」 | 河邑 厚徳, グループ現代 | ビジネス・経済 | Kindleストア | Amazon")は消えないのかと思うと，それはそれで怖ろしい気がする[^ende1]。

[^gkt1]: [memorandum@tumblr. の記事](http://pdl2h.tumblr.com/post/169105851262/)によると「恐ろしいのは、全ての広告は宣伝の機能のほかに、消費者をスクリーニングする機能があって、要するにその広告を「つまらない」とか「馬鹿げている」とか「興味が無い」と思う人は最初からその商材の想定顧客ではない、という点。つまり、騙されるのはこのホワイトペーパーを読んでもその間違いが理解できない程度の金融リテラシーを持つ人たち」ということらしく，ただの厨二文章ではないということですな。
[^ende1]: 念のために言うと私はミヒャエル・エンデ（の作品）の大ファンである。私のニックネームの Spiegel はエンデの『[鏡のなかの鏡](https://www.amazon.co.jp/exec/obidos/ASIN/4006020317/baldandersinf-22/)』から拝借している。

もっと決定的なのが Bitcoin の hard fork だ。
つまり Blockchain の分岐（fork）を許してしまったのだ。
Bitcoin における「信用」とは Blockchain に記述された取引履歴の一貫性と無矛盾性であり，本来はこの「信用」を「fork も merge も許容しない」ことで担保していたのが hard fork によって過去の「信用」を維持したまま別の通貨を生み出してしまったわけだ。
Bitcoin 単体は数学的に総量が決まっているが， hard fork すれば「仮想通貨」全体の量を**恣意的**に増やすことができる。

これは Bitcoin/Blockchain の根本を覆す。
何故なら Bitcoin の思想の根底にあるのは1990年代の PGP から脈々と続く「暗号アナーキズム」であり「支配者のコントロールから逃れる」ための脱中央集権化だったからだ。
しかし ICO や hard fork は守るべき一線を超えてしまった。

{{< fig-quote title="仮想通貨の2018年、熱狂に次ぐ幻滅の先に光明はあるか(楠正憲) - 個人 - Yahoo!ニュース" link="https://news.yahoo.co.jp/byline/kusunokimasanori/20180104-00080086/" >}}
<q>これまでビットコインの運営主体はいないとされてきたが、採掘事業者や開発のコアチームといった運営関係者が存在し、円滑な運営のためにはチームのガバナンスが必要で、仲間割れが起きた時にルールを強制する仕組みが何ら存在しないことも明らかとなった。仮想通貨も中央銀行と同じように、人間が創造し、人間が運用し、人間の意志によっていくらでも台無しにできるのだとすれば、仮想通貨だからといって発行に際して発行主体が負債を立てずに済む根拠をどこに置くのだろうか。<br>
野放図なビットコインの分裂に加えて、ICOの多くが有価証券規制を潜脱するため新たに発行するコインを仮想通貨と位置づけてしまったために、当初ビットコインが勝ち取った仮想通貨の特権的な法的位置づけは風前の灯にある。</q>
{{< /fig-quote >}}

## 「仮想通貨」という名の球根

これは Bitcoin の hard fork に深く関係するのだが，「Blockchain はスケールしない」ことが分かってきた。

おそらく Bitcoin/Blockchain は世界的な為替取引に晒されることは想像していなかったと思われる。
そもそも「作業証明」は計算集約的なものであり費用対効果の面から見た「上限」がある筈だった。
しかし Bitcoin を含めた「仮想通貨」が投機の対象となった時，ある筈だった「上限」を超えてしまった。

たとえば国家が「仮想通貨」を外貨獲得手段と見なせばコスト度外視で「作業証明」を行うことができる。
また malware や Web ページ上のスクリプト等を使って「作業証明」にかかるコストを不特定に転嫁する輩も出てきた。

こうして Bitcoin/Blockchain は限界を迎え， ICO や hard fork へと向かい，更なる投機が生まれるスパイラルになる。
「仮想通貨」はもはや通貨ではなく「チューリンプの球根」になってしまったわけだ。

{{< fig-quote title="暗号通貨バブルはイノベーションを絞め殺している | TechCrunch Japan" link="http://jp.techcrunch.com/2018/01/08/2018-01-07-the-cryptocurrency-bubble-is-strangling-innovation/" >}}
<q>現在、投機以外の暗号通貨トークン・プロジェクトは無期限の冬眠を強制されている。 本当にイノベーティブなサービスを作ろうとしているチームにとって、現在の暗号通貨バブルが弾けることは冬ではなく、むしろ春の到来を告げるものだ。</q>
{{< /fig-quote >}}

## 2018年以降の「仮想通貨」

ちょうど一年前（2017-01-10）には

- [未来のインターネットは、非中央集権型のインターネットだ  |  TechCrunch Japan](http://jp.techcrunch.com/2017/01/10/20170108the-future-is-a-decentralized-internet/)

なんてな記事もあって，あまりに皮肉なタイトルにバカウケしたものだが，もう笑い事ではなくなっているかもしれない。
何せあのザッカーバーグ君も ICO に興味を持ってるみたいだし。
ユーザを囲い込み個人情報の搾取を行う中央集権の権化たる Facebook のトップに立つ人間が

{{< fig-quote title="Mark Zuckerberg - Every year I take on a personal challenge to..." link="https://www.facebook.com/zuck/posts/10104380170714571" lang="en" >}}
<q>There are important counter-trends to this --like encryption and cryptocurrency -- that take power from centralized systems and put it back into people's hands. But they come with the risk of being harder to control. I'm interested to go deeper and study the positive and negative aspects of these technologies, and how best to use them in our services.</q>
{{< /fig-quote >}}

{{< fig-quote title="【全文和訳掲載】マーク・ザッカーバーグ「仮想通貨は権力を人々の手に戻す、フェイスブックでの活用方法探る」新年の抱負で言及 | News | Cointelegraph" link="https://jp.cointelegraph.com/news/zuckerberg-eyeing-out-power-of-cryptocurrencies" >}}
<q>このトレンドに逆流する重要なものがでてきている。例えば暗号化技術と仮想通貨だ。これらは中央集権化されたシステムから権力を奪い人々の手に戻す。だがこれらはコントロールしにくいというリスクを伴う。だからこれら技術のポジティブ・ネガティブ両側面とフェイスブックのサービスにどう活用できるかを深く研究したい。</q>
{{< /fig-quote >}}

とか恐怖しか感じない。
このまま進めば間違いなく「仮想通貨」は FinTech の中央集権化に向けた覇権競争に突入するだろう。
飼い慣らした筈の「ビヒーモス」が再び暴れだそうとしている。
あるいは「ビヒーモス」と「グリゴリ」が手を組んだか。

個人的には [Everipedia] 辺りが試金石になるのかな，と思っているのだが，どうだろうねぇ。

- [オンライン百科事典「Everipedia」がブロックチェーン導入で目指すもの｜WIRED.jp](https://wired.jp/2018/01/05/everipedia-blockchain/)

ダメかな，やっぱり。

## ブックマーク

- [Krugman: Baby Sitting the Economy (経済を子守りしてみると。)](http://cruel.org/krugman/babysitj.html)
- [グリゴリの捕縛 あるいは 情報時代の憲法について](http://orion.mt.tama.hosei.ac.jp/hideaki/kenporon.htm)
- [暗号通貨ブームの裏側で顕在化してきた、その基盤技術の「構造的な問題」｜WIRED.jp](https://wired.jp/2018/01/22/bitcoin-infrastructure/)
- [MIT Tech Review: 「非中央集権型」通貨ビットコインの理想は儚く消えたのか？](https://www.technologyreview.jp/s/71469/bitcoin-and-ethereum-have-a-hidden-power-structure-and-its-just-been-revealed/)
- [Blockchain Graveyard](https://magoo.github.io/Blockchain-Graveyard/)
- [仮想通貨とブロックチェーン、そしてICOの狂乱に思うこと：伊藤穰一｜WIRED.jp](https://wired.jp/2018/02/18/ico-cryptocurrency/)
- [MIT Tech Review: 暗号通貨は「通貨」ではなく「資産」、英中央銀総裁が語る](https://www.technologyreview.jp/nl/cryptocurrency-is-terrible-as-money-but-crypto-assets-are-for-real-says-bank-of-englands-chief/) : つまり「暗号」でも「通貨」でもない，と

[Everipedia]: https://everipedia.org/ "Everipedia, the encyclopedia of everything"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4314009071/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51ZRZ62WKCL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4314009071/baldandersinf-22/">暗号化 プライバシーを救った反乱者たち</a></dt><dd>スティーブン・レビー 斉藤 隆央 </dd><dd>紀伊國屋書店 2002-02-16</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/487593100X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/487593100X.09._SCTHUMBZZZ_.jpg"  alt="ハッカーズ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4105393022/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4105393022.09._SCTHUMBZZZ_.jpg"  alt="暗号解読―ロゼッタストーンから量子暗号まで"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4484111160/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4484111160.09._SCTHUMBZZZ_.jpg"  alt="グーグル ネット覇者の真実 追われる立場から追う立場へ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/410215972X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/410215972X.09._SCTHUMBZZZ_.jpg"  alt="暗号解読〈上〉 (新潮文庫)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4102159738/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4102159738.09._SCTHUMBZZZ_.jpg"  alt="暗号解読 下巻 (新潮文庫 シ 37-3)"  /></a> </p>
<p class="description">20世紀末，暗号技術の世界で何があったのか。知りたかったらこちらを読むべし！</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-03-09">2015/03/09</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
