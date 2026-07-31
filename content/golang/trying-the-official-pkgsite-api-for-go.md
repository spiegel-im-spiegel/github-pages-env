+++
title = "Go 公式のモジュール照会 API を試す"
date =  "2026-05-22T12:37:48+09:00"
description = "GET のみの RESTful API として提供されている。"
isCJKLanguage = true
image = "/images/attention/go-logo_blue.png"
tags = [ "engineering", "golang", "api", "package", "module" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

公開された [Go] パッケージやモジュールのメタデータを照会する API が公開されたそうだ。

- [API Documentation - Go Packages](https://pkg.go.dev/v1beta/api)
- [api package - golang.org/x/pkgsite/internal/api - Go Packages](https://pkg.go.dev/golang.org/x/pkgsite/internal/api)
- [Introducing the pkg.go.dev API - The Go Programming Language](https://go.dev/blog/pkgsite-api)

{{< fig-quote type="markdown" title="Introducing the pkg.go.dev API" link="https://go.dev/blog/pkgsite-api" lang="en" >}}
Built for stability and efficient caching, the API uses a stateless, GET-only architecture. Primary endpoints are currently hosted under the /v1beta path. Following a period of community feedback and confirmed stability, we intend to transition toward a formal v1 release.
{{< /fig-quote >}}

というわけで現時点では `v1beta` として提供されている。
エンドポイントは以下の通り：

{{< fig-quote class="nobox" type="markdown" title="Introducing the pkg.go.dev API" link="https://go.dev/blog/pkgsite-api" lang="en" >}}
| Endpoint | Description |
|---|---|
| `/v1beta/package/{path}` | Information about the package at `{path}`. |
| `/v1beta/module/{path}` | Information about the module at `{path}`. |
| `/v1beta/versions/{path}` | Versions of the module at `{path}`. |
| `/v1beta/packages/{path}` | Information about packages of the module at `{path}`. |
| `/v1beta/search?q={query}` | Search results for a given query. |
| `/v1beta/symbols/{path}` | List of symbols declared by the package at `{path}`. |
| `/v1beta/imported-by/{path}` | Paths of packages importing the package at `{path}`. |
| `/v1beta/vulns/{path}` | Vulnerabilities of the module or package at `{path}`. |
{{< /fig-quote >}}

たとえば標準の [`errors`] パッケージの情報を取得するには，以下の URL を指定すればいい。

```text
$ curl -s https://pkg.go.dev/v1beta/package/errors | jq .
{
  "modulePath": "std",
  "version": "v1.26.3",
  "isLatest": true,
  "isStandardLibrary": true,
  "goos": "all",
  "goarch": "all",
  "path": "errors",
  "name": "errors",
  "synopsis": "Package errors implements functions to manipulate errors.",
  "isRedistributable": true
}
```

もちろん標準以外の公開パッケージも取得できる。

```text
$ curl -s https://pkg.go.dev/v1beta/package/github.com/goark/errs | jq .
{
  "modulePath": "github.com/goark/errs",
  "version": "v1.3.4",
  "isLatest": true,
  "isStandardLibrary": false,
  "goos": "all",
  "goarch": "all",
  "path": "github.com/goark/errs",
  "name": "errs",
  "synopsis": "Package errs implements functions to manipulate error instances.",
  "isRedistributable": true
}
```

モジュールの情報も取れる。

```text
$ curl -s https://pkg.go.dev/v1beta/module/github.com/goark/errs | jq .
{
  "path": "github.com/goark/errs",
  "version": "v1.3.4",
  "commitTime": "2026-05-09T04:07:47Z",
  "isLatest": true,
  "isRedistributable": true,
  "isStandardLibrary": false,
  "hasGoMod": true,
  "repoUrl": "https://github.com/goark/errs"
}
```

面白いところでは脆弱性情報も取得できるようだ。
たとえば昨年の [`golang.org/x/crypto` v0.45.0 における脆弱性](https://groups.google.com/g/golang-announce/c/w-oX3UxNcZA "Vulnerabilities in golang.org/x/crypto")情報であれば

```text
$ curl -s https://pkg.go.dev/v1beta/vulns/golang.org/x/crypto?version=v0.45.0 | jq .
{
  "items": [
    {
      "id": "GO-2026-5005",
      "summary": "",
      "details": "Invoking key constraints not enforced in golang.org/x/crypto/ssh/agent",
      "fixedVersion": ""
    },
    {
      "id": "GO-2026-5006",
      "summary": "",
      "details": "Invoking agent constraints dropped when forwarding keys in golang.org/x/crypto/ssh/agent",
      "fixedVersion": ""
    },

    ...

    {
      "id": "GO-2026-5033",
      "summary": "",
      "details": "Invoking pathological inputs can lead to client panic in golang.org/x/crypto/ssh/agent",
      "fixedVersion": ""
    }
  ],
  "total": 13
}
```

という感じに取得できる（数が多いので途中を端折ってる。あしからず）。
CVE ID ベースじゃないんだな。

CLI ツールも提供されているようだ。
[Go] のビルド環境があれば以下のコマンドでインストールできる。

```text
$ go install golang.org/x/pkgsite/cmd/internal/pkgsite-cli@latest
```

ただ，パッケージのパスを見れば分かるように，あくまで API 利用のリファレンス実装という位置づけっぽい。

自作ツールで依存関係を可視化する [goark/depm](https://github.com/goark/depm "goark/depm: Visualize dependency packages and modules") があるのだが，こいつは内部で `go` コマンドを呼んでいる。
これを API ベースに書き換えたら面白いかな。
いや，ますます遅くなりそうだな。
ちょっと考えるところだ。

[Go]: https://go.dev/
[`errors`]: https://pkg.go.dev/errors "errors package - errors - Go Packages"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400535" %}} <!-- 効率的なGo : Efficient Go -->
{{% review-paapi "B0DNYMMBBQ" %}} <!-- Go言語で学ぶ並行プログラミング -->
