+++
title = "NYAGOS 4.3 で GopherLua が採用される"
date = "2018-04-30T18:49:39+09:00"
update = "2018-05-13T13:24:15+09:00"
description = "NYAGOS 4.3 で Lua の Go 言語実装のひとつである GopherLua が採用された。"
image = "/images/attention/tools.png"
tags  = [ "tools", "nyagos", "shell", "windows", "lua", "golang" ]

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

（正式版が出たので内容を更新した 2018-05-07）

[NYAGOS] 4.3 で [Lua] の [Go 言語]実装のひとつである [GopherLua] が採用された。

- [nyagos で lua53.dll のかわりに GopherLua を使おう - Qiita](https://qiita.com/zetamatta/items/112484eb7fdae87830a0)
    - [続・nyagos で lua53.dll のかわりに GopherLua を使おう - Qiita](https://qiita.com/zetamatta/items/18597ed77c4574796c7b)
- [Release 4.3_beta · zetamatta/nyagos](https://github.com/zetamatta/nyagos/releases/tag/4.3_beta)
- [Release 4.3_beta2 · zetamatta/nyagos · GitHub](https://github.com/zetamatta/nyagos/releases/tag/4.3_beta2)
- [Release 4.3.0_0 · zetamatta/nyagos · GitHub](https://github.com/zetamatta/nyagos/releases/tag/4.3.0_0)

以下に変更点をまとめて紹介する。

{{% fig-quote title="Release 4.3.0_0" link="https://github.com/zetamatta/nyagos/releases/tag/4.3.0_0" lang="en" %}}
- **Use Gopher-Lua instead of lua53.dll** #300
    - nyagos with lua53.dll (mains.exe) can be built with `cd mains ; go build`
    - nyagos with no Lua (ngs.exe) can be built with `cd ngs ; go build`
- Made `nyagos.option.cleanup_buffer` (default=false). When it is true, clean up console input buffer before readline.
- `set -o OPTION_NAME` and `set +o OPTION_NAME` (=`nyagos.option.OPTION_NAME=` on Lua)
- Buffer console-output ( go-colorable and bufio.Writer )
^ Fix: Typing C-o looks to raise hang up until Enter or ESCAPE is typed (on 4.3beta) #303
    - Fix the library: [go-box](https://github.com/zetamatta/go-box/commit/322b2318471f1ad3ce99a3531118b7095cdf3842)
- Fix: chcp did not work. (`chcp` was aliaes to update memory of screen width)
- Add `ls -L` which shows information for the file refernces rather than for the link it self. (because ls could not show the directories in OneDrive )

----

- **lua53.dll のかわりに Gopher-Lua を採用** #300
    - 旧来の lua53.dll 版 nyagos (mains.exe) は `cd mains ; go build` でビルド可能
    - Lua無し版 nyagos (ngs.exe) は `cd ngs ; go build` でビルド可能
- `nyagos.option.cleanup_buffer` を追加(デフォルトは false)。true の場合、一行入力の前にコンソールバッファをクリアする
- `set -o OPTION_NAME` と `set +o OPTION_NAME` を新設(`nyagos.option.OPTION_NAME=` on Lua と等価)
- コンソール出力をバッファリングするようにした ( go-colorable and bufio.Writer )
- C-o を押すと Enter か Escape が押されるまでハングしたように見える不具合を修正
    - (ライブラリを修正: [go-box](https://github.com/zetamatta/go-box/commit/322b2318471f1ad3ce99a3531118b7095cdf3842))
- chcp が動作しない不具合を修正 (同コマンドは画面幅取得のため別名定義していた)
- シンボリックリンクの先を参照するオプション `ls -L` を追加（OneDrive内のディレクトリーが表示できなかったため）
{{% /fig-quote %}}

起動時のバージョン表記は以下のようになる。

```text
Nihongo Yet Another GOing Shell 4.3.0_2-amd64 by go1.10.2
Powered by GopherLua 0.1
(c) 2014-2018 NYAOS.ORG <http://www.nyaos.org>
```

私も以前 [Lua] の [Go 言語]実装についてちょろんと調べたことがあるのだが

- [Go 言語による Lua 実装を試してみた]({{< relref "golang/lua.md" >}})

このときは別のパッケージでしかも args などのグローバル変数の設定の仕方がよく分からなくて放置していたのだ（ゴメンペコン）。

[Go 言語]は組込み用途のコードも書けるので，そこに [Lua] エンジンを（[Go 言語]ネイティブで）組み込めるというのは意義が大きいと思う。
例えば現在のツールの多くは設定情報を YAML や TOML などで記述するが，  [Lua] で設定を記述できるのなら，ちょっとしたロジックを組み込むことも可能になる。
実際に [NYAGOS] ではコマンドの alias に [Lua] コードを組み込むことができる。

[NYAGOS] で [GopherLua] を組み込んでくれるのなら，是非ソースコードも読んでノウハウを勉強したい。

{{% div-box %}}
**【追記 2018-05-07】** 4.3.0 にアップグレードして起動した際に

```text
C:\Users\username\AppData\Roaming\NYAOS_ORG\amd64.nyagos.luac line:1(column:1) near '←':   Invalid token
```

と警告が出た。
古いファイルなのかな？

この場合は `C:\Users\username\AppData\Roaming\NYAOS_ORG` フォルダの中身を掃除すればいいようだ。
なお，このフォルダにある `nyagos.history` ファイルは名前の通りコマンド履歴なのでご注意を。
{{% /div-box %}}

## 【追記 2018-05-07】 [NYAGOS] 4.3.0_1 がリリース

- [Release 4.3.0_1 · zetamatta/nyagos · GitHub](https://github.com/zetamatta/nyagos/releases/tag/4.3.0_1)

不具合の修正のみ。

{{% fig-quote title="Release 4.3.0_1" link="https://github.com/zetamatta/nyagos/releases/tag/4.3.0_1" lang="en" %}}
- Fix: nyagos.d/start.lua did not worked because the member `rawargs` of alias-function's argument was not implemented.
- Fix: the return value of alias-function was not evaluted.
- Fix: for the script in -e option, `arg[]` was not assinged.
- Fix: On -f & -e option, warned as `getRegInt: could not find shell in Lua instanc e`
- Fix: batchfile cound not return the value of `exit /b` as ERRORLEVEL

----

- nyagos.d/start.lua が動作していなかった不具合を修正 (エイリアス関数の `rawargs` パラメータが実装されていなかった)
- alias 関数の戻り値が評価されていなかった不具合を修正
- -e オプションのスクリプト向けに、`arg[]` に引数が代入されていなかった
- -e,-f オプションで、`getRegInt: could not find shell in Lua instance` が表示される不具合を修正
- バッチファイルが `exit /b` の値を ERRORLEVEL として返せなかった不具合を修正
{{% /fig-quote %}}

## 【追記 2018-05-07】 [NYAGOS] 4.3.0_2 がリリース

- [Release 4.3.0_2 · zetamatta/nyagos · GitHub](https://github.com/zetamatta/nyagos/releases/tag/4.3.0_2)

不具合の修正のみ。

{{% fig-quote title="Release 4.3.0_2" link="https://github.com/zetamatta/nyagos/releases/tag/4.3.0_2" lang="en" %}}
- #305: Fix issue that user's .nyagos was not loaded again (Thx! @erw7)

----

- #305: ユーザの .nyagos が二回目以降ロードされない不具合を修正(Thx! @erw7)
{{% /fig-quote %}}

## 【追記 2018-05-10】 [NYAGOS] 4.3.0_3 がリリース

- [Release 4.3.0_3 · zetamatta/nyagos](https://github.com/zetamatta/nyagos/releases/tag/4.3.0_3)

不具合の修正。
着々と改善されています。

{{% fig-quote title="Release 4.3.0_3" link="https://github.com/zetamatta/nyagos/releases/tag/4.3.0_3" lang="en" %}}
- Fix: forgot implement nyagos.setalias , nyagos.getalias (`alias { CMD=XXX}` did not work.)
- Fix: that the element [0] of the table value returned by alias-function was not used as the new command name to evaluate.
- Fix: `doc/09-Build_*.md` about how to download sourcefiles from github

----

- nyagos.setalias, nyagos.getalias の実装が漏れており、`alias { CMD=XXX}` が動かなくなっていた
- エイリアスの戻り値でテーブルが与えられた時、コマンド名として解釈すべき、要素[0]が使われていなかった不具合を修正
- `doc/09-Build_*.md`: github からのソースダウンロード方法についてドキュメント更新
{{% /fig-quote %}}

## 【追記 2018-05-13】 [NYAGOS] 4.3.0_4 がリリース

- [Release 4.3.0_4 · zetamatta/nyagos](https://github.com/zetamatta/nyagos/releases/tag/4.3.0_4)

不具合の修正。
着々と改善されています。

{{% fig-quote title="Release 4.3.0_3" link="https://github.com/zetamatta/nyagos/releases/tag/4.3.0_3" lang="en" %}}
- Fix: #309 nyagos.getkey() raised panic (Thx @nocd5)
- Fix: error-message when command `lnk`'s target is not `*.lnk` nor exist.
- Fix: the cursor blink was switched to off on the child process.

----

- Fix: #309 nyagos.getkey() が使えない不具合を修正 (Thx @nocd5)
- `lnk` コマンドの宛先が `*.lnk` でなかったり存在しなかった時のエラーメッセージを修正
- 子プロセスのカーソルがオフになってしまう不具合を修正
{{% /fig-quote %}}

## ブックマーク

- [inforno :: LuaのGo言語実装を公開しました](http://inforno.net/articles/2015/02/15/gopher-lua-released)
- [nyagos 4.3でもmigemoでディレクトリ移動したい! - Qiita](https://qiita.com/nocd5/items/1736064cd9ee652d5920)

[NYAGOS]: https://github.com/zetamatta/nyagos/ "zetamatta/nyagos: NYAGOS - The hybrid UNIXLike Commandline Shell for Windows"
[GopherLua]: https://github.com/yuin/gopher-lua "yuin/gopher-lua: GopherLua: VM and compiler for Lua in Go"
[Lua]: https://www.lua.org/ "The Programming Language Lua"
[Go 言語]: https://golang.org/ "The Go Programming Language"
