+++
title = "「ChatGPT は安全か？」"
date =  "2026-03-07T20:47:23+09:00"
description = "AI と内緒話をしてはいけない"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "security", "risk", "management" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

[前回]の記事のあとにこれだと「またぞろ政治の話か」と思われそうだが，今回は純粋にリスク・マネージメントの話。
[Proton のブログ](https://proton.me/blog "The Proton Blog - News from the front lines of privacy and security | Proton")より。

- [Is ChatGPT safe to use? Here’s what you should know. | Proton](https://proton.me/blog/is-chatgpt-safe)

ChatGPT に限らないが生成 AI サービスの大半はユーザが入力したデータやプロンプトをサービスプロバイダ側が入手できる。
実際にサービスプロバイダの多くはユーザの入力（データもプロンプトも）を保持して学習に使ってるし，インシデントの際は（政府を含む）第三者にそれを開示している。
また広告・分析会社に情報を売っていた例もあるし ChatGPT との[会話内容が Google 検索に拾われてしまった](https://rocket-boys.co.jp/security-measures-lab/chatgpt-conversations-exposed-on-search-engines-potential-personal-confidential-info-leak/ "ChatGPTの一部会話が検索エンジンで公開状態に-重要な個人情報や機密情報が漏洩する可能性|セキュリティとAIのニュース-セキュリティ対策 Lab")例もあったそうな。

そうしたリスクを改めて思い起こさせたのが[前回]紹介した事例で，これによって[400万人以上のユーザ][ChatGPT]が ChatGPT を止めると宣言しているらしい[^qg1]。

[^qg1]: 2026-03-07 時点。

件の記事では ChatGPT に対して「個人のプライバシー」「ビジネスリスク」「透明性の欠如」の3つのリスクを挙げている。
詳しい内容は（結構ボリュームがあるので）ここでは割愛する。

{{< fig-quote type="markdown" title="Is ChatGPT safe to use? Here’s what you should know." link="https://proton.me/blog/is-chatgpt-safe" lang="en" >}}
Even though ChatGPT can be helpful, you should never treat it like a secure vault for sensitive information. Avoid entering anything that could harm you, your company, or others if it were stored, reviewed, or accidentally exposed:
{{< /fig-quote >}}

それでも ChatGPT を使いたいならどうすりゃいいの？ ということで，入力してはいけないものとして以下を挙げている。

{{< fig-img-quote src="./is-chatgpt-safe-diagram.webp" title="Is ChatGPT safe to use? Here’s what you should know." link="https://proton.me/blog/is-chatgpt-safe" width="725" lang="en" >}}

具体的にはこんな感じ：

- **パスワードと認証データ：** アカウントのパスワード，二要素認証（2FA）コード，バックアップ認証コード，プライベートAPIキーなど
- **政府発行の識別番号：** 社会保障番号，国民ID，パスポート，運転免許証，納税者番号など
- **金融および銀行情報：** クレジットカードまたはデビットカードの番号，IBAN，オンラインバンキングの認証情報，投資口座のログイン情報，ビットコインウォレットの秘密鍵など
- **機密性の高い個人データ：** あなたやあなたの家族を特定または追跡するために使用される可能性のある情報。自宅の住所や電話番号，生年月日，プライベートな写真や文書など
- **健康情報：** 医療報告書，診断記録，保険番号，患者ID，またはあなたの身元に結びついた詳細な健康履歴など
- **機密性の高い業務データや企業データ：** 独自のソースコード，社内戦略文書，機密契約書，顧客データベース，クライアント詳細情報，財務予測，未公開レポート，NDAなど
- **法的情報および秘匿特権のある情報：** 弁護士と依頼人間の通信，訴訟戦略，証拠書類，秘密の和解交渉など

いや，まぁ， AI サービスでなくても当たり前なんだけどね，これ。
こういう情報をうっかり漏らしてしまうというのは，やっぱ生成 AI サービスはデータ構造が分かりにくい（というか構造がない）というのが根底にあるんじゃないだろうか，と思ったり。

さらに ChatGPT を安全に利用するために以下の習慣を身につけると良いとある。

- 保存，レビュー，または公に公開されたくない機密情報の共有は避ける
- 個人を特定できる情報を削除するか，プレースホルダーや架空の例に置き換える
- 機密情報や個人情報を含まないファイルのみをアップロードする
- AIチャットは，他の人に見られる可能性のあるメールやサポートチケットのように扱う
- プライバシー設定を確認し，チャット履歴，メモリ，AIトレーニングなどの設定は無効にする
- 不要になった会話を削除し，アカウントに関連付けられた個人情報の量を減らす

もちろん生成 AI サービスを乗換える手もある。

(明らかに宣伝だがw) Proton の [Lumo](https://lumo.proton.me/ "Lumo: Privacy-first AI assistant where chats stay confidential") はユーザデータを「[ゼロアクセス暗号化](https://proton.me/ja/learn/encryption/types-of-encryption/zero-access "ゼロアクセス暗号化とは？ | Proton")」することを売りにしていて，履歴を一切保存しないゴーストモードもあるらしい[^ps1]。

[^ps1]: [Lumo] に対する PrivacySpy のスコアはないが Proton VPS については [8.1](https://privacyspy.org/product/protonvpn/ "ProtonVPN | PrivacySpy") (10点満点) と評価されている。データ侵害が発生した場合のユーザへの対応がイマイチらしい。まぁ，高スコアではあるので [Lumo] もこのくらいならいいんだけど。

私が常用している [Kagi Assistant] は匿名性を売りにしている。
また「ユーザデータは負債である」というポリシーを持っていて，できるだけユーザデータをサービス側に保存しないようにしているそうな[^ps2]。

[^ps2]: [Kagi Assistant] に対する PrivacySpy のスコアはないが Kagi 検索サービスについては [7.5](https://privacyspy.org/product/kagi/ "Kagi | PrivacySpy") と評価されている。一般的なセキュリティ・ポリシーが曖昧なのと，法執行機関の要請への対応が明記されてない点がマイナス要素らしい。

いっそのこと SaaS など使わずオンプレミスで AI 環境を構築するという手もある。
企業・組織はこのほうがいいかも知れない。
軍事関係もね（笑）

（Grok や Google は元から悪いが）[今回の騒動][前回]で Anthropic や OpenAI といった大手のサービスはセキュリティやプライバシーのリスク観点からは旗色が悪くなったように思う。
各サービスにおける運用のポリシーや透明性を見極めて，賢く使うのがいいと思う。
所詮は道具なんだから。

## ブックマーク

- [AI は正解を答えない]({{< ref "/remark/2026/02/ai-doesnt-provide-definitive-answers.md" >}})
- [「プロンプトウェア・キルチェーン」]({{< ref "/remark/2026/02/the-promptware-kill-chain.md" >}})

[前回]: {{< relref "./military-use-of-ai.md" >}} "AI の軍事利用"
[ChatGPT]: https://quitgpt.org/ "QuitGPT — ChatGPT takes Trump's killer robot deal"
[Lumo]: https://lumo.proton.me/ "Lumo: Privacy-first AI assistant where chats stay confidential"
[Kagi Assistant]: https://kagi.com/assistant "Kagi Assistant"

## 参考

{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
{{% review-paapi "B0CK19L1HC" %}} <!-- ハッキング思考 Kindle 版 -->
