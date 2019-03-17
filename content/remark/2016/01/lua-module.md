+++
date = "2016-01-04T22:44:13+09:00"
description = "Lua は不案内なので知らなかったのだが module() 関数は Lua 5.2 で廃止されていたらしい。"
draft = false
tags = ["lua", "module", "nyagos"]
title = "Lua のモジュール"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[NYAGOS] 4.1.0_0 がリリースされた。
主な変更点は `ln` コマンドが追加されたことのようだ。

- [Release 4.1.0_0 · zetamatta/nyagos](https://github.com/zetamatta/nyagos/releases/tag/4.1.0_0)

ただ，私の場合はより切実な問題があって， `-f` オプションで [Lua] のスクリプトを実行させた場合に `module()` 関数が使えなくなった。
[Lua] は不案内なので知らなかったのだが `module()` 関数は [Lua] 5.2 で deprecated になっていたらしい。
逆になんで今まで使えてたのかは分からない。

- [Luaのモジュールを穴が空くまで見つめてみる - プログラミングの魔物](http://p-monster.hatenablog.com/entry/2013/02/13/205419)

`module()` 関数が使えないので `require()` で外部ファイルを呼び出すとファイル内の記述がそのまま実行される。

以前なら `module1.lua` に

```lua
module("module1", package.seeall)

function method1()
    return "Method 1"
end

function method2()
    return "Method 2"
end
```

と定義しておけば

```lua
require("module1")

nyagos.write(module1.method1().."\n")
nyagos.write(module1.method2().."\n")
```

と記述できた。
もし同じように機能させたいなら `module1.lua` を

```lua
module1 = {}

module1.method1 = function()
    return "Method 1"
end

module1.method2 = function()
    return "Method 2"
end
```

と記述するのが一番簡単なようだ。
`module1` を関数テーブルとして定義するわけだ。

あるいは `module1.lua` を

```lua
local module1 = {}

module1.method1 = function()
    return "Method 1"
end

module1.method2 = function()
    return "Method 2"
end

return module1
```

としておいて，呼び出し側を

```lua
local module1 = require("module1")

nyagos.write(module1.method1().."\n")
nyagos.write(module1.method2().."\n")
```

とすればグローバル領域を汚さずに済むだろう。

- [Lua Unofficial FAQ (uFAQ)](http://www.luafaq.org/) : “1.37.2 Life after module()?” の項が参考になる
- [c++ - Using the 'module' function in Lua 5.2? - Stack Overflow](http://stackoverflow.com/questions/16849422/using-the-module-function-in-lua-5-2)
- [その４ 会得必須！Luaの真髄「テーブル」](http://marupeke296.com/LUA_No4_Table.html)
- [第 4 回: Lua のオブジェクト指向について紹介する — WTOPIA v1.0 documentation](http://www.ie.u-ryukyu.ac.jp/~e085739/lua.hajime.4.html)
- [Luaのモジュール徹底解説（Lua 5.1〜5.3対応） - Qiita](https://qiita.com/mod_poppo/items/ef3d8a6fe03f7f426426)

[NYAGOS]: http://www.nyaos.org/index.cgi?p=NYAGOS "NYAOS.ORG - NYAGOS"
[Lua]: http://www.lua.org/ "The Programming Language Lua"
