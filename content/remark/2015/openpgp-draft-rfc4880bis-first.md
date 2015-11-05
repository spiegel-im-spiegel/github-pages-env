+++
date = "2015-11-05T16:44:46+09:00"
description = "おおっ！ 次期 OpenPGP の最初のドラフトが登場している。"
draft = true
tags = ["security", "cryptography", "openpgp", "rfc"]
title = "OpenPGP: First RFC4880bis Draft"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

おおっ！ 次期 OpenPGP の最初のドラフトが登場している。

- [First 4880bis drafts](https://mailarchive.ietf.org/arch/msg/openpgp/uUKa8eQzWOh3quNElu0BDNrKi2o)
    - [OpenPGP Message Format draft-koch-openpgp-rfc4880bis-00](https://tools.ietf.org/id/draft-koch-openpgp-rfc4880bis-00.txt) （[差分](https://tools.ietf.org/rfcdiff?url2=draft-koch-openpgp-rfc4880bis-00.txt)）
    - [OpenPGP Message Format draft-koch-openpgp-rfc4880bis-01](https://tools.ietf.org/id/draft-koch-openpgp-rfc4880bis-01.txt) （[差分](https://tools.ietf.org/rfcdiff?url2=draft-koch-openpgp-rfc4880bis-01.txt)）

差分を簡単に見れるのは便利だな。

Repository も公開されている。
以下の URI で取得すればよい。

```bash
$ git clone git://git.gnupg.org/people/wk/rfc4880bis.git
```

- [git.gnupg.org Git - people/wk/rfc4880bis.git/summary](http://git.gnupg.org/cgi-bin/gitweb.cgi?p=people/wk/rfc4880bis.git)

[RFC 4880](https://tools.ietf.org/html/rfc4880) が[リリースされたのが2007年](http://www.baldanders.info/spiegel/log2/000356.shtml)で，その後に登場した日本国産の Camellia 暗号（[RFC 5581](https://tools.ietf.org/html/rfc5581)）や ECC（Elliptic Curve Cryptography; [RFC 6637](https://tools.ietf.org/html/rfc6637)）が今回統合される。
以前 [SHA-3 が OpenPGP に組み込まれるかも](http://www.baldanders.info/spiegel/log2/000866.shtml)，という話があったが，今のところドラフト版には記述がない感じ。
Fingerprint を SHA-2 ベースにするとかいう議論もあったような気がするが，これもないか？

まぁ週末にでもゆっくり読むか。

## 参考ページ

- [わかる！ OpenPGP 暗号 — Baldanders.info](http://www.baldanders.info/spiegel/archive/pgpdump/openpgp.shtml)

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">第3版出た！ てか，もう Kindle 版出てるのか。紙の本買うのはやまったかなぁ。 SHA-3 や BitCoin/BlockChain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
