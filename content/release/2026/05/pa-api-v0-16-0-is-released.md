+++
title = "Amazon PA-API 改め Creators API v0.16.0 をリリース"
date =  "2026-05-06T12:39:52+09:00"
description = "PA-APIv5 から Amazon Creators API への移行，型変更，バリデーション強化について。"
isCJKLanguage = true
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "package", "module", "pa-api", "amazon" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

現行の Amazon Product Advertising API (PA-API) v5 が 2026-05-15 で廃止になるのに伴い [Creators API] への移行が必要なのだが，拙作の [`goark/pa-api`] パッケージのユーザの方に素晴らしい [pull request](https://github.com/goark/pa-api/pull/47 "Migrate PAAPI v5 to Amazon Creators API by Oakes6 · Pull Request #47 · goark/pa-api") をいただいた。
これをマージして v0.16.0 としてリリースした。

- [Release v0.16.0 · goark/pa-api · GitHub](https://github.com/goark/pa-api/releases/tag/v0.16.0)

v1.x ならメジャーバージョンを上げなきゃいけないところだったけど v0 系なのでマイナーバージョンを上げるだけにした。
まだ v1 に上げる勇気がない（笑）

パッケージ名がアレなのだが，パッケージ名を変えようとすると別リポジトリになってしまうので，README のサンプルコードに

```go
import creatorsapi "github.com/goark/pa-api"
```

と書いてお茶を濁している。

完全に言い訳なのだが，私は今 Amazon Associate の API が使えないのよ。
なんか月間で10アイテム以上売り上げないと API が使えないそうで，うちみたいにネットの端っこで細々と運営しているブログサイトじゃ無理だっての。
Amazon Associate の利用目的だって，単に書影を合法的に使いたいからという，超消極的なものだし。

というわけで [Creators API] の対応はほとんど諦めてたんだけど，まさか PR を貰えるとは。
今回は貰った PR をそのままマージして GitHub Copilot にレビュー&修正を指示して[^pr1]， README とかのドキュメントも Copilot に直させて。
私は1行もコードを書いてない。

[^pr1]: ちなみに PR に対して Copilot にレビューさせてみたんだけど，どうにも「分かってない」感じがして丸無視してしまった。今後もあるから `copilot-instructions.md` ファイルを作った。これで次回以降はもう少しまともなレビューにならないかなぁ。

ホンマこのパッケージは，文字通り皆さんの貢献（contributions）のおかげで成り立ってると実感できる。
今後ともよろしくお願いします {{% emoji "ペコン" %}}

## 利用者向けメモ

README にも書いてるけど，移行に関するメモを記しておく。

{{< fig-quote type="markdown" title="移行チェックリスト" link="https://github.com/goark/pa-api#quick-migration-checklist" >}}
既存の PA-API v5 呼び出し箇所を移行する際は，この手順に従う:

1. インポートパスは `github.com/goark/pa-api` のままにする（必要であれば，例では `creatorsapi` としてエイリアス可）
1. クライアント認証情報を置き換える: AWS Access Key / Secret Key → Creators API Credential ID / Credential Secret
1. マーケットプレイスはサーバー／クライアント側で設定する（`WithMarketplace`）。クエリごとのマーケットプレイス引数には依存しないこと
1. V1 の `offers` 利用を `OffersV2` に置き換える（`EnableOffersV2`。`EnableOffers` は互換エイリアスとして残している）
1. `Merchant`，`OfferCount`，`PartnerType` が効くという前提は外すこと。これらは無視される
1. Credential Version（2.1/2.2/2.3）が呼び出すマーケットプレイスグループと一致していることを確認する
1. 429 および一時的な 5xx 応答に対して再試行とレート制御を追加する
1. PR を作成する前に，プロジェクトの標準的なテスト／リントのワークフローでローカル検証を実行すること
{{< /fig-quote >}}

## ブックマーク

- [Creators API](https://affiliate-program.amazon.com/creatorsapi/docs/en-us/introduction)
  - [Migrating to Creators API from Product Advertising API](https://affiliate-program.amazon.com/creatorsapi/docs/en-us/migrating-to-creatorsapi-from-paapi)
- [Associates Central — Creators API portal](https://affiliate-program.amazon.com/creatorsapi)

[Creators API]: https://affiliate-program.amazon.com/creatorsapi/docs/en-us/introduction "Creators API"
[`goark/pa-api`]: https://github.com/goark/pa-api "goark/pa-api: Go client for the Amazon Creators API"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B0CFL1DK8Q" %}} <!-- Go言語 100Tips -->
{{% review-paapi "B0DNYMMBBQ" %}} <!-- Go言語で学ぶ並行プログラミング -->
