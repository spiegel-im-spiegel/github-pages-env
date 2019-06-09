+++
title = "TeX 日本語環境で「源ノ」フォントを使ってみた"
date =  "2017-10-03T20:10:56+09:00"
description = "第二弾は TeX で「源ノ」フォントを使ってみる である。"
tags        = [ "tex", "font", "luatex", "japanese" ]

[scripts]
  mathjax = true
  mermaidjs = false
+++

さて，[TeX Live 2017 インストール]({{< ref "/remark/2017/09/install-tex-live-2017-for-windows.md" >}})後の「$\mathrm{\TeX}$ で遊ぼう」第二弾は

- $\mathrm{\TeX}$ で「源ノ」フォントを使ってみる

である。

「源ノ」フォントとは「[源ノ角ゴシック]（Source Han Sans）」および「[源ノ明朝]（Source Han Serif）」で構成される CJK 統一デザインの書体を指す。
「源ノ」フォントは Adobe と Google が共同で開発していて，かつオープンソースで公開されている。
この成果の一部は Google により NOTO プロジェクトに取り込まれている。

- [Google Noto Fonts](https://www.google.com/get/noto/)
    - [Google Developers Japan: オープンソースの美しい Noto フォントファミリーに日本語、中国語、韓国語が加わりました。](https://developers-jp.googleblog.com/2014/07/noto.html)
    - [Google Developers Japan: Noto Serif CJK が登場！](https://developers-jp.googleblog.com/2017/04/noto-serif-cjk-is-here.html)

いわゆる CJK 漢字圏の文字は漢字だけでも非常に数が多く ☂ や ☃ みたいな非漢字も合わせると想像するだけでも目眩がする。
開発に携われた方々には足を向けて寝られない（何処に足を向けたらいいのだろう`w`）。
そして勿論，世の $\mathrm{\TeX}$nician たちがこれに指を咥えて見ているはずはないのだ。
なので私もその成果に便乗することにする。

皆さん，ありがとう！

## 「源ノ」フォントのインストール {#install}

早速インストールをはじめるが，[源ノ角ゴシック]および[源ノ明朝]は GitHub で公開されている。

- [adobe-fonts/source-han-sans: Source Han Sans | 思源黑体 | 思源黑體 | 源ノ角ゴシック | 본고딕](https://github.com/adobe-fonts/source-han-sans)
- [adobe-fonts/source-han-serif: Source Han Serif | 思源宋体 | 思源宋體 | 源ノ明朝 | 본명조](https://github.com/adobe-fonts/source-han-serif)

これらのリポジトリから最新リリース（zip または tar.gz ファイル）をダウンロードする。
1.6 から 2 GB ほどの巨大パッケージなのでダウンロードの際はご注意を。
取得したパッケージの内，今回使うのは `OTF/Japanese` フォルダにある以下のファイルである。

- [源ノ角ゴシック]
    1. `SourceHanSans-ExtraLight.otf`
    2. `SourceHanSans-Light.otf`
    3. `SourceHanSans-Normal.otf`
    4. `SourceHanSans-Regular.otf`
    5. `SourceHanSans-Medium.otf`
    6. `SourceHanSans-Bold.otf`
    7. `SourceHanSans-Heavy.otf`
- [源ノ明朝]
    1. `SourceHanSerif-ExtraLight.otf`
    2. `SourceHanSerif-Light.otf`
    3. `SourceHanSerif-Regular.otf`
    4. `SourceHanSerif-Medium.otf`
    5. `SourceHanSerif-SemiBold.otf`
    6. `SourceHanSerif-Bold.otf`
    7. `SourceHanSerif-Heavy.otf`

それぞれのファイルの違いはウェイトの違いで，軽いものから昇順に並べている。
さらに各パッケージ内の `Resources` フォルダにある以下の CMap ファイル[^cmap1] も使用する。

[^cmap1]: CMap とは Unicode 等の文字集合とフォントの CID (Character IDentifier) を対応付けるマッピング情報である。 OpenType フォントのフルセットは Unicode を含む複数の文字集合の字形を包括している。たとえば JIS X 0213 では「葛（U+845B）」の字について2000年版と現行の2004年版で字形が異なるという混乱があるが（厳密には例示字体で示される字の形が異なる。異体字のことではない）， OpenType フォント（のフルセット）はそういった差異も包括し CID で識別することができる。従って CMap は文字集合やそのバージョンごとに存在することになる。

- [源ノ角ゴシック]
    - `UniSourceHanSansJP-UTF16-H`
- [源ノ明朝]
    - `UniSourceHanSerifJP-UTF16-H`

これらのファイルを，以下のフォルダを作成して放り込む[^fnt1]。

[^fnt1]: `adobe`, `sourcehansansjp`, `sourcehanjp` のフォルダ名には特に意味はなく，自分で管理しやすい名前にすればいいらしい。

- [源ノ角ゴシック]： `{$TEXMFLOCAL}/fonts/opentype/adobe/sourcehansansjp/`
- [源ノ明朝]： `{$TEXMFLOCAL}/fonts/opentype/adobe/sourcehanserifjp/`
- CMap ファイル： `{$TEXMFLOCAL}/fonts/cmap/sourcehanjp/`

ちなみに `$TEXMFLOCAL` の場所は `kpsewhich` コマンドで調べられる。
いつもの場所だね。

```text
$ kpsewhich -var-value=TEXMFLOCAL
C:/texlive/texmf-local
```

コピーできたら `mktexlsr` コマンドを実行して `ls-R` を更新しておくことを忘れずに。

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

これで「源ノ」フォントのインストールは完了。
一応 `kpsewhich` コマンドでインストールできたかどうか確認しておくとよい。

```text
$ kpsewhich SourceHanSerif-Regular.otf
c:/texlive/texmf-local/fonts/opentype/adobe/sourcehanserifjp/SourceHanSerif-Regular.otf
```

## upLaTeX で「源ノ」フォントを使ってみる {#uplatex}

では文書ファイルを用意しよう。
こんな感じでどうだろうか。

{{< highlight tex "hl_lines=1-2" >}}
\documentclass[uplatex,a4paper]{jsarticle}
\usepackage[deluxe]{otf}

\begin{document}

{\mcfamily\ltseries ☂ は夜更け過ぎに ☃ へとかわるだろう。（明朝体・細字）}\par
{\mcfamily          ☂ は夜更け過ぎに ☃ へとかわるだろう。（明朝体・中字）}\par
{\mcfamily\bfseries ☂ は夜更け過ぎに ☃ へとかわるだろう。（明朝体・太字）}\par

{\gtfamily          ☂ は夜更け過ぎに ☃ へとかわるだろう。（ゴシック体・中字）}\par
{\gtfamily\bfseries ☂ は夜更け過ぎに ☃ へとかわるだろう。（ゴシック体・太字）}\par
{\gtfamily\ebseries ☂ は夜更け過ぎに ☃ へとかわるだろう。（ゴシック体・極太）}

森鷗外と内田百閒が髙島屋に行った。

\UTF{845B}飾区の\UTF{20BB7}野家  % 葛飾区の𠮷野家

\end{document}
{{< /highlight >}}

タイプセット処理は $\mathrm{up\LaTeX}$ を選択した。
$\mathrm{up\LaTeX}$ で処理したあと `dvipdfmx` コマンドで PDF に変換しないといけないんだけど，「源ノ」フォント用の map ファイルは [TeX Live] 2017 には存在しないようなので，自前で用意しないといけない。
他の書体の map ファイルを参考に自作してみた（正しいかどうかわからないが）。

```text
% CID
otf-cjml-h    Identity-H    SourceHanSerif-Light.otf
otf-cjml-v    Identity-V    SourceHanSerif-Light.otf
otf-cjmr-h    Identity-H    SourceHanSerif-Regular.otf
otf-cjmr-v    Identity-V    SourceHanSerif-Regular.otf
otf-cjmb-h    Identity-H    SourceHanSerif-Bold.otf
otf-cjmb-v    Identity-V    SourceHanSerif-Bold.otf
otf-cjgr-h    Identity-H    SourceHanSans-Regular.otf
otf-cjgr-v    Identity-V    SourceHanSans-Regular.otf
otf-cjgb-h    Identity-H    SourceHanSans-Bold.otf
otf-cjgb-v    Identity-V    SourceHanSans-Bold.otf
otf-cjge-h    Identity-H    SourceHanSans-Heavy.otf
otf-cjge-v    Identity-V    SourceHanSans-Heavy.otf
otf-cjmgr-h   Identity-H    SourceHanSans-Heavy.otf
otf-cjmgr-v   Identity-V    SourceHanSans-Heavy.otf

% Unicode 90JIS
otf-ujml-h    UniSourceHanSerifJP-UTF16-H   SourceHanSerif-Light.otf
otf-ujml-v    UniSourceHanSerifJP-UTF16-V   SourceHanSerif-Light.otf
otf-ujmr-h    UniSourceHanSerifJP-UTF16-H   SourceHanSerif-Regular.otf
otf-ujmr-v    UniSourceHanSerifJP-UTF16-V   SourceHanSerif-Regular.otf
otf-ujmb-h    UniSourceHanSerifJP-UTF16-H   SourceHanSerif-Bold.otf
otf-ujmb-v    UniSourceHanSerifJP-UTF16-V   SourceHanSerif-Bold.otf
otf-ujgr-h    UniSourceHanSansJP-UTF16-H    SourceHanSans-Regular.otf
otf-ujgr-v    UniSourceHanSansJP-UTF16-V    SourceHanSans-Regular.otf
otf-ujgb-h    UniSourceHanSansJP-UTF16-H    SourceHanSans-Bold.otf
otf-ujgb-v    UniSourceHanSansJP-UTF16-V    SourceHanSans-Bold.otf
otf-ujge-h    UniSourceHanSansJP-UTF16-H    SourceHanSans-Heavy.otf
otf-ujge-v    UniSourceHanSansJP-UTF16-V    SourceHanSans-Heavy.otf
otf-ujmgr-h   UniSourceHanSansJP-UTF16-H    SourceHanSans-Heavy.otf
otf-ujmgr-v   UniSourceHanSansJP-UTF16-V    SourceHanSans-Heavy.otf

% Unicode JIS04
otf-ujmln-h    UniJIS2004-UTF16-H    SourceHanSerif-Light.otf
otf-ujmln-v    UniJIS2004-UTF16-V    SourceHanSerif-Light.otf
otf-ujmrn-h    UniJIS2004-UTF16-H    SourceHanSerif-Regular.otf
otf-ujmrn-v    UniJIS2004-UTF16-V    SourceHanSerif-Regular.otf
otf-ujmbn-h    UniJIS2004-UTF16-H    SourceHanSerif-Bold.otf
otf-ujmbn-v    UniJIS2004-UTF16-V    SourceHanSerif-Bold.otf
otf-ujgrn-h    UniJIS2004-UTF16-H    SourceHanSans-Regular.otf
otf-ujgrn-v    UniJIS2004-UTF16-V    SourceHanSans-Regular.otf
otf-ujgbn-h    UniJIS2004-UTF16-H    SourceHanSans-Bold.otf
otf-ujgbn-v    UniJIS2004-UTF16-V    SourceHanSans-Bold.otf
otf-ujgen-h    UniJIS2004-UTF16-H    SourceHanSans-Heavy.otf
otf-ujgen-v    UniJIS2004-UTF16-V    SourceHanSans-Heavy.otf
otf-ujmgrn-h   UniJIS2004-UTF16-H    SourceHanSans-Heavy.otf
otf-ujmgrn-v   UniJIS2004-UTF16-V    SourceHanSans-Heavy.otf

% TEXT, 90JIS
uphminl-h   UniSourceHanSerifJP-UTF16-H SourceHanSerif-Light.otf
uphminl-v   UniSourceHanSerifJP-UTF16-V SourceHanSerif-Light.otf
uphminr-h   UniSourceHanSerifJP-UTF16-H SourceHanSerif-Regular.otf
uphminr-v   UniSourceHanSerifJP-UTF16-V SourceHanSerif-Regular.otf
uphminb-h   UniSourceHanSerifJP-UTF16-H SourceHanSerif-Bold.otf
uphminb-v   UniSourceHanSerifJP-UTF16-V SourceHanSerif-Bold.otf
uphgothr-h  UniSourceHanSansJP-UTF16-H  SourceHanSans-Regular.otf
uphgothr-v  UniSourceHanSansJP-UTF16-V  SourceHanSans-Regular.otf
uphgothb-h  UniSourceHanSansJP-UTF16-H  SourceHanSans-Bold.otf
uphgothb-v  UniSourceHanSansJP-UTF16-V  SourceHanSans-Bold.otf
uphgotheb-h UniSourceHanSansJP-UTF16-H  SourceHanSans-Heavy.otf
uphgotheb-v UniSourceHanSansJP-UTF16-V  SourceHanSans-Heavy.otf
uphmgothr-h UniSourceHanSansJP-UTF16-H  SourceHanSans-Heavy.otf
uphmgothr-v UniSourceHanSansJP-UTF16-V  SourceHanSans-Heavy.otf

% TEXT, JIS04
uphminln-h  UniJIS2004-UTF16-H  SourceHanSerif-Light.otf
uphminln-v  UniJIS2004-UTF16-V  SourceHanSerif-Light.otf
uphminrn-h  UniJIS2004-UTF16-H  SourceHanSerif-Regular.otf
uphminrn-v  UniJIS2004-UTF16-V  SourceHanSerif-Regular.otf
uphminbn-h  UniJIS2004-UTF16-H  SourceHanSerif-Bold.otf
uphminbn-v  UniJIS2004-UTF16-V  SourceHanSerif-Bold.otf
uphgothrn-h UniJIS2004-UTF16-H  SourceHanSans-Regular.otf
uphgothrn-v UniJIS2004-UTF16-V  SourceHanSans-Regular.otf
uphgothbn-h UniJIS2004-UTF16-H  SourceHanSans-Bold.otf
uphgothbn-v UniJIS2004-UTF16-V  SourceHanSans-Bold.otf
uphgothebn-h    UniJIS2004-UTF16-H  SourceHanSans-Heavy.otf
uphgothebn-v    UniJIS2004-UTF16-V  SourceHanSans-Heavy.otf
uphmgothrn-h    UniJIS2004-UTF16-H  SourceHanSans-Heavy.otf
uphmgothrn-v    UniJIS2004-UTF16-V  SourceHanSans-Heavy.otf

urml    UniSourceHanSerifJP-UTF16-H SourceHanSerif-Regular.otf
urmlv   UniSourceHanSerifJP-UTF16-V SourceHanSerif-Regular.otf
ugbm    UniSourceHanSansJP-UTF16-H  SourceHanSans-Medium.otf
ugbmv   UniSourceHanSansJP-UTF16-V  SourceHanSans-Medium.otf
uprml-h UniSourceHanSerifJP-UTF16-H SourceHanSerif-Regular.otf
uprml-v UniSourceHanSerifJP-UTF16-V SourceHanSerif-Regular.otf
upgbm-h UniSourceHanSansJP-UTF16-H  SourceHanSans-Medium.otf
upgbm-v UniSourceHanSansJP-UTF16-V  SourceHanSans-Medium.otf
uprml-hq    UniJIS-UCS2-H   SourceHanSerif-Regular.otf
upgbm-hq    UniJIS-UCS2-H   SourceHanSans-Medium.otf

% TEXT, 90JIS
hminl-h     H    SourceHanSerif-Light.otf
hminl-v     V    SourceHanSerif-Light.otf
hminr-h     H    SourceHanSerif-Regular.otf
hminr-v     V    SourceHanSerif-Regular.otf
hminb-h     H    SourceHanSerif-Bold.otf
hminb-v     V    SourceHanSerif-Bold.otf
hgothr-h    H    SourceHanSans-Regular.otf
hgothr-v    V    SourceHanSans-Regular.otf
hgothb-h    H    SourceHanSans-Bold.otf
hgothb-v    V    SourceHanSans-Bold.otf
hgotheb-h   H    SourceHanSans-Heavy.otf
hgotheb-v   V    SourceHanSans-Heavy.otf
hmgothr-h   H    SourceHanSans-Heavy.otf
hmgothr-v   V    SourceHanSans-Heavy.otf

% TEXT, JIS04
hminln-h    H    SourceHanSerif-Light.otf
hminln-v    V    SourceHanSerif-Light.otf
hminrn-h    H    SourceHanSerif-Regular.otf
hminrn-v    V    SourceHanSerif-Regular.otf
hminbn-h    H    SourceHanSerif-Bold.otf
hminbn-v    V    SourceHanSerif-Bold.otf
hgothrn-h   H    SourceHanSans-Regular.otf
hgothrn-v   V    SourceHanSans-Regular.otf
hgothbn-h   H    SourceHanSans-Bold.otf
hgothbn-v   V    SourceHanSans-Bold.otf
hgothebn-h  H    SourceHanSans-Heavy.otf
hgothebn-v  V    SourceHanSans-Heavy.otf
hmgothrn-h  H    SourceHanSans-Heavy.otf
hmgothrn-v  V    SourceHanSans-Heavy.otf
```

分かりにくいとは思うが，フォントの選択は以下のように対応付けている。

| 字体とウエイト   | フォント名 |
|:-----------------|:-----------|
| 明朝体・細字     | `SourceHanSerif-Light.otf` |
| 明朝体・中字     | `SourceHanSerif-Regular.otf` |
| 明朝体・太字     | `SourceHanSerif-Bold.otf` |
| ゴシック体・中字 | `SourceHanSans-Regular.otf` |
| ゴシック体・太字 | `SourceHanSans-Bold.otf` |
| ゴシック体・極太 | `SourceHanSans-Heavy.otf` |

また `UniSourceHanSerifJP-UTF16-V` や `UniSourceHanSansJP-UTF16-V` という CMap ファイルは存在しないが，まぁ縦書きは使わないから，私。
縦書き用の CMap ファイルはネットを探したらあるかもしれない。
`UniJIS2004-UTF16-H` や `UniJIS-UCS2-H` は他のフォントの設定をそのまま流用してるけど，正しいのかどうかよく分からない。

`dvipdfmx` コマンドで map ファイルを直接指定する必要があるため `.latexmkrc` ファイルを以下のように設定した。

{{< highlight perl "hl_lines=6" >}}
#!/usr/bin/env perl
$latex        = 'uplatex -synctex=1';
$latex_silent = 'uplatex -synctex=1 -interaction=batchmode';
$bibtex       = 'upbibtex';
$biber        = 'biber --bblencoding=utf8 -u -U --output_safechars';
$dvipdf       = 'dvipdfmx -z9 -V 4 -f sourcehanjp.map %O -o %D %S';
$makeindex    = 'mendex %O -o %D %S';
$max_repeat   = 5;
$pdf_mode     = 3; # generates pdf via dvipdfmx

$pvc_view_file_via_temporary = 0;
{{< /highlight >}}

ようやく準備完了。
やっとビルドできるよ。

```text
$ latexmk otf-sample.tex
```

結果はこんな感じ。

{{< fig-img src="https://photo.baldanders.info/flickr/image/36799783933_m.png" title="Output PDF by upLaTeX with Source Han Fonts (1)" link="https://photo.baldanders.info/flickr/36799783933/" >}}

ちなみに埋め込まれたフォントの状態は以下の通り。

{{< fig-img src="https://photo.baldanders.info/flickr/image/23617581598_m.png" title="Output PDF by upLaTeX with Source Han Fonts (2)" link="https://photo.baldanders.info/flickr/23617581598/" >}}

## LuaLaTeX で「源ノ」フォントを使ってみる {#lualatex}

さて，では同じことを $\mathrm{Lua\LaTeX}$-ja でやってみようと思う。

さきほど[フォントファイルをインストールした]({{< relref "#install" >}})が， $\mathrm{Lua\LaTeX}$ で使うにはそのままではダメで，以下のコマンドを実行しておく必要がある。

```text
$ luaotfload-tool --update
```

上手くいったかどうかは以下のコマンドで確認できる。

```text
$ luaotfload-tool --fuzzy --find="sourcehanserif"
luaotfload | resolve : Font "sourcehanserif" found!
luaotfload | resolve : Resolved file name "c:/texlive/texmf-local/fonts/opentype/adobe/sourcehanserifjp/sourcehanserif-regular.otf"
```

では，文書ファイルを作ってみよう。

{{< highlight tex "hl_lines=1-3" >}}
\documentclass{ltjsarticle}
\usepackage[sourcehan,deluxe]{luatexja-preset}
\usepackage{luatexja-otf}

\begin{document}

{\mcfamily\ltseries ☂ は夜更け過ぎに ☃ へとかわるだろう。（明朝体・細字）}\par
{\mcfamily          ☂ は夜更け過ぎに ☃ へとかわるだろう。（明朝体・中字）}\par
{\mcfamily\bfseries ☂ は夜更け過ぎに ☃ へとかわるだろう。（明朝体・太字）}\par

{\gtfamily          ☂ は夜更け過ぎに ☃ へとかわるだろう。（ゴシック体・中字）}\par
{\gtfamily\bfseries ☂ は夜更け過ぎに ☃ へとかわるだろう。（ゴシック体・太字）}\par
{\gtfamily\ebseries ☂ は夜更け過ぎに ☃ へとかわるだろう。（ゴシック体・極太）}

森鷗外と内田百閒が髙島屋に行った。

\UTF{845B}飾区の\UTF{20BB7}野家  % 葛飾区の𠮷野家

\end{document}
{{< /highlight >}}

`luatexja-preset` パッケージでは「源ノ」フォントを指定できる。

```tex
\usepackage[sourcehan]{luatexja-preset}
```

また `\UTF` コマンドが使えるように `luatexja-otf` パッケージを読み込んでいる。
いやぁ，すんごい楽ちん。

タイプセットも前準備なしで簡単にできる（だだし，めがっさ遅い）。

```text
$ lualatex otf-sample.tex
```

結果は以下の通り。

{{< fig-img src="https://photo.baldanders.info/flickr/image/36800606433_m.png" title="Output PDF by LuaLaTeX with Source Han Fonts (1)" link="https://photo.baldanders.info/flickr/36800606433/" >}}

んー。
ウェイトがさっきと違う？

{{< fig-img src="https://photo.baldanders.info/flickr/image/36800606553_m.png" title="Output PDF by LuaLaTeX with Source Han Fonts (2)" link="https://photo.baldanders.info/flickr/36800606553/" >}}

いや，同じか。
まぁいいや。
「極太」なんてまず使わんし。

他に気になることと言えば， $\mathrm{up\LaTeX}$ ＋ dvipdfmx のときもそうだったんだけど，「葛（U+845B）」の字形が JIS X 0213:2004 だよね。
どちらも `jis2004` オプションは付けてないんだが。
やはり `jis2004` オプションは相当特殊な気がする。
これからは JIS X 0213:2004 前提で考えて，むしろ `jis2000` オプションを作ったほうがいいんじゃないだろうか。

ま，それはともかく

**時代は $\mathrm{Lua\TeX}$ で「源ノ」フォント** ってことですね！

ちなみに，[前回]({{< ref "/remark/2017/10/lualatex-ja-on-cloud.md" >}})紹介した [Cloud LaTeX] で同じことをしようとしたら「`sourcehan` なんぞ知らん！」と怒られた（笑）

```text
/usr/local/texlive/2016/texmf-dist/tex/latex/fontspec/fontspec-luatex.sty
line 392
! LaTeX Error: Unknown option `sourcehan' for package `fontspec-luatex'.
```

まぁしょうがないね。
[IPAex] が使えるのなら支障はないし。

## ブックマーク

- [The Typekit Blog | Source Han Sansの紹介：オープンソースのPan-CJK書体](https://blog.typekit.com/alternate/source-han-sans-jp/)
- [源ノ明朝](https://source.typekit.com/source-han-serif/jp/)
- [GoogleとMonotypeが全言語対応フォントのNotoを公開 | TechCrunch Japan](http://jp.techcrunch.com/2016/10/08/20161006google-and-monotype-unveil-the-noto-projects-unified-font-for-all-languages/)
- [日中韓に対応したグーグルの新フォント「Noto Serif CJK」は、なぜ生まれたか｜WIRED.jp](https://wired.jp/2017/04/08/noto-serif-cjk/)
- [フォント - TeX Wiki](https://texwiki.texjp.org/?%E3%83%95%E3%82%A9%E3%83%B3%E3%83%88)
- [源ノ明朝/源ノ角ゴシックをLuaLaTeXで使用する (Windows)](https://jb102.blogspot.jp/2017/04/11-luatex.html)
- [upLaTeX文書で源ノ明朝／Noto Serif CJKを簡単に使う方法（最新のdvipdfmxとpxchfonを使用） - Qiita](https://qiita.com/zr_tex8r/items/9dfeafecca2d091abd02)
- [小形克宏の「文字の海、ビットの舟」――文字コードが私たちに問いかけるもの](http://internet.watch.impress.co.jp/www/column/ogata/index.htm)
- [updmap and kanji - TeX Live - TeX Users Group](https://www.tug.org/texlive/updmap-kanji.html)

[源ノ角ゴシック]: https://github.com/adobe-fonts/source-han-sans "adobe-fonts/source-han-sans: Source Han Sans | 思源黑体 | 思源黑體 | 源ノ角ゴシック | 본고딕"
[源ノ明朝]: https://github.com/adobe-fonts/source-han-serif "adobe-fonts/source-han-serif: Source Han Serif | 思源宋体 | 思源宋體 | 源ノ明朝 | 본명조"
[TeX Live]: http://www.tug.org/texlive/ "TeX Live - TeX Users Group"
[Cloud LaTeX]: https://cloudlatex.io/ "Cloud LaTeX | Build your own LaTeX environment, in seconds"
[IPAex]: http://ipafont.ipa.go.jp/node26#ja "IPAexフォント Ver.003.01(IPAexFont Ver.003.01) | IPAexフォント/IPAフォント"

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
