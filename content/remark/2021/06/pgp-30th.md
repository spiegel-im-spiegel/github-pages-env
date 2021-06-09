+++
title = "PGP は30周年だった"
date =  "2021-06-09T20:04:38+09:00"
description = "PGP の最初のバージョンがリリースされたのが1991年だから，そうか，そうだよな。"
image = "/images/attention/kitten.jpg"
tags = ["cryptography", "openpgp"]
pageType = "text"
draft = true

[scripts]
  mathjax = false
  mermaidjs = false
+++

元祖 PGP (Pretty Good Privacy) の作者である Phil. Zimmermann 氏が30周年の記念ブログ記事（？）を上げている。

- [PGP 30th Anniversary](https://philzimmermann.com/EN/essays/PGP_30th/)
- [ブログ: PGP誕生30周年](https://okuranagaimo.blogspot.com/2021/06/pgp30.html)

PGP の最初のバージョンがリリースされたのが1991年[^rfc1991] だから，そうか，そうだよな。

[^rfc1991]: 当時の PGP の仕様をまとめたものが [RFC 1991] として残されている。

PGP にまつわる話については Steven Levy 原著作の『[暗号化](https://www.amazon.co.jp/dp/4314009071?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』を参照するといいだろう。
つっても絶版みたいなので古本を漁るしかないようだが。
PGP に限らず暗号周りについて勉強したいなら，基礎教養としてこの本を読むことを強くお勧めする。

私が PGP を使い出したのは1990年代後半で，既に PGP 5.x にリリースされていた。
初期の PGP が問題だったのは，作者が暗号特許について迂闊だったこと，当時の暗号技術はまだ軍事物資として輸出制限がかけられていたことである。
これらの問題が根本的に解消されるには西暦2000年まで待つ必要があったのだ。

ちなみに PGP 5.x 国際版では，特許に抵触する暗号アルゴリズム（RSA や IDEA など）は実装から外され，書籍の形でリリースされていた（書籍は輸出制限の対象外だった）。
メンテナンス・配布・ビルドの一連は世界中のボランティアが行っていた。
こういう経緯があったので，私は今だに RSA 暗号は使いたくないんだよなぁ。

現在は PGP の設計コンセプトは OpenPGP に継承され [RFC 4880]/[RFC 5581]/[RFC 6637] として標準化されている。
また AEAD (Authenticated Encryption with Associated Data) 等をサポートする次期 OpenPGP が [RFC 4880bis] として検討されている。
実装に関しては，事実上のリファレンス実装である [GnuPG](https://gnupg.org/ "The GNU Privacy Guard") 以外にも [OpenPGP.js](http://openpgpjs.org/ "OpenPGP.js | OpenPGP JavaScript Implementation") や Rust 製の [Sequoia-PGP](https://sequoia-pgp.org/) といったものがある。

そういや Symantec 社って Broadcom 社 → Accenture 社とたらい回しにされてる感じだが PGP はどうなってるんだろう。

## ブックマーク

- [OpenPGP の実装]({{< rlnk "/openpgp/" >}})

[RFC 1847]: https://tools.ietf.org/html/rfc1847 "RFC 1847 - Security Multiparts for MIME: Multipart/Signed and Multipart/Encrypted"
[RFC 1991]: https://tools.ietf.org/html/rfc1991 "RFC 1991 - PGP Message Exchange Formats"
[RFC 2015]: https://tools.ietf.org/html/rfc2015 "RFC 2015 - MIME Security with Pretty Good Privacy (PGP)"
[RFC 2440]: https://tools.ietf.org/html/rfc2440 "RFC 2440 - OpenPGP Message Format"
[RFC 3156]: https://tools.ietf.org/html/rfc3156 "RFC 3156 - MIME Security with OpenPGP"
[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 5581]: https://tools.ietf.org/html/rfc5581 "RFC 5581 - The Camellia Cipher in OpenPGP"
[RFC 6637]: https://tools.ietf.org/html/rfc6637 "RFC 6637 - Elliptic Curve Cryptography (ECC) in OpenPGP"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "4900900028" %}} <!-- PGP―暗号メールと電子署名 -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
