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

[前回]: {{< relref "./installing-vscode-again.md" >}} "パソコンに Visual Studio Code を導入する（再チャレンジ）"
[VS Code]: https://code.visualstudio.com/ "Visual Studio Code - Code Editing. Redefined"
[ATOM]: https://atom.io/
[Go]: https://golang.org/ "The Go Programming Language"
[go-plus]: https://atom.io/packages/go-plus
[gopls]: https://pkg.go.dev/golang.org/x/tools/gopls "gopls · pkg.go.dev"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
