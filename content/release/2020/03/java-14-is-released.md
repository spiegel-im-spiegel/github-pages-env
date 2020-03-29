+++
title = "Java 14 がリリースされた"
date =  "2020-03-22T16:29:03+09:00"
description = "まぁ，私は PlantUML が問題なく動けばいいので（笑）"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "ubuntu", "plantuml" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

ちょっと前に予定通り Java 14 がリリースされた。
[OpenJDK] および同系列の Java のみのショート・サイクルのバージョンアップである。

- [JDK 14](https://openjdk.java.net/projects/jdk/14/)
- [JDK 14 GA Release](https://jdk.java.net/14/)
- [JDK 14 Release Notes](http://jdk.java.net/14/release-notes)

まぁ，私は [PlantUML] が問題なく動けばいいので[^puml1]（笑）

[^puml1]: 2020-03-19 に [PlantUML] V1.2020.5 が[リリース](http://plantuml.com/changes)されている。

[OpenJDK] を APT で管理するメリットはないので，実行バイナリを[リリースページ](https://jdk.java.net/13/)から直接ダウンロードして配置する。

```text
$ cd /usr/local/src
$ sudo curl "https://download.java.net/java/GA/jdk14/076bab302c7b4508975440c56f6cc26a/36/GPL/openjdk-14_linux-x64_bin.tar.gz" -O
$ cd ..
$ sudo unlink java # 以前のバージョンの Java 環境がある場合
$ sudo tar xvf src/openjdk-14_linux-x64_bin.tar.gz
$ sudo ln -s jdk-14 java
$ java -version # すでに PATH が通っている場合
openjdk version "14" 2020-03-17
OpenJDK Runtime Environment (build 14+36-1461)
OpenJDK 64-Bit Server VM (build 14+36-1461, mixed mode, sharing)
```

[PlantUML] のほうも試してみるかな。

{{< fig-img src="./factory-method-pattern.png" link="./factory-method-pattern.puml" title="Factory Method Pattern" width="1165" >}}

よーし，うむうむ，よーし。

## ブックマーク

- [Java 14正式版が登場。テキストブロック、インストーラー作成ツールなど新機能。Solaris/SPARC版はついに引退 － Publickey](https://www.publickey1.jp/blog/20/java_14solarissparc.html)
- [Java 14新機能まとめ - Qiita](https://qiita.com/nowokay/items/ec85d97a7cecaaac8123)
- [真面目に PlantUML]({{< ref "/remark/2018/12/plantuml-1.md" >}})

[OpenJDK]: http://openjdk.java.net/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "B00I8AT1EU" %}} <!-- Java言語で学ぶリファクタリング入門 -->
{{% review-paapi "B00I8ATHGW" %}} <!-- 増補改訂版 Java言語で学ぶデザインパターン入門 -->
{{% review-paapi "B00I8AT1BS" %}} <!-- Java言語で学ぶデザインパターン入門 マルチスレッド編 -->
