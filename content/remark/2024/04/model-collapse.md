+++
title = "自家中毒を起こす AI と SF に取り憑かれたビッグテック"
date =  "2024-04-15T22:12:54+09:00"
description = "botshit とモデル崩壊"
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "code", "engineering" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

久しぶりに [P2Ptk.org](https://p2ptk.org/ "P2Pとかその辺のお話R | Sharing is Caring") の翻訳記事を見た気がする。

- [食糞AIがもたらす危機：botshitの肥溜めと化すインターネットの未来 | p2ptk[.]org](https://p2ptk.org/monopoly/4484)

またコリイ・ドクトロウさんってばお下品な言葉を（笑）と思ったが “botshit” というのは2024年1月のこの論文が元ネタらしい。

- [Beware of Botshit: How to Manage the Epistemic Risks of Generative Chatbots by Tim Hannigan, Ian P. McCarthy, Andre Spicer :: SSRN](https://papers.ssrn.com/sol3/papers.cfm?abstract_id=4678265)

Abstract によると

{{< fig-quote type="markdown" title="Beware of Botshit: How to Manage the Epistemic Risks of Generative Chatbots" link="https://papers.ssrn.com/sol3/papers.cfm?abstract_id=4678265" lang="en" >}}
Generative chatbots do this work by ‘predicting’ responses rather than ‘knowing’ the meaning of their responses. This means chatbots can produce coherent sounding but inaccurate or fabricated content, referred to as ‘hallucinations’. When humans use this untruthful content for tasks, it becomes what we call ‘botshit’.
{{< /fig-quote >}}

なんだそうだ。
そしてコリイ・ドクトロウさんによると「[すでにインターネットはbotshitの肥溜めと化している](https://p2ptk.org/monopoly/4484 "食糞AIがもたらす危機：botshitの肥溜めと化すインターネットの未来 | p2ptk[.]org")（[internet is already drowning in it](https://pluralistic.net/2024/03/14/inhuman-centipede/#enshittibottification "Pluralistic: The Coprophagic AI crisis (14 Mar 2024) – Pluralistic: Daily links from Cory Doctorow")）」ということらしい。
つか，訳文のほうがお下品だな（笑）

これによって起こりうるのがモデル崩壊（Model Collapse）である。
以下は2023年5月の論文。

- [[2305.17493] The Curse of Recursion: Training on Generated Data Makes Models Forget](https://arxiv.org/abs/2305.17493)

{{< fig-quote type="markdown" title="The Curse of Recursion: Training on Generated Data Makes Models Forget" link="https://arxiv.org/abs/2305.17493" lang="en" >}}
We find that use of model-generated content in training causes irreversible defects in the resulting models, where tails of the original content distribution disappear. We refer to this effect as Model Collapse and show that it can occur in Variational Autoencoders, Gaussian Mixture Models and LLMs.  We build theoretical intuition behind the phenomenon and portray its ubiquity amongst all learned generative models. We demonstrate that it has to be taken seriously if we are to sustain the benefits of training from large-scale data scraped from the web.
{{< /fig-quote >}}

なんちうか，これって自家中毒みたいだよな（喩えが雑）。
こっちの記事のほうがイメージしやすいかな。

- [Will GPT models choke on their own exhaust? | Light Blue Touchpaper](https://www.lightbluetouchpaper.org/2023/06/06/will-gpt-models-choke-on-their-own-exhaust/)

{{< fig-quote type="markdown" title="Will GPT models choke on their own exhaust?" link="https://www.lightbluetouchpaper.org/2023/06/06/will-gpt-models-choke-on-their-own-exhaust/" lang="en" >}}
Just as we’ve strewn the oceans with plastic trash and filled the atmosphere with carbon dioxide, so we’re about to fill the Internet with blah. This will make it harder to train newer models by scraping the web, giving an advantage to firms which already did that, or which control access to human interfaces at scale. Indeed, we already see AI startups [hammering the Internet Archive](https://blog.archive.org/2023/05/29/let-us-serve-you-but-dont-bring-us-down/) for training data.
{{< /fig-quote >}}

“AI startups [hammering the Internet Archive](https://blog.archive.org/2023/05/29/let-us-serve-you-but-dont-bring-us-down/) for training data” の部分を示す記事がこれだね。

- [Let us serve you, but don’t bring us down | Internet Archive Blogs](https://blog.archive.org/2023/05/29/let-us-serve-you-but-dont-bring-us-down/)

{{< fig-quote type="markdown" title="Let us serve you, but don’t bring us down" link="https://blog.archive.org/2023/05/29/let-us-serve-you-but-dont-bring-us-down/" lang="en" >}}
Tens of thousands of requests per second for our public domain OCR files were launched from 64 virtual hosts on amazon’s AWS services. (Even by web standards,10’s of thousands of requests per second is a lot.)
{{< /fig-quote >}}

これは酷い。
1年前にこんなことが起こってたのか。

モデル崩壊については，私も耳にしていたが，昨年前半は仕事が徐々に忙しくなっていて LLM への関心が薄れていたんだよな。

{{< fig-quote type="markdown" title="食糞AIがもたらす危機：botshitの肥溜めと化すインターネットの未来" link="https://p2ptk.org/monopoly/4484" >}}
皮肉なことに、AI企業自身がこの問題の火種を作っている。GoogleやMicrosoftによる「AI検索」の全面的な推進は、検索エンジンがウェブページへのリンクを返すのではなく、そのコンテンツを要約する未来を想定している。しかし、そうなれば誰がウェブを書くだろうか。あなたの書いたものを見つけられるのはAIのクローラーだけで、しかもそのAIはあなたの書いたものを自分のトレーニングの餌にするだけで、読者にあなたの書いたものを紹介する気は毛頭ない。AIが検索を支配すれば、オープンウェブはAIの工業的畜産場（CAFO）となり、検索クローラーはますます肥溜めからクソを吸い上げるようになるだろう。
{{< /fig-quote >}}

大昔，仕事とかでメールのやり取りに辟易してた頃， AI がエージェントとなってメール送受信を仲立ちしてくれれば，マナーと称した下らない時候の挨拶とか書かなくて済むし，分かりにくい文面も要約してくれるんじゃないか？ と夢想したものだが，人同士のやり取りを LLM が仲立ちして肩代わりするようになったらどんな恐ろしいことが起きるんだろうねぇ（笑）

ところで，この記事の最初の方に

{{< fig-quote type="markdown" title="食糞AIがもたらす危機：botshitの肥溜めと化すインターネットの未来" link="https://p2ptk.org/monopoly/4484" >}}
かつては、一部の作家や読者がSFを予言と勘違いしても、さして問題にはならなかった。SF＝予言という妄想に駆られた人々に、社会を間違った方向に再構築する力がなかったからだ。しかし、SFに取り憑かれたテック億万長者たちが、我先に「人類苦悩化システム（torment nexus）」の発明に乗り出すに至っては、SF作家たちは作り物のお話と予言とを明確に区別しなくてはならなくなった（例「サイバーパンクは警告であって提案ではない」）。
{{< /fig-quote >}}

とか書かれていて思わず笑ってしまったのだが，似たような話を yomoyomo さんが紹介していた。

- [SFが未来を方向づけるのか？ 当代の人気SF作家が答える - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20240415/science-fiction-shape-the-future)

この記事でも

{{< fig-quote type="markdown" title="SFが未来を方向づけるのか？ 当代の人気SF作家が答える" link="https://yamdas.hatenablog.com/entry/20240415/science-fiction-shape-the-future" >}}
コリイ・ドクトロウは、実際にものを作る人の多くが SF にインスピレーションを受けているのは否定できないと認めながらも、物語の寓意をそのまま受け取りすぎるのは、[「プラトンの洞窟」](https://ja.wikipedia.org/wiki/%E6%B4%9E%E7%AA%9F%E3%81%AE%E6%AF%94%E5%96%A9)を本当に探すようなものと語っている。
{{< /fig-quote >}}

と紹介している。
ほかにも

{{< fig-quote type="markdown" title="SFが未来を方向づけるのか？ 当代の人気SF作家が答える" link="https://yamdas.hatenablog.com/entry/20240415/science-fiction-shape-the-future" >}}
結局のところ、SF が我々が築く未来の青写真になるんですかね？ という最初の問いに、N・K・ジェミシンが、そんなわけない。他の文学ジャンルと変わらん。ちょっとしたことを一つ正しく言い当てて、他はすべてハズレなのが「先見の明がある」ヴィジョンと言えるか？ と答えていて受けた。
{{< /fig-quote >}}

と紹介されていて，これが一番面白かった。
「SFに取り憑かれたテック億万長者たち」のせいで SF 作家がわざわざこういうことを言わないといけない世の中になったのかねぇ。

## ブックマーク

- [学校教育とAI：生成AIの使用禁止とAI検出ツールは生徒たちを害する | p2ptk[.]org](https://p2ptk.org/ai/4486)
- [AI栄えてQAエンジニアが儲かる？ 話はそんな単純ではない - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20240415/quality-assurance-and-ai)
- [アーヴィンド・ナラヤナンらの「インチキAI」本が遂に出る - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20240415/ai-snake-oil)
- [GPT-4超え精度でスマホ上実行できるオンデバイス生成AI「Octopus v2」、Google「生成AIは大きければいいってものではない」など重要論文5本を解説（生成AIウィークリー） | テクノエッジ TechnoEdge](https://www.techno-edge.net/article/2024/04/08/3130.html)
- [AI黙示録は“簡単に”回避できる | p2ptk[.]org](https://p2ptk.org/ai/4500)
- [TESCREALふたたび：AGIが約束するユートピアはSF脳のディストピアなのか？ – WirelessWire News](https://wirelesswire.jp/2024/05/86454/)
