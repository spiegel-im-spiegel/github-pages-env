+++
title = "GNKF: Network Kanji Filter by Golang"
date =  "2020-07-30T22:41:38+09:00"
description = "昔からある nkf コマンドの劣化コピー版と思っていただければ概ね間違いない（笑）"
image = "/images/attention/go-logo_blue.png"
tags  = [ "tools", "gnkf", "golang", "character", "encoding", "unicode", "normalization", "base64" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

- [spiegel-im-spiegel/gnkf: Network Kanji Filter by Golang](https://github.com/spiegel-im-spiegel/gnkf)

[gnkf] は，コマンドラインで動作する，特に日本語テキストを操作するためのフィルタプログラムである。
昔からある nkf コマンドの劣化コピー版と思っていただければ概ね間違いない（笑）

[![check vulns](https://github.com/spiegel-im-spiegel/gnkf/workflows/vulns/badge.svg)](https://github.com/spiegel-im-spiegel/gnkf/actions)
[![lint status](https://github.com/spiegel-im-spiegel/gnkf/workflows/lint/badge.svg)](https://github.com/spiegel-im-spiegel/gnkf/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/gnkf/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/spiegel-im-spiegel/gnkf.svg)](https://github.com/spiegel-im-spiegel/gnkf/releases/latest)

[gnkf] には以下の機能を組み込んでいる。

- `guess`: [文字エンコーディングの推測]({{< relref "#guess" >}})
- `enc`: [文字エンコーディング変換]({{< relref "#enc" >}})
- `newline`: [改行コードの一括変換]({{< relref "#newline" >}})
- `norm`: [Unicode 正規化]({{< relref "#norm" >}})
- `width`: [全角・半角変換]({{< relref "#width" >}})
- `kana`: [かなカナ変換]({{< relref "#kana" >}})
- `base64`; [BASE64 符号化]({{< relref "#base64" >}})
- `hash`; [Hash 符号化と検証]({{< relref "#hash" >}})
- `remove-bom` : [BOM の除去]({{< relref "#remove-bom" >}})
- `dump`: [16進ダンプ]({{< relref "#dump" >}})

以降からもう少し詳しく紹介する。

## 文字エンコーディングの推測{#guess}

```text
$ gnkf guess -h
Guess character encoding of the text

Usage:
  gnkf guess [flags]

Aliases:
  guess, g

Flags:
      --all           print all guesses
  -f, --file string   path of input text file
  -h, --help          help for guess

Global Flags:
      --debug   for debug
```

たとえば以下のように使う。

```text
$ echo こんにちは，世界 | gnkf g
UTF-8
```

元祖 nkf コマンドは改行コードの推測も行うが， [gnkf] では文字エンコーディングの推測のみ行う。
テキスト全文を一旦メモリ内に読み込むので，巨大ファイルを扱う場合はご注意を。

文字エンコーディングの推測には [saintfish/chardet] パッケージを使っているが，正直に言ってあまり精度がよくない（ただの plain text なら [golang.org/x/net/html/charset] よりはイケてる感じだが）。

そこで `--all` オプションを使って複数の候補を表示できるようにしている。
こんな感じで蓋然性が高い順に表示している。

```text
$ echo こんにちは，世界 | gnkf g --all
UTF-8
windows-1255
windows-1253
Big5
GB-18030
Shift_JIS
```

## 文字エンコーディング変換{#enc}

```text
$ gnkf enc -h
Convert character encoding of the text.
 Using MIME and IANA name as the character encoding name.
 Refer: http://www.iana.org/assignments/character-sets/character-sets.xhtml

Usage:
  gnkf enc [flags]

Aliases:
  enc, encoding, e

Flags:
  -d, --dst-encoding string   character encoding name of output text (default "utf-8")
  -f, --file string           path of input text file
  -g, --guess                 guess character encoding of source text
  -h, --help                  help for enc
  -o, --output string         path of output file
  -b, --remove-bom            remove BOM character in source text (UTF-8 only)
  -s, --src-encoding string   character encoding name of source text (default "utf-8")

Global Flags:
      --debug   for debug
```

`-s` および `-d` オプションで文字エンコーディング名を指定する際は [IANA の登録名称](https://www.iana.org/assignments/character-sets/character-sets.xhtml "Character Sets")を使う（大文字・小文字は区別しない）。
たとえば，こんな感じ。

```text
$ echo こんにちは，世界 | gnkf e -d shift_jis
```

ただし [golang.org/x/text] 以下のサブパッケージで変換器が用意されていない文字エンコーディングについては変換できないのであしからず。

日本語であれば [`japanese`]`.EUCJP`, [`japanese`]`.ISO2022JP`, [`japanese`]`.ShiftJIS` 等が用意されているので，それぞれ `"euc-jp"`, `"iso-2022-jp"`, `"shift_jis"` と指定すれば，いい感じに変換してくれる。

`-g` オプションを付けると `gnkf guess` コマンドと同じロジックで変換元テキストの文字エンコーディングを推測する。
が，前節で述べたように，よく間違うので過信は禁物である。

## 改行コードの一括変換{#newline}

```text
$ gnkf newline -h
Convert newline form in the text.

Usage:
  gnkf newline [flags]

Aliases:
  newline, nwln, nl

Flags:
  -f, --file string           path of input text file
  -h, --help                  help for newline
  -n, --newline-form string   newline form: [lf|cr|crlf] (default "lf")
  -o, --output string         path of output file

Global Flags:
      --debug   for debug
```

こんな感じで使う。

```
$ echo こんにちは，世界 | gnkf nl -n crlf
```

テキスト全文を一旦メモリ内に読み込んで一括変換するため，巨大ファイルを扱う場合はご注意を。
一応，複数の改行コードが混在している場合でも上手く処理してくれる（筈）

## Unicode 正規化{#norm}

```text
$ gnkf norm -h
Unicode normalization of the text (UTF-8 encoding only).

Usage:
  gnkf norm [flags]

Aliases:
  norm, normalize, nrm, nm

Flags:
  -f, --file string        path of input text file
  -h, --help               help for norm
  -k, --kangxi-radicals    normalize kangxi radicals only (with nfkc or nfkd form)
  -n, --norm-form string   Unicode normalization form: [nfc|nfd|nfkc|nfkd] (default "nfc")
  -o, --output string      path of output file
  -b, --remove-bom         remove BOM character

Global Flags:
      --debug   for debug
```

こんな感じで使う。

```text
$ echo ﾍﾟﾝｷﾞﾝ | gnkf nm -n nfkc
ペンギン
```

[Unicode 正規化は色々と副作用]({{< ref "/golang/unicode-normalization.md" >}} "Go 言語と Unicode 正規化")があるのでご注意を。
全角・半角変換については[後述]({{< relref "#width" >}})の `gnkf width` コマンドを使うのがオススメである。

### 康熙部首の正規化

NFKC または NFKD 正規化を行う際に `-k` オプションを付けることで康熙部首の正規化のみを行うようにした。

たとえば「㈱埼⽟」の「⽟（U+2f5f）」のみを正規化したい場合は

```text
$ echo ㈱埼⽟ | gnkf nm -n nfkc -k
```

などとすればよい。

## 全角・半角変換{#width}

```text
$ gnkf width -h
Convert character width in the text (UTF-8 encoding only).

Usage:
  gnkf width [flags]

Aliases:
  width, wdth, w

Flags:
  -c, --conversion-form string   conversion form: [fold|narrow|widen] (default "fold")
  -f, --file string              path of input text file
  -h, --help                     help for width
  -o, --output string            path of output file
  -b, --remove-bom               remove BOM character

Global Flags:
      --debug   for debug
```

こんな感じで使う。

```
$ echo ペンギン | gnkf w -c narrow
ﾍﾟﾝｷﾞﾝ
```

`"fold"` フォームを使うと，英数字は半角にカナは全角に，といい感じに変換してくれる。

変換には [golang.org/x/text/width] パッケージを使っているが，濁点・半濁点を含む文字や全角と半角で一対一に対応していない文字は上手く扱えないため，例外部分については内部で強制的に置換している。

また，このためにテキスト全文を一旦メモリ内に読み込んでいる。
巨大ファイルを扱う場合はご注意を。

## かなカナ変換{#kana}

```text
$ gnkf kana -h
Convert kana characters in the text.
 UTF-8 encoding only.
 "hiragana" and "katakana" forms are valid only for full-width kana character.

Usage:
  gnkf kana [flags]

Aliases:
  kana, k

Flags:
  -c, --conversion-form string   conversion form: [hiragana|katakana|chokuon] (default "katakana")
  -f, --file string              path of input text file
      --fold                     convert character width by fold form
  -h, --help                     help for kana
  -o, --output string            path of output file
  -b, --remove-bom               remove BOM character

Global Flags:
      --debug   for debug
```

こんな感じで使う。

```text
$ echo こんにちは | gnkf k -c katakana
コンニチハ
```

かなカナ変換は全角文字のみが対象だが `--fold` オプションを付けることで全角・半角変換（`"fold"` フォーム）も同時に行うことができる。

```text
$ echo 123 ｺﾝﾆﾁﾊ | gnkf k -c hiragana --fold
123 こんにちは
```

テキスト全文を一旦メモリ内に読み込むので，巨大ファイルを扱う場合はご注意を。

### 直音への変換

`"chokuon"` フォームを指定して{{< ruby "ようおん" >}}拗音{{< /ruby >}}（小さい `'ゃ'`, `'ゅ'`, `'ょ'` など）や促音（小さい `'っ'`）を直音（`'や'`, `'ゆ'`, `'よ'`, `'つ'`）に変換する。

```text
$ echo ニッポン | gnkf k -c chokuon
ニツポン
```

半角カナも変換可能。

## BASE64 符号化{#base64}

```text
$ gnkf base64 -h
Encode/Decode BASE64.

Usage:
  gnkf base64 [flags]

Aliases:
  base64, b64

Flags:
  -d, --decode          decode BASE64 string
  -f, --file string     path of input text file
  -u, --for-url         encoding/decoding defined in RFC 4648
  -h, --help            help for base64
  -p, --no-padding      no padding
  -o, --output string   path of output file

Global Flags:
      --debug   for debug
```

BASE64 符号化および復号機能。
Linux 系の `base64` コマンドや `openssl base64` コマンドの代わりに使える。

```text
$ echo Hello World | gnkf b64
SGVsbG8gV29ybGQK

$ echo SGVsbG8gV29ybGQK | gnkf b64 -d
Hello World
```

これで `base64` や `openssl` コマンドがない環境でも大丈夫。

## Hash 符号化と検証 {#hash}

```text
$ gnkf hash -h
Print or check hash value.
  Support algorithm:
  MD5, SHA-1, SHA-224, SHA-256, SHA-384, SHA-512, SHA-512/224, SHA-512/256

Usage:
  gnkf hash [flags] [file]

Aliases:
  hash, h

Flags:
  -a, --algorithm string   hash algorithm (default "SHA-256")
  -c, --check              don't fail or report status for missing files
  -h, --help               help for hash
      --ignore-missing     don't fail or report status for missing files (with check option)
      --quiet              don't print OK for each successfully verified file (with check option)

Global Flags:
      --debug   for debug
```

ハッシュ関数を使ったデータの要約と検証。
Linux 系の `sha256sum` コマンド等の代わりに使える。

たとえば SHA-256 アルゴリズムでハッシュ値を取得するには

```text
$ echo Hello World | gnkf h
d2a84f4b8b650937ec8f73cd8be2c74add5a911ba64df27458ed8229da804a26  -
```

という感じ。
SHA-256 以外で，たとえば MD5 アルゴリズムを使いたい場合はアルゴリズムを指定して

```text
$ echo Hello World | gnkf h -a md5
e59ff97941044f85df5297e1c302d260  -
```

などとする。
ファイルを指定すると

```text
$ gnkf h gnkf_0.5.0_Linux_64bit.tar.gz 
c64ef70a7ed23261b9ec9000e075ee7c6c441b604e00609b9d10078b123b7cb5  gnkf_0.5.0_Linux_64bit.tar.gz
```

といった感じに出力する。

また

```text
$ cat gnkf_0.5.0_checksums.txt 
01904595df1a438caac6bdecdfa7c9befb4aa43cdb246cbdfba7b7e3bb7153db  gnkf_0.5.0_Linux_ARMv6.tar.gz
02e695cd16cc591eb8094c999ee326a6a15817f857596c6a644cecbb416c3d74  gnkf_0.5.0_Windows_64bit.zip
322734ceb9e927503f620d9ed67fe1b1b0f8df3f09fca15b0d94786890ec07b6  gnkf_0.5.0_FreeBSD_64bit.tar.gz
33d1d7ee4a6417e56ea76cf874b7175b6e5c47ba64216ba8f2f5f347a3bb635d  gnkf_0.5.0_FreeBSD_ARM64.tar.gz
7df75e6d4b458017f4f05cd4ba7007a3852c9f95607c48d2848b383ff463d79f  gnkf_0.5.0_Windows_ARMv6.zip
95528de1b027ac292faacca8d31ae13c7678d2dd15fda3f72744eccaa414dcd4  gnkf_0.5.0_macOS_ARM64.tar.gz
9e32933a3f96bd31a5df80712589beeb37df7d5caa0d0d1de82a30660a816fd6  gnkf_0.5.0_FreeBSD_ARMv6.tar.gz
b46912edf2de327ae2ecbc86d8d27f41a2be977da4adf4e3edbc37ae88d23cd3  gnkf_0.5.0_macOS_64bit.tar.gz
c64ef70a7ed23261b9ec9000e075ee7c6c441b604e00609b9d10078b123b7cb5  gnkf_0.5.0_Linux_64bit.tar.gz
dbaf68c77197780389b714a7fbbbe553874e9f733c5126c31737767d107e0835  gnkf_0.5.0_Linux_ARM64.tar.gz
```

などとハッシュ値の一覧が用意されている場合は `-c` オプションを使って

```text
$ gnkf h --ignore-missing -c gnkf_0.5.0_checksums.txt 
gnkf_0.5.0_Linux_64bit.tar.gz: OK
WARNING in 9 items: no such file or directory
```

とすればハッシュ値が正しいかどうか検証することができる。

## BOM の除去{#remove-bom}

```text
$ gnkf remove-bom -h
Remove BOM character in UTF-8 string.

Usage:
  gnkf remove-bom [flags]

Aliases:
  remove-bom, rbom, rb

Flags:
  -f, --file string     path of input text file
  -h, --help            help for remove-bom
  -o, --output string   path of output file

Global Flags:
      --debug   for debug
```

UTF-8 テキストに BOM (Byte Order Mark; U+FEFF) が含まれている場合

```text
$ echo ﻿Hello | gnkf dump
0xef, 0xbb, 0xbf, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x0a
```

これを除去する。

```text
$ echo ﻿Hello | gnkf remove-bom | gnkf dump
0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x0a
```

先頭だけでなく，文字列中の全ての BOM を除去する。
テキスト全文を一旦メモリ内に読み込むので，巨大ファイルを扱う場合はご注意を。

## 16進ダンプ{#dump}

```text
$ gnkf dump -h
Hexadecimal view of octet data stream with C language array style.

Usage:
  gnkf dump [flags]

Aliases:
  dump, hexdump, d, hd

Flags:
  -f, --file string   path of input text file
  -h, --help          help for dump
  -u, --unicode       print by Unicode code point (UTF-8 only)

Global Flags:
      --debug   for debug
```

デバッグ用に用意した機能（笑） 以下のようにオクテット単位でダンプ表示する。

```text
$ echo ㈱埼⽟ | gnkf d
0xe3, 0x88, 0xb1, 0xe5, 0x9f, 0xbc, 0xe2, 0xbd, 0x9f, 0x0a
```

更に `-u` オプションを付けると Unicode 符号点に変換して表示してくれる。

```text
$ echo ㈱埼⽟ | gnkf d -u
0x3231, 0x57fc, 0x2f5f, 0x000a
```

これを使って，たとえば

```text
$ echo ㈱埼⽟ | gnkf nm -n nfkc -k | gnkf d -u
0x3231, 0x57fc, 0x7389, 0x000a
```

てな感じにパイプで繋いで動作確認できる。
我ながらありがたい機能だ（笑）

## ブックマーク

- [文字エンコーディング変換]({{< ref "/golang/transform-character-encoding.md" >}})
- [Go 言語と Unicode 正規化]({{< ref "/golang/unicode-normalization.md" >}})
- [文字エンコーディングの検出]({{< ref "/golang/detecting-character-encoding.md" >}})
- [Go 言語による Unicode 半角/全角変換]({{< ref "/golang/character-width-for-unicode.md" >}})
- [かなカナ変換]({{< ref "/golang/kana-conversion.md" >}})
- [こんな埼「玉」修正してやるぅ]({{< ref "/golang/unicode-kangxi-radical.md" >}})

[gnkf]: https://github.com/spiegel-im-spiegel/gnkf "spiegel-im-spiegel/gnkf: Network Kanji Filter by Golang"
[saintfish/chardet]: https://github.com/saintfish/chardet "saintfish/chardet: Charset detector library for golang derived from ICU"
[golang.org/x/net/html/charset]: https://pkg.go.dev/golang.org/x/net/html/charset "charset package · pkg.go.dev"
[golang.org/x/text]: https://pkg.go.dev/mod/golang.org/x/text "golang.org/x/text module · pkg.go.dev"
[golang.org/x/text/width]: https://pkg.go.dev/golang.org/x/text/width "width package · pkg.go.dev"
[`japanese`]: https://pkg.go.dev/golang.org/x/text@v0.3.3/encoding/japanese "japanese package · pkg.go.dev"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
