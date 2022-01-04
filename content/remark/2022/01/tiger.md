+++
title = "2022年は寅年らしいので"
date =  "2022-01-04T15:16:30+09:00"
description = "おめでとうございます〜"
image = "/remark/2022/01/tiger/tiger-lualatex.png"
tags = [ "tex", "luatex" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

2022年は寅年らしい。
トラといえば昔の $\mathrm{\LaTeX}$ 本で EPS ファイルを貼り込むときによく使われる虎の絵（`tiger.eps`）って今もあるんだろうか。

と思って[『LaTeX2ε美文書作成入門』の第8版](https://www.amazon.co.jp/dp/4297117126?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "[改訂第8版]LaTeX2ε美文書作成入門 | 奥村晴彦, 黒木裕介 |本 | 通販 | Amazon")を見てみたら今も御健在らしい。
実際に TeX Live のインストール先ディレクトリを検索してみたらあちこちにあった（笑）

というわけで $\mathrm{Lua\LaTeX}$ で虎の絵を表示してみる。
こんな感じでどうかな。

```latex
\documentclass[report]{jlreq}
\usepackage[haranoaji,jis2004,deluxe,match,jfm_yoko=jlreq,jfm_tate=jlreqv]{luatexja-preset} % Japanese fonts
\usepackage{graphicx}

\begin{document}

\begin{figure}
  \centering
  \includegraphics[width=\columnwidth,clip]{./tiger.eps}
  \caption{2022年は寅年らしい}
  \label{fig:tiger}
\end{figure}

\end{document}
```

`jlreq` は最近流行りの汎用ドキュメントクラスで W3C の「[日本語組版処理の要件](https://www.w3.org/TR/jlreq/)」にほぼ準拠しているそうな。
しかも $\mathrm{p\LaTeX}$, $\mathrm{up\LaTeX}$, $\mathrm{Lua\LaTeX}$ で透過的に使えるのがありがたい（まぁ，フォントやドライバの指定の仕方が違うので全く同じにはならないけど）。

これを処理した結果は以下の通り。

{{< fig-img src="./tiger-lualatex.png" title="tiger.pdf" link="./tiger.pdf" width="949" >}}

おめでとうございます〜

## ブックマーク

- [改めて TeX Live を Ubuntu に（APT を使わずに）導入する]({{< ref "/remark/2019/04/install-texlive-in-ubuntu.md" >}})

## 参考図書

{{% review-paapi "4297117126" %}} <!-- LaTeX2ε美文書作成入門 -->
