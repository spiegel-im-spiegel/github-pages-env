+++
title = "Java 13 がリリース"
date =  "2019-09-18T21:42:42+09:00"
description = "個人的に注目点はなし。"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "ubuntu", "plantuml" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

予定通り Java 13 がリリースされた。
[OpenJDK] および同系列の Java のみのショート・サイクルのバージョンアップである。

- [Java 13 / JDK 13: General Availability](https://mail.openjdk.java.net/pipermail/jdk-dev/2019-September/003335.html)

個人的に注目点はなし。
ミリ秒オーダーの GC とか今どきの処理系では遅すぎるだろう。

そういえば，自宅マシンを [Ubuntu] に換装して以降，初めてのメジャーバージョンアップか。
結局，仕事以外で Java でプログラミングを行うことは殆どなかったし，これからもしないと思う[^work1]。

[^work1]: 仕事以外で Java コードを書かなかったのは，うっかり守秘義務に抵触するコードを公開するのを避けるため。 Java 言語自身の問題ではない。これからも書かないというのは，今となっては「[Java はやめておけ]({{< ref "/remark/2016/07/java.md" >}})」と思うから（笑）

[OpenJDK] を APT で管理するメリットはないので，実行バイナリを[リリースページ](https://jdk.java.net/13/)から直接ダウンロードして配置する。

```text
$ cd /usr/local/src
$ sudo curl "https://download.java.net/java/GA/jdk13/5b8a42f3905b406298b72d750b6919f6/33/GPL/openjdk-13_linux-x64_bin.tar.gz" -O
$ cd ..
$ sudo unlink java # 以前のバージョンの Java 環境がある場合
$ sudo tar xvf src/openjdk-13_linux-x64_bin.tar.gz
$ sudo ln -s jdk-13 java
$ java -version # すでに PATH が通っている場合
openjdk version "13" 2019-09-17
OpenJDK Runtime Environment (build 13+33)
OpenJDK 64-Bit Server VM (build 13+33, mixed mode, sharing)
```

私としては [PlantUML] が動けばいいので，試しておく[^puml1]。

[^puml1]: 2019-09-14 に [PlantUML] V1.2019.10 が[リリース](http://plantuml.com/ja/changes)されている。

{{< fig-quote title="Go 言語で Factory Method Pattern を構成できるか" link="/golang/factory-method-pattern/" >}}
{{< fig-img src="./factory-method-pattern.png" link="./factory-method-pattern.puml" width="1165" >}}
{{< /fig-quote >}}

うむうむ。
ちゃんと動くな。

## ブックマーク

- [［速報］Java 13が登場。ZGCの改善やSwitch式の実現など新機能。 Oracle Code One 2019 － Publickey](https://www.publickey1.jp/blog/19/java_13zgcswitch_oracle_code_one_2019.html)
- [Oracle、「Java 13」を発表 ～GCの改良やテキストブロック構文の追加などの機能改善 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1207982.html)

- [Java 環境のリリースとサポートに関する覚え書き]({{< ref "/remark/2018/02/release-cycle-of-java-environment.md" >}})
- [結局 OpenJDK をインストールし直すことにした]({{< ref "/remark/2019/07/reinstalling-openjdk.md" >}})

[OpenJDK]: http://openjdk.java.net/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "B00I8AT1EU" %}} <!-- Java言語で学ぶリファクタリング入門 -->
{{% review-paapi "B00I8ATHGW" %}} <!-- 増補改訂版 Java言語で学ぶデザインパターン入門 -->
{{% review-paapi "B00I8AT1BS" %}} <!-- Java言語で学ぶデザインパターン入門 マルチスレッド編 -->
