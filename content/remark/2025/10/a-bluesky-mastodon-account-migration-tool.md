+++
title = "Bluesky と Mastodon の間でアカウントの移行が可能らしい"
date =  "2025-10-21T09:28:29+09:00"
description = " 活動拠点を（どちらかに）移そうかなと考えてる人には一考の余地があるかな"
image = "/images/attention/kitten.jpg"
tags = [ "communication", "bluesky", "mastodon", "activitypub", "atproto" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

Bluesky の TL で見かけた記事：

- [Launch: Bounce from Mastodon to an existing Bluesky, Blacksky, or other ATProto Account](https://blog.anew.social/bounce-beta-now-live-2/)

へぇ， [Bounce] というサービスがあるのか。

FAQ 形式の[ドキュメント](https://bounce.anew.social/docs)を読むと Bluesky など [AT (Authenticated Transfer) Protocol](https://atproto.com/ "The AT Protocol") (aka atproto) 準拠のサービスと Mastodon および Pixelfed との間でアカウントを移行できるらしい。
具体的には移行先にあらかじめアカウントを作っておいて [Bridgy Fed] で接続するというもの。
移行先にはプロフィール，フォロワーおよびフォローの一部が反映される。
既存のポストは移行されないが，移行後のポストは（[Bridgy Fed] の機能で）両方のサービスに反映される。
また移行元のアカウントは [Bridgy Fed] の管理下に入るためユーザは直接制御できなくなる。

...という感じかな。
まだベータ版で，今のところ移行後は元に戻せないそうなのでご注意を。

既に [Bridgy Fed] で Mastodon-Bluesky 間を[接続]({{< ref "/remark/2024/10/bridgy-fed-for-bluesky.md" >}} "Fediverse と Bluesky を相互接続する")できてる人にはあまり旨味はないかもしれない。
活動拠点を（どちらかに）移そうかなと考えてる人には一考の余地があるだろう。
まだリスキーな感じもするけど。

あと Blacksky なんてサービスがあるんだな。
ググったら真っ先に[BlackSky](https://www.blacksky.com/ "Real-Time Space-Based Intelligence") という情報収集企業が出てきたが，これじゃないよな（笑） atproto 準拠のサービスならこっちか：

- [Blacksky](https://blacksky.community/)

名前被ってるけど大丈夫かなぁ。

まぁ，順調に分散化が進んでるってことだろう。

## ブックマーク

- [BlueskyやThreadsに受け継がれたネット原住民の叡智 – WirelessWire News](https://wirelesswire.jp/2024/04/86389/)

[Bounce]: https://bounce.anew.social/ "Bounce"
[Bridgy Fed]: https://fed.brid.gy/ "Bridgy Fed"

## 参考

{{% review-tatsujin "infoshare2" %}} <!-- 続・情報共有の未来 -->
{{% review-paapi "B00FW5SSCK" %}} <!-- ソーシャル・ネットワーク -->
