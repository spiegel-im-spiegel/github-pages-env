+++
title = "Hugo でスライド・サイトを立てる実験"
date =  "2019-12-21T22:38:36+09:00"
description = "ちょろんとググってみたら reveal-hugo なる Hugo テーマがあって，内部では reveal.js を使ってスライドを制御しているらしい。"
image = "/images/attention/kitten.jpg"
tags = [ "hugo", "presentation", "site", "revealjs" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[来月のイベント](https://shimane-go.connpass.com/event/159977/ "Shimane.go#03 - connpass")用に資料をスライド形式で作ろうと思うのだが， MS Office は大昔に捨ててしまったし LibreOffice の Impress で書くというのも今更だしなぁ。

できれば markdown で書いて VCS でドキュメント管理したい。

で，ちょろんとググってみたら [reveal-hugo] なる [Hugo] テーマがあって，内部では [reveal.js] を使ってスライドを制御しているらしい。
これだよ，これ。

[Hugo] は静的サイト・ジェネレータなのでデプロイ先の自由度が高い。
さくらのレンタルサーバみたいな老舗でも GitHub Page や最近流行り（？）の [Netlify] でも大丈夫。
しかも [Hugo] にはサーバ・モードがあるのでローカル環境で試すこともできる。

[Reveal.js] テーマを自作するのでなければ通常版の [Hugo] で問題ない。
また [reveal-hugo] には [reveal.js] 一式がまるっと含まれているので別途のインストールは不要。
Node.js をインストールする必要もない。
簡単！

では，さっそく試してみよう。

## [Hugo] 環境の構築

まずは適当なディレクトリで [Hugo] 環境を作成する。

```text
$ hugo new site hugo-slide
Congratulations! Your new Hugo site is created in /home/username/hugo-env/hugo-slide.

Just a few more steps and you're ready to go:

1. Download a theme into the same-named folder.
   Choose a theme from https://themes.gohugo.io/ or
   create your own with the "hugo new theme <THEMENAME>" command.
2. Perhaps you want to add some content. You can add single files
   with "hugo new <SECTIONNAME>/<FILENAME>.<FORMAT>".
3. Start the built-in live server via "hugo server".

Visit https://gohugo.io/ for quickstart guide and full documentation.
```

これで `hugo-slide` ディレクトリが作成される。
このディレクトリに入り git リポジトリとして初期化する。

```text
$ cd hugo-slide
$ git init
```

この時点でコミットしておけばいつでも元に戻せるので安心である。

次に [reveal-hugo] を git のサブモジュールとして導入する。

```text
$ git submodule add git@github.com:dzello/reveal-hugo.git themes/reveal-hugo
Cloning into '/home/username/hugo-env/hugo-slide/themes/reveal-hugo'...
remote: Enumerating objects: 1713, done.
remote: Total 1713 (delta 0), reused 0 (delta 0), pack-reused 1713
Receiving objects: 100% (1713/1713), 6.61 MiB | 1.83 MiB/s, done.
Resolving deltas: 100% (848/848), done.
```

マジ簡単。
[Reveal-hugo] をアップデートする場合は

```text
$ git submodule update --init
```

でよい。
以降 [reveal-hugo] の中身は弄らない方針で作業を進める。

### config.toml の設定

[Hugo] 環境作成直後の `config.toml` の中身はこんな感じになっている。

```toml
baseURL = "http://example.org/"
languageCode = "en-us"
title = "My New Hugo Site"
```

これを以下のように書き換える。


```toml
baseURL = "http://example.org/"
languageCode = "ja_JP"
title = "スライド用デモ・サイト"
theme = "reveal-hugo"

[outputFormats.Reveal]
  baseName = "index"
  mediaType = "text/html"
  isHTML = true

[markup]
  defaultMarkdownHandler = "blackfriday"
  [markup.highlight]
    codeFences = false
```

[Hugo] で作成したサイトをデプロイするなら `baseURL` にデプロイ先の URL を設定する。

`[outputFormats.Reveal]` 項目は [reveal.js] 用の出力設定で，取り敢えずおまじないのようなものだと思っておけばよい。

最後の4行は [Hugo] 0.60 以降では必須の設定である。
これも取り敢えずおまじないとして唱えておく（笑）

## スライド用のドキュメントを作成・表示する

`hugo-slide/content` ディレクトリに `_index.md` ファイルを作成する。
お試しなので，中身はこんな感じ。

```markdown
+++
title = "Hello world!"
outputs = ["Reveal"]
+++

# Hello world!

This is my first slide.

---

# Hello Mars!

This is my second slide.
```

これで2ページ分のスライドができた。

早速 [Hugo] をサーバモードで起動して表示してみよう。

```text
$ hugo server
```

この状態でブラウザで `http://localhost:1313/` にアクセスすれば

{{< fig-img src="./first-page.png" link="./first-page.png" width="850" >}}

おおっ，できたできた。
ちなみに `[F]` キー押下で全画面表示になる（元に戻すには `[ESC]` キー押下）。
`[Page Down]` キーを押すか右下の `>` をクリックすれば次ページに遷移する。

{{< fig-img src="./second-page.png" link="./second-page.png" width="850" >}}

よーし，うむうむ，よーし。

## セクションでスライドを分ける

[来月のイベント](https://shimane-go.connpass.com/event/159977/ "Shimane.go#03 - connpass")は技術系なのでコードも書けるようにしておきたい。

では `hugo-slide/content` ディレクトリに `hello` ディレクトリを切ってその中に `_index.md` ファイルを作成してみよう。
中身はこんな感じ。

{{< highlight markdown >}}
+++
title = "みんな大好き Hello World"
outputs = ["Reveal"]
+++

# みんな大好き Hello World

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```
{{< /highlight >}}

これで `http://localhost:1313/hello/` にアクセスすれば

{{< fig-img src="./hello-default.png" link="./hello-default.png" width="850" >}}

と表示される。
このようにディレクトリを切ってセクション毎に分けることで複数のスライドを置くことができる。

ただし [Hugo] はセクションの階層化はできない。
更にいうと [reveal-hugo] ではひとつのセクションに複数の原稿ファイルをおいてもひとつのスライドにまとめられてしまうようだ。
ご注意を。

さて，次は見た目をちょっと弄ってみようか。

## スライドテーマの指定

[Reveal-hugo] ではサイト単位，あるいは front matter 単位で [reveal.js] テーマを指定できる。
`config.toml` でサイト単位で指定する場合は

```toml
[params.reveal_hugo]
  theme = "sky"
```

という感じに， front matter で指定するなら

```toml
[reveal_hugo]
  theme = "sky"
```

てな感じに指定できる。

[Reveal.js] が用意しているテーマには以下のものがある。

| Name        | Description                                                                   |
| ----------- | ----------------------------------------------------------------------------- |
| `black`     | Black background, white text, blue links (default theme)                      |
| `white`     | White background, black text, blue links                                      |
| `league`    | Gray background, white text, blue links (default theme for reveal.js < 3.0.0) |
| `beige`     | Beige background, dark text, brown links                                      |
| `sky`       | Blue background, thin dark text, blue links                                   |
| `night`     | Black background, thick white text, orange links                              |
| `serif`     | Cappuccino background, gray text, brown links                                 |
| `simple`    | White background, black text, blue links                                      |
| `solarized` | Cream-colored background, dark green text, blue links                         |

ちなみに `sky` だとこんな感じ。

{{< fig-img src="./hello-sky.png" link="./hello-sky.png" width="850" >}}

更に [reveal-hugo] ではカスタム・テーマとして `robot-lung` と `sunblind` が用意されている。
カスタム・テーマは

```toml
[params.reveal_hugo]
  custom_theme = "reveal-hugo/themes/robot-lung.css"
```

という感じに指定できる。
もちろん自作のテーマを組み込むことも可能である。

## Syntax Highlight の指定

コード部分の Syntax Highlight には [highlight.js] が使われている。
[Highlight.js] は [reveal.js] のプラグインとして既定で組み込まれている。

[Highlight.js] のスタイルは

```toml
[params.reveal_hugo]
  highlight_theme = "github"
```

てな感じで指定できる。
[highlight.js] で指定可能なスタイルは[デモ・ページ](https://highlightjs.org/static/demo/ "highlight.js demo")で確認できる。

### [Hugo] の Syntax Highlight を使う

ところで [Hugo] は自前で Syntax Highlight の機能を持っている。
これを使わない手はない。

まず [highlight.js] プラグインを無効にするのだが，個別にプラグインを無効化する方法がないため，`config.toml` で

```toml
[params.reveal_hugo]
  load_default_plugins = false
  plugins = [
    "reveal-js/plugin/zoom-js/zoom.js",
    "reveal-js/plugin/notes/notes.js",
  ]
```

という感じに，いったん既定のプラグインを全て無効にしてから必要なものを個別に組み込む。

[Hugo] 側の Syntax Highlight の設定は

```toml
[markup.highlight]
  codeFences = true
  hl_Lines = ""
  lineNoStart = 1
  lineNos = false
  lineNumbersInTable = true
  noClasses = true
  style = "tango"
  tabWidth = 4
```

などとすればよいだろう。
`codeFences` が `true` になっている点に注意。

`style` のサンプルは以下を参考にどうぞ。

- [Short snippets](https://xyproto.github.io/splash/docs/all.html)
- [Long snippets](https://xyproto.github.io/splash/docs/longer/all.html)

たとえば `tango` スタイルを使えば

{{< fig-img src="./hello-sky-tango.png" link="./hello-sky-tango.png" width="850" >}}

てな感じになる。

## Web フォントを使いたい！

Web で公開することを考えるのなら Web フォントを使いたいところである。
それにタイトルがちょいとデカすぎるしコードは小さすぎるよね。

[Reveal-hugo] では `layouts/partials/reveal-hugo/head.html` を使ってページの `<head>` 要素に任意の記述を割り込ませることができる。
これを使ってフォントの調整をしよう。

たとえばこんな記述を割り込ませてみる。

```html
<link rel='stylesheet' href='//fonts.googleapis.com/css?family=Noto+Sans%7cNoto+Sans+JP%3a400,700%7cInconsolata%3a400,700' type='text/css'>
<link rel='stylesheet' href='/css/font-family.css' type='text/css'>
```

ちなみに `font-family.css` はこんな内容にしてみた[^css1]。

[^css1]: `font-family.css` ファイルは `static/css/` ディレクトリに置くこと。 

```css
.reveal {
  font-family: 'Noto Sans', 'Noto Sans JP', sans-serif;
}
.reveal h1,
.reveal h2,
.reveal h3,
.reveal h4,
.reveal h5,
.reveal h6 {
  font-family: 'Noto Sans', 'Noto Sans JP', sans-serif;
}
.reveal code {
  font-family: 'Inconsolata', 'Noto Sans JP', monospace;
}
.reveal h1 {
  font-size: 1.6em;
}
.reveal h2 {
  font-size: 1.3em;
}
.reveal h3 {
  font-size: 1em;
}
.reveal h4 {
  font-size: 1em;
}
.reveal pre {
  font-size: 0.8em;
}
```

これで実際の表示は

{{< fig-img src="./hello-webfonts.png" link="./hello-webfonts.png" width="850" >}}

てな感じになった。
こんなもんかな。

[Reveal-hugo] では他にも fragment など [reveal.js] のギミックが利用可能だが，今回は割愛する。

## ブックマーク

- [hakimel/reveal.js: The HTML Presentation Framework](https://github.com/hakimel/reveal.js/)
- [dzello/reveal-hugo: 📽️ Create rich HTML-based presentations with Hugo and Reveal.js](https://github.com/dzello/reveal-hugo)
    - [Reveal Hugo | Hugo Themes](https://themes.gohugo.io/reveal-hugo/)

- [さくらのレンタルサーバ上で Hugo によるサイト管理を行う]({{< ref "/remark/2019/01/sakura-and-hugo.md" >}})
- [Hugo v0.60 から既定の Markdown パーサが Goldmark になったようだ]({{< ref "/release/2019/11/hugo-v0_60-with-goldmark-parser.md" >}})

[reveal-hugo]: https://reveal-hugo.dzello.com/
[Reveal-hugo]: https://reveal-hugo.dzello.com/
[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[Netlify]: https://www.netlify.com/ "Netlify: All-in-one platform for automating modern web projects."
[reveal.js]: https://revealjs.com/ "reveal.js – The HTML Presentation Framework"
[Reveal.js]: https://revealjs.com/ "reveal.js – The HTML Presentation Framework"
[highlight.js]: https://highlightjs.org/
[Highlight.js]: https://highlightjs.org/
<!-- eof -->
