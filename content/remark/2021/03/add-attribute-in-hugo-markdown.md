+++
title = "Hugo Markdown でブロック要素にクラス属性を付与する"
date =  "2021-03-22T17:14:37+09:00"
description = "お手軽装飾"
image = "/images/attention/kitten.jpg"
tags = [ "hugo", "markdown" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Hugo] 0.81 からの機能らしいのだが， markdown テキスト中のブロック要素に対して class 属性を付けられるようだ（[goldmark] パーサを使う場合）。
これを有効にするには `config.toml` ファイルで

```toml
[markup.goldmark.parser.attribute]
  block = true
```

と指定する（既定値は `false`）。
YAML (`config.yaml`) で書くなら

```yaml
markup:
  goldmark:
    parser:
      attribute:
        block: true
```

てな感じかな。

たとえば

```markdown
- One
- Two
- Three
{ .cloud .center }
```

と箇条書きブロックの最後に `{ .cloud .center }` とクラス名を付けると

```html
<ul class="cloud center">
<li>One</li>
<li>Two</li>
<li>Three</li>
</ul>
```

という感じに展開してくれる。
ここで CSS 側を

```css
ul.cloud {
  list-style: none;
  padding: 0
}
ul.cloud > li {
  display: inline-block;
  margin: 0 0.5rem;
}
.center {
  text-align: center;
}
```

などと定義すれば

{{< div-box type="markdown" >}}
- One
- Two
- Three
{ .cloud .center }
{{< /div-box >}}

とできる。

これは HTML のブロック要素に相当するものであれば何でも適用できるみたいで，通常の段落でも

```markdown
*「人はなぜ大丈夫じゃないときに限って『大丈夫？』と訊くのだろう」*
<br>（某ラノベ作品より）
{.center}
```

などと書けば

{{< div-box type="markdown" >}}
*「人はなぜ大丈夫じゃないときに限って『大丈夫？』と訊くのだろう」*
<br>（某ラノベ作品より）
{.center}
{{< /div-box >}}

とできる。

頻繁に使う表現なら [shortcode] を組んだほうがいいだろうが，ちょっとした装飾ならこちらのほうが手軽でいいかもしれない。

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[CommonMark]: https://commonmark.org/
[goldmark]: https://github.com/yuin/goldmark/ "yuin/goldmark: A markdown parser written in Go. Easy to extend, standard(CommonMark) compliant, well structured."
[shortcode]: https://gohugo.io/extras/shortcodes/ "Shortcodes | Hugo"
<!-- eof -->
