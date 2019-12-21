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

と記述すればよい。
NOTO Sans の太字もご所望なら

```html
<link rel="stylesheet" href="//fonts.googleapis.com/css?family=Noto+Sans+JP:400,700|Noto+Sans:400,700|Noto+Serif|Noto+Serif+JP&display=swap">
```

という感じにウェイトを指定すればいいだろう。

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
- [Font Awesome 5.0.11 で Creative Commons アイコンに完全対応した]({{< ref "/release/2018/05/creative-commons-icons-by-font-awesome.md" >}})

<!-- eof -->
