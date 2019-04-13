+++
title = "NYAGOS 4.2.4_0 Released"
date = "2018-03-11T15:28:00+09:00"
update = "2018-03-27T13:50:56+09:00"
description = "細かい機能の変更や修正がメイン。"
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

[NYAGOS] 4.2.4_0 がリリースされた。

- [Release 4.2.4_0 · zetamatta/nyagos](https://github.com/zetamatta/nyagos/releases/tag/4.2.4_0)

細かい機能の変更や修正がメイン。

{{% fig-quote type="md" title="Release 4.2.4_0" link="https://github.com/zetamatta/nyagos/releases/tag/4.2.4_0" lang="en" %}}
- lua: ole: `variable = OLE.property` is avaliable instead of `OLE:_get('property')`
- lua: ole: `OLE.property = value` is avaliable instead of `OLE:_set('property',value)`
- Load `nyagos.d/*.ny` as batchlike file
- #266: `lua_e "nyagos.option.noclobber = true"` forbides overwriting existing file by redirect.
- #269: `>| FILENAME` and `>! FILENAME` enable to overwrite the file already existing by redirect even if `nyagos.option.noclobber = true`
- #270: Console input buffer has been cleaned up when prompt is drawn.
- #228: Completion supports $ENV[TAB]... by native
- #275: Fix: history substitution like `!str:$` , `!?str?:$` did not work.
- The error `event not found` is caused when the event pointed !y does note exists.
- #285: Not wait GUI-process not using pipeline terminating like CMD.EXE (Call them with ShellExecute() instead of CreateProcess() )
- (Replaced `bytes.Buffer` to `strings.Builder` and Go 1.10 is required to build)
- When more than one are to be executed with `open` at once, display error: `open: ambiguous shellexecute`
- Fix that `nyagos.alias.NAME = nil` could not remove the alias.

----

- lua: ole: `variable = OLE.property` が `OLE:_get('property')` のかわりに使えるようになった
- lua: ole: `OLE.property = value` が `OLE:_set('property',value)` のかわりに使えるようになった
- `nyagos.d/*.ny` のコマンドファイルも読み込むようにした
- #266: `lua_e "nyagos.option.noclobber = true"` でリダイレクトでのファイル上書きを禁止
- #269: `>| FILENAME` もしくは `>! FILENAME` で、`nyagos.option.noclobber = true` の時も上書きできるようにした
- #270: プロンプト表示時にコンソール入力バッファをクリアするようにした
- #228: $ENV[TAB] という補完をネイティブでサポート
- #275: !str:$ や !str?:$ といったヒストリ置換が機能しない不具合を修正
- ! で指定されるヒストリが存在しない時「event not found」エラーを出させるようにした
- #285: パイプラインを使っていない GUIプログラムは CMD.EXE 同様終了を待たないようにした (CreateProcess ではなく ShellExecute を使用する)
- (`bytes.Buffer` を `strings.Builder` に置き換えた。Go 1.10 が必要になった)
- 複数のファイルが「open」で一度に開こうとした時、`open: ambiguous shellexecute` とエラーを表示するようにした。
- `nyagos.alias.NAME = nil` で alias を削除できなかった動作を修正
{{% /fig-quote %}}

`nyagos.d/*.ny` は今回のリリース・パッケージには含まれてなかったのだけど，次回以降で入ってくるってことかなぁ？ それとも自前のスクリプトは `*.lua` じゃなくて `*.ny` にしろってことなのだろうか。

（**追記：** `*.ny` は単純にコマンドを列挙したバッチ処理を格納できるファイルだそうだ）

アップデートは上書きコピーでOK。

## ブックマーク

- [一見、NYAGOSの不具合に見える事例について（随時追記） - Qiita](https://qiita.com/zetamatta/items/441ff50da7c8f3338260)

- [ようやく ConEmu と NYAGOS を導入した]({{< ref "/remark/2015/conemu-and-nyagos.md">}})
- [NYAGOS で Lua]({{< ref "/remark/2015/nyagos-and-lua.md">}})

[NYAGOS]: https://github.com/zetamatta/nyagos/ "zetamatta/nyagos: NYAGOS - The hybrid UNIXLike Commandline Shell for Windows"
