+++
title = "CAVOC Web API で遊ぶ"
date =  "2021-05-29T16:58:27+09:00"
description = "CVO 情報を取得する Web API があるみたいなので，これを検証するためのコマンドライン・ツールを作ってみた。"
image = "/images/attention/kitten.jpg"
tags = [ "engineering" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

突然だが「[共通農業語彙（Common Agricaulturel VOcabulary; CAVOC）](http://cavoc.org/)」というサイトがあるのだが，この中に農作物の語彙体系を整理したデータベースがある。

- [農作物語彙体系(CVO, Crop VOcabulary)](http://cavoc.org/cvo.php)

この CVO 情報を取得する Web API があるみたいなので，これを検証するためのコマンドライン・ツールを作ってみた。

- [spiegel-im-spiegel/gcavoc: Common Agricaulturel Vocabulary API Client by Golang](https://github.com/spiegel-im-spiegel/gcavoc)

動作例を挙げると「せろり」の標準名は

```text
$ gcavoc std せろり
{"term":"セロリ"}
```

てな感じに取得できる。
また，セロリの Wikipedia ページを取得したいなら

```text
$ gcavoc wikipedia セロリ
https://ja.wikipedia.org/wiki/%E3%82%BB%E3%83%AB%E3%83%AA%E3%82%A2%E3%83%83%E3%82%AF
```

という感じ。
いや，なんで「[セルリアック](https://ja.wikipedia.org/wiki/%E3%82%BB%E3%83%AB%E3%83%AA%E3%82%A2%E3%83%83%E3%82%AF)」のページなんだ？

とまぁ微妙なんだよな。

たとえば，同じ「セロリ」にしても収穫・出荷する部位によって異なる「農作物」になるのだが Web API ではそこまで細かく制御できない感じ。
元データはオープンデータで公開されているので，真面目にやりたいなら自前で何とかしろってことかもしれない。
