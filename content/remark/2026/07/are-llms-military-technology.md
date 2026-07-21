+++
title = "LLM は軍事技術か？"
date =  "2026-07-16T16:30:30+09:00"
description = "LLM が国家の統制下に入るのと企業の独占所有物になるのと「どちらがいいか」と問われれば..."
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "security", "risk", "politics", "law", "privacy" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

さて，いつものように Mastodon/Bluesky の自 TL で見かけたポストより。
以下の yomoyomo さんの記事を起点に話を始めよう。

{{< linkcard "7290fdc7e8801c80a4189323c672c9423bd5ec04" >}} <!-- https://wirelesswire.jp/2026/07/94018/ 公益テクノロジーは国家に奉仕するのか -->

Bernie Sanders 米上院議員の話から始まっているが，この記事のメインは暗号学者でセキュリティ専門家（で私が勝手に先生呼ばわりしている） Bruce Schneier 氏と Palantir Technologies 社共同設立者兼 CEO である Alexander C. Karp 氏の対比だろう。
もう少し詳しくいうと Schneier 氏等の “[Rewiring Democracy]” と Karp 氏の “[Technological Republic]” との対比だ。

ちなみに記事中で

{{< fig-quote type="markdown" title="公益テクノロジーは国家に奉仕するのか――AI時代のテクノロジーと権力" link="https://wirelesswire.jp/2026/07/94018/" >}}
しかし、『テクノロジカル・リパブリック』と『Rewiring Democracy』を並べて論じている人をワタシは見たことがありません。
{{< /fig-quote >}}

と書かれていて，今まであまりそういう考察はされていないのかも知れない。
もっとも私は国際政治に疎いので「リパブリック」と聞いても先日話題になった某国大統領渾身のギャグのほうを思い浮かべてしまったり（笑）

{{< fig-img src="./islamic-republic-of-japan.png" title="アレについて Proton Lumo に漫画絵を描かせてみた" link="https://lumo.proton.me/" width="1280" >}}

話を戻して...

実際に両方の書籍を読んだわけではないので正しく理解しているか怪しいが，まずは最初の記事を更に [Proton Lumo] に要約してもらった。
こんな感じ：

{{< div-ai type="markdown" >}}
**シュナイアー vs. カープ（パランティア）の比較**

記事の核心は，シュナイアーの「公益テクノロジー」論と，パランティア共同創業者アレックス・カープの著書『テクノロジカル・リパブリック』の比較です。両者は「技術者が社会課題に関与すべき」という認識を共有しながらも，重要な違いがあります：

- **シュナイアー**：「公共」を広く捉え，社会正義や共通善を重視。国家による監視や権力集中に敏感。
- **カープ**：明確に国家（特に米国）を中心に据え，地政学的競争に勝つための技術者の関与を重視。強い軍事防衛能力を肯定。
{{< /div-ai >}}

Schneier 氏側を「国家による監視や権力集中に敏感」と評すのはよく分かる。

{{< fig-quote type="markdown" title="公益テクノロジーは国家に奉仕するのか――AI時代のテクノロジーと権力" link="https://wirelesswire.jp/2026/07/94018/" >}}
暗号学者というシュナイアーの出自を考えれば当然ですが、サイバーセキュリティやプライバシーの観点で、安全の名の下での国家による国民の監視や権力集中に対して彼は敏感です。
{{< /fig-quote >}}

