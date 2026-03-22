+++
title = "令状の外側で動く位置情報監視"
date =  "2026-03-22T11:31:39+09:00"
description = "「力なき正義は悪の前に無力」かもしれないが，正義なき力は悪そのものである。米国はどこまでもダークサイドへ堕ちていく。日本もきっと追従するのだろう。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "security", "privacy", "risk", "surveillance-capitalism", "tracking", "politics" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

今回は先日見かけた [Proton blog] の記事を起点にして書いてみる[^p1]。

[^p1]: [Proton blog] はなべて広告記事なので，読むときにはその辺を割り引いて読む必要がある。

- [The FBI is buying location data to track people. Here’s how data brokers made it possible. | Proton](https://proton.me/blog/fbi-location-data)

とりあえず [Kagi Assistant] に要約させてみた。こんな感じ：

{{< div-box type="markdown" >}}
- **FBIによる位置データの購入**
    - FBIがデータブローカーからアメリカ国民の位置情報を購入していることを認めました。
    - この手法により、政府は従来の令状（捜査令状）を必要とする手続きを経ることなく、機密性の高い情報にアクセスが可能になります。
- **議会での証言と批判**
    - FBI長官のカッシュ・パテル氏が上院情報委員会の公聴会にて、捜査のために「市販されている情報」を購入していることを認めました。
    - これに対し、ロン・ワイデン上院議員などの議員からは、憲法上の重大な懸念があるとして即座に批判の声が上がっています。
- **背景**
    - FBIがこの種のデータを積極的に購入していることが確認されたのは、2023年以来初めてのことです。
{{< /div-box >}}

国によるのだろうが，警察組織による捜査目的であれ，建前上はインターネット上のサービスにあるユーザの行動データを取得するには裁判所から事前に令状を取る必要がある。
これは日本でも（手続きの違いはあれ）同じ（筈）である。

ところが [POLITICO の記事](https://www.politico.com/news/2026/03/18/fbi-buying-data-track-people-patel-00834080 "FBI is buying data that can be used to track people, Patel says - POLITICO")によると FBI はデータブローカーから位置情報を購入していることを認めたそうだ。
FBI の言い分は「市販データを買うただけやし，ええやろ（←超意訳）」というものらしい。
しかも上院情報委員会の Tom Cotton 議長（共和党・アーカンソー州）もこれを追認する姿勢を見せているようだ。

{{< fig-quote type="markdown" title="FBI is buying data that can be used to track people, Patel says - POLITICO" link="https://www.politico.com/news/2026/03/18/fbi-buying-data-track-people-patel-00834080" lang="en" >}}
Committee Chair Tom Cotton (R-Ark.) defended the practice at the hearing.

“The key words are commercially available. If any other person can buy it, and the FBI can buy it, and it helps them locate a depraved child molester or savage cartel leader, I would certainly hope the FBI is doing anything it can to keep Americans safe,” he said.
{{< /fig-quote >}}

さすが共和党！ おれたちにできない事を平然とやってのけるッ シビれもあこがれもしないが（笑） さらにさらに国防情報局の長官も便乗したのか

{{< fig-quote type="markdown" title="FBI is buying data that can be used to track people, Patel says - POLITICO" link="https://www.politico.com/news/2026/03/18/fbi-buying-data-track-people-patel-00834080" lang="en" >}}
Defense Intelligence Agency Director James Adams told senators at the hearing that his agency also purchases commercially available information.
{{< /fig-quote >}}

とか言い出す始末。

いわゆるオープンソース・インテリジェンス（OSINT）とは公開されているデータを収集・分析する技術で，諜報機関などでは大昔からやっているらしい。
市販されているデータなら「オープンソース」だから OSINT の対象だ，ということなのだろう。

[件の Proton 記事][件の記事]ではこうした情報流通の要としてデータブローカーの存在を挙げている。

{{< fig-quote type="markdown" title="The FBI is buying location data to track people. Here’s how data brokers made it possible." link="https://proton.me/blog/fbi-location-data" lang="en" >}}
A data broker gathers information from apps, websites, and third-party partners. Location data is a central part of that system, often collected through [routine app permissions](https://proton.me/blog/ad-tech-privacy).

That information is combined with other signals such as browsing activity, purchases, and inferred interests. The result is a detailed profile that can be sold to a wide range of buyers, including government agencies.
{{< /fig-quote >}}

さらに生成 AI の台頭が事態を深刻化している。

{{< fig-quote type="markdown" title="The FBI is buying location data to track people. Here’s how data brokers made it possible." link="https://proton.me/blog/fbi-location-data" lang="en" >}}
This data is used beyond surveillance. It can shape advertising and influence political messaging in ways that [undermine democracy](https://proton.me/blog/data-brokers-democracy). These datasets continue to expand and are increasingly [analyzed using AI](https://proton.me/blog/data-brokers-ai), which makes it easier to cross-reference data and uncover deeper patterns about individuals, amplifying existing biases and enabling more precise manipulation at scale.
{{< /fig-quote >}}

異なるドメインの情報が AI によって結び付けられ「データの相互参照や個人に関するより深いパターンの発見が容易になる」わけだ。
これらの情報の利用は犯罪者等の追跡に限らない。
高度情報社会で重要な関心経済（attention economy）において優位に立つための強力な（特に政治的な）武器になる。

令状を迂回して市販の情報を買うというのは『[ハッキング思考]』で言うところの「法システムのハッキング」というやつだ。

{{< fig-quote type="markdown" title="『ハッキング思考』29章 ハッキングと権力" link="https://www.amazon.co.jp/dp/B0CK19L1HC?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
ハッキングとは以前からの権力に対して優位に立とうとして敗者の側がしかけるものと考えがちだ。だが、権力者が自らの優位をさらに強化しようとして手がけるほうが、実はずっと多い。
{{< /fig-quote >}}

これを権力者でもない「敗者」の我々が防ぐのは容易ではない。
インターネット上のサービスは日常生活に深く浸透していて簡単に排除できるものではない。
[Proton 記事][件の記事]では消極策としながら以下の手段を挙げている。

- アプリの権限、特に位置情報へのアクセスを制限する
- 使用していないアプリを削除する
- トラッキングに依存するサービスを避ける
- VPN を利用して IP アドレスを隠しインターネットトラフィックを暗号化する

最後の VPN を利用するやり方は，昨今ではあまり推奨されない。
VPN プロバイダがユーザの情報をデータブローカーに売ったり，所属する国からの要請を断れず情報を提出するケースが見られるからだ（特に無料の VPN サービス）。
Proton はユーザのプライバシーを守る製品構成を売りにしているが，それでも所属する国からの正規の要請は断れない。

- [プライバシー重視をうたうメールサービス「Proton Mail」が個人情報をスイス政府に提供、FBIの手にまで渡ったことが判明 - GIGAZINE](https://gigazine.net/news/20260306-proton-mail-helped-fbi/)

「[力なき正義は悪の前に無力](https://www.afpbb.com/articles/-/3627611 "ネタニヤフ氏、「キリストはチンギスハンに劣る」と発言 力なき正義は悪の前に無力と主張　写真10枚　国際ニュース：AFPBB News")」かもしれないが，正義なき力は悪そのものである。
米国はどこまでもダークサイドへ堕ちていく。
日本もきっと追従するのだろう。

## ブックマーク

- [FBIはアメリカ国民の位置情報データを購入していると長官が発言 - GIGAZINE](https://gigazine.net/news/20260321-fbi-location-data/)

[Proton blog]: https://proton.me/blog "The Proton Blog - News from the front lines of privacy and security | Proton"
[件の記事]: https://proton.me/blog/fbi-location-data "The FBI is buying location data to track people. Here’s how data brokers made it possible. | Proton"
[Kagi Search]: https://kagi.com/ "Kagi Search - A Premium Search Engine"
[Kagi Translate]: https://translate.kagi.com/ "Kagi Translate"
[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"
[ハッキング思考]: https://www.amazon.co.jp/dp/B0CK19L1HC?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Amazon.co.jp: ハッキング思考　強者はいかにしてルールを歪めるのか、それを正すにはどうしたらいいのか eBook : ブルース・シュナイアー, 高橋 聡: Kindleストア"

## 参考

{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
{{% review-paapi "B01MZGVHOA" %}} <!-- 超監視社会 -->
{{% review-paapi "B0CK19L1HC" %}} <!-- ハッキング思考 Kindle 版 -->
