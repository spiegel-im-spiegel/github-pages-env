+++
title = "クラウドで LuaLaTeX-ja"
date =  "2017-10-02T22:13:48+09:00"
update = "2017-10-03T15:47:46+09:00"
description = "ところで『LaTeX2ε美文書作成入門』を改めて眺めていたが， Cloud LaTeX というサービスが有るらしい。pLaTeX, upLaTeX だけじゃなく LuaLaTeX にも対応しているようだ。"
tags        = [ "tex", "luatex", "cloud", "pdf", "japanese" ]

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
  mathjax = true
  mermaidjs = false
+++

[TeX Live 2017 のインストール]({{< relref "remark/2017/09/install-tex-live-2017-for-windows.md" >}})が終わったので，早速 $\mathrm{Lua\TeX}$-ja で遊ぼう。
今回の目標は

- $\mathrm{Lua\TeX}$-ja を使い PDF/A フォーマットで出力する。

である。

$\mathrm{Lua\TeX}$ は2016年にバージョン1.0がリリースされたらしい。
めでたい！ [TeX Live] 2017 に収録されている $\mathrm{Lua\TeX}$ のバージョンは以下の通り。

```text
$ luatex -v
This is LuaTeX, Version 1.0.4 (TeX Live 2017/W32TeX)

Execute  'luatex --credits'  for credits and version details.

There is NO warranty. Redistribution of this software is covered by
the terms of the GNU General Public License, version 2 or (at your option)
any later version. For more information about these matters, see the file
named COPYING and the LuaTeX source.

LuaTeX is Copyright 2017 Taco Hoekwater and the LuaTeX Team.
```

