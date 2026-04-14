+++
title = "Streamplace と Livepeer"
date =  "2026-04-12T22:23:13+09:00"
description = "Livepeer を基盤とする Streamplace の仕組みと，分散型動画配信インフラの可能性・課題を整理してみる。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "media", "tools", "streaming", "blockchain", "atproto", "market", "fintech" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = true
  jsx = false
+++

[前回]の続き。

というか，[前回]の記事が思ったより長くなったのと，話の中心が [Streamplace] のほうに横滑りしそうになったので記事を分けたのだった。

## Livepeer について

[Streamplace] の話をする前にバックエンド（基盤技術）のひとつである [Livepeer] について説明しておこう。

[Livepeer] は Ethereum blockchain 上に構築された分散型のビデオインフラストラクチャ・プラットフォーム，だそうだ。
Blockchain 関連は Ethereum が登場するあたりからどんどん興味が薄れていったのだが，こんなまともそうなものもあるんだ。

[Livepeer] は以下の疑問が出発点となり，2017年に設立されたらしい。

{{< fig-quote type="markdown" title="Livepeer Story - Livepeer Docs" link="https://docs.livepeer.org/v2/home/about-livepeer/vision" lang="en" >}}
What if video - the most captivating, expressive medium online - ran on open, permissionless infrastructure instead of expensive, proprietary systems?
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="Livepeer Story - Livepeer Docs" link="https://docs.livepeer.org/v2/home/about-livepeer/vision" lang="en" >}}
Since its launch in 2017, Livepeer’s mission has been to deliver affordable, performant open video & AI infrastructure rooted in decentralised video technology enabled by cryptoeconomic primitives.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="Livepeer Story - Livepeer Docs" link="https://docs.livepeer.org/v2/home/about-livepeer/vision" lang="en" >}}
Doug and Eric set out to revolutionise video streaming infrastructure from the hardware infrastructure up by:

- Turning idle GPUs around the world into a competitive marketplace,
- Using crypto-economic incentives to guarantee reliability and quality,
- Delivering measurable cost reductions via decentralised economies of scale.
{{< /fig-quote >}}

[Livepeer] の特徴について [Kagi Assistant] にまとめてもらった。

{{< div-box type="markdown" >}}
- **分散型トランスコーディング**
    - 動画配信で最もコストがかかる「トランスコーディング（動画を視聴者の環境に合わせた形式に変換する処理）」を、世界中のコンピューター（ノード）に分散して実行させます
    - これにより YouTube や AWS などの巨大な中央集権的サーバーを利用するよりも大幅に低コストで配信インフラを構築できます。
- **ネットワークの参加者**
    - **オーケストレーター（Orchestrators）**: 自分のコンピューターのリソース（CPU/GPU）を提供して動画処理を行い報酬を得る参加者です。
    - **デリゲーター（Delegators）**: 自分が持つ LPT トークンを信頼できるオーケストレーターに預け（ステーキング）、ネットワークのセキュリティに貢献しながら報酬の一部を受け取る参加者です。
- **LPT (Livepeer Token)**
    - Livepeer の独自トークン「LPT」は、ネットワークの調整、ガバナンス、およびセキュリティ維持のために使用されます。
- **AI との統合**
    - 最近ではビデオ配信だけでなく、分散型 GPU リソースを活用した **AI 推論（リアルタイムのビデオ生成や加工など）** のインフラとしても機能を拡張しています。
- **開発者向けツール**
    - 開発者が自分のアプリに簡単にビデオ機能を組み込めるよう、SDK や「Livepeer Studio」といったホスト型プラットフォームも提供されています。
{{< /div-box >}}

だいたい合ってる？ [Streamplace] や [Livepeer] 周辺で Web3 が云々とか聞こえてきたのはそういう訳だったんだねぇ。

## Streamplace について

