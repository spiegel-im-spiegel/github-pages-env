+++
date = "2015-12-09T15:10:27+09:00"
description = "git 各コマンドに関する覚え書き。思い出したら追記予定。"
draft = true
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
2回目以降， submodule の参照ポイントが変更された場合は

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

```bash
$ git remote set-url origin new_repo.git
```

### Remote Repository との接続を追加する

```bash
$ git remote add upstream up_repo.git
```

### Remote Repository との接続を削除する

```bash
$ git remote rm upstream up_repo.git
```

## Repository の分離

Repository の特定のディレクトリ `subdir` を commit tree を維持したまま分離したい場合。

まず元の repository `org_repo.git` を clone する。

```bash
$ git clone org_repo.git
```

Clone 元の bare repository がない場合にはローカルの repository を丸ごとどっかにコピーすればよい（commit 済みであること）。

で， clone した repository 内に入って以下のコマンドを入力する。

```bash
$ cd org_repo
$ git filter-branch --subdirectory-filter subdir HEAD
Rewrite **************************************** (999/999)
Ref 'refs/heads/master' was rewritten
```

このとき `subdir` 以下のファイル・ディレクトリが repository のトップ・ディレクトリに移動するので注意が必要。
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

## ブックマーク

- [Git - サブモジュール](https://git-scm.com/book/ja/v2/Git-%E3%81%AE%E3%81%95%E3%81%BE%E3%81%96%E3%81%BE%E3%81%AA%E3%83%84%E3%83%BC%E3%83%AB-%E3%82%B5%E3%83%96%E3%83%A2%E3%82%B8%E3%83%A5%E3%83%BC%E3%83%AB)
- [Gitのsubmoduleをお手軽に削除する - Steel Dragon 14106](http://raimon49.github.io/2015/04/04/git-submodule-deinit.html)
- [Gitリポジトリ中のサブディレクトリを別のリポジトリにする - 北海道苫小牧市出身のPGが書くブログ](http://d.hatena.ne.jp/hiratara/20091112/1258023732)
- [Gitリポジトリのディレクトリ構成を変えて別リポジトリにする - ごずろぐ](http://gozuk16.hatenablog.com/entry/2015/04/24/145714)
- [Git リポジトリに上がっているファイルを履歴ごと消すには？ - Qiita](http://qiita.com/go_astrayer/items/6e39d3ab16ae8094496c)

[git]: https://git-scm.com/ "Git"