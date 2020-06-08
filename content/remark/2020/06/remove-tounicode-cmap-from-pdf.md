+++
title = "LuaTeX で生成した PDF から ToUnicode CMap を除去する"
date =  "2020-06-08T17:31:44+09:00"
description = "原ノ味フォントの作者が pdf-rm-tuc というツールを公開されている。ありがたや。"
image = "/images/attention/kitten.jpg"
tags = [ "luatex", "pdf", "unicode", "font", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

[以前に紹介]({{< ref "/remark/2020/04/haranoaji-fonts-with-texlive-2020.md" >}} "TeX Live 2020 で原ノ味フォントを使う")した[原ノ味フォント]だが， $\mathrm{Lua\LaTeX}$ で組版 & PDF 出力する際に Adobe-Identity-0 フォントとして埋め込まれてしまうそうで，フォントの性能を上手く活かしきれてないらしい。
しかも

{{< fig-quote type="markdown" title="PDF から ToUnicode CMap を削除するツール" link="https://github.com/trueroad/pdf-rm-tuc/blob/master/README.ja.md" >}}
{{% quote %}}LuaTeX はフォントの cmap テーブルから逆変換で ToUnicode CMap を生成しているようで、縦書き用グリフを使った場合や異字体を使った場合など、テキスト抽出すると意図しない結果になることがあります{{% /quote %}}
{{< /fig-quote >}}

そこで[原ノ味フォント]の作者が [`pdf-rm-tuc`] というツールを公開されている。

- [trueroad/pdf-rm-tuc: Remove ToUnicode CMap from PDF](https://github.com/trueroad/pdf-rm-tuc)

このツールを使うと

{{< fig-quote type="markdown" title="PDF から ToUnicode CMap を削除するツール" link="https://github.com/trueroad/pdf-rm-tuc/blob/master/README.ja.md" >}}
{{% quote %}}PDF から原ノ味フォントの ToUnicode CMap を削除し、原ノ味フォントの ROS を Adobe-Identity-0 から 元の Adobe-Japan1-7 に変更します。 これによってテキスト抽出で意図しない結果になることを防げると考えています{{% /quote %}}
{{< /fig-quote >}}

更に言うと PDF ファイルのサイズがちょびっとだけ小さくなる。
ありがたや `m(_ _)m`

## [pdf-rm-tuc][`pdf-rm-tuc`] のビルドとインストール

[`pdf-rm-tuc`] はソースコードのみの公開なので GCC を使って自前でビルドする必要がある[^msys2]。
ここでは [Ubuntu] プラットフォームを前提に紹介する。

[^msys2]: Windows であれば [MSYS2 等の環境]({{< ref "/remark/2016/03/gcc-msys2-1.md" >}} "MSYS2 による gcc 開発環境の構築 ― MSYS2 のインストールから初期化処理まで")が必要になるかもしれない。あるいは [TeX Live] ひっくるめて WSL2 に引っ越すか（笑）

### 事前準備

たまに GCC が入ってないことがあるので，入ってなければ APT でインストールしておく。

```text
$ sudo apt install build-essential
```

更に [`pdf-rm-tuc`] の GitHub リポジトリを clone して直接ビルドする場合は Autoconf および Automake が必要になる。
もし入ってなければ，これもインストールしておこう。

```text
$ sudo apt install autoconf automake
```

更に更にビルドの際に lib[qpdf](https://github.com/qpdf/qpdf "qpdf/qpdf: Primary QPDF source code and documentation") が必要らしいので，これもインストールしておく。

```text
$ sudo apt install libqpdf-dev
```

これで準備完了。

### ビルドとインストール

では，ビルドからインストールまで一気にやってしまおう。

```text
$ git clone https://github.com/trueroad/pdf-rm-tuc.git
$ cd pdf-rm-tuc
$ ./autogen.sh
$ mkdir build
$ cd build
$ ../configure
$ make
$ make check # TeX Live 2020 が導入済みであること
$ sudo make install
$ pdf-rm-tuc -V
Remove ToUnicode CMap from PDF 1.0.0
Copyright (C) 2019 Masamichi Hosoda. All rights reserved.
License: BSD-2-Clause

https://github.com/trueroad/pdf-rm-tuc
```

インストールが成功すれば [`pdf-rm-tuc`] コマンドが以下に置かれる。

```text
$ which pdf-rm-tuc
/usr/local/bin/pdf-rm-tuc
```

## [pdf-rm-tuc][`pdf-rm-tuc`] を試してみる

入力テキストは以下の通り。
[TeX Live] 2020 が導入されていることが前提ね。

```latex
\documentclass{ltjsarticle}
\usepackage[deluxe]{luatexja-preset}
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

これを組版すると以下のような結果になる。

{{< fig-img src="/remark/2020/04/haranoaji-fonts-with-texlive-2020/lualatex-sample.png" link="/remark/2020/04/haranoaji-fonts-with-texlive-2020/lualatex-sample.png" width="1388" >}}

このときのフォント情報はこんな風になっている[^poppler1]。

[^poppler1]: `pdffonts` は [Poppler] のコマンドのひとつ。 Windows 版 [TeX Live] には既定で同梱されているらしい。他のプラットフォームについては「[Poppler - TeX Wiki](https://texwiki.texjp.org/?Poppler)」を参考に導入すればいいだろう。

```text
$ pdffonts lualatex-sample.pdf 
name                                 type              encoding         emb sub uni object ID
------------------------------------ ----------------- ---------------- --- --- --- ---------
OLUZKZ+HaranoAjiMincho-Light         CID Type 0C       Identity-H       yes yes yes      4  0
PTNBJH+HaranoAjiMincho-Regular       CID Type 0C       Identity-H       yes yes yes      5  0
BMAPJQ+HaranoAjiMincho-Bold          CID Type 0C       Identity-H       yes yes yes      6  0
HOSFGF+HaranoAjiGothic-Regular       CID Type 0C       Identity-H       yes yes yes      7  0
ULINFN+HaranoAjiGothic-Bold          CID Type 0C       Identity-H       yes yes yes      8  0
PJPUYK+HaranoAjiGothic-Heavy         CID Type 0C       Identity-H       yes yes yes      9  0
JFRMQG+LMRoman10-Regular             CID Type 0C       Identity-H       yes yes yes     10  0
```

`uni` の項目が ToUnicode CMap の有無を示すもので，[原ノ味フォント]の全てに ToUnicode CMap があるのが分かる。

では [`pdf-rm-tuc`] を実行してみよう。

```text
$ pdf-rm-tuc lualatex-sample.pdf lualatex-sample-notuc.pdf
```

実行結果は以下の通り。

```text
$ pdffonts lualatex-sample-notuc.pdf 
name                                 type              encoding         emb sub uni object ID
------------------------------------ ----------------- ---------------- --- --- --- ---------
JFRMQG+LMRoman10-Regular             CID Type 0C       Identity-H       yes yes yes     10  0
PTNBJH+HaranoAjiMincho-Regular       CID Type 0C       Identity-H       yes yes no       5  0
OLUZKZ+HaranoAjiMincho-Light         CID Type 0C       Identity-H       yes yes no       4  0
BMAPJQ+HaranoAjiMincho-Bold          CID Type 0C       Identity-H       yes yes no       6  0
HOSFGF+HaranoAjiGothic-Regular       CID Type 0C       Identity-H       yes yes no       7  0
ULINFN+HaranoAjiGothic-Bold          CID Type 0C       Identity-H       yes yes no       8  0
PJPUYK+HaranoAjiGothic-Heavy         CID Type 0C       Identity-H       yes yes no       9  0
```

[原ノ味フォント]の `uni` 項目が全て `no` になっていることが確認できた。
よーし，うむうむ，よーし。

## 注意事項

- $\mathrm{up\LaTeX}$ 等で組版したものを `dvipdfmx` で PDF に出力する場合は Adobe-Japan1 フォントであれば ToUnicode CMap の生成・埋め込みはしないので [`pdf-rm-tuc`] は不要
- PDF ビュア側で Adobe-Japan1-UCS2 等の ToUnicode CMap を持っていない場合は PDF 側の ToUnicode CMap を削除すると日本語のテキスト抽出が上手く行かないらしい。 [Ubuntu] に標準で入ってる [Evince](https://wiki.gnome.org/Apps/Evince) は画面から普通にコピペできたけど，大丈夫ってこと？
- PDF/A を構成する場合，適合レベルによっては ToUnicode CMap を削除すると拙い場合があるのだが [Adobe-Japan1 フォントについては例外]({{< relref "./pdfa-with-luatex.md" >}} "LuaLaTeX で PDF/A を構成する")らしい？ [veraPDF] 等の Validator で確認したほうがいいかも

## ブックマーク

- [trueroad/tr-TeXConf2019: TeXConf 2019 一般講演「原ノ味フォントと ToUnicode CMap」関連資料](https://github.com/trueroad/tr-TeXConf2019)

[原ノ味フォント]: https://github.com/trueroad/HaranoAjiFonts "trueroad/HaranoAjiFonts: 原ノ味フォント"
[`pdf-rm-tuc`]: https://github.com/trueroad/pdf-rm-tuc "trueroad/pdf-rm-tuc: Remove ToUnicode CMap from PDF"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[TeX Live]: http://www.tug.org/texlive/ "TeX Live - TeX Users Group"
[Poppler]: https://poppler.freedesktop.org/
[veraPDF]: https://verapdf.org/ "veraPDF | Industry Supported PDF/A Validation"

## 参考図書

{{% review-paapi "4774187054" %}} <!-- [改訂第7版]LaTeX2ε美文書作成入門 -->
