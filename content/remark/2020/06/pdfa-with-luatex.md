+++
title = "LuaLaTeX で PDF/A を構成する"
date =  "2020-06-09T00:14:11+09:00"
description = "pdfx パッケージを使って PDF/A-2u を構成してみる。"
image = "/images/attention/kitten.jpg"
tags = [ "luatex", "pdf" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

今回は $\mathrm{Lua\LaTeX}$ で PDF/A を構成してみる。

## PDF/A とは

PDF/A とは自己完結性と永続性を考慮した PDF 構成で， ISO 19005 シリーズとして定義されている。
PDF/A の種別は大まかに

- PDF/A-1 : ISO 19005-1 (PDF 1.4 相当)
    - PDF/A-1a : PDF/A-1 レベルA
    - PDF/A-1b : PDF/A-1 レベルB
- PDF/A-2 : ISO 19005-2 (PDF 1.7 相当)
    - PDF/A-2a : PDF/A-2 レベルA
    - PDF/A-2b : PDF/A-2 レベルB
    - PDF/A-2u : PDF/A-2 レベルU
- PDF/A-3 : ISO 19005-3 (PDF/A-2 拡張)
    - PDF/A-2a : PDF/A-2 レベルA
    - PDF/A-2b : PDF/A-2 レベルB
    - PDF/A-2u : PDF/A-2 レベルU

といった感じに分かれている。

まず PDF/A-1 については

- カラーの再現性の保証
- フォント埋め込み
- XMPメタデータの埋め込み 

が要求され，逆に暗号化によるアクセス制限やコードの埋め込み等は禁止されている。

PDF/A-2 では PDF/A 文書の添付や圧縮機能の一部などが許容されている。
更に PDF/A-3 では任意のドキュメント・データの埋め込みが許容されているが，さすがにこれは自己完結性の観点からは要らない機能と言えるだろう。

これらの条件に加えてレベルUでは ToUnicode CMap による Unicode 対応が要求され（テキスト抽出等に必要），最高レベルのAではタグによる論理構造の埋め込みも要求される。

論理構造の埋め込みはちょっと辛いし（論理構造をきちんと設計するのは大変）そもそも今回使用する `pdfx` パッケージではレベルAの要求を完全には満たせないらしいので，レベルBまたはUの PDF/A を構成することを考える。

## 最初の LuaLaTeX テキスト

手始めに以下の $\mathrm{Lua\LaTeX}$ テキストを用意してみる。

```latex
\documentclass{ltjsarticle}
\usepackage[deluxe,nfssonly]{luatexja-preset}
\renewcommand{\emph}[1]{\textsf{\textgt{#1}}} % 強調をゴシック体と Sans Serif に変更する

\title{\emph{はじめてのLua\TeX-ja}}
\author{Spiegel}
\date{2020-06-08}

\begin{document}

\maketitle

\section{はじめてのLua\TeX-ja}

ちゃんとLua\TeX-jaで日本語が出るかな？

\subsection{出たかな？}

長い文章を入力するとちゃんと右端のところで折り返されるかな？
大丈夫そうな気がするけど．ちょっと不安だけど何事も挑戦だよね．

\end{document}
```

これを処理した結果がこんな感じ。

{{< fig-img src="./output.png" link="./output.png" width="815" >}}

ちなみに `pdfinfo` で中身を見ると，こんな感じになっている[^poppler1]。

[^poppler1]: `pdfinfo` および `pdffonts` は [Poppler] のコマンド。 Windows 版 [TeX Live] には既定で同梱されているらしい。他のプラットフォームについては「[Poppler - TeX Wiki](https://texwiki.texjp.org/?Poppler)」を参考に導入すればいいだろう。

```text
$ pdfinfo sample.pdf 
Creator:        TeX
Producer:       LuaTeX-1.12.0
CreationDate:   Mon Jun  8 20:36:09 2020 JST
ModDate:        Mon Jun  8 20:36:09 2020 JST
Tagged:         no
UserProperties: no
Suspects:       no
Form:           none
JavaScript:     no
Pages:          1
Encrypted:      no
Page size:      595.276 x 841.89 pts (A4)
Page rot:       0
File size:      26856 bytes
Optimized:      no
PDF version:    1.5
```

PDF version が 1.5 である点に注目。

## pdfx パッケージとメタデータの追加

PDF/A を構成するために `pdfx` パッケージとメタデータを追加する。
こんな感じ。

```latex {hl_lines=["1-7", "13-14"]}
% XMPメタデータ
\RequirePackage{filecontents}
\begin{filecontents*}{\jobname.xmpdata}
  \Title{はじめてのLuaTeX-ja}
  \Subject{ちゃんとLuaTeX-jaで日本語が出るかな？}
  \Author{Spiegel}
\end{filecontents*}

\documentclass{ltjsarticle}
\usepackage[deluxe,nfssonly]{luatexja-preset}
\renewcommand{\emph}[1]{\textsf{\textgt{#1}}} % 強調をゴシック体と Sans Serif に変更する

\usepackage[a-2u]{pdfx}   % PDF/A-2u を構成
\pdfvariable omitcidset=1 % LuaTeX で PDF/A-2 を作る際に必要

\title{\emph{はじめてのLua\TeX-ja}}
\author{Spiegel}
\date{2020-06-08}

\begin{document}
...
```

XMPメタデータの指定を `\documentclass` の前に記述するのがポイント。

これを処理して中身を見るとこんな感じになった。

```text {hl_lines=["2-4", "20-26"]}
$ pdfinfo sample.pdf 
Title:          はじめてのLuaTeX-ja
Subject:        ちゃんとLuaTeX-jaで日本語が出るかな？
Author:         Spiegel
Creator:        LaTeX with hyperref
Producer:       LuaTeX
CreationDate:   Mon Jun  8 20:36:09 2020 JST
ModDate:        Mon Jun  8 20:36:09 2020 JST
Tagged:         no
UserProperties: no
Suspects:       no
Form:           none
JavaScript:     no
Pages:          1
Encrypted:      no
Page size:      595.276 x 841.89 pts (A4)
Page rot:       0
File size:      40669 bytes
Optimized:      no
PDF version:    1.4
PDF subtype:    PDF/A-2u:2010
    Title:         ISO 19005 - Electronic document file format for long-term preservation (PDF/A)
    Abbreviation:  PDF/A-2
    Subtitle:      Part 2: Use of ISO 32000-1
    Standard:      ISO 19005-2
    Conformance:   Level U, Unicode support
```

Title, Subject, および Author のメタデータが追加され， PDF version が 1.4 の PDF/A-2u として構成されているのが分かる。

### pdfx パッケージのオプション

`pdfx` パッケージでは PDF/E や PDF/X 等も構成できるが，ここでは PDF/A のオプションに限って紹介しておく。

| オプション | 構成                    |
| ---------- | ----------------------- |
| `a-1a`     | PDF/A-1a : ただし不完全 |
| `a-1b`     | PDF/A-1b                |
| `a-2a`     | PDF/A-2a : ただし不完全 |
| `a-2b`     | PDF/A-2b                |
| `a-2u`     | PDF/A-2u                |
| `a-3a`     | PDF/A-3a : ただし不完全 |
| `a-3b`     | PDF/A-3b                |
| `a-3u`     | PDF/A-3u                |

レベルAは使わないほうがいいだろう（笑） 一般的には最小構成の PDF/A-1b で十分なようだ。
ただし `pdfx` パッケージはレベルB指定でも ToUnicode CMap を埋め込んだままにするらしいので，今回のように，敢えて PDF/A-2u にする手もある。

### pdfx パッケージで設定可能な XMP メタデータ項目

`filecontents` 環境

```latex
\begin{filecontents*}{\jobname.xmpdata}
  ...
\end{filecontents*}
```

内に記述する XMP メタデータ項目のうち，主なものは以下の通り。

| 項目名       | XMP 要素         | 備考                                |
| ------------ | ---------------- | ----------------------------------- |
| `\Author`    | `dc:creator`     | `\sep` で複数指定可能               |
| `\Title`     | `dc:title`       |                                     |
| `\Language`  | `dc:language`    | `ja-JP` など，`\sep` で複数指定可能 |
| `\Keywords`  | `dc:subject`     | `\sep` で複数指定可能               |
| `\Publisher` | `dc:publisher`   | `\sep` で複数指定可能               |
| `\Subject`   | `dc:description` |                                     |

著作権情報を載せる場合には以下の項目も使える[^copy1]。

[^copy1]: 基本的に © マークや “Copyright” の文言は不要。その代わり著作（権）者名と公開年と許諾条件を表記するとよい（他のメタデータで代替できるのであればそれも不要だが）。 CC Licenses のようなライセンス・ツールを使うのであれば `\CopyrightURL` 項目に（コモンズ証などの）許諾条件を記した Web ページの URL をセットすればよい。なお CC Licenses について詳しくは拙文「[改訂3版： CC Licenses について](/cc-licenses/)」を参考にどうぞ。 PDF/A はアクセス制限を禁止するので CC Licenses と相性がいいよね♪

| 項目名            | XMP 要素                 | 備考                                          |
| ----------------- | ------------------------ | --------------------------------------------- |
| `\Copyright`      | `dc:rights`              | 利用規約等                                    |
| `\CopyrightURL`   | `xmpRights:WebStatement` |                                               |
| `\Copyrighted`    | `xmpRights:Marked`       | 公有の場合は `False` をセットする             |
| `\Owner`          | `xmpRights:Owner`        | 著作権者が別にいる場合，`\sep` で複数指定可能 |
| `\CertificateURL` | `xmpRights:Certificate`  |                                               |
| `\Date`           | `dc:date`                | `YYYY-MM-DD` または `YYYY-MM`                 |
| `\Relation`       | `dc:relation`            |                                               |
| `\URLlink`        | `dc:identifier`          |                                               |

その他，設定可能な項目については [`pdfx` パッケージのドキュメント](https://www.ctan.org/pkg/pdfx)を参照のこと。

なお `filecontents` 環境の内容は `*.xmpdata` ファイルに吐き出されるが，既にファイルがある場合は上書き保存されないため，メタデータを変更したらこのファイルを削除してから再処理する必要がある。
もし `.latexmkrc` ファイルでビルドの制御を行っているなら

```perl
$clean_ext = "xmpdata";
```

の記述を追加することで

```text
$ latexmk -c
```

コマンドで，他の一時ファイルと共に `*.xmpdata` ファイルも削除してくれる。

実際に PDF ファイルに埋め込まれる XMP メタデータの内容は `pdfa.xmpi` に出力されているので参考になると思う。
また，すでに生成済みの PDF に対しては

```text
$ pdfinfo -meta sample.pdf
```

で XMP メタデータを抽出できる。

## 日本語の ToUnicode CMap はなくても大丈夫（らしい）

レベルUの要件として PDF に ToUnicode CMap が埋め込まれている必要があるが，[原ノ味フォント]のような Adobe-Japan1 フォントについては例外のようだ。

試しに [`pdf-rm-tuc`] コマンドで[原ノ味フォント]の ToUnicode CMap を削除してみる[^prt1]。

[^prt1]: [`pdf-rm-tuc`] コマンドの導入については拙文「[LuaTeX で生成した PDF から ToUnicode CMap を除去する]({{< ref "/remark/2020/06/remove-tounicode-cmap-from-pdf.md" >}})」を参考にどうぞ。

```text
$ pdf-rm-tuc --newline-before-endstream sample.pdf sample-tuc.pdf
```

PDF/A ドキュメントを処理する場合は `--newline-before-endstream` オプションが必須になるらしいので注意。

これで

```text {hl_lines=[5, 7]}
$ pdffonts sample-tuc.pdf 
name                                 type              encoding         emb sub uni object ID
------------------------------------ ----------------- ---------------- --- --- --- ---------
JKVWEI+LMRoman10-Regular             CID Type 0C       Identity-H       yes yes yes     22  0
IECCEF+HaranoAjiMincho-Regular       CID Type 0C       Identity-H       yes yes no      23  0
XKBUGM+LMSans17-Regular              CID Type 0C       Identity-H       yes yes yes     24  0
JYWTWW+HaranoAjiGothic-Regular       CID Type 0C       Identity-H       yes yes no      25  0
RHMEBC+LMRoman12-Regular             CID Type 0C       Identity-H       yes yes yes     26  0
RHDCZK+LMSans12-Regular              CID Type 0C       Identity-H       yes yes yes     27  0

$ pdfinfo sample-tuc.pdf 
Title:          はじめてのLuaTeX-ja
Subject:        ちゃんとLuaTeX-jaで日本語が出るかな？
Author:         Spiegel
Creator:        LaTeX with hyperref
Producer:       LuaTeX
CreationDate:   Mon Jun  8 20:36:09 2020 JST
ModDate:        Mon Jun  8 20:36:09 2020 JST
Tagged:         no
UserProperties: no
Suspects:       no
Form:           none
JavaScript:     no
Pages:          1
Encrypted:      no
Page size:      595.276 x 841.89 pts (A4)
Page rot:       0
File size:      39154 bytes
Optimized:      no
PDF version:    1.4
PDF subtype:    PDF/A-2u:2010
    Title:         ISO 19005 - Electronic document file format for long-term preservation (PDF/A)
    Abbreviation:  PDF/A-2
    Subtitle:      Part 2: Use of ISO 32000-1
    Standard:      ISO 19005-2
    Conformance:   Level U, Unicode support
```

のように PDF/A-2u を維持しつつ[原ノ味フォント]の ToUnicode CMap を削除することができた。
念のため，この PDF ファイルを [veraPDF] で検査してみたが “PDF/A-2U validation profile” でちゃんとパスしたので本当に問題ないのだろう。

さらに [`pdf-rm-tuc`] コマンドに `--linearize` および `--object-streams=generate` オプションを付けて実行すると

```text
$ pdf-rm-tuc --newline-before-endstream --linearize --object-streams=generate sample.pdf sample-tuc.pdf
```

PDF version 1.5 のドキュメントとして出力される。

```text {hl_lines=["19-20"]}
$ pdfinfo sample-tuc.pdf 
Title:          はじめてのLuaTeX-ja
Subject:        ちゃんとLuaTeX-jaで日本語が出るかな？
Author:         Spiegel
Creator:        LaTeX with hyperref
Producer:       LuaTeX
CreationDate:   Mon Jun  8 20:36:09 2020 JST
ModDate:        Mon Jun  8 20:36:09 2020 JST
Tagged:         no
UserProperties: no
Suspects:       no
Form:           none
JavaScript:     no
Pages:          1
Encrypted:      no
Page size:      595.276 x 841.89 pts (A4)
Page rot:       0
File size:      34742 bytes
Optimized:      yes
PDF version:    1.5
PDF subtype:    PDF/A-2u:2010
    Title:         ISO 19005 - Electronic document file format for long-term preservation (PDF/A)
    Abbreviation:  PDF/A-2
    Subtitle:      Part 2: Use of ISO 32000-1
    Standard:      ISO 19005-2
    Conformance:   Level U, Unicode support
```

よーし，うむうむ，よーし。

## ブックマーク

- [PDF/Aとはなにか | アンテナハウス PDF資料室](https://www.antenna.co.jp/pdf/reference/PDFA.htm)
- [PDF/A - Wikipedia](https://ja.wikipedia.org/wiki/PDF/A)

[TeX Live]: http://www.tug.org/texlive/ "TeX Live - TeX Users Group"
[Poppler]: https://poppler.freedesktop.org/
[原ノ味フォント]: https://github.com/trueroad/HaranoAjiFonts "trueroad/HaranoAjiFonts: 原ノ味フォント"
[`pdf-rm-tuc`]: https://github.com/trueroad/pdf-rm-tuc "trueroad/pdf-rm-tuc: Remove ToUnicode CMap from PDF"
[veraPDF]: https://verapdf.org/ "veraPDF | Industry Supported PDF/A Validation"

## 参考図書

{{% review-paapi "4774187054" %}} <!-- [改訂第7版]LaTeX2ε美文書作成入門 -->
