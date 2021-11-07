+++
title = "それはワンタイム・パスワードの問題ではない"
date =  "2021-11-07T18:28:00+09:00"
description = "道具は適切に組み合わせないと所定の性能を発揮できない。 "
image = "/images/attention/kitten.jpg"
tags = [ "security", "authentication", "risk", "management", "nist" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

いやさ

- [狙われるワンタイムパスワード、多要素認証を破る闇サービスが浮上：この頃、セキュリティ界隈で（1/2 ページ） - ITmedia NEWS](https://www.itmedia.co.jp/news/articles/2111/05/news052.html)

て酷い釣りタイトルだよな，と思いつつモニタにツッコんでしまったよ（笑）

ハッキリ言おう。
それはワンタイム・パスワードの問題ではない。

ワンタイム・パスワードに問題がないわけではない。
[NIST SP 800-63B](https://pages.nist.gov/800-63-3/sp800-63b.html "NIST Special Publication 800-63B") によればスマホの TOTP アプリや専用のワンタイム・パスワード機器は「単要素 OTP デバイス（Single-Factor One-Time Password (OTP) Device）」と呼ばれる。
これはタイプとしては「知識」ではなく「所有」に分類される。
所有型の Authenticator には紛失・盗難のリスクが伴う。
また，基本的に OTP は認証する側とされる側との間で最初にシークレットを共有する必要があり（特にアプリでは）シークレットの受け渡しと管理の問題が発生するが，認証の段階でこのシークレットをやり取りすることはない。

しかし，上の記事の「ワンタイム・パスワード」はこれとは異なる。
SMS やチャット・アプリ等を通してサービスプロバイダから使い捨てシークレットを通知し，シークレットをもらったユーザは通知を受けたチャネルとは別のチャネル（大抵はスマホ・アプリかブラウザで表示される Web ページ）でシークレットを返す。
送ったシークレットと返ってきたシークレットを比較して認証を行うわけだ。
こうした仕組みは [NIST SP 800-63B] の分類では「経路外デバイス（Out-of-Band Devices）」と呼ばれている。

経路外デバイスが筋が悪いのは攻撃者から見て攻撃ポイントが多いことだ。
上の記事で示される事例はまさにそこを突かれて認証を突破されている。

以前に拙文「[Authenticator と AAL]({{< ref "/remark/2020/09/authenticator-and-aal.md" >}})」で紹介したが，元々 NIST は SMS を使った認証を非推奨（または禁止）にするつもりだった。

- [SMSを使った二要素認証を非推奨〜禁止へ、米国立技術規格研究所NISTの新ガイダンス案 | TechCrunch Japan](http://jp.techcrunch.com/2016/07/26/20160725nist-declares-the-age-of-sms-based-2-factor-authentication-over/)

しかしその後，色々あったようで，最終的には “[RESTRICTED Authenticator](https://pages.nist.gov/800-63-3/sp800-63b.html#restricted)” という位置づけまで緩和されてしまった。

{{< fig-quote title="NIST Special Publication 800-63B" link="https://pages.nist.gov/800-63-3/sp800-63b.html#restricted" lang="en" >}}
{{% quote %}}The use of a RESTRICTED authenticator requires that the implementing organization assess, understand, and accept the risks associated with that RESTRICTED authenticator and acknowledge that risk will likely increase over time. It is the responsibility of the organization to determine the level of acceptable risk for their system(s) and associated data and to define any methods for mitigating excessive risks. If at any time the organization determines that the risk to any party is unacceptable, then that authenticator SHALL NOT be used{{% /quote %}}.
{{< /fig-quote >}}

でも，結局その「温情措置」が重大なセキュリティ・インシデントを招いているのだから「なんだかなぁ」という感じである。
電子メールや VoIP が経路外デバイスとして NG なら SMS だって NG だろう。

Authenticator による認証には3つのレベル（Authenticator Assurance Level; AAL）があるが，最初の記事の事例にあるような「[暗号資産]({{< ref "/remark/2018/12/crypto-assets.md" >}} "暗号資産」とやら")」を扱う重要な決済システムなら AAL3 は必須だろうし，そうであるなら認証手段として経路外デバイスを選択するのはあり得ない。
これは「闇」でも何でもなく，単にサービス・プロバイダが間抜けで迂闊だったというだけの話である。

道具は適切に組み合わせないと所定の性能を発揮できない。
私達エンジニアはこのことを肝に銘じておくべきだし，利用者としても「利便性はセキュリティとトレードオフできない」ことは知っておくべきだろう。

[SP 800-63-3]: https://pages.nist.gov/800-63-3/ "NIST SP 800-63 Digital Identity Guidelines"
[NIST SP 800-63B]: https://pages.nist.gov/800-63-3/sp800-63b.html "NIST Special Publication 800-63B"

## 参考図書

{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
{{% review-paapi "4757143044" %}} <!-- 信頼と裏切りの社会 -->
