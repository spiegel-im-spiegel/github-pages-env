+++
title       = "ゼロから始める Hugo — Section と Pagination"
description = "Hugo に関するメモ2回目。今回は Section と Pagination について。"
date        = "2015-09-09T00:00:00"
tags        = [ "hugo", "section", "pagination" ]
slug        = "hugo-section"

[author]
name        = "Spiegel"
profileurl  = "http://www.baldanders.info/spiegel/profile/"
avatar      = "/images/avatar.jpg"
license     = "by-sa"
github      = "spiegel-im-spiegel"
twitter     = "spiegel_2007"
medium      = "@spiegel"
instagram   = "spiegel_2007"
facebook    = "spiegel.im.spiegel"
linkedin    = "spiegelimspiegel"
flattr      = "spiegel"
+++

Hugo に関するメモ。[前回]に続き2回目。今回は Section と Pagination について。

## Categories と Tags

その前に categories と tags に触れておく。

[Hugo] では記事にカテゴリとタグを指定できる。これらの指定は front matter で行う。こんな感じ。

```markdown:content\hello.md
+++
date = "2015-09-05T16:40:41+09:00"
draft = false
title = "Hello!"
categories = [ "hugo" ]
tags = [ "hello", "world" ]
+++

ようこそ， [Hugo](http://gohugo.io/) の世界へ！
```

早速ビルド。

```shell
C:\hugo-env\www>hugo
0 draft content
0 future content
1 pages created
0 paginator pages created
2 tags created
1 categories created
in 14 ms

C:\hugo-env\www>tree /f .
C:\HUGO-ENV\WWW
│  .editorconfig
│  .gitignore
│  config.toml
│  README.md
│
├─archetypes
├─content
│      hello.md
│
├─data
├─layouts
│  │  404.html
│  │  index.html
│  │
│  └─_default
│          single.html
│
├─public
│  │  404.html
│  │  index.html
│  │  index.xml
│  │  sitemap.xml
│  │
│  ├─categories
│  │  └─hugo
│  │          index.html
│  │          index.xml
│  │
│  ├─hello
│  │      index.html
│  │
│  └─tags
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

`categories` フォルダと `tags` フォルダが作成され，それぞれの中にキーワードのフォルダも作成されていることがお分かりだろうか。各フォルダには `index.html` と `index.xml` が作成されている。つまり，カテゴリやタグのキーワードごとに feed が作成される。

この時点では `index.html` は（テンプレートがないため）空である。これらのテンプレートは `layouts/_default` フォルダに `list.html` を置く。名前からしてカテゴリ・タグごとのリストを表示するページを作るよう期待されているようだ（笑）

```html:layouts/_default/single.html
<!DOCTYPE html>
{{ with .Site.LanguageCode }}<html lang="{{ . }}">{{ else }}<html>{{ end }}
<head>
<meta charset="utf-8">
<title>List of {{ .Title }} -- {{ .Site.Title }}</title>
</head>
<body>
<h1>List of {{ .Title }}</h1>

<ul style="list-style:none;">
{{ range .Data.Pages }}
	<li><a href="{{ .Permalink }}">{{ .Title }}</a> (<time pubdate="{{ .Date.Format "2006-01-02" }}">{{ .Date.Format "2006-01-02" }}</time>){{ if .Draft }} #Draft{{ end }}</li>
{{ end }}
</ul>

</body>
<html>
```

これをビルドした結果はこんな感じ。

```html:public/categories/hugo/index.html
<!DOCTYPE html>
<html lang="jp">
<head>
<meta charset="utf-8">
<title>List of Hugo -- Hello World!</title>
</head>
<body>
<h1>List of Hugo</h1>

<ul style="list-style:none;">

	<li><a href="http://hello.example.com/hello/">Hello!</a> (<time pubdate="2015-09-05">2015-09-05</time>)</li>

</ul>

