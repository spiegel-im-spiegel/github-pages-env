+++
date = "2015-10-04T20:21:54+09:00"
update      = "2016-12-14T10:20:58+09:00"
description = "以前の回で紹介してなかったのだが， Categories や Tags の一覧ページを作る機能があるので紹介する。"
draft = false
tags = [ "hugo", "categories", "tags" ]
title = "Categories や Tags の一覧ページを作る"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

今回は補足なので簡単に。

「[Categories, Tags そして Section]({{< relref "hugo/section.md" >}})」の回で紹介してなかったのだが[^a]， Categories や Tags の一覧ページを作る機能があるので紹介する。

[^a]: というか，今までやり方が分かってなかった。

Categories や Tags の一覧ページを作るには `layouts/_default` フォルダに `terms.html` テンプレートを作成する[^b]。
たとえば中身はこんな感じ。

[^b]: Tags や Categories などごとに個別にテンプレートを作りたいのであれば `layouts/taxonomy/tag.terms.html` といった感じでファイルを作るとよい。

```html
<!DOCTYPE html>
{{ with .Site.LanguageCode }}<html lang="{{ . }}">{{ else }}<html>{{ end }}
<head>
<meta charset="utf-8">
<title>{{ .Title }} Cloud -- {{ .Site.Title }}</title>
<style>
ul.cloud {
    list-style: none;
    padding: 0
}
ul.cloud > li {
    display: inline-block;
    margin: 0 0.5rem;
}
</style>
</head>
<body>
<h1>{{ .Title }} Cloud</h1>

<ul class="cloud">{{ $plural := .Data.Plural }}
{{ range $key, $value := .Data.Terms }}
    <li><a href="/{{ $plural }}/{{ $key | urlize }}">{{ $key }}</a> ({{ len $value }})</li>
{{ end }}
</ul>

</body>
<html>
```

[以前の記事]({{< relref "hugo/section.md" >}})の [Taxonomy]({{< relref "hugo/section.md#" >}}) の記述に似ているので，それほど難しくないだろう。
これをビルドするとこんな感じにファイルが展開される。

```text
C:\hugo-env\www>hugo
0 draft content
0 future content
1 pages created
0 paginator pages created
2 tags created
1 categories created
in 17 ms

C:\hugo-env\www>tree /f .
C:\HUGO-ENV\WWW
│  config.toml
│
├─archetypes
├─content
│  └─practice
│          hello.md
│
├─data
├─layouts
│  │  404.html
│  │  index.html
│  │
│  ├─practice
│  │      single.html
│  │
│  ├─section
│  │      practice.html
│  │
│  └─_default
│          list.html
│          single.html
│          terms.html
│
├─public
│  │  404.html
│  │  index.html
│  │  index.xml
│  │  sitemap.xml
│  │
│  ├─categories
│  │  │  index.html
│  │  │
│  │  └─hugo
│  │          index.html
│  │          index.xml
│  │
│  ├─practice
│  │  │  index.html
│  │  │  index.xml
│  │  │
│  │  └─hello
│  │          index.html
│  │
│  └─tags
│      │  index.html
│      │
│      ├─hello
│      │      index.html
│      │      index.xml
│      │
│      └─world
│              index.html
│              index.xml
│
└─static
```

`public/categories` および `public/tags` フォルダの直下に `index.html` ファイルが生成されているのがお分かりだろうか。
また `public/tags/index.html` ファイルの中身を見てみると以下のように Tags 情報が展開されている。

```html
<!DOCTYPE html>
<html lang="jp">
<head>
<meta charset="utf-8">
<title>Tags Cloud -- Hello World!</title>
<style>
ul.cloud {
    list-style: none;
    padding: 0
}
ul.cloud > li {
    display: inline-block;
    margin: 0 0.5rem;
}
</style>
</head>
<body>
<h1>Tags Cloud</h1>

<ul class="cloud">

    <li><a href="/tags/hello">hello</a> (1)</li>

    <li><a href="/tags/world">world</a> (1)</li>

</ul>

</body>
<html>
```

## ブックマーク{#bookmark}

- [Taxonomy Terms Template](https://gohugo.io/templates/terms/)
- [Hugoサイト構築 | Watanabe-DENKI Inc. 渡辺電気株式会社](http://wdkk.co.jp/lab/hugo/) : お勧め！
- [Hugoでタグやカテゴリのリンク切れが起こる - Qiita](http://qiita.com/_shun_sato_/items/87888fa8425e55b1c758)

[Hugo に関するブックマークはこちら]({{< ref "hugo/bookmark.md" >}})。

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
