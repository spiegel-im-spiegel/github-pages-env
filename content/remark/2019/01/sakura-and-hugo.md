+++
title = "さくらのレンタルサーバ上で Hugo によるサイト管理を行う"
date = "2019-01-21T23:03:56+09:00"
description = "サーバ側で動く CMS なんか今さら使う気にもならないし，でもこれで多少は楽にサイト管理ができそうである。"
image = "/images/attention/kitten.jpg"
tags = [ "site", "web", "hugo" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

前回， [Flickr から引き揚げた写真を展示するためにサイトを公開した話]({{< relref "escape-from-flickr.md" >}} "Flickr から写真を引き揚げました")をしたが，[このサイト](https://photo.baldanders.info/ "photo.Baldanders.info")は静的サイト・ジェネレータである [Hugo] で管理を行っている。
今回は「[さくらのレンタルサーバ]」上で [Hugo] によるサイト管理を行う手順について覚え書きとして記しておく。

## 「[さくらのレンタルサーバ]」について

まず「[さくらのレンタルサーバ]」でサイトを公開するまでの手順については以下のページが参考になる。

- [rinopo/sakura-init: さくらのレンタルサーバを借りたとき最初にすること](https://github.com/rinopo/sakura-init)

リンク先の記事は特にマルチドメインでサイトを運用する際に有用な情報となるだろう。
今回の記事では `~/www/` ディレクトリ直下にサイトを公開する場合の手順について書くが，マルチドメインでもそれほど変わらないはずである。

前提として「[さくらのレンタルサーバ]」上で [git] が動作する必要がある。
昔は自前で [git] をインストールする必要があったが，今は予め用意されているようである。
ただし，私が利用しているサーバではエラく古いバージョンだった（なんで？）。

```text
$ which git
/usr/local/bin/git

$ git version
git version 2.7.0
```

しかも git-submodule すら入っていない。

```text
$ git submodule version
fatal: git was built without support for git-submodule (NO_PERL=1).
```

というわけで，サイト管理用のリポジトリを構成する場合には注意が必要である。

## [Hugo] 用にベア・リポジトリを作成する

まずは [Hugo] 用にベア・リポジトリを「[さくらのレンタルサーバ]」上に作成する。

```text
$ git init --bare ~/repos/hugo-env.git
```

これで `~/repos/hugo-env.git` ディレクトリに [Hugo] 用のベア・リポジトリが作成される。
当然ながら `~/www/` ディレクトリ以下には作らないこと。

このベア・リポジトリをリモートから clone するには以下のコマンドラインでよい（`example.com` ドメインで構成している場合[^sakura1]）。

```text
[client] $ git clone ssh://example.com/home/username/repos/hugo-env.git
```

[^sakura1]: 独自ドメインを使わない場合は `ssh://username@username.sakura.ne.jp/home/...` のようにドメイン名の前に必ず「ユーザ名@」を付記すること。

あとは任意の環境でコンテンツを記述できる。
[Hugo] にはサーバモード（簡易サーバとして構成することも可能）があるので

```text
[client] $ hugo server -D -w
```

などとするとよい。
`-D` (または `--buildDrafts`) オプションはドラフト・ページを含めてビルドしてくれる。
また `-w` (または `--watch`) オプションを付けておけば [Hugo] 環境内のファイルが更新されるたびに自動的にリビルドしてくれる。

サーバモードは本当に便利で，特に [Hugo] のサーバモードは一瞬でサイトを更新してくれるので，ブラウザで出来上がりをリアルタイムで見ながら編集できる[^ll1]。
これのおかげでテキスト・エディタの markdown previewer 機能が要らなくなりました（笑）

[^ll1]: ブラウザ側は [livereload-js](https://github.com/livereload/livereload-js "livereload/livereload-js: LiveReload JavaScript code that communicates with the server and implements reloading") で自動的に更新しているようだ。

## サーバ側に Deploy 用のリポジトリを作成する

[Hugo] でビルドしたコンテンツを deploy できるように，サーバ側でリポジトリを clone する。

```text
$ git clone ~/hugo-env.git
```

これも当然だが `~/www/` ディレクトリ以下には作らないこと。

更に [Hugo] の実行バイナリを `$PATH` の通っているディレクトリに入れる。
[Hugo] の実行バイナリは[リリースページ](https://github.com/gohugoio/hugo/releases "Releases · gohugoio/hugo")で取得できる。

「[さくらのレンタルサーバ]」は FreeBSD で構成されているので `hugo_x.xx_FreeBSD-64bit.tar.gz` (`x.xx` はバージョン番号) をダウンロードし中にある `hugo` ファイルを使えばよい。 
Go 製ツールはシングル・バイナリで動くのが素敵。

Build & deploy は `hugo` コマンド一発でできる。

```text
$ cd ~/hugo-env
$ hugo -d ~/www
```

これでビルドした結果をまるっと `~/www/` ディレクトリに出力してくれる。

## Deploy の自動化

コンテンツを修正するたびにサーバに入って手動で pull & build & deploy というのも芸がないので，ベア・リポジトリに対して自動化の設定を行う。

具体的にはベア・リポジトリの `hooks/` ディレクトリに以下の内容で `post-update` を追加する（実行権限の付与を忘れないこと）。

```bash
#!/bin/sh
cd /home/username/hugo-env || exit
unset GIT_DIR
git pull origin master || exit
hugo -d /home/username/www
```

これでリモートからの push でベア・リポジトリの内容が更新されるたびに上の shell スクリプトが走って deploy まで自動的に行ってくれる。
今回作成した [photo.Baldanders.info](https://photo.baldanders.info/) は regular page で2,900ページ余りあるのだが，上書き更新なら5秒前後でビルドが完了する。
不安なら

```bash
hugo -d /home/username/www &> /dev/nul &
```

とかでもいいかも。

本当はコンテナを起動して云々とかやれば CI (Continuous Integration) ぽくてカッコよかったかも知れないが，この構成でも特に不自由しないし，問題ないだろう。

「[さくらのレンタルサーバ]」はスタンダード・プランでもストレージが100GBもあるので「もっとリッチに使えれば」と常々思っていたが，[サーバ側で動く CMS なんか今さら使う気にもならない]({{< ref "/remark/2016/07/cms.md" >}})し，でもこれで多少は楽にサイト管理ができそうである。

## ブックマーク

- [Hugoサイト構築 | Watanabe-DENKI Inc. 渡辺電気株式会社](https://wdkk.co.jp/lab/hugo/) : [Hugo] の基本的な使い方はこちらがオススメ
- [ゼロから始める Hugo]({{< rlnk "/hugo/" >}}) : 拙文ご容赦

- [さくらのレンタルサーバ上でJekyllブログを公開 | Dev SandBox](http://utwang.io/2013/01/04/jekyll-on-sakura/)

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[さくらのレンタルサーバ]: https://www.sakura.ne.jp/ "さくらのレンタルサーバ | 高速・安定WordPressなら！無料2週間お試し"
[git]: https://git-scm.com/
[Git]: https://git-scm.com/
<!-- eof -->
