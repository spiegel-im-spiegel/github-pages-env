+++
title = "書籍データ取得ツール books-data v0.3.0 をリリースした"
date =  "2019-08-12T15:10:56+09:00"
description = "ようやく青空文庫 API に対応したですよ。これで個人的に欲しい機能は一通り揃ったかな。"
image = "/images/attention/tools.png"
tags = [ "tools", "book", "books-data", "pa-api", "openbd", "aozora", "aozora-api" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

書籍データ取得ツール [books-data] v0.3.0 をリリースした。

- [Release v0.3.0 · spiegel-im-spiegel/books-data · GitHub](https://github.com/spiegel-im-spiegel/books-data/releases/tag/v0.3.0)

ようやく[青空文庫] API に対応したですよ。

{{< highlight text "hl_lines=17" >}}
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
      --config string          Config file (default $HOME/.books-data.yaml)
      --debug                  for debug
  -h, --help                   help for books-data
  -i, --isbn string            ISBN code
      --marketplace string     Config: PA-API Marketplace (default "webservices.amazon.co.jp")
  -l, --review-log string      Config: Review log file (JSON format)
      --secret-key string      Config: PA-API Secret Access Key
  -t, --template-file string   Template file for formatted output

Use "books-data [command] --help" for more information about a command.
{{< /highlight >}}

これに伴い [aozora-api] パッケージも公開し [books-data] に組み込んでいる。
ちなみに [aozora-api] パッケージは[青空文庫] API をフルサポートしている。

たとえば谷崎潤一郎さんの「陰翳礼讃」の図書カードNo.は 56642 なので

```text
$ books-data search -c 56642 | jq .
{
  "Type": "aozora",
  "ID": "56642",
  "Title": "陰翳礼讃",
  "URL": "https://www.aozora.gr.jp/cards/001383/card56642.html",
  "Image": {
    "URL": ""
  },
  "ProductType": "青空文庫",
  "Authors": [
    "谷崎 潤一郎"
  ],
  "Codes": [
    {
      "Name": "図書カードNo.",
      "Value": "56642"
    }
  ],
  "PublicationDate": "2016-08-23",
  "LastRelease": "2016-06-10",
  "PublicDomain": true,
  "FirstAppearance": "「経済往来」1933（昭和8）年12月号、1934（昭和9）年1月号",
  "Service": {
    "Name": "aozorahack",
    "URL": "https://aozorahack.org/"
  }
}
```

てな感じで情報が取れる。
翻訳作品も対応している。
たとえば「ニャルラトホテプ」なら

```text
$ books-data search -c 56839 | jq .
{
  "Type": "aozora",
  "ID": "56839",
  "Title": "ニャルラトホテプ",
  "OriginalTitle": "NYARLATHOTEP",
  "URL": "https://www.aozora.gr.jp/cards/001699/card56839.html",
  "Image": {
    "URL": ""
  },
  "ProductType": "青空文庫",
  "Authors": [
    "ラヴクラフト ハワード・フィリップス"
  ],
  "Creators": [
    {
      "Name": "大久保 ゆう",
      "Role": "翻訳"
    }
  ],
  "Codes": [
    {
      "Name": "図書カードNo.",
      "Value": "56839"
    }
  ],
  "PublicationDate": "2014-04-04",
  "LastRelease": "2015-08-19",
  "Service": {
    "Name": "aozorahack",
    "URL": "https://aozorahack.org/"
  }
}
```

という感じ。

[青空文庫]は[書影データを持ってない]({{< ref "/remark/2019/03/book-cover-for-aozora-bunko.md" >}} "「青空文庫」用の書影がほしい")のが寂しい限りだが，自前で用意できるならレビューデータ作成時に `--image-url` オプションを使って

```text
$ books-data review -c 56642 --image-url https://text.baldanders.info/images/aozora/card56839.svg -r 4 "SAN 値が下がる。" | jq .
{
  "Book": {
    "Type": "aozora",
    "ID": "56642",
    "Title": "陰翳礼讃",
    "URL": "https://www.aozora.gr.jp/cards/001383/card56642.html",
    "Image": {
      "URL": "https://text.baldanders.info/images/aozora/card56839.svg"
    },
    "ProductType": "青空文庫",
    "Authors": [
      "谷崎 潤一郎"
    ],
    "Codes": [
      {
        "Name": "図書カードNo.",
        "Value": "56642"
      }
    ],
    "PublicationDate": "2016-08-23",
    "LastRelease": "2016-06-10",
    "PublicDomain": true,
    "FirstAppearance": "「経済往来」1933（昭和8）年12月号、1934（昭和9）年1月号",
    "Service": {
      "Name": "aozorahack",
      "URL": "https://aozorahack.org/"
    }
  },
  "Date": "2019-08-12",
  "Rating": 4,
  "Star": [
    true,
    true,
    true,
    true,
    false
  ],
  "Description": "SAN 値が下がる。"
}
```

などとすれば書影データを補完できる。
これで適当なテンプレートを噛ませれば

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.aozora.gr.jp/cards/001699/card56839.html"><img src="https://text.baldanders.info/images/aozora/card56839.svg" width="110" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.aozora.gr.jp/cards/001699/card56839.html">ニャルラトホテプ</a></dt>
    <dd>原題: NYARLATHOTEP</dd>
    <dd>ラヴクラフト ハワード・フィリップス</dd>
    <dd>大久保 ゆう (翻訳)</dd>
    <dd> 2014-04-04 (Release 2015-08-19)</dd>
    <dd>青空文庫</dd>
    <dd>56839 (図書カードNo.)</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">SAN 値が下がる。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-03-28">2019-03-28</abbr> (powered by <a href="https://aozorahack.org/">aozorahack</a>)</p>
</div>

てな感じに仕上げることができる。

これで個人的に欲しい機能は一通り揃ったかな。

## ブックマーク

- [#aozorahack に関する覚え書き]({{< ref "/remark/2019/08/about-aozorahack.md" >}})
- [「青空文庫」用の書影がほしい]({{< ref "/remark/2019/03/book-cover-for-aozora-bunko.md" >}})
- [Go 言語用青空文庫 API クライアント・パッケージ]({{< ref "/release/aozora-api-package-for-golang.md" >}})

[books-data]: https://github.com/spiegel-im-spiegel/books-data "spiegel-im-spiegel/books-data: Search for Books Data"
[PA-API]: https://affiliate.amazon.co.jp/assoc_credentials/home "Product Advertising API"
[openBD]: https://openbd.jp/ "openBD | 書誌情報・書影を自由に"
[青空文庫]: https://www.aozora.gr.jp/ "青空文庫　Aozora Bunko"
[aozora-api]: https://github.com/spiegel-im-spiegel/aozora-api "spiegel-im-spiegel/aozora-api: APIs for Aozora-bunko RESTful Service by Golang"
