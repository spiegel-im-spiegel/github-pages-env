+++
title = "Go 言語用 PA-API v5 クライアント・パッケージ v0.7.0 をリリースした"
date =  "2020-06-30T14:26:02+09:00"
description = "CustomerReviews 項目を指定できるようにしてもらった。"
image = "/images/attention/go-logo_blue.png"
tags = [ "package", "module", "golang", "amazon", "pa-api" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

またも [pull request] をいただきまして，マージして v0.7.0 をリリースしました。

- [Release v0.7.0 · spiegel-im-spiegel/pa-api · GitHub](https://github.com/spiegel-im-spiegel/pa-api/releases/tag/v0.7.0)

ついでにコメントの不備も直してもらっちゃいました。

ありがたや `m(_ _)m`

[PA-API v5 のドキュメント](https://webservices.amazon.com/paapi5/documentation/ "Product Advertising API 5.0 Documentation")には書かれてないのだが，問い合わせ時の resources parameter に `CustomerReviews` の項目があるみたいで  [`paapi5`]`/query` サブパッケージで指定できるようにしてもらった。

こんな感じ。

```go {hl_lines=[30]}
package main

import (
    "bytes"
    "fmt"
    "io"
    "os"

    paapi5 "github.com/spiegel-im-spiegel/pa-api"
    "github.com/spiegel-im-spiegel/pa-api/query"
)

func main() {
    //Create client
    client := paapi5.New(
        paapi5.WithMarketplace(paapi5.LocaleJapan),
    ).CreateClient(
        "mytag-20",
        "AKIAIOSFODNN7EXAMPLE",
        "1234567890",
    )

    //Make query
    q := query.NewGetItems(
		client.Marketplace(),
		client.PartnerTag(),
		client.PartnerType(),
	).
		ASINs([]string{"B07NVMYB7K"}).
		EnableCustomerReviews()

    //Requet and response
    body, err := client.Request(q)
    if err != nil {
        fmt.Printf("%+v\n", err)
        return
    }
    io.Copy(os.Stdout, bytes.NewReader(body))
}
```

これでリクエスト側のペイロードの内容は

```json {hl_lines=["6-7"]}
{
 "ItemIds": [
  "B07NVMYB7K"
 ],
 "Resources": [
  "CustomerReviews.Count",
  "CustomerReviews.StarRating"
 ],
 "PartnerTag": "mytag-20",
 "PartnerType": "Associates",
 "Marketplace": "www.amazon.co.jp",
 "Operation": "GetItems"
}
```

となり，返ってきた値も

{{< fig-img src="./paapi5-itemresult.png" title="Added undocumented Customer Rating to API" link="https://github.com/spiegel-im-spiegel/pa-api/pull/13" >}}

となる筈，なのだが，手元では `CustomerReviews` 項目が上手く取れなかった（指定してもエラーにはならない）。
いただいた [pull request] には

{{< fig-quote type="markdown" title="Added undocumented Customer Rating to API" link="https://github.com/spiegel-im-spiegel/pa-api/pull/13" lang="en" >}}
{{% quote %}}It seems you have to request this functionality directly at Amazon to be able to use the resource.  My organization got approval and therefore needs this resource in this API.{{% /quote %}}
{{< /fig-quote >}}

とあるので `CustomerReviews` 項目を使えるようにする何らかの手続きが必要なのかもしれない（日本の marketplace が対応してないだけかもしれないが）。
まぁ，私個人は使わない項目なので，困らないのが困りもの（笑）

ちうわけで，日本の環境で `CustomerReviews` 周りの情報を知っておられる方は是非教えてください。

## ブックマーク

- [Go 言語用 PA-API v5 クライアント・パッケージ]({{< ref "/release/pa-api-v5.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[PA-API]: https://affiliate.amazon.co.jp/assoc_credentials/home "Product Advertising API"
[spiegel-im-spiegel/pa-api]: https://github.com/spiegel-im-spiegel/pa-api "spiegel-im-spiegel/pa-api: APIs for Amazon Product Advertising API v5 by Golang"
[`paapi5`]: https://github.com/spiegel-im-spiegel/pa-api "spiegel-im-spiegel/pa-api: APIs for Amazon Product Advertising API v5 by Golang"
[pull request]: https://github.com/spiegel-im-spiegel/pa-api/pull/13 "Added undocumented Customer Rating to API by hackmac89 · Pull Request #13 · spiegel-im-spiegel/pa-api"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
