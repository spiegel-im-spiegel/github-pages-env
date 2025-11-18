+++
title = "そもそも「AI ブラウザ」ってなんなん？"
date =  "2025-11-18T23:12:31+09:00"
description = "Mozilla Firefox はどこへ行く"
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "google", "firefox", "market", "web", "security", "privacy", "risk" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

最近は他人様のブログ記事や SNS ポストの感想文みたいな記事ばっかり書いてるのだが，今回も yomoyomo さんのこの記事から。

- [「Mozillaさん、誰もFirefoxにAIなんか望んでないと思うよ」よりも前向きな提言 - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20251118/ai-in-firefox)

## Mozilla Firefox はどこへ行く？

なんでも Mozilla が Firefox に “AI Window” なる機能を組み込む[計画](https://blog.mozilla.org/en/firefox/ai-window/ "Introducing AI, the Firefox way: A look at what we're working on and how you can help shape it")があるそうで

{{< fig-quote type="markdown" title="FirefoxにAIタスク用の「AIウィンドウ」が追加される予定 - GIGAZINE" link="https://gigazine.net/news/20251114-firefox-ai-window/" >}}
AIウィンドウにはAIアシスタントとチャットしたりブラウジングを補助してもらったりできる機能が搭載される計画です。

[...]

AIウィンドウの機能はオプトイン形式で提供され、利用するかしないかはユーザーが自分で選択できるようになります。Firefox開発チームは「AI企業が開発したブラウザは、ユーザーに対してAIを常に使うかまったく使わないかの難しい選択を迫ります」と指摘し、FirefoxのAIウィンドウならAIを使いたいときにだけ使えるとアピールしています。
{{< /fig-quote >}}

これに対して「[要らんよ，そんなもん](https://manualdousuario.net/en/mozilla-firefox-window-ai/ "I think nobody wants AI in Firefox, Mozilla ⁄ Manual do Usuário")」という意見が出てくるのだが，さらにその意見に対して「[そうだけどそうじゃねーよ](https://www.anildash.com//2025/11/14/wanting-not-to-want-ai/ "I know you don’t want them to want AI, but… - Anil Dash")」と反論（？）しているのが「前向きな提言」というわけだ。

Mosaic 系列の Web ブラウザとして有名な Netscape Navigator と Internet Explorer だったが，後者が Windows 標準ブラウザとしてカーネルの深いところに組み込まれ，パソコン OS の世界市場独占を背景に猛威を振るう一方で，前者は最初こそ大きなシェアを誇っていたものの市場的には失敗に終わっている。
ユーザからは Netscape の後継プロジェクトと目されていた Firefox (当時は Phoenix と呼ばれていた) には IE のカウンタとして期待が寄せられていたが，最終的に IE に止めを刺したのは「[プライバシーに敵対的な企業]({{< ref "/remark/2018/04/handling-privacy.md" >}} "誰がプライバシーを支配するのか")」として名を馳せる Google の Chrome だったというのは皮肉としか言いようがない。

これは完全に私個人の偏見だが， IE が完全に倒れたことで Mozilla Firefox は存在意義を失い迷走しているように見える。
Mozilla は Firefox をセキュリティやユーザのプライバシー保護に配慮したブラウザとして宣伝するけれど，それにしては今だに検索窓と連携する検索サービスのデフォルトが ([メタクソ化](https://en.wikipedia.org/wiki/Enshittification "Enshittification - Wikipedia")して久しい) Google 検索なのは何故だろう[^m1]。
私はこの一点で Mozilla を完全には信用していない。
これで

[^m1]: もちろんそれが Mozilla の大きな収入源のひとつだからなのだろうが。

{{< fig-quote type="markdown" title="Introducing AI, the Firefox way: A look at what we're working on and how you can help shape it" link="https://blog.mozilla.org/en/firefox/ai-window/" lang="en" >}}
With AI becoming a more widely adopted interface to the web, the principles of transparency, accountability, and respect for user agency are critical to keeping it free, open, and accessible to all. As an independent browser, we are well positioned to uphold these principles.
{{< /fig-quote >}}

とか言われても説得力が... って感じではある。

## そもそも「AI ブラウザ」って何？

ここまで来てこう思わなかっただろうか。

{{< div-gen class="center" type="markdown" >}}
**そもそも「AI ブラウザ（AI-driven browser）」ってなんなん？**
{{< /div-gen >}}

たとえば [ChatGPT Atlas](https://openai.com/ja-JP/index/introducing-chatgpt-atlas/ "ChatGPT 搭載のブラウザー、ChatGPT Atlas が登場 | OpenAI") は，開発元の OpenAI によると「ChatGPT を中核に据えて構築した新しいウェブブラウザー」だそうだが，実際には[アンチ Web ブラウザ](https://www.anildash.com/2025/10/22/atlas-anti-web-browser/ "ChatGPT's Atlas: The Browser That's Anti-Web - Anil Dash")とでも言うべきものだろう。
だって Web にアクセスしないで，それらしいものを「生成」しているだけなんだもん。

{{< fig-quote type="markdown" title="ChatGPT's Atlas: The Browser That's Anti-Web - Anil Dash" link="https://www.anildash.com/2025/10/22/atlas-anti-web-browser/" lang="en" >}}
You're the agent for the browser, it's not being an agent for you
{{< /fig-quote >}}

じゃあ「AI ブラウザ」って何なのだろう。
ちゃんとした定義があるのだろうか。

「[AIブラウザとは？次世代“ブラウザ×AI”で何が変わるか](https://momo-gpt.com/column/ai-browser/ "AIブラウザとは？次世代“ブラウザ×AI”で何が変わるか｜機能・設定・活用ガイド - 株式会社MoMo")」によると

{{< fig-quote type="markdown" title="AIブラウザとは？次世代“ブラウザ×AI”で何が変わるか｜機能・設定・活用ガイド - 株式会社MoMo" link="https://momo-gpt.com/column/ai-browser/" >}}
AIブラウザは、以下のような多様な機能を一つの画面内で提供します：

- ページ要約・翻訳・解説：長文記事を瞬時に要点化、外国語ページの即時翻訳も
- 質問応答と並行作業：画面を見ながら「この製品のメリットは？」「比較して」と尋ねるだけ
- 情報抽出・整理：複数ページから必要情報を抜き出し、表形式で整理
- フォーム入力補助・リライト：SNS投稿、メール作成などの文章推敲をリアルタイム支援
- 自動操作（Agent Mode）：AIがタブを開き、リンクをクリックし、必要事項を記入する

単なる情報収集の補助ではなく、意思決定・実行レベルにまでAIの支援が拡張されています。
{{< /fig-quote >}}

と書かれている。
また「[AIブラウザで情報収集を時短｜仕事で使えるおすすめ7選と活用事例を紹介](https://ai.cloudcircus.jp/media/column/ai-browser-work "AIブラウザで情報収集を時短｜仕事で使えるおすすめ7選と活用事例を紹介 | コラム")」には「単なるWeb閲覧だけでなく、ユーザーと対話しながら情報収集や作業をサポートしてくれるのが大きな特徴」とあり，さらに

{{< fig-quote type="markdown" title="AIブラウザで情報収集を時短｜仕事で使えるおすすめ7選と活用事例を紹介 | コラム" link="https://ai.cloudcircus.jp/media/column/ai-browser-work" >}}
検索結果の要約、翻訳、文章のポイント抽出、比較、さらには自然言語での質問応答や資料作成の補助まで、さまざまな作業をブラウザ上で完結できます。ユーザーの関心や行動履歴に応じて関連情報を提案する機能もあり、必要な情報に素早くアクセスできるのも大きな魅力です。

さらに、音声操作や自動化機能も備えており、作業の手間を減らし、よりスムーズで効率的な情報収集が可能になります。情報収集から業務効率化まで幅広く活用でき、ビジネスシーンにおける活用の期待も高まっている注目のツールです。
{{< /fig-quote >}}

などと書かれている。

結局「定義」と呼べるものはないみたいだが，強いて言うならグレッグ・イーガンの『[万物理論](https://www.amazon.co.jp/dp/4488711022?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "万物理論 | グレッグ・イーガン, 山岸 真 |本 | 通販 | Amazon")』に出てくる「シジフォス」みたいなもんか（笑）

## AI ブラウザの議論と実装

Web と AI あるいは AI ブラウザについては W3C でも議論があるらしい。
以下，タイトルだけ挙げておく。

- [AI Agent Protocol Community Group](https://www.w3.org/community/agentprotocol/)
- [AI & the Web: Understanding and managing the impact of Machine Learning models on the Web](https://www.w3.org/reports/ai-web-impact/)
- [Machine Learning for the Web Community Group](https://www.w3.org/community/webmachinelearning/)
- [Web AI standardization in W3C](https://www.w3.org/2025/Talks/0906-webai-plh/)
- {{< pdf-file title="Threat Modeling Agentic AI Web Browsers" link="https://www.w3.org/2025/Talks/agentic-ai-web-browser-tm-simone.pdf" >}}
- [Web Machine Learning | Working Groups | Discover W3C groups | W3C](https://www.w3.org/groups/wg/webmachinelearning/)

既存の主要ブラウザの実装状況は以下の通り：

- Google Chrome は Gemini を統合している。最近 “[Private AI Compute](https://k-tai.watch.impress.co.jp/docs/news/2062738.html "グーグルが「プライベートAIコンピュート」発表、AI処理でプライバシー保護とクラウドの能力を両立 - ケータイ Watch")” と称してオンデバイス処理とクラウド・サービスを組み合わせた機能を発表した
- Microsoft Edge は Copilot モードによって AI 機能を提供している。 AI 検索については Bing Chat が統合されている
- Mozilla Firefox はオンデバイスの AI 機能を搭載している。チャット機能については各種 AI サービスと連携できる。 AI 検索機能についても他社との協業を検討しているらしい。前節の AI Window については詳細不明
- Apple Safari はオンデバイス処理と自社のクラウド・サービスを組み合わせた Apple Intelligence が統合されている。 AI 検索機能については他社との協業を検討しているらしい

ChatGPT/Copilot や Gemini のようなサービスを持っていない Mozilla や Apple はプライバシー重視の象徴としてオンデバイス処理を強調しているように見える。
Google の Private AI Compute はこれに対抗する動きと見られるが，どうやって「[クラウドで動作する高性能モデルの速度とパワーを活用しつつ、個人のデータを非公開に保つ](https://k-tai.watch.impress.co.jp/docs/news/2062738.html "グーグルが「プライベートAIコンピュート」発表、AI処理でプライバシー保護とクラウドの能力を両立 - ケータイ Watch")」のかはよく分からない（経路の暗号化だけじゃ駄目なんだよ）。
むしろ，オンデバイス OCR のように，プライバシー云々というより，即応性を上げるための分散処理なんじゃないの？ と思えなくもない。

## AI ブラウザのリスク

生成 AI をトリガーとして Web を含むユーザ体験（User Experience; UX）は確実に変わるだろう。
少なくとも検索サービスや SNS のタイムラインを情報ハブとしたユーザの今までの行動パターンからは既に変わりつつある。

機械に人生相談した挙げ句に[自殺をそそのかされる](https://gigazine.net/news/20251108-seven-families-suing-openai-chatgpt-suicides/ "「ChatGPTが自殺や悪い妄想をかきたてた」として複数の家族がOpenAIを訴える - GIGAZINE")みたいな事例も出ている。
また AI ブラウザが情報漏洩リスクになっている点も[指摘](https://gigazine.net/news/20251117-browser-copy-paste-data-leak/ "企業からのデータ流出ルートは「ウェブブラウザでのコピペ」が最多 - GIGAZINE")されている。

{{< fig-quote type="markdown" title="Why The Browser Has Become the Enterprise’s Most Overlooked Endpoint - LayerX" link="https://layerxsecurity.com/blog/why-the-browser-has-become-the-enterprises-most-overlooked-endpoint/" lang="en" >}}
AI browsers access sensitive corporate content through session data, cookies, and SaaS tabs to personalize results. Every tab, copy/paste, and login could quietly feed external AI models, creating what we call an “invisible AI endpoint.”

Unlike traditional browsers, these AI-driven ones operate outside enterprise visibility and DLP controls, turning session memory, auto-prompting, and cookie sharing into new exfiltration paths. And because employees adopt them alongside Chrome or Edge, most security tools never see them.
{{< /fig-quote >}}

先に挙げた W3C の “{{< pdf-file title="Threat Modeling Agentic AI Web Browsers" link="https://www.w3.org/2025/Talks/agentic-ai-web-browser-tm-simone.pdf" >}}” のような観点での議論がもっと必要だろう。

それでも私達はもう「生成 AI」より前には戻れないのだ。
そこに (Google や Microsoft などとは異なる立場に立てる) Mozilla Firefox がリーダーシップを取れる余地があると私は思っている。

{{< fig-quote type="markdown" title="I know you don’t want them to want AI, but… - Anil Dash" link="https://www.anildash.com//2025/11/14/wanting-not-to-want-ai/" lang="en" >}}
Market Firefox as “The best AI browser for people who hate Big AI”. Regular users have no idea how creepy the Big AI companies are — they’ve just heard their local news talk about how AI is the inevitable future. If Mozilla can warn me [how to protect my privacy from ChatGPT](https://www.mozillafoundation.org/en/privacynotincluded/articles/how-to-protect-your-privacy-from-chatgpt-and-other-ai-chatbots), then it can also mention that ChatGPT tells children how to self-harm, and should be aggressive in engaging with the community on how to build tools that help mitigate those kinds of harms — how do we catalyze that innovation?
{{< /fig-quote >}}

というわけで，最初に戻る。
果たして Mozilla Firefox はどこへ行くのだろう？

このまま Mozilla がどっちつかずの態度のままブラウザ市場におけるプレゼンスを縮小させていくのであれば，いっそ今からでも [Orion](https://kagi.com/orion/ "Orion Browser by Kagi") か [Zen](https://zen-browser.app/) あたりに鞍替えするのが無難だろうか[^o1]。

[^o1]: 手元にある [MacBook Air]({{< ref "/remark/2024/05/get-a-used-pc-from-workplace.md" >}} "勤務先からの払い下げ PC") には既に [Orion](https://kagi.com/orion/ "Orion Browser by Kagi") を入れている。 Linux 版の登場を待ってます！

## ブックマーク

- [Web のコストは誰が支払うのか]({{< ref "/remark/2025/03/who-pays-for-web-costs.md" >}})

## 参考

{{% review-paapi "4488711022" %}} <!-- 万物理論 -->
{{% review-paapi "B0C9Z7KGRN" %}} <!-- はじめて学ぶ ビデオゲームの心理学 Kindle 版 -->
