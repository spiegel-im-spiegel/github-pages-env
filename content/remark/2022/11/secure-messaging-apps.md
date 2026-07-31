+++
title = "いまさら「安全なメッセージング・アプリ」"
date =  "2022-11-10T21:09:00+09:00"
description = "私から「Signal を入れよう」と持ちかけてもいぢめないであげてください。"
image = "/images/attention/kitten.jpg"
tags = [ "security", "privacy", "risk", "management", "messaging", "communication" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Twitter TL を眺めてたら Mastodon で [{{% quote lang="en" %}}Direct Messages on Mastodon are NOT encrypted{{% /quote %}}](https://mastodon.world/@ruud/109315777955924243) みたいなのが流れてるんだそうで，そもそも Twitter の DM だって[安全じゃない](https://edition.cnn.com/2022/04/30/tech/twitter-dms-encrypted-elon-musk/ "Elon Musk wants end-to-end encryption for Twitter DMs. It may not be that simple | CNN Business")し，複数のサーバが分散している Mastodon なら尚更だと思うのだが，世の中の認識はどんな感じなのだろう。

大昔のカビの生えた記事で申し訳ないが，2014年に EFF が “[Secure Messaging Scorecard]” という記事を公開していて，この中で「安全なメッセージング・アプリ」の7項目を挙げている。
曰く

1. Encrypted in transit?
1. Encrypted so the provider can't read it?
1. Can you verify contacts' identities?
1. Are past comms secure if your keys are stolen? 
1. Is the code open to independent review?
1. Is security design properly documented?
1. Has there been any recent code audit?

である。
当時は私も[便乗記事](https://baldanders.info/blog/000782/ "安全なメッセージング・アプリとは")を書いていたが，その後 EFF は記事をいったん引っ込めている。
現在，跡地として残っている記事には

{{< fig-quote type="markdown" title="Secure Messaging Scorecard | Electronic Frontier Foundation" link="https://www.eff.org/pages/secure-messaging-scorecard" lang="en" >}}
This is version 1.0 of our scorecard; it is out of date, and is preserved here for purely historical reasons. Please visit [Surveillance Self-Defense](https://ssd.eff.org/) if you're looking for recommendations on specific tools to use to ensure your privacy and security while we work on writing an updated guide to secure messaging. Again, you should not use this scorecard to evaluate the security of any of the listed tools, as many of them have been changed since the scorecard was last updated (some for the better, some for the worse).
{{< /fig-quote >}}

と但し書きがされているのでご注意を。
ただ，まぁ，上に挙げた7項目は「安全なメッセージング・アプリ」を実現するための最低限の条件と言っていいだろう。

Twitter や Mastodon の DM は電子メール・システムに近い。
ユーザ間を直接的に繋ぐ通信ではなく，最低ひとつのサーバにメッセージが保持され，サービス・プロバイダ側の「なにか」がそのメッセージを見ることは（技術的には）可能である。
簡単に言うと，これらは「はがき」と同程度に「検閲可能」なのだ。

安全を確保するためのやり方は色々ある。
たとえば，電子メールでは S/MIME や PGP/MIME といった仕組みでメッセージ本文を暗号化できる。
今は廃れているかもしれないが，かつて Jabber とよばれた XMPP 方式のインスタント・メッセージング・サービスは [OTR (Off-the-Recording)](https://wiki.xmpp.org/web/OTR) を組み込むことで，上述の7項目を満たすことに成功した。

個人的には [Signal] を推す。
上の “[Secure Messaging Scorecard]” 跡地で紹介されている “[Surveillance Self-Defense]” でもメッセージング・アプリとして [Signal] と OTR と WhatsApp が紹介されている。

とはいうものの，8年前の記事でも書いたが，結局メッセージング・アプリで一番大事なのは「相手がいること」なのよね。
一応 [Signal] は入れてるんだよ。
誰も遊んでくれないけど（笑） 仕事では Microsoft Teams を使ってるし，広島の友人との連絡は Facebook Messenger だし，家族との連絡は電話かキャリアメールだ。
まぁ LINE は今だに毛嫌いしてるので，知り合いから「LINE 入れないの？」と言われても頑なに断ってるけど。

というわけで，私から「[Signal] を入れよう」と持ちかけてもいぢめないであげてください。

## ブックマーク

- [メッセージングアプリSignalが暗号化で妥協しない理由を新プレジデントが語る - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20221031/signal-encryption) : ちなみにこの記事で [Signal] の送金機能の話が出ているが，ベータ機能として組み込まれている。私は有効にしてないけど

- [Android 版 Signal が SMS 対応を取りやめる話]({{< ref "/remark/2022/10/signal-will-remove-support-for-sms-text-messages-on-android.md" >}})

[Secure Messaging Scorecard]: https://www.eff.org/pages/secure-messaging-scorecard "Secure Messaging Scorecard | Electronic Frontier Foundation"
[Surveillance Self-Defense]: https://ssd.eff.org/ "Surveillance Self-Defense | Tips, Tools and How-tos for Safer Online Communications"
[Signal]: https://signal.org/

## 参考文献

<!-- eof -->
{{% review-paapi "B01MZGVHOA" %}} <!-- 超監視社会 -->
{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
