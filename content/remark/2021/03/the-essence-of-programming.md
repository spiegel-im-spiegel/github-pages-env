+++
title = "『再発見の発想法』は非エンジニアこそ読んで欲しい"
date =  "2021-03-14T07:31:13+09:00"
description = "3月14日は「円周率の日」で「数学の日」でアインシュタイン博士の誕生日である。"
image = "/images/attention/kitten.jpg"
tags = [ "book", "engineering" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

3月14日は「円周率の日」で「数学の日」でアインシュタイン博士の誕生日である。
世界の数学者と理学者と工学者は今日というめでたい日を盛大に祝うんだ！

というわけで，今回は[結城浩]さん著作の『[再発見の発想法]』の感想文なんぞを書いてみる。
読書感想文は久しぶりだな（笑）

ソフトウェエア・エンジニアというのは不思議な職業で，それ単体では成立し得ない。
何故ならソフトウェアを使う人の多くはソフトウェエア・エンジニア以外の人たちだからだ。

たとえば{{% ruby "こうば" %}}工場{{% /ruby %}}の工員さんやコールセンターの中の人，役場のお役人様やタクシーや長距離トラックの運ちゃん，そういった人たちのために私達エンジニアは日々頭を悩ませている。

特に悩ましいのは業種や職場によって独自の用語や隠語があることだ。
だから私達はシステムを作る際にはまず「用語集」を作る。

『[再発見の発想法]』6.2章に出てくる「ベンチマーク（benchmark）」という単語なんかが典型的だが，同じ言葉なのに業種によって少しずつ意味が異なることがある。
そうした言葉の背後にある差異に気づかず設計を推し進めると，後でとんでもないしっぺ返しを食らったりするのだ。

でも，最初の段階でお互いに言葉をすり合わせて「用語集」を作っておくと，以後のコミュニケーションがスムーズになる。
特に考え方や手段といった目に見えないものに「名前」を付けることは重要で，名前を付けることで議論のアンカーもしくは原点として機能し始める。
私は名前の「正しさ」というものを全く信用していない不遜な人間だが，当事者がお互いに同じ言葉を同じ意味で話すことの重要性は理解しているつもりである。

とはいえ，私達エンジニアが普段どんな言葉を喋ってどんな思考のもとに設計しコードを組み立てていくか知ってほしいというのも本音だ。

そこで，いよいよ『[再発見の発想法]』の登場である。

この本は[結城浩]さんによる「Software Design」誌上の同名連載をまとめたものだそうで，さまざまな「技術用語」について解説したものだ。
しかし，内容は技術系雑誌に掲載されていたとは思えないほど平易な言葉で書かれている。

用語の意味とその背後にある考え方と日常生活になぞらえた例示が絶妙なバランスで配置されていて，この辺のさじ加減の上手さは流石[結城浩]さんと言うしかない。
やはり《例示は理解の試金石》（by 数学ガール）なんだねぇ。

ソフトウェア・エンジニアリングの有名な格言（？）に「推測するな，計測せよ」というのがある。
『[再発見の発想法]』に出てくる技術用語の選択は，まさに 計測→評価→改修 というエンジニアリングの基本サイクルを意識している印象がある。

技術系雑誌の連載が出自だし「再発見」というタイトルには私達エンジニアが普段使っている言葉を見直すという意味が含まれているのだろうが，私としてはこの本は非エンジニアの方々あるいはエンジニアを目指す若い方々にこそ読んで欲しい。
そして，本に出てくる言葉をエンジニアがつぶやいているときは「こういう風に考えてるんだなぁ」と感じていただければ幸いである。

さて，積ん読状態の「数学ガール」シリーズも読み進めないとなぁ。
ようやく「本を読む」余裕が出てきたよ。

[結城浩]: http://hyuki.com/
[再発見の発想法]: https://www.amazon.co.jp/dp/B08S2LY9VG?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "再発見の発想法 | 結城 浩 | 産業研究 | Kindleストア | Amazon"

## 参考図書

{{% review-paapi "B08S2LY9VG" %}} <!-- 再発見の発想法 -->
{{% review-paapi "B079JLW5YN" %}} <!-- プログラマの数学 第2版 -->
{{% review-paapi "B015SAFU42" %}} <!-- イミテーション・ゲーム -->
<!-- eof -->
