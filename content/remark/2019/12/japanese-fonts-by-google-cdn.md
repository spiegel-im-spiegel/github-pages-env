+++
title = "Google Fonts が日本語に対応してた"
date =  "2019-12-22T08:03:36+09:00"
description = "これで日本語 Web フォントを指定するのに Early Access 版を使ったり，自前で巨大フォントファイルを用意しなくてもよくなったんだね。"
image = "/images/attention/kitten.jpg"
tags = ["google", "web", "font", "character", "site"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[前回の記事]({{< relref "./slide-site-by-hugo.md" >}} "Hugo でスライド・サイトを立てる実験")を書いてて気付いたのだが， Google Fonts が正式に日本語に対応してるぢゃん！ おじさん気がつかなかったよ `orz`

- [Google Fonts](https://www.google.com/fonts/)

これで日本語 Web フォントを指定するのに [Early Access](https://fonts.google.com/earlyaccess "Early Access - Google Fonts") 版を使ったり，自前で巨大フォントファイルを用意しなくてもよくなったんだね。

たとえば Web フォントを NOTO フォントで統一するなら `<head>` 要素内に

```html
<link rel="stylesheet" href="//fonts.googleapis.com/css?family=Noto+Sans+JP|Noto+Sans|Noto+Serif|Noto+Serif+JP&display=swap">
```

と記述すればよい[^url1]。
NOTO Sans の太字もご所望なら

[^url1]: URL のパラメータ部（`?` 以降の文字列）に `|` や `:` の文字を使うと [W3C の Validator](https://validator.w3.org/ "The W3C Markup Validation Service") に怒られるのでご注意を。大抵はブラウザが上手く解釈してくれるので問題ないのだが。気になるなら `%7c` や `%3a` にエスケープすればよい。

```html
<link rel="stylesheet" href="//fonts.googleapis.com/css?family=Noto+Sans+JP:400,700|Noto+Sans:400,700|Noto+Serif|Noto+Serif+JP&display=swap">
```

という感じにウェイトを指定すればいいだろう[^fw1]。

[^fw1]: 400 とか 700 とかの値は CSS の `font-weight` に設定する数値。 `font-weight` の `normal` が 400 に相当する。ちなみに `bold` は 700。 `bolder` で 400 → 700 → 900 と太くなる。 `lighter` だと 400 → 100 と細くなる。普通の文章であれば 400 のみで十分だし，太字が必要な場合でも追加で 700 を用意しておけば問題ないだろう。

Web ページ側のスタイルは，たとえば

```css
body {
  font-family: 'Noto Serif', 'Noto Serif JP', serif;
  font-weight: 400; /* normal */
}
h1, h2, h3, h4, h5, h6 {
  font-family: 'Noto Sans', 'Noto Sans JP', sans-serif;
  font-weight: 700; /* bold */
}
```

こんな感じで記述していけばいいだろう。

ちなみにブラウザの設定によっては Google Fonts を追跡コードとみなしてブロックする場合があるので既定値（`serif` や `sans-serif`）の指定を忘れないこと。

よしよし。
これでよーやくフォント周りの懸念がなくなった。

## ブックマーク

- [Web フォントに関する覚え書き]({{< ref "/remark/2015/web-font-family.md" >}})
- [Web フォントに関する覚え書き（明朝体編）]({{< ref "/remark/2016/10/japanese-serif-fonts-by-google-cdn.md" >}})
- [結局 Noto Serif JP を Web フォントとして導入した]({{< ref "/remark/2017/12/installing-noto-serif-jp-in-www_baldanders_info.md" >}}) : 自前で Web フォントをインストールしたときの話
- [Font Awesome 5.0.11 で Creative Commons アイコンに完全対応した]({{< ref "/release/2018/05/creative-commons-icons-by-font-awesome.md" >}})

<!-- eof -->
