+++
date = "2016-04-17T15:38:38+09:00"
update = "2016-04-28T21:52:35+09:00"
description = "2016年熊本地震 / 三半規管に異常あり / ATOM 1.7 と Visual Studio Code 1.0 が登場 / Windows 版 QuickTime がようやく終了 / その他の気になる記事"
draft = false
tags = ["environment", "astronomy", "telescope", "jaxa", "astro-h", "atom", "vscode", "editor", "security", "vulnerability"]
title = "週末スペシャル： 2016年熊本地震"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

疲れが溜まってるのか，ここのところ自宅帰って飯食ったらそのまま寝落ちしていることが多い。
妙に怒りっぽくなってるし（路面電車で隣に人が座っただけでイラッとする）気をつけないとなぁ。

1. [2016年熊本地震]({{< relref "#eq" >}})
1. [三半規管に異常あり]({{< relref "#astro" >}})
1. [ATOM 1.7 と Visual Studio Code 1.0 が登場]({{< relref "#edit" >}})
1. [Windows 版 QuickTime がようやく終了]({{< relref "#qt" >}})
1. [その他の気になる記事]({{< relref "#other" >}})

## 2016年熊本地震{#eq}

まずは被災された方々にはお見舞い申し上げます。
地震関連 tweets をチェックしてたけど直下型のせいか M4, M5 クラスの余震でも震度5とか結構あって，これはちょっとキツいよね。
特に夜中の地震はトラウマで不眠症になったりするのでホンマに大変だと思います。
私自身は今は自分のことで手一杯なので何もできないのですが，まずは事態が落ち着いてくれることを祈ります。

大きな災害なので流石にテレビ報道も見たんだけど，やっぱテレビはクズだね。
ヘリで上空から中継してる映像とかあるんだけど，あれ地上から「ウザい」と思われてるの知っててやってるのかね。
様々な映像を見るだに「日本のマスコミはクソ」としか言いようがない。

