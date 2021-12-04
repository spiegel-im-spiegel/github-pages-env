+++
title = "書籍データ取得ツール books-data"
date =  "2019-09-08T15:22:41+09:00"
description = "本ツールは PA-API, openBD および 青空文庫 API より書籍情報を取得・加工するための CLI ツールである。"
image = "/images/attention/tools.png"
tags = [ "tools", "book", "books-data", "pa-api", "openbd-api", "aozora-api" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

- [spiegel-im-spiegel/books-data: Search for Books Data](https://github.com/spiegel-im-spiegel/books-data)

本ツールは [PA-API], [openBD] および [青空文庫] API より書籍情報を取得・加工するための CLI (Command-Line Interface) ツールである。

[![check vulns](https://github.com/spiegel-im-spiegel/books-data/workflows/vulns/badge.svg)](https://github.com/spiegel-im-spiegel/books-data/actions)
[![lint status](https://github.com/spiegel-im-spiegel/books-data/workflows/lint/badge.svg)](https://github.com/spiegel-im-spiegel/books-data/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/books-data/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/spiegel-im-spiegel/books-data.svg)](https://github.com/spiegel-im-spiegel/books-data/releases/latest)

## ダウンロードとビルド

[books-data] は以下の [Go] コマンドでダウンロードとビルドができる。

```text
$ go get github.com/spiegel-im-spiegel/books-data@latest
```

なおビルドには [Go] 1.13 以上が必要になる。
ご注意を。

各プラットフォーム用のバイナリも用意している。
[最新バイナリはリリースページから取得](https://github.com/spiegel-im-spiegel/books-data/releases/latest)できる。

## 簡単な使い方

`-h` オプションで簡単なヘルプを表示できる。

```text
$ books-data -h
Search for books data

Usage:
  books-data [flags]
  books-data [command]

Available Commands:
  help        Help about any command
  history     Lookup review data from history log
  review      Make review data
  search      Search for books data
  version     Print the version number

Flags:
      --access-key string      Config: PA-API Access Key ID
  -c, --aozora-card string     Aozora-bunko card no.
  -a, --asin string            Amazon ASIN code
      --associate-tag string   Config: PA-API Associate Tag
      --config string          Config file (default /home/username/.config/books-data/config.yaml)
      --debug                  for debug
  -h, --help                   help for books-data
  -i, --isbn string            ISBN code
      --marketplace string     Config: PA-API Marketplace (default "www.amazon.co.jp")
  -l, --review-log string      Config: Review log file (JSON format)
      --secret-key string      Config: PA-API Secret Access Key
  -t, --template-file string   Template file for formatted output

Use "books-data [command] --help" for more information about a command.
```

### 設定ファイル

[books-data] のオプションの一部は設定ファイルであらかじめ指定しておくことができる。
指定可能なオプションは以下の通り

| オプション名    | 既定値             | 内容                                        |
| --------------- | ------------------ | ------------------------------------------- |
| `marketplace`   | `www.amazon.co.jp` | [PA-API] サーバ名                           |
| `associate-tag` | なし               | [PA-API] アソシエイト・タグ（アカウント名） |
| `access-key`    | なし               | [PA-API] アクセスキー                       |
| `secret-key`    | なし               | [PA-API] 秘密キー                           |
| `review-log`    | なし               | レビューデータ作成時の保存ログファイル名    |

設定ファイルの書式は YAML 形式で以下のような感じに記述する。

```text
$ cat ~/.config/books-data/config.yaml
marketplace: www.amazon.co.jp
associate-tag: mytag-20
access-key: AKIAIOSFODNN7EXAMPLE
secret-key: 1234567890
review-log: /home/username/.local/share/books-data/review-log.json
```

設定ファイルは以下のパスに置いておけば [books-data] 起動時に自動的に読み込まれる。

- Windows
    - `%AppData%\books-data\config.yaml`
- Darwin (macOS)
    - `$HOME/Library/Application Support/books-data/config.yaml`
- Plan 9
    - `$home/lib/books-data/config.yaml`
- UNIX 系 OS (Linux 他)
    - `$XDG_CONFIG_HOME/books-data/config.yaml` （`$XDG_CONFIG_HOME` 環境変数が指定されている場合）
    - `$HOME/.config/books-data/config.yaml`

また [books-data] 起動時に `--config` オプションで設定ファイルを直接指定することも可能だ。
設定ファイルの内容には秘密情報が含まれる（[PA-API] を利用する場合）ためファイルのアクセス設定に注意すること。

### 書籍情報の検索

```text
$ books-data search -h
Search for books data

Usage:
  books-data search [flags]

Flags:
  -h, --help   help for search
      --raw    Output raw data from API

Global Flags:
      --access-key string      Config: PA-API Access Key ID
  -c, --aozora-card string     Aozora-bunko card no.
  -a, --asin string            Amazon ASIN code
      --associate-tag string   Config: PA-API Associate Tag
      --config string          Config file (default /home/username/.config/books-data/config.yaml)
      --debug                  for debug
  -i, --isbn string            ISBN code
      --marketplace string     Config: PA-API Marketplace (default "www.amazon.co.jp")
  -l, --review-log string      Config: Review log file (JSON format)
      --secret-key string      Config: PA-API Secret Access Key
  -t, --template-file string   Template file for formatted output
```

たとえば ASIN コード `B07TYKJQFK` のアイテムの情報を取得する場合は以下のコマンドラインで起動する（[PA-API] を利用可能な場合）。

```text
$ books-data search -a B07TYKJQFK
{"Type":"paapi","ID":"B07TYKJQFK","Title":"数学ガールの秘密ノート／ビットとバイナリー","URL":"https://www.amazon.co.jp/dp/B07TYKJQFK?tag=mytag-20\u0026linkCode=ogi\u0026th=1\u0026psc=1","Image":{"URL":"https://m.media-amazon.com/images/I/41Q0Hyr8g3L._SL160_.jpg","Height":160,"Width":111},"ProductType":"Kindle版","Creators":[{"Name":"結城 浩","Role":"著"}],"Publisher":"SBクリエイティブ","Codes":[{"Name":"ASIN","Value":"B07TYKJQFK"}],"PublicationDate":"2019-07-19","LastRelease":"2019-07-20","Service":{"Name":"PA-APIv5","URL":"https://affiliate.amazon.co.jp/assoc_credentials/home"}}
```

実行結果は JSON 形式で出力される。

他に指定可能な検索オプションは以下の通り。

| オプション          | 内容                     |
| ------------------- | ------------------------ |
| `-a, --asin`        | Amazon ASIN コード       |
| `-i, --isbn`        | ISBN 番号                |
| `-c, --aozora-card` | [青空文庫]図書カード No. |

これらは同時に指定可能だが Amazon ASIN コード → ISBN 番号 → [青空文庫]図書カード No. の順で検索を行う。
検索キーとして ISBN コードを指定した場合は [openBD] による検索を行う。

### レビュー・データの作成

```text
$ books-data review -h
Make review data

Usage:
  books-data review [flags] [description]

Flags:
      --bookpage-url string   URL of book page
  -h, --help                  help for review
      --image-url string      URL of book cover image
      --pipe                  Import description from Stdin
  -r, --rating int            Rating of product
      --review-date string    Date of review

Global Flags:
      --access-key string      Config: PA-API Access Key ID
  -c, --aozora-card string     Aozora-bunko card no.
  -a, --asin string            Amazon ASIN code
      --associate-tag string   Config: PA-API Associate Tag
      --config string          Config file (default /home/username/.config/books-data/config.yaml)
      --debug                  for debug
  -i, --isbn string            ISBN code
      --marketplace string     Config: PA-API Marketplace (default "www.amazon.co.jp")
  -l, --review-log string      Config: Review log file (JSON format)
      --secret-key string      Config: PA-API Secret Access Key
  -t, --template-file string   Template file for formatted output
```

ASIN, ISBN または [青空文庫]図書カード No. で指定した書籍情報に対してレビュー情報を作成する。
レビュー情報として以下のオプションを指定できる。

| オプション       | 内容                                    |
| ---------------- | --------------------------------------- |
| `--bookpage-url` | 書籍ページへの URL                      |
| `--image-url`    | 書影（画像データ）への URL              |
| `--rating`       | 評価ポイント (`0-5`) ※ `0`なら評価なし |
| `--review-date`  | レビュー日                              |

コマンドラインはこんな感じで記述する。

```text
$ books-data review -a B07TYKJQFK -r 5 --review-date 2019-09-08 "面白かった！"
{"Book":{"Type":"paapi","ID":"B07TYKJQFK","Title":"数学ガールの秘密ノート／ビットとバイナリー","URL":"https://www.amazon.co.jp/dp/B07TYKJQFK?tag=mytag-20\u0026linkCode=ogi\u0026th=1\u0026psc=1","Image":{"URL":"https://m.media-amazon.com/images/I/41Q0Hyr8g3L._SL160_.jpg","Height":160,"Width":111},"ProductType":"Kindle版","Creators":[{"Name":"結城 浩","Role":"著"}],"Publisher":"SBクリエイティブ","Codes":[{"Name":"ASIN","Value":"B07TYKJQFK"}],"PublicationDate":"2019-07-19","LastRelease":"2019-07-20","Service":{"Name":"PA-APIv5","URL":"https://affiliate.amazon.co.jp/assoc_credentials/home"}},"Date":"2019-09-08","Rating":5,"Star":[true,true,true,true,true],"Description":"面白かった！"}
```

実行結果は JSON 形式で出力される。
また `--pipe` オプションを付けることで標準入力から description を入力できる。

```text
$ echo "面白かった！" | books-data review -a B07TYKJQFK -r 5 --review-date 2019-09-08 --pipe
{"Book":{"Type":"paapi","ID":"B07TYKJQFK","Title":"数学ガールの秘密ノート／ビットとバイナリー","URL":"https://www.amazon.co.jp/dp/B07TYKJQFK?tag=mytag-20\u0026linkCode=ogi\u0026th=1\u0026psc=1","Image":{"URL":"https://m.media-amazon.com/images/I/41Q0Hyr8g3L._SL160_.jpg","Height":160,"Width":111},"ProductType":"Kindle版","Creators":[{"Name":"結城 浩","Role":"著"}],"Publisher":"SBクリエイティブ","Codes":[{"Name":"ASIN","Value":"B07TYKJQFK"}],"PublicationDate":"2019-07-19","LastRelease":"2019-07-20","Service":{"Name":"PA-APIv5","URL":"https://affiliate.amazon.co.jp/assoc_credentials/home"}},"Date":"2019-09-08","Rating":5,"Star":[true,true,true,true,true],"Description":"面白かった！\n"}
```

### レビュー履歴の参照

```text
$ books-data history -h
Lookup review data from history log

Usage:
  books-data history [flags]

Flags:
  -h, --help   help for history

Global Flags:
      --access-key string      Config: PA-API Access Key ID
  -c, --aozora-card string     Aozora-bunko card no.
  -a, --asin string            Amazon ASIN code
      --associate-tag string   Config: PA-API Associate Tag
      --config string          Config file (default /home/spiegel/.config/books-data/config.yaml)
      --debug                  for debug
  -i, --isbn string            ISBN code
      --marketplace string     Config: PA-API Marketplace (default "www.amazon.co.jp")
  -l, --review-log string      Config: Review log file (JSON format)
      --secret-key string      Config: PA-API Secret Access Key
  -t, --template-file string   Template file for formatted output
```

`--review-log` オプションでレビュー結果を保存している場合は `history` コマンドで過去の履歴を呼び出せる。

```text
$ books-data history -a B07TYKJQFK
{"Book":{"Type":"paapi","ID":"B07TYKJQFK","Title":"数学ガールの秘密ノート／ビットとバイナリー","URL":"https://www.amazon.co.jp/dp/B07TYKJQFK?tag=mytag-20\u0026linkCode=ogi\u0026th=1\u0026psc=1","Image":{"URL":"https://m.media-amazon.com/images/I/41Q0Hyr8g3L._SL160_.jpg","Height":160,"Width":111},"ProductType":"Kindle版","Creators":[{"Name":"結城 浩","Role":"著"}],"Publisher":"SBクリエイティブ","Codes":[{"Name":"ASIN","Value":"B07TYKJQFK"}],"PublicationDate":"2019-07-19","LastRelease":"2019-07-20","Service":{"Name":"PA-APIv5","URL":"https://affiliate.amazon.co.jp/assoc_credentials/home"}},"Date":"2019-09-08","Rating":5,"Star":[true,true,true,true,true],"Description":"面白かった！\n"}
```

### 出力の整形

`-t` または `--template-file` オプションででテンプレートファイルを指定することにより `search`, `review`, `history` 各コマンドの実行結果の出力を整形できる。

たとえば以下のようなテンプレートファイルを用意すれば

```text
$ cat reviews/template.bib.txt 
@BOOK{Book:{{ .Book.ID }},
    TITLE = "{{ .Book.Title }}",
    AUTHOR = "{{ with .Book.Creators }}{{ range $i, $v := . }}{{ if ne $i 0 }} and {{ end }}{{ . }}{{ end }}{{ end }}"{{ if .Book.Publisher }},
    PUBLISHER = {{ "{" }}{{ .Book.Publisher }}{{ "}" }}{{ end }}{{ if gt .Book.PublicationDate.Year 1 }},
    YEAR = {{ .Book.PublicationDate.Year }}{{ end }}
}
```

このような出力になる。

```text
$ books-data history -a B07TYKJQFK -t reviews/template.bib.txt 
@BOOK{Book:B07TYKJQFK,
    TITLE = "数学ガールの秘密ノート／ビットとバイナリー",
    AUTHOR = "結城 浩 (著)",
    PUBLISHER = {SBクリエイティブ},
    YEAR = 2019
}
```

`search` コマンドで取得した書籍データのフォーマットは以下の通り。

```go
//Book is entity class of information for book
type Book struct {
    Type            string
    ID              string
    Title           string
    SubTitle        string `json:",omitempty"`
    SeriesTitle     string `json:",omitempty"`
    OriginalTitle   string `json:",omitempty"`
    URL             string `json:",omitempty"`
    Image           BookCover
    ProductType     string    `json:",omitempty"`
    Creators        []Creator `json:",omitempty"`
    Publisher       string    `json:",omitempty"`
    Codes           []Code
    PublicationDate values.Date
    LastRelease     values.Date
    PublicDomain    bool   `json:",omitempty"`
    FirstAppearance string `json:",omitempty"`
    Service         Service
}

//Code is entity class of book code
type Code struct {
    Name  string
    Value string
}

//Creator is entity class of creator info.
type Creator struct {
    Name string
    Role string `json:",omitempty"`
}

//BookCover is entity class of book cover image info.
type BookCover struct {
    URL    string
    Height uint16 `json:",omitempty"`
    Width  uint16 `json:",omitempty"`
}

//Service is entity class of API service info.
type Service struct {
    Name string
    URL  string
}
```

また `review`, `history` コマンドで取得したレビュー・データのフォーマットは以下の通り。

```go
//Review is entity class for review info.
type Review struct {
    Book        *entity.Book
    Date        values.Date
    Rating      int
    Star        [MAX_STAR]bool
    Description string `json:",omitempty"`
}
```

ちなみに，このブログにおけるレビューカードのテンプレートは以下の内容になっている。

```text
<div class="hreview">{{ if .Book.Image.URL }}
  <div class="photo">{{ if .Book.URL }}<a class="item url" href="{{ .Book.URL }}">{{ end }}<img src="{{ .Book.Image.URL }}" width="{{ with .Book.Image.Width }}{{ . }}{{ else }}110{{ end }}" alt="photo">{{ if .Book.URL }}</a>{{ end }}</div>{{ end }}
  <dl class="fn">
    <dt>{{ if .Book.URL }}<a href="{{ .Book.URL }}">{{ end }}{{ .Book.Title }}{{ with .Book.SubTitle }} {{ . }}{{ end }}{{ with .Book.SeriesTitle }} ({{ . }}){{ end }}{{ if .Book.URL }}</a>{{ end }}</dt>{{ if .Book.OriginalTitle }}
    <dd>原題: {{ .Book.OriginalTitle }}</dd>{{ end }}{{ if .Book.FirstAppearance }}
    <dd>（初出: {{ .Book.FirstAppearance }}）</dd>{{ end }}{{ if .Book.Creators }}
    <dd>{{ range $i, $v := .Book.Creators }}{{ if ne $i 0 }}, {{ end }}{{ $v }}{{ end }}</dd>{{ end }}
    <dd>{{ .Book.Publisher }}{{ if not .Book.PublicationDate.IsZero }} {{ .Book.PublicationDate }}{{ end }}{{ if not .Book.LastRelease.IsZero }} (Release {{ .Book.LastRelease }}){{ end }}</dd>
    <dd>{{ .Book.ProductType }}{{ if .Book.PublicDomain }} (Public Domain){{ end }}</dd>{{ if .Book.Codes }}
    <dd>{{ range $i, $v := .Book.Codes }}{{ if ne $i 0 }}, {{ end }}{{ $v }}{{ end }}</dd>{{ end }}{{ if gt .Rating 0 }}
    <dd>評価<abbr class="rating fa-sm" title="{{ .Rating }}">{{ range .Star }}&nbsp;{{ if . }}<i class="fas fa-star"></i>{{ else }}<i class="far fa-star"></i>{{ end }}{{ end }}</abbr></dd>{{ end }}
  </dl>
  <p class="description">{{ .Description }}</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="{{ .Date }}">{{ .Date }}</abbr> (powered by {{ if .Book.Service.URL }}<a href="{{ .Book.Service.URL }}">{{ end }}{{ .Book.Service.Name }}{{ if .Book.Service.URL }}</a>{{ end }})</p>
</div>
```

テンプレート・ファイルを作成する際の参考にどうぞ。

## 参考情報

- [Go 言語用 CLI プログラミング支援パッケージ]({{< ref "/release/gocli-package-for-golang.md" >}})
- [Go 言語用エラーハンドリング・パッケージ]({{< ref "/release/errs-package-for-golang.md" >}})
- [Go 言語用青空文庫 API クライアント・パッケージ]({{< ref "/release/aozora-api-package-for-golang.md" >}})
- [Go 言語用 openBD クライアント・パッケージ]({{< ref "/release/openbd-api-package-for-golang.md" >}})
- [Go 言語用 PA-API v5 クライアント・パッケージ]({{< ref "/release/pa-api-v5.md" >}})

[Go]: https://go.dev/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[books-data]: https://github.com/spiegel-im-spiegel/books-data "spiegel-im-spiegel/books-data: Search for Books Data"
[PA-API]: https://affiliate.amazon.co.jp/assoc_credentials/home "Product Advertising API"
[openBD]: https://openbd.jp/ "openBD | 書誌情報・書影を自由に"
[青空文庫]: https://www.aozora.gr.jp/ "青空文庫　Aozora Bunko"
