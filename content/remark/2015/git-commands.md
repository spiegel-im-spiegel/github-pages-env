+++
date = "2015-12-09T20:14:59+09:00"
update = "2016-11-16T13:06:24+09:00"
description = "git 各コマンドに関する覚え書き。思い出したら追記予定。"
draft = false
tags = ["git", "tools"]
title = "あまり使わないけど，たまに使おうとすると忘れてる Git コマンド集"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

[git] 各コマンドに関する覚え書き。
思い出したら追記予定。

## Submodule の作成から削除まで

### Submodule の作成

Repository に別の repository `sub_repo.git` を submodule として `subdir` に追加する場合は以下のコマンドを実行する。

```bash
$ git submodule add sub_repo.git subdir
```

### Submodule の初期化

Submodule を含む repository を clone してきた場合は

```bash
$ git clone repo.git
$ cd repo/subdir
$ git submodule init
$ git submodule update
```

または

```bash
$ git clone --recursive repo.git
```

で初期化できる。
これ以降， remote 側から fetch/maerge した際に submodule の参照ポイントが変更されていた場合は

```bash
$ git submodule update --init --recursive
```

で初期化できる[^sub1]。

[^sub1]: `--recursive` オプションなしで cloneした 直後は `git update` では初期化できないので注意。

### Submodule の更新

Submodule の remote 側の変更を pull する場合は以下のコマンドを実行する。

```bash
$ git submodule update --remote subdir
```

Fetch したデータを merge する場合は `--merge` オプションを， rebase する場合は `--rebase` オプションを付ける[^rb]。

[^rb]: Rebase は歴史の改変なので取り扱いに注意。

### Submodule の削除

Repository から submodule を削除する場合は以下のコマンドで削除できる。

```bash
$ git submodule deinit subdir
$ git rm subdir
```

## Remote Repository の接続設定

### Remote Repository への URI を変更する

Remote repository （通常は `origin`）の接続先 URI を変更するには以下のコマンドを実行する。

```bash
$ git remote set-url origin new_repo.git
```

### Remote Repository との接続を追加する

新たに remote repository との接続を追加する場合には以下のコマンドを実行する。

```bash
$ git remote add upstream up_repo.git
```

ここでは追加した remote repository に `upstream` と名前をつけている。
たとえば fork した repository で作業する際に fork 元の repository の変更も取り込みたい場合などに有効である。

### Remote Repository との接続を削除する

Remote repository との接続を削除する場合には以下のコマンドを実行する。

```bash
$ git remote rm upstream
```

## Repository の分離

Repository の特定のディレクトリ `subdir` を commit tree を維持したまま分離したいときの手順。

まず元の repository `org_repo.git` を clone する。

```bash
$ git clone org_repo.git
```

Clone 元の bare repository がない場合にはローカルの repository を丸ごとどっかにコピーすればよい（commit 済みであること）。

で， clone した repository 内に入って `git filter-branch` コマンドを実行する。

```bash
$ cd org_repo
$ git filter-branch --subdirectory-filter subdir HEAD
Rewrite **************************************** (999/999)
Ref 'refs/heads/master' was rewritten
```

このとき `subdir` 以下のファイル・ディレクトリが repository のトップ・ディレクトリに移動するので注意が必要である。
また，オリジナルの commit tree が `refs/original/refs/heads/master` として残っているので以下のコマンドで削除する。

```bash
$ git update-ref -d refs/original/refs/heads/master
```

作成した repository はそのまま使うなり新しい bare repository に push するなりすればいい。

```bash
$ git remote set-url origin new_repo.git
$ git push -u origin master
```

元の repository に push しようとすると怒られる[^sf]。

[^sf]: `-f` オプションを付けて強制的に push することは可能。ただしこれは（rebase と同じで）歴史の改変になるため取り扱いに注意。

## Remote Repository の branch を削除する

[GitHub] とかで余計な branch を作っちゃって削除したい時。
ローカルの clone から以下の引数で push すればいいようだ。

```bash
$ git push origin :wrong-branch
```

