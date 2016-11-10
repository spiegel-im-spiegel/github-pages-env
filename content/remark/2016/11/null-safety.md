+++
date = "2016-11-10T16:48:59+09:00"
description = "description"
draft = true
tags = ["programming"]
title = "「null 安全」について"

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

- [null安全でない言語は、もはやレガシー言語だ - Qiita](http://qiita.com/koher/items/e4835bd429b88809ab33)
- [null安全を誤解している人達へのメッセージ - Qiita](http://qiita.com/omochimetaru/items/ee29d4c6eb0d78f02b15)

「null 安全（null safety）」についての上の記事はなかなか興味深かった。
特に「[null安全を誤解している人達へのメッセージ](http://qiita.com/omochimetaru/items/ee29d4c6eb0d78f02b15)」は事実上の FAQ になっているので是非読んでみてほしい。

以下は個人的なメモである。

「参照」は昔から悩ましい問題である。
「null 参照」は「参照が無効である」ことを示すものだ。
どんなプログラムであれ「null 参照」が存在するのであれば，それを異常系としてどこかで始末しなければならない。

オブジェクト指向以前，たとえばアセンブラや C 言語の時代では値と参照は明確に区別されていなかった。
「それ」を値と見なすのか「参照（ポインタ）」と見なすのかは完全にプログラマの責任だったのである。

それからオブジェクト指向プログラミングが産業分野でも台頭してきたのだが，このパラダイムシフトの過程で「参照」が言語レベルで明確に意味を持つものとなった。
ここで「null 参照」をどう扱うか（そもそも「null 参照」を許容するのか）が重要になってくる。

「null 参照」はどこにでもある。
たとえば SQL 文法や JSON データ表現では値が存在しないことを null で表す。
プログラマは「値が存在しないこと」を安直に「null 参照」で表現しようとする[^vo]。

[^vo]: もちろん賢明なプログラマは安易な「null 参照」ではなく，適切な value object を使うだろう？

もちろんプログラマはみんな（みんなと言い切ってしまうが）「null 参照」が厄介であることは分かっている。
オブジェクト指向プログラミングは状態の詳細を隠ぺいしてしまう。
ロジックの深いところで発生した「null 参照」はしばしばプログラマが予測しないところで牙をむく。
例外処理で捕まえようとしてもいつの間にかすり抜けてしまう。
このことによる経済的損失を10億ドルと見積もった人[^ah]もいる。

[^ah]: 「[アントニー・ホーア - Wikipedia](https://ja.wikipedia.org/wiki/%E3%82%A2%E3%83%B3%E3%83%88%E3%83%8B%E3%83%BC%E3%83%BB%E3%83%9B%E3%83%BC%E3%82%A2)」参照。

「null 安全）」な言語では nullable な（null かもしれない）参照と non-null な（null を許容しない）参照をコンパイルレベルで区別する。
この「コンパイルレベルで」というのがすごく重要。

もちろん設計やテストで「null 参照」に起因するバグや脆弱性を回避するのは重要である。
でも，それでも「null 参照」に起因するバグや脆弱性は発生し続けているのである。

「null 安全）」はもともとプログラマの責任であった「null 参照」の始末を


## ブックマーク

- [なぜNullはダメか | To Be Decided](https://tbd.kaitoy.xyz/2015/07/26/why-null-is-bad/)
