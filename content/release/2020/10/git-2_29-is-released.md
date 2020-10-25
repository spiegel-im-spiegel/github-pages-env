+++
title = "Git v2.29 がリリースされた"
date =  "2020-10-25T18:09:33+09:00"
description = "SHA-2 コミット・ハッシュの実験的サポート / Windows 版 Git Credential Manager のアップグレード"
image = "/images/attention/tools.png"
tags  = [ "git", "tools", "hash", "sha-1", "ubuntu", "windows" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先日 [Git][git] v.2.29 がリリースされた。

- [[ANNOUNCE] Git v2.29.0](https://lore.kernel.org/git/xmqqy2k2t77l.fsf@gitster.c.googlers.com/)

## SHA-2 コミット・ハッシュの実験的サポート

v2.29 ではコミット・ハッシュに関する重大な仕様変更がある。

{{< fig-quote type="markdown" title="Highlights from Git 2.29 - The GitHub Blog" link="https://github.blog/2020-10-19-git-2-29-released/" lang="en" >}}
{{% quote %}}Git 2.29 includes experimental support for writing your repository’s objects using a SHA-256 hash of their contents, instead of using SHA-1{{% /quote %}}.
{{< /fig-quote >}}

といっても今回は実験的なサポート（experimental support）で，試すのであれば以下のような感じでできるらしい。

```text
$ git init --object-format=sha256 sample-repo
Initialized empty Git repository in /home/username/sample-repo/.git/

$ cd sample-repo
$ echo 'Hello, SHA-256!' >README.md
$ git add README.md
$ git commit -m "README.md: initial commit"
[main (root-commit) 6d45449] README.md: initial commit
 1 file changed, 1 insertion(+)
 create mode 100644 README.md

 $ git rev-parse HEAD
 6d45449028a8e76500adbfe7330e779d5dc4a3a14fca58ff08ec354c58727b2c
```

当然ながら SHA-1 ベースのコミット・ハッシュと SHA-2 (SHA256) ベースのコミット・ハッシュとの間には互換性がない。
当面は2系統のリポジトリを使い分けて運用することになりそうだ。
ただし，将来的には両者を相互運用できるようにするらしい。

{{< fig-quote type="markdown" title="Highlights from Git 2.29 - The GitHub Blog" link="https://github.blog/2020-10-19-git-2-29-released/" lang="en" >}}
{{% quote %}}In future releases, Git will support interoperating between repositories with different object formats by computing both a SHA-1 and SHA-256 hash of each object it writes, and storing a translation table between them. This will eventually allow repositories that store their objects using SHA-256 to interact with (sufficiently up-to-date) SHA-1 clients, and vice-versa. It will also allow converted SHA-256 repositories to have their references to older SHA-1 commits still function as normal (e.g., if I write a commit whose message references an earlier commit by its SHA-1 name, then Git will still be able to follow that reference even after the repository is converted to use SHA-256 by consulting the translation table){{% /quote %}}.
{{< /fig-quote >}}

ここまで到達すれば SHA-2 ベースへ本格的に切り替えていってもいいかも知れない。

## Windows 版 Git Credential Manager のアップグレード

[Git for Windows] では，もうひとつ重大な変更がある。

{{< fig-quote type="markdown" title="Release Git for Windows 2.29.0 · git-for-windows/git" link="https://github.com/git-for-windows/git/releases/tag/v2.29.0.windows.1" lang="en" >}}
{{% quote %}}This version upgrades existing users of [Git Credential Manager for Windows](https://github.com/microsoft/Git-Credential-Manager-for-Windows/) (which was just deprecated) to [Git Credential Manager Core](https://github.com/microsoft/Git-Credential-Manager-Core) ("GCM Core", which is the designated successor of the former). This is necessary because [GitHub deprecated password-based authentication](https://github.blog/changelog/2019-08-08-password-based-http-basic-authentication-deprecation-and-removal/) and intends to remove support for it soon, and GCM Core is prepared for this change.{{% /quote %}}.
{{< /fig-quote >}}

GitHub リポジトリに HTTPS でアクセスしている場合は注意が必要かも知れない（[Git for Windows] を使ってないので，どの程度影響するか分かってない。ゴメンペコン {{% emoji "ゴメン" %}}）。

## ブックマーク

- [[ANNOUNCE] Git v2.29.1](https://lore.kernel.org/git/xmqq4kmlj9q9.fsf@gitster.c.googlers.com/)
- [Highlights from Git 2.29 - The GitHub Blog](https://github.blog/2020-10-19-git-2-29-released/)
- [Release Git for Windows 2.29.0 · git-for-windows/git · GitHub](https://github.com/git-for-windows/git/releases/tag/v2.29.0.windows.1)
- [Release Git for Windows 2.29.1 · git-for-windows/git · GitHub](https://github.com/git-for-windows/git/releases/tag/v2.29.1.windows.1)
- [「Git for Windows 2.29.0」が公開 ～セットアップ時にデフォルトブランチ名を設定可能 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1284871.html)

- [PPA から Git をインストールする]({{< ref "/remark/2019/04/install-git-from-ppa.md" >}})

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[git]: https://git-scm.com/
[Git for Windows]: https://gitforwindows.org/
[PPA]: https://launchpad.net/ubuntu/+ppas "Personal Package Archives : Ubuntu"
