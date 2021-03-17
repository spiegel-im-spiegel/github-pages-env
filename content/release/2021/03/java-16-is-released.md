+++
title = "Java 16 がリリースされた"
date =  "2021-03-17T19:32:20+09:00"
description = "OpenJDK および同系列 Java のショート・サイクルのバージョンアップ"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "ubuntu", "plantuml" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

予定通り Java 16 がリリースされた。
[OpenJDK] および同系列 Java におけるショート・サイクルのバージョンアップである。

- [JDK 16](https://openjdk.java.net/projects/jdk/16/)
- [JDK 16 GA Release](https://jdk.java.net/16/)
- [JDK 16 Release Notes](https://jdk.java.net/16/release-notes)

[OpenJDK] を APT で管理するメリットはないので，実行バイナリを[リリースページ](https://jdk.java.net/16/)から直接ダウンロードして配置する。
以下は完全手動でのインストール（笑）

```text
$ cd /usr/local/src
$ sudo curl -L "https://download.java.net/java/GA/jdk16/7863447f0ab643c585b9bdebf67c69db/36/GPL/openjdk-16_linux-x64_bin.tar.gz" -O
$ cd ..
$ sudo unlink java # 以前のバージョンの Java 環境がある場合
$ sudo tar xvf src/openjdk-16_linux-x64_bin.tar.gz
$ sudo ln -s jdk-16 java
$ java -version # すでに PATH が通っている場合
openjdk version "16" 2021-03-16
OpenJDK Runtime Environment (build 16+36-2231)
OpenJDK 64-Bit Server VM (build 16+36-2231, mixed mode, sharing)
```

私としては [PlantUML] が動けばいいので，試しておく[^puml1]。

[^puml1]: 2021-03-07 に [PlantUML] Version 1.2021.2 が[リリース](http://plantuml.com/changes)されている。 [PlantUML] の使い方等については拙文「[真面目に PlantUML]({{< ref "/remark/2018/12/plantuml-1.md" >}})」シリーズを参考にどうぞ。

{{< fig-img src="./factory-method-pattern.png" title="Factory Method Pattern" link="./factory-method-pattern.puml" width="1160" >}}

うむうむ。
ちゃんと動くな。



[OpenJDK]: http://openjdk.java.net/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "4621303252" %}} <!-- Effective Java 第3版 -->
