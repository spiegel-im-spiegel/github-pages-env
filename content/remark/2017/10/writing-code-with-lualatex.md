+++
title = "LuaLaTeX でコードを書いてみる"
date =  "2017-10-04T18:19:51+09:00"
update = "2018-01-04T20:01:07+09:00"
description = "とりあえず Go 言語による以下のコードを LuaLaTeX で書くことを考える。"
tags        = [ "tex", "luatex", "font", "golang", "programming", "language" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
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

（なんか，すっかりシリーズ化しちゃったなぁ。別 section にすべきだったか？）

さて，[TeX Live 2017 インストール]({{< ref "/remark/2017/09/install-tex-live-2017-for-windows.md" >}})後の「$\mathrm{\TeX}$ で遊ぼう」第三弾は

- $\mathrm{Lua\LaTeX}$ でコードを書いてみる

である。
といっても $\mathrm{Lua\LaTeX}$ 特有の部分は少ないと思うので，後半のフォント指定以外は $\mathrm{p\LaTeX}$ とかでもいけるだろう。

とりあえず [Go 言語]による以下のコードを $\mathrm{Lua\LaTeX}$ で書くことを考える。

```go
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        fmt.Println("Hello, world")
    }
}
```

安直に `verbatim` 環境を使う手もあるが，世の中には [`listings`] パッケージなる便利なものがあるらしい。
これを使ってまずは

{{< highlight tex "hl_lines=4-9 15 25" >}}
\documentclass{ltjsarticle}
\usepackage[no-math,sourcehan]{luatexja-preset} % Japanese fonts

\usepackage{listings}
\lstset{
    frame=single,
    basicstyle=\small\ttfamily,
    tabsize=4
}

\begin{document}

\section{Go 言語による Hello World}

\begin{lstlisting}
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        fmt.Println("Hello, world")
    }
}
\end{lstlisting}

\end{document}
{{< /highlight >}}

と書いてみる。
結果はこんな感じ。

{{< fig-img src="https://farm5.staticflickr.com/4505/37440154426_9d1a97aa3d.jpg" title="code with LuaLaTeX (1)"  link="https://www.flickr.com/photos/spiegel/37440154426/" >}}

悪かないけど，キーワードの強調表示くらい欲しいよね。
[`listings`] パッケージでは `language=C++` みたいな感じに言語を指定できるらしい。
じゃあ，ってんで早速 `language=Go` とかやってみたんだけど「知らんがな」って怒られた。

```text
! Package Listings Error: Couldn't load requested language.
```

んー。
ってことは誰か [Go 言語]用の補助パッケージを作ってくれていれば... というわけで探したらありましたよ。

- [julienc91/listings-golang: Golang support for the listings package in LaTeX](https://github.com/julienc91/listings-golang )

神！！！

というわけで，ここにある `listings-golang.sty` を取ってきて作業フォルダに置けばいいんだけど，折角なら定常的に使いたいので，以下のフォルダを掘って，そこにリポジトリを git clone してしまった。

- `{$TEXMFLOCAL}/tex/latex/`

ちなみに `$TEXMFLOCAL` の場所は `kpsewhich` コマンドで調べられる。
いつもの場所だね。

```text
$ kpsewhich -var-value=TEXMFLOCAL
C:/texlive/texmf-local
```

`mktexlsr` コマンドで `ls-R` を更新するのも忘れずに。

では，先程の文書を [`listings-golang`] パッケージを使って以下のように書き換える[^glst1]。

[^glst1]: [`listings-golang`] パッケージは `\RequirePackage` コマンドで [`listings`] パッケージを内部で呼び出しているため，記述上は [`listings`] パッケージを置き換えることができる。のだが，今回は README にしたがっている。

{{< highlight tex "hl_lines=4 6 11-12 19" >}}
\documentclass{ltjsarticle}
\usepackage[no-math,sourcehan]{luatexja-preset} % Japanese fonts

\usepackage{graphicx,xcolor}
\usepackage{listings}
\usepackage{listings-golang} % import this package after listings
\lstset{
    frame=single,
    basicstyle=\small\ttfamily,
    tabsize=4,
    keywordstyle=\color{red},
    stringstyle=\color{blue}
}

\begin{document}

\section{Go 言語による Hello World}

\begin{lstlisting}[language=Golang]
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        fmt.Println("Hello, world")
    }
}
\end{lstlisting}

\end{document}
{{< /highlight >}}

結果はこんな感じ。

