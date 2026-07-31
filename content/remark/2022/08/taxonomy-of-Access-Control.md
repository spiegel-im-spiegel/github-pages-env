+++
title = "アクセス制御状態の4分類"
date =  "2022-08-16T12:41:39+09:00"
description = "Bruce Schneier 先生の記事より"
image = "/images/attention/kitten.jpg"
tags = [ "security", "risk", "management" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Bruce Schneier](https://www.schneier.com/) 先生の記事より：

- [A Taxonomy of Access Control - Schneier on Security](https://www.schneier.com/blog/archives/2022/08/a-taxonomy-of-access-control.html)

元ネタは Cryptocurrency Wallet (何となく嫌なので日本語で書いてやらんw) に関する以下の論文：

- {{< pdf-file title="On Cryptocurrency Wallet Design" link="https://eprint.iacr.org/2021/1522.pdf" >}}

Abstract には

{{< fig-quote type="markdown" title="On Cryptocurrency Wallet Design" link="https://eprint.iacr.org/2021/1522.pdf" lang="en" >}}
The security of individual keys was widely studied with practical solutions available, from mnemonic phrases to dedicated hardware. There are also techniques for securing funds by requiring combinations of multiple keys. However, to the best of our knowledge, a crucial question was never addressed: How is wallet security affected by the number of keys, their types, and how they are combined? This is the focus of this work.

We present a model where each key has certain probabilities for being *safe*, *lost*, *leaked*, or *stolen* (available only to an attacker). The number of possible wallets for a given number of keys is the Dedekind number, prohibiting an exhaustive search with many keys. Nonetheless, we bound optimal-wallet failure probabilities with an evolutionary algorithm.
{{< /fig-quote >}}

とあるが， Bruce Schneier さんが[指摘](https://www.schneier.com/blog/archives/2022/08/a-taxonomy-of-access-control.html "A Taxonomy of Access Control - Schneier on Security")される通り，これはもっと一般的なアクセス制御に関わるアカウント情報や暗号鍵の状態についても応用可能となる。
すなわち：

{{< fig-quote class="nobox" title="On Cryptocurrency Wallet Design" link="https://eprint.iacr.org/2021/1522.pdf" lang="en" >}}
<table>
<tr><th>safe</th><td>Only the user has access,</td></tr>
<tr><th>loss</th><td>No one has access,</td></tr>
<tr><th>leak</th><td>Both the user and the adversary have access, or</td></tr>
<tr><th>theft</th><td>Only the adversary has access</td></tr>
</table>
{{< /fig-quote >}}

の4つの状態を定義することで

{{< fig-quote type="markdown" title="A Taxonomy of Access Control" link="https://www.schneier.com/blog/archives/2022/08/a-taxonomy-of-access-control.html" lang="en" >}}
Once you know these states, you can assign probabilities of transitioning from one state to another (someone hacks your account and locks you out, you forgot your own password, etc.) and then build optimal security and reliability to deal with it.
{{< /fig-quote >}}

というわけだ。
件の論文はこの4状態とその関係を定量的に評価している点で素晴らしく，他のセキュリティ評価にも使えるため “[It’s a truly elegant way of conceptualizing the problem](https://www.schneier.com/blog/archives/2022/08/a-taxonomy-of-access-control.html "A Taxonomy of Access Control - Schneier on Security")” と絶賛されているのだろう。


というわけで，覚え書きとして記しておく。

## 参考図書

{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
