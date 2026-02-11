+++
title = "Leaflet と The Atmosphere"
date =  "2026-02-11T18:45:00+09:00"
description = "入院中の暇つぶしに遊んだブログサービス Leaflet を紹介してみる。"
image = "/images/attention/kitten.jpg"
tags = [ "atproto", "bluesky", "tools", "blog" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

無事に退院できました。
入院中の話（後編）はそのうち書くとして，今回は入院中の暇つぶしに遊んだブログサービス [Leaflet] を紹介してみる。

## Leaflet (leaflet.pub)

[Leaflet (leaflet.pub)][Leaflet] は [AT (Authenticated Transfer) Protocol][atproto] (以降 [atproto]) ベースのブログサービスである。
さっそく試してみた。

{{< fig-img-quote src="./leaflet.png" title="Leaflet" link="https://leaflet.pub/" width="900" lang="en">}}

そのまま “Write a Doc” で書き始めることもできるけど[^l1] “Start a Publication and blog in the Atmosphere” で Bluesky アカウントと紐づけることもできる。

{{< fig-img-quote src="./leaflet-sign-up-1.png" title="Log In or Sign Up" link="https://leaflet.pub/" width="900" lang="en">}}

“use an ATProto handle” を選ぶと Bluesky 以外の [atproto] ベースのサービスのアカウントでログインできるらしいのだが，私は Bluesky 以外にアカウントを持ってないので省略。
ここは素直に “Log In/Sign Up with Bluesky” ボタンを押す。

{{< fig-img-quote src="./leaflet-sign-up-2.png" title="Log In or Sign Up" link="https://leaflet.pub/" width="900" lang="en">}}

Bluesky のハンドル名（`user.bsky.social` など）を入力して認証を開始する。
ちなみに OAuth 認証なのでアプリパスワードではなく通常のパスワードを使う必要がある。
認証が成功すると OAuth のアクセス承認画面がでる。

{{< fig-img-quote src="./oauth-approve.png" title="承認する" link="https://leaflet.pub/" width="900" lang="en">}}

接続先（`leaflet.pub`）等をよく確認すること。
ここで “承認する” を選ぶと Bluesky アカウントと紐付けられ Publication 作成画面に移動する。

{{< fig-img-quote src="./new-publication.png" title="Create Your Publication!" link="https://leaflet.pub/" width="900" lang="en">}}

こんな感じに Publication 名（ブログ名）やドメイン名等を決めて “Create Publication!” ボタンを押すと Publication が作成される。

{{< fig-img-quote src="./publication.png" title="Create Your Publication!" link="https://leaflet.pub/" width="900" lang="en">}}

ひとつのアカウントに対して複数の Publication を作成することができる。
さっそく “New Draft” ボタンを押して記事を書いてみよう。

{{< fig-img-quote src="./draft-1.png" title="Draft" link="https://leaflet.pub/" width="900" lang="en">}}

編集は WYSIWYG (What You See Is What You Get) スタイルで，今どきの markdown 記法とかは使えない。
最初「絵とかどうやって入れるの？」と途方に暮れたが，行頭に `/` を入力するとメニューが出てくるようだ。
絵図を挿入する以外にもコードや数式やリンクカードなども挿入できる。

{{< fig-img-quote src="./draft-2.png" title="Draft" link="https://leaflet.pub/" width="900" lang="en">}}

でも表は書けないっぽい。

せっかく「お散歩カメラ」とタイトルを付けたので，ここは写真を貼り付けておく。
こんな感じ。

{{< fig-img-quote src="./draft-3.png" title="Draft" link="https://leaflet.pub/" width="900" lang="en">}}

記事の保存は自動で行われる。
公開する際は左上の “Publish!” ボタンを押す。

{{< fig-img-quote src="./draft-4.png" title="Draft" link="https://leaflet.pub/" width="900" lang="en">}}

すると以下の確認画面に移動する。

{{< fig-img-quote src="./“publish.png" title="Publish!" link="https://leaflet.pub/" width="900" lang="en">}}

“Share on Bluesky” を選ぶと記事へのリンクが Bluesky にポストされる。
“Publish this Post!” ボタンを押すと記事が公開される。
Bluesky へのポストはこんな感じ。

{{< fig-gen >}}
<blockquote class="bluesky-embed" data-bluesky-uri="at://did:plc:w7fwp2mtlaaffnx42wsc76wt/app.bsky.feed.post/3meibm6q6oc2r" data-bluesky-cid="bafyreigopqsviwupg4sawxqt3tqlgxrhf2qxbx76fw3r4zd3wlwkbtkzky" data-bluesky-embed-color-mode="system"><p lang="">2026-02-08 のお散歩カメラ<br><br><a href="https://bsky.app/profile/did:plc:w7fwp2mtlaaffnx42wsc76wt/post/3meibm6q6oc2r?ref_src=embed">[image or embed]</a></p>&mdash; Spiegel＠巡遊者 (<a href="https://bsky.app/profile/did:plc:w7fwp2mtlaaffnx42wsc76wt?ref_src=embed">@baldanders.info</a>) <a href="https://bsky.app/profile/did:plc:w7fwp2mtlaaffnx42wsc76wt/post/3meibm6q6oc2r?ref_src=embed">February 10, 2026 at 2:54 PM</a></blockquote><script async src="https://embed.bsky.app/static/embed.js" charset="utf-8"></script>
{{< /fig-gen >}}

実際に公開した記事はこんな感じ。

{{< fig-img src="./posted.png" title="雪やこんこ - お散歩カメラ" link="https://spiegel.leaflet.pub/3meibm4cbj22r" width="900">}}

凝ったカスタマイズはできなさそうだが，基本的な機能は押さえてあるし，シンプルで使いやすいサービスだと思う。

## Leaflet と The Atmosphere

最初に述べたように [Leaflet] は [atproto] ベースのブログサービスである。
[atproto] の構成はこんな感じ。

{{< fig-img-quote class="lightmode" src="./Bluesky–AT_Protocol_federation_architecture.svg.png" title="File:Bluesky–AT Protocol federation architecture.svg - Wikimedia Commons (CC BY 4.0)" link="https://commons.wikimedia.org/wiki/File:Bluesky%E2%80%93AT_Protocol_federation_architecture.svg" width="960" lang="en">}}

前に紹介した[三行日記]({{< ref "/remark/2026/01/trilinesat.md" >}} "三行日記はじめました")と同じように [Leaflet] も（おそらく[^l1]）独自のストレージを持っているわけではなく PDS (Personal Data Server) の各ユーザのリポジトリに記事データを格納している。

[^l1]: [Leaflet] ではサインインせずに記事を書いて公開することができるのだが，どういう仕組みか分からない。

さらに [Leaflet] は最近 [standard.site] と呼ばれる（ブログなどの）長文記事用の Lexicon を[採用](https://lab.leaflet.pub/3md4qsktbms24 "Leaflet, standard.site, and open social publishing! - Leaflet Lab Notes")した[^l2]。
[Standard.site][standard.site] を採用しているサービスとしては他に [pckt](https://pckt.blog/ "pckt.blog") や [Offprint](https://offprint.app/ "Own your words. Own your audience. | Offprint") といったサービスがあるらしい。
同じ Lexicon を持つこれらのサービスでは相互運用が容易になる。

[^l2]: Lexicon は直訳的には「語彙（集）」だが，ここではデータおよびそのアクセスのための構造（schema）を指す。ちなみに Bluesky の lexicon は bsky.app で，これは Bluesky の App View を兼ねている（筈）。

[Leaflet] に関連して以下の記事を見つけた。

- [The Everything Account](https://www.augment.ink/the-everything-account/)

この記事は以下の問題意識から始まる。

{{< fig-quote type="markdown" title="The Everything Account" link="https://www.augment.ink/the-everything-account/" lang="en" >}}
The problem with our accounts is that they aren't our accounts. They're accounts owned by Google, Instagram, Reddit, and X. And when any of these organizations wants to go rogue, all you can do is say "wow, that sucks," and keep using it because starting from scratch with another service is hard.

And what makes it hard is simple: they have our data, including our posts, our relationships, and so much more, and we can't take it all elsewhere. The data you've created in an app is directly tied to the app itself, and the companies behind it hoard the data because they know it's what gives them the power to hold on to you.
{{< /fig-quote >}}

「私のアカウント」「私のデータ」だと思っているものは実際にはアプリと不可分で，サービス・プロバイダが所有し囲っており，アカウントおよびそれに紐づくデータがサービス間で共有・移動することはあまりない。
プロバイダの気が向けばデータを共有するための API を提供したりインポート・エクスポート機能を提供したりすることもあるが，それが他サービスとマッチするとは限らない。

中国の WeChat は “The Everything App” (万能アプリ) と呼ばれる仕組みを持っている。
あるひとつのアプリのアカウントに対して多様なミニアプリがぶら下がり，それらの間でシームレスにデータを共有できる仕組みらしい。

{{< fig-quote type="markdown" title="The Everything Account" link="https://www.augment.ink/the-everything-account/" lang="en" >}}
here's a common concept known as "The Everything App": a single app that contains a series of sub-apps, all under a single account. An example is the popular Chinese app [WeChat](https://en.wikipedia.org/wiki/WeChat), which offers messaging, payments, a full social network, and mini-apps that other companies can build to live inside it.

This simplifies the model we're used to by not requiring you to create a new account for every app. The Everything App stores your data in a shared location within it, and other apps can use it as long as you give them permission.
{{< /fig-quote >}}

確かにこの方法であれば “The Everything App” の裁量下で多様なサービスとアカウントとデータを共有できる。
ただし，それでも “The Everything App” のアカウント自体は「わたしのもの」ではなく “The Everything App” を所有するプロバイダのものである。

では本当の意味で「わたしのアカウント」「わたしのデータ」を実現するにはどうすればいいのか。
これを解決する（解決できそう）なのが [atproto] である。
先に挙げた [atproto] の構成をもう一度見返してみよう。

「わたしのアカウント」「わたしのデータ」を格納している PDS はアプリから独立しており， Bluesky を含む各アプリは API を通じて PDS にアクセスする。
これが “The Everything Account” (万能アカウント) である。
そして “The Everything Account” を中心にしたエコシステムを “The Atmosphere” と呼ぶ。

{{< fig-quote type="markdown" title="The Everything Account" link="https://www.augment.ink/the-everything-account/" lang="en" >}}

{{< fig-img src="./Screenshot_20260204_121413_Samsung-Notes.jpg" width="1000" lang="en">}}

The Atmosphere, as you can see above, already has countless apps you can use. There are Twitter-like apps such as [Bluesky](https://bsky.app/) and [Blacksky](https://blacksky.community/); blogging services like [Leaflet](https://leaflet.pub/), [Offprint](https://offprint.app/), and [pckt](https://pckt.blog/); collection and annotation tools like [Semble](https://semble.so/), [Margin](http://margin.at/), and [Seams](https://seams.so/); and I can go on and on because the ecosystem is expanding by the day. And this is just a small portion of the existing Atmosphere - I couldn't fit all of the different apps because there are just. so. many. Perhaps in a separate post.
{{< /fig-quote >}}

不勉強ながら [atproto] を中心としたいわゆる ATmosphere が既にこれほどの多様性を産んでいることを初めて知った。
真面目に [atproto] の勉強をしないといけないなぁ。

以上の話を併せると [atproto] の ATmosphere が実装よりの話で “The Atmosphere” はもう少し概念よりの話ということになるだろうか。
実際には稼働し利用されている PDS のほとんどは Bluesky が提供している `bsky.social`[^p1] らしいのでまだ理想の状況とは言えないかもしれないが PDS 自体は自前で立てることも可能らしい。
[Hostdon](https://hostdon.jp/ "Hostdon | 分散SNSホスティングサービス") でお一人様サーバを立てるみたいに，簡単に PDS サーバを立てられるようになると面白いんだけどねぇ。

[^p1]: 厳密には Bluesky の PDS は多数のサーバで構成されているそうだが，ユーザから見た PDS へのアクセスは `bsky.social` で[統合されている](https://docs.bsky.app/docs/advanced-guides/entryway "PDS Entryway | BlueskyPDS Entryway | Bluesky")。

{{< fig-quote type="markdown" title="The Everything Account" link="https://www.augment.ink/the-everything-account/" lang="en" >}}
The Atmosphere is an ecosystem that respects your agency as a user, and it's one of many ways we can start taking back control of our online experiences. Whether you're a user or a builder, you no longer need to hold onto hope that a giant company does the right thing for you. You can take your Everything Account across The Atmosphere without the permission of any other entity.

[...]

Your Everything Account – your account on The Atmosphere – is your account. And it's about damn time we had an account that actually is.
{{< /fig-quote >}}

## ブックマーク

- [Leaflet Lab Notes](https://lab.leaflet.pub/)
- [[2402.03239] Bluesky and the AT Protocol: Usable Decentralized Social Media](https://arxiv.org/abs/2402.03239)
- [AT Protocol - Wikipedia](https://ja.wikipedia.org/wiki/AT_Protocol)

- [ATProtocol上でブログ風長文投稿を行うLeaflet – あたしンちのおとうさんの独り言](https://atasinti.chu.jp/dad3/archives/86927)
  - [続・ATProtocol上でブログ風長文投稿を行うLeaflet – あたしンちのおとうさんの独り言](https://atasinti.chu.jp/dad3/archives/86992)

[Leaflet]: https://leaflet.pub/ "Leaflet"
[atproto]: https://atproto.com/ "The AT Protocol"
[standard.site]: https://standard.site/ "Standard.site - One schema. Every platform."
