+++
title = "TeX Live 2017 から 2018 へのアップグレード"
date = "2018-09-24T17:24:14+09:00"
update = "2018-09-24T17:33:15+09:00"
description = "今回は update-tlmgr コマンドを使って2017年版から2018年版へアップグレードする方法を考える。"
image = "/images/attention/kitten.jpg"
tags = [ "tex", "install" ]

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

更新が滞っております。
ごめんペコン。

- [そうだ， TeX Live 2017 (for Windows) をインストールしよう！]({{< ref "/remark/2017/09/install-tex-live-2017-for-windows.md" >}})

1年前に [TeX Live] 2017 をインストールしたのだけど，現在の最新は [TeX Live] 2018 であり，2017年版のリポジトリは既にメンテナンスされていないようである。

```text
$ tlmgr update --self --all
tlmgr.pl: Remote repository is newer than local (2017 < 2018)
Cross release updates are only supported with
  update-tlmgr-latest(.sh/.exe) --update
Please see https://tug.org/texlive/upgrade.html for details.
```

したがって2018年版（[TeX Live] 2018）にアップグレードする必要がある。

手っ取り早くやるなら最新版のインストーラを取ってきて実行すればいいのだが（かなり巨大），今回は `update-tlmgr` コマンドを使って2017年版から2018年版へアップグレードする方法を考える。

詳しくは以下のリンク先を参照していただくとして，この記事では，ごく簡単にやり方を紹介する。

