+++
date = "2015-09-25T11:14:44+09:00"
description = "今回は Template について，もう少しだけ詳しく紹介してみる。"
draft = false
tags = [ "hugo", "template" ]
title = "Template の部品化"

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

今回は Template について，もう少しだけ詳しく紹介してみる。

## Theme “hugo-theme-text” を導入する

[前回]予告した通り，拙作の Theme [spiegel-im-spiegel/hugo-theme-text] を導入する。

```
C:\hugo-env\www>git clone https://github.com/spiegel-im-spiegel/hugo-theme-text.git themes/hugo-theme-text
Cloning into 'themes/hugo-theme-text'...
remote: Counting objects: 174, done.
remote: Compressing objects: 100% (133/133), done.
Receiving objecemote: Total 174 (delta 82), reused 123 (delta 35), pack-reused 0
Receiving objects:  49%
Receiving objects: 100% (174/174), 25.34 KiB | 0 bytes/s, done.
Resolving deltas: 100% (82/82), done.
Checking connectivity... done.
```

作業環境が git 管理下にある場合は submodule として導入するとよい。

```
C:\hugo-env\www>git submodule add https://github.com/spiegel-im-spiegel/hugo-theme-text.git themes/hugo-theme-text
Cloning into 'themes/hugo-theme-text'...
remote: Counting objects: 282, done.
remote: Compressing objects: 100% (11/11), done.
rRemote: Total 282 (delta 4), reused 0 (delta 0), pack-reused 271eceiving object
Receiving objects: 100% (282/282), 37.54 KiB | 0 bytes/s, done.
Resolving deltas: 100% (141/141), done.
Checking connectivity... done.
```

では，例によっていきなりビルド。
の前に，作業環境の `layouts` フォルダを空っぽにする。

ではビルド。

```
C:\hugo-env\www>hugo --theme="hugo-theme-text"
0 draft content
0 future content
1 pages created
0 paginator pages created
2 tags created
1 categories created
in 21 ms
```

これでこんな感じの `index.html` ファイルができれば成功。

```html
<!DOCTYPE html>
<html lang="jp">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="generator" content="Hugo 0.15-DEV" />

<link rel="stylesheet" href="//fonts.googleapis.com/css?family=Droid+Serif:400,400italic,700,700italic|Open+Sans:400,400italic,700,700italic|Inconsolata:400,700&subset=latin,latin-ext">
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.4.0/css/font-awesome.min.css">
<link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/cc-icons/1.2.1/css/cc-icons.min.css">
<link rel="stylesheet" href="/css/main.css">
<link rel="stylesheet" href="/css/text-baldanders.css">

<link rel="alternate" href="/index.xml" type="application/rss+xml" title="Hello World!">

<title>Hello World!</title>

</head>
<body>



<div id='container'>
<header role="banner">
<h1><a href="http://hello.example.com/">Hello World!</a> <a href='/index.xml' title='Feed'><i class='fa fa-rss'></i></a></h1>

</header>


<main role="main">
<article itemscope itemtype='http://schema.org/Blog' itemref='maker'>
<h2 itemprop='name'>All Entries</h2>
<ul style="list-style:none;"><li itemprop='blogPost' itemscope itemtype='http://schema.org/BlogPosting'>
	[<a href="/practice">practice</a>] <a href="http://hello.example.com/practice/hello/" itemprop='url'><strong itemprop='name'>Hello!</strong></a>
	<span style="font-size:smaller;">(<time>2015-09-05</time>)
	 #<a href="/tags/hello">hello</a> #<a href="/tags/world">world</a>
	</span>
</li>
</ul>
</article>
</main>

<footer>

<nav role="navigation">
<ul class='cloud center'>
<li>Powered by <a href='http://gohugo.io/'>Hugo 0.15-DEV</a> and <a href="https://github.com/spiegel-im-spiegel/hugo-theme-text">Text</a>.</li>
</ul>
</nav>

</footer>

</div>

</body>
</html>
```

だいぶ賑やかになった。

この `index.html` のテンプレートはこんな感じになっている。

```html
{{ partial "top.html" . }}
<head>
{{ partial "header.html" . }}
<title>{{ .Site.Title }}</title>
{{ with .Site.Params.description }}<meta name="description" content="{{ . }}">{{ end }}
</head>
<body>
{{ partial "prepare.html" . }}
<div id='container'><!-- Container Area -->
{{ partial "banner.html" . }}

<main role="main">
<article itemscope itemtype='http://schema.org/Blog' itemref='maker'>
<h2 itemprop='name'>All Entries</h2>
<ul style="list-style:none;">{{ range .Site.Pages }}{{ partial "li.html" . }}{{ end }}</ul>
</article>
</main>

{{ partial "footer.html" . }}
</div><!-- Container Area -->
{{ partial "cleanup.html" . }}
</body>
{{ partial "end.html" . }}
```

