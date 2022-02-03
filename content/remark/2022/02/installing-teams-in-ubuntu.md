+++
title = "Ubuntu に Microsoft Teams をインストールする"
date =  "2022-02-03T18:48:48+09:00"
description = "公式のインストール手順の説明が簡潔すぎて分かりにくい（笑）"
image = "/images/attention/kitten.jpg"
tags = [ "tools", "install", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

ちょいと訳あって自宅の Ubuntu 機に急遽 [Microsoft Teams][Teams] をインストールすることになったので，忘れる前に覚え書き。

Linux 用の [Teams] クライアントは DEB または RPM 形式で[インストールできる](https://docs.microsoft.com/ja-jp/microsoftteams/get-clients#linux "Microsoft Teams のクライアントを取得する - Microsoft Teams | Microsoft Docs")とあるのだが，説明が簡潔すぎて分かりにくいので以下の記事を参考にした。

- [ubuntu 20.04にてteamsをインストールする - Qiita](https://qiita.com/berukokoko/items/d5a4f3a1a7930aaf9bd2)

まずは APT に電子署名検証用の公開鍵を追加する。

```text
$ curl https://packages.microsoft.com/keys/microsoft.asc | sudo apt-key add -
```

続いて [Teams] 用のリポジトリを登録。

```text
$ sudo sh -c 'echo "deb [arch=amd64] https://packages.microsoft.com/repos/ms-teams stable main" > /etc/apt/sources.list.d/teams.list'
```

ここまでできれば

```text
$ sudo apt update
$ sudo apt install teams
```

でインストール完了である。
ただしアイコンが

{{< fig-img src="./teams-icon.png" link="./teams-icon.png" >}}

てな感じになってるのでプレビュー版ってことかな？ Windows 版に比べると機能がしょぼい気がする。
まぁ，今回は最低限リモート会議に参加できれば OK ちうことで。

## ブックマーク

- [非推奨となったapt-keyの代わりにsigned-byとgnupgを使う方法 - 2021-05-05 - ククログ](https://www.clear-code.com/blog/2021/5/5.html)
- [第675回　apt-keyはなぜ廃止予定となったのか：Ubuntu Weekly Recipe｜gihyo.jp … 技術評論社](https://gihyo.jp/admin/serial/01/ubuntu-recipe/0675)

[Teams]: https://www.microsoft.com/ja-jp/microsoft-teams/group-chat-software "リモート ワーク - コラボレーション ツール | Microsoft Teams"
<!-- eof -->
