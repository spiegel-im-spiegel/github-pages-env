+++
title = "Integrity が問われる時代"
date =  "2025-06-28T19:48:41+09:00"
description = "Bruce Schneier さんのエッセイより"
image = "/images/attention/kitten.jpg"
tags = [ "code", "security", "risk", "management", "artificial-intelligence" ]

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

セキュリティの3要件をご存知だろうか。
すなわち

- 機密性 (Confidentiality)
- 完全性 (Integrity)
- 可用性 (Availability)

である。
このうち機密性は言わずもがなで，可用性についても Web 2.0 以降から重要視されるようになった。
一方で完全性についてはこれまで機密性を担保するための（データ破壊や改竄を防ぐ）要件として考えられることが多かったように思う。

ここで我らが Bruce Schneier 先生が以下のエッセイを公開された。

- [The Age of Integrity - Schneier on Security](https://www.schneier.com/blog/archives/2025/06/the-age-of-integrity.html)

この記事ではまず，これまでの狭義の「完全性（Integrity）」を拡張する。

{{< fig-quote type="markdown" title="The Age of Integrity" link="https://www.schneier.com/blog/archives/2025/06/the-age-of-integrity.html" lang="en" >}}
More broadly, integrity refers to ensuring that data is correct and accurate from the point it is collected, through all the ways it is used, modified, transformed, and eventually deleted. Integrity-related incidents include malicious actions, but also inadvertent mistakes.
{{< /fig-quote >}}

つまり，利用からはじめ変更・変換そして削除に至るまで，データのライフサイクル全体で「整合性[^i1]（Integrity）」を担保していくわけだ。
こうした「整合性」の問題は，これからの AI 時代に特に重要になっていくという。

[^i1]: 広義の “Integrity” の場合は「完全性」と訳すより「整合性」のほうが日本語的にピッタリくる？ Google 翻訳や Kagi 翻訳では「整合性」と訳していた。

{{< fig-quote type="markdown" title="The Age of Integrity" link="https://www.schneier.com/blog/archives/2025/06/the-age-of-integrity.html" lang="en" >}}
Integrity has always been important, but as we start using massive amounts of data to both train and operate AI systems, data integrity will become more critical than ever.

Most of the attacks against AI systems are integrity attacks. Affixing small stickers on road signs to fool AI driving systems is an integrity violation. Prompt injection attacks are another integrity violation. In both cases, the AI model can’t distinguish between legitimate data and malicious input: visual in the first case, text instructions in the second. Even worse, the AI model can’t distinguish between legitimate data and malicious commands.

Any attacks that manipulate the training data, the model, the input, the output, or the feedback from the interaction back into the model is an integrity violation. If you’re building an AI system, integrity is your biggest security problem. And it’s one we’re going to need to think about, talk about, and figure out how to solve.
{{< /fig-quote >}}

さらに Bruce Schneier さんは “confidentiality” に対する “confidential” や “availability” に対する “available” のように “integrity” に対しても “integrous” という言葉を普及させてはどうかと提案している。
ただし辞書にはない単語らしいけど（笑）

{{< fig-quote type="markdown" title="The Age of Integrity" link="https://www.schneier.com/blog/archives/2025/06/the-age-of-integrity.html" lang="en" >}}
There are deep questions here, deep as the internet. Back in the 1960s, the internet was designed to answer a basic security question: Can we build an available network in a world of availability failures? More recently, we turned to the question of privacy: Can we build a confidential network in a world of confidentiality failures? I propose that the current version of this question needs to be this: Can we build an **integrous** network in a world of integrity failures?
{{< /fig-quote >}}

（強調は私がやりました）

AI だけでなく，近年 Web 3.0 などと言われる分散型・非中央集権型・インテリジェンスなネットワークにおいても「整合性」の問題は重要な課題と言える。

{{< fig-quote type="markdown" title="The Age of Integrity" link="https://www.schneier.com/blog/archives/2025/06/the-age-of-integrity.html" lang="en" >}}
We need research into integrous system design.
{{< /fig-quote >}}

こうした議論についても今後注目していく必要があるかな。

## ブックマーク

- [IEEE Security & Privacy: The Age of Integrity](https://www.computer.org/csdl/magazine/sp/2025/03/11038984/27COaJtjDOM)

- [CC Signals と AI]({{< relref "./cc-signals.md" >}})

## 参考

{{% review-paapi "B0CK19L1HC" %}} <!-- ハッキング思考 Kindle 版 -->
