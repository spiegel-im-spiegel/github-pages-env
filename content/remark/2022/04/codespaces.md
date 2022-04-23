+++
title = "やっと Codespaces が使える"
date =  "2022-04-23T22:45:12+09:00"
description = "Go が動くぞ！"
image = "/images/attention/kitten.jpg"
tags = [ "github", "codespaces", "vscode" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

1. [パソコンに Visual Studio Code を導入する（再チャレンジ）]({{< ref "/remark/2021/02/installing-vscode-again.md" >}})
2. [Go と VS Code]({{< ref "/remark/2021/02/golang-with-vscode.md" >}})
3. [Markdown と VS Code]({{< ref "/remark/2021/02/markdown-with-vscode.md" >}})
4. [Java と VS Code]({{< ref "/remark/2021/08/java-with-vscode.md" >}})
5. [やっと Codespaces が使える]({{< ref "/remark/2022/04/codespaces.md" >}}) ← イマココ

さて，[独り GitHub Team]({{< ref "/remark/2022/03/github-team.md" >}}) も契約したし，よーやく [Codespaces] が使えるようになったよ。

今のところ [Codespaces] は GitHub Team または GitHub Enterprise Cloud のメンバで利用できる。
2022-04-23 時点の料金（米ドル建て）は以下の通り。

{{< fig-img-quote src="./priceing.png" title="About billing for Codespaces - GitHub Docs" link="https://docs.github.com/billing/managing-billing-for-github-codespaces/about-billing-for-codespaces" lang="en" width="626" >}}

ちなみに，使いすぎないよう上限を設定することも可能。

## [Codespaces] クラウドに立つ

[Codespaces] の起動にはまず，ブラウザでリポジトリ・ページを開き，`[Code]` ボタンを押下する。

{{< fig-img src="./code.png" title="Create codespace on master" link="./code.png" lang="en" >}}

`[Create codespace on master]` ボタン押下で [Codespaces] のインスタンス生成が始まる。
上手く起動すればこんな感じに表示される。

{{< fig-img src="./codespae-on-the-browser.png" title="Create codespace on master" link="./codespae-on-the-browser.png" lang="en" width="1922" >}}

拡張機能も普通に入れられる。

{{< fig-img src="./extensions.png" title="Extensions" link="./extensions.png" lang="en" >}}

設定はユーザごとに保持して同期させることもできるようだ。

## Go が動くぞ！

[Codespaces] 上のターミナルで試しに Go コンパイラを動かしてみたら普通に動いた。

{{< fig-img src="./go-on-the-terminal.png" title="こいつ・・・動くぞ！" link="./go-on-the-terminal.png" width="961" >}}

## [GitHub Codespaces] 拡張機能

[VS Code] の拡張機能に [GitHub Codespaces] というのがあって，これを使うとローカルの [VS Code] 上で [Codespaces] のリソースにアクセスできる。
コマンドパレットから `Codespaces: Connect to Codespace` を選択すると生成済みのインスタンスの一覧が表示されるので，選択して接続すると [Codespaces] に接続した [VS Code] が起動する。

{{< fig-img src="./codespace-on-the-vscode.png" title="Codespace on the VS Code" link="./codespace-on-the-vscode.png" lang="en" width="1595" >}}

パソコンで作業する場合はこっちのほうがいいかもねぇ。

## Andorid タブレットでも動いた

{{< fig-img src="./codespace-on-the-android-browser.jpg" title="Codespace on the Android Browser" link="./codespace-on-the-android-browser.jpg" lang="en" width="1920" >}}

おわ。
ターミナルの状態も共有できるのか。
凄いな。

ちなみに Android の Firefox では [Codespaces] を起動できなかった。
冷遇されてるなぁ（笑）

## 後始末

[Codespaces] を終了する場合はコマンドパレットから `Codespaces: Stop Current Codespace` を選択して明示的に停止させること。
まぁ，最悪でも30分非活性状態なら自動的に停止するらしいけど。

[Codespaces] のインスタンスはリポジトリごとに生成されるのでご注意を。
また，使わなくなったインスタンスはマメに削除することをお勧めする。
インスタンスの削除は “Your codespaces” でできる。

{{< fig-img src="./your-codespaces.png" title="Your codespaces" link="./your-codespaces.png" lang="en" >}}

## ブックマーク

- [GitHub Codespaces をつかって 3分で始めるサービス開発 | Wantedly, Inc.](https://www.wantedly.com/companies/wantedly/post_articles/355862)
- [GitHub Codespaces で開発する - Wantedly Engineering Handbook](https://docs.wantedly.dev/fields/dev-tools/codespaces)

[Codespaces]: https://docs.github.com/codespaces "GitHub Codespaces Documentation - GitHub Docs"
[VS Code]: https://code.visualstudio.com/ "Visual Studio Code - Code Editing. Redefined"
[GitHub Codespaces]: https://code.visualstudio.com/docs/remote/codespaces "Developing with GitHub Codespaces"
