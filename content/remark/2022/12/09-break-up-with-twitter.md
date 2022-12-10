+++
title = "5年後に Mastodon に出戻る（Advent Calendar）"
date =  "2022-12-09T00:00:02+09:00"
description = "もう少し様子見かなぁ，と思っています。"
image = "/images/attention/kitten.jpg"
tags = [ "mastodon" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

（本ページは「[Fediverse (2) Advent Calendar 2022](https://adventar.org/calendars/7719 "Fediverse (2) Advent Calendar 2022 - Adventar")」9日目の記事です）

はじめましての方ははじめまして。
Spiegel と申します。
簡単なプロフィールは[こちら](https://baldanders.info/profile/ "公開履歴書 | Baldanders.info")。

実は私は，かつて Mastodon が日本で台頭しはじめた2017年に「[GW 過ぎたらみんな忘れてるに100カノッサ]({{< ref "/remark/2017/04/mastodon.md" >}})」と書いて賭けに惨敗した者です。
その辺を割り引いて読んでいただければ幸いです。

----

正直に言って Twitter の「お家騒動」がここまで騒ぎになるとは思いませんでした。
海の向こうではアンパンマンのごとくオーナーや CEO の頭が挿げ替えられることは珍しくなく，またそのことによって逆に業績が悪化する事態も珍しいことではないため，完全に舐めてかかってました。

しかし Signal が

{{< fig-gen lang="en" >}}
<blockquote class="twitter-tweet"><p lang="en" dir="ltr">Hello, Mastodon - signalapp@mastodon.world</p>&mdash; Signal (@signalapp) <a href="https://twitter.com/signalapp/status/1593678164319997953?ref_src=twsrc%5Etfw">November 18, 2022</a></blockquote>
{{< /fig-gen >}}

と tweet したのにビビッて，私もほぼ5年間放置していた mstdn.jp のアカウントを発掘し復活させました。
あとで気が付いたのですが [EFF](https://mastodon.social/@eff) も [FSF](https://hostux.social/@fsf) も以前から Mastodon アカウントを持ってるんですね。

以後の自ブログの記事はこんな感じです。

- [Mastodon の復活]({{< ref "/remark/2022/11/the-return-of-mastodon.md" >}})
- [もうちょこっと Mastodon]({{< ref "/remark/2022/11/some-little-more-mastodon.md" >}})
- [“Gopher at Mastodon”]({{< ref "/remark/2022/11/gopher-at-mastodon.md" >}})
- [「プラットフォームとフェデレーションとの競争」]({{< ref "/remark/2022/11/new-spark-of-competition-between-platforms-and-federations.md" >}})
- [Prismo はじめました，他]({{< ref "/remark/2022/11/prismo-fedibird.md" >}})

5年経った Mastodon 界隈を眺めて驚いたのは [ActivityPub] を要として Mastodon 以外のサービスとも緩く繋がっている状況（いわゆる Fediverse）でした。
この時点でようやく「Mastodon はもはや Twitter クローンではない」という認識に変わります。
むしろ文章，写真・絵画，動画といったメディア・サービスを（包摂や併呑ではなく）自律[^ai1] 的に繋ぐハブとして機能しているのが Mastodon であると気付きました。
たった5年で気分は「浦島太郎」ですよ。

[^ai1]: 『[そろそろ、人工知能の真実を話そう](https://www.amazon.co.jp/dp/B071FHBGW8?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』によると「自律」というのは元々哲学用語で「自らが行動する際の基準と目的を明確を持ち、自ら規範を作り出すことができることをいう」のだそうです。機械的な「自立」ではないということですね。

この状況では Flickr が「[うちも参画したほうがええんとちゃうん？](https://twitter.com/DonMacAskill/status/1594945727255699457)」と検討し始めるのも納得です。
むしろ Tumblr や Flickr といった老舗サービスが遅まきながらでも [ActivityPub] 連携を検討し始めたことはいいことで，それだけでも某マスク氏が Twitter を買い取った意義があったというものです（笑）

更に更に EFF (Electronic Frontier Foundation; 電子フロンティア財団) までが Mastodon に関連する記事を公開しています。

- [Leaving Twitter's Walled Garden | Electronic Frontier Foundation](https://www.eff.org/deeplinks/2022/11/leaving-twitters-walled-garden)
  - [Twitterの“囲い込み”から離脱するということ――あるいは Fediverse、Federation、Mastodonとは何か？ | p2ptk[.]org](https://p2ptk.org/freedom-of-speech/4167) : 翻訳記事
- [The Fediverse Could Be Awesome (If We Don’t Screw It Up) | Electronic Frontier Foundation](https://www.eff.org/deeplinks/2022/11/fediverse-could-be-awesome-if-we-dont-screw-it)
  - [フェディバースは（“私たち”次第で）素晴らしいものになる | p2ptk[.]org](https://p2ptk.org/freedom-of-speech/4165) : 翻訳記事
- [Is Mastodon Private and Secure? Let’s Take a Look | Electronic Frontier Foundation](https://www.eff.org/deeplinks/2022/11/mastodon-private-and-secure-lets-take-look)
  - [Mastodonのプライバシー、セキュリティ、モデレーションってどんな感じなの？ | p2ptk[.]org](https://p2ptk.org/privacy/4194) : 翻訳記事
- [How to Make a Mastodon Account and Join the Fediverse | Electronic Frontier Foundation](https://www.eff.org/deeplinks/2022/12/how-make-mastodon-account-and-join-fediverse)
  - [Mastodonアカウントの作成とフェディバースへの参加方法 | p2ptk[.]org](https://p2ptk.org/freedom-of-speech/4173) : 翻訳記事

それだけの社会現象になっているということでしょう（今更）

Mastodon を含む [ActivityPub] 連合がどうなるかは分かりません（なんせ前回は賭けで大負けしてますので）。
そのうち沈静化して再びコアな方々だけで続いていくかもしれませんし， [ActivityPub] をも超える新たな展開を迎えるかもしれません。
私自身は（少なくとも Flickr が去就を決めるまでは）もう少し様子を見ながら情報収集を中心に活動していく所存です。

[ActivityPub]: https://www.w3.org/TR/activitypub/

## 参考文献

{{% review-paapi "B0B54Z1LQ1" %}} <!-- シャーロック・ホームズの復活 -->
