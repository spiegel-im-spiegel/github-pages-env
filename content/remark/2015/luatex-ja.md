+++
title = "LuaTeX-ja に関する覚え書き"
date = "2015-09-27T01:32:49+09:00"
description = "LuaTeX では pdfTeX と同等のことができ， DVI ファイルではなく PDF ファイルを直接出力する。なおかつ callback を記述することにより内部処理に割り込みをかけ機能拡張することが可能になっている。LuaTeX-ja はこの機能拡張を使って日本語組版を LuaTeX の上で実現する。"
tags = ["lua", "tex", "luatex", "japanese"]

[scripts]
  mathjax = true
  mermaidjs = false
+++

（これは[2014年9月12日に公開した記事](https://baldanders.info/mdwiki/#!luatexja.md)を再構成したものです）

$\mathrm{Lua\TeX}$ では $\mathrm{pdf\TeX}$ と同等のことができ， DVI ファイルではなく PDF ファイルを直接出力する。
なおかつ callback を記述することにより内部処理に割り込みをかけ機能拡張することが可能になっている。
$\mathrm{Lua\TeX}$-ja はこの機能拡張を使って日本語組版を $\mathrm{Lua\TeX}$ の上で実現する。

## 参考

- [LuaTeX-ja Wiki - LuaTeX-ja - SourceForge.JP](http://sourceforge.jp/projects/luatex-ja/wiki/FrontPage)
  - [LuaTeX-jaの使い方 - LuaTeX-ja Wiki - LuaTeX-ja - SourceForge.JP](http://sourceforge.jp/projects/luatex-ja/wiki/LuaTeX-ja%E3%81%AE%E4%BD%BF%E3%81%84%E6%96%B9)
- [XMP™ (Extensible Metadata Platform)仕様についてのメモ](http://www.antenna.co.jp/xml/xmllist/XMP/AboutXMP.htm)

## pTeX からの主な変更点

- 和文フォントは(小塚明朝,IPA 明朝などの)実際のフォント，和文フォントメトリック(JFMと呼ぶ )の組である
- 日本語の文書中では改行はほとんどどこでも許されるので， $\mathrm{p\TeX}$ では和文文字直後の改行は無視される(スペースが入らない)ようになっていた。しかし， $\mathrm{Lua\TeX}$-ja では $\mathrm{Lua\TeX}$ の仕様のためにこの機能は完全には実装されていない
- 2つの和文文字の間や，和文文字と欧文文字の間に入るグルー/カーン(両者をあわせて JAglue と呼ぶ)の挿入処理が 0 から書き直されている
- $\mathrm{Lua\TeX}$-ja では， $\mathrm{p\TeX}$ と同様に漢字・仮名を制御綴内に用いることができ， `\西暦` などが正しく動作するようにしている。但し，制御綴中に使える和文文字が $\mathrm{p\TeX}$, $\mathrm{up\TeX}$ と全く同じではないことに注意すること

## LuaTeX のバージョン（in TeX Live 2017）

```text
$ lualatex -v
This is LuaTeX, Version 1.0.4 (TeX Live 2017/W32TeX)

Execute  'luatex --credits'  for credits and version details.

There is NO warranty. Redistribution of this software is covered by
the terms of the GNU General Public License, version 2 or (at your option)
any later version. For more information about these matters, see the file
named COPYING and the LuaTeX source.

LuaTeX is Copyright 2017 Taco Hoekwater and the LuaTeX Team.
```

[TeX Live] で最新版を取得するには

```text
$ tlmgr update --self --all
```

などとする。

## 大雑把な解説

### 最初の一歩

最低限の $\mathrm{\LaTeX}$ 文書ファイルはこんな感じ。入力は UTF-8 で行う。

```tex
\documentclass{ltjsarticle}
% \usepackage{luatexja} % no need to use luatexja with ltjclasses or ltjsclasses classes

\begin{document}

\section{はじめてのLua\TeX-ja}

ちゃんとLua\TeX-jaで日本語が出るかな？

\subsection{出たかな？}

長い文章を入力するとちゃんと右端のところで折り返されるかな？
大丈夫そうな気がするけど．ちょっと不安だけど何事も挑戦だよね。

\end{document}
```

`ltjsarticle` クラスを用いるのであれば `\usepackage{luatexja}` の記述はなくても問題ない。

[^cls1]: `ltjsarticle` は `jsarticle` 互換クラスとして機能する。同様に `jsbook` 互換の `ltjsbook` クラスもある。

これで

```text
> luatatex hoge.tex
```

で問題なく処理できる。

### 和文フォントの埋め込み

現在のバージョンでは既定で IPAex フォントを埋め込むよう設定されている。
明示的に和文フォントを埋め込む場合は `luatexja-preset` パッケージで書体単位でまとめてフォントを指定できる。

```tex
\usepackage[sourcehan]{luatexja-preset} %和文フォントに Source Han フォントを指定する
```

プリセットオプションは以下のとおり

- `kozuka-pro`
- `kozuka-pr6`
- `kozuka-pr6n`
- `hiragino-pro`
- `hiragino-pron`
- `morisawa-pro`
- `morisawa-pr6n`
- `yu-win` （[游明朝/游ゴシック](http://blog.petitboys.com/archives/yugothic-yumincho-jiyukobo.html); Win8.1 以降）
- `yu-osx` （[游明朝/游ゴシック](http://blog.petitboys.com/archives/yugothic-yumincho-jiyukobo.html); macOS）
- `ipa`, `ipaex`, `ms`
- `ipa-hg`, `ipaex-hg`, `ms-hg` （Office 付属フォントを利用）
- `moga-mobo`
- `sourcehan` [^sh1]
- `noembed` （フォントを埋め込まない）

[^sh1]: [源ノ明朝]および[源ノ角ゴシック]の仕様例として「[TeX 日本語環境で「源ノ」フォントを使ってみた]({{< ref "/remark/2017/10/using-source-han-fonts-by-japanese-tex.md" >}})」を参考にどうぞ。

追加で以下のオプションも使用できる

- `nodeluxe`: 明朝体・ゴシック体を各 1 ウェイトで使用する（規定値）
- `deluxe`: 明朝体2ウェイト・ゴシック体3ウェイトと，丸ゴシック体を使用可能にする（otf パッケージ互換）
- `expert`: 横組専用仮名を用いる（otf パッケージ互換）
- `bold`: 「明朝の太字」をゴシック体の太字によって代替する（otf パッケージ互換）
- `90jis`: 可能ならば 90JIS 字形を使う
- `jis2004`: 可能ならば JIS2004 字形を使う
- `jis`: jfm-jis.lua を JFM として用いる（JIS フォントメトリックに近い結果が得られる）
- `no-math`: `fontspec` パッケージによる数式フォント置換が不都合な場合に指定する

### graphicx および xcolor パッケージ

`graphicx` および `xcolor` パッケージはドライバ指定なしでOK。（自動検出される）

```tex
\usepackage{graphicx,xcolor}
```

明示的に指定するのであれば `luatex` を指定する[^def1]。

```tex
\usepackage[luatex]{graphicx,xcolor}
```

[^def1]: $\mathrm{Lua\TeX}$ 0.95 以降では `luatex.def` ドライバが新設され，ドライバ依存のパッケージではこれを指定すればいいようになった。

###  hyperref パッケージ

`hyperref` パッケージも同様にドライバ指定なし大丈夫だが，そのままでは PDF の目次等が文字化けしてしまうので以下のパラメータを指定する。

```tex
\usepackage[pdfencoding=auto]{hyperref}
```

または

```tex
\usepackage[unicode=true]{hyperref}
```

更に `pdfa` オプションをつけると PDF/A-1b 準拠の PDF を出力する。

```tex
\usepackage[pdfencoding=auto,pdfa]{hyperref} % PDF/A compatible
```

`hyperref` パッケージでは PDF metadata 用に以下のオプションが指定できる。

- `baseurl`
- `pdfauthor`
- `pdfkeywords`
- `pdflang`
- `pdfproducer`
- `pdfsubject`
- `pdftitle`

###  hyperxmp パッケージ

`hyperxmp` パッケージを使うと XMP（Extensible Metadata Platform）によるメタデータを埋め込むことができる。
これは `hyperref` パッケージと組み合わせて使う。

```tex
\usepackage{hyperxmp} % XMP support with hyperref
\usepackage[pdfencoding=auto,pdfa]{hyperref} % PDF/A compatible

\hypersetup{% hyperref options (and metadata)
    pdflang={jp},
    pdftitle={はじめての LuaTeX-ja},
    pdfsubject={ちゃんとLuaTeX-jaで日本語が出るかな？},
    pdfauthor={Spiegel},
    pdfkeywords={LuaTeX-ja, PDF/A},
    pdfcopyright={Written by Spiegel on 2014,2017, and licensed under CC-BY.},
    pdflicenseurl={https://creativecommons.org/licenses/by/4.0/},
    pdfcontacturl={https://baldanders.info/},
    pdfcontactcity={Hiroshima},
    pdfcontactcountry={Japan},
    pdfcontactregion={JA},
    pdfcaptionwriter={Spiegel},
    baseurl={https://text.baldanders.info/remark/2015/luatex-ja/},
    draft=false,
    bookmarks=true,
    bookmarksnumbered=true,
    bookmarksopen=false,
    colorlinks=true,
    linkcolor=red,
    citecolor=green,
    filecolor=magenta,
    urlcolor=cyan
}
```

`hyperxmp` パッケージで追加されるパラメータは以下のとおり

- `pdfauthortitle`
- `pdfcaptionwriter`
- `pdfcontactaddress`
- `pdfcontactcity`
- `pdfcontactcountry`
- `pdfcontactemail`
- `pdfcontactphone`
- `pdfcontactpostcode`
- `pdfcontactregion`
- `pdfcontacturl`
- `pdfcopyright`
- `pdflicenseurl`
- `pdfmetalang` （ない場合は `pdflang` を参照する）

どういうわけか `hyperxmp` パッケージを使ってもいわゆる「タグ入り PDF」として Adobe Reader で認識されない。
[Evince](https://wiki.gnome.org/Apps/Evince) では著作権情報は読み取れているみたい。

{{< fig-img src="https://photo.baldanders.info/flickr/image/21544622778_m.png" title="PDF property (Evince)" link="https://photo.baldanders.info/flickr/21544622778/" >}}

###  参考文献の管理

$\mathrm{bib\TeX}$, $\mathrm{bib\LaTeX}$ は $\mathrm{p\TeX}$, $\mathrm{up\TeX}$ で使っていたものを流用できる。
ただし，入出力は UTF-8 になること。

```tex
\usepackage[backend=biber, style=numeric]{biblatex}
\addbibresource{refer.bib}
```

### LuaLaTeX 用 .latexmkrc ファイルの準備

```perl
#!/usr/bin/env perl

# LaTeX commands
$pdflatex            = 'lualatex %O -synctex=1 %S';
$latex               = 'uplatex %O -synctex=1 %S';
$latex_silent_switch = '-interaction=batchmode -c-style-errors';

# bibTeX commands
$bibtex    = 'upbibtex %O %B';
$biber     = 'biber %O --bblencoding=utf8 -u -U --output_safechars %B';
$makeindex = 'mendex %O -o %D %S';

# Device Driver
$dvipdf = 'dvipdfmx %O -z9 -V 7 -o %D %S';
$dvips = 'dvips %O -z -f %S | convbkmk -u > %D';
$ps2pdf = 'ps2pdf14 -dPDFA -dPDFACompatibilityPolicy=1 -sProcessColorModel=DeviceCMYK %O %S %D';

# Typeset mode (generate a PDF)
$pdf_mode = 1; # 0: do not generate a pdf , 1: using $pdflatex , 2: using $ps2pdf , 3: using $dvipdf

# Other configuration
$pvc_view_file_via_temporary = 0;
$max_repeat                  = 5;
```

###  最終形

```tex
\documentclass{ltjsarticle}
\usepackage[no-math,sourcehan]{luatexja-preset} %和文フォントに Source Han フォントを指定する

%% Hyperref 設定
\usepackage{hyperxmp} % XMP support with hyperref
\usepackage[pdfencoding=auto,pdfa]{hyperref} % PDF/A compatible

%% その他
\renewcommand{\emph}[1]{\textsf{\textgt{#1}}} % 強調をゴシック体と Sans Serif に変更する

%% 文書情報
\title{\emph{はじめてのLua\TeX-ja}}
\author{Spiegel}
\date{\today}

\hypersetup{% hyperref options (and metadata)
    pdflang={jp},
    pdftitle={はじめての LuaTeX-ja},
    pdfsubject={ちゃんとLuaTeX-jaで日本語が出るかな？},
    pdfauthor={Spiegel},
    pdfkeywords={LuaTeX-ja, PDF/A},
    pdfcopyright={Written by Spiegel on 2014,2017, and licensed under CC-BY.},
    pdflicenseurl={https://creativecommons.org/licenses/by/4.0/},
    pdfcontacturl={https://baldanders.info/},
    pdfcontactcity={Hiroshima},
    pdfcontactcountry={Japan},
    pdfcontactregion={JA},
    pdfcaptionwriter={Spiegel},
    baseurl={https://text.baldanders.info/remark/2015/luatex-ja/},
    draft=false,
    bookmarks=true,
    bookmarksnumbered=true,
    bookmarksopen=false,
    colorlinks=true,
    linkcolor=red,
    citecolor=green,
    filecolor=magenta,
    urlcolor=cyan
}

\begin{document}

\maketitle

\section{はじめてのLua\TeX-ja}

ちゃんとLua\TeX-jaで日本語が出るかな？

\subsection{出たかな？}

長い文章を入力するとちゃんと右端のところで折り返されるかな？
大丈夫そうな気がするけど．ちょっと不安だけど何事も挑戦だよね．

\end{document}
```

## ブックマーク

- [The Typekit Blog | Source Han Sansの紹介：オープンソースのPan-CJK書体](https://blog.typekit.com/alternate/source-han-sans-jp/)
- [源ノ明朝](https://source.typekit.com/source-han-serif/jp/)
- [TeX Live 2016 の新しい LuaTeX あれこれ - Acetaminophen’s diary](http://acetaminophen.hatenablog.com/entry/2016/04/23/141922)
- [LuaTeXで使える！便利なLaTeX関連ツール - 0番染色体](http://0-chromosome.hatenablog.jp/entry/2015/08/10/214938)
- [upLaTeXからLuaTeX-jaに移行するのに必要そうなことをまとめた - Qiita](https://qiita.com/Naughie/items/461d5a3a1cc77394834d)

[源ノ角ゴシック]: https://github.com/adobe-fonts/source-han-sans "adobe-fonts/source-han-sans: Source Han Sans | 思源黑体 | 思源黑體 | 源ノ角ゴシック | 본고딕"
[源ノ明朝]: https://github.com/adobe-fonts/source-han-serif "adobe-fonts/source-han-serif: Source Han Serif | 思源宋体 | 思源宋體 | 源ノ明朝 | 본명조"
[TeX Live]: http://www.tug.org/texlive/ "TeX Live - TeX Users Group"

## 参考図書

{{% review-paapi "4774187054" %}} <!-- [改訂第7版]LaTeX2ε美文書作成入門 -->
