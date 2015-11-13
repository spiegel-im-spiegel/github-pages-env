+++
date = "2015-11-13T10:27:05+09:00"
description = "よろしい，ならば自作しよう。"
draft = true
tags = ["atom", "editor", "tools"]
title = "ATOM Editor で現在日時を挿入する"

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

- [Atomに現在日時を挿入するコマンドを追加する - Qiita](http://qiita.com/toruot/items/b26fde1a898bb52985e1)

上の記事で， [date](https://atom.io/packages/date) パッケージなるものがあるというので早速試してみるが，フォーマットの指定の仕方が分からない。

- [日付を挿入するキーマップをAtomに追加 | Jun Nishii](http://bcl.sci.yamaguchi-u.ac.jp/~jun/ja/blog/150221-insert_date_keymap_to_atom)

この記事を見ると，変な場所を触らないといけないらしい。

よろしい，ならば~~戦争だ~~自作しよう。

「[Atomに現在日時を挿入するコマンドを追加する](http://qiita.com/toruot/items/b26fde1a898bb52985e1)」を参考に， Init Script (`init.coffee`) に以下の記述を追加した。

```coffee
# 現在日時を挿入するコマンドを追加
# refs https://github.com/dannyfritz/atom-date
# refs https://github.com/JerrySievert/date-utils
# refs http://qiita.com/toruot/items/b26fde1a898bb52985e1
daysAbbr = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat']
#daysKanji = ['日', '月', '火', '水', '木', '金', '土']

paddingZero = (str, length) ->
  str = String(str)
  str = '0' + str  while(str.length < length)
  str

timezone = (offset) ->
  if offset == 0
    'Z'
  if offset < 0
    '+' + paddingZero(-offset / 60, 2) + ':00'
  else
    '-' + paddingZero(offset / 60, 2) + ':00'

dateOrTime = (kind) ->
  now = new Date()
  yyyy = now.getYear() + 1900
  mm = paddingZero(now.getMonth() + 1, 2)
  dd = paddingZero(now.getDate(), 2)
  ddd = daysAbbr[now.getDay()]
  #ddd = daysKanji[now.getDay()]
  hh24 = paddingZero(now.getHours(), 2)
  mi = paddingZero(now.getMinutes(), 2)
  ss = paddingZero(now.getSeconds(), 2)
  tz = timezone(now.getTimezoneOffset())
  if kind == 1
    "#{yyyy}/#{mm}/#{dd} (#{ddd})"
  else if kind == 2
    "#{hh24}:#{mi}:#{ss}"
  else if kind == 3
    "#{yyyy}/#{mm}/#{dd} (#{ddd}) #{hh24}:#{mi}:#{ss}"
  else
    "#{yyyy}-#{mm}-#{dd}T#{hh24}:#{mi}:#{ss}#{tz}"

insertText = (str) ->
  return unless editor = atom.workspace.getActiveTextEditor()
  selection = editor.getLastSelection()
  selection.insertText(str)

atom.commands.add 'atom-text-editor', 'my-date:date', ->
  insertText(dateOrTime(1))
atom.commands.add 'atom-text-editor', 'my-date:time', ->
  insertText(dateOrTime(2))
atom.commands.add 'atom-text-editor', 'my-date:datetime', ->
  insertText(dateOrTime(3))
atom.commands.add 'atom-text-editor', 'my-date:rfc3339', ->
  insertText(dateOrTime(0))
```

元のコードでは timezone の項目がなかったので追加し，更に [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) のフォーマットを追加した。
そういや JavaScript の [`Date.getTimezoneOffset()`](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Global_Objects/Date/getTimezoneOffset) って ± が逆になるんだっけ。
忘れてた orz

なるほど，自分で（パッケージにするまでもない）簡単なコマンドを作る場合はこうすればいいんだな。
勉強になりました（谷啓風）。
しかし [CoffeeScript](http://coffeescript.org/) は慣れん。
Arrow functions は確かに便利だけど（ES6 に採用されるくらいだし）。
