+++
title = "AI の軍事利用"
date =  "2026-03-07T11:50:35+09:00"
description = "最近思うのだが，今年はひとつくらい大手の AI 企業が倒れたほうが世の中のためなんじゃないだろうか。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "engineering", "politics" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

最初に私の基本的な考えを述べておくと，生成 AI を含む AI 技術の軍事利用に反対する理由はないと思っている。
前世紀の暗号の歴史を見ても最先端ソフトウェア技術の軍事利用は普通に行われていることだし，下手な正義感を振りかざして国際軍事バランスを傾けるのも馬鹿げた話ではある[^mb1]。

[^mb1]: まぁ，実際には露国や米国といった大国が率先して国際軍事バランスを傾けているのだが（笑）

その上で米国国防総省[^dod1] と AI サービスプロバイダである Anthropic および OpenAI との間の顛末（絶賛継続中）について Bruce Schneier & Nathan E. Sanders 両氏が書かれた記事が面白かったので   紹介してみる。

[^dod1]: 2025年9月にトランプ大統領は国防総省（Department of Defense）を戦争省（Department of War）に改称するための大統領令に署名した。正式な改称には議会の承認が必要だが，トランプ大統領が先走ったためか戦争省の名称は既に使われているらしい。本稿では Bruce Schneier & Nathan E. Sanders 両氏の[記事](https://www.schneier.com/blog/archives/2026/03/anthropic-and-the-pentagon.html "Anthropic and the Pentagon - Schneier on Security")に従って国防総省の名称を使うことにする。

- [Anthropic and the Pentagon - Schneier on Security](https://www.schneier.com/blog/archives/2026/03/anthropic-and-the-pentagon.html)
- [Don’t bet that the Pentagon – or Anthropic – is acting in the public interest | Nathan E Sanders and Bruce Schneier | The Guardian](https://www.theguardian.com/commentisfree/2026/mar/03/anthropic-openai-pentagon-ethics)

（どちらも内容は同じ）

顛末について簡単に紹介すると，まず Anthropic 社が自社の AI モデルを大規模監視（mass surveillance）や完全自律型兵器[^a1]（fully autonomous weapons）に利用することを認めないと主張したことで，国防長官である Pete Hegseth 氏が「[意識高い系（woke）](https://www.npr.org/2026/02/24/nx-s1-5725327/pentagon-anthropic-hegseth-safety "Hegseth threatens to blacklist Anthropic over 'woke AI' concerns : NPR")」だと皮肉り，これに乗った[反知性主義]({{< ref "/remark/2016/09/anti-intellectualism.md" >}} "ようやく『反知性主義』を読んだ")な大統領が連邦政府機関に対し Anthropic 製モデルの使用を中止するよう[命令を出した](https://www.theguardian.com/us-news/2026/feb/27/trump-anthropic-ai-federal-agencies "Trump orders US agencies to stop use of Anthropic technology amid dispute over ethics of AI | US news | The Guardian")らしい。
更に，この機に乗じた OpenAI 社がちゃっかり[割り込んで](https://www.nytimes.com/2026/02/27/technology/openai-agreement-pentagon-ai.html "OpenAI Reaches A.I. Agreement With Defense Dept. After Anthropic Clash - The New York Times")数億ドル規模の[政府契約](https://www.nytimes.com/2026/02/27/technology/anthropic-trump-pentagon-silicon-valley.html "Silicon Valley Rallies Behind Anthropic in A.I. Clash With Trump - The New York Times")を結ぼうかというところまで来ているそうな。
モノポリーゲーム（笑）

[^a1]: いつも同じことを言っていて恐縮だが，単に人間のオペレーションから離れて独立に動くことを指すのであれば「自律機械」ではなく「自立機械」あるいは「自動機械」とすべき。掃除機のルンバを自律機械とは言わないだろう？ 「自律」というからには自らの倫理・道徳観念に基づいて評価・判断・行動できることが要件で，現在そのような自律機械は AI を含めてまだ存在しない。将来は知らないが。

この顛末に対して記事では

{{< fig-quote type="markdown" title="Anthropic and the Pentagon - Schneier on Security" link="https://www.schneier.com/blog/archives/2026/03/anthropic-and-the-pentagon.html" lang="en" >}}
Despite the histrionics, this is probably the best outcome for Anthropic—and for the Pentagon. In our free-market economy, both are, and should be, free to sell and buy what they want with whom they want, subject to longstanding federal [rules](https://www.acquisition.gov/far/subpart-9.4) on contracting, acquisitions, and blacklisting.
{{< /fig-quote >}}

と評価し，加えて “The only factor out of place here are the Pentagon’s vindictive threats” と批判している。
執念深い脅迫... まぁ，確かに。

Anthropic の態度は明らかに顧客に向けた（ちゃんと貴方のほうを向いてますよという）アピールと分かるが OpenAI のほうは（顧客の側から見て）不透明さを残している。

{{< fig-quote type="markdown" title="Anthropic and the Pentagon - Schneier on Security" link="https://www.schneier.com/blog/archives/2026/03/anthropic-and-the-pentagon.html" lang="en" >}}
Despite the histrionics, this is probably the best outcome for Anthropic—and for the Pentagon. In our free-market economy, both are, and should be, free to sell and buy what they want with whom they want, subject to longstanding federal [rules](https://www.acquisition.gov/far/subpart-9.4) on contracting, acquisitions, and blacklisting.
{{< /fig-quote >}}

だからといって Anthropic マンセーかといえばそうでもなくて

{{< fig-quote type="markdown" title="Anthropic and the Pentagon - Schneier on Security" link="https://www.schneier.com/blog/archives/2026/03/anthropic-and-the-pentagon.html" lang="en" >}}
We can admire Amodei’s stance, but, to be sure, it is primarily posturing. Anthropic knew what they were getting into when they [agreed to a defense department partnership](https://www.anthropic.com/news/anthropic-and-the-department-of-defense-to-advance-responsible-ai-in-defense-operations) for $200m last year. And when they [signed a partnership](https://www.anthropic.com/news/anthropic-and-palantir-partnership) with the surveillance company Palantir in 2024.
{{< /fig-quote >}}

そりゃあダブスタちゃうんかい！ というわけだ。

国防総省の方も強引に AI 企業を従わせる必要があるかといえば，それも疑問。

{{< fig-quote type="markdown" title="Anthropic and the Pentagon - Schneier on Security" link="https://www.schneier.com/blog/archives/2026/03/anthropic-and-the-pentagon.html" lang="en" >}}
The Pentagon, meanwhile, has plenty of options. Even if no big tech company was willing to supply it with AI, the department has already deployed dozens of [open weight](https://www.wired.com/story/open-ai-artificial-intelligence-open-weight-model/) models—whose parameters are public and are often licensed permissively for government use.
{{< /fig-quote >}}

民間で出回っている AI を軍事転用するというのは（技術的な部分を除いても）簡単なことではない。

{{< fig-quote type="markdown" title="Anthropic and the Pentagon - Schneier on Security" link="https://www.schneier.com/blog/archives/2026/03/anthropic-and-the-pentagon.html" lang="en" >}}
The Pentagon is not a normal customer; it buys products that kill people all the time. Tanks, artillery pieces, and hand grenades are not products with ethical guard rails. The Pentagon’s needs reasonably involve weapons of lethal force, and those weapons are continuing on a steady, if potentially [catastrophic](https://thebulletin.org/2026/02/anthropics-showdown-with-the-us-department-of-war-may-literally-mean-life-or-death-for-all-of-us/), [path](https://www.theguardian.com/news/2020/oct/15/dangerous-rise-of-military-ai-drone-swarm-autonomous-weapons) of [increasing](https://www.wired.com/story/us-military-robot-drone-guns/) [automation](https://fsi.stanford.edu/sipr/content/lethal-autonomous-weapons-next-frontier-international-security-and-arms-control).
{{< /fig-quote >}}

ここで，最初の（国防総省やトランプ政権に向けた）批判に戻ってくるわけだ。

{{< fig-quote type="markdown" title="Anthropic and the Pentagon - Schneier on Security" link="https://www.schneier.com/blog/archives/2026/03/anthropic-and-the-pentagon.html" lang="en" >}}
But, of course, this is the Trump administration, so it doesn’t stop there. Hegseth has threatened Anthropic not just with loss of government contracts. The administration has, at least until the inevitable lawsuits force the courts to sort things out, [designated the company](https://www.nytimes.com/2026/02/27/us/politics/anthropic-military-ai.html) as “a supply-chain risk to national security,” a designation previously only ever applied to foreign companies. This prevents not only government agencies, but also their own contractors and suppliers, from contracting with Anthropic.

The government has incompatibly also threatened to invoke the [Defense Production Act](https://www.lawfaremedia.org/article/what-the-defense-production-act-can-and-can't-do-to-anthropic), which could force Anthropic to remove contractual provisions the department had previously agreed to, or perhaps to fundamentally modify its AI models to remove in-built safety guardrails. The government’s demands, Anthropic’s response, and the legal context in which they are acting will undoubtedly all change over the coming weeks.
{{< /fig-quote >}}

なんちうか，メチャクチャだなぁ。
こんなことされたら Anthropic は交渉のテーブルに再び乗らざるを得ない。
さすがマフィア国家は一味違う。
シビれもあこがれもしないが（笑）

最近思うのだが，今年はひとつくらい大手の AI 企業が倒れたほうが世の中のためなんじゃないだろうか。
そうすれば今の狂躁状態から目が覚めるんじゃないの？

それはともかく，この顛末から得られる教訓は「企業は民衆のヒーローにはなり得ない」ということだ。
当たり前だけどね。
私達は GAFA 等のビッグテックがメタクソ化する前に掲げた「企業倫理」がいかに薄っぺらいか，この20年で嫌というほど経験してきたではないか。

もし「大規模監視」や「完全自律型兵器」への AI 軍事利用を民衆が良しとしないのであれば，そうした軍事活動に対して法的制限を設ける必要がある，と記事は主張する。
同様に企業製品の安全でない利用を強要するために政府の権力が行使されることに私達が不快感を抱くのであれば，政府調達に関する法的保護を強化すべきとも主張している（激しく同意）。
政府や企業の「良心」に期待するなというわけだ。

{{< fig-quote type="markdown" title="Anthropic and the Pentagon - Schneier on Security" link="https://www.schneier.com/blog/archives/2026/03/anthropic-and-the-pentagon.html" lang="en" >}}
The Pentagon should maximize its warfighting capabilities, subject to the law. And private companies like Anthropic should posture to gain consumer and buyer confidence. But we should not rest on our laurels, thinking that either is doing so in the public’s interest.
{{< /fig-quote >}}

ダークヒーローに良心回路は組み込まれない。

## ブックマーク

- [AI は正解を答えない]({{< ref "/remark/2026/02/ai-doesnt-provide-definitive-answers.md" >}})
- [「プロンプトウェア・キルチェーン」]({{< ref "/remark/2026/02/the-promptware-kill-chain.md" >}})

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B0CK19L1HC" %}} <!-- ハッキング思考 Kindle 版 -->
{{% review-paapi "B00F5P454W" %}} <!-- キカイダー02 -->
