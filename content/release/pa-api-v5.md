+++
title = "Go 言語用 PA-API v5 クライアント・パッケージ"
date =  "2019-10-06T15:46:03+09:00"
description = "本パッケージは Amazon Product Advertising API v5 へアクセスできる Go 言語用クライアント・パッケージだ。 API を通じて Ammazon で取り扱っている商品の情報を取得できる。"
image = "/images/attention/go-logo_blue.png"
tags = [ "package", "module", "golang", "amazon", "pa-api", "programming" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

- [spiegel-im-spiegel/pa-api: APIs for Amazon Product Advertising API v5 by Golang](https://github.com/spiegel-im-spiegel/pa-api)

本パッケージは Amazon [Product Advertising API] v5 へアクセスできる [Go 言語]用クライアント・パッケージだ。
API を通じて Amazon で取り扱っている商品の情報を取得できる。

なお [spiegel-im-spiegel/pa-api] パッケージは [Go] 1.13 以上を要求する。
ご注意を。

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/pa-api.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/pa-api)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/pa-api/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/spiegel-im-spiegel/pa-api.svg)](https://github.com/spiegel-im-spiegel/pa-api/releases/latest)

## インポート

```go
import "github.com/spiegel-im-spiegel/pa-api"
```

## 簡単な使い方

### [PA-API] アクセス用のパラメータ

使い方の前に，この記事で使用する [PA-API] アクセス用のパラメータを以下に例示しておく。

| パラメータ名       | 値                     |
| ------------------ | ---------------------- |
| マーケットプレイス | `www.amazon.co.jp`     |
| アソシエイト・タグ | `mytag-20`             |
| アクセス・キー     | `AKIAIOSFODNN7EXAMPLE` |
| シークレット・キー | `1234567890`           |

もちろん実際には使えないのでご安心を（笑）

### サンプル・コード

以下に簡単なコード例を示す。

```go
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
        paapi5.WithContext(context.Background()),
        paapi5.WithHttpClient(http.DefaultClient),
    )

    //Make query
    q := query.NewGetItems(client.Marketplace(), client.PartnerTag(), client.PartnerType())
    q.ASINs([]string{"B07YCM5K55"}).EnableImages(true).EnableParentASIN(true)

    //Requet and response
    body, err := client.Request(q)
    if err != nil {
        fmt.Printf("%+v\n", err)
        return
    }
    io.Copy(os.Stdout, bytes.NewReader(body))
}
```

このコードの実行結果はこんな感じ。

```json
$ go run sample.go | jq .
{
  "ItemsResult": {
    "Items": [
      {
        "ASIN": "B07YCM5K55",
        "DetailPageURL": "https://www.amazon.co.jp/dp/B07YCM5K55?tag=mytag-20&linkCode=ogi&th=1&psc=1",
        "Images": {
          "Primary": {
            "Large": {
              "Height": 500,
              "URL": "https://m.media-amazon.com/images/I/51Ef9EF+HaL.jpg",
              "Width": 352
            },
            "Medium": {
              "Height": 160,
              "URL": "https://m.media-amazon.com/images/I/51Ef9EF+HaL._SL160_.jpg",
              "Width": 113
            },
            "Small": {
              "Height": 75,
              "URL": "https://m.media-amazon.com/images/I/51Ef9EF+HaL._SL75_.jpg",
              "Width": 53
            }
          }
        },
        "ItemInfo": {
          "ByLineInfo": {
            "Contributors": [
              {
                "Locale": "ja_JP",
                "Name": "宮成楽",
                "Role": "著"
              }
            ],
            "Manufacturer": {
              "DisplayValue": "竹書房",
              "Label": "Manufacturer",
              "Locale": "ja_JP"
            }
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
              "DisplayValue": 116,
              "Label": "NumberOfPages",
              "Locale": "en_US"
            },
            "PublicationDate": {
              "DisplayValue": "2019-09-27T00:00:00.000Z",
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
              "DisplayValue": "2019-09-27T00:00:00.000Z",
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
            "DisplayValue": "晴れのちシンデレラ　（１２） (バンブーコミックス MOMOセレクション)",
            "Label": "Title",
            "Locale": "ja_JP"
          }
        }
      }
    ]
  }
}
```

### リクエスト実行インスタンスの生成

上のコード例の

```go
client := paapi5.New(
    paapi5.WithMarketplace(paapi5.LocaleJapan),
).CreateClient(
    "mytag-20",
    "AKIAIOSFODNN7EXAMPLE",
    "1234567890",
    paapi5.WithContext(context.Background()),
    paapi5.WithHttpClient(http.DefaultClient),
)
```

がリクエストを実行する [`paapi5`]`.Client` 型のインスタンスを生成している部分である。
ちなみに

```go
client := paapi5.DefaultClient("mytag-20", "AKIAIOSFODNN7EXAMPLE", "1234567890") //Create default client
```

と既定値で簡単に [`paapi5`]`.Client` インスタンスを生成することもできるが，マーケットプレイスが米国（`www.amazon.com`）になるのでご注意を。

### クエリの生成とリクエストの実行

リクエストを実行する関数

```go
func (c *Client) Request(q Query) ([]byte, error) {
    ...
}
```

の引数 [`paapi5`]`.Query` は interface 型で以下のように定義している。

```go
type Query interface {
    Operation() Operation
    Payload() ([]byte, error)
}
```

[`paapi5`]`.Query.Operation()` 関数の返り値の [`paapi5`]`.Operation` 型は以下のように定義している。

```go
type Operation int

const (
    NullOperation Operation = iota
    GetVariations
    GetItems
    SearchItems
)
```

[`paapi5`]`.Query.Payload()` 関数はリクエストのペイロードにセットする JSON データを返す。
JSON データの内容は，例えばこんな感じ。

```go
{
 "ItemIds": [
  "B07YCM5K55"
 ],
 "Resources": [
  "Images.Primary.Small",
  "Images.Primary.Medium",
  "Images.Primary.Large",
  "ItemInfo.ByLineInfo",
  "ItemInfo.ContentInfo",
  "ItemInfo.Classifications",
  "ItemInfo.ExternalIds",
  "ItemInfo.ProductInfo",
  "ItemInfo.Title"
 ],
 "PartnerTag": "mytag-20",
 "PartnerType": "Associates",
 "Marketplace": "www.amazon.co.jp",
 "Operation": "GetItems"
}
```

適切な [`paapi5`]`.Operation` 値と JSON データを出力できるのであれば，利用者側でクエリ・オブジェクトを自由に設計できる。

{{< div-box type="md" >}}
#### 余談だが...

ある型が特定の interface 型に適合するかコンパイル時点でチェックするには以下の記述を加えるとよい。

```go
var _ paapi5.Query = (*CustomQuery)(nil)
```
{{< /div-box >}}

### クエリの実例とレスポンスの取り込み例

[spiegel-im-spiegel/pa-api] パッケージではクエリ用のサンプルとして [`paapi5`]`.query` サブパッケージを用意している。
最初のコード例の

```go
q := query.NewGetItems(client.Marketplace(), client.PartnerTag(), client.PartnerType())
q.ASINs([]string{"B07YCM5K55"}).EnableImages(true).EnableParentASIN(true)
```

の部分がそれである。
また [`paapi5`]`.Client.Request()` 関数の出力結果を構造体に落とし込むための [`paapi5`]`.entity` サブパッケージも用意した。
どちらもそのままではあまり使い勝手がいいとは言えないが，コード例として自由に利用していただいて構わない。

なお [spiegel-im-spiegel/pa-api] パッケージは [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0) でライセンスしている。

## ブックマーク

- [Product Advertising API 5.0 Documentation](https://webservices.amazon.com/paapi5/documentation/)
- [PA-API V5 への移行]({{< ref "/remark/2019/10/pa-api-v5.md" >}})

[spiegel-im-spiegel/pa-api]: https://github.com/spiegel-im-spiegel/pa-api "spiegel-im-spiegel/pa-api: APIs for Amazon Product Advertising API v5 by Golang"
[`paapi5`]: https://github.com/spiegel-im-spiegel/pa-api "spiegel-im-spiegel/pa-api: APIs for Amazon Product Advertising API v5 by Golang"
[Product Advertising API]: https://affiliate.amazon.co.jp/assoc_credentials/home
[PA-API]: https://affiliate.amazon.co.jp/assoc_credentials/home "Product Advertising API"
[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`context`]: https://golang.org/pkg/context/ "context - The Go Programming Language"
[`http`]: https://golang.org/pkg/net/http/ "http - The Go Programming Language"

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