ここで [Streamplace] が登場する。
[Streamplace] は2024年後半にパイロット版が公開され，2025年の ATmosphereConf のライブ配信に採用されて[注目を集めた](https://au.finance.yahoo.com/news/beyond-bluesky-apps-building-social-150000953.html "Beyond Bluesky: These are the apps building social experiences on the AT Protocol")。
[2026年の ATmosphereConf](https://atprotocol.dev/announcing-atmosphereconf-vancouver-march-2026/ "Announcing ATmosphereConf: Vancouver, March 2026") でも [Streamplace] によるライブ配信が行われていた（まぁ，私はそれで知ったのだが）。

[Streamplace] は，この [Livepeer] が提供する強力な分散型ビデオ処理能力をバックエンド（基盤）として利用することで特定の企業に依存しない自由な配信プラットフォームを実現する，ということのようだ。
[前回]の記事でも挙げたが「特定の企業」のひとつである YouTube と [Streamplace] との比較を再び挙げておく（AI による要約なので鵜呑みにしないように）。

{{< div-box type="markdown" >}}
| 比較項目 | Streamplace | YouTube |
| :--- | :--- | :--- |
| **運営形態** | 分散型（特定の企業に依存しない） | 中央集権型（Google が運営） |
| **基盤技術** | AT Protocol / Livepeer ネットワーク | Google 独自のクローズドなインフラ |
| **ソースコード** | オープンソース（誰でも利用・改善可能） | 非公開（プロプライエタリ） |
| **データの所有権** | ユーザー自身（自己主権型アイデンティティ） | 運営企業（Google） |
| **主な目的** | 分散型 SNS 向けのビデオインフラ提供 | 広告収益を主とした動画共有サービス |
| **検閲・制限** | プロトコルレベルでの検閲は困難 | 運営企業のポリシーにより削除・停止あり |
| **主な利用者層** | Web3 開発者，分散型 SNS ユーザー | 一般消費者，クリエイター，広告主 |
{{< /div-box >}}

もちろん，基盤技術として [AT Protocol (Authenticated Transfer Protocol)][atproto] が使われているのもポイントで，つまり [Streamplace] は「[万能アカウント（The Everything Account）](https://www.augment.ink/the-everything-account/)」上のアプリケーション（のひとつ）である，とも言えるわけだ（分散サービスではあるが，匿名サービスではない）。

{{< fig-quote type="markdown" title="The Everything Account" link="https://www.augment.ink/the-everything-account/" lang="en" >}}

{{< fig-img src="/remark/2026/02/leaflet-and-the-atmosphere/Screenshot_20260204_121413_Samsung-Notes.jpg" width="1000" lang="en">}}

The Atmosphere, as you can see above, already has countless apps you can use. There are Twitter-like apps such as [Bluesky](https://bsky.app/) and [Blacksky](https://blacksky.community/); blogging services like [Leaflet](https://leaflet.pub/), [Offprint](https://offprint.app/), and [pckt](https://pckt.blog/); collection and annotation tools like [Semble](https://semble.so/), [Margin](http://margin.at/), and [Seams](https://seams.so/); and I can go on and on because the ecosystem is expanding by the day. And this is just a small portion of the existing Atmosphere - I couldn't fit all of the different apps because there are *just. so. many.*
{{< /fig-quote >}}

なお VOD (Video on Demand) については，この記事を書いてる時点でまだ開発・テスト中の段階らしい。
特にフロントエンド側が未整備のようで，私達一般ユーザが利用できるのはまだ先のようだ。

[Livepeer] が分散ネットワークであるということは [Streamplace] はセルフホストしないといけないのか？ と一瞬思ったが，一般ユーザ向けには [Streamplace] 公式の環境（[`stream.place`][Streamplace]）が用意されている（使い方は[前回]記事参照）。
組織向けや自前で動画配信プラットフォームを構築したいユーザはセルフホストもできますよ，というスタンスみたい[^sh1]。
[Streamplace] がもっと注目されれば（今 PDS サーバがボコボコ立ち始めているように）一般ユーザが選択可能なインスタンスがもっと増えるのかも知れない。

[^sh1]: なお，セルフホスト用に Docker 環境が用意されているそうな。 WebRTC を利用するので（UDP ポートをいっぱい開ける），その辺の知識がないと難しいかも。

## Streamplace のマネタイズってどうなってるの？

でも，そうなると [`stream.place`][Streamplace] のマネタイズってどうなってるのか気になる。
ベンチャーキャピタルから資金調達しまくって，挙げ句の果てにメタクソ化か？ とか意地悪なことも考えたが，ちょっと違うらしい。

[Streamplace] は [Livepeer Treasury](https://docs.livepeer.org/v2/about/protocol/treasury "Livepeer Treasury - Livepeer Docs") から助成金を受けて開発を行っているらしい（[LPT 払い？](https://forum.livepeer.org/t/agent-spe-phase-2-pre-proposal/2828 "Agent SPE: Phase 2 Pre-Proposal - Treasury - Livepeer Forum")）。
[Livepeer] にとって [Streamplace] はエコシステム上の「公共財（public good）」と[見なされて](https://forum.livepeer.org/t/streamplace-2-0-solving-video-for-everybody-forever/2847 "Streamplace 2.0: Solving Video for Everybody Forever - Treasury - Livepeer Forum")いて， [Streamplace] を育成することが [Livepeer] エコシステムの拡大に繋がると考えているようだ。

一方 [Streamplace] 側は Gateway ビジネスモデルによるマネタイズを将来的に考えているらしい。
アーキテクチャと支払いフローはこんな感じ。

{{< fig-quote class="mermaid" title="Gateway Architecture - Livepeer Docs" link="https://docs.livepeer.org/v2/gateways/concepts/architecture" lang="en" >}}
flowchart LR
    subgraph ApplicationLayer [Application Layer]
        direction TD
        A1["User Apps\nWeb, Mobile,\nTouchDesigner"]
    end

    subgraph GatewayLayer [Gateway Layer]
        direction TD
        A1 --> J[Job Intake]
        J --> A1
        G1["Capability Discovery"]
        G2["Pricing & Selection"]
        G3["Routing / Session Mgmt"]
    end

    subgraph OrchestratorLayer [Orchestrator Layer]
        direction TD
        J --> GW[GPU Worker]
        GW --> J
        AI["AI Models / BYOC\nContainers"]
        T[Transcoder]
    end

    style ApplicationLayer stroke:#3cb540,stroke-width:2px
    style GatewayLayer stroke:#3cb540,stroke-width:2px
    style OrchestratorLayer stroke:#3cb540,stroke-width:2px
{{< /fig-quote >}}

{{< fig-quote class="mermaid" title="Gateway Business Model - Livepeer Docs" link="https://docs.livepeer.org/v2/gateways/concepts/business-model" lang="en" >}}
flowchart LR
    A["End users / \nApplications"]
    G["Gateway"]
    O["Orchestrator"]
    W["Transcoder / \nAI Runner"]
    C["Arbitrum \n(on-chain)"]
    R["Gateway revenue"]

    A -- "Pay gateway's rate \n(business layer)" --> G
    G -- "Pay network rate \n(ETH payment tickets)" --> O
    O -- "Pay workers" --> W
    O -- "Redeem tickets" --> C
    G -- "Keep margin" --> R
{{< /fig-quote >}}

この “End users” は配信者を指すのか？ 視聴者も含むのか？

上の図は [Streamplace] を含む [Livepeer] エコシステムに対する支払いモデルだ。
でも，これだと最終的にかかるコストを配信者と視聴者の間でどう配分するかについて不透明な感じがする（見落としがあったらゴメン）。
一応

1. アプリ開発者や配信者が使用したリソース分を支払う従量課金モデル（クラウド資源利用とかでよく見る方式）
1. 高度な機能や安定した帯域を保証する代わりに視聴者が一定額を支払うサブスクリプションモデル（Amazon Prime Video や Netflix みたいな感じ）

とかいった方法で収益化を行うことを考えてるみたいだけど...

YouTube などは基本的にオーディエンスの規模が大きいほど（視聴者への接触機会が多いほど）配信者が稼げる構造になっているが，広告モデルであるが故に配信者も視聴者もコストを払っている感覚が薄い（スパチャのような投げ銭に近いものもあるが）。
そういう意味でも YouTube は現代版「テレビ」なんだよなぁ。

動画配信に限らないけど，こういうのって{{% ruby "トレンド" %}}流れ{{% /ruby %}}を激変させるほどのキラーコンテンツが現れるかどうかなんだよな。
そう考えると [Streamplace] はまだ「整ってない」感じがするし，整ったとしてもバズるかどうかは運次第といったところもある。

P2P や Web3 などに興味のある人にとっては注目すべき点もあろうが，一般のユーザにとっては「もうしばらく様子見かな」といったところか。
やっぱ VOD 機能が正式にサービスインされないと是とも非とも言えないなぁ。

でも，私みたいに今まで「配信」とは無縁だった周回遅れの人が「ちょっとやってみようかな」と思える程度にはなっているかな。

## 暗号通貨/資産はゼロサムゲームか？

私個人の感覚で申し訳ないが， [Livepeer] が Ethereum blockchain 上に構築されていると聞いて身構えてしまうのは「暗号通貨/資産って，結局のところゼロサムゲームなんじゃないの？」と思ってしまうからだ。
特に NFT (Non-Fungible Token) とか極まってるよね。
生成 AI が登場して，そちらに興味が移ったら，あっという間に廃れたけど（笑）

...ていう雑談を [Kagi Assistant] としてたのだが AI 曰く

{{< div-box type="markdown" >}}
投資家同士のマネーゲーム（富の移転）で終わるのではなく、そのプロジェクトが **「世の中をどう便利にするか」「どんな問題を解決するか」** という実需が伴うことで、市場全体が成長するポジティブサムな関係が成立します。
{{< /div-box >}}

なんだそうな。
確かに NFT には「実需」はなかったな。
LPT はどうなんだろう。
[Livepeer] エコシステムを通じて何かしらの「実需」を創造できるのだろうか。

## ブックマーク

- [About Stream.Place - Livepeer Docs](https://docs.livepeer.org/v2/solutions/streamplace/overview)
- [Gateway Business Model - Livepeer Docs](https://docs.livepeer.org/v2/gateways/concepts/business-model)
- [Livepeer Story - Livepeer Docs](https://docs.livepeer.org/v2/home/about-livepeer/vision)
- [About Livepeer: Protocol & Network - Livepeer Docs](https://docs.livepeer.org/v2/about/portal)
  - [Livepeer Overview - Livepeer Docs](https://docs.livepeer.org/v2/about/concepts/livepeer-overview)
  - [Protocol Overview - Livepeer Docs](https://docs.livepeer.org/v2/about/protocol/overview)
  - [Livepeer Network Overview - Livepeer Docs](https://docs.livepeer.org/v2/about/network/overview)

- Livepeer 技術・仕組みの解説
  - [Livepeerとは何ですか？どのように機能しますか？ | Gate Learn](https://www.gate.com/ja/learn/articles/what-is-livepeer/419)
  - [仮想通貨Livepeer(LPT)とは？所有するメリット・デメリットや将来性、特徴を解説！ - TokenNews｜仮想通貨ビットコインニュース・投資情報](https://tokennews.jp/livepeerlpts/)
- 仮想通貨・金融メディアによる解説
  - [ライブピア（LPT）の特徴と将来性、おすすめ取引所と購入方法を解説](https://coinpost.jp/?p=608128)
  - [仮想通貨LPT(ライブピア)とは？特徴や将来性、価格動向を徹底解説！ | CRYPTO INSIGHT powered by ダイヤモンド・ザイ](https://diamond.jp/crypto/market/lpt/)
  - [仮想通貨LPT(Live peer/ライブピア)とは？特徴と将来性、購入できる取引所を解説 | JinaCoin](https://jinacoin.ne.jp/lpt/)
- ブログ・ガイド記事
  - [【Web3時代の分散型動画配信インフラ】仮想通貨LPT（ライブピア）が築く「Livepeer Networkと動画ストリーミングの未来」とは？｜きままなヤース](https://note.com/gentle_gibbon604/n/n0721226ae7bf)
  - [仮想通貨LPT（Livepeer）完全ガイド：動画配信を変える仕組み・ステーキング・将来性  |  bitCurrent](https://crypto.appmatch.jp/4535-2/)
  - [Streamplaceの登場は、Livepeerエコシステムにおいて興味深いマイルストーンを示しています。 |Gate 広場のWeb3Ronin](https://www.gate.com/ja/post/status/17790209)

- [金融セクターの社会的無用さ » p2ptk[.]org](https://p2ptk.org/notes/5542)

[前回]: {{< ref "./gogh-streaming-on-streamplace.md" >}} "Streamplace で Gogh 作業配信を行う"
[gogh]: https://gogh.gg/ "gogh - Focus with Your Avatar"
[OBS]: https://obsproject.com/ "Open Broadcaster Software | OBS"
[Streamplace]: https://stream.place/ "Streamplace"
[atproto]: https://atproto.com/ "The AT Protocol"
[Livepeer]: https://livepeer.org/ "Livepeer — The open network for real-time AI video"
[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"
