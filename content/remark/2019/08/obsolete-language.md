+++
title = "「死んでない」プログラミング言語とは"
date =  "2019-08-05T22:50:03+09:00"
description = "この手の記事に反応してしまうとか，我ながらミーハーなことである。反省はしない（笑）"
image = "/images/attention/kitten.jpg"
tags = [ "programming", "language" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

最近あまりネット巡回をしなくなったせいか，どうしてもネタ元が偏ってしまう（笑） ちうわけで今回も [yomoyomo さんのブログ記事](https://yamdas.hatenablog.com/entry/20190805/doomed-programming-language "おそらく先がない5つのプログラミング言語？ - YAMDAS現更新履歴")から。

- [5 Programming Languages That Are Probably Doomed](https://insights.dice.com/2019/07/29/5-programming-languages-probably-doomed/)

この手の記事に反応してしまうとか，我ながらミーハーなことである。
反省はしない（笑）

この記事では「先のないプログラム言語」として Ruby, Haskell, Objective-C, R, Perl の5つを挙げている。
個人的には納得がいくものだ。
いや Haskell にはもうちょっと頑張って欲しい気持ちはあるのだが。

私が Ruby 勉強会に参加したのはもう[9年も前のこと](https://baldanders.info/blog/000502/ "https://baldanders.info/blog/000502/")だが，実は Ruby よりも Erlang のほうが印象に残ってしまった妙な勉強会であった（今思い出しても酷い司会だったよな`w`）。
だいたい2010年前後からプログラミング言語のトレンドが大きく変わっているような気がする。

「死んでいない」ことと「生きている」ことには彼岸と此岸ほどの差がある。
レガシー・システムというのは廃炉原発みたいなもんで（喩えが悪くてスマン），本当に「終わる」までには随分な時間とコストを要する。
だからって「あと10年は戦える」とか言ったところで釣り銭を誤魔化しているようにしか聞こえない。

たとえば（汎用機システムに過剰に最適化してしまった） COBOL が死んでないからといって前途ある若者に「COBOL を勉強しろ」とか言うか？ 言わないだろう。
そういう観点で考えると，最初に挙がってた5つの言語を指して “doomed” というのに納得してしまうのである（いや Haskell にはもうちょっと頑張って欲しいのだが）。

私は「手続き型言語」では Go 言語が究極だと思っていて，もし「それ以上」を望むなら最早「手続き型言語でないもの」にシフトするしかないと考える。
それは単なるトレンドの変化ではなくもっと大きなパラダイムシフトである。
その「手続き型言語でないもの」が次の世代の「[母国語]」になれば，その言語は本当に「あと10年は戦える」と言っていいんじゃないだろうか。

## ブックマーク

- [5 Powerful Programming Languages to Stretch Your Brain - DEV Community 👩‍💻👨‍💻](https://dev.to/jacobherrington/5-programming-languages-to-stretch-your-brain-mmg)

[母国語]: {{< ref "/remark/2015/programming-language.md" >}} "プログラミング言語との付き合い方"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN), 4621300253 (ISBN), 9784621300251 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/B07VPSXF6N?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/51jif840ScL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/B07VPSXF6N?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">改訂2版 みんなのGo言語</a></dt>
    <dd>松木 雅幸 (著), mattn (著), 藤原 俊一郎 (著), 中島 大一 (著), 上田 拓也 (著), 牧 大輔 (著), 鈴木 健太 (著)</dd>
    <dd>技術評論社 2019-08-01 (Release 2019-08-01)</dd>
    <dd>Kindle版</dd>
    <dd>B07VPSXF6N (ASIN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">改訂2版の目玉は7章の「データベースの扱い方」が追加されたことだろう。他の章では，大まかな構成は1版と同じだが細かい部分が変わっていて Go 1.12 への言及まであるのには驚いた。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-08-12">2019-08-12</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
