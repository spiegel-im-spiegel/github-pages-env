+++
title = "TeX Live を Ubuntu に（APT を使わずに）導入する"
date =  "2019-04-30T17:49:27+09:00"
description = "Windows と異なり，設定用の各コマンドは管理者権限（sudo）で起動する点に気をつける必要がある。"
image = "/images/attention/kitten.jpg"
tags = [ "tex", "linux", "ubuntu", "install", "font" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

[Ubuntu] 環境に [TeX Live] を導入するには APT (Advanced Package Tool) を使う方法と [TeX Live] が提供する `install-tl` を使う方法がある。 

{{< fig-quote title="Linux - TeX Wiki" link="https://texwiki.texjp.org/?Linux" >}}
<q>前者の場合は，他のパッケージと同様に統一的な管理ができますが，ディストリビューションによっては提供されているパッケージのバージョンが古いことがあります．後者の場合は，パッケージ管理システムによる管理からは外れてしまいますが，tlmgr を使って最新の状態にアップデートし続けることが可能です．</q>
{{< /fig-quote >}}

なかなか悩ましい選択だが，今回は `install-tl` を使って導入と運用をしてみよう。
なお，対象は [TeX Live] 2018 で（2019 はまだリリースされていない）。

## [TeX Live] 2018 のインストール

今回は初めてのインストールなので関係ないが，以前の環境があれば削除しておく。

```text
$ rm -rf /usr/local/texlive/2018
$ rm -rf ~/.texlive2018
```

ダウンロードページから `install-tl-unx.tar.gz` をダウンロードし中身を展開する。
その後 `install-tl` を管理者権限で起動する。

```text
$ tar xvf install-tl-unx.tar.gz

$ ls install-tl*
install-tl-unx.tar.gz

install-tl-20190227:
LICENSE.CTAN  LICENSE.TL  install-tl  release-texlive.txt  texmf-dist  tlpkg

$ cd install-tl-20190227

$ sudo ./install-tl
```

ミラーサイトのリポジトリを使うなら

```text
$ sudo ./install-tl --repository http://mirror.ctan.org/systems/texlive/tlnet/
```

とするらしい（後で気がついた）。

起動直後の画面はこんな感じ。

```text
======================> TeX Live installation procedure <=====================

======>   Letters/digits in <angle brackets> indicate   <=======
======>   menu items for actions or customizations      <=======

 Detected platform: GNU/Linux on x86_64
 
 <B> set binary platforms: 1 out of 17

 <S> set installation scheme: scheme-full

 <C> set installation collections:
     40 collections out of 41, disk space required: 5806 MB

 <D> set directories:
   TEXDIR (the main TeX directory):
     !! default location: /usr/local/texlive/2018
     !! is not writable or not allowed, please select a different one!
   TEXMFLOCAL (directory for site-wide local files):
     /usr/local/texlive/texmf-local
   TEXMFSYSVAR (directory for variable and automatically generated data):
     /usr/local/texlive/2018/texmf-var
   TEXMFSYSCONFIG (directory for local config):
     /usr/local/texlive/2018/texmf-config
   TEXMFVAR (personal directory for variable and automatically generated data):
     ~/.texlive2018/texmf-var
   TEXMFCONFIG (personal directory for local config):
     ~/.texlive2018/texmf-config
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
 <H> help
 <Q> quit

Enter command: 
```

オプション等を見て問題なければ `I` を入力してインストールを開始する。
おそらく既定のままで大丈夫。

さてここからは長いのでしばらく放置で大丈夫だろう。
今回は（[前回の経験]({{< ref "/remark/2017/09/install-tex-live-2017-for-windows.md" >}} "そうだ， TeX Live 2017 をインストールしよう！")を活かし）寝る直前にインストーラを仕掛けた。
私はやればできる子なのだ（自画自賛）。

インストールが無事に完了すれば最後に以下のメッセージが表示される。

```text
Welcome to TeX Live!


See /usr/local/texlive/2018/index.html for links to documentation.
The TeX Live web site (https://tug.org/texlive/) contains any updates and
corrections. TeX Live is a joint project of the TeX user groups around the
world; please consider supporting it by joining the group best for you. The
list of groups is available on the web at https://tug.org/usergroups.html.


Add /usr/local/texlive/2018/texmf-dist/doc/man to MANPATH.
Add /usr/local/texlive/2018/texmf-dist/doc/info to INFOPATH.
Most importantly, add /usr/local/texlive/2018/bin/x86_64-linux
to your PATH for current and future sessions.
Logfile: /usr/local/texlive/2018/install-tl.log
```

## パス設定

パス設定については手動で行ってもいいのだが `tlmgr` コマンドを使うほうが簡単である。

```text
$ sudo /usr/local/texlive/2018/bin/x86_64-linux/tlmgr path add
```

具体的には `/usr/local/texlive/2018/` 以下の各種ファイルに対して `/usr/local/bin/` 等へシンボリック・リンクを張っているらしい。
逆にリンクを削除する場合は

```text
$ sudo tlmgr path remove
```

とすれば安全に処理できる。

{{< div-box type="markdown" >}}

### 環境変数の指定

[TeX Live] は1年毎にアップグレードされるし，その度にリンクを切った張ったは微妙に嫌だったので， `/etc/profile.d/` ディレクトリに以下の内容を書いたファイル `texlive-paths.sh` を置いてみた（ファイル名は適当）。

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

ログイン時にこのファイルが実行され環境変数がセットされる。
この方法の欠点は管理者権限での実行時にはパスが通らないため

```text
$ sudo tlmgr version
sudo: tlmgr: コマンドが見つかりません
```

と言われてしまう点だろう。
普段使いなら問題ないんだけどね。
結局，シンボリック・リンクを張ったほうがいいみたい。

[TeX Live] を削除する際は `tlmgr path remove` でリンクを削除するのを忘れずに。

[TeX Live]: http://www.tug.org/texlive/ "TeX Live - TeX Users Group"
{{< /div-box >}}

## [TeX Live] をアップデートしようとしたが...

まずは参照するリポジトリを指定をしておこう。

```text
$ sudo tlmgr option repository http://mirror.ctan.org/systems/texlive/tlnet
tlmgr: setting default package repository to http://mirror.ctan.org/systems/texlive/tlnet
```

このあと `tlmgr` でアップデートを行おうとしたが

```text
$ sudo tlmgr update --self --all
TeX Live 2018 is frozen forever and will no
longer be updated.  This happens in preparation for a new release.

If you're interested in helping to pretest the new release (when
pretests are available), please read https://tug.org/texlive/pretest.html.
Otherwise, just wait, and the new release will be ready in due time.
tlmgr: package repository http://ctan.math.washington.edu/tex-archive/systems/texlive/tlnet (verified)
tlmgr: saving backups to /usr/local/texlive/2018/tlpkg/backups
```

と言われてしまった。
そうか。
夏には 2019 が出るから現行版の更新は既に凍結されているのか。
迂闊だった `orz`

## 各種設定変更

Windows と異なり，設定用の各コマンドは管理者権限（`sudo`）で起動する点に気をつける必要がある。

### 自動実行可能な外部コマンドの指定

まずは `shell_escape_commands` の値を変更する。
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
$ sudo mktexlsr
mktexlsr: Updating /usr/local/texlive/2018/texmf-config/ls-R... 
mktexlsr: Updating /usr/local/texlive/2018/texmf-dist/ls-R... 
mktexlsr: Updating /usr/local/texlive/2018/texmf-var/ls-R... 
mktexlsr: Updating /usr/local/texlive/texmf-local/ls-R... 
mktexlsr: Done.
```

### 日本語フォントの埋め込み設定

日本語フォントの埋め込み設定はどうなってるんだろうと思ったが

```text
$ kanji-config-updmap status
CURRENT family for ja: ipaex
Standby family : ipa
```

ありゃりゃ， IPA フォントしかないのか。
そりゃそうか。
Linux だもんな。

ちなみに日本語フォントの埋め込みを行わない（非推奨）場合は

```text
$ kanji-config-updmap --user nofont
```

とする。
変更結果は `~/.texlive2018/` ディレクトリに反映される。

`--user` ではなく `--sys` オプションをつけてシステム全体の設定を変更する場合は管理者権限で起動すること[^kanji1]。

[^kanji1]: `kanji-config-updmap-sys` コマンドでも同様。つか `kanji-config-updmap-sys` コマンドは内部で `kanji-config-updmap --sys` を起動しているだけなんだけど。

そういえば [IPAex] フォントの Ver.004.01 がリリースされている。
あの[負の遺産]({{< ref "/remark/2017/12/character-of-the-new-era-name.md" >}} "新元号「文字」という技術的負債")のひとつ「元号の合成文字」である「令和（U+32FF “SQUARE ERA NAME REIWA”）」が追加されている。

- [リリースノート(Release Note) | IPAexフォント/IPAフォント](https://ipafont.ipa.go.jp/node21)
- [情報処理推進機構、新元号“令和”の合字に対応した「IPAexフォント」v004.01を公開 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1182604.html)

とりあえず `/usr/local/texlive/2018/texmf-dist/fonts/truetype/public/ipaex` ディレクトリにあるフォントファイルを差し替えればいいかな。
たぶん 2019 かその次の 2020 あたりで正式にアップデートされるだろう。

その他フォントの設定については以下の拙文を参照のこと。

- [TeX 日本語環境で「源ノ」フォントを使ってみた]({{< ref "/remark/2017/10/using-source-han-fonts-by-japanese-tex.md" >}})
- [数式用フォントで遊ぶ]({{< ref "/remark/2017/10/math-fonts.md" >}})

## 試しに何かタイプセットしてみる

ちうわけでいつものあの文書ですね。
20世紀な内容でゴメン。

- [spiegel-im-spiegel/charset_document: 「文字コードとその実装」 upLaTeX ドキュメント](https://github.com/spiegel-im-spiegel/charset_document)

```
$ latexmk charset.tex
...
Latexmk: All targets (charset.pdf charset.dvi) are up-to-date
```

よし。
ちゃんと動いた。
Lunux 環境でも問題なさそうだな。
ついでに PDF/A への変換もやっておこう。

```text
$ ps2pdf14 -dPDFA -dPDFACompatibilityPolicy=1 -sProcessColorModel=DeviceCMYK charset.pdf charset-pdfa.pdf
```

これもちゃんと動いた。
フォントも正しく埋め込まれているようだ。

{{< fig-img src="./font-property.png" link="./font-property.png" width="570" >}}

もうひとつ， $\mathrm{Lua\LaTeX}$ で以下のタイプセットも試してみる。

- [LuaLaTeX でコードを書いてみる]({{< ref "/remark/2017/10/writing-code-with-lualatex.md" >}})

結果はこんな感じ。

{{< fig-img src="./output.png" link="./output.png" width="686" >}}

よーし，うむうむ，よーし。

## ブックマーク

- [Quick install - TeX Live - TeX Users Group](http://www.tug.org/texlive/quickinstall.html)
- [Ubuntu 16.04 に TeX Live を入れる - 毎日もくもく](http://xartaky.hatenablog.jp/entry/2016-12-27-texlive-on-ubuntu)

- [そうだ， TeX Live 2017 (for Windows) をインストールしよう！]({{< ref "/remark/2017/09/install-tex-live-2017-for-windows.md" >}})

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[TeX Live]: http://www.tug.org/texlive/ "TeX Live - TeX Users Group"
[IPAex]: https://ipafont.ipa.go.jp/ "IPAexフォント/IPAフォント | IPAフォントのダウンロードサイトです"

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
