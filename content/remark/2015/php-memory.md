+++
date = "2015-09-23T10:42:51+09:00"
description = "うひゃあ。PHP って今そんなことになっているのか。今度仕事で PHP やる機会があっても最初から勉強しなおしだな，こりゃ。"
draft = true
tags = ["engineering", "php"]
title = "PHP の思ひ出"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

（この記事は「プリキュア・メモリ」を脳内 BGM にお送りしています）

- [モダンPHPアンチパターン - Qiita](http://qiita.com/tadsan/items/157969b338fd8b782b21)

うひゃあ。
[PHP] って今そんなことになっているのか。
今度仕事で [PHP] やる機会があっても最初から勉強しなおしだな，こりゃ。

## ASP の劣化コピーだろ，これ

[PHP] は最初から気に食わない言語だった。
だって ASP （ASP.NET じゃなくて ASP）の劣化コピーにしか見えなかったんだもの。
MS ですら ASP.NET に移行していたあの当時に [PHP] が魅力的に見えるはずがなかった。

## PHP で制御を行う

[最後に PHP の仕事をしたのは2010年](http://www.baldanders.info/spiegel/log2/000481.shtml)かな。
ネットワークに繋がっている機器を [PHP] で（もちろん Web インタフェースで）制御するという仕事で，初めて [PHP] が面白いと思った仕事だった。
ちなみにこの時のフレームワークは [Smarty](http://www.smarty.net/) だった。

今まで C/C++ でゴリゴリと制御を書いてきた身としてはまさに「目からうろこが落ちる」状態だった。
スクリプト言語で制御ができるなんて！

当時の機器は制御ロジック自体は内部に埋め込まれていた。
だから programmable な部分は

1. 機器に制御のトリガ・イベントを送る
1. 機器間の同期をとる

ことであり，トリガ・イベントを送るだけでいいのならスクリプト言語でも構わないわけだ。

IoT たらいうバズワードでもてはやされる最近の「スマート家電」の特徴は，制御だけでなくセンサも備えていることだが，センシング・データを「情報」として解釈して自律的に動けるほど「スマート」なものは少ない[^a]。
結局，データの解釈とその結果として何らかの制御トリガを蹴ったくるのは「外部のなにか」であり，その「外部のなにか」の真正性が問題になる[^b]。

[^a]: それができるなら，もう家電じゃなくてロボットだけどね。ほとんどロボットみたいな家電は既にあるが。
[^b]: そもそも機器自体（というかその背後にいる存在）の悪意の有無は，実際に被害に遭うまで分かりようがない。

## PHP は死なず？

これは印象論だが，ゼロからシステムを組むのであれば [PHP] を使う必然性は全くない。
Web サービス限定であっても「後ろから前まで JavaScript でいいじゃん」な感じである。
デスクトップアプリですら [Electron](http://electron.atom.io/) で組めるんだから。

まぁでも，安いレンタルサーバを借りてる場合は選択肢は少ないし，既に [PHP] が稼働している環境なら [PHP] のほうが手軽なことも多々ある[^c]。

[^c]: 昔， [Perl で RSS を JSON に変換しようとしてどえら苦労した](http://www.baldanders.info/spiegel/log2/000236.shtml)が， PHP なら数行の記述でほぼ同じことができる。

あるいは[日本以外にも PHP が人気な国もある](http://qiita.com/naru0504/items/9bd56998a187d101a777)そうなので，「ただ消え去る（fade away）のみ」とはならないかもしれない。
ただ，時々舞い込んでくる「汎用機の COBOL システムをオープン系のシステムに replace する」案件に関わってると，似たようなことが [PHP] 界隈でも起きないか心配なところではある。

[PHP]: https://secure.php.net/ "PHP: Hypertext Preprocessor"
