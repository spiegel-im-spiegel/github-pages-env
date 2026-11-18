+++
title = "ZonePane で Bluesky に長文を投稿する（mochott 連携）"
date =  "2026-11-18T16:35:03+09:00"
description = "ブログとソーシャル・メディアをシームレスに繋ぐというアプローチは，非常に面白い"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "tools", "zonepane", "bluesky", "atproto", "blog", "communication" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

改めて紹介すると [ZonePane]（ぞーぺん）は Mastodon/Misskey/Bluesky/タイッツー/mixi2 に対応したマルチ SNS クライアントである。
Android および iOS/iPad に対応したスマホアプリと Windows/macOS/Linux 向けのデスクトップアプリが提供されている。

この [ZonePane] の v49 が昨日リリースされた。

- [ZonePane v49 をリリースしました](https://note.com/panecraft/n/n5241b20b38e3)

主な変更点は以下の通り（ブログ記事から抜粋）。

- Bluesky: 300文字を超える長文を投稿できるようになった（mochott 対応）
- デスクトップ版（Windows / macOS / Linux）が正式版に昇格
- Mastodon: サーバー側のフィルターをアプリ内で追加・編集できるようになった
- Mastodon: 「トレンド」タブで話題の投稿を表示するようになった
- 投稿画面でアカウントを切り替えても、添付ファイルを引き継げるようになった
- Bluesky: 通知のカテゴリ別タブ・プロフィールリンクの編集などに対応した
- Mastodon・Misskey: 投稿の編集で添付ファイルを再アップロードしないようにした
- タイッツー・mixi2 の改善
- その他の改善・不具合修正

まずはデスクトップ版が（[お試し版]({{< ref "/remark/2026/03/zonepane-desktop.md" >}} "デスクトップ版ぞーぺんに期待！")から）正式版に昇格したとのことで，おめでとうございます。
これで大手を振ってデスクトップ版が使える（いや，今までも普通に使ってたがw）。

以降では [mochott] との連携機能に絞ってメモがてら紹介する。

[mochott] は [Bluesky] 用 Web クライアント [TOKIMEKI] の[作者](https://bsky.app/profile/holybea.blue "ほりべあ (@holybea.blue) — Bluesky")が提供するブログサービスで，元々 [TOKIMEKI] と [mochott] はシームレスに連携しているそうだ[^t1]。
今回の [ZonePane] のバージョンアップでは [TOKIMEKI] と同じように [mochott] とシームレスに連携できるようになった。

[^t1]: "[Start Bluesky with TOKIMEKI](https://portal.tokimeki.blue/ "TOKIMEKI & Friends — Start Bluesky with TOKIMEKI")" によると，他にも TOKIMEKI Diary, TOKIMEKI Poll, tomarigi, TOKIMEKI SPACE といったアプリがあるようで，これらを [TOKIMEKI] の UI で統合する形になってるっぽい。ちなみに [TOKIMEKI] のアンケート投票機能は [ZonePane] でも利用できる。

早速デスクトップ版で試してみよう。
入力はこんな感じ。

{{< fig-img-quote src="./zonepane-1.png" title="ZonePane Desktop v49.0.2 (Linux) (1)" link="./zonepane-1.png" lang="en" >}}

下の方に「長文投稿（mochott）」のマークが表示されているのがお分かりだろうか。
これをポストすると以下のようなダイアログが表示される。

{{< fig-img-quote src="./zonepane-2.png" title="ZonePane Desktop v49.0.2 (Linux) (2)" link="./zonepane-2.png" width="565" lang="en" >}}

ポストした内容は [Bluesky] 公式 Web クライアントではこのように見える。

<blockquote class="bluesky-embed" data-bluesky-uri="at://did:plc:w7fwp2mtlaaffnx42wsc76wt/app.bsky.feed.post/3n2j5q72lbtjo" data-bluesky-cid="bafyreihfj26towuhnnleqxicxaqpq5vznbllnmwi4y77lhopfnp2xhwiji" data-bluesky-embed-color-mode="system"><p lang="ja">「ZonePane v49 リリース！デスクトップ版が正式版に。Bluesky の長文投稿(mochott)にも対応しました」

&gt;ZonePane v49 リリース！デスクトップ版が正式版に。Bluesky の長文投稿(mochott)にも対応しました｜takke
&gt; note.com/panecraft/n/n5241b...

まずはデスクトップ版の正式版おめでとうございます！

今回の目玉は，なんといっても mochott との連携だろう

{{< fig-gen >}}
&gt;mochott は、TOKIMEKI の作者のほりべあさんが作られている、Bluesky 向けの長文投稿サービスです。長文は moc…もっと見る<br><br><a href="https://bsky.app/profile/did:plc:w7fwp2mtlaaffnx42wsc76wt/post/3n2j5q72lbtjo?ref_src=embed">[image or embed]</a></p>&mdash; Spiegel＠ナナシビト (<a href="https://bsky.app/profile/did:plc:w7fwp2mtlaaffnx42wsc76wt?ref_src=embed">@baldanders.info</a>) <a href="https://bsky.app/profile/did:plc:w7fwp2mtlaaffnx42wsc76wt/post/3n2j5q72lbtjo?ref_src=embed">2026年11月17日 22:32</a></blockquote><script async src="https://embed.bsky.app/static/embed.js" charset="utf-8"></script>
{{< /fig-gen >}}

「もっと見る」のリンクから [mochott] 側の[投稿ページ](https://mochott.site/baldanders.info/3n2j5q72lb3jo "「ZonePane v49 リリース！デスクトップ版が正式版に。Bluesky の長文投稿(mochott)にも対応しま… by baldanders.info | mochott")に行ける。

同じ記事を [ZonePane] で表示するとこんな感じ。

{{< fig-img-quote src="./zonepane-3.png" title="ZonePane Desktop v49.0.2 (Linux) (3)" link="./zonepane-3.png" width="811" lang="en" >}}

返信部分は無視してほしい。
「全文を表示」ボタンを押すと [mochott] に投稿した全文が表示される。

{{< fig-img-quote src="./zonepane-4.png" title="ZonePane Desktop v49.0.2 (Linux) (4)" link="./zonepane-4.png" width="811" lang="en" >}}

おそらく PDS にある [mochott] のデータ（[standard.site] Lexicon）を参照しているのだろう。
ちなみに [mochott] 側の記事は後から修正することができるが，この場合 [ZonePane] 側の全文表示では「投稿後に編集されています」と表示される。
素晴らしい！

なお [ZonePane] の目玉である複数サービスへのクロスポスト機能を使う場合には [mochott] 連携が使えない。
ポストするサービスの最短文字数でカウントして，それを超えるとポストできない仕様になっている。
たとえば mixi2 は149.3文字で[^m1]，こちらに合わせる場合はしょうがないけど，一方で Misskey は3K文字書けるし Mastodon も（サーバによるけど）1K文字以上書けるサーバが多い（[うち](https://goark.fedicity.net/ "Goark")は5K文字まで書ける。 [Fedibird](https://fedibird.com/ "Fedibird") は560文字）。

[^m1]: 149.3を逆さから読むと3941（みくしい）となる。要らないダジャレ `orz` まぁ，旧 Twitter は140文字だったし，それよりはマシかも知れないが。昔は如何に140文字以内に情報をおさめるかという，実社会では役に立たなそうなスキルが求められていた。昭和のバブル時代に言われた「名刺の裏にアイデアが書けるか」みたいな話か（笑）

[mochott] との連携で [Bluesky] の文字制限は事実上なくなったので，クロスポストでも対応してほしいけど... いや，実装が面倒くさそうか？

私は Mastodon というか Fediverse のアプリケーション・サービスが ActivityPub で緩く繋がってる状態をとても気に入っているが（私の Mastodon サーバが「ぽつんと一軒家」みたいな状態でも外と繋がっていられるのはこれのおかげ），[atproto] で繋がっているアプリケーション・サービスはユーザアカウントにデータが集約されるため，より密な連携が期待できる。
また [ZonePane] のようにサービス・プラットフォームを跨いで閲覧・投稿できるクライアントアプリは，今までにない景色を見せてくれる。

私は最近，サービス・プラットフォーム内部で多様性を担保するのは無理だし，そこに固執する必要はないと思えるようになってきた。
まぁ {{% emoji "X" %}} のように何かと言いがかりをつけて[気まぐれに BAN する]({{< relref "./04-stories.md#x-ban" >}} "𝕏 アカウントを BAN された")ようなサービスは論外だが，（ひとつのプラットフォームに億単位の莫大なユーザを包摂するのではなく）サービス・プラットフォームを跨いで{{% ruby "peer" %}}横向き{{% /ruby %}}に繋がっていくプロセスこそが多様性であり，World Wide Web が黎明期に掲げた理念に近づける道筋なんじゃないだろうか。

というわけで，ブログとソーシャル・メディアをシームレスに繋ぐというアプローチは，非常に面白いと思っている。
[TOKIMEKI] も含めて，今後の展開に期待したい。

## ブックマーク

- [デスクトップ版ZonePaneが正式リリースされBlueskyの300文字以上の投稿に対応 – あたしンちのおとうさんの独り言](https://atasinti.chu.jp/dad4/archives/96815)
- [Leaflet と The Atmosphere]({{< ref "/remark/2026/02/leaflet-and-the-atmosphere.md" >}})
- [AT Protocol & ActivityPub 関連サービスの利用について]({{< ref "/atproto-based-services.md" >}})

[ZonePane]: https://zonepane.com/ "ZonePane - マルチSNSクライアント"
[mochott]: https://mochott.site/ "mochott"
[TOKIMEKI]: https://tokimeki.blue/ "TOKIMEKI - Bluesky client project"
[Bluesky]: https://bsky.app/ "Bluesky"
[atproto]: https://atproto.com/ "The AT Protocol"
[standard.site]: https://standard.site/ "Standard.site - One schema. Every platform."

## 参考

{{% review-tatsujin "infoshare2" %}} <!-- 続・情報共有の未来 -->
{{< linkcard "80446738630e3d6e2fd656927bc64838731bc1e6" >}} <!-- THIS IS FOR EVERYONE Tim Berners-Lee -->
