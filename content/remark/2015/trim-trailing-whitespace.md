+++
date = "2015-11-23T20:42:06+09:00"
description = "行末の空白文字を削除する小ネタ。いや，たまに使おうとすると忘れてるんだよね。"
draft = false
tags = ["tools", "editor", "atom", "hidemaru", "sakura"]
title = "行末の空白文字を削除する"

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
  url = "https://baldanders.info/spiegel/profile/"
+++

テキストエディタで行末の空白文字を削除する小ネタ。
いや，たまに使おうとすると忘れてるんだよね。

## ATOM Editor の場合

以前は行末の処理は [editorconfig] でできてたはずなんだけど，いつの間にか `trim_trailing_whitespace` と `insert_final_newline` は “doesn't work yet” になってる。
それとも私の勘違いだったのか。

気を取り直して。

行末の空白を一気に削除したいならコマンドパレットで `whitespace` を検索すると “Remove Trailing Whitespace” があるので，これを実行する。
保存するたびに自動的に行末の空白文字を削除したいのであれば [whitespace] の Settings にある “Remove Trailing Whitespace” を有効にする。

{{< fig-img src="https://photo.baldanders.info/flickr/image/23214924646_m.png" title="settings for whitespace (ATOM)" link="https://photo.baldanders.info/flickr/23214924646/" >}}

おまけだが，同じ Settings 画面で “Ensure Single Trailing Newline” を有効にすると，テキストファイル末尾が改行になっていない場合は改行を補ってくれる。
他にもこまごまとした設定があり，しかもファイルタイプごとに設定できるようだ。
でもファイルタイプではなく [editorconfig] でコントロールできるようになってほしい。

もうひとつ余談だが， [Go 言語環境]({{< ref "/golang/golang-with-atom.md" >}})ではコードの整形を行う際に [whitespace] の設定に関係なく行末の空白を削除してくれる。
他の言語でも整形ツールがあれば同様にできるかも。

## 秀丸またはサクラエディタの場合

現在，ほとんどの作業は [ATOM] に移行できているのだが，巨大ファイルを扱う場合などはまだ[秀丸]等[^a] のお世話になっている。
[秀丸]ではファイルタイプごとに保存時に行末の空白を削除するかどうか設定できる。
あるいは置換機能を使う方法もある。
手順は以下の通り。

[^a]: 職場では[秀丸]が NG の場合もあるので，その場合は[サクラエディタ]で代替えしている。

1. 「検索文字列」に **`[ 　\t]+$`** を指定する（行末の半角空白，全角空白，タブ文字を正規表現で指定する）
1. 「置換文字列」はブランクにする（何もセットしない）
1. 「正規表現」の項目にチェックを入れて置換を開始する

ちなみに同様の手順は[サクラエディタ]でも使える[^b]。
てか， [ATOM] も含め置換処理に正規表現が使えるエディタならたいてい使える[^c]。

[^b]: [サクラエディタ]では「タイプ別設定」の「スクリーン → インデント」の項目で「改行時に末尾の空白を削除」にチェックを入れることで行末の空白を自動的に削除できる。なんでこんな変なところに設定項目があるんだろう。
[^c]: でも正規表現ってあまり得意じゃないので，普段は忘れてるんだよなぁ。昔，正規表現を簡単に組み立てることのできるツールがあって重宝していたが，マシンを replace していくうちに行方不明になっちゃった。

[ATOM]: https://atom.io/ "Atom"
[editorconfig]: https://atom.io/packages/editorconfig "editorconfig"
[whitespace]: https://atom.io/packages/whitespace "whitespace"
[go-plus]: https://atom.io/packages/go-plus "go-plus"
[秀丸]: http://hide.maruo.co.jp/software/hidemaru.html "秀まるおのホームページ(サイトー企画)－秀丸エディタ"
[サクラエディタ]: http://sakura-editor.sourceforge.net/ "サクラエディタ"
