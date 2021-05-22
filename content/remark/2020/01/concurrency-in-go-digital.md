+++
title = "『Go 言語による並行処理』はEブック版がオススメ"
date =  "2020-01-13T11:28:13+09:00"
description = "私も完全に失念していたのだが『Go 言語による並行処理』ってEブック版があるんだよね。"
image = "/images/attention/kitten.jpg"
tags = [ "book", "e-book", "golang", "concurrency" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Twitter で

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">並行（concurrent）と並列（parallel）の違いは『Go 言語による並行処理』できちんと説明されている。 <a href="https://twitter.com/hashtag/golang?src=hash&amp;ref_src=twsrc%5Etfw">#golang</a> で並行処理を書くなら『Go 言語による並行処理』は必読書なので、是非とも手にしていただきたいところである<br> <a href="https://t.co/wvUiI94qOh">https://t.co/wvUiI94qOh</a> <a href="https://t.co/zAkd9anE8u">https://t.co/zAkd9anE8u</a></p>&mdash; Der Spiegel im Spiegel (@spiegel_2007) <a href="https://twitter.com/spiegel_2007/status/1214391644616658944?ref_src=twsrc%5Etfw">January 7, 2020</a></blockquote>
{{< /fig-gen >}}

とか呟いたら微妙に反応があったみたいなので追加情報を記しておく[^cing1]。

[^cing1]: 「並行性と並列性の違い」については2.1章で3ページ分使って詳しく説明されている。ありがたや。

つか，私も完全に失念していたのだが『Go 言語による並行処理』ってEブック版があるんだよね。
[O'Reilly Japan Ebook Store](https://www.oreilly.co.jp/ebook/) で買える（PayPal 決済）。

- [O'Reilly Japan - Go言語による並行処理](https://www.oreilly.co.jp/books/9784873118468/)

この手の技術解説書を Kindle で買うのは基本的にオススメできない。
理由は大きく2つ：

1. コードをコピペして検証できない
2. 索引が使えない

ことである。
まぁ，1番目については，最近では GitHub リポジトリ等にサンプルコードが置かれてたりするのだが。
あと，最初から索引が貧弱な本は技術解説書としては論外。
読み物として面白ければ，それでもいいが。

O'Reilly のEブックは PDF, ePub, mobi の3つのフォーマットで提供されていて，しかも要らん DRM がかかっていない。
適切な PDF リーダ等を使えば語句検索が利用できるし， [Ubuntu] 環境であれば標準の [Evince](https://wiki.gnome.org/action/show/Apps/Evince "Apps/Evince - GNOME Wiki!") で充分足りる[^abd1]。

[^abd1]: [Adobe なんて要らん](https://pdfreaders.org/ "Get a Free Software PDF reader!")ですよ。偉い人にはそれが分からんのです（笑） いや，真面目な話，官庁が出す PDF を読むのに Adobe Reader をダウンロードさせるのマジで止めて欲しい。特定企業を優遇するとか！

というわけで，『Go 言語による並行処理』を買うならEブック版がオススメという話でした。

## ブックマーク

- [『Go 言語による並行処理』は Go 言語プログラマ必読書だろう]({{< ref "/remark/2018/11/concurrency-in-go.md" >}})

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4908686033" %}} <!-- Goならわかるシステムプログラミング -->
