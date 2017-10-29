+++
title = "Twitter Card メタデータに対応した"
date =  "2017-10-29T16:00:38+09:00"
update =  "2017-10-29T21:51:56+09:00"
description = "もう Semantic Web なんか誰も見向きもしなくなってるみたいだし， Web コンテンツを解析する手段は AI 技術を利用したものへシフトしてるようだし，もう（どうでも）いいかな，と。"
tags        = [ "site", "semantic", "web", "twitter", "open-graph" ]

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

なんとなく思いついてこのブログをいわゆる “[Twitter Card]” に対応させた。

実は [Twitter Card] が登場し始めたときに Tumblr などで適用していたのだが，あまりに酷い仕様で [HTML Validator] にかけるとエラーの嵐になるし，そもそも Semantic Web を無視した設計に腹が立って忌避していたのだ。
しかし，まぁ，もう Semantic Web なんか誰も見向きもしなくなってるみたいだし， Web コンテンツを解析する手段は AI 技術を利用したものへシフトしてるようだし，もう（どうでも）いいかな，と。

[Twitter Card] の仕様も随分シンプルになった。
まず，カードのタイプが以下の4種類のみになった。

- [Summary card]
- [Summary with large image]
- [Player card]
- [App card]

[Player card] は動画やスライドショウのページ用， [App card] はアプリページ用なので，それ以外の Web ページでは [Summary card] か [Summary with large image] を選択することになる。

[Summary with large image] はアイキャッチ[^ec1] 用の大きめの画像（300×157から4096×4096）を含む Card で，メディア・サイトなどが多用するあの**ウザい**やつである。
とはいえ，写真やイラストなどを中心としたサイトでは [Summary with large image] が向いているだろう。

[Summary with large image] にしないのなら [Summary card] を選択する。
うちのブログはもちろん~~オリーブオイル~~ [Summary card] で[^swli1]。

[^ec1]: この記事を書くにあたってちょっと調べたのだが「アイキャッチ」というのは和製英語なんだそうだ。ただし，コメントで頂いた情報では “eye-cacher” という言葉はあるらしい（thx!）。なお “eye” は “attention” に置き換えることができるそうで，その場合は [“catch the reader’s attention” みたいな言い回し](アイキャッチ画像の「アイキャッチ」とは？正しい英語ではどういう？ "アイキャッチ画像の「アイキャッチ」とは？正しい英語ではどういう？")になるとか。ふぅ。英語は難しいぜ。
[^swli1]: 実は一度 [Summary with large image] を試したのだが，自分で眺めてやっぱりウザかったので [Summary card] にした。

[Summary card] で必須のメタデータは以下の2つである。
これを指定しないと，そもそもカードが表示されない。

- `twitter:card`
- `twitter:title`

具体的には `<meta>` 要素を使って以下のように記述する。

```html
<meta name="twitter:card" content="summary">
<meta name="twitter:title" content="Codic API を利用するパッケージを作ってみた — プログラミング言語 Go | text.Baldanders.info">
```

しかし，このままではページの説明（description）やアイコン画像が表示されないため（何故 optional なのに無理くり表示しようとするのだろう），以下のメタデータも追加する。

- `twitter:description`
- `twitter:image`

アイコン画像には144×144から4096×4096までのサイズが使える。
これらを合わせると以下の記述が最低限必要と言える。

```html
<meta name="twitter:card" content="summary">
<meta name="twitter:title" content="Codic API を利用するパッケージを作ってみた — プログラミング言語 Go | text.Baldanders.info">
<meta name="twitter:description" content="spf13/viper を使ってみたかったのだ。">
<meta name="twitter:image" content="http://text.baldanders.info/images/attention/go-code.png">
```

以上のメタデータを `<head>` 要素内に設置する。
これで以下のように表示される（筈）。

