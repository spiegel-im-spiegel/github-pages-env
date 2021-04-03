+++
title = "gpgpdump v0.12.1 をリリースした"
date =  "2021-04-03T11:41:35+09:00"
description = "golang.org/x/crypto が OpenPGP のサポートを止めたので，サードパーティのものに置き換えた。"
image = "/images/attention/openpgp.png"
tags = ["tools", "openpgp", "gpgpdump", "golang"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP] パケットの内容を可視化する [gpgpdump] の v0.12.2 をリリースした。

- [Release v0.12.2 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.12.2)

[gpgpdump] でも利用している [`golang.org/x/crypto/openpgp`] パッケージが凍結されてしまったようだ。

- [x/crypto/openpgp: mark as frozen and deprecated · Issue #44226 · golang/go](https://github.com/golang/go/issues/44226)

まぁ，色々と御託が並べられているが，要は「[OpenPGP] は触りたくない」と言うことだろう。
触りたくない使いもしないパッケージを（事実上）捨てるのは間違ってない。

幸いなことに [`golang.org/x/crypto`] 互換で活況そうな [`ProtonMail/go-crypto`] パッケージがそのまま使えるようなので（テストも問題なく通った），今回はこれに置き換えることで応急措置とした。
機能上の修正・変更はない。

といっても [gpgpdump] で [`golang.org/x/crypto/openpgp`] パッケージを使ってたのはパケットの切り出しに便利だったからというだけで，暗号周りの機能は一切使ってない。
最悪は自分で組んで置き換えるかなぁ...

まぁ，[個人のプライバシーに敵対的な企業]({{< ref "/remark/2018/04/handling-privacy.md" >}} "誰がプライバシーを支配するのか")がホストする言語がいくら [OpenPGP] を DIS っても微塵も刺さらないところがにんともかんとも（笑） 「坊主憎けりゃ袈裟まで憎い」などと言うつもりはないので，利用できる部分は利用して賢くやりくりしましょう。

それでにしても先週の mimemagic の騒ぎといい，特定のフレームワークやライブラリに依存しすぎるのは考えものだねぇ。
というわけで最近『[Clean Architecture](https://www.amazon.co.jp/dp/B07FSBHS2V?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』を読んでいる。

## ブックマーク

- [世界一わかりやすいClean Architecture - nuits.jp blog](https://www.nuits.jp/entry/easiest-clean-architecture-2019-09)

- [OpenPGP の実装]({{< rlnk "openpgp/" >}})
- [OpenPGP パケットを可視化する gpgpdump]({{< ref "/release/gpgpdump.md" >}})

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[OpenPGP]: http://openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[RFC 4880]: https://tools.ietf.org/html/rfc4880
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/
[Go]: https://golang.org/ "The Go Programming Language"
[`golang.org/x/crypto`]: https://pkg.go.dev/golang.org/x/crypto "crypto · pkg.go.dev"
[`golang.org/x/crypto/openpgp`]: https://pkg.go.dev/golang.org/x/crypto/openpgp "openpgp · pkg.go.dev"
[`ProtonMail/go-crypto`]: https://github.com/ProtonMail/go-crypto "GitHub - ProtonMail/go-crypto: [mirror] Go supplementary cryptography libraries"

## 参考図書

{{% review-paapi "B07FSBHS2V" %}} <!-- Clean Architecture -->
{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
