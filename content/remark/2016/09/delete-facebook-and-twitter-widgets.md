+++
title = "Facebook と Twitter の Widget を削除した"
date = "2016-09-25T15:11:48+09:00"
description = "まっ，気にしない人は気にしないんだろうけど，こんな辺境の個人サイトにまで追跡コードがくっついてるなんて（私が）ヤな感じなので外すことにした。"
tags = ["facebook", "twitter", "tracking"]

[scripts]
  mathjax = false
  mermaidjs = false
+++

いいわけだけど，あんまり自分のサイトとかちゃんと見ないじゃないですか。
で，最近気がついたんだけど， [F-Secure](https://www.f-secure.com/) の [Freedome](https://www.f-secure.com/freedome) 越し，あるいは [Firefox](https://www.mozilla.org/firefox/) の [private browsing](https://support.mozilla.org/kb/private-browsing-use-firefox-without-history) だと Twitter や Facebook のコードをブロックするんだね。

まっ，気にしない人は気にしないんだろうけど，こんな辺境の個人サイトにまで追跡コードがくっついてるなんて（私が）ヤな感じなので外すことにした。
ちなみに Twitter の tweet ボタンは以下のようにした。

```html
<a href="http://twitter.com/share?text={{ $.Title }}&amp;url={{ $.Permalink }}" rel=”nofollow” target="_blank"><i class="fab fa-twitter-square fa-2x" aria-hidden="true"></i></a>
```

Facebook の share ボタンはこんな感じ。

```html
<a href="http://www.facebook.com/share.php?u={{ $.Permalink }}" rel=”nofollow” target="_blank"><i class="fab fa-facebook-square fa-2x" aria-hidden="true"></i></a>
```

あっ。
`{{ $.Title }}` とか `{{ $.Permalink }}` とかは [Hugo] （つか [Go 言語]の）の[テンプレート](https://golang.org/pkg/text/template/ "template - The Go Programming Language")文法なので，あまり気にしないように。
アイコンは [Font Awesome](https://fortawesome.github.io/Font-Awesome/ "Font Awesome, the iconic font and CSS toolkit") で。

まぁ，こんな感じでいいだろう。

[本家サイト](https://baldanders.info/ "Baldanders.info")は申し訳ないが放置。
全ページ手動で直さないといけないので今は無理。

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[Go 言語]: https://golang.org/ "The Go Programming Language"
