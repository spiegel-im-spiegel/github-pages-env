+++
title = "「リモートデスクトップに乗っ取り可能な脆弱性」とか何の冗句かと..."
date =  "2017-12-12T09:01:50+09:00"
description = "今回の脆弱性は乗っ取る側が別の第3者にリモート制御を取られる可能性がある，というわけやね。ミイラ取りがミイラに（笑）"
image = "/images/attention/remark.jpg"
tags        = [ "security", "vulnerability", "risk" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

- [リモートデスクトップソフト「TeamViewer」に乗っ取り可能な脆弱性、修正版が公開 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1096259.html)

いや，この記事に文句があるわけではなく，「リモートデスクトップに乗っ取り可能な脆弱性」とか何の冗句かと笑ってしまったのだ。

で， Facebook の TL に「そもそもリモートデスクトップって相手のマシンを乗っ取るツールだろう」と呟いた[^fb1] ら反応をいただいて，実際にそういうツールを使ったテクニカルサポート詐欺の事例が結構あるらしい。
YouTube の検索で “[Tech support scam](https://www.youtube.com/results?search_query=Tech+support+scam "Tech support scam - YouTube")” と入れるとわらわら出て来る，と教えていただいた。

[^fb1]: 私が主に呟く場所は Facebook の TL です。栃木弁の某芸人さんではないので呟きを公衆空間にばら撒く勇気はないです。最近は少し Twitter でも呟くようになりましたが。まぁ，嘘や間違いを書けばお叱りを受けるのはどちらでも同じですが，少なくともいきなりマサカリが飛んで来ることはない（笑）

{{< fig-youtube id="pDqUQupEcE4" title="Fake Tech Support Scammer - $1,000 One Time Fix (Event Viewer Scam) - YouTube" >}}

日本でも昨年話題になったようだ。

- [最近発生している遠隔操作サポート詐欺に対する注意喚起 - 空気を読まない中杜カズサ](http://nakamorikzs.net/entry/20160930/remotescam)
- [不安をあおって電話でだます「サポート詐欺」の手口を追う | トレンドマイクロ セキュリティブログ](http://blog.trendmicro.co.jp/archives/13970)

当時は忙しかったし，セキュリティ・エリアで仕事しててネットとは切り離されたりしてたからなぁ。

今回の脆弱性は乗っ取る側が別の第3者にリモート制御を取られる可能性がある，というわけやね。
ミイラ取りがミイラに（笑）

うーむ。
世の中なかなか侮れん。
