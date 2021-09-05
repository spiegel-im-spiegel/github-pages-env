+++
title = "数式用フォントで遊ぶ"
date =  "2017-10-05T06:58:54+09:00"
update = "2018-01-04T20:01:07+09:00"
description = "数式用フォントに OpenType フォントを指定する場合は unicode-math パッケージを使う。"
tags        = [ "tex", "luatex", "font", "math" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/profile/"
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

[TeX Live 2017 インストール]({{< ref "/remark/2017/09/install-tex-live-2017-for-windows.md" >}})後の「$\mathrm{\TeX}$ で遊ぼう」第四弾は

- 数式用フォントで遊ぼう

である。

## 既定状態での数式表示

$\mathrm{\TeX}$ および $\mathrm{\LaTeX}$ は元々が数式表現に優れている。
これは $\mathrm{Lua\TeX}$, $\mathrm{Lua\LaTeX}$ でも同じで，たとえばプリアンブルに特に何も書かなくても（『[LaTeX2ε美文書作成入門]』より引用）

```tex
\documentclass[fleqn]{ltjsarticle}

\begin{document}

``?`!`But aren't Kafka's \emph{Schlo{\ss}} and {\AE}sop's {\OE}uvres
often na\"{\i}ve vis-\`a-vis the d{\ae}monic ph{\oe}nix's
official r\^ole in fluffy souffl\'es!?''

\[
  \left( \int_0^\infty \frac{\sin x}{\sqrt x} dx \right)^2 =
  \sum_{k=0}^\infty \frac{(2k)!}{2^{2k}(k!)^2} \frac{1}{2k+1} =
  \prod_{k=1}^\infty \frac{4k^2}{4k^2-1} =
  \frac{\pi}{2}
\]

\end{document}
```

以下のような感じで適切に組んでくれる。

{{< fig-img src="https://photo.baldanders.info/flickr/image/37446584446_m.png" title="equation (1)" link="https://photo.baldanders.info/flickr/37446584446/" >}}

もう少し変わった式やシンボルを使う場合には `amsmath` および `amssymb` パッケージを使って（同じく『[LaTeX2ε美文書作成入門]』より引用）

```tex
\documentclass[fleqn]{ltjsarticle}

\usepackage{amsmath,amssymb} % amsmath packages

\begin{document}

\subsection*{連分数（amsmathパッケージ）}

\begin{equation}
    b_0 + \cfrac{c_1}{b_1 +
          \cfrac{c_2}{b_2 +
          \cfrac{c_3}{b_3 +
          \cfrac{c_4}{b_4 + \cdots}}}}
\end{equation}

\subsection*{黒板文字（amssybパッケージ）}

$\mathbb{ABCDEFGHIJELMN}$

\end{document}
```

このように記述すれば，タイプセットの結果は

{{< fig-img src="https://photo.baldanders.info/flickr/image/36825338243_m.png" title="equation (2)" link="https://photo.baldanders.info/flickr/36825338243/" >}}

こんな感じになる。
ちなみに本文の欧文書体は Latin Modern，数式は Computer Modern で，`amssymb` パッケージで提供される黒板太文字は AMSFonts と呼ばれるものだ[^ams]。

[^ams]: `amsmath` および `amssymb` パッケージは AMS (AmericanMathenatical Society; 米国数学会) が数学論文記述用に開発した $\mathrm{\TeX}$ 用のパッケージ及び文字・記号がベースになっていて，このうちの文字・記号のコレクションが AMSFonts と呼ばれている。 AMSFonts は `amssymb` パッケージを指定することで使用できる。また $\mathrm{\TeX}$ 用の OpenType フォントには AMSFonts を収録したものもある。

## 本文の欧文書体を OpenType フォントに変更する

$\mathrm{Lua\LaTeX}$ で本文の書体を OpenType フォントに変更するのはとても簡単になった。
本文の和文フォントを変更する方法は[前々回に紹介した]({{< ref "/remark/2017/10/using-source-han-fonts-by-japanese-tex.md" >}} "TeX 日本語環境で「源ノ」フォントを使ってみた")が，たとえば欧文書体として $\mathrm{\TeX}$ Gyre をセットする場合は

{{< highlight tex "hl_lines=2 5-7" >}}
\documentclass[fleqn]{ltjsarticle}
\usepackage[no-math,sourcehan]{luatexja-preset} % Japanese fonts

% Import fonts
\setmainfont[Scale=MatchLowercase]{TeXGyreTermes}
\setsansfont[Scale=MatchLowercase]{TeXGyreHeros}
\setmonofont[AutoFakeSlant,BoldItalicFeatures={FakeSlant},Scale=MatchLowercase]{Inconsolatazi4}

\begin{document}

...

\end{document}
{{< /highlight >}}

のように指定する。
欧文フォントの指定には `fontspec` パッケージが必要だが，和文フォント設定用の `luatexja-preset` パッケージに含まれているため，改めて `fontspec` パッケージを明示する必要はない。

ちなみにタイプライタ文字には [Inconsolata] を指定している[^tt1]。

書体を選択する場合は本文の書体と後述する数式用の書体があまりかけ離れないようバランスをとることが重要である。
『[LaTeX2ε美文書作成入門]』には欧文書体についてかなりページを割いて解説されているので，一読することをお薦めする。

[^tt1]: 前回の「[LuaLaTeX でコードを書いてみる]({{< ref "/remark/2017/10/writing-code-with-lualatex.md" >}})」を参照のこと。

## 数式用の書体を変更する

### 既存のパッケージを使う（Concrete ＋ Euler の場合）

たとえば，数学好きに人気の高い Concrete ＋ Euler フォントの組み合わだが，これは既存のパッケージを使ったほうがよさそうである（[後述]({{< relref "#opt" >}})）。
フォントの指定も従来通り。

{{< highlight tex "hl_lines=4" >}}
\documentclass[fleqn]{ltjsarticle}

% Import fonts
\usepackage{ccfonts,eulervm}

\begin{document}

\[
  \left( \int_0^\infty \frac{\sin x}{\sqrt x} dx \right)^2 =
  \sum_{k=0}^\infty \frac{(2k)!}{2^{2k}(k!)^2} \frac{1}{2k+1} =
  \prod_{k=1}^\infty \frac{4k^2}{4k^2-1} =
  \frac{\pi}{2}
\]

\end{document}
{{< /highlight >}}

結果は以下の通り。

{{< fig-img src="https://photo.baldanders.info/flickr/image/37464960082_m.png" title="equation (Concrete + Euler)" link="https://photo.baldanders.info/flickr/37464960082/" >}}

### OpenType フォントを使う（unicode-math パッケージ）

数式用の書体に OpenType フォントを指定する場合は [`unicode-math`] パッケージを使うのが便利だ。
たとえば [$\mathrm{\TeX}$ Gyre Pagella Math](https://ctan.org/pkg/tex-gyre-math-pagella "CTAN: Package tex-gyre-math-pagella") を使う場合は以下のように指定する[^ams1]。

[^ams1]: `amsmath` パッケージを併用する場合は [`unicode-math`] パッケージの前に`amsmath` パッケージを指定しないとタイプセット時にエラーになる。なお `amssymb` パッケージは必要ない。

{{< highlight tex "hl_lines=2 5" >}}
\documentclass[fleqn]{ltjsarticle}
\usepackage{unicode-math} % using unicode/OpenType maths fonts

% Import fonts
\setmathfont{TeXGyrePagella-Math}

\begin{document}

\[
  \left( \int_0^\infty \frac{\sin x}{\sqrt x} dx \right)^2 =
  \sum_{k=0}^\infty \frac{(2k)!}{2^{2k}(k!)^2} \frac{1}{2k+1} =
  \prod_{k=1}^\infty \frac{4k^2}{4k^2-1} =
  \frac{\pi}{2}
\]

\end{document}
{{< /highlight >}}

結果は以下の通り。

{{< fig-img src="https://photo.baldanders.info/flickr/image/23644212118_m.png" title="equation (TeXGyrePagella-Math)" link="https://photo.baldanders.info/flickr/23644212118/" >}}

他の書体も紹介しよう。
[Asana Math](https://ctan.org/pkg/asana-math "CTAN: Package asana-math") を利用する場合は，以下のように指定する。

```tex
\setmathfont{Asana-Math}
```

結果は以下の通り。

{{< fig-img src="https://photo.baldanders.info/flickr/image/23644821888_m.png" title="equation (Asana-Math)" link="https://photo.baldanders.info/flickr/23644821888/" >}}

[Neo Euler] を利用する場合は，[フォントを取得]({{< relref "#elr-inst" >}})して以下のように指定する。

```tex
\setmathfont{Neo Euler}[math-style=upright]
```

`math-style` オプションについては[後述]({{< relref "#opt" >}})する。
結果は以下の通り。

{{< fig-img src="https://photo.baldanders.info/flickr/image/37239500550_m.png" title="equation (Neo Euler)" link="https://photo.baldanders.info/flickr/37239500550/" >}}

[STIX] については [TeX Live] 2017 に収録されているものはデザインがイマイチな気がする。
なので[バージョン2を取得]({{< relref "#stx2-inst" >}})して以下のように指定する。

```tex
\setmathfont{STIX Two Math}
```

結果は以下の通り。

{{< fig-img src="https://photo.baldanders.info/flickr/image/36788335984_m.png" title="equation (STIX2 Math)" link="https://photo.baldanders.info/flickr/36788335984/" >}}

[STIX] の fork とも言える（？） [XITS](https://github.com/khaledhosny/xits "khaledhosny/xits: XITS - OpenType implementation of STIX fonts with math support") を利用する場合は，以下のように指定する。

```tex
\setmathfont{XITS-Math}
```

結果は以下の通り。

{{< fig-img src="https://photo.baldanders.info/flickr/image/37450915146_m.png" title="equation (XITS Math)" link="https://photo.baldanders.info/flickr/37450915146/" >}}

## unicode-math パッケージのオプション {#opt}

`unicode-math` パッケージのオプションはいくつかあるが，今回はよく使うと思われるものについて紹介する。

`math-style`, `bold-style`, `sans-style` 各オプションは数式の文字や記号のスタイルを指定する。
具体的には ラテン小文字，ラテン大文字，ギリシア小文字，ギリシア大文字 それぞれに対して「立体（upright）」または「斜体（slant）」の組み合わせを指定する。

数式全体（`math-style`）のスタイル設定は以下の通り。

| `math-style` | L小文字 | L大文字 | G小文字 | G大文字 |
|:-------------|:-------:|:-------:|:-------:|:-------:|
| `ISO`        | 斜体    | 斜体    | 斜体    | 斜体    |
| `TeX`        | 斜体    | 斜体    | 斜体    | 立体    |
| `french`     | 斜体    | 立体    | 立体    | 立体    |
| `upright`    | 立体    | 立体    | 立体    | 立体    |

`\symbf` 指定時（`bold-style`）のスタイル設定は以下の通り。

| `bold-style` | L小文字 | L大文字 | G小文字 | G大文字 |
|:-------------|:-------:|:-------:|:-------:|:-------:|
| `ISO`        | 斜体    | 斜体    | 斜体    | 斜体    |
| `TeX`        | 立体    | 立体    | 斜体    | 立体    |
| `upright`    | 立体    | 立体    | 立体    | 立体    |

`\symsf` 指定時（`sans-style`）のスタイル設定は以下の通り。

| `sans-style` | L小文字 | L大文字 | G小文字 | G大文字 |
|:-------------|:-------:|:-------:|:-------:|:-------:|
| `italic`     | 斜体    | 斜体    | 斜体    | 斜体    |
| `upright`    | 立体    | 立体    | 立体    | 立体    |

オプションの指定方法は以下の通り。

```tex
\usepackage[math-style=ISO,bold-style=ISO]{unicode-math}
```

または

```tex
\usepackage{unicode-math}
\unimathsetup{%
    math-style=ISO,%
    bold-style=ISO%
}
```

フォント指定時にオプションを指定することもできる。

```tex
\setmathfont{Neo Euler}[math-style=upright]
```

### Neo Euler 書体について

[Neo Euler] のギリシア文字には斜体がないようで， `math-style` の値を `ISO` または `TeX` に設定するとギリシア文字が表示されない。

{{< fig-img src="https://photo.baldanders.info/flickr/image/23648392388_m.png" title="equation (Neo Euler 2)" link="https://photo.baldanders.info/flickr/23648392388/" >}}

他にも AMSFonts にあるシンボル（先程の黒板太文字とか）が存在しなかったりするようだ。
Euler フォントに関しては，使うのであれば， OpenType ではなく既存のパッケージを使うほうが簡単だろう。

## 【付録1】 euler.otf ファイルのインストール {#elr-inst}

`euler.otf` ファイルは [TeX Live] 2017 には収録されていないため手動でインストールする必要がある。

といっても簡単で，まず `{$TEXMFLOCAL}/fonts/opentype/public/` フォルダを作成し，そこに [`euler-otf`] リポジトリを `git clone` すればOK。
ちなみに `$TEXMFLOCAL` の場所は `kpsewhich` コマンドで調べられる。
いつもの場所だね。

```text
$ kpsewhich -var-value=TEXMFLOCAL
C:/texlive/texmf-local
```

コピーできたら `mktexlsr` コマンドを実行して `ls-R` を更新し，さらに `luaotfload-tool` も実行すること。

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

$ luaotfload-tool --update
```

ちゃんとインストールできたか確認も忘れずに。

```text
$ luaotfload-tool --fuzzy --find="neo euler"
luaotfload | resolve : Font "neo euler" found!
luaotfload | resolve : Resolved file name "c:/texlive/texmf-local/fonts/opentype/public/euler-otf/euler.otf"
```

## 【付録2】 STIX2 のインストール {#stx2-inst}

[STIX] バージョン2（zip ファイル）をダウンロード[^sf] し， `Fonts/OTF/` フォルダにある以下のファイルを `{$TEXMFLOCAL}/fonts/opentype/public/STIX2/` フォルダにコピーする。

[^sf]: ダウンロード・ページは広告だらけで且つ[色々と悪名高い SourceForge](http://gigazine.net/news/20150722-sourceforge/ "オープンソースソフトウェアの老舗サイト「SourceForge」はいかにして堕ちていったのか - GIGAZINE") なのでご注意を。

- `STIX2Math.otf`
- `STIX2Text-Bold.otf`
- `STIX2Text-BoldItalic.otf`
- `STIX2Text-Italic.otf`
- `STIX2Text-Regular.otf`

コピーできたら `mktexlsr` コマンドを実行して `ls-R` を更新し，さらに `luaotfload-tool` も実行すること。

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

$ luaotfload-tool --update
```

ちゃんとインストールできたか確認も忘れずに。

```text
$ luaotfload-tool --fuzzy --find="stix two math"
luaotfload | resolve : Font "stix two math" found!
luaotfload | resolve : Resolved file name "c:/texlive/texmf-local/fonts/opentype/public/stix2/stix2math.otf"
```

## ブックマーク

- {{< pdf-file title="Experimental Unicode mathematical typesetting: The unicode-math package" link="http://mirrors.ibiblio.org/CTAN/macros/latex/contrib/unicode-math/unicode-math.pdf" >}} 
- [LuaLaTeX で unicode-math を使う場合の書体変更コマンド | org-技術](https://org-technology.com/posts/lualatex-unicode-math.html)

[TeX Live]: http://www.tug.org/texlive/ "TeX Live - TeX Users Group"
[LaTeX2ε美文書作成入門]: https://www.amazon.co.jp/exec/obidos/ASIN/4774187054/baldandersinf-22/ "Amazon | [改訂第7版]LaTeX2ε美文書作成入門 | 奥村 晴彦, 黒木 裕介 通販"
[Inconsolata]: http://levien.com/type/myfonts/inconsolata.html
[Neo Euler]: https://github.com/khaledhosny/euler-otf "khaledhosny/euler-otf: An abandoned OpenType/Unicode math port of AMS Euler font"
[STIX]: http://www.stixfonts.org/ "STIX Fonts Project Website"
[`unicode-math`]: https://github.com/wspr/unicode-math "wspr/unicode-math: XeLaTeX/LuaLaTeX package for using unicode/OpenType maths fonts"
[`euler-otf`]: https://github.com/khaledhosny/euler-otf "khaledhosny/euler-otf: An abandoned OpenType/Unicode math port of AMS Euler font"

## 参考図書 {#books}

{{% review-paapi "4297117126" %}} <!-- LaTeX2ε美文書作成入門 -->
