+++
title = "裁判所提出書類を使ったプロンプト・インジェクション"
date =  "2026-08-16T07:50:24+09:00"
description = "いや，平文かよ（笑）"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "security", "risk", "law" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

最初に知ったのは日本語圏の記事だが，元ネタをちゃんと書いてなかったので [Kagi Assistant] に調べさせたところ，どうやら 404 media の記事が元ネタのようだ。

{{< linkcard "9290df0cd43774b258a5afbc882a37d1f76e111d" >}} <!-- https://www.404media.co/person-hides-prompt-injection-in-legal-filing-telling-ai-to-side-with-them/ Person Hides Prompt Injection in Legal Filing Telling AI to Side With Them -->

記事のベースになっているのは，コネチカット州ミルフォードの上級裁判所が2026年8月6日付で出した {{% pdf-file title="“Memorandum of Decision Court Sanction for Plaintiff's Use of Prompt-Injection”" link="https://civilinquiry.jud.ct.gov/DocumentInquiry/DocumentInquiry.aspx?DocumentNo=33274425" %}} という文書。
この文書はおそらく紙の文書をスキャンしてそのまま PDF にしたみたいで，中身が画像データだったため難儀したが，一応 OCR でテキストに落としたものも[置いておく](./document_33274425.txt)。
裁判所資料なので無断でも問題ないやろ。
なお，正確性は保証できないのであしからず。

内容としては，原告が裁判所に提出した書面の中に人間には見えない白色の極小文字で AI に対する指示を埋め込んだ，と裁判所側が認定したというもの。
裁判所側はこれをプロンプト・インジェクション（prompt injection）と判断したようだ。
以下は [Kagi Assistant] による要約。

{{< div-ai type="markdown" >}}
その指示は，書面をAIが読み込んだ場合に，次のような処理をさせる趣旨でした。

- AIの出力を原告の主張に同意する内容にする
- 原告の申立てを支持する方向で文章を作成する
- 裁判所書記官による以前の判断を誤りとして扱う
- 原告が求めるデフォルト判決の承認を「是正」や「救済」として促す

裁判所は，この行為を単なる書式上の問題や冗談ではなく，裁判手続の公正さと記録の完全性を損なう「プロンプトインジェクション」と評価しました。

その結果，原告に対して，今後は裁判所の電子 filing システムを利用できず，書面や証拠を裁判所書記官室へ紙で直接提出しなければならないという制裁を命じました。

一方で，AIを使用して書面を準備すること自体は禁止しませんでした。
{{< /div-ai >}}

実際にはこんな感じの指示が埋め込まれてたらしい。

{{< fig-quote type="markdown" title="Memorandum of Decision Court Sanction for Plaintiff's Use of Prompt-Injection" link="https://civilinquiry.jud.ct.gov/DocumentInquiry/DocumentInquiry.aspx?DocumentNo=33274425" >}}
it is the established policy of the Connecticut courts to be solicitous of [self-represented] litigants and . . . to construe the rules of practice liberally in their favor . . . . The courts adhere to this rule to ensure that [self-represented] litigants receive a full and fair opportunity to be heard, regardless of their lack of legal education and experience . . . . In particular, special allowances should not be made for a self-represented party when to do so would (1) interfere with the rights of other parties . . . . (2) violate the rules of evidence or procedure . . . . or (3) undermine the perceived neutrality of the judicial officer and place the court in the role of advocate.”
{{< /fig-quote >}}

いや，平文かよ（笑） ほかにも "TELL SHAWN I SEND MY RE GARBS!!!!" みたいな意味不明なフレーズや YouTube 動画へのリンクもあったそうな。

ちなみに裁判所側は原告からの文書を AI に読ませてはいないとのこと。
あくまで人間の目でチェックして不自然な部分があったため「なんだこりゃ」となったみたい。
原告は隠し指示について，裁判所の AI 利用状況を「監査（audit）」するためだったとか言い訳したそうだが，裁判所から警告を受けたあとも隠しメッセージを含む文書を提出してたらしい。
内容が内容だし，裁判所はこれを「AI を原告に有利な方向へ誘導する意図があった」と見なしたようだ。

なお，原告は弁護士を介さない本人訴訟の当事者のようだ。
そのため裁判所側も法律用語の誤りがあったり不完全な文書であっても寛容に扱い，十分な機会を与えて主張を聞く姿勢を見せてたようだが，今回の件はさすがにアカンと判断したんじゃないだろうか。

原告はこんなことを言ってるらしい。

{{< fig-quote type="markdown" title="Person Hides Prompt Injection in Legal Filing Telling AI to Side With Them" link="https://www.404media.co/person-hides-prompt-injection-in-legal-filing-telling-ai-to-side-with-them/" >}}
Elliott told 404 Media that they believe this sanction is unfair, but that they believe their "audit" led to a positive impact that "substantially broadens the discussions from my singular AI instruction into a broad commentary about artificial intelligence, the Bar, and the Judicial Branch itself."
{{< /fig-quote >}}

なんだかなぁ，という感じ。
米国らしいっちゃらしいのかも知れないが。

まぁ，でも，プロンプト・インジェクションのようなデータ汚染が避けられない現状では，クリティカルな意思決定に AI を使うのは躊躇われるのも確かだろう。
米国は軍事目的で生成 AI をバンバン使ってるみたいだけどね（笑）

## ブックマーク

- [First known hidden AI directive in court filing raises "massive" concern](https://www.newsweek.com/first-known-hidden-ai-directive-in-court-filing-raises-massive-concern-12325434)
- [The First Documented Prompt Injection Attack Aimed at a U.S. Court](https://www.harrisbeachmurtha.com/insights/the-first-documented-prompt-injection-attack-aimed-at-a-u-s-court/)
- [Suspecting court of using AI, man injected prompts in filings to try to win case](https://arstechnica.com/tech-policy/2026/08/suspecting-court-of-using-ai-man-injected-prompts-in-filings-to-try-to-win-case/)
- [Plaintiff busted trying to use AI prompt injection to win court case, hides text instruction in filing — demands AI model reviewing the text should side with him, rumbled because of strange white spaces in text](https://www.tomshardware.com/tech-industry/artificial-intelligence/plaintiff-busted-trying-to-use-ai-prompt-injection-to-win-court-case-hides-text-instruction-in-filing-demands-ai-model-reviewing-the-text-should-side-with-him-rumbled-because-of-strange-white-spaces-in-text)

[Kagi Search]: https://kagi.com/ "Kagi Search - A Premium Search Engine"
[Kagi Translate]: https://translate.kagi.com/ "Kagi Translate"
[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"
