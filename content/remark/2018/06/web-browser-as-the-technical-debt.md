+++
title = "技術的負債としての Web ブラウザ"
date = "2018-06-04T21:12:37+09:00"
update = "2018-12-10T22:28:16+09:00"
description = "しかし，結局ユーザはプライバシーより利便性を選択し，同じ口で Facebook を避難するのだ。"
image = "/images/attention/kitten.jpg"
tags = [ "web", "code", "privacy", "engineering", "firefox" ]

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

- [Bye, Chrome: Why I’m switching to Firefox and you should too](https://www.fastcodesign.com/90174010/bye-chrome-why-im-switching-to-firefox-and-you-should-too)
- [今こそブラウザをChromeからFirefoxに乗り換えるべき理由 - YAMDAS現更新履歴](http://d.hatena.ne.jp/yomoyomo/20180604/chrometofirefox)

Firefox を使いだしたのっていつだっけ？ と昔の日記を掘り返してみたら，どうやら2003年頃らしい。
たった15年前か。

- [火の鳥導入 しっぽのさきっちょ 2003年04月 -- Spiegel's Trunk](http://www.baldanders.info/spiegel/log/200304.html#d29)

（ちなみに当時は Phoenix とか Firebird とかなかなか名前が定着しなかった）

それまでは Opera を（ちゃんとお金を払って）使っていたのだが，あまりのサポートのダメさ加減に見限ったのだった。
決して Firefox が優れていたから乗り換えたわけじゃない。

はっきり言うが Mozilla/Firefox がユーザのプライバシーに配慮しているというのは嘘っぱちである。
もし本当に Firefox がユーザのプライバシーに配慮していると言うなら既定で 3rd party cookie を無効にしているはずだし検索窓の標準も [DuckDuckGo] になっている筈である。
昨年リリースされた Quantum だって既定でトラッキングをブロックする設定にすべきだったのだ[^ff1]。
そうしなかったのは何故か。
営利企業ではない彼らもやはり市場原理の下に行動せざるを得ないからだ。

[^ff1]: 実際に携帯端末用の [Firefox Focus](https://play.google.com/store/apps/details?id=org.mozilla.focus "Firefox Focus: プライバシー保護ブラウザー - Google Play") は最初からトラッキングをブロックする設定になっている。

何よりも Mozilla は [Campbridge Analytica の件]({{< ref "/remark/2018/03/name-identification.md" >}} "善悪の葛藤")が発覚するまで Facebook に広告を出していたのである。
その広告から広告主たる Mozilla が何を得ていたか知らなかったとは言わせない。
あの件以降，そそくさと広告を引っ込めた企業・団体の名前は覚えておいたほうがいい。
連中は Facebook よりも卑劣である（目くそ鼻くそだけど）。

私たちに示されている選択肢はそう多くない。
企業製以外のブラウザが欲しければ Firefox か Firefox より fork した強化ブラウザを使うしかない。
検索サービスに至っては今や [DuckDuckGo] 一択である[^ddg1]。
しかし，結局ユーザはプライバシーより利便性を選択し，同じ口で Facebook を非難するのだ。
茶番にも程がある。

[^ddg1]: そういえば最初に "[Schneier on Security](https://www.schneier.com/)” を見たとき，サイト内検索窓が [DuckDuckGo] なのを見つけて「やっぱちゃんとした人はちゃんとしてるんだな」と改めて思ったものである。

こうしてブラウザは（市場原理の名の下に）いつまでもその技術的負債を背負い続けるのであろう。

## ブックマーク

- [GitHub - Eloston/ungoogled-chromium: Google Chromium, sans integration with Google](https://github.com/Eloston/ungoogled-chromium)
    - [Google Chromeからプライバシーを侵す機能を全削除したブラウザ「ungoogled-chromium」 - GIGAZINE](https://gigazine.net/news/20161109-ungoogled-chromium/)

[DuckDuckGo]: https://duckduckgo.com/ "DuckDuckGo — Privacy, simplified."
