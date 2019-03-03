+++
title = "JSON-LD に対応してみた"
date =  "2017-11-17T23:40:05+09:00"
update =  "2017-11-18T12:50:11+09:00"
description = "以前，このサイトを Twitter Cards に対応させたのだが，今回も思いつきで JSON-LD に対応させることにした。"
image = "/images/attention/remark.jpg"
tags        = [ "site", "semantic ", "web", "metadata", "json" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

以前，このサイトを [Twitter Cards に対応させた]({{< ref "/remark/2017/10/twitter-card-metadata.md" >}} "Twitter Card メタデータに対応した")のだが，今回も思いつきで [JSON-LD] に対応させることにした。

- [JSON-LD - JSON for Linking Data](https://json-ld.org/)

[JSON-LD] の特徴は複数の情報のかたまりを IRI[^iri] で繋ぐことで複雑なデータ構造を記述できる点にある。
これを使ってネット上のリソースのメタデータとその関係を表現するわけだ。

[^iri]: IRI (Internationalized Resource Indicator) というのは [JSON-LD] 独自の用語のようだが，ここでは大雑把に URL (Uniform Resource Locator) や URI (Uniform Resource Indicator) と同じようにネット上のリソースを一意に識別する識別子であると理解しておけばよい。なお，[テストツール]を見るかぎり Google 検索では IRI からメタデータの場所（URL または URI）を推測して取得するといったことはしないようだ。したがって Web ページに [JSON-LD] を設置する場合には，関連する全てのメタデータを記述する必要がある。冗長すぎるやろ。

ただし Web ページの場合は， RDFa や Microdata などと違って，語彙を HTML の要素に埋め込むことが出来ないため冗長にならざるを得ない。
まぁ CMS を使ってサイトや Web ページを管理しているなら，一度テンプレート等を作ってしまえば済む話なので，大した手間ではないかもしれないが。

逆に RESTful API でメタデータを提供する場合は [JSON-LD] はかなり有力な手段となるだろう。

真面目に [JSON-LD] を扱うとなると相当凝ったことができるみたいだが[^jsnld1]，今回は軽めに Google の検索サービスが解釈できる範囲で調整してみようと思う。

[^jsnld1]: [JSON-LD] の仕様を見るかぎり，Microdata や RDFa の語彙を流用したり自前で語彙を作ったりできるっぽい。

- [Introduction to Structured Data | Search | Google Developers](https://developers.google.com/search/docs/guides/intro-structured-data)
- [構造化データ テストツール](https://search.google.com/structured-data/testing-tool)

なお [JSON-LD] を導入するのなら Microdata の記述はページから削除することをお勧めする。
Microdata は未完成のまま開発が終了しており，もはや推奨されない。

## Web ページに [JSON-LD] を埋め込む

[JSON-LD] を Web ページに埋め込む際は，以下のように `<script>` 要素で囲む。

```html
<script type="application/ld+json">
{
  ...
}
</script>
```

メディア・タイプに注意すること。
どうやらこの記述はページ内にいくつ置いてもいいようだ（少なくとも Google の[テストツール]には怒られなかった）。

```html
<script type="application/ld+json">
{
  ...
}
</script>

<script type="application/ld+json">
{
  ...
}
</script>
```

## [JSON-LD] の詳細

[JSON-LD] の中身だが，まずは `@context` と `@type` を指定する。

```javascript
{
  "@context": "http://schema.org",
  "@type": "WebSite"
}
```

`@context` で語彙を定義するのだが，どうやら Google は Microdata の語彙（のひとつ）である [schema.org] を前提にしているらしい。
`@type` には [schema.org] で定義される content type を指定する。
このサイトは一応ブログサイトなので，以下の content type を使うことにする。

- [WebSite]
- [Blog]
- [BlogPosting]
- [BreadcrumbList]

更に，これらが参照する content type として以下も使用する。

- [Organization]
- [Person]
- [ImageObject]

まずは [WebSite] から。

```html
<script type="application/ld+json">
{
  "@context": "http://schema.org",
  "@type": "WebSite",
  "@id": "https://text.baldanders.info/",
  "inLanguage": "ja",
  "name": "text.Baldanders.info",
  "url": "https://text.baldanders.info/",
  "publisher": {
    "@id": "https://text.baldanders.info/#org"
  },
  "author": {
    "@type": "Person",
    "@id": "https://text.baldanders.info/#maker",
    "name": "Spiegel",
    "url": "https://baldanders.info/spiegel/profile/",
    "image": "https://text.baldanders.info/images/avatar.jpg"
  },
  "image": "https://text.baldanders.info/images/avatar.jpg",
  "description": "帰ってきた「しっぽのさきっちょ」"
}
</script>
<script type="application/ld+json">
{
  "@context": "http://schema.org",
  "@type": "Organization",
  "@id": "https://text.baldanders.info/#org",
  "name": "text.Baldanders.info",
  "logo": {
    "@type": "ImageObject",
    "@id": "https://text.baldanders.info/#logo",
    "url": "https://text.baldanders.info/images/avatar.jpg"
  }
}
</script>
```

`@id` は IRI を定義または参照する。
`@id` によって [Organization] のデータを `publisher` から参照しているのがお分かりだろうか。
これを[テストツール]にかけるとこんな感じになる。

{{< fig-img src="https://photo.baldanders.info/flickr/image/38448664942_m.png" title="Structured Data Testing Tool (1)" link="https://photo.baldanders.info/flickr/38448664942/" >}}

`publisher` に [Organization] のデータが入っているのがわかると思う。

本来の [schema.org] の定義では `publisher` には [Organization] と [Person] のどちらも有効な筈なのだが，[テストツール]は [Organization] しか受け付けないようだ。
これではうちのような個人サイトは大変に困るのだが，しょうがないので [Organization] にテキトーな情報を入れてお茶を濁している。
何とかしてよ Google 先生！

次は [Blog] の情報をセットする。

```html
<script type="application/ld+json">
{
  "@context": "http://schema.org",
  "@type": "Blog",
  "@id": "https://text.baldanders.info/remark/",
  "url": "https://text.baldanders.info/remark/",
  "inLanguage": "ja",
  "name": "しっぽのさきっちょ",
  "description": "とりとめのない四方山話。",
  "image": "https://text.baldanders.info/images/attention/remark.jpg",
  "publisher": {
    "@id": "https://text.baldanders.info/#org"
  },
  "author": {
    "@type": "Person",
    "@id": "https://text.baldanders.info/remark/#maker",
    "name": "Spiegel",
    "url": "https://baldanders.info/spiegel/profile/",
    "image": "https://text.baldanders.info/images/avatar.jpg"
  }
}
</script>
```

`publisher` の参照先は [WebSite] と同じなので省略する。
これを[テストツール]にかけるとこんな感じになる。

{{< fig-img src="https://photo.baldanders.info/flickr/image/37593823405_m.png" title="Structured Data Testing Tool (2)" link="https://photo.baldanders.info/flickr/37593823405/" >}}

どんどん行こう。
[BlogPosting] はこんな感じ。

```html
<script type="application/ld+json">
{
  "@context": "http://schema.org",
  "@type": "BlogPosting",
  "@id": "https://text.baldanders.info/remark/2017/11/qiitadon/",
  "url": "https://text.baldanders.info/remark/2017/11/qiitadon/",
  "mainEntityOfPage": "https://text.baldanders.info/remark/2017/11/qiitadon/",
  "inLanguage": "ja",
  "name": "Qiita って Mastodon やってたのか",
  "description": "私にとって今年最初の大外しは「GW 過ぎたら Mastodon のことなんかみんな忘れてる」だったが，本当に忘れていたのは私だけだったようだ。",
  "headline": "私にとって今年最初の大外しは「GW 過ぎたら Mastodon のことなんかみんな忘れてる」だったが，本当に忘れていたのは私だけだったようだ。",
  "keywords": "mastodon, communication",
  "image": "https://text.baldanders.info/images/attention/remark.jpg",
  "datePublished": "2017-11-10T13:49:58+09:00",
  "dateModified": "2017-11-16T10:09:40+09:00",
  "publisher": {
    "@id": "https://text.baldanders.info/#org"
  },
  "author": {
    "@type": "Person",
    "@id": "https://text.baldanders.info/remark/2017/11/qiitadon/#maker",
    "name": "Spiegel",
    "url": "https://baldanders.info/spiegel/profile/",
    "image": "https://text.baldanders.info/images/avatar.jpg"
  },
  "license": "https://creativecommons.org/licenses/by-sa/4.0/"
}
</script>
```

Google 検索は [BlogPosting] の内容を参照している。
Google 検索が [BlogPosting] で要求するプロパティは以下の通り。

| Properties              |          Data Type          |     AMP     |   non-AMP   |
|:----------------------- |:---------------------------:|:-----------:|:-----------:|
| `mainEntityOfPage`      |            `URL`            | recommended |   ignored   |
| `headline`              |           `Text`            |  required   | recommended |
| `image`                 |   ` ImageObject` or `URL`   |  required   | recommended |
| `publisher`             |       ` Organization`       |  required   |   ignored   |
| `publisher.name`        |           ` Text`           |  required   |   ignored   |
| `publisher.logo`        |       ` ImageObject`        |  required   |   ignored   |
| `publisher.logo.url`    |           ` URL`            |  required   |   ignored   |
| `publisher.logo.height` |          ` Number`          |  required   |   ignored   |
| `publisher.logo.width`  |          ` Number`          |  required   |   ignored   |
| `datePublished`         |         ` DateTime`         |  required   |   ignored   |
| `dateModified`          |         ` DateTime`         |  required   |   ignored   |
| `author`                | ` Person` or `Organization` |  required   |   ignored   |
| `author.name`           |           ` Text`           |  required   |   ignored   |
| `description`           |           ` Text`           | recommended |   ignored   |

AMP (Accelerated Mobile Pages) と non-AMP で要求が異なるが，[テストツール]は AMP を前提にしているようで，これらのプロパティがないとエラーまたは警告を吐く[^tst1]。

[^tst1]: ただし `publisher.logo.height` と `publisher.logo.width` は例外のようで，これらのプロパティがなくても[テストツール]ではエラーにならなかった。まぁ画像のサイズは画像データから読み取れるからね。

これも[テストツール]にかけてみよう。
こんな感じ。

{{< fig-img src="https://photo.baldanders.info/flickr/image/37594061745_m.png" title="Structured Data Testing Tool (3)" link="https://photo.baldanders.info/flickr/37594061745/" >}}

最後に [BreadcrumbList] （パンくずリスト[^bcl1]）はこんな感じになる。

[^bcl1]: 一応説明しておくと「パンくずリスト」とはサイト内のページ位置を見失わないようにするためのナビゲーション表示を指す。名前で想像される通り「ヘンゼルとグレーテル」の童話が語源になっているらしい。まぁ，いまどきパンくずリストがないと迷子になってしまうようなダサい構成のサイトは少なくなってると思うけど。

```html
<script type="application/ld+json">
{
  "@context": "http://schema.org",
  "@type": "BreadcrumbList",
  "@id": "https://text.baldanders.info/remark/2017/11/qiitadon/#breadcrumb-list",
  "itemListElement": [
    {
      "@type": "ListItem",
      "position": 1,
      "item": {
        "@id": "https://text.baldanders.info/"
      }
    },
    {
      "@type": "ListItem",
      "position": 2,
      "item": {
        "@id": "https://text.baldanders.info/remark/"
      }
    }
  ]
}
</script>
```

Google 検索は [BreadcrumbList] の内容も参照している。
Google 検索が [BreadcrumbList] で要求するプロパティは以下の通り。

| Properties | Data Type | Requirement |
|:---------- |:---------:|:-----------:|
| `image`    |   `URL`   |  optional   |
| `item`     |  `Thing`  |  required   |
| `name`     |  `Text`   |  required   |
| `position` | `Integer` |  required   |

上述のコードでは `item` を `@id` の参照先と繋げている。
具体的には，最初の階層に [WebSite] データの `@id` を，2番目の階層に [Blog] データの `@id` を指定している。
これによって `item` の中に `name` や `image` が含まれるため，不足なく情報を網羅できている。

## ブックマーク

- [Google推奨「JSON-LD」で構造化マークアップ - Qiita](https://qiita.com/narumana/items/b66969b80cce848b2ddf)
- [Google リッチカードの導入 - Qiita](https://qiita.com/ko8@github/items/d2cf8786032972d7d090)
- [schema.orgのパンくずリストをようやくGoogleがサポート開始 | 海外SEO情報ブログ](https://www.suzukikenichi.com/blog/google-finally-supports-breadcrumbs-with-schema-org/)
- [JSON-LDによるブログの構造化データをHugoで実現する | TRIAL DANCE](https://blog.mizukmb.net/post/apply-json-ld/)
- [Hugo で JSON-LD 内の URL がなぜかエスケープされる問題 - Qiita](https://qiita.com/1000k/items/2f2258cbbe9c9493551e)
- [HUGOで作れるCMSっぽいパーツ:関連記事・目次・JSON-LDなど - Qiita](https://qiita.com/y_hokkey/items/f9d8b66b3770a82d4c1c)

[JSON-LD]: https://json-ld.org/ "JSON-LD - JSON for Linking Data"
[テストツール]: https://search.google.com/structured-data/testing-tool "構造化データ テストツール"
[schema.org]: http://schema.org/
[WebSite]: http://schema.org/WebSite "WebSite - schema.org"
[Blog]: http://schema.org/Blog "Blog - schema.org"
[Article]: http://schema.org/Article "Article - schema.org"
[BlogPosting]: http://schema.org/BlogPosting "BlogPosting - schema.org"
[BreadcrumbList]: http://schema.org/BreadcrumbList "BreadcrumbList - schema.org"
[Organization]: http://schema.org/Organization "Organization - schema.org"
[Person]: http://schema.org/Person "Person - schema.org"
[ImageObject]: http://schema.org/ImageObject "ImageObject - schema.org"
