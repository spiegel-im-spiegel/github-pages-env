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

接続手順だけ以下に抜き出しておくと

{{< fig-quote type="markdown" title="BlueskyとFediverse(MastodonやMisskey等)がつながった！簡単設定方法まとめ ～変わりつつあるSNS～" link="https://hiramy.hateblo.jp/entry/2024/10/04/230000" >}}
1. Bluesky側で、ブリッジ用アカウント([@ap.brid.gy](https://bsky.app/profile/ap.brid.gy))をフォローします。
2. Fediverse側で『`@bsky.brid.gy@bsky.brid.gy`』をフォローします。
3. BlueskyからFediverseのアカウントをフォローする場合は『`@アカウント名+ドメイン名+ap.brid.gy`』で検索します。
   例えば私のMisskeyのアカウントだと「`@hiramy.misskey.io.ap.brid.gy`」になります。
4. FediverseからBlueskyのアカウントをフォローする場合は『`@アカウント名+bsky.social(または他のドメイン)+@bsky.brid.gy`』で検索します。
   例えば私のBlueskyのアカウントだと「`@hiramy.bsky.social@bsky.bird.gy`」となります。
{{< /fig-quote >}}

だそうだ。
最初にブリッジ用の Bluesky アカウントおよび Mastodon アカウントをフォローするのがポイント。

私の場合なら Bluesky から 私の Masdodon アカウント（[`@spiegel@goark.fedicity.net`](https://goark.fedicity.net/@spiegel)）をフォローするには

-  [`@spiegel.goark.fedicity.net.ap.brid.gy`](https://bsky.app/profile/spiegel.goark.fedicity.net.ap.brid.gy)

からフォローできる。
逆に私の Bluesky アカウント（[`@baldanders.info`](https://bsky.app/profile/baldanders.info)）を[^bsky1] Mastodon 等でフォローする場合は

[^bsky1]: 私の Bluesky アカウント（[@baldanders.info](https://bsky.app/profile/baldanders.info)）はカスタム・アカウントである。カスタムアカウントの設定については「[Bluesky に旗を立てろ！]({{< ref "/remark/2023/04/plant-your-flag-mark-your-bluesky-territory.md" >}})」で紹介しているが，なにぶん古い記事なので，あまり参考にならないかも知れない。

-  `@baldanders.info@bsky.brid.gy`

で検索すればいいようだ。
なお Mastodon から見て上のアカウントはボットアカウントとして運用されている。
サーバによってはこの手のボットアカウントに（主にトラフィックの面などで）不寛容なところもあるのでご注意を。

ちなみに私は Mastodon でも Bluesky でも同じような投稿をしているので，どちらかでフォローすれば基本的に大丈夫っス。
Mastodon のほうが制約が少ないので長文になる傾向があるくらいかな。

詳しい使い方は[公式のドキュメント](https://fed.brid.gy/docs)を参照のこと。

昔 Twitter (現 {{% emoji "X" %}}) アカウントを Mastodon でフォローするためのブリッジ・サービスがあった気がするのだが，いつの間にかなくなっちゃったんだよね。
{{% emoji "X" %}} は基本的にボットアカウントに不寛容なのでしょうがないっちゃあしょうがない。

今回紹介した [Bridgy Fed] も将来的に（大人の事情で）なくなる可能性もあるけど，これからも継続的に運用されることを期待しよう。

## ブックマーク

- [Fediverse 関連のブックマーク]({{< ref "/bookmarks/fediverse.md" >}})

[Bridgy Fed]: https://fed.brid.gy/ "Bridgy Fed"
