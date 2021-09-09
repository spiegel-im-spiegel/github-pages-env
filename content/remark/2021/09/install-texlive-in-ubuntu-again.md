+++
title = "改めて TeX Live を Ubuntu に（APT を使わずに）導入する"
date =  "2021-09-09T22:20:36+09:00"
description = "LuaLaTeX でソースコードを含む文書を PDF 出力するところまで。"
image = "/images/attention/kitten.jpg"
tags = [ "tex", "linux", "ubuntu", "install", "luatex", "vscode" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

6月にパソコンを買い換えてから新しいマシンには $\mathrm{\TeX}$ 環境を入れてなかったのだが，先日『[LaTeX2ε美文書作成入門]』第8版の[読書感想文を書いた]({{< relref "./latex-primer.md" >}} "改訂第8版『LaTeX2ε美文書作成入門』を眺める")ばかりなので，調子に乗って [TeX Live] 2021 を入れてしまおうと思い立った。

『[LaTeX2ε美文書作成入門]』の付録Aには

{{< fig-quote type="markdown" title="［改訂第8版］LaTeX2ε美文書作成入門" link="https://www.amazon.co.jp/dp/B08MZ98Z1Q?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
{{% quote %}}あらかじめ `/usr/local/texlive` というディレクトリを作成し，インストールする人の権限で書き込めるようにしておくのが簡単です{{% /quote %}}
{{< /fig-quote >}}

とか書かれてあって「やっぱそーなのか」と納得してしまった。
激しくダサい気がするがしょうがないところか。

サーバ機などで複数のユーザが使う可能性がある場合は，インストール完了後に

```text
$ cd /usr/local/texlive/2021/bin/x86_64-linux/
$ sudo ./tlmgr path add
```

とすれば `/usr/local/bin/` 等へシンボリック・リンクを張ってくれる。
まぁ，今回は個人パソコンなので

```text
$ cd /usr/local/
$ sudo mkdir texlive
sudo chown -R spiegel:spiegel texlive
```

としてしおう。

## OpenPGP 公開鍵を取ってくる

[TeX Live] のサイトでは検証用に OpenPGP 公開鍵を公開している。
これを取ってきて自分の鍵束にインポートしてしまおう。

```text
$ gpg --fetch-key http://www.tug.org/texlive/files/texlive.asc
gpg: 鍵を'http://www.tug.org/texlive/files/texlive.asc'から要求
gpg: 鍵0D5E5D9106BAB6BC: 公開鍵"TeX Live Distribution <tex-live@tug.org>"をインポートしました
gpg: 処理数の合計: 1
gpg:               インポート: 1

$ gpg --list-keys 0D5E5D9106BAB6BC
pub   rsa2048 2016-03-19 [SC]
      C78B82D8C79512F79CC0D7C80D5E5D9106BAB6BC
uid           [  不明  ] TeX Live Distribution <tex-live@tug.org>
sub   rsa2048 2016-03-19 [E]
sub   rsa2048 2016-03-19 [S] [有効期限: 2022-07-27]
```

これはアレだな。
[TeX Live] がアップグレードされるたびに更新しろってことだな。

## インストーラのダウンロードと検証

では続きを。

```text
$ curl -L "https://mirror.ctan.org/systems/texlive/tlnet/install-tl-unx.tar.gz" -O
$ curl -L "https://mirror.ctan.org/systems/texlive/tlnet/install-tl-unx.tar.gz.sha512" -O
$ curl -L "https://mirror.ctan.org/systems/texlive/tlnet/install-tl-unx.tar.gz.sha512.asc" -O
$ gpg -d install-tl-unx.tar.gz.sha512.asc
gpg: 署名されたデータが'install-tl-unx.tar.gz.sha512'にあると想定します
gpg: 2021年09月08日 08時53分40秒 JSTに施された署名
gpg:                RSA鍵4CE1877E19438C70を使用
gpg: "TeX Live Distribution <tex-live@tug.org>"からの正しい署名 [不明の]
gpg: *警告*: この鍵は信用できる署名で証明されていません!
gpg:          この署名が所有者のものかどうかの検証手段がありません。
主鍵フィンガープリント: C78B 82D8 C795 12F7 9CC0  D7C8 0D5E 5D91 06BA B6BC
     副鍵フィンガープリント: D8F2 F860 57A8 57E4 2A88  106A 4CE1 877E 1943 8C70

$ sha512sum -c install-tl-unx.tar.gz.sha512
install-tl-unx.tar.gz: OK
```

なんちう回りくどい[^ds1]。
普通にダウンロードファイルに署名すりゃいいぢゃん `orz`

[^ds1]: 電子署名の検証で「`*警告*: この鍵は信用できる署名で証明されていません!`」と表示されているが，気にしなくてよい。その前の「`"TeX Live Distribution <tex-live@tug.org>"からの正しい署名`」が表示されていればOK。警告が出るのはインポートした公開鍵に自鍵で署名したり有効度を設定したりしてないからだが，直接手渡しされた鍵でもないのに安易に信用するのは危険であると言っておこう。もちろん何らかの手段で鍵と所有者が確定できるのであれば署名するなり有効度を設定するなりすればよい。

ちなみに拙作の [gnkf] を使っても

```text
$ gnkf hash -a SHA-512 -c install-tl-unx.tar.gz.sha512
install-tl-unx.tar.gz: OK
```

てな感じで検証できる。
`sha512sum` コマンドがない環境でどうぞ。
宣伝でした（笑）

## よーやくインストール開始

```text
$ tar xvf install-tl-unx.tar.gz
$ cd install-tl-20210908/
$ ./install-tl
```

リポジトリを指定する必要はないみたい。
適当に近場を探してくれてるようだ。

```text
======================> TeX Live installation procedure <=====================

======>   Letters/digits in <angle brackets> indicate   <=======
======>   menu items for actions or customizations      <=======
= help>   https://tug.org/texlive/doc/install-tl.html   <=======

 Detected platform: GNU/Linux on x86_64
 
 <B> set binary platforms: 1 out of 16

 <S> set installation scheme: scheme-full

 <C> set installation collections:
     40 collections out of 41, disk space required: 7135 MB

 <D> set directories:
   TEXDIR (the main TeX directory):
     /usr/local/texlive/2021
   TEXMFLOCAL (directory for site-wide local files):
     /usr/local/texlive/texmf-local
   TEXMFSYSVAR (directory for variable and automatically generated data):
     /usr/local/texlive/2021/texmf-var
   TEXMFSYSCONFIG (directory for local config):
     /usr/local/texlive/2021/texmf-config
   TEXMFVAR (personal directory for variable and automatically generated data):
     ~/.texlive2021/texmf-var
   TEXMFCONFIG (personal directory for local config):
     ~/.texlive2021/texmf-config
   TEXMFHOME (directory for user-specific files):
     ~/texmf

 <O> options:
   [ ] use letter size instead of A4 by default
   [X] allow execution of restricted list of programs via \write18
   [X] create all format files
   [X] install macro/font doc tree
   [X] install macro/font source tree
   [ ] create symlinks to standard directories

 <V> set up for portable installation

Actions:
 <I> start installation to hard disk
 <P> save installation profile to 'texlive.profile' and exit
 <Q> quit

Enter command: 
```

必要に応じて設定を変えて（既定のままでも無問題）問題なければ `I` を入力してインストールを開始する。
私の環境では1時間近くかかった。
お茶菓子を用意しておくか（笑）

最後に

```text
 ----------------------------------------------------------------------
 The following environment variables contain the string "tex"
 (case-independent).  If you're doing anything but adding personal
 directories to the system paths, they may well cause trouble somewhere
 while running TeX.  If you encounter problems, try unsetting them.
 Please ignore spurious matches unrelated to TeX.

    INFOPATH=:/usr/local/texlive/2020/texmf-dist/doc/info
    MANPATH=:/usr/local/texlive/2020/texmf-dist/doc/man
 ----------------------------------------------------------------------
```

と表示されれば成功である。

環境変数については `~/.profile` ファイルに

```bash
# Expand $PATH to include the directory where TeX Live applications go.
texlive_path="/usr/local/texlive/2021"
texlive_bin_path="${texlive_path}/bin/x86_64-linux"
if [ -n "${PATH##*${texlive_bin_path}}" -a -n "${PATH##*${texlive_bin_path}:*}" ]; then
    export MANPATH=${MANPATH}:${texlive_path}/texmf-dist/doc/man
    export INFOPATH=${INFOPATH}:${texlive_path}/texmf-dist/doc/info
    export PATH=${PATH}:${texlive_bin_path}
fi
```

と追記しておけばいいかな。
一応，動作確認しておく。

```text
$ lualatex -v
This is LuaHBTeX, Version 1.13.2 (TeX Live 2021)

Execute  'luahbtex --credits'  for credits and version details.

There is NO warranty. Redistribution of this software is covered by
the terms of the GNU General Public License, version 2 or (at your option)
any later version. For more information about these matters, see the file
named COPYING and the LuaTeX source.

LuaTeX is Copyright 2021 Taco Hoekwater and the LuaTeX Team.
```

よーし，うむうむ，よーし。

## 自動実行可能な外部コマンドの指定

とりあえず `shell_escape_commands` の値を変更しておく。
インストール直後は

```text
$ kpsewhich -var-value=shell_escape_commands
bibtex,bibtex8,extractbb,gregorio,kpsewhich,makeindex,repstopdf,texosquery-jre8,
```

となっているので `/usr/ocal/texlive/texmf-local/web2c/texmf.cnf` ファイルを作成し以下を記述する。

```text
shell_escape_commands = \
bibtex,bibtex8,pbibtex,jbibtex,\
extractbb,\
gregorio,\
kpsewhich,\
makeindex,mendex,\
repstopdf,epspdf,\
texosquery-jre8,\
```

これで `shell_escape_commands` の値が上書きされて

```text
$ kpsewhich -var-value=shell_escape_commands
bibtex,bibtex8,pbibtex,jbibtex,extractbb,gregorio,kpsewhich,makeindex,mendex,repstopdf,epspdf,texosquery-jre8,
```

となる。

設定を変更したら `mktexlsr` で状態を更新しておくこと。

```text
$ mktexlsr
mktexlsr: Updating /usr/local/texlive/2021/texmf-config/ls-R... 
mktexlsr: Updating /usr/local/texlive/2021/texmf-dist/ls-R... 
mktexlsr: Updating /usr/local/texlive/2021/texmf-var/ls-R... 
mktexlsr: Updating /usr/local/texlive/texmf-local/ls-R... 
mktexlsr: Done.
```

## LuaLaTeX でなんか出力してみる

そうだ。
プログラムコードを出してみよう。

- [LuaLaTeX でコードを書いてみる]({{< ref "/remark/2017/10/writing-code-with-lualatex.md" >}})

その前に [`listings`] パッケージで Go コードの syntax highlight が効くよう [`listings-golang`] を取ってくる。

```text
$ cd /usr/local/texlive/texmf-local/tex/latex/
$ git clone git@github.com:julienc91/listings-golang.git
$ mktexlsr
```

そうそう。
$\mathrm{Lua\LaTeX}$ を使うならフォントキャッシュも念の為にアップデートしておくか。

```text
$ luaotfload-tool -fu
```

用意したテキストはこんな感じ。

```latex
\documentclass{jlreq}
\usepackage[jis2004,deluxe]{luatexja-preset} % Japanese fonts
\setmonofont[AutoFakeSlant,BoldItalicFeatures={FakeSlant},Scale=MatchLowercase]{Inconsolatazi4} % use Inconsolata

\usepackage{graphicx,xcolor}
\usepackage{listings-golang}
\lstset{
    frame=single,
    basicstyle=\small\ttfamily,
    tabsize=4,
    keywordstyle=\color{red}\bfseries,
    stringstyle=\color{blue}
}

\begin{document}

\section{Go 言語による Hello World}

\begin{lstlisting}[language=Golang]
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        fmt.Println("Hello, world") //こんにちは World
    }
}
\end{lstlisting}

\section{シェルスクリプト}

\begin{lstlisting}[language=sh]
#!/bin/sh
for i in `seq 100`; do
  j="$i"
  if [ `expr $i % 3` == 0 ]; then echo -n 'Fizz'; j=''; fi
  if [ `expr $i % 5` == 0 ]; then echo -n 'Buzz'; j=''; fi
  echo "$j"
done
\end{lstlisting}

\end{document}
```

へっへっへ。
`jlreq` ドキュメントクラスを使ってみたぜ。

これを処理した結果は以下の通り。

{{< fig-img src="./sample.png" title="sample.pdf" link="./sample.pdf" width="1298" >}}

うんうん。
こんなもんだろう。

## [VS Code] に [LaTeX Workshop] を導入する

『[LaTeX2ε美文書作成入門]』でも紹介されていた [VS Code] 用の [LaTeX Workshop] を導入してみた。

```text
$ code --install-extension James-Yu.latex-workshop
```

たとえば，以下の内容で `.latexmkrc` ファイルを用意しておく。

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

したら $\mathrm{\LaTeX}$ テキストを保存するたびに上の設定でタイプセットが自動実行されて PDF まで作ってしまうわけですよ。
めっさ便利。

## ブックマーク

- [Installing TeX Live over the Internet - TeX Users Group](http://www.tug.org/texlive/acquire-netinstall.html)
- [TeX Wiki](https://texwiki.texjp.org/)

- [TeX Live を Ubuntu に（APT を使わずに）導入する]({{< ref "/remark/2019/04/install-texlive-in-ubuntu.md" >}})
- [TeX Live 2018 から 2019 へのアップグレード]({{< ref "/remark/2019/06/upgrade-texlive-from-2018-to-2019.md" >}})
- [TeX Live 2020 へのアップグレード]({{< ref "/remark/2020/04/upgrade-texlive-2020.md" >}})

[LaTeX2ε美文書作成入門]: https://www.amazon.co.jp/dp/4297117126?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "[改訂第8版]LaTeX2ε美文書作成入門 | 奥村晴彦, 黒木裕介 |本 | 通販 | Amazon"
[TeX Live]: http://www.tug.org/texlive/ "TeX Live - TeX Users Group"
[gnkf]: https://github.com/spiegel-im-spiegel/gnkf "spiegel-im-spiegel/gnkf: Network Kanji Filter by Golang"
[`listings`]: https://ctan.org/tex-archive/macros/latex/contrib/listings "CTAN: /tex-archive/macros/latex/contrib/listings"
[`listings-golang`]: https://github.com/julienc91/listings-golang "julienc91/listings-golang: Golang support for the listings package in LaTeX"
[Inconsolata]: http://levien.com/type/myfonts/inconsolata.html
[LaTeX Workshop]: https://marketplace.visualstudio.com/items?itemName=James-Yu.latex-workshop "LaTeX Workshop - Visual Studio Marketplace"
[VS Code]: https://code.visualstudio.com/ "Visual Studio Code - Code Editing. Redefined"

## 参考図書

{{% review-paapi "4297117126" %}} <!-- LaTeX2ε美文書作成入門 -->
