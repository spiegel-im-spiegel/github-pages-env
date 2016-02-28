+++
date = "2016-02-28T14:25:09+09:00"
description = "MSYS の後継（？）とも言える MSYS2 では開発環境を 32bit と 64bit で併設できるらしい。"
draft = true
tags = ["msys2", "gcc", "tools"]
title = "MSYS2 による gcc 開発環境の構築 ― インストールから初期化処理まで"

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

1. [MSYS2 による gcc 開発環境の構築 ― インストールから初期化処理まで]({{< relref "remark/2016/02/gcc-msys2-1.md" >}})
2. [MSYS2 による gcc 開発環境の構築 ― gcc パッケージ群の導入]({{< relref "remark/2016/02/gcc-msys2-2.md" >}})
3. [MSYS2 による gcc 開発環境の構築 ― pgpdump をビルドする]({{< relref "remark/2016/02/gcc-msys2-3.md" >}})

## MSYS2 のインストール

MSYS の後継（？）とも言える [MSYS2] では開発環境を 32bit と 64bit で併設できるらしい。

- [MSYS2 installer](http://msys2.github.io/)
- [msys2](https://github.com/msys2) : GitHub ページ

早速 64bit 版インストーラ（今回は `msys2-x86_64-20160205.exe` を使用）でインストールを行ってみる。

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

[MSYS2] の起動バッチには以下の3つがある（64bit 版をインストールした場合）

- `mingw32_shell.bat`
- `mingw64_shell.bat`
- `msys2_shell.bat`

中身はターミナルエミュレータ（既定で [mintty]）の場所を探して起動するだけだが，環境変数 `MSYSTEM` にそれぞれ以下の値をセットしている

| バッチファイル      | セットする値 |
|:--------------------|:-------------|
| `mingw32_shell.bat` | `MINGW32`    |
| `mingw64_shell.bat` | `MINGW64`    |
| `msys2_shell.bat`   | `MSYS`       |

環境変数 `MSYSTEM` は `/etc/profile` 内で参照される。

```bash
MSYS2_PATH="/usr/local/bin:/usr/bin:/bin"
MANPATH="/usr/local/man:/usr/share/man:/usr/man:/share/man:${MANPATH}"
INFOPATH="/usr/local/info:/usr/share/info:/usr/info:/share/info:${INFOPATH}"
MINGW_MOUNT_POINT=
if [ -n "$MSYSTEM" ]
then
  case "$MSYSTEM" in
    MINGW32)
      MINGW_MOUNT_POINT=/mingw32
      PATH="${MINGW_MOUNT_POINT}/bin:${MSYS2_PATH}:${PATH}"
      PKG_CONFIG_PATH="${MINGW_MOUNT_POINT}/lib/pkgconfig:${MINGW_MOUNT_POINT}/share/pkgconfig"
      ACLOCAL_PATH="${MINGW_MOUNT_POINT}/share/aclocal:/usr/share/aclocal"
      MANPATH="${MINGW_MOUNT_POINT}/share/man:${MANPATH}"
    ;;
    MINGW64)
      MINGW_MOUNT_POINT=/mingw64
      PATH="${MINGW_MOUNT_POINT}/bin:${MSYS2_PATH}:${PATH}"
      PKG_CONFIG_PATH="${MINGW_MOUNT_POINT}/lib/pkgconfig:${MINGW_MOUNT_POINT}/share/pkgconfig"
      ACLOCAL_PATH="${MINGW_MOUNT_POINT}/share/aclocal:/usr/share/aclocal"
      MANPATH="${MINGW_MOUNT_POINT}/share/man:${MANPATH}"
    ;;
    MSYS)
      PATH="${MSYS2_PATH}:/opt/bin:${PATH}"
      PKG_CONFIG_PATH="/usr/lib/pkgconfig:/usr/share/pkgconfig:/lib/pkgconfig"
    ;;
    *)
      PATH="${MSYS2_PATH}:${PATH}"
    ;;
  esac
else
  PATH="${MSYS2_PATH}:${PATH}"
fi
```

`MINGW32` または `MINGW64` の場合は，新たに `/mingw32` または `/mingw64` がパス等に追加されているのが分かるだろう。
また，各モードは[次回]説明するリポジトリと関連している。

### ConEmu から MSYS2 bash を起動する

以上から，環境変数 `MSYSTEM` を適切に設定すれば [MSYS2] の起動バッチを介さずに shell を起動しても構わないことが分かるだろう。
ここでは [ConEmu] から [MSYS2] の bash を起動することを考える。
以下のように Tasks 設定で [MSYS2] の bash を起動するシーケンスを設定すればよい。

{{< fig-img flickr="true" src="https://farm1.staticflickr.com/698/22388493089_73bb752b56.jpg" alt="ConEmu Setting" title="ConEmu Setting" link="https://www.flickr.com/photos/spiegel/22388493089/" >}}

起動シーケンスは以下のとおり。

```text
set MSYSTEM=MSYS & C:\msys64\usr\bin\bash.exe --login -i
```

前半で環境変数を設定し，後半で実際に bash を呼び出している。

### 環境変数 PATH のチューニング

Windows の環境変数は [MSYS2] にも引き継がれる。
`PATH` も同様。
ただし Windows の `PATH` 上のアプリケーションを [MSYS2] 上で動いてほしくない場合もある。
この場合は起動時に `PATH` を書き換える。
先ほどの [ConEmu] の起動シーケンスなら

```text
set PATH=%SystemRoot%System32 & set MSYSTEM=MSYS & C:\msys64\usr\bin\bash.exe --login -i
```

でいいだろう。
Windows 上の特定のツールを使いたい場合はフルパスまたはその alias で指定するほうがスマートである。

## 初期設定

では本題に戻ろう。
`msys2_shell.bat` または環境変数 `MSYSTEM` に `MSYS` をセットして shell を起動し，まずはコア・パッケージを最新のものに更新する。

```text
$ update-core
==> Update package databases...
:: パッケージデータベースの同期中...
 mingw32                             232.6 KiB   720K/s 00:00 [##################] 100%
 mingw32.sig                          96.0   B  93.8K/s 00:00 [##################] 100%
 mingw64                             231.6 KiB   740K/s 00:00 [##################] 100%
 mingw64.sig                          96.0   B  0.00B/s 00:00 [##################] 100%
 msys                                130.7 KiB   165K/s 00:01 [##################] 100%
 msys.sig                             96.0   B  93.8K/s 00:00 [##################] 100%
==> Checking if there are critical packages to upgrade.
==> No updates for core packages.
```

ありゃ，更新対象がなかったか。
ここでパッケージの更新が行われた場合，いったん shell を終了させないといけないらしい。

次に他のインストール済みのパッケージを更新しよう[^pm]。

[^pm]: pacman コマンドについては[次回]説明する。

```text
$ pacman -Su
:: システム全体の更新を開始...
依存関係を解決しています...
衝突するパッケージがないか確認しています...

パッケージ (7) curl-7.47.1-1  file-5.25-1  gnupg-1.4.20-1  libcurl-7.47.1-1
               libopenssl-1.0.2.f-1  mintty-1~2.2.3-1  openssl-1.0.2.f-1

合計ダウンロード容量:   4.40 MiB
合計インストール容量:  20.60 MiB
最終的なアップグレード容量:   0.12 MiB

:: インストールを行いますか？ [Y/n] y
:: パッケージを取得します ...
 libopenssl-1.0.2.f-1-x86_64         815.3 KiB   999K/s 00:01 [##################] 100%
 openssl-1.0.2.f-1-x86_64           1336.9 KiB   918K/s 00:01 [##################] 100%
 libcurl-7.47.1-1-x86_64             185.0 KiB   582K/s 00:00 [##################] 100%
 curl-7.47.1-1-x86_64                595.3 KiB   752K/s 00:01 [##################] 100%
 file-5.25-1-x86_64                  396.5 KiB   831K/s 00:00 [##################] 100%
 gnupg-1.4.20-1-x86_64              1026.9 KiB   806K/s 00:01 [##################] 100%
 mintty-1~2.2.3-1-x86_64             147.2 KiB   462K/s 00:00 [##################] 100%
(7/7) キーリングのキーを確認                                  [##################] 100%
(7/7) パッケージの整合性をチェック                            [##################] 100%
(7/7) パッケージファイルのロード                              [##################] 100%
(7/7) ファイルの衝突をチェック                                [##################] 100%
(7/7) 空き容量を確認                                          [##################] 100%
:: パッケージの変更を処理しています...
(1/7) 更新 libopenssl                                         [##################] 100%
(2/7) 更新 openssl                                            [##################] 100%
(3/7) 更新 libcurl                                            [##################] 100%
(4/7) 更新 curl                                               [##################] 100%
(5/7) 更新 file                                               [##################] 100%
(6/7) 更新 gnupg                                              [##################] 100%
gpg: WARNING: standard input reopened
gpg: /etc/pacman.d/gnupg/trustdb.gpg: 信用データベースができました
gpg: 究極的に信用する鍵が見つかりません
gpg: WARNING: standard input reopened
gpg: Generating pacman keyring master key...
gpg: skipping control `%no-protection' ()
...+++++
..+++++
gpg: 鍵A52980BFを究極的に信用するよう記録しました
gpg: Done
==> 信用データベースを更新...
gpg: WARNING: standard input reopened
gpg: 最小の「ある程度の信用」3、最小の「全面的信用」1、PGP信用モデル
gpg: 深さ: 0  有効性:   1  署名:   0  信用: 0-, 0q, 0n, 0m, 0f, 1u
==> msys2.gpg からキーを追加...
gpg: WARNING: standard input reopened
gpg: 最小の「ある程度の信用」3、最小の「全面的信用」1、PGP信用モデル
gpg: 深さ: 0  有効性:   1  署名:   0  信用: 0-, 0q, 0n, 0m, 0f, 1u
==> キーリングの信頼されたキーに署名...
  -> キーに署名 D55E7A6D7CE9BA1587C0ACACF40D263ECA25678A...
  -> キーに署名 123D4D51A1793859C2BE916BBBE514E53E0D0813...
  -> キーに署名 B91BCF3303284BF90CC043CA9F418C233E652008...
  -> キーに署名 9DD0D4217D75A33B896159E6DA7EF2ABAEEA755C...
==> 所有者信頼値をインポート...
gpg: WARNING: standard input reopened
gpg: setting ownertrust to 4
gpg: setting ownertrust to 4
gpg: setting ownertrust to 4
gpg: inserting ownertrust of 4
==> 信用データベースを更新...
gpg: WARNING: standard input reopened
gpg: 最小の「ある程度の信用」3、最小の「全面的信用」1、PGP信用モデル
gpg: 深さ: 0  有効性:   1  署名:   4  信用: 0-, 0q, 0n, 0m, 0f, 1u
gpg: 深さ: 1  有効性:   4  署名:   3  信用: 0-, 0q, 0n, 4m, 0f, 0u
gpg: 深さ: 2  有効性:   3  署名:   0  信用: 3-, 0q, 0n, 0m, 0f, 0u
(7/7) 更新 mintty                                             [##################] 100%
```

これで初期化は終了。
[次回]へ続く。

## 【おまけ】 Proxy サーバ越しのアクセス

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
pub   2048R/A52980BF 2016-02-28
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

$ gpg --homedir /etc/pacman.d/gnupg --list-keys
```

もしかしたらこちらの `gpg.conf` ファイルも修正する必要があるかもしれない。

[MSYS2]: http://msys2.github.io/ "MSYS2 installer"
[mintty]: https://mintty.github.io/ "Mintty — Cygwin Terminal emulator"
[ConEmu]: https://conemu.github.io/ "ConEmu - Handy Windows Terminal"
[次回]: {{< relref "remark/2016/02/gcc-msys2-2.md" >}} "MSYS2 による gcc 開発環境の構築 ― gcc パッケージ群の導入"
