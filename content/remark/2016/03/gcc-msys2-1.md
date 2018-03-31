+++
date = "2016-03-01T22:02:41+09:00"
update = "2018-03-31T23:56:06+09:00"
description = "MSYS の後継（？）とも言える MSYS2 では開発環境を 32bit と 64bit で併設できるらしい。"
draft = false
tags = ["msys2", "mingw", "gcc", "tools", "windows"]
title = "MSYS2 による gcc 開発環境の構築 ― MSYS2 のインストールから初期化処理まで"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

これは Qiita の以下の記事からの再構成である。

- [MSYS2 による gcc 開発環境の構築 - Qiita](http://qiita.com/spiegel-im-spiegel/items/ba4e8d2418bdfe0c8049)
- [MSYS2 による gcc 開発環境の構築（2） - Qiita](http://qiita.com/spiegel-im-spiegel/items/877cbfb970509b272fc1)

ここでは以下の3部構成になっている。

1. [MSYS2 のインストールから初期化処理まで]({{< relref "remark/2016/03/gcc-msys2-1.md" >}}) （← イマココ）
2. [gcc パッケージ群の導入]({{< relref "remark/2016/03/gcc-msys2-2.md" >}})
3. [pgpdump をビルドする]({{< relref "remark/2016/03/gcc-msys2-3.md" >}})

## gcc と MSYS2

MSYS の後継（？）とも言える [MSYS2] では gcc を含む開発環境を 32bit と 64bit で併設できるらしい。

- [MSYS2 installer](http://msys2.github.io/)
- [msys2](https://github.com/msys2) : GitHub ページ

今回は [MSYS2] を使って gcc 開発環境を構築する。

なお，単に gcc が欲しいのであれば [MinGW-w64] から Windows 用のバイナリを取得するほうがお勧めである[^mgw1]。

[^mgw1]: この記事では今まで [TDM-GCC] を勧めてきたが，どうも何年もメンテナンスされていないようだ。したがって今後は [MinGW-w64] を推すことにする。

- [MinGW-w64 を導入する]({{< relref "remark/2018/03/mingw-w64.md" >}})

そうではなく autotools などを含む UNIX 互換の環境が必要なのであれば，今回紹介する [MSYS2] が適切と思われる（UNIX 系のコマンドを Windows にポーティングする場合など）。
また開発環境は要らないが UNIX 互換の動作環境が欲しいだけなら [Git for Windows] を導入する手もある[^gfw]。

[^gfw]: [Git for Windows] には bash などの [MSYS2] サブセットを含んでいるため， [Git for Windows] の bash （通称 git bash）を起動することで UNIX 互換の動作環境を得られる。なお [Git for Windows] の環境は git を動かすための最小限のツールしか入っていないため， Git for Windows SDK も併せて導入し，その中の pacman コマンドで必要に応じてのツールを追加する。 pacman については[次回]説明する。

## MSYS2 のインストール

では早速 64bit 版インストーラ（今回は `msys2-x86_64-20161025.exe` を使用）でインストールを行ってみる。

{{< fig-img flickr="true" src="https://farm2.staticflickr.com/1462/25210363812_7cd554cc78.jpg" title="MSYS2 Installing (1)" link="https://www.flickr.com/photos/spiegel/25210363812/" >}}

わお！ 日本語だ。

{{< fig-img flickr="true" src="https://farm2.staticflickr.com/1540/25032985150_5bc67480e8.jpg" title="MSYS2 Installing (2)" link="https://www.flickr.com/photos/spiegel/25032985150/" >}}

{{< fig-img flickr="true" src="https://farm2.staticflickr.com/1594/24960938719_0dba3ebf95.jpg" title="MSYS2 Installing (3)" link="https://www.flickr.com/photos/spiegel/24960938719/" >}}

{{< fig-img flickr="true" src="https://farm2.staticflickr.com/1621/25210363972_b14f16a013.jpg" title="MSYS2 Installing (4)" link="https://www.flickr.com/photos/spiegel/25210363972/" >}}

{{< fig-img flickr="true" src="https://farm2.staticflickr.com/1473/25235500391_d1bf4d5f3b.jpg" title="MSYS2 Installing (5)" link="https://www.flickr.com/photos/spiegel/25235500391/" >}}

このまま「完了」するとターミナル（[mintty]）が起動する。
のだが，ここでちょっと寄り道する。
「MSYS2 64bit を実行中」のチェックを外して「完了」しよう。

### MSYS2 起動モード

[MSYS2] には3つの起動モジュールがある（64bit 版をインストールした場合）。

- `mingw32.exe`
- `mingw64.exe`
- `msys2.exe`

またそれぞれに ini ファイルが用意されている。
たとえば `msys2.exe` であればこんな感じ。

```ini
#MSYS=winsymlinks:nativestrict
#MSYS=error_start:mingw64/bin/qtcreator.exe|-debug|<process-id>
#CHERE_INVOKING=1
#MSYS2_PATH_TYPE=inherit
MSYSTEM=MSYS
```

どうやら ini ファイルの内容をそのまま環境変数として渡しているようだ。
このうち `MSYSTEM` に注目すると以下のようになっている。

| バッチファイル | セットする値 |
|:---------------|:-------------|
| `mingw32.exe`  | `MINGW32`    |
| `mingw64.exe`  | `MINGW64`    |
| `msys2.exe`    | `MSYS`       |

`MSYSTEM` は `/etc/profile` 内で参照される。
以下に `/etc/profile` の一部を引用する。

```test
MSYS2_PATH="/usr/local/bin:/usr/bin:/bin"
MANPATH='/usr/local/man:/usr/share/man:/usr/man:/share/man'
INFOPATH='/usr/local/info:/usr/share/info:/usr/info:/share/info'

case "${MSYS2_PATH_TYPE:-minimal}" in
  strict)
    # Do not inherit any path configuration, and allow for full customization
    # of external path. This is supposed to be used in special cases such as
    # debugging without need to change this file, but not daily usage.
    unset ORIGINAL_PATH
    ;;
  inherit)
    # Inherit previous path. Note that this will make all of the Windows path
    # available in current shell, with possible interference in project builds.
    ORIGINAL_PATH="${ORIGINAL_PATH:-${PATH}}"
    ;;
  *)
    # Do not inherit any path configuration but configure a default Windows path
    # suitable for normal usage with minimal external interference.
    WIN_ROOT="$(PATH=${MSYS2_PATH} cygpath -Wu)"
    ORIGINAL_PATH="${WIN_ROOT}/System32:${WIN_ROOT}:${WIN_ROOT}/System32/Wbem:${WIN_ROOT}/System32/WindowsPowerShell/v1.0/"
esac

unset MINGW_MOUNT_POINT
source '/etc/msystem'
case "${MSYSTEM}" in
MINGW32)
  MINGW_MOUNT_POINT="${MINGW_PREFIX}"
  PATH="${MINGW_MOUNT_POINT}/bin:${MSYS2_PATH}${ORIGINAL_PATH:+:${ORIGINAL_PATH}}"
  PKG_CONFIG_PATH="${MINGW_MOUNT_POINT}/lib/pkgconfig:${MINGW_MOUNT_POINT}/share/pkgconfig"
  ACLOCAL_PATH="${MINGW_MOUNT_POINT}/share/aclocal:/usr/share/aclocal"
  MANPATH="${MINGW_MOUNT_POINT}/share/man:${MANPATH}"
  ;;
MINGW64)
  MINGW_MOUNT_POINT="${MINGW_PREFIX}"
  PATH="${MINGW_MOUNT_POINT}/bin:${MSYS2_PATH}${ORIGINAL_PATH:+:${ORIGINAL_PATH}}"
  PKG_CONFIG_PATH="${MINGW_MOUNT_POINT}/lib/pkgconfig:${MINGW_MOUNT_POINT}/share/pkgconfig"
  ACLOCAL_PATH="${MINGW_MOUNT_POINT}/share/aclocal:/usr/share/aclocal"
  MANPATH="${MINGW_MOUNT_POINT}/share/man:${MANPATH}"
  ;;
*)
  PATH="${MSYS2_PATH}:/opt/bin${ORIGINAL_PATH:+:${ORIGINAL_PATH}}"
  PKG_CONFIG_PATH="/usr/lib/pkgconfig:/usr/share/pkgconfig:/lib/pkgconfig"
esac
```

`MINGW32` または `MINGW64` の場合は，新たに `/mingw32` または `/mingw64` がパス等に追加されているのがお分かりだろうか。

### ConEmu から MSYS2 bash を起動する

以上から，環境変数 `MSYSTEM` を適切に設定すれば [MSYS2] の起動モジュールを介さずに shell を起動しても構わないことが分かる。
ここでは [ConEmu] から [MSYS2] の bash を起動することを考える。
以下のように Tasks 設定で [MSYS2] の bash を起動するシーケンスを設定すればよい。

{{< fig-img flickr="true" src="https://farm1.staticflickr.com/698/22388493089_73bb752b56.jpg" alt="ConEmu Setting" title="ConEmu Setting" link="https://www.flickr.com/photos/spiegel/22388493089/" >}}

起動シーケンスは以下のとおり。

```text
set MSYSTEM=MSYS & chcp 65001 & C:\msys64\usr\bin\bash.exe --login -i -new_console:C:"C:\msys64\msys2.ico"
```

前半で環境変数とコードページ（65001 は UTF-8）を設定し，後半で実際に bash を呼び出している。

## 初期設定

では本題に戻ろう。
`msys2.exe` または環境変数 `MSYSTEM` に `MSYS` をセットして shell を起動し，まずはコア・パッケージを最新のものに更新する。

```text
$ pacman -Syu
:: パッケージデータベースの同期中...
 mingw32                                              232.6 KiB   720K/s 00:00 [##################] 100%
 mingw32.sig                                           96.0   B  93.8K/s 00:00 [##################] 100%
 mingw64                                              231.6 KiB   740K/s 00:00 [##################] 100%
 mingw64.sig                                           96.0   B  0.00B/s 00:00 [##################] 100%
 msys                                                 130.7 KiB   165K/s 00:01 [##################] 100%
 msys.sig                                              96.0   B  93.8K/s 00:00 [##################] 100%
==> Checking if there are critical packages to upgrade.
pacman 5.0.0.6348.cc5a8f1-1 -> 5.0.1.6388.dcb2397-1
==> Core packages require updating.
==> Please close all other MSYS2 derived windows (e.g. terminal
==> windows, Bash sessions, etc) before proceeding.
==> 警告: When the update has completed, you MUST close this MSYS2 window
==> 警告: (use Alt-F4 or red [ X ], etc.), rather than 'exit'!!!
Press [Enter] key when ready to start update...
==> Updating core packages...
警告: bash-4.3.042-4 は最新です -- スキップ
警告: msys2-runtime-2.4.1.16860.40c26fc-1 は最新です -- スキップ
依存関係を解決しています...
衝突するパッケージがないか確認しています...

パッケージ (2) msys2-runtime-devel-2.4.1.16860.40c26fc-1  pacman-5.0.1.6388.dcb2397-1

合計ダウンロード容量:  10.34 MiB
合計インストール容量:  54.72 MiB
最終的なアップグレード容量:  21.12 MiB

:: インストールを行いますか？ [Y/n]
:: パッケージを取得します ...
 pacman-5.0.1.6388.dcb2397-1-x86_64                     6.8 MiB  1248K/s 00:06 [##################] 100%
 msys2-runtime-devel-2.4.1.16860.40c26fc-1-x86_64       3.5 MiB  1027K/s 00:04 [##################] 100%
(2/2) キーリングのキーを確認                                                   [##################] 100%
(2/2) パッケージの整合性をチェック                                             [##################] 100%
(2/2) パッケージファイルのロード                                               [##################] 100%
(2/2) ファイルの衝突をチェック                                                 [##################] 100%
(2/2) 空き容量を確認                                                           [##################] 100%
:: パッケージの変更を処理しています...
(1/2) 更新 pacman                                                              [##################] 100%
(2/2) インストール msys2-runtime-devel                                         [##################] 100%
Please close this window.
```

おや。
ツールが更新された。
`pacman -Syu` で更新があった場合，いったん shell を閉じて起動し直す。
このとき `exit` コマンドで終了するのではなく，強制終了する。

Shell を再起動したら他のインストール済みのパッケージを更新しよう。

```text
$ pacman -Su
:: システム全体の更新を開始...
依存関係を解決しています...
衝突するパッケージがないか確認しています...

パッケージ (7) curl-7.47.1-1  file-5.25-1  gnupg-1.4.20-1  libcurl-7.47.1-1  libopenssl-1.0.2.f-1
               mintty-1~2.2.3-1  openssl-1.0.2.f-1

合計ダウンロード容量:   4.40 MiB
合計インストール容量:  20.60 MiB
最終的なアップグレード容量:   0.12 MiB

:: インストールを行いますか？ [Y/n] y
:: パッケージを取得します...
 libopenssl-1.0.2.f-1-x86_64                          815.3 KiB   796K/s 00:01 [##################] 100%
 openssl-1.0.2.f-1-x86_64                            1336.9 KiB   977K/s 00:01 [##################] 100%
 libcurl-7.47.1-1-x86_64                              185.0 KiB   269K/s 00:01 [##################] 100%
 curl-7.47.1-1-x86_64                                 595.3 KiB   855K/s 00:01 [##################] 100%
 file-5.25-1-x86_64                                   396.5 KiB   767K/s 00:01 [##################] 100%
 gnupg-1.4.20-1-x86_64                               1026.9 KiB   672K/s 00:02 [##################] 100%
 mintty-1~2.2.3-1-x86_64                              147.2 KiB   436K/s 00:00 [##################] 100%
(7/7) キーリングのキーを確認                                                   [##################] 100%
(7/7) パッケージの整合性をチェック                                             [##################] 100%
(7/7) パッケージファイルのロード                                               [##################] 100%
(7/7) ファイルの衝突をチェック                                                 [##################] 100%
(7/7) 空き容量を確認                                                           [##################] 100%
:: パッケージの変更を処理しています...
(1/7) 更新 libopenssl                                                          [##################] 100%
(2/7) 更新 openssl                                                             [##################] 100%
(3/7) 更新 libcurl                                                             [##################] 100%
(4/7) 更新 curl                                                                [##################] 100%
(5/7) 更新 file                                                                [##################] 100%
(6/7) 更新 gnupg                                                               [##################] 100%
gpg: /etc/pacman.d/gnupg/trustdb.gpg: 信用データベースができました
gpg: 究極的に信用する鍵が見つかりません
gpg: Generating pacman keyring master key...
gpg: skipping control `%no-protection' ()
.+++++
+++++
gpg: 鍵CD81616Cを究極的に信用するよう記録しました
gpg: Done
==> 信用データベースを更新...
gpg: 最小の「ある程度の信用」3、最小の「全面的信用」1、PGP信用モデル
gpg: 深さ: 0  有効性:   1  署名:   0  信用: 0-, 0q, 0n, 0m, 0f, 1u
==> msys2.gpg からキーを追加...
gpg: 最小の「ある程度の信用」3、最小の「全面的信用」1、PGP信用モデル
gpg: 深さ: 0  有効性:   1  署名:   0  信用: 0-, 0q, 0n, 0m, 0f, 1u
==> キーリングの信頼されたキーに署名...
  -> キーに署名 D55E7A6D7CE9BA1587C0ACACF40D263ECA25678A...
  -> キーに署名 123D4D51A1793859C2BE916BBBE514E53E0D0813...
  -> キーに署名 B91BCF3303284BF90CC043CA9F418C233E652008...
  -> キーに署名 9DD0D4217D75A33B896159E6DA7EF2ABAEEA755C...
==> 所有者信頼値をインポート...
gpg: setting ownertrust to 4
gpg: setting ownertrust to 4
gpg: setting ownertrust to 4
gpg: inserting ownertrust of 4
==> 信用データベースを更新...
gpg: 最小の「ある程度の信用」3、最小の「全面的信用」1、PGP信用モデル
gpg: 深さ: 0  有効性:   1  署名:   4  信用: 0-, 0q, 0n, 0m, 0f, 1u
gpg: 深さ: 1  有効性:   4  署名:   3  信用: 0-, 0q, 0n, 4m, 0f, 0u
gpg: 深さ: 2  有効性:   3  署名:   0  信用: 3-, 0q, 0n, 0m, 0f, 0u
(7/7) 更新 mintty                                                              [##################] 100%
```

これで初期化は終了。
[次回]へ続く。

## 【おまけ1】 ホームディレクトリについて

[MSYS2] のホームディレクトリは，インストールフォルダ直下の `home\username` フォルダになるが， Windows 上で自前のホームディレクトリを作っている場合は，その場所を `/etc/fstab` でマウントする手もある。

```text
C:/home	/home
```

この場合，元のフォルダにある `.profile` ファイル等の移動も忘れないこと。

## 【おまけ2】 Proxy サーバ越しのアクセス

Proxy サーバ越しに curl を使う場合は `~/.curlrc` に以下の記述を追加する。

```text
proxy-user = "username:password"
proxy = "http://proxy.example.com:8080"
```

パッケージ管理コマンドである `pacman` を Proxy サーバ越しに使う場合は，上述の通りに curl の設定をしたうえで `/etc/pacman.conf` ファイルの以下の記述を探し，コメントを解除する。

```text
#XferCommand = /usr/bin/curl -C - -f %u > %o
```

GnuPG で鍵サーバにアクセスする場合も， Proxy サーバ越しにアクセスするのであれば， `~/.gnupg/gpg.conf` に以下のオプションを追記する。

```text:gpg.conf
keyserver-options http-proxy=http://username:password@proxy.example.com:8080/
```

実は `/etc/pacman.d/gnupg` ディレクトリにも `gpg.conf` ファイルがあって `pacman-key` コマンドはこちらを見てるっぽい[^gpg]。

[^gpg]: ちなみに `/etc/pacman.d/gnupg` ディレクトリにある鍵束を GnuPG から直接見る場合には `gpg --homedir /etc/pacman.d/gnupg --list-keys` とする。

```text
$ pacman-key -l
/etc/pacman.d/gnupg/pubring.gpg
-------------------------------
pub   2048R/CD81616C 2016-03-01
uid                  Pacman Keyring Master Key <pacman@localhost>

pub   2048R/CA25678A 2014-09-28
uid                  Alexey Pavlov (Alexpux) <alexey.pawlow@gmail.com>
sub   2048R/2BBF340E 2014-09-28

pub   4096R/AEEA755C 2014-10-04
uid                  Martell Malone (martell) <martellmalone@gmail.com>
sub   4096R/2A292C03 2014-10-04

pub   4096R/3E0D0813 2014-09-28
uid                  Ray Donnelly (MSYS2 Developer - master key) <mingw.android@gmail.com>
sub   4096R/8603AA9D 2014-09-28

pub   2048R/3E652008 2014-09-29
uid                  Ignacio Casal Quinteiro <icquinteiro@gmail.com>
sub   2048R/64D62A76 2014-09-29

pub   2048D/A47D45A1 2013-11-11
uid                  Alexey Pavlov (Alexpux) <alexpux@gmail.com>
sub   2048g/31CF7700 2013-11-11

pub   4096R/2C51581E 2015-07-22
uid                  Martell Malone (MSYS2 Developer) <martellmalone@gmail.com>
sub   4096R/282D6707 2015-07-22

pub   4096R/4CA56930 2014-09-28
uid                  Ray Donnelly (MSYS2 Developer) <mingw.android@gmail.com>
sub   4096R/576CF231 2014-09-28
```

もしかしたらこちらの `gpg.conf` ファイルも修正する必要があるかもしれない（未確認）。

## 関連するブックマーク

- [MSYS2で快適なターミナル生活 - Qiita](http://qiita.com/Ted-HM/items/4f2feb9fdacb6c72083c)

[MSYS2]: http://msys2.github.io/ "MSYS2 installer"
[TDM-GCC]: http://tdm-gcc.tdragon.net/ "TDM-GCC"
[MinGW-w64]: http://mingw-w64.org/ "Mingw-w64 - GCC for Windows 64 & 32 bits [mingw-w64]"
[Git for Windows]: http://git-for-windows.github.io/ "Git for Windows"
[mintty]: https://mintty.github.io/ "Mintty — Cygwin Terminal emulator"
[ConEmu]: https://conemu.github.io/ "ConEmu - Handy Windows Terminal"
[次回]: {{< relref "remark/2016/03/gcc-msys2-2.md" >}} "MSYS2 による gcc 開発環境の構築 ― gcc パッケージ群の導入"
