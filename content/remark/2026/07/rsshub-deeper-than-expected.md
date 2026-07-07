+++
title = "RSSHub について調べようとしたら結構おおきな話だった"
date =  "2026-07-07T10:52:33+09:00"
description = "RSSHub と Folo / RSS3 / Web3 はちゃんと “Web” になっている？"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "feed", "web", "blockchain", "market", "activitypub", "nostr" ]
pageType = "text"
draft = true

[scripts]
  mathjax = false
  mermaidjs = true
  jsx = false
+++

いつものように Mastodon の自 TL に流れてきたポストから。
EFF (Electronic Frontier Foundation) の記事ですな。

- [Hate “The Algorithm?” RSS Is One of the Tools You’ve Been Looking For](https://www.eff.org/deeplinks/2026/06/hate-algorithm-rss-one-tools-youve-been-looking)

もの凄く端折っていうと，ソーシャルメディア等のアルゴリズムに任せるんじゃなくて，枯れた技術とはいえ RSS があるんだから，これを使って自分で情報管理しようぜ！ といった内容。

{{< fig-quote type="markdown" title="Hate “The Algorithm?” RSS Is One of the Tools You’ve Been Looking For" link="https://www.eff.org/deeplinks/2026/06/hate-algorithm-rss-one-tools-youve-been-looking" lang="en" >}}
RSS is one of the best examples we have of the open web, where we can design and customize how we experience the internet, not the other way around. RSS has come in and out of fashion, been declared dead, and has come back, every time. Open systems are the best way forward to a free, equitable internet, and the resilience and continued reinvention of RSS has shown just how creative the web community can be with open protocols.
{{< /fig-quote >}}

この記事の最後の方に

{{< fig-quote type="markdown" title="Hate “The Algorithm?” RSS Is One of the Tools You’ve Been Looking For" link="https://www.eff.org/deeplinks/2026/06/hate-algorithm-rss-one-tools-youve-been-looking" lang="en" >}}
And even if a website doesn’t have an RSS feed, there are workarounds from tools like [RSSHub](https://docs.rsshub.app/), [RSS-Bridge](https://github.com/rss-bridge/rss-bridge), and [RSS.app](https://rss.app/) that require varying levels of technical expertise or a willingness to pay subscription fees.
{{< /fig-quote >}}

などと書かれていて，「へぇ。フィードがなくても生成してくれるんだ」と思いながら，とりあえず [RSSHub のサイト][RSSHub]を覗いてみたのだが，何が書いてあるかよく分からない。
Folo？ RSSHub の話じゃねーの？ 自前で実装するのか？

というわけで，いつものように [Kagi Assistant] とやり取りしながら調べてみた。

## RSSHub と Folo

[RSSHub] は「Everything is RSSible（全ては RSS 化できる）」を掲げる「世界最大のRSSネットワーク」らしい。
[Kagi Assistant] にまとめさせた特徴は以下の通り。

{{< div-ai type="markdown" >}}
| 特徴 | 内容 |
|------|------|
| **オープンソース** | AGPL-3.0ライセンスで公開，オープンスタンダードとプロトコルに準拠 |
| **大規模コミュニティ** | 1,300人以上のコントリビューター・メンテナーが活躍 |
| **豊富なルート** | 多数の事前設定済みルート（変換ルール）とパブリックインスタンスが利用可能 |
| **強力なAPI** | 様々なシーンに対応するAPIとエコシステムプロジェクトを提供 |
{{< /div-ai >}}

Docker を使って環境を自前で構築できる他，公式を含む[パブリックインスタンス](https://docs.rsshub.app/guide/instances "Public Instances | RSSHub")も提供されている。
このインスタンスのひとつが [Folo] ということのようだ。

{{< fig-quote type="markdown" title="Public Instances" link="https://docs.rsshub.app/guide/instances" lang="en" >}}
[Folo](https://folo.is/) is an AI RSS reader and discovery platform that works especially well with RSSHub. It offers numerous user-shared instances that you can browse, switch between, and use with just one click inside Folo. These instances are not compatible with external readers. You can also share your own instance on Folo to earn tokens.
{{< /fig-quote >}}

「自身のインスタンスを Folo 上で共有してトークンを稼ぐ」？ とりあえず，そこは後回しだ。

### 直接 RSS フィードを提供する場合

まずは一番単純なパターン。
あらかじめ RSS フィードを提供している Web サイトから Folo に直接フィードを提供する場合。

{{< div-ai type="markdown" >}}
{{< fig-mermaid >}}
graph TB
    subgraph "WWW (Site)"
        RSS_YES[<strong>Web2 Site</strong><br/>RSSフィードあり<br/>ブログ，ニュースサイト等]
    end
    subgraph "Folo"
        FOLO[<strong>Folo</strong><br/>RSSリーダー]
    end
    %% RSSフィードを直接 FOLO に提供する
    RSS_YES -.->|RSSフィード提供| FOLO
{{< /fig-mermaid >}}
{{< /div-ai >}}

これは分かりやすい。

### RSSHub サーバ経由で RSS フィードを提供する場合

古き良き（笑）テキストサイト等や Web 2.0 時代のサービスでも RSS フィードを提供していない SNS サービス等は [RSSHub] サーバ経由でフィードを生成する。

{{< div-ai type="markdown" >}}
{{< fig-mermaid >}}
graph TB
    subgraph "WWW (Site)"
        RSS_NO[<strong>Web1/Web2 Site</strong><br/>RSSフィードなし<br/>静的サイト，SNS等]
    end
    RH_STANDALONE[<strong>RSSHub</strong><br/>独立サーバー]
    subgraph "Folo"
        FOLO[<strong>Folo</strong><br/>RSSリーダー]
    end
    %% 独立サーバー（RSS3不要）
    RSS_NO -.->|スクレイピング/独自API| RH_STANDALONE
    RH_STANDALONE -.->|RSSフィード提供| FOLO
{{< /fig-mermaid >}}
{{< /div-ai >}}

これもまだ分かる。
[RSSHub] が頑張って仲立ちしてるってことだよね。

## RSS3

ここからだよ，{{% ruby "おおごと" %}}大事{{% /ruby %}}なのは。

{{< div-ai type="markdown" >}}
{{< fig-mermaid >}}
graph TB
    subgraph "WWW (Site)"
        RSS_YES[<strong>Web2 Site</strong><br/>RSSフィードあり<br/>ブログ等]
        RSS_NO[<strong>Web1/Web2 Site</strong><br/>RSSフィードなし<br/>静的サイト，SNS等]
    end
    subgraph "Fediverse"
        AP[<strong>ActivityPub</strong><br/>Mastodon等]
    end
    subgraph "Web3 (Blockchain)"
        ETH[<strong>Ethereum</strong><br/>トランザクション]
        NFT[<strong>NFT</strong><br/>ERC-721]
        DEFICH[<strong>DeFi</strong><br/>Uniswap/Aave]
        ARWEAVE[<strong>Arweave</strong><br/>分散型ストレージ]
    end
    subgraph "RSS3 Node"
        direction TB
        subgraph "RSS Component"
            RH_INTERNAL[<strong>RSSHub</strong><br/>内部コンポーネント]
        end
        subgraph "Federated Component"
            FED[<strong>ActivityPub</strong><br/>インデクサー]
        end
        subgraph "Decentralized Component"
            DEC[<strong>Blockchain</strong><br/>インデクサー]
        end
        DSL[<strong>DSL</strong><br/>Data Sublayer<br/>インデックス・保存]
    end
    subgraph "RSS3 Network"
        NETWORK[<strong>RSS3 Network</strong><br/>分散型ノード群]
    end
    subgraph "Folo"
        FOLO[<strong>Folo</strong><br/>RSSリーダー]
    end
    %% Web2 → RSSComponent
    RSS_YES -->|RSSフィード取得| RH_INTERNAL
    RSS_NO -->|スクレイピング/API| RH_INTERNAL
    RH_INTERNAL -->|フィードデータ| DSL
    %% Fediverse → FederatedComponent
    AP -->|ActivityPubデータ| FED
    FED -->|インデックス| DSL
    %% Web3 → DecentralizedComponent
    ETH -->|ブロックチェーンデータ| DEC
    NFT -->|NFT取引データ| DEC
    DEFICH -->|DeFiトランザクション| DEC
    ARWEAVE -->|分散型データ| DEC
    DEC -->|インデックス| DSL
    %% DSL → ネットワーク → Folo
    DSL -->|統合データ| NETWORK
    NETWORK -->|Web1/2/3統合| FOLO
{{< /fig-mermaid >}}
{{< /div-ai >}}

なんと [RSS3] ノードを経由することで Fediverse (ActivityPub) や Web3 のデータも [Folo] で閲覧できるらしい。
ActivityPub とか Blockchain とかしれっと出てきて，私は驚いたよ。
世の中そんな風になってるんだねぇ。

[RSS3] は “An open web, readable again” をビジョンとして（開かれたウェブですってよ，奥さん），「Web1/Web2/Web3 にまたがる情報をインデックス・配信する分散型情報インフラ」だそうだ。
分散ネットワークで検閲耐性が高いと言われている[^dc1]。
[RSS3] のエコシステムにおいて [RSSHub] および [Folo] は中核プロダクトと見なされているらしい[^rh1]。

[^dc1]: P2P など，本当に分散化したシステムではデータ（コンテンツ）の「場所」が不定となり遍在化するため検閲が難しくなる。コピーはいくらでもあるし，アクセスの迂回も容易だからだ。
[^rh1]: 2024年1月に [RSSHub] と [RSS3] との間でパートナーシップが結ばれた，と [Kagi Assistant] が言っていたが，その説明の Web ページが 404 でなくなってるみたいで，経緯とかは確認できなかった。AI の知識にはあるけど裏付ける情報がない，ということのようだ。これからこういうのも増えそうだな。

### RSS3 と $RSS3

[RSS3] 経済圏の通貨が [$RSS3] トークンで，以下のように機能しているそうな。

{{< div-ai type="markdown" >}}
{{< fig-mermaid >}}
graph TB
    subgraph "アプリケーション層"
        APP[アプリケーション開発者]
    end
    subgraph "ネットワーク層"
        NODE[ノード運営者]
        DELEGATOR[トークン保有者<br/>委任者]
    end
    subgraph "インフラ層"
        DSL[DSL<br/>Data Sublayer<br/>情報インデックス]
        VSL[VSL<br/>Value Sublayer<br/>Ethereum L2]
    end
    subgraph "トークンエコノミー"
        TOKEN[$RSS3 トークン]
        POOL[Operation Pools]
    end
    %% トークンの流れ
    APP -->|"① API利用料支払い"| TOKEN
    TOKEN -->|"② クエリ料金分配"| POOL
    POOL -->|"③ 報酬分配"| NODE
    POOL -->|"④ 委任報酬"| DELEGATOR
    %% データの流れ
    NODE -->|"データインデックス"| DSL
    DSL -->|"検証済みデータ"| VSL
    VSL -->|"チェーン上データ"| APP
    %% 委任関係
    DELEGATOR -->|"トークン委任"| NODE
{{< /fig-mermaid >}}

| 番号 | フロー | 説明 |
|------|--------|------|
| **①** | アプリケーション → トークン | 開発者がAPIキーを取得し $RSS3 でクエリ料を支払う |
| **②** | トークン → Operation Pools | 支払いがプールに集約される |
| **③** | プール → ノード運営者 | データインデックス作業の報酬として分配 |
| **④** | プール → 委任者 | ノードに委任したトークン保有者にも報酬分配 |
{{< /div-ai >}}

これ，実際にデータを使うエンドユーザはエコシステムの外側にいるし，ユーザから見て直接データを提供するアプリケーション（開発者）も，この流通サイクルの中では [$RSS3] トークンを（払うだけで）得る手段がないことになる。
トークンがこのエコシステムの外側でも流通すればいいんだろうけど（以下は [Kagi Assistant] が示した [RSS3] エコシステムの外側との関係）。

{{< div-ai type="markdown" >}}
{{< fig-mermaid >}}
graph TB
    subgraph "RSS3 の役割"
        RSS3[RSS3<br/>分散型情報インフラ]
        API[RSS3 API<br/>データ提供]
    end
    subgraph "取引所の活用"
        OKX[OKX<br/>取引ボット，ポートフォリオ管理]
        ARB[Arbitrum<br/>クロスチェーンエクスプローラー]
        AI[AI DApps<br/>分散型エージェント]
    end
    subgraph "$RSS3 トークン"
        TOKEN[$RSS3<br/>取引所で取引可能]
        BITGET[Bitget]
        LBANK[LBank]
        COINW[CoinW]
    end
    RSS3 -->|データ提供| API
    API --> OKX
    API --> ARB
    API --> AI
    TOKEN --> BITGET
    TOKEN --> LBANK
    TOKEN --> COINW
{{< /fig-mermaid >}}
{{< /div-ai >}}

この手のネットワーク通貨の問題って，結局ゼロ年代に流行った「[地域通貨](https://cruel.org/krugman/babysitj.html "経済を子守りしてみると。")」の問題と同じなんだよなぁ。 ...まぁ，いいや。
ちなみに，一定額の [$RSS3] トークンを Staking[^s1] すれば誰でも [RSS3] ノードを運用できるらしい。

[^s1]: Staking とは，トークンを（ウォレットとかではなく） Blockchain ネットワーク上に一定期間預けておくことを指す。 Ethereum のコンセンサスアルゴリズムである Proof of Stake (PoS) はトークンを多く長く預けている人ほどトランザクション承認の権利が得やすくなるため，結果として信用度や貢献度が高いと見なされ，報酬も得やすくなる。簡単にいうと金遣いが荒い人は信用されないってことだよね。この Staking の仕組みを使って投資を行う人も結構いるらしい。

### 分散型ソーシャルプロトコル統合

[RSS3] は複数の分散型ソーシャルプロトコルを Universal Social API で統合しているらしい。

{{< div-ai type="markdown" >}}
{{< fig-mermaid >}}
graph TB
    subgraph "分散型ソーシャルプロトコル"
        FC[<strong>Farcaster</strong><br/>Ethereumベース]
        LENS[<strong>Lens Protocol</strong><br/>Polygonベース]
        NOSTR[<strong>Nostr</strong><br/>シンプルプロトコル]
        MIRROR[<strong>Mirror</strong><br/>Web3パブリッシング]
        CB[<strong>Crossbell</strong><br/>ソーシャルプロトコル]
    end
    subgraph "RSS3"
        RSS3api[<strong>RSS3</strong><br/>Universal Social API]
        API[<strong>RSS3 API</strong><br/>データ統合]
    end
    subgraph "アプリケーション"
        HOOT[<strong>Hoot.it</strong><br/>ソーシャル検索]
        FOLO[<strong>Folo</strong><br/>RSSリーダー]
        DEV[<strong>開発者アプリ</strong>]
    end
    FC --> RSS3api
    LENS --> RSS3api
    NOSTR --> RSS3api
    MIRROR --> RSS3api
    CB --> RSS3api
    RSS3api --> API
    API --> HOOT
    API --> FOLO
    API --> DEV
{{< /fig-mermaid >}}

**各プロトコルの特徴**

| プロトコル | ブロックチェーン | 特徴 |
|-----------|----------------|------|
| **Farcaster** | Ethereum | 「十分な分散化」，柔軟なネームスペース，グローバル状態 |
| **Lens Protocol** | Polygon | Aave創設者Stani Kulechovによる，合成可能なソーシャルグラフ |
| **Nostr** | なし（リレー方式） | シンプル，検閲耐性，Jack Dorseyが支援 |
| **Mirror** | Ethereum | Web3パブリッシングプラットフォーム |
| **Crossbell** | - | ソーシャルプロトコル |

**具体的な機能**

| 機能 | 説明 |
|------|------|
| **クロスプロトコル検索** | Farcaster，Lens，Nostr のデータを横断して検索 |
| **ソーシャルグラフ統合** | 異なるプロトコルのユーザー関係を統合 |
| **コンテンツ相互運用性** | プロトコル間のコンテンツ相互利用を可能にする |
{{< /div-ai >}}

Nostr もここにかかってくるんだねぇ。
これもびっくりした。
ただ Blockchain ベースのプロトコルは分かるけど Nostr は（Blockchain ではないので）どうやって統合してるのか分からない。
[Kagi Assistant] は色々憶測を述べてたけど（RSS Component 経由で RSS フィードとして取り込まれるんじゃないか，とか言ってくさる），実際の技術詳細は公開されてないらしい。

アプリケーション例のうち Hoot.it は「[RSS3] が開発したオープン Web 検索エンジン」だそうで Nostr データの閲覧も可能なんだと。

{{< div-ai type="markdown" >}}
**Hoot.it の機能**

| 機能 | 説明 |
|------|------|
| **Nostr アドレス検索** | Nostr アドレスを入力してソーシャルデータを表示 |
| **ChatGPT 統合** | 2024年5月に ChatGPT との接続を発表 |
| **Web3 コンテンツ検索** | 分散型ネットワーク上のコンテンツを検索 |
{{< /div-ai >}}

ActivityPub が Federated Component で，Nostr が Universal Social API で [RSS3] に統合されてるとなると [atproto] が取り残されてる感じかねぇ。
でも，以前に[紹介]({{< ref "/remark/2026/04/streamplace-and-livepeer.md" >}} "Streamplace と Livepeer")した Streamplace/Livepeer なんかは独自のトークンを運用してるって話だし，またぞろ面倒くさいことになってる気がするなぁ。

## Web3 はちゃんと “Web” になっている？

Blockchain 周りは NFT が登場したあたりから完全に興味が失せてしまったのだけど，現時点においてちゃんと “Web” としての体裁を整えてきているという印象を受けた。
そこに私のようなネットの辺境にいるような人間が入り込めるかは分からないけど [Folo] のようなサービスはちゃんと触っておいたほうがいい気がする。

まぁ，その辺はおいおい（笑）

[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"
[RSSHub]: https://docs.rsshub.app/ "RSSHub"
[Folo]: https://folo.is/ "Folo — AI RSS Reader for deep, noise-free reading with contextual AI."
[RSS3]: https://rss3.io/ "RSS3"
[$RSS3]: https://rss3.io/blog/about-rss3.html "About $RSS3"
[atproto]: https://atproto.com/ "The AT Protocol"

## 参考

{{% review-tatsujin "infoshare" %}} <!-- 情報共有の未来 -->
{{% review-tatsujin "infoshare2" %}} <!-- 続・情報共有の未来 -->
