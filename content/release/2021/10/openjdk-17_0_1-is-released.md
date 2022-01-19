+++
title = "OpenJDK 17.0.1 のリリース【セキュリティ・アップデート】"
date =  "2021-10-20T19:36:55+09:00"
description = "アップデートは計画的に。"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "ubuntu", "plantuml", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

定例の Java マイナー・バージョンアップが行われた。

- [OpenJDK Vulnerability Advisory: 2021/10/19](https://openjdk.java.net/groups/vulnerability/advisories/2021-10-19)

CVE ID ベースで4つの脆弱性修正がある。
影響を受けるバージョンは 17.0.0, 16.0.2, 15.0.4, 13.0.8, 11.0.12, 8u302, 7u311 およびそれ以前。

{{< fig-quote class="nobox" title="OpenJDK Vulnerability Advisory: 2021/10/19" link="https://openjdk.java.net/groups/vulnerability/advisories/2021-10-19" lang="en" >}}
<table class="risk-matrix center smaller" summary="Risk matrix">
<tr>
<th rowspan="2">CVE ID</th>
<th rowspan="2">Component</th>
<th rowspan="2"><a href="https://www.first.org/cvss/">CVSSv3.1</a><br>Score</th>
<th colspan="6">Affects ...</th>
</tr>
<tr>
<th>7</th>
<th>8</th>
<th>11</th>
<th>13</th>
<th>15</th>
<th>17</th>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2021-35567">CVE-2021-35567</a></td>
<td style="text-align:left;"><code>security-libs/java.security</code></td>
<td>6.8</td>
<td>&nbsp;</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2021-35550">CVE-2021-35550</a></td>
<td style="text-align:left;"><code>security-libs/javax.net.ssl</code></td>
<td>5.9</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>&nbsp;</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2021-35586">CVE-2021-35586</a></td>
<td style="text-align:left;"><code>client-libs/javax.imageio</code></td>
<td>5.3</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2021-35564">CVE-2021-35564</a></td>
<td style="text-align:left;"><code>security-libs/java.security</code></td>
<td>5.3</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2021-35561">CVE-2021-35561</a></td>
<td style="text-align:left;"><code>core-libs/java.util</code></td>
<td>5.3</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2021-35565">CVE-2021-35565</a></td>
<td style="text-align:left;"><code>core-libs/java.net</code></td>
<td>5.3</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>&nbsp;</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2021-35559">CVE-2021-35559</a></td>
<td style="text-align:left;"><code>client-libs/javax.swing</code></td>
<td>5.3</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2021-35578">CVE-2021-35578</a></td>
<td style="text-align:left;"><code>security-libs/javax.net.ssl</code></td>
<td>5.3</td>
<td>&nbsp;</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2021-35556">CVE-2021-35556</a></td>
<td style="text-align:left;"><code>client-libs/javax.swing</code></td>
<td>5.3</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2021-35603">CVE-2021-35603</a></td>
<td style="text-align:left;"><code>security-libs/javax.net.ssl</code></td>
<td>3.7</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2021-35588">CVE-2021-35588</a></td>
<td style="text-align:left;"><code>hotspot/runtime</code></td>
<td>3.1</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2021-3517">CVE-2021-3517</a></td>
<td style="text-align:left;"><code>javafx/web</code></td>
<td>8.6</td>
<td>&nbsp;</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2021-3522">CVE-2021-3522</a></td>
<td style="text-align:left;"><code>javafx/web</code></td>
<td>5.5</td>
<td>&nbsp;</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>{{% icons "check" %}}</td>
</tr>

</table>
{{< /fig-quote >}}

[OpenJDK] を APT で管理するメリットはないので，実行バイナリを[リリースページ](https://jdk.java.net/17/)から直接ダウンロードして配置する。
以下は完全手動でのインストール（笑）

```text
$ cd /usr/local/src
$ sudo curl -L "https://download.java.net/java/GA/jdk17.0.1/2a2082e5a09d4267845be086888add4f/12/GPL/openjdk-17.0.1_linux-x64_bin.tar.gz" -O
$ cd ..
$ sudo unlink java # 以前のバージョンの Java 環境がある場合
$ sudo tar xvf src/openjdk-17.0.1_linux-x64_bin.tar.gz
$ sudo ln -s jdk-17.0.1 java
$ java -version # すでに PATH が通っている場合
openjdk version "17.0.1" 2021-10-19
OpenJDK Runtime Environment (build 17.0.1+12-39)
OpenJDK 64-Bit Server VM (build 17.0.1+12-39, mixed mode, sharing)
```

私としては [PlantUML] が動けばいいので，試しておく[^puml1]。

[^puml1]: 2021-10-05 に [PlantUML] V1.2021.12 が[リリース](http://plantuml.com/changes)されている。 [PlantUML] の使い方等については拙文「[真面目に PlantUML]({{< ref "/remark/2018/12/plantuml-1.md" >}})」シリーズを参考にどうぞ。

{{< fig-img src="./factory-method-pattern.png" title="Factory Method Pattern" link="./factory-method-pattern.puml" width="1160" >}}

うむうむ。
ちゃんと動くな。

アップデートは計画的に。

## ブックマーク

- [Oracle Critical Patch Update Advisory - October 2021](https://www.oracle.com/security-alerts/cpuoct2021.html)
- [Oracle Java の脆弱性対策について(CVE-2021-3517等)：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/ciadr/vul/20211020-jre.html)
- [2021年10月Oracle製品のクリティカルパッチアップデートに関する注意喚起](https://www.jpcert.or.jp/at/2021/at210046.html)

[OpenJDK]: http://openjdk.java.net/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "B07CKHR8C1" %}} <!-- Spring Data JPAプログラミング入門 -->
{{% review-paapi "4621303252" %}} <!-- Effective Java 第3版 -->
{{% review-paapi "B0893LQ5KY" %}} <!-- Spring Boot 2 入門 -->
