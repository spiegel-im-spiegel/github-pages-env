+++
title = "いまさら調べるデジタル主権"
date =  "2026-08-27T19:18:49+09:00"
description = "米国の政変をきっかけに国際政治的には色々と変化が起きているが，個人のレベルに落とし込んだときに，デジタル主権を背負うってのは簡単なことじゃないとは思うよ。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "security", "politics", "internet", "privacy", "cloud", "risk", "international" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

中国の兵法書に「兵法三十六計」というのがあって，この中で劣勢の場合に用いる奇策として「{{% ruby "にぐるをじょうとなす" %}}走為上{{% /ruby %}}」というのがある。
日本で皮肉っぽく言われる「逃げるが勝ち」とか近年流行った「逃げるは恥だが役に立つ」というのとは異なり，自身が不利なら無理なく撤退し「次」のために戦力を温存すべき，という積極的な撤退策である。

{{< fig-quote type="markdown" title="走為上 - Wikipedia" link="https://ja.wikipedia.org/wiki/%E8%B5%B0%E7%82%BA%E4%B8%8A" >}}
{{< fig-quote type="markdown" >}}
全師避敵　左次無咎　未失常也
{{< /fig-quote >}}
将兵を全うしたまま強敵を回避する。これを咎なく（秩序立って）行えれば、失うものもなく、兵法の常道に適う。
{{< /fig-quote >}}

## 国家による制裁とデジタル主権の実践

{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}

国際刑事裁判所（International Criminal Court; ICC）の赤根智子所長[^icc1] が（大統領令に基づき）米国外務省によって SDN リスト（Specially Designated Nationals and Blocked Persons List; 特別指定国民および資格停止者リスト）に追加された件について，産経新聞のインタビュー（有料記事）に答えて曰く

[^icc1]: 赤根智子氏は2024年から ICC 所長に就任している。

{{< fig-quote type="markdown" title="ICC赤根所長 会見要旨「クレジットカード、メール使えず」「EUで制裁の無効化訴え」 - 産経ニュース（無料パート）" link="https://www.sankei.com/article/20260824-4MUPB6UV4VINFAIUGYKLOU5D6I/" >}}
「米社VISAのクレジットカードは日本で発行したものを含めて、使えないと思う。制裁により、私用のGmailも使えなくなる。昨年、米国がICC制裁に動いた後、ICCでは米マイクロソフトから欧州拠点のシステムへのデータ移転を進めており、日常業務に支障はない。金融機関では制裁の波及に対する懸念が強いが、欧州で複数の銀行が制裁を受けたICC職員を支援し、取引を継続してくれている。出張の旅行手配、給与支払いなどは、これらの銀行を通じてやっている」
{{< /fig-quote >}}

私は「政治的無関心」を年間目標にしてるので（日本政府の対応を含め）この件について何らかの立場を表明することはしないが，重要なポイントは ICC が欧州サービスへのデータ移転を進めていることと，欧州の複数の銀行が支援している点だろう。

これについて Bluesky で以下のポストを見かけた。

{{< fig-gen >}}
<blockquote class="bluesky-embed" data-bluesky-uri="at://did:plc:plqlg756md5grgwwynvcp5vt/app.bsky.feed.post/3mtvl7bcjfc2z" data-bluesky-cid="bafyreifhmsr2voj2lmyjgxsgc2y3hxlsvop6yz7w7tu2dxc7uq3324vtaq" data-bluesky-embed-color-mode="system"><p lang="ja">ICC赤根所長　会見要旨「クレジットカード、メール使えず」「EUで制裁の無効化訴え」 www.sankei.com/article/2026...
「昨年、米国がICC制裁に動いた後、ICCでは米マイクロソフトから欧州拠点のシステムへのデータ移転を進めており、日常業務に支障はない。」
Digital sovereignty の実験場になっている
bsky.app/profile/tuta...<br><br><a href="https://bsky.app/profile/did:plc:plqlg756md5grgwwynvcp5vt/post/3mtvl7bcjfc2z?ref_src=embed">[image or embed]</a></p>&mdash; Shinji R. Yamane (<a href="https://bsky.app/profile/did:plc:plqlg756md5grgwwynvcp5vt?ref_src=embed">@shinjiyamane.bsky.social</a>) <a href="https://bsky.app/profile/did:plc:plqlg756md5grgwwynvcp5vt/post/3mtvl7bcjfc2z?ref_src=embed">August 25, 2026 at 7:08 PM</a></blockquote>
{{< /fig-gen >}}

なるほど。
デジタル主権の実験場か。

## いまさら「デジタル主権」とは

"Digital sovereignty" は日本語では「デジタル主権」と訳されるようだ。
文脈によっては「デジタル自律性」とかいうこともあるらしい。

