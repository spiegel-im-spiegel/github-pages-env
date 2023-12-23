+++
title = "The Social Butterfly"
date =  "2023-12-23T20:08:46+09:00"
description = "Bluesky で，ついにユーザのプロファイルおよび投稿がログインしなくても見れるようになった。"
image = "/images/attention/kitten.jpg"
tags = [ "bluesky", "communication" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今朝 Bluesky アプリのアイコンが蝶になってることに気がついた。
Web UI も含めてアイコンを変えたらしい。

- [A New Look for Bluesky: The Social Butterfly - Bluesky](https://blueskyweb.xyz/blog/12-21-2023-butterfly)

ブランドロゴとしてはどうなるか分からないが，とりあえず絵文字の 🦋 で代替できるっぽい。

他に重要な変更といえば，ユーザのプロファイルおよび投稿がログインしなくても見れるようになった。

{{< fig-quote type="markdown" title="A New Look for Bluesky: The Social Butterfly" link="https://blueskyweb.xyz/blog/12-21-2023-butterfly" lang="en" >}}
We built Bluesky to be a home for public conversation — breaking news, commentary and analysis, jokes and more. And we’re taking one step closer to this goal by releasing a public web view, which means that you don’t have to be logged in to view posts on Bluesky.

Starting today, you can easily view Bluesky posts without being logged in. Sharing will be more convenient — whether it's a joke you want to text a friend, or a post you want to embed in an article.
{{< /fig-quote >}}

不特定から見られたくない場合は Setting の Moderation にある “Discourage apps from showing my account to logged-out users” のスイッチをオンにすればよい（既定はオフ）。

{{< fig-img-quote src="./logged-out-visibility.png" title="Moderation — Bluesky" link="https://bsky.app/moderation" lang="en" width="574" >}}

ただし，この設定に従うかどうかはアプリ次第らしいので，公式のアプリや Web UI 以外では無視される可能性もある。
&#x1D54F; (旧 Twitter) で言うところの鍵垢とは違うので注意が必要である。

というわけで，フッタ部に暫定で Bluesky のアイコンを追加してみた。
そうそう。
招待状が余ってるので，私と直接連絡がつけれる方は言ってください。
融通できます（転売禁止）。
ちなみに [&#x1D54F; での活動は休眠中]({{< ref "/remark/2023/09/suspend-activity-on-twitter.md" >}} "𝕏 (旧 Twitter) の活動を休止します（期間未定）")で殆どアクセスしないのであしからず。

アイコンの蝶ように蛹から羽化して青空に上手く飛び立てればいいねぇ。
個人的には「分散型 SNS」というものに対して今ひとつ懐疑的ではあるが，少なくとも &#x1D54F; の代替としては期待している（Mastodon は &#x1D54F; の代替にはなれないと思う。もっと違う何か）。

## ブックマーク

- [Mastodon と Bluesky でボット運用はじめました]({{< ref "/remark/2023/07/crawler.md" >}})
