+++
title = "米政府は何故 Anthropic モデルを禁じたのか"
date =  "2026-06-17T16:15:19+09:00"
description = "やはり今の米国は，中国やロシアなどと並んでリスクの高い国になったということだろうか。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "security", "risk", "management", "politics", "law" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

先日 Anthropic 社の最新 LLM である Claude Fable 5 および Mythos 5 に対して輸出規制がかけられた。

- [Anthropic、「Fable 5」「Mythos 5」のアクセスを停止、米国政府による輸出管理指令／その他のモデルは利用可能](https://forest.watch.impress.co.jp/docs/news/2116968.html)

私はこのニュースを見て20世紀後半の暗号輸出規制を連想したのだが，どうもそんな話ではないらしい。
色々と情報が錯綜しているが TechCrunch の記事が分かりやすくまとめられている。

- [The US government's Anthropic models ban was never about an AI jailbreak | TechCrunch](https://techcrunch.com/2026/06/15/the-us-governments-anthropic-models-ban-was-never-about-an-ai-jailbreak/)

ホンマ，日本語版がなくなったのは残念だよ。

この記事によると Anthropic 社は今回の措置の理由をガードレールを回避する所謂「AI 脱獄（AI jailbreak）」に関するものだと推測しているが，本当の理由については明らかになっていない模様。
政府と Anthropic 社の間のやり取りは公開されていないそうだ。
というわけで，いろんな憶測が飛び交っている。

{{< fig-quote type="markdown" title="The US government's Anthropic models ban was never about an AI jailbreak" link="https://techcrunch.com/2026/06/15/the-us-governments-anthropic-models-ban-was-never-about-an-ai-jailbreak/" lang="en" >}}
The Trump administration hasn’t confirmed why it invoked its export control directive. Did the officials misread the report and freak out? Did Amazon CEO Andy Jassy [say something to senior government officials](https://techcrunch.com/2026/06/13/amazon-ceo-reportedly-raised-anthropic-model-concerns-before-government-crackdown/) that prompted the reaction, out of caution or spite? Was something lost in translation, or was this a way to pressure Anthropic, with whom the administration already [has a fractious relationship](https://techcrunch.com/2026/03/05/its-official-the-pentagon-has-labeled-anthropic-a-supply-chain-risk/)? It’s possible that the White House was unaware of the far-reaching consequences of the letter’s demand and officials are scrambling to undo the damage of their own making.
{{< /fig-quote >}}

最近 [Amazon のセキュリティ研究者](https://www.wsj.com/tech/ai/anthropic-halts-access-to-top-ai-models-after-u-s-ban-on-foreign-use-a4bca2cc?st=sAwkD3)が Fable 5 におけるガードレール回避を疑う論文（非公開）を書いたそうで Anthropic 社がこれについて著者に連絡してきたとの情報はあるらしい。
当然米国政府にもこの情報は伝わってるだろう。

仮に LLM の脆弱性が理由だとしても，そのために輸出管理命令を出すのは全く筋違いの話だ。
だから “the Trump administration’s directive appears retaliatory” とか書かれちゃうんだろうけど。
セキュリティ研究者や専門家は輸出管理命令を撤回するよう求めているそうな。

{{< fig-quote type="markdown" title="The US government's Anthropic models ban was never about an AI jailbreak" link="https://techcrunch.com/2026/06/15/the-us-governments-anthropic-models-ban-was-never-about-an-ai-jailbreak/" lang="en" >}}
Moussouris and dozens of other top security researchers and experts have since called on the Trump administration to [revoke the export control order](https://techcrunch.com/2026/06/15/cybersecurity-vets-protest-dangerous-us-government-ban-on-anthropics-most-powerful-models/), calling the move to pull advanced cybersecurity capabilities from network defenders in the U.S. as “dangerous.”
{{< /fig-quote >}}

理由が何であれ，これが国際社会に対する信用を貶めるものになることは間違いないだろう。

{{< fig-quote type="markdown" title="The US government's Anthropic models ban was never about an AI jailbreak" link="https://techcrunch.com/2026/06/15/the-us-governments-anthropic-models-ban-was-never-about-an-ai-jailbreak/" lang="en" >}}
Justin Hendrix, the [editor of Tech Policy Press](https://www.techpolicy.press/anthropics-mythos-recall-and-the-white-houses-missing-ai-safety-playbook/), said the Trump administration’s move is “likely to raise alarms in foreign capitals about the reliability of American AI for critical applications.” The message is that AI companies in the United States can’t be trusted to operate without interference from the U.S. government.

[...]

To quote Hendrix, “the climate is one of a cloud of suspicion that senior officials are picking favorites based on personal and political factors.” The aftermath is that the government has set a dangerous precedent about how much control it intends to wield over the release of American-made software.
{{< /fig-quote >}}

やはり今の米国は，中国やロシアなどと並んでリスクの高い国になったということだろうか。

余談だが，生成 AI に対するガードレール回避などの問題は，今までのセキュリティ脆弱性とは性質が異なる。
間違ったコードを改善すれば直るようなものではないからだ。

{{< fig-quote type="markdown" title="The US government's Anthropic models ban was never about an AI jailbreak" link="https://techcrunch.com/2026/06/15/the-us-governments-anthropic-models-ban-was-never-about-an-ai-jailbreak/" lang="en" >}}
“The behavior described in the paper cannot meaningfully be fixed, and any attempt would only weaken the model for defense,” said Moussouris, who criticized the export control directive as hasty, heavy-handed, and misguided.
{{< /fig-quote >}}

生成 AI の登場でセキュリティ・リスク管理も新しい段階に入ったのかもしれない。

## ブックマーク

- [「プロンプトウェア・キルチェーン」]({{< ref "/remark/2026/02/the-promptware-kill-chain.md" >}})
- [欧州は脱米国に突き進むのか]({{< relref "./07-stories.md#breakup" >}} "最近の話題（2026-06-07）")

[Kagi Search]: https://kagi.com/ "Kagi Search - A Premium Search Engine"
[Kagi Translate]: https://translate.kagi.com/ "Kagi Translate"
[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"
[GitHub Copilot]: https://github.com/features/copilot "GitHub Copilot · Your AI pair programmer · GitHub"
<!-- eof -->
