+++
title = "Pckt に新しい PDS レジストリを作る"
date =  "2026-03-04T22:19:00+09:00"
description = "PDS レジストリの作成 / アカウント情報 / ハンドル名を独自ドメイン（サブドメイン）に変更する / なにか記事を書いてみる"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "atproto", "bluesky", "blog" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

[AT Protocol][atproto] (以降 [atproto]) ベースのブログサービスに [pckt] ([@pckt.blog](https://bsky.app/profile/pckt.blog)) というのがある。
先日[紹介]({{< ref "/remark/2026/02/leaflet-and-the-atmosphere.md" >}} "Leaflet と The Atmosphere")した [Leaflet] と同じく [standard.site] Lexicon に基づいた構成になっている。

{{< fig-img-quote src="./pckt.blog.png" title="pckt.blog" link="https://pckt.blog/" width="1202" lang="en">}}

この [pckt] が自前の PDS (Personal Data Server) を launch したらしい。

{{< fig-quote type="markdown" title="Building Our Corner of the Open Social Web: pckt.cafe and More" link="https://devlog.pckt.blog/building-our-corner-of-the-open-social-web-pcktcafe-and-more-21wfbg2" lang="en" >}}
When you [sign up](https://pckt.blog/register) for **pckt** blog, you will get the opportunity to claim a username that's part of our **pckt.cafe** personal data server.

That means you will get to choose a handle like **@your-name-here.pckt.cafe** which will be usable to sign in across the atmosphere for apps like [Bluesky](https://bsky.app/), [Germ Network](https://germ.network/), [Stream.place](https://stream.place/), and any place built with AT Protocol.
{{< /fig-quote >}}

ちょっと試してみよう。

## PDS レジストリの作成 {#create}

[Register](https://pckt.blog/register "Register - pckt") ページから作成できる。

{{< fig-img-quote src="./register.png" title="Register - pckt" link="https://pckt.blog/register" width="1118" lang="en">}}

入力したメールアドレスに対して確認用の URL が送られるのでので必ず実行すること。
上手く登録できるとダッシュボード画面が表示される。

{{< fig-img-quote src="./dashboard-1.png" title="Dashboard - pckt.blog" link="https://pckt.blog/dashboard" width="1118" lang="en">}}

とりあえずブログ名を決めてしまおう。
入力したブログ名がサブドメイン名になるので注意。
ブログ名は後で変更できる。

{{< fig-img-quote src="./dashboard-2.png" title="Dashboard - pckt.blog" link="https://pckt.blog/dashboard" width="1118" lang="en">}}

{{< fig-img-quote src="./dashboard-3.png" title="Dashboard - pckt.blog" link="https://pckt.blog/dashboard" width="1118" lang="en">}}

ダッシュボードの右側に注目すると，サインアップ直後では制限が大きいことが分かる。
制限内で試す分には問題ないが，本格的に利用するなら有料版にアップグレードする必要がある。

{{< fig-img-quote src="./subscription.png" title="Subscription - pckt.blog" link="https://pckt.blog/billing/plans" width="1075" lang="en">}}

## アカウント情報 {#account}

ダッシュボード右上のハンバーガーメニューから “Account” を選択すると以下の画面になる。

{{< fig-img-quote src="./settings-1.png" title="Account Settings - pckt.blog" link="https://pckt.blog/profile" width="1202" lang="en">}}

どうやら PDS レジストリを作成すると自動的に Bluesky のアカウントも作成されるようだ。

{{< fig-img-quote src="./bluesky.png" title="Bluesky" link="https://bsky.app/profile/did:plc:4kzexuq4xieuarc3q6lmkeou" width="1202" lang="en">}}

DID の項目から [PDSls] のサイトへ行ける。

{{< fig-img-quote src="./pdsls-1.png" title="PDSls" link="https://pdsls.dev/at://did:plc:4kzexuq4xieuarc3q6lmkeou#identity" width="1202" lang="en">}}

色々と記述があるが，ここでは DID, Aliases, Services の3点のみ確認できればよい。

- **DID:** `did:plc:4kzexuq4xieuarc3q6lmkeou`
- **Aliases:** `at://spiegel.pckt.cafe`
- **Services:** `https://pds.pckt.cafe`

折角なので，ハンドル名 `spiegel.pckt.cafe` を独自ドメイン名 `photos.baldanders.info` に変更してみる。

## ハンドル名を独自ドメイン（サブドメイン）に変更する {#change-handle}

- [Handle - AT Protocol](https://atproto.com/ja/specs/handle)

ハンドル名にドメイン名を割り当てる方法は以下の2つがある。

1. DNS の `TXT` レコードに DID を設定する
2. `https://domain-name/.well-known/atproto-did` に DID を設定する

そもそもユーザが DNS にアクセスできない場合は2番目の方法をとるしかない。
もうひとつ注意しないといけないのは，サブドメインをハンドル名にしたい場合だ。

私は `baldanders.info` ドメイン以下のレコードに対してフル権限を持っているので（ぞんぞがさばる），サブドメインをいくつでも追加できるし `TXT` レコードも自由に追加できる。
ただし 既に `CNAME` で既にサブドメインを Web サイトに割り当てている（例えばここのブログは [GitHub Pages](https://docs.github.com/ja/pages "GitHub Pages のドキュメント - GitHub ドキュメント") を利用して `CNAME` で独自ドメインを割り当てている）場合は1番目の方法では上手くいかないため2番目の方法をとるしかない[^dns1]。

[^dns1]: `CNAME` レコードと（`TXT` を含む）他のレコードとの関係がよく分かってなくて，しばらく悩んでしまった。アホである。 [RFC 1912](https://datatracker.ietf.org/doc/html/rfc1912 "RFC 1912 - Common DNS Operational and Configuration Errors") も参照のこと。

既にあるサイトに `/.well-known/atproto-did` ファイルをセットする手もあったが，今回は `CNAME` レコード設定を避けて `photos.baldanders.info` をハンドル名とした。
こんな感じ。

{{< fig-img src="./txt-record.png" link="./txt-record.png" width="757" lang="en">}}

`_atproto` レコードは `baldanders.info` ドメインに対応し `_atproto.photos` レコードは `photos.baldanders.info` ドメインに対応している。

nslookup コマンドが使えるのであれば，以下のように確認できる[^cf1]。

[^cf1]: コマンドラインにある `1.1.1.1` は [Cloudflare が運営するパブリック DNS リゾルバ](https://www.cloudflare.com/ja-jp/learning/dns/what-is-1.1.1.1/ "1.1.1.1 とは？DNSリゾルバーとは？ | Cloudflare")である。

```text
$ nslookup -q=txt _atproto.photos.baldanders.info 1.1.1.1
Server:		1.1.1.1
Address:	1.1.1.1#53

Non-authoritative answer:
_atproto.photos.baldanders.info	text = "did=did:plc:4kzexuq4xieuarc3q6lmkeou"

Authoritative answers can be found from:
```

`photos.baldanders.info` ドメインをハンドル名として登録するには Bluesky の設定画面を使うとよい。
最初に `spiegel.pckt.cafe` で Bluesky にサインインしておいて Settings → Account → Handle で以下の画面が表示されるので “*Enter the domain you want to use*” の項目に `photos.baldanders.info` を入力する。

{{< fig-img-quote src="./change-handle-1.png" title="Change Handle - Bluesky" link="https://bsky.app/settings/account" width="1202" lang="en">}}

“`Add the following DNS record to your domain:`” の部分は（既に登録済みなので）無視してよい。
最後の “*This should create a domain record at:*” の部分が

```text
_atproto.photos.baldanders.info
```

になっていることを確認して “Verify DNS Record” をクリックする。

{{< fig-img-quote src="./change-handle-2.png" title="Change Handle - Bluesky" link="https://bsky.app/settings/account" width="1202" lang="en">}}

“Domain verified!” と表示されれば DNS レコードが正しく読み込まれている。
あとは “*Update to photos.baldanders.info*” をクリックして完了である（クリックしないと完了しない）。
完了後は新しいハンドル名でサインインし直すこと[^hp1]。

[^hp1]: サインイン時の Hosting Provider はデフォルトの Bluesky Social ではなく `pds.pckt.cafe` を指定すること。ちなみに認証は OAuth である。

[pckt] のほうもサインインし直してね。

{{< fig-img-quote src="./settings-2.png" title="Account Settings - pckt.blog" link="https://pckt.blog/profile" width="1202" lang="en">}}

[PDSls] で PDS の状態を確認してみる。

{{< fig-img-quote src="./pdsls-2.png" title="PDSls" link="https://pdsls.dev/at://did:plc:4kzexuq4xieuarc3q6lmkeou#identity" width="1202" lang="en">}}

最後に Bluesky プロフィール画面を整えておく。

{{< fig-img-quote src="./bluesky-2.png" title="Bluesky" link="https://bsky.app/profile/photos.baldanders.info" width="716" lang="en">}}

## なにか記事を書いてみる {#write}

アカウントとハンドルについてひととおり設定ができたところで，なにか記事を書いてみる。

その前にブログのカスタマイズをしておく。
ダッシュボード右上のハンバーガーメニューから “Customize” を選択する。
今回はこんな感じにした。

{{< fig-img-quote src="./customize.png" title="Customize - pckt.blog" link="https://pckt.blog/blog-settings" width="1202" lang="en">}}

テーマはプリセットされているものから選択できる（自身でカスタマイズも可能）。
フォントもある程度選べる。
ただし日本語の場合はどれを選んでも Serif か Sans-Srif かくらいの違いしかないみたい。

ダッシュボードから “New Post” で記事を書いてみる。

{{< fig-img-quote src="./write.png" link="./write.png" width="1202" lang="en">}}

編集は WYSIWYG (What You See Is What You Get) スタイルで，今どきの markdown 記法とかは使えない。
行頭に `/` を入力するとメニューが出てくる。

タスクリストが書けるのか。
これは嬉しい。

コードも書ける。

{{< fig-img-quote src="./write-code.png" link="./write-code.png" width="811" lang="en">}}

テーブルも書ける。

{{< fig-img-quote src="./write-table.png" link="./write-table.png" width="811" lang="en">}}

文章中の文字列にハイパーリンクを仕込めない。
その代わり URL を書くとリンクカードに変換してくれる。

{{< fig-img-quote src="./write-link-1.png" link="./write-link-1.png" width="811" lang="en">}}

{{< fig-img-quote src="./write-link-2.png" link="./write-link-2.png" width="811" lang="en">}}

数式は無理そう。

文章中の文字列にハイパーリンクを仕込めないとなると，記事の構成を工夫する必要があるかもしれない。
数式が使えないのは... 人によるか。

完成した記事はこんな感じ。

{{< fig-img-quote src="./article.png" title="こんにちは，Pckt - お散歩カメラ - pckt.blog" link="https://osanpo.pckt.blog/pckt-ya0be17" width="1202" lang="en">}}

Bluesky と連動してるくせに publish と同時に Bluesky へのポストしてくれるとかはないらしい。
Bluesky へのポストは当面手動で行うか。

{{< fig-gen >}}
<blockquote class="bluesky-embed" data-bluesky-uri="at://did:plc:4kzexuq4xieuarc3q6lmkeou/app.bsky.feed.post/3mg7jqytars2l" data-bluesky-cid="bafyreia4d3pxndowr27yvkju2nspxgua3oh7roelktlyw6roiac2rkre3q" data-bluesky-embed-color-mode="system"><p lang="ja">こんにちは，Pckt - お散歩カメラ
osanpo.pckt.blog/pckt-ya0be17<br><br><a href="https://bsky.app/profile/did:plc:4kzexuq4xieuarc3q6lmkeou/post/3mg7jqytars2l?ref_src=embed">[image or embed]</a></p>&mdash; Spiegel with お散歩カメラ (<a href="https://bsky.app/profile/did:plc:4kzexuq4xieuarc3q6lmkeou?ref_src=embed">@photos.baldanders.info</a>) <a href="https://bsky.app/profile/did:plc:4kzexuq4xieuarc3q6lmkeou/post/3mg7jqytars2l?ref_src=embed">March 4, 2026 at 2:17 PM</a></blockquote><script async src="https://embed.bsky.app/static/embed.js" charset="utf-8"></script>
{{< /fig-gen >}}

今日はこの辺で勘弁してやろう（笑）

[atproto]: https://atproto.com/ "The AT Protocol"
[pckt]: https://pckt.blog/ "pckt.blog - pckt"
[Leaflet]: https://leaflet.pub/ "Leaflet"
[standard.site]: https://standard.site/ "Standard.site - One schema. Every platform."
[PDSls]: https://pdsls.dev/ "PDSls"
