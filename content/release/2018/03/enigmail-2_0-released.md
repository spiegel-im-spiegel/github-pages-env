+++
title = "Enigmail 2.0 Released"
date = "2018-03-27T21:13:01+09:00"
description = "って Autocrypt も p≡p もまだ調査してないよ！"
image = "/images/attention/tools.png"
tags = [
  "tools",
  "security",
  "cryptography",
  "openpgp",
  "autocrypt",
  "gnupg",
]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Thunderbird] 用の [OpenPGP] 暗号化アドオンである [Enigmail] が 2.0 になったようだ[^gpg1]。

[^gpg1]: 実際の処理には [GnuPG] が必要。

主な変更点は以下の通り。

- *メッセージの件名の保護：* 件名を暗号データに含めることができるようになった。 PGP/MIME 暗号化でのみ有効
- *暗号化および署名ボタンの改良：* S/MIME 暗号化と統合
- *Autocrypt のサポート：* 新しい鍵配布標準である [Autocrypt] に対応した
- *p≡p (Pretty Easy Privacy) のサポート：* [p≡p] Junior Mode をサポート。ただし現時点では [p≡p] を手動でインストールする必要がある

って [Autocrypt] も [p≡p] もまだ調査してないよ！

- [Email green, secure, simple and ad-free - posteo.de - New: Easy email encryption with Autocrypt and OpenPGP header](https://posteo.de/en/blog/new-easy-email-encryption-with-autocrypt-and-openpgp-header)

そのうち試そう。
鍵管理が楽になるのは嬉しい。

[Thunderbird]: https://www.mozilla.org/thunderbird/ "Thunderbird — Software made to make email easier. — Mozilla"
[Enigmail]: https://addons.mozilla.org/ja/thunderbird/addon/enigmail/ "Enigmail :: Thunderbird 向けアドオン"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenPGP]: http://openpgp.org/
[Autocrypt]: https://autocrypt.org/ "Autocrypt 1.0.0 documentation"
[p≡p]: https://pep-project.org/

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
