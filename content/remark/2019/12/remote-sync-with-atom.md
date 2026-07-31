+++
title = "ATOM エディタでリモートのディレクトリ・ファイルを同期する"
date =  "2019-12-31T19:13:34+09:00"
description = "remote-sync が設定や操作がシンプルなので，これを使うことにした。 "
image = "/images/attention/kitten.jpg"
tags = [ "atom", "openssh" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

まぁ，独りで作業するなら [git] で管理するほうが安全なのだが，リモートのサーバとちょっとしたファイルのやり取りをするために [FileZilla] を起動して接続するのは大袈裟でかったるいなぁ思い始め「そういや [ATOM] のパッケージでリモートのファイルを直接編集できるパッケージがなかったっけ？」と探してみた。

[remote-ftp](https://atom.io/packages/remote-ftp) を推す記事が多かったので最初はこれを導入してみたが，何故か sftp で接続できんのよ。
しょうがないので他のパッケージを漁ってみたのだが [remote-sync] が設定や操作がシンプルなので，これを使うことにした。
厳密にはリモートのファイルを直接編集するわけではないのだが，まぁいいや。

まずはローカル側とリモート側で同期するディレクトリを作成する。

ローカル側で作成したディレクトリをプロジェクトのルートとして [ATOM] を起動し，コマンドパレットから “Remote Sync: Configure” を選択するか Tree View のコンテキストメニューから  “Remote Sync” → “Configure” と選択する。

{{< fig-img src="./remote-sync-context-menu.png" link="./remote-sync-context-menu.png" width="680" >}}

すると以下の設定画面が表示されるので必要な情報を設定して `[Save]` する。

{{< fig-img src="./remote-sync-configure.png" link="./remote-sync-configure.png" width="530" >}}

これで準備完了。
“uploadOnSave” にチェックが入っていればセーブするたびにリモートのファイルを更新してくれる。
その他の機能についてはコマンドパレットまたはコンテキストメニューから起動できる[^diff3]。

[^diff3]: Diff ツールについては初期状態では何も指定されていないので [remote-sync] の設定で任意のツールを指定する。個人的には KDiff3 がオススメ。 [Ubuntu] なら apt コマンドで導入できる。

ちなみに設定内容はロカール側のプロジェクト・ディレクトリ直下の `.remote-sync.json` に格納されている。
中身はこんな感じ。

```json
{
  "logger": {
    "title": "Remote Sync"
  },
  "uploadOnSave": true,
  "useAtomicWrites": true,
  "deleteLocal": true,
  "hostname": "example.com",
  "port": "2222",
  "target": "/home/username/docs/sync",
  "ignore": [
    ".remote-sync.json",
    ".git/**"
  ],
  "username": "username",
  "watch": [],
  "useAgent": true,
  "transport": "scp"
}
```

このファイルさえあれば任意のディレクトリを同期ディレクトリにできるわけだ。

とりあえず，こんなもんかな。

[ATOM]: https://atom.io/
[remote-sync]: https://atom.io/packages/remote-sync
[FileZilla]: https://filezilla-project.org/ "FileZilla - The free FTP solution"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenSSH]: http://www.openssh.com/ "OpenSSH"
[git]: https://git-scm.com/
[Git]: https://git-scm.com/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
