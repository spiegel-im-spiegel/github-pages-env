+++
title = "Firge フォントを使って LuaLaTeX でコードを書く"
date =  "2021-09-30T22:29:34+09:00"
description = "こういうのが欲しかったんだよ。"
image = "/images/attention/kitten.jpg"
tags = [ "tex", "luatex", "font" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

テキスト・エディタやターミナル・エミュレータ用の日本語が使えるフォントとして有名な[白源（HackGen）](https://github.com/yuru7/HackGen)だが，個人的にどうしても思い入れがあって今まで使わなかった[^ez1]。
で，たまたま[作者の方の Qiita 記事](https://qiita.com/tawara_/items/374f3ca0a386fab8b305 "Ricty を神フォントだと崇める僕が、フリーライセンスのプログラミングフォント「白源」を作った話 - Qiita")を目にしたのだが

[^ez1]: いや，大した思い入れでもないんだけど，新人（大昔）の頃は紙のコーディング・シートで書かされたこともあって，手書きのゼロは $\emptyset$ (empty set) なイメージなんだよねぇ。

{{< fig-quote type="markdown" title="Ricty を神フォントだと崇める僕が、フリーライセンスのプログラミングフォント「白源」を作った話 - Qiita" link="https://qiita.com/tawara_/items/374f3ca0a386fab8b305" >}}
HackGen の姉妹フォント "**Firge (ファージ)**" を試験的にリリースしました！<br>
https://github.com/yuru7/Firge/releases

英数字部分に Mozilla 製 Fira Mono、日本語文字部分に HackGen と同じ源柔ゴシック（濁点等の修正も含む）、を適用しています。<br>
HackGen の批判的な意見として多かった以下の点を満たすことを意識しました。

- `0` はドットゼロではなくスラッシュゼロに。
- `|` は破断線ではなく直線に。
{{< /fig-quote >}}

とアナウンスされていた。
ありがとう！ これが欲しかったんだよ。

というわけで早速 Windows Terminal (+ [NYAGOS](https://github.com/nyaosorg/nyagos "nyaosorg/nyagos: NYAGOS - The hybrid Commandline Shell between UNIX & DOS")) と Ubuntu 機のターミナル・エミュレータに導入した。
あっ，あとブラウザの monospace の既定フォントね[^editor1]。

[^editor1]: 私はテキスト・エディタは明朝体じゃないと目がしんどいので，英数字は Inconsolata だけど日本語部分は Noto Serif CJK JP にしている。どなたか明朝体の monospace を作ってください。

ちなみに Ununtu 機にフォントを入れる場合は

| ディレクトリ | 用途 |
| --- | --- |
| `/usr/share/fonts` | システムフォント |
| `/usr/local/share/fonts` | 追加フォントをマシンで共有する場合 |
| `~/.fonts` | 個人で導入する場合 |

のいずれかにフォントファイルをコピーして（ディレクトリを掘ってもOK）

```text
$ fc-cache -fv
```

でキャッシュを更新すれば OK。
確認は

```text
$ fc-list | grep Firge
/usr/local/share/fonts/FirgeNerd/Firge35Nerd-Regular.ttf: Firge35Nerd:style=Regular
/usr/local/share/fonts/FirgeNerd/FirgeNerd-Bold.ttf: FirgeNerd:style=Bold
/usr/local/share/fonts/FirgeNerd/FirgeNerd-Regular.ttf: FirgeNerd:style=Regular
/usr/local/share/fonts/FirgeNerd/Firge35NerdConsole-Regular.ttf: Firge35Nerd Console:style=Regular
/usr/local/share/fonts/FirgeNerd/Firge35NerdConsole-Bold.ttf: Firge35Nerd Console:style=Bold
/usr/local/share/fonts/FirgeNerd/FirgeNerdConsole-Bold.ttf: FirgeNerd Console:style=Bold
/usr/local/share/fonts/FirgeNerd/FirgeNerdConsole-Regular.ttf: FirgeNerd Console:style=Regular
/usr/local/share/fonts/FirgeNerd/Firge35Nerd-Bold.ttf: Firge35Nerd:style=Bold
```

とすればよい。

で，この [Firge] フォントを使って $\mathrm{Lua\LaTeX}$ でコードを表示したいと思うよね。ね。

というわけで試してみよう。
なお [TeX Live] のインストールは先日書いた「[改めて TeX Live を Ubuntu に（APT を使わずに）導入する]({{< relref "./install-texlive-in-ubuntu-again.md" >}})」を参照のこと。

$\mathrm{Lua\LaTeX}$ でフォントを認識させるために `luaotfload-tool --update` でアップデートしておく。

```text
$ luaotfload-tool --update
```

念の為の確認。

```text
$ luaotfload-tool --fuzzy --find="FirgeNerd"
luaotfload | resolve : Font "FirgeNerd" found!
luaotfload | resolve : Resolved file name "/usr/local/share/fonts/FirgeNerd/FirgeNerd-Regular.ttf"
```

よしよし。

システムに [Firge] フォントをインストールしないのであれば `$TEXMFLOCAL/fonts/` ディレクトリに適当に放り込んでおけばよい。
今回は $\mathrm{Lua\LaTeX}$ のみが対象なので CMAP ファイルとか面倒くさそうなのは華麗にスルーする（笑）

なお `$TEXMFLOCAL` ディレクトリは `kpsewhich` コマンドで調べられる。

```text
$ kpsewhich -var-value=TEXMFLOCAL
/usr/local/texlive/texmf-local
```

ファイルを入れたら `ls-R` を更新するのを忘れないように。

```text
$ mktexlsr
mktexlsr: Updating /usr/local/texlive/2021/texmf-config/ls-R... 
mktexlsr: Updating /usr/local/texlive/2021/texmf-dist/ls-R... 
mktexlsr: Updating /usr/local/texlive/2021/texmf-var/ls-R... 
mktexlsr: Updating /usr/local/texlive/texmf-local/ls-R... 
mktexlsr: Done.
```

念の為 `kpsewhich` コマンドで確認しておく。

```text
$ kpsewhich FirgeNerd-Regular.ttf
/usr/local/texlive/texmf-local/fonts/truetype/FirgeNerd/FirgeNerd-Regular.ttf
```

確認したら `luaotfload-tool --update` も忘れないでね。

先日 [TeX Live をインストール]({{< relref "./install-texlive-in-ubuntu-again.md" >}} "改めて TeX Live を Ubuntu に（APT を使わずに）導入する")したときに使ったサンプルコードは以下の通り。

```latex
\documentclass{jlreq}
\usepackage[jis2004,deluxe]{luatexja-preset} % Japanese fonts
\setmonofont[AutoFakeSlant,BoldItalicFeatures={FakeSlant},Scale=MatchLowercase]{Inconsolatazi4} % use Inconsolata

\usepackage{graphicx,xcolor}
% \usepackage{listings}
\usepackage{listings-golang} % import this package after listings
\lstset{
    frame=single,
    basicstyle=\small\ttfamily,
    tabsize=4,
    commentstyle=\color{darkgray},
    keywordstyle=\color{brown}\bfseries,
    stringstyle=\color{blue},
    showstringspaces=false
}

\begin{document}

\section{Go 言語による Hello World}

\begin{lstlisting}[language=Golang]
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        fmt.Println("Hello, world") //Hello, 世界
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

これの組版結果はこちら。

{{< fig-img src="../install-texlive-in-ubuntu-again/sample.png" title="sample.pdf" link="../install-texlive-in-ubuntu-again/sample.pdf" width="1252" >}}

ここで `setmonofont` 指定を以下のように変更する。

```latex
\setmonojfont[AutoFakeSlant,BoldItalicFeatures={FakeSlant},Scale=MatchLowercase]{FirgeNerd} % use FirgeNerd
\setmonofont[AutoFakeSlant,BoldItalicFeatures={FakeSlant},Scale=MatchLowercase]{FirgeNerd} % use FirgeNerd
```

この場合の組版結果は以下の通り。

{{< fig-img src="./sample2.png" title="sample2.pdf" link="./sample2.pdf" width="1253" >}}

小文字のエルとかで比べると分かりやすいだろうか。ただ日本語は [Firge] フォントじゃないっぽい。

{{< fig-img src="./property.png" title="プロパティ" link="./property.png" >}}

これの HaranoAjiGothic-Regular は多分コードの「世界」の部分だよねぇ。理由は不明。

ついでなので [HackGen] フォントでも試してみようか。

```latex
\setmonojfont[AutoFakeSlant,BoldItalicFeatures={FakeSlant},Scale=MatchLowercase]{FirgeNerd} % use FirgeNerd
\setmonofont[AutoFakeSlant,BoldItalicFeatures={FakeSlant},Scale=MatchLowercase]{FirgeNerd} % use FirgeNerd
```

組版結果は以下の通り。

{{< fig-img src="./sample3.png" title="sample3.pdf" link="./sample3.pdf" width="1248" >}}

これなら違いが分かりやすいかな。

ふむむー。
個人の好みの問題なんだろうけど，これなら [Inconsolata] でも別にいいかなぁ。
やはりターミナル・エミュレータ用と割り切って使うほうがいいか。

## ブックマーク

- [LuaLaTeX でコードを書いてみる]({{< ref "/remark/2017/10/writing-code-with-lualatex.md" >}})

[TeX Live]: http://www.tug.org/texlive/ "TeX Live - TeX Users Group"
[HackGen]: https://github.com/yuru7/HackGen "yuru7/HackGen: Hack と源柔ゴシックを合成したプログラミングフォント 白源 (はくげん／HackGen)"
[Firge]: https://github.com/yuru7/Firge "yuru7/Firge: Fira Mono と源真ゴシックを合成したプログラミングフォント Firge (ファージ)"
[Inconsolata]: http://levien.com/type/myfonts/inconsolata.html

## 参考図書

{{% review-paapi "4297117126" %}} <!-- LaTeX2ε美文書作成入門 -->
