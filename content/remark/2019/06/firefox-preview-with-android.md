+++
title = "Firefox Preview を試してみた"
date =  "2019-06-28T22:07:04+09:00"
description = "ここは生暖かく「今後に期待」と見守っておこう。 "
image = "/images/attention/kitten.jpg"
tags = [ "firefox", "android" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Android に [Firefox Preview] なるものが登場したらしい。

- [Android版「Firefox」のテスト用プレビューが公開--高速化、プライバシー強化 - ZDNet Japan](https://japan.zdnet.com/article/35139189/)
- [Mozillaが最大2倍高速なAndroid版Firefoxをプレビュー  |  TechCrunch Japan](https://techcrunch.com/2019/06/27/mozilla-previews-a-redesigned-and-faster-firefox-for-android/)

さっそく導入してみた。
こんな感じ。

{{< fig-img src="./firefox-preview-1.png" link="./firefox-preview-1.png" width="1440" >}}

他の普段使いのブラウザと区別するためにダークモードにしている。
設定画面はこんな感じ。

{{< fig-img src="./firefox-preview-3.png" link="./firefox-preview-3.png" width="1440" >}}

「トラッキング防止」は最初からオンになっているが検索エンジンを DuckDuckGo にできないという相変わらずのクソ野郎ぶりである。

ちなみにパソコン用の Firefox Quantum ではトラッキング防止機能は2段階あって

{{< fig-img src="/release/2019/05/firefox-67-is-released/block-list.png" title="Firefox 67 がリリースされた" link="/release/2019/05/firefox-67-is-released/" width="877" >}}

レベル2に設定するとこのページのフィードバックに使っている Disqus もブロックされるくらい強力に効くのだが [Firefox Preview] ではそこまでに至らないないようだ。
この辺り，何とも中途半端な感じで Mozilla はユーザのプライバシーを重視していないということが今回の出来を見ても分かろうというものである（それでも Chrome よりはマシというのがにんともかんとも）。

はっきり言って [GeckoView] を使ったプライバシー重視のブラウザを使いたいなら [Firefox Focus] を使うことを強くお勧めする[^ddg1]。
一般のユーザが [Firefox Preview] を使うメリットは（現時点では）万にひとつもない。

文字通りの preview 版だとしても，同じブラウザエンジンを使った先行アプリがあるのだから，せめて同程度のセキュリティ・プライバシー設定が可能な程度の機能は実装してほしかった。
どうせ iOS/iPadOS 版は出ない（出せない）だろうから Android 特化でチューニングすればいい。

[^ddg1]: [Firefox Focus] の検索エンジンにも DuckDuckGo は入っていないが，追加して切り換えることはできる。

まぁ，ここは生暖かく「今後に期待」と見守っておこう。
とりあえずアプリは削除だな。

[Firefox Preview]: https://play.google.com/store/apps/details?id=org.mozilla.fenix "Firefox Preview - Google Play"
[GeckoView]: https://wiki.mozilla.org/Mobile/GeckoView "Mobile/GeckoView - MozillaWiki"
[Firefox Focus]: https://play.google.com/store/apps/details?id=org.mozilla.focus
