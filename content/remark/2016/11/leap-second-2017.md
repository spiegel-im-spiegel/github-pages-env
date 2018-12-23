+++
date = "2016-11-29T22:06:59+09:00"
update = "2017-01-12T11:11:28+09:00"
title = "2017年直前の閏秒について"
description = "年も押し迫ってきたし関連情報も出てき始めているので，あらためて記事を再構成し関連リンクを順次追加していくことにする。"
tags = [
  "astronomy",
  "leap-second",
  "engineering",
]
draft = false

[author]
  github = "spiegel-im-spiegel"
  tumblr = ""
  facebook = "spiegel.im.spiegel"
  flattr = ""
  avatar = "/images/avatar.jpg"
  linkedin = "spiegelimspiegel"
  instagram = "spiegel_2007"
  flickr = "spiegel"
  name = "Spiegel"
  license = "by-sa"
  url = "http://www.baldanders.info/spiegel/profile/"
  twitter = "spiegel_2007"

[scripts]
  mathjax = true
  mermaidjs = false
+++

[以前に紹介した]({{< ref "/remark/2016/07/10-stories.md#ls" >}})けど，年も押し迫ってきたし関連情報も出てき始めているので，あらためて記事を再構成し関連リンクを順次追加していくことにする。

- [プレスリリース | 「うるう秒」挿入のお知らせ | NICT-情報通信研究機構](http://www.nict.go.jp/press/2016/07/08-1.html)
- [2016年12月31日にうるう秒が追加へ - CNET Japan](http://japan.cnet.com/news/business/35085582/)

UTC の2017年直前に閏秒が挿入される。
日本時間では，時差があるので，2017年1月1日午前9:00直前に閏秒が挿入されることになる。
これにより UTC と TAI の差は37秒となる。


| 実施年月日     | 調整時間（秒） | $UTC - TAI$（秒） |
|:---------------|---------------:|------------------:|
| 1972年1月1日   |             ― |             $-10$ |
| 1972年7月1日   |           $+1$ |             $-11$ |
| 1973年1月1日   |           $+1$ |             $-12$ |
| 1974年1月1日   |           $+1$ |             $-13$ |
| 1975年1月1日   |           $+1$ |             $-14$ |
| 1976年1月1日   |           $+1$ |             $-15$ |
| 1977年1月1日   |           $+1$ |             $-16$ |
| 1978年1月1日   |           $+1$ |             $-17$ |
| 1979年1月1日   |           $+1$ |             $-18$ |
| 1980年1月1日   |           $+1$ |             $-19$ |
| 1981年7月1日   |           $+1$ |             $-20$ |
| 1982年7月1日   |           $+1$ |             $-21$ |
| 1983年7月1日   |           $+1$ |             $-22$ |
| 1985年7月1日   |           $+1$ |             $-23$ |
| 1988年1月1日   |           $+1$ |             $-24$ |
| 1990年1月1日   |           $+1$ |             $-25$ |
| 1991年1月1日   |           $+1$ |             $-26$ |
| 1992年7月1日   |           $+1$ |             $-27$ |
| 1993年7月1日   |           $+1$ |             $-28$ |
| 1994年7月1日   |           $+1$ |             $-29$ |
| 1996年1月1日   |           $+1$ |             $-20$ |
| 1997年7月1日   |           $+1$ |             $-31$ |
| 1999年1月1日   |           $+1$ |             $-32$ |
| 2006年1月1日   |           $+1$ |             $-33$ |
| 2009年1月1日   |           $+1$ |             $-34$ |
| 2012年7月1日   |           $+1$ |             $-35$ |
| 2015年7月1日   |           $+1$ |             $-36$ |
| 2017年1月1日   |           $+1$ |             $-37$ |

## 閏秒について

日常生活で接する時刻系としては UT （universal time; 世界時）と UTC （coordinated universal time; 協定世界時）の2つが存在する。
2つの時刻系は基準となる物差しが異なる。

UT は恒星時（sidereal time）系の一種[^ut]で，簡単に言うと地球の自転速度を基準にしている。
私たちの日常生活は太陽や月などに大きな影響を受けているので， UT を用いるのは妥当と言える。
一方で UT は観測値であり，しかも地球の自転速度は一定ではなく予測できない[^es] のが欠点だった。
そこで「同じ間隔で時を刻む時刻系」の要求が高まってくる。

[^ut]: もう少し詳しく言うと UT は恒星時系の一種である太陽時系である。恒星時は「春分点の子午線からの時角」であるのに対して UT は「平均太陽のグリニジ子午線からの時角＋12時」となる（子午線を基準にすると昼間に日付が変わるため）。両者は同じ物差し（地球の自転速度）で求められた時刻系であり数学的関係があるため互いに換算可能である。ちなみに平均太陽というのは天球上にある太陽の移動速度を均した仮想太陽である。実際の太陽（真太陽）は季節によって天球上での移動速度が異なる。これは地球の公転軌道が楕円になっているからだ。
[^es]: 地球の自転は（潮汐摩擦などにより）大雑把に言って少しずつ遅くなる傾向にある。実はこれが分かったのって20世紀に入ってからなのだよ（問題の認識は19世紀後半からあった）。

かなりの試行錯誤の末，最終的に原子時（atmic time）系が採用されることになった。
具体的には1958年1月1日0時0分0秒 UT2[^ut2] を原点とした TAI[^tai] （international atomic time; 国際原子時）である。

[^ut2]: 観測から得られた時刻を UT0 と呼ぶ。 UT0 に対し極運動等の補正をかけたものを UT1 と呼ぶ。現在 UT と言う場合にはこの UT1 を指す。 UT2 は UT1 から更に自転速度の季節変動分を均した値である。現在の UTC が採用されるまでは UT2 が主に使われていたらしい。
[^tai]: 現在の TAI の定義は「回転するジオイド上で実現される SI の秒を目盛りの単位とした, 地心座標系で定義される座標時（coordinate time）の目盛り」（[訳は Wikipedia より](https://ja.wikipedia.org/wiki/%E5%9B%BD%E9%9A%9B%E5%8E%9F%E5%AD%90%E6%99%82 "国際原子時 - Wikipedia")）となっている。ちなみに「SI の秒」は「セシウム 133 の原子の基底状態の2つの超微細構造準位の間の遷移に対応する放射の周期の9192631770倍の継続時間」（[訳は Wikipedia より](https://ja.wikipedia.org/wiki/%E7%A7%92 "秒 - Wikipedia")）である。この「SI の秒」を使っている時刻系は全て原子時系とも言える。

そして UT と TAI との間のギャップを埋めるのが UTC なのである。

UTC は TAI と同じ原子時系で TAI との差が整数秒になるように調整される。
また UT に対する差は 0.9 秒以内になるように調整される。
この「調整」を行う手段が閏秒（leap second）である。

閏秒は UTC の1月1日直前または7月1日直前に挿入あるいは削除される[^410]。
UT が観測値である以上「いつ閏秒が発生するか」は長期的には予測できないが，少なくとも半年前には[告知](http://jjy.nict.go.jp/QandA/data/leapsec.html "日本標準時プロジェクト Information of Leap second")される。

[^410]: これ以外には4月1日直前および10月1日直前も第2優先日として閏秒が発生する可能性があるが，過去においてこの日に閏秒が発生したことはない。

## コンピュータ・システムにおける閏秒の問題

コンピュータ・システムにおける閏秒の問題は概ね以下の2つに大別できる。

1. 閏秒の挿入による見かけ上の時刻の巻き戻し
2. タイムスタンプの処理

以下，もう少し詳しく解説する。

### 閏秒の挿入による見かけ上の時刻の巻き戻し

現在，多くのコンピュータシステムは [NTP (Network Time Protocol)](https://tools.ietf.org/html/rfc5905 "RFC 5905 - Network Time Protocol Version 4: Protocol and Algorithms Specification") によって時刻同期を行っている。
閏秒が発生する際に [NTP] では LI (Leap Indicator) をセットするのだが，閏秒を挿入する際には見かけ上の時刻の巻き戻しが発生する。

```text
23:59:59.000000（→ 23:59:59.999999）→ 23:59:59.000000 → 00:00:00.000000
```

サービスやアプリケーションのいくつか[^debug] は「時刻の巻き戻し」が考慮されていないため，巻き戻しのタイミングで重大なエラーが発生する可能性がある。

[^debug]: 2012年の閏秒の挿入時には数多くのサービスに障害が発生したが，その後の改善や回避策により2015年には大きな混乱もなく閏秒を迎えられた。なお Linux 系のシステムには `right/Asia/Tokyo` のような閏秒を考慮したタイムゾーンも存在するが，旧来の `Asia/Tokyo` のようなタイムゾーンとの間で UNIX Time の互換性がなくなるのとタイムスタンプ処理と同じ問題が発生してしまうため推奨されないことも多い。

[NTP] には LI をセットせず時刻の巻き戻しを行わない SLEW モードがあり，このモードにすることで問題を回避できる。
ただし SLEW モードは時刻のギャップが起きた際の同期に時間がかかるため，モードの使い分けが必要になるかもしれない。

### タイムスタンプの処理

タイムスタンプは電子署名などでは欠かせない技術要素である（「否認防止（non-repudiation）」には正しい時刻が必要）。
タイムスタンプを提供するサービスでは厳密に UTC に準拠して運営されている。
したがって閏秒の瞬間に「2016年12月31日 23:59:60 UTC」といったタイムスタンプがセットされる可能性がある。

このタイムスタンプを扱う際に処理系によっては誤動作を起こす可能性があるのだ。

- [【殴り書き】javascriptにおけるうるう秒の調査 - Qiita](http://qiita.com/J_3woo86/items/8641b98c14278569ab94)
- [【殴り書き】C#におけるうるう秒調査 - Qiita](http://qiita.com/J_3woo86/items/31ffd000786273fd05e6)

タイムスタンプ・サービスを行っているプロバイダには，閏秒の前後でサービスを止めて「23:59:60 UTC」なタイムスタンプが発生しないようにしているところもあるようだ。

## ブックマーク

### 技術情報

- [事前にご確認ください – AWSにおける2016年12月31日（日本時間2017年1月1日）のうるう秒 | Amazon Web Services ブログ](https://aws.amazon.com/jp/blogs/news/look-before-you-leap-december-31-2016-leap-second-on-aws/)
    - [Amazon AWS、うるう秒の1秒分を前後12時間の1秒を1/86400長くして吸収 -INTERNET Watch](http://internet.watch.impress.co.jp/docs/news/1036937.html)
- [Google Cloud Platform Blog: Making every (leap) second count with our new public NTP servers](https://cloudplatform.googleblog.com/2016/11/making-every-leap-second-count-with-our-new-public-NTP-servers.html)
    - [Google、うるう秒の1秒分を前後20時間のクロック変更で吸収 -INTERNET Watch](http://internet.watch.impress.co.jp/docs/news/1033101.html)
- [来年の元旦に「うるう秒」挿入、システム障害などのトラブル回避に向けて引き続き警戒を -INTERNET Watch](http://internet.watch.impress.co.jp/docs/news/1027910.html)
- [製品使用上の重要なお知らせ](http://www.hitachi-support.com/alert/ss/HWS16-002/index.htm)
- [VMware製品へのうるう秒調整の影響 - Qiita](http://qiita.com/tsukamoto/items/5bbecd29ac40ac16e039) : 最新版に更新されている
    - [NTP のスルー モードの有効化 (2126101) | VMware KB](https://kb.vmware.com/selfservice/microsites/search.do?language=en_US&cmd=displayKC&externalId=2126101)
- [「ステラナビゲータ10」10.0gアップデータ公開、うるう秒対応や不具合修正など - AstroArts](https://www.astroarts.co.jp/article/hl/a/8687)
- [サマータイム、うるう秒とタイムスタンプの関係｜タイムスタンプ入門｜セイコーサイバータイム](https://www.seiko-cybertime.jp/time/column3.html)
- [RFC 5905 - Network Time Protocol Version 4: Protocol and Algorithms Specification](https://tools.ietf.org/html/rfc5905)
- [RFC 3161 - Internet X.509 Public Key Infrastructure Time Stamp Protocol (TSP)](https://tools.ietf.org/html/rfc3161) （[日本語訳](https://www.ipa.go.jp/security/rfc/RFC3161JA.html)）

### その他の四方山話

- [閏秒(うるう秒)とは何か - 国立天文台暦計算室](http://eco.mtk.nao.ac.jp/koyomi/topics/html/topics2013_1.html)
- {{< pdf-file title="暦象年表の改訂について" link="http://www.nao.ac.jp/contents/about-naoj/reports/report-naoj/11-34-2.pdf" >}}
- [暦の改訂（DE405 から DE430 へ） — Baldanders.info](http://www.baldanders.info/spiegel/log2/000840.shtml)
- [Linux開発者リーナス・トーヴァルズ、来る「うるう秒」を語る｜WIRED.jp](http://wired.jp/2015/06/30/torvalds_leapsecond/) ： 2015年の記事
- [うるう秒は当分存続らしい]({{< ref "/remark/2015/leap-second.md" >}})
- [うるう秒で1秒長かった今年の元日｜ニュース/アーカイブ｜準天頂衛星システム（QZSS）公式サイト - 内閣府](http://qzss.go.jp/news/archive/nict_170110.html) : 地球の自転について現在の観測方法を紹介している

昔書いた以下の記事は内容が古くて地球時（terrestrial time; TT）などを含む座標時系への言及がないけど，閏秒に関する歴史的経緯を知るという意味では参考になるかもしれない。

- [時刻系の話： 閏秒ができるまで － 序章 -- 戯れ言++](http://www.baldanders.info/spiegel/remark/archives/000109.shtml)
- [時刻系の話： 閏秒ができるまで － 恒星時系と世界時系 -- 戯れ言++](http://www.baldanders.info/spiegel/remark/archives/000118.shtml)
- [時刻系の話： 閏秒ができるまで － 暦表時系 -- 戯れ言++](http://www.baldanders.info/spiegel/remark/archives/000127.shtml)
- [時刻系の話： 閏秒ができるまで － 原子時系と閏秒 -- 戯れ言++](http://www.baldanders.info/spiegel/remark/archives/000130.shtml)
- [時刻系の話： 閏秒ができるまで － 新しい暦計算システムと力学時 -- 戯れ言++](http://www.baldanders.info/spiegel/remark/archives/000147.shtml)

[NTP]: https://tools.ietf.org/html/rfc5905 "RFC 5905 - Network Time Protocol Version 4: Protocol and Algorithms Specification"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4805202254/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51mQCyP04rL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4805202254/baldandersinf-22/">天体の位置計算</a></dt><dd>長沢 工 </dd><dd>地人書館 1985-09</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4805206349/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4805206349.09._SCTHUMBZZZ_.jpg"  alt="日の出・日の入りの計算―天体の出没時刻の求め方"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4769908180/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4769908180.09._SCTHUMBZZZ_.jpg"  alt="天文計算入門―一球面三角から軌道計算まで"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4805204141/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4805204141.09._SCTHUMBZZZ_.jpg"  alt="パソコンで見る天体の動き"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4416114710/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4416114710.09._SCTHUMBZZZ_.jpg"  alt="天文年鑑2015年版"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00R4X7R0M/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00R4X7R0M.09._SCTHUMBZZZ_.jpg"  alt="月刊 星ナビ 2015年 02月号 [雑誌]"  /></a> </p>
<p class="description">B1950.0 分点から J2000.0 分点への過渡期に書かれた本なので情報が古いものもあるが，基本的な内容は位置天文学の教科書として充分通用する。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-01-11">2015/01/11</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B01JFLCW5K/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51EnYDL31WL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B01JFLCW5K/baldandersinf-22/">猫暦 ねこよみ コミック 1-3巻セット (ねこぱんちコミックス)</a></dt><dd>ねこしみず 美濃 </dd><dd>少年画報社 2016-07-11</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"></p>
<p class="description">「寛政の改暦」のころの伊能勘解由（忠敬）とその妻とされる「おえい」の物語。感想は<a href="http://text.baldanders.info/remark/2016/05/nekoyomi/">こちら</a>。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-11-29">2016-11-29</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
