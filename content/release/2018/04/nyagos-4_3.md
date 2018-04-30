+++
title = "NYAGOS 4.3 で GopherLua が採用される"
date = "2018-04-30T18:49:39+09:00"
description = "NYAGOS 4.3 で Lua の Go 言語実装のひとつである GopherLua が採用されるようだ。"
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

次に出る [NYAGOS] 4.3 で [Lua] の [Go 言語]実装のひとつである [GopherLua] が採用されるようだ。

- [nyagos で lua53.dll のかわりに GopherLua を使おう - Qiita](https://qiita.com/zetamatta/items/112484eb7fdae87830a0)
- [Release 4.3_beta · zetamatta/nyagos](https://github.com/zetamatta/nyagos/releases/tag/4.3_beta)

[GopherLua] は [Lua] 5.1 がベースになっているのでこれまでのバージョンとの互換性が気になるところではあるが，楽しみでもある。

私も以前 [Lua] の [Go 言語]実装についてちょろんと調べたことがあるのだが

- [Go 言語による Lua 実装を試してみた]({{< relref "golang/lua.md" >}})

このときは別のパッケージでしかも args などのグローバル変数の設定の仕方がよく分からなくて放置していたのだ（ゴメンペコン）。

[Go 言語]は組込み用途のコードも書けるので，そこに [Lua] エンジンを（[Go 言語]ネイティブで）組み込めるというのは意義が大きいと思う。
例えば現在のツールの多くは設定情報を YAML や TOML などで記述するが，  [Lua] で設定を記述できるのなら，ちょっとしたロジックを組み込むことも可能になる。
実際に [NYAGOS] ではコマンドの alias に [Lua] コードを組み込むことができる。

[NYAGOS] で [GopherLua] を組み込んでくれるのなら，是非ソースコードも読んでノウハウを勉強したい。

今後 [NYAGOS] 4.3 の正式版が出たらこの記事を更新・追記する形で紹介します。
あぁ，正式版が楽しみ。

## ブックマーク

- [inforno :: LuaのGo言語実装を公開しました](http://inforno.net/articles/2015/02/15/gopher-lua-released)

[NYAGOS]: https://github.com/zetamatta/nyagos/ "zetamatta/nyagos: NYAGOS - The hybrid UNIXLike Commandline Shell for Windows"
[GopherLua]: https://github.com/yuin/gopher-lua "yuin/gopher-lua: GopherLua: VM and compiler for Lua in Go"
[Lua]: https://www.lua.org/ "The Programming Language Lua"
[Go 言語]: https://golang.org/ "The Go Programming Language"
<!-- eof -->
