+++
title = "Java 22 がリリースされた（2024-04 セキュリティ・アップデートあり）"
date =  "2024-03-20T10:00:51+09:00"
description = "Java 22 は 2024-09 までの短期サポート版。 Java 21 は 2028-09 までの LTS 版"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "ubuntu", "plantuml" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

仕事の忙しさにかまけて色々と放っぽり出してたので，少しずつ回復中。

かーなり遅まきながらであるが，昨年（2023年）9月に Java 21 がリリースされている。
また今月（2024-03）に予定通り Java 22 もリリースされた。

- [OpenJDK JDK 21.0.2 GA Release](https://jdk.java.net/21/)
- [JDK 21.0.2 Release Notes](https://jdk.java.net/21/release-notes)
  - [OpenJDK Vulnerability Advisory: 2023/10/17](https://openjdk.org/groups/vulnerability/advisories/2023-10-17)
  - [OpenJDK Vulnerability Advisory: 2024/01/16](https://openjdk.org/groups/vulnerability/advisories/2024-01-16)
- [OpenJDK JDK 22 GA Release](https://jdk.java.net/22/)
- [JDK 22 Release Notes](https://jdk.java.net/22/release-notes)

Java 21 は LTS 版で 2028-09 までのサポートとなる。
既に2回のセキュリティ・アップデートが行われているけど，内容については割愛する。

## 【2024-04-17 追記】 OpenJDK のセキュリティ・アップデート（2024-04）

予定通り [OpenJDK] がリリースされた。

- [OpenJDK Vulnerability Advisory: 2024/04/16](https://openjdk.org/groups/vulnerability/advisories/2024-04-16)

詳細については割愛する。

[OpenJDK] を APT で管理するメリットはないので，実行バイナリを[リリースページ](https://jdk.java.net/22/)から直接ダウンロードして配置する。
以下は完全手動でインストールした場合（笑）

```text
$ cd /usr/local/src
$ sudo curl -L "https://download.java.net/java/GA/jdk22.0.1/c7ec1332f7bb44aeba2eb341ae18aca4/8/GPL/openjdk-22.0.1_linux-x64_bin.tar.gz" -O
$ cd ..
$ sudo unlink java # 以前のバージョンの Java 環境がある場合
$ sudo tar xvf src/openjdk-22.0.1_linux-x64_bin.tar.gz
$ sudo ln -s jdk-22.0.1 java
$ java -version # すでに PATH が通っている場合
openjdk version "22.0.1" 2024-04-16
OpenJDK Runtime Environment (build 22.0.1+8-16)
OpenJDK 64-Bit Server VM (build 22.0.1+8-16, mixed mode, sharing)
```

私としては [PlantUML] が動けばいいので，試しておく[^puml1]。

[^puml1]: 2024-04-06 に [PlantUML] [v1.2024.4](https://github.com/plantuml/plantuml/releases/tag/v1.2024.4) が[リリース](http://plantuml.com/changes)されている。 [PlantUML] の使い方等については拙文「[真面目に PlantUML]({{< ref "/remark/2018/12/plantuml-1.md" >}})」シリーズを参考にどうぞ。

{{< fig-img src="./factory-method-pattern.png" title="Factory Method Pattern" link="./factory-method-pattern.puml" width="1160" >}}

よーし，うむうむ，よーし。

## Oracle Java のサポート期間

“[Oracle Java SE Support Roadmap](https://www.oracle.com/java/technologies/java-se-support-roadmap.html)” より。

| Release  | GA Date | Premier Support | Extended Support |
| -------- | ------- | --------------- | ---------------- |
| 17 (LTS) | 2021-09 | 2026-09         | 2029-09          |
| 21 (LTS) | 2023-09 | 2028-09         | 2031-09          |
| 22       | 2024-03 | 2024-09         |                  |
| 23       | 2024-09 | 2025-03         |                  |
| 24       | 2025-03 | 2025-09         |                  |
| 25 (LTS) | 2025-09 | 2030-09         | 2033-09          |

2024-03 時点で Premier Support が終了しているものは除いている。
いまだ需要のある Java 8 については [Adoptium](https://adoptium.net/) などで最新バイナリを取得可能。

## ブックマーク

- [Java 22リリース――Project AmberやProject PanamaのスタンダードJEPを含め12のJEPsがアップデート。そして、2025年のJavaOneがベイエリアで開催 | gihyo.jp](https://gihyo.jp/article/2024/03/java22)
- [Oracle Critical Patch Update Advisory - April 2024](https://www.oracle.com/security-alerts/cpuapr2024.html)
- [Oracle Java の脆弱性対策について(CVE-2023-41993等) | 情報セキュリティ | IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/security-alert/2024/0417-jre.html)

[OpenJDK]: http://openjdk.java.net/
[AdoptOpenJDK]: https://adoptopenjdk.net/ "AdoptOpenJDK - Open source, prebuilt OpenJDK binaries"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "B07CKHR8C1" %}} <!-- Spring Data JPAプログラミング入門 -->
{{% review-paapi "4621303252" %}} <!-- Effective Java 第3版 -->
{{% review-paapi "B0893LQ5KY" %}} <!-- Spring Boot 2 入門 -->
