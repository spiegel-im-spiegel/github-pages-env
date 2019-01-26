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
  url       = "https://baldanders.info/spiegel/profile/"
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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE"><img src="https://images-fe.ssl-images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE">暗号技術入門 第3版　秘密の国のアリス</a></dt>
	<dd>結城 浩</dd>
    <dd>SBクリエイティブ 2015-08-25 (Release 2015-09-17)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B015643CPE</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
