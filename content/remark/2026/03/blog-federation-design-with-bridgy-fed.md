+++
title = "ブログを Fediverse & Atmosphere に参加させる"
date =  "2026-03-23T20:42:41+09:00"
description = "とりあえず，このブログを Fediverse & Atmosphere に参加させるところまでやってみた。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "activitypub", "atproto", "mastodon", "bluesky", "web", "site" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = true
  jsx = false
+++

随分前に [Bridgy Fed] を利用して Bluesky と Mastodon を相互接続させる方法を[紹介]({{< ref "/remark/2024/10/bridgy-fed-for-bluesky.md" >}} "Fediverse と Bluesky を相互接続する")したが，同じサービスを使ってブログなどの Web サイト（の更新）を Bluesky/Atmosphere および Mastodon/Fediverse と連携させることが出来るようだ。

{{< fig-mermaid title="Federation with Bridgy Fed" >}}
graph LR
  Website["Web site"]
  BridgyFed(("Bridgy Fed"))
  Mastodon["Mastodon (Fediverse)"]
  Bluesky["Bluesky (Atmosphere)"]

  Website-->BridgyFed
  BridgyFed-->Mastodon
  BridgyFed-->Bluesky
{{< /fig-mermaid >}}

登録自体は簡単で，以下のページでブログの URL を入力するだけ。

{{< fig-img-quote src="./web-site-1.png" title="Enter a web site to bridge:" link="https://fed.brid.gy/web-site" width="1175" lang="en" >}}

成功すれば以下の画面に遷移する。

{{< fig-img-quote src="./web-site-2.png" title="text.baldanders.info profile - Bridgy Fed" link="https://fed.brid.gy/web/text.baldanders.info" width="1175" lang="en" >}}

RSS フィードを備えているサイトであれば問題なく登録できるはず。
サイトの登録は誰でもどのサイトでもできる。
勝手に登録されてしまい取り消したいのであれば，[オプトアウト](https://fed.brid.gy/docs#opt-out)の手続きを行う必要がある。

{{< fig-quote type="markdown" title="How do I opt out and remove my site or account?" link="https://fed.brid.gy/docs#opt-out" lang="en" >}}
If you're on the web, [email us](mailto:feedback@brid.gy) from an address at your web site's domain to show that you own it, or you can put the text `#nobridge` in the [profile on your home page](https://fed.brid.gy/docs#web-profile) and then [update your profile](https://fed.brid.gy/docs#update-profile) on [your user page](https://fed.brid.gy/docs#user-page).
{{< /fig-quote >}}

初期状態では Bluesky ハンドルは `@yourdomain.com.web.brid.gy` に， Mastodon ハンドルは `@yourdomain@web.brid.gy` になっている。

たとえばサイトのドメインが `text.baldanders.info` であれば，それぞれ

- `@text.baldanders.info.web.brid.gy` (Bluesky)
- `@text.baldanders.info@web.brid.gy` (Mastodon)

となる。

このうち Bluesky ハンドルはサイトのドメイン名に変更可能である。
変更方法は[以前書いた記事]({{< relref "./creating-a-new-pds-registry-on-pckt.md#change-handle" >}} "ハンドル名を独自ドメイン（サブドメイン）に変更する")を参照のこと。
DNS の TXT レコードまたは `/.well-known/atproto-did` ファイルに DID を設置すれば勝手に更新してくれるみたい。

ちなみに，ここのブログのように GitHub Pages で独自ドメインにしている場合は DNS の TXT レコードが使えないため `/.well-known/atproto-did` ファイルに DID を書いて設置した。
この場合 `/_config.yml` ファイルに以下の記述を追加する必要がある。

```yaml
include: [".well-known"]
```

最終的には以下のハンドル名になった。

- [`@text.baldanders.info`](https://bsky.app/profile/text.baldanders.info)

{{< fig-img-quote src="./pdsls.png" title="text.baldanders.info - PDSls" link="https://pdsls.dev/at://did:plc:crwrada4cit2ijjxuwz4bsjc#identity" width="549" lang="en" >}}

Mastodon ハンドル名の `web.brid.gy` の部分を[独自ドメインに変える](https://fed.brid.gy/docs#fediverse-enhanced "Can I use my own domain as my fediverse handle?")こともできるようだが，このブログではメリットが薄い（かえってハンドル名が長くなる）のでやらないかな。

Fediverse の各サービスや Bluesky 等とやり取りするために [IndieWeb](https://indieweb.org/) や [Webmention](https://webmention.net/) が推奨されているみたいだが，これは後日調査して可能なら対応するか。

今回はここまで。

## ブックマーク

- [Using Web Mentions in a static site (Hugo) | Modern Web Development with Chrome](https://paul.kinlan.me/using-web-mentions-in-a-static-sitehugo/)
- [このブログがFediverseに対応しました](https://blog.tyage.net/post/2023/2023-07-17-bridgy-fed/)
- [Bridgy-Fedのまとめ](https://ubanis.com/note/bridgy_fed_summary/)
- [Webmention.io](https://webmention.io/)
- [Long-Form Comes To Bridgy Fed](https://blog.anew.social/long-form-comes-to-bridgy-fed/) : Bridgy Fed が長文ドキュメント用の [Standard.site](https://standard.site/ "Standard.site - One schema. Every platform.") に対応

[Bridgy Fed]: https://fed.brid.gy/ "Bridgy Fed"
[Kagi Search]: https://kagi.com/ "Kagi Search - A Premium Search Engine"
[Kagi Translate]: https://translate.kagi.com/ "Kagi Translate"
[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"
