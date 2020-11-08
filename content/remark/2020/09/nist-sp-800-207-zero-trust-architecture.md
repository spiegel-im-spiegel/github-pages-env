+++
title = "NIST SP 800-207: “Zero Trust Architecture”"
date =  "2020-09-17T12:25:07+09:00"
description = "Refactoring することを前提としたシステム設計が大事。"
image = "/images/attention/kitten.jpg"
tags = [ "nist", "trust", "security", "risk", "management" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

毎度の言い訳だが（笑），私がネットワーク管理者でセキュリティ管理者だったのは遥か昔の牧歌的な時代であり，現在時点で参加・所属する企業・組織のポリシーについてどうこう言う権限はないし，その気もない。
ただし，自衛のための知識を摂取し続けることは必要だと思っている。

というわけで，今回は2020年8月に最終版が公開された NIST [SP 800-207] の触りの部分を覚え書きとして記しておく。

- [SP 800-207, Zero Trust Architecture | CSRC][SP 800-207]

## Zero Trust および Zero Trust Architecture の定義

[SP 800-207] では Zero Trust および Zero Trust Architecture は以下のように定義づけられている。

{{< fig-quote type="markdown" title="SP 800-207: Zero Trust Architecture" link="https://csrc.nist.gov/publications/detail/sp/800-207/final" lang="en" >}}
{{% quote %}}<i>Zero trust</i> (ZT) provides a collection of concepts and ideas designed to minimize uncertainty in enforcing accurate, least privilege per-request access decisions in information systems and services in the face of a network viewed as compromised. <i>Zero trust architecture</i> (ZTA) is an enterprise’s cybersecurity plan that utilizes zero trust concepts and encompasses component relationships, workflow planning, and access policies. Therefore, a zero trust enterpriseis the network infrastructure (physical and virtual) and operational policies that are in place for an enterprise as a product of a zero trust architecture plan{{% /quote %}}.
{{< /fig-quote >}}

ポイントは，アクセスを行うリソース，権限，期間が最小となるよう設計することだ。
これは，特定の安全地帯に入る許可さえあれば，内部のリソースにラフにアクセスできる従来の境界型セキュリティ設計とは大きく異なっている。

ちなみに [SP 800-207] では，アクセスする対象を {{% quote lang="en" %}}resource{{% /quote %}} と呼んでいる。
これは単なるデータだけではなく物理的なデバイスも対象となっていることを示す。
つまり ([日本では既に幻滅期に入っている](https://www.gartner.com/jp/newsroom/press-releases/pr-20200910 "ガートナー、「日本における未来志向型インフラ・テクノロジのハイプ・サイクル：2020年」を発表")) IoT も視野に入っているわけだ。

さらに [SP 800-207] では，アクセスを行う主体を {{% quote lang="en" %}}subjects{{% /quote %}} と呼んでいる。
そう呼ぶからには subjects が指すのは人間（ユーザ）だけではなく，アプリケーション，サービス，デバイス等も含まれる。
また subjects は {{% quote lang="en" %}}authorized and approved subjects{{% /quote %}} と {{% quote lang="en" %}}all other subjects{{% /quote %}} で色分けされている。
もちろん {{% quote lang="en" %}}all other subjects{{% /quote %}} の代表は「攻撃者（attackers）」である。

つまり，あるリソースに対して認証・承認されない actor は，システム上の役割に関わらず，**全て敵**である（笑） この辺が「ゼロトラスト」と呼ばれる所以なのだろう。

ZT/ZTA が重視される理由としては以下の2つが挙げられると思う。

1. 企業・組織への攻撃が巧妙化していて，セキュリティ管理の比重が防御から監視へシフトした
2. クラウド上の XaaS リソースは「境界型」では管理できない

できれば安直にクラウドに繋がろうとするスマート家電もなんとかしてほしいのだが...

## Zero Trust Architecture の基本理念

[SP 800-207] では ZTA の基本理念として，以下の7つの項目を挙げている。

{{< fig-quote type="markdown" title="SP 800-207: Zero Trust Architecture" link="https://csrc.nist.gov/publications/detail/sp/800-207/final" lang="en" >}}
1. All data sources and computing services are considered resources.
2. All communication is secured regardless of network location.
3. Access to individual enterprise resources is granted on a per-session basis.
4. Access to resources is determined by dynamic policy—including the observable state of client identity, application/service, and the requesting asset—and may include other behavioral and environmental attributes.
5. The enterprise monitors and measures the integrity and security posture of all owned and associated assets.
6. All resource authentication and authorization are dynamic and strictly enforced before access is allowed. 
7. The enterprise collects as much information as possible about the current state of assets, network infrastructure and communications and uses it to improve its security posture.
{{< /fig-quote >}}

[@IT の記事](https://www.atmarkit.co.jp/ait/articles/2009/15/news007.html "NISTによる「ゼロトラストにおける7つの基本原則」と従来の境界型防御との関係")にこの7項目の抄訳が載ってたので以下に挙げておく（ちなみに私は @IT のアカウントを持ってないので記事自体は読んでない`w`）。

{{< fig-quote type="markdown" title="NISTによる「ゼロトラストにおける7つの基本原則」と従来の境界型防御との関係：働き方改革時代の「ゼロトラスト」セキュリティ（6） - ＠IT" link="https://www.atmarkit.co.jp/ait/articles/2009/15/news007.html" >}}
1. データソースとコンピュータサービスは、全てリソースと見なす
2. 「ネットワークの場所」に関係なく、通信は全て保護される
3. 組織のリソースへのアクセスは、全て個別のセッションごとに許可される
4. リソースへのアクセスは動的なポリシーによって決定される
5. 組織が保有するデバイスは、全て正しくセキュリティが保たれているように継続的に監視する
6. リソースの認証と認可は、全てアクセスが許可される前に動的かつ厳密に実施される
7. 資産・ネットワーク・通信の状態について可能な限り多くの情報を収集し、セキュリティを高めるために利用する
{{< /fig-quote >}}

面白いのは ZTA に最初から「監視」が組み込まれていること，常に状況をフィードバックして「改善」のサイクルを構築するところまでがセットになっていることだろう。

セキュリティに於いても PDCA サイクル，いや今なら OODA ループか，が重要ということやね（笑）

## 大変なのは...

ZT を組み込むこと自体は，そう難しくないだろう。
おそらくは既存のシステムに ZT の仕組みをラッピングすることで構成可能なはずだ。

{{< fig-img-quote src="./zero-trust-access.png" title="SP 800-207: Zero Trust Architecture" link="https://csrc.nist.gov/publications/detail/sp/800-207/final" lang="en" width="1183" >}}

大変なのは ZTA におけるリソースとサブジェクト（の権限）の定義・運用・評価だろう。
これ，かなり細かい要求分析が必要だと思うよ。

当然ながら人間組織の役職で権限を決めるわけにはいかない。
システム管理者やセキュリティ管理者（セキュリティ企業も含めて）であっても「アクセスしてはいけないリソース」はある。
サブジェクトやリソースの杜撰な管理で [Class Break を引き起こした Twitter]({{< ref "/remark/2020/07/class-breaks-from-twitter.md" >}} "Twitter から始まる Class Break") の事例は耳に新しいだろう。
日本での最近の Class Break 事例は「ドコモロ系事案[^d1]」か（笑）

[^d1]: [キャッシュレス決済を使った不正利用に関する一連のインシデント](https://piyolog.hatenadiary.jp/entry/2020/09/16/064653 "不正利用が発生した電子決済サービスについてまとめてみた - piyolog")のこと。 Facebook の TL で見かけた「ドコモロ系事案」のフレーズが面白かったので使ってみた（笑）

だからこそループを回して「改善」していかなければならないんだろうけど。
Refactoring することを前提としたシステム設計が大事。

## ブックマーク

- [Lightning Q&A: DevSecOps in five with Maya Kaczorowski - The GitHub Blog](https://github.blog/2020-09-24-lightning-qa-devsecops-in-five-with-maya-kaczorowski/)
- [【超図解】ゼロトラスト: NECセキュリティブログ | NEC](https://jpn.nec.com/cybersecurity/blog/201016/index.html)

[SP 800-207]: https://csrc.nist.gov/publications/detail/sp/800-207/final "SP 800-207, Zero Trust Architecture | CSRC"

## 参考図書

{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
{{% review-paapi "4757143044" %}} <!-- 信頼と裏切りの社会 -->
{{% review-paapi "B07ND6QTN4" %}} <!-- OODA LOOP -->
