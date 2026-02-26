+++
title = "Mastodon/Bluesky で共有するためのボタンを追加・更新"
date =  "2026-02-26T15:32:46+09:00"
description = "そろそろ重い腰をあげようかな，ということで。"
image = "/images/attention/kitten.jpg"
tags = [ "site", "mastodon", "bluesky" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

そろそろ重い腰をあげようかな，ということで。

このブログ各記事ページの右上には共有のためのボタンがある（目立たないようにわざと小さくしているが）。
今まで Bluesky の共有ボタンを付けてなかったのだが，今回これを追加した。
ボタンに対応する URL の組み立ては以下の公式ドキュメントを参考にした。

- [Action Intent Links | Bluesky](https://docs.bsky.app/docs/advanced-guides/intent-links)

ついでに Mastodon の共有ボタンも更新した。
第三者サービスを経由してポストするのはそろそろやめようかなと思って（信用してないわけじゃないけど）。
参考にしたのは以下の記事。

- [サイトに「Mastodonで共有」ボタンを設置する。 - freefielder.jp](https://freefielder.jp/blog/2024/11/share-on-mastodon-button.html)

このページで公開している CSS および JavaScript コードをそのまま使わせてもらった。
権利の帰属先および利用ライセンスは[こちら]({{< rlnk "/mstdn_share/license.txt" >}})。

ありがたや {{% emoji "ペコン" %}}

なお [Pocket](https://getpocket.com/) への共有ボタンは外した。
私が使わなくなったので。
はてな への共有ボタンも付けない。
そもそも はてな のアカウント持ってないしな。
