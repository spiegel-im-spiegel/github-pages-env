+++
title = "Hugo で簡易データベースを構成する"
date =  "2020-01-04T22:53:33+09:00"
description = "拙作 books-data にはレビュー結果を JSON 形式のログファイルに出力する機能があるので，そのファイルを Hugo の data ディレクトリに置いて Hugo 環境で利用する。"
image = "/images/attention/kitten.jpg"
tags = [ "hugo", "books-data", "book", "site" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Hugo] と拙作「[書籍データ取得ツール books-data]({{< ref "/release/books-data.md" >}})」を組み合わせて書籍等のレビュー・データを格納する簡易データベースを構成することを考える。

つっても難しい話じゃ全然なくて， [books-data] にはレビュー結果を JSON 形式のログファイルに出力する機能があるので，そのファイルを [Hugo] の `data` ディレクトリに置いて [Hugo] 環境で利用するだけの話。
随分前からそうしようと思ってたんだけど，一度構築した環境を変えるのはなかなかに面倒な気がして踏ん切りがつかなかったのよ。

## [books-data] のレビュー結果を data ディレクトリに格納する

拙作 [books-data] の `review` コマンドの使い方は以下の通り（[books-data] の詳しい使い方は「[書籍データ取得ツール books-data]({{< ref "/release/books-data.md" >}})」を参考にどうぞ）。

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
      --config string          Config file (default /home/spiegel/.config/books-data/config.yaml)
      --debug                  for debug
  -i, --isbn string            ISBN code
      --marketplace string     Config: PA-API Marketplace (default "www.amazon.co.jp")
  -l, --review-log string      Config: Review log file (JSON format)
      --secret-key string      Config: PA-API Secret Access Key
  -t, --template-file string   Template file for formatted output
```

`review` コマンドに `-l` オプションを付けるか設定ファイルに以下のように `review-log` を設定することでレビュー結果をファイルに格納できる。

```yaml {hl_lines=[5]}
marketplace: www.amazon.co.jp
associate-tag: mytag-20
access-key: AKIAIOSFODNN7EXAMPLE
secret-key: 1234567890
review-log: /home/username/hugo-work/data/reviews.json
```

ここでは [Hugo] 環境の `data` ディレクトリに `reviews.json` のファイル名でレビュー結果を格納している。

ちなみに `reviews.json` ファイルの中身は，例えばこんな感じになる。

```json
[
  {
    "Book": {
      "Type": "paapi",
      "ID": "4621300253",
      "Title": "プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)",
      "URL": "https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22\u0026linkCode=ogi\u0026th=1\u0026psc=1",
      "Image": {
        "URL": "https://m.media-amazon.com/images/I/41meaSLNFfL._SL160_.jpg",
        "Height": 160,
        "Width": 123
      },
      "ProductType": "単行本（ソフトカバー）",
      "Creators": [
        {
          "Name": "Alan A.A. Donovan",
          "Role": "著"
        },
        {
          "Name": "Brian W. Kernighan",
          "Role": "著"
        },
        {
          "Name": "柴田 芳樹",
          "Role": "翻訳"
        }
      ],
      "Publisher": "丸善出版",
      "Codes": [
        {
          "Name": "ASIN",
          "Value": "4621300253"
        },
        {
          "Name": "EAN",
          "Value": "9784621300251"
        },
        {
          "Name": "ISBN",
          "Value": "4621300253"
        },
        {
          "Name": "ISBN",
          "Value": "9784621300251"
        }
      ],
      "PublicationDate": "2016-06-20",
      "LastRelease": "",
      "Service": {
        "Name": "PA-APIv5",
        "URL": "https://affiliate.amazon.co.jp/assoc_credentials/home"
      }
    },
    "Date": "2018-10-20",
    "Rating": 5,
    "Star": [
      true,
      true,
      true,
      true,
      true
    ],
    "Description": "著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K\u0026amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。"
  }
]
```

## レビュー結果の一覧を表示する

たとえば [Hugo] のテンプレートファイルでアイテムの名前の一覧を表示したいのであれば，こんな感じに記述すれば

```html
<ul>{{ range .Site.Data.reviews }}
  <li>{{ .Book.Title }}</li>
{{ end }}</ul>
```

以下のような出力を得られる。

```html
<ul>
  <li>プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</li>
