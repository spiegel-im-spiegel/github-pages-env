+++
title = "知識蒸留とモデル崩壊"
date =  "2026-07-23T18:33:51+09:00"
description = "知識蒸留と蒸留攻撃 / 知識蒸留とモデル崩壊"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "risk", "security", "law" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

[原文記事](https://www.technologyreview.com/2026/07/20/1140675/chinas-ai-models-have-trumps-ai-world-at-war-with-itself/ "China’s AI models have Trump’s AI world at war with itself | MIT Technology Review")の紹介ポストを Bluesky で見かけたのだが，ポップアップが煩くてまともに読む気がしなくて。
日本語訳がリリースされたので，こっちを今回の起点としよう。

{{% linkcard "b814392d6be5cb7959000f17b9484358de2ce622" %}} <!-- https://www.technologyreview.jp/s/386339/chinas-ai-models-have-trumps-ai-world-at-war-with-itself/ -->

## 知識蒸留と蒸留攻撃

Kimi はいわゆる中華 AI で Moonshot AI (月之暗面[^st1]) が提供するチャットボット・サービスおよびその基盤モデルを指す。
AI モデルの方は[オープンウェイト]({{< ref "/remark/2026/06/g7-definition-of-open-ai.md" >}} "G7 による『オープンな AI』の定義")・ライセンスで提供されているそうで，最新の K3 は性能的に [Claude Fable 5 を一部超えてる](https://www.watch.impress.co.jp/docs/news/2126048.html "2.8兆パラメータのAI「Kimi K3」　一部性能でFable 5超え - Impress Watch")とまで言われている。

[^st1]: Moonshot AI は中国では「AI 六小虎（Six AI Tigers）」と呼ばれるトップ AI スタートアップ企業の一角と評されている。「AI 六小虎」には他に「智谱 AI（Zhipu AI）」「稀宇科技（MiniMax）」「百川智能（Baichuan Intelligence）」「阶跃星辰（StepFun）」「零一万物（01.AI）」といった名前が挙がっている。もっとも Zhipu AI と Moonshot AI 以外は勢いも落ちてパッとしないみたい。まぁ，中国は今不景気だからねぇ。ちなみに中国では Alibaba，Tencent，ByteDance，Baidu といった既存の IT 大企業とは別に「AI四小龍（Four Dragons）」「AI六小虎（Six Tigers）」と呼ばれるスタートアップのカテゴライズがあるそうで，さらにそれとは別に DeepSeek のような企業もある。

Anthropic や OpenAI といった米国企業が高性能 AI モデルをリリースしてそれほど間を置かずに，それらに匹敵する性能を持つ中華 AI モデルが登場することで，それらの企業および（米国 AI 企業を推す）米国政権にとっての脅威となっている。
ただ脅威に対してどう対処すべきかという点で米国政権内で対立があるそうで，もの凄く大雑把にいうと「政府が（特定の AI 企業の利益を守るために）オープンソース競争を排除すべきではない」という立場と「AI モデルの国家安全保障上のリスクを強調し，政府による厳格なライセンス制度や介入が必要」という立場に分かれている，というのが記事のあらましである。

{{< fig-quote type="markdown" title="中国製無料AIモデル「Kimi」に揺れる米政権、対応をめぐり論争" link="https://www.technologyreview.jp/s/386339/chinas-ai-models-have-trumps-ai-world-at-war-with-itself/" >}}
しかし、Kimiはすでに公開され、無料で利用できる。その性能は、米国政府が国家安全保障上の脅威になり得るとして[一時的に公開を停止](https://www.technologyreview.jp/s/385116/three-things-to-watch-amid-anthropics-latest-feud-with-the-government/)させたアンソロピックのモデルにほぼ匹敵する。先週末の論争は、トランプ大統領の周辺にいる多くの人々が、これを重大な警鐘として受け止めていることを示している。しかし、それが何に対する警鐘なのかについては、誰一人として意見が一致していない。
{{< /fig-quote >}}

あまり深く触れられていないが，この記事のポイントは「知識蒸留（Knowledge Distillation）」だと思う。

{{< fig-quote type="markdown" title="中国製無料AIモデル「Kimi」に揺れる米政権、対応をめぐり論争" link="https://www.technologyreview.jp/s/386339/chinas-ai-models-have-trumps-ai-world-at-war-with-itself/" >}}
その過程では、「蒸留（distillation）」と呼ばれる手法が用いられた可能性がある。これは、既存のAIモデルが生成した出力を利用して別のAIモデルを学習させる手法である。オープンAIとアンソロピックは、中国のAI企業がこの手法を利用していると以前から主張しており、それを阻止するための政府による支援を求めてきた。4月、トランプ政権は蒸留を抑制するための一連の[施策](https://www.npr.org/2026/04/24/g-s1-118582/administration-crackdown-on-chinese-firms-exploiting-u-s-ai-models)を発表し、両社の要望はようやく実現した。
{{< /fig-quote >}}

ここで言う「蒸留」とは

{{< fig-quote type="markdown" title="知識の抽出とは| IBM" link="https://www.ibm.com/jp-ja/think/topics/knowledge-distillation" >}}
知識蒸留とは、大規模な事前トレーニング済みモデル（「教師モデル」）の学習内容を、より小規模な「生徒モデル」に転送することを目的とした機械学習アプローチです。これは、特に大規模なディープ・ニューラル・ネットワーク向けのモデル圧縮と知識転送の形式として、ディープラーニングで使用されます。
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="知識の抽出とは| IBM" link="https://www.ibm.com/jp-ja/think/topics/knowledge-distillation" >}}
近年、知識蒸留の研究は、大規模言語モデル（LLM）にとって特に重要となっています。LLMでは、KDは、よく使用されている有料モデルから、より小規模でアクセスしやすいオープンソース・モデルに高度機能を移行するのに効果的な手段とみなされています。
{{< /fig-quote >}}

といったことらしい。
実際に Meta や OpenAI では自社の小型モデルに対して知識蒸留を行っているそうで，業界全体でも知識蒸留自体は一般的な手法として認知されているようだ。

ただ中華 AI モデルの場合はやり方が問題のようで「蒸留攻撃（distillation attack）」などと呼ばれている。

{{< fig-quote type="markdown" title="Anthropic、中国AI企業による蒸留攻撃を公表　安全保障上の対応を要望 - Impress Watch" link="https://www.watch.impress.co.jp/docs/news/2088037.html" >}}
Anthropicは23日、DeepSeek、Moonshot、MiniMaxなどの中国のAIモデル各社により、Claudeの機能を不正に抽出する蒸留攻撃を検出していると発表した。これらの研究所は、約2.4万の不正アカウントを通じてClaudeと1,600万件以上のやり取りを生成し、Anthropicの利用規約や地域アクセス制限に違反したという。 

[...]

AIモデルのKimiを展開するMoonshot AIからは、340万件以上のやり取りが確認され、主体的推論とツール利用、コーディングとデータ分析などについて、数百の不正アカウントから組織的に実行されていたという。MiniMaxからは、1,300万件以上のやり取りが確認されている。

Anthropicは、国家安全保障上の理由から、中国と中国国外にある企業子会社に対してClaudeの商用アクセスを提供していないが、これを回避するために各社は商用プロキシサービスを利用。APIやサードパーティのクラウドプラットフォームにトラフィックを分散させ、不正な利用を試みているという。
{{< /fig-quote >}}

こうした状況に対して Anthropic 社は2026年2月のブログ記事にて

{{< fig-quote type="markdown" title="Detecting and preventing distillation attacks \ Anthropic" link="https://www.anthropic.com/news/detecting-and-preventing-distillation-attacks" >}}
But no company can solve this alone. As we noted above, distillation attacks at this scale require a coordinated response across the AI industry, cloud providers, and policymakers.
{{< /fig-quote >}}

と述べている。
「助けて！ トラえも〜ん」って感じだろうか（笑）

となると，先月の[「輸出規制」]({{< ref "/remark/2026/06/why-us-government-banned-anthropic-models.md" >}} "米政府は何故 Anthropic モデルを禁じたのか")の件も見方が変わってくる。
つまり，あれは企業自ら望んだ結果ではないのか。
もっと意地の悪い見方をすれば政府と企業の癒着の結果と見ることもできる。

「約2.4万の不正アカウント」と言うが，実際には商用プロキシサービスを経由した（誰かのアカウントを乗っ取ったわけではない）真っ当なものに見えるし，それだけのアカウントが知識蒸留のためにサービスにアクセスしたのなら Anthropic 社はさぞかし儲かったであろうと皮肉のひとつも言いたくなる。
その上で利用規約違反として金銭的なペナルティを課せば更に大儲けだ（笑）

そもそも米国企業の高性能 AI モデルだって元はと言えば人間の社会活動による「出力」を教師データとして事前学習に無断で利用[^fu1] してるわけで，それって「蒸留」と何が違うというのか。
今回の記事を読んでも，正直 Anthropic 社等に対しては「おまえがゆーな！」という感想しか湧かなかった。
米国に対しては... あそこはマフィアで海賊な国だし（笑）

[^fu1]: 「[連邦判事は2025年6月、Claudeの学習に作家の著作を利用したことはフェアユースに当たると判断した](https://www.itmedia.co.jp/aiplus/article/2607/21/2000000210/ "Anthropic、著作権訴訟で史上最大「2400億円」和解金支払いへ　学習利用は「フェアユース」認定 - ITmedia AI＋")」。つまり，フェアユースの要件を満たす限り著作物の AI 学習のための無断利用は合法ということだ（[和解金云々の話](https://www.itmedia.co.jp/aiplus/article/2607/21/2000000210/ "Anthropic、著作権訴訟で史上最大「2400億円」和解金支払いへ　学習利用は「フェアユース」認定 - ITmedia AI＋")はフェアユースの要件を超えていると見なされた）。ちなみに，米国のフェアユースでは「利用の目的や本質」「原作品の本質」「抜粋の量や実質性」「原作品の価値への影響」の4つの観点に基づいて「公正」かどうかが議論される。有名な例としては [Google Books の Library Book Scan に関する訴訟]({{< ref "/remark/2015/google-books-library-project.md" >}} "Google Books の Library Book Scan すら Fair Use と言われたのに...")が参考になるだろう。

とはいえ AI 技術界隈がルール無用のバトルロイヤル状態なのは確かだし，個人的には [Rewiring Democracy の夢が叶えばいいなぁ]({{< relref "./are-llms-military-technology.md" >}} "LLM は軍事技術か？")，とは思うけど。

## 知識蒸留とモデル崩壊

ところで「知識蒸留ってモデル崩壊を起こさないの？」と思い，いつものように [Kagi Assistant] に訊いてみたところ，（生徒モデルをさらに別のモデルの教師モデルにする）再帰的な蒸留を繰り返すことで世代ごとに劣化していく可能性はあるらしい。
曰く：

{{< div-ai type="markdown" >}}
**再帰的蒸留で劣化する理由**

1. **情報損失の蓄積** <br>各世代の蒸留で：
   - 教師モデルの出力分布の「希少なパターン（rare events）」が失われる 
   - 生徒モデルは**教師の能力を超えられない**
   - 世代を重ねるごとに**表現の多様性が徐々に減少**
2. **誤差の伝播**
   - 教師モデルの**バイアスや限界**がそのまま次世代に引き継がれる
   - 教師モデルの**間違いやノイズ**も学習データとして取り込まれる
   - 世代ごとに**誤差が累積**していく
3. **分布の収束**
   - モデルが**狭い表現空間に収束**していく 
   - 最終的には**無意味な出力**を生成するようになる 

ただし，現実のシナリオでは完全な劣化は起きにくい場合もある

1. **実データの混合**
   - 各世代で**実データ（人間が生成したデータ）を混ぜる**と，モデル崩壊は回避できる 
   - 純粋な合成データのみで学習し続ける場合に劣化が顕著

2. **教師モデルの更新**
   - 各世代で**新しい高性能モデルを教師**として使う場合
   - 例：GPT-4 → GPT-5 → GPT-6 と教師を更新し続ければ，劣化は抑制される

3. **蒸留の手法**
   - 単に出力をコピーするだけでなく，**中間層の表現や確率分布**を伝達する手法もある
   - このような手法では情報損失が少ない
{{< /div-ai >}}

とのこと。
あぁ，これ，裏取ってないので鵜呑みにしないでね。
でも再帰的な蒸留を繰り返すとモデル崩壊する可能性がある，というのは納得できる。

## ブックマーク

- [Anthropic joins OpenAI in flagging 'industrial-scale' distillation campaigns by Chinese AI firms](https://www.cnbc.com/2026/02/24/anthropic-openai-china-firms-distillation-deepseek.html)
- [Anthropicの『蒸留攻撃』公表で、なぜコミュニティはここまで割れたのか | ブログ](https://merchmindai.net/blog/ja/post/anthropic-distillation-attacks-community-reactions)
- [公益テクノロジーは国家に奉仕するのか――AI時代のテクノロジーと権力](https://wirelesswire.jp/2026/07/94018/)

[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"
[GitHub Copilot]: https://github.com/features/copilot "GitHub Copilot · Your AI pair programmer · GitHub"
[Proton Lumo]: https://lumo.proton.me/ "Lumo: Privacy-first AI assistant where chats stay confidential"

## 参考

{{< linkcard "3c73c645059fc28681687d3e8eca2cf4f7cf7725" >}} <!-- https://www.schneier.com/books/rewiring-democracy/ Rewiring Democracy ブルース・シュナイアー Bruce Schneier -->
