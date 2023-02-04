+++
title = "DSA は NIST 電子署名標準から外れるようだ"
date =  "2021-02-17T20:53:55+09:00"
description = "ドラフト版 FIPS 186-5 の DSA の節の内容がまるっと削られている。"
image = "/images/attention/kitten.jpg"
tags = [ "security", "cryptography", "risk", "ecc", "nist", "dsa" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go 1.16 のリリースノート](https://golang.org/doc/go1.16 "Go 1.16 Release Notes - The Go Programming Language")を眺めていて今さら気付いたのだが，いまだ[ドラフト中の FIPS 186-5][FIPS 186-5] によると， DSA を「電子署名生成用途としては」電子署名標準（Digital Signature Standard; DSS）から外すつもりのようだ。
ドラフト版 [FIPS 186-5] の DSA の節はまるっと削られて以下の文章のみ掲載されている。

{{< fig-quote type="markdown" title="FIPS 186-5 (Draft), Digital Signature Standard (DSS)" link="https://csrc.nist.gov/publications/detail/fips/186/5/draft" lang="en" >}}
{{% quote %}}Prior versions of this standard specified the DSA. This standard no longer approves DSA for digital signature generation. DSA may be used to verify signatures generated prior to the implementation date of this standard. See FIPS 186-4 [20] for the specifications for DSA{{% /quote %}}.
{{< /fig-quote >}}

鍵長の問題（現在は鍵長を変数化することで対処している）以外はアルゴリズム自体に危殆化問題があるわけではないと思うが，出自がアレだし実装面の複雑さ[^dsa1] もあって評判がよろしくないのは聞いていた。
まぁ「設計限界」というやつなんだろう。

[^dsa1]: DSA は実装に必要な要素技術が多く，そのうちのひとつでも瑕疵があれば全体のセキュリティ強度が下がってしまう。

その代わりといってはナニだが楕円曲線暗号のひとつである EdDSA (楕円曲線 ed25519/ed488) が [FIPS 186-5] から電子署名標準のひとつとして加わることになっている。
これについては以前に記事を書いた（なんでそのときに DSA に気がつかなかったのか）。

- [Edwards-curve Digital Signature Algorithm]({{< ref "/remark/2020/06/eddsa.md" >}})

まぁ，時代は楕円曲線暗号ということなのだろう。
次期 [OpenPGP] である [RFC 4880bis] が正式にリリースされるタイミングでメインの鍵を替えようかと目論んでいるのだが，前倒ししたほうがいいのかねぇ。

## ブックマーク

- [crypto/dsa: deprecate and remove from crypto/x509 and x/crypto/ssh · Issue #40337 · golang/go · GitHub](https://github.com/golang/go/issues/40337)
- [Cryptography Dispatches: DSA Is Past Its Prime](https://buttondown.email/cryptography-dispatches/archive/cryptography-dispatches-dsa-is-past-its-prime/)
- [そろそろ GnuPG でも ECC を標準で使うのがいいんじゃないかな]({{< ref "/openpgp/using-ecc-with-gnupg.md" >}})
- [Go 1.16 がリリースされた]({{< ref "/release/2021/02/go-1_16-is-released.md" >}})： Go 1.16 では [`crypto/dsa`] パッケージが非推奨になった。
- [NIST FIPS 186-5 および SP 800-186 正式版がリリースされた]({{< ref "/remark/2023/02/nist-fips-186-5.md" >}})

[FIPS 186-5]: https://csrc.nist.gov/publications/detail/fips/186/5/draft "FIPS 186-5 (Draft), Digital Signature Standard (DSS) | CSRC"
[Go]: https://go.dev/
[`crypto/dsa`]: https://golang.org/pkg/crypto/dsa/ "dsa - The Go Programming Language"
[OpenPGP]: http://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
