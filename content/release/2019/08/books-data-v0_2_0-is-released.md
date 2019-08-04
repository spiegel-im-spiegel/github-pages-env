+++
title = "書籍データ取得ツール books-data v0.2.0 をリリースした"
date =  "2019-08-04T19:48:15+09:00"
description = "v0.1.0 の使い勝手が微妙に悪かったのでコマンド構成を変えた。"
image = "/images/attention/tools.png"
tags = [ "tools", "book", "books-data", "pa-api", "openbd" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

書籍データ取得ツール [books-data] v0.2.0 をリリースした。

- [Release v0.2.0 · spiegel-im-spiegel/books-data · GitHub](https://github.com/spiegel-im-spiegel/books-data/releases/tag/v0.2.0)

[v0.1.0]({{< ref "/release/2019/07/first-release-books-data.md" >}} "書籍データ取得ツール books-data の最初のリリース") の使い勝手が微妙に悪かったのでコマンド構成を変えた。

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
  -a, --asin string            Amazon ASIN code
      --associate-tag string   Config: PA-API Associate Tag
      --config string          Config file (default $HOME/.books-data.yaml)
      --debug                  for debug
  -h, --help                   help for books-data
  -i, --isbn string            ISBN code
      --marketplace string     Config: PA-API Marketplace (default "webservices.amazon.co.jp")
  -l, --review-log string      Config: Review log file (JSON format)
      --secret-key string      Config: PA-API Secret Access Key
  -t, --template-file string   Template file for formatted output

Use "books-data [command] --help" for more information about a command.
```

`search` コマンドで [PA-API] または [openBD] から書籍情報を検索する。
ASIN コードを指定する場合はこんな感じ（Web ページへのリンク URL は長ったらしいので省略している，あしからず）。

```text
$ books-data search -a B07TYKJQFK | jq .
{
  "Type": "paapi",
  "ID": "B07TYKJQFK",
  "Title": "数学ガールの秘密ノート／ビットとバイナリー",
  "URL": "https://www.amazon.co.jp/exec/obidos/ASIN/B07TYKJQFK/mytag-20/",
  "Image": {
    "URL": "https://images-fe.ssl-images-amazon.com/images/I/41Q0Hyr8g3L._SL160_.jpg",
    "Height": 160,
    "Width": 111
  },
  "ProductType": "Kindle版",
  "Authors": [
    "結城 浩"
  ],
  "Publisher": "SBクリエイティブ",
  "Codes": [
    {
      "Name": "ASIN",
      "Value": "B07TYKJQFK"
    }
  ],
  "PublicationDate": "2019-07-19",
  "LastRelease": "2019-07-20",
  "Service": {
    "Name": "PA-API",
    "URL": "https://affiliate.amazon.co.jp/assoc_credentials/home"
  }
}
```

[PA-API] へアクセスするためのアクセスキー等の情報は設定ファイル（既定は `$HOME/.books-data.yaml`）で指定できる。

```yaml
marketplace: webservices.amazon.co.jp
associate-tag: mytag-20
access-key: AKIAIOSFODNN7EXAMPLE
secret-key: 1234567890
```

[PA-API] のアクセスキーがない場合は `-i` オプションで ISBN を指定して [openBD] から検索できる。

```text
$ books-data search -i 9784797391398 | jq .
{
  "Type": "openbd",
  "ID": "9784797391398",
  "Title": "数学ガールの秘密ノート／ビットとバイナリー",
  "Image": {
    "URL": "https://cover.openbd.jp/9784797391398.jpg",
    "Height": 0,
    "Width": 0
  },
  "ProductType": "Book",
  "Authors": [
    "結城浩／著"
  ],
  "Publisher": "ＳＢクリエイティブ",
  "Codes": [
    {
      "Name": "ISBN",
      "Value": "9784797391398"
    }
  ],
  "PublicationDate": "2019-07-22",
  "LastRelease": "0001-01-01",
  "Service": {
    "Name": "openBD",
    "URL": "https://openbd.jp/"
  }
}
```

[前回]({{< ref "/release/2019/07/first-release-books-data.md" >}} "書籍データ取得ツール books-data の最初のリリース")書いたとおり，まるでやる気の感じられない情報量だけど（笑）

レビュー情報を作成する場合は `review` コマンドを使う。

```text
$ books-data review -a B07TYKJQFK -r 5 "この本はオススメ" | jq .
{
  "Book": {
    "Type": "paapi",
    "ID": "B07TYKJQFK",
    "Title": "数学ガールの秘密ノート／ビットとバイナリー",
    "URL": "https://www.amazon.co.jp/exec/obidos/ASIN/B07TYKJQFK/mytag-20/",
    "Image": {
      "URL": "https://images-fe.ssl-images-amazon.com/images/I/41Q0Hyr8g3L._SL160_.jpg",
      "Height": 160,
      "Width": 111
    },
    "ProductType": "Kindle版",
    "Authors": [
      "結城 浩"
    ],
    "Publisher": "SBクリエイティブ",
    "Codes": [
      {
        "Name": "ASIN",
        "Value": "B07TYKJQFK"
      }
    ],
    "PublicationDate": "2019-07-19",
    "LastRelease": "2019-07-20",
    "Service": {
      "Name": "PA-API",
      "URL": "https://affiliate.amazon.co.jp/assoc_credentials/home"
    }
  },
  "Date": "2019-08-04",
  "Rating": 5,
  "Star": [
    true,
    true,
    true,
    true,
    true
  ],
  "Description": "この本はオススメ"
}
```

作成したデータ（JSON 形式）はログファイルにも蓄積される。
ログファイルは `review-log` オプションで設定ファイルまたはコマンドラインで指定する。

```yaml
review-log: /home/username/review-log.json
```

ログファイルが設定されていれば `history` コマンドでレビュー情報を取り出すことができる。

```text
$ books-data history -a B07TYKJQFK | jq .
{
  "Book": {
    "Type": "paapi",
    "ID": "B07TYKJQFK",
    "Title": "数学ガールの秘密ノート／ビットとバイナリー",
    "URL": "https://www.amazon.co.jp/exec/obidos/ASIN/B07TYKJQFK/mytag-20/",
    "Image": {
      "URL": "https://images-fe.ssl-images-amazon.com/images/I/41Q0Hyr8g3L._SL160_.jpg",
      "Height": 160,
      "Width": 111
    },
    "ProductType": "Kindle版",
    "Authors": [
      "結城 浩"
    ],
    "Publisher": "SBクリエイティブ",
    "Codes": [
      {
        "Name": "ASIN",
        "Value": "B07TYKJQFK"
      }
    ],
    "PublicationDate": "2019-07-19",
    "LastRelease": "2019-07-20",
    "Service": {
      "Name": "PA-API",
      "URL": "https://affiliate.amazon.co.jp/assoc_credentials/home"
    }
  },
  "Date": "2019-08-04",
  "Rating": 5,
  "Star": [
    true,
    true,
    true,
    true,
    true
  ],
  "Description": "この本はオススメ"
}
```

ログファイルの中身は以下のような JSON 形式の配列として格納されているので，ぶっちゃけテキスト・エディタ等で編集できる（笑）

```json
[
  {
    "Book": {
      "Type": "paapi",
      "ID": "B07TYKJQFK",
      "Title": "数学ガールの秘密ノート／ビットとバイナリー",
      "URL": "https://www.amazon.co.jp/exec/obidos/ASIN/B07TYKJQFK/mytag-20/",
      "Image": {
        "URL": "https://images-fe.ssl-images-amazon.com/images/I/41Q0Hyr8g3L._SL160_.jpg",
        "Height": 160,
        "Width": 111
      },
      "ProductType": "Kindle版",
      "Authors": [
        "結城 浩"
      ],
      "Publisher": "SBクリエイティブ",
      "Codes": [
        {
          "Name": "ASIN",
          "Value": "B07TYKJQFK"
        }
      ],
      "PublicationDate": "2019-07-19",
      "LastRelease": "2019-07-20",
      "Service": {
        "Name": "PA-API",
        "URL": "https://affiliate.amazon.co.jp/assoc_credentials/home"
      }
    },
    "Date": "2019-08-04",
    "Rating": 5,
    "Star": [
      true,
      true,
      true,
      true,
      true
    ],
    "Description": "この本はオススメ"
  }
]
```

いや，最初は SQLite とか使おうと思ったのよ。
でも途中までテーブル設計してたら思ったより面倒な感じがしたのでしばらく先送りすることにした。
他に追加したい機能もあったりするし，データ構造が固まるまでは JSON 形式のままアドホックな運用でもいいかな，と。

以上，いいわけでした。

なお `search` および `review`, `history` コマンドではテンプレートファイルを指定して出力を整形できる。
こんな感じ。

```text
$ books-data search -a B07TYKJQFK -t testdata/book-template/template.bib.txt 
@BOOK{Book:B07TYKJQFK,
    TITLE = "数学ガールの秘密ノート／ビットとバイナリー",
    AUTHOR = "結城 浩",
    PUBLISHER = {SBクリエイティブ},
    YEAR = 2019
}
```

よし。
次はいよいよ[青空文庫APIサーバ](https://qiita.com/ksato9700/items/626cc82c007ba8337034 "青空文庫APIサーバーのご紹介 - Qiita")への対応に取り掛かるぞ。

[books-data]: https://github.com/spiegel-im-spiegel/books-data "spiegel-im-spiegel/books-data: Search for Books Data"
[PA-API]: https://affiliate.amazon.co.jp/assoc_credentials/home "Product Advertising API"
[openBD]: https://openbd.jp/ "openBD | 書誌情報・書影を自由に"

## 今回紹介した書籍

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%95%B0%E5%AD%A6%E3%82%AC%E3%83%BC%E3%83%AB%E3%81%AE%E7%A7%98%E5%AF%86%E3%83%8E%E3%83%BC%E3%83%88%EF%BC%8F%E3%83%93%E3%83%83%E3%83%88%E3%81%A8%E3%83%90%E3%82%A4%E3%83%8A%E3%83%AA%E3%83%BC-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B07TYKJQFK?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B07TYKJQFK"><img src="https://images-fe.ssl-images-amazon.com/images/I/41Q0Hyr8g3L._SL160_.jpg" width="111" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%95%B0%E5%AD%A6%E3%82%AC%E3%83%BC%E3%83%AB%E3%81%AE%E7%A7%98%E5%AF%86%E3%83%8E%E3%83%BC%E3%83%88%EF%BC%8F%E3%83%93%E3%83%83%E3%83%88%E3%81%A8%E3%83%90%E3%82%A4%E3%83%8A%E3%83%AA%E3%83%BC-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B07TYKJQFK?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B07TYKJQFK">数学ガールの秘密ノート／ビットとバイナリー</a></dt>
    <dd>結城 浩</dd>
    <dd>SBクリエイティブ 2019-07-19 (Release 2019-07-20)</dd>
    <dd>Kindle版</dd>
    <dd>B07TYKJQFK (ASIN)</dd>
  </dl>
  <p class="description">まだ読んでる途中。後日，感想を書く予定。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-08-04">2019-08-04</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home" >PA-API</a>)</p>
</div>
