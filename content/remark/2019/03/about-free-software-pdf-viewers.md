+++
title = "FOSS な PDF Viewer for Windows (Evince も導入できた) "
date = "2019-03-09T12:39:15+09:00"
description = "この中で Windows 環境かつ比較的活況な製品を選ぶならやっぱ XpdfReader か MuPDF ってところか。"
image = "/images/attention/kitten.jpg"
tags = [ "pdf", "tools", "msys2" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

ちうわけで，[前回]の続き。

PDF Viewer は FOSS な製品からプロプライエタリな製品まで色々あるし，全部を比較するのは無理なので，[前回]に引き続き [PDFreaders.org] に掲載されている製品からピックアップしてみよう。

各製品のスペックは以下の通り（2019年3月時点）。

| 製品名       | 動作環境                     | フォーマット   | ライセンス | 最近のリリース状況   |
| ------------ | ---------------------------- | -------------- | ---------- | -------------------- |
| [Evince]     | Linux                        | PDF, ...       | GPL        | 3.31 (2019-02-19)    |
| [Okular]     | Linux, macOS, Windows[^kde1] | PDF, EPUB, ... | GPL        | 1.6 (2018-12-13)     |
| [SumatraPDF] | Windows                      | PDF, EPUB, ... | GPL        | 3.1.2 (2016-09-23)   |
| [XpdfReader] | Linux, Windows               | PDF            | GPL        | 4.01 (2019-02-18)    |
| [PDF.js]     | Firefox 19+                  | PDF            | Apache-2.0 | 2.1.266 (2019-02-17) |
| [qpdfview]   | Linux, NetBSD, FreeBSD       | PDF, ...       | GPL        | 0.4.18 (2018-11-17)  |
| [GNU GV]     | Linux macOS                  | PDF, PS        | GPL        | 3.7.4 (2013-03-17)   |
| [ePDFView]   | Linux                        | PDF            | GPL        | 0.1.8 (2011-05-28)   |
| [MuPDF]      | Linux, windows, Android      | PDF, EPUB, ... | AGPL       | 1.14.0 (2018-10-04)  |
| [zathura]    | Linux                        | PDF, ...       | zlib       | 0.4.3 (2018-12-22)   |

[^kde1]: [Okular] を Windows で動作させるには [KDE on Windows](http://windows.kde.org/) 環境が必要。同様に macOS では [KDE on Mac OS X](https://community.kde.org/Mac) 環境が必要となる。 

んー。
この中で Windows 環境かつ比較的活況な製品を選ぶならやっぱ [XpdfReader] か [MuPDF] ってところか。
EPUB 文書も見たいなら [MuPDF] 一択だな。
ちなみに [MuPDF] は [SyncTeX] に対応しているらしい。

そうそう。
[MuPDF] の Android 版で EPUB 文書を開いてみたが，見事に日本語文字がトーフ（□）になった。

{{< fig-img title="EPUB版『続・情報共有の未来』を Android 版 MuPDF で開いてみた（笑）" src="./Screenshot_20190309-114206.png" link="Screenshot_20190309-114206.png" width="1440" >}}

実は EPUB の仕様はよく知らないのだけど[^epub1]，おそらく対応しているフォントがないんだろう（Windows 版はちゃんと表示できる）。

[^epub1]: EPUB 興味ないよなぁ。日本では業界主導で（縦書きがどうとか）なんか騒いでいたみたいだけど，所詮は「紙の延長」だし，それなら別に PDF でいいぢゃんと思ってしまう。そのためにも PDF Viewer が（無料じゃなくて）自由なライセンスで提供されることが重要なんだけどね。「紙」という制約が不要ならそれこそ AsciiDoc とか markdown とか Org mode とかプリミティブに HTML & CSS とか色々あるわけじゃん。

## [Evince] for [MSYS2] を導入する

[前回]の記事で [Evince] の Windows 用バイナリがなくなったと書いたが [MSYS2] 経由なら導入できそうである。
ただし [MSYS2] は万人におすすめの方法ではないので，既に [MSYS2] を導入している方のみという感じだろうか。
[MSYS2] の導入手順は以下の拙文を参考にどうぞ。

- [MSYS2 による gcc 開発環境の構築 ― MSYS2 のインストールから初期化処理まで]({{< ref "/remark/2016/03/gcc-msys2-1.md" >}})
- [MSYS2 による gcc 開発環境の構築 ― gcc パッケージ群の導入]({{< ref "/remark/2016/03/gcc-msys2-2.md" >}})

ね。
万人向けじゃないでしょ。

{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}， `pacman` コマンドを使って [Evince] を探してみると

```text
$ pacman -Ss evince
mingw32/mingw-w64-i686-evince 3.28.2-3
    Document (PostScript, PDF) viewer (mingw-w64)
mingw64/mingw-w64-x86_64-evince 3.28.2-3
    Document (PostScript, PDF) viewer (mingw-w64)
```

おおっ！ あったあった。
ふむ。
`mingw32` か `mingw64` のいずれかのリポジトリから取るのか。
ちょいとバージョンが古いが，まぁいいか。

ちなみに `mingw32` からインストールした場合には `/mingw32` ディレクトリ以下に `mingw64` からインストールした場合には `/mingw64` ディレクトリ以下にインストールされる。
Shell として bash を起動する際にはモードにご注意を。

```text
C:> SET MSYSTEM=MINGW64
C:> C:\msys64\usr\bin\bash.exe --login -i
```

じゃあ導入していこう。

```text
$ pacman -S mingw-w64-x86_64-evince
依存関係を解決しています...
衝突するパッケージがないか確認しています...
警告: 循環依存が検出されました:
警告: mingw-w64-x86_64-harfbuzz は依存パッケージ mingw-w64-x86_64-freetype の前にインストールされます

パッケージ (74) mingw-w64-x86_64-adwaita-icon-theme-3.30.1-1  mingw-w64-x86_64-atk-2.30.0-1
                mingw-w64-x86_64-brotli-1.0.7-1  mingw-w64-x86_64-c-ares-1.15.0-1
                mingw-w64-x86_64-cairo-1.16.0-1  mingw-w64-x86_64-curl-7.64.0-2
                mingw-w64-x86_64-dbus-1.12.12-1  mingw-w64-x86_64-djvulibre-3.5.27-3
                mingw-w64-x86_64-fontconfig-2.13.1-1  mingw-w64-x86_64-freeglut-3.0.0-4
                mingw-w64-x86_64-freetype-2.9.1-1  mingw-w64-x86_64-fribidi-1.0.5-1
                mingw-w64-x86_64-gdk-pixbuf2-2.38.1-1  mingw-w64-x86_64-ghostscript-9.26-1
                mingw-w64-x86_64-glib2-2.58.3-1  mingw-w64-x86_64-graphene-1.8.4-1
                mingw-w64-x86_64-graphite2-1.3.13-1 mingw-w64-x86_64-gsettings-desktop-schemas-3.28.1-1
                mingw-w64-x86_64-gsl-2.5-1  mingw-w64-x86_64-gst-plugins-base-1.14.4-1
                mingw-w64-x86_64-gstreamer-1.14.4-1  mingw-w64-x86_64-gtk3-3.24.5-1
                mingw-w64-x86_64-harfbuzz-2.3.1-1  mingw-w64-x86_64-hicolor-icon-theme-0.17-1
                mingw-w64-x86_64-icu-62.1-1  mingw-w64-x86_64-jansson-2.12-1
                mingw-w64-x86_64-jasper-2.0.14-1  mingw-w64-x86_64-jemalloc-5.1.0-3
                mingw-w64-x86_64-json-glib-1.4.4-1  mingw-w64-x86_64-lcms2-2.9-1
                mingw-w64-x86_64-libarchive-3.3.3-2  mingw-w64-x86_64-libcroco-0.6.12-1
                mingw-w64-x86_64-libdatrie-0.2.12-1  mingw-w64-x86_64-libepoxy-1.5.3-1
                mingw-w64-x86_64-libgcrypt-1.8.4-1  mingw-w64-x86_64-libgpg-error-1.35-1
                mingw-w64-x86_64-libgxps-0.3.1-1  mingw-w64-x86_64-libidn-1.35-1
                mingw-w64-x86_64-libidn2-2.1.1a-1  mingw-w64-x86_64-libjpeg-turbo-2.0.2-1
                mingw-w64-x86_64-libmetalink-0.1.3-3  mingw-w64-x86_64-libogg-1.3.3-1
                mingw-w64-x86_64-libpaper-1.1.24-2  mingw-w64-x86_64-libpng-1.6.36-1
                mingw-w64-x86_64-libpsl-0.20.2-3  mingw-w64-x86_64-librsvg-2.40.20-1
                mingw-w64-x86_64-libspectre-0.2.8-2  mingw-w64-x86_64-libssh2-1.8.0-3
                mingw-w64-x86_64-libthai-0.1.28-2  mingw-w64-x86_64-libtheora-1.1.1-4
                mingw-w64-x86_64-libtiff-4.0.10-1  mingw-w64-x86_64-libunistring-0.9.10-1
                mingw-w64-x86_64-libvorbis-1.3.6-1  mingw-w64-x86_64-libvorbisidec-svn-r19643-1
                mingw-w64-x86_64-libxml2-2.9.9-1  mingw-w64-x86_64-libxslt-1.1.33-1
                mingw-w64-x86_64-lz4-1.8.3-1  mingw-w64-x86_64-lzo2-2.10-1
                mingw-w64-x86_64-nettle-3.4.1-1  mingw-w64-x86_64-nghttp2-1.36.0-1
                mingw-w64-x86_64-nspr-4.20-1  mingw-w64-x86_64-nss-3.42.1-1
                mingw-w64-x86_64-openjpeg2-2.3.0-2  mingw-w64-x86_64-opus-1.3-1
                mingw-w64-x86_64-orc-0.4.28-1  mingw-w64-x86_64-pango-1.43.0-2
                mingw-w64-x86_64-pcre-8.43-1  mingw-w64-x86_64-pixman-0.38.0-1
                mingw-w64-x86_64-poppler-0.74.0-1  mingw-w64-x86_64-poppler-data-0.4.9-1
                mingw-w64-x86_64-shared-mime-info-1.12-1  mingw-w64-x86_64-wineditline-2.205-1
                mingw-w64-x86_64-zstd-1.3.8-1  mingw-w64-x86_64-evince-3.28.2-3

合計ダウンロード容量:   99.07 MiB
合計インストール容量:  522.88 MiB

:: インストールを行いますか？ [Y/n]
```

うひぃ，大事になっとるがな。
まぁいいや。
このまま `Y` を返してインストールしてしまおう（以降の操作は割愛）。

### [Evince] の起動

Explorer のコンテキストメニュー「プログラムから開く」で導入した `evince.exe` を使って PDF ファイルを開こうとするが，私の環境では何故か上手くいかないので， shell から

```text
$ C:\msys64\mingw64\bin\evince.exe foo.pdf
```

のように起動するしかなさそうだ（引数なしならラウンチゃが開く）。
まぁ [MSYS2] ファイルシステム依存でなくてよかった。

私は [NYAGOS を使っている]({{< ref "/remark/2015/conemu-and-nyagos.md" >}} "ようやく ConEmu と NYAGOS を導入した")のでホームディレクトリの `.nyagos` ファイルに

```lua
nyagos.alias.evince = "c:/msys64/mingw64/bin/evince.exe"
```

とか記述しておけば

```text
$ evince foo.pdf
```

で起動できる。

ちょっと扱いが面倒になってしまったが仕方がない。
普段遣いは [MuPDF] にして， [MuPDF] で上手く表示できない場合や PDF ファイルのプロパティを確認する際は [Evince] を使うって感じで。

早く Windows 捨てたい。

## ブックマーク

- [Windows で evince – 山本昌志のブログ](http://www.yamamo10.jp/yamamoto/wordpress/?p=842)
- [Windowsコマンドプロンプトでaliasを使いたい - ばくのエンジニア日誌](http://bakunyo.hatenablog.com/entry/2014/02/16/Windows%E3%82%B3%E3%83%9E%E3%83%B3%E3%83%89%E3%83%97%E3%83%AD%E3%83%B3%E3%83%97%E3%83%88%E3%81%A7alias%E3%82%92%E4%BD%BF%E3%81%84%E3%81%9F%E3%81%84)

[前回]: {{< relref "./get-a-free-software-pdf-reader.md" >}} "いつの間にか Evince Windows 版がなくなっていた"
[PDFreaders.org]: https://pdfreaders.org/ "Get a Free Software PDF reader!"
[Evince]: http://projects.gnome.org/evince/
[Okular]: https://okular.kde.org/ "Okular - more than a reader"
[SumatraPDF]: https://www.sumatrapdfreader.org/ "Free PDF Reader - Sumatra PDF"
[XpdfReader]: http://www.xpdfreader.com/
[PDF.js]: https://github.com/mozilla/pdf.js "mozilla/pdf.js: PDF Reader in JavaScript"
[qpdfview]: https://launchpad.net/qpdfview "qpdfview in Launchpad"
[GNU GV]: http://www.gnu.org/software/gv/ "GV - GNU Project - Free Software Foundation"
[ePDFView]: http://freshmeat.sourceforge.net/projects/epdfview/ "ePDFView – Freecode"
[MuPDF]: https://mupdf.com/
[zathura]: https://pwmt.org/projects/zathura/
[SyncTeX]: https://texwiki.texjp.org/?SyncTeX "SyncTeX - TeX Wiki"
[MSYS2]: http://msys2.github.io/ "MSYS2 installer"
<!-- eof -->
