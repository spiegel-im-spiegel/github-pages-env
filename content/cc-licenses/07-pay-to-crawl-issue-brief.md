+++
title = "「クローリング課金」問題の概要"
date =  "2025-11-27T19:54:53+09:00"
description = "この記事は Creative Commons が2025年11月に公開した PDF 文書 ”Pay-To-Crawl Issue Brief” を試みに翻訳したものです。"
image = "/images/attention/cc-logo.png"
tags = ["creative-commons", "copyright", "artificial-intelligence", "web"]
pageType = "text"
license = "by"

[scripts]
  mathjax = false
  mermaidjs = false
+++

この記事は [Creative Commons] が2025年11月に公開した PDF 文書 {{% quote lang="en" %}}{{< pdf-file title="Pay-To-Crawl Issue Brief" link="https://creativecommons.org/wp-content/uploads/2025/11/Pay-to-Crawl-Issue-Brief-Nov-2025.pdf" >}}{{% /quote %}} を試みに翻訳したものです。
原文は Jack Hardinges 氏により [CC-BY {{< cc-syms "cc" "by" >}} 4.0](https://creativecommons.org/licenses/by/4.0/) ライセンス下で公開されています。

翻訳に際して [GitHub Copilot] (model: GPT-5 mini) および [Kagi Transrate] の助けを多く借りています。
なお，翻訳に関する間違いは全て翻訳者である [Spiegel](https://baldanders.info/profile/) に帰します。
もし間違いがあればフィードバックして頂けると助かります。

では，本編をどうぞ。

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

<!-- ### How Pay-To-Crawl Systems Work

Not all pay-to-crawl systems work in the same way. Some systems, for example, are focused more on blocking machines from accessing content than making them pay. However, most systems tend to involve some combination of the following components:

- **Authentication:** Pay-to-crawl systems require the person, organization, or product operating a machine seeking to access content to identify themselves. Many systems use cryptographic authentication rather than methods that operators of machines have proven able to circumvent in the past, such as user agent strings or IP addresses.
- **Access Control:** Pay-to-crawl systems set granular and functional rules to define which machines can access content, under what conditions, and whether access is freely enabled, blocked, or billed. Some systems allow websites to set rate limits, rather than fully block off access.
- **Pricing & Contracting:** Pay-to-crawl systems define compensation for access, such as per page or by data volume, or on a subscription basis. Contracts are generally automated, sometimes using standardized licenses or terms. Terms are not only or always financial—they can involve attribution and other reuse obligations. Some systems enable collective bargaining on behalf of groups of websites.
- **Payment:** Pay-to-crawl systems provide mechanisms for payment, often using secure third-party processing services. Payment can be made directly and immediately to the website, or taken by an operator of the system, such as a web services provider, on their behalf.
- **Content Delivery:** Upon authentication and payment, pay-to-crawl systems enable access to content, typically in formats optimized for machine consumption. Some systems enable encrypted access to non-public content.
- **Metering & Logging:** Pay-to-crawl systems often log information related to access and use of content to enable billing and some degree of auditability.

The role of websites in developing and using pay-to-crawl systems varies. Some may choose to deploy such a system themselves, using emerging protocols and code. In other cases, pay-to-crawl systems are being developed as specialist, paid-for products, or introduced by web services providers (such as domain hosts and content delivery networks) on behalf of websites. As a result, pay-to-access systems vary in terms of their openness, standardization, and interoperability, as well as the permissiveness of their access controls and payment terms. -->

### クローリング課金システムの仕組み

すべてのクローリング課金システムが同じ方式で動作するわけではありません。
例えば，システムによっては，機械に料金を支払わせることよりも，コンテンツへのアクセスをブロックすることに重点を置いているものもあります。
しかし，大抵のシステムは以下の要素の組み合わせを含む傾向があります：

- **認証：** クローリング課金システムは，コンテンツにアクセスしようとする機械を操作する個人・組織・製品を識別することを要求します。多くのシステムは，ユーザーエージェントや IP アドレスのように回避されやすい方法ではなく，暗号化された認証を用いています。
- **アクセス制御：** クローリング課金システムは，コンテンツにアクセスできる条件（無料で許可するか，ブロックするか，課金するか，といった詳細かつ機能的なルール）を機械ごとに定義します。システムによっては，ウェブサイト側で（アクセスを完全に遮断するのではなく）アクセス頻度の上限を設定できるようになっているものもあります。
- **価格設定と契約：** クローリング課金システムは，ページ単位やデータ量による課金，あるいはサブスクリプションなど，アクセスに対する対価を定義します。契約は一般的に自動化されており，標準化されたライセンスや規約が用いられることもあります。条件は必ずしも金銭的なものだけでなく，帰属表示やその他の再利用に関する義務を伴う場合があります。ウェブサイトのグループを代表して集団交渉を行うことを可能とするシステムもあります。
- **支払い：** クローリング課金システムは，多くの場合，安全な第三者決済サービスを利用した支払いの仕組みを提供します。支払いはウェブサイトに直接かつ即時に行われるか，ウェブサービス事業者のようなシステムの運営者がウェブサイトに代わって受け取ることも可能です。
- **コンテンツ配信：** 認証と支払いが完了すると，クローリング課金システムは（通常の場合）機械での利用（consumption）に最適化された形式でコンテンツへのアクセスを可能にします。非公開コンテンツへの暗号化されたアクセスを可能とするシステムもあります。
- **計測と記録：** 課金とある程度の監査可能性（auditability）のために，クローリング課金システムは，コンテンツへのアクセスと利用に関連する情報を記録することが多いです。

クローリング課金システムの開発と利用におけるウェブサイトの役割は様々です。
一部のシステムでは，新興のプロトコルやコードを自ら展開することを選択するかもしれません。
あるいは，専門的な有料製品として開発されたり，ウェブサービス事業者（ドメインホストやコンテンツ配信ネットワークなど）がウェブサイトに代わって導入する場合もあります。
その結果，アクセス課金（pay-to-access）システムは，公開性，標準化，相互運用性，さらにはアクセス制御の許容度や支払い条件においても差異があります。

<!-- ## Considerations

In the face of unprecedented consumption of digital content by large AI models—both in scale and in impact—the use of pay-to-crawl systems may help websites sustain the creation and publication of content, or tackle what they consider to be substitutive uses of their works. However, overbroad and indiscriminate use of pay-to-crawl systems could block off access to digital content for researchers, nonprofits, cultural heritage institutions, educators, and other actors working in the public interest; obstruct legitimate uses of content protected by copyright or other laws; and create new walled gardens, web gatekeepers, and excesses of power. Wide adoption of pay-to-crawl could ultimately represent a shift away from the spirit of the open web towards a more tightly controlled and monetized content ecosystem. -->

## 考慮事項

大規模な AI モデルによるデジタルコンテンツの消費が，その規模においても影響においても前例のないものとなる中で，クローリング課金システムの利用は，ウェブサイトがコンテンツの制作と公開を維持したり作品の代替的な利用（substitutive uses）と見なされるものに対処したりする助けとなるかもしれません。
しかし，クローリング課金の過度に広範で無差別な利用は，研究者，非営利団体，文化遺産機関，教育者，その他公益のために活動する人々によるデジタルコンテンツへのアクセスを遮断し，著作権法やその他の法律で保護された正当な利用を妨げ，さらに新たなウォールド・ガーデン（walled gardens）やウェブ・ゲートキーパーや権力の濫用を生む可能性があります。
クローリング課金が広く採用されれば，最終的にオープンウェブの精神からより厳格に管理・収益化されたコンテンツエコシステムへのシフトをもたらすかもしれません。

<!-- ## Examples

Examples of pay-to-crawl systems and related initiatives include: [Pay Per Crawl](https://blog.cloudflare.com/introducing-pay-per-crawl/) by Cloudflare, [AI RevShare](https://www.valyu.network/rev-share-partner-programme) by Valyu, [GistAttribution](https://gist.ai/#attribution) by ProRata, [Open Licensing Protocol](https://rslstandard.org/) by RSL and [TollBit](https://tollbit.com/). -->

## 事例

クローリング課金システムや関連する取り組みの例としては Cloudflare の [Pay Per Crawl](https://blog.cloudflare.com/introducing-pay-per-crawl/), Valyu の [AI RevShare](https://www.valyu.network/rev-share-partner-programme), ProRata の [GistAttribution](https://gist.ai/#attribution), RSL の [Open Licensing Protocol](https://rslstandard.org/), そして [TollBit](https://tollbit.com/) などがあります。

<!-- ## Notes on Terminology

We’re choosing to use *pay-to-crawl* to describe these systems on account of the term already being widely used. We generally prefer the broader term *pay-to-access*, given that, technically speaking, there are many purposes and forms of machine access to content beyond crawling. Crawling does not, for example, adequately describe the process of extracting and making copies of content (often referred to as scraping), nor analyzing them to derive insights or patterns (text and data mining).

“Websites” is a broad category. The term *publisher* might be more appropriate to describe the entity responsible for the content and the user of a pay-to-crawl system, especially in domains such as news, academia, and the media. The user of a pay-to-crawl system, regardless of how they are described, is not always the original creator or owner of the content such a system is used to manage access to.

In this context, *machines* refers to the systematic access and use of digital content using code and automated programs, rather than typical human browsing and consumption. *Bots* is also sometimes used. This shouldn’t obscure the fact that code and automated programs are ultimately operated by humans. -->

## 用語に関する注記

用語として既に広く使われているため，私達は「クローリング課金（pay-to-crawl）」をこれらのシステムを表現するために用いています。
技術的に言えば，機械によるコンテンツへのアクセスにはクローリング以外にも多くの目的や形態が存在するため，私達は通常，より広い意味での「アクセス課金（pay-to-access）」を推奨します。
例えば，クローリングという言葉では，コンテンツを抽出して複製するプロセス（しばしばスクレイピングと呼ばれる）や，洞察やパターンを導き出す分析（テキスト・データマイニング）を十分に説明できません。

「ウェブサイト（“Website）」は広範なカテゴリーです。
特にニュース，学術，メディアといった領域においては，コンテンツに責任を負う主体でありクローリング課金システムの利用者でもある存在を説明するには，「パブリッシャー（publisher）」という用語の方がより適切かもしれません。
クローリング課金システムの利用者は，どのように表現されるかに関わらず，そのようなシステムがアクセス管理の対象とするコンテンツの本来の制作者や所有者であるとは限りません。

本稿の文脈において「機械（machine）」はコードや自動化プログラムを用いた体系的なアクセスと利用を指し，人間による典型的な閲覧や利用（consumption）とは区別されます。
「ボット（bot）」という言葉が使われることもあります。
コードや自動化されたプログラムは最終的には人間によって操作されているという事実を曖昧にすべきではありません。

<!-- ## License

This brief by Jack Hardinges is licensed under [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/). -->

## ライセンス

Jack Hardinges による本稿は [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/) の下にライセンスされています。

----

以上， {{% quote lang="en" %}}{{< pdf-file title="Pay-To-Crawl Issue Brief" link="https://creativecommons.org/wp-content/uploads/2025/11/Pay-to-Crawl-Issue-Brief-Nov-2025.pdf" >}}{{% /quote %}} の翻訳でした。

## ブックマーク

- [AI and the Commons: A Reading List - Creative Commons](https://creativecommons.org/2025/09/03/ai-and-the-commons-a-reading-list/)
- [Where CC Stands on Pay-to-Crawl - Creative Commons](https://creativecommons.org/2025/12/12/where-cc-stands-on-pay-to-crawl/) : Creative Commons 名義による公式見解
  - [Creative Commonsの「ペイ・トゥ・クロール（Pay-to-Crawl）」に関する立場（Where CC Stands on Pay-to-Crawl 日本語訳）](https://www.yamdas.org/column/technique/where-cc-stands-on-pay-to-crawlj.html)

- [CC Signals と AI]({{< ref "/remark/2025/06/cc-signals.md" >}})

[著作権法]: https://elaws.e-gov.go.jp/document?lawid=345AC0000000048 "著作権法"
[Creative Commons]: https://creativecommons.org/ "Creative Commons"
[CC Licenses]: https://creativecommons.org/licenses/ "ライセンスについて - Creative Commons"
[GitHub Copilot]: https://github.com/features/copilot "GitHub Copilot · Your AI pair programmer · GitHub"
[Kagi Transrate]: https://translate.kagi.com/ "Kagi Transrate"
