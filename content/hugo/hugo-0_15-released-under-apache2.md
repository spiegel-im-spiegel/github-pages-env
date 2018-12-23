+++
date = "2015-11-26T20:26:45+09:00"
update = "2017-06-21T16:10:38+09:00"
description = "追加された機能が盛りだくさんあり詳しくは上のリンク先を確認していただくとして，実は今回の目玉はライセンスのアップグレードである。"
draft = false
tags = ["hugo", "license", "apache2", "contribution"]
title = "Hugo 0.15 が Apache License 2.0 下でリリース"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

[Repository](https://github.com/gohugoio/hugo/) ではしばらく前からリリース準備が行われていて wktk 状態で待っていたのだが，ようやくリリースの運びとなった。

- [Release v0.15 · gohugoio/hugo](https://github.com/gohugoio/hugo/releases/tag/v0.15)

```bash
C:> hugo version
Hugo Static Site Generator v0.15 BuildDate: 2015-11-25T14:35:20+09:00
```

追加された機能が盛りだくさんあり詳しくは上のリンク先を確認していただくとして，実は今回の目玉はライセンスのアップグレードである。

[Hugo] のライセンスは [Simple Public License (SimPL) 2.0](https://opensource.org/licenses/Simple-2.0) だったのだが，これを企業ユーザでも利用しやすい一般的なライセンスにしてくれという要望というか議論は随分前からあったらしい。

- [Upgrade the license to an open one · Issue #201 · gohugoio/hugo](https://github.com/gohugoio/hugo/issues/201)

[SimPL-2.0] はオープンソース・ライセンスのひとつで GPLv2 を（文字通り）よりシンプルに記述したもののようである。

{{< fig-quote title="Simple Public License (SimPL-2.0)" link="https://opensource.org/licenses/Simple-2.0" lang="en" >}}
<q>This Simple Public License 2.0 (SimPL-2.0 for short) is a plain language implementation of GPL 2.0.  The words are different, but the goal is the same - to guarantee for all users the freedom to share and change software.  If anyone wonders about the meaning of the SimPL, they should interpret it as consistent with GPL 2.0.</q>
{{< /fig-quote >}}

当初は GPL や MIT ライセンスなどが候補に挙がっていたようだが，最終的には [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0) （[日本語訳](https://osdn.jp/projects/opensource/wiki/licenses%2FApache_License_2.0)）に決まった。
[Apache-2.0] は copyleft ではないがオープンソース・ライセンスのひとつであり，特許の終了規定やコード等の寄与（contribution）に関する規定があるのが特徴。
さらに [GPLv3 と互換性のあるライセンス](http://www.gnu.org/licenses/license-list.ja.html#apache2)として認められている。

[Hugo] ではユーザによるコードの寄与が結構あるため，この点が [Apache-2.0] を選択する決め手になったようだ。

[GitHub] などは pull request で気軽にコードを寄与できるメリットがあるが，寄与されたコードの扱いが問題になることもある。
Copyleft であれば寄与されたコードも自動的に元のライセンスに従うため混乱は少ないが，そうでない場合は [Apache-2.0] のような明示的な規定が必要になるだろう[^a]。

[^a]: 更に言うと， [Go 言語]では気軽にパッケージを import できるが，最終的な製品のライセンスをどうするかは結構重要な問題である。 [Hugo] のライセンス・アップグレードの際には，この辺の確認でも時間がかかった感じだ。事ほど左様にライセンスの互換性とは重要な問題なのである。ちなみに [Hugo] のようなコード・ジェネレータあるいはもっと広く CASE (Computer Aided Software Engineering) ツールが生成するコードが誰に帰属するのかは，また別の問題である。

## ブックマーク{#bookmark}

[Hugo に関するブックマークはこちら]({{< ref "/hugo/bookmark.md" >}})。

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[SimPL-2.0]: https://opensource.org/licenses/Simple-2.0 "Simple Public License (SimPL-2.0) | Open Source Initiative"
[Apache-2.0]: http://www.apache.org/licenses/LICENSE-2.0 "Apache License, Version 2.0"
[GitHub]: https://github.com/ "GitHub"
[Go 言語]: https://golang.org/ "The Go Programming Language"
