+++
title = "Signal グループチャット内容の漏洩と NSA"
date =  "2025-04-02T01:36:55+09:00"
description = "というわけでメッセージング・アプリは米国政府関係者からテロリストまで多くの人に愛される Signal を使いましょう。"
image = "/images/attention/kitten.jpg"
tags = [ "security", "management", "risk", "cryptography", "communication", "signal" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

先週の話であるが [Signal] のグループチャットから国家機密情報が漏洩していたというニュースがあった。

- [The Trump Administration Accidentally Texted Me Its War Plans - The Atlantic](https://www.theatlantic.com/politics/archive/2025/03/trump-administration-accidentally-texted-me-its-war-plans/682151/)

この記事は購読しないと見れないが，私は Bluesky の TL で[見かけて](https://bsky.app/profile/taineko399.bsky.social/post/3ll5ru4pebc2g "朝起きてトランプ政権に驚かされる事はよくあるけれど、今朝のこれは凄いな。...")知った。

{{< fig-gen >}}
<blockquote class="bluesky-embed" data-bluesky-uri="at://did:plc:gknmy4fk7o2yocdywvtue5il/app.bsky.feed.post/3ll5ru4pebc2g" data-bluesky-cid="bafyreiemvh5lrqrbj4t45aasrmz2hbyofpi22o3po7vtddcqy7ywolpp44" data-bluesky-embed-color-mode="system"><p lang="ja">朝起きてトランプ政権に驚かされる事はよくあるけれど、今朝のこれは凄いな。

トランプ政権の重要閣僚がイエメンのフーシ派攻撃計画をシグナルで話し合っており、グループチャットに誤ってThe Atlanticの編集長を招待していたという話。
ゴールドバーグ編集長は最初何かの罠なのか、それともいたずらかと無言で成り行きを見守っていたが、参加者の言動があまりにも「らしい」上に、ユーザーネームが「ピート・ヘグセス」の発言通りにイエメンが爆撃され始めて、重大な国家機密の漏洩に意図せず巻き込まれた事を悟ったと。　1/3

www.theatlantic.com/politics/arc...<br><br><a href="https://bsky.app/profile/did:plc:gknmy4fk7o2yocdywvtue5il/post/3ll5ru4pebc2g?ref_src=embed">[image or embed]</a></p>&mdash; 鯛猫 (<a href="https://bsky.app/profile/did:plc:gknmy4fk7o2yocdywvtue5il?ref_src=embed">@taineko399.bsky.social</a>) <a href="https://bsky.app/profile/did:plc:gknmy4fk7o2yocdywvtue5il/post/3ll5ru4pebc2g?ref_src=embed">March 25, 2025 at 7:16 AM</a></blockquote><script async src="https://embed.bsky.app/static/embed.js" charset="utf-8"></script>
{{< /fig-gen >}}

笑うのが「誤ってThe Atlanticの編集長を招待」してグループチャットを開始したのが国家安全保障担当補佐官 (US National Security Advisor) の Mike Waltz 氏というところだろうか。
彼は “The Atlantic” の編集長をグループチャットに加えたことを認めておらず

{{< fig-quote type="markdown" title="Trump admin's shifting explanations for how a journalist was added to Signal chat - ABC News" link="https://abcnews.go.com/Politics/trump-admins-shifting-explanations-journalist-added-signal-chat/story?id=120179649" lang="en" >}}
"Of course, I didn't see this loser in the group. It looked like someone else. Now, whether he did it deliberately or it happened in some other technical mean, is something we're trying to figure out," Waltz said, without evidence, on Fox News.
{{< /fig-quote >}}

などと語ってるらしい（言いがかり？）。
もちろん “The Atlantic” の編集長は

{{< fig-quote type="markdown" title="Trump admin's shifting explanations for how a journalist was added to Signal chat - ABC News" link="https://abcnews.go.com/Politics/trump-admins-shifting-explanations-journalist-added-signal-chat/story?id=120179649" lang="en" >}}
"I didn't hack into anyone's phone," Goldberg told ABC News Live anchor Kyra Phillips on Wednesday. "Mike Waltz invited me to Signal and then he invited me to a group. I don't know how to say it more simply than that."
{{< /fig-quote >}}

と反論している。

これに関連して面白かったのが，遡ること2025年2月に NSA (National Security Agency; 国家安全保障局) が自身の職員に対し [Signal] に脆弱性があると警告していたというニュースだ（これ自体は最初に挙げたニュースの翌日に流された）。

{{< fig-quote type="markdown" title="NSA warned of vulnerabilities in Signal app a month before Houthi strike chat - CBS News" link="https://www.cbsnews.com/news/nsa-signal-app-vulnerabilities-before-houthi-strike-chat/" lang="en" >}}
The National Security Agency sent out an operational security special bulletin to its employees in February 2025 warning them of vulnerabilities in using the encrypted messaging application Signal, according to internal NSA documents obtained by CBS News.

[...]

The bulletin warned of [Russian](https://www.cbsnews.com/news/trump-envoy-steve-witkoff-signal-text-group-chat-russia-putin/) professional hacking groups employing phishing scams to gain access to encrypted conversations, bypassing the end-to-end encryption the application uses. 

The bulletin also underscored to NSA employees that third-party messaging applications such as [Signal](https://www.cbsnews.com/news/what-is-signal-app-messaging/) and Whatsapp are permitted for certain "unclassified accountability/recall exercises" but not for communicating more sensitive information.
{{< /fig-quote >}}

ちなみに「[Signal] に脆弱性がある」という点について [Signal] は「[それは脆弱性じゃなくて Phishing だろ](https://x.com/signalapp/status/1904666111989166408 "The memo used the term ‘vulnerability’ in relation to Signal—but it had nothing to do with Signal’s core tech. It was warning against phishing scams targeting Signal users. ...")」と反論している[^s1]。
それはともかく，今回は国家安全保障担当補佐官氏の PON だと思うけどな（いや，そんな可愛らしいもんじゃないが`w`）。

[^s1]: [Signal] 側はグループチャットへのソーシャル・エンジニアリング（Phishing）に対抗する手段を2025年2月の時点で既に[実装](https://www.wired.com/story/russia-signal-qr-code-phishing-attack/)している。 NSA 内部に向けた警告はこれに絡む話だったのかもしれない。いずれにしろ NSA 内部で自前でないサービスは使わないだろうけど。ちなみに [Signal] のコードは AGPLv3 ライセンス下で[公開](https://github.com/signalapp "Signal")されている。また [EFF (Electronic Frontier Foundation)](https://www.eff.org/) での評価も高く，かつてあった “[Secure Messaging Scorecard](https://www.eff.org/pages/secure-messaging-scorecard "Secure Messaging Scorecard | Electronic Frontier Foundation")” でも満点のスコアを叩き出したほか，今でも “[Surveillance Self-Defense](https://ssd.eff.org/ "Surveillance Self-Defense")” に “[How to: Use Signal](https://ssd.eff.org/module/how-to-use-signal "How to: Use Signal | Surveillance Self-Defense")” のページがある。

とまぁ，ここまでが前説（なげーよ！）。
ここから Bruce Schneier 先生のエッセイが登場する。

- [The Signal Chat Leak and the NSA - Schneier on Security](https://www.schneier.com/blog/archives/2025/03/the-signal-chat-leak-and-the-nsa.html)

米国に於いて NSA には2つの役割がある。

{{< fig-quote type="markdown" title="The Signal Chat Leak and the NSA - Schneier on Security" link="https://www.schneier.com/blog/archives/2025/03/the-signal-chat-leak-and-the-nsa.html" lang="en" >}}
It’s common knowledge that the NSA’s mission is breaking into and eavesdropping on other countries’ networks. (During President George W. Bush’s administration, the NSA conducted warrantless taps into domestic communications as well—surveillance that several district courts ruled to be illegal before those decisions were later overturned by appeals courts. To this day, many legal experts maintain that the program violated federal privacy protections.) But the organization has a secondary, complementary responsibility: to protect US communications from others who want to spy on them. That is to say: While one part of the NSA is listening into foreign communications, another part is stopping foreigners from doing the same to Americans.
{{< /fig-quote >}}

つまり NSA は諜報活動として他国や敵対組織の通信を傍受する必要があり，同時に自国の通信の秘密は（他国や敵対組織から）守る必要がある。
ここで「セキュリティのトレードオフ」が発生するわけだ。

{{< fig-quote type="markdown" title="The Signal Chat Leak and the NSA - Schneier on Security" link="https://www.schneier.com/blog/archives/2025/03/the-signal-chat-leak-and-the-nsa.html" lang="en" >}}
Previously, it might have preferred to keep the flaw quiet and use it to listen to adversaries. Now, if the agency does that, it risks someone else finding the same vulnerability and using it against the US government. And if it was later disclosed that the NSA could have fixed the problem and didn’t, then the results might be catastrophic for the agency.
{{< /fig-quote >}}

特に興味深かったのがこの話（ちなみに2024年11月のロイター記事）：

{{< fig-quote type="markdown" title="China-linked hackers stole surveillance data from telecom companies, US says | Reuters" link="https://www.reuters.com/technology/cybersecurity/china-affiliated-actors-compromised-networks-multiple-telecom-companies-us-says-2024-11-13/" lang="en" >}}
WASHINGTON, Nov 13 (Reuters) - China-linked hackers have intercepted surveillance data intended for American law enforcement agencies after breaking in to an unspecified number of telecom companies, U.S. authorities said on Wednesday.

[...]

The announcement confirms the broad outlines of previous media reports, [especially those in the Wall Street Journal](https://www.reuters.com/technology/cybersecurity/chinese-hackers-breached-us-court-wiretap-systems-wsj-reports-2024-10-06/), that Chinese hackers were feared to have opened a back door into the interception systems used by law enforcement to surveil Americans' telecommunications.
{{< /fig-quote >}}

つまり米国が通信傍受のために仕掛けたバックドアに中国のハッカーがバックドアを開けたんじゃないか？ ということらしい。
裏口の裏口とか（笑）

スマートフォンおよびスマートフォン上のアプリについても同じことが言える。
というか，スマートフォンは重要な機密を守るという点において適切なデバイスとは言えない。

{{< fig-quote type="markdown" title="The Signal Chat Leak and the NSA - Schneier on Security" link="https://www.schneier.com/blog/archives/2025/03/the-signal-chat-leak-and-the-nsa.html" lang="en" >}}
The biggest risk of eavesdropping on a Signal conversation comes from the individual phones that the app is running on. While it’s largely unclear whether the US officials involved had downloaded the app onto personal or government-issued phones—although Witkoff suggested on X that the program was on his "[personal devices](https://x.com/SteveWitkoff/status/1904886084879720683)"—smartphones are consumer devices, not at all suitable for classified US government conversations.
{{< /fig-quote >}}

米国現政権を見てつくづく思うのは「真正の馬鹿は害悪にしかならない」ということだったが，今回のグループチャット内容の漏洩の件に関しては，ある意味で朗報だったかもしれない。
何故なら，セキュリティというのはもっとも脆弱なところから崩れていくものであり，セキュリティ管理はそこに基準を合わせる必要があることを再確認できたから。

{{< fig-quote type="markdown" title="The Signal Chat Leak and the NSA - Schneier on Security" link="https://www.schneier.com/blog/archives/2025/03/the-signal-chat-leak-and-the-nsa.html" lang="en" >}}
As long as smartphones are in the pocket of every government official, police officer, judge, CEO, and nuclear power plant operator—and now that they are being used for what the White House now calls calls  "sensitive," if not outright classified conversations among cabinet members—we need them to be as secure as possible. And that means no government-mandated backdoors.
{{< /fig-quote >}}

そして NSA は [Signal] を含むスマホアプリの脆弱性について「黙して語らず」というわけにはいかなくなった。

{{< fig-quote type="markdown" title="The Signal Chat Leak and the NSA - Schneier on Security" link="https://www.schneier.com/blog/archives/2025/03/the-signal-chat-leak-and-the-nsa.html" lang="en" >}}
Other governments, possibly including US allies, will now have much more incentive to break Signal’s security than they did in the past, and more incentive to hack US government smartphones than they did before March 24.

For just the same reason, the US government has urgent incentives to protect them.
{{< /fig-quote >}}

というわけでメッセージング・アプリは米国政府関係者からテロリストまで多くの人に愛される [Signal] を使いましょう。

## ブックマーク

- [The war on encryption is dangerous](https://www.ft.com/content/a934150f-e0f5-4e75-a2d1-a3671ea52ca0)
- [政府高官が軍事計画の情報共有に使っていたメッセージアプリ「Signal」の脆弱性についてNSAが職員に警告していたことが判明 - GIGAZINE](https://gigazine.net/news/20250326-nsa-warned-signal-app-vulnerabilities/)
- [暗号化の勝利：フランス議会がバックドア義務化を否決 » p2ptk[.]org](https://p2ptk.org/privacy/5397)

[Signal]: https://signal.org/ "Signal"

## 参考文献

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
