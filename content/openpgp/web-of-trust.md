+++
title = "OpenPGP の電子署名は「ユーザーの身元を保証し」ない"
date = "2019-03-21T00:28:40+09:00"
description = "つまり公開鍵や電子署名で「ユーザーの身元を保証」するのではなく「身元の保証されたユーザ」同士で鍵と電子署名を運用するのである。"
image = "/images/attention/openpgp.png"
tags = [
  "cryptography",
  "openpgp",
  "management",
  "pki",
  "trust",
  "gnupg",
  "git",
  "github",
  "certification",
]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

重箱の隅を突っつくような記事で申し訳ないのだが

- [「GitKraken 5.0」がリリース ～GPGコミットや“Interactive Rebase”をサポート - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1175019.html)

という記事で

{{< fig-quote title="「GitKraken 5.0」がリリース ～GPGコミットや“Interactive Rebase”をサポート" link="https://forest.watch.impress.co.jp/docs/news/1175019.html" >}}
<q>メジャーアップデートとなる本バージョンでは、“GNU Privacy Guard (GPG)”による署名付きのコミットがサポートされた。ユーザーの身元を保証し、他のユーザーによるなりすましを防止することができる。</q>
{{< /fig-quote >}}

などと書いてあって「それはちゃうやろ」という話。

## 暗号機能の4要素

いわゆる暗号機能の4要素は以下の通り

- 機密性（Confidentiality）
- 完全性（Integrity）
- 認証（Authentication）
- 否認防止（Non-repudiation）

このうちデータへの電子署名では主に完全性と否認防止を行う。
否認防止という言葉はちょっと耳慣れないかもしれないが，要するに「あなたはこのデータに署名した」という事実を否認することが出来ない，という意味である。

これらを達成するためには電子署名に使う公開鍵がユーザと正しく紐付いている必要がある。
おそらく最初の記事は「公開鍵がユーザと正しく紐付いている」という前提で「ユーザーの身元を保証」などと書いているのかもしれないが，話はそう簡単ではないのだ。

## 公開鍵基盤

### X.509

いわゆる「公開鍵基盤（Public-Key Infrastructure; PKI）」として最も有名なのは HTTPS 接続で使われている X.509 であろう。
X.509 の「信用モデル（trust model）」では hierarchical な CA を構成し，その CA が公開鍵を「証明（certification）」することで成り立っている。

