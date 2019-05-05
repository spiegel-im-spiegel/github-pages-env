+++
title = "LaTeX で履歴書を書こう"
date = "2018-11-12T17:07:15+09:00"
description = "履歴書を LaTeX で書くことのメリットは入力ファイル自体がデータベースとして機能することなのよ。"
image = "/images/attention/kitten.jpg"
tags = [ "tex" ]

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
