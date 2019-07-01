+++
date = "2015-11-03T22:05:33+09:00"
update = "2016-12-26T21:07:45+09:00"
description = "人工知能で弁護士は絶滅するか / AI は「トロッコ問題」をどう解くか / TVer に拒否られた / Node v5.0.0 (Stable)"
draft = false
tags = ["artificial-intelligence", "media", "nodejs"]
title = "今日の戯れ言： 人工知能で弁護士は絶滅するか"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "https://baldanders.info/profile/"
+++

週末に書く予定だった記事。

1. [人工知能で弁護士は絶滅するか]({{< relref "#lawyer" >}})
1. [AI は「トロッコ問題」をどう解くか]({{< relref "#trolley-problem" >}})
1. [TVer に拒否られた]({{< relref "#tver" >}})
1. [Node v5.0.0 (Stable)]({{< relref "#node" >}})

## 人工知能で弁護士は絶滅するか{#lawyer}

- [人工知能は弁護士を絶滅させていく：米国での調査結果 « WIRED.jp](http://wired.jp/2015/10/28/computers-replacing-lawyers/)

いやぁ？ たぶん「弁護士」という名前の意味が変わるだけじゃないのかな。

[「ロボット」の語源は「奴隷（robata）」であり「労働（robota）」でもある](https://wirelesswire.jp/2015/10/47216/)そうだ。

近代における「奴隷解放」は単に奴隷を解放しただけではなく，それまで奴隷がやっていた労働を一般に開放することを意味する。
これにより「労働」の意味が変わってしまった。

{{< fig-youtube id="Ps6ck1ejoAw" width="500" height="375" title="Smile, Charlie Chaplin , Modern Times, 1936 - YouTube" >}}

さらに「ロボット」の台頭は，人々から，その（奴隷解放以後の）「労働」の一部を奪った。
しかも単に人間から「労働」の一部を奪っただけでなく，「労働」の意味が再び変節する。
現在「ロボット」に対抗できる労働者は，それこそ奴隷的な扱いを受けている人たち（たとえば日本なら黒い企業の従業員）だけだろう。

{{< fig-quote title="「海外生産が安い」はもう古い、エプソンの国内回帰戦略 - 次世代工場 - 日経テクノロジーオンライン" link="http://techon.nikkeibp.co.jp/atcl/news/15/102300871/" >}}
「2013年以降、労働集約型の海外工場に比べて、自動化設備を積極的に導入した国内工場の方が労務費の面で安く抑えられている」
{{< /fig-quote >}}

同様に AI が本格的に台頭してくれば「労働」の意味が三たび変わる。
特に今まで「知的労働」と呼ばれていたものの一部はネガティブな意味になるかもしれない[^a]。
どれだけ知識や経験があろうとも，ただ入力に反応しているだけの「オートマトン（automaton）」な人は AI に絶滅させられる。

[^a]: 「プログラマ」とかね。まぁ日本では既にプログラマを「IT 土方」とか言ったりするけど（笑）

これからは自分で問いを立て，そこから何かを生み出せる人でなければ食いっぱぐれる（または機械以下の階層で使役されるだけの存在になる）。
[もう一度引用する](https://medium.com/@spiegel/-8fcccfb661 "続ける（続いてる）だけではダメ — Medium")が

> 「知識のストックに価値があった時代」から、ネットの普及による「知識や価値を生み出す力なしには生き残れない時代」への大転換

を迫られている，ということだ。
年寄りにはキツい話だが，しょうがない。

## AI は「トロッコ問題」をどう解くか{#trolley-problem}

- [完全自動運転自動車とトロッコ問題について](http://blogos.com/article/142284/)

あー。
これって「[われはロボット](https://www.amazon.co.jp/exec/obidos/ASIN/B00O1VK072/baldandersinf-22/)」だよね。

明らかに「正しい解」がない場合，いくつかの近似解の中から妥当と思われるものを選ぶしかない[^b]。
でも，どの解を選んでも結局「正しい解」ではないのだ。
だから人は葛藤し，さらに後悔する。
「緊急避難」というのは「正しい解」が存在しない場合に「近似解でいいんだよ」ということを法的に担保するものだ[^c]。

[^b]: もちろん「何も選ばない」というのも選択肢のひとつである。
[^c]: しかし，法的に担保されているからといって倫理・道徳的に問題がないとは限らない。しかも倫理観・道徳観念というのは，特に個人主義が進んだ現代では，かつての「大きな物語（meta-narrative）」ほどには機能しない。

AI は（今のところ）近似解に葛藤したりしない。
もちろん後悔だってしない。
その解に辿り着いたのは，機械が自ら考えたのではなく，あくまでも設定された（あるいは構築された）論理に沿って必然的に導かれたものだから（たとえその道筋が人には理解できないものだとしても）。
じゃあ機械が導き出した解を実行した結果の責任は誰が取るの？ ってことである。

もし AI が近似解に葛藤することがあるなら，それは「知的生命」が新しい階梯に進んだことを意味する。

### 参考

- [自動運転車の法律問題を総括すると見えてくる難解な課題 - 風観羽　情報空間を羽のように舞い本質を観る](http://d.hatena.ne.jp/ta26/20151104)

## TVer に拒否られた{#tver}

- [民放公式テレビポータル「TVer（ティーバー）」 - Google Play](https://play.google.com/store/apps/details?id=jp.hamitv.hamiand1)
- [民放キー局5社の無料見逃し配信「TVer」アプリが登場 | アプリオ](http://appllio.com/20151026-7241-tver-app)

まだ対応している番組は少ないそうだが，面白そうなのでインストールしようとしたら Nexus 7 は対応してないと言われた。
どうやらスマホだけ？ らしい。
なにそれ。
やっぱテレビは糞だな。
スマホでちまちま動画なんか見てられねーよ！

まぁ Hulu である程度見れるし，見れない番組は大きな声では言えないサイトで見ればいいや。

## Node v5.0.0 (Stable){#node}

- [Node v5.0.0 (Stable) | Node.js](https://nodejs.org/en/blog/release/v5.0.0/)
    - [Node v5.0.0 (stable) - Qiita](http://qiita.com/zakiko/items/f1e6db17e243667b8513)
- [Node.js 5.0がリリース。奇数バージョンは最新機能版、偶数バージョンは長期サポート版 － Publickey](http://www.publickey1.jp/blog/15/nodejs_50.html)

[LTS (Long-term Support)](https://github.com/nodejs/LTS/) の 4.2（11月4日時点の最新は v4.2.2）とは別に v5.0.0 が登場。
LTS 版って Firefox で言うところの [ESR (Extended Support Release) 版](http://www.mozilla.jp/business/downloads/)みたいなものかな？

まだ v4 すら試してないんだよなぁ。
そもそも ECMAScript 2015 (ES6) に脳みそが追いついてない。
なんとかしないとなぁ。

- [JavaScript ES6 (ES2015) 移行に関する覚書 - Qiita](http://qiita.com/LightSpeedC/items/9dd46c456e7bbdb1c857)
- [Node v4.2.2 (LTS) | Node.js](https://nodejs.org/en/blog/release/v4.2.2/)

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%82%8F%E3%82%8C%E3%81%AF%E3%83%AD%E3%83%9C%E3%83%83%E3%83%88%E3%80%94%E6%B1%BA%E5%AE%9A%E7%89%88%E3%80%95-%E3%82%A2%E3%82%A4%E3%82%B6%E3%83%83%E3%82%AF-%E3%82%A2%E3%82%B7%E3%83%A2%E3%83%95-ebook/dp/B00O1VK072?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00O1VK072"><img src="https://images-fe.ssl-images-amazon.com/images/I/51UzGYXJ70L._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%82%8F%E3%82%8C%E3%81%AF%E3%83%AD%E3%83%9C%E3%83%83%E3%83%88%E3%80%94%E6%B1%BA%E5%AE%9A%E7%89%88%E3%80%95-%E3%82%A2%E3%82%A4%E3%82%B6%E3%83%83%E3%82%AF-%E3%82%A2%E3%82%B7%E3%83%A2%E3%83%95-ebook/dp/B00O1VK072?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00O1VK072">われはロボット〔決定版〕</a></dt>
	<dd>アイザック アシモフ</dd>
	<dd>小尾芙佐 (翻訳)</dd>
    <dd>早川書房 2014-04-25 (Release 2014-09-30)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00O1VK072</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">ロボットや AI の SF ならこれが古典で定番か？ 面白かったら続けて『<a href="https://www.amazon.co.jp/%E9%8B%BC%E9%89%84%E9%83%BD%E5%B8%82-%E3%82%A2%E3%82%A4%E3%82%B6%E3%83%83%E3%82%AF-%E3%82%A2%E3%82%B7%E3%83%A2%E3%83%95-ebook/dp/B00O2O7JFY?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00O2O7JFY">鋼鉄都市</a>』も読むとよい。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-11-03">2015-11-03</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>