{{< fig-img src="https://farm5.staticflickr.com/4443/37488315861_e3ecc8b840.jpg" title="code with LuaLaTeX (2)"  link="https://www.flickr.com/photos/spiegel/37488315861/" >}}

おっとお。
原色は派手すぎたか。
まぁ，ここでは分かりやすさ優先ってことで。

実は [`listings-golang`] パッケージの中身はそれほど難しくなくて（でも自分でやれって言われたらやっぱり面倒くさいんだけど）

```tex
%% Golang definition for listings
%% http://github.io/julienc91/lstlistings-golang
%%
\RequirePackage{listings}

\lstdefinelanguage{Golang}%
  {morekeywords=[1]{package,import,func,type,struct,return,defer,panic,%
     recover,select,var,const,iota,},%
   morekeywords=[2]{string,uint,uint8,uint16,uint32,uint64,int,int8,int16,%
     int32,int64,bool,float32,float64,complex64,complex128,byte,rune,uintptr,%
     error,interface},%
   morekeywords=[3]{map,slice,make,new,nil,len,cap,copy,close,true,false,%
     delete,append,real,imag,complex,chan,},%
   morekeywords=[4]{for,break,continue,range,goto,switch,case,fallthrough,if,%
     else,default,},%
   morekeywords=[5]{Println,Printf,Error,},%
   sensitive=true,%
   morecomment=[l]{//},%
   morecomment=[s]{/*}{*/},%
   morestring=[b]',%
   morestring=[b]",%
   morestring=[s]{`}{`},%
}
```

という感じでキーワードと文字列とコメントの3つの定義でできている。
このようにして [`listings`] パッケージが対応していない言語にも対応できる。

最後に，やっぱタイプライタ文字なら [Inconsolata] だよね，ってことで，タイプライタ文字を以下のように変更する[^z4]。

[^z4]: 厳密に言うと，今回使うのはオリジナルの [Inconsolata] を改良した zi4 版である。

{{< highlight tex "hl_lines=3" >}}
\documentclass{ltjsarticle}
\usepackage[no-math,sourcehan]{luatexja-preset} % Japanese fonts
\setmonofont[AutoFakeSlant,BoldItalicFeatures={FakeSlant},Scale=MatchLowercase]{Inconsolatazi4}

\usepackage{graphicx,xcolor}
\usepackage{listings}
\usepackage{listings-golang} % import this package after listings
\lstset{
    frame=single,
    basicstyle=\small\ttfamily,
    tabsize=4,
    keywordstyle=\color{red},
    stringstyle=\color{blue}
}

\begin{document}

\section{Go 言語による Hello World}

\begin{lstlisting}[language=Golang]
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        fmt.Println("Hello, world")
    }
}
\end{lstlisting}

\end{document}
{{< /highlight >}}

昔の $\mathrm{Lua\TeX}$-ja では，欧文フォントと和文フォントが干渉するので，指定の順番とかうるさかったが，エラく簡単になったものである。

結果はこんな感じ。

{{< fig-img src="https://farm5.staticflickr.com/4451/37441473856_10ccde1a0c.jpg" title="code with LuaLaTeX (4)"  link="https://www.flickr.com/photos/spiegel/37441473856/" >}}

数字のゼロに斜線が入っているのがお分かりだろうか。
以前は[ゼロやバッククォート等の表示で苦労してた](http://d.hatena.ne.jp/zrbabbler/20130421/1366560678 "LaTeX の inconsolata がアレなので何とかする - マクロツイーター")気がするが，最近のものなら問題ないようだ。

ちなみに，フォント指定にある `AutoFakeSlant` や `BoldItalicFeatures={FakeSlant}` といった記述は斜体の定義である。
[Inconsolata] フォントには斜体やイタリック体はないので $\mathrm{Lua\LaTeX}$ 内部で擬似的な斜体を作るよう指示しているのだ。

## ブックマーク

- [LaTeXでソースコードを書く - しがないプログラマ（仮）のブログ](http://turgure.hatenablog.com/entry/2016/08/19/183501)
- [listings.sty: LaTeX パッケージ](http://www.biwako.shiga-u.ac.jp/sensei/kumazawa/tex/listings.html) : `\lstset` について簡単な解説

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`listings`]: https://ctan.org/tex-archive/macros/latex/contrib/listings "CTAN: /tex-archive/macros/latex/contrib/listings"
[`listings-golang`]: https://github.com/julienc91/listings-golang "julienc91/listings-golang: Golang support for the listings package in LaTeX"
[Inconsolata]: http://levien.com/type/myfonts/inconsolata.html
