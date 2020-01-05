+++
title = "Java 13.0.1 がリリース【セキュリティ・アップデート】"
date =  "2019-10-16T09:39:08+09:00"
description = "予定通り Java 13 のマイナー・バージョンアップが行われた。脆弱性の修正がてんこ盛りである。"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "ubuntu", "plantuml", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

予定通り Java 13 のマイナー・バージョンアップが行われた。
緊急性が高いものはないようだが，脆弱性の修正がてんこ盛りである。

- [OpenJDK Vulnerability Advisory: 2019/10/15](https://openjdk.java.net/groups/vulnerability/advisories/2019-10-15)

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
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2019-2949">CVE-2019-2949</a></td>
<td>security-libs/javax.net.ssl</td>
<td>6.8</td>
<td></td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2019-2989">CVE-2019-2989</a></td>
<td>core-libs/java.net</td>
<td>6.8</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2019-2958">CVE-2019-2958</a></td>
<td>core-libs/java.lang</td>
<td>5.9</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2019-2975">CVE-2019-2975</a></td>
<td>core-libs/javax.script</td>
<td>4.8</td>
<td></td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2019-2977">CVE-2019-2977</a></td>
<td>hotspot/compiler</td>
<td>4.8</td>
<td></td>
<td></td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2019-2999">CVE-2019-2999</a></td>
<td>tools/javadoc(tool)</td>
<td>4.7</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2019-2981">CVE-2019-2981</a></td>
<td>xml/jaxp</td>
<td>3.7</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2019-2973">CVE-2019-2973</a></td>
<td>xml/jaxp</td>
<td>3.7</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2019-2983">CVE-2019-2983</a></td>
<td>client-libs/2d</td>
<td>3.7</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2019-2988">CVE-2019-2988</a></td>
<td>client-libs/2d</td>
<td>3.7</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2019-2978">CVE-2019-2978</a></td>
<td>core-libs/java.net</td>
<td>3.7</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2019-2964">CVE-2019-2964</a></td>
<td>core-libs/java.util.regex</td>
<td>3.7</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2019-2992">CVE-2019-2992</a></td>
<td>client-libs/2d</td>
<td>3.7</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2019-2962">CVE-2019-2962</a></td>
<td>client-libs/2d</td>
<td>3.7</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2019-2987">CVE-2019-2987</a></td>
<td>client-libs/2d</td>
<td>3.7</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2019-2894">CVE-2019-2894</a></td>
<td>security-libs/javax.net.ssl</td>
<td>3.7</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2019-2933">CVE-2019-2933</a></td>
<td>core-libs</td>
<td>3.1</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
<tr>
<td><a href="https://nvd.nist.gov/vuln/detail/CVE-2019-2945">CVE-2019-2945</a></td>
<td>core-libs/java.net</td>
<td>3.1</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>
</table>
{{< /fig-quote >}}


[OpenJDK] を APT で管理するメリットはないので，実行バイナリを[リリースページ](https://jdk.java.net/13/)から直接ダウンロードして配置する。

```text
$ cd /usr/local/src
$ sudo curl "https://download.java.net/java/GA/jdk13.0.1/cec27d702aa74d5a8630c65ae61e4305/9/GPL/openjdk-13.0.1_linux-x64_bin.tar.gz" -O
$ cd ..
$ sudo unlink java # 以前のバージョンの Java 環境がある場合
$ sudo tar xvf src/openjdk-13.0.1_linux-x64_bin.tar.gz
$ sudo ln -s jdk-13.0.1 java
$ java -version # すでに PATH が通っている場合
openjdk version "13.0.1" 2019-10-15
OpenJDK Runtime Environment (build 13.0.1+9)
OpenJDK 64-Bit Server VM (build 13.0.1+9, mixed mode, sharing)
```

私としては [PlantUML] が動けばいいので，試しておく[^puml1]。

[^puml1]: 2019-09-22 に [PlantUML] Version 1.2019.11 が[リリース](http://plantuml.com/ja/changes)されている。

{{< fig-quote title="Go 言語で Factory Method Pattern を構成できるか" link="/golang/factory-method-pattern/" >}}
{{< fig-img src="/release/2019/09/java-13-is-released/factory-method-pattern.png" link="/release/2019/09/java-13-is-released/factory-method-pattern.puml" width="1165" >}}
{{< /fig-quote >}}

うむうむ。
ちゃんと動くな。

## ブックマーク

- [Oracle Critical Patch Update - October 2019](https://www.oracle.com/technetwork/security-advisory/cpuoct2019-5072832.html)

- [Java 環境のリリースとサポートに関する覚え書き]({{< ref "/remark/2018/02/release-cycle-of-java-environment.md" >}})
- [結局 OpenJDK をインストールし直すことにした]({{< ref "/remark/2019/07/reinstalling-openjdk.md" >}})

[OpenJDK]: http://openjdk.java.net/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "B00I8AT1EU" %}} <!-- Java言語で学ぶリファクタリング入門 -->
{{% review-paapi "B00I8ATHGW" %}} <!-- 増補改訂版 Java言語で学ぶデザインパターン入門 -->
{{% review-paapi "B00I8AT1BS" %}} <!-- Java言語で学ぶデザインパターン入門 マルチスレッド編 -->
