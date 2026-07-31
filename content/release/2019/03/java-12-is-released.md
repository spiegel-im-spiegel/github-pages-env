+++
title = "Java 12 がリリース"
date = "2019-03-20T22:58:01+09:00"
description = "Switch 文を式として評価できるのは羨ましい。"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "engineering" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

予定通り Java 12 がリリースされた。
[OpenJDK] および同系列の Java のみのショート・サイクルのバージョンアップである。

- [Java 12 / JDK 12: General Availability](https://mail.openjdk.java.net/pipermail/jdk-dev/2019-March/002718.html)
- [Java 12正式版がリリース、大きな変更などはなし。新ガベージコレクタの実験的導入、Switch文の拡張がプレビューなど － Publickey](https://www.publickey1.jp/blog/19/java_12switch.html)

Windows 版だとこんな感じ。

```text
$ java --version
openjdk 12 2019-03-19
OpenJDK Runtime Environment (build 12+33)
OpenJDK 64-Bit Server VM (build 12+33, mixed mode, sharing)
```

「[Kotlin の予備学習]({{< ref "/remark/2018/12/kotlin-book.md" >}})」をしているときに switch 文を式（expression）として評価できると書いてあって「なにそれ羨ましい」と思ったものだが Java 側も追従したようである。
[Go 言語]なんかは意図的に式として評価できる範囲を限定しているので，ちょっと羨ましいんだよなぁ。

アップデートは計画的に。

## ブックマーク

- [Java12が出たので、とりあえずswitch式を試してみた - Qiita](https://qiita.com/dhirabayashi/items/2999b04a369379f41675)

- [Java 環境のリリースとサポートに関する覚え書き]({{< ref "/remark/2018/02/release-cycle-of-java-environment.md" >}})

[OpenJDK]: http://openjdk.java.net/
[Go 言語]: https://golang.org/ "The Go Programming Language"
