+++
date = "2015-11-05T20:39:13+09:00"
description = "おおっ！ 次期 OpenPGP の最初のドラフトが登場している。"
draft = false
tags = ["security", "cryptography", "openpgp", "rfc"]
title = "OpenPGP: First RFC4880bis Draft"

[scripts]
  mathjax = false
  mermaidjs = false
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

現行の [RFC 4880](https://tools.ietf.org/html/rfc4880) が[2007年にリリース](https://baldanders.info/blog/000356/)され，その後に追加された日本国産の Camellia 暗号（[RFC 5581](https://tools.ietf.org/html/rfc5581)）や ECC（Elliptic Curve Cryptography; [RFC 6637](https://tools.ietf.org/html/rfc6637)）が今回ひとつに統合される。
以前 [SHA-3 が OpenPGP に組み込まれるかも](https://baldanders.info/blog/000866/)，という話があったが，今のところドラフト版には記述がない感じ。
Fingerprint を SHA-2 ベースにするとかいう議論もあったような気がするが，これもないか？

まぁ週末にでもゆっくり読むか。

## 参考ページ

- [わかる！ OpenPGP 暗号 — Baldanders.info](https://baldanders.info/spiegel/cc-license/)

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
