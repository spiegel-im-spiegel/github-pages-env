+++
title = "OpenJDK 13.0.2 のリリース【セキュリティ・アップデート】"
date =  "2020-01-15T09:11:22+09:00"
description = "今回は深刻度が高いセキュリティ・アップデートを含んでいるようだ。"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "ubuntu", "plantuml", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

予定通り Java 13 のマイナー・バージョンアップが行われた。
今回は深刻度が高いセキュリティ・アップデートを含んでいるようだ。

- [OpenJDK Vulnerability Advisory: 2020/01/14](https://openjdk.java.net/groups/vulnerability/advisories/2020-01-14)

以下に一覧を挙げておく。
なお Java 9, 10 および 12 は既にサポート外なので注意すること。

{{< fig-quote title="OpenJDK Vulnerability Advisory: 2019/10/15" link="https://openjdk.java.net/groups/vulnerability/advisories/2019-10-15" lang="en" >}}
<table class="risk-matrix center" summary="Risk matrix">
<tr>
<th rowspan="2">CVE ID</th>
<th rowspan="2">Component</th>
<th rowspan="2"><a href="https://www.first.org/cvss/">CVSSv3</a><br>Score</th>
<th colspan="4">Affects ...</th>
</tr>
<tr>
<th>7</th>
<th>8</th>
<th>11</th>
<th>13</th>
</tr>

<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-2604">CVE-2020-2604</a></td>
<td>core-libs/java.io:serialization</td>
<td>8.1</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-2601">CVE-2020-2601</a></td>
<td>security-libs/java.security</td>
<td>6.8</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-2655">CVE-2020-2655</a></td>
<td>security-libs/javax.net.ssl</td>
<td>4.8</td>
<td></td>
<td></td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-2593">CVE-2020-2593</a></td>
<td>core-libs/java.net</td>
<td>4.8</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-2654">CVE-2020-2654</a></td>
<td>security-libs/java.security</td>
<td>3.7</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-2590">CVE-2020-2590</a></td>
<td>security-libs/org.ietf.jgss</td>
<td>3.7</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-2659">CVE-2020-2659</a></td>
<td>core-libs/java.nio</td>
<td>3.7</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td></td>
<td></td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-2583">CVE-2020-2583</a></td>
<td>client-libs/java.beans</td>
<td>3.7</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>

</table>
{{< /fig-quote >}}

相変わらず溜め込んでるなぁ（笑）

[OpenJDK] を APT で管理するメリットはないので，実行バイナリを[リリースページ](https://jdk.java.net/13/)から直接ダウンロードして配置する。

```text
$ cd /usr/local/src
$ sudo curl "https://download.java.net/java/GA/jdk13.0.2/d4173c853231432d94f001e99d882ca7/8/GPL/openjdk-13.0.2_linux-x64_bin.tar.gz" -O
$ cd ..
$ sudo unlink java # 以前のバージョンの Java 環境がある場合
$ sudo tar xvf src/openjdk-13.0.2_linux-x64_bin.tar.gz
$ sudo ln -s jdk-13.0.2 java
$ java -version # すでに PATH が通っている場合
openjdk version "13.0.2" 2020-01-14
OpenJDK Runtime Environment (build 13.0.2+8)
OpenJDK 64-Bit Server VM (build 13.0.2+8, mixed mode, sharing)
```

私としては [PlantUML] が動けばいいので，試しておく[^puml1]。

[^puml1]: 2020-01-12 に [PlantUML] VV1.2020.0 が[リリース](http://plantuml.com/ja/changes)されている。

{{< fig-quote title="Go 言語で Factory Method Pattern を構成できるか" link="/golang/factory-method-pattern/" >}}
{{< fig-img src="/release/2019/09/java-13-is-released/factory-method-pattern.png" link="/release/2019/09/java-13-is-released/factory-method-pattern.puml" width="1165" >}}
{{< /fig-quote >}}

うむうむ。
ちゃんと動くな。

## ブックマーク

- [2020年1月 Oracle 製品のクリティカルパッチアップデートに関する注意喚起](https://www.jpcert.or.jp/at/2020/at200002.html) : Oracle Java を含む Oracle 製品のアップデート情報

- [Java 環境のリリースとサポートに関する覚え書き]({{< ref "/remark/2018/02/release-cycle-of-java-environment.md" >}})
- [結局 OpenJDK をインストールし直すことにした]({{< ref "/remark/2019/07/reinstalling-openjdk.md" >}})

[OpenJDK]: http://openjdk.java.net/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "B00I8AT1EU" %}} <!-- Java言語で学ぶリファクタリング入門 -->
{{% review-paapi "B00I8ATHGW" %}} <!-- 増補改訂版 Java言語で学ぶデザインパターン入門 -->
{{% review-paapi "B00I8AT1BS" %}} <!-- Java言語で学ぶデザインパターン入門 マルチスレッド編 -->
