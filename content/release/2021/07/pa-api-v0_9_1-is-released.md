+++
title = "PA-APIv5 クライアント・パッケージ v0.9.1 をリリースした"
date =  "2021-07-16T21:18:45+09:00"
description = "Pull request をいただいたので反映した。"
image = "/images/attention/go-logo_blue.png"
tags = [ "package", "module", "golang", "amazon", "pa-api" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

ちょっとした改修の pull request をいただいたので，これを反映した Amazon [Product Advertising API][PA-API] v5 (PA-APIv5) クライアント・パッケージ [spiegel-im-spiegel/pa-api] の v0.9.1 をリリースした。
と思ったのだが，何故か v0.9.0 のアナウンスをしてなかったみたいなので，併せて紹介する。

- [Release v0.9.0 · spiegel-im-spiegel/pa-api · GitHub](https://github.com/spiegel-im-spiegel/pa-api/releases/tag/v0.9.0)
- [Release v0.9.1 · spiegel-im-spiegel/pa-api · GitHub](https://github.com/spiegel-im-spiegel/pa-api/releases/tag/v0.9.1)

[v0.8.0 をリリースした]({{< ref "/release/2020/11/pa-api-v0_8_0-is-released.md" >}} "PA-APIv5 クライアント・パッケージ v0.8.0 をリリースした")ときに「いっそ interface 型で汎化してしまうか」と書いたのだが，実際に v0.9.0 では `Marketplace` 型を

```go
type Marketplace interface {
    String() string   // marketplace name (www.amazon.com)
    HostName() string // hostname (webservices.amazon.com)
    Region() string   // region (us-east-1)
    Language() string // language code (en_US)
}
```

と interface 型で定義しなおすことで自前でマーケットプレイス情報を定義できるようにした。
なお，元々の列挙型は `MarketplaceEnum` で再定義している。

私が公開しているパッケージでは珍しく pull request をいただくことが多いのだが，使っていただけるのなら幸いである。

[Go]: https://golang.org/ "The Go Programming Language"
[spiegel-im-spiegel/pa-api]: https://github.com/spiegel-im-spiegel/pa-api "spiegel-im-spiegel/pa-api: APIs for Amazon Product Advertising API v5 by Golang"
[PA-API]: https://affiliate.amazon.co.jp/assoc_credentials/home "Product Advertising API"

## 参考図書

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
