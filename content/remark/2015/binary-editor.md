+++
date = "2015-12-04T23:14:34+09:00"
update = "2015-12-05T17:12:51+09:00"
description = "Windows で EBCDIC を扱えるバイナリ・エディタを紹介。"
draft = false
tags = ["tools", "editor", "ebcdic", "windows"]
title = "Windows 用バイナリ・エディタ"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"
+++

長らく組込み開発から離れているせいもあるが，昨今めっきりバイナリデータを触る機会が減った。
データを dump out することはあっても[^dump]，直に弄ることは少なくなった。
それでも，全くなくなったわけではなく，そうなった時にいつも「ええつと，バイナリ・エディタってどこにあるっけ？」と探しまわることになるのだ。

[^dump]: データを dump out するだけならいくらでも手段がある。懐かしいところだと UNIX 系の `od` コマンドとかあるし（`-tx1z` とかオプションを付けると幸せ），大抵のスクリプト言語なら手軽に dump out できる。

特に最近は EBCDIC，それも PACKED DECIMAL を弄らないといけなくて，たまにモニタを殴りたくなる。

- [EBICDICで符号付数値型がなんだって？ « motoama's](http://motoama.chu.jp/program/905)
- [Packed-Decimal Format, Description and Discussion](http://www.simotime.com/datapk01.htm)

そうでなくても（ASCII コードに慣れてる身としては） EBCDIC は直感的じゃないのに。

## Windows で使えるバイナリ・エディタ

というわけで， Windows で使えるバイナリ・エディタ。
実は色々あるのだが， EBCDIC を扱えるとなると以下の3つが妥当だろう。

- [Stirling]
- [Bz]
- [xedit]

[Stirling] はかなり古いアプリケーションで，最新版の 1.31 が出たのは1999年のようだ。
古いのがいけないわけではないが（実際，現在でも遜色ないほど高機能で職場で使ってる人もいる），流石に15年以上前のものを使うのはねぇ...

というわけで，残りの2つが個人的におすすめなのだが，今回は [Bz] を紹介する[^xe]。

[^xe]: [xedit] は組込み向けの開発に向いているが，今回は割愛する。ちなみに EBCDIC 対応でもカナや漢字をまともに扱えるものは存在しない。そもそもカナや漢字は汎用機ごとに方言がキツい（たとえ Shift-JIS でも旧 JIS だったりする）ので，汎用のツールはないと考えたほうがいい。

## Binary Editor Bz

知る人ぞ知るだが， [Bz] のオリジナルはあの [c.mos](http://www.vcraft.jp/) さんによるものである[^cmos]。
で，そのオリジナルのコードを [devil.tamachan](https://github.com/devil-tamachan) さんが改造したのが今回紹介する [Bz] である。

[^cmos]: [c.mos](http://www.vcraft.jp/) さんといえば Vz Editor。私たちの世代から見れば神のようなお方である。

- [Binary Editor BZ](http://devil-tamachan.github.io/BZDoc/) : ヘルプページ

現時点での最新版は [1.9.7.1](https://github.com/devil-tamachan/binaryeditorbz/releases/tag/v1.9.7.1)。
Portable 版の zip ファイルの中身を適当なフォルダにコピーして起動すればよい。
簡単！

ただし，最新のバージョンは `EBCDIC.def` ファイルがないため，そのままでは EBCDIC で表示できない。
`EBCDIC.def` ファイルはググれば見つかると思うが，一応[ここ](/material/bz/EBCDIC.def)にも置いておく。

[Bz] では ASCII や EBCDIC の他， Shift-JIS, JIS, EUC や Unicode 系の文字エンコードィングにも（一応）対応している。
その他の特徴としては

1. 既定で書き込み禁止（間違って弄らないようにするため）
1. 構造体解析が可能
1. データをビットマップで俯瞰できる
1. オフセットジャンプ（カーソル位置の値分だけジャンプする）
1. 画面を2分割まで可能

ってとこかな。
個人的にはこれで必要十分。
贅沢を言えば 10GB 程度のファイル[^file] を読み込んでもへこたれないでほしいものだが，まぁそこは無理は言うまい[^size]。

[^file]: そういうのがあるんだってば，どってんばってん。
[^size]: [Bz] は今のところ 4GB までしか扱えない。

[Bz]: https://github.com/devil-tamachan/binaryeditorbz "devil-tamachan/binaryeditorbz"
[Stirling]: http://www.vector.co.jp/soft/win95/util/se079072.html "Stirlingの詳細情報 : Vector ソフトを探す！"
[xedit]: http://www002.upp.so-net.ne.jp/janus/xedit.html "ROM化支援バイナリエディタ - xedit -"
