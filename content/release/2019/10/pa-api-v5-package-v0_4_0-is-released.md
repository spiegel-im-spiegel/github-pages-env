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

引数のオプション関数の仕様を変更した。
マーケットプレイスとして以下のシンボルをセットできる。

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

- [Go 言語用 PA-API v5 クライアント・パッケージ]({{< ref "/release/pa-api-v5.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[PA-API]: https://affiliate.amazon.co.jp/assoc_credentials/home "Product Advertising API"
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
