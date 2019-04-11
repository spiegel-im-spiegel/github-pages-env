+++
title = "NYAGOS 4.2.5 のリリースと環境変数の扱い"
date = "2018-03-27T19:47:59+09:00"
update = "2018-04-18T09:57:00+09:00"
description = "このバージョンからバッチファイル実行時の環境変数の扱いが変わるようだ。"
image = "/images/attention/tools.png"
tags  = [ "tools", "nyagos", "shell", "windows" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
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
  mathjax = false
  mermaidjs = false
+++

（正式版がリリースされたので改題しました）

[NYAGOS] 4.2.5 がリリースされた。

- [Release 4.2.5_beta · zetamatta/nyagos](https://github.com/zetamatta/nyagos/releases/tag/4.2.5_beta)
- [Release 4.2.5_beta2 · zetamatta/nyagos](https://github.com/zetamatta/nyagos/releases/tag/4.2.5_beta2)
- [Release 4.2.5_0 · zetamatta/nyagos](https://github.com/zetamatta/nyagos/releases/tag/4.2.5_0)

以下に変更点をまとめて紹介する。

{{% fig-quote-md title="Release 4.2.5" link="https://github.com/zetamatta/nyagos/releases/tag/4.2.5_0" lang="en" %}}
- Read the value of environment variables and the current directory that a batchfile changed like CMD.EXE.
- And refactored a lot of source files
- Fix: #296 the batchfile could not be executed when the username contains multibyte-character.
    - Fix that the encoding of the temporary batchfile was UTF8.
    - Fix that the end of the each line of the temporary batchfile was LF not CRLF.
- Fix: #297 running the batchfile includes exit without /b option, an error occurs
- Add Lua-flag: nyagos.option.usesource. When it is false, batchfiles can not change nyagos's environment variables and directory.(default:true)

----

- CMD.EXE と同様に、バッチファイルが変更した環境変数・カレントディレクトリを読み取るようにした。
- ソースの幾つかを派手にリファクタリングした。
- #296 ユーザ名にマルチバイト文字が入っていると、バッチが正常動作しない不具合を修正
    - 一時バッチファイルのエンコーディングが UTF8 になっていた
    - 一時バッチファイルの改行コードが CRLF ではなく LF になっていた
- #297 /b なしの exit をバッチファイルが実行した時、一時ファイルが無い旨のエラーがでていた
- luaフラグ nyagos.option.usesource を追加。false の時、バッチファイルは NYAGOS の環境変数を変更できなくなる(default:true)
{{% /fig-quote-md %}}

というわけで，このバージョンからバッチファイル（`*.bat`, `*.cmd`）実行時の環境変数の扱いが変わるようだ。
詳しくは以下を参照のこと。

- [NYAGOS 4.2.5βが、いかにしてバッチファイルでの環境変数の変更取り込みを可能としたか - Qiita](https://qiita.com/zetamatta/items/efff93d92ac2150192fb)

個人的にはバッチファイルで環境変数が汚れるのは好みではなかったので `source` コマンドのみで環境変数を変えられるという仕様は結構気に入っていたのだが，まぁいいか。
問題ない。

{{% md-box %}}
**【追記 2018-03-31】** `nyagos.option.usesource` オプションを追加していただいた。
これを `false` にすれば従来どおり `source` コマンドのみで環境変数を変更できる。
ホームディレクトリの `.nyagos` に追記しておけばいいだろう。
{{% /md-box %}}

## 【追記 2018-04-18】 [NYAGOS] 4.2.5_1 がリリース

- [Release 4.2.5_1 · zetamatta/nyagos](https://github.com/zetamatta/nyagos/releases/tag/4.2.5_1)

不具合の修正のみ。

{{% fig-quote-md title="Release 4.2.5_1" link="https://github.com/zetamatta/nyagos/releases/tag/4.2.5_1" lang="en" %}}
- Fix: `if [not] errorlevel N` did not work on block-if.
- Fix: that `ls -1F` did not show the indicator such as `/`,`*` or `@`.
- Fix: the problem that executables reparse-pointed but not symbolic-linked can not be found. ★
- Fix: `ls -F` marked '`@`' to files and directories which ar reparse-pointed but not symbolic-link nor junction
- Changed the error message when the command history is called in `_nyagos`

★ This is the reason that executables in OneDrive can not be called.

----

- ブロックif で `if [not] errorlevel N` が動かなかった不具合を修正
- リパースポイント先の実行ファイルが見付からなくなっている問題を修正（※）
- `ls -1F` が `/`,`*` や `@` といったインジケーターを出力しない問題を修正
- `ls -F` が「リパースポイントではあるが、ジャンクション、シンボリックリンクでないファイル・ディレクトリ」に `@` マークをつけていた不具合を修正
- `_nyagos` で history コマンドを使った時のエラーメッセージを変更

（※ OneDrive の下においた実行ファイルが実行できなくなる問題の原因がコレです）
{{% /fig-quote-md %}}

## 【付録】 環境変数の汚染を防ぐには

バッチファイルで環境変数を汚さないようにするには `setlocal ... endlocal` で囲めばよい[^sl1]。
しかも入れ子にできる。
たとえばフィボナッチ数を数え上げる処理はこんな感じ[^de1]。

[^sl1]: 時々 `setlocal` と `endlocal` がペアになっていない記述を見かける（特に遅延環境変数の説明とか）。 `endlocal` がなくても別にエラーになったりはしないしちゃんと動くのだが（バッチ終了時に全スコープが閉じられるので），プログラマならスコープは正しく書こうね。
[^de1]: コード内の `enabledelayedexpansion` は遅延環境変数を有効にするオプション。 `!fib!` などがこれに該当する。 `for` 文や `if` 文の `( ... )` で囲まれた部分で環境変数を操作する場合に必要。遅延環境変数は `setlocal ... endlocal` スコープ内でしか有効にできない点に注意。

```bat
@echo off
setlocal

setlocal enabledelayedexpansion

set /a n = %~1
set /a b1 = 0
set /a b2 = 0

if %n% == 1 echo 1: %b2% && goto :end
echo 1: %b2%
set /a b2 = 1
if %n% == 2 echo 2: %b2% && goto :end
echo 2: %b2%
for /l %%i in (3, 1, %n%) do (
    set /a fib = !b1! + !b2!
    set /a b1 = !b2!
    set /a b2 = !fib!
    echo %%i: !b2!
)

:end
endlocal && set /a fib = %b2%

echo %~1th Fibonacci number is %fib%
endlocal 
```

`endlocal && set /a fib = %b2%` で `setlocal ... endlocal` スコープの外に値を持ち出している点に注目。

これを実行すると以下のようになる。

```text
C:> fibonacci.cmd 10
1: 0
2: 1
3: 1
4: 2
5: 3
6: 5
7: 8
8: 13
9: 21
10: 34
10th Fibonacci number is 34

C:> set fib
環境変数 fib が定義されていません

C:> set b2
環境変数 b2 が定義されていません
```

## ブックマーク

- [バッチファイルで、setlocal～endlocal内での変数の値を外部に引き継ぎたい！ - IIJIMASの日記](http://d.hatena.ne.jp/IIJIMAS/20101023/1287772847)
- [.bat（バッチファイル）のforコマンド解説。 - Qiita](https://qiita.com/sawa_tsuka/items/67be34bab1fdf3fb87f9)
- [バッチファイル界の魔境『遅延環境変数』に挑む（おまけもあるよ） - Qiita](https://qiita.com/sawa_tsuka/items/c7c477cacf8c97792e17)

[NYAGOS]: https://github.com/zetamatta/nyagos/ "zetamatta/nyagos: NYAGOS - The hybrid UNIXLike Commandline Shell for Windows"
