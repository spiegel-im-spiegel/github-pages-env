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
  url       = "http://www.baldanders.info/spiegel/profile/"
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

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
