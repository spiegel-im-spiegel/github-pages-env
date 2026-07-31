+++
title = "Chromebook を導入する 6 — Git & Go コンパイラ"
date =  "2024-05-06"
description = "GitHub のリポジトリから Go のコードを取ってきて修正&コミット&プッシュするところまで。"
image = "/images/attention/tools.png"
tags = [ "chromebook", "linux", "gnupg", "openssh", "git", "golang" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

- [Chromebook を導入する 1]({{< ref "/remark/2024/03/chromebook-1.md" >}})
- [Chromebook を導入する 2 — Linux サブシステム]({{< ref "/remark/2024/04/chromebook-2.md" >}})
- [Chromebook を導入する 3 — GnuPG & OpenSSH]({{< ref "/remark/2024/04/chromebook-3.md" >}})
- [Chromebook を導入する 4 — Flatpak で Firefox を導入する]({{< ref "/remark/2024/04/chromebook-4.md" >}})
- [Chromebook を導入する 5 — APT で Firefox を導入する]({{< ref "/remark/2024/04/chromebook-5.md" >}})
- [Chromebook を導入する 6 — Git & Go コンパイラ]({{< ref "/remark/2024/05/chromebook-6.md" >}}) ← イマココ
- [Chromebook を導入する 7 — VS Code の導入]({{< ref "/remark/2024/05/chromebook-7.md" >}})

さて。
色々と脱線したが「[Chromebook を導入する 3]({{< ref "/remark/2024/04/chromebook-3.md" >}})」の続きでようやく git の設定を始めるですよ。
その前に Linux サブシステムの[日本語化]({{< ref "/remark/2024/04/chromebook-2.md#jp" >}})と [`pinentry-gnome3` の導入]({{< ref "/remark/2024/04/chromebook-3.md#pinentry" >}})を済ませておく。
以降はこれらの設定が終わってることを前提に書いていく。

## Git のセットアップ

[Git] は最初から入っているのだが，見ての通り最新バージョンではない（2024-05-06 時点の最新は 2.45.0）。

{{< fig-img src="./git-version.png" title="git version" link="./git-version.png" width="986" >}}

Ubuntu であれば [PPA から最新をインストール]({{< ref "/remark/2019/04/install-git-from-ppa.md" >}} "PPA から Git をインストールする")できるのだが， Debian にはそういうのはないらしい。
最新バージョンが欲しければソースコードをビルドして自前で用意しないとダメなようだ。
なんだかなぁ。
Debian 系ってそういう文化なのかね。

というわけで [git][Git] のバージョンアップは諦めて，とっととセットアップを済ませてしまおう。

最初に `~/.config/git/config` ファイルを作っておく。
中身は空でよい。

```text
$ touch ~/.config/git/config
```

こうしないと勝手に `~/.gitconfig` ファイルを作ってくれやがるので，ホームディレクトリが煩くなるのだ。

ファイルを作ったら初期設定を行う。
初期設定なんか滅多にやらないので忘れてしまってるよ（笑） まずは基本情報から。

```text
$ git config --global user.name "username"
$ git config --global user.email "username@example.com"
```

実際には自分の名前とEメールアドレスで置き換えること。

次は電子署名関連のの設定

```text
$ git config --global user.signingkey 4F56CCE34EB6819E
$ git config --global commit.gpgSign true
$ git config --global tag.gpgSign true
```

これでコミットおよびタグ付与の際に電子署名を付けてくれる。
`user.signingkey` の後ろの `4F56CCE34EB6819E` の部分は各自の電子署名鍵の鍵 ID で置き換えること。
鍵 ID は `gpg --list-keys` コマンドの出力で分かる。

以下の設定はお好みで。

```text
$ git config --global core.autocrlf input
$ git config --global init.defaultBranch main
```

私はリモートの [Git] リポジトリへのアクセスに ssh を使うことを想定している（「[Chromebook を導入する 3]({{< ref "/remark/2024/04/chromebook-3.md" >}})」を参照のこと）。
もし HTTPS アクセスを使う場合は credenntial 情報を管理するためのツールの導入や設定が必要になるかも知れない。

設定した内容は

```text
$ git config --global -l
```

で確認できる。

## Go コンパイラのインストール

Chromebook の Linux サブシステムには GCC や [Go] コンパイラは入ってない。
いずれも APT でインストールできるのだが APT の [Go] コンパイラはバージョンが古いので（そういうとこよ），[公式サイト](https://go.dev/dl/ "All releases - The Go Programming Language")からバイナリを取ってきて展開する。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.22.2.linux-amd64.tar.gz" -O
$ cd ..
$ sudo tar xvf src/go1.22.2.linux-amd64.tar.gz
$ sudo mv go go1.22.2
$ sudo ln -s go1.22.2 go
```

[Go] コンパイラへの PATH は `/etc/profile` に `/usr/local/go/bin` を追記しておけばいいだろう。
これで

```text
$ go version
go version go1.22.2 linux/amd64
```

と起動が確認できた。

## Git と Go コンパイラの動作を確認する

今回はあらかじめ GitHub に動作確認用のリポジトリを作っておいた。

- [GitHub - spiegel-im-spiegel/hello: "Hello World" for golang](https://github.com/spiegel-im-spiegel/hello)

これをローカルに `git clone` する。

{{< fig-img src="./hello-world-1.png" title="Hello World" link="./hello-world-1.png" width="986" >}}

`go run hello.go` 実行で “Hello World!” が表示されるところまで確認できた。
ソースコードはこんな感じ。

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```

このコードの “World” の部分を “Chromebook” に書き換えて GitHub にプッシュするところまでやってみる。

{{< fig-img src="./hello-world-2.png" title="Hello Chromebook" link="./hello-world-2.png" width="986" >}}

上手く行ったかな？ ちなみにこの時点ではエディタに vim を使っている。
GitHub サイトで確認してみよう。

{{< fig-img src="./commit-history.png" title="Hello Chromebook" link="./commit-history.png" width="1099" >}}

うんうん。
ちゃんとコミットに電子署名が付いてるね。
これでよーやく VS Code を導入する準備ができたよ。

## 【2024-05-12 追記】 Go の開発環境を外部ストレージに構築する {#goenv}

内部ストレージを十分備えている Chromebook なら以下の設定は不要だと思う。
たとえば [ASUS CX34 シリーズ](https://www.asus.com/jp/laptops/for-home/chromebook/asus-chromebook-plus-cx34-cx3402/ "ASUS Chromebook Plus CX34 (CX3402) | Chromebook | 法人向けノートパソコン | ノートパソコン | ASUS日本")なら十分なストレージ容量を持つ（これで税込8万円切るもんなー）。
まぁ，私の場合は14インチという時点で対象外だが。

[Go] の環境設定については以下の拙文が参考になる。

- [Go のモジュール管理【バージョン 1.17 改訂版】](https://zenn.dev/spiegel/articles/20210223-go-module-aware-mode)

「[Chromebook を導入する 2]({{< ref "/remark/2024/04/chromebook-2.md" >}} "Chromebook を導入する 2 — Linux サブシステム")」で

```text
$ ln -s /mnt/chromeos/removable/SD\ Card/linux ~/ws
```

てな感じに外部ストレージのディレクトリをシンボリックファイルで繋いだので， `~/.profile` に以下の内容を追記してみた。

```bash
if [ -d "$HOME/ws/go" ] ; then
    export GOPATH="$HOME/ws/go"
    export GOCACHE="$GOPATH/go-build"
    export GOBIN="$GOPATH/bin"
    PATH="$GOBIN:$PATH"
fi
```

この設定が効いてる状態で `go env` コマンドの結果を見てみる。

{{< fig-img src="./go-env.png" title="go env" link="./go-env.png" width="986" >}}

注目する変数は `GOPATH`, `GOBIN`, `GOCACHE`, `GOCMODACHE` あたり。
これらの設定が外部ストレージを指示していれば問題ない。

### 外部ストレージのファイルモード

外部ストレージのファイルにファイルを置くと何故か実行権限 `x` が付いてしまう。
しかも `chmod` コマンドで変更できないようだ。
ちなみに `chown` コマンドも効かない。

軽くググってみたら以下の記事を見つけた。

- [【Chromebook/Linux】共有フォルダ「/mnt/chromeos」やchmodが「Operation not permitted」エラーになる現象についてのメモ](https://did2memo.net/2020/10/30/chromebook-mnt-chromeos-chmod-operation-not-permitted/)

こちらは外部ストレージファイルを実行できないというものだが，私の環境を見ると `/etc/mtab` を見ると

```text
9p /mnt/chromeos 9p rw,nosuid,nodev,relatime,access=any,msize=65560,trans=fd,rfd=18,wfd=18 0 0
```

となっていて `noexec` が外されているようだ。
これによって `/mnt/chromeos` 以下の外部ストレージの全ファイルが実行可能ファイルになってしまうっぽい？

{{< fig-quote type="markdown" title="【Chromebook/Linux】共有フォルダ「/mnt/chromeos」やchmodが「Operation not permitted」エラーになる現象についてのメモ" link="https://did2memo.net/2020/10/30/chromebook-mnt-chromeos-chmod-operation-not-permitted/" >}}
いろいろ追加で調べている中で、/mnt/chromeosをマウントする際に指定されていたファイルシステム「9P（Plan 9 Filesystem Protocol）」（/etc/mtab に記載されていたファイルシステム）について調べつつ、Chromium OSのissue trackerを確認したところ、どうやらこの9p（p9）を利用することと、chownが利用できないことに関連があるようでした（「by design」とあるので、不具合というよりそもそもの仕様として、というようなイメージ）。
{{< /fig-quote >}}

これはちょっと取り扱い注意かも。
それにしても，なんで安全でない方に倒すかね。

念のため [git][Git] にも以下の設定を加えておく。

```text
$ git config --global core.filemode false
$ git config --global core.ignorecase false
```

これでうっかり実行可能ファイルとしてコミットしてしまうこともないだろう。
ちなみにファイル名の大文字小文字も区別しないので `core.ignorecase` も `false` にしておく（これは外部ストレージのフォーマットのせいかな）。

## ブックマーク

- [Git 初期設定 #Git - Qiita](https://qiita.com/ucan-lab/items/aadbedcacbc2ac86a2b3)
- [GitHub アカウントへの新しい SSH キーの追加 - GitHub Docs](https://docs.github.com/ja/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account)
- [GitHub アカウントに GPG キーを追加する - GitHub Docs](https://docs.github.com/ja/authentication/managing-commit-signature-verification/adding-a-gpg-key-to-your-github-account)

- [GnuPG チートシート（鍵作成から失効まで）]({{< ref "/openpgp/gnupg-cheat-sheet.md" >}})
- [Git Commit で OpenPGP 署名を行う]({{< ref "/openpgp/git-commit-with-openpgp-signature.md" >}})
- [OpenSSH の認証鍵を GunPG で作成・管理する]({{< ref "/openpgp/ssh-key-management-with-gnupg.md" >}})

[Git]: https://git-scm.com/
[Go]: https://go.dev/ "The Go Programming Language"

## 参考

{{% review-paapi "B0BKKF7JHV" %}} <!-- ASUS Chromebook -->
{{% review-paapi "B079MCPJGH" %}} <!-- カメラ 目隠し シャッター -->
{{% review-paapi "B08LMYWKZD" %}} <!-- Bluetooth 無線静音マウス -->
{{% review-paapi "B0CH2X5LBX" %}} <!-- micoroSD カード -->
{{% review-paapi "B07KJWYQJW" %}} <!-- ANKER PowerExpand USB メディアハブ -->
{{% review-paapi "B09BMPZ3BZ" %}} <!-- Chromebook仕事術 -->
{{% review-paapi "4295013498" %}} <!-- Linuxシステムの仕組み -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B096TH798S" %}} <!-- 改訂2版 わかばちゃんと学ぶ Git使い方入門 -->
