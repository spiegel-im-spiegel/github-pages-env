+++
title = "Java 19 がリリースされた"
date =  "2022-09-21T05:13:14+09:00"
description = "2023-03 までの短期サポート・バージョン"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "ubuntu", "plantuml" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

予定通り Java 19 がリリースされた。
2023-03 までの短期サポート・バージョンである。

- [OpenJDK JDK 19 GA Release](https://jdk.java.net/19/)
- [JDK 19](https://openjdk.org/projects/jdk/19/)
- [JDK 19 Release Notes](https://jdk.java.net/19/release-notes)
- [Overview (Java SE 19 & JDK 19)](https://docs.oracle.com/en/java/javase/19/docs/api/)

主な内容は以下の通り。

{{< fig-quote type="markdown" title="JDK 19" link="https://openjdk.org/projects/jdk/19/" lang="en" >}}
- 405: [Record Patterns (Preview)](https://openjdk.org/jeps/405)
- 422: [Linux/RISC-V Port](https://openjdk.org/jeps/422)
- 424: [Foreign Function & Memory API (Preview)](https://openjdk.org/jeps/424)
- 425: [Virtual Threads (Preview)](https://openjdk.org/jeps/425)
- 426: [Vector API (Fourth Incubator)](https://openjdk.org/jeps/426)
- 427: [Pattern Matching for switch (Third Preview)](https://openjdk.org/jeps/427)
- 428: [Structured Concurrency (Incubator)](https://openjdk.org/jeps/428)
{{< /fig-quote >}}

[OpenJDK] を APT で管理するメリットはないので，実行バイナリを[リリースページ](https://jdk.java.net/19/)から直接ダウンロードして配置する。
以下は完全手動でのインストール（笑）

```text
$ cd /usr/local/src
$ sudo curl -L "https://download.java.net/java/GA/jdk19/877d6127e982470ba2a7faa31cc93d04/36/GPL/openjdk-19_linux-x64_bin.tar.gz" -O
$ cd ..
$ sudo unlink java # 以前のバージョンの Java 環境がある場合
$ sudo tar xvf src/openjdk-19_linux-x64_bin.tar.gz
$ sudo ln -s jdk-19 java
$ java -version # すでに PATH が通っている場合
openjdk version "19" 2022-09-20
OpenJDK Runtime Environment (build 19+36-2238)
OpenJDK 64-Bit Server VM (build 19+36-2238, mixed mode, sharing)
```

LTS 版 Java バイナリが欲しいなら [Adoptium](https://adoptium.net/) で取得できる。

私としては [PlantUML] が動けばいいので，試しておく[^puml1]。

[^puml1]: 2021-08-22 に [PlantUML] [v1.2022.7](https://github.com/plantuml/plantuml/releases/tag/v1.2022.7) が[リリース](http://plantuml.com/changes)されている。 [PlantUML] の使い方等については拙文「[真面目に PlantUML]({{< ref "/remark/2018/12/plantuml-1.md" >}})」シリーズを参考にどうぞ。

{{< fig-img src="./factory-method-pattern.png" title="Factory Method Pattern" link="./factory-method-pattern.puml" width="1160" >}}

## Oracle Java のサポート期間

“[Oracle Java SE Support Roadmap](https://www.oracle.com/java/technologies/java-se-support-roadmap.html)” より。

| Release  | GA Date | Premier Support | Extended Support |
| -------- | ------- | --------------- | ---------------- |
| 11 (LTS) | 2018-09 | 2023-09         | 2026-09          |
| 17 (LTS) | 2021-09 | 2026-09         | 2029-09          |
| 18       | 2022-03 | 2022-09         | -                |
| 19       | 2022-09 | 2023-03         | -                |
| 20       | 2023-03 | 2023-09         | -                |
| 21 (LTS) | 2023-09 | 2028-09         | 2031-09          |

2022-09 時点で Premier Support が終了しているものは除いている。
Java 8 については [Adoptium](https://adoptium.net/) などで最新バイナリを取得可能。

ぶっちゃけ Java 17 の Premier Support が2026年まであるので，少なくとも LTS に関しては 17→21 への換装は進まないと思うなぁ。
11→21 はあるかもしれんけど。
Azure の Web Apps とか今年に入ってようやく Java 17 に対応したんだぜ。
そんで来年は21とか LTS の意味がねー！ 多分その次に出るであろう 25? まで保留だよな。

## ブックマーク

- [Java 19が正式リリース。より軽量な仮想スレッド、RISC-Vへの移植など新機能。1年後のJava 21が次のLTS版に － Publickey](https://www.publickey1.jp/blog/22/java_19risc-v1java_21lts.html)

[OpenJDK]: http://openjdk.java.net/
[AdoptOpenJDK]: https://adoptopenjdk.net/ "AdoptOpenJDK - Open source, prebuilt OpenJDK binaries"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "B07CKHR8C1" %}} <!-- Spring Data JPAプログラミング入門 -->
{{% review-paapi "4621303252" %}} <!-- Effective Java 第3版 -->
{{% review-paapi "B0893LQ5KY" %}} <!-- Spring Boot 2 入門 -->
