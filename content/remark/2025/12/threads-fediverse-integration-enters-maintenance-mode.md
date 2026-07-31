+++
title = "Threads の Fediverse 連携がメンテナンスモードに入った話"
date =  "2025-12-20T17:08:04+09:00"
description = "もう Threads に夢を見るのはやめよう。あそこは別次元の箱庭世界である。"
image = "/images/attention/kitten.jpg"
tags = [ "activitypub", "communication" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

以下の記事の前半で Threads の Fediverse 連携がメンテナンスモードに入ったことが報告されている。

- [Fediverse Report -#147 – Connected Places][件の記事]

{{< fig-quote type="markdown" title="Fediverse Report -#147" link="https://connectedplaces.online/reports/fediverse-report-147/" lang="en" >}}
Threads’ integration with the fediverse is on maintenance mode, says head of Threads Connor Hayes in an interview by Alex Heath. Hayes said about the fediverse: ““It’s something that we’re supporting, it’s something that we’re maintaining, but it’s not the thing that we’re talking about that’s gonna help the app break out”.
{{< /fig-quote >}}

私自身は Threads や Instagram のアカウントを持ってないし（[Instragram アカウントは抹消]({{< ref "/remark/2020/08/divorce-from-instagram.md" >}} "Instagram から足を洗う方法")した）， Mastodon からフォローしている Threads アカウントは [yomoyomo](https://www.threads.com/@yomoyomo/) さんと [gihyo.jp](https://www.threads.com/@gihyojp/) くらいなので，どういう状況になっているか分からないのだが，この記事を書いている時点では ActivityPub 機能が止まっているとかではないらしい。
でも将来的にはどうなるか分からんよ，という感じだろうか。

[上の記事][件の記事]によると Threads の月間アクティブユーザー数（MAU）は4億人らしい。
{{% emoji "X" %}} の MAU が5.7億人だそうなので，今や Threads のユーザ規模は {{% emoji "X" %}} に匹敵すると言っていいだろう[^mau1]。
これだけのユーザ規模にも関わらず Fediverse 連携を有効にしているユーザは25K人程度で[^f1]，さらに Threads からフォローしている mastodon.social サーバのユーザ数は800人程とのこと。

[^mau1]: ちなみに Instagram や Facebook の MAU は30億人， YouTube が20億人， TikTok が15億人らしい。桁が違うね（笑）
[^f1]: Threads では ActivityPub による Fediverse 連携はオプトインで提供されている。ちなみに25K人というのは mastodon.social サーバから見た数値で，実際にはもっと多い可能性がある。

もっとも Threads から他サーバのユーザへのフォローが極端に少ないのは Threads 側の意図的な UX 設計のようだ。

{{< fig-quote type="markdown" title="Fediverse Report -#147" link="https://connectedplaces.online/reports/fediverse-report-147/" lang="en" >}}
Threads deliberately made it difficult to follow fediverse accounts, and later updates deprioritised the fediverse feed even more.
{{< /fig-quote >}}

要するに Threads 側から見て分散化や Fediverse 連携は「客寄せパンダ」みたいなものであり Threads 内では「取るに足らないもの」もしくは「ノイズ」でしかないということなのだろう。

[件の記事]はもっと辛辣で，某マスク氏による Twitter/{{% emoji "X" %}} のドラスティックなメタクソ化（[enshittification](https://en.wikipedia.org/wiki/Enshittification "Enshittification - Wikipedia")）によるユーザの嫌気をうまく利用してサービスを展開し，名ばかりの分散化で規制当局を謀り， Fediverse 内部の対立を煽り，さらには Bluesky ユーザの一部も食ったと高評価（笑）である。

{{< fig-quote type="markdown" title="Fediverse Report -#147" link="https://connectedplaces.online/reports/fediverse-report-147/" lang="en" >}}
My take is that Meta and Threads have played the game well. They immediately capitalised on the moment in 2023 when decentralisation and Twitter-alternatives got large-scale attention, and knew how to say the right buzzwords to ride the wave. It got them in the right light for regulators, and gave them something tangible to point out to say ‘hey we’re doing interoperability now!’. The fediverse turned out to be highly vulnerable to such a strategy, a sitting duck for Big Tech companies to pluck some good PR from. That it turned the fediverse against itself, resulting in vicious and endless arguments about whether servers should block Threads, and whether Threads joining the fediverse validated the movement, was only a nice bonus. By slowly rolling out an implementation over the years Threads built their own positive-PR machine, every slight update worthy of a new article that put ‘Meta’ and ‘Interoperability’ in the headlines again. That nobody ever really used the integration between Threads and the fediverse never really seemed to matter, only the hypothetical future mattered. Nor did the press seem particularly interested in reporting on the fact that marginally few people seemed to be using the connection between the fediverse and Threads. Still, the company found itself a place at the table for protocol conversations about ActivityPub, which might pay dividend in the future if the need arises.

GGWP to Mark Zuckerberg and Adam Mosseri, for the skillful execution that helped Meta significantly, nestled themselves into the fediverse were it ever to become useful again, and taking some of the [wind out of Bluesky’s sails](https://bsky.app/profile/pfrazee.com/post/3macg4psbf22k) as well as a bonus. 
{{< /fig-quote >}}

ここに至ってユーザ規模が {{% emoji "X" %}} に追いついてきたのなら「Fediverse 連携要らんよね」という結論になってもおかしくない。
もう Threads に夢を見るのはやめよう。
あそこは別次元の箱庭世界である。
今フォローしてる Threads アカウントはどうしようかな。
gihyo.jp の Mastodon アカウントってないのかな？

今回の件を切っ掛けとしてソーシャルメディアの分散化についてもう一度考えるべきなんじゃないかね。
ほとんどのユーザにとって分散化は関心領域外だろうし，関心のある人でも本当に欲しいのは分散化じゃなくて「いざというときの避難場所」だろう。
Mastodon だってサーバによってルールや文化が違うし，本当に好き勝手にやりたいなら，私のように「おひとりさまサーバ」に引きこもるしかない。

ソーシャルメディアの分散化のお題目で連想するのは，かつてリベラルが掲げ失敗した「多文化主義」なんだよな。
やはり歴史は繰り返すのか？

[件の記事]: https://connectedplaces.online/reports/fediverse-report-147/ "Fediverse Report -#147 – Connected Places"

## ブックマーク

- [2025年12月版！性別・年齢別 SNSユーザー数（X、Instagram、TikTokなど13媒体）  | 株式会社ガイアックス](https://gaiax-socialmedialab.jp/socialmedia/435)

- [ソーシャルメディアは億万長者から逃げ切れるか？]({{< ref "/remark/2025/01/escape-from-billionaires.md" >}})
