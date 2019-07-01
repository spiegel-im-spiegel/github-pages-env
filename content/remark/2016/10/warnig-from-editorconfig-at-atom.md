+++
description = "EditorConfig と Whitespace のどちらを優先するかは人によって違うだろうが，少なくともコード書きなら EditorConfig を優先することを強くお勧めする。"
tags = [
  "atom",
  "editor",
]
draft = false
date = "2016-10-29T22:27:00+09:00"
title = "【ATOM Editor】 EditorConfig を使うなら Whitespace は不要"

[author]
  twitter = "spiegel_2007"
  flickr = "spiegel"
  name = "Spiegel"
  avatar = "/images/avatar.jpg"
  flattr = ""
  linkedin = "spiegelimspiegel"
  url = "https://baldanders.info/profile/"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  tumblr = ""
  facebook = "spiegel.im.spiegel"
+++

日本シリーズ楽しんでますか？ やっぱ日ハムは強いねぇ。
指揮官が無能で選手の勢いだけで優勝したどこぞのチームとは格が違う。
まぁ，最初の2連勝で夢を見させていただいただけでもよしとしましょう。

ところで [ATOM] の [EditorConfig] が 1.7 になって面白いワーニングを吐くようになった。

{{< fig-img src="https://photo.baldanders.info/flickr/image/30550165591_m.png" title="warning from editorconfig @atom" link="https://photo.baldanders.info/flickr/30550165591/" >}}

どうやら `trim_trailing_whitespace` と `insert_final_newline` に対応したらしいんだけど， Core Package の [whitespace] と conflict しているようだ。
この場合は [whitespace] を Disable にすれば解消する（[whitespace] は Core Package なので削除できない）。

[EditorConfig] と [whitespace] のどちらを優先するかは人によって違うだろうが，少なくともコード書きなら [EditorConfig] を優先することを強くお勧めする。

[EditorConfig] は入力テキストの文字エンコーディングや改行コードなどを設定する機能で，プロジェクトのルート・フォルダに `.editorconfig` ファイルを設置しておけば **利用環境に関係なく** 設定を統一できる。
対応しているエディタも多く， Vim や Emacs や Notepad++, Sublime Text といった定番のテキストエディタはもちろん， Eclipse, Visual Studio (Code), IntelliJ IDEA, WebStorm といった IDE (Integrated Development Environment) でも対応している。
もちろん [ATOM] も[対応している](https://atom.io/packages/editorconfig "editorconfig")。

設定可能な項目は以下の6つ。

1. `indent_style` はインデントのスタイルを指定する。 `tab` または `space` を指定する
2. `indent_size` はインデントの幅を指定する
3. `end_of_line` は改行コードを指定する。 `lf`, `cr`, `crlf` から選択できる
4. `chaset` は文字エンコーディングを指定する。 `latin1`, `utf-8`, `utf-8-bom`, `utf-16be` or `utf-16le` から選択できる。残念ながらこれ以外の文字エンコーディングについては動作不定（エディタ側の実装による）である[^u]
5. `trim_trailing_whitespace` を `true` にすると行末の空白文字を削除してくれる
6. `insert_final_newline` を `true` にするとファイルの末尾が改行文字ではない場合に補完してくれる

[^u]: まぁ Unicode が使えない国と地域以外で今更 Unicode を使わない手はないと思うけどね。[面倒なことも多い](http://qiita.com/kawasima/items/41632dbd423dc0445e14 "Shift_JIS文化からUTF-8への移行ガイド - Qiita")けど。

これらの項目をファイルの種別ごとに設定できる。
ちなみに私が仕事以外でよく使う設定はこんな感じ。

```toml
root = true

[*]
end_of_line = lf
charset = utf-8
indent_style = tab
indent_size = 4
trim_trailing_whitespace = true
insert_final_newline = true

[*.html]
insert_final_newline = false

[*.md]
indent_style = space
indent_size = 4
trim_trailing_whitespace = false
```

チームでこうした設定を決めてプロジェクトに組み込んでおけば文字エンコーディングや改行コードやインデント幅といった馬鹿らしいことで悩むことなく作業に専念できる。

- [ATOM Editor に関するメモ]({{< ref "/remark/2015/atom-editor.md" >}})


[ATOM]: https://atom.io/ "Atom"
[EditorConfig]: http://editorconfig.org/ "EditorConfig"
[whitespace]: https://atom.io/packages/whitespace
