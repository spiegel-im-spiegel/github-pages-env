+++
title = "Skyreader: AT Protocol で駆動する Feed Reader"
date =  "2026-02-25T17:33:49+09:00"
description = "同一の Lexicon でサービス間を水平展開できるというのは，色々と夢が広がる話である。"
image = "/images/attention/kitten.jpg"
tags = [ "tools", "atproto", "bluesky" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

[Leaflet] で見かけた記事。

- [Skyreader update - standard.site feeds, read-it-later workflow, and sustainability - Skyreader Dev Log](https://skyreader-dev.leaflet.pub/3mfmuhvwoos2g)

[Skyreader] が [standard.site] に対応したのか。
これは試してみるしかあるまい。

## Bluesky でログイン

{{< fig-img-quote src="./skyreader-0.png" title="Skyreader" link="https://skyreader.app/" width="757" lang="en">}}

利用には Bluesky アカウントが必要なようだ。
OAuth 認証。

{{< fig-img-quote src="./login.png" title="Skyreader: login" link="https://skyreader.app/" width="757" lang="en">}}

{{< fig-img-quote src="./oauth.png" title="Skyreader: OAuth" link="https://skyreader.app/" width="757" lang="en">}}

ドメインを確認すること。

## フィードの登録

当然ながら初期状態は空っぽ。

{{< fig-img-quote src="./skyreader-1.png" title="Skyreader" link="https://skyreader.app/" width="806" lang="en">}}

左上の “All” の部分をクリックするとメニューが出てくる。

{{< fig-img-quote src="./skyreader-2.png" title="Skyreader" link="https://skyreader.app/" width="806" lang="en">}}

メニュー右上の “＋” の部分をクリックして “Add Feed” を選択するとフィードを登録できる。

{{< fig-img-quote src="./skyreader-3.png" title="Skyreader" link="https://skyreader.app/" width="806" lang="en">}}

おぉ！ [Leaflet] で購読（subscribe）したブログが一覧表示されているな。

どうやら [standard.site] 対応の他サービスで購読したブログを読み取って取り込めるようになっているらしい。
素晴らしい！

[standard.site] に対応しているサービスであれば使用しているハンドル名を入力すれば購読できるらしい。
例えば [atproto.com のブログ](https://atproto.com/ja/blog "Blog - AT Protocol")は [standard.site] に対応しているので [`@atproto.com`](https://bsky.app/profile/atproto.com "AT Protocol Developers (@atproto.com) — Bluesky") と入力すれば

{{< fig-img-quote src="./skyreader-4a.png" title="Skyreader" link="https://skyreader.app/" width="806" lang="en">}}

{{< fig-img-quote src="./skyreader-4b.png" title="Skyreader" link="https://skyreader.app/" width="806" lang="en">}}

こんな感じに有効なフィードが表示されるので選択する（この場合はどちらも同じ内容）。

通常の RSS フィードも登録できる。
この場合はブログページの URL を指定すればよい。

{{< fig-img-quote src="./skyreader-5.png" title="Skyreader" link="https://skyreader.app/" width="806" lang="en">}}

[`https://go.dev/blog/`](https://go.dev/blog/ "The Go Blog - The Go Programming Language") のフィードはひとつのみなので，そのまま取り込める。

{{< fig-img-quote src="./skyreader-6.png" title="Skyreader" link="https://skyreader.app/" width="806" lang="en">}}

表示される各記事に対しては

{{< fig-img-quote src="./skyreader-7.png" title="Skyreader" link="https://skyreader.app/" width="806" lang="en">}}

以下の操作ができる。

- **Read**: 既読・未読を切り替える
- **Save**: 後で読むために保存する
- **Share**: 記事を共有する
- **Open**: 別タブで Web ページを開く
- **Full**: 記事の内容をフル表示する
- **Tag**: タグを付ける

共有とタグについてはどうやって使うのか分からない。
今後の実装になるのかな？

「後で読む」については次節で。

## 後で読む

記事を「後で読む」ために保存することができる。
前節に書いたようにフィードの各記事に対して “Save” することもできるし URL を直接指定して “Save” することもできる。

{{< fig-img-quote src="./skyreader-2.png" title="Skyreader" link="https://skyreader.app/" width="806" lang="en">}}

メニューの右上の “＋” の部分をクリック。
“ Save Article by URL” を選択して URL を入力すれば保存できる。

{{< fig-img-quote src="./skyreader-8.png" title="Skyreader" link="https://skyreader.app/" width="806" lang="en">}}

見終わったページは “Archive” に移動できる。

{{< fig-img-quote src="./skyreader-9.png" title="Skyreader" link="https://skyreader.app/" width="806" lang="en">}}

これはありがたい機能。

## Skyreader の制限

[Skyreader] は現在ベータ版で無料で利用できるが，登録できるフィードと保存できる URL に制限がある。

{{< fig-img-quote src="./settings.png" title="Skyreader: Settings" link="https://skyreader.app/settings" width="806" lang="en">}}

他所の Feed Reader サービスをガッツリ使ってる人には物足りないスペックだと思うが，正式リリースの際はサブスクリプション制にする可能性があるらしい（現在も支援を募集している）。
他にも色々と機能を追加する予定のようだ。

{{< fig-quote type="markdown" title="Skyreader update - standard.site feeds, read-it-later workflow, and sustainability" link="https://skyreader-dev.leaflet.pub/3mfmuhvwoos2g" lang="en" >}}
Skyreader is starting to become a complete reading environment: subscribe to anything (RSS, standard.site, more coming soon), save and manage your reading list, and control your data on atproto. Browser extensions, more integrations, and full PDS sync for highlights and reading position coming soon.
{{< /fig-quote >}}

楽しみにしつつ，試してみるつもりだ。
しかし [standard.site] はめっちゃ使い勝手がいい。
同一の Lexicon でサービス間を水平展開できるというのは，色々と夢が広がる話である。

## ブックマーク

- [Leaflet と The Atmosphere]({{< relref "./leaflet-and-the-atmosphere.md" >}})

[Skyreader]: https://skyreader.app/ "Skyreader"
[Leaflet]: https://leaflet.pub/ "Leaflet"
[atproto]: https://atproto.com/ "The AT Protocol"
[standard.site]: https://standard.site/ "Standard.site - One schema. Every platform."
