+++
title = "Java 環境のリリースとサポートに関する覚え書き"
date = "2018-02-25T19:31:42+09:00"
update = "2018-05-13T13:24:15+09:00"
description = "特にこの記事では嘘や間違いがあればフィードバック等いただけると有り難い。随時加筆・修正する予定。"
image = "/images/attention/kitten.jpg"
tags        = [ "java", "engineering", "license", "tools" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先日 Java のリリースサイクルとサポートについて訊かれたのだが，1年半近く Java から離れていたこともあって，まったくノーチェックだった。
そこで，改めて Java 環境について調べてみることにした。

この記事はその時の覚え書きとして残しておく。
なお，特にこの記事では嘘や間違いがあればフィードバック等いただけると有り難い。
Oracle の料金体系とかイマイチよく分かっていないのだ。

## Java 9 は後方互換性がない？

2017年9月にリリースされた Java 9 の最大の特徴は [Project Jigsaw] の成果が取り込まれモジュール化に対応したことだが，他にも

{{< fig-quote title="Java 9が正式リリース、Javaをモジュール化するProject Jigsawがついに実現。今後のJavaは6カ月ごとタイムベースのアップデートへ" link="http://www.publickey1.jp/blog/17/java_9_release_project_jigsaw.html" >}}
<q>一方でJava 9は日付や通貨のデフォルトフォーマットが変更され、いくつかの構文や演算子の変更や廃止が行われるなど、Java 8以前との互換性は保証されていません。</q>
{{< /fig-quote >}}

ということで，これまで保証されていた後方互換性（backward compatibility）はなくなってしまった。
したがって Java 9 以降へのアップグレードについては十分な調査と場合によっては改修作業が必要になる。

- [Java Platform, Standard Edition Oracle JDK 9 Migration Guide, Release 9](https://docs.oracle.com/javase/9/migrate/toc.htm)
    - [The Hitchhiker's Migration Guide to JDK 9 新機能の影で消えた機能](https://kasecato.github.io/migrating2Jdk9/)

### Java 8 の無償サポートは2019年1月まで

一方， Java 8 の無償サポート期間が2019年1月までとなっていて期限が迫っている[^java8_1]。
有償サポートは Java SE Advanced の Named User Plus[^nup1] で2,460円/ユーザ， Processor License[^pl1] で132K円/コアとなる。

[^java8_1]: ただし「[個人で Java SE を利用する（non-corporate desktop use）場合、2020年12月末まで引き続きアップデートを受け取ることができます](http://www.oracle.com/technetwork/jp/java/eol-135779-ja.html)」。
[^nup1]: Named User Plus は利用するユーザ数に応じて料金がかかる。サポート料金とは別に12K円/ユーザ必要。
[^pl1]: Processor License は Web サービスなど不特定多数のユーザが利用する場合のライセンス。なお MPU 単位ではなくコア単位になるため，コア4つの MPU なら料金は4倍になる。サポート料金とは別に600K円/コア必要。単一コア単一プロセッサのみの構成ならユーザ数が50人を超えれば Processor License のほうが割安になる。

- {{< pdf-file title="Oracle Technology Global Price List" link="http://www.oracle.com/jp/corporate/pricing/e-pl101005-101005a-n-176288-ja.pdf" >}}
- [「Java SE 8」のパブリックアップデートが延長、新規提供は2019年1月まで継続される - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1103999.html)

## Java のリリースサイクル

Java 9 以降は6ヶ月毎に機能の追加・変更を含むメジャー・バージョンアップが行われる。
また1月，4月，7月，10月にバグや脆弱性の修正を含むマイナー・バージョンアップが行われる。

|   月 | バージョンアップ・イベント |
| ----:| -------------------------- |
|  1月 | マイナー・バージョンアップ |
|  2月 |                            |
|  3月 | メジャー・バージョンアップ |
|  4月 | マイナー・バージョンアップ |
|  5月 |                            |
|  6月 |                            |
|  7月 | マイナー・バージョンアップ |
|  8月 |                            |
|  9月 | メジャー・バージョンアップ |
| 10月 | マイナー・バージョンアップ |
| 11月 |                            |
| 12月 |                            |

マイナー・バージョンアップは基本的に現行メジャー・バージョンに対してのみ行われ，過去のバージョンについてはサポート外となる。
ただし LTS (Long Term Support) 対象のバージョンについては新バージョン・リリース後もサポートが行われる。
（LTS については次節で解説）

- [［速報］Java 9が正式リリース、Javaをモジュール化するProject Jigsawがついに実現。今後のJavaは6カ月ごとタイムベースのアップデートへ － Publickey](http://www.publickey1.jp/blog/17/java_9_release_project_jigsaw.html)

### Oracle Java と [OpenJDK]

Java および JDK (Java Development Kit) の実装としては Oracle Java と [OpenJDK] が双璧と言える。
また Oracle は JDK を GPL ライセンスでオープンソース化し [OpenJDK] との統合を行っている。

{{< fig-quote title="来月にはJava 10が登場し、9月にはJava 11が登場予定。新しいリリースモデルを採用した今後のJava、入手方法やサポート期間はこう変わる" link="http://www.publickey1.jp/blog/17/java_9_release_project_jigsaw.html" >}}
<q>商用ライセンスのOracle JDKには、無償版にはないJavaのプロファイラのFlight RecorderやMission Controlといったツールが提供されています。<br>
オラクルは現在、こうしたツールをオープンソースのJava実装であるOpenJDKにも含めるように作業を進めています。<br>
予定ではJava 11がリリースされる2018年9月にこの作業が完了し、商用ディストリビューションのOracle JDKとオープンソース実装のOpenJDKは機能、品質両面で同じものになる予定です。</q>
{{< /fig-quote >}}

今のところ Oracle Java と [OpenJDK] は協力関係にあり Java 9 以降の各バージョンのリリースを同時期に行うことになっている[^ocl1]。
リリース時期とサポート期間は以下の通り。

[^ocl1]: ただし開発の主導権は Oracle 側が握ってる印象。 [OpenJDK] では，メジャー・バージョンのアナウンスはあれど，不具合や脆弱性に関する周知プロセスが皆無なので，結局のところ Oracle の動向を見ないと何も分からない。これで Oracle が LTS に入ったらどうするつもりなのか...

|         | Oracle Java                | [OpenJDK]             |
| ------- | -------------------------- | --------------------- |
| Java 9  | 2017年9月 - 2018年3月      | 2017年9月 - 2018年3月 |
| Java 10 | 2018年3月 - 2018年9月      | 2018年3月 - 2018年9月 |
| Java 11 | 2018年9月 - 2026年9月[^s1] | 2018年9月 - 2019年3月 |
| Java 12 | -                          | 2019年3月 - 2019年9月 |
| Java 13 | -                          | 2019年9月 - 2020年3月 |
| Java 14 | -                          | 2020年3月 - 2020年9月 |
| Java 15 | -                          | 2020年9月 - 2021年3月 |
| Java 16 | -                          | 2021年3月 - 2021年9月 |
| Java 17 | 2021年9月 - 2029年9月[^s1] | 2021年9月 - 2022年3月 |

[^s1]: 5年間の Premier Support および3年間の Extended Support を併せた期間。更に延長する場合は Sustaining Support の契約が必要。

Oracle Java は Java 11 以降は有償の LTS のみとなる。
LTS 対象のバージョンは3年毎にリリースされるが，サポートは（メジャー・バージョンアップしなくてもお金さえ払えば）継続可能である。

[OpenJDK] はオープンソース（[GPLv2 + Classpath Exception](http://openjdk.java.net/legal/gplv2+ce.html "OpenJDK: GPLv2 + Classpath Exception")）なのでそれ自体にサポート料金は発生しないが（ディストリビュータ等が有償でサポートを請け負う場合は別），サポート期間が（次のメジャー・バージョンがリリースされるまでの）半年のみとなる。

{{< fig-quote title="来月にはJava 10が登場し、9月にはJava 11が登場予定。新しいリリースモデルを採用した今後のJava、入手方法やサポート期間はこう変わる" link="http://www.publickey1.jp/blog/17/java_9_release_project_jigsaw.html" >}}
<q>ただしOpenJDKがオープンソースプロジェクトの判断としてLTSを設定し、LTSに対するメンテナンスバージョンを継続してリリースする動きがあるとの指摘をいただきました。<br>
[...]<br>
これが実行されれば、Java 11、Java 17などLTSが設定されたバージョンのOpenJDKに対しては、数年間（スライドでは次のLTSのメジャーバージョンが登場するまでの3年間）、無償のOpenJDKでもメンテナンスパッチが提供されることになります。</q>
{{< /fig-quote >}}

- [Oracle Java SE サポート･ロードマップ](http://www.oracle.com/technetwork/jp/java/eol-135779-ja.html)
- [OracleがJDKの全ての機能をオープンソース化し、Java EEの欠点に取り組む計画を発表した](https://www.infoq.com/jp/news/2017/11/javaone-opening)
- [来月にはJava 10が登場し、9月にはJava 11が登場予定。新しいリリースモデルを採用した今後のJava、入手方法やサポート期間はこう変わる（OpenJDKに関する追記あり） － Publickey](http://www.publickey1.jp/blog/18/java_109java_11java.html)

## Java 関連のサービスやアプリケーション

Java 関連のサービスやアプリケーションについて以下にメモ書きしておく。
随時加筆・修正する予定。

### IDE (Integrated Development Environment)

もっとも有名な IDE である [Eclipse] については [Oxygen].1a で正式に Java 9 をサポートしているようだ。

- [Configure Eclipse for Java 9 - Eclipsepedia](https://wiki.eclipse.org/Configure_Eclipse_for_Java_9)
- [Eclipse Oxygen.1a is out — with support for Java 9 and JUnit 5 - JAXenter](https://jaxenter.com/eclipse-oxygen-1a-java-9-junit-5-138113.html)
- [Eclipse 4.7 Oxygen 新機能 30+ / Java 9 を試そう！ - Qiita](https://qiita.com/cypher256/items/e308d920dfaf15892baa)

[Eclipse] のプラグインが Java 9 に追従しているかどうかは調査する必要がある。
現在 [Eclipse] のリリースサイクルは1年（毎年6月頃）であり，今後の Java のリリースサイクルに追従できない可能性もあるが，どうなっているのだろう。

もうひとつの有名な IDE として [IntelliJ IDEA] があるが，こちらも Java 9 への対応はできているようだ。

- [Java 9 and IntelliJ IDEA | IntelliJ IDEA Blog](https://blog.jetbrains.com/idea/2017/09/java-9-and-intellij-idea/)

### Java 9 Ready な FOSS

[Quality Outreach](https://wiki.openjdk.java.net/display/quality/Quality+Outreach "Quality Outreach - Quality - OpenJDK Wiki") より。
（自己申告らしいので他にもあるかもしれない）

- Akka, Lightbend
- Apache Derby
- Apache Log4j
- Apache Lucene/SOLR
- Apache Maven
- Apache Kafka
- Apache PDFBox
- Apache POI
- BootstrapFX
- bt
- Byte Buddy
- CruiseControl
- Eclipse Collections
- EqualsVerifier
- GraphHopper
- HeapStats
- Hibernate
- Ikonli
- Jackson
- JaCoCo
- JITWatch
- jOOQ
- JOSM
- JSilhouette
- JUnit 5
- Lillith
- LWJGL
- Mockito
- Rapidoid
- RedHat Wildfly
- RxJava
- Spotbugs
- Spring Framework
- Woodstox

[Apache Tomcat] は Java 9 Ready じゃない？

## ブックマーク

- [ヌーラボのアカウント基盤を Java 9 にマイグレーションして起きた問題と解決法 | ヌーラボ](https://nulab-inc.com/ja/blog/nulab/java9-migration/)
- [Java9勘所集 - Qiita](https://qiita.com/shiroma_yuki/items/8725a73493e3fe75155d)
- [Java 9のモジュール機能「Project Jigsaw」の基本を紹介 (1/2)：CodeZine（コードジン）](https://codezine.jp/article/detail/10524)
- [第2回　JDKの新しいリリースモデルに要注目 OpenJDKとOracle JDKの違いにも注意が必要［JavaOne 2017］：Java 9のその先へ～JavaOne Conference 2017レポート｜gihyo.jp … 技術評論社](http://gihyo.jp/news/report/01/JavaOne2017/0002)
- [Javaエンジニアが Java 11 リリースに向けて備えておくべきこと - Qiita](https://qiita.com/mao172/items/dbf41e7a246c3f87dcbe)

[OpenJDK]: http://openjdk.java.net/
[Project Jigsaw]: http://openjdk.java.net/projects/jigsaw/ "Project Jigsaw"
[Eclipse]: https://www.eclipse.org/ "Eclipse - The Eclipse Foundation open source community website."
[Oxygen]: http://www.eclipse.org/oxygen/ "Eclipse Oxygen"
[IntelliJ IDEA]: https://www.jetbrains.com/idea/ "IntelliJ IDEA: The Java IDE for Professional Developers by JetBrains"
[Apache Tomcat]: https://tomcat.apache.org/ "Apache Tomcat® - Welcome!"
