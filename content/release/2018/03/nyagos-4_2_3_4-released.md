+++
title = "NYAGOS 4.2.3_4 Released"
date =  "2018-03-05T09:30:02+09:00"
description = "今回は不具合の修正がメインのようだ。"
image = "/images/attention/tools.png"
tags  = [ "tools", "nyagos", "shell" ]
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

[NYAGOS] 4.2.3_4 がリリースされた。

- [Release 4.2.3_4 (bug fixes) · zetamatta/nyagos](https://github.com/zetamatta/nyagos/releases/tag/4.2.3_4)

今回は不具合の修正がメインのようだ。

{{% fig-quote title="Release 4.2.3_4 (bug fixes)" link="https://github.com/zetamatta/nyagos/releases/tag/4.2.3_4" lang="en" %}}
- `ls -?` for help instead of `ls -h`
- Building with go build instead of make.cmd, print version as `snapshot-GOARCH`
- Show an error when `type DIRECTORY` is executed.
- Made error message simple on `del NOTEXISTFILE`
- Fix: #279 Substitution on Environment variable (%VAR:OLD=NEW%) did not ignore case
- Fix: #281 `cd \\server\folder ; open` ->` C:\Windows\system32` was open.
- Fix: #286 A tilde(~) after whitespace enclosed with double quotations was interpreted same as %USERPROFILE%
- #287 On the last entry of the history, do nothing for typing ARROW-DOWN

----

- `ls -h` のかわりに `ls -?` をヘルプに用意した
- make.cmd のかわりに go build でビルドした時、バージョンを `snapshot-GOARCH` と表示するようにした
- `type DIRECTORY` が実行された時にエラーにするようにした。
- `del 存在しないファイル` を実行した時のエラーをシンプルにした.
- #279 環境変数置換(%VAR:OLD=NEW%)で、英大文字/小文字を区別していた不具合を修正
- #281 `cd \\server\folder ; open` で `C:\Windows\system32` 開く不具合を修正
- #286 Fix: 二重引用符内の空白に続く ~ が %USERPROFILE% と解釈されていた不具合を修正
- #287 ヒストリの最後のエントリの時、↓をタイプしても何もしないようにした
{{% /fig-quote %}}

アップデートは上書きコピーでOK。

## ブックマーク

- [ようやく ConEmu と NYAGOS を導入した]({{< relref "remark/2015/conemu-and-nyagos.md">}})
- [NYAGOS で Lua]({{< relref "remark/2015/nyagos-and-lua.md">}})

[NYAGOS]: https://github.com/zetamatta/nyagos/ "zetamatta/nyagos: NYAGOS - The hybrid UNIXLike Commandline Shell for Windows"