でも国土地理院の [UAV (Unmanned Aerial Vehicle)](http://psgsv2.gsi.go.jp/koukyou/public/uav/) で空撮した映像はいいね。

- [平成２８年熊本地震に関する情報｜国土地理院](http://www.gsi.go.jp/BOUSAI/H27-kumamoto-earthquake-index.html)

{{< fig-youtube id="DXTAAvVB2M8" width="500" height="281" title="【国土地理院】　南阿蘇村河陽周辺の断層 - YouTube" >}}

人が容易に入れないようなところでも drone なら入れるし，軽い荷物なら配送できる drone とかもあると聞くので，もっとこういう技術を活用したらいいと思うよ。

危機状態に求められるリーダーと平時のリーダーとでは要件が異なる。
「政治判断」しかできない官僚・政治家や官僚的体質の企業・組織は危機状態に機能しないどころか障害になる。
彼らをバイパスして必要な情報や判断を必要な人に行き渡らせ有機的に連携できるか。
Crisis Management って結局はそういうことなんだよね。

小川一水さんの『復活の地』でも読んで勉強しなはれ（笑）

{{< fig-gen >}}
<blockquote class="twitter-tweet" data-lang="ja"><p lang="ja" dir="ltr">【拡散希望】『阪神大震災被災者からのお願い』　嬉しいんです！嬉しいんですけどその援助、もうちょっとだけ待って下さい！ <a href="https://t.co/2BgwW00mRx">pic.twitter.com/2BgwW00mRx</a></p>&mdash; 澤田 圭 ｷｬﾗｸﾀｰﾃﾞｻﾞｲﾅｰ (@keisawada) <a href="https://twitter.com/keisawada/status/721306607389253632">2016年4月16日</a></blockquote>
{{< /fig-gen >}}

### 関連ブックマーク（随時追記予定）

- [平成28年（2016年）熊本地震（M6.5）の地学的背景と布田川断層帯・日奈久断層帯について](http://www.eri.u-tokyo.ac.jp/?page_id=183&id=12595)
- {{< pdf-file title="2016年4月17日地震調査委員会評価文（熊本地方地震）" link="http://www.mext.go.jp/component/a_menu/other/detail/__icsFiles/afieldfile/2016/04/17/1369832_002.pdf" >}}
- {{< pdf-file title="平成28 年4 月17日地震調査委員会図表（熊本地方地震）" link="http://www.mext.go.jp/component/a_menu/other/detail/__icsFiles/afieldfile/2016/04/17/1369832_003.pdf" >}}
- [避難所、給水、配給、空いているスーパーなどについてgooglemapまとめ | 熊本地震 情報掲示板](http://kumamoto-jishin.info/map.html)
- [熊本地震で観測した電子基準点の変位を、国土地理院がいち早く公表｜ニュース/アーカイブ｜準天頂衛星システム（QZSS）公式サイト - 内閣府](http://qzss.go.jp/news/archive/gsi_kumamoto_160418.html)
- [広島大学　熊本県熊本地方を震源とする地震に関する記者説明会を開催しました](http://www.hiroshima-u.ac.jp/news/show/id/25510)
- [2016年4月15日ニュース「熊本の大地震は横ずれ断層型 震源浅く揺れ大きく」 | SciencePortal](http://scienceportal.jst.go.jp/news/newsflash_review/newsflash/2016/04/20160415_01.html)
- [2016年4月18日ニュース「本震と余震は別の断層帯で発生 地震調査委と気象庁が見解」 | SciencePortal](http://scienceportal.jst.go.jp/news/newsflash_review/newsflash/2016/04/20160418_01.html)
- [2016年4月18日ニュース「被害状況を動画で公開 国土地理院がドローンで調査」 | SciencePortal](http://scienceportal.jst.go.jp/news/newsflash_review/newsflash/2016/04/20160418_02.html)
- [平成28年熊本地震に伴う被災地救援等のために使用する車両の取り扱いについて / 熊本県](http://www.pref.kumamoto.jp/kiji_15425.html?type=top)
- [オープンデータで避難所地図を表示する「ヒナンパス」が、熊本地震に暫定対応｜利用者向け情報｜準天頂衛星システム（QZSS）公式サイト - 内閣府](http://qzss.go.jp/usage/userreport/hinanpass_160420.html)
- [熊本地震で、ITSJapanが乗用車・トラックの通行実績情報を提供｜ニュース/アーカイブ｜準天頂衛星システム（QZSS）公式サイト - 内閣府](http://qzss.go.jp/news/archive/its-jp_160419.html)
- [熊本地震への対応について | OpenStreetMap Japan](https://openstreetmap.jp/node/779)
- [災害ボランティアの受入について / 熊本市ホームページ](http://www.city.kumamoto.jp/hpkiji/pub/detail.aspx?c_id=5&type=top&id=12541)

## 三半規管に異常あり{#astro}

- {{< pdf-file title="Ｘ線天文衛星「ひとみ」（ASTRO-H）の状況について" link="http://fanfun.jaxa.jp/jaxatv/files/20160415_hitomi.pdf" >}}
- [X線天文衛星「ひとみ」、2重のトラブルで「自分で回った」と推定 | Sorae.jp : 宇宙（そら）へのポータルサイト](http://sorae.jp/space/2016_04_15_asrtoh.html)
- [JAXA | X線天文衛星「ひとみ」（ASTRO-H）の状況について](http://www.jaxa.jp/press/2016/04/20160419_hitomi_j.html)
- [2016年4月20日ニュース「姿勢制御装置が誤作動 通信途絶えた天文衛星『ひとみ』」 | SciencePortal](http://scienceportal.jst.go.jp/news/newsflash_review/newsflash/2016/04/20160420_01.html)
- [X線天文衛星「ひとみ」、浮かび上がった3つの問題点 | sorae.jp : 宇宙（そら）へのポータルサイト](http://sorae.jp/02/2016_04_26_astroh.html)

小天体かデブリにでもぶち当たったのかと思ったが，擬人化するなら三半規管の異常で「ピヨった」状態らしい。
太陽電池パネルの一部が千切れるほどの高速回転って相当だと思うけど。

JAXA は通信が回復する可能性を期待して受信体制を続けるようだが，これは難しいんじゃないかなぁ。
宇宙じゃ「ちょいと行ってきて直す」とかできないからねぇ[^ht]。

[^ht]: ハッブル望遠鏡の時はスペースシャトルが運用されていた時代なので，望遠鏡のところまで行って修理したけど。

## ATOM 1.7 と Visual Studio Code 1.0 が登場{#edit}

- [GitHub、オープンソースのテキストエディター「Atom」の最新正式版v1.7を公開 - 窓の杜](http://www.forest.impress.co.jp/docs/news/20160414_753303.html)
- [Microsoft製の無償コードエディター「Visual Studio Code」がv1.0.0に - 窓の杜](http://www.forest.impress.co.jp/docs/news/20160415_753468.html)
- [オープンソースの開発ツール「Visual Studio Code」が正式版となるバージョン1.0に到達 － Publickey](http://www.publickey1.jp/blog/16/visual_studio_code10.html)
- [1.0になったVisualStudioCodeの強みと弱点 - Qiita](http://qiita.com/74th/items/12521790a1c680af934c)

私は Windows 版の [ATOM] を使ってるんだけど， tree view からのファイル削除をしくじるんだよなぁ（[ATOM] というより [ATOM] にインストールしているパッケージのせいかもしれんけど）。
多分あれって Windows のファイルシステムのせいだよなぁ。

- [Big Sky :: Windows ユーザは cmd.exe で生きるべき。](http://mattn.kaoriya.net/software/why-i-use-cmd-on-windows.htm)

[Visual Studio Code] はいい製品に育ったねぇ。
私は既に [ATOM] をメインに使ってるけど，これは「ちょっと使ってみようかな」と思わせる製品になった。

最近の Microsoft は .NET Core と [Visual Studio Code] がいい感じである。
逆に Windows 10 は過剰なコントロールでどんどんダメな OS になっていく。
もう Microsoft も Windows は事実上見捨ててるのだろう。

## Windows 版 QuickTime がようやく終了{#qt}

- [Apple Ends Support for QuickTime for Windows; New Vulnerabilities Announced](https://www.us-cert.gov/ncas/alerts/TA16-105A)
- [JVNTA#92371676: QuickTime for Windows に複数のヒープバッファオーバフローの脆弱性](http://jvn.jp/ta/JVNTA92371676/)

Mac 版とかはどうか知らないが， Windows 版の QuickTime は最初からクソッタレな実装だった。
しかも当時は iTunes と抱き合わせでインストールされるので容易に捨てられず往生した覚えがある。

CVSSv3 基本値は 6.3 なので「要注意」レベルだが，サポートされない製品を入れておく理由はない。
危険物はとっとと捨てましょう。

## その他の気になる記事{#other}

- [Linking to Pirated Content Is Not Copyright Infringement, Says EU Court Adviser - TorrentFreak](https://torrentfreak.com/linking-to-pirated-content-is-not-copyright-infringement-160407/)
- [カナダ警察、「BlackBerry」のマスター暗号化キーを入手していた--メッセージの解読に利用 - CNET Japan](http://japan.cnet.com/news/service/35081244/)
- [White House Source Code Policy a Big Win for Open Government | Electronic Frontier Foundation](https://www.eff.org/deeplinks/2016/04/white-house-source-code-policy-big-win-open-government)
- [【注意喚起】ランサムウェア感染を狙った攻撃に注意：IPA 独立行政法人 情報処理推進機構](http://www.ipa.go.jp/security/topics/alert280413.html)
- [Working on HTML5.1 | W3C Blog](https://www.w3.org/blog/2016/04/working-on-html5-1/)
    - [w3c/html: Working Draft of the HTML specification](https://github.com/w3c/html)
    - [W3C、「HTML5.1」を今年9月に勧告とする計画。仕様はGitHubで公開 － Publickey](http://www.publickey1.jp/blog/16/w3chtml519github.html)
- [中国は「オーウェル風ディストピア」？「社会信用制度」とは · Global Voices 日本語](https://jp.globalvoices.org/2016/04/13/40418/)
- [At Japanese Beatmaking Event, Producers Create CC Remixes in Just Four Hours - Creative Commons blog - Creative Commons](https://blog.creativecommons.org/2016/04/14/japanese-beatmaking-event-producers-create-cc-remixes-just-four-hours/) : 日本のイベントに本家 CC が反応していることに驚いた。昨年のソウルのイベントでは CCjp はほぼ無視してたのにね
- [[Announce] Libgcrypt 1.7.0 released](https://lists.gnupg.org/pipermail/gnupg-announce/2016q2/000386.html) : SHA-3 や ChaCha など新しいアルゴリズムが登場

[ATOM]: https://atom.io/ "Atom"
[Visual Studio Code]: https://code.visualstudio.com/ "Visual Studio Code - Code Editing. Redefined"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B00GJOESS6/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51ymtvyHUmL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B00GJOESS6/baldandersinf-22/">復活の地１</a></dt><dd>小川一水 </dd><dd>早川書房 2012-09-15</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00GJMUKEY/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00GJMUKEY.09._SCTHUMBZZZ_.jpg"  alt="復活の地２"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00GJMUKG2/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00GJMUKG2.09._SCTHUMBZZZ_.jpg"  alt="復活の地３"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00GJMUKDK/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00GJMUKDK.09._SCTHUMBZZZ_.jpg"  alt="第六大陸２"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00GJMUKDU/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00GJMUKDU.09._SCTHUMBZZZ_.jpg"  alt="第六大陸1"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00GJMUKYO/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00GJMUKYO.09._SCTHUMBZZZ_.jpg"  alt="天涯の砦"  /></a> </p>
<p class="description">コミカライズ版もある。てか，コミカライズ版を最初に読んだ（笑） 大きな災害がある度にこの作品を思い出す。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-04-17">2016-04-17</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
