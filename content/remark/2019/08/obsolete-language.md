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

[母国語]: {{< ref "/remark/2015/programming-language.md" >}} "プログラミング言語との付き合い方"

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253"><img src="https://images-fe.ssl-images-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan, Brian W. Kernighan</dd>
    <dd>柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN)</dd>
    <dd>Rating<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home" >PA-API</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%94%B9%E8%A8%822%E7%89%88-%E3%81%BF%E3%82%93%E3%81%AA%E3%81%AEGo%E8%A8%80%E8%AA%9E-%E6%9D%BE%E6%9C%A8-%E9%9B%85%E5%B9%B8-ebook/dp/B07VPSXF6N?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B07VPSXF6N"><img src="https://images-fe.ssl-images-amazon.com/images/I/51jif840ScL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%94%B9%E8%A8%822%E7%89%88-%E3%81%BF%E3%82%93%E3%81%AA%E3%81%AEGo%E8%A8%80%E8%AA%9E-%E6%9D%BE%E6%9C%A8-%E9%9B%85%E5%B9%B8-ebook/dp/B07VPSXF6N?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B07VPSXF6N">改訂2版 みんなのGo言語</a></dt>
    <dd>松木 雅幸, mattn, 藤原 俊一郎, 中島 大一, 上田 拓也, 牧 大輔, 鈴木 健太</dd>
    <dd>技術評論社 2019-08-01 (Release 2019-08-01)</dd>
    <dd>Kindle版</dd>
    <dd>B07VPSXF6N (ASIN), 9784297107284 (EISBN)</dd>
  </dl>
  <p class="description">今読んでるとこ。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-08-05">2019-08-05</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home" >PA-API</a>)</p>
</div>
