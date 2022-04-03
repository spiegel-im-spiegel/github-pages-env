+++
title = "Go 言語用青空文庫 API クライアント・パッケージ"
date =  "2019-09-07T22:10:02+09:00"
description = "本パッケージは青空文庫 API を通じて青空文庫に収録されている作品情報等を取得できる。"
image = "/images/attention/go-logo_blue.png"
tags = [ "package", "module", "golang", "aozora", "aozora-api", "book" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

- [goark/aozora-api: APIs for Aozora-bunko RESTful Service by Golang](https://github.com/goark/aozora-api)

本パッケージは[青空文庫] API へアクセスできる [Go 言語]用クライアント・パッケージだ。
API を通じて[青空文庫]に収録されている作品情報等を取得できる。

なお [goark/aozora-api][`aozora`] パッケージは [Go] 1.16 以上を要求する。
ご注意を。

[![check vulns](https://github.com/goark/aozora-api/workflows/vulns/badge.svg)](https://github.com/goark/aozora-api/actions)
[![lint status](https://github.com/goark/aozora-api/workflows/lint/badge.svg)](https://github.com/goark/aozora-api/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/goark/aozora-api/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/goark/aozora-api.svg)](https://github.com/goark/aozora-api/releases/latest)

## インポート

```go
import "github.com/goark/aozora-api"
```

## 簡単な使い方

### 作品情報の取得

たとえばタイトル名「天に積む宝」著者名「富田倫生」を検索キーとして[青空文庫] API から作品情報を取得してみる。

```go
package main

import (
    "bytes"
    "fmt"
    "io"
    "os"

    "github.com/goark/aozora-api"
)

func main() {
    b, err := aozora.DefaultClient().SearchBooksRaw(
        aozora.WithBookTitle("/天に積む宝/"),
        aozora.WithBookAuthor("富田倫生"),
    )
    if err != nil {
        fmt.Println(err)
        return
    }
    io.Copy(os.Stdout, bytes.NewReader(b))
}
```

検索オプションは以下の関数と連動している。
これらの関数を[`aozora`]`.Client.SearchBooksRaw()` 関数の引数に指定する。

| オプション名 | 対応する関数                                  |
| ------------:| --------------------------------------------- |
|      `title` | [`aozora`]`.WithBookTitle(string)`            |
|     `author` | [`aozora`]`.WithBookAuthor(string)`           |
|     `fields` | [`aozora`]`.WithBookFields(string)`[^fields1] |
|      `limit` | [`aozora`]`.WithBookLimit(int)`               |
|       `skip` | [`aozora`]`.WithBookSkip(int)`                |
|      `after` | [`aozora`]`.WithBookAfter(time.Time)`         |

[^fields1]: `fields` オプションは効いてない感じ？

[`aozora`]`.Client.SearchBooksRaw()` 関数は API を通じて取得した結果（JSON 形式）をそのまま返す。

```text
$ go run sample.go | jq .
[
  {
    "book_id": 59489,
    "title": "「天に積む宝」のふやし方、へらし方",
    "title_yomi": "「てんにつむたから」のふやしかた、へらしかた",
    "title_sort": "てんにつむたからのふやしかたへらしかた",
    "subtitle": "著作権保護期間延長が青空文庫にもたらすもの",
    "subtitle_yomi": "ちょさくけんほごきかんえんちょうがあおぞらぶんこにもたらすもの",
    "original_title": "",
    "first_appearance": "",
    "ndc_code": "",
    "font_kana_type": "新字新仮名",
    "copyright": true,
    "release_date": "2019-01-01T00:00:00.000Z",
    "last_modified": "2018-12-24T00:00:00.000Z",
    "card_url": "https://www.aozora.gr.jp/cards/000055/card59489.html",
    "base_book_1": "インターネット図書館　青空文庫",
    "base_book_1_publisher": "はる書房",
    "base_book_1_1st_edition": "2005（平成17）年11月15日",
    "base_book_1_edition_input": "2005（平成17）年11月15日初版第1刷",
    "base_book_1_edition_proofing": "2005（平成17）年11月15日初版第1刷",
    "base_book_1_parent": "",
    "base_book_1_parent_publisher": "",
    "base_book_1_parent_1st_edition": "",
    "base_book_2": "",
    "base_book_2_publisher": "",
    "base_book_2_1st_edition": "",
    "base_book_2_edition_input": "",
    "base_book_2_edition_proofing": "",
    "base_book_2_parent": "",
    "base_book_2_parent_publisher": "",
    "base_book_2_parent_1st_edition": "",
    "input": "富田晶子",
    "proofing": "雪森",
    "text_url": "https://www.aozora.gr.jp/cards/000055/files/59489_txt_66663.zip",
    "text_last_modified": "2018-12-24T00:00:00.000Z",
    "text_encoding": "ShiftJIS",
    "text_charset": "JIS X 0208",
    "text_updated": 0,
    "html_url": "https://www.aozora.gr.jp/cards/000055/files/59489_66714.html",
    "html_last_modified": "2018-12-24T00:00:00.000Z",
    "html_encoding": "ShiftJIS",
    "html_charset": "JIS X 0208",
    "html_updated": 0,
    "authors": [
      {
        "person_id": 55,
        "last_name": "富田",
        "first_name": "倫生"
      }
    ]
  }
]
```

図書カード No. が分かっている場合は [`aozora`]`.Client.LookupBookRaw()` 関数で作品情報を取得できる。

{{< highlight go "hl_lines=2" >}}
func main() {
    b, err := aozora.DefaultClient().LookupBookRaw(59489)
    if err != nil {
        fmt.Println(err)
        return
    }
    io.Copy(os.Stdout, bytes.NewReader(b))
}
{{< /highlight >}}

[`aozora`]`.Client.SearchBooks()` または [`aozora`]`.Client.LookupBook()` 関数を使うと結果を [`aozora`]`.Book` 構造体で返す。

```go
book, err := aozora.DefaultClient().LookupBook(59489)
```

[`aozora`]`.Book` 構造体の構成は以下の通り。

```go
//Author is entity class of author and translator info.
type Author struct {
    PersonID  int    `json:"person_id"`
    LastName  string `json:"last_name"`
    FirstName string `json:"first_name"`
}

//Book is entity class of book info.
type Book struct {
    BookID                      int      `json:"book_id"`
    Title                       string   `json:"title"`
    TitleYomi                   string   `json:"title_yomi"`
    TitleSort                   string   `json:"title_sort"`
    Subtitle                    string   `json:"subtitle"`
    SubtitleYomi                string   `json:"subtitle_yomi"`
    OriginalTitle               string   `json:"original_title"`
    FirstAppearance             string   `json:"first_appearance"`
    NDCCode                     string   `json:"ndc_code"`
    FontKanaType                string   `json:"font_kana_type"`
    Copyright                   bool     `json:"copyright"`
    ReleaseDate                 Date     `json:"release_date"`
    LastModified                Date     `json:"last_modified"`
    CardURL                     string   `json:"card_url"`
    BaseBook1                   string   `json:"base_book_1"`
    BaseBookPublisher1          string   `json:"base_book_1_publisher"`
    BaseBookFirstEdition1       string   `json:"base_book_1_1st_edition"`
    BaseBookEditionInput1       string   `json:"base_book_1_edition_input"`
    BaseBookEditionProofing1    string   `json:"base_book_1_edition_proofing"`
    BaseBookParent1             string   `json:"base_book_1_parent"`
    BaseBookParentPublisher1    string   `json:"base_book_1_parent_publisher"`
    BaseBookParentFirstEdition1 string   `json:"base_book_1_parent_1st_edition"`
    BaseBook2                   string   `json:"base_book_2"`
    BaseBookPublisher2          string   `json:"base_book_2_publisher"`
    BaseBookFirstEdition2       string   `json:"base_book_2_1st_edition"`
    BaseBookEditionInput2       string   `json:"base_book_2_edition_input"`
    BaseBookEditionProofing2    string   `json:"base_book_2_edition_proofing"`
    BaseBookParent2             string   `json:"base_book_2_parent"`
    BaseBookParentPublisher2    string   `json:"base_book_2_parent_publisher"`
    BaseBookParentFirstEdition2 string   `json:"base_book_2_parent_1st_edition"`
    Input                       string   `json:"input"`
    Proofing                    string   `json:"proofing"`
    TextURL                     string   `json:"text_url"`
    TextLastModified            Date     `json:"text_last_modified"`
    TextEncoding                string   `json:"text_encoding"`
    TextCharset                 string   `json:"text_charset"`
    TextUpdated                 int      `json:"text_updated"`
    HTMLURL                     string   `json:"html_url"`
    HTMLLastModified            Date     `json:"html_last_modified"`
    HTMLEncoding                string   `json:"html_encoding"`
    HTMLCharset                 string   `json:"html_charset"`
    HTMLUpdated                 int      `json:"html_updated"`
    Translators                 []Author `json:"translators"`
    Authors                     []Author `json:"authors"`
}
```

### 作家情報の取得

今度は作家名「富田倫生」を検索キーとして[青空文庫] API から作家情報を取得してみる。

```go
func main() {
    b, err := aozora.DefaultClient().SearchPersonsRaw(
        aozora.WithPersonName("富田倫生"),
    )
    if err != nil {
        fmt.Println(err)
        return
    }
    io.Copy(os.Stdout, bytes.NewReader(b))
}
```

実行結果は以下の通り。

```text
$ go run sample.go | jq .
[
  {
    "person_id": 55,
    "last_name": "富田",
    "first_name": "倫生",
    "last_name_yomi": "とみた",
    "first_name_yomi": "みちお",
    "last_name_sort": "とみた",
    "first_name_sort": "みちお",
    "last_name_roman": "Tomita",
    "first_name_roman": "Michio",
    "date_of_birth": "1952-04-20",
    "date_of_death": "2013-08-16",
    "author_copyright": true
  }
]
```

作家 No. が分かっている場合は [`aozora`]`.Client.LookupPersonRaw()` 関数で作品情報を取得できる。

{{< highlight go "hl_lines=2" >}}
func main() {
    b, err := aozora.DefaultClient().LookupPersonRaw(55)
    if err != nil {
        fmt.Println(err)
        return
    }
    io.Copy(os.Stdout, bytes.NewReader(b))
}
{{< /highlight >}}

[`aozora`]`.Client.SearchPersons()` または [`aozora`]`.Client.LookupPerson()` 関数を使うと結果を [`aozora`]`.Person` 構造体で返す。

```go
person, err := aozora.DefaultClient().LookupPerson(55)
```

[`aozora`]`.Person` 構造体の構成は以下の通り。

```go
//Person is entity class of person info.
type Person struct {
    PersonID        int    `json:"person_id"`
    LastName        string `json:"last_name"`
    FirstName       string `json:"first_name"`
    LastNameYomi    string `json:"last_name_yomi"`
    FirstNameYomi   string `json:"first_name_yomi"`
    LastNameSort    string `json:"last_name_sort"`
    FirstNameSort   string `json:"first_name_sort"`
    LastNameRoman   string `json:"last_name_roman"`
    FirstNameRoman  string `json:"first_name_roman"`
    DateOfBirth     Date   `json:"date_of_birth"`
    DateOfDeath     Date   `json:"date_of_death"`
    AuthorCopyright bool   `json:"author_copyright"`
}
```

### 工作員情報の取得

更に更に工作員名「雪森」を検索キーとして[青空文庫] API から工作員情報を取得してみる。

```go
func main() {
    b, err := aozora.DefaultClient().SearchWorkersRaw(
        aozora.WithWorkerName("é›ªæ£®"),
    )
    if err != nil {
        fmt.Println(err)
        return
    }
    io.Copy(os.Stdout, bytes.NewReader(b))
}
```

実行結果は以下の通り。

```text
$ go run sample.go | jq .
[
  {
    "id": 845,
    "name": "雪森"
  }
]
```

工作員 No. が分かっている場合は [`aozora`]`.Client.LookupWorkerRaw()` 関数で作品情報を取得できる。

{{< highlight go "hl_lines=2" >}}
func main() {
    b, err := aozora.DefaultClient().LookupWorkerRaw(845)
    if err != nil {
        fmt.Println(err)
        return
    }
    io.Copy(os.Stdout, bytes.NewReader(b))
}
{{< /highlight >}}

[`aozora`]`.Client.SearchWorkers()` または [`aozora`]`.Client.LookupWorker()` 関数を使うと結果を [`aozora`]`.Worker` 構造体で返す。

```go
worker, err := aozora.DefaultClient().LookupWorker(845)
```

[`aozora`]`.Worker` 構造体の構成は以下の通り。

```go
//Worker is entity class of worker info.
type Worker struct {
    WorkerID int    `json:"id"`
    Name     string `json:"name"`
}
```

## Server および Client インスタンスの生成

[`aozora`]`.Client` インスタンスの生成は [`aozora`]`.DefaultClient()` 関数で簡単に行えるが，もう少し細かい制御もできる。

### [青空文庫] API サーバを指定する

[`aozora`]`.New()` 関数で [`aozora`]`.Server` インスタンスを生成できるが，引数としてサーバを指定できる。

```go
server := aozora.New(
    aozora.WithScheme("http"),
    aozora.WithServerName("pubserver2.herokuapp.com"),
)
```

これで[青空文庫] API サーバとして `http://pubserver2.herokuapp.com` を指定できた。

### http.Client を指定する

[`aozora`]`.Server.CreateClient()` 関数により [`aozora`]`.Client` インスタンスを生成できるが，引数として `http.Client` インスタンスを指定できる。

```go {hl_lines=[4]}
client := aozora.New(
    aozora.WithScheme("http"),
    aozora.WithServerName("pubserver2.herokuapp.com"),
).CreateClient(aozora.WithHttpClient(&http.Client{}))
```

ちなみに [`aozora`]`.DefaultClient()` 関数は以下の記述と同等である。

```go
client := aozora.New(
    aozora.WithScheme("http"),
    aozora.WithServerName("www.aozorahack.net"),
).CreateClient(aozora.WithHttpClient(&http.Client{}))
```

## context.Context 付きのアクセス

[青空文庫] API アクセスの際に [`context`]`.Context` を付けることができる。

### 作品情報の取得

```go
rawjson, err := aozora.DefaultClient().SearchBooksRawContext(
    context.Background(),
    aozora.WithBookTitle("/天に積む宝/"),
    aozora.WithBookAuthor("富田倫生"),
)
```

```go
books, err := aozora.DefaultClient().SearchBooksContext(
    context.Background(),
    aozora.WithBookTitle("/天に積む宝/"),
    aozora.WithBookAuthor("富田倫生"),
)
```

```go
rawjson, err := aozora.DefaultClient().LookupBookRawContext(context.Background(), 59489)
```

```go
book, err := aozora.DefaultClient().LookupBookContext(context.Background(), 59489)
```

### 作家情報の取得

```go
rawjson, err := aozora.DefaultClient().SearchPersonsRawContext(
    context.Background(),
    aozora.WithPersonName("富田倫生"),
)
```

```go
persons, err := aozora.DefaultClient().SearchPersonsContext(
    context.Background(),
    aozora.WithPersonName("富田倫生"),
)
```

```go
rawjson, err := aozora.DefaultClient().LookupPersonRawContext(context.Background(), 55)
```

```go
person, err := aozora.DefaultClient().LookupPersonContext(context.Background(), 55)
```

### 工作員情報の取得

```go
rawjson, err := aozora.DefaultClient().SearchWorkersRawContext(
    context.Background(),
    aozora.WithWorkerName("é›ªæ£®"),
)
```

```go
workers, err := aozora.DefaultClient().SearchWorkersContext(
    context.Background(),
    aozora.WithWorkerName("é›ªæ£®"),
)
```

```go
rawjson, err := aozora.DefaultClient().LookupWorkerRawContext(context.Background(), 845)
```

```go
worker, err := aozora.DefaultClient().LookupWorkerContext(context.Background(), 845)
```

## ブックマーク

- [#aozorahack に関する覚え書き]({{< ref "/remark/2019/08/about-aozorahack.md" >}})

[Go]: https://go.dev/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[青空文庫]: https://www.aozora.gr.jp/ "青空文庫　Aozora Bunko"
[`aozora`]: https://github.com/goark/aozora-api "goark/aozora-api: APIs for Aozora-bunko RESTful Service by Golang"
[`http`]: https://golang.org/pkg/net/http/ "http - The Go Programming Language"
[`context`]: https://golang.org/pkg/context/ "context - The Go Programming Language"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4899840721" %}} <!-- インターネット図書館 青空文庫 -->
{{% review-paapi "4756117074" %}} <!-- 本の未来 -->
