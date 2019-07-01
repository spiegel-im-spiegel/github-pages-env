+++
date = "2015-11-02T21:41:54+09:00"
update = "2018-03-27T09:22:45+09:00"
description = "前回，ConEmu とともに NYAGOS を導入したけど，今回は NYAGOS のセッティングを中心に。"
draft = false
tags = ["windows", "tools", "nyagos", "shell", "lua"]
title = "NYAGOS で Lua"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "https://baldanders.info/profile/"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[前回]({{< ref "/remark/2015/conemu-and-nyagos.md" >}})， [ConEmu] とともに [NYAGOS] を導入したけど，今回は [NYAGOS] のセッティングを中心に。

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

[Lua] 拡張については[ドキュメント](https://github.com/zetamatta/nyagos/blob/master/Doc/07-LuaFunctions_ja.md)を参照するとよいだろう。
特筆すべきは `alias` 機能で，単純なコマンドの置き換えだけでなく

```
nyagos.alias.cmdname = function(args) ... end
```

のような形式で新しい内部コマンドを作成することもできる。

ところで `nyagos.exe` を普通に起動してプロンプトから `ls -oF` コマンドでファイルの一覧を表示させてみると，拡張子 `.lua` のファイルは実行可能ファイルになっていることが分かる[^rc]。

[^rc]: 拡張子 `.lua` とコマンドの関連付けは [NYAGOS] をインストールしたフォルダにある `_nyagos` ファイルに記述されている。

<!--
ただ，実際に起動しようとすると

```
~> hello.lua
'lua' is not recognized as an internal or external command,
operable program or batch file
```

と怒られた。
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

となっていて，拡張子 `.lua` のファイルが `lua` コマンドに関連付けられているのが分かる。
ってことは，これを `nyagos.exe` に書き換えればいいわけだ。
とはいえ `nyagos.d\suffix.lua` ファイルを直接いじるわけにはいかないので（バージョンアップのたびに上書きされる）， `%HOME%` または `%HOMEPATH%` フォルダにある `.nyagos` ファイルに以下の記述を追加する[^a]。

 [^a]: `suffix` ではなく `alias` で `lua` コマンドを定義する手もある。

 ```lua
 suffix.lua={"nyagos.exe","-f"}
 ```

これで `nyagos.exe` を起動し直して `hello.lua` ファイルを実行してみる。
-->

実際に [NYAGOS] のプロンプト上で動かしてみよう。

```
$ hello.lua
Hello World!
```

おおっ，うまくいった。
これで [Lua] スクリプトを [NYAGOS] 上でバッチファイルのように扱うことができる。

さぁて，これでいよいよもって [Lua] の勉強をしないといけなくなった。
Pascal 系の構文は目が滑るんだよなぁ。

## ブックマーク

### NYAGOS 作者による解説

- [環境変数PATHが長すぎて、追加したパスが有効にならないぜ！ - Qiita](http://qiita.com/zetamatta/items/a49e3a40201511128508)
- [nyagosスクリプト解説 - svn のサブコマンドを勝手に拡張する - Qiita](http://qiita.com/zetamatta/items/c4ad3cc55c5afa74da63)
- [nyagosスクリプト解説 - ezoe.lua「コマンドではない。」 - Qiita](http://qiita.com/zetamatta/items/29a85695813926cafd2c)
- [nyagosスクリプト解説 - 逆クォートによるコマンド出力展開編(backquote.lua) - Qiita](http://qiita.com/zetamatta/items/cdff310f53faf3369e48)
- [nyagosスクリプト解説 - CMD.EXEで化けさせず、nyagosの中だけプロンプトをカラー化 - Qiita](http://qiita.com/zetamatta/items/c08586c85fa73c182a7a)
- [nyagosスクリプト解説 - VisualStudio れんけー - Qiita](http://qiita.com/zetamatta/items/89a907f4bd46d1750c31)
- [git 付属のPerlやunzipとかを使いたいけどsort,find,lnとかは要らない - Qiita](http://qiita.com/zetamatta/items/1fe83f736b0254e02415) [^b]

[^b]: 最近の [Git for Windows](https://git-for-windows.github.io/) は bash 関連のコマンドを `Git\usr\bin` フォルダに集めている。 git コマンドだけが必要なら `Git\cmd` フォルダにのみパスを通せばよい。 bash を使う場合は `Git\bin` フォルダにある `bash.exe` を起動するのが一番安全なようだ。

### Lua の解説

- [Lua の Windows へのインストールと使い方 | プログラマーズ雑記帳](http://yohshiy.blog.fc2.com/blog-entry-291.html)
- [Luaプログラミング入門 | densan-labs.net](http://densan-labs.net/tech/lua/index.html)
- [高速スクリプト言語 Lua を始めよう — WTOPIA v1.0 documentation](http://www.ie.u-ryukyu.ac.jp/~e085739/lua.hajime.html)
- [良いもの。悪いもの。: Lua基礎文法最速マスター](http://handasse.blogspot.com/2010/02/lua.html)
- [紀子さん＠へぼぷろぐらまの日常 | Luaで日付時間操作。](http://noriko3.blog42.fc2.com/blog-entry-128.html)
- [Luaでファイルの読み書きを行なう - Symfoware](http://symfoware.blog68.fc2.com/blog-entry-454.html)
- [Lua のコルーチンの使い方〜基本編〜 : torus solutions!](http://torus.jp/memo/x200907/lua-coroutine.rd.html)

### その他

- [NYAGOSとconemuでキーボード操作の拡張 - Qiita](http://qiita.com/daxanya1/items/7d4b51bba6c8f3a6016b) : [NYAGOS]+[Lua] で動作を定義， [ConEmu] でキーバインドを変更する
- [nyagosでbower searchを便利につかいたかった - Qiita](http://qiita.com/JugnautOnishi/items/7bec6008b6bdb1c1fb9a)
- [【ポエム】 NYAGOSの現在と今後 【2017年版】 - Qiita](http://qiita.com/zetamatta/items/3e83c7bfdfbe7fcc92b5)
    - [【ポエム】NYAGOS 2017年ふりかえり - Qiita](https://qiita.com/zetamatta/items/e2ae6e2ca232a3164214)
- [一見、NYAGOSの不具合に見える事例について（随時追記） - Qiita](https://qiita.com/zetamatta/items/441ff50da7c8f3338260)
    - [NYAGOS で環境設定系バッチファイルが期待どおり動かない原因と対策 - Qiita](https://qiita.com/zetamatta/items/f62bafd711755a4cf8d7)
    - [NYAGOS 4.2.5βが、いかにしてバッチファイルでの環境変数の変更取り込みを可能としたか - Qiita](https://qiita.com/zetamatta/items/efff93d92ac2150192fb)
- [補完候補の既入力部分の文字列をハイライト - Qiita](https://qiita.com/nocd5/items/a5e136285804ba2d02c3)
- [Nyagos リポジトリブランチ名表示 プロンプトの一例 - Qiita](https://qiita.com/tsuyoshi_cho/items/d029825b6d8d3688da92)

[ConEmu]: https://conemu.github.io/ "ConEmu - Handy Windows Terminal"
[NYAGOS]: http://www.nyaos.org/index.cgi?p=NYAGOS "NYAOS.ORG - NYAGOS"
[Lua]: http://www.lua.org/ "The Programming Language Lua"
