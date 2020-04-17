+++
title = "TeX Live 2020 で原ノ味フォントを使う"
date =  "2020-04-17T14:34:30+09:00"
description = "TeX Live 2020 で原ノ味フォントが正式に組み込まれ日本語の既定フォントになったようだ。"
image = "/images/attention/kitten.jpg"
tags = [ "tex", "luatex", "font" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

[原ノ味フォント]は Adobe と Google が共同開発している「源ノ」フォントからの派生フォントで， $\mathrm{\TeX}$/$\mathrm{\LaTeX}$ において日本語を扱いやすいようチューニングされているらしい。
ちなみに「源ノ」フォントと同じく SIL Open Font License 1.1 で提供されている。
ありがたや。

[TeX Live] 2020 で[原ノ味フォント]が正式に組み込まれ日本語の既定フォントになったようだ。
ただし [2019 およびそれ以前からアップグレードした]({{< relref "./upgrade-texlive-2020.md" >}} "TeX Live 2020 へのアップグレード")場合は

```text
$ kanji-config-updmap status
CURRENT family for ja: ipaex (variant: <empty>)
Standby family : haranoaji
Standby family : ipa
```

のように以前の設定を引き継いでいるため，[原ノ味フォント]を使うなら手動で設定変更する必要がある[^kanji1]。

[^kanji1]: 全体設定では `kanji-config-updmap-sys` コマンドを，ユーザ毎の設定では `kanji-config-updmap-user` コマンドで使い分ける。当然ながらユーザ毎の設定のほうが優先されるのでご注意を。

```text
$ kanji-config-updmap-sys --jis2004 haranoaji
```

これで設定が

```text
$ kanji-config-updmap status
CURRENT family for ja: haranoaji (variant: -04)
Standby family : haranoaji
Standby family : ipa
Standby family : ipaex
```

となった。

## upLaTeX による組版

まずは $\mathrm{up\LaTeX}$ で組版を行い，フォントの違いを調べてみる。

### 入力テキスト

$\mathrm{up\LaTeX}$ 用に入力テキストを用意する。
こんな感じ。

```latex
\documentclass[uplatex,a4paper]{jsarticle}
\usepackage[deluxe]{otf}

\begin{document}

{\mcfamily\ltseries ☂ は夜更け過ぎに ☃ へとかわるだろう。（明朝体・細字）}\par
{\mcfamily          ☂ は夜更け過ぎに ☃ へとかわるだろう。（明朝体・中字）}\par
{\mcfamily\bfseries ☂ は夜更け過ぎに ☃ へとかわるだろう。（明朝体・太字）}\par

{\gtfamily          ☂ は夜更け過ぎに ☃ へとかわるだろう。（ゴシック体・中字）}\par
{\gtfamily\bfseries ☂ は夜更け過ぎに ☃ へとかわるだろう。（ゴシック体・太字）}\par
{\gtfamily\ebseries ☂ は夜更け過ぎに ☃ へとかわるだろう。（ゴシック体・極太）}

\end{document}
```

これを `uplatex` コマンドで処理する。

### 組版結果（[IPA]ex フォント使用）

まずは [IPA]ex フォントで組版した結果がこちら。

{{< fig-img src="./ipaex-sample.png" link="./ipaex-sample.png" title="upLaTeX 出力結果（IPAex フォント使用）" width="1381" >}}

[IPA]ex フォントはウェイトのバリエーションがないためこんな感じになる。

### 組版結果（[原ノ味フォント]使用）

[原ノ味フォント]を使った組版結果がこちら。

{{< fig-img src="./haranoaji-sample.png" link="./haranoaji-sample.png" title="upLaTeX 出力結果（原ノ味フォント使用）" width="1391" >}}

各ウェイト毎にフォントが対応しているのが分かるだろうか。
あと，絵文字の字形が違うんだけど，いいのか？

## LuaLaTeX による組版

今度は $\mathrm{Lua\LaTeX}$ で試してみる。

### 入力テキスト

```latex {hl_lines=[2]}
\documentclass{ltjsarticle}
\usepackage[haranoaji,deluxe]{luatexja-preset}
\usepackage{luatexja-otf}

\begin{document}

{\mcfamily\ltseries ☂ は夜更け過ぎに ☃ へとかわるだろう。（明朝体・細字）}\par
{\mcfamily          ☂ は夜更け過ぎに ☃ へとかわるだろう。（明朝体・中字）}\par
{\mcfamily\bfseries ☂ は夜更け過ぎに ☃ へとかわるだろう。（明朝体・太字）}\par

{\gtfamily          ☂ は夜更け過ぎに ☃ へとかわるだろう。（ゴシック体・中字）}\par
{\gtfamily\bfseries ☂ は夜更け過ぎに ☃ へとかわるだろう。（ゴシック体・太字）}\par
{\gtfamily\ebseries ☂ は夜更け過ぎに ☃ へとかわるだろう。（ゴシック体・極太）}

\end{document}
```

2行目の `luatexja-preset` パッケージの指定で[原ノ味フォント]を指定してるのがお分かりだろうか[^kanji2]。
これを `lualatex` コマンドで処理する。

[^kanji2]: [TeX Live] 2020 の $\mathrm{Lua\LaTeX}$ では `luatexja-preset` パッケージの既定が[原ノ味フォント]になっているようで，フォントを指定しない場合は[原ノ味フォント]で組版される。

### 組版結果（[原ノ味フォント]使用）

組版結果はこちら。

{{< fig-img src="./lualatex-sample.png" link="./lualatex-sample.png" title="LuaLaTeX 出力結果（原ノ味フォント使用）" width="1388" >}}

$\mathrm{up\LaTeX}$ と同等の出力になっている。
よーし，うむうむ，よーし。

## ブックマーク

- [TeX Live 2020 released | There and back again](https://www.preining.info/blog/2020/04/tex-live-2020-released/)
- [TeX 日本語環境で「源ノ」フォントを使ってみた]({{< ref "/remark/2017/10/using-source-han-fonts-by-japanese-tex.md" >}})
- [LuaTeX-ja に関する覚え書き]({{< ref "/remark/2015/luatex-ja.md" >}})

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[TeX Live]: http://www.tug.org/texlive/ "TeX Live - TeX Users Group"
[原ノ味フォント]: https://github.com/trueroad/HaranoAjiFonts "trueroad/HaranoAjiFonts: 原ノ味フォント"
[IPA]: https://ipafont.ipa.go.jp/ "IPAexフォント/IPAフォント"

## 参考図書

{{% review-paapi "4774187054" %}} <!-- [改訂第7版]LaTeX2ε美文書作成入門 -->
