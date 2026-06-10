+++
title = "Ubuntu サーバで LLM 構築へ：（2）用語集"
date =  "2026-06-10T17:15:23+09:00"
description = "事前学習に当たって関連する用語を整理していく。随時更新する予定。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "lan-llm", "engineering", "linux", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = true
  jsx = false
+++

[前回]: {{< relref "./ubuntu-llm-01-study-notes.md" >}} "Ubuntu サーバで LLM 構築へ：（1）事前学習"

[前回]の続き。
事前学習に当たって関連する用語を整理していく。

最近はどうやってるか知らないが，私が若い頃はユーザ企業からの受託プロジェクトが多かったので，業務知識の摺り合わせのために用語集を作ることが必須だった。
これをサボると顧客との間で「そんなつもりじゃなかった」みたいなすれ違いが頻出して泥沼化することがある。
思ったより業界・業種ごとに独自の用語や隠語は多いものである。

用語については，まず AI に思いつく限り列挙してもらい，それを私が（勉強がてら）自力で調べる，という形で進めていく。
...という感じに指示したら，いきなり180個近く列挙して下さりやがった（笑）

まぁ，少しずつのんびりいこう。
この記事は随時更新する予定。

---

## 人工知能

### 別表記

Artificial Intelligence, AI

### 意味，定義など

元々は動物や人間の認知機能（の一部）を機械的に模倣（シミュレーション）する研究分野を指していたが，そこから派生した要素技術や製品・サービスも含めて AI と呼ばれるようになった。

## 機械学習

### 別表記

Machine Learning, ML

### 意味，定義など

人工知能（技術）のサブセット。
データから学習し，未知の入力に対する予測・分類・推定を行う技術の総称。

## 統計的機械学習

### 別表記

Statistical Machine Learning, Statistical ML, SML

### 意味，定義など

機械学習を「確率モデル」「推定」「推論」「不確実性」といった確率・統計学の枠組みで捉えて学習や評価を行うという立場（見方）を指す。

### ブックマーク

