+++
date = "2016-01-29T16:46:49+09:00"
description = "description"
draft = true
tags = ["golang", "unicode", "normalization"]
title = "Go 言語と Unicode 正規化"

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

- [Text normalization in Go - The Go Blog](https://blog.golang.org/normalization)
- [norm - GoDoc](https://godoc.org/golang.org/x/text/unicode/norm)
- [Golang: Unicode 正規化 - Sarabande.jp](http://blog.sarabande.jp/post/89636452673)
- [Go で UTF-8 の文字列を扱う - Qiita](http://qiita.com/masakielastic/items/01a4fb691c572dd71a19)
- [文字コード地獄秘話 第3話：後戻りの効かないUnicode正規化 - ALBERT Engineer Blog](http://tech.albert2005.co.jp/blog/2014/11/21/mco-normalize/)
- [本の虫: Linus Torvalds、HFS+に激怒](http://cpplover.blogspot.jp/2015/01/blog-post_14.html)
- [Golang: Unicode 正規化 - Sarabande.jp](http://blog.sarabande.jp/post/89636452673)
- [文字列の表記揺れをUnicode正規化で簡単に解決する方法 - Qiita](http://qiita.com/y-ken/items/d08eb7f66c8fb2fa7d21)
- [Mac OS Xの濁点ファイルがやってきた - miauの避難所](http://d.hatena.ne.jp/miau/20110805/1312555736)
- [Unicode正規化 - Wikipedia](https://ja.wikipedia.org/wiki/Unicode%E6%AD%A3%E8%A6%8F%E5%8C%96)
- [Golang: 文字列のバイト数と文字数を数える - Sarabande.jp](http://blog.sarabande.jp/post/61104546593)

いわゆる「異体字」は正規化されないようだ。
