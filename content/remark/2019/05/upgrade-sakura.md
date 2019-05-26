+++
title = "さくらのレンタルサーバの Git が「使える！」ようになっていた"
date =  "2019-05-26T22:21:14+09:00"
description = "Hugo + Git の組み合わせで git-push から自動デプロイができるところまでは確認済みなので，この構成でのんびりリニューアル作業を行うとしよう。"
image = "/images/attention/kitten.jpg"
tags = [ "site", "web", "sakura", "git", "golang" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

自宅マシンの OS 換装も一段落ついたので，そろそろ[本家サイト]のリニューアルの準備を進めようと思うのだが，その前にホストである「[さくらのレンタルサーバ]」がどうなったか確認しておかないと。

「[さくらのレンタルサーバ]」は今年（2019年）に入ってから大規模な OS アップグレードを行っている。

- [FreeBSDのアップデートに伴う変更点 – さくらのサポート情報](https://help.sakura.ad.jp/hc/ja/articles/360000190121)

私が利用しているサーバは4月にアップグレードが完了していて，ログインすると以下のバージョン表記になった。

```text
FreeBSD 11.2-RELEASE-p9 (GENERIC) #0: Tue Feb  5 15:30:36 UTC 2019
```

嬉しいのは [git] で

```text
$ git version
git version 2.19.2
```

と，かなりまともになった。
アップグレード前は 2.7 とか巫山戯たバージョンで git-submodule も使えない状態だったのでかなりの進歩と言えよう。

これでようやく「[さくらのレンタルサーバ]」に（自前でビルドせずとも）ベア・リポジトリを置いてコンテンツのバージョン管理や自動デプロイができるようになったよ。
スタンダード・プランですらストレージが100GBもあって持て余してたんだよねぇ。

そして [Go] コンパイラが導入された。

```text
$ go version
go version go1.11.1 freebsd/amd64
```

かなり古くてセキュリティ脆弱性とかちゃんと始末してるのか不安なところもあるが，このバージョンならモジュール・モードも（一応）使えるし，まぁいいか。
バックエンド側の処理を [Go] で書くこともできるだろう。

とりあえず [photo.Baldanders.info](https://photo.baldanders.info/) を使って [Hugo] + [Git] の組み合わせで git-push から自動デプロイができるところまでは[確認済み]({{< ref "/remark/2019/01/sakura-and-hugo.md" >}} "さくらのレンタルサーバ上で Hugo によるサイト管理を行う")なので，この構成でのんびりリニューアル作業を行うとしよう。

[さくらのレンタルサーバ]: https://www.sakura.ne.jp/
[本家サイト]: https://baldanders.info/ "Baldanders.info"
[git]: https://git-scm.com/
[Git]: https://git-scm.com/
[Go]: https://golang.org/ "The Go Programming Language"
[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
