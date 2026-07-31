+++
title = "Ubuntu サーバで LLM 構築へ：（1）事前学習"
date =  "2026-06-09T18:26:51+09:00"
description = "まずは AI に構成を提案してもらい，それを起点に勉強していこうかな，と。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "linux", "ubuntu", "lan-llm", "tools", "engineering", "security" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = true
  jsx = false
+++

だいぶ生成 AI を使うことに慣れてきたので，そろそろ自前で LLM サービスを構築することを考えてみる。

いや，既にもうメジャーな AI サービスプロバイダってメタクソ化（enshittified）し始めてるぢゃん。
企業サービスに完全にロックインされないうちに自前で構築できるようになっておかないと，と思ってさ。

今回は，以下の要求で [Kagi Assistant] にサービス構成を提案してもらった。

1. 家庭内 LAN 上の Ubuntu サーバ機に LLM サービスを構築する
2. クライアントは Windows 機と Ubuntu 機を想定
3. 主な用途はプログラミング支援とブログ作成補助
4. VS Code で使えるようにする

以降に [Kagi Assistant] からの提案を載せるけど，まだ内容の検証をしていない。
なので嘘や間違いが紛れている可能性があるのでご注意を。
私としては，この提案を学習の起点（trigger）とするつもりである。

## アーキテクチャ

[Kagi Assistant] が最初に推奨してきた構成：

{{< div-ai >}}

{{< fig-mermaid title="全体アーキテクチャ推奨構成（LiteLLM/Nginx あり）" >}}
graph TB
    subgraph Server["Ubuntu Server"]
        Ollama["Ollama<br/>:11434<br/>(推論エンジン)"]
        LiteLLM["LiteLLM<br/>:4000<br/>(API Proxy)"]
        Nginx["Nginx<br/>:443<br/>(Reverse Proxy)"]
        Embed["nomic-embed-text<br/>(Embedding / RAG用)"]
        PG["PostgreSQL<br/>(LiteLLM 用)"]

        Ollama ---> LiteLLM
        LiteLLM --- PG
        LiteLLM ---> Nginx
        Ollama --- Embed
    end

    subgraph ClientUbuntu["Ubuntu PC"]
        VSCodeU["VS Code + Continue"]
    end

    subgraph ClientWin["Windows PC"]
        VSCodeW["VS Code + Continue"]
    end

    Nginx -- "LAN (HTTPS)" --> ClientUbuntu
    Nginx -- "LAN (HTTPS)" --> ClientWin
{{< /fig-mermaid >}}

{{< /div-ai >}}

推論エンジンとして [Ollama] を提案した理由は以下の通りらしい：

{{< div-ai type="markdown" >}}
個人〜小規模LAN利用ではOllamaが最適。セットアップが最も簡単で、OpenAI互換APIを標準提供し、モデル管理が容易です。

