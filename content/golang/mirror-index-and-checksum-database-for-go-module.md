+++
title = "Go モジュールのミラーリング・サービス【正式版】"
date =  "2019-09-01T19:39:57+09:00"
description = "Google による Go モジュールの公式ミラーリング・サービスが正式リリースした。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "module", "engineering", "management" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

（この記事は「[Go モジュールのミラーリングとインデックス化]({{< ref "/release/2019/05/mirror-index-and-checksum-database-for-go-module.md" >}})」を全面改訂したものです）

Google による [Go] モジュールの公式ミラーリング・サービスが正式リリースした。

- [Module Mirror and Checksum Database Launched - The Go Blog](https://blog.golang.org/module-mirror-launch)
    - [`proxy.golang.org`]
    - [`sum.golang.org`]
    - [`index.golang.org`]

[`proxy.golang.org`]: https://proxy.golang.org/
[`sum.golang.org`]: https://sum.golang.org/
[`index.golang.org`]: https://index.golang.org/

## モジュール・ミラー

ミラー・サーバ [`proxy.golang.org`] は [Go 言語]のモジュールをミラーリングする一種のプロキシ・サーバとして機能する。
ただし，単純な透過プロキシではなく必要なバージョンのみ取ってこれるようになっているようだ。

{{< fig-img src="https://blog.golang.org/module-mirror-launch/proxy-protocol.png" title="Module Mirror and Checksum Database Launched" link="https://blog.golang.org/module-mirror-launch" width="2106" >}}

ミラー・サーバは環境変数 `GOPROXY` で URL を指定する。
`GOPROXY` の既定値は [Go] 1.13 では以下のようになっている。

```text
$ go env | grep GOPROXY
GOPROXY="https://proxy.golang.org,direct"
```

ミラー・サーバは複数指定できる（指定する場合はカンマで区切る）。
ミラーリングを無効にする場合は `GOPROXY` に `direct` のみを指定する。

```text
go env -w GOPROXY=direct
```

ミラーリングが有効な状態で特定の非公開モジュールを使う場合は `GOPRIVATE` に非公開モジュールのパスを指定する。

## チェックサム・データベース

[Go 言語]では利用するモジュールの完全性（integrity）を担保するために `go.sum` ファイルにモジュールの SHA-256 チェックサム値を格納しているが，最初にモジュールをフェッチする場合はチェックサム値が分からないため無条件に信頼せざるを得ない。

公開されているモジュールのチェックサム値がデータベース化されていれば未知のモジュールに対しても（チェックサム・データベースを参照することで）ある程度の信頼性を確保できるだろう。

チェックサム・データベースは環境変数 `GOSUMDB` で指定する。
`GOSUMDB` の既定値は [Go] 1.13 では以下のようになっている。

```text
$ go env | grep GOSUMDB
GOSUMDB="sum.golang.org"
```

チェックサム・データベースを無効にする場合は `GOSUMDB` に `off` をセットする。

```text
go env -w GOSUMDB=off
```

また特定のモジュールを検索対象から除外する場合は，ミラーリング除外のときと同じく `GOPRIVATE` が使える。

公式チェックサム・データベース [`sum.golang.org`] は [Trillian] による[透過ログ](https://research.swtch.com/tlog "research!rsc: Transparent Logs for Skeptical Clients")（追記型データベース）で改竄に強いという特徴がある。

まぁ，最初から悪意のあるモジュールは排除しようがないが malicious code 混入は検知しやすくなるかも知れない。

### 【2019-09-06 追記】 バージョンタグの管理に注意

上で述べたように標準の [`sum.golang.org`] は追記型のデータベースで，既存データの変更はできない仕様になっている。

たとえばパッケージ/モジュールのリポジトリにバージョンタグ（`v0.1.0` など）を付ける際に，うっかり間違ったコミットにタグを付けたとしても，一度 [`sum.golang.org`] に登録されると取り消すことができない。
手動でタグの位置を変えても，モジュール読み込み時に

```text
SECURITY ERROR
This download does NOT match the one reported by the checksum server.
The bits may have been replaced on the origin server, or an attacker may
have intercepted the download attempt.

For more information, see 'go help module-auth'.
```

などと表示されビルドに失敗する。

少なくとも公開されているパッケージ/モジュールでは迂闊にバージョンタグを付け換えないよう，管理は慎重に行う必要がある。

## ミラーリング・サービスのプライバシー・ポリシー

今回，正式稼働したミラーリング・サービスのプライバシー・ポリシーは以下のページある。

- [Privacy: Go modules services](https://proxy.golang.org/privacy)

以前はいきなり Google のページに飛ばされて「なんだかなぁ」という感じだったが，多少マシになったようである（笑）

これによると公式サービスでは以下の情報を収集しているようだ。

{{< fig-quote type="markdown" title="Privacy: Go modules services" link="https://proxy.golang.org/privacy" lang="en" >}}
- Request timestamp
- Client IP address
- Full request URL, including:
    - service domain, e.g. proxy.golang.org
    - URI path being requested 
- Response latency
- Response bytes sent
- Response code sent
- The response returned by the go command when it runs in our systems
- Whether the request hit our frontend cache
- Whether the request hit a cache elsewhere in the system (but not the frontend)
- Name of the Google machine that processed this request, e.g. machine101
{{< /fig-quote >}}

取得した情報については

{{< fig-quote type="markdown" title="Privacy: Go modules services" link="https://proxy.golang.org/privacy" lang="en" >}}
{{< quote >}}We do not store logged personally identifiable information such as IP addresses for more than 30 days. We also do not correlate or combine information from our request logs with any personal information that you have provided Google for other services.{{< /quote >}}
{{< /fig-quote >}}

ということらしいが，最後の

{{< fig-quote type="markdown" title="Privacy: Go modules services" link="https://proxy.golang.org/privacy" lang="en" >}}
{{< quote >}}We intend to aggregate and anonymize usage metrics to measure popularity for Go modules and share this popularity data with the Go community.{{< /quote >}}
{{< /fig-quote >}}

てのがねぇ。
[プライバシーに敵対的な企業]({{< ref "/remark/2018/04/handling-privacy.md" >}} "誰がプライバシーを支配するのか")のサービスだと思うとあまり利用したくない気分なのだが（もちろん偏見），はてさて[^lang1]。

[^lang1]: [Go 言語]にしろ（最近ブームが再燃している） Dart 言語にしろ，言語系のプロダクトやサービスは Google から切り離してほしいよなぁ。 Alphabet の傘下から外れろとまでは言わないから。

## ブックマーク

- [Go 1.13 と 1.14 （Go 2 へ向けて）]({{< ref "/release/2019/06/next-steps-toward-go-2.md" >}})
- [Go 言語の環境変数管理]({{< relref "./go-env.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[Go]: https://golang.org/ "The Go Programming Language"
[Trillian]: https://github.com/google/trillian "google/trillian: A transparent, highly scalable and cryptographically verifiable data store."
[Certificate Transparency]: https://www.certificate-transparency.org/

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
