+++
title = "今日の戯れ言 2024-10-25"
date =  "2024-10-25T22:39:32+09:00"
description = "欲しいのは Bluesky でないもの？ / Linux Kernel コミュニティへのソーシャル・エンジニアリング？"
image = "/images/attention/kitten.jpg"
tags = [ "bluesky", "activitypub", "security", "risk", "linux" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

前回に引き続き Mastodon に書き散らかした戯れ言からピックアップし再構成してお送りする。

やっぱ TL サービスはログが流れて過去に自分が書いたものを探すのも面倒くさい感じになるので，ちゃんとブログに残しておきたい。
と，いつも思うのだが，特にこの1年半ほどは「書く」意欲が減退してて「お散歩カメラ」の記事ばっかりになっている。
まぁ「書きたくないときには書かない」のがこのブログのポリシーなので（笑）

## 欲しいのは Bluesky でないもの？

Bluesky がシリーズAで15Mドルを調達したらしい。

- [Bluesky Announces Series A to Grow Network of 13M+ Users - Bluesky](https://bsky.social/about/blog/10-24-2024-series-a)

めでたい！

最後に書かれているこの文章が印象的である。

{{< fig-quote type="markdown" title="Bluesky Announces Series A to Grow Network of 13M+ Users" link="https://bsky.social/about/blog/10-24-2024-series-a" lang="en" >}}
Traditional social media companies have enclosed the online commons, locked down their APIs to shut out independent developers, and deployed black box algorithms that leave us guessing. This era of old social is over — at Bluesky, we’re returning choice and power to you.
{{< /fig-quote >}}

あくまで私個人の印象で申し訳ないが Bluesky が {{% emoji "X" %}} (旧Twitter) や Facebook/Instagram など既存の SNS と正面切って比較できるようになるのは，ユーザ数が1億を超えたあたりからじゃないかと思っている。
ぶっちゃけ今の Bluesky は {{% emoji "X" %}} がやらかしたときの一時避難先に過ぎず “Traditional social media companies” へのインパクトは小さい。

あと，これは私が見逃してるかも知れないのでそれならゴメンナサイ {{% emoji "ペコン" %}} だけど，確かに Bluesky は本格的に PDS の分散・連合（federation）を始めたけれど，あくまでも Bluesky というアプリケーションの枠内に納まっている。
この辺は ActivityPub ベースの Fediverse が [Pixelfed](https://pixelfed.org/ "Pixelfed - Decentralized social media") や [PeerTube](https://joinpeertube.org/ "What is PeerTube? | JoinPeerTube") などの多様なアプリケーションを内包しているのとは対照的と言える。

でも本当に欲しいのはアプリケーションの分散化（あるいは非中央集権化）ではなくアプリケーションを横断して「緩く繋がる」仕組みじゃないだろうか。
そして今のところそれを実現できているのは ActivityPub ベースの Fediverse のみ，と思っている。

Bluesky というか [AT Protocol] がそこまでを視野に入れているかどうかは分からない（分かるようになるのはまだ先だろう）。
でも何となく Bluesky がこれよりも発展するためには ([AT Protocol] ベースの) Bluesky でないキラーアプリケーションが登場することなんじゃないのかな，とか思ったり。

## Linux Kernel コミュニティへのソーシャル・エンジニアリング？

- [“ロシアのトロールどもに告ぐ、この変更が元に戻ることはない” ―Linus、ロシア系メンテナーの"追放"を認める | gihyo.jp](https://gihyo.jp/article/2024/10/daily-linux-241025)

この記事だけだとピンとこなかったのだが，記事中のリンク先の以下の記事を併せて読んでちょっと理解できたかな。

- [そのパッチは受け取れない ―ロシアの半導体メーカーに在籍する開発者、カーネルパッチを拒否される | gihyo.jp](https://gihyo.jp/article/2023/03/daily-linux-230329)

この記事は2023年3月の話でロシアの半導体メーカーに所属するエンジニアが寄与したパッチを拒否されたらしい。
戦時下のロシアメーカーってのがポイントね。

それに対して最初に挙げた記事はパッチは受け入れたけどメンテナーを拒否したという話らしい。
Linus Torvalds 氏は拒否した彼らを “Russian trolls” と呼んでいるようだ。

{{< fig-quote type="markdown" title="Re: [PATCH] Revert “MAINTAINERS: Remove some entries due to various compliance requirements.” - Linus Torvalds" link="https://lore.kernel.org/all/CAHk-=whNGNVnYHHSXUAsWds_MoZ-iEgRMQMxZZ0z-jY4uHT+Gg@mail.gmail.com/" lang="en" >}}
Ok, lots of Russian trolls out and about.
{{< /fig-quote >}}

www

今回の措置の妥当性は私には分からないが，思い出したのは今年（2024年）春にあった OSS プロジェクトへの貢献者を装ってバックドアを埋め込んだ事例だ。

- [XZ Utils に仕組まれたバックドアに関する覚え書き]({{< ref "/remark/2024/04/xz-utils-backdoor.md" >}})
- [オープンソース・プロジェクトの乗っ取りを試みる]({{< ref "/remark/2024/04/take-over-opensource-project.md" >}})

この手のソーシャル・エンジニアリングを検出するのはホンマに難しい。
きょうび悪人顔の悪人なんかいないしな。

## ブックマーク

- [Bluesky、サブスクリプションサービスを準備中 ｰ ユーザー数は先月から300万人増の1,300万人超に | 気になる、記になる…](https://taisy0.com/2024/10/25/205353.html)
- [Bluesky、1500万ドル調達　追加機能が使えるサブスクプラン開発中 - ITmedia NEWS](https://www.itmedia.co.jp/news/articles/2410/25/news110.html)

[AT Protocol]: https://atproto.com/ "AT Protocol"

## 参考図書

{{% review-paapi "4296001574" %}} <!-- ハッキング思考 -->
{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
