+++
title = "PA-API v5 クライアント・パッケージ v0.5.0 をリリースした"
date =  "2019-10-20T20:29:51+09:00"
description = "コメントとかドキュメントはまだまだ整理中なので，長い目で見てやってください。"
image = "/images/attention/go-logo_blue.png"
tags = [ "package", "module", "golang", "amazon", "pa-api", "programming" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go 言語]用 PA-API v5 クライアント・パッケージ [spiegel-im-spiegel/pa-api] v0.5.0 をリリースした。

- [Release v0.5.0 · spiegel-im-spiegel/pa-api · GitHub](https://github.com/spiegel-im-spiegel/pa-api/releases/tag/v0.5.0)

[以前いただいた pull request]({{< relref "./pa-api-v5-package-v0_4_0-is-released.md" >}}) のコードがなかなかよかったので [`paapi5`]`/query` サブパッケージを全面的に書き直した。

もともと [`paapi5`]`/query` サブパッケージは [`paapi5`].`Query` インタフェースの実装サンプル・コード程度にしか考えてなかったのだが， [PA-API] の `GetItems` および `SearchItems` オペレーションに限れば，そのまま使えそうな感じになってしまった。
PR を送っていただいた方にはホンマに感謝である。

コメントとかドキュメントはまだまだ整理中なので，長い目で見てやってください。
あっ  PR はいつでも歓迎です。

## ブックマーク

- [PA-API v5 への移行]({{< ref "/remark/2019/10/pa-api-v5.md" >}})
- [Go 言語用 PA-API v5 クライアント・パッケージ]({{< ref "/release/pa-api-v5.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[PA-API]: https://affiliate.amazon.co.jp/assoc_credentials/home "Product Advertising API"
[spiegel-im-spiegel/pa-api]: https://github.com/spiegel-im-spiegel/pa-api "spiegel-im-spiegel/pa-api: APIs for Amazon Product Advertising API v5 by Golang"
[`paapi5`]: https://github.com/spiegel-im-spiegel/pa-api "spiegel-im-spiegel/pa-api: APIs for Amazon Product Advertising API v5 by Golang"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN), 4621300253 (ISBN), 9784621300251 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
