+++
tags = ["code", "internet", "cryptography", "drm", "bitcoin", "blockchain", "grigori"]
description = "「ウォルター・アイザックソン「インターネットは壊れている。初めからやり直すなら、私はこのようにそれを正す」 - YAMDAS現更新履歴」より"
date = "2016-12-28T12:49:19+09:00"
update = "2017-10-15T09:23:00+09:00"
title = "ぼくがかんがえたさいきょうのいんたねっと"
draft = false

[author]
  name = "Spiegel"
  flickr = "spiegel"
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  tumblr = ""
  url = "https://baldanders.info/spiegel/profile/"
  instagram = "spiegel_2007"
  twitter = "spiegel_2007"
  github = "spiegel-im-spiegel"
  license = "by-sa"
  flattr = ""
  linkedin = "spiegelimspiegel"
+++

- [ウォルター・アイザックソン「インターネットは壊れている。初めからやり直すなら、私はこのようにそれを正す」 - YAMDAS現更新履歴](http://d.hatena.ne.jp/yomoyomo/20161226/walterisaacson)

この記事の中で以下の2つの記事が紹介されている。

- [The internet is broken. Starting from scratch, here's how I'd fix it. | Walter Isaacson | LinkedIn](https://www.linkedin.com/pulse/internet-broken-starting-from-scratch-heres-how-id-fix-isaacson)
- [Fixing The Internet – AVC](http://avc.com/2016/12/fixing-the-internet/)

最初の記事で「俺ならこう直すぜ」という部分で挙げられている項目は以下の5つ。

1. Create a system that enables content producers to negotiate with aggregators and search engines to get a royalty whenever their content is used, like ASCAP has negotiated for public performances and radio airings of its members’ works. （コンテンツの作者がアグリゲータや検索エンジンと交渉して、コンテンツが使用されるたびに印税を得られるようにするシステムを構築する。ASCAP（米国作曲家作詞家出版者協会）が、その会員の作品のライブ演奏やラジオ放送において行っているような）
2. Embed a simple digital wallet and currency for quick and easy small payments for songs, blogs, articles, and whatever other digital content is for sale. （楽曲、ブログ、記事、その他のデジタルコンテンツが売れる、手早く簡単なスモールペイメントに使えるシンプルなデジタルの財布や通貨を埋め込む）
3. Encode emails with an authenticated return or originating address. （電子メールは認証された相手か送信元アドレスで暗号化）
4. Enforce critical properties and security at the lowest levels of the system possible, such as in the hardware or in the programming language, instead of leaving it to programmers to incorporate security into every line of code they write. （クリティカルなプロパティやセキュリティをシステムの可能な限り最も低い層にする。セキュリティをプログラマが書くすべての行に組み込むようプログラマ任せにするのではなく、ハードウェアやプログラミング言語のレベルで実現させる）
5. Build chips and machines that update the notion of an internet packet. For those who want, their packets could be encoded or tagged with metadata that describe what they contain and give the rules for how it can be used. （インターネットのパケットの概念を更新するチップやマシンを作る。希望する人は、パケットを暗号化し、どのように利用可能かルールを指定するメタデータでタグ付け可能にする）

（日本語訳は [yomoyomo さんの記事]を拝借した。なお [yomoyomo さんの記事]は [BY-NC-SA](https://creativecommons.org/licenses/by-nc-sa/4.0/) で公開されているので取り扱いに注意）

最初の2つはコンテンツ・マネジメントをネットワークレベルで組み込むという話。
3番目は暗号化メッセージングだね。
最後の2つは暗号化の機能をチップレベルで実装しようという話であると理解した。
以降，この3つのテーマについて考えてみよう。

## コンテンツ・マネジメントをネットワークレベルで組み込む

[yomoyomo さんの記事]と同じく私も Xanadu を連想した。
Xanadu については以下の記事が参考になるだろう。

- [リテラリーマシン―ハイパーテキスト原論 - 雑記帳](http://d.hatena.ne.jp/ced/20060516/1147704205)
- [文化庁は和式Xanaduでも作ればいいんじゃない？ - 雑記帳](http://d.hatena.ne.jp/ced/20071220/1198144063)

{{< fig-quote title="リテラリーマシン―ハイパーテキスト原論" link="http://d.hatena.ne.jp/ced/20060516/1147704205" >}}
<q>Xanaduの実現を難しくしている原因の一つは、著作権問題にある。いわゆるマイクロペイメント・システムが実現し、データ通信量に応じて著作権料が支払われるシステムをXanaduは想定しているけれど、今のところこれを可能にするマイクロペイメント・システムは実現されていない。バージョン管理システムも、CVSのようなアプリケーションとしては存在していても、HTMLそのものにバージョン管理システムが埋め込まれているわけではない。そして、XanaduとWorld Wide Web(インターネット)の歴史を隔てる一番の要因は、開発方法そのもの。インターネットは徹底してオープンな姿勢を貫いてきた。勿論例外も存在するけれど、基本的にソースは公開されることが原則で、だからこそGPLやLinuxなどといったものが生まれて来た。それに対して Xanaduはソースの内容を一切公開せず特許化してしまっている。だから、ネルソンの語るXanaduの理想郷がどれだけ素晴らしいものであったとしても、開発グループに入ることができなければ、Xanaduに関わることはできない。HTMLは確かに貧弱なマークアップ言語だが、その記述方法が簡単であったからこそ、ここまで普及することができたのも事実なのである。もし、仮にXanaduがオープンソースであったのなら、もしかしたらHTMLではなくXanaduが今のインターネットの主流になっていたのかもしれない。それでもXanaduはHTMLやXMLで実現できそうにない部分を未だに含んでおり、今後もその理念は必要とされていくのだろう。</q>
{{< /fig-quote >}}

というわけで “[Fixing The Internet]” では「[Mediachain] と Bitcoin があるぢゃん！」と言っているわけである。

{{< fig-quote title="Fixing The Internet" link="http://avc.com/2016/12/fixing-the-internet/" lang="en" >}}
<q>While not “a system to negotiate”, services like our portfolio company <a href="http://www.mediachain.io/">Mediachain</a>‘s platform will provide much of the underlying infrastructure for this to happen.</q>
{{< /fig-quote >}}

[Mediachain] については以下の記事が参考になるかな。

- [その画像は誰のもの？　Mediachainが画像の帰属を検索できるアトリビューション・エンジンを公開 | ビットコインの最新情報 BTCN｜ビットコインニュース](http://btcnews.jp/mediachain-attribution-engine/)

このタイプの DRM (Digital Rights Management) については拙文「[CC Licenses における「技術的保護手段」の扱い]({{< ref "/cc-licenses/05-technological-measures.md" >}})」でも取り上げた。

> もうひとつは DRM が「監視型」に移行したことである。「電子透かし」や「電子指紋」を使ってネット上に流通するコンテンツを比較的容易に監視できるようになった。これがうまく機能すれば一般ユーザのネット上での活動を妨げることなく悪質なものだけを排除することができる。

監視型の DRM の問題は，監視を行う「技術的ゲートキーパー[^tgk]」の力が強くなりすぎることである（またはコンテンツホルダーと癒着しやすい）。
流通するコンテンツを監視することはそれに紐付かれるユーザ（著者とは限らない）を監視することにつながる。
サービス・プロバイダ側が「私たちはコンテンツの監視しかしてませんよ」と無邪気に言ったところで通用するものではないのである。
おそらく監視型 DRM は知財トロルにとって最も効果的な道具になるだろう。

[^tgk]: 「技術的ゲートキーパー」については拙文「[監視をコントロールする](https://baldanders.info/spiegel/log2/000490.shtml)」を参考にどうぞ。

## 暗号化メッセージング

何が言いたいのか分からない。
“[Fixing The Internet]” では

{{< fig-quote title="Fixing The Internet" link="http://avc.com/2016/12/fixing-the-internet/" lang="en" >}}
<q>While not blockchain based, standards like DKIM and SPF in the email sector provide some of this today. I am also excited about a blockchain based identity later like the one being built by our portfolio company <a href="https://blockstack.org/posts/blockchain-identity">Blockstack</a>.</q>
{{< /fig-quote >}}

とあるが，これも意味不明。
それともやっぱり “Encode” を「暗号化」と解釈してはいけないのだろうか。

暗号化メッセージングについては以下で言及しているので，ここでは割愛する。

- [PGP ってゆーな！]({{< ref "/remark/2016/12/do-not-call-it-pgp.md" >}})

ぶっちゃけサーバ間でバケツリレーするメッセージング・サービスは（spam も含めて）攻撃の対象になりやすい。
伝統的な「電子メール」はゆるゆると役目を終えていくと思うよ。

## 暗号化の機能をチップレベルで実装

これには苦笑せざるを得ない。
「それなんて[クリッパーチップ](https://baldanders.info/spiegel/log2/000854.shtml "米国は今もクリッパーチップの夢を見るか，他 — Baldanders.info")？」である。

ネットワーク通信のプロトコルスタックを組み込んだチップというのは既に存在する。
たとえば車載ネットワークではリアルタイム[^rt] が要求されるため CAN/LIN など[^vn] はチップレベルで実装されている。

[^rt]: コンピュータの世界で言う「リアルタイム」は決められた時間以内に特定の処理を完了することが保証されていることを指す。
[^vn]: 他にも欧州規格の [MOST](http://car.watch.impress.co.jp/docs/series/cew/537764.html "【大原雄介のカーエレWatch】車載Networkの話（5）「MOST」 - Car Watch") や Maxim 社製の [GMSL](http://monoist.atmarkit.co.jp/mn/articles/1308/20/news068.html "車載半導体：「同軸ケーブル対応SERDESが最適解」、マキシムが車両内の映像データ伝送で提案 - MONOist（モノイスト）") といったものもある。また近年では，こうした複数のネットワークを Ethernet で統一しようという動きもある。

しかし通信上のセキュリティ要件や暗号要件をチップレベルで実装するのはかなり難しい。
つか現状では無理（軽量暗号の研究はある）。

ちなみにクリッパーチップが挫折した直接の原因は実装した暗号アルゴリズムに脆弱性が発見されたためであった。
セキュリティ要件は（社会の変化に応じて）変化していくし，暗号アルゴリズムもいつか破られる可能性がある[^cr]。
これはプログラミング言語でも同様[^pr]。

[^cr]: 故に暗号システムを実装する際には必ず複数のアルゴリズムを実装または実装可能にする必要がある。
[^pr]: もちろんプログラミング言語側でも努力をしているよ。先日紹介した「[null 安全]({{< ref "/remark/2016/11/null-safety.md" >}} "「null 安全」について")」もうそうした努力のひとつである。

これはぜひ肝に銘じておいた方がいいと思うが，コンピュータ・ネットワークのセキュリティに関しては常に攻撃者の方が有利なのである。
不利な戦いを何とか凌ぐにはチップに籠城するよりは攻撃者に適応してこちらも動的に力を付けていくしかない。

“[Fixing The Internet]” では

{{< fig-quote title="Fixing The Internet" link="http://avc.com/2016/12/fixing-the-internet/" lang="en" >}}
<q>Blockchain based applications can use the underlying security of the blockchain (using sophisticated cryptography) to achieve higher levels of security in their applications.</q>
{{< /fig-quote >}}

と述べている。

ところで，何故かみんな褒め殺すばかりな気がするのだが Bitcoin には大きな問題がある。
みなさん暗号通信の4要件を覚えておられるだろうか。
すなわち

1. 機密性（Confidentiality）
2. 完全性（Integrity）
3. 認証（Authentication）
4. 否認防止（Non-repudiation）

の4つである[^em]。
このうち Bitcoin は完全性しか満たさない。
Bitcoin のアドレス（Blockchain の識別子でもある）は公開鍵暗号の鍵ペアそのものだが，これとユーザを紐付かせる情報を一切持っていない。
確かに Bitcoin に機密性は不要だが，他の認証や否認防止といった要件は意図的に削除されている。
そうすることによってひたすら log の完全性のみを追求したのが Bitcoin でありその中核技術たる Blockchain なのだ。

[^em]: ただし近年は「暗号化メッセージング」について要件が変わってきている。詳しくは拙文「[メッセージングは E2E 暗号化および PFS が肝](https://baldanders.info/spiegel/log2/000675.shtml "メッセージングは E2E 暗号化および PFS が肝 — Baldanders.info")」を参考にどうぞ。今回はコンテンツの取引に絡む話なので「否認防止」は必須要件になる。また「否認防止」を満たすには「完全性」と「認証」が必要条件になる。

- [そろそろ Blockchain について勉強を始めるか — Baldanders.info](https://baldanders.info/spiegel/log2/000827.shtml)

Bitcoin というか Blockchain を使う以上，認証および否認防止は別の手段（不特定多数と取引するのなら何らかの認証基盤）を以って担保する必要がある。
なので少なくとも「Blockchain があればおっけー」とはならないと思う。
まぁ FinTech なるバズワードに関わる人達には自明なんだろうけど。

## ブックマーク

- [誰も教えてくれないけれど、これを読めば分かるビットコインの仕組みと可能性 | TechCrunch Japan](http://jp.techcrunch.com/2015/03/31/bitcoin-essay/)
- [RFC 6962 - Certificate Transparency](https://tools.ietf.org/html/rfc6962)
    - [Certificate Transparency | GMOグローバルサインブログ](https://jp.globalsign.com/blog/2014/certificate_transparency.html)
    - [Google による OpenPGP 鍵配送の解決提案 — Baldanders.info](https://baldanders.info/spiegel/log2/000785.shtml)
- [Blockchain と Smart Contract]({{< ref "/remark/2016/01/blockchain-and-smart-contract.md" >}})

[yomoyomo さんの記事]: http://d.hatena.ne.jp/yomoyomo/20161226/walterisaacson "ウォルター・アイザックソン「インターネットは壊れている。初めからやり直すなら、私はこのようにそれを正す」 - YAMDAS現更新履歴"
[Fixing The Internet]: http://avc.com/2016/12/fixing-the-internet/ "Fixing The Internet – AVC"
[Mediachain]: http://www.mediachain.io/

## 参考図書

<div class="hreview" ><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B00I6IT1U8/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/41PrUBQnZ5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B00I6IT1U8/baldandersinf-22/">ビットコイン解説本</a></dt><dd>足立　明穂 </dd><dd> 2014-02-01</dd><dd>評価<abbr class="rating" title="4"><img src="https://images-fe.ssl-images-amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00DFZLTPM/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B00DFZLTPM.09._SCTHUMBZZZ_.jpg"  alt="クラウドソーシングの衝撃 (NextPublishing)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00L8MXE9I/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B00L8MXE9I.09._SCTHUMBZZZ_.jpg"  alt="ビットコイン解説本2　＜投資編＞"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B007SRLKTS/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B007SRLKTS.09._SCTHUMBZZZ_.jpg"  alt="小心者でもサラリとかわせる「断る」心理テクニック"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00J9EN42O/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B00J9EN42O.09._SCTHUMBZZZ_.jpg"  alt="仏教は宗教ではない ～お釈迦様が教えた完成された科学～"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00JGOROY6/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B00JGOROY6.09._SCTHUMBZZZ_.jpg"  alt="広告の基本　この１冊ですべてわかる"  /></a> </p>
<p class="description" >最初に読んだ Bitcoin 解説本。手軽に読める。</p>
<p class="gtools" >reviewed by <a href="#maker" class="reviewer">Spiegel</a> on <abbr class="dtreviewed" title="2015-04-02">2015/04/02</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html">G-Tools</a>)</p>
</div>

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
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E5%8C%96-%E3%83%97%E3%83%A9%E3%82%A4%E3%83%90%E3%82%B7%E3%83%BC%E3%82%92%E6%95%91%E3%81%A3%E3%81%9F%E5%8F%8D%E4%B9%B1%E8%80%85%E3%81%9F%E3%81%A1-%E3%82%B9%E3%83%86%E3%82%A3%E3%83%BC%E3%83%96%E3%83%B3%E3%83%BB%E3%83%AC%E3%83%93%E3%83%BC/dp/4314009071?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4314009071"><img src="https://images-fe.ssl-images-amazon.com/images/I/51ZRZ62WKCL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E5%8C%96-%E3%83%97%E3%83%A9%E3%82%A4%E3%83%90%E3%82%B7%E3%83%BC%E3%82%92%E6%95%91%E3%81%A3%E3%81%9F%E5%8F%8D%E4%B9%B1%E8%80%85%E3%81%9F%E3%81%A1-%E3%82%B9%E3%83%86%E3%82%A3%E3%83%BC%E3%83%96%E3%83%B3%E3%83%BB%E3%83%AC%E3%83%93%E3%83%BC/dp/4314009071?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4314009071">暗号化 プライバシーを救った反乱者たち</a></dt>
	<dd>スティーブン・レビー</dd>
	<dd>斉藤 隆央 (翻訳)</dd>
    <dd>紀伊國屋書店 2002-02-16</dd>
    <dd>Book 単行本</dd>
    <dd>ASIN: 4314009071, EAN: 9784314009072</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">20世紀末，暗号技術の世界で何があったのか。知りたかったらこちらを読むべし！</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-16">2018-12-16</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
