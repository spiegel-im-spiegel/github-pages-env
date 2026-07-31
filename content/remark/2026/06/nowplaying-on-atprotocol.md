+++
title = "「なうぷれあっと」を試す"
date =  "2026-06-27T20:28:22+09:00"
description = "AT Protocol ベースの音楽共有サービス"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "music", "atproto", "bluesky", "tools" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

[すいばり](https://bsky.app/profile/suibari.com "すいばり (@suibari.com) — Bluesky")さんによる音楽共有サービス「[なうぷれあっと]」を試してみた。
ご本人による紹介記事はこちら

- [日本の音楽シーンをATプロトコルで共有するSNSを作った｜すいばり](https://note.com/suibari/n/n30d0da7aaa13)

[AT Protocol][atproto] ベースの音楽共有サービスとしては [Rocksky] が有名らしいが，あちらのメインボリュームは（当然ながら）英語圏の人達なので，上がってくる楽曲もあちらの音楽が中心になる。
そこで

{{< fig-quote type="markdown" title="日本の音楽シーンをATプロトコルで共有するSNSを作った" link="https://note.com/suibari/n/n30d0da7aaa13" >}}
**日本のユーザーが、日本語の音楽でつながれる場所が欲しかった。**
{{< /fig-quote >}}

というのがサービスを立ち上げた動機だそうな。

うんうん。
よく分かる。
私なんて今はアニソンか VTuber の楽曲しか聴かないため，JPOP ならまだしも，海の向こうの音楽はさっぱり分からない。
なので，こういうサービスは大変ありがたかったり。

というわけで試してみた。
サインアップは Bluesky (というか [atproto]) のアカウントで行う。
設定画面はこんな感じ。

{{< fig-img-quote src="./settings.png" title="設定" link="https://nowplayingat.suibari.com/settings" width="949" >}}

「自動投稿」を有効にして [Last.fm] のユーザ名をセットすれば [Last.fm] の scrobble 情報を使って Bluesky に自動投稿してくれる。
[Last.fm] と [Spotify] などの音楽配信サービスを連携すれば，それらの視聴データも反映される。

私はスマホに [Onkyo HF Player] を[入れていて]({{< ref "/remark/2024/03/high-resolution-audio.md" >}} "スマホで高解像度音楽ファイルを鳴らしてみる")，楽曲の音源を入れてヘヴィローテーションしているのだが，これも [Last.fm] の[アプリ][Last.fm アプリ]で認識できた。

[なうぷれあっと] のデータは各ユーザの PDS リポジトリに公開情報として格納される。
[なうぷれあっと]自体はユーザの視聴データを保持しないが，各 PDS リポジトリのデータをかき集めてトップページに統計情報を表示しているらしい。

{{< fig-quote type="markdown" title="日本の音楽シーンをATプロトコルで共有するSNSを作った" link="https://note.com/suibari/n/n30d0da7aaa13" >}}
ATプロトコルは **分散型アーキテクチャ** です。「再生回数が多い曲」「人気ユーザー」のような集計は、複数のサーバーをまたぐデータを扱う必要があります。

音楽SNSのような中央集権的なサービスが当たり前にやっていることを、分散型の制約の中でどこまで実現できるか試してみたかったです。
{{< /fig-quote >}}

{{< fig-img-quote src="./top.png" title="なうぷれあっと" link="https://nowplayingat.suibari.com/" width="1145" >}}

Bluesky への自動投稿は設定で間引くことができる。
さっき書いたように，私は聴き始めると割とヘヴィローテーションしてしまうので，全部を Bluesky にポストすると TL が `#なうぷれ` で埋め尽くされてしまう。
ポストの頻度についてはチューニング中だが，私は 10% くらいで十分かなぁと思っている。

実は [Last.fm] のことをすっかり忘れていた。
Scrobble の履歴を見たら2016年が最後のアクセスだった。
まだアカウントがあったことにびっくりだよ。
当時の履歴もちゃんと残ってたし。

そういえば[なうぷれあっと]に上がってる楽曲で [Spotify] 等で視聴可能なものについては該当のページに飛べるようになっている。
素晴らしい。
[Spotify] はアカウントだけ取って完全に放置してたんだよな[^s1]。
初めて役に立ったよ（笑）

[^s1]: [Spotify] に限らないけど（アルゴリズムで）機械的に作ったプレイリストやおすすめを聴くのってダルくない？ まぁ，私はアナログレコードや CD の世代で，気に入った楽曲は手元に置いて聴き倒したい人だからしょうがない（←年寄り）。今は FLAC 形式の音源が普通に買えるようになって本当によかった。リッピングしなくて済む。

今回のサービスを試してみて思ったが，ソーシャルなサービスだからといって何でもフラットにすればいいというわけじゃないよな。
グローバルとの繋がりは保ちつつ，やんわりと囲える設計がいいのかもしれない。
そう考えると Mastodon というか ActivityPub はよくできているかも。

[なうぷれあっと]: https://nowplayingat.suibari.com/ "なうぷれあっと"
[atproto]: https://atproto.com/ "The AT Protocol"
[Rocksky]: https://rocksky.app/ "Rocksky"
[Last.fm]: https://www.last.fm/ "Last.fm | Play music, find songs, and discover artists"
[Last.fm アプリ]: https://play.google.com/store/apps/details?id=fm.last.android "Last.fm - Apps on Google Play"
[Onkyo HF Player]: https://play.google.com/store/apps/details?id=com.onkyo.jp.musicplayer "Onkyo HF Player - Apps on Google Play"
[Spotify]: https://open.spotify.com "Spotify - Web Player: Music for everyone"
