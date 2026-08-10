+++
title = "技術的負債と生成 AI"
date =  "2026-08-05T12:00:44+09:00"
description = "顧客に「自分でやったほうがはやい」と思われたら，そこで試合終了"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "engineering", "technical-debt", "management", "risk" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

ふと思いついて [Kagi Assistant] に「生成 AI が世の中に登場してからの大まかな年表って作れる？」って訊いてみたら作ってくれた。
年表自体は公開しないが AI が参考にした Web ページは以下らしい。

- [【年表】生成AIの進化史｜Transformer登場からAGI実現予測まで](https://www.generativeai.tokyo/media/genai-history/)
- [History of Generative AI — AI Timeline](https://aitimeline.world/history/history-of-generative-ai)
- [The History of AI: A Timeline of Artificial Intelligence](https://www.coursera.org/articles/history-of-ai)
- [生成AIの歴史｜多層年表｜GENAI-RON](/notes/history-of-generative-ai/)

ターニングポイントになっているのは2017年の Google による論文 "Attention Is All You Need" の Transformer アーキテクチャと，2022年の OpenAI による一般向けの ChatGPT サービスの登場のようだ。
今年に入ってからはエージェント型 AI が台頭し，完全に仕事の仕方が変わってしまった。

というわけで，以下の記事。

- [AI普及でデザイン業の倒産が前年比2.7倍に　「独自性なき企業は淘汰」　東京商工リサーチ](https://www.itmedia.co.jp/news/article/2608/04/2000000377/)

グラフィックデザインなどの「デザイン業」で一千万円以上の負債を抱えて倒産する企業が増えているという記事。
背景として以下を挙げている。

{{< fig-quote type="markdown" title="AI普及でデザイン業の倒産が前年比2.7倍に　「独自性なき企業は淘汰」　東京商工リサーチ" link="https://www.itmedia.co.jp/news/article/2608/04/2000000377/" >}}
東京商工リサーチは、倒産増加の背景を大きく3つに分類している。1つ目は、生成AIの普及を受け、顧客企業のデザイン業務の内製化が進んだこと。2つ目は、デジタル広告の需要拡大により、紙媒体を主力とする企業が影響を受けたこと。3つ目は、ハウスメーカー（木造建築工事業）の倒産増加に伴い、内装や家具などのインテリアデザインを手掛ける会社に影響が及んだことと分析している。
{{< /fig-quote >}}

私はアートデザインの分野はさっぱり分からないので断言できないのだが，上の記事を読む限り主因は2つ目と3つ目じゃないの？ と思ってしまう。
つまり産業構造の変化だ。

ぶっちゃけてしまうが，顧客に「自分でやったほうがはやい」と思われたら，そこで試合終了なのよ。

自分のところ（IT 業界）の話で申し訳ないが，今世紀に入って，いわゆる受託開発の仕事はどんどん減っている。
というか，もう殆どないかも知れない。
例外は公共事業関連くらいかな。
まぁ，日本は公共事業への依存度が高いんだけど。

理由は色々とあるが，事象としてはまさに「自分でやったほうがはやい」という風潮になり，これに適応できない中小企業はブラック化（極端に安い単価で過剰に受ける）して，最終的には潰れた。
ゼロ年代の話である。

生成 AI，特にエージェント型 AI によって仕事の仕方は大きく変わったが IT 産業の構造としてはとっくにシフト済みで，バカな管理職・経営者が従業員を AI で置き換え（られ）るとか言い出さない限り，業界として大きな構造変化はないと思う。
とはいえ，個々のエンジニアについては（特に年配者は）適応できない人もいるだろうし（そっと耳を塞ぐ），本当に AI 導入を理由にエンジニアを解雇するバカな経営者も（特に米国では）いるみたいだが[^w1]。

[^w1]: 業態が変わって合わないエンジニアを解雇するというのは昔からあるが，AI 導入を理由（名目）にエンジニアを解雇するような企業が，あとで「やっぱ人間要るわ」となって雇い直そうとしても，ヤバすぎて近付きたくないと思ったりするんじゃないのかなぁ。一度浮気をした男は必ず二度三度と繰り返す（笑）

「技術的負債」というものがある。
この言葉が流行ったときは，できの悪いコードや製品を揶揄して言うことが多かったが，本来の意味としては（銀行でお金を借りるように）「より早く実現するために引き換えにするもの」を指す言葉らしい。
なので，この「負債」自体には良し悪しというものはない。

- [お前も技術的負債にしてやろうか！ もしくは技術的負債と和田卓人さんをめぐるシンクロニシティ - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20201210/technical-debt)

ソフトウェア・エンジニアリングにおいて生成 AI を使うメリットは「そこそこのものを早く作れる」点である。
この点で AI は人間より優れていると（たぶん）断言できるし，それ故に生成 AI と技術的負債は相性がいいと言えるだろう。

となると，考えないといけないのは，「そこそこ」で収まらない部分に対して AI が残した負債（そもそも負債と気付かない可能性もあるが）をどうやって返すのか，ということだ[^p1]。
知り合いと話すときは大体ここで（エンジニアの教育も含めて）行き詰まる。
Web 上の技術関連のサイトやページを眺めても，そういったことに言及する記事はほとんど見かけない。
まぁ，それこそが本当の意味で「技術（＝お金を取れるところ）」だからなんだろうけど。
AI のために残業するとか論外だしね。

[^p1]: 本文の受託開発がなくなった話とも絡むが，今は要件定義から製造・デプロイまでの一連の工程を一発で完遂させるようなプロジェクトはないと言っていい。大抵は小さな単位に分割して，持続的にイテレーションを回していくスタイルだ（故にプロジェクトではなくミッションと言ったりする）。「技術的負債を返す」というのはこのスタイルを前提にしている。これも例外は政府・自治体主導の公共事業で，大抵は予算と期間が決まってるし，開発フェーズと運用（保守）フェーズが分かれているので「持続的にイテレーションを回す」ことができない。もっとも今は公共事業でも大手が持ってるパッケージを使うケースが多いかもしれない。これなら自社製品のカスタマイズとして受けられるので，企業側も（持続的でない）公共事業への依存を減らせる。

## ブックマーク

- [AIが変えた震災支援……熊本の被災者がスマホで作った「イマココナビ」活況　“乱立”に課題も](https://www.itmedia.co.jp/news/article/2608/10/2000000486/)

- [技術的負債とハッカー]({{< ref "/remark/2020/12/technical-debt-and-hacker.md" >}})
- [技術的負債と FOSS]({{< ref "/remark/2021/07/technical-debt-and-foss.md" >}})

[Kagi Search]: https://kagi.com/ "Kagi Search - A Premium Search Engine"
[Kagi Translate]: https://translate.kagi.com/ "Kagi Translate"
[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"
[GitHub Copilot]: https://github.com/features/copilot "GitHub Copilot · Your AI pair programmer · GitHub"
[Proton Lumo]: https://lumo.proton.me/ "Lumo: Privacy-first AI assistant where chats stay confidential"
