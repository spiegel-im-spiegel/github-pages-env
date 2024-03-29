+++
title = "#aozorahack に関する覚え書き"
date =  "2019-08-07T23:22:53+09:00"
description = "自作の books-data に青空文庫の書籍情報を取得する機能を組み込もうと思っているのだが，その前に #aozorahack について情報をまとめておく。"
image = "/images/attention/kitten.jpg"
tags = [ "aozora", "book", "aozorahack", "api" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

自作の [books-data] に青空文庫の書籍情報を取得する機能を組み込もうと思っているのだが，その前に #[aozorahack] について情報をまとめておく。

なお，この記事は覚え書きとして随時更新する予定である。

## [青空文庫]と #[aozorahack]

#[aozorahack] は2015年の「[Code for 青空文庫](https://baldanders.info/blog/000846/)」の成果とも言える活動で，[青空文庫]そのものではないが協働関係にある。
基本的にボランティア活動だが「[本の未来基金]」による支援を受けているらしい。
2016年からは「[本の未来基金]」が [`aozorahack.org`](https://aozorahack.org/) ドメインを取得し #[aozorahack] に提供している。

{{< fig-quote type="markdown" title="aozorahack" link="https://aozorahack.org/" >}}
「著作権切れのテキストを中心に、今や1万3000作品の自由な電子本を収蔵する日本最大級のインターネット電子図書館「青空文庫」。青空文庫の運営はボランティアによって受け継がれてきました。1997年のスタートから着実に蔵書数を増やしてきましたが、だんだんとそのシステムは古くなってきてしまいました。

そこで青空文庫を技術の力で支えようとスタートしたのが #aozorahack です。青空文庫とは協業関係にあり、本の未来基金の支援を受けています。当面の目標はサーバの安定運用とシステムの分散ですが、それ以外にも青空文庫が今後も続いていけるようなしくみをみなさんと一緒に考えて、動かしていければと考えています。」
{{< /fig-quote >}}

当面の活動としては，[青空文庫]活動の worlflow を変えることなくシステム周りの強化・運用を支援し，更に[青空文庫]周辺のツールやサービスを充実させていくことを目標としているようだ。

## [青空文庫] API

こうした活動の成果のひとつが「[青空文庫] API」で，サーバ・サイドの [aozorahack/pubserver2] とクライアント・サイドの [aozorahack/aozora-cli] で構成される。
[aozorahack/aozora-cli] は Python で組まれていて中のコードも素朴な作りなので理解は容易だろう。

「[青空文庫] API」では作品（`books`）および作家（`persons`）に関する情報の検索および参照機能を提供する。
API は RESTful な構成になっているため他サービスからも利用しやすい。

### API サーバ

[aozorahack/aozora-cli] の中身を見ると API を提供するサーバのドメインは `www.aozorahack.net` になっていた。
この記事では `www.aozorahack.net` を[青空文庫] API サーバとして解説していく。

### 作品情報の照会

以下の URI で[青空文庫]に収録されている全作品の情報が取れる。

```text
$ http://www.aozorahack.net/api/v0.1/books
```

URI には以下のオプションを付加できる。

| オプション名 | 機能                                                 | 既定値 |
| ------------:| ---------------------------------------------------- | ------ |
|      `title` | タイトル名で絞り込む                                 | なし   |
|     `author` | 著者名で絞り込む                                     | なし   |
|     `fields` | 出力するデータ項目を指定する[^fields1]               | 全項目 |
|      `limit` | 出力するデータ数を指定する                           | 100    |
|       `skip` | データ出力をスキップする                             | 0      |
|      `after` | 指定した日付（`YYYY-MM-DD`）以降のデータのみ出力する | なし   |

[^fields1]: `fields` オプションは効いてない感じ？

たとえば「天に積む宝」というワードを含むタイトルを探す場合には

```
$ curl http://www.aozorahack.net/api/v0.1/books?title=/%E5%A4%A9%E3%81%AB%E7%A9%8D%E3%82%80%E5%AE%9D/ | jq .
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

などとする。
検索キーワードを `/.../` で囲むことで正規表現に対応するようだ。

図書カードNo.（`book_id`）が分かっている場合は

```text
$ curl http://www.aozorahack.net/api/v0.1/books/59489 | jq .
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
```

としてもよい。

### 作家情報の照会

以下の URI で[青空文庫]に収録されている全ての作家情報が取れる。

```text
$ http://www.aozorahack.net/api/v0.1/persons
```

URI には以下のオプションを付加できる。

| オプション名 | 機能             | 既定値 |
| ------------:| ---------------- | ------ |
|       `name` | 作家名で絞り込む | なし   |

たとえば「「天に積む宝」のふやし方、へらし方」を著した「富田倫生」さんの情報を探す場合には

```text
$ curl http://www.aozorahack.net/api/v0.1/persons?name=%E5%AF%8C%E7%94%B0%E5%80%AB%E7%94%9F | jq .
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

などとする。

作家No.（`person_id`）が分かっている場合は

```text
$ curl http://www.aozorahack.net/api/v0.1/persons/55 | jq .
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
```

としてもよい。

## ブックマーク

- [青空文庫APIサーバーのご紹介 - Qiita](https://qiita.com/ksato9700/items/626cc82c007ba8337034)
- [青空文庫のデータ構造について - Qiita](https://qiita.com/ksato9700/items/48fd0eba67316d58b9d6)

[books-data]: https://github.com/spiegel-im-spiegel/books-data "spiegel-im-spiegel/books-data: Search for Books Data"
[aozorahack]: https://aozorahack.org/
[青空文庫]: https://www.aozora.gr.jp/ "青空文庫　Aozora Bunko"
[本の未来基金]: https://www.honnomirai.net/
[aozorahack/pubserver2]: https://github.com/aozorahack/pubserver2 "aozorahack/pubserver2: Pubserver"
[aozorahack/aozora-cli]: https://github.com/aozorahack/aozora-cli

## 参考図書

{{% review-paapi "4899840721" %}} <!-- インターネット図書館 青空文庫 -->
{{% review-paapi "4756117074" %}} <!-- 本の未来 -->
