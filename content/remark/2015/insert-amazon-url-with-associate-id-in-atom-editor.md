+++
date = "2015-11-27T22:20:50+09:00"
description = "昔，結城浩さんが Amazon の商品 URL を変換する秀丸マクロを公開されていて， ATOM Editor では使えないためどうしたものかと思っていたのだが，先日 init.coffee に簡単なコマンドを書く方法を習ったので移植してみた。"
draft = false
tags = ["atom", "editor", "tools", "amazon"]
title = "ATOM Editor で Amazon Associate ID を含んだ商品 URL を生成する"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

昔，結城浩さんが Amazon の商品 URL を変換する秀丸マクロを公開されていて

- [アマゾン・アソシエイトID（アフィリエイト用のID）を含んだ商品URLを生成する秀丸マクロ - 結城浩のはてな日記](http://d.hatena.ne.jp/hyuki/20120413/amazon)

これがとても便利で重宝していたのだが， [ATOM] Editor では使えないためどうしたものかと思っていた。
で，先日 [`init.coffee` に簡単なコマンドを書く方法を習った]({{< relref "remark/2015/insert-datetime-in-atom-editor.md" >}})ので，上の秀丸マクロを [ATOM] Editor に移植してみた。

以下が `init.coffee` に追記する内容。

```coffee
# Amazon Associate ID を含んだ商品 URL を生成する
#  クリップボードの内容を読み込み，変換してセットする
#  http(s)://www.amazon.co.jp/... から始まる文字列を想定
#  /dp/XXXXXXXXXX または /ASIN/XXXXXXXXXX のパタンを探す
#  変換できない場合はクリップボードの内容をそのままセットする
# refs http://d.hatena.ne.jp/hyuki/20120413/amazon
amazonUrl = (id) ->
  url = atom.clipboard.read()
  re = /^htt(?:p|ps):\/\/www.amazon.co.jp\//
  if !re.test(url)
    return url
  result = url.match(/\/(?:dp|ASIN)\/(.{10})/)
  if result == null
    return url
  else if result.length < 2
    return url
  asin = result[1]
  if id == ""
    "http://www.amazon.co.jp/exec/obidos/ASIN/#{asin}/"
  else
    "http://www.amazon.co.jp/exec/obidos/ASIN/#{asin}/#{id}/"

insertText = (str) ->
  return unless editor = atom.workspace.getActiveTextEditor()
  selection = editor.getLastSelection()
  selection.insertText(str)

atom.commands.add 'atom-text-editor', 'my-tools:amazon', ->
  id = '' # Amazon Associate ID
  insertText(amazonUrl(id))
```

コードがやっつけでダサいのはご勘弁ということで[^a]。
`insertText` 関数は[前のとき]({{< relref "remark/2015/insert-datetime-in-atom-editor.md" >}})の使い回し。

[^a]: [CoffeeScript](http://coffeescript.org/) は慣れん。

これで `id` に Associate ID （たとえば私の `baldandersinf-22`）をセットして

```
http://www.amazon.co.jp/%E6%95%B0%E5%AD%A6%E3%82%AC%E3%83%BC%E3%83%AB-%E7%B5%90%E5%9F%8E-%E6%B5%A9/dp/4797341378
```

をクリップボードにコピーした状態でコマンドパレットから「My Tools: Amazon」を起動すると

```
http://www.amazon.co.jp/exec/obidos/ASIN/4797341378/baldandersinf-22/
```

と出力される。
変換できない場合はクリップボードの内容をそのまま出力する。

よーし，うむうむ，よーし。

## 参考

- [Atom API](https://atom.io/docs/api/)
- [正規表現 - JavaScript | MDN](https://developer.mozilla.org/ja/docs/Web/JavaScript/Guide/Regular_Expressions)
- [正規表現チェッカー（JavaScript版） | Softel labs](https://www.softel.co.jp/labs/tools/regex/)

[ATOM]: https://atom.io/ "Atom"
