+++
title = "「LLM は詭弁家である」"
date =  "2025-11-20T18:37:58+09:00"
description = "LLMは考えない; 確率に基づいて行動する"
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "engineering", "politics" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

## 「LLM は詭弁家である」

自ら LLM (生成 AI) による[アシスタント][Assistant]や[翻訳サービス][Translate]を提供している “[Kagi Search]” ([いつもお世話になっています]({{< ref "/remark/2025/04/kagi-assistant-for-all-users.md" >}} "Kagi Assistant が全ユーザに解放")) のブログで面白い記事が上がっていた。

- [LLMs are bullshitters. But that doesn't mean they're not useful | Kagi Blog](https://blog.kagi.com/llms)

記事ではまず「嘘（Lying）」と「デタラメ（Bullshitting）」の違いを提示し

{{< fig-quote type="markdown" title="LLMs are bullshitters. But that doesn't mean they're not useful" link="https://blog.kagi.com/llms" lang="en" >}}
- **Lying** means you have a concept of what is true, and you’re choosing to misrepresent it.
- **Bullshitting** means you’re attempting to persuade without caring for what the truth is.
{{< /fig-quote >}}

その上で LLM は嘘を吐いているのではなく（アルゴリズムに従って）デタラメを言っているに過ぎないと指摘する。

{{< fig-quote type="markdown" title="LLMs are bullshitters. But that doesn't mean they're not useful" link="https://blog.kagi.com/llms" lang="en" >}}
LLMs predict text. That’s it.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="LLMs are bullshitters. But that doesn't mean they're not useful" link="https://blog.kagi.com/llms" lang="en" >}}
LLMs don’t think; they act in probabilities
{{< /fig-quote >}}

そういえば，前に『[そろそろ、人工知能の真実を話そう]』を[読んだ]({{< ref "/remark/2017/09/the-myth-of-the-singularity.md" >}} "『シンギュラリティの神話』を読む")ときにこの単語が出てきたな。

{{< fig-quote type="markdown" title="『シンギュラリティの神話』を読む" link="/remark/2017/09/the-myth-of-the-singularity/" >}}
{{< fig-quote  title="そろそろ、人工知能の真実を話そう" link="https://www.amazon.co.jp/dp/B071FHBGW8?tag=baldandersinf-22&LINKCODE=OGI&TH=1&PSC=1" >}}
<q>近頃は、われわれを翻弄するような科学的な言い回しが非常に多く、未来に対して異なるアプローチをとる三つの概念がよく混同されている。すなわち、可能性（possibilité）、蓋然性（probabilité）、信憑性（plausibilité）―― この三つが区別されていないために混乱が生じているのである。</q>
{{< /fig-quote >}}

信憑性についてはもう少し説明が必要だろう。

plausibilité (plausibility) の意味を Google 先生に訊くと「{{< ruby "ゆうど" >}}尤度{{< /ruby >}}」と答えが返ってきた。
尤度とは統計学の用語らしい。

