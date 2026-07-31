+++
title = "MinGW-w64 を導入する"
date = "2018-03-31T17:40:56+09:00"
update = "2018-12-07T08:41:56+09:00"
description = "MinGW-w64 は Windows 用のバイナリを出力可能な GCC を含む開発環境を提供している。もちろん Windows 用のバイナリも存在する。"
image = "/images/attention/kitten.jpg"
tags = ["mingw", "gcc", "tools", "windows"]

[scripts]
  mathjax = false
  mermaidjs = false
+++

[MinGW-w64] は [MinGW (Minimalist GNU for Windows)](http://www.mingw.org/) の後継とも言えるプロジェクトで Windows 用のバイナリを出力可能な [GCC] を含む開発環境を提供している。
[MSYS2] も [MinGW-w64] の成果を取り込んでいるので馴染みのある人もいるだろう[^msys2]。

[^msys2]: [MSYS2] のインストールについては拙文「[MSYS2 による gcc 開発環境の構築 ― MSYS2 のインストールから初期化処理まで]({{< ref "/remark/2016/03/gcc-msys2-1.md" >}})」を参照のこと。

[MinGW] および [MinGW-w64] はどちらかと言うと Linux 系のプラットフォームで Windows 向けのビルドを行うクロスコンパイラとしての側面が強いが，もちろん Windows 用のバイナリも存在する。
[MinGW-w64] の Windows 用インストーラは以下で取得できる[^sf1]。

[^sf1]: SourceForge からのダウンロードなので取扱いに注意。ページの余計なところを触らないこと（笑）

- [Downloads - Mingw-w64](https://www.mingw-w64.org/downloads/#mingw-builds)

インストーラを起動すると以下のウィザード画面になる。

{{< fig-img src="./mingw-w64-install-1.png" title="Installing Mingw-w64 (1)" link="./mingw-w64-install-1.png" width="527" >}}

このまま `[Next >]`。

{{< fig-img src="./mingw-w64-install-2.png" title="Installing Mingw-w64 (2)" link="./mingw-w64-install-2.png" width="527" >}}

この画面でインストールする [GCC] の種別を指定する。

- *Version* : [GCC] のバージョン。特に理由がない限り最新版でいいだろう（2021-10-01 時点の最新は 8.1.0）
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

{{< fig-img src="./mingw-w64-install-4.png" title="Installing Mingw-w64 (3)" link="./mingw-w64-install-4.png" width="527" >}}

インストール先のフォルダを指定する。
既定のままではめっさ長い名前になるので変更したほうがいいかも？

`[Next >]` でインストールを開始する。

{{< fig-img src="./mingw-w64-install-6.png" title="Installing Mingw-w64 (4)" link="./mingw-w64-install-6.png" width="527" >}}

どうもネットから指定した種別に対応する圧縮ファイルをダウンロードしてバイナリを展開しているようだ。
完了したら `[Next >]`。

{{< fig-img src="./mingw-w64-install-7.png" title="Installing Mingw-w64 (5)" link="./mingw-w64-install-7.png" width="527" >}}

これで完了。
お疲れ様でした。

試しに gcc を起動してみる。
こんな感じ。

```text
$ gcc -v
Using built-in specs.
COLLECT_GCC=gcc
COLLECT_LTO_WRAPPER=C:/Program\ Files/mingw-w64/latest/mingw64/bin/../libexec/gcc/x86_64-w64-mingw32/8.1.0/lto-wrapper.exe
Target: x86_64-w64-mingw32
Configured with: ../../../src/gcc-8.1.0/configure --host=x86_64-w64-mingw32 --build=x86_64-w64-mingw32 --target=x86_64-w64-mingw32 --prefix=/mingw64 --with-sysroot=/c/mingw810/x86_64-810-posix-seh-rt_v6-rev0/mingw64 --enable-shared --enable-static --disable-multilib --enable-languages=c,c++,fortran,lto --enable-libstdcxx-time=yes --enable-threads=posix --enable-libgomp --enable-libatomic --enable-lto --enable-graphite --enable-checking=release --enable-fully-dynamic-string --enable-version-specific-runtime-libs --disable-libstdcxx-pch --disable-libstdcxx-debug --enable-bootstrap --disable-rpath --disable-win32-registry --disable-nls --disable-werror --disable-symvers --with-gnu-as --with-gnu-ld --with-arch=nocona --with-tune=core2 --with-libiconv --with-system-zlib --with-gmp=/c/mingw810/prerequisites/x86_64-w64-mingw32-static --with-mpfr=/c/mingw810/prerequisites/x86_64-w64-mingw32-static --with-mpc=/c/mingw810/prerequisites/x86_64-w64-mingw32-static --with-isl=/c/mingw810/prerequisites/x86_64-w64-mingw32-static --with-pkgversion='x86_64-posix-seh-rev0, Built by MinGW-W64 project' --with-bugurl=https://sourceforge.net/projects/mingw-w64 CFLAGS='-O2 -pipe -fno-ident -I/c/mingw810/x86_64-810-posix-seh-rt_v6-rev0/mingw64/opt/include -I/c/mingw810/prerequisites/x86_64-zlib-static/include -I/c/mingw810/prerequisites/x86_64-w64-mingw32-static/include' CXXFLAGS='-O2 -pipe -fno-ident -I/c/mingw810/x86_64-810-posix-seh-rt_v6-rev0/mingw64/opt/include -I/c/mingw810/prerequisites/x86_64-zlib-static/include -I/c/mingw810/prerequisites/x86_64-w64-mingw32-static/include' CPPFLAGS=' -I/c/mingw810/x86_64-810-posix-seh-rt_v6-rev0/mingw64/opt/include -I/c/mingw810/prerequisites/x86_64-zlib-static/include -I/c/mingw810/prerequisites/x86_64-w64-mingw32-static/include' LDFLAGS='-pipe -fno-ident -L/c/mingw810/x86_64-810-posix-seh-rt_v6-rev0/mingw64/opt/lib -L/c/mingw810/prerequisites/x86_64-zlib-static/lib -L/c/mingw810/prerequisites/x86_64-w64-mingw32-static/lib '
Thread model: posix
gcc version 8.1.0 (x86_64-posix-seh-rev0, Built by MinGW-W64 project)
```

インストーラ自体は環境変数を変更しないのだがインストールしたフォルダに `mingw-w64.bat` というのができていて，これを起動すると PATH を追加してコマンドプロンプトを起動する。
常用するのであれば自前で環境変数を変更すればよい。

削除する際はコントロールパネルの「プログラムと機能」から可能だがメチャメチャ分かりにくい名前になっている。

{{< fig-img src="https://photo.baldanders.info/flickr/image/40235231465_m.png" title="Uninstalling Mingw-w64" link="https://photo.baldanders.info/flickr/40235231465/" >}}

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

## 参考図書

{{% review-paapi "4320026926" %}} <!-- プログラミング言語C -->
{{% review-paapi "4873117364" %}} <!-- Effective Modern C++ -->