{{< fig-img src="https://farm5.staticflickr.com/4510/37949847556_0867f5741b_o.png" title="Twitter Card: Summary"  link="https://www.flickr.com/photos/spiegel/37949847556/" >}}

以前は `twitter:image` 等を property-content 属性の組み合わせで記述させようとしていたが（これのせいで [HTML Validator] がエラーを吐いていた），さすがに改心したようである（笑）

サイトのオーナーやページの作成者が Twitter ユーザの場合は以下のメタデータも使える。

- `twitter:site`
- `twitter:creator`

これも同じように

```html
<meta name="twitter:site" content="@spiegel_2007">
<meta name="twitter:creator" content="@spiegel_2007">
```

とすればよい。
カードの見た目には全く関係ないが Twitter アナリティクスか何かで使うのだろう，多分。

メタデータの幾つかは [OGP] の語彙と置き換えることができる。

一応説明しておくと， [OGP] は Facebook が最初に考えた仕様で，元々はネット上のコンテンツと Facebook のアプリを関連付けて制御する仕組みだったのだが

{{< fig-img src="http://www.baldanders.info/spiegel/archive/rdfa/ogp.svg" title="Open Graph の相関図"  link="http://www.baldanders.info/spiegel/archive/rdfa/ogp.svg" >}}

RDFa の仕様の一部を借用した大変筋の悪いもので，本来の目的は明後日方向に飛んでいき，現在は `<head>` 要素内にメタデータを記述するための迂遠な手段に堕している[^rdfa1]。
ただし [OGP] で記述したメタデータを参照するサービスは多いため，今だに SEO 対策として用いられているようだ。

[^rdfa1]: RDF/RDFa は Web 上の（URI で記述可能な）あらゆるリソース同士の関係を「主語・述語・目的語」の3つ組（triple）で表すことで machine-understandable な「意味」を与える Semantic Web の実装のひとつである。もちろん Facebook はそんな思想背景など微塵も考慮していなかったと思うが。

[Twitter Card] のメタデータと置き換え可能な [OGP] の述語を以下に示す。

| [Twitter Card] | [OGP] |
|:---------------|:------|
| `twitter:card`        | `og:type` |
| `twitter:description` | `og:description` |
| `twitter:title`       | `og:title` |
| `twitter:image`       | `og:image` |

これらの述語を既に使っている場合は [Twitter Card] のメタデータで記述する必要はない。
なお `og:type` は本来はメディア・タイプ（MIME タイプや RDF/RDFa の語彙で定義されるタイプ）を指定するものなので [Twitter Card] 用に使うべきではない。

[OGP] は（一応） RDFa の仕様に従っているのでメタデータ指定には name-content 属性の組み合わせではなく property-content 属性の組み合わせで記述する。
例えば以下の通り[^ogp1]。

[^ogp1]: `og:image` の目的語は URL なので， RDFa 仕様としては property-content 属性ではなく（`<link>` 要素を使って） rel-resource 属性または rel-href 属性の組み合わせとするのが正しい。のだが， [OGP] は馬鹿なので property-content 属性とする（つまりリテラル・データとして扱う）よう求めている。私が間違ってるわけではない，決して。ホントなんだよこのクソ仕様は。ブツブツブツ...

```html
<head  prefix="og: http://ogp.me/ns#">
  ...
  <meta property="og:title" content="Codic API を利用するパッケージを作ってみた — プログラミング言語 Go | text.Baldanders.info">
  <meta property="og:description" content="spf13/viper を使ってみたかったのだ。">
  <meta property="og:image" content="http://text.baldanders.info/images/attention/go-code.png">
  ...
</head>
```

`<head>` 要素に `prefix` 属性を付けるのを忘れずに（[HTML5] の場合）。
なぜ `<html>` 要素ではなく `<head>` 要素に付けるかというと，  [OGP] は `<head>` 要素外での使用を考慮していないからである。
まぁ Twitter 側はそんなこと微塵も気にしてないだろうけど。
おそらく接頭辞を `og` 以外にしたら Twitter 側は認識できなくなるんじゃないのかな（馬鹿らしいので試さない）[^prfx1]。

