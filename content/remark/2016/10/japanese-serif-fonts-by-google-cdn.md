+++
date = "2016-10-16T23:41:57+09:00"
title = "Web フォントに関する覚え書き（明朝体編）"
description = "最近 Google の Early Access のページを見たら随分と日本語の書体が増えている気がする。"
tags = ["web", "font", "google", "character"]
draft = false

[author]
  name = "Spiegel"
  instagram = "spiegel_2007"
  twitter = "spiegel_2007"
  tumblr = "spiegel-im-spiegel"
  facebook = "spiegel.im.spiegel"
  license = "by-sa"
  github = "spiegel-im-spiegel"
  flickr = "spiegel"
  url = "http://www.baldanders.info/spiegel/profile/"
  avatar = "/images/avatar.jpg"
  linkedin = "spiegelimspiegel"
  flattr = "spiegel"
+++

以前，[ここで使っている Web Fonts の話]({{< relref "remark/2015/web-font-family.md" >}} "Web フォントに関する覚え書き")を書いた。
実は今まで日本語の Serif （明朝体）フォントがどうにも気に食わなかったのだが，最近 Google の Early Access のページを見たら随分と日本語の書体が増えている気がする。

- [Early Access - Google Fonts](https://fonts.google.com/earlyaccess)

この中から “Japanese” に分類されているものを挙げてみる。（10月16日時点）

- Nico Moji （[ニコモジ](http://nicofont.pupu.jp/nicomoji-plus.html "丸文字「ニコモジ・プラス」（漢字付きフォント）ダウンロード｜丸文字フォント（ニコ文字）配布所")）
- Mplus 1p （[M+ FONTS](http://mplus-fonts.osdn.jp/index.html)）
- Hannari （[はんなり明朝](http://typingart.net/?p=44 "日本語フォント「はんなり明朝」 - フォント無料ダウンロード｜Typing Art")）
- Rounded Mplus 1c （[Rounded M+](http://jikasei.me/font/rounded-mplus/ "自家製 Rounded M+ ラウンデッド エムプラス | 自家製フォント工房")）
- Nikukyu （[ニクキュウ](http://fontopo.com/?p=85 "ニクキュウ | fontopo")）
- Noto Sans Japanese
- Sawarabi Gothic （[さわらびゴシック](http://sawarabi-fonts.osdn.jp/ "さわらびフォント")）
- Kokoro （[こころ明朝体]）
- Sawarabi Mincho （[さわらび明朝](http://sawarabi-fonts.osdn.jp/ "さわらびフォント")）
- Noto Sans JP

このうち Hannari, Kokoro, Sawarabi Mincho が Serif 相当の書体になる。
それぞれの見本をみて，私は Kokoro を選択した。

[こころ明朝体]はひらがな・カタカナをデザインしたものだが，それ以外の文字は [IPA フォント](http://ipafont.ipa.go.jp/ "IPAexフォント/IPAフォント | IPAフォントのダウンロードサイトです")で補完しているため問題ないと思われる。
あと Noto Sans JP[^jp] と組み合わせた場合に違和感が少ないというのも気に入っている。
もうちょっとだけ線を太くして文字間を詰めてくれるといいんだけどねぇ。

[^jp]: Noto Sans JP は Noto Sans Japanese から 350 のウェイトを削除しているらしい。アホなブラウザが 350 という値をうまくハンドリングできないからだそうな。まぁ文章の中で使うだけなら多くても 400 と 700 のふたつがあれば充分なのでこれで問題ないし，ダウンロードサイズが小さくなるのも魅力である。

Kokoro フォントを導入するには CSS で以下のようにインポートする。

```css
@import url(http://fonts.googleapis.com/earlyaccess/kokoro.css);
```

あるいは各ページのヘッダ部分で

```html
<link rel='stylesheet' href='http://fonts.googleapis.com/earlyaccess/kokoro.css' type='text/css'>
```

としてもよい。
この CSS ファイルの中身は以下のようになっていて

```css
/*
 * Kokoro (Japanese) https://fonts.google.com/earlyaccess
 */
@font-face {
  font-family: 'Kokoro';
  font-style: normal;
  font-weight: 400;
  src: url(//fonts.gstatic.com/ea/kokoro/v1/Kokoro-Regular.eot);
  src: url(//fonts.gstatic.com/ea/kokoro/v1/Kokoro-Regular.eot?#iefix) format('embedded-opentype'),
       url(//fonts.gstatic.com/ea/kokoro/v1/Kokoro-Regular.woff2) format('woff2'),
       url(//fonts.gstatic.com/ea/kokoro/v1/Kokoro-Regular.woff) format('woff'),
       url(//fonts.gstatic.com/ea/kokoro/v1/Kokoro-Regular.ttf) format('truetype');
}
```

Kokoro フォントを指定するには

```css
body {
  font-family: 'Noto Serif', 'Kokoro', serif;
  font-weight: 400; /* normal */
}
```

とすればいいことが分かる。
ウェイトが1種類しかないが，明朝体を太字で表示することは多分ないので，これで OK。

Web フォントを Google のようなところから取得するのはメリットがある。
一度ダウンロードした Web フォントはブラウザ内でキャッシュされるため，みんなが同じ場所にあるフォントを使えばネット全体で通信量を減らすことができるからだ。

本当は Noto Serif で日本語をサポートしてくれると嬉しいんだけどねぇ。

[こころ明朝体]: http://typingart.net/?p=46 "日本語フォント「こころ明朝体」 - フォント無料ダウンロード｜Typing Art"