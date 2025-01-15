+++
title = "ソーシャルメディアは億万長者から逃げ切れるか？"
date =  "2025-01-15T22:41:44+09:00"
description = "タイトルは釣り"
image = "/images/attention/kitten.jpg"
tags = [ "media", "engineering", "activitypub", "atproto", "mastodon", "bluesky", "communication" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

面白い翻訳記事を立て続けに見たので，今回はこれをネタに戯れ言を書いてみる。

- [Free Our Feeds：ソーシャルメディアを億万長者の支配から解放するために力を貸してくれないか » p2ptk[.]org](https://p2ptk.org/freedom-of-speech/5174)
- [「広場」は誰のものでもない、みんなのものであるべきだ » p2ptk[.]org](https://p2ptk.org/freedom-of-speech/5176)

これらの記事を並べて載せる [p2ptk.org] は相変わらず流石だなぁ，と思う。

## AT Protocol を億万長者から守れ！

（原文: [Free Our Feeds – Save Social Media From Billionaire Capture.][Free Our Feeds]）

プラットフォームあるいはアプリケーションとしての Bluesky と基盤技術である AT Protocol は開発体制としては不可分になっている。
今のところはこれでも問題なく運営できているが，プロトコルを私企業に握られている状況は将来に不安があるため AT Protocol を独立させようという動きがあるようだ。
このための資金として，まずは4M米ドルの調達を目指すらしい。

{{< fig-quote type="markdown" title="Free Our Feeds：ソーシャルメディアを億万長者の支配から解放するために力を貸してくれないか" link="https://p2ptk.org/freedom-of-speech/5174" >}}
資金は以下の目的で調達される。

- BlueskyのベースとなるテクノロジーであるAT Protocolを億万長者の支配から守るために活動する公益財団の設立。
- 独立したホスティングインフラ（第2の「リレー」）の構築。これにより、Blueskyのユーザ、開発者、研究者は、今後同社がどのような決定を下そうとも、コンテンツとデータのストリームに完全にアクセスできる。
- 開発者への資金提供。これにより、オープンプロトコル上に多様なソーシャルアプリケーションを構築し、ソーシャルメディアをより健全で幸せな場所にすることができる。

今こそ、ソーシャルメディアを解放する時だ。
{{< /fig-quote >}}

たとえば ActivityPub は2018年早々に [W3C 勧告](https://www.w3.org/TR/2018/REC-activitypub-20180123/ "ActivityPub - W3C Recommendation 23 January 2018")となっている。
AT Protocol も何らかの形で標準化していくのが正しい道筋だと思う。
そのためには実績を作らなくちゃね{{% emoji "はぁと" %}} というところだろう。

Facebook や Instagram を擁する Meta が Pixelfed サーバへのリンクを[ブロックしている](https://gigazine.net/news/20250114-meta-block-link-pixelfed/ "MetaがInstagram風の写真共有SNS「Pixelfed」へのリンクをブロック中 - GIGAZINE")という話も聞くし，マイクロブログだけでなく写真や動画といったものについても AT Protocol ベースのサービスが登場すれば面白いのに，とか思ったりする。

## Mastodon を「所有」から解放する

（原文: [The people should own the town square - Mastodon Blog](https://blog.joinmastodon.org/2025/01/the-people-should-own-the-town-square/)）

もうひとつは Mastodon の話。

{{< fig-quote type="markdown" title="「広場」は誰のものでもない、みんなのものであるべきだ" link="https://p2ptk.org/freedom-of-speech/5176" >}}
端的に言えば、Mastodonのエコシステムとプラットフォームの主要コンポーネント（名称や著作権などの資産を含む）の所有権を新しい非営利組織に移管し、Mastodonが特定の個人によって所有または管理されるべきではないという意図を明確にします。
{{< /fig-quote >}}

個人的には [Fediscovery] プロジェクトの話は知らなかったので興味を引いた。

{{< fig-quote type="markdown" title="「広場」は誰のものでもない、みんなのものであるべきだ" link="https://p2ptk.org/freedom-of-speech/5176" >}}
より直感的で使いやすいインターフェースを目指すとともに、[Fediscovery](https://www.fediscovery.org/)プロジェクトも前進させています。これは、Mastodonが主導し、NGI Searchの支援を受けて開発している、プライバシーに配慮したFediverse全体の検索ツールです。
{{< /fig-quote >}}

Mastodon というか ActivityPub は検索機能が弱く[^m1]，特にサーバ横断的な検索ができないのが不便ではある（公開された投稿であってもインデックス化されるのを嫌うユーザもいるので，一概にダメとは言えないが）。
[Fediscovery] プロジェクトでこの辺が改善されるのであればありがたい。

[^m1]: Mastodon は標準では全文検索機能を持っていない。ただし [Fedibird](https://fedibird.com/ "Fedibird") のように独自に全文検索機能を実装しているサーバもある。

## ソーシャルメディアは億万長者から逃げ切れるか？

[Free Our Feeds] も Mastodon も運営のための寄付を募っている。
億万長者の害意に対抗するにはお金が必要だからね。
日本では税控除の対象にならないだろうけど，よろしかったらどうぞって感じかな。

私は [Fedibird] に対しては継続的に[寄付](https://opencollective.com/fedibird-infrastructure "Fedibird Infrastructure - Open Collective")を行っている。
まぁ，微々たる金額だけど。
こちらは [Fedibird] のサーバ維持に限定したものなので Mastodon 全体の運営に寄付するなら別口で考えないとなぁ。

先日[紹介]({{< relref "/remark/2025/01/2024-mosaic-retrospective.md" >}} "パズルで2024年を振り返る")した “[2024: Mosaic Retrospective]” で遊びながら2024年を振り返ったりしているのだが，現職大統領が辞職間際に有罪判決を受けた自分の息子を恩赦した（2024年当時）とか有罪判決を受けたにも関わらず当選した次期大統領が就任後に自身を恩赦する（2024年当時）とかいうニュースを見て，米国は本当に破落戸の国になったんだなぁと感慨深かった。
政治トップが破落戸なら市場も安心して破落戸になれるというものである（笑） 億万長者が所有するソーシャルメディアも，どこぞの世紀末漫画よろしく「汚物は消毒だ～!!」とばかりに立ち塞がるものを蹴散らしていく。

Bluesky は英国で Mastodon は独国だっけ？ 逃げ切れるといいねぇ。

私は Mastodon 個人サーバという離れ小島でお一人さま生活を満喫しつつ，対岸の火事を眺めることにしよう。

## ブックマーク

- [メタ、「MAGA」になる──ファクトチェック廃止で、ヘイトスピーチ増加に懸念 | WIRED.jp](https://wired.jp/article/plaintext-meta-zuckerberg-maga-trump/)
- [弱者をいたぶるためのMetaのポリシー変更――表現の自由のために真になすべきこと » p2ptk[.]org](https://p2ptk.org/freedom-of-speech/5153)
- [Xがアルゴリズム変更へ　イーロン・マスク氏が優遇したい投稿とは？：Social Media Today - ITmedia マーケティング](https://marketing.itmedia.co.jp/mm/articles/2501/10/news045.html)
- [ソーシャルメディアを億万長者の魔の手から守るキャンペーン「Free Our Feeds」が登場 - GIGAZINE](https://gigazine.net/news/20250114-free-our-feeds/)
- [Mastodonが運営の大部分を非営利団体に移管することを発表 - GIGAZINE](https://gigazine.net/news/20250114-mastodon-transition-nonprofit-structure/)

[p2ptk.org]: https://p2ptk.org/ "P2Pとかその辺のお話R | Sharing is Caring"
[Free Our Feeds]: https://freeourfeeds.com/ "Free Our Feeds"
[Fediscovery]: https://www.fediscovery.org/ "Fediverse Discovery Providers"
[Fedibird]: https://fedibird.com/ "Fedibird"
[2024: Mosaic Retrospective]: https://store.steampowered.com/app/3380760/2024_Mosaic_Retrospective/ "2024: Mosaic Retrospective on Steam"

## 作業中の BGV (メン限配信以外)

- [【民俗学】今年の干支は巳！カオスな蛇の民俗学！【VTuber/ #諸星めぐる 】 - YouTube](https://www.youtube.com/watch?v=uq9V7WlnczQ)
- [【雑談】いろいろ始まりました！ちょっとまつたけ～！！！【儒烏風亭らでん #ReGLOSS 】 - YouTube](https://www.youtube.com/watch?v=gWo7tfVcM7Q)
- [【雑談】ゲリラでごめんね！新成人おめでとう！乾杯！🍻【古代日本史VTuber きら子】 - YouTube](https://www.youtube.com/watch?v=DHjJPACJZnQ)