（強力な）暗号はそもそもが軍事技術であり，つい前の世紀までは，国家によって統制されて然るべきものとされていた。
この辺は Steven Levy 氏の『[暗号化](https://www.kinokuniya.co.jp/f/dsg-01-9784314009072 "暗号化―プライバシーを救った反乱者たち")』を一読されることをおすすめする。

{{< linkcard "0475327ab9b35f0374f4211dfa0cc4e5e8679952" >}} <!-- https://www.kinokuniya.co.jp/f/dsg-01-9784314009072 暗号化 プライバシーを救った反乱者たち ブルース・シュナイアー Bruce Schneier -->

私は1990年代の PGP 愛用者で，国家の所有物だった暗号技術が一般庶民に降りてくる（誤解を恐れずに言うなら）「暗号技術の民主化」とでもいうような変化をリアルに体験してきた。
もし「暗号技術の民主化」がなければ HTTPS すらままならなかったろう。

一方で LLM を中心とする今の AI サービスは，私企業により最初から一般向けに提供されているものだ。
これが予想以上にハッキングに向いていたわけだ（単にハッキング・ツールとして優秀であるだけでなく AI サービス自体が hackable でもある）。

{{< linkcard "f49db55e98f0eb56c864acd2ee6f4da8da10016c" >}} <!-- https://bookplus.nikkei.com/atcl/catalog/23/09/25/01018/ ハッキング思考 ブルース・シュナイアー Bruce Schneier -->

故に国家（特に米国）が LLM を情報戦の（[核兵器に匹敵する](https://wired.jp/article/plaintext-how-to-start-an-ai-panic/ "AIは「核兵器のように危険」という警告の真意 | WIRED.jp")）軍事技術として統制をかけようとするのは自然な成り行きかも知れない。
先月，米国政府が Anthropic 社製モデルに対して[輸出規制]({{< ref "/remark/2026/06/why-us-government-banned-anthropic-models.md" >}} "米政府は何故 Anthropic モデルを禁じたのか")をかけたのも，やはりそういう流れの一部と見てよさそうである[^llm1]。

[^llm1]: Anthropic 社の Claude Fable 5 および Mythos 5 については今月早々に[輸出規制が解除](https://wired.jp/article/trump-administration-lifts-export-controls-on-anthropics-mythos-and-fable-ai-models/ "米政権、AnthropicのAIモデル「Mythos」と「Fable」への輸出規制を解除 | WIRED.jp")されている。一方で OpenAI 社の最新モデル GPT 5.6 も一時[公開制限](https://www.sbbit.jp/article/cont1/185903 "米政府、オープンAI新モデル「GPT-5.6」の公開制限を要請 ｜ビジネス+IT")が行われた。米国では高機能 LLM に対する政府の干渉が[常態化](https://www.businessinsider.jp/article/2606-news-washington-is-knocking-on-openais-doors-with-restrictions/ "ワシントンがオープンAIの扉をたたく。新モデル「事前承認」という厄介な前例 | Business Insider Japan")しつつあるようだ。余談だが Anthropic 社は Palantir Technologies 社と[パートナーシップ](https://investors.palantir.com/news-details/2024/Anthropic-and-Palantir-Partner-to-Bring-Claude-AI-Models-to-AWS-for-U.S.-Government-Intelligence-and-Defense-Operations/ "Anthropic and Palantir Partner to Bring Claude AI Models to AWS for U.S. Government Intelligence and Defense Operations")を結び自社製モデルを提供している。

最初に紹介した記事に戻ると，最後の方のオチが秀逸でちょっと笑ってしまった[^tr1]。

[^tr1]: そもそも Palantir Technologies 社は顧客に CIA, NSA, FBI, DHS といった情報機関や国防総省などの防衛機関を抱えている。また悪名高い ICE の移民強制送還プログラムにも関与していることが知られていて，批判の対象にもなっている。本文で挙げている “[Technological Republic]” についても，記事を読む限り，顧客に対するリップサービスが感じられて，正直共感できないんだよなぁ。まぁ，ナチス時代のドイツの例もあるし，独裁政権やマフィア（あるいは海賊）国家にぶら下がる企業や知識人ってのは，そういうものかも知れないが。

{{< fig-quote type="markdown" title="公益テクノロジーは国家に奉仕するのか――AI時代のテクノロジーと権力" link="https://wirelesswire.jp/2026/07/94018/" >}}
JS・タンは「[私がパランティアの本を読んだんで、あなた方は読む必要はないです](https://www.valueadded.tech/p/i-read-the-palantir-book-so-you-dont)」において、シリコンバレーのリベラルな価値観に対する軽蔑に目がくらんでしまい、この業界を単一のイデオロギー的ブロックとしてしか理解できていない著者たちの近視眼的な見方、2010年代にシリコンバレーの起業家が消費者向けビジネスに夢中だったのは、（シリコンバレーのルーツに米軍があるにせよ）インターネットの台頭とゼロ金利時代を背景に、消費者向けプラットフォームこそ手っ取り早く莫大な利益を生み出すビジネスモデルだったこと、逆に言うと2020年代、特に第二期トランプ政権下でGoogleをはじめとするビッグテックは軍との契約に積極的であり、『テクノロジカル・リパブリック』の主張は既に成立しなくなっている点について批判しています。
{{< /fig-quote >}}

（久しぶりにこんな長いセンテンスを見た（笑） 書籍の文章でも最近はなかなか見かけないよ）

私個人としては LLM が国家の統制下に入るのと企業の独占所有物になるのと「どちらがいいか」と問われれば，どちらもお断りと答える。
そのためにも早く AI サービス依存状態から脱しないとね。
今月半ばで既に [GitHub Copilot] の 10USD クレジットを使い切っちゃったよ `orz`

## ブックマーク

- [WirelessWire News連載更新（公益テクノロジーは国家に奉仕するのか――AI時代のテクノロジーと権力） - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20260715/wirelesswire)
- [China’s AI models have Trump’s AI world at war with itself](https://www.technologyreview.com/2026/07/20/1140675/chinas-ai-models-have-trumps-ai-world-at-war-with-itself/)

- [AI は正解を答えない]({{< ref "/remark/2026/02/ai-doesnt-provide-definitive-answers.md" >}})
- [「プロンプトウェア・キルチェーン」]({{< ref "/remark/2026/02/the-promptware-kill-chain.md" >}})
- [AI の軍事利用]({{< ref "/remark/2026/03/military-use-of-ai.md" >}})

[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"
[GitHub Copilot]: https://github.com/features/copilot "GitHub Copilot · Your AI pair programmer · GitHub"
[Proton Lumo]: https://lumo.proton.me/ "Lumo: Privacy-first AI assistant where chats stay confidential"

[Rewiring Democracy]: https://www.schneier.com/books/rewiring-democracy/ "Rewiring Democracy - Schneier on Security"
[Technological Republic]: https://techrepublicbook.com/ "THE TECHNOLOGICAL REPUBLIC"

## 参考

{{< linkcard "3c73c645059fc28681687d3e8eca2cf4f7cf7725" >}} <!-- https://www.schneier.com/books/rewiring-democracy/ Rewiring Democracy ブルース・シュナイアー Bruce Schneier -->
{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B0CK19L1HC" %}} <!-- ハッキング思考 Kindle 版 -->
{{% review-paapi "4560084122" %}} <!-- ヒトラーと哲学者 -->