この中で `{{ partial "top.html" . }}` となっている部分が partial template を呼び出している部分で，テンプレートの中にテンプレートを入れ子にすることができる。

### Partial Template

Partial template を格納する場所は決まっている。
では， [spiegel-im-spiegel/hugo-theme-text] の構成を見てみよう。

```
C:\hugo-env\www>tree /f themes\hugo-theme-text
C:\HUGO-ENV\WWW\THEMES\HUGO-THEME-TEXT
│  LICENSE
│  README.md
│  theme.toml
│
├─archetypes
│      default.md
│
├─layouts
│  │  index.html
│  │  rss.xml
│  │
│  ├─partials
│  │      banner.html
│  │      cc-license.html
│  │      ccl-rss.html
│  │      cleanup.html
│  │      end.html
│  │      feedback.html
│  │      footer.html
│  │      header-include.html
│  │      header.html
│  │      li.html
│  │      navi.html
│  │      prepare.html
│  │      social.html
│  │      top.html
│  │
│  └─_default
│          list.html
│          single.html
│
└─static
    ├─css
    │      main.css
    │      text-baldanders.css
    │
    └─images
            avatar.png
```

`layouts/partials` フォルダに多くのファイルが格納されているのがお分かりだろうか。
これが partial template である。

たとえば `layouts/partials/top.html` の中身は以下のようになっている。

```html
<!DOCTYPE html>
{{ with .Site.LanguageCode }}<html lang="{{ . }}">{{ else }}<html>{{ end }}
```

また `layouts/partials/header.html` は

```html
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
{{ .Hugo.Generator }}
{{ with .Site.Params.favicon }}<link rel="icon" type="image/x-icon" href="{{ . }}">{{ end }}
<link rel="stylesheet" href="//fonts.googleapis.com/css?family=Droid+Serif:400,400italic,700,700italic|Open+Sans:400,400italic,700,700italic|Inconsolata:400,700&subset=latin,latin-ext">
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.4.0/css/font-awesome.min.css">
<link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/cc-icons/1.2.1/css/cc-icons.min.css">
<link rel="stylesheet" href="/css/main.css">
<link rel="stylesheet" href="/css/text-baldanders.css">
{{ partial "header-include.html" . }}
<link rel="alternate" href="/index.xml" type="application/rss+xml" title="{{ .Site.Title }}">
```

となっていて，内部でさらに `layouts/partials/header-include.html` を呼び出している。
実は `layouts/partials/header-include.html` の中身は空である。
ユーザの作業環境側に `layouts/partials/header-include.html` ファイルを設置すればそちらが優先して読み込まれるため，これでユーザ側がカスタマイズできるようになっている。

### 【おまけ】 Author 情報

[spiegel-im-spiegel/hugo-theme-text] では `config.toml` に以下の author 情報を入れるとページに反映される。

```toml
[params.author]
name      = "Spiegel"
url       = "http://www.baldanders.info/spiegel/profile/"
avatar    = "/images/avatar.jpg"
license   = "by-sa"
github    = "spiegel-im-spiegel"
twitter   = "spiegel_2007"
medium    = "@spiegel"
instagram = "spiegel_2007"
facebook  = "spiegel.im.spiegel"
linkedin  = "spiegelimspiegel"
flattr    = "spiegel"
```

{{< fig-img src="https://farm6.staticflickr.com/5779/21504929379_e29706db7c.jpg" alt="Top Page of “hugo-theme-text”" title="Top Page of “hugo-theme-text”" link="https://www.flickr.com/photos/spiegel/21504929379/" >}}

記事ページでは front matter に `[author]` 項目を加える事で front matter の値が優先して表示される。

```toml
+++
date = "2015-09-05T16:40:41+09:00"
draft = false
title = "Hello!"
categories = [ "hugo" ]
tags = [ "hello", "world" ]

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
```

{{< fig-img src="https://farm1.staticflickr.com/718/21069070064_4a0e51cbbd.jpg" alt="Entry Page of “hugo-theme-text”" title="Entry Page of “hugo-theme-text”" link="https://www.flickr.com/photos/spiegel/21069070064/" >}}

他にも favicon や Disqus などを設置可能。
詳しくは [`README.md`](https://github.com/spiegel-im-spiegel/hugo-theme-text#simple-text-hugo-theme) を参照のこと。

## ブックマーク{#bookmark}

- [Hugoサイト構築 | Watanabe-DENKI Inc. 渡辺電気株式会社](http://wdkk.co.jp/lab/hugo/) : お勧め！

[Hugo に関するブックマークはこちら]({{< ref "hugo/bookmark.md" >}})。

[Hugo]: http://gohugo.io/ "Hugo :: A fast and modern static website engine"
[spiegel-im-spiegel/hugo-theme-text]: https://github.com/spiegel-im-spiegel/hugo-theme-text
[前回]: {{< ref "hugo/theme.md" >}} "Theme を利用する"
