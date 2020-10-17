+++
title = "GitHub リポジトリの既定ブランチ名が main になるらしい"
date =  "2020-08-27T17:53:37+09:00"
description = "サービス・プロバイダやオープンソースな人たちが挙ってやってるやつね。"
image = "/images/attention/kitten.jpg"
tags = [ "politics", "engineering", "git", "github" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Twitter の TL で見かけたんだけどさ。

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">デフォルトブランチを変更していない場合は 2020/10/1 から自動的に新規プロジェクトのデフォルトブランチが master から main に変更されるとのこと。 / “Set the default branch for newly-created repositories - GitHub Changelog” <a href="https://t.co/zkXPp43aRd">https://t.co/zkXPp43aRd</a></p>&mdash; mattn (@mattn_jp) <a href="https://twitter.com/mattn_jp/status/1298803432984207360?ref_src=twsrc%5Etfw">August 27, 2020</a></blockquote>
{{< /fig-gen >}}

不覚にも3秒位「？？？」ってなっちゃったYO。

あ゙ー，あれね。
サービス・プロバイダやオープンソースな人たちが挙って「差別用語を狩ろう」っちう “[Virtue Signaling]({{< ref "/remark/2020/07/virtue-signaling.md" >}} "Virtue Signaling という広告")” なやつの一環ね。

なんでも

{{< fig-quote type="markdown" title="Set the default branch for newly-created repositories - GitHub Changelog" link="https://github.blog/changelog/2020-08-26-set-the-default-branch-for-newly-created-repositories/" lang="en" >}}
{{% quote %}}On **October 1, 2020**, if you haven't changed the default branch for new repositories for your user, organization, or enterprise, it will automatically change from `master` to `main`{{% /quote %}}.
{{< /fig-quote >}}

だそうで，何もしなかったら 2020-10-01 以降の新規作成リポジトリの既定ブランチが `main` になるらしい。
既存のリポジトリには影響しないし[設定画面](https://github.com/settings/repositories "Repositories")で既定ブランチをあらかじめ指定しておくことで回避することも可能である（リポジトリ毎に既定ブランチを指定することもできるみたい）。


個人的にはこういう政治広告に巻き込まれたくないんだけど，名前を変えること自体はさしたる手間ではないし，吝かではないといったところか。
強いていうなら CI サービス周りの対応次第かな。

## Git リポジトリの既定ブランチ名を指定するには

Git の設定で既定ブランチ名を変更するには，以下のコマンドラインで無問題[^git228]。

[^git228]: 既定ブランチ名の指定は [git 2.28 から使える](https://github.blog/2020-07-27-highlights-from-git-2-28/ "Highlights from Git 2.28 - The GitHub Blog")ようだ。

```text
$ git config --global init.defaultBranch foo
```

これで `$HOME/.config/git/config` (または `$HOME/.gitconfig`) ファイルに

```toml
[init]
  defaultBranch = foo
```

が追加される。

また `git init` コマンドで

```text
$ git init --initial-branch=foo myrepo
```

とすれば `foo` ブランチで初期化してくれるので，そのまま既定ブランチとして使えばよい。

## リモートにある Git リポジトリの既定ブランチ名を取得するには

Git の既定ブランチというのは要するに `HEAD` が指すブランチなので，リモートにある git リポジトリの既定ブランチ名を（clone しないで）取得するには

```text
$ git ls-remote --symref https://github.com/git/git HEAD
ref: refs/heads/master	HEAD
a5fa49ff0a8f3252c6bff49f92b85e7683868f8a	HEAD
```

などとすればいいようだ。
これでリポジトリ `https://github.com/git/git` の既定ブランチが `master` であることが分かる。

リポジトリ名のみ取り出したければ，たとえば sed コマンドと組み合わせて

```text
$ echo `git ls-remote --symref https://github.com/git/git HEAD | sed -n 's/^ref: refs\/heads\/\(.*\)\s\+HEAD$/\1/p'`
master
```

とかできるらしい。

## ブックマーク

- [github/renaming: Guidance for changing the default branch name for GitHub repositories](https://github.com/github/renaming)
    - [The default branch for newly-created repositories is now main - GitHub Changelog](https://github.blog/changelog/2020-10-01-the-default-branch-for-newly-created-repositories-is-now-main/)
- [「ブラックハット」も差別連想？--用語変更についてセキュリティコミュニティで議論に - ZDNet Japan](https://japan.zdnet.com/article/35156317/)
- [Twitter、コードやドキュメント内の用語「Whitelist/Blacklist」「Master/Slave」「Dummy value」などを好ましい用語へ置き換え、具体例も発表 － Publickey](https://www.publickey1.jp/blog/20/twitterwhitelistblacklistmasterslavedummy_value.html)
- [Linuxで「マスター／スレイブ」「ブラックリスト／ホワイトリスト」の語句置き換えが決定 - GIGAZINE](https://gigazine.net/news/20200713-linux-replace-master-slave/)
    - [Linuxカーネルでの「master/slave」と「blacklist」禁止、トーバルズ氏が承認 - ITmedia NEWS](https://www.itmedia.co.jp/news/articles/2007/13/news058.html)
- [アップルも「マスター」「スレーブ」などの用語置き換え--スタイルガイド更新 - ZDNet Japan](https://japan.zdnet.com/article/35156986/)
- [差別的な用語を排除へ--IBMとマイクロソフトの自発的な取り組み - CNET Japan](https://japan.cnet.com/article/35158622/)
- [GitHub、これから作成するリポジトリのデフォルトブランチ名が「main」に。「master」から「main」へ変更 － Publickey](https://www.publickey1.jp/blog/20/githubmainmastermain.html)
- [git clone しないでデフォルトブランチ名を取得する](https://zenn.dev/msmhrt/articles/0f530f16cf86fa0baeee)

## 参考図書

{{% review-paapi "B07Y29NV9P" %}} <!-- Virtue Signaling (洋書) -->
{{% review-paapi "4770501994" %}} <!-- ちびくろサンボ -->
