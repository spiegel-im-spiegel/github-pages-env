+++
title = "Fediverse と Bluesky を相互接続する"
date =  "2024-10-06T09:46:04+09:00"
description = "TL が流れて忘れてしまう前に覚え書きとして残しておく。"
image = "/images/attention/kitten.jpg"
tags = [ "activitypub", "mastodon", "bluesky", "communication" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

TL が流れて忘れてしまう前に覚え書きとして残しておく。

Fediverse と Bluesky を相互接続する[サービス][Bridgy Fed]があるらしい。
接続方法とか運用上の制約については以下の記事が参考になる。
感謝！

- [BlueskyとFediverse(MastodonやMisskey等)がつながった！簡単設定方法まとめ ～変わりつつあるSNS～ - ひらみ漂流記](https://hiramy.hateblo.jp/entry/2024/10/04/230000)

接続手順だけ以下に書き出しておこう。
詳しい使い方は[公式のドキュメント](https://fed.brid.gy/docs)を参照のこと。

## Bluesky から Fediverse への接続

[Bridgy Fed] を使った Bluesky から Fediverse への接続では，まずアカウントのオーナーが手順を通して明示的に Fediverse からのフォローを許可する必要がある。
手順自体は簡単で Fediverse に公開したい Bluesky アカウントでブリッジ用アカウント [`@ap.brid.gy`](https://bsky.app/profile/ap.brid.gy) をフォローすればよい。
すると DM にこんなメッセージが来る。

{{< fig-img src="./message-from-bridgy-fed.png" title="Chat — Bluesky" link="./message-from-bridgy-fed.png" width="589" >}}

これは私の Bluesky アカウント（[`@baldanders.info`](https://bsky.app/profile/baldanders.info)）でフォローした場合[^bsky1]。
このアカウント名にサーバ名 `@bsky.brid.gy` を付加した名前が Mastodon などの Fediverse 上のアカウント名になるわけだ。

[^bsky1]: 私の Bluesky アカウント（[@baldanders.info](https://bsky.app/profile/baldanders.info)）はカスタム・アカウントである。カスタムアカウントの設定については「[Bluesky に旗を立てろ！]({{< ref "/remark/2023/04/plant-your-flag-mark-your-bluesky-territory.md" >}})」で紹介しているが，なにぶん古い記事なので，あまり参考にならないかも知れない。

こんな感じ。

{{< div-box type="markdown" >}}
`@baldanders.info@bsky.brid.gy`
{{< /div-box>}}

このアカウント名を公開し Mastodon などのサービス上で検索・フォローしてもらえばよい。

なお Mastodon から見てこのアカウントはボットアカウントとして定義されている。
サーバによってはこの手のボットアカウントに（主にトラフィックの面などで）不寛容なところもあるのでご注意を。

## Fediverse から Bluesky への接続

同じように Fediverse から Bluesky への接続についても，まず Fediverse アカウントのオーナーがブリッジ用アカウントを `@bsky.brid.gy@bsky.brid.gy` をフォローする。
すると DM にこんなメッセージが来る。

{{< fig-img src="./message-from-bridgy-fed-2.png" title="Bridgy Fed for Bluesky: “Welcome to Bridgy Fed! Your …”" link="./message-from-bridgy-fed-2.png" width="573" >}}

これは私の Masdodon アカウント（[`@spiegel@goark.fedicity.net`](https://goark.fedicity.net/@spiegel)）でフォローした場合。

Fadiverse アカウントは `@username@servername` という構造になっているが，これを `@username.servername` に変換し（中の `@` マークをピリオドに置き換える）後ろに `.ap.brid.gy` を付加したものが Bluesky から見えるアカウント名となる。
具体的にはこんな感じ。

{{< div-box type="markdown" >}}
[`@spiegel.goark.fedicity.net.ap.brid.gy`](https://bsky.app/profile/spiegel.goark.fedicity.net.ap.brid.gy)
{{< /div-box>}}

このアカウントを公開し Bluesky でフォローしてもらえばよい。

## 昔は...

昔 Twitter (現 {{% emoji "X" %}}) アカウントを Mastodon でフォローするためのブリッジ・サービスがあった気がするのだが，いつの間にかなくなっちゃったんだよね。
確か個人で運用されてたようだし，そもそも {{% emoji "X" %}} は基本的にボットアカウントに不寛容なのでしょうがないっちゃあしょうがない。

今回紹介した [Bridgy Fed] も将来的に（大人の事情で）なくなる可能性はあるけど，これからも継続的に運用されることを期待しよう。

ちなみに私は Mastodon でも Bluesky でも同じような投稿をしているので，どちらかでフォローすれば基本的に大丈夫っス。
Mastodon のほうが制約が少ないので長文になる傾向があるかな。

## ブックマーク

- [Fediverse 関連のブックマーク]({{< ref "/bookmarks/fediverse.md" >}})

[Bridgy Fed]: https://fed.brid.gy/ "Bridgy Fed"
