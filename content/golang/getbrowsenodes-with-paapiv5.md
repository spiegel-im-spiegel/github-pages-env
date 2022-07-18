+++
title = "goark/pa-api パッケージで GetBrowseNodes オペレーションに対応した"
date =  "2022-07-18T16:01:46+09:00"
description = "まぁ，使ってみてください。"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "pa-api", "amazon" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今更であるが，[Amazon PA-APIv5][PA-API] にアクセスするための自作パッケージ [`goark/pa-api`][github.com/goark/pa-api] は実はコードの半分以上が他人様の PR でできているのだが（感謝！），今回は `GetBrowseNodes` オペレーション追加をリクエストされまして

- [Enable to get browse node Info · Issue #25 · goark/pa-api · GitHub](https://github.com/goark/pa-api/issues/25)

たまには自分でコードを書こうと思い立ったのだった。
ちなみに `GetBrowseNodes` オペレーションの API 仕様は以下の通り。

- [GetBrowseNodes · Product Advertising API 5.0](https://webservices.amazon.com/paapi5/documentation/getbrowsenodes.html)

ここで browse node というのは構造化された商品カテゴリのようなものらしい。

- [Browse Nodes · Product Advertising API 5.0](https://webservices.amazon.com/paapi5/documentation/use-cases/organization-of-items-on-amazon/browse-nodes.html)

[`goark/pa-api`][github.com/goark/pa-api] を使って browse node の情報を取得してみよう。
ちなみに [PA-API] アクセス用パラメータを以下に例示する。

| パラメータ名       | 値                     |
| ------------------ | ---------------------- |
| マーケットプレイス | `www.amazon.co.jp`     |
| アソシエイト・タグ | `mytag-20`             |
| アクセス・キー     | `AKIAIOSFODNN7EXAMPLE` |
| シークレット・キー | `1234567890`           |

もちろん実際には使えないのでご安心を（笑） さっそくコードを書いてみる。

```go
//go:build run
// +build run

package main

import (
    "bytes"
    "context"
    "fmt"
    "io"
    "os"

    paapi5 "github.com/goark/pa-api"
    "github.com/goark/pa-api/query"
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
    q := query.NewGetBrowseNodes(
        client.Marketplace(),
        client.PartnerTag(),
        client.PartnerType(),
    ).BrowseNodeIds([]string{"2291970051"})

    //Requet and response
    body, err := client.RequestContext(context.Background(), q)
    if err != nil {
        fmt.Printf("%+v\n", err)
        return
    }
    io.Copy(os.Stdout, bytes.NewReader(body))
}
```

これを実行してみよう。

```text
$ go run sample1.go | jq .
{
  "BrowseNodesResult": {
    "BrowseNodes": [
      {
        "ContextFreeName": "産業研究",
        "DisplayName": "産業研究",
        "Id": "2291970051",
        "IsRoot": false
      }
    ]
  }
}
```

これで browse node ID が 2291970051 の情報が取得する。
さらに

```go
//Make query
q := query.NewGetBrowseNodes(
    client.Marketplace(),
    client.PartnerTag(),
    client.PartnerType(),
).BrowseNodeIds([]string{"2291970051"}).EnableBrowseNodes()
```

と，リソースパラメータを有効にすると

```text
$ go run sample2.go | jq .
{
  "BrowseNodesResult": {
    "BrowseNodes": [
      {
        "Ancestor": {
          "Ancestor": {
            "Ancestor": {
              "Ancestor": {
                "ContextFreeName": "Kindleストア",
                "DisplayName": "Kindleストア",
                "Id": "2250738051"
              },
              "ContextFreeName": "Kindleストア",
              "DisplayName": "カテゴリー別",
              "Id": "2250739051"
            },
            "ContextFreeName": "Kindle本",
            "DisplayName": "Kindle本",
            "Id": "2275256051"
          },
          "ContextFreeName": "ビジネス・経済",
          "DisplayName": "ビジネス・経済",
          "Id": "2291905051"
        },
        "Children": [
          {
            "ContextFreeName": "コンサルティング",
            "DisplayName": "コンサルティング",
            "Id": "4715053051"
          },
          {
            "ContextFreeName": "経営情報システム",
            "DisplayName": "経営情報システム",
            "Id": "4715061051"
          },
          {
            "ContextFreeName": "銀行・金融業",
            "DisplayName": "銀行・金融業",
            "Id": "2292599051"
          },
          {
            "ContextFreeName": "人事・労務管理",
            "DisplayName": "人事・労務管理",
            "Id": "2291959051"
          },
          {
            "ContextFreeName": "経理・アカウンティング",
            "DisplayName": "会計",
            "Id": "2292083051"
          },
          {
            "ContextFreeName": "e コマース",
            "DisplayName": "e コマース",
            "Id": "2291907051"
          },
          {
            "ContextFreeName": "NGO・NPO",
            "DisplayName": "非営利組織",
            "Id": "2293150051"
          },
          {
            "ContextFreeName": "不動産",
            "DisplayName": "不動産",
            "Id": "2291980051"
          }
        ],
        "ContextFreeName": "産業研究",
        "DisplayName": "産業研究",
        "Id": "2291970051",
        "IsRoot": false
      }
    ]
  }
}
```

という感じに親ノードと子ノードも併せて取得できる。

この browse node ID を使って商品検索結果を絞り込むことができる。
たとえば結城浩さんの著作を 2291970051 で絞り込んでみる。

```go
//Make query
q := query.NewSearchItems(
    client.Marketplace(),
    client.PartnerTag(),
    client.PartnerType(),
).Search(query.Author, "結城浩").Request(query.BrowseNodeID, "2291970051").EnableItemInfo()
```

これを実行すると

```text
$ go run sample3c.go | jq .
{
  "SearchResult": {
    "Items": [
      {
        "ASIN": "B08S2LY9VG",
        "DetailPageURL": "https://www.amazon.co.jp/dp/B08S2LY9VG?tag=mytag-20&linkCode=osi&th=1&psc=1",
        "ItemInfo": {
          "ByLineInfo": {
            "Manufacturer": {
              "DisplayValue": "SBクリエイティブ",
              "Label": "Manufacturer",
              "Locale": "ja_JP"
            },
            "Contributors": [
              {
                "Name": "結城 浩",
                "Locale": "ja_JP",
                "Role": "著"
              }
            ]
          },
          "Classifications": {
            "Binding": {
              "DisplayValue": "Kindle版",
              "Label": "Binding",
              "Locale": "ja_JP"
            },
            "ProductGroup": {
              "DisplayValue": "Digital Ebook Purchas",
              "Label": "ProductGroup",
              "Locale": "ja_JP"
            }
          },
          "ContentInfo": {
            "Languages": {
              "DisplayValues": [
                {
                  "DisplayValue": "日本語",
                  "Type": "発行済み"
                }
              ],
              "Label": "Language",
              "Locale": "ja_JP"
            },
            "PagesCount": {
              "DisplayValue": 218,
              "Label": "NumberOfPages",
              "Locale": "en_US"
            },
            "PublicationDate": {
              "DisplayValue": "2021-02-19T00:00:00Z",
              "Label": "PublicationDate",
              "Locale": "en_US"
            }
          },
          "ProductInfo": {
            "IsAdultProduct": {
              "DisplayValue": false,
              "Label": "IsAdultProduct",
              "Locale": "en_US"
            },
            "ReleaseDate": {
              "DisplayValue": "2021-02-20T00:00:00Z",
              "Label": "ReleaseDate",
              "Locale": "en_US"
            }
          },
          "TechnicalInfo": {
            "Formats": {
              "DisplayValues": [
                "Kindle本"
              ],
              "Label": "Format",
              "Locale": "ja_JP"
            }
          },
          "Title": {
            "DisplayValue": "再発見の発想法",
            "Label": "Title",
            "Locale": "ja_JP"
          }
        }
      }
    ],
    "SearchURL": "https://www.amazon.co.jp/s?rh=n%3A2291970051%2Cp_lbr_three_browse-bin%3A%E7%B5%90%E5%9F%8E%E6%B5%A9%2Cp_n_availability%3A-1&tag=mytag-20&linkCode=osi",
    "TotalResultCount": 1
  }
}
```

という感じに『[再発見の発想法](https://www.amazon.co.jp/dp/B08S2LY9VG?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "再発見の発想法 | 結城 浩 | 産業研究 | Kindleストア | Amazon")』のみ絞り込めた。

まぁ，こんな感じに使ってみてください。

[Go]: https://go.dev/
[github.com/goark/pa-api]: https://github.com/goark/pa-api "goark/pa-api: APIs for Amazon Product Advertising API v5 by Golang"
[`paapi5`]: https://github.com/goark/pa-api "goark/pa-api: APIs for Amazon Product Advertising API v5 by Golang"
[PA-API]: https://affiliate.amazon.co.jp/assoc_credentials/home "Product Advertising API"

## ブックマーク

- [Go 言語用 PA-API v5 クライアント・パッケージ]({{< ref "/release/pa-api-v5.md" >}})
- [amazon APIのBrowseNode ID一覧ってどうやって調べるの？ – ユズムログ](http://yuzum.com/blog/20210124_20210120_api)

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B09C2XBC2F" %}} <!-- Golang Tシャツ -->
{{% review-paapi "B08S2LY9VG" %}} <!-- 再発見の発想法 -->
