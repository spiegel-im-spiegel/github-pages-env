+++
title = "2022-08-20 のブックマーク"
date =  "2022-08-20T10:54:13+09:00"
description = "「WebAssembly化したPostgreSQLをWebブラウザ上で実際に動かして学ぶ「Postgres playground」をCrunchy Dataが公開」他"
image = "/images/attention/bookmarks.jpg"
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

## リリース情報 {#release}

- [Announcing Rust 1.63.0 | Rust Blog](https://blog.rust-lang.org/2022/08/11/Rust-1.63.0.html)
- [Release 9.1.5 · mermaid-js/mermaid · GitHub](https://github.com/mermaid-js/mermaid/releases/tag/9.1.5)
- [Release v0.8.2 · nyaosorg/go-readline-ny · GitHub](https://github.com/nyaosorg/go-readline-ny/releases/tag/v0.8.2)
- [Release gopls/v0.9.4 · golang/tools · GitHub](https://github.com/golang/tools/releases/tag/gopls%2Fv0.9.4)
- [Release v2.0.3 · pelletier/go-toml · GitHub](https://github.com/pelletier/go-toml/releases/tag/v2.0.3)
- [Release 9.1.6 · mermaid-js/mermaid · GitHub](https://github.com/mermaid-js/mermaid/releases/tag/9.1.6)
- [LibreOffice 7.4 Community, a benchmark for interoperability - The Document Foundation Blog](https://blog.documentfoundation.org/blog/2022/08/18/libreoffice-7-4-community/)
  - [「LibreOffice 7.4 Community」が公開、スプレッドシートで16,384列まで扱えるように - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1433116.html)
- [「Thunderbird 91」から「Thunderbird 102」への手動アップデートが開始 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1432440.html)

## セキュリティ＆プライバシー関連 {#security}

- [The Hacking of Starlink Terminals Has Begun | WIRED](https://www.wired.com/story/starlink-internet-dish-hack/)
  - [Hacking Starlink - Schneier on Security](https://www.schneier.com/blog/archives/2022/08/hacking-starlink.html)
  - [ブログ: Starlink端末のハッキングが始まった](https://okuranagaimo.blogspot.com/2022/08/starlink.html)
- [Facebook begins testing default end-to-end encryption on Messenger - The Verge](https://www.theverge.com/2022/8/11/23301275/facebook-messenger-end-to-end-encryption-default-test-e2ee-privacy-security-plans) : Apple みたいに E2E 暗号化を勘違いしてなきゃいいけどね
- [Twitter Exposes Personal Information for 5.4 Million Accounts - Schneier on Security](https://www.schneier.com/blog/archives/2022/08/twitter-exposes-personal-information-for-5-4-million-accounts.html)
- [The new USB Rubber Ducky is more dangerous than ever - The Verge](https://www.theverge.com/23308394/usb-rubber-ducky-review-hack5-defcon-duckyscript)
  - [USB “Rubber Ducky” Attack Tool - Schneier on Security](https://www.schneier.com/blog/archives/2022/08/usb-rubber-ducky-attack-tool.html)
- [セキュリティ、プライバシーに無頓着すぎる保育アプリ業界 | p2ptk[.]org](https://p2ptk.org/privacy/3787)
- [「商業的監視」を標的に定めた米連邦取引委員会――政府による監視と企業による監視の密接な関係 | p2ptk[.]org](https://p2ptk.org/privacy/3798)
- [ジャネット・ジャクソンのヒット曲を再生すると古いHDDがクラッシュする脆弱性が話題に - GIGAZINE](https://gigazine.net/news/20220819-hdd-crash-janet-jackson-music/) : マジな[脆弱性](https://nvd.nist.gov/vuln/detail/CVE-2022-38392 "CVE-2022-38392")

## Go 言語関連 {#golang}

- [Goにおける型によってSQLインジェクションを防ぐ方法](https://zenn.dev/tenntenn/articles/b452911b4200ff) : リテラル文字列以外はエラーにするという荒業。面白い！
- [作ってわかる！ はじめてのgRPC](https://zenn.dev/hsaki/books/golang-grpc-starting)
- [Go で Stack と FIFO - Qiita](https://qiita.com/mattn/items/774280a746c82ee63fc0)
- [1.15kgで14時間駆動のUbuntu搭載モバイルノート  - PC Watch](https://pc.watch.impress.co.jp/docs/news/1423368.html)
- [なぜContextを構造体に含めてはいけないのか、またそれが許される状況について #golang - Qiita](https://qiita.com/sonatard/items/d97279086b24e588a82d) : `context.Context` を構造体を含めてもいい場合も解説
- [Go1.19で採用された Pattern-defeating Quicksort の紹介 - Speaker Deck](https://speakerdeck.com/po3rin/go1-dot-19decai-yong-sareta-pattern-defeating-quicksort-falseshao-jie)
- [gRPCurlでgRPCサーバの動作確認をする](https://zenn.dev/necocat/articles/e0a65f2da065ec)
- [gRPCがフロントエンド通信の第一の選択肢になる時代がやってきたかも？ | フューチャー技術ブログ](https://future-architect.github.io/articles/20220819a/)
- [[Go] ありそうで（あまり）なかったのでローマ字->ひらがな変換パッケージを作った](https://zenn.dev/usk81/articles/d95d2d50541a09)
- [インフラもバックエンドもフロントエンドも Go で書いてみた](https://zenn.dev/kou_pg_0131/articles/gogogo-introduction)

## Java  言語関連 {#java}

- [Release v2.7.3 · spring-projects/spring-boot · GitHub](https://github.com/spring-projects/spring-boot/releases/tag/v2.7.3)
- [Release v2.6.11 · spring-projects/spring-boot · GitHub](https://github.com/spring-projects/spring-boot/releases/tag/v2.6.11)

## その他 {#others}

- [ここさえ抑えればGitHub API v4がわかる! GraphQL入門](https://zenn.dev/hsaki/articles/github-graphql)
- [Steam Deckいよいよ日本上陸。5万9,800円から  - PC Watch](https://pc.watch.impress.co.jp/docs/news/1430013.html) : ゲーム専用タブレットと考えれば安いのかもしれんが，ふむむー
- [The US monkeypox response needs a robust data infrastructure. We don’t have it. - Vox](https://www.vox.com/future-perfect/2022/8/14/23302054/monkeypox-data-health-care-covid-collection)
- [Amazonの監視カメラ付きドアベル「Ring」の面白動画を集めるテレビ番組が制作決定、ディストピアな監視社会が加速するとメディアから拒否反応 - GIGAZINE](https://gigazine.net/news/20220815-ring-nation-amazon/)
- [JavaScript のクロージャーと for 文の let 初期化の例外](http://nmi.jp/2022-08-16-Let-Initializer-In-For-Loop)
- [バービー「『私の価値を決めるのは私』にグッときた」×ブレイディみかこ「貧しいから、女の子だから、守ってあげたくなるなんて描き方には絶対したくなかった」 | ダ・ヴィンチWeb](https://ddnavi.com/interview/1018707/a/)
- [Denoが大幅な方針変更を発表。3カ月以内にnpmパッケージへの対応を実現、最速のJavaScriptランタイムを目指しHTTPサーバを刷新 － Publickey](https://www.publickey1.jp/blog/22/deno3npmjavascripthttp.html)
- [WebAssembly化したPostgreSQLをWebブラウザ上で実際に動かして学ぶ「Postgres playground」をCrunchy Dataが公開 － Publickey](https://www.publickey1.jp/blog/22/webassemblypostgresqlwebpostgres_playgroundcrunchy_data.html)
- [Linking to news doesn’t make Google liable for defamation, Australia court rules | Ars Technica](https://arstechnica.com/tech-policy/2022/08/linking-to-news-doesnt-make-google-liable-for-defamation-australia-court-rules/)
- [PDFを「Excel」スプレッドシートに変換するには - ZDNet Japan](https://japan.zdnet.com/article/35191765/)
- [Centralization, Decentralization, and Internet Standards](https://www.ietf.org/archive/id/draft-nottingham-avoiding-internet-centralization-05.html)
  - [ブログ: 中央集権化、非中央集権化、インターネット標準](https://okuranagaimo.blogspot.com/2022/08/blog-post_14.html)
