+++
title = "call と start"
date =  "2021-05-15T14:49:33+09:00"
description = "先日書いた Zenn Scrap の内容。"
image = "/images/attention/kitten.jpg"
tags = [ "windows", "tools" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先日 [Zenn の Scrap](https://zenn.dev/spiegel/scraps "Spiegelのスクラップ一覧") に書いた記事だが， Scrap は基本的に書き捨てで最終的にはアーカイブ（非表示）化する予定なので，改めてこちらのブログに覚え書きとして記しておく。

いやね。
バッチファイルって使い方を忘れちゃうのよ。
たとえば「コマンドプロンプトのウィンドウを非表示（または最小化）の状態でコマンド実行するのってどうすんだっけ？」って感じ。

起動するコマンドを細かく制御したいならコマンドプロンプト（`cmd.exe`）の内蔵コマンド `start` を使う。

```text
> start /?
指定されたプログラムまたはコマンドを実行するためにウィンドウを開きます。

START ["タイトル"] [/D パス] [/I] [/MIN] [/MAX] [/SEPARATE | /SHARED]
      [/LOW | /NORMAL | /HIGH | /REALTIME | /ABOVENORMAL | /BELOWNORMAL]
      [/NODE <NUMA ノード>] [/AFFINITY <16 進数の関係マスク>] [/WAIT] [/B]
      [コマンド/プログラム] [パラメーター]

    "タイトル"  ウィンドウのタイトル バーに表示するタイトル。
    パス        開始するディレクトリ。
    B           新しいウィンドウを作成せずにアプリケーションを起動します。
                アプリケーションは Ctrl + C を無視します。
                アプリケーションで Ctr l+ C を有効にしていない場合、
                Ctrl + Break がアプリケーションを中断する唯一の方法です。
    I           新しい環境は、現在の環境ではなく、cmd.exe に渡された元の環境に
                なります。
    MIN         ウィンドウを最小化の状態で起動します。
    MAX         ウィンドウを最大表示の状態で起動します。
    SEPARATE    16 ビットの Windows プログラムを別メモリ領域で起動します。
    SHARED      16 ビットの Windows プログラムを共有メモリ領域で起動します。
    LOW         IDLE 優先度クラスでアプリケーションを起動します。
    NORMAL      NORMAL 優先度クラスでアプリケーションを起動します。
    HIGH        HIGH 優先度クラスでアプリケーションを起動します。
    REALTIME    REALTIME 優先度クラスでアプリケーションを起動します。
    ABOVENORMAL ABOVENORMAL 優先度クラスでアプリケーションを起動します。
    BELOWNORMAL BELOWNORMAL 優先度クラスでアプリケーションを起動します。
    NODE        優先 NUMA (Non-Uniform Memory Architecture) ノードを 10 進の
                整数で指定します。
    AFFINITY    プロセッサの関係マスクを 16 進数で指定します。
                プロセスはこれらのプロセッサで実行されるように制限されます。

                /AFFINITY と /NODE を組み合わせると、関係マスクは異なって
                解釈されます。NUMA ノードのプロセッサ マスクを右にシフトして
                ビット 0 で始まるかのように関係マスクを指定します。
                プロセスは、指定した関係マスクと NUMA ノードの間で共通する
                プロセッサ上で実行されるように制限されます。共通するプロセッサ
                がない場合は、プロセスは指定した NUMA ノード上で実行される
                ように制限されます。
    WAIT        アプリケーションを起動し、終了するまで待ちます。
    コマンド/プログラム
                内部コマンドまたはバッチ ファイルの場合、コマンド プロセッサ
                は cmd.exe の /K オプションを使用して実行されます。
                これはコマンドの後でもウィンドウが残ることを意味
                します。

                内部コマンドまたはバッチ ファイルではない場合、そのプログラム
                はウィンドウ モードのアプリケーションまたはコンソール
                アプリケーションとして動作します。

    パラメーター
                コマンド/プログラムに渡すパラメーターです。

注意: SEPARATE および SHARED オプションは 64 ビット プラットフォームでは
サポートされません。
...
```

このように `start` コマンドではウィンドウの表示状態や優先順位等の制御ができる。
たとえば Windows ログオン時に，コマンドプロンプトのウィンドウを最小化した状態で `gpg-agent` を起動させたい場合は

```bat
@echo off
start /min gpg-connect-agent.exe /bye
```

という内容のバッチファイルを作ってスタートアップ・フォルダ[^suf] に放り込んでおけばよい。

[^suf]: ユーザごとのスタートアップフォルダの位置例： `%APPDATA%\Microsoft\Windows\Start Menu\Programs\Startup`

`/wait` オプションを付けるとコマンド終了まで待機してくれる。
`start /wait` と `call` コマンド（`call` も内蔵コマンド）は似ているようで異なった動作をする。
これはバッチ処理をファイル単位で入れ子の構造にするとよく分かる。
簡単な実験をしてみよう。

まず，以下の内容で `alice.cmd` というバッチファイルを作る。

```batch
@echo off
call bob.cmd
echo Hello, Alice!
exit /b
```

`alice.cmd` の中で呼び出している `bob.cmd` の内容は以下の通り。

```bat
@echo off
echo Hello, Bob!
exit /b
```

`alice.cmd` も `bob.cmd` も最終行が `exit /b` になっているのがポイント。
`/b` オプションを付けることでバッチ処理の終了であることを明示し，バッチファイルをモジュール化できる。

これで `alice.cmd` を実行すると

```text
> alice.cmd
Hello, Bob!
Hello, Alice!
```

となった。
`bob.cmd` がサブルーチンのように機能しているのが分かると思う。

このように Windows バッチ処理では `call` と `exit /b` を組み合わせてバッチ処理をモジュール化し再利用可能とすることができる[^call1]。

[^call1]: `call` と `exit /b` の組み合わせは同一ファイル内でも可能。この場合 `call :label` のようにラベルで飛び先を指定する。ラベルをアンカーにしたジャンプなので，どうしても記述がアセンブラっぽくなってしまうのがにんともかんとも。

次に `alice.cmd` の内容を

```bat {hl_lines=[2]}
@echo off
start /b /wait bob.cmd
echo Hello, Alice!
exit /b
```

とし， call コマンドを start コマンドに置き換える。
これを実行すると

```text
> alice.cmd
Hello, Bob!
```

と `bob.cmd` バッチ終了のタイミングでバッチ処理全体が終了してしまう。
`start` コマンドを使うと `cmd.exe /k bob.cmd` 相当のプロセスが走る筈なのだが，子バッチを `/b /wait` オプション付きで呼び出した場合は，親子全体でひとつのバッチ処理として解釈されるのか，子バッチ処理が終了した時点で親バッチ処理もろともプロセスが終了してしまうらしい[^bat1]。
これを回避するための `call` コマンドというわけだ。

[^bat1]: 件の Scrap で教えてもらったが，おそらく DOS 時代のバッチ処理との互換性を維持するためだろう，とのこと。

ちなみに

```bat {hl_lines=[2]}
@echo off
start /b bob.cmd
echo Hello, Alice!
exit /b
```

と `/wait` オプションを外すと，同じコマンドプロンプト・ウィンドウ上で別プロセスとして駆動するため，出力順が不定になる。
たとえば

```text
> alice.cmd
Hello, Alice!
Hello, Bob!
```

てな感じ。

## [NYAGOS] の start

まるっきり余談だが， [NYAGOS] では `start` コマンドは独自内蔵コマンドとして定義されている[^start1]。
いや，うっかり [NYAGOS] 上で `start /?` ってやったら， '?' なんてコマンド知らん（←超意訳）と怒られてしまったのだ（笑）

[^start1]: [NYAGOS] の `start` コマンドは `nyagos.d\start.lua` で定義されている。

あぁ，でも，バッチファイルは [NYAGOS] から起動しても `cmd.exe` 上のプロセスとして走るのでご心配なく。

## ブックマーク

- [/bin/shに慣れた人に贈るバッチファイルの書き方](https://zenn.dev/zetamatta/books/c84cbe23093eee1b5830) : バッチ処理を組む際に参考になる。オススメ
- [GnuPG for Windows : gpg-agent について]({{< ref "/openpgp/using-gnupg-for-windows-2.md" >}})

[NYAGOS]: https://github.com/nyaosorg/nyagos "nyaosorg/nyagos: NYAGOS - The hybrid Commandline Shell betweeeeeeen UNIX & DOS"
