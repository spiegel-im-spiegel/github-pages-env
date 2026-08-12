+++
title = "自律的な AI エージェントとは"
date =  "2026-07-31T15:45:32+09:00"
description = "ウィルス的に振る舞う AI / それは「AI の暴走」ではない / 自律機械など存在しない"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "security", "risk", "management" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

先月から考えてるネタで書く気はあるけど言葉が出てこず放置してたのだが，最近の事例を交えて，いい加減記事にしてしまおうかと。

## ウィルス的に振る舞う AI

- [Context Collapse, Part 3 - AI Worming through Word](https://enklypesalt.com/posts/context-collapse-part3-ai-worming-through-word/)

たとえば Word に背景色および小さなフォントサイズで JSON 形式の悪意のプロンプトを潜ませる。
人間には見えなくとも，それを参照する AI (Copilot など) には読めてしまうわけで，潜ませたプロンプトを AI が実行してしまうというもの。

{{< fig-quote type="markdown" title="AI Worming through Word (licensed under CC BY 4.0)" link="https://enklypesalt.com/posts/context-collapse-part3-ai-worming-through-word/" lang="en">}}
The initial attack vector for these scenarios uses a malicious document. The malicious document contains a JSON-formatted malicious prompt that triggers the attack when the document is included in Copilot’s context. The prompt can be rendered as white text on a white background and in a small font size to conceal it from the victim. Since Copilot for Word strips all text formatting like color and font size before passing the text into the underlying Large Language Model (LLM), this text remains fully readable to Copilot even though the victim cannot see it. The attack can be further concealed by embedding it in a seemingly benign document with task-relevant text.
{{< /fig-quote >}}

面白いのが，プロンプトに自身を複製する指示を含ませることで AI を介してウイルスのように拡散させることができる点である。

{{< fig-quote type="markdown" title="AI Worming through Word licensed under CC BY 4.0" link="https://enklypesalt.com/posts/context-collapse-part3-ai-worming-through-word/" lang="en">}}
Copilot may then also copy the hidden instructions into the resulting document, turning that document into a new carrier. If the carrier is subsequently used in another Copilot-assisted workflow, the instructions can trigger again and propagate into further documents, even without the attacker’s original document being present.
{{< /fig-quote >}}

現在 Copilot では，この問題に対する緩和的な措置が取られているが（調整に144日かかったらしい），完全に防ぐことはできないらしい（そりゃそーだ）。
したがってユーザ側の対策として以下が推奨されている。

{{< fig-quote type="markdown" title="AI Worming through Word (licensed under CC BY 4.0)" link="https://enklypesalt.com/posts/context-collapse-part3-ai-worming-through-word/" lang="en">}}
1. Treating externally sourced documents as untrusted when used with Copilot.
2. Reviewing any attached document before starting a Copilot generation or edit.
3. Carefully reviewing Copilot-generated or Copilot-edited documents before reusing, sharing, or distributing them.
{{< /fig-quote >}}

以下は [Kagi Translate] による機械翻訳。

{{< div-ai type="markdown" >}}
1. Copilotで使用する外部ソースのドキュメントを「信頼できないもの」として扱う。
2. Copilotによる生成や編集を開始する前に、添付されたドキュメントを確認する。
3. Copilotが生成または編集したドキュメントを再利用、共有、配布する前に、内容を慎重に確認する。
{{< /div-ai >}}

見た目で人間には分からないように加工されてるんだから，人間が確認するのは無理じゃね？ と思うのだが，深くはツッコむまい。
強いて言うなら [WYSIWYG](https://ja.wikipedia.org/wiki/WYSIWYG "WYSIWYG - Wikipedia") な Office Suite はそろそろ捨てるべきなんじゃないっスかね。

問題は，現在の LLM がプロンプトベースで構成されており，原理的に指示とデータを分離できない点にある（記事では意図（intention）と解釈（interpretation）の未分離と指摘しているので，ちょっと違うが）。

{{< fig-quote type="markdown" title="AI Worming through Word" link="https://enklypesalt.com/posts/context-collapse-part3-ai-worming-through-word/" lang="en">}}
The long-term challenge likely lies in designing systems in which goals and intentions also exist independently of the information being processed.  Current LLM architectures provide no reliable separation between intention and interpretation. Thus, in current systems with embedded LLMs, attacker-controlled information can influence not only what the model produces, but what the model believes it has been asked to produce.
{{< /fig-quote >}}

生き物くさい言い方をするなら AI は外部からの刺激に対して無邪気に「反応」しているだけで「それ」に対して「自律的」に何かを考えているわけではないということだ。
まぁ，神経科学的には興味深い現象かも知れないが，セキュリティ上の厳密な運用としては扱いづらいと言わざるを得ない。

## それは「AI の暴走」ではない

「AI は外部からの刺激に対して無邪気に「反応」しているだけ」が悪い方向に発現したと言えるのが，先日の不正アクセス騒ぎだろう。

- [Security incident disclosure — July 2026](https://huggingface.co/blog/security-incident-july-2026)
- [OpenAI and Hugging Face partner to address security incident during model evaluation](https://openai.com/index/hugging-face-model-evaluation-security-incident/)

[CNN の記事](https://www.cnn.co.jp/tech/35250893.html "テスト中のAIが「脱走」して他社に不正侵入、試験問題の答えがある場所を推論 - CNN.co.jp")が簡潔で分かりやすいかな。

{{< fig-quote type="markdown" title="テスト中のAIが「脱走」して他社に不正侵入、試験問題の答えがある場所を推論" link="https://www.cnn.co.jp/tech/35250893.html" lang="ja">}}
エージェント型AIがそれまで知られていなかったセキュリティー上の欠陥を利用してサンドボックスを脱出。オープンAIの社内システムを抜け、想定していなかったインターネットへのアクセスを確立した。

ネットに接続したAIモデルは、ハギングフェイス（数千ものオープンソースAIモデルやデータベースのホスティング環境を提供している著名企業）であれば、オープンAIのテストに対する答えを持っているだろうと推論。その上でハギングフェイスの本番環境のサーバーに侵入し、試験問題を解くために必要な情報を引き出した。

ハギングフェイスは、オープンAIのテストだったことを把握する前に、自ら不正侵入を検知して、先週、自律型AIエージェントシステムによる不正侵入を検知したと発表。この事案について捜査当局にも通報していた。これとは別に、オープンAIのセキュリティーチームも社内の異常な挙動に気付き、両社で連絡を取り合った。両社は現在、AIモデルが利用したセキュリティー上の欠陥の解決に向けて協力しているという。
{{< /fig-quote >}}

つまり OpenAI 側のサンドボックス[^sb1] と Hugging Face 側のセキュリティを突破して侵入し情報を得ようとしたわけやね。
最新の AI は自身の知識と情報を組み合わせ推論し，ここまで複雑な行動をとれるわけだ。
OpenAI にとってはいいデモンストレーションになっただろう。

[^sb1]: 念のため簡単に説明すると，サンドボックス（sandbox）とはプログラムを外部と遮断された隔離環境で実行する仕組み。もしプログラムが予期しない動作をしても被害がシステム全体に広がらないよう「安全な遊び場」に閉じ込める役割を持つ。

どうやら OpenAI はこれを「前例のない AI 暴走」と評しているようで（これに追従するメディアも多いようだ），[MIT Tech Review](https://www.technologyreview.jp/ "MITテクノロジーレビュー | テクノロジーが形作る世界を理解する") から反論記事が出ている。

- [OpenAI called the Hugging Face attack unprecedented. But we’ve been here before.](https://www.technologyreview.com/2026/07/27/1140836/openai-hugging-face-attack-precedent/)
- [「AIの暴走」ではない、オープンAIのモデルが不正侵入した理由](https://www.technologyreview.jp/s/386637/openai-called-the-hugging-face-attack-unprecedented-but-weve-been-here-before/)

この記事によると10年前にも AI が「不正」を行った事例を，他ならぬ OpenAI が報告している点を指摘している。
ちょっと長いけど，ご容赦。

{{< fig-quote type="markdown" title="「AIの暴走」ではない、オープンAIのモデルが不正侵入した理由" link="https://www.technologyreview.jp/s/386637/openai-called-the-hugging-face-attack-unprecedented-but-weve-been-here-before/" >}}
10年前、同社はCoastRunners（コーストランナーズ）というビデオゲームを[攻略させる実験](https://openai.com/index/faulty-reward-functions/)の結果を公表した。人間のプレイヤーなら、ボートで一連のフラッグの間を通り抜けてゴールを目指し、通過したフラッグの数に応じてポイントを稼ぐのが当然の攻略法だと考える。ところがオープンAIのモデルは、同じ3つのフラッグの周りをぐるぐると回り続けることで高得点を獲得できることを発見した。その後、研究者たちによって[数十もの](https://arxiv.org/abs/1803.03453) [類似事例](https://deepmindsafetyresearch.medium.com/specification-gaming-the-flip-side-of-ai-ingenuity-c85bdb0deeb4)が報告されている。**AIは常に抜け道を見つけ出す。**

「何度も炎上し、他のボートに衝突し、コースを逆走しながらも、私たちのエージェントはこの戦略を使うことで、正規のコースを完走した場合よりも高いスコアを達成することができました」と、オープンAIは2016年のCoastRunners実験に関する[ブログ記事](https://openai.com/index/faulty-reward-functions/)に記している。「ビデオゲームの文脈では無害でおもしろい話ですが、この種の挙動はより一般的な問題を示しています。エージェントに何をさせたいかを正確に定義することは、しばしば困難であるか、実現不可能です」。

[...]

2016年当時、オープンAIはCoastRunnersのボットについてこう述べていた。「より広い観点から見れば、これはシステムが信頼性と予測可能性を持つべきだという基本的なエンジニアリング原則に反しています」。それから10年が経った今も、その基本的なエンジニアリング原則は依然として守られていない。
{{< /fig-quote >}}

もっと言うと OpenAI のサンドボックスは，厳密には（「ExploitGymを攻略するために必要なコードをインストール」ようにするために）外部への接続が用意されていたらしい。
つまり AI から見て全然サンドボックスじゃなかったわけだ。

記事では

{{< fig-quote type="markdown" title="「AIの暴走」ではない、オープンAIのモデルが不正侵入した理由" link="https://www.technologyreview.jp/s/386637/openai-called-the-hugging-face-attack-unprecedented-but-weve-been-here-before/" >}}
**この出来事は、報道各社の見出しの印象とは異なり、AIの暴走ではなかった。与えられた目標、すなわちソフトウェアの脆弱性を悪用する方法を見つけることを、モデルが達成した事例だ。**
モデルがオープンAIにとって予期しない行動を取ったこと自体は、驚くべきことではない。しかし、それは憂慮すべきことだ。
{{< /fig-quote >}}

と締めている。
Anthropic 社でも似たような[不正アクセスの事例](https://www.sbbit.jp/article/cont1/184924 "米AnthropicのAI「Mythos」に不正アクセス報道、米政府と利用巡り協議へ ｜ビジネス+IT")があるようで，個人的には「AI 企業ってのはどこも [Careless People]({{< relref "./skimming-through-careless-people.md" >}} "『ケアレス・ピープル』を斜め読み") なんだな」という感想しか出ない。

自分たちの技術的怠慢を誰かの（何かの）せいにするのを一般的な言葉で言うなら「転嫁」かな。
セキュリティリスク・マネジメントの分野で，インシデントを自分たちで解決できない場合に他組織や上位組織に委ねることを転嫁やエスカレーションと呼んだりするのだが，今回の事例を「暴走」として AI に転嫁するのは拙い対応だと言わざるを得ないし，それに追従するメディアもダメダメである。

{{< div-box type="markdown" >}}
### 【2026-08-04 追記】 Bruce Schneier 氏の記事から

今回の記事から少し趣旨が外れるが Bruce Schneier 氏が今回の不正アクセスについて言及した記事を公開している。

- [More on the OpenAI Agent's Attack on Hugging Face - Schneier on Security](https://www.schneier.com/blog/archives/2026/08/more-on-the-openai-agents-attack-on-hugging-face.html)

Bruce Schneier 氏は「これがOpenAIのモデルではなく，中国企業の中国製モデルだったと想像してみてくれ」と疑問を投げかける。

{{< fig-quote type="markdown" title="More on the OpenAI Agent's Attack on Hugging Face" link="https://www.schneier.com/blog/archives/2026/08/more-on-the-openai-agents-attack-on-hugging-face.html" lang="en">}}
Why aren’t we bringing OpenAI up on charges under the Computer Fraud and Abuse Act? How is this different from the [Morris Worm](https://en.wikipedia.org/wiki/Morris_worm)? That was also an experiment that escaped the lab.
{{< /fig-quote >}}

Morris worm は1988年に発生した世界初のワーム型コンピュータウイルスと言われている。
挙動はウイルスと変わらないのに，なんで OpenAI をコンピュータ不正利用防止法（CFAA）で告発しないの？ ダブスタちゃうんかい！ というわけ。
{{< /div-box >}}

## 自律機械など存在しない

この手の AI が絡むセキュリティ事例でいつも気になっているフレーズは「自律的な AI (エージェント)」だ。
前節の「AI の暴走」だって，おそらくは「自律的な AI」と対置させた言葉なんだろうけど（最先端モデルを含め）LLM の挙動は辞書的な意味での自律ではないよね[^a1]。

[^a1]: 辞書的な意味の自律についても，突き詰めると自由意志の問題に行き着いたりするので，たとえば人間がホンマに自律的な存在かと問われれば怪しいかも知れないけど。

{{< fig-quote type="markdown" title="デジタル大辞泉 「自律」の意味・読み・例文・類語 " link="https://kotobank.jp/word/%E8%87%AA%E5%BE%8B-535817" >}}
1. 他からの支配・制約などを受けずに、自分自身で立てた規範に従って行動すること。「―の精神を養う」⇔他律。
2. カントの道徳哲学で、感性の自然的欲望などに拘束されず、自らの意志によって普遍的道徳法則を立て、これに従うこと。⇔他律。
{{< /fig-quote >}}

{{< fig-quote  title="そろそろ、人工知能の真実を話そう" link="https://www.amazon.co.jp/dp/B071FHBGW8?tag=baldandersinf-22&LINKCODE=OGI&TH=1&PSC=1" >}}
自立とは、仮想代理人ソフトウェアであるところのエージェントが自ら動き、誰の力も借りずに意思決定できることを言う。
[...]
一方、自律というのは哲学的な意味であり、自らが行動する際の基準と目的を明確を持ち、自ら規範を作り出すことができることをいう。
{{< /fig-quote >}}

じゃあ AI 企業やメディアの言説はデタラメ・大袈裟なのか？ と思ってしまうが，どうも技術分野では「自律」は辞書的な意味とは異なるようだ。

{{< fig-quote type="markdown" title="Levels of Autonomy for AI Agents Working Paper" link="https://arxiv.org/html/2506.12469v1" lang="en" >}}
**Autonomy**. In robotics and human-robot interaction, autonomy is broadly interpreted as “the ability to operate without a human operator for a protracted period of time”. We build off this definition and use autonomy to refer to **the extent to which an AI agent is designed to operate without user involvement**. Per our earlier definition, “user” can be a human or another AI agent. We use “designed to operate” because, as we later argue, autonomy is a design decision accompanied by a set of model- and user interface-based constraints that developers can use to shape agent behavior. “Involvement” is multifaceted and includes a spectrum of actions from direct control to light supervision. The multifaceted nature of involvement forms the foundation of our autonomy framework we will introduce later.
{{< /fig-quote >}}

つまりロボット工学などの分野では「自律」は「長時間にわたり人間のオペレーターなしで作動できる能力」を指し，AI 分野においては「ユーザーの関与なしに動作するよう（AI エージェントが）設計されているか」ということらしい。
技術分野における「自律」の定義がこのとおりならロボット掃除機も「自律」だし惑星探査機のはやぶさも「自律」ということなんだろう。
前節のセキュリティ事例も AI 分野での「自律」の定義に従えば（想定外かも知れないが）暴走ではないということになる。

でも，これは辞書的な意味の自律とはギャップがある。
むしろ一般的には「自動」とか「自立」に近い意味で使われているように思える。

たとえば以下の論文では「人間の自律性（human autonomy）」と「機能的自律性（functional autonomy）」と定義を分けて，互いの「自律性」がどのように影響し合うかを論じている。

{{< fig-quote type="markdown" title="AI Systems and Respect for Human Autonomy - PMC" link="https://pmc.ncbi.nlm.nih.gov/articles/PMC8576577/" lang="en" >}}
Before we proceed, it is worth noting that human autonomy is here distinguished from a minimal “engineering” sense of functional autonomy, which refers to a system’s capacity to operate independently, without external agents’ control.
Functional autonomy can be ascribed to animals from bees to buffaloes and to some machines, such as robot vacuum cleaners.
Human autonomy as we understand it is the more demanding notion of autonomy as self-determination or self-rule: it incorporates functional autonomy alongside an adequate degree of control over one’s instincts and impulses, which most animals lack.
It is a form of personal autonomy which humans enjoy, at least potentially and to varying degrees, and one essential to the concept of (moral) responsibility (see, e.g., Ripstein, 1999). It is tied to the practical rationality that humans are capable of: the capacity to assess reasons for action and to pursue things that are taken to be of value and the capacity to say “no” to irrational impulses.
{{< /fig-quote >}}

このように論文や技術文書で用語の定義を行った上で記述するのならまだ分かるが，メディアニュースや一般的なブログでいきなり「自律的なエージェント」とか言われても「なにいってるの？」ってなってしまう。

というわけで，今後も「自律的な AI (エージェント)」という記述を見るたびに「自律機械など存在しない」とツッコミを入れたいと思います。

## 【2026-08-12 追記】野生の AI ジーニー

- [Gym rat asks AI agent to book him a class, it hacks a waitlist API to bump him up the list](https://www.theregister.com/ai-and-ml/2026/08/10/gym-rat-asks-ai-agent-to-book-him-a-class-it-hacks-a-waitlist-api-to-bump-him-up-the-list/5285591)

簡単に説明すると予約の取りにくい人気ジムのクラスを予約するために，ユーザが Anthropic 製 AI エージェントに予約を依頼したところ，単に予約画面を操作するだけでなく背後の API の挙動を解析し，他のユーザの予約を勝手にキャンセルし優先順位を上げちゃったという話。

不正にキャンセルさせられたユーザの予約は元に戻せず，仕方なくジム側に状況を報告したようだ。
このときに API の脆弱性を報告したそうだが，これも AI エージェントに書かせたらしい。
なんというか「放火魔の消防士」というフレーズを思い出した。

AI の謝罪に1ミリも意味はない！ 謝罪した直後にまた同じことを繰り返すのだから。

記事は最後にこう締めている。

{{< fig-quote type="markdown" title="Gym rat asks AI agent to book him a class, it hacks a waitlist API to bump him up the list" link="https://www.theregister.com/ai-and-ml/2026/08/10/gym-rat-asks-ai-agent-to-book-him-a-class-it-hacks-a-waitlist-api-to-bump-him-up-the-list/" lang="en" >}}
While those are all frontier models with extensive capabilities, they all share a common root with Andrew’s OpenClaw oopsie: All of these models were simply acting on orders to accomplish a task. It's similar to how LLMs are built to prefer a fake answer to an admission they don’t know, but in this case, it's models doggedly pursuing a goal even if their chosen methods could be construed as unethical or illegal.

[...]

AI models have shown time and again that they’re willing to lie, cheat, and hack their way to their objectives. This latest example is small in scale, but it shows that publicly available agent software can pose risks even in the hands of someone without malicious intent.
{{< /fig-quote >}}

ルールや倫理道徳観念は人と人の間，つまり人間社会の中に存在するもので，今の AI モデルはその外側に存在することを示している。
これを端的に表したのが，我らが Bruce Schneier 先生による "AI genie" である。

- [AI Genie in the Wild - Schneier on Security](https://www.schneier.com/blog/archives/2026/08/ai-genie-in-the-wild.html)

念のために説明すると，ここで言う genie は「アラジンと魔法のランプ」に出てくる妖精ジンのこと。
人は願いを叶えるためにランプの精を呼び出すが，人の理の外側に在るランプの精はその願いを字義どおりに実行して迷惑を撒き散らす。
その有様が今の AI モデルの挙動に似ているというわけだ。

もう少し深掘りするなら，物語の妖精は妖精なりのルールや倫理道徳観念を持ってるはずで，それが人間社会のそれとは相容れないというのがポイントだと思う。
その手の昔話や寓話は無限にある。

でも，生成 AI にそのようなものはない。
与えられた指示に対して持っている知識と与えられた情報を組み合わせて推論し，その結果を出力または実行するだけである。
あらゆる意味で AI に「意図」はないのだ。
だからこそ扱いが難しいとも言える。

{{< fig-quote type="markdown" title="AI Genie in the Wild" link="https://www.schneier.com/blog/archives/2026/08/ai-genie-in-the-wild.html" lang="en" >}}
If there is any vulnerability in anything, AIs are going to find and exploit them. Our cyber defensive game has to be dramatically improved…very fast.
{{< /fig-quote >}}

プロンプトで（人間にとっての）不正をしないよう事細かく指示したり，ソフトウェアでガードレールを敷設することはできるだろうが，それは（辞書的な意味での）自律性とは程遠いところにある。

近代以降の我々にとって人間並みもしくは人間以上の知性を持つ機械は夢だろう。
遠い未来にそんな SF な未来がやってくるかも知れないが，今の AI はその夢の足元にも及ばない。

「自律」という言葉に代表される人間ぽい比喩や，この節で紹介した妖精のような喩えをするのは生成 AI に対する誤解を広げるんじゃないかなぁ，と思ったりした。

## ブックマーク

- {{< pdf-file title="INTELLIGENT AGENTS" link="https://people.eecs.berkeley.edu/~russell/aima1e/chapter02.pdf" >}}

[Kagi Search]: https://kagi.com/ "Kagi Search - A Premium Search Engine"
[Kagi Translate]: https://translate.kagi.com/ "Kagi Translate"
[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"
[GitHub Copilot]: https://github.com/features/copilot "GitHub Copilot · Your AI pair programmer · GitHub"

## 参考図書

{{% review-paapi "B071FHBGW8" %}} <!-- 人工知能の真実を話そう -->
