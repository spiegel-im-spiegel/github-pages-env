+++
title = "LuaLaTeX でも履歴書を書きたい！"
date =  "2020-06-11T16:52:45+09:00"
description = "後方互換性が崩れるので pull request を投げていいか思案中。"
image = "/images/attention/kitten.jpg"
tags = [ "tex", "luatex", "pdf" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

[2018年の記事]({{< ref "/remark/2018/11/resume-in-latex.md" >}} "LaTeX で履歴書を書こう")で

- [履歴書スタイルファイル](https://www.tamacom.com/rireki-j.html)
- [shigio/rireki-style: Style file for writing resume.](https://github.com/shigio/rireki-style)

を紹介した。

その後いろいろ試してみて，サンプルの `rireki.tex` の最初の

```latex
\documentclass{jarticle}
\usepackage{rireki}
...
```

の部分を

```latex
\documentclass[b5j]{ltjsarticle}
\usepackage[deluxe,nfssonly]{luatexja-preset}
\usepackage{rireki}
...
```

と書き換えれば一応 $\mathrm{Lua\LaTeX}$ でも通ることが分かったのだが，この状態では顔写真を

```latex
\顔写真{photo.jpg}
```

と指定しても上手く貼り付けられない。

しょうがないので `rireki.sty` を眺めていたら冒頭部分に

```latex
\usepackage[dvipdfmx]{graphicx}
```

とか書かれていた。

ドライバが `dvipdfmx` 決め打ちか。
そりゃあ $\mathrm{Lua\LaTeX}$ で上手くいかんわな。

そこで `rireki.sty` の `graphicx` パッケージの記述をコメントアウトして `rireki.tex` を

```latex {hl_lines=[3]}
\documentclass[b5j]{ltjsarticle}
\usepackage[deluxe,nfssonly]{luatexja-preset}
\usepackage[luatex]{graphicx}
\usepackage{rireki}
...
```

としたら[^drv1] 問題なく顔写真を貼り付けられた。
これでますます $\mathrm{p\LaTeX}$/$\mathrm{up\LaTeX}$ は「要らない子」になるね（笑）

[^drv1]: ドライバ指定を省略して `\usepackage{graphicx}` でも可。

なお，今回改変した `rireki.sty` を使って $\mathrm{p\LaTeX}$/$\mathrm{up\LaTeX}$ で処理する場合は

```latex {hl_lines=[2]}
\documentclass[uplatex,b5j]{jsarticle}
\usepackage[dvipdfmx]{graphicx}
\usepackage{rireki}
...
```

などとすればOK[^pr1]。

[^pr1]: 後方互換性が崩れるので [pull request を投げ](https://github.com/spiegel-im-spiegel/rireki-style)ていいか思案中。

ちなみに

```perl
#!/usr/bin/env perl

# LaTeX commands
$pdflualatex = 'lualatex %O -synctex=1 %S';
$latex       = 'uplatex %O -synctex=1 %S';

# Device Driver
$dvipdf = 'dvipdfmx %O -z9 -p jisb5 -V 7 -o %D %S';

# Typeset mode (generate a PDF)
$pdf_mode = 4; # 0: do not generate a pdf , 1: using $pdflatex , 2: using $ps2pdf , 3: using $dvipdf , 4: using $pdflualatex
```

という内容で `.latexmkrc` ファイルを作成すれば

```perl
$pdf_mode = 3
```

で $\mathrm{p\LaTeX}$/$\mathrm{up\LaTeX}$ 用，

```perl
$pdf_mode = 4
```

で $\mathrm{Lua\LaTeX}$ 用の設定になる。
あとは

```text
$ latexmk
```

で `*.tex` ファイルを探して勝手に全部処理してくれる。

ところで，私が PDF 文書の紙出力に利用している[ネットプリント](http://www.printing.ne.jp/)が 2020-06-15 から値上げするらしいのだが。
今は求職活動で履歴書乱発中なんだよ。
とほほ `orz`

## ブックマーク

- [pdfTeX による見開きPDFの結合・分割 - TeX Alchemist Online](http://doratex.hatenablog.jp/entry/20160610/1465560005)
- [西暦・元号対照表](http://www2.japanriver.or.jp/search_kasen/search_help/refer_year.htm)

## 参考図書

{{% review-paapi "4774187054" %}} <!-- [改訂第7版]LaTeX2ε美文書作成入門 -->
