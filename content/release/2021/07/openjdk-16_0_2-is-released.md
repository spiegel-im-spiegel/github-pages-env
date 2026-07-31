+++
title = "OpenJDK 16.0.2 のリリース【セキュリティ・アップデート】"
date =  "2021-07-22T11:39:03+09:00"
description = "CVE ID ベースで4つの脆弱性修正がある。 "
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "ubuntu", "plantuml", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

定例の Java マイナー・バージョンアップが行われた。

- [OpenJDK Vulnerability Advisory: 2021/07/20](https://openjdk.java.net/groups/vulnerability/advisories/2021-07-20)

CVE ID ベースで4つの脆弱性修正がある。
影響を受けるバージョンは 16.0.1, 15.0.2, 13.0.6, 11.0.10, 8u282, 7u291 およびそれ以前。

{{< fig-quote class="nobox" title="OpenJDK Vulnerability Advisory: 2020/10/20" link="https://openjdk.java.net/groups/vulnerability/advisories/2020-10-20" lang="en" >}}
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
<th>16</th>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2021-2388">CVE-2021-2388</a></td>
<td style="text-align:left;"><code>hotspot/compiler</code></td>
<td>7.3</td>
<td>&nbsp;</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2021-2369">CVE-2021-2369</a></td>
<td style="text-align:left;"><code>security-libs/java.security</code></td>
<td>4.3</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2021-2432">CVE-2021-2432</a></td>
<td style="text-align:left;"><code>core-libs/javax.naming</code></td>
<td>3.7</td>
<td>{{% icons "check" %}}</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2021-2341">CVE-2021-2341</a></td>
<td style="text-align:left;"><code>core-libs/java.net</code></td>
<td>3.1</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

</table>
{{< /fig-quote >}}

[OpenJDK] を APT で管理するメリットはないので，実行バイナリを[リリースページ](https://jdk.java.net/16/)から直接ダウンロードして配置する。

```text
$ cd /usr/local/src
$ sudo curl "https://download.java.net/java/GA/jdk16.0.2/d4a915d82b4c4fbb9bde534da945d746/7/GPL/openjdk-16.0.2_linux-x64_bin.tar.gz" -O
$ cd ..
$ sudo unlink java # 以前のバージョンの Java 環境がある場合
$ sudo tar xvf src/openjdk-16.0.2_linux-x64_bin.tar.gz
$ sudo ln -s jdk-16.0.2 java
$ java -version # すでに PATH が通っている場合
openjdk version "16.0.2" 2021-07-20
OpenJDK Runtime Environment (build 16.0.2+7-67)
OpenJDK 64-Bit Server VM (build 16.0.2+7-67, mixed mode, sharing)
```

私としては [PlantUML] が動けばいいので，試しておく[^puml1]。

[^puml1]: 2021-06-26 に [PlantUML] V1.2021.8 が[リリース](http://plantuml.com/ja/changes)されている。 [PlantUML] の使い方等については拙文「[真面目に PlantUML]({{< ref "/remark/2018/12/plantuml-1.md" >}})」シリーズを参考にどうぞ。

{{< fig-img src="./factory-method-pattern.png" title="Factory Method Pattern" link="./factory-method-pattern.puml" width="1160" >}}

うむうむ。
ちゃんと動くな。

アップデートは計画的に。

## ブックマーク

- [Oracle Critical Patch Update Advisory - July 2021](https://www.oracle.com/security-alerts/cpujul2021.html)
- [2021年7月Oracle製品のクリティカルパッチアップデートに関する注意喚起](https://www.jpcert.or.jp/at/2021/at210032.html)

[OpenJDK]: http://openjdk.java.net/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "4621303252" %}} <!-- Effective Java 第3版 -->