[^prfx1]: ちなみに `prefix` 属性ではなく `vocab` 属性を使えば接頭辞が不要になる。どうなるかなんて試さないよ（笑）

ちなみに `<head>` 要素内の `<title>` 要素や `description` メタデータはまるっと無視するようである。
本当に何でこんな頭の悪い仕様になってるのか知らないが[^ttl1]，最初に書いたように，いまさら Semantic Web なんか気にする人はいないだろうし，どうでもいいか。

[^ttl1]: ちなみに Facebook は `<title>` 要素を認識して使っている。はっきり言って `<title>` 要素と `twitter:title` メタデータが独立して存在してるってのは詐欺の匂いがするんだが，誰も気にしないのかね。アイキャッチ画像につられて詐欺みたいなサイトに誘導されるってのは Facebook でも見られるが（笑）

## ブックマーク

- [Card Validator | Twitter Developers](https://cards-dev.twitter.com/validator)
- [ツイートにページ情報を表示する「Twitterカード（Twitter Cards）」を設定してみた | 株式会社グランフェアズ](http://www.granfairs.com/blog/staff/setting-twitter-cards)
- [【2017年版】Twitterカードとは？使い方と設定方法まとめ](https://saruwakakun.com/html-css/reference/twitter-card)

- [RDFa 入門 — Baldanders.info](http://www.baldanders.info/spiegel/archive/rdfa/)
- [タイムラインの奴隷 - Spiegel's Branch - Scrapbox](https://scrapbox.io/spiegel-branch/%E3%82%BF%E3%82%A4%E3%83%A0%E3%83%A9%E3%82%A4%E3%83%B3%E3%81%AE%E5%A5%B4%E9%9A%B7)

[Twitter Card]: https://developer.twitter.com/en/docs/tweets/optimize-with-cards/overview "Summary card — Twitter Developers"
[Summary card]: https://developer.twitter.com/en/docs/tweets/optimize-with-cards/overview/summary "Summary card — Twitter Developers"
[Summary with large image]: https://developer.twitter.com/en/docs/tweets/optimize-with-cards/overview/summary-card-with-large-image "Summary with large image — Twitter Developers"
[Player card]: https://developer.twitter.com/en/docs/tweets/optimize-with-cards/overview/player-card "Player card — Twitter Developers"
[App card]: https://developer.twitter.com/en/docs/tweets/optimize-with-cards/overview/app-card "App card — Twitter Developers"
[HTML Validator]: https://validator.w3.org/ "The W3C Markup Validation Service"
[OGP]: http://ogp.me/ "The Open Graph protocol"
[HTML5]: https://www.w3.org/TR/html5/

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/483993195X/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51oaN2iq9xL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/483993195X/baldandersinf-22/">セマンティック HTML/XHTML</a></dt><dd>神崎 正英 </dd><dd>毎日コミュニケーションズ 2009-05-28</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4627829310/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4627829310.09._SCTHUMBZZZ_.jpg"  alt="セマンティック・ウェブのためのRDF/OWL入門"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873114527/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873114527.09._SCTHUMBZZZ_.jpg"  alt="セマンティックWeb プログラミング"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4764904276/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4764904276.09._SCTHUMBZZZ_.jpg"  alt="Linked Data: Webをグローバルなデータ空間にする仕組み"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274202925/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274202925.09._SCTHUMBZZZ_.jpg"  alt="オントロジー構築入門"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4501542101/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4501542101.09._SCTHUMBZZZ_.jpg"  alt="トピックマップ入門 (セマンティック技術シリーズ)"  /></a> </p>
<p class="description">残念ながら紙の本は実質的に絶版なんですよねぇ。是非デジタル化を希望します。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2014-08-17">2014/08/17</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
