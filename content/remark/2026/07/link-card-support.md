+++
title = "リンクカードに対応した"
date =  "2026-07-08T20:36:01+09:00"
description = "久しぶりに（AI を使わず）自力でコードを書いた。たまには脳みそを使わないと（笑）"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "site", "tools", "web", "hugo" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

取引量が少なくて [Amazon PA-API にアクセスできなくなってしまった]({{< ref "/remark/2025/11/paapi-access-disabled.md" >}} "Amazon PA-API にアクセスできなくなった")ので，今までの商品レビューカードから Web サイトのリンクカードに切り替えないと。
...と思いながら半年以上経ってしまった（その間に骨折入院とかもあったけど）。

ようやくリンクカードを Hugo の処理で出せるようになった。
こんな感じ。

{{% linkcard "79c9ba36c62cd792801ef499cffe7e614e89670a" %}} <!-- https://text.baldanders.info/ -->

Web ページの情報収集のための[ツール][linkcard]を作るのに時間がかかってしまったのだ。
いや，厳密には，実作業時間では半日ほどなのだが，やる気スイッチがなかなか入らなくて（笑）

あらかじめ[ツール][linkcard]を使って Web ページ情報をこんな感じに Hugo の環境下に保持っておいて

```json
[
  {
    "hash_id": "79c9ba36c62cd792801ef499cffe7e614e89670a",
    "url": "https://text.baldanders.info/",
    "title": "text.Baldanders.info",
    "description": "帰ってきた「しっぽのさきっちょ」",
    "image_url": "https://text.baldanders.info/images/attention/site.jpg",
    "image_path": "/images/thumbs/79c9ba36c62cd792801ef499cffe7e614e89670a.jpg",
    "image_width": 100,
    "stars": [
      false,
      false,
      false,
      false,
      false
    ]
 }
]
```

実際に使うときは Hugo のショートコードで展開してる。
上のリンクカードだとこんな感じ。

```hugo
{{</* linkcard "79c9ba36c62cd792801ef499cffe7e614e89670a" */>}}
```

引数で `hash_id` を指定すれば，そのリンクカードを展開するようにしている。

最近は出版社さんもページの `og:image` 等で書影を指示してくれるので，こちらは `og:image` の中身をサムネイル表示すれば，本の紹介用リンクカードになる。
たとえば[最近読んだ『ケアレス・ピープル』]({{< ref "/remark/2026/07/skimming-through-careless-people.md" >}} "『ケアレス・ピープル』を斜め読み")ならこんな感じ。

{{% linkcard "955ecd76b8158e21c4262c73717f0c1bc2d94351" %}} <!-- ケアレス・ピープル -->

ありがたい話である。

YouTube とかも埋め込みにしないで，こんな感じにリンクカードにすれば煩くないかな。

{{% linkcard "ffea61f5e3dfd0138b8ce8d45abc1acb9e50f134" %}} <!-- 輪堂千速生誕祭2026 -->

今回，[ツール][linkcard]を組むのに意図的に AI エージェントを使わないようにした。
いや，細かく指示するのが面倒になって「自分で書いたほうが速い」となってしまったのだ。
たまには自力でコードを書かないと脳みそが鈍るからな（笑）

あっ，でも，レビューとかは AI エージェントにやってもらってる。
ひとりでコードを書くときって AI が QA 役みたいに振る舞ってくれるので，めっちゃ助かっている。
AI がコードを書くときは私が QA 役かな。

[linkcard]: "https://github.com/spiegel-im-spiegel/linkcard" "spiegel-im-spiegel/linkcard: Information for creating link cards"