でも実は CA が保証しているのは「正当な公開鍵」であることだけで，その鍵の「ユーザーの身元を保証し」ているわけではない。
そこで HTTPS には [EV SSL (Extended Validation SSL) なる奇っ怪な仕組み](https://baldanders.info/spiegel/log2/000277.shtml "Extended Validation SSL — Baldanders.info")が組み込まれた。
これは証明書の「ユーザーの身元を保証」するための仕組みで，ユーザは CA に対して自身の身元を証明するものを提出し CA は公開鍵の管理をより厳格に行う，ということらしい。

正直に言って「屋上屋を架す」仕組みであり CA の責務を逸脱していると思うのだが，まぁ深くは突っ込むまい。

### Web of Trust

これに対して PGP/[OpenPGP] が伝統的に執っている信用モデルは「信用の輪（web of trust）」と呼ばれている。

これは要するに「友達の友達は友達」というやつで，ユーザ同士がお互いの鍵を相互に認証（authentication）することで信用を構成する仕組みである。

何故 PGP/[OpenPGP] が X.509 のような「権威による認証」を採用しなかったかというと，それは PGP の出自に関係がある。
PGP の作者である Phil Zimmermann は，当時は反核運動家であり国家等の「権威」に依らない信用モデルを必要としていたわけだ。

OpenPGP は名前だけならどんな鍵でも作れる。
たとえば Bitcoin の作者と言われる Satoshi Nakamoto の公開鍵は[公開鍵サーバを探せば簡単に見つかる]({{< ref "/remark/2016/05/openpgp-key-of-satoshi-nakamoto.md" >}})が，それが「あの」 Satoshi Nakamoto 本人であると示す証拠は（少なくとも公開鍵自体には）存在しない。
OpenPGP 公開鍵やその電子署名で赤の他人の「ユーザーの身元を保証」することは出来ないわけだ。

## Git を中心としたチーム運営に [OpenPGP] を利用する

じゃあ git commit で OpenPGP 署名を付与することにどんな意義があるかというと，それはチーム運営で威力を発揮する。
つまり公開鍵や電子署名で「ユーザーの身元を保証」するのではなく「身元の保証されたユーザ」同士で鍵と電子署名を運用するのである。
この辺について詳しくは，拙文「[OpenPGP 鍵管理に関する考察]({{< relref "./openpgp-key-management.md" >}})」を参照してほしい。

また GitHub のようにアカウントと公開鍵を紐つけることによってサービス内における強力なポートフォリオとして機能している点は見逃せないだろう。

ホンマ git ってよく出来たシステムだよなぁ。

## ブックマーク

- [Git Commit で OpenPGP 署名を行う]({{< relref "./git-commit-with-openpgp-signature.md" >}})
- [わかる！ OpenPGP 暗号 — Baldanders.info](https://baldanders.info/spiegel/archive/pgpdump/openpgp.shtml)

- [spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer](https://github.com/spiegel-im-spiegel/gpgpdump) : OpenPGP 鍵や電子署名のダンプには拙作をどうぞ（宣伝`w`）

[OpenPGP]: http://openpgp.org/
[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E5%8C%96-%E3%83%97%E3%83%A9%E3%82%A4%E3%83%90%E3%82%B7%E3%83%BC%E3%82%92%E6%95%91%E3%81%A3%E3%81%9F%E5%8F%8D%E4%B9%B1%E8%80%85%E3%81%9F%E3%81%A1-%E3%82%B9%E3%83%86%E3%82%A3%E3%83%BC%E3%83%96%E3%83%B3%E3%83%BB%E3%83%AC%E3%83%93%E3%83%BC/dp/4314009071?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4314009071"><img src="https://images-fe.ssl-images-amazon.com/images/I/51ZRZ62WKCL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E5%8C%96-%E3%83%97%E3%83%A9%E3%82%A4%E3%83%90%E3%82%B7%E3%83%BC%E3%82%92%E6%95%91%E3%81%A3%E3%81%9F%E5%8F%8D%E4%B9%B1%E8%80%85%E3%81%9F%E3%81%A1-%E3%82%B9%E3%83%86%E3%82%A3%E3%83%BC%E3%83%96%E3%83%B3%E3%83%BB%E3%83%AC%E3%83%93%E3%83%BC/dp/4314009071?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4314009071">暗号化 プライバシーを救った反乱者たち</a></dt>
	<dd>スティーブン・レビー</dd>
	<dd>斉藤 隆央 (翻訳)</dd>
    <dd>紀伊國屋書店 2002-02-16</dd>
    <dd>Book 単行本</dd>
    <dd>ASIN: 4314009071, EAN: 9784314009072</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">20世紀末，暗号技術の世界で何があったのか。知りたかったらこちらを読むべし！</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-16">2018-12-16</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/PGP%E2%80%95%E6%9A%97%E5%8F%B7%E3%83%A1%E3%83%BC%E3%83%AB%E3%81%A8%E9%9B%BB%E5%AD%90%E7%BD%B2%E5%90%8D-%E3%82%B7%E3%83%A0%E3%82%BD%E3%83%B3-%E3%82%AC%E3%83%BC%E3%83%95%E3%82%A3%E3%83%B3%E3%82%B1%E3%83%AB/dp/4900900028?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4900900028"><img src="https://images-fe.ssl-images-amazon.com/images/I/5132396FFQL._SL160_.jpg" width="124" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/PGP%E2%80%95%E6%9A%97%E5%8F%B7%E3%83%A1%E3%83%BC%E3%83%AB%E3%81%A8%E9%9B%BB%E5%AD%90%E7%BD%B2%E5%90%8D-%E3%82%B7%E3%83%A0%E3%82%BD%E3%83%B3-%E3%82%AC%E3%83%BC%E3%83%95%E3%82%A3%E3%83%B3%E3%82%B1%E3%83%AB/dp/4900900028?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4900900028">PGP―暗号メールと電子署名</a></dt>
	<dd>シムソン ガーフィンケル</dd>
	<dd>Simson Garfinkel (原著), ユニテック (翻訳)</dd>
    <dd>オライリー・ジャパン 1996-04</dd>
    <dd>Book 単行本</dd>
    <dd>ASIN: 4900900028, EAN: 9784900900028</dd>
    <dd>評価<abbr class="rating fa-sm" title="3">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">良書なのだが，残念ながら内容が古すぎた。 PGP の歴史資料として読むならいいかもしれない。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2014-10-16">2014-10-16</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

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
