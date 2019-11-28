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

- [spiegel-im-spiegel/aozora-api: APIs for Aozora-bunko RESTful Service by Golang](https://github.com/spiegel-im-spiegel/aozora-api)

本パッケージは[青空文庫] API へアクセスできる [Go 言語]用クライアント・パッケージだ。
API を通じて[青空文庫]に収録されている作品情報等を取得できる。

なお [spiegel-im-spiegel/aozora-api] パッケージは [Go] 1.13 以上を要求する。
ご注意を。

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/aozora-api.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/aozora-api)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/aozora-api/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/spiegel-im-spiegel/aozora-api.svg)](https://github.com/spiegel-im-spiegel/aozora-api/releases/latest)

## インポート

```go
import "github.com/spiegel-im-spiegel/aozora-api"
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

    "github.com/spiegel-im-spiegel/aozora-api"
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
これらの関数を[`aozora`].`Client.SearchBooksRaw()` 関数の引数に指定する。

| オプション名 | 対応する関数                                  |
| ------------:| --------------------------------------------- |
|      `title` | [`aozora`].`WithBookTitle(string)`            |
|     `author` | [`aozora`].`WithBookAuthor(string)`           |
|     `fields` | [`aozora`].`WithBookFields(string)`[^fields1] |
|      `limit` | [`aozora`].`WithBookLimit(int)`               |
|       `skip` | [`aozora`].`WithBookSkip(int)`                |
|      `after` | [`aozora`].`WithBookAfter(time.Time)`         |

[^fields1]: `fields` オプションは効いてない感じ？

[`aozora`].`Client.SearchBooksRaw()` 関数は API を通じて取得した結果（JSON 形式）をそのまま返す。

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

図書カード No. が分かっている場合は [`aozora`].`Client.LookupBookRaw()` 関数で作品情報を取得できる。

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

[`aozora`].`Client.SearchBooks()` または [`aozora`].`Client.LookupBook()` 関数を使うと結果を [`aozora`].`Book` 構造体で返す。

```go
book, err := aozora.DefaultClient().LookupBook(59489)
```

[`aozora`].`Book` 構造体の構成は以下の通り。

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

作家 No. が分かっている場合は [`aozora`].`Client.LookupPersonRaw()` 関数で作品情報を取得できる。

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

[`aozora`].`Client.SearchPersons()` または [`aozora`].`Client.LookupPerson()` 関数を使うと結果を [`aozora`].`Person` 構造体で返す。

```go
person, err := aozora.DefaultClient().LookupPerson(55)
```

[`aozora`].`Person` 構造体の構成は以下の通り。

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

工作員 No. が分かっている場合は [`aozora`].`Client.LookupWorkerRaw()` 関数で作品情報を取得できる。

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

[`aozora`].`Client.SearchWorkers()` または [`aozora`].`Client.LookupWorker()` 関数を使うと結果を [`aozora`].`Worker` 構造体で返す。

```go
person, err := aozora.DefaultClient().LookupWorker(845)
```

[`aozora`].`Worker` 構造体の構成は以下の通り。

```go
//Worker is entity class of worker info.
type Worker struct {
    WorkerID int    `json:"id"`
    Name     string `json:"name"`
}
```

## Server および Client インスタンスの生成

[`aozora`].`Client` インスタンスの生成は [`aozora`].`DefaultClient()` 関数で簡単に行えるが，もう少し細かい制御もできる。

### [青空文庫] API サーバを指定する

[`aozora`].`New()` 関数で [`aozora`].`Server` インスタンスを生成できるが，引数としてサーバを指定できる。

```go
server := aozora.New(
    aozora.WithScheme("http"),
    aozora.WithServerName("pubserver2.herokuapp.com"),
)
```

これで[青空文庫] API サーバとして `http://pubserver2.herokuapp.com` を指定できた。

### context.Context および http.Client を指定する

[`aozora`].`Server.CreateClient()` 関数により [`aozora`].`Client` インスタンスを生成できるが，引数として `context.Context` および `http.Client` インスタンスを指定する。

```go
client := aozora.New(
    aozora.WithScheme("http"),
    aozora.WithServerName("pubserver2.herokuapp.com"),
).CreateClient(
    context.Background(),
    &http.Client{},
)
```

ちなみに [`aozora`].`DefaultClient()` 関数は以下の記述と同等である。

```go
client := aozora.New(
    aozora.WithScheme("http"),
    aozora.WithServerName("www.aozorahack.net"),
).CreateClient(
    context.Background(),
    http.DefaultClient,
)
```

## ブックマーク

- [#aozorahack に関する覚え書き]({{< ref "/remark/2019/08/about-aozorahack.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[青空文庫]: https://www.aozora.gr.jp/ "青空文庫　Aozora Bunko"
[spiegel-im-spiegel/aozora-api]: https://github.com/spiegel-im-spiegel/aozora-api "spiegel-im-spiegel/aozora-api: APIs for Aozora-bunko RESTful Service by Golang"
[`aozora`]: https://github.com/spiegel-im-spiegel/aozora-api "spiegel-im-spiegel/aozora-api: APIs for Aozora-bunko RESTful Service by Golang"

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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%BC%E3%83%8D%E3%83%83%E3%83%88%E5%9B%B3%E6%9B%B8%E9%A4%A8-%E9%9D%92%E7%A9%BA%E6%96%87%E5%BA%AB-%E9%87%8E%E5%8F%A3-%E8%8B%B1%E5%8F%B8/dp/4899840721?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4899840721"><img src="https://images-fe.ssl-images-amazon.com/images/I/51V8S7TXJ5L._SL160_.jpg" width="112" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%BC%E3%83%8D%E3%83%83%E3%83%88%E5%9B%B3%E6%9B%B8%E9%A4%A8-%E9%9D%92%E7%A9%BA%E6%96%87%E5%BA%AB-%E9%87%8E%E5%8F%A3-%E8%8B%B1%E5%8F%B8/dp/4899840721?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4899840721">インターネット図書館 青空文庫</a></dt>
    <dd>野口 英司</dd>
    <dd>はる書房 2005-11-01</dd>
    <dd>単行本</dd>
    <dd>4899840721 (ASIN), 9784899840725 (EAN)</dd>
    <dd>Rating<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">青空文庫の活動について紹介。作品を収録した DVD も付いてる！ 巻末に載っている<a href="https://www.tomita-michio.jp/">富田倫生</a>さんの文章は<a href="https://www.aozora.gr.jp/cards/000055/card59489.html">青空文庫に収録</a>されている。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-01-02">2019-01-02</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home" >PA-API</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%9C%AC%E3%81%AE%E6%9C%AA%E6%9D%A5-Ascii-books-%E5%AF%8C%E7%94%B0-%E5%80%AB%E7%94%9F/dp/4756117074?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4756117074"><img src="https://images-fe.ssl-images-amazon.com/images/I/5131GA04AHL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%9C%AC%E3%81%AE%E6%9C%AA%E6%9D%A5-Ascii-books-%E5%AF%8C%E7%94%B0-%E5%80%AB%E7%94%9F/dp/4756117074?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4756117074">本の未来 (Ascii books)</a></dt>
    <dd>富田 倫生</dd>
    <dd>アスキー</dd>
    <dd>単行本</dd>
    <dd>4756117074 (ASIN), 9784756117076 (EAN)</dd>
    <dd>Rating<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">e-book の未来を予見する試みの書。あるいは本とコンピュータの関係について。<a href="https://www.aozora.gr.jp/cards/000055/card56499.html">青空文庫にも収録</a>されている。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-01-02">2019-01-02</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home" >PA-API</a>)</p>
</div>
