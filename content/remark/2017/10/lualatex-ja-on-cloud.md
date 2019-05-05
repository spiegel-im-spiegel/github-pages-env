+++
title = "クラウドで LuaLaTeX-ja"
date =  "2017-10-02T22:13:48+09:00"
description = "ところで『LaTeX2ε美文書作成入門』を改めて眺めていたが， Cloud LaTeX というサービスが有るらしい。pLaTeX, upLaTeX だけじゃなく LuaLaTeX にも対応しているようだ。"
tags        = [ "tex", "luatex", "cloud", "pdf", "japanese" ]

[scripts]
  mathjax = true
  mermaidjs = false
+++

[TeX Live 2017 のインストール]({{< ref "/remark/2017/09/install-tex-live-2017-for-windows.md" >}})が終わったので，早速 $\mathrm{Lua\TeX}$-ja で遊ぼう。
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

```tex
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

{{< fig-img src="https://photo.baldanders.info/flickr/image/37448584791_m.png" title="LuaLaTeX-ja (1)" link="https://photo.baldanders.info/flickr/37448584791/" >}}

余談だが PDF ビュアはオープンソース製品を使おう。
個人的には [Evince] がおすすめ。

- [Get a Free Software PDF reader!](https://pdfreaders.org/)
    - [Apps/Evince - GNOME Wiki!](https://wiki.gnome.org/Apps/Evince)
    - [Evince - TeX Wiki](https://texwiki.texjp.org/?Evince)

で，出力した PDF のプロパティを [Evince] で見るとこんな感じになる。

{{< fig-img src="https://photo.baldanders.info/flickr/image/23596684168_m.png" title="LuaLaTeX-ja (2)" link="https://photo.baldanders.info/flickr/23596684168/" >}}

{{< fig-img src="https://photo.baldanders.info/flickr/image/23596684238_m.png" title="LuaLaTeX-ja (3)" link="https://photo.baldanders.info/flickr/23596684238/" >}}

フォーマットは PDF 1.5。
[IPAex] フォントが埋め込み済みである。
最近の $\mathrm{Lua\TeX}$-ja は既定で [IPAex] フォントを使うようになっているらしい。
フォントの埋め込みを禁止するには以下の記述をプリアンブルに加える。

```tex
\usepackage[noembed]{luatexja-preset}
```

出力フォーマットが PDF 1.5 なので，これを PDF/A に変更する。
PDF/A に対応するには `hyperref` パッケージを使う。

{{< highlight tex "hl_lines=3-4" >}}
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

\section{はじめてのLua\TeX-ja}
ちゃんと日本語が出るかな？

\subsection{出たかな？}
長い文章を入力するとちゃんと右端のところで折り返されるかな？
大丈夫そうな気がするけど．ちょっと不安だけど何事も挑戦だよね．

\end{document}
{{< /highlight >}}

`hyperxmp` パッケージは `hyperref` パッケージに対して XMP（Extensible Metadata Platform）拡張を行う。
これでタイプセットした結果について，先ほどと同じようにプロパティを確認する。

{{< fig-img src="https://photo.baldanders.info/flickr/image/23596684308_m.png" title="LuaLaTeX-ja (4)" link="https://photo.baldanders.info/flickr/23596684308/" >}}

フォーマットが PDF/A-1b になっていることが確認できる。
ライセンス情報もちゃんと付与されてるな。

{{< fig-img src="https://photo.baldanders.info/flickr/image/37418815792_m.png" title="LuaLaTeX-ja (5)" link="https://photo.baldanders.info/flickr/37418815792/" >}}

ところで『[LaTeX2ε美文書作成入門]』を改めて眺めていたが， [Cloud LaTeX] というサービスがあるらしい。

- [Cloud LaTeX | Build your own LaTeX environment, in seconds](https://cloudlatex.io/)

$\mathrm{p\LaTeX}$, $\mathrm{up\LaTeX}$ だけじゃなく $\mathrm{Lua\LaTeX}$ にも対応しているみたい。
早速サインアップしてみた。

{{< fig-img src="https://photo.baldanders.info/flickr/image/37192503250_m.png" title="Sign-up Cloud LaTeX" link="https://photo.baldanders.info/flickr/37192503250/" >}}

Facebook アカウントと Twitter アカウントが使えるようだ。
まぁ，昔はともかく，[今の Twitter アカウントには一貫性も永続性もない](http://www.itmedia.co.jp/news/articles/1710/01/news035.html "「仕事にも支障が」　Twitterを凍結され、日本法人を訪れて抗議したエンジニアに聞く - ITmedia NEWS")ので認証に使うのは避けたほうがいいだろう（笑）

サインアップ後の初期画面はこんな感じ。

{{< fig-img src="https://photo.baldanders.info/flickr/image/36739275654_m.png" title="Cloud LaTeX (1)" link="https://photo.baldanders.info/flickr/36739275654/" >}}

容量は 1 GB，プロジェクトは999個まで作成できるようだ。
バージョン管理みたいなのはできなさそうだが Dropbox と連携が可能だ。
あらかじめいくつかのテンプレートが登録されている。

{{< fig-img src="https://photo.baldanders.info/flickr/image/37419240592_m.png" title="Templete in Cloud LaTeX" link="https://photo.baldanders.info/flickr/37419240592/" >}}

学生や研究者向けのテンプレートがほとんどだが，履歴書のテンプレートもあったりする。

それでは，新しく空のプロジェクトを作ってみよう。

{{< fig-img src="https://photo.baldanders.info/flickr/image/37448667041_m.png" title="Cloud LaTeX (2)" link="https://photo.baldanders.info/flickr/37448667041/" >}}

{{< fig-img src="https://photo.baldanders.info/flickr/image/36739275764_m.png" title="Cloud LaTeX (3)" link="https://photo.baldanders.info/flickr/36739275764/" >}}

プロジェクトに入って先程の文書の内容を入力する。
さらにプロジェクトの設定で処理系を選択する。

{{< fig-img src="https://photo.baldanders.info/flickr/image/36739275864_m.png" title="Cloud LaTeX (4)" link="https://photo.baldanders.info/flickr/36739275864/" >}}

ここでは lualatex を選択する。
この状態で「コンパイル」した結果が以下の通り。

{{< fig-img src="https://photo.baldanders.info/flickr/image/36739275814_m.png" title="Cloud LaTeX (5)" link="https://photo.baldanders.info/flickr/36739275814/" >}}

見た目は OK かな。
生成した PDF をダウンロードしてプロパティを見てみる。

{{< fig-img src="https://photo.baldanders.info/flickr/image/37448667361_m.png" title="Cloud LaTeX (6)" link="https://photo.baldanders.info/flickr/37448667361/" >}}

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

- [LuaTeX-ja に関する覚え書き]({{< ref "/remark/2015/luatex-ja.md" >}})

[TeX Live]: http://www.tug.org/texlive/ "TeX Live - TeX Users Group"
[Evince]: https://wiki.gnome.org/Apps/Evince "Apps/Evince - GNOME Wiki!"
[IPAex]: http://ipafont.ipa.go.jp/node26#ja "IPAexフォント Ver.003.01(IPAexFont Ver.003.01) | IPAexフォント/IPAフォント"
[LaTeX2ε美文書作成入門]: https://www.amazon.co.jp/exec/obidos/ASIN/4774187054/baldandersinf-22/ "Amazon | [改訂第7版]LaTeX2ε美文書作成入門 | 奥村 晴彦, 黒木 裕介 通販"
[Cloud LaTeX]: https://cloudlatex.io/ "Cloud LaTeX | Build your own LaTeX environment, in seconds"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%94%B9%E8%A8%82%E7%AC%AC7%E7%89%88-LaTeX2%CE%B5%E7%BE%8E%E6%96%87%E6%9B%B8%E4%BD%9C%E6%88%90%E5%85%A5%E9%96%80-%E5%A5%A5%E6%9D%91-%E6%99%B4%E5%BD%A6/dp/4774187054?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4774187054"><img src="https://images-fe.ssl-images-amazon.com/images/I/51E5K7B53aL._SL160_.jpg" width="127" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%94%B9%E8%A8%82%E7%AC%AC7%E7%89%88-LaTeX2%CE%B5%E7%BE%8E%E6%96%87%E6%9B%B8%E4%BD%9C%E6%88%90%E5%85%A5%E9%96%80-%E5%A5%A5%E6%9D%91-%E6%99%B4%E5%BD%A6/dp/4774187054?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4774187054">[改訂第7版]LaTeX2ε美文書作成入門</a></dt>
	<dd>奥村 晴彦, 黒木 裕介</dd>
    <dd>技術評論社 2017-01-24</dd>
    <dd>Book 大型本</dd>
    <dd>ASIN: 4774187054, EAN: 9784774187051</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">ついに第7版が登場。紙の本で買って常に側に置いておくのが吉。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-09-27">2017-09-27</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
