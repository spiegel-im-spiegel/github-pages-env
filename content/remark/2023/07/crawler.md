+++
title = "Mastodon と Bluesky でボット運用はじめました【2023-12-04 更新】"
date =  "2023-07-01T12:49:59+09:00"
description = "自作ツールに Web 上の情報を収集する機能を付けて Mastodon/Bluesky 上でボットを構成できるようにした。"
image = "/images/attention/kitten.jpg"
tags = [ "communication", "bluesky", "mastodon", "tools", "golang", "engineering" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

## 自作ツール goark/toolbox

5月頃に思いついてコマンドラインで Mastodon や Bluesky に投稿できる [goark/toolbox] というツールを作った。
ついでに Web 上の情報を収集する機能も付けてボットを構成できるようにした。

実際の運用は自宅 PC で cron を回している（自宅 PC は24時間稼働中）。
[Go] ならシングルバイナリで取り回しできるし，この程度ならクラウドとか Docker とか要らんですよ。

## Bluesky で非公式 APOD 配信ボットを作った

Mastodon には [APOD (Astronomy Picture of the Day)](https://apod.nasa.gov/apod/ "Astronomy Picture of the Day") の非公式配信ボットがいくつかあるのだが（`#apod` で検索するとアホほど出てくる），できたばかりの Bluesky で運用している人はいない様子。
なら作っちゃえ！ というわけで作った。

- [@apodunofficial.bsky.social](https://bsky.app/profile/apodunofficial.bsky.social "Astronmy Picture of the Day (unofficial bot)")

データを収集するために [NASA API](https://api.nasa.gov/ "NASA Open APIs") をハンドリングする機能を [goark/toolbox] に組み込んだ。

- [NASA API を使って “Astronomy Picture of the Day” のデータを取得する]({{< ref "/remark/2023/02/api-for-astronomy-picture-of-the-day.md" >}})

API でクレジットが示されているものは明示しているが， NASA 関連の画像・動画については示されない？

あと API がしょっちゅう 504 で落ちるのね。
Web ページは全然構造化されてなくてスクレイピングする気も起こらない。
というわけで，生暖かく見守っていただければ（笑）

## 自ブログおよび自作パッケージの更新情報も配信

Bluesky や Mastodon に記事を投稿する仕組みが整ったので，フィードを読み込む機能も組み込んで自ブログや Flickr にアップした写真の更新情報を以下の Mastodon/Bluesky のメインアカウントに配信できるようにした。

- [@spiegel@fedibird.com ](https://fedibird.com/@spiegel "Spiegel@がんばらない")
- [@baldanders.info](https://bsky.app/profile/baldanders.info "Spiegel")

ちなみに [Go] でフィードを取得するには [mmcdole/gofeed] パッケージを使うのが便利。

- [フィードを取得する Go 言語パッケージ](https://zenn.dev/spiegel/articles/20201003-feed-with-golang)

これでフィードを自動配信する仕組みも用意できたので，自ブログ以外に [github.com/goark](https://github.com/goark "Playing with Go Language") で公開している自作パッケージの更新情報も自動投稿することにした。

GitHub のリリース情報のフィードは以下の URL で取得できる。

```html
https://github.com/username/repositoryname/releases.atom
```

自作パッケージの更新情報は以下のアカウントで自動投稿している。

- [@goark@goark.fedicity.net](https://goark.fedicity.net/@goark "Goark (@goark@goark.fedicity.net) - Goark")
- [@goark.bsky.social](https://bsky.app/profile/goark.bsky.social "Goark")

[goark.fedicity.net](https://goark.fedicity.net/ "Goark") については「[個人用 Mastodon サーバを立てた]({{< ref "/remark/2023/12/personal-mastodon-server.md" >}})」を参考にどうぞ。

[@goark.bsky.social](https://bsky.app/profile/goark.bsky.social "Goark") は [goark/toolbox] の動作テスト用に取ったアカウントだけど，遊ばせておくのもナニなので半ボットとして運用することにした。
まぁ，今後もテスト用にゴミ投稿することもあると思うけど，そこはご容赦を（笑）

## その他のフィードも配信するぞ

調子に乗って自作以外でお世話になっている [Go] パッケージ（プロダクト）や [Go] 関連記事も以下のアカウントに自動投稿することにした。

- [@goark@goark.fedicity.net](https://goark.fedicity.net/@goark "Goark (@goark@goark.fedicity.net) - Goark")
- [@osanpo.bsky.social](https://bsky.app/profile/osanpo.bsky.social "Spiegel's crawler")

[@osanpo.bsky.social](https://bsky.app/profile/osanpo.bsky.social "Spiegel's crawler") は Bluesky のボット運用のために取ったアカウント。

監視対象は以下のサイト（今後追加予定）：

- [The Go Blog - The Go Programming Language](https://go.dev/blog/)
- [golang/tools: [mirror] Go Tools](https://github.com/golang/tools)
- [go-task/task: A task runner / simpler Make alternative written in Go](https://github.com/go-task/task)
- [goreleaser/goreleaser: Deliver Go binaries as fast and easily as possible](https://github.com/goreleaser/goreleaser)
- [sashabaranov/go-openai: OpenAI ChatGPT, GPT-3, GPT-4, DALL·E, Whisper API wrapper for Go](https://github.com/sashabaranov/go-openai)
- [anchore/syft: CLI tool and library for generating a Software Bill of Materials from container images and filesystems](https://github.com/anchore/syft)
- [golangci/golangci-lint: Fast linters Runner for Go](https://github.com/golangci/golangci-lint)
- [tinygo-org/tinygo: Go compiler for small places. Microcontrollers, WebAssembly (WASM/WASI), and command-line tools. Based on LLVM.](https://github.com/tinygo-org/tinygo)
- [pelletier/go-toml: Go library for the TOML file format](https://github.com/pelletier/go-toml)
- [nyaosorg/go-readline-ny: Readline library for golang , used in nyagos](https://github.com/nyaosorg/go-readline-ny)
- [hymkor/go-multiline-ny: Readline package supporting multi-lines](https://github.com/hymkor/go-multiline-ny)
- [rs/zerolog: Zero Allocation JSON Logger](https://github.com/rs/zerolog)
- [uber-go/zap: Blazing fast, structured, leveled logging in Go.](https://github.com/uber-go/zap)

さらに調子に乗って [Go] 関連以外にも私の関心領域にかかる情報も収集・投稿することにした。
投稿先は以下の通り：

- [@osanpo@goark.fedicity.net](https://goark.fedicity.net/@osanpo "Spiegel's Crawler (@osanpo@goark.fedicity.net) - Goark")
- [@osanpo.bsky.social](https://bsky.app/profile/osanpo.bsky.social "Spiegel's crawler")

監視対象は以下のサイト（今後追加予定）：

- [Schneier on Security](https://www.schneier.com/)
- [情報セキュリティ | IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/index.html)
- [JPCERT コーディネーションセンター](https://www.jpcert.or.jp/)
- [piyolog](https://piyolog.hatenadiary.jp/)
- [Security notices | Ubuntu](https://ubuntu.com/security/notices)

- [National Institute of Standards and Technology](https://www.nist.gov/)
- [P2Pとかその辺のお話R | Sharing is Caring](https://p2ptk.org/)
- [Blog - Creative Commons](https://creativecommons.org/blog/)
- [Flickr Foundation](https://www.flickr.org/)

- [国立天文台（NAOJ）](https://www.nao.ac.jp/)
- [国立天文台 天文情報センター 暦計算室](https://eco.mtk.nao.ac.jp/koyomi/)

- [Publickey － Enterprise IT × Cloud Computing × Web Technology / Blog](https://www.publickey1.jp/)
- [MathJax | Beautiful math in all browsers.](https://www.mathjax.org/)
- [The Thunderbird Blog -](https://blog.thunderbird.net/)
- [Big Sky](https://mattn.kaoriya.net/)

- [nyaosorg/nyagos: NYAGOS - The hybrid Commandline Shell between UNIX & DOS](https://github.com/nyaosorg/nyagos)
- [gohugoio/hugo: The world’s fastest framework for building websites.](https://github.com/gohugoio/hugo)
- [tailscale/tailscale: The easiest, most secure way to use WireGuard and 2FA.](https://github.com/tailscale/tailscale)
- [koki-develop/gat: 🐱 cat alternative written in Go.](https://github.com/koki-develop/gat)

- [mermaid-js/mermaid: Generation of diagrams like flowcharts or sequence diagrams from text in a similar manner as markdown](https://github.com/mermaid-js/mermaid)
- [microsoft/vscode: Visual Studio Code](https://github.com/microsoft/vscode)
- [plantuml/plantuml: Generate diagrams from textual description](https://github.com/plantuml/plantuml)
- [spring-projects/spring-boot: Spring Boot](https://github.com/spring-projects/spring-boot)
- [keepassxreboot/keepassxc: KeePassXC is a cross-platform community-driven port of the Windows application “Keepass Password Safe”.](https://github.com/keepassxreboot/keepassxc)

上のリスト（の一部）は元々 [Slack 上で監視]({{< ref "/remark/2017/01/slack.md" >}} "いまさら聞けない Slack の使い方")していたものだが，プライベートでも仕事でもあまり Slack を使わなくなったので（仕事では主に Teams） Mastodon/Bluesky に移行することにしたのだ。
ただ，今までの反省で，やたら滅多とフィードを食い散らかして自 TL を埋め尽くすのは嬉しくないので，流量についてはチューニングしながら運用する予定である。

本来はフィードのチェックは Feedly とか使うべきだし，本当によく見るものだけに厳選したい。
特に青空文庫（[@aozorabunko.bsky.social](https://bsky.app/profile/aozorabunko.bsky.social "青空文庫(Aozora Bunko)")）とか，自身でブログ記事の更新情報を積極的に上げておられるユーザとかのフィードは外している。

## 今後は...

今の [goark/toolbox] ってフィード情報をキャッシュするのにテキストファイルを使ってるんだよな。
これだとそろそろ耐えられない気がするので SQLite を導入するかなぁ。

Pure [Go] で実装するならこのあたりか？

- [glebarez/go-sqlite: pure-Go SQLite driver for Go (SQLite embedded)](https://github.com/glebarez/go-sqlite)
- [glebarez/sqlite: The pure-Go SQLite driver for GORM](https://github.com/glebarez/sqlite)
- [moul/zapgorm2: ⚡ zap logging driver for gorm v2](https://github.com/moul/zapgorm2)
- [mpalmer/gorm-zerolog: Alternative logging with Zerolog for GORM ⚡️](https://github.com/mpalmer/gorm-zerolog)

まぁ，ぼちぼちやろう。

Twitter ？ 知らんですよ。
IFTTT にも金を払う気はないし，向こうでの活動は最小限に留めたい。

## ブックマーク

- [RSS/Atom フィードを Twitter/Nostr にポストする小さいプログラムを書いた。](https://zenn.dev/mattn/articles/7ac25f3328bde3)

[Go]: https://go.dev/
[goark/toolbox]: https://github.com/goark/toolbox "goark/toolbox: A collection of miscellaneous commands"
[mmcdole/gofeed]: https://github.com/mmcdole/gofeed "mmcdole/gofeed: Parse RSS, Atom and JSON feeds in Go"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
