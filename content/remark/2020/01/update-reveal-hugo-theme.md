+++
title = "Hugo 0.60+ 向けに reveal-hugo がアップデートされたようだ"
date =  "2020-01-19T21:48:34+09:00"
description = "ただし markdown パーサが変わることで shortcodes の挙動が変わるので注意が必要である。"
image = "/images/attention/kitten.jpg"
tags = [ "hugo", "presentation", "revealjs" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Hugo] 用のスライド・テーマ [reveal-hugo] を [Hugo] 0.60+ でも使えるように修正した pull request が{{< ruby "マージ" >}}受理{{< /ruby >}}されたようだ。

- [Hugo 0.60 compatibility by chatziko · Pull Request #69 · dzello/reveal-hugo · GitHub](https://github.com/dzello/reveal-hugo/pull/69)

これまでは markdown パーサとして [russross/blackfriday] にのみ最適化されていたが，今回のマージで既定の [yuin/goldmark] パーサが使えるようになった。
したがって，[Hugo] 環境の設定ファイル（`config.toml`）でも

```toml
[markup]
  defaultMarkdownHandler = "blackfriday"
```

の指定が不要になる。

なお  [yuin/goldmark] パーサのオプションとして

```toml
[markup.goldmark.renderer]
  unsafe = true
```

の指定はしておいたほうがいいだろう（既定の `false` のままだと本文中の HTML 記述が排除されてしまうため）。

ただし markdown パーサが変わることで shortcodes の挙動が変わるので注意が必要である。
挙動が変わりすぎて問題になるようなら [russross/blackfriday] パーサに戻すのも手である。

Shortcodes の挙動については，たとえば “`fragment`” なら，今までの [russross/blackfriday] パーサであれば

```text
{{%/* fragment */%}}foo{{%/* /fragment */%}}

{{%/* fragment */%}}bar{{%/* /fragment */%}}
```

と記述すれば

```html
<p><span class='fragment'>foo</span></p>

<p><span class='fragment'>bar</span></p>
```

と展開されるが， [yuin/goldmark] パーサでは

```html
<span class='fragment'>foo</span>

<span class='fragment'>bar</span>
```

てな感じになり `<p>` タグが付加されない。
ちなみに

```text
<p>{{%/* fragment */%}}foo{{%/* /fragment */%}}</p>

<p>{{%/* fragment */%}}bar{{%/* /fragment */%}}</p>
```

のように強制的に `<p>` タグを付けても，どこかに吸い込まれてしまい，同じ結果になる。

更に行間を空けずに

```text
{{%/* fragment */%}}foo{{%/* /fragment */%}}
{{%/* fragment */%}}bar{{%/* /fragment */%}}
```

とすると，合わせてひとつの段落とみなされるため

```html
<p><span class='fragment'>foo</span>
<span class='fragment'>bar</span></p>
```

のように全体が `<p>` タグで囲まれる。
正しいけど，それがしたいんじゃないんだって！ どうにも shortcodes は挙動が謎だよなぁ。

“`fragment`” の中身を見ると

```text
{{/* Render .Inner before processing the shortcode. */}}
{{ $_hugo_config := `{ "version": 1 }` }}
<span class='fragment {{ .Get "class" }}'
  {{ with .Get "index" }}data-fragment-index='{{ . }}'{{ end }}>
  {{ .Inner }}
</span>
```

とかなっていて，どうやら [0.55 以降にすら対応していない]({{< ref "/release/2019/04/broken-backward-compatibility-by-hugo-0_55.md" >}} "Hugo 0.55 リリースでまた後方互換性が壊れた")ようだ[^bc1]。

[^bc1]: この ``{{ $_hugo_config := `{ "version": 1 }` }}`` は [Hugo] 0.55 より前のバージョンの shortcodes の互換性を確保するための設定である。

とりあえず要らないものは取って

```text
<span class='fragment{{ with .Get "class" }} {{ . }}{{ end }}'{{ with .Get "index" }} data-fragment-index='{{ . }}'{{ end }}>{{ .Inner }}</span>
```

てな感じに1行にしてしまえば [Hugo] 0.60 以降にも対応する気がするが，どういう副作用があるか分からないので，今回は放っておく。
でもこのままでは困るので，私の方で “`paragraph`” shortcode を作ってみた。
中身はこんな感じ。

```text
<p{{ with .Get "class" }} class='{{ . }}'{{ end }}{{ with .Get "lang" }} lang='{{ . }}'{{ end }}>{{ .Inner }}</p>
```

これを使って

```text
{{</* paragraph */>}}{{%/* fragment */%}}foo{{%/* /fragment */%}}{{</* /paragraph */>}}

{{</* paragraph */>}}{{%/* fragment */%}}bar{{%/* /fragment */%}}{{</* /paragraph */>}}
```

とすれば意図通り

```html
<p><span class='fragment'>foo</span></p>

<p><span class='fragment'>bar</span></p>
```

と展開される。

今回はこれで凌ぐか。
準備完了っと。

## ブックマーク

- [Hugo でスライド・サイトを立てる実験]({{< ref "/remark/2019/12/slide-site-by-hugo.md" >}})
- [ぼくがかんがえたさいきょうの Hugo Shortcodes]({{< ref "/remark/2020/01/my-hugo-shortcodes.md" >}})

[reveal-hugo]: https://reveal-hugo.dzello.com/
[Reveal-hugo]: https://reveal-hugo.dzello.com/
[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[yuin/goldmark]: https://github.com/yuin/goldmark/ "yuin/goldmark: A markdown parser written in Go. Easy to extend, standard(CommonMark) compliant, well structured."
[russross/blackfriday]: https://github.com/russross/blackfriday "russross/blackfriday: Blackfriday: a markdown processor for Go"
