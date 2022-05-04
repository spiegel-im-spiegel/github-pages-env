+++
title = "Authenticator と AAL"
date =  "2020-09-27T19:45:38+09:00"
description = "どうも日本の金融界は「リスク感度が鈍い」そうなので，自衛のためにも2017年にリリースされた NIST SP 800-63-3 をベースに少しお勉強しておく。"
image = "/images/attention/kitten.jpg"
tags = [ "security", "authentication", "risk", "management", "nist" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

どうも日本の金融界は「[リスク感度が鈍い](https://japan.cnet.com/article/35160001/ "ゆうちょ池田社長「リスク感度が鈍かった」--被害は約6000万円に拡大、2017年から発生 - CNET Japan")」そうなので，自衛のためにも2017年にリリースされた [NIST SP 800-63-3][SP 800-63-3] をベースに少しお勉強しておく。

- [NIST Special Publication 800-63-3: Digital Identity Guidelines](https://pages.nist.gov/800-63-3/sp800-63-3.html)
- [NIST Special Publication 800-63A: Enrollment and Identity Proofing](https://pages.nist.gov/800-63-3/sp800-63a.html)
- [NIST Special Publication 800-63B: Authentication and Lifecycle Management](https://pages.nist.gov/800-63-3/sp800-63b.html)
- [NIST Special Publication 800-63C: Federation and Assertions](https://pages.nist.gov/800-63-3/sp800-63c.html)

[SP 800-63-3] といえばパスワード運用で当時は話題になった。

- [「パスワードのベストプラクティス」が変わる]({{< ref "/remark/2017/10/changes-in-password-best-practices.md" >}})

このパスワード話が出てくるのが [SP 800-63B] だが，このドキュメントでは Authenticator 全体について色々と書かれている。

## Authenticator

Authenticator について適切な日本語が見当たらないが，強いて言うなら「認証機能」あるいは「認証器」といったところだろうか。
たとえばパスワードも Authenticator だし，スマホにインストールした TOTP アプリも Authenticator だ。
Yubikey なんかの暗号デバイスも Authenticator に含まれる。

[SP 800-63B] では Authenticator を以下の9つに分類している。

| 種別名                                                         |     認証要素     |
| -------------------------------------------------------------- |:----------------:|
| Memorized Secrets<br>記憶シークレット                          |       知識       |
| Look-Up Secrets<br>ルックアップ・シークレット                  |       所有       |
| Out-of-Band Devices<br>経路外デバイス                          |       所有       |
| Single-Factor OTP Device<br>単要素 OTP デバイス                |       所有       |
| Multi-Factor OTP Devices<br>多要素 OTP デバイス                | 所有＋知識／生体 |
| Single-Factor Cryptographic Software<br>単要素暗号ソフトウェア |       所有       |
| Single-Factor Cryptographic Devices<br>単要素暗号デバイス      |       所有       |
| Multi-Factor Cryptographic Software<br>多要素暗号ソフトウェア  | 所有＋知識／生体 |
| Multi-Factor Cryptographic Devices<br>多要素暗号デバイス       | 所有＋知識／生体 |

また，各 Authenticator の例としては以下のものが挙げられる。

| Authenticator              | 具体例                                                                      |
| -------------------------- | ----------------------------------------------------------------------- |
| 記憶シークレット           | パスワード，PINコード                                                   |
| ルックアップ・シークレット | 乱数表，認証失敗時のリカバリコード                                                  |
| 経路外デバイス             | SMS によるコード送信， QR コード（電子メールや VoIP は認められない）    |
| 単要素 OTP デバイス        | アクティベーションを必要としない OTP デバイスまたはソフトウェア         |
| 多要素 OTP デバイス        | アクティベーションを行った上で利用可能な OTP デバイスまたはソフトウェア |
| 単要素暗号ソフトウェア     | セキュアなストレージ上で保護されている暗号鍵                            |
| 単要素暗号デバイス         | FIDO U2F の USB ドングル                                                |
| 多要素暗号ソフトウェア     | 単要素暗号ソフトウェアに対して追加のアクティベーションを必要とするもの  |
| 多要素暗号デバイス         | 単要素暗号デバイスに対して追加のアクティベーションを必要とするもの      |

## Authenticator Assurance Level

さらに [SP 800-63B] では AAL (Authenticator Assurance Level) を定義している。
AAL は 1 〜 3 の3段階あり，それぞれ以下に示す  Authenticator の組み合わせを許容している。

- AAL 1 では9種の Authenticator 全て許容され，単要素の認証で OK
- AAL 2 では以下に示す通り複数の認証要素による多要素認証が必要：
    - 多要素 OTP デバイス
    - 多要素暗号ソフトウェア
    - 多要素暗号デバイス
    - 記憶シークレット＋以下
        - ルックアップ・シークレット
        - 経路外デバイス
        - 単要素 OTP デバイス
        - 単要素暗号ソフトウェア
        - 単要素暗号デバイス
- AAL 3 では以下に示す通り，暗号鍵の所持証明要素とハードウェア関与を含む複数の認証要素による多要素認証が必要：
    - 多要素暗号デバイス
    - 単要素暗号デバイス＋記憶シークレット
    - 多要素OTPデバイス(SW/HW)＋単要素暗号デバイス
    - 多要素OTPデバイス(HW)＋単要素暗号ソフトウェア
    - 単要素OTPデバイス(HW)＋多要素暗号ソフトウェア
    - 単要素OTPデバイス(HW)＋単暗号ソフトウェア＋記憶シークレット

AAL の各レベルごとに要求されるセキュリティ事項（一部）は以下の通り。

| 要求事項                  | AAL 1 | AAL 2 | AAL 3 |
| ------------------------- |:-----:|:-----:|:-----:|
| 中間者攻撃耐性            | 必須  | 必須  | 必須  |
| Verifier なりすまし耐性   | 不要  | 不要  | 必須  |
| Verifier 改ざん耐性       | 不要  | 不要  | 必須  |
| リプレイ耐性              | 不要  | 必須  | 必須  |
| 認証意図（AuthN Inbtent） | 不要  | 推奨  | 必須  |
| レコード保持ポリシー      | 必須  | 必須  | 必須  |
| プライバシー統制          | 必須  | 必須  | 必須  |

金融系サービスの subscriber 確認で乗っ取りやなりすましを防ぎたいなら AAL 3 で何らかの物理暗号デバイスが必要だと思うけどねー。

## 格子型の乱数表は NG

現在は使ってるところはないだろうが，かつてネットバンキングでよく見られた格子型の乱数表はルックアップ・シークレットとしても NG だそうだ。
まぁ，当然だよな。

## SMS 認証は非推奨？

NIST は SMS によるコード送信について， [SP 800-63-3] のドラフト段階では非推奨にするつもりだったらしい。

- [SMSを使った二要素認証を非推奨〜禁止へ、米国立技術規格研究所NISTの新ガイダンス案 | TechCrunch Japan](https://techcrunch.com/2016/07/25/nist-declares-the-age-of-sms-based-2-factor-authentication-over/)

しかしその後，激しい議論があったようで，最終的には “[RESTRICTED Authenticator](https://pages.nist.gov/800-63-3/sp800-63b.html#restricted)” という位置づけまで緩和されたようだ。

{{< fig-quote type="markdown" title="NIST SP 800-63 Digital Identity Guidelines-FAQ" link="https://pages.nist.gov/800-63-FAQ/#q-b01" lang="en" >}}
{{% quote %}}Currently, authenticators leveraging the public switched telephone network, including phone- and Short Message Service (SMS)-based one-time passwords (OTPs) are restricted. Other authenticator types may be added as additional threats emerge. Note that, among other requirements, even when using phone- and SMS-based OTPs, the agency also has to verify that the OTP is being directed to a phone and not an IP address, such as with VoIP, as these accounts are not typically protected with multi-factor authentication{{% /quote %}}.
{{< /fig-quote >}}

（スマホを含む）電話機に依存した認証は，プライバシーも絡めて考えると筋が悪い。
ぶっちゃけ SMS 認証を含む経路外デバイスを使った認証は排除するか（ルックアップ・シークレットのように）優先順位を下げて非常時のみ使えるようにするのがいいと思う。
もちろん[電話番号を広告に流用](https://japan.cnet.com/article/35159898/ "Twitterに集団訴訟--電話番号がターゲティング広告に不正利用された可能性 - CNET Japan")するなど以っての外である。

## 生体情報は Authenticator として使えるか

Authenticator の分類を見れば分かるように，生体情報は単独では認証手段としては使えないという認識のようだ。
そもそも**生体情報は秘密情報ではない**のだから当たり前といえば当たり前かな。

## ブックマーク

- [usnistgov/800-63-3: Home to public development of NIST Special Publication 800-63-3: Digital Authentication Guidelines](https://github.com/usnistgov/800-63-3)
- [NIST SP800-63-3翻訳版63-Bパートの紹介](https://www.slideshare.net/kthrtty/20171027-nist-sp80063bkthrtty-81333156)
- {{< pdf-file title="認証にまつわるセキュリティの新常識" link="https://www.nic.ad.jp/sc-sendai/program/iwsc-sendai-d2-6.pdf" >}}
- [世界の電子認証基準が変わる：NIST SP800-63-3を読み解く – サポート − トラスト・ログイン byGMO【旧SKUID(スクイド)】](https://support.trustlogin.com/hc/ja/articles/115004031154-%E4%B8%96%E7%95%8C%E3%81%AE%E9%9B%BB%E5%AD%90%E8%AA%8D%E8%A8%BC%E5%9F%BA%E6%BA%96%E3%81%8C%E5%A4%89%E3%82%8F%E3%82%8B-NIST-SP800-63-3%E3%82%92%E8%AA%AD%E3%81%BF%E8%A7%A3%E3%81%8F)
- [Phishing Resistant SMS Autofill - The GitHub Blog](https://github.blog/2020-09-25-phishing-resistant-sms-autofill/)

[SP 800-63-3]: https://pages.nist.gov/800-63-3/ "NIST SP 800-63 Digital Identity Guidelines"
[SP 800-63B]: https://pages.nist.gov/800-63-3/sp800-63b.html "[NIST Special Publication 800-63B: Authentication and Lifecycle Management"

## 参考図書

{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
{{% review-paapi "4757143044" %}} <!-- 信頼と裏切りの社会 -->
