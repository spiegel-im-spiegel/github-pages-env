+++
title = "Java 25 LTS がリリースされた"
date =  "2025-09-17T09:55:46+09:00"
description = "Java 25 は LTS 版で 2030-09 までのサポートとなる。"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "ubuntu", "plantuml" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

予定通り Java 25 LTS がリリースされた。

- [Oracle Releases Java 25](https://www.oracle.com/news/announcement/oracle-releases-java-25-2025-09-16/)
- [OpenJDK JDK 25 GA Release](https://jdk.java.net/25/)
- [JDK 25 Release Notes](https://jdk.java.net/25/release-notes)

## Java のサポート期間

以下は Oracle が定めるサポート・ロードマップ[^roadmap]。

[^roadmap]: 記事で挙げたロードマップは Oracle Java によるものであり， [OpenJDK] のサポート期間についてはディストリビューションによるが，大抵の場合 LTS 版であれば Oracle Java より長くサポートされることが多い。

{{< fig-quote class="nobox" type="markdown" title="Oracle Java SE Support Roadmap" link="https://www.oracle.com/java/technologies/java-se-support-roadmap.html" lang="en" >}}
| Release  | GA Date | Premier Support | Extended Support |
| -------- | ------- | --------------- | ---------------- |
| 17 (LTS) | 2021-09 | 2026-09         | 2029-09          |
| 21 (LTS) | 2023-09 | 2028-09         | 2031-09          |
| 24       | 2025-03 | 2025-09         |                  |
| 25 (LTS) | 2025-09 | 2030-09         | 2033-09          |
| 26       | 2026-03 | 2026-09         |                  |
| 27       | 2026-09 | 2027-03         |                  |
{{< /fig-quote >}}

これによると Java 25 は 2030-09 までのサポートとなる。
なお 2025-09 時点で Premier Support が終了しているものは除いている（Java 24 は今月（2025-09）でサポートが切れるので注意）。

いまだ需要のある Java 8 等については [Adoptium](https://adoptium.net/) や [Canonical](https://ubuntu.com/toolchains/java) などで最新バイナリを取得可能である。

## OpenJDK のアップデート

ここでは [OpenJDK] の実行バイナリを[リリースページ](https://jdk.java.net/25/)から直接ダウンロードして配置する方法を挙げる。
以下は完全手動でインストールした場合（笑）

```text
$ cd /usr/local/src
$ sudo curl -LO "https://download.java.net/java/GA/jdk25/bd75d5f9689641da8e1daabeccb5528b/36/GPL/openjdk-25_linux-x64_bin.tar.gz"
$ cd ..
$ sudo unlink java # 以前のバージョンの Java 環境がある場合
$ sudo tar xvf src/openjdk-25_linux-x64_bin.tar.gz
$ sudo ln -s jdk-25 java
$ java -version # すでに PATH が通っている場合
openjdk version "25" 2025-09-16
OpenJDK Runtime Environment (build 25+36-3489)
OpenJDK 64-Bit Server VM (build 25+36-3489, mixed mode, sharing)
```

私としては [PlantUML] が動けばいいので，試しておく[^puml1]。

[^puml1]: 2025-09-09 に [PlantUML] [v1.2025.7](https://github.com/plantuml/plantuml/releases/tag/v1.2025.7) が[リリース](http://plantuml.com/changes)されている。 [PlantUML] の使い方等については拙文「[真面目に PlantUML]({{< ref "/remark/2018/12/plantuml-1.md" >}})」シリーズを参考にどうぞ。

{{< fig-img src="./factory-method-pattern.png" title="Factory Method Pattern" link="./factory-method-pattern.puml" width="1171" >}}

よーし，うむうむ，よーし。

## ブックマーク

- [Introducing Canonical builds of OpenJDK](https://canonical.com/blog/introducing-canonical-builds-of-openjdk)
  - [Ubuntu Linux向けの新しいビルドセット「⁠Canonical builds of OpenJDK」が登場|CodeZine（コードジン）](https://codezine.jp/news/detail/21857)
- [「Java 25」正式リリース、2年振りのLTS版。事前キャッシュによる高速起動、JITの即時ネイティブコード生成など新機能 － Publickey](https://www.publickey1.jp/blog/25/java_252ltsjit.html)

[OpenJDK]: http://openjdk.java.net/
[AdoptOpenJDK]: https://adoptopenjdk.net/ "AdoptOpenJDK - Open source, prebuilt OpenJDK binaries"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "B07CKHR8C1" %}} <!-- Spring Data JPAプログラミング入門 -->
{{% review-paapi "4621303252" %}} <!-- Effective Java 第3版 -->
{{% review-paapi "B0893LQ5KY" %}} <!-- Spring Boot 2 入門 -->
