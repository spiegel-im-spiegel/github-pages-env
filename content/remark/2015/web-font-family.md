+++
date = "2015-10-21T21:41:45+09:00"
update = "2018-05-10T22:46:43+09:00"
description = "このサイトの Web フォントまわりを整理したので，覚え書きとして残しておく。"
draft = false
tags = ["character", "font", "web", "google", "ipa", "site"]
title = "Web フォントに関する覚え書き"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"
+++

このサイトの Web フォントまわりを整理したので，覚え書きとして残しておく。

## Google Noto Fonts

このサイトの Web フォントを Google Noto Fonts で統一した[^a]。

[^a]: “Noto” は “No-Tofu” の意味らしい。昔はフォント・セットに該当文字がない場合に表示される「□」を「トーフ（豆腐）」と呼んでいた。つまり Noto Fonts は世界中のすべての文字を統一デザインで網羅するという壮大なプロジェクトである。

Google Noto Fonts は以下のサイトで取得できる。

- [Google Fonts](https://www.google.com/fonts/)

このサイトでは様々なフォントが選べるが “Noto” をキーワードに以下のフォントを選択する。

- Noto Sans : Normal 400, Bold 700
- Noto Serif : Normal 400, Bold 700

またサブセットとして Latin を選択する。
これで，以下に示す書式で CSS ファイルを HTML の `head` 要素内に指定すればよい[^aa]。

[^aa]: [Google Fonts](https://www.google.com/fonts/) サイトでは `@import` 文を使った書式や JavaScript による読み込みコードも例示されている。

```html
<link href='//fonts.googleapis.com/css?family=Noto+Sans:400,700|Noto+Serif:400,700' rel='stylesheet' type='text/css'>
```

Bold が不要であればもっと軽くできる[^b]。

[^b]: Modern browser では Bold や Italic がなくとも機械的に合成される（Italic は斜体（slant）で代用）。当然ながら綺麗ではない。 Bold や Italic を多用しないのであれば Web フォントから外す手はある。

```html
<link href='//fonts.googleapis.com/css?family=Noto+Sans|Noto+Serif' rel='stylesheet' type='text/css'>
```

CSS でフォントを指定する際は `"Noto Sans"` や `"Noto Serif"` で指定する。

```css
body {
  font-family: 'Noto Serif', serif;
  font-weight: 400; /* normal */
  font-variant: normal;
  font-style: normal;
  font-size: medium;
  line-height: 1.4;
}
h1, h2, h3, h4, h5, h6, address, table th, dl > dt, em, strong, cite {
  font-family: 'Noto Sans', sans-serif;
  font-weight: 400; /* normal */
  font-variant: inherit;
  font-style: inherit;
  font-size: inherit;
}
```

### Noto Sans Japanese

2014年， Noto Sans に Noto Sans CJK が加わった。

- [Noto Sans CJK – Google Noto Fonts](https://www.google.com/get/noto/help/cjk/)
- [オープンソースの美しい Noto フォントファミリーに日本語、中国語、韓国語が加わりました。 - Google Developer Japan Blog](http://googledevjp.blogspot.jp/2014/07/noto.html)

Noto Sans CJK を Web フォントとして使用するなら以下のサイトが参考になる。

- [Google Fonts : Early Access](https://www.google.com/fonts/earlyaccess)

日本語であれば “Noto Sans JP を導入すればよい[^jp]。

[^jp]: Noto Sans JP は Noto Sans Japanese から 350 のウェイトを削除しているらしい。アホなブラウザが 350 という値をうまくハンドリングできないからだそうな。まぁ文章の中で使うだけなら多くても 400 と 700 のふたつがあれば充分なのでこれで問題ないし，ダウンロードサイズが小さくなるのも魅力である。

```html
<link href='//fonts.googleapis.com/earlyaccess/notosansjp.css' rel='stylesheet' type='text/css'>
```

フォント名は `"Noto Sans JP"` で指定する。

```css
body {
  font-family: 'Noto Serif', serif;
  font-weight: 400; /* normal */
  font-variant: normal;
  font-style: normal;
  font-size: medium;
  line-height: 1.4;
}
h1, h2, h3, h4, h5, h6, address, table th, dl > dt, em, strong, cite {
  font-family: 'Noto Sans', 'Noto Sans JP', sans-serif;
  font-weight: 400; /* normal */
  font-variant: inherit;
  font-style: inherit;
  font-size: inherit;
}
```

Noto Sans JP では JIS X 0208 の文字集合を収録しているため， JIS X 0213 第3第4水準漢字はサポートされていない模様。
これらの漢字が必要な場合は Noto Sans CJK の日本語フルセットが必要となる。

### Inconsolata

プログラム・コードの表示用に Inconsolata を導入する。
Inconsolata は [Google Fonts](https://www.google.com/fonts/) を使って導入可能。
最終的に Google Fonts の書式は以下のようになる。

```html
<link href='//fonts.googleapis.com/css?family=Noto+Serif|Noto+Sans|Inconsolata:400,700' rel='stylesheet' type='text/css'>
<link href='//fonts.googleapis.com/earlyaccess/notosansjapanese.css' rel='stylesheet' type='text/css'>
```

また CSS の記述は以下のようになる。

```css
body {
  font-family: 'Noto Serif', serif;
  font-weight: 400; /* normal */
  font-variant: normal;
  font-style: normal;
  font-size: medium;
  line-height: 1.4;
}
h1, h2, h3, h4, h5, h6, address, table th, dl > dt, em, strong, cite {
  font-family: 'Noto Sans', 'Noto Sans JP', sans-serif;
  font-weight: 400; /* normal */
  font-variant: inherit;
  font-style: inherit;
  font-size: inherit;
}
pre, tt, code, var, kbd, samp {
  font-family: 'Inconsolata', 'Noto Sans JP', monospace;
  font-weight: 400; /* normal */
  font-variant: inherit;
  font-style: inherit;
  font-size: inherit;
}
```

### Noto Serif Japanese が欲しい！

残念ながら，現時点では Noto Serif の日本語フォントは存在しない。
有料で明朝体 Web フォントを提供してくれるサービスはあるが，デザイン重視のサイトならともかく，うちのようにただ文字が多いだけのサイトでは「お金を払ってまでねぇ」という感じ。
そこで，いくつかオープンなフォントを探してみる。

#### IPA 明朝

- [IPAexフォント/IPAフォント | IPAフォントのダウンロードサイトです](http://ipafont.ipa.go.jp/)

オープンなフォントといえば IPA フォントだよね。
[OSD (Open Source Definition)](http://www.opensource.jp/osd/osd-japanese.html) に準拠した[ライセンス](http://ipafont.ipa.go.jp/ipa_font_license_v1-html)で提供されているため扱いやすい[^c]。

[^c]: 「[さまざまなライセンスとそれらについての解説 - GNUプロジェクト](http://www.gnu.org/licenses/license-list.ja.html)」では [IPA フォントライセンスは GPL とは両立しない](http://www.gnu.org/licenses/license-list.ja.html#IPAFONT)とある。派生作品ではオリジナルの名前を含めてはいけない，という制限があるため。

ただし IPA フォントは PDF 文書のようなものには向いてるけど， Web フォントには向いていない。
理由は以下の2つ。

- 全体的に線が細く表示環境によっては文字がかすれてしまう。 ウェイトが1種類しかないのもマイナスポイント
- JIS X 0213 の文字集合を全て収録しているのはいいが，データサイズが大きくなってしまう。第1第2水準漢字のみのサブセットがあればいいのに

#### あおぞら明朝

- [あおぞら明朝フォント](http://blueskis.wktk.so/AozoraMincho/ "ホーム | あおぞら明朝フォント")

IPA 明朝（厳密には IPA P明朝）からの fork で，7ウェイトをサポートしているのが特徴。
ライセンスも [IPA フォントライセンス](http://ipafont.ipa.go.jp/ipa_font_license_v1-html)を継承している。
欠点は以下のとおり。

- ベースがプロポーショナル・フォントなので見た目が詰まった感じになってしまう
- IPA フォントと同じく JIS X 0213 の文字集合を全て収録しているのはいいが，データサイズが大きくなってしまう

#### やさしさアンチック

- [フリーフォントやさしさアンチック（漫画用書体・セリフフォント）のダウンロード | フォントな。無料日本語フリーフォント](http://www.fontna.com/blog/1122/)

アンチック体あるいは「アンチゴチ」と呼ばれる書体で，漢字はゴシック体，かな文字等は明朝体で表現されるのが特徴。
IPA フォントおよび [M+ フォント](https://mplus-fonts.osdn.jp/mplus-outline-fonts/)がベースになっているため，ライセンスもこれらに準拠する形になっている。

見出しや短文に向いている。
長文はちょっとしんどいか。
第1第2水準漢字までをサポートしており，サイズ的にも手頃といえる。

### もう WOFF でいいじゃない

複製・再配布が自由なフォントであれば WOFF (Web Open Font Format) に変換[^d] して Web フォントとして利用可能。
WOFF への変換には，以下のツールが使える。

[^d]: 機械的変換であれば翻案ではなく「逐語的コピー」とみなせる。

- [WOFFコンバータ](http://opentype.jp/woffconv.htm) （[武蔵システム](http://opentype.jp/)）

作成した WOFF ファイルを使って CSS 内に `@font-face` を定義する。

```css
@font-face{
    font-family: 'aozora-m';
    src: local('AozoraMinchoMedium'), url('/fonts/AozoraMinchoMedium.woff') format('woff');
    font-weight: 400;
    font-style: normal;
}
```

これで定義したフォント名 `"aozora-m"` を使って CSS に記述することができる。

```css
body {
  font-family: 'Noto Serif', "aozora-m", serif;
  font-weight: 400; /* normal */
  font-variant: normal;
  font-style: normal;
  font-size: medium;
  line-height: 1.4;
}
```

Web フォントの形式はいくつかあるが，少なくとも Modern browser であれば WOFF をサポートしている。
IE のことは忘れましょう（笑）

## その他の Web フォント

Google Noto Fonts 以外で利用しているフォントを列挙しておく。

### Font Awesome

[Font Awesome] は 5 系がリリースされ， SVG と JavaScript によるアイコン表示と制御がサポートされた。
詳しくは以下の記事を参照のこと。

- [Font Awesome 5 がリリースされた]({{< ref "/remark/2017/12/font-awesome-5-released.md" >}})

[Font Awesome]: https://fontawesome.com/

### Creative Commons Icon Font

- [Creative Commons Icon Font](http://cc-icons.github.io/)
- [cc-icons/cc-icons](https://github.com/cc-icons/cc-icons)

CC License のアイコンは [Font Awesome] 5.0.11 以降に含まれた。

```html
<i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-sa fa-2x"></i>
```

表示はこんな感じ。

{{< fig-quote >}}
<i class="fab fa-creative-commons-by fa-2x"></i>&nbsp;<i class="fab fa-creative-commons-sa fa-2x"></i>
{{< /fig-quote >}}

ブラボー！

## 関連（するかもしれない）記事

- [ちょこっと MathJax — Baldanders.info](https://baldanders.info/spiegel/log2/000750.shtml) : MathJax で数式用の Web フォントを使用可能
- [Web フォントに関する覚え書き（明朝体編）]({{< ref "/remark/2016/10/japanese-serif-fonts-by-google-cdn.md" >}}) : 明朝体について追記している
