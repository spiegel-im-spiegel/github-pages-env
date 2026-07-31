+++
title = "G7 による『オープンな AI』の定義"
date =  "2026-06-02T11:34:41+09:00"
description = "この定義が定着することを祈るよ。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "politics", "law" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

例によって Bluesky の TL で見かけた記事より：

- [G7 Agrees On Shared Language Around Open-Source AI, Open Weights AI - Phoronix](https://www.phoronix.com/news/G7-On-Open-Source-AI)

世間で言う所謂「オープンソース AI」について G7[^g7] でちゃんと定義しましょう，ということのようだ。
[GitHub Copilot] に簡単に要約してもらった。
こんな感じ：

[^g7]: G7（Group of Seven; 主要7か国首脳会議）のメンバーは，日本，アメリカ，カナダ，フランス，イギリス，ドイツ，イタリアの7ヶ国および欧州連合（EU）。これにゲスト国や国際機関が参加することもある。

{{< div-ai type="markdown" >}}
**G7 が合意した「オープンな AI」の共通定義**

2026年の G7 デジタル・技術大臣会合（フランス・エヴィアンで開催予定の第52回 G7 サミットに先立ち実施）では，オープンソース AI をめぐる用語の共通定義に合意した。

| 用語 | 定義 |
|------|------|
| **Open Source AI with Open Data** | モデルの重み・デプロイコード・学習コード・**完全な学習データ**を含め，オープンソースライセンスのもとで無償公開された AI |
| **Open Source AI** | モデルの重み・デプロイコード・学習コードを含め，オープンソースライセンスのもとで無償公開された AI。完全な学習データの公開が法的・技術的に不可能な場合は，代わりにデータに関する情報（Data Information）の提供が必要 |
| **Open Weights AI** | 重みとデプロイコードをオープンソースライセンスのもとで無償公開した AI（学習コード・学習データの公開は必須でない） |
| **Weights Available AI** | 重みとデプロイコードを無償公開しているが，商用利用・地域・用途などに制限を設けたライセンスで提供される AI |

これらは「AI オープン性のスペクトラム」として位置づけられており，Open Source AI with Open Data が最もオープンで，Weights Available AI が最も制限が多い。
{{< /div-ai >}}

ちなみに，ここで言う「無償」は “free of charge” のこと。
AI 企業が好き勝手絶頂なことを言ってるからねぇ。
この定義が定着することを祈るよ。

[Kagi Assistant] にも色々と訊いたのだが，今回のこれはポッと出てきたものではなく，2023年の G7 広島サミットで立ち上げられた「[広島 AI プロセス]」と関連しているようだ。

{{< div-ai type="markdown" >}}
| 時期 | 主な出来事 |
| :--- | :--- |
| **2024年5月** | 広島AIプロセスの精神に賛同する国々の自発的な枠組みである「広島AIプロセス・フレンズグループ」が発表される |
| **2025年2月** | フレンズグループの初の対面会合が東京で開催される |
| **2026年3月** | フレンズグループ第2回対面会合が開催され、「広島AIプロセス・フレンズグループ アクションプラン2026」が策定される |
{{< /div-ai >}}

- [お知らせ｜広島AIプロセス](https://www.soumu.go.jp/hiroshimaaiprocess/information.html)
- [Hiroshima AI Process Friends Group Meeting](https://www.soumu.go.jp/hiroshimaaiprocess/meeting/results.html "開催結果｜広島AIプロセス")
- {{< pdf-file title="「AIプロセス・フレンズグループアクションプラン2026」の概要" link="https://www.soumu.go.jp/hiroshimaaiprocess/pdf/document11.pdf" >}}

「[広島 AI プロセス]」は G7 が立ち上げた枠組みだけど，それ以外の国も巻き込むための組織が「フレンズグループ」で， G7 デジタル・技術大臣会合の「意思決定」と広島 AI プロセス・フレンズグループのアクションプランは相補的な関係になっているようだ。
だから G7 側で用語の定義をしたりするんだねぇ。

というわけで，日本もちゃんと AI 外交の仕事をしてるという話でした。
まぁ，市場が突出してて政治が追いついてないというのは，いつものことだが。

[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"
[GitHub Copilot]: https://github.com/features/copilot "GitHub Copilot · Your AI pair programmer · GitHub"
[広島 AI プロセス]: https://www.soumu.go.jp/hiroshimaaiprocess/ "広島AIプロセス"
