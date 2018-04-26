+++
title = "MinGW-w64 を導入する"
date = "2018-03-31T17:40:56+09:00"
update = "2018-04-26T09:13:34+09:00"
description = "MinGW-w64 は Windows 用のバイナリを出力可能な GCC を含む開発環境を提供している。もちろん Windows 用のバイナリも存在する。"
image = "/images/attention/kitten.jpg"
tags = ["mingw", "gcc", "tools", "windows"]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[MinGW-w64] は [MinGW (Minimalist GNU for Windows)](http://www.mingw.org/) の後継とも言えるプロジェクトで Windows 用のバイナリを出力可能な [GCC] を含む開発環境を提供している。
[MSYS2] も [MinGW-w64] の成果を取り込んでいるので馴染みのある人もいるだろう。

[MinGW] および [MinGW-w64] はどちらかと言うと Linux 系のプラットフォームで Windows 向けのビルドを行うクロスコンパイラとしての側面が強いが，もちろん Windows 用のバイナリも存在する。
[MinGW-w64] の Windows 用インストーラは以下で取得できる[^sf1]。

[^sf1]: SourceForge からのダウンロードなので取扱いに注意。ページの余計なところを触らないこと（笑）

- [Mingw-builds](http://mingw-w64.org/doku.php/download/mingw-builds)

インストーラを起動すると以下のウィザード画面になる。

{{< fig-img src="https://farm1.staticflickr.com/868/40234864025_5e6b30727f.jpg" title="Installing Mingw-w64 (1)" link="https://www.flickr.com/photos/spiegel/40234864025/" >}}

このまま `[Next >]`。

{{< fig-img src="https://farm1.staticflickr.com/898/41129922281_1cd3e899a2.jpg" title="Installing Mingw-w64 (2)" link="https://www.flickr.com/photos/spiegel/41129922281/" >}}

この画面でインストールする [GCC] の種別を指定する。

- *Version* : [GCC] のバージョン。特に理由がない限り最新版でいいだろう（2018-03-31 時点の最新は 7.3.0）
- *Architecture* : ターゲットシステムのアーキテクチャ。 `i686` か `x86_64` のいずれかを選択する。 `x86_64` が64ビット版なのでご注意を
- *Thread* : スレッドモデル。 `posix` か `win32` のいずれかを選択する。 `win32` は Windows ネイティブ関数を使っていて速いのだが C++ 11 以降の thread, mutex, future が使えなくなる。特に理由がないのであれば `posix` でいいだろう
- *Exception* : 例外処理ハンドリング。アーキテクチャによって以下を選択できる。
    - `i686`
        - `dwarf` : [DWARF](http://ja.wikipedia.org/wiki/DWARF)
        - `sjlj` : SetJump/LongJump
    - `x86_64`
        - `seh` : Structured Exception Handling （Windows ネイティブ。おすすめ）
        - `sjlj` : SetJump/LongJump
- *Build Version* : ビルドバージョン。特に理由がなければ最新版でいいだろう

種別を指定したら `[Next >]`。

{{< fig-img src="https://farm1.staticflickr.com/798/27256804608_3c74aed9e4.jpg" title="Installing Mingw-w64 (3)" link="https://www.flickr.com/photos/spiegel/27256804608/" >}}

インストール先のフォルダを指定する。
既定のままではめっさ長い名前になるので変更したほうがいいかも？

`[Next >]` でインストールを開始する。

{{< fig-img src="https://farm1.staticflickr.com/885/40234864195_35a1d24993.jpg" title="Installing Mingw-w64 (4)" link="https://www.flickr.com/photos/spiegel/40234864195/" >}}

どうもネットから指定した種別に対応する圧縮ファイルをダウンロードしてバイナリを展開しているようだ。
完了したら `[Next >]`。

{{< fig-img src="https://farm1.staticflickr.com/894/40234864325_ea472a14df.jpg" title="Installing Mingw-w64 (5)" link="https://www.flickr.com/photos/spiegel/40234864325/" >}}

これで完了。
お疲れ様でした。

試しに gcc を起動してみる。
こんな感じ。

```text
$ gcc -v
Using built-in specs.
COLLECT_GCC=gcc.exe
COLLECT_LTO_WRAPPER=C:/Program\ Files/mingw-w64/x86_64-7.3.0-posix-seh-rt_v5-rev0/mingw64/bin/../libexec/gcc/x86_64-w64-mingw32/7.3.0/lto-wrapper.exe
Target: x86_64-w64-mingw32
Configured with: ../../../src/gcc-7.3.0/configure --host=x86_64-w64-mingw32 --build=x86_64-w64-mingw32 --target=x86_64-w64-mingw32 --prefix=/mingw64 --with-sysroot=/c/mingw730/x86_64-730-posix-seh-rt_v5-rev0/mingw64 --enable-shared --enable-static --disable-multilib --enable-languages=c,c++,fortran,lto --enable-libstdcxx-time=yes --enable-threads=posix --enable-libgomp --enable-libatomic --enable-lto --enable-graphite --enable-checking=release --enable-fully-dynamic-string --enable-version-specific-runtime-libs --enable-libstdcxx-filesystem-ts=yes --disable-libstdcxx-pch --disable-libstdcxx-debug --enable-bootstrap --disable-rpath --disable-win32-registry --disable-nls --disable-werror --disable-symvers --with-gnu-as --with-gnu-ld --with-arch=nocona --with-tune=core2 --with-libiconv --with-system-zlib --with-gmp=/c/mingw730/prerequisites/x86_64-w64-mingw32-static --with-mpfr=/c/mingw730/prerequisites/x86_64-w64-mingw32-static --with-mpc=/c/mingw730/prerequisites/x86_64-w64-mingw32-static --with-isl=/c/mingw730/prerequisites/x86_64-w64-mingw32-static --with-pkgversion='x86_64-posix-seh-rev0, Built by MinGW-W64 project' --with-bugurl=https://sourceforge.net/projects/mingw-w64 CFLAGS='-O2 -pipe -fno-ident -I/c/mingw730/x86_64-730-posix-seh-rt_v5-rev0/mingw64/opt/include -I/c/mingw730/prerequisites/x86_64-zlib-static/include -I/c/mingw730/prerequisites/x86_64-w64-mingw32-static/include' CXXFLAGS='-O2 -pipe -fno-ident -I/c/mingw730/x86_64-730-posix-seh-rt_v5-rev0/mingw64/opt/include -I/c/mingw730/prerequisites/x86_64-zlib-static/include -I/c/mingw730/prerequisites/x86_64-w64-mingw32-static/include' CPPFLAGS=' -I/c/mingw730/x86_64-730-posix-seh-rt_v5-rev0/mingw64/opt/include -I/c/mingw730/prerequisites/x86_64-zlib-static/include -I/c/mingw730/prerequisites/x86_64-w64-mingw32-static/include' LDFLAGS='-pipe -fno-ident -L/c/mingw730/x86_64-730-posix-seh-rt_v5-rev0/mingw64/opt/lib -L/c/mingw730/prerequisites/x86_64-zlib-static/lib -L/c/mingw730/prerequisites/x86_64-w64-mingw32-static/lib '
Thread model: posix
gcc version 7.3.0 (x86_64-posix-seh-rev0, Built by MinGW-W64 project)
```

インストーラ自体は環境変数を変更しないのだがインストールしたフォルダに `mingw-w64.bat` というのができていて，これを起動すると PATH を追加してコマンドプロンプトを起動する。
常用するのであれば自前で環境変数を変更すればよい。

削除する際はコントロールパネルの「プログラムと機能」から可能だがメチャメチャ分かりにくい名前になっている。

{{< fig-img src="https://farm1.staticflickr.com/789/40235231465_d253312aa0.jpg" title="Uninstalling Mingw-w64" link="https://www.flickr.com/photos/spiegel/40235231465/" >}}

「プログラムと機能」の右上に検索窓があるのでそこで「mingw」と入力すれば上の画面のようになるので，これで該当プログラムを選択して削除すればいいだろう。

## みんな大好き Hello World

では動作確認。
以下のコードを用意する。

```c
#include "stdio.h"

void main(void) {
    printf("hello world\n");
}
```

これをコンパイルして実行する。

```text
$ gcc hello.c

$ a.exe
hello world
```

よーし，うむうむ，よーし。

[GCC]: https://gcc.gnu.org/ "GCC, the GNU Compiler Collection - GNU Project - Free Software Foundation (FSF)"
[MinGW-w64]: http://mingw-w64.org/ "Mingw-w64 - GCC for Windows 64 & 32 bits [mingw-w64]"
[MinGW]: http://www.mingw.org/ "MinGW | Minimalist GNU for Windows"
[MSYS2]: http://msys2.github.io/ "MSYS2 installer"

## 参考文献

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4320026926/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/41W69WGATNL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4320026926/baldandersinf-22/">プログラミング言語C 第2版 ANSI規格準拠</a></dt><dd>B.W. カーニハン D.M. リッチー 石田 晴久 </dd><dd>共立出版 1989-06-15</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4320027485/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4320027485.09._SCTHUMBZZZ_.jpg"  alt="プログラミング言語Cアンサー・ブック 第2版"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4874084141/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4874084141.09._SCTHUMBZZZ_.jpg"  alt="C言語による最新アルゴリズム事典 (ソフトウェアテクノロジー)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774111422/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774111422.09._SCTHUMBZZZ_.jpg"  alt="C言語ポインタ完全制覇 (標準プログラマーズライブラリ)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797304952/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797304952.09._SCTHUMBZZZ_.jpg"  alt="定本 Cプログラマのためのアルゴリズムとデータ構造 (SOFTBANK BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4900900648/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4900900648.09._SCTHUMBZZZ_.jpg"  alt="C実践プログラミング 第3版"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4781908535/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4781908535.09._SCTHUMBZZZ_.jpg"  alt="工科系の数学 (5)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4781908896/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4781908896.09._SCTHUMBZZZ_.jpg"  alt="工科系の数学〈6〉関数論"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4756136494/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4756136494.09._SCTHUMBZZZ_.jpg"  alt="プログラミング作法"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798030147/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798030147.09._SCTHUMBZZZ_.jpg"  alt="苦しんで覚えるC言語"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798101036/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798101036.09._SCTHUMBZZZ_.jpg"  alt="Cの絵本―C言語が好きになる9つの扉"  /></a> </p>
<p class="description">通称 “K&amp;R”。その筋の人々には「バイブル」と呼ばれる名著（当時は）。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-02-19">2017-02-19</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117364/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51ng4usMVYL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117364/baldandersinf-22/">Effective Modern C++ ―C++11/14プログラムを進化させる42項目</a></dt><dd>Scott Meyers 千住 治郎 </dd><dd>オライリージャパン 2015-09-18</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4048694243/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4048694243.09._SCTHUMBZZZ_.jpg"  alt="C++11/14 コア言語"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797375957/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797375957.09._SCTHUMBZZZ_.jpg"  alt="プログラミング言語C++第4版"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4621066099/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4621066099.09._SCTHUMBZZZ_.jpg"  alt="Effective C++ 第3版 (ADDISON-WESLEY PROFESSIONAL COMPUTI)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797384778/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797384778.09._SCTHUMBZZZ_.jpg"  alt="C++のエッセンス"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4048930516/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4048930516.09._SCTHUMBZZZ_.jpg"  alt="C++によるプログラミングの原則と実践"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797376686/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797376686.09._SCTHUMBZZZ_.jpg"  alt="C++テンプレートテクニック 第2版"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4844338900/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4844338900.09._SCTHUMBZZZ_.jpg"  alt="Game Programming Patterns ソフトウェア開発の問題解決メニュー (impress top gear)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117569/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117569.09._SCTHUMBZZZ_.jpg"  alt="Effective Python ―Pythonプログラムを改良する59項目"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774174084/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774174084.09._SCTHUMBZZZ_.jpg"  alt="改訂新版　C++ポケットリファレンス"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4862462928/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4862462928.09._SCTHUMBZZZ_.jpg"  alt="Computer Graphics Gems JP 2015 - コンピュータグラフィックス技術の最前線 -"  /></a> </p>
<p class="description">C++ 再勉強中。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-09-22">2016-09-22</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
