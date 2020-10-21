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

[^url1]: URL のクエリー部（`?` 以降の文字列）に `|` や `:` の文字を使うと [W3C の Validator](https://validator.w3.org/ "The W3C Markup Validation Service") に怒られるのでご注意を。大抵はブラウザが上手く解釈してくれるので問題ないのだが。気になるなら `%7c` や `%3a` といった[パーセント・エンコーディング](https://tools.ietf.org/html/rfc3986#section-2.1 "RFC 3986 - Uniform Resource Identifier: Generic Syntax")に置き換えればよい。

```html
<link rel="stylesheet" href="//fonts.googleapis.com/css?family=Noto+Sans+JP:400,700|Noto+Sans:400,700|Noto+Serif|Noto+Serif+JP&display=swap">
```

という感じにウェイトを指定できる[^fw1]。

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

## 【付録】 Google Fonts スタイルの中身

たとえば NOTO Sans のみを指定する場合は，スタイルシートの指定は以下のようになるが

```html
<link rel="stylesheet" href="//fonts.googleapis.com/css?family=Noto+Sans&display=swap">
```

中身を見るとこんな感じになっている（一部のみ）。

```css
/* latin-ext */
@font-face {
  font-family: 'Noto Sans';
  font-style: normal;
  font-weight: 400;
  font-display: swap;
  src: local('Noto Sans'), local('NotoSans'), url(https://fonts.gstatic.com/s/notosans/v9/o-0IIpQlx3QUlC5A4PNr6zRAW_0.woff2) format('woff2');
  unicode-range: U+0100-024F, U+0259, U+1E00-1EFF, U+2020, U+20A0-20AB, U+20AD-20CF, U+2113, U+2C60-2C7F, U+A720-A7FF;
}
/* latin */
@font-face {
  font-family: 'Noto Sans';
  font-style: normal;
  font-weight: 400;
  font-display: swap;
  src: local('Noto Sans'), local('NotoSans'), url(https://fonts.gstatic.com/s/notosans/v9/o-0IIpQlx3QUlC5A4PNr5TRA.woff2) format('woff2');
  unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
}
```

このように単一のフォント・データではなく，いくつかのサブセットに分けてダウンロードさせているようだ。
NOTO Sans のラテン文字は上の2つのようだが（他の文字のサブセットもある），日本語の NOTO Sans JP では，なんと！ 120に分割されていた。

更に言うと， `href` 属性で指定される URL のクエリ部に `display=swap` がくっ付いているのに気づかれただろうが，これを指定すると `@font-face` ルールのプロパティとして `font-display: swap;` がセットされるようだ。 

不勉強で `font-display` については知らなかったのだが，このプロパティに `swap` がセットされていると Web フォントがロードされるまでの間，代替フォント（`sans-serif` など）で表示されるらしい。

`font-display` の仕様は割と最近にできたものらしく，対応していないブラウザも一部あるようだ（[ここで対応ブラウザを確認](https://caniuse.com/#search=font-display "Can I use... Support tables for HTML5, CSS3, etc")できる）。

ちなみに Android 版 Firefox も `font-display` に対応しているのだが，フォントが入れ替わるたびに画面がちらつくのが鬱陶しい。
この辺は要改良ってところなんだろう。

日本語フォントは巨大だが，このような仕組みを使って，できるだけストレスのないよう工夫されているわけだ。

ありがとう！

## ブックマーク

- [日本語Webフォントを使用する際のアプローチについて | フロントエンドBlog | ミツエーリンクス](https://www.mitsue.co.jp/knowledge/blog/frontend/201912/19_0000.html) : Web フォントで日本語を設定する際に参考になる
- [デザイナーとフロントエンドエンジニアに知ってほしいWebのフォント周りのお話](https://zenn.dev/tak_dcxi/articles/588fbc205251043dc357)
- [Yaku Han JP](https://yakuhanjp.qranoko.jp/) : NOTO フォント派生。訳物だけ半角サイズになっている
- [Google Fontsは使用する文字を絞り込んでダウンロードできる](https://zenn.dev/neriko/articles/3b55c547a07c8d22c9f1)

- [Web フォントに関する覚え書き]({{< ref "/remark/2015/web-font-family.md" >}})
- [Web フォントに関する覚え書き（明朝体編）]({{< ref "/remark/2016/10/japanese-serif-fonts-by-google-cdn.md" >}})
- [結局 Noto Serif JP を Web フォントとして導入した]({{< ref "/remark/2017/12/installing-noto-serif-jp-in-www_baldanders_info.md" >}}) : 自前で Web フォントをインストールしたときの話
- [Font Awesome 5.0.11 で Creative Commons アイコンに完全対応した]({{< ref "/release/2018/05/creative-commons-icons-by-font-awesome.md" >}})

<!-- eof -->
