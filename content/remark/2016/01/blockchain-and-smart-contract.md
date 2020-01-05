+++
date = "2016-01-07T20:50:45+09:00"
description = "Bitcoin は「誰が」取り引きしているかについて気にしないが，もっと広く smart contract を考えるのなら何らかの形で「誰が」を保証しなければならないし，そのための基盤（infrastructure）が必要となる，と思うのだがどうだろう。"
tags = ["blockchain", "smart-contract", "fintech", "engineering"]
title = "Blockchain と Smart Contract"

[scripts]
  mathjax = false
  mermaidjs = false
+++

年末忙しくて書き損ねてるネタを回収中。

[「YAMDAS現更新履歴」の記事](http://d.hatena.ne.jp/yomoyomo/20151223/blockchain)より。

- [世界の金融機関がフィンテックの本命としてブロックチェーン技術にこぞって投資する理由とは？（全3話中1話目） | コインテレグラフジャパン](http://jp.cointelegraph.com/news/115869/what-is-the-reason-that-a-world-financial-institution-invests-in-a-block-chain-technology-as-a-favorite-of-the-fin-technical-center-all-togetherjp1)
- [世界の金融機関がフィンテックの本命としてブロックチェーン技術にこぞって投資する理由とは？（全3話中2話目） | コインテレグラフジャパン](http://jp.cointelegraph.com/news/115872/what-is-the-reason-that-a-world-financial-institution-invests-in-a-block-chain-technology-as-a-favorite-of-the-fin-technical-center-all-togetherjp2)
- [世界の金融機関がフィンテックの本命としてブロックチェーン技術にこぞって投資する理由とは？（全3話中3話目） | コインテレグラフジャパン](http://jp.cointelegraph.com/news/115873/what-is-the-reason-that-a-world-financial-institution-invests-in-a-block-chain-technology-as-a-favorite-of-the-fin-technical-center-all-togetherjp3)

昨年この記事を斜め読みした時は気づかなかったが

- [さくらインターネットとテックビューロ、ブロックチェーンの実証実験環境「mijinクラウドチェーンβ」を金融機関やITエンジニア向けに無料提供～本日より申込受付を開始し、2016年1月から順次提供～](http://www.sakura.ad.jp/press/2015/1216_mijin_cloud_chain/)
- [「さくらのクラウド」、「ASTERIA WARP」、「mijin」の3社製品･サービスによるプライベート･ブロックチェーン＆IoT実証実験プラットフォームの無償提供開始について＜無償提供期間：2016年1月18日～6月30日＞](http://www.sakura.ad.jp/press/2016/0107_demonstration_platform/)

あぁ，これって広告記事だったのか。
感度低いなぁ，私（笑）

上の記事からいくつか気になる部分を抜き出しておく。

{{< fig-quote title="世界の金融機関がフィンテックの本命としてブロックチェーン技術にこぞって投資する理由とは？（全3話中2話目）" link="http://jp.cointelegraph.com/news/115872/what-is-the-reason-that-a-world-financial-institution-invests-in-a-block-chain-technology-as-a-favorite-of-the-fin-technical-center-all-togetherjp2" >}}
<q>しかし、ブロックチェーン技術の登場により、既存の様々なセキュリティや整合性といった基準を妥協することなく、トランザクションあたりのコストが限りなくゼロに近づいていきます。そこで新しい常識として、「リアルタイムセトルメント」の概念が現実となります。そして締めの集計概念も、そもそもトランザクションに残高（バランス）の概念を持つブロックチェーンの基本機能でまかなえます。<br>
ビジネス慣習が締め処理からリアルタイム決済へパラダイムシフトすることにより、カウンターパーティーリスクも最小限となるだけではなく、既存の様々な頭痛の種が払拭されるでしょう。バッチ処理がされていたものに関してもリアルタイム処理へと移行し、真のリアルタイム24時間サービスの提供など、一般消費者もそこから大きな恩恵を得るに違いありません。</q>
{{< /fig-quote >}}

{{< fig-quote title="世界の金融機関がフィンテックの本命としてブロックチェーン技術にこぞって投資する理由とは？（全3話中2話目）" link="http://jp.cointelegraph.com/news/115872/what-is-the-reason-that-a-world-financial-institution-invests-in-a-block-chain-technology-as-a-favorite-of-the-fin-technical-center-all-togetherjp2" >}}
<q>現在主流のブロックチェーン技術は、処理の負荷を分散することが主目的ではないという点が分散型DBとは大きく違います。あくまでも主目的は物理的分散によるダウンタイムの払拭です。我々はデータ自体の分散や処理の分散の研究も進めておりますが、現在では全てのノードが同じデータを保有し、同じ仕事をすることができます。従って一般的なブロックチェーンの世界では、台数の増加がパフォーマンスの向上に直結しません。よって、「最低何台必要ですか？」の問いに対する答えは、「1台」となります。<br>
現在のブロックチェーン技術が、最新の分散型DBにデータベースとしてのスループット性能では勝つことはできません。ブロックチェーンが持つ大きな利点は、現在の金融システム等で必要とされる処理量であれば、データの整合性を保ったまま、物理的分散によりゼロダウンタイムを実現し、かつそこに同時に劇的なコスト削減をもたらすということに集約されます。</q>
{{< /fig-quote >}}

Blockchain の特徴は完全性と可用性にある，ということのようだ。
「[東雲銀子さん](https://www.amazon.co.jp/exec/obidos/ASIN/B00ORBY3PQ/baldandersinf-22/)」みたいな人は要らなくなるかもしれない。


ちなみに昨年までの私の Bitcoin や Blockchain に関する理解は以下の通り。

- [そろそろ Blockchain について勉強を始めるか — Baldanders.info](https://baldanders.info/blog/000827/)

このときのまとめをもう一度整理して再掲載しておくと

1. Blockchain は「鎖」で繋がれた追記型データベース。「鎖」の途中のデータは取り消しも変更（改竄）もできない
2. Blockchain の追記プロセスには不正の余地がないよう何らかの仕掛けが必要。 Bitcoin の場合は「作業証明（proof-of-work）」がそれにあたる
3. Blockchain は P2P による分散型かつ fault tolerant（過失を許容する）なシステムだが最終的には fork も merge も許容しない
4. Bitcoin のアドレス（実体は公開鍵）の帰属先について Bitcoin/Blockchain は関知しない。Bitcoin が気にするのは Blockchain に記載されるアドレスの一貫性と無矛盾性である。アドレスの証明が必要な場合は外部の PKI を利用するか新たに組み込む必要がある

私の関心領域は「信用モデル（trust model）」にあるため，このようなまとめになっている。
Bitcoin は「誰が」取り引きしているかについて気にしない[^a] が，もっと広く smart contract を考えるのなら何らかの形で「誰が」を保証しなければならないし，そのための基盤（infrastructure）が必要となる，と思うのだがどうだろう。

[^a]: これは匿名とは異なる。 Bitcoin の元帳である Blockchain は同期されたコピーを誰でも閲覧でき，同じ Bitcoin アドレスを使い続けるならそのアドレスに対する取引履歴を抽出することもできる。アドレスや取引履歴をもとに対象を絞り込むのは面倒だが不可能ではない。

私は Bitcoin そのものへの関心が薄いのだが，これは地方から見て Bitcoin がアベノミクスのごとき「彼岸」の話であることと，いわゆる FinTech（Financial Technology） なるバズワードは「畑違い」だから。
でも Blockchain 自身は FinTech 分野におさまらない可能性を持っていると思う。
ので，今後も情報を追いかけていくつもりである。

## 参考ページ

- [誰も教えてくれないけれど、これを読めば分かるビットコインの仕組みと可能性 | TechCrunch Japan](http://jp.techcrunch.com/2015/03/31/bitcoin-essay/)
- [（日銀レビュー）「デジタル通貨」の特徴と国際的な議論 ：日本銀行 Bank of Japan](http://www.boj.or.jp/research/wps_rev/rev_2015/rev15j13.htm/)
- [スマートコントラクトとは何か? Smart Contractsの言葉の定義 - Qiita](http://qiita.com/hshimo/items/093f40b856ba2436fbba)
- [Smart Contract - 暗号通貨のお勉強　～Bitcoin、Rippleを中心に～](http://cryptocoin.hatenablog.com/entry/2015/07/22/001500)
- [BitCoinとBlockChainにまつわる誤解ーそんなことはできない - Qiita](http://qiita.com/tatarou1986/items/9d994896795a4871dc37)

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/UNDERGROUND-MARKET-%E3%83%92%E3%82%B9%E3%83%86%E3%83%AA%E3%82%A2%E3%83%B3%E3%83%BB%E3%82%B1%E3%83%BC%E3%82%B9-%E8%97%A4%E4%BA%95%E5%A4%AA%E6%B4%8B-ebook/dp/B00FONW2V8?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00FONW2V8"><img src="https://images-fe.ssl-images-amazon.com/images/I/51AT2LqRIsL._SL160_.jpg" width="116" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/UNDERGROUND-MARKET-%E3%83%92%E3%82%B9%E3%83%86%E3%83%AA%E3%82%A2%E3%83%B3%E3%83%BB%E3%82%B1%E3%83%BC%E3%82%B9-%E8%97%A4%E4%BA%95%E5%A4%AA%E6%B4%8B-ebook/dp/B00FONW2V8?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00FONW2V8">UNDERGROUND MARKET　ヒステリアン・ケース</a></dt>
	<dd>藤井太洋</dd>
    <dd>朝日新聞出版 2013-11-07 (Release 2013-10-25)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00FONW2V8</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">日本で「仮想通貨」が流行る前に登場した傑作。つかエンジニアは全員「UNDERGROUND MARKET」シリーズを読め！</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-02-04">2018-02-04</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E9%8A%80%E5%AD%90%E3%81%AE%E7%AA%93%E5%8F%A3/dp/B074CJ5VJJ?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B074CJ5VJJ"><img src="https://images-fe.ssl-images-amazon.com/images/I/51UmNb4fC8L._SL160_.jpg" width="160" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E9%8A%80%E5%AD%90%E3%81%AE%E7%AA%93%E5%8F%A3/dp/B074CJ5VJJ?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B074CJ5VJJ">銀子の窓口</a></dt>
	<dd>唐草ミチル</dd>
    <dd></dd>
    <dd>Collections Kindle版</dd>
    <dd>ASIN: B074CJ5VJJ</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">金融界のクール・ビューティ（笑）</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-01-07">2016-01-07</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>
