+++
title = "Go 言語用 openBD クライアント・パッケージ"
date =  "2019-09-07T22:10:02+09:00"
description = "本パッケージは openBD が提供する書籍情報を取得できる。"
image = "/images/attention/go-logo_blue.png"
tags = [ "package", "module", "golang", "openbd", "openbd-api", "book" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

- [goark/openbd-api: APIs for openBD by Golang](https://github.com/goark/openbd-api)

本パッケージは [openBD] へアクセスできる [Go 言語]用クライアント・パッケージだ。
[openBD] が提供する書籍情報を取得できる。

なお [goark/openbd-api] パッケージは [Go] 1.16 以上を要求する。
ご注意を。

[![check vulns](https://github.com/goark/openbd-api/workflows/vulns/badge.svg)](https://github.com/goark/openbd-api/actions)
[![lint status](https://github.com/goark/openbd-api/workflows/lint/badge.svg)](https://github.com/goark/openbd-api/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/goark/openbd-api/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/goark/openbd-api.svg)](https://github.com/goark/openbd-api/releases/latest)

## インポート

```go
import "github.com/goark/openbd-api"
```

## 書籍情報の取得

[openBD] は ISBN をキーとして書籍情報の検索ができる。
[goark/openbd-api] パッケージでは以下のように記述する。

```go
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/goark/openbd-api"
)

func main() {
	b, err := openbd.DefaultClient().LookupBooksRaw([]string{"9784274069321"})
	if err != nil {
		fmt.Println(err)
		return
	}
	io.Copy(os.Stdout, bytes.NewReader(b))
}
```

[`openbd`]`.Client.LookupBooksRaw()` 関数は，引数に複数の ISBN を指定することが可能で [openBD] から取得した結果（JSON 形式）をそのまま返す。

```text
$ go run sample.go | jq .
[
  {
    "onix": {
      "RecordReference": "9784274069321",
      "NotificationType": "03",
      "ProductIdentifier": {
        "ProductIDType": "15",
        "IDValue": "9784274069321"
      },
      "DescriptiveDetail": {
        "ProductComposition": "00",
        "ProductForm": "BZ",
        "Measure": [
          {
            "MeasureType": "01",
            "Measurement": "210",
            "MeasureUnitCode": "mm"
          },
          {
            "MeasureType": "02",
            "Measurement": "0",
            "MeasureUnitCode": "mm"
          }
        ],
        "TitleDetail": {
          "TitleType": "01",
          "TitleElement": {
            "TitleElementLevel": "01",
            "TitleText": {
              "collationkey": "リーン カイハツ ノ ゲンバ : カンバン ニ ヨル ダイキボ プロジェクト ノ ウンエイ",
              "content": "リーン開発の現場 : カンバンによる大規模プロジェクトの運営"
            }
          }
        },
        "Contributor": [
          {
            "SequenceNumber": "1",
            "ContributorRole": [
              "A01"
            ],
            "PersonName": {
              "content": "Kniberg, Henrik"
            }
          },
          {
            "SequenceNumber": "2",
            "ContributorRole": [
              "A01"
            ],
            "PersonName": {
              "content": "オーム社"
            }
          },
          {
            "SequenceNumber": "3",
            "ContributorRole": [
              "A01"
            ],
            "PersonName": {
              "content": "オーム社開発局"
            }
          },
          {
            "SequenceNumber": "4",
            "ContributorRole": [
              "B06"
            ],
            "PersonName": {
              "content": "市谷 聡啓"
            }
          },
          {
            "SequenceNumber": "5",
            "ContributorRole": [
              "B06"
            ],
            "PersonName": {
              "content": "藤原 大"
            }
          },
          {
            "SequenceNumber": "6",
            "ContributorRole": [
              "B06"
            ],
            "PersonName": {
              "content": "角谷 信太郎"
            }
          },
          {
            "SequenceNumber": "7",
            "ContributorRole": [
              "A01"
            ],
            "PersonName": {
              "content": "クニバーグ ヘンリック"
            }
          }
        ],
        "Language": [
          {
            "LanguageRole": "01",
            "LanguageCode": "jpn",
            "CountryCode": "JP"
          }
        ],
        "Extent": [
          {
            "ExtentType": "11",
            "ExtentValue": "190",
            "ExtentUnit": "03"
          }
        ]
      },
      "CollateralDetail": {
        "TextContent": [
          {
            "TextType": "03",
            "ContentAudience": "00",
            "Text": "官公庁の大規模システム開発における、カンバンシステムを軸にしたプロジェクト進行の様子を、著者の経験に基づいて描写。「リーンソフトウェア開発」を現場でどのように適用するかを直裁的に解説しています。"
          },
          {
            "TextType": "04",
            "ContentAudience": "00",
            "Text": "第1部 僕らのやり方を伝えよう(プロジェクトについて\nチーム編成\nデイリーカクテルパーティーに参加しよう\nプロジェクトボード\nカンバンボードをスケールさせる ほか)\n第2部 テクニックを詳しく見る(アジャイルとリーンの概要\nテスト自動化の戦略\nプランニングポーカーによる見積り\n因果関係図\n最後に伝えたいこと)"
          }
        ],
        "SupportingResource": [
          {
            "ResourceContentType": "01",
            "ContentAudience": "01",
            "ResourceMode": "03",
            "ResourceVersion": [
              {
                "ResourceForm": "02",
                "ResourceVersionFeature": [
                  {
                    "ResourceVersionFeatureType": "01",
                    "FeatureValue": "D502"
                  },
                  {
                    "ResourceVersionFeatureType": "04",
                    "FeatureValue": "9784274069321.jpg"
                  }
                ],
                "ResourceLink": "https://cover.openbd.jp/9784274069321.jpg"
              }
            ]
          }
        ]
      },
      "PublishingDetail": {
        "Imprint": {
          "ImprintIdentifier": [
            {
              "ImprintIDType": "19",
              "IDValue": "274"
            }
          ],
          "ImprintName": "オーム社"
        },
        "PublishingDate": [
          {
            "PublishingDateRole": "01",
            "Date": ""
          }
        ]
      },
      "ProductSupply": {
        "SupplyDetail": {
          "ReturnsConditions": {
            "ReturnsCodeType": "04",
            "ReturnsCode": "02"
          },
          "ProductAvailability": "99"
        }
      }
    },
    "hanmoto": {
      "datecreated": "2016-08-24 22:49:25",
      "dateshuppan": "2013-10",
      "datemodified": "2016-08-24 22:49:25"
    },
    "summary": {
      "isbn": "9784274069321",
      "title": "リーン開発の現場 : カンバンによる大規模プロジェクトの運営",
      "volume": "",
      "series": "",
      "publisher": "オーム社",
      "pubdate": "2013-10",
      "cover": "https://cover.openbd.jp/9784274069321.jpg",
      "author": "Kniberg,Henrik／著 オーム社／著 オーム社開発局／著 市谷聡啓／翻訳 ほか"
    }
  }
]
```

[`openbd`]`.Client.LookupBooks()` 関数を使うと結果を [`openbd`]`.Book` 構造体の配列で返す。

```go
books, err := openbd.DefaultClient().LookupBooks([]string{"9784274069321"})
```

[`openbd`]`.Book` 構造体の構成は以下の通り。

```go
//Book is entity class of book info.
type Book struct {
	Onix    Onix    `json:"onix"`
	Hanmoto Hanmoto `json:"hanmoto"`
	Summary Summary `json:"summary"`
}

//Onix is entity class of JPRO-onix items
type Onix struct {
	RecordReference   string //ISBN code (key code)
	NotificationType  string
	ProductIdentifier struct {
		ProductIDType string
		IDValue       string //ISBN ?
	}
	DescriptiveDetail struct {
		ProductComposition string
		ProductForm        string
		Measure            []struct {
			MeasureType     string
			Measurement     string
			MeasureUnitCode string
		} `json:",omitempty"`
		Collection struct {
			CollectionType     string
			CollectionSequence *struct {
				CollectionSequenceType     string `json:",omitempty"`
				CollectionSequenceTypeName string `json:",omitempty"`
				CollectionSequenceNumber   string `json:",omitempty"`
			} `json:",omitempty"`
			TitleDetail *struct {
				TitleType    string `json:",omitempty"`
				TitleElement []struct {
					TitleElementLevel string
					TitleText         struct {
						Content      string `json:"content"`
						CollationKey string `json:"collationkey,omitempty"`
					}
				} `json:",omitempty"`
			} `json:",omitempty"`
		}
		TitleDetail struct {
			TitleType    string
			TitleElement struct {
				TitleElementLevel string
				TitleText         struct {
					Content      string `json:"content"`
					Collationkey string `json:"collationkey,omitempty"`
				}
			}
		}
		Contributor []struct {
			SequenceNumber  string
			ContributorRole []string `json:",omitempty"`
			PersonName      struct {
				Content      string `json:"content"`
				Collationkey string `json:"collationkey,omitempty"`
			}
			BiographicalNote string `json:",omitempty"`
		} `json:",omitempty"`
		Language []struct {
			LanguageRole string
			LanguageCode string
			CountryCode  string
		} `json:",omitempty"`
		Extent []struct {
			ExtentType  string
			ExtentValue string
			ExtentUnit  string
		} `json:",omitempty"`
		Subject []struct {
			SubjectSchemeIdentifier string
			SubjectCode             string
			SubjectHeadingText      string `json:",omitempty"`
		} `json:",omitempty"`
		Audience []struct {
			AudienceCodeType  string
			AudienceCodeValue string
		} `json:",omitempty"`
	}
	CollateralDetail *struct {
		TextContent []struct {
			TextType        string
			ContentAudience string
			Text            string
		} `json:",omitempty"`
		SupportingResource []struct {
			ResourceContentType string
			ContentAudience     string
			ResourceMode        string
			ResourceVersion     []struct {
				ResourceForm           string
				ResourceVersionFeature []struct {
					ResourceVersionFeatureType string
					FeatureValue               string
				} `json:",omitempty"`
				ResourceLink string
			} `json:",omitempty"`
		} `json:",omitempty"`
	} `json:",omitempty"`
	PublishingDetail struct {
		Imprint struct {
			ImprintIdentifier []struct {
				ImprintIDType string
				IDValue       string
			} `json:",omitempty"`
			ImprintName string
		}
		Publisher struct {
			PublisherIdentifier []struct {
				PublisherIDType string
				IDValue         string
			} `json:",omitempty"`
			PublishingRole string
			PublisherName  string
		}
		PublishingDate []struct {
			Date               Date
			PublishingDateRole string
		} `json:",omitempty"`
	}
	ProductSupply struct {
		SupplyDetail struct {
			ReturnsConditions struct {
				ReturnsCodeType string
				ReturnsCode     string
			}
			ProductAvailability string
			Price               []struct {
				PriceType    string
				CurrencyCode string
				PriceAmount  string
			} `json:",omitempty"`
		}
	}
}

//Hanmoto is entity class of Hanmoto dot com items
type Hanmoto struct {
	DatePublished Date `json:"dateshuppan"`
	DateModified  Date `json:"datemodified"`
	DateCreated   Date `json:"datecreated"`
	DateReleased  Date `json:"datekoukai"`
	IsLightNovel  bool `json:"lanove,omitempty"`
	HasReview     bool `json:"hasshohyo,omitempty"`
	Reviews       []struct {
		Reviewer       string `json:"reviewer"`
		Link           string `json:"link"`
		DateAppearance Date   `json:"appearance"`
		SourceKindID   int    `json:"kubun_id"`
		SourceID       int    `json:"source_id"`
		Source         string `json:"source"`
		PaperType      string `json:"choyukan"`
		PostUser       string `json:"post_user"`
		Han            string `json:"han"`
		Gou            string `json:"gou"`
	} `json:"reviews,omitempty"`
	HasSample bool `json:"hastameshiyomi,omitempty"`
}

//Summary is entity class of summary data
type Summary struct {
	ISBN      string `json:"isbn"`
	Title     string `json:"title"`
	Volume    string `json:"volume"`
	Series    string `json:"series"`
	Publisher string `json:"publisher"`
	PubDate   Date   `json:"pubdate"`
	Author    string `json:"author"`
	Cover     string `json:"cover"`
}
```

[`openbd`]`.Book` 構造体はこのままでは使い辛いので以下のヘルパ関数を用意した。
引数がないのでテンプレートでは値のように扱うことができる。

- [`openbd`]`.Book.IsValid() bool`
- [`openbd`]`.Book.Id() string`
- [`openbd`]`.Book.ISBN() string`
- [`openbd`]`.Book.Title() string`
- [`openbd`]`.Book.SubTitle() string` : 現在は空文字列のみ返す
- [`openbd`]`.Book.SeriesTitle() string`
- [`openbd`]`.Book.Label() string`
- [`openbd`]`.Book.ImageURL() string`
- [`openbd`]`.Book.Authors() []string`
- [`openbd`]`.Book.Publisher() string`
- [`openbd`]`.Book.PublicationDate()` [`openbd`]`.Date` 
- [`openbd`]`.Book.Description() string`

## Server および Client インスタンスの生成

[`openbd`]`.Client` インスタンスの生成は [`openbd`]`.DefaultClient()` 関数で簡単に行えるが，もう少し細かい制御もできる。

### [openBD] サーバを指定する

[`openbd`]`.New()` 関数で [`openbd`]`.Server` インスタンスを生成できるが，引数としてサーバを指定できる。

```go
server := openbd.New(
    openbd.WithScheme("https"),
    openbd.WithServerName("api.example.com"),
)
```

これで [openBD] サーバとして `https://api.example.com` を指定できた（実在しない URL なので注意）。

### http.Client を指定する

[`openbd`]`.Server.CreateClient()` 関数により [`openbd`]`.Client` インスタンスを生成できるが，引数として [`http`]`.Client` インスタンスを指定できる。

```go {hl_lines=[4]}
client := openbd.New(
    openbd.WithScheme("https"),
    openbd.WithServerName("api.example.com"),
).CreateClient(openbd.WithHttpClient(&http.Client{}))
```

ちなみに [`openbd`]`.DefaultClient()` 関数は以下の記述と同等である。

```go
client := openbd.New(
    openbd.WithScheme("https"),
    openbd.WithServerName("api.openbd.jp"),
).CreateClient(openbd.WithHttpClient(&http.Client{}))
```

## context.Context 付きのアクセス

[openBD] へのアクセスの際に [`context`]`.Context` を付けることができる。

```go
rawjson, err := openbd.DefaultClient().
    LookupBooksRawContext(context.Background(), []string{"9784274069321"})
```

```go
books, err := openbd.DefaultClient().
    LookupBooksContext(context.Background(), []string{"9784274069321"})
```

[Go]: https://go.dev/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[openBD]: https://openbd.jp/ "openBD | 書誌情報・書影を自由に"
[goark/openbd-api]: https://github.com/goark/openbd-api "goark/openbd-api: APIs for openBD by Golang"
[`openbd`]: https://github.com/goark/openbd-api "goark/openbd-api: APIs for openBD by Golang"
[`http`]: https://golang.org/pkg/net/http/ "http - The Go Programming Language"
[`context`]: https://golang.org/pkg/context/ "context - The Go Programming Language"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