</ul>
```

更に

```text
{{ range .Site.Data.reviews }}
  {{ partial "review-box.html" . }}
{{ end }}
```

として，テンプレート・ファイル `layouts/partials/review-box.html` の中身を

```html
<div class="hreview">{{ if .Book.Image.URL }}
  <div class="photo">{{ if .Book.URL }}<a class="item url" href="{{ .Book.URL }}">{{ end }}<img src="{{ .Book.Image.URL }}" width="{{ with .Book.Image.Width }}{{ . }}{{ else }}110{{ end }}" alt="photo">{{ if .Book.URL }}</a>{{ end }}</div>{{ end }}
  <dl class="fn">
    <dt>{{ if .Book.URL }}<a href="{{ .Book.URL }}">{{ end }}{{ .Book.Title }}{{ with .Book.SubTitle }} {{ . }}{{ end }}{{ with .Book.SeriesTitle }} ({{ . }}){{ end }}{{ if .Book.URL }}</a>{{ end }}</dt>{{ if .Book.OriginalTitle }}
    <dd>原題: {{ .Book.OriginalTitle }}</dd>{{ end }}{{ if .Book.FirstAppearance }}
    <dd>（初出: {{ .Book.FirstAppearance }}）</dd>{{ end }}{{ if .Book.Creators }}
    <dd>{{ range $i, $v := .Book.Creators }}{{ if ne $i 0 }}, {{ end }}{{ $v.Name }} ({{ $v.Role }}){{ end }}</dd>{{ end }}
    <dd>{{ if .Book.Publisher }}{{ .Book.Publisher }}{{ end }}{{ if .Book.PublicationDate }} {{ .Book.PublicationDate }}{{ end }}{{ if .Book.LastRelease }} (Release {{ .Book.LastRelease }}){{ end }}</dd>
    <dd>{{ .Book.ProductType }}{{ if .Book.PublicDomain }} (Public Domain){{ end }}</dd>{{ if .Book.Codes }}
    <dd>{{ range $i, $v := .Book.Codes }}{{ if ne $i 0 }}, {{ end }}{{ $v.Value }} ({{ $v.Name }}){{ end }}</dd>{{ end }}{{ if gt .Rating 0 }}
    <dd>評価<abbr class="rating fa-sm" title="{{ .Rating }}">{{ range .Star }}&nbsp;{{ if . }}<i class="fas fa-star"></i>{{ else }}<i class="far fa-star"></i>{{ end }}{{ end }}</abbr></dd>{{ end }}
  </dl>
  <p class="description">{{ .Description | safeHTML }}</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="{{ .Date }}">{{ .Date }}</abbr> (powered by {{ if .Book.Service.URL }}<a href="{{ .Book.Service.URL }}">{{ end }}{{ .Book.Service.Name }}{{ if .Book.Service.URL }}</a>{{ end }})</p>
</div>
```

とすれば[レビュー・カードの一覧]({{< ref "/reviews.md" >}} "Reviews")ができる。

## アイテムを指定してレビュー結果を表示する

[books-data] の `review` コマンドで出力されたレビュー結果データは `.Book.Type` および `.Book.ID` 要素で一意に識別できる。
そこで，

```text
{{ range where (where $.Site.Data.reviews "Book.ID" "4621300253") "Book.Type" "paapi" }}
  {{ partial "review-box.html" . }}
{{ end }}
```

とすれば

```html
<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&amp;linkCode=ogi&amp;th=1&amp;psc=1"><img src="https://m.media-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&amp;linkCode=ogi&amp;th=1&amp;psc=1">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN), 4621300253 (ISBN), 9784621300251 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
```

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->

のようにアイテムを指定して表示することができる。

これをもう少し一般化して shortcode として `layouts/shortcodes/review-paapi.html` を

```text
{{ $id := index .Params 0 }}
{{ range where (where $.Site.Data.reviews "Book.ID" $id) "Book.Type" "paapi" }}
  {{ partial "review-box.html" . }}
{{else}}<p>No data (paapi:{{ $id }})</p>{{ end }}
```

と定義すれば

```text
{{%/* review-paapi "4621300253" */%}} <!-- プログラミング言語Go -->
```

で全く同じ結果を得られる。

## リモートの JSON データを利用する

今回は `data` ディレクトリ上に JSON ファイルを置いて利用する方法を紹介したが，リモートにある JSON データを直接利用することもできる。
こんな感じで JSON データを取得できるようだ。

```text
{{ $dataJ := getJSON "url" }}
```

ビルドする際にリモートサーバとのやり取りがボトルネックになる気がするが，色々と使い道はあるだろう。

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[books-data]: https://github.com/spiegel-im-spiegel/books-data "spiegel-im-spiegel/books-data: Search for Books Data"