デジタル主権の資料として [Kagi Assistant] に以下の PDF 文書を教えてもらった。

- {{< pdf-file title="Digital sovereignty for Europe" link="https://www.europarl.europa.eu/RegData/etudes/BRIE/2020/651992/EPRS_BRI(2020)651992_EN.pdf" >}}

欧州議会調査局（EPRS）が2020年7月に公表した（政策提案を含む）報告書で，デジタル主権を以下のように定義付けている（[Kagi Assistant] による要約）。

{{< div-ai type="markdown" >}}
報告書におけるデジタル主権とは，**欧州がデジタル世界で独立して行動する能力**です。

ただし，これは外国企業や外国技術を一切排除することを意味しません。報告書は，デジタル主権を次の二つの側面から捉えています。

- 外部からの脅威や過度な依存からデータ，インフラ，市民の権利を守ること
- デジタル・イノベーションを促進し，欧州自身が技術やサービスを生み出すこと

したがって，非EU企業との協力を含めながらも，欧州が重要な技術，データ，規制について自律的に選択し，必要に応じて代替手段を持てる状態を目指す概念です。
{{< /div-ai >}}

2020年のこの資料は，2016年の米国大統領選挙におけるロシア介入疑惑や Cambridge Analytica 問題が絡んでいると思われる。
報告書は「EU が直面する主な課題」として以下を挙げているようだ。

{{< div-ai type="markdown" >}}
**1 デジタル技術の開発競争での遅れ**

AI，5G，クラウド，IoT，量子技術などは，今後の経済成長と安全保障に関わる戦略的技術です。しかし，AI分野では，EUは米国や中国に比べて民間投資，企業や市民による導入，人材の獲得，特許出願などで後れを取っていると報告書は説明しています。

米国はAI人材，研究者，特許，民間投資で優位に立ち，中国はデータの収集・利用やスーパーコンピューターなどのハードウェアで進展しています。EUが外国技術に依存すれば，経済成長だけでなく，国際的な影響力も低下する可能性があります。

**2 市民のデータとプライバシーのコントロール**

Google，Apple，Facebook，Amazon，Microsoftなどの大手プラットフォーム企業は，大量の個人データを収集し，広告などに利用するビジネスモデルを発展させてきました。

報告書は，このようなデータの収集・分析・商業利用が，市民から自分の情報を管理する力を奪う可能性を指摘しています。Cambridge Analytica問題は，オンライン上の個人データが政治的プロファイリングにも利用され得ることを示す事例として扱われています。

**3 クラウドとデータ基盤への依存**

世界のクラウド市場は，米国やアジアの企業が大きく支配していました。そのため，欧州の政府や企業が域外企業のクラウドサービスを利用すると，データの管理や安全性に関する問題が生じる可能性があります。

報告書は，特に米国のCLOUD Actによって，米国当局が一定の条件の下で外国にあるデータへのアクセスを求める可能性があることを，欧州側の懸念として紹介しています。

**4 5Gとサイバーセキュリティ**

中国企業など域外企業への5Gインフラ依存は，供給停止，外国からの影響，サイバー攻撃などのリスクにつながり得ます。また，EU加盟国ごとにサイバーセキュリティ政策が分かれているため，統一された欧州のデジタル空間を構築できていないことも課題とされました。

**5 大規模プラットフォームによる市場支配**

大手プラットフォーム企業は，検索，広告，電子商取引，クラウド，ヘルスケアなど複数の分野へ事業を拡大しています。こうした企業がデータ，顧客基盤，インフラを横断的に支配すると，新興企業が成長しにくくなり，競争が弱まる可能性があります。

また，大企業による新興テクノロジー企業の買収が，潜在的な競争相手を早期に排除する，いわゆる「キラー買収」になる可能性も問題視されています。
{{< /div-ai >}}

ちなみに生成 AI が世界を席巻する前の時点だ。

