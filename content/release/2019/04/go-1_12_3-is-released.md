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

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
