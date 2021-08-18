+++
title = "Go と VS Code"
date =  "2021-02-27T18:47:11+09:00"
description = "ぶっちゃけ ATOM の go-plus よりも出来がいいので，ゆるゆると移行していきますよっと。"
image = "/images/attention/go-logo_blue.png"
tags = ["vscode", "editor", "golang"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

どっちのセクションで書こう悩んだが，所詮ツールの話だし，こっち側で。

1. [パソコンに Visual Studio Code を導入する（再チャレンジ）]({{< ref "/remark/2021/02/installing-vscode-again.md" >}})
2. [Go と VS Code]({{< ref "/remark/2021/02/golang-with-vscode.md" >}}) ← イマココ
3. [Markdown と VS Code]({{< ref "/remark/2021/02/markdown-with-vscode.md" >}})
4. [Java と VS Code]({{< ref "/remark/2021/08/java-with-vscode.md" >}})

それでは[前回]の続き。
[VS Code] に [Go 言語][Go]開発支援環境を入れる。
といっても一択だけどね。

- [Go - Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=golang.go)

インストールはこちら。

```text
$ code --install-extension golang.go
```

一応 Google が公式に参加している拡張機能らしい。

- [The VS Code Go extension joins the Go project - The Go Blog](https://blog.golang.org/vscode-go)
- [Visual Studio Code Go extension joins the Go project](https://code.visualstudio.com/blogs/2020/06/09/go-extension)

しかも最近，こちらも事実上のオフィシャルである [gopls] が，この拡張機能の既定の [language server](https://langserver.org/ "Langserver.org") となったようだ。

- [Gopls on by default in the VS Code Go extension - The Go Blog](https://blog.golang.org/gopls-vscode-go)

個人的にひとつだけ気に入らないところがあって，既定の設定のままではプロジェクト・フォルダを開いたときにトップに `go.mod` ファイルがないと上手く動かないみたいなんだよね（[モジュール対応モード](https://zenn.dev/spiegel/articles/20210223-go-module-aware-mode "Go のモジュール管理【バージョン 1.16 改訂版】")の場合）。

最初は拡張機能側の問題かと思っていたが，どうも [gopls] の制限らしい。

- [tools/workspace.md at master · golang/tools · GitHub](https://github.com/golang/tools/blob/master/gopls/doc/workspace.md)

一応の回避策はあって `settings.json` に以下の[オプションを設定](https://github.com/golang/tools/blob/master/gopls/doc/settings.md#experimentalworkspacemodule-bool)することでトップ以外の複数の `go.mod` ファイルに対応できるようだ。

```json
{
    "gopls": {
        "build.experimentalWorkspaceModule": true
    }
}
```

ただし `go.mod` ファイルで定義したモジュール名は開いたプロジェクト・フォルダ内で一意でなければならない。
名前が被ってるとめっさ怒られる。

まぁ，理屈は分からんでもないが，私の場合，ブログなどのドキュメント環境でサンプルコードを書き散らしていて，それらのモジュール名が大抵 "`sample`” だったりするので被りまくりなんだよなぁ sigh...

ちなみに，この [`experimentalWorkspaceModule` オプション](https://github.com/golang/tools/blob/master/gopls/doc/settings.md#experimentalworkspacemodule-bool)は一時的な措置なので将来的にはなくなるようだ。
それまでには何とかするということだろう。

それ以外の部分ではぶっちゃけ [ATOM] の [go-plus] よりも出来がいいので，ゆるゆると移行していきますよっと。

現時点での `settings.json` の内容はこんな感じ。

```json
{
"go.buildOnSave": "off",
	"go.autocompleteUnimportedPackages": true,
	"go.formatTool": "goimports",
	"go.lintTool": "golangci-lint",
	"go.gotoSymbol.includeGoroot": true,
	"go.testOnSave": true,
	"go.coverageOptions": "showUncoveredCodeOnly",
	"go.coverOnSave": true,
	"go.coverOnSingleTestFile": true,
	"go.coverageDecorator": {
		"type": "gutter"
	},
	"gopls": {
		"build.experimentalWorkspaceModule": true
	}
}
```

## 【2021-03-02 追記】

書き忘れていたが，コマンドパレットから

```text
>Go: Install/Update Tools
```

と入力すると関連ツールのインストールおよびアップデートができる。

{{< fig-img src="./update-tools.png" title="Go: Install/Update Tools" link="./update-tools.png" width="612" >}}

これを実行すると

```text
Tools environment: GOPATH=/home/username/go
Installing 9 tools at /home/username/go/bin in module mode.
  gopkgs
  go-outline
  gotests
  gomodifytags
  impl
  goplay
  dlv
  golangci-lint
  gopls

Installing github.com/uudashr/gopkgs/v2/cmd/gopkgs (/home/username/go/bin/gopkgs) SUCCEEDED
Installing github.com/ramya-rao-a/go-outline (/home/username/go/bin/go-outline) SUCCEEDED
Installing github.com/cweill/gotests/... (/home/username/go/bin/gotests) SUCCEEDED
Installing github.com/fatih/gomodifytags (/home/username/go/bin/gomodifytags) SUCCEEDED
Installing github.com/josharian/impl (/home/username/go/bin/impl) SUCCEEDED
Installing github.com/haya14busa/goplay/cmd/goplay (/home/username/go/bin/goplay) SUCCEEDED
Installing github.com/go-delve/delve/cmd/dlv (/home/username/go/bin/dlv) SUCCEEDED
Installing github.com/golangci/golangci-lint/cmd/golangci-lint (/home/username/go/bin/golangci-lint) SUCCEEDED
Installing golang.org/x/tools/gopls (/home/username/go/bin/gopls) SUCCEEDED

All tools successfully installed. You are ready to Go :).
```

てな感じで `go install` してくれるようだ。
自動でアップデートしてくれるといいのだが...

[前回]: {{< relref "./installing-vscode-again.md" >}} "パソコンに Visual Studio Code を導入する（再チャレンジ）"
[VS Code]: https://code.visualstudio.com/ "Visual Studio Code - Code Editing. Redefined"
[ATOM]: https://atom.io/
[Go]: https://golang.org/ "The Go Programming Language"
[go-plus]: https://atom.io/packages/go-plus
[gopls]: https://pkg.go.dev/golang.org/x/tools/gopls "gopls · pkg.go.dev"

## 参考図書

{{% review-paapi "B08CZ2C3NZ" %}} <!-- Software Design (2020年8月号) -->
{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
