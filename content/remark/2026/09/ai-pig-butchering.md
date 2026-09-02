+++
title = "AI による豚の屠殺詐欺か？"
date =  "2026-09-02T11:02:05+09:00"
description = "今回は我らが Bruce Schneier 先生のブログ記事から。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "security", "risk", "phishing", "spam", "messaging", "communication" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

## 豚の屠殺詐欺？

今回は我らが Bruce Schneier 先生のブログ記事から。

- [What's the Scam? - Schneier on Security](https://www.schneier.com/blog/archives/2026/09/whats-the-scam.html)

彼が提供するニュースレターの購読手順は (1) ウェブページで情報を入力 (2) 自動生成されたメールに返信 というもの。
まぁ一般的な手順。
で，最近そのメールに対して1行だけの個別返信が届くようになったそうな。
曰く

{{< fig-quote type="markdown" title="What's the Scam? - Schneier on Security" link="https://www.schneier.com/blog/archives/2026/09/whats-the-scam.html" lang="en" >}}
Thank you for the positive impact your emails have had on my life.<br>
Your emails are a game-changer.<br>
Your emails are a constant reminder of why I subscribed.<br>
Your emails rock.<br>
Thank you for the time and effort you put into creating these informative emails.<br>
Thank you for the passion and enthusiasm you infuse into your email content.<br>
Your emails consistently exceed my expectations. Thank you for the exceptional value!
{{< /fig-quote >}}

Bruce Schneier 氏は以前から確認メールの返信で個別のメッセージを受け取ることがあったそうで，最初はこれらが偽物であることに気付かなかったらしい。
まぁ，でも，こういうのが一遍にくれば怪しいと思うよね。
しかもよく見れば AI 臭い文章だし（笑），メールアドレスも Gmail でアカウント名がランダムな文字列になってるし。
件のメールアドレス群は最終的にはニュースレターの購読を行っていないそうな。

氏は最初 "pig butchering scam" を疑ったそうだが，それにしては氏の返信以降音沙汰がないので ？？？ な感じらしい。

"Pig butchering scam" は日本語で「[豚の屠殺詐欺](https://blog.kaspersky.co.jp/pig-butchering-crypto-investment-scam/36039/ "豚の屠殺詐欺：大規模な仮想通貨詐欺 | カスペルスキー公式ブログ")」などと訳されているようだ。
SNS やメッセージングアプリなどを使って最初は友好的な態度で接近し，仲良くなってから「儲かる投資話」などを持ちかける手口なんだとか。
これが「豚を肥え太らせて屠殺する」行為に類似してるとして名付けられたらしい。
2022年ごろからこの手の詐欺報告が増えてるそうで [Bruce Schneier 氏も言及](https://www.schneier.com/blog/archives/2023/02/on-pig-butchering-scams.html "On Pig Butchering Scams - Schneier on Security")されていた。
いわゆる「ロマンス詐欺」もこれの一種かな？

## 記事の反響

記事は最後に "Anyone have any idea?" と締めくくっていて，これに対するコメントがなかなか面白かった。

多かったのはメールアカウントのウォームアップまたはエージング仮説。
将来 spam 送信や詐欺に使うため，メールでの会話履歴を作り Google による不正検知を回避しようとしてるのではないか，というもの。
その会話履歴を AI で自動生成し省力化しているわけだ。
これなら会話を継続させる必要もないだろう。

収集したメールアドレスが「生きてる」かどうか確認し，さらに返信内容からユーザのアクティビティを推測しているのではないか，という意見もあった。

面白かったのは AI エージェントによる自律的[^a1] な実行または誤動作の可能性。
特定のコミュニティに属する人々の反応や，購読者限定の情報を集める目的があったのではないか，というもの。
あるいは，単に「メールに自動返信する」というタスクを実行しただけというもの。
さらには，エージェント特有の多段階作業の副作用で与えられたタスクから外れた行動を取ったのではないか，という意見もあった。

[^a1]: 私としては何度でも「自律機械は存在しない」と主張しておこう。詳しくは「[自律的な AI エージェントとは]({{< ref "/remark/2026/07/what-is-an-autonomous-ai-agent.md" >}})」参照。

ちなみに [Kagi Assistant] に AI エージェントが独力で Gmail アカウントを取得することは可能か聞いてみたが，少なくともアカウントの大量取得についてはエージェント単独では無理じゃね？ と言われた。

## AI はチョロいか？

記事やそれに対するコメントをつらつらと読んで妄想したのは，メールであれメッセージングアプリであれ，メッセージの送り手と受け手の双方が AI エージェントを代理に使っていた場合，どのようなやり取りが行われるのだろう，ということだ。
少なくとも送られてきたメールをエージェントが読んで，それらしい返信を行うというのはできるっぽい。
人間にはやりとりの結果だけ知らされるが，実はその何倍もの情報をやり取りしてるってのはありそうな話である。
あるいは AI 間のやり取りでなら「豚の屠殺詐欺」など回りくどいことをせずとも案外チョロいんではないか，などと思ったりする。

というわけで，色々と妄想がはかどる記事であった。

## ブックマーク

- [Internet Crime Complaint Center (IC3) | Cryptocurrency Investment Schemes](https://www.ic3.gov/PSA/2022/PSA221003)
- [Pig butchering - how to spot and report the scam - DFPI](https://dfpi.ca.gov/news/insights/pig-butchering-how-to-spot-and-report-the-scam/)
- [豚の食肉解体詐欺とは【用語集詳細】](https://www.sompocybersecurity.com/column/glossary/pig-butchering-scam)
- [プロンプターとしての人間，演劇者としての AI （「『AIと生きる』を読む」より）]({{< ref "/remark/2026/05/reading-living-with-ai.md#prompter" >}})

[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"
