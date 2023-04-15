+++
title = "Bluesky に旗を立てろ！"
date =  "2023-04-15T18:31:03+09:00"
description = "正直，どう使うかまるで考えてない。"
image = "/images/attention/kitten.jpg"
tags = [ "tools", "communication" ]
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

[Bluesky]: https://blueskyweb.xyz/
[staging.bsky.app]: https://staging.bsky.app/ "Bluesky"
