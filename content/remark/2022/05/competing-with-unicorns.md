+++
title = "『ユニコーン企業のひみつ』読書会（1）"
date =  "2022-05-21T20:28:24+09:00"
description = "アレな GAFA 礼賛本を読むくらいならこっちの方をお勧めする。"
image = "/images/attention/kitten.jpg"
tags = [ "book", "engineering", "management" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

前にどっかで書いたような気がするが，遅まきながらプロジェクト・マネジメントの勉強を始めている。
といっても別にプロマネになりたいわけではなく，単に職場でマネジメントの観点で意見を求められることがあったので，そういう観点で見る目も養おうということで。
つっても PMBOK の基礎的な内容に留まっているが。
身も蓋もない言い方をするなら興味本位（笑）

で，今回『[ユニコーン企業のひみつ]』という本を題材にした読書会を開催するというので，これまた興味本位で参加してみた。

- [第1回『ユニコーン企業のひみつ』オンライン読書会 - connpass](https://technical-book-reading.connpass.com/event/245551/)

この本は知らなくても原著者の Jonathan Rasmusson[^bio1] という名前を聞いたら「あぁ『[アジャイルサムライ](https://www.amazon.co.jp/dp/B00J1XKB6K?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』の人か」と思う人もいるかも知れない。
いや，私は読んでないけど。
ちなみに原書 “[Competing with Unicorns](https://www.amazon.co.jp/dp/B088PBCWBZ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)” は2020年，今回の翻訳本『[ユニコーン企業のひみつ]』は2021年に出版されている。

[^bio1]: Amazon に書かれている著者の経歴によると「世界最大級の革新的なテック企業が世界中にソフトウェアを届けることを支援してきた。エンジニアとしては、Spotifyのインテグレーションを支援した。対象プラットフォームはSony PlayStation、Facebook Messenger、Google Chromecast、iMessage。同様にBMW、Tesla、Fordの自動車にも統合した」とある。『[ユニコーン企業のひみつ]』では Spotify での経験が色濃く反映されている。

## 目次

- [『ユニコーン企業のひみつ』読書会（1）]({{< ref "/remark/2022/05/competing-with-unicorns.md" >}}) ← イマココ
- [『ユニコーン企業のひみつ』読書会（2）]({{< ref "/remark/2022/06/competing-with-unicorns-2.md" >}})

## 「ユニコーン企業」とは

『[ユニコーン企業のひみつ]』の訳注によると，「ユニコーン企業」とは

{{< fig-quote type="markdown" title="『ユニコーン企業のひみつ』まえがき" link="https://www.amazon.co.jp/dp/4873119464?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
2013年にベンチャーキャリタリストの Aileen Lee が「ユニコーン企業」という概念を提唱したときの定義は、創業10年以内、評価額10億ドル以上、非上場、米国拠点のテック企業（当時は39社）。ユニコーンと呼ぶのは、その統計上の希少性を神話的な幻獣になぞらえたもの
{{< /fig-quote >}}

だそうな。
ただ，この本ではもう少し緩い定義で

{{< fig-quote type="markdown" title="『ユニコーン企業のひみつ』まえがき" link="https://www.amazon.co.jp/dp/4873119464?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
本書で「ユニコーン企業」と言うとき、それは評価額が10億ドル規模の企業でありながら、スタートアップみたいに運営されている企業のことだと思ってほしい
{{< /fig-quote >}}

としている。
この定義なら「Google、Amazon、Facebook、Spotify なんかがこれに該当する」（まえがき）と言えるだろう。

なお「スタートアップ」の定義も

{{< fig-quote type="markdown" title="『ユニコーン企業のひみつ』まえがき" link="https://www.amazon.co.jp/dp/4873119464?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
本書においてスタートアップとは、50名未満の小さな企業を想定している
{{< /fig-quote >}}

とだいぶ緩い。

## 「スタートアップはプロダクトがすべて」

スタートアップを特徴づける説明がこれ。

{{< fig-quote type="markdown" title="『ユニコーン企業のひみつ』1.1章" link="https://www.amazon.co.jp/dp/4873119464?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
スタートアップはプロダクトがすべてだ。デモで見せるのはプロダクトだ。新しい顧客をひきつけるのもプロダクトだ。資金調達するのもプロダクトだし、学習するのもプロダクトを通じてだ
{{< /fig-quote >}}

故に「スタートアップはイテレーションを重ねる」（1.2章）ことになる。
これが既存のエンタープライズ企業との最大の違いだろう。
『[ユニコーン企業のひみつ]』ではエンタープライズを「期待に応じる機械」（1.4章）と呼びスタートアップを「学習する機械」（1.2章）と呼んでいる。

「期待に応じる」ためには「プロジェクト駆動」が望ましい。
期待に応じる目標を掲げ，目標を達成するための KGI/KPI を評価しゴールを設定する。
コストを計上し期日を決め作業分割を行いリソース（予算や人員を含む）を確保してスケジュールやリスクの変動を管理（監視）する。
プロジェクトが設定するゴールへ計画通りに到達できたら成功だ。

でもプロダクトを中心に据えるなら「プロジェクト駆動」はうまく行かない。
PDCA や OODA といったサイクルを回すにしても活動がプロジェクト単位で途切れてしまうので，とてもじゃないが「イテレーションを重ねる」ことはできない。
そこでスタートアップはチームに権限と自律性を与える「ミッション」で仕事を進める。

## 「失敗はゲームの一部」

「プロジェクトとミッションってなにがちゃうねん」と思うだろう（実際に読書会でもそういう疑問は出た）。
『[ユニコーン企業のひみつ]』ではプロジェクトとミッションの違いを以下のようにまとめている。

{{< fig-quote class="nobox" type="markdown" title="『ユニコーン企業のひみつ』2.5章" link="https://www.amazon.co.jp/dp/4873119464?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
| プロジェクト                   | ミッション                       |
| ------------------------------ | -------------------------------- |
| 予算がある                     | チームの人数が予算               |
| 終わりがある                   | 期間に定めがない                 |
| 短期間                         | 長期間                           |
| プロジェクトマネージャーがいる | プロジェクトマネージャーがいない |
| 開発だけして引き継ぐ           | 開発もメンテナンスもする         |
| 完成したら解散する             | チームは一緒のまま               |
| 計画にフォーカス               | 顧客にフォーカス                 |
| 期待に応じることが価値         | インパクトが価値                 |
| トップダウン                   | ボトムアップ                     |
{{< /fig-quote >}}

もしプロジェクトからミッションへ切り替えたいならプロダクトに対する「成功」の定義も書き換える必要がある。

{{< fig-quote type="markdown" title="『ユニコーン企業のひみつ』1.5章" link="https://www.amazon.co.jp/dp/4873119464?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
プロダクト開発における成功とは「発見と学習」だ。最初のプロダクトをとにかく早く世に出すのもそのためだ。そしてこれを素早く、何度も何度も繰り返す。失策をおかすこともあるだろう。だが、リリースを重ねるごとにプロダクトは良くなっていく
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="『ユニコーン企業のひみつ』1.5章" link="https://www.amazon.co.jp/dp/4873119464?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
プロダクト開発では失敗はゲームの一部だ。 [...] プロダクト開発は一発勝負じゃない。初回のリリースは始まりに過ぎない
{{< /fig-quote >}}

## プロジェクトマネージャーなどいない

ところで「プロジェクトマネージャーがいない」のにどうやって管理するねん，って思わなかった？ 私は思った。
でもプロジェクトじゃないんだからプロジェクトマネージャーがいないのは当たり前だよな。
その代わり重要なのが「プロダクトマネージャー」と「データサイエンティスト」だ。

データサイエンティストは8章で詳しく説明されるらしいので，とりあえずプロダクトマネージャーについて紹介すると

{{< fig-quote type="markdown" title="『ユニコーン企業のひみつ』3.3章" link="https://www.amazon.co.jp/dp/4873119464?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
プロダクトマネージャーは「このプロダクトは何をすべきか」について、信頼のおける情報源になることでプロダクトのデリバリーを導く。スクワッドと協力して戦略を定義し、ロードマップを策定し、機能の定義を考える。マーケティング、売上予測、損益計算の責任にも何らかの関わりを持つ
{{< /fig-quote >}}

だそうな。
「スクワッドと協力して」というのがポイントだろう。
つまりプロダクトマネージャーはスクワッドの管理者ではないということだ。
故にプロダクトマネージャーはスクワッドに対して「指示」は行わない（アドバイスや提案はあり）。

ちなみに，この本で「スクワッド（Squad）」とは

{{< fig-quote type="markdown" title="『ユニコーン企業のひみつ』3.1章" link="https://www.amazon.co.jp/dp/4873119464?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
少人数で、職能横断（Cross-Function）の、自己組織化されたチームだ。
[...]
自律した小さなチームはテック企業のあらゆる活動の中心だ。新規プロダクトの開発、新たな市場への参入、株式公開の準備。どの場合であっても、テック企業では自律した小さなチームがその中心にいる
{{< /fig-quote >}}

と定義・解説されている[^squad1]。
この定義を見て「公安9課」を連想したのは私だけじゃない筈（と思いたいw）。

[^squad1]: スクワッドはかなり重要な概念なので3章を丸々使って解説されている。 Spotify においてスクワッドがどう機能しているかについて[これ](https://youtu.be/Yvfz4HGtoPc)とか[これ](https://youtu.be/vOt4BbWLWQw)で紹介されている（いずれも15分弱の動画）。

スクワッドが所属する企業に対して自律的に振る舞うためには，お互いに対等な信頼関係が必要だ。
この辺が一番難しいんじゃないかな，と個人的には思う。
そのためのヒントは3.7章に書かれている。

## GAFA 礼賛本を読むくらいなら

というわけで，今回は3章まで読了。
ページ数が少ないので，多分あと2回ほどで終わると思う。
もしこの拙文を見て興味が湧いたなら手にとって読んでみてもいいだろう。

買うなら[版元で扱っている電子版](https://www.oreilly.co.jp/books/9784873119465/ "O'Reilly Japan - ユニコーン企業のひみつ")がお勧め。
「[プログラミング言語Go読書会](https://gpl-reading.connpass.com/)」では紙の本にびっしり書き込みをしているが PDF なら書き込みもデジタルでできる[^pdf1]（笑）

[^pdf1]: Ubuntu に標準で付いている [Evince](https://wiki.gnome.org/action/show/Apps/Evince "Apps/Evince - GNOME Wiki!") は PDF の任意の場所にメモを追加する程度はできる。このくらいの機能で十分だろう。

この本を読んで「お父さん，明日からミッションを掲げてプロダクト駆動で頑張るぞ！」とは行かないかも知れないが，アレな GAFA 礼賛本を読むくらいならこっちの方をお勧めする。

[ユニコーン企業のひみつ]: https://www.amazon.co.jp/dp/4873119464?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "ユニコーン企業のひみつ ―Spotifyで学んだソフトウェアづくりと働き方 | Jonathan Rasmusson, 島田 浩二, 角谷 信太郎 |本 | 通販 - Amazon.co.jp"

## 参考図書

{{% review-paapi "4873119464" %}} <!-- ユニコーン企業のひみつ -->
{{% review-paapi "B01JMEDX8A" %}} <!-- 攻殻機動隊 STAND ALONE COMPLEX (SAC) -->