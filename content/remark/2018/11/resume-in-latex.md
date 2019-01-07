+++
title = "LaTeX で履歴書を書こう"
date = "2018-11-12T17:07:15+09:00"
description = "履歴書を LaTeX で書くことのメリットは入力ファイル自体がデータベースとして機能することなのよ。"
image = "/images/attention/kitten.jpg"
tags = [ "tex" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = true
  mermaidjs = false
+++

久しぶりに履歴書を書くことになって，[宮本友介](http://www.dma.jim.osaka-u.ac.jp/view?u=2103 "研究者詳細 - 宮本 友介")さん作の「[JIS 規格 履歴書 LaTeX クラスファイル](http://koko15.hus.osaka-u.ac.jp/~yusuke/archive/rireki/)」の最新版を取ってこようとしたら，ページが無くなったのか URL が変わったのか, unreachable なのですよ。
これってもうメンテされてないってことなのか？ うーん。
しょうがないので別のスタイルファイルにしよう。

- [履歴書スタイルファイル](https://www.tamacom.com/rireki-j.html)
    - [GitHub - shigio/rireki-style: Style file for writing resume.](https://github.com/shigio/rireki-style)

FreeBSD ライセンスで公開されているので安心して利用できる。
サンプル用に `rireki.tex` が同梱されているのでこれを参考に項目を埋めていけばよい。
顔写真は 縦:横＝4:3 比率の画像データを指定する。

```text
\顔写真{photo.jpg}
```

`README.md` には EPS 形式の画像ファイルを使えってあるけど JPEG でも無問題。

タイプセット手順は以下の通り。

```text
$ platex rireki
$ dvipdfmx -z9 -p jisb5 rireki.dvi
```

そうそう。
入力は UTF-8 なのでご注意を。
いや，まぁ，このご時世に ISO-2022-JP とか Shift-JIS とか「捨て」だと思うけど。

入力が UTF-8 なら `uplatex` で処理したいよね。
`rireki.tex` のプリアンブルには

```latex
\documentclass{jarticle}
```

と書かれているので，これを

```latex
\documentclass[uplatex]{jsarticle}
```

と書き換えれば `uplatex` で処理できる。
用紙は `rireki.sty` 内部で B5 に指定されているので `\documentclass` で指定する必要はない。

ちなみに $\mathrm{Lua\LaTeX}$ で処理しようと

```latex
\documentclass{ltjsarticle}
```

と書き換えて `lualatex` コマンドを叩いてみたんだけど，途中でメモリ不足？ でコケてしまった。
うーん。

$\mathrm{\LaTeX}$ 環境なんか用意できねーよ，って方は [Cloud LaTeX](https://cloudlatex.io/ "Cloud LaTeX | Build your own LaTeX environment, in seconds") を利用する手もある。


履歴書を $\mathrm{\LaTeX}$ で書くことのメリットは，もちろん出力が綺麗なこともあるけど，入力テキストが構造化されているためデータベースとして機能し得ることなのよ。
そういう意味じゃ[宮本友介](http://www.dma.jim.osaka-u.ac.jp/view?u=2103 "研究者詳細 - 宮本 友介")さん作のクラスファイルのほうが（データベースとして）よくできてたんだけど今回のスタイルでも十分つかえる。

履歴書を手書きで要求するという企業が今だにあるらしいが $\mathrm{\LaTeX}$ で出力したものを見て手書きで清書すればよい。
ホント馬鹿らしいよね。

履歴書の用紙も文房具屋で買う必要などない。
顔写真だけ貼り付けたブランクの入力ファイルを用意してタイプセットしたものを必要なだけ印刷・コピーすればいいのだ。
自宅にプリンタがない人は[ネットプリント](http://www.printing.ne.jp/)のようなサービスを利用すればいいだろう。
私はゼロ年代中頃までに自宅のプリンタは捨てた。

試しにブランクの履歴書データを上げておく。

- [`blank.tex`](./blank.tex)
- [`blank.pdf`](./blank.pdf)

出力した PDF ファイルについて見開きにしたい場合もあるだろう（印刷時など）。
この場合は以下のページで紹介しているやり方が参考になる。

- [pdfTeX による見開きPDFの結合・分割 - TeX Alchemist Online](http://doratex.hatenablog.jp/entry/20160610/1465560005)

上の `blank.pdf` を見開き処理したものを以下に挙げておく。

- [`blank-2up.tex`](./blank-2up.tex)
- [`blank-2up.pdf`](./blank-2up.pdf)

楽しくお仕事しましょう。

## ブックマーク

- [学会スタイル等 - TeX Wiki](https://texwiki.texjp.org/?%E5%AD%A6%E4%BC%9A%E3%82%B9%E3%82%BF%E3%82%A4%E3%83%AB%E7%AD%89)
- [LaTeX pdf変換(dvipdfmx)](http://www.yamamo10.jp/~yamamoto/comp/latex/dvipdfmx/dvipdfmx.html)
- [西暦・元号対照表](http://www2.japanriver.or.jp/search_kasen/search_help/refer_year.htm)
