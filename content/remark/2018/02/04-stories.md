+++
title = "週末スペシャル： 2019年の暦象，他"
date = "2018-02-04T02:44:22+09:00"
update = "2018-12-12T15:48:49+09:00"
description = "2019年の暦象 / Coincheck のアレ / インフルエンザ過敏症とトリアージ / 「公衆無線LANセキュリティ分科会報告書（案）に対する意見募集」"
image = "/images/attention/kitten.jpg"
tags        = [ "astronomy", "ephemeris", "security", "medical", "risk", "management", "wireless", "fintech", "market" ]

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

久しぶり，週末の戯れ言特集。
今回は Facebook の TL に書き散らしたものを中心にお送りします。

1. [2019年の暦象]({{< relref "#calendar" >}})
1. [Coincheck のアレ]({{< relref "#coincheck" >}})
1. [インフルエンザ過敏症とトリアージ]({{< relref "#triage" >}})
1. [「公衆無線LANセキュリティ分科会報告書（案）に対する意見募集」]({{< relref "#wireless" >}})

## 2019年の暦象{#calendar}

今年も予定通り国立天文台から来年（2019年）の暦象が公開された。

- [平成31（2019）年暦要項の発表 | 国立天文台(NAOJ)](https://www.nao.ac.jp/news/topics/2018/20180201-rekiyoko.html)

祝日から解説すると，12月23日の天皇誕生日がなくなる。
これは新しい天皇が即位されるのが2019年5月で，その時点から「天皇誕生日」が2月23日に変わるため2019年中は天皇誕生日はないということらしい。
また即位による祝日が5月1日になった場合は前後の平日も祝日になるようだ。

- [天皇陛下の退位に伴い2019年は祝日「天皇誕生日」がなくなる　翌年から2月23日が新たな祝日に - ねとらぼ](http://nlab.itmedia.co.jp/nl/articles/1802/02/news108.html)

暦象については2019年は水星の日面経過がある。
ただし日本では見られない。
また金環日食もあるが日本では部分日食になる。

## Coincheck のアレ{#coincheck}

- [Coincheckで発生した暗号通貨NEMの不正送金事案についてまとめてみた - piyolog](http://d.hatena.ne.jp/Kango/20180126/1517012654)
- [コインチェックからのNEM流出、なぜ安全対策が遅れたのか(楠正憲) - 個人 - Yahoo!ニュース](https://news.yahoo.co.jp/byline/kusunokimasanori/20180128-00080965/)
    - [コインチェック事件から1週間、よく聞かれた疑問と今後の論点(楠正憲) - 個人 - Yahoo!ニュース](https://news.yahoo.co.jp/byline/kusunokimasanori/20180204-00081229/)
- [金融庁がコインチェックへの立入検査、CAMPFIREなどみなし仮想通貨交換業者15社にも報告徴求命令  |  TechCrunch Japan](http://jp.techcrunch.com/2018/02/02/fsa-coincheck/)

一応，重大なセキュリティ・インシデント（かつ現在進行形）なんで挙げておくけど，個人的には[「仮想通貨」]の投機には小指の先ほども興味が無いので完全に他人事である。
だいたい，あんな短期間に価値が乱高下する通貨なんか取引に使えないだろ。
もはやこれを「通貨（currency）」と呼ぶことさえ疑問に感じてしまう。
もう「チューリップ」でいいぢゃん。

以上，終わり。

[「仮想通貨」]: {{< ref "/remark/2018/01/cryptocurrency-are-not-crypto.md" >}} "「暗号通貨」ってゆーな！"

## インフルエンザ過敏症とトリアージ{#triage}

- [Vol.021 インフルエンザで「早めの受診」は間違いです！  |  MRIC by 医療ガバナンス学会](http://medg.jp/mt/?p=8111)
    - [インフルエンザで「早めの受診」は間違いです！：医療ガバナンス学会 | MRICの部屋 | 新潮社　Foresight(フォーサイト) | 会員制国際情報サイト](http://www.fsight.jp/articles/-/43278)
- [インフルエンザの予防と治療について考える - 18 til i die (another phase)](http://k3c.hatenablog.com/entry/2018/02/01/230354)

これね、結局リスク・マネジメントの問題なのよ。

インフルエンザに対してリスクの高い（重症化しやすい）人はいる。
例えば循環器系にリスクがある人，子供や老人，等々[^inf1]。
あるいは受験等の重要なイベントを控えている人などは出来るだけ避けたいと思うだろう。
そういった人（あるいはその家族）がリスクを下げるために予防接種や罹患時に早めの措置を行なうことは意味がある。

[^inf1]: それ以外では劇症型または劇症型が予測されるウイルスの場合とか。これは「新型インフルエンザ」のときに大騒ぎしたのでみんな覚えてると思う。でも劇症型のインフルエンザなんてそうそうないのよ。

でもリスクというのは系（system）全体で最小になるように再配分（trade-off）しなければならない。
医療ってのは一種の公共サービスなんだから個人の思惑が通らないこともある。
リソースは有限なのだから。

で，（急変など）予測を超える事態が起きた場合に備えて（救急車の手配や医師やベッドの確保など）医療機関側の体制を整えておかなきゃいけないのに，実際には医療機関側のリソースは軽症者に使い潰されているわけ。
そこが問題で，そのことを（正しくリスクを再配分すべき）国が全く分かってないというのが批判の核心なのである。

もうね，インフルエンザだからって診断証明を要求するのは止めなよ。
日本の悪しき慣習。
“[This Is Japan](https://www.amazon.co.jp/exec/obidos/ASIN/B01LYTKUPM/baldandersinf-22/ "Amazon.co.jp： THIS IS JAPAN 英国保育士が見た日本 eBook: ブレイディみかこ: Kindleストア")” って奴。
病院をバイオ・ハザードにするつもりか。

そろそろ日本人は「インフルエンザ過敏症」から卒業して正しいリスク感覚を身に着けるべきなんじゃないのか。
何度おなじ失敗を繰り返したら懲りるんだ。

まぁあとは「てうが」ね。

## 「公衆無線LANセキュリティ分科会報告書（案）に対する意見募集」{#wireless}

- [総務省｜公衆無線LANセキュリティ分科会報告書（案）に対する意見募集](http://www.soumu.go.jp/menu_news/s-news/01ryutsu03_02000137.html)

途中まで読んで「公衆無線 LAN 版安全・安心マーク」という語句を見た瞬間に読む気が失せた。
これが日本のネットワーク・セキュリティの現状である。

海外から来られる皆さん，日本政府はあてにならないのでぜひ自衛してください。

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/UNDERGROUND-MARKET-%E3%83%92%E3%82%B9%E3%83%86%E3%83%AA%E3%82%A2%E3%83%B3%E3%83%BB%E3%82%B1%E3%83%BC%E3%82%B9-%E8%97%A4%E4%BA%95%E5%A4%AA%E6%B4%8B-ebook/dp/B00FONW2V8?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00FONW2V8"><img src="https://images-fe.ssl-images-amazon.com/images/I/51AT2LqRIsL._SL160_.jpg" width="116" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/UNDERGROUND-MARKET-%E3%83%92%E3%82%B9%E3%83%86%E3%83%AA%E3%82%A2%E3%83%B3%E3%83%BB%E3%82%B1%E3%83%BC%E3%82%B9-%E8%97%A4%E4%BA%95%E5%A4%AA%E6%B4%8B-ebook/dp/B00FONW2V8?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00FONW2V8">UNDERGROUND MARKET　ヒステリアン・ケース</a></dt>
	<dd>藤井太洋</dd>
    <dd>朝日新聞出版 2013-11-07 (Release 2013-10-25)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00FONW2V8</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">日本で「仮想通貨」が流行る前に登場した傑作。つかエンジニアは全員「UNDERGROUND MARKET」シリーズを読め！</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-02-04">2018-02-04</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>
