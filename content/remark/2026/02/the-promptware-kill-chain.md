+++
title = "「プロンプトウェア・キルチェーン」"
date =  "2026-02-17T10:25:14+09:00"
description = "複数のレイヤできっちり仕事をされる Bruce Schneier 先生マジすげーなと思ったのであった。"
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "security", "risk", "management" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

“Promptware Kill Chain” をどう訳していいか分からなかったので，とりあえずそのままカタカナで「プロンプトウェア・キルチェーン」としてしまった[^kc1]。
きっとそのうち偉い人が適切な訳語を考えてくれるだろう。
英単語をカタカナに伸ばしただけなんてダサいし，それによって言葉の意味が消失するのは拙いと思う[^es1]。
けど，私は英語不得手だからなぁ...

[^kc1]: “kill chain” は元々軍事用語で，敵軍の構造を破壊（切断する）することで自軍を守る先制処置の考え方らしい。これをサイバーセキュリティの分野に持ち込んだのが “Cyber Kill Chain” なんだそうで，攻撃の構造を理解し，各フェーズの攻撃に対して有効な対策をとることで自システムを防御する考え方を指すとのこと。故に “Promptware Kill Chain” はプロンプトウェアの7段階を理解し，各段階それぞれに対して有効な対策をとるべき，という意図なのだろう。
[^es1]: “enshittification” を「[メタクソ化](https://p2ptk.org/monopoly/4366 "メタクソ化するTiktok：プラットフォームが生まれ、成長し、支配し、滅びるまで » p2ptk[.]org")」と訳した [p2ptk.org](https://p2ptk.org/ "P2Pとかその辺のお話R | Sharing is Caring") の中の人は天才だと思う。

- [The Promptware Kill Chain | Lawfare](https://www.lawfaremedia.org/article/the-promptware-kill-chain)
- [The Promptware Kill Chain - Schneier on Security](https://www.schneier.com/blog/archives/2026/02/the-promptware-kill-chain.html)

（内容は同じ）

この記事は Oleg Brodt, Elad Feldman, Bruce Schneier, Ben Nassi による共著で，斜め読みするに以下の（同じ著者による）論文の解説になっているようだ。

- [[2601.09625] The Promptware Kill Chain: How Prompt Injections Gradually Evolved Into a Multistep Malware Delivery Mechanism](https://arxiv.org/abs/2601.09625)

論文の要旨は以下の通り：

{{< fig-quote type="markdown" title="The Promptware Kill Chain: How Prompt Injections Gradually Evolved Into a Multistep Malware Delivery Mechanism" link="https://arxiv.org/abs/2601.09625" lang="en" >}}
Prompt injection was initially framed as the large language model (LLM) analogue of SQL injection. However, over the past three years, attacks labeled as prompt injection have evolved from isolated input-manipulation exploits into multistep attack mechanisms that resemble malware. In this paper, we argue that prompt injections evolved into promptware, a new class of malware execution mechanism triggered through prompts engineered to exploit an application's LLM.
[...]
By moving the conversation from prompt injection to a promptware kill chain, our work provides analytical clarity, enables structured risk assessment, and lays a foundation for systematic security engineering of LLM-based systems.
{{< /fig-quote >}}

つまり，プロンプト・インジェクションやジャイルブレイクといった個々の事象にフォーカスするのではなくマルウェアに類似した多段階の攻撃メカニズムを「プロンプトウェア」と定義し，理解に必要な語彙やフレームワークを政策立案者やセキュリティ専門家に提供することが目的のようだ。
なので，既に AI システムを導入している（あるいはこれから導入する）企業・組織のセキュリティ担当者は，論文のほうを読んでおくといいかもしれない。

プロンプトウェアには以下の7つの段階がある。

1. **Initial Access (初期アクセス)** : プロンプト・インジェクション等（悪意のあるペイロードが AI システムに侵入）
2. **Privilege Escalation (権限昇格)** : ジャイルブレイク等（安全トレーニングやポリシーのガードレールの回避）
3. **Reconnaissance (偵察)** : 資産，接続されているサービス，および能力に関する情報を明らかにさせる
4. **Persistence (永続化)** : 長期記憶に自身を埋め込む，またはエージェントが依存するデータベースを汚染
5. **Command and Control (C2; 指揮統制)** : 攻撃者がその挙動を変更できる制御可能なトロイの木馬へと進化
6. **Lateral Movement (水平展開)** : 他のユーザ，デバイス，またはシステムへと拡散。マルウェア拡散の高速道路を作り出す
7. **Actions on Objective (目的の実行)**

{{< fig-img-quote src="./promptware-kill-chain-660w.jpg" title="The Promptware Kill Chain - Schneier on Security" link="https://www.schneier.com/blog/archives/2026/02/the-promptware-kill-chain.html" lang="en" width="660" >}}

こうやって見るとプロンプト・インジェクションなどは一連の攻撃の取っ掛かりに過ぎないことが分かる。
最終段階では仕込まれたマルウェアが発火して攻撃者の目的が実行される。

{{< fig-quote type="markdown" title="The Promptware Kill Chain - Schneier on Security" link="https://www.schneier.com/blog/archives/2026/02/the-promptware-kill-chain.html" lang="en" >}}
The goal of promptware is not just to make a chatbot say something offensive; it is often to achieve tangible malicious outcomes through data exfiltration, financial fraud, or even physical world impact.
{{< /fig-quote >}}

実際に AI エージェントによって攻撃者のウォレットに[暗号資産が送金](https://crypto.news/aixbt-agent-hacked-losing-55eth-aixbt-token-drops-2025/ "AiXBT agent hacked, loses 55.50 ETH as token tumbles 16%")された事例とかあるらしい。
他にも

- [[2508.12175] Invitation Is All You Need! Promptware Attacks Against LLM-Powered Assistants in Production Are Practical and Dangerous](https://arxiv.org/abs/2508.12175)
- [Here Comes the AI Worm: Preventing the Propagation of Adversarial Self-Replicating Prompts Within GenAI Ecosystems | Proceedings of the 2025 ACM SIGSAC Conference on Computer and Communications Security](https://dl.acm.org/doi/10.1145/3719027.3765196)

といった論文がある。

根本的な問題は LLM 自体のアーキテクチャにあると言う。

{{< fig-quote type="markdown" title="The Promptware Kill Chain - Schneier on Security" link="https://www.schneier.com/blog/archives/2026/02/the-promptware-kill-chain.html" lang="en" >}}
Unlike traditional computing systems that strictly separate executable code from user data, LLMs process all input—whether it is a system command, a user’s email, or a retrieved document—as a single, undifferentiated sequence of tokens. There is no architectural boundary to enforce a distinction between trusted instructions and untrusted data. Consequently, a malicious instruction embedded in a seemingly harmless document is processed with the same authority as a system command.
{{< /fig-quote >}}

つまり LLM はコマンドとデータを区別できず全てを一連のトークン列として処理するため，容易に悪意を忍ばせることができる。

初期アクセス（プロンプト・インジェクション等）を防ぐことが難しいなら，後続のステップで防ぐしかない。
権限昇格の制限，偵察の制約，永続化の防止，C2の妨害，エージェントが実行を許可されるアクションの制限といった対策で多層防衛を構築する必要がある。

某マスク氏率いる xAI について

{{< fig-quote type="markdown" title="Is safety ‘dead’ at xAI? | TechCrunch" link="https://techcrunch.com/2026/02/14/is-safety-is-dead-at-xai/" lang="en" >}}
One source said, “Safety is a dead org at xAI,” while the other said that Musk is “actively is trying to make the model more unhinged because safety means censorship, in a sense, to him.”
{{< /fig-quote >}}

みたいな話もあるようだが，そんな牧歌的な段階はとうに過ぎてるということだろう。
他の大手 AI サービス・プロバイダはどうなんだろうねぇ。

そういえば先日公開された「[AIによる民主主義の巻き返しは可能か](https://wirelesswire.jp/2026/02/92715/ "AIによる民主主義の巻き返しは可能か – WirelessWire & Schrödinger's")」で “[Rewiring Democracy](https://www.schneier.com/books/rewiring-democracy/ "Rewiring Democracy - Schneier on Security")” に関して

{{< fig-quote type="markdown" title="AIによる民主主義の巻き返しは可能か – WirelessWire & Schrödinger's" link="https://wirelesswire.jp/2026/02/92715/" >}}
実際、やはり情報セキュリティ分野の専門家であるベン・ロスキーが本書について、「AIを民主主義のために利用する上で最大の障壁はセキュリティだが、その話が第27章になるまで出てこない」「セキュリティなくして信頼は成立しないのに、著者らは今日のAI開発において、セキュリティは後付けの考慮事項だと言わんばかりだ」と[不満げに書いている](https://www.rsaconference.com/library/blog/bens-book-of-the-month-rewiring-democracy)のを読んで思い至ったのですが...
{{< /fig-quote >}}

という指摘があったが，複数のレイヤできっちり仕事をされる Bruce Schneier 先生マジすげーなと思ったのであった。

## 参考

{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
{{% review-paapi "B0CK19L1HC" %}} <!-- ハッキング思考 Kindle 版 -->
