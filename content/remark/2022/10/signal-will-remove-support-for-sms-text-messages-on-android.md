+++
title = "Android 版 Signal が SMS 対応を取りやめる話"
date =  "2022-10-16T21:51:56+09:00"
description = "まぁ，これはしょうがない。"
image = "/images/attention/kitten.jpg"
tags = [ "security", "privacy", "risk", "communication", "messaging" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先日 [Signal] が Android 版の SMS 対応を取りやめる旨の記事を公開した。

- [Signal >> Blog >> Removing SMS support from Signal Android (soon)](https://signal.org/blog/sms-removal-android/)

私も TextSecure の頃から SMS の既定アプリを切り替えて運用してる。
確かに SMS の通信経路にはセキュリティ・リスクがあるが，メッセージング・アプリは相手がいてはじめて成立するものだし， SMS の通信内容をローカルでは暗号化して保存してくれるので「素の SMS アプリよりはマシ」程度で使っていた。

しかし，ここに来て [Signal] をして {{% quote lang="en" %}}We have now reached the point where SMS support no longer makes sense{{% /quote %}} とまで言わしめている。

[Signal] が SMS への対応を取りやめる理由として3つ挙げている。
まずひとつめ。

{{< fig-quote type="markdown" title="Removing SMS support from Signal Android (soon)" link="https://signal.org/blog/sms-removal-android/" lang="en" >}}
The most important reason for us to remove SMS support from Android is that plaintext SMS messages are [inherently insecure](https://signal.org/blog/goodbye-encrypted-sms). They leak sensitive metadata and place your data in the hands of telecommunications companies. With privacy and security at the heart of what we do, letting a deeply insecure messaging protocol have a place in the Signal interface is inconsistent with our values and with what people expect when they open Signal.
{{< /fig-quote >}}

まぁ，これはそうだよねって話。

ふたつめはコスト。

{{< fig-quote type="markdown" title="Removing SMS support from Signal Android (soon)" link="https://signal.org/blog/sms-removal-android/" lang="en" >}}
Back when we started supporting plaintext SMS messaging things were different. Data plans were much more expensive generally, and were totally inaccessible in many parts of the world. Now, data plans are cheaper and far more ubiquitous than they were nearly a decade ago. In a reversal, the cost of sending SMS is now prohibitively high in many parts of the world. This brings us to our second reason: we’ve heard repeatedly from people who’ve been hit with high messaging fees after assuming that the SMS messages they were sending were Signal messages, only to find out that they were using SMS, and being charged by their telecom provider. This is a terrible experience with real consequences.
{{< /fig-quote >}}

まじか。
そりゃあ，アカンね。

みっつめは UX の観点から。

{{< fig-quote type="markdown" title="Removing SMS support from Signal Android (soon)" link="https://signal.org/blog/sms-removal-android/" lang="en" >}}
Third, there are serious UX and design implications to inviting SMS messages to live beside Signal messages in the Signal interface. It’s important that people don’t mistake SMS messages sent or received via the Signal interface as secure and private when in fact they are not. And while we flag the difference between them in the app, we can only do so much on the design side to prevent such misunderstandings.
{{< /fig-quote >}}

というわけで

{{< fig-quote type="markdown" title="Removing SMS support from Signal Android (soon)" link="https://signal.org/blog/sms-removal-android/" lang="en" >}}
After much discussion, we determined that we can no longer continue to invest in accommodating SMS in the Android app while also dedicating the resources we need to make Signal the best messenger out there.
{{< /fig-quote >}}

という結論に至ったそうな。
まぁ，これはしょうがない。
SMS アプリはキャリアが用意しているものに戻すか。

ところで最初の理由で SMS が {{% quote lang="en" %}}inherently insecure{{% /quote %}} であるとはどういうことか。
これに関しては以下の記事が参考になる。

- [Using Your Phone in Times of Crisis | Electronic Frontier Foundation](https://www.eff.org/deeplinks/2022/03/using-your-phone-times-crisis)
- [危機的状況下でのモバイルネットワークの使用がもたらすセキュリティリスク | p2ptk[.]org](https://p2ptk.org/security/3488)

原文は2022年3月に公開されたものだが，今回の [Signal] の発表を受けて [p2ptk.org](https://p2ptk.org/ "P2Pとかその辺のお話R | Sharing is Caring") の中の人が翻訳記事を公開してくださった。
ありがたや {{% emoji "ペコン" %}}

ぶっちゃけていうと SMS というのは [2G 時代の遺物を継承]({{< ref "/remark/2020/01/5g-security-risk.md" >}} "5G セキュリティ・リスク")したもので「個人のセキュリティやプライバシーなんか知ったことか」というつくりになっている。

{{< fig-quote type="markdown" title="危機的状況下でのモバイルネットワークの使用がもたらすセキュリティリスク" link="https://p2ptk.org/security/3488" >}}
特に、最古の携帯電話ネットワークの[2G通信](https://www.eff.org/deeplinks/2020/06/your-phone-vulnerable-because-2g-it-doesnt-have-be)では、通話やテキストメッセージは簡単に傍受されてしまう。そのため、我々はAppleとGoogleに対し、ユーザが（訳注：AndroidやiOSで）2Gをオフにできる機能を提供するよう要請してきた。Googleは最新の端末にこのオプションを[導入した](https://www.eff.org/deeplinks/2022/01/victory-google-releases-disable-2g-feature-new-android-smartphones)が、ロシアとウクライナではほぼ利用できない。またAppleは未だに対応していない。

我々は可能な限り2Gの利用をやめるよう促してきたが、3G、4G、5Gだからセキュアだというわけではない。とりわけ、音声とテキスト通信にとってはセキュアなオプションとは言えない。これらのネットワークでは、通信がエンドツーエンドで暗号化されていないため、通信事業者を含め、通信の傍受者が通信内容を見聞きできてしまう。
{{< /fig-quote >}}

えー。
Apple って未だにそんなんなの？ まぁ，そういうわけで

{{< fig-quote type="markdown" title="危機的状況下でのモバイルネットワークの使用がもたらすセキュリティリスク" link="https://p2ptk.org/security/3488" >}}
あなたがどこに住んでいようと、特にロシアとウクライナにおいては、政府当局からあなたの通信のプライバシーを守るために、電話やSMSに依存してはならない。
{{< /fig-quote >}}

となるわけだ。
これって必然的に

{{< fig-quote type="markdown" title="危機的状況下でのモバイルネットワークの使用がもたらすセキュリティリスク" link="https://p2ptk.org/security/3488" >}}
[2要素認証（2FA](https://www.eff.org/deeplinks/2016/12/12-days-2fa-how-enable-two-factor-authentication-your-online-accounts)：アカウントにログインする際に取得するコード）については、できる限りSMSではなくアプリを使用することを推奨する。
{{< /fig-quote >}}

という話になるのだが，今だに SMS でシークレットを送りつけるあのサービスやあのサービスは個人のセキュリティやプライバシーについてどう考えているか知りたいものである。
あとは [SP 800-63-3](https://pages.nist.gov/800-63-3/ "NIST SP 800-63 Digital Identity Guidelines") を改訂して，はっきり SMS は Authenticator としては deprecated だと言ってもらえるとみんな諦めがつくと思うのだが（笑）


## ブックマーク

- [Signal will remove support for SMS text messages on Android](https://www.bleepingcomputer.com/news/technology/signal-will-remove-support-for-sms-text-messages-on-android/)
- [Why Signal won’t compromise on encryption, with president Meredith Whittaker - The Verge](https://www.theverge.com/23409716/signal-encryption-messaging-sms-meredith-whittaker-imessage-whatsapp-china)
  - [メッセージングアプリSignalが暗号化で妥協しない理由を新プレジデントが語る - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20221031/signal-encryption)
  - [ブログ: Signalが暗号化で妥協しない理由、メレディス・ウィテカー社長に聞く](https://okuranagaimo.blogspot.com/2022/11/signal.html)

- [Authenticator と AAL]({{< ref "/remark/2020/09/authenticator-and-aal.md" >}})

[Signal]: https://signal.org/

## 参考文献

{{% review-paapi "4822281426" %}} <!-- 情報セキュリティ技術大全 -->
