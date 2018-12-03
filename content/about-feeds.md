+++
date = "2017-01-25T17:28:42+09:00"
update = "2018-05-21T19:04:37+09:00"
title = "更新情報について"
tags = ["site", "feed"]
description = "text.Baldanders.info ではサイトの更新情報を提供しています。"

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[text.Baldanders.info]({{< rlnk "/" >}}) （以降「[本サイト]」と呼びます）ではサイトの更新情報を提供しています。
更新情報の詳細は以下の通りです。
なお，フォーマットはいずれも Atom Syndication 形式で， [Feedly](https://feedly.com/) などお好みの feed reader に取り込めます。

## サイト全体の更新情報

[本サイト]全体の更新情報は以下の URL から取得できます。

- [{{< lnk "index.xml" >}}]({{< rlnk "index.xml" >}})
- [{{< lnk "sitemap.xml" >}}]({{< rlnk "sitemap.xml" >}})

## セクションごとの更新情報

[本サイト]では以下のセクションごとに更新情報を提供しています。
URL は以下の通りです。

| セクション名                                                      | 更新情報の URL                                                                                           |
|:----------------------------------------------------------------- |:-------------------------------------------------------------------------------------------------------- |
| [しっぽのさきっちょ]({{< ref path="/remark" >}})                  | [{{< ref path="/remark" outputFormat="rss" >}}]({{< ref path="/remark" outputFormat="rss" >}})           |
| [ゼロから始める Hugo]({{< ref path="/hugo" >}})                   | [{{< ref path="/hugo" outputFormat="rss" >}}]({{< ref path="/hugo" outputFormat="rss" >}})               |
| [プログラミング言語 Go]({{< ref path="/golang" >}})               | [{{< ref path="/golang" outputFormat="rss" >}}]({{< ref path="/golang" outputFormat="rss" >}})           |
| [改訂3版： CC Licenses について]({{< ref path="/cc-licenses" >}}) | [{{< ref path="/cc-licenses" outputFormat="rss" >}}]({{< ref path="/cc-licenses" outputFormat="rss" >}}) |
| [[OpenPGP の実装]({{< ref path="/openpgp" >}})                    | [{{< ref path="/openpgp" outputFormat="rss" >}}]({{< ref path="/openpgp" outputFormat="rss" >}})         |
| [ブックマーク集]({{< ref path="/bookmarks" >}})                   | [{{< ref path="/bookmarks" outputFormat="rss" >}}]({{< ref path="/bookmarks" outputFormat="rss" >}})     |
| [リリース情報]({{< ref path="/release" >}})                       | [{{< ref path="/release" outputFormat="rss" >}}]({{< ref path="/release" outputFormat="rss" >}})         |

## タグごとの更新情報

[本サイト]では[タグ]({{< ref "/tags" >}})ごとに更新情報を提供しています。
たとえば [astronomy]({{< ref "/tags/astronomy" >}}) タグであれば以下の URL となります。

- [{{< ref "/tags" >}}**astronomy**/index.xml]({{< ref "/tags/astronomy" >}}index.xml)

## ブックマーク

- [RSSリーダーの効率的な使い方「RSSとは」な人に朗報！Feedlyでフィード購読のすすめ](http://millkeyweb.com/rss-feedly/)
- [一番使いやすいRSSリーダーは？Reeder、Sylfeed、Bylineを比較！ – iBitzEdge](https://i-bitzedge.com/ios-apps/the-best-rss-readers-to-use)
- [最強RSSリーダーを選ぶならfeedlyではなくinoreaderが一番おすすめかもしれない１２の理由 - ウェブ企画ラボ](https://webkikaku.co.jp/blog/software/inoreader/)
- [SlackをRSSリーダーにしたら、他のRSSリーダーサービスを使わなくなった話 - Qiita](http://qiita.com/kozyty@github/items/f094ae8fea08b471ae08)

[本サイト]: {{< rlnk "/" >}} "text.Baldanders.info"
