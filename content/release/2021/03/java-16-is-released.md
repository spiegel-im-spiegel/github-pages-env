+++
title = "Java 16 がリリースされた"
date =  "2021-03-17T19:32:20+09:00"
description = "OpenJDK および同系列 Java のショート・サイクルのバージョンアップ"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "ubuntu", "plantuml", "windows" ]
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

## Scoop で [OpenJDK] 16 を入れる

Windows 環境では Scoop を使って [OpenJDK] を入れているのだが，なかなかアップデートされない。
そこでバケットの中を見てみたのだが

```text
$ scoop search openjdk
'java' bucket:
    adoptopenjdk-hotspot-jre (15.0.2-7)
    adoptopenjdk-hotspot (15.0.2-7)
    adoptopenjdk-lts-hotspot-jre (11.0.10-9)
    adoptopenjdk-lts-hotspot (11.0.10-9)
    adoptopenjdk-lts-openj9-jre (11.0.10-9-0.24.0)
    adoptopenjdk-lts-openj9-xl-jre (11.0.10-9-0.24.0)
    adoptopenjdk-lts-openj9-xl (11.0.10-9-0.24.0)
    adoptopenjdk-lts-openj9 (11.0.10-9-0.24.0)
    adoptopenjdk-lts-upstream-jre (11.0.10-9)
    adoptopenjdk-lts-upstream (11.0.10-9)
    adoptopenjdk-openj9-jre (16-36-0.25.0)
    adoptopenjdk-openj9-xl-jre (15.0.2-7-0.24.0)
    adoptopenjdk-openj9-xl (15.0.2-7-0.24.0)
    adoptopenjdk-openj9 (16-36-0.25.0)
    openjdk-ea (16-36)
    openjdk (15.0.2-7)
    openjdk10 (10.0.2-13)
    openjdk11 (11.0.2-9)
    openjdk12 (12.0.2-10)
    openjdk13 (13.0.2-8)
    openjdk14 (14.0.2-12)
    openjdk15 (15.0.2-7)
    openjdk16 (16-36)
    openjdk7-unofficial (7u80-b32)
    openjdk8-redhat-jre (8u282-b08)
    openjdk8-redhat (8u282-b08)
    openjdk9 (9.0.4-12)
```

んー。
`openjdk` だと 16 に上がらないのか。
どうやら `openjdk16` のようにバージョンを指定したほうがいいようだ。

というわけで

```text
$ scoop install openjdk16
Installing 'openjdk16' (16-36) [64bit]
openjdk-16_windows-x64_bin.zip (175.1 MB) [===================================================================] 100%
Checking hash of openjdk-16_windows-x64_bin.zip ... ok.
Extracting openjdk-16_windows-x64_bin.zip ... done.
Linking ~\scoop\apps\openjdk16\current => ~\scoop\apps\openjdk16\16-36

$ scoop uninstall openjdk
Uninstalling 'openjdk' (15.0.2-7).
Unlinking ~\scoop\apps\openjdk\current
Removing ~\scoop\apps\openjdk\current\bin from your path.
Removing older version (15.0.1-9).
'openjdk' was uninstalled.
```

という感じに入れ換えた。
Windows Terminal を起動し直して

```text
$ java -version
openjdk version "16" 2021-03-16
OpenJDK Runtime Environment (build 16+36-2231)
OpenJDK 64-Bit Server VM (build 16+36-2231, mixed mode, sharing)
```

よしよし，ちゃんと 16 になったな。
ちなみに各パッケージの旧バージョンを削除する場合は

```text
$ scoop cleanup openjdk16
```

などとすれば，最新バージョン以外は削除される。

## ブックマーク

- [JDK 16 Security Enhancements](https://seanjmullan.org/blog/2021/03/18/jdk16)
- [Java 16正式リリース。JavaアプリをWin/Mac/Linuxのインストール形式にするパッケージャ登場、OpenJDKソースコードがGitHubへ移行 － Publickey](https://www.publickey1.jp/blog/21/java_16javawinmaclinuxopenjdkgithub.html)
- [Oracle、「Java 16」を発表 ～パターンマッチングinstanceofとrecord型が正式機能に - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1312633.html)

[OpenJDK]: http://openjdk.java.net/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "4621303252" %}} <!-- Effective Java 第3版 -->
