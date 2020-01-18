+++
title = "5G セキュリティ・リスク"
date =  "2020-01-18T21:26:35+09:00"
description = "歴史は繰り返す。 ハムスターの回し車のように…"
image = "/images/attention/kitten.jpg"
tags = [ "security", "rsik", "communication" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

とても残念な話だが，大国間の極めて政治的なやり取りによって 5G のセキュリティ・リスクについて本当のところの議論がなおざりになっている感がある。

というわけで Schneier 先生のエッセイを紹介する。

- [5G Security - Schneier on Security](https://www.schneier.com/blog/archives/2020/01/china_isnt_the_.html)

詳しくはリンク先を見てもらうとして，この記事では Bruce Schneier さんが挙げる3つのセキュリティ問題を挙げておく。

{{< fig-quote type="markdown" title="5G Security" link="https://www.schneier.com/blog/archives/2020/01/china_isnt_the_.html" lang="en" >}}
{{% quote %}}First, the standards are simply too complex to implement securely. This is true for [all software](https://www.schneier.com/essays/archives/1999/11/a_plea_for_simplicit.html), but the 5G protocols offer particular difficulties. Because of how it is designed, the system blurs the wireless portion of the network connecting phones with base stations and the core portion that routes data around the world. Additionally, much of the network is virtualized, meaning that it will rely on software running on dynamically configurable hardware. This design dramatically increases the points vulnerable to attack, as does the expected massive increase in both things connected to the network and the data flying about it{{% /quote %}}.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="5G Security" link="https://www.schneier.com/blog/archives/2020/01/china_isnt_the_.html" lang="en" >}}
{{% quote %}}Second, there's so much backward compatibility built into the 5G network that older vulnerabilities remain. 5G is an evolution of the decade-old 4G network, and most networks will mix generations. Without the ability to do a clean break from 4G to 5G, it will simply be impossible to improve security in some areas. Attackers may be able to force 5G systems to use more vulnerable 4G protocols, for example, and 5G networks will [inherit](https://gcn.com/articles/2019/10/21/5g-security.aspx) many existing problems{{% /quote %}}.
{{< /fig-quote >}}

いや，そんな継承いらんですよ `orz`

{{< fig-quote type="markdown" title="5G Security" link="https://www.schneier.com/blog/archives/2020/01/china_isnt_the_.html" lang="en" >}}
{{% quote %}}Third, the 5G standards committees missed many opportunities to improve security. Many of the new security features in 5G are optional, and network operators can choose not to implement them. The same [happened](https://www.wired.com/story/5g-more-secure-4g-except-when-not/) with 4G; operators even ignored security features defined as mandatory in the standard because implementing them was expensive. But even worse, for 5G, development, performance, cost, and time to market were all prioritized over security, which was treated as an afterthought{{% /quote %}}.
{{< /fig-quote >}}

これで思い出したのは，本家旧ブログで1010年以上前（！）に書いた「[GSM の暗号の解読](https://baldanders.info/blog/000472/)」である。
その中で紹介している『[情報セキュリティ技術大全](https://www.amazon.co.jp/dp/4822281426?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』の中身を少し紹介しよう。

{{< fig-quote type="markdown" title="『情報セキュリティ技術大全』（p.355-356）" link="https://www.amazon.co.jp/dp/4822281426?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
電話会社の観点からいえば、 GSM は成功だった。 Vodafone などの GSM 通信事業者の株主は莫大な利益を得た。その（ほんの）一部は、 GSM のチャレンジレスポンスメカニズムによってクローニングに歯止めがかかったおかげである。暗号の弱点など大した問題ではなかった。それらが悪用されたことはまったく（少なくとも、通信料収入を大きく損なうような形では）なかったからである。

（中略）

犯罪者の観点からも、 GSM は申し分なかった。 GSM は、電話サービスの盗聴に対する歯止めにはならなかった。単に犯罪の手口が変わっただけで、そのつけはクレジットカード会社や、アイデンティティ窃盗や路上での強盗に遭った被害者個人に回された。

（中略）

顧客の観点から見れば、 GSM は本来完全に安全なものとして販売されていた。これは正しかったのだろうか。確かに基地局までの通信の暗号化によって、アナログ電話でよく悩まされていた日常生活の盗聴はなくなった （有名人が被害を受け、 世間の注目を集めた事件がいくつかあった。たとえば，チャールズ英皇太子が離婚前に愛人との会話を盗聴された英国での事件や、 ニュート・ギングリッジ元米下院議長が絡んだ米国での事件などである）。しかし、世界中のほぼすべての電話盗聴は巨大な諜報機関によって行われており、彼等にとって暗号はほとんど問題にならない。
{{< /fig-quote >}}

“[5G Security](https://www.schneier.com/blog/archives/2020/01/china_isnt_the_.html "5G Security - Schneier on Security")” に於いても，以下のように指摘されている。

{{< fig-quote type="markdown" title="5G Security" link="https://www.schneier.com/blog/archives/2020/01/china_isnt_the_.html" lang="en" >}}
{{% quote %}}Chinese, Iranians, North Koreans, and Russians have been breaking into U.S. networks for years without having any control over the hardware, the software, or the companies that produce the devices. (And the U.S. National Security Agency, or NSA, has been breaking into foreign networks for years without having to coerce companies into deliberately adding backdoors.) Nothing in 5G prevents these activities from continuing, even increasing, in the future{{% /quote %}}.
{{< /fig-quote >}}

歴史は繰り返す。
ハムスターの回し車のように...

{{% review-paapi "4822281426" %}} <!-- 情報セキュリティ技術大全 -->
{{% review-paapi "B07925HGCX" %}} <!-- ハムスターの回し車 -->
