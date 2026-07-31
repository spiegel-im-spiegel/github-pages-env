+++
title = "TeX Live 2020 へのアップグレード"
date =  "2020-04-16T14:33:22+09:00"
description = "手元の環境では973個ほど更新されたよ…"
image = "/images/attention/kitten.jpg"
tags = [ "tex", "install", "linux", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

どうやら [TeX Live] 2020 がリリースされたようだ。
今年は早かったな（笑）

{{< fig-quote type="markdown" title="TeX Live - TeX Users Group" link="http://www.tug.org/texlive/" lang="en" >}}
{{< quote >}}TeX Live 2020 is [available over the Internet](http://www.tug.org/texlive/acquire.html) and (after production) [on DVD](http://www.tug.org/texlive/acquire-dvd.html). It was released on 10 April 2020, and [ongoing updates are available](http://www.tug.org/texlive/pkginstall.html){{< /quote >}}.
{{< /fig-quote >}}

ちうわけで，今年も [TeX Live] のアップグレードを行おう。
あっ，念のために言うと，今回のアップグレードは APT を使わず `install-tl` でインストールされていることが前提ね。

- [TeX Live を Ubuntu に（APT を使わずに）導入する]({{< ref "/remark/2019/04/install-texlive-in-ubuntu.md" >}})

## 前準備

まず `tlmgr path add` コマンドでパス設定をしている場合は

```text
$ sudo tlmgr path remove
```

で設定を消しておく。

次に 2019 の環境を 2020 へコピーする。

```text
$ cd /usr/local/texlive
$ sudo cp -a 2019 2020
$ sudo rm 2020/tlpkg/backups/*
```

`update-tlmgr` を使ってアップグレードを行う際に権限のコントロールが上手くないみたいなので `2020/` フォルダ以下のオーナーを一時的に自ユーザに書き換えておく。

```text
$ sudo chown -R username:username 2020
```

`~/.texlive2019/` ディレクトリも `~/.texlive2020/` にコピっとけばいいかな。

```text
$ cd ~
$ cp -a .texlive2019 .texlive2020
```

最後に環境変数 `PATH` や `MANPATH` 等を設定している場合はその設定を変更しておく。
たとえば  `/etc/profile.d/` ディレクトリに以下の内容を書いたファイル `texlive-paths.sh` を置いている場合

```bash
# shellcheck shell=sh

# Expand $PATH to include the directory where TeX Live applications go.
texlive_path="/usr/local/texlive/2019"
texlive_bin_path="${texlive_path}/bin/x86_64-linux"
if [ -n "${PATH##*${texlive_bin_path}}" -a -n "${PATH##*${texlive_bin_path}:*}" ]; then
    export MANPATH=${MANPATH}:${texlive_path}/texmf-dist/doc/man
    export INFOPATH=${INFOPATH}:${texlive_path}/texmf-dist/doc/info
    export PATH=${PATH}:${texlive_bin_path}
fi
```

`2019` の部分を `2020` に置き換える。

```bash {hl_lines=[4]}
# shellcheck shell=sh

# Expand $PATH to include the directory where TeX Live applications go.
texlive_path="/usr/local/texlive/2020"
texlive_bin_path="${texlive_path}/bin/x86_64-linux"
if [ -n "${PATH##*${texlive_bin_path}}" -a -n "${PATH##*${texlive_bin_path}:*}" ]; then
    export MANPATH=${MANPATH}:${texlive_path}/texmf-dist/doc/man
    export INFOPATH=${INFOPATH}:${texlive_path}/texmf-dist/doc/info
    export PATH=${PATH}:${texlive_bin_path}
fi
```

## 2020年版へのアップグレード

ほいじゃまぁ，アップグレードしますかね。
まずは `update-tlmgr` による `tlmgr` のアップグレードから。

```text
$ cd /usr/local/texlive/2020
$ wget http://mirror.ctan.org/systems/texlive/tlnet/update-tlmgr-latest.sh
$ sh update-tlmgr-latest.sh -- --upgrade
```

`tlmgr` の動作確認をしておこう。

```text
$ tlmgr version
tlmgr revision 54446 (2020-03-21 17:45:22 +0100)
tlmgr using installation: /usr/local/texlive/2020
TeX Live (https://tug.org/texlive) version 2020
```

よしよし。

次はアップグレードした `tlmgr` でアップデートを行う。

```text
$ tlmgr option repository http://mirror.ctan.org/systems/texlive/tlnet
tlmgr: setting default package repository to http://mirror.ctan.org/systems/texlive/tlnet
tlmgr: updating /usr/local/texlive/2020/tlpkg/texlive.tlpdb

$ tlmgr update --self --all
tlmgr: package repository http://ftp.yz.yamagata-u.ac.jp/pub/CTAN/systems/texlive/tlnet (verified)
tlmgr: saving backups to /usr/local/texlive/2020/tlpkg/backups
...
```

さて，お茶の時間にするか。

...手元の環境では973個ほど更新されたよ...

$\mathrm{Lua\TeX}$ を使う場合はフォントキャッシュのアップデートも忘れずに。

```text
$ luaotfload-tool -fu
```

最後に `/usr/local/texlive/2020` ディレクトリ以下のオーナーを `root` に戻す。

```text
$ cd /usr/local/texlive
$ sudo chown -R root:root 2020
```

更にパス設定（`/usr/local/bin/` 等へシンボリック・リンクを張る）を行うなら

```text
$ sudo /usr/local/texlive/2020/bin/x86_64-linux/tlmgr path add
```

とする。

## 動作確認

ちょろんと動作確認しておこう。
$\mathrm{Lua\LaTeX}$ でね。

```text {hl_lines=[2]}
$ lualatex -v
This is LuaHBTeX, Version 1.12.0 (TeX Live 2020)

Execute  'luahbtex --credits'  for credits and version details.

There is NO warranty. Redistribution of this software is covered by
the terms of the GNU General Public License, version 2 or (at your option)
any later version. For more information about these matters, see the file
named COPYING and the LuaTeX source.

LuaTeX is Copyright 2020 Taco Hoekwater and the LuaTeX Team.
```

うお！ $\mathrm{LuaHB\TeX}$ ベースになってる。
どうやら $\mathrm{Lua\LaTeX}$ では $\mathrm{LuaHB\TeX}$ ベースになる模様。
$\mathrm{Lua\TeX}$ 自体はあるようで

```text {hl_lines=[2]}
$ luatex -v
This is LuaTeX, Version 1.12.0 (TeX Live 2020)

Execute  'luatex --credits'  for credits and version details.

There is NO warranty. Redistribution of this software is covered by
the terms of the GNU General Public License, version 2 or (at your option)
any later version. For more information about these matters, see the file
named COPYING and the LuaTeX source.

LuaTeX is Copyright 2020 Taco Hoekwater and the LuaTeX Team.
```

となる。

以下のタイプセットも試してみるか。

- [LuaLaTeX でコードを書いてみる]({{< ref "/remark/2017/10/writing-code-with-lualatex.md" >}})

結果はこんな感じ。

{{< fig-img src="./sample.jpg" link="./sample.tex" width="1183" >}}

よーし，うむうむ，よーし。

## ブックマーク

- [TeX Live 2020 released | There and back again](https://www.preining.info/blog/2020/04/tex-live-2020-released/)
- [trueroad/HaranoAjiFonts: 原ノ味フォント](https://github.com/trueroad/HaranoAjiFonts) : [TeX Live] 2020 に組み込まれたそうな

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[TeX Live]: http://www.tug.org/texlive/ "TeX Live - TeX Users Group"

## 参考図書

{{% review-paapi "4297117126" %}} <!-- LaTeX2ε美文書作成入門 -->
