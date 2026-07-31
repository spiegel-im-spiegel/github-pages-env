+++
title = "「Signal は安全か？」"
date =  "2025-12-12T14:11:11+09:00"
description = "もはや米国および米国ビッグテックへの依存は，セキュリティやプライバシーの観点から，リスクと見なされていることが分かる。"
image = "/images/attention/kitten.jpg"
tags = [ "security", "privacy", "cryptography", "messaging", "tools", "signal" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

Proton Mail や Proton VPN などでおなじみ [Proton] が，暗号化メッセージングアプリ [Signal] を評価する記事を公開している。

- [Is Signal safe? What to know about this encrypted messaging app | Proton](https://proton.me/blog/is-signal-safe)

今頃こんな記事が出る背景のひとつとして，今年の春に話題になった通称 “Signalgate” 事件があると思われる。
この辺の話は以前記事にした。

- [Signal グループチャット内容の漏洩と NSA]({{< ref "/remark/2025/04/the-signal-chat-leak-and-the-nsa.md" >}})

“[Is Signal safe?]” でも以下のように言及されている。

{{< fig-quote type="markdown" title="Is Signal safe? What to know about this encrypted messaging app" link="https://proton.me/blog/is-signal-safe" lang="en" >}}
In a stunning breach of national security, a reporter for The Atlantic was inadvertently invited to join the group chat and soon made a partially redacted transcript of it available to the public. Crucially, the entire “Signalgate” scandal was the result of purely human error, not any technical flaw in Signal.
{{< /fig-quote >}}

ことの重大性に比してかなり間抜けな事件だと個人的には思うが，その後も

{{< fig-quote type="markdown" title="Is Signal safe? What to know about this encrypted messaging app" link="https://proton.me/blog/is-signal-safe" lang="en" >}}
A similar issue appears to be responsible for a successful [FBI spying operation](https://www.theguardian.com/us-news/2025/nov/21/fbi-signal-group-chat-immigration) that targeted an immigrants’ rights activists Signal group chat. Although details are not yet clear, the FBI said the information came from a “sensitive source with excellent access”, which strongly suggests that an inside source had been invited to the chat.
{{< /fig-quote >}}

みたいな事例が発覚してるそうで，オンライン・コミュニティにスパイを紛れ込ませて秘匿情報を抜いたり悪意のコードを組み込む手法は，もはや常套手段になっているのかもしれない。

もうひとつ “[Is Signal safe?]” で面白いと思ったのは [Signal] を含めた暗号化メッセージングアプリの比較である。
記事では以下の観点で比較が行われている。

- Messages and calls encrypted in transit and at rest (メッセージと通話は転送中および保存時に暗号化)
- Messages and calls encrypted end to end (メッセージと通話はエンドツーエンドで暗号化)
- Metadata minimized (メタデータの最小化)
- No sharing of your data (データを共有しない)
- Open source and audited (オープンソースで監査済み)
- Jurisdiction (法的管轄)

比較表はこんな感じ。

{{< fig-img-quote src="./how-safe-is-signal_8974500fe6.webp" title="How secure and private is Signal?" link="https://proton.me/blog/is-signal-safe#why" lang="en" width="1024" >}}

まぁ [Signal] が基準なので， これが（Jurisdiction を除いて）満点なのは当然だが，特筆すべきは [Threema] だろう。
[Threema] は WhatsApp の代替アプリとして有名なようで，[件の記事][Is Signal safe?]でも

{{< fig-quote type="markdown" title="Is Signal safe? What to know about this encrypted messaging app" link="https://proton.me/blog/is-signal-safe" lang="en" >}}
All Threema apps use the open-source [NaCl cryptography library](https://nacl.cr.yp.to/) to encrypt messages end to end, and they have been [audited by security professionals](https://threema.com/en/faq/code-audit). Unlike most messaging apps, you don’t need an email address or phone number to register an account, and it’s possible to purchase Threema for Android anonymously using Bitcoin —Threema is a paid app. The company says this allows you to text and make calls anonymously, and it goes to lengths to ensure that it [collects minimum metadata](https://threema.ch/en/blog/posts/metadata).
{{< /fig-quote >}}

と有料アプリながら絶賛している[^t1]。

[^t1]: [Proton] の別の記事 “[Best WhatsApp alternatives for privacy](https://proton.me/blog/whatsapp-alternatives "Best WhatsApp alternatives for privacy | Proton Mail | Proton")” でも [Threema] を最高評価している。

[Threema] と比較してということになるだろうが [Signal] への懸念として

{{< fig-quote type="markdown" title="Is Signal safe? What to know about this encrypted messaging app" link="https://proton.me/blog/is-signal-safe" lang="en" >}}
However, being hosted on AWS servers remains a concern in light of Signal’s reliance on SGX. There are a number of open-source [encrypted messaging apps](https://proton.me/blog/whatsapp-alternatives) like Threema that try to address this and other perceived issues with Signal — such as its reliance on a centralized server and the need to supply a real phone number— some of which show great promise.
{{< /fig-quote >}}

と述べていて，もはや米国および米国ビッグテックへの依存は，セキュリティやプライバシーの観点から，リスクと見なされていることが分かる。

とはいえ，ユーザ規模や外部からの厳密なセキュリティ精査といった面を考えれば，暗号化メッセージングアプリの選択肢として上位にあるのは変わらないだろう。
メッセージングアプリは相手がいて初めて成立するしね。
その上で [Threema] のような選択肢があることも頭の隅に入れておくのがいいかもしれない。

[Proton]: https://proton.me/ "Proton: Privacy by default"
[Is Signal safe?]: https://proton.me/blog/is-signal-safe "Is Signal safe? What to know about this encrypted messaging app | Proton"
[Signal]: https://signal.org/ "Signal"
[Threema]: https://threema.com/en "Threema – Highly Secure Communication For Individuals and Companies – Threema"

## 参考図書

{{% review-paapi "B01MZGVHOA" %}} <!-- 超監視社会 -->
{{% review-tatsujin "infoshare2" %}} <!-- 続・情報共有の未来 -->
