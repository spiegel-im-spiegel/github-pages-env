+++
title = "プラットフォームかプロトコルか"
date =  "2024-04-15T21:07:04+09:00"
description = "日本で2日間（東京と大阪）に渡り Bluesky Meetup が開催された"
image = "/images/attention/kitten.jpg"
tags = [ "internet", "engineering", "communication", "bluesky", "mastodon", "activitypub" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

物凄い偏見であることは承知の上で書いてしまうけど，現代を含む近代というのは「『個人』が台頭する時代」だと思うのよ。
もしくは社会と個人が向き合う時代。
時に協力し時に反目するのが社会と個人との関係。
その文脈の中で，たとえば身分や差別の問題，たとえばジェンダーの問題，たとえば税金と福祉の問題，といったことを捉える必要がある。

ゼロ年代前半の「Web 2.0」で期待されたのは多様な Web サービスの間で個人が緩やかに連携しボトムアップで「ネット社会」を構成することだった筈だ。
しかしそれは「ソーシャルネットワーク・サービス」の登場でぶち壊された。
ユーザは人間関係を「人質」にプラットフォームに囲い込まれ縛り付けられる。
更に2010年代の「アラブの春」がこれを強化した。
日本ではこれに 3.11 を加えてもいいかも知れない。
ネットはプラットフォーム間の覇権ゲームの会場となりユーザを[過剰に飲み込み嘔吐する]({{< ref "/remark/2023/01/internet-bulimia.md" >}} "嘔吐するインターネット")。
そこに個人は存在せず，ユーザは統計上の「数値」でしかない。
この辺が「[デジタル封建主義](https://yamdas.hatenablog.com/entry/20230828/neo-feudalism "時代はデジタル封建主義？ ジョエル・コトキン『新しい封建制がやってくる』が出るぞ - YAMDAS現更新履歴")」などと呼ばれる所以だろう。

これを打ち破りたいならプロトコルから変えるべきだ！ という発想になるのは自然なことかもしれない。

先週末は日本で2日間（東京と大阪）に渡り Bluesky Meetup が開催された。
個人的には関心が薄かったので見てもないのだが yomoyomo さんがこれに絡む記事を [WirelessWire News](https://wirelesswire.jp/ "WirelessWire News – The Technology and Ecosystem of the IoT.") のブログにタイムリーに公開されていて，流石と思ってしまった。

- [BlueskyやThreadsに受け継がれたネット原住民の叡智 – WirelessWire News](https://wirelesswire.jp/2024/04/86389/)

この中で紹介されている以下の記事は読んでおくべきだろう。

- [言論の自由を取り巻く問題を解決する「プロトコルに基づいた仕組み」とは？ - GIGAZINE](https://gigazine.net/news/20210201-free-speech-protocols-approach/)
- [ウェブの35歳の誕生日を祝う：オープンレター（Marking the Web’s 35th Birthday: An Open Letter 日本語訳）](https://www.yamdas.org/column/technique/marking-the-webs-35th-birthday-an-open-letterj.html)

今のところ Bluesky には Bluesky というサービスと，それを駆動する [AT (Authenticated Transfer) Protocol](https://atproto.com/ "The AT Protocol") という2本の柱がある。
おそらく Bluesky 自身は AT プロトコルを実装する実験的な側面を持っていて，本命は AT プロトコルの標準化とそれを使った「Bluesky ではない」サービス群の台頭ではないかと思う。

上で紹介した一連の記事を読んで思い出したのは『[インテンション・エコノミー](https://www.amazon.co.jp/dp/B00DIM6BE6?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』だ。
この本についても yomoyomo さんの[書評](https://www.yamdas.org/booklog/intentioneconomy.html "yomoyomoの読書記録 - ドク・サールズ『インテンション・エコノミー 顧客が支配する経済』（翔泳社）")がある。
一部引用しておこう。

{{< fig-quote type="markdown" title="yomoyomoの読書記録 - ドク・サールズ『インテンション・エコノミー 顧客が支配する経済』（翔泳社）" link="https://www.yamdas.org/booklog/intentioneconomy.html" >}}
　インテンション・エコノミーとは買い手が価値の源泉となる真の意味でオープンな市場であり、VRM（企業関係管理）が CRM（顧客関係管理）にとってかわり、顧客は企業に囲い込まれることなく、つまり消費者として集合的に扱われるのでなく企業との関係はパーソナルなものになる。

　つまり、顧客の側が自分に関するデータの主導権に握り、自らの意思に従い企業との関係を決められる。それには顧客側に立ちそのニーズの代理人として機能する「フォース・パーティ」の存在が必要になるし、顧客と企業の間には対話的でオープンな API が提供されることで市場のオープンさが担保される。
{{< /fig-quote >}}

これを[読んだ](https://baldanders.info/blog/000638/ "『インテンション・エコノミー』を読む")当時は「企業と個人」の関係で考えていたが，これは「社会と個人」の関係で考えることもできる。
もしかしたらこれから求められるのは SRM (社会関係管理) みたいな仕組みかもしれない，などと妄想してみた。
もし AT プロトコルによって SRM が構成できるなら面白いのに（笑）

既に ActivityPub は W3C によって[標準化][ActivityPub]され Mastodon のみならず PeerTube や Pixelfed といった多様なサービスを生み出しつつある（なくなってるぽいものもあるが）。
日本でならこれに Misskey を加えてもいいかも知れない。
Threads は Mastodon ベースだそうだが，もはや Mastodon とは別物だろう。
こういうことが AT プロトコル周りでも起こるようになれば，もう少し真面目に取り組んでもいいかなぁ，とは思っている。

まぁ，でも，標準化されたらされたで，またぞろ覇権ゲームが勃発しそうな気がするけどね。
そしてその勝者はやがて腐っていく。
これも自然の理か？

{{< fig-quote type="markdown" title="メタクソ化するTiktok" link="https://p2ptk.org/monopoly/4366" >}}
プラットフォームはこのように滅びていく。まず、ユーザにとって良き存在になる。次に、ビジネス顧客にとって良き存在になるために、ユーザを虐げる。最後に、ビジネス顧客を虐げて、すべての価値を自分たちに向ける。そうして死んでいく。
{{< /fig-quote >}}

[ActivityPub]: https://www.w3.org/TR/activitypub/

## ブックマーク

- [WirelessWire News連載更新（BlueskyやThreadsに受け継がれたネット原住民の叡智） - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20240415/wirelesswire)
- [プロトコル・ベースのプラットフォーム規制 | PPT](https://www.slideshare.net/masayukihatta/ss-199545007)
- [なぜミレニアル世代はTiktokを離れられないのか（あるいはおじさん・おばさんがFacebookを離れられないのはなぜか） | p2ptk[.]org](https://p2ptk.org/monopoly/4482)

-  [Fediverse 関連のブックマーク]({{< ref "/bookmarks/fediverse.md" >}})

## 参考

{{% review-paapi "B00DIM6BE6" %}} <!-- インテンション・エコノミー -->
{{% review-paapi "B00FW5SSCK" %}} <!-- ソーシャル・ネットワーク -->
{{% review-tatsujin "infoshare2" %}} <!-- 続・情報共有の未来 -->