さらに報告書では EU におけるデジタル主権を強化するための政策として「データ基盤の構築」「信頼できるデジタル環境」「競争政策と規制の適応」の3つの柱を挙げている。
こうやって眺めると欧州が [Eurosky](https://eurosky.tech/ "Eurosky – mu is here. The first of a thousand social apps.") を含む atproto エコシステムに積極的なのも頷けるよねぇ。

## 今こそデジタル主権について考えよう

ICC の件に関連して，欧州でメールサービス，カレンダー，クラウドストレージを提供しているドイツの [Tuta] 社が以下のブログ記事を公開している。

- [なぜ今こそ、デジタル主権に関する枠組み協定を結ぶのに最適な時期なのか！ | Tuta](https://tuta.com/ja/blog/framework-digital-sovereignty)

記事ではドイツ政府と Microsoft との関係を批判しているが

{{< fig-quote type="markdown" title="なぜ今こそ、デジタル主権に関する枠組み協定を結ぶのに最適な時期なのか！" link="https://tuta.com/ja/blog/framework-digital-sovereignty" >}}
2025年、ドイツ連邦政府はマイクロソフトのライセンスに総額5億ユーロを費やした。 これはマイクロソフトのみ、しかも連邦政府だけの話である。マイクロソフト365の価格は、ライセンスの種類に応じて今年7月から5％から43％値上げされたため、これらの費用は今後も上昇し続けるだろう。

その資金を、欧州製のソリューションへの投資に充てれば、より有効に活用できたはずです。
{{< /fig-quote >}}

まぁ，ここから自社製品を勧めてくるのがにんともかんともだけど（笑） 記事の最後に

{{< fig-quote type="markdown" title="なぜ今こそ、デジタル主権に関する枠組み協定を結ぶのに最適な時期なのか！" link="https://tuta.com/ja/blog/framework-digital-sovereignty" >}}
デジタル主権のための枠組み協定が必要かどうかという問題ではなく、いつそれが実現するかという問題です。なぜなら、今日において米国のITソフトウェアを選択することは、それに伴う地政学的リスクを受け入れるという意識的な決断に他ならないからです。
{{< /fig-quote >}}

と煽っている。
まぁ，でも，今や米国がリスクであることは国際的な認識なのかも知れないなぁ。

ちなみに私はスイスの [Proton] 社のサービスを利用している。

## 【おまけ】 非営利団体を見放す Microsoft

[Proton] 社といえば，こんな記事が上がってた。

- [Nonprofits report data loss after Microsoft software grant ends | Proton](https://proton.me/business/blog/microsoft-nonprofit-data-deletion)

ちなみにこれも広告記事なので注意（笑）

この記事によると Microsoft は非営利団体に対して OneDrive を含む Office 365 のフルセットを無償提供していたそうなのだが，これの終了により、小規模組織は数年分のデータからロックアウトされる事態になったらしい。
しかも Microsoft からの警告はほとんどなかったそうな。

{{< fig-quote type="markdown" title="Nonprofits report data loss after Microsoft software grant ends" link="https://proton.me/business/blog/microsoft-nonprofit-data-deletion" lang="en" >}}
Ronald Khosla, who runs the environmental nonprofit Canopy, got an email in May 2025 warning that his organization’s Microsoft grant would be discontinued and his license would expire that October. When he renewed that October, though, Microsoft’s confirmation email said nothing about the phaseout, and told him he would retain access until October 4, 2026. He logged in that June to find years of files gone anyway. A Microsoft representative later told him roughly 171,000 nonprofits had lost their [OneDrive](https://proton.me/business/drive/onedrive-alternative) data the same way, a figure relayed to Khosla by phone.
{{< /fig-quote >}}

今回は非営利団体だけど，そこに限るわけじゃない。
記事では，上場している欧州企業の74%以上が Microsoft や Google などの米国の技術に依存していると警告している。

{{< fig-quote type="markdown" title="Nonprofits report data loss after Microsoft software grant ends" link="https://proton.me/business/blog/microsoft-nonprofit-data-deletion" lang="en" >}}
When that vendor changed course, the data went with it. A free grant is not a guarantee; it’s a product decision, reversible at any time, with an organization’s data as collateral.

The dependency problem isn’t unique to nonprofits either, as over 74% of publicly listed [European companies depend on US tech](https://proton.me/business/europe-tech-watch) like Microsoft and Google.
{{< /fig-quote >}}

前節と同じく，ここでも「うちの製品を使いなよ」となるわけだが，欧州企業だから安心ってわけでもないよね。
ただ XaaS に関しては「無料より高いものはない」って状態になりつつあるので，もしものときのリスクが最小限になるようリソース配分とコスト配分を考えていかないといけないだろう。

ひるがえって，自身のデジタル主権を考える際に，サービスプロバイダとして日本の企業じゃなくて [Tuta] や [Proton] のような欧州企業を真っ先に連想してしまうのは，我ながら「なんだかなぁ」と思ってしまう。

米国の政変（もう政変と言ってしまおう）をきっかけに国際政治的には色々と変化が起きているが，個人のレベルに落とし込んだときに，デジタル主権を背負うってのは簡単なことじゃないとは思うよ。

## ブックマーク

[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"
[Tuta]: https://tuta.com/ "Tuta: Turn ON privacy for free with secure emails, calendars & contacts | Tuta"
[Proton]: https://proton.me "Proton: Privacy by default"

- [秋 NAS は俺に喰わせろ！]({{< ref "/remark/2021/10/nas.md" >}})

## 参考

{{< linkcard "0956cf258b8a07a0dd78159aa90718306b553b51" >}} <!-- 江河の如く 孫子物語 杜康潤 https://www.kadokawa.co.jp/product/301305000217/ -->
