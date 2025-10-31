+++
title = "ハロウィンでプレミアムな Friday Night"
date =  "2025-10-31T09:33:16+09:00"
description = "お祭りな金曜日をお過ごしください。"
image = "/images/attention/kitten.jpg"
tags = [ "calendar" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

いやぁ，昭和生まれなもんでハロウィンとかあまり馴染みがないのですよ。
Wikipedia を見ると「[各地のイベントでハロウィーンにちなんだ仮装が導入されるようになったのは1980年代のバブル期以降](https://ja.wikipedia.org/wiki/%E3%83%8F%E3%83%AD%E3%82%A6%E3%82%A3%E3%83%B3 "ハロウィン - Wikipedia")」とか書いてあるけど，その頃にはお菓子をあげる側だったからなぁ。

ところで Mastodon の TL を眺めてたら

{{< fig-gen >}}
<blockquote class="mastodon-embed" data-embed-url="https://366.koyomi.online/@matsbox/115463736847543793/embed" style="background: #FCF8FF; border-radius: 8px; border: 1px solid #C9C4DA; margin: 0; max-width: 540px; min-width: 270px; overflow: hidden; padding: 0;"> <a href="https://366.koyomi.online/@matsbox/115463736847543793" target="_blank" style="align-items: center; color: #1C1A25; display: flex; flex-direction: column; font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Oxygen, Ubuntu, Cantarell, 'Fira Sans', 'Droid Sans', 'Helvetica Neue', Roboto, sans-serif; font-size: 14px; justify-content: center; letter-spacing: 0.25px; line-height: 20px; padding: 24px; text-decoration: none;"> <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="32" height="32" viewBox="0 0 79 75"><path d="M63 45.3v-20c0-4.1-1-7.3-3.2-9.7-2.1-2.4-5-3.7-8.5-3.7-4.1 0-7.2 1.6-9.3 4.7l-2 3.3-2-3.3c-2-3.1-5.1-4.7-9.2-4.7-3.5 0-6.4 1.3-8.6 3.7-2.1 2.4-3.1 5.6-3.1 9.7v20h8V25.9c0-4.1 1.7-6.2 5.2-6.2 3.8 0 5.8 2.5 5.8 7.4V37.7H44V27.1c0-4.9 1.9-7.4 5.8-7.4 3.5 0 5.2 2.1 5.2 6.2V45.3h8ZM74.7 16.6c.6 6 .1 15.7.1 17.3 0 .5-.1 4.8-.1 5.3-.7 11.5-8 16-15.6 17.5-.1 0-.2 0-.3 0-4.9 1-10 1.2-14.9 1.4-1.2 0-2.4 0-3.6 0-4.8 0-9.7-.6-14.4-1.7-.1 0-.1 0-.1 0s-.1 0-.1 0 0 .1 0 .1 0 0 0 0c.1 1.6.4 3.1 1 4.5.6 1.7 2.9 5.7 11.4 5.7 5 0 9.9-.6 14.8-1.7 0 0 0 0 0 0 .1 0 .1 0 .1 0 0 .1 0 .1 0 .1.1 0 .1 0 .1.1v5.6s0 .1-.1.1c0 0 0 0 0 .1-1.6 1.1-3.7 1.7-5.6 2.3-.8.3-1.6.5-2.4.7-7.5 1.7-15.4 1.3-22.7-1.2-6.8-2.4-13.8-8.2-15.5-15.2-.9-3.8-1.6-7.6-1.9-11.5-.6-5.8-.6-11.7-.8-17.5C3.9 24.5 4 20 4.9 16 6.7 7.9 14.1 2.2 22.3 1c1.4-.2 4.1-1 16.5-1h.1C51.4 0 56.7.8 58.1 1c8.4 1.2 15.5 7.5 16.6 15.6Z" fill="currentColor"/></svg> <div style="color: #787588; margin-top: 16px;">Post by @matsbox@366.koyomi.online</div> <div style="font-weight: 500;">View on Mastodon</div> </a> </blockquote> <script data-allowed-prefixes="https://366.koyomi.online/" async src="https://366.koyomi.online/embed.js"></script>
{{< /fig-gen >}}

とかあって「そういえば今日ってプレミアムフライデーだっけ」と思い出した。
[前にも書いた]({{< ref "/remark/2019/04/hello-ubuntu.md" >}} "Windows とともに平成は去り Ubuntu とともに令和を迎える")が「二千円札」と「プレミアムフライデー」は平成における二大駄政策だと思う。
当時でもプレミアムフライデーに律儀に15時で仕事上がりする人なんかおったんか？ って感じだったし，昨年（2024年）の[新札](https://www.npb.go.jp/ja/n_banknote/ "新しい日本銀行券特設サイト")切り替えでは二千円札なくなっちゃったもんなぁ。

ふと思いついて，次にハロウィン（10月31日）とプレミアムフライデーが重なるのはいつか調べてみた。
こんな感じ。

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    for y := 2026; ; y++ {
        tm := time.Date(y, time.October, 31, 0, 0, 0, 0, time.UTC) // Time zone differences don't affect this, so use UTC for now
        fmt.Printf("%s (%s)\n", tm.Format(time.DateOnly), tm.Weekday())
        if tm.Weekday() == time.Friday {
            break
        }
    }
}
```

これを[実行](https://go.dev/play/p/VWzl7Coi0c8)すると

```text
2026-10-31 (Saturday)
2027-10-31 (Sunday)
2028-10-31 (Tuesday)
2029-10-31 (Wednesday)
2030-10-31 (Thursday)
2031-10-31 (Friday)
```

となり，次は6年後の2031年らしい。
閏年にもよるが大体5,6年ごとに巡ってくる感じ。

それでは皆さま，お祭りな金曜日をお過ごしください。
私は多分[ドラクエI&II](https://store.steampowered.com/app/2893570/III/ "Steam：ドラゴンクエストI＆II")をやってると思う。

## ブックマーク

- [貴重資料展示室061 国の重要文化財『星学手簡』 - 国立天文台暦計算室](https://eco.mtk.nao.ac.jp/koyomi/exhibition/061/)
  - [漫画を担当：国の重要文化財『星学手簡』 解説漫画（国立天文台） | 松箱-matsbox-](https://mats-box.com/2023/10/26/20231026/)
- [拝啓、暦の上から -koyomi.online- | 漫画家・松浦はこによる、暦の擬人化サイトです](https://koyomi.online/)

- [「プレミアムフライデー」を求めるパッケージを作ってみた]({{< ref "/golang/premium-friday.md" >}})

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