- vLLMは同時接続ユーザーが多い（5人以上）本番運用向けで、単一ユーザーには複雑さが増すだけ
  - [vLLM vs. Ollama: When to use each framework](https://www.redhat.com/en/topics/ai/vllm-vs-ollama)
  - [vLLM vs Ollama: Which Local AI Server to Run in 2026](https://www.fosslinux.com/156914/vllm-vs-ollama-which-local-ai-server-to-run-in-2026.htm)
- llama.cppは低レベルすぎて管理が煩雑
{{< /div-ai >}}

[LiteLLM] は [Ollama] の前に置く OpenAPI 互換 API プロキシ。
[Ollama] だけでなく OpenAI や Anthropic などの外部 API も統一的に扱えるらしい。
図では [LiteLLM] に [PostgreSQL] がぶら下がっているが，これは API キー管理・使用量トラッキング用とのこと。

- [What is stored in the DB | liteLLM](https://docs.litellm.ai/docs/proxy/db_info)

[Nginx] はリバースプロキシとして HTTPS 終端，セキュリティ（[Ollama]/[LiteLLM] への直結ポートをサーバの外に晒さない），ストリーミング対応を担っている。
注意点は以下の通り：

{{< div-ai type="markdown" >}}
- LAN 内のみの利用でもHTTPS推奨（Continue.dev が API キーを送るため）
- Let's Encrypt 証明書か自己署名証明書を使用
- ストリーミング（SSE）対応のため proxy_buffering off が必須
{{< /div-ai >}}

nomic-embed-text については後述する。

[LiteLLM] と [Nginx] は入れなきゃダメ？ と訊いてみたら，もう少し簡単な構成もあるらしい。

{{< div-ai >}}

{{< fig-mermaid title="全体アーキテクチャ推奨構成（Ollama 単体, LiteLLM/Nginx なし）" >}}
graph TB
    subgraph Server["Ubuntu Server"]
        Ollama["Ollama<br/>:11434<br/>(推論エンジン)"]
        OpenWebUI["Open WebUI<br/>:8080<br/>(GUIチャット)"]
        Embed["nomic-embed-text<br/>(Embedding / RAG用)"]

        Ollama --- Embed
        Ollama ---> OpenWebUI
    end

    subgraph ClientUbuntu["Ubuntu PC"]
        VSCodeU["VS Code + Continue<br/>ollama run (CUI)<br/>ブラウザGUIチャット"]
    end

    subgraph ClientWin["Windows PC"]
        VSCodeW["VS Code + Continue<br/>ollama run (CUI)<br/>ブラウザGUIチャット"]
    end

    Ollama -- "Ollama API<br/>:11434" --> ClientUbuntu
    OpenWebUI -- "Web UI<br/>:8080" --> ClientUbuntu
    OpenWebUI -- "Web UI<br/>:8080" --> ClientWin
    Ollama -- "Ollama API<br/>:11434" --> ClientWin
{{< /fig-mermaid >}}

{{< /div-ai >}}

クライアントの用途ごとの使い分けは以下の通り。

{{< div-ai type="markdown" >}}
| アクセス先 | ポート | 用途 | クライアント |
|-----------|--------|------|-------------|
| Ollama API | `:11434` | VS Code + Continue（コーディング支援） | Ubuntu / Windows |
| Ollama API | `:11434` | `ollama run`（CUIチャット） | Ubuntu / Windows |
| Open WebUI | `:8080` | ブラウザ GUI チャット（ブログ執筆・雑談） | Ubuntu / Windows |
{{< /div-ai >}}

[VS Code] の拡張機能である [Continue] は [Ollama] API に対応しているそうなので（セキュリティの問題を除けば）直結で OK。
ターミナルエミュレータ等で CUI チャットを行う場合はクライアント側にも [Ollama] をインストールして接続先をサーバ機に向ければいいらしい。

{{< div-ai type="markdown" >}}
```bash
OLLAMA_HOST=http://192.168.x.x:11434 ollama run qwen3.5:9b
```
{{< /div-ai >}}

さらに簡単に [Ollama] API を直に叩く方法もある（自作ツール向け？）。
こんな感じ。

{{< div-ai type="markdown" >}}
```bash
# ワンショット質問
curl http://192.168.x.x:11434/api/chat -d '{
  "model": "qwen3.5:9b",
  "messages": [
    {"role": "user", "content": "ブログのタイトル案を5つ提案して"}
  ],
  "stream": false
}'

# ストリーミング（リアルタイム出力）
curl http://192.168.x.x:11434/api/chat -d '{
  "model": "qwen3.5:9b",
  "messages": [
    {"role": "user", "content": "PythonでFizzBuzzを書いて"}
  ]
}'
```
{{< /div-ai >}}

[Ollama] API って素の RESTful API でやり取りするのか。
確かにこれは（セキュリティ的に）ヤバいな。
[LiteLLM]/[Nginx] を噛ませろって言う筈だわ。
LAN 内に他人（のデバイス）が入ってくる余地がないならギリ大丈夫かな。
まぁ，どのみち malware に侵入されたらアウトだけど。

[Open WebUI] は Web ベースのチャットインターフェースで，以下の特徴があるらしい[^rag1]。

{{< div-ai type="markdown" >}}
| 機能 | 説明 |
|------|------|
| ChatGPT 風 UI | ブラウザでチャット、使い慣れたインターフェース |
| 会話履歴 | すべての会話を保存・検索可能 |
| モデル切替 | ドロップダウンでモデル切替 |
| ドキュメント RAG | ファイルをアップロードして対話 |
| ユーザー管理 | 複数ユーザー・アクセス制御 |
| 完全オフライン | 外部へのデータ送信なし |
{{< /div-ai >}}

OpenAI の API キーを持っていれば [Open WebUI] から OpenAI のモデルも使えるそうな（以下は Docker 上に [Open WebUI] をインストールした場合の起動例）。

{{< div-ai type="markdown" >}}
```bash
# Open WebUI起動時にOpenAI APIも設定
docker run -d \
  --name open-webui \
  --network host \
  -v open-webui:/app/backend/data \
  -e OLLAMA_BASE_URL=http://127.0.0.1:11434 \
  -e OPENAI_API_KEY=sk-your-openai-key \
  --restart always \
  ghcr.io/open-webui/open-webui:main
```
{{< /div-ai >}}

用途別の使い分けについてまとめてもらった。

{{< div-ai type="markdown" >}}
| 用途 | インターフェース | 理由 |
|------|----------------|------|
| **ブログ執筆・雑談** | Open WebUI（ブラウザ） | ChatGPT風で直感的、会話履歴管理 |
| **プログラミング支援** | VS Code + Continue | コード補完・エディタ統合 |
| **サクッと質問** | `ollama run`（CUI） | ターミナルから即座に |
| **スクリプト・自動化** | curl / Python API | バッチ処理やパイプライン |
{{< /div-ai >}}

普段の調べ物は [Kagi Assistant] で必要十分なので [Open WebUI] はあまり使わないかもしれないけど。
[Open WebUI] が不要なら入れない構成もありだな。

## モデル選択

[Ollama] で選択可能な LLM モデルで [Kagi Assistant] が勧めてきたのは以下のもの[^rag1]。

{{< div-ai type="markdown" >}}
| 用途 | 推奨モデル | 役割 |
|------|-----------|------|
| **チャット・コーディング対話** | Qwen 3.6-27B または 35B-A3B | chat, edit, summarize |
| **コード自動補完（FIM）** | Qwen 2.5-Coder（7B/14B/32B） | autocomplete（タブ補完） |
| **Embedding（RAG用）** | nomic-embed-text | codebase検索 |
| **Reranker（RAG精度向上）** | Qwen3-Reranker-8B | 検索結果の再順位付け |

**GPU VRAM別のモデル構成**

| VRAM | チャット用モデル | 自動補完用モデル | 備考 |
|------|----------------|----------------|------|
| **8GB** | Qwen 3.5-9B (~6.6GB) | Qwen 2.5-Coder 7B (~5GB) | 1モデルずつ切り替え運用 |
| **16GB** | Qwen 3.6-35B-A3B (~16-22GB) | Qwen 2.5-Coder 14B (~9GB) | 同時稼働は量子化調整が必要 |
| **24GB** | Qwen 3.6-27B (~17GB) | Qwen 2.5-Coder 32B (~20GB) | ★最も推奨される構成 |
| **32GB+** | Qwen 3.6-27B (Q6/Q8) + Qwen 2.5-Coder 32B | 同時稼働余裕あり | RTX 5090等 |

💡 **24GB VRAM（RTX 3090/4090）の構成が最もバランスが良い**: Qwen 3.6-27Bでチャット・エージェント的コーディング、Qwen 2.5-Coder 32Bでタブ補完。ただし同時稼働はVRAM的に厳しいため、Ollamaはモデルを自動でアンロード/ロードするので、交互に使う運用になります。
{{< /div-ai >}}

あー。
そういや，サーバ機のスペックを教えてなかった。
手持ちの空いてるミニ PC をサーバ機にする予定なのでスペックは高くないのよ。

というわけで Ryzen 5 7430U, RAM 16GB, 外付け GPU なしの構成で再提案してもらった。

{{< div-ai type="markdown" >}}
| 役割 | 推奨モデル | RAM使用量 | 予想速度 | 理由 |
|------|-----------|----------|---------|------|
| **チャット・ブログ・コーディング対話** | **Qwen 3.5-9B** (Q4) | ~6 GB | ~5-8 tok/s | 9Bで品質最優、ブログ執筆に十分な表現力 |
| **コード自動補完（FIM）** | **Qwen 2.5-Coder 3B** (Q4) | ~2.5 GB | ~15-20 tok/s | 補完は高速が命、3Bで品質も十分 |
| **Embedding（RAG用）** | **nomic-embed-text** | ~0.3 GB | 高速 | コードベース検索用 |

💡 **合計RAM使用量: ~9 GB**（OS等の予約分を含めても16GB内に余裕あり）

チャット用を軽くして速度を稼ぐ構成：

| 役割 | モデル | RAM | 予想速度 |
|------|--------|-----|---------|
| チャット | **Qwen 3.5-4B** (Q4) | ~3 GB | ~12-18 tok/s |
| 自動補完 | **Qwen 2.5-Coder 1.5B** (Q4) | ~1.2 GB | ~30+ tok/s |
| Embedding | nomic-embed-text | ~0.3 GB | 高速 |

合計 ~4.5 GB。速度は快適ですが、ブログ執筆の品質は9Bに劣ります。

**🌟 注目の選択肢：Gemma 4 26B-A4B（MoE）**

GoogleのGemma 4 26B-A4Bは **総パラメータ26B・推論時アクティブ3.8B** のMoEモデルです。
Q4量子化で **~15 GB** のRAM使用量。

- ✅ **品質**: 3.8B活性化で26Bクラスの表現力（Qwen 3.5-9B以上の可能性）
- ⚠️ **RAM**: ~15 GBで16GBギリギリ。OSの予約分と競合する可能性
- ⚠️ **速度**: MoEはアクティブパラメータが少なくても全パラメータのメモリアクセスが必要なため、DDR4帯域幅がボトルネックになり **予想3-6 tok/s**

**結論**: 試す価値はありますが、16GBではスワップ発生のリスクが高く、安定性に欠ける可能性があります。まずはQwen 3.5-9Bで安定運用し、余裕があればGemma 4 26B-A4Bも試すのがおすすめ。

**⚡ パフォーマンスを引き上げるTips**

| 施策 | 効果 | 詳細 |
|------|------|------|
| **`num_ctx`を小さく** | 速度向上 | 4096→2048にするとコンテキスト処理が高速化。ブログ執筆なら2048で十分な場面も |
| **モデルを1つずつロード** | RAM節約 | `OLLAMA_MAX_LOADED_MODELS=1`で未使用モデルを自動アンロード |
| **Q3量子化を試す** | RAM削減・速度UP | `ollama run qwen3.5:9b-q3_K_M`でRAM~4.5GB、品質は若干低下 |
| **スワップの無効化** | 安定性向上 | RAM不足時のスワップは極端な低速化の原因。`/etc/fstab`でswapを0に |
| **将来的にGPU追加** | 劇的改善 | Thunderbolt/OCuLink外付けGPUエンクロージャーでRTX 3060(12GB)等を追加すれば、一気に実用的に |
{{< /div-ai >}}

Ryzen 5 7430U には内蔵 GPU (Radeon Graphics) があるが LLM 推論には非力なので CPU-only 推論となる。
そうなるとメモリ帯域幅が重要になるんだけど DDR4-3200 デュアルチャネルは DDR5 半分以下の帯域幅（〜51GB/s）しかない。
RAM のサイズも 32GB 以上は欲しいらしい。

{{< div-ai type="markdown" >}}
| サーバRAM | 推奨構成 | 備考 |
|-----------|---------|------|
| **8GB** | Qwen 3.5-4B (チャット) + Qwen 2.5-Coder 1.5B (補完) | 1モデルずつ切り替え |
| **16GB** | Qwen 2.5-Coder 7B (チャット) + Qwen 2.5-Coder 3B (補完) + nomic-embed | ★最も現実的 |
| **32GB** | **Qwen 3.5-35B-A3B (チャット)** + Qwen 2.5-Coder 7B (補完) + nomic-embed | ★★最高の体験 |
| **64GB** | Qwen 3.5-35B-A3B + Qwen 2.5-Coder 14B + nomic-embed | 余裕あり |
{{< /div-ai >}}

なかなかに貧弱な構成と言われてしまった（笑）

{{< div-gen class="center" type="markdown" >}}
**！だってメモリもプロセッサも高いんだよ，今！**
{{< /div-gen >}}

しょうがないぢゃん `orz`

### nomic-embed-text について

{{< div-ai type="markdown" >}}
nomic-embed-textはOllama上で動く **Embeddingモデル** です。テキストを数値ベクトル（768次元）に変換するだけで、それ自体にデータベースは必要ありません 。

```
テキスト入力 → nomic-embed-text → 768次元ベクトル出力
```
{{< /div-ai >}}

ベクトルデータはデータベース化して RAG (Retrieval-Augmented Generation; 検索拡張生成) に使用する[^rag1]。
[Open WebUI] および [Continue] は自前でベクトルデータベースを持っている。
規模が大きくなれば [PostgreSQL] + pgvector などのベクトルデータベースが別途必要になる。

[^rag1]: RAG の機能は，簡単にいうと LLM に外部の情報を検索させてから回答させる仕組み。これによって組み込まれた知識にない情報についても（extrinsic hallucination を抑えつつ）対応することができるようになる。 Google 検索のようなキーワード検索ではなく，プロンプトによる指示に合わせて意味的検索ができるのが特徴。 [Continue] は nomic-embed-text が出力したベクトル化情報をファイルに保存するらしい。 [Open WebUI] は自前でベクトル化を行い，内蔵の [ChromaDB] に保存するようだ。

{{< div-ai type="markdown" >}}
**🏗️ RAGの規模別アーキテクチャ**

| 規模 | ベクトル保存先 | 構成 |
|------|-------------|------|
| **小規模**（個人・Continue.dev） | ローカルファイル（Continue内蔵） | Ollama + nomic-embed-text のみ |
| **中規模**（チーム・Open WebUI） | Open WebUI内蔵のChromaDB | Ollama + Open WebUI |
| **大規模**（組織・自作パイプライン） | PostgreSQL + pgvector / Qdrant / Milvus | Ollama + 独自バックエンド |
{{< /div-ai >}}

## 今回はここまで

とりあえず [Ollama] と RAG についてはちゃんと勉強せなあかんなぁ，というのは分かった。
あとは実際に組んでみて評価していく感じかな。

薄々わかってはいたが，どうも手持ちの環境は貧弱すぎるみたいだけど，どの程度足らないのかの感覚がわからない。
これは使ってみて確かめるしかないだろう。

仕事が絡むなら vLLM も気になるところだが，そういう仕事が貰えるなら考えるということで（笑）

[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"
[Ollama]: https://ollama.com/ "Ollama"
[LiteLLM]: https://docs.litellm.ai/ "LiteLLM - Getting Started | liteLLM"
[PostgreSQL]: https://www.postgresql.org/ "PostgreSQL: The world's most advanced open source database"
[Nginx]: https://nginx.org/ "nginx"
[VS Code]: https://code.visualstudio.com/ "Visual Studio Code - The open source AI code editor | Your home for multi-agent development"
[Continue]: https://www.continue.dev/ "Continue • Quality control for your software factory. | Continue"
[Open WebUI]: https://docs.openwebui.com/ "Open WebUI"
[ChromaDB]: https://www.trychroma.com/ "Chroma - open-source search infrastructure for AI"
