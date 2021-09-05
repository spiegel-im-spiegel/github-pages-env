+++
title = "TeX Live 2018 から 2019 へのアップグレード"
date =  "2019-06-09T14:58:35+09:00"
description = " 以前に Windows で2018年版にアップグレードしたが，やり方はだいたい同じ。"
image = "/images/attention/kitten.jpg"
tags = [ "tex", "install", "linux", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

そろそろ2019年版のアナウンスが出てるかなぁ，と思って [TeX Live] のサイトへ見に行ったら

{{< fig-quote title="TeX Live - TeX Users Group" link="http://www.tug.org/texlive/" lang="en" >}}
<q>Current release: TeX Live 2019 is <a href="http://www.tug.org/texlive/acquire.html">available over the Internet</a> and <a href="http://www.tug.org/texlive/acquire-dvd.html">on DVD</a>. It was released on 29 April 2019, and <a href="http://www.tug.org/texlive/pkginstall.html">ongoing updates are available</a>.</q>
{{< /fig-quote >}}

とか書いてくさるですよ。

はっ？ 私が[2018年版をインストールしたのって4月末]({{< ref "/remark/2019/04/install-texlive-in-ubuntu.md" >}} "TeX Live を Ubuntu に（APT を使わずに）導入する")なんだけど。

試しに `tlmgr update` してみたら

```
$ sudo tlmgr update --self --all

tlmgr: Remote repository is newer than local (2018 < 2019)
Cross release updates are only supported with
  update-tlmgr-latest(.sh/.exe) --update
Please see https://tug.org/texlive/upgrade.html for details.
```

とか言われた。
以前と表示が変わっている。
タッチの差でアップグレードされたということか `orz`

しょうがない。
[アップグレード](http://www.tug.org/texlive/upgrade.html "Upgrade - TeX Live - TeX Users Group")を行おう。
以前に [Windows で2018年版にアップグレード]({{< ref "/remark/2018/09/upgrade-texlive-from-2017-to-2018.md" >}})したが，やり方はだいたい同じ。

そうそう。
今回のアップグレードは APT を使わず `install-tl` でインストールされていることが前提ね。

- [TeX Live を Ubuntu に（APT を使わずに）導入する]({{< ref "/remark/2019/04/install-texlive-in-ubuntu.md" >}})

## 前準備

まず [TeX Live] のインストール先ディレクトリを `/usr/local/texlive` として話を進める（普通にインストールすればそうなってる筈）。
このディレクトリの中身はこんな感じになっているだろう。

```text
$ ls -F /usr/local/texlive
2018/  texmf-local/
```

最初に `2018/` ディレクトリをコピーしよう。

```text
$ cd /usr/local/texlive
$ sudo cp -a 2018 2019
$ sudo rm 2019/tlpkg/backups/*
```

このあと `update-tlmgr` を使ってアップグレードを行うのだが，権限のコントロールが上手くないみたいなので `2019/` フォルダ以下のオーナーを一時的に自ユーザに書き換えておく。

```text
$ sudo chown -R username:username 2019
```

次にパス設定の削除を行う。
`tlmgr path add` コマンドでパス設定をしている場合は

```text
$ sudo tlmgr path remove
```

で設定を消しておく。

環境変数 `PATH` や `MANPATH` 等を設定している場合はその設定を変更しておく。
たとえば  `/etc/profile.d/` ディレクトリに以下の内容を書いたファイル `texlive-paths.sh` を置いている場合

```bash
# shellcheck shell=sh

# Expand $PATH to include the directory where TeX Live applications go.
texlive_path="/usr/local/texlive/2018"
texlive_bin_path="${texlive_path}/bin/x86_64-linux"
if [ -n "${PATH##*${texlive_bin_path}}" -a -n "${PATH##*${texlive_bin_path}:*}" ]; then
    export MANPATH=${MANPATH}:${texlive_path}/texmf-dist/doc/man
    export INFOPATH=${INFOPATH}:${texlive_path}/texmf-dist/doc/info
    export PATH=${PATH}:${texlive_bin_path}
fi
```

`2018` の部分を `2019` に置き換える。

{{< highlight bash "hl_lines=4" >}}
# shellcheck shell=sh

# Expand $PATH to include the directory where TeX Live applications go.
texlive_path="/usr/local/texlive/2019"
texlive_bin_path="${texlive_path}/bin/x86_64-linux"
if [ -n "${PATH##*${texlive_bin_path}}" -a -n "${PATH##*${texlive_bin_path}:* }" ]; then
    export MANPATH=${MANPATH}:${texlive_path}/texmf-dist/doc/man
    export INFOPATH=${INFOPATH}:${texlive_path}/texmf-dist/doc/info
    export PATH=${PATH}:${texlive_bin_path}
fi
{{< /highlight >}}

最後に `~/.texlive2018/` ディレクトリがあれば `~/.texlive2019/` にコピーすればいいかな。

```text
% cd ~
$ cp -a .texlive2018 .texlive2019
```

## 2019年版へのアップグレード

ほいじゃまぁ，アップグレードしますかね。

まずは `update-tlmgr` による `tlmgr` のアップグレードから

```text
$ cd /usr/local/texlive/2019
$ wget http://mirror.ctan.org/systems/texlive/tlnet/update-tlmgr-latest.sh
$ sh update-tlmgr-latest.sh -- --upgrade
```

`tlmgr` の動作確認をしておこう。

```text
$ tlmgr version
tlmgr revision 51217 (2019-05-24 23:47:41 +0200)
tlmgr using installation: /usr/local/texlive/2019
TeX Live (http://tug.org/texlive) version 2019
```

よしよし。

次はアップグレードした `tlmgr` でアップデートを行う。

```text
$ tlmgr option repository http://mirror.ctan.org/systems/texlive/tlnet
tlmgr: setting default package repository to http://mirror.ctan.org/systems/texlive/tlnet

$ tlmgr update --self --all
tlmgr: package repository http://ftp.jaist.ac.jp/pub/CTAN/systems/texlive/tlnet (verified)
tlmgr: saving backups to /usr/local/texlive/2019/tlpkg/backups
...
```

手元の環境では520個ほど更新された。

$\mathrm{Lua\TeX}$ を使う場合はフォントキャッシュのアップデートも忘れずに。

```text
$ luaotfload-tool -fu
```

最後に `/usr/local/texlive/2019` ディレクトリ以下のオーナーを `root` に戻す。

```text
$ cd /usr/local/texlive
$ sudo chown -R root:root 2019
```

更にパス設定（`/usr/local/bin/` 等へシンボリック・リンクを張る）を行うなら

```text
$ sudo /usr/local/texlive/2019/bin/x86_64-linux/tlmgr path add
```

とする。

## 動作確認

ちょろんと動作確認しておこう。
$\mathrm{Lua\LaTeX}$ でね。

```text
$ lualatex -v
This is LuaTeX, Version 1.10.0 (TeX Live 2019)

Execute  'luatex --credits'  for credits and version details.

There is NO warranty. Redistribution of this software is covered by
the terms of the GNU General Public License, version 2 or (at your option)
any later version. For more information about these matters, see the file
named COPYING and the LuaTeX source.

LuaTeX is Copyright 2019 Taco Hoekwater and the LuaTeX Team.
```

以下のタイプセットも試してみる。

- [LuaLaTeX でコードを書いてみる]({{< ref "/remark/2017/10/writing-code-with-lualatex.md" >}})

結果はこんな感じ。

{{< fig-img src="./typesetting-by-lualatex.png" link="./typesetting-by-lualatex.png" width="1388" >}}

よーし，うむうむ，よーし。

## ブックマーク

- [TeX Live 2019 注目ポイントまとめ (1) - Acetaminophen’s diary](http://acetaminophen.hatenablog.com/entry/tl2019-01)
- [TeX Live 2019 注目ポイントまとめ (2) - Acetaminophen’s diary](http://acetaminophen.hatenablog.com/entry/tl2019-02)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[TeX Live]: http://www.tug.org/texlive/ "TeX Live - TeX Users Group"

## 参考図書

{{% review-paapi "4297117126" %}} <!-- LaTeX2ε美文書作成入門 -->
