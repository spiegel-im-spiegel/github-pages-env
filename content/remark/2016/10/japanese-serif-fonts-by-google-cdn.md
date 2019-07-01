+++
date = "2016-10-16T23:41:57+09:00"
update = "2017-04-22T18:28:05+09:00"
title = "Web フォントに関する覚え書き（明朝体編）"
description = "最近 Google の Early Access のページを見たら随分と日本語の書体が増えている気がする。"
tags = ["web", "font", "google", "character", "site"]
draft = false

[author]
  name = "Spiegel"
  instagram = "spiegel_2007"
  twitter = "spiegel_2007"
  tumblr = ""
  facebook = "spiegel.im.spiegel"
  license = "by-sa"
  github = "spiegel-im-spiegel"
  flickr = "spiegel"
  url = "https://baldanders.info/profile/"
  avatar = "/images/avatar.jpg"
  linkedin = "spiegelimspiegel"
  flattr = ""
+++

以前，[ここで使っている Web フォントの話]({{< ref "/remark/2015/web-font-family.md" >}} "Web フォントに関する覚え書き")を書いた。
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
- Noto Sans JP[^jp]

[^jp]: Noto Sans JP は Noto Sans Japanese から 350 のウェイトを削除しているらしい。アホなブラウザが 350 という値をうまくハンドリングできないからだそうな。まぁ文章の中で使うだけなら多くても 400 と 700 のふたつがあれば充分なのでこれで問題ないし，ダウンロードサイズが小さくなるのも魅力である。

このうち Hannari, Kokoro, Sawarabi Mincho が Serif 相当の書体になる。
それぞれの見本をみて，私は Sawarabi Mincho を選択した。

[さわらびフォント](http://sawarabi-fonts.osdn.jp/)は明朝体ではまだ第一水準漢字も網羅されていないが，とりあえず鼻の先は問題なさそうである[^kokoro]。
Sawarabi Mincho を導入するには CSS で以下のようにインポートする。

<!-- [こころ明朝体]はひらがな・カタカナをデザインしたものだが，それ以外の文字は [IPA フォント](http://ipafont.ipa.go.jp/ "IPAexフォント/IPAフォント | IPAフォントのダウンロードサイトです")で補完しているため問題ないと思われる。
あと Noto Sans JP と組み合わせた場合に違和感が少ないというのも気に入っている。
もうちょっとだけ線を太くして文字間を詰めてくれるといいんだけどねぇ。 -->

[^kokoro]: 本当は[こころ明朝体]にしたかったのだが，どうもひらがなとカタカナしかサポートしていないらしい。最初[こころ明朝体]にしてみて，ケータイでの表示が大変なことになったので引っ込めた。

```css
@import url(http://fonts.googleapis.com/earlyaccess/sawarabimincho.css);
```

あるいは各ページのヘッダ部分で

```html
<link rel='stylesheet' href='http://fonts.googleapis.com/earlyaccess/sawarabimincho.css' type='text/css'>
```

としてもよい。
この CSS ファイルの中身は以下のようになっていて

```css
/*
 * Sawarabi Mincho (Japanese) https://fonts.google.com/earlyaccess
 */
@font-face {
  font-family: 'Sawarabi Mincho';
  font-style: normal;
  font-weight: 400;
  src: url(//fonts.gstatic.com/ea/sawarabimincho/v1/SawarabiMincho-Regular.eot);
  src: url(//fonts.gstatic.com/ea/sawarabimincho/v1/SawarabiMincho-Regular.eot?#iefix) format('embedded-opentype'),
       url(//fonts.gstatic.com/ea/sawarabimincho/v1/SawarabiMincho-Regular.woff2) format('woff2'),
       url(//fonts.gstatic.com/ea/sawarabimincho/v1/SawarabiMincho-Regular.woff) format('woff'),
       url(//fonts.gstatic.com/ea/sawarabimincho/v1/SawarabiMincho-Regular.ttf) format('truetype');
}
```

Sawarabi Mincho を指定するには

```css
body {
  font-family: 'Noto Serif', 'Sawarabi Mincho', serif;
  font-weight: 400; /* normal */
}
```

とすればいいことが分かる。
ウェイトが1種類しかないが，明朝体を太字で表示することは多分ないので，これで OK。

Web フォントを Google のようなところから取得するのはメリットがある。
一度ダウンロードした Web フォントはブラウザ内でキャッシュされるため，みんなが同じ場所にあるフォントを使えばネット全体で通信量を減らすことができるからだ。

本当は Noto Serif で日本語をサポートしてくれると嬉しいんだけどねぇ。
でもフォントの制作は時間とお金と人手がかかる仕事だし，ただ成果を享受している身としてはワガママは言えません。

## ブックマーク

- [Google Developers Japan: Noto Serif CJK が登場！](https://developers-jp.googleblog.com/2017/04/noto-serif-cjk-is-here.html)
    - [Adobe、「源ノ明朝」フォントをリリース ～「源ノ角ゴシック」と対になるセリフ書体 - 窓の杜](http://forest.watch.impress.co.jp/docs/news/1052973.html)
    - [Google、日中韓対応の明朝体フォント「Noto Serif CJK」を無償公開 - 窓の杜](http://forest.watch.impress.co.jp/docs/news/1052998.html)
- [Noto Serif(源ノ明朝)のCSS指定 - Qiita](http://qiita.com/_RJ/items/645adf95ed6f5841eaf6)
- [日中韓に対応したグーグルの新フォント「Noto Serif CJK」は、なぜ生まれたか｜WIRED.jp](http://wired.jp/2017/04/08/noto-serif-cjk/)

[こころ明朝体]: http://typingart.net/?p=46 "日本語フォント「こころ明朝体」 - フォント無料ダウンロード｜Typing Art"
