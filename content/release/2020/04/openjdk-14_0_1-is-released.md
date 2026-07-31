+++
title = "OpenJDK 14.0.1 のリリース【セキュリティ・アップデート】"
date =  "2020-04-16T12:21:08+09:00"
description = "深刻度が高いセキュリティ・アップデートも含まれているので必ず対応すること。"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "ubuntu", "plantuml", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

予定通り Java 14 および LTS である Java 11 のマイナー・バージョンアップが行われた。
深刻度が高い脆弱性も含まれているので必ず対応すること。

- [OpenJDK Vulnerability Advisory: 2020/04/14](https://openjdk.java.net/groups/vulnerability/advisories/2020-04-14)

以下に脆弱性の一覧を挙げておく。

{{< fig-quote title="OpenJDK Vulnerability Advisory: 2020/04/14" link="https://openjdk.java.net/groups/vulnerability/advisories/2020-04-14" lang="en" >}}
<table class="risk-matrix center smaller" summary="Risk matrix">
<tr>
<th rowspan="2">CVE ID</th>
<th rowspan="2">Component</th>
<th rowspan="2"><a href="https://www.first.org/cvss/">CVSSv3</a><br>Score</th>
<th colspan="5">Affects ...</th>
</tr>
<tr>
<th>7</th>
<th>8</th>
<th>11</th>
<th>13</th>
<th>14</th>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-2803">CVE-2020-2803</a></td>
<td style="text-align:left;"><code>core-libs/java.nio</code></td>
<td>8.3</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-2805">CVE-2020-2805</a></td>
<td style="text-align:left;"><code>core-libs/java.io</code></td>
<td>8.3</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-2816">CVE-2020-2816</a></td>
<td style="text-align:left;"><code>security-libs/javax.net.ssl</code></td>
<td>7.5</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-2781">CVE-2020-2781</a></td>
<td style="text-align:left;"><code>security-libs/java.security</code></td>
<td>5.3</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-2830">CVE-2020-2830</a></td>
<td style="text-align:left;"><code>core-libs/java.util</code></td>
<td>5.3</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-2767">CVE-2020-2767</a></td>
<td style="text-align:left;"><code>security-libs/javax.net.ssl</code></td>
<td>4.8</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-2800">CVE-2020-2800</a></td>
<td style="text-align:left;"><code>core-libs/java.net</code></td>
<td>4.8</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-2778">CVE-2020-2778</a></td>
<td style="text-align:left;"><code>security-libs/javax.net.ssl</code></td>
<td>3.7</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-2754">CVE-2020-2754</a></td>
<td style="text-align:left;"><code>core-libs/javax.script</code></td>
<td>3.7</td>
<td>&nbsp;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-2755">CVE-2020-2755</a></td>
<td style="text-align:left;"><code>core-libs/javax.script</code></td>
<td>3.7</td>
<td>&nbsp;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-2773">CVE-2020-2773</a></td>
<td style="text-align:left;"><code>security-libs/javax.xml.crypto</code></td>
<td>3.7</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-2756">CVE-2020-2756</a></td>
<td style="text-align:left;"><code>core-libs/java.io:serialization</code></td>
<td>3.7</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-2757">CVE-2020-2757</a></td>
<td style="text-align:left;"><code>core-libs/java.io:serialization</code></td>
<td>3.7</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&#8226;</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-18197">CVE-2020-18197</a></td>
<td style="text-align:left;"><code>javafx/web</code></td>
<td>8.1</td>
<td>&nbsp;</td>
<td>&#8226;</td>
<td>&#8226;</td>
<td>&nbsp;</td>
<td>&#8226;</td>
</tr>

</table>
{{< /fig-quote >}}

Java 10 以下 および 12, 13 は基本的にサポート期間が切れてる。
Java 11 または 14 へアップグレードするか [Amazon Corretto](https://aws.amazon.com/jp/corretto/) のようなディストリビューションを利用すること。

[OpenJDK] を APT で管理するメリットはないので，実行バイナリを[リリースページ](https://jdk.java.net/14/)から直接ダウンロードして配置する。

```text
$ cd /usr/local/src
$ sudo curl "https://download.java.net/java/GA/jdk14.0.1/664493ef4a6946b186ff29eb326336a2/7/GPL/openjdk-14.0.1_linux-x64_bin.tar.gz" -O
$ cd ..
$ sudo unlink java # 以前のバージョンの Java 環境がある場合
$ sudo tar xvf src/openjdk-14.0.1_linux-x64_bin.tar.gz
$ sudo ln -s jdk-14.0.1 java
$ java -version # すでに PATH が通っている場合
openjdk 14.0.1 2020-04-14
OpenJDK Runtime Environment (build 14.0.1+7)
OpenJDK 64-Bit Server VM (build 14.0.1+7, mixed mode, sharing)
```

私としては [PlantUML] が動けばいいので，試しておく[^puml1]。

[^puml1]: 2020-04-05 に [PlantUML] V1.2020.6 が[リリース](http://plantuml.com/ja/changes)されている。

{{< fig-img src="./factory-method-pattern.png" title="Factory Method Pattern" link="./factory-method-pattern.puml" width="1190" >}}

うむうむ。
ちゃんと動くな。

## ブックマーク

- [Oracle Critical Patch Update Advisory - April 2020](https://www.oracle.com/security-alerts/cpuapr2020.html)
    - [Oracle Java の脆弱性対策について(CVE-2020-2803等)：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/ciadr/vul/20200415-jre.html)
    - [2020年4月 Oracle 製品のクリティカルパッチアップデートに関する注意喚起](https://www.jpcert.or.jp/at/2020/at200017.html)
- [Java 環境のリリースとサポートに関する覚え書き]({{< ref "/remark/2018/02/release-cycle-of-java-environment.md" >}})

[OpenJDK]: http://openjdk.java.net/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "B00I8AT1EU" %}} <!-- Java言語で学ぶリファクタリング入門 -->
{{% review-paapi "B00I8ATHGW" %}} <!-- 増補改訂版 Java言語で学ぶデザインパターン入門 -->
{{% review-paapi "B00I8AT1BS" %}} <!-- Java言語で学ぶデザインパターン入門 マルチスレッド編 -->
