+++
title = "SMS 認証がダメな理由"
date =  "2026-11-13T15:51:38+09:00"
description = "Authenticator と SMS 認証 / SIM スワップと SMS 認証のリスク / SMS 認証に代えて / SIM スワップへの対応は最初の1時間が肝心らしい / 【おまけ】電子メールを Out-of-Band 認証に使ってはいけない"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "security", "authentication", "risk", "management", "nist", "k-tai" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

今回は Proton Blog の以下の記事から。

- [How to protect your business from SIM swapping | Proton][How to protect your business from SIM swapping]

例によって Proton 製品の広告を含んでるので気をつけてね。
記事自体は主に企業向けの内容となってるけど，個人でも知識として参考になる部分は多いと思う。

## Authenticator と SMS 認証

"[NIST SP 800-63 Digital Identity Guidelines][NIST SP 800-63]" では Authenticator (認証器？) を以下のように分類している。

- [Passwords (3.1.1)](https://pages.nist.gov/800-63-4/sp800-63b/authenticators/#password)
- [Look-Up Secrets (3.1.2)](https://pages.nist.gov/800-63-4/sp800-63b/authenticators/#lookupsecrets)
- [Out-of-Band Devices (3.1.3)](https://pages.nist.gov/800-63-4/sp800-63b/authenticators/#out-of-band)
- [Single-Factor OTP (3.1.4)](https://pages.nist.gov/800-63-4/sp800-63b/authenticators/#singlefactorOTP)
- [Multi-Factor OTPs (3.1.5)](https://pages.nist.gov/800-63-4/sp800-63b/authenticators/#multifactorOTP)
- [Single-Factor Cryptographic Authentication (3.1.6)](https://pages.nist.gov/800-63-4/sp800-63b/authenticators/#sfc)
- [Multi-Factor Cryptographic Authentication (3.1.7)](https://pages.nist.gov/800-63-4/sp800-63b/authenticators/#mfc)

私たちが普段「パスキー認証」と呼んでいるものは，ほぼ同期型のパスキーで，[NIST SP 800-63] では "[Syncable Authenticators (B)](https://pages.nist.gov/800-63-4/sp800-63b/syncable/)" に該当するみたい[^pa1]。
上の分類で言うなら "Single-Factor Cryptographic Authentication" または "Multi-Factor Cryptographic Authentication" に相当するようだ（実装による）。

[^pa1]: パスキー認証は FIDO2 (WebAuthn/CTAP2) に基づく認証方式で，秘密鍵をユーザ側の端末（または CTAP2 でやり取りできる Authenticator）内に保持し，公開鍵をサービス・サーバに登録することで認証を行う。実はパスキー認証には同期型とデバイス固定型があり，同期型は端末間で秘密鍵を同期できる。同期型のパスキー認証は実装条件によっては AAL (Authenticator Assurance Level) の 2 は達成し得るが AAL3 としては不適格とされている（実質的に秘密鍵をエクスポートできるため）。私たちが普段「パスキー認証」と呼んでいるものは，この同期型のことを指すことが多いため "Syncable Authenticators" の章が設けられていると思われる。

SMS を含む PSTN (Public Switched Telephone Network) を使った認証は "Out-of-Band Devices" に分類されるが，制限付き認証器（[Restricted Authenticators (3.2.9)](https://pages.nist.gov/800-63-4/sp800-63b/authenticators/#restricted)）とみなされている。また，NIST は SIM 変更や番号移行などのリスク指標を考慮するよう推奨している。
また以下の引用のように

{{< fig-quote type="markdown" title="3.1.3.3. Authentication Using the Public Switched Telephone Network" link="https://pages.nist.gov/800-63-4/sp800-63b/authenticators/#pstnOOB" lang="en" >}}
Some subscribers may be unable to use PSTN to deliver out-of-band authentication secrets in areas with limited telephone coverage, particularly without mobile phone service. Accordingly, verifiers **SHALL** ensure that alternative authenticator types are available to all subscribers and **SHOULD** remind subscribers of this limitation of PSTN out-of-band authenticators before binding one or more devices controlled by the subscriber.
{{< /fig-quote >}}

電話ベースの認証を使えないユーザのために代替手段を用意することも求めている。

と，以上が前置き。
これを踏まえて最初の記事を読んでいく。

## SIM スワップと SMS 認証のリスク

最初の記事では SIM スワップによって SMS 認証が突破されるリスクについて言及されている。

{{< fig-quote type="markdown" title="How to protect your business from SIM swapping" link="https://proton.me/business/blog/sim-swapping" lang="en" >}}
Most security controls within your business can be improved, such as passwords, training, and access permissions. However, SMS-based [two-factor authentication](https://proton.me/blog/what-is-two-factor-authentication-2fa) (2FA) is different, because your business can’t make it stronger. The codes are only as secure as your phone number, and your mobile carrier controls who that number belongs to. Its weakest point is the carrier’s customer support agents.
{{< /fig-quote >}}

攻撃者は携帯電話の担当者（customer support agent）に連絡し，本人（企業の従業員）になりすまして，電話番号を自分の SIM に移行させてしまうらしい。
これが SIM スワップ攻撃の大まかな手口である。

もちろん，これを成功させるには，なりすますための個人情報を詳細に把握している必要がある。

{{< fig-quote type="markdown" title="How to protect your business from SIM swapping" link="https://proton.me/business/blog/sim-swapping" lang="en" >}}
A SIM swap relies more on preparation than technical skill. First, the attacker collects enough personal information to pass the carrier’s identity checks, such as their full name, date of birth, and answers to common security questions.

With this profile, the attacker contacts the carrier and poses as the customer. They give a reason for the request, such as a lost phone, a damaged SIM, or an upgrade. The goal is always to get the support agent to activate the number on a new SIM.
{{< /fig-quote >}}

しかし SIM スワップが成功すれば，SMS 認証だけでなく，電話番号を認証やアカウント復旧に使っているサービスのアカウントを乗っ取られる可能性がある。

{{< fig-quote type="markdown" title="How to protect your business from SIM swapping" link="https://proton.me/business/blog/sim-swapping" lang="en" >}}
Once the swap is complete, the attacker moves fast. Password resets for many accounts, such as [business](https://proton.me/business/mail) email, banking, and software-as-a-service (SaaS) tools, often rely on SMS. The attacker triggers resets on the most valuable accounts first, intercepts each code, and locks the real owner out. This often happens before the employee realizes why their phone has no signal.
{{< /fig-quote >}}

この手口が厄介なのは，なりすまされる当人だけでは防ぎにくい点である。

{{< fig-quote type="markdown" title="How to protect your business from SIM swapping" link="https://proton.me/business/blog/sim-swapping" lang="en" >}}
The employee’s device isn’t hacked. The attacker doesn’t need common methods such as [malware](https://protonvpn.com/blog/what-is-malware), [brute force attacks](https://proton.me/blog/what-is-brute-force-attack), or [phishing](https://proton.me/blog/what-is-phishing). The carrier simply reassigns the number, using its normal process, to someone who has no right to it.

The weak point is the carrier’s identity checks; something your business has no control over.
{{< /fig-quote >}}

SIM スワップは元々，著名人や富裕層（暗号資産の保有者など）を狙う犯罪だったのだが，最近では企業の従業員アカウントがターゲットになっているらしい。
給与計算・メール・クラウドなどの管理権限を持つ従業員アカウントのどれかを攻略できれば，企業の重要な情報や資産にアクセスできる可能性がある。

{{< fig-quote type="markdown" title="How to protect your business from SIM swapping" link="https://proton.me/business/blog/sim-swapping" lang="en" >}}
The numbers show how fast this has grown. Cifas, the UK’s fraud prevention service, recorded a [1,055% rise](https://www.cifas.org.uk/newsroom/huge-surge-see-sim-swaps-hit-telco-and-mobile) in unauthorized SIM swaps in 2024, from 289 reported cases to nearly 3,000. Its [Fraudscape 2026 report](https://www.cifas.org.uk/newsroom/fraudscape2026) shows a further 38% increase in 2025, driven by easy access to stolen personal data and more automated attacks.

Mobile phone accounts now make up most account takeover cases, according to the [Cifas Fraudscape 2026 report](https://www.cifas.org.uk/newsroom/fraudscape2026). A phone number has become the recovery key for almost everything else.
{{< /fig-quote >}}

記事は英国の数字を挙げているが，日本はどうなんだろうね。
ショップとかで厳格化された本人確認がどのくらい効いているか知りたいものである。

念のため [Kagi Assistant] に「eSIM なら SIM スワップのリスクは減るのか」と訊いてみたのだが，SIM スワップ攻撃は基本的にソーシャルエンジニアリングなので eSIM だけで SIM スワップを防げるとは限らない，と言われた。
そりゃ，まぁ，そーだ（笑）

## SMS 認証に代えて

記事では SMS など電話番号に紐づく認証手段から別の手段へ移行することを推奨している。
具体的には

- TOTP 認証アプリ
- ハードウェアセキュリティキーを使った認証
- パスキー（FIDO2）認証

を挙げている。
運用にもよるが，物理キーや（公開鍵認証を行う）パスキー認証は有力な選択肢だろう。
パスワードと TOTP 認証を組み合わせる二要素認証も，SIM スワップによるリスクを低減できる。

最近は TOTP や FIDO2 に対応したパスワード管理ツールも増えているので，スマホやブラウザへの依存を減らしつつ，安全な認証手段を利用できるだろう。
以下は，私が使っているパスワード管理アプリでパスキー認証を行う際に参考となる記事だ。

- [実装に1年、「KeePassXC」が待望の「パスキー」対応 ～フリーのパスワード管理アプリ／v2.7.7がリリース](https://forest.watch.impress.co.jp/docs/news/1575337.html) (2024年の記事)
  - [KeePassXCを使ってLinuxでもパスキーを使えるようにした話 - Qiita](https://qiita.com/hiratannnnnn/items/f07e251d714f1bd2181a) : ブラウザ拡張を強引に入れて動かした話
- [Android デバイス上で KeePassDX を使ってパスキーを管理する](https://www.naney.org/diki/d/2026-07-25-passkey.html)

あとは利用してるサービスの認証方法の洗い出しかね。
企業とかで管理者アカウントを持っている場合は優先順位を決めて対応していくのがいいらしい。

{{< fig-quote type="markdown" title="How to protect your business from SIM swapping" link="https://proton.me/business/blog/sim-swapping" lang="en" >}}
- **Email administration:** Email is the recovery route for almost everything else. An attacker with your email admin account can reset their way into the rest of your systems.
- **Finance and payroll platforms:** These accounts can move money.
- **Domain registrars and cloud consoles:** These control your websites and infrastructure.
- **Everything else:** Lower-risk tools can move on a slower schedule.
{{< /fig-quote >}}

## SIM スワップへの対応は最初の1時間が肝心らしい

SIM スワップの兆候は些細なことから始まる。
電波が届くはずの場所なのに突然圏外になったり（しかも再起動しても直らない）。
同じ頃に覚えのないログイン通知やパスワードリセット通知などが来たら，SIM スワップを疑うべきだろう。

SIM スワップかなと思ったら，最初の1時間が肝心らしい。

{{< fig-quote type="markdown" title="How to protect your business from SIM swapping" link="https://proton.me/business/blog/sim-swapping" lang="en" >}}
The first hour is critical. The attacker is working through account recovery one account at a time, and every minute gives them more access. Act fast:

1. Contact the carrier from another phone or through its online support. Report the unauthorized swap and ask them to freeze and recover the number.
2. Secure the accounts the number protects, starting with email. Change passwords from a trusted device and sign out of all active sessions.
3. In the UK, report the incident to [Action Fraud](https://www.actionfraud.police.uk/), the national reporting center for fraud and cybercrime. Reports support investigations and push carriers and banks to tighten their checks.

Your business can also reduce the risk on the carrier side. Most carriers offer extra protections, such as a SIM lock, port freeze, or account PIN, that add a check before any transfer.
{{< /fig-quote >}}

3番目は英国の話だが，日本だと[警察庁の通報窓口](https://www.npa.go.jp/bureau/cyber/soudan.html "サイバー事案に関する相談窓口｜警察庁Webサイト")があるらしい。
ただし，人命に関わる事案など緊急を要する場合は110番でもOKだそうな。

{{< fig-quote type="markdown" title="How to protect your business from SIM swapping" link="https://proton.me/business/blog/sim-swapping" lang="en" >}}
The lasting fix is to make sure a phone number can’t unlock your business accounts.
{{< /fig-quote >}}

## 【おまけ】電子メールを Out-of-Band 認証に使ってはいけない

SMS 認証はしないほうがいいという話は世間的にもかなり認知されたと思うが，じゃあってんで電子メールを認証に使うサービスがポツポツあるのが気になっている。
[NIST SP 800-63] では，電子メールを Out-of-Band 認証として使用するのは NG である。

{{< fig-quote type="markdown" title="3.1.3.1. Out-of-Band Authenticators" link="https://pages.nist.gov/800-63-4/sp800-63b/authenticators/#ooba" lang="en" >}}
Email **SHALL NOT** be used for out-of-band authentication because it may be vulnerable to:

- Access using only a password
- Interception in transit or at intermediate mail servers
- Rerouting attacks, such as those caused by Domain Name System (DNS) spoofing

> Confirmation codes that are sent to validate email addresses or are issued as recovery codes (see [Sec. 4.2.1.2](https://pages.nist.gov/800-63-4/sp800-63b/events/#issuedrecovery)) are not authentication processes and not affected by the above prohibition.
{{< /fig-quote >}}

例外として，上の引用にあるように，メールアドレスの確認やリカバリコード送信のために電子メールを使うのは問題ない。
ただし，期限を切って運用する必要がある（電子メールの場合は最大でも24時間）。

というわけで，電子メールを Out-of-Band 認証には使ってはいけない。

## ブックマーク

- [NIST SP 800-63-4 最終版がリリース]({{< ref "/remark/2025/08/nist-sp-800-63-4.md" >}})

[How to protect your business from SIM swapping]: https://proton.me/business/blog/sim-swapping "How to protect your business from SIM swapping | Proton"
[NIST SP 800-63]: https://pages.nist.gov/800-63-4/ "NIST SP 800-63 Digital Identity Guidelines"
[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"

## 参考図書

{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
{{% review-paapi "4757143044" %}} <!-- 信頼と裏切りの社会 -->