- [統計的機械学習とは| IBM](https://www.ibm.com/jp-ja/think/topics/statistical-machine-learning)

## 深層学習

### 別表記

深層機械学習, ディープラーニング, Deep Learning, DL

### 意味，定義など

多層（deep）のニューラルネットワークを中核にする機械学習のサブセット。
必ずしもラベル付きデータセットを必要とせず，テキストや画像など非構造化データから特徴を自動的に抽出して学習することができる点が特徴。

深層学習を統計的に見るなら統計的機械学習の文脈に入ると言えるが，深層学習という言葉が指すもの（実装・最適化・表現学習など）は確率論中心で語られるとは限らない。

### ブックマーク

- [AI，機械学習，ディープラーニング，ニューラル・ネットワーク | IBM](https://www.ibm.com/jp-ja/think/topics/ai-vs-machine-learning-vs-deep-learning-vs-neural-networks)

## ニューラルネットワーク

### 別表記

Neural Network, NN

### 意味，定義など

単純なニューロンを層状（入力層，中間層/隠れ層，出力層）に積み重ね，データからパターン認識の重みとバイアスを学習し，入力から出力までをマッピングする機械学習モデル。
脳の神経ネットワークから着想を得ている。

---

## 生成 AI

### 別表記

Generative AI

### 意味，定義など

与えられた知識（訓練データ）を元に文脈（コンテキスト）や指示に従って，テキスト（プログラムコードを含む）や画像，動画，音声などのコンテンツを生成する AI 製品のこと。

従来の AI 製品が情報の分析・評価・予測などを主目的としていたのに対して，生成 AI はコンテンツの生成を主目的としている点が特徴。
生成 AI が社会に与えたインパクトは大きく，近年では AI といえば生成 AI を指すことが多い。

## 大規模言語モデル

### 別表記

Large Language Model, LLM

### 意味，定義など

ニューラルネットワークを大規模化し，トランスフォーマーを採用して「次トークンを予測すること」を事前学習目的の中心にした機械学習モデル。
生成 AI におけるテキスト生成に関する部分の中核的な存在。
大規模言語モデルの「大規模」はモデルが大きい（パラメータ数が多い）ことを指して名付けられている。

まず事前学習（pre-training）として大量のテキストを使って「ラベルなし（自己教師あり）」でモデルに言語の基本を学ばせる。
その後の調整（fine-tuning / instruction tuning など）で特定のタスクに適応させることができる。

{{< fig-mermaid title="全体アーキテクチャ推奨構成（LiteLLM/Nginx あり）" >}}
flowchart TD
  A[事前学習<br/>Pre-training] --> B[微調整<br/>Fine-tuning]
  B --> C[指示追従の微調整<br/>Instruction tuning]
  B --> D[タスク別微調整例<br/>要約/分類/抽出など]
  B --> E[（場合により）さらなる調整<br/>例: 報酬学習/RLHF, 安全性調整]

  A:::pre
  B:::ft
  C:::it
  D:::task
  E:::post

  classDef pre fill:#e8f4ff,stroke:#2b6cb0
  classDef ft fill:#f0fff4,stroke:#2f855a
  classDef it fill:#fff7e6,stroke:#b7791f
  classDef task fill:#faf5ff,stroke:#6b46c1
  classDef post fill:#fff0f0,stroke:#c53030
{{< /fig-mermaid >}}

なぜ生成 AI が LLM 中心になったか：

- **トークン予測**： 単純で汎用的な学習目的で，大規模テキストから共通に学べる
- **トランスフォーマー設計**： attention 中心でスケールしやすい学習器として噛み合う
- **スケーリング則**： 増やしたときの改善が経験的に整理され，投資判断がしやすい
- **スケールでタスク適応が変化**： fine-tuning なしに few-shot で強くなる傾向が示され，普及・実装が加速

### ブックマーク

- [Scaling Laws for Neural Language Models](https://arxiv.org/abs/2001.08361)
- [Training Compute-Optimal Large Language Models](https://arxiv.org/abs/2203.15556)

## トランスフォーマー

### 別表記

Transformer

### 意味，定義など

自己注意機構だけに基づく新しいネットワーク構成。
再帰や畳み込みを使わずに，注意機構だけで構成しているため，並列化が可能で大規模化しやすい。

### ブックマーク

- [Attention Is All You Need](https://arxiv.org/abs/1706.03762)
- [Summary of the models — transformers 3.1.0 documentation](https://huggingface.co/docs/transformers/main/index)

---

## RAG

### 別表記

Retrieval-Augmented Generation

### 意味，定義など

LLM に外部知識（検索結果）を取り込ませてから生成させるアプローチ。
これにより知識不足や幻覚（hallucination）の問題を緩和し，より正確で信頼性の高い生成が可能になる。

実運用では以下のような手順になる。

| 手順 | 役割 |
|---|---|
| 1. 文書をチャンク化 | 検索しやすい単位に分割 |
| 2. チャンクを embedding に変換 | 文書意味を数値化（ベクトル化） |
| 3. ベクトルを Vector DB / Vector Index へ格納 | 類似検索できる形で保存 |
| 4. ユーザ質問を embedding に変換 | 質問も同じ空間で表現（ベクトル化） |
| 5. 類似度（例：cosine similarity）で上位を検索 | 関連文書を取り出す |
| 6. 取得した文書をプロンプトに入れて生成 | LLM の generator が回答 |

### ブックマーク

- [Retrieval-Augmented Generation for Knowledge-Intensive NLP Tasks](https://arxiv.org/abs/2005.11401)

## Embedding

### 別表記

埋め込み，ベクトル化

（複数形 Embeddings で）埋め込みベクトル群

### 意味，定義など

RAG において，外部文書やユーザ質問（プロンプト）などのテキストデータを数値ベクトルに変換する技術。
文書をチャンクと呼ばれる単位に分割し，チャンク化された各データに対して意味的な特徴を捉える多次元のベクトル情報を埋め込む。
ベクトルを埋め込まれたデータは key-value 形式で持つデータストア（Vector DB など）に保存される。

このため，日本語では Embedding をサ変名詞的に「ベクトル化」と訳すことが多い。

## MCP

### 別表記

Model Context Protocol

### 意味，定義など

AI アプリケーションが外部のデータや機能にアクセスするため，外部コンテキスト交換のためのオープンな標準プロトコル。
クライアント・サーバで構成される。
AI アプリケーション側がクライアント。
サーバ側は LLM に渡せるものを以下のプリミティブとして提供する。

| プリミティブ | 概要 | 典型例 |
|---|---|---|
| **Tools** | サーバが公開する「実行可能なツール」。LLM（言語モデル）が文脈に応じて発見・呼び出しできる形で提供されます。 | DB問い合わせ、API呼び出し、計算など |
| **Resources** | クライアント（LLMアプリ）にコンテキストとして渡すためのデータ。ファイルやDBスキーマ、アプリ固有情報などを共有できます。 | 社内文書、スキーマ、参照データ等 |
| **Prompts** | サーバが公開するプロンプトテンプレート。**ユーザが明示的に選べる**ように設計される、と説明されています。 | 業務用の定型指示テンプレなど |

RAG とは相補的な関係になることが多い。
たとえば RAG が外部資料を取得する方法を MCP で提供するなど。

---

## AI エージェント

### 別表記

AI Agent

単に Agent / エージェント と言う場合もある

### 意味，定義など

一般的に AI エージェントは「環境を知覚（perceive）し，環境に働きかけて行動（act）する自律的なシステム」として説明される。

AI エージェントが生成 AI と結びつくと以下のような役割・行動になる。

1. LLMを意思決定の中核にして
1. 計画して（plan）
1. ツールを呼び出し（call tools）
1. 複数ステップの実行に必要な状態（state）を保持して
1. 目的を達成するために動く。または推論と行動を繰り返す

目的別に AI エージェントを例示してみる。

| 目的（ユーザが求めること） | エージェントの核となる動き（イメージ） |
|---|---|
| 会話/問い合わせ対応 | 会話しつつ，必要なら外部情報やツールへ “取りに行って” 回答する（単発Q&Aより行動ループを持つことが多い） |
| 業務自動化（ワークフロー実行） | 次に何をするか計画し，ツール/API を呼んで段階的に完了させる（state 保持が重要） |
| 調査・意思決定支援 | 情報を集め，推論して結論を出す（必要に応じて “推論→行動→再推論” のループ） |
| 開発支援（コード生成/デバッグ） | 目標に向けて反復し，ツール（実行環境，リポジトリ操作等）を使いながら進める（目的達成型の行動） |

また，以下のように行動決定の仕方で分類する方法もある。

| 種類 | 特徴 |
|---|---|
| ツール実行を伴うエージェント | LLM が計画し，ツールを呼び，複数ステップを進める |
| 熟考（deliberative）/プランニング中心 | 世界の明示的なモデルを持ち，記号的推論で「何をするか」を決める |
| BDI エージェント | beliefs（信念）/desires（欲求）/intentions（意図）で分けて行動選択する枠組み |
| マルチエージェント | 複数のエージェントが協力して（autonomousのまま）タスクを進める |

なお，ここで言う「自律的（autonomous / autonomy）」とは「どれくらいユーザ介入なしで動けるように設計されているか」という文脈で使われ，辞書的な意味での自律とは異なる点に注意。

### ブックマーク

- [ReAct: Synergizing Reasoning and Acting in Language Models](https://arxiv.org/abs/2210.03629)
- [Agents SDK | OpenAI API](https://developers.openai.com/api/docs/guides/agents)
- [Deliberative agent - Wikipedia](https://en.wikipedia.org/wiki/Deliberative_agent)
- [Belief–desire–intention software model - Wikipedia](https://en.wikipedia.org/wiki/Belief%E2%80%93desire%E2%80%93intention_software_model)
- [マルチエージェント・システムとは何か| IBM](https://www.ibm.com/jp-ja/think/topics/multiagent-system)


<!-- ## AI 基礎

### 人工知能，AI (Artificial-Intelligence)
### 生成 AI (Generative AI)
### 機械学習（ML）
### 深層学習（Deep Learning）
### ニューラルネットワーク
### トランスフォーマー（Transformer）
### 自己注意機構（Self-Attention）
### トークン（Token）
### コンテキストウィンドウ（Context Window）
### マルチモーダル（Multimodal）

## LLM と推論

### 大規模言語モデル（LLM）
### 小規模言語モデル（SLM）
### 基盤モデル（Foundation Model）
### 推論（Inference）
### 事前学習（Pretraining）
### 事後学習（Post-Training）
### 指示チューニング（Instruction Tuning）
### アライメント（Alignment）
### RLHF
### DPO
### システムプロンプト（System Prompt）
### Temperature
### Top-k
### Top-p
### シード値（Seed）
### ストリーミング応答（Streaming）
### Function Calling
### Tool Calling
### JSON Mode
### 構造化出力（Structured Output）

## モデル管理と最適化

### モデル重み（Weights）
### パラメータ数（Parameter Count）
### VRAM
### RAM
### 量子化（Quantization）
### Q4_K_M
### Q8
### GGUF
### モデルカード（Model Card）
### チェックポイント（Checkpoint）
### LoRA
### QLoRA
### 蒸留（Distillation）
### Mixture of Experts（MoE）
### ファインチューニング（Fine-Tuning）
### 継続事前学習（Continued Pretraining）
### モデルマージ（Model Merge）
### コンテキスト長拡張

## RAG と検索

### RAG（Retrieval-Augmented Generation）
### 埋め込みモデル（Embedding Model）
### ベクトル化（Embedding）
### ベクトルデータベース（Vector DB）
### 近傍探索（ANN）
### コサイン類似度（Cosine Similarity）
### リランキング（Reranking）
### ハイブリッド検索（Hybrid Search）
### チャンク分割（Chunking）
### チャンクサイズ
### オーバーラップ
### メタデータフィルタ
### ナレッジベース（Knowledge Base）
### インデキシング（Indexing）
### 再検索（Retrieval）
### 幻覚（Hallucination）
### 引用付き回答（Cited Answer）

## エージェントと実行制御

### AI エージェント（Agent）
### ReAct
### Plan-and-Execute
### マルチエージェント
### オーケストレーション（Orchestration）
### ワークフロー（Workflow）
### タスク分解（Task Decomposition）
### メモリ（Agent Memory）
### スクラッチパッド（Scratchpad）
### ガードレール（Guardrails）
### ツール権限制御
### 自動承認（Auto-Approve）
### Human-in-the-Loop

## Ubuntu / LAN インフラ

### Ubuntu Server
### systemd
### Docker
### Docker Compose
### Podman
### Nginx
### リバースプロキシ（Reverse Proxy）
### HTTPS 終端（TLS Termination）
### 自己署名証明書
### Let's Encrypt
### LAN セグメント
### 固定IP（Static IP）
### DHCP
### DNS / mDNS
### ファイアウォール（ufw）
### ポート開放
### NAT
### VPN

## ローカル LLM サービス実装

### Ollama
### vLLM
### llama.cpp
### LiteLLM
### Open WebUI
### Continue
### OpenAI 互換 API
### /v1/chat/completions
### /api/chat
### SSE
### API キー管理
### レート制限（Rate Limit）
### 同時接続数（Concurrency）
### キューイング（Queueing）
### モデル自動アンロード

## 運用と監視

### 可観測性（Observability）
### ログ（Logging）
### メトリクス（Metrics）
### トレース（Tracing）
### レイテンシ（Latency）
### スループット（Throughput）
### TTFT（Time to First Token）
### トークン毎秒（tok/s）
### 稼働率（Availability）
### ヘルスチェック（Health Check）
### バックアップ
### リストア
### ローリングアップデート
### カナリアリリース

## セキュリティとガバナンス

### 認証（Authentication）
### 認可（Authorization）
### 最小権限（Least Privilege）
### シークレット管理（Secrets Management）
### 監査ログ（Audit Log）
### データ保持ポリシー
### PII（個人情報）
### プロンプトインジェクション
### データ漏えい（Data Exfiltration）
### Jailbreak
### Supply Chain Security
### SBOM
### ライセンス互換性
### 利用規約準拠

## 評価と改善

### ベンチマーク（Benchmark）
### 回帰テスト（Regression Test）
### 自動評価（LLM-as-a-Judge）
### 人手評価（Human Evaluation）
### 正確性（Accuracy）
### 忠実性（Faithfulness）
### 妥当性（Relevance）
### コスト最適化
### モデルルーティング
### A/B テスト -->
