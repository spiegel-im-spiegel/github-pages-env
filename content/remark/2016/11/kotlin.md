+++
draft = false
date = "2016-11-08T21:02:26+09:00"
update = "2016-11-10T18:45:34+09:00"
title = "Kotlin に関する覚え書き"
description = "今回は試して遊ぶ余裕がないので本当にただのメモ。"
tags = [
  "programming",
  "language",
  "kotlin",
  "java",
  "null-safety",
]

[author]
  name = "Spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  url = "https://baldanders.info/profile/"
  avatar = "/images/avatar.jpg"
  github = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  tumblr = ""
  flickr = "spiegel"
  flattr = ""
  facebook = "spiegel.im.spiegel"
  linkedin = "spiegelimspiegel"
+++

今回は試して遊ぶ余裕がないので本当にただのメモ。

[Kotlin] は IntelliJ IDEA で有名な [JetBrains](http://www.jetbrains.com/ "JetBrains: Development Tools for Professionals and Teams") 社が中心となって開発が行われているオブジェクト指向のプログラミング言語だ。
研究用とかではなく明確に産業利用を目的に作られている。

JavaVM で動作可能なバイトコードを吐くそうで， Android アプリの開発などでもすでに実績があるようだ。
また JavaScript のコードも吐けるらしい。

[Kotlin] の特徴の一つは「null 安全（null safety）」であることのようだ。
「null 安全」とは実行時にいわゆる「ぬるぽ（null pointer error or exception）[^np]」が発生しない言語仕様になっていることで， [Kotlin] の場合はコンパイル時にチェックされる。

[^np]: null の英語での発音は「ナル」に近いのだが，どういうわけか日本の IT 界隈ではドイツ語っぽく「ヌル」と呼ぶ人が圧倒的多数のようだ。私もそうだし職場の半径100m以内で「ナル」と呼ぶ人はいない。まぁ私は TTC (Telecommunication Technology Committee) を「てってーしー」と呼んでいた世代だからな（笑） （参考：[IT業界で横行する恥ずかしい英語発音 - Qiita](http://qiita.com/ryounagaoka/items/290885ee3291b393fe1f)：コメントの議論が面白いので是非）

アプリケーションが「ぬるぽ」で落ちるならまだマシなほうで，最悪の場合は重大な脆弱性問題を引き起こす場合もあるため，「null 安全」という考え方は近年とくに注目されている。
そのため今どき流行りの言語では「null 安全」が取り入れられつつあるらしい[^go]。

[^go]: ちなみに [Go 言語]は全く「null 安全」ではない。「ぬるぽ」を緩和するような仕組みはないこともないが，コンパイラレベルで保証しているわけではない。（参考： [Go言語がダメな理由 | プログラミング | POSTD](http://postd.cc/why-go-is-not-good/)）

更に [Kotlin] は Java からの置き換えを目論み，かつ言語仕様を拡張させている点も特徴である（Scala や Groovy 由来の機能や糖衣構文もある）。
少なくとも [Kotlin] では関数を第一級オブジェクト（first-class object）として扱うことができるのは大きい。
個人的には [Java は終わってる]({{< ref "/remark/2016/07/java.md" >}} "Java はやめておけ")と思ってるので，こういうのは大歓迎である。

## ブックマーク

- [プログラミング言語Kotlin 解説](https://sites.google.com/site/tarokotlin/)
- [Kotlinを１ヶ月使ってみた - Qiita](http://qiita.com/noppefoxwolf/items/b2d93f946f158c7b5852)
- [Kotlinに対する雑感](http://blog.satotaichi.info/first-thoughts-of-kotlin/)
- [null安全でない言語は、もはやレガシー言語だ - Qiita](http://qiita.com/koher/items/e4835bd429b88809ab33)

[Kotlin]: https://kotlinlang.org/ "Kotlin Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
