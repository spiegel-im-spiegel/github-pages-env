+++
title = "正義は公正とは相容れない"
date =  "2022-12-10T16:27:48+09:00"
description = "プラットフォームは詳細である"
image = "/images/attention/kitten.jpg"
tags = [ "code", "media" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

（胡乱なことを口走ってる自覚はあります。閲覧注意）

正義は公正とは相容れない。
もうちょっと簡単に言うと「英雄は裁かれない」ってやつだ。

このフレーズが浮かんだのは [Fediverse Advent Calendar 2022](https://adventar.org/calendars/7371) 8日目の記事「[個人情報保護法令とFediverseの話](https://www.blessedgeeks.com/~h12o/datalaws-and-fediverse.html)」の最後の方に出てくる以下の記事を読んだとき。

- [データやアルゴリズム、そしてユーザーにどう向き合うべきか (あるいはTwitterのホーム表示と時系列表示について) - フジイユウジ::ドットネット](https://fujii-yuji.net/2022/12/06/200337)

この記事は

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">俺ら(おまえら)ヘビーユーザーの方々は、タイムラインの順番が変わったり人のイイネが勝手に入るのいらねぇと思うでしょ？自分もそう思うよ。でもどんだけA/Bテストしてもそっちの方がユーザーの定着率が高い。直感と、直感に反するデータとをどうやって消化していくか…</p>&mdash; nabokov♹ (@nabokov7) <a href="https://twitter.com/nabokov7/status/1589107052282277888?ref_src=twsrc%5Etfw">November 6, 2022</a></blockquote>
{{< /fig-gen >}}

という tweet を起点に考察したものだ。
とはいえ，この tweet 自体は個人的には納得できる。
むしろ Twitter 側が計測もせずに主観でタイムラインを弄ってた，なんて話だったら，そっちのほうが拙い。

それでも Twitter はマシな方で，タイムライン右上のボタン

{{< fig-img-quote src="./timeline-in-twitter.png" link="./timeline-in-twitter.png" width="591" >}}

をクリックすれば「ホーム」と「最新ツイート」を切り換えることができる[^t1]。
しかも永続的に効く。
Facebook も「最新」にできるけど永続的に効かないので，画面を開くたびに切り換えないといけないんだよなぁ。
Instagram はアカウントを捨てたし TikTok はそもそもアカウントを取ってないので知らない（笑）

[^t1]: Twitter は，少し前に「ホーム」と「最新」を[スワイプで切り換える](https://forest.watch.impress.co.jp/docs/news/1394582.html "時系列ツイート表示にした「Twitter」にスワイプで［ホーム］と［最新］で切り替える機能 - 窓の杜")ことができるようスマホアプリの機能を変更したが不評だったらしく，元に戻している。

Twitter にとって主たる「顧客」は株主[^s1] と広告主だ。
さんざん言われていることだが，ユーザは「商品」でしかない。
商品を陳列棚（＝タイムライン）にどう並べらたら利益が上がるか計測（＝監視）するのは営利企業として正しいし，計測結果に従って商品を並べ替え，さらに法を根拠に取り下げる（＝検閲）のも正しい。
フォローもしてない広告アカウントが頻繁にプロモーションツイートを差し込むのも正しいわけだ。
忌々しいことに。

[^s1]: 最近まで Twitter は上場企業だった。某マスク氏が買い取って[上場廃止](https://www.itmedia.co.jp/news/articles/2210/29/news048.html "Twitterのマスク氏による買収完了、上場廃止は11月8日（公式） - ITmedia NEWS")したけど。

このことから分かるように，覇権プラットフォームに於いて私達は「量」として評価され，ひと括りの「ユーザ」として包摂されている。
ユーザを個人として見ているわけではないから個人に対する公正さは考慮されない。
そして「Twitter の正義」からこぼれ落ちる個人は排除される。

包摂と排除はコインの表裏の関係である。
排除されていることに我慢ができないのなら，排除されない別の包摂基準を持つプラットフォームへ移住するか，さもなくば自身がプラットフォーマー（笑）になるか。
国際移民問題と同根だな，これ。

私自身 Twitter は「広告媒体」であると認識しており，情報収集の手段として重宝してはいるものの，仮に某マスク氏が突然「もう閉店ガラガラ」とか言ってシャットダウンしたとしても全く困らない。
そこまではなくても Twitter の独裁化が進むであろうことは容易に想像できるし，そうなればその「孤独な正義」からこぼれ落ちる個人は（私も含め）増えるかもしれない。

EFF も言っているではないか。

{{< fig-quote type="markdown" title="フェディバースは（“私たち”次第で）素晴らしいものになる" link="https://p2ptk.org/freedom-of-speech/4165" >}}
善き独裁者を選んでも、独裁体制を修正することはできない
{{< /fig-quote >}}

と。
社会が「多様性（diversity）」を重視しつつある昨今，プラットフォーム覇権主義はもはや無理筋と言えるだろう。
そこで私達にできることは，陳腐ではあるが，覇権プラットフォームにロックインされないことである。
『[Clean Architecture](https://www.amazon.co.jp/dp/B07FSBHS2V?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』的に言うなら「プラットフォームは詳細である」といったところか（笑）

## ブックマーク

- [メディアは（常に）スパムか？ « マガジン航[kɔː]](https://magazine-k.jp/2016/01/25/spam-and-media/)
- [タイムラインの奴隷 - Spiegel's Branch - Scrapbox](https://scrapbox.io/spiegel-branch/%E3%82%BF%E3%82%A4%E3%83%A0%E3%83%A9%E3%82%A4%E3%83%B3%E3%81%AE%E5%A5%B4%E9%9A%B7)
- [徒然ブログ]({{< ref "/remark/2021/06/weblog.md" >}})

## 参考図書

{{% review-paapi "4873119464" %}} <!-- ユニコーン企業のひみつ -->
{{% review-paapi "B0125TZSZ0" %}} <!-- つながりっぱなしの日常を生きる -->
{{% review-paapi "B01MZGVHOA" %}} <!-- 超監視社会 -->
{{% review-paapi "4903127044" %}} <!-- 排除型社会 -->
{{% review-paapi "B00HY73M16" %}} <!-- Be mine! -->
{{% review-paapi "B07K356C43" %}} <!-- 転スラ Another Colony -->
