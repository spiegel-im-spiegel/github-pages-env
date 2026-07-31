+++
title = "“Astronomy Picture of the Day” はどこで更新されている？"
date =  "2025-10-28T08:47:23+09:00"
description = "情報をお持ちの方がいらっしゃいましたら，ぜひ教えてください 🙇"
image = "/images/attention/kitten.jpg"
tags = [ "astronomy", "photography" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

2025年10月に入ってから NASA の “Astronomy Picture of the Day” (APOD) の更新が止まっていて，ページトップに以下の文言が書かれている。

{{< fig-quote type="markdown" title="Astronomy Picture of the Day" link="https://apod.nasa.gov/apod/" lang="en" >}}
Due to the lapse in federal government funding, NASA is not updating this website. We sincerely regret this inconvenience.
{{< /fig-quote >}}

`orz`

ただ Bluesky の TL とかを眺めていると APOD が更新されているように見える。

{{< fig-gen >}}
<blockquote class="bluesky-embed" data-bluesky-uri="at://did:plc:btzjjptgouva43ko4ip5pl67/app.bsky.feed.post/3m45x7bxzub2q" data-bluesky-cid="bafyreigkomfauu2vxbxtzfdrr5mxwnraamvoqcoalo6sc76d72cxlo2kfq" data-bluesky-embed-color-mode="system"><p lang="">🔭 Two Tails of Comet Lemmon

Image Credit: Massimo Penna

star.ucl.ac.uk/~apod/apod/a...<br><br><a href="https://bsky.app/profile/did:plc:btzjjptgouva43ko4ip5pl67/post/3m45x7bxzub2q?ref_src=embed">[image or embed]</a></p>&mdash; Astronomy Picture of the Day 🪐 (<a href="https://bsky.app/profile/did:plc:btzjjptgouva43ko4ip5pl67?ref_src=embed">@apod.shinyakato.dev</a>) <a href="https://bsky.app/profile/did:plc:btzjjptgouva43ko4ip5pl67/post/3m45x7bxzub2q?ref_src=embed">October 27, 2025 at 5:00 PM</a></blockquote><script async src="https://embed.bsky.app/static/embed.js" charset="utf-8"></script>
{{< /fig-gen >}}

上のポストは[ミラー](https://apod.nasa.gov/apod/lib/about_apod.html "About APOD")のひとつである [UCL] のサイトをソースにしているようだ。

- [`http://star.ucl.ac.uk/~apod/apod/`](http://star.ucl.ac.uk/~apod/apod/)

つまり APOD 自体は止まってないが NASA の本家サイトでの更新が止まってるって感じだろうか。
これらのミラーサイトはどこから情報を取得してるんだろう。
まぁ，いいや。
とりあえず APOD の更新を追いたい人は，こちらのサイトをチェックするとよいだろう。

私も Bluesky 上で APOD の更新を通知する非公式ボットを運用しているのだが

- [Astronomy Picture of the Day (unofficial bot) (@apodunofficial.bsky.social) — Bluesky](https://bsky.app/profile/did:plc:hbzmqswkx5pbg5fhr33pw4iw)

こちらは [NASA API] を使ってデータを取得しているため，本家 APOD と同じく，更新が止まってしまっている。
[NASA API] もどこかで代替運用されてないっスかねぇ。

情報をお持ちの方がいらっしゃいましたら，ぜひ教えてください {{% emoji "ペコン" %}}

## ブックマーク

- [NASA API を使って “Astronomy Picture of the Day” のデータを取得する]({{< ref "/remark/2023/02/api-for-astronomy-picture-of-the-day.md" >}})

[UCL]: https://www.ucl.ac.uk/ "Welcome to UCL | University College London"
[NASA API]: https://api.nasa.gov/ "NASA Open APIs"

## 参考

{{% review-paapi "4416723660" %}} <!-- 天文年鑑 2025年版 -->
{{% review-paapi "4627275110" %}} <!-- 天体物理学 -->
