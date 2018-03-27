+++
title = "NYAGOS 4.2.5_Beta のリリースと環境変数の扱い"
date =  "2018-03-27T09:51:31+09:00"
description = "このバージョンからバッチファイル実行時の環境変数の扱いが変わるようだ。"
image = "/images/attention/tools.png"
tags  = [ "tools", "nyagos", "shell", "windows" ]
draft = true

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

[NYAGOS] 4.2.5_beta がリリースされた。

- [Release 4.2.5_beta · zetamatta/nyagos](https://github.com/zetamatta/nyagos/releases/tag/4.2.5_beta)

{{% fig-quote title="Release 4.2.5_beta" link="https://github.com/zetamatta/nyagos/releases/tag/4.2.5_beta" lang="en" %}}
- Read the value of environment variables and the current directory that a batchfile changed like CMD.EXE.
- And refactored a lot of source files

----

- CMD.EXE と同様に、バッチファイルが変更した環境変数・カレントディレクトリを読み取るようにした。
- ソースの幾つかを派手にリファクタリングした。
{{% /fig-quote %}}

というわけで，このバージョンからバッチファイル（`*.bat`, `*.cmd`）実行時の環境変数の扱いが変わるようだ。
詳しくは以下を参照のこと。

- [NYAGOS 4.2.5βが、いかにしてバッチファイルでの環境変数の変更取り込みを可能としたか - Qiita](https://qiita.com/zetamatta/items/efff93d92ac2150192fb)

個人的にはバッチファイルで環境変数が汚れるのは好みではなかったので `source` コマンドのみで環境変数を変えられるという仕様は結構気に入っていたのだが，まぁ仕方ない。

## 環境変数の汚染を防ぐには

バッチファイルで環境変数を汚さないようにするには `setlocal ... endlocal` で囲めばよい[^sl1]。
しかも入れ子にできる。
たとえばフィボナッチ数を数え上げる処理はこんな感じ[^de1]。

[^sl1]: 時々 `setlocal` はセットしてるのに `endlocal` がない記述を見かける（特に遅延環境変数の説明とか）。 `endlocal` がなくても別にエラーになったりはしないしちゃんと動くのだが（バッチ終了時に全スコープが閉じられるので），プログラマならスコープは正しく書こうね。
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
