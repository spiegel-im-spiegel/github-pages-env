+++
title = "デスクトップ版ぞーぺんに期待！"
date =  "2026-03-08T18:48:24+09:00"
description = "ぞーぺんに慣れると公式 Web アプリには戻れないのよ。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "tools", "Windows", "linux", "mastodon", "bluesky", "zonepane" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

昨日の “[hololive 7th fes. Ridin' on Dreams STAGE3](https://www.youtube.com/watch?v=L3-VtHNm9oA "【チラ見せ】hololive 7th fes. Ridin' on Dreams STAGE3【#hololivefesEXPO26_DAY2】 - YouTube")” よかったねぇ。
FLOW GLOW 初 fes. っスよ。
ライブを観ながら Mastodon の TL をチェックしてたんだけど

{{< fig-gen >}}
<iframe src="https://fedibird.com/@takke/116181260919599160/embed" width="400" allowfullscreen="allowfullscreen" sandbox="allow-scripts allow-same-origin allow-popups allow-popups-to-escape-sandbox allow-forms"></iframe>
{{< /fig-gen >}}

なんと！ そういえば TL で何やら作業をポストされてたな。
とりあえず Windows 版を[ゲット](https://zonepane.com/zonepane_beta.html "ZonePane Beta")して[ミニ PC]({{< ref "/remark/2025/01/win11pro-on-minipc.md" >}} "Mini PC を衝動買いした") にインストールしクロスポストを試してみた。

{{< fig-gen >}}
<blockquote class="mastodon-embed" data-embed-url="https://goark.fedicity.net/@spiegel/116181732370373076/embed" style="background: #FCF8FF; border-radius: 8px; border: 1px solid #C9C4DA; margin: 0; max-width: 540px; min-width: 270px; overflow: hidden; padding: 0;"> <a href="https://goark.fedicity.net/@spiegel/116181732370373076" target="_blank" style="align-items: center; color: #1C1A25; display: flex; flex-direction: column; font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Oxygen, Ubuntu, Cantarell, 'Fira Sans', 'Droid Sans', 'Helvetica Neue', Roboto, sans-serif; font-size: 14px; justify-content: center; letter-spacing: 0.25px; line-height: 20px; padding: 24px; text-decoration: none;"> <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="32" height="32" viewBox="0 0 79 75"><path d="M63 45.3v-20c0-4.1-1-7.3-3.2-9.7-2.1-2.4-5-3.7-8.5-3.7-4.1 0-7.2 1.6-9.3 4.7l-2 3.3-2-3.3c-2-3.1-5.1-4.7-9.2-4.7-3.5 0-6.4 1.3-8.6 3.7-2.1 2.4-3.1 5.6-3.1 9.7v20h8V25.9c0-4.1 1.7-6.2 5.2-6.2 3.8 0 5.8 2.5 5.8 7.4V37.7H44V27.1c0-4.9 1.9-7.4 5.8-7.4 3.5 0 5.2 2.1 5.2 6.2V45.3h8ZM74.7 16.6c.6 6 .1 15.7.1 17.3 0 .5-.1 4.8-.1 5.3-.7 11.5-8 16-15.6 17.5-.1 0-.2 0-.3 0-4.9 1-10 1.2-14.9 1.4-1.2 0-2.4 0-3.6 0-4.8 0-9.7-.6-14.4-1.7-.1 0-.1 0-.1 0s-.1 0-.1 0 0 .1 0 .1 0 0 0 0c.1 1.6.4 3.1 1 4.5.6 1.7 2.9 5.7 11.4 5.7 5 0 9.9-.6 14.8-1.7 0 0 0 0 0 0 .1 0 .1 0 .1 0 0 .1 0 .1 0 .1.1 0 .1 0 .1.1v5.6s0 .1-.1.1c0 0 0 0 0 .1-1.6 1.1-3.7 1.7-5.6 2.3-.8.3-1.6.5-2.4.7-7.5 1.7-15.4 1.3-22.7-1.2-6.8-2.4-13.8-8.2-15.5-15.2-.9-3.8-1.6-7.6-1.9-11.5-.6-5.8-.6-11.7-.8-17.5C3.9 24.5 4 20 4.9 16 6.7 7.9 14.1 2.2 22.3 1c1.4-.2 4.1-1 16.5-1h.1C51.4 0 56.7.8 58.1 1c8.4 1.2 15.5 7.5 16.6 15.6Z" fill="currentColor"/></svg> <div style="color: #787588; margin-top: 16px;">Post by @spiegel@goark.fedicity.net</div> <div style="font-weight: 500;">View on Mastodon</div> </a> </blockquote> <script data-allowed-prefixes="https://goark.fedicity.net/" async src="https://goark.fedicity.net/embed.js"></script>
{{< /fig-gen >}}

画像データの添付ができなかったものの，基本機能は概ね出来てる感じ。
こうなると Linux 版が欲しいよなぁ，と呟いたのだが Go で CLI ならともかくマルチプラットフォームは面倒だよねぇ。
Linux は需要少ないだろうし...

と思ってたら

{{< fig-gen >}}
<iframe src="https://fedibird.com/@takke/116188664136103534/embed" width="400" allowfullscreen="allowfullscreen" sandbox="allow-scripts allow-same-origin allow-popups allow-popups-to-escape-sandbox allow-forms"></iframe>
{{< /fig-gen >}}

まじすか！ ありがとうございます {{% emoji "ペコン" %}}

それじゃあ自宅の Ubuntu 機（Ubuntu Desktop 25.10）に入れて動かしてみるか。

{{< fig-img-quote src="./zonepane-1.png" title="ようこそ、ZonePaneへ！" link="./zonepane-1.png" width="1064" >}}

問題なく起動。
Mastodon と Bluesky のアカウントを追加してクロスポストを試してみる。

{{< fig-img-quote src="./cross-post-1.png" title="クロスポスト" link="./cross-post-1.png" width="1353" >}}

実は IM (Input Method) が上手く動かくて日本語が入力できなかった。
別のエディタで投稿内容を書いてコピペして対応。

{{< fig-img-quote src="./cross-post-3.png" title="投稿" link="./cross-post-3.png" width="1353" >}}

クロスポスト自体は問題なくできた。
ついでに投稿したものに対して返信と引用投稿を試してみた。

{{< fig-img-quote src="./reply-2.png" title="返信" link="./reply-2.png" width="1353" >}}

{{< fig-img-quote src="./quoted-post-2.png" title="引用投稿" link="./quoted-post-2.png" width="1353" >}}

同じように日本語が不自由な以外は問題なさそう。
結果はこんな感じ。

{{< fig-img-quote src="./zonepane-2.png" title="投稿結果" link="./zonepane-2.png" width="1353" >}}

概ね使えてる。
強いて気になるところを挙げると（上で挙げた点も含める）

- クロスポスト時に画像データの添付ができない（普通の投稿であれば問題ない）
- 入力フィールドで IM が効かず日本語が入力できない（Ubuntu 環境のみ）
- タイムラインの最上部に達すると更新マークが表示されるが実際には何も怒らない。アクションボタンの更新は効く
- スワイプが使えないのでタブがたくさんあると一部が隠れてしまう（タブ部分をマウスホイールで横方向にスクロールしないかなぁ）
- タイムラインがマルチカラムにならないかなぁ

といったところか。
まぁ iOS 版をベースにしているそうなので今のところはこんなものかと。
つか iOS 版をベースにここまでできるんだ。
Compose 凄いなぁ。

なによりクロスポストとタイムラインの使い勝手はぞーぺんならではで，これに慣れると公式 Web アプリには戻れないのよ。

- [ぞーぺん（ZonePane）によるクロスポスト]({{< ref "/remark/2025/06/multi-post-with-zonepane.md" >}})
- [ぞーぺんのユーザ体験]({{< ref "/remark/2026/02/zonepane.md" >}})

あとはマネタイズかな。
Windows 版と macOS 版はストアがあるので有料版の配布は難しくないんだろうけど Linux 版はねぇ。
[Open Collective](https://opencollective.com/ "Raise, manage and disburse money with full transparency. - Open Collective") みたいな仕組みを使うなら継続的に払うよ，私は（大金は難しいけど）。
ちなみに Android 版はサブスクリプション契約している。

まぁ，のんびり応援していこう。

## ブックマーク

- [ZonePane(ぞーぺん) (@zonepane@fedibird.com) - Fedibird](https://fedibird.com/@zonepane)
- [ZonePane iOS版をリリースしました！｜takke](https://note.com/panecraft/n/n2b5e96633032)
- [ZonePane for iOS が Mastodon に対応しました！｜takke](https://note.com/panecraft/n/n0bfcfffdcb62)
- [ZonePane for Bluesky&Mastodon - Apps on Google Play](https://play.google.com/store/apps/details?id=com.zonepane)
- [ZonePane – Multi SNS Clientアプリ - App Store](https://apps.apple.com/jp/app/zonepane-multi-sns-client/id6747976082)
