+++
date = "2015-12-09T11:52:34+09:00"
description = "description"
draft = true
tags = ["golang", "programming", "inheritance", "delegation", "embed", "interface"]
title = "継承か委譲か，それが問題だ？"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

- [オブジェクト指向言語としてGolangをやろうとするとハマること - Qiita](http://qiita.com/shibukawa/items/16acb36e94cfe3b02aa1)
    - [オブジェクト指向言語としてGolangをやろうとするとハマる点を整理してみる - Qiita](http://qiita.com/sona-tar/items/2b4b70694fd680f6297c)
- [Go言語に継承は無いんですか【golang】 - DRYな備忘録](http://otiai10.hatenablog.com/entry/2014/01/15/220136)

ただし， interface および embed を組み合わせることによってオブジェクト指向設計における「is-a 関係」（「汎化」および「特化」）を実装することは可能である。
「is-a 関係」を実装する手段は継承に限らないということだね。

継承に必要な機能を実現することはできる。
オブジェクト指向設計における「汎化」の実装は継承か委譲かということは大きな問題ではないということだね。

