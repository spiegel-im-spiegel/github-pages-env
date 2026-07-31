+++
title = "Java 18 がリリースされた"
date =  "2022-03-23T20:32:56+09:00"
description = "2022-09 までの短期サポート・バージョン"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "ubuntu", "plantuml" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

予定通り Java 18 がリリースされた。
2022-09 までの短期サポート・バージョンである。

- [OpenJDK JDK 18 GA Release](https://jdk.java.net/18/)
- [JDK 18](https://openjdk.java.net/projects/jdk/18/)
  - [Java 18 / JDK 18: General Availability](https://mail.openjdk.java.net/pipermail/jdk-dev/2022-March/006458.html)
- [JDK 18 Release Notes](https://jdk.java.net/18/release-notes)
- [Overview (Java SE 18 & JDK 18)](https://docs.oracle.com/en/java/javase/18/docs/api/)

主な内容は以下の通り。

{{< fig-quote type="markdown" title="JDK 18" link="https://openjdk.java.net/projects/jdk/18/" lang="en" >}}
- 400: [UTF-8 by Default](https://openjdk.java.net/jeps/400)
- 408: [Simple Web Server](https://openjdk.java.net/jeps/408)
- 413: [Code Snippets in Java API Documentation](https://openjdk.java.net/jeps/413)
- 416: [Reimplement Core Reflection with Method Handles](https://openjdk.java.net/jeps/416)
- 417: [Vector API (Third Incubator)](https://openjdk.java.net/jeps/417)
- 418: [Internet-Address Resolution SPI](https://openjdk.java.net/jeps/418)
- 419: [Foreign Function & Memory API (Second Incubator)](https://openjdk.java.net/jeps/419)
- 420: [Pattern Matching for switch (Second Preview)](https://openjdk.java.net/jeps/420)
- 421: [Deprecate Finalization for Removal](https://openjdk.java.net/jeps/421)
{{< /fig-quote >}}

[OpenJDK] を APT で管理するメリットはないので，実行バイナリを[リリースページ](https://jdk.java.net/18/)から直接ダウンロードして配置する。
以下は完全手動でのインストール（笑）

```text
$ cd /usr/local/src
$ sudo curl -L "https://download.java.net/java/GA/jdk18/43f95e8614114aeaa8e8a5fcf20a682d/36/GPL/openjdk-18_linux-x64_bin.tar.gz" -O
$ cd ..
$ sudo unlink java # 以前のバージョンの Java 環境がある場合
$ sudo tar xvf src/openjdk-18_linux-x64_bin.tar.gz
$ sudo ln -s jdk-18 java
$ java -version # すでに PATH が通っている場合
openjdk version "18" 2022-03-22
OpenJDK Runtime Environment (build 18+36-2087)
OpenJDK 64-Bit Server VM (build 18+36-2087, mixed mode, sharing)
```

私としては [PlantUML] が動けばいいので，試しておく[^puml1]。

[^puml1]: 2021-03-05 に [PlantUML] V1.2022.2 が[リリース](http://plantuml.com/changes)されている。 [PlantUML] の使い方等については拙文「[真面目に PlantUML]({{< ref "/remark/2018/12/plantuml-1.md" >}})」シリーズを参考にどうぞ。

{{< fig-img src="./factory-method-pattern.png" title="Factory Method Pattern" link="./factory-method-pattern.puml" width="1160" >}}

おりょ？ [PlantUML] のバージョンが上がったせいか？ 角がちょっぴり丸くなってるな。
まぁ，いいか。
ちゃんと動くし。

## ブックマーク

- [Java 18正式リリース。デフォルトのCharsetが「UTF-8」に、シンプルWebサーバ搭載など新機能 － Publickey](https://www.publickey1.jp/blog/22/java_18charsetutf-8web.html)
- [AWS、Java 18対応の独自Javaディストリビューション「Amazon Corretto 18」正式リリース － Publickey](https://www.publickey1.jp/blog/22/awsjava_18javaamazon_corretto_18.html)

- [Oracle Java SE Support Roadmap](https://www.oracle.com/java/technologies/java-se-support-roadmap.html)
- [Support | AdoptOpenJDK - Open source, prebuilt OpenJDK binaries](https://adoptopenjdk.net/support.html)

[OpenJDK]: http://openjdk.java.net/
[AdoptOpenJDK]: https://adoptopenjdk.net/ "AdoptOpenJDK - Open source, prebuilt OpenJDK binaries"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "B07CKHR8C1" %}} <!-- Spring Data JPAプログラミング入門 -->
{{% review-paapi "4621303252" %}} <!-- Effective Java 第3版 -->
{{% review-paapi "B0893LQ5KY" %}} <!-- Spring Boot 2 入門 -->
