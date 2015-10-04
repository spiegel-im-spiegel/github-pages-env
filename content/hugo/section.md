+++
title       = "Categories, Tags そして Section"
description = "前回の続き。今回は Categories, Tags そして Section について書いてみる。"
date        = "2015-09-11T17:58:32+09:00"
update      = "2015-10-04T19:54:00+09:00"
tags        = [ "hugo", "categories", "tags", "taxonomy", "section" ]
draft = false

[author]
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
+++

（初出： [ゼロから始める Hugo — Categories, Tags そして Section - Qiita](http://qiita.com/spiegel-im-spiegel/items/4c5859f7cac877068742)）

[前回]の続き。
今回は Categories, Tags そして Section について書いてみる。

## Categories と Tags{#categories-tags}

まずは Categories と Tags について。

[Hugo] では記事に Categories および Tags を設定することができる。
以下のように記述すれば良い。

```markdown
+++
date = "2015-09-05T16:40:41+09:00"
draft = false
title = "Hello!"
categories = [ "hugo" ]
tags = [ "hello", "world" ]
+++

ようこそ， [Hugo](http://gohugo.io/) の世界へ！
```

このように Categories および Tags のキーワードを配列で列挙する（キーワードがひとつでも配列に入れること）。
これをビルドすると以下のようになる。

```
C:\hugo-env\www>hugo
0 draft content
0 future content
1 pages created
0 paginator pages created
2 tags created
1 categories created
in 20 ms

C:\hugo-env\www>tree /f .
C:\HUGO-ENV\WWW
│  config.toml
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

出力先に `categories` および `tags` フォルダが作成され，更にその下にキーワードのフォルダが作成されているのがおわかりだろうか。
キーワードのフォルダの `index.xml` は feed である。
つまり [Hugo] ではカテゴリ・タグ毎に自動で feed が作成される。

`index.html` は（テンプレートがないため）この時点では空である。テンプレートは `layouts/_default` フォルダに `list.html` ファイルを配置する。
名前からして Categories/Tags 毎に記事を列挙することを期待しているわけやね（笑） とりあえず中身はこんな感じでどうだろう。

```html
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
	<li><a href="{{ .Permalink }}">{{ .Title }}</a> (<time>{{ .Date.Format "2006-01-02" }}</time>){{ if .Draft }} #Draft{{ end }}</li>
{{ end }}
</ul>

</body>
<html>
```

ビルド結果はこんな感じ。

```html
<!DOCTYPE html>
<html lang="ja">
<head>
<meta charset="utf-8">
<title>List of Hugo -- Hello World!</title>
</head>
<body>
<h1>List of Hugo</h1>

<ul style="list-style:none;">

	<li><a href="http://hello.example.com/hello/">Hello!</a> (<time>2015-09-05</time>)</li>

</ul>

</body>
<html>
```

おおう。キーワードの頭文字が勝手に大文字に変換されてるぜ。

[前回]を読んでいる人はトップページのテンプレート `layouts/index.html` とほぼ同じ構成であることに気づくと思う。
違うのは `{{ .Title }}` には Categories/Tags のキーワードが入ることと `{{ range }}` 構文の対象変数が `.Site.Pages` ではなく `.Data.Pages` であることだ。

ついでに記事ページで Categories/Tags を表示できるようにしてみよう。
`layouts/_default/single.html` を以下のように記述する。

```html
<!DOCTYPE html>
{{ with .Site.LanguageCode }}<html lang="{{ . }}">{{ else }}<html>{{ end }}
<head>
<meta charset="utf-8">
<title>{{ .Title }} -- {{ .Site.Title }}</title>
</head>
<body>
<h1>{{ .Title }}</h1>
<nav>
	{{ with .Params.categories }}<div>Categories:{{ range . }} <a href="/categories/{{ . | urlize }}/">{{ . }}</a>{{ end }}</div>{{ end }}
	{{ with .Params.tags }}<div>Tags:{{ range . }} <a href="/tags/{{ . | urlize }}/">#{{ . }}</a>{{ end }}</div>{{ end }}
</nav>

<div>{{ .Content }}</div>
</body>
<html>
```

以下がビルド結果。

```html
<!DOCTYPE html>
<html lang="ja">
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

`{{ with }}` 構文の中に `{{ range }}` 構文が入ってて分かりにくいが，変数のスコープに注意すれば，それほど難しくはないはず。
注意しないといけないのは， Categories/Tags の変数名が `.Categories`, `.Tags` ではなく `.Params.categories`, `.Params.tags` になっている点である。

`{{ . | urlize }}` というのはパイプ機能の一種で， `urlize` であれば左側の値（文字列）を URL として安全な文字列に変換してくれる。

Categories と Tags との間に機能上の違いはない。
名前が違うだけである。
おそらく他のブログサービスとの互換性の為にあるのだろうが，「ゼロから始める」のであれば Categories と Tags を併記することに意味はない。
それなら後述する Section と組み合わせるほうが合理的である。

### .Params のルール{#params}

Front matter で指定する変数は，「[テンプレート変数](http://gohugo.io/templates/variables/)」にある既定のもの以外は `.Params` 以下に自動的に組み換えられる。
なおかつ `.Params` 以下の変数名は小文字になる決まりである。
Categories/Tags は標準機能なのだが，どういうわけかこれだけ `.Params` 以下に組み替えられる。
なんだかなぁ。「歴史的経緯」ってやつだろうか[^a]。

[^a]: Categories/Tags の配置が特殊なのは，これらが [Taxonomy]({{ relref "#taxonomy" }}) として実装されているからのようだ。

ちなみに `config.toml` によるサイト設定では `.Site.Params` への暗黙的な組み換えは行われないため，明示的に記述する必要がある。

```toml
[params]
author = "Spiegel"
```

この非対称性も分かりにくいんだよなぁ。

### Taxonomy{#taxonomy}

[Hugo] には Taxonomy と呼ばれる機能があって，標準では Categories/Tags のリストを取り出すことができる[^b]。
たとえば， `layouts/index.html` をこんな感じに書く。

[^b]: Taxonomy の項目は Categories/Tags 以外にも任意に設定することができる（{{% quote lang="en" %}}[Using Taxonomies](https://gohugo.io/taxonomies/usage/){{% /quote %}} 参照）。

```html
<h2>Taxonomy Terms</h2>
<ul>{{ range $taxonomyname, $taxonomy := .Site.Taxonomies }}
    <li>{{ $taxonomyname }}
        <ul>{{ range $key, $value := $taxonomy }}
            <li><a href="/{{ $taxonomyname | urlize }}/{{ $key | urlize }}">{{ $key }}</a> ({{ $value.Count }})
                <ul>{{ range $value.Pages }}
                    <li><a href="{{ .Permalink }}">{{ .LinkTitle }}</a> </li>
                {{ end }}</ul></li>
        {{ end }}</ul></li>
{{ end }}</ul>
```

これをビルドするとこんな感じに展開される。

```html
<h2>Taxonomy Terms</h2>
<ul>
    <li>categories
        <ul>
            <li><a href="/categories/hugo">hugo</a> (1)
                <ul>
                    <li><a href="http://hello.example.com/hello/">Hello!</a> </li>
                </ul></li>
        </ul></li>

    <li>tags
        <ul>
            <li><a href="/tags/hello">hello</a> (1)
                <ul>
                    <li><a href="http://hello.example.com/hello/">Hello!</a> </li>
                </ul></li>

            <li><a href="/tags/world">world</a> (1)
                <ul>
                    <li><a href="http://hello.example.com/hello/">Hello!</a> </li>
                </ul></li>
        </ul></li>
</ul>
```

Tags の一覧のみを取得したいのであれば，もっと簡単に

```html
<h2>Tags</h2>
<ul>{{ range $key, $value := .Site.Taxonomies.tags.ByCount }}
	<li>#<a href="/tags/{{ $key | urlize }}">{{ $key }}</a> ({{ $value.Count }})</li>
{{ end }}</ul>
```

でもよい。
ちなみに `.Site.Taxonomies.tags.ByCount` は Tags の一覧をタグされている記事数の順で sort したものである。
アルファベット順にするには `.Site.Taxonomies.tags.Alphabetical` とする。

## Section{#section}

`content` フォルダの下に `practice` というフォルダを作り，ここに `hello.md` を移動させてみよう。
新たに作る場合は path 付きで作成すればよい。

```
C:\hugo-env\www>hugo new practice/hello.md
C:\hugo-env\www\content\practice\hello.md created
```

これでビルドしてみる（ファイルを移動した際は出力フォルダの中をいったんクリーンにしてからビルドするとゴミが残らない）。

```
C:\hugo-env\www>hugo
0 draft content
0 future content
1 pages created
0 paginator pages created
2 tags created
1 categories created
in 19 ms

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
│  └─_default
│          list.html
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
│  ├─practice
│  │  │  index.html
│  │  │  index.xml
│  │  │
│  │  └─hello
│  │          index.html
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

`hello/index.html` が `practice/hello/index.html` に配置されるのは予想通りだと思うが， `practice` に `index.html` と `index.xml` が生成されているのがおわかりだろうか。
`practice/index.html` の中身はこんな感じ。

```html
<!DOCTYPE html>
<html lang="ja">
<head>
<meta charset="utf-8">
<title>List of Practices -- Hello World!</title>
</head>
<body>
<h1>List of Practices</h1>

<ul style="list-style:none;">

	<li><a href="http://hello.example.com/practice/hello/">Hello!</a> (<time>2015-09-05</time>)</li>

</ul>

</body>
<html>
```

これは `layouts/_default/list.html` テンプレートで生成されたページだ。
[Hugo] ではフォルダ付きの記事を作成すると，そのフォルダが Section として機能する。

当然だが，ひとつの記事はひとつの Section にしか帰属できない。
これは先の Categories/Tags との大きな違いである。
Section と Categories/Tags を組み合わせれば縦串と横串で記事を指示できるようになる。

ついでに記事ページで Section を表示できるようにしてみよう。
たとえば， `layouts/_default/single.html` をこんな感じに書く。

```html
<!DOCTYPE html>
{{ with .Site.LanguageCode }}<html lang="{{ . }}">{{ else }}<html>{{ end }}
<head>
<meta charset="utf-8">
<title>{{ .Title }}{{ with .Section }} -- {{ . }}{{ end }} -- {{ .Site.Title }}</title>
</head>
<body>
<h1>{{ .Title }}{{ with .Section }} [<a href="/{{ . | urlize }}/">{{ . }}</a>]{{ end }}</h1>
<nav>
	{{ with .Params.categories }}<div>Categories:{{ range . }} <a href="/categories/{{ . | urlize }}/">{{ . }}</a>{{ end }}</div>{{ end }}
	{{ with .Params.tags }}<div>Tags:{{ range . }} <a href="/tags/{{ . | urlize }}/">#{{ . }}</a>{{ end }}</div>{{ end }}
</nav>

<div>{{ .Content }}</div>
</body>
<html>
```

以下がビルド結果。

```html
<!DOCTYPE html>
<html lang="ja">
<head>
<meta charset="utf-8">
<title>Hello! -- practice -- Hello World!</title>
</head>
<body>
<h1>Hello! [<a href="/practice/">practice</a>]</h1>
<nav>
	<div>Categories: <a href="/categories/hugo/">hugo</a></div>
	<div>Tags: <a href="/tags/hello/">#hello</a> <a href="/tags/world/">#world</a></div>
</nav>

<div><p>ようこそ， <a href="http://gohugo.io/">Hugo</a> の世界へ！</p>
</div>
</body>
<html>
```

記事のフォルダ階層はいくらでも深くできるが， Section として認識されるのは直下のフォルダのみのようである。
たとえば `content/practice/hello.md` を `content/practice/firstcode/hello.md` に移動してビルドすると（出力フォルダはクリーンアップしてね）

```
C:\hugo-env\www>tree /f .
C:\HUGO-ENV\WWW
│  config.toml
│
├─archetypes
├─content
│  └─practice
│      └─firstcode
│              hello.md
│
├─data
├─layouts
│  │  404.html
│  │  index.html
│  │
│  └─_default
│          list.html
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
│  ├─practice
│  │  │  index.html
│  │  │  index.xml
│  │  │
│  │  └─firstcode
│  │      └─hello
│  │              index.html
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

```html
<!DOCTYPE html>
<html lang="ja">
<head>
<meta charset="utf-8">
<title>Hello! -- practice -- Hello World!</title>
</head>
<body>
<h1>Hello! [<a href="/practice/">practice</a>]</h1>
<nav>
	<div>Categories: <a href="/categories/hugo/">hugo</a></div>
	<div>Tags: <a href="/tags/hello/">#hello</a> <a href="/tags/world/">#world</a></div>
</nav>

<div><p>ようこそ， <a href="http://gohugo.io/">Hugo</a> の世界へ！</p>
</div>
</body>
<html>
```

というわけで，あくまでも `practice` フォルダが Section になっているのがわかると思う。

### Section ごとのカスタマイズ{#section-custom}

[Hugo] では Section ごとにカスタマイズすることができる。
`layouts` フォルダに `section` フォルダを作成し，その中に `<section name>.html` ファイルを作成すると，そのテンプレートで Section のトップページ（`<section name>/index.html`）を作成する。
今回は `layouts/section/practice.html` を作成してみる。

```html
<!DOCTYPE html>
{{ with .Site.LanguageCode }}<html lang="{{ . }}">{{ else }}<html>{{ end }}
<head>
<meta charset="utf-8">
<title>Hugo の練習 -- {{ .Site.Title }}</title>
</head>
<body>
<h1>Hugo の練習</h1>

<ul style="list-style:none;">
{{ range .Data.Pages }}
	<li><a href="{{ .Permalink }}">{{ .Title }}</a> (<time>{{ .Date.Format "2006-01-02" }}</time>){{ if .Draft }} #Draft{{ end }}</li>
{{ end }}
</ul>

</body>
<html>
```

ビルド結果。変わり映えしなくてすみません。

```html
<!DOCTYPE html>
<html lang="ja">
<head>
<meta charset="utf-8">
<title>Hugo の練習 -- Hello World!</title>
</head>
<body>
<h1>Hugo の練習</h1>

<ul style="list-style:none;">

	<li><a href="http://hello.example.com/practice/hello/">Hello!</a> (<time>2015-09-05</time>)</li>

</ul>

</body>
<html>
```

更に Section 内の記事ページもカスタマイズできる。
これは `layouts` フォルダに Section 名のフォルダを作成し，その中に `single.html` を配置する。

```html
<!DOCTYPE html>
{{ with .Site.LanguageCode }}<html lang="{{ . }}">{{ else }}<html>{{ end }}
<head>
<meta charset="utf-8">
<title>{{ .Title }} -- Hugo の練習 | {{ .Site.Title }}</title>
</head>
<body>
<h1>{{ .Title }} -- Hugo の練習</h1>
<nav>
	{{ with .Params.categories }}<div>Categories:{{ range . }} <a href="/categories/{{ . | urlize }}/">{{ . }}</a>{{ end }}</div>{{ end }}
	{{ with .Params.tags }}<div>Tags:{{ range . }} <a href="/tags/{{ . | urlize }}/">#{{ . }}</a>{{ end }}</div>{{ end }}
</nav>

<div>{{ .Content }}</div>
</body>
<html>
```

```html
<!DOCTYPE html>
<html lang="ja">
<head>
<meta charset="utf-8">
<title>Hello! -- Hugo の練習 | Hello World!</title>
</head>
<body>
<h1>Hello! -- Hugo の練習</h1>
<nav>
	<div>Categories: <a href="/categories/hugo/">hugo</a></div>
	<div>Tags: <a href="/tags/hello/">#hello</a> <a href="/tags/world/">#world</a></div>
</nav>

<div><p>ようこそ， <a href="http://gohugo.io/">Hugo</a> の世界へ！</p>
</div>
</body>
<html>
```

これも変わり映えしなくてすみません。

しかし，なんでこんなテンプレートの構成なんだろう。
これだと `section` や `_default` という名前のセクションはカスタマイズ出来ないことになる（まだ説明してないが `layouts` フォルダには他にも `partials` や `shortcodes` といったフォルダもある）。
そうではなくて，  `layouts/section` フォルダの下にセクション名のフォルダを掘ってその中に `list.html` や `single.html` を配置すればスッキリするのに。

さきほどの `.Params` の話といい，どうも [Hugo] は名前管理がいきあたりばったりな気がする。

### categories/hugo という名前の記事はどうなるの？{#conflict}

これも名前に関する話。

`categories/hugo.md` という名前の記事は `categories/hugo/index.html` に展開される。
これって Categories の機能と丸かぶりである。
実は [Hugo] では名前が衝突した際の挙動は明文化されていない（筈）。
強いて言うなら実装依存で状況依存である。またビルド時にエラーになることもない。

```
C:\hugo-env\www>hugo new categories/hugo.md
C:\hugo-env\www\content\categories\hugo.md created

C:\hugo-env\www>hugo undraft content/categories/hugo.md

C:\hugo-env\www>hugo
0 draft content
0 future content
2 pages created
0 paginator pages created
2 tags created
1 categories created
in 24 ms

C:\hugo-env\www>tree /f .
C:\HUGO-ENV\WWW
│  config.toml
│
├─archetypes
├─content
│  ├─categories
│  │      hugo.md
│  │
│  └─practice
│          hello.md
│
├─data
├─layouts
│  │  404.html
│  │  index.html
│  │
│  └─_default
│          list.html
│          single.html
│
├─public
│  │  404.html
│  │  index.html
│  │  index.xml
│  │  sitemap.xml
│  │
│  ├─categories
│  │  │  index.html
│  │  │  index.xml
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

上の場合は `categories/hugo/index.html` が記事ページになった。
他にも，このフォルダ構成なら `practice.md` も名前が衝突する。
このような衝突を避けるにはユーザ側で名前を管理するしかない。
小規模なサイトなら人間が気をつければいいが，中規模以上で複数人が関わるようになると結構危ないかもしれない。

## ブックマーク{#bookmark}

- [Hugoサイト構築 | Watanabe-DENKI Inc. 渡辺電気株式会社](http://wdkk.co.jp/lab/hugo/) : お勧め！

[Hugo に関するブックマークはこちら]({{< ref "hugo/bookmark.md" >}})。

[Hugo]: http://gohugo.io/ "Hugo :: A fast and modern static website engine"
[前回]: {{< ref "hugo/hello.md" >}} "インストールから Hello World まで"
