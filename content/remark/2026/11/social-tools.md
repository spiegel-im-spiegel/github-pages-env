+++
title = "分散型 SNS をもっと活用する"
date =  "2026-11-16T16:11:28+09:00"
description = "notestock と mochott"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "tools", "atproto", "blog", "activitypub", "search", "mastodon", "bluesky", "communication" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

分散型 SNS 用のサービスをメモがてら紹介する。

## mochott

[mochott] は [AT Protocol][atproto] および [standard.site] Lexicon で駆動するブログサービスである。
[Bluesky] など [atproto] 対応サービスのアカウントでサインアップでき，データは [mochott] のサーバではなく PDS に保存される。

この手のブログサービス自体は珍しくなくなったが，マルチ SNS クライアント「ぞーぺん（[ZonePane]）」が [mochott] への投稿機能を実装しようとしているようである。

{{< fig-gen >}}
<blockquote class="bluesky-embed" data-bluesky-uri="at://did:plc:mm5foq3omooqaqpdjfbjgohe/app.bsky.feed.post/3n2dycfcgdt3b" data-bluesky-cid="bafyreibeacmmtd75ylffm4i6aftetwyxcajhyg5i5ac4hwjdk7oaqffafm" data-bluesky-embed-color-mode="system"><p lang="ja">Blueskyの文字数制限は300文字までです。

これを超える投稿をしたい場合はBluesky本体ならスレッドとして分割投稿することになります。いにしえのTwitterと同じ方式でなじみがありますが、最近のXでは直接長い文章を投稿できたりしますよね。課金要素になっているあたりがいけてないわけですが、収益化とセットで導入されているのでこのあたりに鉱脈があると踏んでいたのがほんのり透けてみえるわけです。

さてそんな中、Bluesky界のトップランナーであるほりべあさんが長文ブログサービスmochottをリリースし、TOKIMEKIとのシームレスな連携を実現しました。TOKIMEKI…もっと見る</p>&mdash; たけうちひろあき (<a href="https://bsky.app/profile/did:plc:mm5foq3omooqaqpdjfbjgohe?ref_src=embed">@takke.jp</a>) <a href="https://bsky.app/profile/did:plc:mm5foq3omooqaqpdjfbjgohe/post/3n2dycfcgdt3b?ref_src=embed">2026年11月15日 21:11</a></blockquote><script async src="https://embed.bsky.app/static/embed.js" charset="utf-8"></script>
{{< /fig-gen >}}

こんな感じで，最初の部分を [Bluesky] へポストしつつ全文は [mochott] に置く，といった構造になるようだ。

これが本当に実装されるかどうかは分からないが，[ぞーぺん][ZonePane]ユーザとしては期待しつつ [mochott] のアカウントを取得しておくことにした。

{{< linkcard "18cd6d2b3f6831d36250aee3000ea1880a587b10" >}} <!-- https://mochott.site/baldanders.info/mochott-zonepane mochott を試してみる -->

継続して使うかどうかは[ぞーぺん][ZonePane]の開発状況次第だが，ちょっと触ってみた感じでは，イマドキの UI で使い勝手はよさそうである。

## notestock

[Mastodon] で最も不満な点は，検索性の悪さである。
これは，まぁ，半分くらいはしょうがない面もある。
Fediverse (federated universe) とは上手いこと言ったもので[^f1]，島宇宙たる各サーバはデータも運用ポリシーもそれぞれ独立している。
このため Fediverse 全体を完全に把握して検索するというのは相当難しい。

[^f1]: ここでいう federation は本来，分散コンピューティング環境における複数のサーバ間の連携を指す IT 用語なのだが， "Fediverse" という造語になって思想っぽいニュアンスが加わっている気がする。

にしても，自身の過去ポストくらいもう少し上手く検索できるようにしてほしい。

そこで [notestock] である。

[notestock] はものすごく簡単に言うと {{% emoji "X" %}} における [Twilog](https://twilog.togetter.com/ "Twitterのつぶやきをブログ形式で保存 - Twilog (ツイログ)") みたいなサービスで，自身の公開ポストを過去に遡って収集・保存してくれる。
検索については全文検索に対応しているので，いつ書いたか忘れてしまった過去のポストを簡単に見つけることができる。

また [notestock] で収集したポストを公開すると「グローバル検索」の対象にすることもできる（既定は非公開）。
自身のポストを公開している人は割といるみたいで，たとえば「メタクソ化」で検索すると

{{< fig-img-quote src="./notestock.png" title="メタクソ化 の検索結果 - notestock" link="./notestock.png" >}}

という感じに結構な数が引っかかる。

### Fedibird のすゝめ

実はもうひとつ全文検索が可能な手段がある。
それは [Fedibird] の検索機能の利用である。

[Fedibird] は標準の [Mastodon] からかなりの魔改造を施していて，全文検索機能に対応している。
さらに日本語圏にあるサーバとしてはユーザ規模も大きく[^fb1]，他サーバのユーザも積極的にフォローしているようだ[^fb2]。
このため，設定で検索対象を広げることで，かなり広範に検索をかけることができる（設定の「その他」にある「検索の対象とする範囲」で検索対象を指定できる）。
これがあるから [Fedibird] のアカウントは手放せないんだよなぁ。

[^fb1]: 2026年11月現在，[Fedibird] は（短期間での急激な流入を防ぐため）招待ユーザのみ受け入れている。主に運用コストの問題で手軽にスケールアップできないというのもあるみたい？
[^fb2]: [Mastodon] は他サーバのデータについて，そのサーバのユーザをフォローすることで，ユーザごとにポストデータをキャッシュしてくれる仕組みになっている。 [Mastodon] の初期の頃は複数のサーバのデータをキャッシュして検索ができるサードパーティのサービスがあったが，現在まで残っているものはほとんどない（はず）。

[Mastodon] を含む Fediverse の検索については（不完全ながらも）上述の方法を使い分けていくのがいいだろう。

[Bluesky]: https://bsky.app/ "Bluesky"
[Mastodon]: https://joinmastodon.org/ "Mastodon - Decentralized social media"
[ZonePane]: https://zonepane.com/ "ZonePane - マルチSNSクライアント"
[mochott]: https://mochott.site/ "mochott"
[notestock]: https://notestock.osa-p.net/ "notestock"
[ActivityPub]: https://www.w3.org/TR/activitypub/
[atproto]: https://atproto.com/ "The AT Protocol"
[standard.site]: https://standard.site/ "Standard.site - One schema. Every platform."
[Fedibird]: https://fedibird.com/ "Fedibird"
