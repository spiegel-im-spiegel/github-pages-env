# [Amakuri] 用テンプレート

## テンプレート・コード（HTML）

```html
<div class="hreview">
  <div class="photo"><a class="item url" href="<%link_url%>"><%image_medium%></a></div>
  <dl class="fn">
    <dt><a href="<%link_url%>"><%title%></a></dt>
    <dd><%author%></dd>
    <dd><%label%></dd>
    <dd>評価&nbsp;<abbr class="rating fa-sm" title="3">
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="far fa-star"></i>
      <i class="far fa-star"></i>
    </abbr></dd>
  </dl>
  <p class="description"><%comment%></p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed"><%posted_date%></abbr> (powered by <a href="<%amakuri_url%>" >Amakuri</a>)</p>
</div>
```

## 主なテンプレート文字列

| タグ                                 | 内容                                                                                               |
| ------------------------------------ | -------------------------------------------------------------------------------------------------- |
| <%title%>                            | 商品名                                                                                             |
| <%author%>                           | 著者名                                                                                             |
| <%label%>                            | 出版社                                                                                             |
| <%ranking%>                          | 売上ランキング順位                                                                                 |
| <%ranking_tag%> ... </%ranking_tag%> | 例: `<%ranking_tag%>売上げランキング： <%ranking%></%ranking_tag%>`                                |
| <%platform%>                         | プラットフォーム                                                                                   |
| <%image_small%>                      | 商品画像(小) : img タグ                                                                            |
| <%image_medium%>                     | 商品画像(中) : img タグ                                                                            |
| <%image_large%>                      | 商品画像(大) : img タグ                                                                            |
| <%comment%>                          | コメント（任意文字列）                                                                             |
| <%link_text%>                        | Amazon への誘導リンクテキスト（固定文字列）                                                        |
| <%link_url%>                         | Amazon へのリンクURL                                                                               |
| <%amakuri%>                          | `Posted with <a href="https://dadadadone.com/amakuri/" target="_blank">Amakuri</a>` （固定文字列） |
| <%amakuri_url%>                      | `https://dadadadone.com/amakuri/` （固定文字列）                                                   |
| <%posted_date%>                      | 作成日（例: 2014.8.28）                                                                            |
| <%hanbai_price%>                     | 販売価格                                                                                           |
| <%hanbai_tag%> ... <%/hanbai_tag%>   | 例: `<%hanbai_tag%>販売価格 <%hanbai_price%><%/hanbai_tag%>`                                       |

## ブックマーク

- [hreview-examples · Microformats Wiki](http://microformats.org/wiki/hreview-examples)

[Amakuri]: https://dadadadone.com/amakuri/ "Amakuri [Amazonアフィリエイトリンク作成ツール]"
