+++
title = "ぼくがかんがえたさいきょうの Hugo Shortcodes"
date =  "2020-01-03T15:02:51+09:00"
description = "Hugo 0.62 用に自作 shortcode を整理した。以下にいくつか紹介しておこう。"
image = "/images/attention/kitten.jpg"
tags = [ "hugo", "shortcodes", "site" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今回は [Hugo] に関する小ネタをいくつかお送りする。

[Hugo] v0.60 から markdown パーサの既定が [yuin/goldmark] になった。

- [Hugo v0.60 から既定の Markdown パーサが Goldmark になったようだ]({{< ref "/release/2019/11/hugo-v0_60-with-goldmark-parser.md" >}})

[yuin/goldmark] はきちんと [CommonMark] に準拠していて，そこには好感が持てるのだが，見方を変えるとあまり融通の効かないパーサであるとも言える。
たとえば [russross/blackfriday] では設定で `target="_blank"` や `rel="nofollow"` といった属性値を `<a>` 要素に付与できたが， [yuin/goldmark] にはこれに対応するオプションがない。

## [Markdown Render Hooks]

そこで [Hugo] v0.62 で “[Markdown Render Hooks]” という仕組みが用意された（[yuin/goldmark] パーサ使用時のみ有効）。
これは，簡単にいうと画像やリンクの HTML レンダリングを横取りして自前のテンプレートを差し込むものである。

例を挙げて説明しよう。

たとえば以下のような markdown 記述がある。

```markdown
[Hugo](https://gohugo.io/ "The world’s fastest framework for building websites") はチョー簡単！
```

これに対して [yuin/goldmark] パーサを使って HTML レンダリングを行うと以下のようになる。

```html
<a href="https://gohugo.io/" title="The world’s fastest framework for building websites">Hugo</a> はチョー簡単！
```

この `<a>` 要素に属性値 `target="_blank"` を付加することを考える。

まずは [Hugo] の環境で `layouts/_default/_markup` ディレクトリを作成し，その中に `render-link.html` ファイルを置く。
`render-link.html` ファイルの中身はこんな感じでどうだろう。

```text
<a href="{{ .Destination | safeURL }}" target="_blank"{{ with .Title }} title="{{ . }}"{{ end }}>{{ .Text | safeHTML }}</a>
```

これで先程の markdown 記述は

```html
<a href="https://gohugo.io/" target="_blank" title="The world’s fastest framework for building websites">Hugo</a> はチョー簡単！
```

となった。
ちゃんと `target="_blank"` が付加されているのがお分かりだろうか。

画像も同様にカスタマイズできる。
たとえば `layouts/_default/_markup/render-image.html` ファイルを用意して

```text
<img src="{{ .Destination | safeURL }}" alt="{{ .PlainText }}"{{ with .PlainText }} title="{{ . }}"{{ end }}>
```

としておけば

```markdown
![avatar](/images/avatar.jpg)
```

という記述に対して

```html
<img src="/images/avatar.jpg" alt="avatar" title="avatar">
```

と `title` 属性を付加することができる。

なお `render-link.html` および `render-image.html` テンプレートファイルは1つのサイトについて（今のところ）ひとつずつしか定義でず，サイト全体に適用されてしまうので注意が必要だ（テンプレートの条件分け等の機能を使うことはできる）。

## .RenderString 関数

任意の文字列を markdown テキストとして HTML 形式へレンダリングする関数としては [`markdownify`] があったが，新たに [Hugo] 0.62 で [`.RenderString`] 関数が追加された。
これまでの [`markdownify`] 関数との違いは以下の通り。

- [`.RenderString`] 関数は [`Page`] 変数のメソッドとして実装されている
- パーサを指定できる（既定値はそのページの形式とパーサ）
    - パーサが [yuin/goldmark] であれば [Markdown Render Hooks] が適用される
- 入力テキストを強制的にブロック要素として展開することができる

たとえば [Org-mode] 形式による

```text
{{ "/italic org mode/" | $.Page.RenderString (dict "display" "block" "markup" "org") }}
```

という記述に対して

```html
<p>
<em>italic org mode</em>
</p>
```

などと展開される。

上の例のように，パラメータには `display` と `markup` があり（[`dict`] 関数を使って）連想配列にして渡す。

`display` が取り得る値は以下の通り。

| display 値 | 内容                                                                                      |
| ---------- | ----------------------------------------------------------------------------------------- |
| `inline`   | 既定値                                                                                    |
| `block`    | インライン・テキストと推定される場合でも強制的に `<p>` などのブロック要素として展開する |

`markup` が取り得る値は以下の通り。

| markup 値     | 内容                                                   |
| ------------- | ------------------------------------------------------ |
| `markdown`    | 既定の markdown パーサ[^md1]                           |
| `goldmark`    | Markdown として [yuin/goldmark] パーサを使用           |
| `blackfriday` | Markdown として [russross/blackfriday] パーサを使用    |
| `org`         | [Org-mode] として [niklasfasching/go-org] パーサを使用 |
| `asciidoc`    | [AsciiDoc] として処理（要 [AsciiDoc] 環境）            |
| `rst`         | [reStructuredText] として処理（要 [RST] 環境）         |
| `pandoc`      | [Pandoc] として処理（要 [Pandoc] 環境）                |
| `html`        | 素通し                                                 |

[^md1]: 既定の markdown パーサは [yuin/goldmark] だが `config.toml` 等の設定で変更することができる。

## ぼくがかんがえたさいきょうの [Hugo] Shortcodes

以上を踏まえて自作 shortcode を整理した。
以下にいくつか紹介しておこう。

### ルビを振る

まずは軽くルビを振るところから。

`shortcodes/ruby.html`:

```text
<ruby><rb>{{ .Inner }}</rb><rp> (</rp><rt>{{ index .Params 0 }}</rt><rp>) </rp></ruby>
```

*入力：*

```text
{{</* ruby "それはさておき" */>}}閑話休題{{</* /ruby */>}}
```

*出力：*

```html
<ruby><rb>閑話休題</rb><rp> (</rp><rt>それはさておき</rt><rp>) </rp></ruby>
```

{{< div-box type="markdown" >}}
{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}
{{< /div-box >}}

### 図を figure 要素で囲む

実は [Hugo] の[組み込み shortcode](https://gohugo.io/content-management/shortcodes/#use-hugos-built-in-shortcodes "Shortcodes | Hugo") に `figure` というのがあって図を `<figure>` 要素で囲んでキャプションを付けれるようになってるんだけど，作りが気に食わないので自作したのだった。

`shortcodes/fig-img.html`:

```text
<figure style='margin:0 auto;text-align:center;'>
{{ with .Get "link" }}<a href="{{ . }}">{{ end }}<img src="{{ .Get "src" }}" srcset="{{ .Get "src" }}{{ with .Get "width" }}{{ . }}{{ else }}500{{ end }}w" sizes="(min-width:600px) 500px, 80vw"{{ if .Get "title" }} alt="{{ .Get "title" }}"{{ else if .Get "alt" }} alt="{{ .Get "alt" }}"{{ else }} alt=""{{ end }}>{{ with .Get "link" }}</a>{{ end }}{{ if .Get "title" }}
<figcaption>{{ with .Get "link" }}<a href="{{ . }}">{{ end }}{{ .Get "title" }}{{ with .Get "link" }}</a>{{ end }}</figcaption>
{{ end }}</figure>
```

*入力：*

```text
{{</* fig-img src="https://photo.baldanders.info/flickr/image/8113077_o.png" title="SHIMENAWA" link="https://photo.baldanders.info/flickr/8113077/" width="768" */>}}
```

*出力：*

```html
<figure style='margin:0 auto;text-align:center;'>
<a href="https://photo.baldanders.info/flickr/8113077/"><img src="https://photo.baldanders.info/flickr/image/8113077_o.png" srcset="https://photo.baldanders.info/flickr/image/8113077_o.png 768w" sizes="(min-width:600px) 500px, 80vw" alt="SHIMENAWA"></a>
<figcaption><a href="https://photo.baldanders.info/flickr/8113077/">SHIMENAWA</a></figcaption>
</figure>
```

{{< fig-img src="https://photo.baldanders.info/flickr/image/8113077_o.png" title="SHIMENAWA" link="https://photo.baldanders.info/flickr/8113077/" width="768" >}}

`<img>` 要素の各属性値はこのブログのレイアウトに合わせてチューニングしているので，そのままでは（多分）使えない。
あしからず。

### 引用を figure 要素で囲んで引用元を指示する

引用ブロックは `<blockquote>` 要素で囲めば「意味」は通るし，実際に markdown の引用（`>`）も `<blockquote>` 要素内に展開するようになっている。
しかしこれでは「引用元」の指示が曖昧になってしまう。

そこで `<blockquote>` 要素を更に `<figure>` 要素で囲んでキャプションを付加することで引用元を指示できるようにした。

`shortcodes/fig-quote.html`:

```text
{{ with .Get "lang" }}<figure lang="{{ . }}">{{ else }}<figure>{{ end }}
<blockquote>{{ with .Get "type" }}{{ $.Inner | $.Page.RenderString (dict "markup" .) }}{{ else }}{{ .Inner }}{{ end }}</blockquote>{{ if .Get "title"}}
<figcaption>via <q>{{ with .Get "link"}}<a href="{{.}}">{{ end }}{{ .Get "title" }}{{ with .Get "link"}}</a>{{ end }}</q></figcaption>
{{ end }}</figure>
```

*入力：*

```text
{{</* fig-quote title="The Two Myths of the Internet | WIRED" link="https://www.wired.com/story/the-two-myths-of-the-internet/" lang="en" */>}}
<q>Technologies determine nothing.
Technologies influence everything</q>.
{{</* /fig-quote */>}}
```

*出力：*

```html
<figure lang="en">
<blockquote>
<q>Technologies determine nothing.
Technologies influence everything</q>.
</blockquote>
<figcaption>via <q><a href="https://www.wired.com/story/the-two-myths-of-the-internet/">The Two Myths of the Internet | WIRED</a></q></figcaption>
</figure>
```

{{< fig-quote title="The Two Myths of the Internet | WIRED" link="https://www.wired.com/story/the-two-myths-of-the-internet/" lang="en" >}}
<q>Technologies determine nothing.
Technologies influence everything</q>.
{{< /fig-quote >}}

これで「引用」の最低要件は満たせるかな。

### 汎用 figure 環境

上記以外に汎用的に `<figure>` 環境を使うための shortcode。

`shortcodes/fig-gen.html`:

```text
<figure style='margin:0 auto;text-align:center;'>{{ with .Get "type" }}{{ $.Inner | $.Page.RenderString (dict "markup" .) }}{{ else }}{{ .Inner }}{{ end }}{{ if .Get "title"}}
<figcaption>{{ with .Get "link"}}<a href="{{.}}">{{ end }}{{ .Get "title" }}{{ with .Get "link"}}</a>{{ end }}</figcaption>
{{ end }}</figure>
```

入出力例は割愛する。

### 汎用 div 環境

これはもっと素朴に `<div>` で囲むだけの shortcode。

`shortcodes/div-gen.html`:

```text
<div{{ with .Get "class" }} class="{{ . }}"{{ end }}{{ with .Get "height" }} style="height:{{ . }};"{{ end }}{{ with .Get "lang" }} lang="{{ . }}"{{ end }}>{{ with .Get "type" }}{{ $.Inner | $.Page.RenderString (dict "markup" .) }}{{ else }}{{ .Inner }}{{ end }}</div>
```

`class` を指定できる他， `height` で高さを指定できる。
ネタバレ防止の行間稼ぎでよく使う手ですな（笑）

これも入出力例は割愛する。

## ブックマーク

- [Hugo Christmas Edition! | Hugo](https://gohugo.io/news/0.62.0-relnotes/)
- [Markdown Render Hooks - リンクや画像のテンプレートを作成する仕組み | Hugo 入門 / 解説 | nasust dev blog](https://nasust.com/hugo/tips/markdown_render_hooks/)
- [.RenderString | Markdownの文字列をHTMLに変換する | Hugo 入門 / 解説 | nasust dev blog](https://nasust.com/hugo/functions/render_string/)
- [A way to add parameters to image and link render hooks - tips & tricks - HUGO](https://discourse.gohugo.io/t/a-way-to-add-parameters-to-image-and-link-render-hooks/23496)

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[CommonMark]: https://commonmark.org/
[yuin/goldmark]: https://github.com/yuin/goldmark/ "yuin/goldmark: A markdown parser written in Go. Easy to extend, standard(CommonMark) compliant, well structured."
[russross/blackfriday]: https://github.com/russross/blackfriday "russross/blackfriday: Blackfriday: a markdown processor for Go"
[Org-mode]: https://orgmode.org/ "Org mode for Emacs – Your Life in Plain Text"
[niklasfasching/go-org]: https://github.com/niklasfasching/go-org "niklasfasching/go-org: Org mode parser with html & pretty printed org rendering"
[AsciiDoc]: http://asciidoc.org/
[reStructuredText]: https://docutils.sourceforge.io/rst.html "reStructuredText"
[RST]: https://docutils.sourceforge.io/rst.html "reStructuredText"
[Pandoc]: https://www.pandoc.org/
[Markdown Render Hooks]: https://gohugo.io/getting-started/configuration-markup/#markdown-render-hooks "Configure Markup | Hugo"
[`markdownify`]: https://gohugo.io/functions/markdownify/ "markdownify | Hugo"
[`dict`]: https://gohugo.io/functions/dict/ "dict | Hugo"
[`.RenderString`]: https://gohugo.io/functions/renderstring/ ".RenderString | Hugo"
[`Page`]: https://gohugo.io/variables/page/ "Page Variables | Hugo"
