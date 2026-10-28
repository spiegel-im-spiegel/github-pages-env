+++
title = "eMudhra 中間 CA 鍵の危殆化と Go の対応"
date =  "2026-10-28T12:38:56+09:00"
description = "Web や gRPC など TLS 接続を行うネットワークサービスにアクセスする可能性のある Go 製ツールは，コンパイラを 1.27.3 (または 1.26.10) 以降に更新してビルドし直す必要がある。"
isCJKLanguage = true
image = "/images/attention/go-logo_blue.png"
tags = [ "security", "vulnerability", "risk", "tls", "pki", "trust", "certification", "golang", "x509" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

## eMudhra 中間 CA 鍵の危殆化

[eMudhra] は，インドを拠点にデジタル ID や電子証明書，PKI などの事業を行う企業グループだそうな[^e1]。
証明書発行事業を担う会社は eMudhra Limited で，インド認証局監督機関（Controller of Certifying Authorities; CCA）の[ライセンスを受けている](https://cca.gov.in/e_mudhra.html "Disclosure Record of e-Mudhra CA | CCA")ようだ。

[^e1]: [eMudhra] をカタカナにすると「イーミュドラ」表記になるようだ。

先日の中間 CA 証明書の秘密鍵の危殆化（compromise）と，それに伴うルート CA 証明書の失効に関する情報については，以下が参考になる。

- [eMudhra Statement on Recent Certificate Security Incident](https://emudhra.com/en/post/certificate-security-incident)
- [2083449 - eMudhra: Private key compromise of Intermediate Certificates](https://bugzilla.mozilla.org/show_bug.cgi?id=2083449)

どうやら，クライアント側で中間 CA 証明書を個別に失効させられない場合に備えた予防措置としてルート CA 証明書も失効することにしたらしい。
主要ブラウザでは既にルート CA 証明書が失効しているはずである。
ただし，独自に CA 証明書をバンドルしているアプリケーション，コンテナ，VPN，組込み機器などは，個別の対応が必要になる。

今回の危殆化に至る経緯が明らかになるにつれ（かなり間抜けだったらしい），同社に対する批判が高まり，グローバル SSL/TLS 証明書事業から段階的に[撤退](https://emudhra.com/en/press-release/emudhra-to-exit-global-ssl-tls-certificate-business "eMudhra to Exit global SSL/TLS Certificate Business")することにしたようだ。

{{< fig-quote type="markdown" title="eMudhra to Exit global SSL/TLS Certificate Business" link="https://emudhra.com/en/press-release/emudhra-to-exit-global-ssl-tls-certificate-business" lang="en" >}}
However, the level of social media amplification around this incident has brought scrutiny that is out of proportion to SSL/TLS’s place in our business. In view of this, after detailed deliberations our board has concluded that withdrawing from global SSL/TLS is the right decision for eMudhra and its shareholders. For many certifying authorities, SSL/TLS is the whole business. For eMudhra it is one part of a much broader digital solutions portfolio. We will put our energy where we create the most value for customers with Sharpened Focus on Core Enterprise Solutions and trust services other than SSL/TLS.
{{< /fig-quote >}}

他のインド国内向けの事業については継続するとのこと。

## Go の対応

今回の件に対応するため [Go] 1.27.3 がリリースされた。
経緯や対応については，以下の公式ブログ記事が参考になる。

- [Our Response To The eMudhra Key Disclosure Incident - The Go Programming Language](https://go.dev/blog/emudhra-incident)

[Go] は標準ライブラリに TLS 証明書の検証機能を持っている。
Windows と macOS では検証を OS の API に移譲しているが Linux では独自実装の検証器を用いている。

{{< fig-quote type="markdown" title="Our Response To The eMudhra Key Disclosure Incident" link="https://go.dev/blog/emudhra-incident" lang="en" >}}
Go ships a TLS certificate verifier in the standard library. When users use [`net/http`](https://pkg.go.dev/net/http) or [`crypto/tls`](https://pkg.go.dev/crypto/tls) to connect to a TLS service, the verifier is called to check if the certificates the other side of the connection presents are valid. On Windows and macOS we delegate these checks to the platform APIs, and on Linux, which lacks platform APIs, we use our own verifier implementation.
{{< /fig-quote >}}

というわけで Web や gRPC など TLS 接続を行うネットワークサービスにアクセスする可能性のある [Go] 製ツールは，[Go] コンパイラを 1.27.3 (または 1.26.10) 以降に更新してビルドし直す必要がある。

{{< fig-quote type="markdown" title="Our Response To The eMudhra Key Disclosure Incident" link="https://go.dev/blog/emudhra-incident" lang="en" >}}
If you run Go binaries which interact with network services using TLS (such as directly via net/http or crypto/tls or via a dependency such as gRPC), ensure you are running on the latest minor releases. Binaries rebuilt with Go 1.27.3 and 1.26.10 will automatically reject certificates issued by the compromised eMudhra intermediates.
{{< /fig-quote >}}

[Go] 以外の製品の情報については次節の[ブックマーク](#bookmarks)を参照のこと（随時更新予定）。
まぁ，証明書の危殆化はビックリするほど珍しいことではないので（多くはないが過去にも危殆化事例はある），粛々と対応するのが吉である。

## ブックマーク {#bookmarks}

- Google: [Protecting Chrome and Android Users from Compromised eMudhra Intermediate CA Certificates](https://blog.google/security/protecting-chrome-and-android-users-from-compromised-emudhra-intermediate-ca-certificates/)
- Apple: [About the compromised eMudhra intermediate CAs - Apple Support](https://support.apple.com/en-us/149555)
- Microsoft: [Microsoft’s response to compromised eMudhra certificate authority keys](https://www.microsoft.com/en-us/msrc/blog/2026/10/microsoft-response-to-compromised-emudhra-certificate-authority-keys)
- Cloudflare: [8 eMudhra certificates are compromised. Here's what you need to know](https://blog.cloudflare.com/emudhra-certificate-compromise/)
- Cisco: [Cisco Informational for Emergency Distrust of eMudhra Root and Intermediate Certificate Authorities](https://sec.cloudapps.cisco.com/security/center/content/CiscoSecurityAdvisory/cisco-sa-emudhra-info-ScS22A1r)

[eMudhra]: https://www.emudhra.com/ "eMudhra - Digital Trust Solutions"
[Go]: https://go.dev/

## 参考図書

{{< linkcard "9e8d8f717b1d1d85b0da7f07cb10a83b78d4227f" >}} <!-- プログラミング言語Go -->
{{< review-paapi "B0CFL1DK8Q" >}} <!-- Go言語 100Tips -->
{{< review-paapi "B0DNYMMBBQ" >}} <!-- Go言語で学ぶ並行プログラミング -->