- [passingloop • 復習 Git: GitHub のブランチを削除する．](http://passingloop.tumblr.com/post/18015576208/remove-remote-branches-by-git-push)
- [GitHubで付けたタグを削除する - アインシュタインの電話番号](http://blog.ruedap.com/2011/02/08/git-github-delete-tag)

## commit の取り消し（しかもリモートに push しちゃった）

まず commit を完全になかったことにするには以下のコマンドを実行する。

```text
$ git reset --hard xxxxxxx
```

`xxxxxxx` は commit ID （ハッシュ）で，取り消すコミットの直前のものを指定する。
これを push すればいいのだが，単に push しただけではエラーになるので

```text
$ git push -f
```

として強制的に push する。
これで貴方の歴史はなかったことになりました（笑） チームで管理している場合はご注意を。

- [【git】git pushを取り消す - tweeeetyのぶろぐ的めも](http://tweeeety.hateblo.jp/entry/2015/06/10/215753)

## Git に関するブックマーク

ついでなので，ローカルで溜め込んでいた [git] に関するブックマークを挙げておく。

- [Git for Windows](https://git-for-windows.github.io/)
    - [Git for Windows で日本語を使いたい - Qiita](http://qiita.com/kumazo@github/items/2169e1ee7be278f82b94)
- [Git Extensions](http://gitextensions.github.io/)
- [【Git入門者向け】イメージで理解するGitコマンド事始め | きのこる庭](http://kinokoru.jp/archives/1017)
- [git pull と git pull –rebase の違いって？図を交えて説明します！ | KRAY Inc](http://kray.jp/blog/git-pull-rebase/)
    - [Git - pull は本当に fetch + merge なの？ - Qiita](http://qiita.com/Teloo/items/95a860ae276b49edb040)
- [Gitのこれやめて！リスト - Qiita](http://qiita.com/doilux/items/b5a9abd95ac91e848a5f)
- [Git LFSが1.0になってGitHubで使えるようになったので試してみた - Qiita](http://qiita.com/kiida/items/0d51c43ac73f14f09f5a)
- [git-lfsは大容量のファイルを扱うもので多量のファイルを扱うものではない - Qiita](http://qiita.com/crifff/items/32ffc824f69ed5632217)
- [gitで重いリポジトリをcloneするとき - webネタ](http://r-h.hatenablog.com/entry/2013/12/07/093423) （[Qiita 版](http://qiita.com/butchi_y/items/cc0fe50acc47c1e3ab32)）
- [Git - サブモジュール](https://git-scm.com/book/ja/v2/Git-%E3%81%AE%E3%81%95%E3%81%BE%E3%81%96%E3%81%BE%E3%81%AA%E3%83%84%E3%83%BC%E3%83%AB-%E3%82%B5%E3%83%96%E3%83%A2%E3%82%B8%E3%83%A5%E3%83%BC%E3%83%AB)
- [Gitのsubmoduleをお手軽に削除する - Steel Dragon 14106](http://raimon49.github.io/2015/04/04/git-submodule-deinit.html)
- [Gitリポジトリ中のサブディレクトリを別のリポジトリにする - 北海道苫小牧市出身のPGが書くブログ](http://d.hatena.ne.jp/hiratara/20091112/1258023732)
- [Gitリポジトリのディレクトリ構成を変えて別リポジトリにする - ごずろぐ](http://gozuk16.hatenablog.com/entry/2015/04/24/145714)
- [Git リポジトリに上がっているファイルを履歴ごと消すには？ - Qiita](http://qiita.com/go_astrayer/items/6e39d3ab16ae8094496c)
- [マージされてない他者のプルリクを取り込む - Qiita](http://qiita.com/hirogw/items/3ea3a321a367740e971a)
- [git commit時に英語でメッセージ書くためのヒントを表示する - Qiita](http://qiita.com/sagaraya/items/60e983856e16dc571f22)
- [【Git】コミットに規約をつくる - Qiita](http://qiita.com/Kenya/items/f72fba8fecc79d1b090c)
- [git reset のやり方 備忘録 - Qiita](http://qiita.com/waterada/items/44f270a659370809e1dc)
- [気をつけて！Git for Windowsにおける改行コード - Qiita](http://qiita.com/uggds/items/00a1974ec4f115616580) : `autocrlf` の値は `input` がお勧め
- [git - 簡単ガイド](http://rogerdudler.github.io/git-guide/index.ja.html)
- [Git Subtree 事始め - Qiita](http://qiita.com/mikakane/items/487ca8b3acddfa5fdb41)
    - [git-subtree移行メモ - Qiita](http://qiita.com/shogo82148/items/04b29b195dbbb373152f)
- [たった 8 つの alias で git を大幅にわかりやすくする - Qiita](http://qiita.com/ticonz/items/2f5d5299e3562e6941a8)
- [Git で履歴のAuthor/Emailを上書きする - Qiita](http://qiita.com/kenju/items/196e8d5df38f19ec2dd3)

### GitHub

- [githubからclone時にerror setting certificate verify locationsがでる | MemeTodo](http://meme.efcl.info/2011/07/gitcloneerror-setting-certificate.html) : うちでもなっていろいろ大変だった
- [GitHub にパスワードとかセンシティブなファイルを push してしまったときの対処法 - Qiita](http://qiita.com/dtan4/items/34e41e3bd40a43fd8cbf) : GitHub などリモートのリポジトリに上げてしまうと，とてつもなく面倒になるので，要注意
- [Github のブランチ保護を使用してリスキーなマージを防止する - Qiita](http://qiita.com/yo1000/items/8ffe225716ba3b064697)
- [Using SSH over the HTTPS port - User Documentation](https://help.github.com/articles/using-ssh-over-the-https-port/)

### Git-flow and GitHub-flow

- [GitHub Flow – Scott Chacon](http://scottchacon.com/2011/08/31/github-flow.html)
- [GitHub Flow (Japanese translation)](https://gist.github.com/Gab-km/3705015)
- [A successful Git branching model » nvie.com](http://nvie.com/posts/a-successful-git-branching-model/)
- [git-flow cheatsheet](http://danielkummer.github.io/git-flow-cheatsheet/) （[日本語](http://danielkummer.github.io/git-flow-cheatsheet/index.ja_JP.html)）
- [git flowとgithub flowざっくりまとめ | KentaKomai Blog](http://komaken.me/blog/2013/09/09/git-flow%E3%81%A8github-flow%E3%81%96%E3%81%A3%E3%81%8F%E3%82%8A%E3%81%BE%E3%81%A8%E3%82%81/)
- [GitHub初心者はForkしない方のPull Requestから入門しよう | qnyp blog](http://blog.qnyp.com/2013/05/28/pull-request-for-github-beginners/)
- [Git for Windows 2.xのシステムコンフィグファイルは2つある - Qiita](http://qiita.com/anqooqie/items/07367b2b5932f3acfc38)
- [小規模開発のgit-flowの導入を楽にするブランチルールと拡張スクリプト配布 - Qiita](http://qiita.com/eaglesakura/items/4fa17bf1f2b6683e6520)

[git]: https://git-scm.com/ "Git"
[GitHub]: https://github.com/