- [Upgrade - TeX Live - TeX Users Group](https://tug.org/texlive/upgrade.html)

## まずは前提条件

以前インストールした [TeX Live] 2017 のインストール先フォルダを `/usr/local/texlive/2017/` とする。
Windows 環境であれば `C:\texlive\2017\` などとなっているかもしれない。
適宜読み替えてほしい。

## 2018年版環境の作成

2018年版の環境を作るにはまず `/usr/local/texlive/2017/` フォルダをまるっとコピーする。

```text
$ cd /usr/local/texlive
$ cp -a 2017 2018
```

これで2018年版用に `/usr/local/texlive/2018` フォルダが作成されたはずである。
またストレージの節約のために `2018/` フォルダ以下の `tlpkg/backups/` フォルダ内のファイルを削除しておくとよい。 

```text
$ cd 2018/tlpkg/backups
$ rm *
```

## 2018年版環境への切り替え

[TeX Live] へのパスを2018年版に切り替える。
具体的には `$PATH` 環境変数の `/usr/local/texlive/2017/bin` 部分を `/usr/local/texlive/`**2018**`/bin`に変更する（環境変数変更後の反映を忘れずに）。

Windows 版ではパス構成がちょっと違っていて， `C:\texlive\2017\bin\win32` のようになっている。
同じく `2017` → `2018` と変更すればよい。

- [Windows 10でPath環境変数を設定／編集する：Tech TIPS - ＠IT](http://www.atmarkit.co.jp/ait/articles/1805/11/news035.html)

### `update-tlmgr` による環境のアップグレード

[TeX Live] の配布サイト内のページ http://mirror.ctan.org/systems/texlive/tlnet/ へアクセスして `update-tlmgr-latest.sh` をダウンロードして実行する（管理者権限での実行が必要な場合があるので注意）。

```text
$ cd /usr/local/texlive/2018
$ wget http://mirror.ctan.org/systems/texlive/tlnet/update-tlmgr-latest.sh
$ sh update-tlmgr-latest.sh - --upgrade
```

ちなみに Windows 版では `update-tlmgr-latest.exe` をダウンロードして実行すればいいようだ[^w64]。

[^w64]: Windows 版は対応してないような記述が見られたが `update-tlmgr-latest.exe` でいけそうな気がするぅ。なお Windows 64bit 用のモジュールについてはインストーラが用意されていないので，「[そうだ， TeX Live 2017 (for Windows) をインストールしよう！]({{< ref "/remark/2017/09/install-tex-live-2017-for-windows.md" >}})」の後半部分を参考に手動でインストールすること。

## 各ファイルのアップデート

ここまでできたら `tlmgr` コマンドで最新版にアップデートできるようになる（管理者権限での実行が必要な場合があるので注意）。

```text
$ tlmgr update --self --all
```

手元の Windows 環境では700個ほど更新された。
あまり使わないからアップデートもサボりがち。

$\mathrm{Lua\TeX}$ を使う場合はフォントキャッシュのアップデートも忘れずに。

```text
$ luaotfload-tool -fu
```

ちょろんと動作確認。

```text
$ luatex -v
This is LuaTeX, Version 1.07.0 (TeX Live 2018/W32TeX)

Execute  'luatex --credits'  for credits and version details.

There is NO warranty. Redistribution of this software is covered by
the terms of the GNU General Public License, version 2 or (at your option)
any later version. For more information about these matters, see the file
named COPYING and the LuaTeX source.

LuaTeX is Copyright 2018 Taco Hoekwater and the LuaTeX Team.
```

よーし，うむうむ，よーし。


あとは実際にタイプセットを行って問題ないなら2017年版は削除してしまっていいだろう。

## 【おまけの注意喚起】 Ghostscript の脆弱性に注意

Ghostscript 9.23 およびそれ以前には任意のコマンドを実行する脆弱性が存在する。
該当するバージョンを使用している方は [9.24](https://www.ghostscript.com/doc/9.24/News.htm) 以降へアップデートを。

- [Vulnerability Note VU#332928 - Ghostscript contains multiple -dSAFER sandbox bypass vulnerabilities](https://www.kb.cert.org/vuls/id/332928)
- [Ghostscript の -dSAFER オプションの脆弱性に関する注意喚起](https://www.jpcert.or.jp/at/2018/at180035.html)

アップデートは計画的に。

## ブックマーク

- [TeX Wiki](https://texwiki.texjp.org/)
    - [TeX Live - TeX Wiki](https://texwiki.texjp.org/?TeX%20Live)
    - [tlmgr - TeX Wiki](https://texwiki.texjp.org/?tlmgr)

[TeX Live]: http://www.tug.org/texlive/ "TeX Live - TeX Users Group"

## 参考図書 {#books}

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4774187054/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/51E5K7B53aL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4774187054/baldandersinf-22/">[改訂第7版]LaTeX2ε美文書作成入門</a></dt><dd>奥村 晴彦 黒木 裕介 </dd><dd>技術評論社 2017-01-24</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798118141/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798118141.09._SCTHUMBZZZ_.jpg"  alt="LaTeX2e辞典~用法・用例逆引きリファレンス (DESKTOP REFERENCE)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4535558752/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4535558752.09._SCTHUMBZZZ_.jpg"  alt="公共政策入門 ミクロ経済学的アプローチ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4320112415/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4320112415.09._SCTHUMBZZZ_.jpg"  alt="Rで楽しむ統計 (Wonderful R 1)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4000298550/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4000298550.09._SCTHUMBZZZ_.jpg"  alt="岩波データサイエンス Vol.5"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797391383/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797391383.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/積分を見つめて (数学ガールの秘密ノートシリーズ)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4000298569/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4000298569.09._SCTHUMBZZZ_.jpg"  alt="岩波データサイエンス Vol.6"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798115363/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798115363.09._SCTHUMBZZZ_.jpg"  alt="独習 LaTeX2ε"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4785315717/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4785315717.09._SCTHUMBZZZ_.jpg"  alt="具体例から学ぶ 多様体"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774193046/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774193046.09._SCTHUMBZZZ_.jpg"  alt="【改訂第3版】基礎からわかる情報リテラシー"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4768704700/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4768704700.09._SCTHUMBZZZ_.jpg"  alt="はじめて学ぶリー群 ―線型代数から始めよう"  /></a> </p>
<p class="description">ついに第7版が登場。紙の本で買って常に側に置いておくのが吉。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-09-27">2017-09-27</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
