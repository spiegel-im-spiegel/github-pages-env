+++
title = "Java 10.0.1 がリリース（セキュリティ・アップデート）"
date = "2018-04-18T21:50:09+09:00"
description = "脆弱性の深刻度が高いので早めのアップデートを推奨。"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "engineering" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

Oracle Critical Patch Update が出た。

- [Oracle Critical Patch Update - April 2018](http://www.oracle.com/technetwork/security-advisory/cpuapr2018-3678067.html)

対象となる製品は多岐にわたるが，今回は Java 10.0.1 について紹介する。

Java の脆弱性については IPA および JPCERT/CC からもアラートが上がっている。

- [Oracle Java の脆弱性対策について(CVE-2018-2814等)：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/ciadr/vul/20180418-jre.html)
- [2018年 4月 Oracle 製品のクリティカルパッチアップデートに関する注意喚起](https://www.jpcert.or.jp/at/2018/at180018.html)

Java 10 だけでなく Java 8 以下もアップデート対象になっているのでご注意を（Java 9 は2018年3月で無償サポートが終了している）。
脆弱性の内容によっては CVSSv3 の基本評価値が 7 から 8 程度（深刻度：重要）あるようなので早めのアップデートが推奨される。

Oracle Java 10.0.1 の JDK は以下からダウンロードできる。

- [Java SE Development Kit 10- - Downloads](http://www.oracle.com/technetwork/java/javase/downloads/jdk10-downloads-4416644.html)

[OpenJDK] については以下から入手可能である。

- [JDK Builds from Oracle](http://jdk.java.net/)
    - [JDK 10.0.1 GA Release](http://jdk.java.net/10/)

[OpenJDK] 版を起動すると以下の表示になる（Windows の場合）。

```text
$ java -version
openjdk version "10.0.1" 2018-04-17
OpenJDK Runtime Environment (build 10.0.1+10)
OpenJDK 64-Bit Server VM (build 10.0.1+10, mixed mode)
```

"18.3” みたいなバージョン表記は完全になくなったんだね。
ふむむ。

次回の Critical Patch Update は2018年7月の予定である（緊急リリース等がなければ）。
アップデートは計画的に。

## ブックマーク

- [Java 環境のリリースとサポートに関する覚え書き]({{< ref "/remark/2018/02/release-cycle-of-java-environment.md" >}})

[OpenJDK]: http://openjdk.java.net/
<!-- eof -->
