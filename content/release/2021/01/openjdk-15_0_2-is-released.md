+++
title = "OpenJDK 15.0.2 がリリースされた"
date =  "2021-01-21T20:32:19+09:00"
description = "今回は OpenJDK に関しては脆弱性の修正がない。"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "ubuntu", "plantuml" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

定例の Java マイナー・バージョンアップが行われた。
なんとビックリなことに今回は [OpenJDK] に関しては脆弱性の修正がない。

- [OpenJDK Vulnerability Advisory: 2021/01/19](https://openjdk.java.net/groups/vulnerability/advisories/2021-01-19)

ただし Oracle Java については Java 8 以下で脆弱性の修正があるらしい。

- [Oracle Java の脆弱性対策について(CVE-2020-14803)：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/ciadr/vul/20210120-jre.html)

これって [OpenJDK] には影響ないのか？ ホンマに？ ...まぁ，いいや。

[OpenJDK] を APT で管理するメリットはないので，実行バイナリを[リリースページ](https://jdk.java.net/15/)から直接ダウンロードして配置する。

```text
$ cd /usr/local/src
$ sudo curl "https://download.java.net/java/GA/jdk15.0.2/0d1cfde4252546c6931946de8db48ee2/7/GPL/openjdk-15.0.2_linux-x64_bin.tar.gz" -O
$ cd ..
$ sudo unlink java # 以前のバージョンの Java 環境がある場合
$ sudo tar xvf src/openjdk-15.0.2_linux-x64_bin.tar.gz
$ sudo ln -s jdk-15.0.2 java
$ java -version # すでに PATH が通っている場合
openjdk version "15.0.2" 2021-01-19
OpenJDK Runtime Environment (build 15.0.2+7-27)
OpenJDK 64-Bit Server VM (build 15.0.2+7-27, mixed mode, sharing)
```

私としては [PlantUML] が動けばいいので，試しておく[^puml1]。

[^puml1]: 2021-01-10 に [PlantUML] V1.2021.0 が[リリース](http://plantuml.com/ja/changes)されている。 [PlantUML] の使い方等については拙文「[真面目に PlantUML]({{< ref "/remark/2018/12/plantuml-1.md" >}})」シリーズを参考にどうぞ。

{{< fig-img src="./factory-method-pattern.png" title="Factory Method Pattern" link="./factory-method-pattern.puml" width="1178" >}}

うむうむ。
ちゃんと動くな。

## ブックマーク

- [2021年1月Oracle製品のクリティカルパッチアップデートに関する注意喚起](https://www.jpcert.or.jp/at/2021/at210003.html)

[OpenJDK]: http://openjdk.java.net/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "B00I8AT1EU" %}} <!-- Java言語で学ぶリファクタリング入門 -->
{{% review-paapi "B00I8ATHGW" %}} <!-- 増補改訂版 Java言語で学ぶデザインパターン入門 -->
{{% review-paapi "B00I8AT1BS" %}} <!-- Java言語で学ぶデザインパターン入門 マルチスレッド編 -->
