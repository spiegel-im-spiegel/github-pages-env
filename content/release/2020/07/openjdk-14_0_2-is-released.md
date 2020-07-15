+++
title = "OpenJDK 14.0.2 のリリース【セキュリティ・アップデート】"
date =  "2020-07-15T11:58:14+09:00"
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

- [OpenJDK Vulnerability Advisory: 2020/07/14](https://openjdk.java.net/groups/vulnerability/advisories/2020-07-14)

以下に脆弱性の一覧を挙げておく。
今回は脆弱性が少なめでよかったね（笑）

{{< fig-gen title="via “OpenJDK Vulnerability Advisory: 2020/07/14”" link="https://openjdk.java.net/groups/vulnerability/advisories/2020-07-14" lang="en" >}}
<table class="risk-matrix center smaller" summary="Risk matrix">
<tr>
<th rowspan="2">CVE ID</th>
<th rowspan="2">Component</th>
<th rowspan="2"><a href="https://www.first.org/cvss/">CVSSv3.1</a><br>Score</th>
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
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-14583">CVE-2020-14583</a></td>
<td style="text-align:left;"><code>core-libs/java.io</code></td>
<td>8.3</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-14593">CVE-2020-14593</a></td>
<td style="text-align:left;"><code>client-libs/2d</code></td>
<td>7.4</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-14562">CVE-2020-14562</a></td>
<td style="text-align:left;"><code>client-libs/javax.imageio</code></td>
<td>5.3</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-14621">CVE-2020-14621</a></td>
<td style="text-align:left;"><code>xml/jaxp</code></td>
<td>5.3</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-14556">CVE-2020-14556</a></td>
<td style="text-align:left;"><code>core-libs/java.util.concurrent</code></td>
<td>4.8</td>
<td>&nbsp;</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-14573">CVE-2020-14573</a></td>
<td style="text-align:left;"><code>hotspot/compiler</code></td>
<td>3.7</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-14578">CVE-2020-14578</a></td>
<td style="text-align:left;"><code>security-libs/java.security</code></td>
<td>3.7</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
</tr>

<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-14579">CVE-2020-14579</a></td>
<td style="text-align:left;"><code>security-libs/java.security</code></td>
<td>3.7</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
</tr>

<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-14581">CVE-2020-14581</a></td>
<td style="text-align:left;"><code>client-libs/2d</code></td>
<td>3.7</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-14577">CVE-2020-14577</a></td>
<td style="text-align:left;"><code>security-libs/javax.net.ssl</code></td>
<td>3.7</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2020-14664">CVE-2020-14664</a></td>
<td style="text-align:left;"><code>javafx/graphics</code></td>
<td>8.3</td>
<td>&nbsp;</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

</table>
{{< /fig-gen >}}

Java 10 以下 および 12, 13 は基本的にサポート期間が切れている。
Java 11 または 14 へアップグレードするか [Amazon Corretto](https://aws.amazon.com/jp/corretto/) のようなディストリビューションを利用することを強く推奨する。

[OpenJDK] を APT で管理するメリットはないので，実行バイナリを[リリースページ](https://jdk.java.net/14/)から直接ダウンロードして配置する。

```text
$ cd /usr/local/src
$ sudo curl "https://download.java.net/java/GA/jdk14.0.2/205943a0976c4ed48cb16f1043c5c647/12/GPL/openjdk-14.0.2_linux-x64_bin.tar.gz" -O
$ cd ..
$ sudo unlink java # 以前のバージョンの Java 環境がある場合
$ sudo tar xvf src/openjdk-14.0.2_linux-x64_bin.tar.gz
$ sudo ln -s jdk-14.0.2 java
$ java -version # すでに PATH が通っている場合
openjdk version "14.0.2" 2020-07-14
OpenJDK Runtime Environment (build 14.0.2+12-46)
OpenJDK 64-Bit Server VM (build 14.0.2+12-46, mixed mode, sharing)
```

私としては [PlantUML] が動けばいいので，試しておく[^puml1]。

[^puml1]: 2020-06-28 に [PlantUML] Version 1.2020.15 が[リリース](http://plantuml.com/ja/changes)されている。 [PlantUML] の使い方等については拙文「[真面目に PlantUML]({{< ref "/remark/2018/12/plantuml-1.md" >}})」シリーズを参考にどうぞ。

{{< fig-img src="./factory-method-pattern.png" title="Factory Method Pattern" link="./factory-method-pattern.puml" width="1193" >}}

うむうむ。
ちゃんと動くな。

## ブックマーク

- [Oracle Critical Patch Update Advisory - July 2020](https://www.oracle.com/security-alerts/cpujul2020.html)
- [AdoptOpenJDKプロジェクトがEclipse Foundationへの合流を発表。合流後の新プロジェクト名は「Eclipse Adoptium」に － Publickey](https://www.publickey1.jp/blog/20/adoptopenjdkeclipse_foundationeclipse_adoptium.html)
- [Java 環境のリリースとサポートに関する覚え書き]({{< ref "/remark/2018/02/release-cycle-of-java-environment.md" >}})

[OpenJDK]: http://openjdk.java.net/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "B00I8AT1EU" %}} <!-- Java言語で学ぶリファクタリング入門 -->
{{% review-paapi "B00I8ATHGW" %}} <!-- 増補改訂版 Java言語で学ぶデザインパターン入門 -->
{{% review-paapi "B00I8AT1BS" %}} <!-- Java言語で学ぶデザインパターン入門 マルチスレッド編 -->