では，さっそく文書ファイルを用意する。
（「[LuaTeX-jaの使い方](https://ja.osdn.net/projects/luatex-ja/wiki/LuaTeX-ja%E3%81%AE%E4%BD%BF%E3%81%84%E6%96%B9 "LuaTeX-jaの使い方 - LuaTeX-ja Wiki - LuaTeX-ja - OSDN")」のコードを流用）

```text
\documentclass{ltjsarticle}

\begin{document}

\section{はじめてのLua\TeX-ja}
ちゃんと日本語が出るかな？

\subsection{出たかな？}
長い文章を入力するとちゃんと右端のところで折り返されるかな？
大丈夫そうな気がするけど．ちょっと不安だけど何事も挑戦だよね．

\end{document}
```

プリアンブルの記述が減ったなぁ。
これをタイプセットしてみる。

```text
$ lualatex hello.tex
```

$\mathrm{Lua\TeX}$ では直接 PDF を吐く。
出力結果はこんな感じ。

{{< fig-img src="https://farm5.staticflickr.com/4513/37448584791_9b74093e59.jpg" title="LuaLaTeX-ja (1)"  link="https://www.flickr.com/photos/spiegel/37448584791/" >}}

余談だが PDF ビュアはオープンソース製品を使おう。
個人的には [Evince] がおすすめ。

- [Get a Free Software PDF reader!](https://pdfreaders.org/)
    - [Apps/Evince - GNOME Wiki!](https://wiki.gnome.org/Apps/Evince)
    - [Evince - TeX Wiki](https://texwiki.texjp.org/?Evince)

で，出力した PDF のプロパティを [Evince] で見るとこんな感じになる。

{{< fig-img src="https://farm5.staticflickr.com/4438/23596684168_c877fd563c_o.png"  link="https://www.flickr.com/photos/spiegel/23596684168/" >}}
{{< fig-img src="https://farm5.staticflickr.com/4373/23596684238_c2f77f7534_o.png"  link="https://www.flickr.com/photos/spiegel/23596684238/" >}}

フォーマットは PDF 1.5。
[IPAex] フォントが埋め込み済みである。
最近の $\mathrm{Lua\TeX}$-ja は既定で [IPAex] フォントを使うようになっているらしい。
フォントの埋め込みを禁止するには以下の記述をプリアンブルに加える。

```text
\usepackage[noembed]{luatexja-preset}
```

出力フォーマットが PDF 1.5 なので，これを PDF/A に変更する。
PDF/A に対応するには `hyperref` パッケージを使う。

{{< highlight text "hl_lines=3-4" >}}
\documentclass{ltjsarticle}

\usepackage{hyperxmp} % XMP support with hyperref
\usepackage[pdfencoding=auto,pdfa]{hyperref} % PDF/A compatible

\hypersetup{% hyperref options (and metadata)
    pdflang={jp},
    pdftitle={はじめての LuaTeX-ja},
    pdfsubject={ちゃんとLuaTeX-jaで日本語が出るかな？},
    pdfauthor={Spiegel},
    pdfkeywords={LuaTeX-ja, PDF/A},
    pdfcopyright={Written by Spiegel on 2017, and licensed under CC-BY.},
    pdflicenseurl={http://creativecommons.org/licenses/by/4.0/},
    pdfcontacturl={http://www.baldanders.info/},
    pdfcontactcity={Hiroshima},
    pdfcontactcountry={Japan},
    pdfcontactregion={JA},
    pdfcaptionwriter={Spiegel},
    baseurl={http://text.baldanders.info/remark/2015/luatex-ja/},
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

\section{はじめてのLua\TeX-ja}
ちゃんと日本語が出るかな？

\subsection{出たかな？}
長い文章を入力するとちゃんと右端のところで折り返されるかな？
大丈夫そうな気がするけど．ちょっと不安だけど何事も挑戦だよね．

\end{document}
{{< /highlight >}}

`hyperxmp` パッケージは `hyperref` パッケージに対して XMP（Extensible Metadata Platform）拡張を行う。
これでタイプセットした結果について，先ほどと同じようにプロパティを確認する。

{{< fig-img src="https://farm5.staticflickr.com/4342/23596684308_4f3600d2b7_o.png" title="LuaLaTeX-ja (4)"  link="https://www.flickr.com/photos/spiegel/23596684308/" >}}

フォーマットが PDF/A-1b になっていることが確認できる。

{{< fig-img src="https://farm5.staticflickr.com/4433/37418815792_8de8637bdb_o.png" title="LuaLaTeX-ja (5)"  link="https://www.flickr.com/photos/spiegel/37418815792/" >}}

ライセンス情報もちゃんと付与されてるな。

ところで『[LaTeX2ε美文書作成入門]』を改めて眺めていたが， [Cloud LaTeX] というサービスがあるらしい。

- [Cloud LaTeX | Build your own LaTeX environment, in seconds](https://cloudlatex.io/)

$\mathrm{p\LaTeX}$, $\mathrm{up\LaTeX}$ だけじゃなく $\mathrm{Lua\LaTeX}$ にも対応しているみたい。
早速サインアップしてみた。

{{< fig-img src="https://farm5.staticflickr.com/4441/37192503250_58c2c8c1d7.jpg" title="Sign-up Cloud LaTeX"  link="https://www.flickr.com/photos/spiegel/37192503250/" >}}

Facebook アカウントと Twitter アカウントが使えるようだ。
まぁ，昔はともかく，[今の Twitter アカウントには一貫性も永続性もない](http://www.itmedia.co.jp/news/articles/1710/01/news035.html "「仕事にも支障が」　Twitterを凍結され、日本法人を訪れて抗議したエンジニアに聞く - ITmedia NEWS")ので認証に使うのは避けたほうがいいだろう（笑）

サインアップ後の初期画面はこんな感じ。

{{< fig-img src="https://farm5.staticflickr.com/4419/36739275654_f2c069d976.jpg" title="Cloud LaTeX (1)"  link="https://www.flickr.com/photos/spiegel/36739275654/" >}}

容量は 1 GB，プロジェクトは999個まで作成できるようだ。
バージョン管理みたいなのはできなさそうだが Dropbox と連携が可能だ。
あらかじめいくつかのテンプレートが登録されている。

{{< fig-img src="https://farm5.staticflickr.com/4411/37419240592_545c19302a.jpg" title="Templete in Cloud LaTeX"  link="https://www.flickr.com/photos/spiegel/37419240592/" >}}

学生や研究者向けのテンプレートがほとんどだが，履歴書のテンプレートもあったりする。

それでは，新しく空のプロジェクトを作ってみよう。

{{< fig-img src="https://farm5.staticflickr.com/4500/37448667041_f279e0e649.jpg" title="Cloud LaTeX (2)"  link="https://www.flickr.com/photos/spiegel/37448667041/" >}}

{{< fig-img src="https://farm5.staticflickr.com/4348/36739275764_bcc40cb274.jpg" title="Cloud LaTeX (3)"  link="https://www.flickr.com/photos/spiegel/36739275764/" >}}

プロジェクトに入って先程の文書の内容を入力する。
さらにプロジェクトの設定で処理系を選択する。

{{< fig-img src="https://farm5.staticflickr.com/4433/36739275864_0420af8d49_o.png" title="Cloud LaTeX (4)"  link="https://www.flickr.com/photos/spiegel/36739275864/" >}}

ここでは lualatex を選択する。
この状態で「コンパイル」した結果が以下の通り。

{{< fig-img src="https://farm5.staticflickr.com/4436/36739275814_59bcea74d8.jpg" title="Cloud LaTeX (5)"  link="https://www.flickr.com/photos/spiegel/36739275814/" >}}

見た目は OK かな。
生成した PDF をダウンロードしてプロパティを見てみる。

{{< fig-img src="https://farm5.staticflickr.com/4431/37448667361_24bf5b262e_o.png" title="Cloud LaTeX (6)"  link="https://www.flickr.com/photos/spiegel/37448667361/" >}}

おっ。
バージョン0.95なのか。

というわけで， [Cloud LaTeX] でも同じものが生成できることを確認した。
ちなみに lualatex は（他の処理系に比べて）めっさ遅いので「自動コンパイル」を有効にする場合には注意が必要である。

作成した文書ファイルは zip ファイルのダウンロードまたは Dropbox との連携で取得できる。
なお， zip ファイルでも Dropbox 連携でも，対象となるファイルは文書関連ファイルのみでコンパイル途中で生成されるファイルや処理結果の PDF は対象にならない。

また他ユーザとの共有は考慮されていないようだ。
せめて GitHub のプライベートリポジトリと連携できると面白いと思うのだが。
でも容量が 1 GB しかないから git 等を使ったバージョン管理は無理かな。

## ブックマーク

- [LuaTeX-ja に関する覚え書き]({{< relref "remark/2015/luatex-ja.md" >}})

[TeX Live]: http://www.tug.org/texlive/ "TeX Live - TeX Users Group"
[Evince]: https://wiki.gnome.org/Apps/Evince "Apps/Evince - GNOME Wiki!"
[IPAex]: http://ipafont.ipa.go.jp/node26#ja "IPAexフォント Ver.003.01(IPAexFont Ver.003.01) | IPAexフォント/IPAフォント"
[LaTeX2ε美文書作成入門]: http://www.amazon.co.jp/exec/obidos/ASIN/4774187054/baldandersinf-22/ "Amazon | [改訂第7版]LaTeX2ε美文書作成入門 | 奥村 晴彦, 黒木 裕介 通販"
[Cloud LaTeX]: https://cloudlatex.io/ "Cloud LaTeX | Build your own LaTeX environment, in seconds"
