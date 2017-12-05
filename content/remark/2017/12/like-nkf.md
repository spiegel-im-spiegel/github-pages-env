+++
title = "「nkf っぽいなにか」を作った"
date =  "2017-12-06T01:20:36+09:00"
description = "思いつきで文字エンコーディングを変換するロジックを考えていたのだが，その副産物で「nkf っぽいなにか」を作ったので，併せてリリースする。"
image = "/images/attention/remark.jpg"
tags = ["tools", "golang", "character", "encoding", "transform"]

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

思いつきで[文字エンコーディングを変換するロジックを考えていた]({{< relref "golang/detecting-character-encoding.md" >}})のだが，その副産物で「nkf っぽいなにか」を作ったので，併せてリリースする。

- [Release v0.3.0 · spiegel-im-spiegel/text · GitHub](https://github.com/spiegel-im-spiegel/text/releases/tag/v0.3.0)

リリースパッケージに含まれる gonkf がそれ。
nkf のように文字エンコーディングを指定して変換を行う。
使い方はこんな感じ。

```text
$ gonkf -h
Usage:
  gonkf [flags]
  gonkf [command]

Available Commands:
  conv        Convert character encoding of text
  guess       Guess character encoding of text
  help        Help about any command
  list        List of available character encoding
  norm        Unicode normalization
  version     Print the version number of gonkf

Flags:
  -h, --help   help for gonkf

Use "gonkf [command] --help" for more information about a command.
```

nikf と異なり，サブコマンド方式にした。
文字エンコーディング変換なら `conv` サブコマンドを使う。

```text
$ gonkf conv -h
Convert character encoding of text

Usage:
  gonkf conv [flags] [text file]

Flags:
  -d, --dst-encoding string   character encoding of destination text
  -h, --help                  help for conv
  -n, --newline string        type of newline
  -o, --output string         output file path
  -s, --src-encoding string   character encoding of source text
```

たとえば Shift-JIS のテキストを UTF-8 に変換したい場合は以下のようにする。

```text
$ gonkf conv -s sjis -d utf8 SHIFT_JIS.txt
こんにちは。世界の国から。
```

元テキストの文字エンコーディング指定を省略すると，元テキストの文字エンコーディングを推測して変換する。

```text
$ gonkf conv -d utf8 SHIFT_JIS.txt
こんにちは。世界の国から。
```

文字エンコーディングの推測のみ実行したい場合は `guess` サブコマンドを使う。

```text
$ gonkf guess SHIFT_JIS.txt
Shift_JIS
```

nkf にない機能（多分）としては Unicode 正規化ができる。
たとえば半角カナの「ﾍﾟﾝｷﾞﾝ」を事前合成形に正規化すると

```text
$ echo ﾍﾟﾝｷﾞﾝ | gonkf norm -f NFKC
ペンギン
```

となる。
Unicode 正規化については以下の記事を参照のこと。

- [Go 言語と Unicode 正規化]({{< relref "golang/unicode-normalization.md" >}})

文字エンコーディング変換で元テキストの文字エンコーディングを推測する場合，文字列が短いと誤判定する確率が跳ね上がるので注意（特に Shift-JIS と EUC）。
また ISO-2022-JP (`jis`) に変換する際，変換ロジックにバグがあるようで，文字列の末尾が改行で終わらない場合に文字セットを US-ASCII に戻す指示シーケンス（1BH 28H 42H）が出力されない[^jis1]。
末尾が改行で終わる場合は大丈夫。

[^jis1]: おそらく「行末に指示シーケンスを出力」するのではなく「改行コードの手前で指示シーケンスを出力」しているのだろう。昔はこの手の安直な実装が結構あって，電子署名の検証が valid にならなかったりと苦労した記憶があるが，今だに残ってるものなんだねぇ。まぁ ISO-2022-JP ってだいぶ廃れてきてると思うので，放置しても実害はないだろう。

nkf みたいなツールは15年以上前なら割と重宝してたけど，今はテキスト・エディタの機能でちょいちょいと変換できるので，あんまり使わなくなったよなぁ。
まぁ，今回は自作のパッケージの使い勝手を確認するための習作なので，こんなもんだろう。
