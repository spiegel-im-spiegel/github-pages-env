+++
date = "2015-11-02T09:41:51+09:00"
description = "前回，ConEmu とともに NYAGOS を導入したけど，今回は NYAGOS のセッティングを中心に。"
draft = true
tags = ["windows", "tools", "nyagos", "lua"]
title = "NYAGOS で Lua"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

[前回]({{< relref "remark/2015/conemu-and-nyagos.md" >}})， [ConEmu] とともに [NYAGOS] を導入したけど，今回は [NYAGOS] のセッティングを中心に。

[NYAGOS] の特徴のひとつに， [Lua] のエンジンを内蔵し解釈できる点がある。
というか [NYAGOS] 自身が機能拡張を [Lua] で実装していて，ユーザも hackable に [NYAGOS] をカスタマイズできる。

また `nyagos.exe` を `-f` オプションを付けて起動すると [Lua] のソースファイルをスクリプトとして実行できる。
たとえば以下のコード `hello.lua` を以下のように記述し

```lua
print "Hello World!"
```

`nyagos.exe` で実行すると

```
C:>nyagos.exe -f hello.lua
Hello World!
```

となる。

これだけだと大したことはないが， `nyagos.exe` でスクリプトを実行した場合は [NYAGOS] の [Lua] 拡張が使える。
例えばカレント・フォルダを取得する `getwd.lua` を以下のように記述し

```lua
nyagos.write(nyagos.getwd().."\n")
```

`nyagos.exe` で実行すると

```
C:\Users\username>nyagos.exe -f getwd.lua
C:\Users\username
```

となる。

[Lua] 拡張については [NYAGOS のドキュメント](https://github.com/zetamatta/nyagos/blob/master/Doc/nyagos_ja.md)を参照するとよいだろう。
特筆すべきは `alias` 機能で，単純なコマンドの置き換えだけでなく

```
nyagos.alias.cmdname = function(args) ... end
```

のような形式で新しい内部コマンドを作成することもできる。

ところで `nyagos.exe` を普通に起動して `ls -oF` コマンドでファイルの一覧を表示させてみると，拡張子 `.lua` のファイルは実行可能ファイルになっていることが分かる。
実際に起動しようとすると

```
~> hello.lua
'lua' is not recognized as an internal or external command,
operable program or batch file
```

と怒られる。
どうやら `lua` コマンドがねーよ，と言っているらしい（確かに [Lua] の実行モジュールは入れてないのだが）。

[NYAGOS] をインストールしたフォルダにある `nyagos.d\suffix.lua` ファイルの末尾あたりを見ると

```lua
suffix.pl="perl"
if nyagos.which("ipy") then
  suffix.py="ipy"
elseif nyagos.which("py") then
  suffix.py="py"
else
  suffix.py="python"
end
suffix.rb="ruby"
suffix.lua="lua"
suffix.awk={"awk","-f"}
suffix.js={"cscript","//nologo"}
suffix.vbs={"cscript","//nologo"}
suffix.ps1={"powershell","-file"}
```

となっていて，実際に拡張子 `.lua` のファイルが `lua` コマンドに関連付けられているのが分かる。
ってことは，これを `nyagos.exe` に書き換えればいいわけだ。
とはいえ `nyagos.d\suffix.lua` ファイルを直接いじるわけにはいかないので（バージョンアップのたびに上書きされる），
 `%HOME%` または `%HOMEPATH%` フォルダにある `.nyagos` ファイルに以下の記述を追加する[^a]。

 [^a]: `suffix` ではなく `alias` で `lua` コマンドを定義する手もある。
 
 ```lua
 suffix.lua={"nyagos.exe","-f"}
 ```

これで `nyagos.exe` を起動し直して `hello.lua` ファイルを実行してみる。

```
~> hello.lua
Hello World!
```

おおっ，うまくいった。
これで [Lua] スクリプトを [NYAGOS] 上でバッチファイルのように扱うことができる。

さぁて，これでいよいよもって [Lua] の勉強をしないといけなくなったなぁ。
Pascal 系の構文は目が滑るんだよなぁ。

- [お気楽 Lua プログラミング超入門](http://www.geocities.jp/m_hiroi/light/lua01.html)
- [Luaプログラミング入門 | densan-labs.net](http://densan-labs.net/tech/lua/index.html)

[ConEmu]: https://conemu.github.io/ "ConEmu - Handy Windows Terminal"
[NYAGOS]: http://www.nyaos.org/index.cgi?p=NYAGOS "NYAOS.ORG - NYAGOS"
[Lua]: http://www.lua.org/ "The Programming Language Lua"
