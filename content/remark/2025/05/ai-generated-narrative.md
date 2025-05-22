+++
title = "AI が描く物語"
date =  "2025-05-22T23:20:31+09:00"
description = "LLM によって機械は「言語的理解」に必要なピースを獲得した。"
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "code" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

オチなし記事。
どんどん行こう（笑）

## AI が描く物語

今回の起点はこれ：

- [grokにファクトチェックさせている人は何も知らない | Books&Apps](https://blog.tinect.jp/?p=89620)

そういえば [{{% emoji "X" %}} 休眠中]({{< ref "/remark/2023/09/suspend-activity-on-twitter.md" >}} "𝕏 (旧 Twitter) の活動を休止します（期間未定）")の私は [Grok] について全く知識がないわ。
ちょっと [Kagi Assistant] に訊いてみた。

{{< fig-quote type="markdown" >}}
Grokは、xAIによって開発されたAIアシスタントで、X（旧Twitter）プラットフォーム上で利用できます。質問への回答、問題解決、ブレインストーミングなどを支援するように設計されています。[^1]

GrokはxAIの最先端の大規模言語モデル（LLM）を搭載しており、『銀河ヒッチハイク・ガイド』や『アイアンマン』のジャーヴィスから着想を得て、ユーモアを交えながら有益で洞察力のある応答を提供することを目指しています。[^1] [^2]

Grokの大きな特徴は、Xプラットフォームを通じてリアルタイムの情報にアクセスできる点です。これにより、他のモデルと比較して大きな優位性を持っています。[^3] [^4] トレンド分析なども可能です。[^5]

2025年2月20日には、Grok 3が無料開放され、有料プランのユーザーだけでなく、無料ユーザーも利用できるようになりました。[^6] [^7] ただし、無料開放はサーバー負荷による制限がある場合もあります。[^7]

Grokは、ほぼすべての質問に回答し、さらにはどのような質問をすべきか提案することも意図されています。[^2] 型破りな性格と、人類の部外者としての視点を持つ、ユニークで楽しいAIコンパニオンとされています。[^1]

参考になるWebページとしては、以下のものがあります。

* **Grok - xAI:** [https://x.ai/grok](https://x.ai/grok)
* **xAI: Welcome:** [https://x.ai/](https://x.ai/)
* **Xのユーモア溢れるAIアシスタント、Grokについて:** [https://help.x.com/ja/using-x/about-grok](https://help.x.com/ja/using-x/about-grok)
* **Grok:** [https://grok.com/](https://grok.com/)

[^1]: [Xのユーモア溢れるAIアシスタント、Grokについて](https://help.x.com/ja/using-x/about-grok)
[^2]: [Announcing Grok! - xAI](https://x.com/xai/status/1721027348970238035?lang=en)
[^3]: [Grok the AI (@groktheai) / X](https://x.com/groktheai?lang=en)
[^4]: [xAI (@xai) / X](https://x.com/xai)
[^5]: [Grok - xAI](https://x.ai/grok)
[^6]: [世界で最も賢いAI、Grok 3が登場しました。 現在どなたでも無料で ...](https://x.com/XcorpJP/status/1892482239742492682?lang=en)
[^7]: [XがGrok 3を無料開放！無料版ユーザー含めて誰でも利用可能に ...](https://usedoor.jp/news/2025-02-20-x-grok3-muryou-kaihou/)
{{< /fig-quote >}}

ええつと？ これって [Grok] はエンターテインメント特化ってことなのか？ ファクトが云々以前の問題じゃね？ よーわからん。

{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}

プログラム・コードなら実際に動かしてみれば間違っていることが分かる。
まぁ悪意のコードを混ぜ込まれればどうなるか分からないが，今のところ AI はそこまで[賢くなく]({{< relref "./dynastic-revolution-in-codebase.md" >}} "コードの易姓革命")（？）既存のコードを（要件に応じて）無難に組み合わせることしかできない。
プログラマにとってはそれだけで「[助かる](https://numan.tokyo/words/tasukaru/ "助かる（たすかる） | numan")」のだが（笑）

でも社会的な事実や真実が何かなど分からなくて当たり前だし，大抵の人は確かめようもない（統計データの読み方とか読解力はあるという前提でも）。
むしろ事実や真実を（都合よく）再構成した「物語（narrative）」の上に存在するのが私達の人間社会である。
人間はコンピュータのようにデータをデータのまま記憶・再利用するわけではないからだ。
前期近代と後期近代（つまり現代）との違いは，その物語がどの程度に受容・共有されているかの差だ。
人の多様性とは物語の多様性である。

LLM によって機械は「言語的理解」に必要なピースを獲得した。

{{< fig-quote type="markdown" title="「日経サイエンス」2025年6月号 脳が興味を生み出すとき 「やる気」の源泉に迫る" link="https://www.amazon.co.jp/dp/B0F64YN7KQ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
現在の大規模言語モデルのように広範な言語情報を獲得したAIは，すでに世界の地図を手にしたAIだ。
{{< /fig-quote >}}

これにより機械は人間に（単なるデータではなく）物語を語れるようになった。

物語を語れる AI がアシストすることにより人間の認知の一部を（機械が）肩代わりできる。
これは人の能力を飛躍的に向上させる期待がある（これ自体がSFの鉄板の物語）一方で，認知する「世界」を容易に歪める事ができる（これもSFの鉄板）。
よく「ナーロッパ」ファンタジーで鑑定スキルが人格を持って主人公と対話を始めるというプロットがあるが，鑑定スキル自体が認知の一部を肩代わりする機能なのに，それが人格を持つとか，なかなかにホラーである（笑）

好むと好まざるに関わらず今の LLM を発端とする何度目かの AI ブームは確実に（情報化）社会の構造を変えるだろう。
かつて SNS の台頭が社会構造を変えたように。

人と人を繋ぐと期待された SNS が企業および政治の広告媒体に成り果て，人はそれらが提供する物語の消費に堕すると2010年代前半[^sns1] までに予測できた人はきっと少ないだろう。

[^sns1]: 2010年代前半といえば「アラブの春」によって SNS への期待が最高潮に盛り上がった頃である。その後，2016年の米国大統領選挙によってその評価は地に落ちるのだが。

- [つながりのテクノロジーはまたしても我々を引き裂く – WirelessWire News](https://wirelesswire.jp/2025/05/88642/)

同様に今の AI ブームが何をもたらすか予測するのは難しい。
言うだけなら（この記事みたいに）誰でもどうとでも言えるし，当たればドヤ顔されるだろうが。

## 【2025-05-23 追記】 AI アシスタントと広告

- [Google is stuffing even more ads into its AI results | The Verge](https://www.theverge.com/news/671514/google-ai-mode-overviews-ads-expansion)

これも [Kagi Assistant] に要約させてみる。

{{< fig-quote type="markdown" >}}
Googleは、AI検索結果にさらに多くの広告を導入しています。具体的には、AI Modeでのテストと、デスクトップ版のAI Overviewsへの展開を進めています[^1]。

AI ModeはGoogle検索の新しいタブであり[^2]、AI Overviewsは検索クエリに基づいてAIが情報を要約し、検索結果の上位に表示する機能です[^3] [^4]。これらのAIが生成する応答の中に、「sponsored」と表示された広告が表示されるようになります[^5]。

[^1]: [Google is stuffing even more ads into its AI results | The Verge](https://www.theverge.com/news/671514/google-ai-mode-overviews-ads-expansion)
[^2]: [Google launches AI Mode to everyone in the US, adds... | The Verge](https://www.theverge.com/google-io/670439/google-ai-mode-search-io-2025)
[^3]: [2024年のGoogle AI Overviews：検索の未来を変える革新的な技術](https://hblab.co.jp/blog/google-ai-overviews-advantages-disadvantages-2024/)
[^4]: [Googleの「AIによる概要」がGemini 2.0にアップデート、より ...](https://gigazine.net/news/20250306-google-ai-overviews-gemini-2-0/)
[^5]: [Google Expands Sponsored Ads in AI-Powered Search Features](https://opentools.ai/news/google-expands-sponsored-ads-in-ai-powered-search-features)
{{< /fig-quote >}}

上の要約にはないが一応米国リージョンでのテストということらしい。
まぁ，全世界展開でないのは広告主との契約関係からだろう。

「さすが Google！ [おれたちにできない事を平然とやってのけるッ](https://dic.pixiv.net/a/%E3%81%9D%E3%81%93%E3%81%AB%E3%82%B7%E3%83%93%E3%82%8C%E3%82%8B%21%E3%81%82%E3%81%93%E3%81%8C%E3%82%8C%E3%82%8B%E3%82%A5%21 "そこにシビれる!あこがれるゥ! (そこにしびれるあこがれるぅ)とは【ピクシブ百科事典】")」って感じだろうか。
シビれもあこがれもしないけど（笑）

広告と（AI が描く）物語は，さぞかし相性がいいことだろう。
私は既に [Kagi Search](https://kagi.com/) に移行済みなので大したインパクトはない。

{{< fig-quote type="markdown" title="Kagi HQ: Kagi products will always be free of ads and trackers... — Bluesky" link="https://bsky.app/profile/kagi.com/post/3lpq6cq7tyk2p" lang="en" >}}
Kagi products will always be free of ads and trackers. In fact, Kagi Search will actively down-rank sites with lots of ads and trackers in the results and promote sites with little or no advertising.
{{< /fig-quote >}}

## ブックマーク

- [「好奇心に好奇心」を読む]({{< ref "/remark/2025/05/curiosity-about-curiosity.md" >}} "好奇心に好奇心")

[Grok]: https://grok.com/ "Grok"
[Kagi Assistant]: {{< ref "/remark/2025/04/kagi-assistant-for-all-users.md" >}} "Kagi Assistant が全ユーザに解放"

## 参考図書

{{% review-paapi "B0F64YN7KQ" %}} <!-- 「日経サイエンス」2025年6月号 -->
{{% review-paapi "4571210450" %}} <!-- はじめて学ぶ ビデオゲームの心理学 -->
{{% review-paapi "B07LG7TG5N" %}} <!-- FACTFULNESS ファクトフルネス -->
{{% review-paapi "B01J1I8PRQ" %}} <!-- 社会は情報化の夢を見る -->
{{% review-paapi "4791772091" %}} <!-- 後期近代の眩暈 -->
