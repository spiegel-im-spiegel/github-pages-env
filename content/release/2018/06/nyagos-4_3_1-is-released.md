+++
title = "NYAGOS 4.3.1 がリリース"
date = "2018-06-03T18:16:56+09:00"
update = "2018-06-24T14:23:49+09:00"
description = "いくつかの起動時オプションが追加された。"
image = "/images/attention/tools.png"
tags  = [ "tools", "nyagos", "shell", "windows" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/profile/"
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

[NYAGOS] 4.3.1 がリリースされた。

- [Release 4.3.1_0 · nyaosorg/nyagos](https://github.com/nyaosorg/nyagos/releases/tag/4.3.1_0)

いくつかの起動時オプションが追加された。

{{% fig-quote type="markdown" title="Release 4.3.1_0" link="https://github.com/nyaosorg/nyagos/releases/tag/4.3.1_0" lang="en" %}}

- Support Windows10's native ESCAPE SEQUENCE processing with `--no-go-colorable` and `--enable-virtual-terminal-processing`
- For #304,#312, added options to search for the executable from the current directory
    - `--look-curdir-first`: do before %PATH% (compatible with CMD.EXE)
    - `--look-curdir-last` : do after %PATH% (compatible with PowerShell)
    - `--look-curdir-never`: never (compatible with UNIX Shells)
- nyagos.prompt can now be assigned string literal as prompt template directly.
- Fix #314 rmdir could not remove junctions.

----

- `--no-go-colorable` と `--enable-virtual-terminal-processing` で、Windows10 ネイティブのエスケープシーケンスをサポート
- #304,#312, カレントディレクトリから実行ファイルを探す時のオプションを追加
    - `--look-curdir-first`: `%PATH%` より前に探す(デフォルト:CMD.EXE互換動作)
    - `--look-curdir-last` : `%PATH%` より後に探す(PowerShell互換動作)
    - `--look-curdir-never`: `%PATH%` だけから実行ファイルを探す(UNIX Shells互換動作)
- nyagos.prompt にプロンプトテンプレートの文字列を直接代入できるようになった。
- #314 rmdir がジャンクションを削除できなかった問題を修正
{{% /fig-quote %}}

{{% div-box type="markdown" %}}
**【追記 2018-06-24】**
下記の件は 4.3.1_3 で改修されたようだ。
`--look-curdir-never` 指定時はカレントディレクトリ直下の `hoge.exe` は `./hoge` でのみ起動する。

{{% div-box type="markdown" %}}
実行ファイルの検索オプションだが `--look-curdir-never` について `--look-curdir-last` との違いがよく分からなかったり。
たとえばカレントディレクトリに `hoge.exe` があって PATH 上に同名のファイルが存在しない場合に，どちらのオプションでも

```text
$ hoge
```

で起動してしまう。
私としては `--look-curdir-never` 指定時は `./hoge` のみで起動することを期待したのだが違うのだろうか。
まぁ，大した問題ではないのでよかろう。
{{% /div-box %}}
{{% /div-box %}}

にしても，コマンドプロンプトと PowerShell ってそんなところで挙動が違うんだなぁ。
いいのか，それ。

## 【追記 2018-06-24】 [NYAGOS] 4.3.1_1 がリリース

不具合等の改修。

{{% fig-quote type="markdown" title="Release 4.3.1_1" link="https://github.com/nyaosorg/nyagos/releases/tag/4.3.1_1" lang="en" %}}
- Remove source code for lua53.dll
- #317: deadlock when `use "subcomplete"` is enabled and rclone.exe is found.
    - See also: [yuin/gopher-lua#181](https://github.com/yuin/gopher-lua/issues/181)
- #318,#319: add compatible functions with lua 5.3
    - bit32.band/bitor/bxor
    - utf8.char/charpattern/codes

----

- lua53.dll 向けのソースコードを削除
- #317: `use subcomplete` が有効で、rclone.exe が見付かった時デッドロックしていた
    - [yuin/gopher-lua#181](https://github.com/yuin/gopher-lua/issues/181) も参照のこと
- #318,#319 下記の Lua 5.3 互換関数を追加
    - bit32.band/bitor/bxor
    - utf8.char/charpattern/codes
{{% /fig-quote %}}

## 【追記 2018-06-24】 [NYAGOS] 4.3.1_2 がリリース

不具合等の改修。

{{% fig-quote type="markdown" title="Release 4.3.1_2" link="https://github.com/nyaosorg/nyagos/releases/tag/4.3.1_2" lang="en" %}}
- #320: fix the imcompatibility: nyagos.rawexec & raweval did not expand tables in arguments.
- --show-version-only enables --norc automatically

----

- #320: nyagos.rawexec & raweval が引数内のテーブルを展開していなかった非互換性を修正
- --show-version-only を指定すると --norc を自動的に有効化するようにした
{{% /fig-quote %}}

## 【追記 2018-06-24】 [NYAGOS] 4.3.1_3 がリリース

不具合等の改修。

{{% fig-quote type="markdown" title="Release 4.3.1_3" link="https://github.com/nyaosorg/nyagos/releases/tag/4.3.1_3" lang="en" %}}
- #316 Fix: zero-length directory-name in %PATH% is regarded as the current directory
- #321 Fix: key function names `previous_history` & `next_history` were not registered.
- Add -h and --help option
- Lines starting with `@` of Lua script are now ignored to embed into batchfile.
- #322 Fix: change the encoding for batchfile's parameters from Thread Codepage to Console Codepage #322
- All of lua variables `nyagos.option.*` are now able to be set by nyagos.exe's command-line option.

----

- #316 %PATH% の中の長さゼロのエントリがカレントディレクトリとみなされていた不具合を修正
- #321 キー機能名の `previous_history` と `next_history` が未登録だった不具合を修正
- -h,--help オプションを追加
- バッチファイル組み込みのため、Luaスクリプトの `@` で始まる行を無視するようにした
- #322 バッチファイルの引数のエンコーディングをスレッドのコードページから、コンソールのコードページへ変更した。
- Lua変数 `nyagos.option.*` の全てを nyagos.exe のコマンドラインオプションで設定できるようにした。
{{% /fig-quote %}}

[NYAGOS]: https://github.com/nyaosorg/nyagos/ "nyaosorg/nyagos: NYAGOS - The hybrid UNIXLike Commandline Shell for Windows"