</body>
<html>
```

おおっ。キーワードの頭文字が勝手に大文字になっている。

[前回]の記事を読んだ方は `layouts/_default/single.html` の構成が `layouts/index.html` とほとんど同じだということに気付かれたと思う。違いは， `{{ .Title }}` にはサイトのタイトルではなくカテゴリ・タグのキーワード名が入ること， `{{ range }}` 構文の対象が `.Site.Pages` ではなく `.Data.Pages` になっていることである。

ついでなので，記事中にカテゴリ・タグを表示できるようにしておこう。たとえばテンプレートはこんな感じ。

```html:layouts/_default/single.html
<!DOCTYPE html>
{{ with .Site.LanguageCode }}<html lang="{{ . }}">{{ else }}<html>{{ end }}
<head>
<meta charset="utf-8">
<title>{{ .Title }} -- {{ .Site.Title }}</title>
</head>
<body>
<h1>{{ .Title }}</h1>
<nav>
	{{ with .Params.categories }}<div>Categories:{{ range . }} <a href="/categories/{{ . }}/">{{ . }}</a>{{ end }}</div>{{ end }}
	{{ with .Params.tags }}<div>Tags:{{ range . }} <a href="/tags/{{ . }}/">#{{ . }}</a>{{ end }}</div>{{ end }}
</nav>

<div>{{ .Content }}</div>
</body>
<html>
```

これをビルドするとこんな感じになる。

```html:public/hello/index.html
<!DOCTYPE html>
<html lang="jp">
<head>
<meta charset="utf-8">
<title>Hello! -- Hello World!</title>
</head>
<body>
<h1>Hello!</h1>
<nav>
	<div>Categories: <a href="/categories/hugo/">hugo</a></div>
	<div>Tags: <a href="/tags/hello/">#hello</a> <a href="/tags/world/">#world</a></div>
</nav>

<div><p>ようこそ， <a href="http://gohugo.io/">Hugo</a> の世界へ！</p>
</div>
</body>
<html>
```

よしよし。ちゃんと小文字になってるな。

`{{ with }}` 構文の中に `{{ range }}` 構文が入ってて分かりにくいけど，変数のスコープに注意すればそれほど難しい記述じゃないはず。注意しないといけないのはカテゴリ・タブの変数が `.Categories`, `.Tags` とかではなく `.Params` の配下に入っていることかな。

ここまで見て分かると思うけど， [Hugo] ではカテゴリとタグの間に機能的な違いはない。名前が違うだけだ。これはおそらく他のブログサービスから移転する際の互換用として用意されているんだろうけど，「ゼロから始める」のであれば両方使うのは無意味だと思う。それよりは，どちらか一方を後述する Section と組み合わせるほうが合理的である。

### .Params の謎

front matter では任意の変数を定義できるけど `{{ .Title }}` や `{{ .Date }}` といった標準の変数以外は `{{ .Params }}` の配下に組み替えられる。たとえば

```toml
+++
author = "Spiegel"
+++
```

と記述した場合，テンプレート上では `{{ .Params.author }}` となる。また `{{ .Params }}` 配下の変数名は必ず小文字になる決まりである（これは `{{ .Site.Params }}` の場合も同じ）。

カテゴリやタブは標準機能なのだが，なぜか変数はオプション扱いだ。なんでこんなヘンテコなことになってるんだろう。「歴史的経緯」ってやつかねぇ。

ちなみにサイト設定の変数である `{{ .Site.Params }}` の場合はこのような暗黙の組み替えは行われず，明示的に

```toml:config.toml
[params]
author = "Spiegel"
```

と書かなければならない。この非対称性も分かりにくいんだよなぁ。

### categories という名前の記事はどうなるの？

もし `categories.md` というファイル名の記事があるとすると，それは `categories/index.html` に展開される。でもこれってカテゴリの機能と丸かぶりである。実はこのように path が衝突する場合のルールは明文化されていない。しいて言うなら実装依存で状況依存である。また，このことによりビルドエラーになることはない。

```shell
C:\hugo-env\www>hugo new categories.md
C:\hugo-env\www\content\categories.md created

C:\hugo-env\www>hugo undraft content/categories.md

C:\hugo-env\www>hugo
0 draft content
0 future content
2 pages created
0 paginator pages created
2 tags created
1 categories created
in 15 ms
```

このビルドの結果 `public/categories/index.html` に何が入るかは天数演繹なみのオカルトだ（私の場合は記事が入っていた）。したがって記事の名前はユーザ自身で気を付けて管理するしかない。これは後述する Section でも同じである。


## ブックマーク

- [Hugoサイト構築 | Watanabe-DENKI Inc. 渡辺電気株式会社](http://wdkk.co.jp/lab/hugo/) : お勧め！

[Hugo]: http://gohugo.io/ "Hugo :: A fast and modern static website engine"
[前回]: http://qiita.com/spiegel-im-spiegel/items/eac7bf2a3c0fc19afcbf "ゼロから始める Hugo — インストールから Hello World まで - Qiita"
