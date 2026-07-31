+++
title = "ぞーぺん（ZonePane）で bsky.social 以外の PDS にサインインする"
date =  "2026-05-08T20:56:43+09:00"
description = "bsky.social 以外の PDS でサインインを試す / アプリケーションとデータストレージの抱き合わせはリスクか？"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "tools", "zonepane", "bluesky", "atproto", "authentication", "privacy", "web" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

以前書いた「[そろそろ PDS を個人で管理する時代か？]({{< ref "/remark/2026/03/pds-for-everyone.md" >}})」で「常用している[ぞーぺん（ZonePane）][ZonePane]も `bsky.social` 以外の PDS に対応してない」と愚痴ったのだが [ZonePane for Desktop](https://zonepane.com/zonepane_beta.html) v2.9.18 で `bsky.social` 以外の PDS (Personal Data Server) にもサインインできるようになった[^zpd1]。
ありがたや。
さっそく試してみよう。

[^zpd1]: [デスクトップ版 ZonePane](https://zonepane.com/zonepane_beta.html) は現在「お試し版」の状態で Windows, macOS, Linux (Debian, Ubuntu) で利用可能。

## bsky.social 以外の PDS でサインインを試す

アカウント追加画面に移動して

{{< fig-img-quote src="./zonepane-add-account.png" title="アカウント追加" link="./zonepane-add-account.png" width="796" >}}

「Bluesky にサインイン」を選択する。

{{< fig-img-quote src="./zonepane-sign-in-1.png" title="サインイン" link="./zonepane-sign-in-1.png" width="795" >}}

折りたたまれた状態ではあるが「サーバー（PDS）」の項目が追加されている。
Bluesky の既定 PDS である `bsky.social` であればこの項目は空欄でもいいが，それ以外の PDS のアカウントの場合は PDS のドメインを入力する必要がある。

{{< fig-img-quote src="./zonepane-sign-in-2.png" title="サインイン" link="./zonepane-sign-in-2.png" width="795" >}}

これで「サインイン」ボタンを押せばブラウザの OAuth 認証・承認画面に飛ぶ。

{{< fig-img-quote src="./oauth.png" title="Authorize" link="./oauth.png" width="1088" >}}

ここで承認すればサインイン完了である。

Bluesky じゃないほうの PDS を利用している方はプロバイダから PDS 名のアナウンスがあるはずなので必ずメモしておくこと。
紛失しちゃった場合は [PDSls] で確認する手もある。

{{< fig-img-quote src="./pdsls.png" title="photos.baldanders.info - PDSls" link="https://pdsls.dev/at://did:plc:4kzexuq4xieuarc3q6lmkeou#identity" width="515" >}}

デスクトップ版 [ZonePane] の Mastodon やら Bluesky やらをゴチャまぜにしてマルチカラムで表示するインタフェースはホンマに重宝していて，もう手放せなくなっている。
これからもよろしくお願いします。

{{< div-box type="markdown" >}}
v2.9.18 は `bsky.social` 以外の PDS で OAuth 認証した場合に再認証で PDS ドメイン名を忘れてしまう不具合があったが 2026-05-08 にリリースされた v3.0.0 では修正されていた。

ありがたや {{% emoji "ペコン" %}}
{{< /div-box >}}

## アプリケーションとデータストレージの抱き合わせはリスクか？

今までは，ユーザから見てアプリケーションとユーザデータ（のストレージ）が一体であることは当たり前だった。
分散型と言われる Mastodon でさえストレージはアプリケーション・インスタンスの一部になっていて，別のインスタンスに移りたければインスタンス間でデータの「引っ越し」が必要である（引っ越しができるだけマシだが）。

ここで [AT Protocol][atproto] が登場し [Eurosky] や [Periwinkle] といったアカウントとデータを専ら扱う PDS サービスが出てきたことで（ユーザから見た）アプリケーションとユーザデータの関係が変わってきているんじゃないだろうか。
たとえば Bluesky は利用したいが Bluesky が擁する PDS に自分のアカウントとデータは置きたくない，といったような。

- [Blueskyが生み出した、世界一奇妙で意味不明な「拘束力ある仲裁条項」 » p2ptk[.]org](https://p2ptk.org/digital-rights/5598)
- [Twitterとの決別――マスクはカエルを火炎放射器で丸焼きにした » p2ptk[.]org](https://p2ptk.org/digital-rights/5452)

（上の記事によると，コリイ・ドクトロウ氏は PDS を自前で用意したようだ）

私が [pckt.blog][pckt] サービスのアカウント（`photos.baldanders.info`）を取得するときには深く考えず「ひとつくらい別 PDS のアカウントがあってもいいだろう」くらいの軽い気持ちだったのだが，最近 [atproto] エコシステムにおいてアプリケーションとデータストレージの抱き合わせは（アプリケーションがなんであれ）リスクなんじゃないかと思い始めている。

というわけで PDS を引っ越そうか悩み中である。

## 【2026-05-10 追記】

Andorid 版および iOS 版の [ZonePane] も最新版で `bsky.social` 以外へのサインインに対応したようだ。

<fig-quote class="nobox">
<iframe src="https://fedibird.com/@zonepane/116543569773674491/embed" class="mastodon-embed" style="max-width: 100%; border: 0" width="400" allowfullscreen="allowfullscreen"></iframe><script src="https://fedibird.com/embed.js" async="async"></script>
</fig-quote>

手元のスマホで試したが問題なくサインインできた。

[ZonePane]: https://twitpane.com/ "ZonePane - マルチSNSクライアント"
[PDSls]: https://pdsls.dev/ "PDSls"
[atproto]: https://atproto.com/ "The AT Protocol"
[Eurosky]: https://eurosky.tech/ "Eurosky - Building a thriving open social web for Europe"
[Periwinkle]: https://periwinkle.social/ "Periwinkle - Get Started"
[pckt]: https://pckt.blog/ "pckt.blog - pckt"
