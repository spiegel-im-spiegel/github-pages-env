+++
title = "Amazon アフィリエイトリンク作成サービスを Amakuri へ移行する"
date = "2018-10-19T17:19:05+09:00"
description = "G-Tools が 2018-10-13 付けでサービスが終わってしまったので Amakuri へ移行することにした。"
image = "/images/attention/kitten.jpg"
tags = [ "amazon", "web", "site", "market" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

このサイトで本を紹介する際に「[Amazon アソシエイト・プログラム](https://affiliate.amazon.co.jp/)」を利用しているのだが[^aa1]，アフィリエイトリンクを作成するのに今までは [G-Tools] を利用してきた。
ところが 2018-10-13 付けでサービスが終わってしまったらしい[^aa2]。

[^aa1]: まぁ金額は微々たるものだけど，気軽に書影を利用できるので重宝している。
[^aa2]: [G-Tools] のサイトには「Amazonの規約変更により、G-Toolsの機能提供を継続することが難しくなったため」と書かれていた。

んで，何か代わりになるサービスはないかなぁ，と探してみたら [Amakuri] という Web サービスが良さげである。

- [Amakuri [Amazonアフィリエイトリンク作成ツール]](https://dadadadone.com/amakuri/)

{{< div-box type="markdown" >}}
**【2021-05-04 追記】**
Amakuri は 2019-01-15 にサービスを閉鎖しています。
この記事の内容はコメントアウトしています。
悪しからずご了承の程を。
{{< /div-box >}}

<!--
基本的な使い方は [G-Tools] と似ていて，商品名, ASIN コード, ISBN 番号などをキーに検索してページに貼り付ける HTML コードを取得するというもの。
リンクに Amazon Associate ID を設定できる。

このサービスが秀逸なのは，出力する HTML コードをカスタマイズできること。
「カスタムテンプレート」と呼ばれるものを最大8つまで登録し（Cookie を使うのかな），任意の出力を得ることができる。

まずは商品を検索する。
今回は『[Go言語による並行処理]』というタイトルで本を探してみる。
検索結果はこんな感じ。

{{< fig-img src="https://photo.baldanders.info/flickr/image/43605887110_m.png" title="Search by Amakuri" link="https://photo.baldanders.info/flickr/43605887110/" >}}

検索結果の「商品リンクを作る」ボタンを押下（クリック）すると以下の画面になる。

{{< fig-img src="https://photo.baldanders.info/flickr/image/45371115622_m.png" title="Amakuri" link="https://photo.baldanders.info/flickr/45371115622/" >}}

Associate ID を持っているなら「アソシエイトID」の項目に入力して「登録」すれば「貼り付け用リンクコード」に出力結果が即時反映される。
なくても無問題。

「テンプレート」の項目に「カスタムテンプレート1」を選択し「カスタムテンプレート1」の項目にテンプレート・コードを入力する。
たとえばこんな感じのコード。

```html
<div class="hreview">
  <div class="photo"><a class="item url" href="<%link_url%>"><%image_medium%></a></div>
  <dl class="fn">
    <dt><a href="<%link_url%>"><%title%></a></dt>
    <dd><%author%></dd>
    <dd><%label%></dd>
  </dl>
  <p class="description"><%comment%></p>
  <p class="powered-by" >powered by <a href="<%amakuri_url%>" >Amakuri</a></p>
</div>
```

入力したら「カスタムテンプレート適用」ボタンを推すと「貼り付け用リンクコード」に出力結果が反映される。
上述のテンプレート・コードに対する出力結果はこんな感じ。

```html
<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4873118468/"><img src="https://images-fe.ssl-images-amazon.com/images/I/51pUKQajnaL._SL160_.jpg" width="125" height="160" alt="Go言語による並行処理"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4873118468/">Go言語による並行処理</a></dt>
    <dd>Katherine Cox-Buday</dd>
    <dd>オライリージャパン</dd>
  </dl>
  <p class="description"></p>
  <p class="powered-by">powered by <a href="https://dadadadone.com/amakuri/" >Amakuri</a></p>
</div>
```

実際の表示はこんな感じになる[^css1]。

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4873118468/baldandersinf-22"><img src="https://images-fe.ssl-images-amazon.com/images/I/51pUKQajnaL._SL160_.jpg" width="125" height="160" alt="Go言語による並行処理"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4873118468/baldandersinf-22">Go言語による並行処理</a></dt>
    <dd>Katherine Cox-Buday</dd>
    <dd>オライリージャパン</dd>
  </dl>
  <p class="description"></p>
  <p class="powered-by">powered by <a href="https://dadadadone.com/amakuri/" >Amakuri</a></p>
</div>

[^css1]: スタイルの指定は [`addon.css`](/css/addon.css) を参考にどうぞ。

テンプレート・コード中の `<%title%>` などのタグが実際の商品データに置き換わる。
テンプレート用の主なタグを以下に挙げる。

| タグ                                   | 内容                                                     |
| -------------------------------------- | -------------------------------------------------------- |
| `<%title%>`                            | 商品名                                                   |
| `<%author%>`                           | 著者名                                                   |
| `<%label%>`                            | 出版社名                                                 |
| `<%ranking%>`                          | 売上ランキング順位                                       |
| `<%ranking_tag%> ... </%ranking_tag%>` | （売上ランキング情報が存在する場合のみ要素内を展開する） |
| `<%platform%>`                         | プラットフォーム名                                       |
| `<%image_small%>`                      | 商品画像(小) : img タグに展開                            |
| `<%image_medium%>`                     | 商品画像(中) : img タグに展開                            |
| `<%image_large%>`                      | 商品画像(大) : img タグに展開                            |
| `<%comment%>`                          | コメント（任意文字列）                                   |
| `<%link_text%>`                        | （固定文字列） Amazon への誘導リンクテキスト             |
| `<%link_url%>`                         | Amazon へのリンクURL                                     |
| `<%amakuri%>`                          | （固定文字列） [Amakuri] へのクレジット表記              |
| `<%amakuri_url%>`                      | （固定文字列） [Amakuri] への                            |
| `<%posted_date%>`                      | 作成日（例: 2014.8.28）                                  |
| `<%hanbai_price%>`                     | 販売価格                                                 |
| `<%hanbai_tag%> ... <%/hanbai_tag%>`   | （販売価格情報が存在する場合のみ要素内を展開する）       |

出版日のタグがないのがなぁ。
せめて年だけでも...

まぁ，こんな感じで運用していくことにしよう。
-->

## ブックマーク

- [AmazonアフィリエイトツールAmakuriの使い方とCSSデザインサンプル](https://naifix.com/amakuri/)
- [hreview-examples · Microformats Wiki](http://microformats.org/wiki/hreview-examples)

[G-Tools]: https://www.goodpic.com/mt/aws/ "G-Tools ブログとAmazon(アマゾン) アソシエイトでアフィリエイト"
[Amakuri]: https://dadadadone.com/amakuri/ "Amakuri [Amazonアフィリエイトリンク作成ツール]"
[Go言語による並行処理]: https://www.amazon.co.jp/exec/obidos/ASIN/4873118468/baldandersinf-22/ "Go言語による並行処理 | Katherine Cox-Buday, 山口 能迪 |本 | 通販 | Amazon"
