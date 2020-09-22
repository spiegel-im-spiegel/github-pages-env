+++
title = "ブログのネタをカンバン方式で管理する"
date =  "2020-09-22T19:32:40+09:00"
description = "これで塩漬け案件が無闇に増えるのを防ぐわけ。このまましばらく運用してみよう。"
image = "/images/attention/kitten.jpg"
tags = [ "site", "blog", "github", "management" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

ここのブログ記事は，大体は思い付きで書いているのだが，ネタとして溜め込んでいるものもそこそこある。
今まではテキストファイルでフラットに管理していたのだが「そういや [GitHub] って Projects 機能があるぢゃん♡」と気がついた。

そこで溜め込んでるネタを [Projects にいったん吐き出してみる][Projects]ことにした。
まずはネタを全部 [Issues] に書き出すところから。
書き出したらそれを [Projects] で整理していく。

[GitHub] の Projects 機能は，いわゆる「カンバン方式」で [Issues] を管理できる。
こんな感じ。

{{< fig-img src="./kanban.png" title="ブログのネタ帳" link="https://github.com/spiegel-im-spiegel/github-pages-env/projects/1" width="1477" >}}

とりあえず [Issues] に上げたものは全部 “Materials” に寄せておいて，その中で着手予定のものを “Issues” へ。
実際に着手を始めたものを “WIP (Work In Progress)” へ移動し，完了したら案件を Close (自動で “Done” に移動) する。

単純工程なこともあり WIP と Done はひと組しかないが，チームでやるシステム開発じゃないんだからこれで必要十分だろう。

ただし [Issues] も [Projects] も期限を管理できないので [Milestones] を設定する。
ここで言う期限はいわゆる「締切」ではない。
文字通りの一里塚。つまり，あるマイルストーンの期限が来たら終了してない案件を次のマイルストーンに回すか止めるか判断するわけだ。

これで塩漬け案件が無闇に増えるのを防げる（筈）。
まぁ，このまましばらく運用してみよう。

## ブックマーク

- [srggrs/assign-one-project-github-action: Automatically add an issue or pull request to specific GitHub Project(s) when you create and/or label them.](https://github.com/srggrs/assign-one-project-github-action)
- [Trello で引っ越し]({{< ref "/remark/2018/12/move-with-trello.md" >}})

[GitHub]: https://github.com/
[Issues]: https://github.com/spiegel-im-spiegel/github-pages-env/issues "Issues · spiegel-im-spiegel/github-pages-env"
[Milestones]: https://github.com/spiegel-im-spiegel/github-pages-env/milestones "Milestones - spiegel-im-spiegel/github-pages-env"
[Projects]: https://github.com/spiegel-im-spiegel/github-pages-env/projects/1 "ブログのネタ帳"

## 参考図書

{{% review-paapi "B01IGW5IIW" %}} <!-- リーン開発の現場 -->
