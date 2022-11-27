+++
title = "GopenPGP"
date = "2022-11-27T17:06:11+09:00"
description = "GopenPGP は OpenPGP の 暗号化・復号・署名およびその検証 といった処理を行うための Go パッケージ"
image = "/images/attention/openpgp.png"
tags = [ "openpgp", "golang", "programming", "package" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

本当は色々と試してから記事にしたかったのだが，しばらく時間が取れそうもないので，今回は簡単な紹介のみで。

- [GopenPGP | An easy-to-use OpenPGP wrapper library written in golang](https://gopenpgp.org/)

[GopenPGP] は [OpenPGP] の 暗号化・復号・署名およびその検証 といった処理を行うための [Go] パッケージで，大きく2つのレイヤに分かれている。

- OpenPGP 暗号化ライブラリ（[`golang.org/x/crypto/openpgp`] からの fork）
  - [`github.com/ProtonMail/go-crypto`]
- 暗号化等の処理を簡単に行うためのラッパー
  - [`github.com/ProtonMail/gopenpgp`]

一応，経緯を紹介すると，2021年に [Go] 側が [`golang.org/x/crypto/openpgp`] のメンテナンスを止めて非推奨（deprecated）にすると言いやがりまして。

- [x/crypto/openpgp: mark as frozen and deprecated · Issue #44226 · golang/go](https://github.com/golang/go/issues/44226)

まぁ，[個人のプライバシーに敵対的な企業]({{< ref "/remark/2018/04/handling-privacy.md" >}} "誰がプライバシーを支配するのか")がホストする言語ですから（笑）

その後 [Proton Mail] が [`golang.org/x/crypto/openpgp`] を fork したリポジトリを公開してメンテナンスを始めた。
これが [`github.com/ProtonMail/go-crypto`] パッケージ。
この頃のことは私も記事を書いている。

- [gpgpdump v0.12.2 をリリースした]({{< ref "/release/2021/04/gpgpdump-v0_12_2-is-released.md" >}})

当時は緊急避難的な措置かと思っていたが， [Proton Mail] 側は本格的に [Go] で [OpenPGP] 暗号化ライブラリを構築しようと考えたようだ。
これが [`github.com/ProtonMail/gopenpgp`] にあたると思われる。
また [`github.com/ProtonMail/go-crypto`] パッケージについても，現行の [RFC 4880] だけでなく，次期 [OpenPGP] ([RFC 4880bis]) でサポート予定の AEAD の EAX モードや OCB モードもサポートしているようで，これを見ても [Proton Mail] 側の本気度が伺えるというものである。

[OpenPGP] についてはリファレンス実装としての [GnuPG] が有名だが，他にも JavaScript や Rust 等による実装が存在する。
こうした中でようやく [Go] でもまともな実装が出てきたというわけだ。

今年のうちは無理かもしれないが来年（2023年）に入ったら，ちょっと真面目に [`golang.org/x/crypto/openpgp`] パッケージを弄ってみようかな，と考える所存であります。
ちゃんと時間が取れるといいなぁ。

## ブックマーク

- [【Golang】Go で OpenPGP のペア鍵を生成する【GopenPGP】 - Qiita](https://qiita.com/KEINOS/items/705f67cf47262f31dd19)
- [draft-ietf-openpgp-crypto-refresh-07](https://datatracker.ietf.org/doc/html/draft-ietf-openpgp-crypto-refresh-07) : 次期 [OpenPGP] ドラフト 2022-10-23 版

[OpenPGP]: http://openpgp.org/
[GopenPGP]: https://gopenpgp.org/ "GopenPGP | An easy-to-use OpenPGP wrapper library written in golang"
[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Go]: https://go.dev/
[`golang.org/x/crypto/openpgp`]: https://pkg.go.dev/golang.org/x/crypto/openpgp "openpgp package - golang.org/x/crypto/openpgp - Go Packages"
[`github.com/ProtonMail/go-crypto`]: https://github.com/ProtonMail/go-crypto "ProtonMail/go-crypto: Fork of go/x/crypto, providing an up-to-date OpenPGP implementation"
[`github.com/ProtonMail/gopenpgp`]: https://github.com/ProtonMail/gopenpgp "ProtonMail/gopenpgp: A high-level OpenPGP library"
[Proton Mail]: https://proton.me/mail "Proton Mail — Get a private, secure, and encrypted email"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
