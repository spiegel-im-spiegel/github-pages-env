+++
title = "Ubuntu で音楽 CD のリッピング"
date =  "2019-06-10T22:59:23+09:00"
description = "Windows のときは文字エンコーディングで苦労したからねぇ。 UNIX 系は楽でいいわ（笑）"
image = "/images/attention/kitten.jpg"
tags = [ "ubuntu", "music", "tools" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

いやぁ，ひーさしぶりに音楽 CD 買ったですよ。
最近はストリーミング・サービスか MP3 で買うかだったもんねぇ。

ちうわけで [Ubuntu] でリッピング・ツールを探したら [Asunder] というのがいいらしい。
インストールは APT でできる。

```text
$ sudo apt install asunder lame
```

`lame` は MP3 エンコーディングを扱うのに必要らしい。

起動したらちゃんと日本語で表示できてた。
えらいえらい！

{{< fig-img src="./asunder.png" link="./asunder.png" width="817" >}}

CDDB のプロキシ・サーバに [`freedbtest.dyndns.org:80`](https://freedbtest.dyndns.org/ "freedb 日本語") を指定すれば日本の CD 情報もちゃんと取れるようだ。

Windows のときは文字エンコーディングで苦労したからねぇ。
UNIX 系は楽でいいわ（笑）

## ブックマーク

- [UbuntuやArch Linuxで音楽CDの取り込み Asunderの紹介!! - BobbyQuineのブログ(備忘録)](http://bobbyquine.hatenablog.com/entry/2018/01/16/005131)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Asunder]: http://www.littlesvr.ca/asunder/

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%82%A2%E3%83%AB%E3%83%86%E3%82%A3%E3%83%A1%E3%83%83%E3%83%88%E2%98%86MAGIC-CD-i%E2%98%86Ris/dp/B07PBB8W1X?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B07PBB8W1X"><img src="https://images-fe.ssl-images-amazon.com/images/I/61fwAZoNBfL._SL160_.jpg" width="160" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%82%A2%E3%83%AB%E3%83%86%E3%82%A3%E3%83%A1%E3%83%83%E3%83%88%E2%98%86MAGIC-CD-i%E2%98%86Ris/dp/B07PBB8W1X?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B07PBB8W1X">アルティメット☆MAGIC *CD</a></dt>
    <dd>エイベックス・ピクチャーズ株式会社(Music) 2019-05-21 (Release 2019-05-22)</dd>
    <dd>Music CD</dd>
    <dd>ASIN: B07PBB8W1X, EAN: 4562475295115</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">アニメ「<a href="http://kenja-no-mago.jp/">賢者の孫</a>」の OP 曲を収録。 CD 版で購入したのだが <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B07RMKKSG9/baldandersinf-22/">MP3 版</a>もあった orz 通勤時に聴くと気合が入る（笑）</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-06-10">2019-06-10</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E9%AD%94%E6%B3%95%E3%81%AE%E5%A4%A9%E4%BD%BF-%E3%82%AF%E3%83%AA%E3%82%A3%E3%83%9F%E3%83%BC%E3%83%9E%E3%83%9F-%E5%85%AC%E5%BC%8F%E3%83%88%E3%83%AA%E3%83%93%E3%83%A5%E3%83%BC%E3%83%88%E3%83%BB%E3%82%A2%E3%83%AB%E3%83%90%E3%83%A0-Various-artists/dp/B007EEGL58?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B007EEGL58"><img src="https://images-fe.ssl-images-amazon.com/images/I/51kqeF4UfiL._SL160_.jpg" width="160" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E9%AD%94%E6%B3%95%E3%81%AE%E5%A4%A9%E4%BD%BF-%E3%82%AF%E3%83%AA%E3%82%A3%E3%83%9F%E3%83%BC%E3%83%9E%E3%83%9F-%E5%85%AC%E5%BC%8F%E3%83%88%E3%83%AA%E3%83%93%E3%83%A5%E3%83%BC%E3%83%88%E3%83%BB%E3%82%A2%E3%83%AB%E3%83%90%E3%83%A0-Various-artists/dp/B007EEGL58?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B007EEGL58">魔法の天使 クリィミーマミ 公式トリビュート・アルバム</a></dt>
	<dd>Various artists (メインアーティスト)</dd>
    <dd>P-VINE (Release 2012-04-01)</dd>
    <dd>Digital Music Album MP3 ダウンロード</dd>
    <dd>ASIN: B007EEGL58</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">往年のクリィミーマミ・ファンが泣いて喜ぶアルバム。 CD 版もあるが MP3 でも買える。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-06-10">2019-06-10</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>
