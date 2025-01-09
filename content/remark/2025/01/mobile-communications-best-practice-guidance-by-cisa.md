+++
title = "CISA による携帯通信のベストプラクティス"
date =  "2025-01-09T21:04:30+09:00"
description = "この記事では覚え書きとして記しておく。"
image = "/images/attention/kitten.jpg"
tags = [ "security", "risk", "management", "k-tai" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

米国 [CISA] が携帯通信のベストプラクティスを公開している。

- [Mobile Communications Best Practice Guidance | CISA](https://www.cisa.gov/resources-tools/resources/mobile-communications-best-practice-guidance)
- {{< pdf-file title="Mobile Communications Best Practice Guidance" link="https://www.cisa.gov/sites/default/files/2024-12/guidance-mobile-communications-best-practices.pdf" >}}

この記事では覚え書きとして記しておく。

## 一般的な推奨事項

“General Recommendations” としては以下の8つ。

1. **Use only end-to-end encrypted communications.**<br>例として [Signal](https://signal.org/ "Signal") を挙げている
2. **Enable Fast Identity Online (FIDO)** [phishing-resistant authentication](https://www.cisa.gov/sites/default/files/publications/fact-sheet-implementing-phishing-resistant-mfa-508c.pdf).<br> [Yubico](https://www.yubico.com/products/) や [Google Titan](https://cloud.google.com/security/products/titan-security-key) のようなハードウェアベースのキーを推奨しているが代替手段として FIDO パスキーも許容されるとある
3. **Migrate away from Short Message Service (SMS)-based MFA.**<br>MFA (Multi-Factor Authentication; 多要素認証) の2番目の要素として SMS は使わないという話。ただしアカウント回復フローとして SMS を使う場合があるため完全な排除は現実的ではないと述べている
4. **Use a password manager** to store all passwords.<br>例として以下のアプリを挙げている
   - [Apple Passwords app](https://support.apple.com/en-us/120758)
   - [LastPass](https://www.lastpass.com/)
   - [1Password](https://1password.com/)
   - [Google Password Manager](https://passwords.google.com/)
   - [Dashlane](https://www.dashlane.com/)
   - [Keeper](https://www.keepersecurity.com/)
   - [Proton Pass](https://proton.me/pass)
5. **Set a Telco PIN.**<br>スマホの画面を開くときの PIN コードではなくキャリアとの間で決めている PIN コード。いわゆる SIM スワッピングに対抗するための重要なステップだと述べている
6. **Regularly update software.**<br>OS とアプリの両方のアップデート。自動更新を有効にしろとある
7. **Opt for the latest hardware version from your cell phone manufacturer.**<br>古いハードは最新のセキュリティ機能が組み込まれていない場合があるため
8. **Do not use a personal virtual private network (VPN).**<br>無料および商用 VPN プロバイダーのセキュリティおよびプライバシーポリシーは怪しいとして，プライベート VPN の利用はセキュリティリスクを ISP から移転しているだけに過ぎないと述べている

## iPhone 固有の推奨事項

私は iPhone も iPad も持ってないので “iPhone-Specific Recommendations” についてはよく分からない。
ので，項目を列挙するだけに留める。

1. **Enable [Lockdown Mode](https://support.apple.com/en-us/105120).**
2. **Disable the following setting** to ensure messages do not send as SMS if iMessage is unavailable.
   - Disable: Settings → Apps → Messages → “Send as Text Message”
3. **Protect your Domain Name System (DNS) queries.**
4. **Enroll in [Apple iCloud Private Relay](https://support.apple.com/en-us/102602).**
5. **Review and [restrict app permissions](https://support.apple.com/guide/iphone/control-access-to-information-in-apps-iph251e92810/18.0/ios/18.0)** through Settings → Privacy & Security.

## 携帯端末の維持は難しい

個々の項目についての感想は控えるが，スマホを含む携帯端末は紛失・盗難リスクが高い上にハードウェアや OS のバージョンアップ・サイクルが短く（自宅機の買い替えが8年単位の私から見て）頻繁に買い替えが必要なのが致命的だと思っている。
これのせいでスマホを認証デバイスとして使うことに躊躇してる。
如何にしてスマホに依存しないオンライン活動をするか，だよなぁ...

## ブックマーク

- [Authenticator と AAL]({{< ref "/remark/2020/09/authenticator-and-aal.md" >}})
- [誰が「パスワードは複雑なものにしろ」と言ったのか]({{< ref "/remark/2024/12/deprecated-password-policy.md" >}})

[CISA]: https://www.cisa.gov/ "CISA"
