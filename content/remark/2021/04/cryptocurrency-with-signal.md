+++
title = "FinTech の曲がり角"
date =  "2021-04-24T15:12:30+09:00"
description = "藪をつついて蛇を出す"
image = "/images/attention/kitten.jpg"
tags = [ "fintech", "security", "privacy", "risk", "messaging" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

2週間以上前の記事で恐縮だが [Signal] がいわゆる暗号通貨（cryptocurrency）機能を組み込むらしい。

- [Signal Adds a Payments Feature—With a Privacy-Focused Cryptocurrency | WIRED](https://www.wired.com/story/signal-mobilecoin-payments-messaging-cryptocurrency/)
- [Signal Adds Cryptocurrency Support - Schneier on Security](https://www.schneier.com/blog/archives/2021/04/wtf-signal-adds-cryptocurrency-support.html)
- [Et tu, Signal?](https://www.stephendiehl.com/blog/signal.html)
- [暗号化チャットのSignalが英国でMobileCoinによる決済機能をテスト中  |  TechCrunch Japan](https://jp.techcrunch.com/2021/04/07/2021-04-06-signal-tests-payments-in-the-uk-using-mobilecoin/)
- [メッセージングアプリSignalの暗号通貨による送金機能の追加にブルース・シュナイアーが苦言 - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20210412/signal-cryptcurrency)

実はこの話はほぼスルーしていたのだが [yomoyomo さんの記事](https://yamdas.hatenablog.com/entry/20210412/signal-cryptcurrency "メッセージングアプリSignalの暗号通貨による送金機能の追加にブルース・シュナイアーが苦言 - YAMDAS現更新履歴")を見て，やはり「ログ」として残しておくのがいいだろうという気分になった。

この問題は Bruce Schneier 先生の以下の一文に集約できるだろう。

{{< fig-quote type="markdown" title="Signal Adds Cryptocurrency Support" link="https://www.schneier.com/blog/archives/2021/04/wtf-signal-adds-cryptocurrency-support.html" lang="en" >}}
{{% quote %}}It’s that adding a cryptocurrency to an end-to-end encrypted app muddies the morality of the product, and invites all sorts of government investigative and regulatory meddling: by the IRS, the SEC, FinCEN, and probably the FBI{{% /quote %}}.
{{< /fig-quote >}}

日本語でいうと「藪をつついて蛇を出す」。
個人的な印象で恐縮だが，信頼できる暗号化メッセージング・アプリケーションはもはや [Signal] が唯一と思っているので，マジで勘弁していただきたいところである。
そんなもん入れられるくらいなら普通に [Signal] に金を払うよ[^sec1]。

[^sec1]: {{% ruby "ただ" %}}無料{{% /ruby %}}ほど高い「安心安全」はない。

そもそも2018年に G20 が「それ」を[暗号資産（crypto-assets）と定義]({{< ref "/remark/2018/12/crypto-assets.md" >}} "「暗号資産」とやら")した時点で，私の中では「試合終了」だったのよ[^curr1]。
最近では「通貨」機能は諦めたのか NFT (Non-Fungible Token) が流行りらしいが， Xanadu 未満の駄アイデアにしか見えない。
追記型データベースとしての Blockchain も意外に使い道がなく，むしろ「環境にやさしくない」とか言われているし（笑）

[^curr1]: 何故 G20 が「暗号資産」と定義したか考えて欲しい。通貨としての信頼と安定がないからだよ。

いずれにしろデジタル・トークンを「通貨」とみなして流通させる季節はとっくに終わってるわけで，あとは「資産」としてのデジタル・トークンの運用で一山当てようという山師的な発想しか見えてこない。

デジタル・トークンの「資産運用」を全面否定するつもりはないが（私は関わりたくないが）， [Signal] のようなセキュリティまたはプライバシー上のクリティカル・サービスで山師的機能を付加する動きについては，今後は厳しく監視していく必要があるかもしれない。

## ブックマーク

- [「暗号通貨」ってゆーな！]({{< ref "/remark/2018/01/cryptocurrency-are-not-crypto.md" >}})
- [階級社会としてのインターネット]({{< ref "/remark/2018/05/internet-as-a-class-system.md" >}})
- [“Blockchain and Trust” by Bruce Schneier]({{< ref "/remark/2019/02/blockchain-and-trust-by-bruce-schneier.md" >}})

[Signal]: https://signal.org/

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
