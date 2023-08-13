+++
title = "Bluesky に旗を立てろ！"
date =  "2023-04-15T18:31:03+09:00"
description = "正直，どう使うかまるで考えてない。"
image = "/images/attention/kitten.jpg"
tags = [ "tools", "communication", "bluesky" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今だによく分かってない [Bluesky]。
なにせ「[AT (Authenticated Transfer) Protocol](https://atproto.com/ "The AT Protocol") って何？[^at1]」ってレベルで分からない（笑）

[^at1]: AT (Absolute Terror) Field の親戚か，思うた（笑）

最近話題の Mastodon や Nostr と何が違うん？ という感じでスルーしてたのだが（総ユーザ数が一億超えたら考えるって言ってたぢゃん＞自分）， Twitter や Mastodon の TL で連日連呼されるし，運良く invitation をいただける機会があったので，サインアップしてしまった。

携帯端末は iPhone しか対応してないみたいなので Web クライアントでサインアップしようとしたが， Web クライアントはサードパーティも含めて色々あるのな。
で，公式の Web クライアントは [staging.bsky.app] だそうな。
ということで，以下でサインアップした。

- [@baldanders.info](https://staging.bsky.app/profile/baldanders.info)

... `xxx.bsky.social` みたいなハンドル名にないってない理由は次節をどうぞ。

フォローはご自由にどうぞ。
正直，どう使うかまるで考えてない。
Nostr と同じで，こちらも “[Plant Your Flag, Mark Your Territory](https://krebsonsecurity.com/2018/06/plant-your-flag-mark-your-territory/ "Plant Your Flag, Mark Your Territory – Krebs on Security")” ってことで，旗を立てるだけにしておこう。

何が生き残るんだろうねぇ。
案外，すったもんだの挙句に残ったのは Twitter だけだった，なんてなオチもありそうだが（笑）

## [Bluesky] のアイデンティティ

Twitter は [Blue をお金で買えるようになった](https://p2ptk.org/freedom-of-speech/4384 "かくしてTwitterの青いチェックマークは「嘲笑」のシンボルとなった | p2ptk[.]org")ので論外だが，たとえば Mastodon なら自サイトに

```html
<link rel="me" href="https://social.example.com/@username">
```

などと記述し，緩くサイト間連携することでアイデンティティを記述できる。
じゃあ [Bluesky] はどうしてるのかと思ったら，ハンドル名をカスタムドメインに変更できるらしい。

- [Domain Names as Handles in Bluesky - Bluesky](https://blueskyweb.xyz/blog/3-6-2023-domain-names-as-handles-in-bluesky)

でも，この記事には手順が書いてないのよね。
ググって最初に見つけた記事では Web API を叩けみたいなことが書いてあったのだが，実際にやってみると `"Method Not Implemented"` と返ってきて設定できない。
他に情報がないかと探してみたら Zenn に投稿してる方がいた。

- [誰でもできる、Blueskyでカスタムドメインを簡単に設定する方法](https://zenn.dev/kato_shinya/articles/lets-set-custom-domain-in-bluesky)

公式 Web クライアントの Settings にあった “Change my handle” って項目はこれのことか！

{{< fig-img src="./settings.png" link="./settings.png" >}}

あとは上のリンク先記事の通りにやれば変更できる[^dn1]。
こうやって Bluesky のユーザは自らのアイデンティティを担保するというわけだ。

[^dn1]: 当然ながら連携するドメインは自前で用意する必要がある。かつ，ドメインに対して `TXT` レコードを設定できる権限が必要。

## ブックマーク

- [Bluesky -- Scrapbox](https://scrapbox.io/Bluesky/)
- [開発視点から見る、新しい分散型SNS「Bluesky」とAT Protocolの可能性 | gihyo.jp](https://gihyo.jp/article/2023/04/bluesky-atprotocol)
- [BlueskyのWebクライイアントSkylightとTokimeki Blueskyの比較 – あたしンちのおとうさんの独り言](https://atasinti.chu.jp/dad3/archives/64132)
- [「Bluesky」はSNSと言い切れない理由、そして世界初のミートアップが東京で行われた事情（1/2 ページ） - ITmedia NEWS](https://www.itmedia.co.jp/news/articles/2304/14/news175.html)
- [分散型SNS｢Bluesky｣って今どうなってるの？ | ギズモード・ジャパン](https://www.gizmodo.jp/2023/04/bluesky-now.html)
- [Blueskyの一ヶ月前史 | 点と接線。](https://riq0h.jp/2023/04/15/110859/)
- [「Bluesky」で本人認証をする方法、Twitterと違って無料＆自分で全部可能 - GIGAZINE](https://gigazine.net/news/20230421-bluesky-handle-domain-name/)
- [The AT Protocol Developer Ecosystem - Bluesky](https://blueskyweb.xyz/blog/4-21-2023-atproto-developer-ecosystem) : アプリケーションごとにパスワードを発行できるようにしたらしい
- [開発視点から見る、新しい分散型SNS「Bluesky」とAT Protocolの可能性 | gihyo.jp](https://gihyo.jp/article/2023/04/bluesky-atprotocol)
  - [GitHub - bluesky-social/indigo: Go source code for Bluesky's atproto services. NOT STABLE (yet)](https://github.com/bluesky-social/indigo) : 一応公式らしい
  - [GitHub - mattn/bsky: A cli application for bluesky social](https://github.com/mattn/bsky) : mattn さんが公開している CLI ツール。フル機能が実装されてるっぽい
  - [nostr と Bluesky に7つ bot を作り k8s で稼働させた](https://zenn.dev/mattn/articles/cbc42abe8be82b)
- [ツイッターの後継として注目度急上昇のBluesky。いま日本人が知っておくべきこと（徳力基彦） - 個人 - Yahoo!ニュース](https://news.yahoo.co.jp/byline/tokurikimotohiko/20230505-00348373)
- [「Bluesky」のAndroid版が登場、Twitter乗り換えの最有力候補の1つと目されTwitter元CEOのジャック・ドーシーがTwitterの反省点を踏まえて開発 - GIGAZINE](https://gigazine.net/news/20230420-bluesky-android/) : まだ最小限の機能しかない。せめて「共有」機能がないと...
  - [Bluesky - Apps on Google Play](https://play.google.com/store/apps/details?id=xyz.blueskyweb.app&hl=en_US)
- [Nostrのrabble氏の投稿の日本語訳 · GitHub](https://gist.github.com/kojira/94079413376d2a37c626d30a51fcb89f)
- [Decentralized Identifiers (DIDs) v1.0](https://www.w3.org/TR/did-core/)
- [Twitter代替有力候補SNS「Bluesky」でGIGAZINEが新着ニュース配信を開始＆どういう仕組みなのか解説 - GIGAZINE](https://gigazine.net/news/20230518-bluesky-gigazine/)
- [BlueskyのAT Protocolでリンクカード付きのpostを投稿する方法](https://zenn.dev/ryo_kawamata/articles/8d1966f6bb0a82)
- [GitHub ActionsからBluesky Socialへ簡単にポストする](https://zenn.dev/kato_shinya/articles/send-post-to-bluesky-via-github-actions)
- [アプリパスワードについて – Nextra](https://tokimekibluesky-docs.vercel.app/ja/apppassword)
- [Twitter代替SNS「Bluesky」で「ブロックを無視して投稿閲覧」「各ユーザーのブロック情報を確認」といった操作が可能な理由とは？ - GIGAZINE](https://gigazine.net/news/20230613-bluesky-block-public/)
- [SNS究極の自由、アルゴリズムもユーザーしだいの次世代Twitter｢Bluesky｣：エンジニアにインタビュー | ギズモード・ジャパン](https://www.gizmodo.jp/2023/06/bluesky-interview.html)
- [BINGO game on Bluesky](https://bingo.b35.jp/)
- [LangChain を使って Hacker News の日本語要約 Bluesky ボットを作ってみた](https://zenn.dev/ryo_kawamata/articles/98b7cc1c67ad0c)
- [SNS「Bluesky」800万ドル調達　「広告モデルではない収益化」へ、ドメイン販売開始 - ITmedia NEWS](https://www.itmedia.co.jp/news/articles/2307/07/news097.html)
- [No.31 「Bluesky」でDNSを使わずユーザー名を独自ドメインにする方法 - sakatori[sm] note](https://www.sakatori.net/note/?postid=31)

[Bluesky]: https://blueskyweb.xyz/
[staging.bsky.app]: https://staging.bsky.app/ "Bluesky"
