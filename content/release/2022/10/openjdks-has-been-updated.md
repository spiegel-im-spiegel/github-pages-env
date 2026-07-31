+++
title = "OpenJDK のセキュリティ・アップデート（2022-10）"
date =  "2022-11-04T13:56:55+09:00"
description = "影響を受けるバージョンは 19, 17.0.4, 15.0.8, 13.0.12, 11.0.16, 8u342, 7u351 およびそれ以前。"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "ubuntu", "plantuml", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

サボり気味ですみません {{% emoji "ペコン" %}}

定例の Java マイナー・バージョンアップが行われた。

- [OpenJDK Vulnerability Advisory: 2022/10/18](https://openjdk.org/groups/vulnerability/advisories/2022-10-18)

CVE ID ベースで6個の脆弱性修正がある。
影響を受けるバージョンは 19, 17.0.4, 15.0.8, 13.0.12, 11.0.16, 8u342, 7u351 およびそれ以前。

{{< fig-quote class="nobox" title="OpenJDK Vulnerability Advisory: 2022/10/18" link="https://openjdk.org/groups/vulnerability/advisories/2022-10-18" lang="en" >}}
<table class="risk-matrix center smaller" summary="Risk matrix">
<tr>
<th rowspan="2">CVE ID</th>
<th rowspan="2">Component</th>
<th rowspan="2"><a href="https://www.first.org/cvss/">CVSSv3.1</a><br>Score</th>
<th colspan="7">Affects ...</th>
</tr>
<tr>
<th>7</th>
<th>8</th>
<th>11</th>
<th>13</th>
<th>15</th>
<th>17</th>
<th>19</th>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2022-21626">CVE-2022-21626</a></td>
<td style="text-align:left;"><code>security-libs/<br>java.security</code></td>
<td>5.3</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td></td>
<td></td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2022-21618">CVE-2022-21618</a></td>
<td style="text-align:left;"><code>security-libs/<br>org.ietf.jgss</code></td>
<td>5.3</td>
<td></td>
<td></td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2022-21628">CVE-2022-21628</a></td>
<td style="text-align:left;"><code>core-libs/<br>java.net</code></td>
<td>5.3</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2022-39399">CVE-2022-39399</a></td>
<td style="text-align:left;"><code>core-libs/<br>java.net</code></td>
<td>3.7</td>
<td></td>
<td></td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2022-21619">CVE-2022-21619</a></td>
<td style="text-align:left;"><code>security-libs/<br>java.security</code></td>
<td>3.7</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2022-21624">CVE-2022-21624</a></td>
<td style="text-align:left;"><code>core-libs/<br>javax.naming</code></td>
<td>3.7</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

</table>
{{< /fig-quote >}}


[OpenJDK] を APT で管理するメリットはないので，実行バイナリを[リリースページ](https://jdk.java.net/19/)から直接ダウンロードして配置する。
以下は完全手動でのインストール（笑）

```text
$ cd /usr/local/src
$ sudo curl -L "https://download.java.net/java/GA/jdk19.0.1/afdd2e245b014143b62ccb916125e3ce/10/GPL/openjdk-19.0.1_linux-x64_bin.tar.gz" -O
$ cd ..
$ sudo unlink java # 以前のバージョンの Java 環境がある場合
$ sudo tar xvf src/openjdk-19.0.1_linux-x64_bin.tar.gz
$ sudo ln -s jdk-19.0.1 java
$ java -version # すでに PATH が通っている場合
openjdk version "19.0.1" 2022-10-18
OpenJDK Runtime Environment (build 19.0.1+10-21)
OpenJDK 64-Bit Server VM (build 19.0.1+10-21, mixed mode, sharing)
```

LTS 版 Java バイナリが欲しいなら [Adoptium](https://adoptium.net/) で取得できる。

私としては [PlantUML] が動けばいいので，試しておく[^puml1]。

[^puml1]: 2021-10-23 に [PlantUML] V1.2022.12 が[リリース](http://plantuml.com/changes)されている。 [PlantUML] の使い方等については拙文「[真面目に PlantUML]({{< ref "/remark/2018/12/plantuml-1.md" >}})」シリーズを参考にどうぞ。

{{< fig-img src="./factory-method-pattern.png" title="Factory Method Pattern" link="./factory-method-pattern.puml" width="1167" >}}

よーし，うむうむ，よーし。

## ブックマーク

- [2022年10月Oracle製品のクリティカルパッチアップデートに関する注意喚起](https://www.jpcert.or.jp/at/2022/at220029.html)

[OpenJDK]: http://openjdk.java.net/
[Adoptium]: https://adoptium.net/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "B07CKHR8C1" %}} <!-- Spring Data JPAプログラミング入門 -->
{{% review-paapi "4621303252" %}} <!-- Effective Java 第3版 -->
{{% review-paapi "B0893LQ5KY" %}} <!-- Spring Boot 2 入門 -->
