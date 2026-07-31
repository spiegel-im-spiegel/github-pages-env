+++
title = "安全なデジタルライフのために"
date =  "2021-11-14T16:12:34+09:00"
description = "例によって Bruce Schneier 先生の記事経由なのだが，なかなか面白かったので，かいつまんで紹介してみる。"
image = "/images/attention/kitten.jpg"
tags = [ "security", "privacy", "risk", "management", "communication" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

例によって [Bruce Schneier 先生の記事](https://www.schneier.com/blog/archives/2021/11/advice-for-personal-digital-security.html "Advice for Personal Digital Security - Schneier on Security")経由なのだが，なかなか面白かったので，かいつまんで紹介してみる。

ちなみに私は「安心」を主観評価，「安全」を量的評価として区別している。
この記事では項目を列挙するだけに留めているが，本来はこの先に「量的評価」が存在する，という点を踏まえて読んでいただければ幸いである。
リスクは系全体で最小となるよう配分すべきである。

## 携帯端末について

まずは前半の記事。

- [Securing your digital life, part one: The basics | Ars Technica](https://arstechnica.com/features/2021/10/securing-your-digital-life-part-1/)

携帯端末は常に紛失・盗難のリスクがつきまとう。
紛失・盗難が起きないよう気をつけることはもちろんだが，万一そうなった場合でも最小限の被害で済むよう備えておく必要がある。

{{< fig-quote type="markdown" title="Securing your digital life, part one: The basics" link="https://arstechnica.com/features/2021/10/securing-your-digital-life-part-1/" lang="en" >}}
{{% quote %}}First, if you're not at home, you should always lock your device before you put it down, no exceptions. Your phone should be locked with the most secure method you're comfortable with—as long as it's not a 4-digit PIN, which isn't exactly useless but is definitely adjacent to uselessness. For better security, use a password or a passcode that's at least six characters long—and preferably longer. If you're using facial recognition or a fingerprint unlock on your phone, this shouldn't be too inconvenient{{% /quote %}}.
{{< /fig-quote >}}

きょうび PIN コードは6桁なんよねー。
携帯端末のアンロックが6桁なのに，いまだ銀行 ATM の暗証番号が4桁なのは何故だろう（笑）

{{< fig-quote type="markdown" title="Securing your digital life, part one: The basics" link="https://arstechnica.com/features/2021/10/securing-your-digital-life-part-1/" lang="en" >}}
{{% quote %}}Second, set your device to require a password immediately after it’s been locked. Delays mean someone who snatches your phone can get to your data if they bring up the screen in time{{% /quote %}}.
{{< /fig-quote >}}

えっ，最近のケータイってそんな要らん機能があるの？ 困るなぁ。

{{< fig-quote type="markdown" title="Securing your digital life, part one: The basics" link="https://arstechnica.com/features/2021/10/securing-your-digital-life-part-1/" lang="en" >}}
{{% quote %}}Additionally, make sure your device is set to erase its contents after 10 bad password attempts at maximum. This is especially important if you haven't set a longer passcode{{% /quote %}}.
{{< /fig-quote >}}

これは，そうだね。
心配なら[顔認識（facial recognition）]({{< ref "/remark/2021/10/glossary-about-face.md" >}} "「顔」に関する用語集")や指紋を使ったアンロックを普段使いにする手もある。

端末の紛失・盗難に備えてバックアップもしとけよ，って話だが

{{< fig-quote type="markdown" title="Securing your digital life, part one: The basics" link="https://arstechnica.com/features/2021/10/securing-your-digital-life-part-1/" lang="en" >}}
{{% quote %}}Also, regularly back up your phone. The safest way to back up data if you’re concerned about privacy is an encrypted backup to your personal computer; however, most iOS device owners can back up their data to iCloud with confidence that it is end-to-end encrypted (as long as they have iOS 13 or later){{% /quote %}}.
{{< /fig-quote >}}

とか書いてあるけど [iCloud へのバックアップは（プライバシー的には）安全とは言えない]({{< ref "/remark/2021/08/apples-mass-surveillance-plans.md" >}} "Apple 監視社会化計画（裏口を穿つ）")からね。
私みたいに [NAS 経由でクラウド・ストレージとの間でデータ暗号化]({{< ref "/remark/2021/10/nas.md" >}} "秋 NAS は俺に喰わせろ！")する手もあるが[^icloud]，あまり一般的じゃないよねぇ。
一番いいのは，メールにせよ写真にせよ動画にせよ，プライバシー的にヤバいものはケータイで扱わないことやね。
その上でクラウド・ストレージなりにバックアップしましょう，と。

[^icloud]: ちなみに Synology NAS の Cloud Sync は iCloud に対応してない。まぁ iCloud は可能な限り使わんことやね。

携帯端末の OS を最新にしましょうという話。

{{< fig-quote type="markdown" title="Securing your digital life, part one: The basics" link="https://arstechnica.com/features/2021/10/securing-your-digital-life-part-1/" lang="en" >}}
{{% quote %}}Along the same lines, make sure you have installed the most recent version of the phone OS available to prevent someone from taking advantage of known security bypasses{{% /quote %}}.
{{< /fig-quote >}}

これに続いて {{% quote lang="en" %}}For iOS, this is generally simple{{% /quote %}} とあるが，私は OS のアップデートを渋る iPhone ユーザを結構知っている。
Apple は機能追加や変更のついでみたいにセキュリティ・アップデートを行うが，これによって端末動作に不具合が起きることがままあるからだ。
まぁ Android の場合はキャリアや端末機器メーカーがアップデートを渋る場合があるので，さらに質が悪いのだが（笑） セキュリティ・リスクを減らすために OS を最新に保ってほしいのなら，ユーザがそのように行動するよう工夫すべきだと思うけどね。

これはいわゆる「野良アプリ」かな。

{{< fig-quote type="markdown" title="Securing your digital life, part one: The basics" link="https://arstechnica.com/features/2021/10/securing-your-digital-life-part-1/" lang="en" >}}
{{% quote %}}Side-loaded apps can also lead to security issues. Never side-load an app from an untrusted source or allow an iOS app that requires a “profile” to be installed on your device if the app isn't one you've created or one provided to you by your employer’s mobile device management (MDM) platform{{% /quote %}}.
{{< /fig-quote >}}

まぁ，これはそのとおり。
そして，こっちはアプリの「権限」の話。

{{< fig-quote type="markdown" title="Securing your digital life, part one: The basics" link="https://arstechnica.com/features/2021/10/securing-your-digital-life-part-1/" lang="en" >}}
{{% quote %}}To mitigate such vulnerabilities via apps, regularly review the permissions that applications request from the device. [...] Avoid apps with sketchy permission asks, and deny anything that seems like overreach—like when Facebook Messenger asks to be your SMS client and then logs all your phone calls to your Facebook account so it can find “friends” for you more efficiently. (Also, for the love of God, don't use Facebook Messenger.) And if there are apps that you don't use, delete them{{% /quote %}}.
{{< /fig-quote >}}

Facebook Messenger めった切りだな（笑） まぁ，私は随分前に Facebook Messenger は捨てたけど。

次は， Wi-Fi や Bluetooth を点けっぱにすんなって話。ちょっと長いけど重要なのでそのまま挙げておく。

{{< fig-quote type="markdown" title="Securing your digital life, part one: The basics" link="https://arstechnica.com/features/2021/10/securing-your-digital-life-part-1/" lang="en" >}}
Besides issues that arise from questionable app behavior, mobile devices can be vulnerable through normal functions like Wi-Fi or Bluetooth. Consider turning off Wi-Fi when you’re away from home. Your device may otherwise be constantly polling for the network SSIDs in its history to reconnect automatically or to connect to anything that looks like a carrier’s Wi-Fi network. When this happens, your device gives away information about networks you’ve seen and might allow a hostile network access point to connect. Also, your phone's Wi-Fi MAC address could be used to fingerprint your device and track it. (Apple randomizes the MAC address of its iOS devices’ Wi-Fi adapters while scanning for networks—but if your home Wi-Fi network’s name is particularly memorable, that may not matter.) When your phone tells you to turn on Wi-Fi to improve location accuracy, ignore it.

The same goes for Bluetooth. If your device has Bluetooth turned on, it’s broadcasting information that could identify it—and you. (I have demonstrated this to journalism classes by calling out students' names that I picked out from the default names of their iPhones.)
{{< /fig-quote >}}

これ，分かってるけどつい忘れちゃうんだよねぇ。
本当は移動中にワイヤレスのヘッドセットを付けるのも止めたほうがいいんだろうけど。

## パソコンとか

パソコン OS や Web ブラウザの最新化は言わずもがなだろう。

{{< fig-quote type="markdown" title="Securing your digital life, part one: The basics" link="https://arstechnica.com/features/2021/10/securing-your-digital-life-part-1/" lang="en" >}}
{{% quote %}}I have found several common themes when things go wrong; the biggest is that malware protection is frequently not up to date or, worse, is disabled. Even allowing Windows Defender to run in the background provides a significant bump in protection over nothing, and disabling it without a very good reason is a very bad idea{{% /quote %}}.
{{< /fig-quote >}}

最低でも Windows Defender は有効にしておけ，ってことやね。

このあと具体的な脅威が列挙されているが，それは本文を見てもらうとして

{{< fig-quote type="markdown" title="Securing your digital life, part one: The basics" link="https://arstechnica.com/features/2021/10/securing-your-digital-life-part-1/" lang="en" >}}
{{% quote %}}The basic fixes for these threats are straightforward but require some behavior modification. And one of the easiest behaviors to modify is how we browse the web. We need to treat this the same way we would a walk home in the dark—with extra, active attention to our surroundings{{% /quote %}}.
{{< /fig-quote >}}

と続く。
面倒な時代になったよねぇ。
Web ブラウジングについては以下のような緩和措置もある。

{{< fig-quote type="markdown" title="Securing your digital life, part one: The basics" link="https://arstechnica.com/features/2021/10/securing-your-digital-life-part-1/" lang="en" >}}
{{% quote %}}Another helpful trick for keeping laptops safe is using a privacy browser plugin such as the [EFF’s Privacy Badger](https://privacybadger.org/) and being very selective about which sites are allowed to use tracking cookies. This will reduce the threat of [malicious ads](https://news.sophos.com/en-us/2020/09/09/faking-it-the-thriving-business-of-fake-alert-web-scams/){{% /quote %}}.
{{< /fig-quote >}}

[Privacy Badger](https://privacybadger.org/) 自体は昔からあるようで，以下の記事でも紹介されている。

- [監視されているウェブ閲覧の足跡、ブラウザ拡張機能で対策を | ギズモード・ジャパン](https://www.gizmodo.jp/2020/09/untracking-browser-extensions.html)
- [Doxxing（晒し）からあなたを守るための手引き | P2Pとかその辺のお話R](https://p2ptk.org/privacy/3425)

単純な広告ブロックではない点がポイントかな。
んー， Android 端末の Firefox でしばらく使ってみようかな。

そして更に以下に続く。

{{< fig-quote type="markdown" title="Securing your digital life, part one: The basics" link="https://arstechnica.com/features/2021/10/securing-your-digital-life-part-1/" lang="en" >}}
{{% quote %}}Another easy way to minimize threats to your PC, first and foremost, is running the most recent fully updated version of the operating system of your choice. I’m not going to advocate for any particular flavor of operating system here, but if you’re on an older release of any OS and connected to the Internet, you’re increasing your risk of compromise. Turn on automatic updates and leave them on. When an update is pending, stop what you're doing and install it immediately. Yes, this can often be inconvenient. Welcome to the modern world of malware. Suck it up and install your updates or risk compromise. (This applies to your web browser, too—stop putting off that Chrome update prompt and do it right now.){{% /quote %}}
{{< /fig-quote >}}

今世紀に入って可用性（availablity）へのインパクトがセキュリティ・リスクとして取り上げられるようになって，迂闊に「すぐに更新しろ」とか言えなくなってしまったが，個人で使ってる携帯端末やパソコンは**可及的速やかに**更新するという考え方でいいだろう。
アップデートは計画的に。

パソコンも携帯端末と同じようにちゃんとパスワードロックしておけよ，という話。

{{< fig-quote type="markdown" title="Securing your digital life, part one: The basics" link="https://arstechnica.com/features/2021/10/securing-your-digital-life-part-1/" lang="en" >}}
{{% quote %}}Just as a phone's solid unlock password prevents data theft, the same is true of enabling password or PIN protection on your notebook computer for sleep mode. When traveling in high-risk areas like airports, power-down your computer when it’s not in use so that the risk of someone playing “[Evil Maid](https://en.wikipedia.org/wiki/Evil_maid_attack)” or surreptitiously gaining access in some other physical way is reduced{{% /quote %}}.
{{< /fig-quote >}}

“[Evil maid attack](https://en.wikipedia.org/wiki/Evil_maid_attack)” なんてな名前があるのか。
知らんかった（笑）

パスワード・マネージャについても言及されている。

{{< fig-quote type="markdown" title="Securing your digital life, part one: The basics" link="https://arstechnica.com/features/2021/10/securing-your-digital-life-part-1/" lang="en" >}}
{{% quote %}}And, finally, use a password manager. An easy-to-guess password renders all other security efforts moot. Whether it’s a password built into your web browser of choice or a standalone program, use one. Chrome, Firefox, and Safari all have reasonably secure password managers, and you can replicate passwords for web accounts across devices. If you don't like the idea of a password manager because you're one of those folks who just uses letmein123! as your password everywhere, you need to decide if the convenience is worth the price you'll eventually pay when you're compromised. (Spoiler alert: it's not.){{% /quote %}}
{{< /fig-quote >}}

この辺は改めて言うことはない。
「[パスワードを覚えるなんて脳みその無駄遣い](https://baldanders.info/blog/000739/ "「パスワードを覚える」なんて脳みその無駄遣い — Baldanders.info")」である。

## 情報管理の Best Practices

それでは後半の記事を見てみよう。

- [Securing your digital life, part two: The bigger picture—and special circumstances | Ars Technica](https://arstechnica.com/information-technology/2021/10/securing-your-digital-life-part-2/)

最初のほうでは “These are some best practices to consider” として箇条書されている。
見出しだけ抜き出して番号をつけてみよう。

{{< div-gen type="markdown" lang="en" >}}
1. Use a password manager that generates strong passwords you don’t have to remember.
2. When possible, use two-factor or multi-factor authentication ("2FA" or "MFA").
3. Set up a separate email address or email alias for your high-value web accounts so that all email regarding them is segmented off from your usual email address.
4. If you're a US resident, make sure to [claim an account for your Social Security number](https://www.irs.gov/payments/view-your-tax-account) from the IRS for tax information access and other purposes.
5. Register for account breach checkups, either through the service provided through your browser (Firefox or Chrome) or through Troy Hunt’s [haveIbeenpwned.com](https://haveibeenpwned.com/) (or both!).
6. Consider locking your credit reports to reduce identity theft risks.
{{< /div-gen>}}

4番目は分かりにくいが，以下の記事が参考になるかも知れない。

- [Plant Your Flag, Mark Your Territory —  Krebs on Security](https://krebsonsecurity.com/2018/06/plant-your-flag-mark-your-territory/)

つまり，なりすましなどに遭う前に「旗を立て」て証明しろというわけだ。
日本で類似の話としては JP の「e転居」を悪用した事例がある。

- [日本郵便のe転居を悪用したストーカー事件についてまとめてみた - piyolog](https://piyolog.hatenadiary.jp/entry/2020/07/17/174628)

まぁ，これは「e転居」の仕組みが ad hoc 過ぎるのが問題だと思うが。

これで連想するのが日本の「個人番号カード」だ。
今はあまり表面化してないが，日本政府が無駄にを推進しようとしているせいで今後「個人番号カード」そのものがセキュリティ・リスクになると思う。

今からでも「個人番号カード」を取得しようという人は「とっとと手続きしなはれ」と言えばいいが，特に年寄りは「個人番号カードなんか分からんし使いこなせん」と思う人もいると思う。
私は役所に「絶対に個人番号カードを発行しない」手続きを行える仕組みを提案したい。

## 2要素認証で不十分なとき

“When 2FA is not enough” については似た話を[先日書いた]({{< ref "/remark/2021/11/out-of-band-devices.md" >}} "それはワンタイム・パスワードの問題ではない")ので，今回は割愛する。

## 公衆への暴露

標的型攻撃の（文字通り）ターゲットにならないようにするには個人活動の露出を抑える必要がある。
これも箇条書きにして抜き出しておこう。

{{< div-gen type="markdown" lang="en" >}}
1. don’t post unredacted pictures of driver's licenses, vaccination cards, credit cards, passports, or other documents with PII on social media.
2. Lock down access to your social media accounts with 2FA and unique, strong passwords to prevent "brute force" breaches and "password reuse" hacks.
3. On Facebook, set the default privacy for posts to “friends only."
4. Do not use “precise location” information on posts that can be used to locate you in realtime.
5. Don’t post pictures with your home address or other identifying information about your residence clearly visible.
6. Don’t drop personal email addresses or phone numbers into public online conversations.
7. Don’t allow dating apps, ride-sharing apps, or any other apps that use your location data to collect that data while you’re not actively using them.
8. If you are sending your location to someone in one of these apps, make sure it is a public place and that a friend or family member is in sight of that location, or at least knows to check in with you shortly after the appointed meeting time.
9. Never take a conversation in one app over to another—say, from a dating app to WhatsApp—before you’ve met a person in person and feel safe.
10. Be aware of links sent in Facebook Messenger and of friend requests claiming to be from people you already know—but coming from new accounts.
11. Don’t download and run anything from Discord without a malware scan.
{{< /div-gen>}}

実はこのあとオチがあって

{{< fig-quote type="markdown" title="Securing your digital life, part two: The bigger picture—and special circumstances" link="https://arstechnica.com/information-technology/2021/10/securing-your-digital-life-part-2/" lang="en" >}}
{{% quote %}}Or, instead of trying to follow all these suggestions, you could mitigate this entire category of risks by never using social media again{{% /quote %}}.
{{< /fig-quote >}}

と結んでいる。
いや，まぁ，それはそうだけどさ（笑）

## 特殊ケース

最後に “Special cases” として VPN (Virtual Private Network) や暗号化コミュニケーションについて以下のように述べている。

{{< fig-quote type="markdown" title="Securing your digital life, part two: The bigger picture—and special circumstances" link="https://arstechnica.com/information-technology/2021/10/securing-your-digital-life-part-2/" lang="en" >}}
But for everyday Internetting, you just don’t need VPNs that much anymore. Transport Layer Security now encrypts a vast majority of Internet traffic, and it’s unlikely that someone is going to grab your credit card data or other personal information off a public Wi-Fi network.

The same is true of the Tor protocol for anonymizing Internet traffic—odds are you won’t need it daily, but there are times when it’s good to have. Tor and VPNs are most useful when you're outside of your home and on a potentially hostile network (or on the Internet in a potentially hostile country).

You’ll also want Tor or a VPN in situations where you’re on a network that has a TLS proxy that breaks traditional HTTPS encryption by using proxy certificates to decrypt traffic in the middle. At least in those scenarios, the worst that can happen is you can’t get an outbound connection.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="Securing your digital life, part two: The bigger picture—and special circumstances" link="https://arstechnica.com/information-technology/2021/10/securing-your-digital-life-part-2/" lang="en" >}}
{{% quote %}}For encrypted and verified communications with another specific person, [Signal](https://signal.org/en/) is the current standard—it’s cross-platform and doesn’t even have an option for unencrypted storage or transmission of text and voice. Lesser-known platforms such as [Keybase](https://keybase.io/) and [Wire](https://wire.com/) offer encrypted text communications as well, but a full discussion of encrypted voice and text communications is a subject for another time{{% /quote %}}.
{{< /fig-quote >}}

ところで，最後の “Peanut Butter Sandwich Anonymization Protocol” ちうのが分からんのだが，どなたか教えてください。

{{< div-box type="markdown" >}}
**【2021-11-15 追記】**

[Twitter で教えてもらった](https://twitter.com/atsushieno/status/1459971971299479556)が “Peanut Butter Sandwich Anonymization Protocol” ってのは以下の話のことらしい？

- [Couple Arrested After Spreading U.S. Navy Secrets Via a Peanut Butter Sandwich - Thrillist](https://www.thrillist.com/news/nation/couple-arrested-after-spreading-us-navy-secrets-via-a-peanut-butter-sandwich)

情報感謝！
{{< /div-box >}}

## 【あとで整理するかも】

- [Securing your digital life, part three: How smartphones make us vulnerable | Ars Technica](https://arstechnica.com/information-technology/2021/11/securing-your-digital-life-part-3/)
  - [Securing Your Smartphone - Schneier on Security](https://www.schneier.com/blog/archives/2021/11/securing-your-smartphone.html)

## 参考図書

{{% review-paapi "B01MZGVHOA" %}} <!-- 超監視社会 -->
{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
