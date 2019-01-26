+++
title = "そうだ， TeX Live 2017 (for Windows) をインストールしよう！"
date =  "2017-10-01T17:06:19+09:00"
update =  "2017-10-02T17:35:32+09:00"
description = "久しぶりに TeX をいじってみようかな，と思い立ったり。私の環境は TeX Live 2014 のままだったので，この機会にいったん環境を削除して TeX Live 2017 を入れ直すことにした。"
tags        = [ "tex", "windows", "install" ]

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

（あれ？ 前にも似たようなタイトルで記事を書いたような...）

最近 [$\mathrm{\TeX}$ 絡みの記事を書いてる関係](/tags/tex/)で『[LaTeX2ε美文書作成入門]』を読み返しているのだが，久しぶりに $\mathrm{\TeX}$ をいじってみようかな，と思い立ったり。
私の環境は [TeX Live] 2014 のままだったので，この機会にいったん環境を削除して [TeX Live] 2017 を入れ直すことにした。

## TeX Live のインストール

Windows の場合は，以下の場所からインストーラ `install-tl-windows.exe` を取得する。

- [Installing TeX Live over the Internet - TeX Users Group](http://www.tug.org/texlive/acquire-netinstall.html)

{{< fig-quote title="Installing TeX Live over the Internet" link="http://www.tug.org/texlive/acquire-netinstall.html" lang="en" >}}
<q>For typical needs, we recommend starting the TeX Live installation by downloading <a href="http://mirror.ctan.org/systems/texlive/tlnet/install-tl-windows.exe">install-tl-windows.exe</a> for Windows (15mb), or <a href="http://mirror.ctan.org/systems/texlive/tlnet/install-tl-unx.tar.gz">install-tl-unx.tar.gz</a> (4mb) for everything else.</q>
{{< /fig-quote >}}

そのままインストーラを実行すればいいのだが，私の環境では Explorer からの起動では途中で処理が失敗してしまった。
私のようにうまくいかない場合は，コマンドプロンプト（[NYAGOS](http://www.nyaos.org/index.cgi?p=NYAGOS "NYAOS.ORG - NYAGOS") でもオッケ）から起動すれば上手くいく。

```text
$ install-tl-windows.exe
```

インストーラを起動して最初の画面がこれ。

{{< fig-img src="https://farm5.staticflickr.com/4495/23550763638_69e23884f3_o.png" title="Install TeX Live 2017 (1)"  link="https://www.flickr.com/photos/spiegel/23550763638/" >}}

特に理由がなければ “Simple install (big)” を選択すればいいようだ。
まぁ simple と言いつつ**全部**入ってしまうんだけど（笑）[^tljp1]

[^tljp1]: 日本語組版に必要なパッケージをインストールするには “Simple install (big)” が手っ取り早い。 [TeX Live] のパッケージ群を熟知している方なら “Custom install” でもいいかもしれないが。

ちなみに [TeX Live] の内容はかなり巨大で，全てをインストールするとなると相当な時間がかかる。
インストールが終わるまで待っているのは時間の浪費なので，他の作業のついでにバックグラウンドで行うようにするのがいいだろう。

“Next” ボタンを押して次の画面へ。

{{< fig-img src="https://farm5.staticflickr.com/4386/37355502656_9aa4243436_o.png" title="Install TeX Live 2017 (2)"  link="https://www.flickr.com/photos/spiegel/37355502656/" >}}

さらに “Install” ボタンを押して次の画面へ。

{{< fig-img src="https://farm5.staticflickr.com/4349/37355502726_7907d3266d_o.png" title="Install TeX Live 2017 (3)"  link="https://www.flickr.com/photos/spiegel/37355502726/" >}}

この画面でファイルを展開し，以下の画面を起動する。
この画面が本当のインストーラである。

{{< fig-img src="https://farm5.staticflickr.com/4419/23550763878_eed5744a31.jpg" title="Install TeX Live 2017 (4)"  link="https://www.flickr.com/photos/spiegel/23550763878/" >}}

既定では「既定リポジトリを変更」のチェックが外れているが，リポジトリを手動で選択できるようチェックを入れる。

「次へ」ボタンを押して次の画面へ。

{{< fig-img src="https://farm5.staticflickr.com/4426/23550763928_fae06af303.jpg" title="Install TeX Live 2017 (5)"  link="https://www.flickr.com/photos/spiegel/23550763928/" >}}

この画面でリポジトリを選択する。
日本国内であれば “Japan” のリポジトリのどれかを選択すればいいだろう。

「次へ」ボタンを押して次の画面へ。

{{< fig-img src="https://farm5.staticflickr.com/4499/23550764378_fbf4f57091.jpg" title="Install TeX Live 2017 (6a)"  link="https://www.flickr.com/photos/spiegel/23550764378/" >}}

この画面でインストール先のフォルダを決定する。
既定は `C:\texlive\2017` フォルダだが，変更するなら「変更」ボタンを押して以下のプロンプトでインストール先のフォルダを変更できる。

{{< fig-img src="https://farm5.staticflickr.com/4441/37355503256_3fe5943b3a_o.png" title="Install TeX Live 2017 (6b)"  link="https://www.flickr.com/photos/spiegel/37355503256/" >}}

インストール先を決定したら「次へ」ボタンを押して次の画面へ。

{{< fig-img src="https://farm5.staticflickr.com/4445/37355503416_a63a2dbf2f.jpg" title="Install TeX Live 2017 (7)"  link="https://www.flickr.com/photos/spiegel/37355503416/" >}}

この画面では各種オプションを指定する。
分からなければ規定値のままでも特に問題はない。

「次へ」ボタンを押して次の画面へ。

{{< fig-img src="https://farm5.staticflickr.com/4451/37355503286_ac02d01b46.jpg" title="Install TeX Live 2017 (8)"  link="https://www.flickr.com/photos/spiegel/37355503286/" >}}

これまでの指定内容で問題なければ「導入」ボタンを押してインストール処理を開始する。

{{< fig-img src="https://farm5.staticflickr.com/4355/37355503336_2ddc73822f.jpg" title="Install TeX Live 2017 (9a)"  link="https://www.flickr.com/photos/spiegel/37355503336/" >}}

上述したようにインストール処理が完了するまでにかなりの時間がかかる。
私の場合は5時間以上もかかった。
休日だったからよかったものの，インストールを開始してから朝食の準備を始めて，掃除・洗濯を終わらせて昼食を食べてもまだ終わらなかったほどだ（笑）

処理が完了すると以下の表示になる。

{{< fig-img src="https://farm5.staticflickr.com/4503/37355503396_6b5430d4fe.jpg" title="Install TeX Live 2017 (9b)"  link="https://www.flickr.com/photos/spiegel/37355503396/" >}}

これで画面下にある「終了」ボタンを押して画面を閉じればインストール完了である。

正常にインストールすれば環境変数 `PATH` にインストール先フォルダが追加される。
たとえば ，インストール先が `C:\texlive\2017` なら `C:\texlive\2017\bin\win32` が `PATH` に追加される[^pth1]。

[^pth1]: Windows には「システム環境変数」と「ユーザ環境変数」の2つがあって，ユーザに固有の環境変数は「ユーザ環境変数」のほうに設定できるようになっている。環境変数 `PATH` は「システム環境変数」と「ユーザ環境変数」の両方に存在し，両方を足し合わせたものが実際の `PATH` として機能する。私が以前にインストールした [TeX Live] 2014 では `PATH` を「システム環境変数」のほうにセットしていたが， 2017 では「ユーザ環境変数」のほうにセットしているようだ。私のように [TeX Live] を入れ直したり，バージョン別に [TeX Live] を使い分けている人は環境変数に注意した方がいいだろう。

この状態で `uplatex` コマンドを起動して動作確認をする。

```text
$ uplatex.exe -v
e-upTeX 3.14159265-p3.7.1-u1.22-161114-2.6 (utf8.uptex) (TeX Live 2017/W32TeX)
kpathsea version 6.2.3
ptexenc version 1.3.5
Copyright 2017 D.E. Knuth.
There is NO warranty.  Redistribution of this software is
covered by the terms of both the e-upTeX copyright and
the Lesser GNU General Public License.
For more information about these matters, see the file
named COPYING and the e-upTeX source.
Primary author of e-upTeX: Peter Breitenlohner.
```

これで無問題。

### 64ビット版

[TeX Live] の実行モジュールには64ビット版が存在するようだ。
インストーラが別途用意されているのではなく， zip ファイル（`tl-win64.zip`）が配布サイトにひっそりと置かれている。
日本国内の ring サーバでは以下の場所にある。

- http://core.ring.gr.jp/pub/text/TeX/ptex-win32/current/TLW64/
- http://ftp.ring.gr.jp/pub/text/TeX/ptex-win32/current/TLW64/
- http://ftp.dnsbalance.ring.gr.jp/pub/text/TeX/ptex-win32/current/TLW64/

取得した zip の中身をインストール先フォルダの直下に展開すればよい。
インストール先フォルダが `C:\texlive\2017` なら `C:\texlive\2017\bin\win64` フォルダに展開されるはずである。
この `C:\texlive\2017\bin\win64` フォルダへ `PATH` を通せばいいのだが，自動ではやってくれないので，手動で環境変数 `PATH` を設定する。

- [Windows 環境変数 Path の設定方法](http://next.matrix.jp/config-path-win7.html)
- [環境変数の設定方法](http://www.k-cube.co.jp/wakaba/server/environ.html)

32ビット版のインストールが正しく終了していれば `PATH` に32ビット版モジュールへのパス `C:\texlive\2017\bin\win32` がセットされているはずだが，この値を変更するのではなく，32ビット版と64ビット版の両方に `PATH` が通るようにする（全てのモジュールが64ビット版に置き換えられてるわけではないらしい）。

```text
SET PATH=...C:\texlive\2017\bin\win64;C:\texlive\2017\bin\win32;...
```

ここまで完了したら以下のコマンドを**必ず**実行して最新状態にしておく。

```text
$ fmtutil-sys --all
```

では，32版と同じように `uplatex` コマンドを起動して動作確認をしてみよう。

```text
$ uplatex.exe -v
e-upTeX 3.14159265-p3.7.2-u1.22-170924-2.6 (utf8.uptex) (TeX Live 2017/W32TeX)
kpathsea version 6.2.3
ptexenc version 1.3.5
Copyright 2017 D.E. Knuth.
There is NO warranty.  Redistribution of this software is
covered by the terms of both the e-upTeX copyright and
the Lesser GNU General Public License.
For more information about these matters, see the file
named COPYING and the e-upTeX source.
Primary author of e-upTeX: Peter Breitenlohner.
```

あれ？ 微妙にバージョンが違わね？ まぁこっちが新しいならいいか。

## tlmgr によるアップデート

[TeX Live] は多数のパッケージで構成され，かつ常にメンテナンスされている。
[TeX Live] の各パッケージを更新するには `tlmgr` コマンド[^tlmgr1] を使う。

[^tlmgr1]: Windows 環境では `tlmgr` コマンドはバッチファイルで提供される。実体は Perl スクリプトである。

- [tlmgr - TeX Live - TeX Users Group](http://www.tug.org/texlive/tlmgr.html)
    - [full documentation for tlmgr](http://www.tug.org/texlive/doc/tlmgr.html)
    - [tlmgr - TeX Wiki](https://texwiki.texjp.org/?tlmgr) : 日本語による簡単な解説

`tlmgr` コマンドでインストールしたパッケージを更新するには `tlmgr update` コマンドを起動する。
以下は全パッケージを更新する場合。

```
$ tlmgr update --self --all
tlmgr.pl: package repository ftp://ftp.kddilabs.jp/CTAN/systems/texlive/tlnet (verified)
tlmgr.pl: saving backups to C:/texlive/2017/tlpkg/backups
[1/9, ??:??/??:??] update: animate [3112k] (45378 -> 45431) ... done
[2/9, 00:30/02:12] update: etoc [1225k] (45380 -> 45429) ... done
[3/9, 00:41/02:09] update: luatexko [255k] (44141 -> 45434) ... done
[4/9, 00:48/02:23] update: media9 [6692k] (45129 -> 45432) ... done
[5/9, 01:16/01:32] update: ocgx2 [14k] (45337 -> 45430) ... done
[6/9, 01:21/01:38] update: xetexko [282k] (43173 -> 45435) ... done
[7/9, 01:25/01:40] update: xltabular [71k] (45412 -> 45433) ... done
[8/9, 01:28/01:43] auto-install: amscls-doc (45436) [2080k] ... done
[9/9, 01:39/01:39] update: collection-langenglish [1k] (45076 -> 45436) ... done
running mktexlsr ...
done running mktexlsr.
running mtxrun --generate ...
done running mtxrun --generate.
tlmgr.pl: package log updated: C:/texlive/2017/texmf-var/web2c/tlmgr.log
```

ちなみに `tlmgr` コマンド自身を手動で更新するには以下のインストーラをダウンロードして実行する。

- [update-tlmgr-latest.exe ](http://mirror.ctan.org/systems/texlive/tlnet/update-tlmgr-latest.exe)

インストールされているパッケージの情報を確認するには `tlmgr info` コマンドを使う。
たとえば $\mathrm{Lua\TeX}$-ja なら以下の通り。

```
$ tlmgr info luatexja
package:     luatexja
category:    Package
shortdesc:   Typeset Japanese with Lua(La)TeX
longdesc:    The package offers support for typesetting Japanese documents with LuaTeX. Either of the Plain and LaTeX2e formats may be used with the package.
installed:   Yes
revision:    45216
sizes:       src: 477k, doc: 3889k, run: 1725k
relocatable: No
cat-version: 20170904.0
cat-date:    2017-09-04 14:48:49 +0200
cat-license: bsd
cat-topics:  japanese luatex class
collection:  collection-langjapanese
```

## texmf.cnf について

[TeX Live] の設定は `texmf.cnf` ファイルに記述されている。
`texmf.cnf` の場所を確認するには `kpsewhich` コマンドを使う。

```text
$ kpsewhich -all texmf.cnf
c:/texlive/2017/texmf.cnf
c:/texlive/2017/texmf-dist/web2c/texmf.cnf
```

[TeX Live] の設定を変更するにはこれらのファイルを変更すればいいんだけど，ファイルの中身を見るに，怖くてあまり触りたくない。
そこで `texmf-local/web2c` フォルダを作ってそこに空の `texmf.cnf` ファイルを放り込む。

```text
$pwd
C:\texlive\2017
$ cd ..
$ mkdir texmf-local
$ cd texmf-local
$ mkdir web2c
$ cd web2c
$ echo>texmf.cnf
$ kpsewhich -all texmf.cnf
c:/texlive/texmf-local/web2c/texmf.cnf
c:/texlive/2017/texmf.cnf
c:/texlive/2017/texmf-dist/web2c/texmf.cnf
```

`texmf-local/web2c/texmf.cnf` はインストール先フォルダのひとつ上のフォルダに作られる点に注意。
このファイルは自前で作ったファイルなので好きに弄っていいわけだ。
ちなみに `texmf.cnf` の設定内容の優先順位は `kpsewhich` コマンド実行結果の表示順のとおり[^pth2]。
つまり `c:/texlive/2017/texmf-dist/web2c/texmf.cnf` の内容を書き換えたければ，書き換えたい部分を `c:/texlive/texmf-local/web2c/texmf.cnf` にコピペして弄ればいいのだ。

[^pth2]: `texmf.cnf` の優先順位に関する完全なリストは `kpsewhich -var-value=TEXMFCNF` コマンドから取得できる。

では，いくつか設定変更を行う。
そうそう。
もし設定変更で構成が変わったら `mktexlsr` を実行しておくこと。

```text
$ mktexlsr
mktexlsr: Updating C:/texlive/texmf-local/ls-R...
mktexlsr: Updated C:/texlive/texmf-local/ls-R.
mktexlsr: Updating C:/texlive/2017/texmf-config/ls-R...
mktexlsr: Updated C:/texlive/2017/texmf-config/ls-R.
mktexlsr: Updating C:/texlive/2017/texmf-var/ls-R...
mktexlsr: Updated C:/texlive/2017/texmf-var/ls-R.
mktexlsr: Updating C:/texlive/2017/texmf-dist/ls-R...
mktexlsr: Updated C:/texlive/2017/texmf-dist/ls-R.
mktexlsr: Done.
```

### 外部コマンドの自動実行（Restricted Shell Escape）

[TeX Live] のコマンドを実行する際，あるコマンドから別のコマンドを自動的に実行する仕組みがある。
しかし，勝手にどんなコマンドでも実行できる，というわけにはいかないので，実行できるコマンドをあらかじめ指定しておく。
自動実行可能な外部コマンドの一覧を指定するのが `shell_escape_commands` で [TeX Live] 2017 では以下のような設定になっている。

```text
shell_escape_commands = \
bibtex,bibtex8,\
extractbb,\
gregorio,\
kpsewhich,\
makeindex,\
repstopdf,\
texosquery-jre8,\
```

設定状態は `kpsewhich` コマンドでも確認できる。

```text
$ kpsewhich -var-value=shell_escape_commands
bibtex,bibtex8,extractbb,gregorio,kpsewhich,makeindex,repstopdf,texosquery-jre8,
```

実はこれには日本語用の `pbibtex` などが含まれていない。
ちなみに，いわゆる「角藤版」と呼ばれる [W32TeX](http://w32tex.org/index-ja.html) では以下のように設定されているらしい。

```text
shell_escape_commands = \
bibtex,pbibtex,jbibtex,repstopdf,epspdf,extractbb,\
makeindex,mendex,kpsewhich
```

そこで `texmf-local/web2c/texmf.cnf` ファイルにこの記述を追加して設定を上書きする。
一応 `kpsewhich` コマンドでも確認しておこう。

```text
$ kpsewhich -var-value=shell_escape_commands
bibtex,pbibtex,jbibtex,repstopdf,epspdf,extractbb,makeindex,mendex,kpsewhich
```

よしよし。

ところで `shell_escape_commands` の設定はセキュリティ上のリスクを伴うものであることをお忘れなく。
過去には `shell_escape_commands` に関する脆弱性問題もあったようだ。

- [JVNDB-2016-008562 TeX Live における任意のコマンドを実行される脆弱性 - JVN iPedia - 脆弱性対策情報データベース](http://jvndb.jvn.jp/ja/contents/2016/JVNDB-2016-008562.html)
- [【重要】TeXシステムの脆弱性の件 - マクロツイーター](http://d.hatena.ne.jp/zrbabbler/20161206/1481039449)

### 文字エンコーディングの自動判別

$\mathrm{p\TeX}$ の処理系には，入力テキストの文字エンコーディングを自動判別する機能があり `texmf.cnf` の以下の記述により有効になっている。

```text
% Guess input encoding (SJIS vs. Unicode, etc.) in pTeX and friends?
% Default is 0, to not guess.
guess_input_kanji_encoding = 1
```

これはこれで問題ないのだが，自動判別機能は完璧ではないので間違うことも稀にある。
少なくとも $\mathrm{up\TeX}$ 系は入力が UTF-8 固定なので自動判別機能は不要である。

そこで，またまた `texmf-local/web2c/texmf.cnf` ファイルに追記して以下のように設定を書き換える。

```text
guess_input_kanji_encoding.uptex = 0
guess_input_kanji_encoding.uplatex = 0
```

一応 `kpsewhich` コマンドで設定を確認しておこう。

```text
$ kpsewhich -var-value=guess_input_kanji_encoding -progname=platex
1
$ kpsewhich -var-value=guess_input_kanji_encoding -progname=uplatex
0
```

$\mathrm{p\LaTeX}$ と $\mathrm{up\LaTeX}$ で設定が異なることが確認できた。

## 漢字の埋め込み設定

DVI から PDF に変換するドライバ `dvipdfmx` 用に漢字の埋め込み設定をしておく。
設定には `updmap` コマンドを使う。

まずは設定状態を確認する。

```text
$ kanji-config-updmap status
CURRENT family for ja: ipaex
Standby family : ipa
Standby family : ms
Standby family : yu-win10
```

この状態で，たとえば [IPAex] 書体から游書体に切り替えたいのであれば以下のコマンドを実行すればよい[^upd1]。
本当に変更されると困るので `--dry-run` オプション付きで。

[^upd1]: Windows 環境でフォントを切り替える場合は `kanji-config-updmap` コマンドではなく `kanji-config-updmap-sys` コマンドを使わないとうまくいかないようだ。ちなみに `-sys` 付きのコマンドではシステム設定を変更する。 `texmf-local` フォルダに環境をまるっとコピーしたらいけるかな？ 試す気はないが。

```text
$ kanji-config-updmap-sys --dry-run yu-win10
Setting up ... ptex-yu-win10.map
updmap --sys --quiet --nomkmap --nohash -setoption jaEmbed yu-win10
updmap --sys --quiet --nomkmap --nohash -setoption jaVariant ""
updmap --sys
```

ちなみにフォントの埋め込みを行わない場合は `nofont` を指定する。

```text
$ kanji-config-updmap-sys --dry-run nofont
Setting up ... ptex-noEmbed.map
updmap --sys --quiet --nomkmap --nohash -setoption jaEmbed noEmbed
updmap --sys --quiet --nomkmap --nohash -setoption jaVariant ""
updmap --sys
```

さらに `--jis2004` オプションを付けて起動すると字形が JIS X 0213:2004 に対応したものになるのだが

```text
$ kanji-config-updmap-sys --dry-run --jis2004 ipaex
Setting up ... ptex-ipaex.map
updmap --sys --quiet --nomkmap --nohash -setoption jaEmbed ipaex
updmap --sys --quiet --nomkmap --nohash -setoption jaVariant -04
updmap --sys
```

IPA/[IPAex] 書体， MS， Windows 用の游書体は JIS X 0213:2000 （あるいはそれ以前）または JIS X 0213:2004 のどちらかの字形しか持っていないため `--jis2004` を付けても意味がない。
`--jis2004` オプションが有効なのは macOS 用の游書体（`yu-osx`），ヒラギノ書体（`hiragino-pron`），モリサワ書体（`morisawa-pr6n`），小塚書体（`kozuka-pr6`, `kozuka-pr6n`）である。

ちなみに [IPAex] フォントは2015年のバージョン3リリースで JIS X 0213:2004 の字形に対応しているらしい（異体字を含む）。

- [リリースノート(Release Note) | IPAexフォント/IPAフォント](http://ipafont.ipa.go.jp/node21#ja)

ところで游書体は Windows 環境では 8.1 から標準搭載されているが， Office ユーザ向けにフォントパックが公開され Windows 7 でも利用可能になっている。

- [Download 游ゴシック 游明朝フォントパック from Official Microsoft Download Center](https://www.microsoft.com/ja-jp/download/details.aspx?id=49116)
- [【レビュー】「Office 2016」の「游明朝」「游ゴシック」を旧Officeに追加する公式フォントパック - 窓の杜](http://forest.watch.impress.co.jp/docs/review/745200.html)

## 試しに何かタイプセットしてみる

ちうわけでいつものあの文書ですね。

- [spiegel-im-spiegel/charset_document: 「文字コードとその実装」 upLaTeX ドキュメント](https://github.com/spiegel-im-spiegel/charset_document)

いやぁ，いい加減これも内容が古すぎてアレなのだが，いまさら書き直す気にはならないし，すっかり $\mathrm{\LaTeX}$ 動作テスト用の文書になってしまった。
早速リポジトリをローカルに clone して，いきなり latexmk を起動する。

```text
$ latexmk charset.tex
...
Latexmk: All targets (charset.dvi charset.pdf) are up-to-date
```

よし！ ちゃんと動いたな。
ついでに PDF を PDF/A に変換しておこう。

```text
$ ps2pdf14 -dPDFA -dPDFACompatibilityPolicy=1 -sProcessColorModel=DeviceCMYK charset.pdf charset-pdfa.pdf
```

生成したファイルをチェックする。

{{< fig-img src="https://farm5.staticflickr.com/4475/23572647868_cb8eb6d545_o.png" title="Output PDF/A"  link="https://www.flickr.com/photos/spiegel/23572647868/" >}}

よしよし。
ちゃんと PDF/A になってるな。

## ブックマーク

- [Japanese TeX Development Community](https://github.com/texjporg)
- [TeX Live - TeX Wiki](https://texwiki.texjp.org/?TeX%20Live)
    - [TeX Live/Windows - TeX Wiki](https://texwiki.texjp.org/?TeX%20Live%2FWindows)
- [texmf.cnf ファイル - TeX Wiki](https://texwiki.texjp.org/?texmf.cnf%20%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB)
    - [外部コマンドの実行 - TeX Wiki](https://texwiki.texjp.org/?%E5%A4%96%E9%83%A8%E3%82%B3%E3%83%9E%E3%83%B3%E3%83%89%E3%81%AE%E5%AE%9F%E8%A1%8C)
- [updmap and kanji - TeX Live - TeX Users Group](http://tug.org/texlive/updmap-kanji.html)

- [TeX Live 2013 のインストールに挑戦 — Baldanders.info](https://baldanders.info/spiegel/log2/000640.shtml)
    - [TeX Live へようこそ！ — Baldanders.info](https://baldanders.info/spiegel/log2/000698.shtml)
- [TeX Live 2014 for Windows — Baldanders.info](https://baldanders.info/spiegel/log2/000730.shtml) : この記事の中で pxjahyper パッケージについて説明しているが [TeX Live] 2017 には既に含まれていた。ブラボー！
- [TeX 覚え書き（upLaTeX から PDF/A まで） — Baldanders.info](https://baldanders.info/spiegel/log2/000731.shtml)
- [LuaTeX の練習（About CC-License） — Baldanders.info](https://baldanders.info/spiegel/log2/000735.shtml)

- [LuaTeX-ja に関する覚え書き]({{< ref "/remark/2015/luatex-ja.md" >}})

[TeX Live]: http://www.tug.org/texlive/ "TeX Live - TeX Users Group"
[LaTeX2ε美文書作成入門]: https://www.amazon.co.jp/exec/obidos/ASIN/4774187054/baldandersinf-22/ "Amazon | [改訂第7版]LaTeX2ε美文書作成入門 | 奥村 晴彦, 黒木 裕介 通販"
[IPAex]: http://ipafont.ipa.go.jp/node26#ja "IPAexフォント Ver.003.01(IPAexFont Ver.003.01) | IPAexフォント/IPAフォント"

## 参考図書 {#books}

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
