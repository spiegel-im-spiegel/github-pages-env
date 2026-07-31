+++
title = "書籍データ取得ツール books-data の最初のリリース"
date =  "2019-07-27T21:46:19+09:00"
description = "ちうわけで PA-API にも openBD にも対応する書籍情報取得ツールを作ったですよ。"
image = "/images/attention/tools.png"
tags = [ "tools", "book", "books-data", "pa-api", "openbd", "golang" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

いやね。
以前[公開]({{< ref "/release/2019/01/amazon-product-advertising-api.md" >}} "しょうがないので Amazon アフィリエイト・リンク作成ツールを作ったですよ")した [PA-API] を使った書籍情報検索ツール [amazon-item] が一時的に使えなくなったのよ。
理由は30日以上売り上げがなかったため。

まぁ，ここはアフィリエイトが目的のブログじゃないし，いつかはこんなことになるだろうとは思っていたけど，やっぱ [PA-API] に依存するのは拙いと思い代わりになるものを探してたんだけど，現状では [openBD] くらいしか選択肢がないらしい。

ちうわけで [PA-API] にも [openBD] にも対応する書籍情報取得ツールを作ったですよ。

- [spiegel-im-spiegel/books-data: Search for Books Data](https://github.com/spiegel-im-spiegel/books-data)
- [Release v0.1.0 · spiegel-im-spiegel/books-data](https://github.com/spiegel-im-spiegel/books-data/releases/tag/v0.1.0)

例によって CLI のツール。

```text
$ books-data -h
Search for books data

Usage:
  books-data [flags]
  books-data [command]

Available Commands:
  help        Help about any command
  openbd      Search for books data by openBD
  paapi       Search for books data by PA-API
  version     Print the version number

Flags:
      --debug   for debug
  -h, --help    help for books-data
      --pipe    Import description from Stdin
      --raw     Output raw data from openBD

Use "books-data [command] --help" for more information about a command.
```

`openbd` および `paapi` コマンドで取得先を切り換える。

## [PA-API] による書籍情報の取得

[PA-API] による書籍情報の取得は以下のコマンドラインで行う。

```text
$ books-data paapi -h
Search for books data by PA-API

Usage:
  books-data paapi [flags] [description]

Flags:
      --access-key string      PA-API: Access Key ID
  -a, --asin string            Amazon ASIN code
      --associate-tag string   PA-API: Associate Tag
      --config string          config file (default $HOME/.paapi.yaml)
  -h, --help                   help for paapi
      --marketplace string     PA-API: Marketplace (default "webservices.amazon.co.jp")
  -r, --rating int             Rating of product
      --review-date string     Date of review
      --secret-key string      PA-API: Secret Access Key
  -t, --template string        Template file

Global Flags:
      --debug   for debug
      --pipe    Import description from Stdin
      --raw     Output raw data from openBD
```

[PA-API] に必要な設定はホーム・ディレクトリの `.paapi.yaml` ファイルに記述する。
こんな感じ。

```text
$ cat ~/.paapi.yaml
associate-tag: mytag-20
access-key: AKIAIOSFODNN7EXAMPLE
secret-key: 1234567890
```

これをセットした上で ASIN コードを指定して書籍情報を取得する。
結果は JSON フォーマットで出力される。
こんな感じ（一部データを端折っているのであしからず）。

```text
$ books-data paapi -a 427406932X | jq .
{
  "Book": {
    "ID": "427406932X",
    "Title": "リーン開発の現場 カンバンによる大規模プロジェクトの運営",
    "URL": "https://www.amazon.co.jp/exec/obidos/ASIN/427406932X/mytag-20",
    "Image": {
      "URL": "https://images-fe.ssl-images-amazon.com/images/I/51llL1uygcL._SL160_.jpg",
      "Height": 160,
      "Width": 116
    },
    "ProductType": "単行本（ソフトカバー）",
    "Authors": [
      "Henrik Kniberg"
    ],
    "Creators": [
      {
        "Name": "角谷 信太郎",
        "Role": "翻訳"
      },
      {
        "Name": "市谷 聡啓",
        "Role": "翻訳"
      },
      {
        "Name": "藤原 大",
        "Role": "翻訳"
      }
    ],
    "Publisher": "オーム社",
    "Codes": [
      {
        "Name": "ASIN",
        "Value": "427406932X"
      },
      {
        "Name": "EAN",
        "Value": "9784274069321"
      }
    ],
    "PublicationDate": "2013-10-26",
    "LastRelease": "0001-01-01",
    "Service": {
      "Name": "PA-API",
      "URL": "https://affiliate.amazon.co.jp/assoc_credentials/home"
    }
  },
  "Date": "2019-07-27",
  "Rating": 0,
  "Star": [
    false,
    false,
    false,
    false,
    false
  ]
}
```

また適当なテンプレートファイルを指定することで任意のフォーマットで指定することも可能。
たとえば

```text
$ cat template/template.html 
<div class="hreview">{{ if .Book.Image.URL }}
  <div class="photo">{{ if .Book.URL }}<a class="item url" href="{{ .Book.URL }}">{{ end }}<img src="{{ .Book.Image.URL }}"{{ with .Book.Image.Width }} width="{{ . }}"{{ end }} alt="photo">{{ if .Book.URL }}</a>{{ end }}</div>{{ end }}
  <dl class="fn">
    <dt>{{ if .Book.URL }}<a href="{{ .Book.URL }}">{{ end }}{{ .Book.Title }}{{ if .Book.URL }}</a>{{ end }}</dt>{{ if .Book.Authors }}
    <dd>{{ range $i, $v := .Book.Authors }}{{ if ne $i 0 }}, {{ end }}{{ $v }}{{ end }}</dd>{{ end }}{{ if .Book.Creators }}
    <dd>{{ range $i, $v := .Book.Creators }}{{ if ne $i 0 }}, {{ end }}{{ $v }}{{ end }}</dd>{{ end }}
    <dd>{{ .Book.Publisher }}{{ if gt .Book.PublicationDate.Year 1 }} {{ .Book.PublicationDate }}{{ end }}{{ if gt .Book.LastRelease.Year 1 }} (Release {{ .Book.LastRelease }}){{ end }}</dd>
    <dd>{{ .Book.ProductType }}</dd>{{ if .Book.Codes }}
    <dd>{{ range $i, $v := .Book.Codes }}{{ if ne $i 0 }}, {{ end }}{{ $v }}{{ end }}</dd>{{ end }}{{ if gt .Rating 0 }}
    <dd>Rating<abbr class="rating fa-sm" title="{{ .Rating }}">{{ range .Star }}&nbsp;{{ if . }}<i class="fas fa-star"></i>{{ else }}<i class="far fa-star"></i>{{ end }}{{ end }}</abbr></dd>{{ end }}
  </dl>
  <p class="description">{{ .Description }}</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="{{ .Date }}">{{ .Date }}</abbr> (powered by {{ if .Book.Service.URL }}<a href="{{ .Book.Service.URL }}" >{{ end }}{{ .Book.Service.Name }}{{ if .Book.Service.URL }}</a>{{ end }})</p>
</div>
```

というテンプレートファイルを使えば

```text
$ books-data paapi -a 427406932X -r 4 -t template/template.html 
<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/427406932X/mytag-20"><img src="https://images-fe.ssl-images-amazon.com/images/I/51llL1uygcL._SL160_.jpg" width="116" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/427406932X/mytag-20">リーン開発の現場 カンバンによる大規模プロジェクトの運営</a></dt>
    <dd>Henrik Kniberg</dd>
    <dd>角谷 信太郎 (翻訳), 市谷 聡啓 (翻訳), 藤原 大 (翻訳)</dd>
    <dd>オーム社 2013-10-26</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>427406932X (ASIN), 9784274069321 (EAN)</dd>
    <dd>Rating<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description"></p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-07-27">2019-07-27</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home" >PA-API</a>)</p>
</div>
```

ってな感じに整形してくれる。

## [openBD] による書籍情報の取得

[openBD] による書籍情報の取得は以下のコマンドラインで行う。

```text
$ books-data openbd -h
Search for books data by openBD

Usage:
  books-data openbd [flags] [description]

Flags:
  -h, --help                 help for openbd
  -i, --isbn string          ISBN code
  -r, --rating int           Rating of product
      --review-date string   Date of review
  -t, --template string      Template file

Global Flags:
      --debug   for debug
      --pipe    Import description from Stdin
      --raw     Output raw data from openBD
```

使い方は `paapi` コマンドとほぼ同じ。
微妙にオプション名が違うのは仕様である。

```text
$ books-data openbd -i 427406932X | jq .
{
  "Book": {
    "ID": "9784274069321",
    "Title": "リーン開発の現場 : カンバンによる大規模プロジェクトの運営",
    "Image": {
      "URL": "https://cover.openbd.jp/9784274069321.jpg",
      "Height": 0,
      "Width": 0
    },
    "ProductType": "Book",
    "Authors": [
      "Kniberg,Henrik／著 オーム社／著 オーム社開発局／著 市谷聡啓／翻訳 ほか"
    ],
    "Publisher": "オーム社",
    "Codes": [
      {
        "Name": "ISBN",
        "Value": "9784274069321"
      }
    ],
    "PublicationDate": "2013-10-01",
    "LastRelease": "0001-01-01",
    "Service": {
      "Name": "openBD",
      "URL": "https://openbd.jp/"
    }
  },
  "Date": "2019-07-27",
  "Rating": 0,
  "Star": [
    false,
    false,
    false,
    false,
    false
  ]
}
```

ご覧の通り，圧倒的に情報量が少ない。
書籍を紹介する Web ページへのリンクもないとか，やる気のないことこの上ない。

たとえばテンプレートファイルを

```text
$ cat template/template.bib.txt 
@BOOK{Book:{{ .Book.ID }},
    TITLE = "{{ .Book.Title }}"{{ if .Book.Authors }},
    AUTHOR = "{{ range $i, $v := .Book.Authors }}{{ if ne $i 0 }} and {{ end }}{{ $v }}{{ end }}{{ if .Book.Creators }}{{ range .Book.Creators }} and {{ . }}{{ end }}{{ end }}"{{ end }}{{ if .Book.Publisher }},
    PUBLISHER = {{ "{" }}{{ .Book.Publisher }}{{ "}" }}{{ end }}{{ if gt .Book.PublicationDate.Year 1 }},
    YEAR = {{ .Book.PublicationDate.Year }}{{ end }}
}
```

みたいにすれば

```text
$ books-data openbd -i 427406932X -t template/template.bib.txt 
@BOOK{Book:9784274069321,
    TITLE = "リーン開発の現場 : カンバンによる大規模プロジェクトの運営",
    AUTHOR = "Kniberg,Henrik／著 オーム社／著 オーム社開発局／著 市谷聡啓／翻訳 ほか",
    PUBLISHER = {オーム社},
    YEAR = 2013
}
```

てな感じに [BiBTeX] 風の出力にもできるので，全く使えないこともない？

## 今後の予定

[青空文庫APIサーバ](https://qiita.com/ksato9700/items/626cc82c007ba8337034 "青空文庫APIサーバーのご紹介 - Qiita")とかあるらしいので，こちらも対応していきたい。
でも[青空文庫は書影がない]({{< ref "/remark/2019/03/book-cover-for-aozora-bunko.md" >}} "「青空文庫」用の書影がほしい")んだよなぁ。

## ブックマーク

- [AmazonのProduct Advertising APIの使い方 | Apitore blog](http://blog.apitore.com/2016/08/01/amazon-product-advertising-api/)
- [「Amazon API」の使い方を紹介します！最安値やランキング取得できるよ①－アソシエイトID（タグ）登録編－ | HPcode](https://haniwaman.com/amazon-api-1/)

[amazon-item]: https://github.com/spiegel-im-spiegel/amazon-item "spiegel-im-spiegel/amazon-item: Searching Amazon Items, Powered by PA-API"
[PA-API]: https://affiliate.amazon.co.jp/assoc_credentials/home "Product Advertising API"
[openBD]: https://openbd.jp/ "openBD | 書誌情報・書影を自由に"
[BiBTeX]: https://texwiki.texjp.org/?BibTeX%E9%96%A2%E9%80%A3%E3%83%84%E3%83%BC%E3%83%AB "BibTeX関連ツール - TeX Wiki"
