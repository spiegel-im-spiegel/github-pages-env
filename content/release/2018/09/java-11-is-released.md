+++
title = "Java 11 がリリース"
date = "2018-09-27T09:57:29+09:00"
update = "2018-10-21T19:16:32+09:00"
description = "Oracle Java に関しては，予告どおり，有償の LTS (Long Term Support) としてのリリースとなる。"
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
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

予定通り Java 11 がリリースされた。

- [JDK 11 GA Release](https://jdk.java.net/11/)
- [JDK 11 Release Notes](https://jdk.java.net/11/release-notes)

Windows 版だとこんな感じ。

```text
$ java --version
openjdk 11 2018-09-25
OpenJDK Runtime Environment 18.9 (build 11+28)
OpenJDK 64-Bit Server VM 18.9 (build 11+28, mixed mode)
```

Oracle Java に関しては，予告どおり，有償の LTS (Long Term Support) としてのリリースとなる。

- [JDK 11 Is Released! | Oracle The Java Tutorials Blog](https://blogs.oracle.com/thejavatutorials/jdk-11-is-released)
- [JDK 11 Release Notes](https://www.oracle.com/technetwork/java/javase/11-relnotes-5012447.html)
- [Oracle Java SE サポート･ロードマップ](https://www.oracle.com/technetwork/jp/java/eol-135779-ja.html)

有償ではあるが2026年9月までの長期サポート[^j11] となるため， Java と共に滅びる勇気があるのであれば，そちらに乗り換えるのも手であろう。
なお Java 8 に関しては2019年1月までは無償でサポートされる[^j8p]。

[^j8p]: 個人ユーザは2020年12月まで無償でサポートされる。ユーザ端末に JRE を入れさせる馬鹿な実装が今だに多いので仕方ないといったところだろうか。
[^j11]: Extended Support の期限。 Sustaining Support であればバージョンに依らず無期限だがお高い。

[OpenJDK] が無償で LTS を行うのではないかという希望的観測もあったが，今のところそういう話は聞こえないようだ。
Red Hat など，ディストリビュータによっては独自の延長サポートもあるようなので調べてみるのもいいだろう。

他には [AdoptOpenJDK] あたりも良いらしい[^j8]。

[^j8]: [Java 8 については少なくとも2023年9月まではサポートされる](https://adoptopenjdk.net/support.html "Support | AdoptOpenJDK - Open source, prebuilt OpenJDK binaries")ようだ。

個人的な偏見で言わせてもらえば Java は1980年代から1990年代にかけて起こった UNIX 機メーカ間戦争の遺児みたいなもので，ライブラリやフレームワークを含めてバイナリ互換を目指すのであれば後方互換性は絶対条件である筈だったのだ。
Java 9 以降でそれを投げ捨ててしまった今， Java の優位性はないと声を大にして言いたい。
バイナリ互換を保証しなくていいのなら他に FOSS で優秀なプログラミング言語が幾らでもあるのだから。

まぁ Java 8 互換バージョンを legacy Java としてできるだけ維持しつつ徐々に Java 縛りから脱却するのが中長期戦略として妥当だと思うのだけどねぇ。

## 【2018-10-21 追記】 Java 11.0.1 がリリース（セキュリティ・リリース）

2018年10月の定期アップデートに伴い脆弱性の報告が行われている。

- [Oracle Java の脆弱性対策について(CVE-2018-3183等)：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/ciadr/vul/20181017-jre.html)
- [2018年10月 Oracle 製品のクリティカルパッチアップデートに関する注意喚起](http://www.jpcert.or.jp/at/2018/at180042.html)

Oracle Java だけでなく [OpenJDK] も更新されているのでアップデートを行うこと。
以下は Windows 64bit 版の場合

```text
$ java --version
openjdk 11.0.1 2018-10-16
OpenJDK Runtime Environment 18.9 (build 11.0.1+13)
OpenJDK 64-Bit Server VM 18.9 (build 11.0.1+13, mixed mode)
```

アップデートは計画的に。

## ブックマーク

- [Java 11正式版がリリース、本バージョンからOracle JDKのサポートは有償に。OpenJDKで無償の長期サポート提供は現時点で期待薄 － Publickey](https://www.publickey1.jp/blog/18/java_11oracle_jdkopenjdk.html)
- [Microsoft Azure上での実行目的ならJavaの長期サポート（LTS）を無料提供、MacやWindowsでの開発用途もOK。マイクロソフトとAzul Systemsが提携で － Publickey](https://www.publickey1.jp/blog/18/microsoft_azurejavaltsmacwindowsokazul_systems.html)
- [【GlassFish勉強会レポート】各JDKベンダの動向を知ってJava 11に備えよう：レポート｜gihyo.jp … 技術評論社](https://gihyo.jp/news/report/2018/10/0501)
- [Java Is Still Free - Google Document](https://docs.google.com/document/d/1nFGazvrCvHMZJgFstlbzoHjpAVwv5DEdnaBr_5pKuHo/edit#heading=h.p3qt2oh5eczi)
- [Oracle、「Java SE 11.0.1」「Java SE 8 Update 191」を公開 ～12件の脆弱性を修正 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1148283.html)
- [Oracle、Javaやデータベースなど301件の脆弱性を修正　速やかに適用を - ITmedia エンタープライズ](http://www.itmedia.co.jp/enterprise/articles/1810/17/news074.html)

- [Java 環境のリリースとサポートに関する覚え書き]({{< ref "/remark/2018/02/release-cycle-of-java-environment.md" >}})

[OpenJDK]: http://openjdk.java.net/
[AdoptOpenJDK]: https://adoptopenjdk.net/ "AdoptOpenJDK - Open source, prebuilt OpenJDK binaries"
