+++
title = "日本版「グレート・ファイアウォール」に関するブックマーク集"
date = "2018-04-25T16:12:41+09:00"
update = "2018-06-03T17:32:14+09:00"
description = "情報収集しつつ個人でどこまで武装できるか，これから少しずつ考えていくことにする。"
image = "/images/attention/kitten.jpg"
tags = [ "censorship", "internet", "code", "law", "politics" ]

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

[2018-04-17 頃に閉鎖したとみられる「漫画村」](https://www.huffingtonpost.jp/2018/04/17/mangamura-server_a_23413077/ "「漫画村」運営側が自ら閉鎖か　サーバーに接続できず")を NTT グループがブロッキングするという茶番を見るに至り，日本という国も本格的にインターネットに敵対することが明白となった。

- [NTT、海賊版サイトへのブロッキング実施を発表 - ZDNet Japan](https://japan.zdnet.com/article/35118199/)
- [ドコモが海賊版サイトのブロッキングを正式発表--ソフトバンクとKDDIは？ - CNET Japan](https://japan.cnet.com/article/35118202/)

とはいえ個人にできることは少ない。
そこで情報収集しつつ個人でどこまで武装できるか，これから少しずつ考えていくことにする。
大雑把に以下のような感じかなぁ。

1. 「自由」な DNS サービスを利用する
2. Tor や VPN といった仮想ネットワークを用い ISP レベルのフィルタリングやブロッキング，あるいはリージョン・コントロールといったものを迂回する
3. なるべく ISP を使わないメッシュ・ネットワーク構築の可能性を探る

なお，この記事は随時・加筆修正していく。

{{% div-box %}}
あらかじめ言っておくが，私は海賊版サイトを擁護するつもりはひと欠片もない。
ただし，[海賊版サイトをゼロにすることは不可能](http://p2ptk.org/copyright/715 "日本では知られていない海賊版の新潮流 – P2Pとかその辺のお話R")であることも知っている。
これはインセンティブの問題なのだ。
今回の問題に限らずリスク感覚に乏しい日本人が無駄なリソースを消費する場面をこれまで何度も見てきた。
できないことをしようとして見当違いの方法を執る（そしてそれを利権に換える）バカさ加減に多くの人が気づくことを祈っている。
{{% /div-box %}}

## DNS サービス関連

- [ネット利用履歴をISPに残さない新ツール「1.1.1.1」--接続も高速化 - CNET Japan](https://japan.cnet.com/article/35117022/)
- [Cloudflare、セキュアで最速な一般向けDNSサービス「1.1.1.1」提供開始 - ITmedia NEWS](http://www.itmedia.co.jp/news/articles/1804/02/news074.html)
- [Cloudflareが1.1.1.1で超高性能DNS始めたし、いっちょ俺のパソコンもDNS over HTTPSしてみる - Qiita](https://qiita.com/onokatio/items/42fb4a2811600680591b)
- [Cloudflareの1.1.1.1が「DNS over Twitter」を開始。ドメイン名をツイートすると即座にIPアドレスを返答 － Publickey](http://www.publickey1.jp/blog/18/cloudflare1111dns_over_twitterip.html)
- [ラズパイとBINDでDNSブロッキングを回避する | 日経 xTECH（クロステック）](http://tech.nikkeibp.co.jp/atcl/nxt/column/18/00079/00012/)

## ネットワーク関連

もともと VPN は公衆無線 LAN のセキュリティ不備や [WPA2 脆弱性]({{< ref "/remark/2017/10/breaking-wpa2-by-forcing-nonce-reuse.md" >}} "WPA2 脆弱性（KRACKs）に関する覚え書き")の回避手段として注目されているが， VPN 内部や出入口ポイントの安全性を確認するのが（一般の人には）難しいという欠点がある。
これについては [TorrentFreak] が時々レビュー記事を載せている。

- [Which VPN Services Keep You Anonymous in 2018? - TorrentFreak](https://torrentfreak.com/vpn-services-keep-anonymous-2018/)

他にも [TorrentFreak] では海賊版サイトの記事やサイト・ブロッキングなどについても多くの記事があるので参考になるだろう。
もちろん日本の話も載っている。

- [Japan ISP Says it Will Voluntarily Block Pirate Sites as Major Portal Disappears - TorrentFreak](https://torrentfreak.com/japan-isp-says-it-will-voluntarily-block-pirate-sites-as-major-portal-disappears-180424/)

個人的には携帯端末では [F-Secure] 社の [FREEDOME] を使っている。
他には [ProtonVPN] も注目しているのだが，まだ評価できていない。

[TorrentFreak]: https://torrentfreak.com/ "TorrentFreak ⋆ Breaking File-sharing, Copyright and Privacy News"
[F-Secure]: https://www.f-secure.com/ "F-Secure | Cyber Security Solutions for your Home and Business"
[FREEDOME]: https://www.f-secure.com/en/web/home_global/freedome "F-Secure FREEDOME VPN — Protect your privacy | F-Secure"
[ProtonVPN]: https://protonvpn.com/ "ProtonVPN: Secure and Free VPN service for protecting your privacy"

### ネットワーク関連ブックマーク

- [Tor Project | Privacy Online](https://www.torproject.org/)
    - [Orbot: Tor for Android – Guardian Project](https://guardianproject.info/apps/orbot/)
    - [Tor Browser](https://www.torproject.org/projects/torbrowser.html.en)

- [“The Shadow Web” （再掲載）]({{< ref "/remark/2016/11/the-shadow-web.md" >}})

## その他ブックマーク

- [マンガ・アニメの海賊版サイト、ブロッキング含め検討＝官房長官 | ロイター](https://jp.reuters.com/article/suga-pirated-copy-idJPKBN1GV0PV)
- [政府による海賊版サイトへのブロッキング要請に反対する緊急声明 - MIAU](https://miau.jp/ja/845)
- [ウェブは危機に瀕している。我々とともにウェブのために戦おう。（The web is under threat. Join us and fight for it. 日本語訳）](https://www.yamdas.org/column/technique/web-birthday-29j.html)
- [発明者が語る、ウェブの三つの課題（Three challenges for the web, according to its inventor 日本語訳）](https://www.yamdas.org/column/technique/web-turns-28-letterj.html)
- [「海賊版サイト」ブロッキングで白熱議論　「抽象論だと結論出ない」「いや、避けて通れない」 - 弁護士ドットコム](https://www.bengo4.com/internet/n_7763/)
- [「漫画村」などの海賊版サイトを潰すために出版業界が行ってきたことは？ - 鷹野凌のデジタル出版最前線 - 窓の杜](https://forest.watch.impress.co.jp/docs/bookwatch/digipub/1118474.html)
- [NTTによるブロッキングの何が許せないのか - Software Transactional Memo](http://kumagi.hatenablog.com/entry/why-ntt-blocking)
- [海賊版サイトブロッキングに関する質問趣意書に「内容のない回答」をする日本政府 – P2Pとかその辺のお話R](http://p2ptk.org/copyright/887)
- [出版業界はブロッキング問題で岐路に立っている «  マガジン航[kɔː]](https://magazine-k.jp/2018/05/01/editors-note-32/)
- [ディズニー、ネットセーフティ、プログラミング教育まで遮断してしまう英国ISPのセーフティ・フィルター – P2Pとかその辺のお話R](http://p2ptk.org/freedom-of-speech/904)
- [デンマーク、ブロッキング実施後も海賊版サイトへのアクセスが67％増加 – P2Pとかその辺のお話R](http://p2ptk.org/copyright/900)
- [最も「安全」なメッセンジャーアプリ「Signal」がGoogleに続いてAmazonからもBANされ一部の国で利用不可能な状態になる - GIGAZINE](https://gigazine.net/news/20180507-amazon-shut-down-signal-account/)
- [「ネット中立性」規則、6月に撤廃へ--通信会社の自由が拡大 - CNET Japan](https://japan.cnet.com/article/35119001/)
- ['Anonymous' Hackers Deface Russian Govt. Site to Protest Web-Blocking (NSFW) - TorrentFreak](https://torrentfreak.com/anonymous-hackers-deface-russian-govt-site-to-protest-web-blocking-nsfw-180512/)
- [ネットの自由を守り、無法地帯にしないために――NTTグループがブロッキングを決断した理由 - ITmedia Mobile](http://www.itmedia.co.jp/mobile/articles/1805/11/news136.html)
- [ISP Telenor Will Block The Pirate Bay in Sweden Without a Shot Fired - TorrentFreak](https://torrentfreak.com/isp-telenor-will-block-the-pirate-bay-in-sweden-without-a-shot-fired-180520/)
- [カナダ権利者団体、裁判によるサイトブロッキングは役に立たないとして民間サイトブロッキングの実施を求める – P2Pとかその辺のお話R](http://p2ptk.org/copyright/951)
- [Spotifyがヘイトスピーチを含む楽曲の検閲ポリシーを撤回 | All Digital Music](http://jaykogami.com/2018/06/15253.html)

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B01CYDGUV8/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/31Q2jh%2B5SgL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B01CYDGUV8/baldandersinf-22/">CODE VERSION 2.0</a></dt><dd>ローレンス・レッシグ 山形浩生 </dd><dd>翔泳社 2007-12-19</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01HPIZ24I/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B01HPIZ24I.09._SCTHUMBZZZ_.jpg"  alt="コモンズ～ネット上の所有権強化は技術革新を殺す"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01DJ5VE0W/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B01DJ5VE0W.09._SCTHUMBZZZ_.jpg"  alt="FREE CULTURE"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01DIV9AI0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B01DIV9AI0.09._SCTHUMBZZZ_.jpg"  alt="REMIX ハイブリッド経済で栄える文化と商業のあり方"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B06Y5KFBMM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B06Y5KFBMM.09._SCTHUMBZZZ_.jpg"  alt="勉強の哲学　来たるべきバカのために (文春e-book)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01N4MYLFN/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B01N4MYLFN.09._SCTHUMBZZZ_.jpg"  alt="あたらしい人工知能の教科書 プロダクト／サービス開発に必要な基礎知識"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01FH3KRTI/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B01FH3KRTI.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01MU9VUTA/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B01MU9VUTA.09._SCTHUMBZZZ_.jpg"  alt="ネットメディア覇権戦争～偽ニュースはなぜ生まれたか～ (光文社新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00ETNHZJS/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00ETNHZJS.09._SCTHUMBZZZ_.jpg"  alt="ヴェニスの商人の資本論 (ちくま学芸文庫)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01LWTQFDN/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B01LWTQFDN.09._SCTHUMBZZZ_.jpg"  alt="プログラマのためのSQLグラフ原論 リレーショナルデータベースで木と階層構造を扱うために"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01J7Q5LB0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B01J7Q5LB0.09._SCTHUMBZZZ_.jpg"  alt="弱いつながり　検索ワードを探す旅 (幻冬舎文庫)"  /></a> </p>
<p class="description">前著『CODE』改訂版。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-04-07">2017-04-07</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
