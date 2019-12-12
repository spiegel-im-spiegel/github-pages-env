+++
title = "#shimanego より： Docker 上で Go 言語コードを実行する"
date =  "2019-12-12T20:29:12+09:00"
description = "description"
image = "/images/attention/kitten.jpg"
tags = [ "remark" ]
pageType = "text"
draft = true

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Shimane.go#02] に参加してきた。
年末でめっさ忙しく参加できないんじゃないかと思ったよ。

[Shimane.go#02] では「go-lang on docker」タイトルで LT が行われた。
まぁ [Go 言語]未経験の人も多いので軽くジャブからというところだろうか。

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

次に [Docker] インストール用のリポジトリを追加する。

```text
$ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
$ sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
```

ただし，現時点（2019-12-12）では [Ubuntu] 19.10 用のリポジトリは用意されていないので， ひとつ前（19.04）のリポジトリでお茶を濁しておく。

```text
$ sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu disco stable"
```

ここまでくればあとは普通に `apt install` すればよい。

```text
$ sudo apt update
$ sudo apt install docker-ce docker-ce-cli containerd.io
```

ここで動作確認しておこう。

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





[connpass]: https://connpass.com/ "connpass - エンジニアをつなぐIT勉強会支援プラットフォーム"
[Shimane.go#02]: https://shimane-go.connpass.com/event/156445/ "Shimane.go#02 - connpass"
[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Docker]: https://www.docker.com/ "Enterprise Container Platform | Docker"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN), 4621300253 (ISBN), 9784621300251 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
