+++
title = "eBird と Merlin"
date =  "2023-05-01T21:29:43+09:00"
description = "野鳥識別アプリを使ってみた。"
image = "/remark/2023/04/30-osanpo-camera/52856884852_d3f335c728_o.jpg"
tags = [ "tools", "photography" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

私は特にバードウォッチャーというわけではないが，[昨年末に望遠専用のコンパクトデジカメを買って]({{< ref "/remark/2022/12/autumn-color-in-matsue-castle.md" >}} "Black Friday の戦利品")以来，その辺にいる野鳥を撮って遊んでたりする。

私の場合，生き物の外見を「見分ける」のが不得手で人間の顔もなかなか覚えられないのだが，せっかく写真に撮ったんだし，それがどんな鳥なのか知りたいじゃない？ 今までは [Google Lens](https://lens.google/) とかで調べてたんだけど， [Merlin] というスマホアプリがあるらしい。

- [Merlin Bird ID – Free, instant bird identification help and guide for thousands of birds – Identify the birds you see](https://merlin.allaboutbirds.org/)

{{< fig-youtube id="K0qPBjLcIPk" title="アプリ「Merlin野鳥識別」の使い方（Explore Merlin Bird ID App） - YouTube" >}}

このアプリの面白いところは，単に野鳥の種類が調べられるだけでなく， [eBird] と連動している点のようだ。

{{< fig-img-quote src="./eBird-pic.png" title="eBirdについて - eBird Japan" link="https://ebird.org/japan/about/" width="3420" >}}

[Merlin] で野鳥を調べるには，本体とは別に「バードパック」というのが必要。
最初に起動した際にリージョンに合ったバードパックを勧めてくるので，通常はそのままダウンロード&インストールすればよい。
のだが，データ量が多く，けっこう時間がかかった。

実際に撮った写真で試してみよう。
たとえば，昨日撮った

{{< fig-img src="/remark/2023/04/30-osanpo-camera/52856884852_81a4e36f93_e.jpg" title="今日のバードウォッチング（松江 楽山公園） | Flickr" link="https://www.flickr.com/photos/spiegel/52856884852/">}}

という写真を調べてみる。

{{< fig-img src="./merline-0.jpg" title="Merlin" link="./merline-0.jpg" >}}

最初のメニューに「写真で識別する」という項目があるので，これを使う。

{{< fig-img src="./merline-1.jpg" title="Merlin" link="./merline-1.jpg" >}}

こんな感じで拡大して「次へ」

{{< fig-img src="./merline-2.jpg" title="Merlin" link="./merline-2.jpg" >}}

こんな感じに場所（地図で指定できる）と見つけた日付を入れて「識別」する。
すると

{{< fig-img src="./merline-3.jpg" title="Merlin" link="./merline-3.jpg" >}}

こんな感じに候補が表示される。
これで件の鳥がカルガモだと分かる。

この状態で「これが私の見つけた鳥です！」ボタンを押すと登録されるのだが，登録するためには [Cornell Lab] のアカウントが必要である。

{{< fig-img src="./merline-4.jpg" title="Merlin" link="./merline-4.jpg" >}}

アカウントの作成は Web ブラウザに移動して行う。
住所とか入力させられるので（必須ではないみたいだが），そういうのが嫌な方は登録を諦めるしかない。
なお [Web でサインイン](https://secure.birds.cornell.edu/cassso)して後からアカウント設定変更も可能で，削除も簡単にできる。

アカウントを作成して登録すれば自身のライフリストに記録が残る。
ただし識別に使った写真は登録されないみたい。
あくまで日付と場所と目撃した鳥の種類（推定）が非公開の個人記録として登録されるだけのようだ（学術研究にも使われない）。
[Merlin] で行った識別を [eBird] 側に登録するには [eBird Mobile](https://ebird.org/about/ebird-mobile/ "eBird Mobile - eBird") というアプリを入れて連携させる必要があるらしい。
[eBird] メインで活動するなら，もう立派にバードウォッチャーかな。

んー。
最初に言ったように私はバードウォッチャーではないので，今のところ [Merlin] で遊ぶくらいでちょうどいいだろう。

## ブックマーク

- [eBird Japan - バードウォッチングの世界が広がる…](https://ebird.org/japan/home)
  - [キャンペーン「愛鳥週間！ 身近な鳥を観察してeBirdに投稿しよう」 ～スマホで参加する、市民科学プロジェクト～ - eBird Japan](https://ebird.org/japan/news/campaign-ebird-20230510) : 5月10〜16日は「愛鳥週間」だそうな
  - [「eBird/Merlin設定と基本操作ガイド」を公開しました - eBird Japan](https://ebird.org/japan/news/setting_guide)
  - [eBird/Merlinで種名を日本語表示にするには - eBird Japan](https://ebird.org/japan/news/specific_name_jp)

[Merlin]: https://merlin.allaboutbirds.org/ "Merlin Bird ID – Free, instant bird identification help and guide for thousands of birds – Identify the birds you see"
[eBird]: https://ebird.org/ "eBird - Discover a new world of birding..."
[Cornell Lab]: https://www.birds.cornell.edu/home/ "Cornell Lab of Ornithology—Home | Birds, Cornell Lab of Ornithology"

## 参考

{{% review-paapi "B08L4WKDZ7" %}} <!-- PowerShot ZOOM -->
{{% review-paapi "B07W7JMHHX" %}} <!-- 「ガッチャマン クラウズ」 Crowds -->
