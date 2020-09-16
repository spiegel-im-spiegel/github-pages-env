+++
title = "Java 15 がリリースされた"
date =  "2020-09-16T10:22:49+09:00"
description = "OpenJDK および同系列の Java のみのショート・サイクルのバージョンアップ"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "ubuntu", "plantuml" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

予定通り Java 15 がリリースされた。
[OpenJDK] および同系列の Java のみのショート・サイクルのバージョンアップである。

- [JDK 15](https://openjdk.java.net/projects/jdk/15/)
- [JDK 15 GA Release](https://jdk.java.net/15/)
- [JDK 15 Release Notes](https://jdk.java.net/15/release-notes)

[OpenJDK] を APT で管理するメリットはないので，実行バイナリを[リリースページ](https://jdk.java.net/15/)から直接ダウンロードして配置する。

```text
$ cd /usr/local/src
$ sudo curl -L "https://download.java.net/java/GA/jdk15/779bf45e88a44cbd9ea6621d33e33db1/36/GPL/openjdk-15_linux-x64_bin.tar.gz" -O
$ cd ..
$ sudo unlink java # 以前のバージョンの Java 環境がある場合
$ sudo tar xvf src/openjdk-15_linux-x64_bin.tar.gz
$ sudo ln -s jdk-15 java
$ java -version # すでに PATH が通っている場合
openjdk version "15" 2020-09-15
OpenJDK Runtime Environment (build 15+36-1562)
OpenJDK 64-Bit Server VM (build 15+36-1562, mixed mode, sharing)
```

私としては [PlantUML] が動けばいいので，試しておく[^puml1]。

[^puml1]: 2020-08-23 に [PlantUML] Version 1.2020.16 が[リリース](http://plantuml.com/changes)されている。 [PlantUML] の使い方等については拙文「[真面目に PlantUML]({{< ref "/remark/2018/12/plantuml-1.md" >}})」シリーズを参考にどうぞ。

{{< fig-img src="./factory-method-pattern.png" title="Factory Method Pattern" link="./factory-method-pattern.puml" width="1193" >}}

うむうむ。
ちゃんと動くな。

## [OpenJDK] のリポジトリが GitHub に移行

[OpenJDK] のリポジトリが GitHub に移行したそうだ。

- [openjdk/jdk: JDK main-line development](https://github.com/openjdk/jdk)

コード管理について，2019年の段階で Mercurial から git への移行が提案されていたらしい。
その上で，自前での管理ではなく，外部のサービスを使うことにしたようだ。

外部サービスとして GitHub を選択した理由は

{{< fig-quote type="markdown" title="OpenJDKのソースコード、GitHubへの移行を完了 － Publickey" link="https://www.publickey1.jp/blog/20/openjdkgithub.html" >}}
GitHub's performance is as good as or superior to other providers, it is the world's largest source-code hosting service (50 million users as of May 2020), and it has one of the most extensive APIs.

GitHubのパフォーマンスは他のプロバイダと同等かそれ以上であり、世界最大のソースコードホスティングサービス（2020年5月時点で5000万人のユーザー）でもあり、最も豊富なAPIを備えています。

GitHub's extensive API has enabled support for GitHub in many tools including text editors, IDEs, command-line tools, and graphical desktop clients.

GitHubの豊富なAPIは、テキストエディタ、IDE、コマンドラインツール、グラフィカルなデスクトップクライアントなど、多くのツールでのサポートを実現しています。
{{< /fig-quote >}}

とのこと。
さもありなん。

ところで [GitHub Discussions](https://github.blog/2020-05-06-new-from-satellite-2020-github-codespaces-github-discussions-securing-code-in-private-repositories-and-more/ "New from Satellite 2020: GitHub Discussions, Codespaces, securing code in private repositories, and more - The GitHub Blog") はいつ一般に開放されるのだろう。
開放されるなら，Disqus なんか捨てて，フィードバックはそっちに移行したいんだけどなぁ。
それとも個人ユーザには開放されないとか？

[OpenJDK]: http://openjdk.java.net/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "B00I8AT1EU" %}} <!-- Java言語で学ぶリファクタリング入門 -->
{{% review-paapi "B00I8ATHGW" %}} <!-- 増補改訂版 Java言語で学ぶデザインパターン入門 -->
{{% review-paapi "B00I8AT1BS" %}} <!-- Java言語で学ぶデザインパターン入門 マルチスレッド編 -->
