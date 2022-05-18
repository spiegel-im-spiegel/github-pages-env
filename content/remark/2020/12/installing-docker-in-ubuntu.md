+++
title = "Ubuntu に Docker を入れる"
date =  "2020-12-20T21:55:25+09:00"
description = "なんか以前より面倒くさくなってないか？"
image = "/images/attention/kitten.jpg"
tags = [ "install", "ubuntu", "docker" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

ちょっと思いついて [onatype-nexus-community/nancy][nancy] を [Docker] で動かそうと思ったのだが，メイン・マシンに入れてないことに気がついた。
そういや[実験用のノートPC]({{< ref "/remark/2019/12/install-ubuntu-to-second-hand-pc.md" >}} "中古 PC に Ubuntu 環境を導入する")に入れただけでメイン・マシンには入れてなかったか。

というわけで，改めて [Ubuntu] に [Docker] を入れようとしたのだが，なんか以前より面倒くさくなってないか？ 一応，覚え書きとして残しておこう。

{{< div-box type="markdown" >}}
**【2022-05-19 追記】** `apt-key` コマンドが非推奨になったためインストール手順が一部変わっている。
詳しくは以下を参照のこと。

- [ついに apt-key コマンドに「非推奨」のワーニングが]({{< ref "/remark/2022/05/apt-key-is-deprecated.md" >}})
{{< /div-box >}}

まず [Ubuntu] の APT 標準リポジトリに入ってるものは微妙に古くてダメぽい感じ。
なので，既に入っている場合はいったん消しておく。

```text
$ sudo apt remove docker docker-engine docker.io containerd runc
```

まっさらになったらリポジトリの登録から始める。

その前に以下のパッケージが未インストールなら先にインストールする。

```text
sudo apt install apt-transport-https ca-certificates curl software-properties-common
```

インストールできたら [Docker] インストール用の OpenPGP 公開鍵をインポートする[^dump1]。

[^dump1]: インポートする前に公開鍵の中身を [gpgpdump] で確認することができる。コマンドラインで `gpgpdump fetch https://download.docker.com/linux/ubuntu/gpg` とすればよい。宣伝でした（笑）

```text
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
```

一応，鍵指紋を確認しておくね。

```text
$ sudo apt-key fingerprint 0EBFCD88
pub   rsa4096 2017-02-22 [SCEA]
      9DC8 5822 9FC7 DD38 854A  E2D8 8D81 803C 0EBF CD88
uid           [  不明  ] Docker Release (CE deb) <docker@docker.com>
sub   rsa4096 2017-02-22 [S]
```

無駄にデカい鍵を使ってる気がするが，まぁいいや。
鍵指紋が `9DC8 5822 9FC7 DD38 854A  E2D8 8D81 803C 0EBF CD88` なら無問題。

そして，ようやくリポジトリの登録である。

```text
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
```

`stable` の他に `nightly` と `test` があるそうだが， [Docker] 自体で遊ぶわけではないので `stable` でええじゃろ。
`sudo apt update` で最新化しておくのをお忘れなく。

あとは普通にインストールすればよい。

```text
sudo apt install docker-ce docker-ce-cli containerd.io
```

[Docker] のインストールが成功すると daemon として起動する。
状態を確認するには以下のコマンドでOK。

```text
$ sudo systemctl status docker
● docker.service - Docker Application Container Engine
     Loaded: loaded (/lib/systemd/system/docker.service; enabled; vendor preset: enabled)
     Active: active (running) since Sun 2020-12-20 00:00:00 UTC; 21min ago
TriggeredBy: ● docker.socket
       Docs: https://docs.docker.com
   Main PID: 849 (dockerd)
      Tasks: 14
     Memory: 112.1M
     CGroup: /system.slice/docker.service
             └─849 /usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock
```

うんうん，ちゃんと動いてるね。

[Docker] を自身のアカウントで動かしたい場合は 

```text
$ sudo usermod -aG docker username
```

などとして docker グループに組み入れる。
ログインし直さないと有効にならないので注意。

ほんじゃあ，いよいよ [onatype-nexus-community/nancy][nancy] を動かしてみますかね。

```text
$ go list -json -m all | docker run --rm -i sonatypecommunity/nancy:latest sleuth -n
┏━━━━━━━━━━━━━━━┓
┃ Summary                      ┃
┣━━━━━━━━━━━━━┳━┫
┃ Audited Dependencies     ┃ 9┃
┣━━━━━━━━━━━━━╋━┫
┃ Vulnerable Dependencies  ┃ 0┃
┗━━━━━━━━━━━━━┻━┛
```

よーし，うむうむ，よーし。

## ブックマーク

- [Install Docker Engine on Ubuntu | Docker Documentation](https://docs.docker.com/engine/install/ubuntu/)
- [How To Install and Use Docker on Ubuntu 20.04 | DigitalOcean](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-20-04)
- [Ubuntuにdockerをインストールする - Qiita](https://qiita.com/tkyonezu/items/0f6da57eb2d823d2611d)
- [Dockerは非推奨じゃないし今すぐ騒ぐのをやめろ - Cloud Penguins](https://jaco.udcp.info/entry/2020/12/03/172843)

- [Go 依存パッケージの脆弱性検査]({{< ref "/golang/check-for-vulns-in-golang-dependencies.md" >}})

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Docker]: https://www.docker.com/ "Empowering App Development for Developers | Docker"
[nancy]: https://github.com/sonatype-nexus-community/nancy "sonatype-nexus-community/nancy: A tool to check for vulnerabilities in your Golang dependencies, powered by Sonatype OSS Index"
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"

## 参考図書

{{% review-paapi "B08PNMRXKN" %}} <!-- イラストでわかるDockerとKubernetes -->
