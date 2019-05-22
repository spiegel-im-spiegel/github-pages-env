+++
title = "Firefox 67 がリリースされた"
date =  "2019-05-22T16:44:00+09:00"
description = "今まではマイニング防止に NoCoin を使っていたが，これからは標準機能でいけるかな。"
image = "/images/attention/tools.png"
tags  = [ "tools", "firefox", "security", "privacy" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Firefox 67 がリリースされた。

- [Firefox  67.0, See All New Features, Updates and Fixes](https://www.mozilla.org/en-US/firefox/67.0/releasenotes/)
- [「Firefox 67」が正式公開 ～暗号通貨採掘・フィンガープリンティング防止機能を搭載 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1185902.html)

修正点や変更点は色々あるが個人的には「暗号通貨マイニング」と「フィンガープリント採取」をブロックできるようになったことだろう。

{{< fig-img src="./firefox67.png" title="Firefox 67" link="https://www.flickr.com/photos/spiegel/47853743812/" width="834" >}}

今まではマイニング防止に [NoCoin] を使っていたが，これからは標準機能でいけるかな。
「フィンガープリント採取」をブロックできるようになったことは僥倖だろう。

[Ubuntu] の場合は APT でアップデートできる。
即時対応で助かったよ。

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[NoCoin]: https://github.com/keraf/NoCoin/ "keraf/NoCoin: No Coin is a tiny browser extension aiming to block coin miners such as Coinhive."