- [【統計学】尤度って何？をグラフィカルに説明してみる。 - Qiita](http://qiita.com/kenmatsu4/items/b28d1b3b3d291d0cc698)

ただし，この本では「信憑性（plausibilité）」をそのような意味では使ってなくて，語源に近いニュアンスで

{{< fig-quote  title="そろそろ、人工知能の真実を話そう" link="https://www.amazon.co.jp/dp/B071FHBGW8?tag=baldandersinf-22&LINKCODE=OGI&TH=1&PSC=1" >}}
<q>つまりそれは一般受けが良く、多くの人が起こると思っているということだ。だが、実際には可能性も蓋然性も保証されていない。</q>
{{< /fig-quote >}}

という感じで使っている。

[そろそろ、人工知能の真実を話そう]: https://www.amazon.co.jp/dp/B071FHBGW8?tag=baldandersinf-22&LINKCODE=OGI&TH=1&PSC=1 "そろそろ、人工知能の真実を話そう (早川書房) | ジャン＝ガブリエル ガナシア, 小林 重裕・他, 伊藤 直子 | コンピュータサイエンス | Kindleストア | Amazon"
{{< /fig-quote >}}

とするなら LLM は “probability” というより “plausibility” に基づいて行動している，と言えるかも知れない。
だから「[ChatGPT はイエスマンだ](https://www.washingtonpost.com/technology/2025/11/12/how-people-use-chatgpt-data/ "How people use ChatGPT, according to 47,000 of its conversations - The Washington Post")」などと言われてるんだろうけど。

この辺の話を聞くと，どうしても落語の「[転失気](https://ja.wikipedia.org/wiki/%E8%BB%A2%E5%A4%B1%E6%B0%97 "転失気 - Wikipedia")」を連想してしまう。

{{< fig-youtube id="PV0cfrurcN4" title="三遊亭圓生・転失気（1968年） - YouTube" >}}

あまりネタバレを言うのもアレだが，知ったかぶりのデタラメを言う「大人」と，それを見抜いてやり込める「子供」の対比が面白いのよ。
その「大人」と LLM を見比べてしまうんだよな。

だから

{{< fig-quote type="markdown" title="LLMs are bullshitters. But that doesn't mean they're not useful" link="https://blog.kagi.com/llms" lang="en" >}}
LLMs are Sophists
{{< /fig-quote >}}

というフレーズにも納得してしまうのだ。
そして件の記事では “and sophists are useful” と続ける。

{{< fig-quote type="markdown" title="LLMs are bullshitters. But that doesn't mean they're not useful" link="https://blog.kagi.com/llms" lang="en" >}}
If I use a LLM to help me find a certain page in a document, or sanity check this post while writing it, I don’t care “why” the LLM did it. I just care that it found that page or caught obvious mistakes in my writing faster than I could have.

[...]

But remember that LLMs are bullshitters: you can use LLMs to get incredible gains in how fast you can do tasks like research, writing code, etc. assuming that you are doing it mindfully with the pitfalls in mind

By all means, use LLMs where they are useful tools: tasks where you can verify the output, where speed matters more than perfection, where the stakes of being wrong are low.
{{< /fig-quote >}}

最後の活用例で挙がっている「出力を検証できる場合，完璧さよりも速さが重要な場合，間違えてもリスクが低い場合」を見ると LLM がどういう性質のものか分かるというもの。
少なくとも死を覚悟するほどの人生の悩みを [LLM に相談](https://gigazine.net/news/20251108-seven-families-suing-openai-chatgpt-suicides/ "「ChatGPTが自殺や悪い妄想をかきたてた」として複数の家族がOpenAIを訴える - GIGAZINE")するのは間違ってると思うよ。

## 「AIと有権者エンゲージメント」

LLM の利用については，もうひとつ面白い記事がある。

- [AI and Voter Engagement - Schneier on Security](https://www.schneier.com/blog/archives/2025/11/ai-and-voter-engagement.html)

このエッセイは Bruce Schneier 氏と Nathan E. Sanders 氏との共著で “The Fulcrum” に[掲載](https://thefulcrum.us/media-technology/artificial-intelligence-in-politics "Who Will Be the First American Candidate To Harness AI - The Fulcrum")されたものらしい。
内容について簡単に言うと， SNS の繋がりを利用して政治的連帯（というか動員？）を呼びかける政治手法から AI を介した Relational Organizing に変わりつつある，ということのようだ。

{{< fig-quote type="markdown" title="AI and Voter Engagement" link="https://www.schneier.com/blog/archives/2025/11/ai-and-voter-engagement.html" lang="en" >}}
So if a campaign hits you at the right time with the right message, they might persuade you to task your AI assistant to ask your friends to donate or volunteer. The result can be something more than a form letter; it could be automatically drafted based on the entirety of your email or text correspondence with that friend. It could include references to your discussions of recent events, or past campaigns, or shared personal experiences. It could sound as authentic as if you’d written it from the heart, but scaled to everyone in your address book.

[Research](https://www.pnas.org/doi/10.1073/pnas.2412815122) suggests that AI can generate and perform written political messaging about as well as humans. AI will surely play a [tactical role](https://prospect.org/power/2025-10-10-ai-artificial-intelligence-campaigns-midterms/) in the 2026 midterm campaigns, and some candidates may even use it for relational organizing in this way.
{{< /fig-quote >}}

この具体例として日本の「チームみらい」が挙げられているのが面白かった。
田舎に住んでる身としては国内のことなのに彼岸の話なんだけど（笑）

{{< fig-quote type="markdown" title="AI and Voter Engagement" link="https://www.schneier.com/blog/archives/2025/11/ai-and-voter-engagement.html" lang="en" >}}
Anno was RECENTLY [elected](https://mainichi.jp/english/articles/20250720/p2a/00m/0na/011000c) to the upper house of the federal legislature as the founder of a new party with a [100 day plan](https://note.com/annotakahiro24/n/nd648962bd411) to bring his vision of a “public listening AI” to the whole country. In the early stages of that plan, they’ve invested their share of Japan’s 32 billion yen in [party grants](https://www.nippon.com/en/japan-data/h02362/)—public subsidies for political parties—to hire engineers building digital civic infrastructure for Japan. They’ve already created platforms to provide [transparency](https://marumie.team-mir.ai/o/team-mirai) for party expenditures, and to use AI to make [legislation](https://gikai.team-mir.ai/) in the Diet easy, and are meeting with engineers from US-based Jigsaw Labs (a Google company) to [learn from international examples](https://note.com/team_mirai_jp/n/n0bbbcc21c752) of how AI can be used to power participatory democracy.

Team Mirai has yet to prove that it can get a second member elected to the Japanese Diet, let alone to win substantial power, but they’re innovating and demonstrating new ways of using AI to give people a way to participate in politics that we believe is likely to spread.
{{< /fig-quote >}}

2026年は米国では中間選挙の年であると同時に建国250周年なんだそうで。
これに合わせ「AI による組織化（Organizing with AI）」についても色々と動きがあるだろうと締めている。

## ブックマーク

- [The Dangers of AI-Generated Romance | Psychology Today](https://www.psychologytoday.com/us/blog/its-not-just-in-your-head/202408/the-dangers-of-ai-generated-romance)
- [Rewiring Democracy - Schneier on Security](https://www.schneier.com/books/rewiring-democracy/)

[Kagi Search]: https://kagi.com/ "Kagi Search"
[Assistant]: https://kagi.com/assistant "The Assistant"
[Translate]: https://translate.kagi.com/ "Kagi Translate"
[そろそろ、人工知能の真実を話そう]: https://www.amazon.co.jp/dp/B071FHBGW8?tag=baldandersinf-22&LINKCODE=OGI&TH=1&PSC=1 "そろそろ、人工知能の真実を話そう (早川書房) | ジャン＝ガブリエル ガナシア, 小林 重裕・他, 伊藤 直子 | コンピュータサイエンス | Kindleストア | Amazon"

## 参考

{{% review-paapi "B071FHBGW8" %}} <!-- 人工知能の真実を話そう -->
