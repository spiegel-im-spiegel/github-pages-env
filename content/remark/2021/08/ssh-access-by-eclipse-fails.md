+++
title = "Eclipse の SSH アクセスで失敗した話"
date =  "2021-08-30T20:55:10+09:00"
description = "これのせいで半日近く作業が止まってしまったよ。"
image = "/images/attention/kitten.jpg"
tags = [ "engineering", "java", "git", "openssh", "editor" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

トラブル対応で自前の環境に [Eclipse] を入れて [GitHub] リポジトリにアクセスする羽目になった。
[GitHub] リポジトリへのアクセスには HTTPS と SSH の2通りある。
このうち HTTPS のほうはアカウントのパスワードが使えなくなり，代替手段として personal access token を取得する必要がある[^pat1]。

[^pat1]: [GitHub Docs の内容](https://docs.github.com/ja/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token "個人アクセストークンを使用する - GitHub Docs")は少し古いみたいで，Settings の構成がちょっと違う。 `Settings > Developer settings > Personal access tokens` の順に辿っていけばよい。ちなみに personal access token の Expiration を無期限に設定すると [GitHub] が物凄い勢いで警告を出してくる（笑）

- [個人アクセストークンを使用する - GitHub Docs](https://docs.github.com/ja/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token)

まぁ personal access token の話はこれくらいにして，本題は SSH 接続のほう。
ちなみに今回の構成は以下の通り。

- OS : Windows 10
- Eclipse : [Pleiades All in One](https://mergedoc.osdn.jp/ "Java 開発環境 - Eclipse 日本語化 Pleiades プラグイン") 2021-06 リリース版

[Eclipse] で Git および SSH をドライブしてるのは [EGit](https://www.eclipse.org/egit/ "EGit | The Eclipse Foundation") と [JSch](http://www.jcraft.com/jsch/ "JSch - Java Secure Channel") だそうだが，ユーザからこれが見えることはない。
「自前で構築している環境（[Git for Windows] と [PuTTY]）は使えないんだろうなぁ。鍵を作るところから始めるか」とため息をつきつつ，以下のページを参考に（感謝！）作業を始めたのだが

- [Eclipseでgit操作(GitHubからcloneしてpushまで/ssh接続) | ITSakura](https://itsakura.com/eclipse-github-clone-push)

最初の1フィートで驚愕してしまった。

{{< fig-img src="eclipse-ssh2.png" link="eclipse-ssh2.png" width="762" >}}

いやいやいや。
DSA か RSA の1024ビット鍵しか作れんのかい！ まさか世の Eclipse ユーザはみんな1024ビット鍵で SSH にアクセスしてるの？ アカンやろ。アカンよね！？

首を捻りながらも設定を済ませ，いざ clone しようとしたら今度は {{% quote lang="en" %}}No supported authentication methods available{{% /quote %}} とか言ってくさる。

{{< div-gen type="markdown" class="center" >}}
[ちょっと何言ってるか分からない](https://dic.nicovideo.jp/a/%E3%81%A1%E3%82%87%E3%81%A3%E3%81%A8%E4%BD%95%E8%A8%80%E3%81%A3%E3%81%A6%E3%82%8B%E3%81%8B%E5%88%86%E3%81%8B%E3%82%89%E3%81%AA%E3%81%84)...
{{< /div-gen >}}

周囲の人たちは問題なくアクセスできてるみたいなので私の環境の問題だと思うが，マジで分からない。
これのせいで半日近く作業が止まってしまったよ。
で，さんざん悩んだ挙句に気が付いた。
これってひょっとして plink で SSH アクセスしようとしている？

念のために説明すると plink は Windows 用の SSH クライアント兼ターミナル・エミュレータである [PuTTY] が他ツールとの連携用に提供している SSH 接続ツールだ。
Git がインストールされている環境で

```
set GIT_SSH=C:\Program Files\PuTTY\plink.exe
```

とか環境設定しておけば OpenSSH ではなく[PuTTY]/plink で接続しようとする。
どうやら [Eclipse] はこれを認識しているのかな？

となると話は早い。
ぶっちゃけ [Git for Windows] & [PuTTY] で既に環境が出来ていれば [Eclipse] 側の SSH 設定は不要ということだ。
というわけでめでたく GitHub リポジトリに SSH でアクセスできた。
ようやくトラブル対応できる。
RSA-1024 とか玩具の南京錠のごとき鍵を作る必要もなく，既に運用している EdDSA 鍵で問題なし。
上のスナップショットで作った鍵はとっとと捨てたよ（笑）

[Eclipse] の git リポジトリ操作も見せてもらったが面倒臭すぎてまともに運用できる気がしない。
いや，これ，マジで VS Code を啓蒙しないとダメかしらん。

## ブックマーク

- [GnuPG for Windows : gpg-agent について]({{< ref "/openpgp/using-gnupg-for-windows-2.md" >}})
- [OpenSSH の認証鍵を GunPG で作成・管理する]({{< ref "/openpgp/ssh-key-management-with-gnupg.md" >}})

[Eclipse]: https://www.eclipse.org/ "Enabling Open Innovation & Collaboration | The Eclipse Foundation"
[GitHub]: https://github.com/
[Git for Windows]: https://gitforwindows.org/
[PuTTY]: https://www.chiark.greenend.org.uk/~sgtatham/putty/ "PuTTY: a free SSH and Telnet client"

## 参考図書

{{% review-paapi "B0893LQ5KY" %}} <!-- Spring Boot 2 入門 -->
{{% review-paapi "4621303252" %}} <!-- Effective Java 第3版 -->
