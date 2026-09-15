+++
title = "昔から「ソフトウェアのバージョンは最新にすべし」と言われたけど..."
date =  "2026-09-15T19:58:13+09:00"
description = "攻撃側が AI 支援型攻撃で迅速に脆弱性を突いてくるのであれば，防御側もそのスピードに対応しなければならない。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "security", "risk", "management", "artificial-intelligence", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

今回は Bruce Schneier 氏の[記事](https://www.schneier.com/blog/archives/2026/09/microsofts-patching.html "Microsoft's Patching - Schneier on Security")経由で，以下を起点に話を始めようか。

{{< linkcard "bf8c2b1da340d9c0e6fb2a5972bd0253dab6a6db" >}} <!-- https://arstechnica.com/security/2026/09/microsoft-patches-a-record-972-vulnerabilities-112-of-them-critical/ Why this month's Microsoft patch release is a doozy -->

先週の，いわゆる Patch Tuesday で Microsoft Update による脆弱性の修正が約972件にのぼり，そのうちの112件が深刻なものとして計上されている。
年単位の集計でも，今年の報告数が（現時点で既に）ダントツらしい。

{{< fig-img-quote src="./Picture1.webp" title="Zero Day Initiative — The September 2026 Security Update Review" link="https://www.zerodayinitiative.com/blog/2026/9/8/the-september-2026-security-update-review" width="1500" lang="en" >}}

こうした脆弱性報告の急増は，ソフトウェアが急にポンコツになったわけではなく，最近の「AI 支援型攻撃（AI-assisted attacks）」が関係しているようだ。

{{< fig-quote type="markdown" title="Why this month's Microsoft patch release is a doozy" link="https://arstechnica.com/security/2026/09/microsoft-patches-a-record-972-vulnerabilities-112-of-them-critical/" lang="en" >}}
Two weeks ago, OpenAI, Anthropic, Amazon Web Services, Google, Microsoft, and 100 companies and organizations published an [open letter](https://openai.com/collective-cyberdefense) warning of a narrowing window for patching vulnerabilities ahead of an expected tsunami of AI-enabled attacks that actively exploit them first. The industry is taking the threat seriously by pumping out unprecedented numbers of patches in their software.
{{< /fig-quote >}}

攻撃される前に穴を塞ぐのは正しい。
問題は発見する穴の数と穴を塞ぐまでの猶予期間だろう。
当然ながら（規模的に）人間の手には負えないので，防御する方も AI による支援を受けて対応する。
これが脆弱性報告が急増した理由らしい。

Bruce Schneier 氏はこれを

{{< fig-quote type="markdown" title="Microsoft's Patching - Schneier on Security" link="https://www.schneier.com/blog/archives/2026/09/microsofts-patching.html" lang="en" >}}
This is the result of AI-powered vulnerability finding, and a good example of AI helping the defenders more than the attackers.
{{< /fig-quote >}}

と肯定的に評価している。

一方で，攻撃側はパッチ情報を解析して利用できる。
これも AI により効率化されていると思われる。

{{< fig-quote type="markdown" title="Microsoft's Patching - Schneier on Security" link="https://www.schneier.com/blog/archives/2026/09/microsofts-patching.html" lang="en" >}}
And Microsoft is right: The window to patch has shrunk to “immediately.” AIs are also good at reverse-engineering exploits from patches, which means that these vulnerabilities will be weaponized as soon as the update is published.
{{< /fig-quote >}}

個人ユーザはともかく，組織的に Windows を利用している場合は，パッチが出てもすぐに適用できないことが多い。
しかし，攻撃側が AI 支援型攻撃で迅速に脆弱性を突いてくるのであれば，防御側もそのスピードに対応しなければならない。

[Kagi Assistant] に件の記事を読み込ませたところ，以下を提言してきた。

{{< div-ai type="markdown" >}}
- パッチ公開後の検証期間を短縮する必要がある。
- インターネットから到達可能なシステムを最優先で更新する必要がある。
- パッチ適用が間に合わない場合は，サービス停止，アクセス制限，仮想パッチ，監視強化などの代替策が必要になる。
- 「深刻度が高いか」だけでなく，「パッチ公開後に攻撃コードが作られやすいか」を考慮する必要がある。
- パッチ適用を定期的な保守作業ではなく，攻撃開始前の緊急対応として扱う必要がある。
{{< /div-ai >}}

情シスな方々は，この状況をどう考えているだろうか。
一朝一夕にはいかないと思うのだけど，先延ばしにできる話でもないだろう。

ここからは余談。

最近 Anthropic 社の CEO が AI 開発スピードを減速すべきとか言ってるらしいが（そして米国の大手 IT 企業が賛同している）

- [AI開発「減速」をアンソロピックCEOが提案　アルトマン氏・マスク氏らも賛同](https://www.watch.impress.co.jp/docs/news/2140440.html)

ぶっちゃけ「おまえがゆーな！」という感じである。
Anthropic 社をはじめとする米国大手 AI 企業に対しては，まず自社の AI サービスのやらかしに対してちゃんとペナルティを支払わせろよ。
話はそれからだっちうの。

しかもこの「提案」に対してかの大統領は

{{< fig-quote type="markdown" title="OpenAIやAnthropicトップが異例の「AI減速」支持　トランプ氏は拒否、中国との競争を優先 - 週刊アスキー" link="https://weekly.ascii.jp/elem/000/004/434/4434468/" >}}
一方、トランプ大統領はロイター通信などの取材に対して、「多くの非常に否定的な勢力が、実現しないであろうAIについての事柄を持ち出している」と主張。中国とのAI開発競争を踏まえ、米国が同分野で先行する状況を維持したいとの考えを示した。
{{< /fig-quote >}}

などと言ってるらしい。
つまり自国企業の野放図な状況を放置する気満々ということである。

ホンマ，この「ならずもの国家[^rs1]」をどうにかしてくれよ！ つか，とっとと AI バブル弾けてくれ！ [グラボが100万円](https://automaton-media.com/articles/newsjp/20260911-466913/ "グラボ高騰止まらず、「NVIDIA GeForce RTX 5090」はついに100万円を超え始める。“さらなるグラボ値上げ”を予告するショップも - AUTOMATON")とかありえねーよ。

[^rs1]: ならずもの国家（rogue state）は1990年代の米国クリントン政権時代に用いられた表現で，北朝鮮やイランなど7ヶ国を指していた。後にこれは子ブッシュ政権（9.11 のときの政権）で「悪の枢軸」などと再定義されている。[日本版 Wikipedia の記事](https://ja.wikipedia.org/wiki/%E3%81%AA%E3%82%89%E3%81%9A%E3%82%82%E3%81%AE%E5%9B%BD%E5%AE%B6 "ならずもの国家 - Wikipedia")には「日本の外務省は違法国家あるいは無責任国家と意訳している」とか書いてある。ホンマか？ クレイジーキャッツやがな（笑） まぁ「無責任国家」というなら，間違いなく米国現政権は rogue state だよな。

[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"

## 参考

{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
{{< linkcard "f49db55e98f0eb56c864acd2ee6f4da8da10016c" >}} <!-- ハッキング思考 -->
{{< linkcard "955ecd76b8158e21c4262c73717f0c1bc2d94351" >}} <!-- ケアレス・ピープル -->
