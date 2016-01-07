+++
date = "2016-01-07T20:50:45+09:00"
description = "Bitcoin は「誰が」取り引きしているかについて気にしないが，もっと広く smart contract を考えるのなら何らかの形で「誰が」を保証しなければならないし，そのための基盤（infrastructure）が必要となる，と思うのだがどうだろう。"
draft = false
tags = ["blockchain", "smart-contract", "fintech", "engineering"]
title = "Blockchain と Smart Contract"

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
「[東雲銀子さん](http://www.amazon.co.jp/exec/obidos/ASIN/B00ORBY3PQ/baldandersinf-22/)」みたいな人は要らなくなるかもしれない。


ちなみに昨年までの私の Bitcoin や Blockchain に関する理解は以下の通り。

- [そろそろ Blockchain について勉強を始めるか — Baldanders.info](http://www.baldanders.info/spiegel/log2/000827.shtml)

このときのまとめをもう一度整理して再掲載しておくと

1. Blockchain は「鎖」で繋がれた追記型データベース。「鎖」の途中のデータは取り消しも変更（改竄）もできない
2. Blockchain の追記プロセスには不正の余地がないよう何らかの仕掛けが必要。 Bitcoin の場合は「作業証明（proof-of-work）」がそれにあたる
3. Blockchain は P2P による分散型かつ fault tolerant（過失を許容する）なシステムだが最終的には fork も merge も許容しない
4. Bitcoin のアドレス（実体は公開鍵）が誰に帰属するか Bitcoin/Blockchain は関知しない。Bitcoin が気にするのは Blockchain に記載されるアドレスの一貫性と無矛盾性である。アドレスの証明が必要な場合は外部の PKI を利用するか新たに組み込む必要がある

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

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W1BX3N2/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51aobp3mndL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W1BX3N2/baldandersinf-22/">アンダーグラウンド・マーケット (朝日新聞出版)</a></dt><dd>藤井太洋 </dd><dd>朝日新聞出版 2015-03-20</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00V7Y7DUS/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00V7Y7DUS.09._SCTHUMBZZZ_.jpg"  alt="ビッグデータ・コネクト (文春文庫)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00I3W45AS/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00I3W45AS.09._SCTHUMBZZZ_.jpg"  alt="オービタル・クラウド"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00CHIFA1M/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00CHIFA1M.09._SCTHUMBZZZ_.jpg"  alt="Gene Mapper -full build-"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B019FV4BGI/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B019FV4BGI.09._SCTHUMBZZZ_.jpg"  alt="天冥の標Ⅸ　PART1　ヒトであるヒトとないヒトと (ハヤカワ文庫JA)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B019FV4BGS/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B019FV4BGS.09._SCTHUMBZZZ_.jpg"  alt="中継ステーション〔新訳版〕 (ハヤカワ文庫SF)"  /></a> </p>
<p class="description">デジタル通貨ならこれかな。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-01-07">2016-01-07</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B00ORBY3PQ/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51d7PAEntoL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B00ORBY3PQ/baldandersinf-22/">銀子の窓口（１） (バンブーコミックス 4コマセレクション)</a></dt><dd>唐草ミチル </dd><dd>竹書房 2014-10-27</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0185M2QAQ/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0185M2QAQ.09._SCTHUMBZZZ_.jpg"  alt="秘書の仕事じゃありません　２巻 (まんがタイムコミックス)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0197POLV2/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0197POLV2.09._SCTHUMBZZZ_.jpg"  alt="軍神ちゃんとよばないで　２巻 (まんがタイムコミックス)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B018K0HAIQ/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B018K0HAIQ.09._SCTHUMBZZZ_.jpg"  alt="私設花野女子怪館（２） (バンブーコミックス MOMOセレクション)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B013FVA8KW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B013FVA8KW.09._SCTHUMBZZZ_.jpg"  alt="腹黒舞子さんとの京生活（２） (バンブーコミックス 4コマセレクション)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B017XKJD84/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B017XKJD84.09._SCTHUMBZZZ_.jpg"  alt="うしろのご先祖さま（２） (バンブーコミックス 4コマセレクション)"  /></a> </p>
<p class="description">金融界のクール・ビューティ。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-01-07">2016-01-07</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
