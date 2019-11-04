+++
title = "Go 1.12.3 がリリースされた"
date = "2019-04-09T21:32:01+09:00"
description = "この前 1.12.2 が出たばかりなんだけどね。 Linux 環境の方は要アップデート。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "linux" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

この前 1.12.2 が出たばかりだが [Go] 1.12.3 リリースされた。
セキュリティ・アップデートはなし。

- [Go 1.12.3 and Go 1.11.8 are released - Google Group](https://groups.google.com/forum/#!topic/golang-announce/o8f9J4WQhKs)

{{< fig-quote title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release.html#go1.12.minor" lang="en" >}}
<q>go1.12.3 (released 2019/04/08) fixes an issue where using the prebuilt binary releases on older versions of GNU/Linux <a href="https://golang.org/issue/31293">led to failures</a> when linking programs that used cgo. Only Linux users who hit this issue need to update. </q>
{{< /fig-quote >}}

というわけで Linux 環境の方はアップデートしたほうがいいだろう。

アップデートは計画的に。

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN), 4621300253 (ISBN), 9784621300251 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
