+++
title = "AT Protocol & ActivityPub 関連サービスの利用について"
date =  "2026-03-06T11:36:06+09:00"
description = "AT Protocol ベースのサービスをいくつか試しているが，そろそろ覚えきれなくなってきたので，整理のためにここでまとめておく。"
isCJKLanguage = true
image = "/images/attention/site.jpg"
tags = [ "site", "atproto", "bluesky", "blog" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = true
+++

## ATProto/Atmosphere 関連サービス

[AT Protocol][atproto] ベースのサービスをいくつか試しているが，そろそろ覚えきれなくなってきたので，整理のためにここでまとめておく。

<table>
  <thead>
    <tr>
      <th>サービス</th>
      <th>アカウント等</th>
      <th>備考</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td rowspan="5" class="nowrap"><a href="https://bsky.app/">Bluesky</a></td>
      <td class="nowrap">
        <a href="https://bsky.app/profile/did:plc:w7fwp2mtlaaffnx42wsc76wt"><code>@baldanders.info</code></a>
      </td>
      <td>
        メインアカウント,
        <a href="https://pdsls.dev/at://did:plc:w7fwp2mtlaaffnx42wsc76wt">PDSls</a>
      </td>
    </tr><tr>
      <td class="nowrap">
        <a href="https://bsky.app/profile/did:plc:hbzmqswkx5pbg5fhr33pw4iw"><code>@fanapod.baldanders.info</code></a>
      </td>
      <td>
        非公式 APOD 配信ボット,
        <a href="https://pdsls.dev/at://did:plc:hbzmqswkx5pbg5fhr33pw4iw">PDSls</a>
      </td>
    </tr><tr>
      <td class="nowrap">
        <a href="https://bsky.app/profile/did:plc:qv4cel4jrf5qqzfq6ddvp6yv"><code>@crawler.baldanders.info</code></a>
      </td>
      <td>
        Web クローラー ボット,
        <a href="https://pdsls.dev/at://did:plc:qv4cel4jrf5qqzfq6ddvp6yv">PDSls</a>
      </td>
    </tr><tr>
      <td class="nowrap">
        <a href="https://bsky.app/profile/did:plc:4kzexuq4xieuarc3q6lmkeou"><code>@photos.baldanders.info</code></a>
      </td>
      <td>
        お散歩カメラ, 実験用,
        <a href="https://pdsls.dev/at://did:plc:4kzexuq4xieuarc3q6lmkeou">PDSls</a>
      </td>
    </tr><tr>
      <td class="nowrap">
        <a href="https://bsky.app/profile/did:plc:26eoqvsiov3hmsfqruoi4d3s"><code>@goark.bsky.social</code></a>
      </td>
      <td>
        開発・実験用,
        <a href="https://pdsls.dev/at://did:plc:26eoqvsiov3hmsfqruoi4d3s">PDSls</a>
      </td>
    </tr><tr>
      <td class="nowrap"><a href="https://leaflet.pub/">Leaflet</a></td>
      <td class="nowrap"><a href="https://spiegel.leaflet.pub/">お散歩カメラ</a></td>
      <td>ブログ試用中, <a href="https://bsky.app/profile/did:plc:w7fwp2mtlaaffnx42wsc76wt">メイン</a>アカウント使用</td>
    </tr><tr>
      <td class="nowrap"><a href="https://pckt.blog/">Pckt</a></td>
      <td class="nowrap"><a href="https://osanpo.pckt.blog">お散歩カメラ</a></td>
      <td>ブログ試用中, <a href="https://bsky.app/profile/did:plc:4kzexuq4xieuarc3q6lmkeou">photos</a>アカウント使用</td>
    </tr><tr>
      <td class="nowrap"><a href="https://trilinesat.suibari.com/">TriLinesAt</a></td>
      <td class="nowrap">&mdash;</td>
      <td>三行日記, <a href="https://bsky.app/profile/did:plc:w7fwp2mtlaaffnx42wsc76wt">メイン</a>アカウント使用</td>
    </tr><tr>
      <td class="nowrap"><a href="https://skyreader.app/">Skyreader</a></td>
      <td class="nowrap">&mdash;</td>
      <td>試用中, <a href="https://bsky.app/profile/did:plc:w7fwp2mtlaaffnx42wsc76wt">メイン</a>アカウント使用</td>
    </tr>
  </tbody>
</table>

他にも面白そうなサービスが色々と登場しているので，これからも増えていくかも知れない。
`bsky.social` 以外の PDS で Bluesky の API を使えるようになれば `pckt.cafe` とかも使いやすくなるのだが... この辺は調査中。

## ActivityPub/Fediverse 関連サービス

ついでに，私が利用している ActivityPub/Fediverse 関連サービスも挙げておく。
といっても今のところは Mastodon しか使ってないのだが。

<table>
  <thead>
    <tr>
      <th>サービス</th>
      <th>アカウント等</th>
      <th>備考</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td rowspan="3" class="nowrap"><a href="https://goark.fedicity.net/">Goark</a></td>
      <td class="nowrap">
        <a href="https://goark.fedicity.net/@spiegel"><code>@spiegel@goark.fedicity.net</code></a>
      </td>
      <td>メインアカウント</td>
    </tr><tr>
      <td class="nowrap">
        <a href="https://goark.fedicity.net/@osanpo"><code>@osanpo@goark.fedicity.net</code></a>
      </td>
      <td>Web クローラー ボット</td>
    </tr><tr>
      <td class="nowrap">
        <a href="https://goark.fedicity.net/@goark"><code>@goark@goark.fedicity.net</code></a>
      </td>
      <td>開発・実験用, 今のところボット運用</td>
    </tr><tr>
      <td class="nowrap"><a href="https://fedibird.com/">Fedibird</a></td>
      <td class="nowrap">
        <a href="https://fedibird.com/@spiegel"><code>@spiegel@fedibird.com</code></a>
      </td>
      <td>予備系</td>
    </tr><tr>
      <td class="nowrap"><a href="https://mstdn.jp/">mstdn.jp</a></td>
      <td class="nowrap">
        <a href="https://mstdn.jp/@spiegel"><code>@spiegel@mstdn.jp</code></a>
      </td>
      <td>休眠中</td>
    </tr>
  </tbody>
</table>

PixcelFed も気になっているのだが，これも API とかどうなってるんだろうと踏み出せてない。
まぁ Mastodon が問題なく使えてるので優先順位は低め。

## Bridgy Fed 連携

[Bridgy Fed] は [Bluesky] や [Mastodon] 等の ActivityPub 連合（Fediverse）を連携させるサービスで，現在 Bluesky → Fediverse と Fediverse → Bluesky のブリッヂが提供されている。
また Web サイトを [Bluesky] や Fediverse へ参加させる機能も提供されている。

[Bridgy Fed] について詳しくは[ドキュメント](https://fed.brid.gy/docs "Bridgy Fed")を参照のこと。
ここでは本サイトで提供している [Bridgy Fed] 連携の設定を挙げておく。

### Bluesky → Fediverse

{{< fig-mermaid >}}
graph LR
  BridgyFed(("Bridgy Fed"))
  Bluesky["Bluesky"]
  Mastodon["Fediverse\n(Mastodon, ...)"]

  Bluesky e2@--> BridgyFed
  BridgyFed e1@--> Mastodon

  e1@{ animate: true }
  e2@{ animate: true }
{{< /fig-mermaid >}}

<table>
  <thead>
    <tr>
      <th>Bluesky</th>
      <th>Fediverse</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td class="nowrap">
        <a href="https://bsky.app/profile/did:plc:w7fwp2mtlaaffnx42wsc76wt"><code>@baldanders.info</code></a>
      </td>
      <td class="nowrap">
        <a href="https://bsky.brid.gy/r/https://bsky.app/profile/baldanders.info"><code>@baldanders.info@bsky.brid.gy</code></a>
      </td>
    </tr><tr>
      <td class="nowrap">
        <a href="https://bsky.app/profile/did:plc:hbzmqswkx5pbg5fhr33pw4iw"><code>@fanapod.baldanders.info</code></a>
      </td>
      <td class="nowrap">
        <a href="https://bsky.brid.gy/r/https://bsky.app/profile/fanapod.baldanders.info"><code>@fanapod.baldanders.info@bsky.brid.gy</code></a>
      </td>
    </tr>
  </tbody>
</table>

### Mastodon → Bluesky

{{< fig-mermaid >}}
graph LR
  BridgyFed(("Bridgy Fed"))
  Bluesky["Bluesky\n(and Atmosphere)"]
  Mastodon["Mastodon"]

  Mastodon e1@--> BridgyFed
  BridgyFed e2@--> Bluesky

  e1@{ animate: true }
  e2@{ animate: true }
{{< /fig-mermaid >}}

<table>
  <thead>
    <tr>
      <th>Mastodon</th>
      <th>Bluesky</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td class="nowrap">
        <a href="https://goark.fedicity.net/@spiegel"><code>@spiegel@goark.fedicity.net</code></a>
      </td>
      <td class="nowrap">
        <a href="https://bsky.app/profile/mstdn.baldanders.info"><code>@mstdn.baldanders.info</code></a>
        (<a href="https://pdsls.dev/at://did:plc:fe7f76dxtltjam5kh5xk3blz">PDSls</a>)
      </td>
    </tr>
  </tbody>
</table>

### Web → Bluesky & Fediverse

{{< fig-mermaid >}}
graph LR
  Website["Web site"]
  BridgyFed(("Bridgy Fed"))
  Bluesky["Bluesky\n(and Atmosphere)"]
  Mastodon["Fediverse\n(Mastodon, ...)"]

  Website e1@--> BridgyFed
  BridgyFed e2@-->Mastodon
  BridgyFed e3@-->Bluesky

  e1@{ animate: true }
  e2@{ animate: true }
  e3@{ animate: true }
{{< /fig-mermaid >}}

<table>
  <thead>
    <tr>
      <th>Web site</th>
      <th>Fediverse / Bluesky</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td rowspan="2" class="nowrap">
        <a href="https://text.baldanders.info/"><code>text.baldanders.info</code></a>
      </td>
      <td class="nowrap">
        <a href="https://web.brid.gy/@text.baldanders.info"><code>@text.baldanders.info@web.brid.gy</code></a>
      </td>
    </tr><tr>
      <td class="nowrap">
        <a href="https://bsky.app/profile/text.baldanders.info"><code>@text.baldanders.info</code></a>
        (<a href="https://pdsls.dev/at://did:plc:crwrada4cit2ijjxuwz4bsjc">PDSls</a>)
      </td>
    </tr>
  </tbody>
</table>

## ブックマーク

- [Mastodon と Bluesky でボット運用はじめました]({{< ref "/remark/2023/07/crawler.md" >}})
- [Fediverse と Bluesky を相互接続する]({{< ref "/remark/2024/10/bridgy-fed-for-bluesky.md" >}})
- [Bluesky と Mastodon の間でアカウントの移行が可能らしい]({{< ref "/remark/2025/10/a-bluesky-mastodon-account-migration-tool.md" >}})
- [三行日記はじめました]({{< ref "/remark/2026/01/trilinesat.md" >}})
- [Leaflet と The Atmosphere]({{< ref "/remark/2026/02/leaflet-and-the-atmosphere.md" >}})
- [Skyreader: AT Protocol で駆動する Feed Reader]({{< ref "/remark/2026/02/skyreader-powered-by-atproto.md" >}})
- [Pckt に新しい PDS レジストリを作る]({{< ref "/remark/2026/03/creating-a-new-pds-registry-on-pckt.md" >}})
- [そろそろ PDS を個人で管理する時代か？]({{< ref "/remark/2026/03/pds-for-everyone.md" >}})

[atproto]: https://atproto.com/ "The AT Protocol"
[PDSls]: https://pdsls.dev/ "PDSls"
[standard.site]: https://standard.site/ "Standard.site - One schema. Every platform."
[pckt]: https://pckt.blog/ "pckt.blog - pckt"
[Leaflet]: https://leaflet.pub/ "Leaflet"
[Bluesky]: https://bsky.app/ "Bluesky"
[Mastodon]: https://joinmastodon.org/ "Mastodon - Decentralized social media"
[Bridgy Fed]: https://fed.brid.gy/ "Bridgy Fed"
