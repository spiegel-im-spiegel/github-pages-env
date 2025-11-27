+++
title = "「クローリング課金」問題の概要"
date =  "2025-11-27T07:49:26+09:00"
description = "description"
image = "/images/attention/cc-logo.png"
tags = ["creative-commons", "copyright", "artificial-intelligence", "web"]
pageType = "text"
license = "by"
draft = true

[scripts]
  mathjax = false
  mermaidjs = false
+++

この記事は [Creative Commons] が2025年11月に公開した PDF 文書 {{% quote lang="en" %}}{{< pdf-file title="Pay-To-Crawl Issue Brief" link="https://creativecommons.org/wp-content/uploads/2025/11/Pay-to-Crawl-Issue-Brief-Nov-2025.pdf" >}}{{% /quote %}} を試みに翻訳したものです。
原文は Jack Hardinges 氏により [CC-BY {{< cc-syms "cc" "by" >}} 4.0](https://creativecommons.org/licenses/by/4.0/) ライセンス下で公開されています。

翻訳に際して [GitHub Copilot] (model: GPT-5 mini) および [Kagi Transrate] の助けを多く借りています。
なお，翻訳に関する間違いは全て翻訳者である [Spiegel](https://baldanders.info/profile/) に帰します。
もし間違いがあればフィードバックして頂けると助かります。

----

<!-- ## Pay-To-Crawl

_A Creative Commons Issue Brief: Backgrounders on topics related to AI & the commons_

_November 2025_ -->

## クローリング課金（Pay‑To‑Crawl）

_Creative Commons Issue Brief： AI とコモンズに関連するトピックの背景説明_

_2025年11月_

<!-- ### Introduction

While machine access of digital content is not entirely new, pay-to-crawl systems have emerged in response to [the disruption caused by large artificial intelligence (AI) models accessing vast amounts of content](https://creativecommons.org/wp-content/uploads/2025/06/Human-Content-to-Machine-Data_Final.pdf) without permission, attribution, or compensation.

Pay-to-crawl systems are described as addressing issues including increased hosting costs for websites, reductions in traffic and visibility brought about by AI-enabled search, and the undermining of referral and advertising-based business models.

Viewed through a wider lens, pay-to-crawl systems represent one of the latest incarnations of content monetization and control, combining elements of paywall, digital rights management (DRM), and micropayment approaches. -->

### はじめに

機械（machine）によるデジタルコンテンツへのアクセス自体は全く新しいものではありませんが，大規模な人工知能（AI）モデルが許可や帰属の明記，対価の支払いなしに[膨大な量のコンテンツにアクセスすることで生じる混乱](https://creativecommons.org/wp-content/uploads/2025/06/Human-Content-to-Machine-Data_Final.pdf)に対応するため，クローリング課金（pay-to-crawl）システムが登場しました。

クローリング課金システムは，ウェブサイト（website）のホスティング費用の増加， AI 検索によるトラフィックと可視性の低下，そして紹介や広告ベースのビジネスモデルの崩壊といった問題に対処するものとして説明されています。


より広い視野で見れば，クローリング課金システムは，ペイウォール，デジタル著作権管理（DRM），そしてマイクロペイメントのアプローチの要素を組み合わせた，コンテンツの収益化と管理における最新の形態のひとつであると言えます。

### How Pay-To-Crawl Systems Work

Not all pay-to-crawl systems work in the same way. Some systems, for example, are focused more on blocking machines from accessing content than making them pay. However, most systems tend to involve some combination of the following components:

- **Authentication:** Pay-to-crawl systems require the person, organization, or product operating a machine seeking to access content to identify themselves. Many systems use cryptographic authentication rather than methods that operators of machines have proven able to circumvent in the past, such as user agent strings or IP addresses.
- **Access Control:** Pay-to-crawl systems set granular and functional rules to define which machines can access content, under what conditions, and whether access is freely enabled, blocked, or billed. Some systems allow websites to set rate limits, rather than fully block off access.
- **Pricing & Contracting:** Pay-to-crawl systems define compensation for access, such as per page or by data volume, or on a subscription basis. Contracts are generally automated, sometimes using standardized licenses or terms. Terms are not only or always financial—they can involve attribution and other reuse obligations. Some systems enable collective bargaining on behalf of groups of websites.

### クローリング課金システムの仕組み

すべてのクローリング課金システムが同じ方式で動作するわけではありません。例えば，システムによっては，機械に料金を支払わせることよりも，コンテンツへのアクセスをブロックすることに重点を置いているものもあります。しかし，大抵のシステムは以下のような要素を組み合わせて運用される傾向があります。

- **認証（Authentication）：** クローリング課金システムは，コンテンツにアクセスしようとする機械を操作する個人・組織・製品を識別することを要求します。多くのシステムは，ユーザーエージェントやIPアドレスのように回避されやすい方法ではなく，暗号化された認証を用いています。
- **アクセス制御（Access Control）：** クローリング課金システムは，コンテンツにアクセスできる条件（無料で許可するか，ブロックするか，課金するか，といった詳細かつ機能的なルール）を機械ごとに定義します。システムによっては，ウェブサイト側がアクセスを完全に遮断するのではなく，アクセス頻度の上限（レートリミット）を設定できるようになっているものもあります。

- **価格設定と契約（Pricing & Contracting）：** クローリング課金システムは，ページ単位やデータ量による課金，あるいはサブスクリプションなど，アクセスに対する対価を定義します。契約は一般的に自動化されており，標準化されたライセンスや規約が用いられることもあります。条件は必ずしも金銭的なものだけでなく，帰属表示やその他の再利用に関する義務を伴う場合があります。ウェブサイトのグループを代表して集団交渉を行うことを可能とするシステムもあります。














[著作権法]: https://elaws.e-gov.go.jp/document?lawid=345AC0000000048 "著作権法"
[Creative Commons]: https://creativecommons.org/ "Creative Commons"
[CC Licenses]: https://creativecommons.org/licenses/ "ライセンスについて - Creative Commons"
[GitHub Copilot]: https://github.com/features/copilot "GitHub Copilot · Your AI pair programmer · GitHub"
[Kagi Transrate]: https://translate.kagi.com/ "Kagi Transrate"
