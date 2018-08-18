+++
title = "結局 Noto Serif JP を Web フォントとして導入した"
date =  "2017-12-11T19:03:01+09:00"
description = "昨日の話で「Noto Serif の和文フォントが Web フォントで登場するまでは「さわらび明朝」で頑張る」と書いたばっかりなのに，日和りました。"
image = "/images/attention/remark.jpg"
tags        = [ "web", "site", "font", "character" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[昨日の話]({{< ref "/remark/2017/12/gods-in-chrome.md" >}} "「神」と「神」と Chrome の文字化け")で「Noto Serif の和文フォントが Web フォントで登場するまでは「[さわらび明朝]」で頑張る」と書いたばっかりなのに，日和りました。
意思が弱いな，私。

厳密には[本家サイト]の方に導入してその設定を丸々こっちにも適用した。
ので，今回はその覚え書き。

## Noto Serif JP の取得

Noto Serif の和文フォントは以下から取得できる。

- [Noto CJK – Google Noto Fonts](https://www.google.com/get/noto/help/cjk/)


このページの後半の方に “Region-specific Subset OpenType/CFF (Subset OTF)” というのがあるので，ここから `NotoSerifJP-[weight].otf` のリンク先の zip ファイルを取得する。
Web フォント用ならこのサブセットでも大きすぎるくらいだが，まぁいいだろう。

Zip ファイルを展開すると，以下の OTF (OpenType Font) ファイルがある。
括弧内の数字はウェイト数である[^fw1]。

[^fw1]: CSS の `font-weight` に設定する数値。 `font-weight` の `normal` が 400 に相当する。ちなみに `bold` は 700。 `bolder` で 400 → 700 → 900 と太くなる。 `lighter` だと 400 → 100 と細くなる。

- `NotoSerifJP-ExtraLight.otf` (200)
- `NotoSerifJP-Light.otf` (300)
- `NotoSerifJP-Regular.otf` (400)
- `NotoSerifJP-Medium.otf` (500)
- `NotoSerifJP-SemiBold.otf` (600)
- `NotoSerifJP-Bold.otf` (700)
- `NotoSerifJP-Black.otf` (900)

はっきり言って通常の文章の中で明朝体の太さを変えることはないが（強調するときはゴシック体にする），今回は念のため `Regular` と `Bold` のふたつのフォントを導入した。

### WOFF/WOFF2 フォントの生成と CSS の作成

取得した OTF ファイルから WOFF (Web Open Font Format) および WOFF2 (Web Open Font Format 2.0) 形式のファイルを生成する。
使用しているマシンが Windows か macOS なら以下のフリーソフトがお勧めである。 

- [WOFFコンバータ](https://opentype.jp/woffconv.htm)

Web サービスだと，以下のページで WOFF の生成ができるようだ

- [Online font converter](http://www.fontconverter.org/)

WOFF2 は比較的新しいフォーマットだが，データの圧縮率が WOFF よりもいいため，こちらのほうがお勧めである。
最近のブラウザならほぼ対応している（IE と Opera mini は無視していいと思う）。

- [WOFF 2.0 - Web Open Font Format](https://caniuse.com/#feat=woff2)

作成したフォント・ファイルを読み込むための CSS ファイルを書く。
こんな感じに書いてみた。

```css
/*
 * Original: Noto Serif CJK (Subset OTF; Japanese) https://www.google.com/get/noto/help/cjk/
 * Converted by WOFF Converter https://opentype.jp/woffconv.htm
 * License: http://www.baldanders.info/fonts/NotoSerifJP/LICENSE_OFL.txt
 * README: http://www.baldanders.info/fonts/NotoSerifJP/README
 */
@font-face {
  font-family: 'Noto Serif JP';
  font-style: normal;
  src: url('http://www.baldanders.info/fonts/NotoSerifJP/NotoSerifJP-Regular.woff2') format('woff2'),
       url('http://www.baldanders.info/fonts/NotoSerifJP/NotoSerifJP-Regular.woff') format('woff');
  font-weight: 400;
}
@font-face {
  font-family: 'Noto Serif JP';
  font-style: normal;
  src: url('http://www.baldanders.info/fonts/NotoSerifJP/NotoSerifJP-Bold.woff2') format('woff2'),
       url('http://www.baldanders.info/fonts/NotoSerifJP/NotoSerifJP-Bold.woff') format('woff');
  font-weight: 700;
}
```

`src` 指定で WOFF2 ファイルを先に書くのがポイントである。
これを読み込ませて

```css
body {
  font-family: 'Noto Serif JP', serif;
}
```

てな感じにフォントを指定すればよい。

## CORS の指定

今回は[本家サイト]に Noto Serif JP を導入し，それをこのブログサイトでも使おうと目論んでいるのだが，単純に

```html
<link rel='stylesheet' href='http://www.baldanders.info/fonts/NotoSerifJP/NotoSerifJP.css' type='text/css'>
```

と Web フォントを記述した CSS ファイルをリンクしただけでは使えない。
JavaScript の `XMLHttpRequest` と同じく， Web フォントも（主にセキュリティ上の理由から）別ドメインからの読み込みが既定で禁止されているためだ。

この制約を解除するには Web フォントを提供する側で HTTP レスポンス・ヘッダに `Access-Control-Allow-Origin` 情報を付加して相手先を指定する必要がある。

さくらのレンタルサーバなら `.htaccess` ファイルに

```text
# Cross-Origin Resource Sharing 
Header append Access-Control-Allow-Origin: http://text.baldanders.info
```

と記述すればヘッダに追加されるようだ。
相手先の指定を `*` にすると任意のサイトから（Web フォントだけでなく JavaScript も）アクセスし放題になるので，ご注意を。

## ローカルのフォントも有効にする

（ブラウザ側の設定などで） Web フォントが使えない場合でも，ローカルマシンに Noto 和文フォントがあれば優先的に参照できるようにした。
まずはこんな感じで CSS を書く。

```css
@font-face {
  font-family: 'NotoSansJP-local';
  font-style: normal;
  src: local("Noto Sans JP"), local('NotoSansJP');
}
@font-face {
  font-family: 'NotoSansJP-local';
  font-style: normal;
  src: local("Noto Sans JP Regular"), local('NotoSansJP-Regular');
  font-weight: 400;
}
@font-face {
  font-family: 'NotoSansJP-local';
  font-style: normal;
  src: local("Noto Sans JP Bold"), local('NotoSansJP-Bold');
  font-weight: 700;
}
@font-face {
  font-family: 'NotoSerifJP-local';
  font-style: normal;
  src: local("Noto Serif JP"), local('NotoSerifJP');
}
@font-face {
  font-family: 'NotoSerifJP-local';
  font-style: normal;
  src: local("Noto Serif JP Regular"), local('NotoSerifJP-Regular');
  font-weight: 400;
}
@font-face {
  font-family: 'NotoSerifJP-local';
  font-style: normal;
  src: local("Noto Serif JP Bold"), local('NotoSerifJP-Bold');
  font-weight: 700;
}
```

これで

```css
body {
  font-family: 'NotoSerifJP-local', 'Noto Serif JP', serif;
}
h1, h2, h3, h4, h5, h6, em, strong {
  font-family: 'NotoSansJP-local', 'Noto Sans JP', sans-serif;
}
```

などとしておけば，ローカルの Noto 和文フォントを優先的に見てくれる。
なお，欧文 Noto フォントの CSS は既にローカルのフォントを優先的に見るよう設定されている。

早く Noto Serif JP が CDN 配信されるようにならないかな。

## ブックマーク

- [WOFF (Web Open Font Format) - ウェブデベロッパーガイド | MDN](https://developer.mozilla.org/ja/docs/Web/Guide/WOFF)
- [HTTP アクセス制御 (CORS) - HTTP | MDN](https://developer.mozilla.org/ja/docs/Web/HTTP/HTTP_access_control)
- [さくらのサーバに置いたWebフォントをはてなブログから使う - にせねこメモ](http://nixeneko.hatenablog.com/entry/2016/10/11/231435)
- [「Noto Serif Japanese」を使ってみたので、自分用メモ。 - Qiita](https://qiita.com/umeume66/items/01291353fc43c17da992)

- [Web フォントに関する覚え書き（明朝体編）]({{< ref "/remark/2016/10/japanese-serif-fonts-by-google-cdn.md" >}})

[本家サイト]: http://www.baldanders.info/ "Baldanders.info"
[さわらび明朝]: http://sawarabi-fonts.osdn.jp/ "さわらびフォント"
