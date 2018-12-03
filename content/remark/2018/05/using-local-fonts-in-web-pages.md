+++
title = "Web ページでローカル・フォントを使う"
date = "2018-05-13T14:52:10+09:00"
description = "なんだよ，それ。馬鹿すぎる。"
image = "/images/attention/kitten.jpg"
tags        = [ "web", "font", "engineering" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

いや，個人的には Web フォントで NOTO フォント使えばいいぢゃんって思ってるけど，仕事だと色々と思惑も絡むので簡単には行かないわけさ。
んで Web ページでローカル・フォントを指定する方法についてちょろんと調べてみたので，以下に覚え書きとして残しておく。

やっぱり問題は日本語フォント。
しかも Windows と macOS の差異だろう。
こんな感じ。

- Windows
    - 游ゴシック（Yu Gothic）
    - 游明朝（Yu Mincho）
    - メイリオ（Meiryo）
    - ＭＳ Ｐゴシック（MS PGothic）
- macOS
    - 游ゴシック体（YuGothic）
    - 游明朝体（YuMincho）
    - ヒラギノ角ゴ Pro W3（Hiragino Kaku Gothic Pro）
    - Osaka

これだけ見たら「游明朝／游ゴシック」で統一すればいいぢゃん，と思うでしょ。
でも「游明朝／游ゴシック」は Windows 7 には標準で入っていない[^win1]。
更に言うと Windows と macOS では「游明朝／游ゴシック」の中身が違うのだ。

[^win1]: Windows Vista, XP は論外ね。つまり IE 11 より前のバージョンはもう考えなくていいと思う。ベストは一刻も早く IE が滅亡することだけど。

- Windows 用の「游明朝／游ゴシック」には一部の異体字が入っていない
- Windows 用の「游明朝／游ゴシック」の ウェイトは Regular サイズが標準だが macOS 用は Regular サイズが存在しない

なんだよ，それ `orz`

まだあるぞ。

Windows の Internet Explorler (IE) で日本語のローカル・フォントを指定する際には日本語名で指定しないといけないが，どうやらこれが IE だけの特殊ルールらしいのだ。

sigh...

結局 Web ページで「游明朝／游ゴシック」を指定したいなら

```css
div.gothic {
  font-family: "游ゴシック Medium", "Yu Gothic Medium", "YuGothic", sans-serif;
}
```

みたいな感じにしないといけないらしい。
馬鹿すぎる。

これではあまりにカオスなので以下のように `@font-face` を定義してみた。

```css
@font-face {
  font-family: 'HiGothic-local';
  font-style: normal;
  src: local("ヒラギノ角ゴ Pro W3"), local("Hiragino Kaku Gothic Pro");
}
@font-face {
  font-family: 'YuGothic-local';
  font-style: normal;
  src: local("游ゴシック Medium"), local("Yu Gothic Medium"), local("游ゴシック体"), local("YuGothic");
}
@font-face {
  font-family: 'MSGothic-local';
  font-style: normal;
  src: local("ＭＳ Ｐゴシック"), local("MS PGothic");
}
@font-face {
  font-family: 'SansJP-local';
  font-style: normal;
  src: local("メイリオ"), local(Meiryo), local(Osaka) ;
}
```

こうしておけば

```css
div.gothic {
  font-family: 'HiGothic-local', 'YuGothic-local', 'SansJP-local', 'MSGothic-local', sans-serif;
}
```

という感じに多少スッキリするのではないかと。

ちなみに iOS はフォントの種類が限られている（日本語はヒラギノ・フォント）ので serif /  sans-serif を指定すれば対応できる。
更に Android には明朝体がない上，ゴシック体についても端末メーカが各々独自の日本語フォントを入れているらしい。
だから日本はガラｐ...

くぁｗせｄｒｆｔｇｙふじこｌｐ

...やっぱ Web フォントでいいと思うなぁ。
ローカル・フォントは serif /  sans-serif で指定すれば十分だろう。

## ブックマーク

- [font-family【CSSリファレンス】 | Webデザインラボ](https://www.webdlab.com/css/font-text-font-family/)
- [font-familyには何を指定すればいい？ - csstux.com](http://www.csstux.com/font-family.html)
- [2017年版個人的ベスト｜CSSフォント指定WIN、MACでキレイに表示（ゴシック、明朝）｜MWORKS](https://mw-s.jp/2017css-font-win-mac/)
- [Android における最適なフォント環境を考えてみる – ミルログ](https://www.mirucon.com/2017/03/04/android-font-family/)
