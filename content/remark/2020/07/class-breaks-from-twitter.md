+++
title = "Twitter から始まる Class Break"
date =  "2020-07-21T13:08:08+09:00"
description = "この件で4年前の米国大統領選挙を連想した人は多いだろうねぇ。"
image = "/images/attention/kitten.jpg"
tags  = [ "security", "vulnerability", "risk", "twitter", "class-break" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

さて，先週 Twitter 界隈で騒ぎになった事件について。
ぼちぼちと情報が出始めているようなので，覚え書きとして記しておこう。

今回の事件について Bruce Schneier 先生によるエッセイが公開されている。

- [On the Twitter Hack - Schneier on Security](https://www.schneier.com/blog/archives/2020/07/on_the_twitter_.html)
- [ブログ: Twitterのハッキングについて](https://okuranagaimo.blogspot.com/2020/07/twitter_21.html)

事件を一言で言うとこんな感じ。

{{< fig-quote type="markdown" title="On the Twitter Hack - Schneier on Security" link="https://www.schneier.com/blog/archives/2020/07/on_the_twitter_.html" lang="en" >}}
{{% quote %}}Someone compromised the entire Twitter network, probably by stealing the log-in credentials of one of Twitter's system administrators{{% /quote %}}.
{{< /fig-quote >}}

件の「誰か」が Twitter のシステム管理者からどうやって認証情報を引き出したのかはよく分かっていないようだ。
従業員への賄賂の可能性やら SIM スワッピング攻撃やら，様々な可能性を考慮して捜査中らしい。

{{< fig-quote type="markdown" title="On the Twitter Hack - Schneier on Security" link="https://www.schneier.com/blog/archives/2020/07/on_the_twitter_.html" lang="en" >}}
{{% quote %}}In the Twitter case this week, the hacker's tactics weren't [particularly sophisticated](https://twitter.com/TwitterSupport/status/1283591846464233474). We will almost certainly learn about security lapses at Twitter that enabled the hack, possibly including a [SIM-swapping attack](https://krebsonsecurity.com/2020/07/whos-behind-wednesdays-epic-twitter-hack/) that targeted an employee's cellular service provider, or maybe even a [bribed insider](https://www.vice.com/en_us/article/jgxd3d/twitter-insider-access-panel-account-hacks-biden-uber-bezos). The [FBI is investigating](https://www.washingtonpost.com/technology/2020/07/16/twitter-security-breach-response/){{% /quote %}}.
{{< /fig-quote >}}

Twitter 側は今回の件に関して，被害に遭っていると思われるアカウントを一時停止した。
その後はほぼ復旧しているらしい。

{{< fig-quote type="markdown" title="An update on our security incident" link="https://blog.twitter.com/en_us/topics/company/2020/an-update-on-our-security-incident.html" lang="en" >}}
{{% quote %}}In addition to our efforts behind the scenes, shortly after we became aware of the ongoing situation, we took preemptive measures to restrict functionality for many accounts on Twitter - this included things like preventing them from Tweeting or changing passwords. We did this to prevent the attackers from further spreading their scam as well as to prevent them from being able to take control of any additional accounts while we were investigating. We also locked accounts where a password had been recently changed out of an abundance of caution. Late on Wednesday, we were able to return Tweeting functionality to many accounts, and as of today, have restored most of the accounts that were locked pending password changes for their owners{{% /quote %}}.
{{< /fig-quote >}}

今回のポイントは通常の認証の問題（パスワードの複雑性や2要素認証など）をすっ飛ばして，直にシステムに入り込みやりたい放題にされてしまったことだろう。
これでは私達一般利用者は対処のしようがない。

{{< fig-quote type="markdown" title="On the Twitter Hack - Schneier on Security" link="https://www.schneier.com/blog/archives/2020/07/on_the_twitter_.html" lang="en" >}}
{{% quote %}}This kind of attack is known as a "[class break](https://www.schneier.com/blog/archives/2017/01/class_breaks.html)." Class breaks are endemic to computerized systems, and they're not something that we as users can defend against with better personal security. It didn't matter whether individual accounts had a complicated and hard-to-remember password, or two-factor authentication. It didn't matter whether the accounts were normally accessed via a Mac or a PC. There was literally nothing any user could do to protect against it{{% /quote %}}.
{{< /fig-quote >}}

{{% quote lang="en" %}}Class Break{{% /quote %}} とは，ひとつまたは関連する複数のシステム間で共通の脆弱性を持ち，その脆弱性を使って全てが一気に破壊される状況を指す。

{{< fig-quote type="markdown" title="On the Twitter Hack - Schneier on Security" link="https://www.schneier.com/blog/archives/2020/07/on_the_twitter_.html" lang="en" >}}
{{% quote %}}Class breaks are security vulnerabilities that break not just one system, but an entire class of systems. They might exploit a vulnerability in a particular operating system that allows an attacker to take remote control of every computer that runs on that system's software. Or a vulnerability in internet-enabled digital video recorders and webcams that allows an attacker to recruit those devices into a massive botnet. Or a single vulnerability in the Twitter network that allows an attacker to take over every account{{% /quote %}}.
{{< /fig-quote >}}

{{% quote lang="en" %}}Class Break{{% /quote %}} については Bruce Schneier 先生のあの本にも言及がある。

{{< fig-quote type="markdown" title="セキュリティはなぜやぶられたのか" link="https://www.amazon.co.jp/dp/4822283100?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
{{% quote %}}技術が進むと標準化が進み、脆弱性が増える。同じ機能を持つ部分すべてを破壊する「クラスブレーク」が可能になるのだ{{% /quote %}}（p.133）
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="セキュリティはなぜやぶられたのか" link="https://www.amazon.co.jp/dp/4822283100?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
{{% quote %}}システムがあまりに複雑になり技術的に高度になった結果、セキュリティが劣るかもしれない新しいものを使うより、クラスブレークのおそれがあっても実績のあるものを使い続けたほうがいいことが多くなったのだ{{% /quote %}}（p.158）
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="セキュリティはなぜやぶられたのか" link="https://www.amazon.co.jp/dp/4822283100?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
{{% quote %}}画一的なシステムはクラスブレークをうけやすく、剛性が高い。多様な場合にもクラスブレークは可能だが、難しくなる。可能なかぎり多様性を選んだほうがいいのだ。画一性自体がセキュリティリスクになるが、そのリスクは、確立されていない方法を選んだ場合よりも小さいことが多い{{% /quote %}}（p.176）
{{< /fig-quote >}}

不幸中の幸いというか，今回は[金銭絡みの比較的穏便（笑）な詐欺](https://piyolog.hatenadiary.jp/entry/2020/07/22/060139 "Twitter社内管理ツールの不正アクセスについてまとめてみた - piyolog")だったが，可能性としてはもっと大規模なこともできた筈である。

{{< fig-quote type="markdown" title="On the Twitter Hack - Schneier on Security" link="https://www.schneier.com/blog/archives/2020/07/on_the_twitter_.html" lang="en" >}}
{{% quote %}}Imagine a government using this sort of attack against another government, coordinating a series of fake tweets from hundreds of politicians and other public figures the day before a major election, to affect the outcome. Or to escalate an [international dispute](https://onezero.medium.com/twitter-security-flaws-pose-a-unique-threat-to-nuclear-diplomacy-experts-say-b509e0eb2aad). Done well, it would be devastating{{% /quote %}}.
{{< /fig-quote >}}

あるいは今回は何かの前準備なのか？ この件で4年前の米国大統領選挙を連想した人は多いだろうねぇ。

{{% quote lang="en" %}}Class Break{{% /quote %}} について面白い喩えがある。

{{< fig-quote type="markdown" title="Class Breaks - Schneier on Security" link="https://www.schneier.com/blog/archives/2017/01/class_breaks.html" lang="en" >}}
{{% quote %}}In a sense, class breaks are not a new concept in risk management. It's the difference between home burglaries and fires, which happen occasionally to different houses in a neighborhood over the course of the year, and floods and earthquakes, which either happen to everyone in the neighborhood or no one. Insurance companies can handle both types of risk, but they are inherently different. The increasing computerization of everything is moving us from a burglary/fire risk model to a flood/earthquake model, which a given threat either affects everyone in town or doesn't happen at all{{% /quote %}}.
{{< /fig-quote >}}

洪水や地震への対策と同じように {{% quote lang="en" %}}Class Break{{% /quote %}} への対策も中長期的なタイム・スケールで考えないといけないということだろう。
先ず，どこに問題があるか洗い出すところからだけど（笑） そういう意味で今回のケースは（ゼロトラスト・セキュリティの観点からも）教訓になる筈である。

## ブックマーク

- [Hackers Convinced Twitter Employee to Help Them Hijack Accounts](https://www.vice.com/en_us/article/jgxd3d/twitter-insider-access-panel-account-hacks-biden-uber-bezos)
    - [Twitter Hackers May Have Bribed an Insider - Schneier on Security](https://www.schneier.com/blog/archives/2020/07/twitter_hackers.html)
- [AppleやマスクCEOなど多数のセレブTwitterアカウントが乗っ取られ、暗号通貨詐欺に悪用される - ITmedia NEWS](https://www.itmedia.co.jp/news/articles/2007/16/news054.html)
- [Twitterの暗号通貨詐欺の元凶は内部ツールに不正アクセスした一人のハッカー  |  TechCrunch Japan](https://jp.techcrunch.com/2020/07/16/2020-07-15-twitter-hacker-admin-scam/)
- [Twitter社内管理ツールの不正アクセスについてまとめてみた - piyolog](https://piyolog.hatenadiary.jp/entry/2020/07/22/060139)

## 参考図書

{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
