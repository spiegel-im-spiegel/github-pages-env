+++
title = "#shimanego より： Docker 上で Go 言語コードを実行する"
date =  "2019-12-12T22:56:50+09:00"
description = "いちから Docker を勉強してみるかな。"
image = "/images/attention/kitten.jpg"
tags = [ "golang", "communication", "shimane", "engineering", "docker" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Shimane.go#02] に参加してきた。
年末でめっさ忙しく，参加できないんじゃないかと思ったよ。
辿り着けてよかった。

[Shimane.go#02] では「go-lang on docker」というタイトルで LT が行われた。
まぁ [Go 言語]未経験の人も多いし，本格的な活動は年が明けてからということなので，今回は軽いジャブというところだろうか。

そういや，[玩具用のパソコン買った]({{< ref "remark/2019/12/install-ubuntu-to-second-hand-pc.md" >}})ら [Docker] で遊ぼうと思ってたっけ。
ちょうどいいや。

## [Docker] のインストール

LT では [Docker] が使える前提で解説されていたが，私はまずインストールするところから。
[Ubuntu] へのインストールは以下が参考になる。

- [Get Docker Engine - Community for Ubuntu | Docker Documentation](https://docs.docker.com/install/linux/docker-ce/ubuntu/)

まずはインストールに必要なパッケージを APT でインストールする。

```text
$ sudo apt install apt-transport-https ca-certificates curl software-properties-common
```

既にインストール済みのパッケージについてはちゃんとスキップしてくれるので大丈夫。

次に [Docker] インストール用のリポジトリと署名検証用の OpenPGP 公開鍵を追加する。

```text
$ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
$ sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
```

ただし，現時点（2019-12-12）では [Ubuntu] 19.10 用のリポジトリは用意されていないので， ひとつ前（19.04）のリポジトリを強制的に追加してお茶を濁しておく。

```text
$ sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu disco stable"
```

[Ubuntu] 19.10 用のリポジトリが出たら入れ替えないと。

ここまでくれば，あとは普通に `apt install` すればよい。

```text
$ sudo apt update
$ sudo apt install docker-ce docker-ce-cli containerd.io
```

ここで動作確認しておこう。
みんな大好き “Hello World” （笑）

```text
$ sudo docker run hello-world
Unable to find image 'hello-world:latest' locally
latest: Pulling from library/hello-world
1b930d010525: Pull complete 
Digest: sha256:4fe721ccc2e8dc7362278a29dc660d833570ec2682f4e4194f4ee23e415e1064
Status: Downloaded newer image for hello-world:latest

Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker ID:
 https://hub.docker.com/

For more examples and ideas, visit:
 https://docs.docker.com/get-started/
```

おおっ，動いた！

## docker グループを追加する。

上の実行例を見れば分かるが，インストール直後は `root` 以外のユーザには `docker` コマンドを動かす権限がない。
ユーザ `username` に権限を付与するには `docker` グループを追加すればよい。
こんな感じ。

```text
$ sudo usermod -aG docker username
```

これでログインし直せば[^login1]

[^login1]: 何故か私の環境ではログインし直しただけではダメで再起動する必要があった。なんで？ これのせいでしばらくハマっちゃったよ。今日のイベントで私の周囲の人はブツブツ独り言をいうおっさんがいてさぞ気持ち悪かったことだろう。ゴメンペコン

```
$ docker run hello-world
```

としても “permission denied” と怒られることはなくなる。

## Docker 上で Go 言語コードを実行する

LT ではサンプルコードとして以下のリポジトリを紹介された。

- [gitkado/docker-run-go-sample](https://github.com/gitkado/docker-run-go-sample)

このリポジトリを取ってきて

```text
$ docker run --rm -v $PWD:/go golang:latest go run sample.go
...
HelloWorld!
```

となれば成功！

思ったより取っつきやすいな，[Docker]。
もっと面倒くさいのかと思ってた。
まぁ，ネットワーク周りを弄りだすと面倒なんだろうけど。

いちからちゃんと勉強してみるかな。

## ブックマーク

- [DockerをLinux(Ubuntu 14.04 LTS)にインストールする方法と解説 | tracpath:Works](https://tracpath.com/works/devops/how_to_install_the_docker/)
- [Ubuntu 19.10にDockerをインストールする - Qiita](https://qiita.com/rarudonet/items/8c5e99f12adc85c73729)

[connpass]: https://connpass.com/ "connpass - エンジニアをつなぐIT勉強会支援プラットフォーム"
[Shimane.go#02]: https://shimane-go.connpass.com/event/156445/ "Shimane.go#02 - connpass"
[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Docker]: https://www.docker.com/ "Enterprise Container Platform | Docker"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
