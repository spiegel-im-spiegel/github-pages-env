+++
title = "Azure App Service に VS Code でデプロイできなかった話"
date =  "2022-04-06T20:37:03+09:00"
description = "当面は VS Code 1.65.x にダウングレードするしかないらしい。"
image = "/images/attention/kitten.jpg"
tags = [ "azure", "vscode", "engineering", "web"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

私は [Azure App Service][App Service] に手動でデプロイする際は [VS Code] の [Azure App Service 拡張機能][Azure App Service]を使ってるんだけど，ここのところ頻繁にデプロイに失敗してたのよ。
最初は [App Service] のほうがやらかしてるんだと思ってたんだけど（だって上手く行くときもあるし），ググったら違ってたようだ。

{{< fig-quote type="" title="ECONNRESET trying azure deploy webapp VS Code · Issue #2194 · microsoft/vscode-azureappservice" link="https://github.com/microsoft/vscode-azureappservice/issues/2194#issuecomment-1087857459" lang="en" >}}
VS Code 1.66 upgraded to Node 16 which has been causing deployment issues for both Functions and App Service. Unfortunately, the only current known workaround is to downgrade VS Code to 1.65.x. We're currently investigating a proper fix.
{{< /fig-quote >}}

まじすか `orz`

ダウングレードってどうやんだ？ と思ったが単純に 1.65.x のインストール・パッケージを拾ってきて上書きインストールすればいいようだ。

- [Visual Studio Code February 2022](https://code.visualstudio.com/updates/v1_65)

Windows 環境では自動更新を無効にするのをお忘れなく。
やれやれ

{{< div-box type="markdown" >}}
**【2022-04-14 追記】**
v1.66.2 で直った。
よかったよかった。

- [Release March 2022 Recovery 2 · microsoft/vscode · GitHub](https://github.com/microsoft/vscode/releases/tag/1.66.2)
{{< /div-box >}}

## ブックマーク

- [Investigate ECONNRESET issue during "zipdeploy" call · Issue #2844 · microsoft/vscode-azurefunctions · GitHub](https://github.com/microsoft/vscode-azurefunctions/issues/2844)

[App Service]: https://azure.microsoft.com/en-us/services/app-service/ "App Service — Build & Host Web Apps | Microsoft Azure"
[VS Code]: https://code.visualstudio.com/ "Visual Studio Code - Code Editing. Redefined"
[Azure App Service]: https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-azureappservice "Azure App Service - Visual Studio Marketplace"
