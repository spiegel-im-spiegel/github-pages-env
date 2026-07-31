+++
title = "Wiki という時代の終焉"
date =  "2017-09-26T11:08:32+09:00"
description = "かつて Hacker なる言葉が時代とともに意味を消失していったように， Wiki も消失し空っぽの名前だけが遺跡のように残るのだろう。"
tags = [ "wiki", "internet", "media" ]

[scripts]
  mathjax = false
  mermaidjs = true
+++

（この記事は Facebook の TL に書きなぐった内容を再構成してお送りしております）

- [wikipediaをwikiと略すな - ぽめcの脳内を晒す場所](http://pomec-blog.hatenablog.com/entry/20170925/1506351240)

実は私も割と最近まで「Wikipedia を Wiki と略すな」派だったのですよ。

若い人は知らないかも知れないけど， Wiki の語源は [WikiWikiWeb] から来ている。
この [WikiWikiWeb] ってのが当時は「とてもクールなサービス」と評価され，クローンがぼこぼこ登場した。
そんで [WikiWikiWeb] のコンセプトを持つ Web サービスを総称して Wiki と呼ぶようになった。

Wikipedia もそんな「クローン」のひとつであり，したがって「Wiki は Wikipedia の略称」なのではなく「Wikipedia は Wiki の一種」というのが厳密に正しいのだ。
「オブジェクト指向」的に言うと Wiki と Wikipedia は is-a 関係にある，とも言える。

{{< fig-mermaid >}}
graph RL
  Wikipedia-->Wiki
{{< /fig-mermaid >}}

つまり Wikipeadia を Wiki と呼ぶことはマクドナルドを（「マクド」でも「マ」でもなく）ハンバーガー屋と呼ぶことと等価なのである。

しかし考えてみて欲しい。
今現在 Wikipedia 以外に社会的に成功している Wiki サービスがあるだろうか（いやない[^wk1]; 反語）。

[^wk1]: いや，もちろん [TeX Wiki](https://texwiki.texjp.org/) や [Scrapbox](https://scrapbox.io/) のような Wiki サービスも存在していることは知ってますよ，念のため。

{{< fig-quote title="wikipediaをwikiと略すな" link="http://pomec-blog.hatenablog.com/entry/20170925/1506351240" >}}
<q>ゲームによくある攻略Wikiも、Wikipediaのゲーム攻略版か何かかと思っていました。</q>
{{< /fig-quote >}}

実は，今やゲーム攻略サイトで Wiki 運営をしているところは殆どないらしい[^gm1]。
名前に「◯◯攻略Wiki」と付いているところでさえも。
もはや Wiki という名称に意味なんてないのだ。
ならば「Wikipedia を Wiki と略す」ことに問題があるのか。

[^gm1]: ゲーム攻略サイトをアフィリエイト化する過程で Wiki 運営を継続することは軋轢にしかならないしね。

――

今や誰も口にしなくなったけど，昔は Wikinomics なる造語も存在していて，要するに Wiki ってのは単に機能やサービスを指す言葉ではなく「バザール型のコラボレーション」活動を指すものに昇華していったのだ。
それはインターネットやオープンソースとかいった「古き良き時代」を象徴するものでもあったりする[^gh1]。

[^gh1]: 現在「バザール型のコラボレーション」活動を象徴するのは GitHub かな。 GitHub のドキュメント記法が Wiki 記法じゃなく Markdown ってのも象徴的だよねぇ（笑） そう考えるとインターネットは「古き良き時代」に固執するクラスタ，「テレビ化」に迎合するクラスタ，社交（SNS）に引きこもるクラスタなどに多極化している，と見なすこともできるかもしれない。

一方 Wiki が衰退し始めた中で台頭してきたのが Twitter である。
Twitter が登場したときのフィーバーぶりも大変なもので，これのクローンもぼこぼこと登場したことを覚えている人もいるだろう（そういう意味じゃ「[マストどん]({{< ref "/remark/2017/04/mastodon.md" >}} "GW 過ぎたらみんな忘れてるに100カノッサ")」なんていまさら過ぎるよね）。
当然ながら Twitter は社会現象となり Twitternomics なる造語まで登場した。
今や Twitter が差別と排除を増幅する{{< ruby "メディア" >}}装置{{< /ruby >}}として機能している現状からは考えられないかもしれないが（笑）

今にして思えば Twitter は，スマートフォンの普及と併せて，「ネットのテレビ化」の急先鋒だったわけだ。

ネット上の Wiki システムが絶滅危惧種化していき，その隙間を埋めるように Twitter をはじめとする「テレビ」が侵食していくさまは，ひとつの時代の終わりを示すものである。
かつて Hacker なる言葉が時代とともに意味を消失していったように， Wiki も消失し空っぽの名前だけが遺跡のように残るのだろう。

なんてな（by いかりや長介）

## ブックマーク

- [Scrapbox または Wiki で再び遊ぶ]({{< ref "/remark/2017/01/scrapbox-and-wiki.md" >}})

[WikiWikiWeb]: http://www.hyuki.com/yukiwiki/wiki.cgi?WikiWikiWeb "WikiWikiWeb - WikiWikiWebとは何か"

## 参考図書

{{% review-paapi "4844319159" %}} <!-- 結城浩のWiki入門 -->
{{% review-paapi "482224587X" %}} <!-- ウィキノミクス -->
{{% review-aozora "227" %}} <!-- 伽藍とバザール -->
