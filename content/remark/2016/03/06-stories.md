+++
date = "2016-03-06T18:20:23+09:00"
description = "公開鍵暗号の研究者がチューリング賞を受賞 / 出版社（者）としての青空文庫と aozorahack / 避難訓練と 3.11 / ユーザが EC に求めること / その他の気になる記事"
draft = false
tags = ["math", "cryptography", "book", "aozora", "environment", "risk", "management", "market"]
title = "週末スペシャル： 公開鍵暗号の研究者がチューリング賞を受賞"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"
+++

2月は逃げました。
3月ですよ。
年度末ですよ。

1. [公開鍵暗号の研究者がチューリング賞を受賞]({{< relref "#dh" >}})
1. [出版社（者）としての青空文庫と aozorahack]({{< relref "#aozora" >}})
1. [避難訓練と 3.11]({{< relref "#risk" >}})
1. [ユーザが EC に求めること]({{< relref "#ec" >}})
1. [その他の気になる記事]({{< relref "#other" >}})

## 公開鍵暗号の研究者がチューリング賞を受賞{#dh}

- [コンピューティング分野のノーベル賞ことチューリング賞、公開鍵暗号の研究者に与えられる : ギズモード・ジャパン](http://www.gizmodo.jp/2016/03/Turing_Award_for_public_key_encryption.html)

「チューリング賞（ACM A. M. Turing Award）」はコンピュータ科学の分野での最高権威の賞と言われている。
この賞に（両氏の名を冠した） “Diffie-Hellman” 暗号方式の発明者である Whitfield Diffie, Martin E. Hellman 両氏が今まで受賞されていなかったことのほうが驚きだが，公開鍵暗号といえば RSA のほうが真っ先に浮かんでしまうからだろうか。
今ごろ受賞というのも政治臭プンプンだが，功績が認められるということはいいことである。

具体的な論文はこちら。

- {{< pdf-file title="New Directions in Cryptography" link="https://www-ee.stanford.edu/~hellman/publications/24.pdf" >}}

20世紀後半は暗号の暗黒時代とも言われている。
もともと暗号技術は軍事技術の一種として使われることが多かったが，第2次世界大戦の前後から重要性が高まり，国家による統制が激しくなった。
これを変えるきっかけとなったもののひとつが公開鍵暗号である。

詳しくは Steven Levy さんの『[暗号化 プライバシーを救った反乱者たち](https://www.amazon.co.jp/exec/obidos/ASIN/4314009071/baldandersinf-22/)』をどうぞ。

## 出版社（者）としての青空文庫と aozorahack{#aozora}

[OSC Tokyo 2016/Spring](http://www.ospn.jp/osc2016-spring/ "オープンソースカンファレンス2016 Tokyo/Spring - オープンソースの文化祭！") のときの発表資料だそうな。

{{< fig-slideshare id="itRRJPdUruVE52" title="aozorahackと青空文庫の現状とこれから (OSC 2016 Tokyo/Spring)" link="http://www.slideshare.net/takahashim/osc-2016-tokyospring-aozorahack" width="425" height="355" >}}

この中で青空文庫の出版社（者）的な機能に注目している。
つまり著作権の切れた書籍に対して入力・校正・公開を行うプロセスは出版（publishing）そのものではないかという指摘である。
で，青空文庫の「図書館」機能と「出版」機能を分離することでより多くの人を巻き込むことができるのではないかという提案のようだ。

具体的には管理用の DB から公開可能なサブセットを公開 DB として使えるようにしたいらしい。

{{< fig-img src="https://farm2.staticflickr.com/1473/24920649843_83bff63a6e.jpg" title="aozorahackと青空文庫の現状とこれから (OSC 2016 Tokyo/Spring)" link="http://www.slideshare.net/takahashim/osc-2016-tokyospring-aozorahack" >}}

多分ファイルも公開可能なものを別にして切り出したほうがいいのだろう（公開していない校正中の作品にはまだ著作権が切れていないものもあるので）。
公開 DB およびファイルにアクセスする API を作れば様々な人が様々なシーンで公有作品を利用できる。
また翻訳作品以外の Free Culture Licenses[^fcl] の作品も多く含めることができれば青空文庫の幅が広がるかもしれない。

[^fcl]: [CC0](https://creativecommons.org/publicdomain/zero/1.0/deed.ja), [by](https://creativecommons.org/licenses/by/4.0/deed.ja), [by-sa](https://creativecommons.org/licenses/by-sa/4.0/deed.ja) を合わせて Free Culture Licenses と呼ぶ。

## 避難訓練と 3.11{#risk}

- [3.11WALK - 3.11は、歩いて帰ろう。](http://311walk.jp/)

大規模災害は滅多に起こらない。
だから，いざ起きた時，つまり「事後」にどう備えるか。
ただ単に情緒的に 3.11 を思い出すのではなく，これから起きる未来を想定して「避難訓練」をしていくことはとても重要だと思う。

「想定外を想定する」のは結構難しいが必要なことである。
何事もね。

## ユーザが EC に求めること{#ec}

- [PayPalとJECCICAが「EC戦略白書」発表、EC企業と顧客の認識差が浮き彫り - ZDNet Japan](http://japan.zdnet.com/article/35078917/)

{{< fig-quote title="PayPalとJECCICAが「EC戦略白書」発表、EC企業と顧客の認識差が浮き彫り" link="http://japan.zdnet.com/article/35078917/" >}}
<q>ユーザーが中小ECサイトを利用しない理由として、（1）存在が知られていない（44.5%）、（2）会員登録が面倒（34.4％）、（3）商品数の少なさ（25.6％）、（4）セキュリティ面の不安（17.5%）--の順で回答が多かった。商品の少なさやセキュリティよりも、「会員登録が面倒」という理由の方が強いことは注目に値する。</q>
{{< /fig-quote >}}

「会員登録が面倒」というのはよく分かる。
ユーザは「一見客」として来ているのに会員登録を要求されるのはウザいことこの上ない。
問題は，そういうユーザの心理をサービス・プロバイダ側が認識していなかったということだろう。
ユーザを囲い込むことに躍起になって，ユーザのことを全く考えていないという証拠でもある。

{{< fig-quote title="PayPalとJECCICAが「EC戦略白書」発表、EC企業と顧客の認識差が浮き彫り" link="http://japan.zdnet.com/article/35078917/" >}}
<q>この調査結果を踏まえて川連氏は、中小EC企業はユーザーの「囲い込み戦略」を止めるべきだと断言した。国内ECサイトの売上トップ100社のうち会員登録を必須とするサイトは7割。それに対して、米国のECサイト売上トップ100社で会員登録が必須なのは2割程度だ。さらにトップ25社に絞ると、国内21社に対して米国は2社となる。「すでに、グローバルトレンドはゲスト購入に移行しつつある」（川連氏）。</q>
{{< /fig-quote >}}

似たような発想は実際の店舗でもあって

- [飲食店の「常連作り」支援に向け、予約台帳のトレタがPOSシステム5社と連携へ | TechCrunch Japan](http://jp.techcrunch.com/2016/03/01/toreta-pos/)

もう「常連」を囲い込みの道具としか見なしていないことがミエミエの最低マーケティングである。

POS などのシステムはユーザを「スマート・モブ（smart mob）」と見なし行動追跡しながら保険統計学的に評価する。
近年はやりの「ビッグ・データ（big data）」も同じである。
これはこれで大事な分析だが，そのような視点から「常連」が生まれることはない。

「固定客」と「常連客」は全くベクトルが違う。
足繁く通っているからといって客側が自身を常連とは認識してないこともあるし（毎朝マックでコーヒーを飲んでるからといって自分がその店の常連だとは思ってないだろう），逆に数年に1度しか来ない「常連客」だっているのだ[^a]。

[^a]: 広島市は「支店都市」なのでだいたい数年単位で人が入れ替わる。故にわざわざ遠方から来る「滅多に来ない常連客」だっているのである。でも「滅多に来ない常連客」がいるというのは，その店が長く続いている証でもある。繁華街では「2年保てば上出来，3年保てば老舗」と言われるくらいお店の入れ替わりも激しい。

かつての「クーポン」戦略が失敗したのは，ユーザを囲い込んでいるつもりで実は「クーポン」という土俵ですべての店舗が相対評価されてしまっている点である。
ユーザは数多ある「クーポン」の中から一番お得なものを都度選択しているだけで，お店に対する愛着も敬意もない。
ただ「クーポンを使えばお得」という感想が残るだけだ。

EC も同じこと。
ユーザから見てその店舗に惹きつけられる何かがあれば囲い込む必要はないし，そこで勝負していかなければ結局は「EC」という括りの中で相対評価されるだけで「固定客」にも「常連客」にもならない。

## その他の気になる記事{#other}

あとで個別に記事にするかもしれないが，とりあえずブックマークのみ。

- [北海道のシャッター通りに本屋をつくる « マガジン航[kɔː]](http://magazine-k.jp/2016/03/02/little-bookstore-in-northern-street/)
- [The Importance of Strong Encryption to Security - Schneier on Security](https://www.schneier.com/blog/archives/2016/02/the_importance_.html)
- [「Ruby」をWindows環境へ簡単導入できる「Rumix 2」が更新。「Ruby 2.2」に対応 - 窓の杜](http://www.forest.impress.co.jp/docs/news/20160229_745961.html)
- [「丸善＆ジュンク堂ネットストア」がネット書店「honto」に統合、hontoでも店舗取り置きサービスが利用可能に -INTERNET Watch](http://internet.watch.impress.co.jp/docs/news/20160301_746190.html)
- [OAuth 2.0 + OpenID Connect のフルスクラッチ実装者が知見を語る - Qiita](http://qiita.com/TakahikoKawasaki/items/f2a0d25a4f05790b3baa)
- [【第二弾】OAuth 2.0 + OpenID Connect のフルスクラッチ実装者が知見を語る - Qiita](http://qiita.com/TakahikoKawasaki/items/30fbd546935cea914e4f)
- [GoTTY 良さそう - Qiita](http://qiita.com/A-I/items/49bcb56ed977b4fb36ff)
    - [GoTTYでブラウザからルータを操作してみた - Qiita](http://qiita.com/kooshin/items/8c7dbfc9c5e8a88d1705)

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
