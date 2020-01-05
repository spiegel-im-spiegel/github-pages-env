+++
title = "Firefox の Cipher Suite"
date = "2018-12-16T11:24:32+09:00"
update = "2018-12-16T17:55:21+09:00"
description = "たまには Firefox のことも思い出してあげてください"
image = "/images/attention/kitten.jpg"
tags = [
  "security",
  "cryptography",
  "firefox",
  "tls",
  "management",
]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

- [J-STAGEがFirefoxでのアクセスを遮断、日本の電子ジャーナルが世界から不可視となった日｜Guest｜note](https://note.mu/note_s/n/n517ff243e083)

この記事を読んで以下のように思わなかっただろうか。

「Firefox 側が `TLS_RSA_WITH_AES_256_CBC_SHA256` あたりを用意すればええんちゃうん？」

私は思った。
[J-STAGE](https://www.jstage.jst.go.jp/) というサイトや運営している [JST](https://www.jst.go.jp/ "国立研究開発法人 科学技術振興機構") という組織をよく知らないので，なんでそこまで悪意たっぷりに書かれるのかよく分からないが，もうちょっと書きようがあるだろうに。

## TLS Cipher Suite

まず TLS (Transport Layer Security) Version 1.2 の仕様を記した [RFC 5246] によると “[Mandatory Cipher Suites](https://tools.ietf.org/html/rfc5246#section-9)” というのがあって，この中で `TLS_RSA_WITH_AES_128_CBC_SHA` を MUST に定めている。
したがって，上述のリンク先の記事でこれがないことについて「TLS仕様違反」と断じている点は間違いない。 

この場合，サーバ側の対応としては， NSS なら `TLS_RSA_WITH_AES_128_CBC_SHA`， OpenSSL なら `AES128-SHA` を有効にすればいいわけだ。

一方， Firefox はどうなっているかというと `about:config` を開いて “`ssl3`” で検索すると出てくる。

{{< fig-img src="./cipher-suite-in-firefox.png" title="Cipher Suite in Firefox" link="./cipher-suite-in-firefox.png" width="643" >}}

ちなみにこれは Firefox Developer Edition 65 の場合である。
一応，抜き書きすると

- `security.ssl3.dhe_rsa_aes_128_sha`
- `security.ssl3.dhe_rsa_aes_256_sha`
- `security.ssl3.ecdhe_ecdsa_aes_128_gcm_sha256`
- `security.ssl3.ecdhe_ecdsa_aes_128_sha`
- `security.ssl3.ecdhe_ecdsa_aes_256_gcm_sha384`
- `security.ssl3.ecdhe_ecdsa_aes_256_sha`
- `security.ssl3.ecdhe_ecdsa_chacha20_poly1305_sha256`
- `security.ssl3.ecdhe_rsa_aes_128_gcm_sha256`
- `security.ssl3.ecdhe_rsa_aes_128_sha`
- `security.ssl3.ecdhe_rsa_aes_256_gcm_sha384`
- `security.ssl3.ecdhe_rsa_aes_256_sha`
- `security.ssl3.ecdhe_rsa_chacha20_poly1305_sha256`
- `security.ssl3.rsa_aes_128_sha`
- `security.ssl3.rsa_aes_256_sha`
- `security.ssl3.rsa_des_ede3_sha`

となっている。
これを見ると `TLS_RSA_...` から始まる古い cipher suite は殆ど対応していないことが分かる（mandatory cipher suites を除けば2つのみ）。
Firefox 側の思惑は不明だが，古い cipher suite は PFS (Perfect Forward Secrecy) に対応してない（できない）ため意図的に外してる可能性もある。

## Cipher Suites for Modern Browser

（おそらくもうサポートされていない）古いブラウザを切り捨てていいのであれば，以下の cipher suite をサポートしておけば問題ないらしい。

| HEX       | IANA                                            | OpenSSL                         |
| --------- | ----------------------------------------------- | ------------------------------- |
| 0xC0,0x2C | `TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384`       | `ECDHE-ECDSA-AES256-GCM-SHA384` |
| 0xC0,0x30 | `TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384`         | `ECDHE-RSA-AES256-GCM-SHA384`   |
| 0xCC,0xA9 | `TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256` | `ECDHE-ECDSA-CHACHA20-POLY1305` |
| 0xCC,0xA8 | `TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256`   | `ECDHE-RSA-CHACHA20-POLY1305`   |
| 0xC0,0x2B | `TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256`       | `ECDHE-ECDSA-AES128-GCM-SHA256` |
| 0xC0,0x2F | `TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256`         | `ECDHE-RSA-AES128-GCM-SHA256`   |
| 0xC0,0x24 | `TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384`       | `ECDHE-ECDSA-AES256-SHA384`     |
| 0xC0,0x28 | `TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384`         | `ECDHE-RSA-AES256-SHA384`       |
| 0xC0,0x23 | `TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256`       | `ECDHE-ECDSA-AES128-SHA256`     |
| 0xC0,0x27 | `TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256`         | `ECDHE-RSA-AES128-SHA256`       |

これらの cipher suite は Firefox 27 以降, Chrome 30 以降, IE 11 以降, Edge, Opera 17 以降, Safari 9 以降, Android 5.0 以降, Java 8 以降であれば対応している。

## たまには Firefox のことも思い出してあげてください

[以前も書いた]({{< ref "/remark/2018/12/software-diversity.md" >}} "オープンソースはコードの多様性を担保しない")が，携帯端末も含めれば Chrome/Chromium と Safari/WebKit でほぼ寡占状態と言ってよい。
そういった状況の中で Firefox も対応しろと言われてもかったるいかもしれないし，実際に私が関わったことのある案件でも Chrome や Edge での動作確認はするのに Firefox はスルーするところも結構あったりした。
特定企業への依存度が減るんだから，むしろ Firefox は推奨されていいと思うんだけどねぇ。

少なくとも HTTPS 接続で Firefox を{{< ruby "ハブ" >}}村八分{{< /ruby >}}にする理由はないと思うので，是非ともよろしくお願いします。

## ブックマーク

- [How can I disable arbitrary SSL/TLS cipher suites in Firefox? | Firefox Support Forum | Mozilla Support](https://support.mozilla.org/en-US/questions/1119007)
- [Security/Server Side TLS - MozillaWiki](https://wiki.mozilla.org/Security/Server_Side_TLS)
- [RFC 7905 - ChaCha20-Poly1305 Cipher Suites for Transport Layer Security (TLS)](https://tools.ietf.org/html/rfc7905)

- [SSL Server Test (Powered by Qualys SSL Labs)](https://www.ssllabs.com/ssltest/)
- [SSL Cipher Suites Supported By Your Browser](https://cc.dcsec.uni-hannover.de/)
- [GitHub - mozilla/cipherscan: A very simple way to find out which SSL ciphersuites are supported by a target.](https://github.com/mozilla/cipherscan)
- [GitHub - square/certigo: A utility to examine and validate certificates in a variety of formats](https://github.com/square/certigo)

- [暗号鍵関連の各種変数について]({{< ref "/remark/2017/10/key-parameters.md" >}})

[RFC 5246]: https://tools.ietf.org/html/rfc5246 "RFC 5246 - The Transport Layer Security (TLS) Protocol Version 1.2"

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
