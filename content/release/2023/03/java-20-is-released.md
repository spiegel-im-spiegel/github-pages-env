+++
title = "Java 20 がリリースされた"
date =  "2023-03-22T18:56:43+09:00"
description = "2023-09 までの短期サポート・バージョン"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "ubuntu", "plantuml" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

予定通り Java 19 がリリースされた。
2023-03 までの短期サポート・バージョンである。

- [OpenJDK JDK 20 GA Release](https://jdk.java.net/20/)
- [JDK 20](https://openjdk.org/projects/jdk/20/)
- [JDK 20 Release Notes](https://jdk.java.net/20/release-notes)

主な内容は以下の通り。

{{< fig-quote type="markdown" title="JDK 20" link="https://openjdk.org/projects/jdk/20/" lang="en" >}}
- 429: [Scoped Values (Incubator)](https://openjdk.org/jeps/429)
- 432: [Record Patterns (Second Preview)](https://openjdk.org/jeps/432)
- 433: [Pattern Matching for switch (Fourth Preview)](https://openjdk.org/jeps/433)
- 434: [Foreign Function & Memory API (Second Preview)](https://openjdk.org/jeps/434)
- 436: [Virtual Threads (Second Preview)](https://openjdk.org/jeps/436)
- 437: [Structured Concurrency (Second Incubator)](https://openjdk.org/jeps/437)
- 438: [Vector API (Fifth Incubator)](https://openjdk.org/jeps/438)
{{< /fig-quote >}}

[OpenJDK] を APT で管理するメリットはないので，実行バイナリを[リリースページ](https://jdk.java.net/20/)から直接ダウンロードして配置する。
以下は完全手動でのインストール（笑）

```text
$ cd /usr/local/src
$ sudo curl -L "https://download.java.net/java/GA/jdk20/bdc68b4b9cbc4ebcb30745c85038d91d/36/GPL/openjdk-20_linux-x64_bin.tar.gz" -O
$ cd ..
$ sudo unlink java # 以前のバージョンの Java 環境がある場合
$ sudo tar xvf src/openjdk-20_linux-x64_bin.tar.gz
$ sudo ln -s jdk-20 java
$ java -version # すでに PATH が通っている場合
openjdk version "20" 2023-03-21
OpenJDK Runtime Environment (build 20+36-2344)
OpenJDK 64-Bit Server VM (build 20+36-2344, mixed mode, sharing)
```

LTS 版 Java バイナリが欲しいなら [Adoptium](https://adoptium.net/) で取得できる。

私としては [PlantUML] が動けばいいので，試しておく[^puml1]。

[^puml1]: 2023-03-09 に [PlantUML] [1.2023.4](https://github.com/plantuml/plantuml/releases/tag/v1.2023.4) が[リリース](http://plantuml.com/changes)されている。 [PlantUML] の使い方等については拙文「[真面目に PlantUML]({{< ref "/remark/2018/12/plantuml-1.md" >}})」シリーズを参考にどうぞ。

{{< fig-img src="./factory-method-pattern.png" title="Factory Method Pattern" link="./factory-method-pattern.puml" width="1160" >}}

よーし，うむうむ，よーし。

## Oracle Java のサポート期間

“[Oracle Java SE Support Roadmap](https://www.oracle.com/java/technologies/java-se-support-roadmap.html)” より。

| Release  | GA Date | Premier Support | Extended Support |
| -------- | ------- | --------------- | ---------------- |
| 11 (LTS) | 2018-09 | 2023-09         | 2026-09          |
| 17 (LTS) | 2021-09 | 2026-09         | 2029-09          |
| 19       | 2022-09 | 2023-03         | -                |
| 20       | 2023-03 | 2023-09         | -                |
| 21 (LTS) | 2023-09 | 2028-09         | 2031-09          |

2023-03 時点で Premier Support が終了しているものは除いている。
Java 8 については [Adoptium](https://adoptium.net/) などで最新バイナリを取得可能。

[OpenJDK]: http://openjdk.java.net/
[AdoptOpenJDK]: https://adoptopenjdk.net/ "AdoptOpenJDK - Open source, prebuilt OpenJDK binaries"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "B07CKHR8C1" %}} <!-- Spring Data JPAプログラミング入門 -->
{{% review-paapi "4621303252" %}} <!-- Effective Java 第3版 -->
{{% review-paapi "B0893LQ5KY" %}} <!-- Spring Boot 2 入門 -->
f -->
