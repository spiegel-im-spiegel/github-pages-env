+++
title = "ブログ記事に OGP article メタデータを追加した"
date =  "2026-08-10T15:20:48+09:00"
description = "OGP は相変わらずセマンティック・ウェブからかけ離れたところにいるのがいい感じである（笑）"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "site", "semantic", "web", "metadata", "html" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

他所のブログ記事を眺めてたら `<head>` 要素内のメタデータに

```html
<meta property="article:tag" content="foo">
```

みたいな記述が見える。
`article` なんてあったか？ と思い，調べてみたら，どうやら [OGP] の中で定義されている語彙（vocabulary）のようだ。

## オブジェクトタイプに article を指定する

`article` を使うためには `prefix` に `article` の語彙を定義し `og:type` に `article` を指定する。

```html
<html>
<head prefix="og: http://ogp.me/ns# article: http://ogp.me/ns/article#">
  ...
  <meta property="og:type" content="article">
  ...
</head>
```

`content` 属性は本来リテラルデータを指すので内容に（RDF 上の）意味はないのだが，どうやらただの文字列ではなく， Web 上のリソースを指定しなければいけないようだ。
したがって `<head>` 要素に

```html
<meta property="og:type" content="article">
```

と書きたければ `<head>` 要素の `prefix` 属性を使って

```html
<head prefix="og: http://ogp.me/ns# article: http://ogp.me/ns/article#"> ... </head>
```

のように記述する必要がある。
まぁ RDFa の設計思想的には間違ってるけどね。
もう，私は気にしないことにしている。

[OGP] では `article` を含め，以下の語彙が定義されている。

{{< div-gen type="markdown" class="smaller" >}}
| IRI | resource | types |
|:---|:---|:---|
| `music` | `https://ogp.me/ns/music#` | `music.song`, `music.album`, `music.playlist`, `music.radio_station` |
| `video` | `https://ogp.me/ns/video#` | `video.movie`, `video.episode`, `video.tv_show`, `video.other` |
| `article` | `http://ogp.me/ns/article#` | `article` |
| `book` | `http://ogp.me/ns/book#` | `book` |
| `payment` | `http://ogp.me/ns/payment#` | `payment.link` |
| `profile` | `http://ogp.me/ns/profile#` | `profile` |
| `website` | `http://ogp.me/ns/website#` | `website` |
{{< /div-gen >}}

これ以外の語彙が欲しいなら独自に定義するしかない。

## article のメタデータ

`article` の語彙で使えるメタデータは以下の通り

| property | content | description |
|:---|:---|:---|
| `article:published_time` | ISO 8601 形式の日時 | 記事の公開日時 |
| `article:modified_time` | ISO 8601 形式の日時 | 記事の最終更新日時 |
| `article:expiration_time` | ISO 8601 形式の日時 | 記事の有効期限 |
| `article:author` | `profile` リソース | 記事の著者情報 |
| `article:section` | 文字列 | 記事のセクション（カテゴリ） |
| `article:tag` | 文字列 | 記事のタグ |

`article:author` に正しくリソース（通常は URI）を指定している Web ページを見たことがない（笑） みんな名前を文字列で書いてるっぽい。
私も `profile` の語彙でわざわざメタデータを書く気にならないので，名前の文字列をセットしている。
文法的には `content` 属性はリテラル値なので，ただの文字列でも文法上はエラーにならないのだが。

`article:author` や `article:tag` が複数ある場合は `<meta>` 要素を複数書けばよい。

```html
<meta property="article:tag" content="foo">
<meta property="article:tag" content="bar">
```

`article:section` や `article:tag` には文字列を指定するとあるが，それって意味なくね？ 普通はセクションやタグの情報を示すリソースを示すものなんじゃないのか？

## OGP は，やっぱり OGP だった

以上を踏まえて，このブログサイトも `<head>` 要素に [OGP] の `article` メタデータを追加している。
まぁ `article` メタデータがどのくらい活用されているのかは分からないけど。
しばらくは様子見ということで。

[OGP] は（RDFa に則ってるのに）相変わらずセマンティック・ウェブからかけ離れたところにいるのがいい感じである（笑）

## ブックマーク

- [RDFa 1.1 Primer - Third Edition](https://www.w3.org/TR/rdfa-primer/)
- [RDFa Core 1.1 - Third Edition](https://www.w3.org/TR/rdfa-core/)
- [HTML+RDFa 1.1 - Second Edition](https://www.w3.org/TR/html-rdfa/)
- [アナタの知らないOGPタグ](https://webtech.fukushimaku.jp/kiji/ogp-open-graph-protocol-knowledge.html)

- [Twitter Card メタデータに対応した]({{< ref "/remark/2017/10/twitter-card-metadata.md" >}})


[Kagi Search]: https://kagi.com/ "Kagi Search - A Premium Search Engine"
[Kagi Translate]: https://translate.kagi.com/ "Kagi Translate"
[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"
[GitHub Copilot]: https://github.com/features/copilot "GitHub Copilot · Your AI pair programmer · GitHub"
[Proton Lumo]: https://lumo.proton.me/ "Lumo: Privacy-first AI assistant where chats stay confidential"

[OGP]: https://ogp.me/ "The Open Graph protocol"
<!-- eof -->
