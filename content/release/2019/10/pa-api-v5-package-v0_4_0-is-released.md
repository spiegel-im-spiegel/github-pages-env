+++
title = "PA-API v5 クライアント・パッケージ v0.4.0 をリリースした"
date =  "2019-10-15T21:52:44+09:00"
description = "暇にちょっとずつコードもコメント・ドキュメントも整理していく予定である。なので長い目で見てやってください。"
image = "/images/attention/go-logo_blue.png"
tags = [ "package", "module", "golang", "amazon", "pa-api", "programming" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

いや， [Go 言語]用の [PA-API] v5 クライアント・パッケージなんか私以外に需要はないと思ってたのよ。
したら pull request 貰っちゃって大慌て（笑）

貰った PR を受け入れた後，コメント・ドキュメントの整理と若干の機能追加をした v0.4.0 をリリースした。

- [Release v0.4.0 · spiegel-im-spiegel/pa-api · GitHub](https://github.com/spiegel-im-spiegel/pa-api/releases/tag/v0.4.0)
- [Release v0.4.1 · spiegel-im-spiegel/pa-api · GitHub](https://github.com/spiegel-im-spiegel/pa-api/releases/tag/v0.4.1)
- [Release v0.4.2 · spiegel-im-spiegel/pa-api · GitHub](https://github.com/spiegel-im-spiegel/pa-api/releases/tag/v0.4.2)

まずは，アイテム検索用の query 構造体を作ってもらっちゃった。
これを使って，こんな感じにアイテムの検索ができる。

{{< highlight go "hl_lines=25-30" >}}
package main

import (
    "context"
    "fmt"
    "net/http"

    paapi5 "github.com/spiegel-im-spiegel/pa-api"
    "github.com/spiegel-im-spiegel/pa-api/entity"
    "github.com/spiegel-im-spiegel/pa-api/query"
)

func main() {
    //Create client
    client := paapi5.New(
        paapi5.WithMarketplace(paapi5.LocaleJapan),
    ).CreateClient(
        "mytag-20",             //Amazon associate tag
        "AKIAIOSFODNN7EXAMPLE", //access key for PA-API
        "1234567890",           //secret key for PA-API
        paapi5.WithContext(context.Background()),
        paapi5.WithHttpClient(http.DefaultClient),
    )

    //Make query
    q := query.NewSearchItems(
        client.Marketplace(),
        client.PartnerTag(),
        client.PartnerType(),
    ).Search("数学ガール", query.Keywords)

    //Requet and response
    body, err := client.Request(q)
    if err != nil {
        fmt.Printf("%+v\n", err)
        return
    }
    //io.Copy(os.Stdout, bytes.NewReader(body))

    //Decode JSON
    res, err := entity.DecodeResponse(body)
    if err != nil {
        fmt.Printf("%+v\n", err)
        return
    }
    fmt.Println(res.String())
}
{{< /highlight >}}

あっ，もちろんアクセス・キー等は自分のを使ってね。

サーバ側の情報を格納する [`paapi5`]`.Server` インスタンスを生成しているのが以下の部分だが

```go
paapi5.New(
    paapi5.WithMarketplace(paapi5.LocaleJapan),
)
```

引数の[関数オプション]({{< ref "/golang/functional-options-pattern.md" >}})の仕様を変更した。
以下のシンボルをマーケットプレイスとしてセットできる。

```go
//Marketplace is enumeration of locale information
type Marketplace int

const (
    LocaleUnknown            Marketplace = iota //Unknown local
    LocaleAustralia                             //Australia
    LocaleBrazil                                //Brazil
    LocaleCanada                                //Canada
    LocaleFrance                                //France
    LocaleGermany                               //Germany
    LocaleIndia                                 //India
    LocaleItaly                                 //Italy
    LocaleJapan                                 //Japan
    LocaleMexico                                //Mexico
    LocaleSpain                                 //Spain
    LocaleTurkey                                //Turkey
    LocaleUnitedArabEmirates                    //United Arab Emirates
    LocaleUnitedKingdom                         //United Kingdom
    LocaleUnitedStates                          //United States
)
```

これらのシンボルからマーケットプレイス名，サービスサーバ名，リージョン，言語の情報を取得できる。
既定値は [`paapi5`]`.LocaleUnitedStates` とした。

他の人が使うとか勘定に入れてなかったので相当にやっつけコードだったが，余暇にちょっとずつコードもコメント・ドキュメントも整理していく予定である。
なので長い目で見てやってください。
あっ pull request はいつでも歓迎です。

## ブックマーク

- [PA-API v5 への移行]({{< ref "/remark/2019/10/pa-api-v5.md" >}})
- [Go 言語用 PA-API v5 クライアント・パッケージ]({{< ref "/release/pa-api-v5.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[PA-API]: https://affiliate.amazon.co.jp/assoc_credentials/home "Product Advertising API"
[`paapi5`]: https://github.com/spiegel-im-spiegel/pa-api "spiegel-im-spiegel/pa-api: APIs for Amazon Product Advertising API v5 by Golang"

## 参考図書

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
